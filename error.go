/*
Package validation exposes some utility functions to validate strings and
numbers. ErrValidation is returned by validation functions.
*/
package validation

import "fmt"

// ErrValidation is the custom error returned by the Number* and String* family
// of functions.
type ErrValidation struct {
	Args    interface{}
	Code    string
	Field   string
	Message string
	Value   interface{}
}

func (err *ErrValidation) Error() string {
	if (*err).Value == nil {
		return fmt.Sprintf("%v: %v fails validation, %v", (*err).Code, (*err).Field, (*err).Message)
	}

	return fmt.Sprintf("%v: (%v, %v) fails validation, %v", (*err).Code, (*err).Field, (*err).Value, (*err).Message)
}

// NewError creates ErrValidation from code, message, field and value provided.
func NewError(code string, args interface{}, message string, field string, value interface{}) *ErrValidation {
	return &ErrValidation{args, code, field, message, value}
}
