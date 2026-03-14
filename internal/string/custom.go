package string

import (
	"regexp"

	"github.com/501urchin/gopt"
	gverrors "github.com/501urchin/gv/internal/errors"
)



func (s *StringValidator[T]) UUID() *StringValidator[T] {
	return s
}

var emailReg = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func (s *StringValidator[T]) Email() *StringValidator[T] {
	if !emailReg.Match(gopt.StringToBytes(string(s.val))) {
		s.err = gverrors.ErrEmail
	}

	return s
}


func (s *StringValidator[T]) Hex() *StringValidator[T]    { return s }
func (s *StringValidator[T]) URL() *StringValidator[T]    { return s }
func (s *StringValidator[T]) URI() *StringValidator[T]    { return s }
func (s *StringValidator[T]) Alpha() *StringValidator[T]  { return s }
func (s *StringValidator[T]) Lower() *StringValidator[T]  { return s }
func (s *StringValidator[T]) Upper() *StringValidator[T]  { return s }
func (s *StringValidator[T]) Base64() *StringValidator[T] { return s }
