package coreappend

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func MapStringStringAppendMapStringToAnyItems(
	isSkipEmpty bool,
	mainMap map[string]string,
	appendMapItems map[string]interface{},
) map[string]string {
	if len(appendMapItems) == 0 {
		return mainMap
	}

	for key, valInf := range appendMapItems {
		val := fmt.Sprintf(
			constants.SprintValueFormat,
			valInf)

		if isSkipEmpty && val == "" {
			continue
		}

		mainMap[key] = val
	}

	return mainMap
}
