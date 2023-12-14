package codegen

import (
	"fmt"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/convertinteranl"
	"gitlab.com/auk-go/core/iserror"
)

type CodeOutput struct {
	UnitTest             string
	TestCase             string
	StructName, FuncName string
	Error                error
	FileWriter           *chmodhelper.SimpleFileReaderWriter
}

func (it *CodeOutput) IsValid() bool {
	return it != nil &&
		it.Error == nil &&
		it.UnitTest != "" ||
		it.TestCase != ""
}

func (it *CodeOutput) IsInvalid() bool {
	return !it.IsValid()
}

func (it *CodeOutput) HasError() bool {
	return it != nil && it.Error != nil
}

func (it *CodeOutput) IsEmptyError() bool {
	return it == nil && it.Error == nil
}

func (it *CodeOutput) ErrorString() string {
	if it.IsEmptyError() {
		return ""
	}

	return it.Error.Error()
}

func (it *CodeOutput) FormatCode(code string) (string, error) {
	if it.HasError() {
		return code, it.Error
	}

	s, err := convertinteranl.CodeFormatter.Golang(code)

	return s, errcore.StackEnhance.Error(err)
}

func (it *CodeOutput) FormatUnitTest() (string, error) {
	if it.HasUnitTest() {
		return it.FormatCode(it.UnitTest)
	}

	return "", it.Error
}

func (it *CodeOutput) FormatTestCase() (string, error) {
	if it.HasTestCase() {
		return it.FormatCode(it.TestCase)
	}

	return "", it.Error
}

func (it *CodeOutput) HasUnitTest() bool {
	return it != nil &&
		it.UnitTest != ""
}

func (it *CodeOutput) HasTestCase() bool {
	return it != nil &&
		it.TestCase != ""
}

func (it *CodeOutput) Write() errcore.RawErrCollection {
	var rawErrCollection errcore.RawErrCollection

	if it == nil {
		rawErrCollection.AddMsg("code output is nil")

		return rawErrCollection
	}

	if it.IsInvalid() {
		rawErrCollection.Add(it.Error)

		return rawErrCollection
	}

	if it.HasUnitTest() {
		rawErrCollection.Add(it.WriteUnitTestFile())
	}

	if it.HasTestCase() {
		rawErrCollection.Add(it.WriteTestCaseFile())
	}

	return rawErrCollection
}

func (it *CodeOutput) WriteUnitTestFile() error {
	filePath := it.unitTestFileName()
	code, err := it.FormatUnitTest()

	if iserror.Defined(err) {
		return err
	}

	return it.FileWriter.WriteRelativePath(
		it.FileWriter.IsRemoveBeforeWrite,
		filePath,
		[]byte(code),
	)
}

func (it *CodeOutput) WriteTestCaseFile() error {
	filePath := it.testCaseFileName()
	code, err := it.FormatTestCase()

	if iserror.Defined(err) {
		return err
	}

	return it.FileWriter.WriteRelativePath(
		it.FileWriter.IsRemoveBeforeWrite,
		filePath,
		[]byte(code),
	)
}

func (it *CodeOutput) unitTestFileName() string {
	if len(it.StructName) == 0 {
		return fmt.Sprintf(
			"%s_test.go",
			it.FuncName,
		)
	}

	return fmt.Sprintf(
		"%s_%s_test.go",
		it.StructName,
		it.FuncName,
	)
}

func (it *CodeOutput) testCaseFileName() string {
	if len(it.StructName) == 0 {
		return fmt.Sprintf(
			"%s_testcases.go",
			it.FuncName,
		)
	}

	return fmt.Sprintf(
		"%s_%s_testcases.go",
		it.StructName,
		it.FuncName,
	)
}
