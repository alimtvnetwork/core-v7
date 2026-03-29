package corestrtests

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — Segment 6: Basic ops, Add variants, Has/Missing, Filter (L1-700)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovHS1_01_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_CovHS1_01_IsEmpty_HasItems", func() {
		hs := corestr.New.Hashset.Empty()
		if !hs.IsEmpty() {
			t.Fatal("expected empty")
		}
		if hs.HasItems() {
			t.Fatal("expected no items")
		}
		hs.Add("a")
		if hs.IsEmpty() {
			t.Fatal("expected not empty")
		}
		if !hs.HasItems() {
			t.Fatal("expected items")
		}
	})
}

func Test_CovHS1_02_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_02_IsEmptyLock", func() {
		hs := corestr.New.Hashset.Empty()
		if !hs.IsEmptyLock() {
			t.Fatal("expected empty")
		}
		hs.Add("x")
		if hs.IsEmptyLock() {
			t.Fatal("expected not empty")
		}
	})
}

func Test_CovHS1_03_AddCapacities(t *testing.T) {
	safeTest(t, "Test_CovHS1_03_AddCapacities", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		// with capacities
		r := hs.AddCapacities(10, 20)
		if r.Length() != 1 {
			t.Fatal("expected 1")
		}
		// empty capacities
		r2 := hs.AddCapacities()
		if r2.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHS1_04_AddCapacitiesLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_04_AddCapacitiesLock", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		r := hs.AddCapacitiesLock(10)
		if r.Length() != 1 {
			t.Fatal("expected 1")
		}
		// empty
		r2 := hs.AddCapacitiesLock()
		if r2.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHS1_05_Resize(t *testing.T) {
	safeTest(t, "Test_CovHS1_05_Resize", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		// capacity > length
		r := hs.Resize(100)
		if r.Length() != 1 {
			t.Fatal("expected 1")
		}
		// capacity < length (no-op)
		r2 := hs.Resize(0)
		if r2.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHS1_06_ResizeLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_06_ResizeLock", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		r := hs.ResizeLock(100)
		if r.Length() != 1 {
			t.Fatal("expected 1")
		}
		r2 := hs.ResizeLock(0)
		if r2.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHS1_07_Collection(t *testing.T) {
	safeTest(t, "Test_CovHS1_07_Collection", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		col := hs.Collection()
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHS1_08_ConcatNewHashsets(t *testing.T) {
	safeTest(t, "Test_CovHS1_08_ConcatNewHashsets", func() {
		a := corestr.New.Hashset.Empty()
		a.Add("a")
		// empty hashsets — clone
		r := a.ConcatNewHashsets(true)
		if r.Length() != 1 {
			t.Fatal("expected 1")
		}
		// with hashsets
		b := corestr.New.Hashset.Empty()
		b.Add("b")
		r2 := a.ConcatNewHashsets(false, b, nil)
		if r2.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_CovHS1_09_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_CovHS1_09_ConcatNewStrings", func() {
		a := corestr.New.Hashset.Empty()
		a.Add("a")
		// empty
		r := a.ConcatNewStrings(true)
		if r.Length() != 1 {
			t.Fatal("expected 1")
		}
		// with strings
		r2 := a.ConcatNewStrings(false, []string{"b", "c"})
		if r2.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_CovHS1_10_AddPtr(t *testing.T) {
	safeTest(t, "Test_CovHS1_10_AddPtr", func() {
		hs := corestr.New.Hashset.Empty()
		s := "hello"
		hs.AddPtr(&s)
		if !hs.Has("hello") {
			t.Fatal("expected has hello")
		}
	})
}

func Test_CovHS1_11_AddPtrLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_11_AddPtrLock", func() {
		hs := corestr.New.Hashset.Empty()
		s := "x"
		hs.AddPtrLock(&s)
		if !hs.Has("x") {
			t.Fatal("expected has x")
		}
	})
}

func Test_CovHS1_12_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_12_AddWithWgLock", func() {
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hs.AddWithWgLock("a", wg)
		wg.Wait()
		if !hs.Has("a") {
			t.Fatal("expected has a")
		}
	})
}

func Test_CovHS1_13_Add_AddBool(t *testing.T) {
	safeTest(t, "Test_CovHS1_13_Add_AddBool", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		// AddBool — new
		existed := hs.AddBool("b")
		if existed {
			t.Fatal("expected false for new")
		}
		// AddBool — existing
		existed2 := hs.AddBool("a")
		if !existed2 {
			t.Fatal("expected true for existing")
		}
	})
}

func Test_CovHS1_14_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_CovHS1_14_AddNonEmpty", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddNonEmpty("")
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
		hs.AddNonEmpty("a")
		if hs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHS1_15_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_CovHS1_15_AddNonEmptyWhitespace", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddNonEmptyWhitespace("  ")
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
		hs.AddNonEmptyWhitespace("a")
		if hs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHS1_16_AddIf(t *testing.T) {
	safeTest(t, "Test_CovHS1_16_AddIf", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddIf(false, "skip")
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
		hs.AddIf(true, "keep")
		if hs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHS1_17_AddIfMany(t *testing.T) {
	safeTest(t, "Test_CovHS1_17_AddIfMany", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddIfMany(false, "a", "b")
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
		hs.AddIfMany(true, "a", "b")
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovHS1_18_AddFunc(t *testing.T) {
	safeTest(t, "Test_CovHS1_18_AddFunc", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddFunc(func() string { return "generated" })
		if !hs.Has("generated") {
			t.Fatal("expected has generated")
		}
	})
}

func Test_CovHS1_19_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_CovHS1_19_AddFuncErr", func() {
		hs := corestr.New.Hashset.Empty()
		// success
		hs.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(err error) { t.Fatal("should not be called") },
		)
		if !hs.Has("ok") {
			t.Fatal("expected has ok")
		}
		// error
		hs.AddFuncErr(
			func() (string, error) { return "", fmt.Errorf("err") },
			func(err error) {},
		)
	})
}

func Test_CovHS1_20_AddStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_20_AddStringsPtrWgLock", func() {
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hs.AddStringsPtrWgLock([]string{"a", "b"}, wg)
		wg.Wait()
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovHS1_21_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_CovHS1_21_AddHashsetItems", func() {
		hs := corestr.New.Hashset.Empty()
		other := corestr.New.Hashset.Empty()
		other.Add("x")
		hs.AddHashsetItems(other)
		if !hs.Has("x") {
			t.Fatal("expected has x")
		}
		// nil
		hs.AddHashsetItems(nil)
	})
}

func Test_CovHS1_22_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_CovHS1_22_AddItemsMap", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddItemsMap(map[string]bool{"a": true, "b": false})
		if !hs.Has("a") {
			t.Fatal("expected has a")
		}
		if hs.Has("b") {
			t.Fatal("should not have b (disabled)")
		}
		// nil
		hs.AddItemsMap(nil)
	})
}

func Test_CovHS1_23_AddItemsMapWgLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_23_AddItemsMapWgLock", func() {
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		m := map[string]bool{"x": true, "y": false}
		hs.AddItemsMapWgLock(&m, wg)
		wg.Wait()
		if !hs.Has("x") {
			t.Fatal("expected x")
		}
		// nil
		hs.AddItemsMapWgLock(nil, wg)
	})
}

func Test_CovHS1_24_AddHashsetWgLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_24_AddHashsetWgLock", func() {
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		other := corestr.New.Hashset.Empty()
		other.Add("z")
		wg.Add(1)
		hs.AddHashsetWgLock(other, wg)
		wg.Wait()
		if !hs.Has("z") {
			t.Fatal("expected z")
		}
		// nil
		hs.AddHashsetWgLock(nil, wg)
	})
}

func Test_CovHS1_25_AddStrings_Adds(t *testing.T) {
	safeTest(t, "Test_CovHS1_25_AddStrings_Adds", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddStrings([]string{"a", "b"})
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
		hs.AddStrings(nil)
		hs.Adds("c")
		if hs.Length() != 3 {
			t.Fatal("expected 3")
		}
		hs.Adds()
	})
}

func Test_CovHS1_26_AddSimpleSlice(t *testing.T) {
	safeTest(t, "Test_CovHS1_26_AddSimpleSlice", func() {
		hs := corestr.New.Hashset.Empty()
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		hs.AddSimpleSlice(ss)
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
		// empty
		hs.AddSimpleSlice(corestr.Empty.SimpleSlice())
	})
}

func Test_CovHS1_27_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_27_AddStringsLock", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddStringsLock([]string{"a"})
		if hs.Length() != 1 {
			t.Fatal("expected 1")
		}
		hs.AddStringsLock(nil)
	})
}

func Test_CovHS1_28_AddCollection_AddCollections(t *testing.T) {
	safeTest(t, "Test_CovHS1_28_AddCollection_AddCollections", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddCollection(corestr.New.Collection.Strings([]string{"a"}))
		if hs.Length() != 1 {
			t.Fatal("expected 1")
		}
		hs.AddCollection(nil)
		hs.AddCollection(corestr.Empty.Collection())

		hs.AddCollections(corestr.New.Collection.Strings([]string{"b"}), nil)
		if !hs.Has("b") {
			t.Fatal("expected b")
		}
		hs.AddCollections()
	})
}

func Test_CovHS1_29_AddsAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_CovHS1_29_AddsAnyUsingFilter", func() {
		hs := corestr.New.Hashset.Empty()
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}
		hs.AddsAnyUsingFilter(filter, "a", nil, "b")
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
		// nil anys
		hs.AddsAnyUsingFilter(filter)
		// break
		hs2 := corestr.New.Hashset.Empty()
		breakFilter := func(str string, index int) (string, bool, bool) {
			return str, true, true
		}
		hs2.AddsAnyUsingFilter(breakFilter, "a", "b")
		if hs2.Length() != 1 {
			t.Fatal("expected 1 (break)")
		}
	})
}

func Test_CovHS1_30_AddsAnyUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_30_AddsAnyUsingFilterLock", func() {
		hs := corestr.New.Hashset.Empty()
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}
		hs.AddsAnyUsingFilterLock(filter, "a", nil)
		if hs.Length() != 1 {
			t.Fatal("expected 1")
		}
		// nil
		hs.AddsAnyUsingFilterLock(filter)
		// break
		breakFilter := func(str string, index int) (string, bool, bool) {
			return str, true, true
		}
		hs.AddsAnyUsingFilterLock(breakFilter, "x", "y")
	})
}

func Test_CovHS1_31_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_CovHS1_31_AddsUsingFilter", func() {
		hs := corestr.New.Hashset.Empty()
		filter := func(str string, index int) (string, bool, bool) {
			return str, str != "skip", false
		}
		hs.AddsUsingFilter(filter, "a", "skip", "b")
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
		// nil
		hs.AddsUsingFilter(filter)
		// break
		breakFilter := func(str string, index int) (string, bool, bool) {
			return str, true, true
		}
		hs2 := corestr.New.Hashset.Empty()
		hs2.AddsUsingFilter(breakFilter, "a", "b")
		if hs2.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHS1_32_AddLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_32_AddLock", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddLock("a")
		if !hs.Has("a") {
			t.Fatal("expected a")
		}
	})
}

func Test_CovHS1_33_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_CovHS1_33_HasAnyItem", func() {
		hs := corestr.New.Hashset.Empty()
		if hs.HasAnyItem() {
			t.Fatal("expected false")
		}
		hs.Add("a")
		if !hs.HasAnyItem() {
			t.Fatal("expected true")
		}
	})
}

func Test_CovHS1_34_IsMissing_IsMissingLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_34_IsMissing_IsMissingLock", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		if hs.IsMissing("a") {
			t.Fatal("expected found")
		}
		if !hs.IsMissing("b") {
			t.Fatal("expected missing")
		}
		if hs.IsMissingLock("a") {
			t.Fatal("expected found")
		}
	})
}

func Test_CovHS1_35_Has_Contains_HasLock_HasWithLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_35_Has_Contains_HasLock_HasWithLock", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		if !hs.Has("a") {
			t.Fatal("expected true")
		}
		if !hs.Contains("a") {
			t.Fatal("expected true")
		}
		if !hs.HasLock("a") {
			t.Fatal("expected true")
		}
		if !hs.HasWithLock("a") {
			t.Fatal("expected true")
		}
		if hs.Has("missing") {
			t.Fatal("expected false")
		}
	})
}

func Test_CovHS1_36_HasAllStrings_HasAll(t *testing.T) {
	safeTest(t, "Test_CovHS1_36_HasAllStrings_HasAll", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b", "c")
		if !hs.HasAllStrings([]string{"a", "b"}) {
			t.Fatal("expected true")
		}
		if hs.HasAllStrings([]string{"a", "x"}) {
			t.Fatal("expected false")
		}
		if !hs.HasAll("a", "c") {
			t.Fatal("expected true")
		}
		if hs.HasAll("a", "z") {
			t.Fatal("expected false")
		}
	})
}

func Test_CovHS1_37_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_CovHS1_37_HasAllCollectionItems", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		if !hs.HasAllCollectionItems(corestr.New.Collection.Strings([]string{"a", "b"})) {
			t.Fatal("expected true")
		}
		if hs.HasAllCollectionItems(nil) {
			t.Fatal("expected false for nil")
		}
		if hs.HasAllCollectionItems(corestr.Empty.Collection()) {
			t.Fatal("expected false for empty")
		}
	})
}

func Test_CovHS1_38_IsAllMissing(t *testing.T) {
	safeTest(t, "Test_CovHS1_38_IsAllMissing", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		if !hs.IsAllMissing("x", "y") {
			t.Fatal("expected true")
		}
		if hs.IsAllMissing("a", "y") {
			t.Fatal("expected false")
		}
	})
}

func Test_CovHS1_39_HasAny(t *testing.T) {
	safeTest(t, "Test_CovHS1_39_HasAny", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		if !hs.HasAny("x", "a") {
			t.Fatal("expected true")
		}
		if hs.HasAny("x", "y") {
			t.Fatal("expected false")
		}
	})
}

func Test_CovHS1_40_IsEqual_IsEquals(t *testing.T) {
	safeTest(t, "Test_CovHS1_40_IsEqual_IsEquals", func() {
		a := corestr.New.Hashset.Empty()
		a.Adds("a", "b")
		b := corestr.New.Hashset.Empty()
		b.Adds("a", "b")
		if !a.IsEqual(b) {
			t.Fatal("expected equal")
		}
		// same ptr
		if !a.IsEqual(a) {
			t.Fatal("expected equal to self")
		}
		// nil
		if a.IsEquals(nil) {
			t.Fatal("expected false for nil")
		}
		// both empty
		e1 := corestr.New.Hashset.Empty()
		e2 := corestr.New.Hashset.Empty()
		if !e1.IsEquals(e2) {
			t.Fatal("expected empty == empty")
		}
		// one empty
		if a.IsEquals(e1) {
			t.Fatal("expected false")
		}
		// diff length
		c := corestr.New.Hashset.Empty()
		c.Add("a")
		if a.IsEquals(c) {
			t.Fatal("expected false for diff length")
		}
		// same length, diff keys
		d := corestr.New.Hashset.Empty()
		d.Adds("a", "z")
		if a.IsEquals(d) {
			t.Fatal("expected false for diff keys")
		}
	})
}

func Test_CovHS1_41_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_41_IsEqualsLock", func() {
		a := corestr.New.Hashset.Empty()
		a.Add("a")
		b := corestr.New.Hashset.Empty()
		b.Add("a")
		if !a.IsEqualsLock(b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_CovHS1_42_SortedList(t *testing.T) {
	safeTest(t, "Test_CovHS1_42_SortedList", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("c", "a", "b")
		sorted := hs.SortedList()
		if sorted[0] != "a" || sorted[1] != "b" || sorted[2] != "c" {
			t.Fatal("expected sorted asc")
		}
	})
}

func Test_CovHS1_43_Filter(t *testing.T) {
	safeTest(t, "Test_CovHS1_43_Filter", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("apple", "banana", "avocado")
		result := hs.Filter(func(s string) bool {
			return s[0] == 'a'
		})
		if result.Length() != 2 {
			t.Fatalf("expected 2, got %d", result.Length())
		}
	})
}
