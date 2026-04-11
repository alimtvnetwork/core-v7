package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// MapAnyItemDiff — all methods
// ═══════════════════════════════════════════

func Test_MapAnyItemDiff_Basic_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1, "b": 2})
	var nilM *coredynamic.MapAnyItemDiff

	// Act
	actual := args.Map{
		"len": m.Length(), "empty": m.IsEmpty(), "has": m.HasAnyItem(), "last": m.LastIndex(),
		"nilLen": nilM.Length(),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"empty": false,
		"has": true,
		"last": 1,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- basic", actual)
}

func Test_MapAnyItemDiff_AllKeysSorted_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff(map[string]any{"b": 2, "a": 1})
	keys := m.AllKeysSorted()

	// Act
	actual := args.Map{"first": keys[0]}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "AllKeysSorted returns correct value -- with args", actual)
}

func Test_MapAnyItemDiff_Raw_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	var nilM *coredynamic.MapAnyItemDiff
	raw := m.Raw()
	nilRaw := nilM.Raw()

	// Act
	actual := args.Map{
		"rawLen": len(raw),
		"nilLen": len(nilRaw),
	}

	// Assert
	expected := args.Map{
		"rawLen": 1,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "Raw returns correct value -- with args", actual)
}

func Test_MapAnyItemDiff_Clear_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	cleared := m.Clear()

	// Act
	actual := args.Map{"len": len(cleared)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clear returns correct value -- with args", actual)
}

func Test_MapAnyItemDiff_Clear_Nil_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItemDiff
	cleared := m.Clear()

	// Act
	actual := args.Map{"len": len(cleared)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clear returns nil -- nil", actual)
}

func Test_MapAnyItemDiff_IsRawEqual_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})

	// Act
	actual := args.Map{"same": m.IsRawEqual(false, map[string]any{"a": 1})}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "IsRawEqual returns correct value -- with args", actual)
}

func Test_MapAnyItemDiff_HasAnyChanges_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})

	// Act
	actual := args.Map{"changes": m.HasAnyChanges(false, map[string]any{"a": 2})}

	// Assert
	expected := args.Map{"changes": true}
	expected.ShouldBeEqual(t, 0, "HasAnyChanges returns correct value -- with args", actual)
}

func Test_MapAnyItemDiff_DiffRaw_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1, "b": 2})
	diff := m.DiffRaw(false, map[string]any{"a": 1, "b": 3})

	// Act
	actual := args.Map{"hasItems": len(diff) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "DiffRaw returns correct value -- with args", actual)
}

func Test_MapAnyItemDiff_HashmapDiffUsingRaw_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	d := m.HashmapDiffUsingRaw(false, map[string]any{"a": 1})

	// Act
	actual := args.Map{"empty": d.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "HashmapDiffUsingRaw returns correct value -- same", actual)
}

func Test_MapAnyItemDiff_MapAnyItems_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	mai := m.MapAnyItems()

	// Act
	actual := args.Map{"len": mai.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- with args", actual)
}

func Test_MapAnyItemDiff_RawMapDiffer_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	d := m.RawMapDiffer()

	// Act
	actual := args.Map{"len": len(d)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawMapDiffer returns correct value -- with args", actual)
}

func Test_MapAnyItemDiff_Json_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	j := m.Json()
	jp := m.JsonPtr()
	pj := m.PrettyJsonString()

	// Act
	actual := args.Map{
		"jLen": j.Length() > 0,
		"jpNN": jp != nil,
		"pjNE": pj != "",
	}

	// Assert
	expected := args.Map{
		"jLen": true,
		"jpNN": true,
		"pjNE": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- Json", actual)
}

func Test_MapAnyItemDiff_DiffJsonMessage_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	msg := m.DiffJsonMessage(false, map[string]any{"a": 2})

	// Act
	actual := args.Map{"ne": msg != ""}

	// Assert
	expected := args.Map{"ne": true}
	expected.ShouldBeEqual(t, 0, "DiffJsonMessage returns correct value -- with args", actual)
}

func Test_MapAnyItemDiff_ShouldDiffMessage_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	msg := m.ShouldDiffMessage(false, "test", map[string]any{"a": 2})

	// Act
	actual := args.Map{"ne": msg != ""}

	// Assert
	expected := args.Map{"ne": true}
	expected.ShouldBeEqual(t, 0, "ShouldDiffMessage returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// MapAnyItems — core methods
// ═══════════════════════════════════════════

func Test_MapAnyItems_Basic_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(5)

	// Act
	actual := args.Map{
		"len": m.Length(),
		"empty": m.IsEmpty(),
		"has": m.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"empty": true,
		"has": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- basic", actual)
}

func Test_MapAnyItems_NewUsingItems_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	e := coredynamic.NewMapAnyItemsUsingItems(map[string]any{})

	// Act
	actual := args.Map{
		"len": m.Length(),
		"eLen": e.Length(),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"eLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "NewMapAnyItemsUsingItems returns correct value -- with args", actual)
}

func Test_MapAnyItems_HasKey_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	var nilM *coredynamic.MapAnyItems

	// Act
	actual := args.Map{
		"has": m.HasKey("a"),
		"miss": m.HasKey("z"),
		"nil": nilM.HasKey("a"),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"miss": false,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "HasKey returns correct value -- with args", actual)
}

func Test_MapAnyItems_AddSet_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(5)
	new1 := m.Add("a", 1)
	new2 := m.Add("a", 2)
	new3 := m.Set("b", 3)

	// Act
	actual := args.Map{
		"new1": new1,
		"new2": new2,
		"new3": new3,
		"len": m.Length(),
	}

	// Assert
	expected := args.Map{
		"new1": true,
		"new2": false,
		"new3": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "Add/Set returns correct value -- with args", actual)
}

func Test_MapAnyItems_GetValue_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	v := m.GetValue("a")
	vMiss := m.GetValue("z")

	// Act
	actual := args.Map{
		"v": v,
		"miss": vMiss == nil,
	}

	// Assert
	expected := args.Map{
		"v": 1,
		"miss": true,
	}
	expected.ShouldBeEqual(t, 0, "GetValue returns correct value -- with args", actual)
}

func Test_MapAnyItems_Get_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	v, has := m.Get("a")
	_, miss := m.Get("z")

	// Act
	actual := args.Map{
		"v": v,
		"has": has,
		"miss": miss,
	}

	// Assert
	expected := args.Map{
		"v": 1,
		"has": true,
		"miss": false,
	}
	expected.ShouldBeEqual(t, 0, "Get returns correct value -- with args", actual)
}

func Test_MapAnyItems_AllKeys_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"b": 2, "a": 1})
	keys := m.AllKeys()
	sorted := m.AllKeysSorted()
	vals := m.AllValues()

	// Act
	actual := args.Map{
		"keysLen": len(keys),
		"sortedFirst": sorted[0],
		"valsLen": len(vals),
	}

	// Assert
	expected := args.Map{
		"keysLen": 2,
		"sortedFirst": "a",
		"valsLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "AllKeys returns correct value -- with args", actual)
}

func Test_MapAnyItems_AllKeys_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{
		"keys": len(m.AllKeys()),
		"sorted": len(m.AllKeysSorted()),
		"vals": len(m.AllValues()),
	}

	// Assert
	expected := args.Map{
		"keys": 0,
		"sorted": 0,
		"vals": 0,
	}
	expected.ShouldBeEqual(t, 0, "AllKeys returns empty -- empty", actual)
}

func Test_MapAnyItems_IsEqual_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m1 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m2 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m3 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 2})
	var nilM *coredynamic.MapAnyItems

	// Act
	actual := args.Map{
		"eq":      m1.IsEqual(m2),
		"neq":     m1.IsEqual(m3),
		"bothNil": nilM.IsEqual(nil),
		"oneNil":  m1.IsEqual(nil),
	}

	// Assert
	expected := args.Map{
		"eq": true,
		"neq": false,
		"bothNil": true,
		"oneNil": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- with args", actual)
}

func Test_MapAnyItems_IsEqualRaw_LenMismatch(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{"v": m.IsEqualRaw(map[string]any{"a": 1, "b": 2})}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsEqualRaw returns correct value -- len mismatch", actual)
}

func Test_MapAnyItems_IsEqualRaw_KeyMismatch(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{"v": m.IsEqualRaw(map[string]any{"b": 1})}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsEqualRaw returns correct value -- key mismatch", actual)
}

func Test_MapAnyItems_IsEqualRaw_ValueMismatch(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{"v": m.IsEqualRaw(map[string]any{"a": 2})}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsEqualRaw returns correct value -- value mismatch", actual)
}

func Test_MapAnyItems_IsEqualRaw_NilBoth(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems

	// Act
	actual := args.Map{"v": m.IsEqualRaw(nil)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEqualRaw returns nil -- nil both", actual)
}

func Test_MapAnyItems_ClearDisposeDeepClear_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m.Clear()

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clear returns correct value -- with args", actual)
	m2 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m2.DeepClear()
	actual2 := args.Map{"len": m2.Length()}
	expected2 := args.Map{"len": 0}
	expected2.ShouldBeEqual(t, 0, "DeepClear returns correct value -- with args", actual2)
	m3 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m3.Dispose()
	actual3 := args.Map{"nil": m3.Items == nil}
	expected3 := args.Map{"nil": true}
	expected3.ShouldBeEqual(t, 0, "Dispose returns correct value -- with args", actual3)
}

func Test_MapAnyItems_ClearDispose_Nil(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems
	m.Clear()
	m.DeepClear()
	m.Dispose()

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Clear/Dispose returns nil -- nil", actual)
}

func Test_MapAnyItems_AddMapResult_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(5)
	m.AddMapResult(map[string]any{"a": 1})
	m.AddMapResult(nil)

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddMapResult returns correct value -- with args", actual)
}

func Test_MapAnyItems_AddMapResultOption(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m.AddMapResultOption(true, map[string]any{"a": 2})

	// Act
	actual := args.Map{"v": m.GetValue("a")}

	// Assert
	expected := args.Map{"v": 2}
	expected.ShouldBeEqual(t, 0, "AddMapResultOption returns error -- override", actual)
}

func Test_MapAnyItems_AddMapResultOption_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(5)
	m.AddMapResultOption(true, nil)

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddMapResultOption returns empty -- empty", actual)
}

func Test_MapAnyItems_AddManyMapResultsUsingOption(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(5)
	m.AddManyMapResultsUsingOption(true, map[string]any{"a": 1}, map[string]any{"b": 2})

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AddManyMapResultsUsingOption returns correct value -- with args", actual)
}

func Test_MapAnyItems_AddManyMapResultsUsingOption_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(5)
	m.AddManyMapResultsUsingOption(true)

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddManyMapResultsUsingOption returns empty -- empty", actual)
}

func Test_MapAnyItems_GetNewMapUsingKeys_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2, "c": 3})
	sub := m.GetNewMapUsingKeys(false, "a", "c")
	empty := m.GetNewMapUsingKeys(false)

	// Act
	actual := args.Map{
		"subLen": sub.Length(),
		"emptyLen": empty.Length(),
	}

	// Assert
	expected := args.Map{
		"subLen": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "GetNewMapUsingKeys returns correct value -- with args", actual)
}

func Test_MapAnyItems_JsonString_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	js, err := m.JsonString()
	jm := m.JsonStringMust()

	// Act
	actual := args.Map{
		"ne": js != "",
		"noErr": err == nil,
		"jmNE": jm != "",
	}

	// Assert
	expected := args.Map{
		"ne": true,
		"noErr": true,
		"jmNE": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct value -- with args", actual)
}

func Test_MapAnyItems_Strings_String(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	strs := m.Strings()
	s := m.String()

	// Act
	actual := args.Map{
		"len": len(strs) > 0,
		"sNE": s != "",
	}

	// Assert
	expected := args.Map{
		"len": true,
		"sNE": true,
	}
	expected.ShouldBeEqual(t, 0, "Strings/String returns correct value -- with args", actual)
}

func Test_MapAnyItems_Paging_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(10)
	for i := 0; i < 5; i++ {
		m.Add("k"+string(rune('a'+i)), i)
	}
	ps := m.GetPagesSize(2)
	pz := m.GetPagesSize(0)
	paged := m.GetPagedCollection(2)
	small := m.GetPagedCollection(100)

	// Act
	actual := args.Map{
		"ps": ps,
		"pz": pz,
		"pagedLen": len(paged),
		"smallLen": len(small),
	}

	// Assert
	expected := args.Map{
		"ps": 3,
		"pz": 0,
		"pagedLen": 3,
		"smallLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Paging", actual)
}

func Test_MapAnyItems_Diff_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	m2 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 3})
	diff := m.Diff(false, m2)

	// Act
	actual := args.Map{"diffNN": diff != nil}

	// Assert
	expected := args.Map{"diffNN": true}
	expected.ShouldBeEqual(t, 0, "Diff returns correct value -- with args", actual)
}

func Test_MapAnyItems_IsRawEqual_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{"eq": m.IsRawEqual(false, map[string]any{"a": 1})}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsRawEqual returns correct value -- with args", actual)
}

func Test_MapAnyItems_HasAnyChanges_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{"v": m.HasAnyChanges(false, map[string]any{"a": 2})}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "HasAnyChanges returns correct value -- with args", actual)
}

func Test_MapAnyItems_MapStringAnyDiff_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	d := m.MapStringAnyDiff()

	// Act
	actual := args.Map{"len": len(d)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapStringAnyDiff returns correct value -- with args", actual)
}

func Test_MapAnyItems_RawMapStringAnyDiff_Nil(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems
	d := m.RawMapStringAnyDiff()

	// Act
	actual := args.Map{"len": len(d)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RawMapStringAnyDiff returns nil -- nil", actual)
}

func Test_MapAnyItems_MapAnyItemsSelf_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	self := m.MapAnyItems()

	// Act
	actual := args.Map{"same": self == m}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- self", actual)
}

func Test_MapAnyItems_ClonePtr_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	cloned, err := m.ClonePtr()
	var nilM *coredynamic.MapAnyItems
	_, nilErr := nilM.ClonePtr()

	// Act
	actual := args.Map{
		"clonedNN": cloned != nil,
		"noErr": err == nil,
		"nilErr": nilErr != nil,
	}

	// Assert
	expected := args.Map{
		"clonedNN": true,
		"noErr": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns correct value -- with args", actual)
}

func Test_MapAnyItems_Json_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	j := m.Json()
	jp := m.JsonPtr()
	jm := m.JsonModel()
	jma := m.JsonModelAny()

	// Act
	actual := args.Map{
		"jLen": j.Length() > 0,
		"jpNN": jp != nil,
		"jmNN": jm != nil,
		"jmaNN": jma != nil,
	}

	// Assert
	expected := args.Map{
		"jLen": true,
		"jpNN": true,
		"jmNN": true,
		"jmaNN": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Json", actual)
}

func Test_MapAnyItems_JsonResultOfKey_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	jr := m.JsonResultOfKey("a")
	jrMiss := m.JsonResultOfKey("z")

	// Act
	actual := args.Map{
		"jrNN": jr != nil,
		"missNN": jrMiss != nil,
	}

	// Assert
	expected := args.Map{
		"jrNN": true,
		"missNN": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonResultOfKey returns correct value -- with args", actual)
}

func Test_MapAnyItems_JsonResultOfKeys_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	mr := m.JsonResultOfKeys("a", "b")
	mrEmpty := m.JsonResultOfKeys()

	// Act
	actual := args.Map{
		"mrNN": mr != nil,
		"emptyNN": mrEmpty != nil,
	}

	// Assert
	expected := args.Map{
		"mrNN": true,
		"emptyNN": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonResultOfKeys returns correct value -- with args", actual)
}

func Test_MapAnyItems_JsonMapResults(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	mr, err := m.JsonMapResults()

	// Act
	actual := args.Map{
		"nn": mr != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"nn": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonMapResults returns correct value -- with args", actual)
}

func Test_MapAnyItems_JsonMapResults_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	mr, _ := m.JsonMapResults()

	// Act
	actual := args.Map{"nn": mr != nil}

	// Assert
	expected := args.Map{"nn": true}
	expected.ShouldBeEqual(t, 0, "JsonMapResults returns empty -- empty", actual)
}

func Test_MapAnyItems_JsonResultsCollection(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	rc := m.JsonResultsCollection()
	rpc := m.JsonResultsPtrCollection()

	// Act
	actual := args.Map{
		"rcNN": rc != nil,
		"rpcNN": rpc != nil,
	}

	// Assert
	expected := args.Map{
		"rcNN": true,
		"rpcNN": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonResultsCollection returns correct value -- with args", actual)
}

func Test_MapAnyItems_JsonResultsCollection_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	rc := m.JsonResultsCollection()
	rpc := m.JsonResultsPtrCollection()

	// Act
	actual := args.Map{
		"rcNN": rc != nil,
		"rpcNN": rpc != nil,
	}

	// Assert
	expected := args.Map{
		"rcNN": true,
		"rpcNN": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonResultsCollection returns empty -- empty", actual)
}

func Test_MapAnyItems_Deserialize_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": float64(42)})
	var target float64
	err := m.Deserialize("a", &target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"target": target,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"target": float64(42),
	}
	expected.ShouldBeEqual(t, 0, "Deserialize returns correct value -- with args", actual)
}

func Test_MapAnyItems_Deserialize_Missing(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	var target string
	err := m.Deserialize("z", &target)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns correct value -- missing", actual)
}

func Test_MapAnyItems_ReflectSetTo(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "hello"})
	var target string
	err := m.ReflectSetTo("a", &target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"target": target,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"target": "hello",
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetTo returns correct value -- with args", actual)
}

func Test_MapAnyItems_ReflectSetTo_Missing(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	var target string
	err := m.ReflectSetTo("z", &target)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetTo returns correct value -- missing", actual)
}

func Test_MapAnyItems_AddJsonResultPtr_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(5)
	m.AddJsonResultPtr("a", nil)

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddJsonResultPtr returns nil -- nil", actual)
}

func Test_MapAnyItems_HashmapDiffUsingRaw_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	d := m.HashmapDiffUsingRaw(false, map[string]any{"a": 1})

	// Act
	actual := args.Map{"empty": d.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "HashmapDiffUsingRaw returns correct value -- same", actual)
}

// ═══════════════════════════════════════════
// BytesConverter — additional methods
// ═══════════════════════════════════════════

func Test_BytesConverter_DeserializeMust_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	var target string
	bc.DeserializeMust(&target)

	// Act
	actual := args.Map{"v": target}

	// Assert
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "DeserializeMust returns correct value -- with args", actual)
}

func Test_BytesConverter_ToHashmap_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`{"a":"1"}`))
	hm, err := bc.ToHashmap()

	// Act
	actual := args.Map{
		"nn": hm != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"nn": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToHashmap returns correct value -- with args", actual)
}

func Test_BytesConverter_ToHashmap_Invalid_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`invalid`))
	_, err := bc.ToHashmap()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToHashmap returns error -- invalid", actual)
}

func Test_BytesConverter_ToHashset_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	// Hashset is a struct with unexported 'items map[string]bool' —
	// JSON array ["a","b"] can't unmarshal into that struct, so error expected
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	hs, err := bc.ToHashset()

	// Act
	actual := args.Map{
		"nn": hs == nil,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"nn": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToHashset returns correct value -- with args", actual)
}

func Test_BytesConverter_ToHashset_Invalid_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`invalid`))
	_, err := bc.ToHashset()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToHashset returns error -- invalid", actual)
}

func Test_BytesConverter_ToCollection_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	c, err := bc.ToCollection()

	// Act
	actual := args.Map{
		"nn": c != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"nn": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToCollection returns correct value -- with args", actual)
}

func Test_BytesConverter_ToCollection_Invalid_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`invalid`))
	_, err := bc.ToCollection()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToCollection returns error -- invalid", actual)
}

func Test_BytesConverter_ToSimpleSlice_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	ss, err := bc.ToSimpleSlice()

	// Act
	actual := args.Map{
		"nn": ss != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"nn": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToSimpleSlice returns correct value -- with args", actual)
}

func Test_BytesConverter_ToSimpleSlice_Invalid_FromMapAnyItemDiffBasic(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`invalid`))
	_, err := bc.ToSimpleSlice()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToSimpleSlice returns error -- invalid", actual)
}
