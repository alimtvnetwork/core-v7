package enuminf

type IsAnyValueByteEqualer interface {
	IsAnyValuesEqual(anyByteValues ...byte) bool
}

type IsAnyValueIntegerEqualer interface {
	IsAnyValuesEqual(anyValues ...int) bool
}

type IsAnyValueInteger8Equaler interface {
	IsAnyValuesEqual(anyValues ...int8) bool
}

type IsAnyValueInteger16Equaler interface {
	IsAnyValuesEqual(anyValues ...int16) bool
}

type IsAnyValueInteger32Equaler interface {
	IsAnyValuesEqual(anyValues ...int32) bool
}

type IsValueByteEqualer interface {
	IsByteValueEqual(value byte) bool
}

type IsValueIntegerEqualer interface {
	IsIntegerValueEqual(value int) bool
}

type IsValueInteger8Equaler interface {
	IsInteger8ValueEqual(value int8) bool
}

type IsValueInteger16Equaler interface {
	IsInteger16ValueEqual(value int16) bool
}

type IsValueInteger32Equaler interface {
	IsInteger32ValueEqual(value int32) bool
}

type IsEnumEqualer interface {
	IsEnumEqual(enum BasicEnumer) bool
}

type IsAnyEnumsEqualer interface {
	IsAnyEnumsEqual(enums ...BasicEnumer) bool
}

type IsBothEnumEqualer interface {
	IsEnumEqualer
	IsAnyEnumsEqualer
}
