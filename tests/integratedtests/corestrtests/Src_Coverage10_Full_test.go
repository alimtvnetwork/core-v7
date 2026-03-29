package corestrtests

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Collection extended ──

func Test_SrcC10_Collection_TakeSkip_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Collection_TakeSkip_Verification", func() {
		// Arrange
		tc := srcC10CollectionTakeSkipTestCase
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		s := corestr.New.SimpleSlice.Lines("a", "b", "c", "d", "e")

		// Act
		actual := args.Map{
			"takeLen":     c.Take(2).Length(),
			"skipLen":     c.Skip(2).Length(),
			"limitLen":    len(s.Limit(2)),
			"limitAllLen": len(s.Limit(-1)),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Collection_AddNonEmpty_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Collection_AddNonEmpty_Verification", func() {
		// Arrange
		tc := srcC10CollectionAddNonEmptyTestCase

		// Act
		c1 := corestr.New.Collection.Empty()
		c1.AddNonEmptyStrings("a", "", "b")
		c2 := corestr.New.Collection.Empty()
		c2.AddNonEmptyStringsSlice([]string{"a", "", "b"})
		actual := args.Map{
			"addNonEmptyLen":      c1.Length(),
			"addNonEmptySliceLen": c2.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Collection_NonEmptyList_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Collection_NonEmptyList_Verification", func() {
		// Arrange
		tc := srcC10CollectionNonEmptyListTestCase
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})

		// Act
		actual := args.Map{
			"listLen":    len(c.NonEmptyList()),
			"listPtrLen": len(*c.NonEmptyListPtr()),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Collection_MethodsNoPanic_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Collection_MethodsNoPanic_Verification", func() {
		// Arrange
		tc := srcC10CollectionMethodsNoPanicTestCase

		// Act
		noPanic := !callPanicsSrcC10(func() {
			c := corestr.New.Collection.Strings([]string{"a", "b"})
			_ = c.Items()
			_ = c.ListPtr()
			_ = c.ListCopyPtrLock()
			_ = c.HashsetAsIs()
			_ = c.HashsetWithDoubleLength()
			_ = c.HashsetLock()

			c2 := corestr.New.Collection.Strings([]string{"a", "", "b"})
			_ = c2.NonEmptyItems()
			_ = c2.NonEmptyItemsPtr()

			c3 := corestr.New.Collection.Strings([]string{"a", " ", "b"})
			_ = c3.NonEmptyItemsOrNonWhitespace()
			_ = c3.NonEmptyItemsOrNonWhitespacePtr()

			cf := corestr.New.Collection.Strings([]string{"apple", "banana", "cherry"})
			_ = cf.Filter(func(s string, i int) (string, bool, bool) { return s, len(s) > 5, false })
			_ = cf.FilterLock(func(s string, i int) (string, bool, bool) { return s, len(s) > 1, false })
			_ = cf.FilteredCollection(func(s string, i int) (string, bool, bool) { return s, len(s) > 1, false })
			_ = cf.FilteredCollectionLock(func(s string, i int) (string, bool, bool) { return s, len(s) > 1, false })
			_ = cf.FilterPtr(func(s *string, i int) (*string, bool, bool) { return s, len(*s) > 1, false })
			_ = cf.FilterPtrLock(func(s *string, i int) (*string, bool, bool) { return s, len(*s) > 1, false })

			_ = c.CharCollectionMap()
			c.AddCapacity(10)
			c.Resize(5)

			_ = c.CsvLines()
			_ = c.Csv()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Collection_Has_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Collection_Has_Verification", func() {
		// Arrange
		tc := srcC10CollectionHasTestCase
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		cH := corestr.New.Collection.Strings([]string{"Hello", "World"})

		// Act
		actual := args.Map{
			"has":           c.Has("b"),
			"hasMissing":    c.Has("z"),
			"hasLock":       c.HasLock("a"),
			"hasAll":        c.HasAll("a", "b"),
			"hasAllMissing": c.HasAll("a", "z"),
			"hasSensLower":  cH.HasUsingSensitivity("hello", false),
			"hasSensExact":  cH.HasUsingSensitivity("Hello", true),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Collection_Sorted_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Collection_Sorted_Verification", func() {
		// Arrange
		tc := srcC10CollectionSortedTestCase
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})

		// Act
		actual := args.Map{
			"ascFirst":  c.SortedListAsc()[0],
			"ascCFirst": c.SortedAsc().List()[0],
			"ascLFirst": c.SortedAscLock().List()[0],
			"dscFirst":  c.SortedListDsc()[0],
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Collection_Contains_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Collection_Contains_Verification", func() {
		// Arrange
		tc := srcC10CollectionContainsTestCase
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		v := "a"

		// Act
		hs, hasAll := c.GetHashsetPlusHasAll([]string{"a", "b"})
		actual := args.Map{
			"containsPtr":  c.IsContainsPtr(&v),
			"containsAll":  c.IsContainsAll("a", "b"),
			"containsLock": c.IsContainsAllLock("a", "c"),
			"containsSlc":  c.IsContainsAllSlice([]string{"a", "b"}),
			"hashsetAll":   hasAll && hs != nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Collection_NewExpandMerge_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Collection_NewExpandMerge_Verification", func() {
		// Arrange
		tc := srcC10CollectionNewExpandMergeTestCase

		// Act
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		newC := c1.New("c", "d")

		c2 := corestr.New.Collection.Strings([]string{"a"})
		c2.ExpandSlicePlusAdd([]string{"b", "c"}, func(line string) []string { return []string{line} })

		c3 := corestr.New.Collection.Empty()
		c3.MergeSlicesOfSlice([]string{"a", "b"}, []string{"c"})

		actual := args.Map{
			"newLen":    newC.Length(),
			"expandLen": c2.Length(),
			"mergeLen":  c3.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Collection_Except_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Collection_Except_Verification", func() {
		// Arrange
		tc := srcC10CollectionExceptTestCase
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{
			"exceptColLen": len(c.GetAllExceptCollection(corestr.New.Collection.Strings([]string{"b"}))),
			"exceptLen":    len(c.GetAllExcept([]string{"b"})),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Collection_Joins_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Collection_Joins_Verification", func() {
		// Arrange
		tc := srcC10CollectionJoinsTestCase

		// Act
		c1 := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c2 := corestr.New.Collection.Strings([]string{"a", "", "b"})
		c3 := corestr.New.Collection.Strings([]string{"a", " ", "b"})
		actual := args.Map{
			"joins":         c1.Joins(",") != "",
			"nonEmptyJoins": c2.NonEmptyJoins(",") != "",
			"nonWSJoins":    c3.NonWhitespaceJoins(",") != "",
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Collection_String_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Collection_String_Verification", func() {
		// Arrange
		tc := srcC10CollectionStringTestCase
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		noPanic := !callPanicsSrcC10(func() {
			_ = c.Join(",")
			_ = c.JoinLine()
		})
		actual := args.Map{
			"string":     c.String() != "",
			"stringLock": c.StringLock() != "",
			"summary":    c.SummaryString(1) != "",
			"summaryH":   c.SummaryStringWithHeader("header") != "",
			"noPanic":    noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Collection_Json_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Collection_Json_Verification", func() {
		// Arrange
		tc := srcC10CollectionJsonTestCase

		// Act
		noPanic := !callPanicsSrcC10(func() {
			c := corestr.New.Collection.Strings([]string{"a"})
			_ = c.JsonModelAny()
			_ = c.Json()
			_ = c.JsonPtr()
			_ = c.AsJsonMarshaller()
			_ = c.AsJsonContractsBinder()
			_, _ = c.Serialize()
			var out corestr.Collection
			_ = out.Deserialize(c.JsonPtr())
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Collection_ClearDispose_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Collection_ClearDispose_Verification", func() {
		// Arrange
		tc := srcC10CollectionClearDisposeTestCase
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		c.Clear()
		clearLen := c.Length()
		noPanic := !callPanicsSrcC10(func() {
			c.Add("x")
			c.Dispose()
		})
		actual := args.Map{
			"clearLen": clearLen,
			"noPanic":  noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Collection_AddFunc_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Collection_AddFunc_Verification", func() {
		// Arrange
		tc := srcC10CollectionAddFuncTestCase

		// Act
		c1 := corestr.New.Collection.Empty()
		c1.AddFuncResult(func() string { return "hello" })

		c2 := corestr.New.Collection.Empty()
		c2.AddStringsByFuncChecking([]string{"a", "bb", "c", "dd"}, func(s string) bool { return len(s) > 1 })

		actual := args.Map{
			"funcLen":   c1.Length(),
			"filterLen": c2.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Collection_ParseInject_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Collection_ParseInject_Verification", func() {
		// Arrange
		tc := srcC10CollectionParseInjectTestCase

		// Act
		noPanic := !callPanicsSrcC10(func() {
			c := corestr.New.Collection.Strings([]string{"a"})
			jsonResult := c.JsonPtr()
			c2 := corestr.New.Collection.Empty()
			_, _ = c2.ParseInjectUsingJson(jsonResult)
			c3 := corestr.New.Collection.Empty()
			_ = c3.JsonParseSelfInject(c.JsonPtr())
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── Hashmap extended ──

func Test_SrcC10_Hashmap_Basic_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_Basic_Verification", func() {
		// Arrange
		tc := srcC10HashmapBasicTestCase

		// Act
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		emptyHm := corestr.New.Hashmap.Cap(5)
		actual := args.Map{
			"hasItems":    hm.HasItems(),
			"colNotNil":   hm.Collection() != nil,
			"isEmptyLock": emptyHm.IsEmptyLock(),
			"addLockLen": func() int {
				hm2 := corestr.New.Hashmap.Cap(5)
				hm2.AddOrUpdateLock("a", "1")
				return hm2.Length()
			}(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_Contains_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_Contains_Verification", func() {
		// Arrange
		tc := srcC10HashmapContainsTestCase
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")

		// Act
		actual := args.Map{
			"contains":       hm.Contains("a"),
			"containsLock":   hm.ContainsLock("a"),
			"notMissing":     hm.IsKeyMissing("a"),
			"notMissingLock": hm.IsKeyMissingLock("a"),
			"hasLock":        hm.HasLock("a"),
			"hasAllStr":      hm.HasAllStrings("a", "b"),
			"hasAllStrFail":  hm.HasAllStrings("a", "c"),
			"hasAll":         hm.HasAll("a"),
			"hasAny":         hm.HasAny("a", "z"),
			"hasWithLock":    hm.HasWithLock("a"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_KeysVals_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_KeysVals_Verification", func() {
		// Arrange
		tc := srcC10HashmapKeysValsTestCase

		// Act
		noPanic := !callPanicsSrcC10(func() {
			hm := corestr.New.Hashmap.Cap(5)
			hm.AddOrUpdate("a", "1")
			_ = hm.Keys()
			_ = hm.KeysCollection()
			_ = hm.AllKeys()
			_ = hm.KeysLock()
			_ = hm.ValuesList()
			_ = hm.ValuesCollection()
			_ = hm.ValuesHashset()
			_ = hm.ValuesCollectionLock()
			_ = hm.ValuesHashsetLock()
			_, _ = hm.KeysValuesCollection()
			_, _ = hm.KeysValuesList()
			_ = hm.KeysValuePairs()
			_ = hm.KeysValuePairsCollection()
			_, _ = hm.KeysValuesListLock()
			_ = hm.LengthLock()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_Remove_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_Remove_Verification", func() {
		// Arrange
		tc := srcC10HashmapRemoveTestCase

		// Act
		hm1 := corestr.New.Hashmap.Cap(5)
		hm1.AddOrUpdate("a", "1")
		hm1.Remove("a")
		hm2 := corestr.New.Hashmap.Cap(5)
		hm2.AddOrUpdate("a", "1")
		hm2.RemoveWithLock("a")
		actual := args.Map{
			"removeLen":     hm1.Length(),
			"removeLockLen": hm2.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_String_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_String_Verification", func() {
		// Arrange
		tc := srcC10HashmapStringTestCase
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{
			"string":     hm.String() != "",
			"stringLock": hm.StringLock() != "",
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_ItemsCopy_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_ItemsCopy_Verification", func() {
		// Arrange
		tc := srcC10HashmapItemsCopyTestCase

		// Act
		noPanic := !callPanicsSrcC10(func() {
			hm := corestr.New.Hashmap.Cap(5)
			hm.AddOrUpdate("a", "1")
			_ = hm.ItemsCopyLock()
			_ = hm.SafeItems()
			_ = hm.ValuesListCopyLock()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_Mutate_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_Mutate_Verification", func() {
		// Arrange
		tc := srcC10HashmapMutateTestCase

		// Act
		noPanic := !callPanicsSrcC10(func() {
			hm := corestr.New.Hashmap.Cap(5)
			hm.AddOrUpdate("a", "HELLO")
			hm.ValuesToLower()
			hm2 := corestr.New.Hashmap.Cap(5)
			hm2.AddOrUpdate("ABC", "1")
			hm2.KeysToLower()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_EqualClone_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_EqualClone_Verification", func() {
		// Arrange
		tc := srcC10HashmapEqualCloneTestCase
		hm1 := corestr.New.Hashmap.Cap(5)
		hm1.AddOrUpdate("a", "1")
		hm2 := corestr.New.Hashmap.Cap(5)
		hm2.AddOrUpdate("a", "1")

		// Act
		cloneVal := hm1.Clone()
		actual := args.Map{
			"isEqual":     hm1.IsEqual(*hm2),
			"isEqualLock": hm1.IsEqualPtrLock(hm2),
			"cloneLen":    (&cloneVal).Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_GetValue_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_GetValue_Verification", func() {
		// Arrange
		tc := srcC10HashmapGetValueTestCase
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		v, found := hm.GetValue("a")
		actual := args.Map{
			"value":        v,
			"found":        found,
			"joinNonEmpty": hm.Join(",") != "",
			"keysNonEmpty": hm.JoinKeys(",") != "",
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_Dispose_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_Dispose_Verification", func() {
		// Arrange
		tc := srcC10HashmapDisposeTestCase

		// Act
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("err", "something")
		hasErr := hm.ToError(",") != nil
		hasDefErr := hm.ToDefaultError() != nil
		noPanic := !callPanicsSrcC10(func() {
			hm.Dispose()
		})
		actual := args.Map{
			"noPanic":   noPanic,
			"hasError":  hasErr,
			"hasDefErr": hasDefErr,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_Json_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_Json_Verification", func() {
		// Arrange
		tc := srcC10HashmapJsonTestCase

		// Act
		noPanic := !callPanicsSrcC10(func() {
			hm := corestr.New.Hashmap.Cap(5)
			hm.AddOrUpdate("a", "1")
			_ = hm.JsonModelAny()
			_ = hm.Json()
			_ = hm.JsonPtr()
			_ = hm.AsJsoner()
			_ = hm.AsJsonContractsBinder()
			_ = hm.AsJsonParseSelfInjector()
			_ = hm.AsJsonMarshaller()
			_, _ = hm.Serialize()
			var hm2 corestr.Hashmap
			_ = hm2.Deserialize(hm.JsonPtr())
			hm3 := corestr.New.Hashmap.Empty()
			_, _ = hm3.ParseInjectUsingJson(hm.Json().Ptr())
			hm4 := corestr.New.Hashmap.Empty()
			_ = hm4.JsonParseSelfInject(hm.JsonPtr())
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_AddVariants_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_AddVariants_Verification", func() {
		// Arrange
		tc := srcC10HashmapAddVariantsTestCase

		// Act
		noPanic := !callPanicsSrcC10(func() {
			hm := corestr.New.Hashmap.Cap(5)
			hm.AddOrUpdate("a", "1")
			hm2 := corestr.New.Hashmap.Cap(5)
			hm2.AddOrUpdateHashmap(hm)
			hm3 := corestr.New.Hashmap.Cap(5)
			hm3.AddOrUpdateMap(map[string]string{"a": "1", "b": "2"})
			hm4 := corestr.New.Hashmap.Cap(5)
			hm4.AddsOrUpdates(corestr.KeyValuePair{Key: "a", Value: "1"}, corestr.KeyValuePair{Key: "b", Value: "2"})
			hm5 := corestr.New.Hashmap.Cap(5)
			hm5.AddOrUpdateCollection(
				corestr.New.Collection.Strings([]string{"a"}),
				corestr.New.Collection.Strings([]string{"1"}),
			)
			hm6 := corestr.New.Hashmap.Cap(5)
			var wg sync.WaitGroup
			wg.Add(1)
			hm6.AddOrUpdateWithWgLock("a", "1", &wg)
			wg.Wait()
			hm7 := corestr.New.Hashmap.Cap(5)
			var wg2 sync.WaitGroup
			wg2.Add(1)
			hm7.AddOrUpdateStringsPtrWgLock(&wg2, []string{"a", "b"}, []string{"1", "2"})
			wg2.Wait()
			hm8 := corestr.New.Hashmap.Cap(5)
			hm8.AddOrUpdateKeyAnyValues(corestr.KeyAnyValuePair{Key: "a", Value: 1}, corestr.KeyAnyValuePair{Key: "b", Value: "hello"})
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_Set_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_Set_Verification", func() {
		// Arrange
		tc := srcC10HashmapSetTestCase

		// Act
		hm1 := corestr.New.Hashmap.Cap(5)
		hm1.Set("a", "1")
		hm2 := corestr.New.Hashmap.Cap(5)
		hm2.SetTrim(" a ", " 1 ")
		trimVal, _ := hm2.GetValue("a")
		hm3 := corestr.New.Hashmap.Cap(5)
		hm3.SetBySplitter("key=value", "=")
		splitVal, _ := hm3.GetValue("key")
		actual := args.Map{
			"setLen":   hm1.Length(),
			"trimVal":  trimVal,
			"splitVal": splitVal,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_Diff_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_Diff_Verification", func() {
		// Arrange
		tc := srcC10HashmapDiffTestCase

		// Act
		noPanic := !callPanicsSrcC10(func() {
			hm1 := corestr.New.Hashmap.Cap(5)
			hm1.AddOrUpdate("a", "1")
			hm1.AddOrUpdate("b", "2")
			hm2 := corestr.New.Hashmap.Cap(5)
			hm2.AddOrUpdate("a", "1")
			_ = hm1.DiffRaw(hm2.Items())
			_ = hm1.Diff(hm2)
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_FilterExcept_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_FilterExcept_Verification", func() {
		// Arrange
		tc := srcC10HashmapFilterExceptTestCase
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")

		// Act
		actual := args.Map{
			"hasAllCol": hm.HasAllCollectionItems(corestr.New.Collection.Strings([]string{"a"})),
			"filterLen": len(hm.GetKeysFilteredItems(func(k string, _ int) (string, bool, bool) {
				return k, k == "a", false
			})),
			"filterColNN": hm.GetKeysFilteredCollection(func(k string, _ int) (string, bool, bool) {
				return k, k == "a", false
			}) != nil,
			"exceptVals": len(hm.GetValuesExceptKeysInHashset(corestr.New.Hashset.Strings([]string{"a"}))),
			"exceptKeys": len(hm.GetValuesKeysExcept([]string{"a"})),
			"exceptCol":  hm.GetAllExceptCollection(corestr.New.Collection.Strings([]string{"a"})) != nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_CompilerConcat_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_CompilerConcat_Verification", func() {
		// Arrange
		tc := srcC10HashmapCompilerConcatTestCase
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")
		hm2 := corestr.New.Hashmap.Cap(5)
		hm2.AddOrUpdate("b", "2")

		// Act
		actual := args.Map{
			"compilerLen":  len(hm.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })),
			"concatLen":    hm.ConcatNew(false, hm2).Length(),
			"concatMapLen": hm.ConcatNewUsingMaps(false, map[string]string{"b": "2"}).Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_FilterFunc_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_FilterFunc_Verification", func() {
		// Arrange
		tc := srcC10HashmapFilterFuncTestCase
		hm := corestr.New.Hashmap.Cap(5)

		// Act
		hm.AddsOrUpdatesUsingFilter(func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Value, pair.Key != "skip", false
		}, corestr.KeyValuePair{Key: "a", Value: "1"}, corestr.KeyValuePair{Key: "skip", Value: "2"}, corestr.KeyValuePair{Key: "b", Value: "3"})
		filterLen := hm.Length()
		noPanic := !callPanicsSrcC10(func() {
			hm2 := corestr.New.Hashmap.Cap(5)
			hm2.AddsOrUpdatesAnyUsingFilter(func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
				return fmt.Sprintf("%v", pair.Value), pair.Key == "a", false
			}, corestr.KeyAnyValuePair{Key: "a", Value: "1"}, corestr.KeyAnyValuePair{Key: "b", Value: 2})
			hm3 := corestr.New.Hashmap.Cap(5)
			hm3.AddsOrUpdatesAnyUsingFilterLock(func(pair corestr.KeyAnyValuePair) (string, bool, bool) {
				return fmt.Sprintf("%v", pair.Value), true, false
			}, corestr.KeyAnyValuePair{Key: "a", Value: "1"})
		})
		actual := args.Map{
			"filterLen": filterLen,
			"noPanic":   noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_AddTyped_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_AddTyped_Verification", func() {
		// Arrange
		tc := srcC10HashmapAddTypedTestCase

		// Act
		noPanic := !callPanicsSrcC10(func() {
			hm := corestr.New.Hashmap.Cap(5)
			hm.AddOrUpdateKeyStrValInt("a", 42)
			hm.AddOrUpdateKeyStrValFloat("b", 3.14)
			hm.AddOrUpdateKeyStrValFloat64("c", 3.14)
			hm.AddOrUpdateKeyStrValAny("d", "hello")
			hm.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "e", Value: 42})
			hm.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "f", Value: "1"})
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_Hashmap_KeyValLines_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashmap_KeyValLines_Verification", func() {
		// Arrange
		tc := srcC10HashmapKeyValLinesTestCase
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{
			"linesLen": len(hm.KeyValStringLines()),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── Hashset extended ──

func Test_SrcC10_Hashset_Extended_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_Hashset_Extended_Verification", func() {
		// Arrange
		tc := srcC10HashsetExtendedTestCase

		// Act
		noPanic := !callPanicsSrcC10(func() {
			hs := corestr.New.Hashset.Strings([]string{"a", "b"})
			_ = hs.Has("a")
			hs.Add("c")
			_ = hs.HasAll("a", "b")
			_ = hs.HasAny("a", "z")
			hs2 := corestr.New.Hashset.Strings([]string{"a", "b"})
			hs2.Remove("a")
			hs3 := corestr.New.Hashset.Empty()
			hs3.Adds("a", "b", "c")
			_ = hs3.List()
			_ = hs3.ListPtrSortedAsc()
			_ = hs3.Collection()
			_ = hs3.ListCopyLock()
			_ = hs3.String()
			_ = hs3.Join(",")
			_ = hs3.JoinLine()
			hs3.Dispose()
			hs4 := corestr.New.Hashset.Strings([]string{"a"})
			_ = hs4.Json()
			_ = hs4.JsonPtr()
			_ = hs4.AsJsoner()
			_ = hs4.AsJsonContractsBinder()
			_ = hs4.AsJsonParseSelfInjector()
			_ = hs4.AsJsonMarshaller()
			_, _ = hs4.Serialize()
			var hs5 corestr.Hashset
			_ = hs5.Deserialize(hs4.JsonPtr())
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── newCreator paths ──

func Test_SrcC10_NewCreator_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_NewCreator_Verification", func() {
		// Arrange
		tc := srcC10NewCreatorTestCase

		// Act
		noPanic := !callPanicsSrcC10(func() {
			_ = corestr.New.Collection.Empty()
			_ = corestr.New.Collection.CloneStrings([]string{"a", "b"})
			_ = corestr.New.Collection.Create([]string{"a", "b"})
			_ = corestr.New.Collection.StringsPlusCap(10, []string{"a"})
			_ = corestr.New.Collection.CapStrings(10, []string{"a", "b"})
			_ = corestr.New.Collection.LenCap(0, 10)
			_ = corestr.New.Collection.LineDefault("a\nb\nc")
			_ = corestr.New.Collection.LineUsingSep("a,b,c", ",")
			_ = corestr.New.Collection.StringsOptions(true, []string{"a", "b"})
			_ = corestr.New.Hashmap.Empty()
			_ = corestr.New.Hashmap.MapWithCap(5, map[string]string{"a": "1"})
			_ = corestr.New.Hashset.Empty()
			_ = corestr.New.Hashset.StringsOption(0, true, "a")
			_ = corestr.New.Hashset.StringsSpreadItems("a", "b")
			_ = corestr.New.Hashset.UsingCollection(corestr.New.Collection.Strings([]string{"a"}))
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── emptyCreator paths ──

func Test_SrcC10_EmptyCreator_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_EmptyCreator_Verification", func() {
		// Arrange
		tc := srcC10EmptyCreatorTestCase

		// Act
		noPanic := !callPanicsSrcC10(func() {
			_ = corestr.Empty.LinkedList()
			_ = corestr.Empty.KeyValuePair()
			_ = corestr.Empty.KeyAnyValuePair()
			_ = corestr.Empty.KeyValueCollection()
			_ = corestr.Empty.LinkedCollections()
			_ = corestr.Empty.LeftRight()
			_ = corestr.Empty.SimpleStringOnce()
			_ = corestr.Empty.SimpleStringOncePtr()
			_ = corestr.Empty.HashsetsCollection()
			_ = corestr.Empty.CharCollectionMap()
			_ = corestr.Empty.KeyValuesCollection()
			_ = corestr.Empty.CollectionsOfCollection()
			_ = corestr.Empty.CharHashsetMap()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── Standalone functions ──

func Test_SrcC10_AnyToString_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_AnyToString_Verification", func() {
		// Arrange
		tc := srcC10AnyToStringTestCase

		// Act
		noPanic := !callPanicsSrcC10(func() {
			_ = corestr.AnyToString(false, nil)
		})
		actual := args.Map{
			"intNonEmpty": corestr.AnyToString(false, 42) != "",
			"stringVal":   corestr.AnyToString(false, "hello"),
			"noPanic":     noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_AllIndividualStringsLen_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_AllIndividualStringsLen_Verification", func() {
		// Arrange
		tc := srcC10AllIndividualStringsLenTestCase

		// Act
		input := [][]string{{"a", "b"}, {"c"}}
		actual := args.Map{
			"length": corestr.AllIndividualStringsOfStringsLength(&input),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC10_AllIndividualsSimpleSlices_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC10_AllIndividualsSimpleSlices_Verification", func() {
		// Arrange
		tc := srcC10AllIndividualsSimpleSlicesTestCase
		ss1 := corestr.New.SimpleSlice.Lines("a", "b")
		ss2 := corestr.New.SimpleSlice.Lines("c")

		// Act
		actual := args.Map{
			"length": corestr.AllIndividualsLengthOfSimpleSlices(ss1, ss2),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC10(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}
