package corestrtests

import (
	"errors"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Segment 1: Basic methods, Add variants, Remove, Capacity, Equals
// Covers ~200 uncovered statements from Collection.go lines 27-700
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovP1_01_JsonString(t *testing.T) {
	safeTest(t, "Test_CovP1_01_JsonString", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		s := col.JsonString()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty JSON string", actual)
	})
}

func Test_CovP1_02_JsonStringMust(t *testing.T) {
	safeTest(t, "Test_CovP1_02_JsonStringMust", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"x"})
		s := col.JsonStringMust()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty JSON string", actual)
	})
}

func Test_CovP1_03_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_CovP1_03_HasAnyItem", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": col.HasAnyItem()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		empty := corestr.Empty.Collection()
		actual = args.Map{"result": empty.HasAnyItem()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovP1_04_LastIndex(t *testing.T) {
	safeTest(t, "Test_CovP1_04_LastIndex", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"result": col.LastIndex() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP1_05_HasIndex(t *testing.T) {
	safeTest(t, "Test_CovP1_05_HasIndex", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": col.HasIndex(0)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for index 0", actual)
		actual = args.Map{"result": col.HasIndex(1)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for index 1", actual)
		actual = args.Map{"result": col.HasIndex(2)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for index 2", actual)
		actual = args.Map{"result": col.HasIndex(-1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for -1", actual)
	})
}

func Test_CovP1_06_ListStringsPtr(t *testing.T) {
	safeTest(t, "Test_CovP1_06_ListStringsPtr", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": len(col.ListStringsPtr()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovP1_07_ListStrings(t *testing.T) {
	safeTest(t, "Test_CovP1_07_ListStrings", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"x", "y"})

		// Act
		actual := args.Map{"result": len(col.ListStrings()) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP1_08_StringJSON(t *testing.T) {
	safeTest(t, "Test_CovP1_08_StringJSON", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": col.StringJSON() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_CovP1_09_RemoveAt(t *testing.T) {
	safeTest(t, "Test_CovP1_09_RemoveAt", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok := col.RemoveAt(1)

		// Act
		actual := args.Map{"result": ok}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected success", actual)
		actual = args.Map{"result": col.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// out of range
		ok = col.RemoveAt(-1)
		actual = args.Map{"result": ok}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for -1", actual)
		ok = col.RemoveAt(100)
		actual = args.Map{"result": ok}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for 100", actual)
	})
}

func Test_CovP1_10_Count(t *testing.T) {
	safeTest(t, "Test_CovP1_10_Count", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": col.Count() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovP1_11_Capacity(t *testing.T) {
	safeTest(t, "Test_CovP1_11_Capacity", func() {
		// Arrange
		col := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"result": col.Capacity() < 10}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 10", actual)
		empty := corestr.Empty.Collection()
		_ = empty.Capacity()
	})
}

func Test_CovP1_12_Length(t *testing.T) {
	safeTest(t, "Test_CovP1_12_Length", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": col.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP1_13_LengthLock(t *testing.T) {
	safeTest(t, "Test_CovP1_13_LengthLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": col.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovP1_14_IsCollectionPrecheckEqual_AllPaths(t *testing.T) {
	safeTest(t, "Test_CovP1_14_IsCollectionPrecheckEqual_AllPaths", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})

		// same content

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)

		// same pointer
		actual = args.Map{"result": a.IsEquals(a)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal to self", actual)

		// both empty
		e1 := corestr.Empty.Collection()
		e2 := corestr.Empty.Collection()
		actual = args.Map{"result": e1.IsEquals(e2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty equals empty", actual)

		// one empty
		actual = args.Map{"result": a.IsEquals(e1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)

		// diff length
		c := corestr.New.Collection.Strings([]string{"a"})
		actual = args.Map{"result": a.IsEquals(c)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal for diff length", actual)
	})
}

func Test_CovP1_15_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_CovP1_15_IsEqualsWithSensitive_CaseInsensitive", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"Hello", "World"})
		b := corestr.New.Collection.Strings([]string{"hello", "world"})

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(false, b) != true}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected case-insensitive equal", actual)
		actual = args.Map{"result": a.IsEqualsWithSensitive(true, b) != false}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected case-sensitive not equal", actual)

		// mismatch fold
		c := corestr.New.Collection.Strings([]string{"hello", "OTHER"})
		actual = args.Map{"result": a.IsEqualsWithSensitive(false, c)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)

		// mismatch sensitive
		d := corestr.New.Collection.Strings([]string{"Hello", "other"})
		actual = args.Map{"result": a.IsEqualsWithSensitive(true, d)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_CovP1_16_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_CovP1_16_IsEmptyLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": col.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
		empty := corestr.Empty.Collection()
		actual = args.Map{"result": empty.IsEmptyLock()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_CovP1_17_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_CovP1_17_IsEmpty_HasItems", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": col.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
		actual = args.Map{"result": col.HasItems()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has items", actual)
	})
}

func Test_CovP1_18_AddLock(t *testing.T) {
	safeTest(t, "Test_CovP1_18_AddLock", func() {
		// Arrange
		col := corestr.New.Collection.Cap(2)
		col.AddLock("a")
		col.AddLock("b")

		// Act
		actual := args.Map{"result": col.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP1_19_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_CovP1_19_AddNonEmpty", func() {
		// Arrange
		col := corestr.New.Collection.Cap(2)
		col.AddNonEmpty("")

		// Act
		actual := args.Map{"result": col.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for empty string", actual)
		col.AddNonEmpty("a")
		actual = args.Map{"result": col.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovP1_20_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_CovP1_20_AddNonEmptyWhitespace", func() {
		// Arrange
		col := corestr.New.Collection.Cap(2)
		col.AddNonEmptyWhitespace("   ")

		// Act
		actual := args.Map{"result": col.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for whitespace", actual)
		col.AddNonEmptyWhitespace("a")
		actual = args.Map{"result": col.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovP1_21_AddError(t *testing.T) {
	safeTest(t, "Test_CovP1_21_AddError", func() {
		// Arrange
		col := corestr.New.Collection.Cap(2)
		col.AddError(nil)

		// Act
		actual := args.Map{"result": col.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for nil error", actual)
		col.AddError(errors.New("test"))
		actual = args.Map{"result": col.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovP1_22_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_CovP1_22_AsDefaultError", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"err1", "err2"})
		err := col.AsDefaultError()

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil error", actual)
	})
}

func Test_CovP1_23_AsError(t *testing.T) {
	safeTest(t, "Test_CovP1_23_AsError", func() {
		// Arrange
		empty := corestr.Empty.Collection()

		// Act
		actual := args.Map{"result": empty.AsError(",") != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil for empty", actual)
		col := corestr.New.Collection.Strings([]string{"a"})
		actual = args.Map{"result": col.AsError(",") == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_CovP1_24_AddIf(t *testing.T) {
	safeTest(t, "Test_CovP1_24_AddIf", func() {
		// Arrange
		col := corestr.New.Collection.Cap(2)
		col.AddIf(false, "skip")

		// Act
		actual := args.Map{"result": col.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		col.AddIf(true, "keep")
		actual = args.Map{"result": col.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovP1_25_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_CovP1_25_EachItemSplitBy", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a,b", "c,d"})
		result := col.EachItemSplitBy(",")

		// Act
		actual := args.Map{"result": len(result) != 4}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_CovP1_26_ConcatNew(t *testing.T) {
	safeTest(t, "Test_CovP1_26_ConcatNew", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		// no addingStrings
		newCol := col.ConcatNew(0)

		// Act
		actual := args.Map{"result": newCol.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// with addingStrings
		newCol2 := col.ConcatNew(0, "c", "d")
		actual = args.Map{"result": newCol2.Length() != 4}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_CovP1_27_ToError(t *testing.T) {
	safeTest(t, "Test_CovP1_27_ToError", func() {
		col := corestr.New.Collection.Strings([]string{"e1"})
		err := col.ToError(",")
		_ = err
	})
}

func Test_CovP1_28_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_CovP1_28_ToDefaultError", func() {
		col := corestr.New.Collection.Strings([]string{"e1"})
		_ = col.ToDefaultError()
	})
}

func Test_CovP1_29_AddIfMany(t *testing.T) {
	safeTest(t, "Test_CovP1_29_AddIfMany", func() {
		// Arrange
		col := corestr.New.Collection.Cap(5)
		col.AddIfMany(false, "a", "b")

		// Act
		actual := args.Map{"result": col.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		col.AddIfMany(true, "a", "b")
		actual = args.Map{"result": col.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP1_30_AddFunc(t *testing.T) {
	safeTest(t, "Test_CovP1_30_AddFunc", func() {
		// Arrange
		col := corestr.New.Collection.Cap(2)
		col.AddFunc(func() string { return "computed" })

		// Act
		actual := args.Map{"result": col.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovP1_31_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_CovP1_31_AddFuncErr", func() {
		// Arrange
		col := corestr.New.Collection.Cap(2)
		// success
		col.AddFuncErr(
			func() (string, error) { return "ok", nil },

		// Assert
			func(err error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) },
		)

		// Act
		actual = args.Map{"result": col.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// error
		errCalled := false
		col.AddFuncErr(
			func() (string, error) { return "", errors.New("fail") },
			func(err error) { errCalled = true },
		)
		actual = args.Map{"result": errCalled}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected error handler called", actual)
		actual = args.Map{"result": col.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 1", actual)
	})
}

func Test_CovP1_32_AddsLock(t *testing.T) {
	safeTest(t, "Test_CovP1_32_AddsLock", func() {
		// Arrange
		col := corestr.New.Collection.Cap(5)
		col.AddsLock("a", "b")

		// Act
		actual := args.Map{"result": col.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP1_33_Adds(t *testing.T) {
	safeTest(t, "Test_CovP1_33_Adds", func() {
		// Arrange
		col := corestr.New.Collection.Cap(5)
		col.Adds("a", "b", "c")

		// Act
		actual := args.Map{"result": col.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_CovP1_34_AddStrings(t *testing.T) {
	safeTest(t, "Test_CovP1_34_AddStrings", func() {
		// Arrange
		col := corestr.New.Collection.Cap(5)
		col.AddStrings([]string{"x", "y"})

		// Act
		actual := args.Map{"result": col.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP1_35_AddCollection(t *testing.T) {
	safeTest(t, "Test_CovP1_35_AddCollection", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"b", "c"})
		a.AddCollection(b)

		// Act
		actual := args.Map{"result": a.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// empty collection
		a.AddCollection(corestr.Empty.Collection())
		actual = args.Map{"result": a.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 3", actual)
	})
}

func Test_CovP1_36_AddCollections(t *testing.T) {
	safeTest(t, "Test_CovP1_36_AddCollections", func() {
		// Arrange
		col := corestr.New.Collection.Cap(10)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		col.AddCollections(c1, c2, corestr.Empty.Collection())

		// Act
		actual := args.Map{"result": col.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP1_37_AddPointerCollectionsLock(t *testing.T) {
	safeTest(t, "Test_CovP1_37_AddPointerCollectionsLock", func() {
		// Arrange
		col := corestr.New.Collection.Cap(10)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		col.AddPointerCollectionsLock(c1)

		// Act
		actual := args.Map{"result": col.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovP1_38_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_CovP1_38_AddHashmapsValues", func() {
		// Arrange
		col := corestr.New.Collection.Cap(10)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "v1")
		col.AddHashmapsValues(hm)

		// Act
		actual := args.Map{"result": col.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// nil hashmaps
		col.AddHashmapsValues()
		// nil hashmap item
		col.AddHashmapsValues(nil)
	})
}

func Test_CovP1_39_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_CovP1_39_AddHashmapsKeys", func() {
		// Arrange
		col := corestr.New.Collection.Cap(10)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "v1")
		col.AddHashmapsKeys(hm)

		// Act
		actual := args.Map{"result": col.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// nil
		col.AddHashmapsKeys()
		col.AddHashmapsKeys(nil)
	})
}

func Test_CovP1_40_AddHashmapsKeysValues(t *testing.T) {
	safeTest(t, "Test_CovP1_40_AddHashmapsKeysValues", func() {
		// Arrange
		col := corestr.New.Collection.Cap(10)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "v1")
		col.AddHashmapsKeysValues(hm)

		// Act
		actual := args.Map{"result": col.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// nil
		col2 := corestr.New.Collection.Cap(10)
		col2.AddHashmapsKeysValues()
		col2.AddHashmapsKeysValues(nil)
	})
}

func Test_CovP1_41_AddHashmapsKeysValuesUsingFilter(t *testing.T) {
	safeTest(t, "Test_CovP1_41_AddHashmapsKeysValuesUsingFilter", func() {
		// Arrange
		col := corestr.New.Collection.Cap(10)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "v1")
		hm.AddOrUpdate("k2", "v2")
		col.AddHashmapsKeysValuesUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Key + "=" + pair.Value, true, false
			},
			hm,
		)

		// Act
		actual := args.Map{"result": col.Length() < 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
		// nil hashmaps
		col.AddHashmapsKeysValuesUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) { return "", false, false },
		)
		// nil hashmap item
		col.AddHashmapsKeysValuesUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) { return "", false, false },
			nil,
		)
		// break
		col2 := corestr.New.Collection.Cap(10)
		col2.AddHashmapsKeysValuesUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Key, true, true // break on first
			},
			hm,
		)
		// not keep
		col3 := corestr.New.Collection.Cap(10)
		col3.AddHashmapsKeysValuesUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return "", false, false
			},
			hm,
		)
	})
}

func Test_CovP1_42_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_CovP1_42_AddWithWgLock", func() {
		// Arrange
		col := corestr.New.Collection.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		col.AddWithWgLock(wg, "item")
		wg.Wait()

		// Act
		actual := args.Map{"result": col.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovP1_43_IndexAt(t *testing.T) {
	safeTest(t, "Test_CovP1_43_IndexAt", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": col.IndexAt(0) != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_CovP1_44_SafeIndexAtUsingLength(t *testing.T) {
	safeTest(t, "Test_CovP1_44_SafeIndexAtUsingLength", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": col.SafeIndexAtUsingLength("def", 2, 0) != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": col.SafeIndexAtUsingLength("def", 2, 5) != "def"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default", actual)
	})
}

func Test_CovP1_45_First(t *testing.T) {
	safeTest(t, "Test_CovP1_45_First", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"first", "second"})

		// Act
		actual := args.Map{"result": col.First() != "first"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected first", actual)
	})
}

func Test_CovP1_46_Single(t *testing.T) {
	safeTest(t, "Test_CovP1_46_Single", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"only"})

		// Act
		actual := args.Map{"result": col.Single() != "only"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected only", actual)
		// panic on non-single
		defer func() { recover() }()
		col2 := corestr.New.Collection.Strings([]string{"a", "b"})
		col2.Single()
		actual = args.Map{"result": false}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	})
}

func Test_CovP1_47_Last(t *testing.T) {
	safeTest(t, "Test_CovP1_47_Last", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "last"})

		// Act
		actual := args.Map{"result": col.Last() != "last"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected last", actual)
	})
}

func Test_CovP1_48_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_CovP1_48_LastOrDefault", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": col.LastOrDefault() != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
		empty := corestr.Empty.Collection()
		actual = args.Map{"result": empty.LastOrDefault() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty string", actual)
	})
}

func Test_CovP1_49_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_CovP1_49_FirstOrDefault", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": col.FirstOrDefault() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		empty := corestr.Empty.Collection()
		actual = args.Map{"result": empty.FirstOrDefault() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty string", actual)
	})
}

func Test_CovP1_50_ResizeForItems_LargeResize(t *testing.T) {
	safeTest(t, "Test_CovP1_50_ResizeForItems_LargeResize", func() {
		// Arrange
		col := corestr.New.Collection.Cap(5)
		// Add many items to trigger resize logic
		items := make([]string, 300)
		for i := range items {
			items[i] = "x"
		}
		col.AddStrings(items)

		// Act
		actual := args.Map{"result": col.Length() != 300}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 300", actual)
	})
}
