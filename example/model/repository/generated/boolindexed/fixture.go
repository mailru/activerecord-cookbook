// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.4.2-4-g78facfc (Commit: 78facfc7)
package boolindexed

import (
	"context"
	"log"

	"gopkg.in/yaml.v3"

	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
)

type BoolindexedFT struct {
	Code      string `yaml:"code"`
	Invisible bool   `yaml:"invisible"`
}

func MarshalFixtures(objs []*Boolindexed) ([]byte, error) {
	fts := make([]BoolindexedFT, 0, len(objs))
	for _, obj := range objs {
		fts = append(fts, BoolindexedFT{
			Code:      obj.GetCode(),
			Invisible: obj.GetInvisible(),
		})
	}
	return yaml.Marshal(fts)
}

func UnmarshalFixtures(source []byte) []*Boolindexed {
	var fixtures []BoolindexedFT

	if err := yaml.Unmarshal(source, &fixtures); err != nil {
		log.Fatalf("unmarshal BoolindexedFT fixture: %v", err)
	}

	objs := make([]*Boolindexed, 0, len(fixtures))

	for _, ft := range fixtures {

		o := New(context.Background())
		if err := o.SetCode(ft.Code); err != nil {
			log.Fatalf("can't set value %v to field Code of Boolindexed fixture: %s", ft.Code, err)
		}
		if err := o.SetInvisible(ft.Invisible); err != nil {
			log.Fatalf("can't set value %v to field Invisible of Boolindexed fixture: %s", ft.Invisible, err)
		}

		objs = append(objs, o)
	}

	return objs
}

func (objs BoolindexedList) String() string {
	o, err := MarshalFixtures(objs)
	if err != nil {
		activerecord.Logger().Fatal(context.Background(), err)
	}
	return string(o)
}

type BoolindexedUpdateFT struct {
	Code string `yaml:"code"`

	UpdateOptions []BoolindexedUpdateFixtureOptions `yaml:"update_options"`
}

type BoolindexedUpdateFixtureOptions struct {
	Invisible *BoolindexedInvisibleUpdateFixtureOption `yaml:"invisible"`
}

type BoolindexedInvisibleUpdateFixtureOption struct {
	Value bool `yaml:"set_value"`
}

func UnmarshalUpdateFixtures(source []byte) []*Boolindexed {
	var fixtures []BoolindexedUpdateFT

	if err := yaml.Unmarshal(source, &fixtures); err != nil {
		log.Fatalf("unmarshal BoolindexedUpdateFT fixture: %v", err)
	}

	objs := make([]*Boolindexed, 0, len(fixtures))

	for _, ft := range fixtures {
		obj := New(context.Background())

		if err := obj.SetCode(ft.Code); err != nil {
			log.Fatalf("error SetCode: %v", err)
		}
		obj.BaseField.Exists = true
		obj.BaseField.UpdateOps = []octopus.Ops{}

		SetFixtureUpdateOptions(obj, ft.UpdateOptions)

		objs = append(objs, obj)
	}

	return objs
}

func UnmarshalInsertReplaceFixtures(source []byte) []*Boolindexed {
	return UnmarshalFixtures(source)
}

func SetFixtureUpdateOptions(obj *Boolindexed, updateOptions []BoolindexedUpdateFixtureOptions) {
	for priority, updateOption := range updateOptions {

		// Если опции нет, то ее нет в списке на установку
		if updateOption.Invisible != nil {
			if err := obj.SetInvisible(updateOption.Invisible.Value); err != nil {
				log.Fatalf("SetInvisible[priority: %d] error: %v", priority, err)
			}
		}

	}
}
