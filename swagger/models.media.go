package swagger

type Media struct {
	Schema   *Schema `json:",omitempty"`
	Example  interface{}
	Examples map[string]Example
	Encoding map[string]Encoding
}
