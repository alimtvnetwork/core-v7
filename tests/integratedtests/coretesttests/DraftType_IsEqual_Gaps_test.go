package coretesttests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage3 — coretests remaining gaps
//
// Target 1: BaseTestCaseAssertions.go:88-92 — isFailed log branch (test failure only)
// Target 2: DraftType.go:148 — isIncludingInnerFields && f1String differs (unexported)
// Target 3: DraftType.go:174,184 — json.Marshal panic (dead code)
// Target 4: SkipOnUnix.go:12-14 — platform-dependent (Unix only)
// ══════════════════════════════════════════════════════════════════════════════

func Test_DraftType_IsEqual_WithInnerFields_Same(t *testing.T) {
	// Arrange — exported fields match, inner fields are zero-valued and equal
	left := &coretests.DraftType{
		SampleString1: "sample1",
		SampleString2: "sample2",
		SampleInteger: 10,
	}
	right := &coretests.DraftType{
		SampleString1: "sample1",
		SampleString2: "sample2",
		SampleInteger: 10,
	}

	// Act
	result := left.IsEqual(true, right)

	// Assert
	convey.Convey("DraftType.IsEqual returns true when all fields match including inner", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}

// Coverage note:
// - BaseTestCaseAssertions.go:88-92 — only reachable on assertion failure inside convey.
//   Not safely testable. Documented as defensive logging gap.
// - DraftType.go:148,152 — requires different unexported fields (f1String, f2Integer).
//   Cannot be set from external tests. Needs internal test or is accepted gap.
// - DraftType.go:174,184 — json.Marshal panic. Dead code.
// - SkipOnUnix.go:12-14 — platform-dependent. Only runs on Unix.
