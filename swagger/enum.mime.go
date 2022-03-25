package swagger

type MimeEnumOption string

// Used in the Schema and PathItem models
const (
	MimeJson MimeEnumOption = "application/json"
	MimeXml  MimeEnumOption = "application/xml"
)
