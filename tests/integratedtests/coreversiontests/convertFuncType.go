package coreversiontests

func convertFuncType(i interface{}) (resultFunc isBoolCheckerFunc) {
	if f, ok := i.(func(x interface{}) bool); ok {
		return f
	}

	return nil
}
