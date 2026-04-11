package regexnewtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

func Test_Create_Verification(t *testing.T) {
	for caseIndex, testCase := range createTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// Act
		regex, err := regexnew.New.DefaultLock(pattern)

		actual := args.Map{
			params.regexNotNil: regex != nil,
			params.hasError:    err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Create_IsMatchLock_Verification(t *testing.T) {
	for caseIndex, testCase := range createIsMatchLockTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		actual := args.Map{
			params.isMatch: regexnew.IsMatchLock(pattern, compareInput),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Create_IsMatchFailed_Verification(t *testing.T) {
	for caseIndex, testCase := range createIsMatchFailedTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		actual := args.Map{
			params.isFailed: regexnew.IsMatchFailed(pattern, compareInput),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// Test: MatchError
// ==========================================================================

func Test_MatchError_Match_FromCreate(t *testing.T) {
	tc := matchErrorMatchTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)

	// Act
	err := regexnew.MatchError(pattern, compareInput)

	actual := args.Map{
		params.isNoError: err == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MatchError_Mismatch(t *testing.T) {
	tc := matchErrorMismatchTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)

	// Act
	err := regexnew.MatchError(pattern, compareInput)

	actual := args.Map{
		params.isNoError: err == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: MatchErrorLock
// ==========================================================================

func Test_MatchErrorLock_Match_FromCreate(t *testing.T) {
	tc := matchErrorLockMatchTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)

	// Act
	err := regexnew.MatchErrorLock(pattern, compareInput)

	actual := args.Map{
		params.isNoError: err == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MatchErrorLock_Mismatch(t *testing.T) {
	tc := matchErrorLockMismatchTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)

	// Act
	err := regexnew.MatchErrorLock(pattern, compareInput)

	actual := args.Map{
		params.isNoError: err == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
