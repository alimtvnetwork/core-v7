package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — Segment 7: Remaining methods (L700-1469)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovHS2_01_OrderedList(t *testing.T) {
	safeTest(t, "Test_CovHS2_01_OrderedList", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("c", "a", "b")
		list := hs.OrderedList()
		if len(list) != 3 {
			t.Fatal("expected 3")
		}
		// empty
		e := corestr.New.Hashset.Empty()
		list2 := e.OrderedList()
		if len(list2) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovHS2_02_SafeStrings(t *testing.T) {
	safeTest(t, "Test_CovHS2_02_SafeStrings", func() {
		hs := corestr.New.Hashset.Empty()
		if len(hs.SafeStrings()) != 0 {
			t.Fatal("expected 0")
		}
		hs.Add("a")
		if len(hs.SafeStrings()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHS2_03_Lines(t *testing.T) {
	safeTest(t, "Test_CovHS2_03_Lines", func() {
		hs := corestr.New.Hashset.Empty()
		if len(hs.Lines()) != 0 {
			t.Fatal("expected 0")
		}
		hs.Add("a")
		if len(hs.Lines()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHS2_04_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_CovHS2_04_SimpleSlice", func() {
		hs := corestr.New.Hashset.Empty()
		ss := hs.SimpleSlice()
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
		hs.Add("a")
		ss2 := hs.SimpleSlice()
		if ss2.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHS2_05_GetFilteredItems(t *testing.T) {
	safeTest(t, "Test_CovHS2_05_GetFilteredItems", func() {
		hs := corestr.New.Hashset.Empty()
		// empty
		r := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if len(r) != 0 {
			t.Fatal("expected 0")
		}
		// with items, keep all
		hs.Adds("a", "b")
		r2 := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if len(r2) != 2 {
			t.Fatalf("expected 2, got %d", len(r2))
		}
		// skip
		r3 := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, false, false
		})
		if len(r3) != 0 {
			t.Fatal("expected 0")
		}
		// break
		r4 := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})
		if len(r4) != 1 {
			t.Fatal("expected 1 (break)")
		}
	})
}

func Test_CovHS2_06_GetFilteredCollection(t *testing.T) {
	safeTest(t, "Test_CovHS2_06_GetFilteredCollection", func() {
		hs := corestr.New.Hashset.Empty()
		// empty
		col := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
		// with items
		hs.Adds("a", "b")
		col2 := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if col2.Length() != 2 {
			t.Fatal("expected 2")
		}
		// break
		col3 := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})
		if col3.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHS2_07_GetAllExceptHashset(t *testing.T) {
	safeTest(t, "Test_CovHS2_07_GetAllExceptHashset", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b", "c")
		except := corestr.New.Hashset.Empty()
		except.Add("b")
		result := hs.GetAllExceptHashset(except)
		if len(result) != 2 {
			t.Fatalf("expected 2, got %d", len(result))
		}
		// nil
		r2 := hs.GetAllExceptHashset(nil)
		if len(r2) != 3 {
			t.Fatal("expected 3")
		}
		// empty
		r3 := hs.GetAllExceptHashset(corestr.New.Hashset.Empty())
		if len(r3) != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_CovHS2_08_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_CovHS2_08_GetAllExcept", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		r := hs.GetAllExcept([]string{"a"})
		if len(r) != 1 {
			t.Fatal("expected 1")
		}
		// nil
		r2 := hs.GetAllExcept(nil)
		if len(r2) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovHS2_09_GetAllExceptSpread(t *testing.T) {
	safeTest(t, "Test_CovHS2_09_GetAllExceptSpread", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		r := hs.GetAllExceptSpread("a")
		if len(r) != 1 {
			t.Fatal("expected 1")
		}
		// nil
		r2 := hs.GetAllExceptSpread()
		if len(r2) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovHS2_10_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_CovHS2_10_GetAllExceptCollection", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		r := hs.GetAllExceptCollection(corestr.New.Collection.Strings([]string{"a"}))
		if len(r) != 1 {
			t.Fatal("expected 1")
		}
		// nil
		r2 := hs.GetAllExceptCollection(nil)
		if len(r2) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovHS2_11_Items(t *testing.T) {
	safeTest(t, "Test_CovHS2_11_Items", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		items := hs.Items()
		if !items["a"] {
			t.Fatal("expected a")
		}
	})
}
func Test_CovHS2_13_MapStringAny_MapStringAnyDiff(t *testing.T) {
	safeTest(t, "Test_CovHS2_13_MapStringAny_MapStringAnyDiff", func() {
		hs := corestr.New.Hashset.Empty()
		// empty
		m := hs.MapStringAny()
		if len(m) != 0 {
			t.Fatal("expected 0")
		}
		hs.Add("a")
		m2 := hs.MapStringAny()
		if len(m2) != 1 {
			t.Fatal("expected 1")
		}
		_ = hs.MapStringAnyDiff()
	})
}

func Test_CovHS2_14_JoinSorted(t *testing.T) {
	safeTest(t, "Test_CovHS2_14_JoinSorted", func() {
		hs := corestr.New.Hashset.Empty()
		if hs.JoinSorted(",") != "" {
			t.Fatal("expected empty")
		}
		hs.Adds("b", "a")
		s := hs.JoinSorted(",")
		if s != "a,b" {
			t.Fatalf("expected 'a,b', got '%s'", s)
		}
	})
}

func Test_CovHS2_15_ListPtrSortedAsc_Dsc(t *testing.T) {
	safeTest(t, "Test_CovHS2_15_ListPtrSortedAsc_Dsc", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("c", "a", "b")
		asc := hs.ListPtrSortedAsc()
		if asc[0] != "a" {
			t.Fatal("expected a first")
		}
		dsc := hs.ListPtrSortedDsc()
		if dsc[0] != "c" {
			t.Fatal("expected c first")
		}
	})
}

func Test_CovHS2_16_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_CovHS2_16_Clear_Dispose", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		hs.Clear()
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
		hs2 := corestr.New.Hashset.Empty()
		hs2.Add("x")
		hs2.Dispose()
	})
}

func Test_CovHS2_17_Remove_SafeRemove_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_CovHS2_17_Remove_SafeRemove_RemoveWithLock", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b", "c")
		hs.Remove("a")
		if hs.Has("a") {
			t.Fatal("expected removed")
		}
		hs.SafeRemove("b")
		if hs.Has("b") {
			t.Fatal("expected removed")
		}
		hs.SafeRemove("nonexist")
		hs.RemoveWithLock("c")
		if hs.Has("c") {
			t.Fatal("expected removed")
		}
	})
}

func Test_CovHS2_18_String_StringLock(t *testing.T) {
	safeTest(t, "Test_CovHS2_18_String_StringLock", func() {
		hs := corestr.New.Hashset.Empty()
		s := hs.String()
		if s == "" {
			t.Fatal("expected non-empty (NoElements)")
		}
		hs.Add("a")
		s2 := hs.String()
		if s2 == "" {
			t.Fatal("expected non-empty")
		}
		_ = hs.StringLock()
		_ = corestr.New.Hashset.Empty().StringLock()
	})
}

func Test_CovHS2_19_Join_JoinLine(t *testing.T) {
	safeTest(t, "Test_CovHS2_19_Join_JoinLine", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		s := hs.Join(",")
		if s != "a" {
			t.Fatalf("expected 'a', got '%s'", s)
		}
		_ = hs.JoinLine()
	})
}

func Test_CovHS2_20_NonEmptyJoins_NonWhitespaceJoins(t *testing.T) {
	safeTest(t, "Test_CovHS2_20_NonEmptyJoins_NonWhitespaceJoins", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		_ = hs.NonEmptyJoins(",")
		_ = hs.NonWhitespaceJoins(",")
	})
}

func Test_CovHS2_21_ToLowerSet(t *testing.T) {
	safeTest(t, "Test_CovHS2_21_ToLowerSet", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("ABC", "XYZ")
		lower := hs.ToLowerSet()
		if !lower.Has("abc") {
			t.Fatal("expected abc")
		}
	})
}

func Test_CovHS2_22_Length_LengthLock(t *testing.T) {
	safeTest(t, "Test_CovHS2_22_Length_LengthLock", func() {
		hs := corestr.New.Hashset.Empty()
		if hs.Length() != 0 {
			t.Fatal("expected 0")
		}
		if hs.LengthLock() != 0 {
			t.Fatal("expected 0")
		}
		hs.Add("a")
		if hs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHS2_23_JsonModel_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_CovHS2_23_JsonModel_JsonModelAny", func() {
		hs := corestr.New.Hashset.Empty()
		m := hs.JsonModel()
		if len(m) != 0 {
			t.Fatal("expected 0")
		}
		hs.Add("a")
		m2 := hs.JsonModel()
		if len(m2) != 1 {
			t.Fatal("expected 1")
		}
		_ = hs.JsonModelAny()
	})
}

func Test_CovHS2_24_MarshalJSON_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovHS2_24_MarshalJSON_UnmarshalJSON", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		data, err := hs.MarshalJSON()
		if err != nil {
			t.Fatal("unexpected error")
		}
		hs2 := corestr.New.Hashset.Empty()
		err2 := hs2.UnmarshalJSON(data)
		if err2 != nil {
			t.Fatal("unexpected error")
		}
		if hs2.Length() != 2 {
			t.Fatal("expected 2")
		}
		// invalid
		err3 := hs2.UnmarshalJSON([]byte("invalid"))
		if err3 == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_CovHS2_25_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_CovHS2_25_Json_JsonPtr", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		_ = hs.Json()
		_ = hs.JsonPtr()
	})
}

func Test_CovHS2_26_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_CovHS2_26_ParseInjectUsingJson", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		jr := hs.JsonPtr()
		hs2 := corestr.New.Hashset.Empty()
		result, err := hs2.ParseInjectUsingJson(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
		if result.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovHS2_27_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovHS2_27_ParseInjectUsingJsonMust", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		jr := hs.JsonPtr()
		hs2 := corestr.New.Hashset.Empty()
		r := hs2.ParseInjectUsingJsonMust(jr)
		if r.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHS2_28_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovHS2_28_JsonParseSelfInject", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		jr := hs.JsonPtr()
		hs2 := corestr.New.Hashset.Empty()
		err := hs2.JsonParseSelfInject(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovHS2_29_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_CovHS2_29_AsInterfaces", func() {
		hs := corestr.New.Hashset.Empty()
		_ = hs.AsJsonContractsBinder()
		_ = hs.AsJsoner()
		_ = hs.AsJsonParseSelfInjector()
		_ = hs.AsJsonMarshaller()
	})
}

func Test_CovHS2_30_DistinctDiffLinesRaw(t *testing.T) {
	safeTest(t, "Test_CovHS2_30_DistinctDiffLinesRaw", func() {
		hs := corestr.New.Hashset.Empty()
		// both empty
		r := hs.DistinctDiffLinesRaw()
		if len(r) != 0 {
			t.Fatal("expected 0")
		}
		// left empty, right has items
		r2 := hs.DistinctDiffLinesRaw("a")
		if len(r2) != 1 {
			t.Fatal("expected 1")
		}
		// left has items, right empty
		hs.Add("x")
		r3 := hs.DistinctDiffLinesRaw()
		if len(r3) != 1 {
			t.Fatal("expected 1")
		}
		// both have items
		r4 := hs.DistinctDiffLinesRaw("a", "x")
		if len(r4) != 1 {
			t.Fatalf("expected 1 (only 'a'), got %d", len(r4))
		}
	})
}

func Test_CovHS2_31_DistinctDiffHashset(t *testing.T) {
	safeTest(t, "Test_CovHS2_31_DistinctDiffHashset", func() {
		a := corestr.New.Hashset.Empty()
		a.Adds("a", "b")
		b := corestr.New.Hashset.Empty()
		b.Adds("b", "c")
		diff := a.DistinctDiffHashset(b)
		if len(diff) != 2 {
			t.Fatalf("expected 2, got %d", len(diff))
		}
	})
}

func Test_CovHS2_32_DistinctDiffLines(t *testing.T) {
	safeTest(t, "Test_CovHS2_32_DistinctDiffLines", func() {
		hs := corestr.New.Hashset.Empty()
		// both empty
		r := hs.DistinctDiffLines()
		if len(r) != 0 {
			t.Fatal("expected 0")
		}
		// left not empty, right empty
		hs.Add("x")
		r2 := hs.DistinctDiffLines()
		if len(r2) != 1 {
			t.Fatal("expected 1")
		}
		// left empty, right not empty
		e := corestr.New.Hashset.Empty()
		r3 := e.DistinctDiffLines("a")
		if len(r3) != 1 {
			t.Fatal("expected 1")
		}
		// both have items
		r4 := hs.DistinctDiffLines("a", "x")
		if len(r4) != 1 {
			t.Fatalf("expected 1, got %d", len(r4))
		}
	})
}

func Test_CovHS2_33_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_CovHS2_33_Serialize_Deserialize", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		_, err := hs.Serialize()
		if err != nil {
			t.Fatal("unexpected error")
		}
		target := corestr.New.Hashset.Empty()
		err2 := hs.Deserialize(target)
		if err2 != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovHS2_34_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_CovHS2_34_WrapDoubleQuote", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		r := hs.WrapDoubleQuote()
		if !r.Has(`"a"`) {
			t.Fatal("expected wrapped")
		}
	})
}

func Test_CovHS2_35_WrapDoubleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_CovHS2_35_WrapDoubleQuoteIfMissing", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		r := hs.WrapDoubleQuoteIfMissing()
		if !r.Has(`"a"`) {
			t.Fatal("expected wrapped")
		}
	})
}

func Test_CovHS2_36_WrapSingleQuote(t *testing.T) {
	safeTest(t, "Test_CovHS2_36_WrapSingleQuote", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		r := hs.WrapSingleQuote()
		if !r.Has("'a'") {
			t.Fatal("expected wrapped")
		}
	})
}

func Test_CovHS2_37_WrapSingleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_CovHS2_37_WrapSingleQuoteIfMissing", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		r := hs.WrapSingleQuoteIfMissing()
		if !r.Has("'a'") {
			t.Fatal("expected wrapped")
		}
	})
}

func Test_CovHS2_38_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_CovHS2_38_Transpile_Empty", func() {
		hs := corestr.New.Hashset.Empty()
		r := hs.Transpile(func(s string) string { return s })
		if r.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}
