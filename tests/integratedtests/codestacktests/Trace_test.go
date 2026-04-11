package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Trace_BasicProperties(t *testing.T) {
	for caseIndex, testCase := range traceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pkgName, _ := input.GetAsString("packageName")
		methodName, _ := input.GetAsString("methodName")
		pkgMethod, _ := input.GetAsString("pkgMethod")
		filePath, _ := input.GetAsString("filePath")
		line, _ := input.GetAsInt("line")

		// Act
		trace := &codestack.Trace{
			PackageName:       pkgName,
			MethodName:        methodName,
			PackageMethodName: pkgMethod,
			FilePath:          filePath,
			Line:              line,
			IsOkay:            true,
		}

		actual := args.Map{
			"packageName": trace.PackageName,
			"methodName":  trace.MethodName,
			"pkgMethod":   trace.PackageMethodName,
			"filePath":    trace.FullFilePath(),
			"lineNumber":  trace.LineNumber(),
			"isNil":       trace.IsNil(),
			"isNotNil":    trace.IsNotNil(),
			"hasIssues":   trace.HasIssues(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Trace_Nil(t *testing.T) {
	for caseIndex, testCase := range traceNilTestCases {
		// Arrange
		var trace *codestack.Trace

		// Act
		actual := args.Map{
			"isNil":    trace.IsNil(),
			"isNotNil": trace.IsNotNil(),
			"string":   trace.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Trace_Dispose_FromTrace(t *testing.T) {
	for caseIndex, testCase := range traceDisposeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pkgName, _ := input.GetAsString("packageName")
		methodName, _ := input.GetAsString("methodName")
		pkgMethod, _ := input.GetAsString("pkgMethod")
		filePath, _ := input.GetAsString("filePath")
		line, _ := input.GetAsInt("line")

		trace := &codestack.Trace{
			PackageName:       pkgName,
			MethodName:        methodName,
			PackageMethodName: pkgMethod,
			FilePath:          filePath,
			Line:              line,
			IsOkay:            true,
		}

		// Act
		trace.Dispose()

		actual := args.Map{
			"packageName": trace.PackageName,
			"methodName":  trace.MethodName,
			"pkgMethod":   trace.PackageMethodName,
			"filePath":    trace.FilePath,
			"lineNumber":  trace.Line,
			"isOkay":      trace.IsOkay,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Trace_Clone_FromTrace(t *testing.T) {
	for caseIndex, testCase := range traceCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pkgName, _ := input.GetAsString("packageName")
		methodName, _ := input.GetAsString("methodName")
		pkgMethod, _ := input.GetAsString("pkgMethod")
		filePath, _ := input.GetAsString("filePath")
		line, _ := input.GetAsInt("line")

		trace := &codestack.Trace{
			PackageName:       pkgName,
			MethodName:        methodName,
			PackageMethodName: pkgMethod,
			FilePath:          filePath,
			Line:              line,
			IsOkay:            true,
		}

		// Act
		cloned := trace.Clone()

		actual := args.Map{
			"packageName": cloned.PackageName,
			"methodName":  cloned.MethodName,
			"filePath":    cloned.FilePath,
			"lineNumber":  cloned.Line,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Trace_ClonePtr_Nil_FromTrace(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act
	cloned := trace.ClonePtr()

	// Assert
	actual := args.Map{"result": cloned != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil ClonePtr to return nil", actual)
}

func Test_Trace_Message_And_ShortString(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{
		PackageName:       "mypkg",
		MethodName:        "DoWork",
		PackageMethodName: "mypkg.DoWork",
		FilePath:          "/src/mypkg/work.go",
		Line:              55,
		IsOkay:            true,
	}

	// Act
	msg := trace.Message()
	shortStr := trace.ShortString()
	str := trace.String()

	// Assert
	actual := args.Map{"result": msg == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Message to not be empty", actual)
	actual = args.Map{"result": shortStr == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ShortString to not be empty", actual)
	actual = args.Map{"result": str == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected String to not be empty", actual)
}

func Test_Trace_FileWithLine_FromTrace(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{
		FilePath: "/src/file.go",
		Line:     10,
		IsOkay:   true,
	}

	// Act
	fwl := trace.FileWithLine()

	// Assert
	actual := args.Map{"result": fwl.FilePath != "/src/file.go"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected FilePath '/src/file.go', got ''", actual)
	actual = args.Map{"result": fwl.Line != 10}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Line 10", actual)
}

func Test_Trace_FileName_FromTrace(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{
		FilePath: "/src/mypkg/handler.go",
		IsOkay:   true,
	}

	// Act
	fileName := trace.FileName()

	// Assert
	actual := args.Map{"result": fileName != "handler.go"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected FileName 'handler.go', got ''", actual)
}

func Test_Trace_HasIssues_FromTrace(t *testing.T) {
	// Arrange - empty method name
	trace := &codestack.Trace{
		PackageName:       "pkg",
		PackageMethodName: "",
		IsOkay:            true,
	}

	// Act & Assert
	actual := args.Map{"result": trace.HasIssues()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasIssues=true when PackageMethodName is empty", actual)
}

func Test_Trace_FileWithLineString_FromTrace(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{
		FilePath: "/src/file.go",
		Line:     25,
	}

	// Act
	result := trace.FileWithLineString()

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected FileWithLineString to not be empty", actual)
}

func Test_FileWithLine_StringMethods(t *testing.T) {
	for caseIndex, testCase := range fileWithLineStringMethodTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		file, _ := input.GetAsString("file")
		line, _ := input.GetAsInt("line")

		fwl := &codestack.FileWithLine{
			FilePath: file,
			Line:     line,
		}

		// Act
		actual := args.Map{
			"isNil":    fwl.IsNil(),
			"isNotNil": fwl.IsNotNil(),
			"hasLine":  fwl.LineNumber() > 0,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_FileWithLine_NilString(t *testing.T) {
	// Arrange
	var fwl *codestack.FileWithLine

	// Act
	result := fwl.String()

	// Assert
	actual := args.Map{"result": result != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil FileWithLine String to be empty, got ''", actual)
}

func Test_FileWithLine_JsonString_FromTrace(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{
		FilePath: "/src/file.go",
		Line:     42,
	}

	// Act
	jsonStr := fwl.JsonString()

	// Assert
	actual := args.Map{"result": jsonStr == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected JsonString to not be empty", actual)
}

func Test_FileWithLine_AsFileLiner_FromTrace(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{
		FilePath: "/src/test.go",
		Line:     5,
	}

	// Act
	liner := fwl.AsFileLiner()

	// Assert
	actual := args.Map{"result": liner == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected AsFileLiner to not be nil", actual)
	actual = args.Map{"result": liner.FullFilePath() != "/src/test.go"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected path '/src/test.go', got ''", actual)
}

func Test_Trace_AsFileLiner_FromTrace(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{
		FilePath: "/src/trace.go",
		Line:     33,
		IsOkay:   true,
	}

	// Act
	liner := trace.AsFileLiner()

	// Assert
	actual := args.Map{"result": liner == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected AsFileLiner to not be nil", actual)
}

func Test_Trace_StringUsingFmt_FromTrace(t *testing.T) {
	// Arrange
	trace := codestack.Trace{
		PackageMethodName: "pkg.Method",
		Line:              10,
	}

	// Act
	result := trace.StringUsingFmt(func(tr codestack.Trace) string {
		return tr.PackageMethodName
	})

	// Assert
	actual := args.Map{"result": result != "pkg.Method"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'pkg.Method', got ''", actual)
}

func Test_FileWithLine_StringUsingFmt_FromTrace(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{
		FilePath: "/file.go",
		Line:     7,
	}

	// Act
	result := fwl.StringUsingFmt(func(f codestack.FileWithLine) string {
		return f.FilePath
	})

	// Assert
	actual := args.Map{"result": result != "/file.go"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '/file.go', got ''", actual)
}
