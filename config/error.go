package config

import (
	"fmt"
	"net/http"
)

// Error that may occur during the processing of API request.
type Error struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Status  int    `json:"status,omitempty"`
}

// NewError instantiates a default Error.
func NewError() error {
	return &Error{
		Name:    "unrecognized_error",
		Message: "There is no response body and the status code is in the range of 400-599.",
		Status:  http.StatusInternalServerError,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Name, e.Message)
}
