package loggerinf

import "gitlab.com/evatix-go/core/coreinterface/errcoreinf"

type StandardLogger interface {
	Info(args ...interface{}) StandardLogger
	Trace(args ...interface{}) StandardLogger
	Debug(args ...interface{}) StandardLogger
	Warn(args ...interface{}) StandardLogger
	Error(args ...interface{}) StandardLogger
	Fatal(args ...interface{}) StandardLogger
	Panic(args ...interface{}) StandardLogger

	InfoStackSkip(stackSkipIndex int, args ...interface{}) StandardLogger
	TraceStackSkip(stackSkipIndex int, args ...interface{}) StandardLogger
	DebugStackSkip(stackSkipIndex int, args ...interface{}) StandardLogger
	WarnStackSkip(stackSkipIndex int, args ...interface{}) StandardLogger
	ErrorStackSkip(stackSkipIndex int, args ...interface{}) StandardLogger
	FatalStackSkip(stackSkipIndex int, args ...interface{}) StandardLogger
	PanicStackSkip(stackSkipIndex int, args ...interface{}) StandardLogger

	InfoFmt(format string, args ...interface{}) StandardLogger
	TraceFmt(format string, args ...interface{}) StandardLogger
	DebugFmt(format string, args ...interface{}) StandardLogger
	WarnFmt(format string, args ...interface{}) StandardLogger
	ErrorFmt(format string, args ...interface{}) StandardLogger
	FatalFmt(format string, args ...interface{}) StandardLogger
	PanicFmt(format string, args ...interface{}) StandardLogger

	InfoFmtStackSkip(stackSkipIndex int, format string, args ...interface{}) StandardLogger
	TraceFmtStackSkip(stackSkipIndex int, format string, args ...interface{}) StandardLogger
	DebugFmtStackSkip(stackSkipIndex int, format string, args ...interface{}) StandardLogger
	WarnFmtStackSkip(stackSkipIndex int, format string, args ...interface{}) StandardLogger
	ErrorFmtStackSkip(stackSkipIndex int, format string, args ...interface{}) StandardLogger
	FatalFmtStackSkip(stackSkipIndex int, format string, args ...interface{}) StandardLogger
	PanicFmtStackSkip(stackSkipIndex int, format string, args ...interface{}) StandardLogger

	Log(message string) StandardLogger
	LogStackSkip(stackSkipIndex int, message string) StandardLogger
	LogFmtStackSkip(stackSkipIndex int, format string, message string) StandardLogger
	LogAttr(message, attr string) StandardLogger
	LogAttrStackSkip(stackSkipIndex int, message, attr string) StandardLogger

	ErrorMessage(message string) StandardLogger
	ErrorMessageAttr(message, attr string) StandardLogger
	ErrorMessageStackSkip(stackSkipIndex int, message string) StandardLogger
	ErrorMessageAttrStackSkip(stackSkipIndex int, message, attr string) StandardLogger
	ErrorMessageFmtStackSkip(stackSkipIndex int, format string, message string) StandardLogger

	DebugMessage(message string) StandardLogger
	DebugMessageAttr(message, attr string) StandardLogger
	Err(err error) StandardLogger
	// ErrorStackTraces
	//
	// Includes stack-traces
	ErrorStackTraces(err error) StandardLogger
	// DebugStackTraces
	//
	// Includes stack-traces
	DebugStackTraces(message string) StandardLogger
	// DebugAttrStackTraces
	//
	// Includes stack-traces
	DebugAttrStackTraces(message, attr string) StandardLogger
	StackTraces() StandardLogger
	StackTracesSkip(stackSkipIndex int) StandardLogger
	TitleStackTraces(title string) StandardLogger
	TitleStackTracesSkip(stackSkipIndex int, title string) StandardLogger

	FullTraceAsAttr(
		title string,
		attrFullStringWithTraces errcoreinf.FullStringWithTracesGetter,
	) StandardLogger
	FullTraceAsAttrStackSkip(
		stackSkipIndex int,
		title string,
		attrFullStringWithTraces errcoreinf.FullStringWithTracesGetter,
	) StandardLogger
	FullStringWithTracesOptions(
		logType LogTypeChecker,
		fullStringWithTraces errcoreinf.FullStringWithTracesGetter,
	) StandardLogger

	// FullStringWithTraces Log as error
	FullStringWithTraces(
		fullStringWithTraces errcoreinf.FullStringWithTracesGetter,
	) StandardLogger

	BaseRawErrCollectionDefiner(
		rawErrCollection errcoreinf.BaseRawErrCollectionDefiner,
	) StandardLogger
	BasicErrWrapper(basicErrWrapper errcoreinf.BasicErrWrapper) StandardLogger
	BasicErrWrapperOptions(
		logType LogTypeChecker,
		basicErrWrapper errcoreinf.BasicErrWrapper,
		attributes string,
	) StandardLogger
	BasicErrWrapperOptionsStackSkip(
		stackSkipIndex int,
		logType LogTypeChecker,
		basicErrWrapper errcoreinf.BasicErrWrapper,
		attributes string,
	) StandardLogger

	ConditionalStandardLogger
}
