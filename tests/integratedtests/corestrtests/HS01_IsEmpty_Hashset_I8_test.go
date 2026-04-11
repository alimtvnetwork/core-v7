package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// Hashset — Core operations
// =============================================================================

func Test_HS01_IsEmpty(t *testing.T) {
	safeTest(t, "Test_I8_HS01_IsEmpty", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{"result": h.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		h.Add("a")
		actual = args.Map{"result": h.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_HS02_HasItems(t *testing.T) {
	safeTest(t, "Test_I8_HS02_HasItems", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{"result": h.HasItems()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		h.Add("a")
		actual = args.Map{"result": h.HasItems()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_HS03_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_I8_HS03_HasAnyItem", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{"result": h.HasAnyItem()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_HS04_Add(t *testing.T) {
	safeTest(t, "Test_I8_HS04_Add", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.Add("a").Add("b")

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_HS05_AddBool(t *testing.T) {
	safeTest(t, "Test_I8_HS05_AddBool", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		existed := h.AddBool("a")

		// Act
		actual := args.Map{"result": existed}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for new key", actual)
		existed2 := h.AddBool("a")
		actual = args.Map{"result": existed2}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for existing key", actual)
	})
}

func Test_HS06_AddPtr(t *testing.T) {
	safeTest(t, "Test_I8_HS06_AddPtr", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		s := "test"
		h.AddPtr(&s)

		// Act
		actual := args.Map{"result": h.Has("test")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected to have 'test'", actual)
	})
}

func Test_HS07_AddPtrLock(t *testing.T) {
	safeTest(t, "Test_I8_HS07_AddPtrLock", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		s := "test"
		h.AddPtrLock(&s)

		// Act
		actual := args.Map{"result": h.Has("test")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected to have 'test'", actual)
	})
}

func Test_HS08_AddLock(t *testing.T) {
	safeTest(t, "Test_I8_HS08_AddLock", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddLock("a")

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HS09_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_I8_HS09_AddNonEmpty", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddNonEmpty("")

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		h.AddNonEmpty("a")
		actual = args.Map{"result": h.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HS10_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_I8_HS10_AddNonEmptyWhitespace", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddNonEmptyWhitespace("  ")

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		h.AddNonEmptyWhitespace("a")
		actual = args.Map{"result": h.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HS11_AddIf(t *testing.T) {
	safeTest(t, "Test_I8_HS11_AddIf", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddIf(false, "skip")

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		h.AddIf(true, "keep")
		actual = args.Map{"result": h.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HS12_AddIfMany(t *testing.T) {
	safeTest(t, "Test_I8_HS12_AddIfMany", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddIfMany(false, "a", "b")

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		h.AddIfMany(true, "a", "b")
		actual = args.Map{"result": h.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_HS13_AddFunc(t *testing.T) {
	safeTest(t, "Test_I8_HS13_AddFunc", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddFunc(func() string { return "x" })

		// Act
		actual := args.Map{"result": h.Has("x")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 'x'", actual)
	})
}

func Test_HS14_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_I8_HS14_AddFuncErr", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HS15_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_I8_HS15_AddWithWgLock", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddWithWgLock("a", wg)
		wg.Wait()

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HS16_Adds(t *testing.T) {
	safeTest(t, "Test_I8_HS16_Adds", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.Adds("a", "b", "c")

		// Act
		actual := args.Map{"result": h.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_HS17_AddStrings(t *testing.T) {
	safeTest(t, "Test_I8_HS17_AddStrings", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddStrings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_HS18_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_I8_HS18_AddStringsLock", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddStringsLock([]string{"a"})

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HS19_AddCollection(t *testing.T) {
	safeTest(t, "Test_I8_HS19_AddCollection", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		c := corestr.New.Collection.Strings([]string{"x", "y"})
		h.AddCollection(c)

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		h.AddCollection(nil)
		actual = args.Map{"result": h.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 2", actual)
	})
}

func Test_HS20_AddCollections(t *testing.T) {
	safeTest(t, "Test_I8_HS20_AddCollections", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		h.AddCollections(c1, c2, nil)

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_HS21_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_I8_HS21_AddHashsetItems", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		other := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.AddHashsetItems(other)

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		h.AddHashsetItems(nil)
		actual = args.Map{"result": h.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 2", actual)
	})
}

func Test_HS22_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_I8_HS22_AddItemsMap", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		m := map[string]bool{"a": true, "b": false, "c": true}
		h.AddItemsMap(m)

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 (b=false excluded)", actual)
		h.AddItemsMap(nil)
	})
}

func Test_HS23_AddSimpleSlice(t *testing.T) {
	safeTest(t, "Test_I8_HS23_AddSimpleSlice", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		h.AddSimpleSlice(ss)

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// =============================================================================
// Hashset — Query operations
// =============================================================================

func Test_HS24_Has(t *testing.T) {
	safeTest(t, "Test_I8_HS24_Has", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": h.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": h.Has("z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_HS25_HasLock(t *testing.T) {
	safeTest(t, "Test_I8_HS25_HasLock", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": h.HasLock("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_HS26_HasAll(t *testing.T) {
	safeTest(t, "Test_I8_HS26_HasAll", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"result": h.HasAll("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": h.HasAll("a", "z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_HS27_HasAny(t *testing.T) {
	safeTest(t, "Test_I8_HS27_HasAny", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": h.HasAny("a", "z")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": h.HasAny("x", "y")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_HS28_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_I8_HS28_HasAllStrings", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": h.HasAllStrings([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_HS29_IsMissing(t *testing.T) {
	safeTest(t, "Test_I8_HS29_IsMissing", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": h.IsMissing("a")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": h.IsMissing("z")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_HS30_IsMissingLock(t *testing.T) {
	safeTest(t, "Test_I8_HS30_IsMissingLock", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": h.IsMissingLock("a")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_HS31_IsAllMissing(t *testing.T) {
	safeTest(t, "Test_I8_HS31_IsAllMissing", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": h.IsAllMissing("x", "y") != true}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": h.IsAllMissing("a", "y") != false}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_HS32_Contains(t *testing.T) {
	safeTest(t, "Test_I8_HS32_Contains", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": h.Contains("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_HS33_IsEqual(t *testing.T) {
	safeTest(t, "Test_I8_HS33_IsEqual", func() {
		// Arrange
		a := corestr.New.Hashset.Strings([]string{"a", "b"})
		b := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": a.IsEqual(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

// =============================================================================
// Hashset — List, Sort, JSON
// =============================================================================

func Test_HS34_List(t *testing.T) {
	safeTest(t, "Test_I8_HS34_List", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": len(h.List()) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_HS35_SortedList(t *testing.T) {
	safeTest(t, "Test_I8_HS35_SortedList", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
		sorted := h.SortedList()

		// Act
		actual := args.Map{"result": sorted[0] != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a' first", actual)
	})
}

func Test_HS36_OrderedList(t *testing.T) {
	safeTest(t, "Test_I8_HS36_OrderedList", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"c", "a"})
		ordered := h.OrderedList()

		// Act
		actual := args.Map{"result": ordered[0] != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a' first", actual)
	})
}

func Test_HS37_ListPtrSortedAsc(t *testing.T) {
	safeTest(t, "Test_I8_HS37_ListPtrSortedAsc", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"b", "a"})
		list := h.ListPtrSortedAsc()

		// Act
		actual := args.Map{"result": list[0] != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
	})
}

func Test_HS38_ListPtrSortedDsc(t *testing.T) {
	safeTest(t, "Test_I8_HS38_ListPtrSortedDsc", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		list := h.ListPtrSortedDsc()

		// Act
		actual := args.Map{"result": list[0] != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'b'", actual)
	})
}

func Test_HS39_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_I8_HS39_SimpleSlice", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		ss := h.SimpleSlice()

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HS40_SafeStrings(t *testing.T) {
	safeTest(t, "Test_I8_HS40_SafeStrings", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": len(h.SafeStrings()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HS41_Collection(t *testing.T) {
	safeTest(t, "Test_I8_HS41_Collection", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		c := h.Collection()

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HS42_String(t *testing.T) {
	safeTest(t, "Test_I8_HS42_String", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		s := h.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_HS43_StringLock(t *testing.T) {
	safeTest(t, "Test_I8_HS43_StringLock", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		s := h.StringLock()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_HS44_Json(t *testing.T) {
	safeTest(t, "Test_I8_HS44_Json", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		j := h.Json()

		// Act
		actual := args.Map{"result": j.JsonString() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_HS45_JsonModel(t *testing.T) {
	safeTest(t, "Test_I8_HS45_JsonModel", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		m := h.JsonModel()

		// Act
		actual := args.Map{"result": len(m) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HS46_MapStringAny(t *testing.T) {
	safeTest(t, "Test_I8_HS46_MapStringAny", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		m := h.MapStringAny()

		// Act
		actual := args.Map{"result": len(m) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HS47_JoinSorted(t *testing.T) {
	safeTest(t, "Test_I8_HS47_JoinSorted", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"b", "a"})
		s := h.JoinSorted(",")

		// Act
		actual := args.Map{"result": s != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
	})
}

// =============================================================================
// Hashset — Remove, Clear, Dispose, Resize
// =============================================================================

func Test_HS48_Remove(t *testing.T) {
	safeTest(t, "Test_I8_HS48_Remove", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.Remove("a")

		// Act
		actual := args.Map{"result": h.Has("a")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed", actual)
	})
}

func Test_HS49_SafeRemove(t *testing.T) {
	safeTest(t, "Test_I8_HS49_SafeRemove", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.SafeRemove("a")
		h.SafeRemove("nonexistent")

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_HS50_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_I8_HS50_RemoveWithLock", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.RemoveWithLock("a")

		// Act
		actual := args.Map{"result": h.Has("a")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed", actual)
	})
}

func Test_HS51_Clear(t *testing.T) {
	safeTest(t, "Test_I8_HS51_Clear", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.Clear()

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_HS52_Dispose(t *testing.T) {
	safeTest(t, "Test_I8_HS52_Dispose", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.Dispose()

		// Act
		actual := args.Map{"result": h.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_HS53_Resize(t *testing.T) {
	safeTest(t, "Test_I8_HS53_Resize", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		h.Resize(100)

		// Act
		actual := args.Map{"result": h.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HS54_AddCapacities(t *testing.T) {
	safeTest(t, "Test_I8_HS54_AddCapacities", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddCapacities(10, 20)
		_ = h
	})
}

func Test_HS55_ConcatNewHashsets(t *testing.T) {
	safeTest(t, "Test_I8_HS55_ConcatNewHashsets", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		other := corestr.New.Hashset.Strings([]string{"b"})
		result := h.ConcatNewHashsets(true, other)

		// Act
		actual := args.Map{"result": result.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_HS56_ConcatNewHashsets_Empty(t *testing.T) {
	safeTest(t, "Test_I8_HS56_ConcatNewHashsets_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.ConcatNewHashsets(true)

		// Act
		actual := args.Map{"result": result.Length() < 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_HS57_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_I8_HS57_ConcatNewStrings", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.ConcatNewStrings(true, []string{"b", "c"})

		// Act
		actual := args.Map{"result": result.Length() < 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	})
}

func Test_HS58_Filter(t *testing.T) {
	safeTest(t, "Test_I8_HS58_Filter", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"abc", "de", "f"})
		filtered := h.Filter(func(s string) bool { return len(s) > 1 })

		// Act
		actual := args.Map{"result": filtered.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_HS59_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_I8_HS59_AddsUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddsUsingFilter(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		}, "a", "bb", "ccc")

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_HS60_AddsAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_I8_HS60_AddsAnyUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddsAnyUsingFilter(func(s string, i int) (string, bool, bool) {
			return s, true, false
		}, "hello", 42, nil)

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 (nil skipped)", actual)
	})
}

func Test_HS61_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_I8_HS61_IsEmptyLock", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{"result": h.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_HS62_ListCopyLock(t *testing.T) {
	safeTest(t, "Test_I8_HS62_ListCopyLock", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		list := h.ListCopyLock()

		// Act
		actual := args.Map{"result": len(list) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HS63_LengthLock(t *testing.T) {
	safeTest(t, "Test_I8_HS63_LengthLock", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": h.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HS64_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_I8_HS64_ParseInjectUsingJson", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		jr := h.JsonPtr()
		h2 := corestr.New.Hashset.Cap(1)
		_, err := h2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_HS65_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_I8_HS65_ParseInjectUsingJson_Error", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(1)
		bad := corejson.NewResult.UsingString(`invalid`)
		_, err := h.ParseInjectUsingJson(bad)

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_HS66_HasWithLock(t *testing.T) {
	safeTest(t, "Test_I8_HS66_HasWithLock", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": h.HasWithLock("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_HS67_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_I8_HS67_HasAllCollectionItems", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": h.HasAllCollectionItems(c)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": h.HasAllCollectionItems(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_HS68_MapStringAnyDiff(t *testing.T) {
	safeTest(t, "Test_I8_HS68_MapStringAnyDiff", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		_ = h.MapStringAnyDiff()
	})
}

func Test_HS69_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_I8_HS69_WrapDoubleQuote", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		wrapped := h.WrapDoubleQuote()
		list := wrapped.SortedList()

		// Act
		actual := args.Map{"result": len(list) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_HS70_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_I8_HS70_IsEquals_BothEmpty", func() {
		// Arrange
		a := corestr.New.Hashset.Cap(1)
		b := corestr.New.Hashset.Cap(1)

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}
