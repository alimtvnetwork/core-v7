package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// LeftRight — all methods
// ═══════════════════════════════════════════

func Test_Cov13_LeftRight_Empty(t *testing.T) {
	lr := &coredynamic.LeftRight{}
	var nilLR *coredynamic.LeftRight
	actual := args.Map{"empty": lr.IsEmpty(), "has": lr.HasAnyItem(), "nilEmpty": nilLR.IsEmpty()}
	expected := args.Map{"empty": true, "has": false, "nilEmpty": true}
	expected.ShouldBeEqual(t, 0, "LeftRight returns empty -- empty", actual)
}

func Test_Cov13_LeftRight_HasLeftRight(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}
	actual := args.Map{
		"hasL": lr.HasLeft(), "hasR": lr.HasRight(),
		"lEmpty": lr.IsLeftEmpty(), "rEmpty": lr.IsRightEmpty(),
	}
	expected := args.Map{"hasL": true, "hasR": true, "lEmpty": false, "rEmpty": false}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- HasLeft/Right", actual)
}

func Test_Cov13_LeftRight_NilHasLeftRight(t *testing.T) {
	var lr *coredynamic.LeftRight
	actual := args.Map{"hasL": lr.HasLeft(), "hasR": lr.HasRight(), "lEmpty": lr.IsLeftEmpty(), "rEmpty": lr.IsRightEmpty()}
	expected := args.Map{"hasL": false, "hasR": false, "lEmpty": true, "rEmpty": true}
	expected.ShouldBeEqual(t, 0, "LeftRight returns nil -- nil HasLeft/Right", actual)
}

func Test_Cov13_LeftRight_ReflectSet(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "hello", Right: "world"}
	var l, r string
	errL := lr.LeftReflectSet(&l)
	errR := lr.RightReflectSet(&r)
	actual := args.Map{"l": l, "r": r, "noErrL": errL == nil, "noErrR": errR == nil}
	expected := args.Map{"l": "hello", "r": "world", "noErrL": true, "noErrR": true}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- ReflectSet", actual)
}

func Test_Cov13_LeftRight_ReflectSet_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	errL := lr.LeftReflectSet(nil)
	errR := lr.RightReflectSet(nil)
	actual := args.Map{"noErrL": errL == nil, "noErrR": errR == nil}
	expected := args.Map{"noErrL": true, "noErrR": true}
	expected.ShouldBeEqual(t, 0, "LeftRight returns nil -- ReflectSet nil", actual)
}

func Test_Cov13_LeftRight_Deserialize(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}
	dl := lr.DeserializeLeft()
	dr := lr.DeserializeRight()
	actual := args.Map{"dlNN": dl != nil, "drNN": dr != nil}
	expected := args.Map{"dlNN": true, "drNN": true}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Deserialize", actual)
}

func Test_Cov13_LeftRight_Deserialize_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	dl := lr.DeserializeLeft()
	dr := lr.DeserializeRight()
	actual := args.Map{"dlNil": dl == nil, "drNil": dr == nil}
	expected := args.Map{"dlNil": true, "drNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight returns nil -- Deserialize nil", actual)
}

func Test_Cov13_LeftRight_ToDynamic(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}
	ld := lr.LeftToDynamic()
	rd := lr.RightToDynamic()
	actual := args.Map{"ldNN": ld != nil, "rdNN": rd != nil}
	expected := args.Map{"ldNN": true, "rdNN": true}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- ToDynamic", actual)
}

func Test_Cov13_LeftRight_ToDynamic_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	ld := lr.LeftToDynamic()
	rd := lr.RightToDynamic()
	actual := args.Map{"ldNil": ld == nil, "rdNil": rd == nil}
	expected := args.Map{"ldNil": true, "rdNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight returns nil -- ToDynamic nil", actual)
}

func Test_Cov13_LeftRight_TypeStatus(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}
	ts := lr.TypeStatus()
	actual := args.Map{"same": ts.IsSame}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- TypeStatus", actual)
}

func Test_Cov13_LeftRight_TypeStatus_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	ts := lr.TypeStatus()
	actual := args.Map{"same": ts.IsSame}
	expected := args.Map{"same": true} // both nil
	expected.ShouldBeEqual(t, 0, "LeftRight returns nil -- TypeStatus nil", actual)
}

// ═══════════════════════════════════════════
// KeyValCollection — core methods
// ═══════════════════════════════════════════

func Test_Cov13_KeyValCollection_Basic(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(5)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	kvc.AddPtr(&coredynamic.KeyVal{Key: "b", Value: 2})
	kvc.AddPtr(nil)
	kvc.AddMany(coredynamic.KeyVal{Key: "c", Value: 3})
	kvc.AddManyPtr(&coredynamic.KeyVal{Key: "d", Value: 4}, nil)
	actual := args.Map{"len": kvc.Length(), "empty": kvc.IsEmpty(), "has": kvc.HasAnyItem()}
	expected := args.Map{"len": 4, "empty": false, "has": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- basic", actual)
}

func Test_Cov13_KeyValCollection_Items_Nil(t *testing.T) {
	var kvc *coredynamic.KeyValCollection
	actual := args.Map{"nil": kvc.Items() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns nil -- Items nil", actual)
}

func Test_Cov13_KeyValCollection_AllKeys(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(5)
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	keys := kvc.AllKeys()
	sorted := kvc.AllKeysSorted()
	vals := kvc.AllValues()
	actual := args.Map{"keysLen": len(keys), "sortedFirst": sorted[0], "valsLen": len(vals)}
	expected := args.Map{"keysLen": 2, "sortedFirst": "a", "valsLen": 2}
	expected.ShouldBeEqual(t, 0, "AllKeys returns correct value -- with args", actual)
}

func Test_Cov13_KeyValCollection_AllKeys_Empty(t *testing.T) {
	kvc := coredynamic.EmptyKeyValCollection()
	actual := args.Map{"keys": len(kvc.AllKeys()), "sorted": len(kvc.AllKeysSorted()), "vals": len(kvc.AllValues())}
	expected := args.Map{"keys": 0, "sorted": 0, "vals": 0}
	expected.ShouldBeEqual(t, 0, "AllKeys returns empty -- empty", actual)
}

func Test_Cov13_KeyValCollection_String(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	s := kvc.String()
	var nilKvc *coredynamic.KeyValCollection
	ns := nilKvc.String()
	actual := args.Map{"ne": s != "", "nil": ns}
	expected := args.Map{"ne": true, "nil": ""}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- with args", actual)
}

func Test_Cov13_KeyValCollection_MapAnyItems(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	m := kvc.MapAnyItems()
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- with args", actual)
}

func Test_Cov13_KeyValCollection_MapAnyItems_Empty(t *testing.T) {
	kvc := coredynamic.EmptyKeyValCollection()
	m := kvc.MapAnyItems()
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- empty", actual)
}

func Test_Cov13_KeyValCollection_Json(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	j := kvc.Json()
	jp := kvc.JsonPtr()
	jm := kvc.JsonModel()
	jma := kvc.JsonModelAny()
	actual := args.Map{"jLen": j.Length() > 0, "jpNN": jp != nil, "jmNN": jm != nil, "jmaNN": jma != nil}
	expected := args.Map{"jLen": true, "jpNN": true, "jmNN": true, "jmaNN": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Json", actual)
}
func Test_Cov13_KeyValCollection_Serialize(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	b, err := kvc.Serialize()
	js, jsErr := kvc.JsonString()
	// KeyValCollection.Json() serializes JsonModel(), so JsonString is non-empty.
	actual := args.Map{"bLen": len(b) > 0, "noErr": err == nil, "jsNE": js != "", "jsNoErr": jsErr == nil}
	expected := args.Map{"bLen": true, "noErr": true, "jsNE": true, "jsNoErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns correct value -- with args", actual)
}

func Test_Cov13_KeyValCollection_Paging(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(10)
	for i := 0; i < 5; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	ps := kvc.GetPagesSize(2)
	pz := kvc.GetPagesSize(0)
	paged := kvc.GetPagedCollection(2)
	small := kvc.GetPagedCollection(100)
	actual := args.Map{"ps": ps, "pz": pz, "pagedLen": len(paged), "smallLen": len(small)}
	expected := args.Map{"ps": 3, "pz": 0, "pagedLen": 3, "smallLen": 1}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Paging", actual)
}

func Test_Cov13_KeyValCollection_JsonMapResults(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	mr, err := kvc.JsonMapResults()
	actual := args.Map{"mrNN": mr != nil, "noErr": err == nil}
	expected := args.Map{"mrNN": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "JsonMapResults returns correct value -- with args", actual)
}

func Test_Cov13_KeyValCollection_JsonMapResults_Empty(t *testing.T) {
	kvc := coredynamic.EmptyKeyValCollection()
	mr, _ := kvc.JsonMapResults()
	actual := args.Map{"nn": mr != nil}
	expected := args.Map{"nn": true}
	expected.ShouldBeEqual(t, 0, "JsonMapResults returns empty -- empty", actual)
}

func Test_Cov13_KeyValCollection_JsonResultsCollection(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	rc := kvc.JsonResultsCollection()
	rpc := kvc.JsonResultsPtrCollection()
	actual := args.Map{"rcNN": rc != nil, "rpcNN": rpc != nil}
	expected := args.Map{"rcNN": true, "rpcNN": true}
	expected.ShouldBeEqual(t, 0, "JsonResultsCollection returns correct value -- with args", actual)
}

func Test_Cov13_KeyValCollection_JsonResultsCollection_Empty(t *testing.T) {
	kvc := coredynamic.EmptyKeyValCollection()
	rc := kvc.JsonResultsCollection()
	rpc := kvc.JsonResultsPtrCollection()
	actual := args.Map{"rcNN": rc != nil, "rpcNN": rpc != nil}
	expected := args.Map{"rcNN": true, "rpcNN": true}
	expected.ShouldBeEqual(t, 0, "JsonResultsCollection returns empty -- empty", actual)
}

func Test_Cov13_KeyValCollection_AddMany_Nil(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.AddMany()
	kvc.AddManyPtr()
	actual := args.Map{"len": kvc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddMany returns nil -- nil", actual)
}
