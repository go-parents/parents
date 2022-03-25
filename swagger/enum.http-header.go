package swagger

type HttpHeaderEnumOption string

// Used in the Response model
const (
	HeaderCacheControl HttpHeaderEnumOption = "Cache-Control"
	HeaderEtag         HttpHeaderEnumOption = "ETag"

	// todo: expand this list with the most common http headers for RestFul JSON APIs
)
