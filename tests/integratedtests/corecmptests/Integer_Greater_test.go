package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// Integer — Greater branch
// ═══════════════════════════════════════════

func Test_Integer_Greater_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer(10, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- greater", actual)
}

// ═══════════════════════════════════════════
// Integer8 — Greater branch
// ═══════════════════════════════════════════

func Test_Integer8_Greater_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8(10, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- greater", actual)
}

func Test_Integer8Ptr_LeftNil_FromIntegerGreater(t *testing.T) {
	// Arrange
	r := int8(5)

	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- left nil", actual)
}

func Test_Integer8Ptr_Equal_FromIntegerGreater(t *testing.T) {
	// Arrange
	l, r := int8(5), int8(5)

	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns correct value -- equal", actual)
}

// ═══════════════════════════════════════════
// Integer16 — all branches
// ═══════════════════════════════════════════

func Test_Integer16_Equal_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16(5, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- equal", actual)
}

func Test_Integer16_Less_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16(3, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- less", actual)
}

func Test_Integer16_Greater_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16(10, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- greater", actual)
}

func Test_Integer16Ptr_LeftNil_FromIntegerGreater(t *testing.T) {
	// Arrange
	r := int16(5)

	// Act
	actual := args.Map{"result": corecmp.Integer16Ptr(nil, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- left nil", actual)
}

func Test_Integer16Ptr_Equal_FromIntegerGreater(t *testing.T) {
	// Arrange
	l, r := int16(5), int16(5)

	// Act
	actual := args.Map{"result": corecmp.Integer16Ptr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns correct value -- equal", actual)
}

// ═══════════════════════════════════════════
// Integer32 — all branches
// ═══════════════════════════════════════════

func Test_Integer32_Equal_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32(5, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- equal", actual)
}

func Test_Integer32_Less_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32(3, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- less", actual)
}

func Test_Integer32_Greater_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32(10, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- greater", actual)
}

func Test_Integer32Ptr_LeftNil_FromIntegerGreater(t *testing.T) {
	// Arrange
	r := int32(5)

	// Act
	actual := args.Map{"result": corecmp.Integer32Ptr(nil, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- left nil", actual)
}

func Test_Integer32Ptr_Equal_FromIntegerGreater(t *testing.T) {
	// Arrange
	l, r := int32(5), int32(5)

	// Act
	actual := args.Map{"result": corecmp.Integer32Ptr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns correct value -- equal", actual)
}

// ═══════════════════════════════════════════
// Integer64 — Greater branch
// ═══════════════════════════════════════════

func Test_Integer64_Greater_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64(10, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- greater", actual)
}

func Test_Integer64Ptr_LeftNil_FromIntegerGreater(t *testing.T) {
	// Arrange
	r := int64(5)

	// Act
	actual := args.Map{"result": corecmp.Integer64Ptr(nil, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- left nil", actual)
}

func Test_Integer64Ptr_Equal_FromIntegerGreater(t *testing.T) {
	// Arrange
	l, r := int64(5), int64(5)

	// Act
	actual := args.Map{"result": corecmp.Integer64Ptr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns correct value -- equal", actual)
}

// ═══════════════════════════════════════════
// BytePtr — remaining branches
// ═══════════════════════════════════════════

func Test_BytePtr_RightNil_FromIntegerGreater(t *testing.T) {
	// Arrange
	l := byte(5)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(&l, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- right nil", actual)
}

func Test_BytePtr_Equal_FromIntegerGreater(t *testing.T) {
	// Arrange
	l, r := byte(5), byte(5)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "BytePtr returns correct value -- equal", actual)
}

func Test_BytePtr_Less(t *testing.T) {
	// Arrange
	l, r := byte(3), byte(5)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "BytePtr returns correct value -- less", actual)
}

func Test_BytePtr_Greater(t *testing.T) {
	// Arrange
	l, r := byte(10), byte(5)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "BytePtr returns correct value -- greater", actual)
}

// ═══════════════════════════════════════════
// AnyItem — Inconclusive
// ═══════════════════════════════════════════

func Test_AnyItem_RightNil_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem(5, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- right nil", actual)
}

func Test_AnyItem_Inconclusive_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem(5, 10)}

	// Assert
	expected := args.Map{"result": corecomparator.Inconclusive}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- inconclusive", actual)
}

// ═══════════════════════════════════════════
// IsStringsEqual — NotEqual items
// ═══════════════════════════════════════════

func Test_IsStringsEqual_NotEqualItems(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "c"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- not equal items", actual)
}

func Test_IsStringsEqual_RightNil_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a"}, nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- right nil", actual)
}

// ═══════════════════════════════════════════
// IsStringsEqualPtr — DiffLen, RightNil
// ═══════════════════════════════════════════

func Test_IsStringsEqualPtr_RightNil_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualPtr([]string{"a"}, nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns nil -- right nil", actual)
}

func Test_IsStringsEqualPtr_DiffLen_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns correct value -- diff len", actual)
}

// ═══════════════════════════════════════════
// IsStringsEqualWithoutOrder — all branches
// ═══════════════════════════════════════════

func Test_IsStringsEqualWithoutOrder_BothNil_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns nil -- both nil", actual)
}

func Test_IsStringsEqualWithoutOrder_LeftNil_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns nil -- left nil", actual)
}

func Test_IsStringsEqualWithoutOrder_RightNil_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns nil -- right nil", actual)
}

func Test_IsStringsEqualWithoutOrder_DiffLen_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- diff len", actual)
}

func Test_IsStringsEqualWithoutOrder_Equal_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"b", "a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- equal", actual)
}

func Test_IsStringsEqualWithoutOrder_NotEqual_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a", "b"}, []string{"a", "c"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- not equal", actual)
}

// ═══════════════════════════════════════════
// IsIntegersEqual — LeftNil
// ═══════════════════════════════════════════

func Test_IsIntegersEqual_LeftNil_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqual(nil, []int{1})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns nil -- left nil", actual)
}

func Test_IsIntegersEqual_RightNil_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqual([]int{1}, nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns nil -- right nil", actual)
}

// ═══════════════════════════════════════════
// IsIntegersEqualPtr — RightNil, DiffLen
// ═══════════════════════════════════════════

func Test_IsIntegersEqualPtr_RightNil_FromIntegerGreater(t *testing.T) {
	// Arrange
	l := []int{1}

	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&l, nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- right nil", actual)
}

func Test_IsIntegersEqualPtr_DiffLen_FromIntegerGreater(t *testing.T) {
	// Arrange
	l := []int{1}
	r := []int{1, 2}

	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&l, &r)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns correct value -- diff len", actual)
}

// ═══════════════════════════════════════════
// VersionSliceByte — all branches
// ═══════════════════════════════════════════

func Test_VersionSliceByte_BothNil_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns nil -- both nil", actual)
}

func Test_VersionSliceByte_LeftNil_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte(nil, []byte{1})}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns nil -- left nil", actual)
}

func Test_VersionSliceByte_Equal_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- equal", actual)
}

func Test_VersionSliceByte_LeftLess_SameLen(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2, 4})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- left less same len", actual)
}

func Test_VersionSliceByte_LeftGreater_SameLen(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2, 4}, []byte{1, 2, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- left greater same len", actual)
}

func Test_VersionSliceByte_LeftLess_DiffLen(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 2, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- left less diff len", actual)
}

func Test_VersionSliceByte_LeftGreater_DiffLen(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- left greater diff len", actual)
}

// ═══════════════════════════════════════════
// VersionSliceInteger — all branches
// ═══════════════════════════════════════════

func Test_VersionSliceInteger_BothNil_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- both nil", actual)
}

func Test_VersionSliceInteger_LeftNil_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger(nil, []int{1})}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- left nil", actual)
}

func Test_VersionSliceInteger_Equal_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- equal", actual)
}

func Test_VersionSliceInteger_LeftLess_SameLen(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2, 4})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- left less same len", actual)
}

func Test_VersionSliceInteger_LeftGreater_SameLen(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2, 4}, []int{1, 2, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- left greater same len", actual)
}

func Test_VersionSliceInteger_LeftLess_DiffLen(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 2, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- left less diff len", actual)
}

func Test_VersionSliceInteger_LeftGreater_DiffLen(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- left greater diff len", actual)
}

// ═══════════════════════════════════════════
// Time — all branches
// ═══════════════════════════════════════════

func Test_Time_Equal_FromIntegerGreater(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"result": corecmp.Time(now, now)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- equal", actual)
}

func Test_Time_Less_FromIntegerGreater(t *testing.T) {
	// Arrange
	now := time.Now()
	later := now.Add(time.Hour)

	// Act
	actual := args.Map{"result": corecmp.Time(now, later)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- less", actual)
}

func Test_Time_Greater_FromIntegerGreater(t *testing.T) {
	// Arrange
	now := time.Now()
	earlier := now.Add(-time.Hour)

	// Act
	actual := args.Map{"result": corecmp.Time(now, earlier)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- greater", actual)
}

// ═══════════════════════════════════════════
// TimePtr — all branches
// ═══════════════════════════════════════════

func Test_TimePtr_BothNil_FromIntegerGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.TimePtr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- both nil", actual)
}

func Test_TimePtr_LeftNil_FromIntegerGreater(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"result": corecmp.TimePtr(nil, &now)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- left nil", actual)
}

func Test_TimePtr_RightNil_FromIntegerGreater(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"result": corecmp.TimePtr(&now, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- right nil", actual)
}

func Test_TimePtr_Equal_FromIntegerGreater(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"result": corecmp.TimePtr(&now, &now)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "TimePtr returns correct value -- equal", actual)
}

func Test_TimePtr_Less(t *testing.T) {
	// Arrange
	now := time.Now()
	later := now.Add(time.Hour)

	// Act
	actual := args.Map{"result": corecmp.TimePtr(&now, &later)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "TimePtr returns correct value -- less", actual)
}

func Test_TimePtr_Greater(t *testing.T) {
	// Arrange
	now := time.Now()
	earlier := now.Add(-time.Hour)

	// Act
	actual := args.Map{"result": corecmp.TimePtr(&now, &earlier)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "TimePtr returns correct value -- greater", actual)
}
