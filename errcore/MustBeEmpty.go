package errcore

func MustBeEmpty(err error) {
	if err == nil {
		return
	}

	panic(err.Error())
}
