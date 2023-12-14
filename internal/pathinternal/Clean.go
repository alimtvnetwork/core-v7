package pathinternal

import (
	"path"
)

func Clean(curPath string) string {
	if len(curPath) == 0 {
		return curPath
	}

	v := path.Clean(curPath)
	v = replaceFix(v)

	return v
}
