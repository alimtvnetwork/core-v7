package corecmptests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ============================================================================
// Integer16
// ============================================================================

func Test_Integer16_Equal_FromInteger16Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16(5, 5) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16 returns Equal -- same values", actual)
}

func Test_Integer16_LeftLess_FromInteger16Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16(3, 5) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16 returns LeftLess -- left < right", actual)
}

func Test_Integer16_LeftGreater_FromInteger16Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16(7, 5) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16 returns LeftGreater -- left > right", actual)
}

// ============================================================================
// Integer64
// ============================================================================

func Test_Integer64_Equal_FromInteger16Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64(100, 100) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64 returns Equal -- same values", actual)
}

func Test_Integer64_LeftLess_FromInteger16Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64(50, 100) == corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64 returns LeftLess -- left < right", actual)
}

func Test_Integer64_LeftGreater_FromInteger16Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64(200, 100) == corecomparator.LeftGreater}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64 returns LeftGreater -- left > right", actual)
}

// ============================================================================
// IntegerPtr
// ============================================================================

func Test_IntegerPtr_BothNil_FromInteger16Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns Equal -- both nil", actual)
}

func Test_IntegerPtr_LeftNil_FromInteger16Equal(t *testing.T) {
	// Arrange
	v := 5

	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(nil, &v) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns NotEqual -- left nil", actual)
}

func Test_IntegerPtr_RightNil_FromInteger16Equal(t *testing.T) {
	// Arrange
	v := 5

	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(&v, nil) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns NotEqual -- right nil", actual)
}

func Test_IntegerPtr_BothEqual(t *testing.T) {
	// Arrange
	a, b := 5, 5

	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(&a, &b) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns Equal -- both 5", actual)
}

// ============================================================================
// BytePtr
// ============================================================================

func Test_BytePtr_BothNil_FromInteger16Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.BytePtr(nil, nil) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BytePtr returns Equal -- both nil", actual)
}

func Test_BytePtr_LeftNil_FromInteger16Equal(t *testing.T) {
	// Arrange
	v := byte(5)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(nil, &v) == corecomparator.NotEqual}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BytePtr returns NotEqual -- left nil", actual)
}

func Test_BytePtr_BothEqual(t *testing.T) {
	// Arrange
	a, b := byte(5), byte(5)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(&a, &b) == corecomparator.Equal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BytePtr returns Equal -- both 5", actual)
}

// ============================================================================
// IsIntegersEqual
// ============================================================================

func Test_IsIntegersEqual_BothNil_FromInteger16Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqual(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns true -- both nil", actual)
}

func Test_IsIntegersEqual_LeftNil_FromInteger16Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqual(nil, []int{1})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns false -- left nil", actual)
}

func Test_IsIntegersEqual_Same_FromInteger16Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 2})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns true -- same values", actual)
}

func Test_IsIntegersEqual_Different_FromInteger16Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 3})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns false -- different values", actual)
}

// ============================================================================
// IsStringsEqualPtr
// ============================================================================

func Test_IsStringsEqualPtr_BothNil_FromInteger16Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualPtr(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns true -- both nil", actual)
}

func Test_IsStringsEqualPtr_LeftNil_FromInteger16Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualPtr(nil, []string{"a"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns false -- left nil", actual)
}

func Test_IsStringsEqualPtr_DiffLen_FromInteger16Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns false -- different lengths", actual)
}

func Test_IsStringsEqualPtr_Same_FromInteger16Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualPtr([]string{"a", "b"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns true -- same values", actual)
}
