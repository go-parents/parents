package swagger

type Encoding struct {
	ContentType   string            `json:",omitempty"`
	Headers       map[string]Header `json:",omitempty"`
	Style         string            `json:",omitempty"`
	Explode       bool
	AllowReserved bool
}
