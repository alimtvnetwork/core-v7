package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// =============================================================================
// Hashmap — Core operations
// =============================================================================

func Test_I8_HM01_IsEmpty(t *testing.T) {
	safeTest(t, "Test_I8_HM01_IsEmpty", func() {
		h := corestr.New.Hashmap.Cap(5)
		if !h.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_I8_HM02_HasItems(t *testing.T) {
	safeTest(t, "Test_I8_HM02_HasItems", func() {
		h := corestr.New.Hashmap.Cap(5)
		if h.HasItems() {
			t.Fatal("expected false")
		}
		h.AddOrUpdate("k", "v")
		if !h.HasItems() {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_HM03_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_I8_HM03_IsEmptyLock", func() {
		h := corestr.New.Hashmap.Cap(5)
		if !h.IsEmptyLock() {
			t.Fatal("expected empty")
		}
	})
}

func Test_I8_HM04_AddOrUpdate(t *testing.T) {
	safeTest(t, "Test_I8_HM04_AddOrUpdate", func() {
		h := corestr.New.Hashmap.Cap(5)
		isNew := h.AddOrUpdate("k", "v")
		if !isNew {
			t.Fatal("expected new")
		}
		isNew2 := h.AddOrUpdate("k", "v2")
		if isNew2 {
			t.Fatal("expected not new")
		}
	})
}

func Test_I8_HM05_Set(t *testing.T) {
	safeTest(t, "Test_I8_HM05_Set", func() {
		h := corestr.New.Hashmap.Cap(5)
		isNew := h.Set("k", "v")
		if !isNew {
			t.Fatal("expected new")
		}
	})
}

func Test_I8_HM06_SetTrim(t *testing.T) {
	safeTest(t, "Test_I8_HM06_SetTrim", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.SetTrim(" k ", " v ")
		if !h.Has("k") {
			t.Fatal("expected trimmed key")
		}
	})
}

func Test_I8_HM07_SetBySplitter(t *testing.T) {
	safeTest(t, "Test_I8_HM07_SetBySplitter", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.SetBySplitter("=", "key=value")
		if !h.Has("key") {
			t.Fatal("expected key")
		}
		// no splitter
		h.SetBySplitter("=", "nosplit")
		if !h.Has("nosplit") {
			t.Fatal("expected nosplit")
		}
	})
}

func Test_I8_HM08_AddOrUpdateKeyStrValInt(t *testing.T) {
	safeTest(t, "Test_I8_HM08_AddOrUpdateKeyStrValInt", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateKeyStrValInt("k", 42)
		if !h.Has("k") {
			t.Fatal("expected key")
		}
	})
}

func Test_I8_HM09_AddOrUpdateKeyStrValFloat(t *testing.T) {
	safeTest(t, "Test_I8_HM09_AddOrUpdateKeyStrValFloat", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateKeyStrValFloat("k", 3.14)
		if !h.Has("k") {
			t.Fatal("expected key")
		}
	})
}

func Test_I8_HM10_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	safeTest(t, "Test_I8_HM10_AddOrUpdateKeyStrValFloat64", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateKeyStrValFloat64("k", 3.14)
		if !h.Has("k") {
			t.Fatal("expected key")
		}
	})
}

func Test_I8_HM11_AddOrUpdateKeyStrValAny(t *testing.T) {
	safeTest(t, "Test_I8_HM11_AddOrUpdateKeyStrValAny", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateKeyStrValAny("k", "val")
		if !h.Has("k") {
			t.Fatal("expected key")
		}
	})
}

func Test_I8_HM12_AddOrUpdateKeyVal(t *testing.T) {
	safeTest(t, "Test_I8_HM12_AddOrUpdateKeyVal", func() {
		h := corestr.New.Hashmap.Cap(5)
		isNew := h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k", Value: "v"})
		if !isNew {
			t.Fatal("expected new")
		}
	})
}

func Test_I8_HM13_AddOrUpdateKeyValueAny(t *testing.T) {
	safeTest(t, "Test_I8_HM13_AddOrUpdateKeyValueAny", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "k", Value: 42})
		if !h.Has("k") {
			t.Fatal("expected key")
		}
	})
}

func Test_I8_HM14_AddOrUpdateLock(t *testing.T) {
	safeTest(t, "Test_I8_HM14_AddOrUpdateLock", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateLock("k", "v")
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HM15_AddOrUpdateWithWgLock(t *testing.T) {
	safeTest(t, "Test_I8_HM15_AddOrUpdateWithWgLock", func() {
		h := corestr.New.Hashmap.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateWithWgLock("k", "v", wg)
		wg.Wait()
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HM16_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_I8_HM16_AddOrUpdateHashmap", func() {
		h := corestr.New.Hashmap.Cap(5)
		other := corestr.New.Hashmap.Cap(2)
		other.AddOrUpdate("a", "1")
		h.AddOrUpdateHashmap(other)
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
		h.AddOrUpdateHashmap(nil)
		if h.Length() != 1 {
			t.Fatal("expected still 1")
		}
	})
}

func Test_I8_HM17_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_I8_HM17_AddOrUpdateMap", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateMap(map[string]string{"a": "1", "b": "2"})
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
		h.AddOrUpdateMap(nil)
	})
}

func Test_I8_HM18_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_I8_HM18_AddsOrUpdates", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddsOrUpdates(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_HM19_AddOrUpdateKeyAnyValues(t *testing.T) {
	safeTest(t, "Test_I8_HM19_AddOrUpdateKeyAnyValues", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateKeyAnyValues(
			corestr.KeyAnyValuePair{Key: "a", Value: 1},
		)
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HM20_AddOrUpdateKeyValues(t *testing.T) {
	safeTest(t, "Test_I8_HM20_AddOrUpdateKeyValues", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateKeyValues(
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HM21_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_I8_HM21_AddOrUpdateCollection", func() {
		h := corestr.New.Hashmap.Cap(5)
		keys := corestr.New.Collection.Strings([]string{"a", "b"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})
		h.AddOrUpdateCollection(keys, vals)
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
		// length mismatch
		vals2 := corestr.New.Collection.Strings([]string{"1"})
		h.AddOrUpdateCollection(keys, vals2)
		// nil keys
		h.AddOrUpdateCollection(nil, vals)
	})
}

// =============================================================================
// Hashmap — Query operations
// =============================================================================

func Test_I8_HM22_Has(t *testing.T) {
	safeTest(t, "Test_I8_HM22_Has", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("k", "v")
		if !h.Has("k") {
			t.Fatal("expected true")
		}
		if h.Has("z") {
			t.Fatal("expected false")
		}
	})
}

func Test_I8_HM23_Contains(t *testing.T) {
	safeTest(t, "Test_I8_HM23_Contains", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("k", "v")
		if !h.Contains("k") {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_HM24_ContainsLock(t *testing.T) {
	safeTest(t, "Test_I8_HM24_ContainsLock", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("k", "v")
		if !h.ContainsLock("k") {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_HM25_IsKeyMissing(t *testing.T) {
	safeTest(t, "Test_I8_HM25_IsKeyMissing", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("k", "v")
		if h.IsKeyMissing("k") {
			t.Fatal("expected false")
		}
		if !h.IsKeyMissing("z") {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_HM26_IsKeyMissingLock(t *testing.T) {
	safeTest(t, "Test_I8_HM26_IsKeyMissingLock", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("k", "v")
		if h.IsKeyMissingLock("k") {
			t.Fatal("expected false")
		}
	})
}

func Test_I8_HM27_HasLock(t *testing.T) {
	safeTest(t, "Test_I8_HM27_HasLock", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("k", "v")
		if !h.HasLock("k") {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_HM28_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_I8_HM28_HasAllStrings", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Set("b", "2")
		if !h.HasAllStrings("a", "b") {
			t.Fatal("expected true")
		}
		if h.HasAllStrings("a", "z") {
			t.Fatal("expected false")
		}
	})
}

func Test_I8_HM29_HasAll(t *testing.T) {
	safeTest(t, "Test_I8_HM29_HasAll", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		if !h.HasAll("a") {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_HM30_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_I8_HM30_HasAnyItem", func() {
		h := corestr.New.Hashmap.Cap(5)
		if h.HasAnyItem() {
			t.Fatal("expected false")
		}
		h.Set("a", "1")
		if !h.HasAnyItem() {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_HM31_HasAny(t *testing.T) {
	safeTest(t, "Test_I8_HM31_HasAny", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		if !h.HasAny("a", "z") {
			t.Fatal("expected true")
		}
		if h.HasAny("x", "y") {
			t.Fatal("expected false")
		}
	})
}

func Test_I8_HM32_HasWithLock(t *testing.T) {
	safeTest(t, "Test_I8_HM32_HasWithLock", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		if !h.HasWithLock("a") {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_HM33_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_I8_HM33_HasAllCollectionItems", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Set("b", "2")
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !h.HasAllCollectionItems(c) {
			t.Fatal("expected true")
		}
		if h.HasAllCollectionItems(nil) {
			t.Fatal("expected false for nil")
		}
	})
}

// =============================================================================
// Hashmap — Get, Values, Keys, Diff, Filter
// =============================================================================

func Test_I8_HM34_Items(t *testing.T) {
	safeTest(t, "Test_I8_HM34_Items", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		if len(h.Items()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HM35_SafeItems(t *testing.T) {
	safeTest(t, "Test_I8_HM35_SafeItems", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		if len(h.SafeItems()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HM36_ValuesList(t *testing.T) {
	safeTest(t, "Test_I8_HM36_ValuesList", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		if len(h.ValuesList()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HM37_ValuesCollection(t *testing.T) {
	safeTest(t, "Test_I8_HM37_ValuesCollection", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		c := h.ValuesCollection()
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HM38_ValuesHashset(t *testing.T) {
	safeTest(t, "Test_I8_HM38_ValuesHashset", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		hs := h.ValuesHashset()
		if hs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HM39_KeysValuesCollection(t *testing.T) {
	safeTest(t, "Test_I8_HM39_KeysValuesCollection", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		keys, values := h.KeysValuesCollection()
		if keys.Length() != 1 || values.Length() != 1 {
			t.Fatal("expected 1 each")
		}
	})
}

func Test_I8_HM40_KeysValuesList(t *testing.T) {
	safeTest(t, "Test_I8_HM40_KeysValuesList", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		keys, values := h.KeysValuesList()
		if len(keys) != 1 || len(values) != 1 {
			t.Fatal("expected 1 each")
		}
	})
}

func Test_I8_HM41_KeysValuePairs(t *testing.T) {
	safeTest(t, "Test_I8_HM41_KeysValuePairs", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		pairs := h.KeysValuePairs()
		if len(pairs) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HM42_GetKeysFilteredItems(t *testing.T) {
	safeTest(t, "Test_I8_HM42_GetKeysFilteredItems", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("abc", "1")
		h.Set("de", "2")
		result := h.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 2, false
		})
		if len(result) != 1 {
			t.Fatalf("expected 1, got %d", len(result))
		}
	})
}

func Test_I8_HM43_GetKeysFilteredCollection(t *testing.T) {
	safeTest(t, "Test_I8_HM43_GetKeysFilteredCollection", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("abc", "1")
		c := h.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HM44_ConcatNew(t *testing.T) {
	safeTest(t, "Test_I8_HM44_ConcatNew", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		other := corestr.New.Hashmap.Cap(2)
		other.Set("b", "2")
		result := h.ConcatNew(true, other)
		if result.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_I8_HM45_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_I8_HM45_ConcatNew_Empty", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		result := h.ConcatNew(true)
		if result.Length() < 1 {
			t.Fatal("expected at least 1")
		}
	})
}

func Test_I8_HM46_ConcatNewUsingMaps(t *testing.T) {
	safeTest(t, "Test_I8_HM46_ConcatNewUsingMaps", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		m := map[string]string{"b": "2"}
		result := h.ConcatNewUsingMaps(true, m)
		if result.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_I8_HM47_Diff(t *testing.T) {
	safeTest(t, "Test_I8_HM47_Diff", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Set("b", "2")
		other := corestr.New.Hashmap.Cap(5)
		other.Set("a", "1")
		diff := h.Diff(other)
		_ = diff
	})
}

func Test_I8_HM48_DiffRaw(t *testing.T) {
	safeTest(t, "Test_I8_HM48_DiffRaw", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		rawDiff := h.DiffRaw(map[string]string{"a": "1"})
		_ = rawDiff
	})
}

func Test_I8_HM49_AddsOrUpdatesUsingFilter(t *testing.T) {
	safeTest(t, "Test_I8_HM49_AddsOrUpdatesUsingFilter", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddsOrUpdatesUsingFilter(func(p corestr.KeyValuePair) (string, bool, bool) {
			return p.Value, true, false
		}, corestr.KeyValuePair{Key: "a", Value: "1"})
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HM50_AddsOrUpdatesAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_I8_HM50_AddsOrUpdatesAnyUsingFilter", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddsOrUpdatesAnyUsingFilter(func(p corestr.KeyAnyValuePair) (string, bool, bool) {
			return p.ValueString(), true, false
		}, corestr.KeyAnyValuePair{Key: "a", Value: 42})
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HM51_Json(t *testing.T) {
	safeTest(t, "Test_I8_HM51_Json", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		j := h.Json()
		if j.JsonString() == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_I8_HM52_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_I8_HM52_ParseInjectUsingJson", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		jr := h.JsonPtr()
		h2 := corestr.New.Hashmap.Cap(1)
		_, err := h2.ParseInjectUsingJson(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_I8_HM53_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_I8_HM53_ParseInjectUsingJson_Error", func() {
		h := corestr.New.Hashmap.Cap(1)
		bad := corejson.NewResult.UsingString(`invalid`)
		_, err := h.ParseInjectUsingJson(bad)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_I8_HM54_String(t *testing.T) {
	safeTest(t, "Test_I8_HM54_String", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		s := h.String()
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_I8_HM55_Clear(t *testing.T) {
	safeTest(t, "Test_I8_HM55_Clear", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Clear()
		if h.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_I8_HM56_Dispose(t *testing.T) {
	safeTest(t, "Test_I8_HM56_Dispose", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Dispose()
		if h.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_I8_HM57_ItemsCopyLock(t *testing.T) {
	safeTest(t, "Test_I8_HM57_ItemsCopyLock", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		m := h.ItemsCopyLock()
		if len(*m) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HM58_Collection(t *testing.T) {
	safeTest(t, "Test_I8_HM58_Collection", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		c := h.Collection()
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HM59_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_I8_HM59_AddOrUpdateStringsPtrWgLock", func() {
		h := corestr.New.Hashmap.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateStringsPtrWgLock(wg, []string{"a"}, []string{"1"})
		wg.Wait()
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_HM60_KeysValuePairsCollection(t *testing.T) {
	safeTest(t, "Test_I8_HM60_KeysValuePairsCollection", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		kvc := h.KeysValuePairsCollection()
		if kvc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}
