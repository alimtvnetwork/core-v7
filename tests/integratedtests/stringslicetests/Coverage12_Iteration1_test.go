package stringslicetests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// Test_Cov12_MergeSlicesOfSlices_AllEmpty tests MergeSlicesOfSlices with all-empty slices.
func Test_Cov12_MergeSlicesOfSlices_AllEmpty(t *testing.T) {
	// Arrange / Act
	actual := stringslice.MergeSlicesOfSlices([]string{}, []string{})

	// Assert
	actual := args.Map{"result": len(actual) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices with all empty: got len, want 0", actual)
}

// Test_Cov12_RegexTrimmedSplitNonEmptyAll_EmptyResult tests empty content with regex split.
func Test_Cov12_RegexTrimmedSplitNonEmptyAll_EmptyResult(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`\s+`)

	// Act
	actual := stringslice.RegexTrimmedSplitNonEmptyAll(re, "   ")

	// Assert
	actual := args.Map{"result": len(actual) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RegexTrimmedSplitNonEmptyAll with whitespace-only: got len, want 0", actual)
}

// Test_Cov12_SplitTrimmedNonEmpty_ZeroCount tests SplitTrimmedNonEmpty with n=0.
func Test_Cov12_SplitTrimmedNonEmpty_ZeroCount(t *testing.T) {
	// Arrange / Act
	actual := stringslice.SplitTrimmedNonEmpty("a,b,c", ",", 0)

	// Assert
	actual := args.Map{"result": len(actual) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmpty with n=0: got len, want 0", actual)
}

// Test_Cov12_SplitTrimmedNonEmptyAll_EmptyContent tests with empty splitter result.
func Test_Cov12_SplitTrimmedNonEmptyAll_EmptyContent(t *testing.T) {
	// Arrange / Act
	actual := stringslice.SplitTrimmedNonEmptyAll("", "")

	// Assert
	actual := args.Map{"result": len(actual) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmptyAll empty/empty: got len, want 0", actual)
}
