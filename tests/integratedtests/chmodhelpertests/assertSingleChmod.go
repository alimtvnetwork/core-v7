package chmodhelpertests

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/errcore"
)

// assertSingleChmod , expectedChmodRwxFullString 10 chars "-rwxrwxrwx"
func assertSingleChmod(
	t *testing.T,
	testHeader string,
	createPath *chmodhelper.DirFilesWithRwxPermission,
	expectedChmodRwxFullString string,
) {
	fileChmodMap := createPath.GetFilesChmodMap()
	for filePath, chmodValueString := range fileChmodMap.Items() {
		convey.Convey(testHeader, t, func() {
			isEqual := chmodValueString == expectedChmodRwxFullString

			if !isEqual {
				fmt.Println(
					errcore.Expecting(
						filePath,
						expectedChmodRwxFullString,
						chmodValueString))

			}

			convey.So(isEqual, convey.ShouldBeTrue)
		})
	}
}
