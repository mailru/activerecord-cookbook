// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.5.3-7-g1454a87 (Commit: 1454a870)
package arobj

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"strings"

	"math"
	"strconv"

	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/iproto/iproto"
	"github.com/mailru/activerecord/pkg/octopus"
)

type ArObj struct {
	octopus.BaseField
	fieldID        int32
	fieldName      string
	fieldAnotherID int32
	fieldType      string
	fieldFlags     uint32
}

type ArObjList []*ArObj

const (
	namespace uint32 = 5
	cntFields uint32 = 5
)

func New(ctx context.Context) *ArObj {
	newObj := ArObj{}
	newObj.BaseField.UpdateOps = []octopus.Ops{}
	newObj.BaseField.ExtraFields = [][]byte{}
	newObj.BaseField.Objects = map[string][]octopus.ModelStruct{}

	return &newObj
}

func TupleToStruct(ctx context.Context, tuple octopus.TupleData) (*ArObj, error) {
	np := New(ctx)

	valID, err := UnpackID(bytes.NewReader(tuple.Data[0]))
	if err != nil {
		return nil, err
	}

	np.SetID(valID)
	valName, err := UnpackName(bytes.NewReader(tuple.Data[1]))
	if err != nil {
		return nil, err
	}

	np.SetName(valName)
	valAnotherID, err := UnpackAnotherID(bytes.NewReader(tuple.Data[2]))
	if err != nil {
		return nil, err
	}

	np.SetAnotherID(valAnotherID)
	valType, err := UnpackType(bytes.NewReader(tuple.Data[3]))
	if err != nil {
		return nil, err
	}

	np.SetType(valType)
	valFlags, err := UnpackFlags(bytes.NewReader(tuple.Data[4]))
	if err != nil {
		return nil, err
	}

	np.SetFlags(valFlags)

	np.BaseField.Exists = true
	np.BaseField.UpdateOps = []octopus.Ops{}

	if tuple.Cnt > cntFields {
		logger := activerecord.Logger()

		logger.Warn(ctx, "ArObj", np.PrimaryString(), "Extra fields")

		np.BaseField.ExtraFields = tuple.Data[cntFields:]
	}

	return np, nil
}

func NewFromBox(ctx context.Context, tuples []octopus.TupleData) ([]*ArObj, error) {
	logger := activerecord.Logger()

	logger.Debug(ctx, "ArObj", fmt.Sprintf("Cnt tuples %d", len(tuples)))

	ret := make([]*ArObj, 0, len(tuples))

	for num, tuple := range tuples {
		var repaired bool

		if tuple.Cnt < cntFields {
			return nil, fmt.Errorf("not enought selected fields %d in response tuple: %d but expected %d fields", tuple.Cnt, num, cntFields)
		}

		np, err := TupleToStruct(ctx, tuple)
		if err != nil {
			logger.Error(ctx, "ArObj", fmt.Sprintf("error unpack tuple %s", err))
			return nil, err
		}

		np.BaseField.Repaired = repaired
		ret = append(ret, np)
	}

	return ret, nil
}
func packID(w []byte, ID int32) ([]byte, error) {
	pvar := uint32(ID)

	return iproto.PackUint32(w, pvar, iproto.ModeDefault), nil
}

func UnpackID(r *bytes.Reader) (ret int32, errRet error) {
	var ID uint32

	err := iproto.UnpackUint32(r, &ID, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field ID in tuple: '%w'", err)
		return
	}

	bvar := int32(ID)

	svar := bvar

	return svar, nil
}

func (obj *ArObj) GetID() int32 {
	return obj.fieldID
}

func (obj *ArObj) SetID(ID int32) error {
	if obj.BaseField.Exists {
		return fmt.Errorf("can't modify field included in primary key")
	}

	data, err := packID([]byte{}, ID)
	if err != nil {
		return err
	}

	obj.BaseField.UpdateOps = append(obj.BaseField.UpdateOps, octopus.Ops{Field: 0, Op: octopus.OpSet, Value: data})
	obj.fieldID = ID

	return nil
}

func packName(w []byte, Name string) ([]byte, error) {
	pvar := Name

	return octopus.PackString(w, pvar, iproto.ModeDefault), nil
}

func UnpackName(r *bytes.Reader) (ret string, errRet error) {
	var Name string

	err := octopus.UnpackString(r, &Name, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field Name in tuple: '%w'", err)
		return
	}

	bvar := Name

	svar := bvar

	return svar, nil
}

func (obj *ArObj) GetName() string {
	return obj.fieldName
}

func (obj *ArObj) SetName(Name string) error {
	data, err := packName([]byte{}, Name)
	if err != nil {
		return err
	}

	if len(data) > 256 {
		return fmt.Errorf("max length of field 'ArObj.Name' is '%d' (received '%d')", 256, len(data))
	}

	obj.BaseField.UpdateOps = append(obj.BaseField.UpdateOps, octopus.Ops{Field: 1, Op: octopus.OpSet, Value: data})
	obj.fieldName = Name

	return nil
}

func packAnotherID(w []byte, AnotherID int32) ([]byte, error) {
	pvar := uint32(AnotherID)

	return iproto.PackUint32(w, pvar, iproto.ModeDefault), nil
}

func UnpackAnotherID(r *bytes.Reader) (ret int32, errRet error) {
	var AnotherID uint32

	err := iproto.UnpackUint32(r, &AnotherID, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field AnotherID in tuple: '%w'", err)
		return
	}

	bvar := int32(AnotherID)

	svar := bvar

	return svar, nil
}

func (obj *ArObj) GetAnotherID() int32 {
	return obj.fieldAnotherID
}

func (obj *ArObj) SetAnotherID(AnotherID int32) error {
	data, err := packAnotherID([]byte{}, AnotherID)
	if err != nil {
		return err
	}

	obj.BaseField.UpdateOps = append(obj.BaseField.UpdateOps, octopus.Ops{Field: 2, Op: octopus.OpSet, Value: data})
	obj.fieldAnotherID = AnotherID

	return nil
}

func (obj *ArObj) IncAnotherID(mutArg int32) error {
	if mutArg == 0 {
		return nil
	}

	if uint64(math.MaxInt32-obj.fieldAnotherID) < uint64(mutArg) {
		return fmt.Errorf("overflow type 'int32' after Inc %d", mutArg)
	}

	data := iproto.PackUint32([]byte{}, uint32(mutArg), iproto.ModeDefault)

	obj.BaseField.UpdateOps = append(obj.BaseField.UpdateOps, octopus.Ops{Field: 2, Op: octopus.OpAdd, Value: data})
	obj.fieldAnotherID += int32(mutArg)

	return nil
}

func (obj *ArObj) DecAnotherID(mutArg int32) error {
	if mutArg == 0 {
		return nil
	}

	if uint64(obj.fieldAnotherID-math.MinInt32) < uint64(mutArg) {
		return fmt.Errorf("overflow type 'int32' after Dec %d", mutArg)
	}

	data := iproto.PackUint32([]byte{}, uint32(-mutArg), iproto.ModeDefault)

	obj.BaseField.UpdateOps = append(obj.BaseField.UpdateOps, octopus.Ops{Field: 2, Op: octopus.OpAdd, Value: data})
	obj.fieldAnotherID -= int32(mutArg)

	return nil
}

func packType(w []byte, Type string) ([]byte, error) {
	pvar := Type

	return octopus.PackString(w, pvar, iproto.ModeDefault), nil
}

func UnpackType(r *bytes.Reader) (ret string, errRet error) {
	var Type string

	err := octopus.UnpackString(r, &Type, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field Type in tuple: '%w'", err)
		return
	}

	bvar := Type

	svar := bvar

	return svar, nil
}

func (obj *ArObj) GetType() string {
	return obj.fieldType
}

func (obj *ArObj) SetType(Type string) error {
	data, err := packType([]byte{}, Type)
	if err != nil {
		return err
	}

	if len(data) > 64 {
		return fmt.Errorf("max length of field 'ArObj.Type' is '%d' (received '%d')", 64, len(data))
	}

	obj.BaseField.UpdateOps = append(obj.BaseField.UpdateOps, octopus.Ops{Field: 3, Op: octopus.OpSet, Value: data})
	obj.fieldType = Type

	return nil
}

func packFlags(w []byte, Flags uint32) ([]byte, error) {
	pvar := Flags

	return iproto.PackUint32(w, pvar, iproto.ModeDefault), nil
}

func UnpackFlags(r *bytes.Reader) (ret uint32, errRet error) {
	var Flags uint32

	err := iproto.UnpackUint32(r, &Flags, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field Flags in tuple: '%w'", err)
		return
	}

	bvar := Flags

	svar := bvar

	return svar, nil
}

func (obj *ArObj) GetFlags() uint32 {
	return obj.fieldFlags
}

func (obj *ArObj) SetFlags(Flags uint32) error {
	data, err := packFlags([]byte{}, Flags)
	if err != nil {
		return err
	}

	obj.BaseField.UpdateOps = append(obj.BaseField.UpdateOps, octopus.Ops{Field: 4, Op: octopus.OpSet, Value: data})
	obj.fieldFlags = Flags

	return nil
}

func (obj *ArObj) SetBitFlags(mutArg uint32) error {
	if mutArg == 0 || obj.fieldFlags|mutArg == obj.fieldFlags {
		return nil
	}

	data := iproto.PackUint32([]byte{}, uint32(mutArg), iproto.ModeDefault)

	obj.BaseField.UpdateOps = append(obj.BaseField.UpdateOps, octopus.Ops{Field: 4, Op: octopus.OpOr, Value: data})
	obj.fieldFlags |= mutArg

	return nil
}

func (obj *ArObj) ClearBitFlags(mutArg uint32) error {
	if mutArg == 0 || obj.fieldFlags & ^mutArg == obj.fieldFlags {
		return nil
	}

	data := iproto.PackUint32([]byte{}, uint32(^mutArg), iproto.ModeDefault)

	obj.BaseField.UpdateOps = append(obj.BaseField.UpdateOps, octopus.Ops{Field: 4, Op: octopus.OpAnd, Value: data})
	obj.fieldFlags &= ^mutArg

	return nil
}

func selectBox(ctx context.Context, indexnum uint32, keysPacked [][][]byte, limiter activerecord.SelectorLimiter) ([]*ArObj, error) {
	logger := activerecord.Logger()
	ctx = logger.SetLoggerValueToContext(ctx, activerecord.ValueLogPrefix{"limiter": limiter.String()})
	metricTimer := activerecord.Metric().Timer("octopus", "ArObj")
	metricStatCnt := activerecord.Metric().StatCount("octopus", "ArObj")
	metricErrCnt := activerecord.Metric().ErrorCount("octopus", "ArObj")

	w := octopus.PackSelect(namespace, indexnum, limiter.Offset(), limiter.Limit(), keysPacked)

	metricTimer.Timing(ctx, "select_pack")
	metricStatCnt.Inc(ctx, "select_keys", float64(len(keysPacked)))

	logger.Debug(ctx, fmt.Sprintf("Select packed tuple: '% X'", w))

	connection, err := octopus.Box(ctx, 0, activerecord.ReplicaOrMasterInstanceType)
	if err != nil {
		metricErrCnt.Inc(ctx, "select_preparebox", 1)
		logger.Error(ctx, fmt.Sprintf("Error get box '%s'", err))

		return nil, err
	}

	respBytes, errCall := connection.Call(ctx, octopus.RequestTypeSelect, w)
	if errCall != nil {
		metricErrCnt.Inc(ctx, "select_box", 1)
		logger.Error(ctx, "Error select from box", errCall, connection.Info())

		return nil, errCall
	}

	metricTimer.Timing(ctx, "select_box")

	logger.Debug(ctx, fmt.Sprintf("Response from box '%X'", respBytes))

	tuplesData, err := octopus.ProcessResp(respBytes, 0)
	if err != nil {
		metricErrCnt.Inc(ctx, "select_resp", 1)
		logger.Error(ctx, "Error parse response: ", err)

		return nil, err
	}

	metricTimer.Timing(ctx, "select_process")
	metricStatCnt.Inc(ctx, "select_tuples_res", float64(len(tuplesData)))

	nps, err := NewFromBox(ctx, tuplesData)
	if err != nil {
		metricErrCnt.Inc(ctx, "select_preparebox", 1)
		logger.Error(ctx, "Error in response: ", err)

		return nil, err
	}

	metricTimer.Timing(ctx, "select_newobj")

	if limiter.FullfillWarn() && len(nps) == int(limiter.Limit()) {
		logger.Warn(ctx, "Select limit reached. Result may less than db records.")
	}

	mode, ok := connection.InstanceMode().(octopus.ServerModeType)
	if !ok || activerecord.ServerModeType(mode) == activerecord.ModeReplica {
		if !ok {
			logger.Error(ctx, "Invalid server mode type: %T", connection.InstanceMode())
		}

		for npNum := range nps {
			nps[npNum].IsReplica = true
			nps[npNum].Readonly = true
		}
	}

	logger.Debug(ctx, "Success select")

	metricTimer.Finish(ctx, "select")

	return nps, nil
}

// indexes

func (obj *ArObj) Primary() int32 {

	return obj.GetID()
}

func SelectByPrimary(ctx context.Context, pk int32) (*ArObj, error) {
	return SelectByID(ctx, pk)
}

func PackKeyIndexID(ctx context.Context, keys []int32) ([][][]byte, error) {
	keysPacked := [][][]byte{}

	for _, key := range keys {
		keysField := [][]byte{}
		keysField = append(keysField, iproto.PackUint32([]byte{}, uint32(key), iproto.ModeDefault))
		keysPacked = append(keysPacked, keysField)
	}

	return keysPacked, nil
}
func UnpackKeyIndexID(packedKeys [][][]byte) ([]int32, error) {
	ret := []int32{}

	for _, packedKey := range packedKeys {

		newIField, err := UnpackID(bytes.NewReader(packedKey[0]))
		if err != nil {
			return nil, fmt.Errorf("can't unpack index: %s", err)
		}

		ret = append(ret, newIField)
	}

	return ret, nil
}

/*
		keysPacked := [][][]byte{}

		for _, key := range keys {
			keysField := [][]byte{}
			keysField = append(keysField, iproto.PackUint32([]byte{}, uint32(key), iproto.ModeDefault))
			keysPacked = append(keysPacked, keysField)
		}

		return keysPacked, nil
	}
*/
func SelectByIDs(ctx context.Context, keys []int32) ([]*ArObj, error) {
	ctx = activerecord.Logger().SetLoggerValueToContext(ctx, map[string]interface{}{"SelectByIDs": keys, "Repo": "ArObj"})

	keysPacked, err := PackKeyIndexID(ctx, keys)
	if err != nil {
		return nil, fmt.Errorf("can't pack index key: %s", err)
	}

	limiter := activerecord.EmptyLimiter()

	res, err := selectBox(ctx, 0, keysPacked, limiter)
	if err != nil {
		return res, err
	}

	activerecord.Logger().CollectQueries(ctx, SelectByIDMockerLogger(keys, ArObjList(res)))

	return res, err
}

func SelectByID(ctx context.Context, key int32) (*ArObj, error) {
	selected, err := SelectByIDs(ctx, []int32{key})
	if err != nil {
		return nil, err
	}

	if len(selected) > 0 {
		if len(selected) > 1 {
			activerecord.Logger().Error(ctx, "ArObj", "More than one tuple for uniq key ID '%s': %d", key, len(selected))
		}

		return selected[0], nil
	}

	return nil, nil
}
func PackKeyIndexType(ctx context.Context, keys []string) ([][][]byte, error) {
	keysPacked := [][][]byte{}

	for _, key := range keys {
		keysField := [][]byte{}
		keysField = append(keysField, octopus.PackString([]byte{}, key, iproto.ModeDefault))
		keysPacked = append(keysPacked, keysField)
	}

	return keysPacked, nil
}
func UnpackKeyIndexType(packedKeys [][][]byte) ([]string, error) {
	ret := []string{}

	for _, packedKey := range packedKeys {

		newIField, err := UnpackType(bytes.NewReader(packedKey[1]))
		if err != nil {
			return nil, fmt.Errorf("can't unpack index: %s", err)
		}

		ret = append(ret, newIField)
	}

	return ret, nil
}

/*
		keysPacked := [][][]byte{}

		for _, key := range keys {
			keysField := [][]byte{}
			keysField = append(keysField, octopus.PackString([]byte{}, key, iproto.ModeDefault))
			keysPacked = append(keysPacked, keysField)
		}

		return keysPacked, nil
	}
*/
func SelectByTypes(ctx context.Context, keys []string, limiter activerecord.SelectorLimiter) ([]*ArObj, error) {
	ctx = activerecord.Logger().SetLoggerValueToContext(ctx, map[string]interface{}{"SelectByTypes": keys, "Repo": "ArObj"})

	keysPacked, err := PackKeyIndexType(ctx, keys)
	if err != nil {
		return nil, fmt.Errorf("can't pack index key: %s", err)
	}

	res, err := selectBox(ctx, 1, keysPacked, limiter)
	if err != nil {
		return res, err
	}

	activerecord.Logger().CollectQueries(ctx, SelectByTypeMockerLogger(keys, ArObjList(res), limiter))

	return res, err
}

func SelectByType(ctx context.Context, key string, limiter activerecord.SelectorLimiter) ([]*ArObj, error) {
	selected, err := SelectByTypes(ctx, []string{key}, limiter)
	if err != nil {
		return nil, err
	}

	return selected, nil
}

type TypeIDIndexType struct {
	Type string

	ID int32
}

func PackKeyIndexTypeID(ctx context.Context, keys []TypeIDIndexType) ([][][]byte, error) {
	keysPacked := [][][]byte{}

	for _, key := range keys {
		keysField := [][]byte{}

		keysField = append(keysField, octopus.PackString([]byte{}, key.Type, iproto.ModeDefault))

		keysField = append(keysField, iproto.PackUint32([]byte{}, uint32(key.ID), iproto.ModeDefault))
		keysPacked = append(keysPacked, keysField)
	}

	return keysPacked, nil
}
func UnpackKeyIndexTypeID(packedKeys [][][]byte) ([]TypeIDIndexType, error) {
	ret := []TypeIDIndexType{}

	for _, packedKey := range packedKeys {

		newIField := TypeIDIndexType{}

		var err error
		newIField.Type, err = UnpackType(bytes.NewReader(packedKey[0]))
		if err != nil {
			return nil, fmt.Errorf("can't unpack index: %s", err)
		}

		newIField.ID, err = UnpackID(bytes.NewReader(packedKey[1]))
		if err != nil {
			return nil, fmt.Errorf("can't unpack index: %s", err)
		}

		ret = append(ret, newIField)
	}

	return ret, nil
}

/*
		keysPacked := [][][]byte{}

		for _, key := range keys {
			keysField := [][]byte{}



			keysField = append(keysField, octopus.PackString([]byte{}, key.Type, iproto.ModeDefault))



			keysField = append(keysField, iproto.PackUint32([]byte{}, uint32(key.ID), iproto.ModeDefault))
				keysPacked = append(keysPacked, keysField)
		}

		return keysPacked, nil
	}
*/
func SelectByTypeIDs(ctx context.Context, keys []TypeIDIndexType) ([]*ArObj, error) {
	ctx = activerecord.Logger().SetLoggerValueToContext(ctx, map[string]interface{}{"SelectByTypeIDs": keys, "Repo": "ArObj"})

	keysPacked, err := PackKeyIndexTypeID(ctx, keys)
	if err != nil {
		return nil, fmt.Errorf("can't pack index key: %s", err)
	}

	limiter := activerecord.EmptyLimiter()

	res, err := selectBox(ctx, 2, keysPacked, limiter)
	if err != nil {
		return res, err
	}

	activerecord.Logger().CollectQueries(ctx, SelectByTypeIDMockerLogger(keys, ArObjList(res)))

	return res, err
}

func SelectByTypeID(ctx context.Context, key TypeIDIndexType) (*ArObj, error) {
	selected, err := SelectByTypeIDs(ctx, []TypeIDIndexType{key})
	if err != nil {
		return nil, err
	}

	if len(selected) > 0 {
		if len(selected) > 1 {
			activerecord.Logger().Error(ctx, "ArObj", "More than one tuple for uniq key ID '%s': %d", key, len(selected))
		}

		return selected[0], nil
	}

	return nil, nil
}
func PackKeyIndexTypePart(ctx context.Context, keys []string) ([][][]byte, error) {
	keysPacked := [][][]byte{}

	for _, key := range keys {
		keysField := [][]byte{}
		keysField = append(keysField, octopus.PackString([]byte{}, key, iproto.ModeDefault))
		keysPacked = append(keysPacked, keysField)
	}

	return keysPacked, nil
}
func UnpackKeyIndexTypePart(packedKeys [][][]byte) ([]string, error) {
	ret := []string{}

	for _, packedKey := range packedKeys {

		newIField, err := UnpackType(bytes.NewReader(packedKey[3]))
		if err != nil {
			return nil, fmt.Errorf("can't unpack index: %s", err)
		}

		ret = append(ret, newIField)
	}

	return ret, nil
}

/*
		keysPacked := [][][]byte{}

		for _, key := range keys {
			keysField := [][]byte{}
			keysField = append(keysField, octopus.PackString([]byte{}, key, iproto.ModeDefault))
			keysPacked = append(keysPacked, keysField)
		}

		return keysPacked, nil
	}
*/
func SelectByTypeParts(ctx context.Context, keys []string, limiter activerecord.SelectorLimiter) ([]*ArObj, error) {
	ctx = activerecord.Logger().SetLoggerValueToContext(ctx, map[string]interface{}{"SelectByTypeParts": keys, "Repo": "ArObj"})

	keysPacked, err := PackKeyIndexTypePart(ctx, keys)
	if err != nil {
		return nil, fmt.Errorf("can't pack index key: %s", err)
	}

	res, err := selectBox(ctx, 2, keysPacked, limiter)
	if err != nil {
		return res, err
	}

	activerecord.Logger().CollectQueries(ctx, SelectByTypePartMockerLogger(keys, ArObjList(res), limiter))

	return res, err
}

func SelectByTypePart(ctx context.Context, key string, limiter activerecord.SelectorLimiter) ([]*ArObj, error) {
	selected, err := SelectByTypeParts(ctx, []string{key}, limiter)
	if err != nil {
		return nil, err
	}

	return selected, nil
}

// end indexes

func (obj *ArObj) Equal(anotherObjI any) bool {
	anotherObj, ok := anotherObjI.(*ArObj)
	if !ok {
		return false
	}

	var dataObj []byte
	var dataAnotherObj []byte
	var err error
	dataObj, err = packID([]byte{}, obj.GetID())
	if err != nil {
		return false
	}

	dataAnotherObj, err = packID([]byte{}, anotherObj.GetID())
	if err != nil {
		return false
	}

	if string(dataObj) != string(dataAnotherObj) {
		return false
	}

	dataObj, err = packName([]byte{}, obj.GetName())
	if err != nil {
		return false
	}

	dataAnotherObj, err = packName([]byte{}, anotherObj.GetName())
	if err != nil {
		return false
	}

	if string(dataObj) != string(dataAnotherObj) {
		return false
	}

	dataObj, err = packAnotherID([]byte{}, obj.GetAnotherID())
	if err != nil {
		return false
	}

	dataAnotherObj, err = packAnotherID([]byte{}, anotherObj.GetAnotherID())
	if err != nil {
		return false
	}

	if string(dataObj) != string(dataAnotherObj) {
		return false
	}

	dataObj, err = packType([]byte{}, obj.GetType())
	if err != nil {
		return false
	}

	dataAnotherObj, err = packType([]byte{}, anotherObj.GetType())
	if err != nil {
		return false
	}

	if string(dataObj) != string(dataAnotherObj) {
		return false
	}

	dataObj, err = packFlags([]byte{}, obj.GetFlags())
	if err != nil {
		return false
	}

	dataAnotherObj, err = packFlags([]byte{}, anotherObj.GetFlags())
	if err != nil {
		return false
	}

	if string(dataObj) != string(dataAnotherObj) {
		return false
	}

	return true
}

func (obj *ArObj) PrimaryString() string {
	ret := []string{
		strconv.FormatInt(int64(obj.GetID()), 10),
	}

	return strings.Join(ret, ", ")
}

func (obj *ArObj) packPk() ([][]byte, error) {
	packedPk := [][]byte{}

	var (
		data []byte
		err  error
	)

	data, err = packID([]byte{}, obj.GetID())
	if err != nil {
		return [][]byte{}, err
	}

	packedPk = append(packedPk, data)

	return packedPk, nil
}

func (obj *ArObj) Delete(ctx context.Context) error {
	logger := activerecord.Logger()
	metricTimer := activerecord.Metric().Timer("octopus", "ArObj")
	metricStatCnt := activerecord.Metric().StatCount("octopus", "ArObj")
	metricErrCnt := activerecord.Metric().ErrorCount("octopus", "ArObj")

	metricStatCnt.Inc(ctx, "delete_request", 1)

	if !obj.BaseField.Exists {
		return fmt.Errorf("can't delete not exists object")
	}

	pk, err := obj.packPk()
	if err != nil {
		metricErrCnt.Inc(ctx, "delete_pack", 1)
		return fmt.Errorf("error delete: %w", err)
	}

	w := octopus.PackDelete(namespace, pk)
	log.Printf("Delete packed tuple: '%X'\n", w)

	connection, err := octopus.Box(ctx, 0, activerecord.MasterInstanceType)
	if err != nil {
		metricErrCnt.Inc(ctx, "delete_preparebox", 1)
		logger.Error(ctx, "ArObj", obj.PrimaryString(), fmt.Sprintf("Error get box '%s'", err))

		return err
	}

	respBytes, errCall := connection.Call(ctx, octopus.RequestTypeDelete, w)
	if errCall != nil {
		metricErrCnt.Inc(ctx, "delete_box", 1)
		logger.Error(ctx, "ArObj", obj.PrimaryString(), "Error delete from box", errCall, connection.Info())

		return errCall
	}

	metricTimer.Timing(ctx, "delete_box")

	logger.Debug(ctx, "ArObj", obj.PrimaryString(), fmt.Sprintf("Response from box '% X'", respBytes))

	_, err = octopus.ProcessResp(respBytes, octopus.NeedRespFlag|octopus.UniqRespFlag)
	if err != nil {
		metricErrCnt.Inc(ctx, "delete_resp", 1)
		logger.Error(ctx, "ArObj", obj.PrimaryString(), "Error parse response: ", err)

		return err
	}

	metricStatCnt.Inc(ctx, "delete_success", 1)

	obj.BaseField.Exists = false
	obj.BaseField.UpdateOps = []octopus.Ops{}

	logger.Debug(ctx, "ArObj", obj.PrimaryString(), "Success delete")

	metricTimer.Finish(ctx, "delete")

	return nil
}

func (obj *ArObj) Update(ctx context.Context) error {
	logger := activerecord.Logger()
	metricTimer := activerecord.Metric().Timer("octopus", "ArObj")
	metricStatCnt := activerecord.Metric().StatCount("octopus", "ArObj")
	metricErrCnt := activerecord.Metric().ErrorCount("octopus", "ArObj")

	metricStatCnt.Inc(ctx, "update_request", 1)

	if !obj.BaseField.Exists {
		metricErrCnt.Inc(ctx, "update_notexists", 1)
		return fmt.Errorf("can't update not exists object")
	}

	if obj.BaseField.Repaired {
		metricStatCnt.Inc(ctx, "update_repaired", 1)
		logger.Debug(ctx, "", obj.PrimaryString(), "Flag 'Repaired' is true! Insert instead Update")

		return obj.Replace(ctx)
	}

	if len(obj.BaseField.UpdateOps) == 0 {
		metricStatCnt.Inc(ctx, "update_empty", 1)
		logger.Debug(ctx, "", obj.PrimaryString(), "Empty update")

		return nil
	}

	pk, err := obj.packPk()
	if err != nil {
		metricErrCnt.Inc(ctx, "update_packpk", 1)
		return fmt.Errorf("error update: %w", err)
	}

	w := octopus.PackUpdate(namespace, pk, obj.BaseField.UpdateOps)

	log.Printf("Update packed tuple: '%X'\n", w)

	connection, err := octopus.Box(ctx, 0, activerecord.MasterInstanceType)
	if err != nil {
		metricErrCnt.Inc(ctx, "update_preparebox", 1)
		logger.Error(ctx, "ArObj", obj.PrimaryString(), fmt.Sprintf("Error get box '%s'", err))
		return err
	}

	respBytes, errCall := connection.Call(ctx, octopus.RequestTypeUpdate, w)
	if errCall != nil {
		metricErrCnt.Inc(ctx, "update_box", 1)
		logger.Error(ctx, "ArObj", obj.PrimaryString(), "Error update ia a box", errCall, connection.Info())
		return errCall
	}

	metricTimer.Timing(ctx, "update_box")

	logger.Debug(ctx, "ArObj", obj.PrimaryString(), fmt.Sprintf("Response from box '%X'", respBytes))

	_, err = octopus.ProcessResp(respBytes, octopus.NeedRespFlag|octopus.UniqRespFlag)
	if err != nil {
		metricErrCnt.Inc(ctx, "update_resp", 1)
		logger.Error(ctx, "ArObj", obj.PrimaryString(), "Error parse response: ", err)
		return err
	}

	obj.BaseField.UpdateOps = []octopus.Ops{}

	logger.Debug(ctx, "ArObj", obj.PrimaryString(), "Success update")

	metricStatCnt.Inc(ctx, "update_success", 1)
	metricTimer.Finish(ctx, "update")

	return nil
}

func (obj *ArObj) Insert(ctx context.Context) error {
	metricStatCnt := activerecord.Metric().StatCount("octopus", "ArObj")
	metricErrCnt := activerecord.Metric().ErrorCount("octopus", "ArObj")

	metricStatCnt.Inc(ctx, "insert_request", 1)

	if obj.BaseField.Exists {
		metricErrCnt.Inc(ctx, "insert_exists", 1)
		return fmt.Errorf("can't insert already exists object")
	}

	err := obj.insertReplace(ctx, octopus.InsertModeInsert)

	if err == nil {
		metricStatCnt.Inc(ctx, "insert_success", 1)
	}

	return err
}

func (obj *ArObj) Replace(ctx context.Context) error {
	metricStatCnt := activerecord.Metric().StatCount("octopus", "ArObj")
	metricErrCnt := activerecord.Metric().ErrorCount("octopus", "ArObj")

	metricStatCnt.Inc(ctx, "replace_request", 1)

	if !obj.BaseField.Exists {
		metricErrCnt.Inc(ctx, "replace_notexists", 1)
		return fmt.Errorf("can't replace not exists object")
	}

	err := obj.insertReplace(ctx, octopus.InsertModeReplace)

	if err == nil {
		metricStatCnt.Inc(ctx, "replace_success", 1)
	}

	return err
}

func (obj *ArObj) InsertOrReplace(ctx context.Context) error {
	metricStatCnt := activerecord.Metric().StatCount("octopus", "ArObj")

	metricStatCnt.Inc(ctx, "insertorreplace_request", 1)

	err := obj.insertReplace(ctx, octopus.InsertModeInserOrReplace)

	if err == nil {
		metricStatCnt.Inc(ctx, "insertorreplace_success", 1)
	}

	return err
}

func (obj *ArObj) insertReplace(ctx context.Context, insertMode octopus.InsertMode) error {
	var (
		err   error
		tuple [][]byte
		data  []byte
	)

	metricTimer := activerecord.Metric().Timer("octopus", "ArObj")
	metricErrCnt := activerecord.Metric().ErrorCount("octopus", "ArObj")

	data, err = packID([]byte{}, obj.GetID())
	if err != nil {
		metricErrCnt.Inc(ctx, "insertreplace_packfield", 1)
		return err
	}

	tuple = append(tuple, data)

	data, err = packName([]byte{}, obj.GetName())
	if err != nil {
		metricErrCnt.Inc(ctx, "insertreplace_packfield", 1)
		return err
	}

	tuple = append(tuple, data)

	data, err = packAnotherID([]byte{}, obj.GetAnotherID())
	if err != nil {
		metricErrCnt.Inc(ctx, "insertreplace_packfield", 1)
		return err
	}

	tuple = append(tuple, data)

	data, err = packType([]byte{}, obj.GetType())
	if err != nil {
		metricErrCnt.Inc(ctx, "insertreplace_packfield", 1)
		return err
	}

	tuple = append(tuple, data)

	data, err = packFlags([]byte{}, obj.GetFlags())
	if err != nil {
		metricErrCnt.Inc(ctx, "insertreplace_packfield", 1)
		return err
	}

	tuple = append(tuple, data)

	metricTimer.Timing(ctx, "insertreplace_packtuple")

	if len(obj.BaseField.ExtraFields) > 0 {
		tuple = append(tuple, obj.BaseField.ExtraFields...)
	}

	w := octopus.PackInsertReplace(namespace, insertMode, tuple)
	logger := activerecord.Logger()

	metricTimer.Timing(ctx, "insertreplace_pack")
	logger.Trace(ctx, "ArObj", obj.PrimaryString(), fmt.Sprintf("Insert packed tuple: '%X'", w))

	connection, err := octopus.Box(ctx, 0, activerecord.MasterInstanceType)
	if err != nil {
		metricErrCnt.Inc(ctx, "insertreplace_preparebox", 1)
		logger.Error(ctx, "ArObj", obj.PrimaryString(), fmt.Sprintf("Error get box '%s'", err))

		return err
	}

	respBytes, errCall := connection.Call(ctx, octopus.RequestTypeInsert, w)
	if errCall != nil {
		metricErrCnt.Inc(ctx, "insertreplace_box", 1)
		logger.Error(ctx, "ArObj", obj.PrimaryString(), "Error insert into box", errCall, connection.Info())

		return errCall
	}

	metricTimer.Timing(ctx, "insertreplace_box")

	logger.Trace(ctx, "ArObj", obj.PrimaryString(), fmt.Sprintf("Response from box '%X'", respBytes))

	tuplesData, err := octopus.ProcessResp(respBytes, octopus.NeedRespFlag|octopus.UniqRespFlag)
	if err != nil {
		metricErrCnt.Inc(ctx, "insertreplace_prespreparebox", 1)
		logger.Error(ctx, "ArObj", obj.PrimaryString(), "Error parse response: ", err)

		return err
	}

	_, err = NewFromBox(ctx, tuplesData)
	if err != nil {
		metricErrCnt.Inc(ctx, "insertreplace_obj", 1)
		logger.Error(ctx, "ArObj", obj.PrimaryString(), "Error in response: ", err)

		return err
	}

	obj.BaseField.Exists = true
	obj.BaseField.UpdateOps = []octopus.Ops{}
	obj.BaseField.Repaired = false

	logger.Debug(ctx, "ArObj", obj.PrimaryString(), "Success insert")

	metricTimer.Finish(ctx, "insertreplace")

	return nil
}
