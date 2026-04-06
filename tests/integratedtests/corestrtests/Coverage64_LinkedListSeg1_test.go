package corestrtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// LinkedList.go — Full coverage (~403 uncovered stmts, 1141 lines)
// =============================================================================

// ── Head / Tail / Length ──

func Test_Cov64_LinkedList_HeadTail_Empty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_HeadTail_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"headNil": ll.Head() == nil, "tailNil": ll.Tail() == nil, "len": ll.Length()}
		expected := args.Map{"headNil": true, "tailNil": true, "len": 0}
		expected.ShouldBeEqual(t, 0, "Head/Tail nil on empty", actual)
	})
}

func Test_Cov64_LinkedList_HeadTail_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_HeadTail_NonEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a").Add("b")
		actual := args.Map{"head": ll.Head().Element, "tail": ll.Tail().Element, "len": ll.Length()}
		expected := args.Map{"head": "a", "tail": "b", "len": 2}
		expected.ShouldBeEqual(t, 0, "Head/Tail after adds", actual)
	})
}

func Test_Cov64_LinkedList_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_LengthLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"len": ll.LengthLock()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthLock", actual)
	})
}

// ── IsEmpty / HasItems / IsEmptyLock ──

func Test_Cov64_LinkedList_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_IsEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"empty": ll.IsEmpty(), "hasItems": ll.HasItems()}
		expected := args.Map{"empty": true, "hasItems": false}
		expected.ShouldBeEqual(t, 0, "IsEmpty on new", actual)
	})
}

func Test_Cov64_LinkedList_HasItems(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_HasItems", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"empty": ll.IsEmpty(), "hasItems": ll.HasItems()}
		expected := args.Map{"empty": false, "hasItems": true}
		expected.ShouldBeEqual(t, 0, "HasItems after add", actual)
	})
}

func Test_Cov64_LinkedList_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_IsEmptyLock", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"empty": ll.IsEmptyLock()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock", actual)
	})
}

// ── Add variants ──

func Test_Cov64_LinkedList_Add_First(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Add_First", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("first")
		actual := args.Map{"head": ll.Head().Element, "len": ll.Length()}
		expected := args.Map{"head": "first", "len": 1}
		expected.ShouldBeEqual(t, 0, "Add first item", actual)
	})
}

func Test_Cov64_LinkedList_Add_Multiple(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Add_Multiple", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a").Add("b").Add("c")
		actual := args.Map{"len": ll.Length(), "tail": ll.Tail().Element}
		expected := args.Map{"len": 3, "tail": "c"}
		expected.ShouldBeEqual(t, 0, "Add multiple", actual)
	})
}

func Test_Cov64_LinkedList_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddLock("a")
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddLock", actual)
	})
}

func Test_Cov64_LinkedList_Adds(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Adds", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Adds", actual)
	})
}

func Test_Cov64_LinkedList_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Adds_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds()
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Adds empty", actual)
	})
}

func Test_Cov64_LinkedList_AddsLock(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddsLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddsLock("a", "b")
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsLock", actual)
	})
}

func Test_Cov64_LinkedList_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddStrings", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings([]string{"a", "b"})
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStrings", actual)
	})
}

func Test_Cov64_LinkedList_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddStrings_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings([]string{})
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStrings empty", actual)
	})
}

func Test_Cov64_LinkedList_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddNonEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmpty("")
		ll.AddNonEmpty("a")
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonEmpty skips empty", actual)
	})
}

func Test_Cov64_LinkedList_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddNonEmptyWhitespace", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmptyWhitespace("  ")
		ll.AddNonEmptyWhitespace("a")
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyWhitespace skips whitespace", actual)
	})
}

func Test_Cov64_LinkedList_AddIf_True(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddIf_True", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddIf(true, "a")
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddIf true", actual)
	})
}

func Test_Cov64_LinkedList_AddIf_False(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddIf_False", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddIf(false, "a")
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddIf false", actual)
	})
}

func Test_Cov64_LinkedList_AddsIf_True(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddsIf_True", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddsIf(true, "a", "b")
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsIf true", actual)
	})
}

func Test_Cov64_LinkedList_AddsIf_False(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddsIf_False", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddsIf(false, "a", "b")
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsIf false", actual)
	})
}

func Test_Cov64_LinkedList_AddFunc(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddFunc", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFunc(func() string { return "x" })
		actual := args.Map{"head": ll.Head().Element}
		expected := args.Map{"head": "x"}
		expected.ShouldBeEqual(t, 0, "AddFunc", actual)
	})
}

func Test_Cov64_LinkedList_AddFuncErr_NoError(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddFuncErr_NoError", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(err error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) },
		)
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddFuncErr no error", actual)
	})
}

func Test_Cov64_LinkedList_AddFuncErr_Error(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddFuncErr_Error", func() {
		ll := corestr.New.LinkedList.Create()
		errCalled := false
		ll.AddFuncErr(
			func() (string, error) { return "", errors.New("fail") },
			func(err error) { errCalled = true },
		)
		actual := args.Map{"len": ll.Length(), "errCalled": errCalled}
		expected := args.Map{"len": 0, "errCalled": true}
		expected.ShouldBeEqual(t, 0, "AddFuncErr with error", actual)
	})
}

func Test_Cov64_LinkedList_Push(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Push", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Push("a")
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Push", actual)
	})
}

func Test_Cov64_LinkedList_PushBack(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_PushBack", func() {
		ll := corestr.New.LinkedList.Create()
		ll.PushBack("a")
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "PushBack", actual)
	})
}

func Test_Cov64_LinkedList_PushFront(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_PushFront", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("b")
		ll.PushFront("a")
		actual := args.Map{"head": ll.Head().Element, "len": ll.Length()}
		expected := args.Map{"head": "a", "len": 2}
		expected.ShouldBeEqual(t, 0, "PushFront", actual)
	})
}

func Test_Cov64_LinkedList_AddFront_Empty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddFront_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddFront("a")
		actual := args.Map{"head": ll.Head().Element, "len": ll.Length()}
		expected := args.Map{"head": "a", "len": 1}
		expected.ShouldBeEqual(t, 0, "AddFront empty", actual)
	})
}

func Test_Cov64_LinkedList_AddFront_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddFront_NonEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("b").Add("c")
		ll.AddFront("a")
		actual := args.Map{"head": ll.Head().Element, "len": ll.Length()}
		expected := args.Map{"head": "a", "len": 3}
		expected.ShouldBeEqual(t, 0, "AddFront non-empty", actual)
	})
}

func Test_Cov64_LinkedList_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddItemsMap", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{"a": true, "b": false, "c": true})
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddItemsMap", actual)
	})
}

func Test_Cov64_LinkedList_AddItemsMap_Empty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddItemsMap_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{})
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddItemsMap empty", actual)
	})
}

// ── AppendNode / AppendChainOfNodes ──

func Test_Cov64_LinkedList_AppendNode_Empty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AppendNode_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		node := &corestr.LinkedListNode{Element: "a"}
		ll.AppendNode(node)
		actual := args.Map{"len": ll.Length(), "head": ll.Head().Element}
		expected := args.Map{"len": 1, "head": "a"}
		expected.ShouldBeEqual(t, 0, "AppendNode empty", actual)
	})
}

func Test_Cov64_LinkedList_AppendNode_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AppendNode_NonEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := &corestr.LinkedListNode{Element: "b"}
		ll.AppendNode(node)
		actual := args.Map{"len": ll.Length(), "tail": ll.Tail().Element}
		expected := args.Map{"len": 2, "tail": "b"}
		expected.ShouldBeEqual(t, 0, "AppendNode non-empty", actual)
	})
}

func Test_Cov64_LinkedList_AddBackNode(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddBackNode", func() {
		ll := corestr.New.LinkedList.Create()
		node := &corestr.LinkedListNode{Element: "x"}
		ll.AddBackNode(node)
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddBackNode", actual)
	})
}

// ── AddCollection / AddPointerStringsPtr ──

func Test_Cov64_LinkedList_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddCollection", func() {
		ll := corestr.New.LinkedList.Create()
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		ll.AddCollection(col)
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddCollection", actual)
	})
}

func Test_Cov64_LinkedList_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddCollection_Nil", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddCollection(nil)
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollection nil", actual)
	})
}

func Test_Cov64_LinkedList_AddPointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddPointerStringsPtr", func() {
		ll := corestr.New.LinkedList.Create()
		s1 := "a"
		ll.AddPointerStringsPtr([]*string{&s1, nil})
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddPointerStringsPtr skips nil", actual)
	})
}

// ── InsertAt ──

func Test_Cov64_LinkedList_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_InsertAt_Front", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("b", "c")
		ll.InsertAt(0, "a")
		actual := args.Map{"head": ll.Head().Element}
		expected := args.Map{"head": "a"}
		expected.ShouldBeEqual(t, 0, "InsertAt front", actual)
	})
}

func Test_Cov64_LinkedList_InsertAt_Negative(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_InsertAt_Negative", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("b")
		ll.InsertAt(-1, "a")
		actual := args.Map{"head": ll.Head().Element}
		expected := args.Map{"head": "a"}
		expected.ShouldBeEqual(t, 0, "InsertAt negative inserts front", actual)
	})
}

func Test_Cov64_LinkedList_InsertAt_Middle(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_InsertAt_Middle", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "c")
		ll.InsertAt(1, "b")
		list := ll.List()
		actual := args.Map{"second": list[1]}
		expected := args.Map{"second": "b"}
		expected.ShouldBeEqual(t, 0, "InsertAt middle", actual)
	})
}

// ── AttachWithNode ──

func Test_Cov64_LinkedList_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AttachWithNode_NilCurrent", func() {
		ll := corestr.New.LinkedList.Create()
		err := ll.AttachWithNode(nil, &corestr.LinkedListNode{Element: "a"})
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "AttachWithNode nil current", actual)
	})
}

func Test_Cov64_LinkedList_AttachWithNode_NextNotNil(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AttachWithNode_NextNotNil", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		err := ll.AttachWithNode(ll.Head(), &corestr.LinkedListNode{Element: "x"})
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "AttachWithNode next not nil", actual)
	})
}

func Test_Cov64_LinkedList_AttachWithNode_Success(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AttachWithNode_Success", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		err := ll.AttachWithNode(ll.Tail(), &corestr.LinkedListNode{Element: "b"})
		actual := args.Map{"noErr": err == nil, "len": ll.Length()}
		expected := args.Map{"noErr": true, "len": 2}
		expected.ShouldBeEqual(t, 0, "AttachWithNode success", actual)
	})
}

// ── AddAfterNode ──

func Test_Cov64_LinkedList_AddAfterNode(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddAfterNode", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.AddAfterNode(ll.Head(), "b")
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddAfterNode", actual)
	})
}

// ── AddStringsToNode ──

func Test_Cov64_LinkedList_AddStringsToNode_Empty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddStringsToNode_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.AddStringsToNode(true, ll.Head(), []string{})
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddStringsToNode empty items", actual)
	})
}

func Test_Cov64_LinkedList_AddStringsToNode_NilSkip(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddStringsToNode_NilSkip", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddStringsToNode(true, nil, []string{"a"})
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStringsToNode nil node skip", actual)
	})
}

func Test_Cov64_LinkedList_AddStringsToNode_Single(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddStringsToNode_Single", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.AddStringsToNode(false, ll.Head(), []string{"b"})
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsToNode single", actual)
	})
}

func Test_Cov64_LinkedList_AddStringsToNode_Multiple(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddStringsToNode_Multiple", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.AddStringsToNode(false, ll.Head(), []string{"b", "c"})
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AddStringsToNode multiple", actual)
	})
}

func Test_Cov64_LinkedList_AddStringsPtrToNode_Nil(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddStringsPtrToNode_Nil", func() {
		ll := corestr.New.LinkedList.Create()
		ll.AddStringsPtrToNode(true, nil, nil)
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStringsPtrToNode nil", actual)
	})
}

func Test_Cov64_LinkedList_AddStringsPtrToNode_NonNil(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddStringsPtrToNode_NonNil", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		items := []string{"b"}
		ll.AddStringsPtrToNode(false, ll.Head(), &items)
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsPtrToNode non-nil", actual)
	})
}

// ── AddCollectionToNode ──

func Test_Cov64_LinkedList_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AddCollectionToNode", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		col := corestr.New.Collection.Strings([]string{"b", "c"})
		ll.AddCollectionToNode(false, ll.Head(), col)
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AddCollectionToNode", actual)
	})
}

// ── Loop ──

func Test_Cov64_LinkedList_Loop(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Loop", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})
		actual := args.Map{"count": count}
		expected := args.Map{"count": 3}
		expected.ShouldBeEqual(t, 0, "Loop visits all", actual)
	})
}

func Test_Cov64_LinkedList_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Loop_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		called := false
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			called = true
			return false
		})
		actual := args.Map{"called": called}
		expected := args.Map{"called": false}
		expected.ShouldBeEqual(t, 0, "Loop empty", actual)
	})
}

func Test_Cov64_LinkedList_Loop_Break(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Loop_Break", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return true // break immediately
		})
		actual := args.Map{"count": count}
		expected := args.Map{"count": 1}
		expected.ShouldBeEqual(t, 0, "Loop break first", actual)
	})
}

func Test_Cov64_LinkedList_Loop_BreakSecond(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Loop_BreakSecond", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return arg.Index == 1
		})
		actual := args.Map{"count": count}
		expected := args.Map{"count": 2}
		expected.ShouldBeEqual(t, 0, "Loop break second", actual)
	})
}

// ── Filter ──

func Test_Cov64_LinkedList_Filter(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Filter", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: arg.Node.Element != "b", IsBreak: false}
		})
		actual := args.Map{"len": len(nodes)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Filter", actual)
	})
}

func Test_Cov64_LinkedList_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Filter_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true}
		})
		actual := args.Map{"len": len(nodes)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Filter empty", actual)
	})
}

func Test_Cov64_LinkedList_Filter_BreakFirst(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Filter_BreakFirst", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})
		actual := args.Map{"len": len(nodes)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Filter break first", actual)
	})
}

func Test_Cov64_LinkedList_Filter_BreakSecond(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Filter_BreakSecond", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: arg.Index == 1}
		})
		actual := args.Map{"len": len(nodes)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Filter break second", actual)
	})
}

// ── GetNextNodes / GetAllLinkedNodes ──

func Test_Cov64_LinkedList_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_GetNextNodes", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		nodes := ll.GetNextNodes(2)
		actual := args.Map{"len": len(nodes)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetNextNodes", actual)
	})
}

func Test_Cov64_LinkedList_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_GetAllLinkedNodes", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		nodes := ll.GetAllLinkedNodes()
		actual := args.Map{"len": len(nodes)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllLinkedNodes", actual)
	})
}

// ── IndexAt / SafeIndexAt ──

func Test_Cov64_LinkedList_IndexAt_Zero(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_IndexAt_Zero", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		node := ll.IndexAt(0)
		actual := args.Map{"elem": node.Element}
		expected := args.Map{"elem": "a"}
		expected.ShouldBeEqual(t, 0, "IndexAt 0", actual)
	})
}

func Test_Cov64_LinkedList_IndexAt_Middle(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_IndexAt_Middle", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		node := ll.IndexAt(1)
		actual := args.Map{"elem": node.Element}
		expected := args.Map{"elem": "b"}
		expected.ShouldBeEqual(t, 0, "IndexAt middle", actual)
	})
}

func Test_Cov64_LinkedList_IndexAt_Negative(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_IndexAt_Negative", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := ll.IndexAt(-1)
		actual := args.Map{"nil": node == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "IndexAt negative", actual)
	})
}

func Test_Cov64_LinkedList_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_SafeIndexAt", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		node := ll.SafeIndexAt(1)
		actual := args.Map{"elem": node.Element}
		expected := args.Map{"elem": "b"}
		expected.ShouldBeEqual(t, 0, "SafeIndexAt", actual)
	})
}

func Test_Cov64_LinkedList_SafeIndexAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_SafeIndexAt_OutOfRange", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := ll.SafeIndexAt(5)
		actual := args.Map{"nil": node == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "SafeIndexAt out of range", actual)
	})
}

func Test_Cov64_LinkedList_SafeIndexAt_Negative(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_SafeIndexAt_Negative", func() {
		ll := corestr.New.LinkedList.Create()
		node := ll.SafeIndexAt(-1)
		actual := args.Map{"nil": node == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "SafeIndexAt negative", actual)
	})
}

func Test_Cov64_LinkedList_SafeIndexAt_Zero(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_SafeIndexAt_Zero", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := ll.SafeIndexAt(0)
		actual := args.Map{"elem": node.Element}
		expected := args.Map{"elem": "a"}
		expected.ShouldBeEqual(t, 0, "SafeIndexAt zero", actual)
	})
}

func Test_Cov64_LinkedList_SafeIndexAtLock(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_SafeIndexAtLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := ll.SafeIndexAtLock(0)
		actual := args.Map{"elem": node.Element}
		expected := args.Map{"elem": "a"}
		expected.ShouldBeEqual(t, 0, "SafeIndexAtLock", actual)
	})
}

func Test_Cov64_LinkedList_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_SafePointerIndexAt", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ptr := ll.SafePointerIndexAt(0)
		actual := args.Map{"nonNil": ptr != nil, "val": *ptr}
		expected := args.Map{"nonNil": true, "val": "a"}
		expected.ShouldBeEqual(t, 0, "SafePointerIndexAt", actual)
	})
}

func Test_Cov64_LinkedList_SafePointerIndexAt_Nil(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_SafePointerIndexAt_Nil", func() {
		ll := corestr.New.LinkedList.Create()
		ptr := ll.SafePointerIndexAt(5)
		actual := args.Map{"nil": ptr == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "SafePointerIndexAt nil", actual)
	})
}

func Test_Cov64_LinkedList_SafePointerIndexAtUsingDefault(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_SafePointerIndexAtUsingDefault", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		val := ll.SafePointerIndexAtUsingDefault(0, "def")
		actual := args.Map{"val": val}
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "SafePointerIndexAtUsingDefault found", actual)
	})
}

func Test_Cov64_LinkedList_SafePointerIndexAtUsingDefault_Default(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_SafePointerIndexAtUsingDefault_Default", func() {
		ll := corestr.New.LinkedList.Create()
		val := ll.SafePointerIndexAtUsingDefault(5, "def")
		actual := args.Map{"val": val}
		expected := args.Map{"val": "def"}
		expected.ShouldBeEqual(t, 0, "SafePointerIndexAtUsingDefault default", actual)
	})
}

func Test_Cov64_LinkedList_SafePointerIndexAtUsingDefaultLock(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_SafePointerIndexAtUsingDefaultLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		val := ll.SafePointerIndexAtUsingDefaultLock(0, "def")
		actual := args.Map{"val": val}
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "SafePointerIndexAtUsingDefaultLock", actual)
	})
}

// ── RemoveNodeByElementValue ──

func Test_Cov64_LinkedList_RemoveNodeByElementValue_First(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_RemoveNodeByElementValue_First", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByElementValue("a", true, false)
		actual := args.Map{"head": ll.Head().Element, "len": ll.Length()}
		expected := args.Map{"head": "b", "len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByElementValue first", actual)
	})
}

func Test_Cov64_LinkedList_RemoveNodeByElementValue_Middle(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_RemoveNodeByElementValue_Middle", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByElementValue("b", true, false)
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByElementValue middle", actual)
	})
}

func Test_Cov64_LinkedList_RemoveNodeByElementValue_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_RemoveNodeByElementValue_CaseInsensitive", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("Apple", "banana")
		ll.RemoveNodeByElementValue("apple", false, false)
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByElementValue case insensitive", actual)
	})
}

// ── RemoveNodeByIndex ──

func Test_Cov64_LinkedList_RemoveNodeByIndex_First(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_RemoveNodeByIndex_First", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByIndex(0)
		actual := args.Map{"head": ll.Head().Element, "len": ll.Length()}
		expected := args.Map{"head": "b", "len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndex first", actual)
	})
}

func Test_Cov64_LinkedList_RemoveNodeByIndex_Last(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_RemoveNodeByIndex_Last", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByIndex(2)
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndex last", actual)
	})
}

func Test_Cov64_LinkedList_RemoveNodeByIndex_Middle(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_RemoveNodeByIndex_Middle", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByIndex(1)
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndex middle", actual)
	})
}

// ── RemoveNodeByIndexes ──

func Test_Cov64_LinkedList_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_RemoveNodeByIndexes", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c", "d")
		ll.RemoveNodeByIndexes(false, 1, 3)
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndexes", actual)
	})
}

func Test_Cov64_LinkedList_RemoveNodeByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_RemoveNodeByIndexes_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.RemoveNodeByIndexes(false)
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndexes no indexes", actual)
	})
}

// ── RemoveNode ──

func Test_Cov64_LinkedList_RemoveNode_Nil(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_RemoveNode_Nil", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.RemoveNode(nil)
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveNode nil", actual)
	})
}

func Test_Cov64_LinkedList_RemoveNode_First(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_RemoveNode_First", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.RemoveNode(ll.Head())
		actual := args.Map{"len": ll.Length(), "head": ll.Head().Element}
		expected := args.Map{"len": 1, "head": "b"}
		expected.ShouldBeEqual(t, 0, "RemoveNode first", actual)
	})
}

func Test_Cov64_LinkedList_RemoveNode_Second(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_RemoveNode_Second", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		second := ll.IndexAt(1)
		ll.RemoveNode(second)
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveNode second", actual)
	})
}

// ── GetCompareSummary ──

func Test_Cov64_LinkedList_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_GetCompareSummary", func() {
		a := corestr.New.LinkedList.Create()
		a.Add("x")
		b := corestr.New.LinkedList.Create()
		b.Add("x")
		s := a.GetCompareSummary(b, "left", "right")
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "GetCompareSummary", actual)
	})
}

// ── IsEquals / IsEqualsWithSensitive ──

func Test_Cov64_LinkedList_IsEquals_Same(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_IsEquals_Same", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"eq": ll.IsEquals(ll)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals same ptr", actual)
	})
}

func Test_Cov64_LinkedList_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_IsEquals_BothEmpty", func() {
		a := corestr.New.LinkedList.Create()
		b := corestr.New.LinkedList.Create()
		actual := args.Map{"eq": a.IsEquals(b)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals both empty", actual)
	})
}

func Test_Cov64_LinkedList_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_IsEquals_OneEmpty", func() {
		a := corestr.New.LinkedList.Create()
		a.Add("a")
		b := corestr.New.LinkedList.Create()
		actual := args.Map{"eq": a.IsEquals(b)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals one empty", actual)
	})
}

func Test_Cov64_LinkedList_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_IsEquals_DiffLen", func() {
		a := corestr.New.LinkedList.Create()
		a.Adds("a", "b")
		b := corestr.New.LinkedList.Create()
		b.Add("a")
		actual := args.Map{"eq": a.IsEquals(b)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff len", actual)
	})
}

func Test_Cov64_LinkedList_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_IsEquals_Nil", func() {
		a := corestr.New.LinkedList.Create()
		a.Add("a")
		actual := args.Map{"eq": a.IsEquals(nil)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals nil", actual)
	})
}

func Test_Cov64_LinkedList_IsEquals_Equal(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_IsEquals_Equal", func() {
		a := corestr.New.LinkedList.Create()
		a.Adds("a", "b")
		b := corestr.New.LinkedList.Create()
		b.Adds("a", "b")
		actual := args.Map{"eq": a.IsEquals(b)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals equal content", actual)
	})
}

func Test_Cov64_LinkedList_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_IsEqualsWithSensitive_CaseInsensitive", func() {
		a := corestr.New.LinkedList.Create()
		a.Add("Apple")
		b := corestr.New.LinkedList.Create()
		b.Add("apple")
		actual := args.Map{"eq": a.IsEqualsWithSensitive(b, false)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsWithSensitive case insensitive", actual)
	})
}

// ── ToCollection / List ──

func Test_Cov64_LinkedList_ToCollection(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_ToCollection", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		col := ll.ToCollection(0)
		actual := args.Map{"len": col.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ToCollection", actual)
	})
}

func Test_Cov64_LinkedList_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_ToCollection_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		col := ll.ToCollection(5)
		actual := args.Map{"len": col.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ToCollection empty", actual)
	})
}

func Test_Cov64_LinkedList_List(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_List", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		actual := args.Map{"len": len(ll.List())}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "List", actual)
	})
}

func Test_Cov64_LinkedList_List_Empty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_List_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"len": len(ll.List())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "List empty", actual)
	})
}

func Test_Cov64_LinkedList_ListPtr(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_ListPtr", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"len": len(ll.ListPtr())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListPtr", actual)
	})
}

func Test_Cov64_LinkedList_ListLock(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_ListLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"len": len(ll.ListLock())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListLock", actual)
	})
}

func Test_Cov64_LinkedList_ListPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_ListPtrLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"len": len(ll.ListPtrLock())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListPtrLock", actual)
	})
}

// ── String / StringLock / Join / JoinLock / Joins ──

func Test_Cov64_LinkedList_String(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_String", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"nonEmpty": ll.String() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String", actual)
	})
}

func Test_Cov64_LinkedList_String_Empty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_String_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"nonEmpty": ll.String() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String empty", actual)
	})
}

func Test_Cov64_LinkedList_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_StringLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"nonEmpty": ll.StringLock() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock", actual)
	})
}

func Test_Cov64_LinkedList_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_StringLock_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"nonEmpty": ll.StringLock() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock empty", actual)
	})
}

func Test_Cov64_LinkedList_Join(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Join", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		actual := args.Map{"val": ll.Join(",")}
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Join", actual)
	})
}

func Test_Cov64_LinkedList_JoinLock(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_JoinLock", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		actual := args.Map{"val": ll.JoinLock(",")}
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "JoinLock", actual)
	})
}

func Test_Cov64_LinkedList_Joins(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Joins", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		r := ll.Joins(",", "b", "c")
		actual := args.Map{"nonEmpty": r != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Joins", actual)
	})
}

func Test_Cov64_LinkedList_Joins_NilItems(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Joins_NilItems", func() {
		ll := corestr.New.LinkedList.Create()
		r := ll.Joins(",")
		actual := args.Map{"val": r}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Joins nil items empty", actual)
	})
}

// ── JSON ──

func Test_Cov64_LinkedList_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_JsonModel", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		actual := args.Map{"len": len(ll.JsonModel())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel", actual)
	})
}

func Test_Cov64_LinkedList_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_JsonModelAny", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"nonNil": ll.JsonModelAny() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny", actual)
	})
}

func Test_Cov64_LinkedList_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_MarshalJSON", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		b, err := ll.MarshalJSON()
		actual := args.Map{"noErr": err == nil, "nonEmpty": len(b) > 0}
		expected := args.Map{"noErr": true, "nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "MarshalJSON", actual)
	})
}

func Test_Cov64_LinkedList_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_UnmarshalJSON", func() {
		ll := corestr.New.LinkedList.Create()
		err := ll.UnmarshalJSON([]byte(`["a","b"]`))
		actual := args.Map{"noErr": err == nil, "len": ll.Length()}
		expected := args.Map{"noErr": true, "len": 2}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON", actual)
	})
}

func Test_Cov64_LinkedList_UnmarshalJSON_Error(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_UnmarshalJSON_Error", func() {
		ll := corestr.New.LinkedList.Create()
		err := ll.UnmarshalJSON([]byte("invalid"))
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON error", actual)
	})
}

func Test_Cov64_LinkedList_Json(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Json", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		r := ll.Json()
		actual := args.Map{"nonEmpty": r.JsonString() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Json", actual)
	})
}

func Test_Cov64_LinkedList_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_JsonPtr", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		r := ll.JsonPtr()
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonPtr", actual)
	})
}

func Test_Cov64_LinkedList_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_ParseInjectUsingJson", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jr := ll.JsonPtr()
		ll2 := corestr.New.LinkedList.Create()
		r, err := ll2.ParseInjectUsingJson(jr)
		actual := args.Map{"noErr": err == nil, "nonNil": r != nil}
		expected := args.Map{"noErr": true, "nonNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson", actual)
	})
}

func Test_Cov64_LinkedList_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_ParseInjectUsingJson_Error", func() {
		ll := corestr.New.LinkedList.Create()
		jr := &corejson.Result{Error: errors.New("fail")}
		_, err := ll.ParseInjectUsingJson(jr)
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson error", actual)
	})
}

func Test_Cov64_LinkedList_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_ParseInjectUsingJsonMust", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jr := ll.JsonPtr()
		ll2 := corestr.New.LinkedList.Create()
		r := ll2.ParseInjectUsingJsonMust(jr)
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust", actual)
	})
}

func Test_Cov64_LinkedList_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_ParseInjectUsingJsonMust_Panics", func() {
		ll := corestr.New.LinkedList.Create()
		jr := &corejson.Result{Error: errors.New("fail")}
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			ll.ParseInjectUsingJsonMust(jr)
		}()
		actual := args.Map{"panicked": panicked}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust panics", actual)
	})
}

func Test_Cov64_LinkedList_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_JsonParseSelfInject", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jr := ll.JsonPtr()
		ll2 := corestr.New.LinkedList.Create()
		err := ll2.JsonParseSelfInject(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject", actual)
	})
}

func Test_Cov64_LinkedList_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AsJsonMarshaller", func() {
		ll := corestr.New.LinkedList.Create()
		actual := args.Map{"nonNil": ll.AsJsonMarshaller() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonMarshaller", actual)
	})
}

// ── Clear / RemoveAll ──

func Test_Cov64_LinkedList_Clear(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Clear", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.Clear()
		actual := args.Map{"empty": ll.IsEmpty(), "len": ll.Length()}
		expected := args.Map{"empty": true, "len": 0}
		expected.ShouldBeEqual(t, 0, "Clear", actual)
	})
}

func Test_Cov64_LinkedList_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_Clear_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Clear()
		actual := args.Map{"empty": ll.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Clear empty", actual)
	})
}

func Test_Cov64_LinkedList_RemoveAll(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_RemoveAll", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.RemoveAll()
		actual := args.Map{"empty": ll.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "RemoveAll", actual)
	})
}

// ── AppendChainOfNodes ──

func Test_Cov64_LinkedList_AppendChainOfNodes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AppendChainOfNodes_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		chain := corestr.New.LinkedList.Create()
		chain.Adds("a", "b")
		ll.AppendChainOfNodes(chain.Head())
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendChainOfNodes empty list", actual)
	})
}

func Test_Cov64_LinkedList_AppendChainOfNodes_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov64_LinkedList_AppendChainOfNodes_NonEmpty", func() {
		ll := corestr.New.LinkedList.Create()
		ll.Add("x")
		chain := corestr.New.LinkedList.Create()
		chain.Adds("a", "b")
		ll.AppendChainOfNodes(chain.Head())
		actual := args.Map{"len": ll.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AppendChainOfNodes non-empty list", actual)
	})
}
