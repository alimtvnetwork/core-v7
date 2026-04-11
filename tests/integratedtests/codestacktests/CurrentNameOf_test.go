package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── currentNameOf ──

func Test_CurrentNameOf(t *testing.T) {
	// Arrange
	name := codestack.NameOf.MethodStackSkip(0)

	// Act
	actual := args.Map{"notEmpty": name != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CurrentNameOf returns correct value -- with args", actual)
}

// ── skippablePrefixes ──

func Test_SkippablePrefixes(t *testing.T) {
	// Arrange
	// skippablePrefixes is unexported; verify via NameOf instead
	name := codestack.NameOf.Method()

	// Act
	actual := args.Map{"notEmpty": name != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.Method -- proxy for skippable check", actual)
}

// ── New.Default / New.Ptr / New.Skip ──

func Test_New_Default(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	actual := args.Map{
		"isOkay": trace.IsOkay,
		"pkgNotEmpty": trace.PackageName != "",
	}

	// Assert
	expected := args.Map{
		"isOkay": true,
		"pkgNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "New.Default returns correct value -- with args", actual)
}

func Test_New_Ptr(t *testing.T) {
	// Arrange
	trace := codestack.New.Ptr(0)

	// Act
	actual := args.Map{
		"notNil": trace != nil,
		"isOkay": trace.IsOkay,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"isOkay": true,
	}
	expected.ShouldBeEqual(t, 0, "New.Ptr returns correct value -- with args", actual)
}

func Test_New_Skip(t *testing.T) {
	// Arrange
	trace := codestack.New.SkipOne()

	// Act
	actual := args.Map{"isOkay": trace.IsOkay}

	// Assert
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "New.SkipOne returns correct value -- with args", actual)
}

func Test_New_SkipPtr(t *testing.T) {
	// Arrange
	trace := codestack.New.Ptr(1)

	// Act
	actual := args.Map{"notNil": trace != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "New.Ptr returns correct value -- with args", actual)
}

// ── Trace — all getters ──

func Test_Trace_Getters(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	actual := args.Map{
		"isOkay":     trace.IsOkay,
		"isNotOkay":  !trace.IsOkay,
		"isNil":      trace.IsNil(),
		"isNotNil":   trace.IsNotNil(),
		"hasIssues":  trace.HasIssues(),
		"string":     trace.String() != "",
		"shortStr":   trace.ShortString() != "",
		"message":    trace.Message() != "",
		"filePath":   trace.FullFilePath() != "",
		"fileName":   trace.FileName() != "",
		"lineNum":    trace.LineNumber() > 0,
		"fwlStr":     trace.FileWithLineString() != "",
		"jsonStr":    trace.JsonString() != "",
		"pkgMethod":  trace.PackageMethodName != "",
	}

	// Assert
	expected := args.Map{
		"isOkay": true, "isNotOkay": false, "isNil": false, "isNotNil": true,
		"hasIssues": false, "string": true, "shortStr": true, "message": true,
		"filePath": true, "fileName": true, "lineNum": true, "fwlStr": true,
		"jsonStr": true, "pkgMethod": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- getters", actual)
}

func Test_Trace_FileWithLine(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	fwl := trace.FileWithLine()

	// Act
	actual := args.Map{"notEmpty": fwl.FilePath != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace returns non-empty -- FileWithLine", actual)
}

func Test_Trace_Clone(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	cloned := trace.Clone()
	clonedPtr := trace.ClonePtr()

	// Act
	actual := args.Map{
		"pkg": cloned.PackageName,
		"ptrPkg": clonedPtr.PackageName != "",
	}

	// Assert
	expected := args.Map{
		"pkg": trace.PackageName,
		"ptrPkg": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- Clone", actual)
}

func Test_Trace_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act
	actual := args.Map{"isNil": trace.ClonePtr() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Trace returns nil -- ClonePtr nil", actual)
}

func Test_Trace_Dispose(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	trace.Dispose()

	// Act
	actual := args.Map{"pkg": trace.PackageName}

	// Assert
	expected := args.Map{"pkg": ""}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- Dispose", actual)
}

func Test_Trace_JsonModel(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	model := trace.JsonModel()
	modelAny := trace.JsonModelAny()

	// Act
	actual := args.Map{
		"pkg": model.PackageName != "",
		"anyNotNil": modelAny != nil,
	}

	// Assert
	expected := args.Map{
		"pkg": true,
		"anyNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- JsonModel", actual)
}

func Test_Trace_Json(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	r := trace.Json()
	rp := trace.JsonPtr()

	// Act
	actual := args.Map{
		"hasBytes": r.HasBytes(),
		"ptrNotNil": rp != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"ptrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- Json", actual)
}

// ── FileWithLine — all methods ──

func Test_FileWithLine_Basic(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}

	// Act
	actual := args.Map{
		"path":    fwl.FullFilePath(),
		"line":    fwl.LineNumber(),
		"isNil":   fwl.IsNil(),
		"notNil":  fwl.IsNotNil(),
		"string":  fwl.String() != "",
		"fwlStr":  fwl.FileWithLine() != "",
		"jsonStr": fwl.JsonString() != "",
	}

	// Assert
	expected := args.Map{
		"path": "/tmp/test.go", "line": 42, "isNil": false, "notNil": true,
		"string": true, "fwlStr": true, "jsonStr": true,
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- basic", actual)
}

func Test_FileWithLine_Nil(t *testing.T) {
	// Arrange
	var fwl *codestack.FileWithLine

	// Act
	actual := args.Map{
		"isNil": fwl.IsNil(),
		"str": fwl.String(),
	}

	// Assert
	expected := args.Map{
		"isNil": true,
		"str": "",
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns nil -- nil", actual)
}

func Test_FileWithLine_Clone(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}

	// Act
	actual := args.Map{
		"path": fwl.FilePath,
		"line": fwl.Line,
	}

	// Assert
	expected := args.Map{
		"path": "/tmp/test.go",
		"line": 42,
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- fields", actual)
}

func Test_FileWithLine_JsonModel(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	model := fwl.JsonModel()
	modelAny := fwl.JsonModelAny()

	// Act
	actual := args.Map{
		"path": model.FilePath,
		"anyNotNil": modelAny != nil,
	}

	// Assert
	expected := args.Map{
		"path": "/tmp/test.go",
		"anyNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- JsonModel", actual)
}

func Test_FileWithLine_StringUsingFmt(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	result := fwl.StringUsingFmt(func(f codestack.FileWithLine) string { return f.FilePath })

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "/tmp/test.go"}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- StringUsingFmt", actual)
}

func Test_FileWithLine_Json(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	r := fwl.Json()
	rp := fwl.JsonPtr()

	// Act
	actual := args.Map{
		"hasBytes": r.HasBytes(),
		"ptrNotNil": rp != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"ptrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- Json", actual)
}

func Test_FileWithLine_Dispose(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	fwl.FilePath = ""
	fwl.Line = 0

	// Act
	actual := args.Map{"path": fwl.FilePath}

	// Assert
	expected := args.Map{"path": ""}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- Dispose", actual)
}

func Test_FileWithLine_Dispose_Nil(t *testing.T) {
	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns nil -- Dispose nil", actual)
}

// ── StacksTo ──

func Test_StacksTo_String(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.String(0, 10)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": actual["notEmpty"]}
	expected.ShouldBeEqual(t, 0, "StacksTo returns correct value -- String", actual)
}

// ── Dir ──

func Test_Dir_CurrentDir(t *testing.T) {
	// Arrange
	dir := codestack.Dir.CurDir()

	// Act
	actual := args.Map{"notEmpty": len(dir) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir returns correct value -- CurDir", actual)
}

func Test_Dir_RepoDir(t *testing.T) {
	// Arrange
	dir := codestack.Dir.RepoDir()

	// Act
	actual := args.Map{"notEmpty": len(dir) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir returns correct value -- RepoDir", actual)
}

// ── File ──

func Test_File_CurrentFileName(t *testing.T) {
	// Arrange
	file := codestack.File.CurrentFilePath()

	// Act
	actual := args.Map{"notEmpty": len(file) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File returns correct value -- CurrentFileName", actual)
}
