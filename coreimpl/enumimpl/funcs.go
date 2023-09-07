package enumimpl

type (
	LooperFunc        func(index int, name string, anyVal interface{}) (isBreak bool)
	LooperIntegerFunc func(index int, name string, anyVal int) (isBreak bool)
)
