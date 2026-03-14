package string

import (
	"errors"
	"testing"
)

func TestStringStandard(t *testing.T) {
	t.Run("custom type", func(t *testing.T) {
		type cx string
		type ct cx

		var str ct = "hello"
		err := NewStringValidator(str).Min(3).Validate()

		if err != nil {
			t.Error(err)
		}
	})
	
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

	t.Run("cant not contain", func(t *testing.T) {
		err := NewStringValidator("2").CantContain("1").Validate()
		if err != nil {
			t.Error("func returned a err while it wasnt supposed to")
		}

		err = NewStringValidator("1").CantContain("1").Validate()
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

		var str string = "hello"
		customErr := errors.New("custom")
		err := NewStringValidator(str).Custom(func(val string) error {
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

func BenchmarkStandard(b *testing.B) {
	var err error
	for b.Loop() {
		err = NewStringValidator("hello ").Required().Min(1).Max(100).MustContain("h").CantContain("brrr").NoWhitespace().Validate()
	}
	_ = err

}
