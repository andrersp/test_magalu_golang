package genericresponse

import "fmt"

type Error struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Message: %s Detail: %s", e.Message, e.Detail)
}

func NewErrorResponse(message string, detail string) error {
	err := new(Error)
	err.Detail = detail
	err.Message = message

	return err
}
