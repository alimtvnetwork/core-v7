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
	actual := args.Map{"result": issetter.CombinedBooleans(true, true, true) != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected True", actual)
}

func Test_C8_CombinedBooleans_OneFalse(t *testing.T) {
	actual := args.Map{"result": issetter.CombinedBooleans(true, false, true) != issetter.False}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected False", actual)
}

func Test_C8_CombinedBooleans_Empty(t *testing.T) {
	actual := args.Map{"result": issetter.CombinedBooleans() != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected True for empty", actual)
}

// ═══════════════════════════════════════════════
// GetSet / GetSetByte / GetSetUnset / GetSetterByComparing
// ═══════════════════════════════════════════════

func Test_C8_GetSet_True(t *testing.T) {
	actual := args.Map{"result": issetter.GetSet(true, issetter.Set, issetter.Unset) != issetter.Set}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSet_False(t *testing.T) {
	actual := args.Map{"result": issetter.GetSet(false, issetter.Set, issetter.Unset) != issetter.Unset}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetByte_True(t *testing.T) {
	r := issetter.GetSetByte(true, 1, 2)
	actual := args.Map{"result": r.Value() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetByte_False(t *testing.T) {
	r := issetter.GetSetByte(false, 1, 2)
	actual := args.Map{"result": r.Value() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetUnset_True(t *testing.T) {
	actual := args.Map{"result": issetter.GetSetUnset(true) != issetter.Set}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetUnset_False(t *testing.T) {
	actual := args.Map{"result": issetter.GetSetUnset(false) != issetter.Unset}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetterByComparing_Match(t *testing.T) {
	r := issetter.GetSetterByComparing(issetter.True, issetter.False, "a", "a", "b")
	actual := args.Map{"result": r != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetterByComparing_NoMatch(t *testing.T) {
	r := issetter.GetSetterByComparing(issetter.True, issetter.False, "c", "a", "b")
	actual := args.Map{"result": r != issetter.False}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

// ═══════════════════════════════════════════════
// IsCompareResult — all switch branches
// ═══════════════════════════════════════════════

func Test_C8_IsCompareResult_Equal(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsCompareResult(1, corecomparator.Equal)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsCompareResult_LeftGreater(t *testing.T) {
	actual := args.Map{"result": issetter.Set.IsCompareResult(1, corecomparator.LeftGreater)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsCompareResult_LeftGreaterEqual(t *testing.T) {
	actual := args.Map{"result": issetter.Set.IsCompareResult(4, corecomparator.LeftGreaterEqual)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsCompareResult_LeftLess(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsCompareResult(4, corecomparator.LeftLess)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsCompareResult_LeftLessEqual(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsCompareResult(1, corecomparator.LeftLessEqual)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsCompareResult_NotEqual(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsCompareResult(0, corecomparator.NotEqual)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsCompareResult_DefaultPanic(t *testing.T) {
	defer func() {
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic for out of range comparator", actual)
	}()
	issetter.True.IsCompareResult(1, corecomparator.Compare(99))
}

// ═══════════════════════════════════════════════
// IsOutOfRange / Max / Min / MaxByte / MinByte
// ═══════════════════════════════════════════════

func Test_C8_IsOutOfRange_True(t *testing.T) {
	actual := args.Map{"result": issetter.IsOutOfRange(200)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsOutOfRange_False(t *testing.T) {
	actual := args.Map{"result": issetter.IsOutOfRange(1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Max(t *testing.T) {
	actual := args.Map{"result": issetter.Max() != issetter.Wildcard}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Min(t *testing.T) {
	actual := args.Map{"result": issetter.Min() != issetter.Uninitialized}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MaxByte(t *testing.T) {
	actual := args.Map{"result": issetter.MaxByte() != issetter.Set.Value()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MinByte(t *testing.T) {
	actual := args.Map{"result": issetter.MinByte() != issetter.Uninitialized.Value()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

// ═══════════════════════════════════════════════
// New / NewBool / NewBooleans / NewMust / RangeNamesCsv
// ═══════════════════════════════════════════════

func Test_C8_New_Valid(t *testing.T) {
	v, err := issetter.New("True")
	actual := args.Map{"result": err != nil || v != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_New_Invalid(t *testing.T) {
	_, err := issetter.New("NotExist")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_NewBool_True(t *testing.T) {
	actual := args.Map{"result": issetter.NewBool(true) != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_NewBool_False(t *testing.T) {
	actual := args.Map{"result": issetter.NewBool(false) != issetter.False}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_NewBooleans_AllTrue(t *testing.T) {
	actual := args.Map{"result": issetter.NewBooleans(true, true) != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_NewMust_Valid(t *testing.T) {
	v := issetter.NewMust("Set")
	actual := args.Map{"result": v != issetter.Set}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_RangeNamesCsv(t *testing.T) {
	csv := issetter.RangeNamesCsv()
	actual := args.Map{"result": csv == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

// ═══════════════════════════════════════════════
// Value methods — all uncovered from coverage.out
// ═══════════════════════════════════════════════

func Test_C8_AllNameValues(t *testing.T) {
	names := issetter.True.AllNameValues()
	actual := args.Map{"result": len(names) != 6}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 6", actual)
}

func Test_C8_OnlySupportedErr_NoNames(t *testing.T) {
	err := issetter.True.OnlySupportedErr()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_OnlySupportedErr_WithNames(t *testing.T) {
	err := issetter.True.OnlySupportedErr("True", "False")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for unsupported names", actual)
}

func Test_C8_OnlySupportedErr_AllNames(t *testing.T) {
	err := issetter.True.OnlySupportedErr(
		"Uninitialized", "True", "False", "Unset", "Set", "Wildcard")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "all names should be supported", actual)
}

func Test_C8_OnlySupportedMsgErr_NoError(t *testing.T) {
	err := issetter.True.OnlySupportedMsgErr("prefix: ")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_OnlySupportedMsgErr_WithError(t *testing.T) {
	err := issetter.True.OnlySupportedMsgErr("prefix: ", "True")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ValueUInt16(t *testing.T) {
	actual := args.Map{"result": issetter.True.ValueUInt16() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IntegerEnumRanges(t *testing.T) {
	r := issetter.True.IntegerEnumRanges()
	actual := args.Map{"result": len(r) != 6}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MinMaxAny(t *testing.T) {
	min, max := issetter.True.MinMaxAny()
	actual := args.Map{"result": min == nil || max == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MinValueString(t *testing.T) {
	actual := args.Map{"result": issetter.True.MinValueString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MaxValueString(t *testing.T) {
	actual := args.Map{"result": issetter.True.MaxValueString() == ""}
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
	m := issetter.True.RangesDynamicMap()
	actual := args.Map{"result": len(m) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsValueEqual(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsValueEqual(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_RangeNamesCsvMethod(t *testing.T) {
	actual := args.Map{"result": issetter.True.RangeNamesCsv() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsByteValueEqual(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsByteValueEqual(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsOn(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsOn()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsOff(t *testing.T) {
	actual := args.Map{"result": issetter.False.IsOff()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsLater(t *testing.T) {
	actual := args.Map{"result": issetter.Uninitialized.IsLater()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsNot(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsNot(issetter.False)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsNo(t *testing.T) {
	actual := args.Map{"result": issetter.False.IsNo()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsAsk(t *testing.T) {
	actual := args.Map{"result": issetter.Uninitialized.IsAsk()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsIndeterminate(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.IsIndeterminate()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsAccept(t *testing.T) {
	actual := args.Map{"result": issetter.Set.IsAccept()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsReject(t *testing.T) {
	actual := args.Map{"result": issetter.Unset.IsReject()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsFailed(t *testing.T) {
	actual := args.Map{"result": issetter.False.IsFailed()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsSuccess(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsSuccess()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsSkip(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.IsSkip()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_NameValue(t *testing.T) {
	nv := issetter.True.NameValue()
	actual := args.Map{"result": nv == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsNameEqual(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsNameEqual("True")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsAnyNamesOf_Match(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsAnyNamesOf("True", "False")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsAnyNamesOf_NoMatch(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsAnyNamesOf("False", "Unset")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ToNumberString(t *testing.T) {
	actual := args.Map{"result": issetter.True.ToNumberString() != "1"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ValueByte(t *testing.T) {
	actual := args.Map{"result": issetter.True.ValueByte() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ValueInt(t *testing.T) {
	actual := args.Map{"result": issetter.True.ValueInt() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ValueInt8(t *testing.T) {
	actual := args.Map{"result": issetter.True.ValueInt8() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ValueInt16(t *testing.T) {
	actual := args.Map{"result": issetter.True.ValueInt16() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ValueInt32(t *testing.T) {
	actual := args.Map{"result": issetter.True.ValueInt32() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ValueString(t *testing.T) {
	actual := args.Map{"result": issetter.True.ValueString() != "1"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Format(t *testing.T) {
	s := issetter.True.Format("{name}={value}")
	actual := args.Map{"result": s != "True=1"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "got", actual)
}

func Test_C8_EnumType(t *testing.T) {
	actual := args.Map{"result": issetter.True.EnumType() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_StringValue(t *testing.T) {
	actual := args.Map{"result": issetter.True.StringValue() != "1"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_String(t *testing.T) {
	actual := args.Map{"result": issetter.True.String() != "True"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsTrueOrSet(t *testing.T) {
	actual := args.Map{"result": issetter.Set.IsTrueOrSet()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsSet(t *testing.T) {
	actual := args.Map{"result": issetter.Set.IsSet()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsUnset(t *testing.T) {
	actual := args.Map{"result": issetter.Unset.IsUnset()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_HasInitialized(t *testing.T) {
	actual := args.Map{"result": issetter.True.HasInitialized()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_HasInitializedAndSet(t *testing.T) {
	actual := args.Map{"result": issetter.Set.HasInitializedAndSet()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_HasInitializedAndTrue(t *testing.T) {
	actual := args.Map{"result": issetter.True.HasInitializedAndTrue()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsWildcard(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.IsWildcard()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsInit(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsInit()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsInitBoolean(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsInitBoolean()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual := args.Map{"result": issetter.False.IsInitBoolean()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsDefinedBoolean(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsDefinedBoolean()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsInitBooleanWild(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.IsInitBooleanWild()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsInitSet(t *testing.T) {
	actual := args.Map{"result": issetter.Set.IsInitSet()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual := args.Map{"result": issetter.Unset.IsInitSet()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsInitSetWild(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.IsInitSetWild()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsYes(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsYes()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Boolean(t *testing.T) {
	actual := args.Map{"result": issetter.True.Boolean()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsOnLogically(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsOnLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsOffLogically(t *testing.T) {
	actual := args.Map{"result": issetter.False.IsOffLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsAccepted(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsAccepted()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsRejected(t *testing.T) {
	actual := args.Map{"result": issetter.False.IsRejected()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsDefinedLogically(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsDefinedLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual := args.Map{"result": issetter.Uninitialized.IsDefinedLogically()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsUndefinedLogically(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.IsUndefinedLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsInvalid(t *testing.T) {
	actual := args.Map{"result": issetter.Uninitialized.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsValid(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetBoolOnInvalid_AlreadyDefined(t *testing.T) {
	v := issetter.True
	r := v.GetSetBoolOnInvalid(false)
	actual := args.Map{"result": r}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetBoolOnInvalid_SetNew(t *testing.T) {
	v := issetter.Uninitialized
	r := v.GetSetBoolOnInvalid(true)
	actual := args.Map{"result": r}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetBoolOnInvalidFunc_AlreadyDefined(t *testing.T) {
	v := issetter.False
	r := v.GetSetBoolOnInvalidFunc(func() bool { return true })
	actual := args.Map{"result": r}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetSetBoolOnInvalidFunc_SetNew(t *testing.T) {
	v := issetter.Uninitialized
	r := v.GetSetBoolOnInvalidFunc(func() bool { return true })
	actual := args.Map{"result": r}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ToBooleanValue(t *testing.T) {
	actual := args.Map{"result": issetter.Set.ToBooleanValue() != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ToSetUnsetValue(t *testing.T) {
	actual := args.Map{"result": issetter.True.ToSetUnsetValue() != issetter.Set}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_LazyEvaluateBool_NotCalled(t *testing.T) {
	v := issetter.True
	called := v.LazyEvaluateBool(func() { t.Fatal("should not call") })
	actual := args.Map{"result": called}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_LazyEvaluateBool_Called(t *testing.T) {
	v := issetter.Uninitialized
	executed := false
	called := v.LazyEvaluateBool(func() { executed = true })
	actual := args.Map{"result": called || !executed}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_LazyEvaluateSet_NotCalled(t *testing.T) {
	v := issetter.Set
	called := v.LazyEvaluateSet(func() { t.Fatal("should not call") })
	actual := args.Map{"result": called}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_LazyEvaluateSet_Called(t *testing.T) {
	v := issetter.Uninitialized
	executed := false
	called := v.LazyEvaluateSet(func() { executed = true })
	actual := args.Map{"result": called || !executed}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsWildcardOrBool_Wildcard(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.IsWildcardOrBool(false)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsWildcardOrBool_TrueTrue(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsWildcardOrBool(true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsWildcardOrBool_FalseFalse(t *testing.T) {
	actual := args.Map{"result": issetter.False.IsWildcardOrBool(false)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ToByteCondition(t *testing.T) {
	actual := args.Map{"result": issetter.True.ToByteCondition(10, 20, 30) != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual := args.Map{"result": issetter.False.ToByteCondition(10, 20, 30) != 20}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual := args.Map{"result": issetter.Uninitialized.ToByteCondition(10, 20, 30) != 30}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_ToByteConditionWithWildcard(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.ToByteConditionWithWildcard(99, 10, 20, 30) != 99}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual := args.Map{"result": issetter.True.ToByteConditionWithWildcard(99, 10, 20, 30) != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_WildcardApply_Wildcard(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.WildcardApply(true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_WildcardApply_Defined(t *testing.T) {
	actual := args.Map{"result": issetter.True.WildcardApply(false)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_WildcardValueApply(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.WildcardValueApply(issetter.True)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual := args.Map{"result": issetter.True.WildcardValueApply(issetter.False)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_OrBool(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.OrBool(true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual := args.Map{"result": issetter.True.OrBool(false)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_OrValue(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.OrValue(issetter.True)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual := args.Map{"result": issetter.True.OrValue(issetter.False)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_AndBool(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.AndBool(true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual := args.Map{"result": issetter.True.AndBool(false)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_And(t *testing.T) {
	r := issetter.Wildcard.And(issetter.True)
	actual := args.Map{"result": r != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	r2 := issetter.True.And(issetter.False)
	actual := args.Map{"result": r2 != issetter.False}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsUninitialized(t *testing.T) {
	actual := args.Map{"result": issetter.Uninitialized.IsUninitialized()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsInitialized(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsInitialized()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsUnSetOrUninitialized(t *testing.T) {
	actual := args.Map{"result": issetter.Unset.IsUnSetOrUninitialized()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsNegative(t *testing.T) {
	actual := args.Map{"result": issetter.Uninitialized.IsNegative()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual := args.Map{"result": issetter.False.IsNegative()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsPositive(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsPositive()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual := args.Map{"result": issetter.Set.IsPositive()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsBetween(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsBetween(0, 5)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsBetweenInt(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsBetweenInt(0, 5)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Add(t *testing.T) {
	r := issetter.True.Add(1)
	actual := args.Map{"result": r != issetter.False}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Is(t *testing.T) {
	actual := args.Map{"result": issetter.True.Is(issetter.True)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsEqual(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsEqual(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsGreater(t *testing.T) {
	actual := args.Map{"result": issetter.Set.IsGreater(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsGreaterEqual(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsGreaterEqual(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsLess(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsLess(4)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsLessEqual(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsLessEqual(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsEqualInt(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsEqualInt(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsGreaterInt(t *testing.T) {
	actual := args.Map{"result": issetter.Set.IsGreaterInt(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsGreaterEqualInt(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsGreaterEqualInt(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsLessInt(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsLessInt(4)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsLessEqualInt(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsLessEqualInt(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_PanicOnOutOfRange_InRange(t *testing.T) {
	issetter.True.PanicOnOutOfRange(1, "msg")
}

func Test_C8_PanicOnOutOfRange_OutOfRange(t *testing.T) {
	defer func() {
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	issetter.True.PanicOnOutOfRange(200, "out of range")
}

func Test_C8_GetErrorOnOutOfRange_InRange(t *testing.T) {
	err := issetter.True.GetErrorOnOutOfRange(1, "msg")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_GetErrorOnOutOfRange_OutOfRange(t *testing.T) {
	err := issetter.True.GetErrorOnOutOfRange(200, "msg")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Name(t *testing.T) {
	actual := args.Map{"result": issetter.True.Name() != "True"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_YesNoMappedValue_Uninit(t *testing.T) {
	actual := args.Map{"result": issetter.Uninitialized.YesNoMappedValue() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_YesNoMappedValue_True(t *testing.T) {
	actual := args.Map{"result": issetter.True.YesNoMappedValue() != "yes"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_YesNoMappedValue_False(t *testing.T) {
	actual := args.Map{"result": issetter.False.YesNoMappedValue() != "no"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_YesNoLowercaseName(t *testing.T) {
	actual := args.Map{"result": issetter.True.YesNoLowercaseName() != "yes"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_YesNoName(t *testing.T) {
	actual := args.Map{"result": issetter.True.YesNoName() != "Yes"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_TrueFalseName(t *testing.T) {
	actual := args.Map{"result": issetter.True.TrueFalseName() != "True"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_OnOffLowercaseName(t *testing.T) {
	actual := args.Map{"result": issetter.True.OnOffLowercaseName() != "on"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_OnOffName(t *testing.T) {
	actual := args.Map{"result": issetter.True.OnOffName() != "On"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_TrueFalseLowercaseName(t *testing.T) {
	actual := args.Map{"result": issetter.True.TrueFalseLowercaseName() != "true"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_SetUnsetLowercaseName(t *testing.T) {
	actual := args.Map{"result": issetter.True.SetUnsetLowercaseName() != "set"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MarshalJSON(t *testing.T) {
	b, err := issetter.True.MarshalJSON()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_UnmarshalJSON_Valid(t *testing.T) {
	v := issetter.Uninitialized
	err := v.UnmarshalJSON([]byte(`"True"`))
	actual := args.Map{"result": err != nil || v != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_UnmarshalJSON_Nil(t *testing.T) {
	v := issetter.Uninitialized
	err := v.UnmarshalJSON(nil)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_UnmarshalJSON_Invalid(t *testing.T) {
	v := issetter.Uninitialized
	err := v.UnmarshalJSON([]byte(`"BOGUS"`))
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Serialize(t *testing.T) {
	b, err := issetter.True.Serialize()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_TypeName(t *testing.T) {
	actual := args.Map{"result": issetter.True.TypeName() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsAnyValuesEqual_Match(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsAnyValuesEqual(0, 1, 2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_IsAnyValuesEqual_NoMatch(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsAnyValuesEqual(0, 2, 3)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_UnmarshallEnumToValue(t *testing.T) {
	v, err := issetter.Uninitialized.UnmarshallEnumToValue([]byte(`"Set"`))
	actual := args.Map{"result": err != nil || v != issetter.Set.Value()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Deserialize_Valid(t *testing.T) {
	v, err := issetter.Uninitialized.Deserialize([]byte(`"True"`))
	actual := args.Map{"result": err != nil || v != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_Deserialize_Invalid(t *testing.T) {
	_, err := issetter.Uninitialized.Deserialize([]byte(`"BOGUS"`))
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MaxByteMethod(t *testing.T) {
	actual := args.Map{"result": issetter.True.MaxByte() != issetter.Wildcard.ValueByte()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_MinByteMethod(t *testing.T) {
	actual := args.Map{"result": issetter.True.MinByte() != issetter.Uninitialized.ValueByte()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C8_RangesByte_Panics(t *testing.T) {
	defer func() {
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	issetter.True.RangesByte()
}

func Test_C8_ToPtr(t *testing.T) {
	p := issetter.True.ToPtr()
	actual := args.Map{"result": p == nil || *p != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}
