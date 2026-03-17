// Package string implements methods for validating strings
package string

type StringValidator[T ~string] struct {
	val T
	err error
}

func NewStringValidator[T ~string](v T) *StringValidator[T] {
	return &StringValidator[T]{
		val: v,
		err: nil,
	}
}

func (s *StringValidator[T]) Validate(customErr ...error) error {
	return s.err
}
