package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ===== Hashmap =====

func Test_C36_Hashmap_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_IsEmpty", func() {
		h := corestr.New.Hashmap.Empty()
		if !h.IsEmpty() {
			t.Error("expected empty")
		}
	})
}

func Test_C36_Hashmap_HasItems(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_HasItems", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		if !h.HasItems() {
			t.Error("expected has items")
		}
	})
}

func Test_C36_Hashmap_Collection(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Collection", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		c := h.Collection()
		if c.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_IsEmptyLock", func() {
		h := corestr.New.Hashmap.Empty()
		if !h.IsEmptyLock() {
			t.Error("expected empty")
		}
	})
}

func Test_C36_Hashmap_AddOrUpdateWithWgLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateWithWgLock", func() {
		h := corestr.New.Hashmap.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateWithWgLock("k", "v", wg)
		wg.Wait()
		if h.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_AddOrUpdateKeyStrValInt(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateKeyStrValInt", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyStrValInt("k", 42)
		v, _ := h.Get("k")
		if v != "42" {
			t.Error("expected 42")
		}
	})
}

func Test_C36_Hashmap_AddOrUpdateKeyStrValFloat(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateKeyStrValFloat", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyStrValFloat("k", 3.14)
		if h.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateKeyStrValFloat64", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyStrValFloat64("k", 2.71)
		if h.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_AddOrUpdateKeyStrValAny(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateKeyStrValAny", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyStrValAny("k", "hello")
		if h.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_AddOrUpdateKeyValueAny(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateKeyValueAny", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "k", Value: "v"})
		if h.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_AddOrUpdateKeyVal(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateKeyVal", func() {
		h := corestr.New.Hashmap.Empty()
		isNew := h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v"})
		if !isNew {
			t.Error("expected new")
		}
		isNew2 := h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v2"})
		if isNew2 {
			t.Error("expected not new")
		}
	})
}

func Test_C36_Hashmap_AddOrUpdate(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdate", func() {
		h := corestr.New.Hashmap.Empty()
		isNew := h.AddOrUpdate("k", "v")
		if !isNew {
			t.Error("expected new")
		}
	})
}

func Test_C36_Hashmap_Set(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Set", func() {
		h := corestr.New.Hashmap.Empty()
		h.Set("k", "v")
		if h.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_SetTrim(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_SetTrim", func() {
		h := corestr.New.Hashmap.Empty()
		h.SetTrim(" k ", " v ")
		_, found := h.Get("k")
		if !found {
			t.Error("expected found")
		}
	})
}

func Test_C36_Hashmap_SetBySplitter(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_SetBySplitter", func() {
		h := corestr.New.Hashmap.Empty()
		h.SetBySplitter("=", "key=value")
		v, _ := h.Get("key")
		if v != "value" {
			t.Error("expected value")
		}
		// single item
		h.SetBySplitter("=", "onlykey")
		v2, _ := h.Get("onlykey")
		if v2 != "" {
			t.Error("expected empty value")
		}
	})
}

func Test_C36_Hashmap_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateStringsPtrWgLock", func() {
		h := corestr.New.Hashmap.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateStringsPtrWgLock(wg, []string{"a", "b"}, []string{"1", "2"})
		wg.Wait()
		if h.Length() != 2 {
			t.Error("expected 2")
		}
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
		if h.Length() != 2 {
			t.Error("expected 2")
		}
		h.AddOrUpdateHashmap(nil)
	})
}

func Test_C36_Hashmap_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateMap", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateMap(map[string]string{"a": "1"})
		if h.Length() != 1 {
			t.Error("expected 1")
		}
		h.AddOrUpdateMap(nil)
	})
}

func Test_C36_Hashmap_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddsOrUpdates", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddsOrUpdates(corestr.KeyValuePair{Key: "a", Value: "1"})
		if h.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_AddOrUpdateKeyAnyValues(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateKeyAnyValues", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyAnyValues(corestr.KeyAnyValuePair{Key: "k", Value: "v"})
		if h.Length() != 1 {
			t.Error("expected 1")
		}
		h.AddOrUpdateKeyAnyValues()
	})
}

func Test_C36_Hashmap_AddOrUpdateKeyValues(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateKeyValues", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyValues(corestr.KeyValuePair{Key: "k", Value: "v"})
		if h.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateCollection", func() {
		h := corestr.New.Hashmap.Empty()
		keys := corestr.New.Collection.Strings([]string{"a"})
		vals := corestr.New.Collection.Strings([]string{"1"})
		h.AddOrUpdateCollection(keys, vals)
		if h.Length() != 1 {
			t.Error("expected 1")
		}
		// nil/empty
		h.AddOrUpdateCollection(nil, nil)
		// mismatch
		keys2 := corestr.New.Collection.Strings([]string{"a", "b"})
		vals2 := corestr.New.Collection.Strings([]string{"1"})
		h2 := corestr.New.Hashmap.Empty()
		h2.AddOrUpdateCollection(keys2, vals2)
		if h2.Length() != 0 {
			t.Error("expected 0 for mismatch")
		}
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
		if h.Length() != 1 {
			t.Error("expected 1")
		}
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
		if h.Length() != 1 {
			t.Error("expected 1 due to break")
		}
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
		if h.Length() != 1 {
			t.Error("expected 1")
		}
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
		if h.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ConcatNew", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		other := corestr.New.Hashmap.Empty()
		other.AddOrUpdate("b", "2")
		newH := h.ConcatNew(false, other)
		if newH.Length() < 2 {
			t.Error("expected >= 2")
		}
		// no args
		clone := h.ConcatNew(true)
		if clone.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_ConcatNewUsingMaps(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ConcatNewUsingMaps", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		newH := h.ConcatNewUsingMaps(false, map[string]string{"b": "2"})
		if newH.Length() < 2 {
			t.Error("expected >= 2")
		}
		// empty
		clone := h.ConcatNewUsingMaps(true)
		if clone.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_AddOrUpdateLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AddOrUpdateLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateLock("k", "v")
		if h.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_Has_Contains_IsKeyMissing(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Has_Contains_IsKeyMissing", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		if !h.Has("k") {
			t.Error("expected has")
		}
		if !h.Contains("k") {
			t.Error("expected contains")
		}
		if !h.ContainsLock("k") {
			t.Error("expected contains lock")
		}
		if h.IsKeyMissing("k") {
			t.Error("expected not missing")
		}
		if !h.IsKeyMissing("z") {
			t.Error("expected missing")
		}
		if !h.IsKeyMissingLock("z") {
			t.Error("expected missing lock")
		}
		if !h.HasLock("k") {
			t.Error("expected has lock")
		}
	})
}

func Test_C36_Hashmap_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_HasAllStrings", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		if !h.HasAllStrings("a", "b") {
			t.Error("expected has all")
		}
		if h.HasAllStrings("a", "z") {
			t.Error("expected false")
		}
	})
}

func Test_C36_Hashmap_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_HasAllCollectionItems", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		c := corestr.New.Collection.Strings([]string{"a"})
		if !h.HasAllCollectionItems(c) {
			t.Error("expected true")
		}
		if h.HasAllCollectionItems(nil) {
			t.Error("expected false for nil")
		}
	})
}

func Test_C36_Hashmap_HasAll(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_HasAll", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		if !h.HasAll("a") {
			t.Error("expected true")
		}
	})
}

func Test_C36_Hashmap_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_HasAnyItem", func() {
		h := corestr.New.Hashmap.Empty()
		if h.HasAnyItem() {
			t.Error("expected false")
		}
		h.AddOrUpdate("a", "1")
		if !h.HasAnyItem() {
			t.Error("expected true")
		}
	})
}

func Test_C36_Hashmap_HasAny(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_HasAny", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		if !h.HasAny("a", "z") {
			t.Error("expected true")
		}
		if h.HasAny("z") {
			t.Error("expected false")
		}
	})
}

func Test_C36_Hashmap_HasWithLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_HasWithLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		if !h.HasWithLock("a") {
			t.Error("expected true")
		}
	})
}

func Test_C36_Hashmap_GetKeysFilteredItems(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_GetKeysFilteredItems", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		items := h.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if len(items) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_GetKeysFilteredCollection(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_GetKeysFilteredCollection", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		c := h.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if c.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_Items_SafeItems(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Items_SafeItems", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		if len(h.Items()) != 1 {
			t.Error("expected 1")
		}
		if len(h.SafeItems()) != 1 {
			t.Error("expected 1")
		}
		var nilH *corestr.Hashmap
		if nilH.SafeItems() != nil {
			t.Error("expected nil")
		}
	})
}

func Test_C36_Hashmap_ItemsCopyLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ItemsCopyLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		cp := h.ItemsCopyLock()
		if len(*cp) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_ValuesCollection(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ValuesCollection", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		if h.ValuesCollection().Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_ValuesHashset(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ValuesHashset", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		if h.ValuesHashset().Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_ValuesCollectionLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ValuesCollectionLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		if h.ValuesCollectionLock().Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_ValuesHashsetLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ValuesHashsetLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		if h.ValuesHashsetLock().Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_ValuesList(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ValuesList", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		if len(h.ValuesList()) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_KeysValuesCollection(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_KeysValuesCollection", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		keys, values := h.KeysValuesCollection()
		if keys.Length() != 1 || values.Length() != 1 {
			t.Error("expected 1 each")
		}
	})
}

func Test_C36_Hashmap_KeysValuesList(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_KeysValuesList", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		keys, values := h.KeysValuesList()
		if len(keys) != 1 || len(values) != 1 {
			t.Error("expected 1 each")
		}
	})
}

func Test_C36_Hashmap_KeysValuePairs(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_KeysValuePairs", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		pairs := h.KeysValuePairs()
		if len(pairs) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_KeysValuePairsCollection(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_KeysValuePairsCollection", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		c := h.KeysValuePairsCollection()
		if c.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_KeysValuesListLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_KeysValuesListLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		keys, values := h.KeysValuesListLock()
		if len(keys) != 1 || len(values) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_AllKeys_Keys_KeysCollection_KeysLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AllKeys_Keys_KeysCollection_KeysLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		if len(h.AllKeys()) != 1 {
			t.Error("expected 1")
		}
		if len(h.Keys()) != 1 {
			t.Error("expected 1")
		}
		if h.KeysCollection().Length() != 1 {
			t.Error("expected 1")
		}
		if len(h.KeysLock()) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_ValuesListCopyLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ValuesListCopyLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		if len(h.ValuesListCopyLock()) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_KeysToLower(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_KeysToLower", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("KEY", "v")
		lower := h.KeysToLower()
		_, found := lower.Get("key")
		if !found {
			t.Error("expected lowercase key")
		}
	})
}
func Test_C36_Hashmap_Length_LengthLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Length_LengthLock", func() {
		h := corestr.New.Hashmap.Empty()
		if h.Length() != 0 {
			t.Error("expected 0")
		}
		if h.LengthLock() != 0 {
			t.Error("expected 0")
		}
		var nilH *corestr.Hashmap
		if nilH.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C36_Hashmap_IsEqual_IsEqualPtr(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_IsEqual_IsEqualPtr", func() {
		a := corestr.New.Hashmap.Empty()
		a.AddOrUpdate("k", "v")
		b := corestr.New.Hashmap.Empty()
		b.AddOrUpdate("k", "v")
		if !a.IsEqualPtr(b) {
			t.Error("expected equal")
		}
		if !a.IsEqual(*b) {
			t.Error("expected equal")
		}
		if !a.IsEqualPtrLock(b) {
			t.Error("expected equal lock")
		}
		// same ptr
		if !a.IsEqualPtr(a) {
			t.Error("same ptr equal")
		}
		// both empty
		e1 := corestr.New.Hashmap.Empty()
		e2 := corestr.New.Hashmap.Empty()
		if !e1.IsEqualPtr(e2) {
			t.Error("both empty equal")
		}
		// diff values
		c := corestr.New.Hashmap.Empty()
		c.AddOrUpdate("k", "other")
		if a.IsEqualPtr(c) {
			t.Error("expected not equal")
		}
		// nil cases
		var nilH *corestr.Hashmap
		if !nilH.IsEqualPtr(nil) {
			t.Error("both nil equal")
		}
		if nilH.IsEqualPtr(a) {
			t.Error("nil vs non-nil")
		}
	})
}

func Test_C36_Hashmap_Remove_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Remove_RemoveWithLock", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.Remove("a")
		if h.Length() != 0 {
			t.Error("expected 0")
		}
		h.AddOrUpdate("b", "2")
		h.RemoveWithLock("b")
		if h.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C36_Hashmap_String_StringLock(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_String_StringLock", func() {
		h := corestr.New.Hashmap.Empty()
		if h.String() == "" {
			t.Error("expected non-empty for empty")
		}
		h.AddOrUpdate("a", "1")
		if h.String() == "" {
			t.Error("expected non-empty")
		}
		if h.StringLock() == "" {
			t.Error("expected non-empty lock")
		}
	})
}

func Test_C36_Hashmap_GetValuesExceptKeysInHashset(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_GetValuesExceptKeysInHashset", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := h.GetValuesExceptKeysInHashset(hs)
		if len(result) != 1 {
			t.Error("expected 1")
		}
		all := h.GetValuesExceptKeysInHashset(nil)
		if len(all) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C36_Hashmap_GetValuesKeysExcept(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_GetValuesKeysExcept", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		result := h.GetValuesKeysExcept([]string{"a"})
		if len(result) != 0 {
			t.Error("expected 0")
		}
		all := h.GetValuesKeysExcept(nil)
		if len(all) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_GetAllExceptCollection", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		c := corestr.New.Collection.Strings([]string{"a"})
		result := h.GetAllExceptCollection(c)
		if len(result) != 0 {
			t.Error("expected 0")
		}
		all := h.GetAllExceptCollection(nil)
		if len(all) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_Join_JoinKeys(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Join_JoinKeys", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		if h.Join(",") == "" {
			t.Error("expected non-empty")
		}
		if h.JoinKeys(",") == "" {
			t.Error("expected non-empty")
		}
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
		if h.JsonModel() == nil {
			t.Error("expected non-nil")
		}
		if h.JsonModelAny() == nil {
			t.Error("expected non-nil")
		}
		data, err := h.MarshalJSON()
		if err != nil {
			t.Error("expected no error")
		}
		h2 := &corestr.Hashmap{}
		if h2.UnmarshalJSON(data) != nil {
			t.Error("expected no error")
		}
		if h2.UnmarshalJSON([]byte("invalid")) == nil {
			t.Error("expected error")
		}
	})
}

func Test_C36_Hashmap_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Json_JsonPtr", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		r := h.Json()
		if r.HasError() {
			t.Error("expected no error")
		}
		rp := h.JsonPtr()
		if rp == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C36_Hashmap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ParseInjectUsingJson", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		jr := h.JsonPtr()
		h2 := &corestr.Hashmap{}
		result, err := h2.ParseInjectUsingJson(jr)
		if err != nil || result.Length() == 0 {
			t.Error("expected success")
		}
		badJson := corejson.NewPtr("invalid{")
		_, err2 := h2.ParseInjectUsingJson(badJson)
		if err2 == nil {
			t.Error("expected error")
		}
	})
}
func Test_C36_Hashmap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ParseInjectUsingJsonMust", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		jr := h.JsonPtr()
		h2 := &corestr.Hashmap{}
		result := h2.ParseInjectUsingJsonMust(jr)
		if result.Length() == 0 {
			t.Error("expected non-empty")
		}
	})
}

func Test_C36_Hashmap_ToError_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_ToError_ToDefaultError", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		if h.ToError(",") == nil {
			t.Error("expected error")
		}
		if h.ToDefaultError() == nil {
			t.Error("expected error")
		}
	})
}

func Test_C36_Hashmap_KeyValStringLines(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_KeyValStringLines", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		lines := h.KeyValStringLines()
		if len(lines) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Clear_Dispose", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		h.Clear()
		if h.Length() != 0 {
			t.Error("expected 0")
		}
		var nilH *corestr.Hashmap
		if nilH.Clear() != nil {
			t.Error("nil clear returns nil")
		}
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
		if len(lines) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_Hashmap_AsJsoner_JsonParseSelfInject_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_AsJsoner_JsonParseSelfInject_AsJsonContractsBinder", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		if h.AsJsoner() == nil {
			t.Error("expected non-nil")
		}
		jr := h.JsonPtr()
		h2 := &corestr.Hashmap{}
		if h2.JsonParseSelfInject(jr) != nil {
			t.Error("expected no error")
		}
		if h.AsJsonContractsBinder() == nil {
			t.Error("expected non-nil")
		}
		if h.AsJsonParseSelfInjector() == nil {
			t.Error("expected non-nil")
		}
		if h.AsJsonMarshaller() == nil {
			t.Error("expected non-nil")
		}
	})
}
func Test_C36_Hashmap_Get_GetValue(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Get_GetValue", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		v, found := h.Get("k")
		if !found || v != "v" {
			t.Error("expected v")
		}
		v2, found2 := h.GetValue("k")
		if !found2 || v2 != "v" {
			t.Error("expected v")
		}
	})
}

func Test_C36_Hashmap_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_C36_Hashmap_Serialize_Deserialize", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k", "v")
		data, err := h.Serialize()
		if err != nil || len(data) == 0 {
			t.Error("expected success")
		}
		var target map[string]string
		if h.Deserialize(&target) != nil {
			t.Error("expected success")
		}
	})
}

	// ===== newHashmapCreator =====

func Test_C36_NewHashmap_KeyAnyValues(t *testing.T) {
	safeTest(t, "Test_C36_NewHashmap_KeyAnyValues", func() {
		h := corestr.New.Hashmap.KeyAnyValues(corestr.KeyAnyValuePair{Key: "k", Value: "v"})
		if h.Length() != 1 {
			t.Error("expected 1")
		}
		empty := corestr.New.Hashmap.KeyAnyValues()
		if empty == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C36_NewHashmap_KeyValues(t *testing.T) {
	safeTest(t, "Test_C36_NewHashmap_KeyValues", func() {
		h := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k", Value: "v"})
		if h.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_NewHashmap_KeyValuesCollection(t *testing.T) {
	safeTest(t, "Test_C36_NewHashmap_KeyValuesCollection", func() {
		keys := corestr.New.Collection.Strings([]string{"a"})
		vals := corestr.New.Collection.Strings([]string{"1"})
		h := corestr.New.Hashmap.KeyValuesCollection(keys, vals)
		if h.Length() != 1 {
			t.Error("expected 1")
		}
		e := corestr.New.Hashmap.KeyValuesCollection(nil, nil)
		if e.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C36_NewHashmap_KeyValuesStrings(t *testing.T) {
	safeTest(t, "Test_C36_NewHashmap_KeyValuesStrings", func() {
		h := corestr.New.Hashmap.KeyValuesStrings([]string{"a"}, []string{"1"})
		if h.Length() != 1 {
			t.Error("expected 1")
		}
		e := corestr.New.Hashmap.KeyValuesStrings(nil, nil)
		if e.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C36_NewHashmap_UsingMap(t *testing.T) {
	safeTest(t, "Test_C36_NewHashmap_UsingMap", func() {
		h := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		if h.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_NewHashmap_UsingMapOptions(t *testing.T) {
	safeTest(t, "Test_C36_NewHashmap_UsingMapOptions", func() {
		h := corestr.New.Hashmap.UsingMapOptions(true, 5, map[string]string{"a": "1"})
		if h.Length() != 1 {
			t.Error("expected 1")
		}
		e := corestr.New.Hashmap.UsingMapOptions(false, 0, nil)
		if e.Length() != 0 {
			t.Error("expected 0")
		}
		noClone := corestr.New.Hashmap.UsingMapOptions(false, 0, map[string]string{"a": "1"})
		if noClone.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_NewHashmap_MapWithCap(t *testing.T) {
	safeTest(t, "Test_C36_NewHashmap_MapWithCap", func() {
		h := corestr.New.Hashmap.MapWithCap(5, map[string]string{"a": "1"})
		if h.Length() != 1 {
			t.Error("expected 1")
		}
		e := corestr.New.Hashmap.MapWithCap(5, nil)
		if e.Length() != 0 {
			t.Error("expected 0")
		}
		noCap := corestr.New.Hashmap.MapWithCap(0, map[string]string{"a": "1"})
		if noCap.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

	// ===== HashmapDiff =====

func Test_C36_HashmapDiff_Length(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_Length", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		if d.Length() != 1 {
			t.Error("expected 1")
		}
		var nilD *corestr.HashmapDiff
		if nilD.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C36_HashmapDiff_IsEmpty_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_IsEmpty_HasAnyItem", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		if d.IsEmpty() {
			t.Error("expected not empty")
		}
		if !d.HasAnyItem() {
			t.Error("expected has any item")
		}
	})
}

func Test_C36_HashmapDiff_LastIndex(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_LastIndex", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		if d.LastIndex() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C36_HashmapDiff_AllKeysSorted(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_AllKeysSorted", func() {
		d := corestr.HashmapDiff(map[string]string{"b": "2", "a": "1"})
		keys := d.AllKeysSorted()
		if len(keys) != 2 || keys[0] != "a" {
			t.Error("expected sorted keys")
		}
	})
}

func Test_C36_HashmapDiff_MapAnyItems(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_MapAnyItems", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		m := d.MapAnyItems()
		if len(m) != 1 {
			t.Error("expected 1")
		}
		var nilD *corestr.HashmapDiff
		m2 := nilD.MapAnyItems()
		if len(m2) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C36_HashmapDiff_HasAnyChanges(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_HasAnyChanges", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		if !d.HasAnyChanges(map[string]string{"a": "2"}) {
			t.Error("expected has changes")
		}
	})
}

func Test_C36_HashmapDiff_IsRawEqual(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_IsRawEqual", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		if !d.IsRawEqual(map[string]string{"a": "1"}) {
			t.Error("expected equal")
		}
	})
}

func Test_C36_HashmapDiff_HashmapDiffUsingRaw(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_HashmapDiffUsingRaw", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		result := d.HashmapDiffUsingRaw(map[string]string{"a": "1"})
		if len(result) != 0 {
			t.Error("expected empty diff")
		}
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
		if msg == "" {
			t.Error("expected non-empty")
		}
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
		if msg == "" {
			t.Error("expected non-empty")
		}
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
		if len(raw) != 1 {
			t.Error("expected 1")
		}
		var nilD *corestr.HashmapDiff
		if len(nilD.Raw()) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C36_HashmapDiff_RawMapStringAnyDiff(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_RawMapStringAnyDiff", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		m := d.RawMapStringAnyDiff()
		if len(m) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_HashmapDiff_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDiff_Serialize_Deserialize", func() {
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		data, err := d.Serialize()
		if err != nil || len(data) == 0 {
			t.Error("expected success")
		}
		var target map[string]string
		if d.Deserialize(&target) != nil {
			t.Error("expected success")
		}
	})
}

	// ===== HashmapDataModel / HashsetDataModel =====

func Test_C36_HashmapDataModel(t *testing.T) {
	safeTest(t, "Test_C36_HashmapDataModel", func() {
		dm := &corestr.HashmapDataModel{Items: map[string]string{"k": "v"}}
		h := corestr.NewHashmapUsingDataModel(dm)
		if h.Length() != 1 {
			t.Error("expected 1")
		}
		dm2 := corestr.NewHashmapsDataModelUsing(h)
		if len(dm2.Items) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_HashsetDataModel(t *testing.T) {
	safeTest(t, "Test_C36_HashsetDataModel", func() {
		dm := &corestr.HashsetDataModel{Items: map[string]bool{"k": true}}
		h := corestr.NewHashsetUsingDataModel(dm)
		if h.Length() != 1 {
			t.Error("expected 1")
		}
		dm2 := corestr.NewHashsetsDataModelUsing(h)
		if len(dm2.Items) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_CharCollectionDataModel(t *testing.T) {
	safeTest(t, "Test_C36_CharCollectionDataModel", func() {
		items := map[byte]*corestr.Collection{
			'a': corestr.New.Collection.Strings([]string{"apple"}),
		}
		dm := &corestr.CharCollectionDataModel{Items: items, EachCollectionCapacity: 10}
		cm := corestr.NewCharCollectionMapUsingDataModel(dm)
		if cm.Length() != 1 {
			t.Error("expected 1")
		}
		dm2 := corestr.NewCharCollectionMapDataModelUsing(cm)
		if len(dm2.Items) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C36_CharHashsetDataModel(t *testing.T) {
	safeTest(t, "Test_C36_CharHashsetDataModel", func() {
		items := map[byte]*corestr.Hashset{
			'a': corestr.New.Hashset.Strings([]string{"apple"}),
		}
		dm := &corestr.CharHashsetDataModel{Items: items, EachHashsetCapacity: 10}
		cm := corestr.NewCharHashsetMapUsingDataModel(dm)
		if cm.Length() != 1 {
			t.Error("expected 1")
		}
		dm2 := corestr.NewCharHashsetMapDataModelUsing(cm)
		if len(dm2.Items) != 1 {
			t.Error("expected 1")
		}
	})
}
