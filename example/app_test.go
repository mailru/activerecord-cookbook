package main

import (
	"context"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/mailru/activerecord-cookbook/example/model/repository/generated/category"
	"github.com/mailru/activerecord-cookbook/example/model/repository/generated/foo"
	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
	"gotest.tools/assert"

	"github.com/mailru/activerecord-cookbook/example/model/ds"
	repository "github.com/mailru/activerecord-cookbook/example/model/repository/generated"
	"github.com/mailru/activerecord-cookbook/example/model/repository/generated/reward"
	"github.com/mailru/activerecord-cookbook/example/testutil/fixture"
)

var octopusMockServer *octopus.MockServer

func TestMain(m *testing.M) {
	var err error
	octopusMockServer, err = octopus.InitMockServer(octopus.WithLogger(&octopus.DefaultLogger{DebugMeta: repository.NamespacePackages}))
	if err != nil {
		log.Printf("can't start server: %s", err)

		return
	}

	err = octopusMockServer.Start()
	if err != nil {
		log.Printf("can't start server: %s", err)

		return
	}

	defer func() {
		err = octopusMockServer.Stop()
		if err != nil {
			log.Printf("can't stop server: %s", err)
		}
	}()

	activerecord.InitActiveRecord(
		activerecord.WithConfig(activerecord.NewDefaultConfigFromMap(map[string]interface{}{
			"arcfg/master":   octopusMockServer.GetServerHostPort(),
			"arcfg/replica":  octopusMockServer.GetServerHostPort(),
			"arcfg.PoolSize": 1,
			"arcfg.Timeout":  time.Millisecond * 200,
		})),
	)
	activerecord.Logger().SetLogLevel(activerecord.DebugLoggerLevel)

	m.Run()
}

func Test_callNoParamsProc(t *testing.T) {
	ctx := context.Background()
	octopusMockServer.SetFixtures([]octopus.FixtureType{
		fixture.GetCategoryProcedureMocker().ByFixture(ctx),
	})
	res, err := category.Call(ctx)
	assert.NilError(t, err)

	assert.Equal(t, 100500, res.GetAll())
}

func Test_callProc(t *testing.T) {
	ctx := context.Background()

	params := foo.FooParams{
		SearchQuery: map[string]string{"name": "my name"},
	}

	octopusMockServer.SetFixtures([]octopus.FixtureType{
		fixture.GetFooProcedureMocker().ByFixtureParams(ctx, params),
	})

	res, err := foo.Call(ctx, params)
	assert.NilError(t, err)

	assert.Equal(t, 200, res.GetStatus())
	assert.Equal(t, 123, res.GetJsonRawData().ID)
	assert.Equal(t, strings.Join([]string{"bar", "foo"}, ","), strings.Join(res.GetJsonRawData().List, ","))
	assert.Equal(t, "efij", res.GetTraceID())

}

func Test_main(t *testing.T) {
	ctx := context.Background()

	rewardByCodeMocker := fixture.GetRewardByCodeMocker()

	octopusMockServer.SetFixtures([]octopus.FixtureType{rewardByCodeMocker.ByFixtureCode(ctx, "64G_android")})

	rew, err := reward.SelectByCode(ctx, "64G_android")
	if err != nil {
		t.Errorf("Can't select reward: %s", err)
	}

	ret, err := reward.SelectByCode(ctx, "64G_androi")
	if err == nil {
		t.Errorf("select wrong reward. Returned: %+v, expected error", ret)
	}

	octopusMockServer.SetFixtures([]octopus.FixtureType{})

	ret, err = reward.SelectByCode(ctx, "64G_android")
	if err == nil {
		t.Errorf("select reward after reset fixtures return: %+v, expected error", ret)
	}

	err = rew.SetPartner("part")
	if err != nil {
		t.Errorf("can't set partner field: %s", err)
	}

	err = rew.Update(ctx)
	if err == nil {
		t.Errorf("update reward unexpected error: %s", err)
	}

	err = rew.Delete(ctx)
	if err == nil {
		t.Errorf("delete reward unexpected error: %s", err)
	}

	description := "new description"
	newRew := reward.New(ctx)

	if err = newRew.SetCode("new_code"); err != nil {
		t.Errorf("can't set code: %s", err)
	}

	if err = newRew.SetDescription(&description); err != nil {
		t.Errorf("can't set description: %s", err)
	}

	if err = newRew.SetExtra(&ds.Extra{}); err != nil {
		t.Errorf("can't set extra: %s", err)
	}

	if err = newRew.SetFlags(map[string]interface{}{"a": "b"}); err != nil {
		t.Errorf("can't set flags: %s", err)
	}

	err = newRew.Insert(ctx)
	if err == nil {
		t.Errorf("update reward expected error: %s", err)
	}
}
