package stringslicetests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Src_IsEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range srcIsEmptyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw, _ := input.Get("input")

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.IsEmpty(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_Src_HasAnyItem_Verification(t *testing.T) {
	for caseIndex, tc := range srcHasAnyItemTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw, _ := input.Get("input")

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.HasAnyItem(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_Src_Empty_Verification(t *testing.T) {
	for caseIndex, tc := range srcEmptyTestCases {
		// Arrange (no input)

		// Act
		result := stringslice.Empty()

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"length": len(result),
		})
	}
}

	// Act
	result := stringslice.Make(0, expectedCap)

	// Assert
	if cap(result) != expectedCap {
		t.Fatalf("expected cap %d, got %d", expectedCap, cap(result))
	}
}

func Test_Src_MakeDefault_Verification(t *testing.T) {
	// Arrange
	expectedCap := 5

	// Act
	result := stringslice.MakeDefault(expectedCap)

	// Assert
	if cap(result) < expectedCap {
		t.Fatalf("expected cap >= %d, got %d", expectedCap, cap(result))
	}
}
	// Arrange
	expectedLen := 5

	// Act
	result := stringslice.MakeLen(expectedLen)

	// Assert
	if len(result) != expectedLen {
		t.Fatalf("expected len %d, got %d", expectedLen, len(result))
	}
}
	// Arrange (nil)

	// Act
	result := stringslice.CloneSimpleSliceToPointers(nil)

	// Assert
	if result == nil || len(*result) != 0 {
		t.Fatal("expected non-nil ptr with 0 len")
	}

	// Arrange (non-nil)
	input := []string{"a"}

	// Act
	result2 := stringslice.CloneSimpleSliceToPointers(input)

	// Assert
	if len(*result2) != 1 {
		t.Fatalf("expected 1, got %d", len(*result2))
	}
}

// Suppresses unused import
var _ = fmt.Sprint
