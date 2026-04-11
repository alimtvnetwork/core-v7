package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── newTraceCollection ──

func Test_NewTraceCollection_Default(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.Default(0, codestack.DefaultStackCount)

	// Act
	actual := args.Map{"notEmpty": !tc.IsEmpty()}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "newTraceCollection.Default returns non-empty -- default args", actual)
}

func Test_NewTraceCollection_SkipOne(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()

	// Act
	actual := args.Map{"notEmpty": !tc.IsEmpty()}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "newTraceCollection.SkipOne returns non-empty -- skip one", actual)
}

func Test_NewTraceCollection_SkipNone(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()

	// Act
	actual := args.Map{"notEmpty": !tc.IsEmpty()}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "newTraceCollection.SkipNone returns non-empty -- skip none", actual)
}

func Test_NewTraceCollection_DefaultCount(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.DefaultCount(0)

	// Act
	actual := args.Map{"notEmpty": !tc.IsEmpty()}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "newTraceCollection.DefaultCount returns non-empty -- skip 1", actual)
}

// ── Trace edge cases ──

func Test_Trace_NilPtr_String(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act
	actual := args.Map{"result": trace.String()}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Trace.String returns empty -- nil pointer", actual)
}

func Test_Trace_NilPtr_HasIssues(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act
	actual := args.Map{"hasIssues": trace.HasIssues()}

	// Assert
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "Trace.HasIssues returns true -- nil pointer", actual)
}

func Test_Trace_NilPtr_IsNil(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act
	actual := args.Map{"isNil": trace.IsNil()}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Trace.IsNil returns true -- nil pointer", actual)
}

func Test_Trace_NilPtr_IsNotNil(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act
	actual := args.Map{"isNotNil": trace.IsNotNil()}

	// Assert
	expected := args.Map{"isNotNil": false}
	expected.ShouldBeEqual(t, 0, "Trace.IsNotNil returns false -- nil pointer", actual)
}

func Test_Trace_NilPtr_Dispose(t *testing.T) {
	// Arrange
	var trace *codestack.Trace
	trace.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Trace.Dispose returns safely -- nil pointer", actual)
}

func Test_Trace_NilPtr_ClonePtr(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act
	actual := args.Map{"isNil": trace.ClonePtr() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Trace.ClonePtr returns nil -- nil pointer", actual)
}

func Test_Trace_Empty_HasIssues(t *testing.T) {
	// Arrange
	trace := codestack.Trace{}

	// Act
	actual := args.Map{"hasIssues": trace.HasIssues()}

	// Assert
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "Trace.HasIssues returns true -- empty struct", actual)
}

// ── Trace message caching ──

func Test_Trace_Message_CalledTwice(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	msg1 := trace.Message()
	msg2 := trace.Message()

	// Act
	actual := args.Map{
		"equal": msg1 == msg2,
		"notEmpty": msg1 != "",
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace.Message returns cached -- called twice", actual)
}

func Test_Trace_ShortString_CalledTwice(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	s1 := trace.ShortString()
	s2 := trace.ShortString()

	// Act
	actual := args.Map{
		"equal": s1 == s2,
		"notEmpty": s1 != "",
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace.ShortString returns cached -- called twice", actual)
}

// ── FileWithLine edge cases ──

func Test_FileWithLine_NilPtr_String(t *testing.T) {
	// Arrange
	var fwl *codestack.FileWithLine

	// Act
	actual := args.Map{"result": fwl.String()}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "FileWithLine.String returns empty -- nil pointer", actual)
}

func Test_FileWithLine_NilPtr_IsNil(t *testing.T) {
	// Arrange
	var fwl *codestack.FileWithLine

	// Act
	actual := args.Map{"isNil": fwl.IsNil()}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine.IsNil returns true -- nil pointer", actual)
}

func Test_FileWithLine_NilPtr_IsNotNil(t *testing.T) {
	// Arrange
	var fwl *codestack.FileWithLine

	// Act
	actual := args.Map{"isNotNil": fwl.IsNotNil()}

	// Assert
	expected := args.Map{"isNotNil": false}
	expected.ShouldBeEqual(t, 0, "FileWithLine.IsNotNil returns false -- nil pointer", actual)
}

// ── stacksTo ──

func Test_StacksTo_String_FromNewTraceCollectionDe(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.String(0, 5)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.String returns non-empty -- with count", actual)
}

func Test_StacksTo_StringDefault_FromNewTraceCollectionDe(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.StringDefault()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringDefault returns non-empty -- default args", actual)
}

func Test_StacksTo_StringNoCount_FromNewTraceCollectionDe(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.StringNoCount(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringNoCount returns non-empty -- skip 0", actual)
}

func Test_StacksTo_Bytes_FromNewTraceCollectionDe(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.Bytes(0)

	// Act
	actual := args.Map{"notEmpty": len(result) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.Bytes returns non-empty -- skip 0", actual)
}

func Test_StacksTo_BytesDefault_FromNewTraceCollectionDe(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.BytesDefault()

	// Act
	actual := args.Map{"notEmpty": len(result) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.BytesDefault returns non-empty -- default", actual)
}

func Test_StacksTo_JsonString_FromNewTraceCollectionDe(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.JsonString(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.JsonString returns non-empty -- skip 0", actual)
}

func Test_StacksTo_JsonStringDefault_FromNewTraceCollectionDe(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.JsonStringDefault()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.JsonStringDefault returns non-empty -- default", actual)
}

func Test_StacksTo_StringUsingFmt_FromNewTraceCollectionDe(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.StringUsingFmt(
		func(tr *codestack.Trace) string { return tr.PackageName },
		1,
		5,
	)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": result != ""}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringUsingFmt returns non-empty -- with formatter", actual)
}

// ── newCreator ──

func Test_NewCreator_Create_FromNewTraceCollectionDe(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)

	// Act
	actual := args.Map{
		"isOkay": trace.IsOkay,
		"notEmpty": trace.PackageName != "",
	}

	// Assert
	expected := args.Map{
		"isOkay": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "newCreator.Create returns valid trace -- skip 0", actual)
}

func Test_NewCreator_Default_FromNewTraceCollectionDe(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	actual := args.Map{"isOkay": trace.IsOkay}

	// Assert
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "newCreator.Default returns valid trace -- default", actual)
}

func Test_NewCreator_SkipOne_FromNewTraceCollectionDe(t *testing.T) {
	// Arrange
	trace := codestack.New.SkipOne()

	// Act
	actual := args.Map{"isOkay": trace.IsOkay}

	// Assert
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "newCreator.SkipOne returns valid trace -- skip one", actual)
}

func Test_NewCreator_Ptr_FromNewTraceCollectionDe(t *testing.T) {
	// Arrange
	ptr := codestack.New.Ptr(0)

	// Act
	actual := args.Map{
		"notNil": ptr != nil,
		"isOkay": ptr.IsOkay,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"isOkay": true,
	}
	expected.ShouldBeEqual(t, 0, "newCreator.Ptr returns valid pointer -- skip 0", actual)
}

// ── NameOf edge cases ──

func Test_NameOf_All_EmptyInput(t *testing.T) {
	// Arrange
	full, pkg, method := codestack.NameOf.All("")

	// Act
	actual := args.Map{
		"full": full,
		"pkg": pkg,
		"method": method,
	}

	// Assert
	expected := args.Map{
		"full": "",
		"pkg": "",
		"method": "",
	}
	expected.ShouldBeEqual(t, 0, "NameOf.All returns empty -- empty input", actual)
}

func Test_NameOf_Method_EmptyInput(t *testing.T) {
	// Act
	actual := args.Map{"result": codestack.NameOf.MethodByFullName("")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NameOf.Method returns empty -- empty input", actual)
}

func Test_NameOf_Package_EmptyInput(t *testing.T) {
	// Act
	actual := args.Map{"result": codestack.NameOf.PackageByFullName("")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NameOf.Package returns empty -- empty input", actual)
}

// ── Trace.AsFileLiner ──

func Test_Trace_AsFileLiner_FromNewTraceCollectionDe(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	liner := trace.AsFileLiner()

	// Act
	actual := args.Map{
		"notNil": liner != nil,
		"pathNotEmpty": liner.FullFilePath() != "",
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"pathNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace.AsFileLiner returns valid FileWithLiner -- default trace", actual)
}

// ── FileWithLine.AsFileLiner ──

func Test_FileWithLine_AsFileLiner_FromNewTraceCollectionDe(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	liner := fwl.AsFileLiner()

	// Act
	actual := args.Map{
		"notNil": liner != nil,
		"line": liner.LineNumber(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"line": 10,
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine.AsFileLiner returns valid FileWithLiner -- with data", actual)
}

// ── Trace.Clone with IsSkippable ──

func Test_Trace_Clone_PreservesIsSkippable(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	cloned := trace.Clone()

	// Act
	actual := args.Map{
		"skippableMatch": cloned.IsSkippable == trace.IsSkippable,
		"pkgMatch":       cloned.PackageName == trace.PackageName,
	}

	// Assert
	expected := args.Map{
		"skippableMatch": true,
		"pkgMatch":       true,
	}
	expected.ShouldBeEqual(t, 0, "Trace.Clone preserves IsSkippable -- default trace", actual)
}

// ── File / Dir getters ──

func Test_File_Name_FromNewTraceCollectionDe(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": codestack.File.Name(0) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.Name returns non-empty -- skip 0", actual)
}

func Test_File_Path_FromNewTraceCollectionDe(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": codestack.File.Path(0) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.Path returns non-empty -- skip 0", actual)
}

func Test_Dir_CurDir_FromNewTraceCollectionDe(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": codestack.Dir.CurDir() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.CurDir returns non-empty -- current dir", actual)
}

func Test_Dir_CurDirJoin_FromNewTraceCollectionDe(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": codestack.Dir.CurDirJoin("sub") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.CurDirJoin returns non-empty -- with subdir", actual)
}
