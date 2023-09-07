package corecmp

import "gitlab.com/auk-go/core/corecomparator"

func VersionSliceByte(leftVersions, rightVersions []byte) corecomparator.Compare {
	if leftVersions == nil && rightVersions == nil {
		return corecomparator.Equal
	}

	if leftVersions == nil || rightVersions == nil {
		return corecomparator.NotEqual
	}

	leftLen := len(leftVersions)
	rightLen := len(rightVersions)
	minLength := corecomparator.MinLength(
		leftLen,
		rightLen)

	for i := 0; i < minLength; i++ {
		cLeft := leftVersions[i]
		cRight := rightVersions[i]

		if cLeft == cRight {
			continue
		} else if cLeft < cRight {
			return corecomparator.LeftLess
		} else if cLeft > cRight {
			return corecomparator.LeftGreater
		}
	}

	if leftLen == rightLen {
		return corecomparator.Equal
	} else if leftLen < rightLen {
		return corecomparator.LeftLess
	} else if leftLen > rightLen {
		return corecomparator.LeftGreater
	}

	return corecomparator.NotEqual
}
