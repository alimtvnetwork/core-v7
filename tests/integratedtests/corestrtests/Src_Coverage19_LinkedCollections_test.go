package corestrtests

import (
	"encoding/json"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_SrcC19_Creators_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_Creators_Verification", func() {
		// Arrange
		tc := srcC19CreatorsTestCase

		// Act
		s1, s2 := "a", "b"
		items := []*string{&s1, &s2}
		actual := args.Map{
			"createNN":    corestr.New.LinkedCollection.Create() != nil,
			"createEmpty": !corestr.New.LinkedCollection.Create().HasItems(),
			"emptyLen":    corestr.New.LinkedCollection.Empty().Length(),
			"stringsLen":  corestr.New.LinkedCollection.Strings("a", "b").Length(),
			"stringsEE":   !corestr.New.LinkedCollection.Strings().HasItems(),
			"ptrLen":      corestr.New.LinkedCollection.PointerStringsPtr(&items).Length(),
			"ptrNilE":     !corestr.New.LinkedCollection.PointerStringsPtr(nil).HasItems(),
			"usingLen":    corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"})).Length(),
			"usingNilNN":  corestr.New.LinkedCollection.UsingCollections(nil) != nil,
			"emptyLCe":    !corestr.Empty.LinkedCollections().HasItems(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC19_HeadTail_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_HeadTail_Verification", func() {
		// Arrange
		tc := srcC19HeadTailTestCase
		lc := corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"}))

		// Act
		actual := args.Map{
			"headNN":     lc.Head() != nil,
			"tailNN":     lc.Tail() != nil,
			"firstLen":   lc.First().Length(),
			"lastLen":    lc.Last().Length(),
			"singleLen":  lc.Single().Length(),
			"fodEmptyNN": corestr.Empty.LinkedCollections().FirstOrDefault() != nil,
			"lodEmptyNN": corestr.Empty.LinkedCollections().LastOrDefault() != nil,
			"fodHasLen":  corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"x"})).FirstOrDefault().Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC19_Length_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_Length_Verification", func() {
		// Arrange
		tc := srcC19LengthTestCase

		// Act
		actual := args.Map{
			"lengthLock": corestr.New.LinkedCollection.Strings("a").LengthLock(),
			"allItems":   corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a", "b"}), corestr.New.Collection.Strings([]string{"c"})).AllIndividualItemsLength(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC19_State_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_State_Verification", func() {
		// Arrange
		tc := srcC19StateTestCase

		// Act
		actual := args.Map{
			"emptyIsEmpty": corestr.Empty.LinkedCollections().IsEmpty(),
			"emptyHasIt":   corestr.Empty.LinkedCollections().HasItems(),
			"emptyLock":    corestr.Empty.LinkedCollections().IsEmptyLock(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC19_Add_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_Add_Verification", func() {
		// Arrange
		tc := srcC19AddTestCase

		// Act
		lc1 := corestr.New.LinkedCollection.Create(); lc1.Add(corestr.New.Collection.Strings([]string{"a"})); lc1.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc2 := corestr.New.LinkedCollection.Create(); lc2.AddLock(corestr.New.Collection.Strings([]string{"a"}))
		lc3 := corestr.New.LinkedCollection.Create(); lc3.AddStrings("a", "b")
		lc4 := corestr.New.LinkedCollection.Create(); lc4.AddStrings()
		lc5 := corestr.New.LinkedCollection.Create(); lc5.AddStringsLock("a")
		lc6 := corestr.New.LinkedCollection.Create(); lc6.AddStringsLock()
		actual := args.Map{
			"addLen1":       1, // after first add
			"addLen2":       lc1.Length(),
			"addLockLen":    lc2.Length(),
			"addStrLen":     lc3.Length(),
			"addStrEE":      !lc4.HasItems(),
			"addStrLockLen": lc5.Length(),
			"addStrLockEE":  !lc6.HasItems(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC19_AddFront_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_AddFront_Verification", func() {
		// Arrange
		tc := srcC19AddFrontTestCase

		// Act
		lc1 := corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"b"}))
		lc1.AddFront(corestr.New.Collection.Strings([]string{"a"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddFront(corestr.New.Collection.Strings([]string{"a"}))
		lc3 := corestr.New.LinkedCollection.Strings("b")
		lc3.AddFrontLock(corestr.New.Collection.Strings([]string{"a"}))
		lc4 := corestr.New.LinkedCollection.Strings("b")
		lc4.PushFront(corestr.New.Collection.Strings([]string{"a"}))
		lc5 := corestr.New.LinkedCollection.Create()
		lc5.PushBack(corestr.New.Collection.Strings([]string{"a"}))
		lc6 := corestr.New.LinkedCollection.Create()
		lc6.PushBackLock(corestr.New.Collection.Strings([]string{"a"}))
		lc7 := corestr.New.LinkedCollection.Create()
		lc7.Push(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{
			"frontFirst":    lc1.First().List()[0],
			"frontEmptyLen": lc2.Length(),
			"frontLockLen":  lc3.Length(),
			"pushFrontLen":  lc4.Length(),
			"pushBackLen":   lc5.Length(),
			"pushBackLkLen": lc6.Length(),
			"pushLen":       lc7.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC19_AppendNode_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_AppendNode_Verification", func() {
		// Arrange
		tc := srcC19AppendNodeTestCase

		// Act
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AppendNode(&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})})
		lc1.AppendNode(&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})})
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddBackNode(&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})})
		actual := args.Map{
			"appendLen2": lc1.Length(),
			"addBackLen": lc2.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC19_InsertAt_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_InsertAt_Verification", func() {
		// Arrange
		tc := srcC19InsertAtTestCase

		// Act
		lc1 := corestr.New.LinkedCollection.Strings("b")
		lc1.InsertAt(0, corestr.New.Collection.Strings([]string{"a"}))
		lc2 := corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"c"}))
		lc2.InsertAt(1, corestr.New.Collection.Strings([]string{"b"}))
		actual := args.Map{
			"frontFirst": lc1.First().List()[0],
			"middleLen":  lc2.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC19_Attach_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_Attach_Verification", func() {
		// Arrange
		tc := srcC19AttachTestCase

		// Act
		lc1 := corestr.New.LinkedCollection.Create()
		err1 := lc1.AttachWithNode(nil, &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"x"})})
		lc2 := corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"}))
		err2 := lc2.AttachWithNode(lc2.Head(), &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"x"})})
		lc3 := corestr.New.LinkedCollection.Strings("a")
		err3 := lc3.AttachWithNode(lc3.Tail(), &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})})
		actual := args.Map{
			"nilCurrErr":    err1 != nil,
			"nonNilNextErr": err2 != nil,
			"successOk":     err3 == nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC19_AddAnotherColl_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_AddAnotherColl_Verification", func() {
		// Arrange
		tc := srcC19AddAnotherCollTestCase

		// Act
		lc1 := corestr.New.LinkedCollection.Strings("a"); lc1.AddAnother(corestr.New.LinkedCollection.Strings("b"))
		lc2 := corestr.New.LinkedCollection.Strings("a"); lc2.AddAnother(nil)
		lc3 := corestr.New.LinkedCollection.Create(); lc3.AddCollection(corestr.New.Collection.Strings([]string{"a"}))
		lc4 := corestr.New.LinkedCollection.Create(); lc4.AddCollection(nil)
		lc5 := corestr.New.LinkedCollection.Create(); lc5.AddCollectionsPtr([]*corestr.Collection{corestr.New.Collection.Strings([]string{"a"})})
		lc6 := corestr.New.LinkedCollection.Create(); lc6.AddCollections([]*corestr.Collection{corestr.New.Collection.Strings([]string{"a"})})
		lc7 := corestr.New.LinkedCollection.Create(); lc7.AddCollections([]*corestr.Collection{})
		lc8 := corestr.New.LinkedCollection.Create(); lc8.AppendCollections(true, corestr.New.Collection.Strings([]string{"a"}), nil)
		lc9 := corestr.New.LinkedCollection.Create(); lc9.AppendCollections(true, nil)
		cols := []*corestr.Collection{corestr.New.Collection.Strings([]string{"a"}), nil}
		lc10 := corestr.New.LinkedCollection.Create(); lc10.AppendCollectionsPointers(true, &cols)
		cols2 := []*corestr.Collection{corestr.New.Collection.Strings([]string{"a"})}
		lc11 := corestr.New.LinkedCollection.Create(); lc11.AppendCollectionsPointersLock(true, &cols2)
		actual := args.Map{
			"anotherLen":     lc1.Length(),
			"anotherNilLen":  lc2.Length(),
			"colLen":         lc3.Length(),
			"colNilE":        !lc4.HasItems(),
			"colsPtrLen":     lc5.Length(),
			"colsLen":        lc6.Length(),
			"colsEmptyE":     !lc7.HasItems(),
			"appendLen":      lc8.Length(),
			"appendNilE":     !lc9.HasItems(),
			"appendPtrLen":   lc10.Length(),
			"appendPtrLkLen": lc11.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC19_LoopFilter_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_LoopFilter_Verification", func() {
		// Arrange
		tc := srcC19LoopFilterTestCase

		// Act
		count1 := 0
		corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"})).Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count1++
			return false
		})
		count2 := 0
		corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"})).Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count2++
			return true
		})
		emptyOk := true
		corestr.New.LinkedCollection.Create().Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			emptyOk = false
			return false
		})
		nodes1 := corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"})).Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		nodes2 := corestr.New.LinkedCollection.Create().Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		col1 := corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a", "b"}), corestr.New.Collection.Strings([]string{"c"})).FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		}, 0)
		col2 := corestr.New.LinkedCollection.Create().FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		}, 0)
		cols := corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"})).FilterAsCollections(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		actual := args.Map{
			"loopCount":      count1,
			"loopBreak":      count2,
			"loopEmptyOk":    emptyOk,
			"filterLen":      len(nodes1),
			"filterEmptyLen": len(nodes2),
			"filterColLen":   col1.Length(),
			"filterColEE":    !col2.HasItems(),
			"filterColsLen":  len(cols),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC19_Remove_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_Remove_Verification", func() {
		// Arrange
		tc := srcC19RemoveTestCase

		// Act
		lc1 := corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"}), corestr.New.Collection.Strings([]string{"c"}))
		lc2 := corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"}))
		lc2.RemoveNodeByIndex(0)
		lc3 := corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"}))
		lc3.RemoveNodeByIndex(1)
		lc4 := corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"}), corestr.New.Collection.Strings([]string{"c"}))
		lc4.RemoveNodeByIndex(1)
		lc5 := corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"}), corestr.New.Collection.Strings([]string{"c"}))
		lc5.RemoveNodeByIndexes(true, 0, 2)
		lc6 := corestr.New.LinkedCollection.Strings("a")
		lc6.RemoveNodeByIndexes(true)
		lc7 := corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"}))
		lc7.RemoveNode(lc7.Head())
		lc8 := corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"}), corestr.New.Collection.Strings([]string{"c"}))
		lc8.RemoveNode(lc8.Head().Next())
		lc9 := corestr.New.LinkedCollection.Strings("a")
		lc9.AddAfterNode(lc9.Head(), corestr.New.Collection.Strings([]string{"b"}))
		actual := args.Map{
			"nextNodesLen":   len(lc1.GetNextNodes(2)),
			"allNodesLen":    len(corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"})).GetAllLinkedNodes()),
			"rmIdxFirstLen":  lc2.Length(),
			"rmIdxLastLen":   lc3.Length(),
			"rmIdxMidLen":    lc4.Length(),
			"rmIdxesLen":     lc5.Length(),
			"rmIdxesEmpLen":  lc6.Length(),
			"rmNodeFirstLen": lc7.Length(),
			"rmNodeMidLen":   lc8.Length(),
			"afterLen":       lc9.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC19_ConcatIndex_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_ConcatIndex_Verification", func() {
		// Arrange
		tc := srcC19ConcatIndexTestCase

		// Act
		lc := corestr.New.LinkedCollection.Strings("a")
		actual := args.Map{
			"concatLen":      corestr.New.LinkedCollection.Strings("a").ConcatNew(false, corestr.New.LinkedCollection.Strings("b")).Length(),
			"concatCloneLen": corestr.New.LinkedCollection.Strings("a").ConcatNew(true).Length(),
			"concatSameRef":  lc.ConcatNew(false) == lc,
			"idxAt1NN":       corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"})).IndexAt(1) != nil,
			"idxAt0NN":       corestr.New.LinkedCollection.Strings("a").IndexAt(0) != nil,
			"idxNegNil":      corestr.New.LinkedCollection.Strings("a").IndexAt(-1) == nil,
			"safeAt1NN":      corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"})).SafeIndexAt(1) != nil,
			"safeOorNil":     corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"})).SafeIndexAt(5) == nil,
			"safeNegNil":     corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"})).SafeIndexAt(-1) == nil,
			"ptrAt0NN":       corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"})).SafePointerIndexAt(0) != nil,
			"ptrOorNil":      corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"})).SafePointerIndexAt(5) == nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}
func Test_SrcC19_AddStrOfStrAsync_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_AddStrOfStrAsync_Verification", func() {
		// Arrange
		tc := srcC19AddStrOfStrAsyncTestCase

		// Act
		lc1 := corestr.New.LinkedCollection.Create(); lc1.AddStringsOfStrings(false, []string{"a"}, []string{"b"})
		lc2 := corestr.New.LinkedCollection.Create(); lc2.AddStringsOfStrings(false)
		lc3 := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc3.AddAsyncFuncItems(wg, false, func() []string { return []string{"a", "b"} })
		lc4 := corestr.New.LinkedCollection.Create()
		wg2 := &sync.WaitGroup{}
		wg2.Add(1)
		lc4.AddAsyncFuncItems(wg2, false, func() []string { return []string{} })
		lc5 := corestr.New.LinkedCollection.Create()
		wg3 := &sync.WaitGroup{}
		lc5.AddAsyncFuncItems(wg3, false)
		lc6 := corestr.New.LinkedCollection.Create()
		wg4 := &sync.WaitGroup{}
		wg4.Add(1)
		lc6.AddAsyncFuncItemsPointer(wg4, false, func() []string { return []string{"x"} })
		lc7 := corestr.New.LinkedCollection.Create()
		wg5 := &sync.WaitGroup{}
		lc7.AddAsyncFuncItemsPointer(wg5, false)
		actual := args.Map{
			"sosLen":       lc1.Length(),
			"sosEmptyE":    !lc2.HasItems(),
			"asyncLen":     lc3.Length(),
			"asyncEmptyE":  !lc4.HasItems(),
			"asyncNilE":    !lc5.HasItems(),
			"asyncPtrLen":  lc6.Length(),
			"asyncPtrNilE": !lc7.HasItems(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}
func Test_SrcC19_EqualsCompare_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_EqualsCompare_Verification", func() {
		// Arrange
		tc := srcC19EqualsCompareTestCase

		// Act
		lc := corestr.New.LinkedCollection.Strings("a")
		actual := args.Map{
			"equalSame":    corestr.New.LinkedCollection.Strings("a", "b").IsEqualsPtr(corestr.New.LinkedCollection.Strings("a", "b")),
			"equalNil":     corestr.New.LinkedCollection.Strings("a").IsEqualsPtr(nil),
			"equalSameRef": lc.IsEqualsPtr(lc),
			"equalBothE":   corestr.New.LinkedCollection.Create().IsEqualsPtr(corestr.New.LinkedCollection.Create()),
			"equalOneE":    corestr.New.LinkedCollection.Strings("a").IsEqualsPtr(corestr.New.LinkedCollection.Create()),
			"equalDiffLen": corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Strings([]string{"b"})).IsEqualsPtr(corestr.New.LinkedCollection.Strings("a")),
			"summaryNonE":  corestr.New.LinkedCollection.Strings("a").GetCompareSummary(corestr.New.LinkedCollection.Strings("b"), "left", "right") != "",
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC19_Json_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_Json_Verification", func() {
		// Arrange
		tc := srcC19JsonTestCase

		// Act
		noPanic := !callPanicsSrcC19(func() {
			lc := corestr.New.LinkedCollection.Strings("a", "b")
			_ = lc.JsonModel()
			_ = lc.JsonModelAny()
			data, _ := json.Marshal(lc)
			lc2 := corestr.New.LinkedCollection.Create()
			_ = json.Unmarshal(data, lc2)
			r := lc.Json()
			_ = r.Error == nil
			jr := lc.JsonPtr()
			lc3 := corestr.New.LinkedCollection.Create()
			_, _ = lc3.ParseInjectUsingJson(jr)
			lc4 := corestr.New.LinkedCollection.Create()
			_ = lc4.ParseInjectUsingJsonMust(jr)
			lc5 := corestr.New.LinkedCollection.Create()
			_ = lc5.JsonParseSelfInject(jr)
			_ = lc.AsJsonContractsBinder()
			_ = lc.AsJsoner()
			_ = lc.AsJsonParseSelfInjector()
			_ = lc.AsJsonMarshaller()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC19_Clear_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_Clear_Verification", func() {
		// Arrange
		tc := srcC19ClearTestCase

		// Act
		lc1 := corestr.New.LinkedCollection.Strings("a"); lc1.RemoveAll()
		lc2 := corestr.New.LinkedCollection.Strings("a"); lc2.Clear()
		lc3 := corestr.New.LinkedCollection.Create(); lc3.Clear()
		actual := args.Map{
			"removeAllE":  !lc1.HasItems(),
			"clearE":      !lc2.HasItems(),
			"clearEmptyOk": true,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}
func Test_SrcC19_NonChainedEmpty_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC19_NonChainedEmpty_Verification", func() {
		// Arrange
		tc := srcC19NonChainedEmptyTestCase

		// Act
		nc := &corestr.NonChainedLinkedCollectionNodes{}
		actual := args.Map{
			"fodNil": nc.FirstOrDefault() == nil,
			"lodNil": nc.LastOrDefault() == nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC19(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
	}
