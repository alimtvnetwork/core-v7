package loggerinf

import (
	"fmt"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface/entityinf"
	"gitlab.com/auk-go/core/coreinterface/enuminf"
	"gitlab.com/auk-go/core/coreinterface/errcoreinf"
	"gitlab.com/auk-go/core/coreinterface/serializerinf"
)

type SingleLogger interface {
	enuminf.LoggerTyperGetter

	IsSilent() bool

	Stack() MetaAttributesStacker
	StackTitle(title string) MetaAttributesStacker

	On(isLog bool) SingleLogger
	StackSkip(stackSkipIndex int) SingleLogger
	OnString(input, expected string) SingleLogger

	Title(message string) SingleLogger
	Msg(message string) SingleLogger
	TitleAttr(message, attr string) SingleLogger
	Log(message string) SingleLogger
	LogAttr(message, attr string) SingleLogger
	Str(title, val string) SingleLogger
	Strings(title string, values []string) SingleLogger
	StringsSpread(title string, values ...string) SingleLogger
	Stringer(title string, stringer fmt.Stringer) SingleLogger
	Stringers(title string, stringers ...fmt.Stringer) SingleLogger
	Byte(title string, val byte) SingleLogger
	Bytes(title string, values []byte) SingleLogger
	Hex(title string, val []byte) SingleLogger
	RawJson(title string, rawJson []byte) SingleLogger
	Err(err error) SingleLogger
	AnErr(title string, err error) SingleLogger

	SimpleBytesResulter(
		title string,
		result serializerinf.SimpleBytesResulter,
	) MetaAttributesStacker

	BaseJsonResulter(
		title string,
		result serializerinf.BaseJsonResulter,
	) MetaAttributesStacker

	BasicJsonResulter(
		title string,
		result serializerinf.BasicJsonResulter,
	) MetaAttributesStacker
	JsonResulter(
		title string,
		result serializerinf.JsonResulter,
	) MetaAttributesStacker

	ErrWithType(title string, errType errcoreinf.BasicErrorTyper, err error) SingleLogger
	Meta(title string, metaAttr MetaAttributesCompiler) SingleLogger

	MapBool(title string, mapInt map[string]bool) SingleLogger
	MapInt(title string, mapInt map[string]int) SingleLogger
	MapAnyAny(title string, mapAny map[interface{}]interface{}) SingleLogger
	MapAny(title string, mapAny map[string]interface{}) SingleLogger
	MapIntAny(title string, mapAny map[int]interface{}) SingleLogger
	MapIntString(title string, mapAny map[int]string) SingleLogger
	MapJsonResult(title string, mapAny map[string]corejson.Result) SingleLogger

	DefaultStackTraces() SingleLogger
	ErrWithTypeTraces(title string, errType errcoreinf.BasicErrorTyper, err error) SingleLogger
	ErrorsWithTypeTraces(title string, errType errcoreinf.BasicErrorTyper, errorItems ...error) SingleLogger
	StackTraces(stackSkipIndex int, title string) SingleLogger
	OnErrStackTraces(err error) SingleLogger
	OnErrWrapperOrCollectionStackTraces(errWrapperOrCollection errcoreinf.BaseErrorOrCollectionWrapper) SingleLogger

	FullStringer(
		fullStringer errcoreinf.FullStringer,
	) SingleLogger

	FullStringerTitle(
		title string,
		fullStringer errcoreinf.FullStringer,
	) SingleLogger
	FullTraceAsAttr(
		title string,
		attrFullStringWithTraces errcoreinf.FullStringWithTracesGetter,
	) SingleLogger

	BasicErrWrapper(basicErrWrapper errcoreinf.BasicErrWrapper) SingleLogger
	BaseRawErrCollectionDefiner(rawErrCollection errcoreinf.BaseRawErrCollectionDefiner) SingleLogger
	BaseErrorWrapperCollectionDefiner(errWrapperCollection errcoreinf.BaseErrorWrapperCollectionDefiner) SingleLogger
	ErrWrapperOrCollection(errWrapperOrCollection errcoreinf.BaseErrorOrCollectionWrapper) SingleLogger
	RawErrCollection(title string, err errcoreinf.BaseRawErrCollectionDefiner) SingleLogger
	CompiledBasicErrWrapper(compiler errcoreinf.CompiledBasicErrWrapper) SingleLogger

	Namer(title string, namer enuminf.Namer) SingleLogger
	Enum(title string, enum enuminf.BasicEnumer) SingleLogger
	Enums(title string, enums ...enuminf.BasicEnumer) SingleLogger

	OnlyNamer(namer enuminf.Namer) SingleLogger
	OnlyEnum(enum enuminf.BasicEnumer) SingleLogger
	OnlyEnums(enums ...enuminf.BasicEnumer) SingleLogger
	OnlyError(err error) SingleLogger
	OnlyString(value string) SingleLogger
	OnlyStrings(values ...string) SingleLogger
	OnlyMetaAttr(metaAttr MetaAttributesCompiler) SingleLogger

	OnlyStringer(stringer fmt.Stringer) SingleLogger
	OnlyStringers(stringers ...fmt.Stringer) SingleLogger

	OnlyIntegers(values ...int) SingleLogger
	OnlyBooleans(values ...bool) SingleLogger
	OnlyBytes(rawBytes []byte) SingleLogger
	OnlyRawJson(rawBytes []byte) SingleLogger
	OnlyBytesErr(rawBytes []byte, err error) SingleLogger

	OnlyAny(anyItem interface{}) SingleLogger
	OnlyAnyItems(values ...interface{}) SingleLogger
	OnlyAnyIf(isLog bool, anyItem interface{}) SingleLogger
	OnlyAnyItemsIf(isLog bool, anyItems ...interface{}) SingleLogger

	Bool(title string, isResult bool) SingleLogger
	Booleans(title string, isResults ...bool) SingleLogger

	OnlyMapBool(mapInt map[string]bool) SingleLogger
	OnlyMapInt(mapInt map[string]int) SingleLogger
	OnlyMapAny(mapAny map[string]interface{}) SingleLogger
	OnlyMapAnyAny(mapAny map[interface{}]interface{}) SingleLogger
	OnlyMapIntAny(mapAny map[int]interface{}) SingleLogger
	OnlyMapIntString(mapAny map[int]string) SingleLogger
	OnlyMapJsonResult(mapAny map[string]corejson.Result) SingleLogger

	OnlySimpleBytesResulter(
		result serializerinf.SimpleBytesResulter,
	) SingleLogger

	OnlyBaseJsonResulter(
		result serializerinf.BaseJsonResulter,
	) SingleLogger

	OnlyBasicJsonResulter(
		result serializerinf.BasicJsonResulter,
	) SingleLogger
	OnlyJsonResulter(
		result serializerinf.JsonResulter,
	) SingleLogger

	AnyJsonLog(anyItem interface{}) SingleLogger
	Any(anyItem interface{}) SingleLogger
	AnyIf(isLog bool, anyItem interface{}) SingleLogger
	AnyItems(anyItems ...interface{}) SingleLogger
	AnyItemsIf(isLog bool, anyItems ...interface{}) SingleLogger

	OnlyJson(json *corejson.Result) SingleLogger
	OnlyJsons(jsons ...*corejson.Result) SingleLogger

	Jsoner(title string, jsoner corejson.Jsoner) SingleLogger
	Jsoners(jsoners ...corejson.Jsoner) SingleLogger
	OnlyJsoner(jsoner corejson.Jsoner) SingleLogger

	Serializer(serializer Serializer) SingleLogger
	Serializers(serializers ...Serializer) SingleLogger
	SerializerFunc(serializerFunc func() ([]byte, error)) SingleLogger
	SerializerFunctions(serializerFunctions ...func() ([]byte, error)) SingleLogger

	StandardTaskEntityDefiner(entity entityinf.StandardTaskEntityDefiner) SingleLogger
	TaskEntityDefiner(entity entityinf.TaskEntityDefiner) SingleLogger

	StandardTaskEntityDefinerTitle(title string, entity entityinf.StandardTaskEntityDefiner) SingleLogger
	TaskEntityDefinerTitle(title string, entity entityinf.TaskEntityDefiner) SingleLogger

	LogModel(model SingleLogModeler) SingleLogger
	LogModelTitle(title string, model SingleLogModeler) SingleLogger

	Int(title string, i int) SingleLogger
	Integers(title string, integerItems ...int) SingleLogger

	FmtIf(isLog bool, format string, v ...interface{}) SingleLogger
	Fmt(format string, v ...interface{}) SingleLogger
	AttrFmt(title string, attrFormat string, attrValues ...interface{}) SingleLogger

	RawPayloadsGetter(payloadsGetter RawPayloadsGetter) SingleLogger
	RawPayloadsGetterTitle(title string, payloadsGetter RawPayloadsGetter) SingleLogger

	Logger() StandardLogger
}
