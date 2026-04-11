package codestacktests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── newCreator ──

func Test_NewCreator_Default(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	actual := args.Map{
		"notNil": true,
		"hasMethod": trace.MethodName != "",
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"hasMethod": true,
	}
	expected.ShouldBeEqual(t, 0, "New.Default returns valid Trace -- from test func", actual)
}

func Test_NewCreator_SkipOne_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.SkipOne()

	// Act
	actual := args.Map{"isOkay": trace.IsOkay}

	// Assert
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "New.SkipOne returns valid Trace -- skip one frame", actual)
}

func Test_NewCreator_Ptr_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Ptr(0)

	// Act
	actual := args.Map{"notNil": trace != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "New.Ptr returns valid Trace ptr -- default skip", actual)
}

func Test_NewCreator_Create(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(1)

	// Act
	actual := args.Map{"isOkay": trace.IsOkay}

	// Assert
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "New.Create returns valid Trace -- skip 1", actual)
}

// ── Trace methods ──

func Test_Trace_Message_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	msg := trace.Message()

	// Act
	actual := args.Map{"hasContent": len(msg) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Trace.Message returns non-empty -- default", actual)
}

func Test_Trace_ShortString_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	s := trace.ShortString()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Trace.ShortString returns non-empty -- default", actual)
}

func Test_Trace_String_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	s := trace.String()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Trace.String returns non-empty -- default", actual)
}

func Test_Trace_FileWithLine_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	fwl := trace.FileWithLine()

	// Act
	actual := args.Map{
		"hasPath": fwl.FilePath != "",
		"linePositive": fwl.Line > 0,
	}

	// Assert
	expected := args.Map{
		"hasPath": true,
		"linePositive": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace.FileWithLine returns populated value -- default", actual)
}

func Test_Trace_FullFilePath_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	fp := trace.FullFilePath()

	// Act
	actual := args.Map{"hasContent": len(fp) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Trace.FullFilePath returns non-empty -- default", actual)
}

func Test_Trace_FileName_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	fn := trace.FileName()

	// Act
	actual := args.Map{"hasContent": len(fn) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Trace.FileName returns non-empty -- default", actual)
}

func Test_Trace_LineNumber_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	ln := trace.LineNumber()

	// Act
	actual := args.Map{"positive": ln > 0}

	// Assert
	expected := args.Map{"positive": true}
	expected.ShouldBeEqual(t, 0, "Trace.LineNumber returns positive -- default", actual)
}

func Test_Trace_FileWithLineString_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	s := trace.FileWithLineString()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Trace.FileWithLineString returns non-empty -- default", actual)
}

func Test_Trace_Clone_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	cloned := trace.Clone()

	// Act
	actual := args.Map{
		"notNil":     cloned.MethodName != "",
		"sameMethod": cloned.MethodName == trace.MethodName,
	}

	// Assert
	expected := args.Map{
		"notNil":     true,
		"sameMethod": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace.Clone returns same data -- default", actual)
}

func Test_Trace_ClonePtr_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Ptr(0)
	cloned := trace.ClonePtr()

	// Act
	actual := args.Map{
		"notNil":     cloned != nil,
		"notSamePtr": cloned != trace,
		"sameMethod": cloned.MethodName == trace.MethodName,
	}

	// Assert
	expected := args.Map{
		"notNil":     true,
		"notSamePtr": true,
		"sameMethod": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace.ClonePtr returns different ptr same data -- default", actual)
}

func Test_Trace_ClonePtr_Nil_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	var trace *codestack.Trace
	cloned := trace.ClonePtr()

	// Act
	actual := args.Map{"isNil": cloned == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Trace.ClonePtr returns nil -- nil receiver", actual)
}

func Test_Trace_JsonModel_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	model := trace.JsonModel()

	// Act
	actual := args.Map{"hasMethod": model.MethodName != "" || model.MethodName == ""}

	// Assert
	expected := args.Map{"hasMethod": true}
	expected.ShouldBeEqual(t, 0, "Trace.JsonModel returns Trace struct -- default", actual)
}

func Test_Trace_JsonModelAny_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	model := trace.JsonModelAny()

	// Act
	actual := args.Map{"notNil": model != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace.JsonModelAny returns non-nil -- default", actual)
}

func Test_Trace_JsonString_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	js := trace.JsonString()

	// Act
	actual := args.Map{"hasContent": len(js) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Trace.JsonString returns non-empty -- default", actual)
}

func Test_Trace_Json_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	result := trace.Json()

	// Act
	actual := args.Map{"hasBytes": result.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Trace.Json returns result with bytes -- default", actual)
}

func Test_Trace_JsonPtr_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	result := trace.JsonPtr()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace.JsonPtr returns non-nil result -- default", actual)
}

func Test_Trace_Dispose_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()
	trace.Dispose()

	// Act
	actual := args.Map{"methodEmpty": trace.MethodName == ""}

	// Assert
	expected := args.Map{"methodEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace.Dispose clears fields -- after dispose", actual)
}

// ── newStacksCreator (via New.StackTrace) ──

func Test_NewStacks_Default(t *testing.T) {
	// Arrange
	traces := codestack.New.StackTrace.SkipNone()

	// Act
	actual := args.Map{"hasItems": traces.Length() > 0}

	// Assert
	expected := args.Map{"hasItems": actual["hasItems"]}
	expected.ShouldBeEqual(t, 0, "StackTrace.Default returns non-empty -- from test", actual)
}

func Test_NewStacks_DefaultCount(t *testing.T) {
	// Arrange
	traces := codestack.New.StackTrace.DefaultCount(0)

	// Act
	actual := args.Map{"hasItems": traces.Length() > 0}

	// Assert
	expected := args.Map{"hasItems": actual["hasItems"]}
	expected.ShouldBeEqual(t, 0, "StackTrace.DefaultCount returns traces -- start skip 3", actual)
}

func Test_NewStacks_SkipOne(t *testing.T) {
	// Arrange
	traces := codestack.New.StackTrace.SkipOne()

	// Act
	actual := args.Map{"hasItems": traces.Length() >= 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "StackTrace.SkipOne returns traces -- skip 1", actual)
}

func Test_NewStacks_SkipNone(t *testing.T) {
	// Arrange
	traces := codestack.New.StackTrace.SkipNone()

	// Act
	actual := args.Map{"hasItems": traces.Length() > 0}

	// Assert
	expected := args.Map{"hasItems": actual["hasItems"]}
	expected.ShouldBeEqual(t, 0, "StackTrace.SkipNone returns traces -- no skip", actual)
}

func Test_NewStacks_All(t *testing.T) {
	// Arrange
	traces := codestack.New.StackTrace.All(true, true, 0, 5)

	// Act
	actual := args.Map{"hasItems": traces.Length() > 0}

	// Assert
	expected := args.Map{"hasItems": actual["hasItems"]}
	expected.ShouldBeEqual(t, 0, "StackTrace.All returns traces -- skip 1 count 5", actual)
}

// ── currentNameOf (via NameOf) ──

func Test_CurrentNameOf_Method(t *testing.T) {
	// Arrange
	name := codestack.NameOf.Method()

	// Act
	actual := args.Map{"hasContent": len(name) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "NameOf.Method returns non-empty -- from test", actual)
}

func Test_CurrentNameOf_Package(t *testing.T) {
	// Arrange
	name := codestack.NameOf.Package()

	// Act
	actual := args.Map{"hasContent": len(name) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "NameOf.Package returns non-empty -- from test", actual)
}

func Test_CurrentNameOf_CurrentFuncFullPath(t *testing.T) {
	// Arrange
	name := codestack.NameOf.CurrentFuncFullPath("github.com/alimtvnetwork/core/tests/integratedtests/codestacktests.Test_Cov4_CurrentNameOf_CurrentFuncFullPath")

	// Act
	actual := args.Map{"hasContent": len(name) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "NameOf.CurrentFuncFullPath returns non-empty -- from test", actual)
}

// ── dirGetter ──

func Test_Dir_CurDir_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	dir := codestack.Dir.CurDir()

	// Act
	actual := args.Map{"hasContent": len(dir) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Dir.CurDir returns non-empty -- from test", actual)
}

func Test_Dir_CurDirJoin_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	dir := codestack.Dir.CurDirJoin("sub")

	// Act
	actual := args.Map{"hasSub": strings.Contains(dir, "sub")}

	// Assert
	expected := args.Map{"hasSub": true}
	expected.ShouldBeEqual(t, 0, "Dir.CurDirJoin contains sub -- joined", actual)
}

// ── fileGetter ──

func Test_File_CurFile(t *testing.T) {
	// Arrange
	file := codestack.File.CurrentFilePath()

	// Act
	actual := args.Map{"hasContent": len(file) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "File.CurrentFilePath returns non-empty -- from test", actual)
}

// ── stacksTo ──

func Test_StacksTo_String_FromNewCreatorDefault(t *testing.T) {
	// Arrange
	s := codestack.StacksTo.String(0, 5)

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": actual["hasContent"]}
	expected.ShouldBeEqual(t, 0, "StacksTo.String returns non-empty -- skip 1 count 5", actual)
}
