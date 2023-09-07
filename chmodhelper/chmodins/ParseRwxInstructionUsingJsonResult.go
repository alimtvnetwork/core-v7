package chmodins

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/errcore"
)

func ParseRwxInstructionUsingJsonResult(
	result *corejson.Result,
) (*RwxInstruction, error) {
	if result == nil {
		return nil,
			errcore.BytesAreNilOrEmptyType.Error(
				"ParseRwxInstructionUsingJsonResult", nil)
	}

	if result.IsEmptyJsonBytes() || result.HasError() {
		return nil, result.MeaningfulError()
	}

	var rwxInstruction RwxInstruction
	err := result.Unmarshal(&rwxInstruction)

	if err != nil {
		return nil, errcore.MeaningfulError(
			errcore.FailedToParseType,
			"ParseRwxInstructionUsingJsonResult",
			err)
	}

	return &rwxInstruction, err
}
