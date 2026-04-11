package corestrtests

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Collection.go — comprehensive coverage for remaining uncovered methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov21_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_IndexAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.IndexAt(0) != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong index", actual)
	})
}

func Test_Cov21_Collection_SafeIndexAtUsingLength(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_SafeIndexAtUsingLength", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.SafeIndexAtUsingLength("def", 1, 0) != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
		actual = args.Map{"result": c.SafeIndexAtUsingLength("def", 1, 5) != "def"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default", actual)
	})
}

func Test_Cov21_Collection_First_Last(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_First_Last", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"result": c.First() != "a" || c.Last() != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
	})
}

func Test_Cov21_Collection_FirstOrDefault_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_FirstOrDefault_LastOrDefault", func() {
		// Arrange
		c := corestr.New.Collection.Empty()

		// Act
		actual := args.Map{"result": c.FirstOrDefault() != "" || c.LastOrDefault() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		c.Add("x")
		actual = args.Map{"result": c.FirstOrDefault() != "x" || c.LastOrDefault() != "x"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
	})
}

func Test_Cov21_Collection_Take(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Take", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		tk := c.Take(2)

		// Act
		actual := args.Map{"result": tk.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
		full := c.Take(10)
		actual = args.Map{"result": full.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
		zero := c.Take(0)
		actual = args.Map{"result": zero.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
	})
}

func Test_Cov21_Collection_Skip(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Skip", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		sk := c.Skip(1)

		// Act
		actual := args.Map{"result": sk.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
		same := c.Skip(0)
		actual = args.Map{"result": same.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
	})
}

func Test_Cov21_Collection_Reverse(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Reverse", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.Reverse()

		// Act
		actual := args.Map{"result": c.First() != "c" || c.Last() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
		// 2 items
		c2 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2.Reverse()
		actual = args.Map{"result": c2.First() != "b"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
		// 1 item
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c1.Reverse()
	})
}

func Test_Cov21_Collection_Paging(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Paging", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})

		// Act
		actual := args.Map{"result": c.GetPagesSize(2) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong pages", actual)
		actual = args.Map{"result": c.GetPagesSize(0) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		pages := c.GetPagedCollection(2)
		actual = args.Map{"result": pages.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		page := c.GetSinglePageCollection(2, 2)
		actual = args.Map{"result": page.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong page", actual)
	})
}

func Test_Cov21_Collection_SortedOps(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_SortedOps", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		asc := c.SortedListAsc()

		// Act
		actual := args.Map{"result": asc[0] != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong sort", actual)
		dsc := c.SortedListDsc()
		actual = args.Map{"result": dsc[0] != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong sort", actual)
		c.SortedAsc()
		actual = args.Map{"result": c.First() != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
	})
}

func Test_Cov21_Collection_SortedAscLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_SortedAscLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"b", "a"})
		c.SortedAscLock()

		// Act
		actual := args.Map{"result": c.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
	})
}

func Test_Cov21_Collection_HasUsingSensitivity(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_HasUsingSensitivity", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"Hello"})

		// Act
		actual := args.Map{"result": c.HasUsingSensitivity("Hello", true)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": c.HasUsingSensitivity("hello", true)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for case sensitive", actual)
		actual = args.Map{"result": c.HasUsingSensitivity("hello", false)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for case insensitive", actual)
	})
}

func Test_Cov21_Collection_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_IsContains", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		str := "a"

		// Act
		actual := args.Map{"result": c.IsContainsPtr(&str)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": c.IsContainsPtr(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_Cov21_Collection_IsContainsAll(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_IsContainsAll", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"result": c.IsContainsAll("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": c.IsContainsAll("a", "x")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": c.IsContainsAll(nil...)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_Cov21_Collection_IsContainsAllSlice(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_IsContainsAllSlice", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.IsContainsAllSlice([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": c.IsContainsAllSlice([]string{})}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_Cov21_Collection_IsContainsAllLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_IsContainsAllLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.IsContainsAllLock("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov21_Collection_HasAll(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_HasAll", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.HasAll("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov21_Collection_HasLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_HasLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.HasLock("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov21_Collection_Has_HasPtr(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Has_HasPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		s := "a"
		actual = args.Map{"result": c.HasPtr(&s)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": c.HasPtr(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_Cov21_Collection_GetHashsetPlusHasAll(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_GetHashsetPlusHasAll", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs, hasAll := c.GetHashsetPlusHasAll([]string{"a"})

		// Act
		actual := args.Map{"result": hs == nil || !hasAll}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
		_, hasAll2 := c.GetHashsetPlusHasAll(nil)
		actual = args.Map{"result": hasAll2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_Cov21_Collection_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_GetAllExcept", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.GetAllExcept([]string{"b"})

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		all := c.GetAllExcept(nil)
		actual = args.Map{"result": len(all) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov21_Collection_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_GetAllExceptCollection", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		exc := corestr.New.Collection.Strings([]string{"a"})
		result := c.GetAllExceptCollection(exc)

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		all := c.GetAllExceptCollection(nil)
		actual = args.Map{"result": len(all) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov21_Collection_UniqueBoolMap(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_UniqueBoolMap", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "a", "b"})
		m := c.UniqueBoolMap()

		// Act
		actual := args.Map{"result": len(m) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_UniqueBoolMapLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_UniqueBoolMapLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "a"})
		m := c.UniqueBoolMapLock()

		// Act
		actual := args.Map{"result": len(m) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov21_Collection_UniqueList(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_UniqueList", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "a", "b"})
		list := c.UniqueList()

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_UniqueListLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_UniqueListLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "a"})
		list := c.UniqueListLock()

		// Act
		actual := args.Map{"result": len(list) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov21_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Filter", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "bb", "ccc"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_FilteredCollection(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_FilteredCollection", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		fc := c.FilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})

		// Act
		actual := args.Map{"result": fc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov21_Collection_NonEmptyList(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_NonEmptyList", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		list := c.NonEmptyList()

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		_ = c.NonEmptyListPtr()
	})
}

func Test_Cov21_Collection_NonEmptyItems(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_NonEmptyItems", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		items := c.NonEmptyItems()

		// Act
		actual := args.Map{"result": len(items) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_HashsetOps(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_HashsetOps", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		_ = c.HashsetAsIs()
		_ = c.HashsetWithDoubleLength()
	})
}

func Test_Cov21_Collection_Items_List_ListPtr(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Items_List_ListPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.Items()
		_ = c.List()
		_ = c.ListPtr()
	})
}

func Test_Cov21_Collection_ListCopyPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_ListCopyPtrLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		list := c.ListCopyPtrLock()

		// Act
		actual := args.Map{"result": len(list) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov21_Collection_Join_Joins(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Join_Joins", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.Join(",") != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong join", actual)
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
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		actual = args.Map{"result": c.StringLock() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov21_Collection_JsonOps(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_JsonOps", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.JsonModel()
		_ = c.JsonModelAny()
		_ = c.JsonString()
		_ = c.JsonStringMust()
		_ = c.StringJSON()
		j := c.Json()

		// Act
		actual := args.Map{"result": j.HasError()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		jp := c.JsonPtr()
		actual = args.Map{"result": jp.HasError()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_Cov21_Collection_MarshalUnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_MarshalUnmarshalJSON", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		b, err := c.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil || len(b) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
		c2 := corestr.New.Collection.Empty()
		err = c2.UnmarshalJSON(b)
		actual = args.Map{"result": err != nil || c2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov21_Collection_ParseInject(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_ParseInject", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		// Use json.Marshal with pointer to bypass value receiver issue on JsonPtr
		b, _ := json.Marshal(c)
		jr := &corejson.Result{Bytes: b}
		target := corestr.New.Collection.Empty()
		_, err := target.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		target2 := corestr.New.Collection.Empty()
		_ = target2.ParseInjectUsingJsonMust(jr)
	})
}

func Test_Cov21_Collection_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_ClearDispose", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Clear()

		// Act
		actual := args.Map{"result": c.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		c.Add("b")
		c.Dispose()
	})
}

func Test_Cov21_Collection_Capacity_Resize(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_Capacity_Resize", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"result": c.Capacity() < 10}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected capacity >= 10", actual)
		c.Resize(20)
		actual = args.Map{"result": c.Capacity() < 20}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected capacity >= 20", actual)
		c.AddCapacity(5)
	})
}

func Test_Cov21_Collection_SerializeDeserialize(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_SerializeDeserialize", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		b, err := c.Serialize()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
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
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		// Use json.Marshal with pointer to bypass value receiver issue on JsonPtr
		b, _ := json.Marshal(c)
		jr := &corejson.Result{Bytes: b}
		target := corestr.New.Collection.Empty()
		err := target.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov21_Collection_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddWithWgLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.AddWithWgLock(wg, "a")
		wg.Wait()

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov21_Collection_AddPointerCollectionsLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddPointerCollectionsLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c2 := corestr.New.Collection.Strings([]string{"a"})
		c.AddPointerCollectionsLock(c2)

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov21_Collection_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddHashmapsValues", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.Cap(2)
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsValues(hm)

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov21_Collection_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddHashmapsKeys", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.Cap(2)
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsKeys(hm)

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov21_Collection_AddHashmapsKeysValues(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddHashmapsKeysValues", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.Cap(2)
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsKeysValues(hm)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AppendAnys", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AppendAnys("a", nil, 42)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_AppendNonEmptyAnys(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AppendNonEmptyAnys", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AppendNonEmptyAnys("a", nil, "")

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov21_Collection_AddsNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddsNonEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddsNonEmpty("a", "", "b")

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_New(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_New", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		newC := c.New("a", "b")

		// Act
		actual := args.Map{"result": newC.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_AddNonEmptyStrings(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddNonEmptyStrings", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStrings("a", "", "b")

		// Act
		actual := args.Map{"length": c.Length()}

		// Assert
		expected := args.Map{"length": 2}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyStrings returns 2 -- filters empty strings", actual)
	})
}

func Test_Cov21_Collection_AddFuncResult(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddFuncResult", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddFuncResult(func() string { return "a" })

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov21_Collection_AddStringsByFuncChecking(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AddStringsByFuncChecking", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddStringsByFuncChecking([]string{"a", "", "b"}, func(s string) bool { return s != "" })

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_ExpandSlicePlusAdd(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_ExpandSlicePlusAdd", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.ExpandSlicePlusAdd([]string{"a,b"}, func(line string) []string {
			return []string{line + "1", line + "2"}
		})

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_MergeSlicesOfSlice(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_MergeSlicesOfSlice", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.MergeSlicesOfSlice([]string{"a"}, []string{"b"})

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_IsEqualsWithSensitive(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_IsEqualsWithSensitive", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"Hello"})
		b := corestr.New.Collection.Strings([]string{"hello"})

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(true, b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for case sensitive", actual)
		actual = args.Map{"result": a.IsEqualsWithSensitive(false, b)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for case insensitive", actual)
	})
}

func Test_Cov21_Collection_AppendAnysUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AppendAnysUsingFilter", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AppendAnysUsingFilter(func(s string, i int) (string, bool, bool) {
			return s, s != "", false
		}, "a", nil, "b")

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_ChainRemoveAt(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_ChainRemoveAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.ChainRemoveAt(1)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_CharCollectionMap", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"abc", "bcd"})
		ccm := c.CharCollectionMap()

		// Act
		actual := args.Map{"result": ccm == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov21_Collection_SummaryString(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_SummaryString", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.SummaryString(1)

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov21_Collection_SummaryStringWithHeader(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_SummaryStringWithHeader", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		s := c.SummaryStringWithHeader("header")

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
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
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.Single() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
	})
}

func Test_Cov21_Collection_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_HasAnyItem", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.HasAnyItem()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
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
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.RemoveItemsIndexes(true, 1)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_AppendCollectionPtr(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AppendCollectionPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c.AppendCollectionPtr(c2)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_AppendCollections(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_AppendCollections", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c.AppendCollections(c1, c2)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_NonEmptyItemsOrNonWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_NonEmptyItemsOrNonWhitespace", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "", "  ", "b"})
		items := c.NonEmptyItemsOrNonWhitespace()

		// Act
		actual := args.Map{"result": len(items) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_HashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_HashsetLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs := c.HashsetLock()

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov21_Collection_FilterLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_FilterLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		result := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov21_Collection_FilterPtr(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_FilterPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		result := c.FilterPtr(func(s *string, i int) (*string, bool, bool) {
			return s, len(*s) > 1, false
		})

		// Act
		actual := args.Map{"result": len(*result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov21_Collection_FilterPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov21_Collection_FilterPtrLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		result := c.FilterPtrLock(func(s *string, i int) (*string, bool, bool) {
			return s, len(*s) > 1, false
		})

		// Act
		actual := args.Map{"result": len(*result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}
