package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ===== Hashmap =====

func Test_C36_Hashmap_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_IsEmpty", func() {
		h := corestr.New.Hashmap.Empty()
		actual := args.Map{"result": h.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C36_Hashmap_HasItems(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_HasItems", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		actual := args.Map{"result": h.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has items", actual)
	})
}

func Test_C36_Hashmap_Collection(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Collection", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		c := h.Collection()
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_IsEmptyLock", func() {
		h := corestr.New.Hashmap.Empty()
		actual := args.Map{"result": h.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C36_Hashmap_AddOrUpdateWithWgLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateWithWgLock", func() {
		h := corestr.New.Hashmap.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateWithWgLock("k", "v", wg)
		wg.Wait()
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_AddOrUpdateKeyStrValInt(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateKeyStrValInt", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyStrValInt("k", 42)
		v, _ := h.Get("k")
		actual := args.Map{"result": v != "42"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
	})
}

func Test_C36_Hashmap_AddOrUpdateKeyStrValFloat(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateKeyStrValFloat", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyStrValFloat("k", 3.14)
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateKeyStrValFloat64", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyStrValFloat64("k", 2.71)
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_AddOrUpdateKeyStrValAny(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateKeyStrValAny", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyStrValAny("k", "hello")
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_AddOrUpdateKeyValueAny(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateKeyValueAny", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "k", Value: "v"})
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_AddOrUpdateKeyVal(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateKeyVal", func() {
		h := corestr.New.Hashmap.Empty()
		isNew := h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v"})
		actual := args.Map{"result": isNew}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
		isNew2 := h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v2"})
		actual := args.Map{"result": isNew2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not new", actual)
	})
}

func Test_C36_Hashmap_AddOrUpdate(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdate", func() {
		h := corestr.New.Hashmap.Empty()
		isNew := h.AddOrUpdate("k", "v")
		actual := args.Map{"result": isNew}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
	})
}

func Test_C36_Hashmap_Set(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Set", func() {
		h := corestr.New.Hashmap.Empty()
		h.Set("k", "v")
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_SetTrim(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_SetTrim", func() {
		h := corestr.New.Hashmap.Empty()
		h.SetTrim(" k ", " v ")
		_, found := h.Get("k")
		actual := args.Map{"result": found}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
	})
}

func Test_C36_Hashmap_SetBySplitter(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_SetBySplitter", func() {
		h := corestr.New.Hashmap.Empty()
		h.SetBySplitter("=", "key=value")
		v, _ := h.Get("key")
		actual := args.Map{"result": v != "value"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected value", actual)
		// single item
		h.SetBySplitter("=", "onlykey")
		v2, _ := h.Get("onlykey")
		actual := args.Map{"result": v2 != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty value", actual)
	})
}

func Test_C36_Hashmap_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateStringsPtrWgLock", func() {
		h := corestr.New.Hashmap.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateStringsPtrWgLock(wg, []string{"a", "b"}, []string{"1", "2"})
		wg.Wait()
		actual := args.Map{"result": h.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C36_Hashmap_AddOrUpdateStringsPtrWgLock_Empty(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateStringsPtrWgLock_Empty", func() {
		h := corestr.New.Hashmap.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateStringsPtrWgLock(wg, []string{}, []string{})
		wg.Wait()
	})
}

func Test_C36_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateHashmap", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		other := corestr.New.Hashmap.Empty()
		other.AddOrUpdate("b", "2")
		h.AddOrUpdateHashmap(other)
		actual := args.Map{"result": h.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		h.AddOrUpdateHashmap(nil)
	})
}

func Test_C36_Hashmap_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateMap", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateMap(map[string]string{"a": "1"})
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		h.AddOrUpdateMap(nil)
	})
}

func Test_C36_Hashmap_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddsOrUpdates", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddsOrUpdates(corestr.KeyValuePair{Key: "a", Value: "1"})
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_AddOrUpdateKeyAnyValues(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateKeyAnyValues", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyAnyValues(corestr.KeyAnyValuePair{Key: "k", Value: "v"})
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		h.AddOrUpdateKeyAnyValues()
	})
}

func Test_C36_Hashmap_AddOrUpdateKeyValues(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateKeyValues", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyValues(corestr.KeyValuePair{Key: "k", Value: "v"})
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateCollection", func() {
		h := corestr.New.Hashmap.Empty()
		keys := corestr.New.Collection.Strings([]string{"a"})
		vals := corestr.New.Collection.Strings([]string{"1"})
		h.AddOrUpdateCollection(keys, vals)
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// nil/empty
		h.AddOrUpdateCollection(nil, nil)
		// mismatch
		keys2 := corestr.New.Collection.Strings([]string{"a", "b"})
		vals2 := corestr.New.Collection.Strings([]string{"1"})
		h2 := corestr.New.Hashmap.Empty()
		h2.AddOrUpdateCollection(keys2, vals2)
		actual := args.Map{"result": h2.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for mismatch", actual)
	})
}

func Test_C36_Hashmap_AddsOrUpdatesAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddsOrUpdatesAnyUsingFilter", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddsOrUpdatesAnyUsingFilter(
			func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
				return "filtered", true, false
			},
			corestr.KeyAnyValuePair{Key: "k", Value: "v"},
		)
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_AddsOrUpdatesAnyUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddsOrUpdatesAnyUsingFilter_Break", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddsOrUpdatesAnyUsingFilter(
			func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
				return "v", true, true
			},
			corestr.KeyAnyValuePair{Key: "k1", Value: "v"},
			corestr.KeyAnyValuePair{Key: "k2", Value: "v"},
		)
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 due to break", actual)
	})
}

func Test_C36_Hashmap_AddsOrUpdatesAnyUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddsOrUpdatesAnyUsingFilterLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddsOrUpdatesAnyUsingFilterLock(
			func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
				return "v", true, false
			},
			corestr.KeyAnyValuePair{Key: "k", Value: "v"},
		)
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_AddsOrUpdatesUsingFilter(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddsOrUpdatesUsingFilter", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddsOrUpdatesUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Value, true, false
			},
			corestr.KeyValuePair{Key: "k", Value: "v"},
		)
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ConcatNew", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		other := corestr.New.Hashmap.Empty()
		other.AddOrUpdate("b", "2")
		newH := h.ConcatNew(false, other)
		actual := args.Map{"result": newH.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected >= 2", actual)
		// no args
		clone := h.ConcatNew(true)
		actual := args.Map{"result": clone.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_ConcatNewUsingMaps(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ConcatNewUsingMaps", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		newH := h.ConcatNewUsingMaps(false, map[string]string{"b": "2"})
		actual := args.Map{"result": newH.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected >= 2", actual)
		// empty
		clone := h.ConcatNewUsingMaps(true)
		actual := args.Map{"result": clone.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_AddOrUpdateLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateLock("k", "v")
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_Has_Contains_IsKeyMissing(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Has_Contains_IsKeyMissing", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		actual := args.Map{"result": h.Has("k")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has", actual)
		actual := args.Map{"result": h.Contains("k")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected contains", actual)
		actual := args.Map{"result": h.ContainsLock("k")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected contains lock", actual)
		actual := args.Map{"result": h.IsKeyMissing("k")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not missing", actual)
		actual := args.Map{"result": h.IsKeyMissing("z")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected missing", actual)
		actual := args.Map{"result": h.IsKeyMissingLock("z")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected missing lock", actual)
		actual := args.Map{"result": h.HasLock("k")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has lock", actual)
	})
}

func Test_C36_Hashmap_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_HasAllStrings", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		actual := args.Map{"result": h.HasAllStrings("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has all", actual)
		actual := args.Map{"result": h.HasAllStrings("a", "z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C36_Hashmap_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_HasAllCollectionItems", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": h.HasAllCollectionItems(c)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": h.HasAllCollectionItems(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_C36_Hashmap_HasAll(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_HasAll", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		actual := args.Map{"result": h.HasAll("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C36_Hashmap_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_HasAnyItem", func() {
		h := corestr.New.Hashmap.Empty()
		actual := args.Map{"result": h.HasAnyItem()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		h.AddOrUpdate("a", "1")
		actual := args.Map{"result": h.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C36_Hashmap_HasAny(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_HasAny", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		actual := args.Map{"result": h.HasAny("a", "z")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": h.HasAny("z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C36_Hashmap_HasWithLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_HasWithLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		actual := args.Map{"result": h.HasWithLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C36_Hashmap_GetKeysFilteredItems(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_GetKeysFilteredItems", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		items := h.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"result": len(items) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_GetKeysFilteredCollection(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_GetKeysFilteredCollection", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		c := h.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_Items_SafeItems(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Items_SafeItems", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		actual := args.Map{"result": len(h.Items()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": len(h.SafeItems()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		var nilH *corestr.Hashmap
		actual := args.Map{"result": nilH.SafeItems() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C36_Hashmap_ItemsCopyLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ItemsCopyLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		cp := h.ItemsCopyLock()
		actual := args.Map{"result": len(*cp) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_ValuesCollection(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ValuesCollection", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		actual := args.Map{"result": h.ValuesCollection().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_ValuesHashset(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ValuesHashset", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		actual := args.Map{"result": h.ValuesHashset().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_ValuesCollectionLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ValuesCollectionLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		actual := args.Map{"result": h.ValuesCollectionLock().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_ValuesHashsetLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ValuesHashsetLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		actual := args.Map{"result": h.ValuesHashsetLock().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_ValuesList(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ValuesList", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		actual := args.Map{"result": len(h.ValuesList()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_KeysValuesCollection(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_KeysValuesCollection", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		keys, values := h.KeysValuesCollection()
		actual := args.Map{"result": keys.Length() != 1 || values.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 each", actual)
	})
}

func Test_C36_Hashmap_KeysValuesList(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_KeysValuesList", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		keys, values := h.KeysValuesList()
		actual := args.Map{"result": len(keys) != 1 || len(values) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 each", actual)
	})
}

func Test_C36_Hashmap_KeysValuePairs(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_KeysValuePairs", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		pairs := h.KeysValuePairs()
		actual := args.Map{"result": len(pairs) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_KeysValuePairsCollection(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_KeysValuePairsCollection", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		c := h.KeysValuePairsCollection()
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_KeysValuesListLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_KeysValuesListLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		keys, values := h.KeysValuesListLock()
		actual := args.Map{"result": len(keys) != 1 || len(values) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_AllKeys_Keys_KeysCollection_KeysLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AllKeys_Keys_KeysCollection_KeysLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		actual := args.Map{"result": len(h.AllKeys()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": len(h.Keys()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": h.KeysCollection().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": len(h.KeysLock()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_ValuesListCopyLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ValuesListCopyLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		actual := args.Map{"result": len(h.ValuesListCopyLock()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_KeysToLower(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_KeysToLower", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("KEY", "v")
		lower := h.KeysToLower()
		_, found := lower.Get("key")
		actual := args.Map{"result": found}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected lowercase key", actual)
	})
}

func Test_C36_Hashmap_ValuesToLower(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ValuesToLower", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("KEY", "v")
		lower := h.ValuesToLower()
		actual := args.Map{"result": lower.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_Length_LengthLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Length_LengthLock", func() {
		h := corestr.New.Hashmap.Empty()
		actual := args.Map{"result": h.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual := args.Map{"result": h.LengthLock() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		var nilH *corestr.Hashmap
		actual := args.Map{"result": nilH.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C36_Hashmap_IsEqual_IsEqualPtr(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_IsEqual_IsEqualPtr", func() {
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("k", "v")
		b := corestr.New.Hashmap.Empty()
		b.AddOrUpdate("k", "v")
		actual := args.Map{"result": a.IsEqualPtr(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		actual := args.Map{"result": a.IsEqual(*b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		actual := args.Map{"result": a.IsEqualPtrLock(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal lock", actual)
		// same ptr
		actual := args.Map{"result": a.IsEqualPtr(a)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "same ptr equal", actual)
		// both empty
		e1 := corestr.New.Hashmap.Empty()
		e2 := corestr.New.Hashmap.Empty()
		actual := args.Map{"result": e1.IsEqualPtr(e2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "both empty equal", actual)
		// diff values
		c := corestr.New.Hashmap.Empty()
		c.AddOrUpdate("k", "other")
		actual := args.Map{"result": a.IsEqualPtr(c)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
		// nil cases
		var nilH *corestr.Hashmap
		actual := args.Map{"result": nilH.IsEqualPtr(nil)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "both nil equal", actual)
		actual := args.Map{"result": nilH.IsEqualPtr(a)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil vs non-nil", actual)
	})
}

func Test_C36_Hashmap_Remove_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Remove_RemoveWithLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.Remove("a")
		actual := args.Map{"result": h.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		h.AddOrUpdate("b", "2")
		h.RemoveWithLock("b")
		actual := args.Map{"result": h.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C36_Hashmap_String_StringLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_String_StringLock", func() {
		h := corestr.New.Hashmap.Empty()
		actual := args.Map{"result": h.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty for empty", actual)
		h.AddOrUpdate("a", "1")
		actual := args.Map{"result": h.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		actual := args.Map{"result": h.StringLock() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty lock", actual)
	})
}

func Test_C36_Hashmap_GetValuesExceptKeysInHashset(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_GetValuesExceptKeysInHashset", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := h.GetValuesExceptKeysInHashset(hs)
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		all := h.GetValuesExceptKeysInHashset(nil)
		actual := args.Map{"result": len(all) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C36_Hashmap_GetValuesKeysExcept(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_GetValuesKeysExcept", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		result := h.GetValuesKeysExcept([]string{"a"})
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		all := h.GetValuesKeysExcept(nil)
		actual := args.Map{"result": len(all) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_GetAllExceptCollection", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		c := corestr.New.Collection.Strings([]string{"a"})
		result := h.GetAllExceptCollection(c)
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		all := h.GetAllExceptCollection(nil)
		actual := args.Map{"result": len(all) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_Join_JoinKeys(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Join_JoinKeys", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		actual := args.Map{"result": h.Join(",") == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		actual := args.Map{"result": h.JoinKeys(",") == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C36_Hashmap_DiffRaw(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_DiffRaw", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		_ = h.DiffRaw(map[string]string{"a": "2"})
	})
}

func Test_C36_Hashmap_Diff(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Diff", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		other := corestr.New.Hashmap.Empty()
		other.AddOrUpdate("a", "2")
		_ = h.Diff(other)
	})
}

func Test_C36_Hashmap_JsonModel_MarshalJSON_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_JsonModel_MarshalJSON_UnmarshalJSON", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		actual := args.Map{"result": h.JsonModel() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		actual := args.Map{"result": h.JsonModelAny() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		data, err := h.MarshalJSON()
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
		h2 := &corestr.Hashmap{}
		actual := args.Map{"result": h2.UnmarshalJSON(data) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
		actual := args.Map{"result": h2.UnmarshalJSON([]byte("invalid")) == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C36_Hashmap_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Json_JsonPtr", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		r := h.Json()
		actual := args.Map{"result": r.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
		rp := h.JsonPtr()
		actual := args.Map{"result": rp == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C36_Hashmap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ParseInjectUsingJson", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		jr := h.JsonPtr()
		h2 := &corestr.Hashmap{}
		result, err := h2.ParseInjectUsingJson(jr)
		actual := args.Map{"result": err != nil || result.Length() == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected success", actual)
		badJson := corejson.NewPtr("invalid{")
		_, err2 := h2.ParseInjectUsingJson(badJson)
		actual := args.Map{"result": err2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}
func Test_C36_Hashmap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ParseInjectUsingJsonMust", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		jr := h.JsonPtr()
		h2 := &corestr.Hashmap{}
		result := h2.ParseInjectUsingJsonMust(jr)
		actual := args.Map{"result": result.Length() == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C36_Hashmap_ToError_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ToError_ToDefaultError", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		actual := args.Map{"result": h.ToError(",") == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
		actual := args.Map{"result": h.ToDefaultError() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C36_Hashmap_KeyValStringLines(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_KeyValStringLines", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		lines := h.KeyValStringLines()
		actual := args.Map{"result": len(lines) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Clear_Dispose", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		h.Clear()
		actual := args.Map{"result": h.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		var nilH *corestr.Hashmap
		actual := args.Map{"result": nilH.Clear() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil clear returns nil", actual)
		h2 := corestr.New.Hashmap.Empty()
		h2.Dispose()
		var nilH2 *corestr.Hashmap
		nilH2.Dispose()
	})
}

func Test_C36_Hashmap_ToStringsUsingCompiler(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ToStringsUsingCompiler", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		lines := h.ToStringsUsingCompiler(func(k, v string) string {
			return k + "=" + v
		})
		actual := args.Map{"result": len(lines) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_Hashmap_AsJsoner_JsonParseSelfInject_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AsJsoner_JsonParseSelfInject_AsJsonContractsBinder", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		actual := args.Map{"result": h.AsJsoner() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		jr := h.JsonPtr()
		h2 := &corestr.Hashmap{}
		actual := args.Map{"result": h2.JsonParseSelfInject(jr) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
		actual := args.Map{"result": h.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		actual := args.Map{"result": h.AsJsonParseSelfInjector() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		actual := args.Map{"result": h.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C36_Hashmap_Clone_ClonePtr(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Clone_ClonePtr", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		cloned := h.Clone()
		actual := args.Map{"result": cloned.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		cp := h.ClonePtr()
		actual := args.Map{"result": cp == nil || cp.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		var nilH *corestr.Hashmap
		actual := args.Map{"result": nilH.ClonePtr() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil clone returns nil", actual)
		// empty clone
		e := corestr.New.Hashmap.Empty()
		ec := e.Clone()
		actual := args.Map{"result": ec.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C36_Hashmap_Get_GetValue(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Get_GetValue", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		v, found := h.Get("k")
		actual := args.Map{"result": found || v != "v"}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected v", actual)
		v2, found2 := h.GetValue("k")
		actual := args.Map{"result": found2 || v2 != "v"}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected v", actual)
	})
}

func Test_C36_Hashmap_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Serialize_Deserialize", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		data, err := h.Serialize()
		actual := args.Map{"result": err != nil || len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected success", actual)
		var target map[string]string
		actual := args.Map{"result": h.Deserialize(&target) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected success", actual)
	})
}

	// ===== newHashmapCreator =====

func Test_C36_NewHashmap_KeyAnyValues(t *testing.T) {
	safeTest(t, "Test_C36_NewHashmap_KeyAnyValues", func() {
		h := corestr.New.Hashmap.KeyAnyValues(corestr.KeyAnyValuePair{Key: "k", Value: "v"})
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		empty := corestr.New.Hashmap.KeyAnyValues()
		actual := args.Map{"result": empty == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C36_NewHashmap_KeyValues(t *testing.T) {
	safeTest(t, "Test_C36_NewHashmap_KeyValues", func() {
		h := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k", Value: "v"})
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_NewHashmap_KeyValuesCollection(t *testing.T) {
	safeTest(t, "Test_C36_NewHashmap_KeyValuesCollection", func() {
		keys := corestr.New.Collection.Strings([]string{"a"})
		vals := corestr.New.Collection.Strings([]string{"1"})
		h := corestr.New.Hashmap.KeyValuesCollection(keys, vals)
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		e := corestr.New.Hashmap.KeyValuesCollection(nil, nil)
		actual := args.Map{"result": e.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C36_NewHashmap_KeyValuesStrings(t *testing.T) {
	safeTest(t, "Test_C36_NewHashmap_KeyValuesStrings", func() {
		h := corestr.New.Hashmap.KeyValuesStrings([]string{"a"}, []string{"1"})
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		e := corestr.New.Hashmap.KeyValuesStrings(nil, nil)
		actual := args.Map{"result": e.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C36_NewHashmap_UsingMap(t *testing.T) {
	safeTest(t, "Test_C36_NewHashmap_UsingMap", func() {
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_NewHashmap_UsingMapOptions(t *testing.T) {
	safeTest(t, "Test_C36_NewHashmap_UsingMapOptions", func() {
		h := corestr.New.Hashmap.UsingMapOptions(true, 5, map[string]string{"a": "1"})
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		e := corestr.New.Hashmap.UsingMapOptions(false, 0, nil)
		actual := args.Map{"result": e.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		noClone := corestr.New.Hashmap.UsingMapOptions(false, 0, map[string]string{"a": "1"})
		actual := args.Map{"result": noClone.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_NewHashmap_MapWithCap(t *testing.T) {
	safeTest(t, "Test_C36_NewHashmap_MapWithCap", func() {
		h := corestr.New.Hashmap.MapWithCap(5, map[string]string{"a": "1"})
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		e := corestr.New.Hashmap.MapWithCap(5, nil)
		actual := args.Map{"result": e.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		noCap := corestr.New.Hashmap.MapWithCap(0, map[string]string{"a": "1"})
		actual := args.Map{"result": noCap.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

	// ===== HashmapDiff =====

func Test_C36_HashmapDiff_Length(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_Length", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		actual := args.Map{"result": d.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		var nilD *corestr.HashmapDiff
		actual := args.Map{"result": nilD.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C36_HashmapDiff_IsEmpty_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_IsEmpty_HasAnyItem", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		actual := args.Map{"result": d.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
		actual := args.Map{"result": d.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has any item", actual)
	})
}

func Test_C36_HashmapDiff_LastIndex(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_LastIndex", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		actual := args.Map{"result": d.LastIndex() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C36_HashmapDiff_AllKeysSorted(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_AllKeysSorted", func() {
		d := corestr.HashmapDiff(map[string]string{"b": "2", "a": "1"})
		keys := d.AllKeysSorted()
		actual := args.Map{"result": len(keys) != 2 || keys[0] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected sorted keys", actual)
	})
}

func Test_C36_HashmapDiff_MapAnyItems(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_MapAnyItems", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		m := d.MapAnyItems()
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		var nilD *corestr.HashmapDiff
		m2 := nilD.MapAnyItems()
		actual := args.Map{"result": len(m2) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C36_HashmapDiff_HasAnyChanges(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_HasAnyChanges", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		actual := args.Map{"result": d.HasAnyChanges(map[string]string{"a": "2"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has changes", actual)
	})
}

func Test_C36_HashmapDiff_IsRawEqual(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_IsRawEqual", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		actual := args.Map{"result": d.IsRawEqual(map[string]string{"a": "1"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_C36_HashmapDiff_HashmapDiffUsingRaw(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_HashmapDiffUsingRaw", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		result := d.HashmapDiffUsingRaw(map[string]string{"a": "1"})
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty diff", actual)
	})
}

func Test_C36_HashmapDiff_DiffRaw(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_DiffRaw", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		_ = d.DiffRaw(map[string]string{"a": "2"})
	})
}

func Test_C36_HashmapDiff_DiffJsonMessage(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_DiffJsonMessage", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		msg := d.DiffJsonMessage(map[string]string{"a": "2"})
		actual := args.Map{"result": msg == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C36_HashmapDiff_ToStringsSliceOfDiffMap(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_ToStringsSliceOfDiffMap", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		diffMap := map[string]string{"a": "changed"}
		slice := d.ToStringsSliceOfDiffMap(diffMap)
		_ = slice
	})
}

func Test_C36_HashmapDiff_ShouldDiffMessage(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_ShouldDiffMessage", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		msg := d.ShouldDiffMessage("test", map[string]string{"a": "2"})
		actual := args.Map{"result": msg == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C36_HashmapDiff_LogShouldDiffMessage(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_LogShouldDiffMessage", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		msg := d.LogShouldDiffMessage("test", map[string]string{"a": "2"})
		_ = msg
	})
}

func Test_C36_HashmapDiff_Raw(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_Raw", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		raw := d.Raw()
		actual := args.Map{"result": len(raw) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		var nilD *corestr.HashmapDiff
		actual := args.Map{"result": len(nilD.Raw()) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C36_HashmapDiff_RawMapStringAnyDiff(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_RawMapStringAnyDiff", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		m := d.RawMapStringAnyDiff()
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_HashmapDiff_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_Serialize_Deserialize", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		data, err := d.Serialize()
		actual := args.Map{"result": err != nil || len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected success", actual)
		var target map[string]string
		actual := args.Map{"result": d.Deserialize(&target) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected success", actual)
	})
}

	// ===== HashmapDataModel / HashsetDataModel =====

func Test_C36_HashmapDataModel(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDataModel", func() {
		dm := &corestr.HashmapDataModel{Items: map[string]string{"k": "v"}}
		h := corestr.NewHashmapUsingDataModel(dm)
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		dm2 := corestr.NewHashmapsDataModelUsing(h)
		actual := args.Map{"result": len(dm2.Items) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_HashsetDataModel(t *testing.T) {
	safeTest(t, "Test_C36_HashsetDataModel", func() {
		dm := &corestr.HashsetDataModel{Items: map[string]bool{"k": true}}
		h := corestr.NewHashsetUsingDataModel(dm)
		actual := args.Map{"result": h.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		dm2 := corestr.NewHashsetsDataModelUsing(h)
		actual := args.Map{"result": len(dm2.Items) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_CharCollectionDataModel(t *testing.T) {
	safeTest(t, "Test_C36_CharCollectionDataModel", func() {
		items := map[byte]*corestr.Collection{
			'a': corestr.New.Collection.Strings([]string{"apple"}),
		}
		dm := &corestr.CharCollectionDataModel{Items: items, EachCollectionCapacity: 10}
		cm := corestr.NewCharCollectionMapUsingDataModel(dm)
		actual := args.Map{"result": cm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		dm2 := corestr.NewCharCollectionMapDataModelUsing(cm)
		actual := args.Map{"result": len(dm2.Items) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C36_CharHashsetDataModel(t *testing.T) {
	safeTest(t, "Test_C36_CharHashsetDataModel", func() {
		items := map[byte]*corestr.Hashset{
			'a': corestr.New.Hashset.Strings([]string{"apple"}),
		}
		dm := &corestr.CharHashsetDataModel{Items: items, EachHashsetCapacity: 10}
		cm := corestr.NewCharHashsetMapUsingDataModel(dm)
		actual := args.Map{"result": cm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		dm2 := corestr.NewCharHashsetMapDataModelUsing(cm)
		actual := args.Map{"result": len(dm2.Items) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}
