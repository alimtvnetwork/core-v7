package stringslicetests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ============================================================================
// NonEmptySlice
// ============================================================================

func Test_NonEmptySlice_Mixed_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlice([]string{"a", "", "b", ""})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice filters empty -- mixed", actual)
}

func Test_NonEmptySlice_Empty_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlice([]string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice returns empty -- empty input", actual)
}

// ============================================================================
// NonEmptyStrings
// ============================================================================

func Test_NonEmptyStrings_Mixed(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyStrings([]string{"a", "", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings filters empty -- mixed", actual)
}

func Test_NonEmptyStrings_Nil_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyStrings(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings returns empty -- nil", actual)
}

func Test_NonEmptyStrings_Empty_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyStrings([]string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings returns empty -- empty", actual)
}

// ============================================================================
// NonWhitespace
// ============================================================================

func Test_NonWhitespace_Mixed(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespace([]string{"a", "  ", "b", ""})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonWhitespace filters ws and empty -- mixed", actual)
}

func Test_NonWhitespace_Nil_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespace(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespace returns empty -- nil", actual)
}

func Test_NonWhitespace_Empty_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespace([]string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespace returns empty -- empty", actual)
}

// ============================================================================
// Clone
// ============================================================================

func Test_Clone_NonEmpty_FromNonEmptySliceMixed(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "Clone copies slice -- non-empty", actual)
}

func Test_Clone_Empty_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.Clone([]string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clone returns empty -- empty", actual)
}

// ============================================================================
// FirstLastDefault
// ============================================================================

func Test_FirstLastDefault_Empty_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	first, last := stringslice.FirstLastDefault([]string{})

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

func Test_FirstLastDefault_Single_FromNonEmptySliceMixed(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "FirstLastDefault returns first only -- single", actual)
}

func Test_FirstLastDefault_Multiple(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "FirstLastDefault returns first and last -- multiple", actual)
}

// ============================================================================
// FirstLastDefaultStatus
// ============================================================================

func Test_FirstLastDefaultStatus_Empty_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	s := stringslice.FirstLastDefaultStatus([]string{})

	// Act
	actual := args.Map{
		"valid": s.IsValid,
		"hasFirst": s.HasFirst,
		"hasLast": s.HasLast,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"hasFirst": false,
		"hasLast": false,
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus returns invalid -- empty", actual)
}

func Test_FirstLastDefaultStatus_Single_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	s := stringslice.FirstLastDefaultStatus([]string{"only"})

	// Act
	actual := args.Map{
		"valid": s.IsValid,
		"hasFirst": s.HasFirst,
		"hasLast": s.HasLast,
		"first": s.First,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"hasFirst": true,
		"hasLast": false,
		"first": "only",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus returns partial -- single", actual)
}

func Test_FirstLastDefaultStatus_Multiple_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	s := stringslice.FirstLastDefaultStatus([]string{"a", "b"})

	// Act
	actual := args.Map{
		"valid": s.IsValid,
		"hasFirst": s.HasFirst,
		"hasLast": s.HasLast,
		"first": s.First,
		"last": s.Last,
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"hasFirst": true,
		"hasLast": true,
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus returns valid -- multiple", actual)
}

// ============================================================================
// SafeIndexAt
// ============================================================================

func Test_SafeIndexAt_Valid_FromNonEmptySliceMixed(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.SafeIndexAt([]string{"a", "b"}, 1)}

	// Assert
	expected := args.Map{"result": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns element -- valid index", actual)
}

func Test_SafeIndexAt_OutOfBounds(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.SafeIndexAt([]string{"a"}, 5)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns empty -- out of bounds", actual)
}

func Test_SafeIndexAt_Negative_FromNonEmptySliceMixed(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.SafeIndexAt([]string{"a"}, -1)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns empty -- negative index", actual)
}

func Test_SafeIndexAt_Empty_FromNonEmptySliceMixed(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.SafeIndexAt([]string{}, 0)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns empty -- empty slice", actual)
}

// ============================================================================
// SafeIndexAtWith
// ============================================================================

func Test_SafeIndexAtWith_Valid_FromNonEmptySliceMixed(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.SafeIndexAtWith([]string{"a", "b"}, 1, "def")}

	// Assert
	expected := args.Map{"result": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWith returns element -- valid index", actual)
}

func Test_SafeIndexAtWith_Default(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.SafeIndexAtWith([]string{"a"}, 5, "def")}

	// Assert
	expected := args.Map{"result": "def"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWith returns default -- out of bounds", actual)
}

// ============================================================================
// NonEmptyJoin
// ============================================================================

func Test_NonEmptyJoin_Mixed_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoin([]string{"a", "", "b"}, ",")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin filters and joins -- mixed", actual)
}

func Test_NonEmptyJoin_Nil_FromNonEmptySliceMixed(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.NonEmptyJoin(nil, ",")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- nil", actual)
}

func Test_NonEmptyJoin_Empty_FromNonEmptySliceMixed(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.NonEmptyJoin([]string{}, ",")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- empty", actual)
}

// ============================================================================
// NonWhitespaceJoin
// ============================================================================

func Test_NonWhitespaceJoin_Mixed_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespaceJoin([]string{"a", "  ", "b"}, ",")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin filters ws and joins -- mixed", actual)
}

func Test_NonWhitespaceJoin_Nil_FromNonEmptySliceMixed(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.NonWhitespaceJoin(nil, ",")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns empty -- nil", actual)
}

func Test_NonWhitespaceJoin_Empty_FromNonEmptySliceMixed(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.NonWhitespaceJoin([]string{}, ",")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns empty -- empty", actual)
}

// ============================================================================
// PrependNew / PrependLineNew
// ============================================================================

func Test_PrependNew_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.PrependNew([]string{"b", "c"}, "a")

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
	expected.ShouldBeEqual(t, 0, "PrependNew prepends items -- one prepend", actual)
}

func Test_PrependNew_Empty(t *testing.T) {
	// Arrange
	result := stringslice.PrependNew(nil)

	// Act
	actual := args.Map{"len": len(*result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "PrependNew returns empty -- nil both", actual)
}

func Test_PrependLineNew_FromNonEmptySliceMixed(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "PrependLineNew prepends line -- one line", actual)
}

// ============================================================================
// TrimmedEachWords
// ============================================================================

func Test_TrimmedEachWords_Mixed(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWords([]string{"  a  ", "  ", " b "})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWords trims and filters -- mixed", actual)
}

func Test_TrimmedEachWords_Nil_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWords(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWords returns empty -- nil", actual)
}

func Test_TrimmedEachWords_Empty_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWords([]string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWords returns empty -- empty", actual)
}

// ============================================================================
// ExpandByFunc
// ============================================================================

func Test_ExpandByFunc_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.ExpandByFunc([]string{"a,b", "c"}, func(line string) []string {
		return strings.Split(line, ",")
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc expands -- comma split", actual)
}

func Test_ExpandByFunc_Empty_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.ExpandByFunc([]string{}, func(line string) []string { return nil })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc returns empty -- empty input", actual)
}

func Test_ExpandByFunc_NilReturn(t *testing.T) {
	// Arrange
	result := stringslice.ExpandByFunc([]string{"a"}, func(line string) []string { return nil })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc skips nil returns -- nil expand", actual)
}

// ============================================================================
// ExpandBySplits
// ============================================================================

func Test_ExpandBySplits_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplits([]string{"a,b"}, ",", ";")

	// Act
	actual := args.Map{"hasItems": len(result) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits expands -- comma and semi", actual)
}

func Test_ExpandBySplits_Empty_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplits([]string{}, ",")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits returns empty -- empty input", actual)
}

// ============================================================================
// SplitTrimmedNonEmpty
// ============================================================================

func Test_SplitTrimmedNonEmpty_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.SplitTrimmedNonEmpty(" a , b , ", ",", -1)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmpty trims and filters -- comma", actual)
}

// ============================================================================
// SplitContentsByWhitespace
// ============================================================================

func Test_SplitContentsByWhitespace_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.SplitContentsByWhitespace("  hello  world  ")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "hello",
		"last": "world",
	}
	expected.ShouldBeEqual(t, 0, "SplitContentsByWhitespace splits by ws -- two words", actual)
}

// ============================================================================
// NonNullStrings
// ============================================================================

func Test_NonNullStrings_Nil_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.NonNullStrings(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonNullStrings returns empty -- nil", actual)
}

func Test_NonNullStrings_NonNil(t *testing.T) {
	// Arrange
	result := stringslice.NonNullStrings([]string{"a"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NonNullStrings passes through -- non-nil", actual)
}

// ============================================================================
// CloneIf — JoinWith / Joins
// ============================================================================

func Test_JoinWith_NonEmpty_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.JoinWith(",", "a", "b")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ",a,b"}
	expected.ShouldBeEqual(t, 0, "JoinWith prepends joiner -- comma", actual)
}

func Test_JoinWith_Empty_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.JoinWith(",")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "JoinWith returns empty -- no items", actual)
}

func Test_Joins_NonEmpty_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.Joins(",", "a", "b")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "Joins joins items -- comma", actual)
}

func Test_Joins_Empty_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.Joins(",")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Joins returns empty -- no items", actual)
}

// ============================================================================
// CloneIf — more branches
// ============================================================================

func Test_CloneIf_NilNoClone_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.CloneIf(false, 0, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CloneIf nil no-clone returns empty -- nil+false", actual)
}

func Test_CloneIf_Clone_FromNonEmptySliceMixed(t *testing.T) {
	// Arrange
	result := stringslice.CloneIf(true, 5, []string{"a", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CloneIf clone returns copy -- true", actual)
}

// ============================================================================
// NonEmptySlicePtr
// ============================================================================

func Test_NonEmptySlicePtr_Mixed(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlicePtr([]string{"a", "", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptySlicePtr filters empty -- mixed", actual)
}

// ============================================================================
// NonEmptySlice — extra
// ============================================================================

func Test_NonEmptySlice_AllEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlice([]string{"", ""})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice returns empty -- all empty", actual)
}

// ============================================================================
// FirstOrDefault / FirstOrDefaultPtr / LastOrDefault / LastOrDefaultPtr
// ============================================================================

func Test_FirstOrDefault_NonEmpty_FromNonEmptySliceMixed(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.FirstOrDefault([]string{"a", "b"})}

	// Assert
	expected := args.Map{"result": "a"}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault returns first -- non-empty", actual)
}

func Test_FirstOrDefault_Empty_FromNonEmptySliceMixed(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.FirstOrDefault(nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault returns empty -- nil", actual)
}

func Test_LastOrDefault_NonEmpty_FromNonEmptySliceMixed(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.LastOrDefault([]string{"a", "b"})}

	// Assert
	expected := args.Map{"result": "b"}
	expected.ShouldBeEqual(t, 0, "LastOrDefault returns last -- non-empty", actual)
}

func Test_LastOrDefault_Empty_FromNonEmptySliceMixed(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.LastOrDefault(nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "LastOrDefault returns empty -- nil", actual)
}
