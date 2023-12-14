package coretests

type SimpleTestCaseWrapper interface {
	CaseTitle() string
	Input() interface{}
	Expected() interface{}
	Actual() interface{}
	SetActual(actual interface{})
}
