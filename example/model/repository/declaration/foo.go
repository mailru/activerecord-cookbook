package repository

import (
	"github.com/mailru/activerecord-cookbook/example/model/s2s"
)

//ar:serverConf:arcfg
//ar:namespace:foo
//ar:backend:octopus
type ProcFieldsFoo struct {
	SearchQuery string `ar:"input"`
	TraceID     string `ar:"input;output"`
	Status      int    `ar:"output"`
	JsonRawData string `ar:"output;serializer:ServiceResponse"`
}

type SerializersFoo struct {
	ServiceResponse s2s.ServiceResponse `ar:"unmarshaler:JSONUnmarshal;marshaler:JSONMarshal"`
}
