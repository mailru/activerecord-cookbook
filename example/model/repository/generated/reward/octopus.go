// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.5.3-18-g3247b15 (Commit: 3247b15e)
package reward

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"strings"

	"github.com/mailru/activerecord-cookbook/example/model/ds"
	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/iproto/iproto"
	"github.com/mailru/activerecord/pkg/octopus"
	serializerExtra "github.com/mailru/activerecord/pkg/serializer"
)

type Reward struct {
	octopus.BaseField
	fieldCode        string
	fieldServices    *ds.Services
	fieldPartner     string
	fieldExtra       *ds.Extra
	fieldFlags       map[string]interface{}
	fieldUnlocked    ds.ServiceUnlocked
	fieldDescription *string
}

type RewardList []*Reward

const (
	namespace uint32 = 24
	cntFields uint32 = 7
)

func New(ctx context.Context) *Reward {
	newObj := Reward{}
	newObj.BaseField.UpdateOps = []octopus.Ops{}
	newObj.BaseField.ExtraFields = [][]byte{}
	newObj.BaseField.Objects = map[string][]octopus.ModelStruct{}

	return &newObj
}

func TupleToStruct(ctx context.Context, tuple octopus.TupleData) (*Reward, error) {
	np := New(ctx)

	valCode, err := UnpackCode(bytes.NewReader(tuple.Data[0]))
	if err != nil {
		return nil, err
	}

	np.SetCode(valCode)
	valServices, err := UnpackServices(bytes.NewReader(tuple.Data[1]))
	if err != nil {
		return nil, err
	}

	np.SetServices(valServices)
	valPartner, err := UnpackPartner(bytes.NewReader(tuple.Data[2]))
	if err != nil {
		return nil, err
	}

	np.SetPartner(valPartner)
	valExtra, err := UnpackExtra(bytes.NewReader(tuple.Data[3]))
	if err != nil {
		return nil, err
	}

	np.SetExtra(valExtra)
	valFlags, err := UnpackFlags(bytes.NewReader(tuple.Data[4]))
	if err != nil {
		return nil, err
	}

	np.SetFlags(valFlags)
	valUnlocked, err := UnpackUnlocked(bytes.NewReader(tuple.Data[5]))
	if err != nil {
		return nil, err
	}

	np.SetUnlocked(valUnlocked)
	valDescription, err := UnpackDescription(bytes.NewReader(tuple.Data[6]))
	if err != nil {
		return nil, err
	}

	np.SetDescription(valDescription)

	np.BaseField.Exists = true
	np.BaseField.UpdateOps = []octopus.Ops{}

	if tuple.Cnt > cntFields {
		logger := activerecord.Logger()

		logger.Warn(ctx, "Reward", np.PrimaryString(), "Extra fields")

		np.BaseField.ExtraFields = tuple.Data[cntFields:]
	}

	return np, nil
}

func NewFromBox(ctx context.Context, tuples []octopus.TupleData) ([]*Reward, error) {
	logger := activerecord.Logger()

	logger.Debug(ctx, "Reward", fmt.Sprintf("Cnt tuples %d", len(tuples)))

	ret := make([]*Reward, 0, len(tuples))

	for num, tuple := range tuples {
		var repaired bool

		if tuple.Cnt < cntFields {
			return nil, fmt.Errorf("not enought selected fields %d in response tuple: %d but expected %d fields", tuple.Cnt, num, cntFields)
		}

		np, err := TupleToStruct(ctx, tuple)
		if err != nil {
			logger.Error(ctx, "Reward", fmt.Sprintf("error unpack tuple %s", err))
			return nil, err
		}

		np.BaseField.Repaired = repaired
		ret = append(ret, np)
	}

	return ret, nil
}
func packCode(w []byte, Code string) ([]byte, error) {
	pvar := Code

	return octopus.PackString(w, pvar, iproto.ModeDefault), nil
}

func UnpackCode(r *bytes.Reader) (ret string, errRet error) {
	var Code string

	err := octopus.UnpackString(r, &Code, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field Code in tuple: '%w'", err)
		return
	}

	bvar := Code

	svar := bvar

	return svar, nil
}

func (obj *Reward) GetCode() string {
	return obj.fieldCode
}

func (obj *Reward) SetCode(Code string) error {
	if obj.BaseField.Exists {
		return fmt.Errorf("can't modify field included in primary key")
	}

	data, err := packCode([]byte{}, Code)
	if err != nil {
		return err
	}

	logger := activerecord.Logger()

	logger.Warn(context.TODO(), "Reward", obj.PrimaryString(), fmt.Sprintf("Size for field 'Code' not set. Cur field size: %d. Object: 'Reward'", len(data)))

	obj.BaseField.UpdateOps = append(obj.BaseField.UpdateOps, octopus.Ops{Field: 0, Op: octopus.OpSet, Value: data})
	obj.fieldCode = Code

	return nil
}

func packServices(w []byte, Services *ds.Services) ([]byte, error) {
	pvar, err := serializerExtra.MapstructureMarshal(Services)
	if err != nil {
		return nil, fmt.Errorf("error marshal field Services: %w", err)
	}

	return octopus.PackString(w, pvar, iproto.ModeDefault), nil
}

func UnpackServices(r *bytes.Reader) (ret *ds.Services, errRet error) {
	var Services string

	err := octopus.UnpackString(r, &Services, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field Services in tuple: '%w'", err)
		return
	}

	bvar := Services

	var svar ds.Services

	err = serializerExtra.MapstructureUnmarshal(bvar, &svar)
	if err != nil {
		errRet = fmt.Errorf("error unmarshal field Services: %w", err)
		return
	}

	return &svar, nil
}

func (obj *Reward) GetServices() *ds.Services {
	return obj.fieldServices
}

func (obj *Reward) SetServices(Services *ds.Services) error {
	data, err := packServices([]byte{}, Services)
	if err != nil {
		return err
	}

	logger := activerecord.Logger()

	logger.Warn(context.TODO(), "Reward", obj.PrimaryString(), fmt.Sprintf("Size for field 'Services' not set. Cur field size: %d. Object: 'Reward'", len(data)))

	obj.BaseField.UpdateOps = append(obj.BaseField.UpdateOps, octopus.Ops{Field: 1, Op: octopus.OpSet, Value: data})
	obj.fieldServices = Services

	return nil
}

func packPartner(w []byte, Partner string) ([]byte, error) {
	pvar := Partner

	return octopus.PackString(w, pvar, iproto.ModeDefault), nil
}

func UnpackPartner(r *bytes.Reader) (ret string, errRet error) {
	var Partner string

	err := octopus.UnpackString(r, &Partner, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field Partner in tuple: '%w'", err)
		return
	}

	bvar := Partner

	svar := bvar

	return svar, nil
}

func (obj *Reward) GetPartner() string {
	return obj.fieldPartner
}

func (obj *Reward) SetPartner(Partner string) error {
	data, err := packPartner([]byte{}, Partner)
	if err != nil {
		return err
	}

	logger := activerecord.Logger()

	logger.Warn(context.TODO(), "Reward", obj.PrimaryString(), fmt.Sprintf("Size for field 'Partner' not set. Cur field size: %d. Object: 'Reward'", len(data)))

	obj.BaseField.UpdateOps = append(obj.BaseField.UpdateOps, octopus.Ops{Field: 2, Op: octopus.OpSet, Value: data})
	obj.fieldPartner = Partner

	return nil
}

func packExtra(w []byte, Extra *ds.Extra) ([]byte, error) {
	pvar, err := serializerExtra.MapstructureMarshal(Extra)
	if err != nil {
		return nil, fmt.Errorf("error marshal field Extra: %w", err)
	}

	return octopus.PackString(w, pvar, iproto.ModeDefault), nil
}

func UnpackExtra(r *bytes.Reader) (ret *ds.Extra, errRet error) {
	var Extra string

	err := octopus.UnpackString(r, &Extra, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field Extra in tuple: '%w'", err)
		return
	}

	bvar := Extra

	var svar ds.Extra

	err = serializerExtra.MapstructureUnmarshal(bvar, &svar)
	if err != nil {
		errRet = fmt.Errorf("error unmarshal field Extra: %w", err)
		return
	}

	return &svar, nil
}

func (obj *Reward) GetExtra() *ds.Extra {
	return obj.fieldExtra
}

func (obj *Reward) SetExtra(Extra *ds.Extra) error {
	data, err := packExtra([]byte{}, Extra)
	if err != nil {
		return err
	}

	logger := activerecord.Logger()

	logger.Warn(context.TODO(), "Reward", obj.PrimaryString(), fmt.Sprintf("Size for field 'Extra' not set. Cur field size: %d. Object: 'Reward'", len(data)))

	obj.BaseField.UpdateOps = append(obj.BaseField.UpdateOps, octopus.Ops{Field: 3, Op: octopus.OpSet, Value: data})
	obj.fieldExtra = Extra

	return nil
}

func packFlags(w []byte, Flags map[string]interface{}) ([]byte, error) {
	pvar, err := serializerExtra.JSONMarshal(Flags)
	if err != nil {
		return nil, fmt.Errorf("error marshal field Flags: %w", err)
	}

	return octopus.PackString(w, pvar, iproto.ModeDefault), nil
}

func UnpackFlags(r *bytes.Reader) (ret map[string]interface{}, errRet error) {
	var Flags string

	err := octopus.UnpackString(r, &Flags, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field Flags in tuple: '%w'", err)
		return
	}

	bvar := Flags

	var svar map[string]interface{}

	err = serializerExtra.JSONUnmarshal(bvar, &svar)
	if err != nil {
		errRet = fmt.Errorf("error unmarshal field Flags: %w", err)
		return
	}

	return svar, nil
}

func (obj *Reward) GetFlags() map[string]interface{} {
	return obj.fieldFlags
}

func (obj *Reward) SetFlags(Flags map[string]interface{}) error {
	data, err := packFlags([]byte{}, Flags)
	if err != nil {
		return err
	}

	logger := activerecord.Logger()

	logger.Warn(context.TODO(), "Reward", obj.PrimaryString(), fmt.Sprintf("Size for field 'Flags' not set. Cur field size: %d. Object: 'Reward'", len(data)))

	obj.BaseField.UpdateOps = append(obj.BaseField.UpdateOps, octopus.Ops{Field: 4, Op: octopus.OpSet, Value: data})
	obj.fieldFlags = Flags

	return nil
}

func packUnlocked(w []byte, Unlocked ds.ServiceUnlocked) ([]byte, error) {
	pvar, err := serializerExtra.JSONMarshal(Unlocked)
	if err != nil {
		return nil, fmt.Errorf("error marshal field Unlocked: %w", err)
	}

	return octopus.PackString(w, pvar, iproto.ModeDefault), nil
}

func UnpackUnlocked(r *bytes.Reader) (ret ds.ServiceUnlocked, errRet error) {
	var Unlocked string

	err := octopus.UnpackString(r, &Unlocked, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field Unlocked in tuple: '%w'", err)
		return
	}

	bvar := Unlocked

	var svar ds.ServiceUnlocked

	err = serializerExtra.JSONUnmarshal(bvar, &svar)
	if err != nil {
		errRet = fmt.Errorf("error unmarshal field Unlocked: %w", err)
		return
	}

	return svar, nil
}

func (obj *Reward) GetUnlocked() ds.ServiceUnlocked {
	return obj.fieldUnlocked
}

func (obj *Reward) SetUnlocked(Unlocked ds.ServiceUnlocked) error {
	data, err := packUnlocked([]byte{}, Unlocked)
	if err != nil {
		return err
	}

	logger := activerecord.Logger()

	logger.Warn(context.TODO(), "Reward", obj.PrimaryString(), fmt.Sprintf("Size for field 'Unlocked' not set. Cur field size: %d. Object: 'Reward'", len(data)))

	obj.BaseField.UpdateOps = append(obj.BaseField.UpdateOps, octopus.Ops{Field: 5, Op: octopus.OpSet, Value: data})
	obj.fieldUnlocked = Unlocked

	return nil
}

func packDescription(w []byte, Description *string) ([]byte, error) {
	pvar, err := serializerExtra.JSONMarshal(Description)
	if err != nil {
		return nil, fmt.Errorf("error marshal field Description: %w", err)
	}

	return octopus.PackString(w, pvar, iproto.ModeDefault), nil
}

func UnpackDescription(r *bytes.Reader) (ret *string, errRet error) {
	var Description string

	err := octopus.UnpackString(r, &Description, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field Description in tuple: '%w'", err)
		return
	}

	bvar := Description

	var svar string

	err = serializerExtra.JSONUnmarshal(bvar, &svar)
	if err != nil {
		errRet = fmt.Errorf("error unmarshal field Description: %w", err)
		return
	}

	return &svar, nil
}

func (obj *Reward) GetDescription() *string {
	return obj.fieldDescription
}

func (obj *Reward) SetDescription(Description *string) error {
	data, err := packDescription([]byte{}, Description)
	if err != nil {
		return err
	}

	logger := activerecord.Logger()

	logger.Warn(context.TODO(), "Reward", obj.PrimaryString(), fmt.Sprintf("Size for field 'Description' not set. Cur field size: %d. Object: 'Reward'", len(data)))

	obj.BaseField.UpdateOps = append(obj.BaseField.UpdateOps, octopus.Ops{Field: 6, Op: octopus.OpSet, Value: data})
	obj.fieldDescription = Description

	return nil
}

func selectBox(ctx context.Context, indexnum uint32, keysPacked [][][]byte, limiter activerecord.SelectorLimiter) ([]*Reward, error) {
	logger := activerecord.Logger()
	ctx = logger.SetLoggerValueToContext(ctx, activerecord.ValueLogPrefix{"limiter": limiter.String()})
	metricTimer := activerecord.Metric().Timer("octopus", "Reward")
	metricStatCnt := activerecord.Metric().StatCount("octopus", "Reward")
	metricErrCnt := activerecord.Metric().ErrorCount("octopus", "Reward")

	w := octopus.PackSelect(namespace, indexnum, limiter.Offset(), limiter.Limit(), keysPacked)

	metricTimer.Timing(ctx, "select_pack")
	metricStatCnt.Inc(ctx, "select_keys", float64(len(keysPacked)))

	logger.Debug(ctx, fmt.Sprintf("Select packed tuple: '% X'", w))

	connection, err := octopus.Box(ctx, 0, activerecord.ReplicaOrMasterInstanceType, "arcfg", nil)
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

func (obj *Reward) Primary() string {

	return obj.GetCode()
}

func SelectByPrimary(ctx context.Context, pk string) (*Reward, error) {
	return SelectByCode(ctx, pk)
}

func PackKeyIndexCode(ctx context.Context, keys []string) ([][][]byte, error) {
	keysPacked := [][][]byte{}

	for _, key := range keys {
		keysField := [][]byte{}
		keysField = append(keysField, octopus.PackString([]byte{}, key, iproto.ModeDefault))
		keysPacked = append(keysPacked, keysField)
	}

	return keysPacked, nil
}
func UnpackKeyIndexCode(packedKeys [][][]byte) ([]string, error) {
	ret := []string{}

	for _, packedKey := range packedKeys {

		newIField, err := UnpackCode(bytes.NewReader(packedKey[0]))
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
func SelectByCodes(ctx context.Context, keys []string) ([]*Reward, error) {
	ctx = activerecord.Logger().SetLoggerValueToContext(ctx, map[string]interface{}{"SelectByCodes": keys, "Repo": "Reward"})

	keysPacked, err := PackKeyIndexCode(ctx, keys)
	if err != nil {
		return nil, fmt.Errorf("can't pack index key: %s", err)
	}

	limiter := activerecord.EmptyLimiter()

	res, err := selectBox(ctx, 0, keysPacked, limiter)
	if err != nil {
		return res, err
	}

	activerecord.Logger().CollectQueries(ctx, SelectByCodeMockerLogger(keys, RewardList(res)))

	return res, err
}

func SelectByCode(ctx context.Context, key string) (*Reward, error) {
	selected, err := SelectByCodes(ctx, []string{key})
	if err != nil {
		return nil, err
	}

	if len(selected) > 0 {
		if len(selected) > 1 {
			activerecord.Logger().Error(ctx, "Reward", "More than one tuple for uniq key ID '%s': %d", key, len(selected))
		}

		return selected[0], nil
	}

	return nil, nil
}
func PackKeyIndexPartner(ctx context.Context, keys []string) ([][][]byte, error) {
	keysPacked := [][][]byte{}

	for _, key := range keys {
		keysField := [][]byte{}
		keysField = append(keysField, octopus.PackString([]byte{}, key, iproto.ModeDefault))
		keysPacked = append(keysPacked, keysField)
	}

	return keysPacked, nil
}
func UnpackKeyIndexPartner(packedKeys [][][]byte) ([]string, error) {
	ret := []string{}

	for _, packedKey := range packedKeys {

		newIField, err := UnpackPartner(bytes.NewReader(packedKey[1]))
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
func SelectByPartners(ctx context.Context, keys []string, limiter activerecord.SelectorLimiter) ([]*Reward, error) {
	ctx = activerecord.Logger().SetLoggerValueToContext(ctx, map[string]interface{}{"SelectByPartners": keys, "Repo": "Reward"})

	keysPacked, err := PackKeyIndexPartner(ctx, keys)
	if err != nil {
		return nil, fmt.Errorf("can't pack index key: %s", err)
	}

	res, err := selectBox(ctx, 1, keysPacked, limiter)
	if err != nil {
		return res, err
	}

	activerecord.Logger().CollectQueries(ctx, SelectByPartnerMockerLogger(keys, RewardList(res), limiter))

	return res, err
}

func SelectByPartner(ctx context.Context, key string, limiter activerecord.SelectorLimiter) ([]*Reward, error) {
	selected, err := SelectByPartners(ctx, []string{key}, limiter)
	if err != nil {
		return nil, err
	}

	return selected, nil
}

// end indexes

func (obj *Reward) Equal(anotherObjI any) bool {
	anotherObj, ok := anotherObjI.(*Reward)
	if !ok {
		return false
	}

	var dataObj []byte
	var dataAnotherObj []byte
	var err error
	dataObj, err = packCode([]byte{}, obj.GetCode())
	if err != nil {
		return false
	}

	dataAnotherObj, err = packCode([]byte{}, anotherObj.GetCode())
	if err != nil {
		return false
	}

	if string(dataObj) != string(dataAnotherObj) {
		return false
	}

	dataObj, err = packServices([]byte{}, obj.GetServices())
	if err != nil {
		return false
	}

	dataAnotherObj, err = packServices([]byte{}, anotherObj.GetServices())
	if err != nil {
		return false
	}

	if string(dataObj) != string(dataAnotherObj) {
		return false
	}

	dataObj, err = packPartner([]byte{}, obj.GetPartner())
	if err != nil {
		return false
	}

	dataAnotherObj, err = packPartner([]byte{}, anotherObj.GetPartner())
	if err != nil {
		return false
	}

	if string(dataObj) != string(dataAnotherObj) {
		return false
	}

	dataObj, err = packExtra([]byte{}, obj.GetExtra())
	if err != nil {
		return false
	}

	dataAnotherObj, err = packExtra([]byte{}, anotherObj.GetExtra())
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

	dataObj, err = packUnlocked([]byte{}, obj.GetUnlocked())
	if err != nil {
		return false
	}

	dataAnotherObj, err = packUnlocked([]byte{}, anotherObj.GetUnlocked())
	if err != nil {
		return false
	}

	if string(dataObj) != string(dataAnotherObj) {
		return false
	}

	dataObj, err = packDescription([]byte{}, obj.GetDescription())
	if err != nil {
		return false
	}

	dataAnotherObj, err = packDescription([]byte{}, anotherObj.GetDescription())
	if err != nil {
		return false
	}

	if string(dataObj) != string(dataAnotherObj) {
		return false
	}

	return true
}

func (obj *Reward) PrimaryString() string {
	ret := []string{
		obj.GetCode(),
	}

	return strings.Join(ret, ", ")
}

func (obj *Reward) packPk() ([][]byte, error) {
	packedPk := [][]byte{}

	var (
		data []byte
		err  error
	)

	data, err = packCode([]byte{}, obj.GetCode())
	if err != nil {
		return [][]byte{}, err
	}

	packedPk = append(packedPk, data)

	return packedPk, nil
}

func (obj *Reward) Delete(ctx context.Context) error {
	logger := activerecord.Logger()
	metricTimer := activerecord.Metric().Timer("octopus", "Reward")
	metricStatCnt := activerecord.Metric().StatCount("octopus", "Reward")
	metricErrCnt := activerecord.Metric().ErrorCount("octopus", "Reward")

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

	connection, err := octopus.Box(ctx, 0, activerecord.MasterInstanceType, "arcfg", nil)
	if err != nil {
		metricErrCnt.Inc(ctx, "delete_preparebox", 1)
		logger.Error(ctx, "Reward", obj.PrimaryString(), fmt.Sprintf("Error get box '%s'", err))

		return err
	}

	respBytes, errCall := connection.Call(ctx, octopus.RequestTypeDelete, w)
	if errCall != nil {
		metricErrCnt.Inc(ctx, "delete_box", 1)
		logger.Error(ctx, "Reward", obj.PrimaryString(), "Error delete from box", errCall, connection.Info())

		return errCall
	}

	metricTimer.Timing(ctx, "delete_box")

	logger.Debug(ctx, "Reward", obj.PrimaryString(), fmt.Sprintf("Response from box '% X'", respBytes))

	_, err = octopus.ProcessResp(respBytes, octopus.NeedRespFlag|octopus.UniqRespFlag)
	if err != nil {
		metricErrCnt.Inc(ctx, "delete_resp", 1)
		logger.Error(ctx, "Reward", obj.PrimaryString(), "Error parse response: ", err)

		return err
	}

	metricStatCnt.Inc(ctx, "delete_success", 1)

	obj.BaseField.Exists = false
	obj.BaseField.UpdateOps = []octopus.Ops{}

	logger.Debug(ctx, "Reward", obj.PrimaryString(), "Success delete")

	metricTimer.Finish(ctx, "delete")

	return nil
}

func (obj *Reward) Update(ctx context.Context) error {
	logger := activerecord.Logger()
	metricTimer := activerecord.Metric().Timer("octopus", "Reward")
	metricStatCnt := activerecord.Metric().StatCount("octopus", "Reward")
	metricErrCnt := activerecord.Metric().ErrorCount("octopus", "Reward")

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

	connection, err := octopus.Box(ctx, 0, activerecord.MasterInstanceType, "arcfg", nil)
	if err != nil {
		metricErrCnt.Inc(ctx, "update_preparebox", 1)
		logger.Error(ctx, "Reward", obj.PrimaryString(), fmt.Sprintf("Error get box '%s'", err))
		return err
	}

	respBytes, errCall := connection.Call(ctx, octopus.RequestTypeUpdate, w)
	if errCall != nil {
		metricErrCnt.Inc(ctx, "update_box", 1)
		logger.Error(ctx, "Reward", obj.PrimaryString(), "Error update ia a box", errCall, connection.Info())
		return errCall
	}

	metricTimer.Timing(ctx, "update_box")

	logger.Debug(ctx, "Reward", obj.PrimaryString(), fmt.Sprintf("Response from box '%X'", respBytes))

	_, err = octopus.ProcessResp(respBytes, octopus.NeedRespFlag|octopus.UniqRespFlag)
	if err != nil {
		metricErrCnt.Inc(ctx, "update_resp", 1)
		logger.Error(ctx, "Reward", obj.PrimaryString(), "Error parse response: ", err)
		return err
	}

	obj.BaseField.UpdateOps = []octopus.Ops{}

	logger.Debug(ctx, "Reward", obj.PrimaryString(), "Success update")

	metricStatCnt.Inc(ctx, "update_success", 1)
	metricTimer.Finish(ctx, "update")

	return nil
}

func (obj *Reward) Insert(ctx context.Context) error {
	metricStatCnt := activerecord.Metric().StatCount("octopus", "Reward")
	metricErrCnt := activerecord.Metric().ErrorCount("octopus", "Reward")

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

func (obj *Reward) Replace(ctx context.Context) error {
	metricStatCnt := activerecord.Metric().StatCount("octopus", "Reward")
	metricErrCnt := activerecord.Metric().ErrorCount("octopus", "Reward")

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

func (obj *Reward) InsertOrReplace(ctx context.Context) error {
	metricStatCnt := activerecord.Metric().StatCount("octopus", "Reward")

	metricStatCnt.Inc(ctx, "insertorreplace_request", 1)

	err := obj.insertReplace(ctx, octopus.InsertModeInserOrReplace)

	if err == nil {
		metricStatCnt.Inc(ctx, "insertorreplace_success", 1)
	}

	return err
}

func (obj *Reward) insertReplace(ctx context.Context, insertMode octopus.InsertMode) error {
	var (
		err   error
		tuple [][]byte
		data  []byte
	)

	metricTimer := activerecord.Metric().Timer("octopus", "Reward")
	metricErrCnt := activerecord.Metric().ErrorCount("octopus", "Reward")

	data, err = packCode([]byte{}, obj.GetCode())
	if err != nil {
		metricErrCnt.Inc(ctx, "insertreplace_packfield", 1)
		return err
	}

	tuple = append(tuple, data)

	data, err = packServices([]byte{}, obj.GetServices())
	if err != nil {
		metricErrCnt.Inc(ctx, "insertreplace_packfield", 1)
		return err
	}

	tuple = append(tuple, data)

	data, err = packPartner([]byte{}, obj.GetPartner())
	if err != nil {
		metricErrCnt.Inc(ctx, "insertreplace_packfield", 1)
		return err
	}

	tuple = append(tuple, data)

	data, err = packExtra([]byte{}, obj.GetExtra())
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

	data, err = packUnlocked([]byte{}, obj.GetUnlocked())
	if err != nil {
		metricErrCnt.Inc(ctx, "insertreplace_packfield", 1)
		return err
	}

	tuple = append(tuple, data)

	data, err = packDescription([]byte{}, obj.GetDescription())
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
	logger.Trace(ctx, "Reward", obj.PrimaryString(), fmt.Sprintf("Insert packed tuple: '%X'", w))

	connection, err := octopus.Box(ctx, 0, activerecord.MasterInstanceType, "arcfg", nil)
	if err != nil {
		metricErrCnt.Inc(ctx, "insertreplace_preparebox", 1)
		logger.Error(ctx, "Reward", obj.PrimaryString(), fmt.Sprintf("Error get box '%s'", err))

		return err
	}

	respBytes, errCall := connection.Call(ctx, octopus.RequestTypeInsert, w)
	if errCall != nil {
		metricErrCnt.Inc(ctx, "insertreplace_box", 1)
		logger.Error(ctx, "Reward", obj.PrimaryString(), "Error insert into box", errCall, connection.Info())

		return errCall
	}

	metricTimer.Timing(ctx, "insertreplace_box")

	logger.Trace(ctx, "Reward", obj.PrimaryString(), fmt.Sprintf("Response from box '%X'", respBytes))

	tuplesData, err := octopus.ProcessResp(respBytes, octopus.NeedRespFlag|octopus.UniqRespFlag)
	if err != nil {
		metricErrCnt.Inc(ctx, "insertreplace_prespreparebox", 1)
		logger.Error(ctx, "Reward", obj.PrimaryString(), "Error parse response: ", err)

		return err
	}

	_, err = NewFromBox(ctx, tuplesData)
	if err != nil {
		metricErrCnt.Inc(ctx, "insertreplace_obj", 1)
		logger.Error(ctx, "Reward", obj.PrimaryString(), "Error in response: ", err)

		return err
	}

	obj.BaseField.Exists = true
	obj.BaseField.UpdateOps = []octopus.Ops{}
	obj.BaseField.Repaired = false

	logger.Debug(ctx, "Reward", obj.PrimaryString(), "Success insert")

	metricTimer.Finish(ctx, "insertreplace")

	return nil
}
