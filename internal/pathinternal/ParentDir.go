package pathinternal

import "path/filepath"

func ParentDir(curPath string) string {
	if len(curPath) == 0 {
		return curPath
	}

	v := filepath.Dir(curPath)

	return Clean(v)
}
