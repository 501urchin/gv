package gv

import "errors"

// First runs the checks and returns on the first func that returns a error
func First(v ...error) error {
	for _, fn := range v {
		if err := fn; err != nil {
			return err
		}
	}

	return nil
}

// Last runs the checks and returns on the last func that returns a error
func Last(v ...error) (r error) {
	for _, fn := range v {
		if err := fn; err != nil {
			r = err
		}
	}

	return r
}

// Join runs the checks and joins all errors into a single error
func Join(v ...error) (r error) {
	for _, fn := range v {
		if err := fn; err != nil {
			r = errors.Join(r, err)
		}
	}

	return r
}
