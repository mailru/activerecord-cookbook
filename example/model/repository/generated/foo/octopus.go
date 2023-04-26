// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.5.3-9-geaa00ca (Commit: eaa00caf)
package foo

import (
	"bytes"
	"context"
	"fmt"

	"github.com/mailru/activerecord-cookbook/example/model/s2s"
	serializerSearchQuery "github.com/mailru/activerecord-cookbook/example/model/serializer"
	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/iproto/iproto"
	"github.com/mailru/activerecord/pkg/octopus"
	serializerServiceResponse "github.com/mailru/activerecord/pkg/serializer"
)

// proc struct
type Foo struct {
	params           FooParams
	fieldTraceID     string
	fieldStatus      int
	fieldJsonRawData s2s.ServiceResponse
}

type FooList []*Foo

const (
	procName     string = "foo"
	cntOutFields uint32 = 3
)

func (obj *Foo) GetTraceID() string {
	return obj.fieldTraceID
}

func (obj *Foo) GetStatus() int {
	return obj.fieldStatus
}

func (obj *Foo) GetJsonRawData() s2s.ServiceResponse {
	return obj.fieldJsonRawData
}

type FooParams struct {
	SearchQuery map[string]string
	TraceID     string
}

func (obj *Foo) GetParams() FooParams {
	return obj.params
}

func (obj *Foo) setParams(params FooParams) error {
	obj.params = params

	return nil
}

func (obj *FooParams) arrayValues() ([]string, error) {
	ret := []string{}
	pvar, err := serializerSearchQuery.SearchQueryMarshal(obj.SearchQuery)
	if err != nil {
		return nil, fmt.Errorf("error marshal param field SearchQuery: %w", err)
	}

	ret = append(ret, pvar...)
	ret = append(ret, obj.TraceID)

	return ret, nil
}

func (obj FooParams) PK() string {

	return fmt.Sprint(
		obj.SearchQuery,
		obj.TraceID,
	)

}

func Call(ctx context.Context, params FooParams) (*Foo, error) {
	logger := activerecord.Logger()
	ctx = logger.SetLoggerValueToContext(ctx, map[string]interface{}{"LuaProc": procName})
	metricTimer := activerecord.Metric().Timer("octopus", "Foo")
	metricErrCnt := activerecord.Metric().ErrorCount("octopus", "Foo")

	metricTimer.Timing(ctx, "call_proc")

	connection, err := octopus.Box(ctx, 0, activerecord.ReplicaInstanceType)
	if err != nil {
		metricErrCnt.Inc(ctx, "call_proc_preparebox", 1)
		logger.Error(ctx, fmt.Sprintf("Error get box '%s'", err))

		return nil, err
	}

	var args []string

	args, err = params.arrayValues()
	if err != nil {
		metricErrCnt.Inc(ctx, "call_proc_preparebox", 1)
		return nil, fmt.Errorf("Error parse args of procedure %s: %w", procName, err)
	}

	td, err := octopus.CallLua(ctx, connection, procName, args...)
	if err != nil {
		metricErrCnt.Inc(ctx, "call_proc", 1)
		return nil, fmt.Errorf("call lua procedure %s: %w", procName, err)
	}

	if len(td) != 1 {
		return nil, fmt.Errorf("invalid response len from lua call: %d. Only one tuple supported", len(td))
	}

	ret, err := TupleToStruct(ctx, td[0])
	if err != nil {
		metricErrCnt.Inc(ctx, "call_proc_preparebox", 1)
		logger.Error(ctx, "Error in response: ", err)

		return nil, err
	}

	metricTimer.Finish(ctx, "call_proc")

	activerecord.Logger().CollectQueries(ctx, CallMockerLogger(params, FooList([]*Foo{ret})))

	return ret, nil
}

func TupleToStruct(ctx context.Context, tuple octopus.TupleData) (*Foo, error) {
	if tuple.Cnt < cntOutFields {
		return nil, fmt.Errorf("not enought selected fields %d in response tuple: %d but expected %d fields", tuple.Cnt, tuple.Cnt, cntOutFields)
	}

	np := Foo{}

	valTraceID, err := UnpackTraceID(bytes.NewReader(tuple.Data[0]))
	if err != nil {
		return nil, err
	}

	np.fieldTraceID = valTraceID
	valStatus, err := UnpackStatus(bytes.NewReader(tuple.Data[1]))
	if err != nil {
		return nil, err
	}

	np.fieldStatus = valStatus
	valJsonRawData, err := UnpackJsonRawData(bytes.NewReader(tuple.Data[2]))
	if err != nil {
		return nil, err
	}

	np.fieldJsonRawData = valJsonRawData

	return &np, nil
}

func (obj *Foo) SetTraceID(TraceID string) error {
	obj.fieldTraceID = TraceID

	return nil
}

func UnpackTraceID(r *bytes.Reader) (ret string, errRet error) {
	var TraceID string

	err := octopus.UnpackString(r, &TraceID, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field TraceID in tuple: '%w'", err)
		return
	}

	bvar := TraceID

	svar := bvar

	return svar, nil
}

func packTraceID(w []byte, TraceID string) ([]byte, error) {
	pvar := TraceID

	return octopus.PackString(w, pvar, iproto.ModeDefault), nil
}
func (obj *Foo) SetStatus(Status int) error {
	obj.fieldStatus = Status

	return nil
}

func UnpackStatus(r *bytes.Reader) (ret int, errRet error) {
	var Status uint32

	err := iproto.UnpackUint32(r, &Status, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field Status in tuple: '%w'", err)
		return
	}

	bvar := int(Status)

	svar := bvar

	return svar, nil
}

func packStatus(w []byte, Status int) ([]byte, error) {
	pvar := uint32(Status)

	return iproto.PackUint32(w, pvar, iproto.ModeDefault), nil
}
func (obj *Foo) SetJsonRawData(JsonRawData s2s.ServiceResponse) error {
	obj.fieldJsonRawData = JsonRawData

	return nil
}

func UnpackJsonRawData(r *bytes.Reader) (ret s2s.ServiceResponse, errRet error) {
	var JsonRawData string

	err := octopus.UnpackString(r, &JsonRawData, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field JsonRawData in tuple: '%w'", err)
		return
	}

	bvar := JsonRawData

	var svar s2s.ServiceResponse

	err = serializerServiceResponse.JSONUnmarshal(bvar, &svar)
	if err != nil {
		errRet = fmt.Errorf("error unmarshal field JsonRawData: %w", err)
		return
	}

	return svar, nil
}

func packJsonRawData(w []byte, JsonRawData s2s.ServiceResponse) ([]byte, error) {
	pvar, err := serializerServiceResponse.JSONMarshal(JsonRawData)
	if err != nil {
		return nil, fmt.Errorf("error marshal field JsonRawData: %w", err)
	}

	return octopus.PackString(w, pvar, iproto.ModeDefault), nil
}

// end proc struct

func New(ctx context.Context) *Foo {
	newObj := Foo{}
	return &newObj
}

// end indexes
