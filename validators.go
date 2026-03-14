// Package gv contains type specific validator to make your life easier.
// Each validator returns itself allowing you to chain methods.
// Validation stops at the first failure and returns a error when Validate is called.
// 
// Example:
// 	v := "bad"
// 	err := gv.String(v).Required().Min(5).Max(25).Contains("b").Validate()
//	if err != nil {
//		// handle validation error
//	}
package gv

import (
	is "github.com/501urchin/gv/internal/string"
)

func String[T ~string](val T) *is.StringValidator[T] {
	return is.NewStringValidator(val)
}
