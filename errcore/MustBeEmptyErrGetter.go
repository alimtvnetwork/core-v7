package errcore

import "gitlab.com/evatix-go/core/coreinterface"

func MustBeEmptyErrGetter(errorGetter coreinterface.ErrorGetter) {
	if errorGetter == nil {
		return
	}

	err := errorGetter.Error()
	if err == nil {
		return
	}

	panic(err.Error())
}
