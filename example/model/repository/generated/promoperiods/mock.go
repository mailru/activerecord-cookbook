// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.5.3-5-g90e9b6c (Commit: 90e9b6c3)
package promoperiods

import (
	"context"
	"fmt"

	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
	serializerJSON "github.com/mailru/activerecord/pkg/serializer"
)

func (obj *Promoperiods) MockSelectResponse() ([][]byte, error) {
	tuple := [][]byte{}

	var data []byte

	var err error

	data, err = packID([]byte{}, obj.GetID())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packCode([]byte{}, obj.GetCode())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packEmail([]byte{}, obj.GetEmail())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packStart([]byte{}, obj.GetStart())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packFinish([]byte{}, obj.GetFinish())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packAction([]byte{}, obj.GetAction())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packPlatform([]byte{}, obj.GetPlatform())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packPromobunch([]byte{}, obj.GetPromobunch())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packPlatforms([]byte{}, obj.GetPlatforms())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packPlanID([]byte{}, obj.GetPlanID())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packPlanType([]byte{}, obj.GetPlanType())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packPrice([]byte{}, obj.GetPrice())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	return tuple, nil
}

//indexes

func (obj *Promoperiods) MockSelectByIDRequest(ctx context.Context) []byte {
	key := []string{
		obj.GetID(),
	}
	return obj.MockSelectByIDsRequest(ctx, key)
}

func (obj *Promoperiods) MockSelectByIDsRequest(ctx context.Context, keys []string) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockSelectByIDsRequest": keys, "Repo": "Promoperiods"})
	keysPacked := [][][]byte{}

	var (
		data []byte
		err  error
	)
	for _, key := range keys {
		keysField := [][]byte{}

		data, err = packID([]byte{}, key)
		if err != nil {
			log.Fatal(ctx, err)
			return nil
		}

		keysField = append(keysField, data)
		keysPacked = append(keysPacked, keysField)
	}

	log.Debug(ctx, fmt.Sprintf("Packed Keys: '% X'", keysPacked))

	return octopus.PackSelect(namespace, 0, 0, 0, keysPacked)

}

func SelectByIDMockerLogger(keys []string, res PromoperiodsList) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {
		fixturesKey := ""
		for _, key := range keys {

			fixturesKey += "\"" + key + "\""

			fixturesKey += ",\n"
		}

		mockerName := "mockerPromoperiodsByID"
		mocker := "fixture.GetPromoperiodsByIDMocker("

		mocker += ")"

		fixture := mockerName

		if res != nil && len(res) != 0 {
			pks := "[]string{\n"

			for _, r := range res {

				pks += "\"" + r.Primary() + "\",\n"

			}

			pks += "}"

			fixture += ".ByFixturePKWithKeys(ctx, " + pks + ", []string{" + fixturesKey + "})"
		} else {
			fixture += ".EmptyByKeys(ctx, " + fixturesKey + ")"
		}

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "promoperiods", Results: res}, nil
	}
}

func (obj *Promoperiods) MockSelectByCodeRequest(ctx context.Context, limiter activerecord.SelectorLimiter) []byte {
	key := []string{
		obj.GetCode(),
	}
	return obj.MockSelectByCodesRequest(ctx, key, limiter)
}

func (obj *Promoperiods) MockSelectByCodesRequest(ctx context.Context, keys []string, limiter activerecord.SelectorLimiter) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockSelectByCodesRequest": keys, "Repo": "Promoperiods"})
	keysPacked := [][][]byte{}

	var (
		data []byte
		err  error
	)
	for _, key := range keys {
		keysField := [][]byte{}

		data, err = packCode([]byte{}, key)
		if err != nil {
			log.Fatal(ctx, err)
			return nil
		}

		keysField = append(keysField, data)
		keysPacked = append(keysPacked, keysField)
	}

	log.Debug(ctx, fmt.Sprintf("Packed Keys: '% X'", keysPacked))

	return octopus.PackSelect(namespace, 1, limiter.Offset(), limiter.Limit(), keysPacked)

}

func SelectByCodeMockerLogger(keys []string, res PromoperiodsList, limiter activerecord.SelectorLimiter) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {
		fixturesKey := ""
		for _, key := range keys {

			fixturesKey += "\"" + key + "\""

			fixturesKey += ",\n"
		}

		mockerName := "mockerPromoperiodsByCode"
		mocker := "fixture.GetPromoperiodsByCodeMocker("

		mocker += fmt.Sprintf("activerecord.NewLimitOffset(%d, %d)", limiter.Limit(), limiter.Offset())
		mockerName += fmt.Sprintf("_%d_%d", limiter.Limit(), limiter.Offset())

		mocker += ")"

		fixture := mockerName

		if res != nil && len(res) != 0 {
			pks := "[]string{\n"

			for _, r := range res {

				pks += "\"" + r.Primary() + "\",\n"

			}

			pks += "}"

			fixture += ".ByFixturePKWithKeys(ctx, " + pks + ", []string{" + fixturesKey + "})"
		} else {
			fixture += ".EmptyByKeys(ctx, " + fixturesKey + ")"
		}

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "promoperiods", Results: res}, nil
	}
}

func (obj *Promoperiods) MockSelectByEmailRequest(ctx context.Context, limiter activerecord.SelectorLimiter) []byte {
	key := []string{
		obj.GetEmail(),
	}
	return obj.MockSelectByEmailsRequest(ctx, key, limiter)
}

func (obj *Promoperiods) MockSelectByEmailsRequest(ctx context.Context, keys []string, limiter activerecord.SelectorLimiter) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockSelectByEmailsRequest": keys, "Repo": "Promoperiods"})
	keysPacked := [][][]byte{}

	var (
		data []byte
		err  error
	)
	for _, key := range keys {
		keysField := [][]byte{}

		data, err = packEmail([]byte{}, key)
		if err != nil {
			log.Fatal(ctx, err)
			return nil
		}

		keysField = append(keysField, data)
		keysPacked = append(keysPacked, keysField)
	}

	log.Debug(ctx, fmt.Sprintf("Packed Keys: '% X'", keysPacked))

	return octopus.PackSelect(namespace, 2, limiter.Offset(), limiter.Limit(), keysPacked)

}

func SelectByEmailMockerLogger(keys []string, res PromoperiodsList, limiter activerecord.SelectorLimiter) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {
		fixturesKey := ""
		for _, key := range keys {

			fixturesKey += "\"" + key + "\""

			fixturesKey += ",\n"
		}

		mockerName := "mockerPromoperiodsByEmail"
		mocker := "fixture.GetPromoperiodsByEmailMocker("

		mocker += fmt.Sprintf("activerecord.NewLimitOffset(%d, %d)", limiter.Limit(), limiter.Offset())
		mockerName += fmt.Sprintf("_%d_%d", limiter.Limit(), limiter.Offset())

		mocker += ")"

		fixture := mockerName

		if res != nil && len(res) != 0 {
			pks := "[]string{\n"

			for _, r := range res {

				pks += "\"" + r.Primary() + "\",\n"

			}

			pks += "}"

			fixture += ".ByFixturePKWithKeys(ctx, " + pks + ", []string{" + fixturesKey + "})"
		} else {
			fixture += ".EmptyByKeys(ctx, " + fixturesKey + ")"
		}

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "promoperiods", Results: res}, nil
	}
}

func (obj *Promoperiods) MockSelectByPlanTypePriceRequest(ctx context.Context) []byte {
	key := []PlanTypePriceIndexType{{
		PlanType: obj.GetPlanType(),
		Price:    obj.GetPrice(),
	},
	}
	return obj.MockSelectByPlanTypePricesRequest(ctx, key)
}

func serializePrice(ctx context.Context, vPrice float64) string {
	v, err := serializerJSON.PrintfMarshal("%.2f", vPrice)
	if err != nil {
		activerecord.Logger().Fatal(ctx, err)
	}
	return v
}
func (obj *Promoperiods) MockSelectByPlanTypePricesRequest(ctx context.Context, keys []PlanTypePriceIndexType) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockSelectByPlanTypePricesRequest": keys, "Repo": "Promoperiods"})
	keysPacked := [][][]byte{}

	var (
		data []byte
		err  error
	)
	for _, key := range keys {
		keysField := [][]byte{}

		data, err = packPlanType([]byte{}, key.PlanType)
		if err != nil {
			log.Fatal(ctx, err)
		}

		keysField = append(keysField, data)

		data, err = packPrice([]byte{}, key.Price)
		if err != nil {
			log.Fatal(ctx, err)
		}

		keysField = append(keysField, data)
		keysPacked = append(keysPacked, keysField)
	}

	log.Debug(ctx, fmt.Sprintf("Packed Keys: '% X'", keysPacked))

	return octopus.PackSelect(namespace, 3, 0, 0, keysPacked)

}

func SelectByPlanTypePriceMockerLogger(keys []PlanTypePriceIndexType, res PromoperiodsList) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {
		fixturesKey := ""
		for _, key := range keys {
			fixturesKey += "promoperiods.PlanTypePriceIndexType{\n"

			fixturesKey += "PlanType: \"" + key.PlanType + "\",\n"

			skey, err := serializerJSON.PrintfMarshal("%.2f", key.Price)
			if err != nil {
				return activerecord.MockerLogger{}, err
			}

			fixturesKey += "Price: \"" + skey + "\",\n"

			fixturesKey += "},\n"
		}

		mockerName := "mockerPromoperiodsByPlanTypePrice"
		mocker := "fixture.GetPromoperiodsByPlanTypePriceMocker("

		mocker += ")"

		fixture := mockerName

		if res != nil && len(res) != 0 {
			pks := "[]string{\n"

			for _, r := range res {

				pks += "\"" + r.Primary() + "\",\n"

			}

			pks += "}"

			fixture += ".ByFixturePKWithKeys(ctx, " + pks + ", []PlanTypePriceIndexType{" + fixturesKey + "})"
		} else {
			fixture += ".EmptyByKeys(ctx, " + fixturesKey + ")"
		}

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "promoperiods", Results: res}, nil
	}
}

func (obj *Promoperiods) MockSelectByEmailCodeRequest(ctx context.Context) []byte {
	key := []EmailCodeIndexType{{
		Email: obj.GetEmail(),
		Code:  obj.GetCode(),
	},
	}
	return obj.MockSelectByEmailCodesRequest(ctx, key)
}

func (obj *Promoperiods) MockSelectByEmailCodesRequest(ctx context.Context, keys []EmailCodeIndexType) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockSelectByEmailCodesRequest": keys, "Repo": "Promoperiods"})
	keysPacked := [][][]byte{}

	var (
		data []byte
		err  error
	)
	for _, key := range keys {
		keysField := [][]byte{}

		data, err = packEmail([]byte{}, key.Email)
		if err != nil {
			log.Fatal(ctx, err)
		}

		keysField = append(keysField, data)

		data, err = packCode([]byte{}, key.Code)
		if err != nil {
			log.Fatal(ctx, err)
		}

		keysField = append(keysField, data)
		keysPacked = append(keysPacked, keysField)
	}

	log.Debug(ctx, fmt.Sprintf("Packed Keys: '% X'", keysPacked))

	return octopus.PackSelect(namespace, 4, 0, 0, keysPacked)

}

func SelectByEmailCodeMockerLogger(keys []EmailCodeIndexType, res PromoperiodsList) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {
		fixturesKey := ""
		for _, key := range keys {
			fixturesKey += "promoperiods.EmailCodeIndexType{\n"

			fixturesKey += "Email: \"" + key.Email + "\",\n"

			fixturesKey += "Code: \"" + key.Code + "\",\n"

			fixturesKey += "},\n"
		}

		mockerName := "mockerPromoperiodsByEmailCode"
		mocker := "fixture.GetPromoperiodsByEmailCodeMocker("

		mocker += ")"

		fixture := mockerName

		if res != nil && len(res) != 0 {
			pks := "[]string{\n"

			for _, r := range res {

				pks += "\"" + r.Primary() + "\",\n"

			}

			pks += "}"

			fixture += ".ByFixturePKWithKeys(ctx, " + pks + ", []EmailCodeIndexType{" + fixturesKey + "})"
		} else {
			fixture += ".EmptyByKeys(ctx, " + fixturesKey + ")"
		}

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "promoperiods", Results: res}, nil
	}
}

func (obj *Promoperiods) MockSelectByEmailActionRequest(ctx context.Context) []byte {
	key := []EmailActionIndexType{{
		Email:  obj.GetEmail(),
		Action: obj.GetAction(),
	},
	}
	return obj.MockSelectByEmailActionsRequest(ctx, key)
}

func (obj *Promoperiods) MockSelectByEmailActionsRequest(ctx context.Context, keys []EmailActionIndexType) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockSelectByEmailActionsRequest": keys, "Repo": "Promoperiods"})
	keysPacked := [][][]byte{}

	var (
		data []byte
		err  error
	)
	for _, key := range keys {
		keysField := [][]byte{}

		data, err = packEmail([]byte{}, key.Email)
		if err != nil {
			log.Fatal(ctx, err)
		}

		keysField = append(keysField, data)

		data, err = packAction([]byte{}, key.Action)
		if err != nil {
			log.Fatal(ctx, err)
		}

		keysField = append(keysField, data)
		keysPacked = append(keysPacked, keysField)
	}

	log.Debug(ctx, fmt.Sprintf("Packed Keys: '% X'", keysPacked))

	return octopus.PackSelect(namespace, 5, 0, 0, keysPacked)

}

func SelectByEmailActionMockerLogger(keys []EmailActionIndexType, res PromoperiodsList) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {
		fixturesKey := ""
		for _, key := range keys {
			fixturesKey += "promoperiods.EmailActionIndexType{\n"

			fixturesKey += "Email: \"" + key.Email + "\",\n"

			fixturesKey += "Action: \"" + key.Action + "\",\n"

			fixturesKey += "},\n"
		}

		mockerName := "mockerPromoperiodsByEmailAction"
		mocker := "fixture.GetPromoperiodsByEmailActionMocker("

		mocker += ")"

		fixture := mockerName

		if res != nil && len(res) != 0 {
			pks := "[]string{\n"

			for _, r := range res {

				pks += "\"" + r.Primary() + "\",\n"

			}

			pks += "}"

			fixture += ".ByFixturePKWithKeys(ctx, " + pks + ", []EmailActionIndexType{" + fixturesKey + "})"
		} else {
			fixture += ".EmptyByKeys(ctx, " + fixturesKey + ")"
		}

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "promoperiods", Results: res}, nil
	}
}

func (obj *Promoperiods) MockSelectByEmailPartRequest(ctx context.Context, limiter activerecord.SelectorLimiter) []byte {
	key := []string{
		obj.GetEmail(),
	}
	return obj.MockSelectByEmailPartsRequest(ctx, key, limiter)
}

func (obj *Promoperiods) MockSelectByEmailPartsRequest(ctx context.Context, keys []string, limiter activerecord.SelectorLimiter) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockSelectByEmailPartsRequest": keys, "Repo": "Promoperiods"})
	keysPacked := [][][]byte{}

	var (
		data []byte
		err  error
	)
	for _, key := range keys {
		keysField := [][]byte{}

		data, err = packEmail([]byte{}, key)
		if err != nil {
			log.Fatal(ctx, err)
			return nil
		}

		keysField = append(keysField, data)
		keysPacked = append(keysPacked, keysField)
	}

	log.Debug(ctx, fmt.Sprintf("Packed Keys: '% X'", keysPacked))

	return octopus.PackSelect(namespace, 4, limiter.Offset(), limiter.Limit(), keysPacked)

}

func SelectByEmailPartMockerLogger(keys []string, res PromoperiodsList, limiter activerecord.SelectorLimiter) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {
		fixturesKey := ""
		for _, key := range keys {

			fixturesKey += "\"" + key + "\""

			fixturesKey += ",\n"
		}

		mockerName := "mockerPromoperiodsByEmailPart"
		mocker := "fixture.GetPromoperiodsByEmailPartMocker("

		mocker += fmt.Sprintf("activerecord.NewLimitOffset(%d, %d)", limiter.Limit(), limiter.Offset())
		mockerName += fmt.Sprintf("_%d_%d", limiter.Limit(), limiter.Offset())

		mocker += ")"

		fixture := mockerName

		if res != nil && len(res) != 0 {
			pks := "[]string{\n"

			for _, r := range res {

				pks += "\"" + r.Primary() + "\",\n"

			}

			pks += "}"

			fixture += ".ByFixturePKWithKeys(ctx, " + pks + ", []string{" + fixturesKey + "})"
		} else {
			fixture += ".EmptyByKeys(ctx, " + fixturesKey + ")"
		}

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "promoperiods", Results: res}, nil
	}
}

// indexes
func (obj *Promoperiods) RepoSelector(ctx context.Context) (any, error) {
	data, err := SelectByPrimary(ctx, obj.Primary())
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, fmt.Errorf("select Promoperiods with pk %v: %w", obj.Primary(), activerecord.ErrNoData)
	}
	return data, err
}

func (obj *Promoperiods) MockUpdate(ctx context.Context) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockUpdate": obj.BaseField.UpdateOps, "Repo": "Promoperiods"})

	//todo repaired logic not implemented

	pk, err := obj.packPk()
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	w := octopus.PackUpdate(namespace, pk, obj.BaseField.UpdateOps)

	log.Debug(ctx, fmt.Sprintf("Update packed tuple: '%X'\n", w))

	return w
}

func (obj *Promoperiods) MockInsertOrReplace(ctx context.Context) []byte {
	return obj.mockInsertReplace(ctx, octopus.InsertModeInserOrReplace)
}

func (obj *Promoperiods) MockInsert(ctx context.Context) []byte {
	obj.BaseField.Exists = false

	return obj.mockInsertReplace(ctx, octopus.InsertModeInsert)
}

func (obj *Promoperiods) MockReplace(ctx context.Context) []byte {
	obj.BaseField.Exists = true

	return obj.mockInsertReplace(ctx, octopus.InsertModeReplace)
}

func (obj *Promoperiods) mockInsertReplace(ctx context.Context, insertMode octopus.InsertMode) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockInsertReplacePrimary": obj.PrimaryString(), "MockInsertReplaceMode": insertMode, "Repo": "Promoperiods"})

	var tuple [][]byte
	var data []byte
	var err error

	// ID
	data, err = packID([]byte{}, obj.GetID())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// Code
	data, err = packCode([]byte{}, obj.GetCode())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// Email
	data, err = packEmail([]byte{}, obj.GetEmail())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// Start
	data, err = packStart([]byte{}, obj.GetStart())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// Finish
	data, err = packFinish([]byte{}, obj.GetFinish())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// Action
	data, err = packAction([]byte{}, obj.GetAction())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// Platform
	data, err = packPlatform([]byte{}, obj.GetPlatform())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// Promobunch
	data, err = packPromobunch([]byte{}, obj.GetPromobunch())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// Platforms
	data, err = packPlatforms([]byte{}, obj.GetPlatforms())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// PlanID
	data, err = packPlanID([]byte{}, obj.GetPlanID())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// PlanType
	data, err = packPlanType([]byte{}, obj.GetPlanType())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// Price
	data, err = packPrice([]byte{}, obj.GetPrice())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	if len(obj.BaseField.ExtraFields) > 0 {
		tuple = append(tuple, obj.BaseField.ExtraFields...)
	}

	w := octopus.PackInsertReplace(namespace, insertMode, tuple)

	log.Debug(ctx, fmt.Sprintf("insertReplace packed tuple: '%X'\n", w))

	return w
}
