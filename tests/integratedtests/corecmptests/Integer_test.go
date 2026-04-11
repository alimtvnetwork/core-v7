package corecmptests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Integer ──

func Test_Integer(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":   corecmp.Integer(5, 5),
		"less":    corecmp.Integer(3, 5),
		"greater": corecmp.Integer(7, 5),
	}

	// Assert
	expected := args.Map{
		"equal": corecomparator.Equal, "less": corecomparator.LeftLess, "greater": corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "Integer returns correct value -- with args", actual)
}

// ── IntegerPtr ──

func Test_IntegerPtr_BothNil_FromInteger(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- both nil", actual)
}

func Test_IntegerPtr_LeftNil_FromInteger(t *testing.T) {
	// Arrange
	r := 5

	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(nil, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns nil -- left nil", actual)
}

func Test_IntegerPtr_Equal_FromInteger(t *testing.T) {
	// Arrange
	l, r := 5, 5

	// Act
	actual := args.Map{"result": corecmp.IntegerPtr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "IntegerPtr returns correct value -- equal", actual)
}

// ── Integer8 / Integer8Ptr ──

func Test_Integer8(t *testing.T) {
	// Act
	actual := args.Map{
		"equal": corecmp.Integer8(5, 5),
		"less":  corecmp.Integer8(3, 5),
	}

	// Assert
	expected := args.Map{
		"equal": corecomparator.Equal,
		"less": corecomparator.LeftLess,
	}
	expected.ShouldBeEqual(t, 0, "Integer8 returns correct value -- with args", actual)
}

func Test_Integer8Ptr_BothNil_FromInteger(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer8Ptr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer8Ptr returns nil -- both nil", actual)
}

// ── Integer16 / Integer16Ptr ──

func Test_Integer16(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.Integer16(5, 5)}

	// Assert
	expected := args.Map{"equal": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16 returns correct value -- with args", actual)
}

func Test_Integer16Ptr_BothNil_FromInteger(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer16Ptr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer16Ptr returns nil -- both nil", actual)
}

// ── Integer32 / Integer32Ptr ──

func Test_Integer32(t *testing.T) {
	// Act
	actual := args.Map{"equal": corecmp.Integer32(5, 5)}

	// Assert
	expected := args.Map{"equal": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32 returns correct value -- with args", actual)
}

func Test_Integer32Ptr_BothNil_FromInteger(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer32Ptr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer32Ptr returns nil -- both nil", actual)
}

// ── Integer64 / Integer64Ptr ──

func Test_Integer64(t *testing.T) {
	// Act
	actual := args.Map{
		"equal": corecmp.Integer64(5, 5),
		"less": corecmp.Integer64(3, 5),
	}

	// Assert
	expected := args.Map{
		"equal": corecomparator.Equal,
		"less": corecomparator.LeftLess,
	}
	expected.ShouldBeEqual(t, 0, "Integer64 returns correct value -- with args", actual)
}

func Test_Integer64Ptr_BothNil_FromInteger(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.Integer64Ptr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- both nil", actual)
}

func Test_Integer64Ptr_RightNil_FromInteger(t *testing.T) {
	// Arrange
	l := int64(5)

	// Act
	actual := args.Map{"result": corecmp.Integer64Ptr(&l, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "Integer64Ptr returns nil -- right nil", actual)
}

// ── Byte / BytePtr ──

func Test_Byte(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":   corecmp.Byte(5, 5),
		"less":    corecmp.Byte(3, 5),
		"greater": corecmp.Byte(7, 5),
	}

	// Assert
	expected := args.Map{
		"equal": corecomparator.Equal,
		"less": corecomparator.LeftLess,
		"greater": corecomparator.LeftGreater,
	}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- with args", actual)
}

func Test_BytePtr_BothNil_FromInteger(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.BytePtr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- both nil", actual)
}

func Test_BytePtr_LeftNil_FromInteger(t *testing.T) {
	// Arrange
	r := byte(5)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(nil, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "BytePtr returns nil -- left nil", actual)
}

// ── IsIntegersEqual / IsIntegersEqualPtr ──

func Test_IsIntegersEqual_FromInteger(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":   corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 2}),
		"notEq":   corecmp.IsIntegersEqual([]int{1, 2}, []int{1, 3}),
		"diffLen": corecmp.IsIntegersEqual([]int{1}, []int{1, 2}),
		"bothNil": corecmp.IsIntegersEqual(nil, nil),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEq": false,
		"diffLen": false,
		"bothNil": true,
	}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqual returns correct value -- with args", actual)
}

func Test_IsIntegersEqualPtr_FromInteger(t *testing.T) {
	// Arrange
	left := []int{1, 2}
	right := []int{1, 2}
	rightSingle := []int{1}

	// Act
	actual := args.Map{
		"equal":   corecmp.IsIntegersEqualPtr(&left, &right),
		"bothNil": corecmp.IsIntegersEqualPtr(nil, nil),
		"leftNil": corecmp.IsIntegersEqualPtr(nil, &rightSingle),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"bothNil": true,
		"leftNil": false,
	}
	expected.ShouldBeEqual(t, 0, "IsIntegersEqualPtr returns correct value -- with args", actual)
}

// ── IsStringsEqual / IsStringsEqualPtr ──

func Test_IsStringsEqual_FromInteger(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":   corecmp.IsStringsEqual([]string{"a", "b"}, []string{"a", "b"}),
		"bothNil": corecmp.IsStringsEqual(nil, nil),
		"leftNil": corecmp.IsStringsEqual(nil, []string{"a"}),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"bothNil": true,
		"leftNil": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns correct value -- with args", actual)
}

func Test_IsStringsEqualPtr_FromInteger(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":   corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a"}),
		"bothNil": corecmp.IsStringsEqualPtr(nil, nil),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"bothNil": true,
	}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns correct value -- with args", actual)
}

// ── AnyItem ──

func Test_AnyItem(t *testing.T) {
	// Act
	actual := args.Map{
		"bothNil": corecmp.AnyItem(nil, nil),
		"leftNil": corecmp.AnyItem(nil, 5),
		"equal":   corecmp.AnyItem(5, 5),
		"notEq":   corecmp.AnyItem(5, 6),
	}

	// Assert
	expected := args.Map{
		"bothNil": corecomparator.Equal,
		"leftNil": corecomparator.NotEqual,
		"equal":   corecomparator.Equal,
		"notEq":   corecomparator.Inconclusive,
	}
	expected.ShouldBeEqual(t, 0, "AnyItem returns correct value -- with args", actual)
}

// ── VersionSliceInteger — RightNil ──

func Test_VersionSliceInteger_RightNil_FromInteger(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1}, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns nil -- right nil", actual)
}
