// package gverrors contains errors that are returned by the validators
package gverrors

import "errors"

var (
	ErrRequired      = errors.New("required field is empty")
	ErrMin           = errors.New("field is smaller than allowed minimum")
	ErrMax           = errors.New("field is bigger than allowed maximum")
	ErrContains      = errors.New("field contains a value that is not allowed")
	ErrHasWhitespace = errors.New("field contains whitespace")
	ErrUpper         = errors.New("field contains uppercased letters")
	ErrLower         = errors.New("field contains lowercased letters")
)
