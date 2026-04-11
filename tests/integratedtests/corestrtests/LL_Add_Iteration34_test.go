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

func Test_LL_Add(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Add", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a").Add("b")

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LL_Head_Tail(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Head_Tail", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")

		// Act
		actual := args.Map{"result": ll.Head().Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": ll.Tail().Element != "b"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_LL_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": ll.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": ll.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
	})
}

func Test_LL_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEmptyLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": ll.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_LL_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_LengthLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_PushBack(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_PushBack", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.PushBack("x")

		// Act
		actual := args.Map{"result": ll.Length() != 1 || ll.Head().Element != "x"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_LL_Push(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Push", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Push("x")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_LL_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddLock("x")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_LL_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddNonEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmpty("")
		ll.AddNonEmpty("a")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddNonEmptyWhitespace", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmptyWhitespace("  ")
		ll.AddNonEmptyWhitespace("a")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_AddIf(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddIf", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddIf(false, "no")
		ll.AddIf(true, "yes")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_AddsIf_True(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddsIf_True", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddsIf(true, "a", "b")

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LL_AddsIf_False(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddsIf_False", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddsIf(false, "a", "b")

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LL_AddFunc(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddFunc", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddFunc(func() string { return "hello" })

		// Act
		actual := args.Map{"result": ll.Length() != 1 || ll.Head().Element != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_LL_AddFuncErr_NoErr(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddFuncErr_NoErr", func() {
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
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_AddFuncErr_WithErr(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddFuncErr_WithErr", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		called := false
		ll.AddFuncErr(
			func() (string, error) { return "", json.Unmarshal([]byte("bad"), &struct{}{}) },
			func(err error) { called = true },
		)

		// Act
		actual := args.Map{"result": called}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected error handler to be called", actual)
		actual = args.Map{"result": ll.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LL_Adds(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Adds", func() {
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

func Test_LL_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Adds_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds()

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LL_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddStrings", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings([]string{"x", "y"})

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LL_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddStrings_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings(nil)

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LL_AddsLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddsLock", func() {
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

func Test_LL_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddItemsMap", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{"a": true, "b": false, "c": true})

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LL_AddItemsMap_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddItemsMap_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(nil)

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LL_AddFront(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddFront", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("b")
		ll.AddFront("a")

		// Act
		actual := args.Map{"result": ll.Head().Element != "a" || ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_LL_AddFront_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddFront_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddFront("a")

		// Act
		actual := args.Map{"result": ll.Head().Element != "a" || ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_LL_PushFront(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_PushFront", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("b")
		ll.PushFront("a")

		// Act
		actual := args.Map{"result": ll.Head().Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_LL_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddCollection", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		ll := corestr.New.LinkedList.Create()
		ll.AddCollection(col)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LL_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddCollection_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddCollection(nil)

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LL_AddPointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddPointerStringsPtr", func() {
		// Arrange
		a, b := "a", "b"
		ll := corestr.New.LinkedList.Create()
		ll.AddPointerStringsPtr([]*string{&a, nil, &b})

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── InsertAt ──

func Test_LL_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_InsertAt_Front", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("b", "c")
		ll.InsertAt(0, "a")

		// Act
		actual := args.Map{"result": ll.Head().Element != "a" || ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_LL_InsertAt_Middle(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_InsertAt_Middle", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "c")
		ll.InsertAt(1, "b")
		list := ll.List()

		// Act
		actual := args.Map{"result": len(list) != 3 || list[1] != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

// ── AppendNode / AppendChainOfNodes / AddBackNode ──

func Test_LL_AppendNode_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AppendNode_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		node := ll.Head() // nil
		_ = node
		ll.AppendNode(&corestr.LinkedListNode{Element: "x"})

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_AppendNode_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AppendNode_NonEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.AppendNode(&corestr.LinkedListNode{Element: "b"})

		// Act
		actual := args.Map{"result": ll.Length() != 2 || ll.Tail().Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_LL_AddBackNode(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddBackNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddBackNode(&corestr.LinkedListNode{Element: "x"})

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AppendChainOfNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		chain := corestr.New.LinkedList.SpreadStrings("b", "c")
		ll.AppendChainOfNodes(chain.Head())

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_LL_AppendChainOfNodes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AppendChainOfNodes_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		chain := corestr.New.LinkedList.SpreadStrings("a", "b")
		ll.AppendChainOfNodes(chain.Head())

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── AttachWithNode ──

func Test_LL_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AttachWithNode_NilCurrent", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		err := ll.AttachWithNode(nil, &corestr.LinkedListNode{Element: "x"})

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_LL_AttachWithNode_NextNotNil(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AttachWithNode_NextNotNil", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		err := ll.AttachWithNode(ll.Head(), &corestr.LinkedListNode{Element: "x"})

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error because head.next is not nil", actual)
	})
}

// ── Loop ──

func Test_LL_Loop(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Loop", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
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

func Test_LL_Loop_Break(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Loop_Break", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return true
		})

		// Act
		actual := args.Map{"result": count != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Loop_Empty", func() {
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

func Test_LL_Loop_BreakSecond(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Loop_BreakSecond", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
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

// ── Filter ──

func Test_LL_Filter(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Filter", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: arg.Node.Element != "b"}
		})

		// Act
		actual := args.Map{"result": len(nodes) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LL_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Filter_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true}
		})

		// Act
		actual := args.Map{"result": len(nodes) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LL_Filter_BreakFirst(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Filter_BreakFirst", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})

		// Act
		actual := args.Map{"result": len(nodes) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_Filter_BreakSecond(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Filter_BreakSecond", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: arg.Index == 1}
		})

		// Act
		actual := args.Map{"result": len(nodes) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── RemoveNodeByElementValue ──

func Test_LL_RemoveByElem_CaseSensitive(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveByElem_CaseSensitive", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		ll.RemoveNodeByElementValue("b", true, false)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LL_RemoveByElem_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveByElem_CaseInsensitive", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "B", "c")
		ll.RemoveNodeByElementValue("b", false, false)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LL_RemoveByElem_First(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveByElem_First", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		ll.RemoveNodeByElementValue("a", true, false)

		// Act
		actual := args.Map{"result": ll.Length() != 1 || ll.Head().Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

// ── RemoveNodeByIndex ──

func Test_LL_RemoveByIndex_First(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveByIndex_First", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		ll.RemoveNodeByIndex(0)

		// Act
		actual := args.Map{"result": ll.Length() != 2 || ll.Head().Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_LL_RemoveByIndex_Last(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveByIndex_Last", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		ll.RemoveNodeByIndex(2)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LL_RemoveByIndex_Middle(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveByIndex_Middle", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		ll.RemoveNodeByIndex(1)
		list := ll.List()

		// Act
		actual := args.Map{"result": len(list) != 2 || list[0] != "a" || list[1] != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

// ── RemoveNodeByIndexes ──

func Test_LL_RemoveByIndexes(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveByIndexes", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c", "d")
		ll.RemoveNodeByIndexes(true, 1, 3)
		list := ll.List()

		// Act
		actual := args.Map{"result": len(list) != 2 || list[0] != "a" || list[1] != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_LL_RemoveByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveByIndexes_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.RemoveNodeByIndexes(true)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── RemoveNode ──

func Test_LL_RemoveNode_Nil(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveNode_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.RemoveNode(nil)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_RemoveNode_Head(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveNode_Head", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		ll.RemoveNode(ll.Head())

		// Act
		actual := args.Map{"result": ll.Length() != 1 || ll.Head().Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_LL_RemoveNode_NonHead(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveNode_NonHead", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		node := ll.IndexAt(1) // "b"
		ll.RemoveNode(node)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── IndexAt / SafeIndexAt / SafePointerIndexAt ──

func Test_LL_IndexAt(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IndexAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		node := ll.IndexAt(1)

		// Act
		actual := args.Map{"result": node.Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_LL_IndexAt_Zero(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IndexAt_Zero", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")

		// Act
		actual := args.Map{"result": ll.IndexAt(0).Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_LL_IndexAt_Negative(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IndexAt_Negative", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.IndexAt(-1) != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_LL_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_SafeIndexAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")

		// Act
		actual := args.Map{"result": ll.SafeIndexAt(1).Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_LL_SafeIndexAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_SafeIndexAt_OutOfRange", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.SafeIndexAt(5) != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_LL_SafeIndexAt_Negative(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_SafeIndexAt_Negative", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.SafeIndexAt(-1) != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_LL_SafeIndexAtLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_SafeIndexAtLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.SafeIndexAtLock(0).Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_LL_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_SafePointerIndexAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ptr := ll.SafePointerIndexAt(0)

		// Act
		actual := args.Map{"result": ptr == nil || *ptr != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_LL_SafePointerIndexAt_Nil(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_SafePointerIndexAt_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.SafePointerIndexAt(5) != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_LL_SafePointerIndexAtUsingDefault(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_SafePointerIndexAtUsingDefault", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.SafePointerIndexAtUsingDefault(5, "def") != "def"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected def", actual)
		actual = args.Map{"result": ll.SafePointerIndexAtUsingDefault(0, "def") != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_LL_SafePointerIndexAtUsingDefaultLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_SafePointerIndexAtUsingDefaultLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.SafePointerIndexAtUsingDefaultLock(5, "x") != "x"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
	})
}

// ── GetNextNodes / GetAllLinkedNodes ──

func Test_LL_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_GetNextNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c", "d")
		nodes := ll.GetNextNodes(2)

		// Act
		actual := args.Map{"result": len(nodes) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LL_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_GetAllLinkedNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		nodes := ll.GetAllLinkedNodes()

		// Act
		actual := args.Map{"result": len(nodes) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── ToCollection / List / ListPtr / ListLock / ListPtrLock ──

func Test_LL_ToCollection(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_ToCollection", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		col := ll.ToCollection(0)

		// Act
		actual := args.Map{"result": col.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LL_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_ToCollection_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		col := ll.ToCollection(0)

		// Act
		actual := args.Map{"result": col.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LL_List(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_List", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		list := ll.List()

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LL_List_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_List_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": len(ll.List()) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LL_ListPtr(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_ListPtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": len(ll.ListPtr()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_ListLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_ListLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": len(ll.ListLock()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_ListPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_ListPtrLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": len(ll.ListPtrLock()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── String / StringLock / Join / JoinLock / Joins ──

func Test_LL_String(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_String", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		s := ll.String()

		// Act
		actual := args.Map{"result": strings.Contains(s, "a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_LL_String_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_String_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		s := ll.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty (contains NoElements)", actual)
	})
}

func Test_LL_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_StringLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		s := ll.StringLock()

		// Act
		actual := args.Map{"result": strings.Contains(s, "a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_LL_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_StringLock_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		s := ll.StringLock()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_LL_Join(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Join", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")

		// Act
		actual := args.Map{"result": ll.Join(",") != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_LL_JoinLock(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_JoinLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")

		// Act
		actual := args.Map{"result": ll.JoinLock(",") != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_LL_Joins(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Joins", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		result := ll.Joins(",", "c")

		// Act
		actual := args.Map{"result": strings.Contains(result, "a") || !strings.Contains(result, "c")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_LL_Joins_NilItems(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Joins_NilItems", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a")
		result := ll.Joins(",")
		// items is nil, so joins items only
		_ = result
	})
}

// ── IsEquals / IsEqualsWithSensitive ──

func Test_LL_IsEquals_Same(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEquals_Same", func() {
		// Arrange
		a := corestr.New.LinkedList.SpreadStrings("a", "b")
		b := corestr.New.LinkedList.SpreadStrings("a", "b")

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_LL_IsEquals_Diff(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEquals_Diff", func() {
		// Arrange
		a := corestr.New.LinkedList.SpreadStrings("a")
		b := corestr.New.LinkedList.SpreadStrings("b")

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_LL_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEquals_BothEmpty", func() {
		// Arrange
		a := corestr.New.LinkedList.Create()
		b := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_LL_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEquals_DiffLen", func() {
		// Arrange
		a := corestr.New.LinkedList.SpreadStrings("a", "b")
		b := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_LL_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEquals_OneEmpty", func() {
		// Arrange
		a := corestr.New.LinkedList.SpreadStrings("a")
		b := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_LL_IsEquals_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEquals_SamePtr", func() {
		// Arrange
		a := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(a, true)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal (same pointer)", actual)
	})
}

func Test_LL_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEqualsWithSensitive_CaseInsensitive", func() {
		// Arrange
		a := corestr.New.LinkedList.SpreadStrings("A", "B")
		b := corestr.New.LinkedList.SpreadStrings("a", "b")

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(b, false)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal case-insensitive", actual)
	})
}

func Test_LL_IsEqualsWithSensitive_OneNil(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_IsEqualsWithSensitive_OneNil", func() {
		// Arrange
		a := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(nil, true)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

// ── GetCompareSummary ──

func Test_LL_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_GetCompareSummary", func() {
		// Arrange
		a := corestr.New.LinkedList.SpreadStrings("a")
		b := corestr.New.LinkedList.SpreadStrings("b")
		summary := a.GetCompareSummary(b, "left", "right")

		// Act
		actual := args.Map{"result": summary == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty summary", actual)
	})
}

// ── AddStringsToNode / AddStringsPtrToNode / AddCollectionToNode / AddAfterNode ──

func Test_LL_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddStringsToNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "d")
		node := ll.Head()
		ll.AddStringsToNode(false, node, []string{"b", "c"})

		// Act
		actual := args.Map{"result": ll.Length() < 4}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 4", actual)
	})
}

func Test_LL_AddStringsToNode_SingleItem(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddStringsToNode_SingleItem", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "c")
		node := ll.Head()
		ll.AddStringsToNode(false, node, []string{"b"})

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_LL_AddStringsToNode_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddStringsToNode_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.AddStringsToNode(false, ll.Head(), nil)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_AddStringsToNode_NilNode_Skip(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddStringsToNode_NilNode_Skip", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.AddStringsToNode(true, nil, []string{"b"})

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_AddStringsPtrToNode_Nil(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddStringsPtrToNode_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.AddStringsPtrToNode(true, ll.Head(), nil)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_AddStringsPtrToNode(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddStringsPtrToNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		items := []string{"b"}
		ll.AddStringsPtrToNode(true, ll.Head(), &items)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LL_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AddCollectionToNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		col := corestr.New.Collection.Strings([]string{"b"})
		ll.AddCollectionToNode(true, ll.Head(), col)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── JSON / Serialize ──

func Test_LL_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_MarshalJSON", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		data, err := json.Marshal(ll)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual = args.Map{"result": strings.Contains(string(data), "\"a\"")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_LL_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_UnmarshalJSON", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		err := json.Unmarshal([]byte(`["x","y"]`), ll)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual = args.Map{"result": ll.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LL_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_UnmarshalJSON_Invalid", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		err := json.Unmarshal([]byte(`bad`), ll)

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_LL_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_JsonModel", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		model := ll.JsonModel()

		// Act
		actual := args.Map{"result": len(model) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_JsonModelAny", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.JsonModelAny() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_LL_Json(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Json", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		result := ll.Json()

		// Act
		actual := args.Map{"result": result.HasError()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_LL_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_JsonPtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.JsonPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_LL_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_ParseInjectUsingJson", func() {
		// Arrange
		src := corestr.New.LinkedList.SpreadStrings("a", "b")
		ll := corestr.New.LinkedList.Create()
		result, err := ll.ParseInjectUsingJson(src.JsonPtr())

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual = args.Map{"result": result.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LL_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_ParseInjectUsingJsonMust", func() {
		// Arrange
		src := corestr.New.LinkedList.SpreadStrings("a")
		ll := corestr.New.LinkedList.Create()
		result := ll.ParseInjectUsingJsonMust(src.JsonPtr())

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LL_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_JsonParseSelfInject", func() {
		// Arrange
		src := corestr.New.LinkedList.SpreadStrings("x")
		ll := corestr.New.LinkedList.Create()
		err := ll.JsonParseSelfInject(src.JsonPtr())

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_LL_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_AsJsonMarshaller", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.AsJsonMarshaller() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ── Clear / RemoveAll ──

func Test_LL_Clear(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Clear", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		ll.Clear()

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LL_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_Clear_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Clear()

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LL_RemoveAll(t *testing.T) {
	safeTest(t, "Test_Cov38_LL_RemoveAll", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.RemoveAll()

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ══════════════════════════════════════════════════════════════
// LinkedListNode
// ══════════════════════════════════════════════════════════════

func Test_Node_HasNext(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_HasNext", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")

		// Act
		actual := args.Map{"result": ll.Head().HasNext()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ll.Tail().HasNext()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Node_Next(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_Next", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")

		// Act
		actual := args.Map{"result": ll.Head().Next().Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_Node_EndOfChain(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_EndOfChain", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		end, length := ll.Head().EndOfChain()

		// Act
		actual := args.Map{"result": end.Element != "c" || length != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected end= length=", actual)
	})
}

func Test_Node_LoopEndOfChain(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_LoopEndOfChain", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		count := 0
		end, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})

		// Act
		actual := args.Map{"result": count != 3 || length != 3 || end.Element != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected count= length= end=", actual)
	})
}

func Test_Node_LoopEndOfChain_Break(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_LoopEndOfChain_Break", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		end, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			return true // break immediately
		})

		// Act
		actual := args.Map{"result": length != 1 || end.Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected length= end=", actual)
	})
}

func Test_Node_LoopEndOfChain_BreakSecond(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_LoopEndOfChain_BreakSecond", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		_, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			return arg.Index == 1
		})

		// Act
		actual := args.Map{"result": length != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Node_Clone(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_Clone", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		cloned := ll.Head().Clone()

		// Act
		actual := args.Map{"result": cloned.Element != "a" || cloned.HasNext()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected clone without next", actual)
	})
}

func Test_Node_AddNext(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_AddNext", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "c")
		newNode := ll.Head().AddNext(ll, "b")

		// Act
		actual := args.Map{"result": newNode.Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_Node_AddNextNode(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_AddNextNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "c")
		newNode := &corestr.LinkedListNode{Element: "b"}
		ll.Head().AddNextNode(ll, newNode)

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Node_IsEqual_Same(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqual_Same", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.SpreadStrings("a", "b")
		ll2 := corestr.New.LinkedList.SpreadStrings("a", "b")

		// Act
		actual := args.Map{"result": ll1.Head().IsEqual(ll2.Head())}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Node_IsEqual_Diff(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqual_Diff", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.SpreadStrings("a", "b")
		ll2 := corestr.New.LinkedList.SpreadStrings("a", "c")

		// Act
		actual := args.Map{"result": ll1.Head().IsEqual(ll2.Head())}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Node_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqual_BothNil", func() {
		// Arrange
		var a, b *corestr.LinkedListNode

		// Act
		actual := args.Map{"result": a.IsEqual(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Node_IsEqual_OneNil(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqual_OneNil", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.Head().IsEqual(nil)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Node_IsEqual_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqual_SamePtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.Head().IsEqual(ll.Head())}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Node_IsChainEqual(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsChainEqual", func() {
		// Arrange
		a := corestr.New.LinkedList.SpreadStrings("a", "b")
		b := corestr.New.LinkedList.SpreadStrings("A", "B")

		// Act
		actual := args.Map{"result": a.Head().IsChainEqual(b.Head(), false)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal case-insensitive", actual)
		actual = args.Map{"result": a.Head().IsChainEqual(b.Head(), true)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal case-sensitive", actual)
	})
}

func Test_Node_IsChainEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsChainEqual_BothNil", func() {
		// Arrange
		var a, b *corestr.LinkedListNode

		// Act
		actual := args.Map{"result": a.IsChainEqual(b, true)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Node_IsEqualSensitive(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqualSensitive", func() {
		// Arrange
		a := corestr.New.LinkedList.SpreadStrings("a")
		b := corestr.New.LinkedList.SpreadStrings("A")

		// Act
		actual := args.Map{"result": a.Head().IsEqualSensitive(b.Head(), false)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Node_IsEqualSensitive_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqualSensitive_SamePtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.Head().IsEqualSensitive(ll.Head(), true)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Node_IsEqualSensitive_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqualSensitive_BothNil", func() {
		// Arrange
		var a, b *corestr.LinkedListNode

		// Act
		actual := args.Map{"result": a.IsEqualSensitive(b, true)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Node_IsEqualSensitive_OneNil(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqualSensitive_OneNil", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.Head().IsEqualSensitive(nil, true)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_Node_IsEqualValue(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqualValue", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.Head().IsEqualValue("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Node_IsEqualValueSensitive(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_IsEqualValueSensitive", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": ll.Head().IsEqualValueSensitive("A", false)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true case-insensitive", actual)
		actual = args.Map{"result": ll.Head().IsEqualValueSensitive("A", true)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false case-sensitive", actual)
	})
}

func Test_Node_String(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_String", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("hello")

		// Act
		actual := args.Map{"result": ll.Head().String() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Node_List(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_List", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		list := ll.Head().List()

		// Act
		actual := args.Map{"result": len(list) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Node_ListPtr(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_ListPtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")

		// Act
		actual := args.Map{"result": len(ll.Head().ListPtr()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Node_Join(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_Join", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")

		// Act
		actual := args.Map{"result": ll.Head().Join(",") != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Node_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_CreateLinkedList", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		newLL := ll.Head().CreateLinkedList()

		// Act
		actual := args.Map{"result": newLL.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Node_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_AddStringsToNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.Head().AddStringsToNode(ll, false, []string{"b"})

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Node_AddStringsPtrToNode_Nil(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_AddStringsPtrToNode_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		ll.Head().AddStringsPtrToNode(ll, true, nil)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Node_AddStringsPtrToNode(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_AddStringsPtrToNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		items := []string{"b"}
		ll.Head().AddStringsPtrToNode(ll, false, &items)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Node_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_Cov38_Node_AddCollectionToNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a")
		col := corestr.New.Collection.Strings([]string{"b"})
		ll.Head().AddCollectionToNode(ll, false, col)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ══════════════════════════════════════════════════════════════
// NonChainedLinkedListNodes
// ══════════════════════════════════════════════════════════════

func Test_NCLLN_Basic(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_Basic", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(5)

		// Act
		actual := args.Map{"result": nc.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": nc.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
		actual = args.Map{"result": nc.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_NCLLN_Adds(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_Adds", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(3)
		n1 := &corestr.LinkedListNode{Element: "a"}
		n2 := &corestr.LinkedListNode{Element: "b"}
		nc.Adds(n1, n2)

		// Act
		actual := args.Map{"result": nc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual = args.Map{"result": nc.First().Element != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": nc.Last().Element != "b"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_NCLLN_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_Adds_Nil", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(3)
		nc.Adds()

		// Act
		actual := args.Map{"result": nc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_NCLLN_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_FirstOrDefault_Empty", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(3)

		// Act
		actual := args.Map{"result": nc.FirstOrDefault() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_NCLLN_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_LastOrDefault_Empty", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(3)

		// Act
		actual := args.Map{"result": nc.LastOrDefault() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_NCLLN_Items(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_Items", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(3)
		nc.Adds(&corestr.LinkedListNode{Element: "a"})

		// Act
		actual := args.Map{"result": len(nc.Items()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_NCLLN_IsChainingApplied(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_IsChainingApplied", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(3)

		// Act
		actual := args.Map{"result": nc.IsChainingApplied()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_NCLLN_ApplyChaining(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_ApplyChaining", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(3)
		n1 := &corestr.LinkedListNode{Element: "a"}
		n2 := &corestr.LinkedListNode{Element: "b"}
		nc.Adds(n1, n2)
		nc.ApplyChaining()

		// Act
		actual := args.Map{"result": nc.IsChainingApplied()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": n1.HasNext()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a->b chain", actual)
		actual = args.Map{"result": n2.HasNext()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b to be tail", actual)
	})
}

func Test_NCLLN_ApplyChaining_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_ApplyChaining_Empty", func() {
		nc := corestr.NewNonChainedLinkedListNodes(3)
		nc.ApplyChaining()
		// should not panic
	})
}

func Test_NCLLN_ToChainedNodes(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_ToChainedNodes", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(3)
		nc.Adds(&corestr.LinkedListNode{Element: "a"}, &corestr.LinkedListNode{Element: "b"})
		chained := nc.ToChainedNodes()

		// Act
		actual := args.Map{"result": chained == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_NCLLN_ToChainedNodes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov38_NCLLN_ToChainedNodes_Empty", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(3)
		chained := nc.ToChainedNodes()

		// Act
		actual := args.Map{"result": len(chained) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── newLinkedListCreator factory methods ──

func Test_Creator_Create_FromLLAddIteration34(t *testing.T) {
	safeTest(t, "Test_Cov38_Creator_Create", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": ll == nil || ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Creator_Empty_FromLLAddIteration34(t *testing.T) {
	safeTest(t, "Test_Cov38_Creator_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Empty()

		// Act
		actual := args.Map{"result": ll == nil || ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Creator_Strings_FromLLAddIteration34(t *testing.T) {
	safeTest(t, "Test_Cov38_Creator_Strings", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Creator_SpreadStrings_FromLLAddIteration34(t *testing.T) {
	safeTest(t, "Test_Cov38_Creator_SpreadStrings", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Creator_UsingMap(t *testing.T) {
	safeTest(t, "Test_Cov38_Creator_UsingMap", func() {
		// Arrange
		ll := corestr.New.LinkedList.UsingMap(map[string]bool{"a": true, "b": false})

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Creator_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_Cov38_Creator_PointerStringsPtr", func() {
		// Arrange
		a, b := "a", "b"
		ptrs := []*string{&a, nil, &b}
		ll := corestr.New.LinkedList.PointerStringsPtr(&ptrs)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Creator_PointerStringsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov38_Creator_PointerStringsPtr_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.PointerStringsPtr(nil)

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}
