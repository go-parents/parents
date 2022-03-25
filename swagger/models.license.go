package swagger

type License struct {
	Name string `validate:"required"`
	Url  string `json:",omitempty"`
}
