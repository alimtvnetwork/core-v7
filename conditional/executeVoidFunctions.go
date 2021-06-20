package conditional

func executeVoidFunctions(functions []func()) {
	if len(functions) == 0 {
		return
	}

	for _, currentFunction := range functions {
		if currentFunction == nil {
			continue
		}

		currentFunction()
	}
}
