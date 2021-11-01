package conditional

import "gitlab.com/evatix-go/core/constants"

func StringDefault(
	isTrue bool,
	trueValue string,
) string {
	if isTrue {
		return trueValue
	}

	return constants.EmptyString
}
