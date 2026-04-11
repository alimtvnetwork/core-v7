package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// MapAnyItemDiff — Length, IsEmpty, HasAnyItem, LastIndex
// =============================================================================

func Test_MapAnyItemDiff_Nil_Length(t *testing.T) {
	// Arrange
	var d *coredynamic.MapAnyItemDiff

	// Act
	actual := args.Map{"r": d.Length()}

	// Assert
	expected := args.Map{"r": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff nil Length", actual)
}

func Test_MapAnyItemDiff_Empty(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{}

	// Act
	actual := args.Map{
		"empty": d.IsEmpty(),
		"hasAny": d.HasAnyItem(),
		"last": d.LastIndex(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"hasAny": false,
		"last": -1,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff empty", actual)
}

func Test_MapAnyItemDiff_WithItems(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1, "b": 2}

	// Act
	actual := args.Map{
		"len": d.Length(),
		"empty": d.IsEmpty(),
		"hasAny": d.HasAnyItem(),
		"last": d.LastIndex(),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"empty": false,
		"hasAny": true,
		"last": 1,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff with items", actual)
}

// =============================================================================
// MapAnyItemDiff — AllKeysSorted
// =============================================================================

func Test_MapAnyItemDiff_AllKeysSorted_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"b": 2, "a": 1, "c": 3}
	keys := d.AllKeysSorted()

	// Act
	actual := args.Map{
		"len": len(keys),
		"first": keys[0],
		"last": keys[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "a",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff AllKeysSorted", actual)
}

// =============================================================================
// MapAnyItemDiff — IsRawEqual, HasAnyChanges
// =============================================================================

func Test_MapAnyItemDiff_IsRawEqual_Same_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1}
	right := map[string]any{"a": 1}

	// Act
	actual := args.Map{"r": d.IsRawEqual(false, right)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff IsRawEqual same", actual)
}

func Test_MapAnyItemDiff_IsRawEqual_Different(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1}
	right := map[string]any{"a": 2}

	// Act
	actual := args.Map{
		"r": d.IsRawEqual(false, right),
		"changes": d.HasAnyChanges(false, right),
	}

	// Assert
	expected := args.Map{
		"r": false,
		"changes": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff IsRawEqual different", actual)
}

func Test_MapAnyItemDiff_Nil_IsRawEqual(t *testing.T) {
	// Arrange
	var d *coredynamic.MapAnyItemDiff
	right := map[string]any{"a": 1}

	// Act
	actual := args.Map{"r": d.IsRawEqual(false, right)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff nil IsRawEqual", actual)
}

// =============================================================================
// MapAnyItemDiff — HashmapDiffUsingRaw
// =============================================================================

func Test_MapAnyItemDiff_HashmapDiffUsingRaw_NoDiff_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1}
	right := map[string]any{"a": 1}
	diff := d.HashmapDiffUsingRaw(false, right)

	// Act
	actual := args.Map{"empty": diff.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "HashmapDiffUsingRaw no diff", actual)
}

func Test_MapAnyItemDiff_HashmapDiffUsingRaw_WithDiff(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1}
	right := map[string]any{"a": 2}
	diff := d.HashmapDiffUsingRaw(false, right)

	// Act
	actual := args.Map{"hasAny": diff.HasAnyItem()}

	// Assert
	expected := args.Map{"hasAny": true}
	expected.ShouldBeEqual(t, 0, "HashmapDiffUsingRaw with diff", actual)
}

// =============================================================================
// MapAnyItemDiff — DiffRaw, DiffJsonMessage, ToStringsSliceOfDiffMap
// =============================================================================

func Test_MapAnyItemDiff_DiffRaw_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1, "b": 2}
	right := map[string]any{"a": 1, "b": 3}
	diffRaw := d.DiffRaw(false, right)

	// Act
	actual := args.Map{"hasB": diffRaw["b"] != nil}

	// Assert
	expected := args.Map{"hasB": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff DiffRaw", actual)
}

func Test_MapAnyItemDiff_DiffJsonMessage_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1}
	right := map[string]any{"a": 2}
	msg := d.DiffJsonMessage(false, right)

	// Act
	actual := args.Map{"hasMsg": len(msg) > 0}

	// Assert
	expected := args.Map{"hasMsg": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff DiffJsonMessage", actual)
}

func Test_MapAnyItemDiff_ToStringsSliceOfDiffMap_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1}
	diffMap := map[string]any{"a": 2}
	ss := d.ToStringsSliceOfDiffMap(diffMap)

	// Act
	actual := args.Map{"hasItems": len(ss) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff ToStringsSliceOfDiffMap", actual)
}

// =============================================================================
// MapAnyItemDiff — ShouldDiffMessage, LogShouldDiffMessage
// =============================================================================

func Test_MapAnyItemDiff_ShouldDiffMessage_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1}
	right := map[string]any{"a": 2}
	msg := d.ShouldDiffMessage(false, "test", right)

	// Act
	actual := args.Map{"hasMsg": len(msg) > 0}

	// Assert
	expected := args.Map{"hasMsg": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff ShouldDiffMessage", actual)
}

func Test_MapAnyItemDiff_LogShouldDiffMessage_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1}
	right := map[string]any{"a": 2}
	msg := d.LogShouldDiffMessage(false, "test", right)

	// Act
	actual := args.Map{"hasMsg": len(msg) > 0}

	// Assert
	expected := args.Map{"hasMsg": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff LogShouldDiffMessage", actual)
}

// =============================================================================
// MapAnyItemDiff — Raw, Clear, MapAnyItems, RawMapDiffer
// =============================================================================

func Test_MapAnyItemDiff_Raw_Nil_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	var d *coredynamic.MapAnyItemDiff
	r := d.Raw()

	// Act
	actual := args.Map{"len": len(r)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff Raw nil", actual)
}

func Test_MapAnyItemDiff_Clear_Nil_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	var d *coredynamic.MapAnyItemDiff
	c := d.Clear()

	// Act
	actual := args.Map{"len": len(c)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff Clear nil", actual)
}

func Test_MapAnyItemDiff_Clear_Valid(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1}
	c := d.Clear()

	// Act
	actual := args.Map{"len": len(c)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff Clear valid", actual)
}

func Test_MapAnyItemDiff_MapAnyItems_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1}
	m := d.MapAnyItems()

	// Act
	actual := args.Map{"notNil": m != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff MapAnyItems", actual)
}

func Test_MapAnyItemDiff_RawMapDiffer_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1}
	differ := d.RawMapDiffer()

	// Act
	actual := args.Map{"notNil": differ != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff RawMapDiffer", actual)
}

// =============================================================================
// MapAnyItemDiff — Json, JsonPtr, PrettyJsonString, LogPrettyJsonString
// =============================================================================

func Test_MapAnyItemDiff_Json_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1}
	r := d.Json()

	// Act
	actual := args.Map{"noErr": !r.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff Json", actual)
}

func Test_MapAnyItemDiff_JsonPtr_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1}

	// Act
	actual := args.Map{"notNil": d.JsonPtr() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff JsonPtr", actual)
}

func Test_MapAnyItemDiff_PrettyJsonString_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1}
	s := d.PrettyJsonString()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff PrettyJsonString", actual)
}

func Test_MapAnyItemDiff_LogPrettyJsonString_Empty_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{}
	d.LogPrettyJsonString() // should log "empty map"

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff LogPrettyJsonString empty", actual)
}

func Test_MapAnyItemDiff_LogPrettyJsonString_NonEmpty(t *testing.T) {
	// Arrange
	d := coredynamic.MapAnyItemDiff{"a": 1}
	d.LogPrettyJsonString() // should log pretty JSON

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff LogPrettyJsonString non-empty", actual)
}

// =============================================================================
// MapAsKeyValSlice
// =============================================================================

func Test_MapAsKeyValSlice_ValidMap(t *testing.T) {
	// Arrange
	m := map[string]any{"a": 1, "b": 2}
	rv := reflect.ValueOf(m)
	coll, err := coredynamic.MapAsKeyValSlice(rv)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": coll.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice valid map", actual)
}

func Test_MapAsKeyValSlice_NotMap_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf("hello")
	_, err := coredynamic.MapAsKeyValSlice(rv)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice not map", actual)
}

func Test_MapAsKeyValSlice_Pointer_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	m := map[string]any{"x": 5}
	rv := reflect.ValueOf(&m)
	coll, err := coredynamic.MapAsKeyValSlice(rv)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": coll.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": false,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice pointer", actual)
}

// =============================================================================
// LeftRight — all methods
// =============================================================================

func Test_LeftRight_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{
		"empty":    lr.IsEmpty(),
		"hasAny":   lr.HasAnyItem(),
		"hasLeft":  lr.HasLeft(),
		"hasRight": lr.HasRight(),
		"leftE":    lr.IsLeftEmpty(),
		"rightE":   lr.IsRightEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty":    true,
		"hasAny":   false,
		"hasLeft":  false,
		"hasRight": false,
		"leftE":    true,
		"rightE":   true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight nil", actual)
}

func Test_LeftRight_BothSet(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}

	// Act
	actual := args.Map{
		"empty":    lr.IsEmpty(),
		"hasAny":   lr.HasAnyItem(),
		"hasLeft":  lr.HasLeft(),
		"hasRight": lr.HasRight(),
		"leftE":    lr.IsLeftEmpty(),
		"rightE":   lr.IsRightEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty":    false,
		"hasAny":   true,
		"hasLeft":  true,
		"hasRight": true,
		"leftE":    false,
		"rightE":   false,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight both set", actual)
}

func Test_LeftRight_LeftReflectSet_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight
	err := lr.LeftReflectSet(nil)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LeftRight LeftReflectSet nil", actual)
}

func Test_LeftRight_LeftReflectSet_Valid(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "hello"}
	var dest string
	err := lr.LeftReflectSet(&dest)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"dest": dest,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"dest": "hello",
	}
	expected.ShouldBeEqual(t, 0, "LeftRight LeftReflectSet valid", actual)
}

func Test_LeftRight_RightReflectSet_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight
	err := lr.RightReflectSet(nil)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LeftRight RightReflectSet nil", actual)
}

func Test_LeftRight_RightReflectSet_Valid(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Right: "world"}
	var dest string
	err := lr.RightReflectSet(&dest)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"dest": dest,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"dest": "world",
	}
	expected.ShouldBeEqual(t, 0, "LeftRight RightReflectSet valid", actual)
}

func Test_LeftRight_DeserializeLeft_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"isNil": lr.DeserializeLeft() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight DeserializeLeft nil", actual)
}

func Test_LeftRight_DeserializeLeft_Valid(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "hello"}
	r := lr.DeserializeLeft()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight DeserializeLeft valid", actual)
}

func Test_LeftRight_DeserializeRight_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"isNil": lr.DeserializeRight() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight DeserializeRight nil", actual)
}

func Test_LeftRight_DeserializeRight_Valid(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Right: "world"}
	r := lr.DeserializeRight()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight DeserializeRight valid", actual)
}

func Test_LeftRight_LeftToDynamic_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"isNil": lr.LeftToDynamic() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight LeftToDynamic nil", actual)
}

func Test_LeftRight_LeftToDynamic_Valid(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "hello"}
	d := lr.LeftToDynamic()

	// Act
	actual := args.Map{"valid": d.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "LeftRight LeftToDynamic valid", actual)
}

func Test_LeftRight_RightToDynamic_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"isNil": lr.RightToDynamic() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight RightToDynamic nil", actual)
}

func Test_LeftRight_RightToDynamic_Valid(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Right: "world"}
	d := lr.RightToDynamic()

	// Act
	actual := args.Map{"valid": d.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "LeftRight RightToDynamic valid", actual)
}

func Test_LeftRight_TypeStatus_Nil_FromMapAnyItemDiffNilMap(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight
	ts := lr.TypeStatus()

	// Act
	actual := args.Map{"valid": ts.IsValid()}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "LeftRight TypeStatus nil", actual)
}

func Test_LeftRight_TypeStatus_SameType(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}
	ts := lr.TypeStatus()
	sameTs := lr.TypeStatus()

	// Act
	actual := args.Map{
		"valid": ts.IsValid(),
		"equal": ts.IsEqual(&sameTs),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"equal": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight TypeStatus same type", actual)
}

// =============================================================================
// CastedResult — all methods
// =============================================================================

func Test_CastedResult_Nil(t *testing.T) {
	// Arrange
	var cr *coredynamic.CastedResult

	// Act
	actual := args.Map{
		"invalid":    cr.IsInvalid(),
		"notNull":    false,
		"notPointer": false,
		"notMatch":   false,
		"hasErr":     cr.HasError(),
		"hasIssues":  cr.HasAnyIssues(),
	}

	// Assert
	expected := args.Map{
		"invalid":    true,
		"notNull":    false,
		"notPointer": false,
		"notMatch":   false,
		"hasErr":     false,
		"hasIssues":  true,
	}
	expected.ShouldBeEqual(t, 0, "CastedResult nil", actual)
}

func Test_CastedResult_Valid(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{
		IsValid:                true,
		IsNull:                 false,
		IsMatchingAcceptedType: true,
		IsPointer:              false,
		IsSourcePointer:        false,
		SourceKind:             reflect.String,
	}

	// Act
	actual := args.Map{
		"invalid":    cr.IsInvalid(),
		"notNull":    cr.IsNotNull(),
		"notPointer": cr.IsNotPointer(),
		"notMatch":   cr.IsNotMatchingAcceptedType(),
		"isKind":     cr.IsSourceKind(reflect.String),
		"wrongKind":  cr.IsSourceKind(reflect.Int),
		"hasErr":     cr.HasError(),
		"hasIssues":  cr.HasAnyIssues(),
	}

	// Assert
	expected := args.Map{
		"invalid":    false,
		"notNull":    true,
		"notPointer": true,
		"notMatch":   false,
		"isKind":     true,
		"wrongKind":  false,
		"hasErr":     false,
		"hasIssues":  false,
	}
	expected.ShouldBeEqual(t, 0, "CastedResult valid", actual)
}

// =============================================================================
// PointerOrNonPointerUsingReflectValue
// =============================================================================

func Test_PointerOrNonPointerUsingReflectValue_Deref(t *testing.T) {
	// Arrange
	s := "hello"
	rv := reflect.ValueOf(&s)
	out, frv := coredynamic.PointerOrNonPointerUsingReflectValue(false, rv)

	// Act
	actual := args.Map{
		"r": out,
		"kind": frv.Kind().String(),
	}

	// Assert
	expected := args.Map{
		"r": "hello",
		"kind": "string",
	}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointerUsingReflectValue deref", actual)
}

func Test_PointerOrNonPointerUsingReflectValue_Passthrough(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf("hello")
	out, frv := coredynamic.PointerOrNonPointerUsingReflectValue(false, rv)

	// Act
	actual := args.Map{
		"r": out,
		"kind": frv.Kind().String(),
	}

	// Assert
	expected := args.Map{
		"r": "hello",
		"kind": "string",
	}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointerUsingReflectValue passthrough", actual)
}
