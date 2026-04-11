package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Byte ──

func Test_Byte_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Byte(5, 5) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "5 == 5", actual)
	actual = args.Map{"result": corecmp.Byte(3, 5) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "3 < 5", actual)
	actual = args.Map{"result": corecmp.Byte(5, 3) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "5 > 3", actual)
}

func Test_BytePtr_Coverage(t *testing.T) {
	// Arrange
	a, b := byte(5), byte(3)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(&a, &a) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same should be equal", actual)
	actual = args.Map{"result": corecmp.BytePtr(&a, &b) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "5 > 3", actual)
	actual = args.Map{"result": corecmp.BytePtr(nil, nil) != corecomparator.Equal}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
	actual = args.Map{"result": corecmp.BytePtr(nil, &a) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one nil should be not equal", actual)
	actual = args.Map{"result": corecmp.BytePtr(&a, nil) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one nil should be not equal", actual)
}

// ── Integer ──

func Test_Integer_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer(5, 5) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "5 == 5", actual)
	actual = args.Map{"result": corecmp.Integer(3, 5) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "3 < 5", actual)
	actual = args.Map{"result": corecmp.Integer(5, 3) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "5 > 3", actual)
}

func Test_IntegerPtr_Coverage(t *testing.T) {
	// Arrange
	a, b := 5, 3

	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(&a, &a) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same should be equal", actual)
	actual = args.Map{"result": corecmp.IntegerPtr(nil, nil) != corecomparator.Equal}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
	actual = args.Map{"result": corecmp.IntegerPtr(nil, &a) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one nil should be not equal", actual)
	actual = args.Map{"result": corecmp.IntegerPtr(&a, &b) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "5 > 3", actual)
}

// ── Integer64 ──

func Test_Integer64_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64(5, 5) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "equal", actual)
	actual = args.Map{"result": corecmp.Integer64(3, 5) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "less", actual)
	actual = args.Map{"result": corecmp.Integer64(5, 3) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "greater", actual)
}

// ── Integer16, Integer32, Integer8 ──

func Test_Integer16_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16(5, 5) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "equal", actual)
	actual = args.Map{"result": corecmp.Integer16(3, 5) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "less", actual)
	actual = args.Map{"result": corecmp.Integer16(5, 3) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "greater", actual)
}

func Test_Integer32_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32(5, 5) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "equal", actual)
	actual = args.Map{"result": corecmp.Integer32(3, 5) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "less", actual)
	actual = args.Map{"result": corecmp.Integer32(5, 3) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "greater", actual)
}

func Test_Integer8_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8(5, 5) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "equal", actual)
	actual = args.Map{"result": corecmp.Integer8(3, 5) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "less", actual)
	actual = args.Map{"result": corecmp.Integer8(5, 3) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "greater", actual)
}

// ── Ptr variants for 16/32/64/8 ──

func Test_Integer16Ptr_Coverage(t *testing.T) {
	// Arrange
	a, b := int16(5), int16(3)

	// Act
	actual := args.Map{"result": corecmp.Integer16Ptr(&a, &a) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "equal", actual)
	actual = args.Map{"result": corecmp.Integer16Ptr(nil, nil) != corecomparator.Equal}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil equal", actual)
	actual = args.Map{"result": corecmp.Integer16Ptr(nil, &a) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil not equal", actual)
	actual = args.Map{"result": corecmp.Integer16Ptr(&a, &b) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "greater", actual)
}

func Test_Integer32Ptr_Coverage(t *testing.T) {
	// Arrange
	a, b := int32(5), int32(3)

	// Act
	actual := args.Map{"result": corecmp.Integer32Ptr(&a, &a) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "equal", actual)
	actual = args.Map{"result": corecmp.Integer32Ptr(nil, nil) != corecomparator.Equal}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil equal", actual)
	actual = args.Map{"result": corecmp.Integer32Ptr(nil, &a) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil not equal", actual)
	actual = args.Map{"result": corecmp.Integer32Ptr(&a, &b) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "greater", actual)
}

func Test_Integer64Ptr_Coverage(t *testing.T) {
	// Arrange
	a, b := int64(5), int64(3)

	// Act
	actual := args.Map{"result": corecmp.Integer64Ptr(&a, &a) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "equal", actual)
	actual = args.Map{"result": corecmp.Integer64Ptr(nil, nil) != corecomparator.Equal}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil equal", actual)
	actual = args.Map{"result": corecmp.Integer64Ptr(nil, &a) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil not equal", actual)
	actual = args.Map{"result": corecmp.Integer64Ptr(&a, &b) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "greater", actual)
}

func Test_Integer8Ptr_Coverage(t *testing.T) {
	// Arrange
	a, b := int8(5), int8(3)

	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(&a, &a) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "equal", actual)
	actual = args.Map{"result": corecmp.Integer8Ptr(nil, nil) != corecomparator.Equal}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil equal", actual)
	actual = args.Map{"result": corecmp.Integer8Ptr(nil, &a) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil not equal", actual)
	actual = args.Map{"result": corecmp.Integer8Ptr(&a, &b) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "greater", actual)
}

// ── Time / TimePtr ──

func Test_Time_Coverage(t *testing.T) {
	// Arrange
	now := time.Now()
	earlier := now.Add(-time.Hour)
	later := now.Add(time.Hour)

	// Act
	actual := args.Map{"result": corecmp.Time(now, now) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same time should be equal", actual)
	actual = args.Map{"result": corecmp.Time(earlier, later) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "earlier < later", actual)
	actual = args.Map{"result": corecmp.Time(later, earlier) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "later > earlier", actual)
}

func Test_TimePtr_Coverage(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"result": corecmp.TimePtr(&now, &now) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "equal", actual)
	actual = args.Map{"result": corecmp.TimePtr(nil, nil) != corecomparator.Equal}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil equal", actual)
	actual = args.Map{"result": corecmp.TimePtr(nil, &now) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil not equal", actual)
}

// ── IsStringsEqual / IsStringsEqualPtr / IsStringsEqualWithoutOrder ──

func Test_IsStringsEqual_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same should be equal", actual)
	actual = args.Map{"result": corecmp.IsStringsEqual([]string{"a"}, []string{"b"})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different should not be equal", actual)
	actual = args.Map{"result": corecmp.IsStringsEqual([]string{"a"}, []string{"a", "b"})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different length should not be equal", actual)
	actual = args.Map{"result": corecmp.IsStringsEqual(nil, nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
	actual = args.Map{"result": corecmp.IsStringsEqual(nil, []string{})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs empty should not be equal", actual)
}

func Test_IsStringsEqualPtr_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualPtr(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
	a := []string{"a"}
	actual = args.Map{"result": corecmp.IsStringsEqualPtr(nil, a)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil should not be equal", actual)
	actual = args.Map{"result": corecmp.IsStringsEqualPtr(a, a)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same should be equal", actual)
}

func Test_IsStringsEqualWithoutOrder_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"b", "a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same items different order should be equal", actual)
	actual = args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"b"})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different items should not be equal", actual)
	actual = args.Map{"result": corecmp.IsStringsEqualWithoutOrder(nil, nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
	actual = args.Map{"result": corecmp.IsStringsEqualWithoutOrder(nil, []string{})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs empty should not be equal", actual)
}

// ── IsIntegersEqual / IsIntegersEqualPtr ──

func Test_IsIntegersEqual_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 2})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same should be equal", actual)
	actual = args.Map{"result": corecmp.IsIntegersEqual([]int{1}, []int{2})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different should not be equal", actual)
	actual = args.Map{"result": corecmp.IsIntegersEqual(nil, nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
	actual = args.Map{"result": corecmp.IsIntegersEqual(nil, []int{})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs empty should not be equal", actual)
}

func Test_IsIntegersEqualPtr_Coverage(t *testing.T) {
	// Arrange
	a := []int{1, 2}
	b := []int{1, 2}
	c := []int{3}

	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(&a, &b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same should be equal", actual)
	actual = args.Map{"result": corecmp.IsIntegersEqualPtr(&a, &c)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different should not be equal", actual)
	actual = args.Map{"result": corecmp.IsIntegersEqualPtr(nil, nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
	actual = args.Map{"result": corecmp.IsIntegersEqualPtr(nil, &a)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil should not be equal", actual)
}

// ── AnyItem ──

func Test_AnyItem_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.AnyItem(nil, nil) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
	actual = args.Map{"result": corecmp.AnyItem(nil, 42) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil should be not equal", actual)
	actual = args.Map{"result": corecmp.AnyItem(42, nil) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-nil vs nil should be not equal", actual)
	actual = args.Map{"result": corecmp.AnyItem(42, 42) != corecomparator.Equal}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same should be equal", actual)
	actual = args.Map{"result": corecmp.AnyItem(42, 43) != corecomparator.Inconclusive}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different should be inconclusive", actual)
}

// ── VersionSliceByte ──

func Test_VersionSliceByte_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte(nil, nil) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both nil equal", actual)
	actual = args.Map{"result": corecmp.VersionSliceByte(nil, []byte{1}) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil not equal", actual)
	actual = args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 2}) != corecomparator.Equal}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same should be equal", actual)
	actual = args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 3}) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "1.2 < 1.3", actual)
	actual = args.Map{"result": corecmp.VersionSliceByte([]byte{1, 3}, []byte{1, 2}) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "1.3 > 1.2", actual)
	actual = args.Map{"result": corecmp.VersionSliceByte([]byte{1}, []byte{1, 2}) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "shorter version less", actual)
	actual = args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1}) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "longer version greater", actual)
}

// ── VersionSliceInteger ──

func Test_VersionSliceInteger_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger(nil, nil) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both nil equal", actual)
	actual = args.Map{"result": corecmp.VersionSliceInteger(nil, []int{1}) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil not equal", actual)
	actual = args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 2}) != corecomparator.Equal}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same should be equal", actual)
	actual = args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 3}) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "1.2 < 1.3", actual)
	actual = args.Map{"result": corecmp.VersionSliceInteger([]int{1, 3}, []int{1, 2}) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "1.3 > 1.2", actual)
}
