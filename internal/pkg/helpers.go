package pkg

func DefaultOrCustomError(defaultErr error, custom ...error) error {
	if len(custom) == 0 {
		return defaultErr
	}

	return custom[0]
}
