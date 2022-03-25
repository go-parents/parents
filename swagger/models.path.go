package swagger

type Path struct {
	Tags []string `json:",omitempty"`

	Summary     string         `json:",omitempty"`
	Description MarkdownString `json:",omitempty"`

	ExternalDocs *ExternalDocumentation `json:",omitempty"`

	OperationId OperationIdString `validate:"required" json:",omitempty"`

	Consumes []MimeEnumOption `json:",omitempty"`
	Produces []MimeEnumOption `json:",omitempty"`

	Parameters []Parameter `json:",omitempty"`

	RequestBody *RequestBody `json:",omitempty"`

	Responses map[HttpStatusCodeEnumOption]Response `json:",omitempty"`

	// todo: build support for callbacks
	// Callbacks interface{}

	Deprecated bool

	Security map[SecurityDefinitionKey]SecurityScopeKey `json:",omitempty"`

	Servers []Server `json:",omitempty"`
}
