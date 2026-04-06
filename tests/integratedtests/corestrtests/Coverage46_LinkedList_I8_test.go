package corestrtests

import (
	"encoding/json"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ===================== LinkedListNode =====================

func Test_C46_LinkedListNode_HasNext_NoNext(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_HasNext_NoNext", func() {
		node := &corestr.LinkedListNode{Element: "a"}
		actual := args.Map{"result": node.HasNext()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no next", actual)
	})
}

func Test_C46_LinkedListNode_EndOfChain_Single(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_EndOfChain_Single", func() {
		node := &corestr.LinkedListNode{Element: "a"}
		end, length := node.EndOfChain()
		actual := args.Map{"result": end != node || length != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected self and length 1, got length", actual)
	})
}

func Test_C46_LinkedListNode_Clone(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_Clone", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		node := ll.Head()
		cloned := node.Clone()
		actual := args.Map{"result": cloned.Element != "a" || cloned.HasNext()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "clone should copy element but not next", actual)
	})
}

func Test_C46_LinkedListNode_List(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_List", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		list := ll.Head().List()
		actual := args.Map{"result": len(list) != 3 || list[2] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected list:", actual)
	})
}

func Test_C46_LinkedListNode_ListPtr(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_ListPtr", func() {
		ll := corestr.New.LinkedList.Strings([]string{"x"})
		list := ll.Head().ListPtr()
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C46_LinkedListNode_Join(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_Join", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		j := ll.Head().Join(",")
		actual := args.Map{"result": j != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "got", actual)
	})
}

func Test_C46_LinkedListNode_StringList(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_StringList", func() {
		ll := corestr.New.LinkedList.Strings([]string{"x"})
		s := ll.Head().StringList("H:")
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C46_LinkedListNode_Print(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_Print", func() {
		ll := corestr.New.LinkedList.Strings([]string{"x"})
		ll.Head().Print("test: ")
	})
}

func Test_C46_LinkedListNode_String(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_String", func() {
		node := &corestr.LinkedListNode{Element: "hello"}
		actual := args.Map{"result": node.String() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_C46_LinkedListNode_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_IsEqual_BothNil", func() {
		var n1, n2 *corestr.LinkedListNode
		_ = n1
		_ = n2
		// Can't call method on nil, test via chain
	})
}

func Test_C46_LinkedListNode_IsEqual_Same(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_IsEqual_Same", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		node := ll.Head()
		actual := args.Map{"result": node.IsEqual(node)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "same node should be equal", actual)
	})
}

func Test_C46_LinkedListNode_IsEqual_DifferentValues(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_IsEqual_DifferentValues", func() {
		n1 := &corestr.LinkedListNode{Element: "a"}
		n2 := &corestr.LinkedListNode{Element: "b"}
		actual := args.Map{"result": n1.IsEqual(n2)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "different elements should not be equal", actual)
	})
}

func Test_C46_LinkedListNode_IsEqualSensitive(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_IsEqualSensitive", func() {
		n1 := &corestr.LinkedListNode{Element: "Hello"}
		n2 := &corestr.LinkedListNode{Element: "hello"}
		actual := args.Map{"result": n1.IsEqualSensitive(n2, true)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "case sensitive should fail", actual)
		actual := args.Map{"result": n1.IsEqualSensitive(n2, false)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "case insensitive should pass", actual)
	})
}

func Test_C46_LinkedListNode_IsChainEqual(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_IsChainEqual", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		actual := args.Map{"result": ll1.Head().IsChainEqual(ll2.Head(), true)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "chains should be equal", actual)
	})
}

func Test_C46_LinkedListNode_IsChainEqual_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_IsChainEqual_CaseInsensitive", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"A", "B"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		actual := args.Map{"result": ll1.Head().IsChainEqual(ll2.Head(), false)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "chains should be equal case-insensitive", actual)
	})
}

func Test_C46_LinkedListNode_IsEqualValue(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_IsEqualValue", func() {
		n := &corestr.LinkedListNode{Element: "test"}
		actual := args.Map{"result": n.IsEqualValue("test")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
		actual := args.Map{"result": n.IsEqualValue("other")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not match", actual)
	})
}

func Test_C46_LinkedListNode_IsEqualValueSensitive(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_IsEqualValueSensitive", func() {
		n := &corestr.LinkedListNode{Element: "Test"}
		actual := args.Map{"result": n.IsEqualValueSensitive("test", false)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "case-insensitive should match", actual)
		actual := args.Map{"result": n.IsEqualValueSensitive("test", true)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "case-sensitive should not match", actual)
	})
}

func Test_C46_LinkedListNode_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_CreateLinkedList", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		newLL := ll.Head().CreateLinkedList()
		actual := args.Map{"result": newLL.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C46_LinkedListNode_LoopEndOfChain(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_LoopEndOfChain", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		count := 0
		end, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})
		actual := args.Map{"result": count != 3 || length != 3 || end.Element != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected count= length= end=", actual)
	})
}

func Test_C46_LinkedListNode_LoopEndOfChain_BreakFirst(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_LoopEndOfChain_BreakFirst", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		end, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			return true // break immediately
		})
		actual := args.Map{"result": length != 1 || end.Element != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected break at first, length= end=", actual)
	})
}

func Test_C46_LinkedListNode_AddNext(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_AddNext", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "c"})
		node := ll.Head()
		node.AddNext(ll, "b")
		actual := args.Map{"result": ll.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C46_LinkedListNode_AddNextNode(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_AddNextNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "c"})
		node := ll.Head()
		newNode := &corestr.LinkedListNode{Element: "b"}
		node.AddNextNode(ll, newNode)
		actual := args.Map{"result": ll.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C46_LinkedListNode_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_AddStringsToNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "d"})
		node := ll.Head()
		node.AddStringsToNode(ll, false, []string{"b", "c"})
		actual := args.Map{"result": ll.Length() < 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected items added", actual)
	})
}

func Test_C46_LinkedListNode_AddStringsPtrToNode_Nil(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_AddStringsPtrToNode_Nil", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		node := ll.Head()
		result := node.AddStringsPtrToNode(ll, true, nil)
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil ptr should not add", actual)
	})
}

func Test_C46_LinkedListNode_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_AddCollectionToNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "d"})
		node := ll.Head()
		col := corestr.New.Collection.Strings([]string{"b", "c"})
		node.AddCollectionToNode(ll, false, col)
		actual := args.Map{"result": ll.Length() < 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected items added", actual)
	})
}

// ===================== LinkedList =====================

func Test_C46_LinkedList_Create_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Create_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"result": ll.IsEmpty() || ll.Length() != 0}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual := args.Map{"result": ll.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have items", actual)
	})
}

func Test_C46_LinkedList_Add_Single(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Add_Single", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("hello")
		actual := args.Map{"result": ll.Length() != 1 || ll.Head().Element != "hello" || ll.Tail().Element != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "single add failed", actual)
	})
}

func Test_C46_LinkedList_Add_Multiple(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Add_Multiple", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a").Add("b").Add("c")
		actual := args.Map{"result": ll.Length() != 3 || ll.Head().Element != "a" || ll.Tail().Element != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "multi add failed", actual)
	})
}

func Test_C46_LinkedList_Adds(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Adds", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		actual := args.Map{"result": ll.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "adds failed", actual)
	})
}

func Test_C46_LinkedList_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Adds_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds()
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C46_LinkedList_AddStrings(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddStrings", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings([]string{"x", "y"})
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_AddLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddLock("safe")
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C46_LinkedList_AddsLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddsLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddsLock("a", "b")
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_AddFront_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddFront_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFront("first")
		actual := args.Map{"result": ll.Length() != 1 || ll.Head().Element != "first"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "add front to empty failed", actual)
	})
}

func Test_C46_LinkedList_AddFront_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddFront_NonEmpty", func() {
		ll := corestr.New.LinkedList.Strings([]string{"b", "c"})
		ll.AddFront("a")
		actual := args.Map{"result": ll.Head().Element != "a" || ll.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "add front failed", actual)
	})
}

func Test_C46_LinkedList_PushFront(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_PushFront", func() {
		ll := corestr.New.LinkedList.Strings([]string{"b"})
		ll.PushFront("a")
		actual := args.Map{"result": ll.Head().Element != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "push front failed", actual)
	})
}

func Test_C46_LinkedList_Push(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Push", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Push("x")
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C46_LinkedList_PushBack(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_PushBack", func() {
		ll := corestr.New.LinkedList.Create()
		ll.PushBack("x")
		actual := args.Map{"result": ll.Tail().Element != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "push back failed", actual)
	})
}

func Test_C46_LinkedList_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddNonEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmpty("")
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty string should not be added", actual)
		ll.AddNonEmpty("x")
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "non-empty should be added", actual)
	})
}

func Test_C46_LinkedList_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddNonEmptyWhitespace", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmptyWhitespace("   ")
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "whitespace should not be added", actual)
		ll.AddNonEmptyWhitespace("x")
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "non-whitespace should be added", actual)
	})
}

func Test_C46_LinkedList_AddIf(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddIf", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddIf(false, "skip")
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should skip", actual)
		ll.AddIf(true, "add")
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should add", actual)
	})
}

func Test_C46_LinkedList_AddsIf(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddsIf", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddsIf(false, "a", "b")
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should skip", actual)
		ll.AddsIf(true, "a", "b")
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should add", actual)
	})
}

func Test_C46_LinkedList_AddFunc(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddFunc", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFunc(func() string { return "generated" })
		actual := args.Map{"result": ll.Length() != 1 || ll.Head().Element != "generated"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "add func failed", actual)
	})
}

func Test_C46_LinkedList_AddFuncErr_Success(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddFuncErr_Success", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) })
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C46_LinkedList_AddFuncErr_Error(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddFuncErr_Error", func() {
		ll := corestr.New.LinkedList.Create()
		called := false
		ll.AddFuncErr(
			func() (string, error) { return "", json.Unmarshal([]byte("invalid"), nil) },
			func(e error) { called = true },
		)
		actual := args.Map{"result": called}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "error handler should be called", actual)
	})
}

func Test_C46_LinkedList_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddItemsMap", func() {
		ll := corestr.New.LinkedList.Create()
		m := map[string]bool{"a": true, "b": false, "c": true}
		ll.AddItemsMap(m)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_AddItemsMap_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddItemsMap_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{})
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C46_LinkedList_AppendNode_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AppendNode_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AppendNode(&corestr.LinkedListNode{Element: "x"})
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C46_LinkedList_AppendNode_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AppendNode_NonEmpty", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.AppendNode(&corestr.LinkedListNode{Element: "b"})
		actual := args.Map{"result": ll.Length() != 2 || ll.Tail().Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "append failed", actual)
	})
}

func Test_C46_LinkedList_AddBackNode(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddBackNode", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddBackNode(&corestr.LinkedListNode{Element: "x"})
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C46_LinkedList_AppendChainOfNodes_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AppendChainOfNodes_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		chain := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.AppendChainOfNodes(chain.Head())
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_AppendChainOfNodes_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AppendChainOfNodes_NonEmpty", func() {
		ll := corestr.New.LinkedList.Strings([]string{"x"})
		chain := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.AppendChainOfNodes(chain.Head())
		actual := args.Map{"result": ll.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C46_LinkedList_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_InsertAt_Front", func() {
		ll := corestr.New.LinkedList.Strings([]string{"b", "c"})
		ll.InsertAt(0, "a")
		actual := args.Map{"result": ll.Head().Element != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "insert at front failed", actual)
	})
}

func Test_C46_LinkedList_InsertAt_Middle(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_InsertAt_Middle", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "c"})
		ll.InsertAt(1, "b")
		list := ll.List()
		actual := args.Map{"result": len(list) < 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3+", actual)
	})
}

func Test_C46_LinkedList_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AttachWithNode_NilCurrent", func() {
		ll := corestr.New.LinkedList.Create()
		err := ll.AttachWithNode(nil, &corestr.LinkedListNode{Element: "x"})
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error for nil current node", actual)
	})
}

func Test_C46_LinkedList_AttachWithNode_NextNotNil(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AttachWithNode_NextNotNil", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		err := ll.AttachWithNode(ll.Head(), &corestr.LinkedListNode{Element: "x"})
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error for current.next not nil", actual)
	})
}

func Test_C46_LinkedList_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddCollection_Nil", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddCollection(nil)
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil collection should not add", actual)
	})
}

func Test_C46_LinkedList_AddCollection(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddCollection", func() {
		ll := corestr.New.LinkedList.Create()
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		ll.AddCollection(col)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_AddPointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddPointerStringsPtr", func() {
		ll := corestr.New.LinkedList.Create()
		s1 := "a"
		s2 := "b"
		ll.AddPointerStringsPtr([]*string{&s1, nil, &s2})
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 (skip nil)", actual)
	})
}

func Test_C46_LinkedList_IndexAt(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IndexAt", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		n := ll.IndexAt(1)
		actual := args.Map{"result": n.Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
		n0 := ll.IndexAt(0)
		actual := args.Map{"result": n0.Element != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C46_LinkedList_IndexAt_Negative(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IndexAt_Negative", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		n := ll.IndexAt(-1)
		actual := args.Map{"result": n != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "negative index should return nil", actual)
	})
}

func Test_C46_LinkedList_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_SafeIndexAt", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		n := ll.SafeIndexAt(1)
		actual := args.Map{"result": n == nil || n.Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
		n2 := ll.SafeIndexAt(5)
		actual := args.Map{"result": n2 != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "out of range should return nil", actual)
		n3 := ll.SafeIndexAt(-1)
		actual := args.Map{"result": n3 != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "negative should return nil", actual)
	})
}

func Test_C46_LinkedList_SafeIndexAtLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_SafeIndexAtLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		n := ll.SafeIndexAtLock(0)
		actual := args.Map{"result": n == nil || n.Element != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C46_LinkedList_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_SafePointerIndexAt", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		p := ll.SafePointerIndexAt(0)
		actual := args.Map{"result": p == nil || *p != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		p2 := ll.SafePointerIndexAt(99)
		actual := args.Map{"result": p2 != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "out of range should return nil", actual)
	})
}

func Test_C46_LinkedList_SafePointerIndexAtUsingDefault(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_SafePointerIndexAtUsingDefault", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		v := ll.SafePointerIndexAtUsingDefault(0, "def")
		actual := args.Map{"result": v != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		v2 := ll.SafePointerIndexAtUsingDefault(99, "def")
		actual := args.Map{"result": v2 != "def"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default", actual)
	})
}

func Test_C46_LinkedList_SafePointerIndexAtUsingDefaultLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_SafePointerIndexAtUsingDefaultLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		v := ll.SafePointerIndexAtUsingDefaultLock(0, "def")
		actual := args.Map{"result": v != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C46_LinkedList_Loop(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Loop", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		elements := []string{}
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			elements = append(elements, arg.CurrentNode.Element)
			return false
		})
		actual := args.Map{"result": len(elements) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C46_LinkedList_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Loop_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			actual := args.Map{"result": false}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "should not be called", actual)
			return false
		})
	})
}

func Test_C46_LinkedList_Loop_Break(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Loop_Break", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return true // break immediately
		})
		actual := args.Map{"result": count != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 iteration", actual)
	})
}

func Test_C46_LinkedList_Loop_BreakMiddle(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Loop_BreakMiddle", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return arg.Index == 1
		})
		actual := args.Map{"result": count != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_Filter(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Filter", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "bb", "c"})
		result := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{
				Value:  arg.Node,
				IsKeep: len(arg.Node.Element) == 1,
			}
		})
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Filter_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		result := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true}
		})
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C46_LinkedList_Filter_BreakFirst(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Filter_BreakFirst", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		result := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C46_LinkedList_Filter_BreakSecond(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Filter_BreakSecond", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		result := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: arg.Index == 1}
		})
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_GetNextNodes", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c", "d"})
		nodes := ll.GetNextNodes(2)
		actual := args.Map{"result": len(nodes) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_GetAllLinkedNodes", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		all := ll.GetAllLinkedNodes()
		actual := args.Map{"result": len(all) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_RemoveNodeByIndex_First(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNodeByIndex_First", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByIndex(0)
		actual := args.Map{"result": ll.Head().Element != "b" || ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "remove first failed", actual)
	})
}

func Test_C46_LinkedList_RemoveNodeByIndex_Last(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNodeByIndex_Last", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByIndex(2)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "remove last failed", actual)
	})
}

func Test_C46_LinkedList_RemoveNodeByIndex_Middle(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNodeByIndex_Middle", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByIndex(1)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "remove middle failed", actual)
	})
}

func Test_C46_LinkedList_RemoveNodeByElementValue(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNodeByElementValue", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByElementValue("b", true, false)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_RemoveNodeByElementValue_First(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNodeByElementValue_First", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.RemoveNodeByElementValue("a", true, false)
		actual := args.Map{"result": ll.Head().Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "head should be b", actual)
	})
}

func Test_C46_LinkedList_RemoveNodeByElementValue_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNodeByElementValue_CaseInsensitive", func() {
		ll := corestr.New.LinkedList.Strings([]string{"A", "b"})
		ll.RemoveNodeByElementValue("a", false, false)
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C46_LinkedList_RemoveNode(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		nodeToRemove := ll.IndexAt(1)
		ll.RemoveNode(nodeToRemove)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_RemoveNode_Nil(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNode_Nil", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.RemoveNode(nil)
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil remove should not change", actual)
	})
}

func Test_C46_LinkedList_RemoveNode_First(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNode_First", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.RemoveNode(ll.Head())
		actual := args.Map{"result": ll.Length() != 1 || ll.Head().Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "remove first node failed", actual)
	})
}

func Test_C46_LinkedList_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNodeByIndexes", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c", "d"})
		ll.RemoveNodeByIndexes(false, 0, 2)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_RemoveNodeByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNodeByIndexes_Empty", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.RemoveNodeByIndexes(false)
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "no indexes should not change", actual)
	})
}

func Test_C46_LinkedList_ToCollection(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_ToCollection", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		col := ll.ToCollection(0)
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_ToCollection_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		col := ll.ToCollection(5)
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C46_LinkedList_List(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_List", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		list := ll.List()
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_List_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_List_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		list := ll.List()
		actual := args.Map{"result": len(list) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C46_LinkedList_ListPtr(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_ListPtr", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		list := ll.ListPtr()
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C46_LinkedList_ListLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_ListLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		list := ll.ListLock()
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C46_LinkedList_ListPtrLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_ListPtrLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		list := ll.ListPtrLock()
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C46_LinkedList_LengthLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_LengthLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		actual := args.Map{"result": ll.LengthLock() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IsEmptyLock", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"result": ll.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C46_LinkedList_String_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_String_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		s := ll.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty (NoElements)", actual)
	})
}

func Test_C46_LinkedList_String_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_String_NonEmpty", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		s := ll.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected string", actual)
	})
}

func Test_C46_LinkedList_StringLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_StringLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		s := ll.StringLock()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C46_LinkedList_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_StringLock_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		s := ll.StringLock()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected NoElements", actual)
	})
}

func Test_C46_LinkedList_Join(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Join", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		j := ll.Join(",")
		actual := args.Map{"result": j != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
	})
}

func Test_C46_LinkedList_JoinLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_JoinLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		j := ll.JoinLock(",")
		actual := args.Map{"result": j != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
	})
}

func Test_C46_LinkedList_Joins_WithItems(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Joins_WithItems", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		j := ll.Joins(",", "b", "c")
		actual := args.Map{"result": j == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C46_LinkedList_Joins_NilItems(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Joins_NilItems", func() {
		ll := corestr.New.LinkedList.Create()
		j := ll.Joins(",", "a")
		actual := args.Map{"result": j != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C46_LinkedList_IsEquals_SameRef(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IsEquals_SameRef", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		actual := args.Map{"result": ll.IsEquals(ll)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "same ref should be equal", actual)
	})
}

func Test_C46_LinkedList_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IsEquals_BothEmpty", func() {
		ll1 := corestr.New.LinkedList.Create()
		ll2 := corestr.New.LinkedList.Create()
		actual := args.Map{"result": ll1.IsEquals(ll2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "both empty should be equal", actual)
	})
}

func Test_C46_LinkedList_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IsEquals_OneEmpty", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"a"})
		ll2 := corestr.New.LinkedList.Create()
		actual := args.Map{"result": ll1.IsEquals(ll2)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be equal", actual)
	})
}

func Test_C46_LinkedList_IsEquals_DiffLength(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IsEquals_DiffLength", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"a"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		actual := args.Map{"result": ll1.IsEquals(ll2)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "different length should not be equal", actual)
	})
}

func Test_C46_LinkedList_IsEquals_Same(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IsEquals_Same", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		actual := args.Map{"result": ll1.IsEquals(ll2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "same content should be equal", actual)
	})
}

func Test_C46_LinkedList_IsEqualsWithSensitive_Nil(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IsEqualsWithSensitive_Nil", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		actual := args.Map{"result": ll.IsEqualsWithSensitive(nil, true)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil should not be equal", actual)
	})
}

func Test_C46_LinkedList_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IsEqualsWithSensitive_CaseInsensitive", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"A"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a"})
		actual := args.Map{"result": ll1.IsEqualsWithSensitive(ll2, false)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "case insensitive should be equal", actual)
	})
}

func Test_C46_LinkedList_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_GetCompareSummary", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"a"})
		ll2 := corestr.New.LinkedList.Strings([]string{"b"})
		s := ll1.GetCompareSummary(ll2, "left", "right")
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected summary", actual)
	})
}

func Test_C46_LinkedList_Clear(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Clear", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.Clear()
		actual := args.Map{"result": ll.IsEmpty() || ll.Length() != 0}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty after clear", actual)
	})
}

func Test_C46_LinkedList_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Clear_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Clear()
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C46_LinkedList_RemoveAll(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveAll", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.RemoveAll()
		actual := args.Map{"result": ll.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C46_LinkedList_AddStringsToNode_Single(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddStringsToNode_Single", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "d"})
		ll.AddStringsToNode(false, ll.Head(), []string{"b"})
		actual := args.Map{"result": ll.Length() < 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected items added", actual)
	})
}

func Test_C46_LinkedList_AddStringsToNode_Multiple(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddStringsToNode_Multiple", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "d"})
		ll.AddStringsToNode(false, ll.Head(), []string{"b", "c"})
		actual := args.Map{"result": ll.Length() < 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected items added", actual)
	})
}

func Test_C46_LinkedList_AddStringsToNode_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddStringsToNode_Empty", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.AddStringsToNode(false, ll.Head(), []string{})
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty items should not add", actual)
	})
}

func Test_C46_LinkedList_AddStringsToNode_NilNodeSkip(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddStringsToNode_NilNodeSkip", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.AddStringsToNode(true, nil, []string{"b"})
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil node with skip should not add", actual)
	})
}

func Test_C46_LinkedList_AddStringsPtrToNode_Nil(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddStringsPtrToNode_Nil", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.AddStringsPtrToNode(false, ll.Head(), nil)
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil ptr should not add", actual)
	})
}

func Test_C46_LinkedList_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddCollectionToNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "d"})
		col := corestr.New.Collection.Strings([]string{"b", "c"})
		ll.AddCollectionToNode(false, ll.Head(), col)
		actual := args.Map{"result": ll.Length() < 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected items added", actual)
	})
}

func Test_C46_LinkedList_AddAfterNode(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddAfterNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "c"})
		ll.AddAfterNode(ll.Head(), "b")
		list := ll.List()
		actual := args.Map{"result": len(list) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

// JSON

func Test_C46_LinkedList_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_MarshalJSON", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		data, err := json.Marshal(ll)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": string(data) != `["a","b"]`}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected json:", actual)
	})
}

func Test_C46_LinkedList_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_UnmarshalJSON", func() {
		ll := corestr.New.LinkedList.Create()
		err := json.Unmarshal([]byte(`["x","y"]`), ll)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": ll.Length() != 2 || ll.Head().Element != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unmarshal failed", actual)
	})
}

func Test_C46_LinkedList_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_UnmarshalJSON_Invalid", func() {
		ll := corestr.New.LinkedList.Create()
		err := json.Unmarshal([]byte(`invalid`), ll)
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C46_LinkedList_JsonModel(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_JsonModel", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		m := ll.JsonModel()
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C46_LinkedList_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_JsonModelAny", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		a := ll.JsonModelAny()
		actual := args.Map{"result": a == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C46_LinkedList_Json(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Json", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		j := ll.Json()
		actual := args.Map{"result": j.Error}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "j.Error", actual)
	})
}

func Test_C46_LinkedList_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_JsonPtr", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		j := ll.JsonPtr()
		actual := args.Map{"result": j == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C46_LinkedList_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_ParseInjectUsingJson", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		j := ll.Json()
		ll2 := corestr.New.LinkedList.Create()
		result, err := ll2.ParseInjectUsingJson(&j)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_LinkedList_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_ParseInjectUsingJson_Error", func() {
		badResult := corejson.Result{Error: json.Unmarshal([]byte("bad"), nil)}
		ll := corestr.New.LinkedList.Create()
		_, err := ll.ParseInjectUsingJson(&badResult)
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C46_LinkedList_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_ParseInjectUsingJsonMust", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		j := ll.Json()
		ll2 := corestr.New.LinkedList.Create()
		result := ll2.ParseInjectUsingJsonMust(&j)
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C46_LinkedList_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_JsonParseSelfInject", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		j := ll.Json()
		ll2 := corestr.New.LinkedList.Create()
		err := ll2.JsonParseSelfInject(&j)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C46_LinkedList_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AsJsonMarshaller", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		m := ll.AsJsonMarshaller()
		actual := args.Map{"result": m == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// Creators

func Test_C46_NewLinkedListCreator_Create(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_Create", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"result": ll.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C46_NewLinkedListCreator_Empty(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_Empty", func() {
		ll := corestr.New.LinkedList.Empty()
		actual := args.Map{"result": ll.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C46_NewLinkedListCreator_Strings(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_Strings", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_NewLinkedListCreator_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_Strings_Empty", func() {
		ll := corestr.New.LinkedList.Strings([]string{})
		actual := args.Map{"result": ll.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C46_NewLinkedListCreator_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_SpreadStrings", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_NewLinkedListCreator_SpreadStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_SpreadStrings_Empty", func() {
		ll := corestr.New.LinkedList.SpreadStrings()
		actual := args.Map{"result": ll.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C46_NewLinkedListCreator_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_PointerStringsPtr", func() {
		s1, s2 := "a", "b"
		ptrs := []*string{&s1, &s2}
		ll := corestr.New.LinkedList.PointerStringsPtr(&ptrs)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C46_NewLinkedListCreator_PointerStringsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_PointerStringsPtr_Nil", func() {
		ll := corestr.New.LinkedList.PointerStringsPtr(nil)
		actual := args.Map{"result": ll.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C46_NewLinkedListCreator_UsingMap(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_UsingMap", func() {
		m := map[string]bool{"a": true, "b": false}
		ll := corestr.New.LinkedList.UsingMap(m)
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C46_NewLinkedListCreator_UsingMap_Nil(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_UsingMap_Nil", func() {
		ll := corestr.New.LinkedList.UsingMap(nil)
		actual := args.Map{"result": ll.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// NonChainedLinkedListNodes

func Test_C46_NonChainedLinkedListNodes_Basic(t *testing.T) {
	safeTest(t, "Test_C46_NonChainedLinkedListNodes_Basic", func() {
		nc := corestr.NewNonChainedLinkedListNodes(5)
		actual := args.Map{"result": nc.IsEmpty() || nc.Length() != 0 || nc.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)

		n1 := &corestr.LinkedListNode{Element: "a"}
		n2 := &corestr.LinkedListNode{Element: "b"}
		nc.Adds(n1, n2)
		actual := args.Map{"result": nc.Length() != 2 || nc.IsEmpty() || !nc.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 items", actual)
		actual := args.Map{"result": nc.First() != n1 || nc.Last() != n2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "first/last mismatch", actual)
	})
}

func Test_C46_NonChainedLinkedListNodes_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C46_NonChainedLinkedListNodes_FirstOrDefault_Empty", func() {
		nc := corestr.NewNonChainedLinkedListNodes(0)
		actual := args.Map{"result": nc.FirstOrDefault() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C46_NonChainedLinkedListNodes_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C46_NonChainedLinkedListNodes_LastOrDefault_Empty", func() {
		nc := corestr.NewNonChainedLinkedListNodes(0)
		actual := args.Map{"result": nc.LastOrDefault() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C46_NonChainedLinkedListNodes_ApplyChaining(t *testing.T) {
	safeTest(t, "Test_C46_NonChainedLinkedListNodes_ApplyChaining", func() {
		nc := corestr.NewNonChainedLinkedListNodes(3)
		nc.Adds(
			&corestr.LinkedListNode{Element: "a"},
			&corestr.LinkedListNode{Element: "b"},
			&corestr.LinkedListNode{Element: "c"},
		)
		nc.ApplyChaining()
		actual := args.Map{"result": nc.IsChainingApplied()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "chaining should be applied", actual)
		actual := args.Map{"result": nc.First().HasNext()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "first should have next after chaining", actual)
	})
}

func Test_C46_NonChainedLinkedListNodes_ApplyChaining_Empty(t *testing.T) {
	safeTest(t, "Test_C46_NonChainedLinkedListNodes_ApplyChaining_Empty", func() {
		nc := corestr.NewNonChainedLinkedListNodes(0)
		nc.ApplyChaining()
	})
}

func Test_C46_NonChainedLinkedListNodes_ToChainedNodes(t *testing.T) {
	safeTest(t, "Test_C46_NonChainedLinkedListNodes_ToChainedNodes", func() {
		nc := corestr.NewNonChainedLinkedListNodes(2)
		nc.Adds(
			&corestr.LinkedListNode{Element: "a"},
			&corestr.LinkedListNode{Element: "b"},
		)
		chained := nc.ToChainedNodes()
		_ = chained
	})
}

func Test_C46_NonChainedLinkedListNodes_ToChainedNodes_Empty(t *testing.T) {
	safeTest(t, "Test_C46_NonChainedLinkedListNodes_ToChainedNodes_Empty", func() {
		nc := corestr.NewNonChainedLinkedListNodes(0)
		chained := nc.ToChainedNodes()
		actual := args.Map{"result": len(chained) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C46_NonChainedLinkedListNodes_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_C46_NonChainedLinkedListNodes_Adds_Nil", func() {
		nc := corestr.NewNonChainedLinkedListNodes(0)
		nc.Adds(nil)
		// nil node is still appended - just verify no panic
	})
}

// EmptyCreator linked list
func Test_C46_EmptyCreator_LinkedList(t *testing.T) {
	safeTest(t, "Test_C46_EmptyCreator_LinkedList", func() {
		ll := corestr.Empty.LinkedList()
		actual := args.Map{"result": ll.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// Concurrent test
func Test_C46_LinkedList_ConcurrentAdds(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_ConcurrentAdds", func() {
		ll := corestr.New.LinkedList.Create()
		wg := &sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				ll.AddLock("item")
			}(i)
		}
		wg.Wait()
		actual := args.Map{"result": ll.LengthLock() != 10}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
	})
}
