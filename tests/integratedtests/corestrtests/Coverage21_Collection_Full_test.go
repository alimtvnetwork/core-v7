package corestrtests

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// Collection.go — comprehensive coverage for remaining uncovered methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov21_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_IndexAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.IndexAt(0) != "a" {
			t.Fatal("wrong index")
		}
	})
}

func Test_Cov21_Collection_SafeIndexAtUsingLength(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_SafeIndexAtUsingLength", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.SafeIndexAtUsingLength("def", 1, 0) != "a" {
			t.Fatal("wrong")
		}
		if c.SafeIndexAtUsingLength("def", 1, 5) != "def" {
			t.Fatal("expected default")
		}
	})
}

func Test_Cov21_Collection_First_Last(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_First_Last", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		if c.First() != "a" || c.Last() != "c" {
			t.Fatal("wrong")
		}
	})
}

func Test_Cov21_Collection_FirstOrDefault_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_FirstOrDefault_LastOrDefault", func() {
		c := corestr.New.Collection.Empty()
		if c.FirstOrDefault() != "" || c.LastOrDefault() != "" {
			t.Fatal("expected empty")
		}
		c.Add("x")
		if c.FirstOrDefault() != "x" || c.LastOrDefault() != "x" {
			t.Fatal("wrong")
		}
	})
}

func Test_Cov21_Collection_Take(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Take", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		tk := c.Take(2)
		if tk.Length() != 2 {
			t.Fatal("wrong")
		}
		full := c.Take(10)
		if full.Length() != 3 {
			t.Fatal("wrong")
		}
		zero := c.Take(0)
		if zero.Length() != 0 {
			t.Fatal("wrong")
		}
	})
}

func Test_Cov21_Collection_Skip(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Skip", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		sk := c.Skip(1)
		if sk.Length() != 2 {
			t.Fatal("wrong")
		}
		same := c.Skip(0)
		if same.Length() != 3 {
			t.Fatal("wrong")
		}
	})
}

func Test_Cov21_Collection_Reverse(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Reverse", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.Reverse()
		if c.First() != "c" || c.Last() != "a" {
			t.Fatal("wrong")
		}
		// 2 items
		c2 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2.Reverse()
		if c2.First() != "b" {
			t.Fatal("wrong")
		}
		// 1 item
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c1.Reverse()
	})
}

func Test_Cov21_Collection_Paging(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Paging", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		if c.GetPagesSize(2) != 3 {
			t.Fatal("wrong pages")
		}
		if c.GetPagesSize(0) != 0 {
			t.Fatal("expected 0")
		}
		pages := c.GetPagedCollection(2)
		if pages.Length() != 3 {
			t.Fatalf("expected 3, got %d", pages.Length())
		}
		page := c.GetSinglePageCollection(2, 2)
		if page.Length() != 2 {
			t.Fatal("wrong page")
		}
	})
}

func Test_Cov21_Collection_SortedOps(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_SortedOps", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		asc := c.SortedListAsc()
		if asc[0] != "a" {
			t.Fatal("wrong sort")
		}
		dsc := c.SortedListDsc()
		if dsc[0] != "c" {
			t.Fatal("wrong sort")
		}
		c.SortedAsc()
		if c.First() != "a" {
			t.Fatal("wrong")
		}
	})
}

func Test_Cov21_Collection_SortedAscLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_SortedAscLock", func() {
		c := corestr.New.Collection.Strings([]string{"b", "a"})
		c.SortedAscLock()
		if c.First() != "a" {
			t.Fatal("wrong")
		}
	})
}

func Test_Cov21_Collection_HasUsingSensitivity(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_HasUsingSensitivity", func() {
		c := corestr.New.Collection.Strings([]string{"Hello"})
		if !c.HasUsingSensitivity("Hello", true) {
			t.Fatal("expected true")
		}
		if c.HasUsingSensitivity("hello", true) {
			t.Fatal("expected false for case sensitive")
		}
		if !c.HasUsingSensitivity("hello", false) {
			t.Fatal("expected true for case insensitive")
		}
	})
}

func Test_Cov21_Collection_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_IsContains", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		str := "a"
		if !c.IsContainsPtr(&str) {
			t.Fatal("expected true")
		}
		if c.IsContainsPtr(nil) {
			t.Fatal("expected false for nil")
		}
	})
}

func Test_Cov21_Collection_IsContainsAll(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_IsContainsAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		if !c.IsContainsAll("a", "b") {
			t.Fatal("expected true")
		}
		if c.IsContainsAll("a", "x") {
			t.Fatal("expected false")
		}
		if c.IsContainsAll(nil...) {
			t.Fatal("expected false for nil")
		}
	})
}

func Test_Cov21_Collection_IsContainsAllSlice(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_IsContainsAllSlice", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !c.IsContainsAllSlice([]string{"a", "b"}) {
			t.Fatal("expected true")
		}
		if c.IsContainsAllSlice([]string{}) {
			t.Fatal("expected false for empty")
		}
	})
}

func Test_Cov21_Collection_IsContainsAllLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_IsContainsAllLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !c.IsContainsAllLock("a", "b") {
			t.Fatal("expected true")
		}
	})
}

func Test_Cov21_Collection_HasAll(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_HasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !c.HasAll("a", "b") {
			t.Fatal("expected true")
		}
	})
}

func Test_Cov21_Collection_HasLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_HasLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if !c.HasLock("a") {
			t.Fatal("expected true")
		}
	})
}

func Test_Cov21_Collection_Has_HasPtr(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Has_HasPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if !c.Has("a") {
			t.Fatal("expected true")
		}
		s := "a"
		if !c.HasPtr(&s) {
			t.Fatal("expected true")
		}
		if c.HasPtr(nil) {
			t.Fatal("expected false for nil")
		}
	})
}

func Test_Cov21_Collection_GetHashsetPlusHasAll(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_GetHashsetPlusHasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs, hasAll := c.GetHashsetPlusHasAll([]string{"a"})
		if hs == nil || !hasAll {
			t.Fatal("unexpected")
		}
		_, hasAll2 := c.GetHashsetPlusHasAll(nil)
		if hasAll2 {
			t.Fatal("expected false for nil")
		}
	})
}

func Test_Cov21_Collection_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_GetAllExcept", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.GetAllExcept([]string{"b"})
		if len(result) != 2 {
			t.Fatal("expected 2")
		}
		all := c.GetAllExcept(nil)
		if len(all) != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_Cov21_Collection_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_GetAllExceptCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		exc := corestr.New.Collection.Strings([]string{"a"})
		result := c.GetAllExceptCollection(exc)
		if len(result) != 2 {
			t.Fatal("expected 2")
		}
		all := c.GetAllExceptCollection(nil)
		if len(all) != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_Cov21_Collection_UniqueBoolMap(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_UniqueBoolMap", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a", "b"})
		m := c.UniqueBoolMap()
		if len(m) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_UniqueBoolMapLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_UniqueBoolMapLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a"})
		m := c.UniqueBoolMapLock()
		if len(m) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov21_Collection_UniqueList(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_UniqueList", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a", "b"})
		list := c.UniqueList()
		if len(list) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_UniqueListLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_UniqueListLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a"})
		list := c.UniqueListLock()
		if len(list) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov21_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Filter", func() {
		c := corestr.New.Collection.Strings([]string{"a", "bb", "ccc"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})
		if len(result) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_FilteredCollection(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_FilteredCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		fc := c.FilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})
		if fc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}
func Test_Cov21_Collection_NonEmptyItems(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_NonEmptyItems", func() {
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		items := c.NonEmptyItems()
		if len(items) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_HashsetOps(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_HashsetOps", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		_ = c.HashsetAsIs()
		_ = c.HashsetWithDoubleLength()
	})
}
func Test_Cov21_Collection_ListCopyPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_ListCopyPtrLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		list := c.ListCopyPtrLock()
		if len(list) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov21_Collection_Join_Joins(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Join_Joins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.Join(",") != "a,b" {
			t.Fatal("wrong join")
		}
		_ = c.JoinLine()
		_ = c.Joins(",", "c")
		_ = c.NonEmptyJoins(",")
		_ = c.NonWhitespaceJoins(",")
	})
}

func Test_Cov21_Collection_Csv(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Csv", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		_ = c.Csv()
		_ = c.CsvOptions(true)
		_ = c.CsvLines()
		_ = c.CsvLinesOptions(true)
	})
}

func Test_Cov21_Collection_String_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_String_StringLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.String() == "" {
			t.Fatal("expected non-empty")
		}
		if c.StringLock() == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_Cov21_Collection_JsonOps(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_JsonOps", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.JsonModel()
		_ = c.JsonModelAny()
		_ = c.JsonString()
		_ = c.JsonStringMust()
		_ = c.StringJSON()
		j := c.Json()
		if j.HasError() {
			t.Fatal("unexpected error")
		}
		jp := c.JsonPtr()
		if jp.HasError() {
			t.Fatal("unexpected error")
		}
	})
}

func Test_Cov21_Collection_MarshalUnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_MarshalUnmarshalJSON", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		b, err := c.MarshalJSON()
		if err != nil || len(b) == 0 {
			t.Fatal("unexpected")
		}
		c2 := corestr.New.Collection.Empty()
		err = c2.UnmarshalJSON(b)
		if err != nil || c2.Length() != 2 {
			t.Fatal("unexpected")
		}
	})
}

func Test_Cov21_Collection_ParseInject(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_ParseInject", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		// Use json.Marshal with pointer to bypass value receiver issue on JsonPtr
		b, _ := json.Marshal(c)
		jr := &corejson.Result{Bytes: b}
		target := corestr.New.Collection.Empty()
		_, err := target.ParseInjectUsingJson(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
		target2 := corestr.New.Collection.Empty()
		_ = target2.ParseInjectUsingJsonMust(jr)
	})
}

func Test_Cov21_Collection_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_ClearDispose", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Clear()
		if c.Length() != 0 {
			t.Fatal("expected 0")
		}
		c.Add("b")
		c.Dispose()
	})
}

func Test_Cov21_Collection_Capacity_Resize(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Capacity_Resize", func() {
		c := corestr.New.Collection.Cap(10)
		if c.Capacity() < 10 {
			t.Fatal("expected capacity >= 10")
		}
		c.Resize(20)
		if c.Capacity() < 20 {
			t.Fatal("expected capacity >= 20")
		}
		c.AddCapacity(5)
	})
}

func Test_Cov21_Collection_SerializeDeserialize(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_SerializeDeserialize", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		b, err := c.Serialize()
		if err != nil {
			t.Fatal("unexpected")
		}
		var target []string
		err = c.Deserialize(&target)
		_ = err
		_ = b
	})
}

func Test_Cov21_Collection_InterfaceMethods(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_InterfaceMethods", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.AsJsonMarshaller()
		_ = c.AsJsonContractsBinder()
	})
}

func Test_Cov21_Collection_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_JsonParseSelfInject", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		// Use json.Marshal with pointer to bypass value receiver issue on JsonPtr
		b, _ := json.Marshal(c)
		jr := &corejson.Result{Bytes: b}
		target := corestr.New.Collection.Empty()
		err := target.JsonParseSelfInject(jr)
		if err != nil {
			t.Fatal("unexpected")
		}
	})
}

func Test_Cov21_Collection_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddWithWgLock", func() {
		c := corestr.New.Collection.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.AddWithWgLock(wg, "a")
		wg.Wait()
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov21_Collection_AddPointerCollectionsLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddPointerCollectionsLock", func() {
		c := corestr.New.Collection.Empty()
		c2 := corestr.New.Collection.Strings([]string{"a"})
		c.AddPointerCollectionsLock(c2)
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov21_Collection_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddHashmapsValues", func() {
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.Cap(2)
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsValues(hm)
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov21_Collection_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddHashmapsKeys", func() {
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.Cap(2)
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsKeys(hm)
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov21_Collection_AddHashmapsKeysValues(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddHashmapsKeysValues", func() {
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.Cap(2)
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsKeysValues(hm)
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AppendAnys", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnys("a", nil, 42)
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_AppendNonEmptyAnys(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AppendNonEmptyAnys", func() {
		c := corestr.New.Collection.Empty()
		c.AppendNonEmptyAnys("a", nil, "")
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov21_Collection_AddsNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddsNonEmpty", func() {
		c := corestr.New.Collection.Empty()
		c.AddsNonEmpty("a", "", "b")
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_New(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_New", func() {
		c := corestr.New.Collection.Empty()
		newC := c.New("a", "b")
		if newC.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_AddNonEmptyStrings(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddNonEmptyStrings", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStrings("a", "", "b")
		if c.Length() != 2 { // AddNonEmptyStrings filters empty strings
			t.Fatalf("expected 2, got %d", c.Length())
		}
	})
}

func Test_Cov21_Collection_AddFuncResult(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddFuncResult", func() {
		c := corestr.New.Collection.Empty()
		c.AddFuncResult(func() string { return "a" })
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov21_Collection_AddStringsByFuncChecking(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddStringsByFuncChecking", func() {
		c := corestr.New.Collection.Empty()
		c.AddStringsByFuncChecking([]string{"a", "", "b"}, func(s string) bool { return s != "" })
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_ExpandSlicePlusAdd(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_ExpandSlicePlusAdd", func() {
		c := corestr.New.Collection.Empty()
		c.ExpandSlicePlusAdd([]string{"a,b"}, func(line string) []string {
			return []string{line + "1", line + "2"}
		})
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_MergeSlicesOfSlice(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_MergeSlicesOfSlice", func() {
		c := corestr.New.Collection.Empty()
		c.MergeSlicesOfSlice([]string{"a"}, []string{"b"})
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_IsEqualsWithSensitive(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_IsEqualsWithSensitive", func() {
		a := corestr.New.Collection.Strings([]string{"Hello"})
		b := corestr.New.Collection.Strings([]string{"hello"})
		if a.IsEqualsWithSensitive(true, b) {
			t.Fatal("expected false for case sensitive")
		}
		if !a.IsEqualsWithSensitive(false, b) {
			t.Fatal("expected true for case insensitive")
		}
	})
}

func Test_Cov21_Collection_AppendAnysUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AppendAnysUsingFilter", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnysUsingFilter(func(s string, i int) (string, bool, bool) {
			return s, s != "", false
		}, "a", nil, "b")
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_ChainRemoveAt(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_ChainRemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.ChainRemoveAt(1)
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_CharCollectionMap", func() {
		c := corestr.New.Collection.Strings([]string{"abc", "bcd"})
		ccm := c.CharCollectionMap()
		if ccm == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_Cov21_Collection_SummaryString(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_SummaryString", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.SummaryString(1)
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_Cov21_Collection_SummaryStringWithHeader(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_SummaryStringWithHeader", func() {
		c := corestr.New.Collection.Empty()
		s := c.SummaryStringWithHeader("header")
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_Cov21_Collection_Capacity(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Capacity", func() {
		var c *corestr.Collection
		_ = fmt.Sprint(c) // should not panic
	})
}

func Test_Cov21_Collection_Single(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Single", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.Single() != "a" {
			t.Fatal("wrong")
		}
	})
}

func Test_Cov21_Collection_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_HasAnyItem", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if !c.HasAnyItem() {
			t.Fatal("expected true")
		}
	})
}

func Test_Cov21_Collection_InsertAt(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_InsertAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "c"})
		c.InsertAt(1, "b")
	})
}

func Test_Cov21_Collection_RemoveItemsIndexes(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_RemoveItemsIndexes", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.RemoveItemsIndexes(true, 1)
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_AppendCollectionPtr(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AppendCollectionPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c.AppendCollectionPtr(c2)
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_AppendCollections(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AppendCollections", func() {
		c := corestr.New.Collection.Empty()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c.AppendCollections(c1, c2)
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_NonEmptyItemsOrNonWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_NonEmptyItemsOrNonWhitespace", func() {
		c := corestr.New.Collection.Strings([]string{"a", "", "  ", "b"})
		items := c.NonEmptyItemsOrNonWhitespace()
		if len(items) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_HashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_HashsetLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs := c.HashsetLock()
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov21_Collection_FilterLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_FilterLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		result := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})
		if len(result) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov21_Collection_FilterPtr(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_FilterPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		result := c.FilterPtr(func(s *string, i int) (*string, bool, bool) {
			return s, len(*s) > 1, false
		})
		if len(*result) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov21_Collection_FilterPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_FilterPtrLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		result := c.FilterPtrLock(func(s *string, i int) (*string, bool, bool) {
			return s, len(*s) > 1, false
		})
		if len(*result) != 1 {
			t.Fatal("expected 1")
		}
	})
}
