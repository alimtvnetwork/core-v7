package chmodhelper

func GetExistingChmodRwxWrapperMustPtr(
	location string,
) *RwxWrapper {
	wrapperPtr, err := GetExistingChmodRwxWrapperPtr(location)

	if err != nil {
		panic(err)
	}

	return wrapperPtr
}
