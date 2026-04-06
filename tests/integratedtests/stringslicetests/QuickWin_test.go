package stringslicetests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_MergeSlicesOfSlices_Empty(t *testing.T) {
	// Arrange
	result := stringslice.MergeSlicesOfSlices()

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_QW_RegexTrimmedSplitNonEmptyAll_Empty(t *testing.T) {
	re := regexp.MustCompile(`\s+`)
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, "")
	_ = result
}

func Test_QW_SplitTrimmedNonEmpty_Empty(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmpty("", ",", -1)
	_ = result
}

func Test_QW_SplitTrimmedNonEmptyAll_Empty(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmptyAll("", ",")
	_ = result
}
