package gv

import (
	"testing"
)

func BenchmarkHelpers(b *testing.B) {
	type Data struct {
		Name     string
		Username string
		Job      string
		Age      int
		Tags     []string
	}
	var info = Data{
		Name:     "J",
		Username: "jack 123",
		Job:      "sodifbovhbsodhfbvohsbdfhuvohuvsboudfhbovhbfouhbvhsbovhjbsodhbvodhsobfhvodhsfvbodfhboh",
		Age:      24,
		Tags:     []string{"hello", "world"},
	}

	b.Run("first", func(b *testing.B) {
		var err error
		for b.Loop() {
			err = First(
				String(info.Name).Required().Min(3).Max(25),
				String(info.Username).Required().Min(3).Max(25).NoWhitespace(),
				String(info.Job).Max(25),
				Numeric(info.Age).Required().Max(100).Finite().Positive(),
				Slice(info.Tags).NotEmpty().Max(5).Element(func(e string) error {
					return String(e).Required().Validate()
				}),
			)
		}

		_ = err
	})
	b.Run("Last", func(b *testing.B) {
		var err error
		for b.Loop() {
			err = Last(
				String(info.Name).Required().Min(3).Max(25),
				String(info.Username).Required().Min(3).Max(25).NoWhitespace(),
				String(info.Job).Max(25),
				Numeric(info.Age).Required().Max(100).Finite().Positive(),
				Slice(info.Tags).NotEmpty().Max(5).Element(func(e string) error {
					return String(e).Required().Validate()
				}),
			)
		}

		_ = err
	})
	b.Run("join", func(b *testing.B) {
		var err error
		for b.Loop() {
			err = Join(
				String(info.Name).Required().Min(3).Max(25),
				String(info.Username).Required().Min(3).Max(25).NoWhitespace(),
				String(info.Job).Max(25),
				Numeric(info.Age).Required().Max(100).Finite().Positive(),
				Slice(info.Tags).NotEmpty().Max(5).Element(func(e string) error {
					return String(e).Required().Validate()
				}),
			)
		}

		_ = err
	})
}
