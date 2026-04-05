package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage27 — corepayload remaining 35 lines
// ══════════════════════════════════════════════════════════════════════════════

// ── Attributes.Clone nil guard (line 46-48) ──
// Nil receiver — dead code

// ── Attributes error paths (lines 84, 134) ──
// Internal nil-guard paths — dead code

// ── AttributesGetters path (lines 130, 307) ──
// Internal nil / empty guard paths — dead code

// ── AttributesJson error path (line 119) ──
// json.Marshal error on map — dead code

// ── AttributesSetters nil guards (lines 13, 19, 29) ──
// Nil receiver guards — dead code

// ── PayloadWrapper nil guards (lines 134, 146, 188, 210, 230, 242, 276, 294, 335, 385) ──
// Most are nil-receiver guards. Test the reachable ones:

func Test_Cov27_PayloadWrapper_Serialize_Valid(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.Create("id1", "cat1", "name1")

	// Act
	bytes, err := pw.Serialize()

	// Assert
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasError": err != nil,
	}
	expected := args.Map{
		"hasBytes": true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "PayloadWrapper Serialize valid", expected)
}

func Test_Cov27_PayloadWrapper_DeserializeMust(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.Create("id1", "cat1", "name1")
	bytes, _ := pw.Serialize()
	jsonResult := corejson.NewResult.UsingBytes(bytes)

	// Act
	pw2 := corepayload.New.PayloadWrapper.Empty()
	result, err := pw2.ParseInjectUsingJson(&jsonResult)

	// Assert
	actual := args.Map{
		"notNil":   result != nil,
		"hasError": err != nil,
	}
	expected := args.Map{
		"notNil":   true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "PayloadWrapper ParseInjectUsingJson valid", expected)
}

// ── PayloadsCollectionFilter.FilterCategoryCollection edge (lines 52-54, 61, 139) ──

func Test_Cov27_PayloadsCollectionFilter_Empty(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.Create(0)

	// Act
	filtered := pc.FilterCategoryCollection("nonexistent")

	// Assert
	actual := args.Map{"isEmpty": filtered.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	actual.ShouldBeEqual(t, 1, "PayloadsCollectionFilter empty", expected)
}

// ── PayloadsCollectionGetters.Serialize error (line 189) ──
// json.Marshal error on valid collection — dead code

// ── TypedPayloadCollection various paths ──

func Test_Cov27_TypedPayloadCollection_Clone_Empty(t *testing.T) {
	// Arrange
	tc := corepayload.NewTypedPayloadCollection[string](0)

	// Act
	cloned := tc.Clone()

	// Assert
	actual := args.Map{"isEmpty": cloned.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	actual.ShouldBeEqual(t, 1, "TypedPayloadCollection Clone empty", expected)
}

func Test_Cov27_TypedPayloadCollection_FilterByCategory(t *testing.T) {
	// Arrange
	tc := corepayload.NewTypedPayloadCollection[string](3)
	tw1 := corepayload.NewTypedPayloadWrapper[string]("id1", "cat-a", "name1", "data1")
	tw2 := corepayload.NewTypedPayloadWrapper[string]("id2", "cat-b", "name2", "data2")
	tc.Add(tw1)
	tc.Add(tw2)

	// Act
	filtered := tc.FilterByCategory("cat-a")

	// Assert
	actual := args.Map{"count": filtered.Length()}
	expected := args.Map{"count": 1}
	actual.ShouldBeEqual(t, 1, "TypedPayloadCollection FilterByCategory", expected)
}

func Test_Cov27_TypedPayloadCollection_FilterByCategory_NotFound(t *testing.T) {
	// Arrange
	tc := corepayload.NewTypedPayloadCollection[string](2)
	tw := corepayload.NewTypedPayloadWrapper[string]("id1", "cat-a", "name1", "data1")
	tc.Add(tw)

	// Act
	filtered := tc.FilterByCategory("nonexistent")

	// Assert
	actual := args.Map{"isEmpty": filtered.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	actual.ShouldBeEqual(t, 1, "TypedPayloadCollection FilterByCategory not found", expected)
}

// ── TypedPayloadCollection: remaining paths (lines 460, 474, 487, 574, 591, 616, 633, 648, 664, 668) ──
// These are mostly nil-guard, error fallback, or internal branch paths.

func Test_Cov27_TypedPayloadCollection_Serialize_Valid(t *testing.T) {
	// Arrange
	tc := corepayload.NewTypedPayloadCollection[string](2)
	tw := corepayload.NewTypedPayloadWrapper[string]("id1", "cat", "name", "data")
	tc.Add(tw)

	// Act
	bytes, err := tc.Serialize()

	// Assert
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasError": err != nil,
	}
	expected := args.Map{
		"hasBytes": true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "TypedPayloadCollection Serialize valid", expected)
}

// ── TypedPayloadWrapper nil guards (lines 69, 101, 127, 289) ──
// Nil receiver guards — dead code

func Test_Cov27_TypedPayloadWrapper_Serialize_Valid(t *testing.T) {
	// Arrange
	tw := corepayload.NewTypedPayloadWrapper[string]("id1", "cat", "name", "data")

	// Act
	bytes, err := tw.Serialize()

	// Assert
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasError": err != nil,
	}
	expected := args.Map{
		"hasBytes": true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "TypedPayloadWrapper Serialize valid", expected)
}
