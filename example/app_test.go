package main

import (
	"context"
	"testing"
	"time"

	"github.com/mailru/activerecord-cookbook/example/model/repository/generated/foo"
	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
	"gotest.tools/assert"

	"github.com/mailru/activerecord-cookbook/example/model/ds"
	repository "github.com/mailru/activerecord-cookbook/example/model/repository/generated"
	"github.com/mailru/activerecord-cookbook/example/model/repository/generated/reward"
	"github.com/mailru/activerecord-cookbook/example/testutil/fixture"
)

func _Test_callProc(t *testing.T) {
	ctx := context.Background()

	activerecord.InitActiveRecord(
		activerecord.WithConfig(activerecord.NewDefaultConfigFromMap(map[string]interface{}{
			"arcfg/master":   "127.0.0.1:11011",
			"arcfg/replica":  "127.0.0.1:11011",
			"arcfg.PoolSize": 1,
			"arcfg.Timeout":  time.Millisecond * 200,
		})),
	)

	params := foo.FooParams{
		SearchQuery: "who are you",
		TraceID:     "",
	}

	res, err := foo.Call(ctx, params)
	assert.NilError(t, err)

	if res != nil {
		resultStatus := res.GetStatus()
		assert.Equal(t, "200", resultStatus)

		rawJson := res.GetJsonRawData()
		assert.Equal(t, `{"data": []}`, rawJson)

		reqID := res.GetTraceID()
		assert.Equal(t, "", reqID)
	}

}

func Test_main(t *testing.T) {
	ctx := context.Background()

	octopusMockServer, err := octopus.InitMockServer(octopus.WithLogger(&octopus.DefaultLogger{DebugMeta: repository.NamespacePackages}))
	if err != nil {
		t.Errorf("can't start server: %s", err)
	}

	err = octopusMockServer.Start()
	if err != nil {
		t.Errorf("can't start server: %s", err)
	}

	defer func() {
		err = octopusMockServer.Stop()
		if err != nil {
			t.Errorf("can't stop server: %s", err)
		}
	}()

	activerecord.InitActiveRecord(
		activerecord.WithConfig(activerecord.NewDefaultConfigFromMap(map[string]interface{}{
			"arcfg/master":   "127.0.0.1:11011",
			"arcfg/replica":  "127.0.0.1:11011",
			"arcfg.PoolSize": 1,
			"arcfg.Timeout":  time.Millisecond * 200,
		})),
	)

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
