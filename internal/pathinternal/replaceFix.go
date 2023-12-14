package pathinternal

import "strings"

func replaceFix(curPath string) string {
	if len(curPath) == 0 {
		return curPath
	}

	for _, replacer := range replacers {
		curPath = strings.ReplaceAll(
			curPath,
			replacer.from,
			replacer.to,
		)
	}

	return curPath
}
