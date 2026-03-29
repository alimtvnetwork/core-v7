package corestrtests

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════
// S10a — Hashset.go Lines 1-700 — Core, Add, Has methods
// ══════════════════════════════════════════════════════════════

// ── IsEmpty / HasItems / IsEmptyLock ─────────────────────────

func Test_S10_01_Hashset_IsEmpty(t *testing.T) {
	safeTest(t, "Test_S10_01_Hashset_IsEmpty", func() {
		// Arrange
		empty := corestr.Empty.Hashset()
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		if !empty.IsEmpty() {
			t.Fatal("expected empty")
		}
		if hs.IsEmpty() {
			t.Fatal("expected not empty")
		}
	})
}

func Test_S10_02_Hashset_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_S10_02_Hashset_IsEmpty_Nil", func() {
		// Arrange
		var hs *corestr.Hashset

		// Act & Assert
		if !hs.IsEmpty() {
			t.Fatal("expected empty for nil")
		}
	})
}

func Test_S10_03_Hashset_HasItems(t *testing.T) {
	safeTest(t, "Test_S10_03_Hashset_HasItems", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		if !hs.HasItems() {
			t.Fatal("expected has items")
		}
		if corestr.Empty.Hashset().HasItems() {
			t.Fatal("expected no items for empty")
		}
	})
}

func Test_S10_04_Hashset_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_S10_04_Hashset_IsEmptyLock", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act & Assert
		if !hs.IsEmptyLock() {
			t.Fatal("expected empty")
		}
	})
}

// ── AddCapacities / AddCapacitiesLock ────────────────────────

func Test_S10_05_Hashset_AddCapacities(t *testing.T) {
	safeTest(t, "Test_S10_05_Hashset_AddCapacities", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(2)
		hs.Add("a")

		// Act
		result := hs.AddCapacities(10, 5)

		// Assert
		if result.Length() != 1 {
			t.Fatal("expected 1 item preserved")
		}
	})
}

func Test_S10_06_Hashset_AddCapacities_Empty(t *testing.T) {
	safeTest(t, "Test_S10_06_Hashset_AddCapacities_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(2)

		// Act
		result := hs.AddCapacities()

		// Assert
		if result != hs {
			t.Fatal("expected same pointer when no capacities")
		}
	})
}

func Test_S10_07_Hashset_AddCapacitiesLock(t *testing.T) {
	safeTest(t, "Test_S10_07_Hashset_AddCapacitiesLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(2)
		hs.Add("a")

		// Act
		result := hs.AddCapacitiesLock(10)

		// Assert
		if result.Length() != 1 {
			t.Fatal("expected 1 preserved")
		}
	})
}

func Test_S10_08_Hashset_AddCapacitiesLock_Empty(t *testing.T) {
	safeTest(t, "Test_S10_08_Hashset_AddCapacitiesLock_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(2)

		// Act
		result := hs.AddCapacitiesLock()

		// Assert
		if result != hs {
			t.Fatal("expected same pointer")
		}
	})
}

// ── Resize / ResizeLock ──────────────────────────────────────

func Test_S10_09_Hashset_Resize(t *testing.T) {
	safeTest(t, "Test_S10_09_Hashset_Resize", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		result := hs.Resize(10)

		// Assert
		if result.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S10_10_Hashset_Resize_SmallerThanLength(t *testing.T) {
	safeTest(t, "Test_S10_10_Hashset_Resize_SmallerThanLength", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})

		// Act
		result := hs.Resize(1)

		// Assert
		if result.Length() != 3 {
			t.Fatal("expected 3 — no resize when capacity < length")
		}
	})
}

func Test_S10_11_Hashset_ResizeLock(t *testing.T) {
	safeTest(t, "Test_S10_11_Hashset_ResizeLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.ResizeLock(10)

		// Assert
		if result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S10_12_Hashset_ResizeLock_SmallerThanLength(t *testing.T) {
	safeTest(t, "Test_S10_12_Hashset_ResizeLock_SmallerThanLength", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		result := hs.ResizeLock(0)

		// Assert
		if result.Length() != 2 {
			t.Fatal("expected 2 — no resize")
		}
	})
}

// ── Collection ───────────────────────────────────────────────

func Test_S10_13_Hashset_Collection(t *testing.T) {
	safeTest(t, "Test_S10_13_Hashset_Collection", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		col := hs.Collection()

		// Assert
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

// ── ConcatNewHashsets ────────────────────────────────────────

func Test_S10_14_Hashset_ConcatNewHashsets(t *testing.T) {
	safeTest(t, "Test_S10_14_Hashset_ConcatNewHashsets", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		other := corestr.New.Hashset.Strings([]string{"b"})

		// Act
		result := hs.ConcatNewHashsets(true, other)

		// Assert
		if result.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_S10_15_Hashset_ConcatNewHashsets_Empty(t *testing.T) {
	safeTest(t, "Test_S10_15_Hashset_ConcatNewHashsets_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.ConcatNewHashsets(true)

		// Assert
		if result == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S10_16_Hashset_ConcatNewHashsets_NilInList(t *testing.T) {
	safeTest(t, "Test_S10_16_Hashset_ConcatNewHashsets_NilInList", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.ConcatNewHashsets(false, nil)

		// Assert
		if result == nil {
			t.Fatal("expected non-nil")
		}
	})
}

// ── ConcatNewStrings ─────────────────────────────────────────

func Test_S10_17_Hashset_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_S10_17_Hashset_ConcatNewStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.ConcatNewStrings(true, []string{"b", "c"})

		// Assert
		if result.Length() < 3 {
			t.Fatal("expected at least 3")
		}
	})
}

func Test_S10_18_Hashset_ConcatNewStrings_Empty(t *testing.T) {
	safeTest(t, "Test_S10_18_Hashset_ConcatNewStrings_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.ConcatNewStrings(true)

		// Assert
		if result == nil {
			t.Fatal("expected non-nil")
		}
	})
}

// ── Add variants ─────────────────────────────────────────────

func Test_S10_19_Hashset_Add(t *testing.T) {
	safeTest(t, "Test_S10_19_Hashset_Add", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.Add("x")

		// Assert
		if !hs.Has("x") {
			t.Fatal("expected has x")
		}
	})
}

func Test_S10_20_Hashset_AddPtr(t *testing.T) {
	safeTest(t, "Test_S10_20_Hashset_AddPtr", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		key := "ptr-key"

		// Act
		hs.AddPtr(&key)

		// Assert
		if !hs.Has("ptr-key") {
			t.Fatal("expected has ptr-key")
		}
	})
}

func Test_S10_21_Hashset_AddPtrLock(t *testing.T) {
	safeTest(t, "Test_S10_21_Hashset_AddPtrLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		key := "ptr-lock"

		// Act
		hs.AddPtrLock(&key)

		// Assert
		if !hs.Has("ptr-lock") {
			t.Fatal("expected has ptr-lock")
		}
	})
}

func Test_S10_22_Hashset_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_S10_22_Hashset_AddWithWgLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		hs.AddWithWgLock("wg-key", wg)
		wg.Wait()

		// Assert
		if !hs.Has("wg-key") {
			t.Fatal("expected has wg-key")
		}
	})
}

func Test_S10_23_Hashset_AddBool(t *testing.T) {
	safeTest(t, "Test_S10_23_Hashset_AddBool", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		existed := hs.AddBool("k")
		existed2 := hs.AddBool("k")

		// Assert
		if existed {
			t.Fatal("expected not existed on first add")
		}
		if !existed2 {
			t.Fatal("expected existed on second add")
		}
	})
}

func Test_S10_24_Hashset_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_S10_24_Hashset_AddNonEmpty", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddNonEmpty("")
		hs.AddNonEmpty("valid")

		// Assert
		if hs.Length() != 1 {
			t.Fatal("expected 1 — empty string skipped")
		}
	})
}

func Test_S10_25_Hashset_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_S10_25_Hashset_AddNonEmptyWhitespace", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddNonEmptyWhitespace("   ")
		hs.AddNonEmptyWhitespace("valid")

		// Assert
		if hs.Length() != 1 {
			t.Fatal("expected 1 — whitespace skipped")
		}
	})
}

func Test_S10_26_Hashset_AddIf(t *testing.T) {
	safeTest(t, "Test_S10_26_Hashset_AddIf", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddIf(true, "yes")
		hs.AddIf(false, "no")

		// Assert
		if hs.Length() != 1 || !hs.Has("yes") {
			t.Fatal("expected only 'yes'")
		}
	})
}

func Test_S10_27_Hashset_AddIfMany(t *testing.T) {
	safeTest(t, "Test_S10_27_Hashset_AddIfMany", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddIfMany(true, "a", "b")
		hs.AddIfMany(false, "c", "d")

		// Assert
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S10_28_Hashset_AddFunc(t *testing.T) {
	safeTest(t, "Test_S10_28_Hashset_AddFunc", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddFunc(func() string { return "func-val" })

		// Assert
		if !hs.Has("func-val") {
			t.Fatal("expected has func-val")
		}
	})
}

func Test_S10_29_Hashset_AddFuncErr_NoError(t *testing.T) {
	safeTest(t, "Test_S10_29_Hashset_AddFuncErr_NoError", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(err error) { t.Fatal("should not call error handler") },
		)

		// Assert
		if !hs.Has("ok") {
			t.Fatal("expected has ok")
		}
	})
}

func Test_S10_30_Hashset_AddFuncErr_WithError(t *testing.T) {
	safeTest(t, "Test_S10_30_Hashset_AddFuncErr_WithError", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		called := false

		// Act
		hs.AddFuncErr(
			func() (string, error) { return "", fmt.Errorf("simulated error") },
			func(err error) { called = true },
		)

		// Assert
		if !called {
			t.Fatal("expected error handler called")
		}
		if hs.Has("") {
			// it may have "" but the err path was exercised
			_ = 0
		}
	})
}


func Test_S10_31_Hashset_AddStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_S10_31_Hashset_AddStringsPtrWgLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		hs.AddStringsPtrWgLock([]string{"a", "b"}, wg)
		wg.Wait()

		// Assert
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S10_32_Hashset_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_S10_32_Hashset_AddHashsetItems", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		other := corestr.New.Hashset.Strings([]string{"x", "y"})

		// Act
		hs.AddHashsetItems(other)

		// Assert
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S10_33_Hashset_AddHashsetItems_Nil(t *testing.T) {
	safeTest(t, "Test_S10_33_Hashset_AddHashsetItems_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddHashsetItems(nil)

		// Assert
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S10_34_Hashset_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_S10_34_Hashset_AddItemsMap", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddItemsMap(map[string]bool{"a": true, "b": false, "c": true})

		// Assert
		if hs.Length() != 2 {
			t.Fatal("expected 2 — b is disabled")
		}
	})
}

func Test_S10_35_Hashset_AddItemsMap_Nil(t *testing.T) {
	safeTest(t, "Test_S10_35_Hashset_AddItemsMap_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddItemsMap(nil)

		// Assert
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S10_36_Hashset_AddItemsMapWgLock(t *testing.T) {
	safeTest(t, "Test_S10_36_Hashset_AddItemsMapWgLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		m := map[string]bool{"a": true, "b": false}

		// Act
		hs.AddItemsMapWgLock(&m, wg)
		wg.Wait()

		// Assert
		if hs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S10_37_Hashset_AddItemsMapWgLock_Nil(t *testing.T) {
	safeTest(t, "Test_S10_37_Hashset_AddItemsMapWgLock_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddItemsMapWgLock(nil, nil)

		// Assert
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S10_38_Hashset_AddHashsetWgLock(t *testing.T) {
	safeTest(t, "Test_S10_38_Hashset_AddHashsetWgLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		other := corestr.New.Hashset.Strings([]string{"z"})

		// Act
		hs.AddHashsetWgLock(other, wg)
		wg.Wait()

		// Assert
		if hs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S10_39_Hashset_AddHashsetWgLock_Nil(t *testing.T) {
	safeTest(t, "Test_S10_39_Hashset_AddHashsetWgLock_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddHashsetWgLock(nil, nil)

		// Assert
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S10_40_Hashset_AddStrings(t *testing.T) {
	safeTest(t, "Test_S10_40_Hashset_AddStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddStrings([]string{"a", "b"})

		// Assert
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S10_41_Hashset_AddStrings_Nil(t *testing.T) {
	safeTest(t, "Test_S10_41_Hashset_AddStrings_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddStrings(nil)

		// Assert
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S10_42_Hashset_AddSimpleSlice(t *testing.T) {
	safeTest(t, "Test_S10_42_Hashset_AddSimpleSlice", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		hs.AddSimpleSlice(ss)

		// Assert
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S10_43_Hashset_AddSimpleSlice_Empty(t *testing.T) {
	safeTest(t, "Test_S10_43_Hashset_AddSimpleSlice_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		ss := corestr.Empty.SimpleSlice()

		// Act
		hs.AddSimpleSlice(ss)

		// Assert
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S10_44_Hashset_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_S10_44_Hashset_AddStringsLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddStringsLock([]string{"a"})

		// Assert
		if hs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S10_45_Hashset_AddStringsLock_Nil(t *testing.T) {
	safeTest(t, "Test_S10_45_Hashset_AddStringsLock_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddStringsLock(nil)

		// Assert
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S10_46_Hashset_Adds(t *testing.T) {
	safeTest(t, "Test_S10_46_Hashset_Adds", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.Adds("a", "b", "c")

		// Assert
		if hs.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_S10_47_Hashset_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_S10_47_Hashset_Adds_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.Adds(nil...)

		// Assert
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S10_48_Hashset_AddCollection(t *testing.T) {
	safeTest(t, "Test_S10_48_Hashset_AddCollection", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		hs.AddCollection(col)

		// Assert
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S10_49_Hashset_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_S10_49_Hashset_AddCollection_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddCollection(nil)

		// Assert
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S10_50_Hashset_AddCollections(t *testing.T) {
	safeTest(t, "Test_S10_50_Hashset_AddCollections", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})

		// Act
		hs.AddCollections(c1, nil, c2)

		// Assert
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S10_51_Hashset_AddCollections_Nil(t *testing.T) {
	safeTest(t, "Test_S10_51_Hashset_AddCollections_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddCollections(nil...)

		// Assert
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

// ── Filter-based adds ────────────────────────────────────────

func Test_S10_52_Hashset_AddsAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_S10_52_Hashset_AddsAnyUsingFilter", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		hs.AddsAnyUsingFilter(filter, "a", "b")

		// Assert
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S10_53_Hashset_AddsAnyUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_S10_53_Hashset_AddsAnyUsingFilter_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddsAnyUsingFilter(nil, nil...)

		// Assert
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S10_54_Hashset_AddsAnyUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_S10_54_Hashset_AddsAnyUsingFilter_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, true
		}

		// Act
		hs.AddsAnyUsingFilter(filter, "a", "b")

		// Assert
		if hs.Length() != 1 {
			t.Fatal("expected 1 due to break")
		}
	})
}

func Test_S10_55_Hashset_AddsAnyUsingFilter_NilItem(t *testing.T) {
	safeTest(t, "Test_S10_55_Hashset_AddsAnyUsingFilter_NilItem", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		hs.AddsAnyUsingFilter(filter, nil, "b")

		// Assert
		if hs.Length() != 1 {
			t.Fatal("expected 1 — nil skipped")
		}
	})
}

func Test_S10_56_Hashset_AddsAnyUsingFilter_Skip(t *testing.T) {
	safeTest(t, "Test_S10_56_Hashset_AddsAnyUsingFilter_Skip", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		filter := func(str string, index int) (string, bool, bool) {
			return str, false, false
		}

		// Act
		hs.AddsAnyUsingFilter(filter, "a")

		// Assert
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S10_57_Hashset_AddsAnyUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_S10_57_Hashset_AddsAnyUsingFilterLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		hs.AddsAnyUsingFilterLock(filter, "a")

		// Assert
		if hs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S10_58_Hashset_AddsAnyUsingFilterLock_Nil(t *testing.T) {
	safeTest(t, "Test_S10_58_Hashset_AddsAnyUsingFilterLock_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddsAnyUsingFilterLock(nil, nil...)

		// Assert
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S10_59_Hashset_AddsAnyUsingFilterLock_BreakAndSkip(t *testing.T) {
	safeTest(t, "Test_S10_59_Hashset_AddsAnyUsingFilterLock_BreakAndSkip", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		callCount := 0
		filter := func(str string, index int) (string, bool, bool) {
			callCount++
			if callCount == 1 {
				return "", false, false // skip
			}
			return str, true, true // keep + break
		}

		// Act
		hs.AddsAnyUsingFilterLock(filter, "a", "b", "c")

		// Assert
		if hs.Length() != 1 {
			t.Fatalf("expected 1, got %d", hs.Length())
		}
	})
}

func Test_S10_60_Hashset_AddsAnyUsingFilterLock_NilItem(t *testing.T) {
	safeTest(t, "Test_S10_60_Hashset_AddsAnyUsingFilterLock_NilItem", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		hs.AddsAnyUsingFilterLock(filter, nil, "b")

		// Assert
		if hs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S10_61_Hashset_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_S10_61_Hashset_AddsUsingFilter", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		hs.AddsUsingFilter(filter, "a", "b")

		// Assert
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S10_62_Hashset_AddsUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_S10_62_Hashset_AddsUsingFilter_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddsUsingFilter(nil, nil...)

		// Assert
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S10_63_Hashset_AddsUsingFilter_BreakAndSkip(t *testing.T) {
	safeTest(t, "Test_S10_63_Hashset_AddsUsingFilter_BreakAndSkip", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		callCount := 0
		filter := func(str string, index int) (string, bool, bool) {
			callCount++
			if callCount == 1 {
				return "", false, false
			}
			return str, true, true
		}

		// Act
		hs.AddsUsingFilter(filter, "a", "b", "c")

		// Assert
		if hs.Length() != 1 {
			t.Fatalf("expected 1, got %d", hs.Length())
		}
	})
}

// ── AddLock ──────────────────────────────────────────────────

func Test_S10_64_Hashset_AddLock(t *testing.T) {
	safeTest(t, "Test_S10_64_Hashset_AddLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddLock("k")

		// Assert
		if !hs.Has("k") {
			t.Fatal("expected has k")
		}
	})
}

// ── Has / Contains / HasLock / HasWithLock / HasAnyItem ──────

func Test_S10_65_Hashset_Has(t *testing.T) {
	safeTest(t, "Test_S10_65_Hashset_Has", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		if !hs.Has("a") {
			t.Fatal("expected true")
		}
		if hs.Has("z") {
			t.Fatal("expected false")
		}
	})
}

func Test_S10_66_Hashset_Contains(t *testing.T) {
	safeTest(t, "Test_S10_66_Hashset_Contains", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		if !hs.Contains("a") {
			t.Fatal("expected true")
		}
	})
}

func Test_S10_67_Hashset_HasLock(t *testing.T) {
	safeTest(t, "Test_S10_67_Hashset_HasLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		if !hs.HasLock("a") {
			t.Fatal("expected true")
		}
	})
}

func Test_S10_68_Hashset_HasWithLock(t *testing.T) {
	safeTest(t, "Test_S10_68_Hashset_HasWithLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		if !hs.HasWithLock("a") {
			t.Fatal("expected true")
		}
	})
}

func Test_S10_69_Hashset_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_S10_69_Hashset_HasAnyItem", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		if !hs.HasAnyItem() {
			t.Fatal("expected true")
		}
		if corestr.Empty.Hashset().HasAnyItem() {
			t.Fatal("expected false for empty")
		}
	})
}

// ── IsMissing / IsMissingLock ────────────────────────────────

func Test_S10_70_Hashset_IsMissing(t *testing.T) {
	safeTest(t, "Test_S10_70_Hashset_IsMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		if hs.IsMissing("a") {
			t.Fatal("expected false")
		}
		if !hs.IsMissing("z") {
			t.Fatal("expected true")
		}
	})
}

func Test_S10_71_Hashset_IsMissingLock(t *testing.T) {
	safeTest(t, "Test_S10_71_Hashset_IsMissingLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		if hs.IsMissingLock("a") {
			t.Fatal("expected false")
		}
	})
}

// ── HasAllStrings / HasAll / HasAllCollectionItems ────────────

func Test_S10_72_Hashset_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_S10_72_Hashset_HasAllStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act & Assert
		if !hs.HasAllStrings([]string{"a", "b"}) {
			t.Fatal("expected true")
		}
		if hs.HasAllStrings([]string{"a", "c"}) {
			t.Fatal("expected false")
		}
	})
}

func Test_S10_73_Hashset_HasAll(t *testing.T) {
	safeTest(t, "Test_S10_73_Hashset_HasAll", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act & Assert
		if !hs.HasAll("a", "b") {
			t.Fatal("expected true")
		}
		if hs.HasAll("a", "c") {
			t.Fatal("expected false")
		}
	})
}

func Test_S10_74_Hashset_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_S10_74_Hashset_HasAllCollectionItems", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act & Assert
		if !hs.HasAllCollectionItems(col) {
			t.Fatal("expected true")
		}
	})
}

func Test_S10_75_Hashset_HasAllCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_S10_75_Hashset_HasAllCollectionItems_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		if hs.HasAllCollectionItems(nil) {
			t.Fatal("expected false")
		}
	})
}

func Test_S10_76_Hashset_HasAllCollectionItems_Empty(t *testing.T) {
	safeTest(t, "Test_S10_76_Hashset_HasAllCollectionItems_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		if hs.HasAllCollectionItems(corestr.Empty.Collection()) {
			t.Fatal("expected false")
		}
	})
}

// ── IsAllMissing / HasAny ────────────────────────────────────

func Test_S10_77_Hashset_IsAllMissing(t *testing.T) {
	safeTest(t, "Test_S10_77_Hashset_IsAllMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		if !hs.IsAllMissing("x", "y") {
			t.Fatal("expected true")
		}
		if hs.IsAllMissing("a", "y") {
			t.Fatal("expected false — a exists")
		}
	})
}

func Test_S10_78_Hashset_HasAny(t *testing.T) {
	safeTest(t, "Test_S10_78_Hashset_HasAny", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		if !hs.HasAny("x", "a") {
			t.Fatal("expected true")
		}
		if hs.HasAny("x", "y") {
			t.Fatal("expected false")
		}
	})
}

// ── IsEqual / IsEquals / IsEqualsLock ────────────────────────

func Test_S10_79_Hashset_IsEqual(t *testing.T) {
	safeTest(t, "Test_S10_79_Hashset_IsEqual", func() {
		// Arrange
		a := corestr.New.Hashset.Strings([]string{"a"})
		b := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		if !a.IsEqual(b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S10_80_Hashset_IsEquals_BothNil(t *testing.T) {
	safeTest(t, "Test_S10_80_Hashset_IsEquals_BothNil", func() {
		// Arrange
		var a *corestr.Hashset
		var b *corestr.Hashset

		// Act & Assert
		if !a.IsEquals(b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S10_81_Hashset_IsEquals_OneNil(t *testing.T) {
	safeTest(t, "Test_S10_81_Hashset_IsEquals_OneNil", func() {
		// Arrange
		a := corestr.New.Hashset.Strings([]string{"a"})
		var b *corestr.Hashset

		// Act & Assert
		if a.IsEquals(b) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S10_82_Hashset_IsEquals_SamePtr(t *testing.T) {
	safeTest(t, "Test_S10_82_Hashset_IsEquals_SamePtr", func() {
		// Arrange
		a := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		if !a.IsEquals(a) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S10_83_Hashset_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S10_83_Hashset_IsEquals_BothEmpty", func() {
		// Arrange
		a := corestr.Empty.Hashset()
		b := corestr.Empty.Hashset()

		// Act & Assert
		if !a.IsEquals(b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S10_84_Hashset_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_S10_84_Hashset_IsEquals_OneEmpty", func() {
		// Arrange
		a := corestr.New.Hashset.Strings([]string{"a"})
		b := corestr.Empty.Hashset()

		// Act & Assert
		if a.IsEquals(b) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S10_85_Hashset_IsEquals_DiffLength(t *testing.T) {
	safeTest(t, "Test_S10_85_Hashset_IsEquals_DiffLength", func() {
		// Arrange
		a := corestr.New.Hashset.Strings([]string{"a"})
		b := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act & Assert
		if a.IsEquals(b) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S10_86_Hashset_IsEquals_DiffKeys(t *testing.T) {
	safeTest(t, "Test_S10_86_Hashset_IsEquals_DiffKeys", func() {
		// Arrange
		a := corestr.New.Hashset.Strings([]string{"a"})
		b := corestr.New.Hashset.Strings([]string{"b"})

		// Act & Assert
		if a.IsEquals(b) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S10_87_Hashset_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_S10_87_Hashset_IsEqualsLock", func() {
		// Arrange
		a := corestr.New.Hashset.Strings([]string{"a"})
		b := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		if !a.IsEqualsLock(b) {
			t.Fatal("expected equal")
		}
	})
}
