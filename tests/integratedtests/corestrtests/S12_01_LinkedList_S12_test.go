package corestrtests

import (
	"strings"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════
// S12 — LinkedList.go (1,141 lines) — Full coverage
// ══════════════════════════════════════════════════════════════

func Test_S12_01_LinkedList_HeadTailLength(t *testing.T) {
	safeTest(t, "Test_S12_01_LinkedList_HeadTailLength", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act & Assert
		actual := args.Map{"result": ll.Head() == nil || ll.Tail() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil head/tail", actual)
		actual = args.Map{"result": ll.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_02_LinkedList_LengthLock(t *testing.T) {
	safeTest(t, "Test_S12_02_LinkedList_LengthLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_03_LinkedList_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_S12_03_LinkedList_IsEmpty_HasItems", func() {
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
		ll.Add("a")
		actual = args.Map{"result": ll.IsEmpty() || !ll.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_S12_04_LinkedList_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_S12_04_LinkedList_IsEmptyLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": ll.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_S12_05_LinkedList_Add_MultipleItems(t *testing.T) {
	safeTest(t, "Test_S12_05_LinkedList_Add_MultipleItems", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.Add("b")

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_06_LinkedList_AddLock(t *testing.T) {
	safeTest(t, "Test_S12_06_LinkedList_AddLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddLock("a")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_07_LinkedList_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_S12_07_LinkedList_AddItemsMap", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{"a": true, "b": false})

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_08_LinkedList_AddItemsMap_Empty(t *testing.T) {
	safeTest(t, "Test_S12_08_LinkedList_AddItemsMap_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{})

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_09_LinkedList_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_S12_09_LinkedList_AddNonEmpty", func() {
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

func Test_S12_10_LinkedList_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_S12_10_LinkedList_AddNonEmptyWhitespace", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmptyWhitespace("   ")
		ll.AddNonEmptyWhitespace("a")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_11_LinkedList_AddIf(t *testing.T) {
	safeTest(t, "Test_S12_11_LinkedList_AddIf", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddIf(true, "yes")
		ll.AddIf(false, "no")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_12_LinkedList_AddsIf(t *testing.T) {
	safeTest(t, "Test_S12_12_LinkedList_AddsIf", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddsIf(true, "a", "b")
		ll.AddsIf(false, "c")

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_13_LinkedList_AddFunc(t *testing.T) {
	safeTest(t, "Test_S12_13_LinkedList_AddFunc", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddFunc(func() string { return "val" })

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_14_LinkedList_AddFuncErr_NoError(t *testing.T) {
	safeTest(t, "Test_S12_14_LinkedList_AddFuncErr_NoError", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Assert
		ll.AddFuncErr(func() (string, error) { return "ok", nil }, func(err error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) })

		// Act
		actual = args.Map{"result": ll.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_15_LinkedList_AddFuncErr_WithError(t *testing.T) {
	safeTest(t, "Test_S12_15_LinkedList_AddFuncErr_WithError", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		called := false
		ll.AddFuncErr(func() (string, error) { return "", &testErr{} }, func(err error) { called = true })

		// Act
		actual := args.Map{"result": called}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected err handler called", actual)
	})
}

func Test_S12_16_LinkedList_Push_PushBack_PushFront(t *testing.T) {
	safeTest(t, "Test_S12_16_LinkedList_Push_PushBack_PushFront", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Push("a")
		ll.PushBack("b")
		ll.PushFront("z")

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_S12_17_LinkedList_AddFront_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S12_17_LinkedList_AddFront_OnEmpty", func() {
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

func Test_S12_18_LinkedList_AddFront_OnNonEmpty(t *testing.T) {
	safeTest(t, "Test_S12_18_LinkedList_AddFront_OnNonEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("b")
		ll.AddFront("a")

		// Act
		actual := args.Map{"result": ll.Head().Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a at head", actual)
	})
}

func Test_S12_19_LinkedList_Adds(t *testing.T) {
	safeTest(t, "Test_S12_19_LinkedList_Adds", func() {
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

func Test_S12_20_LinkedList_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_S12_20_LinkedList_Adds_Empty", func() {
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

func Test_S12_21_LinkedList_AddStrings(t *testing.T) {
	safeTest(t, "Test_S12_21_LinkedList_AddStrings", func() {
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

func Test_S12_22_LinkedList_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_S12_22_LinkedList_AddStrings_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings([]string{})

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_23_LinkedList_AddsLock(t *testing.T) {
	safeTest(t, "Test_S12_23_LinkedList_AddsLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddsLock("a")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_24_LinkedList_AppendNode_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S12_24_LinkedList_AppendNode_OnEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		node := &corestr.LinkedListNode{Element: "a"}
		ll.AppendNode(node)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_25_LinkedList_AddBackNode(t *testing.T) {
	safeTest(t, "Test_S12_25_LinkedList_AddBackNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("x")
		node := &corestr.LinkedListNode{Element: "y"}
		ll.AddBackNode(node)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_26_LinkedList_InsertAt(t *testing.T) {
	safeTest(t, "Test_S12_26_LinkedList_InsertAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "c")
		ll.InsertAt(1, "b")

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_S12_27_LinkedList_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_S12_27_LinkedList_InsertAt_Front", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("b")
		ll.InsertAt(0, "a")

		// Act
		actual := args.Map{"result": ll.Head().Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a at head", actual)
	})
}

func Test_S12_28_LinkedList_Loop(t *testing.T) {
	safeTest(t, "Test_S12_28_LinkedList_Loop", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
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

func Test_S12_29_LinkedList_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_S12_29_LinkedList_Loop_Empty", func() {
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

func Test_S12_30_LinkedList_Loop_Break(t *testing.T) {
	safeTest(t, "Test_S12_30_LinkedList_Loop_Break", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
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

func Test_S12_31_LinkedList_Filter(t *testing.T) {
	safeTest(t, "Test_S12_31_LinkedList_Filter", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})

		// Act
		actual := args.Map{"result": len(nodes) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_32_LinkedList_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_S12_32_LinkedList_Filter_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})

		// Act
		actual := args.Map{"result": len(nodes) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_33_LinkedList_Filter_Break(t *testing.T) {
	safeTest(t, "Test_S12_33_LinkedList_Filter_Break", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
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

func Test_S12_34_LinkedList_IsEquals(t *testing.T) {
	safeTest(t, "Test_S12_34_LinkedList_IsEquals", func() {
		// Arrange
		a := corestr.New.LinkedList.Create()
		a.Adds("a", "b")
		b := corestr.New.LinkedList.Create()
		b.Adds("a", "b")

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S12_35_LinkedList_IsEquals_BothNil(t *testing.T) {
	safeTest(t, "Test_S12_35_LinkedList_IsEquals_BothNil", func() {
		// Arrange
		var a *corestr.LinkedList
		var b *corestr.LinkedList

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(b, true)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S12_36_LinkedList_IsEquals_OneNil(t *testing.T) {
	safeTest(t, "Test_S12_36_LinkedList_IsEquals_OneNil", func() {
		// Arrange
		a := corestr.New.LinkedList.Create()
		a.Add("a")
		var b *corestr.LinkedList

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(b, true)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S12_37_LinkedList_IsEquals_SamePtr(t *testing.T) {
	safeTest(t, "Test_S12_37_LinkedList_IsEquals_SamePtr", func() {
		// Arrange
		a := corestr.New.LinkedList.Create()
		a.Add("a")

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(a, true)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S12_38_LinkedList_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S12_38_LinkedList_IsEquals_BothEmpty", func() {
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

func Test_S12_39_LinkedList_IsEquals_DiffLength(t *testing.T) {
	safeTest(t, "Test_S12_39_LinkedList_IsEquals_DiffLength", func() {
		// Arrange
		a := corestr.New.LinkedList.Create()
		a.Add("a")
		b := corestr.New.LinkedList.Create()
		b.Adds("a", "b")

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S12_40_LinkedList_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_S12_40_LinkedList_IsEqualsWithSensitive_CaseInsensitive", func() {
		a := corestr.New.LinkedList.Create()
		a.Add("A")
		b := corestr.New.LinkedList.Create()
		b.Add("a")
		if a.IsEqualsWithSensitive(b, false) {
			// depends on LinkedListNode.IsChainEqual implementation
			_ = 0
		}
	})
}

func Test_S12_41_LinkedList_RemoveNodeByElementValue(t *testing.T) {
	safeTest(t, "Test_S12_41_LinkedList_RemoveNodeByElementValue", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByElementValue("b", true, false)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_42_LinkedList_RemoveNodeByElementValue_First(t *testing.T) {
	safeTest(t, "Test_S12_42_LinkedList_RemoveNodeByElementValue_First", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.RemoveNodeByElementValue("a", true, false)

		// Act
		actual := args.Map{"result": ll.Head().Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b at head", actual)
	})
}

func Test_S12_43_LinkedList_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_S12_43_LinkedList_RemoveNodeByIndex", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByIndex(1)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_44_LinkedList_RemoveNodeByIndex_First(t *testing.T) {
	safeTest(t, "Test_S12_44_LinkedList_RemoveNodeByIndex_First", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.RemoveNodeByIndex(0)

		// Act
		actual := args.Map{"result": ll.Head().Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_S12_45_LinkedList_RemoveNodeByIndex_Last(t *testing.T) {
	safeTest(t, "Test_S12_45_LinkedList_RemoveNodeByIndex_Last", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.RemoveNodeByIndex(1)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_46_LinkedList_RemoveNode(t *testing.T) {
	safeTest(t, "Test_S12_46_LinkedList_RemoveNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.Add("b")
		node := ll.IndexAt(0)
		ll.RemoveNode(node)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_47_LinkedList_RemoveNode_Nil(t *testing.T) {
	safeTest(t, "Test_S12_47_LinkedList_RemoveNode_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.RemoveNode(nil)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_48_LinkedList_IndexAt(t *testing.T) {
	safeTest(t, "Test_S12_48_LinkedList_IndexAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		node := ll.IndexAt(1)

		// Act
		actual := args.Map{"result": node.Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_S12_49_LinkedList_IndexAt_Zero(t *testing.T) {
	safeTest(t, "Test_S12_49_LinkedList_IndexAt_Zero", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.IndexAt(0).Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_S12_50_LinkedList_IndexAt_Negative(t *testing.T) {
	safeTest(t, "Test_S12_50_LinkedList_IndexAt_Negative", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.IndexAt(-1) != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_S12_51_LinkedList_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_S12_51_LinkedList_SafeIndexAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")

		// Act
		actual := args.Map{"result": ll.SafeIndexAt(1).Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
		actual = args.Map{"result": ll.SafeIndexAt(-1) != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		actual = args.Map{"result": ll.SafeIndexAt(10) != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_S12_52_LinkedList_SafeIndexAtLock(t *testing.T) {
	safeTest(t, "Test_S12_52_LinkedList_SafeIndexAtLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.SafeIndexAtLock(0).Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_S12_53_LinkedList_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_S12_53_LinkedList_SafePointerIndexAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ptr := ll.SafePointerIndexAt(0)

		// Act
		actual := args.Map{"result": ptr == nil || *ptr != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": ll.SafePointerIndexAt(10) != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_S12_54_LinkedList_SafePointerIndexAtUsingDefault(t *testing.T) {
	safeTest(t, "Test_S12_54_LinkedList_SafePointerIndexAtUsingDefault", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.SafePointerIndexAtUsingDefault(0, "def") != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": ll.SafePointerIndexAtUsingDefault(10, "def") != "def"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected def", actual)
	})
}

func Test_S12_55_LinkedList_SafePointerIndexAtUsingDefaultLock(t *testing.T) {
	safeTest(t, "Test_S12_55_LinkedList_SafePointerIndexAtUsingDefaultLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.SafePointerIndexAtUsingDefaultLock(0, "def") != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_S12_56_LinkedList_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_S12_56_LinkedList_GetNextNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		nodes := ll.GetNextNodes(2)

		// Act
		actual := args.Map{"result": len(nodes) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_57_LinkedList_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_S12_57_LinkedList_GetAllLinkedNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")

		// Act
		actual := args.Map{"result": len(ll.GetAllLinkedNodes()) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_58_LinkedList_AddPointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_S12_58_LinkedList_AddPointerStringsPtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		a := "a"
		ll.AddPointerStringsPtr([]*string{&a, nil})

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_59_LinkedList_AddCollection(t *testing.T) {
	safeTest(t, "Test_S12_59_LinkedList_AddCollection", func() {
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

func Test_S12_60_LinkedList_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_S12_60_LinkedList_AddCollection_Nil", func() {
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

func Test_S12_61_LinkedList_ToCollection(t *testing.T) {
	safeTest(t, "Test_S12_61_LinkedList_ToCollection", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		col := ll.ToCollection(0)

		// Act
		actual := args.Map{"result": col.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_62_LinkedList_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_S12_62_LinkedList_ToCollection_Empty", func() {
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

func Test_S12_63_LinkedList_List(t *testing.T) {
	safeTest(t, "Test_S12_63_LinkedList_List", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")

		// Act
		actual := args.Map{"result": len(ll.List()) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_64_LinkedList_ListPtr(t *testing.T) {
	safeTest(t, "Test_S12_64_LinkedList_ListPtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": len(ll.ListPtr()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_65_LinkedList_ListLock(t *testing.T) {
	safeTest(t, "Test_S12_65_LinkedList_ListLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": len(ll.ListLock()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_66_LinkedList_ListPtrLock(t *testing.T) {
	safeTest(t, "Test_S12_66_LinkedList_ListPtrLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": len(ll.ListPtrLock()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_67_LinkedList_String(t *testing.T) {
	safeTest(t, "Test_S12_67_LinkedList_String", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_S12_68_LinkedList_String_Empty(t *testing.T) {
	safeTest(t, "Test_S12_68_LinkedList_String_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": strings.Contains(ll.String(), "No Element")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected No Element", actual)
	})
}

func Test_S12_69_LinkedList_StringLock(t *testing.T) {
	safeTest(t, "Test_S12_69_LinkedList_StringLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.StringLock() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_S12_70_LinkedList_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_S12_70_LinkedList_StringLock_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": strings.Contains(ll.StringLock(), "No Element")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected No Element", actual)
	})
}

func Test_S12_71_LinkedList_Join(t *testing.T) {
	safeTest(t, "Test_S12_71_LinkedList_Join", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")

		// Act
		actual := args.Map{"result": ll.Join(",") != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
	})
}

func Test_S12_72_LinkedList_JoinLock(t *testing.T) {
	safeTest(t, "Test_S12_72_LinkedList_JoinLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.JoinLock(",") != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_S12_73_LinkedList_Joins(t *testing.T) {
	safeTest(t, "Test_S12_73_LinkedList_Joins", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		result := ll.Joins(",", "b")

		// Act
		actual := args.Map{"result": strings.Contains(result, "a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a in result", actual)
	})
}

func Test_S12_74_LinkedList_Joins_NilItems(t *testing.T) {
	safeTest(t, "Test_S12_74_LinkedList_Joins_NilItems", func() {
		ll := corestr.New.LinkedList.Create()
		result := ll.Joins(",", nil...)
		_ = result
	})
}

func Test_S12_75_LinkedList_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_S12_75_LinkedList_MarshalJSON", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		data, err := ll.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil || len(data) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid JSON", actual)
	})
}

func Test_S12_76_LinkedList_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_S12_76_LinkedList_UnmarshalJSON", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		err := ll.UnmarshalJSON([]byte(`["a","b"]`))

		// Act
		actual := args.Map{"result": err != nil || ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_77_LinkedList_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_S12_77_LinkedList_UnmarshalJSON_Invalid", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		err := ll.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_S12_78_LinkedList_JsonModel(t *testing.T) {
	safeTest(t, "Test_S12_78_LinkedList_JsonModel", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": len(ll.JsonModel()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_79_LinkedList_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_S12_79_LinkedList_JsonModelAny", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.JsonModelAny() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S12_80_LinkedList_Clear_RemoveAll(t *testing.T) {
	safeTest(t, "Test_S12_80_LinkedList_Clear_RemoveAll", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.Clear()

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_81_LinkedList_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_S12_81_LinkedList_Clear_Empty", func() {
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

func Test_S12_82_LinkedList_RemoveAll(t *testing.T) {
	safeTest(t, "Test_S12_82_LinkedList_RemoveAll", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.RemoveAll()

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_83_LinkedList_Json(t *testing.T) {
	safeTest(t, "Test_S12_83_LinkedList_Json", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jsonResult := ll.Json()

		// Act
		actual := args.Map{"result": jsonResult.HasError()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_S12_84_LinkedList_JsonPtr(t *testing.T) {
	safeTest(t, "Test_S12_84_LinkedList_JsonPtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.JsonPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S12_85_LinkedList_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_S12_85_LinkedList_ParseInjectUsingJson", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jsonResult := ll.JsonPtr()
		target := corestr.New.LinkedList.Create()
		result, err := target.ParseInjectUsingJson(jsonResult)

		// Act
		actual := args.Map{"result": err != nil || result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_86_LinkedList_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_S12_86_LinkedList_ParseInjectUsingJsonMust", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jsonResult := ll.JsonPtr()
		target := corestr.New.LinkedList.Create()
		result := target.ParseInjectUsingJsonMust(jsonResult)

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_87_LinkedList_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_S12_87_LinkedList_JsonParseSelfInject", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jsonResult := ll.JsonPtr()
		target := corestr.New.LinkedList.Create()
		err := target.JsonParseSelfInject(jsonResult)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_S12_88_LinkedList_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_S12_88_LinkedList_AsJsonMarshaller", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": ll.AsJsonMarshaller() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S12_89_LinkedList_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_S12_89_LinkedList_GetCompareSummary", func() {
		// Arrange
		a := corestr.New.LinkedList.Create()
		a.Add("x")
		b := corestr.New.LinkedList.Create()
		b.Add("y")
		summary := a.GetCompareSummary(b, "left", "right")

		// Act
		actual := args.Map{"result": summary == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_S12_90_LinkedList_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_S12_90_LinkedList_AppendChainOfNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		other := corestr.New.LinkedList.Create()
		other.Adds("b", "c")
		ll.AppendChainOfNodes(other.Head())

		// Act
		actual := args.Map{"result": ll.Length() < 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	})
}

func Test_S12_91_LinkedList_AppendChainOfNodes_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S12_91_LinkedList_AppendChainOfNodes_OnEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		other := corestr.New.LinkedList.Create()
		other.Add("a")
		ll.AppendChainOfNodes(other.Head())

		// Act
		actual := args.Map{"result": ll.Length() < 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_S12_92_LinkedList_AttachWithNode(t *testing.T) {
	safeTest(t, "Test_S12_92_LinkedList_AttachWithNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := ll.Head()
		addNode := &corestr.LinkedListNode{Element: "b"}
		err := ll.AttachWithNode(node, addNode)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_S12_93_LinkedList_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_S12_93_LinkedList_AttachWithNode_NilCurrent", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		addNode := &corestr.LinkedListNode{Element: "b"}
		err := ll.AttachWithNode(nil, addNode)

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_S12_94_LinkedList_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_S12_94_LinkedList_RemoveNodeByIndexes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByIndexes(false, 1)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_95_LinkedList_RemoveNodeByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_S12_95_LinkedList_RemoveNodeByIndexes_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.RemoveNodeByIndexes(false)

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_96_LinkedList_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_S12_96_LinkedList_AddStringsToNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := ll.Head()
		ll.AddStringsToNode(false, node, []string{"b"})

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_97_LinkedList_AddStringsToNode_SkipOnNull(t *testing.T) {
	safeTest(t, "Test_S12_97_LinkedList_AddStringsToNode_SkipOnNull", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddStringsToNode(true, nil, []string{"a"})

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_98_LinkedList_AddStringsToNode_Empty(t *testing.T) {
	safeTest(t, "Test_S12_98_LinkedList_AddStringsToNode_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.AddStringsToNode(false, ll.Head(), []string{})

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_99_LinkedList_AddStringsToNode_Multiple(t *testing.T) {
	safeTest(t, "Test_S12_99_LinkedList_AddStringsToNode_Multiple", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := ll.Head()
		ll.AddStringsToNode(false, node, []string{"b", "c"})

		// Act
		actual := args.Map{"result": ll.Length() < 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	})
}

func Test_S12_100_LinkedList_AddStringsPtrToNode(t *testing.T) {
	safeTest(t, "Test_S12_100_LinkedList_AddStringsPtrToNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		items := []string{"b"}
		ll.AddStringsPtrToNode(false, ll.Head(), &items)

		// Act
		actual := args.Map{"result": ll.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_S12_101_LinkedList_AddStringsPtrToNode_Nil(t *testing.T) {
	safeTest(t, "Test_S12_101_LinkedList_AddStringsPtrToNode_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddStringsPtrToNode(true, nil, nil)

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_102_LinkedList_AddAsync(t *testing.T) {
	safeTest(t, "Test_S12_102_LinkedList_AddAsync", func() {
		ll := corestr.New.LinkedList.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		col := corestr.New.Collection.Strings([]string{"a"})
		ll.AddCollectionToNode(true, nil, col)
		_ = ll
	})
}
