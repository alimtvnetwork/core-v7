package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Segment 10: Add/Update variants, Has, Filter (L1-700)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovHM1_01_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_CovHM1_01_IsEmpty_HasItems", func() {
		hm := corestr.Empty.Hashmap()
		if !hm.IsEmpty() {
			t.Fatal("expected empty")
		}
		if hm.HasItems() {
			t.Fatal("expected no items")
		}
		hm.AddOrUpdate("a", "1")
		if hm.IsEmpty() {
			t.Fatal("expected not empty")
		}
		if !hm.HasItems() {
			t.Fatal("expected items")
		}
	})
}

func Test_CovHM1_02_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_CovHM1_02_IsEmptyLock", func() {
		hm := corestr.Empty.Hashmap()
		if !hm.IsEmptyLock() {
			t.Fatal("expected empty")
		}
	})
}

func Test_CovHM1_03_Collection(t *testing.T) {
	safeTest(t, "Test_CovHM1_03_Collection", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		col := hm.Collection()
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHM1_04_AddOrUpdateWithWgLock(t *testing.T) {
	safeTest(t, "Test_CovHM1_04_AddOrUpdateWithWgLock", func() {
		hm := corestr.Empty.Hashmap()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hm.AddOrUpdateWithWgLock("a", "1", wg)
		wg.Wait()
		if !hm.Has("a") {
			t.Fatal("expected a")
		}
	})
}

func Test_CovHM1_05_AddOrUpdateKeyStrValInt(t *testing.T) {
	safeTest(t, "Test_CovHM1_05_AddOrUpdateKeyStrValInt", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateKeyStrValInt("age", 25)
		v, _ := hm.Get("age")
		if v != "25" {
			t.Fatalf("expected '25', got '%s'", v)
		}
	})
}

func Test_CovHM1_06_AddOrUpdateKeyStrValFloat(t *testing.T) {
	safeTest(t, "Test_CovHM1_06_AddOrUpdateKeyStrValFloat", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateKeyStrValFloat("f", 1.5)
		_, ok := hm.Get("f")
		if !ok {
			t.Fatal("expected found")
		}
	})
}

func Test_CovHM1_07_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	safeTest(t, "Test_CovHM1_07_AddOrUpdateKeyStrValFloat64", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateKeyStrValFloat64("f", 2.5)
		_, ok := hm.Get("f")
		if !ok {
			t.Fatal("expected found")
		}
	})
}

func Test_CovHM1_08_AddOrUpdateKeyStrValAny(t *testing.T) {
	safeTest(t, "Test_CovHM1_08_AddOrUpdateKeyStrValAny", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateKeyStrValAny("k", 42)
		if !hm.Has("k") {
			t.Fatal("expected k")
		}
	})
}

func Test_CovHM1_09_AddOrUpdateKeyValueAny(t *testing.T) {
	safeTest(t, "Test_CovHM1_09_AddOrUpdateKeyValueAny", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "k", Value: 42})
		if !hm.Has("k") {
			t.Fatal("expected k")
		}
	})
}

func Test_CovHM1_10_AddOrUpdateKeyVal(t *testing.T) {
	safeTest(t, "Test_CovHM1_10_AddOrUpdateKeyVal", func() {
		hm := corestr.Empty.Hashmap()
		isNew := hm.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v"})
		if !isNew {
			t.Fatal("expected new")
		}
		isNew2 := hm.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v2"})
		if isNew2 {
			t.Fatal("expected not new")
		}
	})
}

func Test_CovHM1_11_AddOrUpdate_Set(t *testing.T) {
	safeTest(t, "Test_CovHM1_11_AddOrUpdate_Set", func() {
		hm := corestr.Empty.Hashmap()
		isNew := hm.AddOrUpdate("a", "1")
		if !isNew {
			t.Fatal("expected new")
		}
		isNew2 := hm.Set("a", "2")
		if isNew2 {
			t.Fatal("expected not new")
		}
	})
}

func Test_CovHM1_12_SetTrim(t *testing.T) {
	safeTest(t, "Test_CovHM1_12_SetTrim", func() {
		hm := corestr.Empty.Hashmap()
		hm.SetTrim(" key ", " val ")
		if !hm.Has("key") {
			t.Fatal("expected trimmed key")
		}
	})
}

func Test_CovHM1_13_SetBySplitter(t *testing.T) {
	safeTest(t, "Test_CovHM1_13_SetBySplitter", func() {
		hm := corestr.Empty.Hashmap()
		hm.SetBySplitter("=", "key=value")
		v, _ := hm.Get("key")
		if v != "value" {
			t.Fatalf("expected 'value', got '%s'", v)
		}
		// no value
		hm.SetBySplitter("=", "onlykey")
		v2, _ := hm.Get("onlykey")
		if v2 != "" {
			t.Fatal("expected empty value")
		}
	})
}

func Test_CovHM1_14_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_CovHM1_14_AddOrUpdateStringsPtrWgLock", func() {
		hm := corestr.Empty.Hashmap()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hm.AddOrUpdateStringsPtrWgLock(wg, []string{"a", "b"}, []string{"1", "2"})
		wg.Wait()
		if hm.Length() != 2 {
			t.Fatal("expected 2")
		}
		// empty
		wg.Add(1)
		hm.AddOrUpdateStringsPtrWgLock(wg, []string{}, []string{})
		wg.Wait()
	})
}

func Test_CovHM1_15_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_CovHM1_15_AddOrUpdateHashmap", func() {
		hm := corestr.Empty.Hashmap()
		other := corestr.Empty.Hashmap()
		other.AddOrUpdate("x", "1")
		hm.AddOrUpdateHashmap(other)
		if !hm.Has("x") {
			t.Fatal("expected x")
		}
		hm.AddOrUpdateHashmap(nil)
	})
}

func Test_CovHM1_16_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_CovHM1_16_AddOrUpdateMap", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateMap(map[string]string{"a": "1"})
		if !hm.Has("a") {
			t.Fatal("expected a")
		}
		hm.AddOrUpdateMap(nil)
	})
}

func Test_CovHM1_17_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_CovHM1_17_AddsOrUpdates", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddsOrUpdates(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)
		if hm.Length() != 2 {
			t.Fatal("expected 2")
		}
		hm.AddsOrUpdates()
	})
}

func Test_CovHM1_18_AddOrUpdateKeyAnyValues(t *testing.T) {
	safeTest(t, "Test_CovHM1_18_AddOrUpdateKeyAnyValues", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateKeyAnyValues(
			corestr.KeyAnyValuePair{Key: "k", Value: 1},
		)
		if !hm.Has("k") {
			t.Fatal("expected k")
		}
		hm.AddOrUpdateKeyAnyValues()
	})
}

func Test_CovHM1_19_AddOrUpdateKeyValues(t *testing.T) {
	safeTest(t, "Test_CovHM1_19_AddOrUpdateKeyValues", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateKeyValues(
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)
		if !hm.Has("a") {
			t.Fatal("expected a")
		}
		hm.AddOrUpdateKeyValues()
	})
}

func Test_CovHM1_20_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_CovHM1_20_AddOrUpdateCollection", func() {
		hm := corestr.Empty.Hashmap()
		keys := corestr.New.Collection.Strings([]string{"a", "b"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})
		hm.AddOrUpdateCollection(keys, vals)
		if hm.Length() != 2 {
			t.Fatal("expected 2")
		}
		// nil
		hm.AddOrUpdateCollection(nil, nil)
		// length mismatch
		hm.AddOrUpdateCollection(
			corestr.New.Collection.Strings([]string{"a"}),
			corestr.New.Collection.Strings([]string{"1", "2"}),
		)
	})
}

func Test_CovHM1_21_AddsOrUpdatesAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_CovHM1_21_AddsOrUpdatesAnyUsingFilter", func() {
		hm := corestr.Empty.Hashmap()
		filter := func(p corestr.KeyAnyValuePair) (string, bool, bool) {
			return p.ValueString(), true, false
		}
		hm.AddsOrUpdatesAnyUsingFilter(filter, corestr.KeyAnyValuePair{Key: "a", Value: 1})
		if !hm.Has("a") {
			t.Fatal("expected a")
		}
		// nil
		hm.AddsOrUpdatesAnyUsingFilter(filter)
		// break
		breakFilter := func(p corestr.KeyAnyValuePair) (string, bool, bool) {
			return "", true, true
		}
		hm.AddsOrUpdatesAnyUsingFilter(breakFilter, corestr.KeyAnyValuePair{Key: "b", Value: 2})
	})
}

func Test_CovHM1_22_AddsOrUpdatesAnyUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_CovHM1_22_AddsOrUpdatesAnyUsingFilterLock", func() {
		hm := corestr.Empty.Hashmap()
		filter := func(p corestr.KeyAnyValuePair) (string, bool, bool) {
			return p.ValueString(), true, false
		}
		hm.AddsOrUpdatesAnyUsingFilterLock(filter, corestr.KeyAnyValuePair{Key: "a", Value: 1})
		if !hm.Has("a") {
			t.Fatal("expected a")
		}
		hm.AddsOrUpdatesAnyUsingFilterLock(filter)
		// break
		breakFilter := func(p corestr.KeyAnyValuePair) (string, bool, bool) {
			return "", true, true
		}
		hm.AddsOrUpdatesAnyUsingFilterLock(breakFilter, corestr.KeyAnyValuePair{Key: "x", Value: 1})
	})
}

func Test_CovHM1_23_AddsOrUpdatesUsingFilter(t *testing.T) {
	safeTest(t, "Test_CovHM1_23_AddsOrUpdatesUsingFilter", func() {
		hm := corestr.Empty.Hashmap()
		filter := func(p corestr.KeyValuePair) (string, bool, bool) {
			return p.Value, true, false
		}
		hm.AddsOrUpdatesUsingFilter(filter, corestr.KeyValuePair{Key: "a", Value: "1"})
		if !hm.Has("a") {
			t.Fatal("expected a")
		}
		hm.AddsOrUpdatesUsingFilter(filter)
		// break
		breakFilter := func(p corestr.KeyValuePair) (string, bool, bool) {
			return "", true, true
		}
		hm.AddsOrUpdatesUsingFilter(breakFilter, corestr.KeyValuePair{Key: "x", Value: "1"})
	})
}

func Test_CovHM1_24_ConcatNew(t *testing.T) {
	safeTest(t, "Test_CovHM1_24_ConcatNew", func() {
		a := corestr.Empty.Hashmap()
		a.AddOrUpdate("a", "1")
		// empty
		r := a.ConcatNew(true)
		if r.Length() != 1 {
			t.Fatal("expected 1")
		}
		// with hashmap
		b := corestr.Empty.Hashmap()
		b.AddOrUpdate("b", "2")
		r2 := a.ConcatNew(false, b, nil)
		if r2.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_CovHM1_25_ConcatNewUsingMaps(t *testing.T) {
	safeTest(t, "Test_CovHM1_25_ConcatNewUsingMaps", func() {
		a := corestr.Empty.Hashmap()
		a.AddOrUpdate("a", "1")
		// empty
		r := a.ConcatNewUsingMaps(true)
		if r.Length() != 1 {
			t.Fatal("expected 1")
		}
		// with map
		r2 := a.ConcatNewUsingMaps(false, map[string]string{"b": "2"}, nil)
		if r2.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_CovHM1_26_AddOrUpdateLock(t *testing.T) {
	safeTest(t, "Test_CovHM1_26_AddOrUpdateLock", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdateLock("a", "1")
		if !hm.Has("a") {
			t.Fatal("expected a")
		}
	})
}

func Test_CovHM1_27_Has_Contains_HasLock_HasWithLock(t *testing.T) {
	safeTest(t, "Test_CovHM1_27_Has_Contains_HasLock_HasWithLock", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		if !hm.Has("a") {
			t.Fatal("expected true")
		}
		if !hm.Contains("a") {
			t.Fatal("expected true")
		}
		if !hm.ContainsLock("a") {
			t.Fatal("expected true")
		}
		if !hm.HasLock("a") {
			t.Fatal("expected true")
		}
		if !hm.HasWithLock("a") {
			t.Fatal("expected true")
		}
	})
}

func Test_CovHM1_28_IsKeyMissing_Lock(t *testing.T) {
	safeTest(t, "Test_CovHM1_28_IsKeyMissing_Lock", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		if hm.IsKeyMissing("a") {
			t.Fatal("expected found")
		}
		if !hm.IsKeyMissing("z") {
			t.Fatal("expected missing")
		}
		if hm.IsKeyMissingLock("a") {
			t.Fatal("expected found")
		}
	})
}

func Test_CovHM1_29_HasAllStrings_HasAll(t *testing.T) {
	safeTest(t, "Test_CovHM1_29_HasAllStrings_HasAll", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		if !hm.HasAllStrings("a", "b") {
			t.Fatal("expected true")
		}
		if hm.HasAllStrings("a", "z") {
			t.Fatal("expected false")
		}
		if !hm.HasAll("a", "b") {
			t.Fatal("expected true")
		}
	})
}

func Test_CovHM1_30_HasAnyItem_HasAny(t *testing.T) {
	safeTest(t, "Test_CovHM1_30_HasAnyItem_HasAny", func() {
		hm := corestr.Empty.Hashmap()
		if hm.HasAnyItem() {
			t.Fatal("expected false")
		}
		hm.AddOrUpdate("a", "1")
		if !hm.HasAnyItem() {
			t.Fatal("expected true")
		}
		if !hm.HasAny("z", "a") {
			t.Fatal("expected true")
		}
		if hm.HasAny("x", "y") {
			t.Fatal("expected false")
		}
	})
}

func Test_CovHM1_31_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_CovHM1_31_HasAllCollectionItems", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		if !hm.HasAllCollectionItems(corestr.New.Collection.Strings([]string{"a"})) {
			t.Fatal("expected true")
		}
		if hm.HasAllCollectionItems(nil) {
			t.Fatal("expected false")
		}
	})
}

func Test_CovHM1_32_DiffRaw_Diff(t *testing.T) {
	safeTest(t, "Test_CovHM1_32_DiffRaw_Diff", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		other := corestr.Empty.Hashmap()
		other.AddOrUpdate("b", "2")
		other.AddOrUpdate("c", "3")
		_ = hm.DiffRaw(other.Items())
		_ = hm.Diff(other)
	})
}

func Test_CovHM1_33_GetKeysFilteredItems(t *testing.T) {
	safeTest(t, "Test_CovHM1_33_GetKeysFilteredItems", func() {
		hm := corestr.Empty.Hashmap()
		// empty
		r := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if len(r) != 0 {
			t.Fatal("expected 0")
		}
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		// keep all
		r2 := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if len(r2) != 2 {
			t.Fatal("expected 2")
		}
		// skip
		r3 := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, false, false
		})
		if len(r3) != 0 {
			t.Fatal("expected 0")
		}
		// break
		r4 := hm.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})
		if len(r4) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHM1_34_GetKeysFilteredCollection(t *testing.T) {
	safeTest(t, "Test_CovHM1_34_GetKeysFilteredCollection", func() {
		hm := corestr.Empty.Hashmap()
		// empty
		col := hm.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
		hm.AddOrUpdate("a", "1")
		// keep
		col2 := hm.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if col2.Length() != 1 {
			t.Fatal("expected 1")
		}
		// break
		hm.AddOrUpdate("b", "2")
		col3 := hm.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})
		if col3.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}
