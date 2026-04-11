package regexnewtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// ==========================================================================
// Test: IsMatch
// ==========================================================================

func Test_LazyRegex_IsMatch_FullDigit(t *testing.T) {
	tc := lazyRegexIsMatchFullDigitTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)
	lazyRegex := regexnew.New.LazyLock(pattern)

	// Act
	actual := args.Map{
		params.isMatch: lazyRegex.IsMatch(compareInput),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LazyRegex_IsMatch_PartialMismatch(t *testing.T) {
	tc := lazyRegexIsMatchPartialMismatchTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)
	lazyRegex := regexnew.New.LazyLock(pattern)

	// Act
	actual := args.Map{
		params.isMatch: lazyRegex.IsMatch(compareInput),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: IsFailedMatch
// ==========================================================================

func Test_LazyRegex_IsFailedMatch_FromLazyRegexPatternMatc(t *testing.T) {
	tc := lazyRegexIsFailedMatchTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)
	lazyRegex := regexnew.New.LazyLock(pattern)

	// Act
	actual := args.Map{
		params.isFailedMatch: lazyRegex.IsFailedMatch(compareInput),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: FirstMatchLine
// ==========================================================================

func Test_LazyRegex_FirstMatchLine_Found(t *testing.T) {
	tc := lazyRegexFirstMatchLineFoundTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)
	lazyRegex := regexnew.New.LazyLock(pattern)

	// Act
	firstMatch, isInvalid := lazyRegex.FirstMatchLine(compareInput)

	actual := args.Map{
		params.firstMatch:     firstMatch,
		params.isInvalidMatch: isInvalid,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LazyRegex_FirstMatchLine_NotFound(t *testing.T) {
	tc := lazyRegexFirstMatchLineNotFoundTestCase

	// Arrange
	pattern, _ := tc.Input.GetAsString(params.pattern)
	compareInput, _ := tc.Input.GetAsString(params.compareInput)
	lazyRegex := regexnew.New.LazyLock(pattern)

	// Act
	firstMatch, isInvalid := lazyRegex.FirstMatchLine(compareInput)

	actual := args.Map{
		params.firstMatch:     firstMatch,
		params.isInvalidMatch: isInvalid,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
