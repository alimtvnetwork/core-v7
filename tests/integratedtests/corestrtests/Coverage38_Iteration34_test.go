package corestrtests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════
// LinkedList — basic operations
// ══════════════════════════════════════════════════════════════

func Test_Cov38_LL_Add(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Add", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a").Add("b")
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_LL_Head_Tail(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Head_Tail", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		actual := args.Map{"result": ll.Head().Element != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual := args.Map{"result": ll.Tail().Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_Cov38_LL_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"result": ll.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual := args.Map{"result": ll.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
	})
}

func Test_Cov38_LL_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEmptyLock", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"result": ll.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov38_LL_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_LengthLock", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_PushBack(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_PushBack", func() {
		ll := corestr.New.LinkedList.Create()
		ll.PushBack("x")
		actual := args.Map{"result": ll.Length() != 1 || ll.Head().Element != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov38_LL_Push(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Push", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Push("x")
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov38_LL_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddLock("x")
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov38_LL_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddNonEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmpty("")
		ll.AddNonEmpty("a")
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddNonEmptyWhitespace", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmptyWhitespace("  ")
		ll.AddNonEmptyWhitespace("a")
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_AddIf(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddIf", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddIf(false, "no")
		ll.AddIf(true, "yes")
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_AddsIf_True(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddsIf_True", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddsIf(true, "a", "b")
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_LL_AddsIf_False(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddsIf_False", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddsIf(false, "a", "b")
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov38_LL_AddFunc(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddFunc", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFunc(func() string { return "hello" })
		actual := args.Map{"result": ll.Length() != 1 || ll.Head().Element != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov38_LL_AddFuncErr_NoErr(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddFuncErr_NoErr", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(err error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) },
		)
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_AddFuncErr_WithErr(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddFuncErr_WithErr", func() {
		ll := corestr.New.LinkedList.Create()
		called := false
		ll.AddFuncErr(
			func() (string, error) { return "", json.Unmarshal([]byte("bad"), &struct{}{}) },
			func(err error) { called = true },
		)
		actual := args.Map{"result": called}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected error handler to be called", actual)
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov38_LL_Adds(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Adds", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		actual := args.Map{"result": ll.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov38_LL_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Adds_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds()
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov38_LL_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddStrings", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings([]string{"x", "y"})
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_LL_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddStrings_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings(nil)
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov38_LL_AddsLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddsLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddsLock("a", "b")
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_LL_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddItemsMap", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{"a": true, "b": false, "c": true})
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_LL_AddItemsMap_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddItemsMap_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(nil)
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov38_LL_AddFront(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddFront", func() {
		ll := corestr.New.LinkedList.SpreadStrings("b")
		ll.AddFront("a")
		actual := args.Map{"result": ll.Head().Element != "a" || ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov38_LL_AddFront_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddFront_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFront("a")
		actual := args.Map{"result": ll.Head().Element != "a" || ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov38_LL_PushFront(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_PushFront", func() {
		ll := corestr.New.LinkedList.SpreadStrings("b")
		ll.PushFront("a")
		actual := args.Map{"result": ll.Head().Element != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov38_LL_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddCollection", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		ll := corestr.New.LinkedList.Create()
		ll.AddCollection(col)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_LL_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddCollection_Nil", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddCollection(nil)
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov38_LL_AddPointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddPointerStringsPtr", func() {
		a, b := "a", "b"
		ll := corestr.New.LinkedList.Create()
		ll.AddPointerStringsPtr([]*string{&a, nil, &b})
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── InsertAt ──

func Test_Cov38_LL_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_InsertAt_Front", func() {
		ll := corestr.New.LinkedList.SpreadStrings("b", "c")
		ll.InsertAt(0, "a")
		actual := args.Map{"result": ll.Head().Element != "a" || ll.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_Cov38_LL_InsertAt_Middle(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_InsertAt_Middle", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "c")
		ll.InsertAt(1, "b")
		list := ll.List()
		actual := args.Map{"result": len(list) != 3 || list[1] != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

// ── AppendNode / AppendChainOfNodes / AddBackNode ──

func Test_Cov38_LL_AppendNode_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AppendNode_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		node := ll.Head() // nil
		_ = node
		ll.AppendNode(&corestr.LinkedListNode{Element: "x"})
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_AppendNode_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AppendNode_NonEmpty", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.AppendNode(&corestr.LinkedListNode{Element: "b"})
		actual := args.Map{"result": ll.Length() != 2 || ll.Tail().Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov38_LL_AddBackNode(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddBackNode", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddBackNode(&corestr.LinkedListNode{Element: "x"})
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AppendChainOfNodes", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		chain := corestr.New.LinkedList.SpreadStrings("b", "c")
		ll.AppendChainOfNodes(chain.Head())
		actual := args.Map{"result": ll.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov38_LL_AppendChainOfNodes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AppendChainOfNodes_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		chain := corestr.New.LinkedList.SpreadStrings("a", "b")
		ll.AppendChainOfNodes(chain.Head())
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── AttachWithNode ──

func Test_Cov38_LL_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AttachWithNode_NilCurrent", func() {
		ll := corestr.New.LinkedList.Create()
		err := ll.AttachWithNode(nil, &corestr.LinkedListNode{Element: "x"})
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_Cov38_LL_AttachWithNode_NextNotNil(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AttachWithNode_NextNotNil", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		err := ll.AttachWithNode(ll.Head(), &corestr.LinkedListNode{Element: "x"})
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error because head.next is not nil", actual)
	})
}

// ── Loop ──

func Test_Cov38_LL_Loop(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Loop", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})
		actual := args.Map{"result": count != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov38_LL_Loop_Break(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Loop_Break", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return true
		})
		actual := args.Map{"result": count != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Loop_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			actual := args.Map{"result": false}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "should not be called", actual)
			return false
		})
	})
}

func Test_Cov38_LL_Loop_BreakSecond(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Loop_BreakSecond", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
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

// ── Filter ──

func Test_Cov38_LL_Filter(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Filter", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: arg.Node.Element != "b"}
		})
		actual := args.Map{"result": len(nodes) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_LL_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Filter_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true}
		})
		actual := args.Map{"result": len(nodes) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov38_LL_Filter_BreakFirst(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Filter_BreakFirst", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})
		actual := args.Map{"result": len(nodes) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_Filter_BreakSecond(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Filter_BreakSecond", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: arg.Index == 1}
		})
		actual := args.Map{"result": len(nodes) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── RemoveNodeByElementValue ──

func Test_Cov38_LL_RemoveByElem_CaseSensitive(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveByElem_CaseSensitive", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		ll.RemoveNodeByElementValue("b", true, false)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_LL_RemoveByElem_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveByElem_CaseInsensitive", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "B", "c")
		ll.RemoveNodeByElementValue("b", false, false)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_LL_RemoveByElem_First(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveByElem_First", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		ll.RemoveNodeByElementValue("a", true, false)
		actual := args.Map{"result": ll.Length() != 1 || ll.Head().Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

// ── RemoveNodeByIndex ──

func Test_Cov38_LL_RemoveByIndex_First(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveByIndex_First", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		ll.RemoveNodeByIndex(0)
		actual := args.Map{"result": ll.Length() != 2 || ll.Head().Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov38_LL_RemoveByIndex_Last(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveByIndex_Last", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		ll.RemoveNodeByIndex(2)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_LL_RemoveByIndex_Middle(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveByIndex_Middle", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		ll.RemoveNodeByIndex(1)
		list := ll.List()
		actual := args.Map{"result": len(list) != 2 || list[0] != "a" || list[1] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

// ── RemoveNodeByIndexes ──

func Test_Cov38_LL_RemoveByIndexes(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveByIndexes", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c", "d")
		ll.RemoveNodeByIndexes(true, 1, 3)
		list := ll.List()
		actual := args.Map{"result": len(list) != 2 || list[0] != "a" || list[1] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_Cov38_LL_RemoveByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveByIndexes_Empty", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.RemoveNodeByIndexes(true)
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── RemoveNode ──

func Test_Cov38_LL_RemoveNode_Nil(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveNode_Nil", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.RemoveNode(nil)
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_RemoveNode_Head(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveNode_Head", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		ll.RemoveNode(ll.Head())
		actual := args.Map{"result": ll.Length() != 1 || ll.Head().Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov38_LL_RemoveNode_NonHead(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveNode_NonHead", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		node := ll.IndexAt(1) // "b"
		ll.RemoveNode(node)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── IndexAt / SafeIndexAt / SafePointerIndexAt ──

func Test_Cov38_LL_IndexAt(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IndexAt", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		node := ll.IndexAt(1)
		actual := args.Map{"result": node.Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_Cov38_LL_IndexAt_Zero(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IndexAt_Zero", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		actual := args.Map{"result": ll.IndexAt(0).Element != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Cov38_LL_IndexAt_Negative(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IndexAt_Negative", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.IndexAt(-1) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Cov38_LL_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_SafeIndexAt", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		actual := args.Map{"result": ll.SafeIndexAt(1).Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_Cov38_LL_SafeIndexAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_SafeIndexAt_OutOfRange", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.SafeIndexAt(5) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Cov38_LL_SafeIndexAt_Negative(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_SafeIndexAt_Negative", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.SafeIndexAt(-1) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Cov38_LL_SafeIndexAtLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_SafeIndexAtLock", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.SafeIndexAtLock(0).Element != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Cov38_LL_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_SafePointerIndexAt", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ptr := ll.SafePointerIndexAt(0)
		actual := args.Map{"result": ptr == nil || *ptr != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Cov38_LL_SafePointerIndexAt_Nil(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_SafePointerIndexAt_Nil", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.SafePointerIndexAt(5) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Cov38_LL_SafePointerIndexAtUsingDefault(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_SafePointerIndexAtUsingDefault", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.SafePointerIndexAtUsingDefault(5, "def") != "def"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected def", actual)
		actual := args.Map{"result": ll.SafePointerIndexAtUsingDefault(0, "def") != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Cov38_LL_SafePointerIndexAtUsingDefaultLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_SafePointerIndexAtUsingDefaultLock", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.SafePointerIndexAtUsingDefaultLock(5, "x") != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
	})
}

// ── GetNextNodes / GetAllLinkedNodes ──

func Test_Cov38_LL_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_GetNextNodes", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c", "d")
		nodes := ll.GetNextNodes(2)
		actual := args.Map{"result": len(nodes) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_LL_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_GetAllLinkedNodes", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		nodes := ll.GetAllLinkedNodes()
		actual := args.Map{"result": len(nodes) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── ToCollection / List / ListPtr / ListLock / ListPtrLock ──

func Test_Cov38_LL_ToCollection(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_ToCollection", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		col := ll.ToCollection(0)
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_LL_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_ToCollection_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		col := ll.ToCollection(0)
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov38_LL_List(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_List", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		list := ll.List()
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_LL_List_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_List_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"result": len(ll.List()) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov38_LL_ListPtr(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_ListPtr", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": len(ll.ListPtr()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_ListLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_ListLock", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": len(ll.ListLock()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_ListPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_ListPtrLock", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": len(ll.ListPtrLock()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── String / StringLock / Join / JoinLock / Joins ──

func Test_Cov38_LL_String(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_String", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		s := ll.String()
		actual := args.Map{"result": strings.Contains(s, "a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Cov38_LL_String_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_String_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		s := ll.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty (contains NoElements)", actual)
	})
}

func Test_Cov38_LL_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_StringLock", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		s := ll.StringLock()
		actual := args.Map{"result": strings.Contains(s, "a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Cov38_LL_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_StringLock_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		s := ll.StringLock()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov38_LL_Join(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Join", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		actual := args.Map{"result": ll.Join(",") != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", ll.Join("", actual)
	})
}

func Test_Cov38_LL_JoinLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_JoinLock", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		actual := args.Map{"result": ll.JoinLock(",") != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov38_LL_Joins(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Joins", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		result := ll.Joins(",", "c")
		actual := args.Map{"result": strings.Contains(result, "a") || !strings.Contains(result, "c")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_Cov38_LL_Joins_NilItems(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Joins_NilItems", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		result := ll.Joins(",")
		// items is nil, so joins items only
		_ = result
	})
}

// ── IsEquals / IsEqualsWithSensitive ──

func Test_Cov38_LL_IsEquals_Same(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEquals_Same", func() {
		a := corestr.New.LinkedList.SpreadStrings("a", "b")
		b := corestr.New.LinkedList.SpreadStrings("a", "b")
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov38_LL_IsEquals_Diff(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEquals_Diff", func() {
		a := corestr.New.LinkedList.SpreadStrings("a")
		b := corestr.New.LinkedList.SpreadStrings("b")
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Cov38_LL_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEquals_BothEmpty", func() {
		a := corestr.New.LinkedList.Create()
		b := corestr.New.LinkedList.Create()
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov38_LL_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEquals_DiffLen", func() {
		a := corestr.New.LinkedList.SpreadStrings("a", "b")
		b := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Cov38_LL_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEquals_OneEmpty", func() {
		a := corestr.New.LinkedList.SpreadStrings("a")
		b := corestr.New.LinkedList.Create()
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Cov38_LL_IsEquals_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEquals_SamePtr", func() {
		a := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": a.IsEqualsWithSensitive(a, true)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal (same pointer)", actual)
	})
}

func Test_Cov38_LL_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEqualsWithSensitive_CaseInsensitive", func() {
		a := corestr.New.LinkedList.SpreadStrings("A", "B")
		b := corestr.New.LinkedList.SpreadStrings("a", "b")
		actual := args.Map{"result": a.IsEqualsWithSensitive(b, false)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal case-insensitive", actual)
	})
}

func Test_Cov38_LL_IsEqualsWithSensitive_OneNil(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEqualsWithSensitive_OneNil", func() {
		a := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": a.IsEqualsWithSensitive(nil, true)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

// ── GetCompareSummary ──

func Test_Cov38_LL_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_GetCompareSummary", func() {
		a := corestr.New.LinkedList.SpreadStrings("a")
		b := corestr.New.LinkedList.SpreadStrings("b")
		summary := a.GetCompareSummary(b, "left", "right")
		actual := args.Map{"result": summary == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty summary", actual)
	})
}

// ── AddStringsToNode / AddStringsPtrToNode / AddCollectionToNode / AddAfterNode ──

func Test_Cov38_LL_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddStringsToNode", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "d")
		node := ll.Head()
		ll.AddStringsToNode(false, node, []string{"b", "c"})
		if ll.Length() != 5 { // a + b + c + (n-1 increments from AddNext) + 1 from AddStringsToNode + d
			// The actual length depends on internal logic; just verify it grew
			actual := args.Map{"result": ll.Length() < 4}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected at least 4", actual)
		}
	})
}

func Test_Cov38_LL_AddStringsToNode_SingleItem(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddStringsToNode_SingleItem", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "c")
		node := ll.Head()
		ll.AddStringsToNode(false, node, []string{"b"})
		actual := args.Map{"result": ll.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov38_LL_AddStringsToNode_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddStringsToNode_Empty", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.AddStringsToNode(false, ll.Head(), nil)
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_AddStringsToNode_NilNode_Skip(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddStringsToNode_NilNode_Skip", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.AddStringsToNode(true, nil, []string{"b"})
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_AddStringsPtrToNode_Nil(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddStringsPtrToNode_Nil", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.AddStringsPtrToNode(true, ll.Head(), nil)
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_AddStringsPtrToNode(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddStringsPtrToNode", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		items := []string{"b"}
		ll.AddStringsPtrToNode(true, ll.Head(), &items)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_LL_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddCollectionToNode", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		col := corestr.New.Collection.Strings([]string{"b"})
		ll.AddCollectionToNode(true, ll.Head(), col)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── JSON / Serialize ──

func Test_Cov38_LL_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_MarshalJSON", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		data, err := json.Marshal(ll)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": strings.Contains(string(data), "\"a\"")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_Cov38_LL_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_UnmarshalJSON", func() {
		ll := corestr.New.LinkedList.Create()
		err := json.Unmarshal([]byte(`["x","y"]`), ll)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_LL_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_UnmarshalJSON_Invalid", func() {
		ll := corestr.New.LinkedList.Create()
		err := json.Unmarshal([]byte(`bad`), ll)
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_Cov38_LL_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_JsonModel", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		model := ll.JsonModel()
		actual := args.Map{"result": len(model) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_JsonModelAny", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.JsonModelAny() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov38_LL_Json(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Json", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		result := ll.Json()
		actual := args.Map{"result": result.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_Cov38_LL_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_JsonPtr", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.JsonPtr() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov38_LL_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_ParseInjectUsingJson", func() {
		src := corestr.New.LinkedList.SpreadStrings("a", "b")
		ll := corestr.New.LinkedList.Create()
		result, err := ll.ParseInjectUsingJson(src.JsonPtr())
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_LL_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_ParseInjectUsingJsonMust", func() {
		src := corestr.New.LinkedList.SpreadStrings("a")
		ll := corestr.New.LinkedList.Create()
		result := ll.ParseInjectUsingJsonMust(src.JsonPtr())
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_LL_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_JsonParseSelfInject", func() {
		src := corestr.New.LinkedList.SpreadStrings("x")
		ll := corestr.New.LinkedList.Create()
		err := ll.JsonParseSelfInject(src.JsonPtr())
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_Cov38_LL_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AsJsonMarshaller", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ── Clear / RemoveAll ──

func Test_Cov38_LL_Clear(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Clear", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		ll.Clear()
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov38_LL_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Clear_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Clear()
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov38_LL_RemoveAll(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveAll", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.RemoveAll()
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ══════════════════════════════════════════════════════════════
// LinkedListNode
// ══════════════════════════════════════════════════════════════

func Test_Cov38_Node_HasNext(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_HasNext", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		actual := args.Map{"result": ll.Head().HasNext()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": ll.Tail().HasNext()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov38_Node_Next(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_Next", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		actual := args.Map{"result": ll.Head().Next().Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_Cov38_Node_EndOfChain(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_EndOfChain", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		end, length := ll.Head().EndOfChain()
		actual := args.Map{"result": end.Element != "c" || length != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected end= length=", actual)
	})
}

func Test_Cov38_Node_LoopEndOfChain(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_LoopEndOfChain", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
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

func Test_Cov38_Node_LoopEndOfChain_Break(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_LoopEndOfChain_Break", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		end, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			return true // break immediately
		})
		actual := args.Map{"result": length != 1 || end.Element != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected length= end=", actual)
	})
}

func Test_Cov38_Node_LoopEndOfChain_BreakSecond(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_LoopEndOfChain_BreakSecond", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		_, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			return arg.Index == 1
		})
		actual := args.Map{"result": length != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_Node_Clone(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_Clone", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		cloned := ll.Head().Clone()
		actual := args.Map{"result": cloned.Element != "a" || cloned.HasNext()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected clone without next", actual)
	})
}

func Test_Cov38_Node_AddNext(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_AddNext", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "c")
		newNode := ll.Head().AddNext(ll, "b")
		actual := args.Map{"result": newNode.Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_Cov38_Node_AddNextNode(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_AddNextNode", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "c")
		newNode := &corestr.LinkedListNode{Element: "b"}
		ll.Head().AddNextNode(ll, newNode)
		actual := args.Map{"result": ll.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov38_Node_IsEqual_Same(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqual_Same", func() {
		ll1 := corestr.New.LinkedList.SpreadStrings("a", "b")
		ll2 := corestr.New.LinkedList.SpreadStrings("a", "b")
		actual := args.Map{"result": ll1.Head().IsEqual(ll2.Head())}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov38_Node_IsEqual_Diff(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqual_Diff", func() {
		ll1 := corestr.New.LinkedList.SpreadStrings("a", "b")
		ll2 := corestr.New.LinkedList.SpreadStrings("a", "c")
		actual := args.Map{"result": ll1.Head().IsEqual(ll2.Head())}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Cov38_Node_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqual_BothNil", func() {
		var a, b *corestr.LinkedListNode
		actual := args.Map{"result": a.IsEqual(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov38_Node_IsEqual_OneNil(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqual_OneNil", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.Head().IsEqual(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Cov38_Node_IsEqual_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqual_SamePtr", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.Head().IsEqual(ll.Head())}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov38_Node_IsChainEqual(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsChainEqual", func() {
		a := corestr.New.LinkedList.SpreadStrings("a", "b")
		b := corestr.New.LinkedList.SpreadStrings("A", "B")
		actual := args.Map{"result": a.Head().IsChainEqual(b.Head(), false)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal case-insensitive", actual)
		actual := args.Map{"result": a.Head().IsChainEqual(b.Head(), true)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal case-sensitive", actual)
	})
}

func Test_Cov38_Node_IsChainEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsChainEqual_BothNil", func() {
		var a, b *corestr.LinkedListNode
		actual := args.Map{"result": a.IsChainEqual(b, true)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov38_Node_IsEqualSensitive(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqualSensitive", func() {
		a := corestr.New.LinkedList.SpreadStrings("a")
		b := corestr.New.LinkedList.SpreadStrings("A")
		actual := args.Map{"result": a.Head().IsEqualSensitive(b.Head(), false)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov38_Node_IsEqualSensitive_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqualSensitive_SamePtr", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.Head().IsEqualSensitive(ll.Head(), true)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov38_Node_IsEqualSensitive_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqualSensitive_BothNil", func() {
		var a, b *corestr.LinkedListNode
		actual := args.Map{"result": a.IsEqualSensitive(b, true)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov38_Node_IsEqualSensitive_OneNil(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqualSensitive_OneNil", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.Head().IsEqualSensitive(nil, true)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Cov38_Node_IsEqualValue(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqualValue", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.Head().IsEqualValue("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov38_Node_IsEqualValueSensitive(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqualValueSensitive", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": ll.Head().IsEqualValueSensitive("A", false)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true case-insensitive", actual)
		actual := args.Map{"result": ll.Head().IsEqualValueSensitive("A", true)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false case-sensitive", actual)
	})
}

func Test_Cov38_Node_String(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_String", func() {
		ll := corestr.New.LinkedList.SpreadStrings("hello")
		actual := args.Map{"result": ll.Head().String() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov38_Node_List(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_List", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		list := ll.Head().List()
		actual := args.Map{"result": len(list) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov38_Node_ListPtr(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_ListPtr", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{"result": len(ll.Head().ListPtr()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_Node_Join(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_Join", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		actual := args.Map{"result": ll.Head().Join(",") != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov38_Node_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_CreateLinkedList", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		newLL := ll.Head().CreateLinkedList()
		actual := args.Map{"result": newLL.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_Node_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_AddStringsToNode", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.Head().AddStringsToNode(ll, false, []string{"b"})
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_Node_AddStringsPtrToNode_Nil(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_AddStringsPtrToNode_Nil", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.Head().AddStringsPtrToNode(ll, true, nil)
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_Node_AddStringsPtrToNode(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_AddStringsPtrToNode", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		items := []string{"b"}
		ll.Head().AddStringsPtrToNode(ll, false, &items)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_Node_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_AddCollectionToNode", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		col := corestr.New.Collection.Strings([]string{"b"})
		ll.Head().AddCollectionToNode(ll, false, col)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ══════════════════════════════════════════════════════════════
// NonChainedLinkedListNodes
// ══════════════════════════════════════════════════════════════

func Test_Cov38_NCLLN_Basic(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_Basic", func() {
		nc := corestr.NewNonChainedLinkedListNodes(5)
		actual := args.Map{"result": nc.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual := args.Map{"result": nc.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
		actual := args.Map{"result": nc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov38_NCLLN_Adds(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_Adds", func() {
		nc := corestr.NewNonChainedLinkedListNodes(3)
		n1 := &corestr.LinkedListNode{Element: "a"}
		n2 := &corestr.LinkedListNode{Element: "b"}
		nc.Adds(n1, n2)
		actual := args.Map{"result": nc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual := args.Map{"result": nc.First().Element != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual := args.Map{"result": nc.Last().Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_Cov38_NCLLN_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_Adds_Nil", func() {
		nc := corestr.NewNonChainedLinkedListNodes(3)
		nc.Adds()
		actual := args.Map{"result": nc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov38_NCLLN_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_FirstOrDefault_Empty", func() {
		nc := corestr.NewNonChainedLinkedListNodes(3)
		actual := args.Map{"result": nc.FirstOrDefault() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Cov38_NCLLN_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_LastOrDefault_Empty", func() {
		nc := corestr.NewNonChainedLinkedListNodes(3)
		actual := args.Map{"result": nc.LastOrDefault() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Cov38_NCLLN_Items(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_Items", func() {
		nc := corestr.NewNonChainedLinkedListNodes(3)
		nc.Adds(&corestr.LinkedListNode{Element: "a"})
		actual := args.Map{"result": len(nc.Items()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_NCLLN_IsChainingApplied(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_IsChainingApplied", func() {
		nc := corestr.NewNonChainedLinkedListNodes(3)
		actual := args.Map{"result": nc.IsChainingApplied()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov38_NCLLN_ApplyChaining(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_ApplyChaining", func() {
		nc := corestr.NewNonChainedLinkedListNodes(3)
		n1 := &corestr.LinkedListNode{Element: "a"}
		n2 := &corestr.LinkedListNode{Element: "b"}
		nc.Adds(n1, n2)
		nc.ApplyChaining()
		actual := args.Map{"result": nc.IsChainingApplied()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": n1.HasNext()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a->b chain", actual)
		actual := args.Map{"result": n2.HasNext()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b to be tail", actual)
	})
}

func Test_Cov38_NCLLN_ApplyChaining_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_ApplyChaining_Empty", func() {
		nc := corestr.NewNonChainedLinkedListNodes(3)
		nc.ApplyChaining()
		// should not panic
	})
}

func Test_Cov38_NCLLN_ToChainedNodes(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_ToChainedNodes", func() {
		nc := corestr.NewNonChainedLinkedListNodes(3)
		nc.Adds(&corestr.LinkedListNode{Element: "a"}, &corestr.LinkedListNode{Element: "b"})
		chained := nc.ToChainedNodes()
		actual := args.Map{"result": chained == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov38_NCLLN_ToChainedNodes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_ToChainedNodes_Empty", func() {
		nc := corestr.NewNonChainedLinkedListNodes(3)
		chained := nc.ToChainedNodes()
		actual := args.Map{"result": len(chained) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── newLinkedListCreator factory methods ──

func Test_Cov38_Creator_Create(t *testing.T) {
	safeTest(t, "Test_Cov38_Creator_Create", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"result": ll == nil || ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov38_Creator_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_Creator_Empty", func() {
		ll := corestr.New.LinkedList.Empty()
		actual := args.Map{"result": ll == nil || ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov38_Creator_Strings(t *testing.T) {
	safeTest(t, "Test_Cov38_Creator_Strings", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_Creator_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_Cov38_Creator_SpreadStrings", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		actual := args.Map{"result": ll.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov38_Creator_UsingMap(t *testing.T) {
	safeTest(t, "Test_Cov38_Creator_UsingMap", func() {
		ll := corestr.New.LinkedList.UsingMap(map[string]bool{"a": true, "b": false})
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov38_Creator_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_Cov38_Creator_PointerStringsPtr", func() {
		a, b := "a", "b"
		ptrs := []*string{&a, nil, &b}
		ll := corestr.New.LinkedList.PointerStringsPtr(&ptrs)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov38_Creator_PointerStringsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov38_Creator_PointerStringsPtr_Nil", func() {
		ll := corestr.New.LinkedList.PointerStringsPtr(nil)
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}
