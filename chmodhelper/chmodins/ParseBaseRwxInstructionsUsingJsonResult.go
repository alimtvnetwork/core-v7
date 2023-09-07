package chmodins

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/errcore"
)

func ParseBaseRwxInstructionsUsingJsonResult(
	result *corejson.Result,
) (*BaseRwxInstructions, error) {
	if result == nil {
		return nil,
			errcore.BytesAreNilOrEmptyType.Error(
				"ParseBaseRwxInstructionsUsingJsonResult", nil)
	}

	if result.IsEmptyJsonBytes() || result.HasError() {
		return nil, result.MeaningfulError()
	}

	var baseRwxInstructions BaseRwxInstructions
	err := result.Unmarshal(&baseRwxInstructions)

	if err != nil {
		return nil, errcore.MeaningfulError(
			errcore.FailedToParseType,
			"ParseBaseRwxInstructionsUsingJsonResult",
			err)
	}

	return &baseRwxInstructions, nil
}
