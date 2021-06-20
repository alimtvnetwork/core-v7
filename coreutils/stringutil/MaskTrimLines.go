package stringutil

import "strings"

func MaskTrimLines(
	mask string,
	lines ...string,
) []string {
	length := len(lines)
	slice := make([]string, length)

	if length == 0 {
		return slice
	}

	for i, line := range lines {
		line2 := strings.TrimSpace(line)

		if len(line2) > len(mask) || len(mask) == 0 {
			slice[i] = line2

			continue
		}

		slice[i] = line2 + mask[len(line2):]
	}

	return slice
}
