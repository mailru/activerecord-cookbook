// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.5.3-9-geaa00ca (Commit: eaa00caf)
package arobj

import (
	"context"
	"fmt"

	"strconv"

	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
)

func (obj *ArObj) MockSelectResponse() ([][]byte, error) {
	tuple := [][]byte{}

	var data []byte

	var err error

	data, err = packID([]byte{}, obj.GetID())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packName([]byte{}, obj.GetName())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packAnotherID([]byte{}, obj.GetAnotherID())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packType([]byte{}, obj.GetType())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	data, err = packFlags([]byte{}, obj.GetFlags())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	return tuple, nil
}

func (obj *ArObj) MockSelectByIDRequest(ctx context.Context) []byte {
	key := []int32{
		obj.GetID(),
	}
	return obj.MockSelectByIDsRequest(ctx, key)
}

func (obj *ArObj) MockSelectByIDsRequest(ctx context.Context, keys []int32) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockSelectByIDsRequest": keys, "Repo": "ArObj"})
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

func SelectByIDMockerLogger(keys []int32, res ArObjList) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {
		fixturesKey := ""
		for _, key := range keys {

			fixturesKey += strconv.FormatInt(int64(key), 10)

			fixturesKey += ",\n"
		}

		mockerName := "mockerArObjByID"
		mocker := "fixture.GetArObjByIDMocker("

		mocker += ")"

		fixture := mockerName

		if res != nil && len(res) != 0 {
			pks := "[]int32{\n"

			for _, r := range res {

				pks += strconv.FormatInt(int64(r.Primary()), 10) + ",\n"

			}

			pks += "}"

			fixture += ".ByFixturePKWithKeys(ctx, " + pks + ", []int32{" + fixturesKey + "})"
		} else {
			fixture += ".EmptyByKeys(ctx, " + fixturesKey + ")"
		}

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "arobj", Results: res}, nil
	}
}

func (obj *ArObj) MockSelectByTypeRequest(ctx context.Context, limiter activerecord.SelectorLimiter) []byte {
	key := []string{
		obj.GetType(),
	}
	return obj.MockSelectByTypesRequest(ctx, key, limiter)
}

func (obj *ArObj) MockSelectByTypesRequest(ctx context.Context, keys []string, limiter activerecord.SelectorLimiter) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockSelectByTypesRequest": keys, "Repo": "ArObj"})
	keysPacked := [][][]byte{}

	var (
		data []byte
		err  error
	)
	for _, key := range keys {
		keysField := [][]byte{}

		data, err = packType([]byte{}, key)
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

func SelectByTypeMockerLogger(keys []string, res ArObjList, limiter activerecord.SelectorLimiter) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {
		fixturesKey := ""
		for _, key := range keys {

			fixturesKey += "\"" + key + "\""

			fixturesKey += ",\n"
		}

		mockerName := "mockerArObjByType"
		mocker := "fixture.GetArObjByTypeMocker("

		mocker += fmt.Sprintf("activerecord.NewLimitOffset(%d, %d)", limiter.Limit(), limiter.Offset())
		mockerName += fmt.Sprintf("_%d_%d", limiter.Limit(), limiter.Offset())

		mocker += ")"

		fixture := mockerName

		if res != nil && len(res) != 0 {
			pks := "[]int32{\n"

			for _, r := range res {

				pks += strconv.FormatInt(int64(r.Primary()), 10) + ",\n"

			}

			pks += "}"

			fixture += ".ByFixturePKWithKeys(ctx, " + pks + ", []string{" + fixturesKey + "})"
		} else {
			fixture += ".EmptyByKeys(ctx, " + fixturesKey + ")"
		}

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "arobj", Results: res}, nil
	}
}

func (obj *ArObj) MockSelectByTypeIDRequest(ctx context.Context) []byte {
	key := []TypeIDIndexType{{
		Type: obj.GetType(),
		ID:   obj.GetID(),
	},
	}
	return obj.MockSelectByTypeIDsRequest(ctx, key)
}

func (obj *ArObj) MockSelectByTypeIDsRequest(ctx context.Context, keys []TypeIDIndexType) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockSelectByTypeIDsRequest": keys, "Repo": "ArObj"})
	keysPacked := [][][]byte{}

	var (
		data []byte
		err  error
	)
	for _, key := range keys {
		keysField := [][]byte{}

		data, err = packType([]byte{}, key.Type)
		if err != nil {
			log.Fatal(ctx, err)
		}

		keysField = append(keysField, data)

		data, err = packID([]byte{}, key.ID)
		if err != nil {
			log.Fatal(ctx, err)
		}

		keysField = append(keysField, data)
		keysPacked = append(keysPacked, keysField)
	}

	log.Debug(ctx, fmt.Sprintf("Packed Keys: '% X'", keysPacked))

	return octopus.PackSelect(namespace, 2, 0, 0, keysPacked)

}

func SelectByTypeIDMockerLogger(keys []TypeIDIndexType, res ArObjList) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {
		fixturesKey := ""
		for _, key := range keys {
			fixturesKey += "arobj.TypeIDIndexType{\n"

			fixturesKey += "Type: \"" + key.Type + "\",\n"

			fixturesKey += "ID: " + strconv.FormatInt(int64(key.ID), 10) + ",\n"

			fixturesKey += "},\n"
		}

		mockerName := "mockerArObjByTypeID"
		mocker := "fixture.GetArObjByTypeIDMocker("

		mocker += ")"

		fixture := mockerName

		if res != nil && len(res) != 0 {
			pks := "[]int32{\n"

			for _, r := range res {

				pks += strconv.FormatInt(int64(r.Primary()), 10) + ",\n"

			}

			pks += "}"

			fixture += ".ByFixturePKWithKeys(ctx, " + pks + ", []TypeIDIndexType{" + fixturesKey + "})"
		} else {
			fixture += ".EmptyByKeys(ctx, " + fixturesKey + ")"
		}

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "arobj", Results: res}, nil
	}
}

func (obj *ArObj) MockSelectByTypePartRequest(ctx context.Context, limiter activerecord.SelectorLimiter) []byte {
	key := []string{
		obj.GetType(),
	}
	return obj.MockSelectByTypePartsRequest(ctx, key, limiter)
}

func (obj *ArObj) MockSelectByTypePartsRequest(ctx context.Context, keys []string, limiter activerecord.SelectorLimiter) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockSelectByTypePartsRequest": keys, "Repo": "ArObj"})
	keysPacked := [][][]byte{}

	var (
		data []byte
		err  error
	)
	for _, key := range keys {
		keysField := [][]byte{}

		data, err = packType([]byte{}, key)
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

func SelectByTypePartMockerLogger(keys []string, res ArObjList, limiter activerecord.SelectorLimiter) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {
		fixturesKey := ""
		for _, key := range keys {

			fixturesKey += "\"" + key + "\""

			fixturesKey += ",\n"
		}

		mockerName := "mockerArObjByTypePart"
		mocker := "fixture.GetArObjByTypePartMocker("

		mocker += fmt.Sprintf("activerecord.NewLimitOffset(%d, %d)", limiter.Limit(), limiter.Offset())
		mockerName += fmt.Sprintf("_%d_%d", limiter.Limit(), limiter.Offset())

		mocker += ")"

		fixture := mockerName

		if res != nil && len(res) != 0 {
			pks := "[]int32{\n"

			for _, r := range res {

				pks += strconv.FormatInt(int64(r.Primary()), 10) + ",\n"

			}

			pks += "}"

			fixture += ".ByFixturePKWithKeys(ctx, " + pks + ", []string{" + fixturesKey + "})"
		} else {
			fixture += ".EmptyByKeys(ctx, " + fixturesKey + ")"
		}

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "arobj", Results: res}, nil
	}
}

func (obj *ArObj) RepoSelector(ctx context.Context) (any, error) {
	data, err := SelectByPrimary(ctx, obj.Primary())
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, fmt.Errorf("select ArObj with pk %v: %w", obj.Primary(), activerecord.ErrNoData)
	}
	return data, err
}

func (obj *ArObj) MockUpdate(ctx context.Context) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockUpdate": obj.BaseField.UpdateOps, "Repo": "ArObj"})

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

func (obj *ArObj) MockInsertOrReplace(ctx context.Context) []byte {
	return obj.mockInsertReplace(ctx, octopus.InsertModeInserOrReplace)
}

func (obj *ArObj) MockInsert(ctx context.Context) []byte {
	obj.BaseField.Exists = false

	return obj.mockInsertReplace(ctx, octopus.InsertModeInsert)
}

func (obj *ArObj) MockReplace(ctx context.Context) []byte {
	obj.BaseField.Exists = true

	return obj.mockInsertReplace(ctx, octopus.InsertModeReplace)
}

func (obj *ArObj) mockInsertReplace(ctx context.Context, insertMode octopus.InsertMode) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockInsertReplacePrimary": obj.PrimaryString(), "MockInsertReplaceMode": insertMode, "Repo": "ArObj"})

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

	// Name
	data, err = packName([]byte{}, obj.GetName())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// AnotherID
	data, err = packAnotherID([]byte{}, obj.GetAnotherID())
	if err != nil {
		log.Fatal(ctx, err)
		return nil
	}

	tuple = append(tuple, data)

	// Type
	data, err = packType([]byte{}, obj.GetType())
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

	if len(obj.BaseField.ExtraFields) > 0 {
		tuple = append(tuple, obj.BaseField.ExtraFields...)
	}

	w := octopus.PackInsertReplace(namespace, insertMode, tuple)

	log.Debug(ctx, fmt.Sprintf("insertReplace packed tuple: '%X'\n", w))

	return w
}
