package codegen

const (
	testPkgHeaderTemplate = `
package $packageName

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"

	$newPackages
)
`

	funcTemplate = `
func Test_$FuncName_$Behaviour(t *testing.T) {
	for caseIndex, testCase := range $testCaseName {
		// Arrange
		input := testCase.
			ArrangeInput.($ArrangeType)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap($linesPossible)
		$variablesSetup

		// Act
		actFunc$FuncName := $directFuncInvoke
		$outArgs := actFunc$FuncName($inArgs)

		actualSlice.AppendFmt(
			"$fmtJoin",
			caseIndex,
			$fmtOutputs,
		)

		finalActLines := actualSlice.Strings()
		actualSlice.Dispose()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
`

	fullTestCaseTemplate = `
	$testCaseName = []coretestcases.CaseV1{
		$caseItem
	}
`

	fullTestCaseFileTemplate = `
var (
	$testCases
)
`

	testCaseItemTemplate = `
		{
			Title: $title,
			ArrangeInput: $ArrangeType {
				$arrangeSetup,
			},
			ExpectedInput: []string{
				$expectedLines,
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf($VerifyTypeOf),
		},
`

	argSingleTemplate = "%s : %s"
)
