package numeric

import gverrors "github.com/501urchin/gv/internal/errors"

func (n *NumericValidator[T]) Min(v T) *NumericValidator[T] {
	if n.val < v {
		n.err = gverrors.ErrMin
	}

	return n
}

func (n *NumericValidator[T]) Max(v T) *NumericValidator[T] {
	if n.val > v {
		n.err = gverrors.ErrMax
	}

	return n
}

func (n *NumericValidator[T]) Required() *NumericValidator[T] {
	if n.err != nil {
		return n
	}

	if n.val == 0 {
		n.err = gverrors.ErrRequired
	}

	return n
}

func (n *NumericValidator[T]) Negative() *NumericValidator[T] {
	if n.err != nil {
		return n
	}

	if n.val > 0 {
		n.err = gverrors.ErrNotNegative
	}

	return n
}
func (n *NumericValidator[T]) Equal(v T) *NumericValidator[T] {
	if n.err != nil {
		return n
	}

	if n.val != v {
		n.err = gverrors.ErrNotEqual
	}

	return n
}
func (n *NumericValidator[T]) NotEqual(v T) *NumericValidator[T] {
	if n.err != nil {
		return n
	}

	if n.val == v {
		n.err = gverrors.ErrEqual
	}

	return n
}

func (n *NumericValidator[T]) Positive() *NumericValidator[T] {
	if n.err != nil {
		return n
	}

	if n.val < 0 {
		n.err = gverrors.ErrNotPositive
	}

	return n
}
func (n *NumericValidator[T]) Custom(fn func(v T) error) *NumericValidator[T] {
	if n.err != nil {
		return n
	}

	if err := fn(n.val); err != nil {
		n.err = err
	}

	return n
}
func (n *NumericValidator[T]) Validate() error {
	return n.err
}
