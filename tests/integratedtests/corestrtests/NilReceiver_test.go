package corestrtests

import "testing"

// ==========================================
// Hashset — CaseNilSafe pattern
// ==========================================

// Note: Hashset NilReceiver tests remain in Hashset_NilReceiver_testcases.go
//       with runner in the existing test file.

// ==========================================
// Hashmap.Clear nil — CaseNilSafe pattern
// ==========================================

func Test_Hashmap_Clear_NilReceiver(t *testing.T) {
	safeTest(t, "Test_Hashmap_Clear_NilReceiver", func() {
		for caseIndex, tc := range hashmapClearNilSafeTestCases {
		// Assert
			tc.ShouldBeSafe(t, caseIndex)
		}
	})
}
