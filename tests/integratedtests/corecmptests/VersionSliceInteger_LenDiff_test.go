package corecmptests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_VersionSliceInteger_LenDiff_Cov2(t *testing.T) {
	// Act
	actual := args.Map{
		"shorterVsLonger": fmt.Sprintf("%v", corecmp.VersionSliceInteger([]int{1}, []int{1, 2})),
		"longerVsShorter": fmt.Sprintf("%v", corecmp.VersionSliceInteger([]int{1, 2}, []int{1})),
	}

	// Assert
	expected := args.Map{
		"shorterVsLonger": fmt.Sprintf("%v", corecomparator.LeftLess),
		"longerVsShorter": fmt.Sprintf("%v", corecomparator.LeftGreater),
	}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger_LenDiff returns correct value -- with args", actual)
}

func Test_VersionSliceInteger_LeftNil_Cov2(t *testing.T) {
	// Act
	actual := args.Map{
		"result": fmt.Sprintf("%v", corecmp.VersionSliceInteger([]int{1}, nil)),
	}

	// Assert
	expected := args.Map{
		"result": fmt.Sprintf("%v", corecomparator.NotEqual),
	}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger_LeftNil returns nil -- with args", actual)
}

func Test_VersionSliceByte_LeftNil_Cov2(t *testing.T) {
	// Act
	actual := args.Map{
		"result": fmt.Sprintf("%v", corecmp.VersionSliceByte([]byte{1}, nil)),
	}

	// Assert
	expected := args.Map{
		"result": fmt.Sprintf("%v", corecomparator.NotEqual),
	}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte_LeftNil returns nil -- with args", actual)
}

func Test_IsStringsEqualPtr_SameLenDiffContent_Cov2(t *testing.T) {
	// Arrange
	a := []string{"a", "b"}
	b := []string{"a", "c"}

	// Act
	actual := args.Map{"isEqual": corecmp.IsStringsEqualPtr(a, b)}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr_SameLenDiffContent returns correct value -- with args", actual)
}

func Test_IsStringsEqualPtr_DiffLen_Cov2(t *testing.T) {
	// Arrange
	a := []string{"a"}
	b := []string{"a", "b"}

	// Act
	actual := args.Map{"isEqual": corecmp.IsStringsEqualPtr(a, b)}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr_DiffLen returns correct value -- with args", actual)
}

func Test_IsStringsEqualPtr_LeftNil_Cov2(t *testing.T) {
	// Arrange
	b := []string{"a"}

	// Act
	actual := args.Map{"isEqual": corecmp.IsStringsEqualPtr(nil, b)}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr_LeftNil returns nil -- with args", actual)
}

func Test_IsIntegersEqualPtr_SameLenDiff_Cov2(t *testing.T) {
	// Arrange
	a := []int{1, 2}
	b := []int{1, 3}

	// Act
	actual := args.Map{"isEqual": corecmp.IsIntegersEqualPtr(&a, &b)}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr_SameLenDiff returns correct value -- with args", actual)
}

func Test_IsIntegersEqualPtr_DiffLen_Cov2(t *testing.T) {
	// Arrange
	a := []int{1}
	b := []int{1, 2}

	// Act
	actual := args.Map{"isEqual": corecmp.IsIntegersEqualPtr(&a, &b)}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr_DiffLen returns correct value -- with args", actual)
}

func Test_IsIntegersEqualPtr_LeftNil_Cov2(t *testing.T) {
	// Arrange
	b := []int{1}

	// Act
	actual := args.Map{"isEqual": corecmp.IsIntegersEqualPtr(nil, &b)}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr_LeftNil returns nil -- with args", actual)
}

func Test_IsStringsEqualWithoutOrder_DiffLen_Cov2(t *testing.T) {
	// Act
	actual := args.Map{"isEqual": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder_DiffLen returns non-empty -- with args", actual)
}

func Test_IsStringsEqualWithoutOrder_LeftNil_Cov2(t *testing.T) {
	// Act
	actual := args.Map{"isEqual": corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"})}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder_LeftNil returns nil -- with args", actual)
}

func Test_IsStringsEqualWithoutOrder_RightNil_Cov2(t *testing.T) {
	// Act
	actual := args.Map{"isEqual": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, nil)}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder_RightNil returns nil -- with args", actual)
}

func Test_BytePtr_LeftNilRightNotNil_Cov2(t *testing.T) {
	// Arrange
	b := byte(5)

	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", corecmp.BytePtr(nil, &b))}

	// Assert
	expected := args.Map{"result": fmt.Sprintf("%v", corecomparator.NotEqual)}
	expected.ShouldBeEqual(t, 0, "BytePtr_LeftNilRightNotNil returns nil -- with args", actual)
}

func Test_BytePtr_LeftNotNilRightNil_Cov2(t *testing.T) {
	// Arrange
	a := byte(5)

	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", corecmp.BytePtr(&a, nil))}

	// Assert
	expected := args.Map{"result": fmt.Sprintf("%v", corecomparator.NotEqual)}
	expected.ShouldBeEqual(t, 0, "BytePtr_LeftNotNilRightNil returns nil -- with args", actual)
}

func Test_BytePtr_LeftLess_Cov2(t *testing.T) {
	// Arrange
	a, b := byte(3), byte(5)

	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", corecmp.BytePtr(&a, &b))}

	// Assert
	expected := args.Map{"result": fmt.Sprintf("%v", corecomparator.LeftLess)}
	expected.ShouldBeEqual(t, 0, "BytePtr_LeftLess returns correct value -- with args", actual)
}

func Test_IntegerPtr_LeftNil_Cov2(t *testing.T) {
	// Arrange
	b := 5

	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", corecmp.IntegerPtr(nil, &b))}

	// Assert
	expected := args.Map{"result": fmt.Sprintf("%v", corecomparator.NotEqual)}
	expected.ShouldBeEqual(t, 0, "IntegerPtr_LeftNil returns nil -- with args", actual)
}

func Test_IntegerPtr_LeftLess_Cov2(t *testing.T) {
	// Arrange
	a, b := 3, 5

	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", corecmp.IntegerPtr(&a, &b))}

	// Assert
	expected := args.Map{"result": fmt.Sprintf("%v", corecomparator.LeftLess)}
	expected.ShouldBeEqual(t, 0, "IntegerPtr_LeftLess returns correct value -- with args", actual)
}

func Test_TimePtr_LeftNil_Cov2(t *testing.T) {
	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", corecmp.TimePtr(nil, nil))}

	// Assert
	expected := args.Map{"result": fmt.Sprintf("%v", corecomparator.Equal)}
	expected.ShouldBeEqual(t, 0, "TimePtr_BothNil returns nil -- with args", actual)
}

func Test_Integer16Ptr_LeftNil_Cov2(t *testing.T) {
	// Arrange
	b := int16(5)

	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", corecmp.Integer16Ptr(nil, &b))}

	// Assert
	expected := args.Map{"result": fmt.Sprintf("%v", corecomparator.NotEqual)}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr_LeftNil returns nil -- with args", actual)
}

func Test_Integer16Ptr_LeftLess_Cov2(t *testing.T) {
	// Arrange
	a, b := int16(3), int16(5)

	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", corecmp.Integer16Ptr(&a, &b))}

	// Assert
	expected := args.Map{"result": fmt.Sprintf("%v", corecomparator.LeftLess)}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr_LeftLess returns correct value -- with args", actual)
}

func Test_Integer32Ptr_LeftNil_Cov2(t *testing.T) {
	// Arrange
	b := int32(5)

	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", corecmp.Integer32Ptr(nil, &b))}

	// Assert
	expected := args.Map{"result": fmt.Sprintf("%v", corecomparator.NotEqual)}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr_LeftNil returns nil -- with args", actual)
}

func Test_Integer32Ptr_LeftLess_Cov2(t *testing.T) {
	// Arrange
	a, b := int32(3), int32(5)

	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", corecmp.Integer32Ptr(&a, &b))}

	// Assert
	expected := args.Map{"result": fmt.Sprintf("%v", corecomparator.LeftLess)}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr_LeftLess returns correct value -- with args", actual)
}

func Test_Integer64Ptr_LeftLess_Cov2(t *testing.T) {
	// Arrange
	a, b := int64(3), int64(5)

	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", corecmp.Integer64Ptr(&a, &b))}

	// Assert
	expected := args.Map{"result": fmt.Sprintf("%v", corecomparator.LeftLess)}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr_LeftLess returns correct value -- with args", actual)
}

func Test_Integer8Ptr_LeftLess_Cov2(t *testing.T) {
	// Arrange
	a, b := int8(3), int8(5)

	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", corecmp.Integer8Ptr(&a, &b))}

	// Assert
	expected := args.Map{"result": fmt.Sprintf("%v", corecomparator.LeftLess)}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr_LeftLess returns correct value -- with args", actual)
}

func Test_IsIntegersEqual_LeftNil_Cov2(t *testing.T) {
	// Act
	actual := args.Map{"isEqual": corecmp.IsIntegersEqual(nil, []int{1})}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual_LeftNil returns nil -- with args", actual)
}

func Test_IsIntegersEqual_RightNil_Cov2(t *testing.T) {
	// Act
	actual := args.Map{"isEqual": corecmp.IsIntegersEqual([]int{1}, nil)}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual_RightNil returns nil -- with args", actual)
}

func Test_IsStringsEqual_LeftNil_Cov2(t *testing.T) {
	// Act
	actual := args.Map{"isEqual": corecmp.IsStringsEqual(nil, []string{"a"})}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual_LeftNil returns nil -- with args", actual)
}

func Test_IsStringsEqual_RightNil_Cov2(t *testing.T) {
	// Act
	actual := args.Map{"isEqual": corecmp.IsStringsEqual([]string{"a"}, nil)}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual_RightNil returns nil -- with args", actual)
}
