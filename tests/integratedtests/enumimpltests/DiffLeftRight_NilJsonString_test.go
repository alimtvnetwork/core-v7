package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── DiffLeftRight ──

func Test_DiffLeftRight_NilJsonString(t *testing.T) {
	// Arrange
	var dlr *enumimpl.DiffLeftRight

	// Act
	actual := args.Map{"result": dlr.JsonString()}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "JsonString returns empty -- nil receiver", actual)
}

func Test_DiffLeftRight_SameValues(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "abc", Right: "abc"}

	// Act
	actual := args.Map{
		"isSame":           dlr.IsSame(),
		"isSameType":       dlr.IsSameTypeSame(),
		"isEqual":          dlr.IsEqual(false),
		"isEqualRegardless": dlr.IsEqual(true),
		"isNotEqual":       dlr.IsNotEqual(),
		"diffString":       dlr.DiffString(),
	}

	// Assert
	expected := args.Map{
		"isSame":           true,
		"isSameType":       true,
		"isEqual":          true,
		"isEqualRegardless": true,
		"isNotEqual":       false,
		"diffString":       "",
	}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight all same -- equal strings", actual)
}

func Test_DiffLeftRight_DifferentValues(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "abc", Right: "xyz"}

	// Act
	actual := args.Map{
		"isSame":       dlr.IsSame(),
		"isNotEqual":   dlr.IsNotEqual(),
		"hasMismatch":  dlr.HasMismatch(false),
		"hasMismatchR": dlr.HasMismatch(true),
		"hasMismatchRegardless": dlr.HasMismatchRegardlessOfType(),
		"diffNotEmpty": dlr.DiffString() != "",
	}

	// Assert
	expected := args.Map{
		"isSame":       false,
		"isNotEqual":   true,
		"hasMismatch":  true,
		"hasMismatchR": true,
		"hasMismatchRegardless": true,
		"diffNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight all different -- different strings", actual)
}

func Test_DiffLeftRight_DifferentTypes(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: 1, Right: "1"}

	// Act
	actual := args.Map{
		"sameTypeSame":       dlr.IsSameTypeSame(),
		"sameRegardless":     dlr.IsSameRegardlessOfType(),
	}

	// Assert
	expected := args.Map{
		"sameTypeSame":       false,
		"sameRegardless":     true,
	}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight different types same value -- int vs string", actual)
}

func Test_DiffLeftRight_SpecificFullString_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	l, r := dlr.SpecificFullString()

	// Act
	actual := args.Map{
		"lNotEmpty": l != "",
		"rNotEmpty": r != "",
	}

	// Assert
	expected := args.Map{
		"lNotEmpty": true,
		"rNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "SpecificFullString returns non-empty -- both sides", actual)
}

func Test_DiffLeftRight_Types(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: 1, Right: "x"}
	l, r := dlr.Types()

	// Act
	actual := args.Map{
		"lNotNil": l != nil,
		"rNotNil": r != nil,
		"different": l != r,
	}

	// Assert
	expected := args.Map{
		"lNotNil": true,
		"rNotNil": true,
		"different": true,
	}
	expected.ShouldBeEqual(t, 0, "Types returns reflect types -- different types", actual)
}

func Test_DiffLeftRight_StringNonNil(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}

	// Act
	actual := args.Map{"notEmpty": dlr.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns json -- non-nil", actual)
}

// ── DynamicMap — ConcatNew ──

func Test_DynamicMap_ConcatNew_OverrideExisting(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	other := enumimpl.DynamicMap{"b": 99, "c": 3}
	result := dm.ConcatNew(true, other)

	// Act
	actual := args.Map{
		"b": result["b"],
		"c": result["c"],
		"a": result["a"],
	}

	// Assert
	expected := args.Map{
		"b": 99,
		"c": 3,
		"a": 1,
	}
	expected.ShouldBeEqual(t, 0, "ConcatNew overrides existing -- b becomes 99", actual)
}

func Test_DynamicMap_ConcatNew_NoOverride(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	other := enumimpl.DynamicMap{"b": 99, "c": 3}
	result := dm.ConcatNew(false, other)

	// Act
	actual := args.Map{
		"b": result["b"],
		"c": result["c"],
	}

	// Assert
	expected := args.Map{
		"b": 2,
		"c": 3,
	}
	expected.ShouldBeEqual(t, 0, "ConcatNew no override -- b stays 2", actual)
}

func Test_DynamicMap_ConcatNew_BothEmpty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}
	other := enumimpl.DynamicMap{}
	result := dm.ConcatNew(true, other)

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ConcatNew both empty -- zero length", actual)
}

// ── DynamicMap — StringsUsingFmt ──

func Test_DynamicMap_StringsUsingFmt_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	result := dm.StringsUsingFmt(func(index int, key string, val any) string {
		return key + "=" + enumimpl.NameWithValue(val)
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringsUsingFmt returns formatted -- two items", actual)
}

func Test_DynamicMap_StringsUsingFmt_Empty(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}
	result := dm.StringsUsingFmt(func(index int, key string, val any) string {
		return ""
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringsUsingFmt returns empty -- empty map", actual)
}

// ── DynamicMap — Serialize ──

func Test_DynamicMap_Serialize_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	b, err := dm.Serialize()

	// Act
	actual := args.Map{
		"hasBytes": len(b) > 0,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize returns json bytes -- single item", actual)
}

// ── DynamicMap — IsStringEqual ──

func Test_DynamicMap_IsStringEqual_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{
		"selfEqual":    dm.IsStringEqual(dm.String()),
		"notEqual":     dm.IsStringEqual("nope"),
	}

	// Assert
	expected := args.Map{
		"selfEqual":    true,
		"notEqual":     false,
	}
	expected.ShouldBeEqual(t, 0, "IsStringEqual matches own string -- true/false", actual)
}

// ── DynamicMap — IsKeysEqualOnly ──

func Test_DynamicMap_IsKeysEqualOnly_Match(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	right := map[string]any{"a": 99, "b": 88}

	// Act
	actual := args.Map{"equal": dm.IsKeysEqualOnly(right)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsKeysEqualOnly true -- same keys different values", actual)
}

func Test_DynamicMap_IsKeysEqualOnly_Mismatch(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 99, "b": 88}

	// Act
	actual := args.Map{"equal": dm.IsKeysEqualOnly(right)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsKeysEqualOnly false -- different key count", actual)
}

func Test_DynamicMap_IsKeysEqualOnly_BothNil(t *testing.T) {
	// Arrange
	var dm *enumimpl.DynamicMap

	// Act
	actual := args.Map{"equal": dm.IsKeysEqualOnly(nil)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsKeysEqualOnly true -- both nil", actual)
}

func Test_DynamicMap_IsKeysEqualOnly_OneNil(t *testing.T) {
	// Arrange
	var dm *enumimpl.DynamicMap
	right := map[string]any{"a": 1}

	// Act
	actual := args.Map{"equal": dm.IsKeysEqualOnly(right)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsKeysEqualOnly false -- left nil right non-nil", actual)
}

func Test_DynamicMap_IsKeysEqualOnly_MissingKey(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "c": 3}
	right := map[string]any{"a": 99, "b": 88}

	// Act
	actual := args.Map{"equal": dm.IsKeysEqualOnly(right)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsKeysEqualOnly false -- key mismatch", actual)
}

// ── DynamicMap — KeyValue / KeyValueString ──

func Test_DynamicMap_KeyValue_Found(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"x": 42}
	val, found := dm.KeyValue("x")

	// Act
	actual := args.Map{
		"val": val,
		"found": found,
	}

	// Assert
	expected := args.Map{
		"val": 42,
		"found": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValue found -- existing key", actual)
}

func Test_DynamicMap_KeyValue_NotFound(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"x": 42}
	_, found := dm.KeyValue("z")

	// Act
	actual := args.Map{"found": found}

	// Assert
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "KeyValue not found -- missing key", actual)
}

func Test_DynamicMap_KeyValueString_NotFound(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"x": 42}
	val, found := dm.KeyValueString("z")

	// Act
	actual := args.Map{
		"val": val,
		"found": found,
	}

	// Assert
	expected := args.Map{
		"val": "",
		"found": false,
	}
	expected.ShouldBeEqual(t, 0, "KeyValueString returns empty -- missing key", actual)
}

// ── DynamicMap — Add ──

func Test_DynamicMap_Add_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	dm.Add("b", 2)

	// Act
	actual := args.Map{"len": dm.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Add adds key -- length increases", actual)
}

// ── DynamicMap — HasAllKeys / HasAnyKeys ──

func Test_DynamicMap_HasAllKeys_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2, "c": 3}

	// Act
	actual := args.Map{
		"allPresent": dm.HasAllKeys("a", "b"),
		"oneMissing": dm.HasAllKeys("a", "z"),
	}

	// Assert
	expected := args.Map{
		"allPresent": true,
		"oneMissing": false,
	}
	expected.ShouldBeEqual(t, 0, "HasAllKeys correct -- present and missing", actual)
}

func Test_DynamicMap_HasAnyKeys_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	// Act
	actual := args.Map{
		"onePresent": dm.HasAnyKeys("z", "a"),
		"nonPresent": dm.HasAnyKeys("x", "y"),
	}

	// Assert
	expected := args.Map{
		"onePresent": true,
		"nonPresent": false,
	}
	expected.ShouldBeEqual(t, 0, "HasAnyKeys correct -- found and not found", actual)
}

// ── DynamicMap — Set / AddNewOnly ──

func Test_DynamicMap_AddOrUpdate(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	isNew := dm.AddOrUpdate("b", 2)
	isUpdate := dm.AddOrUpdate("a", 99)

	// Act
	actual := args.Map{
		"isNew": isNew,
		"isUpdate": isUpdate,
		"aVal": dm["a"],
	}

	// Assert
	expected := args.Map{
		"isNew": true,
		"isUpdate": false,
		"aVal": 99,
	}
	expected.ShouldBeEqual(t, 0, "AddOrUpdate new then update -- correct flags", actual)
}

// ── DynamicMap — Diff with LeftRightDiffCheckerImpl ──

func Test_DynamicMap_DiffJsonMessageLeftRight_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	right := map[string]any{"a": 1, "b": 2}
	result := dm.DiffJsonMessageLeftRight(false, right)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "DiffJsonMessageLeftRight empty -- no diff", actual)
}

func Test_DynamicMap_DiffJsonMessageLeftRight_HasDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	right := map[string]any{"a": 1, "c": 3}
	result := dm.DiffJsonMessageLeftRight(false, right)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DiffJsonMessageLeftRight non-empty -- has diff", actual)
}

func Test_DynamicMap_LogShouldDiffLeftRightMessage_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 1}
	result := dm.LogShouldDiffLeftRightMessage(false, "test", right)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "LogShouldDiffLeftRightMessage empty -- no diff", actual)
}

// ── DynamicMap — ExpectingMessage ──

func Test_DynamicMap_ExpectingMessage_Equal(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	result := dm.ExpectingMessage("test", map[string]any{"a": 1})

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingMessage empty -- maps equal", actual)
}

func Test_DynamicMap_ExpectingMessage_NotEqual(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	result := dm.ExpectingMessage("test", map[string]any{"a": 2})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingMessage non-empty -- maps differ", actual)
}

func Test_DynamicMap_LogExpectingMessage_Equal(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	dm.LogExpectingMessage("test", map[string]any{"a": 1})
	// no panic = pass

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LogExpectingMessage no panic -- equal maps", actual)
}

// ── DynamicMap — IsMismatch / IsRawMismatch ──

func Test_DynamicMap_IsMismatch(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	right := enumimpl.DynamicMap{"a": 2}

	// Act
	actual := args.Map{"mismatch": dm.IsMismatch(false, &right)}

	// Assert
	expected := args.Map{"mismatch": true}
	expected.ShouldBeEqual(t, 0, "IsMismatch true -- values differ", actual)
}

func Test_DynamicMap_IsRawMismatch_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 2}

	// Act
	actual := args.Map{"mismatch": dm.IsRawMismatch(false, right)}

	// Assert
	expected := args.Map{"mismatch": true}
	expected.ShouldBeEqual(t, 0, "IsRawMismatch true -- values differ", actual)
}

// ── DynamicMap — IsEqual edge cases ──

func Test_DynamicMap_IsEqual_SamePointer_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"equal": dm.IsEqual(false, &dm)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqual true -- same pointer", actual)
}

func Test_DynamicMap_IsEqual_Regardless(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	right := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"equal": dm.IsEqual(true, &right)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqual regardless true -- same values", actual)
}

func Test_DynamicMap_IsRawEqual_MissingKey_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	right := map[string]any{"a": 1, "c": 3}

	// Act
	actual := args.Map{"equal": dm.IsRawEqual(false, right)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsRawEqual false -- key mismatch", actual)
}

// ── DynamicMap — DiffRaw nil cases ──

func Test_DynamicMap_DiffRawUsingDifferChecker_LeftNil(t *testing.T) {
	// Arrange
	var dm *enumimpl.DynamicMap
	right := map[string]any{"a": 1}
	result := dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, right)

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DiffRaw returns right -- left nil", actual)
}

func Test_DynamicMap_DiffRawUsingDifferChecker_RightNil(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	result := dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DiffRaw returns left -- right nil", actual)
}

func Test_DynamicMap_DiffRawLeftRight_BothNil(t *testing.T) {
	// Arrange
	var dm *enumimpl.DynamicMap
	l, r := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)

	// Act
	actual := args.Map{
		"lLen": l.Length(),
		"rLen": r.Length(),
	}

	// Assert
	expected := args.Map{
		"lLen": 0,
		"rLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "DiffRawLeftRight both empty -- both nil", actual)
}

func Test_DynamicMap_DiffRawLeftRight_LeftNil(t *testing.T) {
	// Arrange
	var dm *enumimpl.DynamicMap
	right := map[string]any{"a": 1}
	l, r := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, right)

	// Act
	actual := args.Map{
		"lLen": l.Length(),
		"rLen": r.Length(),
	}

	// Assert
	expected := args.Map{
		"lLen": 1,
		"rLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "DiffRawLeftRight returns right as lDiff -- left nil", actual)
}

func Test_DynamicMap_DiffRawLeftRight_RightNil(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	l, r := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)

	// Act
	actual := args.Map{
		"lLen": l.Length(),
		"rLen": r.Length(),
	}

	// Assert
	expected := args.Map{
		"lLen": 1,
		"rLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "DiffRawLeftRight returns left as lDiff -- right nil", actual)
}

// ── DynamicMap — ShouldDiff messages ──

func Test_DynamicMap_ShouldDiffMessage_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 1}
	result := dm.ShouldDiffMessage(false, "title", right)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ShouldDiffMessage empty -- no diff", actual)
}

func Test_DynamicMap_ShouldDiffMessage_HasDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 2}
	result := dm.ShouldDiffMessage(false, "title", right)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldDiffMessage non-empty -- has diff", actual)
}

func Test_DynamicMap_ShouldDiffLeftRightMessage_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 1}
	result := dm.ShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.LeftRightDiffCheckerImpl, false, "title", right,
	)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ShouldDiffLeftRightMessage empty -- no diff", actual)
}

func Test_DynamicMap_ShouldDiffLeftRightMessage_HasDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"b": 2}
	result := dm.ShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.LeftRightDiffCheckerImpl, false, "title", right,
	)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldDiffLeftRightMessage non-empty -- has diff", actual)
}

// ── DynamicMap — LogShouldDiffMessageUsingDifferChecker ──

func Test_DynamicMap_LogShouldDiffMessageUsingDifferChecker_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 1}
	result := dm.LogShouldDiffMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "title", right,
	)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "LogShouldDiffMessage empty -- no diff", actual)
}

// ── DynamicMap — DiffJsonMessage ──

func Test_DynamicMap_DiffJsonMessage_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 1}
	result := dm.DiffJsonMessage(false, right)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "DiffJsonMessage empty -- no diff", actual)
}

func Test_DynamicMap_DiffJsonMessageUsingDifferChecker_NoDiff(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 1}
	result := dm.DiffJsonMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, right,
	)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "DiffJsonMessageUsingDifferChecker empty -- no diff", actual)
}

// ── DynamicMap — AllValuesIntegers / AllValuesStrings ──

func Test_DynamicMap_AllValuesIntegers_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	result := dm.AllValuesIntegers()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AllValuesIntegers returns integers -- two items", actual)
}

func Test_DynamicMap_AllValuesStringsSorted(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"b": 2, "a": 1}
	result := dm.AllValuesStringsSorted()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AllValuesStringsSorted returns sorted -- two items", actual)
}

// ── DynamicMap — HasIndex / LastIndex / Count ──

func Test_DynamicMap_HasIndex(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	// Act
	actual := args.Map{
		"hasIdx0": dm.HasIndex(0),
		"hasIdx1": dm.HasIndex(1),
		"hasIdx5": dm.HasIndex(5),
	}

	// Assert
	expected := args.Map{
		"hasIdx0": true,
		"hasIdx1": true,
		"hasIdx5": false,
	}
	expected.ShouldBeEqual(t, 0, "HasIndex correct -- two items", actual)
}

func Test_DynamicMap_Count(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"count": dm.Count()}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Count returns 1 -- single item", actual)
}

// ── DynamicMap — SortedKeyValues with string values ──

func Test_DynamicMap_SortedKeyAnyValues_StringValues_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"b": "beta", "a": "alpha"}
	result := dm.SortedKeyAnyValues()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SortedKeyAnyValues returns sorted -- string values", actual)
}

// ── DynamicMap — MapIntegerString with string values ──

func Test_DynamicMap_MapIntegerString_StringValues(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": "alpha", "b": "beta"}
	rangeMap, sortedKeys := dm.MapIntegerString()

	// Act
	actual := args.Map{
		"mapLen": len(rangeMap),
		"keysLen": len(sortedKeys),
	}

	// Assert
	expected := args.Map{
		"mapLen": 1,
		"keysLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "MapIntegerString handles string values -- two items", actual)
}

// ── DynamicMap — ConvMap variants ──

func Test_DynamicMap_ConvMapStringString_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": "alpha", "b": "beta"}
	result := dm.ConvMapStringString()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ConvMapStringString returns map -- two items", actual)
}

func Test_DynamicMap_ConvMapInt64String(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	result := dm.ConvMapInt64String()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ConvMapInt64String returns map -- two items", actual)
}

func Test_DynamicMap_ConvMapStringInteger(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	result := dm.ConvMapStringInteger()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ConvMapStringInteger returns map -- two items", actual)
}

// ── DynamicMap — KeyValueByte edge cases ──

func Test_DynamicMap_KeyValueByte_NotFound(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	_, isFound, _ := dm.KeyValueByte("z")

	// Act
	actual := args.Map{"found": isFound}

	// Assert
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "KeyValueByte not found -- missing key", actual)
}

func Test_DynamicMap_KeyValueByte_DirectByte(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": byte(42)}
	val, isFound, isFailed := dm.KeyValueByte("a")

	// Act
	actual := args.Map{
		"val": val,
		"found": isFound,
		"failed": isFailed,
	}

	// Assert
	expected := args.Map{
		"val": byte(42),
		"found": true,
		"failed": false,
	}
	expected.ShouldBeEqual(t, 0, "KeyValueByte returns byte -- direct byte value", actual)
}

// ── DynamicMap — KeyValueInt edge cases ──

func Test_DynamicMap_KeyValueInt_DirectInt(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 42}
	val, isFound, isFailed := dm.KeyValueInt("a")

	// Act
	actual := args.Map{
		"val": val,
		"found": isFound,
		"failed": isFailed,
	}

	// Assert
	expected := args.Map{
		"val": 42,
		"found": true,
		"failed": false,
	}
	expected.ShouldBeEqual(t, 0, "KeyValueInt returns int -- direct int value", actual)
}

func Test_DynamicMap_KeyValueInt_NotFound(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	_, isFound, isFailed := dm.KeyValueInt("z")

	// Act
	actual := args.Map{
		"found": isFound,
		"failed": isFailed,
	}

	// Assert
	expected := args.Map{
		"found": false,
		"failed": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValueInt not found -- missing key", actual)
}

func Test_DynamicMap_KeyValueInt_DirectByte(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": byte(5)}
	val, isFound, isFailed := dm.KeyValueInt("a")

	// Act
	actual := args.Map{
		"val": val,
		"found": isFound,
		"failed": isFailed,
	}

	// Assert
	expected := args.Map{
		"val": 5,
		"found": true,
		"failed": false,
	}
	expected.ShouldBeEqual(t, 0, "KeyValueInt converts byte -- direct byte value", actual)
}

// ── DynamicMap — IsValueTypeOf ──

func Test_DynamicMap_IsValueTypeOf_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": "hello"}

	// Act
	actual := args.Map{"isString": dm.IsValueString()}

	// Assert
	expected := args.Map{"isString": true}
	expected.ShouldBeEqual(t, 0, "IsValueString true -- string value", actual)
}

// ── Format / FormatUsingFmt ──

func Test_Format(t *testing.T) {
	// Arrange
	result := enumimpl.Format("MyType", "MyName", "42", "Enum of {type-name} - {name} - {value}")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "Enum of MyType - MyName - 42"}
	expected.ShouldBeEqual(t, 0, "Format compiles template -- all keys replaced", actual)
}

// ── KeyAnyVal ──

func Test_KeyAnyVal_StringMethods(t *testing.T) {
	// Arrange
	kav := enumimpl.KeyAnyVal{Key: "Name", AnyValue: 5}

	// Act
	actual := args.Map{
		"keyString":   kav.KeyString(),
		"anyVal":      kav.AnyVal(),
		"anyValStr":   kav.AnyValString() != "",
		"wrapKey":     kav.WrapKey() != "",
		"wrapValue":   kav.WrapValue() != "",
		"isString":    kav.IsString(),
		"valInt":      kav.ValInt(),
		"stringOut":   kav.String() != "",
	}

	// Assert
	expected := args.Map{
		"keyString":   "Name",
		"anyVal":      5,
		"anyValStr":   true,
		"wrapKey":     true,
		"wrapValue":   true,
		"isString":    false,
		"valInt":      5,
		"stringOut":   true,
	}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal methods -- integer value", actual)
}

func Test_KeyAnyVal_IsString_StringValue(t *testing.T) {
	// Arrange
	kav := enumimpl.KeyAnyVal{Key: "Name", AnyValue: "hello"}

	// Act
	actual := args.Map{
		"isString": kav.IsString(),
		"str": kav.String() != "",
	}

	// Assert
	expected := args.Map{
		"isString": true,
		"str": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal IsString true -- string value", actual)
}

func Test_KeyAnyVal_KeyValInteger_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	kav := enumimpl.KeyAnyVal{Key: "Name", AnyValue: 5}
	kvi := kav.KeyValInteger()

	// Act
	actual := args.Map{
		"key": kvi.Key,
		"val": kvi.ValueInteger,
	}

	// Assert
	expected := args.Map{
		"key": "Name",
		"val": 5,
	}
	expected.ShouldBeEqual(t, 0, "KeyValInteger converts correctly -- int value", actual)
}

// ── KeyValInteger ──

func Test_KeyValInteger_Methods_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	kvi := enumimpl.KeyValInteger{Key: "Name", ValueInteger: 5}
	kav := kvi.KeyAnyVal()

	// Act
	actual := args.Map{
		"wrapKey":    kvi.WrapKey() != "",
		"wrapValue":  kvi.WrapValue() != "",
		"isString":   kvi.IsString(),
		"stringOut":  kvi.String() != "",
		"kavKey":     kav.Key,
	}

	// Assert
	expected := args.Map{
		"wrapKey":    true,
		"wrapValue":  true,
		"isString":   false,
		"stringOut":  true,
		"kavKey":     "Name",
	}
	expected.ShouldBeEqual(t, 0, "KeyValInteger methods -- integer value", actual)
}

// ── KeyAnyValues func ──

func Test_KeyAnyValues_Empty_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	result := enumimpl.KeyAnyValues([]string{}, []byte{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues returns empty -- empty input", actual)
}

func Test_KeyAnyValues_NonEmpty(t *testing.T) {
	// Arrange
	result := enumimpl.KeyAnyValues([]string{"A", "B"}, []byte{0, 1})

	// Act
	actual := args.Map{
		"len": len(result),
		"firstKey": result[0].Key,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"firstKey": "A",
	}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues returns items -- two entries", actual)
}

// ── AllNameValues func ──

func Test_AllNameValues_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	result := enumimpl.AllNameValues([]string{"A", "B"}, []byte{0, 1})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AllNameValues returns strings -- two entries", actual)
}

// ── IntegersRangesOfAnyVal ──

func Test_IntegersRangesOfAnyVal_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	result := enumimpl.IntegersRangesOfAnyVal([]byte{2, 0, 1})

	// Act
	actual := args.Map{
		"first": result[0],
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"first": 0,
		"last": 2,
	}
	expected.ShouldBeEqual(t, 0, "IntegersRangesOfAnyVal sorted -- byte input", actual)
}

// ── PrependJoin / JoinPrependUsingDot ──

func Test_PrependJoin_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	result := enumimpl.PrependJoin(".", "prefix", "a", "b")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "prefix.a.b"}
	expected.ShouldBeEqual(t, 0, "PrependJoin joins with dot -- three parts", actual)
}

func Test_JoinPrependUsingDot_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	result := enumimpl.JoinPrependUsingDot("prefix", "a", "b")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "prefix.a.b"}
	expected.ShouldBeEqual(t, 0, "JoinPrependUsingDot joins -- three parts", actual)
}

// ── NameWithValue func ──

func Test_NameWithValue_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	result := enumimpl.NameWithValue(42)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameWithValue returns formatted -- integer", actual)
}

// ── OnlySupportedErr / UnsupportedNames ──

func Test_OnlySupportedErr_NoUnsupported(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(1, []string{"A", "B"}, "A", "B")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr nil -- all supported", actual)
}

func Test_OnlySupportedErr_HasUnsupported_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(1, []string{"A", "B", "C"}, "A")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr error -- B and C unsupported", actual)
}

func Test_OnlySupportedErr_EmptyAllNames(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(1, []string{}, "A")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr nil -- empty allNames", actual)
}

func Test_UnsupportedNames_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	result := enumimpl.UnsupportedNames([]string{"A", "B", "C"}, "A")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "UnsupportedNames returns B and C -- A supported", actual)
}

// ── BasicString — additional coverage ──

func Test_BasicString_GetNameByIndex_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"Alpha", "Beta", "Gamma"})

	// Act
	actual := args.Map{
		"idx1":      bs.GetNameByIndex(1),
		"idx0":      bs.GetNameByIndex(0),
		"idxNeg":    bs.GetNameByIndex(-1),
		"idxTooBig": bs.GetNameByIndex(99),
	}

	// Assert
	expected := args.Map{
		"idx1":      "Beta",
		"idx0":      "",
		"idxNeg":    "",
		"idxTooBig": "",
	}
	expected.ShouldBeEqual(t, 0, "GetNameByIndex returns name or empty -- boundary checks", actual)
}

func Test_BasicString_GetIndexByName_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"Alpha", "Beta"})

	// Act
	actual := args.Map{
		"alpha":   bs.GetIndexByName("Alpha"),
		"unknown": bs.GetIndexByName("Unknown"),
		"empty":   bs.GetIndexByName(""),
	}

	// Assert
	expected := args.Map{
		"alpha":   0,
		"unknown": -1,
		"empty":   -1,
	}
	expected.ShouldBeEqual(t, 0, "GetIndexByName returns index or -1 -- various inputs", actual)
}

func Test_BasicString_RangesIntegers(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B", "C"})
	result := bs.RangesIntegers()

	// Act
	actual := args.Map{
		"len": len(result),
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"last": 2,
	}
	expected.ShouldBeEqual(t, 0, "RangesIntegers returns 0-based -- three items", actual)
}

func Test_BasicString_Hashset(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})

	// Act
	actual := args.Map{
		"hasItems": len(bs.Hashset()) > 0,
		"ptrNotNil": bs.HashsetPtr() != nil,
	}

	// Assert
	expected := args.Map{
		"hasItems": true,
		"ptrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Hashset and HashsetPtr non-empty -- two items", actual)
}

func Test_BasicString_IsAnyOf_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})

	// Act
	actual := args.Map{
		"emptyCheck": bs.IsAnyOf("A"),
		"found":      bs.IsAnyOf("A", "A", "B"),
		"notFound":   bs.IsAnyOf("A", "X", "Y"),
	}

	// Assert
	expected := args.Map{
		"emptyCheck": true,
		"found":      true,
		"notFound":   false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyOf correct -- various inputs", actual)
}

func Test_BasicString_IsAnyNamesOf_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})

	// Act
	actual := args.Map{
		"found":    bs.IsAnyNamesOf("A", "A"),
		"notFound": bs.IsAnyNamesOf("A", "X"),
	}

	// Assert
	expected := args.Map{
		"found":    true,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyNamesOf correct -- found and not found", actual)
}

func Test_BasicString_HasAnyItem(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A"})

	// Act
	actual := args.Map{"hasAny": bs.HasAnyItem()}

	// Assert
	expected := args.Map{"hasAny": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem true -- one item", actual)
}

func Test_BasicString_MaxIndex(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B", "C"})

	// Act
	actual := args.Map{"maxIdx": bs.MaxIndex()}

	// Assert
	expected := args.Map{"maxIdx": 2}
	expected.ShouldBeEqual(t, 0, "MaxIndex returns 2 -- three items", actual)
}

func Test_BasicString_NameWithIndexMap(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	m := bs.NameWithIndexMap()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "NameWithIndexMap returns map -- two items", actual)
}

func Test_BasicString_ToEnumJsonBytes_Found(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	b, err := bs.ToEnumJsonBytes("A")

	// Act
	actual := args.Map{
		"hasBytes": len(b) > 0,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToEnumJsonBytes found -- valid value", actual)
}

func Test_BasicString_ToEnumJsonBytes_NotFound_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	_, err := bs.ToEnumJsonBytes("UNKNOWN")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToEnumJsonBytes error -- unknown value", actual)
}

func Test_BasicString_UnmarshallToValue_NilNotMapped_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	_, err := bs.UnmarshallToValue(false, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue error -- nil not mapped", actual)
}

func Test_BasicString_UnmarshallToValue_NilMapped_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	val, err := bs.UnmarshallToValue(true, nil)

	// Act
	actual := args.Map{
		"val": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": "A",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns min -- nil mapped", actual)
}

func Test_BasicString_UnmarshallToValue_EmptyMapped_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	val, err := bs.UnmarshallToValue(true, []byte(`""`))

	// Act
	actual := args.Map{
		"val": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": "A",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns min -- empty mapped", actual)
}

func Test_BasicString_EnumType_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A"})

	// Act
	actual := args.Map{"isString": bs.EnumType() != 0}

	// Assert
	expected := args.Map{"isString": true}
	expected.ShouldBeEqual(t, 0, "EnumType returns non-zero -- String type", actual)
}

func Test_BasicString_IsValidRange(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})

	// Act
	actual := args.Map{
		"valid":   bs.IsValidRange("A"),
		"invalid": bs.IsValidRange("UNKNOWN"),
	}

	// Assert
	expected := args.Map{
		"valid":   true,
		"invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "IsValidRange correct -- valid and invalid", actual)
}

func Test_BasicString_AppendPrependJoinValue_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	result := bs.AppendPrependJoinValue(".", "B", "A")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AppendPrependJoinValue non-empty -- valid values", actual)
}

func Test_BasicString_OnlySupportedErr_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B", "C"})
	err := bs.OnlySupportedErr("A")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr error -- B and C unsupported", actual)
}

func Test_BasicString_OnlySupportedMsgErr(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B", "C"})
	err := bs.OnlySupportedMsgErr("custom msg", "A")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedMsgErr error -- with custom message", actual)
}

// ── numberEnumBase — additional coverage ──

func Test_NumberEnumBase_NameWithValueOption_WithQuotation_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	// Access via BasicByte which embeds numberEnumBase
	result := bb.NameWithValueOption(byte(0), true)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameWithValueOption with quotation -- non-empty", actual)
}

func Test_NumberEnumBase_NameWithValueOption_NoQuotation(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.NameWithValueOption(byte(0), false)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameWithValueOption without quotation -- non-empty", actual)
}

func Test_NumberEnumBase_Format_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.Format("Enum {type-name} - {name} - {value}", byte(0))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Format compiles template -- byte enum", actual)
}

func Test_NumberEnumBase_OnlySupportedMsgErr_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B", "C"})
	err := bb.OnlySupportedMsgErr("custom", "A")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedMsgErr error -- with message", actual)
}

func Test_NumberEnumBase_RangesMap_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.RangesMap()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RangesMap returns map -- two items", actual)
}

func Test_NumberEnumBase_NamesHashset(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.NamesHashset()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NamesHashset returns hashset -- two items", actual)
}

func Test_NumberEnumBase_JsonString_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.JsonString(byte(0))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns non-empty -- byte value", actual)
}

func Test_NumberEnumBase_KeyValIntegers_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.KeyValIntegers()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyValIntegers returns slice -- two items", actual)
}

func Test_NumberEnumBase_Loop_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B", "C"})
	count := 0
	bb.Loop(func(index int, name string, anyVal any) (isBreak bool) {
		count++
		return index == 1 // break after second
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "Loop breaks early -- after second item", actual)
}

func Test_NumberEnumBase_LoopInteger(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B", "C"})
	count := 0
	bb.LoopInteger(func(index int, name string, anyVal int) (isBreak bool) {
		count++
		return index == 0 // break after first
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "LoopInteger breaks early -- after first item", actual)
}

func Test_NumberEnumBase_RangesIntegerStringMap_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.RangesIntegerStringMap()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RangesIntegerStringMap returns map -- two items", actual)
}

// ── ConvEnumAnyValToInteger — type switch branches ──

func Test_ConvEnumAnyValToInteger_String_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	result := enumimpl.ConvEnumAnyValToInteger("hello")

	// Act
	actual := args.Map{"isMinInt": result < 0}

	// Assert
	expected := args.Map{"isMinInt": true}
	expected.ShouldBeEqual(t, 0, "ConvEnumAnyValToInteger returns MinInt -- string input", actual)
}

func Test_ConvEnumAnyValToInteger_Int_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	result := enumimpl.ConvEnumAnyValToInteger(42)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ConvEnumAnyValToInteger returns 42 -- int input", actual)
}

// ── BasicByte — AppendPrependJoinNamer ──

type mockNamer struct{ name string }

func (m mockNamer) Name() string { return m.name }

func Test_BasicByte_AppendPrependJoinNamer(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.AppendPrependJoinNamer(".", mockNamer{"B"}, mockNamer{"A"})

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "A.B"}
	expected.ShouldBeEqual(t, 0, "AppendPrependJoinNamer joins names -- dot separator", actual)
}

// ── BasicByte — AsBasicByter ──

func Test_BasicByte_AsBasicByter_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	byter := bb.AsBasicByter()

	// Act
	actual := args.Map{"notNil": byter != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsBasicByter returns non-nil -- valid enum", actual)
}

// ── BasicByte — ToNumberString ──

func Test_BasicByte_ToNumberString_FromDiffLeftRightNilJson(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.ToNumberString(byte(1))

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "1"}
	expected.ShouldBeEqual(t, 0, "ToNumberString returns string -- byte 1", actual)
}

// ── DynamicMap — BasicByte / BasicString creation ──

func Test_DynamicMap_BasicByte(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": byte(0), "B": byte(1)}
	bb := dm.BasicByte("TestDM")

	// Act
	actual := args.Map{
		"notNil": bb != nil,
		"typeName": bb.TypeName(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"typeName": "TestDM",
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap.BasicByte creates enum -- two items", actual)
}

func Test_DynamicMap_BasicString(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": "alpha", "B": "beta"}
	bs := dm.BasicString("TestDM")

	// Act
	actual := args.Map{
		"notNil": bs != nil,
		"typeName": bs.TypeName(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"typeName": "TestDM",
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap.BasicString creates enum -- two items", actual)
}

func Test_DynamicMap_BasicInt8(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 0, "B": 1}
	bi := dm.BasicInt8("TestDM")

	// Act
	actual := args.Map{"notNil": bi != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicMap.BasicInt8 creates enum -- two items", actual)
}

func Test_DynamicMap_BasicInt16(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 0, "B": 1}
	bi := dm.BasicInt16("TestDM")

	// Act
	actual := args.Map{"notNil": bi != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicMap.BasicInt16 creates enum -- two items", actual)
}

func Test_DynamicMap_BasicInt32(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 0, "B": 1}
	bi := dm.BasicInt32("TestDM")

	// Act
	actual := args.Map{"notNil": bi != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicMap.BasicInt32 creates enum -- two items", actual)
}

func Test_DynamicMap_BasicUInt16(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 0, "B": 1}
	bi := dm.BasicUInt16("TestDM")

	// Act
	actual := args.Map{"notNil": bi != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicMap.BasicUInt16 creates enum -- two items", actual)
}
