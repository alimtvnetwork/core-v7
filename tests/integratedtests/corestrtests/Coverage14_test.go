package corestrtests

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_Hashmap_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_BasicOps", func() {
		h := corestr.New.Hashmap.Cap(10)
		h.AddOrUpdate("a", "1")
		h.Set("b", "2")
		h.SetTrim(" c ", " 3 ")

		if h.Length() != 3 || h.IsEmpty() || !h.HasItems() || !h.HasAnyItem() {
			t.Fatal("expected 3")
		}

		if !h.Has("a") || !h.Contains("b") || h.IsKeyMissing("a") || !h.IsKeyMissing("z") {
			t.Fatal("Has/Contains/IsKeyMissing failed")
		}

		if !h.HasAll("a", "b") || h.HasAll("a", "z") {
			t.Fatal("HasAll failed")
		}

		if !h.HasAny("a", "z") || h.HasAny("x", "z") {
			t.Fatal("HasAny failed")
		}
	})
}

func Test_Cov14_Hashmap_SetBySplitter(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_SetBySplitter", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.SetBySplitter("=", "key=value")
		h.SetBySplitter("=", "onlykey")

		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov14_Hashmap_AddOrUpdateVariants(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_AddOrUpdateVariants", func() {
		h := corestr.New.Hashmap.Cap(10)

		h.AddOrUpdateKeyStrValInt("k1", 42)
		h.AddOrUpdateKeyStrValFloat("k2", 3.14)
		h.AddOrUpdateKeyStrValFloat64("k3", 2.71)
		h.AddOrUpdateKeyStrValAny("k4", "any")
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "k5", Value: "val5"})
		h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "k6", Value: "v6"})

		if h.Length() != 6 {
			t.Fatal("expected 6")
		}
	})
}

func Test_Cov14_Hashmap_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_AddsOrUpdates", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddsOrUpdates(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		if h.Length() != 2 {
			t.Fatal("expected 2")
		}

		h.AddsOrUpdates()
	})
}

func Test_Cov14_Hashmap_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_AddOrUpdateMap", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateMap(map[string]string{"a": "1", "b": "2"})

		if h.Length() != 2 {
			t.Fatal("expected 2")
		}

		h.AddOrUpdateMap(nil)
	})
}

func Test_Cov14_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_AddOrUpdateHashmap", func() {
		h1 := corestr.New.Hashmap.Cap(5)
		h1.Set("a", "1")

		h2 := corestr.New.Hashmap.Cap(5)
		h2.Set("b", "2")

		h1.AddOrUpdateHashmap(h2)
		h1.AddOrUpdateHashmap(nil)

		if h1.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov14_Hashmap_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_AddOrUpdateCollection", func() {
		h := corestr.New.Hashmap.Cap(5)
		keys := corestr.New.Collection.Strings([]string{"k1", "k2"})
		vals := corestr.New.Collection.Strings([]string{"v1", "v2"})

		h.AddOrUpdateCollection(keys, vals)
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}

		// Mismatched lengths
		h.AddOrUpdateCollection(keys, corestr.New.Collection.Strings([]string{"v1"}))

		// Nil
		h.AddOrUpdateCollection(nil, nil)
	})
}

func Test_Cov14_Hashmap_Get(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_Get", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")

		v, found := h.Get("a")
		if !found || v != "1" {
			t.Fatal("expected 1")
		}

		_, found = h.Get("z")
		if found {
			t.Fatal("expected not found")
		}

		v2, found2 := h.GetValue("a")
		if !found2 || v2 != "1" {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov14_Hashmap_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_ConcatNew", func() {
		h1 := corestr.New.Hashmap.Cap(5)
		h1.Set("a", "1")

		h2 := corestr.New.Hashmap.Cap(5)
		h2.Set("b", "2")

		result := h1.ConcatNew(true, h2)
		if result.Length() < 2 {
			t.Fatal("expected at least 2")
		}

		// Empty
		result2 := h1.ConcatNew(true)
		if result2.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov14_Hashmap_ConcatNewUsingMaps(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_ConcatNewUsingMaps", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")

		result := h.ConcatNewUsingMaps(true, map[string]string{"b": "2"})
		if result.Length() < 2 {
			t.Fatal("expected at least 2")
		}

		result2 := h.ConcatNewUsingMaps(true)
		_ = result2
	})
}

func Test_Cov14_Hashmap_Keys_Values(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_Keys_Values", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Set("b", "2")

		_ = h.Keys()
		_ = h.AllKeys()
		_ = h.ValuesList()
		_ = h.Items()
		_ = h.SafeItems()
		_ = h.Collection()
		_ = h.ValuesCollection()
		_ = h.ValuesHashset()
		_ = h.KeysCollection()
		_ = h.KeysValuePairs()
		_ = h.KeysValuePairsCollection()
	})
}

func Test_Cov14_Hashmap_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_IsEqual", func() {
		h1 := corestr.New.Hashmap.Cap(5)
		h1.Set("a", "1")

		h2 := corestr.New.Hashmap.Cap(5)
		h2.Set("a", "1")

		if !h1.IsEqualPtr(h2) {
			t.Fatal("expected equal")
		}

		h3 := corestr.New.Hashmap.Cap(5)
		h3.Set("a", "2")

		if h1.IsEqualPtr(h3) {
			t.Fatal("expected not equal")
		}

		// Same ptr
		if !h1.IsEqualPtr(h1) {
			t.Fatal("expected same ptr equal")
		}

		// Both empty
		e1 := corestr.Empty.Hashmap()
		e2 := corestr.Empty.Hashmap()
		if !e1.IsEqualPtr(e2) {
			t.Fatal("expected empty equal")
		}
	})
}

func Test_Cov14_Hashmap_Remove(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_Remove", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Remove("a")

		if h.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_Cov14_Hashmap_KeysToLower(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_KeysToLower", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("HELLO", "world")

		lowered := h.KeysToLower()
		if !lowered.Has("hello") {
			t.Fatal("expected hello")
		}
	})
}

func Test_Cov14_Hashmap_String(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_String", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		_ = h.String()

		empty := corestr.Empty.Hashmap()
		_ = empty.String()
	})
}

func Test_Cov14_Hashmap_Join(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_Join", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		_ = h.Join(",")
		_ = h.JoinKeys(",")
	})
}

func Test_Cov14_Hashmap_JSON(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_JSON", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")

		b, err := json.Marshal(h)
		if err != nil {
			t.Fatal("marshal failed")
		}

		h2 := corestr.Empty.Hashmap()
		err = json.Unmarshal(b, h2)
		if err != nil {
			t.Fatal("unmarshal failed")
		}

		if h2.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}
func Test_Cov14_Hashmap_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_Clear_Dispose", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Clear()

		if h.Length() != 0 {
			t.Fatal("expected 0")
		}

		h2 := corestr.New.Hashmap.Cap(5)
		h2.Set("a", "1")
		h2.Dispose()
	})
}

func Test_Cov14_Hashmap_ToError(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_ToError", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		_ = h.ToError(",")
		_ = h.ToDefaultError()
	})
}

func Test_Cov14_Hashmap_KeyValStringLines(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_KeyValStringLines", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		_ = h.KeyValStringLines()
	})
}

func Test_Cov14_Hashmap_ToStringsUsingCompiler(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_ToStringsUsingCompiler", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")

		lines := h.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })
		if len(lines) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov14_Hashmap_GetKeysFilteredItems(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_GetKeysFilteredItems", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("ab", "1")
		h.Set("cd", "2")

		result := h.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, s == "ab", false
		})

		if len(result) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov14_Hashmap_GetKeysFilteredCollection(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_GetKeysFilteredCollection", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("ab", "1")

		result := h.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		if result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov14_Hashmap_GetValuesExcept(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_GetValuesExcept", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Set("b", "2")

		result := h.GetValuesKeysExcept([]string{"a"})
		if len(result) != 1 {
			t.Fatal("expected 1")
		}

		result2 := h.GetValuesKeysExcept(nil)
		if len(result2) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov14_Hashmap_Creators(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_Creators", func() {
		_ = corestr.New.Hashmap.Empty()
		_ = corestr.New.Hashmap.Cap(5)
		_ = corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		_ = corestr.New.Hashmap.UsingMapOptions(true, 5, map[string]string{"a": "1"})
		_ = corestr.New.Hashmap.UsingMapOptions(false, 0, map[string]string{"a": "1"})
		_ = corestr.New.Hashmap.UsingMapOptions(true, 0, map[string]string{})
		_ = corestr.New.Hashmap.MapWithCap(5, map[string]string{"a": "1"})
		_ = corestr.New.Hashmap.MapWithCap(0, map[string]string{"a": "1"})
		_ = corestr.New.Hashmap.MapWithCap(5, map[string]string{})
		_ = corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		_ = corestr.New.Hashmap.KeyValues()
		_ = corestr.New.Hashmap.KeyAnyValues(corestr.KeyAnyValuePair{Key: "a", Value: 1})
		_ = corestr.New.Hashmap.KeyAnyValues()

		keys := corestr.New.Collection.Strings([]string{"k"})
		vals := corestr.New.Collection.Strings([]string{"v"})
		_ = corestr.New.Hashmap.KeyValuesCollection(keys, vals)
		_ = corestr.New.Hashmap.KeyValuesCollection(nil, nil)
		_ = corestr.New.Hashmap.KeyValuesStrings([]string{"k"}, []string{"v"})
		_ = corestr.New.Hashmap.KeyValuesStrings(nil, nil)
	})
}

func Test_Cov14_Hashmap_AddsOrUpdatesFilter(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_AddsOrUpdatesFilter", func() {
		h := corestr.New.Hashmap.Cap(5)

		h.AddsOrUpdatesUsingFilter(func(kv corestr.KeyValuePair) (string, bool, bool) {
			return kv.Value, true, false
		}, corestr.KeyValuePair{Key: "a", Value: "1"})

		if h.Length() != 1 {
			t.Fatal("expected 1")
		}

		h.AddsOrUpdatesUsingFilter(nil)
	})
}

func Test_Cov14_Hashmap_AddsOrUpdatesAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashmap_AddsOrUpdatesAnyUsingFilter", func() {
		h := corestr.New.Hashmap.Cap(5)

		h.AddsOrUpdatesAnyUsingFilter(func(kav corestr.KeyAnyValuePair) (string, bool, bool) {
			return kav.ValueString(), true, false
		}, corestr.KeyAnyValuePair{Key: "a", Value: 1})

		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// HashmapDiff — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_HashmapDiff_AllMethods(t *testing.T) {
	safeTest(t, "Test_Cov14_HashmapDiff_AllMethods", func() {
		diff := corestr.HashmapDiff(map[string]string{"a": "1", "b": "2"})

		if diff.Length() != 2 || diff.IsEmpty() || !diff.HasAnyItem() {
			t.Fatal("expected 2")
		}

		if diff.LastIndex() != 1 {
			t.Fatal("expected 1")
		}

		_ = diff.AllKeysSorted()
		_ = diff.MapAnyItems()
		_ = diff.Raw()
		_ = diff.RawMapStringAnyDiff()

		// IsRawEqual
		if !diff.IsRawEqual(map[string]string{"a": "1", "b": "2"}) {
			t.Fatal("expected equal")
		}

		if diff.HasAnyChanges(map[string]string{"a": "1", "b": "2"}) {
			t.Fatal("expected no changes")
		}

		_ = diff.DiffRaw(map[string]string{"a": "1", "c": "3"})
		_ = diff.DiffJsonMessage(map[string]string{"a": "2"})
		_ = diff.HashmapDiffUsingRaw(map[string]string{"a": "1"})
		_ = diff.ShouldDiffMessage("test", map[string]string{"a": "1"})
		_ = diff.LogShouldDiffMessage("test", map[string]string{"a": "1"})
		_ = diff.ToStringsSliceOfDiffMap(map[string]string{"a": "1"})

		_, _ = diff.Serialize()

		// Nil receiver
		var nilDiff *corestr.HashmapDiff
		if nilDiff.Length() != 0 {
			t.Fatal("expected 0")
		}
		_ = nilDiff.Raw()
		_ = nilDiff.MapAnyItems()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_Hashset_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_BasicOps", func() {
		h := corestr.New.Hashset.Cap(10)
		h.Add("a")
		h.Add("b")

		if h.Length() != 2 || h.IsEmpty() || !h.HasItems() || !h.HasAnyItem() {
			t.Fatal("expected 2")
		}

		if !h.Has("a") || !h.Contains("b") || h.IsMissing("a") || !h.IsMissing("z") {
			t.Fatal("Has/Contains/IsMissing failed")
		}

		if !h.HasAll("a", "b") || h.HasAll("a", "z") {
			t.Fatal("HasAll failed")
		}

		if !h.HasAny("a", "z") || h.HasAny("x", "z") {
			t.Fatal("HasAny failed")
		}

		if !h.IsAllMissing("x", "y") || h.IsAllMissing("a") {
			t.Fatal("IsAllMissing failed")
		}
	})
}

func Test_Cov14_Hashset_AddVariants(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_AddVariants", func() {
		h := corestr.New.Hashset.Cap(10)
		h.AddNonEmpty("")
		h.AddNonEmpty("a")
		h.AddNonEmptyWhitespace("  ")
		h.AddNonEmptyWhitespace("b")
		h.AddIf(false, "skip")
		h.AddIf(true, "c")
		h.AddIfMany(false, "s1")
		h.AddIfMany(true, "d", "e")
		h.AddFunc(func() string { return "f" })
		h.AddBool("g")
		h.AddBool("g") // Already exists

		str := "h"
		h.AddPtr(&str)

		if h.Length() != 8 {
			t.Fatalf("expected 8 got %d", h.Length())
		}
	})
}

func Test_Cov14_Hashset_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_AddFuncErr", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov14_Hashset_AddStrings_Adds(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_AddStrings_Adds", func() {
		h := corestr.New.Hashset.Cap(10)
		h.AddStrings([]string{"a", "b"})
		h.Adds("c", "d")

		if h.Length() != 4 {
			t.Fatal("expected 4")
		}
	})
}

func Test_Cov14_Hashset_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_AddHashsetItems", func() {
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"b"})

		h1.AddHashsetItems(h2)
		h1.AddHashsetItems(nil)

		if h1.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov14_Hashset_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_AddItemsMap", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddItemsMap(map[string]bool{"a": true, "b": false})

		if h.Length() != 1 {
			t.Fatal("expected 1 (false items skipped)")
		}
	})
}

func Test_Cov14_Hashset_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_AddCollection", func() {
		h := corestr.New.Hashset.Cap(5)
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		h.AddCollection(c)

		if h.Length() != 2 {
			t.Fatal("expected 2")
		}

		h.AddCollection(nil)
	})
}

func Test_Cov14_Hashset_AddSimpleSlice(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_AddSimpleSlice", func() {
		h := corestr.New.Hashset.Cap(5)
		s := corestr.New.SimpleSlice.Lines("a", "b")
		h.AddSimpleSlice(s)

		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov14_Hashset_ConcatNewHashsets(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_ConcatNewHashsets", func() {
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"b"})

		result := h1.ConcatNewHashsets(true, h2)
		if result.Length() != 2 {
			t.Fatal("expected 2")
		}

		// Empty
		result2 := h1.ConcatNewHashsets(true)
		if result2.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov14_Hashset_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_ConcatNewStrings", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.ConcatNewStrings(true, []string{"b", "c"})

		if result.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_Cov14_Hashset_List_OrderedList_SortedList(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_List_OrderedList_SortedList", func() {
		h := corestr.New.Hashset.Strings([]string{"c", "a", "b"})

		_ = h.List()
		_ = h.OrderedList()
		_ = h.SortedList()
		_ = h.SafeStrings()
		_ = h.Lines()
		_ = h.SimpleSlice()
	})
}

func Test_Cov14_Hashset_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_IsEquals", func() {
		h1 := corestr.New.Hashset.Strings([]string{"a", "b"})
		h2 := corestr.New.Hashset.Strings([]string{"a", "b"})

		if !h1.IsEquals(h2) || !h1.IsEqual(h2) {
			t.Fatal("expected equal")
		}

		// Same ptr
		if !h1.IsEquals(h1) {
			t.Fatal("expected same ptr equal")
		}

		// Both empty
		e1 := corestr.Empty.Hashset()
		e2 := corestr.Empty.Hashset()
		if !e1.IsEquals(e2) {
			t.Fatal("expected empty equal")
		}
	})
}

func Test_Cov14_Hashset_Filter(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_Filter", func() {
		h := corestr.New.Hashset.Strings([]string{"ab", "cd", "ef"})

		result := h.Filter(func(s string) bool { return s == "ab" })
		if result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov14_Hashset_GetFilteredItems(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_GetFilteredItems", func() {
		h := corestr.New.Hashset.Strings([]string{"ab", "cd"})

		result := h.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		if len(result) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov14_Hashset_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_GetAllExcept", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b", "c"})

		result := h.GetAllExcept([]string{"b"})
		if len(result) != 2 {
			t.Fatal("expected 2")
		}

		result2 := h.GetAllExcept(nil)
		if len(result2) != 3 {
			t.Fatal("expected 3")
		}

		_ = h.GetAllExceptSpread("a")
	})
}

func Test_Cov14_Hashset_Remove(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_Remove", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.Remove("a")

		if h.Length() != 1 {
			t.Fatal("expected 1")
		}

		h.SafeRemove("z") // no-op
		h.SafeRemove("b")

		if h.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_Cov14_Hashset_ToLowerSet(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_ToLowerSet", func() {
		h := corestr.New.Hashset.Strings([]string{"HELLO"})
		lowered := h.ToLowerSet()

		if !lowered.Has("hello") {
			t.Fatal("expected hello")
		}
	})
}

func Test_Cov14_Hashset_Resize(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_Resize", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		h.Resize(100)
		h.AddCapacities(50)

		if !h.Has("a") {
			t.Fatal("expected a after resize")
		}
	})
}

func Test_Cov14_Hashset_String(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_String", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		_ = h.String()
		_ = h.Join(",")
		_ = h.JoinSorted(",")
		_ = h.JoinLine()

		empty := corestr.Empty.Hashset()
		_ = empty.String()
	})
}

func Test_Cov14_Hashset_WrapQuotes(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_WrapQuotes", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		_ = h.WrapDoubleQuote()
		_ = h.WrapSingleQuote()
		_ = h.WrapDoubleQuoteIfMissing()
		_ = h.WrapSingleQuoteIfMissing()
	})
}

func Test_Cov14_Hashset_Transpile(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_Transpile", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.Transpile(func(s string) string { return s + "!" })

		if result.Length() < 1 {
			t.Fatal("expected at least 1")
		}

		// Empty
		empty := corestr.Empty.Hashset()
		_ = empty.Transpile(func(s string) string { return s })
	})
}

func Test_Cov14_Hashset_JSON(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_JSON", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})

		b, err := json.Marshal(h)
		if err != nil {
			t.Fatal("marshal failed")
		}

		h2 := corestr.Empty.Hashset()
		err = json.Unmarshal(b, h2)
		if err != nil {
			t.Fatal("unmarshal failed")
		}
	})
}

func Test_Cov14_Hashset_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_DistinctDiff", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})

		diffLines := h.DistinctDiffLinesRaw("b", "c")
		if len(diffLines) != 2 {
			t.Fatal("expected 2")
		}

		diffMap := h.DistinctDiffLines("b", "c")
		if len(diffMap) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov14_Hashset_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_Clear_Dispose", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.Clear()
		if h.Length() != 0 {
			t.Fatal("expected 0")
		}

		h2 := corestr.New.Hashset.Strings([]string{"a"})
		h2.Dispose()
	})
}

func Test_Cov14_Hashset_MapStringAny(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_MapStringAny", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		m := h.MapStringAny()

		if len(m) != 1 {
			t.Fatal("expected 1")
		}

		_ = h.MapStringAnyDiff()
	})
}

func Test_Cov14_Hashset_Creators(t *testing.T) {
	safeTest(t, "Test_Cov14_Hashset_Creators", func() {
		_ = corestr.New.Hashset.Empty()
		_ = corestr.New.Hashset.Cap(5)
		_ = corestr.New.Hashset.Strings([]string{"a"})
		_ = corestr.New.Hashset.Strings(nil)
		_ = corestr.New.Hashset.StringsSpreadItems("a", "b")
		_ = corestr.New.Hashset.StringsOption(5, true, "a")
		_ = corestr.New.Hashset.StringsOption(0, false)
		_ = corestr.New.Hashset.StringsOption(5, false)
		_ = corestr.New.Hashset.UsingMap(map[string]bool{"a": true})
		_ = corestr.New.Hashset.UsingMap(map[string]bool{})
		_ = corestr.New.Hashset.UsingMapOption(5, true, map[string]bool{"a": true})
		_ = corestr.New.Hashset.UsingMapOption(5, false, map[string]bool{"a": true})
		_ = corestr.New.Hashset.UsingMapOption(5, false, map[string]bool{})

		c := corestr.New.Collection.Strings([]string{"a"})
		_ = corestr.New.Hashset.UsingCollection(c)
		_ = corestr.New.Hashset.UsingCollection(nil)

		s := corestr.New.SimpleSlice.Lines("a")
		_ = corestr.New.Hashset.SimpleSlice(s)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedList — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_LinkedList_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_BasicOps", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a").Add("b").Add("c")

		if ll.Length() != 3 || ll.IsEmpty() || !ll.HasItems() {
			t.Fatal("expected 3")
		}

		if ll.Head().Element != "a" || ll.Tail().Element != "c" {
			t.Fatal("Head/Tail failed")
		}
	})
}

func Test_Cov14_LinkedList_AddVariants(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_AddVariants", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmpty("")
		ll.AddNonEmpty("a")
		ll.AddNonEmptyWhitespace("  ")
		ll.AddNonEmptyWhitespace("b")
		ll.AddIf(false, "skip")
		ll.AddIf(true, "c")
		ll.AddsIf(false, "s1")
		ll.AddsIf(true, "d")
		ll.AddFunc(func() string { return "e" })
		ll.Push("f")
		ll.PushBack("g")
		ll.PushFront("front")
		ll.AddFront("front2")

		if ll.Length() < 7 {
			t.Fatal("expected at least 7")
		}
	})
}

func Test_Cov14_LinkedList_Adds_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_Adds_AddStrings", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.AddStrings([]string{"d", "e"})

		if ll.Length() != 5 {
			t.Fatal("expected 5")
		}
	})
}

func Test_Cov14_LinkedList_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_AddCollection", func() {
		ll := corestr.New.LinkedList.Create()
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		ll.AddCollection(c)
		ll.AddCollection(nil)

		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov14_LinkedList_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_IsEquals", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a", "b"})

		if !ll1.IsEquals(ll2) {
			t.Fatal("expected equal")
		}

		// Case insensitive
		ll3 := corestr.New.LinkedList.Strings([]string{"A", "B"})
		if !ll1.IsEqualsWithSensitive(ll3, false) {
			t.Fatal("expected equal insensitive")
		}
	})
}

func Test_Cov14_LinkedList_List_ToCollection(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_List_ToCollection", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		_ = ll.List()
		_ = ll.ToCollection(0)
	})
}

func Test_Cov14_LinkedList_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_SafeIndexAt", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})

		node := ll.SafeIndexAt(1)
		if node == nil || node.Element != "b" {
			t.Fatal("expected b")
		}

		if ll.SafeIndexAt(-1) != nil || ll.SafeIndexAt(10) != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_Cov14_LinkedList_SafePointerIndexAtUsingDefault(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_SafePointerIndexAtUsingDefault", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})

		if ll.SafePointerIndexAtUsingDefault(0, "def") != "a" {
			t.Fatal("expected a")
		}

		if ll.SafePointerIndexAtUsingDefault(5, "def") != "def" {
			t.Fatal("expected def")
		}
	})
}

func Test_Cov14_LinkedList_Loop(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_Loop", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		count := 0

		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})

		if count != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_Cov14_LinkedList_String(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_String", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.String()
		_ = ll.Join(",")

		empty := corestr.New.LinkedList.Create()
		_ = empty.String()
	})
}

func Test_Cov14_LinkedList_JSON(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_JSON", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})

		b, err := json.Marshal(ll)
		if err != nil {
			t.Fatal("marshal failed")
		}

		ll2 := corestr.New.LinkedList.Create()
		err = json.Unmarshal(b, ll2)
		if err != nil {
			t.Fatal("unmarshal failed")
		}

		if ll2.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov14_LinkedList_Clear(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_Clear", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.Clear()

		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_Cov14_LinkedList_Creators(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_Creators", func() {
		_ = corestr.New.LinkedList.Create()
		_ = corestr.New.LinkedList.Empty()
		_ = corestr.New.LinkedList.Strings([]string{"a"})
		_ = corestr.New.LinkedList.Strings(nil)
		_ = corestr.New.LinkedList.SpreadStrings("a")
		_ = corestr.New.LinkedList.SpreadStrings()
		_ = corestr.New.LinkedList.UsingMap(map[string]bool{"a": true})
		_ = corestr.New.LinkedList.UsingMap(nil)
	})
}

func Test_Cov14_LinkedList_AppendNode(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_AppendNode", func() {
		ll := corestr.New.LinkedList.Create()
		node := &corestr.LinkedListNode{Element: "a"}
		ll.AppendNode(node)

		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov14_LinkedList_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_GetNextNodes", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		nodes := ll.GetNextNodes(2)

		if len(nodes) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov14_LinkedList_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_GetAllLinkedNodes", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		nodes := ll.GetAllLinkedNodes()

		if len(nodes) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov14_LinkedList_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_GetCompareSummary", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"a"})
		ll2 := corestr.New.LinkedList.Strings([]string{"b"})
		_ = ll1.GetCompareSummary(ll2, "left", "right")
	})
}

func Test_Cov14_LinkedList_Joins(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedList_Joins", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		result := ll.Joins(",", "c")

		if result == "" {
			t.Fatal("expected non-empty")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedListNode — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_LinkedListNode_AllMethods(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedListNode_AllMethods", func() {
		node := &corestr.LinkedListNode{Element: "a"}

		if node.HasNext() {
			t.Fatal("expected no next")
		}

		if node.String() != "a" {
			t.Fatal("expected a")
		}

		_ = node.Clone()
		_ = node.IsEqualValue("a")
		_ = node.IsEqualValueSensitive("A", false)
		_ = node.List()
		_ = node.Join(",")
		_ = node.StringList("header: ")

		end, length := node.EndOfChain()
		if end != node || length != 1 {
			t.Fatal("expected self, 1")
		}

		if !node.IsEqual(node) {
			t.Fatal("expected equal to self")
		}

		node2 := &corestr.LinkedListNode{Element: "a"}
		if !node.IsEqual(node2) {
			t.Fatal("expected equal")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// NonChainedLinkedListNodes — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_NonChainedLinkedListNodes(t *testing.T) {
	safeTest(t, "Test_Cov14_NonChainedLinkedListNodes", func() {
		nodes := corestr.NewNonChainedLinkedListNodes(5)
		n1 := &corestr.LinkedListNode{Element: "a"}
		n2 := &corestr.LinkedListNode{Element: "b"}
		nodes.Adds(n1, n2)

		if nodes.Length() != 2 || nodes.IsEmpty() || !nodes.HasItems() {
			t.Fatal("expected 2")
		}

		if nodes.First().Element != "a" || nodes.Last().Element != "b" {
			t.Fatal("First/Last failed")
		}

		_ = nodes.FirstOrDefault()
		_ = nodes.LastOrDefault()
		_ = nodes.IsChainingApplied()
		_ = nodes.Items()

		nodes.ApplyChaining()
		if !nodes.IsChainingApplied() {
			t.Fatal("expected chaining applied")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// NonChainedLinkedCollectionNodes — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_NonChainedLinkedCollectionNodes(t *testing.T) {
	safeTest(t, "Test_Cov14_NonChainedLinkedCollectionNodes", func() {
		nodes := corestr.NewNonChainedLinkedCollectionNodes(5)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		n1 := &corestr.LinkedCollectionNode{Element: c1}
		nodes.Adds(n1)

		if nodes.Length() != 1 || nodes.IsEmpty() {
			t.Fatal("expected 1")
		}

		_ = nodes.First()
		_ = nodes.FirstOrDefault()
		_ = nodes.Last()
		_ = nodes.LastOrDefault()
		_ = nodes.Items()

		nodes.ApplyChaining()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValueCollection — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_KeyValueCollection_AllMethods(t *testing.T) {
	safeTest(t, "Test_Cov14_KeyValueCollection_AllMethods", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1")
		kvc.Add("k2", "v2")
		kvc.AddIf(true, "k3", "v3")
		kvc.AddIf(false, "s", "s")

		if kvc.Length() != 3 || kvc.Count() != 3 || kvc.IsEmpty() {
			t.Fatal("expected 3")
		}

		if !kvc.HasAnyItem() || kvc.LastIndex() != 2 || !kvc.HasIndex(2) || kvc.HasIndex(5) {
			t.Fatal("index checks failed")
		}

		_ = kvc.First()
		_ = kvc.FirstOrDefault()
		_ = kvc.Last()
		_ = kvc.LastOrDefault()

		if !kvc.HasKey("k1") || !kvc.IsContains("k1") {
			t.Fatal("HasKey failed")
		}

		v, found := kvc.Get("k1")
		if !found || v != "v1" {
			t.Fatal("Get failed")
		}

		_ = kvc.AllKeys()
		_ = kvc.AllKeysSorted()
		_ = kvc.AllValues()
		_ = kvc.SafeValueAt(0)
		_ = kvc.SafeValueAt(100)
		_ = kvc.SafeValuesAtIndexes(0, 1)
		_ = kvc.Strings()
		_ = kvc.String()
		_ = kvc.Compile()
		_ = kvc.Join(",")
		_ = kvc.JoinKeys(",")
		_ = kvc.JoinValues(",")
		_ = kvc.Hashmap()
		_ = kvc.Map()
	})
}

func Test_Cov14_KeyValueCollection_Adds(t *testing.T) {
	safeTest(t, "Test_Cov14_KeyValueCollection_Adds", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Adds(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		kvc.AddMap(map[string]string{"c": "3"})
		kvc.AddHashsetMap(map[string]bool{"d": true})

		h := corestr.New.Hashset.Strings([]string{"e"})
		kvc.AddHashset(h)

		hm := corestr.New.Hashmap.Cap(5)
		hm.Set("f", "6")
		kvc.AddsHashmap(hm)
		kvc.AddsHashmaps(hm)

		if kvc.Length() != 7 {
			t.Fatalf("expected 7 got %d", kvc.Length())
		}
	})
}

func Test_Cov14_KeyValueCollection_Find(t *testing.T) {
	safeTest(t, "Test_Cov14_KeyValueCollection_Find", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("a", "1")
		kvc.Add("b", "2")

		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, kv.Key == "a", false
		})

		if len(found) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov14_KeyValueCollection_JSON(t *testing.T) {
	safeTest(t, "Test_Cov14_KeyValueCollection_JSON", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("a", "1")

		b, err := json.Marshal(kvc)
		if err != nil {
			t.Fatal("marshal failed")
		}

		kvc2 := corestr.Empty.KeyValueCollection()
		err = json.Unmarshal(b, kvc2)
		if err != nil {
			t.Fatal("unmarshal failed")
		}
	})
}

func Test_Cov14_KeyValueCollection_StringsUsingFormat(t *testing.T) {
	safeTest(t, "Test_Cov14_KeyValueCollection_StringsUsingFormat", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("a", "1")
		result := kvc.StringsUsingFormat("%s=%s")

		if len(result) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov14_KeyValueCollection_AddStringBySplit(t *testing.T) {
	safeTest(t, "Test_Cov14_KeyValueCollection_AddStringBySplit", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddStringBySplit("=", "key=value")
		kvc.AddStringBySplitTrim("=", " key = value ")

		if kvc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov14_KeyValueCollection_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov14_KeyValueCollection_Clear_Dispose", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("a", "1")
		kvc.Clear()

		if kvc.Length() != 0 {
			t.Fatal("expected 0")
		}

		kvc2 := corestr.Empty.KeyValueCollection()
		kvc2.Add("x", "y")
		kvc2.Dispose()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnce — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_SimpleStringOnce_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_BasicOps", func() {
		s := corestr.New.SimpleStringOnce.Init("hello")

		if !s.IsInitialized() || !s.IsDefined() || s.IsUninitialized() || s.IsInvalid() {
			t.Fatal("expected initialized")
		}

		if s.Value() != "hello" || s.SafeValue() != "hello" {
			t.Fatal("expected hello")
		}

		if s.IsEmpty() || s.IsWhitespace() {
			t.Fatal("expected non-empty")
		}

		if !s.HasValidNonEmpty() || !s.HasValidNonWhitespace() || !s.HasSafeNonEmpty() {
			t.Fatal("expected valid")
		}

		_ = s.Trim()
		_ = s.ValueBytes()
		_ = s.ValueBytesPtr()
		_ = s.String()
		_ = s.StringPtr()
		_ = s.NonPtr()
		_ = s.Ptr()
	})
}

func Test_Cov14_SimpleStringOnce_SetOnUninitialized(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_SetOnUninitialized", func() {
		s := corestr.New.SimpleStringOnce.Uninitialized("")

		err := s.SetOnUninitialized("val")
		if err != nil {
			t.Fatal("expected no error")
		}

		err = s.SetOnUninitialized("another")
		if err == nil {
			t.Fatal("expected error for already initialized")
		}
	})
}

func Test_Cov14_SimpleStringOnce_GetSetOnce(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_GetSetOnce", func() {
		s := corestr.New.SimpleStringOnce.Uninitialized("")

		val := s.GetSetOnce("first")
		if val != "first" {
			t.Fatal("expected first")
		}

		val2 := s.GetSetOnce("second")
		if val2 != "first" {
			t.Fatal("expected first (already set)")
		}
	})
}

func Test_Cov14_SimpleStringOnce_GetOnce(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_GetOnce", func() {
		s := corestr.New.SimpleStringOnce.Uninitialized("")
		val := s.GetOnce()

		if val != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_Cov14_SimpleStringOnce_GetOnceFunc(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_GetOnceFunc", func() {
		s := corestr.New.SimpleStringOnce.Uninitialized("")
		val := s.GetOnceFunc(func() string { return "computed" })

		if val != "computed" {
			t.Fatal("expected computed")
		}

		val2 := s.GetOnceFunc(func() string { return "other" })
		if val2 != "computed" {
			t.Fatal("expected computed (cached)")
		}
	})
}

func Test_Cov14_SimpleStringOnce_SetOnceIfUninitialized(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_SetOnceIfUninitialized", func() {
		s := corestr.New.SimpleStringOnce.Uninitialized("")

		if !s.SetOnceIfUninitialized("val") {
			t.Fatal("expected set")
		}

		if s.SetOnceIfUninitialized("other") {
			t.Fatal("expected not set")
		}
	})
}

func Test_Cov14_SimpleStringOnce_Reset_Invalidate(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_Reset_Invalidate", func() {
		s := corestr.New.SimpleStringOnce.Init("val")
		s.Reset()

		if s.IsInitialized() {
			t.Fatal("expected uninitialized")
		}

		s2 := corestr.New.SimpleStringOnce.Init("val2")
		s2.Invalidate()

		if s2.IsInitialized() {
			t.Fatal("expected uninitialized")
		}
	})
}

func Test_Cov14_SimpleStringOnce_Conversions(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_Conversions", func() {
		s := corestr.New.SimpleStringOnce.Init("42")

		if s.Int() != 42 || s.ValueInt(0) != 42 || s.ValueDefInt() != 42 {
			t.Fatal("expected 42")
		}

		_ = s.Byte()
		_ = s.Int16()
		_ = s.Int32()
		_ = s.ValueByte(0)
		_ = s.ValueDefByte()

		f := corestr.New.SimpleStringOnce.Init("3.14")
		_ = f.ValueFloat64(0)
		_ = f.ValueDefFloat64()

		b := corestr.New.SimpleStringOnce.Init("true")
		if !b.Boolean(false) || !b.BooleanDefault() || !b.IsValueBool() {
			t.Fatal("expected true")
		}

		y := corestr.New.SimpleStringOnce.Init("yes")
		if !y.Boolean(false) {
			t.Fatal("expected true for yes")
		}

		_ = s.IsSetter(false)
		_ = s.IsSetter(true)
	})
}

func Test_Cov14_SimpleStringOnce_WithinRange(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_WithinRange", func() {
		s := corestr.New.SimpleStringOnce.Init("50")

		val, inRange := s.WithinRange(true, 0, 100)
		if !inRange || val != 50 {
			t.Fatal("expected 50 in range")
		}

		val2, inRange2 := s.WithinRange(true, 60, 100)
		if inRange2 || val2 != 60 {
			t.Fatal("expected min boundary")
		}

		_, _ = s.Uint16()
		_, _ = s.Uint32()
	})
}

func Test_Cov14_SimpleStringOnce_Is_Contains_Regex(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_Is_Contains_Regex", func() {
		s := corestr.New.SimpleStringOnce.Init("hello world")

		if !s.Is("hello world") || s.Is("other") {
			t.Fatal("Is failed")
		}

		if !s.IsAnyOf("hello world", "other") || !s.IsAnyOf() {
			t.Fatal("IsAnyOf failed")
		}

		if !s.IsContains("hello") {
			t.Fatal("IsContains failed")
		}

		if !s.IsAnyContains("hello", "xyz") || !s.IsAnyContains() {
			t.Fatal("IsAnyContains failed")
		}

		if !s.IsEqualNonSensitive("HELLO WORLD") {
			t.Fatal("expected equal")
		}

		re := regexp.MustCompile(`hello`)
		if !s.IsRegexMatches(re) || s.IsRegexMatches(nil) {
			t.Fatal("regex failed")
		}

		_ = s.RegexFindString(re)
		_ = s.RegexFindString(nil)
		_, _ = s.RegexFindAllStringsWithFlag(re, -1)
		_ = s.RegexFindAllStrings(re, -1)
	})
}

func Test_Cov14_SimpleStringOnce_Split(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_Split", func() {
		s := corestr.New.SimpleStringOnce.Init("a,b,c")

		_ = s.Split(",")
		_ = s.SplitNonEmpty(",")
		_ = s.SplitTrimNonWhitespace(",")
		_, _ = s.SplitLeftRight(",")
		_, _ = s.SplitLeftRightTrim(",")
		_ = s.LinesSimpleSlice()
		_ = s.SimpleSlice(",")
	})
}

func Test_Cov14_SimpleStringOnce_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_ConcatNew", func() {
		s := corestr.New.SimpleStringOnce.Init("hello")
		r := s.ConcatNew(" world")

		if r.Value() != "hello world" {
			t.Fatal("expected hello world")
		}

		r2 := s.ConcatNewUsingStrings("-", "world")
		_ = r2
	})
}
func Test_Cov14_SimpleStringOnce_JSON(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_JSON", func() {
		s := corestr.New.SimpleStringOnce.Init("val")

		b, err := json.Marshal(&s)
		if err != nil {
			t.Fatal("marshal failed")
		}

		s2 := corestr.Empty.SimpleStringOnce()
		err = json.Unmarshal(b, &s2)
		if err != nil {
			t.Fatal("unmarshal failed")
		}
	})
}

func Test_Cov14_SimpleStringOnce_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_Dispose", func() {
		s := corestr.New.SimpleStringOnce.Init("val")
		s.Dispose()
	})
}

func Test_Cov14_SimpleStringOnce_Creators(t *testing.T) {
	safeTest(t, "Test_Cov14_SimpleStringOnce_Creators", func() {
		_ = corestr.New.SimpleStringOnce.Init("val")
		_ = corestr.New.SimpleStringOnce.InitPtr("val")
		_ = corestr.New.SimpleStringOnce.Uninitialized("val")
		_ = corestr.New.SimpleStringOnce.Create("val", true)
		_ = corestr.New.SimpleStringOnce.CreatePtr("val", true)
		_ = corestr.New.SimpleStringOnce.Empty()
		_ = corestr.New.SimpleStringOnce.Any(true, 42, true)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionsOfCollection — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_CollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_Cov14_CollectionsOfCollection", func() {
		coc := corestr.Empty.CollectionsOfCollection()

		if !coc.IsEmpty() || coc.HasItems() {
			t.Fatal("expected empty")
		}

		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		coc.Add(c1)

		if coc.Length() != 1 || coc.AllIndividualItemsLength() != 2 {
			t.Fatal("expected 1/2")
		}

		_ = coc.Items()
		_ = coc.List(0)
		_ = coc.ToCollection()
		_ = coc.String()

		coc.AddStrings(false, []string{"c", "d"})
		coc.AddsStringsOfStrings(false, []string{"e"}, []string{"f"})
	})
}

func Test_Cov14_CollectionsOfCollection_JSON(t *testing.T) {
	safeTest(t, "Test_Cov14_CollectionsOfCollection_JSON", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)

		b, err := json.Marshal(coc)
		if err != nil {
			t.Fatal("marshal failed")
		}

		coc2 := corestr.Empty.CollectionsOfCollection()
		err = json.Unmarshal(b, coc2)
		if err != nil {
			t.Fatal("unmarshal failed")
		}
	})
}
func Test_Cov14_HashsetsCollection_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov14_HashsetsCollection_IsEqual", func() {
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"a"})

		hc1 := corestr.Empty.HashsetsCollection()
		hc1.Add(h1)

		hc2 := corestr.Empty.HashsetsCollection()
		hc2.Add(h2)

		if !hc1.IsEqualPtr(hc2) {
			t.Fatal("expected equal")
		}
	})
}

func Test_Cov14_HashsetsCollection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov14_HashsetsCollection_ConcatNew", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.Empty.HashsetsCollection()
		hc.Add(h)

		hc2 := corestr.Empty.HashsetsCollection()
		hc2.Add(corestr.New.Hashset.Strings([]string{"b"}))

		result := hc.ConcatNew(hc2)
		if result.Length() != 2 {
			t.Fatal("expected 2")
		}

		// No args
		result2 := hc.ConcatNew()
		if result2.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov14_HashsetsCollection_JSON(t *testing.T) {
	safeTest(t, "Test_Cov14_HashsetsCollection_JSON", func() {
		hc := corestr.Empty.HashsetsCollection()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))

		b, err := json.Marshal(hc)
		if err != nil {
			t.Fatal("marshal failed")
		}

		hc2 := corestr.Empty.HashsetsCollection()
		err = json.Unmarshal(b, hc2)
		if err != nil {
			t.Fatal("unmarshal failed")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedCollections — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_LinkedCollections_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedCollections_BasicOps", func() {
		lc := corestr.Empty.LinkedCollections()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})

		lc.Add(c1)
		lc.Add(c2)

		if lc.Length() != 2 || lc.IsEmpty() || !lc.HasItems() {
			t.Fatal("expected 2")
		}

		_ = lc.Head()
		_ = lc.Tail()
		_ = lc.First()
		_ = lc.Last()
		_ = lc.FirstOrDefault()
		_ = lc.LastOrDefault()
		_ = lc.AllIndividualItemsLength()
	})
}

func Test_Cov14_LinkedCollections_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedCollections_AddStrings", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.AddStrings("a", "b")

		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov14_LinkedCollections_AddFront(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedCollections_AddFront", func() {
		lc := corestr.Empty.LinkedCollections()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})

		lc.Add(c1)
		lc.AddFront(c2)

		if lc.First().First() != "b" {
			t.Fatal("expected b first")
		}
	})
}

func Test_Cov14_LinkedCollections_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedCollections_IsEquals", func() {
		lc1 := corestr.Empty.LinkedCollections()
		lc1.AddStrings("a")

		lc2 := corestr.Empty.LinkedCollections()
		lc2.AddStrings("a")

		if !lc1.IsEqualsPtr(lc2) {
			t.Fatal("expected equal")
		}
	})
}

func Test_Cov14_LinkedCollections_AddAnother(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedCollections_AddAnother", func() {
		lc1 := corestr.Empty.LinkedCollections()
		lc1.AddStrings("a")

		lc2 := corestr.Empty.LinkedCollections()
		lc2.AddStrings("b")

		lc1.AddAnother(lc2)
		if lc1.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedCollectionNode — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_LinkedCollectionNode(t *testing.T) {
	safeTest(t, "Test_Cov14_LinkedCollectionNode", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		node := &corestr.LinkedCollectionNode{Element: c}

		if node.IsEmpty() || !node.HasElement() || node.HasNext() {
			t.Fatal("expected valid")
		}

		_ = node.String()
		_ = node.List()
		_ = node.Join(",")
		_ = node.Clone()
		_ = node.IsEqual(node)

		end, length := node.EndOfChain()
		if end != node || length != 1 {
			t.Fatal("expected self, 1")
		}

		_ = node.CreateLinkedList()
		_ = node.IsEqualValue(c)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CharCollectionMap — basic coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_CharCollectionMap_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov14_CharCollectionMap_BasicOps", func() {
		ccm := corestr.Empty.CharCollectionMap()

		if !ccm.IsEmpty() || ccm.HasItems() {
			t.Fatal("expected empty")
		}

		_ = ccm.Length()
		_ = ccm.AllLengthsSum()
		_ = ccm.GetChar("hello")
		_ = ccm.GetChar("")
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CharHashsetMap — basic coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_CharHashsetMap_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov14_CharHashsetMap_BasicOps", func() {
		chm := corestr.Empty.CharHashsetMap()

		if !chm.IsEmpty() || chm.HasItems() {
			t.Fatal("expected empty")
		}

		_ = chm.Length()
		_ = chm.AllLengthsSum()
		_ = chm.GetChar("hello")
		_ = chm.GetChar("")
		_ = chm.GetCharOf("hello")
		_ = chm.GetCharOf("")
		_ = chm.Has("hello")
		_ = chm.LengthOf('h')
		_ = chm.LengthOfHashsetFromFirstChar("hello")
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValues Creator — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov14_KeyValuesCreator(t *testing.T) {
	safeTest(t, "Test_Cov14_KeyValuesCreator", func() {
		_ = corestr.New.KeyValues.Cap(5)
	})
}
