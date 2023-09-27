package enumimpl

type (
	LooperFunc                func(index int, name string, anyVal interface{}) (isBreak bool)
	LooperIntegerFunc         func(index int, name string, anyVal int) (isBreak bool)
	IsEqualCheckerFunc        func(isRegardless bool, l, r interface{}) bool
	GetSingleDiffResultFunc   func(isLeft bool, l, r interface{}) interface{}
	GetOnKeyMissingResultFunc func(lKey string, lVal interface{}) interface{}
)
