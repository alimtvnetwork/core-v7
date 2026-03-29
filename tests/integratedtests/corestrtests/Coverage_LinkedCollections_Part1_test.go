package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// LinkedCollections — Segment 4: Basic ops, Add, Loop, Filter, Remove (L1-800)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovLC1_01_Tail_Head(t *testing.T) {
	safeTest(t, "Test_CovLC1_01_Tail_Head", func() {
		lc := corestr.Empty.LinkedCollections()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c1)
		if lc.Tail() == nil {
			t.Fatal("expected non-nil tail")
		}
		if lc.Head() == nil {
			t.Fatal("expected non-nil head")
		}
	})
}

func Test_CovLC1_02_First_Single_Last(t *testing.T) {
	safeTest(t, "Test_CovLC1_02_First_Single_Last", func() {
		lc := corestr.Empty.LinkedCollections()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c1)
		if lc.First().Length() != 1 {
			t.Fatal("expected 1")
		}
		if lc.Single().Length() != 1 {
			t.Fatal("expected 1")
		}
		if lc.Last().Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLC1_03_LastOrDefault_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_CovLC1_03_LastOrDefault_FirstOrDefault", func() {
		lc := corestr.Empty.LinkedCollections()
		if lc.LastOrDefault().Length() != 0 {
			t.Fatal("expected empty")
		}
		if lc.FirstOrDefault().Length() != 0 {
			t.Fatal("expected empty")
		}
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c1)
		if lc.LastOrDefault().Length() != 1 {
			t.Fatal("expected 1")
		}
		if lc.FirstOrDefault().Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLC1_04_Length(t *testing.T) {
	safeTest(t, "Test_CovLC1_04_Length", func() {
		lc := corestr.Empty.LinkedCollections()
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLC1_05_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_CovLC1_05_AllIndividualItemsLength", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		if lc.AllIndividualItemsLength() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_CovLC1_06_LengthLock(t *testing.T) {
	safeTest(t, "Test_CovLC1_06_LengthLock", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.LengthLock() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLC1_07_IsEqualsPtr(t *testing.T) {
	safeTest(t, "Test_CovLC1_07_IsEqualsPtr", func() {
		a := corestr.Empty.LinkedCollections()
		a.Add(corestr.New.Collection.Strings([]string{"x", "y"}))
		b := corestr.Empty.LinkedCollections()
		b.Add(corestr.New.Collection.Strings([]string{"x", "y"}))

		// same ptr
		if !a.IsEqualsPtr(a) {
			t.Fatal("expected equal to self")
		}
		// nil
		if a.IsEqualsPtr(nil) {
			t.Fatal("expected false for nil")
		}
		// both empty
		e1 := corestr.Empty.LinkedCollections()
		e2 := corestr.Empty.LinkedCollections()
		if !e1.IsEqualsPtr(e2) {
			t.Fatal("expected empty == empty")
		}
		// one empty
		if a.IsEqualsPtr(e1) {
			t.Fatal("expected false")
		}
		// diff length
		c := corestr.Empty.LinkedCollections()
		c.Add(corestr.New.Collection.Strings([]string{"x"}))
		c.Add(corestr.New.Collection.Strings([]string{"y"}))
		// same content different structure
		if a.IsEqualsPtr(b) != true {
			t.Fatal("expected equal")
		}
	})
}

func Test_CovLC1_08_IsEmptyLock_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_CovLC1_08_IsEmptyLock_IsEmpty_HasItems", func() {
		lc := corestr.Empty.LinkedCollections()
		if !lc.IsEmptyLock() {
			t.Fatal("expected empty")
		}
		if !lc.IsEmpty() {
			t.Fatal("expected empty")
		}
		if lc.HasItems() {
			t.Fatal("expected no items")
		}
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.IsEmpty() {
			t.Fatal("expected not empty")
		}
		if !lc.HasItems() {
			t.Fatal("expected has items")
		}
	})
}

func Test_CovLC1_09_InsertAt(t *testing.T) {
	safeTest(t, "Test_CovLC1_09_InsertAt", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		// insert at front
		lc.InsertAt(0, corestr.New.Collection.Strings([]string{"front"}))
		if lc.Length() != 3 {
			t.Fatal("expected 3")
		}
		// insert in middle
		lc.InsertAt(1, corestr.New.Collection.Strings([]string{"mid"}))
		if lc.Length() != 4 {
			t.Fatal("expected 4")
		}
	})
}

func Test_CovLC1_10_AddAsync(t *testing.T) {
	safeTest(t, "Test_CovLC1_10_AddAsync", func() {
		lc := corestr.Empty.LinkedCollections()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsync(corestr.New.Collection.Strings([]string{"a"}), wg)
		wg.Wait()
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLC1_11_AddsAsyncOnComplete(t *testing.T) {
	safeTest(t, "Test_CovLC1_11_AddsAsyncOnComplete", func() {
		lc := corestr.Empty.LinkedCollections()
		done := make(chan bool, 1)
		lc.AddsAsyncOnComplete(
			func(lc *corestr.LinkedCollections) { done <- true },
			true,
			corestr.New.Collection.Strings([]string{"a"}),
		)
		<-done
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLC1_12_AddsUsingProcessorAsyncOnComplete(t *testing.T) {
	safeTest(t, "Test_CovLC1_12_AddsUsingProcessorAsyncOnComplete", func() {
		lc := corestr.Empty.LinkedCollections()
		done := make(chan bool, 1)
		lc.AddsUsingProcessorAsyncOnComplete(
			func(lc *corestr.LinkedCollections) { done <- true },
			func(a any, i int) *corestr.Collection {
				return corestr.New.Collection.Strings([]string{a.(string)})
			},
			true,
			"hello", nil,
		)
		<-done
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}

		// nil anys with skip
		lc2 := corestr.Empty.LinkedCollections()
		done2 := make(chan bool, 1)
		lc2.AddsUsingProcessorAsyncOnComplete(
			func(lc *corestr.LinkedCollections) { done2 <- true },
			func(a any, i int) *corestr.Collection { return nil },
			true,
		)
		<-done2
	})
}

func Test_CovLC1_13_AddsUsingProcessorAsync(t *testing.T) {
	safeTest(t, "Test_CovLC1_13_AddsUsingProcessorAsync", func() {
		lc := corestr.Empty.LinkedCollections()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddsUsingProcessorAsync(
			wg,
			func(a any, i int) *corestr.Collection {
				return corestr.New.Collection.Strings([]string{a.(string)})
			},
			true,
			"x",
		)
		wg.Wait()
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}

		// nil anys with skip
		lc2 := corestr.Empty.LinkedCollections()
		wg2 := &sync.WaitGroup{}
		wg2.Add(1)
		lc2.AddsUsingProcessorAsync(wg2,
			func(a any, i int) *corestr.Collection { return nil },
			true,
		)
		wg2.Wait()
	})
}

func Test_CovLC1_14_AddLock(t *testing.T) {
	safeTest(t, "Test_CovLC1_14_AddLock", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.AddLock(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLC1_15_Add(t *testing.T) {
	safeTest(t, "Test_CovLC1_15_Add", func() {
		lc := corestr.Empty.LinkedCollections()
		// first add sets head
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		// second add sets tail
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLC1_16_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_CovLC1_16_AddStringsLock", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.AddStringsLock("a", "b")
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
		// empty
		lc.AddStringsLock()
	})
}

func Test_CovLC1_17_AddStrings(t *testing.T) {
	safeTest(t, "Test_CovLC1_17_AddStrings", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.AddStrings("a", "b")
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
		lc.AddStrings()
	})
}

func Test_CovLC1_18_AddBackNode_AppendNode(t *testing.T) {
	safeTest(t, "Test_CovLC1_18_AddBackNode_AppendNode", func() {
		lc := corestr.Empty.LinkedCollections()
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		lc.AddBackNode(node)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
		// append to non-empty
		node2 := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		lc.AppendNode(node2)
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLC1_19_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_CovLC1_19_AppendChainOfNodes", func() {
		lc := corestr.Empty.LinkedCollections()
		node1 := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		// empty list
		lc.AppendChainOfNodes(node1)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
		// non-empty
		node2 := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		lc.AppendChainOfNodes(node2)
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLC1_20_AppendChainOfNodesAsync(t *testing.T) {
	safeTest(t, "Test_CovLC1_20_AppendChainOfNodesAsync", func() {
		lc := corestr.Empty.LinkedCollections()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		lc.AppendChainOfNodesAsync(node, wg)
		wg.Wait()
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLC1_21_PushBackLock_PushBack_Push_PushFront(t *testing.T) {
	safeTest(t, "Test_CovLC1_21_PushBackLock_PushBack_Push_PushFront", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.PushBackLock(corestr.New.Collection.Strings([]string{"a"}))
		lc.PushBack(corestr.New.Collection.Strings([]string{"b"}))
		lc.Push(corestr.New.Collection.Strings([]string{"c"}))
		lc.PushFront(corestr.New.Collection.Strings([]string{"front"}))
		if lc.Length() != 4 {
			t.Fatal("expected 4")
		}
	})
}

func Test_CovLC1_22_AddFrontLock_AddFront(t *testing.T) {
	safeTest(t, "Test_CovLC1_22_AddFrontLock_AddFront", func() {
		lc := corestr.Empty.LinkedCollections()
		// empty — falls through to Add
		lc.AddFront(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
		// non-empty
		lc.AddFrontLock(corestr.New.Collection.Strings([]string{"front"}))
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLC1_23_AddAnother(t *testing.T) {
	safeTest(t, "Test_CovLC1_23_AddAnother", func() {
		a := corestr.Empty.LinkedCollections()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		b := corestr.Empty.LinkedCollections()
		b.Add(corestr.New.Collection.Strings([]string{"b"}))
		b.Add(corestr.New.Collection.Strings([]string{"c"}))
		a.AddAnother(b)
		if a.Length() != 3 {
			t.Fatal("expected 3")
		}
		// nil
		a.AddAnother(nil)
		// empty
		a.AddAnother(corestr.Empty.LinkedCollections())
	})
}

func Test_CovLC1_24_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_CovLC1_24_GetNextNodes", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		nodes := lc.GetNextNodes(2)
		if len(nodes) != 2 {
			t.Fatalf("expected 2, got %d", len(nodes))
		}
	})
}

func Test_CovLC1_25_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_CovLC1_25_GetAllLinkedNodes", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		nodes := lc.GetAllLinkedNodes()
		if len(nodes) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLC1_26_Loop(t *testing.T) {
	safeTest(t, "Test_CovLC1_26_Loop", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return false
		})
		if count != 2 {
			t.Fatal("expected 2")
		}
		// empty
		corestr.Empty.LinkedCollections().Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			t.Fatal("should not be called")
			return false
		})
		// break
		lc2 := corestr.Empty.LinkedCollections()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2.Add(corestr.New.Collection.Strings([]string{"b"}))
		breakCount := 0
		lc2.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			breakCount++
			return true // break on first
		})
		if breakCount != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLC1_27_Filter(t *testing.T) {
	safeTest(t, "Test_CovLC1_27_Filter", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		result := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{
				Value:   arg.Node,
				IsKeep:  true,
				IsBreak: false,
			}
		})
		if len(result) != 2 {
			t.Fatal("expected 2")
		}
		// empty
		empty := corestr.Empty.LinkedCollections()
		r := empty.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		if len(r) != 0 {
			t.Fatal("expected 0")
		}
		// break
		r2 := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})
		if len(r2) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLC1_28_FilterAsCollection(t *testing.T) {
	safeTest(t, "Test_CovLC1_28_FilterAsCollection", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		col := lc.FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		}, 0)
		if col.Length() != 3 {
			t.Fatalf("expected 3, got %d", col.Length())
		}
		// empty result
		col2 := lc.FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: false}
		}, 0)
		if col2.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovLC1_29_FilterAsCollections(t *testing.T) {
	safeTest(t, "Test_CovLC1_29_FilterAsCollections", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		cols := lc.FilterAsCollections(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		if len(cols) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLC1_30_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_CovLC1_30_RemoveNodeByIndex", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		// remove first
		lc.RemoveNodeByIndex(0)
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
		// remove last
		lc.RemoveNodeByIndex(1)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
		// remove middle (rebuild)
		lc2 := corestr.Empty.LinkedCollections()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc2.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc2.RemoveNodeByIndex(1)
		if lc2.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLC1_31_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_CovLC1_31_RemoveNodeByIndexes", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc.RemoveNodeByIndexes(true, 0, 2)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
		// empty indexes
		lc.RemoveNodeByIndexes(true)
	})
}

func Test_CovLC1_32_RemoveNode(t *testing.T) {
	safeTest(t, "Test_CovLC1_32_RemoveNode", func() {
		lc := corestr.Empty.LinkedCollections()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.Add(c1)
		lc.Add(c2)
		head := lc.Head()
		lc.RemoveNode(head) // remove first
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
		// remove non-first
		lc2 := corestr.Empty.LinkedCollections()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2.Add(corestr.New.Collection.Strings([]string{"b"}))
		tail := lc2.Tail()
		lc2.RemoveNode(tail)
		if lc2.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLC1_33_AppendCollections(t *testing.T) {
	safeTest(t, "Test_CovLC1_33_AppendCollections", func() {
		lc := corestr.Empty.LinkedCollections()
		lc.AppendCollections(true, corestr.New.Collection.Strings([]string{"a"}), nil)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
		// nil with skip
		lc.AppendCollections(true)
	})
}

func Test_CovLC1_34_AppendCollectionsPointersLock(t *testing.T) {
	safeTest(t, "Test_CovLC1_34_AppendCollectionsPointersLock", func() {
		lc := corestr.Empty.LinkedCollections()
		cols := []*corestr.Collection{corestr.New.Collection.Strings([]string{"a"}), nil}
		lc.AppendCollectionsPointersLock(true, &cols)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
		lc.AppendCollectionsPointersLock(true, nil)
	})
}

func Test_CovLC1_35_AppendCollectionsPointers(t *testing.T) {
	safeTest(t, "Test_CovLC1_35_AppendCollectionsPointers", func() {
		lc := corestr.Empty.LinkedCollections()
		cols := []*corestr.Collection{corestr.New.Collection.Strings([]string{"a"}), nil}
		lc.AppendCollectionsPointers(true, &cols)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
		lc.AppendCollectionsPointers(true, nil)
	})
}
