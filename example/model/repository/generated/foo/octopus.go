// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.5.3-4-gd9702e9 (Commit: d9702e9f)
package foo

import (
	"bytes"
	"context"
	"fmt"

	"github.com/mailru/activerecord-cookbook/example/model/s2s"
	serializerServiceResponse "github.com/mailru/activerecord-cookbook/example/model/serializer"
	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/iproto/iproto"
	"github.com/mailru/activerecord/pkg/octopus"
)

// proc struct
type Foo struct {
	params           FooParams
	fieldTraceID     string
	fieldJsonRawData *s2s.ServiceResponse
}

type FooList []*Foo

const (
	procName string = "foo"
)

func (obj *Foo) GetTraceID() string {
	return obj.fieldTraceID
}

func (obj *Foo) GetJsonRawData() *s2s.ServiceResponse {
	return obj.fieldJsonRawData
}

type FooParams struct {
	SearchQuery string
	TraceID     string
}

func (obj *Foo) GetParams() FooParams {
	return obj.params
}

func (obj *FooParams) arrayValues() []string {
	return []string{
		obj.SearchQuery,
		obj.TraceID,
	}
}

func Call(ctx context.Context, params FooParams) (*Foo, error) {
	logger := activerecord.Logger()
	ctx = logger.SetLoggerValueToContext(ctx, map[string]interface{}{"LuaProc": procName})
	metricTimer := activerecord.Metric().Timer("octopus", "Foo")
	metricErrCnt := activerecord.Metric().ErrorCount("octopus", "Foo")

	metricTimer.Timing(ctx, "call_proc")

	connection, err := box(ctx, 0, activerecord.ReplicaInstanceType)
	if err != nil {
		metricErrCnt.Inc(ctx, "call_proc_preparebox", 1)
		logger.Error(ctx, fmt.Sprintf("Error get box '%s'", err))

		return nil, err
	}

	td, err := octopus.CallLua(ctx, connection, procName, params.arrayValues()...)
	if err != nil {
		metricErrCnt.Inc(ctx, "call_proc", 1)
		return nil, fmt.Errorf("call lua procedure %s: %w", procName, err)
	}

	if len(td) != 1 {
		return nil, fmt.Errorf("invalid response len from lua call: %d", len(td))
	}

	ret, err := TupleToStruct(ctx, td[0])
	if err != nil {
		metricErrCnt.Inc(ctx, "call_proc_preparebox", 1)
		logger.Error(ctx, "Error in response: ", err)

		return nil, err
	}

	metricTimer.Finish(ctx, "call_proc")

	return ret, nil
}

func TupleToStruct(ctx context.Context, tuple octopus.TupleData) (*Foo, error) {
	np := Foo{}

	valTraceID, err := UnpackTraceID(bytes.NewReader(tuple.Data[1-1]))
	if err != nil {
		return nil, err
	}

	np.fieldTraceID = valTraceID

	valJsonRawData, err := UnpackJsonRawData(bytes.NewReader(tuple.Data[2-1]))
	if err != nil {
		return nil, err
	}

	np.fieldJsonRawData = valJsonRawData

	return &np, nil
}

func UnpackSearchQuery(r *bytes.Reader) (ret string, errRet error) {
	var SearchQuery string

	err := octopus.UnpackString(r, &SearchQuery, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field SearchQuery in tuple: '%w'", err)
		return
	}

	bvar := SearchQuery

	svar := bvar

	return svar, nil
}

func packSearchQuery(w []byte, SearchQuery string) ([]byte, error) {
	pvar := SearchQuery

	return octopus.PackString(w, pvar, iproto.ModeDefault), nil
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
func (obj *Foo) SetJsonRawData(JsonRawData *s2s.ServiceResponse) error {
	obj.fieldJsonRawData = JsonRawData

	return nil
}

func UnpackJsonRawData(r *bytes.Reader) (ret *s2s.ServiceResponse, errRet error) {
	var JsonRawData string

	err := octopus.UnpackString(r, &JsonRawData, iproto.ModeDefault)
	if err != nil {
		errRet = fmt.Errorf("error unpack field JsonRawData in tuple: '%w'", err)
		return
	}

	bvar := JsonRawData

	var svar s2s.ServiceResponse

	err = serializerServiceResponse.ServiceResponseUnmarshal(bvar, &svar)
	if err != nil {
		errRet = fmt.Errorf("error unmarshal field JsonRawData: %w", err)
		return
	}

	return &svar, nil
}

func packJsonRawData(w []byte, JsonRawData *s2s.ServiceResponse) ([]byte, error) {
	pvar, err := serializerServiceResponse.ServiceResponseMarshal(JsonRawData)
	if err != nil {
		return nil, fmt.Errorf("error marshal field JsonRawData: %w", err)
	}

	return octopus.PackString(w, pvar, iproto.ModeDefault), nil
}

// end proc struct

// box - возвращает коннектор для БД
// TODO
// - унести в пакет pkg/octopus тут общий код нет смысла его нагенеривать
// - сделать статистику по используемым инстансам
// - прикрутить локальный пингер и исключать недоступные инстансы
func box(ctx context.Context, shard int, instType activerecord.ShardInstanceType) (*octopus.Connection, error) {
	configPath := "arcfg"

	clusterInfo, err := activerecord.ConfigCacher().Get(
		ctx,
		configPath,
		activerecord.MapGlobParam{
			Timeout:  octopus.DefaultConnectionTimeout,
			PoolSize: octopus.DefaultPoolSize,
		},
		func(sic activerecord.ShardInstanceConfig) (activerecord.OptionInterface, error) {
			return octopus.NewOptions(
				sic.Addr,
				octopus.ServerModeType(sic.Mode),
				octopus.WithTimeout(sic.Timeout, sic.Timeout),
				octopus.WithPoolSize(sic.PoolSize),
			)
		},
	)
	if err != nil {
		return nil, fmt.Errorf("can't get cluster %s info: %w", configPath, err)
	}

	if len(clusterInfo) < int(shard) {
		return nil, fmt.Errorf("invalid shard num %d, max = %d", shard, len(clusterInfo))
	}

	var configBox activerecord.ShardInstance

	switch instType {
	case activerecord.ReplicaInstanceType:
		if len(clusterInfo[shard].Replicas) == 0 {
			return nil, fmt.Errorf("replicas not set")
		}

		configBox = clusterInfo[shard].NextReplica()
	case activerecord.ReplicaOrMasterInstanceType:
		if len(clusterInfo[shard].Replicas) != 0 {
			configBox = clusterInfo[shard].NextReplica()
			break
		}

		fallthrough
	case activerecord.MasterInstanceType:
		configBox = clusterInfo[shard].NextMaster()
	}

	conn, err := activerecord.ConnectionCacher().GetOrAdd(configBox, func(options interface{}) (activerecord.ConnectionInterface, error) {
		octopusOpt, ok := options.(*octopus.ConnectionOptions)
		if !ok {
			return nil, fmt.Errorf("invalit type of options %T, want Options", options)
		}

		return octopus.GetConnection(ctx, octopusOpt)
	})
	if err != nil {
		return nil, fmt.Errorf("error from connectionCacher: %w", err)
	}

	box, ok := conn.(*octopus.Connection)
	if !ok {
		return nil, fmt.Errorf("invalid connection type %T, want *octopus.Connection", conn)
	}

	return box, nil
}

func New(ctx context.Context) *Foo {
	newObj := Foo{}
	return &newObj
}

// end indexes