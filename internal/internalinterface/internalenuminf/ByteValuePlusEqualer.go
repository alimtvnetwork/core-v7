package internalenuminf

type ByteValuePlusEqualer interface {
	Value() byte
	IsByteValueEqual(value byte) bool
}
