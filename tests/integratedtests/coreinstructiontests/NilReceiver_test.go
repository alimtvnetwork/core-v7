package coreinstructiontests

import "testing"

// ==========================================
// StringCompare — CaseNilSafe pattern
// ==========================================

// Note: StringCompare NilReceiver tests remain in StringCompare_test.go

// ==========================================
// IdentifiersWithGlobals — CaseNilSafe pattern
// ==========================================

func Test_IdentifiersWithGlobals_NilReceiver(t *testing.T) {
	for caseIndex, tc := range identifiersWithGlobalsNilSafeTestCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// FromTo — CaseNilSafe pattern
// ==========================================

func Test_FromTo_NilReceiver(t *testing.T) {
	for caseIndex, tc := range fromToNilSafeTestCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}
