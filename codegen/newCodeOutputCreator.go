package codegen

import (
	"errors"

	"gitlab.com/auk-go/core/chmodhelper"
)

type newCodeOutputCreator struct{}

func (it newCodeOutputCreator) Default(
	structName, funcName,
	unit, testCase string,
	fileWriter *chmodhelper.SimpleFileReaderWriter,
) *CodeOutput {
	return &CodeOutput{
		UnitTest:   unit,
		TestCase:   testCase,
		StructName: structName,
		FuncName:   funcName,
		FileWriter: fileWriter,
	}
}

func (it newCodeOutputCreator) All(
	structName, funcName,
	unit, testCase string,
	fileWriter *chmodhelper.SimpleFileReaderWriter,
) *CodeOutput {
	return &CodeOutput{
		UnitTest:   unit,
		TestCase:   testCase,
		StructName: structName,
		FuncName:   funcName,
		FileWriter: fileWriter,
	}
}

func (it newCodeOutputCreator) Invalid(
	err error,
) *CodeOutput {
	return &CodeOutput{
		Error: err,
	}
}

func (it newCodeOutputCreator) InvalidMsg(
	msg string,
) *CodeOutput {
	return &CodeOutput{
		Error: errors.New(msg),
	}
}
