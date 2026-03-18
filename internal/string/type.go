// Package string implements methods for validating strings
package string

type StringValidator[T ~string] struct {
	optional bool
	val      T
	err      error
}

func NewStringValidator[T ~string](v T) *StringValidator[T] {
	return &StringValidator[T]{
		optional: false,
		val:      v,
		err:      nil,
	}
}

func (s *StringValidator[T]) Validate(customErr ...error) error {
	if s.optional {
		s.err = nil
	}
	
	return s.err
}
