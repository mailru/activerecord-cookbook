package repository

import "github.com/mailru/activerecord-cookbook/example/model/dictionary"

//ar:serverHost:127.0.0.1;serverPort:11011;serverTimeout:500
//ar:namespace:6
//ar:backend:octopus
type FieldsPromoperiods struct {
	ID         string `ar:"primary_key;unique;size:36"`
	Code       string `ar:"size:128;selector:SelectByCode"`
	Email      string `ar:"size:256;selector:SelectByEmail"`
	Start      int32  `ar:""`
	Finish     uint64 `ar:"serializer:Product"`
	Action     string `ar:""`
	Platform   string `ar:"serializer:JSON;size:64"`
	Promobunch uint32 `ar:"mutators:set_bit,clear_bit"`
	Platforms  uint32 `ar:""`
	PlanID     int32  `ar:""`
	PlanType   string `ar:""`
	Price      string `ar:"serializer:Printf,%.2f"`
}

type FieldsObjectPromoperiods struct {
	Plan  bool   `ar:"key:ID;object:arobj;field:PlanID"`
	Plans []bool `ar:"key:Type;object:arobj;field:PlanType"`
}

type (
	IndexesPromoperiods struct {
		PlanTypePrice bool `ar:"fields:PlanType,Price;unique"`
		EmailCode     bool `ar:"fields:Email,Code;unique"`
		EmailAction   bool `ar:"fields:Email,Action;unique"`
	}
	IndexPartsPromoperiods struct {
		EmailPart bool `ar:"index:EmailCode;fieldnum:1"`
	}
)

type SerializersPromoperiods struct {
	JSON    map[string]interface{} `ar:""`
	Product *dictionary.Product    `ar:"pkg:github.com/mailru/activerecord-cookbook/example/model/serializer"`
	Printf  float64                `ar:""`
}

type TriggersPromoperiods struct {
	RepairTuple bool `ar:"pkg:github.com/mailru/activerecord-cookbook/example/model/repository/repair;func:Promoperiod;param:Defaults"`
}

type FlagsPromoperiods struct {
	Platforms bool `ar:"flags:web,ios,android,huawei"`
}
