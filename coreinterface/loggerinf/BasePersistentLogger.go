package loggerinf

import (
	"io"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
)

type BasePersistentLogger interface {
	Attr() AttrPersistentLogger

	Success(message string) BasePersistentLogger
	Info(message string) BasePersistentLogger
	Trace(message string) BasePersistentLogger
	Debug(message string) BasePersistentLogger
	Warn(message string) BasePersistentLogger
	Error(message string) BasePersistentLogger
	Fatal(message string) BasePersistentLogger
	Panic(message string) BasePersistentLogger

	SuccessAttr(message, attr string) BasePersistentLogger
	InfoAttr(message, attr string) BasePersistentLogger
	TraceAttr(message, attr string) BasePersistentLogger
	DebugAttr(message, attr string) BasePersistentLogger
	WarnAttr(message, attr string) BasePersistentLogger
	ErrorAttr(message, attr string) BasePersistentLogger
	FatalAttr(message, attr string) BasePersistentLogger
	PanicAttr(message, attr string) BasePersistentLogger

	InfoStackSkip(stackSkipIndex int, message string) BasePersistentLogger
	TraceStackSkip(stackSkipIndex int, message string) BasePersistentLogger
	DebugStackSkip(stackSkipIndex int, message string) BasePersistentLogger
	WarnStackSkip(stackSkipIndex int, message string) BasePersistentLogger
	ErrorStackSkip(stackSkipIndex int, message string) BasePersistentLogger
	FatalStackSkip(stackSkipIndex int, message string) BasePersistentLogger
	PanicStackSkip(stackSkipIndex int, message string) BasePersistentLogger

	SuccessJson(jsonResult *corejson.Result) BasePersistentLogger
	InfoJson(jsonResult *corejson.Result) BasePersistentLogger
	ErrorJson(jsonResult *corejson.Result) BasePersistentLogger
	WarnJson(jsonResult *corejson.Result) BasePersistentLogger
	DebugJson(jsonResult *corejson.Result) BasePersistentLogger
	TypeJson(
		logType LogTypeChecker,
		jsonResult *corejson.Result,
	) BasePersistentLogger

	SuccessJsoner(jsoner corejson.Jsoner) BasePersistentLogger
	InfoJsoner(jsoner corejson.Jsoner) BasePersistentLogger
	ErrorJsoner(jsoner corejson.Jsoner) BasePersistentLogger
	WarnJsoner(jsoner corejson.Jsoner) BasePersistentLogger
	DebugJsoner(jsoner corejson.Jsoner) BasePersistentLogger
	TypeJsoner(
		logType LogTypeChecker,
		jsoner corejson.Jsoner,
	) BasePersistentLogger

	InfoBytes(rawBytes []byte) BasePersistentLogger
	ErrorBytes(rawBytes []byte) BasePersistentLogger
	DebugBytes(rawBytes []byte) BasePersistentLogger
	SuccessBytes(rawBytes []byte) BasePersistentLogger

	InfoTitleBytes(title string, rawBytes []byte) BasePersistentLogger
	ErrorTitleBytes(title string, rawBytes []byte) BasePersistentLogger
	DebugTitleBytes(title string, rawBytes []byte) BasePersistentLogger
	SuccessTitleBytes(title string, rawBytes []byte) BasePersistentLogger

	Log(message string) BasePersistentLogger
	LogRaw(logType LogTypeChecker, message, attr string) BasePersistentLogger
	LogRawStackSkip(
		stackSkipIndex int,
		logType LogTypeChecker,
		message, attr string,
	) BasePersistentLogger
	JsonResultOptions(
		logType LogTypeChecker,
		message string,
		jsonResult *corejson.Result,
	) BasePersistentLogger
	JsonResultOptionsStackSkip(
		stackSkipIndex int,
		logType LogTypeChecker,
		message string,
		jsonResult *corejson.Result,
	) BasePersistentLogger

	LogStackSkip(stackSkipIndex int, message string) BasePersistentLogger
	LogFmtStackSkip(
		stackSkipIndex int,
		format string,
		v ...interface{},
	) BasePersistentLogger

	LogAttr(message, attr string) BasePersistentLogger
	LogAttrStackSkip(stackSkipIndex int, message, attr string) BasePersistentLogger

	AnErr(err error) BasePersistentLogger
	ErrorMessage(message string) BasePersistentLogger
	ErrorMessageAttr(message, attr string) BasePersistentLogger
	ErrorMessageStackSkip(stackSkipIndex int, message string) BasePersistentLogger
	ErrorMessageAttrStackSkip(stackSkipIndex int, message, attr string) BasePersistentLogger
	ErrorMessageFmtStackSkip(stackSkipIndex int, format string, message string) BasePersistentLogger

	DebugMessage(message string) BasePersistentLogger
	DebugMessageAttr(message, attr string) BasePersistentLogger
	Err(err error) BasePersistentLogger

	FullTraceAsAttr(
		title string,
		attrFullStringWithTraces errcoreinf.FullStringWithTracesGetter,
	) BasePersistentLogger

	FullTraceAsAttrStackSkip(
		stackSkipIndex int,
		title string,
		attrFullStringWithTraces errcoreinf.FullStringWithTracesGetter,
	) BasePersistentLogger

	FullStringWithTracesOptions(
		logType LogTypeChecker,
		fullStringWithTraces errcoreinf.FullStringWithTracesGetter,
	) BasePersistentLogger

	// FullStringWithTraces Log as error
	FullStringWithTraces(
		fullStringWithTraces errcoreinf.FullStringWithTracesGetter,
	) BasePersistentLogger
	BaseRawErrCollectionDefiner(
		rawErrCollection errcoreinf.BaseRawErrCollectionDefiner,
	) BasePersistentLogger
	BasicErrWrapper(basicErrWrapper errcoreinf.BasicErrWrapper) BasePersistentLogger
	BasicErrWrapperOptions(
		logType LogTypeChecker,
		basicErrWrapper errcoreinf.BasicErrWrapper,
		attributes string,
	) BasePersistentLogger
	BasicErrWrapperOptionsStackSkip(
		stackSkipIndex int,
		logType LogTypeChecker,
		basicErrWrapper errcoreinf.BasicErrWrapper,
		attributes string,
	) BasePersistentLogger
	ErrOptions(
		logType LogTypeChecker,
		err error,
		attributes string,
	) BasePersistentLogger
	persistentAllStacktraceLogger

	NewGeneralWriter
	AllLogWriter
	persistentAllParamsLogger
	ConditionalBasePersistentLogger

	SuccessAnyAttr(anyItem interface{}, attr string) BasePersistentLogger
	InfoAnyAttr(anyItem interface{}, attr string) BasePersistentLogger
	TraceAnyAttr(anyItem interface{}, attr string) BasePersistentLogger
	DebugAnyAttr(anyItem interface{}, attr string) BasePersistentLogger
	WarnAnyAttr(anyItem interface{}, attr string) BasePersistentLogger
	ErrorAnyAttr(anyItem interface{}, attr string) BasePersistentLogger
	FatalAnyAttr(anyItem interface{}, attr string) BasePersistentLogger
	PanicAnyAttr(anyItem interface{}, attr string) BasePersistentLogger

	DebugAnyAttrAny(anyItem, attr interface{}) BasePersistentLogger

	InfoFmt(formatter string, v ...interface{}) BasePersistentLogger
	DebugFmt(formatter string, v ...interface{}) BasePersistentLogger
	ErrorFmt(formatter string, v ...interface{}) BasePersistentLogger
	FatalFmt(formatter string, v ...interface{}) BasePersistentLogger

	LogFmt(
		logType LoggerTyper,
		formatter string,
		v ...interface{},
	) BasePersistentLogger

	InfoAny(anyItem interface{}) BasePersistentLogger
	TraceAny(anyItem interface{}) BasePersistentLogger
	DebugAny(anyItem interface{}) BasePersistentLogger
	WarnAny(anyItem interface{}) BasePersistentLogger
	ErrorAny(anyItem interface{}) BasePersistentLogger
	FatalAny(anyItem interface{}) BasePersistentLogger
	PanicAny(anyItem interface{}) BasePersistentLogger
	Session(name string) (newLogger BasePersistentLogger)
	SessionName() string

	SessionAllString() string
	SessionAll() ModelCollectioner
	SessionAllByType(logType LoggerTyper) ModelCollectioner
	SessionAllSuccess(logType LoggerTyper) ModelCollectioner

	GenericSubscribe(subscriber func(modeler SingleLogModeler))
	SpecificSubscribe(logType LoggerTyper, subscriber func(modeler SingleLogModeler))
	SpecificFilterSubscribe(filterType enuminf.BasicEnumer, subscriber func(modeler SingleLogModeler))
	SpecificValueSubscribe(logType LoggerTyper, subscriber func(valuer SpecificValuer))

	CompleteSession(completionTyper enuminf.BasicEnumer)

	SuccessComplete()
	ErrorComplete(errWp errcoreinf.BasicErrWrapper)

	ModelEntry(
		stackSkip int,
		modeler SingleLogModeler,
	) BasePersistentLogger

	InfoAnyStackSkip(
		stackSkip int,
		anyItem interface{},
	) BasePersistentLogger
	TraceAnyStackSkip(
		stackSkip int,
		anyItem interface{},
	) BasePersistentLogger
	DebugAnyStackSkip(
		stackSkip int,
		anyItem interface{},
	) BasePersistentLogger
	WarnAnyStackSkip(
		stackSkip int,
		anyItem interface{},
	) BasePersistentLogger
	ErrorAnyStackSkip(
		stackSkip int,
		anyItem interface{},
	) BasePersistentLogger
	FatalAnyStackSkip(
		stackSkip int,
		anyItem interface{},
	) BasePersistentLogger
	PanicAnyStackSkip(
		stackSkip int,
		anyItem interface{},
	) BasePersistentLogger
}

type persistentAllParamsLogger interface {
	LogAll(
		logType LoggerTyper,
		message, attributes string,
	) BasePersistentLogger
	LogAllUsingStackSkip(
		stackSkipIndex int,
		logType LoggerTyper,
		message, attributes string,
	) BasePersistentLogger
	LogAllUsingConfig(
		config Configurer,
		message, attributes string,
	) BasePersistentLogger
}

type NewGeneralWriter interface {
	NewGeneralWriter(writeConfigurer WriterConfigurer) io.Writer
}

type Configurer interface {
	LoggerTyper() LoggerTyper
	StackSkipIndex() int
}

type WriterConfigurer interface {
	Configurer
	AdditionalConfigProcessor
}

type AdditionalConfigProcessor interface {
	AdditionalConfigBytes() []byte
	AdditionalConfigProcess() error
}
