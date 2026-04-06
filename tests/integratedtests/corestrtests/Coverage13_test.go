package corestrtests

import (
	"encoding/json"
	"errors"
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Collection — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_Collection_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_BasicOps", func() {
		c := corestr.New.Collection.Cap(10)

		c.Add("a").Add("b").Add("c")

		actual := args.Map{"result": c.Length() != 3 || c.Count() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)

		actual := args.Map{"result": c.HasAnyItem() || c.IsEmpty() || !c.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

		actual := args.Map{"result": c.LastIndex() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		actual := args.Map{"result": c.HasIndex(2) || c.HasIndex(3)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasIndex failed", actual)

		actual := args.Map{"result": c.First() != "a" || c.Last() != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "First/Last failed", actual)

		actual := args.Map{"result": c.FirstOrDefault() != "a" || c.LastOrDefault() != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault/LastOrDefault failed", actual)

		actual := args.Map{"result": c.Capacity() < 10}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected cap >= 10", actual)
	})
}

func Test_Cov13_Collection_EmptyDefaults(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_EmptyDefaults", func() {
		c := corestr.Empty.Collection()

		actual := args.Map{"result": c.FirstOrDefault() != "" || c.LastOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)

		actual := args.Map{"result": c.IsEmpty() || c.HasItems() || c.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
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

		actual := args.Map{"result": c.Length() != 7}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 7", actual)
	})
}

func Test_Cov13_Collection_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddFuncErr", func() {
		c := corestr.New.Collection.Cap(5)

		// No error
		c.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)

		// With error
		c.AddFuncErr(func() (string, error) { return "", errors.New("fail") }, func(e error) {})
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 1", actual)
	})
}

func Test_Cov13_Collection_Adds(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Adds", func() {
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")
		c.AddStrings([]string{"d", "e"})

		actual := args.Map{"result": c.Length() != 5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5", actual)
	})
}

func Test_Cov13_Collection_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddCollection", func() {
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"c", "d"})

		c1.AddCollection(c2)
		actual := args.Map{"result": c1.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)

		// Empty collection
		c1.AddCollection(corestr.Empty.Collection())
		actual := args.Map{"result": c1.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 4", actual)
	})
}

func Test_Cov13_Collection_AddCollections(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddCollections", func() {
		c := corestr.New.Collection.Cap(10)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})

		c.AddCollections(c1, c2, corestr.Empty.Collection())
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov13_Collection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_ConcatNew", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.ConcatNew(5, "c", "d")

		actual := args.Map{"result": result.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)

		// Empty additionalStrings
		result2 := c.ConcatNew(0)
		actual := args.Map{"result": result2.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov13_Collection_AsError(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AsError", func() {
		c := corestr.New.Collection.Strings([]string{"err1", "err2"})
		err := c.AsError(",")

		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)

		empty := corestr.Empty.Collection()
		actual := args.Map{"result": empty.AsError(",") != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)

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

		actual := args.Map{"result": c.RemoveAt(1)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected success", actual)

		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		// Out of range
		actual := args.Map{"result": c.RemoveAt(-1) || c.RemoveAt(10)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected failure", actual)
	})
}

func Test_Cov13_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_EachItemSplitBy", func() {
		c := corestr.New.Collection.Strings([]string{"a.b", "c.d"})
		result := c.EachItemSplitBy(".")

		actual := args.Map{"result": len(result) != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_Cov13_Collection_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_IsEquals", func() {
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"a", "b"})
		c3 := corestr.New.Collection.Strings([]string{"a", "B"})

		actual := args.Map{"result": c1.IsEquals(c2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)

		actual := args.Map{"result": c1.IsEquals(c3)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal case-sensitive", actual)

		actual := args.Map{"result": c1.IsEqualsWithSensitive(false, c3)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal case-insensitive", actual)

		// Same ptr
		actual := args.Map{"result": c1.IsEquals(c1)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected same ptr equal", actual)

		// Both empty
		e1 := corestr.Empty.Collection()
		e2 := corestr.Empty.Collection()
		actual := args.Map{"result": e1.IsEquals(e2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty equals", actual)

		// Different length
		c4 := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c1.IsEquals(c4)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal different length", actual)
	})
}

func Test_Cov13_Collection_Take_Skip(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Take_Skip", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})

		taken := c.Take(2)
		actual := args.Map{"result": taken.Length() != 2 || taken.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Take failed", actual)

		// Take more than length
		taken2 := c.Take(10)
		actual := args.Map{"result": taken2.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected full", actual)

		// Take 0
		taken3 := c.Take(0)
		actual := args.Map{"result": taken3.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)

		// Skip
		skipped := c.Skip(2)
		actual := args.Map{"result": skipped.Length() != 2 || skipped.First() != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Skip failed", actual)

		// Skip 0
		actual := args.Map{"result": c.Skip(0) != c}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Skip 0 should return self", actual)
	})
}

func Test_Cov13_Collection_Reverse(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Reverse", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.Reverse()

		actual := args.Map{"result": c.First() != "c" || c.Last() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Reverse failed", actual)

		// 2 items
		c2 := corestr.New.Collection.Strings([]string{"x", "y"})
		c2.Reverse()

		actual := args.Map{"result": c2.First() != "y"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected y first", actual)

		// 1 item
		c3 := corestr.New.Collection.Strings([]string{"z"})
		c3.Reverse()

		actual := args.Map{"result": c3.First() != "z"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected z", actual)
	})
}

func Test_Cov13_Collection_GetPagesSize(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_GetPagesSize", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})

		actual := args.Map{"result": c.GetPagesSize(2) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)

		actual := args.Map{"result": c.GetPagesSize(0) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for 0 page size", actual)
	})
}

func Test_Cov13_Collection_Has(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Has", func() {
		c := corestr.New.Collection.Strings([]string{"hello", "world"})

		actual := args.Map{"result": c.Has("hello") || c.Has("missing")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Has failed", actual)

		str := "hello"
		actual := args.Map{"result": c.HasPtr(&str)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasPtr failed", actual)

		actual := args.Map{"result": c.HasAll("hello", "world")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAll failed", actual)

		actual := args.Map{"result": c.HasAll("hello", "missing")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasAll should fail", actual)

		actual := args.Map{"result": c.HasUsingSensitivity("HELLO", false)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected case-insensitive match", actual)

		actual := args.Map{"result": c.HasUsingSensitivity("HELLO", true)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected case-sensitive mismatch", actual)
	})
}

func Test_Cov13_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Filter", func() {
		c := corestr.New.Collection.Strings([]string{"ab", "cd", "ef"})

		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, s == "ab" || s == "ef", false
		})

		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov13_Collection_SortedListAsc_Dsc(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_SortedListAsc_Dsc", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})

		asc := c.SortedListAsc()
		actual := args.Map{"result": asc[0] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a first", actual)

		dsc := c.SortedListDsc()
		actual := args.Map{"result": dsc[0] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c first", actual)
	})
}

func Test_Cov13_Collection_UniqueList(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_UniqueList", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "a", "c"})
		uniques := c.UniqueList()

		actual := args.Map{"result": len(uniques) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3 unique", actual)
	})
}

func Test_Cov13_Collection_HashsetAsIs(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_HashsetAsIs", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		h := c.HashsetAsIs()

		actual := args.Map{"result": h.Has("a") || !h.Has("b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected both", actual)
	})
}

func Test_Cov13_Collection_NonEmptyList(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_NonEmptyList", func() {
		c := corestr.New.Collection.Strings([]string{"", "a", "", "b"})
		result := c.NonEmptyList()

		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov13_Collection_IsContainsAll(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_IsContainsAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		actual := args.Map{"result": c.IsContainsAll("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)

		actual := args.Map{"result": c.IsContainsAll("a", "z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)

		actual := args.Map{"result": c.IsContainsAll()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_Cov13_Collection_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_GetAllExcept", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.GetAllExcept([]string{"b"})

		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		// nil items
		result2 := c.GetAllExcept(nil)
		actual := args.Map{"result": len(result2) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3 copy", actual)
	})
}

func Test_Cov13_Collection_Join(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Join", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		actual := args.Map{"result": c.Join(",") != "a,b,c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Join failed", actual)

		actual := args.Map{"result": c.JoinLine() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "JoinLine failed", actual)

		empty := corestr.Empty.Collection()
		actual := args.Map{"result": empty.Join(",") != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov13_Collection_String(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_String", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.String()

		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "marshal failed", actual)

		c2 := corestr.Empty.Collection()
		err = json.Unmarshal(b, c2)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unmarshal failed", actual)

		actual := args.Map{"result": c2.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov13_Collection_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Serialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_, err := c.Serialize()
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "serialize failed", actual)
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

		actual := args.Map{"result": c.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 after clear", actual)

		c2 := corestr.New.Collection.Strings([]string{"x"})
		c2.Dispose()
	})
}

func Test_Cov13_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AppendAnys", func() {
		c := corestr.New.Collection.Cap(10)
		c.AppendAnys(42, "hello", nil)

		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 (nil skipped)", actual)
	})
}

func Test_Cov13_Collection_AppendNonEmptyAnys(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AppendNonEmptyAnys", func() {
		c := corestr.New.Collection.Cap(10)
		c.AppendNonEmptyAnys(42, nil)

		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov13_Collection_AddsNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddsNonEmpty", func() {
		c := corestr.New.Collection.Cap(10)
		c.AddsNonEmpty("a", "", "b")

		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov13_Collection_Joins(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_Joins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// With extra items
		r := c.Joins(",", "c", "d")
		actual := args.Map{"result": r == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

		// Without extra items
		r2 := c.Joins(",")
		actual := args.Map{"result": r2 != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
	})
}

func Test_Cov13_Collection_New(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_New", func() {
		c := corestr.New.Collection.Cap(5)
		result := c.New("a", "b")

		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov13_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_IndexAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		actual := args.Map{"result": c.IndexAt(1) != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_Cov13_Collection_SafeIndexAtUsingLength(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_SafeIndexAtUsingLength", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		actual := args.Map{"result": c.SafeIndexAtUsingLength("def", 2, 5) != "def"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default", actual)

		actual := args.Map{"result": c.SafeIndexAtUsingLength("def", 2, 1) != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
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

		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov13_Collection_AppendCollectionPtr(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AppendCollectionPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b", "c"})
		c.AppendCollectionPtr(c2)

		actual := args.Map{"result": c.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov13_Collection_AddFuncResult(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddFuncResult", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddFuncResult(func() string { return "hello" })

		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov13_Collection_AddStringsByFuncChecking(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddStringsByFuncChecking", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddStringsByFuncChecking([]string{"a", "bb", "c"}, func(s string) bool {
			return len(s) == 1
		})

		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov13_Collection_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_GetAllExceptCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		except := corestr.New.Collection.Strings([]string{"b"})
		result := c.GetAllExceptCollection(except)

		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		// nil
		result2 := c.GetAllExceptCollection(nil)
		actual := args.Map{"result": len(result2) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov13_Collection_SummaryString(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_SummaryString", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.SummaryString(1)

		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov13_Collection_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddHashmapsValues", func() {
		c := corestr.New.Collection.Cap(10)
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("k1", "v1")
		h.AddOrUpdate("k2", "v2")

		c.AddHashmapsValues(h)
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		c.AddHashmapsValues(nil)
	})
}

func Test_Cov13_Collection_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddHashmapsKeys", func() {
		c := corestr.New.Collection.Cap(10)
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("k1", "v1")

		c.AddHashmapsKeys(h)
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov13_Collection_AddHashmapsKeysValues(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_AddHashmapsKeysValues", func() {
		c := corestr.New.Collection.Cap(10)
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("k", "v")

		c.AddHashmapsKeysValues(h)
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov13_Collection_GetHashsetPlusHasAll(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_GetHashsetPlusHasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		h, ok := c.GetHashsetPlusHasAll([]string{"a", "b"})

		actual := args.Map{"result": ok || h == nil}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)

		_, ok2 := c.GetHashsetPlusHasAll(nil)
		actual := args.Map{"result": ok2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
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

		actual := args.Map{"result": paged.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3 pages", actual)
	})
}

func Test_Cov13_Collection_GetSinglePageCollection(t *testing.T) {
	safeTest(t, "Test_Cov13_Collection_GetSinglePageCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})

		page := c.GetSinglePageCollection(2, 1)
		actual := args.Map{"result": page.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		page2 := c.GetSinglePageCollection(2, 3)
		actual := args.Map{"result": page2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)

		// When length < eachPageSize
		small := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": small.GetSinglePageCollection(10, 1) != small}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected self", actual)
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

		actual := args.Map{"result": s.Length() != 3 || s.Count() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)

		actual := args.Map{"result": s.IsEmpty() || !s.HasAnyItem()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

		actual := args.Map{"result": s.LastIndex() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		actual := args.Map{"result": s.HasIndex(2) || s.HasIndex(3)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasIndex failed", actual)

		actual := args.Map{"result": s.First() != "a" || s.Last() != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "First/Last failed", actual)

		actual := args.Map{"result": s.FirstOrDefault() != "a" || s.LastOrDefault() != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault/LastOrDefault failed", actual)
	})
}

func Test_Cov13_SimpleSlice_EmptyDefaults(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_EmptyDefaults", func() {
		s := corestr.New.SimpleSlice.Empty()

		actual := args.Map{"result": s.FirstOrDefault() != "" || s.LastOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
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

		actual := args.Map{"result": s.Length() != 9}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 9", actual)
	})
}

func Test_Cov13_SimpleSlice_AddStruct(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_AddStruct", func() {
		s := corestr.New.SimpleSlice.Cap(5)
		s.AddStruct(true, struct{ Name string }{"hello"})
		s.AddStruct(true, nil)

		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov13_SimpleSlice_AddPointer(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_AddPointer", func() {
		s := corestr.New.SimpleSlice.Cap(5)
		val := "test"
		s.AddPointer(true, &val)
		s.AddPointer(true, nil)

		actual := args.Map{"result": s.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
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

		actual := args.Map{"result": s.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
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

		actual := args.Map{"result": s.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_Cov13_SimpleSlice_InsertAt(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_InsertAt", func() {
		s := corestr.New.SimpleSlice.Lines("a", "c")
		s.InsertAt(1, "b")

		actual := args.Map{"result": s.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)

		// Out of range
		s.InsertAt(-1, "x")
		s.InsertAt(100, "x")
	})
}

func Test_Cov13_SimpleSlice_Skip_Take(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_Skip_Take", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "c", "d")

		skipped := s.Skip(2)
		actual := args.Map{"result": len(skipped) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		taken := s.Take(2)
		actual := args.Map{"result": len(taken) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		// Skip more than length
		skippedAll := s.Skip(10)
		actual := args.Map{"result": len(skippedAll) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)

		// Take more
		takenAll := s.Take(10)
		actual := args.Map{"result": len(takenAll) != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)

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

		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)

		_ = s.AsDefaultError()

		empty := corestr.New.SimpleSlice.Empty()
		actual := args.Map{"result": empty.AsError(",") != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Cov13_SimpleSlice_Join(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_Join", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")

		actual := args.Map{"result": s.Join(",") != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)

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

		actual := args.Map{"result": s.IsContains("hello") || s.IsContains("missing")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsContains failed", actual)

		actual := args.Map{"result": s.IndexOf("world") != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IndexOf failed", actual)

		actual := args.Map{"result": s.IndexOf("missing") != -1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
	})
}

func Test_Cov13_SimpleSlice_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_IsContainsFunc", func() {
		s := corestr.New.SimpleSlice.Lines("hello", "world")

		found := s.IsContainsFunc("hello", func(item, searching string) bool {
			return item == searching
		})

		actual := args.Map{"result": found}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
	})
}

func Test_Cov13_SimpleSlice_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_IndexOfFunc", func() {
		s := corestr.New.SimpleSlice.Lines("hello", "world")

		idx := s.IndexOfFunc("world", func(item, searching string) bool {
			return item == searching
		})

		actual := args.Map{"result": idx != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov13_SimpleSlice_CountFunc(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_CountFunc", func() {
		s := corestr.New.SimpleSlice.Lines("a", "bb", "c")

		count := s.CountFunc(func(i int, item string) bool {
			return len(item) == 1
		})

		actual := args.Map{"result": count != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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

		actual := args.Map{"result": (*result)[0] != "a!"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a!", actual)

		// TranspileJoin
		joined := s.TranspileJoin(func(s string) string { return s }, ",")
		actual := args.Map{"result": joined != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
	})
}

func Test_Cov13_SimpleSlice_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_EachItemSplitBy", func() {
		s := corestr.New.SimpleSlice.Lines("a.b", "c.d")
		result := s.EachItemSplitBy(".")

		actual := args.Map{"result": result.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_Cov13_SimpleSlice_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_IsEqual", func() {
		s1 := corestr.New.SimpleSlice.Lines("a", "b")
		s2 := corestr.New.SimpleSlice.Lines("a", "b")
		s3 := corestr.New.SimpleSlice.Lines("a", "c")

		actual := args.Map{"result": s1.IsEqual(s2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)

		actual := args.Map{"result": s1.IsEqual(s3)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)

		actual := args.Map{"result": s1.IsEqualLines([]string{"a", "b"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov13_SimpleSlice_IsUnorderedEqual(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_IsUnorderedEqual", func() {
		s1 := corestr.New.SimpleSlice.Lines("b", "a")
		s2 := corestr.New.SimpleSlice.Lines("a", "b")

		actual := args.Map{"result": s1.IsUnorderedEqual(true, s2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov13_SimpleSlice_IsDistinctEqual(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_IsDistinctEqual", func() {
		s1 := corestr.New.SimpleSlice.Lines("a", "b", "a")
		s2 := corestr.New.SimpleSlice.Lines("b", "a")

		actual := args.Map{"result": s1.IsDistinctEqual(s2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov13_SimpleSlice_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_ConcatNew", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.ConcatNew("c", "d")

		actual := args.Map{"result": result.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)

		_ = s.ConcatNewStrings("c")
	})
}

func Test_Cov13_SimpleSlice_ConcatNewSimpleSlices(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_ConcatNewSimpleSlices", func() {
		s1 := corestr.New.SimpleSlice.Lines("a")
		s2 := corestr.New.SimpleSlice.Lines("b")
		result := s1.ConcatNewSimpleSlices(s2)

		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov13_SimpleSlice_PrependAppend(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_PrependAppend", func() {
		s := corestr.New.SimpleSlice.Lines("b")
		s.PrependAppend([]string{"a"}, []string{"c"})

		actual := args.Map{"result": s.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov13_SimpleSlice_PrependJoin(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_PrependJoin", func() {
		s := corestr.New.SimpleSlice.Lines("b", "c")
		result := s.PrependJoin(",", "a")

		actual := args.Map{"result": result != "a,b,c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b,c", actual)
	})
}

func Test_Cov13_SimpleSlice_AppendJoin(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_AppendJoin", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.AppendJoin(",", "c")

		actual := args.Map{"result": result != "a,b,c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b,c", actual)
	})
}

func Test_Cov13_SimpleSlice_Sort_Reverse(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_Sort_Reverse", func() {
		s := corestr.New.SimpleSlice.Lines("c", "a", "b")
		s.Sort()

		actual := args.Map{"result": s.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a first", actual)

		s.Reverse()
		actual := args.Map{"result": s.First() != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c first", actual)
	})
}

func Test_Cov13_SimpleSlice_Clone(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_Clone", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		cloned := s.Clone(true)

		actual := args.Map{"result": cloned.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		_ = s.DeepClone()
		_ = s.ShadowClone()
		_ = s.ClonePtr(true)
	})
}

func Test_Cov13_SimpleSlice_RemoveIndexes(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_RemoveIndexes", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result, err := s.RemoveIndexes(1)

		actual := args.Map{"result": err != nil || result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		// Empty slice
		empty := corestr.New.SimpleSlice.Empty()
		_, err = empty.RemoveIndexes(0)
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_Cov13_SimpleSlice_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_DistinctDiff", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.DistinctDiffRaw("b", "c")

		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov13_SimpleSlice_AddedRemovedLinesDiff(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_AddedRemovedLinesDiff", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		added, removed := s.AddedRemovedLinesDiff("b", "c")

		actual := args.Map{"result": len(added) != 1 || len(removed) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 added, 1 removed", actual)
	})
}

func Test_Cov13_SimpleSlice_JSON(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_JSON", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")

		b, err := json.Marshal(s)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "marshal failed", actual)

		s2 := corestr.New.SimpleSlice.Empty()
		err = json.Unmarshal(b, s2)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unmarshal failed", actual)
	})
}

func Test_Cov13_SimpleSlice_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_Clear_Dispose", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		s.Clear()

		actual := args.Map{"result": s.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)

		s2 := corestr.New.SimpleSlice.Lines("x")
		s2.Dispose()
	})
}

func Test_Cov13_SimpleSlice_Collection(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_Collection", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		c := s.Collection(false)

		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov13_SimpleSlice_Hashset(t *testing.T) {
	safeTest(t, "Test_Cov13_SimpleSlice_Hashset", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		h := s.Hashset()

		actual := args.Map{"result": h.Has("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
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

		actual := args.Map{"result": ok}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
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

func Test_Cov13_ValidValue_AllMethods(t *testing.T) {
	safeTest(t, "Test_Cov13_ValidValue_AllMethods", func() {
		v := corestr.NewValidValue("42")

		actual := args.Map{"result": v.IsEmpty() || !v.IsValid || !v.HasValidNonEmpty() || !v.HasSafeNonEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid non-empty", actual)

		actual := args.Map{"result": v.IsWhitespace() || !v.HasValidNonWhitespace()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-whitespace", actual)

		_ = v.ValueBytesOnce()
		_ = v.ValueBytesOncePtr()
		_ = v.Trim()

		actual := args.Map{"result": v.ValueInt(0) != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)

		actual := args.Map{"result": v.ValueDefInt() != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)

		_ = v.ValueByte(0)
		_ = v.ValueDefByte()

		fv := corestr.NewValidValue("3.14")
		_ = fv.ValueFloat64(0)
		_ = fv.ValueDefFloat64()

		bv := corestr.NewValidValue("true")
		actual := args.Map{"result": bv.ValueBool()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)

		actual := args.Map{"result": v.Is("42") || v.Is("43")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Is failed", actual)

		actual := args.Map{"result": v.IsAnyOf("42", "43")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)

		actual := args.Map{"result": v.IsAnyOf()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty", actual)

		actual := args.Map{"result": v.IsContains("4")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected contains", actual)

		actual := args.Map{"result": v.IsAnyContains("4", "x")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)

		actual := args.Map{"result": v.IsEqualNonSensitive("42")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)

		re := regexp.MustCompile(`\d+`)
		actual := args.Map{"result": v.IsRegexMatches(re)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected match", actual)
		actual := args.Map{"result": v.IsRegexMatches(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)

		_ = v.RegexFindString(re)
		_ = v.RegexFindString(nil)
		_, _ = v.RegexFindAllStringsWithFlag(re, -1)
		_ = v.RegexFindAllStrings(re, -1)

		_ = v.Split(",")
		_ = v.SplitNonEmpty(",")
		_ = v.SplitTrimNonWhitespace(",")

		_ = v.Clone()
		_ = v.String()
		_ = v.FullString()

		v.Clear()
		v.Dispose()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValues — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_ValidValues_AllMethods(t *testing.T) {
	safeTest(t, "Test_Cov13_ValidValues_AllMethods", func() {
		vv := corestr.NewValidValues(5)
		vv.Add("a")
		vv.AddFull(true, "b", "msg")

		actual := args.Map{"result": vv.Length() != 2 || vv.Count() != 2 || vv.IsEmpty() || !vv.HasAnyItem()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		actual := args.Map{"result": vv.LastIndex() != 1 || !vv.HasIndex(1) || vv.HasIndex(5)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "index check failed", actual)

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
		actual := args.Map{"result": len(found) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov13_ValidValues_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov13_ValidValues_ConcatNew", func() {
		vv := corestr.NewValidValues(2)
		vv.Add("a")

		vv2 := corestr.NewValidValues(2)
		vv2.Add("b")

		result := vv.ConcatNew(true, vv2)
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		// Empty concat
		result2 := vv.ConcatNew(true)
		actual := args.Map{"result": result2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)

		result3 := vv.ConcatNew(false)
		actual := args.Map{"result": result3 != vv}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same ptr", actual)
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
		actual := args.Map{"result": vs.ValueValid.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)

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

		actual := args.Map{"result": kv.KeyName() != "name" || kv.VariableName() != "name"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "key failed", actual)

		actual := args.Map{"result": kv.ValueString() != "42"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)

		actual := args.Map{"result": kv.IsVariableNameEqual("name") || !kv.IsValueEqual("42")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "equality failed", actual)

		actual := args.Map{"result": kv.IsKeyEmpty() || kv.IsValueEmpty() || kv.IsKeyValueEmpty() || kv.IsKeyValueAnyEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

		actual := args.Map{"result": kv.HasKey() || !kv.HasValue()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has", actual)

		_ = kv.TrimKey()
		_ = kv.TrimValue()
		_ = kv.String()
		_ = kv.Compile()
		_ = kv.FormatString("%s=%s")

		actual := args.Map{"result": kv.ValueInt(0) != 42 || kv.ValueDefInt() != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)

		_ = kv.ValueByte(0)
		_ = kv.ValueDefByte()

		kvBool := corestr.KeyValuePair{Key: "k", Value: "true"}
		actual := args.Map{"result": kvBool.ValueBool()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)

		kvFloat := corestr.KeyValuePair{Key: "k", Value: "3.14"}
		_ = kvFloat.ValueFloat64(0)
		_ = kvFloat.ValueDefFloat64()

		_ = kv.ValueValid()
		_ = kv.ValueValidOptions(true, "msg")

		actual := args.Map{"result": kv.Is("name", "42") || !kv.IsKey("name") || !kv.IsVal("42")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Is failed", actual)

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

		actual := args.Map{"result": kav.KeyName() != "test" || kav.VariableName() != "test"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "key failed", actual)

		actual := args.Map{"result": kav.ValueAny() != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)

		actual := args.Map{"result": kav.IsVariableNameEqual("test")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)

		_ = kav.ValueString()
		_ = kav.String()
		_ = kav.Compile()

		actual := args.Map{"result": kav.IsValueNull() || !kav.HasNonNull() || !kav.HasValue()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-null", actual)

		actual := args.Map{"result": kav.IsValueEmptyString() || kav.IsValueWhitespace()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty string", actual)

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

		actual := args.Map{"result": tl.HasLineNumber() || tl.IsInvalidLineNumber()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected valid line number", actual)

		actual := args.Map{"result": tl.Length() != 5 || tl.IsEmpty() || tl.IsEmptyText()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

		// Empty
		empty := &corestr.TextWithLineNumber{}
		actual := args.Map{"result": empty.IsEmptyTextLineBoth()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CloneSlice / CloneSliceIf — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_CloneSlice(t *testing.T) {
	safeTest(t, "Test_Cov13_CloneSlice", func() {
		result := corestr.CloneSlice([]string{"a", "b"})

		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		empty := corestr.CloneSlice([]string{})
		actual := args.Map{"result": len(empty) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov13_CloneSliceIf(t *testing.T) {
	safeTest(t, "Test_Cov13_CloneSliceIf", func() {
		result := corestr.CloneSliceIf(true, "a", "b")
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		// No clone
		result2 := corestr.CloneSliceIf(false, "a", "b")
		actual := args.Map{"result": len(result2) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		// Empty
		result3 := corestr.CloneSliceIf(true)
		actual := args.Map{"result": len(result3) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// utils — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_Utils_WrapMethods(t *testing.T) {
	safeTest(t, "Test_Cov13_Utils_WrapMethods", func() {
		u := corestr.StringUtils

		actual := args.Map{"result": u.WrapDouble("a") != `"a"`}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "WrapDouble failed", actual)

		actual := args.Map{"result": u.WrapSingle("a") != "'a'"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "WrapSingle failed", actual)

		actual := args.Map{"result": u.WrapTilda("a") != "`a`"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "WrapTilda failed", actual)

		actual := args.Map{"result": u.WrapDoubleIfMissing(`"a"`) != `"a"`}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "already wrapped", actual)

		actual := args.Map{"result": u.WrapDoubleIfMissing("a") != `"a"`}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrap missing", actual)

		actual := args.Map{"result": u.WrapDoubleIfMissing("") != `""`}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty wrap", actual)

		actual := args.Map{"result": u.WrapSingleIfMissing("'a'") != "'a'"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "already wrapped", actual)

		actual := args.Map{"result": u.WrapSingleIfMissing("a") != "'a'"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrap missing", actual)

		actual := args.Map{"result": u.WrapSingleIfMissing("") != "''"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty wrap", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyToString / reflectInterfaceVal — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_AnyToString(t *testing.T) {
	safeTest(t, "Test_Cov13_AnyToString", func() {
		r := corestr.AnyToString(true, 42)
		actual := args.Map{"result": r == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

		r2 := corestr.AnyToString(false, "hello")
		actual := args.Map{"result": r2 == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

		// Empty
		r3 := corestr.AnyToString(true, "")
		actual := args.Map{"result": r3 != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
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
		actual := args.Map{"result": length != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)

		actual := args.Map{"result": corestr.AllIndividualsLengthOfSimpleSlices() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_AllIndividualStringsOfStringsLength(t *testing.T) {
	safeTest(t, "Test_Cov13_AllIndividualStringsOfStringsLength", func() {
		items := [][]string{{"a", "b"}, {"c"}}
		length := corestr.AllIndividualStringsOfStringsLength(&items)

		actual := args.Map{"result": length != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)

		actual := args.Map{"result": corestr.AllIndividualStringsOfStringsLength(nil) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_LeftRight_AllMethods(t *testing.T) {
	safeTest(t, "Test_Cov13_LeftRight_AllMethods", func() {
		lr := corestr.NewLeftRight("left", "right")

		actual := args.Map{"result": lr.Left != "left" || lr.Right != "right"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected left/right", actual)

		_ = lr.LeftBytes()
		_ = lr.RightBytes()
		_ = lr.LeftTrim()
		_ = lr.RightTrim()

		actual := args.Map{"result": lr.IsLeftEmpty() || lr.IsRightEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

		actual := args.Map{"result": lr.IsLeftWhitespace() || lr.IsRightWhitespace()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-whitespace", actual)

		actual := args.Map{"result": lr.HasValidNonEmptyLeft() || !lr.HasValidNonEmptyRight()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected valid non-empty", actual)

		actual := args.Map{"result": lr.HasValidNonWhitespaceLeft() || !lr.HasValidNonWhitespaceRight()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected valid non-whitespace", actual)

		actual := args.Map{"result": lr.HasSafeNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected safe non-empty", actual)

		actual := args.Map{"result": lr.IsLeft("left") || !lr.IsRight("right") || !lr.Is("left", "right")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Is failed", actual)

		lr2 := corestr.NewLeftRight("left", "right")
		actual := args.Map{"result": lr.IsEqual(lr2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)

		_ = lr.Clone()
		_ = lr.NonPtr()
		_ = lr.Ptr()

		re := regexp.MustCompile("left")
		actual := args.Map{"result": lr.IsLeftRegexMatch(re)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected match", actual)
		actual := args.Map{"result": lr.IsLeftRegexMatch(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)

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
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a=b", actual)

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

		actual := args.Map{"result": lmr.IsLeftEmpty() || lmr.IsRightEmpty() || lmr.IsMiddleEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

		actual := args.Map{"result": lmr.IsLeftWhitespace() || lmr.IsRightWhitespace() || lmr.IsMiddleWhitespace()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-whitespace", actual)

		actual := args.Map{"result": lmr.HasValidNonEmptyLeft() || !lmr.HasValidNonEmptyRight() || !lmr.HasValidNonEmptyMiddle()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected valid", actual)

		actual := args.Map{"result": lmr.HasValidNonWhitespaceLeft() || !lmr.HasValidNonWhitespaceRight() || !lmr.HasValidNonWhitespaceMiddle()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected valid", actual)

		actual := args.Map{"result": lmr.HasSafeNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected safe", actual)

		actual := args.Map{"result": lmr.IsAll("a", "b", "c") || !lmr.Is("a", "c")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "Is failed", actual)

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
		actual := args.Map{"result": lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "split failed", actual)

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
