package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_StringsOnce_Core(t *testing.T) {
	for caseIndex, tc := range stringsOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringsOncePtr(func() []string { return initVal })

		// Act
		actual := args.Map{
			"length":     once.Length(),
			"isEmpty":    once.IsEmpty(),
			"hasAnyItem": once.HasAnyItem(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringsOnce_Contains(t *testing.T) {
	for caseIndex, tc := range stringsOnceContainsTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringsOncePtr(func() []string { return initVal })

		// Act
		actual := args.Map{
			"hasX":       once.Has("x"),
			"containsY":  once.IsContains("y"),
			"hasMissing": once.Has("missing"),
			"hasAllXY":   once.HasAll("x", "y"),
			"hasAllXW":   once.HasAll("x", "w"),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringsOnce_Sorted_FromStringsOnce(t *testing.T) {
	for caseIndex, tc := range stringsOnceSortedTestCases {
		// Arrange
		initVal := make([]string, len(tc.InitValue))
		copy(initVal, tc.InitValue)
		once := coreonce.NewStringsOncePtr(func() []string { return initVal })

		// Act
		sorted := once.Sorted()

		actual := args.Map{
			"first": sorted[0],
			"last":  sorted[len(sorted)-1],
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringsOnce_Maps(t *testing.T) {
	for caseIndex, tc := range stringsOnceMapTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringsOncePtr(func() []string { return initVal })

		// Act
		uniqueMap := once.UniqueMap()
		rangesMap := once.RangesMap()

		actual := args.Map{
			"uniqueLen":    len(uniqueMap),
			"rangesMapLen": len(rangesMap),
			"uniqueHasA":   uniqueMap["a"],
			"uniqueHasB":   uniqueMap["b"],
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringsOnce_IsEqual_FromStringsOnce(t *testing.T) {
	for caseIndex, tc := range stringsOnceIsEqualTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringsOncePtr(func() []string { return initVal })

		// Act
		actual := args.Map{
			"isEqualSame":    once.IsEqual("a", "b"),
			"isEqualDiff":    once.IsEqual("a", "c"),
			"isEqualDiffLen": once.IsEqual("a"),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringsOnce_Caching(t *testing.T) {
	for caseIndex, tc := range stringsOnceCachingTestCases {
		// Arrange
		callCount := 0
		initVal := tc.InitValue
		once := coreonce.NewStringsOncePtr(func() []string {
			callCount++

			return initVal
		})

		// Act
		_ = once.Value()
		_ = once.Value()

		actual := args.Map{
			"callCount": callCount,
			"length":    once.Length(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringsOnce_Json(t *testing.T) {
	for caseIndex, tc := range stringsOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringsOncePtr(func() []string { return initVal })

		// Act
		data, err := once.MarshalJSON()

		actual := args.Map{
			"noError":        err == nil,
			"marshaledValue": string(data),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringsOnce_Constructor(t *testing.T) {
	for caseIndex, tc := range stringsOnceConstructorTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringsOnce(func() []string { return initVal })

		// Act
		actual := args.Map{
			"length": once.Length(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
