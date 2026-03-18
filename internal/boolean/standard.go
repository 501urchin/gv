package boolean

import (
	gverrors "github.com/501urchin/gv/internal/errors"
	"github.com/501urchin/gv/internal/pkg"
)

func (b *boolValidator[T]) IsFalse(customErr ...error) *boolValidator[T] {
	if b.err != nil {
		return b
	}

	if b.val {
		b.err = pkg.DefaultOrCustomError(gverrors.ErrNotFalse, customErr...)
	}

	return b
}
func (b *boolValidator[T]) IsTrue(customErr ...error) *boolValidator[T] {
	if b.err != nil {
		return b
	}

	if !b.val {
		b.err = pkg.DefaultOrCustomError(gverrors.ErrNotTrue, customErr...)
	}

	return b
}

func (b *boolValidator[T]) Default(v T) *boolValidator[T] {
	if !b.val {
		b.val = v
	}

	return b
}

func (b *boolValidator[T]) When(v bool, fn func(val T) error, customErr ...error) *boolValidator[T] {
	if b.err != nil || !v {
		return b
	}

	if err := fn(b.val); err != nil {
		b.err = pkg.DefaultOrCustomError(err, customErr...)
	}

	return b
}

func (b *boolValidator[T]) Equals(v T, customErr ...error) *boolValidator[T] {
	if b.err != nil {
		return b
	}

	if b.val != v {
		b.err = pkg.DefaultOrCustomError(gverrors.ErrNotEqual, customErr...)
	}

	return b
}
func (b *boolValidator[T]) NotEquals(v T, customErr ...error) *boolValidator[T] {
	if b.err != nil {
		return b
	}

	if b.val == v {
		b.err = pkg.DefaultOrCustomError(gverrors.ErrEqual, customErr...)
	}

	return b
}
