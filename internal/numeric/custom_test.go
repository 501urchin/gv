package numeric

import (
	"errors"
	"math"
	"testing"

	gverrors "github.com/501urchin/gv/internal/errors"
)

func TestCustom(t *testing.T) {
	t.Run("one of", func(t *testing.T) {
		err := NewNumericValidator(2).OneOf(1, 2, 3).Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewNumericValidator(5).OneOf(1, 2, 3).Validate()
		if !errors.Is(err, gverrors.ErrNotOneOF) {
			t.Error("func didnt return ErrNotOneOF while it was supposed to")
		}
	})

	t.Run("not one of", func(t *testing.T) {
		err := NewNumericValidator(5).NotOneOf(1, 2, 3).Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewNumericValidator(2).NotOneOf(1, 2, 3).Validate()
		if !errors.Is(err, gverrors.ErrOneOf) {
			t.Error("func didnt return ErrOneOf while it was supposed to")
		}
	})

	t.Run("finite", func(t *testing.T) {
		err := NewNumericValidator(5.0).Finite().Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewNumericValidator(math.Inf(1)).Finite().Validate()
		if !errors.Is(err, gverrors.ErrInfinite) {
			t.Error("func didnt return ErrInfinite while it was supposed to")
		}
	})
}
