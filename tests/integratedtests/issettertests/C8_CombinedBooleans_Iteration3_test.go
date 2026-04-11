package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/issetter"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════
// CombinedBooleans
// ═══════════════════════════════════════════════

func Test_C8_CombinedBooleans_AllTrue(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.CombinedBooleans(true, true, true) != issetter.True}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected True", actual)
}

func Test_C8_CombinedBooleans_OneFalse(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.CombinedBooleans(true, false, true) != issetter.False}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected False", actual)
}

func Test_C8_CombinedBooleans_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.CombinedBooleans() != issetter.True}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected True for empty", actual)
}

// ═══════════════════════════════════════════════
// GetSet / GetSetByte / GetSetUnset / GetSetterByComparing
// ═══════════════════════════════════════════════

func Test_C8_GetSet_True(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.GetSet(true, issetter.Set, issetter.Unset) != issetter.Set}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSet_False(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.GetSet(false, issetter.Set, issetter.Unset) != issetter.Unset}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetByte_True(t *testing.T) {
	// Arrange
	r := issetter.GetSetByte(true, 1, 2)

	// Act
	actual := args.Map{"result": r.Value() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetByte_False(t *testing.T) {
	// Arrange
	r := issetter.GetSetByte(false, 1, 2)

	// Act
	actual := args.Map{"result": r.Value() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetUnset_True(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.GetSetUnset(true) != issetter.Set}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetUnset_False(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.GetSetUnset(false) != issetter.Unset}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetterByComparing_Match(t *testing.T) {
	// Arrange
	r := issetter.GetSetterByComparing(issetter.True, issetter.False, "a", "a", "b")

	// Act
	actual := args.Map{"result": r != issetter.True}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetterByComparing_NoMatch(t *testing.T) {
	// Arrange
	r := issetter.GetSetterByComparing(issetter.True, issetter.False, "c", "a", "b")

	// Act
	actual := args.Map{"result": r != issetter.False}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

// ═══════════════════════════════════════════════
// IsCompareResult — all switch branches
// ═══════════════════════════════════════════════

func Test_C8_IsCompareResult_Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsCompareResult(1, corecomparator.Equal)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsCompareResult_LeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Set.IsCompareResult(1, corecomparator.LeftGreater)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsCompareResult_LeftGreaterEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Set.IsCompareResult(4, corecomparator.LeftGreaterEqual)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsCompareResult_LeftLess(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsCompareResult(4, corecomparator.LeftLess)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsCompareResult_LeftLessEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsCompareResult(1, corecomparator.LeftLessEqual)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsCompareResult_NotEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsCompareResult(0, corecomparator.NotEqual)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsCompareResult_DefaultPanic(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic for out of range comparator", actual)
	}()
	issetter.True.IsCompareResult(1, corecomparator.Compare(99))
}

// ═══════════════════════════════════════════════
// IsOutOfRange / Max / Min / MaxByte / MinByte
// ═══════════════════════════════════════════════

func Test_C8_IsOutOfRange_True(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.IsOutOfRange(200)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsOutOfRange_False(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.IsOutOfRange(1)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Max(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Max() != issetter.Wildcard}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Min(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Min() != issetter.Uninitialized}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MaxByte(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.MaxByte() != issetter.Set.Value()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MinByte(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.MinByte() != issetter.Uninitialized.Value()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

// ═══════════════════════════════════════════════
// New / NewBool / NewBooleans / NewMust / RangeNamesCsv
// ═══════════════════════════════════════════════

func Test_C8_New_Valid(t *testing.T) {
	// Arrange
	v, err := issetter.New("True")

	// Act
	actual := args.Map{"result": err != nil || v != issetter.True}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_New_Invalid(t *testing.T) {
	// Arrange
	_, err := issetter.New("NotExist")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_NewBool_True(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.NewBool(true) != issetter.True}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_NewBool_False(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.NewBool(false) != issetter.False}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_NewBooleans_AllTrue(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.NewBooleans(true, true) != issetter.True}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_NewMust_Valid(t *testing.T) {
	// Arrange
	v := issetter.NewMust("Set")

	// Act
	actual := args.Map{"result": v != issetter.Set}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_RangeNamesCsv(t *testing.T) {
	// Arrange
	csv := issetter.RangeNamesCsv()

	// Act
	actual := args.Map{"result": csv == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

// ═══════════════════════════════════════════════
// Value methods — all uncovered from coverage.out
// ═══════════════════════════════════════════════

func Test_C8_AllNameValues(t *testing.T) {
	// Arrange
	names := issetter.True.AllNameValues()

	// Act
	actual := args.Map{"result": len(names) != 6}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 6", actual)
}

func Test_C8_OnlySupportedErr_NoNames(t *testing.T) {
	// Arrange
	err := issetter.True.OnlySupportedErr()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_OnlySupportedErr_WithNames(t *testing.T) {
	// Arrange
	err := issetter.True.OnlySupportedErr("True", "False")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for unsupported names", actual)
}

func Test_C8_OnlySupportedErr_AllNames(t *testing.T) {
	// Arrange
	err := issetter.True.OnlySupportedErr(
		"Uninitialized", "True", "False", "Unset", "Set", "Wildcard")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "all names should be supported", actual)
}

func Test_C8_OnlySupportedMsgErr_NoError(t *testing.T) {
	// Arrange
	err := issetter.True.OnlySupportedMsgErr("prefix: ")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_OnlySupportedMsgErr_WithError(t *testing.T) {
	// Arrange
	err := issetter.True.OnlySupportedMsgErr("prefix: ", "True")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ValueUInt16(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ValueUInt16() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IntegerEnumRanges(t *testing.T) {
	// Arrange
	r := issetter.True.IntegerEnumRanges()

	// Act
	actual := args.Map{"result": len(r) != 6}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MinMaxAny(t *testing.T) {
	// Arrange
	min, max := issetter.True.MinMaxAny()

	// Act
	actual := args.Map{"result": min == nil || max == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MinValueString(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.MinValueString() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MaxValueString(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.MaxValueString() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MaxInt(t *testing.T) {
	_ = issetter.True.MaxInt()
}

func Test_C8_MinInt(t *testing.T) {
	_ = issetter.True.MinInt()
}

func Test_C8_RangesDynamicMap(t *testing.T) {
	// Arrange
	m := issetter.True.RangesDynamicMap()

	// Act
	actual := args.Map{"result": len(m) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsValueEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsValueEqual(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_RangeNamesCsvMethod(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.RangeNamesCsv() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsByteValueEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsByteValueEqual(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsOn(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsOn()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsOff(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.False.IsOff()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsLater(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Uninitialized.IsLater()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsNot(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsNot(issetter.False)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsNo(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.False.IsNo()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsAsk(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Uninitialized.IsAsk()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsIndeterminate(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.IsIndeterminate()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsAccept(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Set.IsAccept()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsReject(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Unset.IsReject()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsFailed(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.False.IsFailed()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsSuccess(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsSuccess()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsSkip(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.IsSkip()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_NameValue(t *testing.T) {
	// Arrange
	nv := issetter.True.NameValue()

	// Act
	actual := args.Map{"result": nv == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsNameEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsNameEqual("True")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsAnyNamesOf_Match(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsAnyNamesOf("True", "False")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsAnyNamesOf_NoMatch(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsAnyNamesOf("False", "Unset")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ToNumberString(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ToNumberString() != "1"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ValueByte(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ValueByte() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ValueInt(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ValueInt() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ValueInt8(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ValueInt8() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ValueInt16(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ValueInt16() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ValueInt32(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ValueInt32() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ValueString(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ValueString() != "1"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Format(t *testing.T) {
	// Arrange
	s := issetter.True.Format("{name}={value}")

	// Act
	actual := args.Map{"result": s != "True=1"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "got", actual)
}

func Test_C8_EnumType(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.EnumType() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_StringValue(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.StringValue() != "1"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_String(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.String() != "True"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsTrueOrSet(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Set.IsTrueOrSet()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsSet(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Set.IsSet()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsUnset(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Unset.IsUnset()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_HasInitialized(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.HasInitialized()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_HasInitializedAndSet(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Set.HasInitializedAndSet()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_HasInitializedAndTrue(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.HasInitializedAndTrue()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsWildcard(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.IsWildcard()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsInit(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsInit()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsInitBoolean(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsInitBoolean()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": issetter.False.IsInitBoolean()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsDefinedBoolean(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsDefinedBoolean()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsInitBooleanWild(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.IsInitBooleanWild()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsInitSet(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Set.IsInitSet()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": issetter.Unset.IsInitSet()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsInitSetWild(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.IsInitSetWild()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsYes(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsYes()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Boolean(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.Boolean()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsOnLogically(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsOnLogically()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsOffLogically(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.False.IsOffLogically()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsAccepted(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsAccepted()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsRejected(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.False.IsRejected()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsDefinedLogically(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsDefinedLogically()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": issetter.Uninitialized.IsDefinedLogically()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsUndefinedLogically(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.IsUndefinedLogically()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsInvalid(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Uninitialized.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsValid(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetBoolOnInvalid_AlreadyDefined(t *testing.T) {
	// Arrange
	v := issetter.True
	r := v.GetSetBoolOnInvalid(false)

	// Act
	actual := args.Map{"result": r}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetBoolOnInvalid_SetNew(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	r := v.GetSetBoolOnInvalid(true)

	// Act
	actual := args.Map{"result": r}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetBoolOnInvalidFunc_AlreadyDefined(t *testing.T) {
	// Arrange
	v := issetter.False
	r := v.GetSetBoolOnInvalidFunc(func() bool { return true })

	// Act
	actual := args.Map{"result": r}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetBoolOnInvalidFunc_SetNew(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	r := v.GetSetBoolOnInvalidFunc(func() bool { return true })

	// Act
	actual := args.Map{"result": r}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ToBooleanValue(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Set.ToBooleanValue() != issetter.True}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ToSetUnsetValue(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ToSetUnsetValue() != issetter.Set}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_LazyEvaluateBool_NotCalled(t *testing.T) {
	// Arrange
	v := issetter.True

	// Assert
	called := v.LazyEvaluateBool(func() { actual := args.Map{"called": true}; expected := args.Map{"called": false}; expected.ShouldBeEqual(t, 0, "LazyEvaluateBool should not call", actual) })

	// Act
	actual := args.Map{"result": called}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_LazyEvaluateBool_Called(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	executed := false
	called := v.LazyEvaluateBool(func() { executed = true })

	// Act
	actual := args.Map{"result": called || !executed}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_LazyEvaluateSet_NotCalled(t *testing.T) {
	// Arrange
	v := issetter.Set

	// Assert
	called := v.LazyEvaluateSet(func() { actual := args.Map{"called": true}; expected := args.Map{"called": false}; expected.ShouldBeEqual(t, 0, "LazyEvaluateSet should not call", actual) })

	// Act
	actual := args.Map{"result": called}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_LazyEvaluateSet_Called(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	executed := false
	called := v.LazyEvaluateSet(func() { executed = true })

	// Act
	actual := args.Map{"result": called || !executed}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsWildcardOrBool_Wildcard(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.IsWildcardOrBool(false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsWildcardOrBool_TrueTrue(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsWildcardOrBool(true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsWildcardOrBool_FalseFalse(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.False.IsWildcardOrBool(false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ToByteCondition(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ToByteCondition(10, 20, 30) != 10}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": issetter.False.ToByteCondition(10, 20, 30) != 20}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": issetter.Uninitialized.ToByteCondition(10, 20, 30) != 30}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ToByteConditionWithWildcard(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.ToByteConditionWithWildcard(99, 10, 20, 30) != 99}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": issetter.True.ToByteConditionWithWildcard(99, 10, 20, 30) != 10}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_WildcardApply_Wildcard(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.WildcardApply(true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_WildcardApply_Defined(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.WildcardApply(false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_WildcardValueApply(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.WildcardValueApply(issetter.True)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": issetter.True.WildcardValueApply(issetter.False)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_OrBool(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.OrBool(true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": issetter.True.OrBool(false)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_OrValue(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.OrValue(issetter.True)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": issetter.True.OrValue(issetter.False)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_AndBool(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.AndBool(true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": issetter.True.AndBool(false)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_And(t *testing.T) {
	// Arrange
	r := issetter.Wildcard.And(issetter.True)

	// Act
	actual := args.Map{"result": r != issetter.True}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	r2 := issetter.True.And(issetter.False)
	actual = args.Map{"result": r2 != issetter.False}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsUninitialized(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Uninitialized.IsUninitialized()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsInitialized(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsInitialized()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsUnSetOrUninitialized(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Unset.IsUnSetOrUninitialized()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsNegative(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Uninitialized.IsNegative()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": issetter.False.IsNegative()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsPositive(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsPositive()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual = args.Map{"result": issetter.Set.IsPositive()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsBetween(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsBetween(0, 5)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsBetweenInt(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsBetweenInt(0, 5)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Add(t *testing.T) {
	// Arrange
	r := issetter.True.Add(1)

	// Act
	actual := args.Map{"result": r != issetter.False}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Is(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.Is(issetter.True)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsEqual(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Set.IsGreater(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsGreaterEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsGreaterEqual(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsLess(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsLess(4)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsLessEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsLessEqual(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsEqualInt(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsEqualInt(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsGreaterInt(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Set.IsGreaterInt(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsGreaterEqualInt(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsGreaterEqualInt(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsLessInt(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsLessInt(4)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsLessEqualInt(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsLessEqualInt(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_PanicOnOutOfRange_InRange(t *testing.T) {
	issetter.True.PanicOnOutOfRange(1, "msg")
}

func Test_C8_PanicOnOutOfRange_OutOfRange(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	issetter.True.PanicOnOutOfRange(200, "out of range")
}

func Test_C8_GetErrorOnOutOfRange_InRange(t *testing.T) {
	// Arrange
	err := issetter.True.GetErrorOnOutOfRange(1, "msg")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetErrorOnOutOfRange_OutOfRange(t *testing.T) {
	// Arrange
	err := issetter.True.GetErrorOnOutOfRange(200, "msg")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Name(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.Name() != "True"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_YesNoMappedValue_Uninit(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Uninitialized.YesNoMappedValue() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_YesNoMappedValue_True(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.YesNoMappedValue() != "yes"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_YesNoMappedValue_False(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.False.YesNoMappedValue() != "no"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_YesNoLowercaseName(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.YesNoLowercaseName() != "yes"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_YesNoName(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.YesNoName() != "Yes"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_TrueFalseName(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.TrueFalseName() != "True"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_OnOffLowercaseName(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.OnOffLowercaseName() != "on"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_OnOffName(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.OnOffName() != "On"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_TrueFalseLowercaseName(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.TrueFalseLowercaseName() != "true"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_SetUnsetLowercaseName(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.SetUnsetLowercaseName() != "set"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MarshalJSON(t *testing.T) {
	// Arrange
	b, err := issetter.True.MarshalJSON()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_UnmarshalJSON_Valid(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	err := v.UnmarshalJSON([]byte(`"True"`))

	// Act
	actual := args.Map{"result": err != nil || v != issetter.True}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_UnmarshalJSON_Nil(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	err := v.UnmarshalJSON(nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_UnmarshalJSON_Invalid(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	err := v.UnmarshalJSON([]byte(`"BOGUS"`))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Serialize(t *testing.T) {
	// Arrange
	b, err := issetter.True.Serialize()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_TypeName(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.TypeName() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsAnyValuesEqual_Match(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsAnyValuesEqual(0, 1, 2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsAnyValuesEqual_NoMatch(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsAnyValuesEqual(0, 2, 3)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_UnmarshallEnumToValue(t *testing.T) {
	// Arrange
	v, err := issetter.Uninitialized.UnmarshallEnumToValue([]byte(`"Set"`))

	// Act
	actual := args.Map{"result": err != nil || v != issetter.Set.Value()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Deserialize_Valid(t *testing.T) {
	// Arrange
	v, err := issetter.Uninitialized.Deserialize([]byte(`"True"`))

	// Act
	actual := args.Map{"result": err != nil || v != issetter.True}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Deserialize_Invalid(t *testing.T) {
	// Arrange
	_, err := issetter.Uninitialized.Deserialize([]byte(`"BOGUS"`))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MaxByteMethod(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.MaxByte() != issetter.Wildcard.ValueByte()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MinByteMethod(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.MinByte() != issetter.Uninitialized.ValueByte()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_RangesByte_Panics(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	issetter.True.RangesByte()
}

func Test_C8_ToPtr(t *testing.T) {
	// Arrange
	p := issetter.True.ToPtr()

	// Act
	actual := args.Map{"result": p == nil || *p != issetter.True}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}
