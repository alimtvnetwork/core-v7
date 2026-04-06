package coredynamictests

import "testing"

// ==========================================
// Dynamic — CaseNilSafe pattern
// ==========================================

func Test_Dynamic_NilReceiver(t *testing.T) {
	for caseIndex, tc := range dynamicNilSafeTestCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// LeftRight — CaseNilSafe pattern
// ==========================================

func Test_LeftRight_NilReceiver(t *testing.T) {
	for caseIndex, tc := range leftRightNilSafeTestCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// CastedResult — CaseNilSafe pattern
// ==========================================

func Test_CastedResult_NilReceiver(t *testing.T) {
	for caseIndex, tc := range castedResultNilSafeTestCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// MapAnyItems — CaseNilSafe pattern
// ==========================================

func Test_MapAnyItems_NilReceiver(t *testing.T) {
	for caseIndex, tc := range mapAnyItemsNilSafeTestCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}
