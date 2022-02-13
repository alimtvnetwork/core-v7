package internalenuminf

type CompareMethodsTyper interface {
	IsEqual() bool
	IsLess() bool
	IsLessEqual() bool
	IsGreater() bool
	IsGreaterEqual() bool
	IsNotEqual() bool
	BasicEnumer
	ByteValuePlusEqualer
}
