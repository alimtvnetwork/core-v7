package corestrtests

import (
	"strings"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════
// S08b — Collection Part 2: Lines 700–1600
// ══════════════════════════════════════════════════════════════

// ── AppendAnysLock / AppendAnys ──────────────────────────────

func Test_S08b_01_Collection_AppendAnysLock(t *testing.T) {
	safeTest(t, "Test_S08b_01_Collection_AppendAnysLock", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendAnysLock("a", 42, nil)

		// Assert
		if col.Length() != 2 {
			t.Fatalf("expected 2, got %d", col.Length())
		}
	})
}

func Test_S08b_02_Collection_AppendAnysLock_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_02_Collection_AppendAnysLock_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendAnysLock()

		// Assert
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S08b_03_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_S08b_03_Collection_AppendAnys", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendAnys("hello", 123)

		// Assert
		if col.Length() != 2 {
			t.Fatalf("expected 2, got %d", col.Length())
		}
	})
}

func Test_S08b_04_Collection_AppendAnys_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_04_Collection_AppendAnys_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendAnys()

		// Assert
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S08b_05_Collection_AppendAnys_WithNil(t *testing.T) {
	safeTest(t, "Test_S08b_05_Collection_AppendAnys_WithNil", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendAnys(nil, "ok", nil)

		// Assert
		if col.Length() != 1 {
			t.Fatalf("expected 1, got %d", col.Length())
		}
	})
}

// ── AppendAnysUsingFilter / AppendAnysUsingFilterLock ────────

func Test_S08b_06_Collection_AppendAnysUsingFilter(t *testing.T) {
	safeTest(t, "Test_S08b_06_Collection_AppendAnysUsingFilter", func() {
		// Arrange
		col := corestr.Empty.Collection()
		filter := func(str string, index int) (string, bool, bool) {
			return strings.ToUpper(str), true, false
		}

		// Act
		col.AppendAnysUsingFilter(filter, "hello", nil)

		// Assert
		if col.Length() != 1 {
			t.Fatalf("expected 1, got %d", col.Length())
		}
	})
}

func Test_S08b_07_Collection_AppendAnysUsingFilter_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_07_Collection_AppendAnysUsingFilter_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendAnysUsingFilter(nil)

		// Assert
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S08b_08_Collection_AppendAnysUsingFilter_Skip(t *testing.T) {
	safeTest(t, "Test_S08b_08_Collection_AppendAnysUsingFilter_Skip", func() {
		// Arrange
		col := corestr.Empty.Collection()
		filter := func(str string, index int) (string, bool, bool) {
			return "", false, false // skip
		}

		// Act
		col.AppendAnysUsingFilter(filter, "a")

		// Assert
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S08b_09_Collection_AppendAnysUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_S08b_09_Collection_AppendAnysUsingFilter_Break", func() {
		// Arrange
		col := corestr.Empty.Collection()
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, true // keep and break
		}

		// Act
		col.AppendAnysUsingFilter(filter, "a", "b", "c")

		// Assert
		if col.Length() != 1 {
			t.Fatalf("expected 1, got %d", col.Length())
		}
	})
}

func Test_S08b_10_Collection_AppendAnysUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_S08b_10_Collection_AppendAnysUsingFilterLock", func() {
		// Arrange
		col := corestr.Empty.Collection()
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		col.AppendAnysUsingFilterLock(filter, "x", nil, "y")

		// Assert
		if col.Length() != 2 {
			t.Fatalf("expected 2, got %d", col.Length())
		}
	})
}

func Test_S08b_11_Collection_AppendAnysUsingFilterLock_Nil(t *testing.T) {
	safeTest(t, "Test_S08b_11_Collection_AppendAnysUsingFilterLock_Nil", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendAnysUsingFilterLock(nil, nil)

		// Assert
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S08b_12_Collection_AppendAnysUsingFilterLock_SkipAndBreak(t *testing.T) {
	safeTest(t, "Test_S08b_12_Collection_AppendAnysUsingFilterLock_SkipAndBreak", func() {
		// Arrange
		col := corestr.Empty.Collection()
		callCount := 0
		filter := func(str string, index int) (string, bool, bool) {
			callCount++
			if callCount == 1 {
				return "", false, false // skip
			}
			return str, true, true // keep and break
		}

		// Act
		col.AppendAnysUsingFilterLock(filter, "a", "b", "c")

		// Assert
		if col.Length() != 1 {
			t.Fatalf("expected 1, got %d", col.Length())
		}
	})
}

// ── AppendNonEmptyAnys ───────────────────────────────────────

func Test_S08b_13_Collection_AppendNonEmptyAnys(t *testing.T) {
	safeTest(t, "Test_S08b_13_Collection_AppendNonEmptyAnys", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendNonEmptyAnys("hello", nil, "", "world")

		// Assert
		if col.Length() != 2 {
			t.Fatalf("expected 2, got %d", col.Length())
		}
	})
}

func Test_S08b_14_Collection_AppendNonEmptyAnys_Nil(t *testing.T) {
	safeTest(t, "Test_S08b_14_Collection_AppendNonEmptyAnys_Nil", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AppendNonEmptyAnys(nil)

		// Assert
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

// ── AddsAsync ────────────────────────────────────────────────

func Test_S08b_15_Collection_AddsAsync(t *testing.T) {
	safeTest(t, "Test_S08b_15_Collection_AddsAsync", func() {
		// Arrange
		col := corestr.Empty.Collection()
		wg := &sync.WaitGroup{}
		wg.Add(1)

		// Act
		col.AddsAsync(wg, "a", "b")
		wg.Wait()

		// Assert
		if col.Length() < 1 {
			t.Fatal("expected items added")
		}
	})
}

func Test_S08b_16_Collection_AddsAsync_Nil(t *testing.T) {
	safeTest(t, "Test_S08b_16_Collection_AddsAsync_Nil", func() {
		// Arrange
		col := corestr.Empty.Collection()
		wg := &sync.WaitGroup{}

		// Act
		col.AddsAsync(wg, nil...)

		// Assert
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

// ── AddsNonEmpty ─────────────────────────────────────────────

func Test_S08b_17_Collection_AddsNonEmpty(t *testing.T) {
	safeTest(t, "Test_S08b_17_Collection_AddsNonEmpty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddsNonEmpty("a", "", "b", "")

		// Assert
		if col.Length() != 2 {
			t.Fatalf("expected 2, got %d", col.Length())
		}
	})
}

func Test_S08b_18_Collection_AddsNonEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_S08b_18_Collection_AddsNonEmpty_Nil", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddsNonEmpty(nil...)

		// Assert
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

// ── AddsNonEmptyPtrLock ──────────────────────────────────────

func Test_S08b_19_Collection_AddsNonEmptyPtrLock(t *testing.T) {
	safeTest(t, "Test_S08b_19_Collection_AddsNonEmptyPtrLock", func() {
		// Arrange
		col := corestr.Empty.Collection()
		s1 := "hello"
		s2 := ""

		// Act
		col.AddsNonEmptyPtrLock(&s1, nil, &s2)

		// Assert
		if col.Length() != 1 {
			t.Fatalf("expected 1, got %d", col.Length())
		}
	})
}

func Test_S08b_20_Collection_AddsNonEmptyPtrLock_Nil(t *testing.T) {
	safeTest(t, "Test_S08b_20_Collection_AddsNonEmptyPtrLock_Nil", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddsNonEmptyPtrLock(nil)

		// Assert
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

// ── UniqueBoolMap / UniqueList / UniqueBoolMapLock / UniqueListLock ──

func Test_S08b_21_Collection_UniqueBoolMap(t *testing.T) {
	safeTest(t, "Test_S08b_21_Collection_UniqueBoolMap", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "a"})

		// Act
		m := col.UniqueBoolMap()

		// Assert
		if len(m) != 2 {
			t.Fatalf("expected 2, got %d", len(m))
		}
	})
}

func Test_S08b_22_Collection_UniqueBoolMapLock(t *testing.T) {
	safeTest(t, "Test_S08b_22_Collection_UniqueBoolMapLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"x", "x"})

		// Act
		m := col.UniqueBoolMapLock()

		// Assert
		if len(m) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S08b_23_Collection_UniqueList(t *testing.T) {
	safeTest(t, "Test_S08b_23_Collection_UniqueList", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "a"})

		// Act
		list := col.UniqueList()

		// Assert
		if len(list) != 2 {
			t.Fatalf("expected 2, got %d", len(list))
		}
	})
}

func Test_S08b_24_Collection_UniqueListLock(t *testing.T) {
	safeTest(t, "Test_S08b_24_Collection_UniqueListLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "a"})

		// Act
		list := col.UniqueListLock()

		// Assert
		if len(list) != 1 {
			t.Fatal("expected 1")
		}
	})
}

// ── List ─────────────────────────────────────────────────────

func Test_S08b_25_Collection_List(t *testing.T) {
	safeTest(t, "Test_S08b_25_Collection_List", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		list := col.List()

		// Assert
		if len(list) != 2 {
			t.Fatal("expected 2")
		}
	})
}

// ── Filter / FilterLock / FilteredCollection / FilteredCollectionLock ──

func Test_S08b_26_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_S08b_26_Collection_Filter", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"apple", "banana", "avocado"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, strings.HasPrefix(str, "a"), false
		}

		// Act
		result := col.Filter(filter)

		// Assert
		if len(result) != 2 {
			t.Fatalf("expected 2, got %d", len(result))
		}
	})
}

func Test_S08b_27_Collection_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_27_Collection_Filter_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		result := col.Filter(func(str string, index int) (string, bool, bool) {
			return str, true, false
		})

		// Assert
		if len(result) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S08b_28_Collection_Filter_Break(t *testing.T) {
	safeTest(t, "Test_S08b_28_Collection_Filter_Break", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, index == 0 // break after first
		}

		// Act
		result := col.Filter(filter)

		// Assert
		if len(result) != 1 {
			t.Fatalf("expected 1, got %d", len(result))
		}
	})
}

func Test_S08b_29_Collection_FilterLock(t *testing.T) {
	safeTest(t, "Test_S08b_29_Collection_FilterLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		result := col.FilterLock(filter)

		// Assert
		if len(result) != 2 {
			t.Fatalf("expected 2, got %d", len(result))
		}
	})
}

func Test_S08b_30_Collection_FilterLock_Break(t *testing.T) {
	safeTest(t, "Test_S08b_30_Collection_FilterLock_Break", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, true // keep and break
		}

		// Act
		result := col.FilterLock(filter)

		// Assert
		if len(result) != 1 {
			t.Fatalf("expected 1, got %d", len(result))
		}
	})
}

func Test_S08b_31_Collection_FilteredCollection(t *testing.T) {
	safeTest(t, "Test_S08b_31_Collection_FilteredCollection", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"x", "y", "z"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, str != "y", false
		}

		// Act
		result := col.FilteredCollection(filter)

		// Assert
		if result.Length() != 2 {
			t.Fatalf("expected 2, got %d", result.Length())
		}
	})
}

func Test_S08b_32_Collection_FilteredCollectionLock(t *testing.T) {
	safeTest(t, "Test_S08b_32_Collection_FilteredCollectionLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}

		// Act
		result := col.FilteredCollectionLock(filter)

		// Assert
		if result.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

// ── FilterPtrLock / FilterPtr ────────────────────────────────

func Test_S08b_33_Collection_FilterPtrLock(t *testing.T) {
	safeTest(t, "Test_S08b_33_Collection_FilterPtrLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		filter := func(sp *string, index int) (*string, bool, bool) {
			return sp, true, false
		}

		// Act
		result := col.FilterPtrLock(filter)

		// Assert
		if len(*result) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S08b_34_Collection_FilterPtrLock_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_34_Collection_FilterPtrLock_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()
		filter := func(sp *string, index int) (*string, bool, bool) {
			return sp, true, false
		}

		// Act
		result := col.FilterPtrLock(filter)

		// Assert
		if len(*result) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S08b_35_Collection_FilterPtrLock_Break(t *testing.T) {
	safeTest(t, "Test_S08b_35_Collection_FilterPtrLock_Break", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		filter := func(sp *string, index int) (*string, bool, bool) {
			return sp, true, true
		}

		// Act
		result := col.FilterPtrLock(filter)

		// Assert
		if len(*result) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S08b_36_Collection_FilterPtr(t *testing.T) {
	safeTest(t, "Test_S08b_36_Collection_FilterPtr", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		filter := func(sp *string, index int) (*string, bool, bool) {
			return sp, true, false
		}

		// Act
		result := col.FilterPtr(filter)

		// Assert
		if len(*result) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S08b_37_Collection_FilterPtr_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_37_Collection_FilterPtr_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()
		filter := func(sp *string, index int) (*string, bool, bool) {
			return sp, true, false
		}

		// Act
		result := col.FilterPtr(filter)

		// Assert
		if len(*result) != 0 {
			t.Fatal("expected 0")
		}
	})
}
func Test_S08b_39_Collection_NonEmptyList(t *testing.T) {
	safeTest(t, "Test_S08b_39_Collection_NonEmptyList", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "", "b", ""})

		// Act
		list := col.NonEmptyList()

		// Assert
		if len(list) != 2 {
			t.Fatalf("expected 2, got %d", len(list))
		}
	})
}

func Test_S08b_40_Collection_NonEmptyList_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_40_Collection_NonEmptyList_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		list := col.NonEmptyList()

		// Assert
		if len(list) != 0 {
			t.Fatal("expected 0")
		}
	})
}
func Test_S08b_42_Collection_HashsetAsIs(t *testing.T) {
	safeTest(t, "Test_S08b_42_Collection_HashsetAsIs", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "a"})

		// Act
		hs := col.HashsetAsIs()

		// Assert
		if hs.Length() != 2 {
			t.Fatalf("expected 2 unique, got %d", hs.Length())
		}
	})
}

func Test_S08b_43_Collection_HashsetWithDoubleLength(t *testing.T) {
	safeTest(t, "Test_S08b_43_Collection_HashsetWithDoubleLength", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		hs := col.HashsetWithDoubleLength()

		// Assert
		if hs == nil {
			t.Fatal("expected non-nil")
		}
	})
}
func Test_S08b_45_Collection_NonEmptyItems(t *testing.T) {
	safeTest(t, "Test_S08b_45_Collection_NonEmptyItems", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "", "b"})

		// Act
		items := col.NonEmptyItems()

		// Assert
		if len(items) != 2 {
			t.Fatalf("expected 2, got %d", len(items))
		}
	})
}
func Test_S08b_47_Collection_NonEmptyItemsOrNonWhitespace(t *testing.T) {
	safeTest(t, "Test_S08b_47_Collection_NonEmptyItemsOrNonWhitespace", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "  ", "", "b"})

		// Act
		items := col.NonEmptyItemsOrNonWhitespace()

		// Assert
		if len(items) != 2 {
			t.Fatalf("expected 2, got %d", len(items))
		}
	})
}
func Test_S08b_49_Collection_Items(t *testing.T) {
	safeTest(t, "Test_S08b_49_Collection_Items", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act & Assert
		if len(col.Items()) != 1 {
			t.Fatal("expected 1")
		}
	})
}
func Test_S08b_51_Collection_ListCopyPtrLock(t *testing.T) {
	safeTest(t, "Test_S08b_51_Collection_ListCopyPtrLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		items := col.ListCopyPtrLock()

		// Assert
		if len(items) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S08b_52_Collection_ListCopyPtrLock_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_52_Collection_ListCopyPtrLock_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		items := col.ListCopyPtrLock()

		// Assert
		if len(items) != 0 {
			t.Fatal("expected 0")
		}
	})
}

// ── Has / HasLock / HasPtr / HasAll / HasUsingSensitivity ────

func Test_S08b_53_Collection_Has(t *testing.T) {
	safeTest(t, "Test_S08b_53_Collection_Has", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act & Assert
		if !col.Has("a") {
			t.Fatal("expected true")
		}
		if col.Has("z") {
			t.Fatal("expected false")
		}
	})
}

func Test_S08b_54_Collection_Has_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_54_Collection_Has_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act & Assert
		if col.Has("a") {
			t.Fatal("expected false for empty")
		}
	})
}

func Test_S08b_55_Collection_HasLock(t *testing.T) {
	safeTest(t, "Test_S08b_55_Collection_HasLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act & Assert
		if !col.HasLock("a") {
			t.Fatal("expected true")
		}
	})
}

func Test_S08b_56_Collection_HasPtr(t *testing.T) {
	safeTest(t, "Test_S08b_56_Collection_HasPtr", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"x"})
		s := "x"
		missing := "z"

		// Act & Assert
		if !col.HasPtr(&s) {
			t.Fatal("expected true")
		}
		if col.HasPtr(&missing) {
			t.Fatal("expected false")
		}
		if col.HasPtr(nil) {
			t.Fatal("expected false for nil")
		}
	})
}

func Test_S08b_57_Collection_HasPtr_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_57_Collection_HasPtr_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()
		s := "a"

		// Act & Assert
		if col.HasPtr(&s) {
			t.Fatal("expected false for empty")
		}
	})
}

func Test_S08b_58_Collection_HasAll(t *testing.T) {
	safeTest(t, "Test_S08b_58_Collection_HasAll", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act & Assert
		if !col.HasAll("a", "b") {
			t.Fatal("expected true")
		}
		if col.HasAll("a", "z") {
			t.Fatal("expected false")
		}
	})
}

func Test_S08b_59_Collection_HasAll_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_59_Collection_HasAll_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act & Assert
		if col.HasAll("a") {
			t.Fatal("expected false for empty")
		}
	})
}

func Test_S08b_60_Collection_HasUsingSensitivity(t *testing.T) {
	safeTest(t, "Test_S08b_60_Collection_HasUsingSensitivity", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"Hello", "World"})

		// Act & Assert
		if col.HasUsingSensitivity("hello", true) {
			t.Fatal("expected false for case-sensitive")
		}
		if !col.HasUsingSensitivity("hello", false) {
			t.Fatal("expected true for case-insensitive")
		}
		if col.HasUsingSensitivity("missing", false) {
			t.Fatal("expected false for missing item")
		}
	})
}

// ── IsContainsPtr / IsContainsAll / IsContainsAllSlice / IsContainsAllLock ──

func Test_S08b_61_Collection_IsContainsPtr(t *testing.T) {
	safeTest(t, "Test_S08b_61_Collection_IsContainsPtr", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		s := "a"

		// Act & Assert
		if !col.IsContainsPtr(&s) {
			t.Fatal("expected true")
		}
		if col.IsContainsPtr(nil) {
			t.Fatal("expected false for nil")
		}
	})
}

func Test_S08b_62_Collection_IsContainsAllSlice(t *testing.T) {
	safeTest(t, "Test_S08b_62_Collection_IsContainsAllSlice", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act & Assert
		if !col.IsContainsAllSlice([]string{"a", "b"}) {
			t.Fatal("expected true")
		}
		if col.IsContainsAllSlice([]string{"a", "z"}) {
			t.Fatal("expected false")
		}
		if col.IsContainsAllSlice([]string{}) {
			t.Fatal("expected false for empty")
		}
	})
}

func Test_S08b_63_Collection_IsContainsAll(t *testing.T) {
	safeTest(t, "Test_S08b_63_Collection_IsContainsAll", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act & Assert
		if !col.IsContainsAll("a", "b") {
			t.Fatal("expected true")
		}
		if col.IsContainsAll(nil...) {
			t.Fatal("expected false for nil")
		}
	})
}

func Test_S08b_64_Collection_IsContainsAllLock(t *testing.T) {
	safeTest(t, "Test_S08b_64_Collection_IsContainsAllLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act & Assert
		if !col.IsContainsAllLock("a") {
			t.Fatal("expected true")
		}
		if col.IsContainsAllLock(nil...) {
			t.Fatal("expected false for nil")
		}
	})
}

func Test_S08b_65_Collection_IsContainsAllSlice_EmptyCollection(t *testing.T) {
	safeTest(t, "Test_S08b_65_Collection_IsContainsAllSlice_EmptyCollection", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act & Assert
		if col.IsContainsAllSlice([]string{"a"}) {
			t.Fatal("expected false for empty collection")
		}
	})
}

// ── GetHashsetPlusHasAll ─────────────────────────────────────

func Test_S08b_66_Collection_GetHashsetPlusHasAll(t *testing.T) {
	safeTest(t, "Test_S08b_66_Collection_GetHashsetPlusHasAll", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		hs, hasAll := col.GetHashsetPlusHasAll([]string{"a", "b"})

		// Assert
		if !hasAll {
			t.Fatal("expected true")
		}
		if hs == nil {
			t.Fatal("expected hashset")
		}
	})
}

func Test_S08b_67_Collection_GetHashsetPlusHasAll_NilItems(t *testing.T) {
	safeTest(t, "Test_S08b_67_Collection_GetHashsetPlusHasAll_NilItems", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		_, hasAll := col.GetHashsetPlusHasAll(nil)

		// Assert
		if hasAll {
			t.Fatal("expected false for nil items")
		}
	})
}

func Test_S08b_68_Collection_GetHashsetPlusHasAll_EmptyCollection(t *testing.T) {
	safeTest(t, "Test_S08b_68_Collection_GetHashsetPlusHasAll_EmptyCollection", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		_, hasAll := col.GetHashsetPlusHasAll([]string{"a"})

		// Assert
		if hasAll {
			t.Fatal("expected false for empty collection")
		}
	})
}

// ── SortedListAsc / SortedAsc / SortedAscLock / SortedListDsc ──

func Test_S08b_69_Collection_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_S08b_69_Collection_SortedListAsc", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"c", "a", "b"})

		// Act
		sorted := col.SortedListAsc()

		// Assert
		if sorted[0] != "a" || sorted[2] != "c" {
			t.Fatal("expected sorted ascending")
		}
	})
}

func Test_S08b_70_Collection_SortedListAsc_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_70_Collection_SortedListAsc_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		sorted := col.SortedListAsc()

		// Assert
		if len(sorted) != 0 {
			t.Fatal("expected empty")
		}
	})
}

func Test_S08b_71_Collection_SortedAsc(t *testing.T) {
	safeTest(t, "Test_S08b_71_Collection_SortedAsc", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"b", "a"})

		// Act
		col.SortedAsc()

		// Assert
		if col.First() != "a" {
			t.Fatal("expected 'a' first")
		}
	})
}

func Test_S08b_72_Collection_SortedAsc_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_72_Collection_SortedAsc_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		result := col.SortedAsc()

		// Assert
		if result != col {
			t.Fatal("expected same pointer for empty")
		}
	})
}

func Test_S08b_73_Collection_SortedAscLock(t *testing.T) {
	safeTest(t, "Test_S08b_73_Collection_SortedAscLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"b", "a"})

		// Act
		col.SortedAscLock()

		// Assert
		if col.First() != "a" {
			t.Fatal("expected 'a' first")
		}
	})
}

func Test_S08b_74_Collection_SortedAscLock_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_74_Collection_SortedAscLock_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		result := col.SortedAscLock()

		// Assert
		if result != col {
			t.Fatal("expected same pointer for empty")
		}
	})
}

func Test_S08b_75_Collection_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_S08b_75_Collection_SortedListDsc", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "c", "b"})

		// Act
		sorted := col.SortedListDsc()

		// Assert
		if sorted[0] != "c" || sorted[2] != "a" {
			t.Fatal("expected sorted descending")
		}
	})
}

// ── New ──────────────────────────────────────────────────────

func Test_S08b_76_Collection_New(t *testing.T) {
	safeTest(t, "Test_S08b_76_Collection_New", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		newCol := col.New("a", "b")

		// Assert
		if newCol.Length() != 2 {
			t.Fatalf("expected 2, got %d", newCol.Length())
		}
	})
}

func Test_S08b_77_Collection_New_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_77_Collection_New_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		newCol := col.New()

		// Assert
		if newCol.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

// ── AddNonEmptyStrings / AddNonEmptyStringsSlice ─────────────

func Test_S08b_78_Collection_AddNonEmptyStrings(t *testing.T) {
	safeTest(t, "Test_S08b_78_Collection_AddNonEmptyStrings", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddNonEmptyStrings("a", "", "b")

		// Assert — AddNonEmptyStrings filters empty strings, so only "a" and "b" are added
		if col.Length() != 2 {
			t.Fatalf("expected 2, got %d", col.Length())
		}
	})
}

func Test_S08b_79_Collection_AddNonEmptyStrings_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_79_Collection_AddNonEmptyStrings_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddNonEmptyStrings()

		// Assert
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

// ── AddFuncResult ────────────────────────────────────────────

func Test_S08b_80_Collection_AddFuncResult(t *testing.T) {
	safeTest(t, "Test_S08b_80_Collection_AddFuncResult", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddFuncResult(
			func() string { return "a" },
			func() string { return "b" },
		)

		// Assert
		if col.Length() != 2 {
			t.Fatalf("expected 2, got %d", col.Length())
		}
	})
}

func Test_S08b_81_Collection_AddFuncResult_Nil(t *testing.T) {
	safeTest(t, "Test_S08b_81_Collection_AddFuncResult_Nil", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddFuncResult(nil...)

		// Assert
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

// ── AddStringsByFuncChecking ─────────────────────────────────

func Test_S08b_82_Collection_AddStringsByFuncChecking(t *testing.T) {
	safeTest(t, "Test_S08b_82_Collection_AddStringsByFuncChecking", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.AddStringsByFuncChecking(
			[]string{"hello", "", "world"},
			func(line string) bool { return line != "" },
		)

		// Assert
		if col.Length() != 2 {
			t.Fatalf("expected 2, got %d", col.Length())
		}
	})
}

// ── ExpandSlicePlusAdd ───────────────────────────────────────

func Test_S08b_83_Collection_ExpandSlicePlusAdd(t *testing.T) {
	safeTest(t, "Test_S08b_83_Collection_ExpandSlicePlusAdd", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.ExpandSlicePlusAdd(
			[]string{"a,b", "c,d"},
			func(line string) []string { return strings.Split(line, ",") },
		)

		// Assert
		if col.Length() != 4 {
			t.Fatalf("expected 4, got %d", col.Length())
		}
	})
}

// ── MergeSlicesOfSlice ───────────────────────────────────────

func Test_S08b_84_Collection_MergeSlicesOfSlice(t *testing.T) {
	safeTest(t, "Test_S08b_84_Collection_MergeSlicesOfSlice", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		col.MergeSlicesOfSlice([]string{"a"}, []string{"b", "c"})

		// Assert
		if col.Length() != 3 {
			t.Fatalf("expected 3, got %d", col.Length())
		}
	})
}

// ── GetAllExcept / GetAllExceptCollection ────────────────────

func Test_S08b_85_Collection_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_S08b_85_Collection_GetAllExceptCollection", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		except := corestr.New.Collection.Strings([]string{"b"})

		// Act
		result := col.GetAllExceptCollection(except)

		// Assert
		if len(result) != 2 {
			t.Fatalf("expected 2, got %d", len(result))
		}
	})
}

func Test_S08b_86_Collection_GetAllExceptCollection_NilExcept(t *testing.T) {
	safeTest(t, "Test_S08b_86_Collection_GetAllExceptCollection_NilExcept", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		result := col.GetAllExceptCollection(nil)

		// Assert
		if len(result) != 2 {
			t.Fatalf("expected 2, got %d", len(result))
		}
	})
}

func Test_S08b_87_Collection_GetAllExceptCollection_EmptyExcept(t *testing.T) {
	safeTest(t, "Test_S08b_87_Collection_GetAllExceptCollection_EmptyExcept", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})
		except := corestr.Empty.Collection()

		// Act
		result := col.GetAllExceptCollection(except)

		// Assert
		if len(result) != 1 {
			t.Fatalf("expected 1, got %d", len(result))
		}
	})
}

func Test_S08b_88_Collection_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_S08b_88_Collection_GetAllExcept", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		result := col.GetAllExcept([]string{"a"})

		// Assert
		if len(result) != 2 {
			t.Fatalf("expected 2, got %d", len(result))
		}
	})
}

func Test_S08b_89_Collection_GetAllExcept_NilSlice(t *testing.T) {
	safeTest(t, "Test_S08b_89_Collection_GetAllExcept_NilSlice", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		result := col.GetAllExcept(nil)

		// Assert
		if len(result) != 2 {
			t.Fatalf("expected 2, got %d", len(result))
		}
	})
}

// ── CharCollectionMap ────────────────────────────────────────

func Test_S08b_90_Collection_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_S08b_90_Collection_CharCollectionMap", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"apple", "banana", "avocado"})

		// Act
		ccm := col.CharCollectionMap()

		// Assert
		if ccm == nil {
			t.Fatal("expected non-nil")
		}
	})
}

// ── SummaryString / SummaryStringWithHeader / String ─────────

func Test_S08b_91_Collection_SummaryString(t *testing.T) {
	safeTest(t, "Test_S08b_91_Collection_SummaryString", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		s := col.SummaryString(1)

		// Assert
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S08b_92_Collection_SummaryStringWithHeader(t *testing.T) {
	safeTest(t, "Test_S08b_92_Collection_SummaryStringWithHeader", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		s := col.SummaryStringWithHeader("Header:")

		// Assert
		if !strings.HasPrefix(s, "Header:") {
			t.Fatal("expected header prefix")
		}
	})
}

func Test_S08b_93_Collection_SummaryStringWithHeader_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_93_Collection_SummaryStringWithHeader_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		s := col.SummaryStringWithHeader("Header:")

		// Assert
		if !strings.Contains(s, "No Element") {
			t.Fatal("expected No Element")
		}
	})
}

func Test_S08b_94_Collection_String(t *testing.T) {
	safeTest(t, "Test_S08b_94_Collection_String", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		s := col.String()

		// Assert
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S08b_95_Collection_String_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_95_Collection_String_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		s := col.String()

		// Assert
		if !strings.Contains(s, "No Element") {
			t.Fatal("expected No Element")
		}
	})
}

func Test_S08b_96_Collection_StringLock(t *testing.T) {
	safeTest(t, "Test_S08b_96_Collection_StringLock", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		s := col.StringLock()

		// Assert
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S08b_97_Collection_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_97_Collection_StringLock_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		s := col.StringLock()

		// Assert
		if !strings.Contains(s, "No Element") {
			t.Fatal("expected No Element")
		}
	})
}

// ── Csv / CsvOptions / CsvLines / CsvLinesOptions ───────────

func Test_S08b_98_Collection_Csv(t *testing.T) {
	safeTest(t, "Test_S08b_98_Collection_Csv", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		csv := col.Csv()

		// Assert
		if csv == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S08b_99_Collection_Csv_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_99_Collection_Csv_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		csv := col.Csv()

		// Assert
		if csv != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_S08b_100_Collection_CsvOptions(t *testing.T) {
	safeTest(t, "Test_S08b_100_Collection_CsvOptions", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		csv := col.CsvOptions(true)

		// Assert
		if csv == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S08b_101_Collection_CsvOptions_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_101_Collection_CsvOptions_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		csv := col.CsvOptions(true)

		// Assert
		if csv != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_S08b_102_Collection_CsvLines(t *testing.T) {
	safeTest(t, "Test_S08b_102_Collection_CsvLines", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		lines := col.CsvLines()

		// Assert
		if len(lines) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S08b_103_Collection_CsvLinesOptions(t *testing.T) {
	safeTest(t, "Test_S08b_103_Collection_CsvLinesOptions", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		lines := col.CsvLinesOptions(true)

		// Assert
		if len(lines) != 1 {
			t.Fatal("expected 1")
		}
	})
}

// ── AddCapacity / Resize ─────────────────────────────────────

func Test_S08b_104_Collection_AddCapacity(t *testing.T) {
	safeTest(t, "Test_S08b_104_Collection_AddCapacity", func() {
		// Arrange
		col := corestr.New.Collection.Cap(5)

		// Act
		col.AddCapacity(10)

		// Assert
		if col.Capacity() < 15 {
			t.Fatal("expected capacity >= 15")
		}
	})
}

func Test_S08b_105_Collection_AddCapacity_Nil(t *testing.T) {
	safeTest(t, "Test_S08b_105_Collection_AddCapacity_Nil", func() {
		// Arrange
		col := corestr.New.Collection.Cap(5)

		// Act
		col.AddCapacity(nil...)

		// Assert — no panic
	})
}

func Test_S08b_106_Collection_Resize(t *testing.T) {
	safeTest(t, "Test_S08b_106_Collection_Resize", func() {
		// Arrange
		col := corestr.New.Collection.Cap(5)

		// Act
		col.Resize(20)

		// Assert
		if col.Capacity() < 20 {
			t.Fatal("expected capacity >= 20")
		}
	})
}

func Test_S08b_107_Collection_Resize_SmallerThanExisting(t *testing.T) {
	safeTest(t, "Test_S08b_107_Collection_Resize_SmallerThanExisting", func() {
		// Arrange
		col := corestr.New.Collection.Cap(20)

		// Act
		col.Resize(5)

		// Assert — should not shrink
		if col.Capacity() < 20 {
			t.Fatal("expected capacity unchanged")
		}
	})
}

// ── Joins / NonEmptyJoins / NonWhitespaceJoins ───────────────

func Test_S08b_108_Collection_Joins(t *testing.T) {
	safeTest(t, "Test_S08b_108_Collection_Joins", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		s := col.Joins(",")

		// Assert
		if s != "a,b" {
			t.Fatalf("expected 'a,b', got '%s'", s)
		}
	})
}

func Test_S08b_109_Collection_Joins_WithExtra(t *testing.T) {
	safeTest(t, "Test_S08b_109_Collection_Joins_WithExtra", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		s := col.Joins(",", "b", "c")

		// Assert
		if !strings.Contains(s, "b") || !strings.Contains(s, "c") {
			t.Fatalf("expected b and c in result, got '%s'", s)
		}
	})
}

func Test_S08b_110_Collection_NonEmptyJoins(t *testing.T) {
	safeTest(t, "Test_S08b_110_Collection_NonEmptyJoins", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "", "b"})

		// Act
		s := col.NonEmptyJoins(",")

		// Assert
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S08b_111_Collection_NonWhitespaceJoins(t *testing.T) {
	safeTest(t, "Test_S08b_111_Collection_NonWhitespaceJoins", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "  ", "b"})

		// Act
		s := col.NonWhitespaceJoins(",")

		// Assert
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

// ── JSON methods ─────────────────────────────────────────────

func Test_S08b_112_Collection_JsonModel(t *testing.T) {
	safeTest(t, "Test_S08b_112_Collection_JsonModel", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		model := col.JsonModel()

		// Assert
		if len(model) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S08b_113_Collection_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_S08b_113_Collection_JsonModelAny", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		model := col.JsonModelAny()

		// Assert
		if model == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S08b_114_Collection_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_S08b_114_Collection_MarshalJSON", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		data, err := col.MarshalJSON()

		// Assert
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(data) == 0 {
			t.Fatal("expected non-empty bytes")
		}
	})
}

func Test_S08b_115_Collection_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_S08b_115_Collection_UnmarshalJSON", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		err := col.UnmarshalJSON([]byte(`["a","b"]`))

		// Assert
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S08b_116_Collection_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_S08b_116_Collection_UnmarshalJSON_Invalid", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		err := col.UnmarshalJSON([]byte(`invalid`))

		// Assert
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_S08b_117_Collection_Json(t *testing.T) {
	safeTest(t, "Test_S08b_117_Collection_Json", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		result := col.Json()

		// Assert
		if result.HasError() {
			t.Fatal("expected no error")
		}
	})
}

func Test_S08b_118_Collection_JsonPtr(t *testing.T) {
	safeTest(t, "Test_S08b_118_Collection_JsonPtr", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		result := col.JsonPtr()

		// Assert
		if result == nil {
			t.Fatal("expected non-nil")
		}
	})
}

// ── ParseInjectUsingJson / ParseInjectUsingJsonMust / JsonParseSelfInject ──

func Test_S08b_119_Collection_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_S08b_119_Collection_ParseInjectUsingJson", func() {
		// Arrange
		original := corestr.New.Collection.Strings([]string{"a", "b"})
		jsonResult := original.JsonPtr()
		target := corestr.Empty.Collection()

		// Act
		result, err := target.ParseInjectUsingJson(jsonResult)

		// Assert
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S08b_120_Collection_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_S08b_120_Collection_ParseInjectUsingJsonMust", func() {
		// Arrange
		original := corestr.New.Collection.Strings([]string{"a"})
		jsonResult := original.JsonPtr()
		target := corestr.Empty.Collection()

		// Act
		result := target.ParseInjectUsingJsonMust(jsonResult)

		// Assert
		if result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S08b_121_Collection_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_S08b_121_Collection_JsonParseSelfInject", func() {
		// Arrange
		original := corestr.New.Collection.Strings([]string{"x"})
		jsonResult := original.JsonPtr()
		target := corestr.Empty.Collection()

		// Act
		err := target.JsonParseSelfInject(jsonResult)

		// Assert
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

// ── Clear / Dispose ──────────────────────────────────────────

func Test_S08b_122_Collection_Clear(t *testing.T) {
	safeTest(t, "Test_S08b_122_Collection_Clear", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		col.Clear()

		// Assert
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S08b_123_Collection_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_S08b_123_Collection_Clear_Nil", func() {
		// Arrange
		var col *corestr.Collection

		// Act
		result := col.Clear()

		// Assert
		if result != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_S08b_124_Collection_Dispose(t *testing.T) {
	safeTest(t, "Test_S08b_124_Collection_Dispose", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		col.Dispose()

		// Assert
		if col.Length() != 0 {
			t.Fatal("expected 0 after dispose")
		}
	})
}

func Test_S08b_125_Collection_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_S08b_125_Collection_Dispose_Nil", func() {
		// Arrange
		var col *corestr.Collection

		// Act — should not panic
		col.Dispose()
	})
}

// ── AsJsonMarshaller / AsJsonContractsBinder ─────────────────

func Test_S08b_126_Collection_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_S08b_126_Collection_AsJsonMarshaller", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		m := col.AsJsonMarshaller()

		// Assert
		if m == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S08b_127_Collection_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_S08b_127_Collection_AsJsonContractsBinder", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		b := col.AsJsonContractsBinder()

		// Assert
		if b == nil {
			t.Fatal("expected non-nil")
		}
	})
}

// ── Serialize / Deserialize ──────────────────────────────────

func Test_S08b_128_Collection_Serialize(t *testing.T) {
	safeTest(t, "Test_S08b_128_Collection_Serialize", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		data, err := col.Serialize()

		// Assert
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(data) == 0 {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S08b_129_Collection_Deserialize(t *testing.T) {
	safeTest(t, "Test_S08b_129_Collection_Deserialize", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})
		var target []string

		// Act
		err := col.Deserialize(&target)

		// Assert
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(target) != 1 {
			t.Fatal("expected 1")
		}
	})
}

// ── Join / JoinLine ──────────────────────────────────────────

func Test_S08b_130_Collection_Join(t *testing.T) {
	safeTest(t, "Test_S08b_130_Collection_Join", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		s := col.Join(",")

		// Assert
		if s != "a,b" {
			t.Fatalf("expected 'a,b', got '%s'", s)
		}
	})
}

func Test_S08b_131_Collection_Join_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_131_Collection_Join_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		s := col.Join(",")

		// Assert
		if s != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_S08b_132_Collection_JoinLine(t *testing.T) {
	safeTest(t, "Test_S08b_132_Collection_JoinLine", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		s := col.JoinLine()

		// Assert
		if !strings.Contains(s, "\n") {
			t.Fatal("expected newline")
		}
	})
}

func Test_S08b_133_Collection_JoinLine_Empty(t *testing.T) {
	safeTest(t, "Test_S08b_133_Collection_JoinLine_Empty", func() {
		// Arrange
		col := corestr.Empty.Collection()

		// Act
		s := col.JoinLine()

		// Assert
		if s != "" {
			t.Fatal("expected empty")
		}
	})
}
