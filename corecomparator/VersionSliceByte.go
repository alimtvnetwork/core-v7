package corecomparator

func VersionSliceByte(leftVersions, rightVersions []byte) Compare {
	if leftVersions == nil && rightVersions == nil {
		return Equal
	}

	if leftVersions == nil || rightVersions == nil {
		return NotEqual
	}

	leftLen := len(leftVersions)
	rightLen := len(rightVersions)
	minLength := MinLength(
		leftLen,
		rightLen)

	for i := 0; i < minLength; i++ {
		cLeft := leftVersions[i]
		cRight := rightVersions[i]

		if cLeft == cRight {
			continue
		} else if cLeft < cRight {
			return LeftLess
		} else if cLeft > cRight {
			return LeftGreater
		}
	}

	if leftLen == rightLen {
		return Equal
	} else if leftLen < rightLen {
		return LeftLess
	} else if leftLen > rightLen {
		return LeftGreater
	}

	return NotEqual
}
