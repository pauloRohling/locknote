package error

type HTTPError struct {
	Type      string `json:"type,omitempty"`      // URL of the error type. Can be used to lookup the error in a documentation
	Title     string `json:"title,omitempty"`     // Short title of the error
	Status    int    `json:"status,omitempty"`    // HTTP status code. Example: 403
	Detail    string `json:"detail,omitempty"`    // Human-readable error message
	Instance  string `json:"instance,omitempty"`  // URL to the request that produced the error
	RequestID string `json:"requestId,omitempty"` // ID of the request that produced the error
}
