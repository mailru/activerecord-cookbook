// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.8.5-1-gaa389f8 (Commit: aa389f82)
package arobj

import (
	"context"
	"log"

	"gopkg.in/yaml.v3"

	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
)

func (objs ArObjList) String() string {
	o, err := MarshalFixtures(objs)
	if err != nil {
		activerecord.Logger().Fatal(context.Background(), err)
	}
	return string(o)
}

type ArObjFT struct {
	ID        int32  `yaml:"id"`
	Name      string `yaml:"name"`
	AnotherID int32  `yaml:"another_id"`
	Type      string `yaml:"type"`
	Flags     uint32 `yaml:"flags"`
}

func MarshalFixtures(objs []*ArObj) ([]byte, error) {
	fts := make([]ArObjFT, 0, len(objs))
	for _, obj := range objs {
		fts = append(fts, ArObjFT{
			ID:        obj.GetID(),
			Name:      obj.GetName(),
			AnotherID: obj.GetAnotherID(),
			Type:      obj.GetType(),
			Flags:     obj.GetFlags(),
		})
	}
	return yaml.Marshal(fts)
}

func UnmarshalFixtures(source []byte) []*ArObj {
	var fixtures []ArObjFT

	if err := yaml.Unmarshal(source, &fixtures); err != nil {
		log.Fatalf("unmarshal ArObjFT fixture: %v", err)
	}

	objs := make([]*ArObj, 0, len(fixtures))

	for _, ft := range fixtures {

		o := New(context.Background())
		if err := o.SetID(ft.ID); err != nil {
			log.Fatalf("can't set value %v to field ID of ArObj fixture: %s", ft.ID, err)
		}
		if err := o.SetName(ft.Name); err != nil {
			log.Fatalf("can't set value %v to field Name of ArObj fixture: %s", ft.Name, err)
		}
		if err := o.SetAnotherID(ft.AnotherID); err != nil {
			log.Fatalf("can't set value %v to field AnotherID of ArObj fixture: %s", ft.AnotherID, err)
		}
		if err := o.SetType(ft.Type); err != nil {
			log.Fatalf("can't set value %v to field Type of ArObj fixture: %s", ft.Type, err)
		}
		if err := o.SetFlags(ft.Flags); err != nil {
			log.Fatalf("can't set value %v to field Flags of ArObj fixture: %s", ft.Flags, err)
		}

		objs = append(objs, o)
	}

	return objs
}

type ArObjUpdateFT struct {
	ID int32 `yaml:"id"`

	UpdateOptions []ArObjUpdateFixtureOptions `yaml:"update_options"`
}

type ArObjUpdateFixtureOptions struct {
	Name *ArObjNameUpdateFixtureOption `yaml:"name"`

	AnotherID *ArObjAnotherIDUpdateFixtureOption `yaml:"another_id"`

	Type *ArObjTypeUpdateFixtureOption `yaml:"type"`

	Flags *ArObjFlagsUpdateFixtureOption `yaml:"flags"`
}

type ArObjNameUpdateFixtureOption struct {
	Value string `yaml:"set_value"`
}

type ArObjAnotherIDUpdateFixtureOption struct {
	Value int32 `yaml:"set_value"`
}

type ArObjTypeUpdateFixtureOption struct {
	Value string `yaml:"set_value"`
}

type ArObjFlagsUpdateFixtureOption struct {
	Value uint32 `yaml:"set_value"`
}

func UnmarshalUpdateFixtures(source []byte) []*ArObj {
	var fixtures []ArObjUpdateFT

	if err := yaml.Unmarshal(source, &fixtures); err != nil {
		log.Fatalf("unmarshal ArObjUpdateFT fixture: %v", err)
	}

	objs := make([]*ArObj, 0, len(fixtures))

	for _, ft := range fixtures {
		obj := New(context.Background())

		if err := obj.SetID(ft.ID); err != nil {
			log.Fatalf("error SetID: %v", err)
		}
		obj.BaseField.Exists = true
		obj.BaseField.UpdateOps = []octopus.Ops{}

		SetFixtureUpdateOptions(obj, ft.UpdateOptions)

		objs = append(objs, obj)
	}

	return objs
}

func UnmarshalInsertReplaceFixtures(source []byte) []*ArObj {
	return UnmarshalFixtures(source)
}

func SetFixtureUpdateOptions(obj *ArObj, updateOptions []ArObjUpdateFixtureOptions) {
	for priority, updateOption := range updateOptions {

		// Если опции нет, то ее нет в списке на установку
		if updateOption.Name != nil {
			if err := obj.SetName(updateOption.Name.Value); err != nil {
				log.Fatalf("SetName[priority: %d] error: %v", priority, err)
			}
		}

		// Если опции нет, то ее нет в списке на установку
		if updateOption.AnotherID != nil {
			if err := obj.SetAnotherID(updateOption.AnotherID.Value); err != nil {
				log.Fatalf("SetAnotherID[priority: %d] error: %v", priority, err)
			}
		}

		// Если опции нет, то ее нет в списке на установку
		if updateOption.Type != nil {
			if err := obj.SetType(updateOption.Type.Value); err != nil {
				log.Fatalf("SetType[priority: %d] error: %v", priority, err)
			}
		}

		// Если опции нет, то ее нет в списке на установку
		if updateOption.Flags != nil {
			if err := obj.SetFlags(updateOption.Flags.Value); err != nil {
				log.Fatalf("SetFlags[priority: %d] error: %v", priority, err)
			}
		}

	}
}

func UnmarshalDeleteFixtures(source []byte) []*ArObj {
	return UnmarshalFixtures(source)
}
