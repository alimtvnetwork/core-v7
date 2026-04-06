package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ==========================================
// Constructors
// ==========================================

func Test_StrHashmap_Empty(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Empty", func() {
		hm := corestr.New.Hashmap.Empty()
		if !hm.IsEmpty() {
			t.Error("Empty hashmap should be empty")
		}
		if hm.Length() != 0 {
			t.Errorf("Empty hashmap length: expected 0, got %d", hm.Length())
		}
	})
}

func Test_StrHashmap_Cap(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Cap", func() {
		hm := corestr.New.Hashmap.Cap(50)
		if !hm.IsEmpty() {
			t.Error("Cap hashmap should be empty initially")
		}
	})
}

func Test_StrHashmap_UsingMap(t *testing.T) {
	safeTest(t, "Test_StrHashmap_UsingMap", func() {
		m := map[string]string{"k1": "v1", "k2": "v2"}
		hm := corestr.New.Hashmap.UsingMap(m)
		if hm.Length() != 2 {
			t.Errorf("UsingMap: expected 2, got %d", hm.Length())
		}
	})
}

func Test_StrHashmap_KeyValuesStrings(t *testing.T) {
	safeTest(t, "Test_StrHashmap_KeyValuesStrings", func() {
		hm := corestr.New.Hashmap.KeyValuesStrings(
			[]string{"a", "b"},
			[]string{"1", "2"},
		)
		if hm.Length() != 2 {
			t.Errorf("KeyValuesStrings: expected 2, got %d", hm.Length())
		}
		val, found := hm.Get("a")
		if !found || val != "1" {
			t.Errorf("Get('a'): expected '1', got '%s' (found=%v)", val, found)
		}
	})
}

func Test_StrHashmap_KeyValuesStrings_EmptyKeys(t *testing.T) {
	safeTest(t, "Test_StrHashmap_KeyValuesStrings_EmptyKeys", func() {
		hm := corestr.New.Hashmap.KeyValuesStrings([]string{}, []string{})
		if !hm.IsEmpty() {
			t.Error("KeyValuesStrings with empty keys should be empty")
		}
	})
}

// ==========================================
// Set / AddOrUpdate
// ==========================================

func Test_StrHashmap_Set_NewKey(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Set_NewKey", func() {
		hm := corestr.New.Hashmap.Empty()
		isNew := hm.Set("key", "val")
		if !isNew {
			t.Error("Set on new key should return true")
		}
		if hm.Length() != 1 {
			t.Errorf("After Set: expected 1, got %d", hm.Length())
		}
	})
}

func Test_StrHashmap_Set_ExistingKey(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Set_ExistingKey", func() {
		hm := corestr.New.Hashmap.Empty()
		hm.Set("key", "val1")
		isNew := hm.Set("key", "val2")
		if isNew {
			t.Error("Set on existing key should return false")
		}
		val, _ := hm.Get("key")
		if val != "val2" {
			t.Errorf("Set should overwrite: expected 'val2', got '%s'", val)
		}
	})
}

func Test_StrHashmap_AddOrUpdate_NewKey(t *testing.T) {
	safeTest(t, "Test_StrHashmap_AddOrUpdate_NewKey", func() {
		hm := corestr.New.Hashmap.Empty()
		isNew := hm.AddOrUpdate("k", "v")
		if !isNew {
			t.Error("AddOrUpdate on new key should return true")
		}
	})
}

func Test_StrHashmap_AddOrUpdate_ExistingKey(t *testing.T) {
	safeTest(t, "Test_StrHashmap_AddOrUpdate_ExistingKey", func() {
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v1")
		isNew := hm.AddOrUpdate("k", "v2")
		if isNew {
			t.Error("AddOrUpdate on existing key should return false")
		}
	})
}

func Test_StrHashmap_SetTrim(t *testing.T) {
	safeTest(t, "Test_StrHashmap_SetTrim", func() {
		hm := corestr.New.Hashmap.Empty()
		hm.SetTrim("  key  ", "  val  ")
		val, found := hm.Get("key")
		if !found || val != "val" {
			t.Errorf("SetTrim: expected trimmed key/val, got '%s' (found=%v)", val, found)
		}
	})
}

func Test_StrHashmap_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_StrHashmap_AddOrUpdateHashmap", func() {
		hm1 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		hm2 := corestr.New.Hashmap.UsingMap(map[string]string{"b": "2", "a": "override"})
		hm1.AddOrUpdateHashmap(hm2)
		if hm1.Length() != 2 {
			t.Errorf("After merge: expected 2, got %d", hm1.Length())
		}
		val, _ := hm1.Get("a")
		if val != "override" {
			t.Errorf("Merge should overwrite: expected 'override', got '%s'", val)
		}
	})
}

func Test_StrHashmap_AddOrUpdateHashmap_Nil(t *testing.T) {
	safeTest(t, "Test_StrHashmap_AddOrUpdateHashmap_Nil", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		result := hm.AddOrUpdateHashmap(nil)
		if result != hm || hm.Length() != 1 {
			t.Error("AddOrUpdateHashmap(nil) should be no-op")
		}
	})
}

func Test_StrHashmap_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_StrHashmap_AddOrUpdateMap", func() {
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdateMap(map[string]string{"x": "1", "y": "2"})
		if hm.Length() != 2 {
			t.Errorf("AddOrUpdateMap: expected 2, got %d", hm.Length())
		}
	})
}

func Test_StrHashmap_AddOrUpdateMap_Empty(t *testing.T) {
	safeTest(t, "Test_StrHashmap_AddOrUpdateMap_Empty", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		hm.AddOrUpdateMap(map[string]string{})
		if hm.Length() != 1 {
			t.Errorf("AddOrUpdateMap empty: expected 1, got %d", hm.Length())
		}
	})
}

// ==========================================
// Has / Contains / HasAll / HasAny
// ==========================================

func Test_StrHashmap_Has_Existing(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Has_Existing", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		if !hm.Has("k") {
			t.Error("Has should find existing key")
		}
	})
}

func Test_StrHashmap_Has_Missing(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Has_Missing", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		if hm.Has("missing") {
			t.Error("Has should not find missing key")
		}
	})
}

func Test_StrHashmap_Contains_AliasForHas(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Contains_AliasForHas", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		if hm.Contains("k") != hm.Has("k") {
			t.Error("Contains should match Has")
		}
	})
}

func Test_StrHashmap_IsKeyMissing(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsKeyMissing", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		if hm.IsKeyMissing("k") {
			t.Error("IsKeyMissing should return false for existing key")
		}
		if !hm.IsKeyMissing("z") {
			t.Error("IsKeyMissing should return true for missing key")
		}
	})
}

func Test_StrHashmap_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_StrHashmap_HasAllStrings", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2", "c": "3"})
		if !hm.HasAllStrings("a", "c") {
			t.Error("HasAllStrings should return true when all present")
		}
		if hm.HasAllStrings("a", "z") {
			t.Error("HasAllStrings should return false when one missing")
		}
	})
}

func Test_StrHashmap_HasAll(t *testing.T) {
	safeTest(t, "Test_StrHashmap_HasAll", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		if !hm.HasAll("a", "b") {
			t.Error("HasAll should return true when all present")
		}
		if hm.HasAll("a", "x") {
			t.Error("HasAll should return false when one missing")
		}
	})
}

func Test_StrHashmap_HasAny_OnePresent(t *testing.T) {
	safeTest(t, "Test_StrHashmap_HasAny_OnePresent", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		if !hm.HasAny("z", "a") {
			t.Error("HasAny should return true when at least one present")
		}
	})
}

func Test_StrHashmap_HasAny_NonePresent(t *testing.T) {
	safeTest(t, "Test_StrHashmap_HasAny_NonePresent", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		if hm.HasAny("x", "y") {
			t.Error("HasAny should return false when none present")
		}
	})
}

// ==========================================
// Get
// ==========================================

func Test_StrHashmap_Get_Existing(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Get_Existing", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		val, found := hm.Get("k")
		if !found {
			t.Error("Get should find existing key")
		}
		if val != "v" {
			t.Errorf("Get: expected 'v', got '%s'", val)
		}
	})
}

func Test_StrHashmap_Get_Missing(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Get_Missing", func() {
		hm := corestr.New.Hashmap.Empty()
		val, found := hm.Get("missing")
		if found {
			t.Error("Get should return false for missing key")
		}
		if val != "" {
			t.Errorf("Get missing: expected empty string, got '%s'", val)
		}
	})
}

// ==========================================
// Remove
// ==========================================

func Test_StrHashmap_Remove(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Remove", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		hm.Remove("a")
		if hm.Has("a") {
			t.Error("Remove should delete key")
		}
		if hm.Length() != 1 {
			t.Errorf("After remove: expected 1, got %d", hm.Length())
		}
	})
}

func Test_StrHashmap_Remove_Missing_NoEffect(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Remove_Missing_NoEffect", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		hm.Remove("z")
		if hm.Length() != 1 {
			t.Errorf("Remove missing: expected 1, got %d", hm.Length())
		}
	})
}

// ==========================================
// Clear / Dispose
// ==========================================

func Test_StrHashmap_Clear(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Clear", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		hm.Clear()
		if !hm.IsEmpty() {
			t.Error("Clear should make hashmap empty")
		}
	})
}

func Test_StrHashmap_Clear_NilReceiver(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Clear_NilReceiver", func() {
		var hm *corestr.Hashmap
		result := hm.Clear()
		if result != nil {
			t.Error("Clear on nil should return nil")
		}
	})
}

func Test_StrHashmap_Dispose(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Dispose", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		hm.Dispose()
		if hm.Length() != 0 {
			t.Errorf("After Dispose: expected 0, got %d", hm.Length())
		}
	})
}

// ==========================================
// IsEqualPtr
// ==========================================

func Test_StrHashmap_IsEqualPtr_BothNil(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsEqualPtr_BothNil", func() {
		var a, b *corestr.Hashmap
		if !a.IsEqualPtr(b) {
			t.Error("Two nil hashmaps should be equal")
		}
	})
}

func Test_StrHashmap_IsEqualPtr_OneNil(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsEqualPtr_OneNil", func() {
		hm := corestr.New.Hashmap.Empty()
		var nilHm *corestr.Hashmap
		if hm.IsEqualPtr(nilHm) {
			t.Error("Non-nil vs nil should not be equal")
		}
	})
}

func Test_StrHashmap_IsEqualPtr_SamePointer(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsEqualPtr_SamePointer", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		if !hm.IsEqualPtr(hm) {
			t.Error("Same pointer should be equal")
		}
	})
}

func Test_StrHashmap_IsEqualPtr_SameContent(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsEqualPtr_SameContent", func() {
		a := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		b := corestr.New.Hashmap.UsingMap(map[string]string{"b": "2", "a": "1"})
		if !a.IsEqualPtr(b) {
			t.Error("Same content should be equal")
		}
	})
}

func Test_StrHashmap_IsEqualPtr_DifferentValues(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsEqualPtr_DifferentValues", func() {
		a := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		b := corestr.New.Hashmap.UsingMap(map[string]string{"a": "2"})
		if a.IsEqualPtr(b) {
			t.Error("Different values should not be equal")
		}
	})
}

func Test_StrHashmap_IsEqualPtr_DifferentKeys(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsEqualPtr_DifferentKeys", func() {
		a := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		b := corestr.New.Hashmap.UsingMap(map[string]string{"b": "1"})
		if a.IsEqualPtr(b) {
			t.Error("Different keys should not be equal")
		}
	})
}

func Test_StrHashmap_IsEqualPtr_DifferentLength(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsEqualPtr_DifferentLength", func() {
		a := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		b := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		if a.IsEqualPtr(b) {
			t.Error("Different length should not be equal")
		}
	})
}
func Test_StrHashmap_KeysToLower(t *testing.T) {
	safeTest(t, "Test_StrHashmap_KeysToLower", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"ABC": "val1", "Def": "val2"})
		lower := hm.KeysToLower()
		if !lower.Has("abc") || !lower.Has("def") {
			t.Error("KeysToLower should lowercase all keys")
		}
		if lower.Has("ABC") {
			t.Error("KeysToLower should not retain original case keys")
		}
		val, _ := lower.Get("abc")
		if val != "val1" {
			t.Errorf("KeysToLower should preserve values: expected 'val1', got '%s'", val)
		}
	})
}
func Test_StrHashmap_ValuesList_CacheInvalidatedAfterSet(t *testing.T) {
	safeTest(t, "Test_StrHashmap_ValuesList_CacheInvalidatedAfterSet", func() {
		hm := corestr.New.Hashmap.Empty()
		hm.Set("a", "1")
		list1 := hm.ValuesList()
		if len(list1) != 1 {
			t.Errorf("Initial ValuesList: expected 1, got %d", len(list1))
		}
		hm.Set("b", "2")
		list2 := hm.ValuesList()
		if len(list2) != 2 {
			t.Errorf("After Set, ValuesList should reflect new item: expected 2, got %d", len(list2))
		}
	})
}

func Test_StrHashmap_ValuesList_CacheInvalidatedAfterRemove(t *testing.T) {
	safeTest(t, "Test_StrHashmap_ValuesList_CacheInvalidatedAfterRemove", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		_ = hm.ValuesList() // populate cache
		hm.Remove("a")
		list := hm.ValuesList()
		if len(list) != 1 {
			t.Errorf("After Remove, ValuesList should reflect removal: expected 1, got %d", len(list))
		}
	})
}

// ==========================================
// Keys / Items
// ==========================================

func Test_StrHashmap_Keys(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Keys", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		keys := hm.Keys()
		if len(keys) != 2 {
			t.Errorf("Keys: expected 2, got %d", len(keys))
		}
	})
}

func Test_StrHashmap_Items(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Items", func() {
		m := map[string]string{"a": "1"}
		hm := corestr.New.Hashmap.UsingMap(m)
		items := hm.Items()
		if len(items) != 1 {
			t.Errorf("Items: expected 1, got %d", len(items))
		}
	})
}

// ==========================================
// Nil receiver guards
// ==========================================

func Test_StrHashmap_NilReceiver_IsEmpty(t *testing.T) {
	safeTest(t, "Test_StrHashmap_NilReceiver_IsEmpty", func() {
		var hm *corestr.Hashmap
		if !hm.IsEmpty() {
			t.Error("nil.IsEmpty() should return true")
		}
	})
}

func Test_StrHashmap_NilReceiver_Length(t *testing.T) {
	safeTest(t, "Test_StrHashmap_NilReceiver_Length", func() {
		var hm *corestr.Hashmap
		if hm.Length() != 0 {
			t.Error("nil.Length() should return 0")
		}
	})
}

func Test_StrHashmap_NilReceiver_HasItems(t *testing.T) {
	safeTest(t, "Test_StrHashmap_NilReceiver_HasItems", func() {
		var hm *corestr.Hashmap
		if hm.HasItems() {
			t.Error("nil.HasItems() should return false")
		}
	})
}
			t.Errorf("ConcatNew: expected 2, got %d", result.Length())
		}
		// original should not be mutated
		if hm1.Length() != 1 {
			t.Error("ConcatNew should not mutate original")
		}
	})
}

func Test_StrHashmap_ConcatNew_EmptyArgs(t *testing.T) {
	safeTest(t, "Test_StrHashmap_ConcatNew_EmptyArgs", func() {
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		result := hm.ConcatNew(true)
		if result.Length() != 1 {
			t.Errorf("ConcatNew empty: expected 1, got %d", result.Length())
		}
	})
}
