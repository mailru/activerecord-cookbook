// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.5.3-7-g1454a87 (Commit: 1454a870)
package foo

import (
	"context"
	"fmt"

	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
)

func (obj *Foo) MockSelectResponse() ([][]byte, error) {
	tuple := [][]byte{}

	var data []byte

	var err error

	data, err = packTraceID([]byte{}, obj.GetTraceID())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)
	data, err = packStatus([]byte{}, obj.GetStatus())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)
	data, err = packJsonRawData([]byte{}, obj.GetJsonRawData())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	return tuple, nil
}

func MockCallRequest(ctx context.Context, params FooParams) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"MockCallRequest": params, "Proc": "Foo"})

	return octopus.PackLua(procName, params.arrayValues()...)
}

func (obj *Foo) RepoSelector(ctx context.Context) (any, error) {
	data, err := Call(ctx, obj.params)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, fmt.Errorf("call Foo with params %v: %w", obj.params, activerecord.ErrNoData)
	}

	return data, err
}

func CallMockerLogger(params FooParams, res FooList) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {

		mockerName := "mockerFooByParams"
		mocker := "fixture.GetFooProcedureMocker()"

		fixture := "ps := FooParams{ \n"
		fixture += "SearchQuery: params.SearchQuery,\n"
		fixture += "TraceID: params.TraceID,\n"
		fixture += "}\n"
		fixture += mocker
		fixture += ".ByFixtureParams(ctx, ps)"

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "foo", Results: res}, nil
	}
}
