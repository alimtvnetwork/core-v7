package corestrtests

import (
	"strings"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════
// S13 — LinkedCollections.go (1,551 lines) — Full coverage
// ══════════════════════════════════════════════════════════════

func Test_S13_01_LinkedCollections_HeadTailLength(t *testing.T) {
	safeTest(t, "Test_S13_01_LinkedCollections_HeadTailLength", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(col)

		// Act & Assert
		actual := args.Map{"result": lc.Head() == nil || lc.Tail() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		actual = args.Map{"result": lc.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_02_LinkedCollections_LengthLock(t *testing.T) {
	safeTest(t, "Test_S13_02_LinkedCollections_LengthLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_03_LinkedCollections_FirstSingleLast(t *testing.T) {
	safeTest(t, "Test_S13_03_LinkedCollections_FirstSingleLast", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(col)

		// Act
		actual := args.Map{"result": lc.First() == nil || lc.Single() == nil || lc.Last() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S13_04_LinkedCollections_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_S13_04_LinkedCollections_FirstOrDefault_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": lc.FirstOrDefault() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S13_05_LinkedCollections_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_S13_05_LinkedCollections_LastOrDefault_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": lc.LastOrDefault() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S13_06_LinkedCollections_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_S13_06_LinkedCollections_IsEmpty_HasItems", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": lc.IsEmpty() || lc.HasItems()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual = args.Map{"result": lc.IsEmpty() || !lc.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_S13_07_LinkedCollections_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_S13_07_LinkedCollections_IsEmptyLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": lc.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_S13_08_LinkedCollections_Add_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S13_08_LinkedCollections_Add_OnEmpty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(col)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_09_LinkedCollections_Add_Multiple(t *testing.T) {
	safeTest(t, "Test_S13_09_LinkedCollections_Add_Multiple", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))

		// Act
		actual := args.Map{"result": lc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S13_10_LinkedCollections_AddLock(t *testing.T) {
	safeTest(t, "Test_S13_10_LinkedCollections_AddLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddLock(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_11_LinkedCollections_AddStrings(t *testing.T) {
	safeTest(t, "Test_S13_11_LinkedCollections_AddStrings", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_12_LinkedCollections_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_S13_12_LinkedCollections_AddStrings_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings()

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_13_LinkedCollections_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_S13_13_LinkedCollections_AddStringsLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock("a")

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_14_LinkedCollections_AddStringsLock_Empty(t *testing.T) {
	safeTest(t, "Test_S13_14_LinkedCollections_AddStringsLock_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock()

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_15_LinkedCollections_AddFront(t *testing.T) {
	safeTest(t, "Test_S13_15_LinkedCollections_AddFront", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.AddFront(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.First().Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_16_LinkedCollections_AddFront_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S13_16_LinkedCollections_AddFront_OnEmpty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddFront(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_17_LinkedCollections_AddFrontLock(t *testing.T) {
	safeTest(t, "Test_S13_17_LinkedCollections_AddFrontLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddFrontLock(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_18_LinkedCollections_Push_PushBack_PushFront_PushBackLock(t *testing.T) {
	safeTest(t, "Test_S13_18_LinkedCollections_Push_PushBack_PushFront_PushBackLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Push(col)
		lc.PushBack(corestr.New.Collection.Strings([]string{"b"}))
		lc.PushFront(corestr.New.Collection.Strings([]string{"z"}))
		lc.PushBackLock(corestr.New.Collection.Strings([]string{"c"}))

		// Act
		actual := args.Map{"result": lc.Length() != 4}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_S13_19_LinkedCollections_AppendNode_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S13_19_LinkedCollections_AppendNode_OnEmpty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		lc.AppendNode(node)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_20_LinkedCollections_AddBackNode(t *testing.T) {
	safeTest(t, "Test_S13_20_LinkedCollections_AddBackNode", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		lc.AddBackNode(node)

		// Act
		actual := args.Map{"result": lc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S13_21_LinkedCollections_AddAnother(t *testing.T) {
	safeTest(t, "Test_S13_21_LinkedCollections_AddAnother", func() {
		// Arrange
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		b := corestr.New.LinkedCollection.Create()
		b.Add(corestr.New.Collection.Strings([]string{"b"}))
		a.AddAnother(b)

		// Act
		actual := args.Map{"result": a.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S13_22_LinkedCollections_AddAnother_Nil(t *testing.T) {
	safeTest(t, "Test_S13_22_LinkedCollections_AddAnother_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAnother(nil)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_23_LinkedCollections_AddCollection(t *testing.T) {
	safeTest(t, "Test_S13_23_LinkedCollections_AddCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollection(corestr.New.Collection.Strings([]string{"a"}))
		lc.AddCollection(nil)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_24_LinkedCollections_Loop(t *testing.T) {
	safeTest(t, "Test_S13_24_LinkedCollections_Loop", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
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
	})
}

func Test_S13_25_LinkedCollections_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_S13_25_LinkedCollections_Loop_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {

		// Act
			actual := args.Map{"result": false}

		// Assert
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "should not be called", actual)
			return false
		})
	})
}

func Test_S13_26_LinkedCollections_Loop_Break(t *testing.T) {
	safeTest(t, "Test_S13_26_LinkedCollections_Loop_Break", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return true
		})

		// Act
		actual := args.Map{"result": count != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_27_LinkedCollections_Filter(t *testing.T) {
	safeTest(t, "Test_S13_27_LinkedCollections_Filter", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		nodes := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})

		// Act
		actual := args.Map{"result": len(nodes) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S13_28_LinkedCollections_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_S13_28_LinkedCollections_Filter_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		nodes := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})

		// Act
		actual := args.Map{"result": len(nodes) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_29_LinkedCollections_Filter_Break(t *testing.T) {
	safeTest(t, "Test_S13_29_LinkedCollections_Filter_Break", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		nodes := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})

		// Act
		actual := args.Map{"result": len(nodes) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_30_LinkedCollections_FilterAsCollection(t *testing.T) {
	safeTest(t, "Test_S13_30_LinkedCollections_FilterAsCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		result := lc.FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		}, 0)

		// Act
		actual := args.Map{"result": result.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S13_31_LinkedCollections_FilterAsCollections(t *testing.T) {
	safeTest(t, "Test_S13_31_LinkedCollections_FilterAsCollections", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		result := lc.FilterAsCollections(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_32_LinkedCollections_IsEqualsPtr(t *testing.T) {
	safeTest(t, "Test_S13_32_LinkedCollections_IsEqualsPtr", func() {
		// Arrange
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		b := corestr.New.LinkedCollection.Create()
		b.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": a.IsEqualsPtr(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S13_33_LinkedCollections_IsEqualsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_S13_33_LinkedCollections_IsEqualsPtr_Nil", func() {
		// Arrange
		a := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": a.IsEqualsPtr(nil)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S13_34_LinkedCollections_IsEqualsPtr_SamePtr(t *testing.T) {
	safeTest(t, "Test_S13_34_LinkedCollections_IsEqualsPtr_SamePtr", func() {
		// Arrange
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": a.IsEqualsPtr(a)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S13_35_LinkedCollections_IsEqualsPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S13_35_LinkedCollections_IsEqualsPtr_BothEmpty", func() {
		// Arrange
		a := corestr.New.LinkedCollection.Create()
		b := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": a.IsEqualsPtr(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S13_36_LinkedCollections_IsEqualsPtr_DiffLength(t *testing.T) {
	safeTest(t, "Test_S13_36_LinkedCollections_IsEqualsPtr_DiffLength", func() {
		// Arrange
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		b := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": a.IsEqualsPtr(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S13_37_LinkedCollections_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_S13_37_LinkedCollections_AllIndividualItemsLength", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))

		// Act
		actual := args.Map{"result": lc.AllIndividualItemsLength() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_S13_38_LinkedCollections_AppendCollections(t *testing.T) {
	safeTest(t, "Test_S13_38_LinkedCollections_AppendCollections", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.AppendCollections(true, c1, nil, c2)

		// Act
		actual := args.Map{"result": lc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S13_39_LinkedCollections_AppendCollections_NilSlice(t *testing.T) {
	safeTest(t, "Test_S13_39_LinkedCollections_AppendCollections_NilSlice", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AppendCollections(true, nil...)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_40_LinkedCollections_AddStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_S13_40_LinkedCollections_AddStringsOfStrings", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false, []string{"a"}, nil, []string{"b"})

		// Act
		actual := args.Map{"result": lc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S13_41_LinkedCollections_AddStringsOfStrings_Empty(t *testing.T) {
	safeTest(t, "Test_S13_41_LinkedCollections_AddStringsOfStrings_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_42_LinkedCollections_ConcatNew(t *testing.T) {
	safeTest(t, "Test_S13_42_LinkedCollections_ConcatNew", func() {
		// Arrange
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		b := corestr.New.LinkedCollection.Create()
		b.Add(corestr.New.Collection.Strings([]string{"b"}))
		result := a.ConcatNew(true, b)

		// Act
		actual := args.Map{"result": result.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_S13_43_LinkedCollections_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_S13_43_LinkedCollections_ConcatNew_EmptyClone", func() {
		// Arrange
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		result := a.ConcatNew(true)

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_44_LinkedCollections_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_S13_44_LinkedCollections_ConcatNew_EmptyNoClone", func() {
		// Arrange
		a := corestr.New.LinkedCollection.Create()
		result := a.ConcatNew(false)

		// Act
		actual := args.Map{"result": result != a}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same pointer", actual)
	})
}

func Test_S13_45_LinkedCollections_ToCollection(t *testing.T) {
	safeTest(t, "Test_S13_45_LinkedCollections_ToCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		col := lc.ToCollection(0)

		// Act
		actual := args.Map{"result": col.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S13_46_LinkedCollections_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_S13_46_LinkedCollections_ToCollection_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		col := lc.ToCollection(0)

		// Act
		actual := args.Map{"result": col.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_47_LinkedCollections_ToCollectionSimple(t *testing.T) {
	safeTest(t, "Test_S13_47_LinkedCollections_ToCollectionSimple", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.ToCollectionSimple().Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_48_LinkedCollections_ToStrings(t *testing.T) {
	safeTest(t, "Test_S13_48_LinkedCollections_ToStrings", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": len(lc.ToStrings()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_49_LinkedCollections_ToStringsPtr(t *testing.T) {
	safeTest(t, "Test_S13_49_LinkedCollections_ToStringsPtr", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.ToStringsPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S13_50_LinkedCollections_ToCollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_S13_50_LinkedCollections_ToCollectionsOfCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		coc := lc.ToCollectionsOfCollection(0)

		// Act
		actual := args.Map{"result": coc == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S13_51_LinkedCollections_ToCollectionsOfCollection_Empty(t *testing.T) {
	safeTest(t, "Test_S13_51_LinkedCollections_ToCollectionsOfCollection_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		coc := lc.ToCollectionsOfCollection(0)

		// Act
		actual := args.Map{"result": coc == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S13_52_LinkedCollections_ItemsOfItems(t *testing.T) {
	safeTest(t, "Test_S13_52_LinkedCollections_ItemsOfItems", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		items := lc.ItemsOfItems()

		// Act
		actual := args.Map{"result": len(items) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_53_LinkedCollections_ItemsOfItems_Empty(t *testing.T) {
	safeTest(t, "Test_S13_53_LinkedCollections_ItemsOfItems_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": len(lc.ItemsOfItems()) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_54_LinkedCollections_ItemsOfItemsCollection(t *testing.T) {
	safeTest(t, "Test_S13_54_LinkedCollections_ItemsOfItemsCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		items := lc.ItemsOfItemsCollection()

		// Act
		actual := args.Map{"result": len(items) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_55_LinkedCollections_ItemsOfItemsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_S13_55_LinkedCollections_ItemsOfItemsCollection_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": len(lc.ItemsOfItemsCollection()) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_56_LinkedCollections_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_S13_56_LinkedCollections_SimpleSlice", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		ss := lc.SimpleSlice()

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_57_LinkedCollections_List(t *testing.T) {
	safeTest(t, "Test_S13_57_LinkedCollections_List", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": len(lc.List()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_58_LinkedCollections_List_Empty(t *testing.T) {
	safeTest(t, "Test_S13_58_LinkedCollections_List_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": len(lc.List()) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_59_LinkedCollections_ListPtr(t *testing.T) {
	safeTest(t, "Test_S13_59_LinkedCollections_ListPtr", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.ListPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S13_60_LinkedCollections_String(t *testing.T) {
	safeTest(t, "Test_S13_60_LinkedCollections_String", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_S13_61_LinkedCollections_String_Empty(t *testing.T) {
	safeTest(t, "Test_S13_61_LinkedCollections_String_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": strings.Contains(lc.String(), "No Element")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected No Element", actual)
	})
}

func Test_S13_62_LinkedCollections_StringLock(t *testing.T) {
	safeTest(t, "Test_S13_62_LinkedCollections_StringLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.StringLock() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_S13_63_LinkedCollections_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_S13_63_LinkedCollections_StringLock_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": strings.Contains(lc.StringLock(), "No Element")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected No Element", actual)
	})
}

func Test_S13_64_LinkedCollections_Join(t *testing.T) {
	safeTest(t, "Test_S13_64_LinkedCollections_Join", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.Join(",") == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_S13_65_LinkedCollections_Joins(t *testing.T) {
	safeTest(t, "Test_S13_65_LinkedCollections_Joins", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		result := lc.Joins(",", "b")

		// Act
		actual := args.Map{"result": strings.Contains(result, "a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_S13_66_LinkedCollections_Joins_NilItems(t *testing.T) {
	safeTest(t, "Test_S13_66_LinkedCollections_Joins_NilItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		_ = lc.Joins(",", nil...)
	})
}

func Test_S13_67_LinkedCollections_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_S13_67_LinkedCollections_MarshalJSON", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		data, err := lc.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil || len(data) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid JSON", actual)
	})
}

func Test_S13_68_LinkedCollections_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_S13_68_LinkedCollections_UnmarshalJSON", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		err := lc.UnmarshalJSON([]byte(`["a","b"]`))

		// Act
		actual := args.Map{"result": err != nil || lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_69_LinkedCollections_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_S13_69_LinkedCollections_UnmarshalJSON_Invalid", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		err := lc.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_S13_70_LinkedCollections_JsonModel(t *testing.T) {
	safeTest(t, "Test_S13_70_LinkedCollections_JsonModel", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": len(lc.JsonModel()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_71_LinkedCollections_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_S13_71_LinkedCollections_JsonModelAny", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.JsonModelAny() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S13_72_LinkedCollections_Clear_RemoveAll(t *testing.T) {
	safeTest(t, "Test_S13_72_LinkedCollections_Clear_RemoveAll", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Clear()

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_73_LinkedCollections_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_S13_73_LinkedCollections_Clear_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Clear()

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_74_LinkedCollections_RemoveAll(t *testing.T) {
	safeTest(t, "Test_S13_74_LinkedCollections_RemoveAll", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.RemoveAll()

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_75_LinkedCollections_Json(t *testing.T) {
	safeTest(t, "Test_S13_75_LinkedCollections_Json", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		jsonResult := lc.Json()

		// Act
		actual := args.Map{"result": jsonResult.HasError()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_S13_76_LinkedCollections_JsonPtr(t *testing.T) {
	safeTest(t, "Test_S13_76_LinkedCollections_JsonPtr", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.JsonPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S13_77_LinkedCollections_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_S13_77_LinkedCollections_ParseInjectUsingJson", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		jsonResult := lc.JsonPtr()
		target := corestr.New.LinkedCollection.Create()
		result, err := target.ParseInjectUsingJson(jsonResult)

		// Act
		actual := args.Map{"result": err != nil || result.Length() < 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_S13_78_LinkedCollections_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_S13_78_LinkedCollections_ParseInjectUsingJsonMust", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		jsonResult := lc.JsonPtr()
		target := corestr.New.LinkedCollection.Create()
		result := target.ParseInjectUsingJsonMust(jsonResult)

		// Act
		actual := args.Map{"result": result.Length() < 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_S13_79_LinkedCollections_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_S13_79_LinkedCollections_JsonParseSelfInject", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		jsonResult := lc.JsonPtr()
		target := corestr.New.LinkedCollection.Create()
		err := target.JsonParseSelfInject(jsonResult)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_S13_80_LinkedCollections_AsJsoner(t *testing.T) {
	safeTest(t, "Test_S13_80_LinkedCollections_AsJsoner", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": lc.AsJsoner() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S13_81_LinkedCollections_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_S13_81_LinkedCollections_AsJsonContractsBinder", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": lc.AsJsonContractsBinder() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S13_82_LinkedCollections_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_S13_82_LinkedCollections_AsJsonParseSelfInjector", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": lc.AsJsonParseSelfInjector() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S13_83_LinkedCollections_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_S13_83_LinkedCollections_AsJsonMarshaller", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": lc.AsJsonMarshaller() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S13_84_LinkedCollections_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_S13_84_LinkedCollections_GetCompareSummary", func() {
		// Arrange
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"x"}))
		b := corestr.New.LinkedCollection.Create()
		b.Add(corestr.New.Collection.Strings([]string{"y"}))
		summary := a.GetCompareSummary(b, "left", "right")

		// Act
		actual := args.Map{"result": summary == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_S13_85_LinkedCollections_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_S13_85_LinkedCollections_GetNextNodes", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
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

func Test_S13_86_LinkedCollections_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_S13_86_LinkedCollections_GetAllLinkedNodes", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": len(lc.GetAllLinkedNodes()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_87_LinkedCollections_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_S13_87_LinkedCollections_SafeIndexAt", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))

		// Act
		actual := args.Map{"result": lc.SafeIndexAt(1) == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		actual = args.Map{"result": lc.SafeIndexAt(-1) != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		actual = args.Map{"result": lc.SafeIndexAt(10) != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_S13_88_LinkedCollections_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_S13_88_LinkedCollections_SafePointerIndexAt", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.SafePointerIndexAt(0) == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		actual = args.Map{"result": lc.SafePointerIndexAt(10) != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_S13_89_LinkedCollections_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_S13_89_LinkedCollections_RemoveNodeByIndex", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc.RemoveNodeByIndex(1)

		// Act
		actual := args.Map{"result": lc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S13_90_LinkedCollections_RemoveNode(t *testing.T) {
	safeTest(t, "Test_S13_90_LinkedCollections_RemoveNode", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		node := lc.Head()
		lc.RemoveNode(node)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_91_LinkedCollections_AddAsync(t *testing.T) {
	safeTest(t, "Test_S13_91_LinkedCollections_AddAsync", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
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

func Test_S13_92_LinkedCollections_AddStringsAsync(t *testing.T) {
	safeTest(t, "Test_S13_92_LinkedCollections_AddStringsAsync", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddStringsAsync(wg, []string{"a"})
		wg.Wait()

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_93_LinkedCollections_AddStringsAsync_Nil(t *testing.T) {
	safeTest(t, "Test_S13_93_LinkedCollections_AddStringsAsync_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsAsync(nil, nil)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_94_LinkedCollections_AddCollectionsPtr(t *testing.T) {
	safeTest(t, "Test_S13_94_LinkedCollections_AddCollectionsPtr", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		cols := []*corestr.Collection{corestr.New.Collection.Strings([]string{"a"})}
		lc.AddCollectionsPtr(cols)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_95_LinkedCollections_AddCollectionsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_S13_95_LinkedCollections_AddCollectionsPtr_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollectionsPtr(nil)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_96_LinkedCollections_AddCollections(t *testing.T) {
	safeTest(t, "Test_S13_96_LinkedCollections_AddCollections", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		cols := []*corestr.Collection{nil, corestr.New.Collection.Strings([]string{"a"})}
		lc.AddCollections(cols)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_97_LinkedCollections_AddCollections_Empty(t *testing.T) {
	safeTest(t, "Test_S13_97_LinkedCollections_AddCollections_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollections(nil)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_98_LinkedCollections_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_S13_98_LinkedCollections_AppendChainOfNodes", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		other := corestr.New.LinkedCollection.Create()
		other.Add(corestr.New.Collection.Strings([]string{"b"}))
		other.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc.AppendChainOfNodes(other.Head())

		// Act
		actual := args.Map{"result": lc.Length() < 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	})
}

func Test_S13_99_LinkedCollections_AppendChainOfNodes_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S13_99_LinkedCollections_AppendChainOfNodes_OnEmpty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		other := corestr.New.LinkedCollection.Create()
		other.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.AppendChainOfNodes(other.Head())

		// Act
		actual := args.Map{"result": lc.Length() < 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_S13_100_LinkedCollections_InsertAt(t *testing.T) {
	safeTest(t, "Test_S13_100_LinkedCollections_InsertAt", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc.InsertAt(1, corestr.New.Collection.Strings([]string{"b"}))

		// Act
		actual := args.Map{"result": lc.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_S13_101_LinkedCollections_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_S13_101_LinkedCollections_InsertAt_Front", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.InsertAt(0, corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": lc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S13_102_LinkedCollections_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_S13_102_LinkedCollections_RemoveNodeByIndexes", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc.RemoveNodeByIndexes(false, 1)

		// Act
		actual := args.Map{"result": lc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S13_103_LinkedCollections_RemoveNodeByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_S13_103_LinkedCollections_RemoveNodeByIndexes_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.RemoveNodeByIndexes(false)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_104_LinkedCollections_AddAsyncFuncItems(t *testing.T) {
	safeTest(t, "Test_S13_104_LinkedCollections_AddAsyncFuncItems", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItems(wg, false, func() []string { return []string{"a"} })

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_105_LinkedCollections_AddAsyncFuncItems_Nil(t *testing.T) {
	safeTest(t, "Test_S13_105_LinkedCollections_AddAsyncFuncItems_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAsyncFuncItems(nil, false, nil...)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_106_LinkedCollections_AddAsyncFuncItemsPointer(t *testing.T) {
	safeTest(t, "Test_S13_106_LinkedCollections_AddAsyncFuncItemsPointer", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItemsPointer(wg, false, func() []string { return []string{"a"} })

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S13_107_LinkedCollections_AddAsyncFuncItemsPointer_Nil(t *testing.T) {
	safeTest(t, "Test_S13_107_LinkedCollections_AddAsyncFuncItemsPointer_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAsyncFuncItemsPointer(nil, false, nil...)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_108_LinkedCollections_AttachWithNode(t *testing.T) {
	safeTest(t, "Test_S13_108_LinkedCollections_AttachWithNode", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		node := lc.Head()
		addNode := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		err := lc.AttachWithNode(node, addNode)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_S13_109_LinkedCollections_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_S13_109_LinkedCollections_AttachWithNode_NilCurrent", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		addNode := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		err := lc.AttachWithNode(nil, addNode)

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_S13_110_LinkedCollections_AddCollectionsToNode(t *testing.T) {
	safeTest(t, "Test_S13_110_LinkedCollections_AddCollectionsToNode", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.AddCollectionsToNode(true, lc.Head(), corestr.New.Collection.Strings([]string{"b"}))

		// Act
		actual := args.Map{"result": lc.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_S13_111_LinkedCollections_AddCollectionsToNode_Nil(t *testing.T) {
	safeTest(t, "Test_S13_111_LinkedCollections_AddCollectionsToNode_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollectionsToNode(true, nil, nil...)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S13_112_LinkedCollections_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_S13_112_LinkedCollections_AddCollectionToNode", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		col := corestr.New.Collection.Strings([]string{"b"})
		lc.AddCollectionToNode(true, lc.Head(), col)

		// Act
		actual := args.Map{"result": lc.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}
