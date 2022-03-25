package swagger

type ServerVariable struct {
	Enum        []string `json:",omitempty"`
	Default     string   `validate:"required"`
	Description string   `json:",omitempty"`
}
