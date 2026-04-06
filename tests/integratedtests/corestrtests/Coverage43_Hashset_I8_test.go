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

func Test_I8_HS01_IsEmpty(t *testing.T) {
	safeTest(t, "Test_I8_HS01_IsEmpty", func() {
		h := corestr.New.Hashset.Cap(5)
		actual := args.Map{"result": h.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		h.Add("a")
		actual := args.Map{"result": h.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_I8_HS02_HasItems(t *testing.T) {
	safeTest(t, "Test_I8_HS02_HasItems", func() {
		h := corestr.New.Hashset.Cap(5)
		actual := args.Map{"result": h.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		h.Add("a")
		actual := args.Map{"result": h.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_I8_HS03_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_I8_HS03_HasAnyItem", func() {
		h := corestr.New.Hashset.Cap(5)
		actual := args.Map{"result": h.HasAnyItem()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_I8_HS04_Add(t *testing.T) {
	safeTest(t, "Test_I8_HS04_Add", func() {
		h := corestr.New.Hashset.Cap(5)
		h.Add("a").Add("b")
		actual := args.Map{"result": h.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_HS05_AddBool(t *testing.T) {
	safeTest(t, "Test_I8_HS05_AddBool", func() {
		h := corestr.New.Hashset.Cap(5)
		existed := h.AddBool("a")
		actual := args.Map{"result": existed}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for new key", actual)
		existed2 := h.AddBool("a")
		actual := args.Map{"result": existed2}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for existing key", actual)
	})
}

func Test_I8_HS06_AddPtr(t *testing.T) {
	safeTest(t, "Test_I8_HS06_AddPtr", func() {
		h := corestr.New.Hashset.Cap(5)
		s := "test"
		h.AddPtr(&s)
		actual := args.Map{"result": h.Has("test")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected to have 'test'", actual)
	})
}

func Test_I8_HS07_AddPtrLock(t *testing.T) {
	safeTest(t, "Test_I8_HS07_AddPtrLock", func() {
		h := corestr.New.Hashset.Cap(5)
		s := "test"
		h.AddPtrLock(&s)
		actual := args.Map{"result": h.Has("test")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected to have 'test'", actual)
	})
}

func Test_I8_HS08_AddLock(t *testing.T) {
	safeTest(t, "Test_I8_HS08_AddLock", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddLock("a")
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_HS09_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_I8_HS09_AddNonEmpty", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddNonEmpty("")
		actual := args.Map{"result": h.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		h.AddNonEmpty("a")
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_HS10_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_I8_HS10_AddNonEmptyWhitespace", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddNonEmptyWhitespace("  ")
		actual := args.Map{"result": h.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		h.AddNonEmptyWhitespace("a")
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_HS11_AddIf(t *testing.T) {
	safeTest(t, "Test_I8_HS11_AddIf", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddIf(false, "skip")
		actual := args.Map{"result": h.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		h.AddIf(true, "keep")
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_HS12_AddIfMany(t *testing.T) {
	safeTest(t, "Test_I8_HS12_AddIfMany", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddIfMany(false, "a", "b")
		actual := args.Map{"result": h.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		h.AddIfMany(true, "a", "b")
		actual := args.Map{"result": h.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_HS13_AddFunc(t *testing.T) {
	safeTest(t, "Test_I8_HS13_AddFunc", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddFunc(func() string { return "x" })
		actual := args.Map{"result": h.Has("x")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 'x'", actual)
	})
}

func Test_I8_HS14_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_I8_HS14_AddFuncErr", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_HS15_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_I8_HS15_AddWithWgLock", func() {
		h := corestr.New.Hashset.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddWithWgLock("a", wg)
		wg.Wait()
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_HS16_Adds(t *testing.T) {
	safeTest(t, "Test_I8_HS16_Adds", func() {
		h := corestr.New.Hashset.Cap(5)
		h.Adds("a", "b", "c")
		actual := args.Map{"result": h.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_I8_HS17_AddStrings(t *testing.T) {
	safeTest(t, "Test_I8_HS17_AddStrings", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddStrings([]string{"a", "b"})
		actual := args.Map{"result": h.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_HS18_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_I8_HS18_AddStringsLock", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddStringsLock([]string{"a"})
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_HS19_AddCollection(t *testing.T) {
	safeTest(t, "Test_I8_HS19_AddCollection", func() {
		h := corestr.New.Hashset.Cap(5)
		c := corestr.New.Collection.Strings([]string{"x", "y"})
		h.AddCollection(c)
		actual := args.Map{"result": h.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		h.AddCollection(nil)
		actual := args.Map{"result": h.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 2", actual)
	})
}

func Test_I8_HS20_AddCollections(t *testing.T) {
	safeTest(t, "Test_I8_HS20_AddCollections", func() {
		h := corestr.New.Hashset.Cap(5)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		h.AddCollections(c1, c2, nil)
		actual := args.Map{"result": h.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_HS21_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_I8_HS21_AddHashsetItems", func() {
		h := corestr.New.Hashset.Cap(5)
		other := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.AddHashsetItems(other)
		actual := args.Map{"result": h.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		h.AddHashsetItems(nil)
		actual := args.Map{"result": h.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 2", actual)
	})
}

func Test_I8_HS22_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_I8_HS22_AddItemsMap", func() {
		h := corestr.New.Hashset.Cap(5)
		m := map[string]bool{"a": true, "b": false, "c": true}
		h.AddItemsMap(m)
		actual := args.Map{"result": h.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 (b=false excluded)", actual)
		h.AddItemsMap(nil)
	})
}

func Test_I8_HS23_AddSimpleSlice(t *testing.T) {
	safeTest(t, "Test_I8_HS23_AddSimpleSlice", func() {
		h := corestr.New.Hashset.Cap(5)
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		h.AddSimpleSlice(ss)
		actual := args.Map{"result": h.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// =============================================================================
// Hashset — Query operations
// =============================================================================

func Test_I8_HS24_Has(t *testing.T) {
	safeTest(t, "Test_I8_HS24_Has", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		actual := args.Map{"result": h.Has("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": h.Has("z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_I8_HS25_HasLock(t *testing.T) {
	safeTest(t, "Test_I8_HS25_HasLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"result": h.HasLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_I8_HS26_HasAll(t *testing.T) {
	safeTest(t, "Test_I8_HS26_HasAll", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
		actual := args.Map{"result": h.HasAll("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": h.HasAll("a", "z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_I8_HS27_HasAny(t *testing.T) {
	safeTest(t, "Test_I8_HS27_HasAny", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"result": h.HasAny("a", "z")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": h.HasAny("x", "y")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_I8_HS28_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_I8_HS28_HasAllStrings", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		actual := args.Map{"result": h.HasAllStrings([]string{"a", "b"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_I8_HS29_IsMissing(t *testing.T) {
	safeTest(t, "Test_I8_HS29_IsMissing", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"result": h.IsMissing("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual := args.Map{"result": h.IsMissing("z")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_I8_HS30_IsMissingLock(t *testing.T) {
	safeTest(t, "Test_I8_HS30_IsMissingLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"result": h.IsMissingLock("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_I8_HS31_IsAllMissing(t *testing.T) {
	safeTest(t, "Test_I8_HS31_IsAllMissing", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"result": h.IsAllMissing("x", "y") != true}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": h.IsAllMissing("a", "y") != false}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_I8_HS32_Contains(t *testing.T) {
	safeTest(t, "Test_I8_HS32_Contains", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"result": h.Contains("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_I8_HS33_IsEqual(t *testing.T) {
	safeTest(t, "Test_I8_HS33_IsEqual", func() {
		a := corestr.New.Hashset.Strings([]string{"a", "b"})
		b := corestr.New.Hashset.Strings([]string{"a", "b"})
		actual := args.Map{"result": a.IsEqual(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

// =============================================================================
// Hashset — List, Sort, JSON
// =============================================================================

func Test_I8_HS34_List(t *testing.T) {
	safeTest(t, "Test_I8_HS34_List", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		actual := args.Map{"result": len(h.List()) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_HS35_SortedList(t *testing.T) {
	safeTest(t, "Test_I8_HS35_SortedList", func() {
		h := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
		sorted := h.SortedList()
		actual := args.Map{"result": sorted[0] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a' first", actual)
	})
}

func Test_I8_HS36_OrderedList(t *testing.T) {
	safeTest(t, "Test_I8_HS36_OrderedList", func() {
		h := corestr.New.Hashset.Strings([]string{"c", "a"})
		ordered := h.OrderedList()
		actual := args.Map{"result": ordered[0] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a' first", actual)
	})
}

func Test_I8_HS37_ListPtrSortedAsc(t *testing.T) {
	safeTest(t, "Test_I8_HS37_ListPtrSortedAsc", func() {
		h := corestr.New.Hashset.Strings([]string{"b", "a"})
		list := h.ListPtrSortedAsc()
		actual := args.Map{"result": list[0] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
	})
}

func Test_I8_HS38_ListPtrSortedDsc(t *testing.T) {
	safeTest(t, "Test_I8_HS38_ListPtrSortedDsc", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		list := h.ListPtrSortedDsc()
		actual := args.Map{"result": list[0] != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'b'", actual)
	})
}

func Test_I8_HS39_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_I8_HS39_SimpleSlice", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		ss := h.SimpleSlice()
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_HS40_SafeStrings(t *testing.T) {
	safeTest(t, "Test_I8_HS40_SafeStrings", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"result": len(h.SafeStrings()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_HS41_Collection(t *testing.T) {
	safeTest(t, "Test_I8_HS41_Collection", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		c := h.Collection()
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_HS42_String(t *testing.T) {
	safeTest(t, "Test_I8_HS42_String", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		s := h.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_I8_HS43_StringLock(t *testing.T) {
	safeTest(t, "Test_I8_HS43_StringLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		s := h.StringLock()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_I8_HS44_Json(t *testing.T) {
	safeTest(t, "Test_I8_HS44_Json", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		j := h.Json()
		actual := args.Map{"result": j.JsonString() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_I8_HS45_JsonModel(t *testing.T) {
	safeTest(t, "Test_I8_HS45_JsonModel", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		m := h.JsonModel()
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_HS46_MapStringAny(t *testing.T) {
	safeTest(t, "Test_I8_HS46_MapStringAny", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		m := h.MapStringAny()
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_HS47_JoinSorted(t *testing.T) {
	safeTest(t, "Test_I8_HS47_JoinSorted", func() {
		h := corestr.New.Hashset.Strings([]string{"b", "a"})
		s := h.JoinSorted(",")
		actual := args.Map{"result": s != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
	})
}

// =============================================================================
// Hashset — Remove, Clear, Dispose, Resize
// =============================================================================

func Test_I8_HS48_Remove(t *testing.T) {
	safeTest(t, "Test_I8_HS48_Remove", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.Remove("a")
		actual := args.Map{"result": h.Has("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed", actual)
	})
}

func Test_I8_HS49_SafeRemove(t *testing.T) {
	safeTest(t, "Test_I8_HS49_SafeRemove", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.SafeRemove("a")
		h.SafeRemove("nonexistent")
		actual := args.Map{"result": h.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_I8_HS50_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_I8_HS50_RemoveWithLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.RemoveWithLock("a")
		actual := args.Map{"result": h.Has("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed", actual)
	})
}

func Test_I8_HS51_Clear(t *testing.T) {
	safeTest(t, "Test_I8_HS51_Clear", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.Clear()
		actual := args.Map{"result": h.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_I8_HS52_Dispose(t *testing.T) {
	safeTest(t, "Test_I8_HS52_Dispose", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.Dispose()
		actual := args.Map{"result": h.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_I8_HS53_Resize(t *testing.T) {
	safeTest(t, "Test_I8_HS53_Resize", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		h.Resize(100)
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_HS54_AddCapacities(t *testing.T) {
	safeTest(t, "Test_I8_HS54_AddCapacities", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddCapacities(10, 20)
		_ = h
	})
}

func Test_I8_HS55_ConcatNewHashsets(t *testing.T) {
	safeTest(t, "Test_I8_HS55_ConcatNewHashsets", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		other := corestr.New.Hashset.Strings([]string{"b"})
		result := h.ConcatNewHashsets(true, other)
		actual := args.Map{"result": result.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_I8_HS56_ConcatNewHashsets_Empty(t *testing.T) {
	safeTest(t, "Test_I8_HS56_ConcatNewHashsets_Empty", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.ConcatNewHashsets(true)
		actual := args.Map{"result": result.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_I8_HS57_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_I8_HS57_ConcatNewStrings", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.ConcatNewStrings(true, []string{"b", "c"})
		actual := args.Map{"result": result.Length() < 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	})
}

func Test_I8_HS58_Filter(t *testing.T) {
	safeTest(t, "Test_I8_HS58_Filter", func() {
		h := corestr.New.Hashset.Strings([]string{"abc", "de", "f"})
		filtered := h.Filter(func(s string) bool { return len(s) > 1 })
		actual := args.Map{"result": filtered.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_HS59_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_I8_HS59_AddsUsingFilter", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddsUsingFilter(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		}, "a", "bb", "ccc")
		actual := args.Map{"result": h.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_I8_HS60_AddsAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_I8_HS60_AddsAnyUsingFilter", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddsAnyUsingFilter(func(s string, i int) (string, bool, bool) {
			return s, true, false
		}, "hello", 42, nil)
		actual := args.Map{"result": h.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 (nil skipped)", actual)
	})
}

func Test_I8_HS61_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_I8_HS61_IsEmptyLock", func() {
		h := corestr.New.Hashset.Cap(5)
		actual := args.Map{"result": h.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_I8_HS62_ListCopyLock(t *testing.T) {
	safeTest(t, "Test_I8_HS62_ListCopyLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		list := h.ListCopyLock()
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_HS63_LengthLock(t *testing.T) {
	safeTest(t, "Test_I8_HS63_LengthLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"result": h.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_HS64_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_I8_HS64_ParseInjectUsingJson", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		jr := h.JsonPtr()
		h2 := corestr.New.Hashset.Cap(1)
		_, err := h2.ParseInjectUsingJson(jr)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_I8_HS65_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_I8_HS65_ParseInjectUsingJson_Error", func() {
		h := corestr.New.Hashset.Cap(1)
		bad := corejson.NewResult.UsingString(`invalid`)
		_, err := h.ParseInjectUsingJson(bad)
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_I8_HS66_HasWithLock(t *testing.T) {
	safeTest(t, "Test_I8_HS66_HasWithLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"result": h.HasWithLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_I8_HS67_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_I8_HS67_HasAllCollectionItems", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": h.HasAllCollectionItems(c)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": h.HasAllCollectionItems(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_I8_HS68_MapStringAnyDiff(t *testing.T) {
	safeTest(t, "Test_I8_HS68_MapStringAnyDiff", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		_ = h.MapStringAnyDiff()
	})
}

func Test_I8_HS69_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_I8_HS69_WrapDoubleQuote", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		wrapped := h.WrapDoubleQuote()
		list := wrapped.SortedList()
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_I8_HS70_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_I8_HS70_IsEquals_BothEmpty", func() {
		a := corestr.New.Hashset.Cap(1)
		b := corestr.New.Hashset.Cap(1)
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}
