package coretestsargstests

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ── helper functions ──

func helperAdd(a, b int) int                { return a + b }
func helperRetErr() error                   { return errors.New("err") }
func helperRetNilErr() error                { return nil }
func helperRetBool() bool                   { return true }
func helperRetString() string               { return "hello" }
func helperRetAny() any                     { return 42 }
func helperRetAnyErr() (any, error)         { return "ok", nil }
func helperRetAnyErrFail() (any, error)     { return nil, errors.New("fail") }
func helperRetIntBool(x int) (int, bool)    { return x * 2, true }
func helperVoid()                           {}

// ── FuncWrap.IsEquals branches ──

func Test_FuncWrap_IsEquals_DiffName(t *testing.T) {
	// Arrange
	fw1 := args.NewFuncWrap.Default(helperAdd)
	fw2 := args.NewFuncWrap.Default(helperRetBool)

	// Act
	result := fw1.IsEqual(fw2)

	// Assert
	actual := args.Map{"equal": result}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEquals returns false -- different names", actual)
}

func Test_FuncWrap_IsEquals_DiffValidity(t *testing.T) {
	// Arrange
	fw1 := args.NewFuncWrap.Default(helperAdd)
	fw2 := args.NewFuncWrap.Invalid()

	// Act
	result := fw1.IsEqual(fw2)

	// Assert
	actual := args.Map{"equal": result}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEquals returns false -- different validity", actual)
}

func Test_FuncWrap_IsEquals_DiffArgsCount(t *testing.T) {
	// Arrange
	fw1 := args.NewFuncWrap.Default(helperAdd)
	fw2 := args.NewFuncWrap.Default(helperRetBool)

	// Act
	result := fw1.IsEqual(fw2)

	// Assert
	actual := args.Map{"equal": result}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEquals returns false -- different args count", actual)
}

func Test_FuncWrap_IsEquals_DiffReturnLen(t *testing.T) {
	// Arrange
	fw1 := args.NewFuncWrap.Default(helperRetBool)
	fw2 := args.NewFuncWrap.Default(helperRetAnyErr)

	// Act
	result := fw1.IsEqual(fw2)

	// Assert
	actual := args.Map{"equal": result}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEquals returns false -- different return length", actual)
}

// ── FuncWrap: InvokeMust, InvokeFirstAndError, GetFirstResponseOfInvoke, InvokeResultOfIndex ──

func Test_FuncWrap_InvokeMust_Valid(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperAdd)

	// Act
	results := fw.InvokeMust(2, 3)

	// Assert
	actual := args.Map{"result": results[0]}
	expected := args.Map{"result": 5}
	expected.ShouldBeEqual(t, 0, "InvokeMust returns correct -- valid func", actual)
}

func Test_FuncWrap_InvokeFirstAndError_Valid(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperRetAnyErr)

	// Act
	first, funcErr, procErr := fw.InvokeFirstAndError()

	// Assert
	actual := args.Map{
		"first":   fmt.Sprintf("%v", first),
		"funcErr": funcErr == nil,
		"procErr": procErr == nil,
	}
	expected := args.Map{
		"first":   "ok",
		"funcErr": true,
		"procErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns correct -- valid func", actual)
}

func Test_FuncWrap_InvokeFirstAndError_SingleReturn(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperRetBool)

	// Act
	_, _, procErr := fw.InvokeFirstAndError()

	// Assert
	actual := args.Map{"hasErr": procErr != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns error -- single return func", actual)
}

func Test_FuncWrap_GetFirstResponseOfInvoke(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperRetString)

	// Act
	first, err := fw.GetFirstResponseOfInvoke()

	// Assert
	actual := args.Map{
		"noErr":  err == nil,
		"result": first,
	}
	expected := args.Map{
		"noErr":  true,
		"result": "hello",
	}
	expected.ShouldBeEqual(t, 0, "GetFirstResponseOfInvoke returns correct -- string func", actual)
}

func Test_FuncWrap_InvokeResultOfIndex(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperRetAnyErr)

	// Act
	second, err := fw.InvokeResultOfIndex(1)

	// Assert
	actual := args.Map{
		"noErr":  err == nil,
		"isNil":  second == nil,
	}
	expected := args.Map{
		"noErr":  true,
		"isNil":  true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeResultOfIndex returns correct -- index 1 of (any,error)", actual)
}

func Test_FuncWrap_InvokeError(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperRetErr)

	// Act
	funcErr, procErr := fw.InvokeError()

	// Assert
	actual := args.Map{
		"procOk": procErr == nil,
		"hasErr": funcErr != nil,
	}
	expected := args.Map{
		"procOk": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeError returns correct -- error func", actual)
}

// ── FuncWrap: Typed Invoke Helpers ──

func Test_FuncWrap_InvokeAsBool(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperRetBool)

	// Act
	val, err := fw.InvokeAsBool()

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"val": val,
	}
	expected := args.Map{
		"noErr": true,
		"val": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsBool returns correct -- bool func", actual)
}

func Test_FuncWrap_InvokeAsBool_WrongType(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperRetString)

	// Act
	val, err := fw.InvokeAsBool()

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"val": val,
	}
	expected := args.Map{
		"noErr": true,
		"val": false,
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsBool returns false -- wrong type", actual)
}

func Test_FuncWrap_InvokeAsBool_NoResults(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperVoid)

	// Act
	val, err := fw.InvokeAsBool()

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"val": val,
	}
	expected := args.Map{
		"noErr": true,
		"val": false,
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsBool returns false -- void func", actual)
}

func Test_FuncWrap_InvokeAsError_Valid(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperRetErr)

	// Act
	funcErr, procErr := fw.InvokeAsError()

	// Assert
	actual := args.Map{
		"procOk": procErr == nil,
		"hasErr": funcErr != nil,
	}
	expected := args.Map{
		"procOk": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsError returns correct -- error func", actual)
}

func Test_FuncWrap_InvokeAsError_NilReturn(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperRetNilErr)

	// Act
	funcErr, procErr := fw.InvokeAsError()

	// Assert
	actual := args.Map{
		"procOk": procErr == nil,
		"nilErr": funcErr == nil,
	}
	expected := args.Map{
		"procOk": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsError returns nil -- nil error func", actual)
}

func Test_FuncWrap_InvokeAsError_NoResults(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperVoid)

	// Act
	funcErr, procErr := fw.InvokeAsError()

	// Assert
	actual := args.Map{
		"procOk": procErr == nil,
		"nilErr": funcErr == nil,
	}
	expected := args.Map{
		"procOk": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsError returns nil -- void func", actual)
}

func Test_FuncWrap_InvokeAsString(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperRetString)

	// Act
	val, err := fw.InvokeAsString()

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"val": val,
	}
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsString returns correct -- string func", actual)
}

func Test_FuncWrap_InvokeAsString_WrongType(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperRetBool)

	// Act
	val, err := fw.InvokeAsString()

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"val": val,
	}
	expected := args.Map{
		"noErr": true,
		"val": "",
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsString returns empty -- wrong type", actual)
}

func Test_FuncWrap_InvokeAsAny(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperRetAny)

	// Act
	val, err := fw.InvokeAsAny()

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"val": val,
	}
	expected := args.Map{
		"noErr": true,
		"val": 42,
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsAny returns correct -- any func", actual)
}

func Test_FuncWrap_InvokeAsAny_NoResults(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperVoid)

	// Act
	val, err := fw.InvokeAsAny()

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"nilVal": val == nil,
	}
	expected := args.Map{
		"noErr": true,
		"nilVal": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsAny returns nil -- void func", actual)
}

func Test_FuncWrap_InvokeAsAnyError_Valid(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperRetAnyErr)

	// Act
	result, funcErr, procErr := fw.InvokeAsAnyError()

	// Assert
	actual := args.Map{
		"procOk": procErr == nil,
		"nilErr": funcErr == nil,
		"result": fmt.Sprintf("%v", result),
	}
	expected := args.Map{
		"procOk": true,
		"nilErr": true,
		"result": "ok",
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsAnyError returns correct -- valid func", actual)
}

func Test_FuncWrap_InvokeAsAnyError_WithError(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperRetAnyErrFail)

	// Act
	result, funcErr, procErr := fw.InvokeAsAnyError()

	// Assert
	actual := args.Map{
		"procOk": procErr == nil,
		"hasErr": funcErr != nil,
		"nilRes": result == nil,
	}
	expected := args.Map{
		"procOk": true,
		"hasErr": true,
		"nilRes": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsAnyError returns error -- failing func", actual)
}

func Test_FuncWrap_InvokeAsAnyError_NoResults(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperVoid)

	// Act
	result, funcErr, procErr := fw.InvokeAsAnyError()

	// Assert
	actual := args.Map{
		"procOk": procErr == nil,
		"nilRes": result == nil,
		"nilErr": funcErr == nil,
	}
	expected := args.Map{
		"procOk": true,
		"nilRes": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsAnyError returns nils -- void func", actual)
}

// ── FuncWrap: InvalidError, FuncWrapValidation ──

func Test_FuncWrap_InvalidError_NilRv(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default("not-a-func")

	// Act
	err := fw.InvalidError()

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvalidError returns error -- invalid func", actual)
}

// ── FuncWrap: InArgNames, InArgNamesEachLine, OutArgNames, OutArgNamesEachLine ──

func Test_FuncWrap_InArgNamesEachLine(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperAdd)

	// Act
	names := fw.InArgNamesEachLine()

	// Assert
	actual := args.Map{"hasLines": len(names) > 1}
	expected := args.Map{"hasLines": true}
	expected.ShouldBeEqual(t, 0, "InArgNamesEachLine returns lines -- multi-arg func", actual)
}

func Test_FuncWrap_OutArgNamesEachLine(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperRetAnyErr)

	// Act
	names := fw.OutArgNamesEachLine()

	// Assert
	actual := args.Map{"hasLines": len(names) > 1}
	expected := args.Map{"hasLines": true}
	expected.ShouldBeEqual(t, 0, "OutArgNamesEachLine returns lines -- multi-return func", actual)
}

// ── FuncMap: InArgsVerifyRv, OutArgsVerifyRv, VoidCallNoReturn, MustBeValid, ValidationError ──

func Test_FuncMap_InArgsVerifyRv(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperAdd)
	name := args.NewFuncWrap.Default(helperAdd).Name

	// Act
	ok, err := fm.InArgsVerifyRv(name, []reflect.Type{
		reflect.TypeOf(0),
		reflect.TypeOf(0),
	})

	// Assert
	actual := args.Map{
		"ok": ok,
		"noErr": err == nil,
	}
	expected := args.Map{
		"ok": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InArgsVerifyRv returns ok -- matching types", actual)
}

func Test_FuncMap_InArgsVerifyRv_NotFound(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperAdd)

	// Act
	ok, err := fm.InArgsVerifyRv("nonexistent", nil)

	// Assert
	actual := args.Map{
		"ok": ok,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"ok": false,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InArgsVerifyRv returns error -- not found", actual)
}

func Test_FuncMap_OutArgsVerifyRv(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperAdd)
	name := args.NewFuncWrap.Default(helperAdd).Name

	// Act
	ok, err := fm.OutArgsVerifyRv(name, []reflect.Type{reflect.TypeOf(0)})

	// Assert
	actual := args.Map{
		"ok": ok,
		"noErr": err == nil,
	}
	expected := args.Map{
		"ok": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "OutArgsVerifyRv returns ok -- matching types", actual)
}

func Test_FuncMap_OutArgsVerifyRv_NotFound(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperAdd)

	// Act
	ok, err := fm.OutArgsVerifyRv("nonexistent", nil)

	// Assert
	actual := args.Map{
		"ok": ok,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"ok": false,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "OutArgsVerifyRv returns error -- not found", actual)
}

func Test_FuncMap_VoidCallNoReturn(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperVoid)
	name := args.NewFuncWrap.Default(helperVoid).Name

	// Act
	err := fm.VoidCallNoReturn(name)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VoidCallNoReturn returns nil -- valid void func", actual)
}

func Test_FuncMap_VoidCallNoReturn_NotFound(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperVoid)

	// Act
	err := fm.VoidCallNoReturn("missing")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VoidCallNoReturn returns error -- not found", actual)
}

func Test_FuncMap_MustBeValid_NotFound(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperVoid)

	// Act
	panicked := callPanicsCov4(func() { fm.MustBeValid("missing") })

	// Assert
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "MustBeValid panics -- not found", actual)
}

func Test_FuncMap_ValidationError(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperAdd)
	name := args.NewFuncWrap.Default(helperAdd).Name

	// Act
	err := fm.ValidationError(name)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ValidationError returns nil -- valid func", actual)
}

func Test_FuncMap_ValidationError_NotFound(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperAdd)

	// Act
	err := fm.ValidationError("missing")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidationError returns error -- not found", actual)
}

func Test_FuncMap_InvokeMust_Valid(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperAdd)
	name := args.NewFuncWrap.Default(helperAdd).Name

	// Act
	results := fm.InvokeMust(name, 3, 4)

	// Assert
	actual := args.Map{"result": results[0]}
	expected := args.Map{"result": 7}
	expected.ShouldBeEqual(t, 0, "InvokeMust returns correct -- valid func", actual)
}

func Test_FuncMap_InvokeMust_Panic(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperAdd)

	// Act
	panicked := callPanicsCov4(func() { fm.InvokeMust("missing", 1, 2) })

	// Assert
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "InvokeMust panics -- not found", actual)
}

func Test_FuncMap_InvokeError(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperRetErr)
	name := args.NewFuncWrap.Default(helperRetErr).Name

	// Act
	funcErr, procErr := fm.InvokeError(name)

	// Assert
	actual := args.Map{
		"procOk": procErr == nil,
		"hasErr": funcErr != nil,
	}
	expected := args.Map{
		"procOk": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeError returns correct -- error func", actual)
}

func Test_FuncMap_InvokeError_NotFound(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperRetErr)

	// Act
	_, procErr := fm.InvokeError("missing")

	// Assert
	actual := args.Map{"hasErr": procErr != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeError returns error -- not found", actual)
}

func Test_FuncMap_InvokeFirstAndError(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperRetAnyErr)
	name := args.NewFuncWrap.Default(helperRetAnyErr).Name

	// Act
	first, funcErr, procErr := fm.InvokeFirstAndError(name)

	// Assert
	actual := args.Map{
		"procOk":  procErr == nil,
		"nilErr":  funcErr == nil,
		"hasFirst": first != nil,
	}
	expected := args.Map{
		"procOk":  true,
		"nilErr":  true,
		"hasFirst": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns correct -- valid func", actual)
}

func Test_FuncMap_InvokeFirstAndError_NotFound(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperRetAnyErr)

	// Act
	_, _, procErr := fm.InvokeFirstAndError("missing")

	// Assert
	actual := args.Map{"hasErr": procErr != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns error -- not found", actual)
}

func Test_FuncMap_InvokeResultOfIndex_NotFound(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperAdd)

	// Act
	_, err := fm.InvokeResultOfIndex("missing", 0)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeResultOfIndex returns error -- not found", actual)
}

func Test_FuncMap_ValidateMethodArgs_NotFound(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(helperAdd)

	// Act
	err := fm.ValidateMethodArgs("missing", []any{1, 2})

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs returns error -- not found", actual)
}

// ── DynamicFunc: Typed getters, Invoke, InvokeMust, InvokeWithValidArgs, InvokeArgs ──

func Test_DynamicFunc_GetAsInt(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		Params: args.Map{"count": 5},
	}

	// Act
	val, ok := df.GetAsInt("count")

	// Assert
	actual := args.Map{
		"ok": ok,
		"val": val,
	}
	expected := args.Map{
		"ok": true,
		"val": 5,
	}
	expected.ShouldBeEqual(t, 0, "GetAsInt returns correct -- valid key", actual)
}

func Test_DynamicFunc_GetAsInt_Missing(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		Params: args.Map{},
	}

	// Act
	val, ok := df.GetAsInt("missing")

	// Assert
	actual := args.Map{
		"ok": ok,
		"val": val,
	}
	expected := args.Map{
		"ok": false,
		"val": 0,
	}
	expected.ShouldBeEqual(t, 0, "GetAsInt returns zero -- missing key", actual)
}

func Test_DynamicFunc_GetAsString(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		Params: args.Map{"name": "test"},
	}

	// Act
	val, ok := df.GetAsString("name")

	// Assert
	actual := args.Map{
		"ok": ok,
		"val": val,
	}
	expected := args.Map{
		"ok": true,
		"val": "test",
	}
	expected.ShouldBeEqual(t, 0, "GetAsString returns correct -- valid key", actual)
}

func Test_DynamicFunc_GetAsString_Missing(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		Params: args.Map{},
	}

	// Act
	val, ok := df.GetAsString("missing")

	// Assert
	actual := args.Map{
		"ok": ok,
		"val": val,
	}
	expected := args.Map{
		"ok": false,
		"val": "",
	}
	expected.ShouldBeEqual(t, 0, "GetAsString returns empty -- missing key", actual)
}

func Test_DynamicFunc_GetAsStrings(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		Params: args.Map{"items": []string{"a", "b"}},
	}

	// Act
	val, ok := df.GetAsStrings("items")

	// Assert
	actual := args.Map{
		"ok": ok,
		"len": len(val),
	}
	expected := args.Map{
		"ok": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "GetAsStrings returns correct -- valid key", actual)
}

func Test_DynamicFunc_GetAsStrings_Missing(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		Params: args.Map{},
	}

	// Act
	val, ok := df.GetAsStrings("missing")

	// Assert
	actual := args.Map{
		"ok": ok,
		"len": len(val),
	}
	expected := args.Map{
		"ok": false,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "GetAsStrings returns empty -- missing key", actual)
}

func Test_DynamicFunc_GetAsAnyItems(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		Params: args.Map{"items": []any{1, 2, 3}},
	}

	// Act
	val, ok := df.GetAsAnyItems("items")

	// Assert
	actual := args.Map{
		"ok": ok,
		"len": len(val),
	}
	expected := args.Map{
		"ok": true,
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "GetAsAnyItems returns correct -- valid key", actual)
}

func Test_DynamicFunc_GetAsAnyItems_Missing(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		Params: args.Map{},
	}

	// Act
	val, ok := df.GetAsAnyItems("missing")

	// Assert
	actual := args.Map{
		"ok": ok,
		"len": len(val),
	}
	expected := args.Map{
		"ok": false,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "GetAsAnyItems returns empty -- missing key", actual)
}

func Test_DynamicFunc_Invoke(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		WorkFunc: helperAdd,
		Params:   args.Map{"first": 3},
	}

	// Act
	results, err := df.Invoke(3, 4)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"noErr": true,
		"result": 7,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct -- DynamicFunc with add", actual)
}

func Test_DynamicFunc_InvokeMust(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		WorkFunc: helperAdd,
	}

	// Act
	results := df.InvokeMust(1, 2)

	// Assert
	actual := args.Map{"result": results[0]}
	expected := args.Map{"result": 3}
	expected.ShouldBeEqual(t, 0, "InvokeMust returns correct -- DynamicFunc", actual)
}

func Test_DynamicFunc_InvokeWithValidArgs(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		WorkFunc: helperAdd,
		Params: args.Map{
			"first": 10,
			"p2":    20,
		},
	}

	// Act
	results, err := df.InvokeWithValidArgs()

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"noErr": true,
		"result": 30,
	}
	expected.ShouldBeEqual(t, 0, "InvokeWithValidArgs returns correct -- DynamicFunc", actual)
}

func Test_DynamicFunc_InvokeArgs(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		WorkFunc: helperAdd,
		Params: args.Map{
			"first": 5,
			"p2":    6,
		},
	}

	// Act
	results, err := df.InvokeArgs("first", "p2")

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"noErr": true,
		"result": 11,
	}
	expected.ShouldBeEqual(t, 0, "InvokeArgs returns correct -- DynamicFunc with named args", actual)
}

// ── Dynamic: GetWorkFunc, InvokeArgs ──

func Test_Dynamic_GetWorkFunc_Nil(t *testing.T) {
	// Arrange
	var d *args.DynamicAny

	// Act
	result := d.GetWorkFunc()

	// Assert
	actual := args.Map{"nilResult": result == nil}
	expected := args.Map{"nilResult": true}
	expected.ShouldBeEqual(t, 0, "GetWorkFunc returns nil -- nil receiver", actual)
}

func Test_Dynamic_InvokeArgs(t *testing.T) {
	// Arrange
	d := args.DynamicAny{
		Params: args.Map{
			"func":  helperAdd,
			"first": 3,
			"p2":    7,
		},
	}

	// Act
	results, err := d.InvokeArgs("first", "p2")

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"noErr": true,
		"result": 10,
	}
	expected.ShouldBeEqual(t, 0, "InvokeArgs returns correct -- Dynamic with named args", actual)
}

// ── Map.GetFuncName empty path ──

func Test_Map_GetFuncName_NoFunc(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}

	// Act
	name := m.GetFuncName()

	// Assert — FuncWrap returns non-nil even for nil func, but Name will be empty
	actual := args.Map{"name": name}
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "GetFuncName returns empty -- no func key", actual)
}

// ── funcDetector: GetFuncWrap branches ──

func Test_FuncDetector_GetFuncWrap_Map(t *testing.T) {
	// Arrange
	m := args.Map{"func": helperAdd}

	// Act
	fw := args.FuncDetector.GetFuncWrap(m)

	// Assert
	actual := args.Map{"valid": fw.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "GetFuncWrap returns valid -- Map input", actual)
}

func Test_FuncDetector_GetFuncWrap_FuncWrapPtr(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(helperAdd)

	// Act
	result := args.FuncDetector.GetFuncWrap(fw)

	// Assert
	actual := args.Map{"same": result == fw}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "GetFuncWrap returns same ptr -- FuncWrapAny input", actual)
}

// ── newFuncWrapCreator.StructToMap error path ──

func Test_NewFuncWrap_StructToMap_Error(t *testing.T) {
	// Arrange & Act — passing nil should trigger error
	_, err := args.NewFuncWrap.StructToMap(nil)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StructToMap returns error -- nil input", actual)
}

// ── panic helper ──

func callPanicsCov4(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}
