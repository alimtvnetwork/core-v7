package errcoretests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ==========================================================================
// Diagnostic Output Regression Tests
//
// These tests verify the formatting of diagnostic output functions to
// prevent regressions in visual structure, alignment, and content.
// ==========================================================================

// ── LineDiff ──

func Test_LineDiff_MatchingLines(t *testing.T) {
	// Act
	actual := []string{"a", "b", "c"}

	// Assert
	expected := []string{"a", "b", "c"}
	results := errcore.LineDiff(actual, expected)
	actual2 := args.Map{
		"len":       len(results),
		"status0":   results[0].Status,
		"status1":   results[1].Status,
		"status2":   results[2].Status,
		"actual0":   results[0].Actual,
		"expected0": results[0].Expected,
	}
	expected2 := args.Map{
		"len": 3, "status0": "  ", "status1": "  ", "status2": "  ",
		"actual0": "a", "expected0": "a",
	}
	expected2.ShouldBeEqual(t, 0, "LineDiff returns all-match -- matching lines", actual2)
}

func Test_LineDiff_MismatchLines(t *testing.T) {
	// Act
	actual := []string{"a", "WRONG", "c"}

	// Assert
	expected := []string{"a", "b", "c"}
	results := errcore.LineDiff(actual, expected)
	actual2 := args.Map{
		"len":          len(results),
		"status1":      results[1].Status,
		"mismatchAct":  results[1].Actual,
		"mismatchExp":  results[1].Expected,
		"lineNum":      results[1].LineNumber,
	}
	expected2 := args.Map{
		"len": 3, "status1": "!!",
		"mismatchAct": "WRONG", "mismatchExp": "b",
		"lineNum": 1,
	}
	expected2.ShouldBeEqual(t, 0, "LineDiff returns mismatch -- different line", actual2)
}

func Test_LineDiff_ExtraActual_FromLineDiffMatchingLine(t *testing.T) {
	// Act
	actual := []string{"a", "b", "c"}

	// Assert
	expected := []string{"a", "b"}
	results := errcore.LineDiff(actual, expected)
	actual2 := args.Map{
		"len":     len(results),
		"status2": results[2].Status,
		"extra":   results[2].Actual,
	}
	expected2 := args.Map{
		"len": 3, "status2": "+", "extra": "c",
	}
	expected2.ShouldBeEqual(t, 0, "LineDiff returns extra-actual -- longer actual", actual2)
}

func Test_LineDiff_MissingExpected_FromLineDiffMatchingLine(t *testing.T) {
	// Act
	actual := []string{"a"}

	// Assert
	expected := []string{"a", "b"}
	results := errcore.LineDiff(actual, expected)
	actual2 := args.Map{
		"len":     len(results),
		"status1": results[1].Status,
		"missing": results[1].Expected,
	}
	expected2 := args.Map{
		"len": 2, "status1": "-", "missing": "b",
	}
	expected2.ShouldBeEqual(t, 0, "LineDiff returns missing-expected -- shorter actual", actual2)
}

func Test_LineDiff_BothEmpty(t *testing.T) {
	// Arrange
	results := errcore.LineDiff([]string{}, []string{})

	// Act
	actual := args.Map{"len": len(results)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LineDiff returns empty -- both empty", actual)
}

// ── HasAnyMismatchOnLines ──

func Test_HasAnyMismatchOnLines_Match_FromLineDiffMatchingLine(t *testing.T) {
	// Arrange
	result := errcore.HasAnyMismatchOnLines(
		[]string{"a", "b"},
		[]string{"a", "b"},
	)

	// Act
	actual := args.Map{"hasMismatch": result}

	// Assert
	expected := args.Map{"hasMismatch": false}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines returns false -- matching", actual)
}

func Test_HasAnyMismatchOnLines_DifferentLength(t *testing.T) {
	// Arrange
	result := errcore.HasAnyMismatchOnLines(
		[]string{"a"},
		[]string{"a", "b"},
	)

	// Act
	actual := args.Map{"hasMismatch": result}

	// Assert
	expected := args.Map{"hasMismatch": true}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines returns true -- different length", actual)
}

func Test_HasAnyMismatchOnLines_ContentDiffers(t *testing.T) {
	// Arrange
	result := errcore.HasAnyMismatchOnLines(
		[]string{"a", "x"},
		[]string{"a", "b"},
	)

	// Act
	actual := args.Map{"hasMismatch": result}

	// Assert
	expected := args.Map{"hasMismatch": true}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines returns true -- content differs", actual)
}

// ── LineDiffToString ──

func Test_LineDiffToString_NoMismatch(t *testing.T) {
	// Arrange
	result := errcore.LineDiffToString(0, "Test", []string{"a"}, []string{"a"})

	// Act
	actual := args.Map{"hasOutput": result != ""}

	// Assert
	expected := args.Map{"hasOutput": true}
	expected.ShouldBeEqual(t, 0, "LineDiffToString returns output -- even when matching", actual)
}

func Test_LineDiffToString_HasMismatch(t *testing.T) {
	// Arrange
	result := errcore.LineDiffToString(1, "TestCase", []string{"a", "x"}, []string{"a", "b"})

	// Act
	actual := args.Map{
		"containsDiff":    strings.Contains(result, "Line-by-Line Diff"),
		"containsCase":    strings.Contains(result, "Case 1"),
		"containsTitle":   strings.Contains(result, "TestCase"),
		"containsMismatch": strings.Contains(result, "MISMATCH"),
		"containsActual":  strings.Contains(result, "actual"),
		"containsExpected": strings.Contains(result, "expected"),
		"containsTotal":   strings.Contains(result, "Total"),
	}

	// Assert
	expected := args.Map{
		"containsDiff": true, "containsCase": true,
		"containsTitle": true, "containsMismatch": true,
		"containsActual": true, "containsExpected": true,
		"containsTotal": true,
	}
	expected.ShouldBeEqual(t, 0, "LineDiffToString returns formatted diff -- mismatch present", actual)
}

func Test_LineDiffToString_ExtraAndMissing(t *testing.T) {
	// Arrange
	result := errcore.LineDiffToString(0, "EdgeCase",
		[]string{"a", "b", "extra"},
		[]string{"a"},
	)

	// Act
	actual := args.Map{
		"containsExtra":  strings.Contains(result, "EXTRA ACTUAL"),
		"containsMismatches": strings.Contains(result, "2 mismatches"),
	}

	// Assert
	expected := args.Map{
		"containsExtra": true, "containsMismatches": true,
	}
	expected.ShouldBeEqual(t, 0, "LineDiffToString returns extra-actual markers -- longer actual", actual)
}

func Test_LineDiffToString_MissingExpected(t *testing.T) {
	// Arrange
	result := errcore.LineDiffToString(0, "Missing",
		[]string{"a"},
		[]string{"a", "b", "c"},
	)

	// Act
	actual := args.Map{
		"containsMissing": strings.Contains(result, "MISSING EXPECTED"),
	}

	// Assert
	expected := args.Map{"containsMissing": true}
	expected.ShouldBeEqual(t, 0, "LineDiffToString returns missing-expected markers -- shorter actual", actual)
}

// ── MapMismatchError ──

func Test_MapMismatchError_Format(t *testing.T) {
	// Arrange
	result := errcore.MapMismatchError(
		"Test_Example",
		0,
		"example title",
		[]string{`"key1": 1,`, `"key2": 2,`},
		[]string{`"key1": 1,`},
	)

	// Act
	actual := args.Map{
		"containsTestMethod": strings.Contains(result, "Test Method"),
		"containsCase":       strings.Contains(result, "Case"),
		"containsTitle":      strings.Contains(result, "Title"),
		"containsSeparator":  strings.Contains(result, "============================>"),
		"containsActual":     strings.Contains(result, "Actual Received"),
		"containsExpected":   strings.Contains(result, "Expected Input"),
		"actualEntries":      strings.Contains(result, "2 entries"),
		"expectedEntries":    strings.Contains(result, "1 entries"),
		"containsKey1":       strings.Contains(result, `"key1": 1,`),
		"containsKey2":       strings.Contains(result, `"key2": 2,`),
		"tabIndented":        strings.Contains(result, "\t\"key1\""),
	}

	// Assert
	expected := args.Map{
		"containsTestMethod": true, "containsCase": true,
		"containsTitle": true, "containsSeparator": true,
		"containsActual": true, "containsExpected": true,
		"actualEntries": true, "expectedEntries": true,
		"containsKey1": true, "containsKey2": true,
		"tabIndented": true,
	}
	expected.ShouldBeEqual(t, 0, "MapMismatchError returns formatted block -- standard format", actual)
}

func Test_MapMismatchError_HeaderAlignment(t *testing.T) {
	// Arrange
	result := errcore.MapMismatchError(
		"Test_Alignment",
		5,
		"alignment check",
		[]string{`"a": 1,`},
		[]string{`"a": 1,`},
	)
	lines := strings.Split(result, "\n")

	// Verify 4-space indent for header lines
	hasTestMethod := false
	hasCase := false
	hasTitle := false
	for _, line := range lines {
		if strings.Contains(line, "Test Method") && strings.HasPrefix(line, "    ") {
			hasTestMethod = true
		}
		if strings.Contains(line, "Case") && strings.Contains(line, ": 5") {
			hasCase = true
		}
		if strings.Contains(line, "Title") && strings.Contains(line, "alignment check") {
			hasTitle = true
		}
	}

	// Act
	actual := args.Map{
		"hasTestMethod": hasTestMethod,
		"hasCase":       hasCase,
		"hasTitle":      hasTitle,
	}

	// Assert
	expected := args.Map{
		"hasTestMethod": true, "hasCase": true, "hasTitle": true,
	}
	expected.ShouldBeEqual(t, 0, "MapMismatchError returns aligned headers -- 4-space indent", actual)
}

// ── SliceDiffSummary ──

func Test_SliceDiffSummary_NoMismatch(t *testing.T) {
	// Arrange
	result := errcore.SliceDiffSummary([]string{"a"}, []string{"a"})

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "all lines match"}
	expected.ShouldBeEqual(t, 0, "SliceDiffSummary returns all-match -- no mismatch", actual)
}

func Test_SliceDiffSummary_HasMismatch(t *testing.T) {
	// Arrange
	result := errcore.SliceDiffSummary([]string{"a", "x"}, []string{"a", "b"})

	// Act
	actual := args.Map{
		"containsCount": strings.Contains(result, "1 mismatches"),
		"containsLine":  strings.Contains(result, "line 1"),
	}

	// Assert
	expected := args.Map{
		"containsCount": true,
		"containsLine": true,
	}
	expected.ShouldBeEqual(t, 0, "SliceDiffSummary returns mismatch summary -- one mismatch", actual)
}

// ── ErrorToLinesLineDiff ──

func Test_ErrorToLinesLineDiff_NilError(t *testing.T) {
	// Arrange
	result := errcore.ErrorToLinesLineDiff(0, "NilErr", nil, []string{"expected"})

	// Act
	actual := args.Map{
		"containsMissing": strings.Contains(result, "MISSING EXPECTED"),
	}

	// Assert
	expected := args.Map{"containsMissing": true}
	expected.ShouldBeEqual(t, 0, "ErrorToLinesLineDiff returns missing-expected -- nil error", actual)
}

func Test_ErrorToLinesLineDiff_MatchingError(t *testing.T) {
	// Arrange
	result := errcore.ErrorToLinesLineDiff(0, "Match", nil, []string{})

	// Act
	actual := args.Map{"isEmpty": result == ""}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorToLinesLineDiff returns empty -- nil error empty expected", actual)
}
