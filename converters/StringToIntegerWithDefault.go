package converters

import (
	"strconv"

	"gitlab.com/auk-go/core/constants"
)

func StringToIntegerWithDefault(
	input string,
	defaultInt int,
) (value int, isSuccess bool) {
	if input == constants.EmptyString {
		return defaultInt, false
	}

	convertedVal, err := strconv.Atoi(input)

	if err != nil {
		return defaultInt, false
	}

	return convertedVal, true
}
