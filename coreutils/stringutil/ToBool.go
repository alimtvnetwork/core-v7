package stringutil

import "strconv"

func ToBool(
	s string,
) bool {
	if s == "" {
		return false
	}

	switch s {
	case "yes", "Yes", "YES", "y", "1":
		return true
	case "no", "NO", "No", "0", "n":
		return false
	}

	isBool, err := strconv.ParseBool(s)

	if err != nil {
		return false
	}

	return isBool
}
