package conditional

import (
	"strconv"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
)

func executeErrorFunctions(functions []func() error) error {
	if len(functions) == 0 {
		return nil
	}

	var sliceErr []string
	for index, currentFunction := range functions {
		if currentFunction == nil {
			continue
		}

		err := currentFunction()

		if err != nil {
			sliceErr = append(sliceErr, err.Error()+"- index of - "+strconv.Itoa(index))
		}
	}

	return errcore.SliceToError(sliceErr)
}

func executeAnyFunctions(
	functions []func() (
		result interface{},
		isTake,
		isBreak bool,
	),
) []interface{} {
	if len(functions) == 0 {
		return nil
	}

	results := make(
		[]interface{},
		constants.Zero,
		len(functions))
	for _, curFunc := range functions {
		if curFunc == nil {
			continue
		}

		result, isTake, isBreak := curFunc()

		if isTake {
			results = append(results, result)
		}

		if isBreak {
			return results
		}
	}

	return results
}
