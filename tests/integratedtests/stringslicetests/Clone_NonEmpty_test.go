package stringslicetests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Clone ──

func Test_Clone_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.Clone([]string{"a", "b"})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "Clone returns empty -- non-empty", actual)
}

func Test_Clone_Empty(t *testing.T) {
	// Arrange
	result := stringslice.Clone(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clone returns nil -- nil", actual)
}

// ── FirstLastDefault ──

func Test_FirstLastDefault_SingleElement(t *testing.T) {
	// Arrange
	first, last := stringslice.FirstLastDefault([]string{"only"})

	// Act
	actual := args.Map{
		"first": first,
		"last": last,
	}

	// Assert
	expected := args.Map{
		"first": "only",
		"last": "",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault returns correct value -- single element", actual)
}

func Test_FirstLastDefault_Empty(t *testing.T) {
	// Arrange
	first, last := stringslice.FirstLastDefault(nil)

	// Act
	actual := args.Map{
		"first": first,
		"last": last,
	}

	// Assert
	expected := args.Map{
		"first": "",
		"last": "",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault returns empty -- empty", actual)
}

func Test_FirstLastDefault_Multi(t *testing.T) {
	// Arrange
	first, last := stringslice.FirstLastDefault([]string{"a", "b", "c"})

	// Act
	actual := args.Map{
		"first": first,
		"last": last,
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault returns correct value -- multi", actual)
}

// ── SafeIndexAt ──

func Test_SafeIndexAt(t *testing.T) {
	// Arrange
	slice := []string{"a", "b", "c"}

	// Act
	actual := args.Map{
		"at0":        stringslice.SafeIndexAt(slice, 0),
		"at2":        stringslice.SafeIndexAt(slice, 2),
		"atNeg":      stringslice.SafeIndexAt(slice, -1),
		"atOutBound": stringslice.SafeIndexAt(slice, 10),
		"emptySlice": stringslice.SafeIndexAt(nil, 0),
	}

	// Assert
	expected := args.Map{
		"at0": "a", "at2": "c",
		"atNeg": "", "atOutBound": "", "emptySlice": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns correct value -- with args", actual)
}

// ── NonEmptyJoin ──

func Test_NonEmptyJoin(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoin([]string{"a", "", "b"}, ",")
	resultNil := stringslice.NonEmptyJoin(nil, ",")
	resultEmpty := stringslice.NonEmptyJoin([]string{}, ",")

	// Act
	actual := args.Map{
		"result":    result,
		"nilResult": resultNil,
		"emptyRes":  resultEmpty,
	}

	// Assert
	expected := args.Map{
		"result": "a,b", "nilResult": "", "emptyRes": "",
	}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- with args", actual)
}

// ── MergeNew ──

func Test_MergeNew(t *testing.T) {
	// Arrange
	result := stringslice.MergeNew([]string{"a"}, "b", "c")
	resultEmpty := stringslice.MergeNew(nil, "b")

	// Act
	actual := args.Map{
		"len":      len(result),
		"emptyLen": len(resultEmpty),
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"emptyLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "MergeNew returns correct value -- with args", actual)
}

// ── InPlaceReverse ──

func Test_InPlaceReverse(t *testing.T) {
	// Arrange
	slice := []string{"a", "b", "c", "d"}
	result := stringslice.InPlaceReverse(&slice)

	// Act
	actual := args.Map{
		"first": (*result)[0],
		"last": (*result)[3],
	}

	// Assert
	expected := args.Map{
		"first": "d",
		"last": "a",
	}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns correct value -- 4 elements", actual)
}

func Test_InPlaceReverse_Two(t *testing.T) {
	// Arrange
	slice := []string{"a", "b"}
	result := stringslice.InPlaceReverse(&slice)

	// Act
	actual := args.Map{
		"first": (*result)[0],
		"last": (*result)[1],
	}

	// Assert
	expected := args.Map{
		"first": "b",
		"last": "a",
	}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns correct value -- 2 elements", actual)
}

func Test_InPlaceReverse_SingleAndNil(t *testing.T) {
	// Arrange
	single := []string{"only"}
	r1 := stringslice.InPlaceReverse(&single)
	r2 := stringslice.InPlaceReverse(nil)

	// Act
	actual := args.Map{
		"singleFirst": (*r1)[0],
		"nilLen":      len(*r2),
	}

	// Assert
	expected := args.Map{
		"singleFirst": "only",
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns nil -- single and nil", actual)
}

// ── SortIf ──

func Test_SortIf(t *testing.T) {
	// Arrange
	slice := []string{"c", "a", "b"}
	sorted := stringslice.SortIf(true, slice)
	notSorted := stringslice.SortIf(false, []string{"c", "a"})

	// Act
	actual := args.Map{
		"sortedFirst":    sorted[0],
		"notSortedFirst": notSorted[0],
	}

	// Assert
	expected := args.Map{
		"sortedFirst": "a",
		"notSortedFirst": "c",
	}
	expected.ShouldBeEqual(t, 0, "SortIf returns correct value -- with args", actual)
}

// ── ExpandByFunc ──

func Test_ExpandByFunc(t *testing.T) {
	// Arrange
	result := stringslice.ExpandByFunc(
		[]string{"a,b", "c,d"},
		func(line string) []string { return strings.Split(line, ",") },
	)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc returns correct value -- with args", actual)
}

func Test_ExpandByFunc_Empty(t *testing.T) {
	// Arrange
	result := stringslice.ExpandByFunc(nil, func(line string) []string { return nil })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc returns empty -- empty", actual)
}

// ── AllElemLengthSlices ──

func Test_AllElemLengthSlices(t *testing.T) {
	// Arrange
	result := stringslice.AllElemLengthSlices(
		[]string{"a", "b"},
		nil,
		[]string{"c"},
	)

	// Act
	actual := args.Map{"total": result}

	// Assert
	expected := args.Map{"total": 3}
	expected.ShouldBeEqual(t, 0, "AllElemLengthSlices returns correct value -- with args", actual)
}

func Test_AllElemLengthSlices_Empty(t *testing.T) {
	// Arrange
	result := stringslice.AllElemLengthSlices()

	// Act
	actual := args.Map{"total": result}

	// Assert
	expected := args.Map{"total": 0}
	expected.ShouldBeEqual(t, 0, "AllElemLengthSlices returns empty -- empty", actual)
}

// ── PrependLineNew ──

func Test_PrependLineNew(t *testing.T) {
	// Arrange
	result := stringslice.PrependLineNew("first", []string{"second", "third"})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "first",
	}
	expected.ShouldBeEqual(t, 0, "PrependLineNew returns correct value -- with args", actual)
}

// ── AppendLineNew ──

func Test_AppendLineNew(t *testing.T) {
	// Arrange
	result := stringslice.AppendLineNew([]string{"a", "b"}, "c")

	// Act
	actual := args.Map{
		"len": len(result),
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "AppendLineNew returns correct value -- with args", actual)
}
