package chmodinstruction

import (
	"encoding/json"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/msgtype"
)

func ParsePathModifierUsingJsonResultMust(
	result *corejson.Result,
) *PathModifier {
	if result == nil {
		panic(msgtype.JsonResultBytesAreNilOrEmpty.String())
	}

	if result.IsEmptyJsonBytes() || result.HasError() {
		panic(result.MeaningfulError())
	}

	var pathModifier PathModifier

	err := json.Unmarshal(*result.Bytes, &pathModifier)

	msgtype.MeaningFulErrorHandle(
		msgtype.FailedToParse,
		"ParsePathModifierUsingJsonResultMust",
		err)

	return &pathModifier
}
