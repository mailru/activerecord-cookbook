// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.5.3-7-g1454a87 (Commit: 1454a870)
package fixture

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"sync"

	"github.com/mailru/activerecord-cookbook/example/model/repository/generated/category"
	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
)

var categoryOnce sync.Once
var categoryStore map[category.CategoryParams]*category.Category

//go:embed data/category.yaml
var categorySource []byte

func initCategory() {
	categoryOnce.Do(func() {
		fixtures := category.UnmarshalFixtures(categorySource)

		categoryStore = map[category.CategoryParams]*category.Category{}
		for _, f := range fixtures {
			if _, ok := categoryStore[f.GetParams()]; ok {
				log.Fatalf("category fixture with params %v are duplicated", f.GetParams())
			}
			categoryStore[f.GetParams()] = f
		}
	})
}

func GetCategoryByParams(params category.CategoryParams) *category.Category {
	initCategory()

	res, ex := categoryStore[params]
	if !ex {
		log.Fatalf("Category fixture with params %v not found", params)
	}

	ctx := activerecord.Logger().SetLoggerValueToContext(context.Background(), map[string]interface{}{"GetCategoryByParams": params, "FixtureStore": "categoryStore"})

	activerecord.Logger().Debug(ctx, category.CategoryList([]*category.Category{res}))

	return res
}

type CategoryProcedureMocker struct{}

func GetCategoryProcedureMocker() CategoryProcedureMocker {
	return CategoryProcedureMocker{}
}

func (m CategoryProcedureMocker) ByFixture(ctx context.Context) octopus.FixtureType {

	return m.ByMocks(ctx,
		[]octopus.MockEntities{
			GetCategoryByParams(category.CategoryParams{}),
		})
}

func (m CategoryProcedureMocker) ByMocks(ctx context.Context, mocks []octopus.MockEntities) octopus.FixtureType {
	oft, err := octopus.CreateCallFixture(
		func(wsubME []octopus.MockEntities) []byte {
			return category.MockCallRequest(ctx)
		},
		mocks,
	)
	if err != nil {
		activerecord.Logger().Fatal(ctx, fmt.Sprintf("Error create mock by params: %s", err))
	}

	return oft
}
