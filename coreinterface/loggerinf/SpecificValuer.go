package loggerinf

type SpecificValuer interface {
	IsAnyValueDefined() bool
	BytesVal() []byte
	StringVal() string
	BooleanVal() bool
	IntegerVal() int

	ByteVal() byte
	HasAnyKeyValuePair() bool
	HasAnyKeyValues() bool

	KeyValuesLength() int
	AnyKeyValuesLength() int

	IsEmptyKeyValuePairs() bool
	IsEmptyAnyKeyValuePairs() bool

	AnyValuesMap() map[string]interface{}
	KeyValuesPairsMap() map[string]string

	GetAnyValue(key string) (val interface{}, isFound bool)
	GetKeyValue(key string) (val string, isFound bool)

	AnyValueReflectSetTo(toPtr interface{}) error
}
