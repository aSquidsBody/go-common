package response

import (
	"encoding/json"
)

// errorResponse is the standard struct for API error responses
type errorResponse struct {
	Code    int      `json:"code,omitempty"`
	Message string   `json:"message,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

// Write error
func writeError(w writer, code int, message string, errs ...error) {
	errStr := make([]string, len(errs))
	for i := range errs {
		errStr[i] = errs[i].Error()
	}

	body := errorResponse{
		Code:    code,
		Message: message,
		Errors:  errStr,
	}

	write(w, code, body)
}

// Write json marshal failure in response attempt
func writeMarshalError(w writer, err error) {
	body := errorResponse{
		Code:    500,
		Message: "Internal Server Error",
		Errors:  []string{err.Error()},
	}

	b, _ := json.Marshal(body)
	w.WriteHeader(500)
	w.Write(b)
}
