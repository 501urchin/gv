package string

import (
	"errors"
	"testing"

	gverrors "github.com/501urchin/gv/internal/errors"
)

func TestCustom(t *testing.T) {
	t.Run("email", func(t *testing.T) {
		err := NewStringValidator("email@email.com").Email().Validate()
		if err != nil {
			t.Error("func threw error when it wasnt supposed to")
		}

		err = NewStringValidator("@email.com").Email().Validate()
		if !errors.Is(err, gverrors.ErrEmail) {
			t.Errorf("excpected %v but got %v", gverrors.ErrEmail, err)
		}
	})

	t.Run("lower", func(t *testing.T) {
		err := NewStringValidator("abcd45").Lower().Validate()
		if err != nil {
			t.Error("func threw error when it wasnt supposed to")
		}

		err = NewStringValidator("@FCDSS").Lower().Validate()
		if !errors.Is(err, gverrors.ErrUpper) {
			t.Errorf("excpected %v but got %v", gverrors.ErrUpper, err)
		}
	})
	t.Run("upper", func(t *testing.T) {
		err := NewStringValidator("SDJKSHFD23").Upper().Validate()
		if err != nil {
			t.Error("func threw error when it wasnt supposed to")
		}

		err = NewStringValidator("kcjnm").Upper().Validate()
		if !errors.Is(err, gverrors.ErrLower) {
			t.Errorf("excpected %v but got %v", gverrors.ErrLower, err)
		}
	})
	t.Run("hex", func(t *testing.T) {
		err := NewStringValidator("abcd45").Hex().Validate()
		if err != nil {
			t.Error("func threw error when it wasnt supposed to")
		}

		err = NewStringValidator("@3d5r").Hex().Validate()
		if !errors.Is(err, gverrors.ErrNotHex) {
			t.Errorf("excpected %v but got %v", gverrors.ErrNotHex, err)
		}
	})
	t.Run("alpha", func(t *testing.T) {
		err := NewStringValidator("abcdefghijklmnopqrstuvwxyz12345678").Alpha().Validate()
		if err != nil {
			t.Error("func threw error when it wasnt supposed to")
		}

		err = NewStringValidator("@3d5r").Alpha().Validate()
		if !errors.Is(err, gverrors.ErrNotAlpha) {
			t.Errorf("excpected %v but got %v", gverrors.ErrNotAlpha, err)
		}
	})

}

func BenchmarkEmail(b *testing.B) {
	var err error
	for b.Loop() {
		err = NewStringValidator("jalzzlk@email.com").Email().Validate()
	}

	_ = err
}
