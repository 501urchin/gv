package numeric

import (
	"errors"
	"testing"

	gverrors "github.com/501urchin/gv/internal/errors"
)

func TestStandard(t *testing.T) {
	t.Run("min", func(t *testing.T) {
		err := NewNumericValidator(5).Min(1).Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewNumericValidator(-5).Min(1).Validate()
		if !errors.Is(err, gverrors.ErrMin) {
			t.Error("func didnt return ErrMin while it was supposed to")
		}
	})

	t.Run("max", func(t *testing.T) {
		err := NewNumericValidator(3).Max(5).Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewNumericValidator(10).Max(5).Validate()
		if !errors.Is(err, gverrors.ErrMax) {
			t.Error("func didnt return ErrMax while it was supposed to")
		}
	})

	t.Run("required", func(t *testing.T) {
		err := NewNumericValidator(10).Required().Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewNumericValidator(0).Required().Validate()
		if !errors.Is(err, gverrors.ErrRequired) {
			t.Error("func didnt return ErrRequired while it was supposed to")
		}
	})

	t.Run("negative", func(t *testing.T) {
		err := NewNumericValidator(-5).Negative().Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewNumericValidator(5).Negative().Validate()
		if !errors.Is(err, gverrors.ErrNotNegative) {
			t.Error("func didnt return ErrNotNegative while it was supposed to")
		}
	})

	t.Run("positive", func(t *testing.T) {
		err := NewNumericValidator(5).Positive().Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewNumericValidator(-5).Positive().Validate()
		if !errors.Is(err, gverrors.ErrNotPositive) {
			t.Error("func didnt return ErrNotPositive while it was supposed to")
		}
	})

	t.Run("equal", func(t *testing.T) {
		err := NewNumericValidator(5).Equal(5).Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewNumericValidator(5).Equal(3).Validate()
		if !errors.Is(err, gverrors.ErrNotEqual) {
			t.Error("func didnt return ErrNotEqual while it was supposed to")
		}
	})

	t.Run("not equal", func(t *testing.T) {
		err := NewNumericValidator(5).NotEqual(3).Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewNumericValidator(5).NotEqual(5).Validate()
		if !errors.Is(err, gverrors.ErrEqual) {
			t.Error("func didnt return ErrEqual while it was supposed to")
		}
	})

	t.Run("custom", func(t *testing.T) {
		customErr := errors.New("custom")

		err := NewNumericValidator(5).Custom(func(v int) error {
			return nil
		}).Validate()

		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewNumericValidator(5).Custom(func(v int) error {
			return customErr
		}).Validate()

		if !errors.Is(err, customErr) {
			t.Error("func didnt return custom error while it was supposed to")
		}
	})

	t.Run("default", func(t *testing.T) {
		err := NewNumericValidator(0).Default(5).Min(1).Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewNumericValidator(1).Default(10).Max(0).Validate()
		if !errors.Is(err, gverrors.ErrMax) {
			t.Error("func didnt keep original value while it was supposed to")
		}
	})

	t.Run("when", func(t *testing.T) {
		customErr := errors.New("custom")

		err := NewNumericValidator(5).When(false, func(val int) error {
			return customErr
		}).Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewNumericValidator(5).When(true, func(val int) error {
			if val == 5 {
				return customErr
			}

			return nil
		}).Validate()
		if !errors.Is(err, customErr) {
			t.Error("func didnt return custom error while it was supposed to")
		}
	})
}
