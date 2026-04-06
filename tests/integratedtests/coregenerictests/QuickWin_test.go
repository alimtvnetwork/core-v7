package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_Collection_Length_NilItems(t *testing.T) {
	// Arrange
	var c *coregeneric.Collection[string]

	// Act
	actual := args.Map{"result": c.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for nil collection", actual)
}

func Test_QW_LinkedList_IndexAt_EndOfList(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[string]()
	ll.Add("a")
	// Access index beyond list length — covers the out-of-range early return
	node := ll.IndexAt(5)

	// Act
	actual := args.Map{"result": node != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for out-of-range index", actual)
}

func Test_QW_MinOf_SecondSmaller(t *testing.T) {
	// Arrange
	// Cover the else branch (a >= b)
	result := coregeneric.MinOf(5, 3)

	// Act
	actual := args.Map{"result": result != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_QW_MaxOf_SecondLarger(t *testing.T) {
	// Arrange
	// Cover the else branch (a <= b)
	result := coregeneric.MaxOf(3, 5)

	// Act
	actual := args.Map{"result": result != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_QW_MinOfSlice_NonMinElements(t *testing.T) {
	// Arrange
	// Cover the case where v < result is false
	result := coregeneric.MinOfSlice([]int{3, 5, 1, 4})

	// Act
	actual := args.Map{"result": result != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}
