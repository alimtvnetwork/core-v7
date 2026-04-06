package corestrtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// LinkedList — Segment 12: Core methods (L1-600)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovLL1_01_IsEmpty_HasItems_Length(t *testing.T) {
	safeTest(t, "Test_CovLL1_01_IsEmpty_HasItems_Length", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()

		// Act
		actual := args.Map{"result": ll.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": ll.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
		actual = args.Map{"result": ll.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ll.Add("a")
		actual = args.Map{"result": ll.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
		actual = args.Map{"result": ll.HasItems()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected items", actual)
		actual = args.Map{"result": ll.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLL1_02_IsEmptyLock_LengthLock(t *testing.T) {
	safeTest(t, "Test_CovLL1_02_IsEmptyLock_LengthLock", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()

		// Act
		actual := args.Map{"result": ll.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": ll.LengthLock() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ll.Add("x")
		actual = args.Map{"result": ll.IsEmptyLock()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
		actual = args.Map{"result": ll.LengthLock() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLL1_03_Head_Tail(t *testing.T) {
	safeTest(t, "Test_CovLL1_03_Head_Tail", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()

		// Act
		actual := args.Map{"result": ll.Head() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		actual = args.Map{"result": ll.Tail() != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		ll.Add("a")
		actual = args.Map{"result": ll.Head() == nil || ll.Head().Element != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": ll.Tail() == nil || ll.Tail().Element != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		ll.Add("b")
		actual = args.Map{"result": ll.Head().Element != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": ll.Tail().Element != "b"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_CovLL1_04_Add_Multiple(t *testing.T) {
	safeTest(t, "Test_CovLL1_04_Add_Multiple", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b").Add("c")

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		items := ll.List()
		actual = args.Map{"result": items[0] != "a" || items[1] != "b" || items[2] != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected order", actual)
	})
}

func Test_CovLL1_05_AddLock(t *testing.T) {
	safeTest(t, "Test_CovLL1_05_AddLock", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.AddLock("a")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLL1_06_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_CovLL1_06_AddItemsMap", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.AddItemsMap(map[string]bool{"a": true, "b": false})

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty
		ll.AddItemsMap(nil)
		ll.AddItemsMap(map[string]bool{})
	})
}

func Test_CovLL1_07_AddFront_PushFront(t *testing.T) {
	safeTest(t, "Test_CovLL1_07_AddFront_PushFront", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("b")
		ll.AddFront("a")
		items := ll.List()

		// Act
		actual := args.Map{"result": items[0] != "a" || items[1] != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected order", actual)
		// AddFront on empty
		ll2 := corestr.Empty.LinkedList()
		ll2.AddFront("x")
		actual = args.Map{"result": ll2.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// PushFront
		ll2.PushFront("y")
		actual = args.Map{"result": ll2.Head().Element != "y"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected y", actual)
	})
}

func Test_CovLL1_08_Push_PushBack(t *testing.T) {
	safeTest(t, "Test_CovLL1_08_Push_PushBack", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Push("a")
		ll.PushBack("b")

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLL1_09_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_CovLL1_09_AddNonEmpty", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.AddNonEmpty("")

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ll.AddNonEmpty("a")
		actual = args.Map{"result": ll.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLL1_10_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_CovLL1_10_AddNonEmptyWhitespace", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.AddNonEmptyWhitespace("   ")

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ll.AddNonEmptyWhitespace("a")
		actual = args.Map{"result": ll.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLL1_11_AddIf(t *testing.T) {
	safeTest(t, "Test_CovLL1_11_AddIf", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.AddIf(false, "a")

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ll.AddIf(true, "a")
		actual = args.Map{"result": ll.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLL1_12_AddsIf(t *testing.T) {
	safeTest(t, "Test_CovLL1_12_AddsIf", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.AddsIf(false, "a", "b")

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ll.AddsIf(true, "a", "b")
		actual = args.Map{"result": ll.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLL1_13_AddFunc(t *testing.T) {
	safeTest(t, "Test_CovLL1_13_AddFunc", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.AddFunc(func() string { return "hello" })

		// Act
		actual := args.Map{"result": ll.Length() != 1 || ll.Head().Element != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_CovLL1_14_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_CovLL1_14_AddFuncErr", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		// success
		ll.AddFuncErr(
			func() (string, error) { return "ok", nil },

		// Assert
			func(err error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) },
		)

		// Act
		actual = args.Map{"result": ll.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// error
		ll.AddFuncErr(
			func() (string, error) { return "", fmt.Errorf("fail") },
			func(err error) {},
		)
		actual = args.Map{"result": ll.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 1", actual)
	})
}

func Test_CovLL1_15_Adds_AddStrings(t *testing.T) {
	safeTest(t, "Test_CovLL1_15_Adds_AddStrings", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		ll.Adds()
		ll2 := corestr.Empty.LinkedList()
		ll2.AddStrings([]string{"x", "y"})
		actual = args.Map{"result": ll2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		ll2.AddStrings(nil)
	})
}

func Test_CovLL1_16_AddsLock(t *testing.T) {
	safeTest(t, "Test_CovLL1_16_AddsLock", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.AddsLock("a", "b")

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLL1_17_InsertAt(t *testing.T) {
	safeTest(t, "Test_CovLL1_17_InsertAt", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "c")
		ll.InsertAt(1, "b")
		items := ll.List()

		// Act
		actual := args.Map{"result": len(items) < 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3 items", actual)
		// index < 1 → AddFront
		ll.InsertAt(-1, "z")
		actual = args.Map{"result": ll.Head().Element != "z"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected z at front", actual)
	})
}

func Test_CovLL1_18_AppendNode_AddBackNode(t *testing.T) {
	safeTest(t, "Test_CovLL1_18_AppendNode_AddBackNode", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		node := &corestr.LinkedListNode{Element: "a"}
		ll.AppendNode(node)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// not empty
		node2 := &corestr.LinkedListNode{Element: "b"}
		ll.AddBackNode(node2)
		actual = args.Map{"result": ll.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLL1_19_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_CovLL1_19_AppendChainOfNodes", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		// build chain via another list
		chain := corestr.Empty.LinkedList()
		chain.Adds("a", "b", "c")
		ll.AppendChainOfNodes(chain.Head())

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// append to non-empty
		chain2 := corestr.Empty.LinkedList()
		chain2.Adds("d", "e")
		ll.AppendChainOfNodes(chain2.Head())
		actual = args.Map{"result": ll.Length() != 5}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5", actual)
	})
}

func Test_CovLL1_20_AddPointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_CovLL1_20_AddPointerStringsPtr", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		a := "a"
		ll.AddPointerStringsPtr([]*string{&a, nil})

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLL1_21_AddCollection(t *testing.T) {
	safeTest(t, "Test_CovLL1_21_AddCollection", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.AddCollection(nil)

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		ll.AddCollection(col)
		actual = args.Map{"result": ll.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLL1_22_AttachWithNode(t *testing.T) {
	safeTest(t, "Test_CovLL1_22_AttachWithNode", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		node := ll.Head()
		addNode := &corestr.LinkedListNode{Element: "b"}
		err := ll.AttachWithNode(node, addNode)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		// nil current
		err2 := ll.AttachWithNode(nil, addNode)
		actual = args.Map{"result": err2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
		// current.next not nil
		err3 := ll.AttachWithNode(node, addNode)
		actual = args.Map{"result": err3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error for non-nil next", actual)
	})
}

func Test_CovLL1_23_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_CovLL1_23_AddStringsToNode", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		node := ll.Head()
		ll.AddStringsToNode(false, node, []string{"b", "c"})

		// Act
		actual := args.Map{"result": ll.Length() < 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
		// single item
		ll2 := corestr.Empty.LinkedList()
		ll2.Add("a")
		ll2.AddStringsToNode(false, ll2.Head(), []string{"b"})
		// empty items
		ll2.AddStringsToNode(false, ll2.Head(), nil)
		// nil node skip
		ll2.AddStringsToNode(true, nil, []string{"x"})
	})
}

func Test_CovLL1_24_AddStringsPtrToNode(t *testing.T) {
	safeTest(t, "Test_CovLL1_24_AddStringsPtrToNode", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		items := []string{"b"}
		ll.AddStringsPtrToNode(false, ll.Head(), &items)
		// nil
		ll.AddStringsPtrToNode(false, ll.Head(), nil)
	})
}

func Test_CovLL1_25_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_CovLL1_25_AddCollectionToNode", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		col := corestr.New.Collection.Strings([]string{"b", "c"})
		ll.AddCollectionToNode(true, ll.Head(), col)

		// Act
		actual := args.Map{"result": ll.Length() < 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	})
}

func Test_CovLL1_26_Loop(t *testing.T) {
	safeTest(t, "Test_CovLL1_26_Loop", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		// empty
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})

		// Act
		actual := args.Map{"result": count != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 iterations", actual)
		// with items
		ll.Adds("a", "b", "c")
		count = 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})
		actual := args.Map{"result": count != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// break
		count = 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return true
		})
		actual := args.Map{"result": count != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// break on second
		count = 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return arg.Index == 1
		})
		actual := args.Map{"result": count != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLL1_27_Filter(t *testing.T) {
	safeTest(t, "Test_CovLL1_27_Filter", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		// empty
		r := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true}
		})

		// Act
		actual := args.Map{"result": len(r) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ll.Adds("a", "b", "c")
		// keep all
		r2 := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true}
		})
		actual := args.Map{"result": len(r2) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// break on first
		r3 := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})
		actual := args.Map{"result": len(r3) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// skip all
		r4 := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: false}
		})
		actual := args.Map{"result": len(r4) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// break on loop iteration
		r5 := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: arg.Index == 1}
		})
		actual := args.Map{"result": len(r5) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLL1_28_RemoveNodeByElementValue(t *testing.T) {
	safeTest(t, "Test_CovLL1_28_RemoveNodeByElementValue", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByElementValue("a", true, false)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// case insensitive
		ll.RemoveNodeByElementValue("B", false, false)
		actual = args.Map{"result": ll.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// not found
		ll.RemoveNodeByElementValue("z", true, true)
		// remove non-first
		ll2 := corestr.Empty.LinkedList()
		ll2.Adds("x", "y", "z")
		ll2.RemoveNodeByElementValue("y", true, false)
		actual = args.Map{"result": ll2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLL1_29_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_CovLL1_29_RemoveNodeByIndex", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		// remove first
		ll.RemoveNodeByIndex(0)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// remove last
		ll.RemoveNodeByIndex(1)
		actual = args.Map{"result": ll.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// remove middle
		ll2 := corestr.Empty.LinkedList()
		ll2.Adds("a", "b", "c")
		ll2.RemoveNodeByIndex(1)
		actual = args.Map{"result": ll2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLL1_30_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_CovLL1_30_RemoveNodeByIndexes", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c", "d")
		ll.RemoveNodeByIndexes(false, 1, 3)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty indexes
		ll.RemoveNodeByIndexes(false)
		// ignore panic on empty
		empty := corestr.Empty.LinkedList()
		empty.RemoveNodeByIndexes(true, 0)
	})
}

func Test_CovLL1_31_RemoveNode(t *testing.T) {
	safeTest(t, "Test_CovLL1_31_RemoveNode", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		// nil → skip
		ll.RemoveNode(nil)

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// remove head
		ll.RemoveNode(ll.Head())
		actual = args.Map{"result": ll.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// remove non-head
		node := ll.Head().Next()
		ll.RemoveNode(node)
		actual = args.Map{"result": ll.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLL1_32_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_CovLL1_32_GetCompareSummary", func() {
		// Arrange
		a := corestr.Empty.LinkedList()
		a.Adds("a", "b")
		b := corestr.Empty.LinkedList()
		b.Adds("a", "b")
		s := a.GetCompareSummary(b, "left", "right")

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_CovLL1_33_IndexAt(t *testing.T) {
	safeTest(t, "Test_CovLL1_33_IndexAt", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		node := ll.IndexAt(0)

		// Act
		actual := args.Map{"result": node.Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		node2 := ll.IndexAt(2)
		actual = args.Map{"result": node2.Element != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c", actual)
		// negative
		n := ll.IndexAt(-1)
		actual = args.Map{"result": n != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_CovLL1_34_SafeIndexAt_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_CovLL1_34_SafeIndexAt_SafePointerIndexAt", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		// found
		node := ll.SafeIndexAt(0)

		// Act
		actual := args.Map{"result": node == nil || node.Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		node1 := ll.SafeIndexAt(1)
		actual = args.Map{"result": node1 == nil || node1.Element != "b"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
		// not found
		n := ll.SafeIndexAt(-1)
		actual = args.Map{"result": n != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		n2 := ll.SafeIndexAt(99)
		actual = args.Map{"result": n2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		// empty
		e := corestr.Empty.LinkedList()
		actual = args.Map{"result": e.SafeIndexAt(0) != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		// pointer
		p := ll.SafePointerIndexAt(0)
		actual = args.Map{"result": p == nil || *p != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		p2 := ll.SafePointerIndexAt(-1)
		actual = args.Map{"result": p2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_CovLL1_35_SafePointerIndexAtUsingDefault(t *testing.T) {
	safeTest(t, "Test_CovLL1_35_SafePointerIndexAtUsingDefault", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		v := ll.SafePointerIndexAtUsingDefault(0, "def")

		// Act
		actual := args.Map{"result": v != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		v2 := ll.SafePointerIndexAtUsingDefault(99, "def")
		actual = args.Map{"result": v2 != "def"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected def", actual)
	})
}

func Test_CovLL1_36_SafeIndexAtLock_SafePointerIndexAtUsingDefaultLock(t *testing.T) {
	safeTest(t, "Test_CovLL1_36_SafeIndexAtLock_SafePointerIndexAtUsingDefaultLock", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		n := ll.SafeIndexAtLock(0)

		// Act
		actual := args.Map{"result": n == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		v := ll.SafePointerIndexAtUsingDefaultLock(0, "def")
		actual = args.Map{"result": v != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_CovLL1_37_GetNextNodes_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_CovLL1_37_GetNextNodes_GetAllLinkedNodes", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		r := ll.GetNextNodes(2)

		// Act
		actual := args.Map{"result": len(r) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		all := ll.GetAllLinkedNodes()
		actual = args.Map{"result": len(all) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}
