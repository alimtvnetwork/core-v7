package loggerinf

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
)

type StandardLogger interface {
	StandardLoggerChecker
	ConditionalStandardLogger

	FullLogger() FullLogger
	EnvOptioner() enuminf.EnvironmentOptioner

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
	ErrInterface(errorWrapperOrCollection errcoreinf.BaseErrorOrCollectionWrapper) StandardLogger
	ErrInterfaceStackTraces(errorWrapperOrCollection errcoreinf.BaseErrorOrCollectionWrapper) StandardLogger

	ReflectSetter

	ErrorJsoner(jsoner corejson.Jsoner) StandardLogger
	DebugJsoner(jsoner corejson.Jsoner) StandardLogger
	ErrorJsonerTitle(title string, jsoner corejson.Jsoner) StandardLogger
	DebugJsonerTitle(title string, jsoner corejson.Jsoner) StandardLogger
}
