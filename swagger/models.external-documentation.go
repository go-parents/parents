package swagger

type ExternalDocumentation struct {
	Description MarkdownString `json:",omitempty"`
	Url         UrlString      `validate:"required"`
}
