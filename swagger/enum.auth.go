package swagger

type AuthEnumOption string

// Used in the Schema model
const (
	AuthBasic  AuthEnumOption = "basic"
	AuthApiKey AuthEnumOption = "apiKey"
	AuthOauth2 AuthEnumOption = "oauth2"
)
