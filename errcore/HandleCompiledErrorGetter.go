package errcore

func HandleCompiledErrorGetter(errGetter compiledErrorGetter) {
	if errGetter == nil {
		return
	}

	err := errGetter.CompiledError()

	if err == nil {
		return
	}

	panic(err.Error())
}
