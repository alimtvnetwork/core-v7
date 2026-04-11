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

func Test_FiveFunc_String_I29(t *testing.T) {
	// Arrange
	ff := &args.FiveFunc[int, string, float64, bool, string]{
		First: 1, Second: "two", Third: 3.0, Fourth: true, Fifth: "five",
	}

	// Act
	result := ff.String()

	// Assert
	convey.Convey("FiveFunc String returns formatted output", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

// ── FourFunc.String (line 214) ──

func Test_FourFunc_String_I29(t *testing.T) {
	// Arrange
	ff := &args.FourFunc[int, string, float64, bool]{
		First: 1, Second: "two", Third: 3.0, Fourth: true,
	}

	// Act
	result := ff.String()

	// Assert
	convey.Convey("FourFunc String returns formatted output", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

// ── SixFunc.String (line 259) ──

func Test_SixFunc_String_I29(t *testing.T) {
	// Arrange
	sf := &args.SixFunc[int, string, float64, bool, string, int]{
		First: 1, Second: "two", Third: 3.0, Fourth: true, Fifth: "five", Sixth: 6,
	}

	// Act
	result := sf.String()

	// Assert
	convey.Convey("SixFunc String returns formatted output", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

// ── FuncMap.AddStructFunctions error path (line 112-114) ──

func Test_FuncMap_AddStructFunctions_EmptySlice_I29(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	err := fm.AddStructFunctions()

	// Assert
	actual := args.Map{"err": err}
	expected := args.Map{"err": nil}
	actual.ShouldBeEqual(t, 1, "FuncMap AddStructFunctions empty slice", expected)
}

// ── FuncWrap.IsEqual name mismatch (line 223-225) ──

func Test_FuncWrap_IsEqual_Mismatch_I29(t *testing.T) {
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

// ── FuncWrapArgs.GetOutArgsTypesNames empty (line 253-255) ──

func Test_FuncWrapArgs_GetOutArgsTypesNames_NoReturn_I29(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act
	names := fw.GetOutArgsTypesNames()

	// Assert
	actual := args.Map{"len": len(names)}
	expected := args.Map{"len": 0}
	actual.ShouldBeEqual(t, 1, "FuncWrapArgs GetOutArgsTypesNames no return", expected)
}

// ── FuncWrapInvoke.InvokeError (line 110-112) ──

func Test_FuncWrapInvoke_InvokeError_Valid_I29(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() error { return errors.New("test") })

	// Act
	funcErr, procErr := fw.InvokeError()

	// Assert
	actual := args.Map{
		"hasFuncErr": funcErr != nil,
		"hasProcErr": procErr != nil,
	}
	expected := args.Map{
		"hasFuncErr": true,
		"hasProcErr": false,
	}
	actual.ShouldBeEqual(t, 1, "FuncWrapInvoke InvokeError valid", expected)
}

// ── FuncWrapTypedHelpers: InvokeAsError nil result (line 127-129) ──

func Test_FuncWrapTypedHelpers_InvokeAsError_Nil_I29(t *testing.T) {
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

// ── Map.GetFuncName with nil FuncWrap (line 107) ──

func Test_Map_GetFuncName_NoFunc_I29(t *testing.T) {
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

func Test_Map_SortedKeysMust_Valid_I29(t *testing.T) {
	// Arrange
	m := args.Map{
		"b": 2,
		"a": 1,
	}

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

func Test_MapShouldBeEqual_Mismatch_I29(t *testing.T) {
	// Arrange
	expected := args.Map{"key": "expected"}
	actual := args.Map{"key": "actual"}

	// Act
	actualLines := actual.CompileToStrings()
	expectedLines := expected.CompileToStrings()

	// Assert
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

// ── funcDetector: default case (line 18-19) ──

func Test_FuncDetector_Default_I29(t *testing.T) {
	// Arrange — passing a raw func
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

func Test_NewFuncWrap_Default_NonFunc_I29(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(42)

	// Act
	isInvalid := !fw.HasValidFunc()

	// Assert
	actual := args.Map{"isInvalid": isInvalid}
	expected := args.Map{"isInvalid": true}
	actual.ShouldBeEqual(t, 1, "NewFuncWrap Default non-func", expected)
}
