package slice

import (
	"errors"
	"testing"

	gverrors "github.com/501urchin/gv/internal/errors"
)

func TestUnique(t *testing.T) {
	t.Run("valid case", func(t *testing.T) {
		err := NewSliceValidator([]int{1, 2, 3}).Unique().Validate()
		if err != nil {
			t.Errorf("func returned a err when i wasnt supposed to: %v", err)
		}
	})
	t.Run("invalid case", func(t *testing.T) {
		err := NewSliceValidator([]int{1, 1, 1.}).Unique().Validate()
		if !errors.Is(err, gverrors.ErrNotUnique) {
			t.Errorf("expected %v but got %v", gverrors.ErrNotUnique, err)
		}
	})
}

func TestUniqueBy(t *testing.T) {
	type Data struct {
		val string
		age int
	}

	t.Run("valid case", func(t *testing.T) {
		err := NewSliceValidator([]Data{{val: "hi", age: 12}, {val: "hello", age: 13}}).UniqueBy(func(v Data) any {
			return v.age
		}).Validate()
		if err != nil {
			t.Errorf("func returned a err when i wasnt supposed to: %v", err)
		}
	})
	t.Run("invalid case", func(t *testing.T) {
		err := NewSliceValidator([]Data{{val: "hi", age: 12}, {val: "hello", age: 12}}).UniqueBy(func(v Data) any {
			return v.age
		}).Validate()
		if !errors.Is(err, gverrors.ErrNotUnique) {
			t.Errorf("expected %v but got %v", gverrors.ErrNotUnique, err)
		}
	})
}

func BenchmarkUnique(b *testing.B) {
	var err error
	dt := []int{1, 2, 3}
	for b.Loop() {
		err = NewSliceValidator(dt).Unique().Validate()
	}
	_ = err
}
