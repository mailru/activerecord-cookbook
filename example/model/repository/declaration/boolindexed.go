package repository

//ar:serverConf:arcfg
//ar:namespace:25
//ar:backend:octopus
type FieldsBoolindexed struct {
	Code      string `ar:"primary_key;size:"`
	Invisible bool   `ar:""`
}

type IndexesBoolindexed struct {
	Invisible bool `ar:"fields:Invisible"`
}
