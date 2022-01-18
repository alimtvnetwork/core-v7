package errcore

func HandleErr(err error) {
	if err == nil {
		return
	}

	panic(err.Error())
}
