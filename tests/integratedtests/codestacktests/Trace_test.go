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

func Test_Trace_Dispose(t *testing.T) {
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

func Test_Trace_Clone(t *testing.T) {
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

func Test_Trace_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act
	cloned := trace.ClonePtr()

	// Assert
	if cloned != nil {
		t.Error("expected nil ClonePtr to return nil")
	}
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
	if msg == "" {
		t.Error("expected Message to not be empty")
	}
	if shortStr == "" {
		t.Error("expected ShortString to not be empty")
	}
	if str == "" {
		t.Error("expected String to not be empty")
	}
}

func Test_Trace_FileWithLine(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{
		FilePath: "/src/file.go",
		Line:     10,
		IsOkay:   true,
	}

	// Act
	fwl := trace.FileWithLine()

	// Assert
	if fwl.FilePath != "/src/file.go" {
		t.Errorf("expected FilePath '/src/file.go', got '%s'", fwl.FilePath)
	}
	if fwl.Line != 10 {
		t.Errorf("expected Line 10, got %d", fwl.Line)
	}
}

func Test_Trace_FileName(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{
		FilePath: "/src/mypkg/handler.go",
		IsOkay:   true,
	}

	// Act
	fileName := trace.FileName()

	// Assert
	if fileName != "handler.go" {
		t.Errorf("expected FileName 'handler.go', got '%s'", fileName)
	}
}

func Test_Trace_HasIssues(t *testing.T) {
	// Arrange - empty method name
	trace := &codestack.Trace{
		PackageName:       "pkg",
		PackageMethodName: "",
		IsOkay:            true,
	}

	// Act & Assert
	if !trace.HasIssues() {
		t.Error("expected HasIssues=true when PackageMethodName is empty")
	}
}

func Test_Trace_FileWithLineString(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{
		FilePath: "/src/file.go",
		Line:     25,
	}

	// Act
	result := trace.FileWithLineString()

	// Assert
	if result == "" {
		t.Error("expected FileWithLineString to not be empty")
	}
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
	if result != "" {
		t.Errorf("expected nil FileWithLine String to be empty, got '%s'", result)
	}
}

func Test_FileWithLine_JsonString(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{
		FilePath: "/src/file.go",
		Line:     42,
	}

	// Act
	jsonStr := fwl.JsonString()

	// Assert
	if jsonStr == "" {
		t.Error("expected JsonString to not be empty")
	}
}

func Test_FileWithLine_AsFileLiner(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{
		FilePath: "/src/test.go",
		Line:     5,
	}

	// Act
	liner := fwl.AsFileLiner()

	// Assert
	if liner == nil {
		t.Error("expected AsFileLiner to not be nil")
	}
	if liner.FullFilePath() != "/src/test.go" {
		t.Errorf("expected path '/src/test.go', got '%s'", liner.FullFilePath())
	}
}

func Test_Trace_AsFileLiner(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{
		FilePath: "/src/trace.go",
		Line:     33,
		IsOkay:   true,
	}

	// Act
	liner := trace.AsFileLiner()

	// Assert
	if liner == nil {
		t.Error("expected AsFileLiner to not be nil")
	}
}

func Test_Trace_StringUsingFmt(t *testing.T) {
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
	if result != "pkg.Method" {
		t.Errorf("expected 'pkg.Method', got '%s'", result)
	}
}

func Test_FileWithLine_StringUsingFmt(t *testing.T) {
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
	if result != "/file.go" {
		t.Errorf("expected '/file.go', got '%s'", result)
	}
}
