package corerangestests

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/errcore"
)

func Test_RangeInt_Valid_WithInRange_Verification(t *testing.T) {
	// Arrange
	validCases := []int{
		5, 13, 5, 10, 25,
	}
	toString := converters.AnyTo.ValueString(validCases)

	// Act, Assert
	title := toString + " -- all these are valid for (range) : " + someRange.String()
	convey.Convey(
		title, t, func() {
			for _, v := range validCases {
				for name, inWithFunc := range isWithInFuncsMap {
					validationErr := rangeValidationError(
						name,
						true,
						inWithFunc,
						v,
					)

					convey.So(
						validationErr,
						should.BeNil,
					)
				}
			}
		},
	)
}

func rangeValidationError(
	name string,
	isExpectValid bool,
	isWithInFunc isWithInDefinitionFunc,
	v int,
) error {
	isInRange := isWithInFunc(v)

	if !isInRange && isExpectValid {
		return errcore.WasExpectingErrorF(
			true,
			false,
			"%s - should be valid and within range : %d",
			name,
			v,
		)
	}

	if isInRange && !isExpectValid {
		return errcore.WasExpectingErrorF(
			true,
			false,
			"%s - should be invalid and within range : %d",
			name,
			v,
		)
	}

	return nil
}

func Test_RangeInt_Invalid_WithInRange_Verification(t *testing.T) {
	// Arrange
	invalidCases := []int{
		265, 311, 4, 26, 100,
	}
	toString := converters.AnyTo.ValueString(invalidCases)

	// Act, Assert
	title := toString + " -- all these are invalid for (range) : " + someRange.String()
	convey.Convey(
		title, t, func() {
			for _, v := range invalidCases {
				for name, inWithFunc := range isWithInFuncsMap {
					validationErr := rangeValidationError(
						name,
						false,
						inWithFunc,
						v,
					)

					convey.So(
						validationErr,
						should.BeNil,
					)
				}
			}
		},
	)
}
