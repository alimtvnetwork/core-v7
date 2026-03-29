package corestrtests

import (
	"errors"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_SrcC12_Collection_AddIfFuncErr_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_AddIfFuncErr_Verification", func() {
		// Arrange
		tc := srcC12CollectionAddIfFuncErrTestCase

		// Act
		cT := corestr.New.Collection.Cap(5)
		cT.AddIf(true, "a")
		cF := corestr.New.Collection.Cap(5)
		cF.AddIf(false, "a")
		cOk := corestr.New.Collection.Cap(5)
		cOk.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
		cFail := corestr.New.Collection.Cap(5)
		errHandled := false
		cFail.AddFuncErr(func() (string, error) { return "", errors.New("fail") }, func(e error) { errHandled = true })
		cErr := corestr.New.Collection.Cap(5)
		cErr.AddError(errors.New("test-err"))
		actual := args.Map{
			"addIfTrue":   cT.Length(),
			"addIfFalse":  cF.Length(),
			"funcErrOk":   cOk.Length(),
			"funcErrFail": cFail.Length(),
			"errHandled":  errHandled,
			"addErrVal":   cErr.First(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_HashmapMethods_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_HashmapMethods_Verification", func() {
		// Arrange
		tc := srcC12CollectionHashmapMethodsTestCase
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})

		// Act
		cV := corestr.New.Collection.Cap(5)
		cV.AddHashmapsValues(hm)
		cVN := corestr.New.Collection.Cap(5)
		cVN.AddHashmapsValues(nil)
		cVNN := corestr.New.Collection.Cap(5)
		cVNN.AddHashmapsValues(nil, nil)
		cK := corestr.New.Collection.Cap(5)
		cK.AddHashmapsKeys(hm)
		cKN := corestr.New.Collection.Cap(5)
		cKN.AddHashmapsKeys(nil)
		cKV := corestr.New.Collection.Cap(5)
		cKV.AddHashmapsKeysValues(hm)
		cKVN := corestr.New.Collection.Cap(5)
		cKVN.AddHashmapsKeysValues(nil)
		cF := corestr.New.Collection.Cap(5)
		cF.AddHashmapsKeysValuesUsingFilter(func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Key + "=" + pair.Value, true, false
		}, hm)
		cFN := corestr.New.Collection.Cap(5)
		cFN.AddHashmapsKeysValuesUsingFilter(nil, nil)
		hm2 := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"}, corestr.KeyValuePair{Key: "b", Value: "2"})
		cFB := corestr.New.Collection.Cap(5)
		cFB.AddHashmapsKeysValuesUsingFilter(func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Key, true, true
		}, hm2)
		actual := args.Map{
			"hmValsLen":     cV.Length(),
			"hmValsNil":     cVN.Length(),
			"hmValsNilNil":  cVNN.Length(),
			"hmKeysLen":     cK.Length(),
			"hmKeysNil":     cKN.Length(),
			"hmKVLen":       cKV.Length(),
			"hmKVNil":       cKVN.Length(),
			"hmFilterLen":   cF.Length(),
			"hmFilterNil":   cFN.Length(),
			"hmFilterBreak": cFB.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_WgLock_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_WgLock_Verification", func() {
		// Arrange
		tc := srcC12CollectionWgLockTestCase
		c := corestr.New.Collection.Cap(5)

		// Act
		wg := sync.WaitGroup{}
		wg.Add(1)
		c.AddWithWgLock(&wg, "x")
		wg.Wait()
		actual := args.Map{
			"length": c.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_Index_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_Index_Verification", func() {
		// Arrange
		tc := srcC12CollectionIndexTestCase
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		cXY := corestr.New.Collection.Strings([]string{"x", "y"})

		// Act
		actual := args.Map{
			"indexAt":      c.IndexAt(1),
			"safeInRange":  corestr.New.Collection.Strings([]string{"a"}).SafeIndexAtUsingLength("def", 1, 0),
			"safeOutRange": corestr.New.Collection.Strings([]string{"a"}).SafeIndexAtUsingLength("def", 1, 5),
			"first":        cXY.First(),
			"last":         cXY.Last(),
			"lastOrDefE":   corestr.New.Collection.Cap(0).LastOrDefault(),
			"lastOrDef":    corestr.New.Collection.Strings([]string{"a"}).LastOrDefault(),
			"firstOrDefE":  corestr.New.Collection.Cap(0).FirstOrDefault(),
			"firstOrDef":   corestr.New.Collection.Strings([]string{"z"}).FirstOrDefault(),
			"single":       corestr.New.Collection.Strings([]string{"only"}).Single(),
			"singlePanics": callPanicsSrcC12(func() { corestr.New.Collection.Strings([]string{"a", "b"}).Single() }),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_TakeSkip_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_TakeSkip_Verification", func() {
		// Arrange
		tc := srcC12CollectionTakeSkipTestCase

		// Act
		actual := args.Map{
			"take2":      corestr.New.Collection.Strings([]string{"a", "b", "c"}).Take(2).Length(),
			"takeMore":   corestr.New.Collection.Strings([]string{"a"}).Take(5).Length(),
			"takeZero":   corestr.New.Collection.Strings([]string{"a"}).Take(0).Length(),
			"skip1":      corestr.New.Collection.Strings([]string{"a", "b", "c"}).Skip(1).Length(),
			"skipZero":   corestr.New.Collection.Strings([]string{"a"}).Skip(0).Length(),
			"skipPanics": callPanicsSrcC12(func() { corestr.New.Collection.Strings([]string{"a"}).Skip(5) }),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_Reverse_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_Reverse_Verification", func() {
		// Arrange
		tc := srcC12CollectionReverseTestCase
		c3 := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c2 := corestr.New.Collection.Strings([]string{"a", "b"})
		c1 := corestr.New.Collection.Strings([]string{"a"})

		// Act
		c3.Reverse()
		c2.Reverse()
		c1.Reverse()
		actual := args.Map{
			"rev3First": c3.First(),
			"rev3Last":  c3.Last(),
			"rev2First": c2.First(),
			"rev1First": c1.First(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_Paging_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_Paging_Verification", func() {
		// Arrange
		tc := srcC12CollectionPagingTestCase
		items := make([]string, 10)
		for i := range items {
			items[i] = "x"
		}
		c := corestr.New.Collection.Strings(items)
		small := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"pagesSize":     c.GetPagesSize(2),
			"pagesZero":     c.GetPagesSize(0),
			"pagesNeg":      c.GetPagesSize(-1),
			"pagedLen":      c.GetPagedCollection(3).Length(),
			"pagedSmall":    small.GetPagedCollection(5).Length(),
			"singlePageLen": c.GetSinglePageCollection(3, 2).Length(),
			"lastPageLen":   c.GetSinglePageCollection(3, 4).Length(),
			"smallPageLen":  small.GetSinglePageCollection(5, 1).Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_InsertRemove_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_InsertRemove_Verification", func() {
		// Arrange
		tc := srcC12CollectionInsertRemoveTestCase

		// Act
		cI := corestr.New.Collection.Cap(5)
		cI.InsertAt(0, "a")
		cCR := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		cCR.ChainRemoveAt(1)
		cRI := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		cRI.RemoveItemsIndexes(true, 1)
		cRN := corestr.New.Collection.Strings([]string{"a"})
		cRN.RemoveItemsIndexes(true)
		actual := args.Map{
			"insertLen":      cI.Length(),
			"chainRemoveLen": cCR.Length(),
			"removeIdxLen":   cRI.Length(),
			"removeIdxNoop":  cRN.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_Append_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_Append_Verification", func() {
		// Arrange
		tc := srcC12CollectionAppendTestCase

		// Act
		cP := corestr.New.Collection.Strings([]string{"a"})
		cP.AppendCollectionPtr(corestr.New.Collection.Strings([]string{"b"}))
		cC := corestr.New.Collection.Cap(5)
		cC.AppendCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"}))
		cE := corestr.New.Collection.Cap(5)
		cE.AppendCollections()
		actual := args.Map{
			"appendPtrLen":  cP.Length(),
			"appendColsLen": cC.Length(),
			"appendEmpty":   cE.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_AppendAnys_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_AppendAnys_Verification", func() {
		// Arrange
		tc := srcC12CollectionAppendAnysTestCase
		f := func(str string, i int) (string, bool, bool) { return str, true, false }
		fSkip := func(str string, i int) (string, bool, bool) { return str, false, false }
		fBreak := func(str string, i int) (string, bool, bool) { return str, true, true }

		// Act
		cA := corestr.New.Collection.Cap(5)
		cA.AppendAnys(42, "hello", nil)
		cAE := corestr.New.Collection.Cap(5)
		cAE.AppendAnys()
		cAL := corestr.New.Collection.Cap(5)
		cAL.AppendAnysLock(42)
		cALE := corestr.New.Collection.Cap(5)
		cALE.AppendAnysLock()
		cF := corestr.New.Collection.Cap(5)
		cF.AppendAnysUsingFilter(f, "a", "b")
		cFS := corestr.New.Collection.Cap(5)
		cFS.AppendAnysUsingFilter(fSkip, "a")
		cFB := corestr.New.Collection.Cap(5)
		cFB.AppendAnysUsingFilter(fBreak, "a", "b")
		cFE := corestr.New.Collection.Cap(5)
		cFE.AppendAnysUsingFilter(nil)
		cFN := corestr.New.Collection.Cap(5)
		cFN.AppendAnysUsingFilter(f, nil)
		cFL := corestr.New.Collection.Cap(5)
		cFL.AppendAnysUsingFilterLock(f, "a")
		cFLN := corestr.New.Collection.Cap(5)
		cFLN.AppendAnysUsingFilterLock(nil)
		cFLB := corestr.New.Collection.Cap(5)
		cFLB.AppendAnysUsingFilterLock(fBreak, "a", "b")
		cFLS := corestr.New.Collection.Cap(5)
		cFLS.AppendAnysUsingFilterLock(fSkip, "a")
		cFLNI := corestr.New.Collection.Cap(5)
		cFLNI.AppendAnysUsingFilterLock(f, nil)
		cNE := corestr.New.Collection.Cap(5)
		cNE.AppendNonEmptyAnys(42, nil)
		cNEN := corestr.New.Collection.Cap(5)
		cNEN.AppendNonEmptyAnys(nil)
		cAN := corestr.New.Collection.Cap(5)
		cAN.AddsNonEmpty("a", "", "b")
		cANN := corestr.New.Collection.Cap(5)
		cANN.AddsNonEmpty()
		actual := args.Map{
			"anysLen":         cA.Length(),
			"anysEmpty":       cAE.Length(),
			"anysLock":        cAL.Length(),
			"anysLockEmpty":   cALE.Length(),
			"filterLen":       cF.Length(),
			"filterSkip":      cFS.Length(),
			"filterBreak":     cFB.Length(),
			"filterEmpty":     cFE.Length(),
			"filterNil":       cFN.Length(),
			"filterLock":      cFL.Length(),
			"filterLockNil":   cFLN.Length(),
			"filterLockBreak": cFLB.Length(),
			"filterLockSkip":  cFLS.Length(),
			"filterLockNilI":  cFLNI.Length(),
			"nonEmptyLen":     cNE.Length(),
			"nonEmptyNil":     cNEN.Length(),
			"addsNonEmpty":    cAN.Length(),
			"addsNonEmptyNil": cANN.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_PtrLock_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_PtrLock_Verification", func() {
		// Arrange
		tc := srcC12CollectionPtrLockTestCase
		s := "hello"
		empty := ""

		// Act
		c := corestr.New.Collection.Cap(5)
		c.AddsNonEmptyPtrLock(&s, nil, &empty)
		cN := corestr.New.Collection.Cap(5)
		cN.AddsNonEmptyPtrLock()
		actual := args.Map{
			"ptrLockLen": c.Length(),
			"ptrLockNil": cN.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_Unique_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_Unique_Verification", func() {
		// Arrange
		tc := srcC12CollectionUniqueTestCase

		// Act
		actual := args.Map{
			"boolMapLen":     len(corestr.New.Collection.Strings([]string{"a", "b", "a"}).UniqueBoolMap()),
			"boolMapLockLen": len(corestr.New.Collection.Strings([]string{"a", "a"}).UniqueBoolMapLock()),
			"uniqueLen":      len(corestr.New.Collection.Strings([]string{"a", "b", "a"}).UniqueList()),
			"uniqueLockLen":  len(corestr.New.Collection.Strings([]string{"a", "b", "a"}).UniqueListLock()),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_FilterDeep_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_FilterDeep_Verification", func() {
		// Arrange
		tc := srcC12CollectionFilterDeepTestCase
		f := func(str string, i int) (string, bool, bool) { return str, str != "b", false }
		fAll := func(str string, i int) (string, bool, bool) { return str, true, false }
		fBreak := func(str string, i int) (string, bool, bool) { return str, true, i == 0 }
		fP := func(sp *string, i int) (*string, bool, bool) { return sp, true, false }
		fPB := func(sp *string, i int) (*string, bool, bool) { return sp, true, i == 0 }

		// Act
		actual := args.Map{
			"filterLen":       len(corestr.New.Collection.Strings([]string{"a", "b", "c"}).Filter(f)),
			"filterEmpty":     len(corestr.New.Collection.Cap(0).Filter(fAll)),
			"filterBreak":     len(corestr.New.Collection.Strings([]string{"a", "b", "c"}).Filter(fBreak)),
			"filterLock":      len(corestr.New.Collection.Strings([]string{"a", "b"}).FilterLock(fAll)),
			"filteredCol":     corestr.New.Collection.Strings([]string{"a", "b"}).FilteredCollection(fAll).Length(),
			"filteredColLock": corestr.New.Collection.Strings([]string{"a", "b"}).FilteredCollectionLock(fAll).Length(),
			"filterPtr":       len(*corestr.New.Collection.Strings([]string{"a", "b"}).FilterPtr(fP)),
			"filterPtrEmpty":  len(*corestr.New.Collection.Cap(0).FilterPtr(fP)),
			"filterPtrBreak":  len(*corestr.New.Collection.Strings([]string{"a", "b"}).FilterPtr(fPB)),
			"filterPtrLock":   len(*corestr.New.Collection.Strings([]string{"a", "b"}).FilterPtrLock(fP)),
			"ptrLockEmpty":    len(*corestr.New.Collection.Cap(0).FilterPtrLock(fP)),
			"ptrLockBreak":    len(*corestr.New.Collection.Strings([]string{"a", "b"}).FilterPtrLock(fPB)),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_NonEmptyDeep_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_NonEmptyDeep_Verification", func() {
		// Arrange
		tc := srcC12CollectionNonEmptyDeepTestCase

		// Act
		actual := args.Map{
			"nonEmptyList":     len(corestr.New.Collection.Strings([]string{"a", "", "b"}).NonEmptyList()),
			"nonEmptyListE":    len(corestr.New.Collection.Cap(0).NonEmptyList()),
			"nonEmptyListPtr":  len(*corestr.New.Collection.Strings([]string{"a", ""}).NonEmptyListPtr()),
			"hashsetAsIs":      !corestr.New.Collection.Strings([]string{"a", "b"}).HashsetAsIs().IsEmpty(),
			"hashsetDouble":    !corestr.New.Collection.Strings([]string{"a"}).HashsetWithDoubleLength().IsEmpty(),
			"hashsetLock":      !corestr.New.Collection.Strings([]string{"a"}).HashsetLock().IsEmpty(),
			"nonEmptyItems":    len(corestr.New.Collection.Strings([]string{"a", "", "b"}).NonEmptyItems()),
			"nonEmptyItemsPtr": len(corestr.New.Collection.Strings([]string{"a", ""}).NonEmptyItemsPtr()),
			"nonEmptyWS":       len(corestr.New.Collection.Strings([]string{"a", "  ", "b"}).NonEmptyItemsOrNonWhitespace()),
			"nonEmptyWSPtr":    len(corestr.New.Collection.Strings([]string{"a", "  "}).NonEmptyItemsOrNonWhitespacePtr()),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_HasDeep_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_HasDeep_Verification", func() {
		// Arrange
		tc := srcC12CollectionHasDeepTestCase
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		a := "a"
		cH := corestr.New.Collection.Strings([]string{"Hello"})

		// Act
		actual := args.Map{
			"has":             c.Has("a"),
			"hasMiss":         c.Has("z"),
			"hasEmpty":        corestr.New.Collection.Cap(0).Has("a"),
			"hasPtrA":         c.HasPtr(&a),
			"hasPtrNil":       c.HasPtr(nil),
			"hasPtrEmpty":     corestr.New.Collection.Cap(0).HasPtr(&a),
			"hasAll":          corestr.New.Collection.Strings([]string{"a", "b", "c"}).HasAll("a", "b"),
			"hasAllMiss":      corestr.New.Collection.Strings([]string{"a", "b", "c"}).HasAll("a", "z"),
			"hasAllEmpty":     corestr.New.Collection.Cap(0).HasAll("a"),
			"hasLock":         corestr.New.Collection.Strings([]string{"a"}).HasLock("a"),
			"sensCaseTrue":    cH.HasUsingSensitivity("Hello", true),
			"sensCaseFalse":   cH.HasUsingSensitivity("hello", true),
			"sensInsensitive": cH.HasUsingSensitivity("hello", false),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_ContainsExcept_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_ContainsExcept_Verification", func() {
		// Arrange
		tc := srcC12CollectionContainsExceptTestCase
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c3 := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		s := "a"

		// Act
		hs, hsOk := c.GetHashsetPlusHasAll([]string{"a", "b"})
		_, hsNilOk := corestr.New.Collection.Strings([]string{"a"}).GetHashsetPlusHasAll(nil)
		_ = hs
		actual := args.Map{
			"containsPtr":       c.IsContainsPtr(&s),
			"containsPtrNil":    c.IsContainsPtr(nil),
			"containsAllSlice":  c.IsContainsAllSlice([]string{"a", "b"}),
			"containsAllSliceF": c.IsContainsAllSlice([]string{"z"}),
			"containsAllSliceE": c.IsContainsAllSlice([]string{}),
			"containsAll":       c.IsContainsAll("a", "b"),
			"containsAllLock":   corestr.New.Collection.Strings([]string{"a"}).IsContainsAllLock("a"),
			"hashsetHasAll":     hsOk,
			"hashsetHasAllNil":  hsNilOk,
			"exceptColLen":      len(c3.GetAllExceptCollection(corestr.New.Collection.Strings([]string{"b"}))),
			"exceptColNil":      len(corestr.New.Collection.Strings([]string{"a"}).GetAllExceptCollection(nil)),
			"exceptLen":         len(c3.GetAllExcept([]string{"a"})),
			"exceptNil":         len(corestr.New.Collection.Strings([]string{"a"}).GetAllExcept(nil)),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_NewAddNonEmpty_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_NewAddNonEmpty_Verification", func() {
		// Arrange
		tc := srcC12CollectionNewAddNonEmptyTestCase

		// Act
		actual := args.Map{
			"newLen":          corestr.New.Collection.Cap(0).New("a", "b").Length(),
			"newEmpty":        corestr.New.Collection.Cap(0).New().Length(),
			"addNonEmptyStr": func() int { c := corestr.New.Collection.Cap(5); c.AddNonEmptyStrings("a", "b"); return c.Length() }(),
			"addNonEmptyStrE": func() int { c := corestr.New.Collection.Cap(5); c.AddNonEmptyStrings(); return c.Length() }(),
			"funcResultLen": func() int {
				c := corestr.New.Collection.Cap(5)
				c.AddFuncResult(func() string { return "a" }, func() string { return "b" })
				return c.Length()
			}(),
			"funcResultNil": func() int { c := corestr.New.Collection.Cap(5); c.AddFuncResult(); return c.Length() }(),
			"funcCheckLen": func() int {
				c := corestr.New.Collection.Cap(5)
				c.AddStringsByFuncChecking([]string{"ok", "bad", "ok2"}, func(l string) bool { return l != "bad" })
				return c.Length()
			}(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_ExpandMergeChar_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_ExpandMergeChar_Verification", func() {
		// Arrange
		tc := srcC12CollectionExpandMergeCharTestCase

		// Act
		cE := corestr.New.Collection.Cap(5)
		cE.ExpandSlicePlusAdd([]string{"a,b", "c,d"}, func(l string) []string { return []string{l + "_exp"} })
		cM := corestr.New.Collection.Cap(5)
		cM.MergeSlicesOfSlice([]string{"a"}, []string{"b"})
		actual := args.Map{
			"expandLen": cE.Length(),
			"mergeLen":  cM.Length(),
			"charMapNN": corestr.New.Collection.Strings([]string{"abc", "def"}).CharCollectionMap() != nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_StringDeep_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_StringDeep_Verification", func() {
		// Arrange
		tc := srcC12CollectionStringDeepTestCase

		// Act
		noPanic := !callPanicsSrcC12(func() {
			c := corestr.New.Collection.Strings([]string{"a"})
			_ = c.String()
			_ = corestr.New.Collection.Cap(0).String()
			_ = c.StringLock()
			_ = corestr.New.Collection.Cap(0).StringLock()
			_ = c.SummaryString(1)
			_ = c.SummaryStringWithHeader("header:")
			_ = corestr.New.Collection.Cap(0).SummaryStringWithHeader("header:")
			_ = corestr.New.Collection.Strings([]string{"a", "b"}).Csv()
			_ = corestr.New.Collection.Cap(0).Csv()
			_ = corestr.New.Collection.Strings([]string{"a"}).CsvOptions(true)
			_ = corestr.New.Collection.Cap(0).CsvOptions(false)
			_ = corestr.New.Collection.Strings([]string{"a"}).CsvLines()
			_ = corestr.New.Collection.Strings([]string{"a"}).CsvLinesOptions(true)
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_JsonDeep_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_JsonDeep_Verification", func() {
		// Arrange
		tc := srcC12CollectionJsonDeepTestCase

		// Act
		noPanic := !callPanicsSrcC12(func() {
			c := corestr.New.Collection.Strings([]string{"a"})
			_ = c.JsonModel()
			_ = c.JsonModelAny()
			_, _ = c.MarshalJSON()
			var out corestr.Collection
			_ = out.UnmarshalJSON([]byte(`["a","b"]`))
			var out2 corestr.Collection
			_ = out2.UnmarshalJSON([]byte(`invalid`))
			_ = c.Json()
			_ = c.JsonPtr()
			jr := c.Json()
			c2 := &corestr.Collection{}
			_, _ = c2.ParseInjectUsingJson(&jr)
			c3 := &corestr.Collection{}
			_ = c3.ParseInjectUsingJsonMust(&jr)
			c4 := &corestr.Collection{}
			_ = c4.JsonParseSelfInject(&jr)
			_, _ = c.Serialize()
			var target []string
			_ = c.Deserialize(&target)
			_ = c.AsJsonMarshaller()
			_ = c.AsJsonContractsBinder()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_ClearDisposeDeep_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_ClearDisposeDeep_Verification", func() {
		// Arrange
		tc := srcC12CollectionClearDisposeDeepTestCase
		var nilC *corestr.Collection

		// Act
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Clear()
		clearLen := c.Length()
		noPanic := !callPanicsSrcC12(func() {
			corestr.New.Collection.Strings([]string{"a"}).Dispose()
			nilC.Dispose()
		})
		actual := args.Map{
			"clearLen":   clearLen,
			"nilClear":   nilC.Clear() == nil,
			"nilDispose": true,
			"noPanic":    noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_Misc_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_Misc_Verification", func() {
		// Arrange
		tc := srcC12CollectionMiscTestCase

		// Act
		cPCL := corestr.New.Collection.Cap(5)
		cPCL.AddPointerCollectionsLock(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{
			"ptrColLock":    cPCL.Length(),
			"listCopy":      len(corestr.New.Collection.Strings([]string{"a"}).ListCopyPtrLock()),
			"listCopyEmpty": len(corestr.New.Collection.Cap(0).ListCopyPtrLock()),
			"itemsLen":      len(corestr.New.Collection.Strings([]string{"a"}).Items()),
			"listPtrLen":    len(corestr.New.Collection.Strings([]string{"a"}).ListPtr()),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_SortedDeep_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_SortedDeep_Verification", func() {
		// Arrange
		tc := srcC12CollectionSortedDeepTestCase

		// Act
		noPanic := !callPanicsSrcC12(func() {
			_ = corestr.New.Collection.Cap(0).SortedAsc()
			_ = corestr.New.Collection.Strings([]string{"c", "a"}).SortedAscLock()
			_ = corestr.New.Collection.Strings([]string{"a", "c", "b"}).SortedListDsc()
		})
		actual := args.Map{
			"ascFirst":    corestr.New.Collection.Strings([]string{"c", "a", "b"}).SortedListAsc()[0],
			"ascEmptyLen": len(corestr.New.Collection.Cap(0).SortedListAsc()),
			"noPanic":     noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_CapResize_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_CapResize_Verification", func() {
		// Arrange
		tc := srcC12CollectionCapResizeTestCase

		// Act
		c := corestr.New.Collection.Cap(5)
		c.AddCapacity(10)
		cR := corestr.New.Collection.Strings([]string{"a"})
		cR.Resize(100)
		noPanic := !callPanicsSrcC12(func() {
			corestr.New.Collection.Cap(5).AddCapacity()
			corestr.New.Collection.Cap(100).Resize(5)
		})
		actual := args.Map{
			"capGe10":  c.Capacity() >= 10,
			"noPanic":  noPanic,
			"resGe100": cR.Capacity() >= 100,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC12_Collection_JoinsDeep_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC12_Collection_JoinsDeep_Verification", func() {
		// Arrange
		tc := srcC12CollectionJoinsDeepTestCase

		// Act
		actual := args.Map{
			"joins":         corestr.New.Collection.Strings([]string{"a", "b"}).Joins(",") != "",
			"joinsExtra":    corestr.New.Collection.Strings([]string{"a"}).Joins(",", "b", "c") != "",
			"nonEmptyJoins": corestr.New.Collection.Strings([]string{"a", "", "b"}).NonEmptyJoins(",") != "",
			"nonWSJoins":    corestr.New.Collection.Strings([]string{"a", "  ", "b"}).NonWhitespaceJoins(",") != "",
			"joinAB":        corestr.New.Collection.Strings([]string{"a", "b"}).Join(","),
			"joinEmpty":     corestr.New.Collection.Cap(0).Join(","),
			"joinLine":      corestr.New.Collection.Strings([]string{"a", "b"}).JoinLine() != "",
			"joinLineEmpty": corestr.New.Collection.Cap(0).JoinLine(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC12(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
	}
