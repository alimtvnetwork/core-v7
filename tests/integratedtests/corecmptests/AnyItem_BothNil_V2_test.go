package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// AnyItem — all 4 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_AnyItem_BothNil(t *testing.T) {
	// Arrange
	r := corecmp.AnyItem(nil, nil)

	// Act
	actual := args.Map{"v": r}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- both nil → Equal", actual)
}

func Test_AnyItem_LeftNilOnly(t *testing.T) {
	// Arrange
	r := corecmp.AnyItem(nil, 1)

	// Act
	actual := args.Map{"v": r}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- left nil → NotEqual", actual)
}

func Test_AnyItem_RightNilOnly(t *testing.T) {
	// Arrange
	r := corecmp.AnyItem(1, nil)

	// Act
	actual := args.Map{"v": r}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- right nil → NotEqual", actual)
}

func Test_AnyItem_SameValue(t *testing.T) {
	// Arrange
	r := corecmp.AnyItem(42, 42)

	// Act
	actual := args.Map{"v": r}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- same → Equal", actual)
}

func Test_AnyItem_DiffValue(t *testing.T) {
	// Arrange
	r := corecmp.AnyItem(1, 2)

	// Act
	actual := args.Map{"v": r}

	// Assert
	expected := args.Map{"v": corecomparator.Inconclusive}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- diff → Inconclusive", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Byte — 3 branches: Equal, LeftLess, LeftGreater
// ══════════════════════════════════════════════════════════════════════════════

func Test_Byte_Equal(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Byte(5, 5)}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- equal", actual)
}

func Test_Byte_Less(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Byte(1, 9)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- less", actual)
}

func Test_Byte_Greater(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Byte(9, 1)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- greater", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BytePtr — 4 branches: BothNil, LeftNil, RightNil, Delegate
// ══════════════════════════════════════════════════════════════════════════════

func Test_BytePtr_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.BytePtr(nil, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- both nil", actual)
}

func Test_BytePtr_LeftNil(t *testing.T) {
	// Arrange
	b := byte(1)

	// Act
	actual := args.Map{"v": corecmp.BytePtr(nil, &b)}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- left nil", actual)
}

func Test_BytePtr_RightNil(t *testing.T) {
	// Arrange
	b := byte(1)

	// Act
	actual := args.Map{"v": corecmp.BytePtr(&b, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- right nil", actual)
}

func Test_BytePtr_Delegate(t *testing.T) {
	// Arrange
	a, b := byte(3), byte(7)

	// Act
	actual := args.Map{"v": corecmp.BytePtr(&a, &b)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "BytePtr returns correct value -- delegate", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer — 3 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Integer_Equal(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer(10, 10)}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- equal", actual)
}

func Test_Integer_Less(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer(-5, 5)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- less", actual)
}

func Test_Integer_Greater(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer(5, -5)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- greater", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IntegerPtr — 4 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_IntegerPtr_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IntegerPtr(nil, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- both nil", actual)
}

func Test_IntegerPtr_LeftNil(t *testing.T) {
	// Arrange
	v := 1

	// Act
	actual := args.Map{"v": corecmp.IntegerPtr(nil, &v)}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- left nil", actual)
}

func Test_IntegerPtr_RightNil(t *testing.T) {
	// Arrange
	v := 1

	// Act
	actual := args.Map{"v": corecmp.IntegerPtr(&v, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- right nil", actual)
}

func Test_IntegerPtr_Delegate(t *testing.T) {
	// Arrange
	a, b := 10, 20

	// Act
	actual := args.Map{"v": corecmp.IntegerPtr(&a, &b)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns correct value -- delegate", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer8 — 3 branches + Integer8Ptr — 4 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Integer8_Equal(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer8(5, 5)}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- equal", actual)
}

func Test_Integer8_Less(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer8(-10, 10)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- less", actual)
}

func Test_Integer8_Greater(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer8(10, -10)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- greater", actual)
}

func Test_Integer8Ptr_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer8Ptr(nil, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- both nil", actual)
}

func Test_Integer8Ptr_LeftNil(t *testing.T) {
	// Arrange
	v := int8(1)

	// Act
	actual := args.Map{"v": corecmp.Integer8Ptr(nil, &v)}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- left nil", actual)
}

func Test_Integer8Ptr_RightNil(t *testing.T) {
	// Arrange
	v := int8(1)

	// Act
	actual := args.Map{"v": corecmp.Integer8Ptr(&v, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- right nil", actual)
}

func Test_Integer8Ptr_Delegate(t *testing.T) {
	// Arrange
	a, b := int8(3), int8(7)

	// Act
	actual := args.Map{"v": corecmp.Integer8Ptr(&a, &b)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns correct value -- delegate", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer16 — 3 branches + Integer16Ptr — 4 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Integer16_Equal(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer16(100, 100)}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- equal", actual)
}

func Test_Integer16_Less(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer16(-100, 100)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- less", actual)
}

func Test_Integer16_Greater(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer16(100, -100)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- greater", actual)
}

func Test_Integer16Ptr_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer16Ptr(nil, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- both nil", actual)
}

func Test_Integer16Ptr_LeftNil(t *testing.T) {
	// Arrange
	v := int16(1)

	// Act
	actual := args.Map{"v": corecmp.Integer16Ptr(nil, &v)}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- left nil", actual)
}

func Test_Integer16Ptr_RightNil(t *testing.T) {
	// Arrange
	v := int16(1)

	// Act
	actual := args.Map{"v": corecmp.Integer16Ptr(&v, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- right nil", actual)
}

func Test_Integer16Ptr_Delegate(t *testing.T) {
	// Arrange
	a, b := int16(3), int16(7)

	// Act
	actual := args.Map{"v": corecmp.Integer16Ptr(&a, &b)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns correct value -- delegate", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer32 — 3 branches + Integer32Ptr — 4 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Integer32_Equal(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer32(1000, 1000)}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- equal", actual)
}

func Test_Integer32_Less(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer32(-1000, 1000)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- less", actual)
}

func Test_Integer32_Greater(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer32(1000, -1000)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- greater", actual)
}

func Test_Integer32Ptr_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer32Ptr(nil, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- both nil", actual)
}

func Test_Integer32Ptr_LeftNil(t *testing.T) {
	// Arrange
	v := int32(1)

	// Act
	actual := args.Map{"v": corecmp.Integer32Ptr(nil, &v)}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- left nil", actual)
}

func Test_Integer32Ptr_RightNil(t *testing.T) {
	// Arrange
	v := int32(1)

	// Act
	actual := args.Map{"v": corecmp.Integer32Ptr(&v, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- right nil", actual)
}

func Test_Integer32Ptr_Delegate(t *testing.T) {
	// Arrange
	a, b := int32(3), int32(7)

	// Act
	actual := args.Map{"v": corecmp.Integer32Ptr(&a, &b)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns correct value -- delegate", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer64 — 3 branches + Integer64Ptr — 4 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Integer64_Equal(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer64(99999, 99999)}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- equal", actual)
}

func Test_Integer64_Less(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer64(-99999, 99999)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- less", actual)
}

func Test_Integer64_Greater(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer64(99999, -99999)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- greater", actual)
}

func Test_Integer64Ptr_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.Integer64Ptr(nil, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- both nil", actual)
}

func Test_Integer64Ptr_LeftNil(t *testing.T) {
	// Arrange
	v := int64(1)

	// Act
	actual := args.Map{"v": corecmp.Integer64Ptr(nil, &v)}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- left nil", actual)
}

func Test_Integer64Ptr_RightNil(t *testing.T) {
	// Arrange
	v := int64(1)

	// Act
	actual := args.Map{"v": corecmp.Integer64Ptr(&v, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- right nil", actual)
}

func Test_Integer64Ptr_Delegate(t *testing.T) {
	// Arrange
	a, b := int64(3), int64(7)

	// Act
	actual := args.Map{"v": corecmp.Integer64Ptr(&a, &b)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns correct value -- delegate", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsIntegersEqual — 4 branches (nil/nil, nil/val, val/nil, delegate)
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsIntegersEqual_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsIntegersEqual(nil, nil)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns nil -- both nil", actual)
}

func Test_IsIntegersEqual_LeftNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsIntegersEqual(nil, []int{1})}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns nil -- left nil", actual)
}

func Test_IsIntegersEqual_RightNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsIntegersEqual([]int{1}, nil)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns nil -- right nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsIntegersEqualPtr — 4 branches (nil/nil, nil/val, diffLen, delegate)
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsIntegersEqualPtr_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsIntegersEqualPtr(nil, nil)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- both nil", actual)
}

func Test_IsIntegersEqualPtr_LeftNil(t *testing.T) {
	// Arrange
	r := []int{1}

	// Act
	actual := args.Map{"v": corecmp.IsIntegersEqualPtr(nil, &r)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- left nil", actual)
}

func Test_IsIntegersEqualPtr_RightNil(t *testing.T) {
	// Arrange
	l := []int{1}

	// Act
	actual := args.Map{"v": corecmp.IsIntegersEqualPtr(&l, nil)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- right nil", actual)
}

func Test_IsIntegersEqualPtr_DiffLen(t *testing.T) {
	// Arrange
	l := []int{1}
	r := []int{1, 2}

	// Act
	actual := args.Map{"v": corecmp.IsIntegersEqualPtr(&l, &r)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns correct value -- diff len", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsStringsEqual — 5 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsStringsEqual_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqual(nil, nil)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- both nil", actual)
}

func Test_IsStringsEqual_LeftNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqual(nil, []string{"a"})}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- left nil", actual)
}

func Test_IsStringsEqual_RightNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqual([]string{"a"}, nil)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- right nil", actual)
}

func Test_IsStringsEqual_DiffLen(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqual([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- diff len", actual)
}

func Test_IsStringsEqual_Same(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- same", actual)
}

func Test_IsStringsEqual_Diff(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "c"})}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- diff", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsStringsEqualPtr — 4 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsStringsEqualPtr_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqualPtr(nil, nil)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns nil -- both nil", actual)
}

func Test_IsStringsEqualPtr_LeftNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqualPtr(nil, []string{"a"})}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns nil -- left nil", actual)
}

func Test_IsStringsEqualPtr_RightNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqualPtr([]string{"a"}, nil)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns nil -- right nil", actual)
}

func Test_IsStringsEqualPtr_DiffLen(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns correct value -- diff len", actual)
}

func Test_IsStringsEqualPtr_Same(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqualPtr([]string{"x", "y"}, []string{"x", "y"})}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns correct value -- same", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsStringsEqualWithoutOrder — 5 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsStringsEqualWithoutOrder_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqualWithoutOrder(nil, nil)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "WithoutOrder returns nil -- both nil", actual)
}

func Test_IsStringsEqualWithoutOrder_LeftNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"})}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "WithoutOrder returns nil -- left nil", actual)
}

func Test_IsStringsEqualWithoutOrder_RightNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, nil)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "WithoutOrder returns nil -- right nil", actual)
}

func Test_IsStringsEqualWithoutOrder_DiffLen(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "WithoutOrder returns non-empty -- diff len", actual)
}

func Test_IsStringsEqualWithoutOrder_Reordered(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqualWithoutOrder([]string{"b", "a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "WithoutOrder returns non-empty -- reordered", actual)
}

func Test_IsStringsEqualWithoutOrder_Mismatch(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.IsStringsEqualWithoutOrder([]string{"a", "b"}, []string{"a", "c"})}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "WithoutOrder returns non-empty -- mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Time — 3 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Time_Equal(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"v": corecmp.Time(now, now)}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- equal", actual)
}

func Test_Time_Less(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"v": corecmp.Time(now, now.Add(time.Hour))}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- less", actual)
}

func Test_Time_Greater(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"v": corecmp.Time(now, now.Add(-time.Hour))}

	// Assert
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- greater", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TimePtr — 4 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_TimePtr_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.TimePtr(nil, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- both nil", actual)
}

func Test_TimePtr_LeftNil(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"v": corecmp.TimePtr(nil, &now)}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- left nil", actual)
}

func Test_TimePtr_RightNil(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"v": corecmp.TimePtr(&now, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- right nil", actual)
}

func Test_TimePtr_Delegate(t *testing.T) {
	// Arrange
	now := time.Now()
	later := now.Add(time.Hour)

	// Act
	actual := args.Map{"v": corecmp.TimePtr(&now, &later)}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "TimePtr returns correct value -- delegate", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// VersionSliceByte — 7 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_VersionSliceByte_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.VersionSliceByte(nil, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VSB returns nil -- both nil", actual)
}

func Test_VersionSliceByte_LeftNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.VersionSliceByte(nil, []byte{1})}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VSB returns nil -- left nil", actual)
}

func Test_VersionSliceByte_RightNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.VersionSliceByte([]byte{1}, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VSB returns nil -- right nil", actual)
}

func Test_VersionSliceByte_EqualSameLen(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2, 3})}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VSB returns correct value -- equal same len", actual)
}

func Test_VersionSliceByte_LoopLeftLess(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.VersionSliceByte([]byte{1, 0, 0}, []byte{1, 0, 1})}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VSB returns correct value -- loop left less", actual)
}

func Test_VersionSliceByte_LoopLeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.VersionSliceByte([]byte{2, 0, 0}, []byte{1, 9, 9})}

	// Assert
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VSB returns correct value -- loop left greater", actual)
}

func Test_VersionSliceByte_ShorterLeft(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.VersionSliceByte([]byte{1, 0}, []byte{1, 0, 0})}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VSB returns correct value -- shorter left", actual)
}

func Test_VersionSliceByte_ShorterRight(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.VersionSliceByte([]byte{1, 0, 0}, []byte{1, 0})}

	// Assert
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VSB returns correct value -- shorter right", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// VersionSliceInteger — 7 branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_VersionSliceInteger_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.VersionSliceInteger(nil, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VSI returns nil -- both nil", actual)
}

func Test_VersionSliceInteger_LeftNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.VersionSliceInteger(nil, []int{1})}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VSI returns nil -- left nil", actual)
}

func Test_VersionSliceInteger_RightNil(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.VersionSliceInteger([]int{1}, nil)}

	// Assert
	expected := args.Map{"v": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VSI returns nil -- right nil", actual)
}

func Test_VersionSliceInteger_EqualSameLen(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2, 3})}

	// Assert
	expected := args.Map{"v": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VSI returns correct value -- equal same len", actual)
}

func Test_VersionSliceInteger_LoopLeftLess(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.VersionSliceInteger([]int{1, 0, 0}, []int{1, 0, 1})}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VSI returns correct value -- loop left less", actual)
}

func Test_VersionSliceInteger_LoopLeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.VersionSliceInteger([]int{2, 0, 0}, []int{1, 9, 9})}

	// Assert
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VSI returns correct value -- loop left greater", actual)
}

func Test_VersionSliceInteger_ShorterLeft(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.VersionSliceInteger([]int{1, 0}, []int{1, 0, 0})}

	// Assert
	expected := args.Map{"v": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VSI returns correct value -- shorter left", actual)
}

func Test_VersionSliceInteger_ShorterRight(t *testing.T) {
	// Act
	actual := args.Map{"v": corecmp.VersionSliceInteger([]int{1, 0, 0}, []int{1, 0})}

	// Assert
	expected := args.Map{"v": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VSI returns correct value -- shorter right", actual)
}
