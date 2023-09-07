package loggerinf

type FormatWarningLogger interface {
	WarnFmt(format string, args ...interface{}) // Warn logs a message at Warning level.
}
