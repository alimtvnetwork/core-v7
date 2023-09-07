package serializerinf

type JsonResulter interface {
	BasicJsonResulter

	SafeBytesTypeNameGetter
	BytesTypeNameGetter

	SafeStringGetter
	JsonStringGetter
	JsonStringPtrGetter
	PrettyJsonBufferGetter
	PrettyJsonStringGetter
	PrettyJsonStringOrErrStringGetter

	LengthGetter
	HasErrorChecker
	ErrorStringGetter
	IsErrorEqualChecker
	SafeBytesGetter
	BytesValuesGetter
	SafeValuesGetter
	SafeValuesPtrGetter
	RawSerializeGetter
	MustRawSerializeGetter
	RawStringSerializeGetter
	MustRawStringSerializeGetter
	RawErrStringGetter
	RawPrettyStringGetter
	MeaningfulErrorMessageGetter
	MeaningfulErrorGetter
	IsEmptyErrorChecker
	HasSafeItemsChecker
	IsAnyNullChecker
	HasIssuesOrEmptyChecker
	ErrorHandler
	MustBeSafer
	ErrorHandlerWithMessager
	HasBytesChecker
	HasJsonBytesChecker
	IsEmptyJsonBytesChecker
	IsEmptyChecker
	IsEmptyJsonChecker

	Deserializer
	MustDeserializer

	MustUnmarshaler
	Unmarshaler

	SkipExistingIssuesSerializer

	SelfSerializer
	MustSelfSerializer

	CombineErrorWithRefString(references ...string) string
	CombineErrorWithRefError(references ...string) error
	Dispose()
}

type SimpleBytesResulter interface {
	LengthGetter

	BytesValuesGetter
	SafeValuesGetter
	MeaningfulErrorGetter

	RawSerializeGetter

	IsEmptyChecker
	HasAnyItemChecker

	HasErrorChecker
}

type BaseJsonResulter interface {
	SimpleBytesResulter

	BytesTypeNameGetter

	SafeStringGetter

	LengthGetter
	HasErrorChecker

	MeaningfulErrorGetter

	BytesValuesGetter
	SafeValuesGetter

	RawSerializeGetter

	IsEmptyChecker

	SelfSerializer
}

type BasicJsonResulter interface {
	BaseJsonResulter

	SafeBytesTypeNameGetter
	BytesTypeNameGetter

	SafeStringGetter
	JsonStringGetter

	LengthGetter
	HasErrorChecker
	ErrorStringGetter
	SafeBytesGetter
	BytesValuesGetter
	SafeValuesGetter

	RawSerializeGetter
	MustRawSerializeGetter

	IsEmptyErrorChecker
	HasSafeItemsChecker
	HasIssuesOrEmptyChecker
	ErrorHandler

	Deserializer
	MustDeserializer

	Unmarshaler

	SkipExistingIssuesSerializer

	SelfSerializer
	MustSelfSerializer
}
