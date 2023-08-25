package repository

import (
	"github.com/mailru/activerecord-cookbook/example/model/ds"
)

//ar:serverConf:arcfg
//ar:namespace:24
//ar:backend:octopus
type FieldsReward struct {
	Code        string `ar:"primary_key;size:"`
	Services    string `ar:"serializer:Services;size:"`
	Partner     string `ar:"selector:SelectByPartner;mutators:PartnerUpdate;size:"`
	Extra       string `ar:"serializer:Extra;mutators:ExtraPartUpdate;size:"`
	Flags       string `ar:"serializer:Flags;size:"`
	Unlocked    string `ar:"serializer:Unlocked;mutators:UnlockedPartUpdate;size:"`
	Description string `ar:"serializer:Description;size:"`
}

type SerializersReward struct {
	Extra       *ds.Extra              `ar:"unmarshaler:MapstructureUnmarshal;marshaler:MapstructureMarshal"`
	Services    *ds.Services           `ar:"unmarshaler:MapstructureUnmarshal;marshaler:MapstructureMarshal"`
	Flags       map[string]interface{} `ar:"unmarshaler:JSONUnmarshal;marshaler:JSONMarshal"`
	Unlocked    ds.ServiceUnlocked     `ar:"unmarshaler:JSONUnmarshal;marshaler:JSONMarshal"`
	Description *string                `ar:"unmarshaler:JSONUnmarshal;marshaler:JSONMarshal"`
}

type MutatorsReward struct {
	ExtraPartUpdate    *ds.Extra          `ar:"update:lua.updateExtra;pkg:github.com/mailru/activerecord-cookbook/example/ds;"`
	UnlockedPartUpdate ds.ServiceUnlocked `ar:"update:lua.updateUnlocked;pkg:github.com/mailru/activerecord-cookbook/example/ds;"`
	PartnerUpdate      string             `ar:"update:lua.updateRewardPartner,param1,param2;replace:lua.replaceRewardPartner;"`
}
