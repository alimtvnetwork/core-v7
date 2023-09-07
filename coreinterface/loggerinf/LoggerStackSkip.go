package loggerinf

type LoggerStackSkip interface {
	FatalStackSkip(
		stackSkipIndex int,
		args ...interface{},
	)
	ErrorStackSkip(
		stackSkipIndex int,
		args ...interface{},
	)

	// ErrorUsingError
	//
	// Skip if no error
	ErrorUsingError(err error)

	// ErrorUsingErrorStackSkip
	//
	// Skip if no error
	ErrorUsingErrorStackSkip(
		stackSkipIndex int,
		err error,
	)

	ErrorIf(isLog bool, args ...interface{})
	DebugFmtIf(
		isLog bool,
		formatter string,
		args ...interface{},
	)
	DebugFmtStackSkip(
		stackSkipIndex int,
		format string,
		args ...interface{},
	)

	DebugIf(isLog bool, args ...interface{}) // Debug logs a message at Debug level.
	DebugStackSkip(
		stackSkipIndex int,
		args ...interface{},
	)

	DebugIncludingStackTracesIf(
		isLog bool,
		stackSkipIndex int,
		args ...interface{},
	)
	DebugIncludingStackTraces(
		stackSkipIndex int,
		args ...interface{},
	)
}
