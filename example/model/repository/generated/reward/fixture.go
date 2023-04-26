// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.5.3-7-g1454a87 (Commit: 1454a870)
package reward

import (
	"context"
	"log"

	"gopkg.in/yaml.v3"

	"github.com/mailru/activerecord-cookbook/example/model/ds"
	"github.com/mailru/activerecord/pkg/activerecord"
	"github.com/mailru/activerecord/pkg/octopus"
)

func (objs RewardList) String() string {
	o, err := MarshalFixtures(objs)
	if err != nil {
		activerecord.Logger().Fatal(context.Background(), err)
	}
	return string(o)
}

type RewardFT struct {
	Code        string                 `yaml:"code"`
	Services    *ds.Services           `yaml:"services"`
	Partner     string                 `yaml:"partner"`
	Extra       *ds.Extra              `yaml:"extra"`
	Flags       map[string]interface{} `yaml:"flags"`
	Unlocked    ds.ServiceUnlocked     `yaml:"unlocked"`
	Description *string                `yaml:"description"`
}

func MarshalFixtures(objs []*Reward) ([]byte, error) {
	fts := make([]RewardFT, 0, len(objs))
	for _, obj := range objs {
		fts = append(fts, RewardFT{
			Code:        obj.GetCode(),
			Services:    obj.GetServices(),
			Partner:     obj.GetPartner(),
			Extra:       obj.GetExtra(),
			Flags:       obj.GetFlags(),
			Unlocked:    obj.GetUnlocked(),
			Description: obj.GetDescription(),
		})
	}
	return yaml.Marshal(fts)
}

func UnmarshalFixtures(source []byte) []*Reward {
	var fixtures []RewardFT

	if err := yaml.Unmarshal(source, &fixtures); err != nil {
		log.Fatalf("unmarshal RewardFT fixture: %v", err)
	}

	objs := make([]*Reward, 0, len(fixtures))

	for _, ft := range fixtures {

		o := New(context.Background())
		if err := o.SetCode(ft.Code); err != nil {
			log.Fatalf("can't set value %v to field Code of Reward fixture: %s", ft.Code, err)
		}
		if err := o.SetServices(ft.Services); err != nil {
			log.Fatalf("can't set value %v to field Services of Reward fixture: %s", ft.Services, err)
		}
		if err := o.SetPartner(ft.Partner); err != nil {
			log.Fatalf("can't set value %v to field Partner of Reward fixture: %s", ft.Partner, err)
		}
		if err := o.SetExtra(ft.Extra); err != nil {
			log.Fatalf("can't set value %v to field Extra of Reward fixture: %s", ft.Extra, err)
		}
		if err := o.SetFlags(ft.Flags); err != nil {
			log.Fatalf("can't set value %v to field Flags of Reward fixture: %s", ft.Flags, err)
		}
		if err := o.SetUnlocked(ft.Unlocked); err != nil {
			log.Fatalf("can't set value %v to field Unlocked of Reward fixture: %s", ft.Unlocked, err)
		}
		if err := o.SetDescription(ft.Description); err != nil {
			log.Fatalf("can't set value %v to field Description of Reward fixture: %s", ft.Description, err)
		}

		objs = append(objs, o)
	}

	return objs
}

type RewardUpdateFT struct {
	Code string `yaml:"code"`

	UpdateOptions []RewardUpdateFixtureOptions `yaml:"update_options"`
}

type RewardUpdateFixtureOptions struct {
	Services *RewardServicesUpdateFixtureOption `yaml:"services"`

	Partner *RewardPartnerUpdateFixtureOption `yaml:"partner"`

	Extra *RewardExtraUpdateFixtureOption `yaml:"extra"`

	Flags *RewardFlagsUpdateFixtureOption `yaml:"flags"`

	Unlocked *RewardUnlockedUpdateFixtureOption `yaml:"unlocked"`

	Description *RewardDescriptionUpdateFixtureOption `yaml:"description"`
}

type RewardServicesUpdateFixtureOption struct {
	Value *ds.Services `yaml:"set_value"`
}

type RewardPartnerUpdateFixtureOption struct {
	Value string `yaml:"set_value"`
}

type RewardExtraUpdateFixtureOption struct {
	Value *ds.Extra `yaml:"set_value"`
}

type RewardFlagsUpdateFixtureOption struct {
	Value map[string]interface{} `yaml:"set_value"`
}

type RewardUnlockedUpdateFixtureOption struct {
	Value ds.ServiceUnlocked `yaml:"set_value"`
}

type RewardDescriptionUpdateFixtureOption struct {
	Value *string `yaml:"set_value"`
}

func UnmarshalUpdateFixtures(source []byte) []*Reward {
	var fixtures []RewardUpdateFT

	if err := yaml.Unmarshal(source, &fixtures); err != nil {
		log.Fatalf("unmarshal RewardUpdateFT fixture: %v", err)
	}

	objs := make([]*Reward, 0, len(fixtures))

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

func UnmarshalInsertReplaceFixtures(source []byte) []*Reward {
	return UnmarshalFixtures(source)
}

func SetFixtureUpdateOptions(obj *Reward, updateOptions []RewardUpdateFixtureOptions) {
	for priority, updateOption := range updateOptions {

		// Если опции нет, то ее нет в списке на установку
		if updateOption.Services != nil {
			if err := obj.SetServices(updateOption.Services.Value); err != nil {
				log.Fatalf("SetServices[priority: %d] error: %v", priority, err)
			}
		}

		// Если опции нет, то ее нет в списке на установку
		if updateOption.Partner != nil {
			if err := obj.SetPartner(updateOption.Partner.Value); err != nil {
				log.Fatalf("SetPartner[priority: %d] error: %v", priority, err)
			}
		}

		// Если опции нет, то ее нет в списке на установку
		if updateOption.Extra != nil {
			if err := obj.SetExtra(updateOption.Extra.Value); err != nil {
				log.Fatalf("SetExtra[priority: %d] error: %v", priority, err)
			}
		}

		// Если опции нет, то ее нет в списке на установку
		if updateOption.Flags != nil {
			if err := obj.SetFlags(updateOption.Flags.Value); err != nil {
				log.Fatalf("SetFlags[priority: %d] error: %v", priority, err)
			}
		}

		// Если опции нет, то ее нет в списке на установку
		if updateOption.Unlocked != nil {
			if err := obj.SetUnlocked(updateOption.Unlocked.Value); err != nil {
				log.Fatalf("SetUnlocked[priority: %d] error: %v", priority, err)
			}
		}

		// Если опции нет, то ее нет в списке на установку
		if updateOption.Description != nil {
			if err := obj.SetDescription(updateOption.Description.Value); err != nil {
				log.Fatalf("SetDescription[priority: %d] error: %v", priority, err)
			}
		}

	}
}
