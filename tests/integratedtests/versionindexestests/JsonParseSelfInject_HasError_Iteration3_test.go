package versionindexestests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/versionindexes"
)

// Test_Cov3_JsonParseSelfInject_HasError covers
// enums/versionindexes/Index.go L197-199: jsonResult.HasError() branch.
func Test_JsonParseSelfInject_HasError(t *testing.T) {
	// Arrange
	idx := versionindexes.Major
	errResult := corejson.NewPtr("invalid-data")
	errResult.Error = fmt.Errorf("simulated parse error")

	// Act
	err := idx.JsonParseSelfInject(errResult)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject returns error -- HasError branch", actual)
}
