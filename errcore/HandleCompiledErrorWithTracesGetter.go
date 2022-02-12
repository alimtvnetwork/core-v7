package errcore

func HandleCompiledErrorWithTracesGetter(errGetter compiledErrorWithTracesGetter) {
	if errGetter == nil {
		return
	}

	err := errGetter.CompiledErrorWithStackTraces()

	if err == nil {
		return
	}

	panic(err.Error())
}
