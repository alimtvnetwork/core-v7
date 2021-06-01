package chmodhelper

func GetExistingChmodWrapperMustPtr(filePath string) *Wrapper {
	wrapperPtr, err := GetExistingChmodWrapperPtr(filePath)

	if err != nil {
		panic(err)
	}

	return wrapperPtr
}
