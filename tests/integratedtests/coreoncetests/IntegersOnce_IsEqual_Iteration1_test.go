package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// Test_Cov12_IntegersOnce_IsEqual_NilReceiver tests IsEqual on nil *IntegersOnce.
func Test_IntegersOnce_IsEqual_NilReceiver(t *testing.T) {
	// Arrange
	var nilOnce *coreonce.IntegersOnce

	// Act
	actual := nilOnce.IsEqual(nil...)

	// Assert
	actual := args.Map{"result": actual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil IntegersOnce.IsEqual(nil) should return true, got false", actual)
}

// Test_Cov12_IntegersOnce_IsEqual_NilReceiverEmpty tests IsEqual nil receiver with empty variadic.
func Test_IntegersOnce_IsEqual_NilReceiverEmpty(t *testing.T) {
	// Arrange
	var nilOnce *coreonce.IntegersOnce

	// Act
	actual := nilOnce.IsEqual()

	// Assert
	actual := args.Map{"result": actual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil IntegersOnce.IsEqual() should return true, got false", actual)
}

// Test_Cov12_MapStringStringOnce_IsEqual_NilReceiver tests IsEqual on nil *MapStringStringOnce.
func Test_MapStringStringOnce_IsEqual_NilReceiver(t *testing.T) {
	// Arrange
	var nilOnce *coreonce.MapStringStringOnce

	// Act
	actual := nilOnce.IsEqual(nil)

	// Assert
	actual := args.Map{"result": actual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil MapStringStringOnce.IsEqual(nil) should return true, got false", actual)
}

// Test_Cov12_MapStringStringOnce_JsonStringMust_Success tests JsonStringMust on valid data.
func Test_MapStringStringOnce_JsonStringMust_Success(t *testing.T) {
	// Arrange
	once := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"key": "value"}
	})

	// Act
	actual := once.JsonStringMust()

	// Assert
	expected := `{"key":"value"}`
	actual := args.Map{"result": actual != expected}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce.JsonStringMust: got, want", actual)
}

// Test_Cov12_StringsOnce_IsEqual_NilReceiver tests IsEqual on nil *StringsOnce.
func Test_StringsOnce_IsEqual_NilReceiver(t *testing.T) {
	// Arrange
	var nilOnce *coreonce.StringsOnce

	// Act
	actual := nilOnce.IsEqual()

	// Assert
	actual := args.Map{"result": actual}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil StringsOnce.IsEqual() should return true, got false", actual)
}

// Test_Cov12_StringsOnce_JsonStringMust_Success tests JsonStringMust on valid data.
func Test_StringsOnce_JsonStringMust_Success(t *testing.T) {
	// Arrange
	once := coreonce.NewStringsOnce(func() []string {
		return []string{"a", "b"}
	})

	// Act
	actual := once.JsonStringMust()

	// Assert
	expected := `["a","b"]`
	actual := args.Map{"result": actual != expected}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringsOnce.JsonStringMust: got, want", actual)
}
