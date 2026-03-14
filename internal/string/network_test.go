package string

import (
	"errors"
	"testing"

	gverrors "github.com/501urchin/gv/internal/errors"
)

func TestIpv4(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		err := NewStringValidator("127.0.0.255").Ipv4().Validate()
		if err != nil {
			t.Error("func threw error when it wasnt supposed to")
		}

		err = NewStringValidator("dd.d.d.dfdd").Ipv4().Validate()
		if !errors.Is(err, gverrors.ErrNotIpv4) {
			t.Errorf("excpected %v but got %v", gverrors.ErrNotIpv4, err)

		}

	})
}

func BenchmarkIpv4(b *testing.B) {
	var err error
	for b.Loop() {
		err = NewStringValidator("127.0.0.").Ipv4().Validate()
	}

	_ = err
}
