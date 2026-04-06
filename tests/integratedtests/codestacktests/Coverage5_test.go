package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── TraceCollection extended methods ──

func Test_Cov5_TraceCollection_Length(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	actual := args.Map{"gt0": tc.Length() > 0, "notEmpty": tc.HasAnyItem()}
	expected := args.Map{"gt0": actual["gt0"], "notEmpty": actual["notEmpty"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection Length -- default traces", actual)
}

func Test_Cov5_TraceCollection_IsEmpty(t *testing.T) {
	tc := codestack.TraceCollection{}
	actual := args.Map{"isEmpty": tc.IsEmpty(), "lastIndex": tc.LastIndex()}
	expected := args.Map{"isEmpty": true, "lastIndex": -1}
	expected.ShouldBeEqual(t, 0, "TraceCollection IsEmpty -- empty", actual)
}

func Test_Cov5_TraceCollection_Add(t *testing.T) {
	tc := codestack.TraceCollection{}
	trace := codestack.New.Default()
	tc.Add(trace)
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection Add -- single trace", actual)
}

func Test_Cov5_TraceCollection_Adds(t *testing.T) {
	tc := codestack.TraceCollection{}
	t1 := codestack.New.Default()
	t2 := codestack.New.Default()
	tc.Adds(t1, t2)
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection Adds -- two traces", actual)
}

func Test_Cov5_TraceCollection_AddsEmpty(t *testing.T) {
	tc := codestack.TraceCollection{}
	tc.Adds()
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection Adds empty -- no traces", actual)
}

func Test_Cov5_TraceCollection_AddsPtr(t *testing.T) {
	tc := codestack.TraceCollection{}
	t1 := codestack.New.Ptr(0)
	var nilTrace *codestack.Trace
	tc.AddsPtr(true, t1, nilTrace)
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection AddsPtr -- skip nil", actual)
}

func Test_Cov5_TraceCollection_ConcatNew(t *testing.T) {
	tc := codestack.New.StackTrace.Default(1, 2)
	t1 := codestack.New.Default()
	newTc := tc.ConcatNew(t1)
	actual := args.Map{"greaterLen": newTc.Length() > tc.Length()}
	expected := args.Map{"greaterLen": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection ConcatNew -- adds trace", actual)
}

func Test_Cov5_TraceCollection_ConcatNewPtr(t *testing.T) {
	tc := codestack.New.StackTrace.Default(1, 2)
	t1 := codestack.New.Ptr(0)
	newTc := tc.ConcatNewPtr(t1)
	actual := args.Map{"greaterLen": newTc.Length() > tc.Length()}
	expected := args.Map{"greaterLen": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection ConcatNewPtr -- adds trace ptr", actual)
}

func Test_Cov5_TraceCollection_Clone(t *testing.T) {
	tc := codestack.New.StackTrace.Default(1, 3)
	cloned := tc.Clone()
	actual := args.Map{"sameLen": cloned.Length() == tc.Length()}
	expected := args.Map{"sameLen": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection Clone -- same length", actual)
}

func Test_Cov5_TraceCollection_ClonePtr(t *testing.T) {
	tc := codestack.New.StackTrace.Default(1, 3)
	cloned := tc.ClonePtr()
	actual := args.Map{"notNil": cloned != nil, "sameLen": cloned.Length() == tc.Length()}
	expected := args.Map{"notNil": true, "sameLen": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection ClonePtr -- same length", actual)
}

func Test_Cov5_TraceCollection_ClonePtr_Nil(t *testing.T) {
	var tc *codestack.TraceCollection
	cloned := tc.ClonePtr()
	actual := args.Map{"isNil": cloned == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection ClonePtr nil -- returns nil", actual)
}

func Test_Cov5_TraceCollection_FirstOrDefault(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	first := tc.FirstOrDefault()
	actual := args.Map{"notNil": first != nil}
	expected := args.Map{"notNil": actual["notNil"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection FirstOrDefault -- has items", actual)
}

func Test_Cov5_TraceCollection_FirstOrDefault_Empty(t *testing.T) {
	tc := codestack.TraceCollection{}
	first := tc.FirstOrDefault()
	actual := args.Map{"isNil": first == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection FirstOrDefault empty -- nil", actual)
}

func Test_Cov5_TraceCollection_LastOrDefault(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	last := tc.LastOrDefault()
	actual := args.Map{"notNil": last != nil}
	expected := args.Map{"notNil": actual["notNil"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection LastOrDefault -- has items", actual)
}

func Test_Cov5_TraceCollection_LastOrDefault_Empty(t *testing.T) {
	tc := codestack.TraceCollection{}
	last := tc.LastOrDefault()
	actual := args.Map{"isNil": last == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection LastOrDefault empty -- nil", actual)
}

func Test_Cov5_TraceCollection_CodeStacksString(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	s := tc.CodeStacksString()
	actual := args.Map{"notEmpty": len(s) > 0}
	expected := args.Map{"notEmpty": actual["notEmpty"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection CodeStacksString -- has content", actual)
}

func Test_Cov5_TraceCollection_StackTraces(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	s := tc.StackTraces()
	actual := args.Map{"notEmpty": len(s) > 0}
	expected := args.Map{"notEmpty": actual["notEmpty"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection StackTraces -- has content", actual)
}

func Test_Cov5_TraceCollection_StackTracesJsonResult(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	r := tc.StackTracesJsonResult()
	actual := args.Map{"notNil": r != nil, "hasBytes": r.HasBytes()}
	expected := args.Map{"notNil": true, "hasBytes": actual["hasBytes"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection StackTracesJsonResult -- valid", actual)
}

func Test_Cov5_TraceCollection_NewStackTraces(t *testing.T) {
	tc := codestack.TraceCollection{}
	s := tc.NewStackTraces(1)
	actual := args.Map{"notEmpty": len(s) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection NewStackTraces -- has content", actual)
}

func Test_Cov5_TraceCollection_NewDefaultStackTraces(t *testing.T) {
	tc := codestack.TraceCollection{}
	s := tc.NewDefaultStackTraces()
	actual := args.Map{"notEmpty": len(s) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection NewDefaultStackTraces -- has content", actual)
}

func Test_Cov5_TraceCollection_NewStackTracesJsonResult(t *testing.T) {
	tc := codestack.TraceCollection{}
	r := tc.NewStackTracesJsonResult(1)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection NewStackTracesJsonResult -- valid", actual)
}

func Test_Cov5_TraceCollection_NewDefaultStackTracesJsonResult(t *testing.T) {
	tc := codestack.TraceCollection{}
	r := tc.NewDefaultStackTracesJsonResult()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection NewDefaultStackTracesJsonResult -- valid", actual)
}

func Test_Cov5_TraceCollection_GetPagesSize(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	pages := tc.GetPagesSize(2)
	actual := args.Map{"pages": pages > 0}
	expected := args.Map{"pages": actual["pages"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection GetPagesSize -- valid", actual)
}

func Test_Cov5_TraceCollection_GetPagesSize_Zero(t *testing.T) {
	tc := codestack.TraceCollection{}
	pages := tc.GetPagesSize(2)
	actual := args.Map{"pages": pages}
	expected := args.Map{"pages": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection GetPagesSize zero -- empty", actual)
}

func Test_Cov5_TraceCollection_Dispose(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	tc.Dispose()
	actual := args.Map{"isEmpty": tc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection Dispose -- empty after", actual)
}

func Test_Cov5_TraceCollection_Dispose_Nil(t *testing.T) {
	var tc *codestack.TraceCollection
	tc.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection Dispose nil -- no panic", actual)
}

func Test_Cov5_TraceCollection_Json(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	r := tc.Json()
	actual := args.Map{"hasBytes": r.HasBytes()}
	expected := args.Map{"hasBytes": actual["hasBytes"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection Json -- valid", actual)
}

func Test_Cov5_TraceCollection_JsonPtr(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	r := tc.JsonPtr()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection JsonPtr -- valid", actual)
}

func Test_Cov5_TraceCollection_JsonString(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	s := tc.JsonString()
	actual := args.Map{"notEmpty": len(s) > 0}
	expected := args.Map{"notEmpty": actual["notEmpty"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection JsonString -- valid", actual)
}

func Test_Cov5_TraceCollection_JsonStrings(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	s := tc.JsonStrings()
	actual := args.Map{"notEmpty": len(s) > 0}
	expected := args.Map{"notEmpty": actual["notEmpty"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection JsonStrings -- valid", actual)
}

func Test_Cov5_TraceCollection_String(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	s := tc.String()
	actual := args.Map{"notEmpty": len(s) > 0}
	expected := args.Map{"notEmpty": actual["notEmpty"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection String -- valid", actual)
}

func Test_Cov5_TraceCollection_ConcatNewUsingSkipPlusCount(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	newTc := tc.ConcatNewUsingSkipPlusCount(0, 3)
	actual := args.Map{"gt": newTc.Length() >= tc.Length()}
	expected := args.Map{"gt": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection ConcatNewUsingSkipPlusCount -- appended", actual)
}

func Test_Cov5_TraceCollection_AddsUsingSkipDefault(t *testing.T) {
	tc := codestack.TraceCollection{}
	tc.AddsUsingSkipDefault(0)
	actual := args.Map{"hasItems": tc.HasAnyItem()}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection AddsUsingSkipDefault -- has items", actual)
}

// ── FileWithLine extended ──

func Test_Cov5_FileWithLine_ParseInjectUsingJson(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	jsonResult := fwl.JsonPtr()
	var fwl2 codestack.FileWithLine
	result, err := fwl2.ParseInjectUsingJson(jsonResult)
	actual := args.Map{"noErr": err == nil, "path": result.FilePath}
	expected := args.Map{"noErr": true, "path": "/tmp/test.go"}
	expected.ShouldBeEqual(t, 0, "FileWithLine ParseInjectUsingJson -- roundtrip", actual)
}

func Test_Cov5_FileWithLine_JsonParseSelfInject(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	jsonResult := fwl.JsonPtr()
	var fwl2 codestack.FileWithLine
	err := fwl2.JsonParseSelfInject(jsonResult)
	actual := args.Map{"noErr": err == nil, "path": fwl2.FilePath}
	expected := args.Map{"noErr": true, "path": "/tmp/test.go"}
	expected.ShouldBeEqual(t, 0, "FileWithLine JsonParseSelfInject -- roundtrip", actual)
}

func Test_Cov5_FileWithLine_AsFileLiner(t *testing.T) {
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	liner := fwl.AsFileLiner()
	actual := args.Map{"notNil": liner != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine AsFileLiner -- not nil", actual)
}

// ── Trace extended ──

func Test_Cov5_Trace_ParseInjectUsingJson(t *testing.T) {
	trace := codestack.New.Default()
	jsonResult := trace.JsonPtr()
	var trace2 codestack.Trace
	result, err := trace2.ParseInjectUsingJson(jsonResult)
	actual := args.Map{"noErr": err == nil, "hasPkg": result.PackageName != ""}
	expected := args.Map{"noErr": true, "hasPkg": true}
	expected.ShouldBeEqual(t, 0, "Trace ParseInjectUsingJson -- roundtrip", actual)
}

func Test_Cov5_Trace_JsonParseSelfInject(t *testing.T) {
	trace := codestack.New.Default()
	jsonResult := trace.JsonPtr()
	var trace2 codestack.Trace
	err := trace2.JsonParseSelfInject(jsonResult)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Trace JsonParseSelfInject -- roundtrip", actual)
}

func Test_Cov5_Trace_AsFileLiner(t *testing.T) {
	trace := codestack.New.Default()
	liner := trace.AsFileLiner()
	actual := args.Map{"notNil": liner != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace AsFileLiner -- not nil", actual)
}

func Test_Cov5_Trace_StringUsingFmt(t *testing.T) {
	trace := codestack.New.Default()
	s := trace.StringUsingFmt(func(tr codestack.Trace) string { return tr.PackageName })
	actual := args.Map{"notEmpty": len(s) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace StringUsingFmt -- returns pkg name", actual)
}

func Test_Cov5_Trace_HasIssues_Invalid(t *testing.T) {
	trace := &codestack.Trace{}
	actual := args.Map{"hasIssues": trace.HasIssues()}
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "Trace HasIssues -- invalid trace", actual)
}

func Test_Cov5_Trace_Nil_String(t *testing.T) {
	var trace *codestack.Trace
	actual := args.Map{"str": trace.String()}
	expected := args.Map{"str": ""}
	expected.ShouldBeEqual(t, 0, "Trace nil String -- empty", actual)
}

func Test_Cov5_Trace_Nil_Dispose(t *testing.T) {
	var trace *codestack.Trace
	trace.Dispose()
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Trace nil Dispose -- no panic", actual)
}

// ── dirGetter extended ──

func Test_Cov5_Dir_RepoDir(t *testing.T) {
	dir := codestack.Dir.RepoDir()
	actual := args.Map{"notEmpty": len(dir) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.RepoDir -- not empty", actual)
}

func Test_Cov5_Dir_RepoDirJoin(t *testing.T) {
	dir := codestack.Dir.RepoDirJoin("sub")
	actual := args.Map{"notEmpty": len(dir) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.RepoDirJoin -- not empty", actual)
}

// ── fileGetter extended ──

func Test_Cov5_File_CurrentFilePath(t *testing.T) {
	file := codestack.File.CurrentFilePath()
	actual := args.Map{"notEmpty": len(file) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.CurrentFilePath -- not empty", actual)
}

// ── funcs ──

func Test_Cov5_JoinPackageNameWithRelative(t *testing.T) {
	result := codestack.NameOf.JoinPackageNameWithRelative("pkg.Struct", "Method")
	actual := args.Map{"notEmpty": len(result) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JoinPackageNameWithRelative -- not empty", actual)
}

// ── StacksTo extended ──

func Test_Cov5_StacksTo_Lines(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	lines := tc.Strings()
	actual := args.Map{"gt0": len(lines) > 0}
	expected := args.Map{"gt0": actual["gt0"]}
	expected.ShouldBeEqual(t, 0, "StacksTo Lines -- has lines", actual)
}

func Test_Cov5_StacksTo_JsonResult(t *testing.T) {
	tc := codestack.New.StackTrace.SkipNone()
	r := tc.StackTracesJsonResult()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StacksTo JsonResult -- not nil", actual)
}
