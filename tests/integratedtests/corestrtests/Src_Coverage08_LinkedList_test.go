package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_SrcC08_LinkedList_Basic_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedList_Basic_Verification", func() {
		// Arrange
		tc := srcC08LinkedListBasicTestCase
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{
			"isEmpty":     ll.IsEmpty(),
			"hasItems":    ll.HasItems(),
			"length":      ll.Length(),
			"isEmptyLock": ll.IsEmptyLock(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedList_Add_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedList_Add_Verification", func() {
		// Arrange
		tc := srcC08LinkedListAddTestCase
		ll := corestr.New.LinkedList.Create()

		// Act
		ll.Add("a").Add("b").Add("c")
		length := ll.Length()
		headEl := ll.Head().Element
		tailEl := ll.Tail().Element
		ll.AddLock("d")
		ll.AddFront("z")
		afterFront := ll.Head().Element
		noPanic := !callPanicsSrcC08(func() {
			ll.AddNonEmpty("")
			ll.AddNonEmpty("e")
			ll.AddNonEmptyWhitespace("   ")
			ll.AddNonEmptyWhitespace("f")
			ll.AddIf(false, "skip")
			ll.AddIf(true, "g")
			ll.AddsIf(false, "x")
			ll.AddsIf(true, "h", "i")
			ll.AddFunc(func() string { return "j" })
			ll.AddFuncErr(func() (string, error) { return "k", nil }, func(e error) {})
			ll.Push("l")
			ll.PushFront("m")
			ll.PushBack("n")
		})
		actual := args.Map{
			"length":      length,
			"headElement": headEl,
			"tailElement": tailEl,
			"afterFront":  afterFront,
			"noPanic":     noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedList_Adds_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedList_Adds_Verification", func() {
		// Arrange
		tc := srcC08LinkedListAddsTestCase

		// Act
		noPanic := !callPanicsSrcC08(func() {
			ll := corestr.New.LinkedList.Create()
			ll.Adds("a", "b")
			ll.Adds()
			ll.AddStrings([]string{"c"})
			ll.AddStrings(nil)
			ll.AddsLock("d")
			ll.AddCollection(corestr.New.Collection.Strings([]string{"e"}))
			ll.AddCollection(nil)
			ll.AddItemsMap(map[string]bool{"f": true, "g": false})
			ll.AddItemsMap(nil)
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedList_List_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedList_List_Verification", func() {
		// Arrange
		tc := srcC08LinkedListListTestCase
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})

		// Act
		noPanic := !callPanicsSrcC08(func() {
			_ = ll.ListPtr()
			_ = ll.ListLock()
			_ = ll.ListPtrLock()
			_ = ll.String()
			_ = ll.StringLock()
			_ = ll.Join(",")
			_ = ll.JoinLock(",")
		})
		actual := args.Map{
			"listLen": len(ll.List()),
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedList_ToCollection_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedList_ToCollection_Verification", func() {
		// Arrange
		tc := srcC08LinkedListToCollectionTestCase
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		empty := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{
			"colLen":      ll.ToCollection(0).Length(),
			"emptyColLen": empty.ToCollection(0).Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedList_Loop_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedList_Loop_Verification", func() {
		// Arrange
		tc := srcC08LinkedListLoopTestCase
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		count := 0

		// Act
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})
		actual := args.Map{
			"count": count,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedList_Filter_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedList_Filter_Verification", func() {
		// Arrange
		tc := srcC08LinkedListFilterTestCase
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})

		// Act
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})
		actual := args.Map{
			"nodesLen": len(nodes),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedList_IndexAt_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedList_IndexAt_Verification", func() {
		// Arrange
		tc := srcC08LinkedListIndexAtTestCase
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})

		// Act
		node := ll.SafeIndexAt(1)
		noPanic := !callPanicsSrcC08(func() {
			_ = ll.SafePointerIndexAt(1)
			_ = ll.SafePointerIndexAt(-1)
			_ = ll.SafePointerIndexAtUsingDefault(1, "def")
			_ = ll.SafePointerIndexAtUsingDefault(-1, "def")
			_ = ll.SafePointerIndexAtUsingDefaultLock(0, "def")
			_ = ll.SafeIndexAtLock(0)
		})
		actual := args.Map{
			"element":     node.Element,
			"negOneIsNil": ll.SafeIndexAt(-1) == nil,
			"outIsNil":    ll.SafeIndexAt(99) == nil,
			"noPanic":     noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedList_NextNodes_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedList_NextNodes_Verification", func() {
		// Arrange
		tc := srcC08LinkedListNextNodesTestCase
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{
			"nextNodesLen": len(ll.GetNextNodes(2)),
			"allNodesLen":  len(ll.GetAllLinkedNodes()),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedList_Equals_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedList_Equals_Verification", func() {
		// Arrange
		tc := srcC08LinkedListEqualsTestCase
		ll1 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"isEquals":    ll1.IsEquals(ll2),
			"isSensitive": ll1.IsEqualsWithSensitive(ll2, false),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedList_Remove_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedList_Remove_Verification", func() {
		// Arrange
		tc := srcC08LinkedListRemoveTestCase
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})

		// Act
		ll.RemoveNodeByIndex(0)
		actual := args.Map{
			"length": ll.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedList_Clear_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedList_Clear_Verification", func() {
		// Arrange
		tc := srcC08LinkedListClearTestCase
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		ll.Clear()
		ll.RemoveAll()
		actual := args.Map{
			"length": ll.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedList_Json_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedList_Json_Verification", func() {
		// Arrange
		tc := srcC08LinkedListJsonTestCase
		ll := corestr.New.LinkedList.Strings([]string{"a"})

		// Act
		noPanic := !callPanicsSrcC08(func() {
			_ = ll.JsonModel()
			_ = ll.JsonModelAny()
			_, _ = ll.MarshalJSON()
			_ = ll.AsJsonMarshaller()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedList_Joins_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedList_Joins_Verification", func() {
		// Arrange
		tc := srcC08LinkedListJoinsTestCase
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"nonEmpty": ll.Joins(",", "c") != "",
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedList_AppendNode_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedList_AppendNode_Verification", func() {
		// Arrange
		tc := srcC08LinkedListAppendNodeTestCase
		ll := corestr.New.LinkedList.Create()

		// Act
		ll.AppendNode(&corestr.LinkedListNode{Element: "a"})
		ll.AppendNode(&corestr.LinkedListNode{Element: "b"})
		actual := args.Map{
			"length": ll.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedList_AppendChain_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedList_AppendChain_Verification", func() {
		// Arrange
		tc := srcC08LinkedListAppendChainTestCase
		// Build chain using a temporary linked list, then get head node
		tmpLL := corestr.New.LinkedList.Strings([]string{"a", "b"})
		chainHead := tmpLL.Head()
		ll := corestr.New.LinkedList.Create()

		// Act
		ll.AppendChainOfNodes(chainHead)
		actual := args.Map{
			"length": ll.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedListNode_Methods_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedListNode_Methods_Verification", func() {
		// Arrange
		tc := srcC08LinkedListNodeTestCase
		n := &corestr.LinkedListNode{Element: "a"}

		// Act
		c := n.Clone()
		noPanic := !callPanicsSrcC08(func() {
			_ = n.List()
			_ = n.ListPtr()
			_ = n.Join(",")
			_ = n.StringList("header: ")
			_ = n.CreateLinkedList()
		})
		actual := args.Map{
			"hasNext":          n.HasNext(),
			"string":           n.String(),
			"isEqualValue":     n.IsEqualValue("a"),
			"isSensitiveTrue":  n.IsEqualValueSensitive("a", true),
			"isSensitiveFalse": n.IsEqualValueSensitive("A", false),
			"cloneElement":     c.Element,
			"noPanic":          noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedListNode_EndOfChain_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedListNode_EndOfChain_Verification", func() {
		// Arrange
		tc := srcC08LinkedListEndOfChainTestCase
		// Build chain via linked list to access chained nodes through public API
		tmpLL := corestr.New.LinkedList.Strings([]string{"a", "b"})
		n1 := tmpLL.Head()

		// Act
		end, length := n1.EndOfChain()
		actual := args.Map{
			"endElement": end.Element,
			"chainLen":   length,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC08_LinkedListNode_Equal_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC08_LinkedListNode_Equal_Verification", func() {
		// Arrange
		tc := srcC08LinkedListNodeEqualTestCase
		n1 := &corestr.LinkedListNode{Element: "a"}
		n2 := &corestr.LinkedListNode{Element: "a"}

		// Act
		actual := args.Map{
			"isEqual":          n1.IsEqual(n2),
			"isChainEqual":     n1.IsChainEqual(n2, true),
			"isEqualSensitive": n1.IsEqualSensitive(n2, true),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC08(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}
