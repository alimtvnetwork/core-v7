package coredatatests

import (
	"sort"
	"testing"

	"github.com/alimtvnetwork/core/coredata"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_PointerStrings_Len(t *testing.T) {
	for caseIndex, tc := range pointerStringsLenTestCases {
		// Arrange
		var ps coredata.PointerStrings
		if tc.ArrangeInput != nil {
			input := tc.ArrangeInput.(args.Map)
			count, _ := input.GetAsInt("count")
			ptrs := make([]*string, count)
			for i := range ptrs {
				v := "item"
				ptrs[i] = &v
			}
			ps = coredata.PointerStrings(ptrs)
		}

		// Act
		actual := args.Map{
			"length": ps.Len(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PointerStrings_Less(t *testing.T) {
	// Arrange
	// Case 0: both non-nil
	{
		a, b := "alpha", "beta"
		ps := coredata.PointerStrings{&a, &b}

	// Act
		actual := args.Map{
			"less01": ps.Less(0, 1),
			"less10": ps.Less(1, 0),
		}

	// Assert
		pointerStringsLessTestCases[0].ShouldBeEqualMap(t, 0, actual)
	}

	// Case 1: nil first
	{
		b := "beta"
		ps := coredata.PointerStrings{nil, &b}
		actual := args.Map{
			"result": ps.Less(0, 1),
		}
		pointerStringsLessTestCases[1].ShouldBeEqualMap(t, 1, actual)
	}

	// Case 2: nil second
	{
		a := "alpha"
		ps := coredata.PointerStrings{&a, nil}
		actual := args.Map{
			"result": ps.Less(0, 1),
		}
		pointerStringsLessTestCases[2].ShouldBeEqualMap(t, 2, actual)
	}

	// Case 3: both nil
	{
		ps := coredata.PointerStrings{nil, nil}
		actual := args.Map{
			"result": ps.Less(0, 1),
		}
		pointerStringsLessTestCases[3].ShouldBeEqualMap(t, 3, actual)
	}
}

func Test_PointerStrings_Sort(t *testing.T) {
	tc := pointerStringsSortTestCases[0]

	// Arrange
	c, a, b := "charlie", "alpha", "beta"
	ps := coredata.PointerStrings{&c, nil, &a, &b}

	// Act
	sort.Sort(ps)

	actual := args.Map{
		"firstIsNil": ps[0] == nil,
		"second":     *ps[1],
		"third":      *ps[2],
		"fourth":     *ps[3],
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
