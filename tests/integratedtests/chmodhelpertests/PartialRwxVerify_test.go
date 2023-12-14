package chmodhelpertests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_PartialRwxVerify(t *testing.T) {
	for i, testCase := range chmodhelpertestwrappers.PartialRwxVerifyTestCases {
		Convey(
			testCase.Header, t, func() {
				// Arrange
				rwx, err := chmodhelper.NewRwxVariableWrapper(testCase.PartialRwxInput1)
				errcore.SimpleHandleErr(err, "rwxVar create failed.")

				// Act
				actual := rwx.IsEqualPartialRwxPartial(testCase.FullRwxVerifyInput2)

				// Assert
				if actual != testCase.IsMatchesExpectation {
					fmt.Println("Input 1 :", testCase.PartialRwxInput1)
					fmt.Println("Input 2 :", testCase.FullRwxVerifyInput2)
					fmt.Println(testCase.Header, " --- Failed. Index :", i)
				}

				So(actual, ShouldEqual, testCase.IsMatchesExpectation)
			},
		)
	}
}
