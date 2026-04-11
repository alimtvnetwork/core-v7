package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── AnyItem ──

func Test_AnyItem_BothNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- both nil", actual)
}

func Test_AnyItem_LeftNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem(nil, 42) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- left nil", actual)
}

func Test_AnyItem_RightNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem(42, nil) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- right nil", actual)
}

func Test_AnyItem_Equal_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem(42, 42) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- equal", actual)
}

func Test_AnyItem_Inconclusive_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem(42, 99) == corecomparator.Inconclusive}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- inconclusive", actual)
}

// ── Integer8 ──

func Test_Integer8_Equal_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8(5, 5) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- equal", actual)
}

func Test_Integer8_LeftLess_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8(3, 5) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- left less", actual)
}

func Test_Integer8_LeftGreater_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8(7, 5) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- left greater", actual)
}

// ── Integer32 ──

func Test_Integer32_Equal_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32(5, 5) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- equal", actual)
}

func Test_Integer32_LeftLess_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32(3, 5) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- left less", actual)
}

func Test_Integer32_LeftGreater_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32(7, 5) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- left greater", actual)
}

// ── Integer8Ptr ──

func Test_Integer8Ptr_BothNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- both nil", actual)
}

func Test_Integer8Ptr_LeftNil_FromAnyItemBothNilV4(t *testing.T) {
	// Arrange
	v := int8(5)

	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, &v) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- left nil", actual)
}

func Test_Integer8Ptr_Equal(t *testing.T) {
	// Arrange
	a, b := int8(5), int8(5)

	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(&a, &b) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns correct value -- equal", actual)
}

// ── Integer16Ptr ──

func Test_Integer16Ptr_BothNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16Ptr(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- both nil", actual)
}

func Test_Integer16Ptr_LeftNil_FromAnyItemBothNilV4(t *testing.T) {
	// Arrange
	v := int16(5)

	// Act
	actual := args.Map{"result": corecmp.Integer16Ptr(nil, &v) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- left nil", actual)
}

func Test_Integer16Ptr_Equal(t *testing.T) {
	// Arrange
	a, b := int16(5), int16(5)

	// Act
	actual := args.Map{"result": corecmp.Integer16Ptr(&a, &b) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns correct value -- equal", actual)
}

// ── Integer32Ptr ──

func Test_Integer32Ptr_BothNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32Ptr(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- both nil", actual)
}

func Test_Integer32Ptr_LeftNil_FromAnyItemBothNilV4(t *testing.T) {
	// Arrange
	v := int32(5)

	// Act
	actual := args.Map{"result": corecmp.Integer32Ptr(nil, &v) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- left nil", actual)
}

func Test_Integer32Ptr_Equal(t *testing.T) {
	// Arrange
	a, b := int32(5), int32(5)

	// Act
	actual := args.Map{"result": corecmp.Integer32Ptr(&a, &b) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns correct value -- equal", actual)
}

// ── Integer64Ptr ──

func Test_Integer64Ptr_BothNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64Ptr(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- both nil", actual)
}

func Test_Integer64Ptr_LeftNil_FromAnyItemBothNilV4(t *testing.T) {
	// Arrange
	v := int64(5)

	// Act
	actual := args.Map{"result": corecmp.Integer64Ptr(nil, &v) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- left nil", actual)
}

func Test_Integer64Ptr_Equal(t *testing.T) {
	// Arrange
	a, b := int64(5), int64(5)

	// Act
	actual := args.Map{"result": corecmp.Integer64Ptr(&a, &b) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns correct value -- equal", actual)
}

// ── Time ──

func Test_Time_Equal_FromAnyItemBothNilV4(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"result": corecmp.Time(now, now) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- equal", actual)
}

func Test_Time_LeftLess_FromAnyItemBothNilV4(t *testing.T) {
	// Arrange
	t1 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

	// Act
	actual := args.Map{"result": corecmp.Time(t1, t2) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- left less", actual)
}

func Test_Time_LeftGreater_FromAnyItemBothNilV4(t *testing.T) {
	// Arrange
	t1 := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	// Act
	actual := args.Map{"result": corecmp.Time(t1, t2) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- left greater", actual)
}

// ── TimePtr ──

func Test_TimePtr_BothNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.TimePtr(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- both nil", actual)
}

func Test_TimePtr_LeftNil_FromAnyItemBothNilV4(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"result": corecmp.TimePtr(nil, &now) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- left nil", actual)
}

func Test_TimePtr_RightNil_FromAnyItemBothNilV4(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"result": corecmp.TimePtr(&now, nil) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- right nil", actual)
}

func Test_TimePtr_Equal_FromAnyItemBothNilV4(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"result": corecmp.TimePtr(&now, &now) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "TimePtr returns correct value -- equal", actual)
}

// ── IsIntegersEqualPtr ──

func Test_IsIntegersEqualPtr_BothNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- both nil", actual)
}

func Test_IsIntegersEqualPtr_LeftNil_FromAnyItemBothNilV4(t *testing.T) {
	// Arrange
	right := []int{1}

	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(nil, &right)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- left nil", actual)
}

func Test_IsIntegersEqualPtr_DiffLen_FromAnyItemBothNilV4(t *testing.T) {
	// Arrange
	left := []int{1}
	right := []int{1, 2}

	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&left, &right)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns correct value -- diff len", actual)
}

func Test_IsIntegersEqualPtr_Same(t *testing.T) {
	// Arrange
	left := []int{1, 2}
	right := []int{1, 2}

	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&left, &right)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns correct value -- same", actual)
}

// ── IsStringsEqual ──

func Test_IsStringsEqual_BothNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- both nil", actual)
}

func Test_IsStringsEqual_LeftNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual(nil, []string{"a"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- left nil", actual)
}

func Test_IsStringsEqual_DiffLen_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- diff len", actual)
}

func Test_IsStringsEqual_Same_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- same", actual)
}

func Test_IsStringsEqual_Different(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "c"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- different", actual)
}

// ── IsStringsEqualWithoutOrder ──

func Test_IsStringsEqualWithoutOrder_BothNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns nil -- both nil", actual)
}

func Test_IsStringsEqualWithoutOrder_LeftNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns nil -- left nil", actual)
}

func Test_IsStringsEqualWithoutOrder_DiffLen_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- diff len", actual)
}

func Test_IsStringsEqualWithoutOrder_SameOrder(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a", "b"}, []string{"b", "a"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- same unordered", actual)
}

// ── VersionSliceByte ──

func Test_VersionSliceByte_BothNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns nil -- both nil", actual)
}

func Test_VersionSliceByte_LeftNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte(nil, []byte{1}) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns nil -- left nil", actual)
}

func Test_VersionSliceByte_Equal_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 2}) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- equal", actual)
}

func Test_VersionSliceByte_LeftLess_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 3}) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- left less", actual)
}

func Test_VersionSliceByte_LeftGreater_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 3}, []byte{1, 2}) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- left greater", actual)
}

func Test_VersionSliceByte_DiffLen_LeftShorter(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1}, []byte{1, 2}) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- left shorter", actual)
}

func Test_VersionSliceByte_DiffLen_LeftLonger(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1}) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- left longer", actual)
}

// ── VersionSliceInteger ──

func Test_VersionSliceInteger_BothNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- both nil", actual)
}

func Test_VersionSliceInteger_LeftNil_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger(nil, []int{1}) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- left nil", actual)
}

func Test_VersionSliceInteger_Equal_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 2}) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- equal", actual)
}

func Test_VersionSliceInteger_LeftLess_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 3}) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- left less", actual)
}

func Test_VersionSliceInteger_LeftGreater_FromAnyItemBothNilV4(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 3}, []int{1, 2}) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- left greater", actual)
}

func Test_VersionSliceInteger_DiffLen_LeftShorter(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1}, []int{1, 2}) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- left shorter", actual)
}

func Test_VersionSliceInteger_DiffLen_LeftLonger(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1}) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- left longer", actual)
}
