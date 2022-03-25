package swagger

type Server struct {
	Url         UrlString                      `validate:"required"`
	Description string                         `json:",omitempty"`
	Variables   map[VariableKey]ServerVariable `json:",omitempty"`
}
