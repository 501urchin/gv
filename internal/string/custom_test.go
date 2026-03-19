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
	t.Run("uuid", func(t *testing.T) {
		err := NewStringValidator("f47ac10b-58cc-4372-a567-0e02b2c3d479").UUID().Validate()
		if err != nil {
			t.Error("func threw error when it wasnt supposed to")
		}

		err = NewStringValidator("@3d5r").UUID().Validate()
		if !errors.Is(err, gverrors.ErrNotUUID) {
			t.Errorf("excpected %v but got %v", gverrors.ErrNotUUID, err)
		}
	})
	t.Run("has upper", func(t *testing.T) {
		err := NewStringValidator("Ad").HasUpper().Validate()
		if err != nil {
			t.Error("func threw error when it wasnt supposed to")
		}

		err = NewStringValidator("kshjd").HasUpper().Validate()
		if !errors.Is(err, gverrors.ErrNoUpper) {
			t.Errorf("excpected %v but got %v", gverrors.ErrNoUpper, err)
		}
	})
	t.Run("has lower", func(t *testing.T) {
		err := NewStringValidator("Ad").HasLower().Validate()
		if err != nil {
			t.Error("func threw error when it wasnt supposed to")
		}

		err = NewStringValidator("ASD").HasLower().Validate()
		if !errors.Is(err, gverrors.ErrNoLower) {
			t.Errorf("excpected %v but got %v", gverrors.ErrNoLower, err)
		}
	})
	t.Run("has number", func(t *testing.T) {
		err := NewStringValidator("A3d").HasNumber().Validate()
		if err != nil {
			t.Error("func threw error when it wasnt supposed to")
		}

		err = NewStringValidator("kshjd").HasNumber().Validate()
		if !errors.Is(err, gverrors.ErrNoNumber) {
			t.Errorf("excpected %v but got %v", gverrors.ErrNoNumber, err)
		}
	})
	t.Run("has symbol", func(t *testing.T) {
		err := NewStringValidator("@4").HasSymbol().Validate()
		if err != nil {
			t.Error("func threw error when it wasnt supposed to")
		}

		err = NewStringValidator("kshjd").HasSymbol().Validate()
		if !errors.Is(err, gverrors.ErrNoSymbol) {
			t.Errorf("excpected %v but got %v", gverrors.ErrNoSymbol, err)
		}
	})
	t.Run("contains any", func(t *testing.T) {
		err := NewStringValidator("1234").ContainsAny("1").Validate()
		if err != nil {
			t.Error("func threw error when it wasnt supposed to")
		}

		err = NewStringValidator("kshjd").ContainsAny("1").Validate()
		if !errors.Is(err, gverrors.ErrNotContains) {
			t.Errorf("excpected %v but got %v", gverrors.ErrNotContains, err)
		}
	})

	t.Run("url", func(t *testing.T) {
		err := NewStringValidator("https://jayac.dev").URL().Validate()
		if err != nil {
			t.Error("func threw error when it wasnt supposed to")
		}

		err = NewStringValidator("badurl").URL().Validate()
		if !errors.Is(err, gverrors.ErrNotURL) {
			t.Errorf("excpected %v but got %v", gverrors.ErrNotURL, err)
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
