package corestrtests

import (
	"errors"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// Collection — JSON and Serialization
// =============================================================================

func Test_I8_C01_Collection_JsonString(t *testing.T) {
	safeTest(t, "Test_I8_C01_Collection_JsonString", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.JsonString()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty json", actual)
	})
}

func Test_I8_C02_Collection_JsonStringMust(t *testing.T) {
	safeTest(t, "Test_I8_C02_Collection_JsonStringMust", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.JsonStringMust()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_I8_C03_Collection_StringJSON(t *testing.T) {
	safeTest(t, "Test_I8_C03_Collection_StringJSON", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"x"})
		s := c.StringJSON()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_I8_C04_Collection_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_I8_C04_Collection_HasAnyItem", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.HasAnyItem()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		e := corestr.Empty.Collection()
		actual := args.Map{"result": e.HasAnyItem()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_I8_C05_Collection_LastIndex(t *testing.T) {
	safeTest(t, "Test_I8_C05_Collection_LastIndex", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"result": c.LastIndex() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C06_Collection_HasIndex(t *testing.T) {
	safeTest(t, "Test_I8_C06_Collection_HasIndex", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.HasIndex(0)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for 0", actual)
		actual := args.Map{"result": c.HasIndex(1)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for 1", actual)
		actual := args.Map{"result": c.HasIndex(2)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for 2", actual)
		actual := args.Map{"result": c.HasIndex(-1)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for -1", actual)
	})
}

func Test_I8_C07_Collection_ListStringsPtr(t *testing.T) {
	safeTest(t, "Test_I8_C07_Collection_ListStringsPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": len(c.ListStringsPtr()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_C08_Collection_ListStrings(t *testing.T) {
	safeTest(t, "Test_I8_C08_Collection_ListStrings", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": len(c.ListStrings()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_C09_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_I8_C09_Collection_RemoveAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok := c.RemoveAt(1)

		// Act
		actual := args.Map{"result": ok}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected success", actual)
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// negative index
		actual := args.Map{"result": c.RemoveAt(-1)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for negative", actual)
		// out of range
		actual := args.Map{"result": c.RemoveAt(100)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for out of range", actual)
	})
}

func Test_I8_C10_Collection_Count(t *testing.T) {
	safeTest(t, "Test_I8_C10_Collection_Count", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.Count() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_C11_Collection_Capacity(t *testing.T) {
	safeTest(t, "Test_I8_C11_Collection_Capacity", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"result": c.Capacity() < 10}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 10", actual)
	})
}

func Test_I8_C12_Collection_Capacity_Nil(t *testing.T) {
	safeTest(t, "Test_I8_C12_Collection_Capacity_Nil", func() {
		c := corestr.New.Collection.Strings(nil)
		_ = c.Capacity()
	})
}

func Test_I8_C13_Collection_LengthLock(t *testing.T) {
	safeTest(t, "Test_I8_C13_Collection_LengthLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.LengthLock() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// =============================================================================
// Collection — Equality
// =============================================================================

func Test_I8_C14_Collection_IsEquals(t *testing.T) {
	safeTest(t, "Test_I8_C14_Collection_IsEquals", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_I8_C15_Collection_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_I8_C15_Collection_IsEqualsWithSensitive_CaseInsensitive", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"Hello"})
		b := corestr.New.Collection.Strings([]string{"hello"})

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(false, b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal case-insensitive", actual)
		actual := args.Map{"result": a.IsEqualsWithSensitive(true, b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal case-sensitive", actual)
	})
}

func Test_I8_C16_Collection_IsEquals_BothNil(t *testing.T) {
	safeTest(t, "Test_I8_C16_Collection_IsEquals_BothNil", func() {
		var a, b *corestr.Collection
		_ = a
		_ = b
		// Can't directly test nil.IsEquals(nil) due to nil receiver panic
		// But we test via isCollectionPrecheckEqual through different path
	})
}

func Test_I8_C17_Collection_IsEquals_DiffLength(t *testing.T) {
	safeTest(t, "Test_I8_C17_Collection_IsEquals_DiffLength", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal for different length", actual)
	})
}

func Test_I8_C18_Collection_IsEquals_SamePtr(t *testing.T) {
	safeTest(t, "Test_I8_C18_Collection_IsEquals_SamePtr", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": a.IsEquals(a)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal for same pointer", actual)
	})
}

func Test_I8_C19_Collection_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_I8_C19_Collection_IsEquals_BothEmpty", func() {
		// Arrange
		a := corestr.Empty.Collection()
		b := corestr.Empty.Collection()

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal for both empty", actual)
	})
}

func Test_I8_C20_Collection_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_I8_C20_Collection_IsEquals_OneEmpty", func() {
		// Arrange
		a := corestr.Empty.Collection()
		b := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

// =============================================================================
// Collection — Add variants
// =============================================================================

func Test_I8_C21_Collection_IsEmpty(t *testing.T) {
	safeTest(t, "Test_I8_C21_Collection_IsEmpty", func() {
		// Arrange
		c := corestr.Empty.Collection()

		// Act
		actual := args.Map{"result": c.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual := args.Map{"result": c.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty with lock", actual)
	})
}

func Test_I8_C22_Collection_HasItems(t *testing.T) {
	safeTest(t, "Test_I8_C22_Collection_HasItems", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.HasItems()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_I8_C23_Collection_AddLock(t *testing.T) {
	safeTest(t, "Test_I8_C23_Collection_AddLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(2)
		c.AddLock("a")

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_C24_Collection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_I8_C24_Collection_AddNonEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(2)
		c.AddNonEmpty("")

		// Act
		actual := args.Map{"result": c.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for empty string", actual)
		c.AddNonEmpty("a")
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_C25_Collection_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_I8_C25_Collection_AddNonEmptyWhitespace", func() {
		// Arrange
		c := corestr.New.Collection.Cap(2)
		c.AddNonEmptyWhitespace("  ")

		// Act
		actual := args.Map{"result": c.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for whitespace", actual)
		c.AddNonEmptyWhitespace("a")
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_C26_Collection_AddError(t *testing.T) {
	safeTest(t, "Test_I8_C26_Collection_AddError", func() {
		// Arrange
		c := corestr.New.Collection.Cap(2)
		c.AddError(nil)

		// Act
		actual := args.Map{"result": c.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for nil error", actual)
		c.AddError(errors.New("test"))
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_C27_Collection_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_I8_C27_Collection_AsDefaultError", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"err1", "err2"})
		err := c.AsDefaultError()

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_I8_C28_Collection_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_I8_C28_Collection_AsError_Empty", func() {
		// Arrange
		c := corestr.Empty.Collection()
		err := c.AsError(",")

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_I8_C29_Collection_AddIf(t *testing.T) {
	safeTest(t, "Test_I8_C29_Collection_AddIf", func() {
		// Arrange
		c := corestr.New.Collection.Cap(2)
		c.AddIf(false, "skip")

		// Act
		actual := args.Map{"result": c.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		c.AddIf(true, "keep")
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_C30_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_I8_C30_Collection_EachItemSplitBy", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a.b", "c.d"})
		result := c.EachItemSplitBy(".")

		// Act
		actual := args.Map{"result": len(result) != 4}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_I8_C31_Collection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_I8_C31_Collection_ConcatNew", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		nc := c.ConcatNew(0, "b", "c")

		// Act
		actual := args.Map{"result": nc.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_I8_C32_Collection_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_I8_C32_Collection_ConcatNew_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		nc := c.ConcatNew(0)

		// Act
		actual := args.Map{"result": nc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_C33_Collection_AddIfMany(t *testing.T) {
	safeTest(t, "Test_I8_C33_Collection_AddIfMany", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddIfMany(false, "skip1", "skip2")

		// Act
		actual := args.Map{"result": c.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		c.AddIfMany(true, "keep1", "keep2")
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C34_Collection_AddFunc(t *testing.T) {
	safeTest(t, "Test_I8_C34_Collection_AddFunc", func() {
		// Arrange
		c := corestr.New.Collection.Cap(2)
		c.AddFunc(func() string { return "hello" })

		// Act
		actual := args.Map{"result": c.Length() != 1 || c.First() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'hello'", actual)
	})
}

func Test_I8_C35_Collection_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_I8_C35_Collection_AddFuncErr", func() {
		// Arrange
		c := corestr.New.Collection.Cap(2)
		// success
		c.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// error
		c.AddFuncErr(func() (string, error) { return "", errors.New("fail") }, func(e error) {})
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 1", actual)
	})
}

func Test_I8_C36_Collection_AddsLock(t *testing.T) {
	safeTest(t, "Test_I8_C36_Collection_AddsLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddsLock("a", "b")

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C37_Collection_AddStrings(t *testing.T) {
	safeTest(t, "Test_I8_C37_Collection_AddStrings", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddStrings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C38_Collection_AddCollection(t *testing.T) {
	safeTest(t, "Test_I8_C38_Collection_AddCollection", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		other := corestr.New.Collection.Strings([]string{"x", "y"})
		c.AddCollection(other)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty collection
		c.AddCollection(corestr.Empty.Collection())
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 2", actual)
	})
}

func Test_I8_C39_Collection_AddCollections(t *testing.T) {
	safeTest(t, "Test_I8_C39_Collection_AddCollections", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c.AddCollections(c1, c2, corestr.Empty.Collection())

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C40_Collection_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_I8_C40_Collection_AddWithWgLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.AddWithWgLock(wg, "a")
		wg.Wait()

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// =============================================================================
// Collection — Access, Sort, Filter
// =============================================================================

func Test_I8_C41_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_I8_C41_Collection_IndexAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.IndexAt(0) != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
	})
}

func Test_I8_C42_Collection_SafeIndexAtUsingLength(t *testing.T) {
	safeTest(t, "Test_I8_C42_Collection_SafeIndexAtUsingLength", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.SafeIndexAtUsingLength("default", 2, 5) != "default"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default", actual)
		actual := args.Map{"result": c.SafeIndexAtUsingLength("default", 2, 0) != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
	})
}

func Test_I8_C43_Collection_First(t *testing.T) {
	safeTest(t, "Test_I8_C43_Collection_First", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
	})
}

func Test_I8_C44_Collection_Last(t *testing.T) {
	safeTest(t, "Test_I8_C44_Collection_Last", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.Last() != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'b'", actual)
	})
}

func Test_I8_C45_Collection_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_I8_C45_Collection_LastOrDefault", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.LastOrDefault() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
		e := corestr.Empty.Collection()
		actual := args.Map{"result": e.LastOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_I8_C46_Collection_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_I8_C46_Collection_FirstOrDefault", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.FirstOrDefault() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
		e := corestr.Empty.Collection()
		actual := args.Map{"result": e.FirstOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_I8_C47_Collection_Take(t *testing.T) {
	safeTest(t, "Test_I8_C47_Collection_Take", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		taken := c.Take(2)

		// Act
		actual := args.Map{"result": taken.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// take more than length
		taken2 := c.Take(10)
		actual := args.Map{"result": taken2.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// take 0
		taken3 := c.Take(0)
		actual := args.Map{"result": taken3.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_I8_C48_Collection_Skip(t *testing.T) {
	safeTest(t, "Test_I8_C48_Collection_Skip", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		skipped := c.Skip(1)

		// Act
		actual := args.Map{"result": skipped.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// skip 0
		skipped2 := c.Skip(0)
		actual := args.Map{"result": skipped2.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_I8_C49_Collection_Reverse(t *testing.T) {
	safeTest(t, "Test_I8_C49_Collection_Reverse", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.Reverse()

		// Act
		actual := args.Map{"result": c.First() != "c" || c.Last() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected reversed", actual)
	})
}

func Test_I8_C50_Collection_Reverse_Two(t *testing.T) {
	safeTest(t, "Test_I8_C50_Collection_Reverse_Two", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Reverse()

		// Act
		actual := args.Map{"result": c.First() != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'b' first", actual)
	})
}

func Test_I8_C51_Collection_Reverse_One(t *testing.T) {
	safeTest(t, "Test_I8_C51_Collection_Reverse_One", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Reverse()

		// Act
		actual := args.Map{"result": c.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
	})
}

func Test_I8_C52_Collection_GetPagesSize(t *testing.T) {
	safeTest(t, "Test_I8_C52_Collection_GetPagesSize", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		pages := c.GetPagesSize(2)

		// Act
		actual := args.Map{"result": pages != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual := args.Map{"result": c.GetPagesSize(0) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for zero page size", actual)
	})
}

func Test_I8_C53_Collection_GetSinglePageCollection(t *testing.T) {
	safeTest(t, "Test_I8_C53_Collection_GetSinglePageCollection", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		page := c.GetSinglePageCollection(2, 2)

		// Act
		actual := args.Map{"result": page.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		page3 := c.GetSinglePageCollection(2, 3)
		actual := args.Map{"result": page3.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 for last page", actual)
	})
}

func Test_I8_C54_Collection_GetPagedCollection(t *testing.T) {
	safeTest(t, "Test_I8_C54_Collection_GetPagedCollection", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		pages := c.GetPagedCollection(2)

		// Act
		actual := args.Map{"result": pages.Length() < 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3 pages", actual)
	})
}

func Test_I8_C55_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_I8_C55_Collection_Filter", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "bb", "ccc"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C56_Collection_Filter_WithBreak(t *testing.T) {
	safeTest(t, "Test_I8_C56_Collection_Filter_WithBreak", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "bb", "ccc"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, true, i == 1
		})

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C57_Collection_FilterLock(t *testing.T) {
	safeTest(t, "Test_I8_C57_Collection_FilterLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		result := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C58_Collection_FilteredCollection(t *testing.T) {
	safeTest(t, "Test_I8_C58_Collection_FilteredCollection", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		fc := c.FilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})

		// Act
		actual := args.Map{"result": fc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_C59_Collection_FilterPtr(t *testing.T) {
	safeTest(t, "Test_I8_C59_Collection_FilterPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.FilterPtr(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})

		// Act
		actual := args.Map{"result": len(*result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C60_Collection_FilterPtrLock(t *testing.T) {
	safeTest(t, "Test_I8_C60_Collection_FilterPtrLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.FilterPtrLock(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})

		// Act
		actual := args.Map{"result": len(*result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// =============================================================================
// Collection — Search, Sort, CSV, etc.
// =============================================================================

func Test_I8_C61_Collection_Has(t *testing.T) {
	safeTest(t, "Test_I8_C61_Collection_Has", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": c.Has("z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_I8_C62_Collection_HasPtr(t *testing.T) {
	safeTest(t, "Test_I8_C62_Collection_HasPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := "a"

		// Act
		actual := args.Map{"result": c.HasPtr(&s)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": c.HasPtr(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_I8_C63_Collection_HasAll(t *testing.T) {
	safeTest(t, "Test_I8_C63_Collection_HasAll", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"result": c.HasAll("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": c.HasAll("a", "z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_I8_C64_Collection_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_I8_C64_Collection_SortedListAsc", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		sorted := c.SortedListAsc()

		// Act
		actual := args.Map{"result": sorted[0] != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a' first", actual)
	})
}

func Test_I8_C65_Collection_SortedAsc(t *testing.T) {
	safeTest(t, "Test_I8_C65_Collection_SortedAsc", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		c.SortedAsc()

		// Act
		actual := args.Map{"result": c.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
	})
}

func Test_I8_C66_Collection_SortedAscLock(t *testing.T) {
	safeTest(t, "Test_I8_C66_Collection_SortedAscLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"c", "a"})
		c.SortedAscLock()

		// Act
		actual := args.Map{"result": c.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
	})
}

func Test_I8_C67_Collection_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_I8_C67_Collection_SortedListDsc", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "c", "b"})
		sorted := c.SortedListDsc()

		// Act
		actual := args.Map{"result": sorted[0] != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'c' first", actual)
	})
}

func Test_I8_C68_Collection_HasUsingSensitivity(t *testing.T) {
	safeTest(t, "Test_I8_C68_Collection_HasUsingSensitivity", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"Hello"})

		// Act
		actual := args.Map{"result": c.HasUsingSensitivity("hello", false)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true case-insensitive", actual)
		actual := args.Map{"result": c.HasUsingSensitivity("hello", true)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false case-sensitive", actual)
	})
}

func Test_I8_C69_Collection_IsContainsAll(t *testing.T) {
	safeTest(t, "Test_I8_C69_Collection_IsContainsAll", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"result": c.IsContainsAll("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": c.IsContainsAll("a", "z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual := args.Map{"result": c.IsContainsAll(nil...)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_I8_C70_Collection_IsContainsAllLock(t *testing.T) {
	safeTest(t, "Test_I8_C70_Collection_IsContainsAllLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.IsContainsAllLock("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_I8_C71_Collection_Csv(t *testing.T) {
	safeTest(t, "Test_I8_C71_Collection_Csv", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		csv := c.Csv()

		// Act
		actual := args.Map{"result": csv == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_I8_C72_Collection_CsvEmpty(t *testing.T) {
	safeTest(t, "Test_I8_C72_Collection_CsvEmpty", func() {
		// Arrange
		c := corestr.Empty.Collection()

		// Act
		actual := args.Map{"result": c.Csv() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_I8_C73_Collection_String(t *testing.T) {
	safeTest(t, "Test_I8_C73_Collection_String", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_I8_C74_Collection_StringLock(t *testing.T) {
	safeTest(t, "Test_I8_C74_Collection_StringLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.StringLock()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_I8_C75_Collection_Clear(t *testing.T) {
	safeTest(t, "Test_I8_C75_Collection_Clear", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Clear()

		// Act
		actual := args.Map{"result": c.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_I8_C76_Collection_Dispose(t *testing.T) {
	safeTest(t, "Test_I8_C76_Collection_Dispose", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Dispose()

		// Act
		actual := args.Map{"result": c.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_I8_C77_Collection_Join(t *testing.T) {
	safeTest(t, "Test_I8_C77_Collection_Join", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.Join(",") != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b'", actual)
	})
}

func Test_I8_C78_Collection_JoinLine(t *testing.T) {
	safeTest(t, "Test_I8_C78_Collection_JoinLine", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.JoinLine()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_I8_C79_Collection_Json(t *testing.T) {
	safeTest(t, "Test_I8_C79_Collection_Json", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		j := c.Json()

		// Act
		actual := args.Map{"result": j.JsonString() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_I8_C80_Collection_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_I8_C80_Collection_ParseInjectUsingJson", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		jr := c.JsonPtr()
		c2 := corestr.Empty.Collection()
		_, err := c2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_I8_C81_Collection_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_I8_C81_Collection_ParseInjectUsingJson_Error", func() {
		// Arrange
		c := corestr.Empty.Collection()
		bad := corejson.NewResult.UsingString(`invalid`)
		_, err := c.ParseInjectUsingJson(bad)

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_I8_C82_Collection_Serialize(t *testing.T) {
	safeTest(t, "Test_I8_C82_Collection_Serialize", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		bytes, err := c.Serialize()

		// Act
		actual := args.Map{"result": err != nil || len(bytes) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected serialization", actual)
	})
}

func Test_I8_C83_Collection_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_I8_C83_Collection_MarshalJSON", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		data, err := c.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil || len(data) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected marshal", actual)
	})
}

func Test_I8_C84_Collection_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_I8_C84_Collection_UnmarshalJSON", func() {
		// Arrange
		c := corestr.Empty.Collection()
		err := c.UnmarshalJSON([]byte(`["a","b"]`))

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// =============================================================================
// Collection — More methods
// =============================================================================

func Test_I8_C85_Collection_NonEmptyList(t *testing.T) {
	safeTest(t, "Test_I8_C85_Collection_NonEmptyList", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		list := c.NonEmptyList()

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C86_Collection_UniqueList(t *testing.T) {
	safeTest(t, "Test_I8_C86_Collection_UniqueList", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "a"})
		list := c.UniqueList()

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C87_Collection_UniqueListLock(t *testing.T) {
	safeTest(t, "Test_I8_C87_Collection_UniqueListLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "a"})
		list := c.UniqueListLock()

		// Act
		actual := args.Map{"result": len(list) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_C88_Collection_UniqueBoolMap(t *testing.T) {
	safeTest(t, "Test_I8_C88_Collection_UniqueBoolMap", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		m := c.UniqueBoolMap()

		// Act
		actual := args.Map{"result": len(m) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C89_Collection_HashsetAsIs(t *testing.T) {
	safeTest(t, "Test_I8_C89_Collection_HashsetAsIs", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs := c.HashsetAsIs()

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C90_Collection_Resize(t *testing.T) {
	safeTest(t, "Test_I8_C90_Collection_Resize", func() {
		// Arrange
		c := corestr.New.Collection.Cap(2)
		c.Add("a")
		c.Resize(100)

		// Act
		actual := args.Map{"result": c.Capacity() < 100}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected capacity >= 100", actual)
	})
}

func Test_I8_C91_Collection_AddCapacity(t *testing.T) {
	safeTest(t, "Test_I8_C91_Collection_AddCapacity", func() {
		// Arrange
		c := corestr.New.Collection.Cap(2)
		c.AddCapacity(50)

		// Act
		actual := args.Map{"result": c.Capacity() < 50}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected capacity >= 50", actual)
	})
}

func Test_I8_C92_Collection_Joins(t *testing.T) {
	safeTest(t, "Test_I8_C92_Collection_Joins", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.Joins(",")

		// Act
		actual := args.Map{"result": s != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b'", actual)
		s2 := c.Joins(",", "c")
		actual := args.Map{"result": s2 == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_I8_C93_Collection_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_I8_C93_Collection_GetAllExcept", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.GetAllExcept([]string{"b"})

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C94_Collection_GetAllExcept_Nil(t *testing.T) {
	safeTest(t, "Test_I8_C94_Collection_GetAllExcept_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.GetAllExcept(nil)

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C95_Collection_New(t *testing.T) {
	safeTest(t, "Test_I8_C95_Collection_New", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		nc := c.New("x", "y")

		// Act
		actual := args.Map{"result": nc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		nc2 := c.New()
		actual := args.Map{"result": nc2.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_I8_C96_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_I8_C96_Collection_AppendAnys", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AppendAnys(42, "hello", nil)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 (nil skipped)", actual)
	})
}

func Test_I8_C97_Collection_AppendAnysLock(t *testing.T) {
	safeTest(t, "Test_I8_C97_Collection_AppendAnysLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysLock(42, "hello")

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C98_Collection_AppendNonEmptyAnys(t *testing.T) {
	safeTest(t, "Test_I8_C98_Collection_AppendNonEmptyAnys", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AppendNonEmptyAnys(42, nil)

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_C99_Collection_AddsNonEmpty(t *testing.T) {
	safeTest(t, "Test_I8_C99_Collection_AddsNonEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddsNonEmpty("a", "", "b")

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_C100_Collection_Single(t *testing.T) {
	safeTest(t, "Test_I8_C100_Collection_Single", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"only"})

		// Act
		actual := args.Map{"result": c.Single() != "only"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'only'", actual)
	})
}
