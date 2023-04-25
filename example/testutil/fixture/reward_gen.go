// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.5.3-7-g0d63e41 (Commit: 0d63e411)
package fixture

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"sync"

	"github.com/mailru/activerecord-cookbook/example/model/ds"
	"github.com/mailru/activerecord-cookbook/example/model/repository/generated/reward"
	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
)

var rewardOnce sync.Once
var rewardStore map[string]*reward.Reward

//go:embed data/reward.yaml
var rewardSource []byte

func initReward() {
	rewardOnce.Do(func() {
		fixtures := reward.UnmarshalFixtures(rewardSource)

		rewardStore = map[string]*reward.Reward{}
		for _, f := range fixtures {
			if _, ok := rewardStore[f.Primary()]; ok {
				log.Fatalf("reward  fixture with Code %v is duplicated", f.Primary())
			}
			rewardStore[f.Primary()] = f
		}
	})
}

func GetRewardByCode(Code string) *reward.Reward {
	initReward()

	res, ex := rewardStore[Code]
	if !ex {
		log.Fatalf("Reward  fixture with Code %v not found", Code)
	}

	ctx := activerecord.Logger().SetLoggerValueToContext(context.Background(), map[string]interface{}{"GetRewardByCode": Code, "FixtureStore": "rewardStore"})

	activerecord.Logger().Debug(ctx, reward.RewardList([]*reward.Reward{res}))

	return res
}

var rewardUpdateOnce sync.Once
var rewardUpdateStore map[string]*reward.Reward

//go:embed data/reward_update.yaml
var rewardUpdateSource []byte

func initUpdateReward() {
	rewardUpdateOnce.Do(func() {
		fixtures := reward.UnmarshalUpdateFixtures(rewardUpdateSource)

		rewardUpdateStore = map[string]*reward.Reward{}
		for _, f := range fixtures {
			if _, ok := rewardUpdateStore[f.Primary()]; ok {
				log.Fatalf("reward Update fixture with Code %v is duplicated", f.Primary())
			}
			rewardUpdateStore[f.Primary()] = f
		}
	})
}

func GetUpdateRewardByCode(Code string) *reward.Reward {
	initUpdateReward()

	res, ex := rewardUpdateStore[Code]
	if !ex {
		log.Fatalf("Reward Update fixture with Code %v not found", Code)
	}

	ctx := activerecord.Logger().SetLoggerValueToContext(context.Background(), map[string]interface{}{"GetUpdateRewardByCode": Code, "FixtureStore": "rewardUpdateStore"})

	activerecord.Logger().Debug(ctx, reward.RewardList([]*reward.Reward{res}))

	return res
}

var rewardInsertReplaceOnce sync.Once
var rewardInsertReplaceStore map[string]*reward.Reward

//go:embed data/reward_insert_replace.yaml
var rewardInsertReplaceSource []byte

func initInsertReplaceReward() {
	rewardInsertReplaceOnce.Do(func() {
		fixtures := reward.UnmarshalInsertReplaceFixtures(rewardInsertReplaceSource)

		rewardInsertReplaceStore = map[string]*reward.Reward{}
		for _, f := range fixtures {
			if _, ok := rewardInsertReplaceStore[f.Primary()]; ok {
				log.Fatalf("reward InsertReplace fixture with Code %v is duplicated", f.Primary())
			}
			rewardInsertReplaceStore[f.Primary()] = f
		}
	})
}

func GetInsertReplaceRewardByCode(Code string) *reward.Reward {
	initInsertReplaceReward()

	res, ex := rewardInsertReplaceStore[Code]
	if !ex {
		log.Fatalf("Reward InsertReplace fixture with Code %v not found", Code)
	}

	ctx := activerecord.Logger().SetLoggerValueToContext(context.Background(), map[string]interface{}{"GetInsertReplaceRewardByCode": Code, "FixtureStore": "rewardInsertReplaceStore"})

	activerecord.Logger().Debug(ctx, reward.RewardList([]*reward.Reward{res}))

	return res
}

func GetUpdateRewardFixtureByCode(ctx context.Context, Code string, trigger func(types []octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := GetUpdateRewardByCode(Code)

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateUpdateFixture(obj.MockUpdate(ctx), wrappedTrigger), promiseIsUsed
}

func GetInsertRewardFixtureByCode(ctx context.Context, Code string, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := GetInsertReplaceRewardByCode(Code)

	return GetInsertRewardFixtureByModel(ctx, obj, trigger)
}

func GetInsertRewardFixtureByModel(ctx context.Context, obj *reward.Reward, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	reqData := obj.MockInsert(ctx)

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateInsertOrReplaceFixture(obj, reqData, wrappedTrigger), promiseIsUsed
}

func GetReplaceRewardFixtureByCode(ctx context.Context, Code string, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := GetInsertReplaceRewardByCode(Code)

	return GetReplaceRewardFixtureByModel(ctx, obj, trigger)
}

func GetReplaceRewardFixtureByModel(ctx context.Context, obj *reward.Reward, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	reqData := obj.MockReplace(ctx)

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateInsertOrReplaceFixture(obj, reqData, wrappedTrigger), promiseIsUsed
}

func GetInsertOrReplaceRewardFixtureByCode(ctx context.Context, Code string, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := GetInsertReplaceRewardByCode(Code)

	return GetInsertOrReplaceRewardFixtureByModel(ctx, obj, trigger)
}

func GetInsertOrReplaceRewardFixtureByModel(ctx context.Context, obj *reward.Reward, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	reqData := obj.MockInsertOrReplace(ctx)

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateInsertOrReplaceFixture(obj, reqData, wrappedTrigger), promiseIsUsed
}

type RewardBuildableFixture struct {
	PrimaryKey string // Code
	updateOps  []reward.RewardUpdateFixtureOptions

	trigger func([]octopus.FixtureType) []octopus.FixtureType
}

func UpdateRewardFixture(Code string) RewardBuildableFixture {
	return RewardBuildableFixture{PrimaryKey: Code}
}

func (bf RewardBuildableFixture) WithUpdatedServices(val *ds.Services) RewardBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		reward.RewardUpdateFixtureOptions{Services: &reward.RewardServicesUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf RewardBuildableFixture) WithUpdatedPartner(val string) RewardBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		reward.RewardUpdateFixtureOptions{Partner: &reward.RewardPartnerUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf RewardBuildableFixture) WithUpdatedExtra(val *ds.Extra) RewardBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		reward.RewardUpdateFixtureOptions{Extra: &reward.RewardExtraUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf RewardBuildableFixture) WithUpdatedFlags(val map[string]interface{}) RewardBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		reward.RewardUpdateFixtureOptions{Flags: &reward.RewardFlagsUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf RewardBuildableFixture) WithUpdatedUnlocked(val ds.ServiceUnlocked) RewardBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		reward.RewardUpdateFixtureOptions{Unlocked: &reward.RewardUnlockedUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf RewardBuildableFixture) WithUpdatedDescription(val *string) RewardBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		reward.RewardUpdateFixtureOptions{Description: &reward.RewardDescriptionUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf RewardBuildableFixture) OnUpdate(trigger func([]octopus.FixtureType) []octopus.FixtureType) RewardBuildableFixture {
	bf.trigger = trigger

	return bf
}

func (bf RewardBuildableFixture) Build(ctx context.Context) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(bf.trigger)

	obj := reward.New(ctx)

	if err := obj.SetCode(bf.PrimaryKey); err != nil {
		log.Fatalf("SetCode error: %v", err)
	}

	obj.BaseField.UpdateOps = []octopus.Ops{}

	reward.SetFixtureUpdateOptions(obj, bf.updateOps)

	return octopus.CreateUpdateFixture(obj.MockUpdate(ctx), wrappedTrigger), promiseIsUsed
}

type RewardBySelectByCodeMocker struct {
}

func GetRewardByCodeMocker() RewardBySelectByCodeMocker {
	return RewardBySelectByCodeMocker{}
}

func (m RewardBySelectByCodeMocker) ByFixtureCode(ctx context.Context, Codes ...string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}
	logger := activerecord.Logger()

	var key *string

	for _, Code := range Codes {
		fix := GetRewardByCode(Code)
		mocks = append(mocks, fix)

		if key == nil {

			ikey := fix.GetCode()
			key = &ikey
		} else if *key != fix.GetCode() {
			logger.Fatal(ctx, "Non unique keys in fixture list")
		}

	}

	return m.ByKeysMocks(ctx, []string{*key}, mocks)
}

func (m RewardBySelectByCodeMocker) EmptyByKeys(ctx context.Context, keys ...string) octopus.FixtureType {
	return m.ByKeysMocks(ctx, keys, []octopus.MockEntities{})
}

func (m RewardBySelectByCodeMocker) ByFixturePKWithKeys(ctx context.Context, Codes []string, keys []string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}

	for _, Code := range Codes {
		fix := GetRewardByCode(Code)
		mocks = append(mocks, fix)
	}

	return m.ByKeysMocks(ctx, keys, mocks)
}

func (m RewardBySelectByCodeMocker) ByKeysMocks(ctx context.Context, keys []string, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateSelectFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return reward.New(ctx).MockSelectByCodesRequest(ctx, keys)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by EmptyByKeys: %s", err))
	}

	return oft
}

type RewardBySelectByPartnerMocker struct {
	limiter activerecord.SelectorLimiter
}

func GetRewardByPartnerMocker(limiter activerecord.SelectorLimiter) RewardBySelectByPartnerMocker {
	return RewardBySelectByPartnerMocker{limiter: limiter}
}

func (m RewardBySelectByPartnerMocker) ByFixtureCode(ctx context.Context, Codes ...string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}
	logger := activerecord.Logger()

	var key *string

	for _, Code := range Codes {
		fix := GetRewardByCode(Code)
		mocks = append(mocks, fix)

		if key == nil {

			ikey := fix.GetPartner()
			key = &ikey
		} else if *key != fix.GetPartner() {
			logger.Fatal(ctx, "Non unique keys in fixture list")
		}

	}

	return m.ByKeysMocks(ctx, []string{*key}, mocks)
}

func (m RewardBySelectByPartnerMocker) EmptyByKeys(ctx context.Context, keys ...string) octopus.FixtureType {
	return m.ByKeysMocks(ctx, keys, []octopus.MockEntities{})
}

func (m RewardBySelectByPartnerMocker) ByFixturePKWithKeys(ctx context.Context, Codes []string, keys []string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}

	for _, Code := range Codes {
		fix := GetRewardByCode(Code)
		mocks = append(mocks, fix)
	}

	return m.ByKeysMocks(ctx, keys, mocks)
}

func (m RewardBySelectByPartnerMocker) ByKeysMocks(ctx context.Context, keys []string, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateSelectFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return reward.New(ctx).MockSelectByPartnersRequest(ctx, keys, m.limiter)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by EmptyByKeys: %s", err))
	}

	return oft
}
