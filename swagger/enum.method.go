package swagger

type MethodEnumOption string

// Used in the Path model
const (
	MethodGet     MethodEnumOption = "get"
	MethodPut     MethodEnumOption = "put"
	MethodPost    MethodEnumOption = "post"
	MethodDelete  MethodEnumOption = "delete"
	MethodOptions MethodEnumOption = "options"
	MethodHead    MethodEnumOption = "head"
	MethodPatch   MethodEnumOption = "patch"
	MethodTrace   MethodEnumOption = "trace"
)
