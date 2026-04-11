package corecmptests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Byte_Direct_Comparison(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Byte(5, 5) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same bytes should be Equal", actual)
	actual = args.Map{"result": corecmp.Byte(3, 7) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "3 < 7 should be LeftLess", actual)
	actual = args.Map{"result": corecmp.Byte(7, 3) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "7 > 3 should be LeftGreater", actual)
}

func Test_BytePtr_Direct_Comparison(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.BytePtr(nil, nil) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both nil should be Equal", actual)
	b := byte(5)
	actual = args.Map{"result": corecmp.BytePtr(nil, &b) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "left nil should be NotEqual", actual)
	actual = args.Map{"result": corecmp.BytePtr(&b, nil) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "right nil should be NotEqual", actual)
	b2 := byte(5)
	actual = args.Map{"result": corecmp.BytePtr(&b, &b2) != corecomparator.Equal}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same values should be Equal", actual)
}

func Test_Integer64_Direct_Comparison(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64(10, 10) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same should be Equal", actual)
	actual = args.Map{"result": corecmp.Integer64(5, 10) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "5 < 10 should be LeftLess", actual)
	actual = args.Map{"result": corecmp.Integer64(10, 5) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "10 > 5 should be LeftGreater", actual)
}

func Test_IntegerPtr_Direct_Comparison(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(nil, nil) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both nil should be Equal", actual)
	val := 5
	actual = args.Map{"result": corecmp.IntegerPtr(nil, &val) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "left nil should be NotEqual", actual)
	val2 := 5
	actual = args.Map{"result": corecmp.IntegerPtr(&val, &val2) != corecomparator.Equal}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same values should be Equal", actual)
}

func Test_IsIntegersEqual_Direct_Verification(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqual(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
	actual = args.Map{"result": corecmp.IsIntegersEqual(nil, []int{1})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil should not be equal", actual)
	actual = args.Map{"result": corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 2})}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same slices should be equal", actual)
	actual = args.Map{"result": corecmp.IsIntegersEqual([]int{1}, []int{2})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different values should not be equal", actual)
}

func Test_VersionSliceByte_Direct_Comparison(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte(nil, nil) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both nil should be Equal", actual)
	actual = args.Map{"result": corecmp.VersionSliceByte(nil, []byte{1}) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil should be NotEqual", actual)
	actual = args.Map{"result": corecmp.VersionSliceByte([]byte{1}, nil) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-nil vs nil should be NotEqual", actual)
	actual = args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 2}) != corecomparator.Equal}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same should be Equal", actual)
	actual = args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 3}) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "1.2 < 1.3 should be LeftLess", actual)
	actual = args.Map{"result": corecmp.VersionSliceByte([]byte{1, 3}, []byte{1, 2}) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "1.3 > 1.2 should be LeftGreater", actual)
	actual = args.Map{"result": corecmp.VersionSliceByte([]byte{1}, []byte{1, 2}) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "shorter left should be LeftLess", actual)
	actual = args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1}) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "longer left should be LeftGreater", actual)
}

func Test_IsStringsEqualWithoutOrder_Direct_Verification(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
	actual = args.Map{"result": corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil should not be equal", actual)
	actual = args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"b", "a"}, []string{"a", "b"})}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same items different order should be equal", actual)
	actual = args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"b"})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different items should not be equal", actual)
	actual = args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different lengths should not be equal", actual)
}

func Test_Integer8_Direct_Comparison(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8(5, 5) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same should be Equal", actual)
	actual = args.Map{"result": corecmp.Integer8(3, 7) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "3 < 7 should be LeftLess", actual)
	actual = args.Map{"result": corecmp.Integer8(7, 3) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "7 > 3 should be LeftGreater", actual)
}

func Test_Integer16_Direct_Comparison(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16(5, 5) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same should be Equal", actual)
	actual = args.Map{"result": corecmp.Integer16(3, 7) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "3 < 7 should be LeftLess", actual)
}

func Test_Integer32_Direct_Comparison(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32(5, 5) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same should be Equal", actual)
	actual = args.Map{"result": corecmp.Integer32(3, 7) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "3 < 7 should be LeftLess", actual)
}

func Test_Integer8Ptr_Direct_Comparison(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, nil) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both nil should be Equal", actual)
	v := int8(5)
	actual = args.Map{"result": corecmp.Integer8Ptr(nil, &v) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "left nil should be NotEqual", actual)
}

func Test_Integer16Ptr_Direct_Comparison(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16Ptr(nil, nil) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both nil should be Equal", actual)
	v := int16(5)
	actual = args.Map{"result": corecmp.Integer16Ptr(nil, &v) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "left nil should be NotEqual", actual)
}

func Test_Integer32Ptr_Direct_Comparison(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32Ptr(nil, nil) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both nil should be Equal", actual)
	v := int32(5)
	actual = args.Map{"result": corecmp.Integer32Ptr(nil, &v) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "left nil should be NotEqual", actual)
}

func Test_Integer64Ptr_Direct_Comparison(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64Ptr(nil, nil) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both nil should be Equal", actual)
	v := int64(5)
	actual = args.Map{"result": corecmp.Integer64Ptr(nil, &v) != corecomparator.NotEqual}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "left nil should be NotEqual", actual)
}

func Test_IsIntegersEqualPtr_Verification(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqualPtr(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
	a := []int{1, 2}
	actual = args.Map{"result": corecmp.IsIntegersEqualPtr(&a, nil)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one nil should not be equal", actual)
	b := []int{1, 2}
	actual = args.Map{"result": corecmp.IsIntegersEqualPtr(&a, &b)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same slices should be equal", actual)
}

func Test_AnyItem_Verification(t *testing.T) {
	// Arrange
	result := corecmp.AnyItem("hello", "hello")

	// Act
	actual := args.Map{"result": result != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same strings should be Equal", actual)
	result2 := corecmp.AnyItem(1, 2)
	actual = args.Map{"result": result2 == corecomparator.Equal}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different ints should not be Equal", actual)
}

func Test_VersionSliceInteger_Direct_Comparison(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger(nil, nil) != corecomparator.Equal}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both nil should be Equal", actual)
	actual = args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 2}) != corecomparator.Equal}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same should be Equal", actual)
	actual = args.Map{"result": corecmp.VersionSliceInteger([]int{1}, []int{1, 2}) != corecomparator.LeftLess}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "shorter should be LeftLess", actual)
}
