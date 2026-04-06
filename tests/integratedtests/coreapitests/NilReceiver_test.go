package coreapitests

import "testing"

// ==========================================
// TypedSimpleGenericRequest — CaseNilSafe pattern
// ==========================================

// Note: TypedSimpleGenericRequest NilReceiver tests remain in TypedConversions_test.go

// ==========================================
// PageRequest — CaseNilSafe pattern
// ==========================================

func Test_PageRequest_NilReceiver(t *testing.T) {
	for caseIndex, tc := range pageRequestNilSafeTestCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}
