// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.5.3-4-gd9702e9 (Commit: d9702e9f)
package fixture

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"sync"

	"github.com/mailru/activerecord-cookbook/example/model/repository/generated/arobj"
	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
)

var arobjOnce sync.Once
var arobjStore map[int32]*arobj.ArObj

//go:embed data/arobj.yaml
var arobjSource []byte

func initArObj() {
	arobjOnce.Do(func() {
		fixtures := arobj.UnmarshalFixtures(arobjSource)

		arobjStore = map[int32]*arobj.ArObj{}
		for _, f := range fixtures {
			if _, ok := arobjStore[f.Primary()]; ok {
				log.Fatalf("arobj  fixture with ID %v is duplicated", f.Primary())
			}
			arobjStore[f.Primary()] = f
		}
	})
}

func GetArObjByID(ID int32) *arobj.ArObj {
	initArObj()

	res, ex := arobjStore[ID]
	if !ex {
		log.Fatalf("ArObj  fixture with ID %v not found", ID)
	}

	ctx := activerecord.Logger().SetLoggerValueToContext(context.Background(), map[string]interface{}{"GetArObjByID": ID, "FixtureStore": "arobjStore"})

	activerecord.Logger().Debug(ctx, arobj.ArObjList([]*arobj.ArObj{res}))

	return res
}

var arobjUpdateOnce sync.Once
var arobjUpdateStore map[int32]*arobj.ArObj

//go:embed data/arobj_update.yaml
var arobjUpdateSource []byte

func initUpdateArObj() {
	arobjUpdateOnce.Do(func() {
		fixtures := arobj.UnmarshalUpdateFixtures(arobjUpdateSource)

		arobjUpdateStore = map[int32]*arobj.ArObj{}
		for _, f := range fixtures {
			if _, ok := arobjUpdateStore[f.Primary()]; ok {
				log.Fatalf("arobj Update fixture with ID %v is duplicated", f.Primary())
			}
			arobjUpdateStore[f.Primary()] = f
		}
	})
}

func GetUpdateArObjByID(ID int32) *arobj.ArObj {
	initUpdateArObj()

	res, ex := arobjUpdateStore[ID]
	if !ex {
		log.Fatalf("ArObj Update fixture with ID %v not found", ID)
	}

	ctx := activerecord.Logger().SetLoggerValueToContext(context.Background(), map[string]interface{}{"GetUpdateArObjByID": ID, "FixtureStore": "arobjUpdateStore"})

	activerecord.Logger().Debug(ctx, arobj.ArObjList([]*arobj.ArObj{res}))

	return res
}

var arobjInsertReplaceOnce sync.Once
var arobjInsertReplaceStore map[int32]*arobj.ArObj

//go:embed data/arobj_insert_replace.yaml
var arobjInsertReplaceSource []byte

func initInsertReplaceArObj() {
	arobjInsertReplaceOnce.Do(func() {
		fixtures := arobj.UnmarshalInsertReplaceFixtures(arobjInsertReplaceSource)

		arobjInsertReplaceStore = map[int32]*arobj.ArObj{}
		for _, f := range fixtures {
			if _, ok := arobjInsertReplaceStore[f.Primary()]; ok {
				log.Fatalf("arobj InsertReplace fixture with ID %v is duplicated", f.Primary())
			}
			arobjInsertReplaceStore[f.Primary()] = f
		}
	})
}

func GetInsertReplaceArObjByID(ID int32) *arobj.ArObj {
	initInsertReplaceArObj()

	res, ex := arobjInsertReplaceStore[ID]
	if !ex {
		log.Fatalf("ArObj InsertReplace fixture with ID %v not found", ID)
	}

	ctx := activerecord.Logger().SetLoggerValueToContext(context.Background(), map[string]interface{}{"GetInsertReplaceArObjByID": ID, "FixtureStore": "arobjInsertReplaceStore"})

	activerecord.Logger().Debug(ctx, arobj.ArObjList([]*arobj.ArObj{res}))

	return res
}

func GetUpdateArObjFixtureByID(ctx context.Context, ID int32, trigger func(types []octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := GetUpdateArObjByID(ID)

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateUpdateFixture(obj.MockUpdate(ctx), wrappedTrigger), promiseIsUsed
}

func GetInsertArObjFixtureByID(ctx context.Context, ID int32, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := GetInsertReplaceArObjByID(ID)

	return GetInsertArObjFixtureByModel(ctx, obj, trigger)
}

func GetInsertArObjFixtureByModel(ctx context.Context, obj *arobj.ArObj, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	reqData := obj.MockInsert(ctx)

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateInsertOrReplaceFixture(obj, reqData, wrappedTrigger), promiseIsUsed
}

func GetReplaceArObjFixtureByID(ctx context.Context, ID int32, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := GetInsertReplaceArObjByID(ID)

	return GetReplaceArObjFixtureByModel(ctx, obj, trigger)
}

func GetReplaceArObjFixtureByModel(ctx context.Context, obj *arobj.ArObj, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	reqData := obj.MockReplace(ctx)

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateInsertOrReplaceFixture(obj, reqData, wrappedTrigger), promiseIsUsed
}

func GetInsertOrReplaceArObjFixtureByID(ctx context.Context, ID int32, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := GetInsertReplaceArObjByID(ID)

	return GetInsertOrReplaceArObjFixtureByModel(ctx, obj, trigger)
}

func GetInsertOrReplaceArObjFixtureByModel(ctx context.Context, obj *arobj.ArObj, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	reqData := obj.MockInsertOrReplace(ctx)

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateInsertOrReplaceFixture(obj, reqData, wrappedTrigger), promiseIsUsed
}

type ArObjBuildableFixture struct {
	PrimaryKey int32 // ID
	updateOps  []arobj.ArObjUpdateFixtureOptions

	trigger func([]octopus.FixtureType) []octopus.FixtureType
}

func UpdateArObjFixture(ID int32) ArObjBuildableFixture {
	return ArObjBuildableFixture{PrimaryKey: ID}
}

func (bf ArObjBuildableFixture) WithUpdatedName(val string) ArObjBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		arobj.ArObjUpdateFixtureOptions{Name: &arobj.ArObjNameUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf ArObjBuildableFixture) WithUpdatedAnotherID(val int32) ArObjBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		arobj.ArObjUpdateFixtureOptions{AnotherID: &arobj.ArObjAnotherIDUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf ArObjBuildableFixture) WithUpdatedType(val string) ArObjBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		arobj.ArObjUpdateFixtureOptions{Type: &arobj.ArObjTypeUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf ArObjBuildableFixture) WithUpdatedFlags(val uint32) ArObjBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		arobj.ArObjUpdateFixtureOptions{Flags: &arobj.ArObjFlagsUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf ArObjBuildableFixture) OnUpdate(trigger func([]octopus.FixtureType) []octopus.FixtureType) ArObjBuildableFixture {
	bf.trigger = trigger

	return bf
}

func (bf ArObjBuildableFixture) Build(ctx context.Context) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(bf.trigger)

	obj := arobj.New(ctx)

	if err := obj.SetID(bf.PrimaryKey); err != nil {
		log.Fatalf("SetID error: %v", err)
	}

	obj.BaseField.UpdateOps = []octopus.Ops{}

	arobj.SetFixtureUpdateOptions(obj, bf.updateOps)

	return octopus.CreateUpdateFixture(obj.MockUpdate(ctx), wrappedTrigger), promiseIsUsed
}

type ArObjBySelectByIDMocker struct {
}

func GetArObjByIDMocker() ArObjBySelectByIDMocker {
	return ArObjBySelectByIDMocker{}
}

func (m ArObjBySelectByIDMocker) ByFixtureID(ctx context.Context, IDs ...int32) octopus.FixtureType {
	mocks := []octopus.MockEntities{}
	logger := activerecord.Logger()

	var key *int32

	for _, ID := range IDs {
		fix := GetArObjByID(ID)
		mocks = append(mocks, fix)

		if key == nil {

			ikey := fix.GetID()
			key = &ikey
		} else if *key != fix.GetID() {
			logger.Fatal(ctx, "Non unique keys in fixture list")
		}

	}

	return m.ByKeysMocks(ctx, []int32{*key}, mocks)
}

func (m ArObjBySelectByIDMocker) EmptyByKeys(ctx context.Context, keys ...int32) octopus.FixtureType {
	return m.ByKeysMocks(ctx, keys, []octopus.MockEntities{})
}

func (m ArObjBySelectByIDMocker) ByFixturePKWithKeys(ctx context.Context, IDs []int32, keys []int32) octopus.FixtureType {
	mocks := []octopus.MockEntities{}

	for _, ID := range IDs {
		fix := GetArObjByID(ID)
		mocks = append(mocks, fix)
	}

	return m.ByKeysMocks(ctx, keys, mocks)
}

func (m ArObjBySelectByIDMocker) ByKeysMocks(ctx context.Context, keys []int32, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateSelectFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return arobj.New(ctx).MockSelectByIDsRequest(ctx, keys)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by EmptyByKeys: %s", err))
	}

	return oft
}

type ArObjBySelectByTypeMocker struct {
	limiter activerecord.SelectorLimiter
}

func GetArObjByTypeMocker(limiter activerecord.SelectorLimiter) ArObjBySelectByTypeMocker {
	return ArObjBySelectByTypeMocker{limiter: limiter}
}

func (m ArObjBySelectByTypeMocker) ByFixtureID(ctx context.Context, IDs ...int32) octopus.FixtureType {
	mocks := []octopus.MockEntities{}
	logger := activerecord.Logger()

	var key *string

	for _, ID := range IDs {
		fix := GetArObjByID(ID)
		mocks = append(mocks, fix)

		if key == nil {

			ikey := fix.GetType()
			key = &ikey
		} else if *key != fix.GetType() {
			logger.Fatal(ctx, "Non unique keys in fixture list")
		}

	}

	return m.ByKeysMocks(ctx, []string{*key}, mocks)
}

func (m ArObjBySelectByTypeMocker) EmptyByKeys(ctx context.Context, keys ...string) octopus.FixtureType {
	return m.ByKeysMocks(ctx, keys, []octopus.MockEntities{})
}

func (m ArObjBySelectByTypeMocker) ByFixturePKWithKeys(ctx context.Context, IDs []int32, keys []string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}

	for _, ID := range IDs {
		fix := GetArObjByID(ID)
		mocks = append(mocks, fix)
	}

	return m.ByKeysMocks(ctx, keys, mocks)
}

func (m ArObjBySelectByTypeMocker) ByKeysMocks(ctx context.Context, keys []string, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateSelectFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return arobj.New(ctx).MockSelectByTypesRequest(ctx, keys, m.limiter)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by EmptyByKeys: %s", err))
	}

	return oft
}

type ArObjBySelectByTypeIDMocker struct {
}

func GetArObjByTypeIDMocker() ArObjBySelectByTypeIDMocker {
	return ArObjBySelectByTypeIDMocker{}
}

func (m ArObjBySelectByTypeIDMocker) ByFixtureID(ctx context.Context, IDs ...int32) octopus.FixtureType {
	mocks := []octopus.MockEntities{}
	logger := activerecord.Logger()

	var key *arobj.TypeIDIndexType

	for _, ID := range IDs {
		fix := GetArObjByID(ID)
		mocks = append(mocks, fix)

		if key == nil {
			key = &arobj.TypeIDIndexType{

				Type: fix.GetType(),

				ID: fix.GetID(),
			}
		} else {

			if fix.GetType() != key.Type {
				logger.Fatal(ctx, "Non unique keys in fixture list")
			}

			if fix.GetID() != key.ID {
				logger.Fatal(ctx, "Non unique keys in fixture list")
			}

		}

	}

	return m.ByKeysMocks(ctx, []arobj.TypeIDIndexType{*key}, mocks)
}

func (m ArObjBySelectByTypeIDMocker) EmptyByKeys(ctx context.Context, keys ...arobj.TypeIDIndexType) octopus.FixtureType {
	return m.ByKeysMocks(ctx, keys, []octopus.MockEntities{})
}

func (m ArObjBySelectByTypeIDMocker) ByFixturePKWithKeys(ctx context.Context, IDs []int32, keys []arobj.TypeIDIndexType) octopus.FixtureType {
	mocks := []octopus.MockEntities{}

	for _, ID := range IDs {
		fix := GetArObjByID(ID)
		mocks = append(mocks, fix)
	}

	return m.ByKeysMocks(ctx, keys, mocks)
}

func (m ArObjBySelectByTypeIDMocker) ByKeysMocks(ctx context.Context, keys []arobj.TypeIDIndexType, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateSelectFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return arobj.New(ctx).MockSelectByTypeIDsRequest(ctx, keys)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by EmptyByKeys: %s", err))
	}

	return oft
}

type ArObjBySelectByTypePartMocker struct {
	limiter activerecord.SelectorLimiter
}

func GetArObjByTypePartMocker(limiter activerecord.SelectorLimiter) ArObjBySelectByTypePartMocker {
	return ArObjBySelectByTypePartMocker{limiter: limiter}
}

func (m ArObjBySelectByTypePartMocker) ByFixtureID(ctx context.Context, IDs ...int32) octopus.FixtureType {
	mocks := []octopus.MockEntities{}
	logger := activerecord.Logger()

	var key *string

	for _, ID := range IDs {
		fix := GetArObjByID(ID)
		mocks = append(mocks, fix)

		if key == nil {

			ikey := fix.GetType()
			key = &ikey
		} else if *key != fix.GetType() {
			logger.Fatal(ctx, "Non unique keys in fixture list")
		}

	}

	return m.ByKeysMocks(ctx, []string{*key}, mocks)
}

func (m ArObjBySelectByTypePartMocker) EmptyByKeys(ctx context.Context, keys ...string) octopus.FixtureType {
	return m.ByKeysMocks(ctx, keys, []octopus.MockEntities{})
}

func (m ArObjBySelectByTypePartMocker) ByFixturePKWithKeys(ctx context.Context, IDs []int32, keys []string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}

	for _, ID := range IDs {
		fix := GetArObjByID(ID)
		mocks = append(mocks, fix)
	}

	return m.ByKeysMocks(ctx, keys, mocks)
}

func (m ArObjBySelectByTypePartMocker) ByKeysMocks(ctx context.Context, keys []string, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateSelectFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return arobj.New(ctx).MockSelectByTypePartsRequest(ctx, keys, m.limiter)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by EmptyByKeys: %s", err))
	}

	return oft
}
