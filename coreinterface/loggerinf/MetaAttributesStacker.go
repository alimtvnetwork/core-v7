package loggerinf

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/coreinterface/entityinf"
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
)

type MetaAttributesStacker interface {
	LoggerTyperGetter

	On(isLog bool) MetaAttributesStacker
	OnTitle(isLog bool, title string) MetaAttributesWithoutTileStacker

	Title(title string) MetaAttributesStacker
	TitleAttr(title, attr string) MetaAttributesStacker
	Str(title, val string) MetaAttributesStacker
	Strings(title string, stringItems ...string) MetaAttributesStacker
	StandardSlicer(title string, standardSlice coreinterface.StandardSlicer) MetaAttributesStacker
	Stringer(title string, stringer fmt.Stringer) MetaAttributesStacker
	Stringers(title string, stringers ...fmt.Stringer) MetaAttributesStacker
	Byte(title string, singleByteValue byte) MetaAttributesStacker
	Bytes(title string, values []byte) MetaAttributesStacker
	Hex(title string, hexValues []byte) MetaAttributesStacker
	RawJson(title string, rawJsonBytes []byte) MetaAttributesStacker
	Error(title string, err error) MetaAttributesStacker
	AnErr(key string, err error) MetaAttributesStacker
	MapAny(title string, mapAny map[string]interface{}) MetaAttributesStacker
	MapIntegerAny(title string, mapAny map[int]interface{}) MetaAttributesStacker
	Meta(title string, metaAttr MetaAttributesCompiler) MetaAttributesStacker

	Json(title string, json *corejson.Result) MetaAttributesStacker
	JsonItems(title string, jsons ...*corejson.Result) MetaAttributesStacker

	Err(err error) MetaAttributesStacker

	DefaultStackTraces() MetaAttributesStacker
	ErrWithTypeTraces(title string, errType errcoreinf.BasicErrorTyper, err error) MetaAttributesStacker
	ErrorsWithTypeTraces(title string, errType errcoreinf.BasicErrorTyper, errorItems ...error) MetaAttributesStacker
	StackTraces(stackSkipIndex int, title string) MetaAttributesStacker
	OnErrStackTraces(err error) MetaAttributesStacker
	OnErrWrapperOrCollectionStackTraces(errWrapperOrCollection errcoreinf.BaseErrorOrCollectionWrapper) MetaAttributesStacker

	FullStringer(
		fullStringer errcoreinf.FullStringer,
	) MetaAttributesStacker

	FullStringerTitle(
		title string,
		fullStringer errcoreinf.FullStringer,
	) MetaAttributesStacker
	FullTraceAsAttr(
		title string,
		attrFullStringWithTraces errcoreinf.FullStringWithTracesGetter,
	) MetaAttributesStacker

	BasicErrWrapper(errWrapperOrCollection errcoreinf.BasicErrWrapper) MetaAttributesStacker
	BaseRawErrCollectionDefiner(errWrapperOrCollection errcoreinf.BaseRawErrCollectionDefiner) MetaAttributesStacker
	BaseErrorWrapperCollectionDefiner(errWrapperOrCollection errcoreinf.BaseErrorWrapperCollectionDefiner) MetaAttributesStacker
	ErrWrapperOrCollection(errWrapperOrCollection errcoreinf.BaseErrorOrCollectionWrapper) MetaAttributesStacker
	RawErrCollection(key string, err errcoreinf.BaseRawErrCollectionDefiner) MetaAttributesStacker
	CompiledBasicErrWrapper(compiler errcoreinf.CompiledBasicErrWrapper) MetaAttributesStacker

	Namer(namer enuminf.Namer) MetaAttributesStacker
	NamerTitle(title string, namer enuminf.Namer) MetaAttributesStacker

	EnumTitleEnum(title enuminf.SimpleEnumer, enum enuminf.BasicEnumer) MetaAttributesStacker
	SimpleEnumTitleEnum(title enuminf.SimpleEnumer, enum enuminf.SimpleEnumer) MetaAttributesStacker
	Enum(title string, enum enuminf.BasicEnumer) MetaAttributesStacker
	Enums(key string, enums ...enuminf.BasicEnumer) MetaAttributesStacker
	OnlyEnum(enum enuminf.BasicEnumer) MetaAttributesStacker
	OnlyEnums(enums ...enuminf.BasicEnumer) MetaAttributesStacker
	OnlyString(value string) MetaAttributesStacker
	OnlyStrings(values ...string) MetaAttributesStacker
	OnlyIntegers(values ...int) MetaAttributesStacker
	OnlyBooleans(values ...bool) MetaAttributesStacker
	OnlyBytes(rawBytes []byte) MetaAttributesStacker
	OnlyRawJson(rawBytes []byte) MetaAttributesStacker
	OnlyBytesErr(rawBytes []byte, err error) MetaAttributesStacker
	OnlyAnyItems(values ...interface{}) MetaAttributesStacker
	OnlyMetaAttr(metaAttr MetaAttributesCompiler) MetaAttributesStacker

	OnlyJson(json *corejson.Result) MetaAttributesStacker
	OnlyJsons(jsons ...*corejson.Result) MetaAttributesStacker

	Bool(title string, isResult bool) MetaAttributesStacker
	Booleans(title string, isResults ...bool) MetaAttributesStacker

	Any(anyItem interface{}) MetaAttributesStacker
	AnyIf(isLog bool, anyItem interface{}) MetaAttributesStacker
	AnyItems(anyItems ...interface{}) MetaAttributesStacker
	AnyItemsIf(isLog bool, anyItems ...interface{}) MetaAttributesStacker

	Jsoner(jsoner corejson.Jsoner) MetaAttributesStacker
	Jsoners(jsoners ...corejson.Jsoner) MetaAttributesStacker
	JsonerTitle(title string, jsoner corejson.Jsoner) MetaAttributesStacker
	JsonerIf(isLog bool, jsoner corejson.Jsoner) MetaAttributesStacker
	JsonersIf(isLog bool, jsoners ...corejson.Jsoner) MetaAttributesStacker

	Serializer(serializer Serializer) MetaAttributesStacker
	Serializers(serializers ...Serializer) MetaAttributesStacker
	SerializerFunc(serializerFunc func() ([]byte, error)) MetaAttributesStacker
	SerializerFunctions(serializerFunctions ...func() ([]byte, error)) MetaAttributesStacker

	StandardTaskEntityDefiner(entity entityinf.StandardTaskEntityDefiner) MetaAttributesStacker
	TaskEntityDefiner(entity entityinf.TaskEntityDefiner) MetaAttributesStacker

	StandardTaskEntityDefinerTitle(title string, entity entityinf.StandardTaskEntityDefiner) MetaAttributesStacker
	TaskEntityDefinerTitle(title string, entity entityinf.TaskEntityDefiner) MetaAttributesStacker

	LoggerModel(loggerModel SingleLogModeler) MetaAttributesStacker
	LoggerModelTitle(title string, loggerModel SingleLogModeler) MetaAttributesStacker

	Int(key string, i int) MetaAttributesStacker
	Integers(key string, integerItems ...int) MetaAttributesStacker
	Fmt(format string, v ...interface{}) MetaAttributesStacker
	FmtIf(isLog bool, format string, v ...interface{}) MetaAttributesStacker

	RawPayloadsGetter(payloadsGetter RawPayloadsGetter) MetaAttributesStacker
	RawPayloadsGetterTitle(title string, payloadsGetter RawPayloadsGetter) MetaAttributesStacker
	RawPayloadsGetterIf(isLog bool, payloadsGetter RawPayloadsGetter) MetaAttributesStacker

	Inject(others ...MetaAttributesStacker) MetaAttributesStacker
	ConcatNew(others ...MetaAttributesStacker) MetaAttributesStacker
	coreinterface.Clearer

	MetaAttributesCompiler
}
