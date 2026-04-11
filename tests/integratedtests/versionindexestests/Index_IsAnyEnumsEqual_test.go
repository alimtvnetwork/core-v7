package versionindexestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/versionindexes"
)

func Test_Index_IsAnyEnumsEqual_NoMatch(t *testing.T) {
	// Arrange
	idx := versionindexes.Major
	minor := versionindexes.Minor

	// Act
	actual := args.Map{
		"result": idx.IsAnyEnumsEqual(&minor),
	}

	// Assert
	expected := args.Map{
		"result": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyEnumsEqual returns false -- no match", actual)
}
