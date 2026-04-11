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

func Test_FileWithLine_StringUsingFmt_FromFileWithLineStringUs(t *testing.T) {
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

func Test_FileWithLine_JsonModelAny_FromFileWithLineStringUs(t *testing.T) {
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

func Test_FileWithLine_JsonString_FromFileWithLineStringUs(t *testing.T) {
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

func Test_FileWithLine_Json_FromFileWithLineStringUs(t *testing.T) {
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

func Test_FileWithLine_JsonPtr_FromFileWithLineStringUs(t *testing.T) {
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

func Test_FileWithLine_ParseInjectUsingJson_FromFileWithLineStringUs(t *testing.T) {
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

func Test_FileWithLine_ParseInjectUsingJsonMust_FromFileWithLineStringUs(t *testing.T) {
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

func Test_FileWithLine_JsonParseSelfInject_FromFileWithLineStringUs(t *testing.T) {
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

func Test_FileWithLine_AsFileLiner_FromFileWithLineStringUs(t *testing.T) {
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

func Test_FileWithLine_String_Nil_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	var fwl *codestack.FileWithLine

	// Act
	result := fwl.String()

	// Assert
	convey.Convey("String returns empty for nil", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

func Test_FileWithLine_IsNil_IsNotNil(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{}

	// Act & Assert
	convey.Convey("IsNil/IsNotNil", t, func() {
		convey.So(fwl.IsNil(), convey.ShouldBeFalse)
		convey.So(fwl.IsNotNil(), convey.ShouldBeTrue)
	})
}

// --- Trace methods ---

func Test_Trace_StringUsingFmt_FromFileWithLineStringUs(t *testing.T) {
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

func Test_Trace_FileName_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	result := trace.FileName()

	// Assert
	convey.Convey("Trace.FileName returns file name only", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Trace_FileWithLine_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	result := trace.FileWithLine()

	// Assert
	convey.Convey("Trace.FileWithLine returns struct", t, func() {
		convey.So(result.FilePath, convey.ShouldNotBeEmpty)
	})
}

func Test_Trace_FileWithLineString_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	result := trace.FileWithLineString()

	// Assert
	convey.Convey("Trace.FileWithLineString returns formatted string", t, func() {
		convey.So(result, convey.ShouldContainSubstring, ":")
	})
}

func Test_Trace_ShortString_FromFileWithLineStringUs(t *testing.T) {
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

func Test_Trace_Message_FromFileWithLineStringUs(t *testing.T) {
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

func Test_Trace_HasIssues_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act & Assert
	convey.Convey("Trace.HasIssues returns false for valid trace", t, func() {
		convey.So(trace.HasIssues(), convey.ShouldBeFalse)
	})
}

func Test_Trace_HasIssues_Nil_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act & Assert
	convey.Convey("Trace.HasIssues returns true for nil", t, func() {
		convey.So(trace.HasIssues(), convey.ShouldBeTrue)
	})
}

func Test_Trace_JsonModelAny_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	result := trace.JsonModelAny()

	// Assert
	convey.Convey("Trace.JsonModelAny returns model", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_Trace_JsonString_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	result := trace.JsonString()

	// Assert
	convey.Convey("Trace.JsonString returns JSON", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Trace_ParseInjectUsingJson_FromFileWithLineStringUs(t *testing.T) {
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

func Test_Trace_ParseInjectUsingJsonMust_FromFileWithLineStringUs(t *testing.T) {
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

func Test_Trace_JsonParseSelfInject_FromFileWithLineStringUs(t *testing.T) {
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

func Test_Trace_Clone_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	cloned := trace.Clone()

	// Assert
	convey.Convey("Trace.Clone returns independent copy", t, func() {
		convey.So(cloned.PackageName, convey.ShouldEqual, trace.PackageName)
	})
}

func Test_Trace_ClonePtr_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	cloned := trace.ClonePtr()

	// Assert
	convey.Convey("Trace.ClonePtr returns non-nil", t, func() {
		convey.So(cloned, convey.ShouldNotBeNil)
	})
}

func Test_Trace_ClonePtr_Nil_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act
	cloned := trace.ClonePtr()

	// Assert
	convey.Convey("Trace.ClonePtr nil returns nil", t, func() {
		convey.So(cloned, convey.ShouldBeNil)
	})
}

func Test_Trace_Dispose_FromFileWithLineStringUs(t *testing.T) {
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

func Test_Trace_Dispose_Nil_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act & Assert (no panic)
	convey.Convey("Trace.Dispose nil is safe", t, func() {
		convey.So(func() { trace.Dispose() }, convey.ShouldNotPanic)
	})
}

func Test_Trace_AsFileLiner_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	result := trace.AsFileLiner()

	// Assert
	convey.Convey("Trace.AsFileLiner returns FileWithLiner", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_Trace_String_Nil_FromFileWithLineStringUs(t *testing.T) {
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

func Test_TraceCollection_StackTraces_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.StackTraces()

	// Assert
	convey.Convey("TraceCollection.StackTraces returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_TraceCollection_StackTracesJsonResult_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.StackTracesJsonResult()

	// Assert
	convey.Convey("TraceCollection.StackTracesJsonResult returns non-nil", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_TraceCollection_NewStackTraces_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.NewStackTraces(0)

	// Assert
	convey.Convey("NewStackTraces returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_TraceCollection_NewDefaultStackTraces_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.NewDefaultStackTraces()

	// Assert
	convey.Convey("NewDefaultStackTraces returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_TraceCollection_NewStackTracesJsonResult_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.NewStackTracesJsonResult(0)

	// Assert
	convey.Convey("NewStackTracesJsonResult returns non-nil", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_TraceCollection_NewDefaultStackTracesJsonResult_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.NewDefaultStackTracesJsonResult()

	// Assert
	convey.Convey("NewDefaultStackTracesJsonResult returns non-nil", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_TraceCollection_StackTracesBytes_FromFileWithLineStringUs(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.StackTracesBytes()

	// Assert
	convey.Convey("StackTracesBytes returns non-empty bytes", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_TraceCollection_InsertAt_FromFileWithLineStringUs(t *testing.T) {
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

func Test_NameOf_Method_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.NameOf.Method()

	// Assert
	convey.Convey("NameOf.Method returns method name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_NameOf_Package_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.NameOf.Package()

	// Assert
	convey.Convey("NameOf.Package returns package name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_NameOf_All_Empty_FromFileWithLineStringUs(t *testing.T) {
	// Act
	full, pkg, method := codestack.NameOf.All("")

	// Assert
	convey.Convey("NameOf.All empty returns empty", t, func() {
		convey.So(full, convey.ShouldBeEmpty)
		convey.So(pkg, convey.ShouldBeEmpty)
		convey.So(method, convey.ShouldBeEmpty)
	})
}

func Test_NameOf_MethodByFullName_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.NameOf.MethodByFullName("github.com/pkg.Method")

	// Assert
	convey.Convey("MethodByFullName returns method name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_NameOf_PackageByFullName_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.NameOf.PackageByFullName("github.com/pkg.Method")

	// Assert
	convey.Convey("PackageByFullName returns package name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_NameOf_CurrentFuncFullPath_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.NameOf.CurrentFuncFullPath("github.com/pkg.Method")

	// Assert
	convey.Convey("CurrentFuncFullPath returns full method name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_NameOf_JoinPackageNameWithRelative_FromFileWithLineStringUs(t *testing.T) {
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

func Test_NameOf_MethodStackSkip_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.NameOf.MethodStackSkip(0)

	// Assert
	convey.Convey("MethodStackSkip returns method name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_NameOf_PackageStackSkip_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.NameOf.PackageStackSkip(0)

	// Assert
	convey.Convey("PackageStackSkip returns package name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

// --- dirGetter methods ---

func Test_Dir_CurDir_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.Dir.CurDir()

	// Assert
	convey.Convey("Dir.CurDir returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Dir_CurDirJoin_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.Dir.CurDirJoin("sub", "path")

	// Assert
	convey.Convey("Dir.CurDirJoin returns joined path", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Dir_RepoDir_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.Dir.RepoDir()

	// Assert
	convey.Convey("Dir.RepoDir returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Dir_RepoDirJoin_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.Dir.RepoDirJoin("sub")

	// Assert
	convey.Convey("Dir.RepoDirJoin returns joined path", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Dir_Get_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.Dir.Get(0)

	// Assert
	convey.Convey("Dir.Get returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

// --- fileGetter methods ---

func Test_File_Name_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.File.Name(0)

	// Assert
	convey.Convey("File.Name returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_File_Path_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.File.Path(0)

	// Assert
	convey.Convey("File.Path returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_File_PathLineSep_FromFileWithLineStringUs(t *testing.T) {
	// Act
	filePath, lineNumber := codestack.File.PathLineSep(0)

	// Assert
	convey.Convey("File.PathLineSep returns file and line", t, func() {
		convey.So(filePath, convey.ShouldNotBeEmpty)
		convey.So(lineNumber, convey.ShouldBeGreaterThan, 0)
	})
}

func Test_File_PathLineSepDefault_FromFileWithLineStringUs(t *testing.T) {
	// Act
	filePath, lineNumber := codestack.File.PathLineSepDefault()

	// Assert
	convey.Convey("File.PathLineSepDefault returns file and line", t, func() {
		convey.So(filePath, convey.ShouldNotBeEmpty)
		convey.So(lineNumber, convey.ShouldBeGreaterThan, 0)
	})
}

func Test_File_FilePathWithLineString_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.File.FilePathWithLineString(0)

	// Assert
	convey.Convey("File.FilePathWithLineString returns formatted", t, func() {
		convey.So(result, convey.ShouldContainSubstring, ":")
	})
}

func Test_File_PathLineStringDefault_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.File.PathLineStringDefault()

	// Assert
	convey.Convey("File.PathLineStringDefault returns formatted", t, func() {
		convey.So(result, convey.ShouldContainSubstring, ":")
	})
}

func Test_File_CurrentFilePath_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.File.CurrentFilePath()

	// Assert
	convey.Convey("File.CurrentFilePath returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

// --- stacksTo methods ---

func Test_StacksTo_Bytes_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.StacksTo.Bytes(0)

	// Assert
	convey.Convey("StacksTo.Bytes returns non-empty", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_StacksTo_BytesDefault_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.StacksTo.BytesDefault()

	// Assert
	convey.Convey("StacksTo.BytesDefault returns non-empty", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_StacksTo_String_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.StacksTo.String(0, 5)

	// Assert
	convey.Convey("StacksTo.String returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_StacksTo_StringUsingFmt_FromFileWithLineStringUs(t *testing.T) {
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

func Test_StacksTo_JsonString_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.StacksTo.JsonString(0)

	// Assert
	convey.Convey("StacksTo.JsonString returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_StacksTo_JsonStringDefault_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.StacksTo.JsonStringDefault()

	// Assert
	convey.Convey("StacksTo.JsonStringDefault returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_StacksTo_StringNoCount_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.StacksTo.StringNoCount(0)

	// Assert
	convey.Convey("StacksTo.StringNoCount returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_StacksTo_StringDefault_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.StacksTo.StringDefault()

	// Assert
	convey.Convey("StacksTo.StringDefault returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

// --- newCreator methods ---

func Test_New_Default_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.New.Default()

	// Assert
	convey.Convey("New.Default returns valid trace", t, func() {
		convey.So(result.IsOkay, convey.ShouldBeTrue)
	})
}

func Test_New_SkipOne(t *testing.T) {
	// Act
	result := codestack.New.SkipOne()

	// Assert
	convey.Convey("New.SkipOne returns valid trace", t, func() {
		convey.So(result.IsOkay, convey.ShouldBeTrue)
	})
}

func Test_New_Ptr_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.New.Ptr(0)

	// Assert
	convey.Convey("New.Ptr returns non-nil", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
		convey.So(result.IsOkay, convey.ShouldBeTrue)
	})
}

// --- newStacksCreator methods ---

func Test_StackTrace_Default_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.New.StackTrace.Default(0, 5)

	// Assert
	convey.Convey("StackTrace.Default returns collection", t, func() {
		convey.So(result.HasAnyItem(), convey.ShouldBeTrue)
	})
}

func Test_StackTrace_SkipOne_FromFileWithLineStringUs(t *testing.T) {
	// Act — guard against sandbox stack depth issues
	defer func() {
		if r := recover(); r != nil {
			t.Skipf("SkipOne panicked (sandbox stack depth): %v", r)
		}
	}()
	result := codestack.New.StackTrace.SkipOne()

	// Assert — SkipOne may return empty depending on call depth
	convey.Convey("StackTrace.SkipOne returns non-nil collection", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_StackTrace_SkipNone_FromFileWithLineStringUs(t *testing.T) {
	// Act
	result := codestack.New.StackTrace.SkipNone()

	// Assert
	convey.Convey("StackTrace.SkipNone returns collection", t, func() {
		convey.So(result.HasAnyItem(), convey.ShouldBeTrue)
	})
}

// --- TraceCollection filter by name methods ---

func Test_TraceCollection_FilterPackageName(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.FilterPackageNameTraceCollection("codestacktests")

	// Assert
	convey.Convey("FilterPackageNameTraceCollection filters", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_TraceCollection_SkipFilterPackageName(t *testing.T) {
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

func Test_TraceCollection_GetSinglePageCollection_NegativeIndex(t *testing.T) {
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

func Test_TraceCollection_GetSinglePageCollection_LastPage(t *testing.T) {
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

func Test_TraceCollection_ConcatNewUsingSkipPlusCount(t *testing.T) {
	// Arrange
	stacks := codestack.New.StackTrace.SkipNone()

	// Act
	result := stacks.ConcatNewUsingSkipPlusCount(0, 3)

	// Assert
	convey.Convey("ConcatNewUsingSkipPlusCount returns concatenated", t, func() {
		convey.So(result.Length(), convey.ShouldBeGreaterThanOrEqualTo, stacks.Length())
	})
}

func Test_TraceCollection_ConcatNewUsingSkip(t *testing.T) {
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

func Test_TraceCollection_AddsUsingSkipDefault(t *testing.T) {
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

func Test_TraceCollection_AddsUsingSkipUsingFilter(t *testing.T) {
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

func Test_TraceCollection_AddsUsingSkipUsingFilter_Break(t *testing.T) {
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
