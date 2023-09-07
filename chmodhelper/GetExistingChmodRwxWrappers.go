package chmodhelper

import "gitlab.com/auk-go/core/errcore"

func GetExistingChmodRwxWrappers(
	isContinueOnError bool,
	locations ...string,
) (filePathToRwxWrapper map[string]*RwxWrapper, err error) {
	results := make(
		map[string]*RwxWrapper,
		len(locations))

	if len(locations) == 0 {
		return results, nil
	}

	if isContinueOnError {
		var sliceErr []string

		for _, location := range locations {
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

		return results, errcore.SliceToError(sliceErr)
	}

	// immediate exit
	for _, location := range locations {
		wrapperPtr, err2 := GetExistingChmodRwxWrapperPtr(
			location)

		if err2 != nil {
			return results, err2
		} else {
			results[location] = wrapperPtr
		}
	}

	return results, nil
}
