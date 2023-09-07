package chmodhelper

import (
	"os"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/fsinternal"
)

// GetFilteredExistsFilesInfosOnly
//
// Warning: File related errors will be swallowed
func GetFilteredExistsFilesInfosOnly(
	locations ...string,
) map[string]os.FileInfo {
	if len(locations) == 0 {
		return map[string]os.FileInfo{}
	}

	results := make(
		map[string]os.FileInfo,
		len(locations)+constants.Capacity5)

	for _, location := range locations {
		info, isExist, _ :=
			fsinternal.GetPathExistStat(location)

		if isExist && info != nil {
			results[location] = info
		}
	}

	return results
}
