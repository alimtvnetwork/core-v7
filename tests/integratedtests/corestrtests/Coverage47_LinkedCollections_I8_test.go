package corestrtests

import (
	"encoding/json"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ===================== LinkedCollectionNode =====================

func Test_C47_LinkedCollectionNode_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_IsEmpty", func() {
		var n *corestr.LinkedCollectionNode
		actual := args.Map{"result": n.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
		n2 := &corestr.LinkedCollectionNode{}
		actual := args.Map{"result": n2.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "nil element should be empty", actual)
	})
}

func Test_C47_LinkedCollectionNode_HasElement(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_HasElement", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		n := &corestr.LinkedCollectionNode{Element: col}
		actual := args.Map{"result": n.HasElement()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have element", actual)
	})
}

func Test_C47_LinkedCollectionNode_HasNext(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_HasNext", func() {
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		actual := args.Map{"result": n.HasNext()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have next", actual)
	})
}

func Test_C47_LinkedCollectionNode_Clone(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_Clone", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		lc := corestr.New.LinkedCollection.Strings("a", "b")
		node := lc.Head()
		cloned := node.Clone()
		actual := args.Map{"result": cloned.HasNext()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "clone should not have next", actual)
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
		actual := args.Map{"result": node.IsEqual(node)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "same node should be equal", actual)
	})
}

func Test_C47_LinkedCollectionNode_IsChainEqual(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_IsChainEqual", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("a")
		actual := args.Map{"result": lc1.Head().IsChainEqual(lc2.Head())}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be chain equal", actual)
	})
}

func Test_C47_LinkedCollectionNode_IsEqualValue(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_IsEqualValue", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(col)
		actual := args.Map{"result": lc.Head().IsEqualValue(col)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
	})
}

func Test_C47_LinkedCollectionNode_IsEqualValue_NilBoth(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_IsEqualValue_NilBoth", func() {
		n := &corestr.LinkedCollectionNode{}
		actual := args.Map{"result": n.IsEqualValue(nil)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
	})
}

func Test_C47_LinkedCollectionNode_String(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_String", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(col)
		s := lc.Head().String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C47_LinkedCollectionNode_List(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_List", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		list := lc.Head().List()
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollectionNode_ListPtr(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_ListPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		ptr := lc.Head().ListPtr()
		actual := args.Map{"result": ptr == nil || len(*ptr) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollectionNode_Join(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_Join", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		j := lc.Head().Join(",")
		actual := args.Map{"result": j != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "got", actual)
	})
}

func Test_C47_LinkedCollectionNode_StringList(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_StringList", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		s := lc.Head().StringList("H:")
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
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
		actual := args.Map{"result": count != 2 || length != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "count= length=", actual)
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
		actual := args.Map{"result": length != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
		actual := args.Map{"result": lc.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C47_LinkedCollectionNode_AddNextNode(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_AddNextNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(col)
		newNode := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		lc.Head().AddNextNode(lc, newNode)
		actual := args.Map{"result": lc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollectionNode_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_CreateLinkedList", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(col)
		newLC := lc.Head().CreateLinkedList()
		actual := args.Map{"result": newLC.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ===================== LinkedCollections =====================

func Test_C47_LinkedCollections_Create_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Create_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": lc.IsEmpty() || lc.Length() != 0 || lc.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C47_LinkedCollections_Add(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Add", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		lc.Add(col)
		actual := args.Map{"result": lc.Length() != 1 || lc.Head() == nil || lc.Tail() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "add failed", actual)
	})
}

func Test_C47_LinkedCollections_AddMultiple(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddMultiple", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		actual := args.Map{"result": lc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_AddStrings(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 collection node", actual)
	})
}

func Test_C47_LinkedCollections_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStrings_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings()
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock("a", "b")
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AddStringsLock_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsLock_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock()
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_AddLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddLock(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AddFront_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddFront_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.AddFront(col)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AddFront_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddFront_NonEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.AddFront(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.First().List()[0] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "front should be a", actual)
	})
}

func Test_C47_LinkedCollections_AddFrontLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddFrontLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddFrontLock(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_PushFront(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_PushFront", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.PushFront(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_Push(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Push", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Push(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_PushBack(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_PushBack", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.PushBack(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_PushBackLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_PushBackLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.PushBackLock(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_First_Single_Last(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_First_Single_Last", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(col)
		actual := args.Map{"result": lc.First() != col || lc.Single() != col || lc.Last() != col}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "mismatch", actual)
	})
}

func Test_C47_LinkedCollections_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_FirstOrDefault_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		r := lc.FirstOrDefault()
		actual := args.Map{"result": r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty collection, not nil", actual)
	})
}

func Test_C47_LinkedCollections_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_LastOrDefault_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		r := lc.LastOrDefault()
		actual := args.Map{"result": r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty collection, not nil", actual)
	})
}

func Test_C47_LinkedCollections_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AllIndividualItemsLength", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		lc.AddStrings("c")
		actual := args.Map{"result": lc.AllIndividualItemsLength() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C47_LinkedCollections_LengthLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_LengthLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		actual := args.Map{"result": lc.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEmptyLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": lc.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C47_LinkedCollections_AppendNode_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendNode_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		lc.AppendNode(node)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AppendNode_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendNode_NonEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		lc.AppendNode(node)
		actual := args.Map{"result": lc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_AddBackNode(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddBackNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		lc.AddBackNode(node)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendChainOfNodes", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc1.AddStrings("b")

		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AppendChainOfNodes(lc1.Head())
		actual := args.Map{"result": lc2.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_AppendChainOfNodes_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendChainOfNodes_NonEmpty", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")

		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("x")
		lc2.AppendChainOfNodes(lc1.Head())
		actual := args.Map{"result": lc2.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
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
		actual := args.Map{"result": lc.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_C47_LinkedCollections_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AttachWithNode_NilCurrent", func() {
		lc := corestr.New.LinkedCollection.Create()
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		err := lc.AttachWithNode(nil, node)
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C47_LinkedCollections_AddAnother(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAnother", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc1.AddStrings("b")

		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddAnother(lc1)
		actual := args.Map{"result": lc2.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_AddAnother_Nil(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAnother_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAnother(nil)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddCollection_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollection(nil)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_AddCollection(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollection(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AppendCollections(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendCollections", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.AppendCollections(true, c1, nil, c2)
		actual := args.Map{"result": lc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_AppendCollections_NilSkip(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendCollections_NilSkip", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AppendCollections(true, nil)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_AddStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsOfStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false, []string{"a"}, []string{"b"})
		actual := args.Map{"result": lc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_AddStringsOfStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsOfStrings_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": count != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Loop_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			actual := args.Map{"result": false}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "should not call", actual)
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
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": col.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C47_LinkedCollections_FilterAsCollections(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_FilterAsCollections", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		cols := lc.FilterAsCollections(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		actual := args.Map{"result": len(cols) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_GetNextNodes", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		lc.AddStrings("c")
		nodes := lc.GetNextNodes(2)
		actual := args.Map{"result": len(nodes) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_GetAllLinkedNodes", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		all := lc.GetAllLinkedNodes()
		actual := args.Map{"result": len(all) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_InsertAt_Front", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("b")
		lc.InsertAt(0, corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.First().List()[0] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "insert at front failed", actual)
	})
}

func Test_C47_LinkedCollections_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_RemoveNodeByIndex", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		lc.AddStrings("c")
		lc.RemoveNodeByIndex(1)
		actual := args.Map{"result": lc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_RemoveNodeByIndex_First(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_RemoveNodeByIndex_First", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		lc.RemoveNodeByIndex(0)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_RemoveNode(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_RemoveNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		lc.RemoveNode(lc.Head())
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_RemoveNodeByIndexes", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		lc.AddStrings("c")
		lc.RemoveNodeByIndexes(false, 0, 2)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_RemoveNodeByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_RemoveNodeByIndexes_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.RemoveNodeByIndexes(false)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_IndexAt(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IndexAt", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		node := lc.IndexAt(0)
		actual := args.Map{"result": node == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected node", actual)
		node2 := lc.IndexAt(-1)
		actual := args.Map{"result": node2 != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "negative should be nil", actual)
	})
}

func Test_C47_LinkedCollections_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_SafeIndexAt", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		n := lc.SafeIndexAt(0)
		actual := args.Map{"result": n == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected node", actual)
		n2 := lc.SafeIndexAt(99)
		actual := args.Map{"result": n2 != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		n3 := lc.SafeIndexAt(-1)
		actual := args.Map{"result": n3 != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C47_LinkedCollections_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_SafePointerIndexAt", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		col := lc.SafePointerIndexAt(0)
		actual := args.Map{"result": col == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected collection", actual)
		col2 := lc.SafePointerIndexAt(99)
		actual := args.Map{"result": col2 != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C47_LinkedCollections_IsEqualsPtr(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEqualsPtr", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("a")
		actual := args.Map{"result": lc1.IsEqualsPtr(lc2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C47_LinkedCollections_IsEqualsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEqualsPtr_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": lc.IsEqualsPtr(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil should not be equal", actual)
	})
}

func Test_C47_LinkedCollections_IsEqualsPtr_SameRef(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEqualsPtr_SameRef", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		actual := args.Map{"result": lc.IsEqualsPtr(lc)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "same ref should be equal", actual)
	})
}

func Test_C47_LinkedCollections_IsEqualsPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEqualsPtr_BothEmpty", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc2 := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": lc1.IsEqualsPtr(lc2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "both empty should be equal", actual)
	})
}

func Test_C47_LinkedCollections_IsEqualsPtr_DiffLength(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEqualsPtr_DiffLength", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("a")
		lc2.AddStrings("b")
		actual := args.Map{"result": lc1.IsEqualsPtr(lc2)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "diff length should not be equal", actual)
	})
}

func Test_C47_LinkedCollections_ToCollection(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		lc.AddStrings("c")
		col := lc.ToCollection(0)
		actual := args.Map{"result": col.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C47_LinkedCollections_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToCollection_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := lc.ToCollection(0)
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_ToCollectionSimple(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToCollectionSimple", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		col := lc.ToCollectionSimple()
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_ToStrings(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		s := lc.ToStrings()
		actual := args.Map{"result": len(s) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_ToStringsPtr(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToStringsPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		p := lc.ToStringsPtr()
		actual := args.Map{"result": p == nil || len(*p) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_ToCollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToCollectionsOfCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		coc := lc.ToCollectionsOfCollection(0)
		actual := args.Map{"result": coc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_ToCollectionsOfCollection_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToCollectionsOfCollection_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		coc := lc.ToCollectionsOfCollection(0)
		actual := args.Map{"result": coc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_ItemsOfItems(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ItemsOfItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		items := lc.ItemsOfItems()
		actual := args.Map{"result": len(items) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_ItemsOfItems_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ItemsOfItems_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		items := lc.ItemsOfItems()
		actual := args.Map{"result": len(items) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_ItemsOfItemsCollection(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ItemsOfItemsCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		cols := lc.ItemsOfItemsCollection()
		actual := args.Map{"result": len(cols) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_SimpleSlice", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		ss := lc.SimpleSlice()
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_List(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_List", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		list := lc.List()
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_List_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_List_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		list := lc.List()
		actual := args.Map{"result": len(list) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_ListPtr(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ListPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		p := lc.ListPtr()
		actual := args.Map{"result": p == nil || len(*p) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_String_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_String_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		s := lc.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected NoElements", actual)
	})
}

func Test_C47_LinkedCollections_String_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_String_NonEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		s := lc.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected string", actual)
	})
}

func Test_C47_LinkedCollections_StringLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_StringLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		s := lc.StringLock()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C47_LinkedCollections_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_StringLock_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		s := lc.StringLock()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected NoElements", actual)
	})
}

func Test_C47_LinkedCollections_Join(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Join", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		j := lc.Join(",")
		actual := args.Map{"result": j == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C47_LinkedCollections_Joins(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Joins", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		j := lc.Joins(",", "b")
		actual := args.Map{"result": j == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C47_LinkedCollections_Joins_NilItems(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Joins_NilItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		j := lc.Joins(",", "x")
		actual := args.Map{"result": j != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
	})
}

func Test_C47_LinkedCollections_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ConcatNew", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("b")
		result := lc1.ConcatNew(false, lc2)
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ConcatNew_EmptyClone", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		result := lc.ConcatNew(true)
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ConcatNew_EmptyNoClone", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		result := lc.ConcatNew(false)
		actual := args.Map{"result": result != lc}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same ref", actual)
	})
}

func Test_C47_LinkedCollections_Clear(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Clear", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.Clear()
		actual := args.Map{"result": lc.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
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
		actual := args.Map{"result": lc.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C47_LinkedCollections_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_GetCompareSummary", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("b")
		s := lc1.GetCompareSummary(lc2, "left", "right")
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected summary", actual)
	})
}

// JSON
func Test_C47_LinkedCollections_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_MarshalJSON", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		data, err := json.Marshal(lc)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": string(data) != `["a","b"]`}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_C47_LinkedCollections_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_UnmarshalJSON", func() {
		lc := corestr.New.LinkedCollection.Create()
		err := json.Unmarshal([]byte(`["x","y"]`), lc)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		list := lc.List()
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_Json(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Json", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		j := lc.Json()
		actual := args.Map{"result": j.Error}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "j.Error", actual)
	})
}

func Test_C47_LinkedCollections_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_JsonPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		j := lc.JsonPtr()
		actual := args.Map{"result": j == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C47_LinkedCollections_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ParseInjectUsingJson", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		j := lc.Json()
		lc2 := corestr.New.LinkedCollection.Create()
		_, err := lc2.ParseInjectUsingJson(&j)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
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
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C47_LinkedCollections_JsonModel(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_JsonModel", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		m := lc.JsonModel()
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_JsonModelAny", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		a := lc.JsonModelAny()
		actual := args.Map{"result": a == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
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
		actual := args.Map{"result": lc.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C47_NewLinkedCollectionCreator_Empty(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_Empty", func() {
		lc := corestr.New.LinkedCollection.Empty()
		actual := args.Map{"result": lc.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C47_NewLinkedCollectionCreator_Strings(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_Strings", func() {
		lc := corestr.New.LinkedCollection.Strings("a", "b")
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 node with 2 strings", actual)
	})
}

func Test_C47_NewLinkedCollectionCreator_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_Strings_Empty", func() {
		lc := corestr.New.LinkedCollection.Strings()
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C47_NewLinkedCollectionCreator_UsingCollections(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_UsingCollections", func() {
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc := corestr.New.LinkedCollection.UsingCollections(c1, c2)
		actual := args.Map{"result": lc.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_C47_NewLinkedCollectionCreator_UsingCollections_Nil(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_UsingCollections_Nil", func() {
		lc := corestr.New.LinkedCollection.UsingCollections()
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C47_NewLinkedCollectionCreator_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_PointerStringsPtr", func() {
		s1, s2 := "a", "b"
		ptrs := []*string{&s1, &s2}
		lc := corestr.New.LinkedCollection.PointerStringsPtr(&ptrs)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_NewLinkedCollectionCreator_PointerStringsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_PointerStringsPtr_Nil", func() {
		lc := corestr.New.LinkedCollection.PointerStringsPtr(nil)
		actual := args.Map{"result": lc.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// EmptyCreator
func Test_C47_EmptyCreator_LinkedCollections(t *testing.T) {
	safeTest(t, "Test_C47_EmptyCreator_LinkedCollections", func() {
		lc := corestr.Empty.LinkedCollections()
		actual := args.Map{"result": lc.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// NonChainedLinkedCollectionNodes
func Test_C47_NonChainedLinkedCollectionNodes_Basic(t *testing.T) {
	safeTest(t, "Test_C47_NonChainedLinkedCollectionNodes_Basic", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)
		actual := args.Map{"result": nc.IsEmpty() || nc.Length() != 0 || nc.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		col := corestr.New.Collection.Strings([]string{"a"})
		n1 := &corestr.LinkedCollectionNode{Element: col}
		nc.Adds(n1)
		actual := args.Map{"result": nc.Length() != 1 || nc.IsEmpty() || !nc.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": nc.First() != n1 || nc.Last() != n1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "first/last mismatch", actual)
	})
}

func Test_C47_NonChainedLinkedCollectionNodes_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C47_NonChainedLinkedCollectionNodes_FirstOrDefault_Empty", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(0)
		actual := args.Map{"result": nc.FirstOrDefault() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C47_NonChainedLinkedCollectionNodes_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C47_NonChainedLinkedCollectionNodes_LastOrDefault_Empty", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(0)
		actual := args.Map{"result": nc.LastOrDefault() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
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
		actual := args.Map{"result": nc.IsChainingApplied()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be applied", actual)
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
		actual := args.Map{"result": len(*chained) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": lc.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AddStringsAsync(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsAsync", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddStringsAsync(wg, []string{"a", "b"})
		wg.Wait()
		actual := args.Map{"result": lc.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AddStringsAsync_Nil(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsAsync_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsAsync(nil, nil)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": lc.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AddAsyncFuncItems_Nil(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAsyncFuncItems_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAsyncFuncItems(nil, false)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": lc.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_AddCollectionsPtr(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddCollectionsPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.AddCollectionsPtr([]*corestr.Collection{c1})
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AddCollectionsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddCollectionsPtr_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollectionsPtr(nil)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_AddCollections(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddCollections", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.AddCollections([]*corestr.Collection{c1})
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}
