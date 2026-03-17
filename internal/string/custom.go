package string

import (
	"regexp"
	"unicode"

	"github.com/501urchin/gopt"
	gverrors "github.com/501urchin/gv/internal/errors"
)

// TODO: move away from regex based validation since its slow compared to the other methods
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

func isHex[T ~string](v T) bool {
	for _, c := range v {
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') && (c < 'A' || c > 'F') {
			return false
		}
	}
	return true
}

func (s *StringValidator[T]) Hex() *StringValidator[T] {
	if s.err != nil {
		return s
	}

	if !isHex(s.val) {
		s.err = gverrors.ErrNotHex
		return s
	}

	return s
}

func isAlpha[T ~string](v T) bool {
	for _, c := range v {
		if (c < '0' || c > '9') && (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return false
		}
	}
	return true
}
func (s *StringValidator[T]) Alpha() *StringValidator[T] {
	if s.err != nil {
		return s
	}

	if !isAlpha(s.val) {
		s.err = gverrors.ErrNotAlpha
		return s
	}

	return s
}

func (s *StringValidator[T]) UUID() *StringValidator[T] {
	if s.err != nil {
		return s
	}

	if len(s.val) != 36 {
		s.err = gverrors.ErrNotUUID
		return s
	}

	if !isHex(s.val[:8]) {
		s.err = gverrors.ErrNotUUID
		return s
	}
	if !isHex(s.val[9:13]) {
		s.err = gverrors.ErrNotUUID
		return s
	}

	if !isHex(s.val[14:18]) {
		s.err = gverrors.ErrNotUUID
		return s
	}

	if !isHex(s.val[19:23]) {
		s.err = gverrors.ErrNotUUID
		return s
	}

	if !isHex(s.val[24:32]) {
		s.err = gverrors.ErrNotUUID
		return s
	}

	return s
}

// func (s *StringValidator[T]) URL() *StringValidator[T] {
// 	_, err := url.Parse(string(s.val))
// 	if err != nil {
// 		s.err = gverrors.ErrNotURL
// 		return s
// 	}

// 	return s
// }
