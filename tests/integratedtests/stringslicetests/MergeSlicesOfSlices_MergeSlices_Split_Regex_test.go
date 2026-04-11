package stringslicetests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_MergeSlicesOfSlices_FromMergeSlicesOfSlicesM(t *testing.T) {
	// Arrange
	result := stringslice.MergeSlicesOfSlices([]string{"a"}, []string{"b", "c"})

	// Act
	actual := args.Map{"result": len(result) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	// with nil slice
	result2 := stringslice.MergeSlicesOfSlices([]string{"a"}, []string{})
	actual = args.Map{"result": len(result2) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	// empty
	result3 := stringslice.MergeSlicesOfSlices()
	actual = args.Map{"result": len(result3) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_SplitTrimmedNonEmptyAll_FromMergeSlicesOfSlicesM(t *testing.T) {
	// Arrange
	result := stringslice.SplitTrimmedNonEmptyAll("a , b , c", ",")

	// Act
	actual := args.Map{"result": len(result) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_SplitTrimmedNonEmpty_FromMergeSlicesOfSlicesM(t *testing.T) {
	// Arrange
	result := stringslice.SplitTrimmedNonEmpty("a,b,c", ",", -1)

	// Act
	actual := args.Map{"result": len(result) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_RegexTrimmedSplitNonEmptyAll_FromMergeSlicesOfSlicesM(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`[,;]`)
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, "a , b ; c")

	// Act
	actual := args.Map{"result": len(result) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}
