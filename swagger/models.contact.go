package swagger

type Contact struct {
	Name  string `validate:"required"`
	Url   string `json:",omitempty"`
	Email string `json:",omitempty"`
}
