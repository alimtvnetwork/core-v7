package stringutil

import "gitlab.com/auk-go/core/constants"

func ToInt8Def(
	s string,
) int8 {
	return ToInt8(s, constants.Zero)
}
