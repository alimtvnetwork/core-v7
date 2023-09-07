package chmodhelpertestwrappers

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/coretests"
)

type RwxInstructionTestWrapper struct {
	RwxInstructions []chmodins.RwxInstruction
	DefaultRwx      *chmodins.RwxOwnerGroupOther
	IsErrorExpected bool
	CreatePaths     []*chmodhelper.DirFilesWithRwxPermission
	funcName        coretests.TestFuncName
	expected        chmodins.RwxOwnerGroupOther
	actual          interface{}
}

func (receiver *RwxInstructionTestWrapper) Actual() interface{} {
	return receiver.actual
}

func (receiver *RwxInstructionTestWrapper) SetActual(actual interface{}) {
	receiver.actual = actual
}

func (receiver *RwxInstructionTestWrapper) FuncName() string {
	return receiver.funcName.Value()
}

func (receiver *RwxInstructionTestWrapper) Value() interface{} {
	return receiver
}

func (receiver *RwxInstructionTestWrapper) Expected() interface{} {
	return receiver.expected
}

func (receiver *RwxInstructionTestWrapper) ExpectedAsRwxOwnerGroupOtherInstruction() chmodins.RwxOwnerGroupOther {
	return receiver.expected
}

func (receiver *RwxInstructionTestWrapper) AsTestCaseMessenger() coretests.TestCaseMessenger {
	var testCaseMessenger coretests.TestCaseMessenger = receiver

	return testCaseMessenger
}
