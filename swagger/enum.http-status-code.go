package swagger

type HttpStatusCodeEnumOption string

// Used in the Path model
const (
	// StatusDefault used as a default response object for all HTTP codes that are not covered individually
	StatusDefault      HttpStatusCodeEnumOption = "default"
	StatusOkay         HttpStatusCodeEnumOption = "200"
	StatusNotFound     HttpStatusCodeEnumOption = "404"
	StatusBadRequest   HttpStatusCodeEnumOption = "400"
	StatusUnauthorized HttpStatusCodeEnumOption = "401"
	StatusForbidden    HttpStatusCodeEnumOption = "403"
	StatusException    HttpStatusCodeEnumOption = "500"
)
