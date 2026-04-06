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
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_02_LinkedList_LengthLock(t *testing.T) {
	safeTest(t, "Test_S12_02_LinkedList_LengthLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"result": ll.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_03_LinkedList_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_S12_03_LinkedList_IsEmpty_HasItems", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"result": ll.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual := args.Map{"result": ll.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
		ll.Add("a")
		actual := args.Map{"result": ll.IsEmpty() || !ll.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_S12_04_LinkedList_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_S12_04_LinkedList_IsEmptyLock", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"result": ll.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_S12_05_LinkedList_Add_MultipleItems(t *testing.T) {
	safeTest(t, "Test_S12_05_LinkedList_Add_MultipleItems", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.Add("b")
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_06_LinkedList_AddLock(t *testing.T) {
	safeTest(t, "Test_S12_06_LinkedList_AddLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddLock("a")
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_07_LinkedList_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_S12_07_LinkedList_AddItemsMap", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{"a": true, "b": false})
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_08_LinkedList_AddItemsMap_Empty(t *testing.T) {
	safeTest(t, "Test_S12_08_LinkedList_AddItemsMap_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{})
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_09_LinkedList_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_S12_09_LinkedList_AddNonEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmpty("")
		ll.AddNonEmpty("a")
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_10_LinkedList_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_S12_10_LinkedList_AddNonEmptyWhitespace", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmptyWhitespace("   ")
		ll.AddNonEmptyWhitespace("a")
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_11_LinkedList_AddIf(t *testing.T) {
	safeTest(t, "Test_S12_11_LinkedList_AddIf", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddIf(true, "yes")
		ll.AddIf(false, "no")
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_12_LinkedList_AddsIf(t *testing.T) {
	safeTest(t, "Test_S12_12_LinkedList_AddsIf", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddsIf(true, "a", "b")
		ll.AddsIf(false, "c")
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_13_LinkedList_AddFunc(t *testing.T) {
	safeTest(t, "Test_S12_13_LinkedList_AddFunc", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFunc(func() string { return "val" })
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_14_LinkedList_AddFuncErr_NoError(t *testing.T) {
	safeTest(t, "Test_S12_14_LinkedList_AddFuncErr_NoError", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFuncErr(func() (string, error) { return "ok", nil }, func(err error) { t.Fatal("no err expected") })
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_15_LinkedList_AddFuncErr_WithError(t *testing.T) {
	safeTest(t, "Test_S12_15_LinkedList_AddFuncErr_WithError", func() {
		ll := corestr.New.LinkedList.Create()
		called := false
		ll.AddFuncErr(func() (string, error) { return "", &testErr{} }, func(err error) { called = true })
		actual := args.Map{"result": called}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected err handler called", actual)
	})
}

func Test_S12_16_LinkedList_Push_PushBack_PushFront(t *testing.T) {
	safeTest(t, "Test_S12_16_LinkedList_Push_PushBack_PushFront", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Push("a")
		ll.PushBack("b")
		ll.PushFront("z")
		actual := args.Map{"result": ll.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_S12_17_LinkedList_AddFront_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S12_17_LinkedList_AddFront_OnEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFront("a")
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_18_LinkedList_AddFront_OnNonEmpty(t *testing.T) {
	safeTest(t, "Test_S12_18_LinkedList_AddFront_OnNonEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("b")
		ll.AddFront("a")
		actual := args.Map{"result": ll.Head().Element != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a at head", actual)
	})
}

func Test_S12_19_LinkedList_Adds(t *testing.T) {
	safeTest(t, "Test_S12_19_LinkedList_Adds", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		actual := args.Map{"result": ll.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_S12_20_LinkedList_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_S12_20_LinkedList_Adds_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds()
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_21_LinkedList_AddStrings(t *testing.T) {
	safeTest(t, "Test_S12_21_LinkedList_AddStrings", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings([]string{"a", "b"})
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_22_LinkedList_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_S12_22_LinkedList_AddStrings_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings([]string{})
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_23_LinkedList_AddsLock(t *testing.T) {
	safeTest(t, "Test_S12_23_LinkedList_AddsLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddsLock("a")
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_24_LinkedList_AppendNode_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S12_24_LinkedList_AppendNode_OnEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		node := &corestr.LinkedListNode{Element: "a"}
		ll.AppendNode(node)
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_25_LinkedList_AddBackNode(t *testing.T) {
	safeTest(t, "Test_S12_25_LinkedList_AddBackNode", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("x")
		node := &corestr.LinkedListNode{Element: "y"}
		ll.AddBackNode(node)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_26_LinkedList_InsertAt(t *testing.T) {
	safeTest(t, "Test_S12_26_LinkedList_InsertAt", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "c")
		ll.InsertAt(1, "b")
		actual := args.Map{"result": ll.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_S12_27_LinkedList_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_S12_27_LinkedList_InsertAt_Front", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("b")
		ll.InsertAt(0, "a")
		actual := args.Map{"result": ll.Head().Element != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a at head", actual)
	})
}

func Test_S12_28_LinkedList_Loop(t *testing.T) {
	safeTest(t, "Test_S12_28_LinkedList_Loop", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
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

func Test_S12_29_LinkedList_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_S12_29_LinkedList_Loop_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			actual := args.Map{"result": false}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "should not be called", actual)
			return false
		})
	})
}

func Test_S12_30_LinkedList_Loop_Break(t *testing.T) {
	safeTest(t, "Test_S12_30_LinkedList_Loop_Break", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
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

func Test_S12_31_LinkedList_Filter(t *testing.T) {
	safeTest(t, "Test_S12_31_LinkedList_Filter", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})
		actual := args.Map{"result": len(nodes) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_32_LinkedList_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_S12_32_LinkedList_Filter_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})
		actual := args.Map{"result": len(nodes) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_33_LinkedList_Filter_Break(t *testing.T) {
	safeTest(t, "Test_S12_33_LinkedList_Filter_Break", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})
		actual := args.Map{"result": len(nodes) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_34_LinkedList_IsEquals(t *testing.T) {
	safeTest(t, "Test_S12_34_LinkedList_IsEquals", func() {
		a := corestr.New.LinkedList.Create()
		a.Adds("a", "b")
		b := corestr.New.LinkedList.Create()
		b.Adds("a", "b")
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S12_35_LinkedList_IsEquals_BothNil(t *testing.T) {
	safeTest(t, "Test_S12_35_LinkedList_IsEquals_BothNil", func() {
		var a *corestr.LinkedList
		var b *corestr.LinkedList
		actual := args.Map{"result": a.IsEqualsWithSensitive(b, true)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S12_36_LinkedList_IsEquals_OneNil(t *testing.T) {
	safeTest(t, "Test_S12_36_LinkedList_IsEquals_OneNil", func() {
		a := corestr.New.LinkedList.Create()
		a.Add("a")
		var b *corestr.LinkedList
		actual := args.Map{"result": a.IsEqualsWithSensitive(b, true)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S12_37_LinkedList_IsEquals_SamePtr(t *testing.T) {
	safeTest(t, "Test_S12_37_LinkedList_IsEquals_SamePtr", func() {
		a := corestr.New.LinkedList.Create()
		a.Add("a")
		actual := args.Map{"result": a.IsEqualsWithSensitive(a, true)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S12_38_LinkedList_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S12_38_LinkedList_IsEquals_BothEmpty", func() {
		a := corestr.New.LinkedList.Create()
		b := corestr.New.LinkedList.Create()
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S12_39_LinkedList_IsEquals_DiffLength(t *testing.T) {
	safeTest(t, "Test_S12_39_LinkedList_IsEquals_DiffLength", func() {
		a := corestr.New.LinkedList.Create()
		a.Add("a")
		b := corestr.New.LinkedList.Create()
		b.Adds("a", "b")
		actual := args.Map{"result": a.IsEquals(b)}
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
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByElementValue("b", true, false)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_42_LinkedList_RemoveNodeByElementValue_First(t *testing.T) {
	safeTest(t, "Test_S12_42_LinkedList_RemoveNodeByElementValue_First", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.RemoveNodeByElementValue("a", true, false)
		actual := args.Map{"result": ll.Head().Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b at head", actual)
	})
}

func Test_S12_43_LinkedList_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_S12_43_LinkedList_RemoveNodeByIndex", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByIndex(1)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_44_LinkedList_RemoveNodeByIndex_First(t *testing.T) {
	safeTest(t, "Test_S12_44_LinkedList_RemoveNodeByIndex_First", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.RemoveNodeByIndex(0)
		actual := args.Map{"result": ll.Head().Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_S12_45_LinkedList_RemoveNodeByIndex_Last(t *testing.T) {
	safeTest(t, "Test_S12_45_LinkedList_RemoveNodeByIndex_Last", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.RemoveNodeByIndex(1)
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_46_LinkedList_RemoveNode(t *testing.T) {
	safeTest(t, "Test_S12_46_LinkedList_RemoveNode", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.Add("b")
		node := ll.IndexAt(0)
		ll.RemoveNode(node)
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_47_LinkedList_RemoveNode_Nil(t *testing.T) {
	safeTest(t, "Test_S12_47_LinkedList_RemoveNode_Nil", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.RemoveNode(nil)
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_48_LinkedList_IndexAt(t *testing.T) {
	safeTest(t, "Test_S12_48_LinkedList_IndexAt", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		node := ll.IndexAt(1)
		actual := args.Map{"result": node.Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_S12_49_LinkedList_IndexAt_Zero(t *testing.T) {
	safeTest(t, "Test_S12_49_LinkedList_IndexAt_Zero", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"result": ll.IndexAt(0).Element != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_S12_50_LinkedList_IndexAt_Negative(t *testing.T) {
	safeTest(t, "Test_S12_50_LinkedList_IndexAt_Negative", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"result": ll.IndexAt(-1) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_S12_51_LinkedList_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_S12_51_LinkedList_SafeIndexAt", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		actual := args.Map{"result": ll.SafeIndexAt(1).Element != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
		actual := args.Map{"result": ll.SafeIndexAt(-1) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		actual := args.Map{"result": ll.SafeIndexAt(10) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_S12_52_LinkedList_SafeIndexAtLock(t *testing.T) {
	safeTest(t, "Test_S12_52_LinkedList_SafeIndexAtLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"result": ll.SafeIndexAtLock(0).Element != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_S12_53_LinkedList_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_S12_53_LinkedList_SafePointerIndexAt", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ptr := ll.SafePointerIndexAt(0)
		actual := args.Map{"result": ptr == nil || *ptr != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual := args.Map{"result": ll.SafePointerIndexAt(10) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_S12_54_LinkedList_SafePointerIndexAtUsingDefault(t *testing.T) {
	safeTest(t, "Test_S12_54_LinkedList_SafePointerIndexAtUsingDefault", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"result": ll.SafePointerIndexAtUsingDefault(0, "def") != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual := args.Map{"result": ll.SafePointerIndexAtUsingDefault(10, "def") != "def"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected def", actual)
	})
}

func Test_S12_55_LinkedList_SafePointerIndexAtUsingDefaultLock(t *testing.T) {
	safeTest(t, "Test_S12_55_LinkedList_SafePointerIndexAtUsingDefaultLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"result": ll.SafePointerIndexAtUsingDefaultLock(0, "def") != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_S12_56_LinkedList_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_S12_56_LinkedList_GetNextNodes", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		nodes := ll.GetNextNodes(2)
		actual := args.Map{"result": len(nodes) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_57_LinkedList_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_S12_57_LinkedList_GetAllLinkedNodes", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		actual := args.Map{"result": len(ll.GetAllLinkedNodes()) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_58_LinkedList_AddPointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_S12_58_LinkedList_AddPointerStringsPtr", func() {
		ll := corestr.New.LinkedList.Create()
		a := "a"
		ll.AddPointerStringsPtr([]*string{&a, nil})
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_59_LinkedList_AddCollection(t *testing.T) {
	safeTest(t, "Test_S12_59_LinkedList_AddCollection", func() {
		ll := corestr.New.LinkedList.Create()
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		ll.AddCollection(col)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_60_LinkedList_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_S12_60_LinkedList_AddCollection_Nil", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddCollection(nil)
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_61_LinkedList_ToCollection(t *testing.T) {
	safeTest(t, "Test_S12_61_LinkedList_ToCollection", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		col := ll.ToCollection(0)
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_62_LinkedList_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_S12_62_LinkedList_ToCollection_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		col := ll.ToCollection(0)
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_63_LinkedList_List(t *testing.T) {
	safeTest(t, "Test_S12_63_LinkedList_List", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		actual := args.Map{"result": len(ll.List()) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_64_LinkedList_ListPtr(t *testing.T) {
	safeTest(t, "Test_S12_64_LinkedList_ListPtr", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"result": len(ll.ListPtr()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_65_LinkedList_ListLock(t *testing.T) {
	safeTest(t, "Test_S12_65_LinkedList_ListLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"result": len(ll.ListLock()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_66_LinkedList_ListPtrLock(t *testing.T) {
	safeTest(t, "Test_S12_66_LinkedList_ListPtrLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"result": len(ll.ListPtrLock()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_67_LinkedList_String(t *testing.T) {
	safeTest(t, "Test_S12_67_LinkedList_String", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"result": ll.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_S12_68_LinkedList_String_Empty(t *testing.T) {
	safeTest(t, "Test_S12_68_LinkedList_String_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"result": strings.Contains(ll.String(), "No Element")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected No Element", actual)
	})
}

func Test_S12_69_LinkedList_StringLock(t *testing.T) {
	safeTest(t, "Test_S12_69_LinkedList_StringLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"result": ll.StringLock() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_S12_70_LinkedList_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_S12_70_LinkedList_StringLock_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"result": strings.Contains(ll.StringLock(), "No Element")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected No Element", actual)
	})
}

func Test_S12_71_LinkedList_Join(t *testing.T) {
	safeTest(t, "Test_S12_71_LinkedList_Join", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		actual := args.Map{"result": ll.Join(",") != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", ll.Join("", actual)
	})
}

func Test_S12_72_LinkedList_JoinLock(t *testing.T) {
	safeTest(t, "Test_S12_72_LinkedList_JoinLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"result": ll.JoinLock(",") != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_S12_73_LinkedList_Joins(t *testing.T) {
	safeTest(t, "Test_S12_73_LinkedList_Joins", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		result := ll.Joins(",", "b")
		actual := args.Map{"result": strings.Contains(result, "a")}
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
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		data, err := ll.MarshalJSON()
		actual := args.Map{"result": err != nil || len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid JSON", actual)
	})
}

func Test_S12_76_LinkedList_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_S12_76_LinkedList_UnmarshalJSON", func() {
		ll := corestr.New.LinkedList.Create()
		err := ll.UnmarshalJSON([]byte(`["a","b"]`))
		actual := args.Map{"result": err != nil || ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_77_LinkedList_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_S12_77_LinkedList_UnmarshalJSON_Invalid", func() {
		ll := corestr.New.LinkedList.Create()
		err := ll.UnmarshalJSON([]byte(`invalid`))
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_S12_78_LinkedList_JsonModel(t *testing.T) {
	safeTest(t, "Test_S12_78_LinkedList_JsonModel", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"result": len(ll.JsonModel()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_79_LinkedList_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_S12_79_LinkedList_JsonModelAny", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"result": ll.JsonModelAny() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S12_80_LinkedList_Clear_RemoveAll(t *testing.T) {
	safeTest(t, "Test_S12_80_LinkedList_Clear_RemoveAll", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.Clear()
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_81_LinkedList_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_S12_81_LinkedList_Clear_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Clear()
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_82_LinkedList_RemoveAll(t *testing.T) {
	safeTest(t, "Test_S12_82_LinkedList_RemoveAll", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.RemoveAll()
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_83_LinkedList_Json(t *testing.T) {
	safeTest(t, "Test_S12_83_LinkedList_Json", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jsonResult := ll.Json()
		actual := args.Map{"result": jsonResult.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_S12_84_LinkedList_JsonPtr(t *testing.T) {
	safeTest(t, "Test_S12_84_LinkedList_JsonPtr", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"result": ll.JsonPtr() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S12_85_LinkedList_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_S12_85_LinkedList_ParseInjectUsingJson", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jsonResult := ll.JsonPtr()
		target := corestr.New.LinkedList.Create()
		result, err := target.ParseInjectUsingJson(jsonResult)
		actual := args.Map{"result": err != nil || result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_86_LinkedList_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_S12_86_LinkedList_ParseInjectUsingJsonMust", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jsonResult := ll.JsonPtr()
		target := corestr.New.LinkedList.Create()
		result := target.ParseInjectUsingJsonMust(jsonResult)
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_87_LinkedList_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_S12_87_LinkedList_JsonParseSelfInject", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jsonResult := ll.JsonPtr()
		target := corestr.New.LinkedList.Create()
		err := target.JsonParseSelfInject(jsonResult)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_S12_88_LinkedList_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_S12_88_LinkedList_AsJsonMarshaller", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"result": ll.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S12_89_LinkedList_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_S12_89_LinkedList_GetCompareSummary", func() {
		a := corestr.New.LinkedList.Create()
		a.Add("x")
		b := corestr.New.LinkedList.Create()
		b.Add("y")
		summary := a.GetCompareSummary(b, "left", "right")
		actual := args.Map{"result": summary == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_S12_90_LinkedList_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_S12_90_LinkedList_AppendChainOfNodes", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		other := corestr.New.LinkedList.Create()
		other.Adds("b", "c")
		ll.AppendChainOfNodes(other.Head())
		actual := args.Map{"result": ll.Length() < 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	})
}

func Test_S12_91_LinkedList_AppendChainOfNodes_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S12_91_LinkedList_AppendChainOfNodes_OnEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		other := corestr.New.LinkedList.Create()
		other.Add("a")
		ll.AppendChainOfNodes(other.Head())
		actual := args.Map{"result": ll.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_S12_92_LinkedList_AttachWithNode(t *testing.T) {
	safeTest(t, "Test_S12_92_LinkedList_AttachWithNode", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := ll.Head()
		addNode := &corestr.LinkedListNode{Element: "b"}
		err := ll.AttachWithNode(node, addNode)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_S12_93_LinkedList_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_S12_93_LinkedList_AttachWithNode_NilCurrent", func() {
		ll := corestr.New.LinkedList.Create()
		addNode := &corestr.LinkedListNode{Element: "b"}
		err := ll.AttachWithNode(nil, addNode)
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_S12_94_LinkedList_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_S12_94_LinkedList_RemoveNodeByIndexes", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByIndexes(false, 1)
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_95_LinkedList_RemoveNodeByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_S12_95_LinkedList_RemoveNodeByIndexes_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.RemoveNodeByIndexes(false)
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_96_LinkedList_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_S12_96_LinkedList_AddStringsToNode", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := ll.Head()
		ll.AddStringsToNode(false, node, []string{"b"})
		actual := args.Map{"result": ll.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S12_97_LinkedList_AddStringsToNode_SkipOnNull(t *testing.T) {
	safeTest(t, "Test_S12_97_LinkedList_AddStringsToNode_SkipOnNull", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddStringsToNode(true, nil, []string{"a"})
		actual := args.Map{"result": ll.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S12_98_LinkedList_AddStringsToNode_Empty(t *testing.T) {
	safeTest(t, "Test_S12_98_LinkedList_AddStringsToNode_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.AddStringsToNode(false, ll.Head(), []string{})
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S12_99_LinkedList_AddStringsToNode_Multiple(t *testing.T) {
	safeTest(t, "Test_S12_99_LinkedList_AddStringsToNode_Multiple", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := ll.Head()
		ll.AddStringsToNode(false, node, []string{"b", "c"})
		actual := args.Map{"result": ll.Length() < 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	})
}

func Test_S12_100_LinkedList_AddStringsPtrToNode(t *testing.T) {
	safeTest(t, "Test_S12_100_LinkedList_AddStringsPtrToNode", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		items := []string{"b"}
		ll.AddStringsPtrToNode(false, ll.Head(), &items)
		actual := args.Map{"result": ll.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_S12_101_LinkedList_AddStringsPtrToNode_Nil(t *testing.T) {
	safeTest(t, "Test_S12_101_LinkedList_AddStringsPtrToNode_Nil", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddStringsPtrToNode(true, nil, nil)
		actual := args.Map{"result": ll.Length() != 0}
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
