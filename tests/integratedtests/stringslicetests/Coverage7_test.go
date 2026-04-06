package stringslicetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Empty ──

func Test_Cov7_Empty(t *testing.T) {
	// Arrange
	result := stringslice.Empty()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- with args", actual)
}

// ── IsEmpty ──

func Test_Cov7_IsEmpty_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.IsEmpty([]string{})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- empty", actual)
}

func Test_Cov7_IsEmpty_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.IsEmpty([]string{"a"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- non-empty", actual)
}

// ── HasAnyItem ──

func Test_Cov7_HasAnyItem_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.HasAnyItem([]string{})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasAnyItem returns empty -- empty", actual)
}

func Test_Cov7_HasAnyItem_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.HasAnyItem([]string{"a"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem returns empty -- non-empty", actual)
}

// ── First ──

func Test_Cov7_First(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.First([]string{"a", "b"})}

	// Assert
	expected := args.Map{"result": "a"}
	expected.ShouldBeEqual(t, 0, "First returns correct value -- with args", actual)
}

// ── Last ──

func Test_Cov7_Last(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.Last([]string{"a", "b"})}

	// Assert
	expected := args.Map{"result": "b"}
	expected.ShouldBeEqual(t, 0, "Last returns correct value -- with args", actual)
}

// ── IndexAt ──

func Test_Cov7_IndexAt(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.IndexAt([]string{"a", "b", "c"}, 1)}

	// Assert
	expected := args.Map{"result": "b"}
	expected.ShouldBeEqual(t, 0, "IndexAt returns correct value -- with args", actual)
}

// ── SafeIndexAt ──

func Test_Cov7_SafeIndexAt_Valid(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.SafeIndexAt([]string{"a", "b"}, 0)}

	// Assert
	expected := args.Map{"result": "a"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns non-empty -- valid", actual)
}

func Test_Cov7_SafeIndexAt_OutOfRange(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.SafeIndexAt([]string{"a"}, 5)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns correct value -- out of range", actual)
}

func Test_Cov7_SafeIndexAt_Negative(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.SafeIndexAt([]string{"a"}, -1)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns correct value -- negative", actual)
}

func Test_Cov7_SafeIndexAt_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.SafeIndexAt([]string{}, 0)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns empty -- empty", actual)
}

// ── MergeNew ──

func Test_Cov7_MergeNew(t *testing.T) {
	// Arrange
	result := stringslice.MergeNew([]string{"a"}, "b", "c")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "a",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "MergeNew returns correct value -- with args", actual)
}

func Test_Cov7_MergeNew_EmptyFirst(t *testing.T) {
	// Arrange
	result := stringslice.MergeNew([]string{}, "b")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "b",
	}
	expected.ShouldBeEqual(t, 0, "MergeNew returns empty -- empty first", actual)
}

// ── MergeNewSimple ──

func Test_Cov7_MergeNewSimple_Empty(t *testing.T) {
	// Arrange
	result := stringslice.MergeNewSimple()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple returns empty -- empty", actual)
}

func Test_Cov7_MergeNewSimple_Multiple(t *testing.T) {
	// Arrange
	result := stringslice.MergeNewSimple([]string{"a"}, []string{"b", "c"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple returns correct value -- multiple", actual)
}

// ── AppendLineNew ──

func Test_Cov7_AppendLineNew(t *testing.T) {
	// Arrange
	result := stringslice.AppendLineNew([]string{"a"}, "b")

	// Act
	actual := args.Map{
		"len": len(result),
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "AppendLineNew returns correct value -- with args", actual)
}

// ── PrependNew ──

func Test_Cov7_PrependNew(t *testing.T) {
	// Arrange
	result := stringslice.PrependNew([]string{"c"}, "a", "b")

	// Act
	actual := args.Map{
		"len": len(*result),
		"first": (*result)[0],
		"last": (*result)[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "a",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "PrependNew returns correct value -- with args", actual)
}

// ── InPlaceReverse ──

func Test_Cov7_InPlaceReverse_Nil(t *testing.T) {
	// Arrange
	result := stringslice.InPlaceReverse(nil)

	// Act
	actual := args.Map{"len": len(*result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns nil -- nil", actual)
}

func Test_Cov7_InPlaceReverse_Single(t *testing.T) {
	// Arrange
	s := []string{"a"}
	result := stringslice.InPlaceReverse(&s)

	// Act
	actual := args.Map{"first": (*result)[0]}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns correct value -- single", actual)
}

func Test_Cov7_InPlaceReverse_Two(t *testing.T) {
	// Arrange
	s := []string{"a", "b"}
	result := stringslice.InPlaceReverse(&s)

	// Act
	actual := args.Map{
		"first": (*result)[0],
		"second": (*result)[1],
	}

	// Assert
	expected := args.Map{
		"first": "b",
		"second": "a",
	}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns correct value -- two", actual)
}

func Test_Cov7_InPlaceReverse_Three(t *testing.T) {
	// Arrange
	s := []string{"a", "b", "c"}
	result := stringslice.InPlaceReverse(&s)

	// Act
	actual := args.Map{
		"first": (*result)[0],
		"last": (*result)[2],
	}

	// Assert
	expected := args.Map{
		"first": "c",
		"last": "a",
	}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns correct value -- three", actual)
}

// ── SortIf ──

func Test_Cov7_SortIf_True(t *testing.T) {
	// Arrange
	result := stringslice.SortIf(true, []string{"c", "a", "b"})

	// Act
	actual := args.Map{
		"first": result[0],
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "SortIf returns non-empty -- true", actual)
}

func Test_Cov7_SortIf_False(t *testing.T) {
	// Arrange
	result := stringslice.SortIf(false, []string{"c", "a", "b"})

	// Act
	actual := args.Map{"first": result[0]}

	// Assert
	expected := args.Map{"first": "c"}
	expected.ShouldBeEqual(t, 0, "SortIf returns non-empty -- false", actual)
}

// ── NonEmptySlice ──

func Test_Cov7_NonEmptySlice_Empty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlice([]string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice returns empty -- empty", actual)
}

func Test_Cov7_NonEmptySlice_Mixed(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlice([]string{"a", "", "b"})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"second": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice returns empty -- mixed", actual)
}

// ── NonEmptyJoin ──

func Test_Cov7_NonEmptyJoin_Nil(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoin(nil, ",")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns nil -- nil", actual)
}

func Test_Cov7_NonEmptyJoin_Empty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoin([]string{}, ",")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- empty", actual)
}

func Test_Cov7_NonEmptyJoin_WithEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoin([]string{"a", "", "b"}, ",")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- with empty", actual)
}

// ── ExpandBySplit ──

func Test_Cov7_ExpandBySplit_Empty(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplit([]string{}, ",")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit returns empty -- empty", actual)
}

func Test_Cov7_ExpandBySplit_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplit([]string{"a,b", "c"}, ",")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit returns empty -- non-empty", actual)
}

// ── CloneIf ──

func Test_Cov7_CloneIf_Clone(t *testing.T) {
	// Arrange
	result := stringslice.CloneIf(true, 0, []string{"a"})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "CloneIf returns correct value -- clone", actual)
}

func Test_Cov7_CloneIf_NoClone(t *testing.T) {
	// Arrange
	result := stringslice.CloneIf(false, 0, []string{"a"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CloneIf returns empty -- no clone", actual)
}

func Test_Cov7_CloneIf_NilNoClone(t *testing.T) {
	// Arrange
	result := stringslice.CloneIf(false, 0, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CloneIf returns nil -- nil no clone", actual)
}

// ── JoinWith ──

func Test_Cov7_JoinWith_Empty(t *testing.T) {
	// Arrange
	result := stringslice.JoinWith(",")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "JoinWith returns empty -- empty", actual)
}

func Test_Cov7_JoinWith_Items(t *testing.T) {
	// Arrange
	result := stringslice.JoinWith(",", "a", "b")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ",a,b"}
	expected.ShouldBeEqual(t, 0, "JoinWith returns non-empty -- items", actual)
}

// ── Joins ──

func Test_Cov7_Joins_Empty(t *testing.T) {
	// Arrange
	result := stringslice.Joins(",")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Joins returns empty -- empty", actual)
}

func Test_Cov7_Joins_Items(t *testing.T) {
	// Arrange
	result := stringslice.Joins(",", "a", "b")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "Joins returns correct value -- items", actual)
}
