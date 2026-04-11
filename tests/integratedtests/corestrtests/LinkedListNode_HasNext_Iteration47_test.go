package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════
// LinkedListNode
// ═══════════════════════════════════════════════════════════════

func Test_LinkedListNode_HasNext_False(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedListNode_HasNext_False", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		tc := caseV1Compat{Name: "Node HasNext false", Expected: false, Actual: ll.Head().HasNext(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedListNode_HasNext_True(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedListNode_HasNext_True", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b")
		tc := caseV1Compat{Name: "Node HasNext true", Expected: true, Actual: ll.Head().HasNext(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedListNode_Next(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedListNode_Next", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b")
		tc := caseV1Compat{Name: "Node Next", Expected: "b", Actual: ll.Head().Next().Element, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedListNode_EndOfChain(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedListNode_EndOfChain", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b").Add("c")
		end, length := ll.Head().EndOfChain()
		tc := caseV1Compat{Name: "Node EndOfChain elem", Expected: "c", Actual: end.Element, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "Node EndOfChain len", Expected: 3, Actual: length, Args: args.Map{}}
		tc2.ShouldBeEqual(t)
	})
}

func Test_LinkedListNode_Clone(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedListNode_Clone", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("x")
		clone := ll.Head().Clone()
		tc := caseV1Compat{Name: "Node Clone", Expected: "x", Actual: clone.Element, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "Node Clone no next", Expected: false, Actual: clone.HasNext(), Args: args.Map{}}
		tc2.ShouldBeEqual(t)
	})
}

func Test_LinkedListNode_IsEqual_Same(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedListNode_IsEqual_Same", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		tc := caseV1Compat{Name: "Node IsEqual same", Expected: true, Actual: ll.Head().IsEqual(ll.Head()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedListNode_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedListNode_IsEqual_BothNil", func() {
		var n *corestr.LinkedListNode
		tc := caseV1Compat{Name: "Node IsEqual both nil", Expected: true, Actual: n.IsEqual(nil), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedListNode_IsEqualValue(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedListNode_IsEqualValue", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("hello")
		tc := caseV1Compat{Name: "Node IsEqualValue", Expected: true, Actual: ll.Head().IsEqualValue("hello"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedListNode_IsEqualValueSensitive_CI(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedListNode_IsEqualValueSensitive_CI", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("Hello")
		tc := caseV1Compat{Name: "Node IsEqualValueSensitive CI", Expected: true, Actual: ll.Head().IsEqualValueSensitive("hello", false), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedListNode_String(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedListNode_String", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("test")
		tc := caseV1Compat{Name: "Node String", Expected: "test", Actual: ll.Head().String(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedListNode_List(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedListNode_List", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b")
		tc := caseV1Compat{Name: "Node List", Expected: 2, Actual: len(ll.Head().List()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedListNode_Join(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedListNode_Join", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b")
		tc := caseV1Compat{Name: "Node Join", Expected: "a,b", Actual: ll.Head().Join(","), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedListNode_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedListNode_CreateLinkedList", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("x").Add("y")
		newLL := ll.Head().CreateLinkedList()
		tc := caseV1Compat{Name: "Node CreateLinkedList", Expected: 2, Actual: newLL.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedListNode_IsChainEqual_True(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedListNode_IsChainEqual_True", func() {
		ll1 := corestr.Empty.LinkedList()
		ll1.Add("a").Add("b")
		ll2 := corestr.Empty.LinkedList()
		ll2.Add("a").Add("b")
		tc := caseV1Compat{Name: "Node IsChainEqual true", Expected: true, Actual: ll1.Head().IsChainEqual(ll2.Head(), true), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedListNode_IsChainEqual_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedListNode_IsChainEqual_CaseInsensitive", func() {
		ll1 := corestr.Empty.LinkedList()
		ll1.Add("A").Add("B")
		ll2 := corestr.Empty.LinkedList()
		ll2.Add("a").Add("b")
		tc := caseV1Compat{Name: "Node IsChainEqual CI", Expected: true, Actual: ll1.Head().IsChainEqual(ll2.Head(), false), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// LinkedList
// ═══════════════════════════════════════════════════════════════

func Test_LinkedList_Add(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_Add", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		tc := caseV1Compat{Name: "LL Add", Expected: 1, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_Head(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_Head", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("first")
		tc := caseV1Compat{Name: "LL Head", Expected: "first", Actual: ll.Head().Element, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_Tail(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_Tail", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b")
		tc := caseV1Compat{Name: "LL Tail", Expected: "b", Actual: ll.Tail().Element, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_IsEmpty_True(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_IsEmpty_True", func() {
		ll := corestr.Empty.LinkedList()
		tc := caseV1Compat{Name: "LL IsEmpty true", Expected: true, Actual: ll.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_IsEmpty_False(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_IsEmpty_False", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		tc := caseV1Compat{Name: "LL IsEmpty false", Expected: false, Actual: ll.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_HasItems(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_HasItems", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		tc := caseV1Compat{Name: "LL HasItems", Expected: true, Actual: ll.HasItems(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_IsEmptyLock", func() {
		ll := corestr.Empty.LinkedList()
		tc := caseV1Compat{Name: "LL IsEmptyLock", Expected: true, Actual: ll.IsEmptyLock(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_LengthLock", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		tc := caseV1Compat{Name: "LL LengthLock", Expected: 1, Actual: ll.LengthLock(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AddLock", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddLock("x")
		tc := caseV1Compat{Name: "LL AddLock", Expected: 1, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AddFront_FromLinkedListNodeHasNex(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AddFront", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("b").AddFront("a")
		tc := caseV1Compat{Name: "LL AddFront", Expected: "a", Actual: ll.Head().Element, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AddFront_Empty(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AddFront_Empty", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddFront("a")
		tc := caseV1Compat{Name: "LL AddFront empty", Expected: "a", Actual: ll.Head().Element, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_PushFront(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_PushFront", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("b").PushFront("a")
		tc := caseV1Compat{Name: "LL PushFront", Expected: "a", Actual: ll.Head().Element, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_PushBack(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_PushBack", func() {
		ll := corestr.Empty.LinkedList()
		ll.PushBack("a")
		tc := caseV1Compat{Name: "LL PushBack", Expected: 1, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_Push(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_Push", func() {
		ll := corestr.Empty.LinkedList()
		ll.Push("a")
		tc := caseV1Compat{Name: "LL Push", Expected: 1, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AddNonEmpty", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddNonEmpty("a").AddNonEmpty("")
		tc := caseV1Compat{Name: "LL AddNonEmpty", Expected: 1, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AddNonEmptyWhitespace", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddNonEmptyWhitespace("a").AddNonEmptyWhitespace("  ")
		tc := caseV1Compat{Name: "LL AddNonEmptyWhitespace", Expected: 1, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AddIf_True(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AddIf_True", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddIf(true, "a")
		tc := caseV1Compat{Name: "LL AddIf true", Expected: 1, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AddIf_False(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AddIf_False", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddIf(false, "a")
		tc := caseV1Compat{Name: "LL AddIf false", Expected: 0, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AddsIf_True(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AddsIf_True", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddsIf(true, "a", "b")
		tc := caseV1Compat{Name: "LL AddsIf true", Expected: 2, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AddsIf_False(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AddsIf_False", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddsIf(false, "a", "b")
		tc := caseV1Compat{Name: "LL AddsIf false", Expected: 0, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AddFunc(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AddFunc", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddFunc(func() string { return "x" })
		tc := caseV1Compat{Name: "LL AddFunc", Expected: "x", Actual: ll.Head().Element, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AddFuncErr_NoErr(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AddFuncErr_NoErr", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddFuncErr(func() (string, error) { return "ok", nil }, func(err error) {})
		tc := caseV1Compat{Name: "LL AddFuncErr no err", Expected: "ok", Actual: ll.Head().Element, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_Adds(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_Adds", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		tc := caseV1Compat{Name: "LL Adds", Expected: 3, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_Adds_Empty", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds()
		tc := caseV1Compat{Name: "LL Adds empty", Expected: 0, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AddStrings", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddStrings([]string{"x", "y"})
		tc := caseV1Compat{Name: "LL AddStrings", Expected: 2, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AddsLock(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AddsLock", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddsLock("a", "b")
		tc := caseV1Compat{Name: "LL AddsLock", Expected: 2, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AddItemsMap", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddItemsMap(map[string]bool{"a": true, "b": false})
		tc := caseV1Compat{Name: "LL AddItemsMap", Expected: 1, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AddCollection_FromLinkedListNodeHasNex(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AddCollection", func() {
		ll := corestr.Empty.LinkedList()
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		ll.AddCollection(col)
		tc := caseV1Compat{Name: "LL AddCollection", Expected: 2, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AddCollection_Nil", func() {
		ll := corestr.Empty.LinkedList()
		ll.AddCollection(nil)
		tc := caseV1Compat{Name: "LL AddCollection nil", Expected: 0, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AppendNode_FromLinkedListNodeHasNex(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AppendNode", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		node := &corestr.LinkedListNode{Element: "b"}
		ll.AppendNode(node)
		tc := caseV1Compat{Name: "LL AppendNode", Expected: 2, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AppendNode_Empty(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AppendNode_Empty", func() {
		ll := corestr.Empty.LinkedList()
		node := &corestr.LinkedListNode{Element: "a"}
		ll.AppendNode(node)
		tc := caseV1Compat{Name: "LL AppendNode empty", Expected: "a", Actual: ll.Head().Element, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_IsEquals_Same(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_IsEquals_Same", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b")
		tc := caseV1Compat{Name: "LL IsEquals same", Expected: true, Actual: ll.IsEquals(ll), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_IsEquals_Equal(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_IsEquals_Equal", func() {
		ll1 := corestr.Empty.LinkedList()
		ll1.Add("a").Add("b")
		ll2 := corestr.Empty.LinkedList()
		ll2.Add("a").Add("b")
		tc := caseV1Compat{Name: "LL IsEquals equal", Expected: true, Actual: ll1.IsEquals(ll2), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_IsEquals_DiffLen", func() {
		ll1 := corestr.Empty.LinkedList()
		ll1.Add("a")
		ll2 := corestr.Empty.LinkedList()
		ll2.Add("a").Add("b")
		tc := caseV1Compat{Name: "LL IsEquals diff len", Expected: false, Actual: ll1.IsEquals(ll2), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_IsEqualsWithSensitive_CI(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_IsEqualsWithSensitive_CI", func() {
		ll1 := corestr.Empty.LinkedList()
		ll1.Add("A")
		ll2 := corestr.Empty.LinkedList()
		ll2.Add("a")
		tc := caseV1Compat{Name: "LL IsEqualsWithSensitive CI", Expected: true, Actual: ll1.IsEqualsWithSensitive(ll2, false), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_List(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_List", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b")
		tc := caseV1Compat{Name: "LL List", Expected: 2, Actual: len(ll.List()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_List_Empty(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_List_Empty", func() {
		ll := corestr.Empty.LinkedList()
		tc := caseV1Compat{Name: "LL List empty", Expected: 0, Actual: len(ll.List()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_ListLock(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_ListLock", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		tc := caseV1Compat{Name: "LL ListLock", Expected: 1, Actual: len(ll.ListLock()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_String_FromLinkedListNodeHasNex(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_String", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		tc := caseV1Compat{Name: "LL String", Expected: true, Actual: len(ll.String()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_String_Empty(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_String_Empty", func() {
		ll := corestr.Empty.LinkedList()
		tc := caseV1Compat{Name: "LL String empty", Expected: true, Actual: len(ll.String()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_Join(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_Join", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b")
		tc := caseV1Compat{Name: "LL Join", Expected: "a,b", Actual: ll.Join(","), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_JoinLock(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_JoinLock", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b")
		tc := caseV1Compat{Name: "LL JoinLock", Expected: "a,b", Actual: ll.JoinLock(","), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_ToCollection(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_ToCollection", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b")
		col := ll.ToCollection(0)
		tc := caseV1Compat{Name: "LL ToCollection", Expected: 2, Actual: col.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_ToCollection_Empty", func() {
		ll := corestr.Empty.LinkedList()
		col := ll.ToCollection(5)
		tc := caseV1Compat{Name: "LL ToCollection empty", Expected: 0, Actual: col.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_SafeIndexAt_FromLinkedListNodeHasNex(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_SafeIndexAt", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b")
		node := ll.SafeIndexAt(1)
		tc := caseV1Compat{Name: "LL SafeIndexAt", Expected: "b", Actual: node.Element, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_SafeIndexAt_OOB(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_SafeIndexAt_OOB", func() {
		ll := corestr.Empty.LinkedList()
		node := ll.SafeIndexAt(0)
		tc := caseV1Compat{Name: "LL SafeIndexAt oob", Expected: true, Actual: node == nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_SafeIndexAt_Negative(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_SafeIndexAt_Negative", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		node := ll.SafeIndexAt(-1)
		tc := caseV1Compat{Name: "LL SafeIndexAt neg", Expected: true, Actual: node == nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_SafePointerIndexAt", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		ptr := ll.SafePointerIndexAt(0)
		tc := caseV1Compat{Name: "LL SafePointerIndexAt", Expected: "a", Actual: *ptr, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_SafePointerIndexAt_Nil(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_SafePointerIndexAt_Nil", func() {
		ll := corestr.Empty.LinkedList()
		ptr := ll.SafePointerIndexAt(0)
		tc := caseV1Compat{Name: "LL SafePointerIndexAt nil", Expected: true, Actual: ptr == nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_SafePointerIndexAtUsingDefault_FromLinkedListNodeHasNex(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_SafePointerIndexAtUsingDefault", func() {
		ll := corestr.Empty.LinkedList()
		val := ll.SafePointerIndexAtUsingDefault(0, "def")
		tc := caseV1Compat{Name: "LL SafePointerIndexAtDefault", Expected: "def", Actual: val, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_SafeIndexAtLock(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_SafeIndexAtLock", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		node := ll.SafeIndexAtLock(0)
		tc := caseV1Compat{Name: "LL SafeIndexAtLock", Expected: "a", Actual: node.Element, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_Clear_FromLinkedListNodeHasNex(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_Clear", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b")
		ll.Clear()
		tc := caseV1Compat{Name: "LL Clear", Expected: 0, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_RemoveAll(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_RemoveAll", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		ll.RemoveAll()
		tc := caseV1Compat{Name: "LL RemoveAll", Expected: true, Actual: ll.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_JsonModel", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		tc := caseV1Compat{Name: "LL JsonModel", Expected: 1, Actual: len(ll.JsonModel()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_MarshalJSON", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		data, err := ll.MarshalJSON()
		tc := caseV1Compat{Name: "LL MarshalJSON", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_UnmarshalJSON", func() {
		ll := corestr.Empty.LinkedList()
		err := ll.UnmarshalJSON([]byte(`["x","y"]`))
		tc := caseV1Compat{Name: "LL UnmarshalJSON", Expected: true, Actual: err == nil && ll.Length() == 2, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_Loop_FromLinkedListNodeHasNex(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_Loop", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b")
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})
		tc := caseV1Compat{Name: "LL Loop", Expected: 2, Actual: count, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_Loop_Break(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_Loop_Break", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b").Add("c")
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return true
		})
		tc := caseV1Compat{Name: "LL Loop break", Expected: 1, Actual: count, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_Filter_FromLinkedListNodeHasNex(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_Filter", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b").Add("c")
		result := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: arg.Node.Element == "b", IsBreak: false}
		})
		tc := caseV1Compat{Name: "LL Filter", Expected: 1, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_GetAllLinkedNodes_FromLinkedListNodeHasNex(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_GetAllLinkedNodes", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b")
		nodes := ll.GetAllLinkedNodes()
		tc := caseV1Compat{Name: "LL GetAllLinkedNodes", Expected: 2, Actual: len(nodes), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_GetNextNodes_FromLinkedListNodeHasNex(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_GetNextNodes", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b").Add("c")
		nodes := ll.GetNextNodes(2)
		tc := caseV1Compat{Name: "LL GetNextNodes", Expected: 2, Actual: len(nodes), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_InsertAt_FromLinkedListNodeHasNex(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_InsertAt", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("c")
		ll.InsertAt(1, "b")
		tc := caseV1Compat{Name: "LL InsertAt", Expected: 3, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_InsertAt_Front", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("b")
		ll.InsertAt(0, "a")
		tc := caseV1Compat{Name: "LL InsertAt front", Expected: "a", Actual: ll.Head().Element, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_RemoveNodeByElementValue(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_RemoveNodeByElementValue", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b").Add("c")
		ll.RemoveNodeByElementValue("b", true, false)
		tc := caseV1Compat{Name: "LL RemoveByElem", Expected: 2, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_RemoveNodeByElementValue_First(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_RemoveNodeByElementValue_First", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b")
		ll.RemoveNodeByElementValue("a", true, false)
		tc := caseV1Compat{Name: "LL RemoveByElem first", Expected: "b", Actual: ll.Head().Element, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_RemoveNodeByIndex", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b").Add("c")
		ll.RemoveNodeByIndex(1)
		tc := caseV1Compat{Name: "LL RemoveByIndex", Expected: 2, Actual: ll.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_StringLock", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		tc := caseV1Compat{Name: "LL StringLock", Expected: true, Actual: len(ll.StringLock()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_Joins_FromLinkedListNodeHasNex(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_Joins", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		result := ll.Joins(",", "b", "c")
		tc := caseV1Compat{Name: "LL Joins", Expected: true, Actual: len(result) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_JsonModelAny", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		tc := caseV1Compat{Name: "LL JsonModelAny", Expected: true, Actual: ll.JsonModelAny() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_Json(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_Json", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		j := ll.Json()
		tc := caseV1Compat{Name: "LL Json", Expected: true, Actual: j.HasSafeItems(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_LinkedList_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov51_LinkedList_AsJsonMarshaller", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		m := ll.AsJsonMarshaller()
		tc := caseV1Compat{Name: "LL AsJsonMarshaller", Expected: true, Actual: m != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// CharCollectionDataModel
// ═══════════════════════════════════════════════════════════════

func Test_CharCollectionDataModel_NewUsing(t *testing.T) {
	safeTest(t, "Test_Cov51_CharCollectionDataModel_NewUsing", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("hello")
		dm := corestr.NewCharCollectionMapDataModelUsing(ccm)
		tc := caseV1Compat{Name: "CharCollDM NewUsing", Expected: true, Actual: dm != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_CharCollectionDataModel_NewUsingDataModel(t *testing.T) {
	safeTest(t, "Test_Cov51_CharCollectionDataModel_NewUsingDataModel", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("hello")
		dm := corestr.NewCharCollectionMapDataModelUsing(ccm)
		restored := corestr.NewCharCollectionMapUsingDataModel(dm)
		tc := caseV1Compat{Name: "CharCollDM restored", Expected: true, Actual: restored.Has("hello"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}
