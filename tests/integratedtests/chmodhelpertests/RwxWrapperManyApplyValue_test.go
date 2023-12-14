package chmodhelpertests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

// Test_RwxWrapperManyApplyValue_Unix
//
//	for directory `-` will be placed not `d`
func Test_RwxWrapperManyApplyValue_Unix(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	createPathInstructions := pathInstructionsV2
	chmodhelper.CreateDirFilesWithRwxPermissionsMust(
		true,
		createPathInstructions,
	)
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
			expectation,
		)

		// Act
		err2 := rwxWrapper.ApplyLinuxChmodOnMany(condition, paths...)
		errcore.SimpleHandleErr(
			err2,
			"rwxWrapper.ApplyLinuxChmodOnMany failed",
		)

		// Assert
		assertSingleChmod(
			t,
			header,
			firstCreationIns,
			expectation,
		)
	}
}
