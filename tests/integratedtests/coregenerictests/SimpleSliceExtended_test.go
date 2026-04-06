package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_SimpleSlice_AddIf_Verification(t *testing.T) {
	for caseIndex, testCase := range simpleSliceAddIfTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isAdd := input.GetAsBoolDefault("isAdd", false)
		item, _ := input.GetAsString("item")
		s := coregeneric.EmptySimpleSlice[string]()

		// Act
		s.AddIf(isAdd, item)

		actual := args.Map{
			"length": s.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_SimpleSlice_AddsIf_Verification(t *testing.T) {
	for caseIndex, testCase := range simpleSliceAddsIfTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isAdd := input.GetAsBoolDefault("isAdd", false)
		s := coregeneric.EmptySimpleSlice[string]()

		// Act
		s.AddsIf(isAdd, "a", "b")

		actual := args.Map{
			"length": s.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_SimpleSlice_Methods_Verification(t *testing.T) {
	for caseIndex, testCase := range simpleSliceMethodsTestCases {
		// Arrange
		s := coregeneric.EmptySimpleSlice[string]()
		s.Add("a").Add("b").Add("c")

		// Act
		forEachCount := 0
		s.ForEach(func(index int, item string) {
			forEachCount++
		})

		countFunc := s.CountFunc(func(index int, item string) bool {
			return item != ""
		})

		actual := args.Map{
			"firstOrDefault": s.FirstOrDefault(),
			"lastOrDefault":  s.LastOrDefault(),
			"count":          s.Count(),
			"hasAnyItem":     s.HasAnyItem(),
			"hasItems":       s.HasItems(),
			"hasIndex1":      s.HasIndex(1),
			"hasIndex5":      s.HasIndex(5),
			"itemsLen":       len(s.Items()),
			"forEachCount":   forEachCount,
			"countFuncGt0":   countFunc > 0,
			"stringNotEmpty": s.String() != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_SimpleSlice_Empty_Methods_Verification(t *testing.T) {
	for caseIndex, testCase := range simpleSliceEmptyMethodsTestCases {
		// Arrange
		s := coregeneric.EmptySimpleSlice[string]()

		// Act
		actual := args.Map{
			"firstOrDefault": s.FirstOrDefault(),
			"lastOrDefault":  s.LastOrDefault(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_SimpleSlice_AddSlice_Verification(t *testing.T) {
	for caseIndex, testCase := range simpleSliceAddSliceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rawItems, _ := input.Get("items")
		items := rawItems.([]string)
		s := coregeneric.EmptySimpleSlice[string]()

		// Act
		s.AddSlice(items)

		actual := args.Map{
			"length": s.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_SimpleSlice_AddFunc_Verification(t *testing.T) {
	for caseIndex, testCase := range simpleSliceAddFuncTestCases {
		// Arrange
		s := coregeneric.EmptySimpleSlice[string]()

		// Act
		s.AddFunc(func() string { return "gen" })

		actual := args.Map{
			"length": s.Length(),
			"first":  s.First(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_SimpleSlice_InsertAt_Verification(t *testing.T) {
	for caseIndex, testCase := range simpleSliceInsertAtTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		index := input.GetAsIntDefault("index", 0)
		item, _ := input.GetAsString("item")
		s := coregeneric.EmptySimpleSlice[string]()
		s.Add("a").Add("b").Add("c")

		// Act
		s.InsertAt(index, item)

		actual := args.Map{
			"length": s.Length(),
		}

		expected := testCase.ExpectedInput.(args.Map)
		if _, has := expected["atIndex"]; has {
			actual["atIndex"] = s.Items()[index]
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LinkedListNode_Extended_Verification(t *testing.T) {
	for caseIndex, testCase := range linkedListNodeExtendedTestCases {
		// Arrange
		list := coregeneric.LinkedListFrom[string]([]string{"a", "b", "c"})
		head := list.Head()

		// Act
		endNode, length := head.EndOfChain()
		cloned := head.Clone()
		listPtr := head.ListPtr()
		str := endNode.String()

		actual := args.Map{
			"endLength":    length,
			"cloneHasNext": cloned.HasNext(),
			"listPtrLen":   len(*listPtr),
			"stringOk":     str != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
