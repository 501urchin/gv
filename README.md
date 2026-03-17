# gv

low allocation go validation library to make your life a little bit easier.

```bash
go get github.com/501urchin/gv
```

## Example

```go
type Data struct {
	Name     string
	Username string
	Job      string
	Age      int
}

var info = Data{
	Name:     "Jack",
	Username: "501urchin",
	Job:      "Bjj coach",
	Age:      24,
}

dataSchema := gv.Schema(func(info *Data) error {
	return gv.First(
		gv.String(info.Name).Required().Min(3).Max(25).Validate(),
		gv.String(info.Username).Required().Min(3).Max(25).NoWhitespace().Validate(),
		gv.String(info.Job).Max(25).Validate(),
		gv.Numeric(info.Age).Required().Max(100).Finite().Positive().Validate(),
	)
})

err := dataSchema.Validate(&info)
if err != nil {
	// do something with error
}
```

## Supported types

the base types gv supports are `string`, `numeric`, and `slice`. you can call them like this:

```go
err := gv.String("hi").Validate() // uses generics with type ~string
if err != nil {
	return err
}

err = gv.Numeric(223).Validate() // supports most numeric go types except complex
if err != nil {
	return err
}

err = gv.Slice([]int{1, 2, 3}).Validate() // []T so you can pass any slice or arr[:]
if err != nil {
	return err
}
```

## Error helpers

if you dont want to check all errors manually gv provides some helpers:

```go
gv.First() // returns the first error that failed
gv.Last()  // returns the last error that failed
gv.Join()  // collects all errors into a single error using errors.Join
```

example:

```go
err := gv.First(
	gv.String("hello").Max(2).Validate(), // will return this error
	gv.Numeric(223).Validate(),
	gv.Slice([]int{1, 2, 3}).Validate(),
)

if err != nil {
	return err
}
```

## Schema

since writing validation logic can become repetitive i took inspiration from zod and wrote a schema helper.

```go
// Schema takes in func(info *T) error.
// T represents the data you pass into Validate, so you can reference it
// inside the validation function and write custom logic per field.
dataSchema := gv.Schema(func(info *Data) error {
	return gv.First(
		gv.String(info.Name).Required().Min(3).Max(25).Validate(),
		gv.String(info.Username).Required().Min(3).Max(25).NoWhitespace().Validate(),
		gv.String(info.Job).Max(25).Validate(),
		gv.Numeric(info.Age).Required().Max(100).Finite().Positive().Validate(),

		// if your data has an embedded struct you can embed another schema here
	)
})

err := dataSchema.Validate(dataPointer)
if err != nil {
	return err
}
```

## Benchmark

| Benchmark                  | Iterations | Time (ns/op) | Memory (B/op) | Allocs (allocs/op) |
| -------------------------- | ---------: | -----------: | ------------: | -----------------: |
| 501urchin/gv-10            | 83,503,381 |        13.95 |             0 |                  0 |
| go-playground/validator-10 |  3,241,236 |        370.1 |           688 |                 11 |
| Oudwins/zog-10             |    824,401 |         1448 |           682 |                 18 |
