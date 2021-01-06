package coretests

type TestFuncName string

func (funcName TestFuncName) Value() string {
	return string(funcName)
}
