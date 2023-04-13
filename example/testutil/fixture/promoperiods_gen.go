// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.4.2 (Commit: b721ca40)
package fixture

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"sync"

	"github.com/mailru/activerecord-cookbook/example/model/dictionary"
	"github.com/mailru/activerecord-cookbook/example/model/repository/generated/promoperiods"
	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
)

var promoperiodsOnce sync.Once
var promoperiodsStore map[string]*promoperiods.Promoperiods

//go:embed data/promoperiods.yaml
var promoperiodsSource []byte

func initPromoperiods() {
	promoperiodsOnce.Do(func() {
		fixtures := promoperiods.UnmarshalFixtures(promoperiodsSource)

		promoperiodsStore = map[string]*promoperiods.Promoperiods{}
		for _, f := range fixtures {
			if _, ok := promoperiodsStore[f.Primary()]; ok {
				log.Fatalf("promoperiods  fixture with ID %v is duplicated", f.Primary())
			}
			promoperiodsStore[f.Primary()] = f
		}
	})
}

func GetPromoperiodsByID(ID string) *promoperiods.Promoperiods {
	initPromoperiods()

	res, ex := promoperiodsStore[ID]
	if !ex {
		log.Fatalf("Promoperiods  fixture with ID %v not found", ID)
	}

	ctx := activerecord.Logger().SetLoggerValueToContext(context.Background(), map[string]interface{}{"GetPromoperiodsByID": ID, "FixtureStore": "promoperiodsStore"})

	activerecord.Logger().Debug(ctx, promoperiods.PromoperiodsList([]*promoperiods.Promoperiods{res}))

	return res
}

var promoperiodsUpdateOnce sync.Once
var promoperiodsUpdateStore map[string]*promoperiods.Promoperiods

//go:embed data/promoperiods_update.yaml
var promoperiodsUpdateSource []byte

func initUpdatePromoperiods() {
	promoperiodsUpdateOnce.Do(func() {
		fixtures := promoperiods.UnmarshalUpdateFixtures(promoperiodsUpdateSource)

		promoperiodsUpdateStore = map[string]*promoperiods.Promoperiods{}
		for _, f := range fixtures {
			if _, ok := promoperiodsUpdateStore[f.Primary()]; ok {
				log.Fatalf("promoperiods Update fixture with ID %v is duplicated", f.Primary())
			}
			promoperiodsUpdateStore[f.Primary()] = f
		}
	})
}

func GetUpdatePromoperiodsByID(ID string) *promoperiods.Promoperiods {
	initUpdatePromoperiods()

	res, ex := promoperiodsUpdateStore[ID]
	if !ex {
		log.Fatalf("Promoperiods Update fixture with ID %v not found", ID)
	}

	ctx := activerecord.Logger().SetLoggerValueToContext(context.Background(), map[string]interface{}{"GetUpdatePromoperiodsByID": ID, "FixtureStore": "promoperiodsUpdateStore"})

	activerecord.Logger().Debug(ctx, promoperiods.PromoperiodsList([]*promoperiods.Promoperiods{res}))

	return res
}

var promoperiodsInsertReplaceOnce sync.Once
var promoperiodsInsertReplaceStore map[string]*promoperiods.Promoperiods

//go:embed data/promoperiods_insert_replace.yaml
var promoperiodsInsertReplaceSource []byte

func initInsertReplacePromoperiods() {
	promoperiodsInsertReplaceOnce.Do(func() {
		fixtures := promoperiods.UnmarshalInsertReplaceFixtures(promoperiodsInsertReplaceSource)

		promoperiodsInsertReplaceStore = map[string]*promoperiods.Promoperiods{}
		for _, f := range fixtures {
			if _, ok := promoperiodsInsertReplaceStore[f.Primary()]; ok {
				log.Fatalf("promoperiods InsertReplace fixture with ID %v is duplicated", f.Primary())
			}
			promoperiodsInsertReplaceStore[f.Primary()] = f
		}
	})
}

func GetInsertReplacePromoperiodsByID(ID string) *promoperiods.Promoperiods {
	initInsertReplacePromoperiods()

	res, ex := promoperiodsInsertReplaceStore[ID]
	if !ex {
		log.Fatalf("Promoperiods InsertReplace fixture with ID %v not found", ID)
	}

	ctx := activerecord.Logger().SetLoggerValueToContext(context.Background(), map[string]interface{}{"GetInsertReplacePromoperiodsByID": ID, "FixtureStore": "promoperiodsInsertReplaceStore"})

	activerecord.Logger().Debug(ctx, promoperiods.PromoperiodsList([]*promoperiods.Promoperiods{res}))

	return res
}

func GetUpdatePromoperiodsFixtureByID(ctx context.Context, ID string, trigger func(types []octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := GetUpdatePromoperiodsByID(ID)

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateUpdateFixture(obj.MockUpdate(ctx), wrappedTrigger), promiseIsUsed
}

func GetInsertPromoperiodsFixtureByID(ctx context.Context, ID string, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := GetInsertReplacePromoperiodsByID(ID)

	return GetInsertPromoperiodsFixtureByModel(ctx, obj, trigger)
}

func GetInsertPromoperiodsFixtureByModel(ctx context.Context, obj *promoperiods.Promoperiods, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	reqData := obj.MockInsert(ctx)

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateInsertOrReplaceFixture(obj, reqData, wrappedTrigger), promiseIsUsed
}

func GetReplacePromoperiodsFixtureByID(ctx context.Context, ID string, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := GetInsertReplacePromoperiodsByID(ID)

	return GetReplacePromoperiodsFixtureByModel(ctx, obj, trigger)
}

func GetReplacePromoperiodsFixtureByModel(ctx context.Context, obj *promoperiods.Promoperiods, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	reqData := obj.MockReplace(ctx)

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateInsertOrReplaceFixture(obj, reqData, wrappedTrigger), promiseIsUsed
}

func GetInsertOrReplacePromoperiodsFixtureByID(ctx context.Context, ID string, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	obj := GetInsertReplacePromoperiodsByID(ID)

	return GetInsertOrReplacePromoperiodsFixtureByModel(ctx, obj, trigger)
}

func GetInsertOrReplacePromoperiodsFixtureByModel(ctx context.Context, obj *promoperiods.Promoperiods, trigger func([]octopus.FixtureType) []octopus.FixtureType) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	reqData := obj.MockInsertOrReplace(ctx)

	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(trigger)

	return octopus.CreateInsertOrReplaceFixture(obj, reqData, wrappedTrigger), promiseIsUsed
}

type PromoperiodsBuildableFixture struct {
	PrimaryKey string // ID
	updateOps  []promoperiods.PromoperiodsUpdateFixtureOptions

	trigger func([]octopus.FixtureType) []octopus.FixtureType
}

func UpdatePromoperiodsFixture(ID string) PromoperiodsBuildableFixture {
	return PromoperiodsBuildableFixture{PrimaryKey: ID}
}

func (bf PromoperiodsBuildableFixture) WithUpdatedCode(val string) PromoperiodsBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		promoperiods.PromoperiodsUpdateFixtureOptions{Code: &promoperiods.PromoperiodsCodeUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf PromoperiodsBuildableFixture) WithUpdatedEmail(val string) PromoperiodsBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		promoperiods.PromoperiodsUpdateFixtureOptions{Email: &promoperiods.PromoperiodsEmailUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf PromoperiodsBuildableFixture) WithUpdatedStart(val int32) PromoperiodsBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		promoperiods.PromoperiodsUpdateFixtureOptions{Start: &promoperiods.PromoperiodsStartUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf PromoperiodsBuildableFixture) WithUpdatedFinish(val *dictionary.Product) PromoperiodsBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		promoperiods.PromoperiodsUpdateFixtureOptions{Finish: &promoperiods.PromoperiodsFinishUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf PromoperiodsBuildableFixture) WithUpdatedAction(val string) PromoperiodsBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		promoperiods.PromoperiodsUpdateFixtureOptions{Action: &promoperiods.PromoperiodsActionUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf PromoperiodsBuildableFixture) WithUpdatedPlatform(val map[string]interface{}) PromoperiodsBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		promoperiods.PromoperiodsUpdateFixtureOptions{Platform: &promoperiods.PromoperiodsPlatformUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf PromoperiodsBuildableFixture) WithUpdatedPromobunch(val uint32) PromoperiodsBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		promoperiods.PromoperiodsUpdateFixtureOptions{Promobunch: &promoperiods.PromoperiodsPromobunchUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf PromoperiodsBuildableFixture) WithUpdatedPlatforms(val uint32) PromoperiodsBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		promoperiods.PromoperiodsUpdateFixtureOptions{Platforms: &promoperiods.PromoperiodsPlatformsUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf PromoperiodsBuildableFixture) WithUpdatedPlanID(val int32) PromoperiodsBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		promoperiods.PromoperiodsUpdateFixtureOptions{PlanID: &promoperiods.PromoperiodsPlanIDUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf PromoperiodsBuildableFixture) WithUpdatedPlanType(val string) PromoperiodsBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		promoperiods.PromoperiodsUpdateFixtureOptions{PlanType: &promoperiods.PromoperiodsPlanTypeUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf PromoperiodsBuildableFixture) WithUpdatedPrice(val float64) PromoperiodsBuildableFixture {
	bf.updateOps = append(
		bf.updateOps,
		promoperiods.PromoperiodsUpdateFixtureOptions{Price: &promoperiods.PromoperiodsPriceUpdateFixtureOption{Value: val}},
	)

	return bf
}

func (bf PromoperiodsBuildableFixture) OnUpdate(trigger func([]octopus.FixtureType) []octopus.FixtureType) PromoperiodsBuildableFixture {
	bf.trigger = trigger

	return bf
}

func (bf PromoperiodsBuildableFixture) Build(ctx context.Context) (fx octopus.FixtureType, promiseIsUsed func() bool) {
	wrappedTrigger, promiseIsUsed := octopus.WrapTriggerWithOnUsePromise(bf.trigger)

	obj := promoperiods.New(ctx)

	if err := obj.SetID(bf.PrimaryKey); err != nil {
		log.Fatalf("SetID error: %v", err)
	}

	obj.BaseField.UpdateOps = []octopus.Ops{}

	promoperiods.SetFixtureUpdateOptions(obj, bf.updateOps)

	return octopus.CreateUpdateFixture(obj.MockUpdate(ctx), wrappedTrigger), promiseIsUsed
}

type PromoperiodsBySelectByIDMocker struct {
}

func GetPromoperiodsByIDMocker() PromoperiodsBySelectByIDMocker {
	return PromoperiodsBySelectByIDMocker{}
}

func (m PromoperiodsBySelectByIDMocker) ByFixtureID(ctx context.Context, IDs ...string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}
	logger := activerecord.Logger()

	var key *string

	for _, ID := range IDs {
		fix := GetPromoperiodsByID(ID)
		mocks = append(mocks, fix)

		if key == nil {

			ikey := fix.GetID()
			key = &ikey
		} else if *key != fix.GetID() {
			logger.Fatal(ctx, "Non unique keys in fixture list")
		}

	}

	return m.ByKeysMocks(ctx, []string{*key}, mocks)
}

func (m PromoperiodsBySelectByIDMocker) EmptyByKeys(ctx context.Context, keys ...string) octopus.FixtureType {
	return m.ByKeysMocks(ctx, keys, []octopus.MockEntities{})
}

func (m PromoperiodsBySelectByIDMocker) ByFixturePKWithKeys(ctx context.Context, IDs []string, keys []string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}

	for _, ID := range IDs {
		fix := GetPromoperiodsByID(ID)
		mocks = append(mocks, fix)
	}

	return m.ByKeysMocks(ctx, keys, mocks)
}

func (m PromoperiodsBySelectByIDMocker) ByKeysMocks(ctx context.Context, keys []string, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateSelectFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return promoperiods.New(ctx).MockSelectByIDsRequest(ctx, keys)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by EmptyByKeys: %s", err))
	}

	return oft
}

type PromoperiodsBySelectByCodeMocker struct {
	limiter activerecord.SelectorLimiter
}

func GetPromoperiodsByCodeMocker(limiter activerecord.SelectorLimiter) PromoperiodsBySelectByCodeMocker {
	return PromoperiodsBySelectByCodeMocker{limiter: limiter}
}

func (m PromoperiodsBySelectByCodeMocker) ByFixtureID(ctx context.Context, IDs ...string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}
	logger := activerecord.Logger()

	var key *string

	for _, ID := range IDs {
		fix := GetPromoperiodsByID(ID)
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

func (m PromoperiodsBySelectByCodeMocker) EmptyByKeys(ctx context.Context, keys ...string) octopus.FixtureType {
	return m.ByKeysMocks(ctx, keys, []octopus.MockEntities{})
}

func (m PromoperiodsBySelectByCodeMocker) ByFixturePKWithKeys(ctx context.Context, IDs []string, keys []string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}

	for _, ID := range IDs {
		fix := GetPromoperiodsByID(ID)
		mocks = append(mocks, fix)
	}

	return m.ByKeysMocks(ctx, keys, mocks)
}

func (m PromoperiodsBySelectByCodeMocker) ByKeysMocks(ctx context.Context, keys []string, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateSelectFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return promoperiods.New(ctx).MockSelectByCodesRequest(ctx, keys, m.limiter)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by EmptyByKeys: %s", err))
	}

	return oft
}

type PromoperiodsBySelectByEmailMocker struct {
	limiter activerecord.SelectorLimiter
}

func GetPromoperiodsByEmailMocker(limiter activerecord.SelectorLimiter) PromoperiodsBySelectByEmailMocker {
	return PromoperiodsBySelectByEmailMocker{limiter: limiter}
}

func (m PromoperiodsBySelectByEmailMocker) ByFixtureID(ctx context.Context, IDs ...string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}
	logger := activerecord.Logger()

	var key *string

	for _, ID := range IDs {
		fix := GetPromoperiodsByID(ID)
		mocks = append(mocks, fix)

		if key == nil {

			ikey := fix.GetEmail()
			key = &ikey
		} else if *key != fix.GetEmail() {
			logger.Fatal(ctx, "Non unique keys in fixture list")
		}

	}

	return m.ByKeysMocks(ctx, []string{*key}, mocks)
}

func (m PromoperiodsBySelectByEmailMocker) EmptyByKeys(ctx context.Context, keys ...string) octopus.FixtureType {
	return m.ByKeysMocks(ctx, keys, []octopus.MockEntities{})
}

func (m PromoperiodsBySelectByEmailMocker) ByFixturePKWithKeys(ctx context.Context, IDs []string, keys []string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}

	for _, ID := range IDs {
		fix := GetPromoperiodsByID(ID)
		mocks = append(mocks, fix)
	}

	return m.ByKeysMocks(ctx, keys, mocks)
}

func (m PromoperiodsBySelectByEmailMocker) ByKeysMocks(ctx context.Context, keys []string, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateSelectFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return promoperiods.New(ctx).MockSelectByEmailsRequest(ctx, keys, m.limiter)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by EmptyByKeys: %s", err))
	}

	return oft
}

type PromoperiodsBySelectByPlanTypePriceMocker struct {
}

func GetPromoperiodsByPlanTypePriceMocker() PromoperiodsBySelectByPlanTypePriceMocker {
	return PromoperiodsBySelectByPlanTypePriceMocker{}
}

func (m PromoperiodsBySelectByPlanTypePriceMocker) ByFixtureID(ctx context.Context, IDs ...string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}
	logger := activerecord.Logger()

	var key *promoperiods.PlanTypePriceIndexType

	for _, ID := range IDs {
		fix := GetPromoperiodsByID(ID)
		mocks = append(mocks, fix)

		if key == nil {
			key = &promoperiods.PlanTypePriceIndexType{

				PlanType: fix.GetPlanType(),

				Price: fix.GetPrice(),
			}
		} else {

			if fix.GetPlanType() != key.PlanType {
				logger.Fatal(ctx, "Non unique keys in fixture list")
			}

			if fix.GetPrice() != key.Price {
				logger.Fatal(ctx, "Non unique keys in fixture list")
			}

		}

	}

	return m.ByKeysMocks(ctx, []promoperiods.PlanTypePriceIndexType{*key}, mocks)
}

func (m PromoperiodsBySelectByPlanTypePriceMocker) EmptyByKeys(ctx context.Context, keys ...promoperiods.PlanTypePriceIndexType) octopus.FixtureType {
	return m.ByKeysMocks(ctx, keys, []octopus.MockEntities{})
}

func (m PromoperiodsBySelectByPlanTypePriceMocker) ByFixturePKWithKeys(ctx context.Context, IDs []string, keys []promoperiods.PlanTypePriceIndexType) octopus.FixtureType {
	mocks := []octopus.MockEntities{}

	for _, ID := range IDs {
		fix := GetPromoperiodsByID(ID)
		mocks = append(mocks, fix)
	}

	return m.ByKeysMocks(ctx, keys, mocks)
}

func (m PromoperiodsBySelectByPlanTypePriceMocker) ByKeysMocks(ctx context.Context, keys []promoperiods.PlanTypePriceIndexType, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateSelectFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return promoperiods.New(ctx).MockSelectByPlanTypePricesRequest(ctx, keys)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by EmptyByKeys: %s", err))
	}

	return oft
}

type PromoperiodsBySelectByEmailCodeMocker struct {
}

func GetPromoperiodsByEmailCodeMocker() PromoperiodsBySelectByEmailCodeMocker {
	return PromoperiodsBySelectByEmailCodeMocker{}
}

func (m PromoperiodsBySelectByEmailCodeMocker) ByFixtureID(ctx context.Context, IDs ...string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}
	logger := activerecord.Logger()

	var key *promoperiods.EmailCodeIndexType

	for _, ID := range IDs {
		fix := GetPromoperiodsByID(ID)
		mocks = append(mocks, fix)

		if key == nil {
			key = &promoperiods.EmailCodeIndexType{

				Email: fix.GetEmail(),

				Code: fix.GetCode(),
			}
		} else {

			if fix.GetEmail() != key.Email {
				logger.Fatal(ctx, "Non unique keys in fixture list")
			}

			if fix.GetCode() != key.Code {
				logger.Fatal(ctx, "Non unique keys in fixture list")
			}

		}

	}

	return m.ByKeysMocks(ctx, []promoperiods.EmailCodeIndexType{*key}, mocks)
}

func (m PromoperiodsBySelectByEmailCodeMocker) EmptyByKeys(ctx context.Context, keys ...promoperiods.EmailCodeIndexType) octopus.FixtureType {
	return m.ByKeysMocks(ctx, keys, []octopus.MockEntities{})
}

func (m PromoperiodsBySelectByEmailCodeMocker) ByFixturePKWithKeys(ctx context.Context, IDs []string, keys []promoperiods.EmailCodeIndexType) octopus.FixtureType {
	mocks := []octopus.MockEntities{}

	for _, ID := range IDs {
		fix := GetPromoperiodsByID(ID)
		mocks = append(mocks, fix)
	}

	return m.ByKeysMocks(ctx, keys, mocks)
}

func (m PromoperiodsBySelectByEmailCodeMocker) ByKeysMocks(ctx context.Context, keys []promoperiods.EmailCodeIndexType, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateSelectFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return promoperiods.New(ctx).MockSelectByEmailCodesRequest(ctx, keys)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by EmptyByKeys: %s", err))
	}

	return oft
}

type PromoperiodsBySelectByEmailActionMocker struct {
}

func GetPromoperiodsByEmailActionMocker() PromoperiodsBySelectByEmailActionMocker {
	return PromoperiodsBySelectByEmailActionMocker{}
}

func (m PromoperiodsBySelectByEmailActionMocker) ByFixtureID(ctx context.Context, IDs ...string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}
	logger := activerecord.Logger()

	var key *promoperiods.EmailActionIndexType

	for _, ID := range IDs {
		fix := GetPromoperiodsByID(ID)
		mocks = append(mocks, fix)

		if key == nil {
			key = &promoperiods.EmailActionIndexType{

				Email: fix.GetEmail(),

				Action: fix.GetAction(),
			}
		} else {

			if fix.GetEmail() != key.Email {
				logger.Fatal(ctx, "Non unique keys in fixture list")
			}

			if fix.GetAction() != key.Action {
				logger.Fatal(ctx, "Non unique keys in fixture list")
			}

		}

	}

	return m.ByKeysMocks(ctx, []promoperiods.EmailActionIndexType{*key}, mocks)
}

func (m PromoperiodsBySelectByEmailActionMocker) EmptyByKeys(ctx context.Context, keys ...promoperiods.EmailActionIndexType) octopus.FixtureType {
	return m.ByKeysMocks(ctx, keys, []octopus.MockEntities{})
}

func (m PromoperiodsBySelectByEmailActionMocker) ByFixturePKWithKeys(ctx context.Context, IDs []string, keys []promoperiods.EmailActionIndexType) octopus.FixtureType {
	mocks := []octopus.MockEntities{}

	for _, ID := range IDs {
		fix := GetPromoperiodsByID(ID)
		mocks = append(mocks, fix)
	}

	return m.ByKeysMocks(ctx, keys, mocks)
}

func (m PromoperiodsBySelectByEmailActionMocker) ByKeysMocks(ctx context.Context, keys []promoperiods.EmailActionIndexType, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateSelectFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return promoperiods.New(ctx).MockSelectByEmailActionsRequest(ctx, keys)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by EmptyByKeys: %s", err))
	}

	return oft
}

type PromoperiodsBySelectByEmailPartMocker struct {
	limiter activerecord.SelectorLimiter
}

func GetPromoperiodsByEmailPartMocker(limiter activerecord.SelectorLimiter) PromoperiodsBySelectByEmailPartMocker {
	return PromoperiodsBySelectByEmailPartMocker{limiter: limiter}
}

func (m PromoperiodsBySelectByEmailPartMocker) ByFixtureID(ctx context.Context, IDs ...string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}
	logger := activerecord.Logger()

	var key *string

	for _, ID := range IDs {
		fix := GetPromoperiodsByID(ID)
		mocks = append(mocks, fix)

		if key == nil {

			ikey := fix.GetEmail()
			key = &ikey
		} else if *key != fix.GetEmail() {
			logger.Fatal(ctx, "Non unique keys in fixture list")
		}

	}

	return m.ByKeysMocks(ctx, []string{*key}, mocks)
}

func (m PromoperiodsBySelectByEmailPartMocker) EmptyByKeys(ctx context.Context, keys ...string) octopus.FixtureType {
	return m.ByKeysMocks(ctx, keys, []octopus.MockEntities{})
}

func (m PromoperiodsBySelectByEmailPartMocker) ByFixturePKWithKeys(ctx context.Context, IDs []string, keys []string) octopus.FixtureType {
	mocks := []octopus.MockEntities{}

	for _, ID := range IDs {
		fix := GetPromoperiodsByID(ID)
		mocks = append(mocks, fix)
	}

	return m.ByKeysMocks(ctx, keys, mocks)
}

func (m PromoperiodsBySelectByEmailPartMocker) ByKeysMocks(ctx context.Context, keys []string, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateSelectFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return promoperiods.New(ctx).MockSelectByEmailPartsRequest(ctx, keys, m.limiter)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by EmptyByKeys: %s", err))
	}

	return oft
}
