package strutilinternal

func MaskLine(
	mask,
	line string,
) string {
	if line == "" {
		return mask
	}

	if len(line) > len(mask) || len(mask) == 0 {
		return line
	}

	return line + mask[len(line):]
}
