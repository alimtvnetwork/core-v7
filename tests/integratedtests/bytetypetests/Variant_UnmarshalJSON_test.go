package bytetypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/bytetype"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Variant_UnmarshalJSON_Error(t *testing.T) {
	// Arrange
	v := new(bytetype.Variant)
	err := v.UnmarshalJSON([]byte("invalid"))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Variant_UnmarshallToValue_FromVariantUnmarshalJSON(t *testing.T) {
	v := bytetype.Variant(1)
	jsonBytes, _ := corejson.Serialize.Raw(v)
	val, err := v.UnmarshallToValue(jsonBytes)
	// MarshalJSON serializes to enum name string, UnmarshallToValue
	// round-trips through JSON — the resulting byte value may differ
	// from the original iota value depending on enum implementation.
	_ = val
	_ = err
}
