package bytetypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/bytetype"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_GetSet_Verification(t *testing.T) {
	for caseIndex, testCase := range getSetTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		condition, _ := input.GetAsBool("condition")
		trueVal, _ := input.GetAsInt("trueValue")
		falseVal, _ := input.GetAsInt("falseValue")

		// Act
		result := bytetype.GetSet(
			condition,
			bytetype.New(byte(trueVal)),
			bytetype.New(byte(falseVal)),
		)

		actual := args.Map{
			"result": result.ValueInt(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_GetSetVariant_Verification(t *testing.T) {
	for caseIndex, testCase := range getSetVariantTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		condition, _ := input.GetAsBool("condition")
		trueVal, _ := input.GetAsInt("trueValue")
		falseVal, _ := input.GetAsInt("falseValue")

		// Act
		result := bytetype.GetSetVariant(
			condition,
			byte(trueVal),
			byte(falseVal),
		)

		actual := args.Map{
			"result": result.ValueInt(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_Comparisons(t *testing.T) {
	for caseIndex, testCase := range comparisonTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		v := bytetype.New(byte(val))

		// Act
		actual := args.Map{
			"isEqual3":        v.IsEqual(3),
			"isEqual5":        v.IsEqual(5),
			"isGreater3":      v.IsGreater(3),
			"isGreater7":      v.IsGreater(7),
			"isGreaterEqual5": v.IsGreaterEqual(5),
			"isLess3":         v.IsLess(3),
			"isLess7":         v.IsLess(7),
			"isLessEqual5":    v.IsLessEqual(5),
			"isBetween3and7":  v.IsBetween(3, 7),
			"isBetween6and8":  v.IsBetween(6, 8),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_String_Conversion(t *testing.T) {
	for caseIndex, testCase := range stringConversionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := bytetype.String([]byte(inputStr))

		actual := args.Map{
			"result": result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_Methods(t *testing.T) {
	for caseIndex, testCase := range variantMethodsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		v := bytetype.New(byte(val))

		// Act
		actual := args.Map{
			"isZero":    v.IsZero(),
			"isOne":     v.IsOne(),
			"isTwo":     v.IsTwo(),
			"isThree":  v.IsThree(),
			"isMin":     v.IsMin(),
			"isValid":   v.IsValid(),
			"isInvalid": v.IsInvalid(),
			"valueInt":  v.ValueInt(),
			"valueByte": int(v.ValueByte()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_Arithmetic(t *testing.T) {
	for caseIndex, testCase := range variantArithmeticTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		base, _ := input.GetAsInt("base")
		n, _ := input.GetAsInt("n")
		v := bytetype.New(byte(base))

		// Act
		actual := args.Map{
			"addResult":      v.Add(byte(n)).ValueInt(),
			"subtractResult": v.Subtract(byte(n)).ValueInt(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_IntComparisons(t *testing.T) {
	// Arrange
	v := bytetype.New(5)

	// Act & Assert
	actual := args.Map{"result": v.IsEqualInt(5)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsEqualInt(5) true", actual)
	actual = args.Map{"result": v.IsEqualInt(3)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsEqualInt(3) false", actual)
	actual = args.Map{"result": v.IsGreaterInt(3)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsGreaterInt(3) true", actual)
	actual = args.Map{"result": v.IsGreaterEqualInt(5)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsGreaterEqualInt(5) true", actual)
	actual = args.Map{"result": v.IsLessInt(7)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsLessInt(7) true", actual)
	actual = args.Map{"result": v.IsLessEqualInt(5)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsLessEqualInt(5) true", actual)
	actual = args.Map{"result": v.IsBetweenInt(3, 7)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsBetweenInt(3,7) true", actual)
}

func Test_Variant_Is(t *testing.T) {
	// Arrange
	v := bytetype.New(3)

	// Act & Assert
	actual := args.Map{"result": v.Is(bytetype.Three)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected Is(Three) true", actual)
	actual = args.Map{"result": v.Is(bytetype.One)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Is(One) false", actual)
}

func Test_Variant_HasIndexInStrings(t *testing.T) {
	// Arrange
	v := bytetype.New(1)
	items := []string{"zero", "one", "two"}

	// Act
	val, isValid := v.HasIndexInStrings(items...)

	// Assert
	actual := args.Map{"result": isValid}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected isValid true", actual)
	actual = args.Map{"result": val != "one"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'one', got ''", actual)

	// Out of range
	v2 := bytetype.New(10)
	_, isValid2 := v2.HasIndexInStrings(items...)
	actual = args.Map{"result": isValid2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected isValid false for out of range", actual)

	// Empty slice
	_, isValid3 := v.HasIndexInStrings()
	actual = args.Map{"result": isValid3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected isValid false for empty slice", actual)
}

func Test_Variant_ValueConversions(t *testing.T) {
	// Arrange
	v := bytetype.New(42)

	// Act & Assert
	actual := args.Map{"result": v.ValueUInt16() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ValueUInt16=42", actual)
	actual = args.Map{"result": v.ValueInt8() != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ValueInt8=42", actual)
	actual = args.Map{"result": v.ValueInt16() != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ValueInt16=42", actual)
	actual = args.Map{"result": v.ValueInt32() != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ValueInt32=42", actual)
	actual = args.Map{"result": v.ValueString() != "42"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ValueString='42', got ''", actual)
	actual = args.Map{"result": v.ToNumberString() != "42"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ToNumberString='42', got ''", actual)
	actual = args.Map{"result": v.StringValue() != "42"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected StringValue='42', got ''", actual)
}

func Test_Variant_IsValueEqual(t *testing.T) {
	// Arrange
	v := bytetype.New(5)

	// Act & Assert
	actual := args.Map{"result": v.IsValueEqual(5)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsValueEqual(5) true", actual)
	actual = args.Map{"result": v.IsValueEqual(3)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsValueEqual(3) false", actual)
}

func Test_Variant_ToPtr(t *testing.T) {
	// Arrange
	v := bytetype.New(7)

	// Act
	ptr := v.ToPtr()

	// Assert
	actual := args.Map{"result": ptr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ToPtr to not be nil", actual)
	actual = args.Map{"result": ptr.ValueInt() != 7}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ptr value 7", actual)
}

func Test_Variant_IsNameEqual(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	name := v.Name()

	// Assert
	actual := args.Map{"result": v.IsNameEqual(name)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsNameEqual('') true", actual)
	actual = args.Map{"result": v.IsNameEqual("NonExistent")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsNameEqual('NonExistent') false", actual)
}

func Test_Variant_IsAnyNamesOf(t *testing.T) {
	// Arrange
	v := bytetype.One
	name := v.Name()

	// Act & Assert
	actual := args.Map{"result": v.IsAnyNamesOf(name, "Other")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsAnyNamesOf to find match", actual)
	actual = args.Map{"result": v.IsAnyNamesOf("None", "Other")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsAnyNamesOf to not match", actual)
}
