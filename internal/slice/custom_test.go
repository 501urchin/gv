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
func TestAny(t *testing.T) {
	t.Run("valid case", func(t *testing.T) {
		err := NewSliceValidator([]int{1, 2, 3}).Any(func(i int) bool {
			return i%2 == 0
		}).Validate()
		if err != nil {
			t.Errorf("func returned a err when i wasnt supposed to: %v", err)
		}
	})
	t.Run("invalid case", func(t *testing.T) {
		err := NewSliceValidator([]int{1, 1, 1}).Any(func(i int) bool {
			return i%2 == 0
		}).Validate()
		if !errors.Is(err, gverrors.ErrNotSatisfied) {
			t.Errorf("expected %v but got %v", gverrors.ErrNotSatisfied, err)
		}
	})
}
func TestAll(t *testing.T) {
	t.Run("valid case", func(t *testing.T) {
		err := NewSliceValidator([]int{2, 2, 2}).All(func(i int) bool {
			return i%2 == 0
		}).Validate()
		if err != nil {
			t.Errorf("func returned a err when i wasnt supposed to: %v", err)
		}
	})
	t.Run("invalid case", func(t *testing.T) {
		err := NewSliceValidator([]int{2, 2, 1}).All(func(i int) bool {
			return i%2 == 0
		}).Validate()
		if !errors.Is(err, gverrors.ErrNotSatisfied) {
			t.Errorf("expected %v but got %v", gverrors.ErrNotSatisfied, err)
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
func BenchmarkAny(b *testing.B) {
	var err error
	dt := []int{2, 2, 1}
	for b.Loop() {
		err = NewSliceValidator(dt).All(func(i int) bool {
			return i%2 == 0
		}).Validate()
	}
	_ = err
}
