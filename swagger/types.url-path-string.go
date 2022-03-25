package swagger

// UrlPathString represents the path portion of a URL
// PathItem variable placeholders as defined with curly braces.
// example `/users/{id}` defines a path placeholder called `id`
// Remember that you must still define parameters for any path variables, this is simply the placeholder definition.
type UrlPathString string
