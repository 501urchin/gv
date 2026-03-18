package pointer

import (
	gverrors "github.com/501urchin/gv/internal/errors"
	"github.com/501urchin/gv/internal/pkg"
)

func (p *pointerValidator[T]) Optional() *pointerValidator[T] {
	p.optional = true

	return p
}
func (p *pointerValidator[T]) NotNil(customErr ...error) *pointerValidator[T] {
	if p.optional || p.err != nil {
		return p
	}

	if p.val == nil {
		p.err = pkg.DefaultOrCustomError(gverrors.ErrIsNil, customErr...)
	}

	return p
}
func (p *pointerValidator[T]) Nil(customErr ...error) *pointerValidator[T] {
	if p.optional || p.err != nil {
		return p
	}

	if p.val != nil {
		p.err = pkg.DefaultOrCustomError(gverrors.ErrIsNotNil, customErr...)
	}

	return p
}

// Deref Will try to dereference the pointer. If the pointer is nil it will skip the validation func
func (p *pointerValidator[T]) Deref(fn func(v T) error) *pointerValidator[T] {
	if p.val == nil {
		return p
	}

	p.err = fn(*p.val)
	return p
}

func (p *pointerValidator[T]) Default(v T) *pointerValidator[T] {
	if p.val == nil {
		p.val = &v
	}

	return p
}
func (p *pointerValidator[T]) When(v bool, fn func(v *T) error) *pointerValidator[T] {
	if v {
		p.err = fn(p.val)
	}

	return p
}
