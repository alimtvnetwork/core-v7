package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage16 — coreonce remaining gaps
//
// Target 1: MapStringStringOnce.JsonStringMust line 309-314
//   json.Marshal error → panic. Dead code — marshalling map[string]string
//   cannot fail.
//
// Target 2: StringsOnce.JsonStringMust line 251-256
//   Same pattern — marshalling []string cannot fail. Dead code.
// ══════════════════════════════════════════════════════════════════════════════

func Test_MapStringStringOnce_JsonStringMust_FromMapStringStringOnceJ(t *testing.T) {
	// Arrange
	once := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{
			"key1": "val1",
			"key2": "val2",
		}
	})

	// Act
	result := once.JsonStringMust()

	// Assert
	convey.Convey("MapStringStringOnce.JsonStringMust returns valid JSON", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
		convey.So(result, convey.ShouldContainSubstring, "key1")
	})
}

func Test_StringsOnce_JsonStringMust_FromMapStringStringOnceJ(t *testing.T) {
	// Arrange
	once := coreonce.NewStringsOnce(func() []string {
		return []string{"a", "b", "c"}
	})

	// Act
	result := once.JsonStringMust()

	// Assert
	convey.Convey("StringsOnce.JsonStringMust returns valid JSON", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
		convey.So(result, convey.ShouldContainSubstring, "a")
	})
}

// Coverage note: Both error branches are dead code — json.Marshal of
// map[string]string and []string cannot fail.
