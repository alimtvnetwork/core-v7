package regexnewtests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// =============================================================================
// 1. Create caching — repeated Create returns same pointer
// =============================================================================

func Test_Create_CachingBehavior(t *testing.T) {
	for caseIndex, testCase := range createCachingTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// Act
		regex1, err1 := regexnew.CreateLock(pattern)
		regex2, err2 := regexnew.CreateLock(pattern)

		actual := args.Map{
			params.samePointer: regex1 == regex2,
			params.hasError:    err1 != nil || err2 != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// 2. ManyUsingLock — empty input returns empty map
// =============================================================================

func Test_ManyUsingLock_EmptyInput(t *testing.T) {
	for caseIndex, testCase := range manyUsingLockEmptyTestCases {
		// Arrange — no patterns

		// Act
		result := regexnew.New.LazyRegex.ManyUsingLock()

		actual := args.Map{
			params.mapLength: len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// 3. ManyUsingLock — valid multi-pattern
// =============================================================================

func Test_ManyUsingLock_ValidPatterns(t *testing.T) {
	for caseIndex, testCase := range manyUsingLockValidTestCases {
		// Arrange
		patterns := []string{`\d+_many1`, `[a-z]+_many1`, `\w+_many1`}

		// Act
		result := regexnew.New.LazyRegex.ManyUsingLock(patterns...)

		actual := args.Map{
			params.mapLength: len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// 4. TwoLock — both valid patterns
// =============================================================================

func Test_TwoLock_BothValid(t *testing.T) {
	for caseIndex, testCase := range twoLockValidTestCases {
		// Arrange
		input := testCase.Input
		pattern1, _ := input.GetAsString(params.pattern)
		pattern2, _ := input.GetAsString(params.compareInput)

		// Act
		first, second := regexnew.New.LazyRegex.TwoLock(pattern1, pattern2)

		actual := args.Map{
			params.isDefined:    first.IsDefined() && second.IsDefined(),
			params.isApplicable: first.IsApplicable() && second.IsApplicable(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// 5. TwoLock — invalid patterns
// =============================================================================

func Test_TwoLock_InvalidPatterns(t *testing.T) {
	tc0 := twoLockInvalidTestCases[0]

	// Arrange
	input := tc0.Input
	pattern1, _ := input.GetAsString(params.pattern)
	pattern2, _ := input.GetAsString(params.compareInput)

	// Act
	first, second := regexnew.New.LazyRegex.TwoLock(pattern1, pattern2)

	actual := args.Map{
		params.isDefined: first.IsDefined() && second.IsDefined(),
	}

	// Assert
	tc0.ShouldBeEqualMapFirst(t, actual)

	// Verify first is applicable, second is not
	actual = args.Map{"result": first.IsApplicable()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "case 0: first pattern should be applicable", actual)
	actual = args.Map{"result": second.IsApplicable()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "case 0: second pattern should NOT be applicable", actual)
}

func Test_TwoLock_BothInvalid(t *testing.T) {
	tc1 := twoLockInvalidTestCases[1]

	// Arrange
	input := tc1.Input
	pattern1, _ := input.GetAsString(params.pattern)
	pattern2, _ := input.GetAsString(params.compareInput)

	// Act
	first, second := regexnew.New.LazyRegex.TwoLock(pattern1, pattern2)

	actual := args.Map{
		params.isDefined: first.IsDefined() && second.IsDefined(),
	}

	// Assert
	tc1.ShouldBeEqualMapFirst(t, actual)

	// Verify both are not applicable
	actual = args.Map{"result": first.IsApplicable()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "case 1: first pattern should NOT be applicable", actual)
	actual = args.Map{"result": second.IsApplicable()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "case 1: second pattern should NOT be applicable", actual)
}

// =============================================================================
// 6. MatchError error message — compile error branch
// =============================================================================

func Test_MatchError_ErrorMessage_CompileError(t *testing.T) {
	for caseIndex, testCase := range matchErrorMessageCompileTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		err := regexnew.MatchErrorLock(pattern, compareInput)

		expectedContains, _ := testCase.Expected.GetAsString(params.errorContains)
		actual := args.Map{
			params.hasError:      err != nil,
			params.errorContains: containsOrEmpty(err, expectedContains),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// 7. MatchError error message — mismatch branch
// =============================================================================

func Test_MatchError_ErrorMessage_Mismatch(t *testing.T) {
	for caseIndex, testCase := range matchErrorMessageMismatchTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		err := regexnew.MatchErrorLock(pattern, compareInput)

		expectedContains, _ := testCase.Expected.GetAsString(params.errorContains)
		actual := args.Map{
			params.hasError:      err != nil,
			params.errorContains: containsOrEmpty(err, expectedContains),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// 8. LazyRegex.MatchError error message verification
// =============================================================================

func Test_LazyRegex_MatchError_ErrorMessages(t *testing.T) {
	for caseIndex, testCase := range lazyMatchErrorMessageTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)
		lazy := regexnew.New.LazyLock(pattern)

		// Act
		err := lazy.MatchError(compareInput)

		expectedContains, _ := testCase.Expected.GetAsString(params.errorContains)
		actual := args.Map{
			params.hasError:      err != nil,
			params.errorContains: containsOrEmpty(err, expectedContains),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// containsOrEmpty returns the expected substring if err contains it,
// or the actual error message if it doesn't (to surface the mismatch in diagnostics).
func containsOrEmpty(err error, expected string) string {
	if err == nil {
		return ""
	}

	msg := err.Error()
	if strings.Contains(msg, expected) {
		return expected
	}

	return msg
}
