package slice

import gverrors "github.com/501urchin/gv/internal/errors"

func (s *SliceValidator[T]) Unique() *SliceValidator[T] {
	if len(s.val) <= 1 {
		return s
	}

	vm := make(map[any]int)

	for _, v := range s.val {
		vm[v]++

		if vm[v] > 1 {
			s.err = gverrors.ErrNotUnique
			return s
		}
	}

	return s
}

// UniqueBy can be used for cases where you have a []struct and want only a specific field to be unqiue
func (s *SliceValidator[T]) UniqueBy(fn func(v T) any) *SliceValidator[T] {
	if len(s.val) <= 1 {
		return s
	}

	vm := make(map[any]int)

	for _, v := range s.val {
		vl := fn(v)
		vm[vl]++

		if vm[vl] > 1 {
			s.err = gverrors.ErrNotUnique
			return s
		}
	}
	return s
}
func (s *SliceValidator[T]) Custom(fn func(v []T) error) *SliceValidator[T] {
	err := fn(s.val)
	if err != nil {
		s.err = err
	}

	return s
}

func (s *SliceValidator[T]) Any(fn func(T) bool) *SliceValidator[T] {
	for _, v := range s.val {
		if fn(v) {
			return s
		}
	}

	s.err = gverrors.ErrNotSatisfied
	return s
}

func (s *SliceValidator[T]) All(fn func(T) bool) *SliceValidator[T] {
	for _, v := range s.val {
		if !fn(v) {
			s.err = gverrors.ErrNotSatisfied
			return s
		}
	}

	return s
}
