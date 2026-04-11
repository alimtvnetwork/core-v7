package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── CollectionLock.LengthLock (line 15) ──

func Test_CollectionLock_LengthLock_I29(t *testing.T) {
	// Arrange
	coll := coredynamic.NewCollection[int](5)
	coll.Add(1)
	coll.Add(2)

	// Act
	length := coll.LengthLock()

	// Assert
	actual := args.Map{"length": length}
	expected := args.Map{"length": 2}
	actual.ShouldBeEqual(t, 1, "CollectionLock LengthLock", expected)
}

// ── CollectionLock.RemoveAtLock invalid index (line 125) ──

func Test_CollectionLock_RemoveAtLock_InvalidIndex_I29(t *testing.T) {
	// Arrange
	coll := coredynamic.NewCollection[string](2)
	coll.Add("a")

	// Act
	removed := coll.RemoveAtLock(99)

	// Assert
	actual := args.Map{"removed": removed}
	expected := args.Map{"removed": false}
	actual.ShouldBeEqual(t, 1, "CollectionLock RemoveAtLock invalid index", expected)
}
