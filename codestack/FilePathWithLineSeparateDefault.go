package codestack

func FilePathWithLineSeparateDefault() (
	filePath string, lineNumber int,
) {
	return FilePathWithLineSeparate(defaultInternalSkip)
}
