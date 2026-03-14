package string

import (
	"regexp"

	"github.com/501urchin/gopt"
	gverrors "github.com/501urchin/gv/internal/errors"
)

func parseOctet[T ~string](b T) (int, bool) {
	if len(b) == 0 || len(b) > 3 {
		return 0, false
	}

	n := 0
	for _, c := range b {
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}

	return n, n <= 255
}
func (s *StringValidator[T]) Ipv4() *StringValidator[T] {
	var o1, o2, o3 = -1, -1, -1

	for i, c := range s.val {
		if o1 == 0 && c == '.' {
			o1 = i
			continue
		}

		if o2 == 0 && c == '.' {
			o2 = i
			continue
		}

		if o3 == 0 && c == '.' {
			o3 = i
			continue
		}
	}

	if o1 == -1 || o2 == -1 || o3 == -1 {
		s.err = gverrors.ErrNotIpv4
		return s
	}

	if _, ok := parseOctet(s.val[0:o1]); !ok {
		s.err = gverrors.ErrNotIpv4
		return s
	}

	if _, ok := parseOctet(s.val[o1+1 : o2]); !ok {
		s.err = gverrors.ErrNotIpv4
		return s
	}

	if _, ok := parseOctet(s.val[o2+1 : o3]); !ok {
		s.err = gverrors.ErrNotIpv4
		return s
	}

	sLen := len(s.val)
	if sLen < o3+1 {
		s.err = gverrors.ErrNotIpv4
		return s
	}

	if _, ok := parseOctet(s.val[o3+1 : sLen]); !ok {
		s.err = gverrors.ErrNotIpv4
		return s
	}

	return s
}

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

func (s *StringValidator[T]) Ipv6() *StringValidator[T] { return s }

func (s *StringValidator[T]) Hex() *StringValidator[T]    { return s }
func (s *StringValidator[T]) URL() *StringValidator[T]    { return s }
func (s *StringValidator[T]) URI() *StringValidator[T]    { return s }
func (s *StringValidator[T]) Alpha() *StringValidator[T]  { return s }
func (s *StringValidator[T]) Lower() *StringValidator[T]  { return s }
func (s *StringValidator[T]) Upper() *StringValidator[T]  { return s }
func (s *StringValidator[T]) Base64() *StringValidator[T] { return s }
