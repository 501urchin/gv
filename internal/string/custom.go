package string

import (
	"regexp"
	"unicode"

	"github.com/501urchin/gopt"
	gverrors "github.com/501urchin/gv/internal/errors"
	"github.com/501urchin/gv/internal/pkg"
)

// TODO: move away from regex based validation since its slow compared to the other methods
var emailReg = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func (s *StringValidator[T]) Email(customErr ...error) *StringValidator[T] {
	if s.err != nil {
		return s
	}

	if !emailReg.Match(gopt.StringToBytes(string(s.val))) {
		s.err = pkg.DefaultOrCustomError(gverrors.ErrEmail, customErr...)
	}

	return s
}

func (s *StringValidator[T]) Lower(customErr ...error) *StringValidator[T] {
	if s.err != nil {
		return s
	}

	for _, c := range s.val {
		if unicode.IsUpper(c) {
			s.err = pkg.DefaultOrCustomError(gverrors.ErrUpper, customErr...)
			return s
		}
	}

	return s
}

func (s *StringValidator[T]) Upper(customErr ...error) *StringValidator[T] {
	if s.err != nil {
		return s
	}

	for _, c := range s.val {
		if unicode.IsLower(c) {
			s.err = pkg.DefaultOrCustomError(gverrors.ErrLower, customErr...)
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

func (s *StringValidator[T]) Hex(customErr ...error) *StringValidator[T] {
	if s.err != nil {
		return s
	}

	if !isHex(s.val) {
		s.err = pkg.DefaultOrCustomError(gverrors.ErrNotHex, customErr...)
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
func (s *StringValidator[T]) Alpha(customErr ...error) *StringValidator[T] {
	if s.err != nil {
		return s
	}

	if !isAlpha(s.val) {
		s.err = pkg.DefaultOrCustomError(gverrors.ErrNotAlpha, customErr...)
		return s
	}

	return s
}

func (s *StringValidator[T]) UUID(customErr ...error) *StringValidator[T] {
	if s.err != nil {
		return s
	}

	if len(s.val) != 36 {
		goto errorCase
	}

	if !isHex(s.val[:8]) {
		goto errorCase
	}
	if !isHex(s.val[9:13]) {
		goto errorCase
	}

	if !isHex(s.val[14:18]) {
		goto errorCase
	}

	if !isHex(s.val[19:23]) {
		goto errorCase
	}

	if !isHex(s.val[24:32]) {
		goto errorCase
	}

	return s

errorCase:
	s.err = pkg.DefaultOrCustomError(gverrors.ErrNotUUID, customErr...)
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
