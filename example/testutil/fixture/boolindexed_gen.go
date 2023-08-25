// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.8.5-1-gaa389f8 (Commit: aa389f82)
package fixture

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"sync"

	"github.com/mailru/activerecord-cookbook/example/model/repository/generated/boolindexed"
	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
)

var boolindexedOnce sync.Once
var boolindexedStore map[string]*boolindexed.Boolindexed

//go:embed data/boolindexed.yaml
var boolindexedSource []byte

func initBoolindexed() {
	boolindexedOnce.Do(func() {
		fixtures := boolindexed.UnmarshalFixtures(boolindexedSource)

		boolindexedStore = map[string]*boolindexed.Boolindexed{}
		for _, f := range fixtures {
			if _, ok := boolindexedStore[f.Primary()]; ok {
				log.Fatalf("boolindexed  fixture with Code %v is duplicated", f.Primary())
			}
			boolindexedStore[f.Primary()] = f
		}
	})
}

func GetBoolindexedByCode(Code string) *boolindexed.Boolindexed {
	initBoolindexed()

	res, ex := boolindexedStore[Code]
	if !ex {
		log.Fatalf("Boolindexed  fixture with Code %v not found", Code)
	}

	ctx := activerecord.Logger().SetLoggerValueToContext(context.Background(), map[string]interface{}{"GetBoolindexedByCode": Code, "FixtureStore": "boolindexedStore"})

	activerecord.Logger().Debug(ctx, boolindexed.BoolindexedList([]*boolindexed.Boolindexed{res}))

	return res
}

var boolindexedUpdateOnce sync.Once
var boolindexedUpdateStore map[string]*boolindexed.Boolindexed

//go:embed data/boolindexed_update.yaml
var boolindexedUpdateSource []byte

func initUpdateBoolindexed() {
	boolindexedUpdateOnce.Do(func() {
		fixtures := boolindexed.UnmarshalUpdateFixtures(boolindexedUpdateSource)

		boolindexedUpdateStore = map[string]*boolindexed.Boolindexed{}
		for _, f := range fixtures {
			if _, ok := boolindexedUpdateStore[f.Primary()]; ok {
				log.Fatalf("boolindexed Update fixture with Code %v is duplicated", f.Primary())
			}
			boolindexedUpdateStore[f.Primary()] = f
		}
	})
}

func GetUpdateBoolindexedByCode(Code string) *boolindexed.Boolindexed {
	initUpdateBoolindexed()

	res, ex := boolindexedUpdateStore[Code]
	if !ex {
		log.Fatalf("Boolindexed Update fixture with Code %v not found", Code)
	}

	ctx := activerecord.Logger().SetLoggerValueToContext(context.Background(), map[string]interface{}{"GetUpdateBoolindexedByCode": Code, "FixtureStore": "boolindexedUpdateStore"})

	activerecord.Logger().Debug(ctx, boolindexed.BoolindexedList([]*boolindexed.Boolindexed{res}))

	return res
}

var boolindexedInsertReplaceOnce sync.Once
var boolindexedInsertReplaceStore map[string]*boolindexed.Boolindexed

//go:embed data/boolindexed_insert_replace.yaml
var boolindexedInsertReplaceSource []byte

func initInsertReplaceBoolindexed() {
	boolindexedInsertReplaceOnce.Do(func() {
		fixtures := boolindexed.UnmarshalInsertReplaceFixtures(boolindexedInsertReplaceSource)

		boolindexedInsertReplaceStore = map[string]*boolindexed.Boolindexed{}
		for _, f := range fixtures {
			if _, ok := boolindexedInsertReplaceStore[f.Primary()]; ok {
				log.Fatalf("boolindexed InsertReplace fixture with Code %v is duplicated", f.Primary())
			}
			boolindexedInsertReplaceStore[f.Primary()] = f
		}
	})
}

func GetInsertReplaceBoolindexedByCode(Code string) *boolindexed.Boolindexed {
	initInsertReplaceBoolindexed()

	res, ex := boolindexedInsertReplaceStore[Code]
	if !ex {
		log.Fatalf("Boolindexed InsertReplace fixture with Code %v not found", Code)
	}

	ctx := activerecord.Logger().SetLoggerValueToContext(context.Background(), map[string]interface{}{"GetInsertReplaceBoolindexedByCode": Code, "FixtureStore": "boolindexedInsertReplaceStore"})

	activerecord.Logger().Debug(ctx, boolindexed.BoolindexedList([]*boolindexed.Boolindexed{res}))

	return res
}

func GetDeleteBoolindexedFixtureByPrimaryKey(ctx context.Context, pk string, trigger func(types []octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := boolindexed.New(ctx)

	if err := obj.SetCode(pk); err != nil {
		log.Fatalf("SetCode error: %v", err)
	}

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateDeleteFixture(obj.MockDelete(ctx), wrappedTrigger), promiseIsUsed
}

func GetUpdateBoolindexedFixtureByCode(ctx context.Context, Code string, trigger func(types []octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := GetUpdateBoolindexedByCode(Code)

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateUpdateFixture(obj.MockUpdate(ctx), wrappedTrigger), promiseIsUsed
}

func GetInsertBoolindexedFixtureByCode(ctx context.Context, Code string, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := GetInsertReplaceBoolindexedByCode(Code)

	return GetInsertBoolindexedFixtureByModel(ctx, obj, trigger)
}

func GetInsertBoolindexedFixtureByModel(ctx context.Context, obj *boolindexed.Boolindexed, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	reqData := obj.MockInsert(ctx)

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateInsertOrReplaceFixture(obj, reqData, wrappedTrigger), promiseIsUsed
}

func GetReplaceBoolindexedFixtureByCode(ctx context.Context, Code string, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := GetInsertReplaceBoolindexedByCode(Code)

	return GetReplaceBoolindexedFixtureByModel(ctx, obj, trigger)
}

func GetReplaceBoolindexedFixtureByModel(ctx context.Context, obj *boolindexed.Boolindexed, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	reqData := obj.MockReplace(ctx)

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateInsertOrReplaceFixture(obj, reqData, wrappedTrigger), promiseIsUsed
}

func GetInsertOrReplaceBoolindexedFixtureByCode(ctx context.Context, Code string, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := GetInsertReplaceBoolindexedByCode(Code)

	return GetInsertOrReplaceBoolindexedFixtureByModel(ctx, obj, trigger)
}

func GetInsertOrReplaceBoolindexedFixtureByModel(ctx context.Context, obj *boolindexed.Boolindexed, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	reqData := obj.MockInsertOrReplace(ctx)

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateInsertOrReplaceFixture(obj, reqData, wrappedTrigger), promiseIsUsed
}

type BoolindexedBuildableFixture struct {
	PrimaryKey string // Code
	updateOps  []boolindexed.BoolindexedUpdateFixtureOptions

	trigger func([]octopus.FixtureType) []octopus.FixtureType
}

func UpdateBoolindexedFixture(Code string) BoolindexedBuildableFixture {
	return BoolindexedBuildableFixture{PrimaryKey: Code}
}

func (bf BoolindexedBuildableFixture) WithUpdatedInvisible(val bool) BoolindexedBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		boolindexed.BoolindexedUpdateFixtureOptions{Invisible: &boolindexed.BoolindexedInvisibleUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf BoolindexedBuildableFixture) OnUpdate(trigger func([]octopus.FixtureType) []octopus.FixtureType) BoolindexedBuildableFixture {
	bf.trigger = trigger

	return bf
}

func (bf BoolindexedBuildableFixture) Build(ctx context.Context) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(bf.trigger)

	obj := boolindexed.New(ctx)

	if err := obj.SetCode(bf.PrimaryKey); err != nil {
		log.Fatalf("SetCode error: %v", err)
	}

	obj.BaseField.UpdateOps = []octopus.Ops{}

	boolindexed.SetFixtureUpdateOptions(obj, bf.updateOps)

	return octopus.CreateUpdateFixture(obj.MockUpdate(ctx), wrappedTrigger), promiseIsUsed
}

type BoolindexedBySelectByCodeMocker struct {
}

func GetBoolindexedByCodeMocker() BoolindexedBySelectByCodeMocker {
	return BoolindexedBySelectByCodeMocker{}
}

func (m BoolindexedBySelectByCodeMocker) ByFixtureCode(ctx context.Context, Codes ...string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}
	logger := activerecord.Logger()

	var key *string

	for _, Code := range Codes {
		fix := GetBoolindexedByCode(Code)
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

func (m BoolindexedBySelectByCodeMocker) EmptyByKeys(ctx context.Context, keys ...string) octopus.FixtureType {
	return m.ByKeysMocks(ctx, keys, []octopus.MockEntities{})
}

func (m BoolindexedBySelectByCodeMocker) ByFixturePKWithKeys(ctx context.Context, Codes []string, keys []string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}

	for _, Code := range Codes {
		fix := GetBoolindexedByCode(Code)
		mocks = append(mocks, fix)
	}

	return m.ByKeysMocks(ctx, keys, mocks)
}

func (m BoolindexedBySelectByCodeMocker) ByKeysMocks(ctx context.Context, keys []string, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateSelectFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return boolindexed.New(ctx).MockSelectByCodesRequest(ctx, keys)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by EmptyByKeys: %s", err))
	}

	return oft
}

type BoolindexedBySelectByInvisibleMocker struct {
	limiter activerecord.SelectorLimiter
}

func GetBoolindexedByInvisibleMocker(limiter activerecord.SelectorLimiter) BoolindexedBySelectByInvisibleMocker {
	return BoolindexedBySelectByInvisibleMocker{limiter: limiter}
}

func (m BoolindexedBySelectByInvisibleMocker) ByFixtureCode(ctx context.Context, Codes ...string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}
	logger := activerecord.Logger()

	var key *bool

	for _, Code := range Codes {
		fix := GetBoolindexedByCode(Code)
		mocks = append(mocks, fix)

		if key == nil {

			ikey := fix.GetInvisible()
			key = &ikey
		} else if *key != fix.GetInvisible() {
			logger.Fatal(ctx, "Non unique keys in fixture list")
		}

	}

	return m.ByKeysMocks(ctx, []bool{*key}, mocks)
}

func (m BoolindexedBySelectByInvisibleMocker) EmptyByKeys(ctx context.Context, keys ...bool) octopus.FixtureType {
	return m.ByKeysMocks(ctx, keys, []octopus.MockEntities{})
}

func (m BoolindexedBySelectByInvisibleMocker) ByFixturePKWithKeys(ctx context.Context, Codes []string, keys []bool) octopus.FixtureType {
	mocks := []octopus.MockEntities{}

	for _, Code := range Codes {
		fix := GetBoolindexedByCode(Code)
		mocks = append(mocks, fix)
	}

	return m.ByKeysMocks(ctx, keys, mocks)
}

func (m BoolindexedBySelectByInvisibleMocker) ByKeysMocks(ctx context.Context, keys []bool, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateSelectFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return boolindexed.New(ctx).MockSelectByInvisiblesRequest(ctx, keys, m.limiter)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by EmptyByKeys: %s", err))
	}

	return oft
}
