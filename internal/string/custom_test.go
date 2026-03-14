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

	t.Run("ipv4", func(t *testing.T) {
		err := NewStringValidator("127.0.0.255").Ipv4().Validate()
		if err != nil {
			t.Error("func threw error when it wasnt supposed to")
		}

		err = NewStringValidator("dd.d.d.dfdd").Ipv4().Validate()
		if !errors.Is(err, gverrors.ErrNotIpv4) {
			t.Errorf("excpected %v but got %v", gverrors.ErrEmail, err)

		}

	})
}

func BenchmarkEmail(b *testing.B) {
	var err error
	for b.Loop() {
		err = NewStringValidator("127.0.0.").Ipv4().Validate()
	}

	_ = err
}
