package tarantool

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/ebirukov/tnt-containers/tarantool"
)

func TestConnectFailover(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	arConnTimeout := 200 * time.Millisecond

	pingInterval := 500 * time.Millisecond

	// включаем микросекунды в std логере
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// поднимает контейнер с экземпляром реплики БД
	os.Setenv("TARANTOOL_REPLICATION_SOURCE", "0.0.0.0:12345")
	replica1, err := tarantool.RunContainer(ctx, tarantool.WithTarantool15("tarantool/tarantool:1.5", time.Second), tarantool.WithEnv("TARANTOOL_REPLICATION"))
	require.NoError(t, err)
	// урл для соединения с БД
	rHost1, err := replica1.ServerHostPort(ctx)
	require.NoError(t, err)

	defer replica1.Terminate(ctx)

	// поднимает контейнер с мастером БД
	master, err := tarantool.RunContainer(ctx, tarantool.WithTarantool15("tarantool/tarantool:1.5", time.Second))
	require.NoError(t, err)
	// урл для соединения с БД
	mHost2, err := master.ServerHostPort(ctx)
	require.NoError(t, err)

	defer master.Terminate(ctx)

	replica3, err := tarantool.RunContainer(ctx, tarantool.WithTarantool15("tarantool/tarantool:1.5", time.Second), tarantool.WithEnv("TARANTOOL_REPLICATION"))
	require.NoError(t, err)
	// урл для соединения с БД
	rHost3, err := replica3.ServerHostPort(ctx)
	require.NoError(t, err)

	cfgName := "arcfg"
	// конфигурируем запуск ar с поднятыми репликами
	arConfig := NewTestConfigFromMap(map[string]interface{}{
		"arcfg/master":   strings.Join([]string{rHost3, "fakehost:11011"}, ","),
		"arcfg/replica":  strings.Join([]string{rHost1, mHost2, "fakehost:11012"}, ","),
		"arcfg/PoolSize": 1,
		"arcfg/Timeout":  arConnTimeout,
	})

	pinger := activerecord.NewPinger(activerecord.WithPingInterval(pingInterval))
	defer pinger.StopWatch()

	logger := activerecord.NewLogger()
	logger.SetLogLevel(activerecord.ErrorLoggerLevel)

	activerecord.ReinitActiveRecord(
		activerecord.WithConfig(arConfig),
		activerecord.WithLogger(logger),
		activerecord.WithConnectionPinger(pinger),
	)

	_, err = activerecord.AddClusterChecker(ctx, cfgName, octopus.ClusterConfigParams)
	require.NoError(t, err)

	// проверяем типы и состав узлов кластера после загрузки конфигурации
	instances := pinger.ObservedInstances(cfgName)
	// все инстансы из конфигурации (включая несуществующие fakehost)
	require.Len(t, instances, 5)

	availableMasters := availableInstances(instances, activerecord.ModeMaster)
	require.Len(t, availableMasters, 1)
	// после загрузки конфигурации mHost2 должен переместится в инстансы мастеров
	require.Equal(t, mHost2, availableMasters[0].Config.Addr)

	// rHost3 должен переместится в инстансы реплик к rHost1
	availableReplicas := availableInstances(instances, activerecord.ModeReplica)
	require.Len(t, availableReplicas, 2)

	qDuration := func(st time.Time) int64 {
		return time.Now().UnixMilli() - st.UnixMilli()
	}

	eg := &errgroup.Group{}
	// асинхронно запускаем серию параллельных запросов в узлы кластера (отлавливаем тормоза и гонки)
	for g := 0; g < 8; g++ {
		g := g
		eg.Go(func() error {
			for i := 0; i < 1000; i++ {
				st := time.Now()
				// чтобы не тротлить пул
				time.Sleep(800 * time.Microsecond)
				err = execute(ctx, cfgName)

				if err != nil {
					// подождем немного и попробуем сделать еще запрос
					time.Sleep(10 * time.Millisecond)
					st := time.Now()

					err = execute(ctx, cfgName)
					if qDuration(st) > 5 {
						log.Printf("'%d' long query %d, time: %d ms\n", g, i, qDuration(st))
					}
					if err != nil {
						log.Printf("'%d' err %d: %s", g, i, err)
					} else {
						if i%10 != 0 {
							continue
						}
						log.Printf("'%d' success after retry %d\n", g, i)
					}
				} else {
					if qDuration(st) > 5 {
						log.Printf("'%d' long query %d, time: %d ms\n", g, i, qDuration(st))
					}

					if i%10 != 0 {
						continue
					}

					//log.Printf("'%d' success %d\n", g, i)
				}
			}

			fmt.Printf("all '%d' request complete\n", g)

			return nil
		})
	}

	// останавливаем мастер ноду
	require.NoError(t, master.Stop(ctx))
	// подождем пока пингер актуализирует кластер после остановки ноды
	time.Sleep(pingInterval + 10*time.Millisecond)

	instances = pinger.ObservedInstances(cfgName)
	// проверяем что остановленный мастер пропал из доступных узлов
	masters := availableInstances(instances, activerecord.ModeMaster)
	require.Len(t, masters, 0)

	replicas := availableInstances(instances, activerecord.ModeReplica)
	// есть 2 доступные реплики
	require.Len(t, replicas, 2)

	// останавливаем одну реплику (но в конфигурации активрекорд она по прежнему присутствует)
	require.NoError(t, replica3.Stop(ctx))
	// подождем пока пингер актуализирует кластер после остановки ноды
	time.Sleep(pingInterval + 10*time.Millisecond)

	instances = pinger.ObservedInstances(cfgName)
	replicas = availableInstances(instances, activerecord.ModeReplica)
	// осталась одна доступная реплика
	require.Len(t, replicas, 1)
	require.Equal(t, rHost1, replicas[0].Config.Addr)

	require.NoError(t, master.Start(ctx))
	masterHost, err := master.ServerHostPort(ctx)
	require.NoError(t, err)

	// обновляем конфигурацию узлов кластера
	arConfig.UpdateFromMap(map[string]interface{}{
		"arcfg/master":  strings.Join([]string{masterHost}, ","),
		"arcfg/replica": strings.Join([]string{rHost1}, ","),
		"arcfg/Timeout": arConnTimeout,
	})

	// подождем пока пингер актуализирует кластер после остановки ноды
	time.Sleep(pingInterval + 10*time.Millisecond)

	// обновленная конфигурация состоит из 2 узлов
	instances = pinger.ObservedInstances(cfgName)
	require.Len(t, instances, 2)

	masters = availableInstances(instances, activerecord.ModeMaster)
	require.Len(t, masters, 1)
	require.Equal(t, masterHost, masters[0].Config.Addr)

	require.Len(t, availableInstances(instances, activerecord.ModeReplica), 1)
	replica := availableInstances(instances, activerecord.ModeReplica)
	require.Equal(t, rHost1, replica[0].Config.Addr)

	fmt.Println("wait for requests")
	// ожидаем завершения всех запросов
	eg.Wait()
}

func execute(ctx context.Context, path string) error {
	box, err := octopus.Box(ctx, 0, activerecord.ReplicaOrMasterInstanceType, path, nil)
	if err != nil {
		return err
	}
	_, err = octopus.CallLua(ctx, box, "box.dostring", "return box.info.status")
	if err != nil {
		return err
	}

	return nil
}

func availableInstances(instances []activerecord.ShardInstance, modeType activerecord.ServerModeType) []activerecord.ShardInstance {
	ret := make([]activerecord.ShardInstance, 0, len(instances))
	for _, instance := range instances {
		if !instance.Offline && instance.Config.Mode == modeType {
			ret = append(ret, instance)
		}
	}

	return ret
}

type TestConfig struct {
	cfg     sync.Map
	created time.Time
}

func NewTestConfigFromMap(cfg map[string]interface{}) *TestConfig {
	dc := &TestConfig{
		cfg:     sync.Map{},
		created: time.Now(),
	}

	dc.UpdateFromMap(cfg)

	return dc
}

func (dc *TestConfig) UpdateFromMap(cfg map[string]interface{}) {
	for k, v := range cfg {
		dc.cfg.Store(k, v)
	}

	dc.cfg.Store("last_update", time.Now())
}

func (dc *TestConfig) GetLastUpdateTime() time.Time {
	updated, ok := dc.cfg.Load("last_update")
	if !ok {
		panic("no last_update")
	}

	return updated.(time.Time)
}

func (dc *TestConfig) GetBool(ctx context.Context, confPath string, dfl ...bool) bool {
	if ret, ok := dc.GetBoolIfExists(ctx, confPath); ok {
		return ret
	}

	if len(dfl) != 0 {
		return dfl[0]
	}

	return false
}

func (dc *TestConfig) GetBoolIfExists(ctx context.Context, confPath string) (value bool, ok bool) {
	if param, ex := dc.cfg.Load(confPath); ex {
		if ret, ok := param.(bool); ok {
			return ret, true
		}
	}

	return false, false
}

func (dc *TestConfig) GetInt(ctx context.Context, confPath string, dfl ...int) int {
	if ret, ok := dc.GetIntIfExists(ctx, confPath); ok {
		return ret
	}

	if len(dfl) != 0 {
		return dfl[0]
	}

	return 0
}

func (dc *TestConfig) GetIntIfExists(ctx context.Context, confPath string) (int, bool) {
	if param, ex := dc.cfg.Load(confPath); ex {
		if ret, ok := param.(int); ok {
			return ret, true
		}
	}

	return 0, false
}

func (dc *TestConfig) GetDuration(ctx context.Context, confPath string, dfl ...time.Duration) time.Duration {
	if ret, ok := dc.GetDurationIfExists(ctx, confPath); ok {
		return ret
	}

	if len(dfl) != 0 {
		return dfl[0]
	}

	return 0
}

func (dc *TestConfig) GetDurationIfExists(ctx context.Context, confPath string) (time.Duration, bool) {
	if param, ex := dc.cfg.Load(confPath); ex {
		if ret, ok := param.(time.Duration); ok {
			return ret, true
		}
	}

	return 0, false
}

func (dc *TestConfig) GetString(ctx context.Context, confPath string, dfl ...string) string {
	if ret, ok := dc.GetStringIfExists(ctx, confPath); ok {
		return ret
	}

	if len(dfl) != 0 {
		return dfl[0]
	}

	return ""
}

func (dc *TestConfig) GetStringIfExists(ctx context.Context, confPath string) (string, bool) {
	if param, ex := dc.cfg.Load(confPath); ex {
		if ret, ok := param.(string); ok {
			return ret, true
		}
	}

	return "", false
}

func (dc *TestConfig) GetStrings(ctx context.Context, confPath string, dfl []string) []string {
	return []string{}
}

func (dc *TestConfig) GetStruct(ctx context.Context, confPath string, valuePtr interface{}) (bool, error) {
	return false, nil
}
