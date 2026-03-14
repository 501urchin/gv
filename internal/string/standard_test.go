package string

import (
	"errors"
	"testing"
)

func TestStringStandard(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		err := NewStringValidator("hello").Required().Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewStringValidator("").Required().Validate()
		if err == nil {
			t.Error("func didnt return a err while it was supposed to")
		}
	})

	t.Run("min", func(t *testing.T) {
		err := NewStringValidator("12").Min(2).Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewStringValidator("1").Min(2).Validate()
		if err == nil {
			t.Error("func didnt return a err while it was supposed to")
		}
	})

	t.Run("max", func(t *testing.T) {
		err := NewStringValidator("1").Max(3).Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewStringValidator("1234").Max(3).Validate()
		if err == nil {
			t.Error("func didnt return a err while it was supposed to")
		}
	})

	t.Run("must contain", func(t *testing.T) {
		err := NewStringValidator("1").MustContain("1").Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewStringValidator("234").MustContain("1").Validate()
		if err == nil {
			t.Error("func didnt return a err while it was supposed to")
		}
	})

	t.Run("must not contain", func(t *testing.T) {
		err := NewStringValidator("2").NotContains("1").Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewStringValidator("1").NotContains("1").Validate()
		if err == nil {
			t.Error("func didnt return a err while it was supposed to")
		}
	})

	t.Run("no whitespace", func(t *testing.T) {
		err := NewStringValidator("2").NoWhitespace().Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewStringValidator("s ").NoWhitespace().Validate()
		if err == nil {
			t.Error("func didnt return a err while it was supposed to")
		}
	})

	t.Run("Custom", func(t *testing.T) {
		type ct string

		var str ct = "hello"
		customErr := errors.New("custom")
		err := NewStringValidator(str).Custom(func(val ct) error {
			if val == str {
				return customErr
			}
			return nil
		}).Validate()

		if err == nil {
			t.Error("func failed to return a custom error")
		}

		if !errors.Is(err, customErr) {
			t.Errorf("failed to return correct error: excpected %v but got %v", customErr, err)
		}
	})

}
