package converters

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func MapStringAnyToMapStringString(
	isSkipEmpty bool,
	additionalMapItems map[string]interface{},
) map[string]string {
	if len(additionalMapItems) == 0 {
		return map[string]string{}
	}

	newMap := make(map[string]string, len(additionalMapItems))

	for key, valInf := range additionalMapItems {
		val := fmt.Sprintf(
			constants.SprintValueFormat,
			valInf)

		if isSkipEmpty && val == "" {
			continue
		}

		newMap[key] = val
	}

	return newMap
}

func CloneMapStringStringPlusAppendMapStringAny(
	isSkipEmpty bool,
	mainMap map[string]interface{},
	additionalMapItems map[string]interface{},
) map[string]string {
	if len(mainMap) == 0 && len(additionalMapItems) == 0 {
		return map[string]string{}
	}

	newMap := make(
		map[string]string,
		len(mainMap)+
			len(additionalMapItems)+
			constants.Capacity3)

	for key, valInf := range additionalMapItems {
		val := fmt.Sprintf(
			constants.SprintValueFormat,
			valInf)

		if isSkipEmpty && val == "" {
			continue
		}

		newMap[key] = val
	}

	return newMap
}
