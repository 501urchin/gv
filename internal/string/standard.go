package string

import (
	"strings"
	"unicode"

	gverrors "github.com/501urchin/gv/internal/errors"
)

func (s *StringValidator[T]) Required() *StringValidator[T] {
	if s.err != nil {
		return s
	}

	if len(s.val) == 0 {
		s.err = gverrors.ErrRequired
	}

	return s
}

func (s *StringValidator[T]) Min(v int) *StringValidator[T] {
	if s.err != nil {
		return s
	}

	if len(s.val) < v {
		s.err = gverrors.ErrMin
	}

	return s
}

func (s *StringValidator[T]) Max(v int) *StringValidator[T] {
	if s.err != nil {
		return s
	}

	if len(s.val) > v {
		s.err = gverrors.ErrMax
	}

	return s
}
func (s *StringValidator[T]) MustContain(v T) *StringValidator[T] {
	if s.err != nil {
		return s
	}

	if !strings.Contains(string(s.val), string(v)) {
		s.err = gverrors.ErrMustContain
	}

	return s
}
func (s *StringValidator[T]) NotContains(v T) *StringValidator[T] {
	if s.err != nil {
		return s
	}

	if strings.Contains(string(s.val), string(v)) {
		s.err = gverrors.ErrNotContains
	}

	return s
}
func (s *StringValidator[T]) NoWhitespace() *StringValidator[T] {
	if s.err != nil {
		return s
	}

	for _, r := range s.val {
		if unicode.IsSpace(r) {
			s.err = gverrors.ErrHasWhitespace
			return s
		}
	}

	return s
}
func (s *StringValidator[T]) Custom(fn func(val T) error) *StringValidator[T] {
	if s.err != nil {
		return s
	}

	if err := fn(s.val); err != nil {
		s.err = err
	}

	return s
}
func (s *StringValidator[T]) Validate() error {
	return s.err
}
