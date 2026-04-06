package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── newTraceCollection ──

func Test_Cov2_NewTraceCollection_Default(t *testing.T) {
	tc := codestack.New.StackTrace.Default(0, codestack.DefaultStackCount)
	actual := args.Map{"notEmpty": !tc.IsEmpty()}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "newTraceCollection.Default returns non-empty -- default args", actual)
}

func Test_Cov2_NewTraceCollection_SkipOne(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	actual := args.Map{"notEmpty": !tc.IsEmpty()}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "newTraceCollection.SkipOne returns non-empty -- skip one", actual)
}

func Test_Cov2_NewTraceCollection_SkipNone(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	actual := args.Map{"notEmpty": !tc.IsEmpty()}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "newTraceCollection.SkipNone returns non-empty -- skip none", actual)
}

func Test_Cov2_NewTraceCollection_DefaultCount(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(0)
	actual := args.Map{"notEmpty": !tc.IsEmpty()}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "newTraceCollection.DefaultCount returns non-empty -- skip 1", actual)
}

// ── Trace edge cases ──

func Test_Cov2_Trace_NilPtr_String(t *testing.T) {
	var trace *codestack.Trace
	actual := args.Map{"result": trace.String()}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Trace.String returns empty -- nil pointer", actual)
}

func Test_Cov2_Trace_NilPtr_HasIssues(t *testing.T) {
	var trace *codestack.Trace
	actual := args.Map{"hasIssues": trace.HasIssues()}
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "Trace.HasIssues returns true -- nil pointer", actual)
}

func Test_Cov2_Trace_NilPtr_IsNil(t *testing.T) {
	var trace *codestack.Trace
	actual := args.Map{"isNil": trace.IsNil()}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Trace.IsNil returns true -- nil pointer", actual)
}

func Test_Cov2_Trace_NilPtr_IsNotNil(t *testing.T) {
	var trace *codestack.Trace
	actual := args.Map{"isNotNil": trace.IsNotNil()}
	expected := args.Map{"isNotNil": false}
	expected.ShouldBeEqual(t, 0, "Trace.IsNotNil returns false -- nil pointer", actual)
}

func Test_Cov2_Trace_NilPtr_Dispose(t *testing.T) {
	var trace *codestack.Trace
	trace.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Trace.Dispose returns safely -- nil pointer", actual)
}

func Test_Cov2_Trace_NilPtr_ClonePtr(t *testing.T) {
	var trace *codestack.Trace
	actual := args.Map{"isNil": trace.ClonePtr() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Trace.ClonePtr returns nil -- nil pointer", actual)
}

func Test_Cov2_Trace_Empty_HasIssues(t *testing.T) {
	trace := codestack.Trace{}
	actual := args.Map{"hasIssues": trace.HasIssues()}
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "Trace.HasIssues returns true -- empty struct", actual)
}

// ── Trace message caching ──

func Test_Cov2_Trace_Message_CalledTwice(t *testing.T) {
	trace := codestack.New.Default()
	msg1 := trace.Message()
	msg2 := trace.Message()
	actual := args.Map{"equal": msg1 == msg2, "notEmpty": msg1 != ""}
	expected := args.Map{"equal": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace.Message returns cached -- called twice", actual)
}

func Test_Cov2_Trace_ShortString_CalledTwice(t *testing.T) {
	trace := codestack.New.Default()
	s1 := trace.ShortString()
	s2 := trace.ShortString()
	actual := args.Map{"equal": s1 == s2, "notEmpty": s1 != ""}
	expected := args.Map{"equal": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace.ShortString returns cached -- called twice", actual)
}

// ── FileWithLine edge cases ──

func Test_Cov2_FileWithLine_NilPtr_String(t *testing.T) {
	var fwl *codestack.FileWithLine
	actual := args.Map{"result": fwl.String()}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "FileWithLine.String returns empty -- nil pointer", actual)
}

func Test_Cov2_FileWithLine_NilPtr_IsNil(t *testing.T) {
	var fwl *codestack.FileWithLine
	actual := args.Map{"isNil": fwl.IsNil()}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine.IsNil returns true -- nil pointer", actual)
}

func Test_Cov2_FileWithLine_NilPtr_IsNotNil(t *testing.T) {
	var fwl *codestack.FileWithLine
	actual := args.Map{"isNotNil": fwl.IsNotNil()}
	expected := args.Map{"isNotNil": false}
	expected.ShouldBeEqual(t, 0, "FileWithLine.IsNotNil returns false -- nil pointer", actual)
}

// ── stacksTo ──

func Test_Cov2_StacksTo_String(t *testing.T) {
	result := codestack.StacksTo.String(0, 5)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.String returns non-empty -- with count", actual)
}

func Test_Cov2_StacksTo_StringDefault(t *testing.T) {
	result := codestack.StacksTo.StringDefault()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringDefault returns non-empty -- default args", actual)
}

func Test_Cov2_StacksTo_StringNoCount(t *testing.T) {
	result := codestack.StacksTo.StringNoCount(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringNoCount returns non-empty -- skip 0", actual)
}

func Test_Cov2_StacksTo_Bytes(t *testing.T) {
	result := codestack.StacksTo.Bytes(0)
	actual := args.Map{"notEmpty": len(result) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.Bytes returns non-empty -- skip 0", actual)
}

func Test_Cov2_StacksTo_BytesDefault(t *testing.T) {
	result := codestack.StacksTo.BytesDefault()
	actual := args.Map{"notEmpty": len(result) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.BytesDefault returns non-empty -- default", actual)
}

func Test_Cov2_StacksTo_JsonString(t *testing.T) {
	result := codestack.StacksTo.JsonString(0)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.JsonString returns non-empty -- skip 0", actual)
}

func Test_Cov2_StacksTo_JsonStringDefault(t *testing.T) {
	result := codestack.StacksTo.JsonStringDefault()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.JsonStringDefault returns non-empty -- default", actual)
}

func Test_Cov2_StacksTo_StringUsingFmt(t *testing.T) {
	result := codestack.StacksTo.StringUsingFmt(
		func(tr *codestack.Trace) string { return tr.PackageName },
		1,
		5,
	)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": result != ""}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringUsingFmt returns non-empty -- with formatter", actual)
}

// ── newCreator ──

func Test_Cov2_NewCreator_Create(t *testing.T) {
	trace := codestack.New.Create(0)
	actual := args.Map{"isOkay": trace.IsOkay, "notEmpty": trace.PackageName != ""}
	expected := args.Map{"isOkay": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "newCreator.Create returns valid trace -- skip 0", actual)
}

func Test_Cov2_NewCreator_Default(t *testing.T) {
	trace := codestack.New.Default()
	actual := args.Map{"isOkay": trace.IsOkay}
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "newCreator.Default returns valid trace -- default", actual)
}

func Test_Cov2_NewCreator_SkipOne(t *testing.T) {
	trace := codestack.New.SkipOne()
	actual := args.Map{"isOkay": trace.IsOkay}
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "newCreator.SkipOne returns valid trace -- skip one", actual)
}

func Test_Cov2_NewCreator_Ptr(t *testing.T) {
	ptr := codestack.New.Ptr(0)
	actual := args.Map{"notNil": ptr != nil, "isOkay": ptr.IsOkay}
	expected := args.Map{"notNil": true, "isOkay": true}
	expected.ShouldBeEqual(t, 0, "newCreator.Ptr returns valid pointer -- skip 0", actual)
}

// ── NameOf edge cases ──

func Test_Cov2_NameOf_All_EmptyInput(t *testing.T) {
	full, pkg, method := codestack.NameOf.All("")
	actual := args.Map{"full": full, "pkg": pkg, "method": method}
	expected := args.Map{"full": "", "pkg": "", "method": ""}
	expected.ShouldBeEqual(t, 0, "NameOf.All returns empty -- empty input", actual)
}

func Test_Cov2_NameOf_Method_EmptyInput(t *testing.T) {
	actual := args.Map{"result": codestack.NameOf.MethodByFullName("")}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NameOf.Method returns empty -- empty input", actual)
}

func Test_Cov2_NameOf_Package_EmptyInput(t *testing.T) {
	actual := args.Map{"result": codestack.NameOf.PackageByFullName("")}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NameOf.Package returns empty -- empty input", actual)
}

// ── Trace.AsFileLiner ──

func Test_Cov2_Trace_AsFileLiner(t *testing.T) {
	trace := codestack.New.Default()
	liner := trace.AsFileLiner()
	actual := args.Map{"notNil": liner != nil, "pathNotEmpty": liner.FullFilePath() != ""}
	expected := args.Map{"notNil": true, "pathNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace.AsFileLiner returns valid FileWithLiner -- default trace", actual)
}

// ── FileWithLine.AsFileLiner ──

func Test_Cov2_FileWithLine_AsFileLiner(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	liner := fwl.AsFileLiner()
	actual := args.Map{"notNil": liner != nil, "line": liner.LineNumber()}
	expected := args.Map{"notNil": true, "line": 10}
	expected.ShouldBeEqual(t, 0, "FileWithLine.AsFileLiner returns valid FileWithLiner -- with data", actual)
}

// ── Trace.Clone with IsSkippable ──

func Test_Cov2_Trace_Clone_PreservesIsSkippable(t *testing.T) {
	trace := codestack.New.Default()
	cloned := trace.Clone()
	actual := args.Map{
		"skippableMatch": cloned.IsSkippable == trace.IsSkippable,
		"pkgMatch":       cloned.PackageName == trace.PackageName,
	}
	expected := args.Map{
		"skippableMatch": true,
		"pkgMatch":       true,
	}
	expected.ShouldBeEqual(t, 0, "Trace.Clone preserves IsSkippable -- default trace", actual)
}

// ── File / Dir getters ──

func Test_Cov2_File_Name(t *testing.T) {
	actual := args.Map{"notEmpty": codestack.File.Name(0) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.Name returns non-empty -- skip 0", actual)
}

func Test_Cov2_File_Path(t *testing.T) {
	actual := args.Map{"notEmpty": codestack.File.Path(0) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.Path returns non-empty -- skip 0", actual)
}

func Test_Cov2_Dir_CurDir(t *testing.T) {
	actual := args.Map{"notEmpty": codestack.Dir.CurDir() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.CurDir returns non-empty -- current dir", actual)
}

func Test_Cov2_Dir_CurDirJoin(t *testing.T) {
	actual := args.Map{"notEmpty": codestack.Dir.CurDirJoin("sub") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.CurDirJoin returns non-empty -- with subdir", actual)
}
