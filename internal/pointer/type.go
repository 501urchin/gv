// Package pointer implements methods for validating pointers
package pointer

type pointerValidator[T any] struct {
	optional bool
	val      *T
	err      error
}

func NewPointerValidator[T any](val *T) *pointerValidator[T] {
	return &pointerValidator[T]{
		val: val,
		err: nil,
	}
}

func (p *pointerValidator[T]) Validate() error {
	if p.optional {
		p.err = nil
	}
	
	return p.err
}
