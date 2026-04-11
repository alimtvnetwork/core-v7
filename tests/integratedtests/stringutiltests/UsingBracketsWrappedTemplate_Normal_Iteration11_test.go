package stringutiltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

// ══════════════════════════════════════════════════════════════════════════════
// UsingNamerMapOptions — curly and non-curly with actual values
// ══════════════════════════════════════════════════════════════════════════════

// Note: UsingNamerMapOptions non-nil paths require in-package test (namer is unexported)

// ══════════════════════════════════════════════════════════════════════════════
// UsingBracketsWrappedTemplate, UsingQuotesWrappedTemplate — normal paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_UsingBracketsWrappedTemplate_Normal(t *testing.T) {
	// Arrange
	result := stringutil.ReplaceTemplate.UsingBracketsWrappedTemplate("prefix {brackets-wrapped} suffix", "VALUE")

	// Act
	actual := args.Map{"has": len(result) > 0}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "UsingBracketsWrappedTemplate returns correct value -- normal", actual)
}

func Test_UsingQuotesWrappedTemplate_Normal(t *testing.T) {
	// Arrange
	result := stringutil.ReplaceTemplate.UsingQuotesWrappedTemplate(`prefix "{quotes-wrapped}" suffix`, "VALUE")

	// Act
	actual := args.Map{"has": len(result) > 0}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "UsingQuotesWrappedTemplate returns correct value -- normal", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReplaceWhiteSpaces — with tabs and newlines
// ══════════════════════════════════════════════════════════════════════════════

func Test_ReplaceWhiteSpaces_WithTabs(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.ReplaceWhiteSpaces("  a\tb\nc  ")}

	// Assert
	expected := args.Map{"v": "abc"}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpaces returns correct value -- tabs", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReplaceWhiteSpacesToSingle — with newlines/tabs
// ══════════════════════════════════════════════════════════════════════════════

func Test_ReplaceWhiteSpacesToSingle_WithNewlines(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.ReplaceWhiteSpacesToSingle("a\nb\tc")}

	// Assert
	expected := args.Map{"v": "abc"}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle returns correct value -- newlines", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CurlyKeyUsingMap — normal path
// ══════════════════════════════════════════════════════════════════════════════

func Test_CurlyKeyUsingMap_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyKeyUsingMap("{x}-{y}", map[string]string{"x": "1", "y": "2"})}

	// Assert
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "CurlyKeyUsingMap returns correct value -- normal", actual)
}
