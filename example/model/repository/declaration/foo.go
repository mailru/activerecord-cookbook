package repository

//ar:serverConf:arcfg
//ar:namespace:foo
//ar:backend:octopus
type ProcFieldsFoo struct {
	SearchQuery string `ar:"input"`
	TraceID     string `ar:"input;output"`
	Status      string `ar:"output"`
	JsonRawData string `ar:"output"`
}
