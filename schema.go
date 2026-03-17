package gv

import gverrors "github.com/501urchin/gv/internal/errors"

type schemaFunc[T any] func(d *T) error

type schemaValidator[T any] struct {
	fn schemaFunc[T]
}

/*
Schema lets you define reusable validation logic.

Example:
	schema := Schema(func(info *T) error {
		return First(
			String(info.Field1).Required().Min(3).Max(25).Validate(),
			String(info.Field2).Required().NoWhitespace().Validate(),
		)
	})

*/
func Schema[T any](fn schemaFunc[T]) *schemaValidator[T] {
	return &schemaValidator[T]{fn: fn}
}

// Validate method runs the validation schema on the data pointer
// 	err := schema.Validate(dataPointer)
func (s *schemaValidator[T]) Validate(data *T) error {
	if data == nil {
		return gverrors.ErrIsNilOrEmpty
	}
	return s.fn(data)
}
