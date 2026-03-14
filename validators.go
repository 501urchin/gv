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

	nv "github.com/501urchin/gv/internal/numeric"
	slv "github.com/501urchin/gv/internal/slice"
	sv "github.com/501urchin/gv/internal/string"
)

type validator interface {
	Validate() error
}

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
func First(v ...validator) error {
	for _, fn := range v {
		if err := fn.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Last runs the checks and returns on the last func that returns a error
func Last(v ...validator) (r error) {
	for _, fn := range v {
		if err := fn.Validate(); err != nil {
			r = err
		}
	}

	return r
}

// Join runs the checks and joins all errors into a single error
func Join(v ...validator) (r error) {
	for _, fn := range v {
		if err := fn.Validate(); err != nil {
			r = errors.Join(r, err)
		}
	}

	return r
}
