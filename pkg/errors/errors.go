package errors

import "fmt"

//FieldValidationError field related errors
type FieldValidationError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

func (e FieldValidationError) Error() string {
	return fmt.Sprintf("FieldValidationError: field -> %s, msg -> %s", e.Field, e.Msg)
}

//NonFieldValidationError Non field related errors
type NonFieldValidationError struct {
	Msg string `json:"msg"`
}

func (e NonFieldValidationError) Error() string {
	return fmt.Sprintf("NonFieldValidationError: msg -> %s", e.Msg)
}
