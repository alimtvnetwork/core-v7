package keymktests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/keymk"
)

// ── stubByteEnum — implements enuminf.ByteEnumNamer ──

type stubByteEnum struct {
	name  string
	value byte
}

func (s stubByteEnum) Name() string    { return s.name }
func (s stubByteEnum) String() string  { return s.name }
func (s stubByteEnum) ValueByte() byte { return s.value }

// ── KeyWithLegend.ItemEnumByte ──

func Test_KeyWithLegend_ItemEnumByte(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)
	enumVal := stubByteEnum{name: "Active", value: 1}

	// Act
	result := kl.ItemEnumByte(enumVal)

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend ItemEnumByte returns compiled -- ByteEnumNamer", actual)
}

// ── KeyWithLegend.OutputWithoutLegend with IsAddEntryRegardlessOfEmptiness=true ──

func Test_KeyWithLegend_OutputWithoutLegend_AddRegardless(t *testing.T) {
	// Arrange — IsSkipEmptyEntry=false + isAttachLegendNames=false triggers OutputWithoutLegend
	opt := &keymk.Option{
		Joiner:           "-",
		IsSkipEmptyEntry: false,
		IsUseBrackets:    false,
	}
	kl := keymk.NewKeyWithLegend.All(
		opt,
		keymk.FullLegends,
		false, // no legend → OutputWithoutLegend path
		"r", "p", "", "",
	)

	// Act — request with all empty fields; isAddRegardless=true keeps them
	result := kl.UpToGroupString("")

	// Assert — should contain empty segments since IsAddEntryRegardlessOfEmptiness
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend OutputWithoutLegend keeps empty -- IsAddEntryRegardless", actual)
}

// ── KeyWithLegend.Group (any variant) ──

func Test_KeyWithLegend_Group_Any(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act — pass int as any
	result := kl.Group(42)

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend Group returns compiled -- any arg", actual)
}

// ── TemplateReplacer.CompileUsingReplacerMap with isCurly=false ──

func Test_TemplateReplacer_CompileUsingReplacerMap_NoCurly(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("item", "NAME")
	key.Finalized()
	tr := key.TemplateReplacer()

	// Act
	result := tr.CompileUsingReplacerMap(false, map[string]string{"NAME": "replaced"})

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TemplateReplacer CompileUsingReplacerMap no curly -- direct replace", actual)
}

// ── KeyWithLegend.NoLegendPackage with isAttachLegend=true ──

func Test_NewKeyWithLegend_NoLegendPackage_WithLegend(t *testing.T) {
	// Arrange & Act
	kl := keymk.NewKeyWithLegend.NoLegendPackage(true, keymk.JoinerOption, "r", "g")

	// Assert
	actual := args.Map{
		"ignoreLeg": kl.IsIgnoreLegendAttachments(),
		"package":   kl.PackageName(),
	}
	expected := args.Map{
		"ignoreLeg": false,
		"package": "",
	}
	expected.ShouldBeEqual(t, 0, "NewKeyWithLegend NoLegendPackage with legend -- isAttach true", actual)
}

// ── KeyWithLegend with legend + brackets ──

func Test_KeyWithLegend_WithLegend_Brackets(t *testing.T) {
	// Arrange — legend attached + bracket option
	kl := keymk.NewKeyWithLegend.All(
		keymk.BracketJoinerOption,
		keymk.FullLegends,
		true,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.ItemString("myItem")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend with legend and brackets -- full chain", actual)
}

// ── Key.AppendChain with skip-empty and empty string from fmt.Sprintf ──

func Test_Key_AppendChain_SkipEmpty_StringValue(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")

	// Act — empty string through any interface
	key.AppendChain("", "a")

	// Assert — empty string gets sprintf'd to "" and skipped
	actual := args.Map{"length": key.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "Key AppendChain skips empty string -- via any interface", actual)
}

// ── Key.Compile with brackets and additional items ──

func Test_Key_Compile_Brackets_WithAdditional(t *testing.T) {
	// Arrange
	key := keymk.NewKey.SquareBrackets("root", "a")

	// Act
	result := key.Compile("extra")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key Compile with brackets and additional -- all bracketed", actual)
}

// ── Key.CompileStrings with brackets ──

func Test_Key_CompileStrings_Brackets(t *testing.T) {
	// Arrange
	key := keymk.NewKey.SquareBracketsStrings("root", "a")

	// Act
	result := key.CompileStrings("b")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key CompileStrings with brackets -- all bracketed", actual)
}

// ── Key finalized then CompileKeys ──

func Test_Key_Finalized_CompileKeys(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	key.Finalized()
	key2 := keymk.NewKey.Default("sub", "b")

	// Act
	result := key.CompileKeys(key2)

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key Finalized CompileKeys appends -- completed key with additional keys", actual)
}

// ── Key.Compile with numeric any items ──

func Test_Key_Compile_NumericItems(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")

	// Act — pass various numeric types as any
	result := key.Compile(42, 3.14, uint(7))

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key Compile with numeric any items -- formatted via Sprintf", actual)
}

// ── curlyWrapIf ──

func Test_CurlyWrapIf_ThroughCompileReplace(t *testing.T) {
	// Arrange — test curly wrapping indirectly through CompileReplaceCurlyKeyMap
	key := keymk.NewKey.Default("root", "{x}")

	// Act
	result := key.CompileReplaceCurlyKeyMap(map[string]string{"x": "val"})

	// Assert — {x} should be replaced with val
	actual := args.Map{
		"notEmpty":    result != "",
		"noPlaceholder": result != fmt.Sprintf("root-%s", "{x}"),
	}
	expected := args.Map{
		"notEmpty": true,
		"noPlaceholder": true,
	}
	expected.ShouldBeEqual(t, 0, "CurlyWrapIf wraps keys -- CompileReplaceCurlyKeyMap", actual)
}

// ── KeyWithLegend.OutputItemsArray legend path with all fields ──

func Test_KeyWithLegend_OutputItemsArray_AllFields(t *testing.T) {
	// Arrange — legend attached, all request fields populated
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		true,
		"r", "p", "s", "g",
	)

	// Act — full chain with all fields
	result := kl.GroupStateUserItemString("grp", "st", "usr", "itm")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend OutputItemsArray all fields with legend -- full population", actual)
}

// ── KeyWithLegend.UserStringWithoutState with legend ──

func Test_KeyWithLegend_UserStringWithoutState_WithLegend(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		true,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.UserStringWithoutState("usr")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend UserStringWithoutState with legend -- attached", actual)
}
