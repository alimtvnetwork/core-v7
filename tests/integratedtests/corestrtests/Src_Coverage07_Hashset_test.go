package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_SrcC07_Hashset_Basic_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_Basic_Verification", func() {
		// Arrange
		tc := srcC07HashsetBasicTestCase
		h := corestr.New.Hashset.Empty()
		var nilH *corestr.Hashset

		// Act
		actual := args.Map{
			"isEmpty":   h.IsEmpty(),
			"hasItems":  h.HasItems(),
			"hasAny":    h.HasAnyItem(),
			"length":    h.Length(),
			"nilLength": nilH.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_Add_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_Add_Verification", func() {
		// Arrange
		tc := srcC07HashsetAddTestCase
		h := corestr.New.Hashset.Cap(5)

		// Act
		noPanic := !callPanicsSrcC07(func() {
			h.Add("a").Add("b")
			h.AddLock("c")
			h.AddNonEmpty("")
			h.AddNonEmpty("d")
			h.AddNonEmptyWhitespace("   ")
			h.AddNonEmptyWhitespace("e")
			h.AddIf(false, "skip")
			h.AddIf(true, "f")
			h.AddIfMany(false, "x", "y")
			h.AddIfMany(true, "g", "h")
			h.AddFunc(func() string { return "i" })
			h.AddFuncErr(func() (string, error) { return "j", nil }, func(e error) {})
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_AddBool_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_AddBool_Verification", func() {
		// Arrange
		tc := srcC07HashsetAddBoolTestCase
		h := corestr.New.Hashset.Empty()

		// Act
		first := h.AddBool("a")
		second := h.AddBool("a")
		actual := args.Map{
			"firstExisted":  first,
			"secondExisted": second,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_AddPtr_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_AddPtr_Verification", func() {
		// Arrange
		tc := srcC07HashsetAddPtrTestCase
		h := corestr.New.Hashset.Empty()
		s := "hello"

		// Act
		h.AddPtr(&s)
		h.AddPtrLock(&s)
		actual := args.Map{
			"length": h.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_Adds_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_Adds_Verification", func() {
		// Arrange
		tc := srcC07HashsetAddsTestCase

		// Act
		noPanic := !callPanicsSrcC07(func() {
			h := corestr.New.Hashset.Empty()
			h.Adds("a", "b")
			h.Adds()
			h.AddStrings([]string{"c"})
			h.AddStrings(nil)
			h.AddStringsLock([]string{"d"})
			h.AddStringsLock(nil)
			h.AddCollection(corestr.New.Collection.Strings([]string{"e"}))
			h.AddCollection(nil)
			h.AddCollections(corestr.New.Collection.Strings([]string{"f"}))
			h.AddCollections()
			ss := corestr.New.SimpleSlice.Lines("g")
			h.AddSimpleSlice(ss)
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_Has_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_Has_Verification", func() {
		// Arrange
		tc := srcC07HashsetHasTestCase
		h := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")

		// Act
		actual := args.Map{
			"has":          h.Has("a"),
			"contains":     h.Contains("a"),
			"hasLock":      h.HasLock("a"),
			"hasWithLock":  h.HasWithLock("a"),
			"notMissing":   h.IsMissing("a"),
			"zMissing":     h.IsMissing("z"),
			"missingLock":  h.IsMissingLock("a"),
			"hasAll":       h.HasAll("a", "b"),
			"hasAllStr":    h.HasAllStrings([]string{"a", "b"}),
			"hasAnyAZ":     h.HasAny("a", "z"),
			"hasAnyXZ":     h.HasAny("x", "z"),
			"isAllMissXZ":  h.IsAllMissing("x", "z"),
			"isAllMissA":   h.IsAllMissing("a"),
			"hasAllCol":    h.HasAllCollectionItems(corestr.New.Collection.Strings([]string{"a"})),
			"hasAllColNil": h.HasAllCollectionItems(nil),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_Equals_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_Equals_Verification", func() {
		// Arrange
		tc := srcC07HashsetEqualsTestCase
		h1 := corestr.New.Hashset.StringsSpreadItems("a", "b")
		h2 := corestr.New.Hashset.StringsSpreadItems("a", "b")

		// Act
		actual := args.Map{
			"isEquals":     h1.IsEquals(h2),
			"isEqual":      h1.IsEqual(h2),
			"isEqualsLock": h1.IsEqualsLock(h2),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_Remove_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_Remove_Verification", func() {
		// Arrange
		tc := srcC07HashsetRemoveTestCase
		h := corestr.New.Hashset.StringsSpreadItems("a", "b")

		// Act
		noPanic := !callPanicsSrcC07(func() {
			h.Remove("a")
			h.SafeRemove("b")
			h.SafeRemove("z")
			h.RemoveWithLock("z")
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_List_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_List_Verification", func() {
		// Arrange
		tc := srcC07HashsetListTestCase
		h := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		noPanic := !callPanicsSrcC07(func() {
			_ = h.List()
			_ = h.ListPtr()
			_ = h.Lines()
			_ = h.SafeStrings()
			_ = h.ListPtrSortedAsc()
			_ = h.ListPtrSortedDsc()
			_ = h.OrderedList()
			_ = h.SortedList()
			_ = h.ListCopyLock()
			_ = h.SimpleSlice()
			_ = h.Items()
			_ = h.Collection()
			_ = h.MapStringAny()
			_ = h.MapStringAnyDiff()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_Filter_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_Filter_Verification", func() {
		// Arrange
		tc := srcC07HashsetFilterTestCase
		h := corestr.New.Hashset.StringsSpreadItems("abc", "def")
		f := func(s string) bool { return s == "abc" }
		sf := func(s string, i int) (string, bool, bool) { return s, s == "abc", false }

		// Act
		actual := args.Map{
			"filterLen":      h.Filter(f).Length(),
			"filteredItems":  len(h.GetFilteredItems(sf)),
			"filteredColLen": h.GetFilteredCollection(sf).Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_Except_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_Except_Verification", func() {
		// Arrange
		tc := srcC07HashsetExceptTestCase
		h := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")

		// Act
		_ = h.GetAllExceptHashset(nil)
		_ = h.GetAllExcept(nil)
		_ = h.GetAllExceptSpread()
		_ = h.GetAllExceptCollection(nil)
		actual := args.Map{
			"hashsetExcLen":    len(h.GetAllExceptHashset(corestr.New.Hashset.StringsSpreadItems("a"))),
			"keysExcLen":       len(h.GetAllExcept([]string{"a"})),
			"spreadExcLen":     len(h.GetAllExceptSpread("a")),
			"collectionExcLen": len(h.GetAllExceptCollection(corestr.New.Collection.Strings([]string{"a"}))),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_Resize_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_Resize_Verification", func() {
		// Arrange
		tc := srcC07HashsetResizeTestCase
		h := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		noPanic := !callPanicsSrcC07(func() {
			h.Resize(100)
			h.ResizeLock(200)
			h.AddCapacities(10, 20)
			h.AddCapacitiesLock(10)
			h.AddCapacities()
			h.AddCapacitiesLock()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_Concat_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_Concat_Verification", func() {
		// Arrange
		tc := srcC07HashsetConcatTestCase
		h := corestr.New.Hashset.StringsSpreadItems("a")
		h2 := corestr.New.Hashset.StringsSpreadItems("b")

		// Act
		r := h.ConcatNewHashsets(false, h2)
		noPanic := !callPanicsSrcC07(func() {
			_ = h.ConcatNewHashsets(true)
			_ = h.ConcatNewStrings(false, []string{"c"})
			_ = h.ConcatNewStrings(true)
		})
		actual := args.Map{
			"concatLen": r.Length() >= 2,
			"noPanic":   noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_StringJson_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_StringJson_Verification", func() {
		// Arrange
		tc := srcC07HashsetStringJsonTestCase
		h := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		noPanic := !callPanicsSrcC07(func() {
			_ = h.Join(",")
			_ = h.NonEmptyJoins(",")
			_ = h.NonWhitespaceJoins(",")
			_ = h.JoinSorted(",")
			_ = h.JsonModel()
			_ = h.JsonModelAny()
			_, _ = h.MarshalJSON()
			_ = h.AsJsoner()
			_ = h.AsJsonContractsBinder()
			_ = h.AsJsonParseSelfInjector()
			_ = h.AsJsonMarshaller()
		})
		actual := args.Map{
			"stringNonEmpty":     h.String() != "",
			"stringLockNonEmpty": h.StringLock() != "",
			"noPanic":            noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_ToLower_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_ToLower_Verification", func() {
		// Arrange
		tc := srcC07HashsetToLowerTestCase
		h := corestr.New.Hashset.StringsSpreadItems("ABC")

		// Act
		actual := args.Map{
			"hasLowercase": h.ToLowerSet().Has("abc"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_ClearDispose_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_ClearDispose_Verification", func() {
		// Arrange
		h := corestr.New.Hashset.StringsSpreadItems("a")
		var nilH *corestr.Hashset

		// Act
		h.Clear()
		h.Dispose()
		nilH.Dispose()

		// Assert — no panic is success
	})
}

func Test_SrcC07_Hashset_DistinctDiff_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_DistinctDiff_Verification", func() {
		// Arrange
		tc := srcC07HashsetDistinctDiffTestCase
		h := corestr.New.Hashset.StringsSpreadItems("a", "b")
		empty := corestr.New.Hashset.Empty()

		// Act
		r := h.DistinctDiffLinesRaw("b", "c")
		noPanic := !callPanicsSrcC07(func() {
			_ = h.DistinctDiffLines("b", "c")
			_ = h.DistinctDiffHashset(corestr.New.Hashset.StringsSpreadItems("b", "c"))
			_ = empty.DistinctDiffLinesRaw()
			_ = empty.DistinctDiffLinesRaw("a")
			_ = h.DistinctDiffLinesRaw()
			_ = empty.DistinctDiffLines()
			_ = empty.DistinctDiffLines("a")
		})
		actual := args.Map{
			"diffLen": len(r),
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_WgLock_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_WgLock_Verification", func() {
		// Arrange
		tc := srcC07HashsetWgLockTestCase
		h := corestr.New.Hashset.Cap(10)

		// Act
		noPanic := !callPanicsSrcC07(func() {
			wg := &sync.WaitGroup{}
			wg.Add(1)
			h.AddWithWgLock("a", wg)
			wg2 := &sync.WaitGroup{}
			wg2.Add(1)
			h.AddStringsPtrWgLock([]string{"b", "c"}, wg2)
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_AddItemsMap_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_AddItemsMap_Verification", func() {
		// Arrange
		tc := srcC07HashsetAddItemsMapTestCase
		h := corestr.New.Hashset.Empty()

		// Act
		h.AddItemsMap(map[string]bool{"a": true, "b": false})
		h.AddItemsMap(nil)
		noPanic := !callPanicsSrcC07(func() {
			wg := &sync.WaitGroup{}
			wg.Add(1)
			m := map[string]bool{"c": true}
			h.AddItemsMapWgLock(&m, wg)
			h.AddItemsMapWgLock(nil, nil)
		})
		actual := args.Map{
			"length":  h.Length(),
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_AddHashset_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_AddHashset_Verification", func() {
		// Arrange
		tc := srcC07HashsetAddHashsetTestCase

		// Act
		noPanic := !callPanicsSrcC07(func() {
			h := corestr.New.Hashset.Empty()
			h2 := corestr.New.Hashset.StringsSpreadItems("a", "b")
			h.AddHashsetItems(h2)
			h.AddHashsetItems(nil)
			wg := &sync.WaitGroup{}
			wg.Add(1)
			h3 := corestr.New.Hashset.StringsSpreadItems("c")
			h.AddHashsetWgLock(h3, wg)
			h.AddHashsetWgLock(nil, nil)
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_AddsFilter_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_AddsFilter_Verification", func() {
		// Arrange
		tc := srcC07HashsetAddsFilterTestCase
		h := corestr.New.Hashset.Empty()
		f := func(s string, i int) (string, bool, bool) { return s, true, false }

		// Act
		noPanic := !callPanicsSrcC07(func() {
			h.AddsUsingFilter(f, "a", "b")
			h.AddsUsingFilter(f)
			h.AddsAnyUsingFilter(f, "c")
			h.AddsAnyUsingFilter(f)
			h.AddsAnyUsingFilterLock(f, "d")
			h.AddsAnyUsingFilterLock(f)
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC07_Hashset_EmptyString_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC07_Hashset_EmptyString_Verification", func() {
		// Arrange
		tc := srcC07HashsetEmptyStringTestCase
		h := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{
			"stringNonEmpty":     h.String() != "",
			"stringLockNonEmpty": h.StringLock() != "",
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC07(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}
