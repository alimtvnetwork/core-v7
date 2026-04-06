package corestrtests

import (
	"errors"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// LinkedCollections.go — Seg-01: Lines 1–800 (~200 uncovered stmts)
// Covers: Head, Tail, First, Single, Last, LastOrDefault, FirstOrDefault,
//         Length, AllIndividualItemsLength, LengthLock, IsEqualsPtr, IsEmptyLock,
//         IsEmpty, HasItems, InsertAt, AddAsync, AddsAsyncOnComplete,
//         AddsUsingProcessorAsyncOnComplete, AddsUsingProcessorAsync,
//         AddLock, Add, AddStringsLock, AddStrings, AddBackNode, AppendNode,
//         AppendChainOfNodes, AppendChainOfNodesAsync, PushBackLock, PushBack,
//         Push, PushFront, AddFrontLock, AddFront, AttachWithNode, AddAnother,
//         AddCollectionToNode, GetNextNodes, GetAllLinkedNodes, Loop, Filter,
//         FilterAsCollection, FilterAsCollections, RemoveNodeByIndex,
//         RemoveNodeByIndexes, RemoveNode, AppendCollections,
//         AppendCollectionsPointersLock, AppendCollectionsPointers,
//         AddCollectionsToNodeAsync, AddCollectionsToNode,
//         AddCollectionsPointerToNode, AddAfterNode, AddAfterNodeAsync
// =============================================================================

func newLC(items ...[]string) *corestr.LinkedCollections {
	lc := corestr.New.LinkedCollection.Create()
	for _, s := range items {
		lc.Add(corestr.New.Collection.Strings(s))
	}
	return lc
}

func Test_Cov60_LC_Head_Tail(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Head_Tail", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		actual := args.Map{
			"headNotNil": lc.Head() != nil,
			"tailNotNil": lc.Tail() != nil,
			"first":      lc.First().First(),
			"last":       lc.Last().First(),
			"single":     lc.Head().Element.First(),
		}
		expected := args.Map{
			"headNotNil": true,
			"tailNotNil": true,
			"first":      "a",
			"last":       "b",
			"single":     "a",
		}
		expected.ShouldBeEqual(t, 0, "Head/Tail/First/Last returns correct nodes", actual)
	})
}

func Test_Cov60_LC_Single(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Single", func() {
		lc := newLC([]string{"x"})
		actual := args.Map{"val": lc.Single().First()}
		expected := args.Map{"val": "x"}
		expected.ShouldBeEqual(t, 0, "Single returns first element", actual)
	})
}

func Test_Cov60_LC_LastOrDefault_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_LastOrDefault_NonEmpty", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		actual := args.Map{"val": lc.LastOrDefault().First()}
		expected := args.Map{"val": "b"}
		expected.ShouldBeEqual(t, 0, "LastOrDefault returns last on non-empty", actual)
	})
}

func Test_Cov60_LC_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_LastOrDefault_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"empty": lc.LastOrDefault().IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "LastOrDefault returns empty on empty", actual)
	})
}

func Test_Cov60_LC_FirstOrDefault_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_FirstOrDefault_NonEmpty", func() {
		lc := newLC([]string{"a"})
		actual := args.Map{"val": lc.FirstOrDefault().First()}
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault returns first on non-empty", actual)
	})
}

func Test_Cov60_LC_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_FirstOrDefault_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"empty": lc.FirstOrDefault().IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault returns empty on empty", actual)
	})
}

func Test_Cov60_LC_Length(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Length", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Length returns count", actual)
	})
}

func Test_Cov60_LC_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AllIndividualItemsLength", func() {
		lc := newLC([]string{"a", "b"}, []string{"c"})
		actual := args.Map{"total": lc.AllIndividualItemsLength()}
		expected := args.Map{"total": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualItemsLength returns sum", actual)
	})
}

func Test_Cov60_LC_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_LengthLock", func() {
		lc := newLC([]string{"a"})
		actual := args.Map{"len": lc.LengthLock()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthLock returns count with lock", actual)
	})
}

func Test_Cov60_LC_IsEqualsPtr_Same(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_IsEqualsPtr_Same", func() {
		lc := newLC([]string{"a"})
		actual := args.Map{"eq": lc.IsEqualsPtr(lc)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsPtr same pointer returns true", actual)
	})
}

func Test_Cov60_LC_IsEqualsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_IsEqualsPtr_Nil", func() {
		lc := newLC([]string{"a"})
		actual := args.Map{"eq": lc.IsEqualsPtr(nil)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualsPtr nil returns false", actual)
	})
}

func Test_Cov60_LC_IsEqualsPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_IsEqualsPtr_BothEmpty", func() {
		a := corestr.New.LinkedCollection.Create()
		b := corestr.New.LinkedCollection.Create()
		actual := args.Map{"eq": a.IsEqualsPtr(b)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsPtr both empty returns true", actual)
	})
}

func Test_Cov60_LC_IsEqualsPtr_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_IsEqualsPtr_OneEmpty", func() {
		a := newLC([]string{"a"})
		b := corestr.New.LinkedCollection.Create()
		actual := args.Map{"eq": a.IsEqualsPtr(b)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualsPtr one empty returns false", actual)
	})
}

func Test_Cov60_LC_IsEqualsPtr_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_IsEqualsPtr_DiffLen", func() {
		a := newLC([]string{"a"})
		b := newLC([]string{"a"}, []string{"b"})
		actual := args.Map{"eq": a.IsEqualsPtr(b)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualsPtr different length returns false", actual)
	})
}

func Test_Cov60_LC_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_IsEmptyLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"empty": lc.IsEmptyLock()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock returns true for empty", actual)
	})
}

func Test_Cov60_LC_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_IsEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"empty": lc.IsEmpty(), "hasItems": lc.HasItems()}
		expected := args.Map{"empty": true, "hasItems": false}
		expected.ShouldBeEqual(t, 0, "IsEmpty/HasItems on empty", actual)
	})
}

func Test_Cov60_LC_HasItems(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_HasItems", func() {
		lc := newLC([]string{"a"})
		actual := args.Map{"hasItems": lc.HasItems()}
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "HasItems returns true on non-empty", actual)
	})
}

func Test_Cov60_LC_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_InsertAt_Front", func() {
		lc := newLC([]string{"b"})
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.InsertAt(0, col)
		actual := args.Map{"len": lc.Length(), "first": lc.First().First()}
		expected := args.Map{"len": 2, "first": "a"}
		expected.ShouldBeEqual(t, 0, "InsertAt 0 adds to front", actual)
	})
}

func Test_Cov60_LC_InsertAt_Middle(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_InsertAt_Middle", func() {
		lc := newLC([]string{"a"}, []string{"c"})
		col := corestr.New.Collection.Strings([]string{"b"})
		lc.InsertAt(1, col)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "InsertAt middle inserts", actual)
	})
}

func Test_Cov60_LC_AddAsync(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddAsync", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.AddAsync(col, wg)
		wg.Wait()
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddAsync adds item asynchronously", actual)
	})
}

func Test_Cov60_LC_AddsAsyncOnComplete(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddsAsyncOnComplete", func() {
		lc := corestr.New.LinkedCollection.Create()
		done := make(chan bool, 1)
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.AddsAsyncOnComplete(func(lc2 *corestr.LinkedCollections) {
			done <- true
		}, false, col)
		<-done
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsAsyncOnComplete adds and calls complete", actual)
	})
}

func Test_Cov60_LC_AddsUsingProcessorAsyncOnComplete(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddsUsingProcessorAsyncOnComplete", func() {
		lc := corestr.New.LinkedCollection.Create()
		done := make(chan bool, 1)
		processor := func(a any, i int) *corestr.Collection {
			return corestr.New.Collection.Strings([]string{a.(string)})
		}
		lc.AddsUsingProcessorAsyncOnComplete(func(lc2 *corestr.LinkedCollections) {
			done <- true
		}, processor, false, "x")
		<-done
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsUsingProcessorAsyncOnComplete processes", actual)
	})
}

func Test_Cov60_LC_AddsUsingProcessorAsyncOnComplete_NilSkip(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddsUsingProcessorAsyncOnComplete_NilSkip", func() {
		lc := corestr.New.LinkedCollection.Create()
		done := make(chan bool, 1)
		processor := func(a any, i int) *corestr.Collection {
			return nil
		}
		var anys []any
		anys = nil
		lc.AddsUsingProcessorAsyncOnComplete(func(lc2 *corestr.LinkedCollections) {
			done <- true
		}, processor, true, anys...)
		<-done
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsUsingProcessorAsyncOnComplete nil skip", actual)
	})
}

func Test_Cov60_LC_AddsUsingProcessorAsync(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddsUsingProcessorAsync", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		processor := func(a any, i int) *corestr.Collection {
			return corestr.New.Collection.Strings([]string{a.(string)})
		}
		lc.AddsUsingProcessorAsync(wg, processor, false, "a")
		wg.Wait()
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsUsingProcessorAsync processes", actual)
	})
}

func Test_Cov60_LC_AddsUsingProcessorAsync_NilSkip(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddsUsingProcessorAsync_NilSkip", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		processor := func(a any, i int) *corestr.Collection {
			return nil
		}
		var anys []any
		anys = nil
		lc.AddsUsingProcessorAsync(wg, processor, true, anys...)
		wg.Wait()
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsUsingProcessorAsync nil skip", actual)
	})
}

func Test_Cov60_LC_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.AddLock(col)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddLock adds with lock", actual)
	})
}

func Test_Cov60_LC_Add_ToEmpty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Add_ToEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(col)
		actual := args.Map{"len": lc.Length(), "first": lc.First().First()}
		expected := args.Map{"len": 1, "first": "a"}
		expected.ShouldBeEqual(t, 0, "Add to empty sets head and tail", actual)
	})
}

func Test_Cov60_LC_Add_ToExisting(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Add_ToExisting", func() {
		lc := newLC([]string{"a"})
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		actual := args.Map{"len": lc.Length(), "last": lc.Last().First()}
		expected := args.Map{"len": 2, "last": "b"}
		expected.ShouldBeEqual(t, 0, "Add to existing appends to tail", actual)
	})
}

func Test_Cov60_LC_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddStringsLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock("a", "b")
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddStringsLock adds strings as collection", actual)
	})
}

func Test_Cov60_LC_AddStringsLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddStringsLock_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock()
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStringsLock no args returns same", actual)
	})
}

func Test_Cov60_LC_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("x", "y")
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddStrings adds strings as collection", actual)
	})
}

func Test_Cov60_LC_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddStrings_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings()
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStrings empty returns same", actual)
	})
}

func Test_Cov60_LC_AddBackNode(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddBackNode", func() {
		lc := newLC([]string{"a"})
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		lc.AddBackNode(node)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddBackNode appends node", actual)
	})
}

func Test_Cov60_LC_AppendNode_ToEmpty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AppendNode_ToEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		lc.AppendNode(node)
		actual := args.Map{"len": lc.Length(), "first": lc.First().First()}
		expected := args.Map{"len": 1, "first": "a"}
		expected.ShouldBeEqual(t, 0, "AppendNode to empty sets head", actual)
	})
}

func Test_Cov60_LC_AppendChainOfNodesAsync(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AppendChainOfNodesAsync", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		lc.AppendChainOfNodesAsync(node, wg)
		wg.Wait()
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AppendChainOfNodesAsync adds chain", actual)
	})
}

func Test_Cov60_LC_PushBackLock(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_PushBackLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.PushBackLock(col)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "PushBackLock adds to back", actual)
	})
}

func Test_Cov60_LC_PushBack(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_PushBack", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.PushBack(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "PushBack adds to back", actual)
	})
}

func Test_Cov60_LC_Push(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Push", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Push(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Push adds to back", actual)
	})
}

func Test_Cov60_LC_PushFront(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_PushFront", func() {
		lc := newLC([]string{"b"})
		lc.PushFront(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"len": lc.Length(), "first": lc.First().First()}
		expected := args.Map{"len": 2, "first": "a"}
		expected.ShouldBeEqual(t, 0, "PushFront adds to front", actual)
	})
}

func Test_Cov60_LC_AddFrontLock(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddFrontLock", func() {
		lc := newLC([]string{"b"})
		lc.AddFrontLock(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"len": lc.Length(), "first": lc.First().First()}
		expected := args.Map{"len": 2, "first": "a"}
		expected.ShouldBeEqual(t, 0, "AddFrontLock adds to front with lock", actual)
	})
}

func Test_Cov60_LC_AddFront_ToEmpty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddFront_ToEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddFront(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddFront to empty adds", actual)
	})
}

func Test_Cov60_LC_AddAnother(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddAnother", func() {
		a := newLC([]string{"a"})
		b := newLC([]string{"b"}, []string{"c"})
		a.AddAnother(b)
		actual := args.Map{"len": a.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AddAnother adds all from other", actual)
	})
}

func Test_Cov60_LC_AddAnother_Nil(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddAnother_Nil", func() {
		a := newLC([]string{"a"})
		a.AddAnother(nil)
		actual := args.Map{"len": a.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddAnother nil returns same", actual)
	})
}

func Test_Cov60_LC_AddAnother_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddAnother_Empty", func() {
		a := newLC([]string{"a"})
		a.AddAnother(corestr.New.LinkedCollection.Create())
		actual := args.Map{"len": a.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddAnother empty returns same", actual)
	})
}

func Test_Cov60_LC_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_GetNextNodes", func() {
		lc := newLC([]string{"a"}, []string{"b"}, []string{"c"})
		nodes := lc.GetNextNodes(2)
		actual := args.Map{"len": len(nodes)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetNextNodes returns limited nodes", actual)
	})
}

func Test_Cov60_LC_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_GetAllLinkedNodes", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		nodes := lc.GetAllLinkedNodes()
		actual := args.Map{"len": len(nodes)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllLinkedNodes returns all nodes", actual)
	})
}

func Test_Cov60_LC_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Loop_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return false
		})
		actual := args.Map{"count": count}
		expected := args.Map{"count": 0}
		expected.ShouldBeEqual(t, 0, "Loop on empty does nothing", actual)
	})
}

func Test_Cov60_LC_Loop_Break(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Loop_Break", func() {
		lc := newLC([]string{"a"}, []string{"b"}, []string{"c"})
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return true
		})
		actual := args.Map{"count": count}
		expected := args.Map{"count": 1}
		expected.ShouldBeEqual(t, 0, "Loop breaks after first", actual)
	})
}

func Test_Cov60_LC_Loop_BreakSecond(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Loop_BreakSecond", func() {
		lc := newLC([]string{"a"}, []string{"b"}, []string{"c"})
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return arg.Index >= 1
		})
		actual := args.Map{"count": count}
		expected := args.Map{"count": 2}
		expected.ShouldBeEqual(t, 0, "Loop breaks on second iteration", actual)
	})
}

func Test_Cov60_LC_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Filter_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		r := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Filter on empty returns empty", actual)
	})
}

func Test_Cov60_LC_Filter_KeepAll(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Filter_KeepAll", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		r := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Filter keeps all", actual)
	})
}

func Test_Cov60_LC_Filter_BreakFirst(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Filter_BreakFirst", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		r := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Filter breaks on first", actual)
	})
}

func Test_Cov60_LC_Filter_BreakSecond(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Filter_BreakSecond", func() {
		lc := newLC([]string{"a"}, []string{"b"}, []string{"c"})
		r := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: arg.Index >= 1}
		})
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Filter breaks on second", actual)
	})
}

func Test_Cov60_LC_Filter_Skip(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Filter_Skip", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		r := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: false}
		})
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Filter skips all", actual)
	})
}

func Test_Cov60_LC_FilterAsCollection(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_FilterAsCollection", func() {
		lc := newLC([]string{"a", "b"}, []string{"c"})
		r := lc.FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		}, 0)
		actual := args.Map{"len": r.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "FilterAsCollection merges filtered nodes", actual)
	})
}

func Test_Cov60_LC_FilterAsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_FilterAsCollection_Empty", func() {
		lc := newLC([]string{"a"})
		r := lc.FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: false}
		}, 0)
		actual := args.Map{"empty": r.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "FilterAsCollection no matches returns empty", actual)
	})
}

func Test_Cov60_LC_FilterAsCollections(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_FilterAsCollections", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		r := lc.FilterAsCollections(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "FilterAsCollections returns collection slice", actual)
	})
}

func Test_Cov60_LC_RemoveNodeByIndex_First(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_RemoveNodeByIndex_First", func() {
		lc := newLC([]string{"a"}, []string{"b"}, []string{"c"})
		lc.RemoveNodeByIndex(0)
		actual := args.Map{"len": lc.Length(), "first": lc.First().First()}
		expected := args.Map{"len": 2, "first": "b"}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndex removes first", actual)
	})
}

func Test_Cov60_LC_RemoveNodeByIndex_Last(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_RemoveNodeByIndex_Last", func() {
		lc := newLC([]string{"a"}, []string{"b"}, []string{"c"})
		lc.RemoveNodeByIndex(2)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndex removes last", actual)
	})
}

func Test_Cov60_LC_RemoveNodeByIndex_Middle(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_RemoveNodeByIndex_Middle", func() {
		lc := newLC([]string{"a"}, []string{"b"}, []string{"c"})
		lc.RemoveNodeByIndex(1)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndex removes middle", actual)
	})
}

func Test_Cov60_LC_RemoveNodeByIndex_Negative_Panics(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_RemoveNodeByIndex_Negative_Panics", func() {
		lc := newLC([]string{"a"})
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			lc.RemoveNodeByIndex(-1)
		}()
		actual := args.Map{"panicked": panicked}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndex negative panics", actual)
	})
}

func Test_Cov60_LC_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_RemoveNodeByIndexes", func() {
		lc := newLC([]string{"a"}, []string{"b"}, []string{"c"}, []string{"d"})
		lc.RemoveNodeByIndexes(true, 1, 3)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndexes removes specified", actual)
	})
}

func Test_Cov60_LC_RemoveNodeByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_RemoveNodeByIndexes_Empty", func() {
		lc := newLC([]string{"a"})
		lc.RemoveNodeByIndexes(true)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndexes no args returns same", actual)
	})
}

func Test_Cov60_LC_RemoveNodeByIndexes_EmptyPanics(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_RemoveNodeByIndexes_EmptyPanics", func() {
		lc := corestr.New.LinkedCollection.Create()
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			lc.RemoveNodeByIndexes(false, 0)
		}()
		actual := args.Map{"panicked": panicked}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndexes empty without ignore panics", actual)
	})
}

func Test_Cov60_LC_RemoveNode(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_RemoveNode", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		node := lc.Head()
		lc.RemoveNode(node)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveNode removes head", actual)
	})
}

func Test_Cov60_LC_RemoveNode_Second(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_RemoveNode_Second", func() {
		lc := newLC([]string{"a"}, []string{"b"}, []string{"c"})
		node := lc.IndexAt(1)
		lc.RemoveNode(node)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveNode removes second node", actual)
	})
}

func Test_Cov60_LC_AppendCollections_SkipNil(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AppendCollections_SkipNil", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.AppendCollections(true, col, nil)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AppendCollections skips nil", actual)
	})
}

func Test_Cov60_LC_AppendCollections_NilArg(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AppendCollections_NilArg", func() {
		lc := corestr.New.LinkedCollection.Create()
		var cols []*corestr.Collection
		lc.AppendCollections(true, cols...)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendCollections nil arg returns same", actual)
	})
}

func Test_Cov60_LC_AppendCollectionsPointersLock(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AppendCollectionsPointersLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		cols := []*corestr.Collection{col, nil}
		lc.AppendCollectionsPointersLock(true, &cols)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AppendCollectionsPointersLock adds with lock", actual)
	})
}

func Test_Cov60_LC_AppendCollectionsPointersLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AppendCollectionsPointersLock_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AppendCollectionsPointersLock(true, nil)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendCollectionsPointersLock nil returns same", actual)
	})
}

func Test_Cov60_LC_AppendCollectionsPointers(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AppendCollectionsPointers", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		cols := []*corestr.Collection{col}
		lc.AppendCollectionsPointers(false, &cols)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AppendCollectionsPointers adds", actual)
	})
}

func Test_Cov60_LC_AppendCollectionsPointers_Nil(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AppendCollectionsPointers_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AppendCollectionsPointers(true, nil)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendCollectionsPointers nil returns same", actual)
	})
}

func Test_Cov60_LC_AttachWithNode_NilPanics(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AttachWithNode_NilPanics", func() {
		lc := corestr.New.LinkedCollection.Create()
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		err := lc.AttachWithNode(nil, node)
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "AttachWithNode nil current returns error", actual)
	})
}

func Test_Cov60_LC_AttachWithNode_NextNotNil(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AttachWithNode_NextNotNil", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		node := lc.Head()
		addNode := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"c"})}
		err := lc.AttachWithNode(node, addNode)
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "AttachWithNode next not nil returns error", actual)
	})
}

func Test_Cov60_LC_AddCollectionsToNodeAsync(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddCollectionsToNodeAsync", func() {
		lc := newLC([]string{"a"})
		wg := &sync.WaitGroup{}
		wg.Add(1)
		node := lc.Head()
		col := corestr.New.Collection.Strings([]string{"b"})
		lc.AddCollectionsToNodeAsync(false, wg, node, col)
		wg.Wait()
		actual := args.Map{"lenGte": lc.Length() >= 2}
		expected := args.Map{"lenGte": true}
		expected.ShouldBeEqual(t, 0, "AddCollectionsToNodeAsync adds async", actual)
	})
}

func Test_Cov60_LC_AddCollectionsToNodeAsync_NilSkip(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddCollectionsToNodeAsync_NilSkip", func() {
		lc := newLC([]string{"a"})
		var cols []*corestr.Collection
		lc.AddCollectionsToNodeAsync(true, nil, nil, cols...)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCollectionsToNodeAsync nil skips", actual)
	})
}

func Test_Cov60_LC_AddCollectionsToNode(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddCollectionsToNode", func() {
		lc := newLC([]string{"a"})
		node := lc.Head()
		col := corestr.New.Collection.Strings([]string{"b"})
		lc.AddCollectionsToNode(false, node, col)
		actual := args.Map{"lenGte": lc.Length() >= 2}
		expected := args.Map{"lenGte": true}
		expected.ShouldBeEqual(t, 0, "AddCollectionsToNode adds after node", actual)
	})
}

func Test_Cov60_LC_AddCollectionsToNode_NilSkip(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddCollectionsToNode_NilSkip", func() {
		lc := newLC([]string{"a"})
		var cols []*corestr.Collection
		lc.AddCollectionsToNode(true, nil, cols...)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCollectionsToNode nil skip returns same", actual)
	})
}

func Test_Cov60_LC_AddCollectionsPointerToNode_NilItems(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddCollectionsPointerToNode_NilItems", func() {
		lc := newLC([]string{"a"})
		lc.AddCollectionsPointerToNode(true, nil, nil)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCollectionsPointerToNode nil items returns same", actual)
	})
}

func Test_Cov60_LC_AddCollectionsPointerToNode_NilNodeSkip(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddCollectionsPointerToNode_NilNodeSkip", func() {
		lc := newLC([]string{"a"})
		col := corestr.New.Collection.Strings([]string{"b"})
		cols := []*corestr.Collection{col}
		lc.AddCollectionsPointerToNode(true, nil, &cols)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCollectionsPointerToNode nil node skip returns same", actual)
	})
}

func Test_Cov60_LC_AddCollectionsPointerToNode_NilNodePanics(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddCollectionsPointerToNode_NilNodePanics", func() {
		lc := newLC([]string{"a"})
		col := corestr.New.Collection.Strings([]string{"b"})
		cols := []*corestr.Collection{col}
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			lc.AddCollectionsPointerToNode(false, nil, &cols)
		}()
		actual := args.Map{"panicked": panicked}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "AddCollectionsPointerToNode nil node panics", actual)
	})
}

func Test_Cov60_LC_AddCollectionsPointerToNode_EmptyItems(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddCollectionsPointerToNode_EmptyItems", func() {
		lc := newLC([]string{"a"})
		node := lc.Head()
		cols := []*corestr.Collection{}
		lc.AddCollectionsPointerToNode(false, node, &cols)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCollectionsPointerToNode empty items returns same", actual)
	})
}

func Test_Cov60_LC_AddCollectionsPointerToNode_SingleItem(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddCollectionsPointerToNode_SingleItem", func() {
		lc := newLC([]string{"a"}, []string{"c"})
		node := lc.Head()
		col := corestr.New.Collection.Strings([]string{"b"})
		cols := []*corestr.Collection{col}
		lc.AddCollectionsPointerToNode(false, node, &cols)
		actual := args.Map{"lenGte": lc.Length() >= 3}
		expected := args.Map{"lenGte": true}
		expected.ShouldBeEqual(t, 0, "AddCollectionsPointerToNode single item adds", actual)
	})
}

func Test_Cov60_LC_AddCollectionsPointerToNode_MultiItems(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddCollectionsPointerToNode_MultiItems", func() {
		lc := newLC([]string{"a"})
		node := lc.Head()
		col1 := corestr.New.Collection.Strings([]string{"b"})
		col2 := corestr.New.Collection.Strings([]string{"c"})
		cols := []*corestr.Collection{col1, col2}
		lc.AddCollectionsPointerToNode(false, node, &cols)
		actual := args.Map{"lenGte": lc.Length() >= 3}
		expected := args.Map{"lenGte": true}
		expected.ShouldBeEqual(t, 0, "AddCollectionsPointerToNode multi items adds chain", actual)
	})
}

func Test_Cov60_LC_AddAfterNode(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddAfterNode", func() {
		lc := newLC([]string{"a"}, []string{"c"})
		node := lc.Head()
		col := corestr.New.Collection.Strings([]string{"b"})
		lc.AddAfterNode(node, col)
		actual := args.Map{"lenGte": lc.Length() >= 3}
		expected := args.Map{"lenGte": true}
		expected.ShouldBeEqual(t, 0, "AddAfterNode inserts after node", actual)
	})
}

func Test_Cov60_LC_AddAfterNodeAsync(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddAfterNodeAsync", func() {
		lc := newLC([]string{"a"})
		wg := &sync.WaitGroup{}
		wg.Add(1)
		node := lc.Head()
		col := corestr.New.Collection.Strings([]string{"b"})
		lc.AddAfterNodeAsync(wg, node, col)
		wg.Wait()
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddAfterNodeAsync inserts async", actual)
	})
}

// =============================================================================
// LinkedCollections.go — Seg-01 Part B: Lines 800–1551
// Covers: ConcatNew, AddAsyncFuncItems, AddAsyncFuncItemsPointer,
//         AddStringsOfStrings, IndexAt, SafePointerIndexAt, SafeIndexAt,
//         AddStringsAsync, AddCollection, AddCollectionsPtr, AddCollections,
//         ToStringsPtr, ToStrings, ToCollectionSimple, ToCollection,
//         ToCollectionsOfCollection, ItemsOfItems, ItemsOfItemsCollection,
//         SimpleSlice, ListPtr, List, String, StringLock, Join, Joins,
//         JsonModel, JsonModelAny, MarshalJSON, UnmarshalJSON, RemoveAll,
//         Clear, Json, JsonPtr, ParseInjectUsingJson, ParseInjectUsingJsonMust,
//         GetCompareSummary, JsonParseSelfInject, AsJsonContractsBinder,
//         AsJsoner, AsJsonParseSelfInjector, AsJsonMarshaller
// =============================================================================

func Test_Cov60_LC_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ConcatNew_EmptyClone", func() {
		lc := newLC([]string{"a"})
		r := lc.ConcatNew(true)
		actual := args.Map{"len": r.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ConcatNew empty with clone returns clone", actual)
	})
}

func Test_Cov60_LC_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ConcatNew_EmptyNoClone", func() {
		lc := newLC([]string{"a"})
		r := lc.ConcatNew(false)
		actual := args.Map{"same": r == lc}
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "ConcatNew empty no clone returns self", actual)
	})
}

func Test_Cov60_LC_ConcatNew_WithOthers(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ConcatNew_WithOthers", func() {
		a := newLC([]string{"a"})
		b := newLC([]string{"b"})
		r := a.ConcatNew(false, b)
		actual := args.Map{"len": r.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ConcatNew combines collections", actual)
	})
}

func Test_Cov60_LC_AddAsyncFuncItems(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddAsyncFuncItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItems(wg, false, func() []string { return []string{"a"} })
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddAsyncFuncItems adds func results", actual)
	})
}

func Test_Cov60_LC_AddAsyncFuncItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddAsyncFuncItems_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAsyncFuncItems(nil, false)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddAsyncFuncItems nil returns same", actual)
	})
}

func Test_Cov60_LC_AddAsyncFuncItems_EmptyResult(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddAsyncFuncItems_EmptyResult", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItems(wg, false, func() []string { return []string{} })
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddAsyncFuncItems empty result skips", actual)
	})
}

func Test_Cov60_LC_AddAsyncFuncItemsPointer(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddAsyncFuncItemsPointer", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItemsPointer(wg, false, func() []string { return []string{"a"} })
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddAsyncFuncItemsPointer adds", actual)
	})
}

func Test_Cov60_LC_AddAsyncFuncItemsPointer_Nil(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddAsyncFuncItemsPointer_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAsyncFuncItemsPointer(nil, false)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddAsyncFuncItemsPointer nil returns same", actual)
	})
}

func Test_Cov60_LC_AddAsyncFuncItemsPointer_EmptyResult(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddAsyncFuncItemsPointer_EmptyResult", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItemsPointer(wg, false, func() []string { return []string{} })
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddAsyncFuncItemsPointer empty result skips", actual)
	})
}

func Test_Cov60_LC_AddStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddStringsOfStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false, []string{"a"}, []string{"b"})
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsOfStrings adds slices", actual)
	})
}

func Test_Cov60_LC_AddStringsOfStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddStringsOfStrings_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStringsOfStrings empty returns same", actual)
	})
}

func Test_Cov60_LC_AddStringsOfStrings_NilSlice(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddStringsOfStrings_NilSlice", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false, nil, []string{"a"})
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddStringsOfStrings skips nil slices", actual)
	})
}

func Test_Cov60_LC_IndexAt_Zero(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_IndexAt_Zero", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		node := lc.IndexAt(0)
		actual := args.Map{"val": node.Element.First()}
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "IndexAt 0 returns head", actual)
	})
}

func Test_Cov60_LC_IndexAt_Last(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_IndexAt_Last", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		node := lc.IndexAt(1)
		actual := args.Map{"val": node.Element.First()}
		expected := args.Map{"val": "b"}
		expected.ShouldBeEqual(t, 0, "IndexAt last returns tail", actual)
	})
}

func Test_Cov60_LC_IndexAt_Negative(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_IndexAt_Negative", func() {
		lc := newLC([]string{"a"})
		node := lc.IndexAt(-1)
		actual := args.Map{"nil": node == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "IndexAt negative returns nil", actual)
	})
}

func Test_Cov60_LC_IndexAt_OutOfRange_Panics(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_IndexAt_OutOfRange_Panics", func() {
		lc := newLC([]string{"a"})
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			lc.IndexAt(5)
		}()
		actual := args.Map{"panicked": panicked}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "IndexAt out of range panics", actual)
	})
}

func Test_Cov60_LC_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_SafePointerIndexAt", func() {
		lc := newLC([]string{"a"})
		r := lc.SafePointerIndexAt(0)
		actual := args.Map{"val": r.First()}
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "SafePointerIndexAt returns element", actual)
	})
}

func Test_Cov60_LC_SafePointerIndexAt_Nil(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_SafePointerIndexAt_Nil", func() {
		lc := newLC([]string{"a"})
		r := lc.SafePointerIndexAt(5)
		actual := args.Map{"nil": r == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "SafePointerIndexAt out of range returns nil", actual)
	})
}

func Test_Cov60_LC_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_SafeIndexAt", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		node := lc.SafeIndexAt(1)
		actual := args.Map{"val": node.Element.First()}
		expected := args.Map{"val": "b"}
		expected.ShouldBeEqual(t, 0, "SafeIndexAt returns node", actual)
	})
}

func Test_Cov60_LC_SafeIndexAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_SafeIndexAt_OutOfRange", func() {
		lc := newLC([]string{"a"})
		node := lc.SafeIndexAt(5)
		actual := args.Map{"nil": node == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "SafeIndexAt out of range returns nil", actual)
	})
}

func Test_Cov60_LC_SafeIndexAt_Negative(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_SafeIndexAt_Negative", func() {
		lc := newLC([]string{"a"})
		node := lc.SafeIndexAt(-1)
		actual := args.Map{"nil": node == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "SafeIndexAt negative returns nil", actual)
	})
}

func Test_Cov60_LC_SafeIndexAt_Zero(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_SafeIndexAt_Zero", func() {
		lc := newLC([]string{"a"})
		node := lc.SafeIndexAt(0)
		actual := args.Map{"val": node.Element.First()}
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "SafeIndexAt 0 returns head", actual)
	})
}

func Test_Cov60_LC_AddStringsAsync(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddStringsAsync", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddStringsAsync(wg, []string{"a"})
		wg.Wait()
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddStringsAsync adds strings async", actual)
	})
}

func Test_Cov60_LC_AddStringsAsync_Nil(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddStringsAsync_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsAsync(nil, nil)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStringsAsync nil returns same", actual)
	})
}

func Test_Cov60_LC_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollection(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCollection adds", actual)
	})
}

func Test_Cov60_LC_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddCollection_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollection(nil)
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollection nil returns same", actual)
	})
}

func Test_Cov60_LC_AddCollectionsPtr(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddCollectionsPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.AddCollectionsPtr([]*corestr.Collection{col})
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCollectionsPtr adds", actual)
	})
}

func Test_Cov60_LC_AddCollectionsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddCollectionsPtr_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollectionsPtr([]*corestr.Collection{})
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollectionsPtr empty returns same", actual)
	})
}

func Test_Cov60_LC_AddCollections_SkipNil(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AddCollections_SkipNil", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.AddCollections([]*corestr.Collection{nil, col})
		actual := args.Map{"len": lc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCollections skips nil", actual)
	})
}

func Test_Cov60_LC_ToStringsPtr(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ToStringsPtr", func() {
		lc := newLC([]string{"a", "b"})
		r := lc.ToStringsPtr()
		actual := args.Map{"len": len(*r)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ToStringsPtr returns pointer to strings", actual)
	})
}

func Test_Cov60_LC_ToStrings(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ToStrings", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		r := lc.ToStrings()
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ToStrings returns all strings", actual)
	})
}

func Test_Cov60_LC_ToCollectionSimple(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ToCollectionSimple", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		r := lc.ToCollectionSimple()
		actual := args.Map{"len": r.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ToCollectionSimple merges all", actual)
	})
}

func Test_Cov60_LC_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ToCollection_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		r := lc.ToCollection(0)
		actual := args.Map{"empty": r.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "ToCollection empty returns empty", actual)
	})
}

func Test_Cov60_LC_ToCollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ToCollectionsOfCollection", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		r := lc.ToCollectionsOfCollection(0)
		actual := args.Map{"len": r.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ToCollectionsOfCollection returns all", actual)
	})
}

func Test_Cov60_LC_ToCollectionsOfCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ToCollectionsOfCollection_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		r := lc.ToCollectionsOfCollection(0)
		actual := args.Map{"empty": r.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "ToCollectionsOfCollection empty returns empty", actual)
	})
}

func Test_Cov60_LC_ItemsOfItems(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ItemsOfItems", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		r := lc.ItemsOfItems()
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ItemsOfItems returns slices", actual)
	})
}

func Test_Cov60_LC_ItemsOfItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ItemsOfItems_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		r := lc.ItemsOfItems()
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ItemsOfItems empty returns empty", actual)
	})
}

func Test_Cov60_LC_ItemsOfItemsCollection(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ItemsOfItemsCollection", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		r := lc.ItemsOfItemsCollection()
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ItemsOfItemsCollection returns collections", actual)
	})
}

func Test_Cov60_LC_ItemsOfItemsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ItemsOfItemsCollection_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		r := lc.ItemsOfItemsCollection()
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ItemsOfItemsCollection empty returns empty", actual)
	})
}

func Test_Cov60_LC_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_SimpleSlice", func() {
		lc := newLC([]string{"a"})
		r := lc.SimpleSlice()
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns SimpleSlice", actual)
	})
}

func Test_Cov60_LC_ListPtr(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ListPtr", func() {
		lc := newLC([]string{"a"})
		r := lc.ListPtr()
		actual := args.Map{"len": len(*r)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListPtr returns pointer to list", actual)
	})
}

func Test_Cov60_LC_List(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_List", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		r := lc.List()
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "List returns all strings", actual)
	})
}

func Test_Cov60_LC_List_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_List_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		r := lc.List()
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "List empty returns empty", actual)
	})
}

func Test_Cov60_LC_String(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_String", func() {
		lc := newLC([]string{"a"})
		s := lc.String()
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String returns formatted", actual)
	})
}

func Test_Cov60_LC_String_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_String_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		s := lc.String()
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String empty returns no-elements marker", actual)
	})
}

func Test_Cov60_LC_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_StringLock", func() {
		lc := newLC([]string{"a"})
		s := lc.StringLock()
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock returns string", actual)
	})
}

func Test_Cov60_LC_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_StringLock_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		s := lc.StringLock()
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock empty returns marker", actual)
	})
}

func Test_Cov60_LC_Join(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Join", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		s := lc.Join(",")
		actual := args.Map{"val": s}
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Join joins all strings", actual)
	})
}

func Test_Cov60_LC_Joins(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Joins", func() {
		lc := newLC([]string{"a"})
		s := lc.Joins(",", "b")
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Joins joins with extra items", actual)
	})
}

func Test_Cov60_LC_Joins_NilItems(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Joins_NilItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		s := lc.Joins(",")
		actual := args.Map{"empty": s == ""}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Joins nil items returns empty", actual)
	})
}

func Test_Cov60_LC_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_JsonModel", func() {
		lc := newLC([]string{"a"})
		r := lc.JsonModel()
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel returns strings", actual)
	})
}

func Test_Cov60_LC_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_JsonModelAny", func() {
		lc := newLC([]string{"a"})
		r := lc.JsonModelAny()
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny returns any", actual)
	})
}

func Test_Cov60_LC_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_MarshalJSON", func() {
		lc := newLC([]string{"a"})
		b, err := lc.MarshalJSON()
		actual := args.Map{"noErr": err == nil, "nonEmpty": len(b) > 0}
		expected := args.Map{"noErr": true, "nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "MarshalJSON returns bytes", actual)
	})
}

func Test_Cov60_LC_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_UnmarshalJSON", func() {
		lc := &corestr.LinkedCollections{}
		err := lc.UnmarshalJSON([]byte(`["a","b"]`))
		actual := args.Map{"noErr": err == nil, "len": lc.Length()}
		expected := args.Map{"noErr": true, "len": 1}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON parses items", actual)
	})
}

func Test_Cov60_LC_UnmarshalJSON_Error(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_UnmarshalJSON_Error", func() {
		lc := &corestr.LinkedCollections{}
		err := lc.UnmarshalJSON([]byte(`invalid`))
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON returns error on bad input", actual)
	})
}

func Test_Cov60_LC_RemoveAll(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_RemoveAll", func() {
		lc := newLC([]string{"a"}, []string{"b"})
		lc.RemoveAll()
		actual := args.Map{"empty": lc.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "RemoveAll clears all", actual)
	})
}

func Test_Cov60_LC_Clear(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Clear", func() {
		lc := newLC([]string{"a"})
		lc.Clear()
		actual := args.Map{"empty": lc.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Clear empties collection", actual)
	})
}

func Test_Cov60_LC_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Clear_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Clear()
		actual := args.Map{"empty": lc.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Clear on empty returns same", actual)
	})
}

func Test_Cov60_LC_Json(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_Json", func() {
		lc := newLC([]string{"a"})
		r := lc.Json()
		actual := args.Map{"nonEmpty": r.JsonString() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Json returns Result", actual)
	})
}

func Test_Cov60_LC_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_JsonPtr", func() {
		lc := newLC([]string{"a"})
		r := lc.JsonPtr()
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonPtr returns Result ptr", actual)
	})
}

func Test_Cov60_LC_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ParseInjectUsingJson", func() {
		lc := &corestr.LinkedCollections{}
		jr := corejson.NewPtr([]string{"a", "b"})
		r, err := lc.ParseInjectUsingJson(jr)
		actual := args.Map{"noErr": err == nil, "nonNil": r != nil}
		expected := args.Map{"noErr": true, "nonNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson parses", actual)
	})
}

var errLinkedTest = errors.New("test error")

func Test_Cov60_LC_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ParseInjectUsingJson_Error", func() {
		lc := &corestr.LinkedCollections{}
		jr := &corejson.Result{Error: errLinkedTest}
		_, err := lc.ParseInjectUsingJson(jr)
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns error", actual)
	})
}

func Test_Cov60_LC_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ParseInjectUsingJsonMust", func() {
		lc := &corestr.LinkedCollections{}
		jr := corejson.NewPtr([]string{"a"})
		r := lc.ParseInjectUsingJsonMust(jr)
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust parses", actual)
	})
}

func Test_Cov60_LC_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_ParseInjectUsingJsonMust_Panics", func() {
		lc := &corestr.LinkedCollections{}
		jr := &corejson.Result{Error: errLinkedTest}
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			lc.ParseInjectUsingJsonMust(jr)
		}()
		actual := args.Map{"panicked": panicked}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust panics on error", actual)
	})
}

func Test_Cov60_LC_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_GetCompareSummary", func() {
		a := newLC([]string{"a"})
		b := newLC([]string{"b"})
		s := a.GetCompareSummary(b, "left", "right")
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "GetCompareSummary returns summary", actual)
	})
}

func Test_Cov60_LC_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_JsonParseSelfInject", func() {
		lc := &corestr.LinkedCollections{}
		jr := corejson.NewPtr([]string{"a"})
		err := lc.JsonParseSelfInject(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject injects", actual)
	})
}

func Test_Cov60_LC_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AsJsonContractsBinder", func() {
		lc := newLC([]string{"a"})
		actual := args.Map{"nonNil": lc.AsJsonContractsBinder() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns interface", actual)
	})
}

func Test_Cov60_LC_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AsJsoner", func() {
		lc := newLC([]string{"a"})
		actual := args.Map{"nonNil": lc.AsJsoner() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsoner returns interface", actual)
	})
}

func Test_Cov60_LC_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AsJsonParseSelfInjector", func() {
		lc := newLC([]string{"a"})
		actual := args.Map{"nonNil": lc.AsJsonParseSelfInjector() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonParseSelfInjector returns interface", actual)
	})
}

func Test_Cov60_LC_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov60_LC_AsJsonMarshaller", func() {
		lc := newLC([]string{"a"})
		actual := args.Map{"nonNil": lc.AsJsonMarshaller() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonMarshaller returns interface", actual)
	})
}
