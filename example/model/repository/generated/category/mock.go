// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.11.0-b (Commit: 6934fae2)
package category

import (
	"context"
	"fmt"

	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
)

func (obj *Category) MockSelectResponse() ([][]byte, error) {
	tuple := [][]byte{}

	var data []byte

	var err error

	data, err = packAll([]byte{}, obj.GetAll())
	if err != nil {
		return nil, err
	}

	tuple = append(tuple, data)

	return tuple, nil
}

func MockCallRequest(ctx context.Context) []byte {
	log := activerecord.Logger()
	ctx = log.SetLoggerValueToContext(ctx, map[string]interface{}{"Proc": "Category"})

	return octopus.PackLua(procName)
}

func (obj *Category) RepoSelector(ctx context.Context) (any, error) {
	data, err := Call(ctx)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, fmt.Errorf("call Category with params %v: %w", obj.params, activerecord.ErrNoData)
	}

	return data, err
}

func CallMockerLogger(res CategoryList) func() (activerecord.MockerLogger, error) {
	return func() (activerecord.MockerLogger, error) {

		mockerName := "mockerCategoryByParams"
		mocker := "fixture.GetCategoryProcedureMocker()"

		fixture := mocker
		fixture += ".ByFixture(ctx)"

		return activerecord.MockerLogger{MockerName: mockerName, Mockers: mocker, FixturesSelector: fixture, ResultName: "category", Results: res}, nil
	}
}
