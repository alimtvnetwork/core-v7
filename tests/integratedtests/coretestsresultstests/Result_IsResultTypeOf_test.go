package coretestsresultstests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/results"
)

// ── Result edge cases ──

func Test_Result_IsResultTypeOf_Nil(t *testing.T) {
	// Arrange
	r := results.ResultAny{Value: nil}

	// Act
	actual := args.Map{"nilExpected": r.IsResultTypeOf(nil)}

	// Assert
	expected := args.Map{"nilExpected": true}
	expected.ShouldBeEqual(t, 0, "IsResultTypeOf returns nil -- nil", actual)
}

func Test_Result_String_Panicked(t *testing.T) {
	// Arrange
	r := results.ResultAny{Panicked: true, PanicValue: "boom"}

	// Act
	actual := args.Map{"notEmpty": r.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String panics -- panicked", actual)
}

func Test_Result_String_Error(t *testing.T) {
	// Arrange
	r := results.ResultAny{Error: errors.New("fail"), ReturnCount: 1}

	// Act
	actual := args.Map{"notEmpty": r.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns error -- error", actual)
}

func Test_Result_String_Normal(t *testing.T) {
	// Arrange
	r := results.ResultAny{Value: "ok", ReturnCount: 1}

	// Act
	actual := args.Map{"notEmpty": r.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- normal", actual)
}

// ── ResultsAny methods ──

func Test_ResultsAny_String_Panicked(t *testing.T) {
	// Arrange
	r := results.ResultsAny{Result: results.ResultAny{Panicked: true, PanicValue: "p"}}

	// Act
	actual := args.Map{"notEmpty": r.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ResultsAny panics -- panicked", actual)
}

func Test_ResultsAny_String_Error(t *testing.T) {
	// Arrange
	r := results.ResultsAny{
		Result:  results.ResultAny{Value: "v", Error: errors.New("e")},
		Result2: "r2",
	}

	// Act
	actual := args.Map{"notEmpty": r.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ResultsAny returns error -- error", actual)
}

func Test_ResultsAny_String_Normal(t *testing.T) {
	// Arrange
	r := results.ResultsAny{
		Result:  results.ResultAny{Value: "v"},
		Result2: "r2",
	}

	// Act
	actual := args.Map{"notEmpty": r.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ResultsAny returns correct value -- normal", actual)
}

func Test_ResultsAny_IsResult2(t *testing.T) {
	// Arrange
	r := results.ResultsAny{Result2: "hello"}

	// Act
	actual := args.Map{
		"match": r.IsResult2("hello"),
		"noMatch": r.IsResult2("x"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "ResultsAny returns correct value -- IsResult2", actual)
}

func Test_ResultsAny_Result2String(t *testing.T) {
	// Arrange
	r := results.ResultsAny{Result2: 99}

	// Act
	actual := args.Map{"val": r.Result2String()}

	// Assert
	expected := args.Map{"val": "99"}
	expected.ShouldBeEqual(t, 0, "ResultsAny returns correct value -- Result2String", actual)
}

// ── FromResultAny edge ──

func Test_FromResultAny_Empty(t *testing.T) {
	// Arrange
	ra := results.ResultAny{AllResults: []any{}}
	r := results.FromResultAny[string, string](ra)

	// Act
	actual := args.Map{
		"val": r.Value,
		"val2": r.Result2,
	}

	// Assert
	expected := args.Map{
		"val": "",
		"val2": "",
	}
	expected.ShouldBeEqual(t, 0, "FromResultAny returns empty -- empty", actual)
}

// ── InvokeWithPanicRecovery extended ──

type extCovTestStruct struct{}

func (s *extCovTestStruct) Hello() string { return "hi" }

func Test_Invoke_NilReceiver_Ext(t *testing.T) {
	// Arrange
	r := results.InvokeWithPanicRecovery((*extCovTestStruct).Hello, nil)
	// Hello() doesn't dereference receiver, so nil receiver works fine

	// Act
	actual := args.Map{
		"panicked": r.Panicked,
		"val": r.Value,
	}

	// Assert
	expected := args.Map{
		"panicked": false,
		"val": "hi",
	}
	expected.ShouldBeEqual(t, 0, "Invoke panics -- nil receiver panics", actual)
}

func Test_Invoke_ValidCall_Ext(t *testing.T) {
	// Arrange
	s := &extCovTestStruct{}
	r := results.InvokeWithPanicRecovery((*extCovTestStruct).Hello, s)

	// Act
	actual := args.Map{
		"value": fmt.Sprintf("%v", r.Value),
		"panicked": r.Panicked,
		"count": r.ReturnCount,
	}

	// Assert
	expected := args.Map{
		"value": "hi",
		"panicked": false,
		"count": 1,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns non-empty -- valid", actual)
}

type extCovErrStruct struct{}

func (s *extCovErrStruct) Fail() error { return errors.New("fail") }
func (s *extCovErrStruct) Ok() error   { return nil }

func Test_Invoke_ErrorReturn_Ext(t *testing.T) {
	// Arrange
	s := &extCovErrStruct{}
	r := results.InvokeWithPanicRecovery((*extCovErrStruct).Fail, s)

	// Act
	actual := args.Map{
		"hasError": r.HasError(),
		"panicked": r.Panicked,
	}

	// Assert
	expected := args.Map{
		"hasError": true,
		"panicked": false,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns error -- error return", actual)
}

func Test_Invoke_NilErrorReturn_Ext(t *testing.T) {
	// Arrange
	s := &extCovErrStruct{}
	r := results.InvokeWithPanicRecovery((*extCovErrStruct).Ok, s)

	// Act
	actual := args.Map{
		"hasError": r.HasError(),
		"panicked": r.Panicked,
	}

	// Assert
	expected := args.Map{
		"hasError": false,
		"panicked": false,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns nil -- nil error return", actual)
}

type extCovVoidStruct struct{}

func (s *extCovVoidStruct) DoNothing() {}

func Test_Invoke_VoidReturn_Ext(t *testing.T) {
	// Arrange
	s := &extCovVoidStruct{}
	r := results.InvokeWithPanicRecovery((*extCovVoidStruct).DoNothing, s)

	// Act
	actual := args.Map{
		"count": r.ReturnCount,
		"panicked": r.Panicked,
	}

	// Assert
	expected := args.Map{
		"count": 0,
		"panicked": false,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct value -- void", actual)
}

type extCovMultiStruct struct{}

func (s *extCovMultiStruct) TwoVals() (string, int) { return "a", 1 }

func Test_Invoke_MultiReturn_Ext(t *testing.T) {
	// Arrange
	s := &extCovMultiStruct{}
	r := results.InvokeWithPanicRecovery((*extCovMultiStruct).TwoVals, s)

	// Act
	actual := args.Map{
		"count":  r.ReturnCount,
		"val":    fmt.Sprintf("%v", r.Value),
		"allLen": len(r.AllResults),
		"hasErr": r.HasError(),
	}

	// Assert
	expected := args.Map{
		"count":  2,
		"val":    "a",
		"allLen": 2,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct value -- multi return", actual)
}

// ── Invoke with nil args ──

type extCovArgStruct struct{}

func (s *extCovArgStruct) WithArg(v any) string {
	return fmt.Sprintf("%v", v)
}

func Test_Invoke_NilArg_Ext(t *testing.T) {
	// Arrange
	s := &extCovArgStruct{}
	r := results.InvokeWithPanicRecovery((*extCovArgStruct).WithArg, s, nil)

	// Act
	actual := args.Map{
		"panicked": r.Panicked,
		"count": r.ReturnCount,
	}

	// Assert
	expected := args.Map{
		"panicked": false,
		"count": 1,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns nil -- nil arg", actual)
}

// ── ExpectAnyError sentinel ──

func Test_ExpectAnyError_Ext(t *testing.T) {
	// Act
	actual := args.Map{"notNil": results.ExpectAnyError != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectAnyError returns error -- sentinel", actual)
}

// ── FilterByFields via ToMap ──

func Test_FilterByFields_MissingKey_Ext(t *testing.T) {
	// Arrange
	r := results.ResultAny{}
	m := r.ToMap()

	// Act
	actual := args.Map{
		"hasValue":    m.Has("value"),
		"hasPanicked": m.Has("panicked"),
	}

	// Assert
	expected := args.Map{
		"hasValue":    true,
		"hasPanicked": true,
	}
	expected.ShouldBeEqual(t, 0, "FilterByFields returns correct value -- with args", actual)
}

// ── MethodName combined ──

func Test_MethodName_Combined_Ext(t *testing.T) {
	// Arrange
	name := results.MethodName((*extCovTestStruct).Hello)
	nilName := results.MethodName(nil)
	nonFunc := results.MethodName("notafunc")

	// Act
	actual := args.Map{
		"name": name,
		"nil": nilName,
		"nonFunc": nonFunc,
	}

	// Assert
	expected := args.Map{
		"name": "Hello",
		"nil": "",
		"nonFunc": "",
	}
	expected.ShouldBeEqual(t, 0, "MethodName returns correct value -- combined", actual)
}
