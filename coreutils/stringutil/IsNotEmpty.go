package stringutil

import "gitlab.com/auk-go/core/constants"

func IsNotEmpty(str string) bool {
	return str != constants.EmptyString
}
