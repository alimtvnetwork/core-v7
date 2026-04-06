package corestrtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// LinkedList — Segment 12: Core methods (L1-600)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovLL1_01_IsEmpty_HasItems_Length(t *testing.T) {
	safeTest(t, "Test_CovLL1_01_IsEmpty_HasItems_Length", func() {
		ll := corestr.Empty.LinkedList()
		if !ll.IsEmpty() {
			t.Fatal("expected empty")
		}
		if ll.HasItems() {
			t.Fatal("expected no items")
		}
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
		ll.Add("a")
		if ll.IsEmpty() {
			t.Fatal("expected not empty")
		}
		if !ll.HasItems() {
			t.Fatal("expected items")
		}
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLL1_02_IsEmptyLock_LengthLock(t *testing.T) {
	safeTest(t, "Test_CovLL1_02_IsEmptyLock_LengthLock", func() {
		ll := corestr.Empty.LinkedList()
		if !ll.IsEmptyLock() {
			t.Fatal("expected empty")
		}
		if ll.LengthLock() != 0 {
			t.Fatal("expected 0")
		}
		ll.Add("x")
		if ll.IsEmptyLock() {
			t.Fatal("expected not empty")
		}
		if ll.LengthLock() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLL1_03_Head_Tail(t *testing.T) {
	safeTest(t, "Test_CovLL1_03_Head_Tail", func() {
		ll := corestr.Empty.LinkedList()
		if ll.Head() != nil {
			t.Fatal("expected nil")
		}
		if ll.Tail() != nil {
			t.Fatal("expected nil")
		}
		ll.Add("a")
		if ll.Head() == nil || ll.Head().Element != "a" {
			t.Fatal("expected a")
		}
		if ll.Tail() == nil || ll.Tail().Element != "a" {
			t.Fatal("expected a")
		}
		ll.Add("b")
		if ll.Head().Element != "a" {
			t.Fatal("expected a")
		}
		if ll.Tail().Element != "b" {
			t.Fatal("expected b")
		}
	})
}

func Test_CovLL1_04_Add_Multiple(t *testing.T) {
	safeTest(t, "Test_CovLL1_04_Add_Multiple", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b").Add("c")
		if ll.Length() != 3 {
			t.Fatal("expected 3")
		}
		items := ll.List()
		if items[0] != "a" || items[1] != "b" || items[2] != "c" {
			t.Fatal("unexpected order")
		}
	})
}

func Test_CovLL1_05_AddLock(t *testing.T) {
	safeTest(t, "Test_CovLL1_05_AddLock", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddLock("a")
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLL1_06_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_CovLL1_06_AddItemsMap", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddItemsMap(map[string]bool{"a": true, "b": false})
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
		// empty
		ll.AddItemsMap(nil)
		ll.AddItemsMap(map[string]bool{})
	})
}

func Test_CovLL1_07_AddFront_PushFront(t *testing.T) {
	safeTest(t, "Test_CovLL1_07_AddFront_PushFront", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("b")
		ll.AddFront("a")
		items := ll.List()
		if items[0] != "a" || items[1] != "b" {
			t.Fatal("unexpected order")
		}
		// AddFront on empty
		ll2 := corestr.Empty.LinkedList()
		ll2.AddFront("x")
		if ll2.Length() != 1 {
			t.Fatal("expected 1")
		}
		// PushFront
		ll2.PushFront("y")
		if ll2.Head().Element != "y" {
			t.Fatal("expected y")
		}
	})
}

func Test_CovLL1_08_Push_PushBack(t *testing.T) {
	safeTest(t, "Test_CovLL1_08_Push_PushBack", func() {
		ll := corestr.Empty.LinkedList()
		ll.Push("a")
		ll.PushBack("b")
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLL1_09_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_CovLL1_09_AddNonEmpty", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddNonEmpty("")
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
		ll.AddNonEmpty("a")
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLL1_10_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_CovLL1_10_AddNonEmptyWhitespace", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddNonEmptyWhitespace("   ")
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
		ll.AddNonEmptyWhitespace("a")
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLL1_11_AddIf(t *testing.T) {
	safeTest(t, "Test_CovLL1_11_AddIf", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddIf(false, "a")
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
		ll.AddIf(true, "a")
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLL1_12_AddsIf(t *testing.T) {
	safeTest(t, "Test_CovLL1_12_AddsIf", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddsIf(false, "a", "b")
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
		ll.AddsIf(true, "a", "b")
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLL1_13_AddFunc(t *testing.T) {
	safeTest(t, "Test_CovLL1_13_AddFunc", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddFunc(func() string { return "hello" })
		if ll.Length() != 1 || ll.Head().Element != "hello" {
			t.Fatal("expected hello")
		}
	})
}

func Test_CovLL1_14_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_CovLL1_14_AddFuncErr", func() {
		ll := corestr.Empty.LinkedList()
		// success
		ll.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(err error) { t.Fatal("unexpected error") },
		)
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
		// error
		ll.AddFuncErr(
			func() (string, error) { return "", fmt.Errorf("fail") },
			func(err error) {},
		)
		if ll.Length() != 1 {
			t.Fatal("expected still 1")
		}
	})
}

func Test_CovLL1_15_Adds_AddStrings(t *testing.T) {
	safeTest(t, "Test_CovLL1_15_Adds_AddStrings", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		if ll.Length() != 3 {
			t.Fatal("expected 3")
		}
		ll.Adds()
		ll2 := corestr.Empty.LinkedList()
		ll2.AddStrings([]string{"x", "y"})
		if ll2.Length() != 2 {
			t.Fatal("expected 2")
		}
		ll2.AddStrings(nil)
	})
}

func Test_CovLL1_16_AddsLock(t *testing.T) {
	safeTest(t, "Test_CovLL1_16_AddsLock", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddsLock("a", "b")
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLL1_17_InsertAt(t *testing.T) {
	safeTest(t, "Test_CovLL1_17_InsertAt", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "c")
		ll.InsertAt(1, "b")
		items := ll.List()
		if len(items) < 3 {
			t.Fatal("expected 3 items")
		}
		// index < 1 → AddFront
		ll.InsertAt(-1, "z")
		if ll.Head().Element != "z" {
			t.Fatal("expected z at front")
		}
	})
}

func Test_CovLL1_18_AppendNode_AddBackNode(t *testing.T) {
	safeTest(t, "Test_CovLL1_18_AppendNode_AddBackNode", func() {
		ll := corestr.Empty.LinkedList()
		node := &corestr.LinkedListNode{Element: "a"}
		ll.AppendNode(node)
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
		// not empty
		node2 := &corestr.LinkedListNode{Element: "b"}
		ll.AddBackNode(node2)
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLL1_19_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_CovLL1_19_AppendChainOfNodes", func() {
		ll := corestr.Empty.LinkedList()
		// build chain via another list
		chain := corestr.Empty.LinkedList()
		chain.Adds("a", "b", "c")
		ll.AppendChainOfNodes(chain.Head())
		if ll.Length() != 3 {
			t.Fatal("expected 3")
		}
		// append to non-empty
		chain2 := corestr.Empty.LinkedList()
		chain2.Adds("d", "e")
		ll.AppendChainOfNodes(chain2.Head())
		if ll.Length() != 5 {
			t.Fatal("expected 5")
		}
	})
}

func Test_CovLL1_20_AddPointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_CovLL1_20_AddPointerStringsPtr", func() {
		ll := corestr.Empty.LinkedList()
		a := "a"
		ll.AddPointerStringsPtr([]*string{&a, nil})
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLL1_21_AddCollection(t *testing.T) {
	safeTest(t, "Test_CovLL1_21_AddCollection", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddCollection(nil)
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		ll.AddCollection(col)
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLL1_22_AttachWithNode(t *testing.T) {
	safeTest(t, "Test_CovLL1_22_AttachWithNode", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		node := ll.Head()
		addNode := &corestr.LinkedListNode{Element: "b"}
		err := ll.AttachWithNode(node, addNode)
		if err != nil {
			t.Fatal("unexpected error")
		}
		// nil current
		err2 := ll.AttachWithNode(nil, addNode)
		if err2 == nil {
			t.Fatal("expected error")
		}
		// current.next not nil
		err3 := ll.AttachWithNode(node, addNode)
		if err3 == nil {
			t.Fatal("expected error for non-nil next")
		}
	})
}

func Test_CovLL1_23_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_CovLL1_23_AddStringsToNode", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		node := ll.Head()
		ll.AddStringsToNode(false, node, []string{"b", "c"})
		if ll.Length() < 3 {
			t.Fatal("expected at least 3")
		}
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
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		col := corestr.New.Collection.Strings([]string{"b", "c"})
		ll.AddCollectionToNode(true, ll.Head(), col)
		if ll.Length() < 3 {
			t.Fatal("expected at least 3")
		}
	})
}

func Test_CovLL1_26_Loop(t *testing.T) {
	safeTest(t, "Test_CovLL1_26_Loop", func() {
		ll := corestr.Empty.LinkedList()
		// empty
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})
		if count != 0 {
			t.Fatal("expected 0 iterations")
		}
		// with items
		ll.Adds("a", "b", "c")
		count = 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})
		if count != 3 {
			t.Fatal("expected 3")
		}
		// break
		count = 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return true
		})
		if count != 1 {
			t.Fatal("expected 1")
		}
		// break on second
		count = 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return arg.Index == 1
		})
		if count != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLL1_27_Filter(t *testing.T) {
	safeTest(t, "Test_CovLL1_27_Filter", func() {
		ll := corestr.Empty.LinkedList()
		// empty
		r := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true}
		})
		if len(r) != 0 {
			t.Fatal("expected 0")
		}
		ll.Adds("a", "b", "c")
		// keep all
		r2 := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true}
		})
		if len(r2) != 3 {
			t.Fatal("expected 3")
		}
		// break on first
		r3 := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})
		if len(r3) != 1 {
			t.Fatal("expected 1")
		}
		// skip all
		r4 := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: false}
		})
		if len(r4) != 0 {
			t.Fatal("expected 0")
		}
		// break on loop iteration
		r5 := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: arg.Index == 1}
		})
		if len(r5) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLL1_28_RemoveNodeByElementValue(t *testing.T) {
	safeTest(t, "Test_CovLL1_28_RemoveNodeByElementValue", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByElementValue("a", true, false)
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
		// case insensitive
		ll.RemoveNodeByElementValue("B", false, false)
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
		// not found
		ll.RemoveNodeByElementValue("z", true, true)
		// remove non-first
		ll2 := corestr.Empty.LinkedList()
		ll2.Adds("x", "y", "z")
		ll2.RemoveNodeByElementValue("y", true, false)
		if ll2.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLL1_29_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_CovLL1_29_RemoveNodeByIndex", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		// remove first
		ll.RemoveNodeByIndex(0)
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
		// remove last
		ll.RemoveNodeByIndex(1)
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
		// remove middle
		ll2 := corestr.Empty.LinkedList()
		ll2.Adds("a", "b", "c")
		ll2.RemoveNodeByIndex(1)
		if ll2.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLL1_30_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_CovLL1_30_RemoveNodeByIndexes", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c", "d")
		ll.RemoveNodeByIndexes(false, 1, 3)
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
		// empty indexes
		ll.RemoveNodeByIndexes(false)
		// ignore panic on empty
		empty := corestr.Empty.LinkedList()
		empty.RemoveNodeByIndexes(true, 0)
	})
}

func Test_CovLL1_31_RemoveNode(t *testing.T) {
	safeTest(t, "Test_CovLL1_31_RemoveNode", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		// nil → skip
		ll.RemoveNode(nil)
		if ll.Length() != 3 {
			t.Fatal("expected 3")
		}
		// remove head
		ll.RemoveNode(ll.Head())
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
		// remove non-head
		node := ll.Head().Next()
		ll.RemoveNode(node)
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLL1_32_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_CovLL1_32_GetCompareSummary", func() {
		a := corestr.Empty.LinkedList()
		a.Adds("a", "b")
		b := corestr.Empty.LinkedList()
		b.Adds("a", "b")
		s := a.GetCompareSummary(b, "left", "right")
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_CovLL1_33_IndexAt(t *testing.T) {
	safeTest(t, "Test_CovLL1_33_IndexAt", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		node := ll.IndexAt(0)
		if node.Element != "a" {
			t.Fatal("expected a")
		}
		node2 := ll.IndexAt(2)
		if node2.Element != "c" {
			t.Fatal("expected c")
		}
		// negative
		n := ll.IndexAt(-1)
		if n != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_CovLL1_34_SafeIndexAt_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_CovLL1_34_SafeIndexAt_SafePointerIndexAt", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		// found
		node := ll.SafeIndexAt(0)
		if node == nil || node.Element != "a" {
			t.Fatal("expected a")
		}
		node1 := ll.SafeIndexAt(1)
		if node1 == nil || node1.Element != "b" {
			t.Fatal("expected b")
		}
		// not found
		n := ll.SafeIndexAt(-1)
		if n != nil {
			t.Fatal("expected nil")
		}
		n2 := ll.SafeIndexAt(99)
		if n2 != nil {
			t.Fatal("expected nil")
		}
		// empty
		e := corestr.Empty.LinkedList()
		if e.SafeIndexAt(0) != nil {
			t.Fatal("expected nil")
		}
		// pointer
		p := ll.SafePointerIndexAt(0)
		if p == nil || *p != "a" {
			t.Fatal("expected a")
		}
		p2 := ll.SafePointerIndexAt(-1)
		if p2 != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_CovLL1_35_SafePointerIndexAtUsingDefault(t *testing.T) {
	safeTest(t, "Test_CovLL1_35_SafePointerIndexAtUsingDefault", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		v := ll.SafePointerIndexAtUsingDefault(0, "def")
		if v != "a" {
			t.Fatal("expected a")
		}
		v2 := ll.SafePointerIndexAtUsingDefault(99, "def")
		if v2 != "def" {
			t.Fatal("expected def")
		}
	})
}

func Test_CovLL1_36_SafeIndexAtLock_SafePointerIndexAtUsingDefaultLock(t *testing.T) {
	safeTest(t, "Test_CovLL1_36_SafeIndexAtLock_SafePointerIndexAtUsingDefaultLock", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		n := ll.SafeIndexAtLock(0)
		if n == nil {
			t.Fatal("expected non-nil")
		}
		v := ll.SafePointerIndexAtUsingDefaultLock(0, "def")
		if v != "a" {
			t.Fatal("expected a")
		}
	})
}

func Test_CovLL1_37_GetNextNodes_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_CovLL1_37_GetNextNodes_GetAllLinkedNodes", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		r := ll.GetNextNodes(2)
		if len(r) != 2 {
			t.Fatal("expected 2")
		}
		all := ll.GetAllLinkedNodes()
		if len(all) != 3 {
			t.Fatal("expected 3")
		}
	})
}
