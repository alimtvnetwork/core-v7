package errcore

func HandleErrorGetter(errGetter errorGetter) {
	if errGetter == nil {
		return
	}

	err := errGetter.Error()

	if err == nil {
		return
	}

	panic(err.Error())
}
