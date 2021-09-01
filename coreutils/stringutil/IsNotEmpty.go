package stringutil

import "gitlab.com/evatix-go/core/constants"

func IsNotEmpty(str string) bool {
	return str != constants.EmptyString
}
