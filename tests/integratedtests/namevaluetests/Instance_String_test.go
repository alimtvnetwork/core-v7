package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/namevalue"
)

func Test_Instance_String_Verification(t *testing.T) {
	for caseIndex, testCase := range instanceStringTestCases {
		// Arrange
		inst := namevalue.Instance[string, string]{Name: "key", Value: "val"}

		// Act
		result := inst.String()

		actual := args.Map{"notEmpty": result != ""}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Instance_JsonString_Verification(t *testing.T) {
	for caseIndex, testCase := range instanceJsonStringTestCases {
		// Arrange
		inst := namevalue.Instance[string, string]{Name: "key", Value: "val"}

		// Act
		result := inst.JsonString()

		actual := args.Map{"notEmpty": result != ""}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Instance_Dispose_Verification(t *testing.T) {
	for caseIndex, testCase := range instanceDisposeTestCases {
		// Arrange
		inst := namevalue.Instance[string, string]{Name: "key", Value: "val"}

		// Act
		inst.Dispose()

		actual := args.Map{
			"nameEmpty":  inst.Name == "",
			"valueEmpty": inst.Value == "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Instance_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range instanceNilTestCases {
		// Arrange
		var inst *namevalue.Instance[string, string]

		// Act
		actual := args.Map{"isNull": inst.IsNull()}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Collection_Operations_Verification(t *testing.T) {
	for caseIndex, testCase := range extCollectionTestCases {
		// Arrange
		col := namevalue.NewGenericCollectionDefault[string, string]()
		col.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
		col.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
		col.Add(namevalue.Instance[string, string]{Name: "c", Value: "3"})

		// Act
		actual := args.Map{
			"length":         col.Length(),
			"hasAnyItem":     col.HasAnyItem(),
			"isEmpty":        col.IsEmpty(),
			"stringsLen":     len(col.Strings()),
			"jsonStringsLen": len(col.JsonStrings()),
			"joinNotEmpty":   col.Join(", ") != "",
			"joinLinesOk":    col.JoinLines() != "",
			"joinCsvOk":      col.JoinCsv() != "",
			"joinCsvLineOk":  col.JoinCsvLine() != "",
			"stringOk":       col.String() != "",
			"jsonStringOk":   col.JsonString() != "",
			"csvStringsLen":  len(col.CsvStrings()),
			"cloneLen":       col.Clone().Length(),
			"errorNotNil":    col.Error() != nil,
			"errorMsgNotNil": col.ErrorUsingMessage("test") != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Collection_PrependAppend_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionPrependAppendTestCases {
		// Arrange
		col := namevalue.NewGenericCollectionDefault[string, string]()
		item1 := namevalue.Instance[string, string]{Name: "a", Value: "1"}
		item2 := namevalue.Instance[string, string]{Name: "b", Value: "2"}
		item3 := namevalue.Instance[string, string]{Name: "c", Value: "3"}
		item4 := namevalue.Instance[string, string]{Name: "d", Value: "4"}
		item5 := namevalue.Instance[string, string]{Name: "e", Value: "5"}

		// Act
		col.Add(item1)
		col.Prepend(item2)
		afterPrepend := col.Length()

		col.Append(item3)
		afterAppend := col.Length()

		col.AppendIf(true, item4)
		appendIfTrue := col.Length()

		col.AppendIf(false, item5)
		appendIfFalse := col.Length()

		col.PrependIf(true, item5)
		prependIfTrue := col.Length()

		col.PrependIf(false, item5)
		prependIfFalse := col.Length()

		actual := args.Map{
			"afterPrepend":   afterPrepend,
			"afterAppend":    afterAppend,
			"appendIfTrue":   appendIfTrue,
			"appendIfFalse":  appendIfFalse,
			"prependIfTrue":  prependIfTrue,
			"prependIfFalse": prependIfFalse,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Collection_IsEqual_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionIsEqualTestCases {
		// Arrange
		col1 := namevalue.NewGenericCollectionDefault[string, string]()
		col1.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
		col2 := col1.Clone()

		// Act
		actual := args.Map{"isEqual": col1.IsEqualByString(col2)}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Collection_ClonePtr_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionClonePtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)

		var result *namevalue.Collection[string, string]
		if isNil {
			var nilCol *namevalue.Collection[string, string]
			result = nilCol.ClonePtr()
		} else {
			col := namevalue.NewGenericCollectionDefault[string, string]()
			col.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
			result = col.ClonePtr()
		}

		// Act
		actual := args.Map{"isNil": result == nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Collection_AddsIf_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionAddsIfTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isAdd := input.GetAsBoolDefault("isAdd", false)
		col := namevalue.NewGenericCollectionDefault[string, string]()

		// Act
		col.AddsIf(isAdd, namevalue.Instance[string, string]{Name: "a", Value: "1"})

		actual := args.Map{"length": col.Length()}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Collection_ConcatNew_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionConcatTestCases {
		// Arrange
		col := namevalue.NewGenericCollectionDefault[string, string]()
		col.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})

		// Act
		newCol := col.ConcatNew(namevalue.Instance[string, string]{Name: "b", Value: "2"})

		actual := args.Map{
			"originalLen": col.Length(),
			"concatLen":   newCol.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Collection_FuncIf_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionFuncIfTestCases {
		// Arrange
		col := namevalue.NewGenericCollectionDefault[string, string]()
		col.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})

		item := namevalue.Instance[string, string]{Name: "x", Value: "y"}

		// Act
		col.PrependUsingFuncIf(true, func() []namevalue.Instance[string, string] {
			return []namevalue.Instance[string, string]{item}
		})
		afterPrependFunc := col.Length()

		col.AppendUsingFuncIf(true, func() []namevalue.Instance[string, string] {
			return []namevalue.Instance[string, string]{item}
		})
		afterAppendFunc := col.Length()

		actual := args.Map{
			"afterPrependFunc": afterPrependFunc,
			"afterAppendFunc":  afterAppendFunc,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Collection_AppendPrependIf_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionAppendPrependIfTestCases {
		// Arrange
		col := namevalue.NewGenericCollectionDefault[string, string]()
		col.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
		prepend := []namevalue.Instance[string, string]{{Name: "p", Value: "0"}}
		appnd := []namevalue.Instance[string, string]{{Name: "z", Value: "9"}}

		// Act
		col.AppendPrependIf(true, prepend, appnd)

		actual := args.Map{"length": col.Length()}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Collection_Clear_And_Dispose(t *testing.T) {
	// Arrange
	col := namevalue.NewGenericCollectionDefault[string, string]()
	col.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})

	// Act
	col.Clear()

	// Assert
	actual := args.Map{"result": col.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clear should make length 0", actual)

	// Arrange 2
	col2 := namevalue.NewGenericCollectionDefault[string, string]()
	col2.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})

	// Act 2
	col2.Dispose()

	// Assert 2 - should not panic
}

func Test_Collection_AddsPtr_FromInstanceString(t *testing.T) {
	// Arrange
	col := namevalue.NewGenericCollectionDefault[string, string]()
	item := &namevalue.Instance[string, string]{Name: "a", Value: "1"}

	// Act
	col.AddsPtr(item, nil)

	// Assert
	actual := args.Map{"result": col.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AddsPtr should add non-nil only", actual)
}

func Test_Collection_CompiledLazyString_FromInstanceString(t *testing.T) {
	// Arrange
	col := namevalue.NewGenericCollectionDefault[string, string]()
	col.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})

	// Act
	str1 := col.CompiledLazyString()
	str2 := col.CompiledLazyString()

	// Assert
	actual := args.Map{"result": str1 != str2 || str1 == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CompiledLazyString should return cached non-empty string", actual)
}

func Test_Collection_ConcatNewPtr_FromInstanceString(t *testing.T) {
	// Arrange
	col := namevalue.NewGenericCollectionDefault[string, string]()
	col.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	item := &namevalue.Instance[string, string]{Name: "b", Value: "2"}

	// Act
	newCol := col.ConcatNewPtr(item)

	// Assert
	actual := args.Map{"result": newCol.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ConcatNewPtr should have 2", actual)
}

func Test_Collection_JoinJsonStrings_FromInstanceString(t *testing.T) {
	// Arrange
	col := namevalue.NewGenericCollectionDefault[string, string]()
	col.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})

	// Act
	result := col.JoinJsonStrings(", ")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JoinJsonStrings should not be empty", actual)
}

func Test_Collection_NewGenericCollectionUsing(t *testing.T) {
	// Arrange
	items := []namevalue.Instance[string, string]{
		{Name: "a", Value: "1"},
		{Name: "b", Value: "2"},
	}

	// Act - no clone
	col1 := namevalue.NewGenericCollectionUsing(false, items...)

	// Act - clone
	col2 := namevalue.NewGenericCollectionUsing(true, items...)

	// Act - nil items
	col3 := namevalue.NewGenericCollectionUsing[string, string](false)

	// Assert
	actual := args.Map{"result": col1.Length() != 2 || col2.Length() != 2 || col3.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NewGenericCollectionUsing failed:", actual)
}
