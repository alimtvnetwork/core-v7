package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ═══ Hashmap ═══

func Test_C38_Hashmap_CRUD(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_CRUD", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("a", "1")
		h.Set("b", "2")
		h.SetTrim("  c  ", "  3  ")
		if h.Length() != 3 { t.Fatal("expected 3") }
		if !h.Has("a") { t.Fatal() }
		if !h.Contains("b") { t.Fatal() }
		if h.IsKeyMissing("a") { t.Fatal() }
		if !h.IsKeyMissing("z") { t.Fatal() }
	})
}

func Test_C38_Hashmap_SetBySplitter(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_SetBySplitter", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.SetBySplitter("=", "key=val")
		if v, ok := h.Get("key"); !ok || v != "val" { t.Fatal() }
		h.SetBySplitter("=", "noequals")
		if v, ok := h.Get("noequals"); !ok || v != "" { t.Fatal(v) }
	})
}

func Test_C38_Hashmap_AddOrUpdateKeyStrValInt(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateKeyStrValInt", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyStrValInt("k", 42)
		v, _ := h.Get("k")
		if v != "42" { t.Fatal("expected 42") }
	})
}

func Test_C38_Hashmap_AddOrUpdateKeyStrValFloat(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateKeyStrValFloat", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyStrValFloat("k", 3.14)
		v, _ := h.Get("k")
		if v == "" { t.Fatal("expected non-empty") }
	})
}

func Test_C38_Hashmap_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateKeyStrValFloat64", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyStrValFloat64("k", 2.71)
		v, _ := h.Get("k")
		if v == "" { t.Fatal() }
	})
}

func Test_C38_Hashmap_AddOrUpdateKeyStrValAny(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateKeyStrValAny", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyStrValAny("k", []int{1, 2})
		if h.Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_AddOrUpdateKeyValueAny(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateKeyValueAny", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "k", Value: 42})
		if h.Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_AddOrUpdateKeyVal(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateKeyVal", func() {
		h := corestr.New.Hashmap.Cap(2)
		isNew := h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v"})
		if !isNew { t.Fatal("expected new") }
		isNew2 := h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v2"})
		if isNew2 { t.Fatal("expected not new") }
	})
}

func Test_C38_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateHashmap", func() {
		h1 := corestr.New.Hashmap.Cap(2)
		h1.AddOrUpdate("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.AddOrUpdate("b", "2")
		h1.AddOrUpdateHashmap(h2)
		if h1.Length() != 2 { t.Fatal() }
		h1.AddOrUpdateHashmap(nil)
	})
}

func Test_C38_Hashmap_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateMap", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateMap(map[string]string{"a": "1"})
		if h.Length() != 1 { t.Fatal() }
		h.AddOrUpdateMap(nil)
	})
}

func Test_C38_Hashmap_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddsOrUpdates", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddsOrUpdates(corestr.KeyValuePair{Key: "a", Value: "1"})
		if h.Length() != 1 { t.Fatal() }
		h.AddsOrUpdates(nil...)
	})
}

func Test_C38_Hashmap_AddOrUpdateKeyAnyValues(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateKeyAnyValues", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyAnyValues(corestr.KeyAnyValuePair{Key: "k", Value: 1})
		if h.Length() != 1 { t.Fatal() }
		h.AddOrUpdateKeyAnyValues()
	})
}

func Test_C38_Hashmap_AddOrUpdateKeyValues(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateKeyValues", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateKeyValues(corestr.KeyValuePair{Key: "k", Value: "v"})
		if h.Length() != 1 { t.Fatal() }
		h.AddOrUpdateKeyValues()
	})
}

func Test_C38_Hashmap_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateCollection", func() {
		h := corestr.New.Hashmap.Cap(2)
		keys := corestr.New.Collection.Strings([]string{"a", "b"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})
		h.AddOrUpdateCollection(keys, vals)
		if h.Length() != 2 { t.Fatal() }
		// nil
		h.AddOrUpdateCollection(nil, nil)
		// mismatch
		h.AddOrUpdateCollection(keys, corestr.New.Collection.Strings([]string{"1"}))
	})
}

func Test_C38_Hashmap_AddOrUpdateWithWgLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateWithWgLock", func() {
		h := corestr.New.Hashmap.Cap(2)
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateWithWgLock("a", "1", &wg)
		wg.Wait()
		if h.Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateStringsPtrWgLock", func() {
		h := corestr.New.Hashmap.Cap(2)
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateStringsPtrWgLock(&wg, []string{"a"}, []string{"1"})
		wg.Wait()
		if h.Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_AddOrUpdateLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddOrUpdateLock", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateLock("a", "1")
		if h.Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_Lock_Methods(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Lock_Methods", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		if h.IsEmptyLock() { t.Fatal() }
		if !h.ContainsLock("a") { t.Fatal() }
		if !h.HasLock("a") { t.Fatal() }
		if h.IsKeyMissingLock("a") { t.Fatal() }
		if !h.HasWithLock("a") { t.Fatal() }
		if h.LengthLock() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_HasAllStrings", func() {
		h := corestr.New.Hashmap.Cap(3)
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		if !h.HasAllStrings("a", "b") { t.Fatal() }
		if h.HasAllStrings("a", "z") { t.Fatal() }
	})
}

func Test_C38_Hashmap_HasAll(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_HasAll", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		if !h.HasAll("a") { t.Fatal() }
		if h.HasAll("z") { t.Fatal() }
	})
}

func Test_C38_Hashmap_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_HasAnyItem", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		if !h.HasAnyItem() { t.Fatal() }
	})
}

func Test_C38_Hashmap_HasAny(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_HasAny", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		if !h.HasAny("z", "a") { t.Fatal() }
		if h.HasAny("x", "y") { t.Fatal() }
	})
}

func Test_C38_Hashmap_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_HasAllCollectionItems", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		c := corestr.New.Collection.Strings([]string{"a"})
		if !h.HasAllCollectionItems(c) { t.Fatal() }
		if h.HasAllCollectionItems(nil) { t.Fatal() }
	})
}

func Test_C38_Hashmap_Diff(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Diff", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		_ = h.DiffRaw(map[string]string{"b": "2"})
	})
}

func Test_C38_Hashmap_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_ConcatNew", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.AddOrUpdate("b", "2")
		cn := h.ConcatNew(true, h2)
		if cn.Length() != 2 { t.Fatal() }
		cn2 := h.ConcatNew(true)
		if cn2.Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_ConcatNewUsingMaps(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_ConcatNewUsingMaps", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		cn := h.ConcatNewUsingMaps(true, map[string]string{"b": "2"})
		if cn.Length() != 2 { t.Fatal() }
		cn2 := h.ConcatNewUsingMaps(true)
		if cn2.Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_GetKeysFilteredItems(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_GetKeysFilteredItems", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		r := h.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) { return s, true, false })
		if len(r) != 1 { t.Fatal() }
		r2 := corestr.Empty.Hashmap().GetKeysFilteredItems(func(s string, i int) (string, bool, bool) { return s, true, false })
		if len(r2) != 0 { t.Fatal() }
	})
}

func Test_C38_Hashmap_GetKeysFilteredCollection(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_GetKeysFilteredCollection", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		fc := h.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, false })
		if fc.Length() != 1 { t.Fatal() }
		fc2 := corestr.Empty.Hashmap().GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, false })
		if fc2.Length() != 0 { t.Fatal() }
	})
}

func Test_C38_Hashmap_AddsOrUpdatesUsingFilter(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddsOrUpdatesUsingFilter", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddsOrUpdatesUsingFilter(
			func(p corestr.KeyValuePair) (string, bool, bool) { return p.Value, true, false },
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)
		if h.Length() != 1 { t.Fatal() }
		h.AddsOrUpdatesUsingFilter(func(p corestr.KeyValuePair) (string, bool, bool) { return "", false, false }, nil...)
	})
}

func Test_C38_Hashmap_AddsOrUpdatesAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddsOrUpdatesAnyUsingFilter", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddsOrUpdatesAnyUsingFilter(
			func(p corestr.KeyAnyValuePair) (string, bool, bool) { return "v", true, false },
			corestr.KeyAnyValuePair{Key: "k", Value: 42},
		)
		if h.Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_AddsOrUpdatesAnyUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_AddsOrUpdatesAnyUsingFilterLock", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddsOrUpdatesAnyUsingFilterLock(
			func(p corestr.KeyAnyValuePair) (string, bool, bool) { return "v", true, false },
			corestr.KeyAnyValuePair{Key: "k", Value: 42},
		)
		if h.Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_Values_Keys(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Values_Keys", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		if len(h.ValuesList()) != 1 { t.Fatal() }
		if len(h.Keys()) != 1 { t.Fatal() }
		if len(h.AllKeys()) != 1 { t.Fatal() }
		if h.ValuesCollection().Length() != 1 { t.Fatal() }
		if h.ValuesHashset().Length() != 1 { t.Fatal() }
		if h.KeysCollection().Length() != 1 { t.Fatal() }
		if len(h.KeysLock()) != 1 { t.Fatal() }
		if len(h.ValuesListCopyLock()) != 1 { t.Fatal() }
		if h.ValuesCollectionLock().Length() != 1 { t.Fatal() }
		if h.ValuesHashsetLock().Length() != 1 { t.Fatal() }
		if h.Collection().Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_KeysValuesList(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_KeysValuesList", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		k, v := h.KeysValuesList()
		if len(k) != 1 || len(v) != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_KeysValuesCollection(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_KeysValuesCollection", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		k, v := h.KeysValuesCollection()
		if k.Length() != 1 || v.Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_KeysValuesListLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_KeysValuesListLock", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		k, v := h.KeysValuesListLock()
		if len(k) != 1 || len(v) != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_KeysValuePairs(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_KeysValuePairs", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		pairs := h.KeysValuePairs()
		if len(pairs) != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_KeysValuePairsCollection(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_KeysValuePairsCollection", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		pc := h.KeysValuePairsCollection()
		if pc.Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_Items_SafeItems_ItemsCopyLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Items_SafeItems_ItemsCopyLock", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		if len(h.Items()) != 1 { t.Fatal() }
		if len(h.SafeItems()) != 1 { t.Fatal() }
		cm := h.ItemsCopyLock()
		if len(*cm) != 1 { t.Fatal() }
		var nilH *corestr.Hashmap
		if nilH.SafeItems() != nil { t.Fatal() }
	})
}

func Test_C38_Hashmap_IsEqual(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_IsEqual", func() {
		h1 := corestr.New.Hashmap.Cap(2)
		h1.AddOrUpdate("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.AddOrUpdate("a", "1")
		if !h1.IsEqualPtr(h2) { t.Fatal() }
		if !h1.IsEqualPtrLock(h2) { t.Fatal() }
		if !h1.IsEqual(*h2) { t.Fatal() }
		// nil checks
		var nilH *corestr.Hashmap
		if !nilH.IsEqualPtr(nil) { t.Fatal() }
		if nilH.IsEqualPtr(h1) { t.Fatal() }
		// same ptr
		if !h1.IsEqualPtr(h1) { t.Fatal() }
		// both empty
		e1, e2 := corestr.Empty.Hashmap(), corestr.Empty.Hashmap()
		if !e1.IsEqualPtr(e2) { t.Fatal() }
		// one empty
		if h1.IsEqualPtr(corestr.Empty.Hashmap()) { t.Fatal() }
		// diff lengths
		h3 := corestr.New.Hashmap.Cap(2)
		h3.AddOrUpdate("a", "1")
		h3.AddOrUpdate("b", "2")
		if h1.IsEqualPtr(h3) { t.Fatal() }
		// diff values
		h4 := corestr.New.Hashmap.Cap(2)
		h4.AddOrUpdate("a", "999")
		if h1.IsEqualPtr(h4) { t.Fatal() }
	})
}

func Test_C38_Hashmap_Remove(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Remove", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		h.Remove("a")
		if h.Length() != 0 { t.Fatal() }
	})
}

func Test_C38_Hashmap_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_RemoveWithLock", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		h.RemoveWithLock("a")
		if h.Length() != 0 { t.Fatal() }
	})
}

func Test_C38_Hashmap_String(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_String", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		if h.String() == "" { t.Fatal() }
		if corestr.Empty.Hashmap().String() == "" { t.Fatal() }
	})
}

func Test_C38_Hashmap_StringLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_StringLock", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		if h.StringLock() == "" { t.Fatal() }
	})
}

func Test_C38_Hashmap_GetValuesExceptKeysInHashset(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_GetValuesExceptKeysInHashset", func() {
		h := corestr.New.Hashmap.Cap(3)
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := h.GetValuesExceptKeysInHashset(hs)
		if len(r) != 1 { t.Fatal() }
		r2 := h.GetValuesExceptKeysInHashset(nil)
		if len(r2) != 2 { t.Fatal() }
	})
}

func Test_C38_Hashmap_GetValuesKeysExcept(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_GetValuesKeysExcept", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		r := h.GetValuesKeysExcept(nil)
		if len(r) != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_GetAllExceptCollection", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		r := h.GetAllExceptCollection(nil)
		if len(r) != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_Join(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Join", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		if h.Join(",") == "" { t.Fatal() }
		if h.JoinKeys(",") == "" { t.Fatal() }
	})
}

func Test_C38_Hashmap_KeysToLower(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_KeysToLower", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("UPPER", "val")
		lower := h.KeysToLower()
		if !lower.Has("upper") { t.Fatal() }
	})
}
func Test_C38_Hashmap_ToError(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_ToError", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		if h.ToError(",") == nil { t.Fatal() }
		if h.ToDefaultError() == nil { t.Fatal() }
	})
}

func Test_C38_Hashmap_KeyValStringLines(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_KeyValStringLines", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		if len(h.KeyValStringLines()) != 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_ToStringsUsingCompiler(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_ToStringsUsingCompiler", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		r := h.ToStringsUsingCompiler(func(k, v string) string { return k + v })
		if len(r) != 1 { t.Fatal() }
		r2 := corestr.Empty.Hashmap().ToStringsUsingCompiler(func(k, v string) string { return "" })
		if len(r2) != 0 { t.Fatal() }
	})
}
func Test_C38_Hashmap_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Clear_Dispose", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		h.Clear()
		if h.Length() != 0 { t.Fatal() }
		var nilH *corestr.Hashmap
		if nilH.Clear() != nil { t.Fatal() }
		h2 := corestr.New.Hashmap.Cap(2)
		h2.Dispose()
		var nilH2 *corestr.Hashmap
		nilH2.Dispose()
	})
}

func Test_C38_Hashmap_JSON(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_JSON", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		j := h.Json()
		if j.HasError() { t.Fatal(j.Error) }
		jp := h.JsonPtr()
		if jp.HasError() { t.Fatal(jp.Error) }
		if h.JsonModelAny() == nil { t.Fatal() }
		b, err := h.MarshalJSON()
		if err != nil { t.Fatal(err) }
		if len(b) == 0 { t.Fatal() }
		h2 := &corestr.Hashmap{}
		err2 := h2.UnmarshalJSON(b)
		if err2 != nil { t.Fatal(err2) }
		// invalid
		err3 := h2.UnmarshalJSON([]byte(`{invalid`))
		if err3 == nil { t.Fatal() }
	})
}

func Test_C38_Hashmap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_ParseInjectUsingJson", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		jr := h.JsonPtr()
		h2 := &corestr.Hashmap{}
		result, err := h2.ParseInjectUsingJson(jr)
		if err != nil { t.Fatal(err) }
		if result.Length() < 1 { t.Fatal() }
	})
}

func Test_C38_Hashmap_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Serialize_Deserialize", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		b, err := h.Serialize()
		if err != nil { t.Fatal(err) }
		if len(b) == 0 { t.Fatal() }
		var target map[string]string
		err2 := h.Deserialize(&target)
		if err2 != nil { t.Fatal(err2) }
	})
}

func Test_C38_Hashmap_Get_GetValue(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_Get_GetValue", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		v, ok := h.Get("a")
		if !ok || v != "1" { t.Fatal() }
		v2, ok2 := h.GetValue("a")
		if !ok2 || v2 != "1" { t.Fatal() }
	})
}

func Test_C38_Hashmap_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_InterfaceCasts", func() {
		h := corestr.New.Hashmap.Cap(2)
		if h.AsJsoner() == nil { t.Fatal() }
		if h.AsJsonContractsBinder() == nil { t.Fatal() }
		if h.AsJsonParseSelfInjector() == nil { t.Fatal() }
		if h.AsJsonMarshaller() == nil { t.Fatal() }
	})
}

func Test_C38_Hashmap_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C38_Hashmap_JsonParseSelfInject", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		jr := h.JsonPtr()
		h2 := &corestr.Hashmap{}
		err := h2.JsonParseSelfInject(jr)
		if err != nil { t.Fatal(err) }
	})
}

	// ── newHashmapCreator ──

func Test_C38_NewHashmapCreator_Methods(t *testing.T) {
	safeTest(t, "Test_C38_NewHashmapCreator_Methods", func() {
		h1 := corestr.New.Hashmap.Empty()
		if h1.Length() != 0 { t.Fatal() }
		h2 := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		if h2.Length() != 1 { t.Fatal() }
		h3 := corestr.New.Hashmap.KeyValues()
		if h3.Length() != 0 { t.Fatal() }
		h4 := corestr.New.Hashmap.KeyAnyValues(corestr.KeyAnyValuePair{Key: "k", Value: 1})
		if h4.Length() != 1 { t.Fatal() }
		h5 := corestr.New.Hashmap.KeyAnyValues()
		if h5.Length() != 0 { t.Fatal() }
		h6 := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		if h6.Length() != 1 { t.Fatal() }
		h7 := corestr.New.Hashmap.UsingMapOptions(true, 5, map[string]string{"a": "1"})
		if h7.Length() != 1 { t.Fatal() }
		h8 := corestr.New.Hashmap.UsingMapOptions(false, 0, nil)
		if h8.Length() != 0 { t.Fatal() }
		h9 := corestr.New.Hashmap.UsingMapOptions(false, 0, map[string]string{"a": "1"})
		if h9.Length() != 1 { t.Fatal() }
		h10 := corestr.New.Hashmap.KeyValuesStrings([]string{"a"}, []string{"1"})
		if h10.Length() != 1 { t.Fatal() }
		h11 := corestr.New.Hashmap.KeyValuesStrings(nil, nil)
		if h11.Length() != 0 { t.Fatal() }
		k := corestr.New.Collection.Strings([]string{"a"})
		v := corestr.New.Collection.Strings([]string{"1"})
		h12 := corestr.New.Hashmap.KeyValuesCollection(k, v)
		if h12.Length() != 1 { t.Fatal() }
		h13 := corestr.New.Hashmap.KeyValuesCollection(nil, nil)
		if h13.Length() != 0 { t.Fatal() }
		h14 := corestr.New.Hashmap.MapWithCap(0, map[string]string{"a": "1"})
		if h14.Length() != 1 { t.Fatal() }
		h15 := corestr.New.Hashmap.MapWithCap(5, map[string]string{"a": "1"})
		if h15.Length() != 1 { t.Fatal() }
		h16 := corestr.New.Hashmap.MapWithCap(5, nil)
		if h16.Length() != 0 { t.Fatal() }
	})
}

	// ═══ Hashset (comprehensive) ═══

func Test_C38_Hashset_CRUD(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_CRUD", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.Add("a")
		hs.AddNonEmpty("")
		hs.AddNonEmpty("b")
		hs.AddNonEmptyWhitespace("  ")
		hs.AddNonEmptyWhitespace("c")
		if hs.Length() != 3 { t.Fatal("expected 3") }
		if !hs.Has("a") { t.Fatal() }
		if !hs.Contains("b") { t.Fatal() }
		if hs.IsMissing("a") { t.Fatal() }
		if !hs.IsMissing("z") { t.Fatal() }
	})
}

func Test_C38_Hashset_AddIf_AddIfMany(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddIf_AddIfMany", func() {
		hs := corestr.New.Hashset.Cap(3)
		hs.AddIf(false, "skip")
		hs.AddIf(true, "keep")
		hs.AddIfMany(false, "x", "y")
		hs.AddIfMany(true, "m", "n")
		if hs.Length() != 3 { t.Fatal() }
	})
}

func Test_C38_Hashset_AddFunc_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddFunc_AddFuncErr", func() {
		hs := corestr.New.Hashset.Cap(2)
		hs.AddFunc(func() string { return "fn" })
		if !hs.Has("fn") { t.Fatal() }
		hs.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
		if !hs.Has("ok") { t.Fatal() }
		hs.AddFuncErr(func() (string, error) { return "", errForTest }, func(e error) {})
	})
}

func Test_C38_Hashset_AddBool(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddBool", func() {
		hs := corestr.New.Hashset.Cap(2)
		existed := hs.AddBool("a")
		if existed { t.Fatal("expected false") }
		existed2 := hs.AddBool("a")
		if !existed2 { t.Fatal("expected true") }
	})
}

func Test_C38_Hashset_Adds_AddStrings(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Adds_AddStrings", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.Adds("a", "b")
		hs.AddStrings([]string{"c"})
		if hs.Length() != 3 { t.Fatal() }
		hs.Adds(nil...)
		hs.AddStrings(nil)
	})
}

func Test_C38_Hashset_AddCollection(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddCollection", func() {
		hs := corestr.New.Hashset.Cap(5)
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs.AddCollection(c)
		if hs.Length() != 2 { t.Fatal() }
		hs.AddCollection(nil)
	})
}

func Test_C38_Hashset_AddCollections(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddCollections", func() {
		hs := corestr.New.Hashset.Cap(5)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		hs.AddCollections(c1, c2, nil)
		if hs.Length() != 2 { t.Fatal() }
	})
}

func Test_C38_Hashset_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddHashsetItems", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs2 := corestr.New.Hashset.StringsSpreadItems("a", "b")
		hs.AddHashsetItems(hs2)
		if hs.Length() != 2 { t.Fatal() }
		hs.AddHashsetItems(nil)
	})
}

func Test_C38_Hashset_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddItemsMap", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddItemsMap(map[string]bool{"a": true, "b": false})
		if hs.Length() != 1 { t.Fatal("expected 1, false items skipped") }
		hs.AddItemsMap(nil)
	})
}

func Test_C38_Hashset_Lock_Methods(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Lock_Methods", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs.AddLock("b")
		if !hs.HasLock("a") { t.Fatal() }
		if hs.IsMissingLock("a") { t.Fatal() }
		if hs.IsEmptyLock() { t.Fatal() }
		if hs.LengthLock() != 2 { t.Fatal() }
		if !hs.HasWithLock("a") { t.Fatal() }
	})
}

func Test_C38_Hashset_AddPtr_AddPtrLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddPtr_AddPtrLock", func() {
		hs := corestr.New.Hashset.Cap(2)
		s := "ptr"
		hs.AddPtr(&s)
		if !hs.Has("ptr") { t.Fatal() }
		s2 := "lock"
		hs.AddPtrLock(&s2)
		if !hs.Has("lock") { t.Fatal() }
	})
}

func Test_C38_Hashset_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddStringsLock", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddStringsLock([]string{"a", "b"})
		if hs.Length() != 2 { t.Fatal() }
		hs.AddStringsLock(nil)
	})
}

func Test_C38_Hashset_AddSimpleSlice(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddSimpleSlice", func() {
		hs := corestr.New.Hashset.Cap(5)
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		hs.AddSimpleSlice(ss)
		if hs.Length() != 2 { t.Fatal() }
	})
}

func Test_C38_Hashset_HasAll_HasAny_IsAllMissing(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_HasAll_HasAny_IsAllMissing", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		if !hs.HasAll("a", "b") { t.Fatal() }
		if hs.HasAll("a", "z") { t.Fatal() }
		if !hs.HasAny("z", "a") { t.Fatal() }
		if hs.HasAny("x", "y") { t.Fatal() }
		if !hs.IsAllMissing("x", "y") { t.Fatal() }
		if hs.IsAllMissing("a") { t.Fatal() }
	})
}

func Test_C38_Hashset_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_HasAllStrings", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		if !hs.HasAllStrings([]string{"a"}) { t.Fatal() }
	})
}

func Test_C38_Hashset_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_HasAllCollectionItems", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		c := corestr.New.Collection.Strings([]string{"a"})
		if !hs.HasAllCollectionItems(c) { t.Fatal() }
		if hs.HasAllCollectionItems(nil) { t.Fatal() }
	})
}

func Test_C38_Hashset_IsEquals(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_IsEquals", func() {
		a := corestr.New.Hashset.StringsSpreadItems("a", "b")
		b := corestr.New.Hashset.StringsSpreadItems("a", "b")
		if !a.IsEquals(b) { t.Fatal() }
		if !a.IsEqualsLock(b) { t.Fatal() }
		if !a.IsEqual(b) { t.Fatal() }
	})
}

func Test_C38_Hashset_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_ConcatNew", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs2 := corestr.New.Hashset.StringsSpreadItems("b")
		cn := hs.ConcatNewHashsets(true, hs2)
		if cn.Length() != 2 { t.Fatal() }
		cn2 := hs.ConcatNewHashsets(true)
		if cn2.Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashset_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_ConcatNewStrings", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		cn := hs.ConcatNewStrings(true, []string{"b"})
		if cn.Length() != 2 { t.Fatal() }
		cn2 := hs.ConcatNewStrings(true)
		if cn2.Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashset_Resize(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Resize", func() {
		hs := corestr.New.Hashset.Cap(2)
		hs.Add("a")
		hs.Resize(100)
		if !hs.Has("a") { t.Fatal() }
		hs.Resize(1) // no-op
	})
}

func Test_C38_Hashset_ResizeLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_ResizeLock", func() {
		hs := corestr.New.Hashset.Cap(2)
		hs.Add("a")
		hs.ResizeLock(100)
		if !hs.Has("a") { t.Fatal() }
	})
}

func Test_C38_Hashset_AddCapacities(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddCapacities", func() {
		hs := corestr.New.Hashset.Cap(2)
		hs.AddCapacities(10, 20)
		hs.AddCapacities()
		hs.AddCapacitiesLock(10)
		hs.AddCapacitiesLock()
	})
}

func Test_C38_Hashset_Filter(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Filter", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("abc", "xyz")
		filtered := hs.Filter(func(s string) bool { return s == "abc" })
		if filtered.Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashset_GetFilteredItems(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_GetFilteredItems", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "bb")
		r := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) { return s, len(s) > 1, false })
		if len(r) != 1 { t.Fatal() }
		r2 := corestr.Empty.Hashset().GetFilteredItems(func(s string, i int) (string, bool, bool) { return s, true, false })
		if len(r2) != 0 { t.Fatal() }
	})
}

func Test_C38_Hashset_GetFilteredCollection(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_GetFilteredCollection", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		fc := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, false })
		if fc.Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashset_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddsUsingFilter", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, false }, "a", "b")
		if hs.Length() != 2 { t.Fatal() }
	})
}

func Test_C38_Hashset_AddsAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_AddsAnyUsingFilter", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsAnyUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, false }, 42, nil)
		if hs.Length() != 1 { t.Fatal() }
	})
}

func Test_C38_Hashset_Remove_SafeRemove(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Remove_SafeRemove", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		hs.Remove("a")
		if hs.Length() != 1 { t.Fatal() }
		hs.SafeRemove("b")
		if hs.Length() != 0 { t.Fatal() }
		hs.SafeRemove("z") // no-op
	})
}

func Test_C38_Hashset_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_RemoveWithLock", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs.RemoveWithLock("a")
		if hs.Length() != 0 { t.Fatal() }
	})
}
func Test_C38_Hashset_Collection_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Collection_SimpleSlice", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		if hs.Collection().Length() != 1 { t.Fatal() }
		if hs.SimpleSlice().Length() != 1 { t.Fatal() }
		if corestr.Empty.Hashset().SimpleSlice().Length() != 0 { t.Fatal() }
	})
}

func Test_C38_Hashset_String_StringLock(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_String_StringLock", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		if hs.String() == "" { t.Fatal() }
		if hs.StringLock() == "" { t.Fatal() }
		if corestr.Empty.Hashset().String() == "" { t.Fatal() }
	})
}

func Test_C38_Hashset_Join_JoinSorted_JoinLine(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Join_JoinSorted_JoinLine", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		_ = hs.Join(",")
		_ = hs.JoinSorted(",")
		_ = hs.JoinLine()
		_ = hs.NonEmptyJoins(",")
		_ = hs.NonWhitespaceJoins(",")
		_ = corestr.Empty.Hashset().JoinSorted(",")
	})
}

func Test_C38_Hashset_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_GetAllExcept", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		r := hs.GetAllExcept([]string{"a"})
		if len(r) != 1 { t.Fatal() }
		r2 := hs.GetAllExcept(nil)
		if len(r2) != 2 { t.Fatal() }
	})
}

func Test_C38_Hashset_GetAllExceptSpread(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_GetAllExceptSpread", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		r := hs.GetAllExceptSpread("a")
		if len(r) != 1 { t.Fatal() }
	})
}

func Test_C38_Hashset_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_GetAllExceptCollection", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		r := hs.GetAllExceptCollection(nil)
		if len(r) != 2 { t.Fatal() }
	})
}

func Test_C38_Hashset_GetAllExceptHashset(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_GetAllExceptHashset", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		r := hs.GetAllExceptHashset(nil)
		if len(r) != 2 { t.Fatal() }
	})
}

func Test_C38_Hashset_MapStringAny(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_MapStringAny", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		m := hs.MapStringAny()
		if len(m) != 1 { t.Fatal() }
		if len(corestr.Empty.Hashset().MapStringAny()) != 0 { t.Fatal() }
		_ = hs.MapStringAnyDiff()
	})
}

func Test_C38_Hashset_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_DistinctDiff", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		r := hs.DistinctDiffLinesRaw("b", "c")
		if len(r) < 1 { t.Fatal() }
		r2 := hs.DistinctDiffLinesRaw()
		if len(r2) != 2 { t.Fatal() }
		e := corestr.Empty.Hashset()
		r3 := e.DistinctDiffLinesRaw("x")
		if len(r3) != 1 { t.Fatal() }
		r4 := e.DistinctDiffLinesRaw()
		if len(r4) != 0 { t.Fatal() }
	})
}

func Test_C38_Hashset_DistinctDiffLines(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_DistinctDiffLines", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := hs.DistinctDiffLines("b")
		if len(r) < 1 { t.Fatal() }
		_ = hs.DistinctDiffHashset(corestr.New.Hashset.StringsSpreadItems("b"))
		r2 := hs.DistinctDiffLines()
		if len(r2) != 1 { t.Fatal() }
		e := corestr.Empty.Hashset()
		r3 := e.DistinctDiffLines("x")
		if len(r3) != 1 { t.Fatal() }
	})
}

func Test_C38_Hashset_ToLowerSet(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_ToLowerSet", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("ABC")
		lower := hs.ToLowerSet()
		if !lower.Has("abc") { t.Fatal() }
	})
}

func Test_C38_Hashset_WrapQuotes(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_WrapQuotes", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		_ = hs.WrapDoubleQuote()
		hs2 := corestr.New.Hashset.StringsSpreadItems("a")
		_ = hs2.WrapSingleQuote()
		hs3 := corestr.New.Hashset.StringsSpreadItems("a")
		_ = hs3.WrapDoubleQuoteIfMissing()
		hs4 := corestr.New.Hashset.StringsSpreadItems("a")
		_ = hs4.WrapSingleQuoteIfMissing()
		// empty transpile
		_ = corestr.Empty.Hashset().WrapDoubleQuote()
	})
}

func Test_C38_Hashset_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Clear_Dispose", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs.Clear()
		if hs.Length() != 0 { t.Fatal() }
		hs.Dispose()
		var nilH *corestr.Hashset
		nilH.Dispose()
	})
}

func Test_C38_Hashset_JSON(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_JSON", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		j := hs.Json()
		if j.HasError() { t.Fatal(j.Error) }
		jp := hs.JsonPtr()
		if jp.HasError() { t.Fatal(jp.Error) }
		if hs.JsonModelAny() == nil { t.Fatal() }
		b, err := hs.MarshalJSON()
		if err != nil { t.Fatal(err) }
		hs2 := &corestr.Hashset{}
		err2 := hs2.UnmarshalJSON(b)
		if err2 != nil { t.Fatal(err2) }
		err3 := hs2.UnmarshalJSON([]byte(`{bad`))
		if err3 == nil { t.Fatal() }
		// empty json model
		_ = corestr.Empty.Hashset().JsonModel()
	})
}
func Test_C38_Hashset_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_ParseInjectUsingJson", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		jr := hs.JsonPtr()
		hs2 := &corestr.Hashset{}
		result, err := hs2.ParseInjectUsingJson(jr)
		if err != nil { t.Fatal(err) }
		if result.Length() < 1 { t.Fatal() }
	})
}

func Test_C38_Hashset_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Serialize_Deserialize", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		b, err := hs.Serialize()
		if err != nil { t.Fatal(err) }
		if len(b) == 0 { t.Fatal() }
		var target map[string]bool
		err2 := hs.Deserialize(&target)
		if err2 != nil { t.Fatal(err2) }
	})
}

func Test_C38_Hashset_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_InterfaceCasts", func() {
		hs := corestr.New.Hashset.Cap(1)
		if hs.AsJsoner() == nil { t.Fatal() }
		if hs.AsJsonContractsBinder() == nil { t.Fatal() }
		if hs.AsJsonParseSelfInjector() == nil { t.Fatal() }
		if hs.AsJsonMarshaller() == nil { t.Fatal() }
		err := hs.JsonParseSelfInject(hs.JsonPtr())
		if err != nil { t.Fatal(err) }
	})
}

func Test_C38_Hashset_Wg_Methods(t *testing.T) {
	safeTest(t, "Test_C38_Hashset_Wg_Methods", func() {
		hs := corestr.New.Hashset.Cap(5)
		wg := sync.WaitGroup{}
		wg.Add(1)
		hs.AddWithWgLock("a", &wg)
		wg.Wait()
		wg2 := sync.WaitGroup{}
		wg2.Add(1)
		hs.AddStringsPtrWgLock([]string{"b"}, &wg2)
		wg2.Wait()
		wg3 := sync.WaitGroup{}
		wg3.Add(1)
		hs.AddHashsetWgLock(corestr.New.Hashset.StringsSpreadItems("c"), &wg3)
		wg3.Wait()
		wg4 := sync.WaitGroup{}
		wg4.Add(1)
		m := map[string]bool{"d": true}
		hs.AddItemsMapWgLock(&m, &wg4)
		wg4.Wait()
		if hs.Length() != 4 { t.Fatal("expected 4") }
	})
}

	// ── newHashsetCreator ──

func Test_C38_NewHashsetCreator_Methods(t *testing.T) {
	safeTest(t, "Test_C38_NewHashsetCreator_Methods", func() {
		h1 := corestr.New.Hashset.Empty()
		if h1.Length() != 0 { t.Fatal() }
		h2 := corestr.New.Hashset.Strings([]string{"a"})
		if h2.Length() != 1 { t.Fatal() }
		h3 := corestr.New.Hashset.Strings(nil)
		if h3.Length() != 0 { t.Fatal() }
		h4 := corestr.New.Hashset.StringsSpreadItems("a", "b")
		if h4.Length() != 2 { t.Fatal() }
		h5 := corestr.New.Hashset.StringsOption(5, true, "a")
		if h5.Length() != 1 { t.Fatal() }
		h6 := corestr.New.Hashset.StringsOption(0, false)
		if h6.Length() != 0 { t.Fatal() }
		h7 := corestr.New.Hashset.StringsOption(5, false)
		if h7.Length() != 0 { t.Fatal() }
		h8 := corestr.New.Hashset.UsingMap(map[string]bool{"a": true})
		if h8.Length() != 1 { t.Fatal() }
		h9 := corestr.New.Hashset.UsingMap(nil)
		if h9.Length() != 0 { t.Fatal() }
		h10 := corestr.New.Hashset.UsingMapOption(5, true, map[string]bool{"a": true})
		if h10.Length() != 1 { t.Fatal() }
		h11 := corestr.New.Hashset.UsingMapOption(0, false, nil)
		if h11.Length() != 0 { t.Fatal() }
		h12 := corestr.New.Hashset.UsingCollection(corestr.New.Collection.Strings([]string{"a"}))
		if h12.Length() != 1 { t.Fatal() }
		h13 := corestr.New.Hashset.UsingCollection(nil)
		if h13.Length() != 0 { t.Fatal() }
		ss := corestr.New.SimpleSlice.Lines("a")
		h14 := corestr.New.Hashset.SimpleSlice(ss)
		if h14.Length() != 1 { t.Fatal() }
		h15 := corestr.New.Hashset.SimpleSlice(corestr.Empty.SimpleSlice())
		if h15.Length() != 0 { t.Fatal() }
	})
}
