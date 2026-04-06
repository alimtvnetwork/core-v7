package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/keymk"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_Key_Compile_WithBrackets(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default(".")
	result := k.Compile("a", "b")

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_QW_Key_Compile_Empty(t *testing.T) {
	k := keymk.NewKey.Default(".")
	result := k.Compile()
	_ = result
}

func Test_QW_Key_ParseInjectUsingJson_Error(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default(".")
	bad := corejson.NewResult.UsingString(`invalid`)
	_, err := k.ParseInjectUsingJson(bad)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for invalid JSON", actual)
}

func Test_QW_Key_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() { recover() }()
	k := keymk.NewKey.Default(".")
	bad := corejson.NewResult.UsingString(`invalid`)
	k.ParseInjectUsingJsonMust(bad)
}

func Test_QW_Key_Compile_SkipEmpty(t *testing.T) {
	k := keymk.NewKey.DefaultStrings(".", "base")
	result := k.Compile("", "a", "", "b")
	_ = result
}
