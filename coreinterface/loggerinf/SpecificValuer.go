package loggerinf

type SpecificValuer interface {
	IsAnyValueDefined() bool
	StringVal() string
	BooleanVal() bool
	IntegerVal() bool
	ByteVal() bool
}
