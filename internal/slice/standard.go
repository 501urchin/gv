package slice

// Element loops over the slice and runs the custom validation func fn. fn needs to return a error
func (s *SliceValidator[T]) Element(fn func(e T) error) *SliceValidator[T] {
	for i := range s.val {
		if err := fn(s.val[i]); err != nil {
			s.err = err
		}
	}

	return s
}
