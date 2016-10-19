package twitter

// APIError .
type APIError struct {
	Errors []ErrorDetail `json:"errors"`
}

// ErrorDetail .
type ErrorDetail struct {
	Message string `json:"message"`
	Core    int    `json:"code"`
}
