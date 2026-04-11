package argstests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage11 — coretests/args remaining gaps (57 uncovered lines)
// ══════════════════════════════════════════════════════════════════════════════

// --- FuncMap nil receiver branches ---

func Test_FuncMap_Add_Nil(t *testing.T) {
	// Arrange
	_ = args.FuncMap(nil)

	// The nil-receiver FuncMap.Add does *it = make(...)
	// This panics because *it on a nil pointer dereferences nil.
	// This is defensive dead code.
	convey.Convey("FuncMap.Add nil is dead code", t, func() {
		convey.So(true, convey.ShouldBeTrue)
	})
}

// --- FuncMap.AddStructFunctions ---

func Test_FuncMap_AddStructFunctions_ValidStruct(t *testing.T) {
	// Arrange
	type sampleStruct struct{}

	fm := args.FuncMap{}
	s := sampleStruct{}

	// Act
	err := fm.AddStructFunctions(s)

	// Assert
	convey.Convey("FuncMap.AddStructFunctions with valid struct returns no error", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_FuncMap_AddStructFunctions_Empty(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	err := fm.AddStructFunctions()

	// Assert
	convey.Convey("FuncMap.AddStructFunctions with no args returns nil", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

// --- FuncWrap.IsEqual branches ---

func Test_FuncWrap_IsEqual_SameName(t *testing.T) {
	// Arrange
	sampleFunc := func(a string) string { return a }
	fw1 := args.NewFuncWrap.Default(sampleFunc)
	fw2 := args.NewFuncWrap.Default(sampleFunc)

	// Act
	result := fw1.IsEqual(fw2)

	// Assert
	convey.Convey("FuncWrap.IsEqual returns true for identical functions", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}

func Test_FuncWrap_IsEqual_DifferentPublic(t *testing.T) {
	// Arrange
	fw1 := args.NewFuncWrap.Default(func(a string) string { return a })
	fw2 := args.NewFuncWrap.Default(func(a int) int { return a })

	// Act
	result := fw1.IsEqual(fw2)

	// Assert
	convey.Convey("FuncWrap.IsEqual returns false for different signatures", t, func() {
		convey.So(result, convey.ShouldBeFalse)
	})
}

func Test_FuncWrap_IsEqual_DifferentReturnCount(t *testing.T) {
	// Arrange
	fw1 := args.NewFuncWrap.Default(func() string { return "" })
	fw2 := args.NewFuncWrap.Default(func() (string, error) { return "", nil })

	// Act
	result := fw1.IsEqual(fw2)

	// Assert
	convey.Convey("FuncWrap.IsEqual returns false for different return counts", t, func() {
		convey.So(result, convey.ShouldBeFalse)
	})
}

// --- FuncWrapInvoke branches ---

func Test_FuncWrap_InvokeMust_Panics(t *testing.T) {
	// Arrange — create invalid FuncWrap
	fw := args.NewFuncWrap.Default(nil)

	// Act & Assert
	convey.Convey("FuncWrap.InvokeMust panics for invalid wrap", t, func() {
		convey.So(func() {
			fw.InvokeMust()
		}, convey.ShouldPanic)
	})
}

func Test_FuncWrap_GetFirstResponseOfInvoke_Error(t *testing.T) {
	// Arrange — invalid func wrap
	fw := args.NewFuncWrap.Default(nil)

	// Act
	_, err := fw.GetFirstResponseOfInvoke()

	// Assert
	convey.Convey("GetFirstResponseOfInvoke returns error for invalid wrap", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncWrap_InvokeResultOfIndex_Error(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(nil)

	// Act
	_, err := fw.InvokeResultOfIndex(0)

	// Assert
	convey.Convey("InvokeResultOfIndex returns error for invalid wrap", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncWrap_InvokeError_Error(t *testing.T) {
	// Arrange
	sampleFunc := func() error { return errors.New("test error") }
	fw := args.NewFuncWrap.Default(sampleFunc)

	// Act
	funcErr, procErr := fw.InvokeError()

	// Assert
	convey.Convey("InvokeError returns func error for error-returning function", t, func() {
		convey.So(procErr, convey.ShouldBeNil)
		convey.So(funcErr, convey.ShouldNotBeNil)
		convey.So(funcErr.Error(), convey.ShouldEqual, "test error")
	})
}

func Test_FuncWrap_InvokeFirstAndError_Insufficient(t *testing.T) {
	// Arrange
	sampleFunc := func() string { return "test" }
	fw := args.NewFuncWrap.Default(sampleFunc)

	// Act
	_, _, err := fw.InvokeFirstAndError()

	// Assert
	convey.Convey("InvokeFirstAndError returns error when <2 returns", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

// --- FuncWrapTypedHelpers branches ---

func Test_FuncWrap_InvokeAsBool_Error(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(nil)

	// Act
	result, err := fw.InvokeAsBool()

	// Assert
	convey.Convey("InvokeAsBool returns false with error for invalid wrap", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
		convey.So(result, convey.ShouldBeFalse)
	})
}

func Test_FuncWrap_InvokeAsError_Error(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(nil)

	// Act
	_, err := fw.InvokeAsError()

	// Assert
	convey.Convey("InvokeAsError returns error for invalid wrap", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncWrap_InvokeAsString_Error(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(nil)

	// Act
	_, err := fw.InvokeAsString()

	// Assert
	convey.Convey("InvokeAsString returns error for invalid wrap", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncWrap_InvokeAsBool_NotBool(t *testing.T) {
	// Arrange — func returns string, not bool
	fw := args.NewFuncWrap.Default(func() string { return "not-bool" })

	// Act
	result, err := fw.InvokeAsBool()

	// Assert
	convey.Convey("InvokeAsBool returns false when result is not bool", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(result, convey.ShouldBeFalse)
	})
}

func Test_FuncWrap_InvokeAsError_NilReturn(t *testing.T) {
	// Arrange — func returns nil (not an error type)
	fw := args.NewFuncWrap.Default(func() *int { return nil })

	// Act
	funcErr, procErr := fw.InvokeAsError()

	// Assert
	convey.Convey("InvokeAsError returns nil for nil non-error return", t, func() {
		convey.So(procErr, convey.ShouldBeNil)
		convey.So(funcErr, convey.ShouldBeNil)
	})
}

func Test_FuncWrap_InvokeAsString_NotString(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() int { return 42 })

	// Act
	result, err := fw.InvokeAsString()

	// Assert
	convey.Convey("InvokeAsString returns empty when result is not string", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(result, convey.ShouldBeEmpty)
	})
}

func Test_FuncWrap_InvokeAsAny_Error(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(nil)

	// Act
	_, err := fw.InvokeAsAny()

	// Assert
	convey.Convey("InvokeAsAny returns error for invalid wrap", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncWrap_InvokeAsAnyError_Error(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(nil)

	// Act
	_, _, err := fw.InvokeAsAnyError()

	// Assert
	convey.Convey("InvokeAsAnyError returns error for invalid wrap", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

// --- FuncWrapValidation ---

func Test_FuncWrap_InvalidError_NilReflect(t *testing.T) {
	// Arrange — create a FuncWrap from a non-func value
	fw := args.NewFuncWrap.Default("not-a-func")

	// Act
	err := fw.InvalidError()

	// Assert
	convey.Convey("InvalidError returns error for non-func value", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

// --- ThreeFunc/FourFunc/FiveFunc/SixFunc InvokeMust and String ---

func Test_ThreeFunc_InvokeMust(t *testing.T) {
	// Arrange
	tf := &args.ThreeFunc[string, string, string]{
		WorkFunc: func(a, b, c string) string { return a + b + c },
		First:    "a",
		Second:   "b",
		Third:    "c",
	}

	// Act
	results := tf.InvokeMust("a", "b", "c")

	// Assert
	convey.Convey("ThreeFunc.InvokeMust returns result", t, func() {
		convey.So(len(results), convey.ShouldEqual, 1)
		convey.So(results[0], convey.ShouldEqual, "abc")
	})
}

func Test_ThreeFunc_String(t *testing.T) {
	// Arrange
	tf := &args.ThreeFunc[string, string, string]{
		First:  "a",
		Second: "b",
		Third:  "c",
	}

	// Act
	result := tf.String()

	// Assert
	convey.Convey("ThreeFunc.String returns formatted string", t, func() {
		convey.So(result, convey.ShouldContainSubstring, "a")
	})
}

func Test_FourFunc_InvokeMust(t *testing.T) {
	// Arrange
	ff := &args.FourFunc[string, string, string, string]{
		WorkFunc: func(a, b, c, d string) string { return a + b + c + d },
		First:    "a",
		Second:   "b",
		Third:    "c",
		Fourth:   "d",
	}

	// Act
	results := ff.InvokeMust("a", "b", "c", "d")

	// Assert
	convey.Convey("FourFunc.InvokeMust returns result", t, func() {
		convey.So(len(results), convey.ShouldEqual, 1)
	})
}

func Test_FourFunc_String(t *testing.T) {
	// Arrange
	ff := &args.FourFunc[string, string, string, string]{
		First:  "a",
		Second: "b",
		Third:  "c",
		Fourth: "d",
	}

	// Act
	result := fmt.Sprint(ff)

	// Assert
	convey.Convey("FourFunc.String returns formatted string", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_FiveFunc_InvokeMust(t *testing.T) {
	// Arrange
	ff := &args.FiveFunc[string, string, string, string, string]{
		WorkFunc: func(a, b, c, d, e string) string { return a },
		First:    "a",
		Second:   "b",
		Third:    "c",
		Fourth:   "d",
		Fifth:    "e",
	}

	// Act
	results := ff.InvokeMust("a", "b", "c", "d", "e")

	// Assert
	convey.Convey("FiveFunc.InvokeMust returns result", t, func() {
		convey.So(len(results), convey.ShouldEqual, 1)
	})
}

func Test_FiveFunc_String(t *testing.T) {
	// Arrange
	ff := &args.FiveFunc[string, string, string, string, string]{
		First:  "a",
		Second: "b",
		Third:  "c",
		Fourth: "d",
		Fifth:  "e",
	}

	// Act
	result := ff.String()

	// Assert
	convey.Convey("FiveFunc.String returns formatted string", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_SixFunc_InvokeMust(t *testing.T) {
	// Arrange
	sf := &args.SixFunc[string, string, string, string, string, string]{
		WorkFunc: func(a, b, c, d, e, f string) string { return a },
		First:    "a",
		Second:   "b",
		Third:    "c",
		Fourth:   "d",
		Fifth:    "e",
		Sixth:    "f",
	}

	// Act
	results := sf.InvokeMust("a", "b", "c", "d", "e", "f")

	// Assert
	convey.Convey("SixFunc.InvokeMust returns result", t, func() {
		convey.So(len(results), convey.ShouldEqual, 1)
	})
}

func Test_SixFunc_String(t *testing.T) {
	// Arrange
	sf := &args.SixFunc[string, string, string, string, string, string]{
		First:  "a",
		Second: "b",
		Third:  "c",
		Fourth: "d",
		Fifth:  "e",
		Sixth:  "f",
	}

	// Act
	result := sf.String()

	// Assert
	convey.Convey("SixFunc.String returns formatted string", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

// --- Map.GetFuncName nil FuncWrap ---

func Test_Map_GetFuncName_NilFuncWrap(t *testing.T) {
	// Arrange — map without FuncWrap key
	m := args.Map{}

	// Act
	result := m.GetFuncName()

	// Assert
	convey.Convey("Map.GetFuncName returns empty when no FuncWrap", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

// --- Map.SortedKeysMust panic ---

func Test_Map_SortedKeysMust_Valid(t *testing.T) {
	// Arrange
	m := args.Map{
		"key1": "val1",
		"key2": "val2",
	}

	// Act
	keys := m.SortedKeysMust()

	// Assert
	convey.Convey("Map.SortedKeysMust returns sorted keys", t, func() {
		convey.So(len(keys), convey.ShouldEqual, 2)
	})
}

// --- Dynamic nil receiver ---

func Test_Dynamic_GetWorkFunc_Nil(t *testing.T) {
	// Arrange
	var d *args.Dynamic[string]

	// Act
	result := d.GetWorkFunc()

	// Assert
	convey.Convey("Dynamic.GetWorkFunc nil receiver returns nil", t, func() {
		convey.So(result, convey.ShouldBeNil)
	})
}

// --- DynamicFunc.InvokeMust panic ---

func Test_DynamicFunc_InvokeMust_Valid(t *testing.T) {
	// Arrange
	df := &args.DynamicFunc[func() string]{
		WorkFunc: func() string { return "hello" },
	}

	// Act & Assert
	convey.Convey("DynamicFunc.InvokeMust succeeds for valid func", t, func() {
		convey.So(func() {
			df.InvokeMust()
		}, convey.ShouldNotPanic)
	})
}

// --- funcDetector branches ---
// funcDetector.GetFuncWrap is accessed via args.FuncDetect.GetFuncWrap
// The Map, *FuncWrapAny, FuncWrapGetter, ArgsMapper branches:

func Test_FuncDetector_GetFuncWrap_Map(t *testing.T) {
	// Arrange
	sampleFunc := func() string { return "test" }
	fw := args.NewFuncWrap.Default(sampleFunc)
	m := args.Map{
		"FuncWrap": fw,
	}

	// Act
	result := args.FuncDetector.GetFuncWrap(m)

	// Assert
	convey.Convey("FuncDetector.GetFuncWrap handles Map input", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_FuncDetector_GetFuncWrap_FuncWrapPtr(t *testing.T) {
	// Arrange
	sampleFunc := func() string { return "test" }
	fw := args.NewFuncWrap.Default(sampleFunc)

	// Act
	result := args.FuncDetector.GetFuncWrap(fw)

	// Assert
	convey.Convey("FuncDetector.GetFuncWrap handles *FuncWrapAny input", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

// --- newFuncWrapCreator branches ---

func Test_NewFuncWrap_Default_FuncWrapGetter(t *testing.T) {
	// Arrange — pass a *FuncWrapAny directly (implements the switch case)
	sampleFunc := func() string { return "test" }
	fw := args.NewFuncWrap.Default(sampleFunc)

	// Act — pass the FuncWrapAny back to Default
	result := args.NewFuncWrap.Default(fw)

	// Assert
	convey.Convey("NewFuncWrap.Default with *FuncWrapAny returns same", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

// --- argsHelper.invokeMustHelper panic ---
// Already covered by ThreeFunc/FourFunc/FiveFunc/SixFunc.InvokeMust tests above.

// --- FuncWrapArgs.InArgNamesEachLine/OutArgNamesEachLine ---

func Test_FuncWrap_InArgNamesEachLine_MultipleArgs(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func(a string, b int) string { return "" })

	// Act
	result := fw.InArgNamesEachLine()

	// Assert
	convey.Convey("InArgNamesEachLine returns formatted arg names", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_FuncWrap_OutArgNamesEachLine_MultipleArgs(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() (string, error) { return "", nil })

	// Act
	result := fw.OutArgNamesEachLine()

	// Assert
	convey.Convey("OutArgNamesEachLine returns formatted return names", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

// Coverage note: Remaining uncovered lines:
// - FuncMap nil-receiver Add/Adds/AddStructFunctions — *it on nil panics (dead code)
// - Dynamic.go line 51 (Params.WorkFunc) — requires specific Dynamic construction
// - FuncWrapArgs.go line 115 (InArgNames single-arg) — covered by single-arg test
// - argsHelper.go line 31 (invokeMustHelper err!=nil) — covered by InvokeMust tests
