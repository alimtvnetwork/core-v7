package serializerinf

import "bytes"

type SafeBytesTypeNameGetter interface {
	SafeBytesTypeName() string
}

type BytesTypeNameGetter interface {
	BytesTypeName() string
}

type SafeStringGetter interface {
	SafeString() string
}

type JsonStringGetter interface {
	JsonString() string
}

type JsonStringPtrGetter interface {
	JsonStringPtr() *string
}

type PrettyJsonBufferGetter interface {
	PrettyJsonBuffer(prefix, indent string) (*bytes.Buffer, error)
}

type PrettyJsonStringGetter interface {
	PrettyJsonString() string
}

type PrettyJsonStringOrErrStringGetter interface {
	PrettyJsonStringOrErrString() string
}

type LengthGetter interface {
	Length() int
}

type HasErrorChecker interface {
	HasError() bool
}

type ErrorStringGetter interface {
	ErrorString() string
}

type IsErrorEqualChecker interface {
	IsErrorEqual(err error) bool
}

type SafeBytesGetter interface {
	SafeBytes() []byte
}

type BytesValuesGetter interface {
	Values() []byte
}

type SafeValuesGetter interface {
	SafeValues() []byte
}

type SafeValuesPtrGetter interface {
	SafeValuesPtr() *[]byte
}

type RawSerializeGetter interface {
	Raw() ([]byte, error)
}

type MustRawSerializeGetter interface {
	RawMust() []byte
}

type RawStringSerializeGetter interface {
	RawString() (jsonString string, err error)
}

type MustRawStringSerializeGetter interface {
	RawStringMust() (jsonString string)
}

type RawErrStringGetter interface {
	RawErrString() (rawJsonBytes []byte, errorMsg string)
}

type RawPrettyStringGetter interface {
	RawPrettyString() (jsonString string, err error)
}

type MeaningfulErrorMessageGetter interface {
	MeaningfulErrorMessage() string
}

type MeaningfulErrorGetter interface {
	MeaningfulError() error
}

type IsEmptyErrorChecker interface {
	IsEmptyError() bool
}

type HasSafeItemsChecker interface {
	HasSafeItems() bool
}

type IsAnyNullChecker interface {
	IsAnyNull() bool
}

type HasIssuesOrEmptyChecker interface {
	HasIssuesOrEmpty() bool
}

type ErrorHandler interface {
	HandleError()
}

type MustBeSafer interface {
	MustBeSafe()
}

type ErrorHandlerWithMessager interface {
	HandleErrorWithMsg(msg string)
}

type HasBytesChecker interface {
	HasBytes() bool
}

type HasJsonBytesChecker interface {
	HasJsonBytes() bool
}

type IsEmptyJsonBytesChecker interface {
	IsEmptyJsonBytes() bool
}

type IsEmptyChecker interface {
	IsEmpty() bool
}

type HasAnyItemChecker interface {
	HasAnyItem() bool
}

type IsEmptyJsonChecker interface {
	IsEmptyJson() bool
}

type Deserializer interface {
	Deserialize(
		anyPointer interface{},
	) error
}

type MustDeserializer interface {
	DeserializeMust(
		anyPointer interface{},
	)
}

type MustUnmarshaler interface {
	UnmarshalMust(
		anyPointer interface{},
	)
}

type Unmarshaler interface {
	Unmarshal(
		anyPointer interface{},
	) error
}

type SkipExistingIssuesSerializer interface {
	SerializeSkipExistingIssues() (
		[]byte, error,
	)
}

type SelfSerializer interface {
	Serialize() ([]byte, error)
}

type MustSelfSerializer interface {
	SerializeMust() []byte
}
