package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── DynamicCollection.JsonStringMust valid (line 426) ──

func Test_DynamicCollection_JsonStringMust_Valid_I29(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(3)
	dc.AddAny("hello", true)

	// Act
	jsonStr := dc.JsonStringMust()

	// Assert
	actual := args.Map{"hasContent": len(jsonStr) > 0}
	expected := args.Map{"hasContent": true}
	actual.ShouldBeEqual(t, 1, "DynamicCollection JsonStringMust valid", expected)
}
