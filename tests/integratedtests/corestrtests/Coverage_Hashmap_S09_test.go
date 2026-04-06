package corestrtests

import (
	"strings"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════
// S09 — Hashmap.go (1,300 lines) — Full coverage
// ══════════════════════════════════════════════════════════════

// ── IsEmpty / HasItems / IsEmptyLock ─────────────────────────

func Test_S09_01_Hashmap_IsEmpty(t *testing.T) {
	safeTest(t, "Test_S09_01_Hashmap_IsEmpty", func() {
		// Arrange
		empty := corestr.Empty.Hashmap()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		if !empty.IsEmpty() {
			t.Fatal("expected empty")
		}
		if hm.IsEmpty() {
			t.Fatal("expected not empty")
		}
	})
}

func Test_S09_02_Hashmap_HasItems(t *testing.T) {
	safeTest(t, "Test_S09_02_Hashmap_HasItems", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		if !hm.HasItems() {
			t.Fatal("expected has items")
		}
		if corestr.Empty.Hashmap().HasItems() {
			t.Fatal("expected no items for empty")
		}
	})
}

func Test_S09_03_Hashmap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_S09_03_Hashmap_IsEmptyLock", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act & Assert
		if !hm.IsEmptyLock() {
			t.Fatal("expected empty")
		}
	})
}

func Test_S09_04_Hashmap_Collection(t *testing.T) {
	safeTest(t, "Test_S09_04_Hashmap_Collection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		col := hm.Collection()

		// Assert
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

// ── AddOrUpdate variants ─────────────────────────────────────

func Test_S09_05_Hashmap_AddOrUpdateWithWgLock(t *testing.T) {
	safeTest(t, "Test_S09_05_Hashmap_AddOrUpdateWithWgLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		hm.AddOrUpdateWithWgLock("k", "v", wg)
		wg.Wait()

		// Assert
		if hm.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_06_Hashmap_AddOrUpdateKeyStrValInt(t *testing.T) {
	safeTest(t, "Test_S09_06_Hashmap_AddOrUpdateKeyStrValInt", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateKeyStrValInt("count", 42)

		// Assert
		v, ok := hm.Get("count")
		if !ok || v != "42" {
			t.Fatalf("expected '42', got '%s'", v)
		}
	})
}

func Test_S09_07_Hashmap_AddOrUpdateKeyStrValFloat(t *testing.T) {
	safeTest(t, "Test_S09_07_Hashmap_AddOrUpdateKeyStrValFloat", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateKeyStrValFloat("pi", 3.14)

		// Assert
		v, ok := hm.Get("pi")
		if !ok || v == "" {
			t.Fatal("expected float value")
		}
	})
}

func Test_S09_08_Hashmap_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	safeTest(t, "Test_S09_08_Hashmap_AddOrUpdateKeyStrValFloat64", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateKeyStrValFloat64("e", 2.71828)

		// Assert
		v, ok := hm.Get("e")
		if !ok || v == "" {
			t.Fatal("expected float64 value")
		}
	})
}

func Test_S09_09_Hashmap_AddOrUpdateKeyStrValAny(t *testing.T) {
	safeTest(t, "Test_S09_09_Hashmap_AddOrUpdateKeyStrValAny", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateKeyStrValAny("key", "anyVal")

		// Assert
		v, ok := hm.Get("key")
		if !ok || v == "" {
			t.Fatal("expected value")
		}
	})
}

func Test_S09_10_Hashmap_AddOrUpdateKeyValueAny(t *testing.T) {
	safeTest(t, "Test_S09_10_Hashmap_AddOrUpdateKeyValueAny", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		pair := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		hm.AddOrUpdateKeyValueAny(pair)

		// Assert
		if hm.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_11_Hashmap_AddOrUpdateKeyVal(t *testing.T) {
	safeTest(t, "Test_S09_11_Hashmap_AddOrUpdateKeyVal", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		isNew := hm.AddOrUpdateKeyVal(kv)

		// Assert
		if !isNew {
			t.Fatal("expected new")
		}
		isNew2 := hm.AddOrUpdateKeyVal(kv)
		if isNew2 {
			t.Fatal("expected not new on second add")
		}
	})
}

func Test_S09_12_Hashmap_AddOrUpdate(t *testing.T) {
	safeTest(t, "Test_S09_12_Hashmap_AddOrUpdate", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		isNew := hm.AddOrUpdate("k", "v")
		isNew2 := hm.AddOrUpdate("k", "v2")

		// Assert
		if !isNew {
			t.Fatal("expected new")
		}
		if isNew2 {
			t.Fatal("expected update not new")
		}
	})
}

func Test_S09_13_Hashmap_Set(t *testing.T) {
	safeTest(t, "Test_S09_13_Hashmap_Set", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		isNew := hm.Set("k", "v")

		// Assert
		if !isNew {
			t.Fatal("expected new")
		}
	})
}

func Test_S09_14_Hashmap_SetTrim(t *testing.T) {
	safeTest(t, "Test_S09_14_Hashmap_SetTrim", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.SetTrim("  key  ", "  val  ")

		// Assert
		v, ok := hm.Get("key")
		if !ok || v != "val" {
			t.Fatalf("expected 'val', got '%s'", v)
		}
	})
}

func Test_S09_15_Hashmap_SetBySplitter(t *testing.T) {
	safeTest(t, "Test_S09_15_Hashmap_SetBySplitter", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.SetBySplitter("=", "key=value")
		hm.SetBySplitter("=", "novalue")

		// Assert
		v, ok := hm.Get("key")
		if !ok || v != "value" {
			t.Fatalf("expected 'value', got '%s'", v)
		}
		v2, ok2 := hm.Get("novalue")
		if !ok2 || v2 != "" {
			t.Fatalf("expected empty value, got '%s'", v2)
		}
	})
}

func Test_S09_16_Hashmap_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_S09_16_Hashmap_AddOrUpdateStringsPtrWgLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		hm.AddOrUpdateStringsPtrWgLock(wg, []string{"a", "b"}, []string{"1", "2"})
		wg.Wait()

		// Assert
		if hm.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S09_17_Hashmap_AddOrUpdateStringsPtrWgLock_Empty(t *testing.T) {
	safeTest(t, "Test_S09_17_Hashmap_AddOrUpdateStringsPtrWgLock_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		wg := &sync.WaitGroup{}

		// Act
		hm.AddOrUpdateStringsPtrWgLock(wg, []string{}, []string{})

		// Assert
		if hm.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S09_18_Hashmap_AddOrUpdateStringsPtrWgLock_Panic(t *testing.T) {
	safeTest(t, "Test_S09_18_Hashmap_AddOrUpdateStringsPtrWgLock_Panic", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		wg := &sync.WaitGroup{}

		// Act & Assert
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("expected panic for mismatched lengths")
			}
		}()
		wg.Add(1)
		hm.AddOrUpdateStringsPtrWgLock(wg, []string{"a"}, []string{"1", "2"})
	})
}

func Test_S09_19_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_S09_19_Hashmap_AddOrUpdateHashmap", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "v1")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("k2", "v2")

		// Act
		hm.AddOrUpdateHashmap(other)

		// Assert
		if hm.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S09_20_Hashmap_AddOrUpdateHashmap_Nil(t *testing.T) {
	safeTest(t, "Test_S09_20_Hashmap_AddOrUpdateHashmap_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateHashmap(nil)

		// Assert
		if hm.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S09_21_Hashmap_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_S09_21_Hashmap_AddOrUpdateMap", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateMap(map[string]string{"a": "1", "b": "2"})

		// Assert
		if hm.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S09_22_Hashmap_AddOrUpdateMap_Empty(t *testing.T) {
	safeTest(t, "Test_S09_22_Hashmap_AddOrUpdateMap_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateMap(map[string]string{})

		// Assert
		if hm.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S09_23_Hashmap_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_S09_23_Hashmap_AddsOrUpdates", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddsOrUpdates(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		// Assert
		if hm.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S09_24_Hashmap_AddsOrUpdates_Nil(t *testing.T) {
	safeTest(t, "Test_S09_24_Hashmap_AddsOrUpdates_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddsOrUpdates(nil...)

		// Assert
		if hm.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S09_25_Hashmap_AddOrUpdateKeyAnyValues(t *testing.T) {
	safeTest(t, "Test_S09_25_Hashmap_AddOrUpdateKeyAnyValues", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateKeyAnyValues(
			corestr.KeyAnyValuePair{Key: "k", Value: 42},
		)

		// Assert
		if hm.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_26_Hashmap_AddOrUpdateKeyAnyValues_Empty(t *testing.T) {
	safeTest(t, "Test_S09_26_Hashmap_AddOrUpdateKeyAnyValues_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateKeyAnyValues()

		// Assert
		if hm.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S09_27_Hashmap_AddOrUpdateKeyValues(t *testing.T) {
	safeTest(t, "Test_S09_27_Hashmap_AddOrUpdateKeyValues", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateKeyValues(
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)

		// Assert
		if hm.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_28_Hashmap_AddOrUpdateKeyValues_Empty(t *testing.T) {
	safeTest(t, "Test_S09_28_Hashmap_AddOrUpdateKeyValues_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateKeyValues()

		// Assert
		if hm.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

// ── AddOrUpdateCollection ────────────────────────────────────

func Test_S09_29_Hashmap_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_S09_29_Hashmap_AddOrUpdateCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		keys := corestr.New.Collection.Strings([]string{"a", "b"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})

		// Act
		hm.AddOrUpdateCollection(keys, vals)

		// Assert
		if hm.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S09_30_Hashmap_AddOrUpdateCollection_NilKeys(t *testing.T) {
	safeTest(t, "Test_S09_30_Hashmap_AddOrUpdateCollection_NilKeys", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		vals := corestr.New.Collection.Strings([]string{"1"})

		// Act
		hm.AddOrUpdateCollection(nil, vals)

		// Assert
		if hm.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S09_31_Hashmap_AddOrUpdateCollection_LengthMismatch(t *testing.T) {
	safeTest(t, "Test_S09_31_Hashmap_AddOrUpdateCollection_LengthMismatch", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		keys := corestr.New.Collection.Strings([]string{"a"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})

		// Act
		hm.AddOrUpdateCollection(keys, vals)

		// Assert
		if hm.Length() != 0 {
			t.Fatal("expected 0 for length mismatch")
		}
	})
}

// ── Filter-based adds ────────────────────────────────────────

func Test_S09_32_Hashmap_AddsOrUpdatesAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_S09_32_Hashmap_AddsOrUpdatesAnyUsingFilter", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return pair.ValueString(), true, false
		}

		// Act
		hm.AddsOrUpdatesAnyUsingFilter(filter,
			corestr.KeyAnyValuePair{Key: "k", Value: "v"},
		)

		// Assert
		if hm.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_33_Hashmap_AddsOrUpdatesAnyUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_S09_33_Hashmap_AddsOrUpdatesAnyUsingFilter_Break", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return pair.ValueString(), true, true
		}

		// Act
		hm.AddsOrUpdatesAnyUsingFilter(filter,
			corestr.KeyAnyValuePair{Key: "a", Value: "1"},
			corestr.KeyAnyValuePair{Key: "b", Value: "2"},
		)

		// Assert
		if hm.Length() != 1 {
			t.Fatal("expected 1 due to break")
		}
	})
}

func Test_S09_34_Hashmap_AddsOrUpdatesAnyUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_S09_34_Hashmap_AddsOrUpdatesAnyUsingFilter_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddsOrUpdatesAnyUsingFilter(nil, nil...)

		// Assert
		if hm.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S09_35_Hashmap_AddsOrUpdatesAnyUsingFilter_Skip(t *testing.T) {
	safeTest(t, "Test_S09_35_Hashmap_AddsOrUpdatesAnyUsingFilter_Skip", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return "", false, false
		}

		// Act
		hm.AddsOrUpdatesAnyUsingFilter(filter,
			corestr.KeyAnyValuePair{Key: "k", Value: "v"},
		)

		// Assert
		if hm.Length() != 0 {
			t.Fatal("expected 0 — all skipped")
		}
	})
}

func Test_S09_36_Hashmap_AddsOrUpdatesAnyUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_S09_36_Hashmap_AddsOrUpdatesAnyUsingFilterLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return pair.ValueString(), true, false
		}

		// Act
		hm.AddsOrUpdatesAnyUsingFilterLock(filter,
			corestr.KeyAnyValuePair{Key: "k", Value: "v"},
		)

		// Assert
		if hm.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_37_Hashmap_AddsOrUpdatesAnyUsingFilterLock_Nil(t *testing.T) {
	safeTest(t, "Test_S09_37_Hashmap_AddsOrUpdatesAnyUsingFilterLock_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddsOrUpdatesAnyUsingFilterLock(nil, nil...)

		// Assert
		if hm.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S09_38_Hashmap_AddsOrUpdatesAnyUsingFilterLock_SkipAndBreak(t *testing.T) {
	safeTest(t, "Test_S09_38_Hashmap_AddsOrUpdatesAnyUsingFilterLock_SkipAndBreak", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		callCount := 0
		filter := func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			callCount++
			if callCount == 1 {
				return "", false, false
			}
			return pair.ValueString(), true, true
		}

		// Act
		hm.AddsOrUpdatesAnyUsingFilterLock(filter,
			corestr.KeyAnyValuePair{Key: "a", Value: "1"},
			corestr.KeyAnyValuePair{Key: "b", Value: "2"},
			corestr.KeyAnyValuePair{Key: "c", Value: "3"},
		)

		// Assert
		if hm.Length() != 1 {
			t.Fatalf("expected 1, got %d", hm.Length())
		}
	})
}

func Test_S09_39_Hashmap_AddsOrUpdatesUsingFilter(t *testing.T) {
	safeTest(t, "Test_S09_39_Hashmap_AddsOrUpdatesUsingFilter", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Value, true, false
		}

		// Act
		hm.AddsOrUpdatesUsingFilter(filter,
			corestr.KeyValuePair{Key: "k", Value: "v"},
		)

		// Assert
		if hm.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_40_Hashmap_AddsOrUpdatesUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_S09_40_Hashmap_AddsOrUpdatesUsingFilter_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddsOrUpdatesUsingFilter(nil, nil...)

		// Assert
		if hm.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S09_41_Hashmap_AddsOrUpdatesUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_S09_41_Hashmap_AddsOrUpdatesUsingFilter_Break", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Value, true, true
		}

		// Act
		hm.AddsOrUpdatesUsingFilter(filter,
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		// Assert
		if hm.Length() != 1 {
			t.Fatal("expected 1 due to break")
		}
	})
}

func Test_S09_42_Hashmap_AddsOrUpdatesUsingFilter_Skip(t *testing.T) {
	safeTest(t, "Test_S09_42_Hashmap_AddsOrUpdatesUsingFilter_Skip", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return "", false, false
		}

		// Act
		hm.AddsOrUpdatesUsingFilter(filter,
			corestr.KeyValuePair{Key: "k", Value: "v"},
		)

		// Assert
		if hm.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

// ── ConcatNew / ConcatNewUsingMaps ───────────────────────────

func Test_S09_43_Hashmap_ConcatNew(t *testing.T) {
	safeTest(t, "Test_S09_43_Hashmap_ConcatNew", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("b", "2")

		// Act
		result := hm.ConcatNew(true, other)

		// Assert
		if result.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_S09_44_Hashmap_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_S09_44_Hashmap_ConcatNew_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		result := hm.ConcatNew(true)

		// Assert
		if result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_45_Hashmap_ConcatNew_NilInList(t *testing.T) {
	safeTest(t, "Test_S09_45_Hashmap_ConcatNew_NilInList", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		result := hm.ConcatNew(false, nil)

		// Assert
		if result.Length() < 1 {
			t.Fatal("expected at least 1")
		}
	})
}

func Test_S09_46_Hashmap_ConcatNewUsingMaps(t *testing.T) {
	safeTest(t, "Test_S09_46_Hashmap_ConcatNewUsingMaps", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		result := hm.ConcatNewUsingMaps(true, map[string]string{"b": "2"})

		// Assert
		if result.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_S09_47_Hashmap_ConcatNewUsingMaps_Empty(t *testing.T) {
	safeTest(t, "Test_S09_47_Hashmap_ConcatNewUsingMaps_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		result := hm.ConcatNewUsingMaps(false)

		// Assert
		if result == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S09_48_Hashmap_ConcatNewUsingMaps_NilInList(t *testing.T) {
	safeTest(t, "Test_S09_48_Hashmap_ConcatNewUsingMaps_NilInList", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		result := hm.ConcatNewUsingMaps(false, nil)

		// Assert
		if result.Length() < 1 {
			t.Fatal("expected at least 1")
		}
	})
}

// ── AddOrUpdateLock ──────────────────────────────────────────

func Test_S09_49_Hashmap_AddOrUpdateLock(t *testing.T) {
	safeTest(t, "Test_S09_49_Hashmap_AddOrUpdateLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddOrUpdateLock("k", "v")

		// Assert
		if hm.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

// ── Has / Contains / ContainsLock / IsKeyMissing / IsKeyMissingLock / HasLock / HasWithLock ──

func Test_S09_50_Hashmap_Has(t *testing.T) {
	safeTest(t, "Test_S09_50_Hashmap_Has", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		if !hm.Has("k") {
			t.Fatal("expected true")
		}
		if hm.Has("missing") {
			t.Fatal("expected false")
		}
	})
}

func Test_S09_51_Hashmap_Contains(t *testing.T) {
	safeTest(t, "Test_S09_51_Hashmap_Contains", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		if !hm.Contains("k") {
			t.Fatal("expected true")
		}
	})
}

func Test_S09_52_Hashmap_ContainsLock(t *testing.T) {
	safeTest(t, "Test_S09_52_Hashmap_ContainsLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		if !hm.ContainsLock("k") {
			t.Fatal("expected true")
		}
	})
}

func Test_S09_53_Hashmap_IsKeyMissing(t *testing.T) {
	safeTest(t, "Test_S09_53_Hashmap_IsKeyMissing", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		if hm.IsKeyMissing("k") {
			t.Fatal("expected false")
		}
		if !hm.IsKeyMissing("missing") {
			t.Fatal("expected true")
		}
	})
}

func Test_S09_54_Hashmap_IsKeyMissingLock(t *testing.T) {
	safeTest(t, "Test_S09_54_Hashmap_IsKeyMissingLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		if hm.IsKeyMissingLock("k") {
			t.Fatal("expected false")
		}
	})
}

func Test_S09_55_Hashmap_HasLock(t *testing.T) {
	safeTest(t, "Test_S09_55_Hashmap_HasLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		if !hm.HasLock("k") {
			t.Fatal("expected true")
		}
	})
}

func Test_S09_56_Hashmap_HasWithLock(t *testing.T) {
	safeTest(t, "Test_S09_56_Hashmap_HasWithLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		if !hm.HasWithLock("k") {
			t.Fatal("expected true")
		}
	})
}

// ── HasAllStrings / HasAll / HasAnyItem / HasAny ─────────────

func Test_S09_57_Hashmap_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_S09_57_Hashmap_HasAllStrings", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")

		// Act & Assert
		if !hm.HasAllStrings("a", "b") {
			t.Fatal("expected true")
		}
		if hm.HasAllStrings("a", "c") {
			t.Fatal("expected false")
		}
	})
}

func Test_S09_58_Hashmap_HasAll(t *testing.T) {
	safeTest(t, "Test_S09_58_Hashmap_HasAll", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("x", "1")

		// Act & Assert
		if !hm.HasAll("x") {
			t.Fatal("expected true")
		}
		if hm.HasAll("x", "y") {
			t.Fatal("expected false")
		}
	})
}

func Test_S09_59_Hashmap_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_S09_59_Hashmap_HasAnyItem", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		if !hm.HasAnyItem() {
			t.Fatal("expected true")
		}
	})
}

func Test_S09_60_Hashmap_HasAny(t *testing.T) {
	safeTest(t, "Test_S09_60_Hashmap_HasAny", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act & Assert
		if !hm.HasAny("a", "z") {
			t.Fatal("expected true")
		}
		if hm.HasAny("x", "y") {
			t.Fatal("expected false")
		}
	})
}

// ── HasAllCollectionItems ────────────────────────────────────

func Test_S09_61_Hashmap_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_S09_61_Hashmap_HasAllCollectionItems", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act & Assert
		if !hm.HasAllCollectionItems(col) {
			t.Fatal("expected true")
		}
	})
}

func Test_S09_62_Hashmap_HasAllCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_S09_62_Hashmap_HasAllCollectionItems_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act & Assert
		if hm.HasAllCollectionItems(nil) {
			t.Fatal("expected false for nil")
		}
	})
}

func Test_S09_63_Hashmap_HasAllCollectionItems_Empty(t *testing.T) {
	safeTest(t, "Test_S09_63_Hashmap_HasAllCollectionItems_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act & Assert
		if hm.HasAllCollectionItems(corestr.Empty.Collection()) {
			t.Fatal("expected false for empty")
		}
	})
}

// ── DiffRaw / Diff ───────────────────────────────────────────

func Test_S09_64_Hashmap_DiffRaw(t *testing.T) {
	safeTest(t, "Test_S09_64_Hashmap_DiffRaw", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		diff := hm.DiffRaw(map[string]string{"a": "2"})

		// Assert
		if diff == nil {
			t.Fatal("expected non-nil diff")
		}
	})
}

func Test_S09_65_Hashmap_Diff(t *testing.T) {
	safeTest(t, "Test_S09_65_Hashmap_Diff", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("a", "2")

		// Act
		diff := hm.Diff(other)

		// Assert
		if diff == nil {
			t.Fatal("expected non-nil")
		}
	})
}

// ── GetKeysFilteredItems / GetKeysFilteredCollection ─────────

func Test_S09_66_Hashmap_GetKeysFilteredItems(t *testing.T) {
	safeTest(t, "Test_S09_66_Hashmap_GetKeysFilteredItems", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("apple", "1")
		hm.AddOrUpdate("banana", "2")
		filter := func(str string, index int) (string, bool, bool) {
			return str, strings.HasPrefix(str, "a"), false
		}

		// Act
		result := hm.GetKeysFilteredItems(filter)

		// Assert
		if len(result) != 1 {
			t.Fatalf("expected 1, got %d", len(result))
		}
	})
}

func Test_S09_67_Hashmap_GetKeysFilteredItems_Empty(t *testing.T) {
	safeTest(t, "Test_S09_67_Hashmap_GetKeysFilteredItems_Empty", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		result := hm.GetKeysFilteredItems(filter)

		// Assert
		if len(result) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S09_68_Hashmap_GetKeysFilteredItems_Break(t *testing.T) {
	safeTest(t, "Test_S09_68_Hashmap_GetKeysFilteredItems_Break", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, true
		}

		// Act
		result := hm.GetKeysFilteredItems(filter)

		// Assert
		if len(result) != 1 {
			t.Fatalf("expected 1 due to break, got %d", len(result))
		}
	})
}

func Test_S09_69_Hashmap_GetKeysFilteredCollection(t *testing.T) {
	safeTest(t, "Test_S09_69_Hashmap_GetKeysFilteredCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("x", "1")
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		result := hm.GetKeysFilteredCollection(filter)

		// Assert
		if result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_70_Hashmap_GetKeysFilteredCollection_Empty(t *testing.T) {
	safeTest(t, "Test_S09_70_Hashmap_GetKeysFilteredCollection_Empty", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		result := hm.GetKeysFilteredCollection(nil)

		// Assert
		if !result.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_S09_71_Hashmap_GetKeysFilteredCollection_Break(t *testing.T) {
	safeTest(t, "Test_S09_71_Hashmap_GetKeysFilteredCollection_Break", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, true
		}

		// Act
		result := hm.GetKeysFilteredCollection(filter)

		// Assert
		if result.Length() != 1 {
			t.Fatalf("expected 1, got %d", result.Length())
		}
	})
}

// ── Items / SafeItems / ItemsCopyLock ────────────────────────

func Test_S09_72_Hashmap_Items(t *testing.T) {
	safeTest(t, "Test_S09_72_Hashmap_Items", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		if len(hm.Items()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_73_Hashmap_SafeItems(t *testing.T) {
	safeTest(t, "Test_S09_73_Hashmap_SafeItems", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act & Assert
		if hm.SafeItems() != nil {
			t.Fatal("expected nil for nil hashmap")
		}
	})
}

func Test_S09_74_Hashmap_ItemsCopyLock(t *testing.T) {
	safeTest(t, "Test_S09_74_Hashmap_ItemsCopyLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		copied := hm.ItemsCopyLock()

		// Assert
		if len(*copied) != 1 {
			t.Fatal("expected 1")
		}
	})
}

// ── ValuesCollection / ValuesHashset / ValuesCollectionLock / ValuesHashsetLock ──

func Test_S09_75_Hashmap_ValuesCollection(t *testing.T) {
	safeTest(t, "Test_S09_75_Hashmap_ValuesCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		col := hm.ValuesCollection()

		// Assert
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_76_Hashmap_ValuesHashset(t *testing.T) {
	safeTest(t, "Test_S09_76_Hashmap_ValuesHashset", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		hs := hm.ValuesHashset()

		// Assert
		if hs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_77_Hashmap_ValuesCollectionLock(t *testing.T) {
	safeTest(t, "Test_S09_77_Hashmap_ValuesCollectionLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		col := hm.ValuesCollectionLock()

		// Assert
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_78_Hashmap_ValuesHashsetLock(t *testing.T) {
	safeTest(t, "Test_S09_78_Hashmap_ValuesHashsetLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		hs := hm.ValuesHashsetLock()

		// Assert
		if hs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

// ── ValuesList / KeysValuesCollection / KeysValuesList ───────

func Test_S09_79_Hashmap_ValuesList(t *testing.T) {
	safeTest(t, "Test_S09_79_Hashmap_ValuesList", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		list := hm.ValuesList()

		// Assert
		if len(list) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_80_Hashmap_KeysValuesCollection(t *testing.T) {
	safeTest(t, "Test_S09_80_Hashmap_KeysValuesCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		keys, values := hm.KeysValuesCollection()

		// Assert
		if keys.Length() != 1 || values.Length() != 1 {
			t.Fatal("expected 1 each")
		}
	})
}

func Test_S09_81_Hashmap_KeysValuesList(t *testing.T) {
	safeTest(t, "Test_S09_81_Hashmap_KeysValuesList", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		keys, values := hm.KeysValuesList()

		// Assert
		if len(keys) != 1 || len(values) != 1 {
			t.Fatal("expected 1 each")
		}
	})
}

func Test_S09_82_Hashmap_KeysValuesListLock(t *testing.T) {
	safeTest(t, "Test_S09_82_Hashmap_KeysValuesListLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		keys, values := hm.KeysValuesListLock()

		// Assert
		if len(keys) != 1 || len(values) != 1 {
			t.Fatal("expected 1 each")
		}
	})
}

// ── KeysValuePairs / KeysValuePairsCollection ────────────────

func Test_S09_83_Hashmap_KeysValuePairs(t *testing.T) {
	safeTest(t, "Test_S09_83_Hashmap_KeysValuePairs", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		pairs := hm.KeysValuePairs()

		// Assert
		if len(pairs) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_84_Hashmap_KeysValuePairsCollection(t *testing.T) {
	safeTest(t, "Test_S09_84_Hashmap_KeysValuePairsCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		kvc := hm.KeysValuePairsCollection()

		// Assert
		if kvc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

// ── AllKeys / Keys / KeysCollection / KeysLock ───────────────

func Test_S09_85_Hashmap_AllKeys(t *testing.T) {
	safeTest(t, "Test_S09_85_Hashmap_AllKeys", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		keys := hm.AllKeys()

		// Assert
		if len(keys) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_86_Hashmap_AllKeys_Empty(t *testing.T) {
	safeTest(t, "Test_S09_86_Hashmap_AllKeys_Empty", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		keys := hm.AllKeys()

		// Assert
		if len(keys) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S09_87_Hashmap_Keys(t *testing.T) {
	safeTest(t, "Test_S09_87_Hashmap_Keys", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act & Assert
		if len(hm.Keys()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_88_Hashmap_KeysCollection(t *testing.T) {
	safeTest(t, "Test_S09_88_Hashmap_KeysCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act & Assert
		if hm.KeysCollection().Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_89_Hashmap_KeysLock(t *testing.T) {
	safeTest(t, "Test_S09_89_Hashmap_KeysLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		keys := hm.KeysLock()

		// Assert
		if len(keys) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_90_Hashmap_KeysLock_Empty(t *testing.T) {
	safeTest(t, "Test_S09_90_Hashmap_KeysLock_Empty", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		keys := hm.KeysLock()

		// Assert
		if len(keys) != 0 {
			t.Fatal("expected 0")
		}
	})
}
func Test_S09_92_Hashmap_KeysToLower(t *testing.T) {
	safeTest(t, "Test_S09_92_Hashmap_KeysToLower", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("KEY", "v")

		// Act
		lowered := hm.KeysToLower()

		// Assert
		if !lowered.Has("key") {
			t.Fatal("expected lowercased key")
		}
	})
}
func Test_S09_94_Hashmap_Length(t *testing.T) {
	safeTest(t, "Test_S09_94_Hashmap_Length", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		if hm.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_95_Hashmap_Length_Nil(t *testing.T) {
	safeTest(t, "Test_S09_95_Hashmap_Length_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act & Assert
		if hm.Length() != 0 {
			t.Fatal("expected 0 for nil")
		}
	})
}

func Test_S09_96_Hashmap_LengthLock(t *testing.T) {
	safeTest(t, "Test_S09_96_Hashmap_LengthLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		if hm.LengthLock() != 1 {
			t.Fatal("expected 1")
		}
	})
}

// ── IsEqual / IsEqualPtr / IsEqualPtrLock ────────────────────

func Test_S09_97_Hashmap_IsEqual_Same(t *testing.T) {
	safeTest(t, "Test_S09_97_Hashmap_IsEqual_Same", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("k", "v")

		// Act & Assert
		if !hm.IsEqualPtr(other) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S09_98_Hashmap_IsEqualPtr_DiffValues(t *testing.T) {
	safeTest(t, "Test_S09_98_Hashmap_IsEqualPtr_DiffValues", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v1")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("k", "v2")

		// Act & Assert
		if hm.IsEqualPtr(other) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S09_99_Hashmap_IsEqualPtr_DiffLength(t *testing.T) {
	safeTest(t, "Test_S09_99_Hashmap_IsEqualPtr_DiffLength", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("a", "1")
		other.AddOrUpdate("b", "2")

		// Act & Assert
		if hm.IsEqualPtr(other) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S09_100_Hashmap_IsEqualPtr_BothNil(t *testing.T) {
	safeTest(t, "Test_S09_100_Hashmap_IsEqualPtr_BothNil", func() {
		// Arrange
		var a *corestr.Hashmap
		var b *corestr.Hashmap

		// Act & Assert
		if !a.IsEqualPtr(b) {
			t.Fatal("expected equal for both nil")
		}
	})
}

func Test_S09_101_Hashmap_IsEqualPtr_OneNil(t *testing.T) {
	safeTest(t, "Test_S09_101_Hashmap_IsEqualPtr_OneNil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		var other *corestr.Hashmap

		// Act & Assert
		if hm.IsEqualPtr(other) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S09_102_Hashmap_IsEqualPtr_SamePtr(t *testing.T) {
	safeTest(t, "Test_S09_102_Hashmap_IsEqualPtr_SamePtr", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		if !hm.IsEqualPtr(hm) {
			t.Fatal("expected equal for same pointer")
		}
	})
}

func Test_S09_103_Hashmap_IsEqualPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S09_103_Hashmap_IsEqualPtr_BothEmpty", func() {
		// Arrange
		a := corestr.Empty.Hashmap()
		b := corestr.Empty.Hashmap()

		// Act & Assert
		if !a.IsEqualPtr(b) {
			t.Fatal("expected equal for both empty")
		}
	})
}

func Test_S09_104_Hashmap_IsEqualPtr_OneEmpty(t *testing.T) {
	safeTest(t, "Test_S09_104_Hashmap_IsEqualPtr_OneEmpty", func() {
		// Arrange
		a := corestr.New.Hashmap.Cap(5)
		a.AddOrUpdate("k", "v")
		b := corestr.Empty.Hashmap()

		// Act & Assert
		if a.IsEqualPtr(b) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S09_105_Hashmap_IsEqualPtr_MissingKey(t *testing.T) {
	safeTest(t, "Test_S09_105_Hashmap_IsEqualPtr_MissingKey", func() {
		// Arrange
		a := corestr.New.Hashmap.Cap(5)
		a.AddOrUpdate("a", "1")
		b := corestr.New.Hashmap.Cap(5)
		b.AddOrUpdate("b", "1")

		// Act & Assert
		if a.IsEqualPtr(b) {
			t.Fatal("expected not equal — different keys")
		}
	})
}

func Test_S09_106_Hashmap_IsEqual(t *testing.T) {
	safeTest(t, "Test_S09_106_Hashmap_IsEqual", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("k", "v")

		// Act & Assert — value receiver
		if !hm.IsEqual(*other) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S09_107_Hashmap_IsEqualPtrLock(t *testing.T) {
	safeTest(t, "Test_S09_107_Hashmap_IsEqualPtrLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		other := corestr.New.Hashmap.Cap(5)
		other.AddOrUpdate("k", "v")

		// Act & Assert
		if !hm.IsEqualPtrLock(other) {
			t.Fatal("expected equal")
		}
	})
}

// ── Remove / RemoveWithLock ──────────────────────────────────

func Test_S09_108_Hashmap_Remove(t *testing.T) {
	safeTest(t, "Test_S09_108_Hashmap_Remove", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		hm.Remove("k")

		// Assert
		if hm.Has("k") {
			t.Fatal("expected removed")
		}
	})
}

func Test_S09_109_Hashmap_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_S09_109_Hashmap_RemoveWithLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		hm.RemoveWithLock("k")

		// Assert
		if hm.Has("k") {
			t.Fatal("expected removed")
		}
	})
}

// ── String / StringLock ──────────────────────────────────────

func Test_S09_110_Hashmap_String(t *testing.T) {
	safeTest(t, "Test_S09_110_Hashmap_String", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		s := hm.String()

		// Assert
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S09_111_Hashmap_String_Empty(t *testing.T) {
	safeTest(t, "Test_S09_111_Hashmap_String_Empty", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		s := hm.String()

		// Assert
		if !strings.Contains(s, "No Element") {
			t.Fatal("expected No Element")
		}
	})
}

func Test_S09_112_Hashmap_StringLock(t *testing.T) {
	safeTest(t, "Test_S09_112_Hashmap_StringLock", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		s := hm.StringLock()

		// Assert
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S09_113_Hashmap_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_S09_113_Hashmap_StringLock_Empty", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		s := hm.StringLock()

		// Assert
		if !strings.Contains(s, "No Element") {
			t.Fatal("expected No Element")
		}
	})
}

// ── GetValuesExceptKeysInHashset / GetValuesKeysExcept / GetAllExceptCollection ──

func Test_S09_114_Hashmap_GetValuesExceptKeysInHashset(t *testing.T) {
	safeTest(t, "Test_S09_114_Hashmap_GetValuesExceptKeysInHashset", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hm.GetValuesExceptKeysInHashset(hs)

		// Assert
		if len(result) != 1 {
			t.Fatalf("expected 1, got %d", len(result))
		}
	})
}

func Test_S09_115_Hashmap_GetValuesExceptKeysInHashset_NilHashset(t *testing.T) {
	safeTest(t, "Test_S09_115_Hashmap_GetValuesExceptKeysInHashset_NilHashset", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		result := hm.GetValuesExceptKeysInHashset(nil)

		// Assert
		if len(result) != 1 {
			t.Fatal("expected all values")
		}
	})
}

func Test_S09_116_Hashmap_GetValuesExceptKeysInHashset_EmptyHashset(t *testing.T) {
	safeTest(t, "Test_S09_116_Hashmap_GetValuesExceptKeysInHashset_EmptyHashset", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		result := hm.GetValuesExceptKeysInHashset(corestr.Empty.Hashset())

		// Assert
		if len(result) != 1 {
			t.Fatal("expected all values")
		}
	})
}

func Test_S09_117_Hashmap_GetValuesKeysExcept(t *testing.T) {
	safeTest(t, "Test_S09_117_Hashmap_GetValuesKeysExcept", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")

		// Act
		result := hm.GetValuesKeysExcept([]string{"a"})

		// Assert
		if len(result) != 1 {
			t.Fatalf("expected 1, got %d", len(result))
		}
	})
}

func Test_S09_118_Hashmap_GetValuesKeysExcept_Nil(t *testing.T) {
	safeTest(t, "Test_S09_118_Hashmap_GetValuesKeysExcept_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		result := hm.GetValuesKeysExcept(nil)

		// Assert
		if len(result) != 1 {
			t.Fatal("expected all values")
		}
	})
}

func Test_S09_119_Hashmap_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_S09_119_Hashmap_GetAllExceptCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		result := hm.GetAllExceptCollection(col)

		// Assert
		if len(result) != 1 {
			t.Fatalf("expected 1, got %d", len(result))
		}
	})
}

func Test_S09_120_Hashmap_GetAllExceptCollection_Nil(t *testing.T) {
	safeTest(t, "Test_S09_120_Hashmap_GetAllExceptCollection_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		result := hm.GetAllExceptCollection(nil)

		// Assert
		if len(result) != 1 {
			t.Fatal("expected all values")
		}
	})
}

// ── Join / JoinKeys ──────────────────────────────────────────

func Test_S09_121_Hashmap_Join(t *testing.T) {
	safeTest(t, "Test_S09_121_Hashmap_Join", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		s := hm.Join(",")

		// Assert
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S09_122_Hashmap_JoinKeys(t *testing.T) {
	safeTest(t, "Test_S09_122_Hashmap_JoinKeys", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		s := hm.JoinKeys(",")

		// Assert
		if s != "k" {
			t.Fatalf("expected 'k', got '%s'", s)
		}
	})
}

// ── JSON methods ─────────────────────────────────────────────

func Test_S09_123_Hashmap_JsonModel(t *testing.T) {
	safeTest(t, "Test_S09_123_Hashmap_JsonModel", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act & Assert
		if len(hm.JsonModel()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_124_Hashmap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_S09_124_Hashmap_JsonModelAny", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act & Assert
		if hm.JsonModelAny() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S09_125_Hashmap_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_S09_125_Hashmap_MarshalJSON", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		data, err := hm.MarshalJSON()

		// Assert
		if err != nil || len(data) == 0 {
			t.Fatal("expected valid JSON bytes")
		}
	})
}

func Test_S09_126_Hashmap_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_S09_126_Hashmap_UnmarshalJSON", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		err := hm.UnmarshalJSON([]byte(`{"k":"v"}`))

		// Assert
		if err != nil {
			t.Fatal("expected no error")
		}
		if hm.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_127_Hashmap_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_S09_127_Hashmap_UnmarshalJSON_Invalid", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		err := hm.UnmarshalJSON([]byte(`invalid`))

		// Assert
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_S09_128_Hashmap_Json(t *testing.T) {
	safeTest(t, "Test_S09_128_Hashmap_Json", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		result := hm.Json()

		// Assert
		if result.HasError() {
			t.Fatal("expected no error")
		}
	})
}

func Test_S09_129_Hashmap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_S09_129_Hashmap_JsonPtr", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)

		// Act & Assert
		if hm.JsonPtr() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S09_130_Hashmap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_S09_130_Hashmap_ParseInjectUsingJson", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		jsonResult := hm.JsonPtr()
		target := corestr.Empty.Hashmap()

		// Act
		result, err := target.ParseInjectUsingJson(jsonResult)

		// Assert
		if err != nil || result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_131_Hashmap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_S09_131_Hashmap_ParseInjectUsingJsonMust", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		jsonResult := hm.JsonPtr()
		target := corestr.Empty.Hashmap()

		// Act
		result := target.ParseInjectUsingJsonMust(jsonResult)

		// Assert
		if result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_132_Hashmap_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_S09_132_Hashmap_JsonParseSelfInject", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		jsonResult := hm.JsonPtr()
		target := corestr.Empty.Hashmap()

		// Act
		err := target.JsonParseSelfInject(jsonResult)

		// Assert
		if err != nil {
			t.Fatal("expected no error")
		}
	})
}

// ── ToError / ToDefaultError / KeyValStringLines ─────────────

func Test_S09_133_Hashmap_ToError(t *testing.T) {
	safeTest(t, "Test_S09_133_Hashmap_ToError", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		err := hm.ToError(",")

		// Assert
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_S09_134_Hashmap_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_S09_134_Hashmap_ToDefaultError", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		err := hm.ToDefaultError()

		// Assert
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_S09_135_Hashmap_KeyValStringLines(t *testing.T) {
	safeTest(t, "Test_S09_135_Hashmap_KeyValStringLines", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		lines := hm.KeyValStringLines()

		// Assert
		if len(lines) != 1 {
			t.Fatal("expected 1")
		}
	})
}

// ── Clear / Dispose ──────────────────────────────────────────

func Test_S09_136_Hashmap_Clear(t *testing.T) {
	safeTest(t, "Test_S09_136_Hashmap_Clear", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		_ = hm.ValuesList() // populate cache

		// Act
		hm.Clear()

		// Assert
		if hm.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S09_137_Hashmap_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_S09_137_Hashmap_Clear_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act
		result := hm.Clear()

		// Assert
		if result != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_S09_138_Hashmap_Dispose(t *testing.T) {
	safeTest(t, "Test_S09_138_Hashmap_Dispose", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		hm.Dispose()

		// Assert
		if hm.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S09_139_Hashmap_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_S09_139_Hashmap_Dispose_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act — should not panic
		hm.Dispose()
	})
}

// ── ToStringsUsingCompiler ───────────────────────────────────

func Test_S09_140_Hashmap_ToStringsUsingCompiler(t *testing.T) {
	safeTest(t, "Test_S09_140_Hashmap_ToStringsUsingCompiler", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		lines := hm.ToStringsUsingCompiler(func(key, val string) string {
			return key + "=" + val
		})

		// Assert
		if len(lines) != 1 || lines[0] != "k=v" {
			t.Fatal("expected 'k=v'")
		}
	})
}

func Test_S09_141_Hashmap_ToStringsUsingCompiler_Empty(t *testing.T) {
	safeTest(t, "Test_S09_141_Hashmap_ToStringsUsingCompiler_Empty", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		lines := hm.ToStringsUsingCompiler(func(key, val string) string {
			return key
		})

		// Assert
		if len(lines) != 0 {
			t.Fatal("expected 0")
		}
	})
}

// ── AsJsoner / AsJsonContractsBinder / AsJsonParseSelfInjector / AsJsonMarshaller ──

func Test_S09_142_Hashmap_AsJsoner(t *testing.T) {
	safeTest(t, "Test_S09_142_Hashmap_AsJsoner", func() {
		hm := corestr.New.Hashmap.Cap(5)
		if hm.AsJsoner() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S09_143_Hashmap_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_S09_143_Hashmap_AsJsonContractsBinder", func() {
		hm := corestr.New.Hashmap.Cap(5)
		if hm.AsJsonContractsBinder() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S09_144_Hashmap_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_S09_144_Hashmap_AsJsonParseSelfInjector", func() {
		hm := corestr.New.Hashmap.Cap(5)
		if hm.AsJsonParseSelfInjector() == nil {
			t.Fatal("expected non-nil")
		}
	})
}
		hm.AddOrUpdate("k", "v")

		// Act
		cloned := hm.Clone()

		// Assert
		if cloned.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S09_149_Hashmap_Clone_Empty(t *testing.T) {
	safeTest(t, "Test_S09_149_Hashmap_Clone_Empty", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		cloned := hm.Clone()

		// Assert
		if cloned.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

// ── Get / GetValue ───────────────────────────────────────────

func Test_S09_150_Hashmap_Get(t *testing.T) {
	safeTest(t, "Test_S09_150_Hashmap_Get", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		v, ok := hm.Get("k")

		// Assert
		if !ok || v != "v" {
			t.Fatal("expected 'v'")
		}
	})
}

func Test_S09_151_Hashmap_Get_Missing(t *testing.T) {
	safeTest(t, "Test_S09_151_Hashmap_Get_Missing", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		_, ok := hm.Get("missing")

		// Assert
		if ok {
			t.Fatal("expected not found")
		}
	})
}

func Test_S09_152_Hashmap_GetValue(t *testing.T) {
	safeTest(t, "Test_S09_152_Hashmap_GetValue", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		v, ok := hm.GetValue("k")

		// Assert
		if !ok || v != "v" {
			t.Fatal("expected 'v'")
		}
	})
}

// ── Serialize / Deserialize ──────────────────────────────────

func Test_S09_153_Hashmap_Serialize(t *testing.T) {
	safeTest(t, "Test_S09_153_Hashmap_Serialize", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")

		// Act
		data, err := hm.Serialize()

		// Assert
		if err != nil || len(data) == 0 {
			t.Fatal("expected valid bytes")
		}
	})
}

func Test_S09_154_Hashmap_Deserialize(t *testing.T) {
	safeTest(t, "Test_S09_154_Hashmap_Deserialize", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		var target map[string]string

		// Act
		err := hm.Deserialize(&target)

		// Assert
		if err != nil || len(target) != 1 {
			t.Fatal("expected 1")
		}
	})
}
