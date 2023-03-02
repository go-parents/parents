package swagger

// Header Same as Parameter just excluding "name" and "in" fields
type Header struct {
	// Name omitted as per spec
	// In omitted as per spec
	Description     MarkdownString `json:",omitempty"`
	Required        bool
	Deprecated      bool
	AllowEmptyValue bool

	Style         StyleEnumOption `json:",omitempty"`
	Explode       bool
	AllowReserved bool
	Schema        *Schema            `json:",omitempty"`
	Example       interface{}        `json:",omitempty"`
	Examples      map[string]Example `json:",omitempty"`

	Content map[string]Media
}
