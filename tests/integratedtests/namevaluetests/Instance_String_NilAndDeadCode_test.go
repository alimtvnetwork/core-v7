package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/namevalue"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage9 — Instance.String nil receiver + Collection.JsonString error branch
//
// Target 1: namevalue/Instance.go:20-22
//   it.IsNull() → return EmptyString (nil receiver)
//
// Target 2: namevalue/Collection.go:385-387
//   err != nil || len(jsonBytes) == 0 → return EmptyString
//   json.Marshal of a valid Collection won't fail; this is defensive dead code.
//   Documented as accepted gap.
// ══════════════════════════════════════════════════════════════════════════════

func Test_Instance_String_NilReceiver(t *testing.T) {
	// Arrange
	var inst *namevalue.Instance[string, string]

	// Act
	result := inst.String()

	// Assert
	convey.Convey("Instance.String returns empty for nil receiver", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

// Coverage note: Collection.JsonString line 385-387 (err != nil || len==0)
// is defensive dead code — json.Marshal on Collection[K,V] with basic types
// cannot fail. Documented as accepted dead-code gap.
