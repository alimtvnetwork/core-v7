package versionindexestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/enums/versionindexes"
)

func Test_QW_Index_JsonParseSelfInject_NilResult(t *testing.T) {
	// Arrange
	idx := versionindexes.Major

	// Act
	err := idx.JsonParseSelfInject(nil)

	// Assert
	actual := args.Map{
		"hasError": err != nil,
	}
	expected := args.Map{
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject returns error -- nil result", actual)
}

func Test_QW_Index_JsonParseSelfInject_ErrorResult(t *testing.T) {
	// Arrange
	idx := versionindexes.Major
	bad := corejson.NewResult.UsingString(`invalid`)

	// Act
	err := idx.JsonParseSelfInject(bad)

	// Assert
	actual := args.Map{
		"hasError": err != nil,
	}
	expected := args.Map{
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject returns error -- invalid json", actual)
}
