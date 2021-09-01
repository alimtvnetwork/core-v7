package corejson

import "gitlab.com/evatix-go/core/constants"

type ResultModel struct {
	Bytes    []byte
	Error    string
	HasError bool
}

func NewModel(result *Result) *ResultModel {
	if result == nil {
		return &ResultModel{
			Bytes: []byte{},
			Error: constants.EmptyString,
		}
	}

	model := &ResultModel{}

	return transpileResultToModel(result, model)
}

func NewFromModel(model *ResultModel) *Result {
	if model == nil {
		return &Result{
			Bytes: nil,
			Error: nil,
		}
	}

	result := &Result{}

	return transpileModelToResult(model, result)
}
