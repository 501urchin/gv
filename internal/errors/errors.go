// Package gverrors contains errors that are returned by the validators
package gverrors

import "errors"

var (
	ErrRequired      = errors.New("required field is empty or contains default value")
	ErrMin           = errors.New("field is smaller than allowed minimum")
	ErrMax           = errors.New("field is bigger than allowed maximum")
	ErrMustContain   = errors.New("field does not contains a required value")
	ErrNotContains   = errors.New("field contains a value that is not allowed")
	ErrHasWhitespace = errors.New("field contains whitespace")
	ErrUpper         = errors.New("field contains uppercased letters")
	ErrLower         = errors.New("field contains lowercased letters")
	ErrNotPositive   = errors.New("field must be a positive number")
	ErrNotNegative   = errors.New("field must be a negative number")
	ErrNotEqual      = errors.New("field does not equal value")
	ErrEqual         = errors.New("field must not equal value")
	ErrNotOneOF      = errors.New("field is not included in set")
	ErrOneOf         = errors.New("field is included in not allowed set")
	ErrInfinite      = errors.New("field must not be infinite")
	ErrEmail         = errors.New("field is not a valid email")
	ErrNotIpv4       = errors.New("field is not a valid ipv4 address")
	ErrIsNil         = errors.New("field is nil")
	ErrEmpty         = errors.New("field is empty")
	ErrNotHex        = errors.New("field is not hex")
	ErrNotAlpha      = errors.New("field is not alphanumerical")
	ErrNotUUID      = errors.New("field is not valid UUID")
)
