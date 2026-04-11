package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// KeyValCollection — constructors / Length / IsEmpty / HasAnyItem
// =============================================================================

func Test_KVC_Empty(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{
		"len": kvc.Length(),
		"empty": kvc.IsEmpty(),
		"hasAny": kvc.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"empty": true,
		"hasAny": false,
	}
	expected.ShouldBeEqual(t, 0, "KVC Empty", actual)
}

func Test_KVC_Length_Nil(t *testing.T) {
	// Arrange
	var kvc *coredynamic.KeyValCollection

	// Act
	actual := args.Map{"len": kvc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC Length nil", actual)
}

func Test_KVC_New(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(5)

	// Act
	actual := args.Map{"len": kvc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC New", actual)
}

// =============================================================================
// Add / AddPtr / AddMany / AddManyPtr
// =============================================================================

func Test_KVC_Add_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})

	// Act
	actual := args.Map{"len": kvc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KVC Add", actual)
}

func Test_KVC_AddPtr_Nil(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.AddPtr(nil)

	// Act
	actual := args.Map{"len": kvc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC AddPtr nil", actual)
}

func Test_KVC_AddPtr_Valid(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.AddPtr(&coredynamic.KeyVal{Key: "a", Value: 1})

	// Act
	actual := args.Map{"len": kvc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KVC AddPtr valid", actual)
}

func Test_KVC_AddMany_Nil(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.AddMany()

	// Act
	actual := args.Map{"len": kvc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC AddMany nil", actual)
}

func Test_KVC_AddMany_Valid(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.AddMany(
		coredynamic.KeyVal{Key: "a", Value: 1},
		coredynamic.KeyVal{Key: "b", Value: 2},
	)

	// Act
	actual := args.Map{"len": kvc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KVC AddMany valid", actual)
}

func Test_KVC_AddManyPtr_Nil(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.AddManyPtr()

	// Act
	actual := args.Map{"len": kvc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC AddManyPtr nil", actual)
}

func Test_KVC_AddManyPtr_MixedNils(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kv := &coredynamic.KeyVal{Key: "a", Value: 1}
	kvc.AddManyPtr(kv, nil, kv)

	// Act
	actual := args.Map{"len": kvc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KVC AddManyPtr mixed nils", actual)
}

// =============================================================================
// Items
// =============================================================================

func Test_KVC_Items_Nil(t *testing.T) {
	// Arrange
	var kvc *coredynamic.KeyValCollection

	// Act
	actual := args.Map{"isNil": kvc.Items() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "KVC Items nil", actual)
}

func Test_KVC_Items_Valid(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})

	// Act
	actual := args.Map{"len": len(kvc.Items())}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KVC Items valid", actual)
}

// =============================================================================
// MapAnyItems
// =============================================================================

func Test_KVC_MapAnyItems_Empty_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	m := kvc.MapAnyItems()

	// Act
	actual := args.Map{"empty": m.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "KVC MapAnyItems empty", actual)
}

func Test_KVC_MapAnyItems_Valid(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	m := kvc.MapAnyItems()

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KVC MapAnyItems valid", actual)
}

// =============================================================================
// AllKeys / AllKeysSorted / AllValues
// =============================================================================

func Test_KVC_AllKeys_Empty_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{"len": len(kvc.AllKeys())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC AllKeys empty", actual)
}

func Test_KVC_AllKeys_Valid(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: 1})
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 2})
	keys := kvc.AllKeys()

	// Act
	actual := args.Map{
		"len": len(keys),
		"first": keys[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "b",
	}
	expected.ShouldBeEqual(t, 0, "KVC AllKeys valid", actual)
}

func Test_KVC_AllKeysSorted_Empty(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{"len": len(kvc.AllKeysSorted())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC AllKeysSorted empty", actual)
}

func Test_KVC_AllKeysSorted_Valid(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "c", Value: 1})
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 2})
	keys := kvc.AllKeysSorted()

	// Act
	actual := args.Map{
		"first": keys[0],
		"second": keys[1],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"second": "c",
	}
	expected.ShouldBeEqual(t, 0, "KVC AllKeysSorted valid", actual)
}

func Test_KVC_AllValues_Empty_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{"len": len(kvc.AllValues())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC AllValues empty", actual)
}

func Test_KVC_AllValues_Valid(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 42})
	vals := kvc.AllValues()

	// Act
	actual := args.Map{"len": len(vals)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KVC AllValues valid", actual)
}

// =============================================================================
// GetPagesSize / GetPagedCollection / GetSinglePageCollection
// =============================================================================

func Test_KVC_GetPagesSize_ZeroPageSize(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{"r": kvc.GetPagesSize(0)}

	// Assert
	expected := args.Map{"r": 0}
	expected.ShouldBeEqual(t, 0, "KVC GetPagesSize zero", actual)
}

func Test_KVC_GetPagesSize_Valid(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	for i := 0; i < 5; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}

	// Act
	actual := args.Map{"r": kvc.GetPagesSize(2)}

	// Assert
	expected := args.Map{"r": 3}
	expected.ShouldBeEqual(t, 0, "KVC GetPagesSize valid", actual)
}

func Test_KVC_GetPagedCollection_SmallSet_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	pages := kvc.GetPagedCollection(10)

	// Act
	actual := args.Map{"pages": len(pages)}

	// Assert
	expected := args.Map{"pages": 1}
	expected.ShouldBeEqual(t, 0, "KVC GetPagedCollection small", actual)
}

func Test_KVC_GetPagedCollection_MultiPage(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	for i := 0; i < 5; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	pages := kvc.GetPagedCollection(2)

	// Act
	actual := args.Map{"pages": len(pages)}

	// Assert
	expected := args.Map{"pages": 3}
	expected.ShouldBeEqual(t, 0, "KVC GetPagedCollection multi", actual)
}

func Test_KVC_GetSinglePageCollection_SmallSet(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	page := kvc.GetSinglePageCollection(10, 1)

	// Act
	actual := args.Map{"len": page.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KVC GetSinglePageCollection small", actual)
}

func Test_KVC_GetSinglePageCollection_Page1(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	for i := 0; i < 5; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	page := kvc.GetSinglePageCollection(2, 1)

	// Act
	actual := args.Map{"len": page.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KVC GetSinglePageCollection page1", actual)
}

// =============================================================================
// String / JsonModel / Json / JsonPtr
// =============================================================================

func Test_KVC_String_Nil_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	var kvc *coredynamic.KeyValCollection

	// Act
	actual := args.Map{"r": kvc.String()}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "KVC String nil", actual)
}

func Test_KVC_String_Valid(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	s := kvc.String()

	// Act
	actual := args.Map{"nonEmpty": s != ""}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "KVC String valid", actual)
}

func Test_KVC_JsonModel_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{"notNil": kvc.JsonModel() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC JsonModel", actual)
}

func Test_KVC_Json_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	r := kvc.Json()

	// Act
	actual := args.Map{"noErr": !r.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KVC Json", actual)
}

func Test_KVC_JsonPtr_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{"notNil": kvc.JsonPtr() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC JsonPtr", actual)
}

// =============================================================================
// JSON round-trip / Serialize / JsonString / JsonStringMust
// =============================================================================

func Test_KVC_Serialize_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	b, err := kvc.Serialize()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "KVC Serialize", actual)
}

func Test_KVC_JsonString_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	s, err := kvc.JsonString()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"nonEmpty": s != "",
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"nonEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "KVC JsonString", actual)
}

func Test_KVC_JsonStringMust_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	s := kvc.JsonStringMust()

	// Act
	actual := args.Map{"nonEmpty": s != ""}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "KVC JsonStringMust", actual)
}

func Test_KVC_ParseInjectUsingJson_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	jr := kvc.Json()
	kvc2 := coredynamic.EmptyKeyValCollection()
	result, err := kvc2.ParseInjectUsingJson(&jr)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": result != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "KVC ParseInjectUsingJson", actual)
}

func Test_KVC_ParseInjectUsingJsonMust_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	jr := kvc.Json()
	kvc2 := coredynamic.EmptyKeyValCollection()
	result := kvc2.ParseInjectUsingJsonMust(&jr)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC ParseInjectUsingJsonMust", actual)
}

func Test_KVC_JsonParseSelfInject_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	jr := kvc.Json()
	kvc2 := coredynamic.EmptyKeyValCollection()
	err := kvc2.JsonParseSelfInject(&jr)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KVC JsonParseSelfInject", actual)
}

// =============================================================================
// JsonMapResults / JsonResultsCollection / JsonResultsPtrCollection
// =============================================================================

func Test_KVC_JsonMapResults_Empty_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	mr, err := kvc.JsonMapResults()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"empty": mr.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "KVC JsonMapResults empty", actual)
}

func Test_KVC_JsonMapResults_Valid(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: "hello"})
	mr, err := kvc.JsonMapResults()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"empty": mr.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"empty": false,
	}
	expected.ShouldBeEqual(t, 0, "KVC JsonMapResults valid", actual)
}

func Test_KVC_JsonResultsCollection_Empty(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	rc := kvc.JsonResultsCollection()

	// Act
	actual := args.Map{"empty": rc.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "KVC JsonResultsCollection empty", actual)
}

func Test_KVC_JsonResultsCollection_Valid(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: "hello"})
	rc := kvc.JsonResultsCollection()

	// Act
	actual := args.Map{"empty": rc.IsEmpty()}

	// Assert
	expected := args.Map{"empty": false}
	expected.ShouldBeEqual(t, 0, "KVC JsonResultsCollection valid", actual)
}

func Test_KVC_JsonResultsPtrCollection_Empty(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	rc := kvc.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"empty": rc.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "KVC JsonResultsPtrCollection empty", actual)
}

func Test_KVC_JsonResultsPtrCollection_Valid(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: "hello"})
	rc := kvc.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"empty": rc.IsEmpty()}

	// Assert
	expected := args.Map{"empty": false}
	expected.ShouldBeEqual(t, 0, "KVC JsonResultsPtrCollection valid", actual)
}

// =============================================================================
// Clone / ClonePtr / NonPtr / Ptr
// =============================================================================

func Test_KVC_Clone_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	c := kvc.Clone()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KVC Clone", actual)
}

func Test_KVC_ClonePtr_Nil_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	var kvc *coredynamic.KeyValCollection

	// Act
	actual := args.Map{"isNil": kvc.ClonePtr() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "KVC ClonePtr nil", actual)
}

func Test_KVC_ClonePtr_Valid(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	c := kvc.ClonePtr()

	// Act
	actual := args.Map{
		"notNil": c != nil,
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "KVC ClonePtr valid", actual)
}

func Test_KVC_NonPtr_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	np := kvc.NonPtr()

	// Act
	actual := args.Map{"len": np.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC NonPtr", actual)
}

func Test_KVC_Ptr_FromKVCEmptyKeyValCollec(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{"notNil": kvc.Ptr() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC Ptr", actual)
}

// =============================================================================
// GetPagingInfo
// =============================================================================

func Test_KVC_GetPagingInfo(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()
	for i := 0; i < 10; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	pi := kvc.GetPagingInfo(3, 1)

	// Act
	actual := args.Map{"skip": pi.SkipItems}

	// Assert
	expected := args.Map{"skip": 0}
	expected.ShouldBeEqual(t, 0, "KVC GetPagingInfo", actual)
}
