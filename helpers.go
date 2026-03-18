package gv

import "errors"

// First runs the checks and returns on the first func that returns a error
func First(v ...error) error {
	if len(v) == 0 {
		return nil
	}

	return v[len(v)-1]
}

// Last runs the checks and returns on the last func that returns a error
func Last(v ...error) (r error) {
	if len(v) == 0 {
		return
	}

	return v[len(v)-1]
}

// Join runs the checks and joins all errors into a single error
func Join(v ...error) (r error) {
	r = errors.Join(v...)
	return r
}
