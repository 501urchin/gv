package string

import (
	gverrors "github.com/501urchin/gv/internal/errors"
)

func parseIpv4Octet[T ~string](b T) (int, bool) {
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
	o := make([]int, 0, 3)
	for i, c := range s.val {
		if c == '.' {
			o = append(o, i)
		}
	}

	if len(o) != 3 {
		s.err = gverrors.ErrNotIpv4
		return s
	}

	if _, ok := parseIpv4Octet(s.val[0:o[0]]); !ok {
		s.err = gverrors.ErrNotIpv4
		return s
	}

	if _, ok := parseIpv4Octet(s.val[o[0]+1 : o[1]]); !ok {
		s.err = gverrors.ErrNotIpv4
		return s
	}

	if _, ok := parseIpv4Octet(s.val[o[1]+1 : o[2]]); !ok {
		s.err = gverrors.ErrNotIpv4
		return s
	}

	sLen := len(s.val)
	if sLen < o[2]+1 {
		s.err = gverrors.ErrNotIpv4
		return s
	}

	if _, ok := parseIpv4Octet(s.val[o[2]+1 : sLen]); !ok {
		s.err = gverrors.ErrNotIpv4
		return s
	}

	return s
}
