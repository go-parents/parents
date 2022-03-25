package swagger

type Parameter struct {
	Name            string `validate:"required"`
	In              string `validate:"required" enum:"query,header,path,cookie"`
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
