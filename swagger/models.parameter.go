package swagger

type Parameter struct {
    Name            string         `validate:"required"`
    In              InEnumOption   `validate:"required"`
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
