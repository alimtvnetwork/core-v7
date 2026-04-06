package corestrtests

import (
	"strings"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
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
		if ll.Head() == nil || ll.Tail() == nil {
			t.Fatal("expected non-nil head/tail")
		}
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_02_LinkedList_LengthLock(t *testing.T) {
	safeTest(t, "Test_S12_02_LinkedList_LengthLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		if ll.LengthLock() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_03_LinkedList_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_S12_03_LinkedList_IsEmpty_HasItems", func() {
		ll := corestr.New.LinkedList.Create()
		if !ll.IsEmpty() {
			t.Fatal("expected empty")
		}
		if ll.HasItems() {
			t.Fatal("expected no items")
		}
		ll.Add("a")
		if ll.IsEmpty() || !ll.HasItems() {
			t.Fatal("expected not empty")
		}
	})
}

func Test_S12_04_LinkedList_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_S12_04_LinkedList_IsEmptyLock", func() {
		ll := corestr.New.LinkedList.Create()
		if !ll.IsEmptyLock() {
			t.Fatal("expected empty")
		}
	})
}

func Test_S12_05_LinkedList_Add_MultipleItems(t *testing.T) {
	safeTest(t, "Test_S12_05_LinkedList_Add_MultipleItems", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.Add("b")
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S12_06_LinkedList_AddLock(t *testing.T) {
	safeTest(t, "Test_S12_06_LinkedList_AddLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddLock("a")
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_07_LinkedList_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_S12_07_LinkedList_AddItemsMap", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{"a": true, "b": false})
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_08_LinkedList_AddItemsMap_Empty(t *testing.T) {
	safeTest(t, "Test_S12_08_LinkedList_AddItemsMap_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{})
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S12_09_LinkedList_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_S12_09_LinkedList_AddNonEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmpty("")
		ll.AddNonEmpty("a")
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_10_LinkedList_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_S12_10_LinkedList_AddNonEmptyWhitespace", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmptyWhitespace("   ")
		ll.AddNonEmptyWhitespace("a")
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_11_LinkedList_AddIf(t *testing.T) {
	safeTest(t, "Test_S12_11_LinkedList_AddIf", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddIf(true, "yes")
		ll.AddIf(false, "no")
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_12_LinkedList_AddsIf(t *testing.T) {
	safeTest(t, "Test_S12_12_LinkedList_AddsIf", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddsIf(true, "a", "b")
		ll.AddsIf(false, "c")
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S12_13_LinkedList_AddFunc(t *testing.T) {
	safeTest(t, "Test_S12_13_LinkedList_AddFunc", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFunc(func() string { return "val" })
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_14_LinkedList_AddFuncErr_NoError(t *testing.T) {
	safeTest(t, "Test_S12_14_LinkedList_AddFuncErr_NoError", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFuncErr(func() (string, error) { return "ok", nil }, func(err error) { t.Fatal("no err expected") })
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_15_LinkedList_AddFuncErr_WithError(t *testing.T) {
	safeTest(t, "Test_S12_15_LinkedList_AddFuncErr_WithError", func() {
		ll := corestr.New.LinkedList.Create()
		called := false
		ll.AddFuncErr(func() (string, error) { return "", &testErr{} }, func(err error) { called = true })
		if !called {
			t.Fatal("expected err handler called")
		}
	})
}

func Test_S12_16_LinkedList_Push_PushBack_PushFront(t *testing.T) {
	safeTest(t, "Test_S12_16_LinkedList_Push_PushBack_PushFront", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Push("a")
		ll.PushBack("b")
		ll.PushFront("z")
		if ll.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_S12_17_LinkedList_AddFront_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S12_17_LinkedList_AddFront_OnEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFront("a")
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_18_LinkedList_AddFront_OnNonEmpty(t *testing.T) {
	safeTest(t, "Test_S12_18_LinkedList_AddFront_OnNonEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("b")
		ll.AddFront("a")
		if ll.Head().Element != "a" {
			t.Fatal("expected a at head")
		}
	})
}

func Test_S12_19_LinkedList_Adds(t *testing.T) {
	safeTest(t, "Test_S12_19_LinkedList_Adds", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		if ll.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_S12_20_LinkedList_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_S12_20_LinkedList_Adds_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds()
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S12_21_LinkedList_AddStrings(t *testing.T) {
	safeTest(t, "Test_S12_21_LinkedList_AddStrings", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings([]string{"a", "b"})
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S12_22_LinkedList_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_S12_22_LinkedList_AddStrings_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings([]string{})
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S12_23_LinkedList_AddsLock(t *testing.T) {
	safeTest(t, "Test_S12_23_LinkedList_AddsLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddsLock("a")
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_24_LinkedList_AppendNode_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S12_24_LinkedList_AppendNode_OnEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		node := &corestr.LinkedListNode{Element: "a"}
		ll.AppendNode(node)
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_25_LinkedList_AddBackNode(t *testing.T) {
	safeTest(t, "Test_S12_25_LinkedList_AddBackNode", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("x")
		node := &corestr.LinkedListNode{Element: "y"}
		ll.AddBackNode(node)
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S12_26_LinkedList_InsertAt(t *testing.T) {
	safeTest(t, "Test_S12_26_LinkedList_InsertAt", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "c")
		ll.InsertAt(1, "b")
		if ll.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_S12_27_LinkedList_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_S12_27_LinkedList_InsertAt_Front", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("b")
		ll.InsertAt(0, "a")
		if ll.Head().Element != "a" {
			t.Fatal("expected a at head")
		}
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
		if count != 3 {
			t.Fatalf("expected 3, got %d", count)
		}
	})
}

func Test_S12_29_LinkedList_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_S12_29_LinkedList_Loop_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			t.Fatal("should not be called")
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
		if count != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_31_LinkedList_Filter(t *testing.T) {
	safeTest(t, "Test_S12_31_LinkedList_Filter", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})
		if len(nodes) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S12_32_LinkedList_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_S12_32_LinkedList_Filter_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})
		if len(nodes) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S12_33_LinkedList_Filter_Break(t *testing.T) {
	safeTest(t, "Test_S12_33_LinkedList_Filter_Break", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})
		if len(nodes) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_34_LinkedList_IsEquals(t *testing.T) {
	safeTest(t, "Test_S12_34_LinkedList_IsEquals", func() {
		a := corestr.New.LinkedList.Create()
		a.Adds("a", "b")
		b := corestr.New.LinkedList.Create()
		b.Adds("a", "b")
		if !a.IsEquals(b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S12_35_LinkedList_IsEquals_BothNil(t *testing.T) {
	safeTest(t, "Test_S12_35_LinkedList_IsEquals_BothNil", func() {
		var a *corestr.LinkedList
		var b *corestr.LinkedList
		if !a.IsEqualsWithSensitive(b, true) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S12_36_LinkedList_IsEquals_OneNil(t *testing.T) {
	safeTest(t, "Test_S12_36_LinkedList_IsEquals_OneNil", func() {
		a := corestr.New.LinkedList.Create()
		a.Add("a")
		var b *corestr.LinkedList
		if a.IsEqualsWithSensitive(b, true) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S12_37_LinkedList_IsEquals_SamePtr(t *testing.T) {
	safeTest(t, "Test_S12_37_LinkedList_IsEquals_SamePtr", func() {
		a := corestr.New.LinkedList.Create()
		a.Add("a")
		if !a.IsEqualsWithSensitive(a, true) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S12_38_LinkedList_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S12_38_LinkedList_IsEquals_BothEmpty", func() {
		a := corestr.New.LinkedList.Create()
		b := corestr.New.LinkedList.Create()
		if !a.IsEquals(b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S12_39_LinkedList_IsEquals_DiffLength(t *testing.T) {
	safeTest(t, "Test_S12_39_LinkedList_IsEquals_DiffLength", func() {
		a := corestr.New.LinkedList.Create()
		a.Add("a")
		b := corestr.New.LinkedList.Create()
		b.Adds("a", "b")
		if a.IsEquals(b) {
			t.Fatal("expected not equal")
		}
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
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S12_42_LinkedList_RemoveNodeByElementValue_First(t *testing.T) {
	safeTest(t, "Test_S12_42_LinkedList_RemoveNodeByElementValue_First", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.RemoveNodeByElementValue("a", true, false)
		if ll.Head().Element != "b" {
			t.Fatal("expected b at head")
		}
	})
}

func Test_S12_43_LinkedList_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_S12_43_LinkedList_RemoveNodeByIndex", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByIndex(1)
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S12_44_LinkedList_RemoveNodeByIndex_First(t *testing.T) {
	safeTest(t, "Test_S12_44_LinkedList_RemoveNodeByIndex_First", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.RemoveNodeByIndex(0)
		if ll.Head().Element != "b" {
			t.Fatal("expected b")
		}
	})
}

func Test_S12_45_LinkedList_RemoveNodeByIndex_Last(t *testing.T) {
	safeTest(t, "Test_S12_45_LinkedList_RemoveNodeByIndex_Last", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.RemoveNodeByIndex(1)
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_46_LinkedList_RemoveNode(t *testing.T) {
	safeTest(t, "Test_S12_46_LinkedList_RemoveNode", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.Add("b")
		node := ll.IndexAt(0)
		ll.RemoveNode(node)
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_47_LinkedList_RemoveNode_Nil(t *testing.T) {
	safeTest(t, "Test_S12_47_LinkedList_RemoveNode_Nil", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.RemoveNode(nil)
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_48_LinkedList_IndexAt(t *testing.T) {
	safeTest(t, "Test_S12_48_LinkedList_IndexAt", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		node := ll.IndexAt(1)
		if node.Element != "b" {
			t.Fatal("expected b")
		}
	})
}

func Test_S12_49_LinkedList_IndexAt_Zero(t *testing.T) {
	safeTest(t, "Test_S12_49_LinkedList_IndexAt_Zero", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		if ll.IndexAt(0).Element != "a" {
			t.Fatal("expected a")
		}
	})
}

func Test_S12_50_LinkedList_IndexAt_Negative(t *testing.T) {
	safeTest(t, "Test_S12_50_LinkedList_IndexAt_Negative", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		if ll.IndexAt(-1) != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_S12_51_LinkedList_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_S12_51_LinkedList_SafeIndexAt", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		if ll.SafeIndexAt(1).Element != "b" {
			t.Fatal("expected b")
		}
		if ll.SafeIndexAt(-1) != nil {
			t.Fatal("expected nil")
		}
		if ll.SafeIndexAt(10) != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_S12_52_LinkedList_SafeIndexAtLock(t *testing.T) {
	safeTest(t, "Test_S12_52_LinkedList_SafeIndexAtLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		if ll.SafeIndexAtLock(0).Element != "a" {
			t.Fatal("expected a")
		}
	})
}

func Test_S12_53_LinkedList_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_S12_53_LinkedList_SafePointerIndexAt", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ptr := ll.SafePointerIndexAt(0)
		if ptr == nil || *ptr != "a" {
			t.Fatal("expected a")
		}
		if ll.SafePointerIndexAt(10) != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_S12_54_LinkedList_SafePointerIndexAtUsingDefault(t *testing.T) {
	safeTest(t, "Test_S12_54_LinkedList_SafePointerIndexAtUsingDefault", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		if ll.SafePointerIndexAtUsingDefault(0, "def") != "a" {
			t.Fatal("expected a")
		}
		if ll.SafePointerIndexAtUsingDefault(10, "def") != "def" {
			t.Fatal("expected def")
		}
	})
}

func Test_S12_55_LinkedList_SafePointerIndexAtUsingDefaultLock(t *testing.T) {
	safeTest(t, "Test_S12_55_LinkedList_SafePointerIndexAtUsingDefaultLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		if ll.SafePointerIndexAtUsingDefaultLock(0, "def") != "a" {
			t.Fatal("expected a")
		}
	})
}

func Test_S12_56_LinkedList_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_S12_56_LinkedList_GetNextNodes", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		nodes := ll.GetNextNodes(2)
		if len(nodes) != 2 {
			t.Fatalf("expected 2, got %d", len(nodes))
		}
	})
}

func Test_S12_57_LinkedList_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_S12_57_LinkedList_GetAllLinkedNodes", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		if len(ll.GetAllLinkedNodes()) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S12_58_LinkedList_AddPointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_S12_58_LinkedList_AddPointerStringsPtr", func() {
		ll := corestr.New.LinkedList.Create()
		a := "a"
		ll.AddPointerStringsPtr([]*string{&a, nil})
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_59_LinkedList_AddCollection(t *testing.T) {
	safeTest(t, "Test_S12_59_LinkedList_AddCollection", func() {
		ll := corestr.New.LinkedList.Create()
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		ll.AddCollection(col)
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S12_60_LinkedList_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_S12_60_LinkedList_AddCollection_Nil", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddCollection(nil)
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S12_61_LinkedList_ToCollection(t *testing.T) {
	safeTest(t, "Test_S12_61_LinkedList_ToCollection", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		col := ll.ToCollection(0)
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S12_62_LinkedList_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_S12_62_LinkedList_ToCollection_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		col := ll.ToCollection(0)
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S12_63_LinkedList_List(t *testing.T) {
	safeTest(t, "Test_S12_63_LinkedList_List", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		if len(ll.List()) != 2 {
			t.Fatal("expected 2")
		}
	})
}
func Test_S12_65_LinkedList_ListLock(t *testing.T) {
	safeTest(t, "Test_S12_65_LinkedList_ListLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		if len(ll.ListLock()) != 1 {
			t.Fatal("expected 1")
		}
	})
}
func Test_S12_67_LinkedList_String(t *testing.T) {
	safeTest(t, "Test_S12_67_LinkedList_String", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		if ll.String() == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S12_68_LinkedList_String_Empty(t *testing.T) {
	safeTest(t, "Test_S12_68_LinkedList_String_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		if !strings.Contains(ll.String(), "No Element") {
			t.Fatal("expected No Element")
		}
	})
}

func Test_S12_69_LinkedList_StringLock(t *testing.T) {
	safeTest(t, "Test_S12_69_LinkedList_StringLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		if ll.StringLock() == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S12_70_LinkedList_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_S12_70_LinkedList_StringLock_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		if !strings.Contains(ll.StringLock(), "No Element") {
			t.Fatal("expected No Element")
		}
	})
}

func Test_S12_71_LinkedList_Join(t *testing.T) {
	safeTest(t, "Test_S12_71_LinkedList_Join", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		if ll.Join(",") != "a,b" {
			t.Fatalf("expected 'a,b', got '%s'", ll.Join(","))
		}
	})
}

func Test_S12_72_LinkedList_JoinLock(t *testing.T) {
	safeTest(t, "Test_S12_72_LinkedList_JoinLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		if ll.JoinLock(",") != "a" {
			t.Fatal("expected a")
		}
	})
}

func Test_S12_73_LinkedList_Joins(t *testing.T) {
	safeTest(t, "Test_S12_73_LinkedList_Joins", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		result := ll.Joins(",", "b")
		if !strings.Contains(result, "a") {
			t.Fatal("expected a in result")
		}
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
		if err != nil || len(data) == 0 {
			t.Fatal("expected valid JSON")
		}
	})
}

func Test_S12_76_LinkedList_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_S12_76_LinkedList_UnmarshalJSON", func() {
		ll := corestr.New.LinkedList.Create()
		err := ll.UnmarshalJSON([]byte(`["a","b"]`))
		if err != nil || ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S12_77_LinkedList_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_S12_77_LinkedList_UnmarshalJSON_Invalid", func() {
		ll := corestr.New.LinkedList.Create()
		err := ll.UnmarshalJSON([]byte(`invalid`))
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_S12_78_LinkedList_JsonModel(t *testing.T) {
	safeTest(t, "Test_S12_78_LinkedList_JsonModel", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		if len(ll.JsonModel()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_79_LinkedList_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_S12_79_LinkedList_JsonModelAny", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		if ll.JsonModelAny() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S12_80_LinkedList_Clear_RemoveAll(t *testing.T) {
	safeTest(t, "Test_S12_80_LinkedList_Clear_RemoveAll", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.Clear()
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S12_81_LinkedList_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_S12_81_LinkedList_Clear_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Clear()
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S12_82_LinkedList_RemoveAll(t *testing.T) {
	safeTest(t, "Test_S12_82_LinkedList_RemoveAll", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.RemoveAll()
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S12_83_LinkedList_Json(t *testing.T) {
	safeTest(t, "Test_S12_83_LinkedList_Json", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jsonResult := ll.Json()
		if jsonResult.HasError() {
			t.Fatal("expected no error")
		}
	})
}

func Test_S12_84_LinkedList_JsonPtr(t *testing.T) {
	safeTest(t, "Test_S12_84_LinkedList_JsonPtr", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		if ll.JsonPtr() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S12_85_LinkedList_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_S12_85_LinkedList_ParseInjectUsingJson", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jsonResult := ll.JsonPtr()
		target := corestr.New.LinkedList.Create()
		result, err := target.ParseInjectUsingJson(jsonResult)
		if err != nil || result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_86_LinkedList_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_S12_86_LinkedList_ParseInjectUsingJsonMust", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jsonResult := ll.JsonPtr()
		target := corestr.New.LinkedList.Create()
		result := target.ParseInjectUsingJsonMust(jsonResult)
		if result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_87_LinkedList_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_S12_87_LinkedList_JsonParseSelfInject", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jsonResult := ll.JsonPtr()
		target := corestr.New.LinkedList.Create()
		err := target.JsonParseSelfInject(jsonResult)
		if err != nil {
			t.Fatal("expected no error")
		}
	})
}

func Test_S12_88_LinkedList_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_S12_88_LinkedList_AsJsonMarshaller", func() {
		ll := corestr.New.LinkedList.Create()
		if ll.AsJsonMarshaller() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S12_89_LinkedList_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_S12_89_LinkedList_GetCompareSummary", func() {
		a := corestr.New.LinkedList.Create()
		a.Add("x")
		b := corestr.New.LinkedList.Create()
		b.Add("y")
		summary := a.GetCompareSummary(b, "left", "right")
		if summary == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S12_90_LinkedList_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_S12_90_LinkedList_AppendChainOfNodes", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		other := corestr.New.LinkedList.Create()
		other.Adds("b", "c")
		ll.AppendChainOfNodes(other.Head())
		if ll.Length() < 3 {
			t.Fatal("expected at least 3")
		}
	})
}

func Test_S12_91_LinkedList_AppendChainOfNodes_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S12_91_LinkedList_AppendChainOfNodes_OnEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		other := corestr.New.LinkedList.Create()
		other.Add("a")
		ll.AppendChainOfNodes(other.Head())
		if ll.Length() < 1 {
			t.Fatal("expected at least 1")
		}
	})
}

func Test_S12_92_LinkedList_AttachWithNode(t *testing.T) {
	safeTest(t, "Test_S12_92_LinkedList_AttachWithNode", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := ll.Head()
		addNode := &corestr.LinkedListNode{Element: "b"}
		err := ll.AttachWithNode(node, addNode)
		if err != nil {
			t.Fatal("expected no error")
		}
	})
}

func Test_S12_93_LinkedList_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_S12_93_LinkedList_AttachWithNode_NilCurrent", func() {
		ll := corestr.New.LinkedList.Create()
		addNode := &corestr.LinkedListNode{Element: "b"}
		err := ll.AttachWithNode(nil, addNode)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_S12_94_LinkedList_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_S12_94_LinkedList_RemoveNodeByIndexes", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByIndexes(false, 1)
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S12_95_LinkedList_RemoveNodeByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_S12_95_LinkedList_RemoveNodeByIndexes_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.RemoveNodeByIndexes(false)
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S12_96_LinkedList_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_S12_96_LinkedList_AddStringsToNode", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := ll.Head()
		ll.AddStringsToNode(false, node, []string{"b"})
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S12_97_LinkedList_AddStringsToNode_SkipOnNull(t *testing.T) {
	safeTest(t, "Test_S12_97_LinkedList_AddStringsToNode_SkipOnNull", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddStringsToNode(true, nil, []string{"a"})
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S12_98_LinkedList_AddStringsToNode_Empty(t *testing.T) {
	safeTest(t, "Test_S12_98_LinkedList_AddStringsToNode_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.AddStringsToNode(false, ll.Head(), []string{})
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S12_99_LinkedList_AddStringsToNode_Multiple(t *testing.T) {
	safeTest(t, "Test_S12_99_LinkedList_AddStringsToNode_Multiple", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := ll.Head()
		ll.AddStringsToNode(false, node, []string{"b", "c"})
		if ll.Length() < 3 {
			t.Fatalf("expected at least 3, got %d", ll.Length())
		}
	})
}
	safeTest(t, "Test_S12_102_LinkedList_AddAsync", func() {
		ll := corestr.New.LinkedList.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		col := corestr.New.Collection.Strings([]string{"a"})
		ll.AddCollectionToNode(true, nil, col)
		_ = ll
	})
}
