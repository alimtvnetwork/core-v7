package coretestcasestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── GenericGherkins Getters ──

func Test_GenericGherkins_IsFailedToMatch_FromGenericGherkinsIsFai(t *testing.T) {
	// Arrange
	g := &coretestcases.StringBoolGherkins{IsMatching: true}
	g2 := &coretestcases.StringBoolGherkins{IsMatching: false}

	// Act
	actual := args.Map{
		"matchingFails": g.IsFailedToMatch(),
		"notMatchFails": g2.IsFailedToMatch(),
	}

	// Assert
	expected := args.Map{
		"matchingFails": false,
		"notMatchFails": true,
	}
	expected.ShouldBeEqual(t, 0, "IsFailedToMatch returns correct value -- with args", actual)
}

func Test_GenericGherkins_HasExtraArgs(t *testing.T) {
	// Arrange
	g := &coretestcases.StringBoolGherkins{ExtraArgs: args.Map{"k": "v"}}
	g2 := &coretestcases.StringBoolGherkins{}
	var gNil *coretestcases.StringBoolGherkins

	// Act
	actual := args.Map{
		"hasExtra":    g.HasExtraArgs(),
		"noExtra":     g2.HasExtraArgs(),
		"nilHasExtra": gNil.HasExtraArgs(),
	}

	// Assert
	expected := args.Map{
		"hasExtra":    true,
		"noExtra":     false,
		"nilHasExtra": false,
	}
	expected.ShouldBeEqual(t, 0, "HasExtraArgs returns correct value -- with args", actual)
}

func Test_GenericGherkins_GetExtra(t *testing.T) {
	// Arrange
	g := &coretestcases.StringBoolGherkins{ExtraArgs: args.Map{"k": "v"}}
	var gNil *coretestcases.StringBoolGherkins

	// Act
	actual := args.Map{
		"found":    g.GetExtra("k"),
		"notFound": g.GetExtra("missing") == nil,
		"nilGet":   gNil.GetExtra("k") == nil,
	}

	// Assert
	expected := args.Map{
		"found":    "v",
		"notFound": true,
		"nilGet":   true,
	}
	expected.ShouldBeEqual(t, 0, "GetExtra returns correct value -- with args", actual)
}

func Test_GenericGherkins_GetExtraAsString(t *testing.T) {
	// Arrange
	g := &coretestcases.StringBoolGherkins{ExtraArgs: args.Map{"k": "v"}}
	var gNil *coretestcases.StringBoolGherkins
	val, ok := g.GetExtraAsString("k")
	nilVal, nilOk := gNil.GetExtraAsString("k")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
		"nilVal": nilVal,
		"nilOk": nilOk,
	}

	// Assert
	expected := args.Map{
		"val": "v",
		"ok": true,
		"nilVal": "",
		"nilOk": false,
	}
	expected.ShouldBeEqual(t, 0, "GetExtraAsString returns correct value -- with args", actual)
}

func Test_GenericGherkins_GetExtraAsBool(t *testing.T) {
	// Arrange
	g := &coretestcases.StringBoolGherkins{ExtraArgs: args.Map{"k": true}}
	var gNil *coretestcases.StringBoolGherkins
	val, ok := g.GetExtraAsBool("k")
	nilVal, nilOk := gNil.GetExtraAsBool("k")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
		"nilVal": nilVal,
		"nilOk": nilOk,
	}

	// Assert
	expected := args.Map{
		"val": true,
		"ok": true,
		"nilVal": false,
		"nilOk": false,
	}
	expected.ShouldBeEqual(t, 0, "GetExtraAsBool returns correct value -- with args", actual)
}

func Test_GenericGherkins_GetExtraAsBoolDefault(t *testing.T) {
	// Arrange
	g := &coretestcases.StringBoolGherkins{ExtraArgs: args.Map{"k": true}}
	var gNil *coretestcases.StringBoolGherkins
	val := g.GetExtraAsBoolDefault("k", false)
	nilVal := gNil.GetExtraAsBoolDefault("k", true)

	// Act
	actual := args.Map{
		"val": val,
		"nilVal": nilVal,
	}

	// Assert
	expected := args.Map{
		"val": true,
		"nilVal": true,
	}
	expected.ShouldBeEqual(t, 0, "GetExtraAsBoolDefault returns correct value -- with args", actual)
}

// ── GenericGherkins Formatting ──

func Test_GenericGherkins_ToString_FromGenericGherkinsIsFai(t *testing.T) {
	// Arrange
	g := &coretestcases.StringBoolGherkins{
		Feature: "f", Given: "g", When: "w", Then: "t",
	}
	result := g.ToString(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToString returns correct value -- with args", actual)
}

func Test_GenericGherkins_String_FromGenericGherkinsIsFai(t *testing.T) {
	// Arrange
	g := &coretestcases.StringBoolGherkins{Feature: "f"}

	// Act
	actual := args.Map{"notEmpty": g.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- with args", actual)
}

func Test_GenericGherkins_GetWithExpectation_FromGenericGherkinsIsFai(t *testing.T) {
	// Arrange
	g := &coretestcases.StringBoolGherkins{Feature: "f"}
	result := g.GetWithExpectation(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetWithExpectation returns non-empty -- with args", actual)
}

func Test_GenericGherkins_GetMessageConditional_FromGenericGherkinsIsFai(t *testing.T) {
	// Arrange
	g := &coretestcases.StringBoolGherkins{Feature: "f"}
	withExp := g.GetMessageConditional(true, 0)
	withoutExp := g.GetMessageConditional(false, 0)

	// Act
	actual := args.Map{
		"withNotEmpty":    withExp != "",
		"withoutNotEmpty": withoutExp != "",
	}

	// Assert
	expected := args.Map{
		"withNotEmpty":    true,
		"withoutNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "GetMessageConditional returns correct value -- with args", actual)
}

func Test_GenericGherkins_FullString(t *testing.T) {
	// Arrange
	g := &coretestcases.StringBoolGherkins{
		Title: "t", Feature: "f", Given: "g", When: "w", Then: "th",
		ExtraArgs: args.Map{"k": "v"},
	}
	result := g.FullString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FullString returns correct value -- with args", actual)
}

func Test_GenericGherkins_FullString_Nil_FromGenericGherkinsIsFai(t *testing.T) {
	// Arrange
	var g *coretestcases.StringBoolGherkins
	result := g.FullString()

	// Act
	actual := args.Map{"isNilMsg": result == "<nil GenericGherkins>"}

	// Assert
	expected := args.Map{"isNilMsg": true}
	expected.ShouldBeEqual(t, 0, "FullString_Nil returns nil -- with args", actual)
}

// ── GenericGherkins CaseTitle ──

func Test_GenericGherkins_CaseTitle(t *testing.T) {
	// Arrange
	gTitle := &coretestcases.StringBoolGherkins{Title: "myTitle", When: "myWhen"}
	gWhen := &coretestcases.StringBoolGherkins{When: "myWhen"}

	// Act
	actual := args.Map{
		"titleResult": gTitle.CaseTitle(),
		"whenResult":  gWhen.CaseTitle(),
	}

	// Assert
	expected := args.Map{
		"titleResult": "myTitle",
		"whenResult":  "myWhen",
	}
	expected.ShouldBeEqual(t, 0, "CaseTitle returns correct value -- with args", actual)
}
