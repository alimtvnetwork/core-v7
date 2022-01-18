package coredynamictestwrappers

import (
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/core/corevalidator"
)

type ReflectSetFromToTestWrapper struct {
	Header                          string
	From, To, ExpectedValue, actual interface{}
	IsUsePointerInFrom              bool
	IsErrorExpected                 bool
	Validator                       corevalidator.TextValidator
}

func (it ReflectSetFromToTestWrapper) CaseTitle() string {
	return it.Header
}

func (it ReflectSetFromToTestWrapper) Input() interface{} {
	return it
}

func (it ReflectSetFromToTestWrapper) Expected() interface{} {
	return it.ExpectedValue
}

func (it ReflectSetFromToTestWrapper) ToFieldToDraftType() *coretests.DraftType {
	return coretests.AnyToDraftType(it.To)
}

func (it ReflectSetFromToTestWrapper) ToFieldToBytes() []byte {
	return coretests.AnyToBytes(it.To)
}

func (it ReflectSetFromToTestWrapper) ExpectedFieldToDraftType() *coretests.DraftType {
	return coretests.AnyToDraftType(it.ExpectedValue)
}

func (it ReflectSetFromToTestWrapper) ExpectedFieldToBytes() []byte {
	return coretests.AnyToBytes(it.ExpectedValue)
}

func (it ReflectSetFromToTestWrapper) SetActual(actual interface{}) {
	it.actual = actual
}

func (it ReflectSetFromToTestWrapper) Actual() interface{} {
	return it.actual
}

func (it ReflectSetFromToTestWrapper) AsSimpleTestCaseWrapper() coretests.SimpleTestCaseWrapper {
	return it
}

func (it *ReflectSetFromToTestWrapper) AsSimpleTestCaseWrapperContractsBinder() coretests.SimpleTestCaseWrapperContractsBinder {
	return it
}
