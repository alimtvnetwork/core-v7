package codestack

func FilePathWithLineStringDefault() string {
	stack := New(Skip1)
	fileWithLine := stack.FileWithLineString()
	stack.Dispose()

	return fileWithLine
}
