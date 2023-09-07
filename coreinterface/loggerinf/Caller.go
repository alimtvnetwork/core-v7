package loggerinf

type Caller interface {
	LineNumber() int
	FullFilePath() string
}
