package loggerinf

type FormatFatalLogger interface {
	// FatalFmt
	//
	// logs a message at Fatal level
	// and process will exit with status set to 1.
	FatalFmt(format string, args ...interface{})
	FatalFmtStackSkip(
		stackSkipIndex int,
		format string,
		args ...interface{},
	)
}
