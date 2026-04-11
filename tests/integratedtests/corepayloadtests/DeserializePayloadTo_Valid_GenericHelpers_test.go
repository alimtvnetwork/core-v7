package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── DeserializePayloadTo ──

func Test_DeserializePayloadTo_Valid(t *testing.T) {
	// Arrange
	type testData struct {
		Name string `json:"Name"`
	}
	jsonBytes, _ := corejson.Serialize.Raw(testData{Name: "Alice"})
	pw := &corepayload.PayloadWrapper{Payloads: jsonBytes}

	// Act
	result, err := corepayload.DeserializePayloadTo[testData](pw)

	// Assert
	actual := args.Map{
		"name":    result.Name,
		"noError": err == nil,
	}
	expected := args.Map{
		"name":    "Alice",
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadTo returns typed -- valid", actual)
}

func Test_DeserializePayloadTo_NilWrapper(t *testing.T) {
	// Arrange
	type testData struct{ Name string }

	// Act
	_, err := corepayload.DeserializePayloadTo[testData](nil)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadTo returns error -- nil wrapper", actual)
}

func Test_DeserializePayloadToMust_Valid(t *testing.T) {
	// Arrange
	type testData struct {
		Name string `json:"Name"`
	}
	jsonBytes, _ := corejson.Serialize.Raw(testData{Name: "Bob"})
	pw := &corepayload.PayloadWrapper{Payloads: jsonBytes}
	panicked := false

	// Act
	var result testData
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		result = corepayload.DeserializePayloadToMust[testData](pw)
	}()

	// Assert
	actual := args.Map{
		"name":     result.Name,
		"panicked": panicked,
	}
	expected := args.Map{
		"name":     "Bob",
		"panicked": false,
	}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadToMust returns typed -- valid", actual)
}

func Test_DeserializePayloadToSlice_Valid(t *testing.T) {
	// Arrange
	type testData struct {
		Name string `json:"Name"`
	}
	items := []testData{{Name: "A"}, {Name: "B"}}
	jsonBytes, _ := corejson.Serialize.Raw(items)
	pw := &corepayload.PayloadWrapper{Payloads: jsonBytes}

	// Act
	result, err := corepayload.DeserializePayloadToSlice[testData](pw)

	// Assert
	actual := args.Map{
		"length":  len(result),
		"noError": err == nil,
	}
	expected := args.Map{
		"length":  2,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadToSlice returns slice -- valid", actual)
}

func Test_DeserializePayloadToSlice_NilWrapper(t *testing.T) {
	// Arrange
	type testData struct{ Name string }

	// Act
	result, err := corepayload.DeserializePayloadToSlice[testData](nil)

	// Assert
	actual := args.Map{
		"length":   len(result),
		"hasError": err != nil,
	}
	expected := args.Map{
		"length":   0,
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadToSlice returns error -- nil", actual)
}

func Test_DeserializePayloadToSliceMust_Panics(t *testing.T) {
	// Arrange
	type testData struct{ Name string }
	panicked := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		corepayload.DeserializePayloadToSliceMust[testData](nil)
	}()

	// Assert
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadToSliceMust panics -- nil", actual)
}

// ── DeserializeAttributesPayloadTo ──

func Test_DeserializeAttributesPayloadTo_Valid(t *testing.T) {
	// Arrange
	type cfg struct {
		Enabled bool `json:"Enabled"`
	}
	jsonBytes, _ := corejson.Serialize.Raw(cfg{Enabled: true})
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes(jsonBytes)

	// Act
	result, err := corepayload.DeserializeAttributesPayloadTo[cfg](attr)

	// Assert
	actual := args.Map{
		"enabled": result.Enabled,
		"noError": err == nil,
	}
	expected := args.Map{
		"enabled": true,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "DeserializeAttributesPayloadTo returns typed -- valid", actual)
}

func Test_DeserializeAttributesPayloadTo_NilAttr(t *testing.T) {
	// Arrange
	type cfg struct{ Enabled bool }

	// Act
	_, err := corepayload.DeserializeAttributesPayloadTo[cfg](nil)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "DeserializeAttributesPayloadTo returns error -- nil", actual)
}

func Test_DeserializeAttributesPayloadToMust_Panics(t *testing.T) {
	// Arrange
	type cfg struct{ Enabled bool }
	panicked := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		corepayload.DeserializeAttributesPayloadToMust[cfg](nil)
	}()

	// Assert
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "DeserializeAttributesPayloadToMust panics -- nil", actual)
}

func Test_DeserializeAttributesPayloadToSlice_Valid(t *testing.T) {
	// Arrange
	type item struct {
		V int `json:"V"`
	}
	jsonBytes, _ := corejson.Serialize.Raw([]item{{V: 1}, {V: 2}})
	attr := corepayload.New.Attributes.UsingDynamicPayloadBytes(jsonBytes)

	// Act
	result, err := corepayload.DeserializeAttributesPayloadToSlice[item](attr)

	// Assert
	actual := args.Map{
		"length":  len(result),
		"noError": err == nil,
	}
	expected := args.Map{
		"length":  2,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "DeserializeAttributesPayloadToSlice returns slice -- valid", actual)
}
