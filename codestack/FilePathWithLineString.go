package codestack

func FilePathWithLineString(
	skipStack int,
) string {
	stack := New(Skip1 + skipStack)
	fileWithLine := stack.FileWithLineString()
	stack.Dispose()

	return fileWithLine
}
