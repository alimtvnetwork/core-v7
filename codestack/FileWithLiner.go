package codestack

type FileWithLiner interface {
	FullFilePath() string
	LineNumber() int
}
