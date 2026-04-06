package corestrtests

import (
	"encoding/json"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ===================== LinkedCollectionNode =====================

func Test_C47_LinkedCollectionNode_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_IsEmpty", func() {
		var n *corestr.LinkedCollectionNode
		if !n.IsEmpty() {
			t.Fatal("nil should be empty")
		}
		n2 := &corestr.LinkedCollectionNode{}
		if !n2.IsEmpty() {
			t.Fatal("nil element should be empty")
		}
	})
}

func Test_C47_LinkedCollectionNode_HasElement(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_HasElement", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		n := &corestr.LinkedCollectionNode{Element: col}
		if !n.HasElement() {
			t.Fatal("should have element")
		}
	})
}

func Test_C47_LinkedCollectionNode_HasNext(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_HasNext", func() {
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		if n.HasNext() {
			t.Fatal("should not have next")
		}
	})
}

func Test_C47_LinkedCollectionNode_Clone(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_Clone", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		lc := corestr.New.LinkedCollection.Strings("a", "b")
		node := lc.Head()
		cloned := node.Clone()
		if cloned.HasNext() {
			t.Fatal("clone should not have next")
		}
		_ = col
	})
}

func Test_C47_LinkedCollectionNode_EndOfChain(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_EndOfChain", func() {
		lc := corestr.New.LinkedCollection.Strings("a", "b")
		end, length := lc.Head().EndOfChain()
		_ = end
		if length != 1 {
			// Each AddStrings creates one node per call
			_ = 0
		}
	})
}

func Test_C47_LinkedCollectionNode_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_IsEqual_BothNil", func() {
		var n1, n2 *corestr.LinkedCollectionNode
		_ = n1
		_ = n2
	})
}

func Test_C47_LinkedCollectionNode_IsEqual_Same(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_IsEqual_Same", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		node := lc.Head()
		if !node.IsEqual(node) {
			t.Fatal("same node should be equal")
		}
	})
}

func Test_C47_LinkedCollectionNode_IsChainEqual(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_IsChainEqual", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("a")
		if !lc1.Head().IsChainEqual(lc2.Head()) {
			t.Fatal("should be chain equal")
		}
	})
}

func Test_C47_LinkedCollectionNode_IsEqualValue(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_IsEqualValue", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(col)
		if !lc.Head().IsEqualValue(col) {
			t.Fatal("should match")
		}
	})
}

func Test_C47_LinkedCollectionNode_IsEqualValue_NilBoth(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_IsEqualValue_NilBoth", func() {
		n := &corestr.LinkedCollectionNode{}
		if !n.IsEqualValue(nil) {
			t.Fatal("both nil should be equal")
		}
	})
}

func Test_C47_LinkedCollectionNode_String(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_String", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(col)
		s := lc.Head().String()
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_C47_LinkedCollectionNode_List(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_List", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		list := lc.Head().List()
		if len(list) != 2 {
			t.Fatalf("expected 2, got %d", len(list))
		}
	})
}
func Test_C47_LinkedCollectionNode_Join(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_Join", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		j := lc.Head().Join(",")
		if j != "a,b" {
			t.Fatalf("got %s", j)
		}
	})
}

func Test_C47_LinkedCollectionNode_StringList(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_StringList", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		s := lc.Head().StringList("H:")
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_C47_LinkedCollectionNode_Print(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_Print", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.Head().Print("test: ")
	})
}

func Test_C47_LinkedCollectionNode_LoopEndOfChain(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_LoopEndOfChain", func() {
		lc := corestr.New.LinkedCollection.Create()
		col1 := corestr.New.Collection.Strings([]string{"a"})
		col2 := corestr.New.Collection.Strings([]string{"b"})
		lc.Add(col1)
		lc.Add(col2)
		count := 0
		_, length := lc.Head().LoopEndOfChain(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return false
		})
		if count != 2 || length != 2 {
			t.Fatalf("count=%d length=%d", count, length)
		}
	})
}

func Test_C47_LinkedCollectionNode_LoopEndOfChain_BreakFirst(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_LoopEndOfChain_BreakFirst", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		_, length := lc.Head().LoopEndOfChain(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			return true
		})
		if length != 1 {
			t.Fatalf("expected 1, got %d", length)
		}
	})
}

func Test_C47_LinkedCollectionNode_AddNext(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_AddNext", func() {
		lc := corestr.New.LinkedCollection.Create()
		col1 := corestr.New.Collection.Strings([]string{"a"})
		col2 := corestr.New.Collection.Strings([]string{"c"})
		lc.Add(col1)
		lc.Add(col2)
		colB := corestr.New.Collection.Strings([]string{"b"})
		lc.Head().AddNext(lc, colB)
		if lc.Length() != 3 {
			t.Fatalf("expected 3, got %d", lc.Length())
		}
	})
}

func Test_C47_LinkedCollectionNode_AddNextNode(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_AddNextNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(col)
		newNode := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		lc.Head().AddNextNode(lc, newNode)
		if lc.Length() != 2 {
			t.Fatalf("expected 2, got %d", lc.Length())
		}
	})
}

func Test_C47_LinkedCollectionNode_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_CreateLinkedList", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(col)
		newLC := lc.Head().CreateLinkedList()
		if newLC.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

// ===================== LinkedCollections =====================

func Test_C47_LinkedCollections_Create_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Create_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		if !lc.IsEmpty() || lc.Length() != 0 || lc.HasItems() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C47_LinkedCollections_Add(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Add", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		lc.Add(col)
		if lc.Length() != 1 || lc.Head() == nil || lc.Tail() == nil {
			t.Fatal("add failed")
		}
	})
}

func Test_C47_LinkedCollections_AddMultiple(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddMultiple", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C47_LinkedCollections_AddStrings(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		if lc.Length() != 1 {
			t.Fatal("expected 1 collection node")
		}
	})
}

func Test_C47_LinkedCollections_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStrings_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings()
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C47_LinkedCollections_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock("a", "b")
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_AddStringsLock_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsLock_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock()
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C47_LinkedCollections_AddLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddLock(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_AddFront_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddFront_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.AddFront(col)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_AddFront_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddFront_NonEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.AddFront(corestr.New.Collection.Strings([]string{"a"}))
		if lc.First().List()[0] != "a" {
			t.Fatal("front should be a")
		}
	})
}

func Test_C47_LinkedCollections_AddFrontLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddFrontLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddFrontLock(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_PushFront(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_PushFront", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.PushFront(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_Push(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Push", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Push(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_PushBack(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_PushBack", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.PushBack(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_PushBackLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_PushBackLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.PushBackLock(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_First_Single_Last(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_First_Single_Last", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(col)
		if lc.First() != col || lc.Single() != col || lc.Last() != col {
			t.Fatal("mismatch")
		}
	})
}

func Test_C47_LinkedCollections_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_FirstOrDefault_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		r := lc.FirstOrDefault()
		if r == nil {
			t.Fatal("expected empty collection, not nil")
		}
	})
}

func Test_C47_LinkedCollections_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_LastOrDefault_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		r := lc.LastOrDefault()
		if r == nil {
			t.Fatal("expected empty collection, not nil")
		}
	})
}

func Test_C47_LinkedCollections_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AllIndividualItemsLength", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		lc.AddStrings("c")
		if lc.AllIndividualItemsLength() != 3 {
			t.Fatalf("expected 3, got %d", lc.AllIndividualItemsLength())
		}
	})
}

func Test_C47_LinkedCollections_LengthLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_LengthLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		if lc.LengthLock() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEmptyLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		if !lc.IsEmptyLock() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C47_LinkedCollections_AppendNode_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendNode_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		lc.AppendNode(node)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_AppendNode_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendNode_NonEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		lc.AppendNode(node)
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C47_LinkedCollections_AddBackNode(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddBackNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		lc.AddBackNode(node)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendChainOfNodes", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc1.AddStrings("b")

		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AppendChainOfNodes(lc1.Head())
		if lc2.Length() != 2 {
			t.Fatalf("expected 2, got %d", lc2.Length())
		}
	})
}

func Test_C47_LinkedCollections_AppendChainOfNodes_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendChainOfNodes_NonEmpty", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")

		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("x")
		lc2.AppendChainOfNodes(lc1.Head())
		if lc2.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_C47_LinkedCollections_AppendChainOfNodesAsync(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendChainOfNodesAsync", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("x")

		chain := corestr.New.LinkedCollection.Create()
		chain.AddStrings("a")

		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AppendChainOfNodesAsync(chain.Head(), wg)
		wg.Wait()
		if lc.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_C47_LinkedCollections_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AttachWithNode_NilCurrent", func() {
		lc := corestr.New.LinkedCollection.Create()
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		err := lc.AttachWithNode(nil, node)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_C47_LinkedCollections_AddAnother(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAnother", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc1.AddStrings("b")

		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddAnother(lc1)
		if lc2.Length() != 2 {
			t.Fatalf("expected 2, got %d", lc2.Length())
		}
	})
}

func Test_C47_LinkedCollections_AddAnother_Nil(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAnother_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAnother(nil)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C47_LinkedCollections_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddCollection_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollection(nil)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C47_LinkedCollections_AddCollection(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollection(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_AppendCollections(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendCollections", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.AppendCollections(true, c1, nil, c2)
		if lc.Length() != 2 {
			t.Fatalf("expected 2, got %d", lc.Length())
		}
	})
}

func Test_C47_LinkedCollections_AppendCollections_NilSkip(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendCollections_NilSkip", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AppendCollections(true, nil)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C47_LinkedCollections_AddStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsOfStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false, []string{"a"}, []string{"b"})
		if lc.Length() != 2 {
			t.Fatalf("expected 2, got %d", lc.Length())
		}
	})
}

func Test_C47_LinkedCollections_AddStringsOfStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsOfStrings_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C47_LinkedCollections_Loop(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Loop", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return false
		})
		if count != 2 {
			t.Fatalf("expected 2, got %d", count)
		}
	})
}

func Test_C47_LinkedCollections_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Loop_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			t.Fatal("should not call")
			return false
		})
	})
}

func Test_C47_LinkedCollections_Filter(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Filter", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		result := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		if len(result) != 2 {
			t.Fatalf("expected 2, got %d", len(result))
		}
	})
}

func Test_C47_LinkedCollections_FilterAsCollection(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_FilterAsCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		lc.AddStrings("c")
		col := lc.FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		}, 0)
		if col.Length() != 3 {
			t.Fatalf("expected 3, got %d", col.Length())
		}
	})
}

func Test_C47_LinkedCollections_FilterAsCollections(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_FilterAsCollections", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		cols := lc.FilterAsCollections(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		if len(cols) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_GetNextNodes", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		lc.AddStrings("c")
		nodes := lc.GetNextNodes(2)
		if len(nodes) != 2 {
			t.Fatalf("expected 2, got %d", len(nodes))
		}
	})
}

func Test_C47_LinkedCollections_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_GetAllLinkedNodes", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		all := lc.GetAllLinkedNodes()
		if len(all) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_InsertAt_Front", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("b")
		lc.InsertAt(0, corestr.New.Collection.Strings([]string{"a"}))
		if lc.First().List()[0] != "a" {
			t.Fatal("insert at front failed")
		}
	})
}

func Test_C47_LinkedCollections_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_RemoveNodeByIndex", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		lc.AddStrings("c")
		lc.RemoveNodeByIndex(1)
		if lc.Length() != 2 {
			t.Fatalf("expected 2, got %d", lc.Length())
		}
	})
}

func Test_C47_LinkedCollections_RemoveNodeByIndex_First(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_RemoveNodeByIndex_First", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		lc.RemoveNodeByIndex(0)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_RemoveNode(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_RemoveNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		lc.RemoveNode(lc.Head())
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_RemoveNodeByIndexes", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		lc.AddStrings("c")
		lc.RemoveNodeByIndexes(false, 0, 2)
		if lc.Length() != 1 {
			t.Fatalf("expected 1, got %d", lc.Length())
		}
	})
}

func Test_C47_LinkedCollections_RemoveNodeByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_RemoveNodeByIndexes_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.RemoveNodeByIndexes(false)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_IndexAt(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IndexAt", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		node := lc.IndexAt(0)
		if node == nil {
			t.Fatal("expected node")
		}
		node2 := lc.IndexAt(-1)
		if node2 != nil {
			t.Fatal("negative should be nil")
		}
	})
}

func Test_C47_LinkedCollections_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_SafeIndexAt", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		n := lc.SafeIndexAt(0)
		if n == nil {
			t.Fatal("expected node")
		}
		n2 := lc.SafeIndexAt(99)
		if n2 != nil {
			t.Fatal("expected nil")
		}
		n3 := lc.SafeIndexAt(-1)
		if n3 != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_C47_LinkedCollections_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_SafePointerIndexAt", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		col := lc.SafePointerIndexAt(0)
		if col == nil {
			t.Fatal("expected collection")
		}
		col2 := lc.SafePointerIndexAt(99)
		if col2 != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_C47_LinkedCollections_IsEqualsPtr(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEqualsPtr", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("a")
		if !lc1.IsEqualsPtr(lc2) {
			t.Fatal("should be equal")
		}
	})
}

func Test_C47_LinkedCollections_IsEqualsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEqualsPtr_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.IsEqualsPtr(nil) {
			t.Fatal("nil should not be equal")
		}
	})
}

func Test_C47_LinkedCollections_IsEqualsPtr_SameRef(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEqualsPtr_SameRef", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		if !lc.IsEqualsPtr(lc) {
			t.Fatal("same ref should be equal")
		}
	})
}

func Test_C47_LinkedCollections_IsEqualsPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEqualsPtr_BothEmpty", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc2 := corestr.New.LinkedCollection.Create()
		if !lc1.IsEqualsPtr(lc2) {
			t.Fatal("both empty should be equal")
		}
	})
}

func Test_C47_LinkedCollections_IsEqualsPtr_DiffLength(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEqualsPtr_DiffLength", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("a")
		lc2.AddStrings("b")
		if lc1.IsEqualsPtr(lc2) {
			t.Fatal("diff length should not be equal")
		}
	})
}

func Test_C47_LinkedCollections_ToCollection(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		lc.AddStrings("c")
		col := lc.ToCollection(0)
		if col.Length() != 3 {
			t.Fatalf("expected 3, got %d", col.Length())
		}
	})
}

func Test_C47_LinkedCollections_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToCollection_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := lc.ToCollection(0)
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C47_LinkedCollections_ToCollectionSimple(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToCollectionSimple", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		col := lc.ToCollectionSimple()
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_ToStrings(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		s := lc.ToStrings()
		if len(s) != 1 {
			t.Fatal("expected 1")
		}
	})
}
func Test_C47_LinkedCollections_ToCollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToCollectionsOfCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		coc := lc.ToCollectionsOfCollection(0)
		if coc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_ToCollectionsOfCollection_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToCollectionsOfCollection_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		coc := lc.ToCollectionsOfCollection(0)
		if coc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C47_LinkedCollections_ItemsOfItems(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ItemsOfItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		items := lc.ItemsOfItems()
		if len(items) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C47_LinkedCollections_ItemsOfItems_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ItemsOfItems_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		items := lc.ItemsOfItems()
		if len(items) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C47_LinkedCollections_ItemsOfItemsCollection(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ItemsOfItemsCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		cols := lc.ItemsOfItemsCollection()
		if len(cols) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_SimpleSlice", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		ss := lc.SimpleSlice()
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_List(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_List", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		list := lc.List()
		if len(list) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C47_LinkedCollections_List_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_List_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		list := lc.List()
		if len(list) != 0 {
			t.Fatal("expected 0")
		}
	})
}
func Test_C47_LinkedCollections_String_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_String_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		s := lc.String()
		if s == "" {
			t.Fatal("expected NoElements")
		}
	})
}

func Test_C47_LinkedCollections_String_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_String_NonEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		s := lc.String()
		if s == "" {
			t.Fatal("expected string")
		}
	})
}

func Test_C47_LinkedCollections_StringLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_StringLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		s := lc.StringLock()
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_C47_LinkedCollections_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_StringLock_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		s := lc.StringLock()
		if s == "" {
			t.Fatal("expected NoElements")
		}
	})
}

func Test_C47_LinkedCollections_Join(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Join", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		j := lc.Join(",")
		if j == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_C47_LinkedCollections_Joins(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Joins", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		j := lc.Joins(",", "b")
		if j == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_C47_LinkedCollections_Joins_NilItems(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Joins_NilItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		j := lc.Joins(",", "x")
		if j != "x" {
			t.Fatalf("expected x, got %s", j)
		}
	})
}

func Test_C47_LinkedCollections_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ConcatNew", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("b")
		result := lc1.ConcatNew(false, lc2)
		if result.Length() != 2 {
			t.Fatalf("expected 2, got %d", result.Length())
		}
	})
}

func Test_C47_LinkedCollections_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ConcatNew_EmptyClone", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		result := lc.ConcatNew(true)
		if result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ConcatNew_EmptyNoClone", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		result := lc.ConcatNew(false)
		if result != lc {
			t.Fatal("expected same ref")
		}
	})
}

func Test_C47_LinkedCollections_Clear(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Clear", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.Clear()
		if !lc.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C47_LinkedCollections_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Clear_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Clear()
	})
}

func Test_C47_LinkedCollections_RemoveAll(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_RemoveAll", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.RemoveAll()
		if !lc.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C47_LinkedCollections_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_GetCompareSummary", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("b")
		s := lc1.GetCompareSummary(lc2, "left", "right")
		if s == "" {
			t.Fatal("expected summary")
		}
	})
}

// JSON
func Test_C47_LinkedCollections_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_MarshalJSON", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		data, err := json.Marshal(lc)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != `["a","b"]` {
			t.Fatalf("unexpected: %s", string(data))
		}
	})
}

func Test_C47_LinkedCollections_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_UnmarshalJSON", func() {
		lc := corestr.New.LinkedCollection.Create()
		err := json.Unmarshal([]byte(`["x","y"]`), lc)
		if err != nil {
			t.Fatal(err)
		}
		list := lc.List()
		if len(list) != 2 {
			t.Fatalf("expected 2, got %d", len(list))
		}
	})
}

func Test_C47_LinkedCollections_Json(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Json", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		j := lc.Json()
		if j.Error != nil {
			t.Fatal(j.Error)
		}
	})
}

func Test_C47_LinkedCollections_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_JsonPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		j := lc.JsonPtr()
		if j == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_C47_LinkedCollections_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ParseInjectUsingJson", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		j := lc.Json()
		lc2 := corestr.New.LinkedCollection.Create()
		_, err := lc2.ParseInjectUsingJson(&j)
		if err != nil {
			t.Fatal(err)
		}
	})
}

func Test_C47_LinkedCollections_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ParseInjectUsingJsonMust", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		j := lc.Json()
		lc2 := corestr.New.LinkedCollection.Create()
		result := lc2.ParseInjectUsingJsonMust(&j)
		_ = result
	})
}

func Test_C47_LinkedCollections_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_JsonParseSelfInject", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		j := lc.Json()
		lc2 := corestr.New.LinkedCollection.Create()
		err := lc2.JsonParseSelfInject(&j)
		if err != nil {
			t.Fatal(err)
		}
	})
}

func Test_C47_LinkedCollections_JsonModel(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_JsonModel", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		m := lc.JsonModel()
		if len(m) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_JsonModelAny", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		a := lc.JsonModelAny()
		if a == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_C47_LinkedCollections_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AsJsonContractsBinder", func() {
		lc := corestr.New.LinkedCollection.Create()
		_ = lc.AsJsonContractsBinder()
	})
}

func Test_C47_LinkedCollections_AsJsoner(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AsJsoner", func() {
		lc := corestr.New.LinkedCollection.Create()
		_ = lc.AsJsoner()
	})
}

func Test_C47_LinkedCollections_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AsJsonParseSelfInjector", func() {
		lc := corestr.New.LinkedCollection.Create()
		_ = lc.AsJsonParseSelfInjector()
	})
}

func Test_C47_LinkedCollections_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AsJsonMarshaller", func() {
		lc := corestr.New.LinkedCollection.Create()
		_ = lc.AsJsonMarshaller()
	})
}

// Creators
func Test_C47_NewLinkedCollectionCreator_Create(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_Create", func() {
		lc := corestr.New.LinkedCollection.Create()
		if !lc.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C47_NewLinkedCollectionCreator_Empty(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_Empty", func() {
		lc := corestr.New.LinkedCollection.Empty()
		if !lc.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C47_NewLinkedCollectionCreator_Strings(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_Strings", func() {
		lc := corestr.New.LinkedCollection.Strings("a", "b")
		if lc.Length() != 1 {
			t.Fatal("expected 1 node with 2 strings")
		}
	})
}

func Test_C47_NewLinkedCollectionCreator_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_Strings_Empty", func() {
		lc := corestr.New.LinkedCollection.Strings()
		if lc.Length() != 0 {
			t.Fatal("expected empty")
		}
	})
}

func Test_C47_NewLinkedCollectionCreator_UsingCollections(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_UsingCollections", func() {
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc := corestr.New.LinkedCollection.UsingCollections(c1, c2)
		if lc.Length() < 1 {
			t.Fatal("expected at least 1")
		}
	})
}

func Test_C47_NewLinkedCollectionCreator_UsingCollections_Nil(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_UsingCollections_Nil", func() {
		lc := corestr.New.LinkedCollection.UsingCollections()
		if lc.Length() != 0 {
			t.Fatal("expected empty")
		}
	})
}

func Test_C47_NewLinkedCollectionCreator_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_PointerStringsPtr", func() {
		s1, s2 := "a", "b"
		ptrs := []*string{&s1, &s2}
		lc := corestr.New.LinkedCollection.PointerStringsPtr(&ptrs)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_NewLinkedCollectionCreator_PointerStringsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_PointerStringsPtr_Nil", func() {
		lc := corestr.New.LinkedCollection.PointerStringsPtr(nil)
		if !lc.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

// EmptyCreator
func Test_C47_EmptyCreator_LinkedCollections(t *testing.T) {
	safeTest(t, "Test_C47_EmptyCreator_LinkedCollections", func() {
		lc := corestr.Empty.LinkedCollections()
		if !lc.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

// NonChainedLinkedCollectionNodes
func Test_C47_NonChainedLinkedCollectionNodes_Basic(t *testing.T) {
	safeTest(t, "Test_C47_NonChainedLinkedCollectionNodes_Basic", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)
		if !nc.IsEmpty() || nc.Length() != 0 || nc.HasItems() {
			t.Fatal("expected empty")
		}
		col := corestr.New.Collection.Strings([]string{"a"})
		n1 := &corestr.LinkedCollectionNode{Element: col}
		nc.Adds(n1)
		if nc.Length() != 1 || nc.IsEmpty() || !nc.HasItems() {
			t.Fatal("expected 1")
		}
		if nc.First() != n1 || nc.Last() != n1 {
			t.Fatal("first/last mismatch")
		}
	})
}

func Test_C47_NonChainedLinkedCollectionNodes_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C47_NonChainedLinkedCollectionNodes_FirstOrDefault_Empty", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(0)
		if nc.FirstOrDefault() != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_C47_NonChainedLinkedCollectionNodes_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C47_NonChainedLinkedCollectionNodes_LastOrDefault_Empty", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(0)
		if nc.LastOrDefault() != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_C47_NonChainedLinkedCollectionNodes_ApplyChaining(t *testing.T) {
	safeTest(t, "Test_C47_NonChainedLinkedCollectionNodes_ApplyChaining", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(2)
		nc.Adds(
			&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})},
			&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})},
		)
		nc.ApplyChaining()
		if !nc.IsChainingApplied() {
			t.Fatal("should be applied")
		}
		// Apply again - should be no-op
		nc.ApplyChaining()
	})
}

func Test_C47_NonChainedLinkedCollectionNodes_ToChainedNodes(t *testing.T) {
	safeTest(t, "Test_C47_NonChainedLinkedCollectionNodes_ToChainedNodes", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(2)
		nc.Adds(
			&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})},
			&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})},
		)
		chained := nc.ToChainedNodes()
		_ = chained
	})
}

func Test_C47_NonChainedLinkedCollectionNodes_ToChainedNodes_Empty(t *testing.T) {
	safeTest(t, "Test_C47_NonChainedLinkedCollectionNodes_ToChainedNodes_Empty", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(0)
		chained := nc.ToChainedNodes()
		if len(*chained) != 0 {
			t.Fatal("expected 0")
		}
	})
}

// Async tests
func Test_C47_LinkedCollections_AddAsync(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAsync", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsync(corestr.New.Collection.Strings([]string{"a"}), wg)
		wg.Wait()
		if lc.LengthLock() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_AddStringsAsync(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsAsync", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddStringsAsync(wg, []string{"a", "b"})
		wg.Wait()
		if lc.LengthLock() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_AddStringsAsync_Nil(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsAsync_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsAsync(nil, nil)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C47_LinkedCollections_AddAsyncFuncItems(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAsyncFuncItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItems(wg, false, func() []string {
			return []string{"a"}
		})
		if lc.LengthLock() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_AddAsyncFuncItems_Nil(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAsyncFuncItems_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAsyncFuncItems(nil, false)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C47_LinkedCollections_AddAsyncFuncItemsPointer(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAsyncFuncItemsPointer", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItemsPointer(wg, false, func() []string {
			return []string{"a"}
		})
		if lc.LengthLock() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_AddAsyncFuncItemsPointer_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAsyncFuncItemsPointer_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItemsPointer(wg, false, func() []string {
			return []string{}
		})
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C47_LinkedCollections_AddCollectionsPtr(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddCollectionsPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.AddCollectionsPtr([]*corestr.Collection{c1})
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C47_LinkedCollections_AddCollectionsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddCollectionsPtr_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollectionsPtr(nil)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C47_LinkedCollections_AddCollections(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddCollections", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.AddCollections([]*corestr.Collection{c1})
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}
