package chmodhelper

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/msgtype"
)

// GetFilesChmodRwxFullMap returns filePath -> "-rwxrwxrwx"
func GetFilesChmodRwxFullMap(
	requestedPaths []string,
) (filePathToRwxMap *corestr.Hashmap, err error) {
	length := len(requestedPaths)
	hashmap := corestr.NewHashmap(length)

	if length == 0 {
		return hashmap, nil
	}

	var sliceErr []string

	for _, filePath := range requestedPaths {
		fileMode, err := GetExistingChmod(filePath)

		if err != nil {
			sliceErr = append(sliceErr, err.Error())

			continue
		}

		hashmap.AddOrUpdate(filePath, fileMode.String())
	}

	return hashmap, msgtype.SliceErrorDefault(&sliceErr)
}
