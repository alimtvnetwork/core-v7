package enumimpl

type BasicByter interface {
	IsAnyOf(
		value byte,
		givenBytes ...byte,
	) bool
	Max() byte
	Min() byte
	GetValueByString(
		jsonValueString string,
	) byte
	GetValueByName(
		name string,
	) (byte, error)
	GetStringValue(
		input byte,
	) string
	Ranges() []byte
	Hashmap() map[string]byte
	HashmapPtr() *map[string]byte
	IsValidRange(
		value byte,
	) bool
	ToEnumJsonBytes(
		value byte,
	) ([]byte, error)
	ToEnumString(
		value byte,
	) string
	AppendPrependJoinValue(
		joiner string,
		appendVal, prependVal byte,
	) string
	AppendPrependJoinNamer(
		joiner string,
		appendVal, prependVal toNamer,
	) string
	ToNumberString(
		valueInRawFormat interface{},
	) string
	// UnmarshallToValue
	//
	//  isMappedToFirstIfEmpty: maps invalid values to first item
	UnmarshallToValue(
		isMappedToFirstIfEmpty bool,
		jsonUnmarshallingValue []byte,
	) (byte, error)
}
