package stringutil

func MaskLines(
	mask string,
	lines ...string,
) []string {
	length := len(lines)
	slice := make([]string, length)

	if length == 0 {
		return slice
	}

	for i, line := range lines {
		if len(line) > len(mask) || len(mask) == 0 {
			slice[i] = line

			continue
		}

		slice[i] = line + mask[len(line):]
	}

	return slice
}
