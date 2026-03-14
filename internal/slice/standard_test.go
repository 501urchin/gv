package slice

import (
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
		err := NewSliceValidator([]int{1}).NotNil().Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		var s []int
		err = NewSliceValidator(s).NotNil().Validate()
		if err == nil {
			t.Error("func didnt return a err while it was supposed to")
		}
	})

	t.Run("not empty", func(t *testing.T) {
		err := NewSliceValidator([]int{1}).NotEmpty().Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewSliceValidator([]int{}).NotEmpty().Validate()
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
		}).Validate()
		if err == nil {
			t.Error("func failed to return custom err")
		}
	})
}
