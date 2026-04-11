package corevalidatortests

import "testing"

// ==========================================
// SliceValidator — CaseNilSafe pattern
// ==========================================

func Test_SliceValidator_NilReceiver(t *testing.T) {
	for caseIndex, tc := range sliceValidatorNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// SliceValidators — CaseNilSafe pattern
// ==========================================

func Test_SliceValidators_NilReceiver(t *testing.T) {
	for caseIndex, tc := range sliceValidatorsNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// TextValidator — CaseNilSafe pattern
// ==========================================

func Test_TextValidator_NilReceiver(t *testing.T) {
	for caseIndex, tc := range textValidatorNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// TextValidators — CaseNilSafe pattern
// ==========================================

func Test_TextValidators_NilReceiver(t *testing.T) {
	for caseIndex, tc := range textValidatorsNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// BaseLinesValidators — CaseNilSafe pattern
// ==========================================

func Test_BaseLinesValidators_NilReceiver(t *testing.T) {
	for caseIndex, tc := range baseLinesValidatorsNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// LinesValidators — CaseNilSafe pattern
// ==========================================

func Test_LinesValidators_NilReceiver_FromNilReceiver(t *testing.T) {
	for caseIndex, tc := range linesValidatorsNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// LineValidator — CaseNilSafe pattern
// ==========================================

func Test_LineValidator_NilReceiver(t *testing.T) {
	for caseIndex, tc := range lineValidatorNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}
