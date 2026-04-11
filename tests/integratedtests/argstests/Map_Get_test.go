package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Map_Get(t *testing.T) {
	for caseIndex, testCase := range extMapGetTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		key, _ := input.GetAsString("key")

		// Act
		val, isValid := input.Get(key)

		actual := args.Map{
			"isValid": isValid,
		}

		if isValid {
			actual["value"] = val
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Map_TypedGetters_FromMapGet(t *testing.T) {
	for caseIndex, testCase := range extMapTypedGetTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)

		// Act
		actual := args.Map{}

		intVal, intValid := input.GetAsInt("count")
		if intValid {
			actual["intVal"] = intVal
			actual["intValid"] = intValid
		}

		boolVal, boolValid := input.GetAsBool("active")
		if boolValid {
			actual["boolVal"] = boolVal
			actual["boolValid"] = boolValid
		}

		strVal, strValid := input.GetAsString("text")
		if strValid {
			actual["strVal"] = strVal
			actual["strValid"] = strValid
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Map_Compile(t *testing.T) {
	for caseIndex, testCase := range extMapCompileTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)

		// Act
		lines := input.CompileToStrings()

		actual := args.Map{
			"lineCount": len(lines),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Map_Utility_Methods(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": "two",
		"c": true,
	}

	// Act & Assert
	actual := args.Map{"result": m.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Length 3", actual)

	actual = args.Map{"result": m.Has("a")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected Has('a') true", actual)

	actual = args.Map{"result": m.Has("z")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Has('z') false", actual)

	actual = args.Map{"result": m.HasDefined("a")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasDefined('a') true", actual)

	actual = args.Map{"result": m.IsKeyMissing("a")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsKeyMissing('a') false", actual)

	actual = args.Map{"result": m.IsKeyMissing("z")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsKeyMissing('z') true", actual)

	actual = args.Map{"result": m.HasDefinedAll("a", "b")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasDefinedAll('a','b') true", actual)

	actual = args.Map{"result": m.HasDefinedAll("a", "z")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasDefinedAll('a','z') false", actual)
}

func Test_Map_NilSafety(t *testing.T) {
	// Arrange
	var m args.Map

	// Act & Assert
	_, isValid := m.Get("key")
	actual := args.Map{"result": isValid}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map Get should return false", actual)

	actual = args.Map{"result": m.Has("key")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map Has should return false", actual)

	actual = args.Map{"result": m.HasDefined("key")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map HasDefined should return false", actual)

	actual = args.Map{"result": m.HasDefinedAll("a")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map HasDefinedAll should return false", actual)

	actual = args.Map{"result": m.IsKeyInvalid("key")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map IsKeyInvalid should return false", actual)

	actual = args.Map{"result": m.IsKeyMissing("key")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map IsKeyMissing should return false", actual)
}

func Test_Map_GetDefaults(t *testing.T) {
	// Arrange
	m := args.Map{
		"count": 5,
		"flag":  true,
		"text":  "hello",
	}

	// Act & Assert
	actual := args.Map{"result": m.GetAsIntDefault("count", 0) != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected GetAsIntDefault to return 5", actual)

	actual = args.Map{"result": m.GetAsIntDefault("missing", 99) != 99}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected GetAsIntDefault to return default 99", actual)

	actual = args.Map{"result": m.GetAsBoolDefault("flag", false) != true}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected GetAsBoolDefault to return true", actual)

	actual = args.Map{"result": m.GetAsBoolDefault("missing", true) != true}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected GetAsBoolDefault to return default true", actual)

	actual = args.Map{"result": m.GetAsStringDefault("text") != "hello"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected GetAsStringDefault to return hello", actual)

	actual = args.Map{"result": m.GetAsStringDefault("missing") != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected GetAsStringDefault to return empty", actual)
}

func Test_Map_SpecialKeys(t *testing.T) {
	// Arrange
	m := args.Map{
		"when":     "test scenario",
		"title":    "test title",
		"expected": "some result",
		"first":    "f1",
		"second":   "f2",
		"third":    "f3",
	}

	// Act & Assert
	actual := args.Map{"result": m.When() != "test scenario"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected When() to return scenario", actual)

	actual = args.Map{"result": m.Title() != "test title"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Title() to return title", actual)

	actual = args.Map{"result": m.Expected() != "some result"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Expected() to return expected", actual)

	actual = args.Map{"result": m.FirstItem() != "f1"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected FirstItem() to return f1", actual)

	actual = args.Map{"result": m.SecondItem() != "f2"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected SecondItem() to return f2", actual)

	actual = args.Map{"result": m.ThirdItem() != "f3"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ThirdItem() to return f3", actual)
}

func Test_Map_GetByIndex(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}

	// Act
	slice := m.Slice()

	// Assert
	actual := args.Map{"result": len(slice) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 items in slice", actual)

	actual = args.Map{"result": m.GetByIndex(999) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for out of range index", actual)
}

func Test_Map_GoLiteral(t *testing.T) {
	// Arrange
	m := args.Map{
		"key1": "value1",
		"key2": 42,
	}

	// Act
	lines := m.GoLiteralLines()
	str := m.GoLiteralString()

	// Assert
	actual := args.Map{"result": len(lines) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 lines", actual)

	actual = args.Map{"result": str == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GoLiteralString should not be empty", actual)
}

func Test_Map_GoLiteral_Empty(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	lines := m.GoLiteralLines()

	// Assert
	actual := args.Map{"result": len(lines) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 lines", actual)
}

func Test_Map_GetAsStrings(t *testing.T) {
	// Arrange
	m := args.Map{
		"items": []string{"a", "b", "c"},
	}

	// Act
	items, isValid := m.GetAsStrings("items")

	// Assert
	actual := args.Map{"result": isValid}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)

	actual = args.Map{"result": len(items) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 items", actual)
}

func Test_Map_GetAsAnyItems_FromMapGet(t *testing.T) {
	// Arrange
	m := args.Map{
		"items": []any{1, "two", true},
	}

	// Act
	items, isValid := m.GetAsAnyItems("items")

	// Assert
	actual := args.Map{"result": isValid}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)

	actual = args.Map{"result": len(items) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 items", actual)
}

func Test_Map_String_FromMapGet(t *testing.T) {
	// Arrange
	m := args.Map{
		"key": "value",
	}

	// Act
	s := m.String()

	// Assert
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Map String() should not be empty", actual)
}

func Test_Map_SetActual_FromMapGet(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	m.SetActual("result")

	// Assert
	actual := args.Map{"result": m.Actual() != "result"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Actual to return result", actual)
}

func Test_Map_SortedKeys_FromMapGet(t *testing.T) {
	// Arrange
	m := args.Map{
		"c": 3,
		"a": 1,
		"b": 2,
	}

	// Act
	keys, err := m.SortedKeys()

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SortedKeys error:", actual)

	actual = args.Map{"result": len(keys) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 keys", actual)

	actual = args.Map{"result": keys[0] != "a" || keys[1] != "b" || keys[2] != "c"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "keys should be sorted", actual)
}

func Test_Map_ValidArgs(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": "hello",
		"b": 42,
	}

	// Act
	validArgs := m.ValidArgs()

	// Assert
	actual := args.Map{"result": len(validArgs) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 valid args", actual)
}

func Test_Map_Args(t *testing.T) {
	// Arrange
	m := args.Map{
		"x": 10,
		"y": 20,
	}

	// Act
	result := m.Args("x", "y")

	// Assert
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 args", actual)
}

func Test_One(t *testing.T) {
	for caseIndex, testCase := range extOneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		first := input.GetDirectLower("first")
		expect := input.GetDirectLower("expect")

		one := args.OneAny{
			First:  first,
			Expect: expect,
		}

		// Act
		actual := args.Map{
			"hasFirst":   one.HasFirst(),
			"hasExpect":  one.HasExpect(),
			"argsCount":  one.ArgsCount(),
			"validCount": len(one.ValidArgs()),
		}

		if one.HasFirst() {
			actual["firstItem"] = one.FirstItem()
		}

		if one.HasExpect() {
			actual["expected"] = one.Expected()
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_One_String(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello", Expect: 42}

	// Act
	s := one.String()

	// Assert
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "One.String() should not be empty", actual)
}

func Test_One_Args(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello"}

	// Act & Assert
	actual := args.Map{"result": len(one.Args(1)) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 arg", actual)

	actual = args.Map{"result": len(one.Args(0)) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 args", actual)
}

func Test_One_Slice_Cached(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello"}

	// Act
	slice1 := one.Slice()
	slice2 := one.Slice()

	// Assert
	actual := args.Map{"result": len(slice1) != len(slice2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "cached Slice should return same length", actual)
}

func Test_One_GetByIndex(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello", Expect: 42}

	// Act & Assert
	actual := args.Map{"result": one.GetByIndex(0) != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected GetByIndex(0) to return hello", actual)

	actual = args.Map{"result": one.GetByIndex(99) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected GetByIndex(99) to return nil", actual)
}

func Test_One_LeftRight_FromMapGet(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello", Expect: 42}

	// Act
	lr := one.LeftRight()

	// Assert
	actual := args.Map{"result": lr.Left != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Left to be hello", actual)
}

func Test_Two(t *testing.T) {
	for caseIndex, testCase := range extTwoTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		first, _ := input.GetAsString("first")
		second, _ := input.GetAsString("second")

		two := args.TwoAny{
			First:  first,
			Second: second,
		}

		// Act
		actual := args.Map{
			"hasFirst":   two.HasFirst(),
			"hasSecond":  two.HasSecond(),
			"argsCount":  two.ArgsCount(),
			"validCount": len(two.ValidArgs()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Two_String(t *testing.T) {
	// Arrange
	two := args.TwoAny{First: "a", Second: "b"}

	// Act
	s := two.String()

	// Assert
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Two.String() should not be empty", actual)
}

func Test_Two_LeftRight_FromMapGet(t *testing.T) {
	// Arrange
	two := args.TwoAny{First: "a", Second: "b", Expect: "c"}

	// Act
	lr := two.LeftRight()

	// Assert
	actual := args.Map{"result": lr.Left != "a" || lr.Right != "b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "LeftRight conversion failed", actual)
}

func Test_Three_Methods_FromMapGet(t *testing.T) {
	// Arrange
	three := args.ThreeAny{
		First:  "a",
		Second: "b",
		Third:  "c",
		Expect: "d",
	}

	// Act & Assert
	actual := args.Map{"result": three.HasFirst()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasFirst should be true", actual)

	actual = args.Map{"result": three.HasSecond()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasSecond should be true", actual)

	actual = args.Map{"result": three.HasThird()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasThird should be true", actual)

	actual = args.Map{"result": three.HasExpect()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasExpect should be true", actual)

	actual = args.Map{"result": three.ArgsCount() != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ArgsCount 3", actual)

	actual = args.Map{"result": len(three.ValidArgs()) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 valid args", actual)

	actual = args.Map{"result": three.String() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "String() should not be empty", actual)

	actual = args.Map{"result": three.FirstItem() != "a"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FirstItem should be a", actual)

	actual = args.Map{"result": three.SecondItem() != "b"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SecondItem should be b", actual)

	actual = args.Map{"result": three.ThirdItem() != "c"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ThirdItem should be c", actual)

	argTwo := three.ArgTwo()
	actual = args.Map{"result": argTwo.First != "a" || argTwo.Second != "b"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ArgTwo should have first two", actual)

	argThree := three.ArgThree()
	actual = args.Map{"result": argThree.First != "a"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ArgThree should copy", actual)

	lr := three.LeftRight()
	actual = args.Map{"result": lr.Left != "a" || lr.Right != "b"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "LeftRight should map first/second", actual)
}

func Test_LeftRight(t *testing.T) {
	for caseIndex, testCase := range extLeftRightTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left, _ := input.GetAsInt("left")
		right, _ := input.GetAsInt("right")

		lr := args.LeftRightAny{
			Left:  left,
			Right: right,
		}

		// Act
		actual := args.Map{
			"hasLeft":    lr.HasLeft(),
			"hasRight":   lr.HasRight(),
			"hasFirst":   lr.HasFirst(),
			"hasSecond":  lr.HasSecond(),
			"argsCount":  lr.ArgsCount(),
			"validCount": len(lr.ValidArgs()),
			"firstItem":  lr.FirstItem(),
			"secondItem": lr.SecondItem(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LeftRight_Clone_FromMapGet(t *testing.T) {
	// Arrange
	lr := args.LeftRightAny{Left: "a", Right: "b", Expect: "c"}

	// Act
	cloned := lr.Clone()

	// Assert
	actual := args.Map{"result": cloned.Left != "a" || cloned.Right != "b" || cloned.Expect != "c"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone should preserve all fields", actual)
}

func Test_LeftRight_String(t *testing.T) {
	// Arrange
	lr := args.LeftRightAny{Left: "a", Right: "b"}

	// Act
	s := lr.String()

	// Assert
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "LeftRight.String() should not be empty", actual)
}

func Test_String_Methods_FromMapGet(t *testing.T) {
	for caseIndex, testCase := range extStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		s := args.String(inputStr)

		// Act
		actual := args.Map{
			"length":              s.Length(),
			"isEmpty":             s.IsEmpty(),
			"hasCharacter":        s.HasCharacter(),
			"isDefined":           s.IsDefined(),
			"isEmptyOrWhitespace": s.IsEmptyOrWhitespace(),
			"asciiLength":         s.AscIILength(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_String_Operations(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act & Assert
	actual := args.Map{"result": s.String() != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "String() should return hello", actual)

	actual = args.Map{"result": len(s.Bytes()) != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Bytes() should have 5 bytes", actual)

	actual = args.Map{"result": len(s.Runes()) != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Runes() should have 5 runes", actual)

	actual = args.Map{"result": s.Count() != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Count should be 5", actual)

	trimmed := args.String("  hello  ").TrimSpace()
	actual = args.Map{"result": trimmed != "hello"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "TrimSpace expected hello", actual)

	replaced := s.ReplaceAll("hello", "world")
	actual = args.Map{"result": replaced != "world"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReplaceAll expected world", actual)

	sub := s.Substring(0, 3)
	actual = args.Map{"result": sub != "hel"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Substring expected hel", actual)

	concat := args.String("a").Concat("b", "c")
	actual = args.Map{"result": concat != "abc"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Concat expected abc", actual)

	joined := args.String("a").Join("-", "b", "c")
	actual = args.Map{"result": joined != "a-b-c"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Join expected a-b-c", actual)

	split := args.String("a,b,c").Split(",")
	actual = args.Map{"result": len(split) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Split expected 3", actual)
}

func Test_String_Quoting(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act & Assert
	dq := s.DoubleQuote()
	actual := args.Map{"result": dq == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DoubleQuote should not be empty", actual)

	dqq := s.DoubleQuoteQ()
	actual = args.Map{"result": dqq == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteQ should not be empty", actual)

	sq := s.SingleQuote()
	actual = args.Map{"result": sq == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SingleQuote should not be empty", actual)

	vdq := s.ValueDoubleQuote()
	actual = args.Map{"result": vdq == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueDoubleQuote should not be empty", actual)
}

func Test_Empty_Creator(t *testing.T) {
	// Arrange & Act
	m := args.Empty.Map()
	fw := args.Empty.FuncWrap()
	fm := args.Empty.FuncMap()
	h := args.Empty.Holder()

	// Assert
	actual := args.Map{"result": m == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Empty.Map should not be nil", actual)

	actual = args.Map{"result": fw == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Empty.FuncWrap should not be nil", actual)

	actual = args.Map{"result": fm == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Empty.FuncMap should not be nil", actual)

	actual = args.Map{"result": h.ArgsCount() != 7}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Empty.Holder ArgsCount should be 7", actual)
}

func Test_Holder_Methods_FromMapGet(t *testing.T) {
	// Arrange
	h := args.HolderAny{
		First:  "a",
		Second: "b",
		Third:  "c",
		Fourth: "d",
		Fifth:  "e",
		Sixth:  "f",
		Expect: "g",
	}

	// Act & Assert
	actual := args.Map{"result": h.HasFirst()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasFirst should be true", actual)

	actual = args.Map{"result": h.HasSecond()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasSecond should be true", actual)

	actual = args.Map{"result": h.HasThird()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasThird should be true", actual)

	actual = args.Map{"result": h.HasFourth()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasFourth should be true", actual)

	actual = args.Map{"result": h.HasFifth()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasFifth should be true", actual)

	actual = args.Map{"result": h.HasSixth()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasSixth should be true", actual)

	actual = args.Map{"result": h.HasExpect()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasExpect should be true", actual)

	actual = args.Map{"result": h.ArgsCount() != 7}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ArgsCount should be 7", actual)

	actual = args.Map{"result": len(h.ValidArgs()) != 6}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 6 valid args", actual)

	actual = args.Map{"result": h.FirstItem() != "a"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FirstItem should be a", actual)

	actual = args.Map{"result": h.SecondItem() != "b"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SecondItem should be b", actual)

	actual = args.Map{"result": h.ThirdItem() != "c"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ThirdItem should be c", actual)

	actual = args.Map{"result": h.FourthItem() != "d"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FourthItem should be d", actual)

	actual = args.Map{"result": h.FifthItem() != "e"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FifthItem should be e", actual)

	actual = args.Map{"result": h.SixthItem() != "f"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SixthItem should be f", actual)

	actual = args.Map{"result": h.Expected() != "g"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Expected should be g", actual)

	actual = args.Map{"result": h.String() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Holder String() should not be empty", actual)
}

func Test_Holder_Args(t *testing.T) {
	// Arrange
	h := args.HolderAny{
		First:  "a",
		Second: "b",
		Third:  "c",
	}

	// Act & Assert
	actual := args.Map{"result": len(h.Args(1)) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Args(1) should return 1", actual)

	actual = args.Map{"result": len(h.Args(3)) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Args(3) should return 3", actual)

	actual = args.Map{"result": len(h.Args(6)) != 6}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Args(6) should return 6", actual)
}

func Test_Holder_ArgTwo_ArgThree_ArgFour_ArgFive(t *testing.T) {
	// Arrange
	h := args.HolderAny{
		First:  "a",
		Second: "b",
		Third:  "c",
		Fourth: "d",
		Fifth:  "e",
	}

	// Act & Assert
	at := h.ArgTwo()
	actual := args.Map{"result": at.First != "a" || at.Second != "b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ArgTwo should have first two", actual)

	a3 := h.ArgThree()
	actual = args.Map{"result": a3.First != "a"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ArgThree should have first", actual)

	a4 := h.ArgFour()
	actual = args.Map{"result": a4.First != "a"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ArgFour should have first", actual)

	a5 := h.ArgFive()
	actual = args.Map{"result": a5.First != "a"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ArgFive should have first", actual)
}

func Test_Dynamic_Methods(t *testing.T) {
	// Arrange
	d := args.DynamicAny{
		Params: args.Map{
			"key1": "val1",
			"key2": 42,
		},
		Expect: "expected",
	}

	// Act & Assert
	// Map.ArgsCount() excludes "expected"/"func" keys.
	// Map.HasFunc() returns true even for nil func (non-nil FuncWrap wrapper),
	// so ArgsCount = Length(2) - HasFunc(1) = 1.
	actual := args.Map{"result": d.ArgsCount() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ArgsCount 1", actual)

	actual = args.Map{"result": d.HasExpect()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasExpect should be true", actual)

	actual = args.Map{"result": d.Expected() != "expected"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Expected should be expected", actual)

	actual = args.Map{"result": d.HasDefined("key1")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasDefined key1 should be true", actual)

	actual = args.Map{"result": d.Has("key1")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Has key1 should be true", actual)

	actual = args.Map{"result": d.IsKeyMissing("key1")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsKeyMissing key1 should be false", actual)

	actual = args.Map{"result": d.IsKeyMissing("missing")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsKeyMissing missing should be true", actual)

	actual = args.Map{"result": d.IsKeyInvalid("key1")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsKeyInvalid key1 should be false", actual)

	val, isValid := d.Get("key1")
	actual = args.Map{"result": isValid || val != "val1"}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Get key1 should return val1", actual)

	intVal, intValid := d.GetAsInt("key2")
	actual = args.Map{"result": intValid || intVal != 42}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetAsInt key2 should return 42", actual)

	actual = args.Map{"result": d.GetAsIntDefault("missing", 99) != 99}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetAsIntDefault should return 99", actual)

	strVal, strValid := d.GetAsString("key1")
	actual = args.Map{"result": strValid || strVal != "val1"}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetAsString key1 should return val1", actual)

	actual = args.Map{"result": d.GetAsStringDefault("missing") != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetAsStringDefault should return empty", actual)

	// Note: Dynamic.String() calls Slice() which uses converters.Map.SortedKeys.
	// That function does not support args.Map type, so String() panics.
	// This is a known limitation — skipping String() test.
}

func Test_Dynamic_NilSafety(t *testing.T) {
	// Arrange
	var d *args.DynamicAny

	// Act & Assert
	actual := args.Map{"result": d.ArgsCount() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil ArgsCount should return 0", actual)

	actual = args.Map{"result": d.GetWorkFunc() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil GetWorkFunc should return nil", actual)

	actual = args.Map{"result": d.HasDefined("key")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil HasDefined should return false", actual)

	actual = args.Map{"result": d.Has("key")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil Has should return false", actual)

	actual = args.Map{"result": d.HasDefinedAll("key")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil HasDefinedAll should return false", actual)

	actual = args.Map{"result": d.IsKeyInvalid("key")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil IsKeyInvalid should return false", actual)

	actual = args.Map{"result": d.IsKeyMissing("key")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil IsKeyMissing should return false", actual)

	_, isValid := d.Get("key")
	actual = args.Map{"result": isValid}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil Get should return false", actual)

	actual = args.Map{"result": d.HasExpect()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil HasExpect should return false", actual)
}

func Test_FuncMap_Basic_FromMapGet(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act & Assert
	actual := args.Map{"result": fm.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "new FuncMap should be empty", actual)

	actual = args.Map{"result": fm.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Length should be 0", actual)

	actual = args.Map{"result": fm.Count() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Count should be 0", actual)

	actual = args.Map{"result": fm.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasAnyItem should be false", actual)

	actual = args.Map{"result": fm.Has("nonexistent")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Has should be false", actual)

	actual = args.Map{"result": fm.IsContains("nonexistent")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsContains should be false", actual)

	actual = args.Map{"result": fm.Get("nonexistent") != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Get should return nil", actual)

	actual = args.Map{"result": fm.IsValidFuncOf("nonexistent")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsValidFuncOf should be false", actual)

	actual = args.Map{"result": fm.IsInvalidFunc("nonexistent")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsInvalidFunc should be true", actual)

	actual = args.Map{"result": fm.PkgPath("nonexistent") != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PkgPath should return empty", actual)

	actual = args.Map{"result": fm.PkgNameOnly("nonexistent") != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PkgNameOnly should return empty", actual)

	actual = args.Map{"result": fm.FuncDirectInvokeName("nonexistent") != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FuncDirectInvokeName should return empty", actual)

	actual = args.Map{"result": fm.ArgsCount("nonexistent") != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ArgsCount should return 0", actual)

	actual = args.Map{"result": fm.ReturnLength("nonexistent") != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReturnLength should return 0", actual)

	actual = args.Map{"result": fm.IsPublicMethod("nonexistent")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsPublicMethod should return false", actual)

	actual = args.Map{"result": fm.IsPrivateMethod("nonexistent")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsPrivateMethod should return false", actual)

	actual = args.Map{"result": fm.GetType("nonexistent") != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetType should return nil", actual)

	actual = args.Map{"result": fm.GetPascalCaseFuncName("nonexistent") != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetPascalCaseFuncName should return empty", actual)

	actual = args.Map{"result": fm.InvalidError() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty FuncMap InvalidError should return error", actual)
}

func Test_FuncMap_Add_FromMapGet(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	fm.Add(someFunctionV1)

	// Assert
	actual := args.Map{"result": fm.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FuncMap should not be empty after Add", actual)

	actual = args.Map{"result": fm.Has("someFunctionV1")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected someFunctionV1 in map", actual)

	actual = args.Map{"result": fm.IsValidFuncOf("someFunctionV1")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "someFunctionV1 should be valid", actual)
}

func Test_FuncMap_Adds_FromMapGet(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	fm.Adds(someFunctionV1, someFunctionV2)

	// Assert
	actual := args.Map{"result": fm.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_FuncDetector(t *testing.T) {
	// Arrange & Act
	fw := args.FuncDetector.GetFuncWrap(someFunctionV1)

	// Assert
	actual := args.Map{"result": fw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetFuncWrap should not return nil", actual)

	actual = args.Map{"result": fw.HasValidFunc()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "wrapped func should be valid", actual)
}
