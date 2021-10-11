package chmodhelpertests

import (
	"fmt"
	"testing"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_RwxWrapperManyApplyValue_Unix(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	createPathInstructions := chmodhelpertestwrappers.CreatePathInstruction2
	chmodhelper.CreateDirFilesWithRwxPermissionsMust(
		true,
		&createPathInstructions)
	firstCreationIns := createPathInstructions[0]
	paths := firstCreationIns.GetPaths()
	condition := chmodins.DefaultAllTrueCondition()
	existingAppliedRwxFull := firstCreationIns.ApplyRwx.String()
	for _, testCase := range chmodhelpertestwrappers.SingleRwxTestCases {
		rwxWrapper, err := testCase.ToDisabledRwxWrapper()
		errcore.SimpleHandleErr(err, "SingleRwx ToDisabledRwxWrapper failed")
		expectation := rwxWrapper.ToFullRwxValueString()

		header := fmt.Sprintf(
			"Existing [%s] Applied by [%s] should result [%s]",
			existingAppliedRwxFull,
			expectation,
			expectation)

		// Act
		err2 := rwxWrapper.ApplyLinuxChmodOnMany(condition, paths...)

		errcore.SimpleHandleErr(err2, "rwxWrapper.ApplyLinuxChmodOnMany failed")

		// for directory `-` will be placed not `d`
		// Assert
		assertSingleChmod(
			t,
			header,
			firstCreationIns,
			expectation)
	}
}
