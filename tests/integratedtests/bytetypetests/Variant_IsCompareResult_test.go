package bytetypetests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/bytetype"
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Variant_IsCompareResult(t *testing.T) {
	compareNameToEnum := map[string]corecomparator.Compare{
		"Equal":             corecomparator.Equal,
		"LeftGreater":       corecomparator.LeftGreater,
		"LeftGreaterEqual":  corecomparator.LeftGreaterEqual,
		"LeftLess":          corecomparator.LeftLess,
		"LeftLessEqual":     corecomparator.LeftLessEqual,
		"NotEqual":          corecomparator.NotEqual,
	}

	for caseIndex, testCase := range extIsCompareResultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		n, _ := input.GetAsInt("n")
		compareName, _ := input.GetAsString("compare")
		v := bytetype.New(byte(val))
		compareEnum := compareNameToEnum[compareName]

		// Act
		result := v.IsCompareResult(byte(n), compareEnum)

		actual := args.Map{
			"result": result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_IsCompareResult_Panic(t *testing.T) {
	// Arrange
	v := bytetype.New(5)

	// Act & Assert
	defer func() {
		r := recover()
		actual := args.Map{"result": r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic for out-of-range comparator", actual)
	}()

	v.IsCompareResult(3, corecomparator.Compare(99))
}

func Test_Variant_EnumMethods(t *testing.T) {
	for caseIndex, testCase := range extEnumMethodsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		v := bytetype.New(byte(val))

		// Act
		actual := args.Map{
			"name":           v.Name(),
			"nameValue":      v.NameValue(),
			"typeName":       v.TypeName(),
			"isValidRange":   v.IsValidRange(),
			"isInvalidRange": v.IsInvalidRange(),
			"stringValue":    v.StringValue(),
			"rangeNamesCsv":  v.RangeNamesCsv(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_IsMax(t *testing.T) {
	for caseIndex, testCase := range extIsMaxTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		v := bytetype.New(byte(val))

		// Act
		actual := args.Map{
			"isMax": v.IsMax(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_IsEnumEqual(t *testing.T) {
	for caseIndex, testCase := range extIsEnumEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		other, _ := input.GetAsInt("other")
		v := bytetype.New(byte(val))
		otherVariant := bytetype.New(byte(other))

		// Act
		// IsEnumEqual takes enuminf.BasicEnumer which requires pointer receiver
		// (UnmarshalJSON has pointer receiver), so pass &otherVariant.
		actual := args.Map{
			"isEnumEqual": v.IsEnumEqual(&otherVariant),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_IsAnyEnumsEqual(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act & Assert
	// IsAnyEnumsEqual takes ...enuminf.BasicEnumer which requires pointer receiver.
	// bytetype.Variant has UnmarshalJSON on pointer receiver, so use variables + &.
	two := bytetype.Two
	one := bytetype.One
	three := bytetype.Three

	actual := args.Map{"result": v.IsAnyEnumsEqual(&two, &one)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsAnyEnumsEqual to find match", actual)

	actual = args.Map{"result": v.IsAnyEnumsEqual(&two, &three)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsAnyEnumsEqual to not match", actual)

	actual = args.Map{"result": v.IsAnyEnumsEqual()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsAnyEnumsEqual with no args to return false", actual)
}

func Test_Variant_MarshalJSON(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	data, err := json.Marshal(v)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MarshalJSON error:", actual)

	actual = args.Map{"result": len(data) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MarshalJSON returned empty bytes", actual)
}

func Test_Variant_UnmarshalJSON(t *testing.T) {
	// Arrange
	v := bytetype.One
	data, _ := json.Marshal(v.Value())

	// Act
	var result bytetype.Variant
	err := json.Unmarshal(data, &result)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON error:", actual)

	actual = args.Map{"result": result.Value() != v.Value()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected value", actual)
}

func Test_Variant_UnmarshallToValue(t *testing.T) {
	// Arrange
	v := bytetype.Two
	data, _ := json.Marshal(v.Value())

	// Act
	val, err := v.UnmarshallToValue(data)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue error:", actual)

	actual = args.Map{"result": val != v.Value()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected", actual)
}

func Test_Variant_AllNameValues(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	names := v.AllNameValues()

	// Assert
	actual := args.Map{"result": len(names) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllNameValues should not be empty", actual)
}

func Test_Variant_IntegerEnumRanges(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	ranges := v.IntegerEnumRanges()

	// Assert
	actual := args.Map{"result": len(ranges) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IntegerEnumRanges should not be empty", actual)
}

func Test_Variant_MinMaxAny(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	min, max := v.MinMaxAny()

	// Assert
	actual := args.Map{"result": min == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "min should not be nil", actual)

	actual = args.Map{"result": max == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "max should not be nil", actual)
}

func Test_Variant_MinMaxStrings(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act & Assert
	actual := args.Map{"result": v.MinValueString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinValueString should not be empty", actual)

	actual = args.Map{"result": v.MaxValueString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaxValueString should not be empty", actual)

	actual = args.Map{"result": v.MinInt() < 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinInt should be >= 0", actual)

	actual = args.Map{"result": v.MaxInt() <= 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaxInt should be > 0", actual)
}

func Test_Variant_RangesDynamicMap(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	m := v.RangesDynamicMap()

	// Assert
	actual := args.Map{"result": len(m) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangesDynamicMap should not be empty", actual)
}

func Test_Variant_Format(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	formatted := v.Format("{name}")

	// Assert
	actual := args.Map{"result": formatted == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Format should return non-empty", actual)
}

func Test_Variant_StringRanges(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act & Assert
	actual := args.Map{"result": len(v.StringRanges()) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringRanges should not be empty", actual)

	actual = args.Map{"result": len(v.StringRangesPtr()) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringRangesPtr should not be empty", actual)
}

func Test_Variant_RangesInvalid(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act & Assert
	actual := args.Map{"result": v.RangesInvalidMessage() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangesInvalidMessage should not be empty", actual)

	actual = args.Map{"result": v.RangesInvalidErr() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangesInvalidErr should not be nil", actual)
}

func Test_Variant_EnumType(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	enumType := v.EnumType()

	// Assert
	actual := args.Map{"result": enumType == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "EnumType should not be nil", actual)
}

func Test_Variant_AsBasicEnumContractsBinder(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	binder := v.AsBasicEnumContractsBinder()

	// Assert
	actual := args.Map{"result": binder == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AsBasicEnumContractsBinder should not be nil", actual)
}

func Test_Variant_JsonString(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	js := v.JsonString()

	// Assert
	actual := args.Map{"result": js == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonString should not be empty", actual)
}

func Test_Variant_OnlySupportedErr(t *testing.T) {
	// Arrange
	v := bytetype.One
	// OnlySupportedErr compares against StringRanges() (plain names like "Zero"),
	// not AllNameValues() which returns "Name(Value)" format like "Zero(0)".
	allNames := v.StringRanges()

	// Act
	err := v.OnlySupportedErr(allNames...)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "all names supported should not error, got:", actual)
}

func Test_Variant_OnlySupportedMsgErr(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	err := v.OnlySupportedMsgErr("test message", "NonExistent")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "OnlySupportedMsgErr with single unsupported name should return error", actual)
}
