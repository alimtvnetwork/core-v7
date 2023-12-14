package corerangestests

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/coredata/corerange"
)

func Test_Int_Ranges_ValidCases(t *testing.T) {
	for _, testCase := range validIntRangeTestCases {
		// Arrange
		arrangeInputs := testCase.Arrange()
		first := arrangeInputs[0]
		rest := arrangeInputs[1:]

		// Act
		actualRanges := first.CreateRanges(rest...)

		// Assert
		convey.Convey(
			testCase.Title, t, func() {
				convey.So(
					actualRanges,
					should.Equal,
					testCase.ExpectedInput,
				)
			},
		)

		convey.Convey(
			testCase.Title+" - type verify", t, func() {
				convey.So(
					testCase.TypeValidationError(),
					should.BeNil,
				)
			},
		)
	}
}

func Test_Int_ExceptRanges_Verify(t *testing.T) {
	// Arrange
	arrangeInput := corerange.MinMaxInt{
		Min: 1,
		Max: 15,
	}

	// Act
	actualRanges := arrangeInput.RangesExcept(
		3, 4, 5,
	)

	// Assert
	convey.Convey(
		"Ranges 1-15, RangesExcept(3, 4, 5), should not contain 3,4,5", t, func() {
			convey.So(
				actualRanges, should.Equal, []int{
					1, 2, 6,
					7, 8, 9,
					10, 11,
					12, 13,
					14, 15,
				},
			)
		},
	)
}

func Test_Int8_Ranges_ValidCases(t *testing.T) {
	for _, testCase := range validInt8RangeTestCases {
		// Arrange
		arrangeInputs := testCase.ArrangeInput.([]corerange.MinMaxInt8)
		first := arrangeInputs[0]
		rest := arrangeInputs[1:]

		// Act
		actualRanges := first.CreateRanges(rest...)

		// Assert
		convey.Convey(
			testCase.Title, t, func() {
				convey.So(
					actualRanges,
					should.Equal,
					testCase.ExpectedInput,
				)
			},
		)

		convey.Convey(
			testCase.Title+" - type verify", t, func() {
				convey.So(
					testCase.TypeValidationError(),
					should.BeNil,
				)
			},
		)
	}
}
