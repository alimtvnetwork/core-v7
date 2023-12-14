package coredynamictestwrappers

import (
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/corevalidator"
)

type FromToTestWrapper struct {
	Header                          string
	From, To, ExpectedValue, actual interface{}
	IsUsePointerInFrom              bool
	IsExpectingError                bool
	HasPanic                        bool
	Validator                       corevalidator.TextValidator
}

func (it FromToTestWrapper) CaseTitle() string {
	return it.Header
}

func (it FromToTestWrapper) Input() interface{} {
	return it.From
}

func (it FromToTestWrapper) Expected() interface{} {
	return it.ExpectedValue
}

func (it FromToTestWrapper) ToFieldToDraftType() *coretests.DraftType {
	return coretests.AnyToDraftType(it.To)
}

func (it FromToTestWrapper) ToFieldToBytes() []byte {
	return coretests.AnyToBytes(it.To)
}

func (it FromToTestWrapper) ExpectedFieldToDraftType() *coretests.DraftType {
	return coretests.AnyToDraftType(it.ExpectedValue)
}

func (it FromToTestWrapper) ExpectedFieldToBytes() []byte {
	return coretests.AnyToBytes(it.ExpectedValue)
}

func (it FromToTestWrapper) SetActual(actual interface{}) {
	it.actual = actual
}

func (it FromToTestWrapper) Actual() interface{} {
	return it.actual
}

func (it FromToTestWrapper) AsSimpleTestCaseWrapper() coretests.SimpleTestCaseWrapper {
	return &it
}

func (it *FromToTestWrapper) AsSimpleTestCaseWrapperContractsBinder() coretests.SimpleTestCaseWrapperContractsBinder {
	return it
}
