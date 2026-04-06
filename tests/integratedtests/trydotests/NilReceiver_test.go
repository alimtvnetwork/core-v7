package trydotests

import "testing"

// ==========================================
// WrappedErr — CaseNilSafe pattern
// ==========================================

func Test_WrappedErr_NilReceiver(t *testing.T) {
	for caseIndex, tc := range wrappedErrNilSafeTestCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}
