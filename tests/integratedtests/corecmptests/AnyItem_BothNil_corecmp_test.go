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

func Test_AnyItem_BothNil_FromAnyItemBothNilcorecm(t *testing.T) {
	// Arrange / Act / Assert
	actual := args.Map{"result": corecmp.AnyItem(nil, nil)}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- both nil", actual)
}

func Test_AnyItem_LeftNil_FromAnyItemBothNilcorecm(t *testing.T) {
	// Arrange / Act / Assert
	actual := args.Map{"result": corecmp.AnyItem(nil, "a")}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- left nil", actual)
}

func Test_AnyItem_RightNil_FromAnyItemBothNilcorecm(t *testing.T) {
	// Arrange / Act / Assert
	actual := args.Map{"result": corecmp.AnyItem("a", nil)}
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- right nil", actual)
}

func Test_AnyItem_Equal_FromAnyItemBothNilcorecm(t *testing.T) {
	// Arrange / Act / Assert
	actual := args.Map{"result": corecmp.AnyItem("hello", "hello")}
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- equal", actual)
}

func Test_AnyItem_Inconclusive_FromAnyItemBothNilcorecm(t *testing.T) {
	// Arrange / Act / Assert
	actual := args.Map{"result": corecmp.AnyItem("a", "b")}
	expected := args.Map{"result": corecomparator.Inconclusive}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- inconclusive", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Byte + BytePtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Byte_AllBranches(t *testing.T) {
	// Arrange / Act / Assert
	actual := args.Map{
		"equal":   corecmp.Byte(5, 5),
		"less":    corecmp.Byte(1, 5),
		"greater": corecmp.Byte(5, 1),
	}
	expected := args.Map{
		"equal":   corecomparator.Equal,
		"less":    corecomparator.LeftLess,
		"greater": corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- all branches", actual)
}

func Test_BytePtr_AllBranches(t *testing.T) {
	// Arrange
	a, b := byte(5), byte(5)
	c := byte(1)
	// Act / Assert
	actual := args.Map{
		"bothNil":  corecmp.BytePtr(nil, nil),
		"leftNil":  corecmp.BytePtr(nil, &a),
		"rightNil": corecmp.BytePtr(&a, nil),
		"equal":    corecmp.BytePtr(&a, &b),
		"less":     corecmp.BytePtr(&c, &a),
	}
	expected := args.Map{
		"bothNil":  corecomparator.Equal,
		"leftNil":  corecomparator.NotEqual,
		"rightNil": corecomparator.NotEqual,
		"equal":    corecomparator.Equal,
		"less":     corecomparator.LeftLess,
	}
	expected.ShouldBeEqual(t, 0, "BytePtr returns correct value -- all branches", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer8 + Integer8Ptr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Integer8_AllBranches(t *testing.T) {
	// Arrange / Act / Assert
	actual := args.Map{
		"equal":   corecmp.Integer8(5, 5),
		"less":    corecmp.Integer8(1, 5),
		"greater": corecmp.Integer8(5, 1),
	}
	expected := args.Map{
		"equal":   corecomparator.Equal,
		"less":    corecomparator.LeftLess,
		"greater": corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- all branches", actual)
}

func Test_Integer8Ptr_AllBranches(t *testing.T) {
	// Arrange
	a, b := int8(5), int8(5)
	c := int8(1)
	// Act / Assert
	actual := args.Map{
		"bothNil":  corecmp.Integer8Ptr(nil, nil),
		"leftNil":  corecmp.Integer8Ptr(nil, &a),
		"rightNil": corecmp.Integer8Ptr(&a, nil),
		"equal":    corecmp.Integer8Ptr(&a, &b),
		"less":     corecmp.Integer8Ptr(&c, &a),
	}
	expected := args.Map{
		"bothNil":  corecomparator.Equal,
		"leftNil":  corecomparator.NotEqual,
		"rightNil": corecomparator.NotEqual,
		"equal":    corecomparator.Equal,
		"less":     corecomparator.LeftLess,
	}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns correct value -- all branches", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer16 + Integer16Ptr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Integer16_AllBranches(t *testing.T) {
	// Arrange / Act / Assert
	actual := args.Map{
		"equal":   corecmp.Integer16(5, 5),
		"less":    corecmp.Integer16(1, 5),
		"greater": corecmp.Integer16(5, 1),
	}
	expected := args.Map{
		"equal":   corecomparator.Equal,
		"less":    corecomparator.LeftLess,
		"greater": corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- all branches", actual)
}

func Test_Integer16Ptr_AllBranches(t *testing.T) {
	// Arrange
	a, b := int16(5), int16(5)
	c := int16(1)
	// Act / Assert
	actual := args.Map{
		"bothNil":  corecmp.Integer16Ptr(nil, nil),
		"leftNil":  corecmp.Integer16Ptr(nil, &a),
		"rightNil": corecmp.Integer16Ptr(&a, nil),
		"equal":    corecmp.Integer16Ptr(&a, &b),
		"less":     corecmp.Integer16Ptr(&c, &a),
	}
	expected := args.Map{
		"bothNil":  corecomparator.Equal,
		"leftNil":  corecomparator.NotEqual,
		"rightNil": corecomparator.NotEqual,
		"equal":    corecomparator.Equal,
		"less":     corecomparator.LeftLess,
	}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns correct value -- all branches", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer32 + Integer32Ptr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Integer32_AllBranches(t *testing.T) {
	// Arrange / Act / Assert
	actual := args.Map{
		"equal":   corecmp.Integer32(5, 5),
		"less":    corecmp.Integer32(1, 5),
		"greater": corecmp.Integer32(5, 1),
	}
	expected := args.Map{
		"equal":   corecomparator.Equal,
		"less":    corecomparator.LeftLess,
		"greater": corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- all branches", actual)
}

func Test_Integer32Ptr_AllBranches(t *testing.T) {
	// Arrange
	a, b := int32(5), int32(5)
	c := int32(1)
	// Act / Assert
	actual := args.Map{
		"bothNil":  corecmp.Integer32Ptr(nil, nil),
		"leftNil":  corecmp.Integer32Ptr(nil, &a),
		"rightNil": corecmp.Integer32Ptr(&a, nil),
		"equal":    corecmp.Integer32Ptr(&a, &b),
		"less":     corecmp.Integer32Ptr(&c, &a),
	}
	expected := args.Map{
		"bothNil":  corecomparator.Equal,
		"leftNil":  corecomparator.NotEqual,
		"rightNil": corecomparator.NotEqual,
		"equal":    corecomparator.Equal,
		"less":     corecomparator.LeftLess,
	}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns correct value -- all branches", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Integer64 + Integer64Ptr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Integer64_AllBranches(t *testing.T) {
	// Arrange / Act / Assert
	actual := args.Map{
		"equal":   corecmp.Integer64(5, 5),
		"less":    corecmp.Integer64(1, 5),
		"greater": corecmp.Integer64(5, 1),
	}
	expected := args.Map{
		"equal":   corecomparator.Equal,
		"less":    corecomparator.LeftLess,
		"greater": corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- all branches", actual)
}

func Test_Integer64Ptr_AllBranches(t *testing.T) {
	// Arrange
	a, b := int64(5), int64(5)
	c := int64(1)
	// Act / Assert
	actual := args.Map{
		"bothNil":  corecmp.Integer64Ptr(nil, nil),
		"leftNil":  corecmp.Integer64Ptr(nil, &a),
		"rightNil": corecmp.Integer64Ptr(&a, nil),
		"equal":    corecmp.Integer64Ptr(&a, &b),
		"less":     corecmp.Integer64Ptr(&c, &a),
	}
	expected := args.Map{
		"bothNil":  corecomparator.Equal,
		"leftNil":  corecomparator.NotEqual,
		"rightNil": corecomparator.NotEqual,
		"equal":    corecomparator.Equal,
		"less":     corecomparator.LeftLess,
	}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns correct value -- all branches", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IntegerPtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_IntegerPtr_AllBranches(t *testing.T) {
	// Arrange
	a, b := 5, 5
	c := 1
	// Act / Assert
	actual := args.Map{
		"bothNil":  corecmp.IntegerPtr(nil, nil),
		"leftNil":  corecmp.IntegerPtr(nil, &a),
		"rightNil": corecmp.IntegerPtr(&a, nil),
		"equal":    corecmp.IntegerPtr(&a, &b),
		"less":     corecmp.IntegerPtr(&c, &a),
	}
	expected := args.Map{
		"bothNil":  corecomparator.Equal,
		"leftNil":  corecomparator.NotEqual,
		"rightNil": corecomparator.NotEqual,
		"equal":    corecomparator.Equal,
		"less":     corecomparator.LeftLess,
	}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns correct value -- all branches", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsIntegersEqual + IsIntegersEqualPtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsIntegersEqual_AllBranches(t *testing.T) {
	// Arrange / Act / Assert
	actual := args.Map{
		"bothNil":  corecmp.IsIntegersEqual(nil, nil),
		"leftNil":  corecmp.IsIntegersEqual(nil, []int{1}),
		"rightNil": corecmp.IsIntegersEqual([]int{1}, nil),
		"equal":    corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 2}),
	}
	expected := args.Map{
		"bothNil": true, "leftNil": false, "rightNil": false, "equal": true,
	}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns correct value -- all branches", actual)
}

func Test_IsIntegersEqualPtr_AllBranches(t *testing.T) {
	// Arrange
	a := []int{1, 2}
	b := []int{1, 2}
	c := []int{1}
	// Act / Assert
	actual := args.Map{
		"bothNil":  corecmp.IsIntegersEqualPtr(nil, nil),
		"leftNil":  corecmp.IsIntegersEqualPtr(nil, &a),
		"rightNil": corecmp.IsIntegersEqualPtr(&a, nil),
		"equal":    corecmp.IsIntegersEqualPtr(&a, &b),
		"diffLen":  corecmp.IsIntegersEqualPtr(&a, &c),
	}
	expected := args.Map{
		"bothNil": true, "leftNil": false, "rightNil": false,
		"equal": true, "diffLen": false,
	}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns correct value -- all branches", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsStringsEqualPtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsStringsEqualPtr_AllBranches(t *testing.T) {
	// Arrange / Act / Assert
	actual := args.Map{
		"bothNil":  corecmp.IsStringsEqualPtr(nil, nil),
		"leftNil":  corecmp.IsStringsEqualPtr(nil, []string{"a"}),
		"rightNil": corecmp.IsStringsEqualPtr([]string{"a"}, nil),
		"diffLen":  corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a", "b"}),
		"equal":    corecmp.IsStringsEqualPtr([]string{"a", "b"}, []string{"a", "b"}),
	}
	expected := args.Map{
		"bothNil": true, "leftNil": false, "rightNil": false,
		"diffLen": false, "equal": true,
	}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns correct value -- all branches", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsStringsEqualWithoutOrder
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsStringsEqualWithoutOrder_AllBranches(t *testing.T) {
	// Arrange / Act / Assert
	actual := args.Map{
		"bothNil":  corecmp.IsStringsEqualWithoutOrder(nil, nil),
		"leftNil":  corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"}),
		"rightNil": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, nil),
		"diffLen":  corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"}),
		"sorted":   corecmp.IsStringsEqualWithoutOrder([]string{"b", "a"}, []string{"a", "b"}),
		"notEqual": corecmp.IsStringsEqualWithoutOrder([]string{"x"}, []string{"y"}),
	}
	expected := args.Map{
		"bothNil": true, "leftNil": false, "rightNil": false,
		"diffLen": false, "sorted": true, "notEqual": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- all branches", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Time + TimePtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Time_AllBranches(t *testing.T) {
	// Arrange
	now := time.Now()
	earlier := now.Add(-time.Hour)
	later := now.Add(time.Hour)
	// Act / Assert
	actual := args.Map{
		"equal":   corecmp.Time(now, now),
		"less":    corecmp.Time(earlier, now),
		"greater": corecmp.Time(later, now),
	}
	expected := args.Map{
		"equal":   corecomparator.Equal,
		"less":    corecomparator.LeftLess,
		"greater": corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- all branches", actual)
}

func Test_TimePtr_AllBranches(t *testing.T) {
	// Arrange
	now := time.Now()
	same := now
	earlier := now.Add(-time.Hour)
	// Act / Assert
	actual := args.Map{
		"bothNil":  corecmp.TimePtr(nil, nil),
		"leftNil":  corecmp.TimePtr(nil, &now),
		"rightNil": corecmp.TimePtr(&now, nil),
		"equal":    corecmp.TimePtr(&now, &same),
		"less":     corecmp.TimePtr(&earlier, &now),
	}
	expected := args.Map{
		"bothNil":  corecomparator.Equal,
		"leftNil":  corecomparator.NotEqual,
		"rightNil": corecomparator.NotEqual,
		"equal":    corecomparator.Equal,
		"less":     corecomparator.LeftLess,
	}
	expected.ShouldBeEqual(t, 0, "TimePtr returns correct value -- all branches", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// VersionSliceByte
// ══════════════════════════════════════════════════════════════════════════════

func Test_VersionSliceByte_AllBranches(t *testing.T) {
	// Arrange / Act / Assert
	actual := args.Map{
		"bothNil":      corecmp.VersionSliceByte(nil, nil),
		"leftNil":      corecmp.VersionSliceByte(nil, []byte{1}),
		"rightNil":     corecmp.VersionSliceByte([]byte{1}, nil),
		"equal":        corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2, 3}),
		"lessElement":  corecmp.VersionSliceByte([]byte{1, 1, 3}, []byte{1, 2, 3}),
		"greatElement": corecmp.VersionSliceByte([]byte{1, 3, 3}, []byte{1, 2, 3}),
		"lessLen":      corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 2, 3}),
		"greatLen":     corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2}),
	}
	expected := args.Map{
		"bothNil":      corecomparator.Equal,
		"leftNil":      corecomparator.NotEqual,
		"rightNil":     corecomparator.NotEqual,
		"equal":        corecomparator.Equal,
		"lessElement":  corecomparator.LeftLess,
		"greatElement": corecomparator.LeftGreater,
		"lessLen":      corecomparator.LeftLess,
		"greatLen":     corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- all branches", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// VersionSliceInteger
// ══════════════════════════════════════════════════════════════════════════════

func Test_VersionSliceInteger_AllBranches(t *testing.T) {
	// Arrange / Act / Assert
	actual := args.Map{
		"bothNil":      corecmp.VersionSliceInteger(nil, nil),
		"leftNil":      corecmp.VersionSliceInteger(nil, []int{1}),
		"rightNil":     corecmp.VersionSliceInteger([]int{1}, nil),
		"equal":        corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2, 3}),
		"lessElement":  corecmp.VersionSliceInteger([]int{1, 1, 3}, []int{1, 2, 3}),
		"greatElement": corecmp.VersionSliceInteger([]int{1, 3, 3}, []int{1, 2, 3}),
		"lessLen":      corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 2, 3}),
		"greatLen":     corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2}),
	}
	expected := args.Map{
		"bothNil":      corecomparator.Equal,
		"leftNil":      corecomparator.NotEqual,
		"rightNil":     corecomparator.NotEqual,
		"equal":        corecomparator.Equal,
		"lessElement":  corecomparator.LeftLess,
		"greatElement": corecomparator.LeftGreater,
		"lessLen":      corecomparator.LeftLess,
		"greatLen":     corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- all branches", actual)
}
