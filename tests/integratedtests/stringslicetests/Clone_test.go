package stringslicetests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Clone / ClonePtr / CloneUsingCap ──

func Test_Clone(t *testing.T) {
	// Arrange
	result := stringslice.Clone([]string{"a", "b"})
	nilResult := stringslice.Clone(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"nilLen": len(nilResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "Clone returns correct -- 2 items", actual)
}

func Test_ClonePtr(t *testing.T) {
	// Arrange
	result := stringslice.ClonePtr([]string{"a"})
	emptyResult := stringslice.ClonePtr(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns correct -- 1 item", actual)
}

func Test_CloneUsingCap(t *testing.T) {
	// Arrange
	result := stringslice.CloneUsingCap(5, []string{"a"})
	empty := stringslice.CloneUsingCap(5, nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(empty),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "CloneUsingCap returns correct -- cap 5", actual)
}

// ── Empty / EmptyPtr ──

func Test_Empty(t *testing.T) {
	// Act
	actual := args.Map{
		"len": len(stringslice.Empty()),
		"ptrLen": len(stringslice.EmptyPtr()),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"ptrLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- zero items", actual)
}

// ── First / FirstPtr / FirstOrDefault / FirstOrDefaultWith ──

func Test_First(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.First([]string{"a", "b"})}

	// Assert
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "First returns first -- 2 items", actual)
}

func Test_FirstPtr(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.FirstPtr([]string{"x"})}

	// Assert
	expected := args.Map{"val": "x"}
	expected.ShouldBeEqual(t, 0, "FirstPtr returns first -- 1 item", actual)
}

func Test_FirstOrDefault(t *testing.T) {
	// Act
	actual := args.Map{
		"found": stringslice.FirstOrDefault([]string{"a"}),
		"empty": stringslice.FirstOrDefault(nil),
	}

	// Assert
	expected := args.Map{
		"found": "a",
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault returns correct -- found and empty", actual)
}

func Test_FirstOrDefaultWith(t *testing.T) {
	// Arrange
	v, ok := stringslice.FirstOrDefaultWith([]string{"a"}, "def")
	v2, ok2 := stringslice.FirstOrDefaultWith(nil, "def")

	// Act
	actual := args.Map{
		"v": v,
		"ok": ok,
		"v2": v2,
		"ok2": ok2,
	}

	// Assert
	expected := args.Map{
		"v": "a",
		"ok": true,
		"v2": "def",
		"ok2": false,
	}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultWith returns correct -- found and default", actual)
}

// ── Last / LastPtr / LastOrDefault ──

func Test_Last(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.Last([]string{"a", "b"})}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "Last returns last -- 2 items", actual)
}

func Test_LastPtr(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.LastPtr([]string{"x", "y"})}

	// Assert
	expected := args.Map{"val": "y"}
	expected.ShouldBeEqual(t, 0, "LastPtr returns last -- 2 items", actual)
}

func Test_LastOrDefault(t *testing.T) {
	// Act
	actual := args.Map{
		"found": stringslice.LastOrDefault([]string{"a", "b"}),
		"empty": stringslice.LastOrDefault(nil),
	}

	// Assert
	expected := args.Map{
		"found": "b",
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "LastOrDefault returns correct -- found and empty", actual)
}

// ── FirstLastDefault / FirstLastDefaultStatus ──

func Test_FirstLastDefault_Single(t *testing.T) {
	// Arrange
	f, l := stringslice.FirstLastDefault([]string{"a"})

	// Act
	actual := args.Map{
		"first": f,
		"last": l,
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"last": "",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault single -- first only", actual)
}

func Test_FirstLastDefaultStatus_Single(t *testing.T) {
	// Arrange
	s := stringslice.FirstLastDefaultStatus([]string{"a"})

	// Act
	actual := args.Map{
		"first": s.First,
		"hasFirst": s.HasFirst,
		"hasLast": s.HasLast,
		"isValid": s.IsValid,
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"hasFirst": true,
		"hasLast": false,
		"isValid": false,
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus single -- first only", actual)
}

func Test_FirstLastDefaultStatus_Multiple(t *testing.T) {
	// Arrange
	s := stringslice.FirstLastDefaultStatus([]string{"a", "b", "c"})

	// Act
	actual := args.Map{
		"first": s.First,
		"last": s.Last,
		"isValid": s.IsValid,
		"hasFirst": s.HasFirst,
		"hasLast": s.HasLast,
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"last": "c",
		"isValid": true,
		"hasFirst": true,
		"hasLast": true,
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus multiple -- all valid", actual)
}

// ── HasAnyItem / HasAnyItemPtr / IsEmpty / IsEmptyPtr / LengthOfPointer / SlicePtr ──

func Test_HasAnyItem(t *testing.T) {
	// Act
	actual := args.Map{
		"yes": stringslice.HasAnyItem([]string{"a"}),
		"no": stringslice.HasAnyItem(nil),
	}

	// Assert
	expected := args.Map{
		"yes": true,
		"no": false,
	}
	expected.ShouldBeEqual(t, 0, "HasAnyItem returns correct value -- returns correct", actual)
}

func Test_HasAnyItemPtr(t *testing.T) {
	// Act
	actual := args.Map{
		"yes": stringslice.HasAnyItemPtr([]string{"a"}),
		"no": stringslice.HasAnyItemPtr(nil),
	}

	// Assert
	expected := args.Map{
		"yes": true,
		"no": false,
	}
	expected.ShouldBeEqual(t, 0, "HasAnyItemPtr returns correct value -- returns correct", actual)
}

func Test_IsEmptyPtr(t *testing.T) {
	// Act
	actual := args.Map{
		"empty": stringslice.IsEmptyPtr(nil),
		"notEmpty": stringslice.IsEmptyPtr([]string{"a"}),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"notEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns empty -- returns correct", actual)
}

func Test_LengthOfPointer(t *testing.T) {
	// Act
	actual := args.Map{"len": stringslice.LengthOfPointer([]string{"a", "b"})}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfPointer returns correct value -- returns 2", actual)
}

func Test_SlicePtr(t *testing.T) {
	// Act
	actual := args.Map{
		"len": len(stringslice.SlicePtr([]string{"a"})),
		"emptyLen": len(stringslice.SlicePtr(nil)),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "SlicePtr returns correct value -- returns correct", actual)
}

// ── IndexAt / SafeIndexAt / SafeIndexAtWith ──

func Test_IndexAt(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.IndexAt([]string{"a", "b"}, 1)}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "IndexAt returns correct -- index 1", actual)
}

func Test_SafeIndexAt_FromClone(t *testing.T) {
	// Act
	actual := args.Map{
		"valid":   stringslice.SafeIndexAt([]string{"a", "b"}, 1),
		"invalid": stringslice.SafeIndexAt([]string{"a"}, 5),
		"neg":     stringslice.SafeIndexAt([]string{"a"}, -1),
		"empty":   stringslice.SafeIndexAt(nil, 0),
	}

	// Assert
	expected := args.Map{
		"valid": "b",
		"invalid": "",
		"neg": "",
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns correct -- all branches", actual)
}

func Test_SafeIndexAtWith(t *testing.T) {
	// Act
	actual := args.Map{
		"valid":   stringslice.SafeIndexAtWith([]string{"a", "b"}, 1, "def"),
		"invalid": stringslice.SafeIndexAtWith([]string{"a"}, 5, "def"),
	}

	// Assert
	expected := args.Map{
		"valid": "b",
		"invalid": "def",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWith returns non-empty -- returns correct", actual)
}

// ── SafeIndexAtUsingLastIndexPtr ──

func Test_SafeIndexAtUsingLastIndexPtr(t *testing.T) {
	// Act
	actual := args.Map{
		"valid":    stringslice.SafeIndexAtUsingLastIndexPtr([]string{"a", "b"}, 1, 0),
		"zeroLast": stringslice.SafeIndexAtUsingLastIndexPtr([]string{"a"}, 0, 0),
		"negLast":  stringslice.SafeIndexAtUsingLastIndexPtr([]string{"a"}, -1, 0),
	}

	// Assert
	expected := args.Map{
		"valid": "a",
		"zeroLast": "",
		"negLast": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndexPtr returns correct value -- returns correct", actual)
}

// ── NonEmpty / NonEmptyIf / NonEmptyStrings / NonNullStrings ──

func Test_NonEmptySlice(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlice([]string{"a", "", "b"})
	nilResult := stringslice.NonEmptySlice(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"nilLen": len(nilResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice returns 2 -- skip empty", actual)
}

func Test_NonEmptyIf(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyIf(true, []string{"a", "", "b"})
	result2 := stringslice.NonEmptyIf(false, []string{"a", "", "b"})

	// Act
	actual := args.Map{
		"trueLen": len(result),
		"falseLen": len(result2),
	}

	// Assert
	expected := args.Map{
		"trueLen": 2,
		"falseLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "NonEmptyIf returns correct -- true and false", actual)
}

func Test_NonEmptyStrings(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyStrings([]string{"a", "", "b"})
	nilResult := stringslice.NonEmptyStrings(nil)
	emptyResult := stringslice.NonEmptyStrings([]string{})

	// Act
	actual := args.Map{
		"len": len(result),
		"nilLen": len(nilResult),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"nilLen": 0,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings returns correct -- all branches", actual)
}

func Test_NonNullStrings(t *testing.T) {
	// Arrange
	result := stringslice.NonNullStrings([]string{"a", "", "b"})
	nilResult := stringslice.NonNullStrings(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"nilLen": len(nilResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "NonNullStrings returns correct value -- returns correct", actual)
}

// ── NonWhitespace / NonWhitespaceJoin ──

func Test_NonWhitespace(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespace([]string{"a", "  ", "", "b"})
	nilResult := stringslice.NonWhitespace(nil)
	emptyResult := stringslice.NonWhitespace([]string{})

	// Act
	actual := args.Map{
		"len": len(result),
		"nilLen": len(nilResult),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"nilLen": 0,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "NonWhitespace returns correct -- all branches", actual)
}

func Test_NonWhitespaceJoin(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespaceJoin([]string{"a", "  ", "b"}, ",")
	nilResult := stringslice.NonWhitespaceJoin(nil, ",")
	emptyResult := stringslice.NonWhitespaceJoin([]string{}, ",")

	// Act
	actual := args.Map{
		"val": result,
		"nilVal": nilResult,
		"emptyVal": emptyResult,
	}

	// Assert
	expected := args.Map{
		"val": "a,b",
		"nilVal": "",
		"emptyVal": "",
	}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns correct value -- returns correct", actual)
}

func Test_NonEmptyJoin_FromClone(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoin([]string{"a", "", "b"}, ",")
	nilResult := stringslice.NonEmptyJoin(nil, ",")
	emptyResult := stringslice.NonEmptyJoin([]string{}, ",")

	// Act
	actual := args.Map{
		"val": result,
		"nilVal": nilResult,
		"emptyVal": emptyResult,
	}

	// Assert
	expected := args.Map{
		"val": "a,b",
		"nilVal": "",
		"emptyVal": "",
	}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- returns correct", actual)
}

// ── InPlaceReverse ──

func Test_InPlaceReverse_FromClone(t *testing.T) {
	// Arrange
	items := []string{"a", "b", "c"}
	result := stringslice.InPlaceReverse(&items)
	two := []string{"x", "y"}
	result2 := stringslice.InPlaceReverse(&two)
	single := []string{"z"}
	result3 := stringslice.InPlaceReverse(&single)
	result4 := stringslice.InPlaceReverse(nil)

	// Act
	actual := args.Map{
		"first": (*result)[0], "last": (*result)[2],
		"twoFirst": (*result2)[0], "singleFirst": (*result3)[0],
		"nilLen": len(*result4),
	}

	// Assert
	expected := args.Map{
		"first": "c",
		"last": "a",
		"twoFirst": "y",
		"singleFirst": "z",
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns correct -- all branches", actual)
}

// ── TrimmedEachWords / TrimmedEachWordsIf ──

func Test_TrimmedEachWords(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWords([]string{"  a  ", "  ", " b "})
	nilResult := stringslice.TrimmedEachWords(nil)
	emptyResult := stringslice.TrimmedEachWords([]string{})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"nilNil": nilResult == nil,
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
		"nilNil": true,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWords returns correct value -- returns correct", actual)
}

func Test_TrimmedEachWordsIf(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWordsIf(true, []string{"  a  ", " "})
	result2 := stringslice.TrimmedEachWordsIf(false, []string{"  a  ", " "})

	// Act
	actual := args.Map{
		"trueLen": len(result),
		"falseLen": len(result2),
	}

	// Assert
	expected := args.Map{
		"trueLen": 1,
		"falseLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsIf returns correct value -- returns correct", actual)
}

// ── MergeNew / MergeNewSimple / MergeSlicesOfSlices / PrependNew / AppendLineNew ──

func Test_MergeNew_FromClone(t *testing.T) {
	// Arrange
	result := stringslice.MergeNew([]string{"a"}, "b", "c")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MergeNew returns 3 -- 1+2", actual)
}

func Test_MergeNewSimple(t *testing.T) {
	// Arrange
	result := stringslice.MergeNewSimple([]string{"a"}, []string{"b", "c"}, nil)
	emptyResult := stringslice.MergeNewSimple()

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple returns correct value -- returns correct", actual)
}

func Test_MergeSlicesOfSlices(t *testing.T) {
	// Arrange
	result := stringslice.MergeSlicesOfSlices([]string{"a"}, nil, []string{"b"})
	emptyResult := stringslice.MergeSlicesOfSlices()

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices returns correct value -- returns correct", actual)
}

func Test_PrependNew(t *testing.T) {
	// Arrange
	result := stringslice.PrependNew([]string{"b"}, "a")

	// Act
	actual := args.Map{
		"first": (*result)[0],
		"len": len(*result),
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "PrependNew returns correct -- prepend a", actual)
}

func Test_AppendLineNew_FromClone(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "AppendLineNew returns correct -- append b", actual)
}

// ── AppendStringsWithMainSlice ──

func Test_AppendStringsWithMainSlice(t *testing.T) {
	// Arrange
	result := stringslice.AppendStringsWithMainSlice(true, []string{"a"}, "b", "", "c")
	noSkip := stringslice.AppendStringsWithMainSlice(false, []string{"a"}, "b", "", "c")
	noAppend := stringslice.AppendStringsWithMainSlice(false, []string{"a"})

	// Act
	actual := args.Map{
		"skipLen": len(result),
		"noSkipLen": len(noSkip),
		"noAppendLen": len(noAppend),
	}

	// Assert
	expected := args.Map{
		"skipLen": 3,
		"noSkipLen": 4,
		"noAppendLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice returns non-empty -- returns correct", actual)
}

// ── AppendStringsWithAnyItems / AppendAnyItemsWithStrings ──

func Test_AppendStringsWithAnyItems_Func(t *testing.T) {
	// Arrange
	result := stringslice.AppendStringsWithAnyItems(true, true, []any{1}, "a", "", "b")
	noAppend := stringslice.AppendStringsWithAnyItems(false, false, []any{1})

	// Act
	actual := args.Map{
		"len": len(result),
		"noAppendLen": len(noAppend),
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"noAppendLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems returns non-empty -- returns correct", actual)
}

func Test_AppendAnyItemsWithStrings_Func(t *testing.T) {
	// Arrange
	result := stringslice.AppendAnyItemsWithStrings(true, true, []string{"a"}, 1, nil, 2)
	noAppend := stringslice.AppendAnyItemsWithStrings(false, false, []string{"a"})

	// Act
	actual := args.Map{
		"len": len(result),
		"noAppendLen": len(noAppend),
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"noAppendLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings returns non-empty -- returns correct", actual)
}

// ── AnyItemsCloneIf / AnyItemsCloneUsingCap ──

func Test_AnyItemsCloneIf(t *testing.T) {
	// Arrange
	result := stringslice.AnyItemsCloneIf(true, 5, []any{1, 2})
	noClone := stringslice.AnyItemsCloneIf(false, 0, []any{1})
	nilNoClone := stringslice.AnyItemsCloneIf(false, 0, nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"noCloneLen": len(noClone),
		"nilLen": len(nilNoClone),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"noCloneLen": 1,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf returns correct value -- returns correct", actual)
}

// ── AllElemLengthSlices ──

func Test_AllElemLengthSlices_FromClone(t *testing.T) {
	// Arrange
	result := stringslice.AllElemLengthSlices([]string{"a"}, nil, []string{"b", "c"})
	emptyResult := stringslice.AllElemLengthSlices()

	// Act
	actual := args.Map{
		"count": result,
		"emptyCount": emptyResult,
	}

	// Assert
	expected := args.Map{
		"count": 3,
		"emptyCount": 0,
	}
	expected.ShouldBeEqual(t, 0, "AllElemLengthSlices returns correct value -- returns correct", actual)
}

// ── SafeRangeItems ──

func Test_SafeRangeItems(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems([]string{"a", "b", "c", "d"}, 1, 3)
	nilResult := stringslice.SafeRangeItems(nil, 0, 1)
	emptyResult := stringslice.SafeRangeItems([]string{}, 0, 1)
	outResult := stringslice.SafeRangeItems([]string{"a"}, 5, 6)

	// Act
	actual := args.Map{
		"len": len(result),
		"nilLen": len(nilResult),
		"emptyLen": len(emptyResult),
		"outLen": len(outResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"nilLen": 0,
		"emptyLen": 0,
		"outLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns correct value -- returns correct", actual)
}

// ── SafeIndexesDefaultWithDetail ──

func Test_SafeIndexesDefaultWithDetail(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexesDefaultWithDetail([]string{"a", "b", "c"}, 0, 2, 99)
	emptyResult := stringslice.SafeIndexesDefaultWithDetail(nil, 0)

	// Act
	actual := args.Map{
		"valuesLen":  len(result.Values),
		"anyMissing": result.IsAnyMissing,
		"isValid":    result.IsValid,
		"emptyValid": emptyResult.IsValid,
	}

	// Assert
	expected := args.Map{
		"valuesLen": 2,
		"anyMissing": true,
		"isValid": true,
		"emptyValid": false,
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexesDefaultWithDetail returns non-empty -- returns correct", actual)
}

// ── IndexesDefault ──

func Test_IndexesDefault(t *testing.T) {
	// Arrange
	result := stringslice.IndexesDefault([]string{"a", "b", "c"}, 0, 2)
	emptyResult := stringslice.IndexesDefault(nil, 0)
	noIndexResult := stringslice.IndexesDefault([]string{"a"})

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
		"noIdxLen": len(noIndexResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
		"noIdxLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "IndexesDefault returns correct value -- returns correct", actual)
}

// ── SplitTrimmedNonEmpty ──

func Test_SplitTrimmedNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.SplitTrimmedNonEmpty("a, , b", ",", -1)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmpty returns 2 -- skip empty", actual)
}

// ── RegexTrimmedSplitNonEmptyAll ──

func Test_RegexTrimmedSplitNonEmptyAll(t *testing.T) {
	// Arrange
	re := regexp.MustCompile("[,;]")
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, "a, ;b; c")

	// Act
	actual := args.Map{"gt0": len(result) > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "RegexTrimmedSplitNonEmptyAll returns empty -- returns items", actual)
}

// ── ExpandByFunc / ExpandBySplit / ExpandBySplits ──

func Test_ExpandByFunc_FromClone(t *testing.T) {
	// Arrange
	result := stringslice.ExpandByFunc([]string{"a,b", "c,d"}, func(s string) []string { return []string{s + "!"} })
	emptyResult := stringslice.ExpandByFunc(nil, nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc returns correct value -- returns correct", actual)
}

func Test_ExpandBySplit(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplit([]string{"a,b", "c,d"}, ",")
	emptyResult := stringslice.ExpandBySplit(nil, ",")

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 4,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit returns correct value -- returns correct", actual)
}

// ── LinesProcess ──

func Test_LinesProcess(t *testing.T) {
	// Arrange
	result := stringslice.LinesProcess([]string{"a", "b", "c"}, func(i int, s string) (string, bool, bool) {
		return s + "!", true, i == 1
	})
	emptyResult := stringslice.LinesProcess(nil, nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "LinesProcess returns correct -- break at 1", actual)
}

// ── MakeDefault / MakeDefaultPtr / MakeLen / MakePtr / MakeLenPtr ──

func Test_MakeDefault(t *testing.T) {
	// Act
	actual := args.Map{
		"defLen":    len(stringslice.MakeDefault(5)),
		"defPtrLen": len(stringslice.MakeDefaultPtr(5)),
	}

	// Assert
	expected := args.Map{
		"defLen": 0,
		"defPtrLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "MakeDefault returns empty -- cap 5", actual)
}

// ── SafeIndexesPtr (deprecated) ──

func Test_SafeIndexesPtr(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexesPtr([]string{"a", "b"}, 0, 1)
	emptyResult := stringslice.SafeIndexesPtr(nil, 0)

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"emptyLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexesPtr returns correct value -- returns correct", actual)
}

// ── NonEmptySlicePtr (deprecated) ──

func Test_NonEmptySlicePtr(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlicePtr([]string{"a", "", "b"})
	emptyResult := stringslice.NonEmptySlicePtr(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "NonEmptySlicePtr returns empty -- returns correct", actual)
}

// ── Joins / JoinWith (from CloneIf.go) ──

func Test_JoinWith(t *testing.T) {
	// Arrange
	result := stringslice.JoinWith(",", "a", "b")
	emptyResult := stringslice.JoinWith(",")

	// Act
	actual := args.Map{
		"val": result,
		"empty": emptyResult,
	}

	// Assert
	expected := args.Map{
		"val": ",a,b",
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "JoinWith returns non-empty -- returns correct", actual)
}

func Test_Joins(t *testing.T) {
	// Arrange
	result := stringslice.Joins(",", "a", "b")
	emptyResult := stringslice.Joins(",")

	// Act
	actual := args.Map{
		"val": result,
		"empty": emptyResult,
	}

	// Assert
	expected := args.Map{
		"val": "a,b",
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "Joins returns correct value -- returns correct", actual)
}

// ── SplitContentsByWhitespace ──

func Test_SplitContentsByWhitespace(t *testing.T) {
	// Arrange
	result := stringslice.SplitContentsByWhitespace("  hello  world  ")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "hello",
	}
	expected.ShouldBeEqual(t, 0, "SplitContentsByWhitespace returns correct value -- returns correct", actual)
}

// ── PrependLineNew ──

func Test_PrependLineNew_FromClone(t *testing.T) {
	// Arrange
	result := stringslice.PrependLineNew("first", []string{"second"})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "first",
	}
	expected.ShouldBeEqual(t, 0, "PrependLineNew returns correct value -- returns correct", actual)
}
