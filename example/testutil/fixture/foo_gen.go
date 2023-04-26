// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.5.3-9-geaa00ca (Commit: eaa00caf)
package fixture

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"sync"

	"github.com/mailru/activerecord-cookbook/example/model/repository/generated/foo"
	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
)

var fooOnce sync.Once
var fooStore map[string]*foo.Foo

//go:embed data/foo.yaml
var fooSource []byte

func initFoo() {
	fooOnce.Do(func() {
		fixtures := foo.UnmarshalFixtures(fooSource)

		fooStore = map[string]*foo.Foo{}
		for _, f := range fixtures {
			if _, ok := fooStore[f.GetParams().PK()]; ok {
				log.Fatalf("foo fixture with params %v are duplicated", f.GetParams())
			}
			fooStore[f.GetParams().PK()] = f
		}
	})
}

func GetFooByParams(params foo.FooParams) *foo.Foo {
	initFoo()

	res, ex := fooStore[params.PK()]
	if !ex {
		log.Fatalf("Foo fixture with params %v not found", params)
	}

	ctx := activerecord.Logger().SetLoggerValueToContext(context.Background(), map[string]interface{}{"GetFooByParams": params, "FixtureStore": "fooStore"})

	activerecord.Logger().Debug(ctx, foo.FooList([]*foo.Foo{res}))

	return res
}

type FooProcedureMocker struct{}

func GetFooProcedureMocker() FooProcedureMocker {
	return FooProcedureMocker{}
}

func (m FooProcedureMocker) ByFixtureParams(ctx context.Context, params foo.FooParams) octopus.FixtureType {

	return m.ByParamsMocks(ctx, params,
		[]octopus.MockEntities{
			GetFooByParams(params),
		})

}

func (m FooProcedureMocker) ByParamsMocks(ctx context.Context, params foo.FooParams, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateCallFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return foo.MockCallRequest(ctx, params)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by params: %s", err))
	}

	return oft
}
