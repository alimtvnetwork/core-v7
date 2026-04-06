package enumimpltests

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

func Test_DynamicMapCreationDiff(t *testing.T) {
	for caseIndex, tc := range dynamicMapDiffCaseV1TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		left := enumimpl.DynamicMap(input["left"].(enumimpl.DynamicMap))
		right := enumimpl.DynamicMap(input["right"].(enumimpl.DynamicMap))

		// Act
		diffMap := left.DiffRaw(true, right)
		mapAnyDiffer := coredynamic.MapAnyItemDiff(left)
		anotherDiff := mapAnyDiffer.DiffRaw(true, right)

		// Assert - verify both diffs produce sorted key:value lines
		actLines := dynamicMapToSortedLines(diffMap)
		tc.ShouldBeEqual(t, caseIndex, actLines...)

		// Assert - verify both diff methods produce equal raw maps
		// Note: This is a cross-validation of two dynamic outputs (actLines as expected).
		// Cannot use CaseV1.ShouldBeEqual because the expected value is dynamic, not from ExpectedInput.
		anotherLines := dynamicMapToSortedLines(enumimpl.DynamicMap(anotherDiff))
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title+" (both diff equal)", anotherLines, actLines)

		lineCountActual := args.Map{"lineCount": len(actLines)}
		lineCountExpected := args.Map{"lineCount": len(anotherLines)}
		lineCountExpected.ShouldBeEqual(t, caseIndex, "both diff methods line count matches", lineCountActual)
	}
}

func Test_DynamicMapCreationDiffMessage(t *testing.T) {
	for caseIndex, tc := range dynamicMapDiffMessageCaseV1TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		left := enumimpl.DynamicMap(input["left"].(map[string]any))
		right := input["right"].(map[string]any)

		// Act
		diffJsonMessage := left.ShouldDiffMessage(
			true,
			tc.Title,
			right,
		)
		actLines := coretests.GetAssert.ToStrings(diffJsonMessage)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_DynamicMapCreationDiffMessageV2(t *testing.T) {
	for caseIndex, tc := range dynamicMapDiffMessageV2CaseV1TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		left := enumimpl.DynamicMap(input["left"].(map[string]any))
		right := input["right"].(map[string]any)
		checker := input["checker"].(enumimpl.DifferChecker)

		// Act
		diffJsonMessage := left.ShouldDiffLeftRightMessageUsingDifferChecker(
			checker,
			true,
			tc.Title,
			right,
		)
		actLines := strings.Split(
			diffJsonMessage,
			constants.NewLineUnix,
		)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func dynamicMapToSortedLines(dm enumimpl.DynamicMap) []string {
	if dm.IsEmpty() {
		return []string{}
	}

	keys := dm.AllKeysSorted()
	sort.Strings(keys)

	lines := make([]string, len(keys))
	for i, k := range keys {
		lines[i] = fmt.Sprintf("%s : %v", k, dm[k])
	}

	return lines
}
