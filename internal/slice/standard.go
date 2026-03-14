package slice

import gverrors "github.com/501urchin/gv/internal/errors"

// Element loops over the slice and runs the custom validation func fn. fn needs to return a error
func (s *SliceValidator[T]) Element(fn func(e T) error) *SliceValidator[T] {
	if s.err != nil {
		return s
	}

	for i := range s.val {
		if err := fn(s.val[i]); err != nil {
			s.err = err
		}
	}

	return s
}
func (s *SliceValidator[T]) NotNil() *SliceValidator[T] {
	if s.err != nil {
		return s
	}

	if s.val == nil {
		s.err = gverrors.ErrIsNil
	}

	return s
}
func (s *SliceValidator[T]) NotEmpty() *SliceValidator[T] {
	if s.err != nil {
		return s
	}

	if len(s.val) == 0 {
		s.err = gverrors.ErrEmpty
	}

	return s
}
func (s *SliceValidator[T]) Min(v int) *SliceValidator[T] {
	if s.err != nil {
		return s
	}

	if len(s.val) < v {
		s.err = gverrors.ErrMin
	}

	return s
}

func (s *SliceValidator[T]) Max(v int) *SliceValidator[T] {
	if s.err != nil {
		return s
	}

	if len(s.val) > v {
		s.err = gverrors.ErrMax
	}

	return s
}

