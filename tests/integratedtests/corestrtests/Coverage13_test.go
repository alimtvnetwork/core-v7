package corestrtests

import (
	"encoding/json"
	"errors"
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// Collection — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_Collection_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_BasicOps", func() {
		c := corestr.New.Collection.Cap(10)

		c.Add("a").Add("b").Add("c")

		if c.Length() != 3 || c.Count() != 3 {
			t.Fatal("expected 3")
		}

		if !c.HasAnyItem() || c.IsEmpty() || !c.HasItems() {
			t.Fatal("expected non-empty")
		}

		if c.LastIndex() != 2 {
			t.Fatal("expected 2")
		}

		if !c.HasIndex(2) || c.HasIndex(3) {
			t.Fatal("HasIndex failed")
		}

		if c.First() != "a" || c.Last() != "c" {
			t.Fatal("First/Last failed")
		}

		if c.FirstOrDefault() != "a" || c.LastOrDefault() != "c" {
			t.Fatal("FirstOrDefault/LastOrDefault failed")
		}

		if c.Capacity() < 10 {
			t.Fatal("expected cap >= 10")
		}
	})
}

func Test_Cov13_Collection_EmptyDefaults(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_EmptyDefaults", func() {
		c := corestr.Empty.Collection()

		if c.FirstOrDefault() != "" || c.LastOrDefault() != "" {
			t.Fatal("expected empty")
		}

		if !c.IsEmpty() || c.HasItems() || c.HasAnyItem() {
			t.Fatal("expected empty")
		}
	})
}

func Test_Cov13_Collection_AddVariants(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddVariants", func() {
		c := corestr.New.Collection.Cap(20)

		c.AddNonEmpty("")
		c.AddNonEmpty("x")
		c.AddNonEmptyWhitespace("  ")
		c.AddNonEmptyWhitespace("y")
		c.AddIf(false, "skip")
		c.AddIf(true, "keep")
		c.AddIfMany(false, "s1", "s2")
		c.AddIfMany(true, "m1", "m2")
		c.AddError(nil)
		c.AddError(errors.New("err1"))
		c.AddFunc(func() string { return "func1" })

		if c.Length() != 7 {
			t.Fatalf("expected 7 got %d", c.Length())
		}
	})
}

func Test_Cov13_Collection_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddFuncErr", func() {
		c := corestr.New.Collection.Cap(5)

		// No error
		c.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}

		// With error
		c.AddFuncErr(func() (string, error) { return "", errors.New("fail") }, func(e error) {})
		if c.Length() != 1 {
			t.Fatal("expected still 1")
		}
	})
}

func Test_Cov13_Collection_Adds(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Adds", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")
		c.AddStrings([]string{"d", "e"})

		if c.Length() != 5 {
			t.Fatal("expected 5")
		}
	})
}

func Test_Cov13_Collection_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddCollection", func() {
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"c", "d"})

		c1.AddCollection(c2)
		if c1.Length() != 4 {
			t.Fatal("expected 4")
		}

		// Empty collection
		c1.AddCollection(corestr.Empty.Collection())
		if c1.Length() != 4 {
			t.Fatal("expected still 4")
		}
	})
}

func Test_Cov13_Collection_AddCollections(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddCollections", func() {
		c := corestr.New.Collection.Cap(10)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})

		c.AddCollections(c1, c2, corestr.Empty.Collection())
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov13_Collection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_ConcatNew", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.ConcatNew(5, "c", "d")

		if result.Length() != 4 {
			t.Fatal("expected 4")
		}

		// Empty additionalStrings
		result2 := c.ConcatNew(0)
		if result2.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov13_Collection_AsError(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AsError", func() {
		c := corestr.New.Collection.Strings([]string{"err1", "err2"})
		err := c.AsError(",")

		if err == nil {
			t.Fatal("expected error")
		}

		empty := corestr.Empty.Collection()
		if empty.AsError(",") != nil {
			t.Fatal("expected nil")
		}

		_ = c.AsDefaultError()
	})
}

func Test_Cov13_Collection_ToError(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_ToError", func() {
		c := corestr.New.Collection.Strings([]string{"e1"})
		_ = c.ToError(",")
		_ = c.ToDefaultError()
	})
}

func Test_Cov13_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_RemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		if !c.RemoveAt(1) {
			t.Fatal("expected success")
		}

		if c.Length() != 2 {
			t.Fatal("expected 2")
		}

		// Out of range
		if c.RemoveAt(-1) || c.RemoveAt(10) {
			t.Fatal("expected failure")
		}
	})
}

func Test_Cov13_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_EachItemSplitBy", func() {
		c := corestr.New.Collection.Strings([]string{"a.b", "c.d"})
		result := c.EachItemSplitBy(".")

		if len(result) != 4 {
			t.Fatal("expected 4")
		}
	})
}

func Test_Cov13_Collection_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_IsEquals", func() {
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"a", "b"})
		c3 := corestr.New.Collection.Strings([]string{"a", "B"})

		if !c1.IsEquals(c2) {
			t.Fatal("expected equal")
		}

		if c1.IsEquals(c3) {
			t.Fatal("expected not equal case-sensitive")
		}

		if !c1.IsEqualsWithSensitive(false, c3) {
			t.Fatal("expected equal case-insensitive")
		}

		// Same ptr
		if !c1.IsEquals(c1) {
			t.Fatal("expected same ptr equal")
		}

		// Both empty
		e1 := corestr.Empty.Collection()
		e2 := corestr.Empty.Collection()
		if !e1.IsEquals(e2) {
			t.Fatal("expected empty equals")
		}

		// Different length
		c4 := corestr.New.Collection.Strings([]string{"a"})
		if c1.IsEquals(c4) {
			t.Fatal("expected not equal different length")
		}
	})
}

func Test_Cov13_Collection_Take_Skip(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Take_Skip", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})

		taken := c.Take(2)
		if taken.Length() != 2 || taken.First() != "a" {
			t.Fatal("Take failed")
		}

		// Take more than length
		taken2 := c.Take(10)
		if taken2.Length() != 4 {
			t.Fatal("expected full")
		}

		// Take 0
		taken3 := c.Take(0)
		if taken3.Length() != 0 {
			t.Fatal("expected empty")
		}

		// Skip
		skipped := c.Skip(2)
		if skipped.Length() != 2 || skipped.First() != "c" {
			t.Fatal("Skip failed")
		}

		// Skip 0
		if c.Skip(0) != c {
			t.Fatal("Skip 0 should return self")
		}
	})
}

func Test_Cov13_Collection_Reverse(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Reverse", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.Reverse()

		if c.First() != "c" || c.Last() != "a" {
			t.Fatal("Reverse failed")
		}

		// 2 items
		c2 := corestr.New.Collection.Strings([]string{"x", "y"})
		c2.Reverse()

		if c2.First() != "y" {
			t.Fatal("expected y first")
		}

		// 1 item
		c3 := corestr.New.Collection.Strings([]string{"z"})
		c3.Reverse()

		if c3.First() != "z" {
			t.Fatal("expected z")
		}
	})
}

func Test_Cov13_Collection_GetPagesSize(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_GetPagesSize", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})

		if c.GetPagesSize(2) != 3 {
			t.Fatal("expected 3")
		}

		if c.GetPagesSize(0) != 0 {
			t.Fatal("expected 0 for 0 page size")
		}
	})
}

func Test_Cov13_Collection_Has(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Has", func() {
		c := corestr.New.Collection.Strings([]string{"hello", "world"})

		if !c.Has("hello") || c.Has("missing") {
			t.Fatal("Has failed")
		}

		str := "hello"
		if !c.HasPtr(&str) {
			t.Fatal("HasPtr failed")
		}

		if !c.HasAll("hello", "world") {
			t.Fatal("HasAll failed")
		}

		if c.HasAll("hello", "missing") {
			t.Fatal("HasAll should fail")
		}

		if !c.HasUsingSensitivity("HELLO", false) {
			t.Fatal("expected case-insensitive match")
		}

		if c.HasUsingSensitivity("HELLO", true) {
			t.Fatal("expected case-sensitive mismatch")
		}
	})
}

func Test_Cov13_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Filter", func() {
		c := corestr.New.Collection.Strings([]string{"ab", "cd", "ef"})

		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, s == "ab" || s == "ef", false
		})

		if len(result) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov13_Collection_SortedListAsc_Dsc(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_SortedListAsc_Dsc", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})

		asc := c.SortedListAsc()
		if asc[0] != "a" {
			t.Fatal("expected a first")
		}

		dsc := c.SortedListDsc()
		if dsc[0] != "c" {
			t.Fatal("expected c first")
		}
	})
}

func Test_Cov13_Collection_UniqueList(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_UniqueList", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "a", "c"})
		uniques := c.UniqueList()

		if len(uniques) != 3 {
			t.Fatal("expected 3 unique")
		}
	})
}

func Test_Cov13_Collection_HashsetAsIs(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_HashsetAsIs", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		h := c.HashsetAsIs()

		if !h.Has("a") || !h.Has("b") {
			t.Fatal("expected both")
		}
	})
}

func Test_Cov13_Collection_NonEmptyList(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_NonEmptyList", func() {
		c := corestr.New.Collection.Strings([]string{"", "a", "", "b"})
		result := c.NonEmptyList()

		if len(result) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov13_Collection_IsContainsAll(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_IsContainsAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		if !c.IsContainsAll("a", "b") {
			t.Fatal("expected true")
		}

		if c.IsContainsAll("a", "z") {
			t.Fatal("expected false")
		}

		if c.IsContainsAll() {
			t.Fatal("expected false for nil")
		}
	})
}

func Test_Cov13_Collection_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_GetAllExcept", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.GetAllExcept([]string{"b"})

		if len(result) != 2 {
			t.Fatal("expected 2")
		}

		// nil items
		result2 := c.GetAllExcept(nil)
		if len(result2) != 3 {
			t.Fatal("expected 3 copy")
		}
	})
}

func Test_Cov13_Collection_Join(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Join", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		if c.Join(",") != "a,b,c" {
			t.Fatal("Join failed")
		}

		if c.JoinLine() == "" {
			t.Fatal("JoinLine failed")
		}

		empty := corestr.Empty.Collection()
		if empty.Join(",") != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_Cov13_Collection_String(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_String", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.String()

		if s == "" {
			t.Fatal("expected non-empty")
		}

		empty := corestr.Empty.Collection()
		_ = empty.String()
	})
}

func Test_Cov13_Collection_Csv(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Csv", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		_ = c.Csv()
		_ = c.CsvOptions(true)
		_ = c.CsvLines()
	})
}

func Test_Cov13_Collection_JSON(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_JSON", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		b, err := json.Marshal(c)
		if err != nil {
			t.Fatal("marshal failed")
		}

		c2 := corestr.Empty.Collection()
		err = json.Unmarshal(b, c2)
		if err != nil {
			t.Fatal("unmarshal failed")
		}

		if c2.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov13_Collection_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Serialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_, err := c.Serialize()
		if err != nil {
			t.Fatal("serialize failed")
		}
	})
}

func Test_Cov13_Collection_Resize_AddCapacity(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Resize_AddCapacity", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddCapacity(100)
		c.Resize(200)

		// Already big enough
		c.Resize(10)
	})
}

func Test_Cov13_Collection_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Clear_Dispose", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Clear()

		if c.Length() != 0 {
			t.Fatal("expected 0 after clear")
		}

		c2 := corestr.New.Collection.Strings([]string{"x"})
		c2.Dispose()
	})
}

func Test_Cov13_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AppendAnys", func() {
		c := corestr.New.Collection.Cap(10)
		c.AppendAnys(42, "hello", nil)

		if c.Length() != 2 {
			t.Fatal("expected 2 (nil skipped)")
		}
	})
}

func Test_Cov13_Collection_AppendNonEmptyAnys(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AppendNonEmptyAnys", func() {
		c := corestr.New.Collection.Cap(10)
		c.AppendNonEmptyAnys(42, nil)

		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov13_Collection_AddsNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddsNonEmpty", func() {
		c := corestr.New.Collection.Cap(10)
		c.AddsNonEmpty("a", "", "b")

		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov13_Collection_Joins(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Joins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// With extra items
		r := c.Joins(",", "c", "d")
		if r == "" {
			t.Fatal("expected non-empty")
		}

		// Without extra items
		r2 := c.Joins(",")
		if r2 != "a,b" {
			t.Fatal("expected a,b")
		}
	})
}

func Test_Cov13_Collection_New(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_New", func() {
		c := corestr.New.Collection.Cap(5)
		result := c.New("a", "b")

		if result.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov13_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_IndexAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		if c.IndexAt(1) != "b" {
			t.Fatal("expected b")
		}
	})
}

func Test_Cov13_Collection_SafeIndexAtUsingLength(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_SafeIndexAtUsingLength", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		if c.SafeIndexAtUsingLength("def", 2, 5) != "def" {
			t.Fatal("expected default")
		}

		if c.SafeIndexAtUsingLength("def", 2, 1) != "b" {
			t.Fatal("expected b")
		}
	})
}

func Test_Cov13_Collection_InsertAt(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_InsertAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "c"})
		c.InsertAt(1, "b")
	})
}

func Test_Cov13_Collection_ChainRemoveAt(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_ChainRemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.ChainRemoveAt(1)

		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov13_Collection_AppendCollectionPtr(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AppendCollectionPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b", "c"})
		c.AppendCollectionPtr(c2)

		if c.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_Cov13_Collection_AddFuncResult(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddFuncResult", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddFuncResult(func() string { return "hello" })

		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov13_Collection_AddStringsByFuncChecking(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddStringsByFuncChecking", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddStringsByFuncChecking([]string{"a", "bb", "c"}, func(s string) bool {
			return len(s) == 1
		})

		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov13_Collection_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_GetAllExceptCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		except := corestr.New.Collection.Strings([]string{"b"})
		result := c.GetAllExceptCollection(except)

		if len(result) != 2 {
			t.Fatal("expected 2")
		}

		// nil
		result2 := c.GetAllExceptCollection(nil)
		if len(result2) != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_Cov13_Collection_SummaryString(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_SummaryString", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.SummaryString(1)

		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_Cov13_Collection_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddHashmapsValues", func() {
		c := corestr.New.Collection.Cap(10)
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("k1", "v1")
		h.AddOrUpdate("k2", "v2")

		c.AddHashmapsValues(h)
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}

		c.AddHashmapsValues(nil)
	})
}

func Test_Cov13_Collection_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddHashmapsKeys", func() {
		c := corestr.New.Collection.Cap(10)
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("k1", "v1")

		c.AddHashmapsKeys(h)
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov13_Collection_AddHashmapsKeysValues(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddHashmapsKeysValues", func() {
		c := corestr.New.Collection.Cap(10)
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("k", "v")

		c.AddHashmapsKeysValues(h)
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov13_Collection_GetHashsetPlusHasAll(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_GetHashsetPlusHasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		h, ok := c.GetHashsetPlusHasAll([]string{"a", "b"})

		if !ok || h == nil {
			t.Fatal("expected true")
		}

		_, ok2 := c.GetHashsetPlusHasAll(nil)
		if ok2 {
			t.Fatal("expected false for nil")
		}
	})
}

func Test_Cov13_Collection_GetPagedCollection(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_GetPagedCollection", func() {
		items := make([]string, 25)
		for i := range items {
			items[i] = "item"
		}

		c := corestr.New.Collection.Strings(items)
		paged := c.GetPagedCollection(10)

		if paged.Length() != 3 {
			t.Fatalf("expected 3 pages got %d", paged.Length())
		}
	})
}

func Test_Cov13_Collection_GetSinglePageCollection(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_GetSinglePageCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})

		page := c.GetSinglePageCollection(2, 1)
		if page.Length() != 2 {
			t.Fatal("expected 2")
		}

		page2 := c.GetSinglePageCollection(2, 3)
		if page2.Length() != 1 {
			t.Fatal("expected 1")
		}

		// When length < eachPageSize
		small := corestr.New.Collection.Strings([]string{"a"})
		if small.GetSinglePageCollection(10, 1) != small {
			t.Fatal("expected self")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Creators
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_CollectionCreator_All(t *testing.T) {
	safeTest(t, "Test_Cov13_CollectionCreator_All", func() {
		_ = corestr.New.Collection.Empty()
		_ = corestr.New.Collection.Cap(5)
		_ = corestr.New.Collection.Create([]string{"a"})
		_ = corestr.New.Collection.Strings([]string{"a"})
		_ = corestr.New.Collection.StringsOptions(true, []string{"a"})
		_ = corestr.New.Collection.StringsOptions(false, []string{"a"})
		_ = corestr.New.Collection.StringsOptions(false, []string{})
		_ = corestr.New.Collection.CloneStrings([]string{"a"})
		_ = corestr.New.Collection.LineUsingSep(",", "a,b")
		_ = corestr.New.Collection.LineDefault("a\nb")
		_ = corestr.New.Collection.StringsPlusCap(5, []string{"a"})
		_ = corestr.New.Collection.StringsPlusCap(0, []string{"a"})
		_ = corestr.New.Collection.CapStrings(5, []string{"a"})
		_ = corestr.New.Collection.CapStrings(0, []string{"a"})
		_ = corestr.New.Collection.LenCap(0, 10)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleSlice — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_SimpleSlice_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_BasicOps", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")

		if s.Length() != 3 || s.Count() != 3 {
			t.Fatal("expected 3")
		}

		if s.IsEmpty() || !s.HasAnyItem() {
			t.Fatal("expected non-empty")
		}

		if s.LastIndex() != 2 {
			t.Fatal("expected 2")
		}

		if !s.HasIndex(2) || s.HasIndex(3) {
			t.Fatal("HasIndex failed")
		}

		if s.First() != "a" || s.Last() != "c" {
			t.Fatal("First/Last failed")
		}

		if s.FirstOrDefault() != "a" || s.LastOrDefault() != "c" {
			t.Fatal("FirstOrDefault/LastOrDefault failed")
		}
	})
}

func Test_Cov13_SimpleSlice_EmptyDefaults(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_EmptyDefaults", func() {
		s := corestr.New.SimpleSlice.Empty()

		if s.FirstOrDefault() != "" || s.LastOrDefault() != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_Cov13_SimpleSlice_AddVariants(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_AddVariants", func() {
		s := corestr.New.SimpleSlice.Cap(10)
		s.Add("a")
		s.AddIf(false, "skip")
		s.AddIf(true, "keep")
		s.Adds("x", "y")
		s.AddsIf(false, "skip1")
		s.AddsIf(true, "z")
		s.AddError(errors.New("err"))
		s.AddError(nil)
		s.AddSplit("a.b.c", ".")

		if s.Length() != 9 {
			t.Fatalf("expected 9 got %d", s.Length())
		}
	})
}

func Test_Cov13_SimpleSlice_AddStruct(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_AddStruct", func() {
		s := corestr.New.SimpleSlice.Cap(5)
		s.AddStruct(true, struct{ Name string }{"hello"})
		s.AddStruct(true, nil)

		if s.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov13_SimpleSlice_AddPointer(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_AddPointer", func() {
		s := corestr.New.SimpleSlice.Cap(5)
		val := "test"
		s.AddPointer(true, &val)
		s.AddPointer(true, nil)

		if s.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov13_SimpleSlice_Append_AppendFmt(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_Append_AppendFmt", func() {
		s := corestr.New.SimpleSlice.Cap(5)
		s.Append("a", "b")
		s.AppendFmt("hello %s", "world")
		s.AppendFmt("", ) // empty format + no values → skip
		s.AppendFmtIf(true, "yes %d", 1)
		s.AppendFmtIf(false, "no")

		if s.Length() != 4 {
			t.Fatalf("expected 4 got %d", s.Length())
		}
	})
}

func Test_Cov13_SimpleSlice_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_AddAsTitleValue", func() {
		s := corestr.New.SimpleSlice.Cap(5)
		s.AddAsTitleValue("Name", "John")
		s.AddAsTitleValueIf(true, "Age", 30)
		s.AddAsTitleValueIf(false, "Skip", nil)
		s.AddAsCurlyTitleWrap("Key", "Val")
		s.AddAsCurlyTitleWrapIf(true, "K", "V")
		s.AddAsCurlyTitleWrapIf(false, "S", "S")

		if s.Length() != 4 {
			t.Fatalf("expected 4 got %d", s.Length())
		}
	})
}

func Test_Cov13_SimpleSlice_InsertAt(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_InsertAt", func() {
		s := corestr.New.SimpleSlice.Lines("a", "c")
		s.InsertAt(1, "b")

		if s.Length() != 3 {
			t.Fatal("expected 3")
		}

		// Out of range
		s.InsertAt(-1, "x")
		s.InsertAt(100, "x")
	})
}

func Test_Cov13_SimpleSlice_Skip_Take(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_Skip_Take", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "c", "d")

		skipped := s.Skip(2)
		if len(skipped) != 2 {
			t.Fatal("expected 2")
		}

		taken := s.Take(2)
		if len(taken) != 2 {
			t.Fatal("expected 2")
		}

		// Skip more than length
		skippedAll := s.Skip(10)
		if len(skippedAll) != 0 {
			t.Fatal("expected 0")
		}

		// Take more
		takenAll := s.Take(10)
		if len(takenAll) != 4 {
			t.Fatal("expected 4")
		}

		_ = s.Limit(2)
		_ = s.SkipDynamic(2)
		_ = s.TakeDynamic(2)
		_ = s.LimitDynamic(2)
	})
}

func Test_Cov13_SimpleSlice_AsError(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_AsError", func() {
		s := corestr.New.SimpleSlice.Lines("e1", "e2")
		err := s.AsError(",")

		if err == nil {
			t.Fatal("expected error")
		}

		_ = s.AsDefaultError()

		empty := corestr.New.SimpleSlice.Empty()
		if empty.AsError(",") != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_Cov13_SimpleSlice_Join(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_Join", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")

		if s.Join(",") != "a,b" {
			t.Fatal("expected a,b")
		}

		_ = s.JoinLine()
		_ = s.JoinLineEofLine()
		_ = s.JoinSpace()
		_ = s.JoinComma()
		_ = s.JoinCsv()
		_ = s.JoinCsvLine()
		_ = s.JoinWith(",")
		_ = s.JoinCsvString(",")
	})
}

func Test_Cov13_SimpleSlice_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_IsContains", func() {
		s := corestr.New.SimpleSlice.Lines("hello", "world")

		if !s.IsContains("hello") || s.IsContains("missing") {
			t.Fatal("IsContains failed")
		}

		if s.IndexOf("world") != 1 {
			t.Fatal("IndexOf failed")
		}

		if s.IndexOf("missing") != -1 {
			t.Fatal("expected -1")
		}
	})
}

func Test_Cov13_SimpleSlice_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_IsContainsFunc", func() {
		s := corestr.New.SimpleSlice.Lines("hello", "world")

		found := s.IsContainsFunc("hello", func(item, searching string) bool {
			return item == searching
		})

		if !found {
			t.Fatal("expected found")
		}
	})
}

func Test_Cov13_SimpleSlice_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_IndexOfFunc", func() {
		s := corestr.New.SimpleSlice.Lines("hello", "world")

		idx := s.IndexOfFunc("world", func(item, searching string) bool {
			return item == searching
		})

		if idx != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov13_SimpleSlice_CountFunc(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_CountFunc", func() {
		s := corestr.New.SimpleSlice.Lines("a", "bb", "c")

		count := s.CountFunc(func(i int, item string) bool {
			return len(item) == 1
		})

		if count != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov13_SimpleSlice_WrapQuotes(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_WrapQuotes", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")

		_ = s.WrapDoubleQuote()
		_ = s.WrapSingleQuote()
		_ = s.WrapTildaQuote()
		_ = s.WrapDoubleQuoteIfMissing()
		_ = s.WrapSingleQuoteIfMissing()
	})
}

func Test_Cov13_SimpleSlice_Transpile(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_Transpile", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.Transpile(func(s string) string { return s + "!" })

		if (*result)[0] != "a!" {
			t.Fatal("expected a!")
		}

		// TranspileJoin
		joined := s.TranspileJoin(func(s string) string { return s }, ",")
		if joined != "a,b" {
			t.Fatal("expected a,b")
		}
	})
}

func Test_Cov13_SimpleSlice_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_EachItemSplitBy", func() {
		s := corestr.New.SimpleSlice.Lines("a.b", "c.d")
		result := s.EachItemSplitBy(".")

		if result.Length() != 4 {
			t.Fatal("expected 4")
		}
	})
}

func Test_Cov13_SimpleSlice_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_IsEqual", func() {
		s1 := corestr.New.SimpleSlice.Lines("a", "b")
		s2 := corestr.New.SimpleSlice.Lines("a", "b")
		s3 := corestr.New.SimpleSlice.Lines("a", "c")

		if !s1.IsEqual(s2) {
			t.Fatal("expected equal")
		}

		if s1.IsEqual(s3) {
			t.Fatal("expected not equal")
		}

		if !s1.IsEqualLines([]string{"a", "b"}) {
			t.Fatal("expected equal")
		}
	})
}

func Test_Cov13_SimpleSlice_IsUnorderedEqual(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_IsUnorderedEqual", func() {
		s1 := corestr.New.SimpleSlice.Lines("b", "a")
		s2 := corestr.New.SimpleSlice.Lines("a", "b")

		if !s1.IsUnorderedEqual(true, s2) {
			t.Fatal("expected equal")
		}
	})
}

func Test_Cov13_SimpleSlice_IsDistinctEqual(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_IsDistinctEqual", func() {
		s1 := corestr.New.SimpleSlice.Lines("a", "b", "a")
		s2 := corestr.New.SimpleSlice.Lines("b", "a")

		if !s1.IsDistinctEqual(s2) {
			t.Fatal("expected equal")
		}
	})
}

func Test_Cov13_SimpleSlice_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_ConcatNew", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.ConcatNew("c", "d")

		if result.Length() != 4 {
			t.Fatal("expected 4")
		}

		_ = s.ConcatNewStrings("c")
	})
}

func Test_Cov13_SimpleSlice_ConcatNewSimpleSlices(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_ConcatNewSimpleSlices", func() {
		s1 := corestr.New.SimpleSlice.Lines("a")
		s2 := corestr.New.SimpleSlice.Lines("b")
		result := s1.ConcatNewSimpleSlices(s2)

		if result.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov13_SimpleSlice_PrependAppend(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_PrependAppend", func() {
		s := corestr.New.SimpleSlice.Lines("b")
		s.PrependAppend([]string{"a"}, []string{"c"})

		if s.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_Cov13_SimpleSlice_PrependJoin(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_PrependJoin", func() {
		s := corestr.New.SimpleSlice.Lines("b", "c")
		result := s.PrependJoin(",", "a")

		if result != "a,b,c" {
			t.Fatalf("expected a,b,c got %s", result)
		}
	})
}

func Test_Cov13_SimpleSlice_AppendJoin(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_AppendJoin", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.AppendJoin(",", "c")

		if result != "a,b,c" {
			t.Fatalf("expected a,b,c got %s", result)
		}
	})
}

func Test_Cov13_SimpleSlice_Sort_Reverse(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_Sort_Reverse", func() {
		s := corestr.New.SimpleSlice.Lines("c", "a", "b")
		s.Sort()

		if s.First() != "a" {
			t.Fatal("expected a first")
		}

		s.Reverse()
		if s.First() != "c" {
			t.Fatal("expected c first")
		}
	})
}
func Test_Cov13_SimpleSlice_RemoveIndexes(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_RemoveIndexes", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result, err := s.RemoveIndexes(1)

		if err != nil || result.Length() != 2 {
			t.Fatal("expected 2")
		}

		// Empty slice
		empty := corestr.New.SimpleSlice.Empty()
		_, err = empty.RemoveIndexes(0)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_Cov13_SimpleSlice_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_DistinctDiff", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.DistinctDiffRaw("b", "c")

		if len(result) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov13_SimpleSlice_AddedRemovedLinesDiff(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_AddedRemovedLinesDiff", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		added, removed := s.AddedRemovedLinesDiff("b", "c")

		if len(added) != 1 || len(removed) != 1 {
			t.Fatal("expected 1 added, 1 removed")
		}
	})
}

func Test_Cov13_SimpleSlice_JSON(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_JSON", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")

		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal("marshal failed")
		}

		s2 := corestr.New.SimpleSlice.Empty()
		err = json.Unmarshal(b, s2)
		if err != nil {
			t.Fatal("unmarshal failed")
		}
	})
}

func Test_Cov13_SimpleSlice_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_Clear_Dispose", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		s.Clear()

		if s.Length() != 0 {
			t.Fatal("expected 0")
		}

		s2 := corestr.New.SimpleSlice.Lines("x")
		s2.Dispose()
	})
}

func Test_Cov13_SimpleSlice_Collection(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_Collection", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		c := s.Collection(false)

		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov13_SimpleSlice_Hashset(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_Hashset", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		h := s.Hashset()

		if !h.Has("a") {
			t.Fatal("expected a")
		}
	})
}

func Test_Cov13_SimpleSlice_String(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_String", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		_ = s.String()

		empty := corestr.New.SimpleSlice.Empty()
		_ = empty.String()
	})
}

func Test_Cov13_SimpleSlice_IsEqualByFunc(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_IsEqualByFunc", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")

		ok := s.IsEqualByFunc(func(i int, l, r string) bool {
			return l == r
		}, "a", "b")

		if !ok {
			t.Fatal("expected true")
		}
	})
}

func Test_Cov13_SimpleSlice_SafeStrings(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_SafeStrings", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		_ = s.SafeStrings()

		empty := corestr.New.SimpleSlice.Empty()
		_ = empty.SafeStrings()
	})
}

func Test_Cov13_SimpleSlice_Creators(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_Creators", func() {
		_ = corestr.New.SimpleSlice.Cap(5)
		_ = corestr.New.SimpleSlice.Default()
		_ = corestr.New.SimpleSlice.Empty()
		_ = corestr.New.SimpleSlice.Lines("a")
		_ = corestr.New.SimpleSlice.SpreadStrings("a")
		_ = corestr.New.SimpleSlice.Create([]string{"a"})
		_ = corestr.New.SimpleSlice.Strings([]string{"a"})
		_ = corestr.New.SimpleSlice.StringsClone([]string{"a"})
		_ = corestr.New.SimpleSlice.StringsClone(nil)
		_ = corestr.New.SimpleSlice.Direct(true, []string{"a"})
		_ = corestr.New.SimpleSlice.Direct(false, []string{"a"})
		_ = corestr.New.SimpleSlice.Direct(true, nil)
		_ = corestr.New.SimpleSlice.UsingLines(true, "a")
		_ = corestr.New.SimpleSlice.UsingLines(false, "a")
		_ = corestr.New.SimpleSlice.Split("a.b", ".")
		_ = corestr.New.SimpleSlice.SplitLines("a\nb")
		_ = corestr.New.SimpleSlice.UsingSeparatorLine(",", "a,b")
		_ = corestr.New.SimpleSlice.UsingLine("a\nb")
		_ = corestr.New.SimpleSlice.StringsOptions(true, []string{"a"})
		_ = corestr.New.SimpleSlice.StringsOptions(false, []string{"a"})
		_ = corestr.New.SimpleSlice.StringsOptions(true, []string{})

		h := corestr.New.Hashset.Strings([]string{"a"})
		_ = corestr.New.SimpleSlice.Hashset(h)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_ValidValue_Constructors(t *testing.T) {
	safeTest(t, "Test_Cov13_ValidValue_Constructors", func() {
		_ = corestr.NewValidValue("hello")
		_ = corestr.NewValidValueEmpty()
		_ = corestr.InvalidValidValue("msg")
		_ = corestr.InvalidValidValueNoMessage()
		_ = corestr.NewValidValueUsingAny(true, true, "val")
		_ = corestr.NewValidValueUsingAnyAutoValid(false, "val")
	})
}
func Test_Cov13_ValidValues_AllMethods(t *testing.T) {
	safeTest(t, "Test_Cov13_ValidValues_AllMethods", func() {
		vv := corestr.NewValidValues(5)
		vv.Add("a")
		vv.AddFull(true, "b", "msg")

		if vv.Length() != 2 || vv.Count() != 2 || vv.IsEmpty() || !vv.HasAnyItem() {
			t.Fatal("expected 2")
		}

		if vv.LastIndex() != 1 || !vv.HasIndex(1) || vv.HasIndex(5) {
			t.Fatal("index check failed")
		}

		_ = vv.SafeValueAt(0)
		_ = vv.SafeValueAt(100)
		_ = vv.SafeValidValueAt(0)
		_ = vv.SafeValidValueAt(100)
		_ = vv.SafeValuesAtIndexes(0, 1)
		_ = vv.SafeValidValuesAtIndexes(0, 1)
		_ = vv.Strings()
		_ = vv.FullStrings()
		_ = vv.String()

		found := vv.Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return v, true, false
		})
		if len(found) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov13_ValidValues_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov13_ValidValues_ConcatNew", func() {
		vv := corestr.NewValidValues(2)
		vv.Add("a")

		vv2 := corestr.NewValidValues(2)
		vv2.Add("b")

		result := vv.ConcatNew(true, vv2)
		if result.Length() != 2 {
			t.Fatal("expected 2")
		}

		// Empty concat
		result2 := vv.ConcatNew(true)
		if result2.Length() != 1 {
			t.Fatal("expected 1")
		}

		result3 := vv.ConcatNew(false)
		if result3 != vv {
			t.Fatal("expected same ptr")
		}
	})
}

func Test_Cov13_ValidValues_Constructors(t *testing.T) {
	safeTest(t, "Test_Cov13_ValidValues_Constructors", func() {
		_ = corestr.EmptyValidValues()
		_ = corestr.NewValidValuesUsingValues(corestr.ValidValue{Value: "a", IsValid: true})
		_ = corestr.NewValidValuesUsingValues()
	})
}

func Test_Cov13_ValidValues_Hashmap_Map(t *testing.T) {
	safeTest(t, "Test_Cov13_ValidValues_Hashmap_Map", func() {
		vv := corestr.NewValidValues(2)
		vv.Add("a")
		_ = vv.Hashmap()
		_ = vv.Map()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValueStatus — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_ValueStatus_All(t *testing.T) {
	safeTest(t, "Test_Cov13_ValueStatus_All", func() {
		vs := corestr.InvalidValueStatus("msg")
		if vs.ValueValid.IsValid {
			t.Fatal("expected invalid")
		}

		vs2 := corestr.InvalidValueStatusNoMessage()
		_ = vs2.Clone()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValuePair — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_KeyValuePair_AllMethods(t *testing.T) {
	safeTest(t, "Test_Cov13_KeyValuePair_AllMethods", func() {
		kv := corestr.KeyValuePair{Key: "name", Value: "42"}

		if kv.KeyName() != "name" || kv.VariableName() != "name" {
			t.Fatal("key failed")
		}

		if kv.ValueString() != "42" {
			t.Fatal("expected 42")
		}

		if !kv.IsVariableNameEqual("name") || !kv.IsValueEqual("42") {
			t.Fatal("equality failed")
		}

		if kv.IsKeyEmpty() || kv.IsValueEmpty() || kv.IsKeyValueEmpty() || kv.IsKeyValueAnyEmpty() {
			t.Fatal("expected non-empty")
		}

		if !kv.HasKey() || !kv.HasValue() {
			t.Fatal("expected has")
		}

		_ = kv.TrimKey()
		_ = kv.TrimValue()
		_ = kv.String()
		_ = kv.Compile()
		_ = kv.FormatString("%s=%s")

		if kv.ValueInt(0) != 42 || kv.ValueDefInt() != 42 {
			t.Fatal("expected 42")
		}

		_ = kv.ValueByte(0)
		_ = kv.ValueDefByte()

		kvBool := corestr.KeyValuePair{Key: "k", Value: "true"}
		if !kvBool.ValueBool() {
			t.Fatal("expected true")
		}

		kvFloat := corestr.KeyValuePair{Key: "k", Value: "3.14"}
		_ = kvFloat.ValueFloat64(0)
		_ = kvFloat.ValueDefFloat64()

		_ = kv.ValueValid()
		_ = kv.ValueValidOptions(true, "msg")

		if !kv.Is("name", "42") || !kv.IsKey("name") || !kv.IsVal("42") {
			t.Fatal("Is failed")
		}

		kv.Clear()
		kv.Dispose()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyAnyValuePair — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_KeyAnyValuePair_AllMethods(t *testing.T) {
	safeTest(t, "Test_Cov13_KeyAnyValuePair_AllMethods", func() {
		kav := corestr.KeyAnyValuePair{Key: "test", Value: 42}

		if kav.KeyName() != "test" || kav.VariableName() != "test" {
			t.Fatal("key failed")
		}

		if kav.ValueAny() != 42 {
			t.Fatal("expected 42")
		}

		if !kav.IsVariableNameEqual("test") {
			t.Fatal("expected equal")
		}

		_ = kav.ValueString()
		_ = kav.String()
		_ = kav.Compile()

		if kav.IsValueNull() || !kav.HasNonNull() || !kav.HasValue() {
			t.Fatal("expected non-null")
		}

		if kav.IsValueEmptyString() || kav.IsValueWhitespace() {
			t.Fatal("expected non-empty string")
		}

		kav.Clear()
		kav.Dispose()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// TextWithLineNumber — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_TextWithLineNumber(t *testing.T) {
	safeTest(t, "Test_Cov13_TextWithLineNumber", func() {
		tl := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}

		if !tl.HasLineNumber() || tl.IsInvalidLineNumber() {
			t.Fatal("expected valid line number")
		}

		if tl.Length() != 5 || tl.IsEmpty() || tl.IsEmptyText() {
			t.Fatal("expected non-empty")
		}

		// Empty
		empty := &corestr.TextWithLineNumber{}
		if !empty.IsEmptyTextLineBoth() {
			t.Fatal("expected empty")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CloneSlice / CloneSliceIf — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_CloneSlice(t *testing.T) {
	safeTest(t, "Test_Cov13_CloneSlice", func() {
		result := corestr.CloneSlice([]string{"a", "b"})

		if len(result) != 2 {
			t.Fatal("expected 2")
		}

		empty := corestr.CloneSlice([]string{})
		if len(empty) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_Cov13_CloneSliceIf(t *testing.T) {
	safeTest(t, "Test_Cov13_CloneSliceIf", func() {
		result := corestr.CloneSliceIf(true, "a", "b")
		if len(result) != 2 {
			t.Fatal("expected 2")
		}

		// No clone
		result2 := corestr.CloneSliceIf(false, "a", "b")
		if len(result2) != 2 {
			t.Fatal("expected 2")
		}

		// Empty
		result3 := corestr.CloneSliceIf(true)
		if len(result3) != 0 {
			t.Fatal("expected 0")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// utils — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_Utils_WrapMethods(t *testing.T) {
	safeTest(t, "Test_Cov13_Utils_WrapMethods", func() {
		u := corestr.StringUtils

		if u.WrapDouble("a") != `"a"` {
			t.Fatal("WrapDouble failed")
		}

		if u.WrapSingle("a") != "'a'" {
			t.Fatal("WrapSingle failed")
		}

		if u.WrapTilda("a") != "`a`" {
			t.Fatal("WrapTilda failed")
		}

		if u.WrapDoubleIfMissing(`"a"`) != `"a"` {
			t.Fatal("already wrapped")
		}

		if u.WrapDoubleIfMissing("a") != `"a"` {
			t.Fatal("wrap missing")
		}

		if u.WrapDoubleIfMissing("") != `""` {
			t.Fatal("empty wrap")
		}

		if u.WrapSingleIfMissing("'a'") != "'a'" {
			t.Fatal("already wrapped")
		}

		if u.WrapSingleIfMissing("a") != "'a'" {
			t.Fatal("wrap missing")
		}

		if u.WrapSingleIfMissing("") != "''" {
			t.Fatal("empty wrap")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyToString / reflectInterfaceVal — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_AnyToString(t *testing.T) {
	safeTest(t, "Test_Cov13_AnyToString", func() {
		r := corestr.AnyToString(true, 42)
		if r == "" {
			t.Fatal("expected non-empty")
		}

		r2 := corestr.AnyToString(false, "hello")
		if r2 == "" {
			t.Fatal("expected non-empty")
		}

		// Empty
		r3 := corestr.AnyToString(true, "")
		if r3 != "" {
			t.Fatal("expected empty")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualsLengthOfSimpleSlices — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_AllIndividualsLengthOfSimpleSlices(t *testing.T) {
	safeTest(t, "Test_Cov13_AllIndividualsLengthOfSimpleSlices", func() {
		s1 := corestr.New.SimpleSlice.Lines("a", "b")
		s2 := corestr.New.SimpleSlice.Lines("c")

		length := corestr.AllIndividualsLengthOfSimpleSlices(s1, s2)
		if length != 3 {
			t.Fatal("expected 3")
		}

		if corestr.AllIndividualsLengthOfSimpleSlices() != 0 {
			t.Fatal("expected 0")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_AllIndividualStringsOfStringsLength(t *testing.T) {
	safeTest(t, "Test_Cov13_AllIndividualStringsOfStringsLength", func() {
		items := [][]string{{"a", "b"}, {"c"}}
		length := corestr.AllIndividualStringsOfStringsLength(&items)

		if length != 3 {
			t.Fatal("expected 3")
		}

		if corestr.AllIndividualStringsOfStringsLength(nil) != 0 {
			t.Fatal("expected 0")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_LeftRight_AllMethods(t *testing.T) {
	safeTest(t, "Test_Cov13_LeftRight_AllMethods", func() {
		lr := corestr.NewLeftRight("left", "right")

		if lr.Left != "left" || lr.Right != "right" {
			t.Fatal("expected left/right")
		}

		_ = lr.LeftBytes()
		_ = lr.RightBytes()
		_ = lr.LeftTrim()
		_ = lr.RightTrim()

		if lr.IsLeftEmpty() || lr.IsRightEmpty() {
			t.Fatal("expected non-empty")
		}

		if lr.IsLeftWhitespace() || lr.IsRightWhitespace() {
			t.Fatal("expected non-whitespace")
		}

		if !lr.HasValidNonEmptyLeft() || !lr.HasValidNonEmptyRight() {
			t.Fatal("expected valid non-empty")
		}

		if !lr.HasValidNonWhitespaceLeft() || !lr.HasValidNonWhitespaceRight() {
			t.Fatal("expected valid non-whitespace")
		}

		if !lr.HasSafeNonEmpty() {
			t.Fatal("expected safe non-empty")
		}

		if !lr.IsLeft("left") || !lr.IsRight("right") || !lr.Is("left", "right") {
			t.Fatal("Is failed")
		}

		lr2 := corestr.NewLeftRight("left", "right")
		if !lr.IsEqual(lr2) {
			t.Fatal("expected equal")
		}

		_ = lr.Clone()
		_ = lr.NonPtr()
		_ = lr.Ptr()

		re := regexp.MustCompile("left")
		if !lr.IsLeftRegexMatch(re) {
			t.Fatal("expected match")
		}
		if lr.IsLeftRegexMatch(nil) {
			t.Fatal("expected false for nil")
		}

		lr.Clear()
		lr.Dispose()
	})
}

func Test_Cov13_LeftRight_Constructors(t *testing.T) {
	safeTest(t, "Test_Cov13_LeftRight_Constructors", func() {
		_ = corestr.InvalidLeftRight("msg")
		_ = corestr.InvalidLeftRightNoMessage()
		_ = corestr.LeftRightUsingSlice([]string{"a", "b"})
		_ = corestr.LeftRightUsingSlice([]string{"a"})
		_ = corestr.LeftRightUsingSlice([]string{})
		_ = corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
		_ = corestr.LeftRightUsingSlicePtr([]string{})
		_ = corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
		_ = corestr.LeftRightTrimmedUsingSlice([]string{"a"})
		_ = corestr.LeftRightTrimmedUsingSlice(nil)
		_ = corestr.Empty.LeftRight()
	})
}

func Test_Cov13_LeftRight_FromSplit(t *testing.T) {
	safeTest(t, "Test_Cov13_LeftRight_FromSplit", func() {
		lr := corestr.LeftRightFromSplit("a=b", "=")
		if lr.Left != "a" || lr.Right != "b" {
			t.Fatal("expected a=b")
		}

		_ = corestr.LeftRightFromSplitTrimmed(" a = b ", "=")
		_ = corestr.LeftRightFromSplitFull("a:b:c", ":")
		_ = corestr.LeftRightFromSplitFullTrimmed(" a : b:c ", ":")
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRight — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_LeftMiddleRight_AllMethods(t *testing.T) {
	safeTest(t, "Test_Cov13_LeftMiddleRight_AllMethods", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		_ = lmr.LeftBytes()
		_ = lmr.RightBytes()
		_ = lmr.MiddleBytes()
		_ = lmr.LeftTrim()
		_ = lmr.RightTrim()
		_ = lmr.MiddleTrim()

		if lmr.IsLeftEmpty() || lmr.IsRightEmpty() || lmr.IsMiddleEmpty() {
			t.Fatal("expected non-empty")
		}

		if lmr.IsLeftWhitespace() || lmr.IsRightWhitespace() || lmr.IsMiddleWhitespace() {
			t.Fatal("expected non-whitespace")
		}

		if !lmr.HasValidNonEmptyLeft() || !lmr.HasValidNonEmptyRight() || !lmr.HasValidNonEmptyMiddle() {
			t.Fatal("expected valid")
		}

		if !lmr.HasValidNonWhitespaceLeft() || !lmr.HasValidNonWhitespaceRight() || !lmr.HasValidNonWhitespaceMiddle() {
			t.Fatal("expected valid")
		}

		if !lmr.HasSafeNonEmpty() {
			t.Fatal("expected safe")
		}

		if !lmr.IsAll("a", "b", "c") || !lmr.Is("a", "c") {
			t.Fatal("Is failed")
		}

		_ = lmr.Clone()
		_ = lmr.ToLeftRight()

		lmr.Clear()
		lmr.Dispose()
	})
}

func Test_Cov13_LeftMiddleRight_Constructors(t *testing.T) {
	safeTest(t, "Test_Cov13_LeftMiddleRight_Constructors", func() {
		_ = corestr.InvalidLeftMiddleRight("msg")
		_ = corestr.InvalidLeftMiddleRightNoMessage()
	})
}

func Test_Cov13_LeftMiddleRight_FromSplit(t *testing.T) {
	safeTest(t, "Test_Cov13_LeftMiddleRight_FromSplit", func() {
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
		if lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c" {
			t.Fatal("split failed")
		}

		_ = corestr.LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")
		_ = corestr.LeftMiddleRightFromSplitN("a:b:c:d", ":")
		_ = corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c:d ", ":")
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Empty creators — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_EmptyCreators(t *testing.T) {
	safeTest(t, "Test_Cov13_EmptyCreators", func() {
		_ = corestr.Empty.Collection()
		_ = corestr.Empty.LinkedList()
		_ = corestr.Empty.SimpleSlice()
		_ = corestr.Empty.KeyAnyValuePair()
		_ = corestr.Empty.KeyValuePair()
		_ = corestr.Empty.KeyValueCollection()
		_ = corestr.Empty.LinkedCollections()
		_ = corestr.Empty.LeftRight()
		_ = corestr.Empty.SimpleStringOnce()
		_ = corestr.Empty.SimpleStringOncePtr()
		_ = corestr.Empty.Hashset()
		_ = corestr.Empty.HashsetsCollection()
		_ = corestr.Empty.Hashmap()
		_ = corestr.Empty.CharCollectionMap()
		_ = corestr.Empty.KeyValuesCollection()
		_ = corestr.Empty.CollectionsOfCollection()
		_ = corestr.Empty.CharHashsetMap()
	})
}
