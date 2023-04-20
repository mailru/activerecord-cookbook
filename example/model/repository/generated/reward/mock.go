// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.5.3 (Commit: e0ffb560)
package reward

import (
	"context"
	"fmt"

	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
)

func (obj *Reward) MockSelectResponse() ([][]byte, error) {
	tuple := [][]byte{}

	var data []byte

	var err error

	data, err = packCode([]byte{}, obj.GetCode())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packServices([]byte{}, obj.GetServices())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packPartner([]byte{}, obj.GetPartner())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packExtra([]byte{}, obj.GetExtra())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packFlags([]byte{}, obj.GetFlags())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packUnlocked([]byte{}, obj.GetUnlocked())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packDescription([]byte{}, obj.GetDescription())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	return tuple, nil
}

func (obj *Reward) MockSelectByCodeRequest(ctx context.Context) []byte {
	key := []string{
		obj.GetCode(),
	}
	return obj.MockSelectByCodesRequest(ctx, key)
}

func (obj *Reward) MockSelectByCodesRequest(ctx context.Context, keys []string) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockSelectByCodesRequest": keys, "Repo": "Reward"})
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

	return octopus.PackSelect(namespace, 0, 0, 0, keysPacked)

}

func SelectByCodeMockerLogger(keys []string, res RewardList) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {
		fixturesKey := ""
		for _, key := range keys {

			fixturesKey += "\"" + key + "\""

			fixturesKey += ",\n"
		}

		mockerName := "mockerRewardByCode"
		mocker := "fixture.GetRewardByCodeMocker("

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

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "reward", Results: res}, nil
	}
}

func (obj *Reward) MockSelectByPartnerRequest(ctx context.Context, limiter activerecord.SelectorLimiter) []byte {
	key := []string{
		obj.GetPartner(),
	}
	return obj.MockSelectByPartnersRequest(ctx, key, limiter)
}

func (obj *Reward) MockSelectByPartnersRequest(ctx context.Context, keys []string, limiter activerecord.SelectorLimiter) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockSelectByPartnersRequest": keys, "Repo": "Reward"})
	keysPacked := [][][]byte{}

	var (
		data []byte
		err  error
	)
	for _, key := range keys {
		keysField := [][]byte{}

		data, err = packPartner([]byte{}, key)
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

func SelectByPartnerMockerLogger(keys []string, res RewardList, limiter activerecord.SelectorLimiter) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {
		fixturesKey := ""
		for _, key := range keys {

			fixturesKey += "\"" + key + "\""

			fixturesKey += ",\n"
		}

		mockerName := "mockerRewardByPartner"
		mocker := "fixture.GetRewardByPartnerMocker("

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

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "reward", Results: res}, nil
	}
}

func (obj *Reward) RepoSelector(ctx context.Context) (any, error) {
	data, err := SelectByPrimary(ctx, obj.Primary())
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, fmt.Errorf("select Reward with pk %v: %w", obj.Primary(), activerecord.ErrNoData)
	}
	return data, err
}

func (obj *Reward) MockUpdate(ctx context.Context) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockUpdate": obj.BaseField.UpdateOps, "Repo": "Reward"})

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

func (obj *Reward) MockInsertOrReplace(ctx context.Context) []byte {
	return obj.mockInsertReplace(ctx, octopus.InsertModeInserOrReplace)
}

func (obj *Reward) MockInsert(ctx context.Context) []byte {
	obj.BaseField.Exists = false

	return obj.mockInsertReplace(ctx, octopus.InsertModeInsert)
}

func (obj *Reward) MockReplace(ctx context.Context) []byte {
	obj.BaseField.Exists = true

	return obj.mockInsertReplace(ctx, octopus.InsertModeReplace)
}

func (obj *Reward) mockInsertReplace(ctx context.Context, insertMode octopus.InsertMode) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockInsertReplacePrimary": obj.PrimaryString(), "MockInsertReplaceMode": insertMode, "Repo": "Reward"})

	var tuple [][]byte
	var data []byte
	var err error

	// Code
	data, err = packCode([]byte{}, obj.GetCode())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// Services
	data, err = packServices([]byte{}, obj.GetServices())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// Partner
	data, err = packPartner([]byte{}, obj.GetPartner())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// Extra
	data, err = packExtra([]byte{}, obj.GetExtra())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// Flags
	data, err = packFlags([]byte{}, obj.GetFlags())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// Unlocked
	data, err = packUnlocked([]byte{}, obj.GetUnlocked())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// Description
	data, err = packDescription([]byte{}, obj.GetDescription())
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
