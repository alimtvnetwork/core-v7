package codestack

func FilePathWithLineSeparate(
	skipStack int,
) (
	filePath string, lineNumber int,
) {
	stack := New(Skip1 + skipStack)
	fileWithLine := stack.FileWithLine()
	filePath = fileWithLine.FullFilePath()
	lineNumber = fileWithLine.LineNumber()

	stack.Dispose()

	return filePath, lineNumber
}
