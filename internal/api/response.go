package api

// Response - the common API response for the system.
type Response struct {
	ErrorCode    interface{} `json:"error_code,omitempty"`
	ErrorMessage interface{} `json:"error_message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}
