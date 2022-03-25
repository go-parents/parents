package swagger

type Response struct {
	Description string                          `validate:"required"`
	Headers     map[HttpHeaderEnumOption]Header `json:",omitempty"`
	Content     map[MimeEnumOption]Media        `json:",omitempty"`
	Links       map[LinkDefinitionKey]Link      `json:",omitempty"`
}
