package stringslicetests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage13 — stringslice remaining gaps
//
// Target 1: MergeSlicesOfSlices.go:13-15 — duplicate len==0 check (dead code,
//   line 7 already catches len==0)
//
// Target 2: RegexTrimmedSplitNonEmptyAll.go:17-19 — split returns empty
//   regexp.Split always returns at least one element, so len==0 is unreachable.
//   Dead code.
// ══════════════════════════════════════════════════════════════════════════════

func Test_MergeSlicesOfSlices_AllEmpty(t *testing.T) {
	// Arrange — pass non-nil but empty slices
	slice1 := []string{}
	slice2 := []string{}

	// Act
	result := stringslice.MergeSlicesOfSlices(slice1, slice2)

	// Assert
	convey.Convey("MergeSlicesOfSlices returns empty when all input slices empty", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

func Test_RegexTrimmedSplitNonEmptyAll_NoMatch(t *testing.T) {
	// Arrange — regex that doesn't match
	re := regexp.MustCompile(`\d+`)

	// Act
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, "hello world")

	// Assert
	convey.Convey("RegexTrimmedSplitNonEmptyAll returns trimmed result when no match", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

// Coverage note: Both uncovered branches are dead code:
// MergeSlicesOfSlices line 13: duplicate of line 7
// RegexTrimmedSplitNonEmptyAll line 17: regexp.Split always returns ≥1 element
