package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/simplewrap"
)

func Test_WithDoubleQuote_Verification(t *testing.T) {
	for caseIndex, tc := range withDoubleQuoteTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := simplewrap.WithDoubleQuote(inputStr)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_WithSingleQuote_Verification(t *testing.T) {
	for caseIndex, tc := range withSingleQuoteTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := simplewrap.WithSingleQuote(inputStr)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_WithCurly_Verification(t *testing.T) {
	for caseIndex, tc := range withCurlyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")

		// Act
		result := simplewrap.WithCurly(inputVal)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_WithParenthesis_Verification(t *testing.T) {
	for caseIndex, tc := range withParenthesisTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")

		// Act
		result := simplewrap.WithParenthesis(inputVal)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_WithBrackets_String_Verification(t *testing.T) {
	for caseIndex, tc := range withBracketsStrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := simplewrap.WithBrackets(inputStr)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_TitleSquare_Verification(t *testing.T) {
	for caseIndex, tc := range titleSquareTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		title, _ := input.GetAsString("title")
		value, _ := input.GetAsString("value")

		// Act
		result := simplewrap.TitleSquare(title, value)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_TitleSquareMeta_Verification(t *testing.T) {
	for caseIndex, tc := range titleSquareMetaTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		title, _ := input.GetAsString("title")
		value, _ := input.GetAsString("value")
		meta, _ := input.GetAsString("meta")

		// Act
		result := simplewrap.TitleSquareMeta(title, value, meta)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_With_Verification(t *testing.T) {
	for caseIndex, tc := range withTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		start, _ := input.GetAsString("start")
		source, _ := input.GetAsString("source")
		end, _ := input.GetAsString("end")

		// Act
		result := simplewrap.With(start, source, end)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_WithPtr_Verification(t *testing.T) {
	for caseIndex, tc := range withPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		startNilRaw, hasStartNil := input.Get("startNil")

		source, _ := input.GetAsString("source")
		end, _ := input.GetAsString("end")

		var startPtr *string
		if hasStartNil && startNilRaw == true {
			startPtr = nil
		} else {
			startVal, _ := input.GetAsString("start")
			startPtr = &startVal
		}

		// Act
		result := simplewrap.WithPtr(startPtr, &source, &end)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, *result)
	}
}

func Test_ToJsonName_Verification(t *testing.T) {
	for caseIndex, tc := range toJsonNameTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")

		// Act
		result := simplewrap.ToJsonName(inputVal)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_ConditionalWrapWith_Verification(t *testing.T) {
	for caseIndex, tc := range conditionalWrapWithTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		startRaw, _ := input.Get("start")
		endRaw, _ := input.Get("end")

		// Act
		result := simplewrap.ConditionalWrapWith(startRaw.(byte), inputStr, endRaw.(byte))

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_WithDoubleQuoteAny_Verification(t *testing.T) {
	for caseIndex, tc := range withDoubleQuoteAnyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")

		// Act
		result := simplewrap.WithDoubleQuoteAny(inputVal)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}
