package gv

import (
	"testing"
)

type embedded struct {
	field string
}

type Data struct {
	Name     string
	Username string
	Job      string
	Age      int
	Tags     []string
	embedded embedded
}

func BenchmarkHelpers(b *testing.B) {
	var info = Data{
		Name:     "J",
		Username: "jack 123",
		Job:      "work",
		Age:      24,
		Tags:     []string{"hello", "world"},
	}

	b.Run("first", func(b *testing.B) {
		var err error
		for b.Loop() {
			err = First(
				String(info.Name).Required().Min(3).Max(25).Validate(),
				String(info.Username).Required().Min(3).Max(25).NoWhitespace().Validate(),
				String(info.Job).Max(25).Validate(),
				Numeric(info.Age).Required().Max(100).Finite().Positive().Validate(),
				Slice(info.Tags).Required().Max(1).Element(func(e string) error {
					return String(e).Required().Validate()
				}).Validate(),
			)
		}

		_ = err
	})

	schema := Schema(func(info *Data) error {
		return First(
			String(info.Name).Required().Min(3).Max(25).Validate(),
			String(info.Username).Required().Min(3).Max(25).NoWhitespace().Validate(),
			String(info.Job).Max(25).Validate(),
			Numeric(info.Age).Required().Max(100).Finite().Positive().Validate(),
			Slice(info.Tags).Required().Max(1).Element(func(e string) error {
				return String(e).Required().Validate()
			}).Validate(),
			Schema(func(d *embedded) error {
				return String(d.field).Max(5).Validate()
			}).Validate(&info.embedded),
		)
	})

	b.Run("schema", func(b *testing.B) {
		var err error
		for b.Loop() {
			err = schema.Validate(&info)
		}

		_ = err
	})
	b.Run("Last", func(b *testing.B) {
		var err error
		for b.Loop() {
			err = Last(
				String(info.Name).Required().Min(3).Max(25).Validate(),
				String(info.Username).Required().Min(3).Max(25).NoWhitespace().Validate(),
				String(info.Job).Max(25).Validate(),
				Numeric(info.Age).Required().Max(100).Finite().Positive().Validate(),
				Slice(info.Tags).Required().Max(5).Element(func(e string) error {
					return String(e).Required().Validate()
				}).Validate(),
			)
		}

		_ = err
	})
	b.Run("join", func(b *testing.B) {
		var err error
		for b.Loop() {
			err = Join(
				String(info.Name).Required().Min(3).Max(25).Validate(),
				String(info.Username).Required().Min(3).Max(25).NoWhitespace().Validate(),
				String(info.Job).Max(25).Validate(),
				Numeric(info.Age).Required().Max(100).Finite().Positive().Validate(),
				Slice(info.Tags).Required().Max(5).Element(func(e string) error {
					return String(e).Required().Validate()
				}).Validate(),
			)
		}

		_ = err
	})
}
