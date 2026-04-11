package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ── String type ──

func Test_String_Concat_FromStringConcat(t *testing.T) {
	// Arrange
	s := args.String("hello")
	result := s.Concat(" ", "world")

	// Act
	actual := args.Map{
		"result": result.String(),
	}

	// Assert
	expected := args.Map{
		"result": "hello world",
	}
	expected.ShouldBeEqual(t, 0, "String_Concat returns correct value -- with args", actual)
}

func Test_String_Join_FromStringConcat(t *testing.T) {
	// Arrange
	s := args.String("hello")
	result := s.Join("-", "world", "go")

	// Act
	actual := args.Map{
		"notEmpty": result.String() != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "String_Join returns correct value -- with args", actual)
}

func Test_String_Split_FromStringConcat(t *testing.T) {
	// Arrange
	s := args.String("a,b,c")
	result := s.Split(",")

	// Act
	actual := args.Map{
		"len": len(result),
	}

	// Assert
	expected := args.Map{
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "String_Split returns correct value -- with args", actual)
}

func Test_String_Quoting_FromStringConcat(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act
	actual := args.Map{
		"doubleQuote":      s.DoubleQuote().String() != "",
		"doubleQuoteQ":     s.DoubleQuoteQ().String() != "",
		"singleQuote":      s.SingleQuote().String() != "",
		"valueDoubleQuote": s.ValueDoubleQuote().String() != "",
	}

	// Assert
	expected := args.Map{
		"doubleQuote":      true,
		"doubleQuoteQ":     true,
		"singleQuote":      true,
		"valueDoubleQuote": true,
	}
	expected.ShouldBeEqual(t, 0, "String_Quoting returns correct value -- with args", actual)
}

func Test_String_Length(t *testing.T) {
	// Arrange
	s := args.String("hello")
	empty := args.String("")

	// Act
	actual := args.Map{
		"length":      s.Length(),
		"count":       s.Count(),
		"asciiLen":    s.AscIILength(),
		"isEmpty":     s.IsEmpty(),
		"emptyIsTrue": empty.IsEmpty(),
		"hasCh":       s.HasCharacter(),
		"isDefined":   s.IsDefined(),
	}

	// Assert
	expected := args.Map{
		"length":      5,
		"count":       5,
		"asciiLen":    5,
		"isEmpty":     false,
		"emptyIsTrue": true,
		"hasCh":       true,
		"isDefined":   true,
	}
	expected.ShouldBeEqual(t, 0, "String_Length returns correct value -- with args", actual)
}

func Test_String_IsEmptyOrWhitespace(t *testing.T) {
	// Arrange
	s := args.String("  ")
	sNonEmpty := args.String("hello")

	// Act
	actual := args.Map{
		"whitespace": s.IsEmptyOrWhitespace(),
		"nonEmpty":   sNonEmpty.IsEmptyOrWhitespace(),
	}

	// Assert
	expected := args.Map{
		"whitespace": true,
		"nonEmpty":   false,
	}
	expected.ShouldBeEqual(t, 0, "String_IsEmptyOrWhitespace returns empty -- with args", actual)
}

func Test_String_TrimSpace_FromStringConcat(t *testing.T) {
	// Arrange
	s := args.String("  hello  ")
	result := s.TrimSpace()

	// Act
	actual := args.Map{
		"result": result.String(),
	}

	// Assert
	expected := args.Map{
		"result": "hello",
	}
	expected.ShouldBeEqual(t, 0, "String_TrimSpace returns correct value -- with args", actual)
}

func Test_String_ReplaceAll_FromStringConcat(t *testing.T) {
	// Arrange
	s := args.String("hello world")
	result := s.ReplaceAll("world", "go")

	// Act
	actual := args.Map{
		"result": result.String(),
	}

	// Assert
	expected := args.Map{
		"result": "hello go",
	}
	expected.ShouldBeEqual(t, 0, "String_ReplaceAll returns correct value -- with args", actual)
}

func Test_String_Substring_FromStringConcat(t *testing.T) {
	// Arrange
	s := args.String("hello world")
	result := s.Substring(0, 5)

	// Act
	actual := args.Map{
		"result": result.String(),
	}

	// Assert
	expected := args.Map{
		"result": "hello",
	}
	expected.ShouldBeEqual(t, 0, "String_Substring returns correct value -- with args", actual)
}

func Test_String_Bytes(t *testing.T) {
	// Arrange
	s := args.String("hi")

	// Act
	actual := args.Map{
		"bytesLen": len(s.Bytes()),
		"runesLen": len(s.Runes()),
	}

	// Assert
	expected := args.Map{
		"bytesLen": 2,
		"runesLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "String_Bytes returns correct value -- with args", actual)
}

// ── Dynamic ──

func Test_Dynamic_Getters_FromStringConcat(t *testing.T) {
	// Arrange
	d := &args.DynamicAny{
		Params: args.Map{
			"first":  "f1",
			"second": "f2",
			"third":  "f3",
			"fourth": "f4",
			"fifth":  "f5",
			"sixth":  "f6",
		},
		Expect: "exp",
	}

	// Act
	actual := args.Map{
		"first":    d.FirstItem(),
		"second":   d.SecondItem(),
		"third":    d.ThirdItem(),
		"fourth":   d.FourthItem(),
		"fifth":    d.FifthItem(),
		"sixth":    d.SixthItem(),
		"expected": d.Expected(),
		"hasFirst": d.HasFirst(),
	}

	// Assert
	expected := args.Map{
		"first":    "f1",
		"second":   "f2",
		"third":    "f3",
		"fourth":   "f4",
		"fifth":    "f5",
		"sixth":    "f6",
		"expected": "exp",
		"hasFirst": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_Getters returns correct value -- with args", actual)
}

func Test_Dynamic_NilSafety_FromStringConcat(t *testing.T) {
	// Arrange
	var d *args.DynamicAny

	// Act
	actual := args.Map{
		"argsCount":     d.ArgsCount(),
		"getWorkFunc":   d.GetWorkFunc() == nil,
		"hasFirst":      d.HasFirst(),
		"hasDefined":    d.HasDefined("key"),
		"has":           d.Has("key"),
		"hasDefinedAll": d.HasDefinedAll("key"),
		"isKeyInvalid":  d.IsKeyInvalid("key"),
		"isKeyMissing":  d.IsKeyMissing("key"),
		"hasExpect":     d.HasExpect(),
	}

	// Assert
	expected := args.Map{
		"argsCount":     0,
		"getWorkFunc":   true,
		"hasFirst":      false,
		"hasDefined":    false,
		"has":           false,
		"hasDefinedAll": false,
		"isKeyInvalid":  false,
		"isKeyMissing":  false,
		"hasExpect":     false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_NilSafety returns nil -- with args", actual)
}

func Test_Dynamic_TypedGetters(t *testing.T) {
	// Arrange
	d := &args.DynamicAny{
		Params: args.Map{
			"count":  5,
			"name":   "test",
			"items":  []string{"a", "b"},
			"anyArr": []any{1, 2},
		},
	}

	intVal, intOk := d.GetAsInt("count")
	strVal, strOk := d.GetAsString("name")
	stringsVal, stringsOk := d.GetAsStrings("items")
	anyItems, anyOk := d.GetAsAnyItems("anyArr")
	intDefault := d.GetAsIntDefault("missing", 99)
	strDefault := d.GetAsStringDefault("missing")

	// Act
	actual := args.Map{
		"intVal":     intVal,
		"intOk":      intOk,
		"strVal":     strVal,
		"strOk":      strOk,
		"stringsLen": len(stringsVal),
		"stringsOk":  stringsOk,
		"anyLen":     len(anyItems),
		"anyOk":      anyOk,
		"intDefault": intDefault,
		"strDefault": strDefault,
	}

	// Assert
	expected := args.Map{
		"intVal":     5,
		"intOk":      true,
		"strVal":     "test",
		"strOk":      true,
		"stringsLen": 2,
		"stringsOk":  true,
		"anyLen":     2,
		"anyOk":      true,
		"intDefault": 99,
		"strDefault": "",
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_TypedGetters returns correct value -- with args", actual)
}

func Test_Dynamic_GetLowerCase_FromStringConcat(t *testing.T) {
	// Arrange
	d := &args.DynamicAny{
		Params: args.Map{"name": "val"},
	}

	item, ok := d.GetLowerCase("Name")
	direct := d.GetDirectLower("Name")

	// Act
	actual := args.Map{
		"item":   item,
		"ok":     ok,
		"direct": direct,
	}

	// Assert
	expected := args.Map{
		"item":   "val",
		"ok":     true,
		"direct": "val",
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_GetLowerCase returns correct value -- with args", actual)
}

func Test_Dynamic_HasDefined_FromStringConcat(t *testing.T) {
	// Arrange
	d := &args.DynamicAny{
		Params: args.Map{
			"key": "val",
			"null": nil,
		},
	}

	// Act
	actual := args.Map{
		"hasDefined":     d.HasDefined("key"),
		"hasNull":        d.HasDefined("null"),
		"hasMissing":     d.HasDefined("missing"),
		"has":            d.Has("key"),
		"hasMissingKey":  d.Has("missing"),
		"isKeyInvalid":   d.IsKeyInvalid("null"),
		"isKeyMissing":   d.IsKeyMissing("missing"),
		"hasDefinedAll":  d.HasDefinedAll("key"),
		"hasDefinedNone": d.HasDefinedAll(),
	}

	// Assert
	expected := args.Map{
		"hasDefined":     true,
		"hasNull":        false,
		"hasMissing":     false,
		"has":            true,
		"hasMissingKey":  false,
		"isKeyInvalid":   true,
		"isKeyMissing":   true,
		"hasDefinedAll":  true,
		"hasDefinedNone": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_HasDefined returns correct value -- with args", actual)
}

func Test_Dynamic_ValidArgs_FromStringConcat(t *testing.T) {
	// Arrange
	d := &args.DynamicAny{
		Params: args.Map{
			"a": 1,
			"b": 2,
		},
	}

	validArgs := d.ValidArgs()
	customArgs := d.Args("a")

	// Act
	actual := args.Map{
		"validLen":  len(validArgs),
		"customLen": len(customArgs),
	}

	// Assert
	expected := args.Map{
		"validLen":  2,
		"customLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_ValidArgs returns non-empty -- with args", actual)
}

func Test_Dynamic_Slice(t *testing.T) {
	// Arrange
	d := &args.DynamicAny{
		Params: args.Map{"a": 1},
		Expect: "exp",
	}

	slice := d.Slice()
	slice2 := d.Slice() // cached

	// Act
	actual := args.Map{
		"sliceLen":  len(slice),
		"cachedLen": len(slice2),
	}

	// Assert
	expected := args.Map{
		"sliceLen":  len(slice),
		"cachedLen": len(slice2),
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_Slice returns correct value -- with args", actual)
}

func Test_Dynamic_String(t *testing.T) {
	// Arrange
	d := &args.DynamicAny{
		Params: args.Map{"a": 1},
	}

	s := d.String()

	// Act
	actual := args.Map{
		"notEmpty": s != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_String returns correct value -- with args", actual)
}

func Test_Dynamic_Actual(t *testing.T) {
	// Arrange
	d := &args.DynamicAny{
		Params: args.Map{
			"actual": "myActual",
			"arrange": "myArrange",
		},
	}

	// Act
	actual := args.Map{
		"actual":  d.Actual(),
		"arrange": d.Arrange(),
	}

	// Assert
	expected := args.Map{
		"actual":  "myActual",
		"arrange": "myArrange",
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_Actual returns correct value -- with args", actual)
}

func Test_Dynamic_Contracts(t *testing.T) {
	// Arrange
	d := args.DynamicAny{
		Params: args.Map{"a": 1},
	}

	// Act
	actual := args.Map{
		"mapper":        d.AsArgsMapper() != nil,
		"funcNameBind":  d.AsArgFuncNameContractsBinder() != nil,
		"baseBind":      d.AsArgBaseContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"mapper":        true,
		"funcNameBind":  true,
		"baseBind":      true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic_Contracts returns correct value -- with args", actual)
}

// ── Holder ──

func Test_Holder_Getters(t *testing.T) {
	// Arrange
	h := &args.HolderAny{
		First:  "f1",
		Second: "f2",
		Third:  "f3",
		Fourth: "f4",
		Fifth:  "f5",
		Sixth:  "f6",
		Expect: "exp",
	}

	// Act
	actual := args.Map{
		"first":    h.FirstItem(),
		"second":   h.SecondItem(),
		"third":    h.ThirdItem(),
		"fourth":   h.FourthItem(),
		"fifth":    h.FifthItem(),
		"sixth":    h.SixthItem(),
		"expected": h.Expected(),
		"count":    h.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first":    "f1",
		"second":   "f2",
		"third":    "f3",
		"fourth":   "f4",
		"fifth":    "f5",
		"sixth":    "f6",
		"expected": "exp",
		"count":    7,
	}
	expected.ShouldBeEqual(t, 0, "Holder_Getters returns correct value -- with args", actual)
}

func Test_Holder_HasMethods(t *testing.T) {
	// Arrange
	h := &args.HolderAny{
		First:  "f1",
		Second: "f2",
		Third:  "f3",
		Fourth: "f4",
		Fifth:  "f5",
		Sixth:  "f6",
		Expect: "exp",
	}

	// Act
	actual := args.Map{
		"hasFirst":  h.HasFirst(),
		"hasSecond": h.HasSecond(),
		"hasThird":  h.HasThird(),
		"hasFourth": h.HasFourth(),
		"hasFifth":  h.HasFifth(),
		"hasSixth":  h.HasSixth(),
		"hasExpect": h.HasExpect(),
		"hasFunc":   h.HasFunc(),
	}

	// Assert
	expected := args.Map{
		"hasFirst":  true,
		"hasSecond": true,
		"hasThird":  true,
		"hasFourth": true,
		"hasFifth":  true,
		"hasSixth":  true,
		"hasExpect": true,
		"hasFunc":   false,
	}
	expected.ShouldBeEqual(t, 0, "Holder_HasMethods returns correct value -- with args", actual)
}

func Test_Holder_ArgTwo_FromStringConcat(t *testing.T) {
	// Arrange
	h := &args.HolderAny{First: "a", Second: "b"}
	two := h.ArgTwo()

	// Act
	actual := args.Map{
		"first":  two.First,
		"second": two.Second,
	}

	// Assert
	expected := args.Map{
		"first":  "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "Holder_ArgTwo returns correct value -- with args", actual)
}

func Test_Holder_ArgThree(t *testing.T) {
	// Arrange
	h := &args.HolderAny{First: "a", Second: "b", Third: "c"}
	three := h.ArgThree()

	// Act
	actual := args.Map{
		"first": three.First,
		"third": three.Third,
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"third": "c",
	}
	expected.ShouldBeEqual(t, 0, "Holder_ArgThree returns correct value -- with args", actual)
}

func Test_Holder_Args_FromStringConcat(t *testing.T) {
	// Arrange
	h := &args.HolderAny{First: "a", Second: "b", Third: "c"}

	args1 := h.Args(1)
	args2 := h.Args(2)
	args3 := h.Args(3)

	// Act
	actual := args.Map{
		"args1Len": len(args1),
		"args2Len": len(args2),
		"args3Len": len(args3),
	}

	// Assert
	expected := args.Map{
		"args1Len": 1,
		"args2Len": 2,
		"args3Len": 3,
	}
	expected.ShouldBeEqual(t, 0, "Holder_Args returns correct value -- with args", actual)
}

func Test_Holder_ValidArgs(t *testing.T) {
	// Arrange
	h := &args.HolderAny{First: "a", Second: "b"}

	va := h.ValidArgs()

	// Act
	actual := args.Map{
		"len": len(va),
	}

	// Assert
	expected := args.Map{
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "Holder_ValidArgs returns non-empty -- with args", actual)
}

func Test_Holder_Slice_FromStringConcat(t *testing.T) {
	// Arrange
	h := &args.HolderAny{First: "a", Expect: "exp"}

	s1 := h.Slice()
	s2 := h.Slice() // cached

	// Act
	actual := args.Map{
		"len":      len(s1),
		"cachedEq": len(s1) == len(s2),
	}

	// Assert
	expected := args.Map{
		"len":      2,
		"cachedEq": true,
	}
	expected.ShouldBeEqual(t, 0, "Holder_Slice returns correct value -- with args", actual)
}

func Test_Holder_GetByIndex(t *testing.T) {
	// Arrange
	h := &args.HolderAny{First: "a"}

	// Act
	actual := args.Map{
		"valid":   h.GetByIndex(0),
		"invalid": h.GetByIndex(99) == nil,
	}

	// Assert
	expected := args.Map{
		"valid":   "a",
		"invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "Holder_GetByIndex returns correct value -- with args", actual)
}

func Test_Holder_String_FromStringConcat(t *testing.T) {
	// Arrange
	h := &args.HolderAny{First: "a"}

	// Act
	actual := args.Map{
		"notEmpty": h.String() != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Holder_String returns correct value -- with args", actual)
}

// ── LeftRight ──

func Test_LeftRight_Methods_FromStringConcat(t *testing.T) {
	// Arrange
	lr := &args.LeftRightAny{
		Left:   "L",
		Right:  "R",
		Expect: "exp",
	}

	// Act
	actual := args.Map{
		"first":     lr.FirstItem(),
		"second":    lr.SecondItem(),
		"expected":  lr.Expected(),
		"argsCount": lr.ArgsCount(),
		"hasFirst":  lr.HasFirst(),
		"hasSecond": lr.HasSecond(),
		"hasLeft":   lr.HasLeft(),
		"hasRight":  lr.HasRight(),
		"hasExpect": lr.HasExpect(),
	}

	// Assert
	expected := args.Map{
		"first":     "L",
		"second":    "R",
		"expected":  "exp",
		"argsCount": 2,
		"hasFirst":  true,
		"hasSecond": true,
		"hasLeft":   true,
		"hasRight":  true,
		"hasExpect": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight_Methods returns correct value -- with args", actual)
}

func Test_LeftRight_ArgTwo_FromStringConcat(t *testing.T) {
	// Arrange
	lr := &args.LeftRightAny{Left: "L", Right: "R"}
	two := lr.ArgTwo()

	// Act
	actual := args.Map{
		"first":  two.First,
		"second": two.Second,
	}

	// Assert
	expected := args.Map{
		"first":  "L",
		"second": "R",
	}
	expected.ShouldBeEqual(t, 0, "LeftRight_ArgTwo returns correct value -- with args", actual)
}

func Test_LeftRight_ValidArgs(t *testing.T) {
	// Arrange
	lr := &args.LeftRightAny{Left: "L", Right: "R"}
	va := lr.ValidArgs()
	a := lr.Args(2)

	// Act
	actual := args.Map{
		"vaLen": len(va),
		"aLen":  len(a),
	}

	// Assert
	expected := args.Map{
		"vaLen": 2,
		"aLen":  2,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight_ValidArgs returns non-empty -- with args", actual)
}

func Test_LeftRight_Clone_FromStringConcat(t *testing.T) {
	// Arrange
	lr := &args.LeftRightAny{Left: "L", Right: "R", Expect: "e"}
	clone := lr.Clone()

	// Act
	actual := args.Map{
		"left":   clone.Left,
		"right":  clone.Right,
		"expect": clone.Expect,
	}

	// Assert
	expected := args.Map{
		"left":   "L",
		"right":  "R",
		"expect": "e",
	}
	expected.ShouldBeEqual(t, 0, "LeftRight_Clone returns correct value -- with args", actual)
}

func Test_LeftRight_String_FromStringConcat(t *testing.T) {
	// Arrange
	lr := &args.LeftRightAny{Left: "L"}

	// Act
	actual := args.Map{
		"notEmpty": lr.String() != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight_String returns correct value -- with args", actual)
}

// ── Two ──

func Test_Two_Methods_FromStringConcat(t *testing.T) {
	// Arrange
	tw := &args.TwoAny{First: "a", Second: "b", Expect: "exp"}

	// Act
	actual := args.Map{
		"first":     tw.FirstItem(),
		"second":    tw.SecondItem(),
		"expected":  tw.Expected(),
		"argsCount": tw.ArgsCount(),
		"hasFirst":  tw.HasFirst(),
		"hasSecond": tw.HasSecond(),
		"hasExpect": tw.HasExpect(),
	}

	// Assert
	expected := args.Map{
		"first":     "a",
		"second":    "b",
		"expected":  "exp",
		"argsCount": 2,
		"hasFirst":  true,
		"hasSecond": true,
		"hasExpect": true,
	}
	expected.ShouldBeEqual(t, 0, "Two_Methods returns correct value -- with args", actual)
}

func Test_Two_ArgTwo_FromStringConcat(t *testing.T) {
	// Arrange
	tw := &args.TwoAny{First: "a", Second: "b"}
	at := tw.ArgTwo()

	// Act
	actual := args.Map{
		"first":  at.First,
		"second": at.Second,
	}

	// Assert
	expected := args.Map{
		"first":  "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "Two_ArgTwo returns correct value -- with args", actual)
}

func Test_Two_LeftRight_FromStringConcat(t *testing.T) {
	// Arrange
	tw := &args.TwoAny{First: "a", Second: "b", Expect: "e"}
	lr := tw.LeftRight()

	// Act
	actual := args.Map{
		"left":   lr.Left,
		"right":  lr.Right,
		"expect": lr.Expect,
	}

	// Assert
	expected := args.Map{
		"left":   "a",
		"right":  "b",
		"expect": "e",
	}
	expected.ShouldBeEqual(t, 0, "Two_LeftRight returns correct value -- with args", actual)
}

func Test_Two_String_FromStringConcat(t *testing.T) {
	// Arrange
	tw := &args.TwoAny{First: "a", Second: "b"}

	// Act
	actual := args.Map{
		"notEmpty": tw.String() != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Two_String returns correct value -- with args", actual)
}

// ── Three ──

func Test_Three_Methods_FromStringConcat(t *testing.T) {
	// Arrange
	th := &args.ThreeAny{First: "a", Second: "b", Third: "c", Expect: "exp"}

	// Act
	actual := args.Map{
		"first":     th.FirstItem(),
		"second":    th.SecondItem(),
		"third":     th.ThirdItem(),
		"expected":  th.Expected(),
		"argsCount": th.ArgsCount(),
		"hasFirst":  th.HasFirst(),
		"hasSecond": th.HasSecond(),
		"hasThird":  th.HasThird(),
		"hasExpect": th.HasExpect(),
	}

	// Assert
	expected := args.Map{
		"first":     "a",
		"second":    "b",
		"third":     "c",
		"expected":  "exp",
		"argsCount": 3,
		"hasFirst":  true,
		"hasSecond": true,
		"hasThird":  true,
		"hasExpect": true,
	}
	expected.ShouldBeEqual(t, 0, "Three_Methods returns correct value -- with args", actual)
}

func Test_Three_ArgTwo(t *testing.T) {
	// Arrange
	th := &args.ThreeAny{First: "a", Second: "b", Third: "c"}
	tw := th.ArgTwo()
	at := th.ArgThree()

	// Act
	actual := args.Map{
		"twoFirst":   tw.First,
		"threeThird": at.Third,
	}

	// Assert
	expected := args.Map{
		"twoFirst":   "a",
		"threeThird": "c",
	}
	expected.ShouldBeEqual(t, 0, "Three_ArgTwo returns correct value -- with args", actual)
}

func Test_Three_LeftRight_FromStringConcat(t *testing.T) {
	// Arrange
	th := &args.ThreeAny{First: "a", Second: "b", Expect: "e"}
	lr := th.LeftRight()

	// Act
	actual := args.Map{
		"left":  lr.Left,
		"right": lr.Right,
	}

	// Assert
	expected := args.Map{
		"left":  "a",
		"right": "b",
	}
	expected.ShouldBeEqual(t, 0, "Three_LeftRight returns correct value -- with args", actual)
}

func Test_Three_String_FromStringConcat(t *testing.T) {
	// Arrange
	th := &args.ThreeAny{First: "a"}

	// Act
	actual := args.Map{
		"notEmpty": th.String() != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Three_String returns correct value -- with args", actual)
}

// ── FuncWrap ──

func Test_FuncWrap_NewTypedFuncWrap(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }
	fw := args.NewTypedFuncWrap(fn)

	// Act
	actual := args.Map{
		"isValid":   fw.IsValid(),
		"hasName":   fw.GetFuncName() != "",
		"argsCount": fw.ArgsCount(),
		"returnLen": fw.ReturnLength(),
	}

	// Assert
	expected := args.Map{
		"isValid":   true,
		"hasName":   true,
		"argsCount": 1,
		"returnLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_NewTypedFuncWrap returns correct value -- with args", actual)
}

func Test_FuncWrap_NilFunc(t *testing.T) {
	// Arrange
	fw := args.NewTypedFuncWrap[any](nil)

	// Act
	actual := args.Map{
		"isInvalid": fw.IsInvalid(),
		"getName":   fw.GetFuncName(),
	}

	// Assert
	expected := args.Map{
		"isInvalid": true,
		"getName":   "",
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_NilFunc returns nil -- with args", actual)
}

func Test_FuncWrap_NonFunc(t *testing.T) {
	// Arrange
	fw := args.NewTypedFuncWrap(42)

	// Act
	actual := args.Map{
		"isInvalid": fw.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_NonFunc returns correct value -- with args", actual)
}

func Test_FuncWrap_Invoke(t *testing.T) {
	// Arrange
	fn := func(a, b int) int { return a + b }
	fw := args.NewTypedFuncWrap(fn)

	results, err := fw.Invoke(3, 4)

	// Act
	actual := args.Map{
		"noErr":  err == nil,
		"result": results[0],
	}

	// Assert
	expected := args.Map{
		"noErr":  true,
		"result": 7,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_Invoke returns correct value -- with args", actual)
}

func Test_FuncWrap_InvokeMust_FromStringConcat(t *testing.T) {
	// Arrange
	fn := func() string { return "ok" }
	fw := args.NewTypedFuncWrap(fn)

	results := fw.InvokeMust()

	// Act
	actual := args.Map{
		"result": results[0],
	}

	// Assert
	expected := args.Map{
		"result": "ok",
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_InvokeMust returns correct value -- with args", actual)
}

func Test_FuncWrap_VoidCall_FromStringConcat(t *testing.T) {
	// Arrange
	called := false
	fn := func() { called = true }
	fw := args.NewTypedFuncWrap(fn)

	_, err := fw.VoidCall()

	// Act
	actual := args.Map{
		"noErr":  err == nil,
		"called": called,
	}

	// Assert
	expected := args.Map{
		"noErr":  true,
		"called": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_VoidCall returns correct value -- with args", actual)
}

func Test_FuncWrap_ValidationError(t *testing.T) {
	// Arrange
	fn := func() {}
	fw := args.NewTypedFuncWrap(fn)
	nilFw := args.NewTypedFuncWrap[any](nil)

	// Act
	actual := args.Map{
		"validNoErr": fw.ValidationError() == nil,
		"nilHasErr":  nilFw.ValidationError() != nil,
	}

	// Assert
	expected := args.Map{
		"validNoErr": true,
		"nilHasErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_ValidationError returns error -- with args", actual)
}

func Test_FuncWrap_InvalidError(t *testing.T) {
	// Arrange
	fn := func() {}
	fw := args.NewTypedFuncWrap(fn)
	nilFw := args.NewTypedFuncWrap[any](nil)

	// Act
	actual := args.Map{
		"validNil": fw.InvalidError() == nil,
		"nilErr":   nilFw.InvalidError() != nil,
	}

	// Assert
	expected := args.Map{
		"validNil": true,
		"nilErr":   true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_InvalidError returns error -- with args", actual)
}

func Test_FuncWrap_IsEqual_FromStringConcat(t *testing.T) {
	// Arrange
	fn := func() {}
	fw1 := args.NewTypedFuncWrap(fn)
	fw2 := args.NewTypedFuncWrap(fn)

	// Act
	actual := args.Map{
		"selfEqual":   fw1.IsEqual(fw1),
		"sameFunc":    fw1.IsEqual(fw2),
		"nilBoth":     (*args.FuncWrapAny)(nil).IsEqual(nil),
		"nilOne":      fw1.IsEqual(nil),
		"isNotEqual":  fw1.IsNotEqual(nil),
		"isEqualVal":  fw1.IsEqualValue(*fw2),
	}

	// Assert
	expected := args.Map{
		"selfEqual":   true,
		"sameFunc":    true,
		"nilBoth":     true,
		"nilOne":      false,
		"isNotEqual":  true,
		"isEqualVal":  true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_IsEqual returns correct value -- with args", actual)
}

func Test_FuncWrap_PkgPath_FromStringConcat(t *testing.T) {
	// Arrange
	fn := func() {}
	fw := args.NewTypedFuncWrap(fn)

	// Act
	actual := args.Map{
		"pkgPath":    fw.PkgPath() != "",
		"pkgName":    fw.PkgNameOnly() != "",
		"directName": fw.FuncDirectInvokeName() != "",
		"getType":    fw.GetType() != nil,
		"pascal":     fw.GetPascalCaseFuncName() != "",
	}

	// Assert
	expected := args.Map{
		"pkgPath":    true,
		"pkgName":    true,
		"directName": true,
		"getType":    true,
		"pascal":     true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_PkgPath returns correct value -- with args", actual)
}

func Test_FuncWrap_GetFirstResponseOfInvoke(t *testing.T) {
	// Arrange
	fn := func() string { return "first" }
	fw := args.NewTypedFuncWrap(fn)

	result, err := fw.GetFirstResponseOfInvoke()

	// Act
	actual := args.Map{
		"result": result,
		"noErr":  err == nil,
	}

	// Assert
	expected := args.Map{
		"result": "first",
		"noErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_GetFirstResponseOfInvoke returns correct value -- with args", actual)
}

func Test_FuncWrap_InvokeResultOfIndex_FromStringConcat(t *testing.T) {
	// Arrange
	fn := func() (string, int) { return "a", 42 }
	fw := args.NewTypedFuncWrap(fn)

	r0, err0 := fw.InvokeResultOfIndex(0)
	r1, err1 := fw.InvokeResultOfIndex(1)

	// Act
	actual := args.Map{
		"r0":    r0,
		"err0":  err0 == nil,
		"r1":    r1,
		"err1":  err1 == nil,
	}

	// Assert
	expected := args.Map{
		"r0":    "a",
		"err0":  true,
		"r1":    42,
		"err1":  true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap_InvokeResultOfIndex returns correct value -- with args", actual)
}

// ── Map CompileToString ──

func Test_Map_CompileToString_FromStringConcat(t *testing.T) {
	// Arrange
	m := args.Map{
		"b": 2,
		"a": 1,
	}
	result := m.CompileToString()

	// Act
	actual := args.Map{
		"notEmpty": result != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Map_CompileToString returns correct value -- with args", actual)
}

func Test_Map_CompileToString_Empty(t *testing.T) {
	// Arrange
	m := args.Map{}
	result := m.CompileToString()

	// Act
	actual := args.Map{
		"isEmpty": result == "",
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Map_CompileToString_Empty returns empty -- with args", actual)
}
