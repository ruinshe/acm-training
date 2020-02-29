package api

import "github.com/gogf/gf/frame/g"

// SuccessfulResponse - the response incidcating the action is successful.
var SuccessfulResponse = Response{
	Data: g.Map{"success": true},
}

// Error - the common API error object used for UI error displaying.
type Error struct {
	ErrorCode    interface{} `json:"error_code,omitempty"`
	ErrorMessage interface{} `json:"error_message,omitempty"`
}

// Response - the common API response for the system.
type Response struct {
	Error *Error      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}
