package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── AnyCollection.JsonStringMust valid (line 495) ──

func Test_AnyCollection_JsonStringMust_Valid_I29(t *testing.T) {
	// Arrange
	coll := coredynamic.NewAnyCollection(3)
	coll.Add("hello")
	coll.Add(42)

	// Act
	jsonStr := coll.JsonStringMust()

	// Assert
	actual := args.Map{"hasContent": len(jsonStr) > 0}
	expected := args.Map{"hasContent": true}
	actual.ShouldBeEqual(t, 1, "AnyCollection JsonStringMust valid", expected)
}
