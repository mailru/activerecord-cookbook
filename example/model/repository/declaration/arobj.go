package repository

//ar:serverConf:box1
//ar:namespace:5
//ar:backend:octopus
type FieldsArObj struct {
	ID        int32  `ar:"primary_key;unique"`
	Name      string `ar:"size:256"`
	AnotherID int32  `ar:"mutators:inc,dec"`
	Type      string `ar:"selector:SelectByType;size:64"`
	Flags     uint32 `ar:"mutators:set_bit,clear_bit"`
}

type (
	IndexesArObj struct {
		TypeID bool `ar:"fields:Type,ID;unique"`
	}
	IndexPartsArObj struct {
		TypePart bool `ar:"index:TypeID;fieldnum:1;selector:SelectByTypePart"`
	}
)
