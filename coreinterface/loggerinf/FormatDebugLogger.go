package loggerinf

type FormatDebugLogger interface {
	DebugFmt(formatter string, args ...interface{}) // Debug logs a message at Debug level.
}
