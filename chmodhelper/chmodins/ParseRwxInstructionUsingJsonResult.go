package chmodins

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/errcore"
)

func ParseRwxInstructionUsingJsonResult(
	result *corejson.Result,
) (*RwxInstruction, error) {
	if result == nil {
		return nil,
			errcore.JsonResultBytesAreNilOrEmpty.Error(
				"ParseRwxInstructionUsingJsonResult", nil)
	}

	if result.IsEmptyJsonBytes() || result.HasError() {
		return nil, result.MeaningfulError()
	}

	var rwxInstruction RwxInstruction
	err := result.Unmarshal(&rwxInstruction)

	if err != nil {
		return nil, errcore.MeaningfulError(
			errcore.FailedToParse,
			"ParseRwxInstructionUsingJsonResult",
			err)
	}

	return &rwxInstruction, err
}
