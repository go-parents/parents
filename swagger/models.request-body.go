package swagger

type RequestBody struct {
	Description MarkdownString           `json:",omitempty"`
	Content     map[MimeEnumOption]Media `validate:"required"`
	Required    bool
}
