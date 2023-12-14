package trydo

func WrapPanic(voidFunc func()) Exception {
	var exception Exception

	Block{
		Try: func() {
			voidFunc()
		},
		Catch: func(e Exception) {
			exception = e
		},
		Finally: nil,
	}.Do()

	return exception
}

func ErrorFuncWrapPanic(errFunc func() error) WrappedErr {
	var exception Exception
	var err error
	var hasThrown bool

	Block{
		Try: func() {
			err = errFunc()
		},
		Catch: func(e Exception) {
			exception = e
			hasThrown = e != nil
		},
		Finally: nil,
	}.Do()

	return WrappedErr{
		Error:     err,
		Exception: exception,
		HasThrown: hasThrown,
		HasError:  err != nil,
	}
}
