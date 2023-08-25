// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.8.7 (Commit: e17c811b)
package category

import (
	"context"
	"log"

	"gopkg.in/yaml.v3"

	"github.com/mailru/activerecord/pkg/activerecord"
)

type CategoryFTPK struct {
}

type CategoryFT struct {
	Params CategoryFTPK `yaml:"params"`
	All    int          `yaml:"all"`
}

func MarshalFixtures(objs []*Category) ([]byte, error) {
	fts := make([]CategoryFT, 0, len(objs))
	for _, obj := range objs {

		pk := CategoryFTPK{}
		fts = append(fts, CategoryFT{
			Params: pk,
			All:    obj.GetAll(),
		})
	}
	return yaml.Marshal(fts)
}

func UnmarshalFixtures(source []byte) []*Category {
	var fixtures []CategoryFT

	if err := yaml.Unmarshal(source, &fixtures); err != nil {
		log.Fatalf("unmarshal CategoryFT fixture: %v", err)
	}

	objs := make([]*Category, 0, len(fixtures))

	for _, ft := range fixtures {

		o := New(context.Background())
		o.setParams(CategoryParams{})
		if err := o.SetAll(ft.All); err != nil {
			log.Fatalf("can't set value %v to field All of Category fixture: %s", ft.All, err)
		}

		objs = append(objs, o)
	}

	return objs
}

func (objs CategoryList) String() string {
	o, err := MarshalFixtures(objs)
	if err != nil {
		activerecord.Logger().Fatal(context.Background(), err)
	}
	return string(o)
}

func UnmarshalDeleteFixtures(source []byte) []*Category {
	return UnmarshalFixtures(source)
}
