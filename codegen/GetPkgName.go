package codegen

import (
	"path/filepath"
	"strings"
)

// GetPkgName
//
// fullImportLine - "gitlab.com/auk-go/core/internal/reflectinternal"
//
// # returns
//   - prefix : gitlab.com/auk-go/core/internal/
//   - pkgName : reflectinternal
func GetPkgName(fullImportLine string) (prefix, pkgName string) {
	trim := strings.TrimSpace(fullImportLine)

	if len(trim) <= 2 {
		return trim, trim
	}

	unWrap := fullImportLine[2 : len(fullImportLine)-1]

	return filepath.Split(unWrap)
}
