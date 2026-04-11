package coreoncetests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// AnyOnce — uncovered paths
// ==========================================================================

func Test_AnyOnce_ValueStringOnly(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOnce(func() any { return "hello" })

	// Act
	actual := args.Map{
		"valueStringOnly": ao.ValueStringOnly(),
		"safeString":      ao.SafeString(),
		"valueStringMust": ao.ValueStringMust(),
		"valueOnly":       ao.ValueOnly(),
	}

	// Assert
	expected := args.Map{
		"valueStringOnly": ao.ValueString(),
		"safeString":      ao.ValueString(),
		"valueStringMust": ao.ValueString(),
		"valueOnly":       "hello",
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce alias methods return expected -- hello", actual)
}

func Test_AnyOnce_ValueString_Cached(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOnce(func() any { return 42 })
	// Call twice to test caching path
	first := ao.ValueString()
	second := ao.ValueString()

	// Act
	actual := args.Map{"same": first == second}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "AnyOnce ValueString cached returns same -- second call", actual)
}

func Test_AnyOnce_CastFail(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOnce(func() any { return 42 })
	_, okStr := ao.CastValueString()
	_, okStrings := ao.CastValueStrings()
	_, okMap := ao.CastValueHashmapMap()
	_, okMapAny := ao.CastValueMapStringAnyMap()
	_, okBytes := ao.CastValueBytes()

	// Act
	actual := args.Map{
		"okStr": okStr, "okStrings": okStrings,
		"okMap": okMap, "okMapAny": okMapAny, "okBytes": okBytes,
	}

	// Assert
	expected := args.Map{
		"okStr": false, "okStrings": false,
		"okMap": false, "okMapAny": false, "okBytes": false,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce Cast methods fail -- wrong type", actual)
}

func Test_AnyOnce_ValueString_NilReturn(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOnce(func() any { return nil })
	result := ao.ValueString()

	// Act
	actual := args.Map{"isAngelBracket": result != ""}

	// Assert
	expected := args.Map{"isAngelBracket": true}
	expected.ShouldBeEqual(t, 0, "AnyOnce ValueString nil returns angel bracket -- nil", actual)
}

// ==========================================================================
// ErrorOnce — with actual error
// ==========================================================================

func Test_ErrorOnce_WithError(t *testing.T) {
	// Arrange
	testErr := errors.New("test error")
	eo := coreonce.NewErrorOnce(func() error { return testErr })

	// Act
	actual := args.Map{
		"hasError": eo.HasError(), "isEmptyError": eo.IsEmptyError(),
		"isEmpty": eo.IsEmpty(), "hasAnyItem": eo.HasAnyItem(),
		"isDefined": eo.IsDefined(), "isInvalid": eo.IsInvalid(),
		"isValid": eo.IsValid(), "isSuccess": eo.IsSuccess(),
		"isFailed": eo.IsFailed(), "isNull": eo.IsNull(),
		"isNullOrEmpty": eo.IsNullOrEmpty(),
		"message": eo.Message(),
		"isMessageEqual": eo.IsMessageEqual("test error"),
		"stringNotEmpty": eo.String() != "",
		"concatNew": eo.ConcatNew("extra") != nil,
		"concatNewStr": eo.ConcatNewString("extra") != "",
	}

	// Assert
	expected := args.Map{
		"hasError": true, "isEmptyError": false,
		"isEmpty": false, "hasAnyItem": true,
		"isDefined": true, "isInvalid": true,
		"isValid": false, "isSuccess": false,
		"isFailed": true, "isNull": false,
		"isNullOrEmpty": false,
		"message": "test error",
		"isMessageEqual": true,
		"stringNotEmpty": true,
		"concatNew": true, "concatNewStr": true,
	}
	expected.ShouldBeEqual(t, 0, "ErrorOnce with error returns expected -- test error", actual)
}

func Test_ErrorOnce_Value(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOnce(func() error { return nil })

	// Act
	actual := args.Map{
		"valueNil": eo.Value() == nil,
		"execute":  eo.Execute() == nil,
	}

	// Assert
	expected := args.Map{
		"valueNil": true,
		"execute": true,
	}
	expected.ShouldBeEqual(t, 0, "ErrorOnce Value/Execute return nil -- no error", actual)
}

// ==========================================================================
// AnyErrorOnce — with error
// ==========================================================================

func Test_AnyErrorOnce_WithError(t *testing.T) {
	// Arrange
	testErr := errors.New("fail")
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return nil, testErr })

	// Act
	actual := args.Map{
		"hasError": aeo.HasError(), "isEmptyError": aeo.IsEmptyError(),
		"isFailed": aeo.IsFailed(), "isSuccess": aeo.IsSuccess(),
		"isNull": aeo.IsNull(), "isEmpty": aeo.IsEmpty(),
		"isValid": aeo.IsValid(), "isInvalid": aeo.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"hasError": true, "isEmptyError": false,
		"isFailed": true, "isSuccess": false,
		"isNull": true, "isEmpty": false,
		"isValid": false, "isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce with error returns expected -- fail", actual)
}

func Test_AnyErrorOnce_ValueString_Nil(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return nil, nil })
	val, err := aeo.ValueString()

	// Act
	actual := args.Map{
		"val": val,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"val": "<nil>",
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce ValueString nil value returns empty -- nil data", actual)
}

func Test_AnyErrorOnce_CastFail(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return 42, nil })
	_, _, okStr := aeo.CastValueString()
	_, _, okStrings := aeo.CastValueStrings()
	_, _, okMap := aeo.CastValueHashmapMap()
	_, _, okMapAny := aeo.CastValueMapStringAnyMap()
	_, _, okBytes := aeo.CastValueBytes()

	// Act
	actual := args.Map{
		"okStr": okStr, "okStrings": okStrings,
		"okMap": okMap, "okMapAny": okMapAny, "okBytes": okBytes,
	}

	// Assert
	expected := args.Map{
		"okStr": false, "okStrings": false,
		"okMap": false, "okMapAny": false, "okBytes": false,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce Cast methods fail -- wrong type", actual)
}

// ==========================================================================
// BytesErrorOnce — with error path
// ==========================================================================

func Test_BytesErrorOnce_WithError(t *testing.T) {
	// Arrange
	testErr := errors.New("bytes error")
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return nil, testErr })

	// Act
	actual := args.Map{
		"hasError": beo.HasError(), "isEmpty": beo.IsEmpty(),
		"isValid": beo.IsValid(), "isInvalid": beo.IsInvalid(),
		"length": beo.Length(),
	}

	// Assert
	expected := args.Map{
		"hasError": true, "isEmpty": false,
		"isValid": false, "isInvalid": true,
		"length": 0,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce with error returns expected -- error", actual)
}

func Test_BytesErrorOnce_Valid(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return []byte("hi"), nil })

	// Act
	actual := args.Map{
		"hasError": beo.HasError(), "isEmpty": beo.IsEmpty(),
		"string": beo.String(), "length": beo.Length(),
	}

	// Assert
	expected := args.Map{
		"hasError": false, "isEmpty": false,
		"string": "hi", "length": 2,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce valid returns expected -- hi", actual)
}

// ==========================================================================
// StringOnce — additional edge cases
// ==========================================================================

func Test_StringOnce_IsEqual_False(t *testing.T) {
	// Arrange
	so := coreonce.NewStringOnce(func() string { return "hello" })

	// Act
	actual := args.Map{
		"notEqual":     so.IsEqual("world"),
		"noPrefix":     so.HasPrefix("xyz"),
		"noSuffix":     so.HasSuffix("xyz"),
		"noContains":   so.IsContains("xyz"),
		"notStartWith": so.IsStartsWith("xyz"),
		"notEndWith":   so.IsEndsWith("xyz"),
	}

	// Assert
	expected := args.Map{
		"notEqual": false, "noPrefix": false, "noSuffix": false,
		"noContains": false, "notStartWith": false, "notEndWith": false,
	}
	expected.ShouldBeEqual(t, 0, "StringOnce negative checks return false -- mismatches", actual)
}

// ==========================================================================
// IntegerOnce — zero value
// ==========================================================================

func Test_IntegerOnce_Zero(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegerOnce(func() int { return 0 })

	// Act
	actual := args.Map{
		"isEmpty": io.IsEmpty(), "isZero": io.IsZero(),
		"isAboveZero": io.IsAboveZero(), "isAboveEqualZero": io.IsAboveEqualZero(),
		"isNegative": io.IsNegative(), "isPositive": io.IsPositive(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true, "isZero": true,
		"isAboveZero": false, "isAboveEqualZero": true,
		"isNegative": false, "isPositive": false,
	}
	expected.ShouldBeEqual(t, 0, "IntegerOnce zero value returns expected -- zero", actual)
}

// ==========================================================================
// ByteOnce — zero value
// ==========================================================================

func Test_ByteOnce_Zero(t *testing.T) {
	// Arrange
	bo := coreonce.NewByteOnce(func() byte { return 0 })

	// Act
	actual := args.Map{
		"isEmpty": bo.IsEmpty(), "isZero": bo.IsZero(),
		"isNegative": bo.IsNegative(), "isPositive": bo.IsPositive(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true, "isZero": true,
		"isNegative": false, "isPositive": false,
	}
	expected.ShouldBeEqual(t, 0, "ByteOnce zero value returns expected -- zero", actual)
}

// ==========================================================================
// IntegersOnce — empty
// ==========================================================================

func Test_IntegersOnce_Empty(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOnce(func() []int { return []int{} })

	// Act
	actual := args.Map{"isEmpty": io.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IntegersOnce empty returns true -- empty slice", actual)
}

// ==========================================================================
// StringsOnce — empty
// ==========================================================================

func Test_StringsOnce_Empty(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOnce(func() []string { return []string{} })

	// Act
	actual := args.Map{
		"isEmpty": so.IsEmpty(), "hasAny": so.HasAnyItem(),
		"length": so.Length(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
		"hasAny": false,
		"length": 0,
	}
	expected.ShouldBeEqual(t, 0, "StringsOnce empty returns expected -- empty slice", actual)
}

// ==========================================================================
// MapStringStringOnce — IsEqual edge: different values
// ==========================================================================

func Test_MapStringStringOnce_Empty(t *testing.T) {
	// Arrange
	mso := coreonce.NewMapStringStringOnce(func() map[string]string { return map[string]string{} })

	// Act
	actual := args.Map{
		"isEmpty": mso.IsEmpty(), "length": mso.Length(),
		"isMissing": mso.IsMissing("any"),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
		"length": 0,
		"isMissing": true,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce empty returns expected -- empty map", actual)
}

func Test_MapStringStringOnce_UnmarshalJSON(t *testing.T) {
	// Arrange
	mso := coreonce.NewMapStringStringOnce(func() map[string]string { return map[string]string{} })
	mb, _ := mso.MarshalJSON()
	err := mso.UnmarshalJSON(mb)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce UnmarshalJSON returns no error -- valid", actual)
}
