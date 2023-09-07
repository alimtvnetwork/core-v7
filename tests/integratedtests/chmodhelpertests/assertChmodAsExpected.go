package chmodhelpertests

import (
	"testing"

	"gitlab.com/auk-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func assertTestCaseChmodAsExpected(
	t *testing.T,
	testCase *chmodhelpertestwrappers.RwxInstructionTestWrapper,
	testHeader string,
) {
	expected := testCase.ExpectedAsRwxOwnerGroupOtherInstruction()
	expectedChmod := expected.String()

	for _, createPath := range testCase.CreatePaths {
		assertSingleChmod(
			t,
			testHeader,
			createPath,
			expectedChmod)
	}
}
