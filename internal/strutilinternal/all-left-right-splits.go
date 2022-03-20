package strutilinternal

import "strings"

func SplitLeftRight(splitter, line string) (l, r string) {
	splits := strings.Split(
		line,
		splitter)

	if len(splits) >= 2 {
		return splits[0], splits[len(splits)-1]
	}

	return splits[0], ""
}

func SplitLeftRightTrim(splitter, line string) (l, r string) {
	l, r = SplitLeftRight(splitter, line)

	return strings.TrimSpace(l), strings.TrimSpace(r)
}
