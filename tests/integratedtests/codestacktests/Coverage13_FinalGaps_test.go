package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage13 — codestack final coverage gaps
// ══════════════════════════════════════════════════════════════════════════════

// --- FileWithLine methods ---

func Test_Cov13_FileWithLine_StringUsingFmt(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{
		FilePath: "/some/path.go",
		Line:     42,
	}

	// Act
	result := fwl.StringUsingFmt(func(f codestack.FileWithLine) string {
		return f.FilePath
	})

	// Assert
	convey.Convey("StringUsingFmt calls formatter", t, func() {
		convey.So(result, convey.ShouldEqual, "/some/path.go")
	})
}

func Test_Cov13_FileWithLine_JsonModelAny(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{
		FilePath: "/path.go",
		Line:     10,
	}

	// Act
	result := fwl.JsonModelAny()

	// Assert
	convey.Convey("JsonModelAny returns FileWithLine", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_Cov13_FileWithLine_JsonString(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{
		FilePath: "/path.go",
		Line:     10,
	}

	// Act
	result := fwl.JsonString()

	// Assert
	convey.Convey("JsonString returns JSON", t, func() {
		convey.So(result, convey.ShouldContainSubstring, "path.go")
	})
}

func Test_Cov13_FileWithLine_Json(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{
		FilePath: "/path.go",
		Line:     10,
	}

	// Act
	result := fwl.Json()

	// Assert
	convey.Convey("Json returns Result", t, func() {
		convey.So(result.JsonString(), convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_FileWithLine_JsonPtr(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{
		FilePath: "/path.go",
		Line:     10,
	}

	// Act
	result := fwl.JsonPtr()

	// Assert
	convey.Convey("JsonPtr returns non-nil", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_Cov13_FileWithLine_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{
		FilePath: "/path.go",
		Line:     10,
	}
	jsonResult := fwl.JsonPtr()

	target := &codestack.FileWithLine{}

	// Act
	parsed, err := target.ParseInjectUsingJson(jsonResult)

	// Assert
	convey.Convey("ParseInjectUsingJson succeeds", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(parsed.FilePath, convey.ShouldEqual, "/path.go")
	})
}

func Test_Cov13_FileWithLine_ParseInjectUsingJsonMust(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{
		FilePath: "/path.go",
		Line:     10,
	}
	jsonResult := fwl.JsonPtr()

	target := &codestack.FileWithLine{}

	// Act
	parsed := target.ParseInjectUsingJsonMust(jsonResult)

	// Assert
	convey.Convey("ParseInjectUsingJsonMust succeeds", t, func() {
		convey.So(parsed.FilePath, convey.ShouldEqual, "/path.go")
	})
}

func Test_Cov13_FileWithLine_JsonParseSelfInject(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{
		FilePath: "/path.go",
		Line:     10,
	}
	jsonResult := fwl.JsonPtr()

	target := &codestack.FileWithLine{}

	// Act
	err := target.JsonParseSelfInject(jsonResult)

	// Assert
	convey.Convey("JsonParseSelfInject succeeds", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(target.FilePath, convey.ShouldEqual, "/path.go")
	})
}

func Test_Cov13_FileWithLine_AsFileLiner(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{
		FilePath: "/path.go",
		Line:     10,
	}

	// Act
	result := fwl.AsFileLiner()

	// Assert
	convey.Convey("AsFileLiner returns FileWithLiner", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
		convey.So(result.FullFilePath(), convey.ShouldEqual, "/path.go")
	})
}

func Test_Cov13_FileWithLine_String_Nil(t *testing.T) {
	// Arrange
	var fwl *codestack.FileWithLine

	// Act
	result := fwl.String()

	// Assert
	convey.Convey("String returns empty for nil", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

func Test_Cov13_FileWithLine_IsNil_IsNotNil(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{}

	// Act & Assert
	convey.Convey("IsNil/IsNotNil", t, func() {
		convey.So(fwl.IsNil(), convey.ShouldBeFalse)
		convey.So(fwl.IsNotNil(), convey.ShouldBeTrue)
	})
}

// --- Trace methods ---

func Test_Cov13_Trace_StringUsingFmt(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	result := trace.StringUsingFmt(func(tr codestack.Trace) string {
		return tr.PackageName
	})

	// Assert
	convey.Convey("Trace.StringUsingFmt calls formatter", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_Trace_FileName(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	result := trace.FileName()

	// Assert
	convey.Convey("Trace.FileName returns file name only", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_Trace_FileWithLine(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	result := trace.FileWithLine()

	// Assert
	convey.Convey("Trace.FileWithLine returns struct", t, func() {
		convey.So(result.FilePath, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_Trace_FileWithLineString(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	result := trace.FileWithLineString()

	// Assert
	convey.Convey("Trace.FileWithLineString returns formatted string", t, func() {
		convey.So(result, convey.ShouldContainSubstring, ":")
	})
}

func Test_Cov13_Trace_ShortString(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	result1 := trace.ShortString()
	result2 := trace.ShortString() // cached

	// Assert
	convey.Convey("Trace.ShortString returns and caches", t, func() {
		convey.So(result1, convey.ShouldNotBeEmpty)
		convey.So(result2, convey.ShouldEqual, result1)
	})
}

func Test_Cov13_Trace_Message(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	result1 := trace.Message()
	result2 := trace.Message() // cached

	// Assert
	convey.Convey("Trace.Message returns and caches", t, func() {
		convey.So(result1, convey.ShouldNotBeEmpty)
		convey.So(result2, convey.ShouldEqual, result1)
	})
}

func Test_Cov13_Trace_HasIssues(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act & Assert
	convey.Convey("Trace.HasIssues returns false for valid trace", t, func() {
		convey.So(trace.HasIssues(), convey.ShouldBeFalse)
	})
}

func Test_Cov13_Trace_HasIssues_Nil(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act & Assert
	convey.Convey("Trace.HasIssues returns true for nil", t, func() {
		convey.So(trace.HasIssues(), convey.ShouldBeTrue)
	})
}

func Test_Cov13_Trace_JsonModelAny(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	result := trace.JsonModelAny()

	// Assert
	convey.Convey("Trace.JsonModelAny returns model", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_Cov13_Trace_JsonString(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	result := trace.JsonString()

	// Assert
	convey.Convey("Trace.JsonString returns JSON", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_Trace_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	jsonResult := trace.JsonPtr()
	target := &codestack.Trace{}

	// Act
	parsed, err := target.ParseInjectUsingJson(jsonResult)

	// Assert
	convey.Convey("Trace.ParseInjectUsingJson succeeds", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(parsed.PackageName, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_Trace_ParseInjectUsingJsonMust(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	jsonResult := trace.JsonPtr()
	target := &codestack.Trace{}

	// Act
	parsed := target.ParseInjectUsingJsonMust(jsonResult)

	// Assert
	convey.Convey("Trace.ParseInjectUsingJsonMust succeeds", t, func() {
		convey.So(parsed.PackageName, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_Trace_JsonParseSelfInject(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	jsonResult := trace.JsonPtr()
	target := &codestack.Trace{}

	// Act
	err := target.JsonParseSelfInject(jsonResult)

	// Assert
	convey.Convey("Trace.JsonParseSelfInject succeeds", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_Cov13_Trace_Clone(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	cloned := trace.Clone()

	// Assert
	convey.Convey("Trace.Clone returns independent copy", t, func() {
		convey.So(cloned.PackageName, convey.ShouldEqual, trace.PackageName)
	})
}

func Test_Cov13_Trace_ClonePtr(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	cloned := trace.ClonePtr()

	// Assert
	convey.Convey("Trace.ClonePtr returns non-nil", t, func() {
		convey.So(cloned, convey.ShouldNotBeNil)
	})
}

func Test_Cov13_Trace_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act
	cloned := trace.ClonePtr()

	// Assert
	convey.Convey("Trace.ClonePtr nil returns nil", t, func() {
		convey.So(cloned, convey.ShouldBeNil)
	})
}

func Test_Cov13_Trace_Dispose(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	trace.Dispose()

	// Assert
	convey.Convey("Trace.Dispose clears fields", t, func() {
		convey.So(trace.PackageName, convey.ShouldBeEmpty)
		convey.So(trace.IsOkay, convey.ShouldBeFalse)
	})
}

func Test_Cov13_Trace_Dispose_Nil(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act & Assert (no panic)
	convey.Convey("Trace.Dispose nil is safe", t, func() {
		convey.So(func() { trace.Dispose() }, convey.ShouldNotPanic)
	})
}

func Test_Cov13_Trace_AsFileLiner(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	result := trace.AsFileLiner()

	// Assert
	convey.Convey("Trace.AsFileLiner returns FileWithLiner", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_Cov13_Trace_String_Nil(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act
	result := trace.String()

	// Assert
	convey.Convey("Trace.String nil returns empty", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

// --- TraceCollection methods ---

func Test_Cov13_TraceCollection_StackTraces(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.StackTraces()

	// Assert
	convey.Convey("TraceCollection.StackTraces returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_TraceCollection_StackTracesJsonResult(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.StackTracesJsonResult()

	// Assert
	convey.Convey("TraceCollection.StackTracesJsonResult returns non-nil", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_Cov13_TraceCollection_NewStackTraces(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.NewStackTraces(0)

	// Assert
	convey.Convey("NewStackTraces returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_TraceCollection_NewDefaultStackTraces(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.NewDefaultStackTraces()

	// Assert
	convey.Convey("NewDefaultStackTraces returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_TraceCollection_NewStackTracesJsonResult(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.NewStackTracesJsonResult(0)

	// Assert
	convey.Convey("NewStackTracesJsonResult returns non-nil", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_Cov13_TraceCollection_NewDefaultStackTracesJsonResult(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.NewDefaultStackTracesJsonResult()

	// Assert
	convey.Convey("NewDefaultStackTracesJsonResult returns non-nil", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_Cov13_TraceCollection_StackTracesBytes(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.StackTracesBytes()

	// Assert
	convey.Convey("StackTracesBytes returns non-empty bytes", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_Cov13_TraceCollection_InsertAt(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()
	trace := codestack.New.Default()

	// Act
	stacks.InsertAt(0, trace)

	// Assert
	convey.Convey("InsertAt adds at index", t, func() {
		convey.So(stacks.Length(), convey.ShouldBeGreaterThan, 0)
	})
}

// --- currentNameOf methods ---

func Test_Cov13_NameOf_Method(t *testing.T) {
	// Act
	result := codestack.NameOf.Method()

	// Assert
	convey.Convey("NameOf.Method returns method name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_NameOf_Package(t *testing.T) {
	// Act
	result := codestack.NameOf.Package()

	// Assert
	convey.Convey("NameOf.Package returns package name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_NameOf_All_Empty(t *testing.T) {
	// Act
	full, pkg, method := codestack.NameOf.All("")

	// Assert
	convey.Convey("NameOf.All empty returns empty", t, func() {
		convey.So(full, convey.ShouldBeEmpty)
		convey.So(pkg, convey.ShouldBeEmpty)
		convey.So(method, convey.ShouldBeEmpty)
	})
}

func Test_Cov13_NameOf_MethodByFullName(t *testing.T) {
	// Act
	result := codestack.NameOf.MethodByFullName("github.com/pkg.Method")

	// Assert
	convey.Convey("MethodByFullName returns method name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_NameOf_PackageByFullName(t *testing.T) {
	// Act
	result := codestack.NameOf.PackageByFullName("github.com/pkg.Method")

	// Assert
	convey.Convey("PackageByFullName returns package name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_NameOf_CurrentFuncFullPath(t *testing.T) {
	// Act
	result := codestack.NameOf.CurrentFuncFullPath("github.com/pkg.Method")

	// Assert
	convey.Convey("CurrentFuncFullPath returns full method name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_NameOf_JoinPackageNameWithRelative(t *testing.T) {
	// Act
	result := codestack.NameOf.JoinPackageNameWithRelative(
		"github.com/pkg.Method",
		"SubMethod",
	)

	// Assert
	convey.Convey("JoinPackageNameWithRelative returns joined name", t, func() {
		convey.So(result, convey.ShouldContainSubstring, "SubMethod")
	})
}

func Test_Cov13_NameOf_MethodStackSkip(t *testing.T) {
	// Act
	result := codestack.NameOf.MethodStackSkip(0)

	// Assert
	convey.Convey("MethodStackSkip returns method name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_NameOf_PackageStackSkip(t *testing.T) {
	// Act
	result := codestack.NameOf.PackageStackSkip(0)

	// Assert
	convey.Convey("PackageStackSkip returns package name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

// --- dirGetter methods ---

func Test_Cov13_Dir_CurDir(t *testing.T) {
	// Act
	result := codestack.Dir.CurDir()

	// Assert
	convey.Convey("Dir.CurDir returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_Dir_CurDirJoin(t *testing.T) {
	// Act
	result := codestack.Dir.CurDirJoin("sub", "path")

	// Assert
	convey.Convey("Dir.CurDirJoin returns joined path", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_Dir_RepoDir(t *testing.T) {
	// Act
	result := codestack.Dir.RepoDir()

	// Assert
	convey.Convey("Dir.RepoDir returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_Dir_RepoDirJoin(t *testing.T) {
	// Act
	result := codestack.Dir.RepoDirJoin("sub")

	// Assert
	convey.Convey("Dir.RepoDirJoin returns joined path", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_Dir_Get(t *testing.T) {
	// Act
	result := codestack.Dir.Get(0)

	// Assert
	convey.Convey("Dir.Get returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

// --- fileGetter methods ---

func Test_Cov13_File_Name(t *testing.T) {
	// Act
	result := codestack.File.Name(0)

	// Assert
	convey.Convey("File.Name returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_File_Path(t *testing.T) {
	// Act
	result := codestack.File.Path(0)

	// Assert
	convey.Convey("File.Path returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_File_PathLineSep(t *testing.T) {
	// Act
	filePath, lineNumber := codestack.File.PathLineSep(0)

	// Assert
	convey.Convey("File.PathLineSep returns file and line", t, func() {
		convey.So(filePath, convey.ShouldNotBeEmpty)
		convey.So(lineNumber, convey.ShouldBeGreaterThan, 0)
	})
}

func Test_Cov13_File_PathLineSepDefault(t *testing.T) {
	// Act
	filePath, lineNumber := codestack.File.PathLineSepDefault()

	// Assert
	convey.Convey("File.PathLineSepDefault returns file and line", t, func() {
		convey.So(filePath, convey.ShouldNotBeEmpty)
		convey.So(lineNumber, convey.ShouldBeGreaterThan, 0)
	})
}

func Test_Cov13_File_FilePathWithLineString(t *testing.T) {
	// Act
	result := codestack.File.FilePathWithLineString(0)

	// Assert
	convey.Convey("File.FilePathWithLineString returns formatted", t, func() {
		convey.So(result, convey.ShouldContainSubstring, ":")
	})
}

func Test_Cov13_File_PathLineStringDefault(t *testing.T) {
	// Act
	result := codestack.File.PathLineStringDefault()

	// Assert
	convey.Convey("File.PathLineStringDefault returns formatted", t, func() {
		convey.So(result, convey.ShouldContainSubstring, ":")
	})
}

func Test_Cov13_File_CurrentFilePath(t *testing.T) {
	// Act
	result := codestack.File.CurrentFilePath()

	// Assert
	convey.Convey("File.CurrentFilePath returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

// --- stacksTo methods ---

func Test_Cov13_StacksTo_Bytes(t *testing.T) {
	// Act
	result := codestack.StacksTo.Bytes(0)

	// Assert
	convey.Convey("StacksTo.Bytes returns non-empty", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_Cov13_StacksTo_BytesDefault(t *testing.T) {
	// Act
	result := codestack.StacksTo.BytesDefault()

	// Assert
	convey.Convey("StacksTo.BytesDefault returns non-empty", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_Cov13_StacksTo_String(t *testing.T) {
	// Act
	result := codestack.StacksTo.String(0, 5)

	// Assert
	convey.Convey("StacksTo.String returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_StacksTo_StringUsingFmt(t *testing.T) {
	// Act
	result := codestack.StacksTo.StringUsingFmt(
		func(trace *codestack.Trace) string {
			return trace.PackageName
		},
		0, 5,
	)

	// Assert
	convey.Convey("StacksTo.StringUsingFmt returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_StacksTo_JsonString(t *testing.T) {
	// Act
	result := codestack.StacksTo.JsonString(0)

	// Assert
	convey.Convey("StacksTo.JsonString returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_StacksTo_JsonStringDefault(t *testing.T) {
	// Act
	result := codestack.StacksTo.JsonStringDefault()

	// Assert
	convey.Convey("StacksTo.JsonStringDefault returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_StacksTo_StringNoCount(t *testing.T) {
	// Act
	result := codestack.StacksTo.StringNoCount(0)

	// Assert
	convey.Convey("StacksTo.StringNoCount returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Cov13_StacksTo_StringDefault(t *testing.T) {
	// Act
	result := codestack.StacksTo.StringDefault()

	// Assert
	convey.Convey("StacksTo.StringDefault returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

// --- newCreator methods ---

func Test_Cov13_New_Default(t *testing.T) {
	// Act
	result := codestack.New.Default()

	// Assert
	convey.Convey("New.Default returns valid trace", t, func() {
		convey.So(result.IsOkay, convey.ShouldBeTrue)
	})
}

func Test_Cov13_New_SkipOne(t *testing.T) {
	// Act
	result := codestack.New.SkipOne()

	// Assert
	convey.Convey("New.SkipOne returns valid trace", t, func() {
		convey.So(result.IsOkay, convey.ShouldBeTrue)
	})
}

func Test_Cov13_New_Ptr(t *testing.T) {
	// Act
	result := codestack.New.Ptr(0)

	// Assert
	convey.Convey("New.Ptr returns non-nil", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
		convey.So(result.IsOkay, convey.ShouldBeTrue)
	})
}

// --- newStacksCreator methods ---

func Test_Cov13_StackTrace_Default(t *testing.T) {
	// Act
	result := codestack.New.StackTrace.Default(0, 5)

	// Assert
	convey.Convey("StackTrace.Default returns collection", t, func() {
		convey.So(result.HasAnyItem(), convey.ShouldBeTrue)
	})
}

func Test_Cov13_StackTrace_SkipOne(t *testing.T) {
	// Act
	result := codestack.New.StackTrace.SkipOne()

	// Assert
	convey.Convey("StackTrace.SkipOne returns collection", t, func() {
		convey.So(result.HasAnyItem(), convey.ShouldBeTrue)
	})
}

func Test_Cov13_StackTrace_SkipNone(t *testing.T) {
	// Act
	result := codestack.New.StackTrace.SkipNone()

	// Assert
	convey.Convey("StackTrace.SkipNone returns collection", t, func() {
		convey.So(result.HasAnyItem(), convey.ShouldBeTrue)
	})
}

// --- TraceCollection filter by name methods ---

func Test_Cov13_TraceCollection_FilterPackageName(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.FilterPackageNameTraceCollection("codestacktests")

	// Assert
	convey.Convey("FilterPackageNameTraceCollection filters", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_Cov13_TraceCollection_SkipFilterPackageName(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.SkipFilterPackageNameTraceCollection("codestacktests")

	// Assert
	convey.Convey("SkipFilterPackageNameTraceCollection skips", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

// --- TraceCollection GetSinglePageCollection negative index ---

func Test_Cov13_TraceCollection_GetSinglePageCollection_NegativeIndex(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()
	for i := 0; i < 20; i++ {
		stacks.Add(codestack.New.Default())
	}

	// Act & Assert
	convey.Convey("GetSinglePageCollection panics on negative index", t, func() {
		convey.So(func() {
			stacks.GetSinglePageCollection(5, 0)
		}, convey.ShouldPanic)
	})
}

func Test_Cov13_TraceCollection_GetSinglePageCollection_LastPage(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()
	for i := 0; i < 25; i++ {
		stacks.Add(codestack.New.Default())
	}

	// Act — request page beyond items
	result := stacks.GetSinglePageCollection(10, 5)

	// Assert
	convey.Convey("GetSinglePageCollection last page handles end > length", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

// --- TraceCollection ConcatNewUsingSkipPlusCount ---

func Test_Cov13_TraceCollection_ConcatNewUsingSkipPlusCount(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.ConcatNewUsingSkipPlusCount(0, 3)

	// Assert
	convey.Convey("ConcatNewUsingSkipPlusCount returns concatenated", t, func() {
		convey.So(result.Length(), convey.ShouldBeGreaterThanOrEqualTo, stacks.Length())
	})
}

func Test_Cov13_TraceCollection_ConcatNewUsingSkip(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.ConcatNewUsingSkip(0)

	// Assert
	convey.Convey("ConcatNewUsingSkip returns concatenated", t, func() {
		convey.So(result.Length(), convey.ShouldBeGreaterThanOrEqualTo, stacks.Length())
	})
}

// --- TraceCollection AddsUsingSkipDefault ---

func Test_Cov13_TraceCollection_AddsUsingSkipDefault(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	stacks.AddsUsingSkipDefault(0)

	// Assert
	convey.Convey("AddsUsingSkipDefault adds traces", t, func() {
		convey.So(stacks.HasAnyItem(), convey.ShouldBeTrue)
	})
}

// --- TraceCollection AddsUsingSkipUsingFilter ---

func Test_Cov13_TraceCollection_AddsUsingSkipUsingFilter(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	stacks.AddsUsingSkipUsingFilter(
		true,
		true,
		0,
		5,
		func(trace *codestack.Trace) (isTake, isBreak bool) {
			return true, false
		},
	)

	// Assert
	convey.Convey("AddsUsingSkipUsingFilter adds traces", t, func() {
		convey.So(stacks.HasAnyItem(), convey.ShouldBeTrue)
	})
}

func Test_Cov13_TraceCollection_AddsUsingSkipUsingFilter_Break(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()
	initialLen := stacks.Length()

	// Act
	stacks.AddsUsingSkipUsingFilter(
		true,
		true,
		0,
		5,
		func(trace *codestack.Trace) (isTake, isBreak bool) {
			return true, true // break immediately
		},
	)

	// Assert
	convey.Convey("AddsUsingSkipUsingFilter breaks early", t, func() {
		convey.So(stacks.Length(), convey.ShouldBeLessThanOrEqualTo, initialLen+2)
	})
}
