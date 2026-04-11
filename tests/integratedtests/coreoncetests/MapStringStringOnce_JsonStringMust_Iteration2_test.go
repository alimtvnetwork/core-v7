package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// MapStringStringOnce.JsonStringMust() — success path
// Covers MapStringStringOnce.go L306-317
// Note: Error panic path (L309-314) is unreachable — MarshalJSON on
// map[string]string never returns an error in standard Go.
// ══════════════════════════════════════════════════════════════════════════════

func Test_MapStringStringOnce_JsonStringMust_FromMapStringStringOnceJ_FromMapStringStringOnceJ(t *testing.T) {
	// Arrange
	tc := cov13MapStringStringOnceJsonStringMustTestCase
	m := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{
			"key1": "val1",
		}
	})

	// Act
	noPanic := true
	var jsonStr string
	func() {
		defer func() {
			if r := recover(); r != nil {
				noPanic = false
			}
		}()
		jsonStr = m.JsonStringMust()
	}()

	actual := args.Map{
		"nonEmpty": jsonStr != "",
		"noPanic":  noPanic,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// StringsOnce.JsonStringMust() — success path
// Covers StringsOnce.go L248-258
// Note: Error panic path (L251-256) is unreachable — MarshalJSON on
// []string never returns an error in standard Go.
// ══════════════════════════════════════════════════════════════════════════════

func Test_StringsOnce_JsonStringMust_FromMapStringStringOnceJ_FromMapStringStringOnceJ(t *testing.T) {
	// Arrange
	tc := cov13StringsOnceJsonStringMustTestCase
	s := coreonce.NewStringsOnce(func() []string {
		return []string{"a", "b", "c"}
	})

	// Act
	noPanic := true
	var jsonStr string
	func() {
		defer func() {
			if r := recover(); r != nil {
				noPanic = false
			}
		}()
		jsonStr = s.JsonStringMust()
	}()

	actual := args.Map{
		"nonEmpty": jsonStr != "",
		"noPanic":  noPanic,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
