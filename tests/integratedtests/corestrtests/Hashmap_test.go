package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Constructors
// ==========================================

func Test_StrHashmap_Empty(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"result": hm.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Empty hashmap should be empty", actual)
		actual = args.Map{"result": hm.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty hashmap length: expected 0", actual)
	})
}

func Test_StrHashmap_Cap(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Cap", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(50)

		// Act
		actual := args.Map{"result": hm.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Cap hashmap should be empty initially", actual)
	})
}

func Test_StrHashmap_UsingMap(t *testing.T) {
	safeTest(t, "Test_StrHashmap_UsingMap", func() {
		// Arrange
		m := map[string]string{"k1": "v1", "k2": "v2"}
		hm := corestr.New.Hashmap.UsingMap(m)

		// Act
		actual := args.Map{"result": hm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "UsingMap: expected 2", actual)
	})
}

func Test_StrHashmap_KeyValuesStrings(t *testing.T) {
	safeTest(t, "Test_StrHashmap_KeyValuesStrings", func() {
		// Arrange
		hm := corestr.New.Hashmap.KeyValuesStrings(
			[]string{"a", "b"},
			[]string{"1", "2"},
		)

		// Act
		actual := args.Map{"result": hm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "KeyValuesStrings: expected 2", actual)
		val, found := hm.Get("a")
		actual = args.Map{"result": found || val != "1"}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Get('a'): expected '1', got '' (found=)", actual)
	})
}

func Test_StrHashmap_KeyValuesStrings_EmptyKeys(t *testing.T) {
	safeTest(t, "Test_StrHashmap_KeyValuesStrings_EmptyKeys", func() {
		// Arrange
		hm := corestr.New.Hashmap.KeyValuesStrings([]string{}, []string{})

		// Act
		actual := args.Map{"result": hm.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "KeyValuesStrings with empty keys should be empty", actual)
	})
}

// ==========================================
// Set / AddOrUpdate
// ==========================================

func Test_StrHashmap_Set_NewKey(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Set_NewKey", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		isNew := hm.Set("key", "val")

		// Act
		actual := args.Map{"result": isNew}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Set on new key should return true", actual)
		actual = args.Map{"result": hm.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "After Set: expected 1", actual)
	})
}

func Test_StrHashmap_Set_ExistingKey(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Set_ExistingKey", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.Set("key", "val1")
		isNew := hm.Set("key", "val2")

		// Act
		actual := args.Map{"result": isNew}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Set on existing key should return false", actual)
		val, _ := hm.Get("key")
		actual = args.Map{"result": val != "val2"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Set should overwrite: expected 'val2', got ''", actual)
	})
}

func Test_StrHashmap_AddOrUpdate_NewKey(t *testing.T) {
	safeTest(t, "Test_StrHashmap_AddOrUpdate_NewKey", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		isNew := hm.AddOrUpdate("k", "v")

		// Act
		actual := args.Map{"result": isNew}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdate on new key should return true", actual)
	})
}

func Test_StrHashmap_AddOrUpdate_ExistingKey(t *testing.T) {
	safeTest(t, "Test_StrHashmap_AddOrUpdate_ExistingKey", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v1")
		isNew := hm.AddOrUpdate("k", "v2")

		// Act
		actual := args.Map{"result": isNew}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddOrUpdate on existing key should return false", actual)
	})
}

func Test_StrHashmap_SetTrim(t *testing.T) {
	safeTest(t, "Test_StrHashmap_SetTrim", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.SetTrim("  key  ", "  val  ")
		val, found := hm.Get("key")

		// Act
		actual := args.Map{"result": found || val != "val"}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "SetTrim: expected trimmed key/val, got '' (found=)", actual)
	})
}

func Test_StrHashmap_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_StrHashmap_AddOrUpdateHashmap", func() {
		// Arrange
		hm1 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		hm2 := corestr.New.Hashmap.UsingMap(map[string]string{"b": "2", "a": "override"})
		hm1.AddOrUpdateHashmap(hm2)

		// Act
		actual := args.Map{"result": hm1.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "After merge: expected 2", actual)
		val, _ := hm1.Get("a")
		actual = args.Map{"result": val != "override"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Merge should overwrite: expected 'override', got ''", actual)
	})
}

func Test_StrHashmap_AddOrUpdateHashmap_Nil(t *testing.T) {
	safeTest(t, "Test_StrHashmap_AddOrUpdateHashmap_Nil", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		result := hm.AddOrUpdateHashmap(nil)

		// Act
		actual := args.Map{"result": result != hm || hm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateHashmap(nil) should be no-op", actual)
	})
}

func Test_StrHashmap_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_StrHashmap_AddOrUpdateMap", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdateMap(map[string]string{"x": "1", "y": "2"})

		// Act
		actual := args.Map{"result": hm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateMap: expected 2", actual)
	})
}

func Test_StrHashmap_AddOrUpdateMap_Empty(t *testing.T) {
	safeTest(t, "Test_StrHashmap_AddOrUpdateMap_Empty", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		hm.AddOrUpdateMap(map[string]string{})

		// Act
		actual := args.Map{"result": hm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateMap empty: expected 1", actual)
	})
}

// ==========================================
// Has / Contains / HasAll / HasAny
// ==========================================

func Test_StrHashmap_Has_Existing(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Has_Existing", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})

		// Act
		actual := args.Map{"result": hm.Has("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Has should find existing key", actual)
	})
}

func Test_StrHashmap_Has_Missing(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Has_Missing", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})

		// Act
		actual := args.Map{"result": hm.Has("missing")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Has should not find missing key", actual)
	})
}

func Test_StrHashmap_Contains_AliasForHas(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Contains_AliasForHas", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})

		// Act
		actual := args.Map{"result": hm.Contains("k") != hm.Has("k")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Contains should match Has", actual)
	})
}

func Test_StrHashmap_IsKeyMissing(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsKeyMissing", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})

		// Act
		actual := args.Map{"result": hm.IsKeyMissing("k")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsKeyMissing should return false for existing key", actual)
		actual = args.Map{"result": hm.IsKeyMissing("z")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsKeyMissing should return true for missing key", actual)
	})
}

func Test_StrHashmap_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_StrHashmap_HasAllStrings", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2", "c": "3"})

		// Act
		actual := args.Map{"result": hm.HasAllStrings("a", "c")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAllStrings should return true when all present", actual)
		actual = args.Map{"result": hm.HasAllStrings("a", "z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasAllStrings should return false when one missing", actual)
	})
}

func Test_StrHashmap_HasAll(t *testing.T) {
	safeTest(t, "Test_StrHashmap_HasAll", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})

		// Act
		actual := args.Map{"result": hm.HasAll("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAll should return true when all present", actual)
		actual = args.Map{"result": hm.HasAll("a", "x")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasAll should return false when one missing", actual)
	})
}

func Test_StrHashmap_HasAny_OnePresent(t *testing.T) {
	safeTest(t, "Test_StrHashmap_HasAny_OnePresent", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"result": hm.HasAny("z", "a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAny should return true when at least one present", actual)
	})
}

func Test_StrHashmap_HasAny_NonePresent(t *testing.T) {
	safeTest(t, "Test_StrHashmap_HasAny_NonePresent", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"result": hm.HasAny("x", "y")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasAny should return false when none present", actual)
	})
}

// ==========================================
// Get
// ==========================================

func Test_StrHashmap_Get_Existing(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Get_Existing", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		val, found := hm.Get("k")

		// Act
		actual := args.Map{"result": found}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Get should find existing key", actual)
		actual = args.Map{"result": val != "v"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Get: expected 'v', got ''", actual)
	})
}

func Test_StrHashmap_Get_Missing(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Get_Missing", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		val, found := hm.Get("missing")

		// Act
		actual := args.Map{"result": found}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Get should return false for missing key", actual)
		actual = args.Map{"result": val != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Get missing: expected empty string, got ''", actual)
	})
}

// ==========================================
// Remove
// ==========================================

func Test_StrHashmap_Remove(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Remove", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		hm.Remove("a")

		// Act
		actual := args.Map{"result": hm.Has("a")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Remove should delete key", actual)
		actual = args.Map{"result": hm.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "After remove: expected 1", actual)
	})
}

func Test_StrHashmap_Remove_Missing_NoEffect(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Remove_Missing_NoEffect", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		hm.Remove("z")

		// Act
		actual := args.Map{"result": hm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Remove missing: expected 1", actual)
	})
}

// ==========================================
// Clear / Dispose
// ==========================================

func Test_StrHashmap_Clear(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Clear", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		hm.Clear()

		// Act
		actual := args.Map{"result": hm.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Clear should make hashmap empty", actual)
	})
}

func Test_StrHashmap_Clear_NilReceiver(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Clear_NilReceiver", func() {
		// Arrange
		var hm *corestr.Hashmap
		result := hm.Clear()

		// Act
		actual := args.Map{"result": result != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Clear on nil should return nil", actual)
	})
}

func Test_StrHashmap_Dispose(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Dispose", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		hm.Dispose()

		// Act
		actual := args.Map{"result": hm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "After Dispose: expected 0", actual)
	})
}

// ==========================================
// IsEqualPtr
// ==========================================

func Test_StrHashmap_IsEqualPtr_BothNil(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsEqualPtr_BothNil", func() {
		// Arrange
		var a, b *corestr.Hashmap

		// Act
		actual := args.Map{"result": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Two nil hashmaps should be equal", actual)
	})
}

func Test_StrHashmap_IsEqualPtr_OneNil(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsEqualPtr_OneNil", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		var nilHm *corestr.Hashmap

		// Act
		actual := args.Map{"result": hm.IsEqualPtr(nilHm)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Non-nil vs nil should not be equal", actual)
	})
}

func Test_StrHashmap_IsEqualPtr_SamePointer(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsEqualPtr_SamePointer", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"result": hm.IsEqualPtr(hm)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Same pointer should be equal", actual)
	})
}

func Test_StrHashmap_IsEqualPtr_SameContent(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsEqualPtr_SameContent", func() {
		// Arrange
		a := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		b := corestr.New.Hashmap.UsingMap(map[string]string{"b": "2", "a": "1"})

		// Act
		actual := args.Map{"result": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Same content should be equal", actual)
	})
}

func Test_StrHashmap_IsEqualPtr_DifferentValues(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsEqualPtr_DifferentValues", func() {
		// Arrange
		a := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		b := corestr.New.Hashmap.UsingMap(map[string]string{"a": "2"})

		// Act
		actual := args.Map{"result": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Different values should not be equal", actual)
	})
}

func Test_StrHashmap_IsEqualPtr_DifferentKeys(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsEqualPtr_DifferentKeys", func() {
		// Arrange
		a := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		b := corestr.New.Hashmap.UsingMap(map[string]string{"b": "1"})

		// Act
		actual := args.Map{"result": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Different keys should not be equal", actual)
	})
}

func Test_StrHashmap_IsEqualPtr_DifferentLength(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsEqualPtr_DifferentLength", func() {
		// Arrange
		a := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		b := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})

		// Act
		actual := args.Map{"result": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Different length should not be equal", actual)
	})
}

func Test_StrHashmap_IsEqualPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_StrHashmap_IsEqualPtr_BothEmpty", func() {
		// Arrange
		a := corestr.New.Hashmap.Empty()
		b := corestr.New.Hashmap.Empty()

		// Act
		actual := args.Map{"result": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Two empty hashmaps should be equal", actual)
	})
}

// ==========================================
// KeysToLower (formerly ValuesToLower)
// ==========================================

func Test_StrHashmap_KeysToLower(t *testing.T) {
	safeTest(t, "Test_StrHashmap_KeysToLower", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"ABC": "val1", "Def": "val2"})
		lower := hm.KeysToLower()

		// Act
		actual := args.Map{"result": lower.Has("abc") || !lower.Has("def")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "KeysToLower should lowercase all keys", actual)
		actual = args.Map{"result": lower.Has("ABC")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "KeysToLower should not retain original case keys", actual)
		val, _ := lower.Get("abc")
		actual = args.Map{"result": val != "val1"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "KeysToLower should preserve values: expected 'val1', got ''", actual)
	})
}

func Test_StrHashmap_ValuesToLower_Deprecated_DelegatesToKeysToLower(t *testing.T) {
	safeTest(t, "Test_StrHashmap_ValuesToLower_Deprecated_DelegatesToKeysToLower", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"KEY": "val"})
		result := hm.ValuesToLower()

		// Act
		actual := args.Map{"result": result.Has("key")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Deprecated ValuesToLower should delegate to KeysToLower", actual)
	})
}

// ==========================================
// ValuesList caching (bug 42 context)
// ==========================================

func Test_StrHashmap_ValuesList_CacheInvalidatedAfterSet(t *testing.T) {
	safeTest(t, "Test_StrHashmap_ValuesList_CacheInvalidatedAfterSet", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.Set("a", "1")
		list1 := hm.ValuesList()

		// Act
		actual := args.Map{"result": len(list1) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Initial ValuesList: expected 1", actual)
		hm.Set("b", "2")
		list2 := hm.ValuesList()
		actual = args.Map{"result": len(list2) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "After Set, ValuesList should reflect new item: expected 2", actual)
	})
}

func Test_StrHashmap_ValuesList_CacheInvalidatedAfterRemove(t *testing.T) {
	safeTest(t, "Test_StrHashmap_ValuesList_CacheInvalidatedAfterRemove", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		_ = hm.ValuesList() // populate cache
		hm.Remove("a")
		list := hm.ValuesList()

		// Act
		actual := args.Map{"result": len(list) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "After Remove, ValuesList should reflect removal: expected 1", actual)
	})
}

// ==========================================
// Keys / Items
// ==========================================

func Test_StrHashmap_Keys(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Keys", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
		keys := hm.Keys()

		// Act
		actual := args.Map{"result": len(keys) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Keys: expected 2", actual)
	})
}

func Test_StrHashmap_Items(t *testing.T) {
	safeTest(t, "Test_StrHashmap_Items", func() {
		// Arrange
		m := map[string]string{"a": "1"}
		hm := corestr.New.Hashmap.UsingMap(m)
		items := hm.Items()

		// Act
		actual := args.Map{"result": len(items) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Items: expected 1", actual)
	})
}

// ==========================================
// Nil receiver guards
// ==========================================

func Test_StrHashmap_NilReceiver_IsEmpty(t *testing.T) {
	safeTest(t, "Test_StrHashmap_NilReceiver_IsEmpty", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act
		actual := args.Map{"result": hm.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "nil.IsEmpty() should return true", actual)
	})
}

func Test_StrHashmap_NilReceiver_Length(t *testing.T) {
	safeTest(t, "Test_StrHashmap_NilReceiver_Length", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act
		actual := args.Map{"result": hm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil.Length() should return 0", actual)
	})
}

func Test_StrHashmap_NilReceiver_HasItems(t *testing.T) {
	safeTest(t, "Test_StrHashmap_NilReceiver_HasItems", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act
		actual := args.Map{"result": hm.HasItems()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil.HasItems() should return false", actual)
	})
}

func Test_StrHashmap_NilReceiver_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_StrHashmap_NilReceiver_HasAnyItem", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act
		actual := args.Map{"result": hm.HasAnyItem()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil.HasAnyItem() should return false", actual)
	})
}

// ==========================================
// ClonePtr
// ==========================================

func Test_StrHashmap_ClonePtr_NilReceiver(t *testing.T) {
	safeTest(t, "Test_StrHashmap_ClonePtr_NilReceiver", func() {
		// Arrange
		var hm *corestr.Hashmap
		result := hm.ClonePtr()

		// Act
		actual := args.Map{"result": result != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ClonePtr on nil should return nil", actual)
	})
}

func Test_StrHashmap_ClonePtr_Independence(t *testing.T) {
	safeTest(t, "Test_StrHashmap_ClonePtr_Independence", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		cloned := hm.ClonePtr()
		cloned.Set("b", "2")

		// Act
		actual := args.Map{"result": hm.Has("b")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Clone should be independent from original", actual)
	})
}

// ==========================================
// ConcatNew
// ==========================================

func Test_StrHashmap_ConcatNew(t *testing.T) {
	safeTest(t, "Test_StrHashmap_ConcatNew", func() {
		// Arrange
		hm1 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		hm2 := corestr.New.Hashmap.UsingMap(map[string]string{"b": "2"})
		result := hm1.ConcatNew(false, hm2)

		// Act
		actual := args.Map{"result": result.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ConcatNew: expected 2", actual)
		// original should not be mutated
		actual = args.Map{"result": hm1.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ConcatNew should not mutate original", actual)
	})
}

func Test_StrHashmap_ConcatNew_EmptyArgs(t *testing.T) {
	safeTest(t, "Test_StrHashmap_ConcatNew_EmptyArgs", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		result := hm.ConcatNew(true)

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ConcatNew empty: expected 1", actual)
	})
}
