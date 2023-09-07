package internalenuminf

type InternalByteEnumer interface {
	AllNameValues() []string
	IntegerEnumRanges() []int
	RangesDynamicMap() map[string]interface{}
	Format(format string) (compiled string)
	IsNameEqual(name string) bool
	IsByteValueEqual(value byte) bool

	Name() string
	NameValue() string
	TypeName() string
	ToNumberString() string
	String() string

	ValueInt() int
	ValueInt8() int8
	ValueInt16() int16
	ValueInt32() int32
	ValueString() string

	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
	RangeNamesCsv() string
	RangesByte() []byte
	ValueByte() byte
	MaxByte() byte
	MinByte() byte

	IsAnyNamesOf(names ...string) bool
	IsValueEqual(value byte) bool
	IsAnyValuesEqual(anyByteValues ...byte) bool
}
