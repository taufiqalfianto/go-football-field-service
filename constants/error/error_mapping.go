package error

func ErrMapping(err error) bool {
	allError := make([]error, 0)
	allError = append(append(UserError[:], GeneralError[:]...))

	for _, e := range allError {
		if err == e {
			return true
		}
	}
	return false
}
