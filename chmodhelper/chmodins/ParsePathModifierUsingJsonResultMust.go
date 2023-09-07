package chmodins

import (
	"gitlab.com/auk-go/core/coredata/corejson"
)

func ParseBaseRwxInstructionsUsingJsonResultMust(
	result *corejson.Result,
) *BaseRwxInstructions {
	baseRwxInstructions, err := ParseBaseRwxInstructionsUsingJsonResult(
		result)

	if err != nil {
		panic(err)
	}

	return baseRwxInstructions
}
