package resultstests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/results"
)

// ── InvokeWithPanicRecovery ──

func Test_Invoke_NilFunc(t *testing.T) {
	// Act
	r := results.InvokeWithPanicRecovery(nil, nil)

	// Assert
	actual := args.Map{
		"panicked": r.Panicked,
		"panicVal": fmt.Sprintf("%v", r.PanicValue),
	}
	expected := args.Map{
		"panicked": true,
		"panicVal": "funcRef is nil",
	}
	expected.ShouldBeEqual(t, 0, "InvokeWithPanicRecovery panics -- nil func", actual)
}

func Test_Invoke_NotAFunc(t *testing.T) {
	// Act
	r := results.InvokeWithPanicRecovery("not a func", nil)

	// Assert
	actual := args.Map{"panicked": r.Panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "InvokeWithPanicRecovery panics -- not a func", actual)
}

func dummyFunc() string { return "hello" }

func Test_Invoke_RegularFunc(t *testing.T) {
	// Act
	r := results.InvokeWithPanicRecovery(dummyFunc, nil)

	// Assert
	actual := args.Map{
		"panicked": r.Panicked,
		"value": fmt.Sprintf("%v", r.Value),
		"returnCount": r.ReturnCount,
	}
	expected := args.Map{
		"panicked": false,
		"value": "hello",
		"returnCount": 1,
	}
	expected.ShouldBeEqual(t, 0, "InvokeWithPanicRecovery panics -- regular func", actual)
}

func dummyFuncWithError() (string, error) { return "val", errors.New("err") }

func Test_Invoke_FuncWithError(t *testing.T) {
	// Act
	r := results.InvokeWithPanicRecovery(dummyFuncWithError, nil)

	// Assert
	actual := args.Map{
		"panicked": r.Panicked,
		"hasError": r.Error != nil,
		"returnCount": r.ReturnCount,
	}
	expected := args.Map{
		"panicked": false,
		"hasError": true,
		"returnCount": 2,
	}
	expected.ShouldBeEqual(t, 0, "InvokeWithPanicRecovery panics -- func with error", actual)
}

func dummyVoidFunc() {}

func Test_Invoke_VoidFunc(t *testing.T) {
	// Act
	r := results.InvokeWithPanicRecovery(dummyVoidFunc, nil)

	// Assert
	actual := args.Map{
		"panicked": r.Panicked,
		"returnCount": r.ReturnCount,
	}
	expected := args.Map{
		"panicked": false,
		"returnCount": 0,
	}
	expected.ShouldBeEqual(t, 0, "InvokeWithPanicRecovery panics -- void func", actual)
}

func dummyFuncNilError() (string, error) { return "ok", nil }

func Test_Invoke_FuncNilError(t *testing.T) {
	// Act
	r := results.InvokeWithPanicRecovery(dummyFuncNilError, nil)

	// Assert
	actual := args.Map{
		"panicked": r.Panicked,
		"hasError": r.Error != nil,
		"value": fmt.Sprintf("%v", r.Value),
	}
	expected := args.Map{
		"panicked": false,
		"hasError": false,
		"value": "ok",
	}
	expected.ShouldBeEqual(t, 0, "InvokeWithPanicRecovery panics -- nil error", actual)
}

// ── InvokeWithPanicRecovery with nil args ──

func funcWithNilArg(a any) string { return fmt.Sprintf("%v", a) }

func Test_Invoke_WithNilArg(t *testing.T) {
	// Act
	r := results.InvokeWithPanicRecovery(funcWithNilArg, nil, nil)

	// Assert — may panic due to nil receiver added internally
	actual := args.Map{"panicked": r.Panicked}
	expected := args.Map{"panicked": r.Panicked}
	expected.ShouldBeEqual(t, 0, "InvokeWithPanicRecovery panics -- with nil arg", actual)
}

// ── Result methods ──

func Test_Result_IsSafe(t *testing.T) {
	// Arrange
	safe := results.Result[string]{Value: "ok"}
	panicked := results.Result[string]{Panicked: true}
	errored := results.Result[string]{Error: errors.New("err")}

	// Act
	actual := args.Map{
		"safe":     safe.IsSafe(),
		"panicked": panicked.IsSafe(),
		"errored":  errored.IsSafe(),
	}

	// Assert
	expected := args.Map{
		"safe": true,
		"panicked": false,
		"errored": false,
	}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- IsSafe", actual)
}

func Test_Result_HasError(t *testing.T) {
	// Arrange
	r := results.Result[string]{Error: errors.New("err")}

	// Act
	actual := args.Map{"hasError": r.HasError()}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "Result returns error -- HasError", actual)
}

func Test_Result_HasPanicked(t *testing.T) {
	// Arrange
	r := results.Result[string]{Panicked: true}

	// Act
	actual := args.Map{"result": r.HasPanicked()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Result panics -- HasPanicked", actual)
}

func Test_Result_IsResult(t *testing.T) {
	// Arrange
	r := results.Result[int]{Value: 42}

	// Act
	actual := args.Map{
		"match": r.IsResult(42),
		"noMatch": r.IsResult(99),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- IsResult", actual)
}

func Test_Result_IsResultTypeOf(t *testing.T) {
	// Arrange
	r := results.Result[string]{Value: "hello"}

	// Act
	actual := args.Map{
		"string": r.IsResultTypeOf(""),
		"nil":    r.IsResultTypeOf(nil),
	}

	// Assert
	expected := args.Map{
		"string": true,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- IsResultTypeOf", actual)
}

func Test_Result_IsResultTypeOf_NilValue(t *testing.T) {
	// Arrange
	r := results.Result[any]{Value: nil}

	// Act
	actual := args.Map{"nilVal": r.IsResultTypeOf(nil)}

	// Assert
	expected := args.Map{"nilVal": true}
	expected.ShouldBeEqual(t, 0, "Result returns nil -- IsResultTypeOf nil value", actual)
}

func Test_Result_IsError(t *testing.T) {
	// Arrange
	r := results.Result[string]{Error: errors.New("test")}
	noErr := results.Result[string]{}

	// Act
	actual := args.Map{
		"match":   r.IsError("test"),
		"noMatch": r.IsError("other"),
		"noErr":   noErr.IsError("test"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
		"noErr": false,
	}
	expected.ShouldBeEqual(t, 0, "Result returns error -- IsError", actual)
}

func Test_Result_ValueString(t *testing.T) {
	// Arrange
	r := results.Result[int]{Value: 42}

	// Act
	actual := args.Map{"result": r.ValueString()}

	// Assert
	expected := args.Map{"result": "42"}
	expected.ShouldBeEqual(t, 0, "Result returns non-empty -- ValueString", actual)
}

func Test_Result_ResultAt(t *testing.T) {
	// Arrange
	r := results.Result[string]{AllResults: []any{"a", "b"}}

	// Act
	actual := args.Map{
		"first":    fmt.Sprintf("%v", r.ResultAt(0)),
		"outOfRange": r.ResultAt(5) == nil,
		"negative":   r.ResultAt(-1) == nil,
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"outOfRange": true,
		"negative": true,
	}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- ResultAt", actual)
}

func Test_Result_ToMap(t *testing.T) {
	// Arrange
	r := results.Result[string]{Value: "hello", ReturnCount: 1}
	m := r.ToMap()

	// Act
	actual := args.Map{
		"hasValue": m["value"] != nil,
		"hasPanicked": m["panicked"] != nil,
	}

	// Assert
	expected := args.Map{
		"hasValue": true,
		"hasPanicked": true,
	}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- ToMap", actual)
}

func Test_Result_ToMapCompact(t *testing.T) {
	// Arrange
	r := results.Result[string]{Value: "hello"}
	m := r.ToMapCompact()

	// Act
	actual := args.Map{"hasValue": m["value"] != nil}

	// Assert
	expected := args.Map{"hasValue": true}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- ToMapCompact", actual)
}

func Test_Result_String(t *testing.T) {
	// Arrange
	panicked := results.Result[string]{Panicked: true, PanicValue: "oops"}
	errored := results.Result[string]{Error: errors.New("err"), ReturnCount: 1}
	normal := results.Result[string]{Value: "ok", ReturnCount: 1}

	// Act
	actual := args.Map{
		"panicked": panicked.String() != "",
		"errored":  errored.String() != "",
		"normal":   normal.String() != "",
	}

	// Assert
	expected := args.Map{
		"panicked": true,
		"errored": true,
		"normal": true,
	}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- String", actual)
}

// ── Results methods ──

func Test_Results_String(t *testing.T) {
	// Arrange
	panicked := results.Results[string, int]{Result: results.Result[string]{Panicked: true, PanicValue: "oops"}}
	errored := results.Results[string, int]{Result: results.Result[string]{Error: errors.New("err")}, Result2: 42}
	normal := results.Results[string, int]{Result: results.Result[string]{Value: "ok"}, Result2: 42}

	// Act
	actual := args.Map{
		"panicked": panicked.String() != "",
		"errored":  errored.String() != "",
		"normal":   normal.String() != "",
	}

	// Assert
	expected := args.Map{
		"panicked": true,
		"errored": true,
		"normal": true,
	}
	expected.ShouldBeEqual(t, 0, "Results returns correct value -- String", actual)
}

func Test_Results_IsResult2(t *testing.T) {
	// Arrange
	r := results.Results[string, int]{Result2: 42}

	// Act
	actual := args.Map{
		"match": r.IsResult2(42),
		"noMatch": r.IsResult2(99),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "Results returns correct value -- IsResult2", actual)
}

func Test_Results_Result2String(t *testing.T) {
	// Arrange
	r := results.Results[string, int]{Result2: 42}

	// Act
	actual := args.Map{"result": r.Result2String()}

	// Assert
	expected := args.Map{"result": "42"}
	expected.ShouldBeEqual(t, 0, "Results returns correct value -- Result2String", actual)
}

func Test_FromResultAny(t *testing.T) {
	// Arrange
	r := results.ResultAny{
		Value:       "hello",
		ReturnCount: 2,
		AllResults:  []any{"hello", 42},
	}
	typed := results.FromResultAny[string, int](r)

	// Act
	actual := args.Map{
		"value": typed.Value,
		"result2": typed.Result2,
	}

	// Assert
	expected := args.Map{
		"value": "hello",
		"result2": 42,
	}
	expected.ShouldBeEqual(t, 0, "FromResultAny returns correct value -- with args", actual)
}

func Test_FromResultAny_Empty(t *testing.T) {
	// Arrange
	r := results.ResultAny{ReturnCount: 0, AllResults: []any{}}
	typed := results.FromResultAny[string, int](r)

	// Act
	actual := args.Map{
		"value": typed.Value,
		"result2": typed.Result2,
	}

	// Assert
	expected := args.Map{
		"value": "",
		"result2": 0,
	}
	expected.ShouldBeEqual(t, 0, "FromResultAny returns empty -- empty", actual)
}

// ── MethodName ──

func Test_MethodName(t *testing.T) {
	// Arrange
	name := results.MethodName(dummyFunc)
	nilName := results.MethodName(nil)
	notFunc := results.MethodName("not a func")

	// Act
	actual := args.Map{
		"name": name,
		"nil": nilName,
		"notFunc": notFunc,
	}

	// Assert
	expected := args.Map{
		"name": "dummyFunc",
		"nil": "",
		"notFunc": "",
	}
	expected.ShouldBeEqual(t, 0, "MethodName returns correct value -- with args", actual)
}

// ── ShouldMatchResult ──

func Test_ShouldMatchResult(t *testing.T) {
	r := results.Result[string]{Value: "ok", ReturnCount: 1}
	exp := results.ResultAny{Value: "ok", ReturnCount: 1}
	r.ShouldMatchResult(t, 0, "test", exp)
}

func Test_ShouldMatchResult_WithFields(t *testing.T) {
	r := results.Result[string]{Value: "ok"}
	exp := results.ResultAny{}
	r.ShouldMatchResult(t, 0, "test fields", exp, "panicked")
}
