package internalenuminf

type CompareMethodsTyper interface {
	BasicEnumer
	ByteValuePlusEqualer

	IsEqual() bool
	IsLess() bool
	IsLessEqual() bool
	IsGreater() bool
	IsGreaterEqual() bool
	IsNotEqual() bool
}
