// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.8.5-1-gaa389f8 (Commit: aa389f82)
package foo

import (
	"context"
	"log"

	"gopkg.in/yaml.v3"

	"github.com/mailru/activerecord-cookbook/example/model/s2s"
	"github.com/mailru/activerecord/pkg/activerecord"
)

type FooFTPK struct {
	SearchQuery map[string]string `yaml:"search_query"`

	TraceID string `yaml:"trace_id"`
}

type FooFT struct {
	Params      FooFTPK             `yaml:"params"`
	Status      int                 `yaml:"status"`
	JsonRawData s2s.ServiceResponse `yaml:"json_raw_data"`
	TraceID     string              `yaml:"trace_id"`
}

func MarshalFixtures(objs []*Foo) ([]byte, error) {
	fts := make([]FooFT, 0, len(objs))
	for _, obj := range objs {
		params := obj.GetParams()

		pk := FooFTPK{
			SearchQuery: params.SearchQuery,
			TraceID:     params.TraceID,
		}
		fts = append(fts, FooFT{
			Params:      pk,
			Status:      obj.GetStatus(),
			JsonRawData: obj.GetJsonRawData(),
			TraceID:     obj.GetTraceID(),
		})
	}
	return yaml.Marshal(fts)
}

func UnmarshalFixtures(source []byte) []*Foo {
	var fixtures []FooFT

	if err := yaml.Unmarshal(source, &fixtures); err != nil {
		log.Fatalf("unmarshal FooFT fixture: %v", err)
	}

	objs := make([]*Foo, 0, len(fixtures))

	for _, ft := range fixtures {

		o := New(context.Background())
		o.setParams(FooParams{
			SearchQuery: ft.Params.SearchQuery,
			TraceID:     ft.Params.TraceID,
		})
		if err := o.SetStatus(ft.Status); err != nil {
			log.Fatalf("can't set value %v to field Status of Foo fixture: %s", ft.Status, err)
		}
		if err := o.SetJsonRawData(ft.JsonRawData); err != nil {
			log.Fatalf("can't set value %v to field JsonRawData of Foo fixture: %s", ft.JsonRawData, err)
		}
		if err := o.SetTraceID(ft.TraceID); err != nil {
			log.Fatalf("can't set value %v to field TraceID of Foo fixture: %s", ft.TraceID, err)
		}

		objs = append(objs, o)
	}

	return objs
}

func (objs FooList) String() string {
	o, err := MarshalFixtures(objs)
	if err != nil {
		activerecord.Logger().Fatal(context.Background(), err)
	}
	return string(o)
}

func UnmarshalDeleteFixtures(source []byte) []*Foo {
	return UnmarshalFixtures(source)
}
