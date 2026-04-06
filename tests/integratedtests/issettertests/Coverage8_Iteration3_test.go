package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/issetter"
)

// ═══════════════════════════════════════════════
// CombinedBooleans
// ═══════════════════════════════════════════════

func Test_C8_CombinedBooleans_AllTrue(t *testing.T) {
	if issetter.CombinedBooleans(true, true, true) != issetter.True {
		t.Fatal("expected True")
	}
}

func Test_C8_CombinedBooleans_OneFalse(t *testing.T) {
	if issetter.CombinedBooleans(true, false, true) != issetter.False {
		t.Fatal("expected False")
	}
}

func Test_C8_CombinedBooleans_Empty(t *testing.T) {
	if issetter.CombinedBooleans() != issetter.True {
		t.Fatal("expected True for empty")
	}
}

// ═══════════════════════════════════════════════
// GetSet / GetSetByte / GetSetUnset / GetSetterByComparing
// ═══════════════════════════════════════════════

func Test_C8_GetSet_True(t *testing.T) {
	if issetter.GetSet(true, issetter.Set, issetter.Unset) != issetter.Set {
		t.Fatal()
	}
}

func Test_C8_GetSet_False(t *testing.T) {
	if issetter.GetSet(false, issetter.Set, issetter.Unset) != issetter.Unset {
		t.Fatal()
	}
}

func Test_C8_GetSetByte_True(t *testing.T) {
	r := issetter.GetSetByte(true, 1, 2)
	if r.Value() != 1 {
		t.Fatal()
	}
}

func Test_C8_GetSetByte_False(t *testing.T) {
	r := issetter.GetSetByte(false, 1, 2)
	if r.Value() != 2 {
		t.Fatal()
	}
}

func Test_C8_GetSetUnset_True(t *testing.T) {
	if issetter.GetSetUnset(true) != issetter.Set {
		t.Fatal()
	}
}

func Test_C8_GetSetUnset_False(t *testing.T) {
	if issetter.GetSetUnset(false) != issetter.Unset {
		t.Fatal()
	}
}

func Test_C8_GetSetterByComparing_Match(t *testing.T) {
	r := issetter.GetSetterByComparing(issetter.True, issetter.False, "a", "a", "b")
	if r != issetter.True {
		t.Fatal()
	}
}

func Test_C8_GetSetterByComparing_NoMatch(t *testing.T) {
	r := issetter.GetSetterByComparing(issetter.True, issetter.False, "c", "a", "b")
	if r != issetter.False {
		t.Fatal()
	}
}

// ═══════════════════════════════════════════════
// IsCompareResult — all switch branches
// ═══════════════════════════════════════════════

func Test_C8_IsCompareResult_Equal(t *testing.T) {
	if !issetter.True.IsCompareResult(1, corecomparator.Equal) {
		t.Fatal()
	}
}

func Test_C8_IsCompareResult_LeftGreater(t *testing.T) {
	if !issetter.Set.IsCompareResult(1, corecomparator.LeftGreater) {
		t.Fatal()
	}
}

func Test_C8_IsCompareResult_LeftGreaterEqual(t *testing.T) {
	if !issetter.Set.IsCompareResult(4, corecomparator.LeftGreaterEqual) {
		t.Fatal()
	}
}

func Test_C8_IsCompareResult_LeftLess(t *testing.T) {
	if !issetter.True.IsCompareResult(4, corecomparator.LeftLess) {
		t.Fatal()
	}
}

func Test_C8_IsCompareResult_LeftLessEqual(t *testing.T) {
	if !issetter.True.IsCompareResult(1, corecomparator.LeftLessEqual) {
		t.Fatal()
	}
}

func Test_C8_IsCompareResult_NotEqual(t *testing.T) {
	if !issetter.True.IsCompareResult(0, corecomparator.NotEqual) {
		t.Fatal()
	}
}

func Test_C8_IsCompareResult_DefaultPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for out of range comparator")
		}
	}()
	issetter.True.IsCompareResult(1, corecomparator.Compare(99))
}

// ═══════════════════════════════════════════════
// IsOutOfRange / Max / Min / MaxByte / MinByte
// ═══════════════════════════════════════════════

func Test_C8_IsOutOfRange_True(t *testing.T) {
	if !issetter.IsOutOfRange(200) {
		t.Fatal()
	}
}

func Test_C8_IsOutOfRange_False(t *testing.T) {
	if issetter.IsOutOfRange(1) {
		t.Fatal()
	}
}

func Test_C8_Max(t *testing.T) {
	if issetter.Max() != issetter.Wildcard {
		t.Fatal()
	}
}

func Test_C8_Min(t *testing.T) {
	if issetter.Min() != issetter.Uninitialized {
		t.Fatal()
	}
}

func Test_C8_MaxByte(t *testing.T) {
	if issetter.MaxByte() != issetter.Set.Value() {
		t.Fatal()
	}
}

func Test_C8_MinByte(t *testing.T) {
	if issetter.MinByte() != issetter.Uninitialized.Value() {
		t.Fatal()
	}
}

// ═══════════════════════════════════════════════
// New / NewBool / NewBooleans / NewMust / RangeNamesCsv
// ═══════════════════════════════════════════════

func Test_C8_New_Valid(t *testing.T) {
	v, err := issetter.New("True")
	if err != nil || v != issetter.True {
		t.Fatal()
	}
}

func Test_C8_New_Invalid(t *testing.T) {
	_, err := issetter.New("NotExist")
	if err == nil {
		t.Fatal()
	}
}

func Test_C8_NewBool_True(t *testing.T) {
	if issetter.NewBool(true) != issetter.True {
		t.Fatal()
	}
}

func Test_C8_NewBool_False(t *testing.T) {
	if issetter.NewBool(false) != issetter.False {
		t.Fatal()
	}
}

func Test_C8_NewBooleans_AllTrue(t *testing.T) {
	if issetter.NewBooleans(true, true) != issetter.True {
		t.Fatal()
	}
}

func Test_C8_NewMust_Valid(t *testing.T) {
	v := issetter.NewMust("Set")
	if v != issetter.Set {
		t.Fatal()
	}
}

func Test_C8_RangeNamesCsv(t *testing.T) {
	csv := issetter.RangeNamesCsv()
	if csv == "" {
		t.Fatal()
	}
}

// ═══════════════════════════════════════════════
// Value methods — all uncovered from coverage.out
// ═══════════════════════════════════════════════

func Test_C8_AllNameValues(t *testing.T) {
	names := issetter.True.AllNameValues()
	if len(names) != 6 {
		t.Fatalf("expected 6, got %d", len(names))
	}
}

func Test_C8_OnlySupportedErr_NoNames(t *testing.T) {
	err := issetter.True.OnlySupportedErr()
	if err != nil {
		t.Fatal()
	}
}

func Test_C8_OnlySupportedErr_WithNames(t *testing.T) {
	err := issetter.True.OnlySupportedErr("True", "False")
	if err == nil {
		t.Fatal("expected error for unsupported names")
	}
}

func Test_C8_OnlySupportedErr_AllNames(t *testing.T) {
	err := issetter.True.OnlySupportedErr(
		"Uninitialized", "True", "False", "Unset", "Set", "Wildcard")
	if err != nil {
		t.Fatal("all names should be supported")
	}
}

func Test_C8_OnlySupportedMsgErr_NoError(t *testing.T) {
	err := issetter.True.OnlySupportedMsgErr("prefix: ")
	if err != nil {
		t.Fatal()
	}
}

func Test_C8_OnlySupportedMsgErr_WithError(t *testing.T) {
	err := issetter.True.OnlySupportedMsgErr("prefix: ", "True")
	if err == nil {
		t.Fatal()
	}
}

func Test_C8_ValueUInt16(t *testing.T) {
	if issetter.True.ValueUInt16() != 1 {
		t.Fatal()
	}
}

func Test_C8_IntegerEnumRanges(t *testing.T) {
	r := issetter.True.IntegerEnumRanges()
	if len(r) != 6 {
		t.Fatal()
	}
}

func Test_C8_MinMaxAny(t *testing.T) {
	min, max := issetter.True.MinMaxAny()
	if min == nil || max == nil {
		t.Fatal()
	}
}

func Test_C8_MinValueString(t *testing.T) {
	if issetter.True.MinValueString() == "" {
		t.Fatal()
	}
}

func Test_C8_MaxValueString(t *testing.T) {
	if issetter.True.MaxValueString() == "" {
		t.Fatal()
	}
}

func Test_C8_MaxInt(t *testing.T) {
	_ = issetter.True.MaxInt()
}

func Test_C8_MinInt(t *testing.T) {
	_ = issetter.True.MinInt()
}

func Test_C8_RangesDynamicMap(t *testing.T) {
	m := issetter.True.RangesDynamicMap()
	if len(m) == 0 {
		t.Fatal()
	}
}

func Test_C8_IsValueEqual(t *testing.T) {
	if !issetter.True.IsValueEqual(1) {
		t.Fatal()
	}
}

func Test_C8_RangeNamesCsvMethod(t *testing.T) {
	if issetter.True.RangeNamesCsv() == "" {
		t.Fatal()
	}
}

func Test_C8_IsByteValueEqual(t *testing.T) {
	if !issetter.True.IsByteValueEqual(1) {
		t.Fatal()
	}
}

func Test_C8_IsOn(t *testing.T) {
	if !issetter.True.IsOn() {
		t.Fatal()
	}
}

func Test_C8_IsOff(t *testing.T) {
	if !issetter.False.IsOff() {
		t.Fatal()
	}
}

func Test_C8_IsLater(t *testing.T) {
	if !issetter.Uninitialized.IsLater() {
		t.Fatal()
	}
}

func Test_C8_IsNot(t *testing.T) {
	if !issetter.True.IsNot(issetter.False) {
		t.Fatal()
	}
}

func Test_C8_IsNo(t *testing.T) {
	if !issetter.False.IsNo() {
		t.Fatal()
	}
}

func Test_C8_IsAsk(t *testing.T) {
	if !issetter.Uninitialized.IsAsk() {
		t.Fatal()
	}
}

func Test_C8_IsIndeterminate(t *testing.T) {
	if !issetter.Wildcard.IsIndeterminate() {
		t.Fatal()
	}
}

func Test_C8_IsAccept(t *testing.T) {
	if !issetter.Set.IsAccept() {
		t.Fatal()
	}
}

func Test_C8_IsReject(t *testing.T) {
	if !issetter.Unset.IsReject() {
		t.Fatal()
	}
}

func Test_C8_IsFailed(t *testing.T) {
	if !issetter.False.IsFailed() {
		t.Fatal()
	}
}

func Test_C8_IsSuccess(t *testing.T) {
	if !issetter.True.IsSuccess() {
		t.Fatal()
	}
}

func Test_C8_IsSkip(t *testing.T) {
	if !issetter.Wildcard.IsSkip() {
		t.Fatal()
	}
}

func Test_C8_NameValue(t *testing.T) {
	nv := issetter.True.NameValue()
	if nv == "" {
		t.Fatal()
	}
}

func Test_C8_IsNameEqual(t *testing.T) {
	if !issetter.True.IsNameEqual("True") {
		t.Fatal()
	}
}

func Test_C8_IsAnyNamesOf_Match(t *testing.T) {
	if !issetter.True.IsAnyNamesOf("True", "False") {
		t.Fatal()
	}
}

func Test_C8_IsAnyNamesOf_NoMatch(t *testing.T) {
	if issetter.True.IsAnyNamesOf("False", "Unset") {
		t.Fatal()
	}
}

func Test_C8_ToNumberString(t *testing.T) {
	if issetter.True.ToNumberString() != "1" {
		t.Fatal()
	}
}

func Test_C8_ValueByte(t *testing.T) {
	if issetter.True.ValueByte() != 1 {
		t.Fatal()
	}
}

func Test_C8_ValueInt(t *testing.T) {
	if issetter.True.ValueInt() != 1 {
		t.Fatal()
	}
}

func Test_C8_ValueInt8(t *testing.T) {
	if issetter.True.ValueInt8() != 1 {
		t.Fatal()
	}
}

func Test_C8_ValueInt16(t *testing.T) {
	if issetter.True.ValueInt16() != 1 {
		t.Fatal()
	}
}

func Test_C8_ValueInt32(t *testing.T) {
	if issetter.True.ValueInt32() != 1 {
		t.Fatal()
	}
}

func Test_C8_ValueString(t *testing.T) {
	if issetter.True.ValueString() != "1" {
		t.Fatal()
	}
}

func Test_C8_Format(t *testing.T) {
	s := issetter.True.Format("{name}={value}")
	if s != "True=1" {
		t.Fatalf("got %s", s)
	}
}

func Test_C8_EnumType(t *testing.T) {
	if issetter.True.EnumType() == nil {
		t.Fatal()
	}
}

func Test_C8_StringValue(t *testing.T) {
	if issetter.True.StringValue() != "1" {
		t.Fatal()
	}
}

func Test_C8_String(t *testing.T) {
	if issetter.True.String() != "True" {
		t.Fatal()
	}
}

func Test_C8_IsTrueOrSet(t *testing.T) {
	if !issetter.Set.IsTrueOrSet() {
		t.Fatal()
	}
}

func Test_C8_IsSet(t *testing.T) {
	if !issetter.Set.IsSet() {
		t.Fatal()
	}
}

func Test_C8_IsUnset(t *testing.T) {
	if !issetter.Unset.IsUnset() {
		t.Fatal()
	}
}

func Test_C8_HasInitialized(t *testing.T) {
	if !issetter.True.HasInitialized() {
		t.Fatal()
	}
}

func Test_C8_HasInitializedAndSet(t *testing.T) {
	if !issetter.Set.HasInitializedAndSet() {
		t.Fatal()
	}
}

func Test_C8_HasInitializedAndTrue(t *testing.T) {
	if !issetter.True.HasInitializedAndTrue() {
		t.Fatal()
	}
}

func Test_C8_IsWildcard(t *testing.T) {
	if !issetter.Wildcard.IsWildcard() {
		t.Fatal()
	}
}

func Test_C8_IsInit(t *testing.T) {
	if !issetter.True.IsInit() {
		t.Fatal()
	}
}

func Test_C8_IsInitBoolean(t *testing.T) {
	if !issetter.True.IsInitBoolean() {
		t.Fatal()
	}
	if !issetter.False.IsInitBoolean() {
		t.Fatal()
	}
}

func Test_C8_IsDefinedBoolean(t *testing.T) {
	if !issetter.True.IsDefinedBoolean() {
		t.Fatal()
	}
}

func Test_C8_IsInitBooleanWild(t *testing.T) {
	if !issetter.Wildcard.IsInitBooleanWild() {
		t.Fatal()
	}
}

func Test_C8_IsInitSet(t *testing.T) {
	if !issetter.Set.IsInitSet() {
		t.Fatal()
	}
	if !issetter.Unset.IsInitSet() {
		t.Fatal()
	}
}

func Test_C8_IsInitSetWild(t *testing.T) {
	if !issetter.Wildcard.IsInitSetWild() {
		t.Fatal()
	}
}

func Test_C8_IsYes(t *testing.T) {
	if !issetter.True.IsYes() {
		t.Fatal()
	}
}

func Test_C8_Boolean(t *testing.T) {
	if !issetter.True.Boolean() {
		t.Fatal()
	}
}

func Test_C8_IsOnLogically(t *testing.T) {
	if !issetter.True.IsOnLogically() {
		t.Fatal()
	}
}

func Test_C8_IsOffLogically(t *testing.T) {
	if !issetter.False.IsOffLogically() {
		t.Fatal()
	}
}

func Test_C8_IsAccepted(t *testing.T) {
	if !issetter.True.IsAccepted() {
		t.Fatal()
	}
}

func Test_C8_IsRejected(t *testing.T) {
	if !issetter.False.IsRejected() {
		t.Fatal()
	}
}

func Test_C8_IsDefinedLogically(t *testing.T) {
	if !issetter.True.IsDefinedLogically() {
		t.Fatal()
	}
	if issetter.Uninitialized.IsDefinedLogically() {
		t.Fatal()
	}
}

func Test_C8_IsUndefinedLogically(t *testing.T) {
	if !issetter.Wildcard.IsUndefinedLogically() {
		t.Fatal()
	}
}

func Test_C8_IsInvalid(t *testing.T) {
	if !issetter.Uninitialized.IsInvalid() {
		t.Fatal()
	}
}

func Test_C8_IsValid(t *testing.T) {
	if !issetter.True.IsValid() {
		t.Fatal()
	}
}

func Test_C8_GetSetBoolOnInvalid_AlreadyDefined(t *testing.T) {
	v := issetter.True
	r := v.GetSetBoolOnInvalid(false)
	if !r {
		t.Fatal()
	}
}

func Test_C8_GetSetBoolOnInvalid_SetNew(t *testing.T) {
	v := issetter.Uninitialized
	r := v.GetSetBoolOnInvalid(true)
	if !r {
		t.Fatal()
	}
}

func Test_C8_GetSetBoolOnInvalidFunc_AlreadyDefined(t *testing.T) {
	v := issetter.False
	r := v.GetSetBoolOnInvalidFunc(func() bool { return true })
	if r {
		t.Fatal()
	}
}

func Test_C8_GetSetBoolOnInvalidFunc_SetNew(t *testing.T) {
	v := issetter.Uninitialized
	r := v.GetSetBoolOnInvalidFunc(func() bool { return true })
	if !r {
		t.Fatal()
	}
}

func Test_C8_ToBooleanValue(t *testing.T) {
	if issetter.Set.ToBooleanValue() != issetter.True {
		t.Fatal()
	}
}

func Test_C8_ToSetUnsetValue(t *testing.T) {
	if issetter.True.ToSetUnsetValue() != issetter.Set {
		t.Fatal()
	}
}

func Test_C8_LazyEvaluateBool_NotCalled(t *testing.T) {
	v := issetter.True
	called := v.LazyEvaluateBool(func() { t.Fatal("should not call") })
	if called {
		t.Fatal()
	}
}

func Test_C8_LazyEvaluateBool_Called(t *testing.T) {
	v := issetter.Uninitialized
	executed := false
	called := v.LazyEvaluateBool(func() { executed = true })
	if !called || !executed {
		t.Fatal()
	}
}

func Test_C8_LazyEvaluateSet_NotCalled(t *testing.T) {
	v := issetter.Set
	called := v.LazyEvaluateSet(func() { t.Fatal("should not call") })
	if called {
		t.Fatal()
	}
}

func Test_C8_LazyEvaluateSet_Called(t *testing.T) {
	v := issetter.Uninitialized
	executed := false
	called := v.LazyEvaluateSet(func() { executed = true })
	if !called || !executed {
		t.Fatal()
	}
}

func Test_C8_IsWildcardOrBool_Wildcard(t *testing.T) {
	if !issetter.Wildcard.IsWildcardOrBool(false) {
		t.Fatal()
	}
}

func Test_C8_IsWildcardOrBool_TrueTrue(t *testing.T) {
	if !issetter.True.IsWildcardOrBool(true) {
		t.Fatal()
	}
}

func Test_C8_IsWildcardOrBool_FalseFalse(t *testing.T) {
	if !issetter.False.IsWildcardOrBool(false) {
		t.Fatal()
	}
}

func Test_C8_ToByteCondition(t *testing.T) {
	if issetter.True.ToByteCondition(10, 20, 30) != 10 {
		t.Fatal()
	}
	if issetter.False.ToByteCondition(10, 20, 30) != 20 {
		t.Fatal()
	}
	if issetter.Uninitialized.ToByteCondition(10, 20, 30) != 30 {
		t.Fatal()
	}
}

func Test_C8_ToByteConditionWithWildcard(t *testing.T) {
	if issetter.Wildcard.ToByteConditionWithWildcard(99, 10, 20, 30) != 99 {
		t.Fatal()
	}
	if issetter.True.ToByteConditionWithWildcard(99, 10, 20, 30) != 10 {
		t.Fatal()
	}
}

func Test_C8_WildcardApply_Wildcard(t *testing.T) {
	if !issetter.Wildcard.WildcardApply(true) {
		t.Fatal()
	}
}

func Test_C8_WildcardApply_Defined(t *testing.T) {
	if !issetter.True.WildcardApply(false) {
		t.Fatal()
	}
}

func Test_C8_WildcardValueApply(t *testing.T) {
	if !issetter.Wildcard.WildcardValueApply(issetter.True) {
		t.Fatal()
	}
	if !issetter.True.WildcardValueApply(issetter.False) {
		t.Fatal()
	}
}

func Test_C8_OrBool(t *testing.T) {
	if !issetter.Wildcard.OrBool(true) {
		t.Fatal()
	}
	if !issetter.True.OrBool(false) {
		t.Fatal()
	}
}

func Test_C8_OrValue(t *testing.T) {
	if !issetter.Wildcard.OrValue(issetter.True) {
		t.Fatal()
	}
	if !issetter.True.OrValue(issetter.False) {
		t.Fatal()
	}
}

func Test_C8_AndBool(t *testing.T) {
	if !issetter.Wildcard.AndBool(true) {
		t.Fatal()
	}
	if issetter.True.AndBool(false) {
		t.Fatal()
	}
}

func Test_C8_And(t *testing.T) {
	r := issetter.Wildcard.And(issetter.True)
	if r != issetter.True {
		t.Fatal()
	}
	r2 := issetter.True.And(issetter.False)
	if r2 != issetter.False {
		t.Fatal()
	}
}

func Test_C8_IsUninitialized(t *testing.T) {
	if !issetter.Uninitialized.IsUninitialized() {
		t.Fatal()
	}
}

func Test_C8_IsInitialized(t *testing.T) {
	if !issetter.True.IsInitialized() {
		t.Fatal()
	}
}

func Test_C8_IsUnSetOrUninitialized(t *testing.T) {
	if !issetter.Unset.IsUnSetOrUninitialized() {
		t.Fatal()
	}
}

func Test_C8_IsNegative(t *testing.T) {
	if !issetter.Uninitialized.IsNegative() {
		t.Fatal()
	}
	if !issetter.False.IsNegative() {
		t.Fatal()
	}
}

func Test_C8_IsPositive(t *testing.T) {
	if !issetter.True.IsPositive() {
		t.Fatal()
	}
	if !issetter.Set.IsPositive() {
		t.Fatal()
	}
}

func Test_C8_IsBetween(t *testing.T) {
	if !issetter.True.IsBetween(0, 5) {
		t.Fatal()
	}
}

func Test_C8_IsBetweenInt(t *testing.T) {
	if !issetter.True.IsBetweenInt(0, 5) {
		t.Fatal()
	}
}

func Test_C8_Add(t *testing.T) {
	r := issetter.True.Add(1)
	if r != issetter.False {
		t.Fatal()
	}
}

func Test_C8_Is(t *testing.T) {
	if !issetter.True.Is(issetter.True) {
		t.Fatal()
	}
}

func Test_C8_IsEqual(t *testing.T) {
	if !issetter.True.IsEqual(1) {
		t.Fatal()
	}
}

func Test_C8_IsGreater(t *testing.T) {
	if !issetter.Set.IsGreater(1) {
		t.Fatal()
	}
}

func Test_C8_IsGreaterEqual(t *testing.T) {
	if !issetter.True.IsGreaterEqual(1) {
		t.Fatal()
	}
}

func Test_C8_IsLess(t *testing.T) {
	if !issetter.True.IsLess(4) {
		t.Fatal()
	}
}

func Test_C8_IsLessEqual(t *testing.T) {
	if !issetter.True.IsLessEqual(1) {
		t.Fatal()
	}
}

func Test_C8_IsEqualInt(t *testing.T) {
	if !issetter.True.IsEqualInt(1) {
		t.Fatal()
	}
}

func Test_C8_IsGreaterInt(t *testing.T) {
	if !issetter.Set.IsGreaterInt(1) {
		t.Fatal()
	}
}

func Test_C8_IsGreaterEqualInt(t *testing.T) {
	if !issetter.True.IsGreaterEqualInt(1) {
		t.Fatal()
	}
}

func Test_C8_IsLessInt(t *testing.T) {
	if !issetter.True.IsLessInt(4) {
		t.Fatal()
	}
}

func Test_C8_IsLessEqualInt(t *testing.T) {
	if !issetter.True.IsLessEqualInt(1) {
		t.Fatal()
	}
}

func Test_C8_PanicOnOutOfRange_InRange(t *testing.T) {
	issetter.True.PanicOnOutOfRange(1, "msg")
}

func Test_C8_PanicOnOutOfRange_OutOfRange(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	issetter.True.PanicOnOutOfRange(200, "out of range")
}

func Test_C8_GetErrorOnOutOfRange_InRange(t *testing.T) {
	err := issetter.True.GetErrorOnOutOfRange(1, "msg")
	if err != nil {
		t.Fatal()
	}
}

func Test_C8_GetErrorOnOutOfRange_OutOfRange(t *testing.T) {
	err := issetter.True.GetErrorOnOutOfRange(200, "msg")
	if err == nil {
		t.Fatal()
	}
}

func Test_C8_Name(t *testing.T) {
	if issetter.True.Name() != "True" {
		t.Fatal()
	}
}

func Test_C8_YesNoMappedValue_Uninit(t *testing.T) {
	if issetter.Uninitialized.YesNoMappedValue() != "" {
		t.Fatal()
	}
}

func Test_C8_YesNoMappedValue_True(t *testing.T) {
	if issetter.True.YesNoMappedValue() != "yes" {
		t.Fatal()
	}
}

func Test_C8_YesNoMappedValue_False(t *testing.T) {
	if issetter.False.YesNoMappedValue() != "no" {
		t.Fatal()
	}
}

func Test_C8_YesNoLowercaseName(t *testing.T) {
	if issetter.True.YesNoLowercaseName() != "yes" {
		t.Fatal()
	}
}

func Test_C8_YesNoName(t *testing.T) {
	if issetter.True.YesNoName() != "Yes" {
		t.Fatal()
	}
}

func Test_C8_TrueFalseName(t *testing.T) {
	if issetter.True.TrueFalseName() != "True" {
		t.Fatal()
	}
}

func Test_C8_OnOffLowercaseName(t *testing.T) {
	if issetter.True.OnOffLowercaseName() != "on" {
		t.Fatal()
	}
}

func Test_C8_OnOffName(t *testing.T) {
	if issetter.True.OnOffName() != "On" {
		t.Fatal()
	}
}

func Test_C8_TrueFalseLowercaseName(t *testing.T) {
	if issetter.True.TrueFalseLowercaseName() != "true" {
		t.Fatal()
	}
}

func Test_C8_SetUnsetLowercaseName(t *testing.T) {
	if issetter.True.SetUnsetLowercaseName() != "set" {
		t.Fatal()
	}
}

func Test_C8_MarshalJSON(t *testing.T) {
	b, err := issetter.True.MarshalJSON()
	if err != nil || len(b) == 0 {
		t.Fatal()
	}
}

func Test_C8_UnmarshalJSON_Valid(t *testing.T) {
	v := issetter.Uninitialized
	err := v.UnmarshalJSON([]byte(`"True"`))
	if err != nil || v != issetter.True {
		t.Fatal()
	}
}

func Test_C8_UnmarshalJSON_Nil(t *testing.T) {
	v := issetter.Uninitialized
	err := v.UnmarshalJSON(nil)
	if err == nil {
		t.Fatal()
	}
}

func Test_C8_UnmarshalJSON_Invalid(t *testing.T) {
	v := issetter.Uninitialized
	err := v.UnmarshalJSON([]byte(`"BOGUS"`))
	if err == nil {
		t.Fatal()
	}
}

func Test_C8_Serialize(t *testing.T) {
	b, err := issetter.True.Serialize()
	if err != nil || len(b) == 0 {
		t.Fatal()
	}
}

func Test_C8_TypeName(t *testing.T) {
	if issetter.True.TypeName() == "" {
		t.Fatal()
	}
}

func Test_C8_IsAnyValuesEqual_Match(t *testing.T) {
	if !issetter.True.IsAnyValuesEqual(0, 1, 2) {
		t.Fatal()
	}
}

func Test_C8_IsAnyValuesEqual_NoMatch(t *testing.T) {
	if issetter.True.IsAnyValuesEqual(0, 2, 3) {
		t.Fatal()
	}
}

func Test_C8_UnmarshallEnumToValue(t *testing.T) {
	v, err := issetter.Uninitialized.UnmarshallEnumToValue([]byte(`"Set"`))
	if err != nil || v != issetter.Set.Value() {
		t.Fatal()
	}
}

func Test_C8_Deserialize_Valid(t *testing.T) {
	v, err := issetter.Uninitialized.Deserialize([]byte(`"True"`))
	if err != nil || v != issetter.True {
		t.Fatal()
	}
}

func Test_C8_Deserialize_Invalid(t *testing.T) {
	_, err := issetter.Uninitialized.Deserialize([]byte(`"BOGUS"`))
	if err == nil {
		t.Fatal()
	}
}

func Test_C8_MaxByteMethod(t *testing.T) {
	if issetter.True.MaxByte() != issetter.Wildcard.ValueByte() {
		t.Fatal()
	}
}

func Test_C8_MinByteMethod(t *testing.T) {
	if issetter.True.MinByte() != issetter.Uninitialized.ValueByte() {
		t.Fatal()
	}
}

func Test_C8_RangesByte_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	issetter.True.RangesByte()
}

func Test_C8_ToPtr(t *testing.T) {
	p := issetter.True.ToPtr()
	if p == nil || *p != issetter.True {
		t.Fatal()
	}
}
