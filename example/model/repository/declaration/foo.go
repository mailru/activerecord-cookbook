package repository

//ar:serverHost:box1
//ar:namespace:bar.foo
//ar:backend:octopus
type ProcFieldsFoo struct {
	SearchQuery string `ar:"input"`
	TraceID     string `ar:"input;output"`
	Status      string `ar:"output"`
	JsonRawData string `ar:"output"`
}
