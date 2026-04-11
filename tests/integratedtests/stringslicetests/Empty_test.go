package stringslicetests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Empty / EmptyPtr ──

func Test_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.Empty()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- returns empty slice", actual)
}

func Test_EmptyPtr_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.EmptyPtr()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "EmptyPtr returns empty -- returns empty slice", actual)
}

// ── IsEmpty / IsEmptyPtr / HasAnyItem / HasAnyItemPtr ──

func Test_IsEmpty_True_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"empty": stringslice.IsEmpty([]string{})}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- true on empty", actual)
}

func Test_IsEmpty_False_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"empty": stringslice.IsEmpty([]string{"a"})}

	// Assert
	expected := args.Map{"empty": false}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- false on non-empty", actual)
}

func Test_IsEmptyPtr_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"empty": stringslice.IsEmptyPtr(nil)}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns nil -- true on nil", actual)
}

func Test_HasAnyItem_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"has": stringslice.HasAnyItem([]string{"x"})}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem returns non-empty -- true", actual)
}

func Test_HasAnyItemPtr_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"has": stringslice.HasAnyItemPtr([]string{"x"})}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItemPtr returns non-empty -- true", actual)
}

func Test_HasAnyItemPtr_Empty(t *testing.T) {
	// Act
	actual := args.Map{"has": stringslice.HasAnyItemPtr(nil)}

	// Assert
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "HasAnyItemPtr returns nil -- false on nil", actual)
}

// ── LengthOfPointer ──

func Test_LengthOfPointer_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"len": stringslice.LengthOfPointer([]string{"a", "b"})}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfPointer returns correct value -- with args", actual)
}

// ── Make / MakePtr / MakeLen / MakeLenPtr / MakeDefault / MakeDefaultPtr ──

func Test_Make_FromEmpty(t *testing.T) {
	// Arrange
	s := stringslice.Make(3, 5)

	// Act
	actual := args.Map{
		"len": len(s),
		"cap": cap(s),
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"cap": 5,
	}
	expected.ShouldBeEqual(t, 0, "Make returns correct value -- with args", actual)
}

func Test_MakePtr_FromEmpty(t *testing.T) {
	// Arrange
	s := stringslice.MakePtr(3, 5)

	// Act
	actual := args.Map{
		"len": len(s),
		"cap": cap(s),
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"cap": 5,
	}
	expected.ShouldBeEqual(t, 0, "MakePtr returns correct value -- with args", actual)
}

func Test_MakeLen_FromEmpty(t *testing.T) {
	// Arrange
	s := stringslice.MakeLen(4)

	// Act
	actual := args.Map{"len": len(s)}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "MakeLen returns correct value -- with args", actual)
}

func Test_MakeLenPtr_FromEmpty(t *testing.T) {
	// Arrange
	s := stringslice.MakeLenPtr(4)

	// Act
	actual := args.Map{"len": len(s)}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "MakeLenPtr returns correct value -- with args", actual)
}

func Test_MakeDefault_FromEmpty(t *testing.T) {
	// Arrange
	s := stringslice.MakeDefault(10)

	// Act
	actual := args.Map{
		"len": len(s),
		"cap": cap(s),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"cap": 10,
	}
	expected.ShouldBeEqual(t, 0, "MakeDefault returns correct value -- with args", actual)
}

func Test_MakeDefaultPtr_FromEmpty(t *testing.T) {
	// Arrange
	s := stringslice.MakeDefaultPtr(10)

	// Act
	actual := args.Map{
		"len": len(s),
		"cap": cap(s),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"cap": 10,
	}
	expected.ShouldBeEqual(t, 0, "MakeDefaultPtr returns correct value -- with args", actual)
}

// ── Clone / ClonePtr / CloneUsingCap ──

func Test_Clone_NonEmpty_FromEmpty(t *testing.T) {
	// Arrange
	original := []string{"a", "b"}
	cloned := stringslice.Clone(original)
	original[0] = "X"

	// Act
	actual := args.Map{
		"len": len(cloned),
		"first": cloned[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "Clone returns empty -- non-empty", actual)
}

func Test_Clone_Empty_FromEmpty(t *testing.T) {
	// Arrange
	cloned := stringslice.Clone(nil)

	// Act
	actual := args.Map{"len": len(cloned)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clone returns nil -- nil", actual)
}

func Test_ClonePtr_FromEmpty(t *testing.T) {
	// Arrange
	original := []string{"a", "b"}
	cloned := stringslice.ClonePtr(original)
	original[0] = "X"

	// Act
	actual := args.Map{
		"len": len(cloned),
		"first": cloned[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns correct value -- with args", actual)
}

func Test_ClonePtr_Empty_FromEmpty(t *testing.T) {
	// Arrange
	cloned := stringslice.ClonePtr(nil)

	// Act
	actual := args.Map{"len": len(cloned)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_CloneUsingCap_NonEmpty(t *testing.T) {
	// Arrange
	cloned := stringslice.CloneUsingCap(5, []string{"a", "b"})

	// Act
	actual := args.Map{
		"len": len(cloned),
		"capGt": cap(cloned) >= 7,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"capGt": true,
	}
	expected.ShouldBeEqual(t, 0, "CloneUsingCap returns empty -- non-empty", actual)
}

func Test_CloneUsingCap_Empty(t *testing.T) {
	// Arrange
	cloned := stringslice.CloneUsingCap(5, nil)

	// Act
	actual := args.Map{"len": len(cloned)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CloneUsingCap returns empty -- empty", actual)
}

// ── CloneIf ──

func Test_CloneIf_True(t *testing.T) {
	// Arrange
	original := []string{"a"}
	cloned := stringslice.CloneIf(true, 0, original)
	original[0] = "X"

	// Act
	actual := args.Map{"first": cloned[0]}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "CloneIf returns non-empty -- true clones", actual)
}

func Test_CloneIf_False_NonNil(t *testing.T) {
	// Arrange
	original := []string{"a"}
	result := stringslice.CloneIf(false, 0, original)

	// Act
	actual := args.Map{"first": result[0]}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "CloneIf returns non-empty -- false returns same", actual)
}

func Test_CloneIf_False_Nil(t *testing.T) {
	// Arrange
	result := stringslice.CloneIf(false, 0, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CloneIf returns nil -- false nil", actual)
}

// ── CloneSimpleSliceToPointers ──

func Test_CloneSimpleSliceToPointers_NonEmpty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.CloneSimpleSliceToPointers([]string{"a", "b"})

	// Act
	actual := args.Map{"len": len(*result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CloneSimpleSliceToPointers returns empty -- non-empty", actual)
}

func Test_CloneSimpleSliceToPointers_Empty(t *testing.T) {
	// Arrange
	result := stringslice.CloneSimpleSliceToPointers(nil)

	// Act
	actual := args.Map{"len": len(*result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CloneSimpleSliceToPointers returns empty -- empty", actual)
}

// ── JoinWith / Joins ──

func Test_JoinWith_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.JoinWith(", ", "a", "b")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ", a, b"}
	expected.ShouldBeEqual(t, 0, "JoinWith returns empty -- non-empty", actual)
}

func Test_JoinWith_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.JoinWith(", ")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "JoinWith returns empty -- empty", actual)
}

func Test_Joins_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.Joins(", ", "a", "b")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "a, b"}
	expected.ShouldBeEqual(t, 0, "Joins returns empty -- non-empty", actual)
}

func Test_Joins_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.Joins(", ")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "Joins returns empty -- empty", actual)
}

// ── First / FirstPtr / FirstOrDefault / FirstOrDefaultPtr / FirstOrDefaultWith ──

func Test_First_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.First([]string{"a", "b"})}

	// Assert
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "First returns correct value -- with args", actual)
}

func Test_FirstPtr_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.FirstPtr([]string{"a", "b"})}

	// Assert
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "FirstPtr returns correct value -- with args", actual)
}

func Test_FirstOrDefault_NonEmpty_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.FirstOrDefault([]string{"a"})}

	// Assert
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault returns empty -- non-empty", actual)
}

func Test_FirstOrDefault_Empty_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.FirstOrDefault(nil)}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault returns empty -- empty", actual)
}

func Test_FirstOrDefaultPtr_NonEmpty_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.FirstOrDefaultPtr([]string{"a"})}

	// Assert
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultPtr returns empty -- non-empty", actual)
}

func Test_FirstOrDefaultPtr_Empty_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.FirstOrDefaultPtr(nil)}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultPtr returns empty -- empty", actual)
}

func Test_FirstOrDefaultWith_Found(t *testing.T) {
	// Arrange
	result, ok := stringslice.FirstOrDefaultWith([]string{"a"}, "def")

	// Act
	actual := args.Map{
		"val": result,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "a",
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultWith returns non-empty -- found", actual)
}

func Test_FirstOrDefaultWith_NotFound(t *testing.T) {
	// Arrange
	result, ok := stringslice.FirstOrDefaultWith(nil, "def")

	// Act
	actual := args.Map{
		"val": result,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "def",
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultWith returns non-empty -- not found", actual)
}

// ── Last / LastPtr / LastOrDefault / LastOrDefaultPtr ──

func Test_Last_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.Last([]string{"a", "b"})}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "Last returns correct value -- with args", actual)
}

func Test_LastPtr_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.LastPtr([]string{"a", "b"})}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "LastPtr returns correct value -- with args", actual)
}

func Test_LastOrDefault_NonEmpty_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.LastOrDefault([]string{"a", "b"})}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "LastOrDefault returns empty -- non-empty", actual)
}

func Test_LastOrDefault_Empty_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.LastOrDefault(nil)}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "LastOrDefault returns empty -- empty", actual)
}

func Test_LastOrDefaultPtr_NonEmpty_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.LastOrDefaultPtr([]string{"a", "b"})}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "LastOrDefaultPtr returns empty -- non-empty", actual)
}

func Test_LastOrDefaultPtr_Empty_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.LastOrDefaultPtr(nil)}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "LastOrDefaultPtr returns empty -- empty", actual)
}

// ── LastIndexPtr / LastSafeIndexPtr ──

func Test_LastIndexPtr_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"idx": stringslice.LastIndexPtr([]string{"a", "b", "c"})}

	// Assert
	expected := args.Map{"idx": 2}
	expected.ShouldBeEqual(t, 0, "LastIndexPtr returns correct value -- with args", actual)
}

func Test_LastSafeIndexPtr_Valid(t *testing.T) {
	// Act
	actual := args.Map{"idx": stringslice.LastSafeIndexPtr([]string{"a", "b"})}

	// Assert
	expected := args.Map{"idx": 1}
	expected.ShouldBeEqual(t, 0, "LastSafeIndexPtr returns non-empty -- valid", actual)
}

func Test_LastSafeIndexPtr_Empty_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"idx": stringslice.LastSafeIndexPtr(nil)}

	// Assert
	expected := args.Map{"idx": -1}
	expected.ShouldBeEqual(t, 0, "LastSafeIndexPtr returns empty -- empty", actual)
}

// ── FirstLastDefault / FirstLastDefaultPtr ──

func Test_FirstLastDefault_Empty_FromEmpty(t *testing.T) {
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

func Test_FirstLastDefault_One(t *testing.T) {
	// Arrange
	first, last := stringslice.FirstLastDefault([]string{"a"})

	// Act
	actual := args.Map{
		"first": first,
		"last": last,
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"last": "",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault returns correct value -- one", actual)
}

func Test_FirstLastDefault_Two(t *testing.T) {
	// Arrange
	first, last := stringslice.FirstLastDefault([]string{"a", "b"})

	// Act
	actual := args.Map{
		"first": first,
		"last": last,
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault returns correct value -- two", actual)
}

func Test_FirstLastDefaultPtr_Empty_FromEmpty(t *testing.T) {
	// Arrange
	first, last := stringslice.FirstLastDefaultPtr(nil)

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
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultPtr returns empty -- empty", actual)
}

func Test_FirstLastDefaultPtr_NonEmpty_FromEmpty(t *testing.T) {
	// Arrange
	first, last := stringslice.FirstLastDefaultPtr([]string{"a", "b"})

	// Act
	actual := args.Map{
		"first": first,
		"last": last,
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultPtr returns empty -- non-empty", actual)
}

// ── FirstLastDefaultStatus / FirstLastDefaultStatusPtr / FirstLastStatus ──

func Test_FirstLastDefaultStatus_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.FirstLastDefaultStatus(nil)

	// Act
	actual := args.Map{
		"valid": result.IsValid,
		"hasFirst": result.HasFirst,
		"hasLast": result.HasLast,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"hasFirst": false,
		"hasLast": false,
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus returns empty -- empty", actual)
}

func Test_FirstLastDefaultStatus_One(t *testing.T) {
	// Arrange
	result := stringslice.FirstLastDefaultStatus([]string{"a"})

	// Act
	actual := args.Map{
		"valid": result.IsValid,
		"hasFirst": result.HasFirst,
		"hasLast": result.HasLast,
		"first": result.First,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"hasFirst": true,
		"hasLast": false,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus returns correct value -- one", actual)
}

func Test_FirstLastDefaultStatus_Two(t *testing.T) {
	// Arrange
	result := stringslice.FirstLastDefaultStatus([]string{"a", "b"})

	// Act
	actual := args.Map{
		"valid": result.IsValid,
		"hasFirst": result.HasFirst,
		"hasLast": result.HasLast,
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"hasFirst": true,
		"hasLast": true,
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus returns correct value -- two", actual)
}

func Test_FirstLastDefaultStatusPtr_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.FirstLastDefaultStatusPtr(nil)

	// Act
	actual := args.Map{"valid": result.IsValid}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatusPtr returns empty -- empty", actual)
}

func Test_FirstLastDefaultStatusPtr_NonEmpty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.FirstLastDefaultStatusPtr([]string{"a", "b"})

	// Act
	actual := args.Map{"valid": result.IsValid}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatusPtr returns empty -- non-empty", actual)
}

func Test_InvalidFirstLastStatus(t *testing.T) {
	// Arrange
	result := stringslice.InvalidFirstLastStatus()

	// Act
	actual := args.Map{
		"valid": result.IsValid,
		"first": result.First,
		"last": result.Last,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"first": "",
		"last": "",
	}
	expected.ShouldBeEqual(t, 0, "InvalidFirstLastStatus returns error -- with args", actual)
}

func Test_InvalidFirstLastStatusForLast(t *testing.T) {
	// Arrange
	result := stringslice.InvalidFirstLastStatusForLast("x")

	// Act
	actual := args.Map{
		"hasFirst": result.HasFirst,
		"hasLast": result.HasLast,
		"first": result.First,
	}

	// Assert
	expected := args.Map{
		"hasFirst": true,
		"hasLast": false,
		"first": "x",
	}
	expected.ShouldBeEqual(t, 0, "InvalidFirstLastStatusForLast returns error -- with args", actual)
}

// ── IndexAt ──

func Test_IndexAt_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.IndexAt([]string{"a", "b", "c"}, 1)}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "IndexAt returns correct value -- with args", actual)
}

// ── IndexesDefault ──

func Test_IndexesDefault_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.IndexesDefault([]string{"a", "b", "c"}, 0, 2)

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
		"second": "c",
	}
	expected.ShouldBeEqual(t, 0, "IndexesDefault returns empty -- non-empty", actual)
}

func Test_IndexesDefault_Empty(t *testing.T) {
	// Arrange
	result := stringslice.IndexesDefault(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "IndexesDefault returns empty -- empty", actual)
}

// ── InvalidIndexValuesDetail ──

func Test_InvalidIndexValuesDetail(t *testing.T) {
	// Arrange
	result := stringslice.InvalidIndexValuesDetail()

	// Act
	actual := args.Map{
		"valid": result.IsValid,
		"missing": result.IsAnyMissing,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"missing": true,
	}
	expected.ShouldBeEqual(t, 0, "InvalidIndexValuesDetail returns error -- with args", actual)
}

// ── SafeIndexAt / SafeIndexAtWith / SafeIndexAtUsingLastIndex ──

func Test_SafeIndexAt_Valid_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.SafeIndexAt([]string{"a", "b"}, 1)}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns non-empty -- valid", actual)
}

func Test_SafeIndexAt_OutOfRange_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.SafeIndexAt([]string{"a"}, 5)}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns correct value -- out of range", actual)
}

func Test_SafeIndexAt_Negative_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.SafeIndexAt([]string{"a"}, -1)}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns correct value -- negative", actual)
}

func Test_SafeIndexAt_Empty_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.SafeIndexAt(nil, 0)}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns empty -- empty", actual)
}

func Test_SafeIndexAtWith_Valid_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.SafeIndexAtWith([]string{"a", "b"}, 1, "def")}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWith returns non-empty -- valid", actual)
}

func Test_SafeIndexAtWith_OutOfRange(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.SafeIndexAtWith([]string{"a"}, 5, "def")}

	// Assert
	expected := args.Map{"val": "def"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWith returns non-empty -- out of range", actual)
}

func Test_SafeIndexAtWith_Negative(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.SafeIndexAtWith([]string{"a"}, -1, "def")}

	// Assert
	expected := args.Map{"val": "def"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWith returns non-empty -- negative", actual)
}

func Test_SafeIndexAtWithPtr_Valid_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.SafeIndexAtWithPtr([]string{"a", "b"}, 1, "def")}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWithPtr returns non-empty -- valid", actual)
}

func Test_SafeIndexAtWithPtr_OutOfRange_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.SafeIndexAtWithPtr(nil, 0, "def")}

	// Assert
	expected := args.Map{"val": "def"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWithPtr returns non-empty -- out of range", actual)
}

func Test_SafeIndexAtUsingLastIndex_Valid_FromEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.SafeIndexAtUsingLastIndex([]string{"a", "b"}, 1, 0)}

	// Assert
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex returns non-empty -- valid", actual)
}

func Test_SafeIndexAtUsingLastIndex_Invalid(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.SafeIndexAtUsingLastIndex([]string{"a"}, 0, 0)}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex returns correct value -- lastIndex=0", actual)
}

func Test_SafeIndexAtUsingLastIndex_NegativeIndex(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.SafeIndexAtUsingLastIndex([]string{"a"}, 1, -1)}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex returns correct value -- negative index", actual)
}

func Test_SafeIndexAtUsingLastIndex_IndexGtLastIndex(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.SafeIndexAtUsingLastIndex([]string{"a"}, 1, 5)}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex returns correct value -- index > lastIndex", actual)
}

func Test_SafeIndexAtUsingLastIndexPtr_Valid(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.SafeIndexAtUsingLastIndexPtr([]string{"a", "b"}, 1, 0)}

	// Assert
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndexPtr returns non-empty -- valid", actual)
}

func Test_SafeIndexAtUsingLastIndexPtr_LastIndexZero(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.SafeIndexAtUsingLastIndexPtr([]string{"a"}, 0, 0)}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndexPtr returns correct value -- lastIndex=0", actual)
}

func Test_SafeIndexAtUsingLastIndexPtr_NegativeLastIndex(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.SafeIndexAtUsingLastIndexPtr([]string{"a"}, -1, 0)}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndexPtr returns correct value -- negative lastIndex", actual)
}

// ── SafeIndexes / SafeIndexesPtr / SafeIndexesDefaultWithDetail ──

func Test_SafeIndexes_Valid(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexes([]string{"a", "b", "c"}, 0, 2)

	// Act
	actual := args.Map{
		"first": result[0],
		"second": result[1],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"second": "c",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexes returns non-empty -- valid", actual)
}

func Test_SafeIndexes_OutOfRange(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexes([]string{"a"}, 0, 5)

	// Act
	actual := args.Map{
		"first": result[0],
		"second": result[1],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"second": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexes returns correct value -- out of range", actual)
}

func Test_SafeIndexes_EmptySlice_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexes(nil, 0)

	// Act
	actual := args.Map{"first": result[0]}

	// Assert
	expected := args.Map{"first": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexes returns empty -- empty slice", actual)
}

func Test_SafeIndexesPtr_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexesPtr([]string{"a", "b"}, 0, 1)

	// Act
	actual := args.Map{
		"first": result[0],
		"second": result[1],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexesPtr returns empty -- non-empty", actual)
}

func Test_SafeIndexesPtr_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexesPtr(nil, 0)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SafeIndexesPtr returns empty -- empty", actual)
}

func Test_SafeIndexesDefaultWithDetail_Valid(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexesDefaultWithDetail([]string{"a", "b", "c"}, 0, 2)

	// Act
	actual := args.Map{
		"valid": result.IsValid,
		"missing": result.IsAnyMissing,
		"len": len(result.Values),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"missing": false,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexesDefaultWithDetail returns non-empty -- valid", actual)
}

func Test_SafeIndexesDefaultWithDetail_WithMissing(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexesDefaultWithDetail([]string{"a"}, 0, 5)

	// Act
	actual := args.Map{
		"valid": result.IsValid,
		"missing": result.IsAnyMissing,
		"missingLen": len(result.MissingIndexes),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"missing": true,
		"missingLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexesDefaultWithDetail returns non-empty -- with missing", actual)
}

func Test_SafeIndexesDefaultWithDetail_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexesDefaultWithDetail(nil, 0)

	// Act
	actual := args.Map{"valid": result.IsValid}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "SafeIndexesDefaultWithDetail returns empty -- empty", actual)
}

// ── SafeIndexRanges ──

func Test_SafeIndexRanges_Valid(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexRanges([]string{"a", "b", "c", "d"}, 1, 3)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "b",
		"last": "d",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns non-empty -- valid", actual)
}

func Test_SafeIndexRanges_NegativeRange_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexRanges([]string{"a"}, 5, 2)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns correct value -- negative range", actual)
}

func Test_SafeIndexRanges_OutOfRange(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexRanges([]string{"a", "b"}, 0, 5)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"second": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 6,
		"first": "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns empty -- out of range pads empty", actual)
}

func Test_SafeIndexRanges_EmptySlice(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexRanges(nil, 0, 0)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns empty -- empty slice", actual)
}

// ── SafeRangeItems / SafeRangeItemsPtr ──

func Test_SafeRangeItems_Valid(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems([]string{"a", "b", "c"}, 0, 2)

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
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns non-empty -- valid", actual)
}

func Test_SafeRangeItems_Nil(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems(nil, 0, 2)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns nil -- nil", actual)
}

func Test_SafeRangeItems_Empty(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems([]string{}, 0, 2)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns empty -- empty", actual)
}

func Test_SafeRangeItems_StartGtLastIndex(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems([]string{"a"}, 5, 10)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns correct value -- start > lastIndex", actual)
}

func Test_SafeRangeItems_EndGtLastIndex(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems([]string{"a", "b", "c"}, 0, -1)

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
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns correct value -- end=-1 clips to lastIndex", actual)
}

func Test_SafeRangeItems_StartInvalid(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems([]string{"a", "b", "c"}, -1, 2)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns correct value -- start=-1 uses [:end]", actual)
}

func Test_SafeRangeItemsPtr_NonEmpty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItemsPtr([]string{"a", "b"}, 0, 1)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SafeRangeItemsPtr returns empty -- non-empty", actual)
}

func Test_SafeRangeItemsPtr_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItemsPtr(nil, 0, 1)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeRangeItemsPtr returns empty -- empty", actual)
}

// ── SlicePtr ──

func Test_SlicePtr_NonEmpty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.SlicePtr([]string{"a"})

	// Act
	actual := args.Map{"first": result[0]}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "SlicePtr returns empty -- non-empty", actual)
}

func Test_SlicePtr_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.SlicePtr(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SlicePtr returns empty -- empty", actual)
}

// ── InPlaceReverse ──

func Test_InPlaceReverse_Nil_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.InPlaceReverse(nil)

	// Act
	actual := args.Map{"len": len(*result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns nil -- nil", actual)
}

func Test_InPlaceReverse_Single_FromEmpty(t *testing.T) {
	// Arrange
	s := []string{"a"}
	result := stringslice.InPlaceReverse(&s)

	// Act
	actual := args.Map{"first": (*result)[0]}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns correct value -- single", actual)
}

func Test_InPlaceReverse_Two_FromEmpty(t *testing.T) {
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

func Test_InPlaceReverse_Three_FromEmpty(t *testing.T) {
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

// ── MergeNew / MergeNewSimple / MergeSlicesOfSlices ──

func Test_MergeNew_FromEmpty(t *testing.T) {
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

func Test_MergeNew_EmptyFirst_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MergeNew(nil, "b")

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

func Test_MergeNew_EmptyAdditional(t *testing.T) {
	// Arrange
	result := stringslice.MergeNew([]string{"a"})

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
	expected.ShouldBeEqual(t, 0, "MergeNew returns empty -- empty additional", actual)
}

func Test_MergeNewSimple_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MergeNewSimple([]string{"a"}, []string{"b"}, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple returns correct value -- with args", actual)
}

func Test_MergeNewSimple_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MergeNewSimple()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple returns empty -- empty", actual)
}

func Test_MergeSlicesOfSlices_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MergeSlicesOfSlices([]string{"a"}, nil, []string{"b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices returns correct value -- with args", actual)
}

func Test_MergeSlicesOfSlices_Empty(t *testing.T) {
	// Arrange
	result := stringslice.MergeSlicesOfSlices()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices returns empty -- empty", actual)
}

func Test_AllElemLengthSlices_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.AllElemLengthSlices([]string{"a", "b"}, nil, []string{"c"})

	// Act
	actual := args.Map{"count": result}

	// Assert
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "AllElemLengthSlices returns correct value -- with args", actual)
}

func Test_AllElemLengthSlices_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.AllElemLengthSlices()

	// Act
	actual := args.Map{"count": result}

	// Assert
	expected := args.Map{"count": 0}
	expected.ShouldBeEqual(t, 0, "AllElemLengthSlices returns empty -- empty", actual)
}

// ── AppendLineNew / PrependLineNew / PrependNew ──

func Test_AppendLineNew_FromEmpty(t *testing.T) {
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

func Test_PrependLineNew_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.PrependLineNew("x", []string{"a"})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "x",
	}
	expected.ShouldBeEqual(t, 0, "PrependLineNew returns correct value -- with args", actual)
}

func Test_PrependNew_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.PrependNew([]string{"a"}, "x", "y")

	// Act
	actual := args.Map{
		"len": len(*result),
		"first": (*result)[0],
		"last": (*result)[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "x",
		"last": "a",
	}
	expected.ShouldBeEqual(t, 0, "PrependNew returns correct value -- with args", actual)
}

func Test_PrependNew_EmptyPrepend(t *testing.T) {
	// Arrange
	result := stringslice.PrependNew([]string{"a"})

	// Act
	actual := args.Map{
		"len": len(*result),
		"first": (*result)[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "PrependNew returns empty -- empty prepend", actual)
}

func Test_PrependNew_EmptySlice(t *testing.T) {
	// Arrange
	result := stringslice.PrependNew(nil, "x")

	// Act
	actual := args.Map{
		"len": len(*result),
		"first": (*result)[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "x",
	}
	expected.ShouldBeEqual(t, 0, "PrependNew returns empty -- empty slice", actual)
}

// ── AppendAnyItemsWithStrings ──

func Test_AppendAnyItemsWithStrings_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.AppendAnyItemsWithStrings(false, false, []string{"a"}, "b", 42)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings returns non-empty -- with args", actual)
}

func Test_AppendAnyItemsWithStrings_SkipEmpty(t *testing.T) {
	// Arrange
	result := stringslice.AppendAnyItemsWithStrings(false, true, []string{"a"}, nil, "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings returns nil -- skip nil", actual)
}

func Test_AppendAnyItemsWithStrings_NoAdditional(t *testing.T) {
	// Arrange
	result := stringslice.AppendAnyItemsWithStrings(false, false, []string{"a"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings returns empty -- no additional", actual)
}

// ── AppendStringsWithAnyItems ──

func Test_AppendStringsWithAnyItems_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.AppendStringsWithAnyItems(false, false, []any{"a"}, "b", "c")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems returns non-empty -- with args", actual)
}

func Test_AppendStringsWithAnyItems_SkipEmpty(t *testing.T) {
	// Arrange
	result := stringslice.AppendStringsWithAnyItems(false, true, []any{"a"}, "", "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems returns empty -- skip empty", actual)
}

func Test_AppendStringsWithAnyItems_NoAdditional(t *testing.T) {
	// Arrange
	result := stringslice.AppendStringsWithAnyItems(false, false, []any{"a"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems returns empty -- no additional", actual)
}

// ── AppendStringsWithMainSlice ──

func Test_AppendStringsWithMainSlice_FromEmpty(t *testing.T) {
	// Arrange
	main := []string{"a"}
	result := stringslice.AppendStringsWithMainSlice(false, main, "b", "c")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice returns non-empty -- with args", actual)
}

func Test_AppendStringsWithMainSlice_SkipEmpty_FromEmpty(t *testing.T) {
	// Arrange
	main := []string{"a"}
	result := stringslice.AppendStringsWithMainSlice(true, main, "", "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice returns empty -- skip empty", actual)
}

func Test_AppendStringsWithMainSlice_NoAdditional(t *testing.T) {
	// Arrange
	main := []string{"a"}
	result := stringslice.AppendStringsWithMainSlice(false, main)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice returns empty -- no additional", actual)
}

// ── AnyItemsCloneIf / AnyItemsCloneUsingCap ──

func Test_AnyItemsCloneIf_True(t *testing.T) {
	// Arrange
	result := stringslice.AnyItemsCloneIf(true, 0, []any{"a"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf returns non-empty -- true", actual)
}

func Test_AnyItemsCloneIf_False_NonNil(t *testing.T) {
	// Arrange
	original := []any{"a"}
	result := stringslice.AnyItemsCloneIf(false, 0, original)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf returns nil -- false non-nil", actual)
}

func Test_AnyItemsCloneIf_False_Nil(t *testing.T) {
	// Arrange
	result := stringslice.AnyItemsCloneIf(false, 0, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf returns nil -- false nil", actual)
}

func Test_AnyItemsCloneUsingCap_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.AnyItemsCloneUsingCap(5, []any{"a"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneUsingCap returns empty -- non-empty", actual)
}

func Test_AnyItemsCloneUsingCap_Empty(t *testing.T) {
	// Arrange
	result := stringslice.AnyItemsCloneUsingCap(5, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneUsingCap returns empty -- empty", actual)
}

// ── NonEmpty / NonEmptySlicePtr / NonEmptyStrings / NonNullStrings ──

func Test_NonEmptySlice_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlice([]string{"a", "", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice returns empty -- with args", actual)
}

func Test_NonEmptySlice_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlice(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice returns empty -- empty", actual)
}

func Test_NonEmptySlicePtr_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlicePtr([]string{"a", "", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptySlicePtr returns empty -- with args", actual)
}

func Test_NonEmptySlicePtr_Empty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlicePtr(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptySlicePtr returns empty -- empty", actual)
}

func Test_NonEmptyStrings_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyStrings([]string{"a", "", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings returns empty -- with args", actual)
}

func Test_NonEmptyStrings_Nil_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyStrings(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings returns nil -- nil", actual)
}

func Test_NonEmptyStrings_Empty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyStrings([]string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings returns empty -- empty", actual)
}

func Test_NonNullStrings_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonNullStrings([]string{"a", "", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonNullStrings returns correct value -- with args", actual)
}

func Test_NonNullStrings_Nil_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonNullStrings(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonNullStrings returns nil -- nil", actual)
}

// ── NonEmptyIf ──

func Test_NonEmptyIf_True(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyIf(true, []string{"a", "", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptyIf returns empty -- true", actual)
}

func Test_NonEmptyIf_False(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyIf(false, []string{"a", "", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptyIf returns empty -- false uses NonNullStrings", actual)
}

// ── NonEmptyJoin / NonEmptyJoinPtr ──

func Test_NonEmptyJoin_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoin([]string{"a", "", "b"}, ", ")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "a, b"}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- with args", actual)
}

func Test_NonEmptyJoin_Nil_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoin(nil, ", ")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns nil -- nil", actual)
}

func Test_NonEmptyJoin_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoin([]string{}, ", ")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- empty", actual)
}

func Test_NonEmptyJoinPtr_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoinPtr([]string{"a", "", "b"}, ", ")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "a, b"}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoinPtr returns empty -- with args", actual)
}

func Test_NonEmptyJoinPtr_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoinPtr(nil, ", ")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoinPtr returns empty -- empty", actual)
}

// ── NonWhitespace / NonWhitespacePtr / NonWhitespaceJoin / NonWhitespaceJoinPtr ──

func Test_NonWhitespace_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespace([]string{"a", "  ", "", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonWhitespace returns correct value -- with args", actual)
}

func Test_NonWhitespace_Nil_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespace(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespace returns nil -- nil", actual)
}

func Test_NonWhitespace_Empty(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespace([]string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespace returns empty -- empty", actual)
}

func Test_NonWhitespacePtr(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespacePtr([]string{"a", "  ", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonWhitespacePtr returns correct value -- with args", actual)
}

func Test_NonWhitespacePtr_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespacePtr(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespacePtr returns empty -- empty", actual)
}

func Test_NonWhitespaceJoin_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespaceJoin([]string{"a", "  ", "b"}, ", ")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "a, b"}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns correct value -- with args", actual)
}

func Test_NonWhitespaceJoin_Nil_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespaceJoin(nil, ", ")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns nil -- nil", actual)
}

func Test_NonWhitespaceJoin_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespaceJoin([]string{}, ", ")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns empty -- empty", actual)
}

func Test_NonWhitespaceJoinPtr_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespaceJoinPtr([]string{"a", "  ", "b"}, ", ")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "a, b"}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoinPtr returns correct value -- with args", actual)
}

func Test_NonWhitespaceJoinPtr_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespaceJoinPtr(nil, ", ")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoinPtr returns empty -- empty", actual)
}

// ── TrimmedEachWords / TrimmedEachWordsPtr / TrimmedEachWordsIf ──

func Test_TrimmedEachWords_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWords([]string{"  a  ", " ", "  b  "})

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
	expected.ShouldBeEqual(t, 0, "TrimmedEachWords returns correct value -- with args", actual)
}

func Test_TrimmedEachWords_Nil(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWords(nil)
	isNil := result == nil

	// Act
	actual := args.Map{"isNil": isNil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWords returns nil -- nil returns nil", actual)
}

func Test_TrimmedEachWords_Empty(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWords([]string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWords returns empty -- empty", actual)
}

func Test_TrimmedEachWordsPtr_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWordsPtr([]string{"  a  "})

	// Act
	actual := args.Map{"first": result[0]}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsPtr returns correct value -- with args", actual)
}

func Test_TrimmedEachWordsPtr_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWordsPtr(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsPtr returns empty -- empty", actual)
}

func Test_TrimmedEachWordsIf_True(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWordsIf(true, []string{"  a  ", " ", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsIf returns non-empty -- true", actual)
}

func Test_TrimmedEachWordsIf_False(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWordsIf(false, []string{"a", "", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsIf returns non-empty -- false uses NonNullStrings", actual)
}

// ── SortIf ──

func Test_SortIf_True_FromEmpty(t *testing.T) {
	// Arrange
	s := []string{"c", "a", "b"}
	result := stringslice.SortIf(true, s)

	// Act
	actual := args.Map{"first": result[0]}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "SortIf returns non-empty -- true sorts", actual)
}

func Test_SortIf_False_FromEmpty(t *testing.T) {
	// Arrange
	s := []string{"c", "a", "b"}
	result := stringslice.SortIf(false, s)

	// Act
	actual := args.Map{"first": result[0]}

	// Assert
	expected := args.Map{"first": "c"}
	expected.ShouldBeEqual(t, 0, "SortIf returns empty -- false no sort", actual)
}

// ── SplitContentsByWhitespace ──

func Test_SplitContentsByWhitespace_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.SplitContentsByWhitespace("hello  world   test")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "hello",
		"last": "test",
	}
	expected.ShouldBeEqual(t, 0, "SplitContentsByWhitespace returns correct value -- with args", actual)
}

// ── SplitTrimmedNonEmpty / SplitTrimmedNonEmptyAll ──

func Test_SplitTrimmedNonEmpty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.SplitTrimmedNonEmpty(" a , b , c ", ",", -1)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmpty returns empty -- with args", actual)
}

func Test_SplitTrimmedNonEmptyAll_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.SplitTrimmedNonEmptyAll(" a , b , c ", ",")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmptyAll returns empty -- with args", actual)
}

// ── RegexTrimmedSplitNonEmptyAll ──

func Test_RegexTrimmedSplitNonEmptyAll_FromEmpty(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`[,;]`)
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, " a , b ; c ")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "RegexTrimmedSplitNonEmptyAll returns empty -- with args", actual)
}

// ── ExpandByFunc / ExpandBySplit / ExpandBySplits ──

func Test_ExpandByFunc_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.ExpandByFunc([]string{"a,b", "c,d"}, func(line string) []string {
		return stringslice.SplitTrimmedNonEmptyAll(line, ",")
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc returns correct value -- with args", actual)
}

func Test_ExpandByFunc_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.ExpandByFunc(nil, func(line string) []string { return nil })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc returns empty -- empty", actual)
}

func Test_ExpandBySplit_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplit([]string{"a,b", "c"}, ",")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit returns correct value -- with args", actual)
}

func Test_ExpandBySplit_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplit(nil, ",")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit returns empty -- empty", actual)
}

func Test_ExpandBySplits(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplits([]string{"a,b;c"}, ",", ";")

	// Act
	actual := args.Map{"lenGt": len(result) > 0}

	// Assert
	expected := args.Map{"lenGt": true}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits returns correct value -- with args", actual)
}

func Test_ExpandBySplits_Empty(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplits(nil, ",")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits returns empty -- empty", actual)
}

func Test_ExpandBySplits_NoSplitters_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplits([]string{"a"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits returns empty -- no splitters", actual)
}

// ── ProcessAsync / ProcessOptionAsync ──

func Test_ProcessAsync(t *testing.T) {
	// Arrange
	result := stringslice.ProcessAsync(func(index int, item any) string {
		return "x"
	}, "a", "b")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "x",
	}
	expected.ShouldBeEqual(t, 0, "ProcessAsync returns correct value -- with args", actual)
}

func Test_ProcessAsync_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.ProcessAsync(func(index int, item any) string { return "" })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ProcessAsync returns empty -- empty", actual)
}

func Test_ProcessOptionAsync_SkipEmpty(t *testing.T) {
	// Arrange
	result := stringslice.ProcessOptionAsync(true, func(index int, item any) string {
		if index == 0 {
			return ""
		}
		return "x"
	}, "a", "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync returns empty -- skip empty", actual)
}

func Test_ProcessOptionAsync_ReturnAll(t *testing.T) {
	// Arrange
	result := stringslice.ProcessOptionAsync(false, func(index int, item any) string {
		if index == 0 {
			return ""
		}
		return "x"
	}, "a", "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync returns correct value -- return all", actual)
}

func Test_ProcessOptionAsync_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.ProcessOptionAsync(true, func(index int, item any) string { return "" })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync returns empty -- empty", actual)
}

// ── LinesProcess / LinesSimpleProcess / LinesSimpleProcessNoEmpty / LinesAsyncProcess ──

func Test_LinesProcess_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.LinesProcess([]string{"a", "b", "c"}, func(i int, lineIn string) (string, bool, bool) {
		return lineIn + "!", true, false
	})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "a!",
	}
	expected.ShouldBeEqual(t, 0, "LinesProcess returns correct value -- with args", actual)
}

func Test_LinesProcess_WithBreak(t *testing.T) {
	// Arrange
	result := stringslice.LinesProcess([]string{"a", "b", "c"}, func(i int, lineIn string) (string, bool, bool) {
		if i == 1 {
			return "", false, true
		}
		return lineIn, true, false
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LinesProcess returns non-empty -- with break", actual)
}

func Test_LinesProcess_SkipItem(t *testing.T) {
	// Arrange
	result := stringslice.LinesProcess([]string{"a", "b"}, func(i int, lineIn string) (string, bool, bool) {
		return lineIn, i == 0, false
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LinesProcess returns correct value -- skip item", actual)
}

func Test_LinesProcess_Empty(t *testing.T) {
	// Arrange
	result := stringslice.LinesProcess(nil, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesProcess returns empty -- empty", actual)
}

func Test_LinesSimpleProcess_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.LinesSimpleProcess([]string{"a", "b"}, func(lineIn string) string {
		return lineIn + "!"
	})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a!",
	}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcess returns correct value -- with args", actual)
}

func Test_LinesSimpleProcess_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.LinesSimpleProcess(nil, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcess returns empty -- empty", actual)
}

func Test_LinesSimpleProcessNoEmpty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.LinesSimpleProcessNoEmpty([]string{"a", "b"}, func(lineIn string) string {
		if lineIn == "b" {
			return ""
		}
		return lineIn
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcessNoEmpty returns empty -- with args", actual)
}

func Test_LinesSimpleProcessNoEmpty_Empty(t *testing.T) {
	// Arrange
	result := stringslice.LinesSimpleProcessNoEmpty(nil, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcessNoEmpty returns empty -- empty", actual)
}

func Test_LinesAsyncProcess(t *testing.T) {
	// Arrange
	result := stringslice.LinesAsyncProcess([]string{"a", "b"}, func(i int, lineIn string) string {
		return lineIn + "!"
	})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a!",
	}
	expected.ShouldBeEqual(t, 0, "LinesAsyncProcess returns correct value -- with args", actual)
}

func Test_LinesAsyncProcess_Empty_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.LinesAsyncProcess(nil, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesAsyncProcess returns empty -- empty", actual)
}

// ── AnyLinesProcessAsyncUsingProcessor ──

func Test_AnyLinesProcessAsync_Slice(t *testing.T) {
	// Arrange
	result := stringslice.AnyLinesProcessAsyncUsingProcessor(
		[]string{"a", "b"},
		func(i int, lineIn any) string { return "x" },
	)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns correct value -- slice", actual)
}

func Test_AnyLinesProcessAsync_Nil_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.AnyLinesProcessAsyncUsingProcessor(nil, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns nil -- nil", actual)
}

func Test_AnyLinesProcessAsync_NotSlice_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.AnyLinesProcessAsyncUsingProcessor("hello", nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns correct value -- not slice", actual)
}

func Test_AnyLinesProcessAsync_EmptySlice_FromEmpty(t *testing.T) {
	// Arrange
	result := stringslice.AnyLinesProcessAsyncUsingProcessor(
		[]string{},
		func(i int, lineIn any) string { return "x" },
	)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns empty -- empty slice", actual)
}

func Test_AnyLinesProcessAsync_Array(t *testing.T) {
	// Arrange
	result := stringslice.AnyLinesProcessAsyncUsingProcessor(
		[2]string{"a", "b"},
		func(i int, lineIn any) string { return "x" },
	)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns correct value -- array", actual)
}
