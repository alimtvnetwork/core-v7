package coretaskinfotests

import "testing"

// ==========================================
// Info — CaseNilSafe pattern
// ==========================================

func Test_Info_NilReceiver(t *testing.T) {
	for caseIndex, tc := range infoNilSafeTestCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}
