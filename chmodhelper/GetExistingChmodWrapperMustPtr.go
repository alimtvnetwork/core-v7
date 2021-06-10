package chmodhelper

func GetExistingChmodWrapperMustPtr(
	filePath string,
) *RwxWrapper {
	wrapperPtr, err := GetExistingChmodWrapperPtr(filePath)

	if err != nil {
		panic(err)
	}

	return wrapperPtr
}
