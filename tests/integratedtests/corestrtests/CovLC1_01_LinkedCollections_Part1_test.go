package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// LinkedCollections — Segment 4: Basic ops, Add, Loop, Filter, Remove (L1-800)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovLC1_01_Tail_Head(t *testing.T) {
	safeTest(t, "Test_CovLC1_01_Tail_Head", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c1)

		// Act
		actual := args.Map{"result": lc.Tail() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil tail", actual)
		actual = args.Map{"result": lc.Head() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil head", actual)
	})
}

func Test_CovLC1_02_First_Single_Last(t *testing.T) {
	safeTest(t, "Test_CovLC1_02_First_Single_Last", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c1)

		// Act
		actual := args.Map{"result": lc.First().Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": lc.Single().Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": lc.Last().Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLC1_03_LastOrDefault_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_CovLC1_03_LastOrDefault_FirstOrDefault", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()

		// Act
		actual := args.Map{"result": lc.LastOrDefault().Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": lc.FirstOrDefault().Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c1)
		actual = args.Map{"result": lc.LastOrDefault().Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": lc.FirstOrDefault().Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLC1_04_Length(t *testing.T) {
	safeTest(t, "Test_CovLC1_04_Length", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual = args.Map{"result": lc.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLC1_05_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_CovLC1_05_AllIndividualItemsLength", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))

		// Act
		actual := args.Map{"result": lc.AllIndividualItemsLength() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_CovLC1_06_LengthLock(t *testing.T) {
	safeTest(t, "Test_CovLC1_06_LengthLock", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLC1_07_IsEqualsPtr(t *testing.T) {
	safeTest(t, "Test_CovLC1_07_IsEqualsPtr", func() {
		// Arrange
		a := corestr.Empty.LinkedCollections()
		a.Add(corestr.New.Collection.Strings([]string{"x", "y"}))
		b := corestr.Empty.LinkedCollections()
		b.Add(corestr.New.Collection.Strings([]string{"x", "y"}))

		// same ptr

		// Act
		actual := args.Map{"result": a.IsEqualsPtr(a)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal to self", actual)
		// nil
		actual = args.Map{"result": a.IsEqualsPtr(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
		// both empty
		e1 := corestr.Empty.LinkedCollections()
		e2 := corestr.Empty.LinkedCollections()
		actual = args.Map{"result": e1.IsEqualsPtr(e2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty == empty", actual)
		// one empty
		actual = args.Map{"result": a.IsEqualsPtr(e1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// diff length
		c := corestr.Empty.LinkedCollections()
		c.Add(corestr.New.Collection.Strings([]string{"x"}))
		c.Add(corestr.New.Collection.Strings([]string{"y"}))
		// same content different structure
		actual = args.Map{"result": a.IsEqualsPtr(b) != true}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_CovLC1_08_IsEmptyLock_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_CovLC1_08_IsEmptyLock_IsEmpty_HasItems", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()

		// Act
		actual := args.Map{"result": lc.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": lc.IsEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": lc.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual = args.Map{"result": lc.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
		actual = args.Map{"result": lc.HasItems()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has items", actual)
	})
}

func Test_CovLC1_09_InsertAt(t *testing.T) {
	safeTest(t, "Test_CovLC1_09_InsertAt", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		// insert at front
		lc.InsertAt(0, corestr.New.Collection.Strings([]string{"front"}))

		// Act
		actual := args.Map{"result": lc.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// insert in middle
		lc.InsertAt(1, corestr.New.Collection.Strings([]string{"mid"}))
		actual = args.Map{"result": lc.Length() != 4}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_CovLC1_10_AddAsync(t *testing.T) {
	safeTest(t, "Test_CovLC1_10_AddAsync", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsync(corestr.New.Collection.Strings([]string{"a"}), wg)
		wg.Wait()

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLC1_11_AddsAsyncOnComplete(t *testing.T) {
	safeTest(t, "Test_CovLC1_11_AddsAsyncOnComplete", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		done := make(chan bool, 1)
		lc.AddsAsyncOnComplete(
			func(lc *corestr.LinkedCollections) { done <- true },
			true,
			corestr.New.Collection.Strings([]string{"a"}),
		)
		<-done

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLC1_12_AddsUsingProcessorAsyncOnComplete(t *testing.T) {
	safeTest(t, "Test_CovLC1_12_AddsUsingProcessorAsyncOnComplete", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		done := make(chan bool, 1)
		lc.AddsUsingProcessorAsyncOnComplete(
			func(lc *corestr.LinkedCollections) { done <- true },
			func(a any, i int) *corestr.Collection {
				return corestr.New.Collection.Strings([]string{a.(string)})
			},
			true,
			"hello", nil,
		)
		<-done

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)

		// nil anys with skip
		lc2 := corestr.Empty.LinkedCollections()
		done2 := make(chan bool, 1)
		lc2.AddsUsingProcessorAsyncOnComplete(
			func(lc *corestr.LinkedCollections) { done2 <- true },
			func(a any, i int) *corestr.Collection { return nil },
			true,
		)
		<-done2
	})
}

func Test_CovLC1_13_AddsUsingProcessorAsync(t *testing.T) {
	safeTest(t, "Test_CovLC1_13_AddsUsingProcessorAsync", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddsUsingProcessorAsync(
			wg,
			func(a any, i int) *corestr.Collection {
				return corestr.New.Collection.Strings([]string{a.(string)})
			},
			true,
			"x",
		)
		wg.Wait()

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)

		// nil anys with skip
		lc2 := corestr.Empty.LinkedCollections()
		wg2 := &sync.WaitGroup{}
		wg2.Add(1)
		lc2.AddsUsingProcessorAsync(wg2,
			func(a any, i int) *corestr.Collection { return nil },
			true,
		)
		wg2.Wait()
	})
}

func Test_CovLC1_14_AddLock(t *testing.T) {
	safeTest(t, "Test_CovLC1_14_AddLock", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.AddLock(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLC1_15_Add(t *testing.T) {
	safeTest(t, "Test_CovLC1_15_Add", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		// first add sets head
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		// second add sets tail
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))

		// Act
		actual := args.Map{"result": lc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLC1_16_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_CovLC1_16_AddStringsLock", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.AddStringsLock("a", "b")

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty
		lc.AddStringsLock()
	})
}

func Test_CovLC1_17_AddStrings(t *testing.T) {
	safeTest(t, "Test_CovLC1_17_AddStrings", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.AddStrings("a", "b")

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		lc.AddStrings()
	})
}

func Test_CovLC1_18_AddBackNode_AppendNode(t *testing.T) {
	safeTest(t, "Test_CovLC1_18_AddBackNode_AppendNode", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		lc.AddBackNode(node)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// append to non-empty
		node2 := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		lc.AppendNode(node2)
		actual = args.Map{"result": lc.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLC1_19_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_CovLC1_19_AppendChainOfNodes", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		node1 := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		// empty list
		lc.AppendChainOfNodes(node1)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// non-empty
		node2 := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		lc.AppendChainOfNodes(node2)
		actual = args.Map{"result": lc.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLC1_20_AppendChainOfNodesAsync(t *testing.T) {
	safeTest(t, "Test_CovLC1_20_AppendChainOfNodesAsync", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		lc.AppendChainOfNodesAsync(node, wg)
		wg.Wait()

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLC1_21_PushBackLock_PushBack_Push_PushFront(t *testing.T) {
	safeTest(t, "Test_CovLC1_21_PushBackLock_PushBack_Push_PushFront", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.PushBackLock(corestr.New.Collection.Strings([]string{"a"}))
		lc.PushBack(corestr.New.Collection.Strings([]string{"b"}))
		lc.Push(corestr.New.Collection.Strings([]string{"c"}))
		lc.PushFront(corestr.New.Collection.Strings([]string{"front"}))

		// Act
		actual := args.Map{"result": lc.Length() != 4}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_CovLC1_22_AddFrontLock_AddFront(t *testing.T) {
	safeTest(t, "Test_CovLC1_22_AddFrontLock_AddFront", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		// empty — falls through to Add
		lc.AddFront(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// non-empty
		lc.AddFrontLock(corestr.New.Collection.Strings([]string{"front"}))
		actual = args.Map{"result": lc.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLC1_23_AddAnother(t *testing.T) {
	safeTest(t, "Test_CovLC1_23_AddAnother", func() {
		// Arrange
		a := corestr.Empty.LinkedCollections()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		b := corestr.Empty.LinkedCollections()
		b.Add(corestr.New.Collection.Strings([]string{"b"}))
		b.Add(corestr.New.Collection.Strings([]string{"c"}))
		a.AddAnother(b)

		// Act
		actual := args.Map{"result": a.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// nil
		a.AddAnother(nil)
		// empty
		a.AddAnother(corestr.Empty.LinkedCollections())
	})
}

func Test_CovLC1_24_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_CovLC1_24_GetNextNodes", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		nodes := lc.GetNextNodes(2)

		// Act
		actual := args.Map{"result": len(nodes) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLC1_25_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_CovLC1_25_GetAllLinkedNodes", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		nodes := lc.GetAllLinkedNodes()

		// Act
		actual := args.Map{"result": len(nodes) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLC1_26_Loop(t *testing.T) {
	safeTest(t, "Test_CovLC1_26_Loop", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return false
		})

		// Act
		actual := args.Map{"result": count != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty
		corestr.Empty.LinkedCollections().Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			actual = args.Map{"result": false}
			expected = args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "should not be called", actual)
			return false
		})
		// break
		lc2 := corestr.Empty.LinkedCollections()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2.Add(corestr.New.Collection.Strings([]string{"b"}))
		breakCount := 0
		lc2.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			breakCount++
			return true // break on first
		})
		actual = args.Map{"result": breakCount != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLC1_27_Filter(t *testing.T) {
	safeTest(t, "Test_CovLC1_27_Filter", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		result := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{
				Value:   arg.Node,
				IsKeep:  true,
				IsBreak: false,
			}
		})

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty
		empty := corestr.Empty.LinkedCollections()
		r := empty.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		actual = args.Map{"result": len(r) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// break
		r2 := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})
		actual = args.Map{"result": len(r2) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLC1_28_FilterAsCollection(t *testing.T) {
	safeTest(t, "Test_CovLC1_28_FilterAsCollection", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		col := lc.FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		}, 0)

		// Act
		actual := args.Map{"result": col.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// empty result
		col2 := lc.FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: false}
		}, 0)
		actual = args.Map{"result": col2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovLC1_29_FilterAsCollections(t *testing.T) {
	safeTest(t, "Test_CovLC1_29_FilterAsCollections", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		cols := lc.FilterAsCollections(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})

		// Act
		actual := args.Map{"result": len(cols) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLC1_30_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_CovLC1_30_RemoveNodeByIndex", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		// remove first
		lc.RemoveNodeByIndex(0)

		// Act
		actual := args.Map{"result": lc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// remove last
		lc.RemoveNodeByIndex(1)
		actual = args.Map{"result": lc.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// remove middle (rebuild)
		lc2 := corestr.Empty.LinkedCollections()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc2.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc2.RemoveNodeByIndex(1)
		actual = args.Map{"result": lc2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLC1_31_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_CovLC1_31_RemoveNodeByIndexes", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc.RemoveNodeByIndexes(true, 0, 2)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty indexes
		lc.RemoveNodeByIndexes(true)
	})
}

func Test_CovLC1_32_RemoveNode(t *testing.T) {
	safeTest(t, "Test_CovLC1_32_RemoveNode", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.Add(c1)
		lc.Add(c2)
		head := lc.Head()
		lc.RemoveNode(head) // remove first

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// remove non-first
		lc2 := corestr.Empty.LinkedCollections()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2.Add(corestr.New.Collection.Strings([]string{"b"}))
		tail := lc2.Tail()
		lc2.RemoveNode(tail)
		actual = args.Map{"result": lc2.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLC1_33_AppendCollections(t *testing.T) {
	safeTest(t, "Test_CovLC1_33_AppendCollections", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		lc.AppendCollections(true, corestr.New.Collection.Strings([]string{"a"}), nil)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// nil with skip
		lc.AppendCollections(true)
	})
}

func Test_CovLC1_34_AppendCollectionsPointersLock(t *testing.T) {
	safeTest(t, "Test_CovLC1_34_AppendCollectionsPointersLock", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		cols := []*corestr.Collection{corestr.New.Collection.Strings([]string{"a"}), nil}
		lc.AppendCollectionsPointersLock(true, &cols)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		lc.AppendCollectionsPointersLock(true, nil)
	})
}

func Test_CovLC1_35_AppendCollectionsPointers(t *testing.T) {
	safeTest(t, "Test_CovLC1_35_AppendCollectionsPointers", func() {
		// Arrange
		lc := corestr.Empty.LinkedCollections()
		cols := []*corestr.Collection{corestr.New.Collection.Strings([]string{"a"}), nil}
		lc.AppendCollectionsPointers(true, &cols)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		lc.AppendCollectionsPointers(true, nil)
	})
}
