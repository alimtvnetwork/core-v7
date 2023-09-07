package loggerinf

type FormatErrorLogger interface {
	ErrorFmt(
		format string,
		args ...interface{},
	)
	ErrorFmtIf(
		isLog bool,
		format string,
		args ...interface{},
	)
	ErrorFmtStackSkip(
		stackSkipIndex int,
		format string,
		args ...interface{},
	)

	// ErrorFmtUsingError
	//
	// Skip if no error
	ErrorFmtUsingError(
		format string,
		err error,
	)

	// ErrorFmtUsingErrorStackSkip
	//
	// Skip if no error
	ErrorFmtUsingErrorStackSkip(
		stackSkipIndex int,
		format string,
		err error,
	)
	WarnFmtStackSkip(
		stackSkipIndex int,
		format string,
		args ...interface{},
	)
	InfoFmtStackSkip(
		stackSkipIndex int,
		format string,
		args ...interface{},
	)
	WarnStackSkip(
		stackSkipIndex int,
		args ...interface{},
	)

	InfoStackSkip(
		stackSkipIndex int,
		args ...interface{},
	)
}
