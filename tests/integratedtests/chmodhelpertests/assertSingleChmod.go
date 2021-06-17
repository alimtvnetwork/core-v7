package chmodhelpertests

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

// assertSingleChmod , expectedChmodRwxFullString 10 chars "-rwxrwxrwx"
func assertSingleChmod(
	t *testing.T,
	testHeader string,
	createPath *chmodhelpertestwrappers.CreatePathsInstruction,
	expectedChmodRwxFullString string,
) {
	fileChmodMap := createPath.GetFilesChmodMap()
	for filePath, chmodValueString := range *fileChmodMap.Items() {
		convey.Convey(testHeader, t, func() {
			isEqual := chmodValueString == expectedChmodRwxFullString

			if !isEqual {
				fmt.Println(
					msgtype.Expecting(
						filePath,
						expectedChmodRwxFullString,
						chmodValueString))

			}

			convey.So(isEqual, convey.ShouldBeTrue)
		})
	}
}
