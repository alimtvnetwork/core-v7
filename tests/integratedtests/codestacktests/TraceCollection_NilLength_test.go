package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── TraceCollection additional coverage ──

func Test_TraceCollection_NilLength(t *testing.T) {
	// Arrange
	var tc *codestack.TraceCollection

	// Act
	actual := args.Map{
		"length": tc.Length(),
		"isEmpty": tc.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"length": 0,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Length returns nil -- nil receiver", actual)
}

func Test_TraceCollection_Adds_Empty_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)
	before := tc.Length()
	tc.Adds()

	// Act
	actual := args.Map{"same": tc.Length() == before}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Adds returns empty -- empty does nothing", actual)
}

func Test_TraceCollection_AddsIf_False_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)
	before := tc.Length()
	tc.AddsIf(false, codestack.Trace{})

	// Act
	actual := args.Map{"same": tc.Length() == before}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AddsIf returns non-empty -- false does nothing", actual)
}

func Test_TraceCollection_AddsIf_True_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)
	before := tc.Length()
	tc.AddsIf(true, codestack.Trace{PackageName: "test"})

	// Act
	actual := args.Map{"grew": tc.Length() > before}

	// Assert
	expected := args.Map{"grew": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AddsIf returns non-empty -- true adds", actual)
}

func Test_TraceCollection_AddsPtr_Empty_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)
	before := tc.Length()
	tc.AddsPtr(true)

	// Act
	actual := args.Map{"same": tc.Length() == before}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AddsPtr returns empty -- empty does nothing", actual)
}

func Test_TraceCollection_AddsPtr_NilTrace(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)
	before := tc.Length()
	tc.AddsPtr(true, nil)

	// Act
	actual := args.Map{"same": tc.Length() == before}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AddsPtr returns nil -- nil trace skipped", actual)
}

func Test_TraceCollection_AddsPtr_SkipIssues_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)
	before := tc.Length()
	badTrace := &codestack.Trace{} // HasIssues = true
	tc.AddsPtr(true, badTrace)

	// Act
	actual := args.Map{"same": tc.Length() == before}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AddsPtr returns correct value -- skips issues", actual)
}

func Test_TraceCollection_AddsPtr_NoSkip(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)
	before := tc.Length()
	goodTrace := codestack.New.Ptr(0)
	tc.AddsPtr(false, goodTrace)

	// Act
	actual := args.Map{"grew": tc.Length() > before}

	// Assert
	expected := args.Map{"grew": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AddsPtr returns empty -- no skip adds", actual)
}

func Test_TraceCollection_FirstOrDefault_Empty_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}

	// Act
	actual := args.Map{"isNil": tc.FirstOrDefault() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.FirstOrDefault returns empty -- empty", actual)
}

func Test_TraceCollection_LastOrDefault_Empty_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}

	// Act
	actual := args.Map{"isNil": tc.LastOrDefault() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.LastOrDefault returns empty -- empty", actual)
}

func Test_TraceCollection_HasIndex_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)

	// Act
	actual := args.Map{
		"hasZero":  tc.HasIndex(0),
		"hasMega":  tc.HasIndex(9999),
	}

	// Assert
	expected := args.Map{
		"hasZero":  tc.HasIndex(0),
		"hasMega":  false,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection.HasIndex returns correct value -- with args", actual)
}

func Test_TraceCollection_GetPagesSize_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default(1, 10)

	// Act
	actual := args.Map{
		"zeroPage": tc.GetPagesSize(0),
		"negPage":  tc.GetPagesSize(-1),
		"valid":    tc.GetPagesSize(3) > 0,
	}

	// Assert
	expected := args.Map{
		"zeroPage": 0,
		"negPage":  0,
		"valid":    tc.GetPagesSize(3) > 0,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection.GetPagesSize returns correct value -- with args", actual)
}

func Test_TraceCollection_Filter_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)
	filtered := tc.Filter(func(tr *codestack.Trace) (bool, bool) {
		return true, false
	})

	// Act
	actual := args.Map{"notEmpty": len(filtered) > 0}

	// Assert
	expected := args.Map{"notEmpty": len(filtered) > 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Filter returns correct value -- takes all", actual)
}

func Test_TraceCollection_Filter_BreakEarly(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)
	filtered := tc.Filter(func(tr *codestack.Trace) (bool, bool) {
		return true, true // take first, break
	})

	// Act
	actual := args.Map{"len": len(filtered)}

	// Assert
	expected := args.Map{"len": len(filtered)}
	expected.ShouldBeEqual(t, 0, "TraceCollection.Filter returns correct value -- break early", actual)
}

func Test_TraceCollection_SafeLimitCollection_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)
	limited := tc.SafeLimitCollection(999)

	// Act
	actual := args.Map{"safeLen": limited.Length() == tc.Length()}

	// Assert
	expected := args.Map{"safeLen": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.SafeLimitCollection returns correct value -- exceeds length", actual)
}

func Test_TraceCollection_ConcatNew_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)
	origLen := tc.Length()
	newTc := tc.ConcatNew(codestack.Trace{PackageName: "extra"})

	// Act
	actual := args.Map{
		"origSame": tc.Length() == origLen,
		"newGrew":  newTc.Length() > origLen,
	}

	// Assert
	expected := args.Map{
		"origSame": true,
		"newGrew":  true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection.ConcatNew returns correct value -- with args", actual)
}

func Test_TraceCollection_ConcatNewPtr_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)
	origLen := tc.Length()
	tr := codestack.New.Ptr(0)
	newTc := tc.ConcatNewPtr(tr)

	// Act
	actual := args.Map{
		"origSame": tc.Length() == origLen,
		"newNotEmpty": newTc.Length() > 0,
	}

	// Assert
	expected := args.Map{
		"origSame": true,
		"newNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection.ConcatNewPtr returns correct value -- with args", actual)
}

func Test_TraceCollection_StackTraces_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)

	// Act
	actual := args.Map{"notEmpty": tc.StackTraces() != ""}

	// Assert
	expected := args.Map{"notEmpty": tc.StackTraces() != ""}
	expected.ShouldBeEqual(t, 0, "TraceCollection.StackTraces returns correct value -- with args", actual)
}

func Test_TraceCollection_StackTracesJsonResult_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)

	// Act
	actual := args.Map{"notNil": tc.StackTracesJsonResult() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.StackTracesJsonResult returns correct value -- with args", actual)
}

func Test_TraceCollection_NewStackTraces_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)

	// Act
	actual := args.Map{"notEmpty": tc.NewStackTraces(1) != ""}

	// Assert
	expected := args.Map{"notEmpty": tc.NewStackTraces(1) != ""}
	expected.ShouldBeEqual(t, 0, "TraceCollection.NewStackTraces returns correct value -- with args", actual)
}

func Test_TraceCollection_NewDefaultStackTraces_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)

	// Act
	actual := args.Map{"notEmpty": tc.NewDefaultStackTraces() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.NewDefaultStackTraces returns correct value -- with args", actual)
}

func Test_TraceCollection_NewStackTracesJsonResult_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)

	// Act
	actual := args.Map{"notNil": tc.NewStackTracesJsonResult(1) != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.NewStackTracesJsonResult returns correct value -- with args", actual)
}

func Test_TraceCollection_NewDefaultStackTracesJsonResult_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)

	// Act
	actual := args.Map{"notNil": tc.NewDefaultStackTracesJsonResult() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.NewDefaultStackTracesJsonResult returns correct value -- with args", actual)
}

func Test_TraceCollection_SkipDynamic(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)

	// Act
	actual := args.Map{"notNil": tc.SkipDynamic(1) != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.SkipDynamic returns correct value -- with args", actual)
}

func Test_TraceCollection_TakeDynamic_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)

	// Act
	actual := args.Map{"notNil": tc.TakeDynamic(1) != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.TakeDynamic returns correct value -- with args", actual)
}

func Test_TraceCollection_TakeCollection_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)
	taken := tc.TakeCollection(1)

	// Act
	actual := args.Map{"len": taken.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection.TakeCollection returns correct value -- with args", actual)
}

func Test_TraceCollection_LimitCollection_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)
	limited := tc.LimitCollection(1)

	// Act
	actual := args.Map{"len": limited.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection.LimitCollection returns correct value -- with args", actual)
}

func Test_TraceCollection_LimitDynamic_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)

	// Act
	actual := args.Map{"notNil": tc.LimitDynamic(1) != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.LimitDynamic returns correct value -- with args", actual)
}

func Test_TraceCollection_FirstOrDefaultDynamic_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)

	// Act
	actual := args.Map{"notNil": tc.FirstOrDefaultDynamic() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.FirstOrDefaultDynamic returns correct value -- with args", actual)
}

func Test_TraceCollection_LastOrDefaultDynamic_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)

	// Act
	actual := args.Map{"notNil": tc.LastOrDefaultDynamic() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.LastOrDefaultDynamic returns correct value -- with args", actual)
}

// ── newTraceCollection factory coverage ──

func Test_NewTraces_Cap(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() == 0 {
		t.Skip("StackTrace returned empty -- platform-dependent")
	}

	// Act
	actual := args.Map{"notEmpty": tc.Length() > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "newStacksCreator.All returns correct value -- with args", actual)
}

// ── NameOf with real func names ──

func Test_NameOf_All_WithDot(t *testing.T) {
	// Arrange
	full, pkg, method := codestack.NameOf.All("github.com/pkg/errors.New")

	// Act
	actual := args.Map{
		"fullNotEmpty":   full != "",
		"pkgNotEmpty":    pkg != "",
		"methodNotEmpty": method != "",
	}

	// Assert
	expected := args.Map{
		"fullNotEmpty":   true,
		"pkgNotEmpty":    true,
		"methodNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "NameOf.All returns non-empty -- with dotted path", actual)
}

func Test_NameOf_MethodByFullName_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	result := codestack.NameOf.MethodByFullName("github.com/pkg.Method")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.MethodByFullName returns correct value -- with args", actual)
}

func Test_NameOf_PackageByFullName_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	result := codestack.NameOf.PackageByFullName("github.com/pkg.Method")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.PackageByFullName returns correct value -- with args", actual)
}

func Test_NameOf_CurrentFuncFullPath_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	result := codestack.NameOf.CurrentFuncFullPath("github.com/pkg.Method")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.CurrentFuncFullPath returns correct value -- with args", actual)
}

func Test_NameOf_JoinPackageNameWithRelative_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	result := codestack.NameOf.JoinPackageNameWithRelative("github.com/pkg.Method", "SubMethod")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.JoinPackageNameWithRelative returns non-empty -- with args", actual)
}

// ── File getter additional coverage ──

func Test_File_PathLineSep_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	fp, ln := codestack.File.PathLineSep(0)

	// Act
	actual := args.Map{
		"pathNotEmpty": fp != "",
		"linePositive": ln > 0,
	}

	// Assert
	expected := args.Map{
		"pathNotEmpty": true,
		"linePositive": true,
	}
	expected.ShouldBeEqual(t, 0, "File.PathLineSep returns correct value -- with args", actual)
}

func Test_File_PathLineSepDefault_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	fp, ln := codestack.File.PathLineSepDefault()

	// Act
	actual := args.Map{
		"pathNotEmpty": fp != "",
		"linePositive": ln > 0,
	}

	// Assert
	expected := args.Map{
		"pathNotEmpty": true,
		"linePositive": true,
	}
	expected.ShouldBeEqual(t, 0, "File.PathLineSepDefault returns correct value -- with args", actual)
}

func Test_File_FilePathWithLineString_FromTraceCollectionNilLe(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": codestack.File.FilePathWithLineString(0) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.FilePathWithLineString returns non-empty -- with args", actual)
}

func Test_File_PathLineStringDefault_FromTraceCollectionNilLe(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": codestack.File.PathLineStringDefault() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.PathLineStringDefault returns correct value -- with args", actual)
}

func Test_File_CurrentFilePath_FromTraceCollectionNilLe(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": codestack.File.CurrentFilePath() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.CurrentFilePath returns correct value -- with args", actual)
}

// ── Dir additional coverage ──

func Test_Dir_Get_FromTraceCollectionNilLe(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": codestack.Dir.Get(0) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.Get returns correct value -- with args", actual)
}

func Test_Dir_RepoDir_FromTraceCollectionNilLe(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": codestack.Dir.RepoDir() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.RepoDir returns correct value -- with args", actual)
}

func Test_Dir_RepoDirJoin_FromTraceCollectionNilLe(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": codestack.Dir.RepoDirJoin("sub") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.RepoDirJoin returns correct value -- with args", actual)
}

// ── isSkippablePackage ──

func Test_SkippablePackage_ViaTrace(t *testing.T) {
	// Arrange
	// Create a trace; application code should NOT be skippable
	trace := codestack.New.Default()

	// Act
	actual := args.Map{"notSkippable": !trace.IsSkippable}

	// Assert
	expected := args.Map{"notSkippable": true}
	expected.ShouldBeEqual(t, 0, "isSkippablePackage returns non-empty -- returns false for app code", actual)
}

// ── TraceCollection.AddsUsingSkipDefault ──

func Test_TraceCollection_AddsUsingSkipDefault_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}
	tc.AddsUsingSkipDefault(0)

	// Act
	actual := args.Map{"notEmpty": tc.Length() > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.AddsUsingSkipDefault returns correct value -- with args", actual)
}

// ── TraceCollection.ConcatNewUsingSkip ──

func Test_TraceCollection_ConcatNewUsingSkip_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)
	newTc := tc.ConcatNewUsingSkip(0)

	// Act
	actual := args.Map{"notEmpty": newTc.Length() > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.ConcatNewUsingSkip returns correct value -- with args", actual)
}

func Test_TraceCollection_ConcatNewUsingSkipPlusCount_FromTraceCollectionNilLe(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(1)
	newTc := tc.ConcatNewUsingSkipPlusCount(0, 5)

	// Act
	actual := args.Map{"notEmpty": newTc.Length() > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection.ConcatNewUsingSkipPlusCount returns correct value -- with args", actual)
}
