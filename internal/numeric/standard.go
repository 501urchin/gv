package numeric

import (
	gverrors "github.com/501urchin/gv/internal/errors"
	"github.com/501urchin/gv/internal/pkg"
)

func (n *NumericValidator[T]) Optional() *NumericValidator[T] {
	if n.val == 0 {
		n.optional = true
	}

	return n
}
func (n *NumericValidator[T]) Min(v T, customErr ...error) *NumericValidator[T] {
	if n.optional || n.err != nil {
		return n
	}

	if n.val < v {
		n.err = pkg.DefaultOrCustomError(gverrors.ErrMin, customErr...)
	}

	return n
}

func (n *NumericValidator[T]) Max(v T, customErr ...error) *NumericValidator[T] {
	if n.optional || n.err != nil {
		return n
	}

	if n.val > v {
		n.err = pkg.DefaultOrCustomError(gverrors.ErrMax, customErr...)
	}

	return n
}

func (n *NumericValidator[T]) Required(customErr ...error) *NumericValidator[T] {
	n.optional = false
	if n.err != nil {
		return n
	}

	if n.val == 0 {
		n.err = pkg.DefaultOrCustomError(gverrors.ErrRequired, customErr...)
	}

	return n
}

func (n *NumericValidator[T]) Negative(customErr ...error) *NumericValidator[T] {
	if n.optional || n.err != nil {
		return n
	}

	if n.val > 0 {
		n.err = pkg.DefaultOrCustomError(gverrors.ErrNotNegative, customErr...)
	}

	return n
}

func (n *NumericValidator[T]) Positive(customErr ...error) *NumericValidator[T] {
	if n.optional || n.err != nil {
		return n
	}

	if n.val < 0 {
		n.err = pkg.DefaultOrCustomError(gverrors.ErrNotPositive, customErr...)
	}

	return n
}
func (n *NumericValidator[T]) Equal(v T, customErr ...error) *NumericValidator[T] {
	if n.optional || n.err != nil {
		return n
	}

	if n.val != v {
		n.err = pkg.DefaultOrCustomError(gverrors.ErrNotEqual, customErr...)
	}

	return n
}
func (n *NumericValidator[T]) NotEqual(v T, customErr ...error) *NumericValidator[T] {
	if n.optional || n.err != nil {
		return n
	}

	if n.val == v {
		n.err = pkg.DefaultOrCustomError(gverrors.ErrEqual, customErr...)
	}

	return n
}

func (n *NumericValidator[T]) Custom(fn func(v T) error) *NumericValidator[T] {
	if n.optional || n.err != nil {
		return n
	}

	if err := fn(n.val); err != nil {
		n.err = err
	}

	return n
}

func (n *NumericValidator[T]) Default(v T) *NumericValidator[T] {
	if n.val == 0 {
		n.val = v
	}

	return n
}

func (n *NumericValidator[T]) When(v bool, fn func(val T) error) *NumericValidator[T] {
	if v {
		n.err = fn(n.val)
	}

	return n
}
