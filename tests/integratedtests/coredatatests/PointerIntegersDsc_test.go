package coredatatests

import (
	"sort"
	"testing"

	"github.com/alimtvnetwork/core/coredata"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_PointerIntegersDsc_Len(t *testing.T) {
	for caseIndex, tc := range pointerIntegersDscLenTestCases {
		// Arrange
		var pi coredata.PointerIntegersDsc
		if tc.ArrangeInput != nil {
			input := tc.ArrangeInput.(args.Map)
			count, _ := input.GetAsInt("count")
			ptrs := make([]*int, count)
			for i := range ptrs {
				v := i + 1
				ptrs[i] = &v
			}
			pi = coredata.PointerIntegersDsc(ptrs)
		}

		// Act
		actual := args.Map{
			"length": pi.Len(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PointerIntegersDsc_Less(t *testing.T) {
	// Arrange
	// Case 0: both non-nil (3, 8) — descending so Less(0,1)=false, Less(1,0)=true
	{
		a, b := 3, 8
		pi := coredata.PointerIntegersDsc{&a, &b}

	// Act
		actual := args.Map{
			"lessIJ": pi.Less(0, 1),
			"lessJI": pi.Less(1, 0),
		}

	// Assert
		pointerIntegersDscLessTestCases[0].ShouldBeEqualMap(t, 0, actual)
	}

	// Case 1: nil-i returns false
	{
		b := 5
		pi := coredata.PointerIntegersDsc{nil, &b}
		actual := args.Map{
			"result": pi.Less(0, 1),
		}
		pointerIntegersDscLessTestCases[1].ShouldBeEqualMap(t, 1, actual)
	}

	// Case 2: nil-j returns true
	{
		a := 5
		pi := coredata.PointerIntegersDsc{&a, nil}
		actual := args.Map{
			"result": pi.Less(0, 1),
		}
		pointerIntegersDscLessTestCases[2].ShouldBeEqualMap(t, 2, actual)
	}

	// Case 3: both nil returns false
	{
		pi := coredata.PointerIntegersDsc{nil, nil}
		actual := args.Map{
			"result": pi.Less(0, 1),
		}
		pointerIntegersDscLessTestCases[3].ShouldBeEqualMap(t, 3, actual)
	}
}

func Test_PointerIntegersDsc_Sort(t *testing.T) {
	tc := pointerIntegersDscSortTestCases[0]

	// Arrange
	a, b, c := 1, 5, 3
	pi := coredata.PointerIntegersDsc{&a, nil, &b, &c}

	// Act
	sort.Sort(pi)

	actual := args.Map{
		"first":     *pi[0],
		"lastIsNil": pi[len(pi)-1] == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
