package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── TraceCollection extended methods ──

func Test_TraceCollection_Length(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()

	// Act
	actual := args.Map{
		"gt0": tc.Length() > 0,
		"notEmpty": tc.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"gt0": actual["gt0"],
		"notEmpty": actual["notEmpty"],
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection Length -- default traces", actual)
}

func Test_TraceCollection_IsEmpty_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}

	// Act
	actual := args.Map{
		"isEmpty": tc.IsEmpty(),
		"lastIndex": tc.LastIndex(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
		"lastIndex": -1,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection IsEmpty -- empty", actual)
}

func Test_TraceCollection_Add_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	trace := codestack.New.Default()
	tc.Add(trace)

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection Add -- single trace", actual)
}

func Test_TraceCollection_Adds_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	t1 := codestack.New.Default()
	t2 := codestack.New.Default()
	tc.Adds(t1, t2)

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection Adds -- two traces", actual)
}

func Test_TraceCollection_AddsEmpty(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.Adds()

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection Adds empty -- no traces", actual)
}

func Test_TraceCollection_AddsPtr_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	t1 := codestack.New.Ptr(0)
	var nilTrace *codestack.Trace
	tc.AddsPtr(true, t1, nilTrace)

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection AddsPtr -- skip nil", actual)
}

func Test_TraceCollection_ConcatNew_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default(1, 2)
	t1 := codestack.New.Default()
	newTc := tc.ConcatNew(t1)

	// Act
	actual := args.Map{"greaterLen": newTc.Length() > tc.Length()}

	// Assert
	expected := args.Map{"greaterLen": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection ConcatNew -- adds trace", actual)
}

func Test_TraceCollection_ConcatNewPtr_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default(1, 2)
	t1 := codestack.New.Ptr(0)
	newTc := tc.ConcatNewPtr(t1)

	// Act
	actual := args.Map{"greaterLen": newTc.Length() > tc.Length()}

	// Assert
	expected := args.Map{"greaterLen": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection ConcatNewPtr -- adds trace ptr", actual)
}

func Test_TraceCollection_Clone_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default(1, 3)
	cloned := tc.Clone()

	// Act
	actual := args.Map{"sameLen": cloned.Length() == tc.Length()}

	// Assert
	expected := args.Map{"sameLen": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection Clone -- same length", actual)
}

func Test_TraceCollection_ClonePtr_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default(1, 3)
	cloned := tc.ClonePtr()

	// Act
	actual := args.Map{
		"notNil": cloned != nil,
		"sameLen": cloned.Length() == tc.Length(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"sameLen": true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection ClonePtr -- same length", actual)
}

func Test_TraceCollection_ClonePtr_Nil_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	var tc *codestack.TraceCollection
	cloned := tc.ClonePtr()

	// Act
	actual := args.Map{"isNil": cloned == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection ClonePtr nil -- returns nil", actual)
}

func Test_TraceCollection_FirstOrDefault(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	first := tc.FirstOrDefault()

	// Act
	actual := args.Map{"notNil": first != nil}

	// Assert
	expected := args.Map{"notNil": actual["notNil"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection FirstOrDefault -- has items", actual)
}

func Test_TraceCollection_FirstOrDefault_Empty_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	first := tc.FirstOrDefault()

	// Act
	actual := args.Map{"isNil": first == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection FirstOrDefault empty -- nil", actual)
}

func Test_TraceCollection_LastOrDefault(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	last := tc.LastOrDefault()

	// Act
	actual := args.Map{"notNil": last != nil}

	// Assert
	expected := args.Map{"notNil": actual["notNil"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection LastOrDefault -- has items", actual)
}

func Test_TraceCollection_LastOrDefault_Empty_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	last := tc.LastOrDefault()

	// Act
	actual := args.Map{"isNil": last == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection LastOrDefault empty -- nil", actual)
}

func Test_TraceCollection_CodeStacksString_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	s := tc.CodeStacksString()

	// Act
	actual := args.Map{"notEmpty": len(s) > 0}

	// Assert
	expected := args.Map{"notEmpty": actual["notEmpty"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection CodeStacksString -- has content", actual)
}

func Test_TraceCollection_StackTraces_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	s := tc.StackTraces()

	// Act
	actual := args.Map{"notEmpty": len(s) > 0}

	// Assert
	expected := args.Map{"notEmpty": actual["notEmpty"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection StackTraces -- has content", actual)
}

func Test_TraceCollection_StackTracesJsonResult_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	r := tc.StackTracesJsonResult()

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"hasBytes": r.HasBytes(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"hasBytes": actual["hasBytes"],
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection StackTracesJsonResult -- valid", actual)
}

func Test_TraceCollection_NewStackTraces_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	s := tc.NewStackTraces(1)

	// Act
	actual := args.Map{"notEmpty": len(s) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection NewStackTraces -- has content", actual)
}

func Test_TraceCollection_NewDefaultStackTraces_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	s := tc.NewDefaultStackTraces()

	// Act
	actual := args.Map{"notEmpty": len(s) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection NewDefaultStackTraces -- has content", actual)
}

func Test_TraceCollection_NewStackTracesJsonResult_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	r := tc.NewStackTracesJsonResult(1)

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection NewStackTracesJsonResult -- valid", actual)
}

func Test_TraceCollection_NewDefaultStackTracesJsonResult_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	r := tc.NewDefaultStackTracesJsonResult()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection NewDefaultStackTracesJsonResult -- valid", actual)
}

func Test_TraceCollection_GetPagesSize_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	pages := tc.GetPagesSize(2)

	// Act
	actual := args.Map{"pages": pages > 0}

	// Assert
	expected := args.Map{"pages": actual["pages"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection GetPagesSize -- valid", actual)
}

func Test_TraceCollection_GetPagesSize_Zero(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	pages := tc.GetPagesSize(2)

	// Act
	actual := args.Map{"pages": pages}

	// Assert
	expected := args.Map{"pages": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection GetPagesSize zero -- empty", actual)
}

func Test_TraceCollection_Dispose_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	tc.Dispose()

	// Act
	actual := args.Map{"isEmpty": tc.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection Dispose -- empty after", actual)
}

func Test_TraceCollection_Dispose_Nil_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	var tc *codestack.TraceCollection
	tc.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection Dispose nil -- no panic", actual)
}

func Test_TraceCollection_Json_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	r := tc.Json()

	// Act
	actual := args.Map{"hasBytes": r.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": actual["hasBytes"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection Json -- valid", actual)
}

func Test_TraceCollection_JsonPtr_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	r := tc.JsonPtr()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection JsonPtr -- valid", actual)
}

func Test_TraceCollection_JsonString_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	s := tc.JsonString()

	// Act
	actual := args.Map{"notEmpty": len(s) > 0}

	// Assert
	expected := args.Map{"notEmpty": actual["notEmpty"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection JsonString -- valid", actual)
}

func Test_TraceCollection_JsonStrings_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	s := tc.JsonStrings()

	// Act
	actual := args.Map{"notEmpty": len(s) > 0}

	// Assert
	expected := args.Map{"notEmpty": actual["notEmpty"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection JsonStrings -- valid", actual)
}

func Test_TraceCollection_String_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	s := tc.String()

	// Act
	actual := args.Map{"notEmpty": len(s) > 0}

	// Assert
	expected := args.Map{"notEmpty": actual["notEmpty"]}
	expected.ShouldBeEqual(t, 0, "TraceCollection String -- valid", actual)
}

func Test_TraceCollection_ConcatNewUsingSkipPlusCount_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	newTc := tc.ConcatNewUsingSkipPlusCount(0, 3)

	// Act
	actual := args.Map{"gt": newTc.Length() >= tc.Length()}

	// Assert
	expected := args.Map{"gt": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection ConcatNewUsingSkipPlusCount -- appended", actual)
}

func Test_TraceCollection_AddsUsingSkipDefault_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.AddsUsingSkipDefault(0)

	// Act
	actual := args.Map{"hasItems": tc.HasAnyItem()}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection AddsUsingSkipDefault -- has items", actual)
}

// ── FileWithLine extended ──

func Test_FileWithLine_ParseInjectUsingJson_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	jsonResult := fwl.JsonPtr()
	var fwl2 codestack.FileWithLine
	result, err := fwl2.ParseInjectUsingJson(jsonResult)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"path": result.FilePath,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"path": "/tmp/test.go",
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine ParseInjectUsingJson -- roundtrip", actual)
}

func Test_FileWithLine_JsonParseSelfInject_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	jsonResult := fwl.JsonPtr()
	var fwl2 codestack.FileWithLine
	err := fwl2.JsonParseSelfInject(jsonResult)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"path": fwl2.FilePath,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"path": "/tmp/test.go",
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine JsonParseSelfInject -- roundtrip", actual)
}

func Test_FileWithLine_AsFileLiner_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	liner := fwl.AsFileLiner()

	// Act
	actual := args.Map{"notNil": liner != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine AsFileLiner -- not nil", actual)
}

// ── Trace extended ──

func Test_Trace_ParseInjectUsingJson_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	jsonResult := trace.JsonPtr()
	var trace2 codestack.Trace
	result, err := trace2.ParseInjectUsingJson(jsonResult)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasPkg": result.PackageName != "",
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasPkg": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace ParseInjectUsingJson -- roundtrip", actual)
}

func Test_Trace_JsonParseSelfInject_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	jsonResult := trace.JsonPtr()
	var trace2 codestack.Trace
	err := trace2.JsonParseSelfInject(jsonResult)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Trace JsonParseSelfInject -- roundtrip", actual)
}

func Test_Trace_AsFileLiner_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	liner := trace.AsFileLiner()

	// Act
	actual := args.Map{"notNil": liner != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace AsFileLiner -- not nil", actual)
}

func Test_Trace_StringUsingFmt_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	s := trace.StringUsingFmt(func(tr codestack.Trace) string { return tr.PackageName })

	// Act
	actual := args.Map{"notEmpty": len(s) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace StringUsingFmt -- returns pkg name", actual)
}

func Test_Trace_HasIssues_Invalid(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{}

	// Act
	actual := args.Map{"hasIssues": trace.HasIssues()}

	// Assert
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "Trace HasIssues -- invalid trace", actual)
}

func Test_Trace_Nil_String(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act
	actual := args.Map{"str": trace.String()}

	// Assert
	expected := args.Map{"str": ""}
	expected.ShouldBeEqual(t, 0, "Trace nil String -- empty", actual)
}

func Test_Trace_Nil_Dispose(t *testing.T) {
	// Arrange
	var trace *codestack.Trace
	trace.Dispose()

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Trace nil Dispose -- no panic", actual)
}

// ── dirGetter extended ──

func Test_Dir_RepoDir_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	dir := codestack.Dir.RepoDir()

	// Act
	actual := args.Map{"notEmpty": len(dir) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.RepoDir -- not empty", actual)
}

func Test_Dir_RepoDirJoin_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	dir := codestack.Dir.RepoDirJoin("sub")

	// Act
	actual := args.Map{"notEmpty": len(dir) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.RepoDirJoin -- not empty", actual)
}

// ── fileGetter extended ──

func Test_File_CurrentFilePath_FromTraceCollectionLengt(t *testing.T) {
	// Arrange
	file := codestack.File.CurrentFilePath()

	// Act
	actual := args.Map{"notEmpty": len(file) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.CurrentFilePath -- not empty", actual)
}

// ── funcs ──

func Test_JoinPackageNameWithRelative(t *testing.T) {
	// Arrange
	result := codestack.NameOf.JoinPackageNameWithRelative("pkg.Struct", "Method")

	// Act
	actual := args.Map{"notEmpty": len(result) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JoinPackageNameWithRelative -- not empty", actual)
}

// ── StacksTo extended ──

func Test_StacksTo_Lines(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	lines := tc.Strings()

	// Act
	actual := args.Map{"gt0": len(lines) > 0}

	// Assert
	expected := args.Map{"gt0": actual["gt0"]}
	expected.ShouldBeEqual(t, 0, "StacksTo Lines -- has lines", actual)
}

func Test_StacksTo_JsonResult(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	r := tc.StackTracesJsonResult()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StacksTo JsonResult -- not nil", actual)
}
