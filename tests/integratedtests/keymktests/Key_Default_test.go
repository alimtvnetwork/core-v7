package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/keymk"
)

// ── Key basic methods ──

func Test_Key_Default_Compile(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a", "b")

	// Act
	actual := args.Map{
		"compile":     key.Compile(),
		"mainName":    key.MainName(),
		"length":      key.Length(),
		"isEmpty":     key.IsEmpty(),
		"isComplete":  key.IsComplete(),
		"string":      key.String(),
		"name":        key.Name(),
		"keyCompiled": key.KeyCompiled(),
	}

	// Assert
	expected := args.Map{
		"compile":     key.Compile(),
		"mainName":    "root",
		"length":      2,
		"isEmpty":     false,
		"isComplete":  false,
		"string":      key.Compile(),
		"name":        key.Compile(),
		"keyCompiled": key.Compile(),
	}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- Default Compile", actual)
}

func Test_Key_CompileDefault(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	result := key.CompileDefault()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- CompileDefault", actual)
}

func Test_Key_Nil(t *testing.T) {
	// Arrange
	var key *keymk.Key

	// Act
	actual := args.Map{
		"length":    key.Length(),
		"keyChains": key.KeyChains() == nil,
		"allRaw":    key.AllRawItems() == nil,
		"hasIn":     key.HasInChains("x"),
	}

	// Assert
	expected := args.Map{
		"length": 0,
		"keyChains": true,
		"allRaw": true,
		"hasIn": false,
	}
	expected.ShouldBeEqual(t, 0, "Key returns nil -- nil", actual)
}

func Test_Key_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var key *keymk.Key

	// Act
	actual := args.Map{"isNil": key.ClonePtr() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Key returns nil -- ClonePtr nil", actual)
}

func Test_Key_AppendChainStrings(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")
	key.AppendChainStrings("a", "b")

	// Act
	actual := args.Map{"length": key.Length()}

	// Assert
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- AppendChainStrings", actual)
}

func Test_Key_AppendChainStrings_SkipEmpty(t *testing.T) {
	// Arrange
	key := keymk.NewKey.AllStrings(keymk.JoinerOption, "root", "a")
	key.AppendChainStrings("", "b")

	// Act
	actual := args.Map{"length": key.Length()}

	// Assert
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "Key returns empty -- AppendChainStrings skip empty", actual)
}

func Test_Key_AppendChainKeys(t *testing.T) {
	// Arrange
	key1 := keymk.NewKey.Default("root", "a")
	key2 := keymk.NewKey.Default("sub", "b")
	key1.AppendChainKeys(key2, nil)

	// Act
	actual := args.Map{"length": key1.Length()}

	// Assert
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- AppendChainKeys", actual)
}

func Test_Key_AppendChainKeys_Empty(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	result := key.AppendChainKeys()

	// Act
	actual := args.Map{"length": result.Length()}

	// Assert
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "Key returns empty -- AppendChainKeys empty", actual)
}

func Test_Key_HasInChains(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a", "b")

	// Act
	actual := args.Map{
		"hasA":    key.HasInChains("a"),
		"hasC":    key.HasInChains("c"),
	}

	// Assert
	expected := args.Map{
		"hasA": true,
		"hasC": false,
	}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- HasInChains", actual)
}

func Test_Key_ConcatNewUsingKeys(t *testing.T) {
	// Arrange
	key1 := keymk.NewKey.Default("root", "a")
	key2 := keymk.NewKey.Default("sub", "b")
	result := key1.ConcatNewUsingKeys(key2)

	// Act
	actual := args.Map{
		"length": result.Length(),
		"originalLen": key1.Length(),
	}

	// Assert
	expected := args.Map{
		"length": 3,
		"originalLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- ConcatNewUsingKeys", actual)
}

func Test_Key_ClonePtr(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	cloned := key.ClonePtr("b")

	// Act
	actual := args.Map{
		"clonedLen": cloned.Length(),
		"originalLen": key.Length(),
	}

	// Assert
	expected := args.Map{
		"clonedLen": 2,
		"originalLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- ClonePtr", actual)
}

func Test_Key_Strings(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a", "b")

	// Act
	actual := args.Map{"len": len(key.Strings())}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- Strings", actual)
}

func Test_Key_CompiledChain_NotComplete(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")

	// Act
	actual := args.Map{"result": key.CompiledChain()}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- CompiledChain not complete", actual)
}

// ── Finalized ──

func Test_Key_Finalized_FromKeyDefault(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	key.Finalized("b")

	// Act
	actual := args.Map{
		"isComplete":    key.IsComplete(),
		"compiledChain": key.CompiledChain() != "",
	}

	// Assert
	expected := args.Map{
		"isComplete": true,
		"compiledChain": true,
	}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- Finalized", actual)
}

func Test_Key_Finalized_Compile_Additional(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	key.Finalized()
	result := key.Compile("extra")

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"containsExtra": len(result) > len(key.CompiledChain()),
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"containsExtra": true,
	}
	expected.ShouldBeEqual(t, 0, "Key returns non-empty -- Finalized Compile with extra", actual)
}

func Test_Key_Finalized_CompileStrings_Additional(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	key.Finalized()
	result := key.CompileStrings("extra")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key returns non-empty -- Finalized CompileStrings with extra", actual)
}

func Test_Key_Finalized_Compile_NoAdditional(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	key.Finalized()
	result := key.Compile()

	// Act
	actual := args.Map{
		"result": result,
		"chain": key.CompiledChain(),
	}

	// Assert
	expected := args.Map{
		"result": key.CompiledChain(),
		"chain": key.CompiledChain(),
	}
	expected.ShouldBeEqual(t, 0, "Key returns empty -- Finalized Compile no additional", actual)
}

// ── CompileKeys ──

func Test_Key_CompileKeys_FromKeyDefault(t *testing.T) {
	// Arrange
	key1 := keymk.NewKey.Default("root")
	key2 := keymk.NewKey.Default("sub", "a")
	result := key1.CompileKeys(key2, nil)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- CompileKeys", actual)
}

func Test_Key_CompileKeys_Empty_FromKeyDefault(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	result := key.CompileKeys()

	// Act
	actual := args.Map{
		"result": result,
		"compile": key.Compile(),
	}

	// Assert
	expected := args.Map{
		"result": key.Compile(),
		"compile": key.Compile(),
	}
	expected.ShouldBeEqual(t, 0, "Key returns empty -- CompileKeys empty", actual)
}

// ── CompileReplaceCurlyKeyMap ──

func Test_Key_CompileReplaceCurlyKeyMap_FromKeyDefault(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "{name}", "{id}")
	result := key.CompileReplaceCurlyKeyMap(map[string]string{"name": "test", "id": "42"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- CompileReplaceCurlyKeyMap", actual)
}

func Test_Key_CompileReplaceMapUsingItemsOption_NoCurly_FromKeyDefault(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "NAME")
	result := key.CompileReplaceMapUsingItemsOption(false, map[string]string{"NAME": "test"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key returns empty -- CompileReplaceMap no curly", actual)
}

func Test_Key_CompileReplaceMapUsingItemsOption_EmptyMap_FromKeyDefault(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	result := key.CompileReplaceMapUsingItemsOption(true, map[string]string{})

	// Act
	actual := args.Map{
		"result": result,
		"compile": key.Compile(),
	}

	// Assert
	expected := args.Map{
		"result": key.Compile(),
		"compile": key.Compile(),
	}
	expected.ShouldBeEqual(t, 0, "Key returns empty -- CompileReplaceMap empty map", actual)
}

// ── JoinUsingJoiner / JoinUsingOption ──

func Test_Key_JoinUsingJoiner(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a", "b")
	result := key.JoinUsingJoiner("/")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- JoinUsingJoiner", actual)
}

func Test_Key_JoinUsingOption_FromKeyDefault(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	result := key.JoinUsingOption(keymk.CurlyBracePathJoinerOption, "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- JoinUsingOption", actual)
}

// ── IntRange / IntRangeEnding ──

func Test_Key_IntRange_FromKeyDefault(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("item")
	result := key.IntRange(1, 3)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- IntRange", actual)
}

func Test_Key_IntRangeEnding_FromKeyDefault(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("item")
	result := key.IntRangeEnding(2)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- IntRangeEnding", actual)
}

// ── JSON ──

func Test_Key_JsonString(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	result := key.JsonString()
	notEmpty := result != ""

	// Act
	actual := args.Map{"notEmpty": notEmpty}

	// Assert
	expected := args.Map{"notEmpty": notEmpty}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- JsonString", actual)
}

func Test_Key_MarshalUnmarshalJSON(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	data, err := key.MarshalJSON()
	var key2 keymk.Key
	err2 := key2.UnmarshalJSON(data)

	// Act
	actual := args.Map{
		"noErr":     err == nil,
		"noErr2":    err2 == nil,
		"notEmpty":  len(data) > 0,
		"mainName":  key2.MainName(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"noErr2": true,
		"notEmpty": true,
		"mainName": "root",
	}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- MarshalUnmarshalJSON", actual)
}

func Test_Key_UnmarshalJSON_Invalid(t *testing.T) {
	// Arrange
	var key keymk.Key
	err := key.UnmarshalJSON([]byte(`invalid`))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Key returns error -- UnmarshalJSON invalid", actual)
}

func Test_Key_Serialize(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	data, err := key.Serialize()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": len(data) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- Serialize", actual)
}

func Test_Key_Json(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")
	j := key.Json()
	notEmpty := j.JsonString() != ""

	// Act
	actual := args.Map{"notEmpty": notEmpty}

	// Assert
	expected := args.Map{"notEmpty": notEmpty}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- Json", actual)
}

func Test_Key_JsonPtr(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")
	j := key.JsonPtr()

	// Act
	actual := args.Map{"notNil": j != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- JsonPtr", actual)
}

func Test_Key_JsonModel(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	model := key.JsonModel()

	// Act
	actual := args.Map{"mainName": model.MainName}

	// Assert
	expected := args.Map{"mainName": "root"}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- JsonModel", actual)
}

func Test_Key_JsonModelAny(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")

	// Act
	actual := args.Map{"notNil": key.JsonModelAny() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- JsonModelAny", actual)
}

func Test_Key_Contracts(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")

	// Act
	actual := args.Map{
		"jsonBinder":     key.AsJsonContractsBinder() != nil,
		"jsoner":         key.AsJsoner() != nil,
		"selfInjector":   key.AsJsonParseSelfInjector() != nil,
		"jsonMarshaller": key.AsJsonMarshaller() != nil,
	}

	// Assert
	expected := args.Map{
		"jsonBinder": true, "jsoner": true,
		"selfInjector": true, "jsonMarshaller": true,
	}
	expected.ShouldBeEqual(t, 0, "Key returns correct value -- contracts", actual)
}

// ── TemplateReplacer ──

func Test_Key_TemplateReplacer_IntRange(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("item")
	tr := key.TemplateReplacer()
	result := tr.IntRange(false, "id", 1, 3)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "TemplateReplacer returns correct value -- IntRange", actual)
}

func Test_Key_TemplateReplacer_RequestIntRange(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("item")
	tr := key.TemplateReplacer()
	req := keymk.TempReplace{KeyName: "id", Range: keymk.Range{Start: 0, End: 2}}
	result := tr.RequestIntRange(false, req)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "TemplateReplacer returns correct value -- RequestIntRange", actual)
}

// ── Factory methods ──

func Test_NewKey_Factories(t *testing.T) {
	// Act
	actual := args.Map{
		"default":         keymk.NewKey.Default("r", "a").MainName(),
		"defaultStrings":  keymk.NewKey.DefaultStrings("r", "a").MainName(),
		"defaultMain":     keymk.NewKey.DefaultMain("r").MainName(),
		"curly":           keymk.NewKey.Curly("r", "a").MainName(),
		"curlyStrings":    keymk.NewKey.CurlyStrings("r", "a").MainName(),
		"squareBrackets":  keymk.NewKey.SquareBrackets("r", "a").MainName(),
		"sqBracketsStr":   keymk.NewKey.SquareBracketsStrings("r", "a").MainName(),
		"parenthesis":     keymk.NewKey.Parenthesis("r", "a").MainName(),
		"parenthesisStr":  keymk.NewKey.ParenthesisStrings("r", "a").MainName(),
		"pathTemplate":    keymk.NewKey.PathTemplate("r", "a").MainName(),
		"pathTemplateDef": keymk.NewKey.PathTemplateDefault("a").MainName() != "",
		"optionMain":      keymk.NewKey.OptionMain(keymk.JoinerOption, "r").MainName(),
	}

	// Assert
	expected := args.Map{
		"default": "r", "defaultStrings": "r", "defaultMain": "r",
		"curly": "r", "curlyStrings": "r",
		"squareBrackets": "r", "sqBracketsStr": "r",
		"parenthesis": "r", "parenthesisStr": "r",
		"pathTemplate": "r", "pathTemplateDef": true,
		"optionMain": "r",
	}
	expected.ShouldBeEqual(t, 0, "NewKey returns correct value -- factories", actual)
}

func Test_NewKey_AllStrings(t *testing.T) {
	// Arrange
	key := keymk.NewKey.AllStrings(keymk.JoinerOption, "root", "a", "b")

	// Act
	actual := args.Map{
		"mainName": key.MainName(),
		"len": key.Length(),
	}

	// Assert
	expected := args.Map{
		"mainName": "root",
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "NewKey returns correct value -- AllStrings", actual)
}

func Test_NewKey_StringsWithOptions(t *testing.T) {
	// Arrange
	key := keymk.NewKey.StringsWithOptions(keymk.JoinerOption, "root", "a")

	// Act
	actual := args.Map{
		"mainName": key.MainName(),
		"len": key.Length(),
	}

	// Assert
	expected := args.Map{
		"mainName": "root",
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "NewKey returns non-empty -- StringsWithOptions", actual)
}

// ── Option ──

func Test_Option_Clone(t *testing.T) {
	// Arrange
	opt := keymk.JoinerOption
	cloned := opt.Clone()

	// Act
	actual := args.Map{"joiner": cloned.Joiner}

	// Assert
	expected := args.Map{"joiner": opt.Joiner}
	expected.ShouldBeEqual(t, 0, "Option returns correct value -- Clone", actual)
}

func Test_Option_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var opt *keymk.Option

	// Act
	actual := args.Map{"isNil": opt.ClonePtr() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Option returns nil -- ClonePtr nil", actual)
}

func Test_Option_IsAddEntryRegardlessOfEmptiness_Nil(t *testing.T) {
	// Arrange
	var opt *keymk.Option

	// Act
	actual := args.Map{"result": opt.IsAddEntryRegardlessOfEmptiness()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Option returns nil -- IsAddEntry nil", actual)
}

// ── Brackets Key ──

func Test_Key_Brackets(t *testing.T) {
	// Arrange
	key := keymk.NewKey.SquareBrackets("root", "a", "b")
	result := key.Compile()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key returns non-empty -- with brackets", actual)
}
