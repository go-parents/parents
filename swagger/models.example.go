package swagger

type Example struct {
	Summary       string      `json:",omitempty"`
	Description   string      `json:",omitempty"`
	Value         interface{} `json:",omitempty"`
	ExternalValue string      `json:",omitempty"`
}
