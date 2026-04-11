package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage2 — corecmp remaining 9 lines (all dead-code NotEqual fallbacks)
// ══════════════════════════════════════════════════════════════════════════════

// The uncovered lines are unreachable NotEqual fallback returns after
// exhaustive if-else-if chains covering ==, <, >. For integers/bytes/time,
// if a == b, a < b, and a > b are ALL false, execution is logically impossible.
// These are documented as accepted dead code.

// However, we can verify the three reachable branches are all covered:

func Test_Byte_AllBranches_FromByteAllBranchesAllGa(t *testing.T) {
	// Arrange / Act / Assert
	actual := args.Map{
		"equal":   corecmp.Byte(5, 5),
		"less":    corecmp.Byte(3, 5),
		"greater": corecmp.Byte(5, 3),
	}
	expected := args.Map{
		"equal":   corecomparator.Equal,
		"less":    corecomparator.LeftLess,
		"greater": corecomparator.LeftGreater,
	}
	actual.ShouldBeEqual(t, 1, "Byte all branches", expected)
}

func Test_Integer_AllBranches(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":   corecmp.Integer(5, 5),
		"less":    corecmp.Integer(3, 5),
		"greater": corecmp.Integer(5, 3),
	}

	// Assert
	expected := args.Map{
		"equal":   corecomparator.Equal,
		"less":    corecomparator.LeftLess,
		"greater": corecomparator.LeftGreater,
	}
	actual.ShouldBeEqual(t, 1, "Integer all branches", expected)
}

func Test_Integer16_AllBranches_FromByteAllBranchesAllGa(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":   corecmp.Integer16(5, 5),
		"less":    corecmp.Integer16(3, 5),
		"greater": corecmp.Integer16(5, 3),
	}

	// Assert
	expected := args.Map{
		"equal":   corecomparator.Equal,
		"less":    corecomparator.LeftLess,
		"greater": corecomparator.LeftGreater,
	}
	actual.ShouldBeEqual(t, 1, "Integer16 all branches", expected)
}

func Test_Integer32_AllBranches_FromByteAllBranchesAllGa(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":   corecmp.Integer32(5, 5),
		"less":    corecmp.Integer32(3, 5),
		"greater": corecmp.Integer32(5, 3),
	}

	// Assert
	expected := args.Map{
		"equal":   corecomparator.Equal,
		"less":    corecomparator.LeftLess,
		"greater": corecomparator.LeftGreater,
	}
	actual.ShouldBeEqual(t, 1, "Integer32 all branches", expected)
}

func Test_Integer64_AllBranches_FromByteAllBranchesAllGa(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":   corecmp.Integer64(5, 5),
		"less":    corecmp.Integer64(3, 5),
		"greater": corecmp.Integer64(5, 3),
	}

	// Assert
	expected := args.Map{
		"equal":   corecomparator.Equal,
		"less":    corecomparator.LeftLess,
		"greater": corecomparator.LeftGreater,
	}
	actual.ShouldBeEqual(t, 1, "Integer64 all branches", expected)
}

func Test_Integer8_AllBranches_FromByteAllBranchesAllGa(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":   corecmp.Integer8(5, 5),
		"less":    corecmp.Integer8(3, 5),
		"greater": corecmp.Integer8(5, 3),
	}

	// Assert
	expected := args.Map{
		"equal":   corecomparator.Equal,
		"less":    corecomparator.LeftLess,
		"greater": corecomparator.LeftGreater,
	}
	actual.ShouldBeEqual(t, 1, "Integer8 all branches", expected)
}

func Test_Time_AllBranches_FromByteAllBranchesAllGa(t *testing.T) {
	// Arrange
	now := time.Now()
	past := now.Add(-time.Hour)
	future := now.Add(time.Hour)

	// Act / Assert
	actual := args.Map{
		"equal":   corecmp.Time(now, now),
		"less":    corecmp.Time(past, future),
		"greater": corecmp.Time(future, past),
	}
	expected := args.Map{
		"equal":   corecomparator.Equal,
		"less":    corecomparator.LeftLess,
		"greater": corecomparator.LeftGreater,
	}
	actual.ShouldBeEqual(t, 1, "Time all branches", expected)
}

func Test_VersionSliceByte_AllBranches_FromByteAllBranchesAllGa(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":      corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 2}),
		"leftLess":   corecmp.VersionSliceByte([]byte{1}, []byte{1, 2}),
		"leftGreater": corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2}),
	}

	// Assert
	expected := args.Map{
		"equal":      corecomparator.Equal,
		"leftLess":   corecomparator.LeftLess,
		"leftGreater": corecomparator.LeftGreater,
	}
	actual.ShouldBeEqual(t, 1, "VersionSliceByte all branches", expected)
}

func Test_VersionSliceInteger_AllBranches_FromByteAllBranchesAllGa(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":      corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 2}),
		"leftLess":   corecmp.VersionSliceInteger([]int{1}, []int{1, 2}),
		"leftGreater": corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2}),
	}

	// Assert
	expected := args.Map{
		"equal":      corecomparator.Equal,
		"leftLess":   corecomparator.LeftLess,
		"leftGreater": corecomparator.LeftGreater,
	}
	actual.ShouldBeEqual(t, 1, "VersionSliceInteger all branches", expected)
}
