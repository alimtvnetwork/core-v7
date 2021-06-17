package chmodhelpertests

import (
	"fmt"
	"testing"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_RwxWrapperManyApplyValue(t *testing.T) {
	// Arrange
	createPathInstructions := chmodhelpertestwrappers.CreatePathInstruction2
	createDefaultPaths(&createPathInstructions)
	firstCreationIns := createPathInstructions[0]
	paths := firstCreationIns.GetPaths()
	condition := chmodins.DefaultAllTrueCondition()
	existingAppliedRwxFull := firstCreationIns.ApplyRwx.String()
	for _, testCase := range chmodhelpertestwrappers.SingleRwxTestCases {
		rwxWrapper, err := testCase.ToDisabledRwxWrapper()
		msgtype.SimpleHandleErr(err, "SingleRwx ToDisabledRwxWrapper failed")
		expectation := rwxWrapper.ToFullRwxValueString()

		header := fmt.Sprintf(
			"Existing [%s] Applied by [%s] should result [%s]",
			existingAppliedRwxFull,
			expectation,
			expectation)

		// Act
		err2 := rwxWrapper.ApplyLinuxChmodOnMany(condition, paths...)

		msgtype.SimpleHandleErr(err2, "rwxWrapper.ApplyLinuxChmodOnMany failed")

		// for directory `-` will be placed not `d`
		assertSingleChmod(
			t,
			header,
			firstCreationIns,
			expectation)
	}
}
