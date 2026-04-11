package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Trace methods ──

func Test_Trace_Message_FromTraceMessage(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	// Act
	msg := trace.Message()
	short := trace.ShortString()
	// Assert
	actual := args.Map{
		"msgNotEmpty": msg != "",
		"shortNotEmpty": short != "",
	}
	expected := args.Map{
		"msgNotEmpty": true,
		"shortNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- Message and ShortString", actual)
}

func Test_Trace_MessageCached(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	// Act — call twice to hit cache
	msg1 := trace.Message()
	msg2 := trace.Message()
	short1 := trace.ShortString()
	short2 := trace.ShortString()
	// Assert
	actual := args.Map{
		"msgSame": msg1 == msg2,
		"shortSame": short1 == short2,
	}
	expected := args.Map{
		"msgSame": true,
		"shortSame": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- cached Message", actual)
}

func Test_Trace_NilChecks(t *testing.T) {
	// Arrange
	var nilTrace *codestack.Trace
	trace := codestack.New.Default()
	// Act & Assert
	actual := args.Map{
		"isNil":       nilTrace.IsNil(),
		"isNotNil":    trace.IsNotNil(),
		"hasIssuesNil": nilTrace.HasIssues(),
		"stringNil":   nilTrace.String(),
	}
	expected := args.Map{
		"isNil": true, "isNotNil": true,
		"hasIssuesNil": true, "stringNil": "",
	}
	expected.ShouldBeEqual(t, 0, "Trace returns nil -- nil checks", actual)
}

func Test_Trace_FileInfo(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	// Act
	fwl := trace.FileWithLine()
	fullPath := trace.FullFilePath()
	fileName := trace.FileName()
	lineNum := trace.LineNumber()
	fwlStr := trace.FileWithLineString()
	// Assert
	actual := args.Map{
		"fwlNotEmpty":     fwl.FilePath != "",
		"fullPathNotEmpty": fullPath != "",
		"fileNameNotEmpty": fileName != "",
		"lineNumGt0":       lineNum > 0,
		"fwlStrNotEmpty":  fwlStr != "",
	}
	expected := args.Map{
		"fwlNotEmpty": true, "fullPathNotEmpty": true,
		"fileNameNotEmpty": true, "lineNumGt0": true, "fwlStrNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- file info", actual)
}

func Test_Trace_StringUsingFmt_FromTraceMessage(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	// Act
	result := trace.StringUsingFmt(func(tr codestack.Trace) string { return tr.PackageName })
	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- StringUsingFmt", actual)
}

func Test_Trace_CloneDispose(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	// Act
	cloned := trace.Clone()
	clonedPtr := trace.ClonePtr()
	// Assert
	actual := args.Map{
		"clonedPkg":    cloned.PackageName == trace.PackageName,
		"clonedPtrNil": clonedPtr != nil,
	}
	expected := args.Map{
		"clonedPkg": true,
		"clonedPtrNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- Clone", actual)
	// Cleanup
	trace.Dispose()
	clonedPtr.Dispose()
}

func Test_Trace_ClonePtr_Nil_FromTraceMessage(t *testing.T) {
	// Arrange
	var nilTrace *codestack.Trace
	// Act
	result := nilTrace.ClonePtr()
	// Assert
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Trace returns nil -- ClonePtr nil", actual)
}

func Test_Trace_Json_FromTraceMessage(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	// Act
	jsonStr := trace.JsonString()
	jsonResult := trace.Json()
	jsonModelAny := trace.JsonModelAny()
	// Assert
	actual := args.Map{
		"jsonNotEmpty": jsonStr != "",
		"jsonValid":    jsonResult.HasError() == false,
		"modelNotNil":  jsonModelAny != nil,
	}
	expected := args.Map{
		"jsonNotEmpty": true, "jsonValid": true, "modelNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- Json", actual)
}

func Test_Trace_AsFileLiner_FromTraceMessage(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	// Act
	fl := trace.AsFileLiner()
	// Assert
	actual := args.Map{"notNil": fl != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- AsFileLiner", actual)
}

// ── FileWithLine ──

func Test_FileWithLine_Methods(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{FilePath: "/test.go", Line: 42}
	// Act & Assert
	actual := args.Map{
		"path":     fwl.FullFilePath(),
		"line":     fwl.LineNumber(),
		"str":      fwl.String(),
		"fwlStr":   fwl.FileWithLine(),
		"isNil":    fwl.IsNil(),
		"isNotNil": fwl.IsNotNil(),
	}
	expected := args.Map{
		"path": "/test.go", "line": 42,
		"str": "/test.go:42", "fwlStr": "/test.go:42",
		"isNil": false, "isNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- methods", actual)
}

func Test_FileWithLine_Nil_FromTraceMessage(t *testing.T) {
	// Arrange
	var fwl *codestack.FileWithLine
	// Act & Assert
	actual := args.Map{
		"str": fwl.String(),
		"isNil": fwl.IsNil(),
		"isNotNil": fwl.IsNotNil(),
	}
	expected := args.Map{
		"str": "",
		"isNil": true,
		"isNotNil": false,
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns nil -- nil", actual)
}

func Test_FileWithLine_StringUsingFmt_FromTraceMessage(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{FilePath: "/test.go", Line: 1}
	// Act
	result := fwl.StringUsingFmt(func(f codestack.FileWithLine) string { return f.FilePath })
	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": "/test.go"}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- StringUsingFmt", actual)
}

func Test_FileWithLine_Json_FromTraceMessage(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{FilePath: "/test.go", Line: 1}
	// Act
	jsonStr := fwl.JsonString()
	jsonModel := fwl.JsonModel()
	jsonModelAny := fwl.JsonModelAny()
	fl := fwl.AsFileLiner()
	// Assert
	actual := args.Map{
		"jsonNotEmpty": jsonStr != "",
		"modelPath":    jsonModel.FilePath,
		"anyNotNil":    jsonModelAny != nil,
		"flNotNil":     fl != nil,
	}
	expected := args.Map{
		"jsonNotEmpty": true, "modelPath": "/test.go",
		"anyNotNil": true, "flNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- Json", actual)
}

// ── TraceCollection basic methods ──

func Test_TraceCollection_Empty_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	// Assert
	actual := args.Map{
		"isEmpty":    tc.IsEmpty(),
		"hasAny":     tc.HasAnyItem(),
		"length":     tc.Length(),
		"count":      tc.Count(),
		"lastIndex":  tc.LastIndex(),
	}
	expected := args.Map{
		"isEmpty": true, "hasAny": false, "length": 0, "count": 0, "lastIndex": -1,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- empty", actual)
}

func Test_TraceCollection_Add_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	trace := codestack.New.Default()
	// Act
	tc.Add(trace)
	// Assert
	actual := args.Map{
		"length": tc.Length(),
		"hasAny": tc.HasAnyItem(),
	}
	expected := args.Map{
		"length": 1,
		"hasAny": true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Add", actual)
}

func Test_TraceCollection_Adds_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	t1 := codestack.New.Default()
	t2 := codestack.New.Default()
	// Act
	tc.Adds(t1, t2)
	// Assert
	actual := args.Map{"length": tc.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Adds", actual)
}

func Test_TraceCollection_Adds_Empty_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	// Act
	tc.Adds()
	// Assert
	actual := args.Map{"length": tc.Length()}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- Adds empty", actual)
}

func Test_TraceCollection_AddsIf(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	t1 := codestack.New.Default()
	// Act
	tc.AddsIf(false, t1)
	lenAfterFalse := tc.Length()
	tc.AddsIf(true, t1)
	lenAfterTrue := tc.Length()
	// Assert
	actual := args.Map{
		"afterFalse": lenAfterFalse,
		"afterTrue": lenAfterTrue,
	}
	expected := args.Map{
		"afterFalse": 0,
		"afterTrue": 1,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- AddsIf", actual)
}

func Test_TraceCollection_FirstLast_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	t1 := codestack.Trace{PackageName: "first"}
	t2 := codestack.Trace{PackageName: "last"}
	tc.Adds(t1, t2)
	// Act & Assert
	actual := args.Map{
		"first":          tc.First().PackageName,
		"last":           tc.Last().PackageName,
		"firstOrDefault": tc.FirstOrDefault() != nil,
		"lastOrDefault":  tc.LastOrDefault() != nil,
	}
	expected := args.Map{
		"first": "first", "last": "last",
		"firstOrDefault": true, "lastOrDefault": true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- First/Last", actual)
}

func Test_TraceCollection_FirstLastOrDefault_Empty(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	// Assert
	actual := args.Map{
		"firstOrDefault": tc.FirstOrDefault() == nil,
		"lastOrDefault":  tc.LastOrDefault() == nil,
	}
	expected := args.Map{
		"firstOrDefault": true,
		"lastOrDefault": true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- FirstLast empty", actual)
}

func Test_TraceCollection_HasIndex_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.Add(codestack.Trace{})
	// Assert
	actual := args.Map{
		"has0": tc.HasIndex(0),
		"has1": tc.HasIndex(1),
	}
	expected := args.Map{
		"has0": true,
		"has1": false,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- HasIndex", actual)
}

func Test_TraceCollection_Strings_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.Add(codestack.Trace{PackageName: "pkg", PackageMethodName: "pkg.Method", FilePath: "/f.go", Line: 1, IsOkay: true})
	// Act
	strs := tc.Strings()
	// Assert
	actual := args.Map{
		"len": len(strs),
		"notEmpty": strs[0] != "",
	}
	expected := args.Map{
		"len": 1,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Strings", actual)
}

func Test_TraceCollection_Nil(t *testing.T) {
	// Arrange
	var tc *codestack.TraceCollection
	// Assert
	actual := args.Map{"length": tc.Length()}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns nil -- nil Length", actual)
}

func Test_TraceCollection_Clone_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.Add(codestack.Trace{PackageName: "pkg"})
	// Act
	cloned := tc.Clone()
	// Assert
	actual := args.Map{
		"len": cloned.Length(),
		"pkg": cloned.First().PackageName,
	}
	expected := args.Map{
		"len": 1,
		"pkg": "pkg",
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Clone", actual)
}

func Test_TraceCollection_ConcatNew_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.Add(codestack.Trace{PackageName: "a"})
	extra := codestack.Trace{PackageName: "b"}
	// Act
	newTc := tc.ConcatNew(extra)
	// Assert
	actual := args.Map{
		"origLen": tc.Length(),
		"newLen": newTc.Length(),
	}
	expected := args.Map{
		"origLen": 1,
		"newLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- ConcatNew", actual)
}

func Test_TraceCollection_Skip_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.Adds(codestack.Trace{PackageName: "a"}, codestack.Trace{PackageName: "b"})
	// Act
	skipped := tc.Skip(1)
	// Assert
	actual := args.Map{
		"len": len(skipped),
		"first": skipped[0].PackageName,
	}
	expected := args.Map{
		"len": 1,
		"first": "b",
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Skip", actual)
}

func Test_TraceCollection_Take_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.Adds(codestack.Trace{PackageName: "a"}, codestack.Trace{PackageName: "b"})
	// Act
	taken := tc.Take(1)
	// Assert
	actual := args.Map{"len": len(taken)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Take", actual)
}

func Test_TraceCollection_Filter_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.Adds(
		codestack.Trace{PackageName: "keep"},
		codestack.Trace{PackageName: "skip"},
	)
	// Act
	filtered := tc.Filter(func(tr *codestack.Trace) (bool, bool) {
		return tr.PackageName == "keep", false
	})
	// Assert
	actual := args.Map{"len": len(filtered)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Filter", actual)
}

func Test_TraceCollection_GetPagesSize_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	for i := 0; i < 15; i++ {
		tc.Add(codestack.Trace{PackageName: "pkg"})
	}
	// Act & Assert
	actual := args.Map{
		"pages10":  tc.GetPagesSize(10),
		"pages0":   tc.GetPagesSize(0),
		"pagesNeg": tc.GetPagesSize(-1),
	}
	expected := args.Map{
		"pages10": 2,
		"pages0": 0,
		"pagesNeg": 0,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- GetPagesSize", actual)
}

// ── currentNameOf ──

func Test_NameOf_All_FromTraceMessage(t *testing.T) {
	// Act
	full, pkg, method := codestack.NameOf.All("github.com/user/repo/pkg.Method")
	emptyFull, emptyPkg, emptyMethod := codestack.NameOf.All("")
	// Assert
	actual := args.Map{
		"full": full, "pkg": pkg, "method": method,
		"emptyFull": emptyFull, "emptyPkg": emptyPkg, "emptyMethod": emptyMethod,
	}
	expected := args.Map{
		"full": "pkg.Method", "pkg": "pkg", "method": "Method",
		"emptyFull": "", "emptyPkg": "", "emptyMethod": "",
	}
	expected.ShouldBeEqual(t, 0, "NameOf.All returns correct value -- with args", actual)
}

func Test_NameOf_Method_FromTraceMessage(t *testing.T) {
	// Act
	method := codestack.NameOf.Method()
	// Assert
	actual := args.Map{"notEmpty": method != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.Method returns correct value -- with args", actual)
}

func Test_NameOf_Package_FromTraceMessage(t *testing.T) {
	// Act
	pkg := codestack.NameOf.Package()
	// Assert
	actual := args.Map{"notEmpty": pkg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.Package returns correct value -- with args", actual)
}

func Test_NameOf_MethodByFullName_FromTraceMessage(t *testing.T) {
	// Act
	result := codestack.NameOf.MethodByFullName("github.com/user/repo/pkg.MyFunc")
	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": "MyFunc"}
	expected.ShouldBeEqual(t, 0, "NameOf.MethodByFullName returns correct value -- with args", actual)
}

func Test_NameOf_PackageByFullName_FromTraceMessage(t *testing.T) {
	// Act
	result := codestack.NameOf.PackageByFullName("github.com/user/repo/pkg.MyFunc")
	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": "pkg"}
	expected.ShouldBeEqual(t, 0, "NameOf.PackageByFullName returns correct value -- with args", actual)
}

func Test_NameOf_JoinPackageNameWithRelative_FromTraceMessage(t *testing.T) {
	// Act
	result := codestack.NameOf.JoinPackageNameWithRelative("github.com/user/repo/pkg.X", "SubPkg.Method")
	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": "pkg.SubPkg.Method"}
	expected.ShouldBeEqual(t, 0, "NameOf.JoinPackageNameWithRelative returns non-empty -- with args", actual)
}

func Test_NameOf_CurrentFuncFullPath_FromTraceMessage(t *testing.T) {
	// Act
	result := codestack.NameOf.CurrentFuncFullPath("github.com/user/repo/pkg.Func")
	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": "pkg.Func"}
	expected.ShouldBeEqual(t, 0, "NameOf.CurrentFuncFullPath returns correct value -- with args", actual)
}

// ── Dir / File ──

func Test_Dir_CurDir_FromTraceMessage(t *testing.T) {
	// Act
	dir := codestack.Dir.CurDir()
	// Assert
	actual := args.Map{"notEmpty": dir != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.CurDir returns correct value -- with args", actual)
}

func Test_Dir_RepoDir_FromTraceMessage(t *testing.T) {
	// Act
	repo := codestack.Dir.RepoDir()
	// Assert
	actual := args.Map{"notEmpty": repo != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.RepoDir returns correct value -- with args", actual)
}

func Test_Dir_RepoDirJoin_FromTraceMessage(t *testing.T) {
	// Act
	joined := codestack.Dir.RepoDirJoin("sub", "path")
	// Assert
	actual := args.Map{"notEmpty": joined != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.RepoDirJoin returns correct value -- with args", actual)
}

func Test_Dir_CurDirJoin_FromTraceMessage(t *testing.T) {
	// Act
	joined := codestack.Dir.CurDirJoin("sub")
	// Assert
	actual := args.Map{"notEmpty": joined != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.CurDirJoin returns correct value -- with args", actual)
}

func Test_File_CurrentFilePath_FromTraceMessage(t *testing.T) {
	// Act
	fp := codestack.File.CurrentFilePath()
	// Assert
	actual := args.Map{"notEmpty": fp != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.CurrentFilePath returns correct value -- with args", actual)
}

func Test_File_Name_FromTraceMessage(t *testing.T) {
	// Act
	name := codestack.File.Name(0)
	// Assert
	actual := args.Map{"notEmpty": name != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.Name returns correct value -- with args", actual)
}

func Test_File_Path_FromTraceMessage(t *testing.T) {
	// Act
	fp := codestack.File.Path(0)
	// Assert
	actual := args.Map{"notEmpty": fp != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.Path returns correct value -- with args", actual)
}

func Test_File_PathLineSep_FromTraceMessage(t *testing.T) {
	// Act
	fp, line := codestack.File.PathLineSep(0)
	// Assert
	actual := args.Map{
		"pathNotEmpty": fp != "",
		"lineGt0": line > 0,
	}
	expected := args.Map{
		"pathNotEmpty": true,
		"lineGt0": true,
	}
	expected.ShouldBeEqual(t, 0, "File.PathLineSep returns correct value -- with args", actual)
}

func Test_File_PathLineSepDefault_FromTraceMessage(t *testing.T) {
	// Act
	fp, line := codestack.File.PathLineSepDefault()
	// Assert
	actual := args.Map{
		"pathNotEmpty": fp != "",
		"lineGt0": line > 0,
	}
	expected := args.Map{
		"pathNotEmpty": true,
		"lineGt0": true,
	}
	expected.ShouldBeEqual(t, 0, "File.PathLineSepDefault returns correct value -- with args", actual)
}

func Test_File_FilePathWithLineString_FromTraceMessage(t *testing.T) {
	// Act
	result := codestack.File.FilePathWithLineString(0)
	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.FilePathWithLineString returns non-empty -- with args", actual)
}

func Test_File_PathLineStringDefault_FromTraceMessage(t *testing.T) {
	// Act
	result := codestack.File.PathLineStringDefault()
	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.PathLineStringDefault returns correct value -- with args", actual)
}

// ── StacksTo ──

func Test_StacksTo_String_FromTraceMessage(t *testing.T) {
	// Act
	result := codestack.StacksTo.String(0, 3)
	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.String returns correct value -- with args", actual)
}

func Test_StacksTo_StringDefault_FromTraceMessage(t *testing.T) {
	// Act
	result := codestack.StacksTo.StringDefault()
	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringDefault returns correct value -- with args", actual)
}

func Test_StacksTo_StringNoCount_FromTraceMessage(t *testing.T) {
	// Act
	result := codestack.StacksTo.StringNoCount(0)
	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringNoCount returns correct value -- with args", actual)
}

func Test_StacksTo_Bytes_FromTraceMessage(t *testing.T) {
	// Act
	result := codestack.StacksTo.Bytes(0)
	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.Bytes returns correct value -- with args", actual)
}

func Test_StacksTo_BytesDefault_FromTraceMessage(t *testing.T) {
	// Act
	result := codestack.StacksTo.BytesDefault()
	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.BytesDefault returns correct value -- with args", actual)
}

func Test_StacksTo_JsonString_FromTraceMessage(t *testing.T) {
	// Act
	result := codestack.StacksTo.JsonString(0)
	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.JsonString returns correct value -- with args", actual)
}

func Test_StacksTo_JsonStringDefault_FromTraceMessage(t *testing.T) {
	// Act
	result := codestack.StacksTo.JsonStringDefault()
	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.JsonStringDefault returns correct value -- with args", actual)
}

// ── New.StackTrace ──

func Test_NewStackTrace_SkipNone(t *testing.T) {
	// Act
	tc := codestack.New.StackTrace.SkipNone()
	// Assert
	actual := args.Map{"hasItems": tc.HasAnyItem()}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "New.StackTrace.SkipNone returns correct value -- with args", actual)
}

func Test_NewStackTrace_SkipOne(t *testing.T) {
	// Act
	tc := codestack.New.StackTrace.SkipOne()
	// Assert
	actual := args.Map{"hasItems": tc.HasAnyItem()}
	expected := args.Map{"hasItems": actual["hasItems"]}
	expected.ShouldBeEqual(t, 0, "New.StackTrace.SkipOne returns correct value -- with args", actual)
}

func Test_NewStackTrace_DefaultCount(t *testing.T) {
	// Act
	tc := codestack.New.StackTrace.DefaultCount(0)
	// Assert
	actual := args.Map{"hasItems": tc.HasAnyItem()}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "New.StackTrace.DefaultCount returns correct value -- with args", actual)
}

// ── newTraceCollection ──

func Test_TraceCollection_SkipCollection_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.Adds(codestack.Trace{PackageName: "a"}, codestack.Trace{PackageName: "b"})
	// Act
	skipped := tc.SkipCollection(1)
	// Assert
	actual := args.Map{"len": skipped.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- SkipCollection", actual)
}

func Test_TraceCollection_TakeCollection_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.Adds(codestack.Trace{PackageName: "a"}, codestack.Trace{PackageName: "b"})
	// Act
	taken := tc.TakeCollection(1)
	// Assert
	actual := args.Map{"len": taken.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- TakeCollection", actual)
}

func Test_TraceCollection_LimitCollection_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.Adds(codestack.Trace{PackageName: "a"}, codestack.Trace{PackageName: "b"})
	// Act
	limited := tc.LimitCollection(1)
	// Assert
	actual := args.Map{"len": limited.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- LimitCollection", actual)
}

func Test_TraceCollection_SafeLimitCollection_FromTraceMessage(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.Add(codestack.Trace{PackageName: "a"})
	// Act — limit larger than length
	limited := tc.SafeLimitCollection(100)
	// Assert
	actual := args.Map{"len": limited.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- SafeLimitCollection", actual)
}

func Test_TraceCollection_Dynamics(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.Adds(codestack.Trace{PackageName: "a"}, codestack.Trace{PackageName: "b"})
	// Act & Assert
	actual := args.Map{
		"firstDynNotNil":    tc.FirstDynamic() != nil,
		"lastDynNotNil":     tc.LastDynamic() != nil,
		"firstOrDefDynNil":  tc.FirstOrDefaultDynamic() != nil,
		"lastOrDefDynNil":   tc.LastOrDefaultDynamic() != nil,
		"skipDynNotNil":     tc.SkipDynamic(1) != nil,
		"takeDynNotNil":     tc.TakeDynamic(1) != nil,
		"limitDynNotNil":    tc.LimitDynamic(1) != nil,
		"limitNotNil":       tc.Limit(1) != nil,
	}
	expected := args.Map{
		"firstDynNotNil": true, "lastDynNotNil": true,
		"firstOrDefDynNil": true, "lastOrDefDynNil": true,
		"skipDynNotNil": true, "takeDynNotNil": true,
		"limitDynNotNil": true, "limitNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- dynamics", actual)
}

func Test_TraceCollection_SkipDynamic_AllSkipped(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.Add(codestack.Trace{})
	// Act
	result := tc.SkipDynamic(5)
	// Assert
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- SkipDynamic all skipped", actual)
}
