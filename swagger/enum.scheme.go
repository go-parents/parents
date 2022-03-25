package swagger

type SchemeEnumOption string

// Used in Schema model
const (
	SchemeHttp            SchemeEnumOption = "http"
	SchemeHttps           SchemeEnumOption = "https"
	SchemeWebsocket       SchemeEnumOption = "ws"
	SchemeWebsocketSecure SchemeEnumOption = "wss"
)
