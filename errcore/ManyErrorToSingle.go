package errcore

func ManyErrorToSingle(errorItems []error) error {
	if errorItems == nil || len(errorItems) == 0 {
		return nil
	}

	sliceErr := make([]string, 0, len(errorItems))

	for _, err := range errorItems {
		if err == nil {
			continue
		}

		sliceErr = append(sliceErr, err.Error())
	}

	return SliceToError(sliceErr)
}
