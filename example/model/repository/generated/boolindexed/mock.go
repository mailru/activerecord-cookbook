// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.5.3-12-g7f6d003 (Commit: 7f6d003b)
package boolindexed

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
)

func (obj *Boolindexed) MockSelectResponse() ([][]byte, error) {
	tuple := [][]byte{}

	var data []byte

	var err error

	data, err = packCode([]byte{}, obj.GetCode())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packInvisible([]byte{}, obj.GetInvisible())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	return tuple, nil
}

func (obj *Boolindexed) MockSelectByCodeRequest(ctx context.Context) []byte {
	key := []string{
		obj.GetCode(),
	}
	return obj.MockSelectByCodesRequest(ctx, key)
}

func (obj *Boolindexed) MockSelectByCodesRequest(ctx context.Context, keys []string) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockSelectByCodesRequest": keys, "Repo": "Boolindexed"})
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

func SelectByCodeMockerLogger(keys []string, res BoolindexedList) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {
		fixturesKey := ""
		for _, key := range keys {

			fixturesKey += "\"" + key + "\""

			fixturesKey += ",\n"
		}

		mockerName := "mockerBoolindexedByCode"
		mocker := "fixture.GetBoolindexedByCodeMocker("

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

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "boolindexed", Results: res}, nil
	}
}

func (obj *Boolindexed) MockSelectByInvisibleRequest(ctx context.Context, limiter activerecord.SelectorLimiter) []byte {
	key := []bool{
		obj.GetInvisible(),
	}
	return obj.MockSelectByInvisiblesRequest(ctx, key, limiter)
}

func (obj *Boolindexed) MockSelectByInvisiblesRequest(ctx context.Context, keys []bool, limiter activerecord.SelectorLimiter) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockSelectByInvisiblesRequest": keys, "Repo": "Boolindexed"})
	keysPacked := [][][]byte{}

	var (
		data []byte
		err  error
	)
	for _, key := range keys {
		keysField := [][]byte{}

		data, err = packInvisible([]byte{}, key)
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

func SelectByInvisibleMockerLogger(keys []bool, res BoolindexedList, limiter activerecord.SelectorLimiter) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {
		fixturesKey := ""
		for _, key := range keys {

			fixturesKey += strconv.FormatBool(key)

			fixturesKey += ",\n"
		}

		mockerName := "mockerBoolindexedByInvisible"
		mocker := "fixture.GetBoolindexedByInvisibleMocker("

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

			fixture += ".ByFixturePKWithKeys(ctx, " + pks + ", []bool{" + fixturesKey + "})"
		} else {
			fixture += ".EmptyByKeys(ctx, " + fixturesKey + ")"
		}

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "boolindexed", Results: res}, nil
	}
}

func (obj *Boolindexed) RepoSelector(ctx context.Context) (any, error) {
	data, err := SelectByPrimary(ctx, obj.Primary())
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, fmt.Errorf("select Boolindexed with pk %v: %w", obj.Primary(), activerecord.ErrNoData)
	}
	return data, err
}

func (obj *Boolindexed) MockUpdate(ctx context.Context) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockUpdate": obj.BaseField.UpdateOps, "Repo": "Boolindexed"})

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

func (obj *Boolindexed) MockInsertOrReplace(ctx context.Context) []byte {
	return obj.mockInsertReplace(ctx, octopus.InsertModeInserOrReplace)
}

func (obj *Boolindexed) MockInsert(ctx context.Context) []byte {
	obj.BaseField.Exists = false

	return obj.mockInsertReplace(ctx, octopus.InsertModeInsert)
}

func (obj *Boolindexed) MockReplace(ctx context.Context) []byte {
	obj.BaseField.Exists = true

	return obj.mockInsertReplace(ctx, octopus.InsertModeReplace)
}

func (obj *Boolindexed) mockInsertReplace(ctx context.Context, insertMode octopus.InsertMode) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockInsertReplacePrimary": obj.PrimaryString(), "MockInsertReplaceMode": insertMode, "Repo": "Boolindexed"})

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

	// Invisible
	data, err = packInvisible([]byte{}, obj.GetInvisible())
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
