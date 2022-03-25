package swagger

type Info struct {
	Title          string           `validate:"required"`
	Description    MarkdownString   `json:",omitempty"`
	TermsOfService string           `json:",omitempty"`
	Contact        *Contact         `json:",omitempty"`
	License        *License         `json:",omitempty"`
	Version        ApiVersionString `validate:"required"`
}
