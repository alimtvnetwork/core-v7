package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ==========================================
// Constructors
// ==========================================

func Test_StrHashset_Empty(t *testing.T) {
	safeTest(t, "Test_StrHashset_Empty", func() {
		hs := corestr.New.Hashset.Empty()
		if !hs.IsEmpty() {
			t.Error("Empty hashset should be empty")
		}
		if hs.Length() != 0 {
			t.Errorf("Empty hashset length: expected 0, got %d", hs.Length())
		}
	})
}

func Test_StrHashset_Cap(t *testing.T) {
	safeTest(t, "Test_StrHashset_Cap", func() {
		hs := corestr.New.Hashset.Cap(50)
		if !hs.IsEmpty() {
			t.Error("Cap hashset should be empty initially")
		}
	})
}

func Test_StrHashset_Strings(t *testing.T) {
	safeTest(t, "Test_StrHashset_Strings", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c", "a"})
		if hs.Length() != 3 {
			t.Errorf("Strings with duplicates: expected 3, got %d", hs.Length())
		}
	})
}

func Test_StrHashset_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_StrHashset_Strings_Empty", func() {
		hs := corestr.New.Hashset.Strings([]string{})
		if !hs.IsEmpty() {
			t.Error("Strings from empty slice should be empty")
		}
	})
}

func Test_StrHashset_StringsSpreadItems(t *testing.T) {
	safeTest(t, "Test_StrHashset_StringsSpreadItems", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("x", "y", "z")
		if hs.Length() != 3 {
			t.Errorf("StringsSpreadItems: expected 3, got %d", hs.Length())
		}
	})
}

func Test_StrHashset_UsingMap(t *testing.T) {
	safeTest(t, "Test_StrHashset_UsingMap", func() {
		m := map[string]bool{"a": true, "b": true}
		hs := corestr.New.Hashset.UsingMap(m)
		if hs.Length() != 2 {
			t.Errorf("UsingMap: expected 2, got %d", hs.Length())
		}
	})
}

// ==========================================
// Add / AddBool — caching behavior (bug 42)
// ==========================================

func Test_StrHashset_Add_SetsHasMapUpdated(t *testing.T) {
	safeTest(t, "Test_StrHashset_Add_SetsHasMapUpdated", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("hello")
		if !hs.Has("hello") {
			t.Error("Add should insert item")
		}
		if hs.Length() != 1 {
			t.Errorf("After Add: expected 1, got %d", hs.Length())
		}
	})
}

func Test_StrHashset_Add_Duplicate_NoIncrease(t *testing.T) {
	safeTest(t, "Test_StrHashset_Add_Duplicate_NoIncrease", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a").Add("a").Add("a")
		if hs.Length() != 1 {
			t.Errorf("Add duplicate: expected 1, got %d", hs.Length())
		}
	})
}

func Test_StrHashset_AddBool_FirstAdd_ReturnsFalse(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddBool_FirstAdd_ReturnsFalse", func() {
		hs := corestr.New.Hashset.Empty()
		existed := hs.AddBool("x")
		if existed {
			t.Error("AddBool first add should return false (did not exist)")
		}
		if hs.Length() != 1 {
			t.Errorf("After AddBool: expected 1, got %d", hs.Length())
		}
	})
}

func Test_StrHashset_AddBool_SecondAdd_ReturnsTrue(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddBool_SecondAdd_ReturnsTrue", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddBool("x")
		existed := hs.AddBool("x")
		if !existed {
			t.Error("AddBool second add should return true (already existed)")
		}
		if hs.Length() != 1 {
			t.Errorf("AddBool duplicate: expected 1, got %d", hs.Length())
		}
	})
}

func Test_StrHashset_AddBool_CacheInvalidation(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddBool_CacheInvalidation", func() {
		// Bug 42 context: hasMapUpdated must be set on AddBool for new items
		hs := corestr.New.Hashset.Empty()
		// Force cache by calling List
		_ = hs.List()
		hs.AddBool("new-item")
		list := hs.List()
		found := false
		for _, v := range list {
			if v == "new-item" {
				found = true
			}
		}
		if !found {
			t.Error("AddBool should invalidate cache so List reflects new item")
		}
	})
}

func Test_StrHashset_AddNonEmpty_SkipsEmpty(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddNonEmpty_SkipsEmpty", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddNonEmpty("")
		if hs.Length() != 0 {
			t.Error("AddNonEmpty should skip empty string")
		}
		hs.AddNonEmpty("valid")
		if hs.Length() != 1 {
			t.Errorf("AddNonEmpty: expected 1, got %d", hs.Length())
		}
	})
}

func Test_StrHashset_AddNonEmptyWhitespace_SkipsWhitespace(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddNonEmptyWhitespace_SkipsWhitespace", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddNonEmptyWhitespace("   ")
		if hs.Length() != 0 {
			t.Error("AddNonEmptyWhitespace should skip whitespace-only")
		}
		hs.AddNonEmptyWhitespace("valid")
		if hs.Length() != 1 {
			t.Errorf("expected 1, got %d", hs.Length())
		}
	})
}

// ==========================================
// Adds / AddStrings / AddIf / AddIfMany
// ==========================================

func Test_StrHashset_Adds_Variadic(t *testing.T) {
	safeTest(t, "Test_StrHashset_Adds_Variadic", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b", "c")
		if hs.Length() != 3 {
			t.Errorf("Adds: expected 3, got %d", hs.Length())
		}
	})
}

func Test_StrHashset_AddStrings(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddStrings", func() {
		hs := corestr.New.Hashset.Cap(10)
		hs.AddStrings([]string{"x", "y", "x"})
		if hs.Length() != 2 {
			t.Errorf("AddStrings with dup: expected 2, got %d", hs.Length())
		}
	})
}

func Test_StrHashset_AddIf_True(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddIf_True", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddIf(true, "yes")
		if !hs.Has("yes") {
			t.Error("AddIf(true) should add item")
		}
	})
}

func Test_StrHashset_AddIf_False(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddIf_False", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddIf(false, "no")
		if hs.Has("no") {
			t.Error("AddIf(false) should not add item")
		}
	})
}

func Test_StrHashset_AddIfMany_True(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddIfMany_True", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddIfMany(true, "a", "b")
		if hs.Length() != 2 {
			t.Errorf("AddIfMany(true): expected 2, got %d", hs.Length())
		}
	})
}

func Test_StrHashset_AddIfMany_False(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddIfMany_False", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddIfMany(false, "a", "b")
		if hs.Length() != 0 {
			t.Errorf("AddIfMany(false): expected 0, got %d", hs.Length())
		}
	})
}

// ==========================================
// AddHashsetItems / AddItemsMap
// ==========================================

func Test_StrHashset_AddHashsetItems_Merge(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddHashsetItems_Merge", func() {
		hs1 := corestr.New.Hashset.StringsSpreadItems("a", "b")
		hs2 := corestr.New.Hashset.StringsSpreadItems("b", "c")
		hs1.AddHashsetItems(hs2)
		if hs1.Length() != 3 {
			t.Errorf("Merge: expected 3, got %d", hs1.Length())
		}
	})
}

func Test_StrHashset_AddHashsetItems_Nil(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddHashsetItems_Nil", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		result := hs.AddHashsetItems(nil)
		if result != hs || hs.Length() != 1 {
			t.Error("AddHashsetItems(nil) should be no-op")
		}
	})
}

func Test_StrHashset_AddItemsMap_OnlyTrueValues(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddItemsMap_OnlyTrueValues", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddItemsMap(map[string]bool{"yes": true, "no": false, "also": true})
		if hs.Length() != 2 {
			t.Errorf("AddItemsMap: expected 2 (only true), got %d", hs.Length())
		}
		if hs.Has("no") {
			t.Error("AddItemsMap should skip false-valued entries")
		}
	})
}

// ==========================================
// Has / HasAll / HasAny
// ==========================================

func Test_StrHashset_Has_Existing(t *testing.T) {
	safeTest(t, "Test_StrHashset_Has_Existing", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("alpha", "beta")
		if !hs.Has("alpha") {
			t.Error("Has should find existing item")
		}
	})
}

func Test_StrHashset_Has_Missing(t *testing.T) {
	safeTest(t, "Test_StrHashset_Has_Missing", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("alpha", "beta")
		if hs.Has("gamma") {
			t.Error("Has should not find missing item")
		}
	})
}

func Test_StrHashset_HasAll_AllPresent(t *testing.T) {
	safeTest(t, "Test_StrHashset_HasAll_AllPresent", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		if !hs.HasAll("a", "c") {
			t.Error("HasAll should return true when all present")
		}
	})
}

func Test_StrHashset_HasAll_OneMissing(t *testing.T) {
	safeTest(t, "Test_StrHashset_HasAll_OneMissing", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		if hs.HasAll("a", "z") {
			t.Error("HasAll should return false when one missing")
		}
	})
}

func Test_StrHashset_HasAll_EmptyArgs(t *testing.T) {
	safeTest(t, "Test_StrHashset_HasAll_EmptyArgs", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		if !hs.HasAll() {
			t.Error("HasAll with no args should return true")
		}
	})
}

func Test_StrHashset_HasAny_OnePresent(t *testing.T) {
	safeTest(t, "Test_StrHashset_HasAny_OnePresent", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		if !hs.HasAny("z", "b", "y") {
			t.Error("HasAny should return true when at least one present")
		}
	})
}

func Test_StrHashset_HasAny_NonePresent(t *testing.T) {
	safeTest(t, "Test_StrHashset_HasAny_NonePresent", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		if hs.HasAny("x", "y") {
			t.Error("HasAny should return false when none present")
		}
	})
}

func Test_StrHashset_HasAny_EmptyArgs(t *testing.T) {
	safeTest(t, "Test_StrHashset_HasAny_EmptyArgs", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		if hs.HasAny() {
			t.Error("HasAny with no args should return false")
		}
	})
}

func Test_StrHashset_IsMissing(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsMissing", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		if !hs.IsMissing("z") {
			t.Error("IsMissing should return true for absent key")
		}
		if hs.IsMissing("a") {
			t.Error("IsMissing should return false for present key")
		}
	})
}

func Test_StrHashset_IsAllMissing(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsAllMissing", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		if !hs.IsAllMissing("x", "y") {
			t.Error("IsAllMissing should return true when all missing")
		}
		if hs.IsAllMissing("x", "a") {
			t.Error("IsAllMissing should return false when any present")
		}
	})
}

// ==========================================
// Remove / SafeRemove
// ==========================================

func Test_StrHashset_Remove(t *testing.T) {
	safeTest(t, "Test_StrHashset_Remove", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		hs.Remove("b")
		if hs.Has("b") {
			t.Error("Remove should delete item")
		}
		if hs.Length() != 2 {
			t.Errorf("After remove: expected 2, got %d", hs.Length())
		}
	})
}

func Test_StrHashset_SafeRemove_Existing(t *testing.T) {
	safeTest(t, "Test_StrHashset_SafeRemove_Existing", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		hs.SafeRemove("a")
		if hs.Has("a") {
			t.Error("SafeRemove should delete existing item")
		}
	})
}

func Test_StrHashset_SafeRemove_Missing(t *testing.T) {
	safeTest(t, "Test_StrHashset_SafeRemove_Missing", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs.SafeRemove("z") // should not panic
		if hs.Length() != 1 {
			t.Errorf("SafeRemove missing: expected 1, got %d", hs.Length())
		}
	})
}

// ==========================================
// Resize
// ==========================================

func Test_StrHashset_Resize_LargerPreservesItems(t *testing.T) {
	safeTest(t, "Test_StrHashset_Resize_LargerPreservesItems", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		hs.Resize(100)
		if hs.Length() != 2 {
			t.Errorf("Resize should preserve items: expected 2, got %d", hs.Length())
		}
		if !hs.Has("a") || !hs.Has("b") {
			t.Error("Resize should preserve all items")
		}
	})
}

func Test_StrHashset_Resize_SmallerIsNoOp(t *testing.T) {
	safeTest(t, "Test_StrHashset_Resize_SmallerIsNoOp", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		hs.Resize(1)
		if hs.Length() != 3 {
			t.Errorf("Resize smaller: expected 3, got %d", hs.Length())
		}
	})
}

func Test_StrHashset_AddCapacities(t *testing.T) {
	safeTest(t, "Test_StrHashset_AddCapacities", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs.AddCapacities(10, 20)
		if hs.Length() != 1 {
			t.Errorf("AddCapacities should preserve items: expected 1, got %d", hs.Length())
		}
	})
}

// ==========================================
// IsEquals
// ==========================================

func Test_StrHashset_IsEquals_BothNil(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsEquals_BothNil", func() {
		var a, b *corestr.Hashset
		if !a.IsEquals(b) {
			t.Error("Two nil hashsets should be equal")
		}
	})
}

func Test_StrHashset_IsEquals_OneNil(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsEquals_OneNil", func() {
		hs := corestr.New.Hashset.Empty()
		var nilHs *corestr.Hashset
		if hs.IsEquals(nilHs) {
			t.Error("Non-nil vs nil should not be equal")
		}
	})
}

func Test_StrHashset_IsEquals_SamePointer(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsEquals_SamePointer", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		if !hs.IsEquals(hs) {
			t.Error("Same pointer should be equal")
		}
	})
}

func Test_StrHashset_IsEquals_SameContent(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsEquals_SameContent", func() {
		a := corestr.New.Hashset.StringsSpreadItems("x", "y")
		b := corestr.New.Hashset.StringsSpreadItems("y", "x")
		if !a.IsEquals(b) {
			t.Error("Same content should be equal")
		}
	})
}

func Test_StrHashset_IsEquals_DifferentContent(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsEquals_DifferentContent", func() {
		a := corestr.New.Hashset.StringsSpreadItems("a", "b")
		b := corestr.New.Hashset.StringsSpreadItems("a", "c")
		if a.IsEquals(b) {
			t.Error("Different content should not be equal")
		}
	})
}

func Test_StrHashset_IsEquals_DifferentLength(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsEquals_DifferentLength", func() {
		a := corestr.New.Hashset.StringsSpreadItems("a")
		b := corestr.New.Hashset.StringsSpreadItems("a", "b")
		if a.IsEquals(b) {
			t.Error("Different length should not be equal")
		}
	})
}

func Test_StrHashset_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_StrHashset_IsEquals_BothEmpty", func() {
		a := corestr.New.Hashset.Empty()
		b := corestr.New.Hashset.Empty()
		if !a.IsEquals(b) {
			t.Error("Two empty hashsets should be equal")
		}
	})
}

// ==========================================
// List caching (bug 42 context)
// ==========================================

func Test_StrHashset_List_CacheInvalidatedAfterAdd(t *testing.T) {
	safeTest(t, "Test_StrHashset_List_CacheInvalidatedAfterAdd", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		list1 := hs.List()
		if len(list1) != 1 {
			t.Errorf("Initial list: expected 1, got %d", len(list1))
		}
		hs.Add("b")
		list2 := hs.List()
		if len(list2) != 2 {
			t.Errorf("After Add, list should reflect new item: expected 2, got %d", len(list2))
		}
	})
}

func Test_StrHashset_List_CacheInvalidatedAfterRemove(t *testing.T) {
	safeTest(t, "Test_StrHashset_List_CacheInvalidatedAfterRemove", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		_ = hs.List() // populate cache
		hs.Remove("a")
		list := hs.List()
		if len(list) != 1 {
			t.Errorf("After Remove, list should reflect removal: expected 1, got %d", len(list))
		}
	})
}

func Test_StrHashset_List_CacheInvalidatedAfterAdds(t *testing.T) {
	safeTest(t, "Test_StrHashset_List_CacheInvalidatedAfterAdds", func() {
		hs := corestr.New.Hashset.Empty()
		_ = hs.List() // populate cache
		hs.Adds("x", "y")
		list := hs.List()
		if len(list) != 2 {
			t.Errorf("After Adds, list should reflect new items: expected 2, got %d", len(list))
		}
	})
}

// ==========================================
// Clear / Dispose
// ==========================================

func Test_StrHashset_Clear(t *testing.T) {
	safeTest(t, "Test_StrHashset_Clear", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		hs.Clear()
		if !hs.IsEmpty() {
			t.Error("Clear should make hashset empty")
		}
	})
}

// ==========================================
// Nil receiver guards (CaseNilSafe pattern)
// ==========================================

func Test_StrHashset_NilReceiver(t *testing.T) {
	safeTest(t, "Test_StrHashset_NilReceiver", func() {
		for caseIndex, tc := range hashsetNilReceiverTestCases {
			// Arrange (implicit — nil receiver)

			// Act & Assert
			tc.ShouldBeSafe(t, caseIndex)
		}
	})
}

// ==========================================
// String-specific methods
// ==========================================

func Test_StrHashset_ToLowerSet(t *testing.T) {
	safeTest(t, "Test_StrHashset_ToLowerSet", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("ABC", "Def")
		lower := hs.ToLowerSet()
		if !lower.Has("abc") || !lower.Has("def") {
			t.Error("ToLowerSet should lowercase all keys")
		}
		if lower.Has("ABC") {
			t.Error("ToLowerSet should not retain original case")
		}
	})
}

func Test_StrHashset_GetAllExceptHashset(t *testing.T) {
	safeTest(t, "Test_StrHashset_GetAllExceptHashset", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		except := corestr.New.Hashset.StringsSpreadItems("b")
		result := hs.GetAllExceptHashset(except)
		if len(result) != 2 {
			t.Errorf("GetAllExceptHashset: expected 2, got %d", len(result))
		}
	})
}

func Test_StrHashset_GetAllExceptHashset_NilExcept(t *testing.T) {
	safeTest(t, "Test_StrHashset_GetAllExceptHashset_NilExcept", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		result := hs.GetAllExceptHashset(nil)
		if len(result) != 2 {
			t.Errorf("GetAllExceptHashset(nil): expected 2, got %d", len(result))
		}
	})
}

func Test_StrHashset_Collection_Conversion(t *testing.T) {
	safeTest(t, "Test_StrHashset_Collection_Conversion", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		col := hs.Collection()
		if col.Length() != 2 {
			t.Errorf("Collection: expected 2, got %d", col.Length())
		}
	})
}

func Test_StrHashset_OrderedList(t *testing.T) {
	safeTest(t, "Test_StrHashset_OrderedList", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("c", "a", "b")
		list := hs.OrderedList()
		if len(list) != 3 {
			t.Errorf("OrderedList: expected 3, got %d", len(list))
		}
		if list[0] != "a" || list[1] != "b" || list[2] != "c" {
			t.Errorf("OrderedList should be sorted asc, got %v", list)
		}
	})
}

func Test_StrHashset_JoinSorted(t *testing.T) {
	safeTest(t, "Test_StrHashset_JoinSorted", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("c", "a", "b")
		result := hs.JoinSorted(",")
		if result != "a,b,c" {
			t.Errorf("JoinSorted: expected 'a,b,c', got '%s'", result)
		}
	})
}
