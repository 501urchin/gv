package string

import (
	"regexp"
	"unicode"

	"github.com/501urchin/gopt"
	gverrors "github.com/501urchin/gv/internal/errors"
)

var emailReg = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func (s *StringValidator[T]) Email() *StringValidator[T] {
	if s.err != nil {
		return s
	}

	if !emailReg.Match(gopt.StringToBytes(string(s.val))) {
		s.err = gverrors.ErrEmail
	}

	return s
}

func (s *StringValidator[T]) Lower() *StringValidator[T] {
	if s.err != nil {
		return s
	}

	for _, c := range s.val {
		if unicode.IsUpper(c) {
			s.err = gverrors.ErrUpper
			return s
		}
	}

	return s
}

func (s *StringValidator[T]) Upper() *StringValidator[T] {
	if s.err != nil {
		return s
	}

	for _, c := range s.val {
		if unicode.IsLower(c) {
			s.err = gverrors.ErrLower
			return s
		}
	}

	return s
}

func (s *StringValidator[T]) Hex() *StringValidator[T] {
	if s.err != nil {
		return s
	}

	for _, c := range s.val {
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') && (c < 'A' || c > 'F') {
			s.err = gverrors.ErrNotHex
			return s
		}
	}

	return s
}

// func (s *StringValidator[T]) UUID() *StringValidator[T]   { return s }
// func (s *StringValidator[T]) URL() *StringValidator[T]    { return s }
// func (s *StringValidator[T]) Alpha() *StringValidator[T]  { return s }
// func (s *StringValidator[T]) Base64() *StringValidator[T] { return s }
