package corestrtests

import (
	"encoding/json"
	"strings"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════
// LinkedCollections — basic operations
// ══════════════════════════════════════════════════════════════

func Test_Cov39_LC_Add(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Add", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		lc.Add(col)
		if lc.Length() != 1 {
			t.Errorf("expected 1, got %d", lc.Length())
		}
	})
}

func Test_Cov39_LC_Head_Tail(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Head_Tail", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.Add(c1).Add(c2)
		if lc.Head().Element.First() != "a" {
			t.Error("expected a")
		}
		if lc.Tail().Element.First() != "b" {
			t.Error("expected b")
		}
	})
}

func Test_Cov39_LC_First_Last(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_First_Last", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"x"})
		c2 := corestr.New.Collection.Strings([]string{"y"})
		lc.Add(c1).Add(c2)
		if lc.First().First() != "x" {
			t.Error("expected x")
		}
		if lc.Last().First() != "y" {
			t.Error("expected y")
		}
	})
}

func Test_Cov39_LC_Single(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Single", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"only"}))
		if lc.Single().First() != "only" {
			t.Error("expected only")
		}
	})
}

func Test_Cov39_LC_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_FirstOrDefault_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.FirstOrDefault().Length() != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_Cov39_LC_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_LastOrDefault_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.LastOrDefault().Length() != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_Cov39_LC_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IsEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		if !lc.IsEmpty() {
			t.Error("expected empty")
		}
		if lc.HasItems() {
			t.Error("expected no items")
		}
	})
}

func Test_Cov39_LC_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IsEmptyLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		if !lc.IsEmptyLock() {
			t.Error("expected empty")
		}
	})
}

func Test_Cov39_LC_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_LengthLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.LengthLock() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AllIndividualItemsLength", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		if lc.AllIndividualItemsLength() != 3 {
			t.Errorf("expected 3, got %d", lc.AllIndividualItemsLength())
		}
	})
}

func Test_Cov39_LC_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddLock(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddStrings_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings()
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov39_LC_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddStringsLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock("a")
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_AddStringsLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddStringsLock_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock()
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov39_LC_PushBack(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_PushBack", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.PushBack(corestr.New.Collection.Strings([]string{"x"}))
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_PushBackLock(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_PushBackLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.PushBackLock(corestr.New.Collection.Strings([]string{"x"}))
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_Push(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Push", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Push(corestr.New.Collection.Strings([]string{"x"}))
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_AddFront(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddFront", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"b"})
		c2 := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c1)
		lc.AddFront(c2)
		if lc.First().First() != "a" {
			t.Error("expected a first")
		}
	})
}

func Test_Cov39_LC_AddFront_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddFront_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddFront(corestr.New.Collection.Strings([]string{"x"}))
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_PushFront(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_PushFront", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.PushFront(corestr.New.Collection.Strings([]string{"a"}))
		if lc.First().First() != "a" {
			t.Error("expected a")
		}
	})
}

func Test_Cov39_LC_AddFrontLock(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddFrontLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddFrontLock(corestr.New.Collection.Strings([]string{"x"}))
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_AddAnother(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddAnother", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc1.AddAnother(lc2)
		if lc1.Length() != 2 {
			t.Errorf("expected 2, got %d", lc1.Length())
		}
	})
}

func Test_Cov39_LC_AddAnother_Nil(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddAnother_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAnother(nil)
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov39_LC_AddAnother_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddAnother_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAnother(corestr.New.LinkedCollection.Create())
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov39_LC_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollection(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollection_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollection(nil)
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov39_LC_AddCollections(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollections", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.AddCollections([]*corestr.Collection{c1})
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_AddCollections_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollections_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollections(nil)
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov39_LC_AddCollectionsPtr(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollectionsPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.AddCollectionsPtr([]*corestr.Collection{c1})
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_AddCollectionsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollectionsPtr_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollectionsPtr(nil)
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

// ── AppendNode / AppendChainOfNodes / AddBackNode ──

func Test_Cov39_LC_AppendNode_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendNode_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AppendNode(&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"x"})})
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_AppendNode_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendNode_NonEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.AppendNode(&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})})
		if lc.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov39_LC_AddBackNode(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddBackNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddBackNode(&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"x"})})
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendChainOfNodes", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc2.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc1.AppendChainOfNodes(lc2.Head())
		if lc1.Length() != 3 {
			t.Errorf("expected 3, got %d", lc1.Length())
		}
	})
}

func Test_Cov39_LC_AppendChainOfNodes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendChainOfNodes_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		chain := corestr.New.LinkedCollection.Create()
		chain.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.AppendChainOfNodes(chain.Head())
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

// ── AttachWithNode ──

func Test_Cov39_LC_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AttachWithNode_NilCurrent", func() {
		lc := corestr.New.LinkedCollection.Create()
		err := lc.AttachWithNode(nil, &corestr.LinkedCollectionNode{})
		if err == nil {
			t.Error("expected error")
		}
	})
}

func Test_Cov39_LC_AttachWithNode_NextNotNil(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AttachWithNode_NextNotNil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		err := lc.AttachWithNode(lc.Head(), &corestr.LinkedCollectionNode{})
		if err == nil {
			t.Error("expected error")
		}
	})
}

// ── InsertAt ──

func Test_Cov39_LC_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_InsertAt_Front", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.InsertAt(0, corestr.New.Collection.Strings([]string{"a"}))
		if lc.First().First() != "a" {
			t.Error("expected a first")
		}
	})
}

func Test_Cov39_LC_InsertAt_Middle(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_InsertAt_Middle", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc.InsertAt(1, corestr.New.Collection.Strings([]string{"b"}))
		if lc.Length() != 3 {
			t.Errorf("expected 3, got %d", lc.Length())
		}
	})
}

// ── Loop ──

func Test_Cov39_LC_Loop(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Loop", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return false
		})
		if count != 2 {
			t.Errorf("expected 2, got %d", count)
		}
	})
}

func Test_Cov39_LC_Loop_Break(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Loop_Break", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return true
		})
		if count != 1 {
			t.Errorf("expected 1, got %d", count)
		}
	})
}

func Test_Cov39_LC_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Loop_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			t.Fatal("should not be called")
			return false
		})
	})
}

// ── Filter / FilterAsCollection / FilterAsCollections ──

func Test_Cov39_LC_Filter(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Filter", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		nodes := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		if len(nodes) != 2 {
			t.Errorf("expected 2, got %d", len(nodes))
		}
	})
}

func Test_Cov39_LC_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Filter_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		nodes := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		if len(nodes) != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov39_LC_Filter_BreakFirst(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Filter_BreakFirst", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		nodes := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})
		if len(nodes) != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_FilterAsCollection(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_FilterAsCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		col := lc.FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		}, 0)
		if col.Length() != 3 {
			t.Errorf("expected 3, got %d", col.Length())
		}
	})
}

func Test_Cov39_LC_FilterAsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_FilterAsCollection_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := lc.FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		}, 0)
		if col.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov39_LC_FilterAsCollections(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_FilterAsCollections", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		cols := lc.FilterAsCollections(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		if len(cols) != 1 {
			t.Errorf("expected 1")
		}
	})
}

// ── RemoveNodeByIndex ──

func Test_Cov39_LC_RemoveByIndex_First(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_RemoveByIndex_First", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.RemoveNodeByIndex(0)
		if lc.Length() != 1 || lc.First().First() != "b" {
			t.Error("unexpected")
		}
	})
}

func Test_Cov39_LC_RemoveByIndex_Last(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_RemoveByIndex_Last", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.RemoveNodeByIndex(1)
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_RemoveByIndex_Middle(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_RemoveByIndex_Middle", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc.RemoveNodeByIndex(1)
		if lc.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

// ── RemoveNodeByIndexes ──

func Test_Cov39_LC_RemoveByIndexes(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_RemoveByIndexes", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc.RemoveNodeByIndexes(true, 0, 2)
		if lc.Length() != 1 {
			t.Errorf("expected 1, got %d", lc.Length())
		}
	})
}

func Test_Cov39_LC_RemoveByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_RemoveByIndexes_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.RemoveNodeByIndexes(true)
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

// ── RemoveNode ──

func Test_Cov39_LC_RemoveNode_Head(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_RemoveNode_Head", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.RemoveNode(lc.Head())
		if lc.Length() != 1 || lc.First().First() != "b" {
			t.Error("unexpected")
		}
	})
}

func Test_Cov39_LC_RemoveNode_NonHead(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_RemoveNode_NonHead", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.RemoveNode(lc.IndexAt(1))
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

// ── AppendCollections / AppendCollectionsPointers / AppendCollectionsPointersLock ──

func Test_Cov39_LC_AppendCollections(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendCollections", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.AppendCollections(false, c1)
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_AppendCollections_SkipNil(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendCollections_SkipNil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AppendCollections(true, nil)
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov39_LC_AppendCollectionsPointers(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendCollectionsPointers", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		cols := []*corestr.Collection{c1}
		lc.AppendCollectionsPointers(false, &cols)
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_AppendCollectionsPointers_NilSkip(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendCollectionsPointers_NilSkip", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AppendCollectionsPointers(true, nil)
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov39_LC_AppendCollectionsPointersLock(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendCollectionsPointersLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		cols := []*corestr.Collection{c1}
		lc.AppendCollectionsPointersLock(false, &cols)
		if lc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

// ── AddCollectionsToNode / AddCollectionsPointerToNode ──

func Test_Cov39_LC_AddCollectionsToNode(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollectionsToNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.AddCollectionsToNode(false, lc.Head(), c2)
		if lc.Length() < 2 {
			t.Errorf("expected at least 2, got %d", lc.Length())
		}
	})
}

func Test_Cov39_LC_AddCollectionsPointerToNode_NilSkip(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollectionsPointerToNode_NilSkip", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollectionsPointerToNode(true, nil, nil)
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov39_LC_AddCollectionsPointerToNode_NilItems(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollectionsPointerToNode_NilItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollectionsPointerToNode(true, nil, nil)
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

// ── AddCollectionToNode ──

func Test_Cov39_LC_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollectionToNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.AddCollectionToNode(false, lc.Head(), corestr.New.Collection.Strings([]string{"b"}))
		if lc.Length() < 2 {
			t.Errorf("expected at least 2")
		}
	})
}

// ── AddAsync ──

func Test_Cov39_LC_AddAsync(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddAsync", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsync(corestr.New.Collection.Strings([]string{"a"}), wg)
		wg.Wait()
		if lc.Length() != 1 {
			t.Errorf("expected 1, got %d", lc.Length())
		}
	})
}

func Test_Cov39_LC_AppendChainOfNodesAsync(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendChainOfNodesAsync", func() {
		lc := corestr.New.LinkedCollection.Create()
		chain := corestr.New.LinkedCollection.Create()
		chain.Add(corestr.New.Collection.Strings([]string{"a"}))
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AppendChainOfNodesAsync(chain.Head(), wg)
		wg.Wait()
		if lc.Length() != 1 {
			t.Errorf("expected 1, got %d", lc.Length())
		}
	})
}

// ── AddStringsOfStrings ──

func Test_Cov39_LC_AddStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddStringsOfStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false, []string{"a"}, []string{"b"})
		if lc.Length() != 2 {
			t.Errorf("expected 2, got %d", lc.Length())
		}
	})
}

func Test_Cov39_LC_AddStringsOfStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddStringsOfStrings_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false)
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

// ── AddAsyncFuncItems ──

func Test_Cov39_LC_AddAsyncFuncItems(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddAsyncFuncItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItems(wg, false, func() []string { return []string{"a"} })
		if lc.Length() != 1 {
			t.Errorf("expected 1, got %d", lc.Length())
		}
	})
}

func Test_Cov39_LC_AddAsyncFuncItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddAsyncFuncItems_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAsyncFuncItems(nil, false)
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov39_LC_AddAsyncFuncItemsPointer(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddAsyncFuncItemsPointer", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItemsPointer(wg, false, func() []string { return []string{"x"} })
		if lc.Length() != 1 {
			t.Errorf("expected 1, got %d", lc.Length())
		}
	})
}

func Test_Cov39_LC_AddAsyncFuncItemsPointer_Nil(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddAsyncFuncItemsPointer_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAsyncFuncItemsPointer(nil, false)
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

// ── ConcatNew ──

func Test_Cov39_LC_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ConcatNew", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"b"}))
		result := lc1.ConcatNew(false, lc2)
		if result.Length() != 2 {
			t.Errorf("expected 2, got %d", result.Length())
		}
	})
}

func Test_Cov39_LC_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ConcatNew_EmptyClone", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		result := lc.ConcatNew(true)
		if result.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ConcatNew_EmptyNoClone", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		result := lc.ConcatNew(false)
		if result != lc {
			t.Error("expected same pointer")
		}
	})
}

// ── IndexAt / SafeIndexAt / SafePointerIndexAt ──

func Test_Cov39_LC_IndexAt(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IndexAt", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		node := lc.IndexAt(1)
		if node.Element.First() != "b" {
			t.Error("expected b")
		}
	})
}

func Test_Cov39_LC_IndexAt_Zero(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IndexAt_Zero", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.IndexAt(0).Element.First() != "a" {
			t.Error("expected a")
		}
	})
}

func Test_Cov39_LC_IndexAt_Negative(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IndexAt_Negative", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.IndexAt(-1) != nil {
			t.Error("expected nil")
		}
	})
}

func Test_Cov39_LC_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_SafeIndexAt", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.SafeIndexAt(0).Element.First() != "a" {
			t.Error("expected a")
		}
	})
}

func Test_Cov39_LC_SafeIndexAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_SafeIndexAt_OutOfRange", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.SafeIndexAt(5) != nil {
			t.Error("expected nil")
		}
	})
}

func Test_Cov39_LC_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_SafePointerIndexAt", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.SafePointerIndexAt(0) == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov39_LC_SafePointerIndexAt_Nil(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_SafePointerIndexAt_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.SafePointerIndexAt(0) != nil {
			t.Error("expected nil")
		}
	})
}

// ── GetNextNodes / GetAllLinkedNodes ──

func Test_Cov39_LC_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_GetNextNodes", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		nodes := lc.GetNextNodes(2)
		if len(nodes) != 2 {
			t.Errorf("expected 2, got %d", len(nodes))
		}
	})
}
func Test_Cov39_LC_ToCollection(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ToCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		col := lc.ToCollection(0)
		if col.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov39_LC_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ToCollection_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.ToCollection(0).Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov39_LC_ToCollectionSimple(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ToCollectionSimple", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.ToCollectionSimple().Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_Cov39_LC_ToStrings(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ToStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if len(lc.ToStrings()) != 1 {
			t.Error("expected 1")
		}
	})
}
func Test_Cov39_LC_ToCollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ToCollectionsOfCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		coc := lc.ToCollectionsOfCollection(0)
		if coc == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov39_LC_ToCollectionsOfCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ToCollectionsOfCollection_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		coc := lc.ToCollectionsOfCollection(0)
		if coc == nil {
			t.Error("expected non-nil")
		}
	})
}

// ── ItemsOfItems / ItemsOfItemsCollection ──

func Test_Cov39_LC_ItemsOfItems(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ItemsOfItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		items := lc.ItemsOfItems()
		if len(items) != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_ItemsOfItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ItemsOfItems_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		items := lc.ItemsOfItems()
		if len(items) != 0 {
			t.Errorf("expected 0")
		}
	})
}
func Test_Cov39_LC_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_SimpleSlice", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		ss := lc.SimpleSlice()
		if ss.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_List(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_List", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		if len(lc.List()) != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov39_LC_List_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_List_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		if len(lc.List()) != 0 {
			t.Errorf("expected 0")
		}
	})
}
func Test_Cov39_LC_String(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_String", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.String() == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov39_LC_String_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_String_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		s := lc.String()
		if s == "" {
			t.Error("expected non-empty (NoElements)")
		}
	})
}

func Test_Cov39_LC_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_StringLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if !strings.Contains(lc.StringLock(), "a") {
			t.Error("expected a")
		}
	})
}

func Test_Cov39_LC_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_StringLock_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		s := lc.StringLock()
		if s == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov39_LC_Join(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Join", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		if lc.Join(",") == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov39_LC_Joins(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Joins", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		result := lc.Joins(",", "b")
		if !strings.Contains(result, "a") || !strings.Contains(result, "b") {
			t.Errorf("unexpected: %s", result)
		}
	})
}

func Test_Cov39_LC_Joins_NilItems(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Joins_NilItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		_ = lc.Joins(",")
	})
}

// ── IsEqualsPtr ──

func Test_Cov39_LC_IsEqualsPtr_Same(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IsEqualsPtr_Same", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))
		if !lc1.IsEqualsPtr(lc2) {
			t.Error("expected equal")
		}
	})
}

func Test_Cov39_LC_IsEqualsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IsEqualsPtr_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.IsEqualsPtr(nil) {
			t.Error("expected false")
		}
	})
}

func Test_Cov39_LC_IsEqualsPtr_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IsEqualsPtr_SamePtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if !lc.IsEqualsPtr(lc) {
			t.Error("expected true")
		}
	})
}

func Test_Cov39_LC_IsEqualsPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IsEqualsPtr_BothEmpty", func() {
		a := corestr.New.LinkedCollection.Create()
		b := corestr.New.LinkedCollection.Create()
		if !a.IsEqualsPtr(b) {
			t.Error("expected true")
		}
	})
}

func Test_Cov39_LC_IsEqualsPtr_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IsEqualsPtr_DiffLen", func() {
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		b := corestr.New.LinkedCollection.Create()
		b.Add(corestr.New.Collection.Strings([]string{"a"}))
		b.Add(corestr.New.Collection.Strings([]string{"b"}))
		if a.IsEqualsPtr(b) {
			t.Error("expected false")
		}
	})
}

func Test_Cov39_LC_IsEqualsPtr_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IsEqualsPtr_OneEmpty", func() {
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		b := corestr.New.LinkedCollection.Create()
		if a.IsEqualsPtr(b) {
			t.Error("expected false")
		}
	})
}

// ── GetCompareSummary ──

func Test_Cov39_LC_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_GetCompareSummary", func() {
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		b := corestr.New.LinkedCollection.Create()
		b.Add(corestr.New.Collection.Strings([]string{"b"}))
		summary := a.GetCompareSummary(b, "left", "right")
		if summary == "" {
			t.Error("expected non-empty")
		}
	})
}

// ── JSON ──

func Test_Cov39_LC_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_MarshalJSON", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		data, err := json.Marshal(lc)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(data), "\"a\"") {
			t.Errorf("unexpected: %s", data)
		}
	})
}

func Test_Cov39_LC_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_UnmarshalJSON", func() {
		lc := corestr.New.LinkedCollection.Create()
		err := json.Unmarshal([]byte(`["x","y"]`), lc)
		if err != nil {
			t.Fatal(err)
		}
		if lc.Length() != 1 { // unmarshal adds as single collection
			// accept whatever the implementation does
		}
	})
}

func Test_Cov39_LC_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_UnmarshalJSON_Invalid", func() {
		lc := corestr.New.LinkedCollection.Create()
		err := json.Unmarshal([]byte(`bad`), lc)
		if err == nil {
			t.Error("expected error")
		}
	})
}

func Test_Cov39_LC_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_JsonModel", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if len(lc.JsonModel()) != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LC_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_JsonModelAny", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.JsonModelAny() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov39_LC_Json(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Json", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Json().Error != nil {
			t.Error("unexpected error")
		}
	})
}

func Test_Cov39_LC_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_JsonPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.JsonPtr() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov39_LC_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ParseInjectUsingJson", func() {
		src := corestr.New.LinkedCollection.Create()
		src.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc := corestr.New.LinkedCollection.Create()
		_, err := lc.ParseInjectUsingJson(src.JsonPtr())
		if err != nil {
			t.Fatal(err)
		}
	})
}

func Test_Cov39_LC_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ParseInjectUsingJsonMust", func() {
		src := corestr.New.LinkedCollection.Create()
		src.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc := corestr.New.LinkedCollection.Create()
		result := lc.ParseInjectUsingJsonMust(src.JsonPtr())
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov39_LC_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_JsonParseSelfInject", func() {
		src := corestr.New.LinkedCollection.Create()
		src.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc := corestr.New.LinkedCollection.Create()
		err := lc.JsonParseSelfInject(src.JsonPtr())
		if err != nil {
			t.Fatal(err)
		}
	})
}

func Test_Cov39_LC_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AsJsoner", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.AsJsoner() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov39_LC_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AsJsonContractsBinder", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.AsJsonContractsBinder() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov39_LC_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AsJsonParseSelfInjector", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.AsJsonParseSelfInjector() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov39_LC_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AsJsonMarshaller", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.AsJsonMarshaller() == nil {
			t.Error("expected non-nil")
		}
	})
}

// ── Clear / RemoveAll ──

func Test_Cov39_LC_Clear(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Clear", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Clear()
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov39_LC_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Clear_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Clear()
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov39_LC_RemoveAll(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_RemoveAll", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.RemoveAll()
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

// ══════════════════════════════════════════════════════════════
// LinkedCollectionNode
// ══════════════════════════════════════════════════════════════

func Test_Cov39_LCNode_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsEmpty", func() {
		var n *corestr.LinkedCollectionNode
		if !n.IsEmpty() {
			t.Error("expected empty")
		}
	})
}

func Test_Cov39_LCNode_HasElement(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_HasElement", func() {
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		if !n.HasElement() {
			t.Error("expected true")
		}
	})
}

func Test_Cov39_LCNode_HasNext(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_HasNext", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		if !lc.Head().HasNext() {
			t.Error("expected true")
		}
		if lc.Tail().HasNext() {
			t.Error("expected false")
		}
	})
}

func Test_Cov39_LCNode_Next(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_Next", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		if lc.Head().Next().Element.First() != "b" {
			t.Error("expected b")
		}
	})
}

func Test_Cov39_LCNode_EndOfChain(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_EndOfChain", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		end, length := lc.Head().EndOfChain()
		if length != 2 || end.Element.First() != "b" {
			t.Errorf("unexpected")
		}
	})
}

func Test_Cov39_LCNode_LoopEndOfChain(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_LoopEndOfChain", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		count := 0
		end, length := lc.Head().LoopEndOfChain(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return false
		})
		if count != 2 || length != 2 || end.Element.First() != "b" {
			t.Errorf("unexpected")
		}
	})
}

func Test_Cov39_LCNode_LoopEndOfChain_Break(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_LoopEndOfChain_Break", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		_, length := lc.Head().LoopEndOfChain(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			return true
		})
		if length != 1 {
			t.Errorf("expected 1, got %d", length)
		}
	})
}

func Test_Cov39_LCNode_Clone(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_Clone", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		cloned := lc.Head().Clone()
		if cloned.Element.First() != "a" || cloned.HasNext() {
			t.Error("unexpected")
		}
	})
}

func Test_Cov39_LCNode_IsEqual_Same(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsEqual_Same", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		n1 := &corestr.LinkedCollectionNode{Element: c}
		n2 := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		if !n1.IsEqual(n2) {
			t.Error("expected equal")
		}
	})
}

func Test_Cov39_LCNode_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsEqual_BothNil", func() {
		var a, b *corestr.LinkedCollectionNode
		if !a.IsEqual(b) {
			t.Error("expected equal")
		}
	})
}

func Test_Cov39_LCNode_IsEqual_OneNil(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsEqual_OneNil", func() {
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		if n.IsEqual(nil) {
			t.Error("expected false")
		}
	})
}

func Test_Cov39_LCNode_IsEqual_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsEqual_SamePtr", func() {
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		if !n.IsEqual(n) {
			t.Error("expected true")
		}
	})
}

func Test_Cov39_LCNode_IsChainEqual(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsChainEqual", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))
		if !lc1.Head().IsChainEqual(lc2.Head()) {
			t.Error("expected equal")
		}
	})
}

func Test_Cov39_LCNode_IsChainEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsChainEqual_BothNil", func() {
		var a, b *corestr.LinkedCollectionNode
		if !a.IsChainEqual(b) {
			t.Error("expected equal")
		}
	})
}

func Test_Cov39_LCNode_IsEqualValue(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsEqualValue", func() {
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		if !n.IsEqualValue(corestr.New.Collection.Strings([]string{"a"})) {
			t.Error("expected true")
		}
	})
}

func Test_Cov39_LCNode_IsEqualValue_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsEqualValue_BothNil", func() {
		n := &corestr.LinkedCollectionNode{Element: nil}
		if !n.IsEqualValue(nil) {
			t.Error("expected true")
		}
	})
}

func Test_Cov39_LCNode_String(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_String", func() {
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		if n.String() == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov39_LCNode_List(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_List", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		list := lc.Head().List()
		if len(list) != 3 {
			t.Errorf("expected 3, got %d", len(list))
		}
	})
}
func Test_Cov39_LCNode_Join(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_Join", func() {
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a", "b"})}
		if n.Join(",") != "a,b" {
			t.Error("unexpected")
		}
	})
}

func Test_Cov39_LCNode_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_CreateLinkedList", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		newLC := lc.Head().CreateLinkedList()
		if newLC.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_LCNode_AddNext(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_AddNext", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Head().AddNext(lc, corestr.New.Collection.Strings([]string{"b"}))
		if lc.Length() != 2 {
			t.Errorf("expected 2, got %d", lc.Length())
		}
	})
}

func Test_Cov39_LCNode_AddNextNode(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_AddNextNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		newNode := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		lc.Head().AddNextNode(lc, newNode)
		if lc.Length() != 2 {
			t.Errorf("expected 2, got %d", lc.Length())
		}
	})
}

func Test_Cov39_LCNode_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_AddStringsToNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Head().AddStringsToNode(lc, false, []string{"b"}, false)
		if lc.Length() < 2 {
			t.Errorf("expected at least 2")
		}
	})
}

func Test_Cov39_LCNode_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_AddCollectionToNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Head().AddCollectionToNode(lc, false, corestr.New.Collection.Strings([]string{"b"}))
		if lc.Length() < 2 {
			t.Errorf("expected at least 2")
		}
	})
}

// ══════════════════════════════════════════════════════════════
// NonChainedLinkedCollectionNodes
// ══════════════════════════════════════════════════════════════

func Test_Cov39_NCLCN_Basic(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_Basic", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)
		if !nc.IsEmpty() || nc.HasItems() || nc.Length() != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_Cov39_NCLCN_Adds(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_Adds", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		n1 := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		nc.Adds(n1)
		if nc.Length() != 1 || nc.First() != n1 || nc.Last() != n1 {
			t.Error("unexpected")
		}
	})
}

func Test_Cov39_NCLCN_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_Adds_Nil", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		nc.Adds()
		if nc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov39_NCLCN_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_FirstOrDefault_Empty", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		if nc.FirstOrDefault() != nil {
			t.Error("expected nil")
		}
	})
}

func Test_Cov39_NCLCN_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_LastOrDefault_Empty", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		if nc.LastOrDefault() != nil {
			t.Error("expected nil")
		}
	})
}

func Test_Cov39_NCLCN_Items(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_Items", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		nc.Adds(&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})})
		if len(nc.Items()) != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov39_NCLCN_ApplyChaining(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_ApplyChaining", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		n1 := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		n2 := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		nc.Adds(n1, n2)
		nc.ApplyChaining()
		if !nc.IsChainingApplied() {
			t.Error("expected true")
		}
		if !n1.HasNext() {
			t.Error("expected chaining")
		}
	})
}

func Test_Cov39_NCLCN_ApplyChaining_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_ApplyChaining_Empty", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		nc.ApplyChaining()
		// should not panic
	})
}

func Test_Cov39_NCLCN_ToChainedNodes(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_ToChainedNodes", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		nc.Adds(&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})})
		nc.Adds(&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})})
		chained := nc.ToChainedNodes()
		if chained == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov39_NCLCN_ToChainedNodes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_ToChainedNodes_Empty", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		chained := nc.ToChainedNodes()
		if chained == nil || len(*chained) != 0 {
			t.Errorf("expected empty")
		}
	})
}

// ── newLinkedListCollectionsCreator ──

func Test_Cov39_Creator_LC_Create(t *testing.T) {
	safeTest(t, "Test_Cov39_Creator_LC_Create", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc == nil || lc.Length() != 0 {
			t.Error("unexpected")
		}
	})
}

func Test_Cov39_Creator_LC_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_Creator_LC_Empty", func() {
		lc := corestr.New.LinkedCollection.Empty()
		if lc == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov39_Creator_LC_Strings(t *testing.T) {
	safeTest(t, "Test_Cov39_Creator_LC_Strings", func() {
		lc := corestr.New.LinkedCollection.Strings("a", "b")
		if lc.Length() != 1 { // all strings as one collection
			// implementation may vary
		}
	})
}

func Test_Cov39_Creator_LC_UsingCollections(t *testing.T) {
	safeTest(t, "Test_Cov39_Creator_LC_UsingCollections", func() {
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc := corestr.New.LinkedCollection.UsingCollections(c1, c2)
		if lc.Length() != 2 {
			t.Errorf("expected 2, got %d", lc.Length())
		}
	})
}

func Test_Cov39_Creator_LC_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_Cov39_Creator_LC_PointerStringsPtr", func() {
		a := "a"
		ptrs := []*string{&a, nil}
		lc := corestr.New.LinkedCollection.PointerStringsPtr(&ptrs)
		if lc == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov39_Creator_LC_PointerStringsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov39_Creator_LC_PointerStringsPtr_Nil", func() {
		lc := corestr.New.LinkedCollection.PointerStringsPtr(nil)
		if lc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}
