// Package gv contains type specific validator to make your life easier.
// Each validator returns itself allowing you to chain methods.
// Validation stops at the first failure and returns a error when Validate is called.
//
// Example:
//
//	v := "bad"
//	err := gv.String(v).Required().Min(5).Max(25).Validate()
//	if err != nil {
//		// handle validation error
//	}
package gv

import (
	"errors"

	gverrors "github.com/501urchin/gv/internal/errors"
	nv "github.com/501urchin/gv/internal/numeric"
	slv "github.com/501urchin/gv/internal/slice"
	sv "github.com/501urchin/gv/internal/string"
)

func String[T ~string](val T) *sv.StringValidator[T] {
	return sv.NewStringValidator(val)
}
func Slice[T any](val []T) *slv.SliceValidator[T] {
	return slv.NewSliceValidator(val)
}
func Numeric[T nv.Numeric](val T) *nv.NumericValidator[T] {
	return nv.NewNumericValidator(val)
}

// First runs the checks and returns on the first func that returns a error
func First(v ...error) error {
	for _, fn := range v {
		if err := fn; err != nil {
			return err
		}
	}

	return nil
}

// Last runs the checks and returns on the last func that returns a error
func Last(v ...error) (r error) {
	for _, fn := range v {
		if err := fn; err != nil {
			r = err
		}
	}

	return r
}

// Join runs the checks and joins all errors into a single error
func Join(v ...error) (r error) {
	for _, fn := range v {
		if err := fn; err != nil {
			r = errors.Join(r, err)
		}
	}

	return r
}

type schemaFunc[T any] func(d *T) error

type schemaValidator[T any] struct {
	fn schemaFunc[T]
}

func Schema[T any](fn schemaFunc[T]) *schemaValidator[T] {
	return &schemaValidator[T]{
		fn: fn,
	}
}

func (s *schemaValidator[T]) Validate(data *T) error {
	if data == nil {
		return gverrors.ErrIsNilOrEmpty
	}

	return s.fn(data)
}
