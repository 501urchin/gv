package string

import (
	"strings"
	"unicode"

	gverrors "github.com/501urchin/gv/internal/errors"
	"github.com/501urchin/gv/internal/pkg"
)

func (s *StringValidator[T]) Required(customErr ...error) *StringValidator[T] {
	s.optional = false
	if s.err != nil {
		return s
	}

	if len(s.val) == 0 {
		s.err = pkg.DefaultOrCustomError(gverrors.ErrRequired, customErr...)
	}

	return s
}
func (s *StringValidator[T]) Optional() *StringValidator[T] {
	if len(s.val) == 0 {
		s.optional = true
	}
	
	return s
}

func (s *StringValidator[T]) Min(v int, customErr ...error) *StringValidator[T] {
	if s.optional || s.err != nil {
		return s
	}

	if len(s.val) < v {
		s.err = pkg.DefaultOrCustomError(gverrors.ErrMin, customErr...)
	}

	return s
}

func (s *StringValidator[T]) Max(v int, customErr ...error) *StringValidator[T] {
	if s.optional || s.err != nil {
		return s
	}

	if len(s.val) > v {
		s.err = pkg.DefaultOrCustomError(gverrors.ErrMax, customErr...)
	}

	return s
}

func (s *StringValidator[T]) MustContain(v T, customErr ...error) *StringValidator[T] {
	if s.optional || s.err != nil {
		return s
	}

	if !strings.Contains(string(s.val), string(v)) {
		s.err = pkg.DefaultOrCustomError(gverrors.ErrMustContain, customErr...)
	}

	return s
}
func (s *StringValidator[T]) CantContain(v T, customErr ...error) *StringValidator[T] {
	if s.optional || s.err != nil {
		return s
	}

	if strings.Contains(string(s.val), string(v)) {
		s.err = pkg.DefaultOrCustomError(gverrors.ErrNotContains, customErr...)
	}

	return s
}
func (s *StringValidator[T]) ContainsAny(v string, customErr ...error) *StringValidator[T] {
	if s.optional || s.err != nil {
		return s
	}

	if !strings.ContainsAny(string(s.val), string(v)) {
		s.err = pkg.DefaultOrCustomError(gverrors.ErrNotContains, customErr...)
	}

	return s
}
func (s *StringValidator[T]) NoWhitespace(customErr ...error) *StringValidator[T] {
	if s.optional || s.err != nil {
		return s
	}

	for _, r := range s.val {
		if unicode.IsSpace(r) {
			s.err = pkg.DefaultOrCustomError(gverrors.ErrHasWhitespace, customErr...)
			return s
		}
	}

	return s
}
func (s *StringValidator[T]) Custom(fn func(val T) error) *StringValidator[T] {
	if s.optional || s.err != nil {
		return s
	}

	if err := fn(s.val); err != nil {
		s.err = err
	}

	return s
}

func (s *StringValidator[T]) Default(v T) *StringValidator[T] {
	if len(s.val) == 0 {
		s.val = v
	}

	return s
}

func (s *StringValidator[T]) When(v bool, fn func(val T) error) *StringValidator[T] {
	if v {
		s.err = fn(s.val)
	}

	return s
}
