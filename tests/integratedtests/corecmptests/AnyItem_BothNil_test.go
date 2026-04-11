package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// AnyItem
// ══════════════════════════════════════════════════════════════════════════════

func Test_AnyItem_BothNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.AnyItem(nil, nil)

	// Act
	actual := args.Map{"equal": result.IsEqual()}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- both nil", actual)
}

func Test_AnyItem_LeftNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.AnyItem(nil, "hello")

	// Act
	actual := args.Map{"notEqual": result.IsNotEqual()}

	// Assert
	expected := args.Map{"notEqual": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- left nil", actual)
}

func Test_AnyItem_RightNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.AnyItem("hello", nil)

	// Act
	actual := args.Map{"notEqual": result.IsNotEqual()}

	// Assert
	expected := args.Map{"notEqual": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- right nil", actual)
}

func Test_AnyItem_Equal_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.AnyItem(42, 42)

	// Act
	actual := args.Map{"equal": result.IsEqual()}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- equal", actual)
}

func Test_AnyItem_Inconclusive_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.AnyItem(42, 99)

	// Act
	actual := args.Map{"inconclusive": result.IsInconclusive()}

	// Assert
	expected := args.Map{"inconclusive": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- inconclusive", actual)
}

func Test_AnyItem_EqualStrings(t *testing.T) {
	// Arrange
	result := corecmp.AnyItem("abc", "abc")

	// Act
	actual := args.Map{"equal": result.IsEqual()}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- equal strings", actual)
}

func Test_AnyItem_DifferentStrings(t *testing.T) {
	// Arrange
	result := corecmp.AnyItem("abc", "xyz")

	// Act
	actual := args.Map{"inconclusive": result.IsInconclusive()}

	// Assert
	expected := args.Map{"inconclusive": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- different strings", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Byte
// ══════════════════════════════════════════════════════════════════════════════

func Test_Byte_Equal_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Byte(5, 5)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- equal", actual)
}

func Test_Byte_LeftLess_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Byte(1, 5)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- left less", actual)
}

func Test_Byte_LeftGreater_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Byte(10, 5)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- left greater", actual)
}

func Test_Byte_Zero(t *testing.T) {
	// Arrange
	result := corecmp.Byte(0, 0)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- zero", actual)
}

func Test_Byte_Max(t *testing.T) {
	// Arrange
	result := corecmp.Byte(255, 0)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- max", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BytePtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_BytePtr_BothNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.BytePtr(nil, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- both nil", actual)
}

func Test_BytePtr_LeftNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	b := byte(5)
	result := corecmp.BytePtr(nil, &b)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- left nil", actual)
}

func Test_BytePtr_RightNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	b := byte(5)
	result := corecmp.BytePtr(&b, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- right nil", actual)
}

func Test_BytePtr_Equal_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	a, b := byte(5), byte(5)
	result := corecmp.BytePtr(&a, &b)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "BytePtr returns correct value -- equal", actual)
}

func Test_BytePtr_LeftLess(t *testing.T) {
	// Arrange
	a, b := byte(1), byte(9)
	result := corecmp.BytePtr(&a, &b)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "BytePtr returns correct value -- left less", actual)
}

func Test_BytePtr_LeftGreater(t *testing.T) {
	// Arrange
	a, b := byte(9), byte(1)
	result := corecmp.BytePtr(&a, &b)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "BytePtr returns correct value -- left greater", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer
// ══════════════════════════════════════════════════════════════════════════════

func Test_Integer_Equal_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer(42, 42)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- equal", actual)
}

func Test_Integer_LeftLess_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer(1, 100)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- left less", actual)
}

func Test_Integer_LeftGreater_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer(100, 1)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- left greater", actual)
}

func Test_Integer_Negative(t *testing.T) {
	// Arrange
	result := corecmp.Integer(-5, 5)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- negative", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IntegerPtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_IntegerPtr_BothNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.IntegerPtr(nil, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- both nil", actual)
}

func Test_IntegerPtr_LeftNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	v := 5
	result := corecmp.IntegerPtr(nil, &v)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- left nil", actual)
}

func Test_IntegerPtr_RightNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	v := 5
	result := corecmp.IntegerPtr(&v, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- right nil", actual)
}

func Test_IntegerPtr_Equal(t *testing.T) {
	// Arrange
	a, b := 42, 42
	result := corecmp.IntegerPtr(&a, &b)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns correct value -- equal", actual)
}

func Test_IntegerPtr_LeftLess(t *testing.T) {
	// Arrange
	a, b := 1, 99
	result := corecmp.IntegerPtr(&a, &b)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns correct value -- left less", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer8 / Integer8Ptr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Integer8_Equal_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer8(5, 5)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- equal", actual)
}

func Test_Integer8_LeftLess_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer8(-10, 10)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- left less", actual)
}

func Test_Integer8_LeftGreater_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer8(10, -10)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- left greater", actual)
}

func Test_Integer8Ptr_BothNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer8Ptr(nil, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- both nil", actual)
}

func Test_Integer8Ptr_LeftNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	v := int8(5)
	result := corecmp.Integer8Ptr(nil, &v)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- left nil", actual)
}

func Test_Integer8Ptr_RightNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	v := int8(5)
	result := corecmp.Integer8Ptr(&v, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- right nil", actual)
}

func Test_Integer8Ptr_Equal_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	a, b := int8(3), int8(3)
	result := corecmp.Integer8Ptr(&a, &b)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns correct value -- equal", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer16 / Integer16Ptr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Integer16_Equal_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer16(100, 100)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- equal", actual)
}

func Test_Integer16_LeftLess_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer16(-100, 100)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- left less", actual)
}

func Test_Integer16_LeftGreater_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer16(100, -100)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- left greater", actual)
}

func Test_Integer16Ptr_BothNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer16Ptr(nil, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- both nil", actual)
}

func Test_Integer16Ptr_LeftNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	v := int16(5)
	result := corecmp.Integer16Ptr(nil, &v)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- left nil", actual)
}

func Test_Integer16Ptr_RightNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	v := int16(5)
	result := corecmp.Integer16Ptr(&v, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- right nil", actual)
}

func Test_Integer16Ptr_Equal_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	a, b := int16(7), int16(7)
	result := corecmp.Integer16Ptr(&a, &b)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns correct value -- equal", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer32 / Integer32Ptr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Integer32_Equal_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer32(1000, 1000)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- equal", actual)
}

func Test_Integer32_LeftLess_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer32(-1000, 1000)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- left less", actual)
}

func Test_Integer32_LeftGreater_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer32(1000, -1000)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- left greater", actual)
}

func Test_Integer32Ptr_BothNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer32Ptr(nil, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- both nil", actual)
}

func Test_Integer32Ptr_LeftNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	v := int32(5)
	result := corecmp.Integer32Ptr(nil, &v)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- left nil", actual)
}

func Test_Integer32Ptr_RightNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	v := int32(5)
	result := corecmp.Integer32Ptr(&v, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- right nil", actual)
}

func Test_Integer32Ptr_Equal_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	a, b := int32(7), int32(7)
	result := corecmp.Integer32Ptr(&a, &b)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns correct value -- equal", actual)
}

func Test_Integer32Ptr_LeftLess(t *testing.T) {
	// Arrange
	a, b := int32(1), int32(99)
	result := corecmp.Integer32Ptr(&a, &b)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns correct value -- left less", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer64 / Integer64Ptr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Integer64_Equal_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer64(100000, 100000)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- equal", actual)
}

func Test_Integer64_LeftLess_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer64(-100000, 100000)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- left less", actual)
}

func Test_Integer64_LeftGreater_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer64(100000, -100000)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- left greater", actual)
}

func Test_Integer64Ptr_BothNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.Integer64Ptr(nil, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- both nil", actual)
}

func Test_Integer64Ptr_LeftNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	v := int64(5)
	result := corecmp.Integer64Ptr(nil, &v)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- left nil", actual)
}

func Test_Integer64Ptr_RightNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	v := int64(5)
	result := corecmp.Integer64Ptr(&v, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- right nil", actual)
}

func Test_Integer64Ptr_Equal_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	a, b := int64(7), int64(7)
	result := corecmp.Integer64Ptr(&a, &b)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns correct value -- equal", actual)
}

func Test_Integer64Ptr_LeftGreater(t *testing.T) {
	// Arrange
	a, b := int64(99), int64(1)
	result := corecmp.Integer64Ptr(&a, &b)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns correct value -- left greater", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsIntegersEqual / IsIntegersEqualPtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsIntegersEqual_BothNil_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsIntegersEqual(nil, nil)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns nil -- both nil", actual)
}

func Test_IsIntegersEqual_LeftNil_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsIntegersEqual(nil, []int{1})}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns nil -- left nil", actual)
}

func Test_IsIntegersEqual_RightNil_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsIntegersEqual([]int{1}, nil)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns nil -- right nil", actual)
}

func Test_IsIntegersEqual_Same(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsIntegersEqual([]int{1, 2, 3}, []int{1, 2, 3})}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns correct value -- same", actual)
}

func Test_IsIntegersEqual_Different(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 3})}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns correct value -- different", actual)
}

func Test_IsIntegersEqual_DifferentLength(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsIntegersEqual([]int{1}, []int{1, 2})}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns correct value -- different length", actual)
}

func Test_IsIntegersEqualPtr_BothNil_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsIntegersEqualPtr(nil, nil)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- both nil", actual)
}

func Test_IsIntegersEqualPtr_LeftNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	right := []int{1}

	// Act
	actual := args.Map{"equal": corecmp.IsIntegersEqualPtr(nil, &right)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- left nil", actual)
}

func Test_IsIntegersEqualPtr_RightNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	left := []int{1}

	// Act
	actual := args.Map{"equal": corecmp.IsIntegersEqualPtr(&left, nil)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- right nil", actual)
}

func Test_IsIntegersEqualPtr_DifferentLength(t *testing.T) {
	// Arrange
	left := []int{1}
	right := []int{1, 2}

	// Act
	actual := args.Map{"equal": corecmp.IsIntegersEqualPtr(&left, &right)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns correct value -- different length", actual)
}

func Test_IsIntegersEqualPtr_Same_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	left := []int{1, 2, 3}
	right := []int{1, 2, 3}

	// Act
	actual := args.Map{"equal": corecmp.IsIntegersEqualPtr(&left, &right)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns correct value -- same", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsStringsEqual / IsStringsEqualPtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsStringsEqual_BothNil_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqual(nil, nil)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- both nil", actual)
}

func Test_IsStringsEqual_LeftNil_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqual(nil, []string{"a"})}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- left nil", actual)
}

func Test_IsStringsEqual_RightNil_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqual([]string{"a"}, nil)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- right nil", actual)
}

func Test_IsStringsEqual_Same_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- same", actual)
}

func Test_IsStringsEqual_Different_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqual([]string{"a"}, []string{"b"})}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- different", actual)
}

func Test_IsStringsEqual_DifferentLength_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqual([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- different length", actual)
}

func Test_IsStringsEqualPtr_BothNil_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqualPtr(nil, nil)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns nil -- both nil", actual)
}

func Test_IsStringsEqualPtr_LeftNil_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqualPtr(nil, []string{"a"})}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns nil -- left nil", actual)
}

func Test_IsStringsEqualPtr_RightNil_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqualPtr([]string{"a"}, nil)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns nil -- right nil", actual)
}

func Test_IsStringsEqualPtr_DifferentLength(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns correct value -- different length", actual)
}

func Test_IsStringsEqualPtr_Same_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqualPtr([]string{"x", "y"}, []string{"x", "y"})}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns correct value -- same", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsStringsEqualWithoutOrder
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsStringsEqualWithoutOrder_BothNil_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqualWithoutOrder(nil, nil)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns nil -- both nil", actual)
}

func Test_IsStringsEqualWithoutOrder_LeftNil_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"})}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns nil -- left nil", actual)
}

func Test_IsStringsEqualWithoutOrder_RightNil_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, nil)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns nil -- right nil", actual)
}

func Test_IsStringsEqualWithoutOrder_DifferentLength(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- different length", actual)
}

func Test_IsStringsEqualWithoutOrder_SameOrder_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqualWithoutOrder([]string{"a", "b"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- same order", actual)
}

func Test_IsStringsEqualWithoutOrder_DifferentOrder(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqualWithoutOrder([]string{"b", "a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- different order", actual)
}

func Test_IsStringsEqualWithoutOrder_Mismatch_FromAnyItemBothNil(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.IsStringsEqualWithoutOrder([]string{"a", "b"}, []string{"a", "c"})}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Time / TimePtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Time_Equal_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	now := time.Now()
	result := corecmp.Time(now, now)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- equal", actual)
}

func Test_Time_LeftLess_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	now := time.Now()
	later := now.Add(time.Hour)
	result := corecmp.Time(now, later)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- left less", actual)
}

func Test_Time_LeftGreater_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	now := time.Now()
	earlier := now.Add(-time.Hour)
	result := corecmp.Time(now, earlier)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- left greater", actual)
}

func Test_TimePtr_BothNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.TimePtr(nil, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- both nil", actual)
}

func Test_TimePtr_LeftNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	now := time.Now()
	result := corecmp.TimePtr(nil, &now)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- left nil", actual)
}

func Test_TimePtr_RightNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	now := time.Now()
	result := corecmp.TimePtr(&now, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- right nil", actual)
}

func Test_TimePtr_Equal_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	now := time.Now()
	same := now
	result := corecmp.TimePtr(&now, &same)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "TimePtr returns correct value -- equal", actual)
}

func Test_TimePtr_LeftLess(t *testing.T) {
	// Arrange
	now := time.Now()
	later := now.Add(time.Hour)
	result := corecmp.TimePtr(&now, &later)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "TimePtr returns correct value -- left less", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// VersionSliceByte
// ══════════════════════════════════════════════════════════════════════════════

func Test_VersionSliceByte_BothNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceByte(nil, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns nil -- both nil", actual)
}

func Test_VersionSliceByte_LeftNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceByte(nil, []byte{1, 0, 0})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns nil -- left nil", actual)
}

func Test_VersionSliceByte_RightNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceByte([]byte{1, 0, 0}, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns nil -- right nil", actual)
}

func Test_VersionSliceByte_Equal_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2, 3})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- equal", actual)
}

func Test_VersionSliceByte_LeftLess_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceByte([]byte{1, 0, 0}, []byte{1, 0, 1})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- left less", actual)
}

func Test_VersionSliceByte_LeftGreater_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceByte([]byte{2, 0, 0}, []byte{1, 9, 9})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- left greater", actual)
}

func Test_VersionSliceByte_ShorterLeft_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceByte([]byte{1, 0}, []byte{1, 0, 0})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- shorter left", actual)
}

func Test_VersionSliceByte_ShorterRight_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceByte([]byte{1, 0, 0}, []byte{1, 0})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- shorter right", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// VersionSliceInteger
// ══════════════════════════════════════════════════════════════════════════════

func Test_VersionSliceInteger_BothNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceInteger(nil, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- both nil", actual)
}

func Test_VersionSliceInteger_LeftNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceInteger(nil, []int{1, 0, 0})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- left nil", actual)
}

func Test_VersionSliceInteger_RightNil_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceInteger([]int{1, 0, 0}, nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- right nil", actual)
}

func Test_VersionSliceInteger_Equal_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2, 3})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- equal", actual)
}

func Test_VersionSliceInteger_LeftLess_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceInteger([]int{1, 0, 0}, []int{1, 0, 1})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- left less", actual)
}

func Test_VersionSliceInteger_LeftGreater_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceInteger([]int{2, 0, 0}, []int{1, 9, 9})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- left greater", actual)
}

func Test_VersionSliceInteger_ShorterLeft_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceInteger([]int{1, 0}, []int{1, 0, 0})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- shorter left", actual)
}

func Test_VersionSliceInteger_ShorterRight_FromAnyItemBothNil(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceInteger([]int{1, 0, 0}, []int{1, 0})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- shorter right", actual)
}

func Test_VersionSliceInteger_Empty(t *testing.T) {
	// Arrange
	result := corecmp.VersionSliceInteger([]int{}, []int{})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns empty -- empty", actual)
}
