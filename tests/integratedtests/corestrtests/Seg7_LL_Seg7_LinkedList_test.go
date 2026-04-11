package corestrtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// LinkedList — Segment 7e
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg7_LL_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_IsEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{
			"empty": ll.IsEmpty(),
			"hasItems": ll.HasItems(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"hasItems": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- true on empty", actual)
	})
}

func Test_Seg7_LL_Add(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_Add", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a").Add("b").Add("c")

		// Act
		actual := args.Map{
			"len": ll.Length(),
			"head": ll.Head().Element,
			"tail": ll.Tail().Element,
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"head": "a",
			"tail": "c",
		}
		expected.ShouldBeEqual(t, 0, "Add -- 3 items", actual)
	})
}

func Test_Seg7_LL_AddLock(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddLock("a")

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddLock -- 1 item", actual)
	})
}

func Test_Seg7_LL_Adds(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_Adds", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Adds -- 3 items", actual)
	})
}

func Test_Seg7_LL_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_Adds_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds()

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Adds empty -- no change", actual)
	})
}

func Test_Seg7_LL_AddsLock(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddsLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddsLock("a", "b")

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsLock -- 2 items", actual)
	})
}

func Test_Seg7_LL_AddStrings(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddStrings", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings([]string{"a", "b"})

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStrings -- 2 items", actual)
	})
}

func Test_Seg7_LL_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddStrings_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings([]string{})

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStrings empty -- no change", actual)
	})
}

func Test_Seg7_LL_AddFront(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddFront", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("b").AddFront("a")

		// Act
		actual := args.Map{
			"head": ll.Head().Element,
			"len": ll.Length(),
		}

		// Assert
		expected := args.Map{
			"head": "a",
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "AddFront -- prepended", actual)
	})
}

func Test_Seg7_LL_AddFront_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddFront_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddFront("a")

		// Act
		actual := args.Map{
			"head": ll.Head().Element,
			"len": ll.Length(),
		}

		// Assert
		expected := args.Map{
			"head": "a",
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "AddFront empty -- added as head", actual)
	})
}

func Test_Seg7_LL_PushFront(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_PushFront", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("b").PushFront("a")

		// Act
		actual := args.Map{"head": ll.Head().Element}

		// Assert
		expected := args.Map{"head": "a"}
		expected.ShouldBeEqual(t, 0, "PushFront -- prepended", actual)
	})
}

func Test_Seg7_LL_PushBack(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_PushBack", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a").PushBack("b")

		// Act
		actual := args.Map{"tail": ll.Tail().Element}

		// Assert
		expected := args.Map{"tail": "b"}
		expected.ShouldBeEqual(t, 0, "PushBack -- appended", actual)
	})
}

func Test_Seg7_LL_Push(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_Push", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Push("a")

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Push -- added", actual)
	})
}

func Test_Seg7_LL_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddNonEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmpty("a").AddNonEmpty("")

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonEmpty -- skips empty", actual)
	})
}

func Test_Seg7_LL_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddNonEmptyWhitespace", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmptyWhitespace("a").AddNonEmptyWhitespace("  ")

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyWhitespace -- skips ws", actual)
	})
}

func Test_Seg7_LL_AddIf(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddIf", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddIf(true, "a").AddIf(false, "b")

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddIf -- only true", actual)
	})
}

func Test_Seg7_LL_AddsIf(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddsIf", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddsIf(true, "a", "b").AddsIf(false, "c")

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsIf -- only true", actual)
	})
}

func Test_Seg7_LL_AddFunc(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddFunc", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddFunc(func() string { return "hello" })

		// Act
		actual := args.Map{"head": ll.Head().Element}

		// Assert
		expected := args.Map{"head": "hello"}
		expected.ShouldBeEqual(t, 0, "AddFunc -- func result", actual)
	})
}

func Test_Seg7_LL_AddFuncErr_Success(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddFuncErr_Success", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(err error) {},
		)

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddFuncErr success -- added", actual)
	})
}

func Test_Seg7_LL_AddFuncErr_Error(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddFuncErr_Error", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		called := false
		ll.AddFuncErr(
			func() (string, error) { return "", fmt.Errorf("err") },
			func(err error) { called = true },
		)

		// Act
		actual := args.Map{
			"len": ll.Length(),
			"called": called,
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"called": true,
		}
		expected.ShouldBeEqual(t, 0, "AddFuncErr error -- handler called", actual)
	})
}

func Test_Seg7_LL_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddItemsMap", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{"a": true, "b": false})

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddItemsMap -- only true", actual)
	})
}

func Test_Seg7_LL_AddItemsMap_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddItemsMap_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{})

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddItemsMap empty -- no change", actual)
	})
}

func Test_Seg7_LL_AppendNode(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AppendNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		node := &corestr.LinkedListNode{Element: "a"}
		ll.AppendNode(node)

		// Act
		actual := args.Map{
			"len": ll.Length(),
			"head": ll.Head().Element,
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"head": "a",
		}
		expected.ShouldBeEqual(t, 0, "AppendNode -- added", actual)
	})
}

func Test_Seg7_LL_AppendNode_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AppendNode_NonEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := &corestr.LinkedListNode{Element: "b"}
		ll.AppendNode(node)

		// Act
		actual := args.Map{
			"len": ll.Length(),
			"tail": ll.Tail().Element,
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"tail": "b",
		}
		expected.ShouldBeEqual(t, 0, "AppendNode non-empty -- appended", actual)
	})
}

func Test_Seg7_LL_AddBackNode(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddBackNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		node := &corestr.LinkedListNode{Element: "a"}
		ll.AddBackNode(node)

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddBackNode -- delegates", actual)
	})
}

func Test_Seg7_LL_InsertAt(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_InsertAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "c")
		ll.InsertAt(1, "b")

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "InsertAt -- inserted", actual)
	})
}

func Test_Seg7_LL_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_InsertAt_Front", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("b")
		ll.InsertAt(0, "a")

		// Act
		actual := args.Map{"head": ll.Head().Element}

		// Assert
		expected := args.Map{"head": "a"}
		expected.ShouldBeEqual(t, 0, "InsertAt 0 -- front", actual)
	})
}

func Test_Seg7_LL_AddCollection(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddCollection", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		ll.AddCollection(c)

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddCollection -- 2 items", actual)
	})
}

func Test_Seg7_LL_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddCollection_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddCollection(nil)

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollection nil -- no change", actual)
	})
}

func Test_Seg7_LL_AddPointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AddPointerStringsPtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		s1 := "a"
		ll.AddPointerStringsPtr([]*string{&s1, nil})

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddPointerStringsPtr -- skips nil", actual)
	})
}

// ── Accessors ───────────────────────────────────────────────────────────────

func Test_Seg7_LL_List(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_List", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")

		// Act
		actual := args.Map{
			"len": len(ll.List()),
			"first": ll.List()[0],
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "List -- correct", actual)
	})
}

func Test_Seg7_LL_List_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_List_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"len": len(ll.List())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "List empty -- 0", actual)
	})
}

func Test_Seg7_LL_ListLock(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_ListLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a")

		// Act
		actual := args.Map{"len": len(ll.ListLock())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListLock -- 1", actual)
	})
}

func Test_Seg7_LL_LengthLock(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_LengthLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"len": ll.LengthLock()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthLock -- 1", actual)
	})
}

func Test_Seg7_LL_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_IsEmptyLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"empty": ll.IsEmptyLock()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock -- true", actual)
	})
}

func Test_Seg7_LL_ToCollection(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_ToCollection", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		c := ll.ToCollection(0)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ToCollection -- 2 items", actual)
	})
}

func Test_Seg7_LL_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_ToCollection_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		c := ll.ToCollection(0)

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "ToCollection empty -- empty", actual)
	})
}

func Test_Seg7_LL_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_SafeIndexAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")

		// Act
		actual := args.Map{
			"at0":   ll.SafeIndexAt(0).Element,
			"at2":   ll.SafeIndexAt(2).Element,
			"neg":   ll.SafeIndexAt(-1) == nil,
			"outOf": ll.SafeIndexAt(10) == nil,
		}

		// Assert
		expected := args.Map{
			"at0": "a",
			"at2": "c",
			"neg": true,
			"outOf": true,
		}
		expected.ShouldBeEqual(t, 0, "SafeIndexAt -- various", actual)
	})
}

func Test_Seg7_LL_SafeIndexAtLock(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_SafeIndexAtLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a")

		// Act
		actual := args.Map{"at0": ll.SafeIndexAtLock(0).Element}

		// Assert
		expected := args.Map{"at0": "a"}
		expected.ShouldBeEqual(t, 0, "SafeIndexAtLock -- found", actual)
	})
}

func Test_Seg7_LL_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_SafePointerIndexAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a")
		p := ll.SafePointerIndexAt(0)
		nilP := ll.SafePointerIndexAt(5)

		// Act
		actual := args.Map{
			"val": *p,
			"nilP": nilP == nil,
		}

		// Assert
		expected := args.Map{
			"val": "a",
			"nilP": true,
		}
		expected.ShouldBeEqual(t, 0, "SafePointerIndexAt -- correct", actual)
	})
}

func Test_Seg7_LL_SafePointerIndexAtUsingDefault(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_SafePointerIndexAtUsingDefault", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a")

		// Act
		actual := args.Map{
			"val": ll.SafePointerIndexAtUsingDefault(0, "def"),
			"def": ll.SafePointerIndexAtUsingDefault(5, "def"),
		}

		// Assert
		expected := args.Map{
			"val": "a",
			"def": "def",
		}
		expected.ShouldBeEqual(t, 0, "SafePointerIndexAtUsingDefault -- correct", actual)
	})
}

func Test_Seg7_LL_SafePointerIndexAtUsingDefaultLock(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_SafePointerIndexAtUsingDefaultLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a")

		// Act
		actual := args.Map{"val": ll.SafePointerIndexAtUsingDefaultLock(0, "def")}

		// Assert
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "SafePointerIndexAtUsingDefaultLock -- correct", actual)
	})
}

func Test_Seg7_LL_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_GetNextNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		nodes := ll.GetNextNodes(2)

		// Act
		actual := args.Map{"len": len(nodes)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetNextNodes -- 2 nodes", actual)
	})
}

func Test_Seg7_LL_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_GetAllLinkedNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		nodes := ll.GetAllLinkedNodes()

		// Act
		actual := args.Map{"len": len(nodes)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllLinkedNodes -- 2 nodes", actual)
	})
}

// ── IsEquals ────────────────────────────────────────────────────────────────

func Test_Seg7_LL_IsEquals(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_IsEquals", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Create()
		ll1.Adds("a", "b")
		ll2 := corestr.New.LinkedList.Create()
		ll2.Adds("a", "b")
		ll3 := corestr.New.LinkedList.Create()
		ll3.Adds("x")

		// Act
		actual := args.Map{
			"eq":   ll1.IsEquals(ll2),
			"neq":  ll1.IsEquals(ll3),
			"self": ll1.IsEquals(ll1),
			"nil":  ll1.IsEquals(nil),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"neq": false,
			"self": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEquals -- various", actual)
	})
}

func Test_Seg7_LL_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_IsEquals_BothEmpty", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Create()
		ll2 := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"eq": ll1.IsEquals(ll2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals both empty -- true", actual)
	})
}

func Test_Seg7_LL_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_IsEquals_OneEmpty", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Create()
		ll1.Add("a")
		ll2 := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"eq": ll1.IsEquals(ll2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals one empty -- false", actual)
	})
}

func Test_Seg7_LL_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_IsEquals_DiffLen", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Create()
		ll1.Adds("a")
		ll2 := corestr.New.LinkedList.Create()
		ll2.Adds("a", "b")

		// Act
		actual := args.Map{"eq": ll1.IsEquals(ll2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff len -- false", actual)
	})
}

func Test_Seg7_LL_IsEqualsWithSensitive(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_IsEqualsWithSensitive", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Create()
		ll1.Adds("ABC")
		ll2 := corestr.New.LinkedList.Create()
		ll2.Adds("abc")

		// Act
		actual := args.Map{
			"sensitive":   ll1.IsEqualsWithSensitive(ll2, true),
			"insensitive": ll1.IsEqualsWithSensitive(ll2, false),
		}

		// Assert
		expected := args.Map{
			"sensitive": false,
			"insensitive": true,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualsWithSensitive -- case matters", actual)
	})
}

// ── Loop / Filter ───────────────────────────────────────────────────────────

func Test_Seg7_LL_Loop(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_Loop", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})

		// Act
		actual := args.Map{"count": count}

		// Assert
		expected := args.Map{"count": 3}
		expected.ShouldBeEqual(t, 0, "Loop -- visits all", actual)
	})
}

func Test_Seg7_LL_Loop_Break(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_Loop_Break", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return true
		})

		// Act
		actual := args.Map{"count": count}

		// Assert
		expected := args.Map{"count": 1}
		expected.ShouldBeEqual(t, 0, "Loop break -- stops at first", actual)
	})
}

func Test_Seg7_LL_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_Loop_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})

		// Act
		actual := args.Map{"count": count}

		// Assert
		expected := args.Map{"count": 0}
		expected.ShouldBeEqual(t, 0, "Loop empty -- no visits", actual)
	})
}

func Test_Seg7_LL_Filter(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_Filter", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		result := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{
				Value:   arg.Node,
				IsKeep:  arg.Node.Element != "b",
				IsBreak: false,
			}
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Filter -- excludes b", actual)
	})
}

func Test_Seg7_LL_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_Filter_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		result := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Filter empty -- 0", actual)
	})
}

func Test_Seg7_LL_Filter_Break(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_Filter_Break", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		result := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Filter break -- 1 item", actual)
	})
}

// ── String / Join / JSON ────────────────────────────────────────────────────

func Test_Seg7_LL_String(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_String", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"nonEmpty": ll.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Seg7_LL_String_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_String_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"nonEmpty": ll.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String empty -- still non-empty (has NoElements)", actual)
	})
}

func Test_Seg7_LL_StringLock(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_StringLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"nonEmpty": ll.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock -- non-empty", actual)
	})
}

func Test_Seg7_LL_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_StringLock_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"nonEmpty": ll.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock empty -- non-empty", actual)
	})
}

func Test_Seg7_LL_Join(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_Join", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")

		// Act
		actual := args.Map{"val": ll.Join(",")}

		// Assert
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Join -- comma separated", actual)
	})
}

func Test_Seg7_LL_JoinLock(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_JoinLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")

		// Act
		actual := args.Map{"val": ll.JoinLock(",")}

		// Assert
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "JoinLock -- comma separated", actual)
	})
}

func Test_Seg7_LL_Joins(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_Joins", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a")

		// Act
		actual := args.Map{"nonEmpty": ll.Joins(",", "b") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Joins -- combined", actual)
	})
}

func Test_Seg7_LL_Joins_NilItems(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_Joins_NilItems", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"val": ll.Joins(",", nil...)}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Joins nil -- empty", actual)
	})
}

func Test_Seg7_LL_Json(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_Json", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		j := ll.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Seg7_LL_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_MarshalJSON", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		b, err := ll.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "MarshalJSON -- success", actual)
	})
}

func Test_Seg7_LL_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_UnmarshalJSON", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		b, _ := ll.MarshalJSON()
		ll2 := corestr.New.LinkedList.Create()
		err := ll2.UnmarshalJSON(b)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": ll2.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_Seg7_LL_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_UnmarshalJSON_Invalid", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		err := ll.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_Seg7_LL_JsonModel(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_JsonModel", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"len": len(ll.JsonModel())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel -- 1 item", actual)
	})
}

func Test_Seg7_LL_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_JsonModelAny", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"notNil": ll.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_Seg7_LL_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_ParseInjectUsingJson", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jr := ll.JsonPtr()
		ll2 := corestr.New.LinkedList.Create()
		_, err := ll2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- success", actual)
	})
}

func Test_Seg7_LL_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_ParseInjectUsingJsonMust", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jr := ll.JsonPtr()
		ll2 := corestr.New.LinkedList.Create()
		result := ll2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_Seg7_LL_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_JsonParseSelfInject", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jr := ll.JsonPtr()
		ll2 := corestr.New.LinkedList.Create()
		err := ll2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

func Test_Seg7_LL_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_AsJsonMarshaller", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"notNil": ll.AsJsonMarshaller() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonMarshaller -- non-nil", actual)
	})
}

// ── Remove / Clear ──────────────────────────────────────────────────────────

func Test_Seg7_LL_RemoveNodeByElementValue(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_RemoveNodeByElementValue", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByElementValue("b", true, false)

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByElementValue -- removed b", actual)
	})
}

func Test_Seg7_LL_RemoveNodeByElementValue_First(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_RemoveNodeByElementValue_First", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.RemoveNodeByElementValue("a", true, false)

		// Act
		actual := args.Map{
			"len": ll.Length(),
			"head": ll.Head().Element,
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"head": "b",
		}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByElementValue first -- removed", actual)
	})
}

func Test_Seg7_LL_RemoveNodeByElementValue_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_RemoveNodeByElementValue_CaseInsensitive", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("ABC", "def")
		ll.RemoveNodeByElementValue("abc", false, false)

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByElementValue case-insensitive -- removed", actual)
	})
}

func Test_Seg7_LL_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_RemoveNodeByIndex", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByIndex(1)

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndex -- removed index 1", actual)
	})
}

func Test_Seg7_LL_RemoveNodeByIndex_First(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_RemoveNodeByIndex_First", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.RemoveNodeByIndex(0)

		// Act
		actual := args.Map{
			"len": ll.Length(),
			"head": ll.Head().Element,
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"head": "b",
		}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndex first -- removed", actual)
	})
}

func Test_Seg7_LL_RemoveNodeByIndex_Last(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_RemoveNodeByIndex_Last", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.RemoveNodeByIndex(1)

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndex last -- removed", actual)
	})
}

func Test_Seg7_LL_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_RemoveNodeByIndexes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c", "d")
		ll.RemoveNodeByIndexes(false, 1, 3)

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndexes -- removed 2", actual)
	})
}

func Test_Seg7_LL_RemoveNodeByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_RemoveNodeByIndexes_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.RemoveNodeByIndexes(false)

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndexes empty -- no change", actual)
	})
}

func Test_Seg7_LL_RemoveNode(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_RemoveNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		node := ll.SafeIndexAt(1)
		ll.RemoveNode(node)

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveNode -- removed", actual)
	})
}

func Test_Seg7_LL_RemoveNode_Nil(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_RemoveNode_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.RemoveNode(nil)

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveNode nil -- no change", actual)
	})
}

func Test_Seg7_LL_RemoveNode_First(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_RemoveNode_First", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		node := ll.Head()
		ll.RemoveNode(node)

		// Act
		actual := args.Map{
			"len": ll.Length(),
			"head": ll.Head().Element,
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"head": "b",
		}
		expected.ShouldBeEqual(t, 0, "RemoveNode first -- removed head", actual)
	})
}

func Test_Seg7_LL_Clear(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_Clear", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.Clear()

		// Act
		actual := args.Map{"empty": ll.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Seg7_LL_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_Clear_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		result := ll.Clear()

		// Act
		actual := args.Map{"same": result == ll}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "Clear empty -- returns self", actual)
	})
}

func Test_Seg7_LL_RemoveAll(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_RemoveAll", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.RemoveAll()

		// Act
		actual := args.Map{"empty": ll.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "RemoveAll -- emptied", actual)
	})
}

func Test_Seg7_LL_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_Seg7_LL_GetCompareSummary", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Create()
		ll1.Adds("a")
		ll2 := corestr.New.LinkedList.Create()
		ll2.Adds("a")

		// Act
		actual := args.Map{"nonEmpty": ll1.GetCompareSummary(ll2, "left", "right") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "GetCompareSummary -- non-empty", actual)
	})
}
