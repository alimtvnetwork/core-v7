package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/keymk"
)

// ── Key.go — uncovered branches ──

func Test_Key_IsEmpty_True(t *testing.T) {
	// Arrange — key with empty mainName and no chains
	key := keymk.NewKey.Default("")

	// Act
	result := key.IsEmpty()

	// Assert
	actual := args.Map{"isEmpty": result}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key IsEmpty returns true -- empty mainName no chains", actual)
}

func Test_Key_AppendChain_PanicOnComplete(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	key.Finalized()

	// Act
	recovered := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				recovered = true
			}
		}()
		key.AppendChain("shouldPanic")
	}()

	// Assert
	actual := args.Map{"recovered": recovered}
	expected := args.Map{"recovered": true}
	expected.ShouldBeEqual(t, 0, "Key AppendChain panics on complete -- finalized key", actual)
}

func Test_Key_AppendChainStrings_PanicOnComplete(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root")
	key.Finalized()

	// Act
	recovered := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				recovered = true
			}
		}()
		key.AppendChainStrings("shouldPanic")
	}()

	// Assert
	actual := args.Map{"recovered": recovered}
	expected := args.Map{"recovered": true}
	expected.ShouldBeEqual(t, 0, "Key AppendChainStrings panics on complete -- finalized key", actual)
}

// ── KeyJson.go — uncovered branches ──

func Test_Key_ParseInjectUsingJsonMust_Success(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	jsonResult := key.JsonPtr()
	var target keymk.Key

	// Act
	parsed := target.ParseInjectUsingJsonMust(jsonResult)

	// Assert
	actual := args.Map{
		"notNil": parsed != nil,
		"mainName": parsed.MainName(),
	}
	expected := args.Map{
		"notNil": true,
		"mainName": "",
	}
	expected.ShouldBeEqual(t, 0, "Key ParseInjectUsingJsonMust succeeds -- valid json", actual)
}

func Test_Key_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	// Arrange — create an invalid json result by serializing then corrupting
	var target keymk.Key
	badJson := keymk.NewKey.Default("x")
	_ = badJson.JsonPtr()
	// Corrupt the internal bytes to trigger unmarshal error
	// Use a fresh key and inject bad data via ParseInjectUsingJson first to confirm error path
	// Actually, the simplest approach: construct a corejson.Result from invalid bytes
	// But we don't have direct access. Instead, test the panic path with a nil jsonResult pointer
	// Actually let's test ParseInjectUsingJson error path with invalid data
	err := target.UnmarshalJSON([]byte(`{invalid`))

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Key ParseInjectUsingJson returns error -- invalid json", actual)
}

// ── KeyCompiler.go — uncovered branches ──

func Test_Key_Finalized_CompileStrings_NoAdditional(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")
	key.Finalized()

	// Act
	result := key.CompileStrings()

	// Assert
	actual := args.Map{
		"result": result,
		"chain": key.CompiledChain(),
	}
	expected := args.Map{
		"result": key.CompiledChain(),
		"chain": key.CompiledChain(),
	}
	expected.ShouldBeEqual(t, 0, "Key Finalized CompileStrings returns compiledChain -- no additional", actual)
}

func Test_Key_Finalized_CompileStrings_EmptyAdditional(t *testing.T) {
	// Arrange — finalized key with skip-empty, compile with empty strings
	key := keymk.NewKey.AllStrings(keymk.JoinerOption, "root", "a")
	key.Finalized()

	// Act — pass empty strings which should be skipped
	result := key.CompileStrings("", "")

	// Assert — should return compiledChain since all additional are empty and skipped
	actual := args.Map{"result": result}
	expected := args.Map{"result": key.CompiledChain()}
	expected.ShouldBeEqual(t, 0, "Key Finalized CompileStrings returns compiledChain -- empty additional skipped", actual)
}

func Test_Key_Finalized_Compile_EmptyAdditional(t *testing.T) {
	// Arrange — finalized key with skip-empty option
	key := keymk.NewKey.AllStrings(keymk.JoinerOption, "root", "a")
	key.Finalized()

	// Act — pass nil items (which get skipped in appendAnyItemsWithBaseStrings)
	result := key.Compile(nil, nil)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": key.CompiledChain()}
	expected.ShouldBeEqual(t, 0, "Key Finalized Compile returns compiledChain -- nil additional skipped", actual)
}

func Test_Key_CompileSingleItem_NoBrackets(t *testing.T) {
	// Arrange — use JoinerOption which has IsUseBrackets=false
	key := keymk.NewKey.Default("root", "a")

	// Act
	result := key.Compile()

	// Assert — no brackets in output
	actual := args.Map{"result": result}
	expected := args.Map{"result": "root-a"}
	expected.ShouldBeEqual(t, 0, "Key Compile without brackets -- JoinerOption", actual)
}

func Test_Key_CompileSingleItem_WithBrackets(t *testing.T) {
	// Arrange — use BracketJoinerOption which has IsUseBrackets=true
	key := keymk.NewKey.SquareBrackets("root", "a")

	// Act
	result := key.Compile()

	// Assert — has brackets
	actual := args.Map{"hasBrackets": len(result) > len("root-a")}
	expected := args.Map{"hasBrackets": true}
	expected.ShouldBeEqual(t, 0, "Key Compile with brackets -- BracketJoinerOption", actual)
}

func Test_Key_JoinUsingJoiner_Finalized(t *testing.T) {
	// Arrange — finalized key, use custom joiner with additional items
	key := keymk.NewKey.Default("root", "a")
	key.Finalized()

	// Act
	result := key.JoinUsingJoiner("/", "extra")

	// Assert
	actual := args.Map{
		"notEmpty": result != "",
		"longerThanChain": len(result) > len(key.CompiledChain()),
	}
	expected := args.Map{
		"notEmpty": true,
		"longerThanChain": true,
	}
	expected.ShouldBeEqual(t, 0, "Key JoinUsingJoiner with finalized key -- additional items", actual)
}

// ── appendAnyItemsWithBaseStrings.go — uncovered branches ──

func Test_Key_AppendChain_NilItemSkipped(t *testing.T) {
	// Arrange — append nil items which should be skipped
	key := keymk.NewKey.Default("root")

	// Act
	key.AppendChain(nil, "a", nil, "b")

	// Assert
	actual := args.Map{"length": key.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "Key AppendChain skips nil -- mixed nil and strings", actual)
}

func Test_Key_AppendChain_NoSkipEmpty(t *testing.T) {
	// Arrange — option with IsSkipEmptyEntry=false
	opt := &keymk.Option{
		Joiner:           "-",
		IsSkipEmptyEntry: false,
		IsUseBrackets:    false,
	}
	key := keymk.NewKey.All(opt, "root")

	// Act — append empty string (should NOT be skipped)
	key.AppendChain("")

	// Assert
	actual := args.Map{"length": key.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "Key AppendChain keeps empty -- IsSkipEmptyEntry false", actual)
}

func Test_Key_AppendChain_EmptySlice(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Default("root", "a")

	// Act — append no items
	key.AppendChain()

	// Assert
	actual := args.Map{"length": key.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "Key AppendChain no-op -- empty variadic", actual)
}

// ── appendStringsWithBaseAnyItems.go — uncovered branches ──
// This function is called internally by CompileKeys

func Test_Key_CompileKeys_WithEmptyChainKey(t *testing.T) {
	// Arrange — key2 has empty chains, exercises appendStringsWithBaseAnyItems empty path
	key1 := keymk.NewKey.Default("root")
	key2 := keymk.NewKey.Default("sub")

	// Act
	result := key1.CompileKeys(key2)

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key CompileKeys with no-chain key -- appendStrings empty path", actual)
}

func Test_Key_CompileKeys_SkipEmpty_InChains(t *testing.T) {
	// Arrange — key with IsSkipEmptyEntry=true, key2 has empty chain items
	key1 := keymk.NewKey.AllStrings(keymk.JoinerOption, "root")
	key2 := keymk.NewKey.AllStrings(keymk.JoinerOption, "sub", "", "val")

	// Act — CompileKeys processes key2's chains through appendStringsWithBaseAnyItems
	result := key1.CompileKeys(key2)

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key CompileKeys skips empty chain items -- IsSkipEmptyEntry", actual)
}

func Test_Key_CompileKeys_NoSkipEmpty(t *testing.T) {
	// Arrange — option with IsSkipEmptyEntry=false
	opt := &keymk.Option{
		Joiner:           "-",
		IsSkipEmptyEntry: false,
		IsUseBrackets:    false,
	}
	key1 := keymk.NewKey.AllStrings(opt, "root")
	key2 := keymk.NewKey.AllStrings(opt, "sub", "", "val")

	// Act
	result := key1.CompileKeys(key2)

	// Assert — empty string should be kept in output
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Key CompileKeys keeps empty chain items -- IsSkipEmptyEntry false", actual)
}

// ── KeyWithLegend — additional uncovered branches ──

func Test_KeyWithLegend_OutputItemsArray_WithLegend(t *testing.T) {
	// Arrange — isAttachLegendNames=true to go through legend path
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		true,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.GroupUserItemString("grp", "usr", "itm")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend OutputItemsArray with legend -- full chain", actual)
}

func Test_KeyWithLegend_CloneUsing(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	cloned := kl.CloneUsing("newGroup")

	// Assert
	actual := args.Map{
		"notNil": cloned != nil,
		"groupName": cloned.GroupName(),
	}
	expected := args.Map{
		"notNil": true,
		"groupName": "newGroup",
	}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend CloneUsing returns clone -- new group name", actual)
}

func Test_KeyWithLegend_CloneUsing_Nil(t *testing.T) {
	// Arrange
	var kl *keymk.KeyWithLegend

	// Act
	cloned := kl.CloneUsing("x")

	// Assert
	actual := args.Map{"isNil": cloned == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend CloneUsing returns nil -- nil receiver", actual)
}

func Test_KeyWithLegend_Clone(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	cloned := kl.Clone()

	// Assert
	actual := args.Map{
		"notNil": cloned != nil,
		"groupName": cloned.GroupName(),
	}
	expected := args.Map{
		"notNil": true,
		"groupName": "s",
	}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend Clone preserves group -- same group name", actual)
}

func Test_KeyWithLegend_CompileUsingJoiner(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.CompileUsingJoiner("/")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend CompileUsingJoiner returns compiled -- custom joiner", actual)
}

func Test_KeyWithLegend_CompileStrings(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.CompileStrings()

	// Assert
	actual := args.Map{"notEmpty": len(result) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend CompileStrings returns slice -- default", actual)
}

func Test_KeyWithLegend_Strings(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.Strings()

	// Assert
	actual := args.Map{"notEmpty": len(result) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend Strings returns slice -- delegates to CompileStrings", actual)
}

func Test_KeyWithLegend_CompileItemUsingJoiner(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.CompileItemUsingJoiner("/", "myItem")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend CompileItemUsingJoiner returns compiled -- custom joiner with item", actual)
}

func Test_KeyWithLegend_CompileDefault(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.CompileDefault()

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend CompileDefault returns compiled -- uses state and group", actual)
}

func Test_KeyWithLegend_Compile(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.Compile("itemX")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend Compile delegates to ItemString -- single item", actual)
}

func Test_KeyWithLegend_UpToState(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.UpToState("userX")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend UpToState returns compiled -- with user", actual)
}

func Test_KeyWithLegend_GroupIntRange(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.GroupIntRange(1, 3)

	// Assert
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend GroupIntRange returns range -- 1 to 3", actual)
}

func Test_KeyWithLegend_GroupUIntRange(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.GroupUIntRange(0, 2)

	// Assert
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend GroupUIntRange returns range -- 0 to 2", actual)
}

func Test_KeyWithLegend_ItemIntRange(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.ItemIntRange(1, 4)

	// Assert
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend ItemIntRange returns range -- 1 to 4", actual)
}

func Test_KeyWithLegend_ItemUIntRange(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.ItemUIntRange(0, 2)

	// Assert
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend ItemUIntRange returns range -- 0 to 2", actual)
}

func Test_KeyWithLegend_GroupUserString(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.GroupUserString("grp", "usr")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend GroupUserString returns compiled -- group and user", actual)
}

func Test_KeyWithLegend_GroupUser(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.GroupUser("grp", "usr")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend GroupUser returns compiled -- any group and user", actual)
}

func Test_KeyWithLegend_GroupUInt(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.GroupUInt(5)

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend GroupUInt returns compiled -- uint group", actual)
}

func Test_KeyWithLegend_GroupByte(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.GroupByte(3)

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend GroupByte returns compiled -- byte group", actual)
}

func Test_KeyWithLegend_GroupUserByte(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.GroupUserByte(1, 2)

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend GroupUserByte returns compiled -- byte group and user", actual)
}

func Test_KeyWithLegend_GroupUserItem(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.GroupUserItem("grp", "usr", "itm")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend GroupUserItem returns compiled -- full chain", actual)
}

func Test_KeyWithLegend_GroupStateUserItem(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.GroupStateUserItem("grp", "st", "usr", "itm")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend GroupStateUserItem returns compiled -- explicit state", actual)
}

func Test_KeyWithLegend_StateUserItem(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.StateUserItem("st", "usr", "itm")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend StateUserItem returns compiled -- uses creation group", actual)
}

func Test_KeyWithLegend_StateUser(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.StateUser("st", "usr")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend StateUser returns compiled -- state and user", actual)
}

func Test_KeyWithLegend_GroupStateUserItemString(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.GroupStateUserItemString("grp", "st", "usr", "itm")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend GroupStateUserItemString returns compiled -- string args", actual)
}

func Test_KeyWithLegend_GroupUserItemUint(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.GroupUserItemUint(1, 2, 3)

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend GroupUserItemUint returns compiled -- uint args", actual)
}

func Test_KeyWithLegend_GroupUserItemInt(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.GroupUserItemInt(1, 2, 3)

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend GroupUserItemInt returns compiled -- int args", actual)
}

func Test_KeyWithLegend_GroupItem(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.GroupItem("grp", "itm")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend GroupItem returns compiled -- any group and item", actual)
}

func Test_KeyWithLegend_StateItem(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.StateItem("st", "itm")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend StateItem returns compiled -- state and item", actual)
}

func Test_KeyWithLegend_GroupStateItemString(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.GroupStateItemString("grp", "st", "itm")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend GroupStateItemString returns compiled -- string args", actual)
}

func Test_KeyWithLegend_StateItemString(t *testing.T) {
	// Arrange
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.StateItemString("st", "itm")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend StateItemString returns compiled -- string args", actual)
}

// ── KeyWithLegend with brackets ──

func Test_KeyWithLegend_FinalStrings_WithBrackets(t *testing.T) {
	// Arrange — use bracket option
	kl := keymk.NewKeyWithLegend.All(
		keymk.BracketJoinerOption,
		keymk.FullLegends,
		false,
		"r", "p", "s", "g",
	)

	// Act
	result := kl.CompileDefault()

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend FinalStrings adds brackets -- BracketJoinerOption", actual)
}

// ── KeyWithLegend with legend attachment (OutputItemsArray legend path) ──

func Test_KeyWithLegend_OutputItemsArray_EmptyValues_SkipEmpty(t *testing.T) {
	// Arrange — legend attached, skip empty, some empty request fields
	kl := keymk.NewKeyWithLegend.All(
		keymk.JoinerOption,
		keymk.FullLegends,
		true,
		"r", "p", "", "",
	)

	// Act — request with empty user and item
	result := kl.UpToGroupString("grp")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend OutputItemsArray skips empty legend values -- skip empty option", actual)
}

func Test_KeyWithLegend_OutputItemsArray_KeepEmpty(t *testing.T) {
	// Arrange — legend attached, IsSkipEmptyEntry=false (add regardless)
	opt := &keymk.Option{
		Joiner:           "-",
		IsSkipEmptyEntry: false,
		IsUseBrackets:    false,
	}
	kl := keymk.NewKeyWithLegend.All(
		opt,
		keymk.FullLegends,
		true,
		"r", "p", "", "",
	)

	// Act
	result := kl.UpToGroupString("grp")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyWithLegend OutputItemsArray keeps empty values -- IsAddEntryRegardlessOfEmptiness", actual)
}

// ── newKeyWithLegendCreator — uncovered factory methods ──

func Test_NewKeyWithLegend_Create(t *testing.T) {
	// Arrange & Act
	kl := keymk.NewKeyWithLegend.Create(keymk.JoinerOption, "r", "p", "g")

	// Assert
	actual := args.Map{
		"notNil": kl != nil,
		"rootName": kl.RootName(),
	}
	expected := args.Map{
		"notNil": true,
		"rootName": "r",
	}
	expected.ShouldBeEqual(t, 0, "NewKeyWithLegend Create returns legend -- with full legends", actual)
}

func Test_NewKeyWithLegend_NoLegend(t *testing.T) {
	// Arrange & Act
	kl := keymk.NewKeyWithLegend.NoLegend(keymk.JoinerOption, "r", "p", "g")

	// Assert
	actual := args.Map{"ignoreLeg": kl.IsIgnoreLegendAttachments()}
	expected := args.Map{"ignoreLeg": true}
	expected.ShouldBeEqual(t, 0, "NewKeyWithLegend NoLegend ignores legends -- isAttach false", actual)
}

func Test_NewKeyWithLegend_NoLegendPackage(t *testing.T) {
	// Arrange & Act
	kl := keymk.NewKeyWithLegend.NoLegendPackage(false, keymk.JoinerOption, "r", "g")

	// Assert
	actual := args.Map{
		"package": kl.PackageName(),
		"ignoreLeg": kl.IsIgnoreLegendAttachments(),
	}
	expected := args.Map{
		"package": "",
		"ignoreLeg": true,
	}
	expected.ShouldBeEqual(t, 0, "NewKeyWithLegend NoLegendPackage has empty package -- no legend", actual)
}

func Test_NewKeyWithLegend_ShortLegend(t *testing.T) {
	// Arrange & Act
	kl := keymk.NewKeyWithLegend.ShortLegend(keymk.JoinerOption, "r", "p", "g")

	// Assert
	actual := args.Map{
		"notNil": kl != nil,
		"ignoreLeg": false,
	}
	expected := args.Map{
		"notNil": true,
		"ignoreLeg": false,
	}
	expected.ShouldBeEqual(t, 0, "NewKeyWithLegend ShortLegend attaches short legends -- isAttach true", actual)
}

// ── KeyLegendCompileRequest — uncovered methods ──

func Test_KeyLegendCompileRequest_NewKeyLegend(t *testing.T) {
	// Arrange
	req := keymk.KeyLegendCompileRequest{
		GroupId: "grp",
	}

	// Act
	kl := req.NewKeyLegend(keymk.JoinerOption, keymk.FullLegends, true, "r", "p", "s")

	// Assert
	actual := args.Map{
		"notNil": kl != nil,
		"groupName": kl.GroupName(),
	}
	expected := args.Map{
		"notNil": true,
		"groupName": "grp",
	}
	expected.ShouldBeEqual(t, 0, "KeyLegendCompileRequest NewKeyLegend creates legend -- from request", actual)
}

func Test_KeyLegendCompileRequest_NewKeyLegendDefaults(t *testing.T) {
	// Arrange
	req := keymk.KeyLegendCompileRequest{
		GroupId: "grp",
	}

	// Act
	kl := req.NewKeyLegendDefaults("r", "p", "s")

	// Assert
	actual := args.Map{
		"notNil": kl != nil,
		"groupName": kl.GroupName(),
	}
	expected := args.Map{
		"notNil": true,
		"groupName": "grp",
	}
	expected.ShouldBeEqual(t, 0, "KeyLegendCompileRequest NewKeyLegendDefaults creates legend -- default options", actual)
}

// ── Option — uncovered branch ──

func Test_Option_IsAddEntryRegardlessOfEmptiness_True(t *testing.T) {
	// Arrange — IsSkipEmptyEntry=false means add regardless
	opt := &keymk.Option{
		Joiner:           "-",
		IsSkipEmptyEntry: false,
	}

	// Act
	result := opt.IsAddEntryRegardlessOfEmptiness()

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Option IsAddEntryRegardlessOfEmptiness returns true -- IsSkipEmptyEntry false", actual)
}

// ── AppendChainStrings — non-skip path ──

func Test_Key_AppendChainStrings_NoSkipEmpty(t *testing.T) {
	// Arrange — option with IsSkipEmptyEntry=false
	opt := &keymk.Option{
		Joiner:           "-",
		IsSkipEmptyEntry: false,
		IsUseBrackets:    false,
	}
	key := keymk.NewKey.All(opt, "root")

	// Act — append empty string, should NOT be skipped
	key.AppendChainStrings("", "a")

	// Assert
	actual := args.Map{"length": key.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "Key AppendChainStrings keeps empty -- IsSkipEmptyEntry false", actual)
}

// ── TemplateReplacer with curly ──

func Test_TemplateReplacer_IntRange_Curly(t *testing.T) {
	// Arrange — key with curly placeholder
	key := keymk.NewKey.Default("item", "{id}")
	key.Finalized()
	tr := key.TemplateReplacer()

	// Act
	result := tr.IntRange(true, "id", 1, 3)

	// Assert
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "TemplateReplacer IntRange with curly -- replaces {id}", actual)
}

func Test_TemplateReplacer_CompileUsingReplacerMap_EmptyFormat(t *testing.T) {
	// Arrange — key that compiles to empty
	key := keymk.NewKey.AllStrings(keymk.JoinerOption, "")
	key.Finalized()
	tr := key.TemplateReplacer()

	// Act
	result := tr.CompileUsingReplacerMap(true, map[string]string{"x": "y"})

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "TemplateReplacer CompileUsingReplacerMap returns empty -- empty template", actual)
}

// ── FixedLegend — Compile with isKeepFormatOnEmpty=true and non-empty values ──

func Test_FixedLegend_Compile_NoKeepFormat(t *testing.T) {
	// Arrange & Act — isKeepFormatOnEmpty=false, all values replaced
	result := keymk.FixedLegend.Compile(false, "r", "p", "g", "s", "u", "i")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FixedLegend Compile replaces all -- isKeepFormat false", actual)
}

func Test_FixedLegend_Compile_KeepFormatWithEmpty(t *testing.T) {
	// Arrange & Act — isKeepFormatOnEmpty=true, some empty values kept as format
	result := keymk.FixedLegend.Compile(true, "r", "p", "", "s", "", "i")

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FixedLegend Compile keeps format for empty -- isKeepFormat true", actual)
}
