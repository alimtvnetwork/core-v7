package corestrtests

import (
	"encoding/json"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ===================== LinkedListNode =====================

func Test_C46_LinkedListNode_HasNext_NoNext(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_HasNext_NoNext", func() {
		node := &corestr.LinkedListNode{Element: "a"}
		if node.HasNext() {
			t.Fatal("expected no next")
		}
	})
}

func Test_C46_LinkedListNode_EndOfChain_Single(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_EndOfChain_Single", func() {
		node := &corestr.LinkedListNode{Element: "a"}
		end, length := node.EndOfChain()
		if end != node || length != 1 {
			t.Fatalf("expected self and length 1, got length %d", length)
		}
	})
}

func Test_C46_LinkedListNode_Clone(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_Clone", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		node := ll.Head()
		cloned := node.Clone()
		if cloned.Element != "a" || cloned.HasNext() {
			t.Fatal("clone should copy element but not next")
		}
	})
}

func Test_C46_LinkedListNode_List(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_List", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		list := ll.Head().List()
		if len(list) != 3 || list[2] != "c" {
			t.Fatalf("unexpected list: %v", list)
		}
	})
}
func Test_C46_LinkedListNode_Join(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_Join", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		j := ll.Head().Join(",")
		if j != "a,b" {
			t.Fatalf("got %s", j)
		}
	})
}

func Test_C46_LinkedListNode_StringList(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_StringList", func() {
		ll := corestr.New.LinkedList.Strings([]string{"x"})
		s := ll.Head().StringList("H:")
		if s == "" {
			t.Fatal("expected non-empty")
		}
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
		if node.String() != "hello" {
			t.Fatal("expected hello")
		}
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
		if !node.IsEqual(node) {
			t.Fatal("same node should be equal")
		}
	})
}

func Test_C46_LinkedListNode_IsEqual_DifferentValues(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_IsEqual_DifferentValues", func() {
		n1 := &corestr.LinkedListNode{Element: "a"}
		n2 := &corestr.LinkedListNode{Element: "b"}
		if n1.IsEqual(n2) {
			t.Fatal("different elements should not be equal")
		}
	})
}

func Test_C46_LinkedListNode_IsEqualSensitive(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_IsEqualSensitive", func() {
		n1 := &corestr.LinkedListNode{Element: "Hello"}
		n2 := &corestr.LinkedListNode{Element: "hello"}
		if n1.IsEqualSensitive(n2, true) {
			t.Fatal("case sensitive should fail")
		}
		if !n1.IsEqualSensitive(n2, false) {
			t.Fatal("case insensitive should pass")
		}
	})
}

func Test_C46_LinkedListNode_IsChainEqual(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_IsChainEqual", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		if !ll1.Head().IsChainEqual(ll2.Head(), true) {
			t.Fatal("chains should be equal")
		}
	})
}

func Test_C46_LinkedListNode_IsChainEqual_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_IsChainEqual_CaseInsensitive", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"A", "B"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		if !ll1.Head().IsChainEqual(ll2.Head(), false) {
			t.Fatal("chains should be equal case-insensitive")
		}
	})
}

func Test_C46_LinkedListNode_IsEqualValue(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_IsEqualValue", func() {
		n := &corestr.LinkedListNode{Element: "test"}
		if !n.IsEqualValue("test") {
			t.Fatal("should match")
		}
		if n.IsEqualValue("other") {
			t.Fatal("should not match")
		}
	})
}

func Test_C46_LinkedListNode_IsEqualValueSensitive(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_IsEqualValueSensitive", func() {
		n := &corestr.LinkedListNode{Element: "Test"}
		if !n.IsEqualValueSensitive("test", false) {
			t.Fatal("case-insensitive should match")
		}
		if n.IsEqualValueSensitive("test", true) {
			t.Fatal("case-sensitive should not match")
		}
	})
}

func Test_C46_LinkedListNode_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_CreateLinkedList", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		newLL := ll.Head().CreateLinkedList()
		if newLL.Length() != 3 {
			t.Fatalf("expected 3, got %d", newLL.Length())
		}
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
		if count != 3 || length != 3 || end.Element != "c" {
			t.Fatalf("unexpected count=%d length=%d end=%s", count, length, end.Element)
		}
	})
}

func Test_C46_LinkedListNode_LoopEndOfChain_BreakFirst(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_LoopEndOfChain_BreakFirst", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		end, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			return true // break immediately
		})
		if length != 1 || end.Element != "a" {
			t.Fatalf("expected break at first, length=%d end=%s", length, end.Element)
		}
	})
}

func Test_C46_LinkedListNode_AddNext(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_AddNext", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "c"})
		node := ll.Head()
		node.AddNext(ll, "b")
		if ll.Length() != 3 {
			t.Fatalf("expected 3, got %d", ll.Length())
		}
	})
}

func Test_C46_LinkedListNode_AddNextNode(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_AddNextNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "c"})
		node := ll.Head()
		newNode := &corestr.LinkedListNode{Element: "b"}
		node.AddNextNode(ll, newNode)
		if ll.Length() != 3 {
			t.Fatalf("expected 3, got %d", ll.Length())
		}
	})
}

func Test_C46_LinkedListNode_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_AddStringsToNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "d"})
		node := ll.Head()
		node.AddStringsToNode(ll, false, []string{"b", "c"})
		if ll.Length() < 3 {
			t.Fatalf("expected items added, got %d", ll.Length())
		}
	})
}
func Test_C46_LinkedListNode_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_C46_LinkedListNode_AddCollectionToNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "d"})
		node := ll.Head()
		col := corestr.New.Collection.Strings([]string{"b", "c"})
		node.AddCollectionToNode(ll, false, col)
		if ll.Length() < 3 {
			t.Fatal("expected items added")
		}
	})
}

// ===================== LinkedList =====================

func Test_C46_LinkedList_Create_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Create_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		if !ll.IsEmpty() || ll.Length() != 0 {
			t.Fatal("expected empty")
		}
		if ll.HasItems() {
			t.Fatal("should not have items")
		}
	})
}

func Test_C46_LinkedList_Add_Single(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Add_Single", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("hello")
		if ll.Length() != 1 || ll.Head().Element != "hello" || ll.Tail().Element != "hello" {
			t.Fatal("single add failed")
		}
	})
}

func Test_C46_LinkedList_Add_Multiple(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Add_Multiple", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a").Add("b").Add("c")
		if ll.Length() != 3 || ll.Head().Element != "a" || ll.Tail().Element != "c" {
			t.Fatal("multi add failed")
		}
	})
}

func Test_C46_LinkedList_Adds(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Adds", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		if ll.Length() != 3 {
			t.Fatal("adds failed")
		}
	})
}

func Test_C46_LinkedList_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Adds_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds()
		if ll.Length() != 0 {
			t.Fatal("expected empty")
		}
	})
}

func Test_C46_LinkedList_AddStrings(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddStrings", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings([]string{"x", "y"})
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C46_LinkedList_AddLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddLock("safe")
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C46_LinkedList_AddsLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddsLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddsLock("a", "b")
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C46_LinkedList_AddFront_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddFront_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFront("first")
		if ll.Length() != 1 || ll.Head().Element != "first" {
			t.Fatal("add front to empty failed")
		}
	})
}

func Test_C46_LinkedList_AddFront_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddFront_NonEmpty", func() {
		ll := corestr.New.LinkedList.Strings([]string{"b", "c"})
		ll.AddFront("a")
		if ll.Head().Element != "a" || ll.Length() != 3 {
			t.Fatal("add front failed")
		}
	})
}

func Test_C46_LinkedList_PushFront(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_PushFront", func() {
		ll := corestr.New.LinkedList.Strings([]string{"b"})
		ll.PushFront("a")
		if ll.Head().Element != "a" {
			t.Fatal("push front failed")
		}
	})
}

func Test_C46_LinkedList_Push(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Push", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Push("x")
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C46_LinkedList_PushBack(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_PushBack", func() {
		ll := corestr.New.LinkedList.Create()
		ll.PushBack("x")
		if ll.Tail().Element != "x" {
			t.Fatal("push back failed")
		}
	})
}

func Test_C46_LinkedList_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddNonEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmpty("")
		if ll.Length() != 0 {
			t.Fatal("empty string should not be added")
		}
		ll.AddNonEmpty("x")
		if ll.Length() != 1 {
			t.Fatal("non-empty should be added")
		}
	})
}

func Test_C46_LinkedList_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddNonEmptyWhitespace", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmptyWhitespace("   ")
		if ll.Length() != 0 {
			t.Fatal("whitespace should not be added")
		}
		ll.AddNonEmptyWhitespace("x")
		if ll.Length() != 1 {
			t.Fatal("non-whitespace should be added")
		}
	})
}

func Test_C46_LinkedList_AddIf(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddIf", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddIf(false, "skip")
		if ll.Length() != 0 {
			t.Fatal("should skip")
		}
		ll.AddIf(true, "add")
		if ll.Length() != 1 {
			t.Fatal("should add")
		}
	})
}

func Test_C46_LinkedList_AddsIf(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddsIf", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddsIf(false, "a", "b")
		if ll.Length() != 0 {
			t.Fatal("should skip")
		}
		ll.AddsIf(true, "a", "b")
		if ll.Length() != 2 {
			t.Fatal("should add")
		}
	})
}

func Test_C46_LinkedList_AddFunc(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddFunc", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFunc(func() string { return "generated" })
		if ll.Length() != 1 || ll.Head().Element != "generated" {
			t.Fatal("add func failed")
		}
	})
}

func Test_C46_LinkedList_AddFuncErr_Success(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddFuncErr_Success", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) { t.Fatal("no err expected") })
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
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
		if !called {
			t.Fatal("error handler should be called")
		}
	})
}

func Test_C46_LinkedList_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddItemsMap", func() {
		ll := corestr.New.LinkedList.Create()
		m := map[string]bool{"a": true, "b": false, "c": true}
		ll.AddItemsMap(m)
		if ll.Length() != 2 {
			t.Fatalf("expected 2, got %d", ll.Length())
		}
	})
}

func Test_C46_LinkedList_AddItemsMap_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddItemsMap_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{})
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C46_LinkedList_AppendNode_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AppendNode_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AppendNode(&corestr.LinkedListNode{Element: "x"})
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C46_LinkedList_AppendNode_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AppendNode_NonEmpty", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.AppendNode(&corestr.LinkedListNode{Element: "b"})
		if ll.Length() != 2 || ll.Tail().Element != "b" {
			t.Fatal("append failed")
		}
	})
}

func Test_C46_LinkedList_AddBackNode(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddBackNode", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddBackNode(&corestr.LinkedListNode{Element: "x"})
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C46_LinkedList_AppendChainOfNodes_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AppendChainOfNodes_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		chain := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.AppendChainOfNodes(chain.Head())
		if ll.Length() != 2 {
			t.Fatalf("expected 2, got %d", ll.Length())
		}
	})
}

func Test_C46_LinkedList_AppendChainOfNodes_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AppendChainOfNodes_NonEmpty", func() {
		ll := corestr.New.LinkedList.Strings([]string{"x"})
		chain := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.AppendChainOfNodes(chain.Head())
		if ll.Length() != 3 {
			t.Fatalf("expected 3, got %d", ll.Length())
		}
	})
}

func Test_C46_LinkedList_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_InsertAt_Front", func() {
		ll := corestr.New.LinkedList.Strings([]string{"b", "c"})
		ll.InsertAt(0, "a")
		if ll.Head().Element != "a" {
			t.Fatal("insert at front failed")
		}
	})
}

func Test_C46_LinkedList_InsertAt_Middle(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_InsertAt_Middle", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "c"})
		ll.InsertAt(1, "b")
		list := ll.List()
		if len(list) < 3 {
			t.Fatalf("expected 3+, got %d", len(list))
		}
	})
}

func Test_C46_LinkedList_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AttachWithNode_NilCurrent", func() {
		ll := corestr.New.LinkedList.Create()
		err := ll.AttachWithNode(nil, &corestr.LinkedListNode{Element: "x"})
		if err == nil {
			t.Fatal("expected error for nil current node")
		}
	})
}

func Test_C46_LinkedList_AttachWithNode_NextNotNil(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AttachWithNode_NextNotNil", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		err := ll.AttachWithNode(ll.Head(), &corestr.LinkedListNode{Element: "x"})
		if err == nil {
			t.Fatal("expected error for current.next not nil")
		}
	})
}

func Test_C46_LinkedList_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddCollection_Nil", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddCollection(nil)
		if ll.Length() != 0 {
			t.Fatal("nil collection should not add")
		}
	})
}

func Test_C46_LinkedList_AddCollection(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddCollection", func() {
		ll := corestr.New.LinkedList.Create()
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		ll.AddCollection(col)
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C46_LinkedList_AddPointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddPointerStringsPtr", func() {
		ll := corestr.New.LinkedList.Create()
		s1 := "a"
		s2 := "b"
		ll.AddPointerStringsPtr([]*string{&s1, nil, &s2})
		if ll.Length() != 2 {
			t.Fatalf("expected 2 (skip nil), got %d", ll.Length())
		}
	})
}

func Test_C46_LinkedList_IndexAt(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IndexAt", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		n := ll.IndexAt(1)
		if n.Element != "b" {
			t.Fatal("expected b")
		}
		n0 := ll.IndexAt(0)
		if n0.Element != "a" {
			t.Fatal("expected a")
		}
	})
}

func Test_C46_LinkedList_IndexAt_Negative(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IndexAt_Negative", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		n := ll.IndexAt(-1)
		if n != nil {
			t.Fatal("negative index should return nil")
		}
	})
}

func Test_C46_LinkedList_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_SafeIndexAt", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		n := ll.SafeIndexAt(1)
		if n == nil || n.Element != "b" {
			t.Fatal("expected b")
		}
		n2 := ll.SafeIndexAt(5)
		if n2 != nil {
			t.Fatal("out of range should return nil")
		}
		n3 := ll.SafeIndexAt(-1)
		if n3 != nil {
			t.Fatal("negative should return nil")
		}
	})
}

func Test_C46_LinkedList_SafeIndexAtLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_SafeIndexAtLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		n := ll.SafeIndexAtLock(0)
		if n == nil || n.Element != "a" {
			t.Fatal("expected a")
		}
	})
}

func Test_C46_LinkedList_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_SafePointerIndexAt", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		p := ll.SafePointerIndexAt(0)
		if p == nil || *p != "a" {
			t.Fatal("expected a")
		}
		p2 := ll.SafePointerIndexAt(99)
		if p2 != nil {
			t.Fatal("out of range should return nil")
		}
	})
}

func Test_C46_LinkedList_SafePointerIndexAtUsingDefault(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_SafePointerIndexAtUsingDefault", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		v := ll.SafePointerIndexAtUsingDefault(0, "def")
		if v != "a" {
			t.Fatal("expected a")
		}
		v2 := ll.SafePointerIndexAtUsingDefault(99, "def")
		if v2 != "def" {
			t.Fatal("expected default")
		}
	})
}

func Test_C46_LinkedList_SafePointerIndexAtUsingDefaultLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_SafePointerIndexAtUsingDefaultLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		v := ll.SafePointerIndexAtUsingDefaultLock(0, "def")
		if v != "a" {
			t.Fatal("expected a")
		}
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
		if len(elements) != 3 {
			t.Fatalf("expected 3, got %d", len(elements))
		}
	})
}

func Test_C46_LinkedList_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Loop_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			t.Fatal("should not be called")
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
		if count != 1 {
			t.Fatalf("expected 1 iteration, got %d", count)
		}
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
		if count != 2 {
			t.Fatalf("expected 2, got %d", count)
		}
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
		if len(result) != 2 {
			t.Fatalf("expected 2, got %d", len(result))
		}
	})
}

func Test_C46_LinkedList_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Filter_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		result := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true}
		})
		if len(result) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C46_LinkedList_Filter_BreakFirst(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Filter_BreakFirst", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		result := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})
		if len(result) != 1 {
			t.Fatalf("expected 1, got %d", len(result))
		}
	})
}

func Test_C46_LinkedList_Filter_BreakSecond(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Filter_BreakSecond", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		result := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: arg.Index == 1}
		})
		if len(result) != 2 {
			t.Fatalf("expected 2, got %d", len(result))
		}
	})
}

func Test_C46_LinkedList_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_GetNextNodes", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c", "d"})
		nodes := ll.GetNextNodes(2)
		if len(nodes) != 2 {
			t.Fatalf("expected 2, got %d", len(nodes))
		}
	})
}

func Test_C46_LinkedList_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_GetAllLinkedNodes", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		all := ll.GetAllLinkedNodes()
		if len(all) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C46_LinkedList_RemoveNodeByIndex_First(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNodeByIndex_First", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByIndex(0)
		if ll.Head().Element != "b" || ll.Length() != 2 {
			t.Fatal("remove first failed")
		}
	})
}

func Test_C46_LinkedList_RemoveNodeByIndex_Last(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNodeByIndex_Last", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByIndex(2)
		if ll.Length() != 2 {
			t.Fatal("remove last failed")
		}
	})
}

func Test_C46_LinkedList_RemoveNodeByIndex_Middle(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNodeByIndex_Middle", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByIndex(1)
		if ll.Length() != 2 {
			t.Fatal("remove middle failed")
		}
	})
}

func Test_C46_LinkedList_RemoveNodeByElementValue(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNodeByElementValue", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByElementValue("b", true, false)
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C46_LinkedList_RemoveNodeByElementValue_First(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNodeByElementValue_First", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.RemoveNodeByElementValue("a", true, false)
		if ll.Head().Element != "b" {
			t.Fatal("head should be b")
		}
	})
}

func Test_C46_LinkedList_RemoveNodeByElementValue_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNodeByElementValue_CaseInsensitive", func() {
		ll := corestr.New.LinkedList.Strings([]string{"A", "b"})
		ll.RemoveNodeByElementValue("a", false, false)
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C46_LinkedList_RemoveNode(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		nodeToRemove := ll.IndexAt(1)
		ll.RemoveNode(nodeToRemove)
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C46_LinkedList_RemoveNode_Nil(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNode_Nil", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.RemoveNode(nil)
		if ll.Length() != 1 {
			t.Fatal("nil remove should not change")
		}
	})
}

func Test_C46_LinkedList_RemoveNode_First(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNode_First", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.RemoveNode(ll.Head())
		if ll.Length() != 1 || ll.Head().Element != "b" {
			t.Fatal("remove first node failed")
		}
	})
}

func Test_C46_LinkedList_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNodeByIndexes", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c", "d"})
		ll.RemoveNodeByIndexes(false, 0, 2)
		if ll.Length() != 2 {
			t.Fatalf("expected 2, got %d", ll.Length())
		}
	})
}

func Test_C46_LinkedList_RemoveNodeByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveNodeByIndexes_Empty", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.RemoveNodeByIndexes(false)
		if ll.Length() != 1 {
			t.Fatal("no indexes should not change")
		}
	})
}

func Test_C46_LinkedList_ToCollection(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_ToCollection", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		col := ll.ToCollection(0)
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C46_LinkedList_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_ToCollection_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		col := ll.ToCollection(5)
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C46_LinkedList_List(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_List", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		list := ll.List()
		if len(list) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C46_LinkedList_List_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_List_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		list := ll.List()
		if len(list) != 0 {
			t.Fatal("expected 0")
		}
	})
}
func Test_C46_LinkedList_ListLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_ListLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		list := ll.ListLock()
		if len(list) != 1 {
			t.Fatal("expected 1")
		}
	})
}
func Test_C46_LinkedList_LengthLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_LengthLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		if ll.LengthLock() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C46_LinkedList_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IsEmptyLock", func() {
		ll := corestr.New.LinkedList.Create()
		if !ll.IsEmptyLock() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C46_LinkedList_String_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_String_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		s := ll.String()
		if s == "" {
			t.Fatal("expected non-empty (NoElements)")
		}
	})
}

func Test_C46_LinkedList_String_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_String_NonEmpty", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		s := ll.String()
		if s == "" {
			t.Fatal("expected string")
		}
	})
}

func Test_C46_LinkedList_StringLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_StringLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		s := ll.StringLock()
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_C46_LinkedList_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_StringLock_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		s := ll.StringLock()
		if s == "" {
			t.Fatal("expected NoElements")
		}
	})
}

func Test_C46_LinkedList_Join(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Join", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		j := ll.Join(",")
		if j != "a,b" {
			t.Fatalf("expected a,b got %s", j)
		}
	})
}

func Test_C46_LinkedList_JoinLock(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_JoinLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		j := ll.JoinLock(",")
		if j != "a,b" {
			t.Fatalf("expected a,b got %s", j)
		}
	})
}

func Test_C46_LinkedList_Joins_WithItems(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Joins_WithItems", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		j := ll.Joins(",", "b", "c")
		if j == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_C46_LinkedList_Joins_NilItems(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Joins_NilItems", func() {
		ll := corestr.New.LinkedList.Create()
		j := ll.Joins(",", "a")
		if j != "a" {
			t.Fatalf("expected a, got %s", j)
		}
	})
}

func Test_C46_LinkedList_IsEquals_SameRef(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IsEquals_SameRef", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		if !ll.IsEquals(ll) {
			t.Fatal("same ref should be equal")
		}
	})
}

func Test_C46_LinkedList_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IsEquals_BothEmpty", func() {
		ll1 := corestr.New.LinkedList.Create()
		ll2 := corestr.New.LinkedList.Create()
		if !ll1.IsEquals(ll2) {
			t.Fatal("both empty should be equal")
		}
	})
}

func Test_C46_LinkedList_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IsEquals_OneEmpty", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"a"})
		ll2 := corestr.New.LinkedList.Create()
		if ll1.IsEquals(ll2) {
			t.Fatal("should not be equal")
		}
	})
}

func Test_C46_LinkedList_IsEquals_DiffLength(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IsEquals_DiffLength", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"a"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		if ll1.IsEquals(ll2) {
			t.Fatal("different length should not be equal")
		}
	})
}

func Test_C46_LinkedList_IsEquals_Same(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IsEquals_Same", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		if !ll1.IsEquals(ll2) {
			t.Fatal("same content should be equal")
		}
	})
}

func Test_C46_LinkedList_IsEqualsWithSensitive_Nil(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IsEqualsWithSensitive_Nil", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		if ll.IsEqualsWithSensitive(nil, true) {
			t.Fatal("nil should not be equal")
		}
	})
}

func Test_C46_LinkedList_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_IsEqualsWithSensitive_CaseInsensitive", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"A"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a"})
		if !ll1.IsEqualsWithSensitive(ll2, false) {
			t.Fatal("case insensitive should be equal")
		}
	})
}

func Test_C46_LinkedList_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_GetCompareSummary", func() {
		ll1 := corestr.New.LinkedList.Strings([]string{"a"})
		ll2 := corestr.New.LinkedList.Strings([]string{"b"})
		s := ll1.GetCompareSummary(ll2, "left", "right")
		if s == "" {
			t.Fatal("expected summary")
		}
	})
}

func Test_C46_LinkedList_Clear(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Clear", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.Clear()
		if !ll.IsEmpty() || ll.Length() != 0 {
			t.Fatal("expected empty after clear")
		}
	})
}

func Test_C46_LinkedList_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Clear_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Clear()
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C46_LinkedList_RemoveAll(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_RemoveAll", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.RemoveAll()
		if !ll.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C46_LinkedList_AddStringsToNode_Single(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddStringsToNode_Single", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "d"})
		ll.AddStringsToNode(false, ll.Head(), []string{"b"})
		if ll.Length() < 3 {
			t.Fatal("expected items added")
		}
	})
}

func Test_C46_LinkedList_AddStringsToNode_Multiple(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddStringsToNode_Multiple", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "d"})
		ll.AddStringsToNode(false, ll.Head(), []string{"b", "c"})
		if ll.Length() < 4 {
			t.Fatal("expected items added")
		}
	})
}

func Test_C46_LinkedList_AddStringsToNode_Empty(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddStringsToNode_Empty", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.AddStringsToNode(false, ll.Head(), []string{})
		if ll.Length() != 1 {
			t.Fatal("empty items should not add")
		}
	})
}

func Test_C46_LinkedList_AddStringsToNode_NilNodeSkip(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddStringsToNode_NilNodeSkip", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.AddStringsToNode(true, nil, []string{"b"})
		if ll.Length() != 1 {
			t.Fatal("nil node with skip should not add")
		}
	})
}
func Test_C46_LinkedList_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddCollectionToNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "d"})
		col := corestr.New.Collection.Strings([]string{"b", "c"})
		ll.AddCollectionToNode(false, ll.Head(), col)
		if ll.Length() < 3 {
			t.Fatal("expected items added")
		}
	})
}

func Test_C46_LinkedList_AddAfterNode(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AddAfterNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "c"})
		ll.AddAfterNode(ll.Head(), "b")
		list := ll.List()
		if len(list) != 3 {
			t.Fatalf("expected 3, got %d", len(list))
		}
	})
}

// JSON

func Test_C46_LinkedList_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_MarshalJSON", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		data, err := json.Marshal(ll)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != `["a","b"]` {
			t.Fatalf("unexpected json: %s", string(data))
		}
	})
}

func Test_C46_LinkedList_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_UnmarshalJSON", func() {
		ll := corestr.New.LinkedList.Create()
		err := json.Unmarshal([]byte(`["x","y"]`), ll)
		if err != nil {
			t.Fatal(err)
		}
		if ll.Length() != 2 || ll.Head().Element != "x" {
			t.Fatal("unmarshal failed")
		}
	})
}

func Test_C46_LinkedList_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_UnmarshalJSON_Invalid", func() {
		ll := corestr.New.LinkedList.Create()
		err := json.Unmarshal([]byte(`invalid`), ll)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_C46_LinkedList_JsonModel(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_JsonModel", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		m := ll.JsonModel()
		if len(m) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C46_LinkedList_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_JsonModelAny", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		a := ll.JsonModelAny()
		if a == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_C46_LinkedList_Json(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_Json", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		j := ll.Json()
		if j.Error != nil {
			t.Fatal(j.Error)
		}
	})
}

func Test_C46_LinkedList_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_JsonPtr", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		j := ll.JsonPtr()
		if j == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_C46_LinkedList_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_ParseInjectUsingJson", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		j := ll.Json()
		ll2 := corestr.New.LinkedList.Create()
		result, err := ll2.ParseInjectUsingJson(&j)
		if err != nil {
			t.Fatal(err)
		}
		if result.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C46_LinkedList_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_ParseInjectUsingJson_Error", func() {
		badResult := corejson.Result{Error: json.Unmarshal([]byte("bad"), nil)}
		ll := corestr.New.LinkedList.Create()
		_, err := ll.ParseInjectUsingJson(&badResult)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_C46_LinkedList_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_ParseInjectUsingJsonMust", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		j := ll.Json()
		ll2 := corestr.New.LinkedList.Create()
		result := ll2.ParseInjectUsingJsonMust(&j)
		if result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C46_LinkedList_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_JsonParseSelfInject", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		j := ll.Json()
		ll2 := corestr.New.LinkedList.Create()
		err := ll2.JsonParseSelfInject(&j)
		if err != nil {
			t.Fatal(err)
		}
	})
}

func Test_C46_LinkedList_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_C46_LinkedList_AsJsonMarshaller", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		m := ll.AsJsonMarshaller()
		if m == nil {
			t.Fatal("expected non-nil")
		}
	})
}

// Creators

func Test_C46_NewLinkedListCreator_Create(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_Create", func() {
		ll := corestr.New.LinkedList.Create()
		if !ll.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C46_NewLinkedListCreator_Empty(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_Empty", func() {
		ll := corestr.New.LinkedList.Empty()
		if !ll.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C46_NewLinkedListCreator_Strings(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_Strings", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C46_NewLinkedListCreator_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_Strings_Empty", func() {
		ll := corestr.New.LinkedList.Strings([]string{})
		if !ll.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C46_NewLinkedListCreator_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_SpreadStrings", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C46_NewLinkedListCreator_SpreadStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_SpreadStrings_Empty", func() {
		ll := corestr.New.LinkedList.SpreadStrings()
		if !ll.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C46_NewLinkedListCreator_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_PointerStringsPtr", func() {
		s1, s2 := "a", "b"
		ptrs := []*string{&s1, &s2}
		ll := corestr.New.LinkedList.PointerStringsPtr(&ptrs)
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C46_NewLinkedListCreator_PointerStringsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_PointerStringsPtr_Nil", func() {
		ll := corestr.New.LinkedList.PointerStringsPtr(nil)
		if !ll.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C46_NewLinkedListCreator_UsingMap(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_UsingMap", func() {
		m := map[string]bool{"a": true, "b": false}
		ll := corestr.New.LinkedList.UsingMap(m)
		if ll.Length() != 1 {
			t.Fatalf("expected 1, got %d", ll.Length())
		}
	})
}

func Test_C46_NewLinkedListCreator_UsingMap_Nil(t *testing.T) {
	safeTest(t, "Test_C46_NewLinkedListCreator_UsingMap_Nil", func() {
		ll := corestr.New.LinkedList.UsingMap(nil)
		if !ll.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

// NonChainedLinkedListNodes

func Test_C46_NonChainedLinkedListNodes_Basic(t *testing.T) {
	safeTest(t, "Test_C46_NonChainedLinkedListNodes_Basic", func() {
		nc := corestr.NewNonChainedLinkedListNodes(5)
		if !nc.IsEmpty() || nc.Length() != 0 || nc.HasItems() {
			t.Fatal("expected empty")
		}

		n1 := &corestr.LinkedListNode{Element: "a"}
		n2 := &corestr.LinkedListNode{Element: "b"}
		nc.Adds(n1, n2)
		if nc.Length() != 2 || nc.IsEmpty() || !nc.HasItems() {
			t.Fatal("expected 2 items")
		}
		if nc.First() != n1 || nc.Last() != n2 {
			t.Fatal("first/last mismatch")
		}
	})
}

func Test_C46_NonChainedLinkedListNodes_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C46_NonChainedLinkedListNodes_FirstOrDefault_Empty", func() {
		nc := corestr.NewNonChainedLinkedListNodes(0)
		if nc.FirstOrDefault() != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_C46_NonChainedLinkedListNodes_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C46_NonChainedLinkedListNodes_LastOrDefault_Empty", func() {
		nc := corestr.NewNonChainedLinkedListNodes(0)
		if nc.LastOrDefault() != nil {
			t.Fatal("expected nil")
		}
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
		if !nc.IsChainingApplied() {
			t.Fatal("chaining should be applied")
		}
		if !nc.First().HasNext() {
			t.Fatal("first should have next after chaining")
		}
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
		if len(chained) != 0 {
			t.Fatal("expected 0")
		}
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
		if !ll.IsEmpty() {
			t.Fatal("expected empty")
		}
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
		if ll.LengthLock() != 10 {
			t.Fatalf("expected 10, got %d", ll.LengthLock())
		}
	})
}
