package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// MapAnyItemDiff — all methods
// ═══════════════════════════════════════════

func Test_Cov14_MapAnyItemDiff_Basic(t *testing.T) {
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1, "b": 2})
	var nilM *coredynamic.MapAnyItemDiff
	actual := args.Map{
		"len": m.Length(), "empty": m.IsEmpty(), "has": m.HasAnyItem(), "last": m.LastIndex(),
		"nilLen": nilM.Length(),
	}
	expected := args.Map{"len": 2, "empty": false, "has": true, "last": 1, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- basic", actual)
}

func Test_Cov14_MapAnyItemDiff_AllKeysSorted(t *testing.T) {
	m := coredynamic.MapAnyItemDiff(map[string]any{"b": 2, "a": 1})
	keys := m.AllKeysSorted()
	actual := args.Map{"first": keys[0]}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "AllKeysSorted returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItemDiff_Raw(t *testing.T) {
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	var nilM *coredynamic.MapAnyItemDiff
	raw := m.Raw()
	nilRaw := nilM.Raw()
	actual := args.Map{"rawLen": len(raw), "nilLen": len(nilRaw)}
	expected := args.Map{"rawLen": 1, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "Raw returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItemDiff_Clear(t *testing.T) {
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	cleared := m.Clear()
	actual := args.Map{"len": len(cleared)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clear returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItemDiff_Clear_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItemDiff
	cleared := m.Clear()
	actual := args.Map{"len": len(cleared)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clear returns nil -- nil", actual)
}

func Test_Cov14_MapAnyItemDiff_IsRawEqual(t *testing.T) {
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	actual := args.Map{"same": m.IsRawEqual(false, map[string]any{"a": 1})}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "IsRawEqual returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItemDiff_HasAnyChanges(t *testing.T) {
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	actual := args.Map{"changes": m.HasAnyChanges(false, map[string]any{"a": 2})}
	expected := args.Map{"changes": true}
	expected.ShouldBeEqual(t, 0, "HasAnyChanges returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItemDiff_DiffRaw(t *testing.T) {
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1, "b": 2})
	diff := m.DiffRaw(false, map[string]any{"a": 1, "b": 3})
	actual := args.Map{"hasItems": len(diff) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "DiffRaw returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItemDiff_HashmapDiffUsingRaw(t *testing.T) {
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	d := m.HashmapDiffUsingRaw(false, map[string]any{"a": 1})
	actual := args.Map{"empty": d.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "HashmapDiffUsingRaw returns correct value -- same", actual)
}

func Test_Cov14_MapAnyItemDiff_MapAnyItems(t *testing.T) {
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	mai := m.MapAnyItems()
	actual := args.Map{"len": mai.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItemDiff_RawMapDiffer(t *testing.T) {
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	d := m.RawMapDiffer()
	actual := args.Map{"len": len(d)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RawMapDiffer returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItemDiff_Json(t *testing.T) {
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	j := m.Json()
	jp := m.JsonPtr()
	pj := m.PrettyJsonString()
	actual := args.Map{"jLen": j.Length() > 0, "jpNN": jp != nil, "pjNE": pj != ""}
	expected := args.Map{"jLen": true, "jpNN": true, "pjNE": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- Json", actual)
}

func Test_Cov14_MapAnyItemDiff_DiffJsonMessage(t *testing.T) {
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	msg := m.DiffJsonMessage(false, map[string]any{"a": 2})
	actual := args.Map{"ne": msg != ""}
	expected := args.Map{"ne": true}
	expected.ShouldBeEqual(t, 0, "DiffJsonMessage returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItemDiff_ShouldDiffMessage(t *testing.T) {
	m := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	msg := m.ShouldDiffMessage(false, "test", map[string]any{"a": 2})
	actual := args.Map{"ne": msg != ""}
	expected := args.Map{"ne": true}
	expected.ShouldBeEqual(t, 0, "ShouldDiffMessage returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// MapAnyItems — core methods
// ═══════════════════════════════════════════

func Test_Cov14_MapAnyItems_Basic(t *testing.T) {
	m := coredynamic.NewMapAnyItems(5)
	actual := args.Map{"len": m.Length(), "empty": m.IsEmpty(), "has": m.HasAnyItem()}
	expected := args.Map{"len": 0, "empty": true, "has": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- basic", actual)
}

func Test_Cov14_MapAnyItems_NewUsingItems(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	e := coredynamic.NewMapAnyItemsUsingItems(map[string]any{})
	actual := args.Map{"len": m.Length(), "eLen": e.Length()}
	expected := args.Map{"len": 1, "eLen": 0}
	expected.ShouldBeEqual(t, 0, "NewMapAnyItemsUsingItems returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_HasKey(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	var nilM *coredynamic.MapAnyItems
	actual := args.Map{"has": m.HasKey("a"), "miss": m.HasKey("z"), "nil": nilM.HasKey("a")}
	expected := args.Map{"has": true, "miss": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "HasKey returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_AddSet(t *testing.T) {
	m := coredynamic.NewMapAnyItems(5)
	new1 := m.Add("a", 1)
	new2 := m.Add("a", 2)
	new3 := m.Set("b", 3)
	actual := args.Map{"new1": new1, "new2": new2, "new3": new3, "len": m.Length()}
	expected := args.Map{"new1": true, "new2": false, "new3": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "Add/Set returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_GetValue(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	v := m.GetValue("a")
	vMiss := m.GetValue("z")
	actual := args.Map{"v": v, "miss": vMiss == nil}
	expected := args.Map{"v": 1, "miss": true}
	expected.ShouldBeEqual(t, 0, "GetValue returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_Get(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	v, has := m.Get("a")
	_, miss := m.Get("z")
	actual := args.Map{"v": v, "has": has, "miss": miss}
	expected := args.Map{"v": 1, "has": true, "miss": false}
	expected.ShouldBeEqual(t, 0, "Get returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_AllKeys(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"b": 2, "a": 1})
	keys := m.AllKeys()
	sorted := m.AllKeysSorted()
	vals := m.AllValues()
	actual := args.Map{"keysLen": len(keys), "sortedFirst": sorted[0], "valsLen": len(vals)}
	expected := args.Map{"keysLen": 2, "sortedFirst": "a", "valsLen": 2}
	expected.ShouldBeEqual(t, 0, "AllKeys returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_AllKeys_Empty(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	actual := args.Map{"keys": len(m.AllKeys()), "sorted": len(m.AllKeysSorted()), "vals": len(m.AllValues())}
	expected := args.Map{"keys": 0, "sorted": 0, "vals": 0}
	expected.ShouldBeEqual(t, 0, "AllKeys returns empty -- empty", actual)
}

func Test_Cov14_MapAnyItems_IsEqual(t *testing.T) {
	m1 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m2 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m3 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 2})
	var nilM *coredynamic.MapAnyItems
	actual := args.Map{
		"eq":      m1.IsEqual(m2),
		"neq":     m1.IsEqual(m3),
		"bothNil": nilM.IsEqual(nil),
		"oneNil":  m1.IsEqual(nil),
	}
	expected := args.Map{"eq": true, "neq": false, "bothNil": true, "oneNil": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_IsEqualRaw_LenMismatch(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	actual := args.Map{"v": m.IsEqualRaw(map[string]any{"a": 1, "b": 2})}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsEqualRaw returns correct value -- len mismatch", actual)
}

func Test_Cov14_MapAnyItems_IsEqualRaw_KeyMismatch(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	actual := args.Map{"v": m.IsEqualRaw(map[string]any{"b": 1})}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsEqualRaw returns correct value -- key mismatch", actual)
}

func Test_Cov14_MapAnyItems_IsEqualRaw_ValueMismatch(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	actual := args.Map{"v": m.IsEqualRaw(map[string]any{"a": 2})}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsEqualRaw returns correct value -- value mismatch", actual)
}

func Test_Cov14_MapAnyItems_IsEqualRaw_NilBoth(t *testing.T) {
	var m *coredynamic.MapAnyItems
	actual := args.Map{"v": m.IsEqualRaw(nil)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEqualRaw returns nil -- nil both", actual)
}

func Test_Cov14_MapAnyItems_ClearDisposeDeepClear(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m.Clear()
	actual := args.Map{"len": m.Length()}
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

func Test_Cov14_MapAnyItems_ClearDispose_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItems
	m.Clear()
	m.DeepClear()
	m.Dispose()
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Clear/Dispose returns nil -- nil", actual)
}

func Test_Cov14_MapAnyItems_AddMapResult(t *testing.T) {
	m := coredynamic.NewMapAnyItems(5)
	m.AddMapResult(map[string]any{"a": 1})
	m.AddMapResult(nil)
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddMapResult returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_AddMapResultOption(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m.AddMapResultOption(true, map[string]any{"a": 2})
	actual := args.Map{"v": m.GetValue("a")}
	expected := args.Map{"v": 2}
	expected.ShouldBeEqual(t, 0, "AddMapResultOption returns error -- override", actual)
}

func Test_Cov14_MapAnyItems_AddMapResultOption_Empty(t *testing.T) {
	m := coredynamic.NewMapAnyItems(5)
	m.AddMapResultOption(true, nil)
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddMapResultOption returns empty -- empty", actual)
}

func Test_Cov14_MapAnyItems_AddManyMapResultsUsingOption(t *testing.T) {
	m := coredynamic.NewMapAnyItems(5)
	m.AddManyMapResultsUsingOption(true, map[string]any{"a": 1}, map[string]any{"b": 2})
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AddManyMapResultsUsingOption returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_AddManyMapResultsUsingOption_Empty(t *testing.T) {
	m := coredynamic.NewMapAnyItems(5)
	m.AddManyMapResultsUsingOption(true)
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddManyMapResultsUsingOption returns empty -- empty", actual)
}

func Test_Cov14_MapAnyItems_GetNewMapUsingKeys(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2, "c": 3})
	sub := m.GetNewMapUsingKeys(false, "a", "c")
	empty := m.GetNewMapUsingKeys(false)
	actual := args.Map{"subLen": sub.Length(), "emptyLen": empty.Length()}
	expected := args.Map{"subLen": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "GetNewMapUsingKeys returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_JsonString(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	js, err := m.JsonString()
	jm := m.JsonStringMust()
	actual := args.Map{"ne": js != "", "noErr": err == nil, "jmNE": jm != ""}
	expected := args.Map{"ne": true, "noErr": true, "jmNE": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_Strings_String(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	strs := m.Strings()
	s := m.String()
	actual := args.Map{"len": len(strs) > 0, "sNE": s != ""}
	expected := args.Map{"len": true, "sNE": true}
	expected.ShouldBeEqual(t, 0, "Strings/String returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_Paging(t *testing.T) {
	m := coredynamic.NewMapAnyItems(10)
	for i := 0; i < 5; i++ {
		m.Add("k"+string(rune('a'+i)), i)
	}
	ps := m.GetPagesSize(2)
	pz := m.GetPagesSize(0)
	paged := m.GetPagedCollection(2)
	small := m.GetPagedCollection(100)
	actual := args.Map{"ps": ps, "pz": pz, "pagedLen": len(paged), "smallLen": len(small)}
	expected := args.Map{"ps": 3, "pz": 0, "pagedLen": 3, "smallLen": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Paging", actual)
}

func Test_Cov14_MapAnyItems_Diff(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	m2 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 3})
	diff := m.Diff(false, m2)
	actual := args.Map{"diffNN": diff != nil}
	expected := args.Map{"diffNN": true}
	expected.ShouldBeEqual(t, 0, "Diff returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_IsRawEqual(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	actual := args.Map{"eq": m.IsRawEqual(false, map[string]any{"a": 1})}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsRawEqual returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_HasAnyChanges(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	actual := args.Map{"v": m.HasAnyChanges(false, map[string]any{"a": 2})}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "HasAnyChanges returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_MapStringAnyDiff(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	d := m.MapStringAnyDiff()
	actual := args.Map{"len": len(d)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapStringAnyDiff returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_RawMapStringAnyDiff_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItems
	d := m.RawMapStringAnyDiff()
	actual := args.Map{"len": len(d)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RawMapStringAnyDiff returns nil -- nil", actual)
}

func Test_Cov14_MapAnyItems_MapAnyItemsSelf(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	self := m.MapAnyItems()
	actual := args.Map{"same": self == m}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- self", actual)
}

func Test_Cov14_MapAnyItems_ClonePtr(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	cloned, err := m.ClonePtr()
	var nilM *coredynamic.MapAnyItems
	_, nilErr := nilM.ClonePtr()
	actual := args.Map{"clonedNN": cloned != nil, "noErr": err == nil, "nilErr": nilErr != nil}
	expected := args.Map{"clonedNN": true, "noErr": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_Json(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	j := m.Json()
	jp := m.JsonPtr()
	jm := m.JsonModel()
	jma := m.JsonModelAny()
	actual := args.Map{"jLen": j.Length() > 0, "jpNN": jp != nil, "jmNN": jm != nil, "jmaNN": jma != nil}
	expected := args.Map{"jLen": true, "jpNN": true, "jmNN": true, "jmaNN": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Json", actual)
}

func Test_Cov14_MapAnyItems_JsonResultOfKey(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	jr := m.JsonResultOfKey("a")
	jrMiss := m.JsonResultOfKey("z")
	actual := args.Map{"jrNN": jr != nil, "missNN": jrMiss != nil}
	expected := args.Map{"jrNN": true, "missNN": true}
	expected.ShouldBeEqual(t, 0, "JsonResultOfKey returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_JsonResultOfKeys(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	mr := m.JsonResultOfKeys("a", "b")
	mrEmpty := m.JsonResultOfKeys()
	actual := args.Map{"mrNN": mr != nil, "emptyNN": mrEmpty != nil}
	expected := args.Map{"mrNN": true, "emptyNN": true}
	expected.ShouldBeEqual(t, 0, "JsonResultOfKeys returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_JsonMapResults(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	mr, err := m.JsonMapResults()
	actual := args.Map{"nn": mr != nil, "noErr": err == nil}
	expected := args.Map{"nn": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "JsonMapResults returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_JsonMapResults_Empty(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	mr, _ := m.JsonMapResults()
	actual := args.Map{"nn": mr != nil}
	expected := args.Map{"nn": true}
	expected.ShouldBeEqual(t, 0, "JsonMapResults returns empty -- empty", actual)
}

func Test_Cov14_MapAnyItems_JsonResultsCollection(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	rc := m.JsonResultsCollection()
	rpc := m.JsonResultsPtrCollection()
	actual := args.Map{"rcNN": rc != nil, "rpcNN": rpc != nil}
	expected := args.Map{"rcNN": true, "rpcNN": true}
	expected.ShouldBeEqual(t, 0, "JsonResultsCollection returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_JsonResultsCollection_Empty(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	rc := m.JsonResultsCollection()
	rpc := m.JsonResultsPtrCollection()
	actual := args.Map{"rcNN": rc != nil, "rpcNN": rpc != nil}
	expected := args.Map{"rcNN": true, "rpcNN": true}
	expected.ShouldBeEqual(t, 0, "JsonResultsCollection returns empty -- empty", actual)
}

func Test_Cov14_MapAnyItems_Deserialize(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": float64(42)})
	var target float64
	err := m.Deserialize("a", &target)
	actual := args.Map{"noErr": err == nil, "target": target}
	expected := args.Map{"noErr": true, "target": float64(42)}
	expected.ShouldBeEqual(t, 0, "Deserialize returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_Deserialize_Missing(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	var target string
	err := m.Deserialize("z", &target)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns correct value -- missing", actual)
}

func Test_Cov14_MapAnyItems_ReflectSetTo(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "hello"})
	var target string
	err := m.ReflectSetTo("a", &target)
	actual := args.Map{"noErr": err == nil, "target": target}
	expected := args.Map{"noErr": true, "target": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectSetTo returns correct value -- with args", actual)
}

func Test_Cov14_MapAnyItems_ReflectSetTo_Missing(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	var target string
	err := m.ReflectSetTo("z", &target)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetTo returns correct value -- missing", actual)
}

func Test_Cov14_MapAnyItems_AddJsonResultPtr(t *testing.T) {
	m := coredynamic.NewMapAnyItems(5)
	m.AddJsonResultPtr("a", nil)
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddJsonResultPtr returns nil -- nil", actual)
}

func Test_Cov14_MapAnyItems_HashmapDiffUsingRaw(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	d := m.HashmapDiffUsingRaw(false, map[string]any{"a": 1})
	actual := args.Map{"empty": d.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "HashmapDiffUsingRaw returns correct value -- same", actual)
}

// ═══════════════════════════════════════════
// BytesConverter — additional methods
// ═══════════════════════════════════════════

func Test_Cov14_BytesConverter_DeserializeMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	var target string
	bc.DeserializeMust(&target)
	actual := args.Map{"v": target}
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "DeserializeMust returns correct value -- with args", actual)
}

func Test_Cov14_BytesConverter_ToHashmap(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`{"a":"1"}`))
	hm, err := bc.ToHashmap()
	actual := args.Map{"nn": hm != nil, "noErr": err == nil}
	expected := args.Map{"nn": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ToHashmap returns correct value -- with args", actual)
}

func Test_Cov14_BytesConverter_ToHashmap_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`invalid`))
	_, err := bc.ToHashmap()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToHashmap returns error -- invalid", actual)
}

func Test_Cov14_BytesConverter_ToHashset(t *testing.T) {
	// Hashset is a struct with unexported 'items map[string]bool' —
	// JSON array ["a","b"] can't unmarshal into that struct, so error expected
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	hs, err := bc.ToHashset()
	actual := args.Map{"nn": hs == nil, "hasErr": err != nil}
	expected := args.Map{"nn": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToHashset returns correct value -- with args", actual)
}

func Test_Cov14_BytesConverter_ToHashset_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`invalid`))
	_, err := bc.ToHashset()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToHashset returns error -- invalid", actual)
}

func Test_Cov14_BytesConverter_ToCollection(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	c, err := bc.ToCollection()
	actual := args.Map{"nn": c != nil, "noErr": err == nil}
	expected := args.Map{"nn": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ToCollection returns correct value -- with args", actual)
}

func Test_Cov14_BytesConverter_ToCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`invalid`))
	_, err := bc.ToCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToCollection returns error -- invalid", actual)
}

func Test_Cov14_BytesConverter_ToSimpleSlice(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	ss, err := bc.ToSimpleSlice()
	actual := args.Map{"nn": ss != nil, "noErr": err == nil}
	expected := args.Map{"nn": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ToSimpleSlice returns correct value -- with args", actual)
}

func Test_Cov14_BytesConverter_ToSimpleSlice_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`invalid`))
	_, err := bc.ToSimpleSlice()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToSimpleSlice returns error -- invalid", actual)
}
