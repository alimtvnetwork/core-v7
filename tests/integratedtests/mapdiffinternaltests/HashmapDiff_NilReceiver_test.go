package mapdiffinternaltests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/internal/mapdiffinternal"
)

// ==========================================================================
// HashmapDiff — nil receiver
// ==========================================================================

var extHashmapDiffNilReceiverTestCases = []coretestcases.CaseV1{
	{
		Title:         "Nil HashmapDiff ptr has length 0",
		ArrangeInput:  args.Map{"method": "Length"},
		ExpectedInput: args.Map{"result": 0},
	},
	{
		Title:         "Nil HashmapDiff IsRawEqual both nil returns true",
		ArrangeInput:  args.Map{"method": "IsRawEqual-both-nil"},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "Nil HashmapDiff IsRawEqual left nil right non-nil returns false",
		ArrangeInput:  args.Map{"method": "IsRawEqual-left-nil"},
		ExpectedInput: args.Map{"result": false},
	},
}

func Test_HashmapDiff_NilReceiver_Cov2(t *testing.T) {
	for caseIndex, testCase := range extHashmapDiffNilReceiverTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		method, _ := input.GetAsString("method")

		var h *mapdiffinternal.HashmapDiff
		var actual args.Map

		// Act
		switch method {
		case "Length":
			actual = args.Map{"result": h.Length()}
		case "IsRawEqual-both-nil":
			actual = args.Map{"result": h.IsRawEqual(nil)}
		case "IsRawEqual-left-nil":
			actual = args.Map{"result": h.IsRawEqual(map[string]string{"a": "1"})}
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// HashmapDiff — LogShouldDiffMessage
// ==========================================================================

var extHashmapDiffLogShouldDiffMessageTestCases = []coretestcases.CaseV1{
	{
		Title: "LogShouldDiffMessage no diff returns empty",
		ArrangeInput: args.Map{
			"left":  map[string]string{"a": "1"},
			"right": map[string]string{"a": "1"},
			"title": "test",
		},
		ExpectedInput: args.Map{"isEmpty": true},
	},
	{
		Title: "LogShouldDiffMessage with diff returns non-empty and prints",
		ArrangeInput: args.Map{
			"left":  map[string]string{"a": "1"},
			"right": map[string]string{"a": "2"},
			"title": "test",
		},
		ExpectedInput: args.Map{"isEmpty": false},
	},
}

func Test_HashmapDiff_LogShouldDiffMessage_Cov2(t *testing.T) {
	for caseIndex, testCase := range extHashmapDiffLogShouldDiffMessageTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left := input["left"].(map[string]string)
		right := input["right"].(map[string]string)
		title := input["title"].(string)
		h := mapdiffinternal.HashmapDiff(left)

		// Act
		result := h.LogShouldDiffMessage(title, right)

		// Assert
		actual := args.Map{"isEmpty": result == ""}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// HashmapDiff — DiffRaw right-to-left value diff branch
// ==========================================================================

var extHashmapDiffDiffRawRightDiffValueTestCases = []coretestcases.CaseV1{
	{
		Title: "DiffRaw right has different value for key not in left diff",
		ArrangeInput: args.Map{
			"left":  map[string]string{"a": "1", "b": "2"},
			"right": map[string]string{"a": "1", "b": "3", "c": "4"},
		},
		ExpectedInput: args.Map{
			"diffLength": 2,
			"hasKey-b":   true,
			"hasKey-c":   true,
		},
	},
}

func Test_HashmapDiff_DiffRaw_RightDiffValue_Cov2(t *testing.T) {
	for caseIndex, testCase := range extHashmapDiffDiffRawRightDiffValueTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left := input["left"].(map[string]string)
		right := input["right"].(map[string]string)
		h := mapdiffinternal.HashmapDiff(left)

		// Act
		diffMap := h.DiffRaw(right)

		// Assert
		actual := args.Map{"diffLength": len(diffMap)}
		for key := range diffMap {
			actual["hasKey-"+key] = true
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// HashmapDiff — HashmapDiffUsingRaw empty diff path
// ==========================================================================

var extHashmapDiffUsingRawTestCases = []coretestcases.CaseV1{
	{
		Title: "HashmapDiffUsingRaw identical maps returns empty",
		ArrangeInput: args.Map{
			"left":  map[string]string{"a": "1"},
			"right": map[string]string{"a": "1"},
		},
		ExpectedInput: args.Map{"length": 0},
	},
	{
		Title: "HashmapDiffUsingRaw different maps returns diff",
		ArrangeInput: args.Map{
			"left":  map[string]string{"a": "1"},
			"right": map[string]string{"a": "2"},
		},
		ExpectedInput: args.Map{"length": 1},
	},
}

func Test_HashmapDiff_HashmapDiffUsingRaw_Cov2(t *testing.T) {
	for caseIndex, testCase := range extHashmapDiffUsingRawTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left := input["left"].(map[string]string)
		right := input["right"].(map[string]string)
		h := mapdiffinternal.HashmapDiff(left)

		// Act
		result := h.HashmapDiffUsingRaw(right)

		// Assert
		actual := args.Map{"length": result.Length()}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// HashmapDiff — ToStringsSliceOfDiffMap
// ==========================================================================

var extHashmapDiffToStringsSliceTestCases = []coretestcases.CaseV1{
	{
		Title: "ToStringsSliceOfDiffMap formats entries",
		ArrangeInput: args.Map{
			"diffMap": map[string]string{"key": "val"},
		},
		ExpectedInput: args.Map{"length": 1},
	},
}

func Test_HashmapDiff_ToStringsSliceOfDiffMap_Cov2(t *testing.T) {
	for caseIndex, testCase := range extHashmapDiffToStringsSliceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		diffMap := input["diffMap"].(map[string]string)
		h := mapdiffinternal.HashmapDiff(map[string]string{})

		// Act
		slice := h.ToStringsSliceOfDiffMap(diffMap)

		// Assert
		actual := args.Map{"length": len(slice)}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// MapStringAnyDiff — nil receiver
// ==========================================================================

var extMapStringAnyDiffNilReceiverTestCases = []coretestcases.CaseV1{
	{
		Title:         "Nil MapStringAnyDiff ptr has length 0",
		ArrangeInput:  args.Map{"method": "Length"},
		ExpectedInput: args.Map{"result": 0},
	},
}

func Test_MapStringAnyDiff_NilReceiver_Cov2(t *testing.T) {
	for caseIndex, testCase := range extMapStringAnyDiffNilReceiverTestCases {
		// Arrange
		var m *mapdiffinternal.MapStringAnyDiff
		var actual args.Map

		// Act
		actual = args.Map{"result": m.Length()}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// MapStringAnyDiff — LogShouldDiffMessage
// ==========================================================================

var extMapStringAnyDiffLogTestCases = []coretestcases.CaseV1{
	{
		Title: "LogShouldDiffMessage no diff returns empty",
		ArrangeInput: args.Map{
			"left":  map[string]any{"a": 1},
			"right": map[string]any{"a": 1},
			"title": "test",
		},
		ExpectedInput: args.Map{"isEmpty": true},
	},
	{
		Title: "LogShouldDiffMessage with diff returns non-empty",
		ArrangeInput: args.Map{
			"left":  map[string]any{"a": 1},
			"right": map[string]any{"a": 2},
			"title": "test",
		},
		ExpectedInput: args.Map{"isEmpty": false},
	},
}

func Test_MapStringAnyDiff_LogShouldDiffMessage_Cov2(t *testing.T) {
	for caseIndex, testCase := range extMapStringAnyDiffLogTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left := input["left"].(map[string]any)
		right := input["right"].(map[string]any)
		title := input["title"].(string)
		m := mapdiffinternal.MapStringAnyDiff(left)

		// Act
		result := m.LogShouldDiffMessage(false, title, right)

		// Assert
		actual := args.Map{"isEmpty": result == ""}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// MapStringAnyDiff — Raw / HasAnyItem / LastIndex non-empty
// ==========================================================================

var extMapStringAnyDiffNonEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "Non-empty MapStringAnyDiff accessors",
		ArrangeInput: args.Map{
			"map": map[string]any{"a": 1, "b": 2, "c": 3},
		},
		ExpectedInput: args.Map{
			"hasAnyItem": true,
			"lastIndex":  2,
			"rawLength":  3,
		},
	},
}

func Test_MapStringAnyDiff_NonEmpty_Cov2(t *testing.T) {
	for caseIndex, testCase := range extMapStringAnyDiffNonEmptyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		raw := input["map"].(map[string]any)
		m := mapdiffinternal.MapStringAnyDiff(raw)

		// Act
		actual := args.Map{
			"hasAnyItem": m.HasAnyItem(),
			"lastIndex":  m.LastIndex(),
			"rawLength":  len(m.Raw()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// MapStringAnyDiff — DiffRaw right-to-left value diff
// ==========================================================================

var extMapStringAnyDiffDiffRawRightValueTestCases = []coretestcases.CaseV1{
	{
		Title: "DiffRaw right value differs for key not in left diff",
		ArrangeInput: args.Map{
			"left":  map[string]any{"a": 1, "b": 2},
			"right": map[string]any{"a": 1, "b": 3, "c": 4},
		},
		ExpectedInput: args.Map{
			"diffLength": 2,
			"hasKey-b":   true,
			"hasKey-c":   true,
		},
	},
}

func Test_MapStringAnyDiff_DiffRaw_RightDiffValue_Cov2(t *testing.T) {
	for caseIndex, testCase := range extMapStringAnyDiffDiffRawRightValueTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left := input["left"].(map[string]any)
		right := input["right"].(map[string]any)
		m := mapdiffinternal.MapStringAnyDiff(left)

		// Act
		diffMap := m.DiffRaw(false, right)

		// Assert
		actual := args.Map{"diffLength": len(diffMap)}
		for key := range diffMap {
			actual["hasKey-"+key] = true
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// MapStringAnyDiff — HashmapDiffUsingRaw
// ==========================================================================

var extMapStringAnyDiffUsingRawTestCases = []coretestcases.CaseV1{
	{
		Title: "HashmapDiffUsingRaw identical returns empty",
		ArrangeInput: args.Map{
			"left":  map[string]any{"a": 1},
			"right": map[string]any{"a": 1},
		},
		ExpectedInput: args.Map{"length": 0},
	},
	{
		Title: "HashmapDiffUsingRaw different returns diff",
		ArrangeInput: args.Map{
			"left":  map[string]any{"a": 1},
			"right": map[string]any{"a": 2},
		},
		ExpectedInput: args.Map{"length": 1},
	},
}

func Test_MapStringAnyDiff_HashmapDiffUsingRaw_Cov2(t *testing.T) {
	for caseIndex, testCase := range extMapStringAnyDiffUsingRawTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left := input["left"].(map[string]any)
		right := input["right"].(map[string]any)
		m := mapdiffinternal.MapStringAnyDiff(left)

		// Act
		result := m.HashmapDiffUsingRaw(false, right)

		// Assert
		actual := args.Map{"length": result.Length()}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// MapStringAnyDiff — DiffRaw right nil left non-nil
// ==========================================================================

var extMapStringAnyDiffDiffRawRightNilTestCases = []coretestcases.CaseV1{
	{
		Title: "DiffRaw right nil returns left",
		ArrangeInput: args.Map{
			"left": map[string]any{"a": 1},
		},
		ExpectedInput: args.Map{"diffLength": 1},
	},
}

func Test_MapStringAnyDiff_DiffRaw_RightNil_Cov2(t *testing.T) {
	for caseIndex, testCase := range extMapStringAnyDiffDiffRawRightNilTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left := input["left"].(map[string]any)
		m := mapdiffinternal.MapStringAnyDiff(left)

		// Act
		diffMap := m.DiffRaw(false, nil)

		// Assert
		actual := args.Map{"diffLength": len(diffMap)}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// MapStringAnyDiff — IsRawEqual nil branches
// ==========================================================================

var extMapStringAnyDiffIsRawEqualNilTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsRawEqual both nil returns true",
		ArrangeInput:  args.Map{"scenario": "both-nil"},
		ExpectedInput: args.Map{"result": "true"},
	},
	{
		Title:         "IsRawEqual left nil returns false",
		ArrangeInput:  args.Map{"scenario": "left-nil"},
		ExpectedInput: args.Map{"result": "false"},
	},
	{
		Title:         "IsRawEqual right nil returns false",
		ArrangeInput:  args.Map{"scenario": "right-nil"},
		ExpectedInput: args.Map{"result": "false"},
	},
	{
		Title:         "IsRawEqual different length returns false",
		ArrangeInput:  args.Map{"scenario": "diff-length"},
		ExpectedInput: args.Map{"result": "false"},
	},
	{
		Title:         "IsRawEqual missing key returns false",
		ArrangeInput:  args.Map{"scenario": "missing-key"},
		ExpectedInput: args.Map{"result": "false"},
	},
}

func Test_MapStringAnyDiff_IsRawEqual_NilBranches_Cov2(t *testing.T) {
	for caseIndex, testCase := range extMapStringAnyDiffIsRawEqualNilTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		scenario, _ := input.GetAsString("scenario")
		var result bool

		// Act
		switch scenario {
		case "both-nil":
			var m *mapdiffinternal.MapStringAnyDiff
			result = m.IsRawEqual(false, nil)
		case "left-nil":
			var m *mapdiffinternal.MapStringAnyDiff
			result = m.IsRawEqual(false, map[string]any{"a": 1})
		case "right-nil":
			m := mapdiffinternal.MapStringAnyDiff{"a": 1}
			result = m.IsRawEqual(false, nil)
		case "diff-length":
			m := mapdiffinternal.MapStringAnyDiff{"a": 1}
			result = m.IsRawEqual(false, map[string]any{"a": 1, "b": 2})
		case "missing-key":
			m := mapdiffinternal.MapStringAnyDiff{"a": 1}
			result = m.IsRawEqual(false, map[string]any{"b": 1})
		}

		// Assert
		actual := args.Map{"result": fmt.Sprintf("%v", result)}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// MapStringAnyDiff — DiffJsonMessage
// ==========================================================================

var extMapStringAnyDiffDiffJsonMessageTestCases = []coretestcases.CaseV1{
	{
		Title: "DiffJsonMessage regardless type same string rep returns empty",
		ArrangeInput: args.Map{
			"left":  map[string]any{"a": 1},
			"right": map[string]any{"a": "1"},
		},
		ExpectedInput: args.Map{"isEmpty": true},
	},
}

func Test_MapStringAnyDiff_DiffJsonMessage_RegardlessType_Cov2(t *testing.T) {
	for caseIndex, testCase := range extMapStringAnyDiffDiffJsonMessageTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left := input["left"].(map[string]any)
		right := input["right"].(map[string]any)
		m := mapdiffinternal.MapStringAnyDiff(left)

		// Act
		result := m.DiffJsonMessage(true, right)

		// Assert
		actual := args.Map{"isEmpty": result == ""}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
