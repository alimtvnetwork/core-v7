package enuminf

// BasicEnumer
//
// EnumFormatter:
//
//  Outputs name and
//  value by given format.
//
// sample-format :
//  - "Enum of {type-name} - {name} - {value}"
//
// sample-format-output :
//  - "Enum of EnumFullName - Invalid - 0"
//
// Key-Meaning :
//  - {type-name} : represents type-name string
//  - {name}      : represents name string
//  - {value}     : represents value string
type BasicEnumer interface {
	BaseEnumer
	EnumFormatter
	MaxMaxAny() (min, max interface{})
	MinValueString() string
	MaxValueString() string
	MaxInt() int
	MinInt() int
	RangesDynamicMapGetter
	AllNameValues() []string
	OnlySupportedErr(names ...string) error
	OnlySupportedMsgErr(message string, names ...string) error
	IntegerEnumRangesGetter
	EnumType() EnumTyper
}
