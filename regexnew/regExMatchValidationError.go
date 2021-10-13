package regexnew

import (
	"fmt"
	"regexp"
)

func regExMatchValidationError(
	regex string,
	comparing string,
	err error,
	regEx *regexp.Regexp,
) error {
	if err != nil {
		return fmt.Errorf(
			"[%q], regex pattern compile failed / invalid cannot match with [%q]",
			err.Error(),
			comparing)
	}

	if regEx == nil {
		return fmt.Errorf(
			"given regex pattern [%q] invalid cannot match with [%q]",
			regex,
			comparing)
	}

	return fmt.Errorf(
		"given regex pattern [%q] doesn't match with [%q]",
		regex,
		comparing)
}
