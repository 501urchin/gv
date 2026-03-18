package string

import (
	gverrors "github.com/501urchin/gv/internal/errors"
	"github.com/501urchin/gv/internal/pkg"
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

func (s *StringValidator[T]) Ipv4(customErr ...error) *StringValidator[T] {
	if s.optional || s.err != nil {
		return s
	}

	o := make([]int, 0, 3)
	for i, c := range s.val {
		if c == '.' {
			o = append(o, i)
		}
	}

	sLen := len(s.val)

	if len(o) != 3 {
		goto errorCase
	}

	if sLen < o[2]+1 {
		goto errorCase
	}

	if _, ok := parseIpv4Octet(s.val[0:o[0]]); !ok {
		goto errorCase
	}

	if _, ok := parseIpv4Octet(s.val[o[0]+1 : o[1]]); !ok {
		goto errorCase
	}

	if _, ok := parseIpv4Octet(s.val[o[1]+1 : o[2]]); !ok {
		goto errorCase
	}

	if _, ok := parseIpv4Octet(s.val[o[2]+1 : sLen]); !ok {
		goto errorCase
	}

	return s
errorCase:
	s.err = pkg.DefaultOrCustomError(gverrors.ErrNotIpv4, customErr...)
	return s

}
