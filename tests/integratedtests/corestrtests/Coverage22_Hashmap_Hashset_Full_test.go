package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — comprehensive coverage for remaining uncovered methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov22_Hashmap_AddMethods(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_AddMethods", func() {
		h := corestr.New.Hashmap.Cap(10)
		h.AddOrUpdate("a", "1")
		h.AddOrUpdateKeyStrValInt("b", 2)
		h.AddOrUpdateKeyStrValFloat("c", 3.0)
		h.AddOrUpdateKeyStrValFloat64("d", 4.0)
		h.AddOrUpdateKeyStrValAny("e", 5)
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "f", Value: "6"})
		h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "g", Value: "7"})
		h.Set("h", "8")
		h.SetTrim(" i ", " 9 ")
		h.SetBySplitter("=", "key=val")
		h.SetBySplitter("=", "onlykey")
		h.AddOrUpdateLock("j", "10")
		if h.Length() < 10 {
			t.Fatalf("expected at least 10, got %d", h.Length())
		}
	})
}

func Test_Cov22_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_AddOrUpdateHashmap", func() {
		h1 := corestr.New.Hashmap.Cap(2)
		h1.AddOrUpdate("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.AddOrUpdate("b", "2")
		h1.AddOrUpdateHashmap(h2)
		if h1.Length() != 2 {
			t.Fatal("expected 2")
		}
		h1.AddOrUpdateHashmap(nil)
	})
}

func Test_Cov22_Hashmap_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_AddOrUpdateMap", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdateMap(map[string]string{"a": "1"})
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov22_Hashmap_AddsOrUpdates(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_AddsOrUpdates", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddsOrUpdates(corestr.KeyValuePair{Key: "a", Value: "1"})
		h.AddOrUpdateKeyValues(corestr.KeyValuePair{Key: "b", Value: "2"})
		h.AddOrUpdateKeyAnyValues(corestr.KeyAnyValuePair{Key: "c", Value: "3"})
	})
}

func Test_Cov22_Hashmap_AddOrUpdateCollection(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_AddOrUpdateCollection", func() {
		h := corestr.New.Hashmap.Cap(5)
		keys := corestr.New.Collection.Strings([]string{"a", "b"})
		vals := corestr.New.Collection.Strings([]string{"1", "2"})
		h.AddOrUpdateCollection(keys, vals)
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
		// mismatched lengths
		h.AddOrUpdateCollection(keys, corestr.New.Collection.Strings([]string{"1"}))
	})
}

func Test_Cov22_Hashmap_Lookups(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_Lookups", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("a", "1")
		if !h.Has("a") || !h.Contains("a") || !h.ContainsLock("a") {
			t.Fatal("expected found")
		}
		if h.IsKeyMissing("a") || h.IsKeyMissingLock("a") {
			t.Fatal("expected not missing")
		}
		if !h.HasLock("a") || !h.HasWithLock("a") {
			t.Fatal("expected found")
		}
		if !h.HasAll("a") || !h.HasAny("a") {
			t.Fatal("expected found")
		}
		if !h.HasAllStrings("a") {
			t.Fatal("expected found")
		}
		if !h.HasAnyItem() {
			t.Fatal("expected has item")
		}
	})
}

func Test_Cov22_Hashmap_Get(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_Get", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		v, ok := h.Get("a")
		if !ok || v != "1" {
			t.Fatal("wrong")
		}
		v2, ok2 := h.GetValue("a")
		if !ok2 || v2 != "1" {
			t.Fatal("wrong")
		}
	})
}

func Test_Cov22_Hashmap_Remove(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_Remove", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		h.Remove("a")
		if h.Has("a") {
			t.Fatal("expected removed")
		}
		h.AddOrUpdate("b", "2")
		h.RemoveWithLock("b")
	})
}

func Test_Cov22_Hashmap_Keys_Values(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_Keys_Values", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		_ = h.Keys()
		_ = h.AllKeys()
		_ = h.KeysCollection()
		_ = h.KeysLock()
		_ = h.ValuesList()
		_ = h.ValuesListCopyLock()
		_ = h.ValuesCollection()
		_ = h.ValuesCollectionLock()
		_ = h.ValuesHashset()
		_ = h.ValuesHashsetLock()
		_, _ = h.KeysValuesCollection()
		_, _ = h.KeysValuesList()
		_, _ = h.KeysValuesListLock()
		_ = h.KeysValuePairs()
		_ = h.KeysValuePairsCollection()
	})
}

func Test_Cov22_Hashmap_FilterOps(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_FilterOps", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("abc", "1")
		h.AddOrUpdate("def", "2")
		items := h.GetKeysFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, s == "abc", false
		})
		if len(items) != 1 {
			t.Fatal("expected 1")
		}
		col := h.GetKeysFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov22_Hashmap_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_ConcatNew", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.AddOrUpdate("b", "2")
		newH := h.ConcatNew(true, h2)
		if newH.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_Cov22_Hashmap_ConcatNewUsingMaps(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_ConcatNewUsingMaps", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		newH := h.ConcatNewUsingMaps(true, map[string]string{"b": "2"})
		if newH.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_Cov22_Hashmap_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_IsEqual", func() {
		h1 := corestr.New.Hashmap.Cap(2)
		h1.AddOrUpdate("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.AddOrUpdate("a", "1")
		if !h1.IsEqualPtr(h2) {
			t.Fatal("expected equal")
		}
		if !h1.IsEqualPtrLock(h2) {
			t.Fatal("expected equal")
		}
	})
}

func Test_Cov22_Hashmap_Diff(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_Diff", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		h2 := corestr.New.Hashmap.Cap(2)
		h2.AddOrUpdate("b", "2")
		_ = h.Diff(h2)
	})
}

func Test_Cov22_Hashmap_KeysToLower(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_KeysToLower", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("ABC", "1")
		lower := h.KeysToLower()
		if !lower.Has("abc") {
			t.Fatal("expected lowercase key")
		}
	})
}

func Test_Cov22_Hashmap_GetExcept(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_GetExcept", func() {
		h := corestr.New.Hashmap.Cap(3)
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		result := h.GetValuesKeysExcept([]string{"a"})
		if len(result) != 1 {
			t.Fatal("expected 1")
		}
		_ = h.GetAllExceptCollection(nil)
	})
}

func Test_Cov22_Hashmap_Join(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_Join", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		_ = h.Join(",")
		_ = h.JoinKeys(",")
	})
}

func Test_Cov22_Hashmap_String(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_String", func() {
		h := corestr.New.Hashmap.Cap(2)
		_ = h.String()
		h.AddOrUpdate("a", "1")
		_ = h.String()
		_ = h.StringLock()
	})
}

func Test_Cov22_Hashmap_JsonOps(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_JsonOps", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		_ = h.JsonModel()
		_ = h.JsonModelAny()
		b, _ := h.MarshalJSON()
		h2 := corestr.New.Hashmap.Cap(2)
		_ = h2.UnmarshalJSON(b)
		_ = h.Json()
		_ = h.JsonPtr()
		_ = h.AsJsoner()
		_ = h.AsJsonContractsBinder()
		_ = h.AsJsonParseSelfInjector()
		_ = h.AsJsonMarshaller()
	})
}

func Test_Cov22_Hashmap_Clone(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_Clone", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		c := h.Clone()
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
		cp := h.ClonePtr()
		if cp == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_Cov22_Hashmap_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_ClearDispose", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		h.Clear()
		if h.Length() != 0 {
			t.Fatal("expected 0")
		}
		h.AddOrUpdate("b", "2")
		h.Dispose()
	})
}

func Test_Cov22_Hashmap_ToError(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_ToError", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		_ = h.ToError(",")
		_ = h.ToDefaultError()
		_ = h.KeyValStringLines()
	})
}

func Test_Cov22_Hashmap_SerializeDeserialize(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_SerializeDeserialize", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		_, _ = h.Serialize()
		var target map[string]string
		_ = h.Deserialize(&target)
	})
}

func Test_Cov22_Hashmap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_ParseInjectUsingJson", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		jr := h.JsonPtr()
		target := corestr.New.Hashmap.Cap(2)
		_, err := target.ParseInjectUsingJson(jr)
		if err != nil {
			t.Fatal("unexpected")
		}
	})
}

func Test_Cov22_Hashmap_AddOrUpdateWithWgLock(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_AddOrUpdateWithWgLock", func() {
		h := corestr.New.Hashmap.Cap(2)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateWithWgLock("a", "1", wg)
		wg.Wait()
	})
}

func Test_Cov22_Hashmap_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_AddOrUpdateStringsPtrWgLock", func() {
		h := corestr.New.Hashmap.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		h.AddOrUpdateStringsPtrWgLock(wg, []string{"a"}, []string{"1"})
		wg.Wait()
	})
}

func Test_Cov22_Hashmap_SafeItems(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_SafeItems", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		_ = h.SafeItems()
		_ = h.Items()
		_ = h.ItemsCopyLock()
	})
}

func Test_Cov22_Hashmap_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_HasAllCollectionItems", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		c := corestr.New.Collection.Strings([]string{"a"})
		if !h.HasAllCollectionItems(c) {
			t.Fatal("expected true")
		}
	})
}

func Test_Cov22_Hashmap_ToStringsUsingCompiler(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_ToStringsUsingCompiler", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		strs := h.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })
		if len(strs) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov22_Hashmap_AddsOrUpdatesFilters(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_AddsOrUpdatesFilters", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddsOrUpdatesAnyUsingFilter(
			func(p corestr.KeyAnyValuePair) (string, bool, bool) { return "v", true, false },
			corestr.KeyAnyValuePair{Key: "a", Value: "1"},
		)
		h.AddsOrUpdatesAnyUsingFilterLock(
			func(p corestr.KeyAnyValuePair) (string, bool, bool) { return "v", true, false },
			corestr.KeyAnyValuePair{Key: "b", Value: "2"},
		)
		h.AddsOrUpdatesUsingFilter(
			func(p corestr.KeyValuePair) (string, bool, bool) { return "v", true, false },
			corestr.KeyValuePair{Key: "c", Value: "3"},
		)
	})
}

func Test_Cov22_Hashmap_Collection(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashmap_Collection", func() {
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("a", "1")
		c := h.Collection()
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — comprehensive coverage for remaining uncovered methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov22_Hashset_AddMethods(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_AddMethods", func() {
		h := corestr.New.Hashset.Cap(10)
		h.Add("a")
		h.AddNonEmpty("")
		h.AddNonEmpty("b")
		h.AddNonEmptyWhitespace("  ")
		h.AddNonEmptyWhitespace("c")
		h.AddIf(false, "skip")
		h.AddIf(true, "d")
		h.AddIfMany(false, "x", "y")
		h.AddIfMany(true, "e", "f")
		h.AddFunc(func() string { return "g" })
		h.AddLock("h")
		h.AddBool("a") // existing
		h.AddBool("i") // new
		s := "j"
		h.AddPtr(&s)
		h.AddPtrLock(&s)
		h.AddStrings([]string{"k", "l"})
		h.AddStringsLock([]string{"m"})
		h.Adds("n", "o")
	})
}

func Test_Cov22_Hashset_Lookups(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_Lookups", func() {
		h := corestr.New.Hashset.Cap(5)
		h.Add("a")
		if !h.Has("a") || !h.Contains("a") {
			t.Fatal("expected found")
		}
		if h.IsMissing("a") || h.IsMissingLock("a") {
			t.Fatal("expected not missing")
		}
		if !h.HasLock("a") || !h.HasWithLock("a") {
			t.Fatal("expected found")
		}
		if !h.HasAll("a") || !h.HasAny("a") {
			t.Fatal("expected found")
		}
		if !h.HasAllStrings([]string{"a"}) {
			t.Fatal("expected found")
		}
		if !h.HasAnyItem() || !h.HasItems() {
			t.Fatal("expected has item")
		}
		if h.IsAllMissing("a") {
			t.Fatal("expected false")
		}
	})
}

func Test_Cov22_Hashset_Collection(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_Collection", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Adds("a", "b")
		_ = h.Collection()
		_ = h.SortedList()
		_ = h.OrderedList()
		_ = h.SafeStrings()
		_ = h.Lines()
		_ = h.SimpleSlice()
		_ = h.Items()
		_ = h.List()
		_ = h.ListPtr()
		_ = h.ListCopyLock()
		_ = h.ListPtrSortedAsc()
		_ = h.ListPtrSortedDsc()
	})
}

func Test_Cov22_Hashset_Filter(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_Filter", func() {
		h := corestr.New.Hashset.Cap(5)
		h.Adds("abc", "def", "ghi")
		filtered := h.Filter(func(s string) bool { return s == "abc" })
		if filtered.Length() != 1 {
			t.Fatal("expected 1")
		}
		items := h.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, s == "abc", false
		})
		if len(items) != 1 {
			t.Fatal("expected 1")
		}
		col := h.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if col.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_Cov22_Hashset_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_GetAllExcept", func() {
		h := corestr.New.Hashset.Cap(3)
		h.Adds("a", "b", "c")
		r := h.GetAllExcept([]string{"a"})
		if len(r) != 2 {
			t.Fatal("expected 2")
		}
		_ = h.GetAllExcept(nil)
		_ = h.GetAllExceptSpread("a")
		_ = h.GetAllExceptCollection(nil)
		_ = h.GetAllExceptHashset(nil)
	})
}

func Test_Cov22_Hashset_Concat(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_Concat", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		h2 := corestr.New.Hashset.Cap(2)
		h2.Add("b")
		newH := h.ConcatNewHashsets(true, h2)
		if newH.Length() < 2 {
			t.Fatal("expected at least 2")
		}
		_ = h.ConcatNewStrings(true, []string{"c"})
	})
}

func Test_Cov22_Hashset_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_IsEquals", func() {
		h1 := corestr.New.Hashset.Cap(2)
		h1.Add("a")
		h2 := corestr.New.Hashset.Cap(2)
		h2.Add("a")
		if !h1.IsEquals(h2) {
			t.Fatal("expected equal")
		}
		if !h1.IsEqual(h2) {
			t.Fatal("expected equal")
		}
		if !h1.IsEqualsLock(h2) {
			t.Fatal("expected equal")
		}
	})
}

func Test_Cov22_Hashset_Remove(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_Remove", func() {
		h := corestr.New.Hashset.Cap(3)
		h.Adds("a", "b")
		h.Remove("a")
		h.SafeRemove("b")
		h.SafeRemove("missing")
		h.Add("c")
		h.RemoveWithLock("c")
	})
}

func Test_Cov22_Hashset_String(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_String", func() {
		h := corestr.New.Hashset.Cap(2)
		_ = h.String()
		h.Add("a")
		_ = h.String()
		_ = h.StringLock()
	})
}

func Test_Cov22_Hashset_Join(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_Join", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Adds("a", "b")
		_ = h.Join(",")
		_ = h.JoinSorted(",")
		_ = h.JoinLine()
		_ = h.NonEmptyJoins(",")
		_ = h.NonWhitespaceJoins(",")
	})
}

func Test_Cov22_Hashset_JsonOps(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_JsonOps", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		_ = h.JsonModel()
		_ = h.JsonModelAny()
		b, _ := h.MarshalJSON()
		h2 := corestr.New.Hashset.Cap(2)
		_ = h2.UnmarshalJSON(b)
		_ = h.Json()
		_ = h.JsonPtr()
		_ = h.AsJsonContractsBinder()
		_ = h.AsJsoner()
		_ = h.AsJsonParseSelfInjector()
		_ = h.AsJsonMarshaller()
	})
}

func Test_Cov22_Hashset_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_ClearDispose", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		h.Clear()
		if h.Length() != 0 {
			t.Fatal("expected 0")
		}
		h.Add("b")
		h.Dispose()
	})
}

func Test_Cov22_Hashset_Resize(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_Resize", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		h.Resize(100)
		h.ResizeLock(200)
		h.AddCapacities(50)
		h.AddCapacitiesLock(50)
	})
}

func Test_Cov22_Hashset_ToLowerSet(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_ToLowerSet", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("ABC")
		lower := h.ToLowerSet()
		if !lower.Has("abc") {
			t.Fatal("expected lowercase")
		}
	})
}

func Test_Cov22_Hashset_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_DistinctDiff", func() {
		h := corestr.New.Hashset.Cap(3)
		h.Adds("a", "b")
		diff := h.DistinctDiffLinesRaw("b", "c")
		if len(diff) != 2 { // "a" and "c"
			t.Fatalf("expected 2, got %d", len(diff))
		}
		diffMap := h.DistinctDiffLines("b", "c")
		if len(diffMap) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov22_Hashset_DistinctDiffHashset(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_DistinctDiffHashset", func() {
		h1 := corestr.New.Hashset.Cap(2)
		h1.Adds("a", "b")
		h2 := corestr.New.Hashset.Cap(2)
		h2.Adds("b", "c")
		diff := h1.DistinctDiffHashset(h2)
		if len(diff) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov22_Hashset_SerializeDeserialize(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_SerializeDeserialize", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		_, _ = h.Serialize()
		var target map[string]bool
		_ = h.Deserialize(&target)
	})
}

func Test_Cov22_Hashset_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_ParseInjectUsingJson", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		jr := h.JsonPtr()
		target := corestr.New.Hashset.Cap(2)
		_, err := target.ParseInjectUsingJson(jr)
		if err != nil {
			t.Fatal("unexpected")
		}
	})
}

func Test_Cov22_Hashset_WrapQuotes(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_WrapQuotes", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		_ = h.WrapDoubleQuote()
		h2 := corestr.New.Hashset.Cap(2)
		h2.Add("b")
		_ = h2.WrapSingleQuote()
		h3 := corestr.New.Hashset.Cap(2)
		h3.Add("c")
		_ = h3.WrapDoubleQuoteIfMissing()
		h4 := corestr.New.Hashset.Cap(2)
		h4.Add("d")
		_ = h4.WrapSingleQuoteIfMissing()
	})
}

func Test_Cov22_Hashset_MapStringAny(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_MapStringAny", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		m := h.MapStringAny()
		if len(m) != 1 {
			t.Fatal("expected 1")
		}
		_ = h.MapStringAnyDiff()
	})
}

func Test_Cov22_Hashset_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_AddCollection", func() {
		h := corestr.New.Hashset.Cap(5)
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		h.AddCollection(c)
		h.AddCollections(c, nil)
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov22_Hashset_AddSimpleSlice(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_AddSimpleSlice", func() {
		h := corestr.New.Hashset.Cap(5)
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		h.AddSimpleSlice(ss)
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov22_Hashset_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_AddHashsetItems", func() {
		h := corestr.New.Hashset.Cap(5)
		h2 := corestr.New.Hashset.Cap(2)
		h2.Adds("a", "b")
		h.AddHashsetItems(h2)
		h.AddHashsetItems(nil)
	})
}

func Test_Cov22_Hashset_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_AddItemsMap", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddItemsMap(map[string]bool{"a": true, "b": false})
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov22_Hashset_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_HasAllCollectionItems", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Adds("a", "b")
		c := corestr.New.Collection.Strings([]string{"a"})
		if !h.HasAllCollectionItems(c) {
			t.Fatal("expected true")
		}
		if h.HasAllCollectionItems(nil) {
			t.Fatal("expected false")
		}
	})
}

func Test_Cov22_Hashset_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_AddFuncErr", func() {
		h := corestr.New.Hashset.Cap(2)
		h.AddFuncErr(
			func() (string, error) { return "a", nil },
			func(err error) {},
		)
		if h.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov22_Hashset_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_AddsUsingFilter", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddsUsingFilter(
			func(s string, i int) (string, bool, bool) { return s, s != "", false },
			"a", "", "b",
		)
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov22_Hashset_AddsAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_AddsAnyUsingFilter", func() {
		h := corestr.New.Hashset.Cap(5)
		h.AddsAnyUsingFilter(
			func(s string, i int) (string, bool, bool) { return s, true, false },
			"a", nil, "b",
		)
		if h.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov22_Hashset_Transpile(t *testing.T) {
	safeTest(t, "Test_Cov22_Hashset_Transpile", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a")
		result := h.Transpile(func(s string) string { return s + "!" })
		if !result.Has("a!") {
			t.Fatal("expected transpiled")
		}
	})
}
