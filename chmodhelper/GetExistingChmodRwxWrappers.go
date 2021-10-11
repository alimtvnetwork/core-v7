package chmodhelper

import "gitlab.com/evatix-go/core/errcore"

func GetExistingChmodRwxWrappers(
	isContinueOnError bool,
	filePaths ...string,
) (filePathToRwxWrapper *map[string]*RwxWrapper, err error) {
	results := make(
		map[string]*RwxWrapper,
		len(filePaths))

	if len(filePaths) == 0 {
		return &results, nil
	}

	if isContinueOnError {
		var sliceErr []string

		for _, location := range filePaths {
			wrapperPtr, err2 := GetExistingChmodRwxWrapperPtr(
				location)

			if err2 != nil {
				sliceErr = append(
					sliceErr,
					err2.Error())
			} else {
				results[location] = wrapperPtr
			}
		}

		return &results, errcore.SliceToError(sliceErr)
	}

	// immediate exit
	for _, location := range filePaths {
		wrapperPtr, err2 := GetExistingChmodRwxWrapperPtr(
			location)

		if err2 != nil {
			return &results, err2
		} else {
			results[location] = wrapperPtr
		}
	}

	return &results, nil
}
