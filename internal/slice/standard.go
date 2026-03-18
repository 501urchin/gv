package slice

import (
	gverrors "github.com/501urchin/gv/internal/errors"
	"github.com/501urchin/gv/internal/pkg"
)

// Element loops over the slice and runs the custom validation func fn. fn needs to return a error
func (s *SliceValidator[T]) Element(fn func(e T) error) *SliceValidator[T] {
	if s.optional || s.err != nil {
		return s
	}

	for i := range s.val {
		if err := fn(s.val[i]); err != nil {
			s.err = err
		}
	}

	return s
}
func (s *SliceValidator[T]) Optional() *SliceValidator[T] {
	if s.val == nil {
		s.optional = true
	}

	return s
}
func (s *SliceValidator[T]) Required(customErr ...error) *SliceValidator[T] {
	s.optional = false
	if s.err != nil {
		return s
	}

	if s.val == nil {
		s.err = pkg.DefaultOrCustomError(gverrors.ErrIsNilOrEmpty, customErr...)
	}

	return s
}

func (s *SliceValidator[T]) Min(v int, customErr ...error) *SliceValidator[T] {
	if s.optional || s.err != nil {
		return s
	}

	if len(s.val) < v {
		s.err = pkg.DefaultOrCustomError(gverrors.ErrMin, customErr...)
	}

	return s
}

func (s *SliceValidator[T]) Max(v int, customErr ...error) *SliceValidator[T] {
	if s.optional || s.err != nil {
		return s
	}

	if len(s.val) > v {
		s.err = pkg.DefaultOrCustomError(gverrors.ErrMax, customErr...)
	}

	return s
}

func (s *SliceValidator[T]) Default(v []T) *SliceValidator[T] {
	if s.val == nil {
		s.val = v
	}

	return s
}

func (s *SliceValidator[T]) When(v bool, fn func(val []T) error) *SliceValidator[T] {
	if v {
		s.err = fn(s.val)
	}

	return s
}
