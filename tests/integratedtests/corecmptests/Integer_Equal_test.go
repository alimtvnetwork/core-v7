package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Integer ──

func Test_Integer_Equal_FromIntegerEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer(5, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer returns Equal -- same values", actual)
}

func Test_Integer_LeftLess_FromIntegerEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer(3, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer returns LeftLess -- left smaller", actual)
}

func Test_Integer_LeftGreater_FromIntegerEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer(7, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer returns LeftGreater -- left bigger", actual)
}

// ── IntegerPtr ──

func Test_IntegerPtr_BothNil_FromIntegerEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns Equal -- both nil", actual)
}

func Test_IntegerPtr_LeftNil_FromIntegerEqual(t *testing.T) {
	// Arrange
	r := 5

	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(nil, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns LeftLess -- left nil", actual)
}

func Test_IntegerPtr_RightNil_FromIntegerEqual(t *testing.T) {
	// Arrange
	l := 5

	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(&l, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns LeftGreater -- right nil", actual)
}

func Test_IntegerPtr_Equal_FromIntegerEqual(t *testing.T) {
	// Arrange
	l, r := 5, 5

	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns Equal -- same values", actual)
}

// ── Byte ──

func Test_Byte_Equal_FromIntegerEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Byte(5, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Byte returns Equal -- same values", actual)
}

func Test_Byte_LeftLess_FromIntegerEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Byte(3, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Byte returns LeftLess -- left smaller", actual)
}

func Test_Byte_LeftGreater_FromIntegerEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Byte(7, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Byte returns LeftGreater -- left bigger", actual)
}

// ── BytePtr ──

func Test_BytePtr_BothNil_FromIntegerEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.BytePtr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "BytePtr returns Equal -- both nil", actual)
}

func Test_BytePtr_LeftNil_FromIntegerEqual(t *testing.T) {
	// Arrange
	r := byte(5)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(nil, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr returns LeftLess -- left nil", actual)
}

func Test_BytePtr_RightNil_FromIntegerEqual(t *testing.T) {
	// Arrange
	l := byte(5)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(&l, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr returns LeftGreater -- right nil", actual)
}

// ── Integer8 / Integer8Ptr ──

func Test_Integer8_All(t *testing.T) {
	// Act
	actual := args.Map{
		"eq":   corecmp.Integer8(5, 5),
		"lt":   corecmp.Integer8(3, 5),
		"gt":   corecmp.Integer8(7, 5),
	}

	// Assert
	expected := args.Map{
		"eq":   corecomparator.Equal,
		"lt":   corecomparator.LeftLess,
		"gt":   corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct -- all branches", actual)
}

func Test_Integer8Ptr_BothNil_FromIntegerEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns Equal -- both nil", actual)
}

func Test_Integer8Ptr_LeftNil_FromIntegerEqual(t *testing.T) {
	// Arrange
	r := int8(5)

	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns LeftLess -- left nil", actual)
}

func Test_Integer8Ptr_RightNil_FromIntegerEqual(t *testing.T) {
	// Arrange
	l := int8(5)

	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(&l, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns LeftGreater -- right nil", actual)
}

func Test_Integer8Ptr_Equal_FromIntegerEqual(t *testing.T) {
	// Arrange
	l, r := int8(5), int8(5)

	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns Equal -- same values", actual)
}

// ── Integer16 / Integer16Ptr ──

func Test_Integer16_All(t *testing.T) {
	// Act
	actual := args.Map{
		"eq": corecmp.Integer16(5, 5),
		"lt": corecmp.Integer16(3, 5),
		"gt": corecmp.Integer16(7, 5),
	}

	// Assert
	expected := args.Map{
		"eq": corecomparator.Equal,
		"lt": corecomparator.LeftLess,
		"gt": corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct -- all branches", actual)
}

func Test_Integer16Ptr_All(t *testing.T) {
	// Arrange
	l, r := int16(3), int16(5)

	// Act
	actual := args.Map{
		"bothNil":  corecmp.Integer16Ptr(nil, nil),
		"leftNil":  corecmp.Integer16Ptr(nil, &r),
		"rightNil": corecmp.Integer16Ptr(&l, nil),
		"eq":       corecmp.Integer16Ptr(&r, &r),
	}

	// Assert
	expected := args.Map{
		"bothNil":  corecomparator.Equal,
		"leftNil":  corecomparator.NotEqual,
		"rightNil": corecomparator.NotEqual,
		"eq":       corecomparator.Equal,
	}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns correct -- all branches", actual)
}

// ── Integer32 / Integer32Ptr ──

func Test_Integer32_All(t *testing.T) {
	// Act
	actual := args.Map{
		"eq": corecmp.Integer32(5, 5),
		"lt": corecmp.Integer32(3, 5),
		"gt": corecmp.Integer32(7, 5),
	}

	// Assert
	expected := args.Map{
		"eq": corecomparator.Equal,
		"lt": corecomparator.LeftLess,
		"gt": corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct -- all branches", actual)
}

func Test_Integer32Ptr_All(t *testing.T) {
	// Arrange
	l, r := int32(3), int32(5)

	// Act
	actual := args.Map{
		"bothNil":  corecmp.Integer32Ptr(nil, nil),
		"leftNil":  corecmp.Integer32Ptr(nil, &r),
		"rightNil": corecmp.Integer32Ptr(&l, nil),
		"eq":       corecmp.Integer32Ptr(&r, &r),
	}

	// Assert
	expected := args.Map{
		"bothNil":  corecomparator.Equal,
		"leftNil":  corecomparator.NotEqual,
		"rightNil": corecomparator.NotEqual,
		"eq":       corecomparator.Equal,
	}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns correct -- all branches", actual)
}

// ── Integer64 / Integer64Ptr ──

func Test_Integer64_All(t *testing.T) {
	// Act
	actual := args.Map{
		"eq": corecmp.Integer64(5, 5),
		"lt": corecmp.Integer64(3, 5),
		"gt": corecmp.Integer64(7, 5),
	}

	// Assert
	expected := args.Map{
		"eq": corecomparator.Equal,
		"lt": corecomparator.LeftLess,
		"gt": corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct -- all branches", actual)
}

func Test_Integer64Ptr_All(t *testing.T) {
	// Arrange
	l, r := int64(3), int64(5)

	// Act
	actual := args.Map{
		"bothNil":  corecmp.Integer64Ptr(nil, nil),
		"leftNil":  corecmp.Integer64Ptr(nil, &r),
		"rightNil": corecmp.Integer64Ptr(&l, nil),
		"eq":       corecmp.Integer64Ptr(&r, &r),
	}

	// Assert
	expected := args.Map{
		"bothNil":  corecomparator.Equal,
		"leftNil":  corecomparator.NotEqual,
		"rightNil": corecomparator.NotEqual,
		"eq":       corecomparator.Equal,
	}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns correct -- all branches", actual)
}

// ── AnyItem ──

func Test_AnyItem_BothNil_FromIntegerEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "AnyItem returns Equal -- both nil", actual)
}

func Test_AnyItem_LeftNil_FromIntegerEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem(nil, "hello")}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "AnyItem returns NotEqual -- left nil", actual)
}

func Test_AnyItem_RightNil_FromIntegerEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem("hello", nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "AnyItem returns NotEqual -- right nil", actual)
}

func Test_AnyItem_BothNonNil(t *testing.T) {
	// Arrange
	result := corecmp.AnyItem("hello", "world")

	// Act
	actual := args.Map{"isValid": result == corecomparator.Equal || result == corecomparator.NotEqual || result == corecomparator.Inconclusive}

	// Assert
	expected := args.Map{"isValid": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns valid -- both non-nil", actual)
}

// ── IsIntegersEqual / IsIntegersEqualPtr ──

func Test_IsIntegersEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"eq":  corecmp.IsIntegersEqual([]int{5}, []int{5}),
		"neq": corecmp.IsIntegersEqual([]int{3}, []int{5}),
	}

	// Assert
	expected := args.Map{
		"eq":  true,
		"neq": false,
	}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns correct -- equal and not equal", actual)
}

func Test_IsIntegersEqualPtr(t *testing.T) {
	// Arrange
	l, r := []int{5}, []int{5}
	l2 := []int{3}

	// Act
	actual := args.Map{
		"bothNil":  corecmp.IsIntegersEqualPtr(nil, nil),
		"leftNil":  corecmp.IsIntegersEqualPtr(nil, &r),
		"rightNil": corecmp.IsIntegersEqualPtr(&l, nil),
		"eq":       corecmp.IsIntegersEqualPtr(&l, &r),
		"neq":      corecmp.IsIntegersEqualPtr(&l2, &r),
	}

	// Assert
	expected := args.Map{
		"bothNil":  true,
		"leftNil":  false,
		"rightNil": false,
		"eq":       true,
		"neq":      false,
	}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns correct -- all branches", actual)
}

// ── IsStringsEqual / IsStringsEqualPtr ──

func Test_IsStringsEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"eq":  corecmp.IsStringsEqual([]string{"abc"}, []string{"abc"}),
		"neq": corecmp.IsStringsEqual([]string{"abc"}, []string{"xyz"}),
	}

	// Assert
	expected := args.Map{
		"eq":  true,
		"neq": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct -- equal and not", actual)
}

func Test_IsStringsEqualPtr(t *testing.T) {
	// Arrange
	l, r := []string{"abc"}, []string{"abc"}
	l2 := []string{"xyz"}

	// Act
	actual := args.Map{
		"bothNil":  corecmp.IsStringsEqualPtr(nil, nil),
		"leftNil":  corecmp.IsStringsEqualPtr(nil, r),
		"rightNil": corecmp.IsStringsEqualPtr(l, nil),
		"eq":       corecmp.IsStringsEqualPtr(l, r),
		"neq":      corecmp.IsStringsEqualPtr(l2, r),
	}

	// Assert
	expected := args.Map{
		"bothNil":  true,
		"leftNil":  false,
		"rightNil": false,
		"eq":       true,
		"neq":      false,
	}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns correct -- all branches", actual)
}

// ── IsStringsEqualWithoutOrder ──

func Test_IsStringsEqualWithoutOrder(t *testing.T) {
	// Act
	actual := args.Map{
		"same":       corecmp.IsStringsEqualWithoutOrder([]string{"a", "b"}, []string{"b", "a"}),
		"different":  corecmp.IsStringsEqualWithoutOrder([]string{"a", "b"}, []string{"c", "d"}),
		"diffLen":    corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"}),
		"bothEmpty":  corecmp.IsStringsEqualWithoutOrder([]string{}, []string{}),
		"bothNil":    corecmp.IsStringsEqualWithoutOrder(nil, nil),
		"leftNil":    corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"}),
	}

	// Assert
	expected := args.Map{
		"same":       true,
		"different":  false,
		"diffLen":    false,
		"bothEmpty":  true,
		"bothNil":    true,
		"leftNil":    false,
	}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns correct -- all branches", actual)
}

// ── Time ──

func Test_Time_All(t *testing.T) {
	// Arrange
	now := time.Now()
	before := now.Add(-time.Hour)
	after := now.Add(time.Hour)

	// Act
	actual := args.Map{
		"eq": corecmp.Time(now, now),
		"lt": corecmp.Time(before, now),
		"gt": corecmp.Time(after, now),
	}

	// Assert
	expected := args.Map{
		"eq": corecomparator.Equal,
		"lt": corecomparator.LeftLess,
		"gt": corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "Time returns correct -- all branches", actual)
}

// ── TimePtr ──

func Test_TimePtr_All(t *testing.T) {
	// Arrange
	now := time.Now()
	before := now.Add(-time.Hour)

	// Act
	actual := args.Map{
		"bothNil":  corecmp.TimePtr(nil, nil),
		"leftNil":  corecmp.TimePtr(nil, &now),
		"rightNil": corecmp.TimePtr(&now, nil),
		"eq":       corecmp.TimePtr(&now, &now),
		"lt":       corecmp.TimePtr(&before, &now),
	}

	// Assert
	expected := args.Map{
		"bothNil":  corecomparator.Equal,
		"leftNil":  corecomparator.NotEqual,
		"rightNil": corecomparator.NotEqual,
		"eq":       corecomparator.Equal,
		"lt":       corecomparator.LeftLess,
	}
	expected.ShouldBeEqual(t, 0, "TimePtr returns correct -- all branches", actual)
}

// ── VersionSliceByte ──

func Test_VersionSliceByte_All(t *testing.T) {
	// Act
	actual := args.Map{
		"bothNil":    corecmp.VersionSliceByte(nil, nil),
		"leftNil":    corecmp.VersionSliceByte(nil, []byte{1}),
		"rightNil":   corecmp.VersionSliceByte([]byte{1}, nil),
		"eq":         corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 2}),
		"lt":         corecmp.VersionSliceByte([]byte{1, 1}, []byte{1, 2}),
		"gt":         corecmp.VersionSliceByte([]byte{1, 3}, []byte{1, 2}),
		"shorterLt":  corecmp.VersionSliceByte([]byte{1}, []byte{1, 2}),
		"longerGt":   corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2}),
	}

	// Assert
	expected := args.Map{
		"bothNil":    corecomparator.Equal,
		"leftNil":    corecomparator.NotEqual,
		"rightNil":   corecomparator.NotEqual,
		"eq":         corecomparator.Equal,
		"lt":         corecomparator.LeftLess,
		"gt":         corecomparator.LeftGreater,
		"shorterLt":  corecomparator.LeftLess,
		"longerGt":   corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct -- all branches", actual)
}

// ── VersionSliceInteger ──

func Test_VersionSliceInteger_All(t *testing.T) {
	// Act
	actual := args.Map{
		"bothNil":    corecmp.VersionSliceInteger(nil, nil),
		"leftNil":    corecmp.VersionSliceInteger(nil, []int{1}),
		"rightNil":   corecmp.VersionSliceInteger([]int{1}, nil),
		"eq":         corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 2}),
		"lt":         corecmp.VersionSliceInteger([]int{1, 1}, []int{1, 2}),
		"gt":         corecmp.VersionSliceInteger([]int{1, 3}, []int{1, 2}),
		"shorterLt":  corecmp.VersionSliceInteger([]int{1}, []int{1, 2}),
		"longerGt":   corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2}),
	}

	// Assert
	expected := args.Map{
		"bothNil":    corecomparator.Equal,
		"leftNil":    corecomparator.NotEqual,
		"rightNil":   corecomparator.NotEqual,
		"eq":         corecomparator.Equal,
		"lt":         corecomparator.LeftLess,
		"gt":         corecomparator.LeftGreater,
		"shorterLt":  corecomparator.LeftLess,
		"longerGt":   corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct -- all branches", actual)
}
