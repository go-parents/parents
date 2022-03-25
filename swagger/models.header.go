package swagger

// Header Same as Parameter just excluding "name" and "in" fields
type Header struct {
	// Name omitted as per spec
	// In omitted as per spec
	Description     string `json:",omitempty"`
	Required        bool
	Deprecated      bool
	AllowEmptyValue bool

	Style         string `json:",omitempty" enum:"matrix,label,form,simple,spaceDelimited,pipeDelimited,deepObject"`
	Explode       bool
	AllowReserved bool
	Schema        *Schema            `json:",omitempty"`
	Example       interface{}        `json:",omitempty"`
	Examples      map[string]Example `json:",omitempty"`

	Content map[string]Media
}
