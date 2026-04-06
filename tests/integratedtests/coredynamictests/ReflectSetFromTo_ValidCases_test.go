package coredynamictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/tests/testwrappers/coredynamictestwrappers"
)

// ==========================================================================
// Helper: common base assertion for ReflectSetFromTo
// ==========================================================================

func assertReflectSetFromToBase(
	t *testing.T,
	tc coredynamictestwrappers.FromToTestWrapper,
) (actLines, expected []string) {
	t.Helper()

	err := coredynamic.ReflectSetFromTo(tc.From, tc.To)

	typeStatus := coredynamic.TypeSameStatus(tc.To, tc.ExpectedValue)
	tc.SetActual(tc.To)

	actLines = []string{
		fmt.Sprintf("%v", err == nil),
		fmt.Sprintf("%v", typeStatus.IsSame),
	}
	expected = []string{"true", "true"}

	return actLines, expected
}

// ==========================================================================
// Test: (null, null) — do nothing
// ==========================================================================

func Test_ReflectSetFromTo_NullNull(t *testing.T) {
	tc := coredynamictestwrappers.ReflectSetFromToValidNullNull

	actLines, expected := assertReflectSetFromToBase(t, tc)

	// Assert
	tc.ShouldBeEqual(t, 0, actLines, expected)
}

// ==========================================================================
// Test: (*DraftType, *DraftType) — pointer to pointer
// ==========================================================================

func Test_ReflectSetFromTo_PtrToPtr(t *testing.T) {
	tc := coredynamictestwrappers.ReflectSetFromToValidPtrToPtr

	actLines, expected := assertReflectSetFromToBase(t, tc)

	// Additional DraftType field verification
	convertedFrom := tc.From.(*coretests.DraftType)
	toField := tc.ToFieldToDraftType()

	// Assert
	expectedField := tc.ExpectedFieldToDraftType()
	toFieldEqualErr := toField.VerifyNotEqualExcludingInnerFieldsErr(expectedField)
	fromFieldEqualErr := convertedFrom.VerifyNotEqualExcludingInnerFieldsErr(expectedField)

	actLines = append(actLines,
		fmt.Sprintf("%v", toFieldEqualErr == nil),
		fmt.Sprintf("%v", fromFieldEqualErr == nil),
	)
	expected = append(expected, "true", "true")

	tc.ShouldBeEqual(t, 0, actLines, expected)
}

// ==========================================================================
// Test: (DraftType, *DraftType) — value to pointer
// ==========================================================================

func Test_ReflectSetFromTo_ValueToPtr(t *testing.T) {
	tc := coredynamictestwrappers.ReflectSetFromToValidValueToPtr

	actLines, expected := assertReflectSetFromToBase(t, tc)

	// Additional DraftType field verification
	convertedFrom := tc.From.(coretests.DraftType)
	toField := tc.ToFieldToDraftType()

	// Assert
	expectedField := tc.ExpectedFieldToDraftType()
	toFieldEqualErr := toField.VerifyNotEqualExcludingInnerFieldsErr(expectedField)
	fromFieldEqualErr := convertedFrom.VerifyNotEqualExcludingInnerFieldsErr(expectedField)

	actLines = append(actLines,
		fmt.Sprintf("%v", toFieldEqualErr == nil),
		fmt.Sprintf("%v", fromFieldEqualErr == nil),
	)
	expected = append(expected, "true", "true")

	tc.ShouldBeEqual(t, 0, actLines, expected)
}

// ==========================================================================
// Test: ([]byte, *DraftType) — bytes to draft type
// ==========================================================================

func Test_ReflectSetFromTo_BytesToDraft(t *testing.T) {
	tc := coredynamictestwrappers.ReflectSetFromToValidBytesToDraft

	actLines, expected := assertReflectSetFromToBase(t, tc)

	// Additional field verification
	toField := tc.ToFieldToDraftType()
	toFieldEqualErr := toField.VerifyNotEqualExcludingInnerFieldsErr(
		tc.ExpectedFieldToDraftType(),
	)

	actLines = append(actLines,
		fmt.Sprintf("%v", toFieldEqualErr == nil),
	)
	expected = append(expected, "true")

	// Assert
	tc.ShouldBeEqual(t, 0, actLines, expected)
}

// ==========================================================================
// Test: (*[]byte, *[]byte) — draft to bytes
// ==========================================================================

func Test_ReflectSetFromTo_DraftToBytes(t *testing.T) {
	tc := coredynamictestwrappers.ReflectSetFromToValidDraftToBytes

	actLines, expected := assertReflectSetFromToBase(t, tc)

	// Assert
	tc.ShouldBeEqual(t, 0, actLines, expected)
}
