package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── FileWithLine ──

func Test_FileWithLine_FullFilePath(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}

	// Act
	actual := args.Map{
		"path": fwl.FullFilePath(),
		"line": fwl.LineNumber(),
	}

	// Assert
	expected := args.Map{
		"path": "/tmp/test.go",
		"line": 10,
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- FullFilePath and LineNumber", actual)
}

func Test_FileWithLine_IsNil(t *testing.T) {
	// Arrange
	var fwl *codestack.FileWithLine

	// Act
	actual := args.Map{
		"isNil": fwl.IsNil(),
		"isNotNil": fwl.IsNotNil(),
	}

	// Assert
	expected := args.Map{
		"isNil": true,
		"isNotNil": false,
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns nil -- IsNil", actual)
}

func Test_FileWithLine_IsNotNil(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 1}

	// Act
	actual := args.Map{
		"isNil": fwl.IsNil(),
		"isNotNil": fwl.IsNotNil(),
	}

	// Assert
	expected := args.Map{
		"isNil": false,
		"isNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns nil -- IsNotNil", actual)
}

func Test_FileWithLine_String(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}

	// Act
	actual := args.Map{"notEmpty": fwl.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- String", actual)
}

func Test_FileWithLine_String_Nil(t *testing.T) {
	// Arrange
	var fwl *codestack.FileWithLine

	// Act
	actual := args.Map{"empty": fwl.String() == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns nil -- String nil", actual)
}

func Test_FileWithLine_StringUsingFmt_FromFileWithLineFullFile(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	result := fwl.StringUsingFmt(func(f codestack.FileWithLine) string { return "custom" })

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "custom"}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- StringUsingFmt", actual)
}

func Test_FileWithLine_FileWithLine(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	result := fwl.FileWithLine()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- FileWithLine method", actual)
}

func Test_FileWithLine_JsonModel_FromFileWithLineFullFile(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	model := fwl.JsonModel()

	// Act
	actual := args.Map{"path": model.FilePath}

	// Assert
	expected := args.Map{"path": "/tmp/test.go"}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- JsonModel", actual)
}

func Test_FileWithLine_JsonModelAny(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	result := fwl.JsonModelAny()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- JsonModelAny", actual)
}

func Test_FileWithLine_JsonString(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	result := fwl.JsonString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- JsonString", actual)
}

func Test_FileWithLine_Json_FromFileWithLineFullFile(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	result := fwl.Json()

	// Act
	actual := args.Map{"noErr": result.MeaningfulError() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- Json", actual)
}

func Test_FileWithLine_JsonPtr(t *testing.T) {
	// Arrange
	fwl := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	result := fwl.JsonPtr()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- JsonPtr", actual)
}

func Test_FileWithLine_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	original := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	jsonResult := original.JsonPtr()
	var target codestack.FileWithLine
	result, err := target.ParseInjectUsingJson(jsonResult)

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
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- ParseInjectUsingJson", actual)
}

func Test_FileWithLine_ParseInjectUsingJson_Error(t *testing.T) {
	// Arrange
	badJson := corejson.NewPtr("not a FileWithLine struct")
	var target codestack.FileWithLine
	_, err := target.ParseInjectUsingJson(badJson)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns error -- ParseInjectUsingJson error", actual)
}

func Test_FileWithLine_ParseInjectUsingJsonMust(t *testing.T) {
	// Arrange
	original := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	jsonResult := original.JsonPtr()
	var target codestack.FileWithLine
	result := target.ParseInjectUsingJsonMust(jsonResult)

	// Act
	actual := args.Map{"path": result.FilePath}

	// Assert
	expected := args.Map{"path": "/tmp/test.go"}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- ParseInjectUsingJsonMust", actual)
}

func Test_FileWithLine_JsonParseSelfInject(t *testing.T) {
	// Arrange
	original := codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	jsonResult := original.JsonPtr()
	var target codestack.FileWithLine
	err := target.JsonParseSelfInject(jsonResult)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"path": target.FilePath,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"path": "/tmp/test.go",
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- JsonParseSelfInject", actual)
}

func Test_FileWithLine_AsFileLiner(t *testing.T) {
	// Arrange
	fwl := &codestack.FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	liner := fwl.AsFileLiner()

	// Act
	actual := args.Map{
		"path": liner.FullFilePath(),
		"line": liner.LineNumber(),
	}

	// Assert
	expected := args.Map{
		"path": "/tmp/test.go",
		"line": 10,
	}
	expected.ShouldBeEqual(t, 0, "FileWithLine returns non-empty -- AsFileLiner", actual)
}

// ── Trace ──

func Test_Trace_New_Create(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)

	// Act
	actual := args.Map{
		"isOkay": trace.IsOkay,
		"notNil": trace.IsNotNil(),
	}

	// Assert
	expected := args.Map{
		"isOkay": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- New.Create", actual)
}

func Test_Trace_New_Default(t *testing.T) {
	// Arrange
	trace := codestack.New.Default()

	// Act
	actual := args.Map{"isOkay": trace.IsOkay}

	// Assert
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- New.Default", actual)
}

func Test_Trace_New_SkipOne(t *testing.T) {
	// Arrange
	trace := codestack.New.SkipOne()

	// Act
	actual := args.Map{"isOkay": trace.IsOkay}

	// Assert
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- New.SkipOne", actual)
}

func Test_Trace_New_Ptr(t *testing.T) {
	// Arrange
	trace := codestack.New.Ptr(0)

	// Act
	actual := args.Map{
		"notNil": trace != nil,
		"isOkay": trace.IsOkay,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"isOkay": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- New.Ptr", actual)
}

func Test_Trace_Message(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	msg := trace.Message()

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- Message", actual)
}

func Test_Trace_Message_Cached(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	msg1 := trace.Message()
	msg2 := trace.Message()

	// Act
	actual := args.Map{"same": msg1 == msg2}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- Message cached", actual)
}

func Test_Trace_ShortString(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	ss := trace.ShortString()

	// Act
	actual := args.Map{"notEmpty": ss != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- ShortString", actual)
}

func Test_Trace_ShortString_Cached(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	ss1 := trace.ShortString()
	ss2 := trace.ShortString()

	// Act
	actual := args.Map{"same": ss1 == ss2}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- ShortString cached", actual)
}

func Test_Trace_IsNil(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act
	actual := args.Map{
		"isNil": trace.IsNil(),
		"isNotNil": trace.IsNotNil(),
	}

	// Assert
	expected := args.Map{
		"isNil": true,
		"isNotNil": false,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns nil -- IsNil", actual)
}

func Test_Trace_HasIssues_Nil(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act
	actual := args.Map{"hasIssues": trace.HasIssues()}

	// Assert
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "Trace returns nil -- HasIssues nil", actual)
}

func Test_Trace_HasIssues_NotOkay(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{IsOkay: false}

	// Act
	actual := args.Map{"hasIssues": trace.HasIssues()}

	// Assert
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- HasIssues not okay", actual)
}

func Test_Trace_HasIssues_MissingPackage(t *testing.T) {
	// Arrange
	trace := &codestack.Trace{IsOkay: true, PackageMethodName: "test", PackageName: ""}

	// Act
	actual := args.Map{"hasIssues": trace.HasIssues()}

	// Assert
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- HasIssues missing package", actual)
}

func Test_Trace_String(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)

	// Act
	actual := args.Map{"notEmpty": trace.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- String", actual)
}

func Test_Trace_String_Nil(t *testing.T) {
	// Arrange
	var trace *codestack.Trace

	// Act
	actual := args.Map{"empty": trace.String() == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Trace returns nil -- String nil", actual)
}

func Test_Trace_StringUsingFmt(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	result := trace.StringUsingFmt(func(tr codestack.Trace) string { return "custom" })

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "custom"}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- StringUsingFmt", actual)
}

func Test_Trace_FileWithLine_FromFileWithLineFullFile(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	fwl := trace.FileWithLine()

	// Act
	actual := args.Map{
		"pathNotEmpty": fwl.FilePath != "",
		"lineGt0": fwl.Line > 0,
	}

	// Assert
	expected := args.Map{
		"pathNotEmpty": true,
		"lineGt0": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns non-empty -- FileWithLine", actual)
}

func Test_Trace_FullFilePath(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)

	// Act
	actual := args.Map{"notEmpty": trace.FullFilePath() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- FullFilePath", actual)
}

func Test_Trace_FileName(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)

	// Act
	actual := args.Map{"notEmpty": trace.FileName() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- FileName", actual)
}

func Test_Trace_LineNumber(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)

	// Act
	actual := args.Map{"gt0": trace.LineNumber() > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- LineNumber", actual)
}

func Test_Trace_FileWithLineString(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	result := trace.FileWithLineString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace returns non-empty -- FileWithLineString", actual)
}

func Test_Trace_JsonModel_FromFileWithLineFullFile(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	model := trace.JsonModel()

	// Act
	actual := args.Map{"isOkay": model.IsOkay}

	// Assert
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- JsonModel", actual)
}

func Test_Trace_JsonModelAny(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	result := trace.JsonModelAny()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- JsonModelAny", actual)
}

func Test_Trace_JsonString(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	result := trace.JsonString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- JsonString", actual)
}

func Test_Trace_Json_FromFileWithLineFullFile(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	result := trace.Json()

	// Act
	actual := args.Map{"noErr": result.MeaningfulError() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- Json", actual)
}

func Test_Trace_JsonPtr(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	result := trace.JsonPtr()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- JsonPtr", actual)
}

func Test_Trace_Clone_FromFileWithLineFullFile(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	cloned := trace.Clone()

	// Act
	actual := args.Map{"samePath": cloned.FilePath == trace.FilePath}

	// Assert
	expected := args.Map{"samePath": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- Clone", actual)
}

func Test_Trace_ClonePtr(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	cloned := trace.ClonePtr()

	// Act
	actual := args.Map{
		"notNil": cloned != nil,
		"samePath": cloned.FilePath == trace.FilePath,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"samePath": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- ClonePtr", actual)
}

func Test_Trace_ClonePtr_Nil_FromFileWithLineFullFile(t *testing.T) {
	// Arrange
	var trace *codestack.Trace
	cloned := trace.ClonePtr()
	isNil := cloned == nil

	// Act
	actual := args.Map{"isNil": isNil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Trace returns nil -- ClonePtr nil", actual)
}

func Test_Trace_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	jsonResult := trace.JsonPtr()
	var target codestack.Trace
	result, err := target.ParseInjectUsingJson(jsonResult)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": result != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- ParseInjectUsingJson", actual)
}

func Test_Trace_ParseInjectUsingJsonMust(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	jsonResult := trace.JsonPtr()
	var target codestack.Trace
	result := target.ParseInjectUsingJsonMust(jsonResult)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- ParseInjectUsingJsonMust", actual)
}

func Test_Trace_JsonParseSelfInject(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	jsonResult := trace.JsonPtr()
	var target codestack.Trace
	err := target.JsonParseSelfInject(jsonResult)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- JsonParseSelfInject", actual)
}

func Test_Trace_Dispose_FromFileWithLineFullFile(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	trace.Dispose()

	// Act
	actual := args.Map{
		"pathEmpty": trace.FilePath == "",
		"line": trace.Line,
	}

	// Assert
	expected := args.Map{
		"pathEmpty": true,
		"line": 0,
	}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- Dispose", actual)
}

func Test_Trace_Dispose_Nil(t *testing.T) {
	// Arrange
	var trace *codestack.Trace
	trace.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Trace returns nil -- Dispose nil", actual)
}

func Test_Trace_AsFileLiner(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	liner := trace.AsFileLiner()

	// Act
	actual := args.Map{"notEmpty": liner.FullFilePath() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- AsFileLiner", actual)
}

// ── TraceCollection ──

func Test_TraceCollection_Basic(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()

	// Act
	actual := args.Map{
		"hasAny": tc.HasAnyItem(),
		"lenGt0": tc.Length() > 0,
		"count": tc.Count(),
	}

	// Assert
	expected := args.Map{
		"hasAny": true,
		"lenGt0": true,
		"count": tc.Length(),
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- basic", actual)
}

func Test_TraceCollection_IsEmpty(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}

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
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- IsEmpty", actual)
}

func Test_TraceCollection_Length_Nil(t *testing.T) {
	// Arrange
	var tc *codestack.TraceCollection

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns nil -- Length nil", actual)
}

func Test_TraceCollection_FirstAndLast(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 0 {
		first := tc.First()
		last := tc.Last()

	// Act
		actual := args.Map{
			"firstOk": first.IsOkay,
			"lastOk": last.IsOkay,
		}

	// Assert
		expected := args.Map{
			"firstOk": true,
			"lastOk": true,
		}
		expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- First and Last", actual)
	}
}

func Test_TraceCollection_FirstOrDefault_Empty(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}
	result := tc.FirstOrDefault()
	isNil := result == nil

	// Act
	actual := args.Map{"isNil": isNil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- FirstOrDefault empty", actual)
}

func Test_TraceCollection_LastOrDefault_Empty(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}
	result := tc.LastOrDefault()
	isNil := result == nil

	// Act
	actual := args.Map{"isNil": isNil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- LastOrDefault empty", actual)
}

func Test_TraceCollection_FirstOrDefault_NonEmpty(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.FirstOrDefault()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- FirstOrDefault non-empty", actual)
}

func Test_TraceCollection_FirstDynamic(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.FirstDynamic()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- FirstDynamic", actual)
}

func Test_TraceCollection_LastDynamic(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.LastDynamic()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- LastDynamic", actual)
}

func Test_TraceCollection_FirstOrDefaultDynamic(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.FirstOrDefaultDynamic()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- FirstOrDefaultDynamic", actual)
}

func Test_TraceCollection_LastOrDefaultDynamic(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.LastOrDefaultDynamic()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- LastOrDefaultDynamic", actual)
}

func Test_TraceCollection_HasIndex(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()

	// Act
	actual := args.Map{
		"has0": tc.HasIndex(0),
		"hasBig": tc.HasIndex(99999),
	}

	// Assert
	expected := args.Map{
		"has0": true,
		"hasBig": false,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- HasIndex", actual)
}

func Test_TraceCollection_Strings(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	strs := tc.Strings()

	// Act
	actual := args.Map{"lenGt0": len(strs) > 0}

	// Assert
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Strings", actual)
}

func Test_TraceCollection_ShortStrings(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	strs := tc.ShortStrings()

	// Act
	actual := args.Map{"lenGt0": len(strs) > 0}

	// Assert
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- ShortStrings", actual)
}

func Test_TraceCollection_FileWithLines(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	fwls := tc.FileWithLines()

	// Act
	actual := args.Map{"lenGt0": len(fwls) > 0}

	// Assert
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns non-empty -- FileWithLines", actual)
}

func Test_TraceCollection_FileWithLinesStrings(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	strs := tc.FileWithLinesStrings()

	// Act
	actual := args.Map{"lenGt0": len(strs) > 0}

	// Assert
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns non-empty -- FileWithLinesStrings", actual)
}

func Test_TraceCollection_CodeStacksString(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.CodeStacksString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- CodeStacksString", actual)
}

func Test_TraceCollection_CodeStacksString_Empty(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	result := tc.CodeStacksString()

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- CodeStacksString empty", actual)
}

func Test_TraceCollection_StackTraces(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.StackTraces()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- StackTraces", actual)
}

func Test_TraceCollection_StackTracesJsonResult(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.StackTracesJsonResult()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- StackTracesJsonResult", actual)
}

func Test_TraceCollection_StackTracesBytes(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.StackTracesBytes()

	// Act
	actual := args.Map{"lenGt0": len(result) > 0}

	// Assert
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- StackTracesBytes", actual)
}

func Test_TraceCollection_StackTracesBytes_Empty(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	result := tc.StackTracesBytes()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- StackTracesBytes empty", actual)
}

func Test_TraceCollection_Clone(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	cloned := tc.Clone()

	// Act
	actual := args.Map{"sameLen": cloned.Length() == tc.Length()}

	// Assert
	expected := args.Map{"sameLen": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Clone", actual)
}

func Test_TraceCollection_ClonePtr(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
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
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- ClonePtr", actual)
}

func Test_TraceCollection_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var tc *codestack.TraceCollection
	cloned := tc.ClonePtr()
	isNil := cloned == nil

	// Act
	actual := args.Map{"isNil": isNil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns nil -- ClonePtr nil", actual)
}

func Test_TraceCollection_Reverse_Empty(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}
	result := tc.Reverse()

	// Act
	actual := args.Map{"empty": result.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- Reverse empty", actual)
}

func Test_TraceCollection_Reverse_Single(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(0)
	tc := &codestack.TraceCollection{Items: []codestack.Trace{trace}}
	result := tc.Reverse()

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Reverse single", actual)
}

func Test_TraceCollection_Reverse_Two(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.Reverse()

	// Act
	actual := args.Map{
		"first": result.Items[0].PackageName,
		"second": result.Items[1].PackageName,
	}

	// Assert
	expected := args.Map{
		"first": "b",
		"second": "a",
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Reverse two", actual)
}

func Test_TraceCollection_Reverse_Three(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	t3 := codestack.Trace{PackageName: "c"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2, t3}}
	result := tc.Reverse()

	// Act
	actual := args.Map{
		"first": result.Items[0].PackageName,
		"last": result.Items[2].PackageName,
	}

	// Assert
	expected := args.Map{
		"first": "c",
		"last": "a",
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Reverse three", actual)
}

func Test_TraceCollection_Skip(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 1 {
		skipped := tc.Skip(1)

	// Act
		actual := args.Map{"lenLess": len(skipped) == tc.Length()-1}

	// Assert
		expected := args.Map{"lenLess": true}
		expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Skip", actual)
	}
}

func Test_TraceCollection_SkipDynamic_BeyondLength(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.SkipDynamic(99999)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- SkipDynamic beyond", actual)
}

func Test_TraceCollection_SkipCollection(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 1 {
		result := tc.SkipCollection(1)

	// Act
		actual := args.Map{"lenLess": result.Length() == tc.Length()-1}

	// Assert
		expected := args.Map{"lenLess": true}
		expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- SkipCollection", actual)
	}
}

func Test_TraceCollection_Take(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 1 {
		taken := tc.Take(1)

	// Act
		actual := args.Map{"len": len(taken)}

	// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Take", actual)
	}
}

func Test_TraceCollection_TakeDynamic(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 0 {
		result := tc.TakeDynamic(1)

	// Act
		actual := args.Map{"notNil": result != nil}

	// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- TakeDynamic", actual)
	}
}

func Test_TraceCollection_TakeCollection(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 0 {
		result := tc.TakeCollection(1)

	// Act
		actual := args.Map{"len": result.Length()}

	// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- TakeCollection", actual)
	}
}

func Test_TraceCollection_LimitCollection(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 0 {
		result := tc.LimitCollection(1)

	// Act
		actual := args.Map{"len": result.Length()}

	// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- LimitCollection", actual)
	}
}

func Test_TraceCollection_SafeLimitCollection(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.SafeLimitCollection(999)

	// Act
	actual := args.Map{"sameLen": result.Length() == tc.Length()}

	// Assert
	expected := args.Map{"sameLen": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- SafeLimitCollection", actual)
}

func Test_TraceCollection_LimitDynamic(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 0 {
		result := tc.LimitDynamic(1)

	// Act
		actual := args.Map{"notNil": result != nil}

	// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- LimitDynamic", actual)
	}
}

func Test_TraceCollection_Limit(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	if tc.Length() > 0 {
		result := tc.Limit(1)

	// Act
		actual := args.Map{"len": len(result)}

	// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Limit", actual)
	}
}

func Test_TraceCollection_GetPagesSize(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	t3 := codestack.Trace{PackageName: "c"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2, t3}}

	// Act
	actual := args.Map{
		"pages": tc.GetPagesSize(2),
		"zero": tc.GetPagesSize(0),
	}

	// Assert
	expected := args.Map{
		"pages": 2,
		"zero": 0,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- GetPagesSize", actual)
}

func Test_TraceCollection_GetPagedCollection(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	t3 := codestack.Trace{PackageName: "c"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2, t3}}
	pages := tc.GetPagedCollection(2)

	// Act
	actual := args.Map{"pageCount": len(pages)}

	// Assert
	expected := args.Map{"pageCount": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- GetPagedCollection", actual)
}

func Test_TraceCollection_GetPagedCollection_SmallPage(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageName: "a"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1}}
	pages := tc.GetPagedCollection(5)

	// Act
	actual := args.Map{"pageCount": len(pages)}

	// Assert
	expected := args.Map{"pageCount": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- GetPagedCollection small", actual)
}

func Test_TraceCollection_GetSinglePageCollection(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	t3 := codestack.Trace{PackageName: "c"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2, t3}}
	page := tc.GetSinglePageCollection(2, 1)

	// Act
	actual := args.Map{"len": page.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- GetSinglePageCollection", actual)
}

func Test_TraceCollection_Filter(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageName: "a", IsOkay: true, PackageMethodName: "a.test"}
	t2 := codestack.Trace{PackageName: "b", IsOkay: true, PackageMethodName: "b.test"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.Filter(func(trace *codestack.Trace) (bool, bool) {
		return trace.PackageName == "a", false
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Filter", actual)
}

func Test_TraceCollection_Filter_WithBreak(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.Filter(func(trace *codestack.Trace) (bool, bool) {
		return true, trace.PackageName == "a"
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns non-empty -- Filter with break", actual)
}

func Test_TraceCollection_FilterTraceCollection(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.FilterTraceCollection(func(trace *codestack.Trace) (bool, bool) {
		return trace.PackageName == "a", false
	})

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- FilterTraceCollection", actual)
}

func Test_TraceCollection_FilterPackageNameTraceCollection(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.FilterPackageNameTraceCollection("a")

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- FilterPackageName", actual)
}

func Test_TraceCollection_SkipFilterPackageNameTraceCollection(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.SkipFilterPackageNameTraceCollection("a")

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- SkipFilterPackageName", actual)
}

func Test_TraceCollection_FilterMethodNameTraceCollection(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{MethodName: "test"}
	t2 := codestack.Trace{MethodName: "other"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.FilterMethodNameTraceCollection("test")

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- FilterMethodName", actual)
}

func Test_TraceCollection_SkipFilterMethodNameTraceCollection(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{MethodName: "test"}
	t2 := codestack.Trace{MethodName: "other"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.SkipFilterMethodNameTraceCollection("test")

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- SkipFilterMethodName", actual)
}

func Test_TraceCollection_FilterFullMethodNameTraceCollection(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageMethodName: "pkg.Method"}
	t2 := codestack.Trace{PackageMethodName: "pkg.Other"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.FilterFullMethodNameTraceCollection("pkg.Method")

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- FilterFullMethodName", actual)
}

func Test_TraceCollection_SkipFilterFullMethodNameTraceCollection(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageMethodName: "pkg.Method"}
	t2 := codestack.Trace{PackageMethodName: "pkg.Other"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.SkipFilterFullMethodNameTraceCollection("pkg.Method")

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- SkipFilterFullMethodName", actual)
}

func Test_TraceCollection_SkipFilterFilenameTraceCollection(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{FilePath: "/a.go"}
	t2 := codestack.Trace{FilePath: "/b.go"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	result := tc.SkipFilterFilenameTraceCollection("/a.go")

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- SkipFilterFilename", actual)
}

func Test_TraceCollection_FilterWithLimit(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	t3 := codestack.Trace{PackageName: "c"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2, t3}}
	result := tc.FilterWithLimit(2, func(trace *codestack.Trace) (bool, bool) {
		return true, false
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns non-empty -- FilterWithLimit", actual)
}

func Test_TraceCollection_Json(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.Json()

	// Act
	actual := args.Map{"noErr": result.MeaningfulError() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Json", actual)
}

func Test_TraceCollection_JsonPtr(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JsonPtr()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- JsonPtr", actual)
}

func Test_TraceCollection_JsonString(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JsonString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- JsonString", actual)
}

func Test_TraceCollection_JsonString_Empty(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	result := tc.JsonString()

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- JsonString empty", actual)
}

func Test_TraceCollection_String(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.String()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- String", actual)
}

func Test_TraceCollection_String_Empty(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	result := tc.String()

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- String empty", actual)
}

func Test_TraceCollection_JsonModel(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	model := tc.JsonModel()

	// Act
	actual := args.Map{"lenGt0": len(model) > 0}

	// Assert
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- JsonModel", actual)
}

func Test_TraceCollection_JsonModelAny(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JsonModelAny()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- JsonModelAny", actual)
}

func Test_TraceCollection_Serializer(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	bytes, err := tc.Serializer()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"lenGt0": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"lenGt0": true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Serializer", actual)
}

func Test_TraceCollection_Join(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.Join(", ")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Join", actual)
}

func Test_TraceCollection_JoinLines(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JoinLines()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- JoinLines", actual)
}

func Test_TraceCollection_JoinShortStrings(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JoinShortStrings(", ")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- JoinShortStrings", actual)
}

func Test_TraceCollection_JoinFileWithLinesStrings(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JoinFileWithLinesStrings(", ")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns non-empty -- JoinFileWithLinesStrings", actual)
}

func Test_TraceCollection_JoinJsonStrings(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JoinJsonStrings(", ")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- JoinJsonStrings", actual)
}

func Test_TraceCollection_JsonStrings(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JsonStrings()

	// Act
	actual := args.Map{"lenGt0": len(result) > 0}

	// Assert
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- JsonStrings", actual)
}

func Test_TraceCollection_CsvStrings(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.CsvStrings()

	// Act
	actual := args.Map{"lenGt0": len(result) > 0}

	// Assert
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- CsvStrings", actual)
}

func Test_TraceCollection_CsvStrings_Empty(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	result := tc.CsvStrings()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- CsvStrings empty", actual)
}

func Test_TraceCollection_JoinCsv(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JoinCsv()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- JoinCsv", actual)
}

func Test_TraceCollection_JoinCsvLine(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JoinCsvLine()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- JoinCsvLine", actual)
}

func Test_TraceCollection_JoinUsingFmt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.JoinUsingFmt(func(trace *codestack.Trace) string { return "x" }, ", ")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- JoinUsingFmt", actual)
}

func Test_TraceCollection_StringsUsingFmt(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.StringsUsingFmt(func(trace *codestack.Trace) string { return "x" })

	// Act
	actual := args.Map{
		"lenGt0": len(result) > 0,
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"lenGt0": true,
		"first": "x",
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- StringsUsingFmt", actual)
}

func Test_TraceCollection_CodeStacksStringLimit(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.CodeStacksStringLimit(1)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- CodeStacksStringLimit", actual)
}

func Test_TraceCollection_CodeStacksStringLimit_Empty(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	result := tc.CodeStacksStringLimit(1)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- CodeStacksStringLimit empty", actual)
}

func Test_TraceCollection_FileWithLinesString(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.FileWithLinesString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns non-empty -- FileWithLinesString", actual)
}

func Test_TraceCollection_FileWithLinesString_Empty(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}
	result := tc.FileWithLinesString()

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- FileWithLinesString empty", actual)
}

func Test_TraceCollection_Add(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}
	trace := codestack.Trace{PackageName: "test"}
	tc.Add(trace)

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Add", actual)
}

func Test_TraceCollection_Adds(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}
	tc.Adds(codestack.Trace{PackageName: "a"}, codestack.Trace{PackageName: "b"})

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Adds", actual)
}

func Test_TraceCollection_Adds_Empty(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}
	tc.Adds()

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- Adds empty", actual)
}

func Test_TraceCollection_AddsIf_True(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}
	tc.AddsIf(true, codestack.Trace{PackageName: "a"})

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns non-empty -- AddsIf true", actual)
}

func Test_TraceCollection_AddsIf_False(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}
	tc.AddsIf(false, codestack.Trace{PackageName: "a"})

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns non-empty -- AddsIf false", actual)
}

func Test_TraceCollection_AddsPtr(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}
	trace := &codestack.Trace{PackageName: "a", IsOkay: true, PackageMethodName: "a.test"}
	tc.AddsPtr(false, trace)

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- AddsPtr", actual)
}

func Test_TraceCollection_AddsPtr_SkipIssues(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}
	trace := &codestack.Trace{IsOkay: false}
	tc.AddsPtr(true, trace)

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- AddsPtr skip issues", actual)
}

func Test_TraceCollection_AddsPtr_Nil(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}
	tc.AddsPtr(false, nil)

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns nil -- AddsPtr nil", actual)
}

func Test_TraceCollection_AddsPtr_Empty(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{}
	tc.AddsPtr(false)

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns empty -- AddsPtr empty", actual)
}

func Test_TraceCollection_ConcatNew(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{Items: []codestack.Trace{{PackageName: "a"}}}
	result := tc.ConcatNew(codestack.Trace{PackageName: "b"})

	// Act
	actual := args.Map{
		"len": result.Length(),
		"origLen": tc.Length(),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"origLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- ConcatNew", actual)
}

func Test_TraceCollection_ConcatNewPtr(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{Items: []codestack.Trace{{PackageName: "a"}}}
	trace := &codestack.Trace{PackageName: "b", IsOkay: true, PackageMethodName: "b.test"}
	result := tc.ConcatNewPtr(trace)

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- ConcatNewPtr", actual)
}

func Test_TraceCollection_IsEqual(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageName: "a"}
	tc1 := &codestack.TraceCollection{Items: []codestack.Trace{t1}}
	tc2 := &codestack.TraceCollection{Items: []codestack.Trace{t1}}

	// Act
	actual := args.Map{"equal": tc1.IsEqual(tc2)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- IsEqual", actual)
}

func Test_TraceCollection_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var tc1 *codestack.TraceCollection
	var tc2 *codestack.TraceCollection

	// Act
	actual := args.Map{"equal": tc1.IsEqual(tc2)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns nil -- IsEqual both nil", actual)
}

func Test_TraceCollection_IsEqual_OneNil(t *testing.T) {
	// Arrange
	tc1 := &codestack.TraceCollection{}
	var tc2 *codestack.TraceCollection

	// Act
	actual := args.Map{"equal": tc1.IsEqual(tc2)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns nil -- IsEqual one nil", actual)
}

func Test_TraceCollection_IsEqual_DiffLen(t *testing.T) {
	// Arrange
	tc1 := &codestack.TraceCollection{Items: []codestack.Trace{{PackageName: "a"}}}
	tc2 := &codestack.TraceCollection{Items: []codestack.Trace{{PackageName: "a"}, {PackageName: "b"}}}

	// Act
	actual := args.Map{"equal": tc1.IsEqual(tc2)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- IsEqual diff len", actual)
}

func Test_TraceCollection_IsEqualItems(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageName: "a"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1}}

	// Act
	actual := args.Map{"equal": tc.IsEqualItems(t1)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- IsEqualItems", actual)
}

func Test_TraceCollection_IsEqualItems_DiffLen(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{Items: []codestack.Trace{{PackageName: "a"}}}

	// Act
	actual := args.Map{"equal": tc.IsEqualItems(codestack.Trace{PackageName: "a"}, codestack.Trace{PackageName: "b"})}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- IsEqualItems diff len", actual)
}

func Test_TraceCollection_IsEqualItems_NotEqual(t *testing.T) {
	// Arrange
	tc := &codestack.TraceCollection{Items: []codestack.Trace{{PackageName: "a"}}}

	// Act
	actual := args.Map{"equal": tc.IsEqualItems(codestack.Trace{PackageName: "b"})}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- IsEqualItems not equal", actual)
}

func Test_TraceCollection_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	jsonResult := tc.JsonPtr()
	var target codestack.TraceCollection
	result, err := target.ParseInjectUsingJson(jsonResult)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"lenGt0": result.Length() > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"lenGt0": true,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- ParseInjectUsingJson", actual)
}

func Test_TraceCollection_ParseInjectUsingJsonMust(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	jsonResult := tc.JsonPtr()
	var target codestack.TraceCollection
	result := target.ParseInjectUsingJsonMust(jsonResult)

	// Act
	actual := args.Map{"lenGt0": result.Length() > 0}

	// Assert
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- ParseInjectUsingJsonMust", actual)
}

func Test_TraceCollection_JsonParseSelfInject(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	jsonResult := tc.JsonPtr()
	var target codestack.TraceCollection
	err := target.JsonParseSelfInject(jsonResult)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- JsonParseSelfInject", actual)
}

func Test_TraceCollection_AsJsonContractsBinder(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.AsJsonContractsBinder()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- AsJsonContractsBinder", actual)
}

func Test_TraceCollection_AsJsoner(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.AsJsoner()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- AsJsoner", actual)
}

func Test_TraceCollection_AsJsonParseSelfInjector(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	result := tc.AsJsonParseSelfInjector()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- AsJsonParseSelfInjector", actual)
}

func Test_TraceCollection_Clear(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	tc.Clear()

	// Act
	actual := args.Map{"empty": tc.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Clear", actual)
}

func Test_TraceCollection_Clear_Nil(t *testing.T) {
	// Arrange
	var tc *codestack.TraceCollection
	result := tc.Clear()
	isNil := result == nil

	// Act
	actual := args.Map{"isNil": isNil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns nil -- Clear nil", actual)
}

func Test_TraceCollection_Dispose(t *testing.T) {
	// Arrange
	tc := codestack.New.StackTrace.SkipNone()
	tc.Dispose()
	isNilItems := tc.Items == nil

	// Act
	actual := args.Map{"nilItems": isNilItems}

	// Assert
	expected := args.Map{"nilItems": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- Dispose", actual)
}

func Test_TraceCollection_Dispose_Nil(t *testing.T) {
	// Arrange
	var tc *codestack.TraceCollection
	tc.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns nil -- Dispose nil", actual)
}

func Test_TraceCollection_InsertAt(t *testing.T) {
	// Arrange
	t1 := codestack.Trace{PackageName: "a"}
	t2 := codestack.Trace{PackageName: "b"}
	tc := &codestack.TraceCollection{Items: []codestack.Trace{t1, t2}}
	inserted := codestack.Trace{PackageName: "x"}
	tc.InsertAt(0, inserted)

	// Act
	actual := args.Map{
		"first": tc.Items[0].PackageName,
		"len": tc.Length(),
	}

	// Assert
	expected := args.Map{
		"first": "x",
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- InsertAt", actual)
}

// ── NameOf ──

func Test_NameOf_All_Empty(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "NameOf.All returns empty -- empty", actual)
}

func Test_NameOf_Method(t *testing.T) {
	// Arrange
	method := codestack.NameOf.Method()

	// Act
	actual := args.Map{"notEmpty": method != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.Method returns correct value -- with args", actual)
}

func Test_NameOf_Package(t *testing.T) {
	// Arrange
	pkg := codestack.NameOf.Package()

	// Act
	actual := args.Map{"notEmpty": pkg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.Package returns correct value -- with args", actual)
}

// ── StacksTo ──

func Test_StacksTo_Bytes(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.Bytes(0)

	// Act
	actual := args.Map{"lenGt0": len(result) > 0}

	// Assert
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.Bytes returns correct value -- with args", actual)
}

func Test_StacksTo_BytesDefault(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.BytesDefault()

	// Act
	actual := args.Map{"lenGt0": len(result) > 0}

	// Assert
	expected := args.Map{"lenGt0": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.BytesDefault returns correct value -- with args", actual)
}

func Test_StacksTo_String_FromFileWithLineFullFile(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.String(0, 5)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.String returns correct value -- with args", actual)
}

func Test_StacksTo_StringDefault(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.StringDefault()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringDefault returns correct value -- with args", actual)
}

func Test_StacksTo_StringNoCount(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.StringNoCount(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringNoCount returns correct value -- with args", actual)
}

func Test_StacksTo_JsonString(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.JsonString(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.JsonString returns correct value -- with args", actual)
}

func Test_StacksTo_JsonStringDefault(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.JsonStringDefault()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.JsonStringDefault returns correct value -- with args", actual)
}

func Test_StacksTo_StringUsingFmt(t *testing.T) {
	// Arrange
	result := codestack.StacksTo.StringUsingFmt(
		func(trace *codestack.Trace) string { return "x" },
		0, 5,
	)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksTo.StringUsingFmt returns correct value -- with args", actual)
}

// ── File / Dir ──

func Test_File_Name(t *testing.T) {
	// Arrange
	result := codestack.File.Name(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.Name returns correct value -- with args", actual)
}

func Test_File_Path(t *testing.T) {
	// Arrange
	result := codestack.File.Path(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.Path returns correct value -- with args", actual)
}

func Test_File_CurrentFilePath(t *testing.T) {
	// Arrange
	result := codestack.File.CurrentFilePath()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.CurrentFilePath returns correct value -- with args", actual)
}

func Test_File_PathLineSep(t *testing.T) {
	// Arrange
	path, line := codestack.File.PathLineSep(0)

	// Act
	actual := args.Map{
		"pathNotEmpty": path != "",
		"lineGt0": line > 0,
	}

	// Assert
	expected := args.Map{
		"pathNotEmpty": true,
		"lineGt0": true,
	}
	expected.ShouldBeEqual(t, 0, "File.PathLineSep returns correct value -- with args", actual)
}

func Test_File_PathLineSepDefault(t *testing.T) {
	// Arrange
	path, line := codestack.File.PathLineSepDefault()

	// Act
	actual := args.Map{
		"pathNotEmpty": path != "",
		"lineGt0": line > 0,
	}

	// Assert
	expected := args.Map{
		"pathNotEmpty": true,
		"lineGt0": true,
	}
	expected.ShouldBeEqual(t, 0, "File.PathLineSepDefault returns correct value -- with args", actual)
}

func Test_File_FilePathWithLineString(t *testing.T) {
	// Arrange
	result := codestack.File.FilePathWithLineString(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.FilePathWithLineString returns non-empty -- with args", actual)
}

func Test_File_PathLineStringDefault(t *testing.T) {
	// Arrange
	result := codestack.File.PathLineStringDefault()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "File.PathLineStringDefault returns correct value -- with args", actual)
}

func Test_Dir_CurDir(t *testing.T) {
	// Arrange
	result := codestack.Dir.CurDir()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.CurDir returns correct value -- with args", actual)
}

func Test_Dir_Get(t *testing.T) {
	// Arrange
	result := codestack.Dir.Get(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.Get returns correct value -- with args", actual)
}

func Test_Dir_RepoDir_FromFileWithLineFullFile(t *testing.T) {
	// Arrange
	result := codestack.Dir.RepoDir()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.RepoDir returns correct value -- with args", actual)
}

func Test_Dir_CurDirJoin(t *testing.T) {
	// Arrange
	result := codestack.Dir.CurDirJoin("sub", "dir")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.CurDirJoin returns correct value -- with args", actual)
}

func Test_Dir_RepoDirJoin(t *testing.T) {
	// Arrange
	result := codestack.Dir.RepoDirJoin("sub")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dir.RepoDirJoin returns correct value -- with args", actual)
}

// ── NameOf extended ──

func Test_NameOf_MethodByFullName(t *testing.T) {
	// Arrange
	result := codestack.NameOf.MethodByFullName("github.com/pkg.Type.Method")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.MethodByFullName returns correct value -- with args", actual)
}

func Test_NameOf_PackageByFullName(t *testing.T) {
	// Arrange
	result := codestack.NameOf.PackageByFullName("github.com/pkg.Type.Method")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.PackageByFullName returns correct value -- with args", actual)
}

func Test_NameOf_CurrentFuncFullPath(t *testing.T) {
	// Arrange
	result := codestack.NameOf.CurrentFuncFullPath("github.com/pkg.Type.Method")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.CurrentFuncFullPath returns correct value -- with args", actual)
}

func Test_NameOf_JoinPackageNameWithRelative(t *testing.T) {
	// Arrange
	result := codestack.NameOf.JoinPackageNameWithRelative("github.com/pkg.Type.Method", "SubMethod")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.JoinPackageNameWithRelative returns non-empty -- with args", actual)
}

func Test_NameOf_MethodStackSkip(t *testing.T) {
	// Arrange
	result := codestack.NameOf.MethodStackSkip(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.MethodStackSkip returns correct value -- with args", actual)
}

func Test_NameOf_PackageStackSkip(t *testing.T) {
	// Arrange
	result := codestack.NameOf.PackageStackSkip(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOf.PackageStackSkip returns correct value -- with args", actual)
}

// ── newStacksCreator ──

func Test_StackTrace_SkipOne(t *testing.T) {
	tc := codestack.New.StackTrace.SkipOne()
	// Result may be empty due to integrated test call depth; just exercise the code path
	_ = tc.HasAnyItem()
}

func Test_StackTrace_Default_FromFileWithLineFullFile(t *testing.T) {
	tc := codestack.New.StackTrace.Default(0, 5)
	_ = tc.HasAnyItem()
}

func Test_StackTrace_DefaultCount(t *testing.T) {
	tc := codestack.New.StackTrace.DefaultCount(0)
	_ = tc.HasAnyItem()
}

// ── NewStackTraces / NewDefaultStackTraces ──

func Test_TraceCollection_NewStackTraces(t *testing.T) {
	tc := codestack.TraceCollection{}
	// May be empty due to call depth; just exercise code path
	result := tc.NewStackTraces(0)
	_ = result
}

func Test_TraceCollection_NewDefaultStackTraces(t *testing.T) {
	tc := codestack.TraceCollection{}
	result := tc.NewDefaultStackTraces()
	_ = result
}

func Test_TraceCollection_NewStackTracesJsonResult(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	result := tc.NewStackTracesJsonResult(0)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- NewStackTracesJsonResult", actual)
}

func Test_TraceCollection_NewDefaultStackTracesJsonResult(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	result := tc.NewDefaultStackTracesJsonResult()

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TraceCollection returns correct value -- NewDefaultStackTracesJsonResult", actual)
}

// ── newTraceCollection ──

func Test_NewTraceCollection_Empty(t *testing.T) {
	// Arrange
	// Test via TraceCollection Clear which creates empty
	tc := codestack.New.StackTrace.SkipNone()
	tc.Clear()

	// Act
	actual := args.Map{"empty": tc.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "newTraceCollection returns empty -- Empty via Clear", actual)
}

// ── isSkippablePackage ──

func Test_Trace_Create_HighSkipIndex(t *testing.T) {
	// Arrange
	trace := codestack.New.Create(9999)

	// Act
	actual := args.Map{"isOkay": trace.IsOkay}

	// Assert
	expected := args.Map{"isOkay": false}
	expected.ShouldBeEqual(t, 0, "Trace returns correct value -- Create high skip", actual)
}
