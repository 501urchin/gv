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

type validator interface {
	Validate() error
}

func Group(v ...validator) error {
	for _, v := range v {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}
