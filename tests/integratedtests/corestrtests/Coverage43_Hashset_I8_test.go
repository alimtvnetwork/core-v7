package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// =============================================================================
// Hashset — Core operations
// =============================================================================

func Test_I8_HS01_IsEmpty(t *testing.T) {
	safeTest(t, "Test_I8_HS01_IsEmpty", func() {
		h := corestr.New.Hashset.Cap(5)
		if !h.IsEmpty() {
			t.Fatal("expected empty")
		}
		h.Add("a")
		if h.IsEmpty() {
			t.Fatal("expected not empty")
		}
	})
}

func Test_I8_HS02_HasItems(t *testing.T) {
	safeTest(t, "Test_I8_HS02_HasItems", func() {
		h := corestr.New.Hashset.Cap(5)
		if h.HasItems() {
			t.Fatal("expected false")
		}
		h.Add("a")
		if !h.HasItems() {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_HS03_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_I8_HS03_HasAnyItem", func() {
		h := corestr.New.Hashset.Cap(5)
		if h.HasAnyItem() {
			t.Fatal("expected false")
		}
	})
}

func Test_I8_HS04_Add(t *testing.T) {
	safeTest(t, "Test_I8_HS04_Add", func() {
		h := corestr.New.Hashset.Cap(5)
		h.Add("a").Add("b")
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_HS05_AddBool(t *testing.T) {
	safeTest(t, "Test_I8_HS05_AddBool", func() {
		h := corestr.New.Hashset.Cap(5)
		existed := h.AddBool("a")
		if existed {
			t.Fatal("expected false for new key")
		}
		existed2 := h.AddBool("a")
		if !existed2 {
			t.Fatal("expected true for existing key")
		}
	})
}

func Test_I8_HS06_AddPtr(t *testing.T) {
	safeTest(t, "Test_I8_HS06_AddPtr", func() {
		h := corestr.New.Hashset.Cap(5)
		s := "test"
		h.AddPtr(&s)
		if !h.Has("test") {
			t.Fatal("expected to have 'test'")
		}
	})
}

func Test_I8_HS07_AddPtrLock(t *testing.T) {
	safeTest(t, "Test_I8_HS07_AddPtrLock", func() {
		h := corestr.New.Hashset.Cap(5)
		s := "test"
		h.AddPtrLock(&s)
		if !h.Has("test") {
			t.Fatal("expected to have 'test'")
		}
	})
}

func Test_I8_HS08_AddLock(t *testing.T) {
	safeTest(t, "Test_I8_HS08_AddLock", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddLock("a")
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HS09_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_I8_HS09_AddNonEmpty", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddNonEmpty("")
		if h.Length() != 0 {
			t.Fatal("expected 0")
		}
		h.AddNonEmpty("a")
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HS10_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_I8_HS10_AddNonEmptyWhitespace", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddNonEmptyWhitespace("  ")
		if h.Length() != 0 {
			t.Fatal("expected 0")
		}
		h.AddNonEmptyWhitespace("a")
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HS11_AddIf(t *testing.T) {
	safeTest(t, "Test_I8_HS11_AddIf", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddIf(false, "skip")
		if h.Length() != 0 {
			t.Fatal("expected 0")
		}
		h.AddIf(true, "keep")
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HS12_AddIfMany(t *testing.T) {
	safeTest(t, "Test_I8_HS12_AddIfMany", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddIfMany(false, "a", "b")
		if h.Length() != 0 {
			t.Fatal("expected 0")
		}
		h.AddIfMany(true, "a", "b")
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_HS13_AddFunc(t *testing.T) {
	safeTest(t, "Test_I8_HS13_AddFunc", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddFunc(func() string { return "x" })
		if !h.Has("x") {
			t.Fatal("expected 'x'")
		}
	})
}

func Test_I8_HS14_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_I8_HS14_AddFuncErr", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HS15_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_I8_HS15_AddWithWgLock", func() {
		h := corestr.New.Hashset.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddWithWgLock("a", wg)
		wg.Wait()
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HS16_Adds(t *testing.T) {
	safeTest(t, "Test_I8_HS16_Adds", func() {
		h := corestr.New.Hashset.Cap(5)
		h.Adds("a", "b", "c")
		if h.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_I8_HS17_AddStrings(t *testing.T) {
	safeTest(t, "Test_I8_HS17_AddStrings", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddStrings([]string{"a", "b"})
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_HS18_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_I8_HS18_AddStringsLock", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddStringsLock([]string{"a"})
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HS19_AddCollection(t *testing.T) {
	safeTest(t, "Test_I8_HS19_AddCollection", func() {
		h := corestr.New.Hashset.Cap(5)
		c := corestr.New.Collection.Strings([]string{"x", "y"})
		h.AddCollection(c)
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
		h.AddCollection(nil)
		if h.Length() != 2 {
			t.Fatal("expected still 2")
		}
	})
}

func Test_I8_HS20_AddCollections(t *testing.T) {
	safeTest(t, "Test_I8_HS20_AddCollections", func() {
		h := corestr.New.Hashset.Cap(5)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		h.AddCollections(c1, c2, nil)
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_HS21_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_I8_HS21_AddHashsetItems", func() {
		h := corestr.New.Hashset.Cap(5)
		other := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.AddHashsetItems(other)
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
		h.AddHashsetItems(nil)
		if h.Length() != 2 {
			t.Fatal("expected still 2")
		}
	})
}

func Test_I8_HS22_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_I8_HS22_AddItemsMap", func() {
		h := corestr.New.Hashset.Cap(5)
		m := map[string]bool{"a": true, "b": false, "c": true}
		h.AddItemsMap(m)
		if h.Length() != 2 {
			t.Fatal("expected 2 (b=false excluded)")
		}
		h.AddItemsMap(nil)
	})
}

func Test_I8_HS23_AddSimpleSlice(t *testing.T) {
	safeTest(t, "Test_I8_HS23_AddSimpleSlice", func() {
		h := corestr.New.Hashset.Cap(5)
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		h.AddSimpleSlice(ss)
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

// =============================================================================
// Hashset — Query operations
// =============================================================================

func Test_I8_HS24_Has(t *testing.T) {
	safeTest(t, "Test_I8_HS24_Has", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		if !h.Has("a") {
			t.Fatal("expected true")
		}
		if h.Has("z") {
			t.Fatal("expected false")
		}
	})
}

func Test_I8_HS25_HasLock(t *testing.T) {
	safeTest(t, "Test_I8_HS25_HasLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		if !h.HasLock("a") {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_HS26_HasAll(t *testing.T) {
	safeTest(t, "Test_I8_HS26_HasAll", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
		if !h.HasAll("a", "b") {
			t.Fatal("expected true")
		}
		if h.HasAll("a", "z") {
			t.Fatal("expected false")
		}
	})
}

func Test_I8_HS27_HasAny(t *testing.T) {
	safeTest(t, "Test_I8_HS27_HasAny", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		if !h.HasAny("a", "z") {
			t.Fatal("expected true")
		}
		if h.HasAny("x", "y") {
			t.Fatal("expected false")
		}
	})
}

func Test_I8_HS28_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_I8_HS28_HasAllStrings", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		if !h.HasAllStrings([]string{"a", "b"}) {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_HS29_IsMissing(t *testing.T) {
	safeTest(t, "Test_I8_HS29_IsMissing", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		if h.IsMissing("a") {
			t.Fatal("expected false")
		}
		if !h.IsMissing("z") {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_HS30_IsMissingLock(t *testing.T) {
	safeTest(t, "Test_I8_HS30_IsMissingLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		if h.IsMissingLock("a") {
			t.Fatal("expected false")
		}
	})
}

func Test_I8_HS31_IsAllMissing(t *testing.T) {
	safeTest(t, "Test_I8_HS31_IsAllMissing", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		if h.IsAllMissing("x", "y") != true {
			t.Fatal("expected true")
		}
		if h.IsAllMissing("a", "y") != false {
			t.Fatal("expected false")
		}
	})
}

func Test_I8_HS32_Contains(t *testing.T) {
	safeTest(t, "Test_I8_HS32_Contains", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		if !h.Contains("a") {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_HS33_IsEqual(t *testing.T) {
	safeTest(t, "Test_I8_HS33_IsEqual", func() {
		a := corestr.New.Hashset.Strings([]string{"a", "b"})
		b := corestr.New.Hashset.Strings([]string{"a", "b"})
		if !a.IsEqual(b) {
			t.Fatal("expected equal")
		}
	})
}

// =============================================================================
// Hashset — List, Sort, JSON
// =============================================================================

func Test_I8_HS34_List(t *testing.T) {
	safeTest(t, "Test_I8_HS34_List", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		if len(h.List()) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_HS35_SortedList(t *testing.T) {
	safeTest(t, "Test_I8_HS35_SortedList", func() {
		h := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
		sorted := h.SortedList()
		if sorted[0] != "a" {
			t.Fatal("expected 'a' first")
		}
	})
}

func Test_I8_HS36_OrderedList(t *testing.T) {
	safeTest(t, "Test_I8_HS36_OrderedList", func() {
		h := corestr.New.Hashset.Strings([]string{"c", "a"})
		ordered := h.OrderedList()
		if ordered[0] != "a" {
			t.Fatal("expected 'a' first")
		}
	})
}

func Test_I8_HS37_ListPtrSortedAsc(t *testing.T) {
	safeTest(t, "Test_I8_HS37_ListPtrSortedAsc", func() {
		h := corestr.New.Hashset.Strings([]string{"b", "a"})
		list := h.ListPtrSortedAsc()
		if list[0] != "a" {
			t.Fatal("expected 'a'")
		}
	})
}

func Test_I8_HS38_ListPtrSortedDsc(t *testing.T) {
	safeTest(t, "Test_I8_HS38_ListPtrSortedDsc", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		list := h.ListPtrSortedDsc()
		if list[0] != "b" {
			t.Fatal("expected 'b'")
		}
	})
}

func Test_I8_HS39_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_I8_HS39_SimpleSlice", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		ss := h.SimpleSlice()
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HS40_SafeStrings(t *testing.T) {
	safeTest(t, "Test_I8_HS40_SafeStrings", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		if len(h.SafeStrings()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HS41_Collection(t *testing.T) {
	safeTest(t, "Test_I8_HS41_Collection", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		c := h.Collection()
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HS42_String(t *testing.T) {
	safeTest(t, "Test_I8_HS42_String", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		s := h.String()
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_I8_HS43_StringLock(t *testing.T) {
	safeTest(t, "Test_I8_HS43_StringLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		s := h.StringLock()
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_I8_HS44_Json(t *testing.T) {
	safeTest(t, "Test_I8_HS44_Json", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		j := h.Json()
		if j.JsonString() == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_I8_HS45_JsonModel(t *testing.T) {
	safeTest(t, "Test_I8_HS45_JsonModel", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		m := h.JsonModel()
		if len(m) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HS46_MapStringAny(t *testing.T) {
	safeTest(t, "Test_I8_HS46_MapStringAny", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		m := h.MapStringAny()
		if len(m) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HS47_JoinSorted(t *testing.T) {
	safeTest(t, "Test_I8_HS47_JoinSorted", func() {
		h := corestr.New.Hashset.Strings([]string{"b", "a"})
		s := h.JoinSorted(",")
		if s != "a,b" {
			t.Fatalf("expected 'a,b', got '%s'", s)
		}
	})
}

// =============================================================================
// Hashset — Remove, Clear, Dispose, Resize
// =============================================================================

func Test_I8_HS48_Remove(t *testing.T) {
	safeTest(t, "Test_I8_HS48_Remove", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.Remove("a")
		if h.Has("a") {
			t.Fatal("expected removed")
		}
	})
}

func Test_I8_HS49_SafeRemove(t *testing.T) {
	safeTest(t, "Test_I8_HS49_SafeRemove", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.SafeRemove("a")
		h.SafeRemove("nonexistent")
		if h.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_I8_HS50_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_I8_HS50_RemoveWithLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.RemoveWithLock("a")
		if h.Has("a") {
			t.Fatal("expected removed")
		}
	})
}

func Test_I8_HS51_Clear(t *testing.T) {
	safeTest(t, "Test_I8_HS51_Clear", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.Clear()
		if h.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_I8_HS52_Dispose(t *testing.T) {
	safeTest(t, "Test_I8_HS52_Dispose", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.Dispose()
		if h.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_I8_HS53_Resize(t *testing.T) {
	safeTest(t, "Test_I8_HS53_Resize", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		h.Resize(100)
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
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
		if result.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_I8_HS56_ConcatNewHashsets_Empty(t *testing.T) {
	safeTest(t, "Test_I8_HS56_ConcatNewHashsets_Empty", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.ConcatNewHashsets(true)
		if result.Length() < 1 {
			t.Fatal("expected at least 1")
		}
	})
}

func Test_I8_HS57_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_I8_HS57_ConcatNewStrings", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.ConcatNewStrings(true, []string{"b", "c"})
		if result.Length() < 3 {
			t.Fatal("expected at least 3")
		}
	})
}

func Test_I8_HS58_Filter(t *testing.T) {
	safeTest(t, "Test_I8_HS58_Filter", func() {
		h := corestr.New.Hashset.Strings([]string{"abc", "de", "f"})
		filtered := h.Filter(func(s string) bool { return len(s) > 1 })
		if filtered.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_HS59_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_I8_HS59_AddsUsingFilter", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddsUsingFilter(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		}, "a", "bb", "ccc")
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_HS60_AddsAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_I8_HS60_AddsAnyUsingFilter", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddsAnyUsingFilter(func(s string, i int) (string, bool, bool) {
			return s, true, false
		}, "hello", 42, nil)
		if h.Length() != 2 {
			t.Fatal("expected 2 (nil skipped)")
		}
	})
}

func Test_I8_HS61_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_I8_HS61_IsEmptyLock", func() {
		h := corestr.New.Hashset.Cap(5)
		if !h.IsEmptyLock() {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_HS62_ListCopyLock(t *testing.T) {
	safeTest(t, "Test_I8_HS62_ListCopyLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		list := h.ListCopyLock()
		if len(list) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HS63_LengthLock(t *testing.T) {
	safeTest(t, "Test_I8_HS63_LengthLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		if h.LengthLock() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HS64_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_I8_HS64_ParseInjectUsingJson", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		jr := h.JsonPtr()
		h2 := corestr.New.Hashset.Cap(1)
		_, err := h2.ParseInjectUsingJson(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_I8_HS65_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_I8_HS65_ParseInjectUsingJson_Error", func() {
		h := corestr.New.Hashset.Cap(1)
		bad := corejson.NewResult.UsingString(`invalid`)
		_, err := h.ParseInjectUsingJson(bad)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_I8_HS66_HasWithLock(t *testing.T) {
	safeTest(t, "Test_I8_HS66_HasWithLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		if !h.HasWithLock("a") {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_HS67_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_I8_HS67_HasAllCollectionItems", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !h.HasAllCollectionItems(c) {
			t.Fatal("expected true")
		}
		if h.HasAllCollectionItems(nil) {
			t.Fatal("expected false for nil")
		}
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
		if len(list) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HS70_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_I8_HS70_IsEquals_BothEmpty", func() {
		a := corestr.New.Hashset.Cap(1)
		b := corestr.New.Hashset.Cap(1)
		if !a.IsEquals(b) {
			t.Fatal("expected equal")
		}
	})
}
