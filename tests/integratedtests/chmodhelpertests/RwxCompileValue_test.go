package chmodhelpertests

import (
	"fmt"
	"log"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_RwxCompileValue(t *testing.T) {
	for _, testCase := range chmodhelpertestwrappers.RwxCompileValueTestCases {
		// Arrange
		existingRwxWrapper, _ :=
			chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
				&testCase.Existing)
		expectedVariableWrapper, _ :=
			chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
				&testCase.Expected)
		existing := testCase.Existing.ToString(false)
		input := testCase.Input.ToString(false)
		expected := testCase.Expected.ToString(false)
		expectedFullRwx := expectedVariableWrapper.
			ToCompileFixedPtr().
			ToFullRwxValueString()

		header := fmt.Sprintf(
			"Existing [%s] Applied by [%s] should result [%s]",
			existing,
			input,
			expected)

		// Act
		actualVarWrapper, _ :=
			chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
				&testCase.Input)
		actualRwxWrapper := actualVarWrapper.
			ToCompileWrapper(existingRwxWrapper.ToCompileFixedPtr())
		actualFullRwx := actualRwxWrapper.ToFullRwxValueString()

		// Assert
		convey.Convey(header, t, func() {
			if actualFullRwx != expectedFullRwx {
				log.Println(header)
			}

			convey.So(actualFullRwx, convey.ShouldEqual, expectedFullRwx)
		})
	}
}
