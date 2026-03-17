package numeric

import (
	"math"

	gverrors "github.com/501urchin/gv/internal/errors"
	"github.com/501urchin/gv/internal/pkg"
)

func (n *NumericValidator[T]) OneOf(values ...T) *NumericValidator[T] {
	if n.err != nil {
		return n
	}

	for _, v := range values {
		if n.val == v {
			return n
		}
	}

	n.err = gverrors.ErrNotOneOF
	return n
}
func (n *NumericValidator[T]) NotOneOf(values ...T) *NumericValidator[T] {
	if n.err != nil {
		return n
	}

	for _, v := range values {
		if n.val == v {
			n.err = gverrors.ErrOneOf
			return n
		}
	}
	return n
}

func (n *NumericValidator[T]) Finite(customErr ...error) *NumericValidator[T] {
	if n.err != nil {
		return n
	}

	if math.IsInf(float64(n.val), 0) {
		n.err = pkg.DefaultOrCustomError(gverrors.ErrInfinite, customErr...)
	}

	return n
}
