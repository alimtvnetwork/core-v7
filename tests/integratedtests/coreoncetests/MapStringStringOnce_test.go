package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_MapStringStringOnce_Core(t *testing.T) {
	for caseIndex, tc := range mapSSOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewMapStringStringOncePtr(func() map[string]string { return initVal })

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

func Test_MapStringStringOnce_Contains(t *testing.T) {
	for caseIndex, tc := range mapSSOnceContainsTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewMapStringStringOncePtr(func() map[string]string { return initVal })

		// Act
		actual := args.Map{
			"hasK1":      once.Has("k1"),
			"containsK2": once.IsContains("k2"),
			"isMissingX": once.IsMissing("x"),
			"hasAllK1K2": once.HasAll("k1", "k2"),
			"getK1":      once.GetValue("k1"),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MapStringStringOnce_KeysValues(t *testing.T) {
	for caseIndex, tc := range mapSSOnceKeysValuesTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewMapStringStringOncePtr(func() map[string]string { return initVal })

		// Act
		keys := once.AllKeys()
		values := once.AllValues()
		sortedKeys := once.AllKeysSorted()
		sortedValues := once.AllValuesSorted()

		actual := args.Map{
			"keysLen":          len(keys),
			"valuesLen":        len(values),
			"sortedFirstKey":   sortedKeys[0],
			"sortedLastKey":    sortedKeys[len(sortedKeys)-1],
			"sortedFirstValue": sortedValues[0],
			"sortedLastValue":  sortedValues[len(sortedValues)-1],
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MapStringStringOnce_IsEqual_FromMapStringStringOnce(t *testing.T) {
	for caseIndex, tc := range mapSSOnceIsEqualTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewMapStringStringOncePtr(func() map[string]string { return initVal })

		// Act
		actual := args.Map{
			"isEqualSame":    once.IsEqual(map[string]string{"a": "1"}),
			"isEqualDiffVal": once.IsEqual(map[string]string{"a": "2"}),
			"isEqualDiffKey": once.IsEqual(map[string]string{"b": "1"}),
			"isEqualDiffLen": once.IsEqual(map[string]string{"a": "1", "b": "2"}),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MapStringStringOnce_Caching(t *testing.T) {
	for caseIndex, tc := range mapSSOnceCachingTestCases {
		// Arrange
		callCount := 0
		initVal := tc.InitValue
		once := coreonce.NewMapStringStringOncePtr(func() map[string]string {
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

func Test_MapStringStringOnce_Json(t *testing.T) {
	for caseIndex, tc := range mapSSOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewMapStringStringOncePtr(func() map[string]string { return initVal })

		// Act
		data, err := once.MarshalJSON()

		actual := args.Map{
			"noError":             err == nil,
			"dataLengthAboveZero": len(data) > 0,
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MapStringStringOnce_Constructor(t *testing.T) {
	for caseIndex, tc := range mapSSOnceConstructorTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewMapStringStringOnce(func() map[string]string { return initVal })

		// Act
		actual := args.Map{
			"length": once.Length(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
