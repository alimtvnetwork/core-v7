package loggerinf

type Caller interface {
	Line() int
	File() string
}
