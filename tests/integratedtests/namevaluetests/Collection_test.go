package namevaluetests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/namevalue"
)

// ==========================================================================
// Test: StringStringCollection
// ==========================================================================

func Test_StringStringCollection_Verification(t *testing.T) {
	for caseIndex, testCase := range stringStringCollectionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.Get("count")
		countInt := count.(int)

		// Act
		col := namevalue.NewGenericCollectionDefault[string, string]()
		for i := 0; i < countInt; i++ {
			col.Add(namevalue.StringString{
				Name:  fmt.Sprintf("key%d", i),
				Value: fmt.Sprintf("val%d", i),
			})
		}
		actual := args.Map{
			"length":   col.Length(),
			"isEmpty":  col.IsEmpty(),
			"hasItems": col.HasAnyItem(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// Test: StringIntCollection
// ==========================================================================

func Test_StringIntCollection_Verification(t *testing.T) {
	for caseIndex, testCase := range stringIntCollectionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.Get("count")
		countInt := count.(int)

		// Act
		col := namevalue.NewGenericCollectionDefault[string, int]()
		for i := 0; i < countInt; i++ {
			col.Add(namevalue.StringInt{
				Name:  fmt.Sprintf("item%d", i),
				Value: i * 10,
			})
		}
		joined := col.Join(", ")
		actual := args.Map{
			"length":          col.Length(),
			"hasFirstItem":    strings.Contains(joined, "item0"),
			"joinContainsAll": strings.Contains(joined, ","),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// Test: Prepend
// ==========================================================================

func Test_Collection_Prepend(t *testing.T) {
	tc := collectionPrependTestCase

	col := namevalue.NewGenericCollectionDefault[string, string]()
	col.Add(namevalue.StringString{Name: "original-0", Value: "v0"})
	col.Add(namevalue.StringString{Name: "original-1", Value: "v1"})

	// Act
	col.Prepend(namevalue.StringString{Name: "prepended", Value: "vp"})

	actual := args.Map{
		"length":    col.Length(),
		"firstItem": col.Items[0].Name,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Append
// ==========================================================================

func Test_Collection_Append(t *testing.T) {
	tc := collectionAppendTestCase

	col := namevalue.NewGenericCollectionDefault[string, string]()
	col.Add(namevalue.StringString{Name: "original-0", Value: "v0"})
	col.Add(namevalue.StringString{Name: "original-1", Value: "v1"})

	// Act
	col.Append(namevalue.StringString{Name: "appended", Value: "va"})

	actual := args.Map{
		"length":   col.Length(),
		"lastItem": col.Items[col.LastIndex()].Name,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: PrependIf false
// ==========================================================================

func Test_Collection_PrependIfFalse(t *testing.T) {
	tc := collectionPrependIfFalseTestCase

	col := namevalue.NewGenericCollectionDefault[string, string]()
	col.Add(namevalue.StringString{Name: "original-0", Value: "v0"})
	col.Add(namevalue.StringString{Name: "original-1", Value: "v1"})

	// Act
	col.PrependIf(false, namevalue.StringString{Name: "skipped", Value: "vs"})

	actual := args.Map{
		"length":    col.Length(),
		"firstItem": col.Items[0].Name,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: AppendIf false
// ==========================================================================

func Test_Collection_AppendIfFalse(t *testing.T) {
	tc := collectionAppendIfFalseTestCase

	col := namevalue.NewGenericCollectionDefault[string, string]()
	col.Add(namevalue.StringString{Name: "original-0", Value: "v0"})
	col.Add(namevalue.StringString{Name: "original-1", Value: "v1"})

	// Act
	col.AppendIf(false, namevalue.StringString{Name: "skipped", Value: "vs"})

	actual := args.Map{
		"length":    col.Length(),
		"firstItem": col.Items[0].Name,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Clone — valid collection
// ==========================================================================

func Test_CollectionClone_Valid(t *testing.T) {
	tc := collectionCloneValidTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	count, _ := input.Get("count")
	countInt := count.(int)

	// Act
	col := namevalue.NewGenericCollectionDefault[string, string]()
	for i := 0; i < countInt; i++ {
		col.Add(namevalue.StringString{
			Name:  fmt.Sprintf("k%d", i),
			Value: fmt.Sprintf("v%d", i),
		})
	}

	cloned := col.Clone()
	col.Add(namevalue.StringString{Name: "extra", Value: "x"})

	actual := args.Map{
		"length":        cloned.Length(),
		"sameContent":   cloned.IsEqualByString(namevalue.NewGenericCollectionUsing[string, string](true, cloned.Items...)),
		"isIndependent": cloned.Length() == countInt,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Clone — nil receiver (CaseNilSafe pattern)
// ==========================================================================

func Test_CollectionClone_NilReceiver(t *testing.T) {
	for caseIndex, tc := range collectionNilSafeTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================================================
// Test: IsEqualByString — Equal
// ==========================================================================

func Test_CollectionIsEqual_Equal(t *testing.T) {
	tc := collectionIsEqualEqualTestCase

	a := namevalue.NewGenericCollectionDefault[string, int]()
	a.Add(namevalue.StringInt{Name: "x", Value: 1})
	b := namevalue.NewGenericCollectionDefault[string, int]()
	b.Add(namevalue.StringInt{Name: "x", Value: 1})

	// Assert
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEqualByString(b)))
}

// ==========================================================================
// Test: IsEqualByString — NotEqual
// ==========================================================================

func Test_CollectionIsEqual_NotEqual(t *testing.T) {
	tc := collectionIsEqualNotEqualTestCase

	a := namevalue.NewGenericCollectionDefault[string, int]()
	a.Add(namevalue.StringInt{Name: "x", Value: 1})
	b := namevalue.NewGenericCollectionDefault[string, int]()
	b.Add(namevalue.StringInt{Name: "x", Value: 99})

	// Assert
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEqualByString(b)))
}

// ==========================================================================
// Test: IsEqualByString — DiffLength
// ==========================================================================

func Test_CollectionIsEqual_DiffLength(t *testing.T) {
	tc := collectionIsEqualDiffLengthTestCase

	a := namevalue.NewGenericCollectionDefault[string, int]()
	a.Add(namevalue.StringInt{Name: "x", Value: 1})
	b := namevalue.NewGenericCollectionDefault[string, int]()
	b.Add(namevalue.StringInt{Name: "x", Value: 1})
	b.Add(namevalue.StringInt{Name: "y", Value: 2})

	// Assert
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEqualByString(b)))
}

// ==========================================================================
// Test: IsEqualByString — BothNils
// ==========================================================================

func Test_CollectionIsEqual_BothNils(t *testing.T) {
	tc := collectionIsEqualBothNilsTestCase

	var a *namevalue.StringIntCollection
	var b *namevalue.StringIntCollection

	// Assert
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", a.IsEqualByString(b)))
}

// ==========================================================================
// Test: Error
// ==========================================================================

func Test_CollectionError_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.Get("count")
		countInt := count.(int)

		col := namevalue.NewGenericCollectionDefault[string, string]()
		for i := 0; i < countInt; i++ {
			col.Add(namevalue.StringString{
				Name:  fmt.Sprintf("err%d", i),
				Value: fmt.Sprintf("msg%d", i),
			})
		}

		err := col.Error()
		errMsg := col.ErrorUsingMessage("failed:")

		// Act
		actual := args.Map{
			"hasError":           err != nil,
			"errorContainsItems": errMsg != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// Test: Dispose
// ==========================================================================

func Test_CollectionDispose_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionDisposeTestCases {
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.Get("count")
		countInt := count.(int)

		col := namevalue.NewGenericCollectionDefault[string, string]()
		for i := 0; i < countInt; i++ {
			col.Add(namevalue.StringString{
				Name:  fmt.Sprintf("d%d", i),
				Value: fmt.Sprintf("v%d", i),
			})
		}

		col.Dispose()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", col.Items == nil))
	}
}

// ==========================================================================
// Test: ConcatNew
// ==========================================================================

func Test_CollectionConcatNew_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionConcatNewTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		original, _ := input.Get("original")
		extra, _ := input.Get("extra")
		originalInt := original.(int)
		extraInt := extra.(int)

		col := namevalue.NewGenericCollectionDefault[string, int]()
		for i := 0; i < originalInt; i++ {
			col.Add(namevalue.StringInt{Name: fmt.Sprintf("o%d", i), Value: i})
		}

		extraItems := make([]namevalue.StringInt, extraInt)
		for i := 0; i < extraInt; i++ {
			extraItems[i] = namevalue.StringInt{Name: fmt.Sprintf("e%d", i), Value: i + 100}
		}

		newCol := col.ConcatNew(extraItems...)

		// Act
		actual := args.Map{
			"mergedLength":   newCol.Length(),
			"originalLength": col.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// Test: StringMapAnyCollection — with values
// ==========================================================================

func Test_StringMapAnyCollection_WithValues(t *testing.T) {
	tc := stringMapAnyCollectionWithValuesTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	mapValues := input["mapValues"].([]map[string]any)

	// Act
	col := namevalue.NewGenericCollectionDefault[string, map[string]any]()
	for i, mapVal := range mapValues {
		col.Add(namevalue.StringMapAny{
			Name:  fmt.Sprintf("map%d", i),
			Value: mapVal,
		})
	}

	actual := args.Map{
		"length":    col.Length(),
		"hasValues": col.HasAnyItem(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: StringMapAnyCollection — nil value
// ==========================================================================

func Test_StringMapAnyCollection_NilValue(t *testing.T) {
	tc := stringMapAnyCollectionNilValueTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	mapValues := input["mapValues"].([]map[string]any)

	// Act
	col := namevalue.NewGenericCollectionDefault[string, map[string]any]()
	for i, mapVal := range mapValues {
		col.Add(namevalue.StringMapAny{
			Name:  fmt.Sprintf("map%d", i),
			Value: mapVal,
		})
	}

	actual := args.Map{
		"length":   col.Length(),
		"isNilMap": col.HasAnyItem(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
