package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── AnyItem ──

func Test_AnyItem_BothNil_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- both nil", actual)
}

func Test_AnyItem_LeftNil(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem(nil, "a") == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- left nil", actual)
}

func Test_AnyItem_RightNil(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem("a", nil) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns nil -- right nil", actual)
}

func Test_AnyItem_Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem("a", "a") == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- equal", actual)
}

func Test_AnyItem_Inconclusive(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem("a", "b") == corecomparator.Inconclusive}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- inconclusive", actual)
}

// ── Byte / BytePtr ──

func Test_Byte_Equal_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Byte(1, 1) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- equal", actual)
}

func Test_Byte_LeftLess(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Byte(1, 2) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- left less", actual)
}

func Test_Byte_LeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Byte(2, 1) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- left greater", actual)
}

func Test_BytePtr_BothNil_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.BytePtr(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- both nil", actual)
}

func Test_BytePtr_LeftNil_FromAnyItemBothNilV3(t *testing.T) {
	// Arrange
	b := byte(1)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(nil, &b) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- left nil", actual)
}

func Test_BytePtr_RightNil_FromAnyItemBothNilV3(t *testing.T) {
	// Arrange
	b := byte(1)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(&b, nil) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- right nil", actual)
}

func Test_BytePtr_Equal(t *testing.T) {
	// Arrange
	a, b := byte(5), byte(5)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(&a, &b) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BytePtr returns correct value -- equal", actual)
}

// ── Integer / IntegerPtr ──

func Test_Integer_Equal_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer(5, 5) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- equal", actual)
}

func Test_Integer_LeftLess(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer(3, 5) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- left less", actual)
}

func Test_Integer_LeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer(5, 3) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- left greater", actual)
}

func Test_IntegerPtr_BothNil_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- both nil", actual)
}

func Test_IntegerPtr_LeftNil_FromAnyItemBothNilV3(t *testing.T) {
	// Arrange
	v := 5

	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(nil, &v) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- left nil", actual)
}

func Test_IntegerPtr_RightNil_FromAnyItemBothNilV3(t *testing.T) {
	// Arrange
	v := 5

	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(&v, nil) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- right nil", actual)
}

// ── Integer8 / Integer8Ptr ──

func Test_Integer8_Equal_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8(5, 5) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- equal", actual)
}

func Test_Integer8_LeftLess(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8(3, 5) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- left less", actual)
}

func Test_Integer8_LeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8(5, 3) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- left greater", actual)
}

func Test_Integer8Ptr_BothNil_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- both nil", actual)
}

func Test_Integer8Ptr_OneNil(t *testing.T) {
	// Arrange
	v := int8(5)

	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, &v) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- one nil", actual)
}

// ── Integer16 / Integer16Ptr ──

func Test_Integer16_Equal_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16(5, 5) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- equal", actual)
}

func Test_Integer16_LeftLess(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16(3, 5) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- left less", actual)
}

func Test_Integer16_LeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16(5, 3) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- left greater", actual)
}

func Test_Integer16Ptr_BothNil_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16Ptr(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- both nil", actual)
}

func Test_Integer16Ptr_OneNil(t *testing.T) {
	// Arrange
	v := int16(5)

	// Act
	actual := args.Map{"result": corecmp.Integer16Ptr(nil, &v) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- one nil", actual)
}

// ── Integer32 / Integer32Ptr ──

func Test_Integer32_Equal_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32(5, 5) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- equal", actual)
}

func Test_Integer32_LeftLess(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32(3, 5) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- left less", actual)
}

func Test_Integer32_LeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32(5, 3) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- left greater", actual)
}

func Test_Integer32Ptr_BothNil_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32Ptr(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- both nil", actual)
}

func Test_Integer32Ptr_OneNil(t *testing.T) {
	// Arrange
	v := int32(5)

	// Act
	actual := args.Map{"result": corecmp.Integer32Ptr(nil, &v) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- one nil", actual)
}

// ── Integer64 / Integer64Ptr ──

func Test_Integer64_Equal_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64(5, 5) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- equal", actual)
}

func Test_Integer64_LeftLess(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64(3, 5) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- left less", actual)
}

func Test_Integer64_LeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64(5, 3) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- left greater", actual)
}

func Test_Integer64Ptr_BothNil_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64Ptr(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- both nil", actual)
}

func Test_Integer64Ptr_OneNil(t *testing.T) {
	// Arrange
	v := int64(5)

	// Act
	actual := args.Map{"result": corecmp.Integer64Ptr(nil, &v) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- one nil", actual)
}

// ── Time / TimePtr ──

func Test_Time_Equal_FromAnyItemBothNilV3(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"result": corecmp.Time(now, now) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- equal", actual)
}

func Test_Time_LeftLess(t *testing.T) {
	// Arrange
	now := time.Now()
	later := now.Add(time.Hour)

	// Act
	actual := args.Map{"result": corecmp.Time(now, later) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- left less", actual)
}

func Test_Time_LeftGreater(t *testing.T) {
	// Arrange
	now := time.Now()
	earlier := now.Add(-time.Hour)

	// Act
	actual := args.Map{"result": corecmp.Time(now, earlier) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Time returns correct value -- left greater", actual)
}

func Test_TimePtr_BothNil_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.TimePtr(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- both nil", actual)
}

func Test_TimePtr_OneNil(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"result": corecmp.TimePtr(nil, &now) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "TimePtr returns nil -- one nil", actual)
}

func Test_TimePtr_Equal(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"result": corecmp.TimePtr(&now, &now) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "TimePtr returns correct value -- equal", actual)
}

// ── IsStringsEqual / IsStringsEqualPtr / IsStringsEqualWithoutOrder ──

func Test_IsStringsEqual_BothNil_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- both nil", actual)
}

func Test_IsStringsEqual_OneNil(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual(nil, []string{"a"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns nil -- one nil", actual)
}

func Test_IsStringsEqual_DifferentLength(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- different length", actual)
}

func Test_IsStringsEqual_Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- equal", actual)
}

func Test_IsStringsEqual_NotEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "c"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- not equal", actual)
}

func Test_IsStringsEqualPtr_BothNil_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualPtr(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns nil -- both nil", actual)
}

func Test_IsStringsEqualPtr_OneNil(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualPtr(nil, []string{"a"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns nil -- one nil", actual)
}

func Test_IsStringsEqualPtr_DiffLen_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns correct value -- diff len", actual)
}

func Test_IsStringsEqualWithoutOrder_Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"b", "a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- equal", actual)
}

func Test_IsStringsEqualWithoutOrder_BothNil_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns nil -- both nil", actual)
}

func Test_IsStringsEqualWithoutOrder_OneNil(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns nil -- one nil", actual)
}

func Test_IsStringsEqualWithoutOrder_DiffLen_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- diff len", actual)
}

func Test_IsStringsEqualWithoutOrder_NotEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a", "b"}, []string{"a", "c"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns non-empty -- not equal", actual)
}

// ── IsIntegersEqual / IsIntegersEqualPtr ──

func Test_IsIntegersEqual_BothNil_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqual(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns nil -- both nil", actual)
}

func Test_IsIntegersEqual_OneNil(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqual(nil, []int{1})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns nil -- one nil", actual)
}

func Test_IsIntegersEqual_Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 2})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns correct value -- equal", actual)
}

func Test_IsIntegersEqualPtr_BothNil_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- both nil", actual)
}

func Test_IsIntegersEqualPtr_OneNil(t *testing.T) {
	// Arrange
	a := []int{1}

	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&a, nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns nil -- one nil", actual)
}

func Test_IsIntegersEqualPtr_DiffLen_FromAnyItemBothNilV3(t *testing.T) {
	// Arrange
	a := []int{1}
	b := []int{1, 2}

	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&a, &b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns correct value -- diff len", actual)
}

// ── VersionSliceByte ──

func Test_VersionSliceByte_Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 2}) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- equal", actual)
}

func Test_VersionSliceByte_LeftLess(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 3}) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- left less", actual)
}

func Test_VersionSliceByte_LeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 3}, []byte{1, 2}) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- left greater", actual)
}

func Test_VersionSliceByte_BothNil_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns nil -- both nil", actual)
}

func Test_VersionSliceByte_OneNil(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte(nil, []byte{1}) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns nil -- one nil", actual)
}

func Test_VersionSliceByte_ShorterLeft_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1}, []byte{1, 2}) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- shorter left", actual)
}

func Test_VersionSliceByte_LongerLeft(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1}) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns correct value -- longer left", actual)
}

// ── VersionSliceInteger ──

func Test_VersionSliceInteger_Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 2}) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- equal", actual)
}

func Test_VersionSliceInteger_LeftLess(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 3}) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- left less", actual)
}

func Test_VersionSliceInteger_LeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 3}, []int{1, 2}) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- left greater", actual)
}

func Test_VersionSliceInteger_BothNil_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- both nil", actual)
}

func Test_VersionSliceInteger_OneNil(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger(nil, []int{1}) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- one nil", actual)
}

func Test_VersionSliceInteger_ShorterLeft_FromAnyItemBothNilV3(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1}, []int{1, 2}) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- shorter left", actual)
}

func Test_VersionSliceInteger_LongerLeft(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1}) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns correct value -- longer left", actual)
}
