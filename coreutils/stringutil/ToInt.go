package stringutil

import (
	"strconv"
)

func ToInt(
	s string,
	defVal int,
) int {
	toInt, err := strconv.Atoi(s)

	if err != nil {
		return defVal
	}

	return toInt
}
