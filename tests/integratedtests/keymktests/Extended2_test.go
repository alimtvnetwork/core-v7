package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/keymk"
)

// ============================================================================
// KeyJson: Serialize, MarshalJSON, UnmarshalJSON, JsonModel, JsonString
// ============================================================================

func Test_KeyJson_Serialize_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")

	// Act
	bytes, err := key.Serialize()

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Serialize should not error:", actual)
	actual := args.Map{"result": len(bytes) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Serialize should return non-empty bytes", actual)
}

func Test_KeyJson_MarshalUnmarshal_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a", "b")

	// Act
	data, err := key.MarshalJSON()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MarshalJSON error:", actual)

	var restored keymk.Key
	err = restored.UnmarshalJSON(data)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON error:", actual)

	// Assert
	actual := args.Map{"result": restored.MainName() != "root"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected main 'root', got ''", actual)
}

func Test_KeyJson_JsonModel_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")

	// Act
	model := key.JsonModel()

	// Assert
	actual := args.Map{"result": model.MainName != "root"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root', got ''", actual)
}

func Test_KeyJson_JsonModelAny_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")

	// Act
	result := key.JsonModelAny()

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonModelAny should not be nil", actual)
}

func Test_KeyJson_JsonString_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")

	// Act
	result := key.JsonString()

	// Assert - just verify no panic
	_ = result
}

func Test_KeyJson_Json_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")

	// Act
	result := key.Json()

	// Assert
	actual := args.Map{"result": result.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Json() should not have error:", actual)
}

func Test_KeyJson_JsonPtr_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")

	// Act
	result := key.JsonPtr()

	// Assert
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonPtr should not be nil", actual)
}

func Test_KeyJson_ParseInjectUsingJson_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	jsonResult := key.JsonPtr()

	// Act
	var target keymk.Key
	parsed, err := target.ParseInjectUsingJson(jsonResult)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson error:", actual)
	actual := args.Map{"result": parsed == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-nil", actual)
}

func Test_KeyJson_ParseInjectUsingJsonMust_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	jsonResult := key.JsonPtr()

	// Act
	var target keymk.Key
	parsed := target.ParseInjectUsingJsonMust(jsonResult)

	// Assert
	actual := args.Map{"result": parsed == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust should return non-nil", actual)
}

func Test_KeyJson_AsJsonContractsBinder_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root")
	binder := key.AsJsonContractsBinder()
	actual := args.Map{"result": binder == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_KeyJson_AsJsoner_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root")
	jsoner := key.AsJsoner()
	actual := args.Map{"result": jsoner == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_KeyJson_JsonParseSelfInject_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	jsonResult := key.JsonPtr()

	var target keymk.Key
	err := target.JsonParseSelfInject(jsonResult)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject error:", actual)
}

func Test_KeyJson_AsJsonParseSelfInjector_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root")
	injector := key.AsJsonParseSelfInjector()
	actual := args.Map{"result": injector == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_KeyJson_AsJsonMarshaller_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root")
	m := key.AsJsonMarshaller()
	actual := args.Map{"result": m == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

// ============================================================================
// TemplateReplacer
// ============================================================================

func Test_TemplateReplacer_IntRange_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Curly("root", "id")
	key.Finalized()
	tr := key.TemplateReplacer()

	// Act
	result := tr.IntRange(true, "id", 0, 2)

	// Assert
	actual := args.Map{"result": len(result) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 items", actual)
}

func Test_TemplateReplacer_RequestIntRange_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Curly("root", "id")
	key.Finalized()
	tr := key.TemplateReplacer()

	// Act
	result := tr.RequestIntRange(true, keymk.TempReplace{
		KeyName: "id",
		Range:   keymk.Range{Start: 1, End: 3},
	})

	// Assert
	actual := args.Map{"result": len(result) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 items", actual)
}

func Test_TemplateReplacer_CompileUsingReplacerMap_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Curly("root", "name")
	key.Finalized()
	tr := key.TemplateReplacer()

	// Act
	result := tr.CompileUsingReplacerMap(true, map[string]string{
		"root": "myRoot",
		"name": "myName",
	})

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_TemplateReplacer_CompileUsingReplacerMap_Empty_Ext2(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Curly("root")
	key.Finalized()
	tr := key.TemplateReplacer()

	// Act
	result := tr.CompileUsingReplacerMap(true, map[string]string{})

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return compiled chain", actual)
}

// ============================================================================
// FixedLegend
// ============================================================================

func Test_FixedLegend_FormatKeyMap_Ext2(t *testing.T) {
	// Act
	format, replacerMap := keymk.FixedLegend.FormatKeyMap(
		"r", "p", "g", "s", "u", "i",
	)

	// Assert
	actual := args.Map{"result": format == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "format should not be empty", actual)
	actual := args.Map{"result": len(replacerMap) != 6}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 6 replacers", actual)
}

func Test_FixedLegend_Compile_Ext2(t *testing.T) {
	// Act
	result := keymk.FixedLegend.Compile(
		false, "r", "p", "g", "s", "u", "i",
	)

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_FixedLegend_CompileKeepFormatOnEmpty_Ext2(t *testing.T) {
	// Act
	result := keymk.FixedLegend.CompileKeepFormatOnEmpty(
		"r", "p", "", "s", "u", "i",
	)

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// ============================================================================
// KeyWithLegend additional methods
// ============================================================================

func Test_KeyWithLegend_NoLegend_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	actual := args.Map{"result": k.IsIgnoreLegendAttachments()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should ignore legend attachments", actual)
}

func Test_KeyWithLegend_Create_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.Create(keymk.JoinerOption, "r", "p", "g")
	actual := args.Map{"result": k.IsIgnoreLegendAttachments()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not ignore legend attachments", actual)
}

func Test_KeyWithLegend_ShortLegend_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.ShortLegend(keymk.JoinerOption, "r", "p", "g")
	actual := args.Map{"result": k.IsIgnoreLegendAttachments()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not ignore legend attachments", actual)
}

func Test_KeyWithLegend_NoLegendPackage_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegendPackage(false, keymk.JoinerOption, "r", "g")
	actual := args.Map{"result": k.IsIgnoreLegendAttachments()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should ignore legend", actual)
}

func Test_KeyWithLegend_Getters_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.All(keymk.JoinerOption, keymk.FullLegends, true, "r", "p", "g", "s")
	actual := args.Map{"result": k.RootName() != "r"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RootName mismatch", actual)
	actual := args.Map{"result": k.PackageName() != "p"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PackageName mismatch", actual)
	actual := args.Map{"result": k.GroupName() != "g"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GroupName mismatch", actual)
	actual := args.Map{"result": k.StateName() != "s"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StateName mismatch", actual)
}

func Test_KeyWithLegend_Item_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.Item("myitem")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_ItemString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemString("myitem")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_ItemInt_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemInt(42)
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_ItemUInt_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemUInt(42)
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_GroupItemIntRange_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupItemIntRange("grp", 0, 2)
	actual := args.Map{"result": len(result) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_KeyWithLegend_GroupUIntRange_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUIntRange(0, 2)
	actual := args.Map{"result": len(result) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_KeyWithLegend_ItemIntRange_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemIntRange(0, 2)
	actual := args.Map{"result": len(result) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_KeyWithLegend_ItemUIntRange_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemUIntRange(0, 2)
	actual := args.Map{"result": len(result) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_KeyWithLegend_Group_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.Group("myg")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_GroupString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupString("myg")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_UpToGroup_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.UpToGroup("myg")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_UpToGroupString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.UpToGroupString("myg")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_ItemWithoutUser_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemWithoutUser("item1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_ItemWithoutUserGroup_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemWithoutUserGroup("item1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_ItemWithoutUserStateGroup_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemWithoutUserStateGroup("item1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_GroupUser_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUser("g1", "u1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_GroupUserString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUserString("g1", "u1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_GroupUInt_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUInt(1)
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_GroupByte_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupByte(1)
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_GroupUserByte_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUserByte(1, 2)
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_GroupUserItem_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUserItem("g1", "u1", "i1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_GroupStateUserItem_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupStateUserItem("g1", "s1", "u1", "i1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_StateUserItem_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.StateUserItem("s1", "u1", "i1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_StateUser_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.StateUser("s1", "u1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_GroupStateUserItemString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupStateUserItemString("g1", "s1", "u1", "i1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_GroupUserItemString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUserItemString("g1", "u1", "i1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_GroupUserItemUint_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUserItemUint(1, 2, 3)
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_GroupUserItemInt_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupUserItemInt(1, 2, 3)
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_GroupItem_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupItem("g1", "i1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_StateItem_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.StateItem("s1", "i1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_GroupItemString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupItemString("g1", "i1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_GroupStateItemString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.GroupStateItemString("g1", "s1", "i1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_StateItemString_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.StateItemString("s1", "i1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_Compile_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.Compile("i1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_CompileDefault_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.CompileDefault()
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_CompileUsingJoiner_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.CompileUsingJoiner("/")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_CompileStrings_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.CompileStrings()
	actual := args.Map{"result": len(result) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty slice", actual)
}

func Test_KeyWithLegend_Strings_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.Strings()
	actual := args.Map{"result": len(result) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty slice", actual)
}

func Test_KeyWithLegend_CompileItemUsingJoiner_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.CompileItemUsingJoiner("/", "i1")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_KeyWithLegend_Clone_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	cloned := k.Clone()
	actual := args.Map{"result": cloned == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone should not be nil", actual)
}

func Test_KeyWithLegend_CloneUsing_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	cloned := k.CloneUsing("newGroup")
	actual := args.Map{"result": cloned == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CloneUsing should not be nil", actual)
	actual := args.Map{"result": cloned.GroupName() != "newGroup"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'newGroup', got ''", actual)
}

func Test_KeyWithLegend_NilCloneUsing_Ext2(t *testing.T) {
	var k *keymk.KeyWithLegend
	cloned := k.CloneUsing("newGroup")
	actual := args.Map{"result": cloned != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil CloneUsing should return nil", actual)
}

func Test_KeyWithLegend_OutputItemsArray_WithLegend_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.All(keymk.JoinerOption, keymk.FullLegends, true, "r", "p", "g", "s")
	request := keymk.KeyLegendCompileRequest{
		GroupId:   "g1",
		StateName: "s1",
		UserId:    "u1",
		ItemId:    "i1",
	}
	result := k.OutputItemsArray(request)
	actual := args.Map{"result": len(result) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty items", actual)
}

func Test_KeyWithLegend_FinalStrings_WithBrackets_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.BracketJoinerOption, "r", "p", "g")
	request := keymk.KeyLegendCompileRequest{
		GroupId: "g1",
		ItemId:  "i1",
	}
	result := k.FinalStrings(request)
	actual := args.Map{"result": len(result) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty items", actual)
}

// ============================================================================
// Key: JoinUsingOption, CompileReplaceCurlyKeyMapUsingItems, CompileReplaceMapUsingItemsOption
// ============================================================================

func Test_Key_JoinUsingOption_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	opt := &keymk.Option{
		Joiner:           "/",
		IsSkipEmptyEntry: true,
	}
	result := key.JoinUsingOption(opt, "b")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Key_CompileReplaceCurlyKeyMapUsingItems_Ext2(t *testing.T) {
	key := keymk.NewKey.Curly("root", "name")
	compiled := key.CompileReplaceCurlyKeyMapUsingItems(
		map[string]string{"root": "myroot", "name": "myname"},
		"extra",
	)
	actual := args.Map{"result": compiled == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Key_CompileReplaceMapUsingItemsOption_NoCurly_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root", "name")
	compiled := key.CompileReplaceMapUsingItemsOption(
		false,
		map[string]string{"root": "myroot"},
	)
	actual := args.Map{"result": compiled == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Key_CompileReplaceMapUsingItemsOption_EmptyMap_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root")
	compiled := key.CompileReplaceMapUsingItemsOption(
		true,
		map[string]string{},
	)
	actual := args.Map{"result": compiled != "root"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root', got ''", actual)
}

// ============================================================================
// Key: Finalized then Compile with additional items
// ============================================================================

func Test_Key_Finalized_CompileWithAdditional_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	key.Finalized()

	// Compile with additional should append
	result := key.Compile("b")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Key_Finalized_CompileStringsWithAdditional_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	key.Finalized()

	result := key.CompileStrings("b")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Key_Finalized_CompileNoAdditional_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	key.Finalized()

	result := key.Compile()
	actual := args.Map{"result": result != "root-a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root-a', got ''", actual)
}

func Test_Key_Finalized_CompileStringsNoAdditional_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root", "a")
	key.Finalized()

	result := key.CompileStrings()
	actual := args.Map{"result": result != "root-a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root-a', got ''", actual)
}

// ============================================================================
// KeyLegendCompileRequest constructors
// ============================================================================

func Test_KeyLegendCompileRequest_NewKeyLegend_Ext2(t *testing.T) {
	req := keymk.KeyLegendCompileRequest{GroupId: "g1"}
	k := req.NewKeyLegend(keymk.JoinerOption, keymk.ShortLegends, false, "r", "p", "s")
	actual := args.Map{"result": k == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_KeyLegendCompileRequest_NewKeyLegendDefaults_Ext2(t *testing.T) {
	req := keymk.KeyLegendCompileRequest{GroupId: "g1"}
	k := req.NewKeyLegendDefaults("r", "p", "s")
	actual := args.Map{"result": k == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

// ============================================================================
// NewKey creators: PathTemplatePrefixRelativeIdDefault, PathTemplatePrefixRelativeIdFileDefault
// ============================================================================

func Test_NewKey_PathTemplatePrefixRelativeIdDefault_Ext2(t *testing.T) {
	key := keymk.NewKey.PathTemplatePrefixRelativeIdDefault()
	actual := args.Map{"result": key.Compile() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_NewKey_PathTemplatePrefixRelativeIdFileDefault_Ext2(t *testing.T) {
	key := keymk.NewKey.PathTemplatePrefixRelativeIdFileDefault()
	actual := args.Map{"result": key.Compile() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_NewKey_CurlyStrings_Ext2(t *testing.T) {
	key := keymk.NewKey.CurlyStrings("root", "a")
	actual := args.Map{"result": key.Compile() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_NewKey_SquareBracketsStrings_Ext2(t *testing.T) {
	key := keymk.NewKey.SquareBracketsStrings("root", "a")
	actual := args.Map{"result": key.Compile() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_NewKey_ParenthesisStrings_Ext2(t *testing.T) {
	key := keymk.NewKey.ParenthesisStrings("root", "a")
	actual := args.Map{"result": key.Compile() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_NewKey_StringsWithOptions_Ext2(t *testing.T) {
	key := keymk.NewKey.StringsWithOptions(keymk.JoinerOption, "root", "a")
	actual := args.Map{"result": key.Compile() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_NewKey_OptionMain_Ext2(t *testing.T) {
	key := keymk.NewKey.OptionMain(keymk.JoinerOption, "root")
	actual := args.Map{"result": key.Compile() != "root"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'root', got ''", actual)
}

// ============================================================================
// Key: ItemEnumByte via KeyWithLegend (need to pass a mock enuminf.ByteEnumNamer)
// ============================================================================

func Test_KeyWithLegend_ItemEnumByte_Ext2(t *testing.T) {
	k := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")
	result := k.ItemEnumByte(mockByteEnumNamer{name: "test-item"})
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// ============================================================================
// Key: AppendChainStrings skip empty
// ============================================================================

func Test_Key_AppendChainStrings_SkipEmpty_Ext2(t *testing.T) {
	key := keymk.NewKey.Default("root")
	key.AppendChainStrings("", "a", "", "b")
	actual := args.Map{
		"length": key.Length(),
	}
	expected := args.Map{
		"length": 2,
	}
	actual := args.Map{"result": actual["length"] != expected["length"]}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 chains (empty skipped)", actual)
}
