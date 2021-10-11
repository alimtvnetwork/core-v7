package errcore

func ToStringPtr(err error) *string {
	if err == nil {
		emptyString := ""

		return &emptyString
	}

	errStr := err.Error()

	return &errStr
}
