// Package slice implements methods for validating slices
package slice

type SliceValidator[T any] struct {
	optional bool
	val      []T
	err      error
}

func NewSliceValidator[T any](v []T) *SliceValidator[T] {
	return &SliceValidator[T]{
		val: v,
		err: nil,
	}
}

func (s *SliceValidator[T]) Validate() error {
	if s.optional {
		s.err = nil
	}
	
	return s.err
}
