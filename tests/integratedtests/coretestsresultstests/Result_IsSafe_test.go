package coretestsresultstests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/results"
)

// ── Result methods ──

func Test_Result_IsSafe(t *testing.T) {
	// Arrange
	safe := results.ResultAny{Value: "ok"}
	unsafe := results.ResultAny{Panicked: true}
	errResult := results.ResultAny{Error: errors.New("e")}

	// Act
	actual := args.Map{
		"safe":   safe.IsSafe(),
		"unsafe": unsafe.IsSafe(),
		"errSafe": errResult.IsSafe(),
	}

	// Assert
	expected := args.Map{
		"safe":   true,
		"unsafe": false,
		"errSafe": false,
	}
	expected.ShouldBeEqual(t, 0, "IsSafe returns correct value -- with args", actual)
}

func Test_Result_HasError(t *testing.T) {
	// Arrange
	r := results.ResultAny{Error: errors.New("e")}

	// Act
	actual := args.Map{"hasErr": r.HasError()}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "HasError returns error -- with args", actual)
}

func Test_Result_HasPanicked(t *testing.T) {
	// Arrange
	r := results.ResultAny{Panicked: true}

	// Act
	actual := args.Map{"panicked": r.HasPanicked()}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HasPanicked panics -- with args", actual)
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
	expected.ShouldBeEqual(t, 0, "IsResult returns correct value -- with args", actual)
}

func Test_Result_IsResultTypeOf(t *testing.T) {
	// Arrange
	r := results.Result[int]{Value: 42}

	// Act
	actual := args.Map{
		"intType":  r.IsResultTypeOf(0),
		"nilType":  r.IsResultTypeOf(nil),
	}

	// Assert
	expected := args.Map{
		"intType":  true,
		"nilType":  false,
	}
	expected.ShouldBeEqual(t, 0, "IsResultTypeOf returns correct value -- with args", actual)
}

func Test_Result_IsError(t *testing.T) {
	// Arrange
	r := results.ResultAny{Error: errors.New("test")}
	noErr := results.ResultAny{}

	// Act
	actual := args.Map{
		"match":   r.IsError("test"),
		"noMatch": r.IsError("other"),
		"noErr":   noErr.IsError("test"),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
		"noErr":   false,
	}
	expected.ShouldBeEqual(t, 0, "IsError returns error -- with args", actual)
}

func Test_Result_ValueString(t *testing.T) {
	// Arrange
	r := results.Result[int]{Value: 42}

	// Act
	actual := args.Map{"val": r.ValueString()}

	// Assert
	expected := args.Map{"val": "42"}
	expected.ShouldBeEqual(t, 0, "ValueString returns non-empty -- with args", actual)
}

func Test_Result_ResultAt(t *testing.T) {
	// Arrange
	r := results.ResultAny{AllResults: []any{"a", "b"}}

	// Act
	actual := args.Map{
		"first":  r.ResultAt(0),
		"second": r.ResultAt(1),
		"outOfBounds": r.ResultAt(5) == nil,
		"negative":    r.ResultAt(-1) == nil,
	}

	// Assert
	expected := args.Map{
		"first":  "a",
		"second": "b",
		"outOfBounds": true,
		"negative":    true,
	}
	expected.ShouldBeEqual(t, 0, "ResultAt returns correct value -- with args", actual)
}

func Test_Result_ToMap(t *testing.T) {
	// Arrange
	r := results.Result[int]{Value: 42, ReturnCount: 1}
	m := r.ToMap()

	// Act
	actual := args.Map{
		"value":       m["value"],
		"panicked":    m["panicked"],
		"returnCount": m["returnCount"],
	}

	// Assert
	expected := args.Map{
		"value":       "42",
		"panicked":    false,
		"returnCount": 1,
	}
	expected.ShouldBeEqual(t, 0, "ToMap returns correct value -- with args", actual)
}

func Test_Result_ToMapCompact(t *testing.T) {
	// Arrange
	r := results.Result[int]{Value: 42}
	m := r.ToMapCompact()

	// Act
	actual := args.Map{
		"value": m["value"],
		"panicked": m["panicked"],
	}

	// Assert
	expected := args.Map{
		"value": "42",
		"panicked": false,
	}
	expected.ShouldBeEqual(t, 0, "ToMapCompact returns correct value -- with args", actual)
}

func Test_Result_String(t *testing.T) {
	// Arrange
	rNormal := results.Result[int]{Value: 42, ReturnCount: 1}
	rPanic := results.Result[int]{Panicked: true, PanicValue: "boom"}
	rErr := results.Result[int]{Value: 0, Error: errors.New("e"), ReturnCount: 1}

	// Act
	actual := args.Map{
		"normalNotEmpty": rNormal.String() != "",
		"panicNotEmpty":  rPanic.String() != "",
		"errNotEmpty":    rErr.String() != "",
	}

	// Assert
	expected := args.Map{
		"normalNotEmpty": true,
		"panicNotEmpty":  true,
		"errNotEmpty":    true,
	}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- with args", actual)
}

// ── Results (two-value) ──

func Test_Results_String(t *testing.T) {
	// Arrange
	rNormal := results.Results[int, bool]{Result: results.Result[int]{Value: 42}, Result2: true}
	rPanic := results.Results[int, bool]{Result: results.Result[int]{Panicked: true, PanicValue: "boom"}}
	rErr := results.Results[int, bool]{Result: results.Result[int]{Error: errors.New("e")}, Result2: false}

	// Act
	actual := args.Map{
		"normalNotEmpty": rNormal.String() != "",
		"panicNotEmpty":  rPanic.String() != "",
		"errNotEmpty":    rErr.String() != "",
	}

	// Assert
	expected := args.Map{
		"normalNotEmpty": true,
		"panicNotEmpty":  true,
		"errNotEmpty":    true,
	}
	expected.ShouldBeEqual(t, 0, "Results_String returns correct value -- with args", actual)
}

func Test_Results_IsResult2(t *testing.T) {
	// Arrange
	r := results.Results[int, string]{Result2: "hello"}

	// Act
	actual := args.Map{
		"match": r.IsResult2("hello"),
		"noMatch": r.IsResult2("other"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsResult2 returns correct value -- with args", actual)
}

func Test_Results_Result2String(t *testing.T) {
	// Arrange
	r := results.Results[int, string]{Result2: "hello"}

	// Act
	actual := args.Map{"val": r.Result2String()}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Result2String returns correct value -- with args", actual)
}

func Test_FromResultAny(t *testing.T) {
	// Arrange
	raw := results.ResultAny{
		Value:       42,
		AllResults:  []any{42, "hello"},
		ReturnCount: 2,
	}
	r := results.FromResultAny[int, string](raw)

	// Act
	actual := args.Map{
		"val1": r.Value,
		"val2": r.Result2,
		"count": r.ReturnCount,
	}

	// Assert
	expected := args.Map{
		"val1": 42,
		"val2": "hello",
		"count": 2,
	}
	expected.ShouldBeEqual(t, 0, "FromResultAny returns correct value -- with args", actual)
}

func Test_FromResultAny_TypeMismatch(t *testing.T) {
	// Arrange
	raw := results.ResultAny{AllResults: []any{"not-int", 42}}
	r := results.FromResultAny[int, string](raw)

	// Act
	actual := args.Map{
		"val1": r.Value,
		"val2": r.Result2,
	}

	// Assert
	expected := args.Map{
		"val1": 0,
		"val2": "",
	}
	expected.ShouldBeEqual(t, 0, "FromResultAny_TypeMismatch returns correct value -- with args", actual)
}

// ── InvokeWithPanicRecovery ──

func Test_InvokeWithPanicRecovery_NilFunc(t *testing.T) {
	// Arrange
	result := results.InvokeWithPanicRecovery(nil, nil)

	// Act
	actual := args.Map{"panicked": result.Panicked}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "NilFunc returns nil -- with args", actual)
}

func Test_InvokeWithPanicRecovery_NotFunc(t *testing.T) {
	// Arrange
	result := results.InvokeWithPanicRecovery(42, nil)

	// Act
	actual := args.Map{"panicked": result.Panicked}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "NotFunc returns correct value -- with args", actual)
}

func Test_InvokeWithPanicRecovery_SimpleFunc(t *testing.T) {
	// Arrange
	fn := func(x int) int { return x * 2 }
	result := results.InvokeWithPanicRecovery(fn, nil, 5)
	// buildCallArgs creates a zero int for nil receiver + arg 5 = 2 args for 1-param func → panic

	// Act
	actual := args.Map{"panicked": result.Panicked}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "SimpleFunc returns correct value -- with args", actual)
}

func Test_InvokeWithPanicRecovery_VoidFunc(t *testing.T) {
	// Arrange
	called := false
	fn := func() { called = true }
	result := results.InvokeWithPanicRecovery(fn, nil)

	// Act
	actual := args.Map{
		"panicked": result.Panicked,
		"called": called,
		"count": result.ReturnCount,
	}

	// Assert
	expected := args.Map{
		"panicked": false,
		"called": true,
		"count": 0,
	}
	expected.ShouldBeEqual(t, 0, "VoidFunc returns correct value -- with args", actual)
}

func Test_InvokeWithPanicRecovery_PanicFunc(t *testing.T) {
	// Arrange
	fn := func() { panic("boom") }
	result := results.InvokeWithPanicRecovery(fn, nil)

	// Act
	actual := args.Map{
		"panicked": result.Panicked,
		"panicVal": fmt.Sprintf("%v", result.PanicValue),
	}

	// Assert
	expected := args.Map{
		"panicked": true,
		"panicVal": "boom",
	}
	expected.ShouldBeEqual(t, 0, "PanicFunc panics -- with args", actual)
}

// ── MethodName ──

func Test_MethodName_Nil(t *testing.T) {
	// Arrange
	result := results.MethodName(nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "MethodName_Nil returns nil -- with args", actual)
}

func Test_MethodName_NotFunc(t *testing.T) {
	// Arrange
	result := results.MethodName(42)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "MethodName_NotFunc returns correct value -- with args", actual)
}

func Test_MethodName_Func(t *testing.T) {
	// Arrange
	result := results.MethodName(fmt.Sprintf)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MethodName_Func returns correct value -- with args", actual)
}

// ── ResultAssert (ShouldMatchResult) ──

func Test_ShouldMatchResult_Basic(t *testing.T) {
	r := results.Result[int]{Value: 42, Panicked: false, ReturnCount: 1}
	exp := results.ResultAny{Value: "42", Panicked: false, ReturnCount: 1}
	// Should not fail
	r.ShouldMatchResult(t, 0, "basic", exp)
}

func Test_ShouldMatchResult_ExplicitFields(t *testing.T) {
	r := results.Result[int]{Value: 42, Panicked: false}
	exp := results.ResultAny{Panicked: false}
	r.ShouldMatchResult(t, 0, "explicit", exp, "panicked")
}
