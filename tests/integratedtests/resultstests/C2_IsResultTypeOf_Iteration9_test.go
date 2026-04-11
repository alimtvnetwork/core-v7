package resultstests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/results"
)

// ── Result.IsResultTypeOf — uncovered branches ──

func Test_C2_IsResultTypeOf_ActualTypeNil(t *testing.T) {
	// Arrange
	// Value is nil (any), expected is non-nil → actualType == nil → false
	r := results.Result[any]{Value: nil}

	// Act
	actual := args.Map{"result": r.IsResultTypeOf("something")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsResultTypeOf returns nil -- nil value vs non-nil expected", actual)
}

func Test_C2_IsResultTypeOf_NilExpected_NonZeroValue(t *testing.T) {
	// Arrange
	// expected=nil, value is non-zero → rv.IsValid()=true, rv.IsZero()=false → false
	r := results.Result[string]{Value: "hello"}

	// Act
	actual := args.Map{"result": r.IsResultTypeOf(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsResultTypeOf returns nil -- nil expected with non-zero value", actual)
}

func Test_C2_IsResultTypeOf_NilExpected_ZeroInt(t *testing.T) {
	// Arrange
	// expected=nil, value=0 → rv.IsZero()=true → true
	r := results.Result[int]{Value: 0}

	// Act
	actual := args.Map{"result": r.IsResultTypeOf(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsResultTypeOf returns nil -- nil expected with zero int", actual)
}

// ── Result.String — ensure all 3 branches are covered ──

func Test_C2_Result_String_Panicked(t *testing.T) {
	// Arrange
	r := results.Result[string]{Panicked: true, PanicValue: "boom"}
	s := r.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C2_Result_String_WithError(t *testing.T) {
	// Arrange
	r := results.Result[string]{Value: "v", Error: errors.New("e"), ReturnCount: 1}
	s := r.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C2_Result_String_Normal(t *testing.T) {
	// Arrange
	r := results.Result[int]{Value: 42, ReturnCount: 1}
	s := r.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ── Results.String — 3 branches ──

func Test_C2_Results_String_Panicked(t *testing.T) {
	// Arrange
	r := results.Results[string, int]{Result: results.Result[string]{Panicked: true, PanicValue: "p"}}

	// Act
	actual := args.Map{"result": r.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C2_Results_String_WithError(t *testing.T) {
	// Arrange
	r := results.Results[string, int]{
		Result:  results.Result[string]{Value: "v", Error: errors.New("e")},
		Result2: 5,
	}

	// Act
	actual := args.Map{"result": r.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C2_Results_String_Normal(t *testing.T) {
	// Arrange
	r := results.Results[string, int]{
		Result:  results.Result[string]{Value: "v"},
		Result2: 10,
	}

	// Act
	actual := args.Map{"result": r.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ── FromResultAny — type assertion failures ──

func Test_C2_FromResultAny_TypeAssertionFails(t *testing.T) {
	// Arrange
	r := results.ResultAny{
		ReturnCount: 2,
		AllResults:  []any{42, "not-int"}, // first is int not string, second is string not int
	}
	typed := results.FromResultAny[string, int](r)
	// Both assertions fail → zero values

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
	expected.ShouldBeEqual(t, 0, "FromResultAny returns correct value -- type assertion fails", actual)
}

func Test_C2_FromResultAny_SingleResult(t *testing.T) {
	// Arrange
	r := results.ResultAny{
		ReturnCount: 1,
		AllResults:  []any{"hello"},
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
		"result2": 0,
	}
	expected.ShouldBeEqual(t, 0, "FromResultAny returns correct value -- single result", actual)
}

// ── InvokeWithPanicRecovery — method with receiver ──

type testReceiver struct{ Name string }

func (tr *testReceiver) Greet() string { return "hello " + tr.Name }
func (tr *testReceiver) GreetErr() (string, error) {
	if tr == nil {
		return "", errors.New("nil receiver")
	}
	return "hi " + tr.Name, nil
}

func Test_C2_Invoke_WithReceiver(t *testing.T) {
	// Arrange
	r := results.InvokeWithPanicRecovery((*testReceiver).Greet, &testReceiver{Name: "World"})

	// Act
	actual := args.Map{
		"panicked": r.Panicked,
		"value": fmt.Sprintf("%v", r.Value),
	}

	// Assert
	expected := args.Map{
		"panicked": false,
		"value": "hello World",
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns non-empty -- with receiver", actual)
}

func Test_C2_Invoke_NilReceiver_PanicRecovery(t *testing.T) {
	// Arrange
	// Calling a method on nil receiver that accesses fields → panic
	r := results.InvokeWithPanicRecovery((*testReceiver).Greet, nil)

	// Act
	actual := args.Map{"panicked": r.Panicked}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "Invoke panics -- nil receiver panic", actual)
}

func Test_C2_Invoke_WithNonNilArgs(t *testing.T) {
	fn := func(a, b string) string { return a + b }
	r := results.InvokeWithPanicRecovery(fn, nil, "hello", " world")
	// fn is a plain func, nil receiver doesn't apply → might panic depending on call
	_ = r
}

func Test_C2_Invoke_FuncReturningOnlyError(t *testing.T) {
	// Arrange
	fn := func() error { return errors.New("fail") }
	r := results.InvokeWithPanicRecovery(fn, nil)

	// Act
	actual := args.Map{
		"hasError": r.HasError(),
		"returnCount": r.ReturnCount,
	}

	// Assert
	expected := args.Map{
		"hasError": true,
		"returnCount": 1,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns error -- func returning only error", actual)
}

func Test_C2_Invoke_FuncReturningNilError(t *testing.T) {
	// Arrange
	fn := func() error { return nil }
	r := results.InvokeWithPanicRecovery(fn, nil)

	// Act
	actual := args.Map{
		"hasError": r.HasError(),
		"returnCount": r.ReturnCount,
	}

	// Assert
	expected := args.Map{
		"hasError": false,
		"returnCount": 1,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns nil -- func returning nil error", actual)
}

func Test_C2_Invoke_FuncReturningIntAndNonError(t *testing.T) {
	// Arrange
	fn := func() (int, string) { return 42, "not-error" }
	r := results.InvokeWithPanicRecovery(fn, nil)

	// Act
	actual := args.Map{
		"panicked":    r.Panicked,
		"returnCount": r.ReturnCount,
		"hasError":    r.HasError(),
	}

	// Assert
	expected := args.Map{
		"panicked": false,
		"returnCount": 2,
		"hasError": false,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns error -- func with non-error second return", actual)
}

func Test_C2_Invoke_PanicFunc(t *testing.T) {
	// Arrange
	fn := func() { panic("deliberate") }
	r := results.InvokeWithPanicRecovery(fn, nil)

	// Act
	actual := args.Map{
		"panicked": r.Panicked,
		"panicVal": fmt.Sprintf("%v", r.PanicValue),
	}

	// Assert
	expected := args.Map{
		"panicked": true,
		"panicVal": "deliberate",
	}
	expected.ShouldBeEqual(t, 0, "Invoke panics -- panic func", actual)
}

// ── ShouldMatchResult with explicit compareFields ──

func Test_C2_ShouldMatchResult_ExplicitFields(t *testing.T) {
	r := results.Result[string]{Value: "ok", ReturnCount: 2}
	exp := results.ResultAny{Value: "ok", ReturnCount: 2}
	r.ShouldMatchResult(t, 0, "explicit fields", exp, "value", "returnCount", "panicked")
}

// ── ShouldMatchResult auto-derive with error ──

func Test_C2_ShouldMatchResult_WithError(t *testing.T) {
	r := results.Result[string]{Value: "v", Error: errors.New("e"), ReturnCount: 1}
	exp := results.ResultAny{Value: "v", Error: results.ExpectAnyError, ReturnCount: 1}
	r.ShouldMatchResult(t, 0, "with error", exp)
}

// ── filterByFields missing key branch ──

func Test_C2_ShouldMatchResult_MissingFieldKey(t *testing.T) {
	r := results.Result[string]{Value: "ok"}
	exp := results.ResultAny{}
	// "isSafe" is not auto-derived so we force it
	r.ShouldMatchResult(t, 0, "missing field", exp, "panicked", "isSafe")
}

// ── MethodName with method expression ──

func Test_C2_MethodName_MethodExpression(t *testing.T) {
	// Arrange
	name := results.MethodName((*testReceiver).Greet)
	// Should extract "Greet" from the full name

	// Act
	actual := args.Map{"result": name == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty method name", actual)
}

// ── ToMap and ToMapCompact completeness ──

func Test_C2_Result_ToMap_AllFields(t *testing.T) {
	// Arrange
	r := results.Result[string]{
		Value:       "test",
		Panicked:    false,
		Error:       errors.New("e"),
		ReturnCount: 1,
	}
	m := r.ToMap()

	// Act
	actual := args.Map{"result": m["value"] != "test" || m["panicked"] != false || m["hasError"] != true || m["returnCount"] != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected map values", actual)
}

func Test_C2_Result_ToMapCompact_Fields(t *testing.T) {
	// Arrange
	r := results.Result[int]{Value: 5, Panicked: true}
	m := r.ToMapCompact()

	// Act
	actual := args.Map{"result": m["value"] != "5" || m["panicked"] != true}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected compact map", actual)
}

// ── ResultAt edge cases ──

func Test_C2_Result_ResultAt_ValidIndex(t *testing.T) {
	// Arrange
	r := results.Result[string]{AllResults: []any{"a", "b", "c"}}

	// Act
	actual := args.Map{"result": r.ResultAt(2) != "c"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected c", actual)
}

// ── IsResult2 and Result2String ──

func Test_C2_Results_IsResult2_NoMatch(t *testing.T) {
	// Arrange
	r := results.Results[string, int]{Result2: 10}

	// Act
	actual := args.Map{"result": r.IsResult2(20)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not match", actual)
}
