// Package slice implements methods for validating slices
package slice

type SliceValidator[T any] struct {
	val []T
	err error
}

func NewSliceValidator[T any](v []T) *SliceValidator[T] {
	return &SliceValidator[T]{
		val: v,
		err: nil,
	}
}
