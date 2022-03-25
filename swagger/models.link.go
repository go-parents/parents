package swagger

type Link struct {
	OperationRef OperationRefString `json:",omitempty"`
	OperationId  OperationIdString  `json:",omitempty"`

	// todo: figure out how to handle this in go
	Parameters map[OperationIdString]interface{}
	// todo: figure out how to handle this in go
	RequestBody interface{}

	Description string `json:",omitempty"`

	Server *Server `json:",omitempty"`
}
