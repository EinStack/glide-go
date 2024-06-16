package glide

import "fmt"

// Error that may occur during the processing of API request.
type Error struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Status  int    `json:"status,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Name, e.Message)
}
