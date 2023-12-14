package pathinternal

import "path/filepath"

func Relative(basePath, targetPath string) string {
	v, _ := filepath.Rel(basePath, targetPath)

	return Clean(v)
}
