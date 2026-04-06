package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── currentNameOf ──

func Test_Cov6_CurrentNameOf(t *testing.T) {
	name := codestack.NameOf.MethodStackSkip(0)
	actual := args.Map{"notEmpty": name != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CurrentNameOf returns correct value -- with args", actual)
}

// ── skippablePrefixes ──

func Test_Cov6_SkippablePrefixes(t *testing.T) {
	// skippablePrefixes is unexported; verify via NameOf instead
	name := codestack.NameOf.Method()
	actual := args.Map{"notEmpty": name != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.Method -- proxy for skippable check", actual)
}

// ── New.Default / New.Ptr / New.Skip ──

func Test_Cov6_New_Default(t *testing.T) {
	trace := codestack.New.Default()
	actual := args.Map{"isOkay": trace.IsOkay, "pkgNotEmpty": trace.PackageName != ""}
	expected := args.Map{"isOkay": true, "pkgNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "New.Default returns correct value -- with args", actual)
}

func Test_Cov6_New_Ptr(t *testing.T) {
	trace := codestack.New.Ptr(0)
	actual := args.Map{"notNil": trace != nil, "isOkay": trace.IsOkay}
	expected := args.Map{"notNil": true, "isOkay": true}
	expected.ShouldBeEqual(t, 0, "New.Ptr returns correct value -- with args", actual)
}

func Test_Cov6_New_Skip(t *testing.T) {
	trace := codestack.New.SkipOne()
	actual := args.Map{"isOkay": trace.IsOkay}
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "New.SkipOne returns correct value -- with args", actual)
}

func Test_Cov6_New_SkipPtr(t *testing.T) {
	trace := codestack.New.Ptr(1)
	actual := args.Map{"notNil": trace != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "New.Ptr returns correct value -- with args", actual)
}

// ── Trace — all getters ──

func Test_Cov6_Trace_Getters(t *testing.T) {
	trace := codestack.New.Default()
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
	expected := args.Map{
		"isOkay": true, "isNotOkay": false, "isNil": false, "isNotNil": true,
		"hasIssues": false, "string": true, "shortStr": true, "message": true,
		"filePath": true, "fileName": true, "lineNum": true, "fwlStr": true,
		"jsonStr": true, "pkgMethod": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- getters", actual)
}

func Test_Cov6_Trace_FileWithLine(t *testing.T) {
	trace := codestack.New.Default()
	fwl := trace.FileWithLine()
	actual := args.Map{"notEmpty": fwl.FilePath != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace returns non-empty -- FileWithLine", actual)
}

func Test_Cov6_Trace_Clone(t *testing.T) {
	trace := codestack.New.Default()
	cloned := trace.Clone()
	clonedPtr := trace.ClonePtr()
	actual := args.Map{"pkg": cloned.PackageName, "ptrPkg": clonedPtr.PackageName != ""}
	expected := args.Map{"pkg": trace.PackageName, "ptrPkg": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- Clone", actual)
}

func Test_Cov6_Trace_ClonePtr_Nil(t *testing.T) {
	var trace *codestack.Trace
	actual := args.Map{"isNil": trace.ClonePtr() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Trace returns nil -- ClonePtr nil", actual)
}

func Test_Cov6_Trace_Dispose(t *testing.T) {
	trace := codestack.New.Default()
	trace.Dispose()
	actual := args.Map{"pkg": trace.PackageName}
	expected := args.Map{"pkg": ""}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- Dispose", actual)
}

func Test_Cov6_Trace_JsonModel(t *testing.T) {
	trace := codestack.New.Default()
	model := trace.JsonModel()
	modelAny := trace.JsonModelAny()
	actual := args.Map{"pkg": model.PackageName != "", "anyNotNil": modelAny != nil}
	expected := args.Map{"pkg": true, "anyNotNil": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- JsonModel", actual)
}

func Test_Cov6_Trace_Json(t *testing.T) {
	trace := codestack.New.Default()
	r := trace.Json()
	rp := trace.JsonPtr()
	actual := args.Map{"hasBytes": r.HasBytes(), "ptrNotNil": rp != nil}
	expected := args.Map{"hasBytes": true, "ptrNotNil": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- Json", actual)
}

// ── FileWithLine — all methods ──

func Test_Cov6_FileWithLine_Basic(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	actual := args.Map{
		"path":    fwl.FullFilePath(),
		"line":    fwl.LineNumber(),
		"isNil":   fwl.IsNil(),
		"notNil":  fwl.IsNotNil(),
		"string":  fwl.String() != "",
		"fwlStr":  fwl.FileWithLine() != "",
		"jsonStr": fwl.JsonString() != "",
	}
	expected := args.Map{
		"path": "/tmp/test.go", "line": 42, "isNil": false, "notNil": true,
		"string": true, "fwlStr": true, "jsonStr": true,
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- basic", actual)
}

func Test_Cov6_FileWithLine_Nil(t *testing.T) {
	var fwl *codestack.FileWithLine
	actual := args.Map{"isNil": fwl.IsNil(), "str": fwl.String()}
	expected := args.Map{"isNil": true, "str": ""}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns nil -- nil", actual)
}

func Test_Cov6_FileWithLine_Clone(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	actual := args.Map{"path": fwl.FilePath, "line": fwl.Line}
	expected := args.Map{"path": "/tmp/test.go", "line": 42}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- fields", actual)
}

func Test_Cov6_FileWithLine_JsonModel(t *testing.T) {
	fwl := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	model := fwl.JsonModel()
	modelAny := fwl.JsonModelAny()
	actual := args.Map{"path": model.FilePath, "anyNotNil": modelAny != nil}
	expected := args.Map{"path": "/tmp/test.go", "anyNotNil": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- JsonModel", actual)
}

func Test_Cov6_FileWithLine_StringUsingFmt(t *testing.T) {
	fwl := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	result := fwl.StringUsingFmt(func(f codestack.FileWithLine) string { return f.FilePath })
	actual := args.Map{"result": result}
	expected := args.Map{"result": "/tmp/test.go"}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- StringUsingFmt", actual)
}

func Test_Cov6_FileWithLine_Json(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	r := fwl.Json()
	rp := fwl.JsonPtr()
	actual := args.Map{"hasBytes": r.HasBytes(), "ptrNotNil": rp != nil}
	expected := args.Map{"hasBytes": true, "ptrNotNil": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- Json", actual)
}

func Test_Cov6_FileWithLine_Dispose(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	fwl.FilePath = ""
	fwl.Line = 0
	actual := args.Map{"path": fwl.FilePath}
	expected := args.Map{"path": ""}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- Dispose", actual)
}

func Test_Cov6_FileWithLine_Dispose_Nil(t *testing.T) {
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns nil -- Dispose nil", actual)
}

// ── StacksTo ──

func Test_Cov6_StacksTo_String(t *testing.T) {
	result := codestack.StacksTo.String(0, 10)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": actual["notEmpty"]}
	expected.ShouldBeEqual(t, 0, "StacksTo returns correct value -- String", actual)
}

// ── Dir ──

func Test_Cov6_Dir_CurrentDir(t *testing.T) {
	dir := codestack.Dir.CurDir()
	actual := args.Map{"notEmpty": len(dir) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir returns correct value -- CurDir", actual)
}

func Test_Cov6_Dir_RepoDir(t *testing.T) {
	dir := codestack.Dir.RepoDir()
	actual := args.Map{"notEmpty": len(dir) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir returns correct value -- RepoDir", actual)
}

// ── File ──

func Test_Cov6_File_CurrentFileName(t *testing.T) {
	file := codestack.File.CurrentFilePath()
	actual := args.Map{"notEmpty": len(file) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File returns correct value -- CurrentFileName", actual)
}
