package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// FuncWrapAny — basic
// ═══════════════════════════════════════════

func Test_FuncWrapAny_NilFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(nil)

	// Act
	actual := args.Map{
		"isInvalid":    fw.IsInvalid(),
		"hasValidFunc": fw.HasValidFunc(),
	}

	// Assert
	expected := args.Map{
		"isInvalid": true,
		"hasValidFunc": false,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrapAny returns nil -- nil func", actual)
}

func Test_FuncWrapAny_ValidFunc(t *testing.T) {
	// Arrange
	fn := func(a, b int) int { return a + b }
	fw := args.NewFuncWrap.Default(fn)

	// Act
	actual := args.Map{
		"nameNE":       fw.Name != "",
		"isInvalid":    fw.IsInvalid(),
		"hasValidFunc": fw.HasValidFunc(),
	}

	// Assert
	expected := args.Map{
		"nameNE": true,
		"isInvalid": false,
		"hasValidFunc": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrapAny returns non-empty -- valid func", actual)
}

func Test_FuncWrapAny_Invoke(t *testing.T) {
	// Arrange
	fn := func(a, b int) int { return a + b }
	fw := args.NewFuncWrap.Default(fn)
	results, err := fw.Invoke(3, 4)

	// Act
	actual := args.Map{
		"errNil":    err == nil,
		"resultLen": len(results),
	}

	// Assert
	expected := args.Map{
		"errNil": true,
		"resultLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrapAny returns correct value -- invoke", actual)
}

func Test_FuncWrapAny_InvokeNilFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(nil)
	_, err := fw.Invoke(1)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FuncWrapAny returns nil -- invoke nil func", actual)
}

// ═══════════════════════════════════════════
// Map — accessors
// ═══════════════════════════════════════════

func Test_Map_GetAs(t *testing.T) {
	// Arrange
	m := args.Map{
		"str": "hello", "int": 42, "bool": true,
		"strs": []string{"a", "b"},
	}
	strVal, strOK := m.GetAsString("str")
	intVal, intOK := m.GetAsInt("int")
	boolVal, boolOK := m.GetAsBool("bool")
	strsVal, strsOK := m.GetAsStrings("strs")
	_, missingOK := m.GetAsString("missing")

	// Act
	actual := args.Map{
		"str":       strVal,
		"strOK":     strOK,
		"int":       intVal,
		"intOK":     intOK,
		"bool":      boolVal,
		"boolOK":    boolOK,
		"strsLen":   len(strsVal),
		"strsOK":    strsOK,
		"missingOK": missingOK,
	}

	// Assert
	expected := args.Map{
		"str": "hello", "strOK": true,
		"int": 42, "intOK": true,
		"bool": true, "boolOK": true,
		"strsLen": 2, "strsOK": true,
		"missingOK": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAs", actual)
}

func Test_Map_GetAsBoolDefault(t *testing.T) {
	// Arrange
	m := args.Map{"bool": true}
	val := m.GetAsBoolDefault("bool", false)
	missing := m.GetAsBoolDefault("missing", false)

	// Act
	actual := args.Map{
		"val": val,
		"missing": missing,
	}

	// Assert
	expected := args.Map{
		"val": true,
		"missing": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsBoolDefault", actual)
}

func Test_Map_ArgsCount(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
		"func": nil,
		"expect": nil,
	}

	// Act
	actual := args.Map{"count": m.ArgsCount()}

	// Assert
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- ArgsCount", actual)
}

func Test_Map_WorkFunc(t *testing.T) {
	// Arrange
	fn := func() string { return "hello" }
	m := args.Map{"func": fn}
	wf := m.WorkFunc()

	// Act
	actual := args.Map{"hasFunc": wf != nil}

	// Assert
	expected := args.Map{"hasFunc": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- WorkFunc", actual)
}

func Test_Map_GetFirstOfNames(t *testing.T) {
	// Arrange
	m := args.Map{
		"input": "hello",
		"when": "hello2",
	}
	first := m.GetFirstOfNames("input", "when")
	missing := m.GetFirstOfNames("missing1", "missing2")

	// Act
	actual := args.Map{
		"first":   first,
		"missing": missing == nil,
	}

	// Assert
	expected := args.Map{
		"first": "hello",
		"missing": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetFirstOfNames", actual)
}

func Test_Map_HasFunc(t *testing.T) {
	// Arrange
	fn := func() {}
	m1 := args.Map{"func": fn}
	m2 := args.Map{"other": 1}

	// Act
	actual := args.Map{
		"hasFunc":   m1.HasFunc(),
		"alsoHas":   m2.HasFunc(),
	}
	// HasFunc() always returns true because FuncWrap.Default(nil) returns non-nil *FuncWrapAny

	// Assert
	expected := args.Map{
		"hasFunc": true,
		"alsoHas": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- HasFunc", actual)
}

func Test_Map_CompileToStrings(t *testing.T) {
	// Arrange
	m := args.Map{
		"b": 2,
		"a": 1,
	}
	lines := m.CompileToStrings()

	// Act
	actual := args.Map{
		"linesLen": len(lines),
		"sorted":   lines[0] < lines[1],
	}

	// Assert
	expected := args.Map{
		"linesLen": 2,
		"sorted": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- CompileToStrings", actual)
}

func Test_Map_GoLiteralLines(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}
	lines := m.GoLiteralLines()

	// Act
	actual := args.Map{"linesLen": len(lines) > 0}

	// Assert
	expected := args.Map{"linesLen": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GoLiteralLines", actual)
}

// ═══════════════════════════════════════════
// One through Six — basic (using *Any aliases for untyped usage)
// ═══════════════════════════════════════════

func Test_One_Basic(t *testing.T) {
	// Arrange
	o := args.OneAny{First: 1}

	// Act
	actual := args.Map{
		"first":  o.First,
		"str":    o.String() != "",
		"count":  o.ArgsCount(),
		"slice":  len(o.Slice()),
	}

	// Assert
	expected := args.Map{
		"first": 1, "str": true, "count": 1, "slice": 1,
	}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- basic", actual)
}

func Test_Two_Basic(t *testing.T) {
	// Arrange
	tw := args.TwoAny{First: 1, Second: 2}

	// Act
	actual := args.Map{
		"first":  tw.First,
		"second": tw.Second,
		"count":  tw.ArgsCount(),
		"slice":  len(tw.Slice()),
	}

	// Assert
	expected := args.Map{
		"first": 1,
		"second": 2,
		"count": 2,
		"slice": 2,
	}
	expected.ShouldBeEqual(t, 0, "Two returns correct value -- basic", actual)
}

func Test_Three_Basic(t *testing.T) {
	// Arrange
	th := args.ThreeAny{First: 1, Second: 2, Third: 3}

	// Act
	actual := args.Map{
		"first":  th.First,
		"second": th.Second,
		"third":  th.Third,
		"count":  th.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first": 1,
		"second": 2,
		"third": 3,
		"count": 3,
	}
	expected.ShouldBeEqual(t, 0, "Three returns correct value -- basic", actual)
}

func Test_Four_Basic_FromFuncWrapAnyNilFunc(t *testing.T) {
	// Arrange
	f := args.FourAny{First: 1, Second: 2, Third: 3, Fourth: 4}

	// Act
	actual := args.Map{
		"first":  f.First,
		"second": f.Second,
		"third":  f.Third,
		"fourth": f.Fourth,
		"count":  f.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first": 1,
		"second": 2,
		"third": 3,
		"fourth": 4,
		"count": 4,
	}
	expected.ShouldBeEqual(t, 0, "Four returns correct value -- basic", actual)
}

func Test_Five_Basic_FromFuncWrapAnyNilFunc(t *testing.T) {
	// Arrange
	f := args.FiveAny{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5}

	// Act
	actual := args.Map{
		"first":  f.First,
		"second": f.Second,
		"third":  f.Third,
		"fourth": f.Fourth,
		"fifth":  f.Fifth,
		"count":  f.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first": 1, "second": 2, "third": 3, "fourth": 4, "fifth": 5, "count": 5,
	}
	expected.ShouldBeEqual(t, 0, "Five returns correct value -- basic", actual)
}

func Test_Six_Basic_FromFuncWrapAnyNilFunc(t *testing.T) {
	// Arrange
	s := args.SixAny{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, Sixth: 6}

	// Act
	actual := args.Map{
		"first":  s.First,
		"second": s.Second,
		"third":  s.Third,
		"fourth": s.Fourth,
		"fifth":  s.Fifth,
		"sixth":  s.Sixth,
		"count":  s.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first": 1, "second": 2, "third": 3, "fourth": 4,
		"fifth": 5, "sixth": 6, "count": 6,
	}
	expected.ShouldBeEqual(t, 0, "Six returns correct value -- basic", actual)
}

// ═══════════════════════════════════════════
// Holder
// ═══════════════════════════════════════════

func Test_Holder_Basic(t *testing.T) {
	// Arrange
	h := args.HolderAny{First: "hello"}

	// Act
	actual := args.Map{
		"first": h.First,
		"count": h.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first": "hello",
		"count": 7,
	}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- basic", actual)
}

// ═══════════════════════════════════════════
// LeftRight — args
// ═══════════════════════════════════════════

func Test_LeftRight_Basic(t *testing.T) {
	// Arrange
	lr := args.LeftRightAny{Left: "hello", Right: 42}

	// Act
	actual := args.Map{
		"left":  lr.Left,
		"right": lr.Right,
		"count": lr.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"left": "hello", "right": 42, "count": 2,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- basic", actual)
}

// ═══════════════════════════════════════════
// Dynamic
// ═══════════════════════════════════════════

func Test_Dynamic_Basic_FromFuncWrapAnyNilFunc(t *testing.T) {
	// Arrange
	d := args.DynamicAny{Params: args.Map{"val": "hello"}, Expect: "expected"}

	// Act
	actual := args.Map{
		"expect": d.Expect,
		"hasVal": d.HasDefined("val"),
	}

	// Assert
	expected := args.Map{
		"expect": "expected",
		"hasVal": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- basic", actual)
}
