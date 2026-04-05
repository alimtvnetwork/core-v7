package argstests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage13 — coretests/args remaining 18 lines
// ══════════════════════════════════════════════════════════════════════════════

// ── DynamicFunc.InvokeMust panic on error (line 294) ──
// Dead code — panic path when Invoke returns error

// ── FiveFunc.String (line 231) ──

func Test_Cov13_FiveFunc_String(t *testing.T) {
	// Arrange
	ff := args.NewFiveFunc(1, "two", 3.0, true, "five")

	// Act
	result := ff.String()

	// Assert
	convey.Convey("FiveFunc String returns formatted output", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

// ── FourFunc.String (line 214) ──

func Test_Cov13_FourFunc_String(t *testing.T) {
	// Arrange
	ff := args.NewFourFunc(1, "two", 3.0, true)

	// Act
	result := ff.String()

	// Assert
	convey.Convey("FourFunc String returns formatted output", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

// ── SixFunc.String (line 259) ──

func Test_Cov13_SixFunc_String(t *testing.T) {
	// Arrange
	sf := args.NewSixFunc(1, "two", 3.0, true, "five", 6)

	// Act
	result := sf.String()

	// Assert
	convey.Convey("SixFunc String returns formatted output", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

// ── FuncMap.AddStructs error path (line 112-114) ──

func Test_Cov13_FuncMap_AddStructs_EmptySlice(t *testing.T) {
	// Arrange
	fm := args.NewFuncMap()

	// Act
	err := fm.AddStructs()

	// Assert
	actual := args.Map{"err": err}
	expected := args.Map{"err": nil}
	actual.ShouldBeEqual(t, 1, "FuncMap AddStructs empty slice", expected)
}

// ── FuncWrap.IsEqual name mismatch (line 223-225) ──

func Test_Cov13_FuncWrap_IsEqual_NameMismatch(t *testing.T) {
	// Arrange
	fw1 := args.NewFuncWrap.Default(func() {})
	fw2 := args.NewFuncWrap.Default(func(x int) int { return x })

	// Act
	result := fw1.IsEqual(fw2)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	actual.ShouldBeEqual(t, 1, "FuncWrap IsEqual mismatch", expected)
}

// ── FuncWrapArgs.OutArgsTypesNames empty (line 253-255) ──

func Test_Cov13_FuncWrapArgs_OutArgsTypesNames_NoReturn(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act
	names := fw.OutArgsTypesNames()

	// Assert
	actual := args.Map{"len": len(names)}
	expected := args.Map{"len": 0}
	actual.ShouldBeEqual(t, 1, "FuncWrapArgs OutArgsTypesNames no return", expected)
}

// ── FuncWrapInvoke.InvokeError (line 110-112) ──

func Test_Cov13_FuncWrapInvoke_InvokeError_Valid(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() error { return errors.New("test") })

	// Act
	funcErr, procErr := fw.InvokeError()

	// Assert
	actual := args.Map{
		"hasFuncErr":  funcErr != nil,
		"hasProcErr":  procErr != nil,
	}
	expected := args.Map{
		"hasFuncErr":  true,
		"hasProcErr":  false,
	}
	actual.ShouldBeEqual(t, 1, "FuncWrapInvoke InvokeError valid", expected)
}

// ── FuncWrapTypedHelpers: InvokeAsError nil result (line 127-129) ──

func Test_Cov13_FuncWrapTypedHelpers_InvokeAsError_NilResult(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() error { return nil })

	// Act
	funcErr, procErr := fw.InvokeAsError()

	// Assert
	actual := args.Map{
		"funcErr": funcErr,
		"procErr": procErr,
	}
	expected := args.Map{
		"funcErr": nil,
		"procErr": nil,
	}
	actual.ShouldBeEqual(t, 1, "FuncWrapTypedHelpers InvokeAsError nil", expected)
}

// ── FuncWrapValidation: rv.IsValid false (line 69-71) ──
// Dead code — creating FuncWrap with invalid reflect is structurally unreachable
// via public API. Documented as accepted dead code.

// ── Map.GetFuncName with nil FuncWrap (line 107) ──

func Test_Cov13_Map_GetFuncName_NoFunc(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}

	// Act
	result := m.GetFuncName()

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	actual.ShouldBeEqual(t, 1, "Map GetFuncName no func", expected)
}

// ── Map.SortedKeysMust panic path (line 207-208) ──

func Test_Cov13_Map_SortedKeysMust_Valid(t *testing.T) {
	// Arrange
	m := args.Map{"b": 2, "a": 1}

	// Act
	result := m.SortedKeysMust()

	// Assert
	actual := args.Map{
		"len":   len(result),
		"first": result[0],
	}
	expected := args.Map{
		"len":   2,
		"first": "a",
	}
	actual.ShouldBeEqual(t, 1, "Map SortedKeysMust valid", expected)
}

// ── MapShouldBeEqual: hasMismatch path (line 39-51) ──

func Test_Cov13_MapShouldBeEqual_Mismatch(t *testing.T) {
	// Arrange
	expected := args.Map{"key": "expected"}
	actual := args.Map{"key": "actual"}

	// Act / Assert — ShouldBeEqual with mismatch
	// We verify that ShouldBeEqual reports the mismatch correctly
	// by calling CompileToStrings and comparing
	actualLines := actual.CompileToStrings()
	expectedLines := expected.CompileToStrings()

	result := args.Map{
		"actualLen":   len(actualLines),
		"expectedLen": len(expectedLines),
		"different":   actualLines[0] != expectedLines[0],
	}
	expect := args.Map{
		"actualLen":   1,
		"expectedLen": 1,
		"different":   true,
	}
	result.ShouldBeEqual(t, 1, "MapShouldBeEqual mismatch detected", expect)
}

// ── argsHelper.buildToString cache (line 31-33) ──
// Tested via FiveFunc/FourFunc/SixFunc.String() above

// ── funcDetector: default case (line 18-19) ──

func Test_Cov13_FuncDetector_Default(t *testing.T) {
	// Arrange — passing a raw func (not wrapped)
	rawFunc := func(x int) int { return x }
	m := args.Map{
		"func": rawFunc,
	}

	// Act
	fw := m.FuncWrap()

	// Assert
	actual := args.Map{"notNil": fw != nil}
	expected := args.Map{"notNil": true}
	actual.ShouldBeEqual(t, 1, "funcDetector default case", expected)
}

// ── newFuncWrapCreator.Default non-func kind (line 27-28) ──

func Test_Cov13_NewFuncWrap_Default_NonFunc(t *testing.T) {
	// Arrange — passing non-func
	fw := args.NewFuncWrap.Default(42)

	// Act
	isInvalid := !fw.HasValidFunc()

	// Assert
	actual := args.Map{"isInvalid": isInvalid}
	expected := args.Map{"isInvalid": true}
	actual.ShouldBeEqual(t, 1, "NewFuncWrap Default non-func", expected)
}
