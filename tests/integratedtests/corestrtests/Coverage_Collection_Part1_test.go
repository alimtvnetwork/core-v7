package corestrtests

import (
	"errors"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Segment 1: Basic methods, Add variants, Remove, Capacity, Equals
// Covers ~200 uncovered statements from Collection.go lines 27-700
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovP1_01_JsonString(t *testing.T) {
	safeTest(t, "Test_CovP1_01_JsonString", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		s := col.JsonString()
		if s == "" {
			t.Fatal("expected non-empty JSON string")
		}
	})
}

func Test_CovP1_02_JsonStringMust(t *testing.T) {
	safeTest(t, "Test_CovP1_02_JsonStringMust", func() {
		col := corestr.New.Collection.Strings([]string{"x"})
		s := col.JsonStringMust()
		if s == "" {
			t.Fatal("expected non-empty JSON string")
		}
	})
}

func Test_CovP1_03_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_CovP1_03_HasAnyItem", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		if !col.HasAnyItem() {
			t.Fatal("expected true")
		}
		empty := corestr.Empty.Collection()
		if empty.HasAnyItem() {
			t.Fatal("expected false")
		}
	})
}

func Test_CovP1_04_LastIndex(t *testing.T) {
	safeTest(t, "Test_CovP1_04_LastIndex", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		if col.LastIndex() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP1_05_HasIndex(t *testing.T) {
	safeTest(t, "Test_CovP1_05_HasIndex", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		if !col.HasIndex(0) {
			t.Fatal("expected true for index 0")
		}
		if !col.HasIndex(1) {
			t.Fatal("expected true for index 1")
		}
		if col.HasIndex(2) {
			t.Fatal("expected false for index 2")
		}
		if col.HasIndex(-1) {
			t.Fatal("expected false for -1")
		}
	})
}

func Test_CovP1_06_ListStringsPtr(t *testing.T) {
	safeTest(t, "Test_CovP1_06_ListStringsPtr", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		if len(col.ListStringsPtr()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovP1_07_ListStrings(t *testing.T) {
	safeTest(t, "Test_CovP1_07_ListStrings", func() {
		col := corestr.New.Collection.Strings([]string{"x", "y"})
		if len(col.ListStrings()) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP1_08_StringJSON(t *testing.T) {
	safeTest(t, "Test_CovP1_08_StringJSON", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		if col.StringJSON() == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_CovP1_09_RemoveAt(t *testing.T) {
	safeTest(t, "Test_CovP1_09_RemoveAt", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok := col.RemoveAt(1)
		if !ok {
			t.Fatal("expected success")
		}
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
		// out of range
		ok = col.RemoveAt(-1)
		if ok {
			t.Fatal("expected false for -1")
		}
		ok = col.RemoveAt(100)
		if ok {
			t.Fatal("expected false for 100")
		}
	})
}

func Test_CovP1_10_Count(t *testing.T) {
	safeTest(t, "Test_CovP1_10_Count", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		if col.Count() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovP1_11_Capacity(t *testing.T) {
	safeTest(t, "Test_CovP1_11_Capacity", func() {
		col := corestr.New.Collection.Cap(10)
		if col.Capacity() < 10 {
			t.Fatal("expected at least 10")
		}
		empty := corestr.Empty.Collection()
		_ = empty.Capacity()
	})
}

func Test_CovP1_12_Length(t *testing.T) {
	safeTest(t, "Test_CovP1_12_Length", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP1_13_LengthLock(t *testing.T) {
	safeTest(t, "Test_CovP1_13_LengthLock", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		if col.LengthLock() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovP1_14_IsCollectionPrecheckEqual_AllPaths(t *testing.T) {
	safeTest(t, "Test_CovP1_14_IsCollectionPrecheckEqual_AllPaths", func() {
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})

		// same content
		if !a.IsEquals(b) {
			t.Fatal("expected equal")
		}

		// same pointer
		if !a.IsEquals(a) {
			t.Fatal("expected equal to self")
		}

		// both empty
		e1 := corestr.Empty.Collection()
		e2 := corestr.Empty.Collection()
		if !e1.IsEquals(e2) {
			t.Fatal("expected empty equals empty")
		}

		// one empty
		if a.IsEquals(e1) {
			t.Fatal("expected not equal")
		}

		// diff length
		c := corestr.New.Collection.Strings([]string{"a"})
		if a.IsEquals(c) {
			t.Fatal("expected not equal for diff length")
		}
	})
}

func Test_CovP1_15_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_CovP1_15_IsEqualsWithSensitive_CaseInsensitive", func() {
		a := corestr.New.Collection.Strings([]string{"Hello", "World"})
		b := corestr.New.Collection.Strings([]string{"hello", "world"})

		if a.IsEqualsWithSensitive(false, b) != true {
			t.Fatal("expected case-insensitive equal")
		}
		if a.IsEqualsWithSensitive(true, b) != false {
			t.Fatal("expected case-sensitive not equal")
		}

		// mismatch fold
		c := corestr.New.Collection.Strings([]string{"hello", "OTHER"})
		if a.IsEqualsWithSensitive(false, c) {
			t.Fatal("expected not equal")
		}

		// mismatch sensitive
		d := corestr.New.Collection.Strings([]string{"Hello", "other"})
		if a.IsEqualsWithSensitive(true, d) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_CovP1_16_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_CovP1_16_IsEmptyLock", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		if col.IsEmptyLock() {
			t.Fatal("expected not empty")
		}
		empty := corestr.Empty.Collection()
		if !empty.IsEmptyLock() {
			t.Fatal("expected empty")
		}
	})
}

func Test_CovP1_17_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_CovP1_17_IsEmpty_HasItems", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		if col.IsEmpty() {
			t.Fatal("expected not empty")
		}
		if !col.HasItems() {
			t.Fatal("expected has items")
		}
	})
}

func Test_CovP1_18_AddLock(t *testing.T) {
	safeTest(t, "Test_CovP1_18_AddLock", func() {
		col := corestr.New.Collection.Cap(2)
		col.AddLock("a")
		col.AddLock("b")
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP1_19_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_CovP1_19_AddNonEmpty", func() {
		col := corestr.New.Collection.Cap(2)
		col.AddNonEmpty("")
		if col.Length() != 0 {
			t.Fatal("expected 0 for empty string")
		}
		col.AddNonEmpty("a")
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovP1_20_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_CovP1_20_AddNonEmptyWhitespace", func() {
		col := corestr.New.Collection.Cap(2)
		col.AddNonEmptyWhitespace("   ")
		if col.Length() != 0 {
			t.Fatal("expected 0 for whitespace")
		}
		col.AddNonEmptyWhitespace("a")
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovP1_21_AddError(t *testing.T) {
	safeTest(t, "Test_CovP1_21_AddError", func() {
		col := corestr.New.Collection.Cap(2)
		col.AddError(nil)
		if col.Length() != 0 {
			t.Fatal("expected 0 for nil error")
		}
		col.AddError(errors.New("test"))
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovP1_22_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_CovP1_22_AsDefaultError", func() {
		col := corestr.New.Collection.Strings([]string{"err1", "err2"})
		err := col.AsDefaultError()
		if err == nil {
			t.Fatal("expected non-nil error")
		}
	})
}

func Test_CovP1_23_AsError(t *testing.T) {
	safeTest(t, "Test_CovP1_23_AsError", func() {
		empty := corestr.Empty.Collection()
		if empty.AsError(",") != nil {
			t.Fatal("expected nil for empty")
		}
		col := corestr.New.Collection.Strings([]string{"a"})
		if col.AsError(",") == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_CovP1_24_AddIf(t *testing.T) {
	safeTest(t, "Test_CovP1_24_AddIf", func() {
		col := corestr.New.Collection.Cap(2)
		col.AddIf(false, "skip")
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
		col.AddIf(true, "keep")
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovP1_25_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_CovP1_25_EachItemSplitBy", func() {
		col := corestr.New.Collection.Strings([]string{"a,b", "c,d"})
		result := col.EachItemSplitBy(",")
		if len(result) != 4 {
			t.Fatalf("expected 4, got %d", len(result))
		}
	})
}

func Test_CovP1_26_ConcatNew(t *testing.T) {
	safeTest(t, "Test_CovP1_26_ConcatNew", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		// no addingStrings
		newCol := col.ConcatNew(0)
		if newCol.Length() != 2 {
			t.Fatal("expected 2")
		}
		// with addingStrings
		newCol2 := col.ConcatNew(0, "c", "d")
		if newCol2.Length() != 4 {
			t.Fatal("expected 4")
		}
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
		col := corestr.New.Collection.Cap(5)
		col.AddIfMany(false, "a", "b")
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
		col.AddIfMany(true, "a", "b")
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP1_30_AddFunc(t *testing.T) {
	safeTest(t, "Test_CovP1_30_AddFunc", func() {
		col := corestr.New.Collection.Cap(2)
		col.AddFunc(func() string { return "computed" })
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovP1_31_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_CovP1_31_AddFuncErr", func() {
		col := corestr.New.Collection.Cap(2)
		// success
		col.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(err error) { t.Fatal("unexpected error handler") },
		)
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
		// error
		errCalled := false
		col.AddFuncErr(
			func() (string, error) { return "", errors.New("fail") },
			func(err error) { errCalled = true },
		)
		if !errCalled {
			t.Fatal("expected error handler called")
		}
		if col.Length() != 1 {
			t.Fatal("expected still 1")
		}
	})
}

func Test_CovP1_32_AddsLock(t *testing.T) {
	safeTest(t, "Test_CovP1_32_AddsLock", func() {
		col := corestr.New.Collection.Cap(5)
		col.AddsLock("a", "b")
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP1_33_Adds(t *testing.T) {
	safeTest(t, "Test_CovP1_33_Adds", func() {
		col := corestr.New.Collection.Cap(5)
		col.Adds("a", "b", "c")
		if col.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_CovP1_34_AddStrings(t *testing.T) {
	safeTest(t, "Test_CovP1_34_AddStrings", func() {
		col := corestr.New.Collection.Cap(5)
		col.AddStrings([]string{"x", "y"})
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP1_35_AddCollection(t *testing.T) {
	safeTest(t, "Test_CovP1_35_AddCollection", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"b", "c"})
		a.AddCollection(b)
		if a.Length() != 3 {
			t.Fatal("expected 3")
		}
		// empty collection
		a.AddCollection(corestr.Empty.Collection())
		if a.Length() != 3 {
			t.Fatal("expected still 3")
		}
	})
}

func Test_CovP1_36_AddCollections(t *testing.T) {
	safeTest(t, "Test_CovP1_36_AddCollections", func() {
		col := corestr.New.Collection.Cap(10)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		col.AddCollections(c1, c2, corestr.Empty.Collection())
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP1_37_AddPointerCollectionsLock(t *testing.T) {
	safeTest(t, "Test_CovP1_37_AddPointerCollectionsLock", func() {
		col := corestr.New.Collection.Cap(10)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		col.AddPointerCollectionsLock(c1)
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovP1_38_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_CovP1_38_AddHashmapsValues", func() {
		col := corestr.New.Collection.Cap(10)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "v1")
		col.AddHashmapsValues(hm)
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
		// nil hashmaps
		col.AddHashmapsValues()
		// nil hashmap item
		col.AddHashmapsValues(nil)
	})
}

func Test_CovP1_39_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_CovP1_39_AddHashmapsKeys", func() {
		col := corestr.New.Collection.Cap(10)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "v1")
		col.AddHashmapsKeys(hm)
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
		// nil
		col.AddHashmapsKeys()
		col.AddHashmapsKeys(nil)
	})
}

func Test_CovP1_40_AddHashmapsKeysValues(t *testing.T) {
	safeTest(t, "Test_CovP1_40_AddHashmapsKeysValues", func() {
		col := corestr.New.Collection.Cap(10)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "v1")
		col.AddHashmapsKeysValues(hm)
		if col.Length() != 2 {
			t.Fatalf("expected 2, got %d", col.Length())
		}
		// nil
		col2 := corestr.New.Collection.Cap(10)
		col2.AddHashmapsKeysValues()
		col2.AddHashmapsKeysValues(nil)
	})
}

func Test_CovP1_41_AddHashmapsKeysValuesUsingFilter(t *testing.T) {
	safeTest(t, "Test_CovP1_41_AddHashmapsKeysValuesUsingFilter", func() {
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
		if col.Length() < 1 {
			t.Fatal("expected at least 1")
		}
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
		col := corestr.New.Collection.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		col.AddWithWgLock(wg, "item")
		wg.Wait()
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovP1_43_IndexAt(t *testing.T) {
	safeTest(t, "Test_CovP1_43_IndexAt", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		if col.IndexAt(0) != "a" {
			t.Fatal("expected a")
		}
	})
}

func Test_CovP1_44_SafeIndexAtUsingLength(t *testing.T) {
	safeTest(t, "Test_CovP1_44_SafeIndexAtUsingLength", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		if col.SafeIndexAtUsingLength("def", 2, 0) != "a" {
			t.Fatal("expected a")
		}
		if col.SafeIndexAtUsingLength("def", 2, 5) != "def" {
			t.Fatal("expected default")
		}
	})
}

func Test_CovP1_45_First(t *testing.T) {
	safeTest(t, "Test_CovP1_45_First", func() {
		col := corestr.New.Collection.Strings([]string{"first", "second"})
		if col.First() != "first" {
			t.Fatal("expected first")
		}
	})
}

func Test_CovP1_46_Single(t *testing.T) {
	safeTest(t, "Test_CovP1_46_Single", func() {
		col := corestr.New.Collection.Strings([]string{"only"})
		if col.Single() != "only" {
			t.Fatal("expected only")
		}
		// panic on non-single
		defer func() { recover() }()
		col2 := corestr.New.Collection.Strings([]string{"a", "b"})
		col2.Single()
		t.Fatal("expected panic")
	})
}

func Test_CovP1_47_Last(t *testing.T) {
	safeTest(t, "Test_CovP1_47_Last", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "last"})
		if col.Last() != "last" {
			t.Fatal("expected last")
		}
	})
}

func Test_CovP1_48_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_CovP1_48_LastOrDefault", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		if col.LastOrDefault() != "b" {
			t.Fatal("expected b")
		}
		empty := corestr.Empty.Collection()
		if empty.LastOrDefault() != "" {
			t.Fatal("expected empty string")
		}
	})
}

func Test_CovP1_49_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_CovP1_49_FirstOrDefault", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		if col.FirstOrDefault() != "a" {
			t.Fatal("expected a")
		}
		empty := corestr.Empty.Collection()
		if empty.FirstOrDefault() != "" {
			t.Fatal("expected empty string")
		}
	})
}

func Test_CovP1_50_ResizeForItems_LargeResize(t *testing.T) {
	safeTest(t, "Test_CovP1_50_ResizeForItems_LargeResize", func() {
		col := corestr.New.Collection.Cap(5)
		// Add many items to trigger resize logic
		items := make([]string, 300)
		for i := range items {
			items[i] = "x"
		}
		col.AddStrings(items)
		if col.Length() != 300 {
			t.Fatal("expected 300")
		}
	})
}
