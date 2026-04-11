package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// LeftRight — all methods
// ═══════════════════════════════════════════

func Test_LeftRight_Empty_FromLeftRightEmpty(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{}
	var nilLR *coredynamic.LeftRight

	// Act
	actual := args.Map{
		"empty": lr.IsEmpty(),
		"has": lr.HasAnyItem(),
		"nilEmpty": nilLR.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"has": false,
		"nilEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns empty -- empty", actual)
}

func Test_LeftRight_HasLeftRight(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}

	// Act
	actual := args.Map{
		"hasL": lr.HasLeft(), "hasR": lr.HasRight(),
		"lEmpty": lr.IsLeftEmpty(), "rEmpty": lr.IsRightEmpty(),
	}

	// Assert
	expected := args.Map{
		"hasL": true,
		"hasR": true,
		"lEmpty": false,
		"rEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- HasLeft/Right", actual)
}

func Test_LeftRight_NilHasLeftRight(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{
		"hasL": lr.HasLeft(),
		"hasR": lr.HasRight(),
		"lEmpty": lr.IsLeftEmpty(),
		"rEmpty": lr.IsRightEmpty(),
	}

	// Assert
	expected := args.Map{
		"hasL": false,
		"hasR": false,
		"lEmpty": true,
		"rEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns nil -- nil HasLeft/Right", actual)
}

func Test_LeftRight_ReflectSet(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "hello", Right: "world"}
	var l, r string
	errL := lr.LeftReflectSet(&l)
	errR := lr.RightReflectSet(&r)

	// Act
	actual := args.Map{
		"l": l,
		"r": r,
		"noErrL": errL == nil,
		"noErrR": errR == nil,
	}

	// Assert
	expected := args.Map{
		"l": "hello",
		"r": "world",
		"noErrL": true,
		"noErrR": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- ReflectSet", actual)
}

func Test_LeftRight_ReflectSet_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight
	errL := lr.LeftReflectSet(nil)
	errR := lr.RightReflectSet(nil)

	// Act
	actual := args.Map{
		"noErrL": errL == nil,
		"noErrR": errR == nil,
	}

	// Assert
	expected := args.Map{
		"noErrL": true,
		"noErrR": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns nil -- ReflectSet nil", actual)
}

func Test_LeftRight_Deserialize(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}
	dl := lr.DeserializeLeft()
	dr := lr.DeserializeRight()

	// Act
	actual := args.Map{
		"dlNN": dl != nil,
		"drNN": dr != nil,
	}

	// Assert
	expected := args.Map{
		"dlNN": true,
		"drNN": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Deserialize", actual)
}

func Test_LeftRight_Deserialize_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight
	dl := lr.DeserializeLeft()
	dr := lr.DeserializeRight()

	// Act
	actual := args.Map{
		"dlNil": dl == nil,
		"drNil": dr == nil,
	}

	// Assert
	expected := args.Map{
		"dlNil": true,
		"drNil": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns nil -- Deserialize nil", actual)
}

func Test_LeftRight_ToDynamic(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}
	ld := lr.LeftToDynamic()
	rd := lr.RightToDynamic()

	// Act
	actual := args.Map{
		"ldNN": ld != nil,
		"rdNN": rd != nil,
	}

	// Assert
	expected := args.Map{
		"ldNN": true,
		"rdNN": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- ToDynamic", actual)
}

func Test_LeftRight_ToDynamic_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight
	ld := lr.LeftToDynamic()
	rd := lr.RightToDynamic()

	// Act
	actual := args.Map{
		"ldNil": ld == nil,
		"rdNil": rd == nil,
	}

	// Assert
	expected := args.Map{
		"ldNil": true,
		"rdNil": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns nil -- ToDynamic nil", actual)
}

func Test_LeftRight_TypeStatus(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}
	ts := lr.TypeStatus()

	// Act
	actual := args.Map{"same": ts.IsSame}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- TypeStatus", actual)
}

func Test_LeftRight_TypeStatus_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight
	ts := lr.TypeStatus()

	// Act
	actual := args.Map{"same": ts.IsSame}

	// Assert
	expected := args.Map{"same": true} // both nil
	expected.ShouldBeEqual(t, 0, "LeftRight returns nil -- TypeStatus nil", actual)
}

// ═══════════════════════════════════════════
// KeyValCollection — core methods
// ═══════════════════════════════════════════

func Test_KeyValCollection_Basic_FromLeftRightEmpty(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(5)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	kvc.AddPtr(&coredynamic.KeyVal{Key: "b", Value: 2})
	kvc.AddPtr(nil)
	kvc.AddMany(coredynamic.KeyVal{Key: "c", Value: 3})
	kvc.AddManyPtr(&coredynamic.KeyVal{Key: "d", Value: 4}, nil)

	// Act
	actual := args.Map{
		"len": kvc.Length(),
		"empty": kvc.IsEmpty(),
		"has": kvc.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len": 4,
		"empty": false,
		"has": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- basic", actual)
}

func Test_KeyValCollection_Items_Nil(t *testing.T) {
	// Arrange
	var kvc *coredynamic.KeyValCollection

	// Act
	actual := args.Map{"nil": kvc.Items() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns nil -- Items nil", actual)
}

func Test_KeyValCollection_AllKeys_FromLeftRightEmpty(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(5)
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	keys := kvc.AllKeys()
	sorted := kvc.AllKeysSorted()
	vals := kvc.AllValues()

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

func Test_KeyValCollection_AllKeys_Empty(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{
		"keys": len(kvc.AllKeys()),
		"sorted": len(kvc.AllKeysSorted()),
		"vals": len(kvc.AllValues()),
	}

	// Assert
	expected := args.Map{
		"keys": 0,
		"sorted": 0,
		"vals": 0,
	}
	expected.ShouldBeEqual(t, 0, "AllKeys returns empty -- empty", actual)
}

func Test_KeyValCollection_String_FromLeftRightEmpty(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	s := kvc.String()
	var nilKvc *coredynamic.KeyValCollection
	ns := nilKvc.String()

	// Act
	actual := args.Map{
		"ne": s != "",
		"nil": ns,
	}

	// Assert
	expected := args.Map{
		"ne": true,
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- with args", actual)
}

func Test_KeyValCollection_MapAnyItems_FromLeftRightEmpty(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	m := kvc.MapAnyItems()

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- with args", actual)
}

func Test_KeyValCollection_MapAnyItems_Empty(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	m := kvc.MapAnyItems()

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- empty", actual)
}

func Test_KeyValCollection_Json_FromLeftRightEmpty(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	j := kvc.Json()
	jp := kvc.JsonPtr()
	jm := kvc.JsonModel()
	jma := kvc.JsonModelAny()

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
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Json", actual)
}

func Test_KeyValCollection_CloneNonPtrPtr(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	cloned := kvc.Clone()
	clonedPtr := kvc.ClonePtr()
	np := kvc.NonPtr()
	p := kvc.Ptr()
	var nilKvc *coredynamic.KeyValCollection
	nilClone := nilKvc.ClonePtr()

	// Act
	actual := args.Map{
		"cLen": cloned.Length(),
		"cpNN": clonedPtr != nil,
		"npLen": np.Length(),
		"pNN": p != nil,
		"nilNil": nilClone == nil,
	}

	// Assert
	expected := args.Map{
		"cLen": 1,
		"cpNN": true,
		"npLen": 1,
		"pNN": true,
		"nilNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Clone/NonPtr/Ptr returns correct value -- with args", actual)
}

func Test_KeyValCollection_Serialize_FromLeftRightEmpty(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	b, err := kvc.Serialize()
	js, jsErr := kvc.JsonString()
	// KeyValCollection.Json() serializes JsonModel(), so JsonString is non-empty.

	// Act
	actual := args.Map{
		"bLen": len(b) > 0,
		"noErr": err == nil,
		"jsNE": js != "",
		"jsNoErr": jsErr == nil,
	}

	// Assert
	expected := args.Map{
		"bLen": true,
		"noErr": true,
		"jsNE": true,
		"jsNoErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize returns correct value -- with args", actual)
}

func Test_KeyValCollection_Paging_FromLeftRightEmpty(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(10)
	for i := 0; i < 5; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	ps := kvc.GetPagesSize(2)
	pz := kvc.GetPagesSize(0)
	paged := kvc.GetPagedCollection(2)
	small := kvc.GetPagedCollection(100)

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
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Paging", actual)
}

func Test_KeyValCollection_JsonMapResults_FromLeftRightEmpty(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	mr, err := kvc.JsonMapResults()

	// Act
	actual := args.Map{
		"mrNN": mr != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"mrNN": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonMapResults returns correct value -- with args", actual)
}

func Test_KeyValCollection_JsonMapResults_Empty(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	mr, _ := kvc.JsonMapResults()

	// Act
	actual := args.Map{"nn": mr != nil}

	// Assert
	expected := args.Map{"nn": true}
	expected.ShouldBeEqual(t, 0, "JsonMapResults returns empty -- empty", actual)
}

func Test_KeyValCollection_JsonResultsCollection_FromLeftRightEmpty(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	rc := kvc.JsonResultsCollection()
	rpc := kvc.JsonResultsPtrCollection()

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

func Test_KeyValCollection_JsonResultsCollection_Empty(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	rc := kvc.JsonResultsCollection()
	rpc := kvc.JsonResultsPtrCollection()

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

func Test_KeyValCollection_AddMany_Nil(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.AddMany()
	kvc.AddManyPtr()

	// Act
	actual := args.Map{"len": kvc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddMany returns nil -- nil", actual)
}
