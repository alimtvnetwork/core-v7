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

func Test_Src_IsEmptyPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcIsEmptyPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw, _ := input.Get("input")

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.IsEmptyPtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_Src_HasAnyItemPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcHasAnyItemPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		sliceRaw, _ := input.Get("input")

		var slice []string
		if sliceRaw != nil {
			slice = sliceRaw.([]string)
		}

		// Act
		result := stringslice.HasAnyItemPtr(slice)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"result": result,
		})
	}
}

func Test_Src_EmptyPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcEmptyPtrTestCases {
		// Arrange (no input)

		// Act
		result := stringslice.EmptyPtr()

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, args.Map{
			"length": len(result),
		})
	}
}

func Test_Src_LengthOfPointer_Verification(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		expect int
	}{
		{"nil", nil, 0},
		{"two items", []string{"a", "b"}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			input := tt.input

			// Act
			result := stringslice.LengthOfPointer(input)

			// Assert
			actual := args.Map{"result": result != tt.expect}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected", actual)
		})
	}
}

func Test_Src_Make_Verification(t *testing.T) {
	// Arrange
	expectedCap := 10

	// Act
	result := stringslice.Make(0, expectedCap)

	// Assert
	actual := args.Map{"result": cap(result) != expectedCap}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cap", actual)
}

func Test_Src_MakeDefault_Verification(t *testing.T) {
	// Arrange
	expectedCap := 5

	// Act
	result := stringslice.MakeDefault(expectedCap)

	// Assert
	actual := args.Map{"result": cap(result) < expectedCap}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cap >=", actual)
}

func Test_Src_MakePtr_Verification(t *testing.T) {
	// Arrange
	expectedCap := 10

	// Act
	result := stringslice.MakePtr(0, expectedCap)

	// Assert
	actual := args.Map{"result": cap(result) != expectedCap}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cap", actual)
}

func Test_Src_MakeDefaultPtr_Verification(t *testing.T) {
	// Arrange
	expectedCap := 5

	// Act
	result := stringslice.MakeDefaultPtr(expectedCap)

	// Assert
	actual := args.Map{"result": cap(result) < expectedCap}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cap >=", actual)
}

func Test_Src_MakeLen_Verification(t *testing.T) {
	// Arrange
	expectedLen := 5

	// Act
	result := stringslice.MakeLen(expectedLen)

	// Assert
	actual := args.Map{"result": len(result) != expectedLen}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected len", actual)
}

func Test_Src_MakeLenPtr_Verification(t *testing.T) {
	// Arrange
	expectedLen := 5

	// Act
	result := stringslice.MakeLenPtr(expectedLen)

	// Assert
	actual := args.Map{"result": len(result) != expectedLen}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected len", actual)
}

func Test_Src_SlicePtr_Verification(t *testing.T) {
	// Arrange
	input := []string{"a"}

	// Act
	result := stringslice.SlicePtr(input)

	// Assert
	actual := args.Map{"result": len(result) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	// Arrange
	var nilInput []string

	// Act
	result2 := stringslice.SlicePtr(nilInput)

	// Assert
	actual := args.Map{"result": len(result2) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Src_CloneSimpleSliceToPointers_Verification(t *testing.T) {
	// Arrange (nil)

	// Act
	result := stringslice.CloneSimpleSliceToPointers(nil)

	// Assert
	actual := args.Map{"result": result == nil || len(*result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil ptr with 0 len", actual)

	// Arrange (non-nil)
	input := []string{"a"}

	// Act
	result2 := stringslice.CloneSimpleSliceToPointers(input)

	// Assert
	actual := args.Map{"result": len(*result2) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// Suppresses unused import
var _ = fmt.Sprint
