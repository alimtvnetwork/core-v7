package corestrtests

import (
	"strings"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
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
		if lc.Head() == nil || lc.Tail() == nil {
			t.Fatal("expected non-nil")
		}
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_02_LinkedCollections_LengthLock(t *testing.T) {
	safeTest(t, "Test_S13_02_LinkedCollections_LengthLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.LengthLock() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_03_LinkedCollections_FirstSingleLast(t *testing.T) {
	safeTest(t, "Test_S13_03_LinkedCollections_FirstSingleLast", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(col)
		if lc.First() == nil || lc.Single() == nil || lc.Last() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S13_04_LinkedCollections_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_S13_04_LinkedCollections_FirstOrDefault_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.FirstOrDefault() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S13_05_LinkedCollections_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_S13_05_LinkedCollections_LastOrDefault_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.LastOrDefault() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S13_06_LinkedCollections_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_S13_06_LinkedCollections_IsEmpty_HasItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		if !lc.IsEmpty() || lc.HasItems() {
			t.Fatal("expected empty")
		}
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.IsEmpty() || !lc.HasItems() {
			t.Fatal("expected not empty")
		}
	})
}

func Test_S13_07_LinkedCollections_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_S13_07_LinkedCollections_IsEmptyLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		if !lc.IsEmptyLock() {
			t.Fatal("expected empty")
		}
	})
}

func Test_S13_08_LinkedCollections_Add_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S13_08_LinkedCollections_Add_OnEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(col)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_09_LinkedCollections_Add_Multiple(t *testing.T) {
	safeTest(t, "Test_S13_09_LinkedCollections_Add_Multiple", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S13_10_LinkedCollections_AddLock(t *testing.T) {
	safeTest(t, "Test_S13_10_LinkedCollections_AddLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddLock(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_11_LinkedCollections_AddStrings(t *testing.T) {
	safeTest(t, "Test_S13_11_LinkedCollections_AddStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_12_LinkedCollections_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_S13_12_LinkedCollections_AddStrings_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings()
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_13_LinkedCollections_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_S13_13_LinkedCollections_AddStringsLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock("a")
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_14_LinkedCollections_AddStringsLock_Empty(t *testing.T) {
	safeTest(t, "Test_S13_14_LinkedCollections_AddStringsLock_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock()
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_15_LinkedCollections_AddFront(t *testing.T) {
	safeTest(t, "Test_S13_15_LinkedCollections_AddFront", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.AddFront(corestr.New.Collection.Strings([]string{"a"}))
		if lc.First().Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_16_LinkedCollections_AddFront_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S13_16_LinkedCollections_AddFront_OnEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddFront(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_17_LinkedCollections_AddFrontLock(t *testing.T) {
	safeTest(t, "Test_S13_17_LinkedCollections_AddFrontLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddFrontLock(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_18_LinkedCollections_Push_PushBack_PushFront_PushBackLock(t *testing.T) {
	safeTest(t, "Test_S13_18_LinkedCollections_Push_PushBack_PushFront_PushBackLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a"})
		lc.Push(col)
		lc.PushBack(corestr.New.Collection.Strings([]string{"b"}))
		lc.PushFront(corestr.New.Collection.Strings([]string{"z"}))
		lc.PushBackLock(corestr.New.Collection.Strings([]string{"c"}))
		if lc.Length() != 4 {
			t.Fatalf("expected 4, got %d", lc.Length())
		}
	})
}

func Test_S13_19_LinkedCollections_AppendNode_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S13_19_LinkedCollections_AppendNode_OnEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		lc.AppendNode(node)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_20_LinkedCollections_AddBackNode(t *testing.T) {
	safeTest(t, "Test_S13_20_LinkedCollections_AddBackNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		lc.AddBackNode(node)
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S13_21_LinkedCollections_AddAnother(t *testing.T) {
	safeTest(t, "Test_S13_21_LinkedCollections_AddAnother", func() {
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		b := corestr.New.LinkedCollection.Create()
		b.Add(corestr.New.Collection.Strings([]string{"b"}))
		a.AddAnother(b)
		if a.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S13_22_LinkedCollections_AddAnother_Nil(t *testing.T) {
	safeTest(t, "Test_S13_22_LinkedCollections_AddAnother_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAnother(nil)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_23_LinkedCollections_AddCollection(t *testing.T) {
	safeTest(t, "Test_S13_23_LinkedCollections_AddCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollection(corestr.New.Collection.Strings([]string{"a"}))
		lc.AddCollection(nil)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_24_LinkedCollections_Loop(t *testing.T) {
	safeTest(t, "Test_S13_24_LinkedCollections_Loop", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return false
		})
		if count != 2 {
			t.Fatalf("expected 2, got %d", count)
		}
	})
}

func Test_S13_25_LinkedCollections_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_S13_25_LinkedCollections_Loop_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			t.Fatal("should not be called")
			return false
		})
	})
}

func Test_S13_26_LinkedCollections_Loop_Break(t *testing.T) {
	safeTest(t, "Test_S13_26_LinkedCollections_Loop_Break", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return true
		})
		if count != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_27_LinkedCollections_Filter(t *testing.T) {
	safeTest(t, "Test_S13_27_LinkedCollections_Filter", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		nodes := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})
		if len(nodes) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S13_28_LinkedCollections_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_S13_28_LinkedCollections_Filter_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		nodes := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		if len(nodes) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_29_LinkedCollections_Filter_Break(t *testing.T) {
	safeTest(t, "Test_S13_29_LinkedCollections_Filter_Break", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		nodes := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})
		if len(nodes) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_30_LinkedCollections_FilterAsCollection(t *testing.T) {
	safeTest(t, "Test_S13_30_LinkedCollections_FilterAsCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		result := lc.FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		}, 0)
		if result.Length() != 2 {
			t.Fatalf("expected 2, got %d", result.Length())
		}
	})
}

func Test_S13_31_LinkedCollections_FilterAsCollections(t *testing.T) {
	safeTest(t, "Test_S13_31_LinkedCollections_FilterAsCollections", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		result := lc.FilterAsCollections(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})
		if len(result) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_32_LinkedCollections_IsEqualsPtr(t *testing.T) {
	safeTest(t, "Test_S13_32_LinkedCollections_IsEqualsPtr", func() {
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		b := corestr.New.LinkedCollection.Create()
		b.Add(corestr.New.Collection.Strings([]string{"a"}))
		if !a.IsEqualsPtr(b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S13_33_LinkedCollections_IsEqualsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_S13_33_LinkedCollections_IsEqualsPtr_Nil", func() {
		a := corestr.New.LinkedCollection.Create()
		if a.IsEqualsPtr(nil) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S13_34_LinkedCollections_IsEqualsPtr_SamePtr(t *testing.T) {
	safeTest(t, "Test_S13_34_LinkedCollections_IsEqualsPtr_SamePtr", func() {
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		if !a.IsEqualsPtr(a) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S13_35_LinkedCollections_IsEqualsPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S13_35_LinkedCollections_IsEqualsPtr_BothEmpty", func() {
		a := corestr.New.LinkedCollection.Create()
		b := corestr.New.LinkedCollection.Create()
		if !a.IsEqualsPtr(b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S13_36_LinkedCollections_IsEqualsPtr_DiffLength(t *testing.T) {
	safeTest(t, "Test_S13_36_LinkedCollections_IsEqualsPtr_DiffLength", func() {
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		b := corestr.New.LinkedCollection.Create()
		if a.IsEqualsPtr(b) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S13_37_LinkedCollections_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_S13_37_LinkedCollections_AllIndividualItemsLength", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		if lc.AllIndividualItemsLength() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_S13_38_LinkedCollections_AppendCollections(t *testing.T) {
	safeTest(t, "Test_S13_38_LinkedCollections_AppendCollections", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.AppendCollections(true, c1, nil, c2)
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S13_39_LinkedCollections_AppendCollections_NilSlice(t *testing.T) {
	safeTest(t, "Test_S13_39_LinkedCollections_AppendCollections_NilSlice", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AppendCollections(true, nil...)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_40_LinkedCollections_AddStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_S13_40_LinkedCollections_AddStringsOfStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false, []string{"a"}, nil, []string{"b"})
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S13_41_LinkedCollections_AddStringsOfStrings_Empty(t *testing.T) {
	safeTest(t, "Test_S13_41_LinkedCollections_AddStringsOfStrings_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_42_LinkedCollections_ConcatNew(t *testing.T) {
	safeTest(t, "Test_S13_42_LinkedCollections_ConcatNew", func() {
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		b := corestr.New.LinkedCollection.Create()
		b.Add(corestr.New.Collection.Strings([]string{"b"}))
		result := a.ConcatNew(true, b)
		if result.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_S13_43_LinkedCollections_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_S13_43_LinkedCollections_ConcatNew_EmptyClone", func() {
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		result := a.ConcatNew(true)
		if result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_44_LinkedCollections_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_S13_44_LinkedCollections_ConcatNew_EmptyNoClone", func() {
		a := corestr.New.LinkedCollection.Create()
		result := a.ConcatNew(false)
		if result != a {
			t.Fatal("expected same pointer")
		}
	})
}

func Test_S13_45_LinkedCollections_ToCollection(t *testing.T) {
	safeTest(t, "Test_S13_45_LinkedCollections_ToCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		col := lc.ToCollection(0)
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S13_46_LinkedCollections_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_S13_46_LinkedCollections_ToCollection_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := lc.ToCollection(0)
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_47_LinkedCollections_ToCollectionSimple(t *testing.T) {
	safeTest(t, "Test_S13_47_LinkedCollections_ToCollectionSimple", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.ToCollectionSimple().Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_48_LinkedCollections_ToStrings(t *testing.T) {
	safeTest(t, "Test_S13_48_LinkedCollections_ToStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if len(lc.ToStrings()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_49_LinkedCollections_ToStringsPtr(t *testing.T) {
	safeTest(t, "Test_S13_49_LinkedCollections_ToStringsPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.ToStringsPtr() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S13_50_LinkedCollections_ToCollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_S13_50_LinkedCollections_ToCollectionsOfCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		coc := lc.ToCollectionsOfCollection(0)
		if coc == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S13_51_LinkedCollections_ToCollectionsOfCollection_Empty(t *testing.T) {
	safeTest(t, "Test_S13_51_LinkedCollections_ToCollectionsOfCollection_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		coc := lc.ToCollectionsOfCollection(0)
		if coc == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S13_52_LinkedCollections_ItemsOfItems(t *testing.T) {
	safeTest(t, "Test_S13_52_LinkedCollections_ItemsOfItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		items := lc.ItemsOfItems()
		if len(items) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_53_LinkedCollections_ItemsOfItems_Empty(t *testing.T) {
	safeTest(t, "Test_S13_53_LinkedCollections_ItemsOfItems_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		if len(lc.ItemsOfItems()) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_54_LinkedCollections_ItemsOfItemsCollection(t *testing.T) {
	safeTest(t, "Test_S13_54_LinkedCollections_ItemsOfItemsCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		items := lc.ItemsOfItemsCollection()
		if len(items) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_55_LinkedCollections_ItemsOfItemsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_S13_55_LinkedCollections_ItemsOfItemsCollection_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		if len(lc.ItemsOfItemsCollection()) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_56_LinkedCollections_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_S13_56_LinkedCollections_SimpleSlice", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		ss := lc.SimpleSlice()
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_57_LinkedCollections_List(t *testing.T) {
	safeTest(t, "Test_S13_57_LinkedCollections_List", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if len(lc.List()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_58_LinkedCollections_List_Empty(t *testing.T) {
	safeTest(t, "Test_S13_58_LinkedCollections_List_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		if len(lc.List()) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_59_LinkedCollections_ListPtr(t *testing.T) {
	safeTest(t, "Test_S13_59_LinkedCollections_ListPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.ListPtr() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S13_60_LinkedCollections_String(t *testing.T) {
	safeTest(t, "Test_S13_60_LinkedCollections_String", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.String() == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S13_61_LinkedCollections_String_Empty(t *testing.T) {
	safeTest(t, "Test_S13_61_LinkedCollections_String_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		if !strings.Contains(lc.String(), "No Element") {
			t.Fatal("expected No Element")
		}
	})
}

func Test_S13_62_LinkedCollections_StringLock(t *testing.T) {
	safeTest(t, "Test_S13_62_LinkedCollections_StringLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.StringLock() == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S13_63_LinkedCollections_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_S13_63_LinkedCollections_StringLock_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		if !strings.Contains(lc.StringLock(), "No Element") {
			t.Fatal("expected No Element")
		}
	})
}

func Test_S13_64_LinkedCollections_Join(t *testing.T) {
	safeTest(t, "Test_S13_64_LinkedCollections_Join", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.Join(",") == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S13_65_LinkedCollections_Joins(t *testing.T) {
	safeTest(t, "Test_S13_65_LinkedCollections_Joins", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		result := lc.Joins(",", "b")
		if !strings.Contains(result, "a") {
			t.Fatal("expected a")
		}
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
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		data, err := lc.MarshalJSON()
		if err != nil || len(data) == 0 {
			t.Fatal("expected valid JSON")
		}
	})
}

func Test_S13_68_LinkedCollections_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_S13_68_LinkedCollections_UnmarshalJSON", func() {
		lc := corestr.New.LinkedCollection.Create()
		err := lc.UnmarshalJSON([]byte(`["a","b"]`))
		if err != nil || lc.Length() != 1 {
			t.Fatalf("expected 1, got %d", lc.Length())
		}
	})
}

func Test_S13_69_LinkedCollections_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_S13_69_LinkedCollections_UnmarshalJSON_Invalid", func() {
		lc := corestr.New.LinkedCollection.Create()
		err := lc.UnmarshalJSON([]byte(`invalid`))
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_S13_70_LinkedCollections_JsonModel(t *testing.T) {
	safeTest(t, "Test_S13_70_LinkedCollections_JsonModel", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if len(lc.JsonModel()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_71_LinkedCollections_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_S13_71_LinkedCollections_JsonModelAny", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.JsonModelAny() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S13_72_LinkedCollections_Clear_RemoveAll(t *testing.T) {
	safeTest(t, "Test_S13_72_LinkedCollections_Clear_RemoveAll", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Clear()
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_73_LinkedCollections_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_S13_73_LinkedCollections_Clear_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Clear()
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_74_LinkedCollections_RemoveAll(t *testing.T) {
	safeTest(t, "Test_S13_74_LinkedCollections_RemoveAll", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.RemoveAll()
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_75_LinkedCollections_Json(t *testing.T) {
	safeTest(t, "Test_S13_75_LinkedCollections_Json", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		jsonResult := lc.Json()
		if jsonResult.HasError() {
			t.Fatal("expected no error")
		}
	})
}

func Test_S13_76_LinkedCollections_JsonPtr(t *testing.T) {
	safeTest(t, "Test_S13_76_LinkedCollections_JsonPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.JsonPtr() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S13_77_LinkedCollections_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_S13_77_LinkedCollections_ParseInjectUsingJson", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		jsonResult := lc.JsonPtr()
		target := corestr.New.LinkedCollection.Create()
		result, err := target.ParseInjectUsingJson(jsonResult)
		if err != nil || result.Length() < 1 {
			t.Fatal("expected at least 1")
		}
	})
}

func Test_S13_78_LinkedCollections_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_S13_78_LinkedCollections_ParseInjectUsingJsonMust", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		jsonResult := lc.JsonPtr()
		target := corestr.New.LinkedCollection.Create()
		result := target.ParseInjectUsingJsonMust(jsonResult)
		if result.Length() < 1 {
			t.Fatal("expected at least 1")
		}
	})
}

func Test_S13_79_LinkedCollections_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_S13_79_LinkedCollections_JsonParseSelfInject", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		jsonResult := lc.JsonPtr()
		target := corestr.New.LinkedCollection.Create()
		err := target.JsonParseSelfInject(jsonResult)
		if err != nil {
			t.Fatal("expected no error")
		}
	})
}

func Test_S13_80_LinkedCollections_AsJsoner(t *testing.T) {
	safeTest(t, "Test_S13_80_LinkedCollections_AsJsoner", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.AsJsoner() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S13_81_LinkedCollections_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_S13_81_LinkedCollections_AsJsonContractsBinder", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.AsJsonContractsBinder() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S13_82_LinkedCollections_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_S13_82_LinkedCollections_AsJsonParseSelfInjector", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.AsJsonParseSelfInjector() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S13_83_LinkedCollections_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_S13_83_LinkedCollections_AsJsonMarshaller", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.AsJsonMarshaller() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S13_84_LinkedCollections_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_S13_84_LinkedCollections_GetCompareSummary", func() {
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"x"}))
		b := corestr.New.LinkedCollection.Create()
		b.Add(corestr.New.Collection.Strings([]string{"y"}))
		summary := a.GetCompareSummary(b, "left", "right")
		if summary == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S13_85_LinkedCollections_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_S13_85_LinkedCollections_GetNextNodes", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		nodes := lc.GetNextNodes(2)
		if len(nodes) != 2 {
			t.Fatalf("expected 2, got %d", len(nodes))
		}
	})
}

func Test_S13_86_LinkedCollections_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_S13_86_LinkedCollections_GetAllLinkedNodes", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if len(lc.GetAllLinkedNodes()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_87_LinkedCollections_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_S13_87_LinkedCollections_SafeIndexAt", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		if lc.SafeIndexAt(1) == nil {
			t.Fatal("expected non-nil")
		}
		if lc.SafeIndexAt(-1) != nil {
			t.Fatal("expected nil")
		}
		if lc.SafeIndexAt(10) != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_S13_88_LinkedCollections_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_S13_88_LinkedCollections_SafePointerIndexAt", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if lc.SafePointerIndexAt(0) == nil {
			t.Fatal("expected non-nil")
		}
		if lc.SafePointerIndexAt(10) != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_S13_89_LinkedCollections_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_S13_89_LinkedCollections_RemoveNodeByIndex", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc.RemoveNodeByIndex(1)
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S13_90_LinkedCollections_RemoveNode(t *testing.T) {
	safeTest(t, "Test_S13_90_LinkedCollections_RemoveNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		node := lc.Head()
		lc.RemoveNode(node)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_91_LinkedCollections_AddAsync(t *testing.T) {
	safeTest(t, "Test_S13_91_LinkedCollections_AddAsync", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsync(corestr.New.Collection.Strings([]string{"a"}), wg)
		wg.Wait()
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_92_LinkedCollections_AddStringsAsync(t *testing.T) {
	safeTest(t, "Test_S13_92_LinkedCollections_AddStringsAsync", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddStringsAsync(wg, []string{"a"})
		wg.Wait()
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_93_LinkedCollections_AddStringsAsync_Nil(t *testing.T) {
	safeTest(t, "Test_S13_93_LinkedCollections_AddStringsAsync_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsAsync(nil, nil)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_94_LinkedCollections_AddCollectionsPtr(t *testing.T) {
	safeTest(t, "Test_S13_94_LinkedCollections_AddCollectionsPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		cols := []*corestr.Collection{corestr.New.Collection.Strings([]string{"a"})}
		lc.AddCollectionsPtr(cols)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_95_LinkedCollections_AddCollectionsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_S13_95_LinkedCollections_AddCollectionsPtr_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollectionsPtr(nil)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_96_LinkedCollections_AddCollections(t *testing.T) {
	safeTest(t, "Test_S13_96_LinkedCollections_AddCollections", func() {
		lc := corestr.New.LinkedCollection.Create()
		cols := []*corestr.Collection{nil, corestr.New.Collection.Strings([]string{"a"})}
		lc.AddCollections(cols)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_97_LinkedCollections_AddCollections_Empty(t *testing.T) {
	safeTest(t, "Test_S13_97_LinkedCollections_AddCollections_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollections(nil)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_98_LinkedCollections_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_S13_98_LinkedCollections_AppendChainOfNodes", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		other := corestr.New.LinkedCollection.Create()
		other.Add(corestr.New.Collection.Strings([]string{"b"}))
		other.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc.AppendChainOfNodes(other.Head())
		if lc.Length() < 3 {
			t.Fatal("expected at least 3")
		}
	})
}

func Test_S13_99_LinkedCollections_AppendChainOfNodes_OnEmpty(t *testing.T) {
	safeTest(t, "Test_S13_99_LinkedCollections_AppendChainOfNodes_OnEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		other := corestr.New.LinkedCollection.Create()
		other.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.AppendChainOfNodes(other.Head())
		if lc.Length() < 1 {
			t.Fatal("expected at least 1")
		}
	})
}

func Test_S13_100_LinkedCollections_InsertAt(t *testing.T) {
	safeTest(t, "Test_S13_100_LinkedCollections_InsertAt", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc.InsertAt(1, corestr.New.Collection.Strings([]string{"b"}))
		if lc.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_S13_101_LinkedCollections_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_S13_101_LinkedCollections_InsertAt_Front", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.InsertAt(0, corestr.New.Collection.Strings([]string{"a"}))
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S13_102_LinkedCollections_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_S13_102_LinkedCollections_RemoveNodeByIndexes", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc.RemoveNodeByIndexes(false, 1)
		if lc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S13_103_LinkedCollections_RemoveNodeByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_S13_103_LinkedCollections_RemoveNodeByIndexes_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.RemoveNodeByIndexes(false)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_104_LinkedCollections_AddAsyncFuncItems(t *testing.T) {
	safeTest(t, "Test_S13_104_LinkedCollections_AddAsyncFuncItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItems(wg, false, func() []string { return []string{"a"} })
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_105_LinkedCollections_AddAsyncFuncItems_Nil(t *testing.T) {
	safeTest(t, "Test_S13_105_LinkedCollections_AddAsyncFuncItems_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAsyncFuncItems(nil, false, nil...)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_106_LinkedCollections_AddAsyncFuncItemsPointer(t *testing.T) {
	safeTest(t, "Test_S13_106_LinkedCollections_AddAsyncFuncItemsPointer", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItemsPointer(wg, false, func() []string { return []string{"a"} })
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S13_107_LinkedCollections_AddAsyncFuncItemsPointer_Nil(t *testing.T) {
	safeTest(t, "Test_S13_107_LinkedCollections_AddAsyncFuncItemsPointer_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAsyncFuncItemsPointer(nil, false, nil...)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_108_LinkedCollections_AttachWithNode(t *testing.T) {
	safeTest(t, "Test_S13_108_LinkedCollections_AttachWithNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		node := lc.Head()
		addNode := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		err := lc.AttachWithNode(node, addNode)
		if err != nil {
			t.Fatal("expected no error")
		}
	})
}

func Test_S13_109_LinkedCollections_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_S13_109_LinkedCollections_AttachWithNode_NilCurrent", func() {
		lc := corestr.New.LinkedCollection.Create()
		addNode := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		err := lc.AttachWithNode(nil, addNode)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_S13_110_LinkedCollections_AddCollectionsToNode(t *testing.T) {
	safeTest(t, "Test_S13_110_LinkedCollections_AddCollectionsToNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.AddCollectionsToNode(true, lc.Head(), corestr.New.Collection.Strings([]string{"b"}))
		if lc.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_S13_111_LinkedCollections_AddCollectionsToNode_Nil(t *testing.T) {
	safeTest(t, "Test_S13_111_LinkedCollections_AddCollectionsToNode_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollectionsToNode(true, nil, nil...)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S13_112_LinkedCollections_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_S13_112_LinkedCollections_AddCollectionToNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		col := corestr.New.Collection.Strings([]string{"b"})
		lc.AddCollectionToNode(true, lc.Head(), col)
		if lc.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}
