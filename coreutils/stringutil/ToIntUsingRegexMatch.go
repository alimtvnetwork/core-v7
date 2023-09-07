package stringutil

import (
	"regexp"
	"strconv"

	"gitlab.com/auk-go/core/constants"
)

func ToIntUsingRegexMatch(
	regEx *regexp.Regexp,
	s string,
) int {
	if regEx == nil || !regEx.MatchString(s) {
		return constants.Zero
	}

	toInt, err := strconv.Atoi(s)

	if err != nil {
		return constants.Zero
	}

	return toInt
}
