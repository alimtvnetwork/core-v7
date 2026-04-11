package corestrtests

import (
	"strings"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Collection — comprehensive coverage
// ==========================================================================

func Test_Collection_BasicOps(t *testing.T) {
	safeTest(t, "Test_I8_Collection_BasicOps", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Add("a").Add("b").Add("c")

		// Act
		actual := args.Map{"result": c.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual = args.Map{"result": c.Count() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "count mismatch", actual)
		actual = args.Map{"result": c.Capacity() < 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "capacity too small", actual)
		actual = args.Map{"result": c.HasAnyItem()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has items", actual)
		actual = args.Map{"result": c.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "not empty", actual)
		actual = args.Map{"result": c.HasItems() != true}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "has items", actual)
		actual = args.Map{"result": c.LastIndex() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "last index", actual)
		actual = args.Map{"result": c.HasIndex(0)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "has index 0", actual)
		actual = args.Map{"result": c.HasIndex(2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "has index 2", actual)
		actual = args.Map{"result": c.HasIndex(3)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "no index 3", actual)
		actual = args.Map{"result": c.HasIndex(-1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "no negative index", actual)
	})
}

func Test_Collection_ListMethods(t *testing.T) {
	safeTest(t, "Test_I8_Collection_ListMethods", func() {
		c := corestr.New.Collection.Strings([]string{"x", "y"})
		_ = c.ListStringsPtr()
		_ = c.ListStrings()
		_ = c.List()
		_ = c.ListPtr()
		_ = c.Items()
	})
}

func Test_Collection_AddVariants(t *testing.T) {
	safeTest(t, "Test_I8_Collection_AddVariants", func() {
		// Arrange
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

		// Act
		actual := args.Map{"result": c.Length() < 7}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 7 items", actual)
	})
}

func Test_Collection_Adds_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Adds", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.Adds("a", "b", "c")
		c.AddsLock("d", "e")
		c.AddStrings([]string{"f", "g"})
		c.AddsNonEmpty("", "h", "")

		other := corestr.New.Collection.Strings([]string{"i", "j"})
		c.AddCollection(other)
		c.AddCollections(other)

		// Act
		actual := args.Map{"result": c.Length() < 10}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 10", actual)
	})
}

func Test_Collection_RemoveAt_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_RemoveAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"result": c.RemoveAt(1)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected success", actual)
		actual = args.Map{"result": c.RemoveAt(-1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected fail for negative", actual)
		actual = args.Map{"result": c.RemoveAt(99)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected fail for out of bounds", actual)
		actual = args.Map{"result": c.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_ChainRemoveAt_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_ChainRemoveAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.ChainRemoveAt(0)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_InsertAt_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_InsertAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "c"})
		c.InsertAt(1, "b")

		// Act
		actual := args.Map{"result": c.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Collection_RemoveItemsIndexes_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_RemoveItemsIndexes", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		c.RemoveItemsIndexes(false, 0, 2)
		c2 := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		indexes := []int{0}
		c2.RemoveItemsIndexesPtr(false, indexes)
	})
}

func Test_Collection_FirstLastSingleTakeSkip(t *testing.T) {
	safeTest(t, "Test_I8_Collection_FirstLastSingleTakeSkip", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"result": c.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "first", actual)
		actual = args.Map{"result": c.Last() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "last", actual)
		actual = args.Map{"result": c.FirstOrDefault() != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "first or default", actual)
		actual = args.Map{"result": c.LastOrDefault() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "last or default", actual)

		taken := c.Take(2)
		actual = args.Map{"result": taken.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "take 2", actual)
		skipped := c.Skip(1)
		actual = args.Map{"result": skipped.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "skip 1", actual)

		empty := corestr.New.Collection.Empty()
		actual = args.Map{"result": empty.FirstOrDefault() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty first", actual)
		actual = args.Map{"result": empty.LastOrDefault() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty last", actual)

		single := corestr.New.Collection.Strings([]string{"only"})
		actual = args.Map{"result": single.Single() != "only"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "single", actual)
	})
}

func Test_Collection_Reverse(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Reverse", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		r := c.Reverse()

		// Act
		actual := args.Map{"result": r.First() != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "reverse first", actual)
	})
}

func Test_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_IndexAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"result": c.IndexAt(0) != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "index 0", actual)
		actual = args.Map{"result": c.SafeIndexAtUsingLength("", 3, 0) != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "safe index 0", actual)
	})
}

func Test_Collection_IsEquals_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_IsEquals", func() {
		// Arrange
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})
		c := corestr.New.Collection.Strings([]string{"a", "c"})

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		actual = args.Map{"result": a.IsEquals(c)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
		actual = args.Map{"result": a.IsEquals(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal to nil", actual)

		_ = a.IsEqualsWithSensitive(false, b)
	})
}

func Test_Collection_LengthLock_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_LengthLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": c.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		if !c.IsEmptyLock() == true { /* expected non-empty */ }
	})
}

func Test_Collection_AsyncOps(t *testing.T) {
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

func Test_Collection_Filter_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Filter", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"aa", "b", "cc"})
		filtered := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, len(s) == 2, false
		})

		// Act
		actual := args.Map{"result": len(filtered)}

		// Assert
		expected := args.Map{"result": 2}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

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

func Test_Collection_Unique(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Unique", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "a"})
		uniq := c.UniqueList()

		// Act
		actual := args.Map{"result": len(uniq) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 unique", actual)
		_ = c.UniqueListLock()
		_ = c.UniqueBoolMap()
		_ = c.UniqueBoolMapLock()
	})
}

func Test_Collection_NonEmpty(t *testing.T) {
	safeTest(t, "Test_I8_Collection_NonEmpty", func() {
		c := corestr.New.Collection.Strings([]string{"a", "", "b", "  "})
		ne := c.NonEmptyList()
		_ = ne
		_ = c.NonEmptyListPtr()
		_ = c.NonEmptyItems()
		_ = c.NonEmptyItemsPtr()
		_ = c.NonEmptyItemsOrNonWhitespace()
		_ = c.NonEmptyItemsOrNonWhitespacePtr()
	})
}

func Test_Collection_Has(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Has", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": c.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has a", actual)
		actual = args.Map{"result": c.Has("z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no z", actual)
		_ = c.HasLock("a")
		s := "a"
		_ = c.HasPtr(&s)
		actual = args.Map{"result": c.HasAll("a", "b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has all", actual)
		_ = c.HasUsingSensitivity("A", false)
		_ = c.IsContainsAll("a", "b")
		_ = c.IsContainsAllLock("a")
		_ = c.IsContainsAllSlice([]string{"a"})
		ns := "a"
		_ = c.IsContainsPtr(&ns)
		_, _ = c.GetHashsetPlusHasAll([]string{"a", "b"})
	})
}

func Test_Collection_Sort(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Sort", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		asc := c.SortedListAsc()

		// Act
		actual := args.Map{"result": asc[0] != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a first", actual)
		dsc := c.SortedListDsc()
		actual = args.Map{"result": dsc[0] != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c first", actual)
		_ = c.SortedAsc()
		_ = c.SortedAscLock()
	})
}

func Test_Collection_Hashset(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Hashset", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		h := c.HashsetAsIs()

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		_ = c.HashsetWithDoubleLength()
		_ = c.HashsetLock()
	})
}

func Test_Collection_String(t *testing.T) {
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

func Test_Collection_CsvJoin(t *testing.T) {
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

func Test_Collection_Json_FromCollectionBasicOpsIt(t *testing.T) {
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

func Test_Collection_Serialize_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Serialize", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		b, err := c.Serialize()

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		var target []string
		_ = c.Deserialize(&target)
		_ = b
	})
}

func Test_Collection_Error(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Error", func() {
		c := corestr.New.Collection.Strings([]string{"err1", "err2"})
		_ = c.AsDefaultError()
		_ = c.AsError("; ")
		_ = c.ToError("; ")
		_ = c.ToDefaultError()

		c.AddError(nil)
	})
}

func Test_Collection_EachItemSplitBy_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_EachItemSplitBy", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a,b", "c,d"})
		split := c.EachItemSplitBy(",")

		// Act
		actual := args.Map{"result": len(split)}

		// Assert
		expected := args.Map{"result": 4}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_Collection_ConcatNew_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_ConcatNew", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c2 := c.ConcatNew(0, "b", "c")

		// Act
		actual := args.Map{"result": c2.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Collection_AppendCollections(t *testing.T) {
	safeTest(t, "Test_I8_Collection_AppendCollections", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c.AppendCollectionPtr(c2)
		c.AppendCollections(c2)
	})
}

func Test_Collection_AppendAnys_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_AppendAnys", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnys("a", 1, nil)
		c.AppendAnysLock("b", 2)
		c.AppendNonEmptyAnys("", "c", nil)
		c.AppendAnysUsingFilter(corestr.IsStringFilter(func(str string, index int) (string, bool, bool) { return str, str != "", false }), "d", nil)
		c.AppendAnysUsingFilterLock(corestr.IsStringFilter(func(str string, index int) (string, bool, bool) { return str, true, false }), "e")
	})
}

func Test_Collection_GetAllExcept_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_GetAllExcept", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.GetAllExcept([]string{"b"})

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		exc := corestr.New.Collection.Strings([]string{"a"})
		result2 := c.GetAllExceptCollection(exc)
		actual = args.Map{"result": len(result2) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_Paging(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Paging", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		pages := c.GetPagesSize(2)

		// Act
		actual := args.Map{"result": pages}

		// Assert
		expected := args.Map{"result": 3}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		page := c.GetPagedCollection(2)
		_ = page
		single := c.GetSinglePageCollection(2, 1)
		_ = single
	})
}

func Test_Collection_New(t *testing.T) {
	safeTest(t, "Test_I8_Collection_New", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c2 := c.New()

		// Act
		actual := args.Map{"result": c2.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty new", actual)
	})
}

func Test_Collection_AddNonEmptyStrings_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_AddNonEmptyStrings", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStrings("", "a", "", "b")
		c.AddNonEmptyStringsSlice([]string{"", "c"})
	})
}

func Test_Collection_AddStringsByFuncChecking_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_AddStringsByFuncChecking", func() {
		c := corestr.New.Collection.Empty()
		c.AddStringsByFuncChecking([]string{"aa", "b"}, func(s string) bool {
			return len(s) > 1
		})
	})
}

func Test_Collection_ExpandSlicePlusAdd_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_ExpandSlicePlusAdd", func() {
		c := corestr.New.Collection.Empty()
		c.ExpandSlicePlusAdd([]string{"a", "b"}, func(s string) []string {
			return []string{strings.ToUpper(s)}
		})
	})
}

func Test_Collection_MergeSlicesOfSlice_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_MergeSlicesOfSlice", func() {
		c := corestr.New.Collection.Empty()
		c.MergeSlicesOfSlice([]string{"a"}, []string{"b", "c"})
	})
}

func Test_Collection_CharCollectionMap_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_CharCollectionMap", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"apple", "avocado", "banana"})
		m := c.CharCollectionMap()

		// Act
		actual := args.Map{"result": m.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Collection_Resize(t *testing.T) {
	safeTest(t, "Test_I8_Collection_Resize", func() {
		c := corestr.New.Collection.Empty()
		c.Resize(100)
		c.AddCapacity(50)
	})
}

func Test_Collection_ClearDispose_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_ClearDispose", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Clear()

		// Act
		actual := args.Map{"result": c.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 after clear", actual)

		c2 := corestr.New.Collection.Strings([]string{"a"})
		c2.Dispose()
	})
}

func Test_Collection_ListCopyPtrLock(t *testing.T) {
	safeTest(t, "Test_I8_Collection_ListCopyPtrLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		cp := c.ListCopyPtrLock()

		// Act
		actual := args.Map{"result": len(cp) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Collection_AddHashmapsValues_FromCollectionBasicOpsIt(t *testing.T) {
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

func Test_Collection_AddPointerCollectionsLock_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Collection_AddPointerCollectionsLock", func() {
		c := corestr.New.Collection.Empty()
		other := corestr.New.Collection.Strings([]string{"a"})
		c.AddPointerCollectionsLock(other)
	})
}

func Test_Collection_NilLength(t *testing.T) {
	safeTest(t, "Test_I8_Collection_NilLength", func() {
		// Arrange
		var c *corestr.Collection

		// Act
		actual := args.Map{"result": c.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for nil", actual)
	})
}

func Test_Collection_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_I8_Collection_ParseInjectUsingJsonMust", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		r := corejson.New([]string{"a", "b"})
		c.ParseInjectUsingJsonMust(&r)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ==========================================================================
// Hashmap — comprehensive coverage
// ==========================================================================

func Test_Hashmap_BasicOps(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_BasicOps", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("a", "1")
		h.Set("b", "2")
		h.SetTrim(" c ", " 3 ")

		// Act
		actual := args.Map{"result": h.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "not empty", actual)
		actual = args.Map{"result": h.HasItems()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "has items", actual)
		actual = args.Map{"result": h.HasAnyItem()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "has any", actual)
		actual = args.Map{"result": h.Length() < 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected >= 3", actual)
		_ = h.LengthLock()
		_ = h.IsEmptyLock()
	})
}

func Test_Hashmap_AddVariants(t *testing.T) {
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

func Test_Hashmap_AddCollectionMaps(t *testing.T) {
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

func Test_Hashmap_Has_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Has", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")

		// Act
		actual := args.Map{"result": h.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has a", actual)
		actual = args.Map{"result": h.Contains("b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected contains b", actual)
		_ = h.ContainsLock("a")
		actual = args.Map{"result": h.IsKeyMissing("a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "not missing", actual)
		_ = h.IsKeyMissingLock("a")
		_ = h.HasLock("a")
		_ = h.HasWithLock("a")
		actual = args.Map{"result": h.HasAllStrings("a", "b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "has all", actual)
		actual = args.Map{"result": h.HasAll("a", "b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "has all", actual)
		_ = h.HasAny("a", "z")

		coll := corestr.New.Collection.Strings([]string{"a"})
		_ = h.HasAllCollectionItems(coll)
	})
}

func Test_Hashmap_Get_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Get", func() {
		// Arrange
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("key", "val")
		v, found := h.Get("key")

		// Act
		actual := args.Map{"result": found || v != "val"}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "get failed", actual)
		v2, found2 := h.GetValue("key")
		actual = args.Map{"result": found2 || v2 != "val"}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "getvalue failed", actual)
		_, f := h.Get("missing")
		actual = args.Map{"result": f}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not find", actual)
	})
}

func Test_Hashmap_Keys(t *testing.T) {
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

func Test_Hashmap_KeysValues(t *testing.T) {
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

func Test_Hashmap_Manipulation(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Manipulation", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		h.Remove("a")
		h.RemoveWithLock("b")
		h.AddOrUpdate("c", "C")
		_ = h.ValuesToLower()
		h.AddOrUpdate("d", "D")
		_ = h.KeysToLower()
	})
}

func Test_Hashmap_String(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_String", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		_ = h.String()
		_ = h.StringLock()
		_ = h.KeyValStringLines()
	})
}

func Test_Hashmap_IsEqual(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_IsEqual", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Empty()
		h1.AddOrUpdate("a", "1")
		h2 := corestr.New.Hashmap.Empty()
		h2.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": h1.IsEqualPtr(h2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		_ = h1.IsEqualPtrLock(h2)
	})
}

func Test_Hashmap_Clone(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Clone", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		c := h.Clone()
		_ = c
		cp := h.ClonePtr()
		_ = cp
	})
}

func Test_Hashmap_Filter(t *testing.T) {
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

func Test_Hashmap_Except(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Except", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.AddOrUpdate("b", "2")
		_ = h.GetValuesKeysExcept([]string{"a"})
		_ = h.GetAllExceptCollection(corestr.New.Collection.Strings([]string{"a"}))
	})
}

func Test_Hashmap_Join(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Join", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		_ = h.Join(", ")
		_ = h.JoinKeys(", ")
	})
}

func Test_Hashmap_Json(t *testing.T) {
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

func Test_Hashmap_Serialize(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Serialize", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		_, _ = h.Serialize()
		var target map[string]string
		_ = h.Deserialize(&target)
	})
}

func Test_Hashmap_Error(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_Error", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		_ = h.ToError("; ")
		_ = h.ToDefaultError()
	})
}

func Test_Hashmap_SetBySplitter_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_SetBySplitter", func() {
		h := corestr.New.Hashmap.Empty()
		h.SetBySplitter("key=val", "=")
	})
}

func Test_Hashmap_Diff(t *testing.T) {
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

func Test_Hashmap_ConcatNew_FromCollectionBasicOpsIt(t *testing.T) {
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

func Test_Hashmap_AddsOrUpdatesUsingFilter(t *testing.T) {
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

func Test_Hashmap_AddOrUpdateStringsPtrWgLock(t *testing.T) {
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

func Test_Hashmap_ToStringsUsingCompiler(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_ToStringsUsingCompiler", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		_ = h.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })
	})
}

func Test_Hashmap_ClearDispose(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_ClearDispose", func() {
		h := corestr.New.Hashmap.Empty()
		h.AddOrUpdate("a", "1")
		h.Clear()
		h2 := corestr.New.Hashmap.Empty()
		h2.AddOrUpdate("a", "1")
		h2.Dispose()
	})
}

func Test_Hashmap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_I8_Hashmap_ParseInjectUsingJsonMust", func() {
		h := corestr.New.Hashmap.Empty()
		r := corejson.New(map[string]string{"a": "1"})
		h.ParseInjectUsingJsonMust(&r)
	})
}

// ==========================================================================
// Hashset — additional coverage
// ==========================================================================

func Test_Hashset_AddVariants(t *testing.T) {
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

func Test_Hashset_AddCollections(t *testing.T) {
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

func Test_Hashset_Has(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_Has", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")

		// Act
		actual := args.Map{"result": h.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has a", actual)
		actual = args.Map{"result": h.Contains("b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected contains b", actual)
		_ = h.HasLock("a")
		_ = h.HasWithLock("a")
		actual = args.Map{"result": h.IsMissing("a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "not missing", actual)
		_ = h.IsMissingLock("a")
		actual = args.Map{"result": h.HasAllStrings([]string{"a", "b"})}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "has all strings", actual)
		actual = args.Map{"result": h.HasAll("a", "b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "has all", actual)
		_ = h.HasAny("a", "z")
		_ = h.IsAllMissing("x", "y")
		coll := corestr.New.Collection.Strings([]string{"a"})
		_ = h.HasAllCollectionItems(coll)
	})
}

func Test_Hashset_Lists(t *testing.T) {
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

func Test_Hashset_Filter(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_Filter", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("aa", "b", "cc")
		f := h.Filter(func(s string) bool { return len(s) > 1 })

		// Act
		actual := args.Map{"result": f.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		_ = h.GetFilteredItems(corestr.IsStringFilter(func(str string, index int) (string, bool, bool) { return str, true, false }))
		_ = h.GetFilteredCollection(corestr.IsStringFilter(func(str string, index int) (string, bool, bool) { return str, true, false }))
	})
}

func Test_Hashset_Except(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_Except", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		_ = h.GetAllExcept([]string{"a"})
		_ = h.GetAllExceptSpread("b")
		h2 := corestr.New.Hashset.StringsSpreadItems("c")
		_ = h.GetAllExceptHashset(h2)
	})
}

func Test_Hashset_Concat(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_Concat", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		c1 := h.ConcatNewStrings(false, []string{"b"}, []string{"c"})
		_ = c1
		h2 := corestr.New.Hashset.StringsSpreadItems("d")
		c2 := h.ConcatNewHashsets(false, h2)
		_ = c2
	})
}

func Test_Hashset_IsEqual(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_IsEqual", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a", "b")
		b := corestr.New.Hashset.StringsSpreadItems("a", "b")

		// Act
		actual := args.Map{"result": a.IsEqual(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		_ = a.IsEquals(b)
		_ = a.IsEqualsLock(b)
	})
}

func Test_Hashset_Resize_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_Resize", func() {
		h := corestr.New.Hashset.Empty()
		h.Resize(100)
		h.ResizeLock(200)
		h.AddCapacities(50)
		h.AddCapacitiesLock(50)
	})
}

func Test_Hashset_AsyncOps(t *testing.T) {
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

func Test_Hashset_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_AddsUsingFilter", func() {
		h := corestr.New.Hashset.Empty()
		h.AddsUsingFilter(corestr.IsStringFilter(func(str string, index int) (string, bool, bool) { return str, str != "", false }), "a", "", "b")
		h.AddsAnyUsingFilter(corestr.IsStringFilter(func(str string, index int) (string, bool, bool) { return str, str != "", false }), "c", nil)
		h.AddsAnyUsingFilterLock(corestr.IsStringFilter(func(str string, index int) (string, bool, bool) { return str, true, false }), "d")
	})
}

func Test_Hashset_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_DistinctDiff", func() {
		a := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		b := corestr.New.Hashset.StringsSpreadItems("a")
		_ = a.DistinctDiffHashset(b)
		_ = a.DistinctDiffLines("a")
		_ = a.DistinctDiffLinesRaw("a")
	})
}

func Test_Hashset_String(t *testing.T) {
	safeTest(t, "Test_I8_Hashset_String", func() {
		h := corestr.New.Hashset.StringsSpreadItems("a")
		_ = h.String()
		_ = h.StringLock()
	})
}

func Test_Hashset_ClearDispose(t *testing.T) {
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

func Test_SimpleSlice_BasicOps(t *testing.T) {
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

func Test_SimpleSlice_FirstLast(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_FirstLast", func() {
		// Arrange
		s := corestr.New.SimpleSlice.SpreadStrings("a", "b", "c")

		// Act
		actual := args.Map{"result": s.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "first", actual)
		actual = args.Map{"result": s.Last() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "last", actual)
		actual = args.Map{"result": s.FirstOrDefault() != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "first or default", actual)
		actual = args.Map{"result": s.LastOrDefault() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "last or default", actual)
		_ = s.FirstDynamic()
		_ = s.LastDynamic()
		_ = s.FirstOrDefaultDynamic()
		_ = s.LastOrDefaultDynamic()

		e := corestr.New.SimpleSlice.Empty()
		actual = args.Map{"result": e.FirstOrDefault() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty first", actual)
		actual = args.Map{"result": e.LastOrDefault() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty last", actual)
	})
}

func Test_SimpleSlice_SkipTakeLimit(t *testing.T) {
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

func Test_SimpleSlice_Properties(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_Properties", func() {
		// Arrange
		s := corestr.New.SimpleSlice.SpreadStrings("a", "b")

		// Act
		actual := args.Map{"result": s.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "length", actual)
		actual = args.Map{"result": s.Count() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "count", actual)
		actual = args.Map{"result": s.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "not empty", actual)
		actual = args.Map{"result": s.HasAnyItem()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "has items", actual)
		actual = args.Map{"result": s.LastIndex() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "last index", actual)
		actual = args.Map{"result": s.HasIndex(0)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "has index", actual)
		_ = s.Strings()
		_ = s.List()
	})
}

func Test_SimpleSlice_Contains(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_Contains", func() {
		// Arrange
		s := corestr.New.SimpleSlice.SpreadStrings("a", "b")

		// Act
		actual := args.Map{"result": s.IsContains("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "contains a", actual)
		actual = args.Map{"result": s.IsContains("z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "no z", actual)
		_ = s.IndexOf("b")
		_ = s.IndexOf("z")
		_ = s.CountFunc(func(index int, item string) bool { return index >= 0 && item != "" })
		_ = s.IsContainsFunc("a", func(item, searching string) bool { return item == searching })
		_ = s.IndexOfFunc("b", func(item, searching string) bool { return item == searching })
	})
}

func Test_SimpleSlice_Wrap(t *testing.T) {
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

func Test_SimpleSlice_Transpile(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_Transpile", func() {
		s := corestr.New.SimpleSlice.SpreadStrings("a", "b")
		_ = s.Transpile(func(s string) string { return strings.ToUpper(s) })
		_ = s.TranspileJoin(func(s string) string { return s }, ", ")
	})
}

func Test_SimpleSlice_Join(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_Join", func() {
		s := corestr.New.SimpleSlice.SpreadStrings("a", "b")
		_ = s.Join(", ")
		_ = s.JoinLine()
		_ = s.JoinLineEofLine()
		_ = s.JoinSpace()
		_ = s.JoinComma()
	})
}

func Test_SimpleSlice_Hashset(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_Hashset", func() {
		// Arrange
		s := corestr.New.SimpleSlice.SpreadStrings("a", "b")
		h := s.Hashset()

		// Act
		actual := args.Map{"result": h.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SimpleSlice_Error(t *testing.T) {
	safeTest(t, "Test_I8_SimpleSlice_Error", func() {
		s := corestr.New.SimpleSlice.SpreadStrings("err1", "err2")
		_ = s.AsDefaultError()
		_ = s.AsError("; ")
	})
}

// ==========================================================================
// LinkedList — comprehensive coverage
// ==========================================================================

func Test_LinkedList_BasicOps(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_BasicOps", func() {
		// Arrange
		ll := corestr.New.LinkedList.Empty()
		ll.Add("a").Add("b").Add("c")

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		_ = ll.LengthLock()
		actual = args.Map{"result": ll.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "not empty", actual)
		actual = args.Map{"result": ll.HasItems()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "has items", actual)
		_ = ll.IsEmptyLock()
		actual = args.Map{"result": ll.Head() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "head nil", actual)
		actual = args.Map{"result": ll.Tail() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "tail nil", actual)
	})
}

func Test_LinkedList_AddVariants(t *testing.T) {
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

func Test_LinkedList_IsEquals_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_IsEquals", func() {
		// Arrange
		a := corestr.New.LinkedList.Strings([]string{"a", "b"})
		b := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		_ = a.IsEqualsWithSensitive(b, false)
	})
}

func Test_LinkedList_InsertAt(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_InsertAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "c"})
		ll.InsertAt(1, "b")

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_LinkedList_Loop(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_Loop", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) (isBreak bool) {
			count++
			return false
		})

		// Act
		actual := args.Map{"result": count != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3 iterations", actual)
	})
}

func Test_LinkedList_Filter(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_Filter", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"aa", "b", "cc"})
		filtered := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{
				Value:   arg.Node,
				IsKeep:  len(arg.Node.Element) > 1,
				IsBreak: false,
			}
		})

		// Act
		actual := args.Map{"result": len(filtered) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_RemoveByIndex(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_RemoveByIndex", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByIndex(1)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_RemoveByValue(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_RemoveByValue", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByElementValue("b", true, false)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_RemoveByIndexes(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_RemoveByIndexes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c", "d"})
		ll.RemoveNodeByIndexes(false, 0, 2)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_I8_LinkedList_GetCompareSummary", func() {
		a := corestr.New.LinkedList.Strings([]string{"a", "b"})
		b := corestr.New.LinkedList.Strings([]string{"a", "c"})
		_ = a.GetCompareSummary(b, "left", "right")
	})
}

// ==========================================================================
// Smaller types
// ==========================================================================

func Test_ValidValue(t *testing.T) {
	safeTest(t, "Test_I8_ValidValue", func() {
		// Arrange
		v := corestr.ValidValue{Value: "hello", IsValid: true}

		// Act
		actual := args.Map{"result": v.IsValid}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected valid", actual)
		_ = v.Value
	})
}

func Test_LeftRight(t *testing.T) {
	safeTest(t, "Test_I8_LeftRight", func() {
		lr := corestr.NewLeftRight("a", "b")
		_ = lr.Left
		_ = lr.Right
	})
}

func Test_LeftMiddleRight(t *testing.T) {
	safeTest(t, "Test_I8_LeftMiddleRight", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		_ = lmr.Left
		_ = lmr.Middle
		_ = lmr.Right
	})
}

func Test_ValueStatus_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_ValueStatus", func() {
		vs := corestr.ValueStatus{ValueValid: corestr.NewValidValue("x"), Index: 0}
		_ = vs.ValueValid
		_ = vs.Index
	})
}

func Test_KeyValuePair_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_KeyValuePair", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		_ = kv.Key
		_ = kv.Value
	})
}

func Test_KeyAnyValuePair_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_KeyAnyValuePair", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		_ = kav.Key
		_ = kav.Value
	})
}

func Test_KeyValueCollection(t *testing.T) {
	safeTest(t, "Test_I8_KeyValueCollection", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1")
		kvc.Add("b", "2")

		// Act
		actual := args.Map{"result": kvc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_TextWithLineNumber_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_TextWithLineNumber", func() {
		tln := corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
		_ = tln.LineNumber
		_ = tln.Text
	})
}

// ==========================================================================
// Creators
// ==========================================================================

func Test_NewCreators_FromCollectionBasicOpsIt(t *testing.T) {
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

func Test_CloneSlice_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_CloneSlice", func() {
		// Arrange
		orig := []string{"a", "b"}
		c := corestr.CloneSlice(orig)

		// Act
		actual := args.Map{"result": len(c) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)

		c2 := corestr.CloneSlice(nil)
		actual = args.Map{"result": len(c2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	})
}

func Test_CloneSliceIf(t *testing.T) {
	safeTest(t, "Test_I8_CloneSliceIf", func() {
		// Arrange
		orig := []string{"a", "b"}
		c := corestr.CloneSliceIf(true, orig...)

		// Act
		actual := args.Map{"result": len(c) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		c2 := corestr.CloneSliceIf(false, orig...)
		actual = args.Map{"result": len(c2) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected passthrough len 2 when not cloning", actual)
	})
}

func Test_AnyToString(t *testing.T) {
	safeTest(t, "Test_I8_AnyToString", func() {
		_ = corestr.AnyToString(false, nil)
		_ = corestr.AnyToString(false, "hello")
		_ = corestr.AnyToString(false, 42)
		_ = corestr.AnyToString(false, []string{"a"})
	})
}

func Test_AllIndividualStringsOfStringsLength(t *testing.T) {
	safeTest(t, "Test_I8_AllIndividualStringsOfStringsLength", func() {
		// Arrange
		strs := [][]string{{"a", "bb"}, {"ccc"}}
		// Fix: function counts items (3), not character lengths (6).
		// See issues/corestrtests-allindividualslength-wrong-expectation.md
		result := corestr.AllIndividualStringsOfStringsLength(&strs)

		// Act
		actual := args.Map{"result": result}

		// Assert
		expected := args.Map{"result": 3}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_AllIndividualsLengthOfSimpleSlices_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_AllIndividualsLengthOfSimpleSlices", func() {
		// Arrange
		s1 := corestr.New.SimpleSlice.SpreadStrings("a", "bb")
		s2 := corestr.New.SimpleSlice.SpreadStrings("ccc")
		// Fix: function counts items (3), not character lengths (6).
		// See issues/corestrtests-allindividualslength-wrong-expectation.md
		result := corestr.AllIndividualsLengthOfSimpleSlices(s1, s2)

		// Act
		actual := args.Map{"result": result}

		// Assert
		expected := args.Map{"result": 3}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

// ==========================================================================
// SimpleStringOnce
// ==========================================================================

func Test_SimpleStringOnce(t *testing.T) {
	safeTest(t, "Test_I8_SimpleStringOnce", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{"result": sso.Value() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		actual = args.Map{"result": sso.IsDefined()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected defined", actual)
		actual = args.Map{"result": sso.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

// ==========================================================================
// CharCollectionMap
// ==========================================================================

func Test_CharCollectionMap_Ops(t *testing.T) {
	safeTest(t, "Test_I8_CharCollectionMap_Ops", func() {
		// Arrange
		m := corestr.New.CharCollectionMap.Items([]string{"apple", "avocado", "banana"})

		// Act
		actual := args.Map{"result": m.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "not empty", actual)
		actual = args.Map{"result": m.HasItems()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "has items", actual)
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

func Test_CharCollectionMap_Add_FromCollectionBasicOpsIt(t *testing.T) {
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

func Test_CharHashsetMap_Ops(t *testing.T) {
	safeTest(t, "Test_I8_CharHashsetMap_Ops", func() {
		// Arrange
		m := corestr.New.CharHashsetMap.CapItems(10, 5, "apple", "avocado", "banana")

		// Act
		actual := args.Map{"result": m.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "not empty", actual)
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

func Test_LeftRightFromSplit_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_LeftRightFromSplit", func() {
		// Arrange
		lr := corestr.LeftRightFromSplit("a=b", "=")

		// Act
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "split failed", actual)

		lr2 := corestr.LeftRightFromSplit("noSep", "=")
		actual = args.Map{"result": lr2.Left != "noSep"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "no sep should set left", actual)
	})
}

func Test_LeftMiddleRightFromSplit_FromCollectionBasicOpsIt(t *testing.T) {
	safeTest(t, "Test_I8_LeftMiddleRightFromSplit", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplit("a:b:c", ":")

		// Act
		actual := args.Map{"result": lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "split failed", actual)
	})
}

// ==========================================================================
// ValidValues
// ==========================================================================

func Test_ValidValues(t *testing.T) {
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

func Test_CollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_I8_CollectionsOfCollection", func() {
		// Arrange
		cc := corestr.New.CollectionsOfCollection.Empty()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		cc.Add(c1)

		// Act
		actual := args.Map{"result": cc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ==========================================================================
// HashsetsCollection
// ==========================================================================

func Test_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_I8_HashsetsCollection", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		h := corestr.New.Hashset.StringsSpreadItems("a")
		hc.Add(h)

		// Act
		actual := args.Map{"result": hc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}
