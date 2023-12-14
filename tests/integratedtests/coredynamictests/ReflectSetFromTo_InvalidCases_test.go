package coredynamictests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/internal/trydo"
	"gitlab.com/auk-go/core/tests/testwrappers/coredynamictestwrappers"
)

// Test_ReflectSetFromTo_ValidCases
//
// Valid Inputs:
//   - From, To: (null, null)                          -- do nothing
//   - From, To: (sameTypePointer, sameTypePointer)    -- try reflection
//   - From, To: (sameTypeNonPointer, sameTypePointer) -- try reflection
//   - From, To: ([]byte or *[]byte, otherType)        -- try unmarshal, reflect
//   - From, To: (otherType, *[]byte)                  -- try marshal, reflect
func Test_ReflectSetFromTo_Invalid_Cases_With_Error_Verifications(t *testing.T) {
	for _, testCase := range coredynamictestwrappers.ReflectSetFromToInvalidTestCases {
		// Act
		wrappedResult := trydo.ErrorFuncWrapPanic(
			func() error {
				return coredynamic.ReflectSetFromTo(
					testCase.From,
					testCase.To,
				)
			},
		)

		err := wrappedResult.Error

		testCase.SetActual(wrappedResult)

		// Assert
		Convey(
			testCase.CaseTitle(), t, func() {
				if testCase.IsExpectingError {
					So(err, ShouldNotBeNil)
				} else {
					So(err, ShouldBeNil)
				}
			},
		)

		var parameter = &corevalidator.Parameter{
			CaseIndex:                  0,
			Header:                     testCase.Header,
			IsSkipCompareOnActualEmpty: false,
			IsAttachUserInputs:         true,
			IsCaseSensitive:            true,
		}

		validator := testCase.Validator

		Convey(
			testCase.CaseTitle()+"-exception-validation", t, func() {
				finalErr := getFinalVerificationError(
					testCase,
					validator,
					parameter,
					wrappedResult,
				)

				So(finalErr, ShouldBeNil)
			},
		)
	}
}

func getFinalVerificationError(
	testCase coredynamictestwrappers.FromToTestWrapper,
	validator corevalidator.TextValidator,
	parameter *corevalidator.Parameter,
	wrappedResult trydo.WrappedErr,
) error {
	if testCase.HasPanic {
		return validator.VerifyDetailError(
			parameter,
			wrappedResult.ExceptionString(),
		)
	}

	if testCase.IsExpectingError {
		return validator.VerifyDetailError(
			parameter,
			wrappedResult.ErrorString(),
		)
	}

	return validator.VerifyDetailError(
		parameter,
		converters.AnyTo.ValueString(testCase.ExpectedValue),
	)
}
