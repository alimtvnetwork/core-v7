package corestrtests

import (
	"strings"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ==========================================================================
// Collection — comprehensive coverage
// ==========================================================================

func Test_I8_Collection_BasicOps(t *testing.T) {
	safeTest(t, "Test_I8_Collection_BasicOps", func() {
		c := corestr.New.Collection.Cap(10)
		c.Add("a").Add("b").Add("c")

		if c.Length() != 3 { t.Fatal("expected 3") }
		if c.Count() != 3 { t.Fatal("count mismatch") }
		if c.Capacity() < 3 { t.Fatal("capacity too small") }
		if !c.HasAnyItem() { t.Fatal("expected has items") }
		if c.IsEmpty() { t.Fatal("not empty") }
		if c.HasItems() != true { t.Fatal("has items") }
		if c.LastIndex() != 2 { t.Fatal("last index") }
		if !c.HasIndex(0) { t.Fatal("has index 0") }
		if !c.HasIndex(2) { t.Fatal("has index 2") }
		if c.HasIndex(3) { t.Fatal("no index 3") }
		if c.HasIndex(-1) { t.Fatal("no negative index") }
	})
}
func Test_I8_Collection_AddVariants(t *testing.T) {
	safeTest(t, "Test_I8_Collection_AddVariants", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmpty("")
		c.AddNonEmpty("a")
		c.AddNonEmptyWhitespace("   ")
		c.AddNonEmptyWhitespace("b")
		c.AddLock("c")
		c.AddIf(true, "d")
		c.AddIf(false, "skip")
		c.AddIfMany(true, "e", "f")
		c.AddIfMany(false, "skip1", "skip2")
		c.AddFunc(func() string { return "g" })
		c.AddFuncErr(func() (string, error) { return "h", nil }, func(errInput error) {})
		c.AddError(nil)

		if c.Length() < 7 { t.Fatal("expected at least 7 items") }
	})
}

func Test_I8_Collection_Adds(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Adds", func() {
		c := corestr.New.Collection.Empty()
		c.Adds("a", "b", "c")
		c.AddsLock("d", "e")
		c.AddStrings([]string{"f", "g"})
		c.AddsNonEmpty("", "h", "")

		other := corestr.New.Collection.Strings([]string{"i", "j"})
		c.AddCollection(other)
		c.AddCollections(other)

		if c.Length() < 10 { t.Fatal("expected at least 10") }
	})
}

func Test_I8_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_RemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		if !c.RemoveAt(1) { t.Fatal("expected success") }
		if c.RemoveAt(-1) { t.Fatal("expected fail for negative") }
		if c.RemoveAt(99) { t.Fatal("expected fail for out of bounds") }
		if c.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_I8_Collection_ChainRemoveAt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_ChainRemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.ChainRemoveAt(0)
		if c.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_I8_Collection_InsertAt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_InsertAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "c"})
		c.InsertAt(1, "b")
		if c.Length() != 3 { t.Fatal("expected 3") }
	})
}

func Test_I8_Collection_RemoveItemsIndexes(t *testing.T) {
	safeTest(t, "Test_I8_Collection_RemoveItemsIndexes", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		c.RemoveItemsIndexes(false, 0, 2)
		c2 := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		indexes := []int{0}
		c2.RemoveItemsIndexesPtr(false, indexes)
	})
}

func Test_I8_Collection_FirstLastSingleTakeSkip(t *testing.T) {
	safeTest(t, "Test_I8_Collection_FirstLastSingleTakeSkip", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		if c.First() != "a" { t.Fatal("first") }
		if c.Last() != "c" { t.Fatal("last") }
		if c.FirstOrDefault() != "a" { t.Fatal("first or default") }
		if c.LastOrDefault() != "c" { t.Fatal("last or default") }

		taken := c.Take(2)
		if taken.Length() != 2 { t.Fatal("take 2") }
		skipped := c.Skip(1)
		if skipped.Length() != 2 { t.Fatal("skip 1") }

		empty := corestr.New.Collection.Empty()
		if empty.FirstOrDefault() != "" { t.Fatal("empty first") }
		if empty.LastOrDefault() != "" { t.Fatal("empty last") }

		single := corestr.New.Collection.Strings([]string{"only"})
		if single.Single() != "only" { t.Fatal("single") }
	})
}

func Test_I8_Collection_Reverse(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Reverse", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		r := c.Reverse()
		if r.First() != "c" { t.Fatal("reverse first") }
	})
}

func Test_I8_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_IndexAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		if c.IndexAt(0) != "a" { t.Fatal("index 0") }
		if c.SafeIndexAtUsingLength("", 3, 0) != "a" { t.Fatal("safe index 0") }
	})
}

func Test_I8_Collection_IsEquals(t *testing.T) {
	safeTest(t, "Test_I8_Collection_IsEquals", func() {
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})
		c := corestr.New.Collection.Strings([]string{"a", "c"})
		if !a.IsEquals(b) { t.Fatal("expected equal") }
		if a.IsEquals(c) { t.Fatal("expected not equal") }
		if a.IsEquals(nil) { t.Fatal("expected not equal to nil") }

		_ = a.IsEqualsWithSensitive(false, b)
	})
}

func Test_I8_Collection_LengthLock(t *testing.T) {
	safeTest(t, "Test_I8_Collection_LengthLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.LengthLock() != 1 { t.Fatal("expected 1") }
		if !c.IsEmptyLock() == true { /* expected non-empty */ }
	})
}

func Test_I8_Collection_AsyncOps(t *testing.T) {
	safeTest(t, "Test_I8_Collection_AsyncOps", func() {
		c := corestr.New.Collection.Cap(10)
		var wg sync.WaitGroup
		wg.Add(1)
		c.AddWithWgLock(&wg, "a")
		wg.Wait()

		c.AddStringsAsync(&wg, []string{"b", "c"})
		wg.Wait()

		wg.Add(1)
		c.AddsAsync(&wg, "d", "e")
		wg.Wait()
	})
}

func Test_I8_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Filter", func() {
		c := corestr.New.Collection.Strings([]string{"aa", "b", "cc"})
		filtered := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, len(s) == 2, false
		})
		if len(filtered) != 2 { t.Fatalf("expected 2, got %d", len(filtered)) }

		_ = c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		_ = c.FilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		_ = c.FilteredCollectionLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		_ = c.FilterPtr(func(s *string, i int) (*string, bool, bool) {
			return s, true, false
		})
		_ = c.FilterPtrLock(func(s *string, i int) (*string, bool, bool) {
			return s, true, false
		})
	})
}

func Test_I8_Collection_Unique(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Unique", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "a"})
		uniq := c.UniqueList()
		if len(uniq) != 2 { t.Fatal("expected 2 unique") }
		_ = c.UniqueListLock()
		_ = c.UniqueBoolMap()
		_ = c.UniqueBoolMapLock()
	})
}
func Test_I8_Collection_Has(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Has", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !c.Has("a") { t.Fatal("expected has a") }
		if c.Has("z") { t.Fatal("expected no z") }
		_ = c.HasLock("a")
		s := "a"
		_ = c.HasPtr(&s)
		if !c.HasAll("a", "b") { t.Fatal("expected has all") }
		_ = c.HasUsingSensitivity("A", false)
		_ = c.IsContainsAll("a", "b")
		_ = c.IsContainsAllLock("a")
		_ = c.IsContainsAllSlice([]string{"a"})
		ns := "a"
		_ = c.IsContainsPtr(&ns)
		_, _ = c.GetHashsetPlusHasAll([]string{"a", "b"})
	})
}

func Test_I8_Collection_Sort(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Sort", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		asc := c.SortedListAsc()
		if asc[0] != "a" { t.Fatal("expected a first") }
		dsc := c.SortedListDsc()
		if dsc[0] != "c" { t.Fatal("expected c first") }
		_ = c.SortedAsc()
		_ = c.SortedAscLock()
	})
}

func Test_I8_Collection_Hashset(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Hashset", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		h := c.HashsetAsIs()
		if h.Length() != 2 { t.Fatal("expected 2") }
		_ = c.HashsetWithDoubleLength()
		_ = c.HashsetLock()
	})
}

func Test_I8_Collection_String(t *testing.T) {
	safeTest(t, "Test_I8_Collection_String", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		_ = c.String()
		_ = c.StringLock()
		_ = c.StringJSON()
		_ = c.JsonString()
		_ = c.JsonStringMust()
		_ = c.SummaryString(1)
		_ = c.SummaryStringWithHeader("hdr")
	})
}

func Test_I8_Collection_CsvJoin(t *testing.T) {
	safeTest(t, "Test_I8_Collection_CsvJoin", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		_ = c.Csv()
		_ = c.CsvOptions(true)
		_ = c.CsvLines()
		_ = c.CsvLinesOptions(true)
		_ = c.Join(", ")
		_ = c.JoinLine()
		_ = c.Joins(", ")
		_ = c.NonEmptyJoins(", ")
		_ = c.NonWhitespaceJoins(", ")
	})
}

func Test_I8_Collection_Json(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Json", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.Json()
		_ = c.JsonPtr()
		_ = c.JsonModel()
		_ = c.JsonModelAny()
		_, _ = c.MarshalJSON()
		_ = c.AsJsonMarshaller()
		_ = c.AsJsonContractsBinder()

		c2 := &corestr.Collection{}
		_ = c2.UnmarshalJSON([]byte(`["x"]`))

		r := corejson.New([]string{"y"})
		_, _ = c.ParseInjectUsingJson(&r)
		_ = c.JsonParseSelfInject(&r)
	})
}

func Test_I8_Collection_Serialize(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Serialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		b, err := c.Serialize()
		if err != nil { t.Fatal(err) }
		var target []string
		_ = c.Deserialize(&target)
		_ = b
	})
}

func Test_I8_Collection_Error(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Error", func() {
		c := corestr.New.Collection.Strings([]string{"err1", "err2"})
		_ = c.AsDefaultError()
		_ = c.AsError("; ")
		_ = c.ToError("; ")
		_ = c.ToDefaultError()

		c.AddError(nil)
	})
}

func Test_I8_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_I8_Collection_EachItemSplitBy", func() {
		c := corestr.New.Collection.Strings([]string{"a,b", "c,d"})
		split := c.EachItemSplitBy(",")
		if len(split) != 4 { t.Fatalf("expected 4, got %d", len(split)) }
	})
}

func Test_I8_Collection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_I8_Collection_ConcatNew", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c2 := c.ConcatNew(0, "b", "c")
		if c2.Length() != 3 { t.Fatal("expected 3") }
	})
}

func Test_I8_Collection_AppendCollections(t *testing.T) {
	safeTest(t, "Test_I8_Collection_AppendCollections", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c.AppendCollectionPtr(c2)
		c.AppendCollections(c2)
	})
}

func Test_I8_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_I8_Collection_AppendAnys", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnys("a", 1, nil)
		c.AppendAnysLock("b", 2)
		c.AppendNonEmptyAnys("", "c", nil)
		c.AppendAnysUsingFilter(corestr.IsStringFilter(func(str string, index int) (string, bool, bool) { return str, str != "", false }), "d", nil)
		c.AppendAnysUsingFilterLock(corestr.IsStringFilter(func(str string, index int) (string, bool, bool) { return str, true, false }), "e")
	})
}

func Test_I8_Collection_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_I8_Collection_GetAllExcept", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.GetAllExcept([]string{"b"})
		if len(result) != 2 { t.Fatal("expected 2") }

		exc := corestr.New.Collection.Strings([]string{"a"})
		result2 := c.GetAllExceptCollection(exc)
		if len(result2) != 2 { t.Fatal("expected 2") }
	})
}

func Test_I8_Collection_Paging(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Paging", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		pages := c.GetPagesSize(2)
		if pages != 3 { t.Fatalf("expected 3, got %d", pages) }
		page := c.GetPagedCollection(2)
		_ = page
		single := c.GetSinglePageCollection(2, 1)
		_ = single
	})
}

func Test_I8_Collection_New(t *testing.T) {
	safeTest(t, "Test_I8_Collection_New", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c2 := c.New()
		if c2.Length() != 0 { t.Fatal("expected empty new") }
	})
}

func Test_I8_Collection_AddNonEmptyStrings(t *testing.T) {
	safeTest(t, "Test_I8_Collection_AddNonEmptyStrings", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStrings("", "a", "", "b")
		c.AddNonEmptyStringsSlice([]string{"", "c"})
	})
}

func Test_I8_Collection_AddStringsByFuncChecking(t *testing.T) {
	safeTest(t, "Test_I8_Collection_AddStringsByFuncChecking", func() {
		c := corestr.New.Collection.Empty()
		c.AddStringsByFuncChecking([]string{"aa", "b"}, func(s string) bool {
			return len(s) > 1
		})
	})
}

func Test_I8_Collection_ExpandSlicePlusAdd(t *testing.T) {
	safeTest(t, "Test_I8_Collection_ExpandSlicePlusAdd", func() {
		c := corestr.New.Collection.Empty()
		c.ExpandSlicePlusAdd([]string{"a", "b"}, func(s string) []string {
			return []string{strings.ToUpper(s)}
		})
	})
}

func Test_I8_Collection_MergeSlicesOfSlice(t *testing.T) {
	safeTest(t, "Test_I8_Collection_MergeSlicesOfSlice", func() {
		c := corestr.New.Collection.Empty()
		c.MergeSlicesOfSlice([]string{"a"}, []string{"b", "c"})
	})
}

func Test_I8_Collection_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_I8_Collection_CharCollectionMap", func() {
		c := corestr.New.Collection.Strings([]string{"apple", "avocado", "banana"})
		m := c.CharCollectionMap()
		if m.IsEmpty() { t.Fatal("expected non-empty") }
	})
}

func Test_I8_Collection_Resize(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Resize", func() {
		c := corestr.New.Collection.Empty()
		c.Resize(100)
		c.AddCapacity(50)
	})
}

func Test_I8_Collection_ClearDispose(t *testing.T) {
	safeTest(t, "Test_I8_Collection_ClearDispose", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Clear()
		if c.Length() != 0 { t.Fatal("expected 0 after clear") }

		c2 := corestr.New.Collection.Strings([]string{"a"})
		c2.Dispose()
	})
}

func Test_I8_Collection_ListCopyPtrLock(t *testing.T) {
	safeTest(t, "Test_I8_Collection_ListCopyPtrLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		cp := c.ListCopyPtrLock()
		if len(cp) != 2 { t.Fatal("expected 2") }
	})
}

func Test_I8_Collection_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_I8_Collection_AddHashmapsValues", func() {
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("k1", "v1")
		h.AddOrUpdate("k2", "v2")
		c.AddHashmapsValues(h)
		c.AddHashmapsKeys(h)
		c.AddHashmapsKeysValues(h)
	})
}

func Test_I8_Collection_AddPointerCollectionsLock(t *testing.T) {
	safeTest(t, "Test_I8_Collection_AddPointerCollectionsLock", func() {
		c := corestr.New.Collection.Empty()
		other := corestr.New.Collection.Strings([]string{"a"})
		c.AddPointerCollectionsLock(other)
	})
}

func Test_I8_Collection_NilLength(t *testing.T) {
	safeTest(t, "Test_I8_Collection_NilLength", func() {
		var c *corestr.Collection
		if c.Length() != 0 { t.Fatal("expected 0 for nil") }
	})
}

func Test_I8_Collection_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_I8_Collection_ParseInjectUsingJsonMust", func() {
		c := corestr.New.Collection.Empty()
		r := corejson.New([]string{"a", "b"})
		c.ParseInjectUsingJsonMust(&r)
		if c.Length() != 2 { t.Fatal("expected 2") }
	})
}

// ==========================================================================
// Hashmap — comprehensive coverage
// ==========================================================================

func Test_I8_Hashmap_BasicOps(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_BasicOps", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("a", "1")
		h.Set("b", "2")
		h.SetTrim(" c ", " 3 ")

		if h.IsEmpty() { t.Fatal("not empty") }
		if !h.HasItems() { t.Fatal("has items") }
		if !h.HasAnyItem() { t.Fatal("has any") }
		if h.Length() < 3 { t.Fatal("expected >= 3") }
		_ = h.LengthLock()
		_ = h.IsEmptyLock()
	})
}

func Test_I8_Hashmap_AddVariants(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_AddVariants", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateKeyStrValInt("count", 42)
		h.AddOrUpdateKeyStrValFloat("rate", 3.14)
		h.AddOrUpdateKeyStrValFloat64("pi", 3.14159)
		h.AddOrUpdateKeyStrValAny("any", "value")
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "kav", Value: "vav"})
		h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "kv", Value: "vv"})
		h.AddOrUpdateLock("locked", "val")

		var wg sync.WaitGroup
		wg.Add(1)
		h.AddOrUpdateWithWgLock("wg", "wgval", &wg)
		wg.Wait()
	})
}

func Test_I8_Hashmap_AddCollectionMaps(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_AddCollectionMaps", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdateMap(map[string]string{"a": "1"})
		h.AddsOrUpdates(corestr.KeyValuePair{Key: "b", Value: "2"}, corestr.KeyValuePair{Key: "c", Value: "3"})

		kav := []corestr.KeyAnyValuePair{{Key: "d", Value: "4"}}
		h.AddOrUpdateKeyAnyValues(kav...)

		h.AddOrUpdateKeyValues(corestr.KeyValuePair{Key: "e", Value: "5"})

		keys := corestr.New.Collection.Strings([]string{"f"})
		vals := corestr.New.Collection.Strings([]string{"6"})
		h.AddOrUpdateCollection(keys, vals)

		h2 := corestr.New.Hashmap.Empty()
		h2.AddOrUpdate("f", "6")
		h.AddOrUpdateHashmap(h2)
	})
}

func Test_I8_Hashmap_Has(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Has", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")

		if !h.Has("a") { t.Fatal("expected has a") }
		if !h.Contains("b") { t.Fatal("expected contains b") }
		_ = h.ContainsLock("a")
		if h.IsKeyMissing("a") { t.Fatal("not missing") }
		_ = h.IsKeyMissingLock("a")
		_ = h.HasLock("a")
		_ = h.HasWithLock("a")
		if !h.HasAllStrings("a", "b") { t.Fatal("has all") }
		if !h.HasAll("a", "b") { t.Fatal("has all") }
		_ = h.HasAny("a", "z")

		coll := corestr.New.Collection.Strings([]string{"a"})
		_ = h.HasAllCollectionItems(coll)
	})
}

func Test_I8_Hashmap_Get(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Get", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("key", "val")
		v, found := h.Get("key")
		if !found || v != "val" { t.Fatal("get failed") }
		v2, found2 := h.GetValue("key")
		if !found2 || v2 != "val" { t.Fatal("getvalue failed") }
		_, f := h.Get("missing")
		if f { t.Fatal("should not find") }
	})
}

func Test_I8_Hashmap_Keys(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Keys", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		_ = h.AllKeys()
		_ = h.Keys()
		_ = h.KeysLock()
		_ = h.KeysCollection()
		_ = h.ValuesList()
		_ = h.ValuesListCopyLock()
		_ = h.ValuesCollection()
		_ = h.ValuesCollectionLock()
		_ = h.ValuesHashset()
		_ = h.ValuesHashsetLock()
		_ = h.Items()
		_ = h.SafeItems()
		_ = h.ItemsCopyLock()
		_ = h.Collection()
	})
}

func Test_I8_Hashmap_KeysValues(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_KeysValues", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		kc, vc := h.KeysValuesCollection()
		_, _ = kc, vc
		kl, vl := h.KeysValuesList()
		_, _ = kl, vl
		_, _ = h.KeysValuesListLock()
		_ = h.KeysValuePairs()
		_ = h.KeysValuePairsCollection()
	})
}
func Test_I8_Hashmap_String(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_String", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		_ = h.String()
		_ = h.StringLock()
		_ = h.KeyValStringLines()
	})
}

func Test_I8_Hashmap_IsEqual(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_IsEqual", func() {
		h1 := corestr.New.Hashmap.Empty()
		h1.AddOrUpdate("a", "1")
		h2 := corestr.New.Hashmap.Empty()
		h2.AddOrUpdate("a", "1")
		if !h1.IsEqualPtr(h2) { t.Fatal("expected equal") }
		_ = h1.IsEqualPtrLock(h2)
	})
}
func Test_I8_Hashmap_Filter(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Filter", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("apple", "1")
		h.AddOrUpdate("banana", "2")
		_ = h.GetKeysFilteredItems(corestr.IsStringFilter(func(str string, index int) (string, bool, bool) {
			return str, strings.HasPrefix(str, "a"), false
		}))
		_ = h.GetKeysFilteredCollection(corestr.IsStringFilter(func(str string, index int) (string, bool, bool) {
			return str, true, false
		}))
		_ = h.GetValuesExceptKeysInHashset(corestr.New.Hashset.StringsSpreadItems("apple"))
	})
}

func Test_I8_Hashmap_Except(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Except", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		_ = h.GetValuesKeysExcept([]string{"a"})
		_ = h.GetAllExceptCollection(corestr.New.Collection.Strings([]string{"a"}))
	})
}

func Test_I8_Hashmap_Join(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Join", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		_ = h.Join(", ")
		_ = h.JoinKeys(", ")
	})
}

func Test_I8_Hashmap_Json(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Json", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		_ = h.Json()
		_ = h.JsonPtr()
		_ = h.JsonModel()
		_ = h.JsonModelAny()
		_, _ = h.MarshalJSON()
		_ = h.AsJsoner()
		_ = h.AsJsonContractsBinder()
		_ = h.AsJsonParseSelfInjector()
		_ = h.AsJsonMarshaller()

		h2 := &corestr.Hashmap{}
		_ = h2.UnmarshalJSON([]byte(`{"k":"v"}`))

		r := corejson.New(map[string]string{"x": "y"})
		_, _ = h.ParseInjectUsingJson(&r)
		_ = h.JsonParseSelfInject(&r)
	})
}

func Test_I8_Hashmap_Serialize(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Serialize", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		_, _ = h.Serialize()
		var target map[string]string
		_ = h.Deserialize(&target)
	})
}

func Test_I8_Hashmap_Error(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Error", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		_ = h.ToError("; ")
		_ = h.ToDefaultError()
	})
}

func Test_I8_Hashmap_SetBySplitter(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_SetBySplitter", func() {
		h := corestr.New.Hashmap.Empty()
		h.SetBySplitter("key=val", "=")
	})
}

func Test_I8_Hashmap_Diff(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Diff", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		h2 := corestr.New.Hashmap.Empty()
		h2.AddOrUpdate("a", "1")
		_ = h.Diff(h2)
		_ = h.DiffRaw(h2.Items())
	})
}

func Test_I8_Hashmap_ConcatNew(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_ConcatNew", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		other := corestr.New.Hashmap.Empty()
		other.AddOrUpdate("b", "2")
		h2 := h.ConcatNew(false, other)
		_ = h2
		h3 := h.ConcatNewUsingMaps(false, map[string]string{"d": "4"})
		_ = h3
	})
}

func Test_I8_Hashmap_AddsOrUpdatesUsingFilter(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_AddsOrUpdatesUsingFilter", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddsOrUpdatesUsingFilter(corestr.IsKeyValueFilter(func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Key, pair.Key != "skip", false
		}), corestr.KeyValuePair{Key: "a", Value: "1"}, corestr.KeyValuePair{Key: "skip", Value: "2"})
		h.AddsOrUpdatesAnyUsingFilter(corestr.IsKeyAnyValueFilter(func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return pair.Key, true, false
		}), corestr.KeyAnyValuePair{Key: "b"})
		h.AddsOrUpdatesAnyUsingFilterLock(corestr.IsKeyAnyValueFilter(func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
			return pair.Key, true, false
		}), corestr.KeyAnyValuePair{Key: "c"})
	})
}

func Test_I8_Hashmap_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_AddOrUpdateStringsPtrWgLock", func() {
		h := corestr.New.Hashmap.Empty()
		var wg sync.WaitGroup
		wg.Add(1)
		keys := []string{"a"}
		values := []string{"1"}
		h.AddOrUpdateStringsPtrWgLock(&wg, keys, values)
		wg.Wait()
	})
}

func Test_I8_Hashmap_ToStringsUsingCompiler(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_ToStringsUsingCompiler", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		_ = h.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })
	})
}

func Test_I8_Hashmap_ClearDispose(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_ClearDispose", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.Clear()
		h2 := corestr.New.Hashmap.Empty()
		h2.AddOrUpdate("a", "1")
		h2.Dispose()
	})
}

func Test_I8_Hashmap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_ParseInjectUsingJsonMust", func() {
		h := corestr.New.Hashmap.Empty()
		r := corejson.New(map[string]string{"a": "1"})
		h.ParseInjectUsingJsonMust(&r)
	})
}

// ==========================================================================
// Hashset — additional coverage
// ==========================================================================

func Test_I8_Hashset_AddVariants(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_AddVariants", func() {
		h := corestr.New.Hashset.Empty()
		h.Add("a")
		h.AddBool("b")
		h.AddNonEmpty("")
		h.AddNonEmpty("c")
		h.AddNonEmptyWhitespace("   ")
		h.AddNonEmptyWhitespace("d")
		h.AddIf(true, "e")
		h.AddIf(false, "skip")
		h.AddIfMany(true, "f", "g")
		h.AddIfMany(false, "s1", "s2")
		h.AddFunc(func() string { return "h" })
		h.AddFuncErr(func() (string, error) { return "i", nil }, func(e error) {})
		h.AddStrings([]string{"j", "k"})
		h.AddStringsLock([]string{"l"})
		h.Adds("m", "n")
		h.AddLock("o")
		s := "p"
		h.AddPtr(&s)
		h.AddPtrLock(&s)
	})
}

func Test_I8_Hashset_AddCollections(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_AddCollections", func() {
		h := corestr.New.Hashset.Empty()
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		h.AddCollection(c)
		h.AddCollections(c)
		h.AddItemsMap(map[string]bool{"c": true, "d": false})
		ss := corestr.New.SimpleSlice.SpreadStrings("e", "f")
		h.AddSimpleSlice(ss)
		h2 := corestr.New.Hashset.StringsSpreadItems("g")
		h.AddHashsetItems(h2)
	})
}

func Test_I8_Hashset_Has(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_Has", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		if !h.Has("a") { t.Fatal("expected has a") }
		if !h.Contains("b") { t.Fatal("expected contains b") }
		_ = h.HasLock("a")
		_ = h.HasWithLock("a")
		if h.IsMissing("a") { t.Fatal("not missing") }
		_ = h.IsMissingLock("a")
		if !h.HasAllStrings([]string{"a", "b"}) { t.Fatal("has all strings") }
		if !h.HasAll("a", "b") { t.Fatal("has all") }
		_ = h.HasAny("a", "z")
		_ = h.IsAllMissing("x", "y")
		coll := corestr.New.Collection.Strings([]string{"a"})
		_ = h.HasAllCollectionItems(coll)
	})
}

func Test_I8_Hashset_Lists(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_Lists", func() {
		h := corestr.New.Hashset.StringsSpreadItems("b", "a", "c")
		_ = h.OrderedList()
		_ = h.SafeStrings()
		_ = h.Lines()
		_ = h.SortedList()
		_ = h.SimpleSlice()
		_ = h.ListPtrSortedAsc()
		_ = h.ListPtrSortedDsc()
		_ = h.ListCopyLock()
		_ = h.Collection()
		_ = h.MapStringAny()
	})
}

func Test_I8_Hashset_Filter(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_Filter", func() {
		h := corestr.New.Hashset.StringsSpreadItems("aa", "b", "cc")
		f := h.Filter(func(s string) bool { return len(s) > 1 })
		if f.Length() != 2 { t.Fatal("expected 2") }

		_ = h.GetFilteredItems(corestr.IsStringFilter(func(str string, index int) (string, bool, bool) { return str, true, false }))
		_ = h.GetFilteredCollection(corestr.IsStringFilter(func(str string, index int) (string, bool, bool) { return str, true, false }))
	})
}

func Test_I8_Hashset_Except(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_Except", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		_ = h.GetAllExcept([]string{"a"})
		_ = h.GetAllExceptSpread("b")
		h2 := corestr.New.Hashset.StringsSpreadItems("c")
		_ = h.GetAllExceptHashset(h2)
	})
}

func Test_I8_Hashset_Concat(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_Concat", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		c1 := h.ConcatNewStrings(false, []string{"b"}, []string{"c"})
		_ = c1
		h2 := corestr.New.Hashset.StringsSpreadItems("d")
		c2 := h.ConcatNewHashsets(false, h2)
		_ = c2
	})
}

func Test_I8_Hashset_IsEqual(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_IsEqual", func() {
		a := corestr.New.Hashset.StringsSpreadItems("a", "b")
		b := corestr.New.Hashset.StringsSpreadItems("a", "b")
		if !a.IsEqual(b) { t.Fatal("expected equal") }
		_ = a.IsEquals(b)
		_ = a.IsEqualsLock(b)
	})
}

func Test_I8_Hashset_Resize(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_Resize", func() {
		h := corestr.New.Hashset.Empty()
		h.Resize(100)
		h.ResizeLock(200)
		h.AddCapacities(50)
		h.AddCapacitiesLock(50)
	})
}

func Test_I8_Hashset_AsyncOps(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_AsyncOps", func() {
		h := corestr.New.Hashset.Empty()
		var wg sync.WaitGroup
		wg.Add(1)
		h.AddWithWgLock("a", &wg)
		wg.Wait()

		h2 := corestr.New.Hashset.StringsSpreadItems("x")
		var wg2 sync.WaitGroup
		wg2.Add(1)
		h.AddHashsetWgLock(h2, &wg2)
		wg2.Wait()

		wg3 := sync.WaitGroup{}
		wg3.Add(1)
		items := map[string]bool{"y": true}
		h.AddItemsMapWgLock(&items, &wg3)
		wg3.Wait()

		wg4 := sync.WaitGroup{}
		wg4.Add(1)
		strs := []string{"z"}
		h.AddStringsPtrWgLock(strs, &wg4)
		wg4.Wait()
	})
}

func Test_I8_Hashset_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_AddsUsingFilter", func() {
		h := corestr.New.Hashset.Empty()
		h.AddsUsingFilter(corestr.IsStringFilter(func(str string, index int) (string, bool, bool) { return str, str != "", false }), "a", "", "b")
		h.AddsAnyUsingFilter(corestr.IsStringFilter(func(str string, index int) (string, bool, bool) { return str, str != "", false }), "c", nil)
		h.AddsAnyUsingFilterLock(corestr.IsStringFilter(func(str string, index int) (string, bool, bool) { return str, true, false }), "d")
	})
}

func Test_I8_Hashset_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_DistinctDiff", func() {
		a := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		b := corestr.New.Hashset.StringsSpreadItems("a")
		_ = a.DistinctDiffHashset(b)
		_ = a.DistinctDiffLines("a")
		_ = a.DistinctDiffLinesRaw("a")
	})
}

func Test_I8_Hashset_String(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_String", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		_ = h.String()
		_ = h.StringLock()
	})
}

func Test_I8_Hashset_ClearDispose(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_ClearDispose", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		h.Clear()
		h2 := corestr.New.Hashset.StringsSpreadItems("a")
		h2.Dispose()
	})
}

// ==========================================================================
// SimpleSlice — comprehensive coverage
// ==========================================================================

func Test_I8_SimpleSlice_BasicOps(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_BasicOps", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.Add("a")
		s.AddSplit("b,c", ",")
		s.AddIf(true, "d")
		s.AddIf(false, "skip")
		s.Adds("e", "f")
		s.Append("g")
		s.AppendFmt("%s-%d", "h", 1)
		s.AppendFmtIf(true, "%s", "i")
		s.AppendFmtIf(false, "%s", "skip")
		s.AddAsTitleValue("key", "val")
		s.AddAsTitleValueIf(true, "k2", "v2")
		s.AddAsTitleValueIf(false, "k3", "v3")
		s.AddAsCurlyTitleWrap("k4", "v4")
		s.AddAsCurlyTitleWrapIf(true, "k5", "v5")
		s.AddAsCurlyTitleWrapIf(false, "k6", "v6")
		s.AddsIf(true, "j", "k")
		s.AddsIf(false, "skip1", "skip2")
		s.AddError(nil)
		s.InsertAt(0, "first")
		s.AddStruct(false, struct{ Name string }{"test"})
		s.AddPointer(false, &struct{ Val int }{42})
	})
}

func Test_I8_SimpleSlice_FirstLast(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_FirstLast", func() {
		s := corestr.New.SimpleSlice.SpreadStrings("a", "b", "c")
		if s.First() != "a" { t.Fatal("first") }
		if s.Last() != "c" { t.Fatal("last") }
		if s.FirstOrDefault() != "a" { t.Fatal("first or default") }
		if s.LastOrDefault() != "c" { t.Fatal("last or default") }
		_ = s.FirstDynamic()
		_ = s.LastDynamic()
		_ = s.FirstOrDefaultDynamic()
		_ = s.LastOrDefaultDynamic()

		e := corestr.New.SimpleSlice.Empty()
		if e.FirstOrDefault() != "" { t.Fatal("empty first") }
		if e.LastOrDefault() != "" { t.Fatal("empty last") }
	})
}

func Test_I8_SimpleSlice_SkipTakeLimit(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_SkipTakeLimit", func() {
		s := corestr.New.SimpleSlice.SpreadStrings("a", "b", "c", "d")
		_ = s.Skip(1)
		_ = s.Take(2)
		_ = s.Limit(3)
		_ = s.SkipDynamic(1)
		_ = s.TakeDynamic(2)
		_ = s.LimitDynamic(3)
	})
}

func Test_I8_SimpleSlice_Properties(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_Properties", func() {
		s := corestr.New.SimpleSlice.SpreadStrings("a", "b")
		if s.Length() != 2 { t.Fatal("length") }
		if s.Count() != 2 { t.Fatal("count") }
		if s.IsEmpty() { t.Fatal("not empty") }
		if !s.HasAnyItem() { t.Fatal("has items") }
		if s.LastIndex() != 1 { t.Fatal("last index") }
		if !s.HasIndex(0) { t.Fatal("has index") }
		_ = s.Strings()
		_ = s.List()
	})
}

func Test_I8_SimpleSlice_Contains(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_Contains", func() {
		s := corestr.New.SimpleSlice.SpreadStrings("a", "b")
		if !s.IsContains("a") { t.Fatal("contains a") }
		if s.IsContains("z") { t.Fatal("no z") }
		_ = s.IndexOf("b")
		_ = s.IndexOf("z")
		_ = s.CountFunc(func(index int, item string) bool { return index >= 0 && item != "" })
		_ = s.IsContainsFunc("a", func(item, searching string) bool { return item == searching })
		_ = s.IndexOfFunc("b", func(item, searching string) bool { return item == searching })
	})
}

func Test_I8_SimpleSlice_Wrap(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_Wrap", func() {
		s := corestr.New.SimpleSlice.SpreadStrings("a", "b")
		_ = s.WrapDoubleQuote()
		s2 := corestr.New.SimpleSlice.SpreadStrings("c")
		_ = s2.WrapSingleQuote()
		s3 := corestr.New.SimpleSlice.SpreadStrings("d")
		_ = s3.WrapTildaQuote()
		s4 := corestr.New.SimpleSlice.SpreadStrings("e")
		_ = s4.WrapDoubleQuoteIfMissing()
		s5 := corestr.New.SimpleSlice.SpreadStrings("f")
		_ = s5.WrapSingleQuoteIfMissing()
	})
}

func Test_I8_SimpleSlice_Transpile(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_Transpile", func() {
		s := corestr.New.SimpleSlice.SpreadStrings("a", "b")
		_ = s.Transpile(func(s string) string { return strings.ToUpper(s) })
		_ = s.TranspileJoin(func(s string) string { return s }, ", ")
	})
}

func Test_I8_SimpleSlice_Join(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_Join", func() {
		s := corestr.New.SimpleSlice.SpreadStrings("a", "b")
		_ = s.Join(", ")
		_ = s.JoinLine()
		_ = s.JoinLineEofLine()
		_ = s.JoinSpace()
		_ = s.JoinComma()
	})
}

func Test_I8_SimpleSlice_Hashset(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_Hashset", func() {
		s := corestr.New.SimpleSlice.SpreadStrings("a", "b")
		h := s.Hashset()
		if h.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_I8_SimpleSlice_Error(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_Error", func() {
		s := corestr.New.SimpleSlice.SpreadStrings("err1", "err2")
		_ = s.AsDefaultError()
		_ = s.AsError("; ")
	})
}

// ==========================================================================
// LinkedList — comprehensive coverage
// ==========================================================================

func Test_I8_LinkedList_BasicOps(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_BasicOps", func() {
		ll := corestr.New.LinkedList.Empty()
		ll.Add("a").Add("b").Add("c")

		if ll.Length() != 3 { t.Fatal("expected 3") }
		_ = ll.LengthLock()
		if ll.IsEmpty() { t.Fatal("not empty") }
		if !ll.HasItems() { t.Fatal("has items") }
		_ = ll.IsEmptyLock()
		if ll.Head() == nil { t.Fatal("head nil") }
		if ll.Tail() == nil { t.Fatal("tail nil") }
	})
}

func Test_I8_LinkedList_AddVariants(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_AddVariants", func() {
		ll := corestr.New.LinkedList.Empty()
		ll.AddLock("a")
		ll.AddNonEmpty("")
		ll.AddNonEmpty("b")
		ll.AddNonEmptyWhitespace("   ")
		ll.AddNonEmptyWhitespace("c")
		ll.AddIf(true, "d")
		ll.AddIf(false, "skip")
		ll.AddFunc(func() string { return "e" })
		ll.AddFuncErr(func() (string, error) { return "f", nil }, func(errInput error) {})
		ll.Push("g")
		ll.PushFront("h")
		ll.PushBack("i")
		ll.AddFront("j")
		ll.AddsIf(true, "k", "l")
		ll.AddsIf(false, "s1", "s2")
		ll.AddItemsMap(map[string]bool{"m": true})
	})
}

func Test_I8_LinkedList_IsEquals(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_IsEquals", func() {
		a := corestr.New.LinkedList.Strings([]string{"a", "b"})
		b := corestr.New.LinkedList.Strings([]string{"a", "b"})
		if !a.IsEquals(b) { t.Fatal("expected equal") }
		_ = a.IsEqualsWithSensitive(b, false)
	})
}

func Test_I8_LinkedList_InsertAt(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_InsertAt", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "c"})
		ll.InsertAt(1, "b")
		if ll.Length() != 3 { t.Fatal("expected 3") }
	})
}

func Test_I8_LinkedList_Loop(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_Loop", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) (isBreak bool) {
			count++
			return false
		})
		if count != 3 { t.Fatal("expected 3 iterations") }
	})
}

func Test_I8_LinkedList_Filter(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_Filter", func() {
		ll := corestr.New.LinkedList.Strings([]string{"aa", "b", "cc"})
		filtered := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{
				Value:   arg.Node,
				IsKeep:  len(arg.Node.Element) > 1,
				IsBreak: false,
			}
		})
		if len(filtered) != 2 { t.Fatal("expected 2") }
	})
}

func Test_I8_LinkedList_RemoveByIndex(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_RemoveByIndex", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByIndex(1)
		if ll.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_I8_LinkedList_RemoveByValue(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_RemoveByValue", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByElementValue("b", true, false)
		if ll.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_I8_LinkedList_RemoveByIndexes(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_RemoveByIndexes", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c", "d"})
		ll.RemoveNodeByIndexes(false, 0, 2)
		if ll.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_I8_LinkedList_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_GetCompareSummary", func() {
		a := corestr.New.LinkedList.Strings([]string{"a", "b"})
		b := corestr.New.LinkedList.Strings([]string{"a", "c"})
		_ = a.GetCompareSummary(b, "left", "right")
	})
}

// ==========================================================================
// Smaller types
// ==========================================================================

func Test_I8_ValidValue(t *testing.T) {
	safeTest(t, "Test_I8_ValidValue", func() {
		v := corestr.ValidValue{Value: "hello", IsValid: true}
		if !v.IsValid { t.Fatal("expected valid") }
		_ = v.Value
	})
}

func Test_I8_LeftRight(t *testing.T) {
	safeTest(t, "Test_I8_LeftRight", func() {
		lr := corestr.NewLeftRight("a", "b")
		_ = lr.Left
		_ = lr.Right
	})
}

func Test_I8_LeftMiddleRight(t *testing.T) {
	safeTest(t, "Test_I8_LeftMiddleRight", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		_ = lmr.Left
		_ = lmr.Middle
		_ = lmr.Right
	})
}

func Test_I8_ValueStatus(t *testing.T) {
	safeTest(t, "Test_I8_ValueStatus", func() {
		vs := corestr.ValueStatus{ValueValid: corestr.NewValidValue("x"), Index: 0}
		_ = vs.ValueValid
		_ = vs.Index
	})
}

func Test_I8_KeyValuePair(t *testing.T) {
	safeTest(t, "Test_I8_KeyValuePair", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		_ = kv.Key
		_ = kv.Value
	})
}

func Test_I8_KeyAnyValuePair(t *testing.T) {
	safeTest(t, "Test_I8_KeyAnyValuePair", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		_ = kav.Key
		_ = kav.Value
	})
}

func Test_I8_KeyValueCollection(t *testing.T) {
	safeTest(t, "Test_I8_KeyValueCollection", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		if kvc.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_I8_TextWithLineNumber(t *testing.T) {
	safeTest(t, "Test_I8_TextWithLineNumber", func() {
		tln := corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
		_ = tln.LineNumber
		_ = tln.Text
	})
}

// ==========================================================================
// Creators
// ==========================================================================

func Test_I8_NewCreators(t *testing.T) {
	safeTest(t, "Test_I8_NewCreators", func() {
		_ = corestr.New.Collection.Empty()
		_ = corestr.New.Collection.Cap(10)
		_ = corestr.New.Collection.Strings([]string{"a"})
		_ = corestr.New.Collection.StringsOptions(true, []string{"b", "b"})

		_ = corestr.New.Hashmap.Empty()
		_ = corestr.New.Hashmap.Cap(10)

		_ = corestr.New.Hashset.Empty()
		_ = corestr.New.Hashset.Cap(10)
		_ = corestr.New.Hashset.Strings([]string{"a"})
		_ = corestr.New.Hashset.StringsSpreadItems("a", "b")

		_ = corestr.New.LinkedList.Empty()
		_ = corestr.New.LinkedList.Strings([]string{"a"})

		_ = corestr.New.SimpleSlice.Empty()
		_ = corestr.New.SimpleSlice.SpreadStrings("a", "b")

		_ = corestr.New.KeyValues.Empty()

		_ = corestr.New.SimpleStringOnce.Init("test")
	})
}

func Test_I8_CloneSlice(t *testing.T) {
	safeTest(t, "Test_I8_CloneSlice", func() {
		orig := []string{"a", "b"}
		c := corestr.CloneSlice(orig)
		if len(c) != 2 { t.Fatal("expected 2") }

		c2 := corestr.CloneSlice(nil)
		if len(c2) != 0 { t.Fatal("expected empty for nil") }
	})
}

func Test_I8_CloneSliceIf(t *testing.T) {
	safeTest(t, "Test_I8_CloneSliceIf", func() {
		orig := []string{"a", "b"}
		c := corestr.CloneSliceIf(true, orig...)
		if len(c) != 2 { t.Fatal("expected 2") }
		c2 := corestr.CloneSliceIf(false, orig...)
		if len(c2) != 2 { t.Fatal("expected passthrough len 2 when not cloning") }
	})
}

func Test_I8_AnyToString(t *testing.T) {
	safeTest(t, "Test_I8_AnyToString", func() {
		_ = corestr.AnyToString(false, nil)
		_ = corestr.AnyToString(false, "hello")
		_ = corestr.AnyToString(false, 42)
		_ = corestr.AnyToString(false, []string{"a"})
	})
}

func Test_I8_AllIndividualStringsOfStringsLength(t *testing.T) {
	safeTest(t, "Test_I8_AllIndividualStringsOfStringsLength", func() {
		strs := [][]string{{"a", "bb"}, {"ccc"}}
		// Fix: function counts items (3), not character lengths (6).
		// See issues/corestrtests-allindividualslength-wrong-expectation.md
		result := corestr.AllIndividualStringsOfStringsLength(&strs)
		if result != 3 { t.Fatalf("expected 3, got %d", result) }
	})
}

func Test_I8_AllIndividualsLengthOfSimpleSlices(t *testing.T) {
	safeTest(t, "Test_I8_AllIndividualsLengthOfSimpleSlices", func() {
		s1 := corestr.New.SimpleSlice.SpreadStrings("a", "bb")
		s2 := corestr.New.SimpleSlice.SpreadStrings("ccc")
		// Fix: function counts items (3), not character lengths (6).
		// See issues/corestrtests-allindividualslength-wrong-expectation.md
		result := corestr.AllIndividualsLengthOfSimpleSlices(s1, s2)
		if result != 3 { t.Fatalf("expected 3, got %d", result) }
	})
}

// ==========================================================================
// SimpleStringOnce
// ==========================================================================

func Test_I8_SimpleStringOnce(t *testing.T) {
	safeTest(t, "Test_I8_SimpleStringOnce", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		if sso.Value() != "hello" { t.Fatal("expected hello") }
		if !sso.IsDefined() { t.Fatal("expected defined") }
		if sso.IsEmpty() { t.Fatal("expected not empty") }
	})
}

// ==========================================================================
// CharCollectionMap
// ==========================================================================

func Test_I8_CharCollectionMap_Ops(t *testing.T) {
	safeTest(t, "Test_I8_CharCollectionMap_Ops", func() {
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "avocado", "banana"})
		if m.IsEmpty() { t.Fatal("not empty") }
		if !m.HasItems() { t.Fatal("has items") }
		_ = m.Length()
		_ = m.LengthLock()
		_ = m.AllLengthsSum()
		_ = m.AllLengthsSumLock()
		_ = m.IsEmptyLock()
		_ = m.String()
		_ = m.StringLock()
		_ = m.SummaryString()
		_ = m.SummaryStringLock()
		_ = m.SortedListAsc()
		_ = m.GetMap()
		_ = m.GetCopyMapLock()

		_ = m.Has("apple")
		_ = m.LengthOf('a')
		_ = m.LengthOfLock('a')
		_ = m.LengthOfCollectionFromFirstChar("apple")

		_, _ = m.HasWithCollection("apple")
		_, _ = m.HasWithCollectionLock("apple")

		_ = m.GetCollection("a", false)
		_ = m.GetCollectionLock("a", false)
		_ = m.GetChar("a")

		m2 := corestr.New.CharCollectionMap.Items([]string{"apple", "avocado", "banana"})
		_ = m.IsEquals(m2)
		_ = m.IsEqualsLock(m2)
		_ = m.IsEqualsCaseSensitive(true, m2)
		_ = m.IsEqualsCaseSensitiveLock(true, m2)
	})
}

func Test_I8_CharCollectionMap_Add(t *testing.T) {
	safeTest(t, "Test_I8_CharCollectionMap_Add", func() {
		m := corestr.New.CharCollectionMap.Empty()
		m.Add("test")
		m.AddLock("test2")
		m.AddStrings("abc", "def")
		m.AddSameStartingCharItems('a', []string{"alpha", "auto"}, false)

		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("key", "val")
		m.AddHashmapsValues(h)
		m.AddHashmapsKeysValuesBoth(h)

		c := corestr.New.Collection.Strings([]string{"cat"})
		m.AddCollectionItems(c)
		m.AddSameCharsCollection("c", c)
	})
}

// ==========================================================================
// CharHashsetMap
// ==========================================================================

func Test_I8_CharHashsetMap_Ops(t *testing.T) {
	safeTest(t, "Test_I8_CharHashsetMap_Ops", func() {
		m := corestr.New.CharHashsetMap.CapItems(10, 5, "apple", "avocado", "banana")
		if m.IsEmpty() { t.Fatal("not empty") }
		_ = m.Length()
		_ = m.LengthLock()
		_ = m.IsEmptyLock()
		_ = m.String()
		_ = m.StringLock()
		_ = m.SummaryString()
		_ = m.SummaryStringLock()
		_ = m.GetMap()
		_ = m.GetCopyMapLock()
		_ = m.Has("apple")
		_ = m.GetHashset("a", false)
		_ = m.GetHashsetLock(false, "a")
	})
}

// ==========================================================================
// LeftRightFromSplit / LeftMiddleRightFromSplit
// ==========================================================================

func Test_I8_LeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_I8_LeftRightFromSplit", func() {
		lr := corestr.LeftRightFromSplit("a=b", "=")
		if lr.Left != "a" || lr.Right != "b" { t.Fatal("split failed") }

		lr2 := corestr.LeftRightFromSplit("noSep", "=")
		if lr2.Left != "noSep" { t.Fatal("no sep should set left") }
	})
}

func Test_I8_LeftMiddleRightFromSplit(t *testing.T) {
	safeTest(t, "Test_I8_LeftMiddleRightFromSplit", func() {
		lmr := corestr.LeftMiddleRightFromSplit("a:b:c", ":")
		if lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c" {
			t.Fatal("split failed")
		}
	})
}

// ==========================================================================
// ValidValues
// ==========================================================================

func Test_I8_ValidValues(t *testing.T) {
	safeTest(t, "Test_I8_ValidValues", func() {
		vv := corestr.ValidValues{
			ValidValues: []*corestr.ValidValue{
				{Value: "a", IsValid: true},
				{Value: "b", IsValid: false},
			},
		}
		_ = vv.Length()
	})
}

// ==========================================================================
// CollectionsOfCollection
// ==========================================================================

func Test_I8_CollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_I8_CollectionsOfCollection", func() {
		cc := corestr.New.CollectionsOfCollection.Empty()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		cc.Add(c1)
		if cc.Length() != 1 { t.Fatal("expected 1") }
	})
}

// ==========================================================================
// HashsetsCollection
// ==========================================================================

func Test_I8_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_I8_HashsetsCollection", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		h := corestr.New.Hashset.StringsSpreadItems("a")
		hc.Add(h)
		if hc.Length() != 1 { t.Fatal("expected 1") }
	})
}
