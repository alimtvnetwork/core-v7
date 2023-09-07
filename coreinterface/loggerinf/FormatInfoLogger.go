package loggerinf

type FormatInfoLogger interface {
	InfoFmt(format string, args ...interface{}) // Info logs a message at Info level.
}
