package stringutil

import "gitlab.com/auk-go/core/constants"

func ToInt32Def(
	s string,
) int32 {
	return ToInt32(s, constants.Zero)
}
