package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

func newTestTrace(pkgName string, line int) codestack.Trace {
	return codestack.Trace{
		PackageName:       pkgName,
		MethodName:        "TestMethod",
		PackageMethodName: pkgName + ".TestMethod",
		FilePath:          "/src/" + pkgName + "/file.go",
		Line:              line,
		IsOkay:            true,
	}
}

func newTestCollection() *codestack.TraceCollection {
	return &codestack.TraceCollection{
		Items: []codestack.Trace{
			newTestTrace("pkg1", 10),
			newTestTrace("pkg2", 20),
			newTestTrace("pkg3", 30),
		},
	}
}

func emptyCollection() *codestack.TraceCollection {
	return &codestack.TraceCollection{
		Items: []codestack.Trace{},
	}
}

func Test_TraceCollection_Basic(t *testing.T) {
	for caseIndex, testCase := range traceCollectionBasicTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.GetAsInt("count")

		// Act
		items := make([]codestack.Trace, 0, count)
		for i := 0; i < count; i++ {
			items = append(items, newTestTrace("pkg", i))
		}
		collection := &codestack.TraceCollection{Items: items}

		actual := args.Map{
			"length":     collection.Length(),
			"isEmpty":    collection.IsEmpty(),
			"hasAnyItem": collection.HasAnyItem(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_FirstLast(t *testing.T) {
	for caseIndex, testCase := range traceCollectionFirstLastTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		actual := args.Map{
			"firstPkg": collection.First().PackageName,
			"lastPkg":  collection.Last().PackageName,
			"lastIdx":  collection.LastIndex(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_SkipTake(t *testing.T) {
	for caseIndex, testCase := range traceCollectionSkipTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		skipVal, _ := input.GetAsInt("skip")
		collection := newTestCollection()

		// Act
		skipped := collection.SkipCollection(skipVal)

		actual := args.Map{
			"length":   skipped.Length(),
			"firstPkg": skipped.First().PackageName,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_Take(t *testing.T) {
	for caseIndex, testCase := range traceCollectionTakeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		takeVal, _ := input.GetAsInt("take")
		collection := newTestCollection()

		// Act
		taken := collection.TakeCollection(takeVal)

		actual := args.Map{
			"length":  taken.Length(),
			"lastPkg": taken.Last().PackageName,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_Reverse(t *testing.T) {
	for caseIndex, testCase := range traceCollectionReverseTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		collection.Reverse()

		actual := args.Map{
			"firstPkg": collection.First().PackageName,
			"lastPkg":  collection.Last().PackageName,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_FilterPackageName(t *testing.T) {
	for caseIndex, testCase := range traceCollectionFilterTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pkgName, _ := input.GetAsString("package")
		collection := newTestCollection()

		// Act
		filtered := collection.FilterPackageNameTraceCollection(pkgName)

		actual := args.Map{
			"length":   filtered.Length(),
			"firstPkg": filtered.First().PackageName,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_Add(t *testing.T) {
	for caseIndex, testCase := range traceCollectionAddTestCases {
		// Arrange
		collection := emptyCollection()

		// Act
		collection.Add(newTestTrace("add1", 1))
		collection.Add(newTestTrace("add2", 2))

		actual := args.Map{
			"length":   collection.Length(),
			"firstPkg": collection.First().PackageName,
			"lastPkg":  collection.Last().PackageName,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_Adds(t *testing.T) {
	for caseIndex, testCase := range traceCollectionAddsTestCases {
		// Arrange
		collection := emptyCollection()
		traces := []codestack.Trace{
			newTestTrace("a1", 1),
			newTestTrace("a2", 2),
			newTestTrace("a3", 3),
		}

		// Act
		collection.Adds(traces...)

		actual := args.Map{
			"length": collection.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_AddsIf_True(t *testing.T) {
	for caseIndex, testCase := range traceCollectionAddsIfTrueTestCases {
		// Arrange
		collection := emptyCollection()

		// Act
		collection.AddsIf(true, newTestTrace("added", 1))

		actual := args.Map{
			"length": collection.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_AddsIf_False(t *testing.T) {
	for caseIndex, testCase := range traceCollectionAddsIfFalseTestCases {
		// Arrange
		collection := emptyCollection()

		// Act
		collection.AddsIf(false, newTestTrace("skipped", 1))

		actual := args.Map{
			"length": collection.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_FirstOrDefault_Empty(t *testing.T) {
	for caseIndex, testCase := range traceCollectionFirstOrDefaultEmptyTestCases {
		// Arrange
		collection := emptyCollection()

		// Act
		result := collection.FirstOrDefault()

		actual := args.Map{
			"isNil": result == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_FirstOrDefault_NonEmpty(t *testing.T) {
	for caseIndex, testCase := range traceCollectionFirstOrDefaultNonEmptyTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		result := collection.FirstOrDefault()

		actual := args.Map{
			"isNil":      result == nil,
			"packageName": result.PackageName,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_LastOrDefault_Empty(t *testing.T) {
	for caseIndex, testCase := range traceCollectionLastOrDefaultEmptyTestCases {
		// Arrange
		collection := emptyCollection()

		// Act
		result := collection.LastOrDefault()

		actual := args.Map{
			"isNil": result == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_LastOrDefault_NonEmpty(t *testing.T) {
	for caseIndex, testCase := range traceCollectionLastOrDefaultNonEmptyTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		result := collection.LastOrDefault()

		actual := args.Map{
			"isNil":      result == nil,
			"packageName": result.PackageName,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_Strings(t *testing.T) {
	for caseIndex, testCase := range traceCollectionStringsTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		strs := collection.Strings()

		actual := args.Map{
			"length": len(strs),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_Filter(t *testing.T) {
	for caseIndex, testCase := range traceCollectionFilterFuncTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		filtered := collection.Filter(func(trace *codestack.Trace) (bool, bool) {
			return trace.Line > 10, false
		})

		actual := args.Map{
			"length": len(filtered),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_SkipFilterPackageName(t *testing.T) {
	for caseIndex, testCase := range traceCollectionSkipFilterPkgTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		filtered := collection.SkipFilterPackageNameTraceCollection("pkg2")

		actual := args.Map{
			"length": filtered.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_FilterMethodName(t *testing.T) {
	for caseIndex, testCase := range traceCollectionFilterMethodTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		filtered := collection.FilterMethodNameTraceCollection("TestMethod")

		actual := args.Map{
			"length": filtered.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_Clone(t *testing.T) {
	for caseIndex, testCase := range traceCollectionCloneTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		cloned := collection.Clone()

		actual := args.Map{
			"length":   cloned.Length(),
			"firstPkg": cloned.First().PackageName,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_ClonePtr(t *testing.T) {
	for caseIndex, testCase := range traceCollectionClonePtrTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		cloned := collection.ClonePtr()

		actual := args.Map{
			"isNil":  cloned == nil,
			"length": cloned.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_ClonePtr_Nil(t *testing.T) {
	for caseIndex, testCase := range traceCollectionClonePtrNilTestCases {
		// Arrange
		var collection *codestack.TraceCollection

		// Act
		cloned := collection.ClonePtr()

		actual := args.Map{
			"isNil": cloned == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_IsEqual(t *testing.T) {
	for caseIndex, testCase := range traceCollectionIsEqualTestCases {
		// Arrange
		c1 := newTestCollection()
		c2 := newTestCollection()

		// Act
		actual := args.Map{
			"isEqual": c1.IsEqual(c2),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_IsEqual_BothNil(t *testing.T) {
	for caseIndex, testCase := range traceCollectionIsEqualBothNilTestCases {
		// Arrange
		var c1 *codestack.TraceCollection
		var c2 *codestack.TraceCollection

		// Act
		actual := args.Map{
			"isEqual": c1.IsEqual(c2),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_IsEqual_OneNil(t *testing.T) {
	for caseIndex, testCase := range traceCollectionIsEqualOneNilTestCases {
		// Arrange
		c1 := newTestCollection()
		var c2 *codestack.TraceCollection

		// Act
		actual := args.Map{
			"isEqual": c1.IsEqual(c2),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_Length_Nil(t *testing.T) {
	for caseIndex, testCase := range traceCollectionLengthNilTestCases {
		// Arrange
		var collection *codestack.TraceCollection

		// Act
		actual := args.Map{
			"length": collection.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_HasIndex(t *testing.T) {
	for caseIndex, testCase := range traceCollectionHasIndexTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		idx, _ := input.GetAsInt("index")
		collection := newTestCollection()

		// Act
		actual := args.Map{
			"hasIndex": collection.HasIndex(idx),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_Reverse_Empty(t *testing.T) {
	for caseIndex, testCase := range traceCollectionReverseEmptyTestCases {
		// Arrange
		collection := emptyCollection()

		// Act
		result := collection.Reverse()

		actual := args.Map{
			"length": result.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_Reverse_Two(t *testing.T) {
	for caseIndex, testCase := range traceCollectionReverseTwoTestCases {
		// Arrange
		collection := &codestack.TraceCollection{
			Items: []codestack.Trace{
				newTestTrace("first", 1),
				newTestTrace("second", 2),
			},
		}

		// Act
		collection.Reverse()

		actual := args.Map{
			"firstPkg": collection.First().PackageName,
			"lastPkg":  collection.Last().PackageName,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_ConcatNew(t *testing.T) {
	for caseIndex, testCase := range traceCollectionConcatNewTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		newCollection := collection.ConcatNew(newTestTrace("new1", 40))

		actual := args.Map{
			"newLength":      newCollection.Length(),
			"originalLength": collection.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_Dispose(t *testing.T) {
	for caseIndex, testCase := range traceCollectionDisposeTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		collection.Dispose()

		actual := args.Map{
			"itemsNil": collection.Items == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_Dispose_Nil(t *testing.T) {
	// Arrange
	var collection *codestack.TraceCollection

	// Act - should not panic
	collection.Dispose()
}

func Test_TraceCollection_Clear(t *testing.T) {
	for caseIndex, testCase := range traceCollectionClearTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		collection.Clear()

		actual := args.Map{
			"length": collection.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_Clear_Nil(t *testing.T) {
	for caseIndex, testCase := range traceCollectionClearNilTestCases {
		// Arrange
		var collection *codestack.TraceCollection

		// Act
		result := collection.Clear()

		actual := args.Map{
			"isNil": result == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_CodeStacksString(t *testing.T) {
	for caseIndex, testCase := range traceCollectionCodeStacksStringTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		stacks := collection.CodeStacksString()

		actual := args.Map{
			"notEmpty": stacks != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_CodeStacksString_Empty(t *testing.T) {
	for caseIndex, testCase := range traceCollectionCodeStacksStringEmptyTestCases {
		// Arrange
		collection := emptyCollection()

		// Act
		stacks := collection.CodeStacksString()

		actual := args.Map{
			"isEmpty": stacks == "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_JsonString(t *testing.T) {
	for caseIndex, testCase := range traceCollectionJsonStringTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		jsonStr := collection.JsonString()

		actual := args.Map{
			"notEmpty": jsonStr != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_JsonString_Empty(t *testing.T) {
	for caseIndex, testCase := range traceCollectionJsonStringEmptyTestCases {
		// Arrange
		collection := emptyCollection()

		// Act
		jsonStr := collection.JsonString()

		actual := args.Map{
			"isEmpty": jsonStr == "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_Serializer(t *testing.T) {
	for caseIndex, testCase := range traceCollectionSerializerTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		bytes, err := collection.Serializer()

		actual := args.Map{
			"hasError": err != nil,
			"notEmpty": len(bytes) > 0,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_StackTracesBytes(t *testing.T) {
	for caseIndex, testCase := range traceCollectionStackTracesBytesTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		bytes := collection.StackTracesBytes()

		actual := args.Map{
			"notEmpty": len(bytes) > 0,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_StackTracesBytes_Empty(t *testing.T) {
	for caseIndex, testCase := range traceCollectionStackTracesBytesEmptyTestCases {
		// Arrange
		collection := emptyCollection()

		// Act
		bytes := collection.StackTracesBytes()

		actual := args.Map{
			"length": len(bytes),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_GetPagesSize(t *testing.T) {
	for caseIndex, testCase := range traceCollectionGetPagesSizeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pageSize, _ := input.GetAsInt("pageSize")
		collection := newTestCollection()

		// Act
		pages := collection.GetPagesSize(pageSize)

		actual := args.Map{
			"pages": pages,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_LimitCollection(t *testing.T) {
	for caseIndex, testCase := range traceCollectionLimitTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		limited := collection.LimitCollection(2)

		actual := args.Map{
			"length": limited.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_SafeLimitCollection(t *testing.T) {
	for caseIndex, testCase := range traceCollectionSafeLimitTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		limited := collection.SafeLimitCollection(10)

		actual := args.Map{
			"length": limited.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_FileWithLines(t *testing.T) {
	for caseIndex, testCase := range traceCollectionFileWithLinesTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		fwls := collection.FileWithLines()

		actual := args.Map{
			"length": len(fwls),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_ShortStrings(t *testing.T) {
	for caseIndex, testCase := range traceCollectionShortStringsTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		strs := collection.ShortStrings()

		actual := args.Map{
			"length": len(strs),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_JoinUsingFmt(t *testing.T) {
	for caseIndex, testCase := range traceCollectionJoinUsingFmtTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		joined := collection.JoinUsingFmt(func(trace *codestack.Trace) string {
			return trace.PackageName
		}, ",")

		actual := args.Map{
			"result": joined,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_CsvStrings(t *testing.T) {
	for caseIndex, testCase := range traceCollectionCsvStringsTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		strs := collection.CsvStrings()

		actual := args.Map{
			"length": len(strs),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_AddsPtr_Nil(t *testing.T) {
	for caseIndex, testCase := range traceCollectionAddsPtrNilTestCases {
		// Arrange
		collection := emptyCollection()
		var nilTrace *codestack.Trace

		// Act
		collection.AddsPtr(false, nilTrace)

		actual := args.Map{
			"length": collection.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_AddsPtr_SkipOnIssues(t *testing.T) {
	for caseIndex, testCase := range traceCollectionAddsPtrSkipTestCases {
		// Arrange
		collection := emptyCollection()
		badTrace := &codestack.Trace{
			PackageName:       "",
			PackageMethodName: "",
			IsOkay:            false,
		}

		// Act
		collection.AddsPtr(true, badTrace)

		actual := args.Map{
			"length": collection.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_String(t *testing.T) {
	for caseIndex, testCase := range traceCollectionStringTestCases {
		// Arrange
		collection := newTestCollection()

		// Act
		str := collection.String()

		actual := args.Map{
			"notEmpty": str != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TraceCollection_String_Empty(t *testing.T) {
	for caseIndex, testCase := range traceCollectionStringEmptyTestCases {
		// Arrange
		collection := emptyCollection()

		// Act
		str := collection.String()

		actual := args.Map{
			"isEmpty": str == "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
