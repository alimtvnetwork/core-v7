package codestacktests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// stacksTo: Reflection-based method discovery
// =============================================================================

func Test_StacksTo_MethodDiscovery_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToMethodDiscoveryTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		rt := rv.Type()

		methodMap := map[string]bool{}
		for i := 0; i < rt.NumMethod(); i++ {
			methodMap[rt.Method(i).Name] = true
		}

		// Act
		actual := args.Map{
			"hasBytesMethod":             methodMap["Bytes"],
			"hasBytesDefaultMethod":      methodMap["BytesDefault"],
			"hasStringMethod":            methodMap["String"],
			"hasStringUsingFmtMethod":    methodMap["StringUsingFmt"],
			"hasJsonStringMethod":        methodMap["JsonString"],
			"hasJsonStringDefaultMethod": methodMap["JsonStringDefault"],
			"hasStringNoCountMethod":     methodMap["StringNoCount"],
			"hasStringDefaultMethod":     methodMap["StringDefault"],
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// stacksTo: Bytes via reflection
// =============================================================================

func Test_StacksTo_Bytes_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToBytesTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		method := rv.MethodByName("Bytes")
		input := testCase.ArrangeInput.(args.Map)
		skipIndex, _ := input.GetAsInt("skipIndex")

		// Act
		results := method.Call([]reflect.Value{
			reflect.ValueOf(skipIndex),
		})
		bytes := results[0].Bytes()

		actual := args.Map{
			"notEmpty": len(bytes) > 0,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// stacksTo: BytesDefault via reflection
// =============================================================================

func Test_StacksTo_BytesDefault_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToBytesDefaultTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		method := rv.MethodByName("BytesDefault")

		// Act
		results := method.Call(nil)
		bytes := results[0].Bytes()

		actual := args.Map{
			"notEmpty": len(bytes) > 0,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// stacksTo: String via reflection
// =============================================================================

func Test_StacksTo_String_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToStringTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		method := rv.MethodByName("String")
		input := testCase.ArrangeInput.(args.Map)
		skipIndex, _ := input.GetAsInt("skipIndex")
		count, _ := input.GetAsInt("count")

		// Act
		results := method.Call([]reflect.Value{
			reflect.ValueOf(skipIndex),
			reflect.ValueOf(count),
		})
		str := results[0].String()

		actual := args.Map{
			"notEmpty": str != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// stacksTo: StringUsingFmt via reflection
// =============================================================================

func Test_StacksTo_StringUsingFmt_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToStringUsingFmtTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		method := rv.MethodByName("StringUsingFmt")

		formatter := codestack.Formatter(func(trace *codestack.Trace) string {
			return fmt.Sprintf(
				"%s:%d",
				trace.PackageMethodName,
				trace.Line,
			)
		})

		// Act
		results := method.Call([]reflect.Value{
			reflect.ValueOf(formatter),
			reflect.ValueOf(0),
			reflect.ValueOf(5),
		})
		str := results[0].String()

		actual := args.Map{
			"notEmpty": str != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// stacksTo: JsonString via reflection
// =============================================================================

func Test_StacksTo_JsonString_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToJsonStringTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		method := rv.MethodByName("JsonString")
		input := testCase.ArrangeInput.(args.Map)
		skipIndex, _ := input.GetAsInt("skipIndex")

		// Act
		results := method.Call([]reflect.Value{
			reflect.ValueOf(skipIndex),
		})
		str := results[0].String()

		actual := args.Map{
			"notEmpty": str != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// stacksTo: JsonStringDefault via reflection
// =============================================================================

func Test_StacksTo_JsonStringDefault_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToJsonStringDefaultTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		method := rv.MethodByName("JsonStringDefault")

		// Act
		results := method.Call(nil)
		str := results[0].String()

		actual := args.Map{
			"notEmpty": str != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// stacksTo: StringNoCount via reflection
// =============================================================================

func Test_StacksTo_StringNoCount_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToStringNoCountTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		method := rv.MethodByName("StringNoCount")
		input := testCase.ArrangeInput.(args.Map)
		skipIndex, _ := input.GetAsInt("skipIndex")

		// Act
		results := method.Call([]reflect.Value{
			reflect.ValueOf(skipIndex),
		})
		str := results[0].String()

		actual := args.Map{
			"notEmpty": str != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// stacksTo: StringDefault via reflection
// =============================================================================

func Test_StacksTo_StringDefault_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extStacksToStringDefaultTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.StacksTo)
		method := rv.MethodByName("StringDefault")

		// Act
		results := method.Call(nil)
		str := results[0].String()

		actual := args.Map{
			"notEmpty": str != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// currentNameOf
// =============================================================================

func Test_NameOf_All_Ext(t *testing.T) {
	for caseIndex, testCase := range extNameOfAllTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		fullName, _ := input.GetAsString("fullName")

		// Act
		fullMethod, pkgName, methodName := codestack.NameOf.All(fullName)

		actual := args.Map{
			"hasFullMethod":  fullMethod != "",
			"hasPackageName": pkgName != "",
			"hasMethodName":  methodName != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NameOf_Method_Ext(t *testing.T) {
	for caseIndex, testCase := range extNameOfMethodTestCases {
		// Act
		methodName := codestack.NameOf.Method()

		actual := args.Map{
			"hasMethodName": methodName != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NameOf_Package_Ext(t *testing.T) {
	for caseIndex, testCase := range extNameOfPackageTestCases {
		// Act
		pkgName := codestack.NameOf.Package()

		actual := args.Map{
			"hasPackageName": pkgName != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NameOf_MethodByFullName_Ext(t *testing.T) {
	// Arrange
	fullName := "github.com/mypackage.(*MyStruct).DoWork"

	// Act
	methodName := codestack.NameOf.MethodByFullName(fullName)

	// Assert
	actual := args.Map{"result": methodName == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MethodByFullName should return non-empty", actual)
}

func Test_NameOf_PackageByFullName_Ext(t *testing.T) {
	// Arrange
	fullName := "github.com/mypackage.(*MyStruct).DoWork"

	// Act
	pkgName := codestack.NameOf.PackageByFullName(fullName)

	// Assert
	actual := args.Map{"result": pkgName == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PackageByFullName should return non-empty", actual)
}

func Test_NameOf_JoinPackageNameWithRelative_Ext(t *testing.T) {
	// Arrange
	fullName := "github.com/mypackage.(*MyStruct).DoWork"

	// Act
	result := codestack.NameOf.JoinPackageNameWithRelative(
		fullName,
		"SubStruct.Method",
	)

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JoinPackageNameWithRelative should return non-empty", actual)
}

func Test_NameOf_CurrentFuncFullPath_Ext(t *testing.T) {
	// Arrange
	fullName := "github.com/mypackage.(*MyStruct).DoWork"

	// Act
	result := codestack.NameOf.CurrentFuncFullPath(fullName)

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CurrentFuncFullPath should return non-empty", actual)
}

func Test_NameOf_MethodStackSkip_Ext(t *testing.T) {
	// Act
	result := codestack.NameOf.MethodStackSkip(0)

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MethodStackSkip should return non-empty", actual)
}

func Test_NameOf_PackageStackSkip_Ext(t *testing.T) {
	// Act
	result := codestack.NameOf.PackageStackSkip(0)

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PackageStackSkip should return non-empty", actual)
}

// =============================================================================
// newCreator
// =============================================================================

func Test_NewCreator_Default_Ext(t *testing.T) {
	for caseIndex, testCase := range extNewCreatorDefaultTestCases {
		// Act
		trace := codestack.New.Default()

		actual := args.Map{
			"isOkay":        trace.IsOkay,
			"hasFilePath":   trace.FilePath != "",
			"hasMethodName": trace.MethodName != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewCreator_SkipOne_Ext(t *testing.T) {
	for caseIndex, testCase := range extNewCreatorSkipOneTestCases {
		// Act
		trace := codestack.New.SkipOne()

		actual := args.Map{
			"isOkay":      trace.IsOkay,
			"hasFilePath": trace.FilePath != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewCreator_Ptr_Ext(t *testing.T) {
	for caseIndex, testCase := range extNewCreatorPtrTestCases {
		// Act
		trace := codestack.New.Ptr(0)

		actual := args.Map{
			"isNil":  trace == nil,
			"isOkay": trace.IsOkay,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// newStacksCreator via reflection
// =============================================================================

func Test_NewStacksCreator_Default_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extNewStacksCreatorDefaultTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.New.StackTrace)
		method := rv.MethodByName("Default")

		// Act
		results := method.Call([]reflect.Value{
			reflect.ValueOf(0),
			reflect.ValueOf(3),
		})
		collection := results[0].Interface().(codestack.TraceCollection)

		actual := args.Map{
			"hasItems": collection.HasAnyItem(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewStacksCreator_SkipOne_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extNewStacksCreatorSkipOneTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.New.StackTrace)
		method := rv.MethodByName("SkipOne")

		// Act
		results := method.Call(nil)
		collection := results[0].Interface().(codestack.TraceCollection)

		actual := args.Map{
			"hasItems": collection.HasAnyItem(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewStacksCreator_SkipNone_Reflection_Ext(t *testing.T) {
	for caseIndex, testCase := range extNewStacksCreatorSkipNoneTestCases {
		// Arrange
		rv := reflect.ValueOf(codestack.New.StackTrace)
		method := rv.MethodByName("SkipNone")

		// Act
		results := method.Call(nil)
		collection := results[0].Interface().(codestack.TraceCollection)

		actual := args.Map{
			"hasItems": collection.HasAnyItem(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewStacksCreator_DefaultCount_Reflection_Ext(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(codestack.New.StackTrace)
	method := rv.MethodByName("DefaultCount")

	// Act
	results := method.Call([]reflect.Value{
		reflect.ValueOf(0),
	})
	collection := results[0].Interface().(codestack.TraceCollection)

	// Assert
	actual := args.Map{"result": collection.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DefaultCount should return items", actual)
}

func Test_NewStacksCreator_All_Reflection_Ext(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(codestack.New.StackTrace)
	method := rv.MethodByName("All")

	// Act
	results := method.Call([]reflect.Value{
		reflect.ValueOf(true),
		reflect.ValueOf(true),
		reflect.ValueOf(0),
		reflect.ValueOf(5),
	})
	collection := results[0].Interface().(codestack.TraceCollection)

	// Assert
	actual := args.Map{"result": collection.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "All should return items", actual)
}

// =============================================================================
// dirGetter
// =============================================================================

func Test_Dir_CurDir_Ext(t *testing.T) {
	for caseIndex, testCase := range extDirGetterCurDirTestCases {
		// Act
		result := codestack.Dir.CurDir()

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Dir_RepoDir_Ext(t *testing.T) {
	for caseIndex, testCase := range extDirGetterRepoDirTestCases {
		// Act
		result := codestack.Dir.RepoDir()

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Dir_Get_Ext(t *testing.T) {
	// Act
	result := codestack.Dir.Get(0)

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Dir.Get should return non-empty", actual)
}

func Test_Dir_CurDirJoin_Ext(t *testing.T) {
	// Act
	result := codestack.Dir.CurDirJoin("sub", "path")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Dir.CurDirJoin should return non-empty", actual)
}

func Test_Dir_RepoDirJoin_Ext(t *testing.T) {
	// Act
	result := codestack.Dir.RepoDirJoin("sub")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Dir.RepoDirJoin should return non-empty", actual)
}

// =============================================================================
// fileGetter
// =============================================================================

func Test_File_Path_Ext(t *testing.T) {
	for caseIndex, testCase := range extFileGetterPathTestCases {
		// Act
		result := codestack.File.Path(0)

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_File_Name_Ext(t *testing.T) {
	for caseIndex, testCase := range extFileGetterNameTestCases {
		// Act
		result := codestack.File.Name(0)

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_File_CurrentFilePath_Ext(t *testing.T) {
	for caseIndex, testCase := range extFileGetterCurrentFilePathTestCases {
		// Act
		result := codestack.File.CurrentFilePath()

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_File_PathLineSep_Ext(t *testing.T) {
	// Act
	filePath, lineNumber := codestack.File.PathLineSep(0)

	// Assert
	actual := args.Map{"result": filePath == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PathLineSep should return non-empty file path", actual)
	actual = args.Map{"result": lineNumber <= 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PathLineSep should return positive line number", actual)
}

func Test_File_PathLineSepDefault_Ext(t *testing.T) {
	// Act
	filePath, lineNumber := codestack.File.PathLineSepDefault()

	// Assert
	actual := args.Map{"result": filePath == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PathLineSepDefault should return non-empty file path", actual)
	actual = args.Map{"result": lineNumber <= 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PathLineSepDefault should return positive line number", actual)
}

func Test_File_FilePathWithLineString_Ext(t *testing.T) {
	// Act
	result := codestack.File.FilePathWithLineString(0)

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FilePathWithLineString should return non-empty", actual)
}

func Test_File_PathLineStringDefault_Ext(t *testing.T) {
	// Act
	result := codestack.File.PathLineStringDefault()

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PathLineStringDefault should return non-empty", actual)
}

// =============================================================================
// TraceCollection: additional uncovered methods via reflection
// =============================================================================

func Test_TraceCollection_StackTracesJsonResult_Reflection_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()
	rv := reflect.ValueOf(collection).Elem()
	method := rv.MethodByName("StackTracesJsonResult")

	// Act
	results := method.Call(nil)

	// Assert
	actual := args.Map{"result": results[0].IsNil()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StackTracesJsonResult should not be nil", actual)
}

func Test_TraceCollection_NewStackTraces_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.NewStackTraces(0)

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NewStackTraces should return non-empty", actual)
}

func Test_TraceCollection_NewDefaultStackTraces_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.NewDefaultStackTraces()

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NewDefaultStackTraces should return non-empty", actual)
}

func Test_TraceCollection_NewStackTracesJsonResult_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.NewStackTracesJsonResult(0)

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NewStackTracesJsonResult should not be nil", actual)
}

func Test_TraceCollection_NewDefaultStackTracesJsonResult_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.NewDefaultStackTracesJsonResult()

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NewDefaultStackTracesJsonResult should not be nil", actual)
}

func Test_TraceCollection_GetPagedCollection_Ext(t *testing.T) {
	// Arrange
	collection := &codestack.TraceCollection{
		Items: []codestack.Trace{
			newTestTrace("p1", 1),
			newTestTrace("p2", 2),
			newTestTrace("p3", 3),
			newTestTrace("p4", 4),
			newTestTrace("p5", 5),
		},
	}

	// Act
	pages := collection.GetPagedCollection(2)

	// Assert
	actual := args.Map{"result": len(pages) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 pages", actual)
}

func Test_TraceCollection_GetSinglePageCollection_Ext(t *testing.T) {
	// Arrange
	collection := &codestack.TraceCollection{
		Items: []codestack.Trace{
			newTestTrace("p1", 1),
			newTestTrace("p2", 2),
			newTestTrace("p3", 3),
			newTestTrace("p4", 4),
			newTestTrace("p5", 5),
		},
	}

	// Act
	page := collection.GetSinglePageCollection(2, 2)

	// Assert
	actual := args.Map{"result": page.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 items", actual)
}

func Test_TraceCollection_GetSinglePageCollection_SmallList_Ext(t *testing.T) {
	// Arrange
	collection := &codestack.TraceCollection{
		Items: []codestack.Trace{
			newTestTrace("p1", 1),
		},
	}

	// Act
	page := collection.GetSinglePageCollection(5, 1)

	// Assert
	actual := args.Map{"result": page.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 item", actual)
}

func Test_TraceCollection_FilterWithLimit_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.FilterWithLimit(
		2,
		func(trace *codestack.Trace) (bool, bool) {
			return true, false
		},
	)

	// Assert
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 items", actual)
}

func Test_TraceCollection_FilterTraceCollection_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.FilterTraceCollection(
		func(trace *codestack.Trace) (bool, bool) {
			return trace.PackageName == "pkg1", false
		},
	)

	// Assert
	actual := args.Map{"result": result.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 item", actual)
}

func Test_TraceCollection_SkipFilterMethodName_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.SkipFilterMethodNameTraceCollection("TestMethod")

	// Assert
	actual := args.Map{"result": result.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 items", actual)
}

func Test_TraceCollection_FilterFullMethodName_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.FilterFullMethodNameTraceCollection("pkg1.TestMethod")

	// Assert
	actual := args.Map{"result": result.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 item", actual)
}

func Test_TraceCollection_SkipFilterFullMethodName_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.SkipFilterFullMethodNameTraceCollection("pkg1.TestMethod")

	// Assert
	actual := args.Map{"result": result.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 items", actual)
}

func Test_TraceCollection_SkipFilterFilename_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.SkipFilterFilenameTraceCollection("/src/pkg1/file.go")

	// Assert
	actual := args.Map{"result": result.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 items", actual)
}

func Test_TraceCollection_FileWithLinesStrings_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	strs := collection.FileWithLinesStrings()

	// Assert
	actual := args.Map{"result": len(strs) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_TraceCollection_FileWithLinesString_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.FileWithLinesString()

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_TraceCollection_FileWithLinesString_Empty_Ext(t *testing.T) {
	// Arrange
	collection := emptyCollection()

	// Act
	result := collection.FileWithLinesString()

	// Assert
	actual := args.Map{"result": result != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return empty", actual)
}

func Test_TraceCollection_JoinShortStrings_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.JoinShortStrings(",")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_TraceCollection_JoinFileWithLinesStrings_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.JoinFileWithLinesStrings(",")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_TraceCollection_JoinJsonStrings_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.JoinJsonStrings(",")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_TraceCollection_JsonStrings_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.JsonStrings()

	// Assert
	actual := args.Map{"result": len(result) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_TraceCollection_Join_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.Join(",")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_TraceCollection_JoinLines_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.JoinLines()

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_TraceCollection_JoinCsv_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.JoinCsv()

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_TraceCollection_JoinCsvLine_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.JoinCsvLine()

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_TraceCollection_CodeStacksStringLimit_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.CodeStacksStringLimit(2)

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_TraceCollection_CodeStacksStringLimit_Empty_Ext(t *testing.T) {
	// Arrange
	collection := emptyCollection()

	// Act
	result := collection.CodeStacksStringLimit(2)

	// Assert
	actual := args.Map{"result": result != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return empty", actual)
}

func Test_TraceCollection_CsvStrings_Empty_Ext(t *testing.T) {
	// Arrange
	collection := emptyCollection()

	// Act
	result := collection.CsvStrings()

	// Assert
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_TraceCollection_FirstDynamic_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.FirstDynamic()

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-nil", actual)
}

func Test_TraceCollection_LastDynamic_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.LastDynamic()

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-nil", actual)
}

func Test_TraceCollection_FirstOrDefaultDynamic_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.FirstOrDefaultDynamic()

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-nil", actual)
}

func Test_TraceCollection_LastOrDefaultDynamic_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.LastOrDefaultDynamic()

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-nil", actual)
}

func Test_TraceCollection_SkipDynamic_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.SkipDynamic(1)

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-nil", actual)
}

func Test_TraceCollection_Skip_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.Skip(1)

	// Assert
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_TraceCollection_TakeDynamic_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.TakeDynamic(2)

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-nil", actual)
}

func Test_TraceCollection_Take_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.Take(2)

	// Assert
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_TraceCollection_LimitDynamic_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.LimitDynamic(2)

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-nil", actual)
}

func Test_TraceCollection_Limit_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.Limit(2)

	// Assert
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_TraceCollection_Count_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	count := collection.Count()

	// Assert
	actual := args.Map{"result": count != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_TraceCollection_JsonModel_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	model := collection.JsonModel()

	// Assert
	actual := args.Map{"result": len(model) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_TraceCollection_JsonModelAny_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.JsonModelAny()

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-nil", actual)
}

func Test_TraceCollection_Json_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.Json()

	// Assert
	actual := args.Map{"result": result.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have error:", actual)
}

func Test_TraceCollection_AsJsonContractsBinder_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	binder := collection.AsJsonContractsBinder()

	// Assert
	actual := args.Map{"result": binder == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_TraceCollection_AsJsoner_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	jsoner := collection.AsJsoner()

	// Assert
	actual := args.Map{"result": jsoner == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_TraceCollection_AsJsonParseSelfInjector_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	injector := collection.AsJsonParseSelfInjector()

	// Assert
	actual := args.Map{"result": injector == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_TraceCollection_ParseInjectUsingJson_Ext(t *testing.T) {
	// Arrange
	original := newTestCollection()
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.TraceCollection
	result, err := target.ParseInjectUsingJson(jsonResult)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": result.Length() != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_TraceCollection_ParseInjectUsingJsonMust_Ext(t *testing.T) {
	// Arrange
	original := newTestCollection()
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.TraceCollection
	result := target.ParseInjectUsingJsonMust(jsonResult)

	// Assert
	actual := args.Map{"result": result.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_TraceCollection_JsonParseSelfInject_Ext(t *testing.T) {
	// Arrange
	original := newTestCollection()
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.TraceCollection
	err := target.JsonParseSelfInject(jsonResult)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_TraceCollection_IsEqualItems_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.IsEqualItems(collection.Items...)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_TraceCollection_IsEqualItems_DiffLength_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.IsEqualItems(newTestTrace("pkg1", 10))

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be equal", actual)
}

func Test_TraceCollection_ConcatNewPtr_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()
	trace := newTestTrace("extra", 99)

	// Act
	result := collection.ConcatNewPtr(&trace)

	// Assert
	actual := args.Map{"result": result.Length() != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_TraceCollection_ConcatNewUsingSkipPlusCount_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.ConcatNewUsingSkipPlusCount(0, 3)

	// Assert
	actual := args.Map{"result": result.Length() < 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
}

func Test_TraceCollection_ConcatNewUsingSkip_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.ConcatNewUsingSkip(0)

	// Assert
	actual := args.Map{"result": result.Length() < 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
}

func Test_TraceCollection_InsertAt_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	inserted := newTestTrace("inserted", 99)
	collection.InsertAt(1, inserted)

	// Assert
	actual := args.Map{"result": collection.Items[1].PackageName != "inserted"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "item should be inserted at index 1", actual)
}

func Test_TraceCollection_AddsPtr_Valid_Ext(t *testing.T) {
	// Arrange
	collection := emptyCollection()
	trace := newTestTrace("valid", 1)

	// Act
	collection.AddsPtr(false, &trace)

	// Assert
	actual := args.Map{"result": collection.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_TraceCollection_Adds_Empty_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	collection.Adds()

	// Assert
	actual := args.Map{"result": collection.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_TraceCollection_AddsPtr_Empty_Ext(t *testing.T) {
	// Arrange
	collection := emptyCollection()

	// Act
	collection.AddsPtr(false)

	// Assert
	actual := args.Map{"result": collection.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_TraceCollection_StringsUsingFmt_Ext(t *testing.T) {
	// Arrange
	collection := newTestCollection()

	// Act
	result := collection.StringsUsingFmt(func(trace *codestack.Trace) string {
		return trace.PackageName
	})

	// Assert
	actual := args.Map{"result": len(result) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_TraceCollection_AddsUsingSkipDefault_Ext(t *testing.T) {
	// Arrange
	collection := emptyCollection()

	// Act
	collection.AddsUsingSkipDefault(0)

	// Assert - should have items from current stack
	actual := args.Map{"result": collection.Length() == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have items from stack", actual)
}

func Test_TraceCollection_AddsUsingSkipUsingFilter_Ext(t *testing.T) {
	// Arrange
	collection := emptyCollection()

	// Act
	collection.AddsUsingSkipUsingFilter(
		true,
		true,
		0,
		10,
		func(trace *codestack.Trace) (bool, bool) {
			return true, false
		},
	)

	// Assert
	actual := args.Map{"result": collection.Length() == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have items from filtered stack", actual)
}

// =============================================================================
// Trace: additional Json methods via reflection
// =============================================================================

func Test_Trace_Json_Ext(t *testing.T) {
	// Arrange
	trace := codestack.Trace{
		PackageName:       "pkg",
		MethodName:        "Method",
		PackageMethodName: "pkg.Method",
		FilePath:          "/file.go",
		Line:              10,
		IsOkay:            true,
	}

	// Act
	result := trace.Json()

	// Assert
	actual := args.Map{"result": result.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have error:", actual)
}

func Test_Trace_JsonPtr_Ext(t *testing.T) {
	// Arrange
	trace := codestack.Trace{
		PackageName: "pkg",
		FilePath:    "/file.go",
		Line:        10,
	}

	// Act
	result := trace.JsonPtr()

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_Trace_JsonModel_Ext(t *testing.T) {
	// Arrange
	trace := codestack.Trace{PackageName: "pkg"}

	// Act
	model := trace.JsonModel()

	// Assert
	actual := args.Map{"result": model.PackageName != "pkg"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return same trace as model", actual)
}

func Test_Trace_JsonModelAny_Ext(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{PackageName: "pkg"}

	// Act
	result := trace.JsonModelAny()

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_Trace_JsonString_Ext(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{
		PackageName: "pkg",
		FilePath:    "/file.go",
		Line:        10,
	}

	// Act
	result := trace.JsonString()

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Trace_ParseInjectUsingJson_Ext(t *testing.T) {
	// Arrange
	original := codestack.Trace{
		PackageName:       "pkg",
		PackageMethodName: "pkg.M",
		FilePath:          "/f.go",
		Line:              10,
	}
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.Trace
	result, err := target.ParseInjectUsingJson(jsonResult)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": result.PackageName != "pkg"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should preserve PackageName", actual)
}

func Test_Trace_ParseInjectUsingJsonMust_Ext(t *testing.T) {
	// Arrange
	original := codestack.Trace{PackageName: "pkg", FilePath: "/f.go"}
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.Trace
	result := target.ParseInjectUsingJsonMust(jsonResult)

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_Trace_JsonParseSelfInject_Ext(t *testing.T) {
	// Arrange
	original := codestack.Trace{PackageName: "pkg", FilePath: "/f.go"}
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.Trace
	err := target.JsonParseSelfInject(jsonResult)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_Trace_NilDispose_Ext(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act - should not panic
	trace.Dispose()
}

// =============================================================================
// FileWithLine: additional methods
// =============================================================================

func Test_FileWithLine_JsonModel_Ext(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{FilePath: "/f.go", Line: 5}

	// Act
	model := fwl.JsonModel()

	// Assert
	actual := args.Map{"result": model.FilePath != "/f.go"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return same model", actual)
}

func Test_FileWithLine_JsonModelAny_Ext(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/f.go", Line: 5}

	// Act
	result := fwl.JsonModelAny()

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_FileWithLine_Json_Ext(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{FilePath: "/f.go", Line: 5}

	// Act
	result := fwl.Json()

	// Assert
	actual := args.Map{"result": result.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_FileWithLine_JsonPtr_Ext(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{FilePath: "/f.go", Line: 5}

	// Act
	result := fwl.JsonPtr()

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_FileWithLine_ParseInjectUsingJson_Ext(t *testing.T) {
	// Arrange
	original := codestack.FileWithLine{FilePath: "/f.go", Line: 5}
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.FileWithLine
	result, err := target.ParseInjectUsingJson(jsonResult)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": result.FilePath != "/f.go"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should preserve FilePath", actual)
}

func Test_FileWithLine_ParseInjectUsingJsonMust_Ext(t *testing.T) {
	// Arrange
	original := codestack.FileWithLine{FilePath: "/f.go", Line: 5}
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.FileWithLine
	result := target.ParseInjectUsingJsonMust(jsonResult)

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_FileWithLine_JsonParseSelfInject_Ext(t *testing.T) {
	// Arrange
	original := codestack.FileWithLine{FilePath: "/f.go", Line: 5}
	jsonResult := original.JsonPtr()

	// Act
	var target codestack.FileWithLine
	err := target.JsonParseSelfInject(jsonResult)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_FileWithLine_FileWithLine_Method_Ext(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/f.go", Line: 42}

	// Act
	result := fwl.FileWithLine()

	// Assert
	actual := args.Map{"result": result != "/f.go:42"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '/f.go:42', got ''", actual)
}

// =============================================================================
// Trace: Message caching via reflection (tests lazy once)
// =============================================================================

func Test_Trace_Message_CachesOnSecondCall_Ext(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{
		PackageName:       "pkg",
		MethodName:        "Method",
		PackageMethodName: "pkg.Method",
		FilePath:          "/file.go",
		Line:              10,
		IsOkay:            true,
	}

	// Act
	msg1 := trace.Message()
	msg2 := trace.Message() // cached path

	// Assert
	actual := args.Map{"result": msg1 != msg2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Message should return same value on second call", actual)
}

func Test_Trace_ShortString_CachesOnSecondCall_Ext(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{
		PackageName:       "pkg",
		MethodName:        "Method",
		PackageMethodName: "pkg.Method",
		FilePath:          "/file.go",
		Line:              10,
		IsOkay:            true,
	}

	// Act
	s1 := trace.ShortString()
	s2 := trace.ShortString() // cached path

	// Assert
	actual := args.Map{"result": s1 != s2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ShortString should return same value on second call", actual)
}
