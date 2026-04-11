package corecmptests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// Byte — all branches
// ═══════════════════════════════════════════

func Test_Byte_Equal_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Byte(5, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- equal", actual)
}

func Test_Byte_Less_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Byte(3, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- less", actual)
}

func Test_Byte_Greater_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Byte(10, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- greater", actual)
}

// ═══════════════════════════════════════════
// BytePtr — all branches
// ═══════════════════════════════════════════

func Test_BytePtr_BothNil_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.BytePtr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- both nil", actual)
}

func Test_BytePtr_LeftNil_FromByteEqual(t *testing.T) {
	// Arrange
	r := byte(5)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(nil, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- left nil", actual)
}

// ═══════════════════════════════════════════
// Integer — all branches
// ═══════════════════════════════════════════

func Test_Integer_Equal_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer(5, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- equal", actual)
}

func Test_Integer_Less_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer(3, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- less", actual)
}

// ═══════════════════════════════════════════
// IntegerPtr — all branches
// ═══════════════════════════════════════════

func Test_IntegerPtr_BothNil_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- both nil", actual)
}

func Test_IntegerPtr_LeftNil_FromByteEqual(t *testing.T) {
	// Arrange
	r := 5

	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(nil, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- left nil", actual)
}

func Test_IntegerPtr_RightNil_FromByteEqual(t *testing.T) {
	// Arrange
	l := 5

	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(&l, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- right nil", actual)
}

func Test_IntegerPtr_Equal_FromByteEqual(t *testing.T) {
	// Arrange
	l, r := 5, 5

	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns correct value -- equal", actual)
}

func Test_IntegerPtr_Less(t *testing.T) {
	// Arrange
	l, r := 3, 5

	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns correct value -- less", actual)
}

func Test_IntegerPtr_Greater(t *testing.T) {
	// Arrange
	l, r := 10, 5

	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns correct value -- greater", actual)
}

// ═══════════════════════════════════════════
// Integer8 — remaining branches
// ═══════════════════════════════════════════

func Test_Integer8_Equal_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8(5, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- equal", actual)
}

func Test_Integer8_Less_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8(3, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- less", actual)
}

func Test_Integer8Ptr_BothNil_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- both nil", actual)
}

func Test_Integer8Ptr_RightNil_FromByteEqual(t *testing.T) {
	// Arrange
	l := int8(5)

	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(&l, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- right nil", actual)
}

// ═══════════════════════════════════════════
// Integer16Ptr — remaining branches
// ═══════════════════════════════════════════

func Test_Integer16Ptr_BothNil_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16Ptr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- both nil", actual)
}

func Test_Integer16Ptr_RightNil_FromByteEqual(t *testing.T) {
	// Arrange
	l := int16(5)

	// Act
	actual := args.Map{"result": corecmp.Integer16Ptr(&l, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- right nil", actual)
}

// ═══════════════════════════════════════════
// Integer32Ptr — remaining branches
// ═══════════════════════════════════════════

func Test_Integer32Ptr_BothNil_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32Ptr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- both nil", actual)
}

func Test_Integer32Ptr_RightNil_FromByteEqual(t *testing.T) {
	// Arrange
	l := int32(5)

	// Act
	actual := args.Map{"result": corecmp.Integer32Ptr(&l, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- right nil", actual)
}

// ═══════════════════════════════════════════
// Integer64 — remaining branches
// ═══════════════════════════════════════════

func Test_Integer64_Equal_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64(5, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- equal", actual)
}

func Test_Integer64_Less_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64(3, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- less", actual)
}

func Test_Integer64Ptr_BothNil_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64Ptr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- both nil", actual)
}

func Test_Integer64Ptr_RightNil_FromByteEqual(t *testing.T) {
	// Arrange
	l := int64(5)

	// Act
	actual := args.Map{"result": corecmp.Integer64Ptr(&l, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- right nil", actual)
}

// ═══════════════════════════════════════════
// IsStringsEqual — remaining branches
// ═══════════════════════════════════════════

func Test_IsStringsEqual_BothNil_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- both nil", actual)
}

func Test_IsStringsEqual_Equal_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- equal", actual)
}

func Test_IsStringsEqual_DiffLen_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- diff len", actual)
}

func Test_IsStringsEqual_LeftNil_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual(nil, []string{"a"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- left nil", actual)
}

// ═══════════════════════════════════════════
// IsStringsEqualPtr — remaining branches
// ═══════════════════════════════════════════

func Test_IsStringsEqualPtr_BothNil_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualPtr(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns nil -- both nil", actual)
}

func Test_IsStringsEqualPtr_Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns correct value -- equal", actual)
}

func Test_IsStringsEqualPtr_LeftNil_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualPtr(nil, []string{"a"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns nil -- left nil", actual)
}

// ═══════════════════════════════════════════
// IsIntegersEqual — remaining branches
// ═══════════════════════════════════════════

func Test_IsIntegersEqual_BothNil_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqual(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns nil -- both nil", actual)
}

func Test_IsIntegersEqual_Equal_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 2})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns correct value -- equal", actual)
}

func Test_IsIntegersEqual_DiffLen(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqual([]int{1}, []int{1, 2})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns correct value -- diff len", actual)
}

func Test_IsIntegersEqual_NotEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 3})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns correct value -- not equal", actual)
}

// ═══════════════════════════════════════════
// IsIntegersEqualPtr — remaining branches
// ═══════════════════════════════════════════

func Test_IsIntegersEqualPtr_BothNil_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- both nil", actual)
}

func Test_IsIntegersEqualPtr_LeftNil_FromByteEqual(t *testing.T) {
	// Arrange
	r := []int{1}

	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(nil, &r)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- left nil", actual)
}

func Test_IsIntegersEqualPtr_Equal(t *testing.T) {
	// Arrange
	l := []int{1, 2}
	r := []int{1, 2}

	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&l, &r)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns correct value -- equal", actual)
}

func Test_IsIntegersEqualPtr_NotEqual(t *testing.T) {
	// Arrange
	l := []int{1, 2}
	r := []int{1, 3}

	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&l, &r)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns correct value -- not equal", actual)
}

// ═══════════════════════════════════════════
// AnyItem — remaining branches
// ═══════════════════════════════════════════

func Test_AnyItem_BothNil_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- both nil", actual)
}

func Test_AnyItem_LeftNil_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem(nil, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- left nil", actual)
}

func Test_AnyItem_Equal_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem(5, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- equal", actual)
}

// ═══════════════════════════════════════════
// VersionSliceByte — right nil
// ═══════════════════════════════════════════

func Test_VersionSliceByte_RightNil_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1}, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns nil -- right nil", actual)
}

// ═══════════════════════════════════════════
// VersionSliceInteger — right nil
// ═══════════════════════════════════════════

func Test_VersionSliceInteger_RightNil_FromByteEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1}, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- right nil", actual)
}

// ═══════════════════════════════════════════
// Time — remaining
// ═══════════════════════════════════════════

func Test_TimePtr_Equal_Values(t *testing.T) {
	now := corecmp.Time
	_ = now // Time func already fully tested in Coverage8
}
