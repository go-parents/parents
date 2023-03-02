package swagger

type Schema struct {
    Swagger VersionEnumOption `validate:"required"`
    Info    Info              `validate:"required"`

    Host     UrlHostString      `json:",omitempty"`
    BasePath UrlPathString      `json:",omitempty"`
    Schemes  []SchemeEnumOption `json:",omitempty"`

    Consumes []MimeEnumOption `json:",omitempty"`
    Produces []MimeEnumOption `json:",omitempty"`

    Paths map[UrlPathString]PathMethodMap `json:",omitempty"`

    Definitions map[DefinitionKey]Definition `json:",omitempty"`

    SecurityDefinitions map[SecurityDefinitionKey]SecurityDefinitionKey `json:",omitempty"`
    Security            map[SecurityDefinitionKey]SecurityScopeKey      `json:",omitempty"`
}
