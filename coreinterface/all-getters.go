package coreinterface

import (
	"reflect"

	"gitlab.com/auk-go/core/internal/internalinterface"
)

type StringRangesGetter interface {
	StringRangesPtr() *[]string
	StringRanges() []string
}

type ReflectTypeGetter interface {
	ReflectType() reflect.Type
	GetErrorOnTypeMismatch(
		typeMatch reflect.Type,
		isIncludeInvalidMessage bool,
	) error
}

type ErrorGetter interface {
	Error() error
}

type MetaGetter interface {
	Meta() string
}

type MessageGetter interface {
	Message() string
}

type MeaningFullErrorGetter interface {
	MeaningFullError() error
}

type RawPayloadsGetter interface {
	internalinterface.RawPayloadsGetter
}

type ListStringsGetter interface {
	ListStrings() []string
}

type LengthGetter interface {
	Length() int
}

type EntityTypeNameGetter interface {
	EntityTypeName() string
}

type TypeNameWithRangeNamesCsvGetter interface {
	StringRangeNamesCsvGetter
	TypeNameGetter
}

type CountGetter interface {
	Count() int
}

type StringsGetter interface {
	Strings() []string
}

type IdentifierGetter interface {
	Identifier() string
}

type InvalidMessageGetter interface {
	// InvalidMessage get invalid message
	InvalidMessage() string
}

type TypeNameGetter interface {
	TypeName() string
}

type TypenameStringGetter interface {
	TypenameString() string
}

type StringsLinesGetter interface {
	Lines() []string
}

type NameGetter interface {
	Name() string
}

type CategoryNameGetter interface {
	CategoryName() string
}

type AnyAttributesGetter interface {
	internalinterface.AnyAttributesGetter
}

// DynamicPageItemsGetter Paging items related methods
type DynamicPageItemsGetter interface {
	// GetDynamicPagedItems returns items for the paged collection.
	GetDynamicPagedItems(perPageItems int, pageIndex int) interface{}
}

type IntRangeLengthGetter interface {
	RangeLength() int
}

type IntRangesGetter interface {
	RangesInt() []int
	Ranges() []int
}

type InvalidDirectErrorGetter interface {
	GetInvalidDirectError() error
}

type InvalidErrorGetter interface {
	// InvalidError get invalid message error
	InvalidError() error
}

type LastIndexGetter interface {
	LastIndex() int
}

type RawMessageBytesGetter interface {
	RawMessageBytes() []byte
}

type ReflectKindGetter interface {
	ReflectKind() reflect.Kind
}

type SafeBytesGetter interface {
	SafeBytes() []byte
}

type SafeBytesPointerGetter interface {
	SafeBytesPtr() *[]byte
}

type StringRangeNamesCsvGetter interface {
	RangeNamesCsv() string
}

type ReflectValueGetter interface {
	ReflectValue() *reflect.Value
}

type SafeStringsGetter interface {
	SafeStrings() []string
}

type StringRangeNamesGetter interface {
	StringRangeNames() []string
}

type UintUserIdGetter interface {
	UserId() uint
}

type UserIdGetter interface {
	UserId() int
}

type UsernameGetter interface {
	Username() string
}

type ValidationMessageGetter interface {
	ValidationMessage() string
}

// ValidationErrorGetter
//
//	Returns validation related error
//	only
type ValidationErrorGetter interface {
	// ValidationError
	//
	//  Returns validation related error
	//  only
	ValidationError() error
}

type ValueAnyItemGetter interface {
	Value() interface{}
}

type ValueByteGetter interface {
	Value() byte
}

type ValueErrorGetter interface {
	Value() error
}

type ValueFloat32Getter interface {
	Value() float32
}

type ValueFloat64Getter interface {
	Value() float64
}

type ValueGetter interface {
	ValueDynamic() interface{}
}

type ValueInt16Getter interface {
	Value() int16
}

type ValueInt32Getter interface {
	Value() int32
}

type ValueInt64Getter interface {
	Value() int64
}

type ValueIntegerGetter interface {
	Value() int
}

type ValueReflectSetter interface {
	ValueReflectSet(setterPtr interface{}) error
}

type ValueStringGetter interface {
	Value() string
}

type ValueStringsGetter interface {
	Value() []string
}

type FilterTextGetter interface {
	FilterText() string
}

type CategoryRevealerGetter interface {
	CategoryRevealer() CategoryRevealer
}

type StackTracesBytesGetter interface {
	StackTracesBytes() []byte
}

type JsonErrorBytesGetter interface {
	JsonErrorBytes() []byte
}

type MapStringAnyGetter interface {
	MapStringAny() map[string]interface{}
}

type MapStringStringGetter interface {
	MapStringString() map[string]string
}

type MapStringFuncGetter interface {
	MapStringString() map[string]func()
}

type MapStringErrorFuncGetter interface {
	MapStringString() map[string]func() error
}

type VariableNameStringGetter interface {
	VariableName() string
}

type VariableNameAnyGetter interface {
	VariableNameAny() interface{}
}

type ValueAnyGetter interface {
	ValueAny() interface{}
}

type ExplicitValueStringGetter interface {
	ValueString() string
}

type KeysHashsetGetter interface {
	KeysHashset() map[string]bool
}

type HashmapGetter interface {
	Hashmap() map[string]string
}

type AllValuesStringsGetter interface {
	AllValuesStrings() []string
}

type AllValuesIntegersGetter interface {
	AllValuesIntegers() []int
}

type KeyStringNameGetter interface {
	KeyName() string
}
