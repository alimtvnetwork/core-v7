package loggerinf

import (
	"io"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/internal/internalinterface"
)

type PersistentLogger interface {
	LogPathExtender() PathExtender
	IsRotating() bool
	IsDbLogger() bool
	IsFileLogger() bool

	Config() interface{}
	ConfigReflectSetTo(toPointer interface{}) error
	internalinterface.IdStringerWithNamer

	Info(message string) PersistentLogger
	Trace(message string) PersistentLogger
	Debug(message string) PersistentLogger
	Warn(message string) PersistentLogger
	Error(message string) PersistentLogger
	Fatal(message string) PersistentLogger
	Panic(message string) PersistentLogger

	InfoAttr(message, attr string) PersistentLogger
	TraceAttr(message, attr string) PersistentLogger
	DebugAttr(message, attr string) PersistentLogger
	WarnAttr(message, attr string) PersistentLogger
	ErrorAttr(message, attr string) PersistentLogger
	FatalAttr(message, attr string) PersistentLogger
	PanicAttr(message, attr string) PersistentLogger

	InfoStackSkip(stackSkipIndex int, message string) PersistentLogger
	TraceStackSkip(stackSkipIndex int, message string) PersistentLogger
	DebugStackSkip(stackSkipIndex int, message string) PersistentLogger
	WarnStackSkip(stackSkipIndex int, message string) PersistentLogger
	ErrorStackSkip(stackSkipIndex int, message string) PersistentLogger
	FatalStackSkip(stackSkipIndex int, message string) PersistentLogger
	PanicStackSkip(stackSkipIndex int, message string) PersistentLogger

	InfoJson(jsonResult *corejson.Result) PersistentLogger
	ErrorJson(jsonResult *corejson.Result) PersistentLogger
	DebugJson(jsonResult *corejson.Result) PersistentLogger

	InfoBytes(rawBytes []byte) PersistentLogger
	ErrorBytes(rawBytes []byte) PersistentLogger
	DebugBytes(rawBytes []byte) PersistentLogger

	InfoTitleBytes(title string, rawBytes []byte) PersistentLogger
	ErrorTitleBytes(title string, rawBytes []byte) PersistentLogger
	DebugTitleBytes(title string, rawBytes []byte) PersistentLogger

	Log(message string) PersistentLogger
	LogRaw(logType LogTypeChecker, message, attr string) PersistentLogger
	LogRawStackSkip(stackSkipIndex int, logType LogTypeChecker, message, attr string) PersistentLogger
	Jsoner(logType LogTypeChecker, message string, jsonResult *corejson.Result) PersistentLogger
	JsonerStackSkip(
		stackSkipIndex int, logType LogTypeChecker, message string, jsonResult *corejson.Result,
	) PersistentLogger

	LogStackSkip(stackSkipIndex int, message string) PersistentLogger
	LogFmtStackSkip(stackSkipIndex int, format string, message string) PersistentLogger
	LogAttr(message, attr string) PersistentLogger
	LogAttrStackSkip(stackSkipIndex int, message, attr string) PersistentLogger

	AnErr(err error) PersistentLogger
	ErrorMessage(message string) PersistentLogger
	ErrorMessageAttr(message, attr string) PersistentLogger
	ErrorMessageStackSkip(stackSkipIndex int, message string) PersistentLogger
	ErrorMessageAttrStackSkip(stackSkipIndex int, message, attr string) PersistentLogger
	ErrorMessageFmtStackSkip(stackSkipIndex int, format string, message string) PersistentLogger

	DebugMessage(message string) PersistentLogger
	DebugMessageAttr(message, attr string) PersistentLogger
	Err(err error) PersistentLogger

	FullTraceAsAttr(
		title string,
		attrFullStringWithTraces errcoreinf.FullStringWithTracesGetter,
	) PersistentLogger

	FullTraceAsAttrStackSkip(
		stackSkipIndex int,
		title string,
		attrFullStringWithTraces errcoreinf.FullStringWithTracesGetter,
	) PersistentLogger

	FullStringWithTracesOptions(
		logType LogTypeChecker,
		fullStringWithTraces errcoreinf.FullStringWithTracesGetter,
	) PersistentLogger

	// FullStringWithTraces Log as error
	FullStringWithTraces(
		fullStringWithTraces errcoreinf.FullStringWithTracesGetter,
	) PersistentLogger
	BaseRawErrCollectionDefiner(
		rawErrCollection errcoreinf.BaseRawErrCollectionDefiner,
	) PersistentLogger
	BasicErrWrapper(basicErrWrapper errcoreinf.BasicErrWrapper) PersistentLogger
	BasicErrWrapperOptions(
		logType LogTypeChecker,
		basicErrWrapper errcoreinf.BasicErrWrapper,
		attributes string,
	) PersistentLogger
	BasicErrWrapperOptionsStackSkip(
		stackSkipIndex int,
		logType LogTypeChecker,
		basicErrWrapper errcoreinf.BasicErrWrapper,
		attributes string,
	) PersistentLogger
	ErrOptions(
		logType LogTypeChecker,
		err error,
		attributes string,
	) PersistentLogger
	persistentAllStacktraceLogger

	NewGeneralWriter
	AllLogWriter
	persistentAllParamsLogger
	ConditionalPersistentLogger

	InfoAnyAttr(anyItem interface{}, attr string)
	TraceAnyAttr(anyItem interface{}, attr string)
	DebugAnyAttr(anyItem interface{}, attr string)
	WarnAnyAttr(anyItem interface{}, attr string)
	ErrorAnyAttr(anyItem interface{}, attr string)
	FatalAnyAttr(anyItem interface{}, attr string)
	PanicAnyAttr(anyItem interface{}, attr string)

	DebugAnyAttrAny(anyItem, attr interface{})

	InfoAny(anyItem interface{})
	TraceAny(anyItem interface{})
	DebugAny(anyItem interface{})
	WarnAny(anyItem interface{})
	ErrorAny(anyItem interface{})
	FatalAny(anyItem interface{})
	PanicAny(anyItem interface{})

	InfoAnyStackSkip(
		stackSkip int,
		anyItem interface{},
	)
	TraceAnyStackSkip(
		stackSkip int,
		anyItem interface{},
	)
	DebugAnyStackSkip(
		stackSkip int,
		anyItem interface{},
	)
	WarnAnyStackSkip(
		stackSkip int,
		anyItem interface{},
	)
	ErrorAnyStackSkip(
		stackSkip int,
		anyItem interface{},
	)
	FatalAnyStackSkip(
		stackSkip int,
		anyItem interface{},
	)
	PanicAnyStackSkip(
		stackSkip int,
		anyItem interface{},
	)
}

type persistentAllParamsLogger interface {
	LogAll(
		logTyper LoggerTyper,
		message, attributes string,
	) PersistentLogger
	LogAllUsingStackSkip(
		stackSkipIndex int,
		logTyper LoggerTyper,
		message, attributes string,
	) PersistentLogger
	LogAllUsingConfig(
		config Configurer,
		message, attributes string,
	) PersistentLogger
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
