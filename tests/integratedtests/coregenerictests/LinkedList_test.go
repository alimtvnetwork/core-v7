package coregenerictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Test: EmptyLinkedList
// ==========================================================================

func Test_LinkedList_Empty(t *testing.T) {
	tc := linkedListEmptyTestCase
	ll := coregeneric.EmptyLinkedList[int]()

	// Act
	actual := args.Map{
		"isEmpty":  ll.IsEmpty(),
		"length":   ll.Length(),
		"hasItems": ll.HasItems(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

// ==========================================================================
// Test: LinkedListFrom
// ==========================================================================

func Test_LinkedList_FromSlice(t *testing.T) {
	tc := linkedListFromSliceTestCase
	ll := coregeneric.LinkedListFrom([]string{"a", "b", "c"})

	// Act
	actual := args.Map{
		"length": ll.Length(),
		"first":  ll.First(),
		"last":   ll.Last(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

func Test_LinkedList_FromEmptySlice(t *testing.T) {
	tc := linkedListFromEmptySliceTestCase
	ll := coregeneric.LinkedListFrom([]int{})

	actLines := []string{fmt.Sprintf("%v", ll.IsEmpty())}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: Add
// ==========================================================================

func Test_LinkedList_AddSingle(t *testing.T) {
	tc := linkedListAddSingleTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Add(42)

	// Act
	actual := args.Map{
		"length": ll.Length(),
		"head":   ll.First(),
		"tail":   ll.Last(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

func Test_LinkedList_AddMultiple(t *testing.T) {
	tc := linkedListAddMultipleTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Add(1).Add(2).Add(3)

	// Act
	actual := args.Map{
		"head":   ll.First(),
		"tail":   ll.Last(),
		"length": ll.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

// ==========================================================================
// Test: AddFront
// ==========================================================================

func Test_LinkedList_AddFrontPrepends(t *testing.T) {
	tc := linkedListAddFrontPrependsTestCase
	ll := coregeneric.LinkedListFrom([]int{2, 3})
	ll.AddFront(1)

	// Act
	actual := args.Map{
		"head":   ll.First(),
		"tail":   ll.Last(),
		"length": ll.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

func Test_LinkedList_AddFrontEmpty(t *testing.T) {
	tc := linkedListAddFrontEmptyTestCase
	ll := coregeneric.EmptyLinkedList[string]()
	ll.AddFront("first")

	// Act
	actual := args.Map{
		"head":   ll.First(),
		"tail":   ll.Last(),
		"length": ll.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

// ==========================================================================
// Test: Adds
// ==========================================================================

func Test_LinkedList_Adds(t *testing.T) {
	tc := linkedListAddsTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Adds(1, 2, 3)

	actLines := []string{fmt.Sprintf("%v", ll.Length())}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: AddSlice
// ==========================================================================

func Test_LinkedList_AddSlice(t *testing.T) {
	tc := linkedListAddSliceTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddSlice([]int{10, 20})

	actLines := []string{fmt.Sprintf("%v", ll.Length())}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: AddIf
// ==========================================================================

func Test_LinkedList_AddIfTrue(t *testing.T) {
	tc := linkedListAddIfTrueTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddIf(true, 5)

	actLines := []string{fmt.Sprintf("%v", ll.Length())}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_AddIfFalse(t *testing.T) {
	tc := linkedListAddIfFalseTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddIf(false, 5)

	actLines := []string{fmt.Sprintf("%v", ll.IsEmpty())}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: AddsIf
// ==========================================================================

func Test_LinkedList_AddsIf_FromLinkedList(t *testing.T) {
	tc := linkedListAddsIfFalseTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddsIf(false, 1, 2, 3)

	actLines := []string{fmt.Sprintf("%v", ll.IsEmpty())}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: AddFunc
// ==========================================================================

func Test_LinkedList_AddFunc_FromLinkedList(t *testing.T) {
	tc := linkedListAddFuncTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddFunc(func() int { return 99 })

	actLines := []string{fmt.Sprintf("%v", ll.First())}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: Push
// ==========================================================================

func Test_LinkedList_Push(t *testing.T) {
	tc := linkedListPushTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Push(1)
	ll.PushBack(2)
	ll.PushFront(0)

	actLines := []string{fmt.Sprintf("%v", ll.Length())}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: FirstOrDefault
// ==========================================================================

func Test_LinkedList_FirstOrDefaultEmpty(t *testing.T) {
	tc := linkedListFirstDefaultEmptyTestCase
	ll := coregeneric.EmptyLinkedList[int]()

	actLines := []string{fmt.Sprintf("%v", ll.FirstOrDefault())}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_FirstOrDefaultNonEmpty(t *testing.T) {
	tc := linkedListFirstDefaultNonEmptyTestCase
	ll := coregeneric.LinkedListFrom([]int{10, 20})

	actLines := []string{fmt.Sprintf("%v", ll.FirstOrDefault())}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: LastOrDefault
// ==========================================================================

func Test_LinkedList_LastOrDefaultEmpty(t *testing.T) {
	tc := linkedListLastDefaultEmptyTestCase
	ll := coregeneric.EmptyLinkedList[string]()

	actLines := []string{ll.LastOrDefault()}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_LastOrDefaultNonEmpty(t *testing.T) {
	tc := linkedListLastDefaultNonEmptyTestCase
	ll := coregeneric.LinkedListFrom([]int{10, 20})

	actLines := []string{fmt.Sprintf("%v", ll.LastOrDefault())}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: Items
// ==========================================================================

func Test_LinkedList_ItemsAll(t *testing.T) {
	tc := linkedListItemsAllTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})

	actLines := []string{fmt.Sprintf("%v", len(ll.Items()))}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_ItemsEmpty(t *testing.T) {
	tc := linkedListItemsEmptyTestCase
	ll := coregeneric.EmptyLinkedList[int]()

	actLines := []string{fmt.Sprintf("%v", len(ll.Items()))}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: Collection
// ==========================================================================

func Test_LinkedList_Collection_FromLinkedList(t *testing.T) {
	tc := linkedListCollectionTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2})

	actLines := []string{fmt.Sprintf("%v", ll.Collection().Length())}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: String
// ==========================================================================

func Test_LinkedList_String_FromLinkedList(t *testing.T) {
	tc := linkedListStringTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})

	actLines := []string{ll.String()}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: IndexAt
// ==========================================================================

func Test_LinkedList_IndexAt_Valid(t *testing.T) {
	tc := linkedListIndexAtValidTestCase
	ll := coregeneric.LinkedListFrom([]string{"a", "b", "c"})
	node := ll.IndexAt(1)

	// Act
	actual := args.Map{
		"isNotNil": node != nil,
		"value":    node.Element,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

func Test_LinkedList_IndexAt_First(t *testing.T) {
	tc := linkedListIndexAtFirstTestCase
	ll := coregeneric.LinkedListFrom([]int{10, 20})

	actLines := []string{fmt.Sprintf("%v", ll.IndexAt(0).Element)}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_IndexAt_Last(t *testing.T) {
	tc := linkedListIndexAtLastTestCase
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})

	actLines := []string{fmt.Sprintf("%v", ll.IndexAt(2).Element)}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_IndexAt_OutOfBounds(t *testing.T) {
	tc := linkedListIndexAtOutOfBoundsTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2})

	// Act
	actual := args.Map{
		"isNil":    ll.IndexAt(5) == nil,
		"hasError": ll.IndexAt(-1) == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

func Test_LinkedList_IndexAt_Empty(t *testing.T) {
	tc := linkedListIndexAtEmptyTestCase
	ll := coregeneric.EmptyLinkedList[int]()

	actLines := []string{fmt.Sprintf("%v", ll.IndexAt(0) == nil)}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: ForEach
// ==========================================================================

func Test_LinkedList_ForEachVisitsAll(t *testing.T) {
	tc := linkedListForEachVisitsAllTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
	sum := 0
	ll.ForEach(func(_ int, item int) { sum += item })

	actLines := []string{fmt.Sprintf("%v", sum)}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_ForEachEmpty(t *testing.T) {
	tc := linkedListForEachEmptyTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	called := false
	ll.ForEach(func(_ int, _ int) { called = true })

	actLines := []string{fmt.Sprintf("%v", called)}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: ForEachBreak
// ==========================================================================

func Test_LinkedList_ForEachBreakStopsEarly(t *testing.T) {
	tc := linkedListForEachBreakStopsEarlyTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3, 4, 5})
	count := 0
	ll.ForEachBreak(func(_ int, item int) bool { count++; return item == 3 })

	actLines := []string{fmt.Sprintf("%v", count)}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_ForEachBreakFirst(t *testing.T) {
	tc := linkedListForEachBreakFirstTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
	count := 0
	ll.ForEachBreak(func(_ int, _ int) bool { count++; return true })

	actLines := []string{fmt.Sprintf("%v", count)}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: Head / Tail
// ==========================================================================

func Test_LinkedList_HeadTail(t *testing.T) {
	tc := linkedListHeadTailTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{
		"head":        ll.Head().Element,
		"tail":        ll.Tail().Element,
		"headHasNext": ll.Head().HasNext(),
		"tailHasNext": ll.Tail().HasNext(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

func Test_LinkedList_NodeNext(t *testing.T) {
	tc := linkedListNodeNextTestCase
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})
	n := ll.Head()
	first := n.Element
	n = n.Next()
	second := n.Element
	n = n.Next()
	third := n.Element
	hasMore := n.HasNext()

	// Act
	actual := args.Map{
		"first":   first,
		"second":  second,
		"third":   third,
		"hasMore": hasMore,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

// ==========================================================================
// Test: Lock variants
// ==========================================================================

func Test_LinkedList_LengthLock_FromLinkedList(t *testing.T) {
	tc := linkedListLengthLockTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2})

	actLines := []string{fmt.Sprintf("%v", ll.LengthLock())}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_IsEmptyLock_FromLinkedList(t *testing.T) {
	tc := linkedListIsEmptyLockTestCase
	ll := coregeneric.EmptyLinkedList[int]()

	actLines := []string{fmt.Sprintf("%v", ll.IsEmptyLock())}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_AddLock_FromLinkedList(t *testing.T) {
	tc := linkedListAddLockTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddLock(1)
	ll.AddLock(2)

	actLines := []string{fmt.Sprintf("%v", ll.Length())}

	// Assert
	tc.ShouldBeEqualFirst(t, actLines...)
}

// Note: Nil receiver tests migrated to NilReceiver_test.go using CaseNilSafe pattern.

// ==========================================================================
// Test: AppendNode
// ==========================================================================

func Test_LinkedList_AppendNodeAppends(t *testing.T) {
	tc := linkedListAppendNodeAppendsTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	ll.AppendNode(&coregeneric.LinkedListNode[int]{Element: 3})

	// Act
	actual := args.Map{
		"length":    ll.Length(),
		"lastValue": ll.Last(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}

func Test_LinkedList_AppendNodeEmpty(t *testing.T) {
	tc := linkedListAppendNodeEmptyTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AppendNode(&coregeneric.LinkedListNode[int]{Element: 99})

	// Act
	actual := args.Map{
		"length": ll.Length(),
		"value":  ll.First(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(
		t,
		actual,
	)
}
