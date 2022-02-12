package errcore

func HandleFullStringsWithTracesGetter(errGetter fullStringWithTracesGetter) {
	if errGetter == nil {
		return
	}

	err := errGetter.FullStringWithTraces()

	if err == nil {
		return
	}

	panic(err.Error())
}
