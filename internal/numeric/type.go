// Package numeric implements methods for numerics
package numeric

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type NumericValidator[T Numeric] struct {
	val T
	err error
}

func NewNumericValidator[T Numeric](v T) *NumericValidator[T] {
	return &NumericValidator[T]{
		val: v,
		err: nil,
	}
}
