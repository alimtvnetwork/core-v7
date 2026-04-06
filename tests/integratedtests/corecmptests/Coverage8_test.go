package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Integer — all branches ──

func Test_Cov8_Integer_Greater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer(10, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- greater", actual)
}

// ── Integer8 — greater branch ──

func Test_Cov8_Integer8_Greater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8(10, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- greater", actual)
}

func Test_Cov8_Integer8Ptr_LeftNil(t *testing.T) {
	// Arrange
	r := int8(5)

	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- left nil", actual)
}

func Test_Cov8_Integer8Ptr_Equal(t *testing.T) {
	// Arrange
	l, r := int8(5), int8(5)

	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns correct value -- equal", actual)
}

// ── Integer16 — all branches ──

func Test_Cov8_Integer16_Less(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16(3, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- less", actual)
}

func Test_Cov8_Integer16_Greater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16(10, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- greater", actual)
}

func Test_Cov8_Integer16Ptr_LeftNil(t *testing.T) {
	// Arrange
	r := int16(5)

	// Act
	actual := args.Map{"result": corecmp.Integer16Ptr(nil, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- left nil", actual)
}

func Test_Cov8_Integer16Ptr_Equal(t *testing.T) {
	// Arrange
	l, r := int16(5), int16(5)

	// Act
	actual := args.Map{"result": corecmp.Integer16Ptr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns correct value -- equal", actual)
}

// ── Integer32 — all branches ──

func Test_Cov8_Integer32_Less(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32(3, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- less", actual)
}

func Test_Cov8_Integer32_Greater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32(10, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- greater", actual)
}

func Test_Cov8_Integer32Ptr_LeftNil(t *testing.T) {
	// Arrange
	r := int32(5)

	// Act
	actual := args.Map{"result": corecmp.Integer32Ptr(nil, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- left nil", actual)
}

func Test_Cov8_Integer32Ptr_Equal(t *testing.T) {
	// Arrange
	l, r := int32(5), int32(5)

	// Act
	actual := args.Map{"result": corecmp.Integer32Ptr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns correct value -- equal", actual)
}

// ── Integer64 — greater branch ──

func Test_Cov8_Integer64_Greater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64(10, 5)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- greater", actual)
}

func Test_Cov8_Integer64Ptr_Equal(t *testing.T) {
	// Arrange
	l, r := int64(5), int64(5)

	// Act
	actual := args.Map{"result": corecmp.Integer64Ptr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns correct value -- equal", actual)
}

func Test_Cov8_Integer64Ptr_LeftNil(t *testing.T) {
	// Arrange
	r := int64(5)

	// Act
	actual := args.Map{"result": corecmp.Integer64Ptr(nil, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- left nil", actual)
}

// ── BytePtr — equal values ──

func Test_Cov8_BytePtr_Equal(t *testing.T) {
	// Arrange
	l, r := byte(5), byte(5)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "BytePtr returns correct value -- equal", actual)
}

func Test_Cov8_BytePtr_RightNil(t *testing.T) {
	// Arrange
	l := byte(5)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(&l, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- right nil", actual)
}

// ── Time ──

func Test_Cov8_Time_Equal(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"result": corecmp.Time(now, now)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- equal", actual)
}

func Test_Cov8_Time_Before(t *testing.T) {
	// Arrange
	t1 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

	// Act
	actual := args.Map{"result": corecmp.Time(t1, t2)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- before", actual)
}

func Test_Cov8_Time_After(t *testing.T) {
	// Arrange
	t1 := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	// Act
	actual := args.Map{"result": corecmp.Time(t1, t2)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- after", actual)
}

// ── TimePtr ──

func Test_Cov8_TimePtr_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.TimePtr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- both nil", actual)
}

func Test_Cov8_TimePtr_LeftNil(t *testing.T) {
	// Arrange
	r := time.Now()

	// Act
	actual := args.Map{"result": corecmp.TimePtr(nil, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- left nil", actual)
}

func Test_Cov8_TimePtr_RightNil(t *testing.T) {
	// Arrange
	l := time.Now()

	// Act
	actual := args.Map{"result": corecmp.TimePtr(&l, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- right nil", actual)
}

func Test_Cov8_TimePtr_Equal(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"result": corecmp.TimePtr(&now, &now)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "TimePtr returns correct value -- equal", actual)
}

// ── IsStringsEqualWithoutOrder ──

func Test_Cov8_IsStringsEqualWithoutOrder(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":   corecmp.IsStringsEqualWithoutOrder([]string{"b", "a"}, []string{"a", "b"}),
		"bothNil": corecmp.IsStringsEqualWithoutOrder(nil, nil),
		"leftNil": corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"}),
		"diffLen": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"}),
		"notEq":   corecmp.IsStringsEqualWithoutOrder([]string{"a", "b"}, []string{"a", "c"}),
	}

	// Assert
	expected := args.Map{
		"equal": true, "bothNil": true, "leftNil": false, "diffLen": false, "notEq": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- with args", actual)
}

// ── IsStringsEqual — not equal same length ──

func Test_Cov8_IsStringsEqual_NotEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"notEq":    corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "c"}),
		"rightNil": corecmp.IsStringsEqual([]string{"a"}, nil),
	}

	// Assert
	expected := args.Map{"notEq": false, "rightNil": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- not equal", actual)
}

// ── IsStringsEqualPtr — not equal ──

func Test_Cov8_IsStringsEqualPtr_NotEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"rightNil": corecmp.IsStringsEqualPtr([]string{"a"}, nil),
		"diffLen":  corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a", "b"}),
	}

	// Assert
	expected := args.Map{"rightNil": false, "diffLen": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns correct value -- not equal", actual)
}

// ── IsIntegersEqual — one nil ──

func Test_Cov8_IsIntegersEqual_OneNil(t *testing.T) {
	// Act
	actual := args.Map{
		"rightNil": corecmp.IsIntegersEqual([]int{1}, nil),
		"leftNil":  corecmp.IsIntegersEqual(nil, []int{1}),
	}

	// Assert
	expected := args.Map{"rightNil": false, "leftNil": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns nil -- one nil", actual)
}

// ── IsIntegersEqualPtr — right nil / diff len ──

func Test_Cov8_IsIntegersEqualPtr_RightNil(t *testing.T) {
	// Arrange
	l := []int{1}

	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&l, nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- right nil", actual)
}

func Test_Cov8_IsIntegersEqualPtr_DiffLen(t *testing.T) {
	// Arrange
	l := []int{1}
	r := []int{1, 2}

	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&l, &r)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns correct value -- diff len", actual)
}

// ── VersionSliceByte — all branches ──

func Test_Cov8_VersionSliceByte_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns nil -- both nil", actual)
}

func Test_Cov8_VersionSliceByte_LeftNil(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte(nil, []byte{1})}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns nil -- left nil", actual)
}

func Test_Cov8_VersionSliceByte_Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- equal", actual)
}

func Test_Cov8_VersionSliceByte_LeftLess(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- left less", actual)
}

func Test_Cov8_VersionSliceByte_LeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 3}, []byte{1, 2})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- left greater", actual)
}

func Test_Cov8_VersionSliceByte_ShorterLeft(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1}, []byte{1, 2})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- shorter left", actual)
}

func Test_Cov8_VersionSliceByte_ShorterRight(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- shorter right", actual)
}

// ── VersionSliceInteger — all branches ──

func Test_Cov8_VersionSliceInteger_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- both nil", actual)
}

func Test_Cov8_VersionSliceInteger_LeftNil(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger(nil, []int{1})}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- left nil", actual)
}

func Test_Cov8_VersionSliceInteger_Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 2})}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- equal", actual)
}

func Test_Cov8_VersionSliceInteger_LeftLess(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- left less", actual)
}

func Test_Cov8_VersionSliceInteger_LeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 3}, []int{1, 2})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- left greater", actual)
}

func Test_Cov8_VersionSliceInteger_ShorterLeft(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1}, []int{1, 2})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- shorter left", actual)
}

func Test_Cov8_VersionSliceInteger_LongerLeft(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- longer left", actual)
}

// ── AnyItem — inconclusive ──

func Test_Cov8_AnyItem_Inconclusive(t *testing.T) {
	// Different non-nil non-comparable values
	// Act
	actual := args.Map{"result": corecmp.AnyItem("hello", "world")}

	// Assert
	expected := args.Map{"result": corecomparator.Inconclusive}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- inconclusive", actual)
}

func Test_Cov8_AnyItem_RightNil(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem(5, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- right nil", actual)
}
