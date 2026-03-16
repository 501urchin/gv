package gv

import gverrors "github.com/501urchin/gv/internal/errors"

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
