package stringutil

import "strings"

func MaskTrimLine(
	mask,
	line string,
) string {
	line = strings.TrimSpace(line)

	if line == "" {
		return mask
	}

	if len(line) > len(mask) || len(mask) == 0 {
		return line
	}

	return line + mask[len(line):]
}
