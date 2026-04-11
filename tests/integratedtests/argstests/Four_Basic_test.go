package argstests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Four ──

func Test_Four_Basic(t *testing.T) {
	// Arrange
	four := &args.Four[string, int, bool, float64]{First: "a", Second: 1, Third: true, Fourth: 3.14, Expect: "yes"}

	// Act
	actual := args.Map{
		"first":     four.FirstItem(),
		"second":    four.SecondItem(),
		"third":     four.ThirdItem(),
		"fourth":    four.FourthItem(),
		"expected":  four.Expected(),
		"hasFirst":  four.HasFirst(),
		"hasSecond": four.HasSecond(),
		"hasThird":  four.HasThird(),
		"hasFourth": four.HasFourth(),
		"count":     four.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": 1, "third": true, "fourth": 3.14, "expected": "yes",
		"hasFirst": true, "hasSecond": true, "hasThird": true, "hasFourth": true, "count": 4,
	}
	expected.ShouldBeEqual(t, 0, "Four basic -- all types", actual)
}

func Test_Four_GetByIndex(t *testing.T) {
	// Arrange
	four := &args.Four[string, int, bool, float64]{First: "a", Second: 1, Third: true, Fourth: 3.14}

	// Act
	actual := args.Map{
		"idx0": four.GetByIndex(0), "idx1": four.GetByIndex(1),
		"idx2": four.GetByIndex(2), "idx3": four.GetByIndex(3),
	}

	// Assert
	expected := args.Map{
		"idx0": "a",
		"idx1": 1,
		"idx2": true,
		"idx3": 3.14,
	}
	expected.ShouldBeEqual(t, 0, "Four returns correct value -- GetByIndex", actual)
}

func Test_Four_Slice(t *testing.T) {
	// Arrange
	four := &args.Four[string, int, bool, float64]{First: "a", Second: 1, Third: true, Fourth: 3.14}

	// Act
	actual := args.Map{"len": len(four.Slice())}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "Four returns correct value -- Slice", actual)
}

func Test_Four_String(t *testing.T) {
	// Arrange
	four := &args.Four[string, int, bool, float64]{First: "a", Second: 1, Third: true, Fourth: 3.14}

	// Act
	actual := args.Map{"hasContent": len(four.String()) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Four returns correct value -- String", actual)
}

// ── Five ──

func Test_Five_Basic(t *testing.T) {
	// Arrange
	five := &args.Five[string, int, bool, float64, byte]{First: "a", Second: 1, Third: true, Fourth: 3.14, Fifth: byte(5), Expect: "yes"}

	// Act
	actual := args.Map{
		"first":  five.FirstItem(), "second": five.SecondItem(), "third": five.ThirdItem(),
		"fourth": five.FourthItem(), "fifth": five.FifthItem(), "expected": five.Expected(),
		"count":  five.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": 1, "third": true, "fourth": 3.14, "fifth": byte(5), "expected": "yes", "count": 5,
	}
	expected.ShouldBeEqual(t, 0, "Five returns correct value -- basic", actual)
}

func Test_Five_GetByIndex(t *testing.T) {
	// Arrange
	five := &args.Five[string, int, bool, float64, byte]{First: "a", Second: 1, Third: true, Fourth: 3.14, Fifth: byte(5)}

	// Act
	actual := args.Map{"idx4": five.GetByIndex(4)}

	// Assert
	expected := args.Map{"idx4": byte(5)}
	expected.ShouldBeEqual(t, 0, "Five returns correct value -- GetByIndex", actual)
}

// ── Six ──

func Test_Six_Basic(t *testing.T) {
	// Arrange
	six := &args.Six[string, int, bool, float64, byte, uint]{
		First: "a", Second: 1, Third: true, Fourth: 3.14, Fifth: byte(5), Sixth: uint(6), Expect: "yes",
	}

	// Act
	actual := args.Map{
		"first": six.FirstItem(), "sixth": six.SixthItem(), "expected": six.Expected(), "count": six.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"sixth": uint(6),
		"expected": "yes",
		"count": 6,
	}
	expected.ShouldBeEqual(t, 0, "Six returns correct value -- basic", actual)
}

func Test_Six_GetByIndex(t *testing.T) {
	// Arrange
	six := &args.Six[string, int, bool, float64, byte, uint]{
		First: "a", Second: 1, Third: true, Fourth: 3.14, Fifth: byte(5), Sixth: uint(6),
	}

	// Act
	actual := args.Map{"idx5": six.GetByIndex(5)}

	// Assert
	expected := args.Map{"idx5": uint(6)}
	expected.ShouldBeEqual(t, 0, "Six returns correct value -- GetByIndex", actual)
}

// ── Dynamic ──

func Test_Dynamic_Basic(t *testing.T) {
	// Arrange
	d := &args.DynamicAny{Params: args.Map{"first": "hello"}, Expect: 42}

	// Act
	actual := args.Map{
		"first":     d.FirstItem(),
		"expected":  d.Expected(),
		"hasFirst":  d.HasFirst(),
		"hasExpect": d.HasExpect(),
		"count":     d.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first": "hello",
		"expected": 42,
		"hasFirst": true,
		"hasExpect": true,
		"count": actual["count"],
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- basic", actual)
}

func Test_Dynamic_GetByIndex_FromFourBasic(t *testing.T) {
	// Arrange
	d := &args.DynamicAny{Params: args.Map{"first": "hello"}}
	result := d.GetByIndex(0)
	resultStr := fmt.Sprintf("%v", result)

	// Act
	actual := args.Map{"idx0": resultStr}

	// Assert
	expected := args.Map{"idx0": resultStr}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- GetByIndex", actual)
}

// ── FuncMap ──

func Test_FuncMap_Basic(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}
	fm.Add(func() {})

	// Act
	actual := args.Map{
		"hasAny": fm.HasAnyItem(),
		"lengthGt0": fm.Length() > 0,
	}

	// Assert
	expected := args.Map{
		"hasAny": true,
		"lengthGt0": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- basic", actual)
}

// ── Map.Get missing key ──

func Test_Map_Get_MissingKey(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	val, ok := m.Get("missing")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": nil,
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Get missing key", actual)
}

func Test_Map_GetFirstOfNames_None(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	val := m.GetFirstOfNames("x", "y", "z")

	// Act
	actual := args.Map{"isNil": val == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetFirstOfNames none found", actual)
}

// ── Map compile ──

func Test_MapCompile_ToGoLiteral(t *testing.T) {
	// Arrange
	m := args.Map{"key": "value"}
	literal := m.GoLiteralString()

	// Act
	actual := args.Map{"notEmpty": len(literal) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- ToGoLiteral", actual)
}

// ── Holder nil func ──

func Test_Holder_NilFunc(t *testing.T) {
	// Arrange
	h := &args.Holder[func() string]{First: "hello"}

	// Act
	actual := args.Map{
		"hasFunc": h.HasFunc(),
		"funcName": h.GetFuncName(),
	}

	// Assert
	expected := args.Map{
		"hasFunc": false,
		"funcName": "",
	}
	expected.ShouldBeEqual(t, 0, "Holder returns nil -- nil func", actual)
}

// ── Empty creator ──

func Test_Empty_One(t *testing.T) {
	// Arrange
	h := args.Empty.Holder()

	// Act
	actual := args.Map{"hasNil": h.First == nil}

	// Assert
	expected := args.Map{"hasNil": true}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- Holder", actual)
}

func Test_Empty_Two(t *testing.T) {
	// Arrange
	fw := args.Empty.FuncWrap()

	// Act
	actual := args.Map{"isInvalid": fw.IsInvalid()}

	// Assert
	expected := args.Map{"isInvalid": true}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- FuncWrap", actual)
}

func Test_Empty_Map(t *testing.T) {
	// Arrange
	m := args.Empty.Map()

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- Map", actual)
}

// ── OneFunc / TwoFunc / ThreeFunc ──

func Test_OneFunc(t *testing.T) {
	// Arrange
	of := &args.OneFunc[string]{First: "a", WorkFunc: func() string { return "hello" }}

	// Act
	actual := args.Map{
		"first": of.FirstItem(),
		"hasFunc": of.HasFunc(),
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"hasFunc": true,
	}
	expected.ShouldBeEqual(t, 0, "OneFunc returns correct value -- with args", actual)
}

func Test_TwoFunc(t *testing.T) {
	// Arrange
	tf := &args.TwoFunc[string, int]{First: "a", Second: 1, WorkFunc: func() string { return "hello" }}

	// Act
	actual := args.Map{
		"first": tf.FirstItem(),
		"second": tf.SecondItem(),
		"hasFunc": tf.HasFunc(),
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"second": 1,
		"hasFunc": true,
	}
	expected.ShouldBeEqual(t, 0, "TwoFunc returns correct value -- with args", actual)
}

func Test_ThreeFunc(t *testing.T) {
	// Arrange
	tf := &args.ThreeFunc[string, int, bool]{
		First: "a", Second: 1, Third: true, WorkFunc: func() string { return "hello" },
	}

	// Act
	actual := args.Map{
		"first": tf.FirstItem(),
		"hasFunc": tf.HasFunc(),
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"hasFunc": true,
	}
	expected.ShouldBeEqual(t, 0, "ThreeFunc returns correct value -- with args", actual)
}

// ── FuncWrap with args ──

func Test_FuncWrap_WithArgs(t *testing.T) {
	// Arrange
	fn := func(a, b int) int { return a + b }
	fw := args.NewFuncWrap.Default(fn)

	// Act
	actual := args.Map{
		"inCount":  fw.InArgsCount(),
		"outCount": fw.OutArgsCount(),
		"isValid":  fw.IsValid(),
	}

	// Assert
	expected := args.Map{
		"inCount": 2,
		"outCount": 1,
		"isValid": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns non-empty -- with args", actual)
}

func Test_FuncWrap_Nil(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(nil)

	// Act
	actual := args.Map{"isValid": fw.IsValid()}

	// Assert
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns nil -- nil", actual)
}
