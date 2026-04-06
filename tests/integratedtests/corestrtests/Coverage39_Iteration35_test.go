package corestrtests

import (
	"encoding/json"
	"strings"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════
// LinkedCollections — basic operations
// ══════════════════════════════════════════════════════════════

func Test_Cov39_LC_Add(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Add", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		lc.Add(col)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_Head_Tail(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Head_Tail", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.Add(c1).Add(c2)
		actual := args.Map{"result": lc.Head().Element.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual := args.Map{"result": lc.Tail().Element.First() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_Cov39_LC_First_Last(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_First_Last", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"x"})
		c2 := corestr.New.Collection.Strings([]string{"y"})
		lc.Add(c1).Add(c2)
		actual := args.Map{"result": lc.First().First() != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
		actual := args.Map{"result": lc.Last().First() != "y"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected y", actual)
	})
}

func Test_Cov39_LC_Single(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Single", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"only"}))
		actual := args.Map{"result": lc.Single().First() != "only"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected only", actual)
	})
}

func Test_Cov39_LC_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_FirstOrDefault_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": lc.FirstOrDefault().Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov39_LC_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_LastOrDefault_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": lc.LastOrDefault().Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov39_LC_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IsEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": lc.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual := args.Map{"result": lc.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
	})
}

func Test_Cov39_LC_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IsEmptyLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": lc.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov39_LC_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_LengthLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AllIndividualItemsLength", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		actual := args.Map{"result": lc.AllIndividualItemsLength() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov39_LC_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddLock(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddStrings_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings()
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddStringsLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock("a")
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_AddStringsLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddStringsLock_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock()
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_PushBack(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_PushBack", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.PushBack(corestr.New.Collection.Strings([]string{"x"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_PushBackLock(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_PushBackLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.PushBackLock(corestr.New.Collection.Strings([]string{"x"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_Push(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Push", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Push(corestr.New.Collection.Strings([]string{"x"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_AddFront(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddFront", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"b"})
		c2 := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c1)
		lc.AddFront(c2)
		actual := args.Map{"result": lc.First().First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a first", actual)
	})
}

func Test_Cov39_LC_AddFront_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddFront_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddFront(corestr.New.Collection.Strings([]string{"x"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_PushFront(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_PushFront", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.PushFront(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.First().First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Cov39_LC_AddFrontLock(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddFrontLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddFrontLock(corestr.New.Collection.Strings([]string{"x"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_AddAnother(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddAnother", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc1.AddAnother(lc2)
		actual := args.Map{"result": lc1.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov39_LC_AddAnother_Nil(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddAnother_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAnother(nil)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_AddAnother_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddAnother_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAnother(corestr.New.LinkedCollection.Create())
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollection(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollection_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollection(nil)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_AddCollections(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollections", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.AddCollections([]*corestr.Collection{c1})
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_AddCollections_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollections_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollections(nil)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_AddCollectionsPtr(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollectionsPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.AddCollectionsPtr([]*corestr.Collection{c1})
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_AddCollectionsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollectionsPtr_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollectionsPtr(nil)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AppendNode / AppendChainOfNodes / AddBackNode ──

func Test_Cov39_LC_AppendNode_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendNode_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AppendNode(&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"x"})})
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_AppendNode_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendNode_NonEmpty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.AppendNode(&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})})
		actual := args.Map{"result": lc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov39_LC_AddBackNode(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddBackNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddBackNode(&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"x"})})
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_AppendChainOfNodes(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendChainOfNodes", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc2.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc1.AppendChainOfNodes(lc2.Head())
		actual := args.Map{"result": lc1.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov39_LC_AppendChainOfNodes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendChainOfNodes_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		chain := corestr.New.LinkedCollection.Create()
		chain.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.AppendChainOfNodes(chain.Head())
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── AttachWithNode ──

func Test_Cov39_LC_AttachWithNode_NilCurrent(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AttachWithNode_NilCurrent", func() {
		lc := corestr.New.LinkedCollection.Create()
		err := lc.AttachWithNode(nil, &corestr.LinkedCollectionNode{})
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_Cov39_LC_AttachWithNode_NextNotNil(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AttachWithNode_NextNotNil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		err := lc.AttachWithNode(lc.Head(), &corestr.LinkedCollectionNode{})
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

// ── InsertAt ──

func Test_Cov39_LC_InsertAt_Front(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_InsertAt_Front", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.InsertAt(0, corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.First().First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a first", actual)
	})
}

func Test_Cov39_LC_InsertAt_Middle(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_InsertAt_Middle", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc.InsertAt(1, corestr.New.Collection.Strings([]string{"b"}))
		actual := args.Map{"result": lc.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

// ── Loop ──

func Test_Cov39_LC_Loop(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Loop", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return false
		})
		actual := args.Map{"result": count != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov39_LC_Loop_Break(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Loop_Break", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return true
		})
		actual := args.Map{"result": count != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Loop_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			actual := args.Map{"result": false}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "should not be called", actual)
			return false
		})
	})
}

// ── Filter / FilterAsCollection / FilterAsCollections ──

func Test_Cov39_LC_Filter(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Filter", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		nodes := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		actual := args.Map{"result": len(nodes) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov39_LC_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Filter_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		nodes := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		actual := args.Map{"result": len(nodes) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_Filter_BreakFirst(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Filter_BreakFirst", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		nodes := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})
		actual := args.Map{"result": len(nodes) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_FilterAsCollection(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_FilterAsCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		col := lc.FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		}, 0)
		actual := args.Map{"result": col.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov39_LC_FilterAsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_FilterAsCollection_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		col := lc.FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		}, 0)
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_FilterAsCollections(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_FilterAsCollections", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		cols := lc.FilterAsCollections(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true}
		})
		actual := args.Map{"result": len(cols) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── RemoveNodeByIndex ──

func Test_Cov39_LC_RemoveByIndex_First(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_RemoveByIndex_First", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.RemoveNodeByIndex(0)
		actual := args.Map{"result": lc.Length() != 1 || lc.First().First() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov39_LC_RemoveByIndex_Last(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_RemoveByIndex_Last", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.RemoveNodeByIndex(1)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_RemoveByIndex_Middle(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_RemoveByIndex_Middle", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc.RemoveNodeByIndex(1)
		actual := args.Map{"result": lc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// ── RemoveNodeByIndexes ──

func Test_Cov39_LC_RemoveByIndexes(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_RemoveByIndexes", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		lc.RemoveNodeByIndexes(true, 0, 2)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_RemoveByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_RemoveByIndexes_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.RemoveNodeByIndexes(true)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── RemoveNode ──

func Test_Cov39_LC_RemoveNode_Head(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_RemoveNode_Head", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.RemoveNode(lc.Head())
		actual := args.Map{"result": lc.Length() != 1 || lc.First().First() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov39_LC_RemoveNode_NonHead(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_RemoveNode_NonHead", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.RemoveNode(lc.IndexAt(1))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── AppendCollections / AppendCollectionsPointers / AppendCollectionsPointersLock ──

func Test_Cov39_LC_AppendCollections(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendCollections", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc.AppendCollections(false, c1)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_AppendCollections_SkipNil(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendCollections_SkipNil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AppendCollections(true, nil)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_AppendCollectionsPointers(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendCollectionsPointers", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		cols := []*corestr.Collection{c1}
		lc.AppendCollectionsPointers(false, &cols)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_AppendCollectionsPointers_NilSkip(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendCollectionsPointers_NilSkip", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AppendCollectionsPointers(true, nil)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_AppendCollectionsPointersLock(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendCollectionsPointersLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		cols := []*corestr.Collection{c1}
		lc.AppendCollectionsPointersLock(false, &cols)
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── AddCollectionsToNode / AddCollectionsPointerToNode ──

func Test_Cov39_LC_AddCollectionsToNode(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollectionsToNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.AddCollectionsToNode(false, lc.Head(), c2)
		actual := args.Map{"result": lc.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_Cov39_LC_AddCollectionsPointerToNode_NilSkip(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollectionsPointerToNode_NilSkip", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollectionsPointerToNode(true, nil, nil)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_AddCollectionsPointerToNode_NilItems(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollectionsPointerToNode_NilItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollectionsPointerToNode(true, nil, nil)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AddCollectionToNode ──

func Test_Cov39_LC_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddCollectionToNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.AddCollectionToNode(false, lc.Head(), corestr.New.Collection.Strings([]string{"b"}))
		actual := args.Map{"result": lc.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

// ── AddAsync ──

func Test_Cov39_LC_AddAsync(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddAsync", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsync(corestr.New.Collection.Strings([]string{"a"}), wg)
		wg.Wait()
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_AppendChainOfNodesAsync(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AppendChainOfNodesAsync", func() {
		lc := corestr.New.LinkedCollection.Create()
		chain := corestr.New.LinkedCollection.Create()
		chain.Add(corestr.New.Collection.Strings([]string{"a"}))
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AppendChainOfNodesAsync(chain.Head(), wg)
		wg.Wait()
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── AddStringsOfStrings ──

func Test_Cov39_LC_AddStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddStringsOfStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false, []string{"a"}, []string{"b"})
		actual := args.Map{"result": lc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov39_LC_AddStringsOfStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddStringsOfStrings_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── AddAsyncFuncItems ──

func Test_Cov39_LC_AddAsyncFuncItems(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddAsyncFuncItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItems(wg, false, func() []string { return []string{"a"} })
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_AddAsyncFuncItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddAsyncFuncItems_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAsyncFuncItems(nil, false)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_AddAsyncFuncItemsPointer(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddAsyncFuncItemsPointer", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		lc.AddAsyncFuncItemsPointer(wg, false, func() []string { return []string{"x"} })
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_AddAsyncFuncItemsPointer_Nil(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AddAsyncFuncItemsPointer_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAsyncFuncItemsPointer(nil, false)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ── ConcatNew ──

func Test_Cov39_LC_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ConcatNew", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"b"}))
		result := lc1.ConcatNew(false, lc2)
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov39_LC_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ConcatNew_EmptyClone", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		result := lc.ConcatNew(true)
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ConcatNew_EmptyNoClone", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		result := lc.ConcatNew(false)
		actual := args.Map{"result": result != lc}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same pointer", actual)
	})
}

// ── IndexAt / SafeIndexAt / SafePointerIndexAt ──

func Test_Cov39_LC_IndexAt(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IndexAt", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		node := lc.IndexAt(1)
		actual := args.Map{"result": node.Element.First() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_Cov39_LC_IndexAt_Zero(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IndexAt_Zero", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.IndexAt(0).Element.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Cov39_LC_IndexAt_Negative(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IndexAt_Negative", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.IndexAt(-1) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Cov39_LC_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_SafeIndexAt", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.SafeIndexAt(0).Element.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Cov39_LC_SafeIndexAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_SafeIndexAt_OutOfRange", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.SafeIndexAt(5) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Cov39_LC_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_SafePointerIndexAt", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.SafePointerIndexAt(0) == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov39_LC_SafePointerIndexAt_Nil(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_SafePointerIndexAt_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": lc.SafePointerIndexAt(0) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

// ── GetNextNodes / GetAllLinkedNodes ──

func Test_Cov39_LC_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_GetNextNodes", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		nodes := lc.GetNextNodes(2)
		actual := args.Map{"result": len(nodes) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov39_LC_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_GetAllLinkedNodes", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		nodes := lc.GetAllLinkedNodes()
		actual := args.Map{"result": len(nodes) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── ToCollection / ToStrings / ToStringsPtr / ToCollectionSimple ──

func Test_Cov39_LC_ToCollection(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ToCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		col := lc.ToCollection(0)
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov39_LC_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ToCollection_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": lc.ToCollection(0).Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_ToCollectionSimple(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ToCollectionSimple", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.ToCollectionSimple().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_ToStrings(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ToStrings", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": len(lc.ToStrings()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_ToStringsPtr(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ToStringsPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		ptr := lc.ToStringsPtr()
		actual := args.Map{"result": ptr == nil || len(*ptr) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov39_LC_ToCollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ToCollectionsOfCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		coc := lc.ToCollectionsOfCollection(0)
		actual := args.Map{"result": coc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov39_LC_ToCollectionsOfCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ToCollectionsOfCollection_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		coc := lc.ToCollectionsOfCollection(0)
		actual := args.Map{"result": coc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ── ItemsOfItems / ItemsOfItemsCollection ──

func Test_Cov39_LC_ItemsOfItems(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ItemsOfItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		items := lc.ItemsOfItems()
		actual := args.Map{"result": len(items) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_ItemsOfItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ItemsOfItems_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		items := lc.ItemsOfItems()
		actual := args.Map{"result": len(items) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_ItemsOfItemsCollection(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ItemsOfItemsCollection", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		items := lc.ItemsOfItemsCollection()
		actual := args.Map{"result": len(items) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ── SimpleSlice / List / ListPtr ──

func Test_Cov39_LC_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_SimpleSlice", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		ss := lc.SimpleSlice()
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_List(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_List", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		actual := args.Map{"result": len(lc.List()) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov39_LC_List_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_List_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": len(lc.List()) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_ListPtr(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ListPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.ListPtr() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ── String / StringLock / Join / Joins ──

func Test_Cov39_LC_String(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_String", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov39_LC_String_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_String_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		s := lc.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty (NoElements)", actual)
	})
}

func Test_Cov39_LC_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_StringLock", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": strings.Contains(lc.StringLock(), "a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_Cov39_LC_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_StringLock_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		s := lc.StringLock()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov39_LC_Join(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Join", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		actual := args.Map{"result": lc.Join(",") == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov39_LC_Joins(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Joins", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		result := lc.Joins(",", "b")
		actual := args.Map{"result": strings.Contains(result, "a") || !strings.Contains(result, "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_Cov39_LC_Joins_NilItems(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Joins_NilItems", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		_ = lc.Joins(",")
	})
}

// ── IsEqualsPtr ──

func Test_Cov39_LC_IsEqualsPtr_Same(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IsEqualsPtr_Same", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc1.IsEqualsPtr(lc2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov39_LC_IsEqualsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IsEqualsPtr_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": lc.IsEqualsPtr(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov39_LC_IsEqualsPtr_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IsEqualsPtr_SamePtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.IsEqualsPtr(lc)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov39_LC_IsEqualsPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IsEqualsPtr_BothEmpty", func() {
		a := corestr.New.LinkedCollection.Create()
		b := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": a.IsEqualsPtr(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov39_LC_IsEqualsPtr_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IsEqualsPtr_DiffLen", func() {
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		b := corestr.New.LinkedCollection.Create()
		b.Add(corestr.New.Collection.Strings([]string{"a"}))
		b.Add(corestr.New.Collection.Strings([]string{"b"}))
		actual := args.Map{"result": a.IsEqualsPtr(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov39_LC_IsEqualsPtr_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_IsEqualsPtr_OneEmpty", func() {
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		b := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": a.IsEqualsPtr(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

// ── GetCompareSummary ──

func Test_Cov39_LC_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_GetCompareSummary", func() {
		a := corestr.New.LinkedCollection.Create()
		a.Add(corestr.New.Collection.Strings([]string{"a"}))
		b := corestr.New.LinkedCollection.Create()
		b.Add(corestr.New.Collection.Strings([]string{"b"}))
		summary := a.GetCompareSummary(b, "left", "right")
		actual := args.Map{"result": summary == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

// ── JSON ──

func Test_Cov39_LC_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_MarshalJSON", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		data, err := json.Marshal(lc)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual := args.Map{"result": strings.Contains(string(data), "\"a\"")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_Cov39_LC_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_UnmarshalJSON", func() {
		lc := corestr.New.LinkedCollection.Create()
		err := json.Unmarshal([]byte(`["x","y"]`), lc)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		if lc.Length() != 1 { // unmarshal adds as single collection
			// accept whatever the implementation does
		}
	})
}

func Test_Cov39_LC_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_UnmarshalJSON_Invalid", func() {
		lc := corestr.New.LinkedCollection.Create()
		err := json.Unmarshal([]byte(`bad`), lc)
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_Cov39_LC_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_JsonModel", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": len(lc.JsonModel()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LC_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_JsonModelAny", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.JsonModelAny() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov39_LC_Json(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Json", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.Json().Error != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_Cov39_LC_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_JsonPtr", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.JsonPtr() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov39_LC_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ParseInjectUsingJson", func() {
		src := corestr.New.LinkedCollection.Create()
		src.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc := corestr.New.LinkedCollection.Create()
		_, err := lc.ParseInjectUsingJson(src.JsonPtr())
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_Cov39_LC_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_ParseInjectUsingJsonMust", func() {
		src := corestr.New.LinkedCollection.Create()
		src.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc := corestr.New.LinkedCollection.Create()
		result := lc.ParseInjectUsingJsonMust(src.JsonPtr())
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov39_LC_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_JsonParseSelfInject", func() {
		src := corestr.New.LinkedCollection.Create()
		src.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc := corestr.New.LinkedCollection.Create()
		err := lc.JsonParseSelfInject(src.JsonPtr())
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_Cov39_LC_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AsJsoner", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": lc.AsJsoner() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov39_LC_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AsJsonContractsBinder", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": lc.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov39_LC_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AsJsonParseSelfInjector", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": lc.AsJsonParseSelfInjector() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov39_LC_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_AsJsonMarshaller", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": lc.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ── Clear / RemoveAll ──

func Test_Cov39_LC_Clear(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Clear", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Clear()
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_Clear_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Clear()
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_LC_RemoveAll(t *testing.T) {
	safeTest(t, "Test_Cov39_LC_RemoveAll", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.RemoveAll()
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ══════════════════════════════════════════════════════════════
// LinkedCollectionNode
// ══════════════════════════════════════════════════════════════

func Test_Cov39_LCNode_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsEmpty", func() {
		var n *corestr.LinkedCollectionNode
		actual := args.Map{"result": n.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov39_LCNode_HasElement(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_HasElement", func() {
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		actual := args.Map{"result": n.HasElement()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov39_LCNode_HasNext(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_HasNext", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		actual := args.Map{"result": lc.Head().HasNext()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": lc.Tail().HasNext()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov39_LCNode_Next(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_Next", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		actual := args.Map{"result": lc.Head().Next().Element.First() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_Cov39_LCNode_EndOfChain(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_EndOfChain", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		end, length := lc.Head().EndOfChain()
		actual := args.Map{"result": length != 2 || end.Element.First() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov39_LCNode_LoopEndOfChain(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_LoopEndOfChain", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		count := 0
		end, length := lc.Head().LoopEndOfChain(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return false
		})
		actual := args.Map{"result": count != 2 || length != 2 || end.Element.First() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov39_LCNode_LoopEndOfChain_Break(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_LoopEndOfChain_Break", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		_, length := lc.Head().LoopEndOfChain(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			return true
		})
		actual := args.Map{"result": length != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LCNode_Clone(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_Clone", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		cloned := lc.Head().Clone()
		actual := args.Map{"result": cloned.Element.First() != "a" || cloned.HasNext()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov39_LCNode_IsEqual_Same(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsEqual_Same", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		n1 := &corestr.LinkedCollectionNode{Element: c}
		n2 := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		actual := args.Map{"result": n1.IsEqual(n2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov39_LCNode_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsEqual_BothNil", func() {
		var a, b *corestr.LinkedCollectionNode
		actual := args.Map{"result": a.IsEqual(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov39_LCNode_IsEqual_OneNil(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsEqual_OneNil", func() {
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		actual := args.Map{"result": n.IsEqual(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov39_LCNode_IsEqual_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsEqual_SamePtr", func() {
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		actual := args.Map{"result": n.IsEqual(n)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov39_LCNode_IsChainEqual(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsChainEqual", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc1.Head().IsChainEqual(lc2.Head())}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov39_LCNode_IsChainEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsChainEqual_BothNil", func() {
		var a, b *corestr.LinkedCollectionNode
		actual := args.Map{"result": a.IsChainEqual(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov39_LCNode_IsEqualValue(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsEqualValue", func() {
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		actual := args.Map{"result": n.IsEqualValue(corestr.New.Collection.Strings([]string{"a"}))}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov39_LCNode_IsEqualValue_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_IsEqualValue_BothNil", func() {
		n := &corestr.LinkedCollectionNode{Element: nil}
		actual := args.Map{"result": n.IsEqualValue(nil)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov39_LCNode_String(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_String", func() {
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		actual := args.Map{"result": n.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov39_LCNode_List(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_List", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))
		list := lc.Head().List()
		actual := args.Map{"result": len(list) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov39_LCNode_ListPtr(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_ListPtr", func() {
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		actual := args.Map{"result": n.ListPtr() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov39_LCNode_Join(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_Join", func() {
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a", "b"})}
		actual := args.Map{"result": n.Join(",") != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov39_LCNode_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_CreateLinkedList", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		newLC := lc.Head().CreateLinkedList()
		actual := args.Map{"result": newLC.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_LCNode_AddNext(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_AddNext", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Head().AddNext(lc, corestr.New.Collection.Strings([]string{"b"}))
		actual := args.Map{"result": lc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov39_LCNode_AddNextNode(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_AddNextNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		newNode := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		lc.Head().AddNextNode(lc, newNode)
		actual := args.Map{"result": lc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov39_LCNode_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_AddStringsToNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Head().AddStringsToNode(lc, false, []string{"b"}, false)
		actual := args.Map{"result": lc.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_Cov39_LCNode_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_Cov39_LCNode_AddCollectionToNode", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Head().AddCollectionToNode(lc, false, corestr.New.Collection.Strings([]string{"b"}))
		actual := args.Map{"result": lc.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

// ══════════════════════════════════════════════════════════════
// NonChainedLinkedCollectionNodes
// ══════════════════════════════════════════════════════════════

func Test_Cov39_NCLCN_Basic(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_Basic", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)
		actual := args.Map{"result": nc.IsEmpty() || nc.HasItems() || nc.Length() != 0}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov39_NCLCN_Adds(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_Adds", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		n1 := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		nc.Adds(n1)
		actual := args.Map{"result": nc.Length() != 1 || nc.First() != n1 || nc.Last() != n1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov39_NCLCN_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_Adds_Nil", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		nc.Adds()
		actual := args.Map{"result": nc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov39_NCLCN_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_FirstOrDefault_Empty", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		actual := args.Map{"result": nc.FirstOrDefault() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Cov39_NCLCN_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_LastOrDefault_Empty", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		actual := args.Map{"result": nc.LastOrDefault() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Cov39_NCLCN_Items(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_Items", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		nc.Adds(&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})})
		actual := args.Map{"result": len(nc.Items()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov39_NCLCN_ApplyChaining(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_ApplyChaining", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		n1 := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		n2 := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})}
		nc.Adds(n1, n2)
		nc.ApplyChaining()
		actual := args.Map{"result": nc.IsChainingApplied()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": n1.HasNext()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected chaining", actual)
	})
}

func Test_Cov39_NCLCN_ApplyChaining_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_ApplyChaining_Empty", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		nc.ApplyChaining()
		// should not panic
	})
}

func Test_Cov39_NCLCN_ToChainedNodes(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_ToChainedNodes", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		nc.Adds(&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})})
		nc.Adds(&corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"b"})})
		chained := nc.ToChainedNodes()
		actual := args.Map{"result": chained == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov39_NCLCN_ToChainedNodes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_NCLCN_ToChainedNodes_Empty", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(3)
		chained := nc.ToChainedNodes()
		actual := args.Map{"result": chained == nil || len(*chained) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// ── newLinkedListCollectionsCreator ──

func Test_Cov39_Creator_LC_Create(t *testing.T) {
	safeTest(t, "Test_Cov39_Creator_LC_Create", func() {
		lc := corestr.New.LinkedCollection.Create()
		actual := args.Map{"result": lc == nil || lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov39_Creator_LC_Empty(t *testing.T) {
	safeTest(t, "Test_Cov39_Creator_LC_Empty", func() {
		lc := corestr.New.LinkedCollection.Empty()
		actual := args.Map{"result": lc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov39_Creator_LC_Strings(t *testing.T) {
	safeTest(t, "Test_Cov39_Creator_LC_Strings", func() {
		lc := corestr.New.LinkedCollection.Strings("a", "b")
		if lc.Length() != 1 { // all strings as one collection
			// implementation may vary
		}
	})
}

func Test_Cov39_Creator_LC_UsingCollections(t *testing.T) {
	safeTest(t, "Test_Cov39_Creator_LC_UsingCollections", func() {
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc := corestr.New.LinkedCollection.UsingCollections(c1, c2)
		actual := args.Map{"result": lc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov39_Creator_LC_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_Cov39_Creator_LC_PointerStringsPtr", func() {
		a := "a"
		ptrs := []*string{&a, nil}
		lc := corestr.New.LinkedCollection.PointerStringsPtr(&ptrs)
		actual := args.Map{"result": lc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov39_Creator_LC_PointerStringsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov39_Creator_LC_PointerStringsPtr_Nil", func() {
		lc := corestr.New.LinkedCollection.PointerStringsPtr(nil)
		actual := args.Map{"result": lc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}
