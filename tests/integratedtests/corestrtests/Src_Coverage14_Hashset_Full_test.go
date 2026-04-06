package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_SrcC14_Hashset_CapResize_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_CapResize_Verification", func() {
		// Arrange
		tc := srcC14HashsetCapResizeTestCase

		// Act
		hs1 := corestr.New.Hashset.Cap(5); hs1.AddCapacitiesLock(10); hs1.Add("a")
		hs2 := corestr.New.Hashset.Cap(5); hs2.AddCapacities(10, 5); hs2.Add("a")
		hs3 := corestr.New.Hashset.Cap(5); hs3.Add("a"); hs3.Resize(20)
		hs4 := corestr.New.Hashset.Cap(5); hs4.Add("a"); hs4.ResizeLock(20)
		actual := args.Map{
			"capLockHas": hs1.Has("a"),
			"capHas":     hs2.Has("a"),
			"resizeHas":  hs3.Has("a"),
			"resizeLHas": hs4.Has("a"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_CollMisc_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_CollMisc_Verification", func() {
		// Arrange
		tc := srcC14HashsetCollMiscTestCase

		// Act
		actual := args.Map{
			"colNonE":       !corestr.New.Hashset.Strings([]string{"a"}).Collection().IsEmpty(),
			"emptyLock":     corestr.New.Hashset.Cap(0).IsEmptyLock(),
			"concatHsGe2":  corestr.New.Hashset.Strings([]string{"a"}).ConcatNewHashsets(false, corestr.New.Hashset.Strings([]string{"b"})).Length() >= 2,
			"concatHsEGe1": corestr.New.Hashset.Strings([]string{"a"}).ConcatNewHashsets(true).Length() >= 1,
			"concatStrGe2": corestr.New.Hashset.Strings([]string{"a"}).ConcatNewStrings(false, []string{"b"}).Length() >= 2,
			"concatStrEGe1": corestr.New.Hashset.Strings([]string{"a"}).ConcatNewStrings(true).Length() >= 1,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_AddVariants_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_AddVariants_Verification", func() {
		// Arrange
		tc := srcC14HashsetAddVariantsTestCase
		s := "a"

		// Act
		hsPtr := corestr.New.Hashset.Cap(5); hsPtr.AddPtr(&s)
		hsPtrL := corestr.New.Hashset.Cap(5); hsPtrL.AddPtrLock(&s)
		hsBool := corestr.New.Hashset.Cap(5); b1 := hsBool.AddBool("a"); b2 := hsBool.AddBool("a")
		hsNE := corestr.New.Hashset.Cap(5); hsNE.AddNonEmpty(""); hsNE.AddNonEmpty("a")
		hsNEWS := corestr.New.Hashset.Cap(5); hsNEWS.AddNonEmptyWhitespace("  "); hsNEWS.AddNonEmptyWhitespace("a")
		hsIf := corestr.New.Hashset.Cap(5); hsIf.AddIf(true, "a"); hsIf.AddIf(false, "b")
		hsIfM := corestr.New.Hashset.Cap(5); hsIfM.AddIfMany(true, "a", "b"); hsIfM.AddIfMany(false, "c")
		hsF := corestr.New.Hashset.Cap(5); hsF.AddFunc(func() string { return "a" })
		hsS := corestr.New.Hashset.Cap(5); hsS.AddStrings([]string{"a", "b"})
		hsSL := corestr.New.Hashset.Cap(5); hsSL.AddStringsLock([]string{"a"})
		hsC := corestr.New.Hashset.Cap(5); hsC.AddCollection(corestr.New.Collection.Strings([]string{"a"}))
		hsCs := corestr.New.Hashset.Cap(5); hsCs.AddCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"}))
		hsHI := corestr.New.Hashset.Cap(5); hsHI.AddHashsetItems(corestr.New.Hashset.Strings([]string{"a"}))
		hsIM := corestr.New.Hashset.Cap(5); hsIM.AddItemsMap(map[string]bool{"a": true, "b": false})
		hsLk := corestr.New.Hashset.Cap(5); hsLk.AddLock("a")
		hsAdds := corestr.New.Hashset.Cap(5); hsAdds.Adds("a", "b")

		actual := args.Map{
			"addPtrHas":       hsPtr.Has("a"),
			"addPtrLockHas":   hsPtrL.Has("a"),
			"addBoolFirst":    b1,
			"addBoolSecond":   b2,
			"nonEmptyLen":     hsNE.Length(),
			"nonEmptyWSLen":   hsNEWS.Length(),
			"addIfLen":        hsIf.Length(),
			"addIfManyLen":    hsIfM.Length(),
			"addFuncLen":      hsF.Length(),
			"addStringsLen":   hsS.Length(),
			"addStringsLkLen": hsSL.Length(),
			"addCollLen":      hsC.Length(),
			"addCollsLen":     hsCs.Length(),
			"addHsItemsHas":  hsHI.Has("a"),
			"addItemsMapHas": hsIM.Has("a"),
			"addLockHas":      hsLk.Has("a"),
			"addsLen":         hsAdds.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_WgLock_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_WgLock_Verification", func() {
		// Arrange
		tc := srcC14HashsetWgLockTestCase
		hs := corestr.New.Hashset.Cap(5)
		wg := sync.WaitGroup{}
		wg.Add(1)

		// Act
		hs.AddWithWgLock("a", &wg)
		wg.Wait()
		actual := args.Map{
			"has": hs.Has("a"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_Query_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_Query_Verification", func() {
		// Arrange
		tc := srcC14HashsetQueryTestCase

		// Act
		actual := args.Map{
			"hasAnyItem":     corestr.New.Hashset.Strings([]string{"a"}).HasAnyItem(),
			"isMissingA":     corestr.New.Hashset.Strings([]string{"a"}).IsMissing("a"),
			"isMissingB":     corestr.New.Hashset.Strings([]string{"a"}).IsMissing("b"),
			"isMissingLockA": corestr.New.Hashset.Strings([]string{"a"}).IsMissingLock("a"),
			"containsA":      corestr.New.Hashset.Strings([]string{"a"}).Contains("a"),
			"isEqual":        corestr.New.Hashset.Strings([]string{"a"}).IsEqual(corestr.New.Hashset.Strings([]string{"a"})),
			"sortedFirst":    corestr.New.Hashset.Strings([]string{"c", "a", "b"}).SortedList()[0],
			"filterLen":      corestr.New.Hashset.Strings([]string{"abc", "def"}).Filter(func(s string) bool { return s == "abc" }).Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_HasVariants_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_HasVariants_Verification", func() {
		// Arrange
		tc := srcC14HashsetHasVariantsTestCase
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"hasLock":     corestr.New.Hashset.Strings([]string{"a"}).HasLock("a"),
			"hasAllStr":   hs.HasAllStrings([]string{"a", "b"}),
			"hasAllColl":  hs.HasAllCollectionItems(corestr.New.Collection.Strings([]string{"a", "b"})),
			"hasAll":      hs.HasAll("a", "b"),
			"allMissingT": corestr.New.Hashset.Strings([]string{"a"}).IsAllMissing("x", "y"),
			"allMissingF": corestr.New.Hashset.Strings([]string{"a"}).IsAllMissing("a"),
			"hasAny":      corestr.New.Hashset.Strings([]string{"a"}).HasAny("a", "z"),
			"hasAnyMiss":  corestr.New.Hashset.Strings([]string{"a"}).HasAny("x", "y"),
			"hasWithLock": corestr.New.Hashset.Strings([]string{"a"}).HasWithLock("a"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_List_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_List_Verification", func() {
		// Arrange
		tc := srcC14HashsetListTestCase

		// Act
		actual := args.Map{
			"orderedFirst":  corestr.New.Hashset.Strings([]string{"c", "a", "b"}).OrderedList()[0],
			"safeStrLen":    len(corestr.New.Hashset.Strings([]string{"a"}).SafeStrings()),
			"linesLen":      len(corestr.New.Hashset.Strings([]string{"a"}).Lines()),
			"simpleSliceNN": corestr.New.Hashset.Strings([]string{"a"}).SimpleSlice() != nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_FilterExcept_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_FilterExcept_Verification", func() {
		// Arrange
		tc := srcC14HashsetFilterExceptTestCase
		fAll := func(s string, i int) (string, bool, bool) { return s, true, false }

		// Act
		actual := args.Map{
			"filteredLen":   len(corestr.New.Hashset.Strings([]string{"a", "b"}).GetFilteredItems(fAll)),
			"filteredColNE": !corestr.New.Hashset.Strings([]string{"a"}).GetFilteredCollection(fAll).IsEmpty(),
			"exceptHsLen":   len(corestr.New.Hashset.Strings([]string{"a", "b", "c"}).GetAllExceptHashset(corestr.New.Hashset.Strings([]string{"b"}))),
			"exceptLen":     len(corestr.New.Hashset.Strings([]string{"a", "b"}).GetAllExcept([]string{"a"})),
			"exceptSpLen":   len(corestr.New.Hashset.Strings([]string{"a", "b"}).GetAllExceptSpread("a")),
			"exceptColLen":  len(corestr.New.Hashset.Strings([]string{"a", "b"}).GetAllExceptCollection(corestr.New.Collection.Strings([]string{"a"}))),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_Items_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_Items_Verification", func() {
		// Arrange
		tc := srcC14HashsetItemsTestCase

		// Act
		noPanic := !callPanicsSrcC14(func() {
			_ = corestr.New.Hashset.Strings([]string{"a"}).MapStringAnyDiff()
		})
		actual := args.Map{
			"itemsLen":     len(corestr.New.Hashset.Strings([]string{"a"}).Items()),
			"listLen":      len(corestr.New.Hashset.Strings([]string{"a"}).List()),
			"mapStrAnyLen": len(corestr.New.Hashset.Strings([]string{"a"}).MapStringAny()),
			"noPanic":      noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_SortJoin_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_SortJoin_Verification", func() {
		// Arrange
		tc := srcC14HashsetSortJoinTestCase

		// Act
		actual := args.Map{
			"joinSortedNE": corestr.New.Hashset.Strings([]string{"b", "a"}).JoinSorted(",") != "",
			"ascFirst":     corestr.New.Hashset.Strings([]string{"c", "a"}).ListPtrSortedAsc()[0],
			"dscFirst":     corestr.New.Hashset.Strings([]string{"a", "c"}).ListPtrSortedDsc()[0],
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_Clear_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_Clear_Verification", func() {
		// Arrange
		tc := srcC14HashsetClearTestCase
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.Clear()

		// Act
		actual := args.Map{
			"clearLen":    hs.Length(),
			"listCopyLen": len(corestr.New.Hashset.Strings([]string{"a"}).ListCopyLock()),
			"lowerHas":    corestr.New.Hashset.Strings([]string{"ABC"}).ToLowerSet().Has("abc"),
			"lengthLock":  corestr.New.Hashset.Strings([]string{"a"}).LengthLock(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_EqualRemove_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_EqualRemove_Verification", func() {
		// Arrange
		tc := srcC14HashsetEqualRemoveTestCase

		// Act
		hs1 := corestr.New.Hashset.Strings([]string{"a"}); hs1.Remove("a")
		hs2 := corestr.New.Hashset.Strings([]string{"a"}); hs2.SafeRemove("a"); hs2.SafeRemove("missing")
		hs3 := corestr.New.Hashset.Strings([]string{"a"}); hs3.RemoveWithLock("a")
		actual := args.Map{
			"isEquals":     corestr.New.Hashset.Strings([]string{"a"}).IsEquals(corestr.New.Hashset.Strings([]string{"a"})),
			"isEqualsLock": corestr.New.Hashset.Strings([]string{"a"}).IsEqualsLock(corestr.New.Hashset.Strings([]string{"a"})),
			"removeOk":     !hs1.Has("a"),
			"safeRemoveOk": !hs2.Has("a"),
			"removeLockOk": !hs3.Has("a"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_String_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_String_Verification", func() {
		// Arrange
		tc := srcC14HashsetStringTestCase

		// Act
		actual := args.Map{
			"strNonE":     corestr.New.Hashset.Strings([]string{"a"}).String() != "",
			"strLockNonE": corestr.New.Hashset.Strings([]string{"a"}).StringLock() != "",
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_Unmarshal_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_Unmarshal_Verification", func() {
		// Arrange
		tc := srcC14HashsetUnmarshalTestCase

		// Act
		hs := &corestr.Hashset{}
		_ = hs.UnmarshalJSON([]byte(`{"a":true}`))
		actual := args.Map{
			"unmarshalLen": hs.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_JsonInterfaces_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_JsonInterfaces_Verification", func() {
		// Arrange
		tc := srcC14HashsetJsonInterfacesTestCase

		// Act
		noPanic := !callPanicsSrcC14(func() {
			_ = corestr.New.Hashset.Cap(0).AsJsonContractsBinder()
			_ = corestr.New.Hashset.Cap(0).AsJsoner()
			_ = corestr.New.Hashset.Cap(0).AsJsonMarshaller()
			_ = corestr.New.Hashset.Cap(0).AsJsonParseSelfInjector()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_Diff_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_Diff_Verification", func() {
		// Arrange
		tc := srcC14HashsetDiffTestCase

		// Act
		actual := args.Map{
			"diffRawLen":   len(corestr.New.Hashset.Strings([]string{"a", "b"}).DistinctDiffLinesRaw("b", "c")),
			"diffRawEmpty": len(corestr.New.Hashset.Cap(0).DistinctDiffLinesRaw()),
			"diffHsLen":    len(corestr.New.Hashset.Strings([]string{"a"}).DistinctDiffHashset(corestr.New.Hashset.Strings([]string{"b"}))),
			"diffLinesLen": len(corestr.New.Hashset.Strings([]string{"a"}).DistinctDiffLines("b")),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_Serialize_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_Serialize_Verification", func() {
		// Arrange
		tc := srcC14HashsetSerializeTestCase
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		b, sErr := hs.Serialize()
		var target map[string]bool
		dErr := hs.Deserialize(&target)
		actual := args.Map{
			"serializeOk":   sErr == nil && len(b) > 0,
			"deserializeOk": dErr == nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_Wrap_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_Wrap_Verification", func() {
		// Arrange
		tc := srcC14HashsetWrapTestCase

		// Act
		noPanic := !callPanicsSrcC14(func() {
			hs := corestr.New.Hashset.Strings([]string{"a"})
			_ = hs.WrapDoubleQuote()
			_ = hs.WrapSingleQuote()
			_ = hs.WrapDoubleQuoteIfMissing()
			_ = hs.WrapSingleQuoteIfMissing()
			_ = corestr.New.Hashset.Cap(0).Transpile(func(s string) string { return s })
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC14_Hashset_FilterAdd_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC14_Hashset_FilterAdd_Verification", func() {
		// Arrange
		tc := srcC14HashsetFilterAddTestCase
		fAll := func(s string, i int) (string, bool, bool) { return s, true, false }

		// Act
		h1 := corestr.New.Hashset.Cap(5); h1.AddsUsingFilter(fAll, "a", "b")
		h2 := corestr.New.Hashset.Cap(5); h2.AddsAnyUsingFilter(fAll, "a", nil)
		h3 := corestr.New.Hashset.Cap(5); h3.AddsAnyUsingFilterLock(fAll, "a")
		h4 := corestr.New.Hashset.Cap(5); h4.AddFuncErr(func() (string, error) { return "a", nil }, func(err error) {})
		h5 := corestr.New.Hashset.Cap(5)
		wg1 := sync.WaitGroup{}; wg1.Add(1)
		h5.AddStringsPtrWgLock([]string{"a"}, &wg1); wg1.Wait()
		h6 := corestr.New.Hashset.Cap(5)
		wg2 := sync.WaitGroup{}; wg2.Add(1)
		h6.AddHashsetWgLock(corestr.New.Hashset.Strings([]string{"a"}), &wg2); wg2.Wait()
		h7 := corestr.New.Hashset.Cap(5); h7.AddSimpleSlice(corestr.New.SimpleSlice.SpreadStrings("a", "b"))
		actual := args.Map{
			"addsFilterLen":    h1.Length(),
			"addsAnyFilterLen": h2.Length(),
			"addsAnyFLockLen":  h3.Length(),
			"funcErrLen":       h4.Length(),
			"strPtrWgHas":      h5.Has("a"),
			"hsWgHas":          h6.Has("a"),
			"simpleSliceLen":   h7.Length(),
			"listPtrLen":       len(corestr.New.Hashset.Strings([]string{"a"}).ListPtr()),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC14(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}
