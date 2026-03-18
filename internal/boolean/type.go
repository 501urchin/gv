// Package boolean implements methods for validating booleans
package boolean

type boolValidator[T ~bool] struct {
	val T
	err error
}

func NewBoolValidator[T ~bool](val T) *boolValidator[T] {
	return &boolValidator[T]{
		val: val,
		err: nil,
	}
}

func (b *boolValidator[T]) Validate() error {
	return b.err
}
