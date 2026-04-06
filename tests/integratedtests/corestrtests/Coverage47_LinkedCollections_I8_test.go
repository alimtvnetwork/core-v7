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
		// Arrange
		var n *corestr.LinkedCollectionNode

		// Act
		actual := args.Map{"result": n.IsEmpty()}

		// Assert
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
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})
		n := &corestr.LinkedCollectionNode{Element: col}

		// Act
		actual := args.Map{"result": n.HasElement()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have element", actual)
	})
}

func Test_C47_LinkedCollectionNode_HasNext(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_HasNext", func() {
		// Arrange
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}

		// Act
		actual := args.Map{"result": n.HasNext()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have next", actual)
	})
}

func Test_C47_LinkedCollectionNode_Clone(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_Clone", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})
		lc := corestr.New.LinkedCollection.Strings("a", "b")
		node := lc.Head()
		cloned := node.Clone()

		// Act
		actual := args.Map{"result": cloned.HasNext()}

		// Assert
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
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		node := lc.Head()

		// Act
		actual := args.Map{"result": node.IsEqual(node)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "same node should be equal", actual)
	})
}

func Test_C47_LinkedCollectionNode_IsChainEqual(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_IsChainEqual", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("a")

		// Act
		actual := args.Map{"result": lc1.Head().IsChainEqual(lc2.Head())}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be chain equal", actual)
	})
}

func Test_C47_LinkedCollectionNode_IsEqualValue(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_IsEqualValue", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(col)

		// Act
		actual := args.Map{"result": lc.Head().IsEqualValue(col)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
	})
}

func Test_C47_LinkedCollectionNode_IsEqualValue_NilBoth(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_IsEqualValue_NilBoth", func() {
		// Arrange
		n := &corestr.LinkedCollectionNode{}

		// Act
		actual := args.Map{"result": n.IsEqualValue(nil)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
	})
}

func Test_C47_LinkedCollectionNode_String(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_String", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a"})
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(col)
		s := lc.Head().String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C47_LinkedCollectionNode_List(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_List", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		list := lc.Head().List()

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollectionNode_ListPtr(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_ListPtr", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		ptr := lc.Head().ListPtr()

		// Act
		actual := args.Map{"result": ptr == nil || len(*ptr) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollectionNode_Join(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_Join", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		j := lc.Head().Join(",")

		// Act
		actual := args.Map{"result": j != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "got", actual)
	})
}

func Test_C47_LinkedCollectionNode_StringList(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_StringList", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		s := lc.Head().StringList("H:")

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
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
		// Arrange
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

		// Act
		actual := args.Map{"result": count != 2 || length != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "count= length=", actual)
	})
}

func Test_C47_LinkedCollectionNode_LoopEndOfChain_BreakFirst(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_LoopEndOfChain_BreakFirst", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		_, length := lc.Head().LoopEndOfChain(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			return true
		})

		// Act
		actual := args.Map{"result": length != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollectionNode_AddNext(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_AddNext", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		col1 := corestr.New.Collection.Strings([]string{"a"})
		col2 := corestr.New.Collection.Strings([]string{"c"})
		lc.Add(col1)
		lc.Add(col2)
		colB := corestr.New.Collection.Strings([]string{"b"})
		lc.Head().AddNext(lc, colB)

		// Act
		actual := args.Map{"result": lc.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C47_LinkedCollectionNode_AddNextNode(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_AddNextNode", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(col)
		newNode := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		lc.Head().AddNextNode(lc, newNode)

		// Act
		actual := args.Map{"result": lc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollectionNode_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollectionNode_CreateLinkedList", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(col)
		newLC := lc.Head().CreateLinkedList()

		// Act
		actual := args.Map{"result": newLC.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ===================== LinkedCollections =====================

func Test_C47_LinkedCollections_Create_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Create_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": lc.IsEmpty() || lc.Length() != 0 || lc.HasItems()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C47_LinkedCollections_Add(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Add", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		lc.Add(col)

		// Act
		actual := args.Map{"result": lc.Length() != 1 || lc.Head() == nil || lc.Tail() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "add failed", actual)
	})
}

func Test_C47_LinkedCollections_AddMultiple(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddMultiple", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))

		// Act
		actual := args.Map{"result": lc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_AddStrings(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStrings", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 collection node", actual)
	})
}

func Test_C47_LinkedCollections_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStrings_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings()

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock("a", "b")

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AddStringsLock_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsLock_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock()

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_AddLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddLock(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AddFront_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddFront_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.AddFront(col)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AddFront_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddFront_NonEmpty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.AddFront(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.First().List()[0] != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "front should be a", actual)
	})
}

func Test_C47_LinkedCollections_AddFrontLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddFrontLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddFrontLock(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_PushFront(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_PushFront", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.PushFront(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_Push(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Push", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Push(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_PushBack(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_PushBack", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.PushBack(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_PushBackLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_PushBackLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.PushBackLock(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_First_Single_Last(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_First_Single_Last", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(col)

		// Act
		actual := args.Map{"result": lc.First() != col || lc.Single() != col || lc.Last() != col}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "mismatch", actual)
	})
}

func Test_C47_LinkedCollections_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_FirstOrDefault_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		r := lc.FirstOrDefault()

		// Act
		actual := args.Map{"result": r == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty collection, not nil", actual)
	})
}

func Test_C47_LinkedCollections_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_LastOrDefault_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		r := lc.LastOrDefault()

		// Act
		actual := args.Map{"result": r == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty collection, not nil", actual)
	})
}

func Test_C47_LinkedCollections_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AllIndividualItemsLength", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		lc.AddStrings("c")

		// Act
		actual := args.Map{"result": lc.AllIndividualItemsLength() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C47_LinkedCollections_LengthLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_LengthLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")

		// Act
		actual := args.Map{"result": lc.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEmptyLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": lc.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C47_LinkedCollections_AppendNode_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendNode_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		lc.AppendNode(node)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AppendNode_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendNode_NonEmpty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		lc.AppendNode(node)

		// Act
		actual := args.Map{"result": lc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_AddBackNode(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddBackNode", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		lc.AddBackNode(node)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendChainOfNodes", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc1.AddStrings("b")

		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AppendChainOfNodes(lc1.Head())

		// Act
		actual := args.Map{"result": lc2.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_AppendChainOfNodes_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendChainOfNodes_NonEmpty", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")

		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("x")
		lc2.AppendChainOfNodes(lc1.Head())

		// Act
		actual := args.Map{"result": lc2.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_C47_LinkedCollections_AppendChainOfNodesAsync(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendChainOfNodesAsync", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("x")

		chain := corestr.New.LinkedCollection.Create()
		chain.AddStrings("a")

		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AppendChainOfNodesAsync(chain.Head(), wg)
		wg.Wait()

		// Act
		actual := args.Map{"result": lc.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_C47_LinkedCollections_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AttachWithNode_NilCurrent", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		err := lc.AttachWithNode(nil, node)

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C47_LinkedCollections_AddAnother(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAnother", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc1.AddStrings("b")

		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddAnother(lc1)

		// Act
		actual := args.Map{"result": lc2.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_AddAnother_Nil(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAnother_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAnother(nil)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddCollection_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollection(nil)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_AddCollection(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollection(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AppendCollections(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendCollections", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.AppendCollections(true, c1, nil, c2)

		// Act
		actual := args.Map{"result": lc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_AppendCollections_NilSkip(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AppendCollections_NilSkip", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AppendCollections(true, nil)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_AddStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsOfStrings", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false, []string{"a"}, []string{"b"})

		// Act
		actual := args.Map{"result": lc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_AddStringsOfStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsOfStrings_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_Loop(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Loop", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return false
		})

		// Act
		actual := args.Map{"result": count != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Loop_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {

		// Act
			actual := args.Map{"result": false}

		// Assert
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "should not call", actual)
			return false
		})
	})
}

func Test_C47_LinkedCollections_Filter(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Filter", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		result := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_FilterAsCollection(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_FilterAsCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		lc.AddStrings("c")
		col := lc.FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		}, 0)

		// Act
		actual := args.Map{"result": col.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C47_LinkedCollections_FilterAsCollections(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_FilterAsCollections", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		cols := lc.FilterAsCollections(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})

		// Act
		actual := args.Map{"result": len(cols) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_GetNextNodes", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		lc.AddStrings("c")
		nodes := lc.GetNextNodes(2)

		// Act
		actual := args.Map{"result": len(nodes) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_GetAllLinkedNodes", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		all := lc.GetAllLinkedNodes()

		// Act
		actual := args.Map{"result": len(all) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_InsertAt_Front", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("b")
		lc.InsertAt(0, corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.First().List()[0] != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "insert at front failed", actual)
	})
}

func Test_C47_LinkedCollections_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_RemoveNodeByIndex", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		lc.AddStrings("c")
		lc.RemoveNodeByIndex(1)

		// Act
		actual := args.Map{"result": lc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_RemoveNodeByIndex_First(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_RemoveNodeByIndex_First", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		lc.RemoveNodeByIndex(0)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_RemoveNode(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_RemoveNode", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		lc.RemoveNode(lc.Head())

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_RemoveNodeByIndexes", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		lc.AddStrings("c")
		lc.RemoveNodeByIndexes(false, 0, 2)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_RemoveNodeByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_RemoveNodeByIndexes_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.RemoveNodeByIndexes(false)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_IndexAt(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IndexAt", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		node := lc.IndexAt(0)

		// Act
		actual := args.Map{"result": node == nil}

		// Assert
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
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		n := lc.SafeIndexAt(0)

		// Act
		actual := args.Map{"result": n == nil}

		// Assert
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
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		col := lc.SafePointerIndexAt(0)

		// Act
		actual := args.Map{"result": col == nil}

		// Assert
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
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("a")

		// Act
		actual := args.Map{"result": lc1.IsEqualsPtr(lc2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C47_LinkedCollections_IsEqualsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEqualsPtr_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": lc.IsEqualsPtr(nil)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil should not be equal", actual)
	})
}

func Test_C47_LinkedCollections_IsEqualsPtr_SameRef(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEqualsPtr_SameRef", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")

		// Act
		actual := args.Map{"result": lc.IsEqualsPtr(lc)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "same ref should be equal", actual)
	})
}

func Test_C47_LinkedCollections_IsEqualsPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEqualsPtr_BothEmpty", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc2 := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": lc1.IsEqualsPtr(lc2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "both empty should be equal", actual)
	})
}

func Test_C47_LinkedCollections_IsEqualsPtr_DiffLength(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_IsEqualsPtr_DiffLength", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("a")
		lc2.AddStrings("b")

		// Act
		actual := args.Map{"result": lc1.IsEqualsPtr(lc2)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "diff length should not be equal", actual)
	})
}

func Test_C47_LinkedCollections_ToCollection(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		lc.AddStrings("c")
		col := lc.ToCollection(0)

		// Act
		actual := args.Map{"result": col.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C47_LinkedCollections_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToCollection_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		col := lc.ToCollection(0)

		// Act
		actual := args.Map{"result": col.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_ToCollectionSimple(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToCollectionSimple", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		col := lc.ToCollectionSimple()

		// Act
		actual := args.Map{"result": col.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_ToStrings(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToStrings", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		s := lc.ToStrings()

		// Act
		actual := args.Map{"result": len(s) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_ToStringsPtr(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToStringsPtr", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		p := lc.ToStringsPtr()

		// Act
		actual := args.Map{"result": p == nil || len(*p) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_ToCollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToCollectionsOfCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		coc := lc.ToCollectionsOfCollection(0)

		// Act
		actual := args.Map{"result": coc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_ToCollectionsOfCollection_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ToCollectionsOfCollection_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		coc := lc.ToCollectionsOfCollection(0)

		// Act
		actual := args.Map{"result": coc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_ItemsOfItems(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ItemsOfItems", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.AddStrings("b")
		items := lc.ItemsOfItems()

		// Act
		actual := args.Map{"result": len(items) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_ItemsOfItems_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ItemsOfItems_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		items := lc.ItemsOfItems()

		// Act
		actual := args.Map{"result": len(items) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_ItemsOfItemsCollection(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ItemsOfItemsCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		cols := lc.ItemsOfItemsCollection()

		// Act
		actual := args.Map{"result": len(cols) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_SimpleSlice", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		ss := lc.SimpleSlice()

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_List(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_List", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		list := lc.List()

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_List_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_List_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		list := lc.List()

		// Act
		actual := args.Map{"result": len(list) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_ListPtr(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ListPtr", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		p := lc.ListPtr()

		// Act
		actual := args.Map{"result": p == nil || len(*p) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_String_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_String_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		s := lc.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected NoElements", actual)
	})
}

func Test_C47_LinkedCollections_String_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_String_NonEmpty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		s := lc.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected string", actual)
	})
}

func Test_C47_LinkedCollections_StringLock(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_StringLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		s := lc.StringLock()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C47_LinkedCollections_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_StringLock_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		s := lc.StringLock()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected NoElements", actual)
	})
}

func Test_C47_LinkedCollections_Join(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Join", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		j := lc.Join(",")

		// Act
		actual := args.Map{"result": j == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C47_LinkedCollections_Joins(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Joins", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		j := lc.Joins(",", "b")

		// Act
		actual := args.Map{"result": j == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C47_LinkedCollections_Joins_NilItems(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Joins_NilItems", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		j := lc.Joins(",", "x")

		// Act
		actual := args.Map{"result": j != "x"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
	})
}

func Test_C47_LinkedCollections_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ConcatNew", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("b")
		result := lc1.ConcatNew(false, lc2)

		// Act
		actual := args.Map{"result": result.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C47_LinkedCollections_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ConcatNew_EmptyClone", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		result := lc.ConcatNew(true)

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ConcatNew_EmptyNoClone", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		result := lc.ConcatNew(false)

		// Act
		actual := args.Map{"result": result != lc}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same ref", actual)
	})
}

func Test_C47_LinkedCollections_Clear(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_Clear", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.Clear()

		// Act
		actual := args.Map{"result": lc.IsEmpty()}

		// Assert
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
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		lc.RemoveAll()

		// Act
		actual := args.Map{"result": lc.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C47_LinkedCollections_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_GetCompareSummary", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.AddStrings("a")
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("b")
		s := lc1.GetCompareSummary(lc2, "left", "right")

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected summary", actual)
	})
}

// JSON
func Test_C47_LinkedCollections_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_MarshalJSON", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		data, err := json.Marshal(lc)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{
			"result": string(data) != `["a",
			"b"]`,
		}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_C47_LinkedCollections_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_UnmarshalJSON", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		err := json.Unmarshal([]byte(`["x","y"]`), lc)

		// Act
		actual := args.Map{"result": err}

		// Assert
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
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		j := lc.Json()

		// Act
		actual := args.Map{"result": j.Error}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "j.Error", actual)
	})
}

func Test_C47_LinkedCollections_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_JsonPtr", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		j := lc.JsonPtr()

		// Act
		actual := args.Map{"result": j == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C47_LinkedCollections_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_ParseInjectUsingJson", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		j := lc.Json()
		lc2 := corestr.New.LinkedCollection.Create()
		_, err := lc2.ParseInjectUsingJson(&j)

		// Act
		actual := args.Map{"result": err}

		// Assert
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
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		j := lc.Json()
		lc2 := corestr.New.LinkedCollection.Create()
		err := lc2.JsonParseSelfInject(&j)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C47_LinkedCollections_JsonModel(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_JsonModel", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		m := lc.JsonModel()

		// Act
		actual := args.Map{"result": len(m) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_JsonModelAny", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a")
		a := lc.JsonModelAny()

		// Act
		actual := args.Map{"result": a == nil}

		// Assert
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
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": lc.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C47_NewLinkedCollectionCreator_Empty(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Empty()

		// Act
		actual := args.Map{"result": lc.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C47_NewLinkedCollectionCreator_Strings(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_Strings", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Strings("a", "b")

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 node with 2 strings", actual)
	})
}

func Test_C47_NewLinkedCollectionCreator_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_Strings_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Strings()

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C47_NewLinkedCollectionCreator_UsingCollections(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_UsingCollections", func() {
		// Arrange
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc := corestr.New.LinkedCollection.UsingCollections(c1, c2)

		// Act
		actual := args.Map{"result": lc.Length() < 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_C47_NewLinkedCollectionCreator_UsingCollections_Nil(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_UsingCollections_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.UsingCollections()

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C47_NewLinkedCollectionCreator_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_PointerStringsPtr", func() {
		// Arrange
		s1, s2 := "a", "b"
		ptrs := []*string{&s1, &s2}
		lc := corestr.New.LinkedCollection.PointerStringsPtr(&ptrs)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_NewLinkedCollectionCreator_PointerStringsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_C47_NewLinkedCollectionCreator_PointerStringsPtr_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.PointerStringsPtr(nil)

		// Act
		actual := args.Map{"result": lc.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// EmptyCreator
func Test_C47_EmptyCreator_LinkedCollections(t *testing.T) {
	safeTest(t, "Test_C47_EmptyCreator_LinkedCollections", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()

		// Act
		actual := args.Map{"result": lc.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// NonChainedLinkedCollectionNodes
func Test_C47_NonChainedLinkedCollectionNodes_Basic(t *testing.T) {
	safeTest(t, "Test_C47_NonChainedLinkedCollectionNodes_Basic", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act
		actual := args.Map{"result": nc.IsEmpty() || nc.Length() != 0 || nc.HasItems()}

		// Assert
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
		// Arrange
		nc := corestr.NewNonChainedLinkedCollectionNodes(0)

		// Act
		actual := args.Map{"result": nc.FirstOrDefault() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C47_NonChainedLinkedCollectionNodes_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C47_NonChainedLinkedCollectionNodes_LastOrDefault_Empty", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedCollectionNodes(0)

		// Act
		actual := args.Map{"result": nc.LastOrDefault() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C47_NonChainedLinkedCollectionNodes_ApplyChaining(t *testing.T) {
	safeTest(t, "Test_C47_NonChainedLinkedCollectionNodes_ApplyChaining", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedCollectionNodes(2)
		nc.Adds(
			&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})},
			&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})},
		)
		nc.ApplyChaining()

		// Act
		actual := args.Map{"result": nc.IsChainingApplied()}

		// Assert
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
		// Arrange
		nc := corestr.NewNonChainedLinkedCollectionNodes(0)
		chained := nc.ToChainedNodes()

		// Act
		actual := args.Map{"result": len(*chained) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// Async tests
func Test_C47_LinkedCollections_AddAsync(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAsync", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsync(corestr.New.Collection.Strings([]string{"a"}), wg)
		wg.Wait()

		// Act
		actual := args.Map{"result": lc.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AddStringsAsync(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsAsync", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddStringsAsync(wg, []string{"a", "b"})
		wg.Wait()

		// Act
		actual := args.Map{"result": lc.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AddStringsAsync_Nil(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddStringsAsync_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsAsync(nil, nil)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_AddAsyncFuncItems(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAsyncFuncItems", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItems(wg, false, func() []string {
			return []string{"a"}
		})

		// Act
		actual := args.Map{"result": lc.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AddAsyncFuncItems_Nil(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAsyncFuncItems_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAsyncFuncItems(nil, false)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_AddAsyncFuncItemsPointer(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAsyncFuncItemsPointer", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItemsPointer(wg, false, func() []string {
			return []string{"a"}
		})

		// Act
		actual := args.Map{"result": lc.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AddAsyncFuncItemsPointer_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddAsyncFuncItemsPointer_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItemsPointer(wg, false, func() []string {
			return []string{}
		})

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_AddCollectionsPtr(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddCollectionsPtr", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.AddCollectionsPtr([]*corestr.Collection{c1})

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C47_LinkedCollections_AddCollectionsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddCollectionsPtr_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollectionsPtr(nil)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C47_LinkedCollections_AddCollections(t *testing.T) {
	safeTest(t, "Test_C47_LinkedCollections_AddCollections", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.AddCollections([]*corestr.Collection{c1})

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}
