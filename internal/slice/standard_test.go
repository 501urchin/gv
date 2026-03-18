package slice

import (
	"errors"
	"fmt"
	"testing"
)

func TestSliceStandard(t *testing.T) {
	t.Run("custom slice type", func(t *testing.T) {
		type cx []int
		type ct cx

		var v ct = []int{1, 2, 3}
		err := NewSliceValidator(v).Min(2).Validate()

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("not nil", func(t *testing.T) {
		err := NewSliceValidator([]int{1}).Required().Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		var s []int
		err = NewSliceValidator(s).Required().Validate()
		if err == nil {
			t.Error("func didnt return a err while it was supposed to")
		}
	})

	t.Run("min", func(t *testing.T) {
		err := NewSliceValidator([]int{1, 2, 3}).Min(2).Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewSliceValidator([]int{1}).Min(2).Validate()
		if err == nil {
			t.Error("func didnt return a err while it was supposed to")
		}
	})

	t.Run("max", func(t *testing.T) {
		err := NewSliceValidator([]int{1, 2}).Max(3).Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewSliceValidator([]int{1, 2, 3, 4}).Max(3).Validate()
		if err == nil {
			t.Error("func didnt return a err while it was supposed to")
		}
	})
	t.Run("per element", func(t *testing.T) {
		err := NewSliceValidator([]int{1, 2, 4}).Element(func(e int) error {
			if e == 4 {
				return fmt.Errorf("custom err")
			}

			return nil
		})
		if err == nil {
			t.Error("func failed to return custom err")
		}
	})

	t.Run("default", func(t *testing.T) {
		var v []int
		err := NewSliceValidator(v).Default([]int{1, 2, 3}).Min(2).Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewSliceValidator([]int{1}).Default([]int{1, 2, 3}).Min(2).Validate()
		if err == nil {
			t.Error("func didnt keep original value while it was supposed to")
		}
	})

	t.Run("when", func(t *testing.T) {
		customErr := errors.New("custom")

		err := NewSliceValidator([]int{1, 2, 3}).When(false, func(val []int) error {
			return customErr
		}).Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewSliceValidator([]int{1, 2, 3}).When(true, func(val []int) error {
			if len(val) == 3 {
				return customErr
			}

			return nil
		}).Validate()
		if !errors.Is(err, customErr) {
			t.Error("func didnt return custom error while it was supposed to")
		}
	})
}
