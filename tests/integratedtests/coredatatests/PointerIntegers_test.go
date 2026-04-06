package coredatatests

import (
	"sort"
	"testing"

	"github.com/alimtvnetwork/core/coredata"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_PointerIntegers_Len(t *testing.T) {
	for caseIndex, tc := range pointerIntegersLenTestCases {
		// Arrange
		var pi coredata.PointerIntegers
		if tc.ArrangeInput != nil {
			input := tc.ArrangeInput.(args.Map)
			count, _ := input.GetAsInt("count")
			ptrs := make([]*int, count)
			for i := range ptrs {
				v := i + 1
				ptrs[i] = &v
			}
			pi = coredata.PointerIntegers(ptrs)
		}

		// Act
		actual := args.Map{
			"length": pi.Len(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PointerIntegers_Less(t *testing.T) {
	// Arrange
	// Case 0: both non-nil
	{
		a, b := 3, 8
		pi := coredata.PointerIntegers{&a, &b}

	// Act
		actual := args.Map{
			"lessIJ": pi.Less(0, 1),
			"lessJI": pi.Less(1, 0),
		}

	// Assert
		pointerIntegersLessTestCases[0].ShouldBeEqualMap(t, 0, actual)
	}

	// Case 1: nil-i
	{
		b := 5
		pi := coredata.PointerIntegers{nil, &b}
		actual := args.Map{
			"result": pi.Less(0, 1),
		}
		pointerIntegersLessTestCases[1].ShouldBeEqualMap(t, 1, actual)
	}

	// Case 2: nil-j
	{
		a := 5
		pi := coredata.PointerIntegers{&a, nil}
		actual := args.Map{
			"result": pi.Less(0, 1),
		}
		pointerIntegersLessTestCases[2].ShouldBeEqualMap(t, 2, actual)
	}

	// Case 3: both nil
	{
		pi := coredata.PointerIntegers{nil, nil}
		actual := args.Map{
			"result": pi.Less(0, 1),
		}
		pointerIntegersLessTestCases[3].ShouldBeEqualMap(t, 3, actual)
	}
}

func Test_PointerIntegers_Sort(t *testing.T) {
	tc := pointerIntegersSortTestCases[0]

	// Arrange
	a, b, c := 5, 1, 3
	pi := coredata.PointerIntegers{&a, nil, &b, &c}

	// Act
	sort.Sort(pi)

	actual := args.Map{
		"firstIsNil": pi[0] == nil,
		"second":     *pi[1],
		"last":       *pi[len(pi)-1],
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
