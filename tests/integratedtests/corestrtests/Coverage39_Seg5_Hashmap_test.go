package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Segment 5a
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg5_HM_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_IsEmpty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)

		// Act
		actual := args.Map{
			"empty": h.IsEmpty(),
			"hasItems": h.HasItems(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"hasItems": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- true on empty", actual)
	})
}

func Test_Seg5_HM_AddOrUpdate(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdate", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		isNew := h.AddOrUpdate("k", "v")
		isNew2 := h.AddOrUpdate("k", "v2")

		// Act
		actual := args.Map{
			"isNew": isNew,
			"isNew2": isNew2,
			"len": h.Length(),
		}

		// Assert
		expected := args.Map{
			"isNew": true,
			"isNew2": false,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "AddOrUpdate -- new then update", actual)
	})
}

func Test_Seg5_HM_Set(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Set", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		isNew := h.Set("a", "1")

		// Act
		actual := args.Map{"isNew": isNew}

		// Assert
		expected := args.Map{"isNew": true}
		expected.ShouldBeEqual(t, 0, "Set -- new key", actual)
	})
}

func Test_Seg5_HM_SetTrim(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_SetTrim", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.SetTrim(" a ", " 1 ")
		val, found := h.Get("a")

		// Act
		actual := args.Map{
			"found": found,
			"val": val,
		}

		// Assert
		expected := args.Map{
			"found": true,
			"val": "1",
		}
		expected.ShouldBeEqual(t, 0, "SetTrim -- trimmed key and val", actual)
	})
}

func Test_Seg5_HM_SetBySplitter(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_SetBySplitter", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.SetBySplitter("=", "key=value")
		val, found := h.Get("key")

		// Act
		actual := args.Map{
			"found": found,
			"val": val,
		}

		// Assert
		expected := args.Map{
			"found": true,
			"val": "value",
		}
		expected.ShouldBeEqual(t, 0, "SetBySplitter -- split key=value", actual)
	})
}

func Test_Seg5_HM_SetBySplitter_NoValue(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_SetBySplitter_NoValue", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.SetBySplitter("=", "keyonly")
		val, found := h.Get("keyonly")

		// Act
		actual := args.Map{
			"found": found,
			"val": val,
		}

		// Assert
		expected := args.Map{
			"found": true,
			"val": "",
		}
		expected.ShouldBeEqual(t, 0, "SetBySplitter no value -- empty val", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateKeyStrValInt(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyStrValInt", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyStrValInt("n", 42)
		val, _ := h.Get("n")

		// Act
		actual := args.Map{"val": val}

		// Assert
		expected := args.Map{"val": "42"}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyStrValInt -- int to string", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateKeyStrValFloat(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyStrValFloat", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyStrValFloat("f", 1.5)
		_, found := h.Get("f")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyStrValFloat -- stored", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyStrValFloat64", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyStrValFloat64("f", 2.5)
		_, found := h.Get("f")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyStrValFloat64 -- stored", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateKeyStrValAny(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyStrValAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyStrValAny("a", 123)
		_, found := h.Get("a")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyStrValAny -- stored", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateKeyValueAny(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyValueAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "k", Value: 99})
		_, found := h.Get("k")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyValueAny -- stored", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateKeyVal(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyVal", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		isNew := h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v"})

		// Act
		actual := args.Map{"isNew": isNew}

		// Assert
		expected := args.Map{"isNew": true}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyVal -- new", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateWithWgLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateWithWgLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateWithWgLock("k", "v", &wg)
		wg.Wait()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateWithWgLock -- added", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateHashmap", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.Set("b", "2")
		h.AddOrUpdateHashmap(h2)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateHashmap -- merged", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateHashmap_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateHashmap_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h.AddOrUpdateHashmap(nil)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateHashmap nil -- no change", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateMap", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateMap(map[string]string{"a": "1", "b": "2"})

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateMap -- 2 items", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateMap_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateMap_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateMap(map[string]string{})

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateMap empty -- no change", actual)
	})
}

func Test_Seg5_HM_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdates", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddsOrUpdates(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdates -- 2 items", actual)
	})
}

func Test_Seg5_HM_AddsOrUpdates_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdates_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddsOrUpdates(nil...)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdates nil -- no change", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateKeyAnyValues(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyAnyValues", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyAnyValues(corestr.KeyAnyValuePair{Key: "k", Value: 1})

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyAnyValues -- 1 item", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateKeyAnyValues_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyAnyValues_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyAnyValues()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyAnyValues empty -- no change", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateKeyValues(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyValues", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyValues -- 1 item", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateKeyValues_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateKeyValues_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyValues()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateKeyValues empty -- no change", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		keys := corestr.New.Collection.Strings([]string{"a", "b"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})
		h.AddOrUpdateCollection(keys, vals)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateCollection -- 2 items", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateCollection_NilKeys(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateCollection_NilKeys", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateCollection(nil, nil)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateCollection nil -- no change", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateCollection_LenMismatch(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateCollection_LenMismatch", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		keys := corestr.New.Collection.Strings([]string{"a"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})
		h.AddOrUpdateCollection(keys, vals)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateCollection mismatch -- no change", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateStringsPtrWgLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateStringsPtrWgLock(&wg, []string{"a", "b"}, []string{"1", "2"})
		wg.Wait()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateStringsPtrWgLock -- 2 items", actual)
	})
}

func Test_Seg5_HM_AddOrUpdateStringsPtrWgLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateStringsPtrWgLock_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateStringsPtrWgLock(&wg, []string{}, []string{})
		wg.Wait()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateStringsPtrWgLock empty -- no change", actual)
	})
}

// ── Has / Contains / Missing ────────────────────────────────────────────────

func Test_Seg5_HM_Has_Contains(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Has_Contains", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{
			"has":      h.Has("a"),
			"miss":     h.Has("z"),
			"contains": h.Contains("a"),
			"missing":  h.IsKeyMissing("z"),
		}

		// Assert
		expected := args.Map{
			"has":      true,
			"miss":     false,
			"contains": true,
			"missing":  true,
		}
		expected.ShouldBeEqual(t, 0, "Has/Contains/Missing -- correct", actual)
	})
}

func Test_Seg5_HM_HasLock_ContainsLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_HasLock_ContainsLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{
			"hasLock":      h.HasLock("a"),
			"containsLock": h.ContainsLock("a"),
			"missingLock":  h.IsKeyMissingLock("z"),
			"hasWithLock":  h.HasWithLock("a"),
		}

		// Assert
		expected := args.Map{
			"hasLock":      true,
			"containsLock": true,
			"missingLock":  true,
			"hasWithLock":  true,
		}
		expected.ShouldBeEqual(t, 0, "HasLock/ContainsLock -- correct", actual)
	})
}

func Test_Seg5_HM_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_HasAllStrings", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h.Set("b", "2")

		// Act
		actual := args.Map{
			"all":  h.HasAllStrings("a", "b"),
			"miss": h.HasAllStrings("a", "z"),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAllStrings -- all and missing", actual)
	})
}

func Test_Seg5_HM_HasAll(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_HasAll", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{
			"has": h.HasAll("a"),
			"miss": h.HasAll("z"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAll -- found and missing", actual)
	})
}

func Test_Seg5_HM_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_HasAllCollectionItems", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"has": h.HasAllCollectionItems(c),
			"nil": h.HasAllCollectionItems(nil),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAllCollectionItems -- found and nil", actual)
	})
}

func Test_Seg5_HM_HasAny(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_HasAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{
			"any": h.HasAny("z", "a"),
			"none": h.HasAny("x", "y"),
		}

		// Assert
		expected := args.Map{
			"any": true,
			"none": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAny -- found and none", actual)
	})
}

func Test_Seg5_HM_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_HasAnyItem", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"has": h.HasAnyItem()}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasAnyItem -- true", actual)
	})
}

// ── Get / GetValue ──────────────────────────────────────────────────────────

func Test_Seg5_HM_Get(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Get", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		val, found := h.Get("a")
		val2, found2 := h.GetValue("a")

		// Act
		actual := args.Map{
			"val": val,
			"found": found,
			"val2": val2,
			"found2": found2,
		}

		// Assert
		expected := args.Map{
			"val": "1",
			"found": true,
			"val2": "1",
			"found2": true,
		}
		expected.ShouldBeEqual(t, 0, "Get/GetValue -- found", actual)
	})
}

// ── Items / Keys / Values ───────────────────────────────────────────────────

func Test_Seg5_HM_Items(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Items", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": len(h.Items())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Items -- 1 item", actual)
	})
}

func Test_Seg5_HM_SafeItems_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_SafeItems_Nil", func() {
		// Arrange
		var h *corestr.Hashmap

		// Act
		actual := args.Map{"nil": h.SafeItems() == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "SafeItems nil -- nil", actual)
	})
}

func Test_Seg5_HM_Keys(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Keys", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{
			"len": len(h.Keys()),
			"allLen": len(h.AllKeys()),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"allLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "Keys/AllKeys -- 1 key", actual)
	})
}

func Test_Seg5_HM_KeysCollection(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": h.KeysCollection().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeysCollection -- 1 key", actual)
	})
}

func Test_Seg5_HM_ValuesList(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ValuesList", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": len(h.ValuesList())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesList -- 1 value", actual)
	})
}

func Test_Seg5_HM_ValuesCollection(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ValuesCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": h.ValuesCollection().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesCollection -- 1 value", actual)
	})
}

func Test_Seg5_HM_ValuesHashset(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ValuesHashset", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": h.ValuesHashset().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesHashset -- 1 value", actual)
	})
}

func Test_Seg5_HM_KeysValuesList(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysValuesList", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		keys, values := h.KeysValuesList()

		// Act
		actual := args.Map{
			"kLen": len(keys),
			"vLen": len(values),
		}

		// Assert
		expected := args.Map{
			"kLen": 1,
			"vLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "KeysValuesList -- 1 each", actual)
	})
}

func Test_Seg5_HM_KeysValuesCollection(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysValuesCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		keys, values := h.KeysValuesCollection()

		// Act
		actual := args.Map{
			"kLen": keys.Length(),
			"vLen": values.Length(),
		}

		// Assert
		expected := args.Map{
			"kLen": 1,
			"vLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "KeysValuesCollection -- 1 each", actual)
	})
}

func Test_Seg5_HM_KeysValuePairs(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysValuePairs", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": len(h.KeysValuePairs())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeysValuePairs -- 1 pair", actual)
	})
}

func Test_Seg5_HM_KeysValuePairsCollection(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysValuePairsCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": h.KeysValuePairsCollection().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeysValuePairsCollection -- 1 pair", actual)
	})
}

func Test_Seg5_HM_KeysLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": len(h.KeysLock())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeysLock -- 1 key", actual)
	})
}

func Test_Seg5_HM_KeysLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysLock_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)

		// Act
		actual := args.Map{"len": len(h.KeysLock())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "KeysLock empty -- 0", actual)
	})
}

func Test_Seg5_HM_KeysValuesListLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysValuesListLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		keys, vals := h.KeysValuesListLock()

		// Act
		actual := args.Map{
			"kLen": len(keys),
			"vLen": len(vals),
		}

		// Assert
		expected := args.Map{
			"kLen": 1,
			"vLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "KeysValuesListLock -- 1 each", actual)
	})
}

func Test_Seg5_HM_ValuesListCopyLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ValuesListCopyLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": len(h.ValuesListCopyLock())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesListCopyLock -- 1", actual)
	})
}

func Test_Seg5_HM_ValuesCollectionLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ValuesCollectionLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": h.ValuesCollectionLock().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesCollectionLock -- 1", actual)
	})
}

func Test_Seg5_HM_ValuesHashsetLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ValuesHashsetLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": h.ValuesHashsetLock().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValuesHashsetLock -- 1", actual)
	})
}

func Test_Seg5_HM_ItemsCopyLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ItemsCopyLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		copied := h.ItemsCopyLock()

		// Act
		actual := args.Map{"len": len(*copied)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ItemsCopyLock -- 1", actual)
	})
}

// ── Filter ──────────────────────────────────────────────────────────────────

func Test_Seg5_HM_GetKeysFilteredItems(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetKeysFilteredItems", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.Set("a", "1")
		h.Set("bb", "2")
		result := h.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, len(s) == 1, false
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredItems -- 1 match", actual)
	})
}

func Test_Seg5_HM_GetKeysFilteredItems_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetKeysFilteredItems_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)
		result := h.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredItems empty -- 0", actual)
	})
}

func Test_Seg5_HM_GetKeysFilteredItems_Break(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetKeysFilteredItems_Break", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.Set("a", "1")
		h.Set("b", "2")
		h.Set("c", "3")
		result := h.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})

		// Act
		actual := args.Map{"hasItems": len(result) > 0}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredItems break -- stops early", actual)
	})
}

func Test_Seg5_HM_GetKeysFilteredCollection(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetKeysFilteredCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.Set("a", "1")
		h.Set("bb", "2")
		result := h.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, len(s) == 1, false
		})

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredCollection -- 1 match", actual)
	})
}

func Test_Seg5_HM_GetKeysFilteredCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetKeysFilteredCollection_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)
		result := h.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"empty": result.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredCollection empty -- empty", actual)
	})
}

func Test_Seg5_HM_GetKeysFilteredCollection_Break(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetKeysFilteredCollection_Break", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.Set("a", "1")
		h.Set("b", "2")
		result := h.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})

		// Act
		actual := args.Map{"hasItems": result.HasAnyItem()}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "GetKeysFilteredCollection break -- stops", actual)
	})
}

// ── AddsOrUpdatesAnyUsingFilter ─────────────────────────────────────────────

func Test_Seg5_HM_AddsOrUpdatesAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdatesAnyUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.AddsOrUpdatesAnyUsingFilter(
			func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
				return pair.ValueString(), true, false
			},
			corestr.KeyAnyValuePair{Key: "a", Value: 1},
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilter -- 1 kept", actual)
	})
}

func Test_Seg5_HM_AddsOrUpdatesAnyUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdatesAnyUsingFilter_Break", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.AddsOrUpdatesAnyUsingFilter(
			func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
				return pair.ValueString(), true, true
			},
			corestr.KeyAnyValuePair{Key: "a", Value: 1},
			corestr.KeyAnyValuePair{Key: "b", Value: 2},
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilter break -- only 1", actual)
	})
}

func Test_Seg5_HM_AddsOrUpdatesAnyUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdatesAnyUsingFilter_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddsOrUpdatesAnyUsingFilter(nil, nil...)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilter nil -- no change", actual)
	})
}

func Test_Seg5_HM_AddsOrUpdatesAnyUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdatesAnyUsingFilterLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.AddsOrUpdatesAnyUsingFilterLock(
			func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
				return pair.ValueString(), true, false
			},
			corestr.KeyAnyValuePair{Key: "a", Value: 1},
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilterLock -- 1 kept", actual)
	})
}

func Test_Seg5_HM_AddsOrUpdatesAnyUsingFilterLock_Break(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdatesAnyUsingFilterLock_Break", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.AddsOrUpdatesAnyUsingFilterLock(
			func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
				return pair.ValueString(), true, true
			},
			corestr.KeyAnyValuePair{Key: "a", Value: 1},
			corestr.KeyAnyValuePair{Key: "b", Value: 2},
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesAnyUsingFilterLock break -- only 1", actual)
	})
}

func Test_Seg5_HM_AddsOrUpdatesUsingFilter(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdatesUsingFilter", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.AddsOrUpdatesUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Value, true, false
			},
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesUsingFilter -- 1 kept", actual)
	})
}

func Test_Seg5_HM_AddsOrUpdatesUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdatesUsingFilter_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddsOrUpdatesUsingFilter(nil, nil...)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesUsingFilter nil -- no change", actual)
	})
}

func Test_Seg5_HM_AddsOrUpdatesUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddsOrUpdatesUsingFilter_Break", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.AddsOrUpdatesUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Value, true, true
			},
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsOrUpdatesUsingFilter break -- only 1", actual)
	})
}

// ── Concat ──────────────────────────────────────────────────────────────────

func Test_Seg5_HM_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ConcatNew", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.Set("b", "2")
		result := h.ConcatNew(true, h2)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ConcatNew -- merged", actual)
	})
}

func Test_Seg5_HM_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ConcatNew_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		result := h.ConcatNew(true)

		// Act
		actual := args.Map{"hasItems": result.HasAnyItem()}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "ConcatNew empty -- cloned", actual)
	})
}

func Test_Seg5_HM_ConcatNewUsingMaps(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ConcatNewUsingMaps", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		result := h.ConcatNewUsingMaps(true, map[string]string{"b": "2"})

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ConcatNewUsingMaps -- merged", actual)
	})
}

func Test_Seg5_HM_ConcatNewUsingMaps_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ConcatNewUsingMaps_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		result := h.ConcatNewUsingMaps(true)

		// Act
		actual := args.Map{"hasItems": result.HasAnyItem()}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "ConcatNewUsingMaps empty -- cloned", actual)
	})
}

// ── Lock variants ───────────────────────────────────────────────────────────

func Test_Seg5_HM_AddOrUpdateLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_AddOrUpdateLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateLock("k", "v")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddOrUpdateLock -- 1 item", actual)
	})
}

func Test_Seg5_HM_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_IsEmptyLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)

		// Act
		actual := args.Map{"empty": h.IsEmptyLock()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock -- true", actual)
	})
}

func Test_Seg5_HM_LengthLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_LengthLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": h.LengthLock()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthLock -- 1", actual)
	})
}

func Test_Seg5_HM_Length_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Length_Nil", func() {
		// Arrange
		var h *corestr.Hashmap

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Length nil -- 0", actual)
	})
}

// ── IsEqual ─────────────────────────────────────────────────────────────────

func Test_Seg5_HM_IsEqualPtr(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_IsEqualPtr", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(2)
		h1.Set("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.Set("a", "1")
		h3 := corestr.New.Hashmap.Cap(2)
		h3.Set("a", "2")

		// Act
		actual := args.Map{
			"eq":      h1.IsEqualPtr(h2),
			"neq":     h1.IsEqualPtr(h3),
			"same":    h1.IsEqualPtr(h1),
			"nilBoth": (*corestr.Hashmap)(nil).IsEqualPtr(nil),
			"nilOne":  h1.IsEqualPtr(nil),
		}

		// Assert
		expected := args.Map{
			"eq":      true,
			"neq":     false,
			"same":    true,
			"nilBoth": true,
			"nilOne":  false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr -- various", actual)
	})
}

func Test_Seg5_HM_IsEqualPtr_DiffLen(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_IsEqualPtr_DiffLen", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(2)
		h1.Set("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.Set("a", "1")
		h2.Set("b", "2")

		// Act
		actual := args.Map{"eq": h1.IsEqualPtr(h2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr diff len -- false", actual)
	})
}

func Test_Seg5_HM_IsEqualPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_IsEqualPtr_BothEmpty", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(0)
		h2 := corestr.New.Hashmap.Cap(0)

		// Act
		actual := args.Map{"eq": h1.IsEqualPtr(h2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr both empty -- true", actual)
	})
}

func Test_Seg5_HM_IsEqualPtr_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_IsEqualPtr_OneEmpty", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(2)
		h1.Set("a", "1")
		h2 := corestr.New.Hashmap.Cap(0)

		// Act
		actual := args.Map{"eq": h1.IsEqualPtr(h2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr one empty -- false", actual)
	})
}

func Test_Seg5_HM_IsEqualPtrLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_IsEqualPtrLock", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(2)
		h1.Set("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.Set("a", "1")

		// Act
		actual := args.Map{"eq": h1.IsEqualPtrLock(h2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualPtrLock -- true", actual)
	})
}

// ── Remove ──────────────────────────────────────────────────────────────────

func Test_Seg5_HM_Remove(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Remove", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h.Remove("a")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Remove -- removed", actual)
	})
}

func Test_Seg5_HM_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_RemoveWithLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h.RemoveWithLock("a")

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "RemoveWithLock -- removed", actual)
	})
}

// ── String / Join ───────────────────────────────────────────────────────────

func Test_Seg5_HM_String(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_String", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"nonEmpty": h.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Seg5_HM_String_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_String_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)

		// Act
		actual := args.Map{"nonEmpty": h.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String empty -- has NoElements text", actual)
	})
}

func Test_Seg5_HM_StringLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_StringLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"nonEmpty": h.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock -- non-empty", actual)
	})
}

func Test_Seg5_HM_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_StringLock_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)

		// Act
		actual := args.Map{"nonEmpty": h.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock empty -- has NoElements text", actual)
	})
}

func Test_Seg5_HM_Join(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Join", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"nonEmpty": h.Join(",") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Join -- non-empty", actual)
	})
}

func Test_Seg5_HM_JoinKeys(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_JoinKeys", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"val": h.JoinKeys(",")}

		// Assert
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "JoinKeys -- a", actual)
	})
}

// ── Except ──────────────────────────────────────────────────────────────────

func Test_Seg5_HM_GetValuesExceptKeysInHashset(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetValuesExceptKeysInHashset", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.Set("a", "1")
		h.Set("b", "2")
		except := corestr.New.Hashset.Strings([]string{"a"})
		result := h.GetValuesExceptKeysInHashset(except)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetValuesExceptKeysInHashset -- 1 remaining", actual)
	})
}

func Test_Seg5_HM_GetValuesExceptKeysInHashset_NilExcept(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetValuesExceptKeysInHashset_NilExcept", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		result := h.GetValuesExceptKeysInHashset(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetValuesExceptKeysInHashset nil -- all values", actual)
	})
}

func Test_Seg5_HM_GetValuesKeysExcept(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetValuesKeysExcept", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.Set("a", "1")
		h.Set("b", "2")
		result := h.GetValuesKeysExcept([]string{"a"})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetValuesKeysExcept -- 1 remaining", actual)
	})
}

func Test_Seg5_HM_GetValuesKeysExcept_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetValuesKeysExcept_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		result := h.GetValuesKeysExcept(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetValuesKeysExcept nil -- all values", actual)
	})
}

func Test_Seg5_HM_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetAllExceptCollection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(4)
		h.Set("a", "1")
		h.Set("b", "2")
		c := corestr.New.Collection.Strings([]string{"a"})
		result := h.GetAllExceptCollection(c)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection -- 1 remaining", actual)
	})
}

func Test_Seg5_HM_GetAllExceptCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_GetAllExceptCollection_Nil", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		result := h.GetAllExceptCollection(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection nil -- all values", actual)
	})
}

// ── KeysToLower / ValuesToLower ─────────────────────────────────────────────

func Test_Seg5_HM_KeysToLower(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeysToLower", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("ABC", "1")
		result := h.KeysToLower()
		_, found := result.Get("abc")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "KeysToLower -- lowered", actual)
	})
}

func Test_Seg5_HM_ValuesToLower(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ValuesToLower", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("ABC", "1")
		result := h.ValuesToLower()
		_, found := result.Get("abc")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "ValuesToLower -- delegates to KeysToLower", actual)
	})
}

// ── Diff ────────────────────────────────────────────────────────────────────

func Test_Seg5_HM_Diff(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Diff", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.Set("a", "2")
		diff := h.Diff(h2)

		// Act
		actual := args.Map{"notNil": diff != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Diff -- non-nil", actual)
	})
}

// ── JSON ────────────────────────────────────────────────────────────────────

func Test_Seg5_HM_Json(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Json", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		j := h.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Seg5_HM_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_MarshalJSON", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		b, err := h.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "MarshalJSON -- success", actual)
	})
}

func Test_Seg5_HM_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_UnmarshalJSON", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)
		err := h.UnmarshalJSON([]byte(`{"a":"1"}`))

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": h.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_Seg5_HM_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_UnmarshalJSON_Invalid", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)
		err := h.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_Seg5_HM_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ParseInjectUsingJson", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		jr := h.JsonPtr()
		h2 := corestr.New.Hashmap.Cap(0)
		result, err := h2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": result.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- round trip", actual)
	})
}

func Test_Seg5_HM_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ParseInjectUsingJsonMust", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		jr := h.JsonPtr()
		h2 := corestr.New.Hashmap.Cap(0)
		result := h2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_Seg5_HM_Serialize(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Serialize", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		b, err := h.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "Serialize -- success", actual)
	})
}

func Test_Seg5_HM_Deserialize(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Deserialize", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		var dest map[string]string
		err := h.Deserialize(&dest)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": len(dest),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "Deserialize -- success", actual)
	})
}

func Test_Seg5_HM_JsonModel(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_JsonModel", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": len(h.JsonModel())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel -- 1 item", actual)
	})
}

func Test_Seg5_HM_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_JsonModelAny", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"notNil": h.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_Seg5_HM_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_InterfaceCasts", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)

		// Act
		actual := args.Map{
			"jsoner":   h.AsJsoner() != nil,
			"binder":   h.AsJsonContractsBinder() != nil,
			"injector": h.AsJsonParseSelfInjector() != nil,
			"marsh":    h.AsJsonMarshaller() != nil,
		}

		// Assert
		expected := args.Map{
			"jsoner":   true,
			"binder":   true,
			"injector": true,
			"marsh":    true,
		}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

func Test_Seg5_HM_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_JsonParseSelfInject", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		jr := h.JsonPtr()
		h2 := corestr.New.Hashmap.Cap(0)
		err := h2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

// ── Error / KeyValStringLines ───────────────────────────────────────────────

func Test_Seg5_HM_ToError(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ToError", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		err := h.ToError("; ")

		// Act
		actual := args.Map{"notNil": err != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ToError -- non-nil", actual)
	})
}

func Test_Seg5_HM_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ToDefaultError", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		err := h.ToDefaultError()

		// Act
		actual := args.Map{"notNil": err != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ToDefaultError -- non-nil", actual)
	})
}

func Test_Seg5_HM_KeyValStringLines(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_KeyValStringLines", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": len(h.KeyValStringLines())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeyValStringLines -- 1 line", actual)
	})
}

func Test_Seg5_HM_ToStringsUsingCompiler(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ToStringsUsingCompiler", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		result := h.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })

		// Act
		actual := args.Map{
			"len": len(result),
			"val": result[0],
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"val": "a=1",
		}
		expected.ShouldBeEqual(t, 0, "ToStringsUsingCompiler -- formatted", actual)
	})
}

func Test_Seg5_HM_ToStringsUsingCompiler_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ToStringsUsingCompiler_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)
		result := h.ToStringsUsingCompiler(func(k, v string) string { return k })

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ToStringsUsingCompiler empty -- empty", actual)
	})
}

// ── Clone / Clear / Dispose ─────────────────────────────────────────────────

func Test_Seg5_HM_Clone(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Clone", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		c := h.Clone()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Clone -- same items", actual)
	})
}

func Test_Seg5_HM_Clone_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Clone_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(0)
		c := h.Clone()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Clone empty -- empty", actual)
	})
}

func Test_Seg5_HM_ClonePtr(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ClonePtr", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		c := h.ClonePtr()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ClonePtr -- same items", actual)
	})
}

func Test_Seg5_HM_ClonePtr_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_ClonePtr_Nil", func() {
		// Arrange
		var h *corestr.Hashmap

		// Act
		actual := args.Map{"nil": h.ClonePtr() == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "ClonePtr nil -- returns nil", actual)
	})
}

func Test_Seg5_HM_Clear(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Clear", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h.Clear()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Seg5_HM_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Clear_Nil", func() {
		// Arrange
		var h *corestr.Hashmap

		// Act
		actual := args.Map{"nil": h.Clear() == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Clear nil -- returns nil", actual)
	})
}

func Test_Seg5_HM_Dispose(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Dispose", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		h.Dispose()

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleaned up", actual)
	})
}

func Test_Seg5_HM_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Dispose_Nil", func() {
		var h *corestr.Hashmap
		h.Dispose() // should not panic
	})
}

// ── Collection ──────────────────────────────────────────────────────────────

func Test_Seg5_HM_Collection(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_Collection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")

		// Act
		actual := args.Map{"len": h.Collection().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection -- 1 item", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// HashmapDataModel
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg5_HM_DataModel(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_DataModel", func() {
		// Arrange
		dm := &corestr.HashmapDataModel{Items: map[string]string{"a": "1"}}
		h := corestr.NewHashmapUsingDataModel(dm)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "NewHashmapUsingDataModel -- 1 item", actual)
	})
}

func Test_Seg5_HM_DataModel_Reverse(t *testing.T) {
	safeTest(t, "Test_Seg5_HM_DataModel_Reverse", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(2)
		h.Set("a", "1")
		dm := corestr.NewHashmapsDataModelUsing(h)

		// Act
		actual := args.Map{"len": len(dm.Items)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "NewHashmapsDataModelUsing -- 1 item", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// HashmapDiff
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg5_HMD_Length(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_Length", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		actual := args.Map{
			"len": d.Length(),
			"empty": d.IsEmpty(),
			"hasAny": d.HasAnyItem(),
			"last": d.LastIndex(),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"empty": false,
			"hasAny": true,
			"last": 0,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff -- basic props", actual)
	})
}

func Test_Seg5_HMD_Length_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_Length_Nil", func() {
		// Arrange
		var d *corestr.HashmapDiff

		// Act
		actual := args.Map{"len": d.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "HashmapDiff nil -- 0", actual)
	})
}

func Test_Seg5_HMD_Raw(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_Raw", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"len": len(d.Raw())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashmapDiff Raw -- 1 item", actual)
	})
}

func Test_Seg5_HMD_Raw_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_Raw_Nil", func() {
		// Arrange
		var d *corestr.HashmapDiff

		// Act
		actual := args.Map{"len": len(d.Raw())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "HashmapDiff Raw nil -- empty", actual)
	})
}

func Test_Seg5_HMD_MapAnyItems(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_MapAnyItems", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"len": len(d.MapAnyItems())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "MapAnyItems -- 1 item", actual)
	})
}

func Test_Seg5_HMD_MapAnyItems_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_MapAnyItems_Nil", func() {
		// Arrange
		var d *corestr.HashmapDiff

		// Act
		actual := args.Map{"len": len(d.MapAnyItems())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "MapAnyItems nil -- empty", actual)
	})
}

func Test_Seg5_HMD_AllKeysSorted(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_AllKeysSorted", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"b": "2", "a": "1"})
		keys := d.AllKeysSorted()

		// Act
		actual := args.Map{"first": keys[0]}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "AllKeysSorted -- sorted", actual)
	})
}

func Test_Seg5_HMD_IsRawEqual(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_IsRawEqual", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		actual := args.Map{
			"eq":  d.IsRawEqual(map[string]string{"a": "1"}),
			"neq": d.IsRawEqual(map[string]string{"a": "2"}),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"neq": false,
		}
		expected.ShouldBeEqual(t, 0, "IsRawEqual -- match and mismatch", actual)
	})
}

func Test_Seg5_HMD_HasAnyChanges(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_HasAnyChanges", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"has": d.HasAnyChanges(map[string]string{"a": "2"})}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasAnyChanges -- true", actual)
	})
}

func Test_Seg5_HMD_DiffRaw(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_DiffRaw", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		diff := d.DiffRaw(map[string]string{"a": "2"})

		// Act
		actual := args.Map{"hasItems": len(diff) > 0}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "DiffRaw -- has diff", actual)
	})
}

func Test_Seg5_HMD_HashmapDiffUsingRaw(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_HashmapDiffUsingRaw", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		diff := d.HashmapDiffUsingRaw(map[string]string{"a": "2"})

		// Act
		actual := args.Map{"hasItems": diff.HasAnyItem()}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "HashmapDiffUsingRaw -- has diff", actual)
	})
}

func Test_Seg5_HMD_HashmapDiffUsingRaw_NoDiff(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_HashmapDiffUsingRaw_NoDiff", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		diff := d.HashmapDiffUsingRaw(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"empty": diff.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashmapDiffUsingRaw no diff -- empty", actual)
	})
}

func Test_Seg5_HMD_DiffJsonMessage(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_DiffJsonMessage", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		msg := d.DiffJsonMessage(map[string]string{"a": "2"})

		// Act
		actual := args.Map{"nonEmpty": msg != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "DiffJsonMessage -- non-empty", actual)
	})
}

func Test_Seg5_HMD_ShouldDiffMessage(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_ShouldDiffMessage", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		msg := d.ShouldDiffMessage("test", map[string]string{"a": "2"})

		// Act
		actual := args.Map{"nonEmpty": msg != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "ShouldDiffMessage -- non-empty", actual)
	})
}

func Test_Seg5_HMD_LogShouldDiffMessage(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_LogShouldDiffMessage", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		msg := d.LogShouldDiffMessage("test", map[string]string{"a": "2"})

		// Act
		actual := args.Map{"nonEmpty": msg != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "LogShouldDiffMessage -- non-empty", actual)
	})
}

func Test_Seg5_HMD_ToStringsSliceOfDiffMap(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_ToStringsSliceOfDiffMap", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		result := d.ToStringsSliceOfDiffMap(map[string]string{"a": "2"})

		// Act
		actual := args.Map{"hasItems": len(result) > 0}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "ToStringsSliceOfDiffMap -- non-empty", actual)
	})
}

func Test_Seg5_HMD_RawMapStringAnyDiff(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_RawMapStringAnyDiff", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		result := d.RawMapStringAnyDiff()

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RawMapStringAnyDiff -- 1 item", actual)
	})
}

func Test_Seg5_HMD_Serialize(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_Serialize", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		b, err := d.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff Serialize -- success", actual)
	})
}

func Test_Seg5_HMD_Deserialize(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_Deserialize", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		var dest map[string]string
		err := d.Deserialize(&dest)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "HashmapDiff Deserialize -- success", actual)
	})
}
