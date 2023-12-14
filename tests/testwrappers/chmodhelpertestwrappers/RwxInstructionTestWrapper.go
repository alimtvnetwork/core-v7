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
	CreatePaths     []chmodhelper.DirFilesWithRwxPermission
	TestFuncName    coretests.TestFuncName
	WhatIsExpected  chmodins.RwxOwnerGroupOther
	actual          interface{}
}

func (it *RwxInstructionTestWrapper) Actual() interface{} {
	return it.actual
}

func (it *RwxInstructionTestWrapper) SetActual(actual interface{}) {
	it.actual = actual
}

func (it *RwxInstructionTestWrapper) GetFuncName() string {
	return it.TestFuncName.Value()
}

func (it *RwxInstructionTestWrapper) Value() interface{} {
	return it
}

func (it *RwxInstructionTestWrapper) Expected() interface{} {
	return it.WhatIsExpected
}

func (it *RwxInstructionTestWrapper) ExpectedAsRwxOwnerGroupOtherInstruction() chmodins.RwxOwnerGroupOther {
	return it.WhatIsExpected
}

func (it *RwxInstructionTestWrapper) AsTestCaseMessenger() coretests.TestCaseMessenger {
	var testCaseMessenger coretests.TestCaseMessenger = it

	return testCaseMessenger
}
