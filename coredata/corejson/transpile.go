package corejson

import (
	"errors"

	"gitlab.com/evatix-go/core/constants"
)

func transpileModelToResult(model *ResultModel, result *Result) *Result {
	var err error
	hasError := model.HasError

	if hasError {
		err = errors.New(
			model.Error)
	}

	result.Bytes = model.Bytes
	result.Error = err

	return result
}

func transpileResultToModel(result *Result, model *ResultModel) *ResultModel {
	errString := constants.EmptyString
	hasError := result.Error != nil

	if hasError {
		errString = result.Error.Error()
	}

	model.Bytes = result.Bytes
	model.Error = errString
	model.HasError = hasError

	return model
}
