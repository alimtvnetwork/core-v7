package enumimpltests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/coreimpl/enumimpl/enumtype"
)

func Test_Enum(t *testing.T) {
	for _, testCase := range enumTestCases {
		Convey(
			testCase.Header, t, func() {
				switch testCase.EnumType {
				case enumtype.Byte:
					assertBasicByte(testCase)
				case enumtype.Integer8:
					assertBasicInteger8(testCase)
				case enumtype.Integer16:
					assertBasicInteger16(testCase)
				case enumtype.Integer32:
					assertBasicInteger32(testCase)
				case enumtype.UnsignedInteger16:
					assertBasicUnsignedInteger16(testCase)
				case enumtype.String:
					assertBasicString(testCase)
				}
			},
		)
	}
}

func assertBasicString(testCase TestWrapper) {
	// Arrange
	enumImpl := testCase.EnumMap.BasicString("unknown type")
	actualMin := enumImpl.Min()
	actualMax := enumImpl.Max()

	// Assert
	So(actualMin, ShouldEqual, "")
	So(actualMax, ShouldEqual, "Something2") // it depends on string max than number max
}

func assertBasicUnsignedInteger16(testCase TestWrapper) {
	// Arrange
	enumImpl := testCase.EnumMap.BasicUInt16("unknown type")
	actualMin := enumImpl.Min()
	actualMax := enumImpl.Max()

	// Assert
	So(actualMin, ShouldEqual, testCase.ExpectedMinMax.Min)
	So(actualMax, ShouldEqual, testCase.ExpectedMinMax.Max)
}

func assertBasicInteger32(testCase TestWrapper) {
	// Arrange
	enumImplementation := testCase.EnumMap.BasicInt32("unknown type")
	actualMin := enumImplementation.Min()
	actualMax := enumImplementation.Max()

	// Assert
	So(actualMin, ShouldEqual, testCase.ExpectedMinMax.Min)
	So(actualMax, ShouldEqual, testCase.ExpectedMinMax.Max)
}

func assertBasicInteger16(testCase TestWrapper) {
	// Arrange
	enumImplementation := testCase.EnumMap.BasicInt16("unknown type")
	actualMin := enumImplementation.Min()
	actualMax := enumImplementation.Max()

	// Assert
	So(actualMin, ShouldEqual, testCase.ExpectedMinMax.Min)
	So(actualMax, ShouldEqual, testCase.ExpectedMinMax.Max)
}

func assertBasicInteger8(testCase TestWrapper) {
	// Arrange
	enumImplementation := testCase.EnumMap.BasicInt8("unknown type")
	actualMin := enumImplementation.Min()
	actualMax := enumImplementation.Max()

	// Assert
	So(actualMin, ShouldEqual, testCase.ExpectedMinMax.Min)
	So(actualMax, ShouldEqual, testCase.ExpectedMinMax.Max)
}

func assertBasicByte(testCase TestWrapper) {
	// Arrange
	enumImplementation := testCase.EnumMap.BasicByte("unknown type")
	actualMin := enumImplementation.Min()
	actualMax := enumImplementation.Max()

	// Assert
	So(actualMin, ShouldEqual, testCase.ExpectedMinMax.Min)
	So(actualMax, ShouldEqual, testCase.ExpectedMinMax.Max)
}
