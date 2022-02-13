package internalenuminf

type ByteValuePlusEqualer interface {
	Value() byte
	IsValueEqual(value byte) bool
}
