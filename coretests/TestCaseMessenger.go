package coretests

type TestCaseMessenger interface {
	GetFuncName() string
	Value() interface{}
	Expected() interface{}
	Actual() interface{}
}
