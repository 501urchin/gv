package slice

import (
	gverrors "github.com/501urchin/gv/internal/errors"
	"github.com/501urchin/gv/internal/pkg"
)

func (s *SliceValidator[T]) Unique(customErr ...error) *SliceValidator[T] {
	if len(s.val) <= 1 {
		return s
	}

	vm := make(map[any]int)

	for _, v := range s.val {
		vm[v]++

		if vm[v] > 1 {
			s.err = pkg.DefaultOrCustomError(gverrors.ErrNotUnique, customErr...)
			return s
		}
	}

	return s
}

// UniqueBy can be used for cases where you have a []struct and want only a specific field to be unqiue
func (s *SliceValidator[T]) UniqueBy(fn func(v T) any, customErr ...error) *SliceValidator[T] {
	if len(s.val) <= 1 {
		return s
	}

	vm := make(map[any]int)

	for _, v := range s.val {
		vl := fn(v)
		vm[vl]++

		if vm[vl] > 1 {
			s.err = pkg.DefaultOrCustomError(gverrors.ErrNotUnique, customErr...)
			return s
		}
	}
	return s
}
func (s *SliceValidator[T]) Custom(fn func(v []T) error, customErr ...error) *SliceValidator[T] {
	err := fn(s.val)
	if err != nil {
		s.err = err
	}

	return s
}

func (s *SliceValidator[T]) Any(fn func(T) bool, customErr ...error) *SliceValidator[T] {
	for _, v := range s.val {
		if fn(v) {
			return s
		}
	}

	s.err = pkg.DefaultOrCustomError(gverrors.ErrNotSatisfied, customErr...)
	return s
}

func (s *SliceValidator[T]) All(fn func(T) bool, customErr ...error) *SliceValidator[T] {
	for _, v := range s.val {
		if !fn(v) {
			s.err = pkg.DefaultOrCustomError(gverrors.ErrNotSatisfied, customErr...)
			return s
		}
	}

	return s
}
