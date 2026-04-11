package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Map basic methods ──

func Test_Map_Length_FromMapLength(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}

	// Act
	actual := args.Map{"length": m.Length()}

	// Assert
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "Map.Length returns 2 -- two entries", actual)
}

func Test_Map_ArgsCount_FromMapLength(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
		"expected": 3,
	}

	// Act
	actual := args.Map{"count": m.ArgsCount()}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Map.ArgsCount excludes expected and func -- 3 entries minus 2", actual)
}

func Test_Map_Expected_FromMapLength(t *testing.T) {
	// Arrange
	m := args.Map{"expected": 42}

	// Act
	actual := args.Map{"val": m.Expected()}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "Map.Expected returns value -- key 'expected'", actual)
}

func Test_Map_HasExpect_FromMapLength(t *testing.T) {
	// Arrange
	m1 := args.Map{"expected": 42}
	m2 := args.Map{"a": 1}

	// Act
	actual := args.Map{
		"has": m1.HasExpect(),
		"notHas": m2.HasExpect(),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"notHas": false,
	}
	expected.ShouldBeEqual(t, 0, "Map.HasExpect returns correct -- with and without", actual)
}

func Test_Map_HasFunc_FromMapLength(t *testing.T) {
	// Arrange
	m1 := args.Map{"func": func() {}}
	m2 := args.Map{"a": 1}

	// Act
	actual := args.Map{
		"has": m1.HasFunc(),
		"notHas": m2.HasFunc(),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"notHas": true,
	}
	expected.ShouldBeEqual(t, 0, "Map.HasFunc returns true -- both defined", actual)
}

func Test_Map_GetAs_FromMapLength(t *testing.T) {
	// Arrange
	m := args.Map{
		"name": "hello",
		"count": 42,
		"flag": true,
	}
	name, _ := m.GetAsString("name")
	count, _ := m.GetAsInt("count")
	flag, _ := m.GetAsBool("flag")

	// Act
	actual := args.Map{
		"name": name,
		"count": count,
		"flag": flag,
	}

	// Assert
	expected := args.Map{
		"name": "hello",
		"count": 42,
		"flag": true,
	}
	expected.ShouldBeEqual(t, 0, "Map.GetAs* returns correct types -- string, int, bool", actual)
}

func Test_Map_GetAsStringSlice(t *testing.T) {
	// Arrange
	m := args.Map{"items": []string{"a", "b"}}
	items, _ := m.GetAsStrings("items")

	// Act
	actual := args.Map{"len": len(items)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Map.GetAsStringSlice returns correct -- 2 items", actual)
}

func Test_Map_GetAsBytes(t *testing.T) {
	// Arrange
	m := args.Map{"data": []byte{1, 2, 3}}
	raw, _ := m.Get("data")
	data, _ := raw.([]byte)

	// Act
	actual := args.Map{"len": len(data)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Map.GetAsBytes returns correct -- 3 bytes", actual)
}

func Test_Map_WorkFunc_FromMapLength(t *testing.T) {
	// Arrange
	fn := func() string { return "hello" }
	m := args.Map{"func": fn}

	// Act
	actual := args.Map{"notNil": m.WorkFunc() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Map.WorkFunc returns non-nil -- has func", actual)
}

func Test_Map_GetFirstOfNames_FromMapLength(t *testing.T) {
	// Arrange
	m := args.Map{"name": "hello"}
	val := m.GetFirstOfNames("missing", "name", "other")

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Map.GetFirstOfNames finds first match -- name key", actual)
}

func Test_Map_SortedKeys_FromMapLength(t *testing.T) {
	// Arrange
	m := args.Map{
		"c": 3,
		"a": 1,
		"b": 2,
	}
	keys, _ := m.SortedKeys()

	// Act
	actual := args.Map{
		"first": keys[0],
		"last": keys[2],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "Map.SortedKeys returns sorted -- 3 keys", actual)
}

func Test_Map_String_FromMapLength(t *testing.T) {
	// Arrange
	m := args.Map{"key": "value"}
	s := m.String()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Map.String returns non-empty -- single entry", actual)
}

// ── One[T1] basic methods ──

func Test_One_Basic_FromMapLength(t *testing.T) {
	// Arrange
	one := &args.One[string]{First: "hello", Expect: 42}

	// Act
	actual := args.Map{
		"first":     one.FirstItem(),
		"expected":  one.Expected(),
		"hasFirst":  one.HasFirst(),
		"hasExpect": one.HasExpect(),
		"count":     one.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first":     "hello",
		"expected":  42,
		"hasFirst":  true,
		"hasExpect": true,
		"count":     1,
	}
	expected.ShouldBeEqual(t, 0, "One basic getters -- string first", actual)
}

func Test_One_ArgTwo_FromMapLength(t *testing.T) {
	// Arrange
	one := &args.One[string]{First: "hello", Expect: 42}
	two := one.ArgTwo()

	// Act
	actual := args.Map{
		"first": two.First,
		"expect": two.Expect,
	}

	// Assert
	expected := args.Map{
		"first": "hello",
		"expect": 42,
	}
	expected.ShouldBeEqual(t, 0, "One.ArgTwo returns copy -- same data", actual)
}

func Test_One_Args_FromMapLength(t *testing.T) {
	// Arrange
	one := &args.One[string]{First: "hello"}
	a := one.Args(1)

	// Act
	actual := args.Map{"len": len(a)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "One.Args returns 1 -- single first", actual)
}

func Test_One_Slice(t *testing.T) {
	// Arrange
	one := &args.One[string]{First: "hello"}
	s := one.Slice()

	// Act
	actual := args.Map{"len": len(s)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "One.Slice returns 1 -- single first", actual)
}

func Test_One_String_FromMapLength(t *testing.T) {
	// Arrange
	one := &args.One[string]{First: "hello"}
	s := one.String()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "One.String returns non-empty -- has first", actual)
}

func Test_One_GetByIndex_FromMapLength(t *testing.T) {
	// Arrange
	one := &args.One[string]{First: "hello"}
	val := one.GetByIndex(0)

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "One.GetByIndex returns first -- index 0", actual)
}

func Test_One_LeftRight_FromMapLength(t *testing.T) {
	// Arrange
	one := &args.One[string]{First: "hello", Expect: "world"}
	lr := one.LeftRight()

	// Act
	actual := args.Map{
		"left": lr.Left,
		"expect": lr.Expect,
	}

	// Assert
	expected := args.Map{
		"left": "hello",
		"expect": "world",
	}
	expected.ShouldBeEqual(t, 0, "One.LeftRight returns left=first expect=expect -- set", actual)
}

func Test_One_AsOneParameter(t *testing.T) {
	// Arrange
	one := &args.One[string]{First: "hello"}
	param := one.AsOneParameter()

	// Act
	actual := args.Map{"notNil": param != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "One.AsOneParameter returns non-nil -- valid", actual)
}

// ── Two[T1, T2] basic methods ──

func Test_Two_Basic_FromMapLength(t *testing.T) {
	// Arrange
	two := &args.Two[string, int]{First: "hello", Second: 42, Expect: true}

	// Act
	actual := args.Map{
		"first":     two.FirstItem(),
		"second":    two.SecondItem(),
		"expected":  two.Expected(),
		"hasFirst":  two.HasFirst(),
		"hasSecond": two.HasSecond(),
		"hasExpect": two.HasExpect(),
		"count":     two.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first":     "hello",
		"second":    42,
		"expected":  true,
		"hasFirst":  true,
		"hasSecond": true,
		"hasExpect": true,
		"count":     2,
	}
	expected.ShouldBeEqual(t, 0, "Two basic getters -- string and int", actual)
}

func Test_Two_Args(t *testing.T) {
	// Arrange
	two := &args.Two[string, int]{First: "hello", Second: 42}
	a := two.Args(2)

	// Act
	actual := args.Map{"len": len(a)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Two.Args returns 2 -- first and second", actual)
}

func Test_Two_Slice(t *testing.T) {
	// Arrange
	two := &args.Two[string, int]{First: "hello", Second: 42}
	s := two.Slice()

	// Act
	actual := args.Map{"len": len(s)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Two.Slice returns 2 -- first and second", actual)
}

func Test_Two_GetByIndex_FromMapLength(t *testing.T) {
	// Arrange
	two := &args.Two[string, int]{First: "hello", Second: 42}

	// Act
	actual := args.Map{
		"idx0": two.GetByIndex(0),
		"idx1": two.GetByIndex(1),
	}

	// Assert
	expected := args.Map{
		"idx0": "hello",
		"idx1": 42,
	}
	expected.ShouldBeEqual(t, 0, "Two.GetByIndex returns correct -- index 0 and 1", actual)
}

func Test_Two_String_FromMapLength(t *testing.T) {
	// Arrange
	two := &args.Two[string, int]{First: "hello", Second: 42}
	s := two.String()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Two.String returns non-empty -- has items", actual)
}

func Test_Two_LeftRight_FromMapLength(t *testing.T) {
	// Arrange
	two := &args.Two[string, int]{First: "hello", Second: 42}
	lr := two.LeftRight()

	// Act
	actual := args.Map{
		"left": lr.Left,
		"right": lr.Right,
	}

	// Assert
	expected := args.Map{
		"left": "hello",
		"right": 42,
	}
	expected.ShouldBeEqual(t, 0, "Two.LeftRight returns first and second -- set", actual)
}

// ── Three[T1, T2, T3] basic methods ──

func Test_Three_Basic_FromMapLength(t *testing.T) {
	// Arrange
	three := &args.Three[string, int, bool]{First: "hello", Second: 42, Third: true, Expect: "yes"}

	// Act
	actual := args.Map{
		"first":     three.FirstItem(),
		"second":    three.SecondItem(),
		"third":     three.ThirdItem(),
		"expected":  three.Expected(),
		"hasFirst":  three.HasFirst(),
		"hasSecond": three.HasSecond(),
		"hasThird":  three.HasThird(),
		"count":     three.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first":     "hello",
		"second":    42,
		"third":     true,
		"expected":  "yes",
		"hasFirst":  true,
		"hasSecond": true,
		"hasThird":  true,
		"count":     3,
	}
	expected.ShouldBeEqual(t, 0, "Three basic getters -- string, int, bool", actual)
}

func Test_Three_Args(t *testing.T) {
	// Arrange
	three := &args.Three[string, int, bool]{First: "a", Second: 1, Third: true}
	a := three.Args(3)

	// Act
	actual := args.Map{"len": len(a)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Three.Args returns 3 -- all three", actual)
}

func Test_Three_GetByIndex_FromMapLength(t *testing.T) {
	// Arrange
	three := &args.Three[string, int, bool]{First: "a", Second: 1, Third: true}

	// Act
	actual := args.Map{
		"idx0": three.GetByIndex(0),
		"idx1": three.GetByIndex(1),
		"idx2": three.GetByIndex(2),
	}

	// Assert
	expected := args.Map{
		"idx0": "a",
		"idx1": 1,
		"idx2": true,
	}
	expected.ShouldBeEqual(t, 0, "Three.GetByIndex returns correct -- all indexes", actual)
}

func Test_Three_String(t *testing.T) {
	// Arrange
	three := &args.Three[string, int, bool]{First: "a", Second: 1, Third: true}
	s := three.String()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Three.String returns non-empty -- has items", actual)
}

// ── FuncWrap (via args.NewFuncWrap) ──

func Test_FuncWrap_Basic(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Default(fn)

	// Act
	actual := args.Map{
		"isValid": fw.IsValid(),
		"name":    len(fw.Name) > 0,
	}

	// Assert
	expected := args.Map{
		"isValid": true,
		"name":    true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap basic -- valid func", actual)
}

func Test_FuncWrap_InOutArgs(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Default(fn)

	// Act
	actual := args.Map{
		"inCount":  fw.InArgsCount(),
		"outCount": fw.OutArgsCount(),
	}

	// Assert
	expected := args.Map{
		"inCount":  1,
		"outCount": 1,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap in/out args -- func(string)int", actual)
}

func Test_FuncWrap_IsStringFunc_FromMapLength(t *testing.T) {
	// Arrange
	fn := func() string { return "hello" }
	fw := args.NewFuncWrap.Default(fn)

	// Act
	actual := args.Map{"isString": fw.IsStringFunc()}

	// Assert
	expected := args.Map{"isString": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap.IsStringFunc returns true -- func()string", actual)
}

func Test_FuncWrap_IsBoolFunc_FromMapLength(t *testing.T) {
	// Arrange
	fn := func() bool { return true }
	fw := args.NewFuncWrap.Default(fn)

	// Act
	actual := args.Map{"isBool": fw.IsBoolFunc()}

	// Assert
	expected := args.Map{"isBool": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap.IsBoolFunc returns true -- func()bool", actual)
}

func Test_FuncWrap_IsVoidFunc_FromMapLength(t *testing.T) {
	// Arrange
	fn := func() {}
	fw := args.NewFuncWrap.Default(fn)

	// Act
	actual := args.Map{"isVoid": fw.IsVoidFunc()}

	// Assert
	expected := args.Map{"isVoid": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap.IsVoidFunc returns true -- func()", actual)
}

func Test_FuncWrap_IsErrorFunc_FromMapLength(t *testing.T) {
	// Arrange
	fn := func() error { return nil }
	fw := args.NewFuncWrap.Default(fn)

	// Act
	actual := args.Map{"isError": fw.IsErrorFunc()}

	// Assert
	expected := args.Map{"isError": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap.IsErrorFunc returns true -- func()error", actual)
}

func Test_FuncWrap_Name(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }
	fw := args.NewFuncWrap.Default(fn)

	// Act
	actual := args.Map{"hasContent": len(fw.Name) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap.Name has content -- valid func", actual)
}

// ── Holder ──

func Test_Holder_Basic_FromMapLength(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }
	h := &args.Holder[func(string) int]{
		First:    "hello",
		Expect:   5,
		WorkFunc: fn,
	}

	// Act
	actual := args.Map{
		"first":     h.FirstItem(),
		"expected":  h.Expected(),
		"hasFirst":  h.HasFirst(),
		"hasExpect": h.HasExpect(),
		"hasFunc":   h.HasFunc(),
	}

	// Assert
	expected := args.Map{
		"first":     "hello",
		"expected":  5,
		"hasFirst":  true,
		"hasExpect": true,
		"hasFunc":   true,
	}
	expected.ShouldBeEqual(t, 0, "Holder basic getters -- string first with func", actual)
}

func Test_Holder_GetFuncName(t *testing.T) {
	// Arrange
	fn := func() string { return "hello" }
	h := &args.Holder[func() string]{WorkFunc: fn}
	name := h.GetFuncName()

	// Act
	actual := args.Map{"hasContent": len(name) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Holder.GetFuncName returns non-empty -- has func", actual)
}

func Test_Holder_FuncWrap(t *testing.T) {
	// Arrange
	fn := func() string { return "hello" }
	h := &args.Holder[func() string]{WorkFunc: fn}
	fw := h.FuncWrap()

	// Act
	actual := args.Map{"isValid": fw.IsValid()}

	// Assert
	expected := args.Map{"isValid": true}
	expected.ShouldBeEqual(t, 0, "Holder.FuncWrap returns valid -- has func", actual)
}

// ── LeftRight ──

func Test_LeftRight_Basic_FromMapLength(t *testing.T) {
	// Arrange
	lr := &args.LeftRight[string, string]{Left: "hello", Right: "world"}

	// Act
	actual := args.Map{
		"left":  lr.Left,
		"right": lr.Right,
	}

	// Assert
	expected := args.Map{
		"left":  "hello",
		"right": "world",
	}
	expected.ShouldBeEqual(t, 0, "LeftRight basic -- both set", actual)
}

func Test_LeftRight_GetByIndex(t *testing.T) {
	// Arrange
	lr := &args.LeftRight[string, string]{Left: "L", Right: "R"}

	// Act
	actual := args.Map{
		"idx0": lr.GetByIndex(0),
		"idx1": lr.GetByIndex(1),
	}

	// Assert
	expected := args.Map{
		"idx0": "L",
		"idx1": "R",
	}
	expected.ShouldBeEqual(t, 0, "LeftRight.GetByIndex returns correct -- both indexes", actual)
}

// ── String type ──

func Test_String_Basic(t *testing.T) {
	// Arrange
	sa := args.String("hello world")

	// Act
	actual := args.Map{
		"str":     sa.String(),
		"len":     sa.Length(),
		"isEmpty": sa.IsEmpty(),
		"hasChr":  sa.HasCharacter(),
	}

	// Assert
	expected := args.Map{
		"str":     "hello world",
		"len":     11,
		"isEmpty": false,
		"hasChr":  true,
	}
	expected.ShouldBeEqual(t, 0, "String type basic -- hello world", actual)
}
