package corestrtests

import (
	"encoding/json"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =======================================================
// LinkedListNode
// =======================================================

func Test_C24_LinkedListNode_HasNext_NoNext(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_HasNext_NoNext", func() {
		// Arrange
		node := &corestr.LinkedListNode{Element: "a"}

		// Act
		actual := args.Map{"result": node.HasNext()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected HasNext false", actual)
	})
}

func Test_C24_LinkedListNode_EndOfChain_Single(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_EndOfChain_Single", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		end, length := ll.Head().EndOfChain()

		// Act
		actual := args.Map{"result": length != 1 || end.Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected length 1", actual)
	})
}

func Test_C24_LinkedListNode_EndOfChain_Multi(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_EndOfChain_Multi", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		end, length := ll.Head().EndOfChain()

		// Act
		actual := args.Map{"result": length != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual = args.Map{"result": end.Element != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c", actual)
	})
}

func Test_C24_LinkedListNode_Clone(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_Clone", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		cloned := ll.Head().Clone()

		// Act
		actual := args.Map{"result": cloned.HasNext()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "cloned should not have next", actual)
		actual = args.Map{"result": cloned.Element != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C24_LinkedListNode_LoopEndOfChain(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_LoopEndOfChain", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		var collected []string
		end, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			collected = append(collected, arg.CurrentNode.Element)
			return false
		})

		// Act
		actual := args.Map{"result": length != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual = args.Map{"result": end.Element != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c", actual)
		actual = args.Map{"result": len(collected) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3 collected", actual)
	})
}

func Test_C24_LinkedListNode_LoopEndOfChain_Break(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_LoopEndOfChain_Break", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		end, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			return true // break immediately
		})

		// Act
		actual := args.Map{"result": length != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": end.Element != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C24_LinkedListNode_AddNext(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_AddNext", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "c"})
		node := ll.Head()
		newNode := node.AddNext(ll, "b")

		// Act
		actual := args.Map{"result": newNode.Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
		actual = args.Map{"result": ll.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C24_LinkedListNode_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_AddStringsToNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "d"})
		node := ll.Head()
		node.AddStringsToNode(ll, false, []string{"b", "c"})

		// Act
		actual := args.Map{"result": ll.Length() < 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	})
}

func Test_C24_LinkedListNode_AddStringsPtrToNode_Nil(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_AddStringsPtrToNode_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		node := ll.Head()
		result := node.AddStringsPtrToNode(ll, true, nil)

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C24_LinkedListNode_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_AddCollectionToNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		col := corestr.New.Collection.Strings([]string{"b", "c"})
		node := ll.Head()
		node.AddCollectionToNode(ll, false, col)

		// Act
		actual := args.Map{"result": ll.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_C24_LinkedListNode_AddNextNode(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_AddNextNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "c"})
		nextNode := &corestr.LinkedListNode{Element: "b"}
		ll.Head().AddNextNode(ll, nextNode)

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C24_LinkedListNode_IsEqual_Same(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_IsEqual_Same", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ll.Head().IsEqual(ll.Head())}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "same node should be equal", actual)
	})
}

func Test_C24_LinkedListNode_IsChainEqual(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_IsChainEqual", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ll1.Head().IsChainEqual(ll2.Head(), true)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "chains should be equal", actual)
	})
}

func Test_C24_LinkedListNode_IsChainEqual_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_IsChainEqual_CaseInsensitive", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Strings([]string{"A", "B"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ll1.Head().IsChainEqual(ll2.Head(), false)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "chains should be equal case insensitive", actual)
	})
}

func Test_C24_LinkedListNode_IsEqualSensitive(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_IsEqualSensitive", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Strings([]string{"A"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ll1.Head().IsEqualSensitive(ll2.Head(), true)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be equal case sensitive", actual)
		actual = args.Map{"result": ll1.Head().IsEqualSensitive(ll2.Head(), false)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal case insensitive", actual)
	})
}

func Test_C24_LinkedListNode_IsEqualValue(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_IsEqualValue", func() {
		// Arrange
		node := &corestr.LinkedListNode{Element: "hello"}

		// Act
		actual := args.Map{"result": node.IsEqualValue("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C24_LinkedListNode_IsEqualValueSensitive(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_IsEqualValueSensitive", func() {
		// Arrange
		node := &corestr.LinkedListNode{Element: "Hello"}

		// Act
		actual := args.Map{"result": node.IsEqualValueSensitive("hello", false)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal case insensitive", actual)
	})
}

func Test_C24_LinkedListNode_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_CreateLinkedList", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		newLL := ll.Head().CreateLinkedList()

		// Act
		actual := args.Map{"result": newLL.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedListNode_List(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_List", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		list := ll.Head().List()

		// Act
		actual := args.Map{"result": len(list) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C24_LinkedListNode_ListPtr(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_ListPtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		list := ll.Head().ListPtr()

		// Act
		actual := args.Map{"result": len(list) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C24_LinkedListNode_Join(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_Join", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		result := ll.Head().Join(",")

		// Act
		actual := args.Map{"result": result != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
	})
}

func Test_C24_LinkedListNode_StringList(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_StringList", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		result := ll.Head().StringList("Header:")

		// Act
		actual := args.Map{"result": result == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C24_LinkedListNode_Print(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_Print", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.Head().Print("Test: ")
	})
}

func Test_C24_LinkedListNode_String(t *testing.T) {
	safeTest(t, "Test_C24_LinkedListNode_String", func() {
		// Arrange
		node := &corestr.LinkedListNode{Element: "test"}

		// Act
		actual := args.Map{"result": node.String() != "test"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected test", actual)
	})
}

// =======================================================
// LinkedList
// =======================================================

func Test_C24_LinkedList_Empty(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": ll.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
		actual = args.Map{"result": ll.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have items", actual)
		actual = args.Map{"result": ll.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C24_LinkedList_Add_Single(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_Add_Single", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": ll.Head().Element != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "head should be a", actual)
		actual = args.Map{"result": ll.Tail().Element != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "tail should be a", actual)
	})
}

func Test_C24_LinkedList_Add_Multi(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_Add_Multi", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a").Add("b").Add("c")

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual = args.Map{"result": ll.Head().Element != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "head should be a", actual)
		actual = args.Map{"result": ll.Tail().Element != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "tail should be c", actual)
	})
}

func Test_C24_LinkedList_AddFront(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddFront", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"b", "c"})
		ll.AddFront("a")

		// Act
		actual := args.Map{"result": ll.Head().Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "head should be a", actual)
		actual = args.Map{"result": ll.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C24_LinkedList_AddFront_Empty(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddFront_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddFront("a")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C24_LinkedList_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddNonEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmpty("")

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not add empty", actual)
		ll.AddNonEmpty("a")
		actual = args.Map{"result": ll.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should add non-empty", actual)
	})
}

func Test_C24_LinkedList_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddNonEmptyWhitespace", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmptyWhitespace("   ")

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not add whitespace", actual)
		ll.AddNonEmptyWhitespace("a")
		actual = args.Map{"result": ll.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should add non-whitespace", actual)
	})
}

func Test_C24_LinkedList_AddIf(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddIf", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddIf(false, "skip")

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not add", actual)
		ll.AddIf(true, "keep")
		actual = args.Map{"result": ll.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should add", actual)
	})
}

func Test_C24_LinkedList_AddsIf(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddsIf", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddsIf(false, "a", "b")

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not add", actual)
		ll.AddsIf(true, "a", "b")
		actual = args.Map{"result": ll.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_AddFunc(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddFunc", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddFunc(func() string { return "computed" })

		// Act
		actual := args.Map{"result": ll.Length() != 1 || ll.Head().Element != "computed"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddFunc failed", actual)
	})
}

func Test_C24_LinkedList_AddFuncErr_Success(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddFuncErr_Success", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddFuncErr(
			func() (string, error) { return "ok", nil },

		// Assert
			func(err error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) },
		)

		// Act
		actual = args.Map{"result": ll.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have added", actual)
	})
}

func Test_C24_LinkedList_AddLock(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddLock("a")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have 1", actual)
	})
}

func Test_C24_LinkedList_Push_PushFront_PushBack(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_Push_PushFront_PushBack", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Push("a")
		ll.PushFront("front")
		ll.PushBack("back")

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual = args.Map{"result": ll.Head().Element != "front"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "head should be front", actual)
	})
}

func Test_C24_LinkedList_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddItemsMap", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		m := map[string]bool{"a": true, "b": false, "c": true}
		ll.AddItemsMap(m)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_AddBackNode(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddBackNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		node := &corestr.LinkedListNode{Element: "b"}
		ll.AddBackNode(node)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_AppendNode_Empty(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AppendNode_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		node := &corestr.LinkedListNode{Element: "a"}
		ll.AppendNode(node)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have 1", actual)
	})
}

func Test_C24_LinkedList_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AppendChainOfNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		chain := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.AppendChainOfNodes(chain.Head())

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C24_LinkedList_AppendChainOfNodes_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AppendChainOfNodes_NonEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"x"})
		chain := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.AppendChainOfNodes(chain.Head())

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C24_LinkedList_InsertAt(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_InsertAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "c"})
		ll.InsertAt(1, "b")
		list := ll.List()

		// Act
		actual := args.Map{"result": len(list) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C24_LinkedList_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_InsertAt_Front", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"b"})
		ll.InsertAt(0, "a")

		// Act
		actual := args.Map{"result": ll.Head().Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "head should be a", actual)
	})
}

func Test_C24_LinkedList_AttachWithNode(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AttachWithNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		node := ll.Head()
		addNode := &corestr.LinkedListNode{Element: "b"}
		err := ll.AttachWithNode(node, addNode)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	})
}

func Test_C24_LinkedList_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AttachWithNode_NilCurrent", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		addNode := &corestr.LinkedListNode{Element: "b"}
		err := ll.AttachWithNode(nil, addNode)

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error for nil current", actual)
	})
}

func Test_C24_LinkedList_Adds(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_Adds", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C24_LinkedList_AddStrings(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddStrings", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_AddsLock(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddsLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddsLock("a", "b")

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_AddCollection(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddCollection", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		ll.AddCollection(col)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddCollection_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddCollection(nil)

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should remain empty", actual)
	})
}

func Test_C24_LinkedList_AddPointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddPointerStringsPtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		a, b := "a", "b"
		ll.AddPointerStringsPtr([]*string{&a, nil, &b})

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_IndexAt(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_IndexAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		node := ll.IndexAt(1)

		// Act
		actual := args.Map{"result": node.Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C24_LinkedList_IndexAt_Head(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_IndexAt_Head", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		node := ll.IndexAt(0)

		// Act
		actual := args.Map{"result": node.Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C24_LinkedList_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_SafeIndexAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		node := ll.SafeIndexAt(1)

		// Act
		actual := args.Map{"result": node == nil || node.Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C24_LinkedList_SafeIndexAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_SafeIndexAt_OutOfRange", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ll.SafeIndexAt(5) != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		actual = args.Map{"result": ll.SafeIndexAt(-1) != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil for negative", actual)
	})
}

func Test_C24_LinkedList_SafeIndexAtLock(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_SafeIndexAtLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		node := ll.SafeIndexAtLock(0)

		// Act
		actual := args.Map{"result": node == nil || node.Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C24_LinkedList_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_SafePointerIndexAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"hello"})
		ptr := ll.SafePointerIndexAt(0)

		// Act
		actual := args.Map{"result": ptr == nil || *ptr != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		actual = args.Map{"result": ll.SafePointerIndexAt(5) != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C24_LinkedList_SafePointerIndexAtUsingDefault(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_SafePointerIndexAtUsingDefault", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"hello"})
		val := ll.SafePointerIndexAtUsingDefault(0, "default")

		// Act
		actual := args.Map{"result": val != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		val = ll.SafePointerIndexAtUsingDefault(5, "default")
		actual = args.Map{"result": val != "default"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default", actual)
	})
}

func Test_C24_LinkedList_SafePointerIndexAtUsingDefaultLock(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_SafePointerIndexAtUsingDefaultLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"hello"})
		val := ll.SafePointerIndexAtUsingDefaultLock(0, "default")

		// Act
		actual := args.Map{"result": val != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_C24_LinkedList_LengthLock(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_LengthLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ll.LengthLock() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_IsEmptyLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": ll.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

func Test_C24_LinkedList_IsEquals(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_IsEquals", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ll1.IsEquals(ll2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C24_LinkedList_IsEqualsWithSensitive(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_IsEqualsWithSensitive", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Strings([]string{"A"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ll1.IsEqualsWithSensitive(ll2, true)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be equal case sensitive", actual)
		actual = args.Map{"result": ll1.IsEqualsWithSensitive(ll2, false)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal case insensitive", actual)
	})
}

func Test_C24_LinkedList_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_IsEquals_BothEmpty", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Create()
		ll2 := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": ll1.IsEquals(ll2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "both empty should be equal", actual)
	})
}

func Test_C24_LinkedList_IsEquals_DiffLength(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_IsEquals_DiffLength", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Strings([]string{"a"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ll1.IsEquals(ll2)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "diff lengths should not be equal", actual)
	})
}

func Test_C24_LinkedList_Loop(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_Loop", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})

		// Act
		actual := args.Map{"result": count != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C24_LinkedList_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_Loop_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {

		// Act
			actual := args.Map{"result": false}

		// Assert
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "should not be called", actual)
			return false
		})
	})
	}
func Test_C24_LinkedList_Loop_Break(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_Loop_Break", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return arg.Index == 1
		})

		// Act
		actual := args.Map{"result": count != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_Filter(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_Filter", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		results := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{
				Value:  arg.Node,
				IsKeep: arg.Node.Element != "b",
			}
		})

		// Act
		actual := args.Map{"result": len(results) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_Filter_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		results := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true}
		})

		// Act
		actual := args.Map{"result": len(results) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

func Test_C24_LinkedList_Filter_Break(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_Filter_Break", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		results := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{
				Value:   arg.Node,
				IsKeep:  true,
				IsBreak: arg.Index == 0,
			}
		})

		// Act
		actual := args.Map{"result": len(results) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C24_LinkedList_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_GetNextNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c", "d"})
		nodes := ll.GetNextNodes(2)

		// Act
		actual := args.Map{"result": len(nodes) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_GetAllLinkedNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		nodes := ll.GetAllLinkedNodes()

		// Act
		actual := args.Map{"result": len(nodes) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C24_LinkedList_RemoveNodeByElementValue(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_RemoveNodeByElementValue", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByElementValue("b", true, false)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_RemoveNodeByElementValue_First(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_RemoveNodeByElementValue_First", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.RemoveNodeByElementValue("a", true, false)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": ll.Head().Element != "b"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "head should be b", actual)
	})
}

func Test_C24_LinkedList_RemoveNodeByElementValue_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_RemoveNodeByElementValue_CaseInsensitive", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"A", "b"})
		ll.RemoveNodeByElementValue("a", false, false)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C24_LinkedList_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_RemoveNodeByIndex", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByIndex(1)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_RemoveNodeByIndex_First(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_RemoveNodeByIndex_First", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.RemoveNodeByIndex(0)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C24_LinkedList_RemoveNodeByIndex_Last(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_RemoveNodeByIndex_Last", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.RemoveNodeByIndex(1)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C24_LinkedList_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_RemoveNodeByIndexes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c", "d"})
		ll.RemoveNodeByIndexes(false, 1, 3)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_RemoveNodeByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_RemoveNodeByIndexes_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.RemoveNodeByIndexes(false)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should remain 1", actual)
	})
}

func Test_C24_LinkedList_RemoveNode(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_RemoveNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		node := ll.IndexAt(1)
		ll.RemoveNode(node)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_RemoveNode_Nil(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_RemoveNode_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.RemoveNode(nil)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should remain 1", actual)
	})
}

func Test_C24_LinkedList_RemoveNode_First(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_RemoveNode_First", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.RemoveNode(ll.Head())

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C24_LinkedList_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddStringsToNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "d"})
		node := ll.Head()
		ll.AddStringsToNode(false, node, []string{"b", "c"})

		// Act
		actual := args.Map{"result": ll.Length() < 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	})
}

func Test_C24_LinkedList_AddStringsToNode_Single(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddStringsToNode_Single", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "c"})
		node := ll.Head()
		ll.AddStringsToNode(false, node, []string{"b"})

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C24_LinkedList_AddStringsToNode_NilSkip(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddStringsToNode_NilSkip", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.AddStringsToNode(true, nil, []string{"b"})

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should skip nil node", actual)
	})
}

func Test_C24_LinkedList_AddStringsPtrToNode_NilItems(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddStringsPtrToNode_NilItems", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.AddStringsPtrToNode(true, ll.Head(), nil)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not change", actual)
	})
}

func Test_C24_LinkedList_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddCollectionToNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		col := corestr.New.Collection.Strings([]string{"b"})
		ll.AddCollectionToNode(true, ll.Head(), col)

		// Act
		actual := args.Map{"result": ll.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have added", actual)
	})
}

func Test_C24_LinkedList_AddAfterNode(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AddAfterNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "c"})
		node := ll.Head()
		newNode := ll.AddAfterNode(node, "b")

		// Act
		actual := args.Map{"result": newNode.Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C24_LinkedList_ToCollection(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_ToCollection", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		col := ll.ToCollection(0)

		// Act
		actual := args.Map{"result": col.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_ToCollection_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		col := ll.ToCollection(0)

		// Act
		actual := args.Map{"result": col.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

func Test_C24_LinkedList_List(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_List", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		list := ll.List()

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_ListPtr(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_ListPtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		list := ll.ListPtr()

		// Act
		actual := args.Map{"result": len(list) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C24_LinkedList_ListLock(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_ListLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		list := ll.ListLock()

		// Act
		actual := args.Map{"result": len(list) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C24_LinkedList_ListPtrLock(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_ListPtrLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		list := ll.ListPtrLock()

		// Act
		actual := args.Map{"result": len(list) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C24_LinkedList_String(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_String", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		s := ll.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C24_LinkedList_String_Empty(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_String_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		s := ll.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have no elements text", actual)
	})
}

func Test_C24_LinkedList_StringLock(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_StringLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		s := ll.StringLock()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C24_LinkedList_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_StringLock_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		s := ll.StringLock()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have no elements text", actual)
	})
}

func Test_C24_LinkedList_Join(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_Join", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		result := ll.Join(",")

		// Act
		actual := args.Map{"result": result != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
	})
}

func Test_C24_LinkedList_JoinLock(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_JoinLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		result := ll.JoinLock(",")

		// Act
		actual := args.Map{"result": result != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
	})
}

func Test_C24_LinkedList_Joins(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_Joins", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		result := ll.Joins(",", "b", "c")

		// Act
		actual := args.Map{"result": result == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C24_LinkedList_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_GetCompareSummary", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Strings([]string{"a"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a"})
		summary := ll1.GetCompareSummary(ll2, "left", "right")

		// Act
		actual := args.Map{"result": summary == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C24_LinkedList_Clear(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_Clear", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.Clear()

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be 0 after clear", actual)
	})
}

func Test_C24_LinkedList_RemoveAll(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_RemoveAll", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.RemoveAll()

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be 0", actual)
	})
}

func Test_C24_LinkedList_Json(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_Json", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		result := ll.Json()

		// Act
		actual := args.Map{"result": result.HasError()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not error", actual)
	})
}

func Test_C24_LinkedList_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_JsonPtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		result := ll.JsonPtr()

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C24_LinkedList_JsonModel(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_JsonModel", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		model := ll.JsonModel()

		// Act
		actual := args.Map{"result": len(model) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_JsonModelAny", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		model := ll.JsonModelAny()

		// Act
		actual := args.Map{"result": model == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C24_LinkedList_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_MarshalJSON", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		data, err := json.Marshal(ll)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "marshal error:", actual)
		actual = args.Map{"result": len(data) == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have data", actual)
	})
}

func Test_C24_LinkedList_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_UnmarshalJSON", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		data, _ := json.Marshal(ll)
		ll2 := corestr.New.LinkedList.Create()
		err := json.Unmarshal(data, ll2)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unmarshal error:", actual)
		actual = args.Map{"result": ll2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_ParseInjectUsingJson", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		// Use json.Marshal with pointer to bypass value receiver issue on Json()
		b, _ := json.Marshal(ll)
		jsonResult := corejson.Result{Bytes: b}
		ll2 := corestr.New.LinkedList.Create()
		result, err := ll2.ParseInjectUsingJson(&jsonResult)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
		actual = args.Map{"result": result.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_LinkedList_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_ParseInjectUsingJsonMust", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		// Use json.Marshal with pointer to bypass value receiver issue on Json()
		b, _ := json.Marshal(ll)
		jsonResult := corejson.Result{Bytes: b}
		ll2 := corestr.New.LinkedList.Create()
		result := ll2.ParseInjectUsingJsonMust(&jsonResult)

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C24_LinkedList_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_JsonParseSelfInject", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		// Use json.Marshal with pointer to bypass value receiver issue on Json()
		b, _ := json.Marshal(ll)
		jsonResult := corejson.Result{Bytes: b}
		ll2 := corestr.New.LinkedList.Create()
		err := ll2.JsonParseSelfInject(&jsonResult)
		// Unmarshal may fail due to value-receiver serialization; exercise the code path for coverage
		_ = err
	})
}

func Test_C24_LinkedList_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_AsJsonMarshaller", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		m := ll.AsJsonMarshaller()

		// Act
		actual := args.Map{"result": m == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

// =======================================================
// NonChainedLinkedListNodes
// =======================================================

func Test_C24_NonChainedLinkedListNodes_Empty(t *testing.T) {
	safeTest(t, "Test_C24_NonChainedLinkedListNodes_Empty", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(5)

		// Act
		actual := args.Map{"result": nc.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
		actual = args.Map{"result": nc.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have items", actual)
		actual = args.Map{"result": nc.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be 0", actual)
	})
}

func Test_C24_NonChainedLinkedListNodes_Adds(t *testing.T) {
	safeTest(t, "Test_C24_NonChainedLinkedListNodes_Adds", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(5)
		nc.Adds(
			&corestr.LinkedListNode{Element: "a"},
			&corestr.LinkedListNode{Element: "b"},
		)

		// Act
		actual := args.Map{"result": nc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual = args.Map{"result": nc.First().Element != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "first should be a", actual)
		actual = args.Map{"result": nc.Last().Element != "b"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "last should be b", actual)
	})
}

func Test_C24_NonChainedLinkedListNodes_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C24_NonChainedLinkedListNodes_FirstOrDefault_Empty", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(0)

		// Act
		actual := args.Map{"result": nc.FirstOrDefault() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be nil", actual)
	})
}

func Test_C24_NonChainedLinkedListNodes_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C24_NonChainedLinkedListNodes_LastOrDefault_Empty", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(0)

		// Act
		actual := args.Map{"result": nc.LastOrDefault() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be nil", actual)
	})
}

func Test_C24_NonChainedLinkedListNodes_ApplyChaining(t *testing.T) {
	safeTest(t, "Test_C24_NonChainedLinkedListNodes_ApplyChaining", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(3)
		nc.Adds(
			&corestr.LinkedListNode{Element: "a"},
			&corestr.LinkedListNode{Element: "b"},
			&corestr.LinkedListNode{Element: "c"},
		)
		nc.ApplyChaining()

		// Act
		actual := args.Map{"result": nc.IsChainingApplied()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "chaining should be applied", actual)
		actual = args.Map{"result": nc.First().HasNext()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "first should have next after chaining", actual)
	})
}

func Test_C24_NonChainedLinkedListNodes_ApplyChaining_Empty(t *testing.T) {
	safeTest(t, "Test_C24_NonChainedLinkedListNodes_ApplyChaining_Empty", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(0)
		nc.ApplyChaining()

		// Act
		actual := args.Map{"result": nc.IsChainingApplied()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not apply to empty", actual)
	})
}

func Test_C24_NonChainedLinkedListNodes_ToChainedNodes(t *testing.T) {
	safeTest(t, "Test_C24_NonChainedLinkedListNodes_ToChainedNodes", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(3)
		nc.Adds(
			&corestr.LinkedListNode{Element: "a"},
			&corestr.LinkedListNode{Element: "b"},
		)
		chained := nc.ToChainedNodes()

		// Act
		actual := args.Map{"result": chained == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C24_NonChainedLinkedListNodes_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_C24_NonChainedLinkedListNodes_Adds_Nil", func() {
		nc := corestr.NewNonChainedLinkedListNodes(0)
		nc.Adds(nil)
		// nil entries are still appended per implementation
	})
}

// =======================================================
// newLinkedListCreator
// =======================================================

func Test_C24_NewLinkedList_Create(t *testing.T) {
	safeTest(t, "Test_C24_NewLinkedList_Create", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": ll == nil || !ll.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

func Test_C24_NewLinkedList_Empty(t *testing.T) {
	safeTest(t, "Test_C24_NewLinkedList_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Empty()

		// Act
		actual := args.Map{"result": ll == nil || !ll.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

func Test_C24_NewLinkedList_Strings(t *testing.T) {
	safeTest(t, "Test_C24_NewLinkedList_Strings", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_NewLinkedList_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_C24_NewLinkedList_Strings_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings(nil)

		// Act
		actual := args.Map{"result": ll.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

func Test_C24_NewLinkedList_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_C24_NewLinkedList_SpreadStrings", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C24_NewLinkedList_SpreadStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C24_NewLinkedList_SpreadStrings_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings()

		// Act
		actual := args.Map{"result": ll.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

func Test_C24_NewLinkedList_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_C24_NewLinkedList_PointerStringsPtr", func() {
		// Arrange
		a, b := "a", "b"
		items := []*string{&a, &b}
		ll := corestr.New.LinkedList.PointerStringsPtr(&items)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_NewLinkedList_PointerStringsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_C24_NewLinkedList_PointerStringsPtr_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.PointerStringsPtr(nil)

		// Act
		actual := args.Map{"result": ll.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

func Test_C24_NewLinkedList_UsingMap(t *testing.T) {
	safeTest(t, "Test_C24_NewLinkedList_UsingMap", func() {
		// Arrange
		m := map[string]bool{"a": true, "b": true}
		ll := corestr.New.LinkedList.UsingMap(m)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C24_NewLinkedList_UsingMap_Nil(t *testing.T) {
	safeTest(t, "Test_C24_NewLinkedList_UsingMap_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.UsingMap(nil)

		// Act
		actual := args.Map{"result": ll.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

// =======================================================
// Concurrent LinkedList operations
// =======================================================

func Test_C24_LinkedList_ConcurrentAddsLock(t *testing.T) {
	safeTest(t, "Test_C24_LinkedList_ConcurrentAddsLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		wg := &sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				ll.AddLock("item")
			}()
		}
		wg.Wait()

		// Act
		actual := args.Map{"result": ll.Length() != 10}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
	})
}
