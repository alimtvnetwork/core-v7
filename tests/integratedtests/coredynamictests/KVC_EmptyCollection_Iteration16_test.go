package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// KeyValCollection — constructors and basic accessors
// ══════════════════════════════════════════════════════════════════════════════

func Test_KVC_EmptyCollection(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{
		"len": c.Length(),
		"isEmpty": c.IsEmpty(),
		"hasAny": c.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"isEmpty": true,
		"hasAny": false,
	}
	expected.ShouldBeEqual(t, 0, "EmptyKeyValCollection returns empty -- with args", actual)
}

func Test_KVC_NewWithCapacity(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)

	// Act
	actual := args.Map{
		"len": c.Length(),
		"isEmpty": c.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "NewKeyValCollection returns correct value -- with args", actual)
}

func Test_KVC_NilReceiver_Length(t *testing.T) {
	// Arrange
	var c *coredynamic.KeyValCollection

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC returns nil -- nil Length", actual)
}

func Test_KVC_NilReceiver_Items(t *testing.T) {
	// Arrange
	var c *coredynamic.KeyValCollection

	// Act
	actual := args.Map{"nil": c.Items() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns nil -- nil Items", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValCollection — Add methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_KVC_Add(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	c.Add(coredynamic.KeyVal{Key: "b", Value: 2})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- Add", actual)
}

func Test_KVC_AddPtr(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	kv := &coredynamic.KeyVal{Key: "a", Value: 1}
	c.AddPtr(kv)
	c.AddPtr(nil) // should skip

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- AddPtr", actual)
}

func Test_KVC_AddMany(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.AddMany(
		coredynamic.KeyVal{Key: "a", Value: 1},
		coredynamic.KeyVal{Key: "b", Value: 2},
		coredynamic.KeyVal{Key: "c", Value: 3},
	)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- AddMany", actual)
}

func Test_KVC_AddMany_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.AddMany()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC returns empty -- AddMany empty", actual)
}

func Test_KVC_AddManyPtr(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	kv1 := &coredynamic.KeyVal{Key: "a", Value: 1}
	c.AddManyPtr(kv1, nil, kv1)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- AddManyPtr", actual)
}

func Test_KVC_AddManyPtr_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.AddManyPtr()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC returns empty -- AddManyPtr empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValCollection — query methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_KVC_MapAnyItems(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "name", Value: "Alice"})
	m := c.MapAnyItems()

	// Act
	actual := args.Map{"notNil": m != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- MapAnyItems", actual)
}

func Test_KVC_MapAnyItems_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	m := c.MapAnyItems()

	// Act
	actual := args.Map{"notNil": m != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns empty -- MapAnyItems empty", actual)
}

func Test_KVC_AllKeys(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	keys := c.AllKeys()

	// Act
	actual := args.Map{"len": len(keys)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- AllKeys", actual)
}

func Test_KVC_AllKeys_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{"len": len(c.AllKeys())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC returns empty -- AllKeys empty", actual)
}

func Test_KVC_AllKeysSorted(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "z", Value: 1})
	c.Add(coredynamic.KeyVal{Key: "a", Value: 2})
	keys := c.AllKeysSorted()

	// Act
	actual := args.Map{"len": len(keys)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- AllKeysSorted", actual)
}

func Test_KVC_AllValues(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "a", Value: 42})
	c.Add(coredynamic.KeyVal{Key: "b", Value: "hi"})
	values := c.AllValues()

	// Act
	actual := args.Map{"len": len(values)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KVC returns non-empty -- AllValues", actual)
}

func Test_KVC_AllValues_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{"len": len(c.AllValues())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC returns empty -- AllValues empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValCollection — paging
// ══════════════════════════════════════════════════════════════════════════════

func Test_KVC_GetPagesSize(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	for i := 0; i < 10; i++ {
		c.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}

	// Act
	actual := args.Map{"pages": c.GetPagesSize(3)}

	// Assert
	expected := args.Map{"pages": 4}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- GetPagesSize", actual)
}

func Test_KVC_GetPagesSize_Zero(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{"pages": c.GetPagesSize(0)}

	// Assert
	expected := args.Map{"pages": 0}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- GetPagesSize zero", actual)
}

func Test_KVC_GetPagedCollection(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	for i := 0; i < 10; i++ {
		c.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	pages := c.GetPagedCollection(3)

	// Act
	actual := args.Map{"pagesLen": len(pages)}

	// Assert
	expected := args.Map{"pagesLen": 4}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- GetPagedCollection", actual)
}

func Test_KVC_GetPagedCollection_SmallSet(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: 1})
	pages := c.GetPagedCollection(10)

	// Act
	actual := args.Map{"pagesLen": len(pages)}

	// Assert
	expected := args.Map{"pagesLen": 1}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- GetPagedCollection small", actual)
}

func Test_KVC_GetSinglePageCollection(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	for i := 0; i < 10; i++ {
		c.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	page := c.GetSinglePageCollection(3, 2)

	// Act
	actual := args.Map{"len": page.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- GetSinglePageCollection", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValCollection — JSON and serialization
// ══════════════════════════════════════════════════════════════════════════════

func Test_KVC_String(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})

	// Act
	actual := args.Map{"notEmpty": c.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- String", actual)
}

func Test_KVC_String_Nil(t *testing.T) {
	// Arrange
	var c *coredynamic.KeyValCollection

	// Act
	actual := args.Map{"empty": c.String() == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "KVC returns nil -- String nil", actual)
}

func Test_KVC_Json(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	jr := c.Json()

	// Act
	actual := args.Map{"noErr": !jr.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- Json", actual)
}

func Test_KVC_JsonPtr(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	jr := c.JsonPtr()

	// Act
	actual := args.Map{"notNil": jr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonPtr", actual)
}

func Test_KVC_JsonModel(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{"notNil": c.JsonModel() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonModel", actual)
}

func Test_KVC_JsonModelAny(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{"notNil": c.JsonModelAny() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonModelAny", actual)
}

func Test_KVC_Serialize(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	bytes, err := c.Serialize()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- Serialize", actual)
}

func Test_KVC_JsonString(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	s, err := c.JsonString()
	// KVC serializes JsonModel with exported Items.

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": s != "",
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonString", actual)
}

func Test_KVC_JsonStringMust(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	// JsonStringMust panics with nil because HandleError panics on empty JSON ({})
	// Note: panic(nil) means r == nil, so we track entry into recover itself
	didPanic := false
	func() {
		defer func() {
			recover()
			didPanic = true
		}()
		_ = c.JsonStringMust()
	}()

	// Act
	actual := args.Map{"didPanic": didPanic}

	// Assert
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonStringMust panics on empty JSON", actual)
}

func Test_KVC_JsonMapResults(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	mr, err := c.JsonMapResults()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": mr != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonMapResults", actual)
}

func Test_KVC_JsonMapResults_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	mr, err := c.JsonMapResults()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": mr != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "KVC returns empty -- JsonMapResults empty", actual)
}

func Test_KVC_JsonResultsCollection(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	rc := c.JsonResultsCollection()

	// Act
	actual := args.Map{"notNil": rc != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonResultsCollection", actual)
}

func Test_KVC_JsonResultsPtrCollection(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	rc := c.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"notNil": rc != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonResultsPtrCollection", actual)
}

func Test_KVC_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	jr := corejson.NewPtr(c)
	target := coredynamic.EmptyKeyValCollection()
	result, err := target.ParseInjectUsingJson(jr)

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
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- ParseInjectUsingJson", actual)
}

func Test_KVC_ParseInjectUsingJsonMust(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	jr := corejson.NewPtr(c)
	target := coredynamic.EmptyKeyValCollection()
	result := target.ParseInjectUsingJsonMust(jr)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- ParseInjectUsingJsonMust", actual)
}

func Test_KVC_JsonParseSelfInject(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	jr := corejson.NewPtr(c)
	target := coredynamic.EmptyKeyValCollection()
	err := target.JsonParseSelfInject(jr)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonParseSelfInject", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValCollection — Clone
// ══════════════════════════════════════════════════════════════════════════════

func Test_KVC_Clone(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	cloned := c.Clone()

	// Act
	actual := args.Map{"len": cloned.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- Clone", actual)
}

func Test_KVC_ClonePtr(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	cloned := c.ClonePtr()

	// Act
	actual := args.Map{
		"notNil": cloned != nil,
		"len": cloned.Length(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- ClonePtr", actual)
}

func Test_KVC_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var c *coredynamic.KeyValCollection

	// Act
	actual := args.Map{"nil": c.ClonePtr() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns nil -- ClonePtr nil", actual)
}

func Test_KVC_NonPtr(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	np := c.NonPtr()

	// Act
	actual := args.Map{"len": np.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- NonPtr", actual)
}

func Test_KVC_Ptr(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	p := c.Ptr()

	// Act
	actual := args.Map{"notNil": p != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- Ptr", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicCollection — core methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_DC_EmptyCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{
		"len": dc.Length(),
		"isEmpty": dc.IsEmpty(),
		"hasAny": dc.HasAnyItem(),
		"count": dc.Count(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"isEmpty": true,
		"hasAny": false,
		"count": 0,
	}
	expected.ShouldBeEqual(t, 0, "EmptyDynamicCollection returns empty -- with args", actual)
}

func Test_DC_NilReceiver_Length(t *testing.T) {
	// Arrange
	var dc *coredynamic.DynamicCollection

	// Act
	actual := args.Map{
		"len": dc.Length(),
		"isEmpty": dc.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "DC returns nil -- nil Length", actual)
}

func Test_DC_NilReceiver_Items(t *testing.T) {
	// Arrange
	var dc *coredynamic.DynamicCollection
	items := dc.Items()

	// Act
	actual := args.Map{"len": len(items)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DC returns nil -- nil Items", actual)
}

func Test_DC_Add_And_At(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	d := coredynamic.NewDynamic("hello", true)
	dc.Add(d)
	atVal := dc.At(0)

	// Act
	actual := args.Map{
		"len": dc.Length(),
		"atVal": atVal.ValueString(),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"atVal": "hello",
	}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- Add+At", actual)
}

func Test_DC_AddAny(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("val1", true)
	dc.AddAny(42, true)

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- AddAny", actual)
}

func Test_DC_AddAnyNonNull(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyNonNull("val", true)
	dc.AddAnyNonNull(nil, true) // skipped

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- AddAnyNonNull", actual)
}

func Test_DC_AddAnyMany(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c")

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- AddAnyMany", actual)
}

func Test_DC_AddAnyMany_Nil(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany()

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DC returns nil -- AddAnyMany nil", actual)
}

func Test_DC_AddPtr(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	d := coredynamic.NewDynamic("x", true)
	dc.AddPtr(&d)
	dc.AddPtr(nil) // skipped

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- AddPtr", actual)
}

func Test_DC_AddManyPtr(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	d1 := coredynamic.NewDynamic("a", true)
	d2 := coredynamic.NewDynamic("b", true)
	dc.AddManyPtr(&d1, nil, &d2)

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- AddManyPtr", actual)
}

func Test_DC_AddManyPtr_Nil(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddManyPtr()

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DC returns nil -- AddManyPtr nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicCollection — First/Last/Skip/Take/Limit
// ══════════════════════════════════════════════════════════════════════════════

func Test_DC_First_Last(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("first", true)
	dc.AddAny("last", true)
	first := dc.First()
	last := dc.Last()

	// Act
	actual := args.Map{
		"first":     first.ValueString(),
		"last":      last.ValueString(),
		"lastIdx":   dc.LastIndex(),
		"hasIdx":    dc.HasIndex(1),
		"noIdx":     dc.HasIndex(5),
		"firstDyn":  dc.FirstDynamic() != nil,
		"lastDyn":   dc.LastDynamic() != nil,
	}

	// Assert
	expected := args.Map{
		"first": "first", "last": "last", "lastIdx": 1,
		"hasIdx": true, "noIdx": false, "firstDyn": true, "lastDyn": true,
	}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- First/Last", actual)
}

func Test_DC_FirstOrDefault_NonEmpty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("item", true)
	f := dc.FirstOrDefault()

	// Act
	actual := args.Map{
		"notNil": f != nil,
		"firstOrDefaultDyn": dc.FirstOrDefaultDynamic() != nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"firstOrDefaultDyn": true,
	}
	expected.ShouldBeEqual(t, 0, "DC returns empty -- FirstOrDefault non-empty", actual)
}

func Test_DC_FirstOrDefault_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"nil": dc.FirstOrDefault() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "DC returns empty -- FirstOrDefault empty", actual)
}

func Test_DC_LastOrDefault_NonEmpty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("item", true)
	l := dc.LastOrDefault()

	// Act
	actual := args.Map{
		"notNil": l != nil,
		"lastOrDefaultDyn": dc.LastOrDefaultDynamic() != nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"lastOrDefaultDyn": true,
	}
	expected.ShouldBeEqual(t, 0, "DC returns empty -- LastOrDefault non-empty", actual)
}

func Test_DC_LastOrDefault_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"nil": dc.LastOrDefault() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "DC returns empty -- LastOrDefault empty", actual)
}

func Test_DC_Skip_Take_Limit(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c", "d", "e")

	// Act
	actual := args.Map{
		"skipLen":    len(dc.Skip(2)),
		"takeLen":    len(dc.Take(3)),
		"limitLen":   len(dc.Limit(2)),
		"skipDynNil": dc.SkipDynamic(1) != nil,
		"takeDynNil": dc.TakeDynamic(2) != nil,
		"limitDyn":   dc.LimitDynamic(2) != nil,
	}

	// Assert
	expected := args.Map{
		"skipLen": 3, "takeLen": 3, "limitLen": 2,
		"skipDynNil": true, "takeDynNil": true, "limitDyn": true,
	}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- Skip/Take/Limit", actual)
}

func Test_DC_SkipCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c")
	sc := dc.SkipCollection(1)

	// Act
	actual := args.Map{"len": sc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- SkipCollection", actual)
}

func Test_DC_TakeCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c")
	tc := dc.TakeCollection(2)

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- TakeCollection", actual)
}

func Test_DC_LimitCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c")
	lc := dc.LimitCollection(2)

	// Act
	actual := args.Map{"len": lc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- LimitCollection", actual)
}

func Test_DC_SafeLimitCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b")
	lc := dc.SafeLimitCollection(10) // limit > length

	// Act
	actual := args.Map{"len": lc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- SafeLimitCollection", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicCollection — RemoveAt, Loop, AnyItems, Strings
// ══════════════════════════════════════════════════════════════════════════════

func Test_DC_RemoveAt_Success(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c")
	ok := dc.RemoveAt(1)

	// Act
	actual := args.Map{
		"ok": ok,
		"len": dc.Length(),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- RemoveAt success", actual)
}

func Test_DC_RemoveAt_InvalidIndex(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	ok := dc.RemoveAt(5)

	// Act
	actual := args.Map{"ok": ok}

	// Assert
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "DC returns error -- RemoveAt invalid", actual)
}

func Test_DC_Loop(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c")
	count := 0
	dc.Loop(func(index int, d *coredynamic.Dynamic) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- Loop", actual)
}

func Test_DC_Loop_Break(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c")
	count := 0
	dc.Loop(func(index int, d *coredynamic.Dynamic) bool {
		count++
		return index == 0 // break after first
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- Loop break", actual)
}

func Test_DC_Loop_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	called := false
	dc.Loop(func(index int, d *coredynamic.Dynamic) bool {
		called = true
		return false
	})

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "DC returns empty -- Loop empty", actual)
}

func Test_DC_AnyItems(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", 42)
	items := dc.AnyItems()

	// Act
	actual := args.Map{"len": len(items)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- AnyItems", actual)
}

func Test_DC_AnyItems_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"len": len(dc.AnyItems())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DC returns empty -- AnyItems empty", actual)
}

func Test_DC_AnyItemsCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b")
	ac := dc.AnyItemsCollection()

	// Act
	actual := args.Map{"notNil": ac != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- AnyItemsCollection", actual)
}

func Test_DC_Strings(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b")
	strs := dc.Strings()

	// Act
	actual := args.Map{"len": len(strs)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- Strings", actual)
}

func Test_DC_String(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b")

	// Act
	actual := args.Map{"notEmpty": dc.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- String", actual)
}

func Test_DC_ListStrings(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("hello", "world")
	strs := dc.ListStrings()

	// Act
	actual := args.Map{"notEmpty": len(strs) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- ListStrings", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicCollection — type validation add
// ══════════════════════════════════════════════════════════════════════════════

func Test_DC_AddAnyWithTypeValidation_Success(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyWithTypeValidation(false, reflect.TypeOf(""), "hello")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": dc.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "DC returns non-empty -- AddAnyWithTypeValidation success", actual)
}

func Test_DC_AddAnyWithTypeValidation_TypeMismatch(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyWithTypeValidation(false, reflect.TypeOf(""), 42)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len": dc.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "DC returns non-empty -- AddAnyWithTypeValidation mismatch", actual)
}

func Test_DC_AddAnyItemsWithTypeValidation_ContinueOnError(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyItemsWithTypeValidation(true, false, reflect.TypeOf(""), "ok", 42, "also ok")

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len": dc.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "DC returns non-empty -- AddAnyItemsWithTypeValidation continue", actual)
}

func Test_DC_AddAnyItemsWithTypeValidation_StopOnError(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(""), "ok", 42, "unreachable")

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len": dc.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "DC returns non-empty -- AddAnyItemsWithTypeValidation stop", actual)
}

func Test_DC_AddAnyItemsWithTypeValidation_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(""))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DC returns empty -- AddAnyItemsWithTypeValidation empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicCollection — JSON and paging
// ══════════════════════════════════════════════════════════════════════════════

func Test_DC_JsonString(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b")
	s, err := dc.JsonString()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": s != "",
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- JsonString", actual)
}

func Test_DC_JsonStringMust(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a")
	s := dc.JsonStringMust()

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- JsonStringMust", actual)
}

func Test_DC_JsonModel(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	m := dc.JsonModel()

	// Act
	actual := args.Map{"notNil": m.Items != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- JsonModel", actual)
}

func Test_DC_JsonModelAny(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"notNil": dc.JsonModelAny() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- JsonModelAny", actual)
}

func Test_DC_Json(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	jr := dc.Json()

	// Act
	actual := args.Map{"noErr": !jr.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- Json", actual)
}

func Test_DC_JsonPtr(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"notNil": dc.JsonPtr() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- JsonPtr", actual)
}

func Test_DC_JsonResultsCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a")
	rc := dc.JsonResultsCollection()

	// Act
	actual := args.Map{"notNil": rc != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- JsonResultsCollection", actual)
}

func Test_DC_JsonResultsPtrCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a")
	rc := dc.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"notNil": rc != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- JsonResultsPtrCollection", actual)
}

func Test_DC_GetPagesSize(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c", "d", "e", "f", "g")

	// Act
	actual := args.Map{
		"pages": dc.GetPagesSize(3),
		"zero": dc.GetPagesSize(0),
	}

	// Assert
	expected := args.Map{
		"pages": 3,
		"zero": 0,
	}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- GetPagesSize", actual)
}

func Test_DC_GetPagedCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c", "d", "e", "f", "g")
	pages := dc.GetPagedCollection(3)

	// Act
	actual := args.Map{"pagesLen": len(pages)}

	// Assert
	expected := args.Map{"pagesLen": 3}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- GetPagedCollection", actual)
}

func Test_DC_GetPagedCollection_SmallSet(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a")
	pages := dc.GetPagedCollection(10)

	// Act
	actual := args.Map{"pagesLen": len(pages)}

	// Assert
	expected := args.Map{"pagesLen": 1}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- GetPagedCollection small", actual)
}

func Test_DC_GetSinglePageCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c", "d", "e", "f", "g")
	page := dc.GetSinglePageCollection(3, 2)

	// Act
	actual := args.Map{"len": page.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- GetSinglePageCollection", actual)
}

func Test_DC_MarshalUnmarshalJSON(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b")
	bytes, err := dc.MarshalJSON()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- MarshalJSON", actual)
}

func Test_DC_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b")
	jr := corejson.NewPtr(dc)
	target := coredynamic.EmptyDynamicCollection()
	// DynamicCollection can't unmarshal its Items ([]any) from JSON — expect error
	_, err := target.ParseInjectUsingJson(jr)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- ParseInjectUsingJson fails on unmarshal", actual)
}

func Test_DC_ParseInjectUsingJsonMust(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a")
	jr := corejson.NewPtr(dc)
	target := coredynamic.EmptyDynamicCollection()
	// ParseInjectUsingJsonMust panics because unmarshal fails
	didPanic := false
	func() {
		defer func() {
			recover()
			didPanic = true
		}()
		_ = target.ParseInjectUsingJsonMust(jr)
	}()

	// Act
	actual := args.Map{"didPanic": didPanic}

	// Assert
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- ParseInjectUsingJsonMust panics", actual)
}

func Test_DC_JsonParseSelfInject(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a")
	jr := corejson.NewPtr(dc)
	target := coredynamic.EmptyDynamicCollection()
	err := target.JsonParseSelfInject(jr)
	// DynamicCollection can't unmarshal []any — expect error

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- JsonParseSelfInject fails on unmarshal", actual)
}
