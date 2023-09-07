package chmodhelper

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/errcore"
)

// GetFilesChmodRwxFullMap returns filePath -> "-rwxrwxrwx"
func GetFilesChmodRwxFullMap(
	requestedPaths []string,
) (filePathToRwxMap *corestr.Hashmap, err error) {
	length := len(requestedPaths)
	hashmap := corestr.New.Hashmap.Cap(length)

	if length == 0 {
		return hashmap, nil
	}

	var sliceErr []string

	for _, filePath := range requestedPaths {
		fileMode, err2 := GetExistingChmod(filePath)

		if err2 != nil {
			sliceErr = append(sliceErr, err2.Error())

			continue
		}

		hashmap.AddOrUpdate(filePath, fileMode.String())
	}

	return hashmap, errcore.SliceErrorDefault(&sliceErr)
}
