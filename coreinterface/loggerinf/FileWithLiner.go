package loggerinf

type FileWithLiner interface {
	FullFilePath() string
	LineNumber() int
}
