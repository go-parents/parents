package swagger

type StyleEnumOption string

// Used in Header and Parameter models
const (
	StyleMatrix         StyleEnumOption = "matrix"
	StyleLabel          StyleEnumOption = "label"
	StyleForm           StyleEnumOption = "form"
	StyleSimple         StyleEnumOption = "simple"
	StyleSpaceDelimited StyleEnumOption = "spaceDelimited"
	StylePipeDelimited  StyleEnumOption = "pipeDelimited"
	StyleDeepObject     StyleEnumOption = "deepObject"
)
