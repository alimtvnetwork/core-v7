package loggerinf

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface/enuminf"
	"gitlab.com/auk-go/core/coreinterface/errcoreinf"
)

type StandardLogger interface {
	StandardLoggerChecker
	ConditionalStandardLogger

	FullLogger() FullLogger
	EnvOptioner() enuminf.EnvironmentOptioner

	TaskNamedLogger(
		taskName string,
	) StandardLogger

	TaskWithPayloadLogger(
		taskName string,
		payloadAny interface{}, // can be bytes, payloadWrapper, can be any
	) StandardLogger

	GetLoggerByTaskName(taskName string) StandardLogger
	GetLoggerByTaskNamer(taskNamer enuminf.Namer) StandardLogger

	Success(args ...interface{}) StandardLogger
	Info(args ...interface{}) StandardLogger
	Trace(args ...interface{}) StandardLogger
	Debug(args ...interface{}) StandardLogger
	Warn(args ...interface{}) StandardLogger
	Error(args ...interface{}) StandardLogger
	Fatal(args ...interface{}) StandardLogger
	Panic(args ...interface{}) StandardLogger

	SuccessFmt(format string, args ...interface{}) StandardLogger
	InfoFmt(format string, args ...interface{}) StandardLogger
	TraceFmt(format string, args ...interface{}) StandardLogger
	DebugFmt(format string, args ...interface{}) StandardLogger
	WarnFmt(format string, args ...interface{}) StandardLogger
	ErrorFmt(format string, args ...interface{}) StandardLogger
	FatalFmt(format string, args ...interface{}) StandardLogger
	PanicFmt(format string, args ...interface{}) StandardLogger

	SuccessExtend() SingleLogger
	InfoExtend() SingleLogger
	TraceExtend() SingleLogger
	DebugExtend() SingleLogger
	WarnExtend() SingleLogger
	FatalExtend() SingleLogger
	PanicExtend() SingleLogger

	ErrorDirect(err error) StandardLogger
	OnErrStackTrace(err error) StandardLogger
	ErrInterface(errInf errcoreinf.BaseErrorOrCollectionWrapper) StandardLogger
	ErrInterfaceStackTraces(errInfWithStackTraces errcoreinf.BaseErrorOrCollectionWrapper) StandardLogger

	ReflectSetter

	InfoOrError(isError bool) SingleLogger
	Log(loggerType enuminf.LoggerTyper) StandardLogger

	ErrorJsoner(jsoner corejson.Jsoner) StandardLogger
	DebugJsoner(jsoner corejson.Jsoner) StandardLogger
	ErrorJsonerTitle(title string, jsoner corejson.Jsoner) StandardLogger
	DebugJsonerTitle(title string, jsoner corejson.Jsoner) StandardLogger
}
