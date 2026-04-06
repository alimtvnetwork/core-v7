package corestrtests

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
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
		actual := args.Map{"result": empty.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual := args.Map{"result": hs.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_S10_02_Hashset_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_S10_02_Hashset_IsEmpty_Nil", func() {
		// Arrange
		var hs *corestr.Hashset

		// Act & Assert
		actual := args.Map{"result": hs.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	})
}

func Test_S10_03_Hashset_HasItems(t *testing.T) {
	safeTest(t, "Test_S10_03_Hashset_HasItems", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": hs.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has items", actual)
		actual := args.Map{"result": corestr.Empty.Hashset().HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items for empty", actual)
	})
}

func Test_S10_04_Hashset_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_S10_04_Hashset_IsEmptyLock", func() {
		// Arrange
		hs := corestr.Empty.Hashset()

		// Act & Assert
		actual := args.Map{"result": hs.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
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
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 item preserved", actual)
	})
}

func Test_S10_06_Hashset_AddCapacities_Empty(t *testing.T) {
	safeTest(t, "Test_S10_06_Hashset_AddCapacities_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(2)

		// Act
		result := hs.AddCapacities()

		// Assert
		actual := args.Map{"result": result != hs}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same pointer when no capacities", actual)
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
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 preserved", actual)
	})
}

func Test_S10_08_Hashset_AddCapacitiesLock_Empty(t *testing.T) {
	safeTest(t, "Test_S10_08_Hashset_AddCapacitiesLock_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(2)

		// Act
		result := hs.AddCapacitiesLock()

		// Assert
		actual := args.Map{"result": result != hs}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same pointer", actual)
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
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S10_10_Hashset_Resize_SmallerThanLength(t *testing.T) {
	safeTest(t, "Test_S10_10_Hashset_Resize_SmallerThanLength", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})

		// Act
		result := hs.Resize(1)

		// Assert
		actual := args.Map{"result": result.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3 — no resize when capacity < length", actual)
	})
}

func Test_S10_11_Hashset_ResizeLock(t *testing.T) {
	safeTest(t, "Test_S10_11_Hashset_ResizeLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.ResizeLock(10)

		// Assert
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S10_12_Hashset_ResizeLock_SmallerThanLength(t *testing.T) {
	safeTest(t, "Test_S10_12_Hashset_ResizeLock_SmallerThanLength", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		result := hs.ResizeLock(0)

		// Assert
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 — no resize", actual)
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
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": result.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_S10_15_Hashset_ConcatNewHashsets_Empty(t *testing.T) {
	safeTest(t, "Test_S10_15_Hashset_ConcatNewHashsets_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.ConcatNewHashsets(true)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S10_16_Hashset_ConcatNewHashsets_NilInList(t *testing.T) {
	safeTest(t, "Test_S10_16_Hashset_ConcatNewHashsets_NilInList", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.ConcatNewHashsets(false, nil)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
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
		actual := args.Map{"result": result.Length() < 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	})
}

func Test_S10_18_Hashset_ConcatNewStrings_Empty(t *testing.T) {
	safeTest(t, "Test_S10_18_Hashset_ConcatNewStrings_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		result := hs.ConcatNewStrings(true)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
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
		actual := args.Map{"result": hs.Has("x")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has x", actual)
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
		actual := args.Map{"result": hs.Has("ptr-key")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has ptr-key", actual)
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
		actual := args.Map{"result": hs.Has("ptr-lock")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has ptr-lock", actual)
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
		actual := args.Map{"result": hs.Has("wg-key")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has wg-key", actual)
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
		actual := args.Map{"result": existed}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not existed on first add", actual)
		actual := args.Map{"result": existed2}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected existed on second add", actual)
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
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 — empty string skipped", actual)
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
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 — whitespace skipped", actual)
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
		actual := args.Map{"result": hs.Length() != 1 || !hs.Has("yes")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected only 'yes'", actual)
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
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S10_28_Hashset_AddFunc(t *testing.T) {
	safeTest(t, "Test_S10_28_Hashset_AddFunc", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddFunc(func() string { return "func-val" })

		// Assert
		actual := args.Map{"result": hs.Has("func-val")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has func-val", actual)
	})
}

func Test_S10_29_Hashset_AddFuncErr_NoError(t *testing.T) {
	safeTest(t, "Test_S10_29_Hashset_AddFuncErr_NoError", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(err error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) },
		)

		// Assert
		actual := args.Map{"result": hs.Has("ok")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has ok", actual)
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
		actual := args.Map{"result": called}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected error handler called", actual)
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
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S10_33_Hashset_AddHashsetItems_Nil(t *testing.T) {
	safeTest(t, "Test_S10_33_Hashset_AddHashsetItems_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddHashsetItems(nil)

		// Assert
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S10_34_Hashset_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_S10_34_Hashset_AddItemsMap", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddItemsMap(map[string]bool{"a": true, "b": false, "c": true})

		// Assert
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 — b is disabled", actual)
	})
}

func Test_S10_35_Hashset_AddItemsMap_Nil(t *testing.T) {
	safeTest(t, "Test_S10_35_Hashset_AddItemsMap_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddItemsMap(nil)

		// Assert
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S10_37_Hashset_AddItemsMapWgLock_Nil(t *testing.T) {
	safeTest(t, "Test_S10_37_Hashset_AddItemsMapWgLock_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddItemsMapWgLock(nil, nil)

		// Assert
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S10_39_Hashset_AddHashsetWgLock_Nil(t *testing.T) {
	safeTest(t, "Test_S10_39_Hashset_AddHashsetWgLock_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddHashsetWgLock(nil, nil)

		// Assert
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S10_40_Hashset_AddStrings(t *testing.T) {
	safeTest(t, "Test_S10_40_Hashset_AddStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddStrings([]string{"a", "b"})

		// Assert
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S10_41_Hashset_AddStrings_Nil(t *testing.T) {
	safeTest(t, "Test_S10_41_Hashset_AddStrings_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddStrings(nil)

		// Assert
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S10_44_Hashset_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_S10_44_Hashset_AddStringsLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddStringsLock([]string{"a"})

		// Assert
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S10_45_Hashset_AddStringsLock_Nil(t *testing.T) {
	safeTest(t, "Test_S10_45_Hashset_AddStringsLock_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddStringsLock(nil)

		// Assert
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S10_46_Hashset_Adds(t *testing.T) {
	safeTest(t, "Test_S10_46_Hashset_Adds", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.Adds("a", "b", "c")

		// Assert
		actual := args.Map{"result": hs.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_S10_47_Hashset_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_S10_47_Hashset_Adds_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.Adds(nil...)

		// Assert
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S10_49_Hashset_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_S10_49_Hashset_AddCollection_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddCollection(nil)

		// Assert
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S10_51_Hashset_AddCollections_Nil(t *testing.T) {
	safeTest(t, "Test_S10_51_Hashset_AddCollections_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddCollections(nil...)

		// Assert
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S10_53_Hashset_AddsAnyUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_S10_53_Hashset_AddsAnyUsingFilter_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddsAnyUsingFilter(nil, nil...)

		// Assert
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 due to break", actual)
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
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 — nil skipped", actual)
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
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S10_58_Hashset_AddsAnyUsingFilterLock_Nil(t *testing.T) {
	safeTest(t, "Test_S10_58_Hashset_AddsAnyUsingFilterLock_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddsAnyUsingFilterLock(nil, nil...)

		// Assert
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S10_62_Hashset_AddsUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_S10_62_Hashset_AddsUsingFilter_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		hs.AddsUsingFilter(nil, nil...)

		// Assert
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
		actual := args.Map{"result": hs.Has("k")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has k", actual)
	})
}

// ── Has / Contains / HasLock / HasWithLock / HasAnyItem ──────

func Test_S10_65_Hashset_Has(t *testing.T) {
	safeTest(t, "Test_S10_65_Hashset_Has", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": hs.Has("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.Has("z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_S10_66_Hashset_Contains(t *testing.T) {
	safeTest(t, "Test_S10_66_Hashset_Contains", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": hs.Contains("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_S10_67_Hashset_HasLock(t *testing.T) {
	safeTest(t, "Test_S10_67_Hashset_HasLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": hs.HasLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_S10_68_Hashset_HasWithLock(t *testing.T) {
	safeTest(t, "Test_S10_68_Hashset_HasWithLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": hs.HasWithLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_S10_69_Hashset_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_S10_69_Hashset_HasAnyItem", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": hs.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": corestr.Empty.Hashset().HasAnyItem()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

// ── IsMissing / IsMissingLock ────────────────────────────────

func Test_S10_70_Hashset_IsMissing(t *testing.T) {
	safeTest(t, "Test_S10_70_Hashset_IsMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": hs.IsMissing("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual := args.Map{"result": hs.IsMissing("z")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_S10_71_Hashset_IsMissingLock(t *testing.T) {
	safeTest(t, "Test_S10_71_Hashset_IsMissingLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": hs.IsMissingLock("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

// ── HasAllStrings / HasAll / HasAllCollectionItems ────────────

func Test_S10_72_Hashset_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_S10_72_Hashset_HasAllStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": hs.HasAllStrings([]string{"a", "b"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.HasAllStrings([]string{"a", "c"})}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_S10_73_Hashset_HasAll(t *testing.T) {
	safeTest(t, "Test_S10_73_Hashset_HasAll", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": hs.HasAll("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.HasAll("a", "c")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_S10_74_Hashset_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_S10_74_Hashset_HasAllCollectionItems", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": hs.HasAllCollectionItems(col)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_S10_75_Hashset_HasAllCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_S10_75_Hashset_HasAllCollectionItems_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": hs.HasAllCollectionItems(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_S10_76_Hashset_HasAllCollectionItems_Empty(t *testing.T) {
	safeTest(t, "Test_S10_76_Hashset_HasAllCollectionItems_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": hs.HasAllCollectionItems(corestr.Empty.Collection())}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

// ── IsAllMissing / HasAny ────────────────────────────────────

func Test_S10_77_Hashset_IsAllMissing(t *testing.T) {
	safeTest(t, "Test_S10_77_Hashset_IsAllMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": hs.IsAllMissing("x", "y")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.IsAllMissing("a", "y")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false — a exists", actual)
	})
}

func Test_S10_78_Hashset_HasAny(t *testing.T) {
	safeTest(t, "Test_S10_78_Hashset_HasAny", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": hs.HasAny("x", "a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.HasAny("x", "y")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

// ── IsEqual / IsEquals / IsEqualsLock ────────────────────────

func Test_S10_79_Hashset_IsEqual(t *testing.T) {
	safeTest(t, "Test_S10_79_Hashset_IsEqual", func() {
		// Arrange
		a := corestr.New.Hashset.Strings([]string{"a"})
		b := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": a.IsEqual(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S10_80_Hashset_IsEquals_BothNil(t *testing.T) {
	safeTest(t, "Test_S10_80_Hashset_IsEquals_BothNil", func() {
		// Arrange
		var a *corestr.Hashset
		var b *corestr.Hashset

		// Act & Assert
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S10_81_Hashset_IsEquals_OneNil(t *testing.T) {
	safeTest(t, "Test_S10_81_Hashset_IsEquals_OneNil", func() {
		// Arrange
		a := corestr.New.Hashset.Strings([]string{"a"})
		var b *corestr.Hashset

		// Act & Assert
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S10_82_Hashset_IsEquals_SamePtr(t *testing.T) {
	safeTest(t, "Test_S10_82_Hashset_IsEquals_SamePtr", func() {
		// Arrange
		a := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": a.IsEquals(a)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S10_83_Hashset_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S10_83_Hashset_IsEquals_BothEmpty", func() {
		// Arrange
		a := corestr.Empty.Hashset()
		b := corestr.Empty.Hashset()

		// Act & Assert
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S10_84_Hashset_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_S10_84_Hashset_IsEquals_OneEmpty", func() {
		// Arrange
		a := corestr.New.Hashset.Strings([]string{"a"})
		b := corestr.Empty.Hashset()

		// Act & Assert
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S10_85_Hashset_IsEquals_DiffLength(t *testing.T) {
	safeTest(t, "Test_S10_85_Hashset_IsEquals_DiffLength", func() {
		// Arrange
		a := corestr.New.Hashset.Strings([]string{"a"})
		b := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S10_86_Hashset_IsEquals_DiffKeys(t *testing.T) {
	safeTest(t, "Test_S10_86_Hashset_IsEquals_DiffKeys", func() {
		// Arrange
		a := corestr.New.Hashset.Strings([]string{"a"})
		b := corestr.New.Hashset.Strings([]string{"b"})

		// Act & Assert
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S10_87_Hashset_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_S10_87_Hashset_IsEqualsLock", func() {
		// Arrange
		a := corestr.New.Hashset.Strings([]string{"a"})
		b := corestr.New.Hashset.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": a.IsEqualsLock(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}
