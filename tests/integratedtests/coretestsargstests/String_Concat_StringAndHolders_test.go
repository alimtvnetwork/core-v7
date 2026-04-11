package coretestsargstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ── args.String methods ──

func Test_String_Concat(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act
	result := s.Concat("-", "world")

	// Assert
	actual := args.Map{"result": result.String()}
	expected := args.Map{"result": "hello-world"}
	expected.ShouldBeEqual(t, 0, "Concat returns joined -- two parts", actual)
}

func Test_String_Join(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act
	result := s.Join("-", "world")

	// Assert
	actual := args.Map{"result": result.String()}
	expected := args.Map{"result": "hello-world"}
	expected.ShouldBeEqual(t, 0, "Join returns joined -- with separator", actual)
}

func Test_String_Split(t *testing.T) {
	// Arrange
	s := args.String("a-b-c")

	// Act
	result := s.Split("-")

	// Assert
	actual := args.Map{"length": len(result)}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "Split returns parts -- 3 segments", actual)
}

func Test_String_Quotes(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act
	dq := s.DoubleQuote()
	sq := s.SingleQuote()
	dqq := s.DoubleQuoteQ()
	vdq := s.ValueDoubleQuote()

	// Assert
	actual := args.Map{
		"dqNotEmpty":  len(dq) > 0,
		"sqNotEmpty":  len(sq) > 0,
		"dqqNotEmpty": len(dqq) > 0,
		"vdqNotEmpty": len(vdq) > 0,
	}
	expected := args.Map{
		"dqNotEmpty":  true,
		"sqNotEmpty":  true,
		"dqqNotEmpty": true,
		"vdqNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Quote methods return non-empty -- all variants", actual)
}

func Test_String_Properties(t *testing.T) {
	// Arrange
	s := args.String("hello")
	empty := args.String("")

	// Act & Assert
	actual := args.Map{
		"bytes":              len(s.Bytes()) > 0,
		"runes":              len(s.Runes()) > 0,
		"length":             s.Length(),
		"count":              s.Count(),
		"asciiLen":           s.AscIILength(),
		"isEmpty":            s.IsEmpty(),
		"emptyIsEmpty":       empty.IsEmpty(),
		"hasChar":            s.HasCharacter(),
		"isDefined":          s.IsDefined(),
		"emptyOrWhitespace":  empty.IsEmptyOrWhitespace(),
	}
	expected := args.Map{
		"bytes":              true,
		"runes":              true,
		"length":             5,
		"count":              5,
		"asciiLen":           5,
		"isEmpty":            false,
		"emptyIsEmpty":       true,
		"hasChar":            true,
		"isDefined":          true,
		"emptyOrWhitespace":  true,
	}
	expected.ShouldBeEqual(t, 0, "String properties return correct -- various methods", actual)
}

func Test_String_TrimAndReplace(t *testing.T) {
	// Arrange
	s := args.String("  hello  ")

	// Act
	trimmed := s.TrimSpace()
	replaced := args.String("hello-world").ReplaceAll("-", "_")
	sub := args.String("abcdef").Substring(1, 4)

	// Assert
	actual := args.Map{
		"trimmed":  trimmed.String(),
		"replaced": replaced.String(),
		"sub":      sub.String(),
	}
	expected := args.Map{
		"trimmed":  "hello",
		"replaced": "hello_world",
		"sub":      "bcd",
	}
	expected.ShouldBeEqual(t, 0, "TrimSpace/ReplaceAll/Substring return correct -- basic ops", actual)
}

func Test_String_TrimReplaceMap(t *testing.T) {
	// Arrange
	s := args.String("{name}-{age}")

	// Act
	result := s.TrimReplaceMap(map[string]string{
		"{name}": "John",
		"{age}":  "30",
	})

	// Assert
	actual := args.Map{"result": result.String()}
	expected := args.Map{"result": "John-30"}
	expected.ShouldBeEqual(t, 0, "TrimReplaceMap returns replaced -- template keys", actual)
}

// ── One methods ──

func Test_One_Methods(t *testing.T) {
	// Arrange
	one := args.OneAny{
		First:  "hello",
		Expect: 42,
	}

	// Act & Assert
	actual := args.Map{
		"firstItem": one.FirstItem(),
		"expected":  one.Expected(),
		"hasFirst":  one.HasFirst(),
		"hasExpect": one.HasExpect(),
		"argsCount": one.ArgsCount(),
		"sliceLen":  len(one.Slice()),
		"strLen":    len(one.String()) > 0,
	}
	expected := args.Map{
		"firstItem": "hello",
		"expected":  42,
		"hasFirst":  true,
		"hasExpect": true,
		"argsCount": 1,
		"sliceLen":  2,
		"strLen":    true,
	}
	expected.ShouldBeEqual(t, 0, "One methods return correct -- all accessors", actual)
}

func Test_One_LeftRight(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello"}

	// Act
	lr := one.LeftRight()

	// Assert
	actual := args.Map{"left": lr.Left}
	expected := args.Map{"left": "hello"}
	expected.ShouldBeEqual(t, 0, "One.LeftRight returns correct -- first as left", actual)
}

func Test_One_Args(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello"}

	// Act
	a1 := one.Args(1)
	a0 := one.Args(0)

	// Assert
	actual := args.Map{
		"args1Len": len(a1),
		"args0Len": len(a0),
	}
	expected := args.Map{
		"args1Len": 1,
		"args0Len": 0,
	}
	expected.ShouldBeEqual(t, 0, "One.Args returns correct -- upTo values", actual)
}

// ── Two methods ──

func Test_Two_Methods(t *testing.T) {
	// Arrange
	two := args.TwoAny{
		First:  "hello",
		Second: 42,
		Expect: true,
	}

	// Act & Assert
	actual := args.Map{
		"firstItem":  two.FirstItem(),
		"secondItem": two.SecondItem(),
		"hasFirst":   two.HasFirst(),
		"hasSecond":  two.HasSecond(),
		"hasExpect":  two.HasExpect(),
		"argsCount":  two.ArgsCount(),
	}
	expected := args.Map{
		"firstItem":  "hello",
		"secondItem": 42,
		"hasFirst":   true,
		"hasSecond":  true,
		"hasExpect":  true,
		"argsCount":  2,
	}
	expected.ShouldBeEqual(t, 0, "Two methods return correct -- all accessors", actual)
}

func Test_Two_LeftRight(t *testing.T) {
	// Arrange
	two := args.TwoAny{First: "a", Second: "b"}

	// Act
	lr := two.LeftRight()

	// Assert
	actual := args.Map{
		"left":  lr.Left,
		"right": lr.Right,
	}
	expected := args.Map{
		"left":  "a",
		"right": "b",
	}
	expected.ShouldBeEqual(t, 0, "Two.LeftRight returns correct -- both fields", actual)
}

// ── Three methods ──

func Test_Three_Methods(t *testing.T) {
	// Arrange
	three := args.ThreeAny{
		First:  "a",
		Second: "b",
		Third:  "c",
	}

	// Act & Assert
	actual := args.Map{
		"thirdItem": three.ThirdItem(),
		"hasThird":  three.HasThird(),
		"argsCount": three.ArgsCount(),
		"args2Len":  len(three.Args(2)),
		"args3Len":  len(three.Args(3)),
	}
	expected := args.Map{
		"thirdItem": "c",
		"hasThird":  true,
		"argsCount": 3,
		"args2Len":  2,
		"args3Len":  3,
	}
	expected.ShouldBeEqual(t, 0, "Three methods return correct -- all accessors", actual)
}

// ── Four methods ──

func Test_Four_Methods(t *testing.T) {
	// Arrange
	four := args.FourAny{
		First:  1,
		Second: 2,
		Third:  3,
		Fourth: 4,
	}

	// Act & Assert
	actual := args.Map{
		"fourthItem": four.FourthItem(),
		"hasFourth":  four.HasFourth(),
		"argsCount":  four.ArgsCount(),
		"args4Len":   len(four.Args(4)),
	}
	expected := args.Map{
		"fourthItem": 4,
		"hasFourth":  true,
		"argsCount":  4,
		"args4Len":   4,
	}
	expected.ShouldBeEqual(t, 0, "Four methods return correct -- all accessors", actual)
}

// ── Five methods ──

func Test_Five_Methods(t *testing.T) {
	// Arrange
	five := args.FiveAny{
		First:  1,
		Second: 2,
		Third:  3,
		Fourth: 4,
		Fifth:  5,
	}

	// Act & Assert
	actual := args.Map{
		"fifthItem": five.FifthItem(),
		"hasFifth":  five.HasFifth(),
		"argsCount": five.ArgsCount(),
		"args5Len":  len(five.Args(5)),
	}
	expected := args.Map{
		"fifthItem": 5,
		"hasFifth":  true,
		"argsCount": 5,
		"args5Len":  5,
	}
	expected.ShouldBeEqual(t, 0, "Five methods return correct -- all accessors", actual)
}

// ── Six methods ──

func Test_Six_Methods(t *testing.T) {
	// Arrange
	six := args.SixAny{
		First:  1,
		Second: 2,
		Third:  3,
		Fourth: 4,
		Fifth:  5,
		Sixth:  6,
	}

	// Act & Assert
	actual := args.Map{
		"sixthItem": six.SixthItem(),
		"hasSixth":  six.HasSixth(),
		"argsCount": six.ArgsCount(),
		"args6Len":  len(six.Args(6)),
	}
	expected := args.Map{
		"sixthItem": 6,
		"hasSixth":  true,
		"argsCount": 6,
		"args6Len":  6,
	}
	expected.ShouldBeEqual(t, 0, "Six methods return correct -- all accessors", actual)
}

// ── LeftRight methods ──

func Test_LeftRight_Clone(t *testing.T) {
	// Arrange
	lr := args.LeftRightAny{
		Left:   "a",
		Right:  "b",
		Expect: true,
	}

	// Act
	cloned := lr.Clone()

	// Assert
	actual := args.Map{
		"left":   cloned.Left,
		"right":  cloned.Right,
		"expect": cloned.Expect,
	}
	expected := args.Map{
		"left":   "a",
		"right":  "b",
		"expect": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight.Clone returns copy -- all fields", actual)
}

func Test_LeftRight_HasLeftRight(t *testing.T) {
	// Arrange
	lr := args.LeftRightAny{Left: "a", Right: "b"}

	// Act & Assert
	actual := args.Map{
		"hasLeft":  lr.HasLeft(),
		"hasRight": lr.HasRight(),
	}
	expected := args.Map{
		"hasLeft":  true,
		"hasRight": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight.HasLeft/HasRight return true -- defined", actual)
}

// ── Holder methods ──

func Test_Holder_Methods(t *testing.T) {
	// Arrange
	h := args.HolderAny{
		First:  "a",
		Second: "b",
		Third:  "c",
		Fourth: "d",
		Fifth:  "e",
		Sixth:  "f",
		Expect: true,
	}

	// Act & Assert
	actual := args.Map{
		"argsCount": h.ArgsCount(),
		"hasFirst":  h.HasFirst(),
		"hasSecond": h.HasSecond(),
		"hasThird":  h.HasThird(),
		"hasFourth": h.HasFourth(),
		"hasFifth":  h.HasFifth(),
		"hasSixth":  h.HasSixth(),
		"hasExpect": h.HasExpect(),
		"sliceLen":  len(h.Slice()) > 0,
		"strLen":    len(h.String()) > 0,
	}
	expected := args.Map{
		"argsCount": 7,
		"hasFirst":  true,
		"hasSecond": true,
		"hasThird":  true,
		"hasFourth": true,
		"hasFifth":  true,
		"hasSixth":  true,
		"hasExpect": true,
		"sliceLen":  true,
		"strLen":    true,
	}
	expected.ShouldBeEqual(t, 0, "Holder methods return correct -- all accessors", actual)
}

func Test_Holder_ArgHelpers(t *testing.T) {
	// Arrange
	h := args.HolderAny{
		First:  "a",
		Second: "b",
		Third:  "c",
		Fourth: "d",
		Fifth:  "e",
	}

	// Act
	two := h.ArgTwo()
	three := h.ArgThree()
	four := h.ArgFour()
	five := h.ArgFive()

	// Assert
	actual := args.Map{
		"twoFirst":   two.First,
		"threeThird": three.Third,
		"fourFourth": four.Fourth,
		"fiveFifth":  five.Fifth,
	}
	expected := args.Map{
		"twoFirst":   "a",
		"threeThird": "c",
		"fourFourth": "d",
		"fiveFifth":  "e",
	}
	expected.ShouldBeEqual(t, 0, "Holder Arg helpers return correct -- decomposed", actual)
}

func Test_Holder_Args(t *testing.T) {
	// Arrange
	h := args.HolderAny{
		First:  "a",
		Second: "b",
		Third:  "c",
		Fourth: "d",
		Fifth:  "e",
		Sixth:  "f",
	}

	// Act
	a3 := h.Args(3)
	a6 := h.Args(6)

	// Assert
	actual := args.Map{
		"args3Len": len(a3),
		"args6Len": len(a6),
	}
	expected := args.Map{
		"args3Len": 3,
		"args6Len": 6,
	}
	expected.ShouldBeEqual(t, 0, "Holder.Args returns correct -- upTo values", actual)
}

// ── Dynamic methods ──

func Test_Dynamic_Methods(t *testing.T) {
	// Arrange
	d := args.DynamicAny{
		Params: args.Map{
			"first": "a",
			"name":  "test",
		},
		Expect: "expected",
	}

	// Act & Assert
	actual := args.Map{
		"hasExpect":     d.HasExpect(),
		"expected":      d.Expected(),
		"hasFirst":      d.HasFirst(),
		"argsCount":     d.ArgsCount(),
		"hasDefined":    d.HasDefined("name"),
		"has":           d.Has("name"),
		"isKeyMissing":  d.IsKeyMissing("missing"),
		"isKeyInvalid":  d.IsKeyInvalid("missing"),
		"hasDefinedAll": d.HasDefinedAll("first", "name"),
	}
	expected := args.Map{
		"hasExpect":     true,
		"expected":      "expected",
		"hasFirst":      true,
		"argsCount":     1,
		"hasDefined":    true,
		"has":           true,
		"isKeyMissing":  true,
		"isKeyInvalid":  true,
		"hasDefinedAll": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic methods return correct -- various accessors", actual)
}

func Test_Dynamic_TypedAccessors(t *testing.T) {
	// Arrange
	d := args.DynamicAny{
		Params: args.Map{
			"count": 5,
			"name":  "test",
			"items": []string{"a", "b"},
		},
	}

	// Act
	intVal, intOk := d.GetAsInt("count")
	intDef := d.GetAsIntDefault("missing", 99)
	strVal, strOk := d.GetAsString("name")
	strDef := d.GetAsStringDefault("missing")
	strs, strsOk := d.GetAsStrings("items")

	// Assert
	actual := args.Map{
		"intVal":  intVal,
		"intOk":   intOk,
		"intDef":  intDef,
		"strVal":  strVal,
		"strOk":   strOk,
		"strDef":  strDef,
		"strsLen": len(strs),
		"strsOk":  strsOk,
	}
	expected := args.Map{
		"intVal":  5,
		"intOk":   true,
		"intDef":  99,
		"strVal":  "test",
		"strOk":   true,
		"strDef":  "",
		"strsLen": 2,
		"strsOk":  true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic typed accessors return correct -- all types", actual)
}

// ── FuncDetector ──

func Test_FuncDetector_GetFuncWrap_FromFunc(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }

	// Act
	fw := args.FuncDetector.GetFuncWrap(fn)

	// Assert
	actual := args.Map{
		"isValid": fw.IsValid(),
	}
	expected := args.Map{
		"isValid": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncDetector.GetFuncWrap returns valid -- from func", actual)
}

func Test_FuncDetector_GetFuncWrap_FromMap(t *testing.T) {
	// Arrange
	m := args.Map{
		"func": func() {},
	}

	// Act
	fw := args.FuncDetector.GetFuncWrap(m)

	// Assert
	actual := args.Map{
		"notNil": fw != nil,
	}
	expected := args.Map{
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncDetector.GetFuncWrap returns wrap -- from Map", actual)
}

// ── emptyCreator ──

func Test_Empty_Creators(t *testing.T) {
	// Arrange & Act
	m := args.Empty.Map()
	fw := args.Empty.FuncWrap()
	fm := args.Empty.FuncMap()
	h := args.Empty.Holder()

	// Assert
	actual := args.Map{
		"mapLen":     len(m),
		"fwInvalid":  fw.IsInvalid(),
		"fmLen":      fm.Length(),
		"holderArgs": h.ArgsCount(),
	}
	expected := args.Map{
		"mapLen":     0,
		"fwInvalid":  true,
		"fmLen":      0,
		"holderArgs": 7,
	}
	expected.ShouldBeEqual(t, 0, "Empty creators return correct -- all types", actual)
}
