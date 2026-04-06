package coredynamictests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Test: IsInvalid
// ==========================================

func Test_CastedResult_IsInvalid_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsInvalidTestCases {
		// Act
		actual := args.Map{"result": tc.CR.IsInvalid()}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: IsNotNull
// ==========================================

func Test_CastedResult_IsNotNull_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsNotNullTestCases {
		// Act
		actual := args.Map{"result": tc.CR.IsNotNull()}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: IsNotPointer
// ==========================================

func Test_CastedResult_IsNotPointer_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsNotPointerTestCases {
		// Act
		actual := args.Map{"result": tc.CR.IsNotPointer()}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: IsNotMatchingAcceptedType
// ==========================================

func Test_CastedResult_IsNotMatchingAcceptedType_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsNotMatchingAcceptedTypeTestCases {
		// Act
		actual := args.Map{"result": tc.CR.IsNotMatchingAcceptedType()}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: IsSourceKind
// ==========================================

func Test_CastedResult_IsSourceKind_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsSourceKindTestCases {
		// Act
		actual := args.Map{"result": tc.CR.IsSourceKind(tc.CheckKind)}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: HasError
// ==========================================

func Test_CastedResult_HasError_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultHasErrorTestCases {
		// Act
		actual := args.Map{"result": tc.CR.HasError()}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: HasAnyIssues
// ==========================================

func Test_CastedResult_HasAnyIssues_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultHasAnyIssuesTestCases {
		// Act
		actual := args.Map{"result": tc.CR.HasAnyIssues()}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: SourceReflectType
// ==========================================

func Test_CastedResult_SourceReflectType_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultSourceReflectTypeTestCases {
		// Act
		actual := args.Map{
			"typeName":     tc.CR.SourceReflectType.Name(),
			"isStringKind": tc.CR.IsSourceKind(reflect.String),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Casted
// ==========================================

func Test_CastedResult_CastedValue_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultCastedValueTestCases {
		// Act
		actual := args.Map{
			"castedValue":  fmt.Sprintf("%v", tc.CR.Casted),
			"hasAnyIssues": tc.CR.HasAnyIssues(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: IsSourcePointer
// ==========================================

func Test_CastedResult_IsSourcePointer_Verification(t *testing.T) {
	for caseIndex, tc := range castedResultIsSourcePointerTestCases {
		// Act
		actual := args.Map{"result": tc.CR.IsSourcePointer}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
