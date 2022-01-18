package errcore

func HandleErrMessage(errMsg string) {
	if errMsg == "" {
		return
	}

	panic(errMsg)
}
