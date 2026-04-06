package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_Collection_Length_NilItems(t *testing.T) {
	var c *coregeneric.Collection[string]
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for nil collection", actual)
}

func Test_QW_LinkedList_IndexAt_EndOfList(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[string]()
	ll.Add("a")
	// Access index beyond list length — covers the out-of-range early return
	node := ll.IndexAt(5)
	actual := args.Map{"result": node != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for out-of-range index", actual)
}

func Test_QW_MinOf_SecondSmaller(t *testing.T) {
	// Cover the else branch (a >= b)
	result := coregeneric.MinOf(5, 3)
	actual := args.Map{"result": result != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_QW_MaxOf_SecondLarger(t *testing.T) {
	// Cover the else branch (a <= b)
	result := coregeneric.MaxOf(3, 5)
	actual := args.Map{"result": result != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_QW_MinOfSlice_NonMinElements(t *testing.T) {
	// Cover the case where v < result is false
	result := coregeneric.MinOfSlice([]int{3, 5, 1, 4})
	actual := args.Map{"result": result != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}
