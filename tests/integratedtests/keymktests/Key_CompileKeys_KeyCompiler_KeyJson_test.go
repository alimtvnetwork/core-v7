package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/keymk"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Key_CompileKeys(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default("root", "a")
	k2 := keymk.NewKey.Default("sub", "b")
	result := k.CompileKeys(k2)

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	// nil key in list
	result2 := k.CompileKeys(nil, k2)
	actual = args.Map{"result": result2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Key_CompileKeys_Empty(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default("root", "a")
	result := k.CompileKeys()

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected compiled", actual)
}

func Test_Key_Finalized(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default("root", "a")
	k.Finalized("extra")
	r1 := k.Compile()
	r2 := k.Compile("more")

	// Act
	actual := args.Map{"result": r1 == "" || r2 == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Key_CompileStrings_Finalized(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default("root")
	k.Finalized()
	r := k.CompileStrings("a", "b")

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r2 := k.CompileStrings()
	actual = args.Map{"result": r2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Key_CompileReplaceCurlyKeyMap(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default("root", "{name}")
	result := k.CompileReplaceCurlyKeyMap(map[string]string{"name": "world"})

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Key_CompileReplaceMapUsingItemsOption_NoCurly(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default("root", "KEY")
	result := k.CompileReplaceMapUsingItemsOption(false, map[string]string{"KEY": "val"})

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Key_CompileReplaceMapUsingItemsOption_EmptyMap(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default("root", "a")
	result := k.CompileReplaceMapUsingItemsOption(true, nil)

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Key_IntRange(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default("item")
	result := k.IntRange(0, 2)

	// Act
	actual := args.Map{"result": len(result) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_Key_IntRangeEnding(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default("item")
	result := k.IntRangeEnding(2)

	// Act
	actual := args.Map{"result": len(result) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_Key_JoinUsingOption(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default("root", "a")
	opt := &keymk.Option{Joiner: "-"}
	result := k.JoinUsingOption(opt, "b")

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_KeyJson_Serialize_Unmarshal(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default("root", "a")
	data, err := k.Serialize()

	// Act
	actual := args.Map{"result": err != nil || len(data) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected serialize success", actual)

	k2 := &keymk.Key{}
	err = k2.UnmarshalJSON(data)
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unmarshal success", actual)
}

func Test_KeyJson_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default("root", "a")
	jr := k.JsonPtr()
	k2 := &keymk.Key{}
	result, err := k2.ParseInjectUsingJson(jr)

	// Act
	actual := args.Map{"result": err != nil || result == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_KeyJson_ParseInjectUsingJsonMust(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default("root")
	jr := k.JsonPtr()
	k2 := &keymk.Key{}
	result := k2.ParseInjectUsingJsonMust(jr)

	// Act
	actual := args.Map{"result": result == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_KeyJson_AsJsonContractsBinder(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default("root")

	// Act
	actual := args.Map{"result": k.AsJsonContractsBinder() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": k.AsJsoner() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": k.AsJsonParseSelfInjector() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": k.AsJsonMarshaller() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_KeyJson_JsonParseSelfInject(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default("root")
	jr := k.JsonPtr()
	k2 := &keymk.Key{}
	err := k2.JsonParseSelfInject(jr)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_KeyJson_TemplateReplacer(t *testing.T) {
	k := keymk.NewKey.Default("root", "{name}")
	tr := k.TemplateReplacer()
	_ = tr
}
