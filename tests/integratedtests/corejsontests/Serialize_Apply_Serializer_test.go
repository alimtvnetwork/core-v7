package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// serializerLogic — Apply
// =============================================================================

func Test_Cov41_Serialize_Apply_Verification(t *testing.T) {
	for caseIndex, tc := range serializeApplyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")

		// Act
		r := corejson.Serialize.Apply(inputVal)
		actual := args.Map{
			"hasError": r.HasError(),
			"hasBytes": r.Length() > 0,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Cov41_Serialize_Apply_Unmarshalable(t *testing.T) {
	tc := serializeApplyUnmarshalableTestCase

	// Arrange
	// (channel type is unmarshalable)

	// Act
	r := corejson.Serialize.Apply(make(chan int))
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// serializerLogic — From* methods
// =============================================================================

func Test_Cov41_Serialize_FromBytes(t *testing.T) {
	tc := serializeFromBytesTestCase

	// Arrange
	// (direct byte input)

	// Act
	r := corejson.Serialize.FromBytes([]byte("hello"))
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_FromStrings(t *testing.T) {
	tc := serializeFromStringsTestCase

	// Arrange
	// (direct string slice)

	// Act
	r := corejson.Serialize.FromStrings([]string{"a"})
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_FromStringsSpread(t *testing.T) {
	tc := serializeFromStringsSpreadTestCase

	// Arrange
	// (variadic strings)

	// Act
	r := corejson.Serialize.FromStringsSpread("a", "b")
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_FromString(t *testing.T) {
	tc := serializeFromStringTestCase

	// Arrange
	// (direct string)

	// Act
	r := corejson.Serialize.FromString("hello")
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_FromInteger(t *testing.T) {
	tc := serializeFromIntegerTestCase

	// Arrange
	// (int value)

	// Act
	r := corejson.Serialize.FromInteger(42)
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_FromInteger64(t *testing.T) {
	tc := serializeFromInteger64TestCase

	// Arrange
	// (int64 value)

	// Act
	r := corejson.Serialize.FromInteger64(99)
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_FromBool(t *testing.T) {
	tc := serializeFromBoolTestCase

	// Arrange
	// (bool true)

	// Act
	r := corejson.Serialize.FromBool(true)
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_FromIntegers(t *testing.T) {
	tc := serializeFromIntegersTestCase

	// Arrange
	// (int slice)

	// Act
	r := corejson.Serialize.FromIntegers([]int{1, 2})
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// serializerLogic — UsingAnyPtr / UsingAny
// =============================================================================

func Test_Cov41_Serialize_UsingAnyPtr_Valid(t *testing.T) {
	tc := serializeUsingAnyPtrValidTestCase

	// Arrange
	// (valid string)

	// Act
	r := corejson.Serialize.UsingAnyPtr("hello")
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_UsingAnyPtr_Unmarshalable(t *testing.T) {
	tc := serializeUsingAnyPtrUnmarshalableTestCase

	// Arrange
	// (channel is unmarshalable)

	// Act
	r := corejson.Serialize.UsingAnyPtr(make(chan int))
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_UsingAny(t *testing.T) {
	tc := serializeUsingAnyTestCase

	// Arrange
	// (valid string)

	// Act
	r := corejson.Serialize.UsingAny("hello")
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// serializerLogic — Raw / Marshal / ApplyMust / ToBytesMust / etc.
// =============================================================================

func Test_Cov41_Serialize_Raw(t *testing.T) {
	tc := serializeRawTestCase

	// Arrange
	// (valid string)

	// Act
	b, err := corejson.Serialize.Raw("hello")
	actual := args.Map{
		"hasError": err != nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_Marshal(t *testing.T) {
	tc := serializeMarshalTestCase

	// Arrange
	// (valid string)

	// Act
	b, err := corejson.Serialize.Marshal("hello")
	actual := args.Map{
		"hasError": err != nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_ApplyMust(t *testing.T) {
	tc := serializeApplyMustTestCase

	// Arrange
	// (valid string)

	// Act
	r := corejson.Serialize.ApplyMust("hello")
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_ToBytesMust(t *testing.T) {
	tc := serializeToBytesMustTestCase

	// Arrange
	// (valid string)

	// Act
	b := corejson.Serialize.ToBytesMust("hello")
	actual := args.Map{
		"hasBytes": len(b) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_ToSafeBytesMust(t *testing.T) {
	tc := serializeToSafeBytesMustTestCase

	// Arrange
	// (valid string)

	// Act
	b := corejson.Serialize.ToSafeBytesMust("hello")
	actual := args.Map{
		"hasBytes": len(b) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// serializerLogic — Swallow / ToString / ToPretty
// =============================================================================

func Test_Cov41_Serialize_ToSafeBytesSwallowErr(t *testing.T) {
	tc := serializeToSafeBytesSwallowErrTestCase

	// Arrange
	// (valid string)

	// Act
	b := corejson.Serialize.ToSafeBytesSwallowErr("hello")
	actual := args.Map{
		"hasBytes": len(b) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_ToBytesSwallowErr(t *testing.T) {
	tc := serializeToBytesSwallowErrTestCase

	// Arrange
	// (valid string)

	// Act
	b := corejson.Serialize.ToBytesSwallowErr("hello")
	actual := args.Map{
		"hasBytes": len(b) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_ToBytesErr(t *testing.T) {
	tc := serializeToBytesErrTestCase

	// Arrange
	// (valid string)

	// Act
	b, err := corejson.Serialize.ToBytesErr("hello")
	actual := args.Map{
		"hasError": err != nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_ToString(t *testing.T) {
	tc := serializeToStringTestCase

	// Arrange
	// (valid input)

	// Act
	s := corejson.Serialize.ToString("hello")
	actual := args.Map{
		"hasContent": len(s) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_ToStringMust(t *testing.T) {
	tc := serializeToStringMustTestCase

	// Arrange
	// (valid input)

	// Act
	s := corejson.Serialize.ToStringMust("hello")
	actual := args.Map{
		"hasContent": len(s) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_ToStringErr(t *testing.T) {
	tc := serializeToStringErrTestCase

	// Arrange
	// (valid input)

	// Act
	s, err := corejson.Serialize.ToStringErr("hello")
	actual := args.Map{
		"hasError":   err != nil,
		"hasContent": len(s) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_ToPrettyStringErr(t *testing.T) {
	tc := serializeToPrettyStringErrTestCase

	// Arrange
	// (map input for pretty output)

	// Act
	s, err := corejson.Serialize.ToPrettyStringErr(map[string]int{"a": 1})
	actual := args.Map{
		"hasError":   err != nil,
		"hasContent": len(s) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_ToPrettyStringIncludingErr(t *testing.T) {
	tc := serializeToPrettyStringIncludingErrTestCase

	// Arrange
	// (valid input)

	// Act
	s := corejson.Serialize.ToPrettyStringIncludingErr("hello")
	actual := args.Map{
		"hasContent": len(s) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov41_Serialize_Pretty(t *testing.T) {
	tc := serializePrettyTestCase

	// Arrange
	// (valid input)

	// Act
	s := corejson.Serialize.Pretty("hello")
	actual := args.Map{
		"hasContent": len(s) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
