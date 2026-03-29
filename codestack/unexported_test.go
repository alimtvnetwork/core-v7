package codestack

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ══════════════════════════════════════════════════════════════════════════════
// Consolidated internal tests for unexported symbols in codestack.
// These tests MUST remain in the source package because they access
// unexported fields/functions: New.traces, isSkippablePackage,
// stackTraceEnhance, DefaultStackCount.
//
// Source: Coverage03_TraceCollection_test.go, Coverage04_Creators_test.go,
//         Coverage05_Iteration15_test.go
// ══════════════════════════════════════════════════════════════════════════════

// ── TraceCollection via New.traces (unexported field) ──

func TestInternal_TraceCollection_BasicOps(t *testing.T) {
	tc := New.traces.Empty()
	if !tc.IsEmpty() {
		t.Fatal("should be empty")
	}
	if tc.HasAnyItem() {
		t.Fatal("should not have items")
	}
	if tc.Length() != 0 {
		t.Fatal("expected 0")
	}
	if tc.Count() != 0 {
		t.Fatal("expected 0")
	}
	if tc.LastIndex() != -1 {
		t.Fatal("expected -1")
	}
}

func TestInternal_TraceCollection_Add(t *testing.T) {
	tc := New.traces.Empty()
	trace := New.Create(0)
	tc.Add(trace)
	if tc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestInternal_TraceCollection_Adds(t *testing.T) {
	tc := New.traces.Empty()
	t1 := New.Create(0)
	t2 := New.Create(0)
	tc.Adds(t1, t2)
	if tc.Length() != 2 {
		t.Fatal("expected 2")
	}
	tc.Adds()
	if tc.Length() != 2 {
		t.Fatal("expected 2 still")
	}
}

func TestInternal_TraceCollection_AddsPtr(t *testing.T) {
	tc := New.traces.Empty()
	p := New.Ptr(0)
	tc.AddsPtr(false, p, nil)
	if tc.Length() != 1 {
		t.Fatal("expected 1")
	}

	tc2 := New.traces.Empty()
	tc2.AddsPtr(true, p, nil)
	if tc2.Length() != 1 {
		t.Fatal("expected 1")
	}

	tc3 := New.traces.Empty()
	tc3.AddsPtr(false)
	if tc3.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func TestInternal_TraceCollection_AddsIf(t *testing.T) {
	tc := New.traces.Empty()
	tr := New.Create(0)
	tc.AddsIf(true, tr)
	if tc.Length() != 1 {
		t.Fatal("expected 1")
	}
	tc.AddsIf(false, tr)
	if tc.Length() != 1 {
		t.Fatal("expected 1 still")
	}
}

func TestInternal_TraceCollection_FirstLast(t *testing.T) {
	tc := New.traces.Empty()
	if tc.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
	if tc.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}
	_ = tc.FirstOrDefaultDynamic()
	_ = tc.LastOrDefaultDynamic()

	tr := New.Create(0)
	tc.Add(tr)
	tc.Add(New.Create(0))
	_ = tc.First()
	_ = tc.Last()
	_ = tc.FirstDynamic()
	_ = tc.LastDynamic()
	if tc.FirstOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
	if tc.LastOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
}

func TestInternal_TraceCollection_SkipTakeLimit(t *testing.T) {
	tc := New.traces.Empty()
	for i := 0; i < 5; i++ {
		tc.Add(New.Create(0))
	}
	s := tc.Skip(2)
	if len(s) != 3 {
		t.Fatal("expected 3")
	}
	_ = tc.SkipDynamic(2)
	_ = tc.SkipDynamic(10)
	_ = tc.SkipCollection(1)
	tk := tc.Take(3)
	if len(tk) != 3 {
		t.Fatal("expected 3")
	}
	_ = tc.TakeDynamic(2)
	_ = tc.TakeCollection(2)
	_ = tc.Limit(3)
	_ = tc.LimitCollection(3)
	_ = tc.LimitDynamic(2)
	_ = tc.SafeLimitCollection(3)
	_ = tc.SafeLimitCollection(100)
}

func TestInternal_TraceCollection_HasIndex(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	if !tc.HasIndex(0) {
		t.Fatal("should have index 0")
	}
	if tc.HasIndex(1) {
		t.Fatal("should not have index 1")
	}
}

func TestInternal_TraceCollection_Strings(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	if len(tc.Strings()) != 1 {
		t.Fatal("expected 1")
	}
}

func TestInternal_TraceCollection_Filter(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	tc.Add(New.Create(0))
	filtered := tc.Filter(func(tr *Trace) (bool, bool) {
		return true, false
	})
	if len(filtered) != 2 {
		t.Fatal("expected 2")
	}
	filtered2 := tc.Filter(func(tr *Trace) (bool, bool) {
		return true, true
	})
	if len(filtered2) != 1 {
		t.Fatal("expected 1")
	}
}

func TestInternal_TraceCollection_FilterWithLimit(t *testing.T) {
	tc := New.traces.Empty()
	for i := 0; i < 10; i++ {
		tc.Add(New.Create(0))
	}
	filtered := tc.FilterWithLimit(3, func(tr *Trace) (bool, bool) {
		return true, false
	})
	if len(filtered) != 3 {
		t.Fatal("expected 3")
	}
}

func TestInternal_TraceCollection_FilterTraceCollection(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	fc := tc.FilterTraceCollection(func(tr *Trace) (bool, bool) {
		return true, false
	})
	if fc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestInternal_TraceCollection_FilterByNames(t *testing.T) {
	tc := New.traces.Empty()
	tr := New.Create(0)
	tc.Add(tr)
	_ = tc.FilterPackageNameTraceCollection(tr.PackageName)
	_ = tc.SkipFilterPackageNameTraceCollection("nonexistent")
	_ = tc.FilterMethodNameTraceCollection(tr.MethodName)
	_ = tc.SkipFilterMethodNameTraceCollection("nonexistent")
	_ = tc.FilterFullMethodNameTraceCollection(tr.PackageMethodName)
	_ = tc.SkipFilterFullMethodNameTraceCollection("nonexistent")
	_ = tc.SkipFilterFilenameTraceCollection("nonexistent")
}

func TestInternal_TraceCollection_FileWithLines(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	if len(tc.FileWithLines()) != 1 {
		t.Fatal("expected 1")
	}
	if len(tc.FileWithLinesStrings()) != 1 {
		t.Fatal("expected 1")
	}
}

func TestInternal_TraceCollection_StringsUsingFmt(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	s := tc.StringsUsingFmt(func(tr *Trace) string {
		return tr.PackageName
	})
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func TestInternal_TraceCollection_JoinUsingFmt(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	s := tc.JoinUsingFmt(func(tr *Trace) string {
		return tr.PackageName
	}, ",")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_TraceCollection_ShortStrings(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	if len(tc.ShortStrings()) != 1 {
		t.Fatal("expected 1")
	}
}

func TestInternal_TraceCollection_JoinShortStrings(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	if tc.JoinShortStrings(",") == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_TraceCollection_Reverse(t *testing.T) {
	tc := New.traces.Empty()
	tc.Reverse()
	tc.Add(New.Create(0))
	tc.Reverse()
	tc.Add(New.Create(0))
	tc.Reverse()
	tc.Add(New.Create(0))
	tc.Reverse()
}

func TestInternal_TraceCollection_JsonStrings(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	if len(tc.JsonStrings()) != 1 {
		t.Fatal("expected 1")
	}
}

func TestInternal_TraceCollection_Joins(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	_ = tc.JoinFileWithLinesStrings(",")
	_ = tc.JoinJsonStrings(",")
	_ = tc.Join(",")
	_ = tc.JoinLines()
	_ = tc.JoinCsv()
	_ = tc.JoinCsvLine()
}

func TestInternal_TraceCollection_CodeStacksString(t *testing.T) {
	tc := New.traces.Empty()
	if tc.CodeStacksString() != "" {
		t.Fatal("expected empty")
	}
	tc.Add(New.Create(0))
	if tc.CodeStacksString() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_TraceCollection_FileWithLinesString(t *testing.T) {
	tc := New.traces.Empty()
	if tc.FileWithLinesString() != "" {
		t.Fatal("expected empty")
	}
	tc.Add(New.Create(0))
	if tc.FileWithLinesString() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_TraceCollection_CodeStacksStringLimit(t *testing.T) {
	tc := New.traces.Empty()
	if tc.CodeStacksStringLimit(5) != "" {
		t.Fatal("expected empty")
	}
	tc.Add(New.Create(0))
	if tc.CodeStacksStringLimit(5) == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_TraceCollection_IsEqual(t *testing.T) {
	tc1 := New.traces.Empty()
	tc2 := New.traces.Empty()
	if !tc1.IsEqual(tc2) {
		t.Fatal("should be equal")
	}
	var nilTC *TraceCollection
	if !nilTC.IsEqual(nil) {
		t.Fatal("both nil should be equal")
	}
	if nilTC.IsEqual(tc1) {
		t.Fatal("nil vs non-nil should not be equal")
	}
	if tc1.IsEqual(nil) {
		t.Fatal("non-nil vs nil should not be equal")
	}
}

func TestInternal_TraceCollection_IsEqualItems(t *testing.T) {
	tc := New.traces.Empty()
	tr := New.Create(0)
	tc.Add(tr)
	if !tc.IsEqualItems(tr) {
		t.Fatal("should be equal")
	}
	var nilTC *TraceCollection
	if !nilTC.IsEqualItems() {
		t.Fatal("nil vs nil should be equal")
	}
}

func TestInternal_TraceCollection_ClearDispose(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	tc.Clear()

	tc2 := New.traces.Empty()
	tc2.Add(New.Create(0))
	tc2.Dispose()

	var nilTC *TraceCollection
	nilTC.Dispose()
}

func TestInternal_TraceCollection_Clone(t *testing.T) {
	tc := TraceCollection{Items: []Trace{New.Create(0)}}
	cl := tc.Clone()
	if cl.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestInternal_TraceCollection_ClonePtr(t *testing.T) {
	tc := &TraceCollection{Items: []Trace{New.Create(0)}}
	cp := tc.ClonePtr()
	if cp == nil {
		t.Fatal("expected non-nil")
	}
	var nilTC *TraceCollection
	if nilTC.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
}

func TestInternal_TraceCollection_Paging(t *testing.T) {
	tc := New.traces.Empty()
	for i := 0; i < 25; i++ {
		tc.Add(New.Create(0))
	}
	if tc.GetPagesSize(10) != 3 {
		t.Fatal("expected 3")
	}
	if tc.GetPagesSize(0) != 0 {
		t.Fatal("expected 0")
	}
	paged := tc.GetPagedCollection(10)
	if len(paged) != 3 {
		t.Fatal("expected 3")
	}
}

func TestInternal_TraceCollection_ConcatNew(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	tc2 := tc.ConcatNew(New.Create(0))
	if tc2.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestInternal_TraceCollection_ConcatNewPtr(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	tc2 := tc.ConcatNewPtr(New.Ptr(0), nil)
	if tc2.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestInternal_TraceCollection_NilLength(t *testing.T) {
	var nilTC *TraceCollection
	if nilTC.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func TestInternal_TraceCollection_Json(t *testing.T) {
	tc := TraceCollection{Items: []Trace{New.Create(0)}}
	j := tc.Json()
	if j.HasError() {
		t.Fatal(j.Error)
	}
	jp := tc.JsonPtr()
	if jp == nil {
		t.Fatal("expected non-nil")
	}
}

func TestInternal_TraceCollection_JsonModel(t *testing.T) {
	tc := TraceCollection{Items: []Trace{New.Create(0)}}
	_ = tc.JsonModel()
	_ = tc.JsonModelAny()
}

func TestInternal_TraceCollection_Serializer(t *testing.T) {
	tc := TraceCollection{Items: []Trace{New.Create(0)}}
	b, err := tc.Serializer()
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func TestInternal_TraceCollection_StackTracesBytes(t *testing.T) {
	tc := TraceCollection{}
	if len(tc.StackTracesBytes()) != 0 {
		t.Fatal("expected empty")
	}
	tc.Items = append(tc.Items, New.Create(0))
	if len(tc.StackTracesBytes()) == 0 {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_TraceCollection_ParseInjectUsingJson(t *testing.T) {
	tc := TraceCollection{Items: []Trace{New.Create(0)}}
	jr := corejson.NewPtr(tc)
	target := New.traces.Empty()
	_, err := target.ParseInjectUsingJson(jr)
	if err != nil {
		t.Fatal(err)
	}
	badJr := corejson.NewResult.UsingBytes([]byte("invalid"))
	_, err2 := target.ParseInjectUsingJson(badJr.Ptr())
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func TestInternal_TraceCollection_ParseInjectUsingJsonMust(t *testing.T) {
	tc := TraceCollection{Items: []Trace{New.Create(0)}}
	jr := corejson.NewPtr(tc)
	target := New.traces.Empty()
	result := target.ParseInjectUsingJsonMust(jr)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func TestInternal_TraceCollection_JsonParseSelfInject(t *testing.T) {
	tc := TraceCollection{Items: []Trace{New.Create(0)}}
	jr := corejson.NewPtr(tc)
	target := New.traces.Empty()
	err := target.JsonParseSelfInject(jr)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInternal_TraceCollection_AsInterfaces(t *testing.T) {
	tc := TraceCollection{}
	_ = tc.AsJsonContractsBinder()
	_ = tc.AsJsoner()
	_ = tc.AsJsonParseSelfInjector()
}

func TestInternal_TraceCollection_JsonString(t *testing.T) {
	tc := TraceCollection{}
	if tc.JsonString() != "" {
		t.Fatal("expected empty")
	}
	tc.Items = append(tc.Items, New.Create(0))
	if tc.JsonString() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_TraceCollection_String(t *testing.T) {
	tc := TraceCollection{}
	if tc.String() != "" {
		t.Fatal("expected empty")
	}
	tc.Items = append(tc.Items, New.Create(0))
	if tc.String() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_TraceCollection_CsvStrings(t *testing.T) {
	tc := TraceCollection{}
	if len(tc.CsvStrings()) != 0 {
		t.Fatal("expected empty")
	}
	tc.Items = append(tc.Items, New.Create(0))
	if len(tc.CsvStrings()) != 1 {
		t.Fatal("expected 1")
	}
}

func TestInternal_TraceCollection_StackTraces(t *testing.T) {
	tc := TraceCollection{Items: []Trace{New.Create(0)}}
	if tc.StackTraces() == "" {
		t.Fatal("expected non-empty")
	}
	if tc.StackTracesJsonResult() == nil {
		t.Fatal("expected non-nil")
	}
}

func TestInternal_TraceCollection_NewStackTraces(t *testing.T) {
	tc := TraceCollection{}
	_ = tc.NewStackTraces(0)
	_ = tc.NewDefaultStackTraces()
	_ = tc.NewStackTracesJsonResult(0)
	_ = tc.NewDefaultStackTracesJsonResult()
}

// ── Creators: NameOf, StacksTo, File, Dir (Coverage04) ──

func TestInternal_NameOf_Method(t *testing.T) {
	if NameOf.Method() == "" {
		t.Fatal("expected method name")
	}
}

func TestInternal_NameOf_Package(t *testing.T) {
	if NameOf.Package() == "" {
		t.Fatal("expected package name")
	}
}

func TestInternal_NameOf_All(t *testing.T) {
	full, pkg, method := NameOf.All("github.com/alimtvnetwork/core/codestack.TestInternal_NameOf_All")
	if full == "" || pkg == "" || method == "" {
		t.Fatal("expected non-empty")
	}
	f2, p2, m2 := NameOf.All("")
	if f2 != "" || p2 != "" || m2 != "" {
		t.Fatal("expected all empty")
	}
}

func TestInternal_NameOf_AllStackSkip(t *testing.T) {
	full, pkg, method := NameOf.AllStackSkip(0)
	if full == "" || pkg == "" || method == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_NameOf_MethodStackSkip(t *testing.T) {
	if NameOf.MethodStackSkip(0) == "" {
		t.Fatal("expected method name")
	}
}

func TestInternal_NameOf_PackageStackSkip(t *testing.T) {
	if NameOf.PackageStackSkip(0) == "" {
		t.Fatal("expected package name")
	}
}

func TestInternal_NameOf_MethodByFullName(t *testing.T) {
	_ = NameOf.MethodByFullName("github.com/alimtvnetwork/core/codestack.TestMethod")
}

func TestInternal_NameOf_PackageByFullName(t *testing.T) {
	_ = NameOf.PackageByFullName("github.com/alimtvnetwork/core/codestack.TestMethod")
}

func TestInternal_NameOf_CurrentFuncFullPath(t *testing.T) {
	_ = NameOf.CurrentFuncFullPath("github.com/alimtvnetwork/core/codestack.TestMethod")
}

func TestInternal_NameOf_JoinPackageNameWithRelative(t *testing.T) {
	s := NameOf.JoinPackageNameWithRelative(
		"github.com/alimtvnetwork/core/codestack.TestMethod",
		"NewFunc",
	)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_StacksCreator_All(t *testing.T) {
	_ = New.StackTrace.All(true, true, 0, 5)
}

func TestInternal_StacksCreator_Default(t *testing.T) {
	_ = New.StackTrace.Default(0, 5)
}

func TestInternal_StacksCreator_DefaultCount(t *testing.T) {
	_ = New.StackTrace.DefaultCount(0)
}

func TestInternal_StacksCreator_SkipOne(t *testing.T) {
	_ = New.StackTrace.SkipOne()
}

func TestInternal_StacksCreator_SkipNone(t *testing.T) {
	_ = New.StackTrace.SkipNone()
}

func TestInternal_StacksTo_String(t *testing.T) {
	if StacksTo.String(0, 3) == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_StacksTo_StringDefault(t *testing.T) {
	_ = StacksTo.StringDefault()
}

func TestInternal_StacksTo_StringNoCount(t *testing.T) {
	_ = StacksTo.StringNoCount(0)
}

func TestInternal_StacksTo_Bytes(t *testing.T) {
	_ = StacksTo.Bytes(0)
}

func TestInternal_StacksTo_BytesDefault(t *testing.T) {
	_ = StacksTo.BytesDefault()
}

func TestInternal_StacksTo_JsonString(t *testing.T) {
	_ = StacksTo.JsonString(0)
}

func TestInternal_StacksTo_JsonStringDefault(t *testing.T) {
	_ = StacksTo.JsonStringDefault()
}

func TestInternal_StacksTo_StringUsingFmt(t *testing.T) {
	_ = StacksTo.StringUsingFmt(func(tr *Trace) string {
		return tr.PackageName
	}, 0, 3)
}

func TestInternal_File_Name(t *testing.T) {
	if File.Name(0) == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_File_Path(t *testing.T) {
	if File.Path(0) == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_File_PathLineSep(t *testing.T) {
	fp, ln := File.PathLineSep(0)
	if fp == "" || ln == 0 {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_File_PathLineSepDefault(t *testing.T) {
	fp, ln := File.PathLineSepDefault()
	if fp == "" || ln == 0 {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_File_FilePathWithLineString(t *testing.T) {
	if File.FilePathWithLineString(0) == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_File_PathLineStringDefault(t *testing.T) {
	if File.PathLineStringDefault() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_File_CurrentFilePath(t *testing.T) {
	if File.CurrentFilePath() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_Dir_CurDir(t *testing.T) {
	if Dir.CurDir() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_Dir_CurDirJoin(t *testing.T) {
	if Dir.CurDirJoin("subdir") == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_Dir_RepoDir(t *testing.T) {
	_ = Dir.RepoDir()
}

func TestInternal_Dir_RepoDirJoin(t *testing.T) {
	_ = Dir.RepoDirJoin("subdir")
}

func TestInternal_Dir_Get(t *testing.T) {
	if Dir.Get(0) == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_NewTraceCollection_Default(t *testing.T) {
	tc := New.traces.Default()
	if tc == nil {
		t.Fatal("expected non-nil")
	}
}

func TestInternal_NewTraceCollection_Using(t *testing.T) {
	tc := New.traces.Using(false, New.Create(0))
	if tc.Length() != 1 {
		t.Fatal("expected 1")
	}
	tc2 := New.traces.Using(true, New.Create(0))
	if tc2.Length() != 1 {
		t.Fatal("expected 1")
	}
	tc3 := New.traces.Using(false)
	if tc3.Length() != 0 {
		t.Fatal("expected 0")
	}
}

// ── isSkippablePackage (unexported function) ──

func TestInternal_SkippablePrefixes(t *testing.T) {
	if !isSkippablePackage("runtime") {
		t.Fatal("runtime should be skippable")
	}
	if isSkippablePackage("mypackage") {
		t.Fatal("mypackage should not be skippable")
	}
}

func TestInternal_TraceCollection_AddsUsingSkipDefault(t *testing.T) {
	tc := New.traces.Empty()
	tc.AddsUsingSkipDefault(0)
}

func TestInternal_TraceCollection_AddsUsingSkipUsingFilter(t *testing.T) {
	tc := New.traces.Empty()
	tc.AddsUsingSkipUsingFilter(true, true, 0, 5, func(tr *Trace) (bool, bool) {
		return true, false
	})
}

func TestInternal_TraceCollection_ConcatNewUsingSkipPlusCount(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	_ = tc.ConcatNewUsingSkipPlusCount(0, 3)
}

func TestInternal_TraceCollection_ConcatNewUsingSkip(t *testing.T) {
	tc := New.traces.Empty()
	tc.Add(New.Create(0))
	_ = tc.ConcatNewUsingSkip(0)
}

// ── Iteration15: Panic/error branches (Coverage05) ──

func TestInternal_I15_FileWithLine_ParseInjectUsingJsonMust_Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	badJson := corejson.NewPtr("{invalid")
	fw := &FileWithLine{}
	fw.ParseInjectUsingJsonMust(badJson)
}

func TestInternal_I15_Trace_ParseInjectUsingJson_Error(t *testing.T) {
	badJson := corejson.NewPtr("{invalid")
	tr := &Trace{}
	_, err := tr.ParseInjectUsingJson(badJson)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestInternal_I15_Trace_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	badJson := corejson.NewPtr("{invalid")
	tr := &Trace{}
	tr.ParseInjectUsingJsonMust(badJson)
}

func TestInternal_I15_TraceCollection_AddsUsingSkip_BreakOnceInvalid(t *testing.T) {
	tc := New.traces.Default()
	tc.AddsUsingSkip(true, true, 9999, 5)
}

func TestInternal_I15_TraceCollection_AddsUsingSkipUsingFilter_BreakOnFilter(t *testing.T) {
	tc := New.traces.Default()
	count := 0
	tc.AddsUsingSkipUsingFilter(true, false, 0, DefaultStackCount, func(tr *Trace) (bool, bool) {
		count++
		if count >= 2 {
			return true, true
		}
		return true, false
	})
}

func TestInternal_I15_TraceCollection_AddsUsingSkipUsingFilter_SkipInvalid(t *testing.T) {
	tc := New.traces.Default()
	tc.AddsUsingSkipUsingFilter(true, true, 9999, 5, func(tr *Trace) (bool, bool) {
		return true, false
	})
}

func TestInternal_I15_TraceCollection_PagedItems_NegativeIndex(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for negative page index")
		}
	}()
	tc := New.traces.Default()
	tc.Adds(New.Create(0))
	tc.GetSinglePageCollection(10, -1)
}

func TestInternal_I15_TraceCollection_FilterWithLimit_BreakBranch(t *testing.T) {
	tc := New.traces.Default()
	tc.Adds(New.Create(0), New.Create(0))
	result := tc.FilterWithLimit(10, func(tr *Trace) (bool, bool) {
		return true, true
	})
	if len(result) != 1 {
		t.Fatal("expected 1")
	}
}

func TestInternal_I15_TraceCollection_IsEqualItems_NilIt(t *testing.T) {
	var tc *TraceCollection
	if !tc.IsEqualItems() {
		t.Fatal("expected nil == nil to be true")
	}
}

func TestInternal_I15_TraceCollection_IsEqualItems_NilVsNonNil(t *testing.T) {
	var tc *TraceCollection
	tr := New.Create(0)
	if tc.IsEqualItems(tr) {
		t.Fatal("expected nil != non-nil")
	}
}

func TestInternal_I15_TraceCollection_ParseInjectUsingJson_Error(t *testing.T) {
	badJson := corejson.NewPtr("{invalid")
	tc := New.traces.Default()
	_, err := tc.ParseInjectUsingJson(badJson)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestInternal_I15_TraceCollection_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	badJson := corejson.NewPtr("{invalid")
	tc := New.traces.Default()
	tc.ParseInjectUsingJsonMust(badJson)
}

func TestInternal_I15_DirGetter_Get_HighSkip(t *testing.T) {
	if Dir.Get(99999) != "" {
		t.Fatal("expected empty for unreachable skip")
	}
}

func TestInternal_I15_FileGetter_Name_HighSkip(t *testing.T) {
	if File.Name(99999) != "" {
		t.Fatal("expected empty for unreachable skip")
	}
}

func TestInternal_I15_FileGetter_Path_HighSkip(t *testing.T) {
	if File.Path(99999) != "" {
		t.Fatal("expected empty for unreachable skip")
	}
}

func TestInternal_I15_NewTraceCollection_Using_NilTraces(t *testing.T) {
	tc := New.traces.Using(false)
	if tc == nil || tc.Length() != 0 {
		t.Fatal("expected empty collection")
	}
}

func TestInternal_I15_NewTraceCollection_Using_Clone(t *testing.T) {
	tr := New.Create(0)
	tc := New.traces.Using(true, tr)
	if tc == nil || tc.Length() != 1 {
		t.Fatal("expected 1 trace")
	}
}

func TestInternal_I15_Trace_JsonRoundtrip(t *testing.T) {
	tr := New.Create(0)
	b, err := json.Marshal(tr)
	if err != nil {
		t.Fatal(err)
	}
	var tr2 Trace
	if err := json.Unmarshal(b, &tr2); err != nil {
		t.Fatal(err)
	}
}
