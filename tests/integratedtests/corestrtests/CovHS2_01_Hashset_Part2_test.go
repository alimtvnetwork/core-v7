package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — Segment 7: Remaining methods (L700-1469)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovHS2_01_OrderedList(t *testing.T) {
	safeTest(t, "Test_CovHS2_01_OrderedList", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("c", "a", "b")
		list := hs.OrderedList()

		// Act
		actual := args.Map{"result": len(list) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// empty
		e := corestr.New.Hashset.Empty()
		list2 := e.OrderedList()
		actual = args.Map{"result": len(list2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovHS2_02_SafeStrings(t *testing.T) {
	safeTest(t, "Test_CovHS2_02_SafeStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"result": len(hs.SafeStrings()) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hs.Add("a")
		actual = args.Map{"result": len(hs.SafeStrings()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS2_03_Lines(t *testing.T) {
	safeTest(t, "Test_CovHS2_03_Lines", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"result": len(hs.Lines()) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hs.Add("a")
		actual = args.Map{"result": len(hs.Lines()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS2_04_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_CovHS2_04_SimpleSlice", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		ss := hs.SimpleSlice()

		// Act
		actual := args.Map{"result": ss.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hs.Add("a")
		ss2 := hs.SimpleSlice()
		actual = args.Map{"result": ss2.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS2_05_GetFilteredItems(t *testing.T) {
	safeTest(t, "Test_CovHS2_05_GetFilteredItems", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// empty
		r := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"result": len(r) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// with items, keep all
		hs.Adds("a", "b")
		r2 := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual = args.Map{"result": len(r2) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// skip
		r3 := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, false, false
		})
		actual = args.Map{"result": len(r3) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// break
		r4 := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})
		actual = args.Map{"result": len(r4) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 (break)", actual)
	})
}

func Test_CovHS2_06_GetFilteredCollection(t *testing.T) {
	safeTest(t, "Test_CovHS2_06_GetFilteredCollection", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// empty
		col := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"result": col.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// with items
		hs.Adds("a", "b")
		col2 := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual = args.Map{"result": col2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// break
		col3 := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})
		actual = args.Map{"result": col3.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS2_07_GetAllExceptHashset(t *testing.T) {
	safeTest(t, "Test_CovHS2_07_GetAllExceptHashset", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b", "c")
		except := corestr.New.Hashset.Empty()
		except.Add("b")
		result := hs.GetAllExceptHashset(except)

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// nil
		r2 := hs.GetAllExceptHashset(nil)
		actual = args.Map{"result": len(r2) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// empty
		r3 := hs.GetAllExceptHashset(corestr.New.Hashset.Empty())
		actual = args.Map{"result": len(r3) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_CovHS2_08_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_CovHS2_08_GetAllExcept", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		r := hs.GetAllExcept([]string{"a"})

		// Act
		actual := args.Map{"result": len(r) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// nil
		r2 := hs.GetAllExcept(nil)
		actual = args.Map{"result": len(r2) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovHS2_09_GetAllExceptSpread(t *testing.T) {
	safeTest(t, "Test_CovHS2_09_GetAllExceptSpread", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		r := hs.GetAllExceptSpread("a")

		// Act
		actual := args.Map{"result": len(r) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// nil
		r2 := hs.GetAllExceptSpread()
		actual = args.Map{"result": len(r2) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovHS2_10_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_CovHS2_10_GetAllExceptCollection", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		r := hs.GetAllExceptCollection(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": len(r) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// nil
		r2 := hs.GetAllExceptCollection(nil)
		actual = args.Map{"result": len(r2) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovHS2_11_Items(t *testing.T) {
	safeTest(t, "Test_CovHS2_11_Items", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		items := hs.Items()

		// Act
		actual := args.Map{"result": items["a"]}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_CovHS2_12_List_ListPtr_ListCopyLock(t *testing.T) {
	safeTest(t, "Test_CovHS2_12_List_ListPtr_ListCopyLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")

		// Act
		actual := args.Map{"result": len(hs.List()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": len(hs.ListPtr()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": len(hs.ListCopyLock()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS2_13_MapStringAny_MapStringAnyDiff(t *testing.T) {
	safeTest(t, "Test_CovHS2_13_MapStringAny_MapStringAnyDiff", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// empty
		m := hs.MapStringAny()

		// Act
		actual := args.Map{"result": len(m) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hs.Add("a")
		m2 := hs.MapStringAny()
		actual = args.Map{"result": len(m2) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		_ = hs.MapStringAnyDiff()
	})
}

func Test_CovHS2_14_JoinSorted(t *testing.T) {
	safeTest(t, "Test_CovHS2_14_JoinSorted", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"result": hs.JoinSorted(",") != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		hs.Adds("b", "a")
		s := hs.JoinSorted(",")
		actual = args.Map{"result": s != "a,b"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
	})
}

func Test_CovHS2_15_ListPtrSortedAsc_Dsc(t *testing.T) {
	safeTest(t, "Test_CovHS2_15_ListPtrSortedAsc_Dsc", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("c", "a", "b")
		asc := hs.ListPtrSortedAsc()

		// Act
		actual := args.Map{"result": asc[0] != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a first", actual)
		dsc := hs.ListPtrSortedDsc()
		actual = args.Map{"result": dsc[0] != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c first", actual)
	})
}

func Test_CovHS2_16_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_CovHS2_16_Clear_Dispose", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		hs.Clear()

		// Act
		actual := args.Map{"result": hs.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hs2 := corestr.New.Hashset.Empty()
		hs2.Add("x")
		hs2.Dispose()
	})
}

func Test_CovHS2_17_Remove_SafeRemove_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_CovHS2_17_Remove_SafeRemove_RemoveWithLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b", "c")
		hs.Remove("a")

		// Act
		actual := args.Map{"result": hs.Has("a")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed", actual)
		hs.SafeRemove("b")
		actual = args.Map{"result": hs.Has("b")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed", actual)
		hs.SafeRemove("nonexist")
		hs.RemoveWithLock("c")
		actual = args.Map{"result": hs.Has("c")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed", actual)
	})
}

func Test_CovHS2_18_String_StringLock(t *testing.T) {
	safeTest(t, "Test_CovHS2_18_String_StringLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		s := hs.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty (NoElements)", actual)
		hs.Add("a")
		s2 := hs.String()
		actual = args.Map{"result": s2 == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		_ = hs.StringLock()
		_ = corestr.New.Hashset.Empty().StringLock()
	})
}

func Test_CovHS2_19_Join_JoinLine(t *testing.T) {
	safeTest(t, "Test_CovHS2_19_Join_JoinLine", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		s := hs.Join(",")

		// Act
		actual := args.Map{"result": s != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a', got ''", actual)
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("ABC", "XYZ")
		lower := hs.ToLowerSet()

		// Act
		actual := args.Map{"result": lower.Has("abc")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected abc", actual)
	})
}

func Test_CovHS2_22_Length_LengthLock(t *testing.T) {
	safeTest(t, "Test_CovHS2_22_Length_LengthLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"result": hs.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": hs.LengthLock() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hs.Add("a")
		actual = args.Map{"result": hs.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS2_23_JsonModel_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_CovHS2_23_JsonModel_JsonModelAny", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		m := hs.JsonModel()

		// Act
		actual := args.Map{"result": len(m) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hs.Add("a")
		m2 := hs.JsonModel()
		actual = args.Map{"result": len(m2) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		_ = hs.JsonModelAny()
	})
}

func Test_CovHS2_24_MarshalJSON_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovHS2_24_MarshalJSON_UnmarshalJSON", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		data, err := hs.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		hs2 := corestr.New.Hashset.Empty()
		err2 := hs2.UnmarshalJSON(data)
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual = args.Map{"result": hs2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// invalid
		err3 := hs2.UnmarshalJSON([]byte("invalid"))
		actual = args.Map{"result": err3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		jr := hs.JsonPtr()
		hs2 := corestr.New.Hashset.Empty()
		result, err := hs2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual = args.Map{"result": result.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovHS2_27_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovHS2_27_ParseInjectUsingJsonMust", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		jr := hs.JsonPtr()
		hs2 := corestr.New.Hashset.Empty()
		r := hs2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"result": r.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS2_28_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovHS2_28_JsonParseSelfInject", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		jr := hs.JsonPtr()
		hs2 := corestr.New.Hashset.Empty()
		err := hs2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// both empty
		r := hs.DistinctDiffLinesRaw()

		// Act
		actual := args.Map{"result": len(r) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// left empty, right has items
		r2 := hs.DistinctDiffLinesRaw("a")
		actual = args.Map{"result": len(r2) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// left has items, right empty
		hs.Add("x")
		r3 := hs.DistinctDiffLinesRaw()
		actual = args.Map{"result": len(r3) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// both have items
		r4 := hs.DistinctDiffLinesRaw("a", "x")
		actual = args.Map{"result": len(r4) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 (only 'a')", actual)
	})
}

func Test_CovHS2_31_DistinctDiffHashset(t *testing.T) {
	safeTest(t, "Test_CovHS2_31_DistinctDiffHashset", func() {
		// Arrange
		a := corestr.New.Hashset.Empty()
		a.Adds("a", "b")
		b := corestr.New.Hashset.Empty()
		b.Adds("b", "c")
		diff := a.DistinctDiffHashset(b)

		// Act
		actual := args.Map{"result": len(diff) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovHS2_32_DistinctDiffLines(t *testing.T) {
	safeTest(t, "Test_CovHS2_32_DistinctDiffLines", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// both empty
		r := hs.DistinctDiffLines()

		// Act
		actual := args.Map{"result": len(r) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// left not empty, right empty
		hs.Add("x")
		r2 := hs.DistinctDiffLines()
		actual = args.Map{"result": len(r2) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// left empty, right not empty
		e := corestr.New.Hashset.Empty()
		r3 := e.DistinctDiffLines("a")
		actual = args.Map{"result": len(r3) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// both have items
		r4 := hs.DistinctDiffLines("a", "x")
		actual = args.Map{"result": len(r4) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS2_33_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_CovHS2_33_Serialize_Deserialize", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		_, err := hs.Serialize()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		target := corestr.New.Hashset.Empty()
		err2 := hs.Deserialize(target)
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_CovHS2_34_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_CovHS2_34_WrapDoubleQuote", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		r := hs.WrapDoubleQuote()

		// Act
		actual := args.Map{"result": r.Has(`"a"`)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
	})
}

func Test_CovHS2_35_WrapDoubleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_CovHS2_35_WrapDoubleQuoteIfMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		r := hs.WrapDoubleQuoteIfMissing()

		// Act
		actual := args.Map{"result": r.Has(`"a"`)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
	})
}

func Test_CovHS2_36_WrapSingleQuote(t *testing.T) {
	safeTest(t, "Test_CovHS2_36_WrapSingleQuote", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		r := hs.WrapSingleQuote()

		// Act
		actual := args.Map{"result": r.Has("'a'")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
	})
}

func Test_CovHS2_37_WrapSingleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_CovHS2_37_WrapSingleQuoteIfMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		r := hs.WrapSingleQuoteIfMissing()

		// Act
		actual := args.Map{"result": r.Has("'a'")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
	})
}

func Test_CovHS2_38_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_CovHS2_38_Transpile_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		r := hs.Transpile(func(s string) string { return s })

		// Act
		actual := args.Map{"result": r.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}
