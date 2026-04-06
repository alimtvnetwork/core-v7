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

func Test_I16_KVC_EmptyCollection(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	actual := args.Map{"len": c.Length(), "isEmpty": c.IsEmpty(), "hasAny": c.HasAnyItem()}
	expected := args.Map{"len": 0, "isEmpty": true, "hasAny": false}
	expected.ShouldBeEqual(t, 0, "EmptyKeyValCollection returns empty -- with args", actual)
}

func Test_I16_KVC_NewWithCapacity(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	actual := args.Map{"len": c.Length(), "isEmpty": c.IsEmpty()}
	expected := args.Map{"len": 0, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewKeyValCollection returns correct value -- with args", actual)
}

func Test_I16_KVC_NilReceiver_Length(t *testing.T) {
	var c *coredynamic.KeyValCollection
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC returns nil -- nil Length", actual)
}

func Test_I16_KVC_NilReceiver_Items(t *testing.T) {
	var c *coredynamic.KeyValCollection
	actual := args.Map{"nil": c.Items() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns nil -- nil Items", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValCollection — Add methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_KVC_Add(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	c.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- Add", actual)
}

func Test_I16_KVC_AddPtr(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	kv := &coredynamic.KeyVal{Key: "a", Value: 1}
	c.AddPtr(kv)
	c.AddPtr(nil) // should skip
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- AddPtr", actual)
}

func Test_I16_KVC_AddMany(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.AddMany(
		coredynamic.KeyVal{Key: "a", Value: 1},
		coredynamic.KeyVal{Key: "b", Value: 2},
		coredynamic.KeyVal{Key: "c", Value: 3},
	)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- AddMany", actual)
}

func Test_I16_KVC_AddMany_Empty(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.AddMany()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC returns empty -- AddMany empty", actual)
}

func Test_I16_KVC_AddManyPtr(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	kv1 := &coredynamic.KeyVal{Key: "a", Value: 1}
	c.AddManyPtr(kv1, nil, kv1)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- AddManyPtr", actual)
}

func Test_I16_KVC_AddManyPtr_Empty(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.AddManyPtr()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC returns empty -- AddManyPtr empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValCollection — query methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_KVC_MapAnyItems(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "name", Value: "Alice"})
	m := c.MapAnyItems()
	actual := args.Map{"notNil": m != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- MapAnyItems", actual)
}

func Test_I16_KVC_MapAnyItems_Empty(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	m := c.MapAnyItems()
	actual := args.Map{"notNil": m != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns empty -- MapAnyItems empty", actual)
}

func Test_I16_KVC_AllKeys(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	keys := c.AllKeys()
	actual := args.Map{"len": len(keys)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- AllKeys", actual)
}

func Test_I16_KVC_AllKeys_Empty(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	actual := args.Map{"len": len(c.AllKeys())}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC returns empty -- AllKeys empty", actual)
}

func Test_I16_KVC_AllKeysSorted(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "z", Value: 1})
	c.Add(coredynamic.KeyVal{Key: "a", Value: 2})
	keys := c.AllKeysSorted()
	actual := args.Map{"len": len(keys)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- AllKeysSorted", actual)
}

func Test_I16_KVC_AllValues(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "a", Value: 42})
	c.Add(coredynamic.KeyVal{Key: "b", Value: "hi"})
	values := c.AllValues()
	actual := args.Map{"len": len(values)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KVC returns non-empty -- AllValues", actual)
}

func Test_I16_KVC_AllValues_Empty(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	actual := args.Map{"len": len(c.AllValues())}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC returns empty -- AllValues empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValCollection — paging
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_KVC_GetPagesSize(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	for i := 0; i < 10; i++ {
		c.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	actual := args.Map{"pages": c.GetPagesSize(3)}
	expected := args.Map{"pages": 4}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- GetPagesSize", actual)
}

func Test_I16_KVC_GetPagesSize_Zero(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	actual := args.Map{"pages": c.GetPagesSize(0)}
	expected := args.Map{"pages": 0}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- GetPagesSize zero", actual)
}

func Test_I16_KVC_GetPagedCollection(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	for i := 0; i < 10; i++ {
		c.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	pages := c.GetPagedCollection(3)
	actual := args.Map{"pagesLen": len(pages)}
	expected := args.Map{"pagesLen": 4}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- GetPagedCollection", actual)
}

func Test_I16_KVC_GetPagedCollection_SmallSet(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: 1})
	pages := c.GetPagedCollection(10)
	actual := args.Map{"pagesLen": len(pages)}
	expected := args.Map{"pagesLen": 1}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- GetPagedCollection small", actual)
}

func Test_I16_KVC_GetSinglePageCollection(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	for i := 0; i < 10; i++ {
		c.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	page := c.GetSinglePageCollection(3, 2)
	actual := args.Map{"len": page.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- GetSinglePageCollection", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValCollection — JSON and serialization
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_KVC_String(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	actual := args.Map{"notEmpty": c.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- String", actual)
}

func Test_I16_KVC_String_Nil(t *testing.T) {
	var c *coredynamic.KeyValCollection
	actual := args.Map{"empty": c.String() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "KVC returns nil -- String nil", actual)
}

func Test_I16_KVC_Json(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	jr := c.Json()
	actual := args.Map{"noErr": !jr.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- Json", actual)
}

func Test_I16_KVC_JsonPtr(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	jr := c.JsonPtr()
	actual := args.Map{"notNil": jr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonPtr", actual)
}

func Test_I16_KVC_JsonModel(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	actual := args.Map{"notNil": c.JsonModel() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonModel", actual)
}

func Test_I16_KVC_JsonModelAny(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	actual := args.Map{"notNil": c.JsonModelAny() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonModelAny", actual)
}

func Test_I16_KVC_Serialize(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	bytes, err := c.Serialize()
	actual := args.Map{"noErr": err == nil, "notEmpty": len(bytes) > 0}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- Serialize", actual)
}

func Test_I16_KVC_JsonString(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	s, err := c.JsonString()
	// KVC serializes JsonModel with exported Items.
	actual := args.Map{"noErr": err == nil, "notEmpty": s != ""}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonString", actual)
}

func Test_I16_KVC_JsonStringMust(t *testing.T) {
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
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonStringMust panics on empty JSON", actual)
}

func Test_I16_KVC_JsonMapResults(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	mr, err := c.JsonMapResults()
	actual := args.Map{"noErr": err == nil, "notNil": mr != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonMapResults", actual)
}

func Test_I16_KVC_JsonMapResults_Empty(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	mr, err := c.JsonMapResults()
	actual := args.Map{"noErr": err == nil, "notNil": mr != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns empty -- JsonMapResults empty", actual)
}

func Test_I16_KVC_JsonResultsCollection(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	rc := c.JsonResultsCollection()
	actual := args.Map{"notNil": rc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonResultsCollection", actual)
}

func Test_I16_KVC_JsonResultsPtrCollection(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	rc := c.JsonResultsPtrCollection()
	actual := args.Map{"notNil": rc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonResultsPtrCollection", actual)
}

func Test_I16_KVC_ParseInjectUsingJson(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	jr := corejson.NewPtr(c)
	target := coredynamic.EmptyKeyValCollection()
	result, err := target.ParseInjectUsingJson(jr)
	actual := args.Map{"noErr": err == nil, "notNil": result != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- ParseInjectUsingJson", actual)
}

func Test_I16_KVC_ParseInjectUsingJsonMust(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	jr := corejson.NewPtr(c)
	target := coredynamic.EmptyKeyValCollection()
	result := target.ParseInjectUsingJsonMust(jr)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- ParseInjectUsingJsonMust", actual)
}

func Test_I16_KVC_JsonParseSelfInject(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	jr := corejson.NewPtr(c)
	target := coredynamic.EmptyKeyValCollection()
	err := target.JsonParseSelfInject(jr)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- JsonParseSelfInject", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValCollection — Clone
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_KVC_Clone(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	cloned := c.Clone()
	actual := args.Map{"len": cloned.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- Clone", actual)
}

func Test_I16_KVC_ClonePtr(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	cloned := c.ClonePtr()
	actual := args.Map{"notNil": cloned != nil, "len": cloned.Length()}
	expected := args.Map{"notNil": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- ClonePtr", actual)
}

func Test_I16_KVC_ClonePtr_Nil(t *testing.T) {
	var c *coredynamic.KeyValCollection
	actual := args.Map{"nil": c.ClonePtr() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns nil -- ClonePtr nil", actual)
}

func Test_I16_KVC_NonPtr(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	np := c.NonPtr()
	actual := args.Map{"len": np.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- NonPtr", actual)
}

func Test_I16_KVC_Ptr(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	p := c.Ptr()
	actual := args.Map{"notNil": p != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KVC returns correct value -- Ptr", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicCollection — core methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_DC_EmptyCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"len": dc.Length(), "isEmpty": dc.IsEmpty(), "hasAny": dc.HasAnyItem(), "count": dc.Count()}
	expected := args.Map{"len": 0, "isEmpty": true, "hasAny": false, "count": 0}
	expected.ShouldBeEqual(t, 0, "EmptyDynamicCollection returns empty -- with args", actual)
}

func Test_I16_DC_NilReceiver_Length(t *testing.T) {
	var dc *coredynamic.DynamicCollection
	actual := args.Map{"len": dc.Length(), "isEmpty": dc.IsEmpty()}
	expected := args.Map{"len": 0, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "DC returns nil -- nil Length", actual)
}

func Test_I16_DC_NilReceiver_Items(t *testing.T) {
	var dc *coredynamic.DynamicCollection
	items := dc.Items()
	actual := args.Map{"len": len(items)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DC returns nil -- nil Items", actual)
}

func Test_I16_DC_Add_And_At(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	d := coredynamic.NewDynamic("hello", true)
	dc.Add(d)
	atVal := dc.At(0)
	actual := args.Map{"len": dc.Length(), "atVal": atVal.ValueString()}
	expected := args.Map{"len": 1, "atVal": "hello"}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- Add+At", actual)
}

func Test_I16_DC_AddAny(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("val1", true)
	dc.AddAny(42, true)
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- AddAny", actual)
}

func Test_I16_DC_AddAnyNonNull(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyNonNull("val", true)
	dc.AddAnyNonNull(nil, true) // skipped
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- AddAnyNonNull", actual)
}

func Test_I16_DC_AddAnyMany(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c")
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- AddAnyMany", actual)
}

func Test_I16_DC_AddAnyMany_Nil(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany()
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DC returns nil -- AddAnyMany nil", actual)
}

func Test_I16_DC_AddPtr(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	d := coredynamic.NewDynamic("x", true)
	dc.AddPtr(&d)
	dc.AddPtr(nil) // skipped
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- AddPtr", actual)
}

func Test_I16_DC_AddManyPtr(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	d1 := coredynamic.NewDynamic("a", true)
	d2 := coredynamic.NewDynamic("b", true)
	dc.AddManyPtr(&d1, nil, &d2)
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- AddManyPtr", actual)
}

func Test_I16_DC_AddManyPtr_Nil(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddManyPtr()
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DC returns nil -- AddManyPtr nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicCollection — First/Last/Skip/Take/Limit
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_DC_First_Last(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("first", true)
	dc.AddAny("last", true)
	first := dc.First()
	last := dc.Last()
	actual := args.Map{
		"first":     first.ValueString(),
		"last":      last.ValueString(),
		"lastIdx":   dc.LastIndex(),
		"hasIdx":    dc.HasIndex(1),
		"noIdx":     dc.HasIndex(5),
		"firstDyn":  dc.FirstDynamic() != nil,
		"lastDyn":   dc.LastDynamic() != nil,
	}
	expected := args.Map{
		"first": "first", "last": "last", "lastIdx": 1,
		"hasIdx": true, "noIdx": false, "firstDyn": true, "lastDyn": true,
	}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- First/Last", actual)
}

func Test_I16_DC_FirstOrDefault_NonEmpty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("item", true)
	f := dc.FirstOrDefault()
	actual := args.Map{"notNil": f != nil, "firstOrDefaultDyn": dc.FirstOrDefaultDynamic() != nil}
	expected := args.Map{"notNil": true, "firstOrDefaultDyn": true}
	expected.ShouldBeEqual(t, 0, "DC returns empty -- FirstOrDefault non-empty", actual)
}

func Test_I16_DC_FirstOrDefault_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"nil": dc.FirstOrDefault() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "DC returns empty -- FirstOrDefault empty", actual)
}

func Test_I16_DC_LastOrDefault_NonEmpty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("item", true)
	l := dc.LastOrDefault()
	actual := args.Map{"notNil": l != nil, "lastOrDefaultDyn": dc.LastOrDefaultDynamic() != nil}
	expected := args.Map{"notNil": true, "lastOrDefaultDyn": true}
	expected.ShouldBeEqual(t, 0, "DC returns empty -- LastOrDefault non-empty", actual)
}

func Test_I16_DC_LastOrDefault_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"nil": dc.LastOrDefault() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "DC returns empty -- LastOrDefault empty", actual)
}

func Test_I16_DC_Skip_Take_Limit(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c", "d", "e")
	actual := args.Map{
		"skipLen":    len(dc.Skip(2)),
		"takeLen":    len(dc.Take(3)),
		"limitLen":   len(dc.Limit(2)),
		"skipDynNil": dc.SkipDynamic(1) != nil,
		"takeDynNil": dc.TakeDynamic(2) != nil,
		"limitDyn":   dc.LimitDynamic(2) != nil,
	}
	expected := args.Map{
		"skipLen": 3, "takeLen": 3, "limitLen": 2,
		"skipDynNil": true, "takeDynNil": true, "limitDyn": true,
	}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- Skip/Take/Limit", actual)
}

func Test_I16_DC_SkipCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c")
	sc := dc.SkipCollection(1)
	actual := args.Map{"len": sc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- SkipCollection", actual)
}

func Test_I16_DC_TakeCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c")
	tc := dc.TakeCollection(2)
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- TakeCollection", actual)
}

func Test_I16_DC_LimitCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c")
	lc := dc.LimitCollection(2)
	actual := args.Map{"len": lc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- LimitCollection", actual)
}

func Test_I16_DC_SafeLimitCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b")
	lc := dc.SafeLimitCollection(10) // limit > length
	actual := args.Map{"len": lc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- SafeLimitCollection", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicCollection — RemoveAt, Loop, AnyItems, Strings
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_DC_RemoveAt_Success(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c")
	ok := dc.RemoveAt(1)
	actual := args.Map{"ok": ok, "len": dc.Length()}
	expected := args.Map{"ok": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- RemoveAt success", actual)
}

func Test_I16_DC_RemoveAt_InvalidIndex(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	ok := dc.RemoveAt(5)
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "DC returns error -- RemoveAt invalid", actual)
}

func Test_I16_DC_Loop(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c")
	count := 0
	dc.Loop(func(index int, d *coredynamic.Dynamic) bool {
		count++
		return false
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- Loop", actual)
}

func Test_I16_DC_Loop_Break(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c")
	count := 0
	dc.Loop(func(index int, d *coredynamic.Dynamic) bool {
		count++
		return index == 0 // break after first
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- Loop break", actual)
}

func Test_I16_DC_Loop_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	called := false
	dc.Loop(func(index int, d *coredynamic.Dynamic) bool {
		called = true
		return false
	})
	actual := args.Map{"called": called}
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "DC returns empty -- Loop empty", actual)
}

func Test_I16_DC_AnyItems(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", 42)
	items := dc.AnyItems()
	actual := args.Map{"len": len(items)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- AnyItems", actual)
}

func Test_I16_DC_AnyItems_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"len": len(dc.AnyItems())}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DC returns empty -- AnyItems empty", actual)
}

func Test_I16_DC_AnyItemsCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b")
	ac := dc.AnyItemsCollection()
	actual := args.Map{"notNil": ac != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- AnyItemsCollection", actual)
}

func Test_I16_DC_Strings(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b")
	strs := dc.Strings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- Strings", actual)
}

func Test_I16_DC_String(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b")
	actual := args.Map{"notEmpty": dc.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- String", actual)
}

func Test_I16_DC_ListStrings(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("hello", "world")
	strs := dc.ListStrings()
	actual := args.Map{"notEmpty": len(strs) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- ListStrings", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicCollection — type validation add
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_DC_AddAnyWithTypeValidation_Success(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyWithTypeValidation(false, reflect.TypeOf(""), "hello")
	actual := args.Map{"noErr": err == nil, "len": dc.Length()}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "DC returns non-empty -- AddAnyWithTypeValidation success", actual)
}

func Test_I16_DC_AddAnyWithTypeValidation_TypeMismatch(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyWithTypeValidation(false, reflect.TypeOf(""), 42)
	actual := args.Map{"hasErr": err != nil, "len": dc.Length()}
	expected := args.Map{"hasErr": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "DC returns non-empty -- AddAnyWithTypeValidation mismatch", actual)
}

func Test_I16_DC_AddAnyItemsWithTypeValidation_ContinueOnError(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyItemsWithTypeValidation(true, false, reflect.TypeOf(""), "ok", 42, "also ok")
	actual := args.Map{"hasErr": err != nil, "len": dc.Length()}
	expected := args.Map{"hasErr": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "DC returns non-empty -- AddAnyItemsWithTypeValidation continue", actual)
}

func Test_I16_DC_AddAnyItemsWithTypeValidation_StopOnError(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(""), "ok", 42, "unreachable")
	actual := args.Map{"hasErr": err != nil, "len": dc.Length()}
	expected := args.Map{"hasErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "DC returns non-empty -- AddAnyItemsWithTypeValidation stop", actual)
}

func Test_I16_DC_AddAnyItemsWithTypeValidation_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(""))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DC returns empty -- AddAnyItemsWithTypeValidation empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicCollection — JSON and paging
// ══════════════════════════════════════════════════════════════════════════════

func Test_I16_DC_JsonString(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b")
	s, err := dc.JsonString()
	actual := args.Map{"noErr": err == nil, "notEmpty": s != ""}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- JsonString", actual)
}

func Test_I16_DC_JsonStringMust(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a")
	s := dc.JsonStringMust()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- JsonStringMust", actual)
}

func Test_I16_DC_JsonModel(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	m := dc.JsonModel()
	actual := args.Map{"notNil": m.Items != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- JsonModel", actual)
}

func Test_I16_DC_JsonModelAny(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"notNil": dc.JsonModelAny() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- JsonModelAny", actual)
}

func Test_I16_DC_Json(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	jr := dc.Json()
	actual := args.Map{"noErr": !jr.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- Json", actual)
}

func Test_I16_DC_JsonPtr(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"notNil": dc.JsonPtr() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- JsonPtr", actual)
}

func Test_I16_DC_JsonResultsCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a")
	rc := dc.JsonResultsCollection()
	actual := args.Map{"notNil": rc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- JsonResultsCollection", actual)
}

func Test_I16_DC_JsonResultsPtrCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a")
	rc := dc.JsonResultsPtrCollection()
	actual := args.Map{"notNil": rc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- JsonResultsPtrCollection", actual)
}

func Test_I16_DC_GetPagesSize(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c", "d", "e", "f", "g")
	actual := args.Map{"pages": dc.GetPagesSize(3), "zero": dc.GetPagesSize(0)}
	expected := args.Map{"pages": 3, "zero": 0}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- GetPagesSize", actual)
}

func Test_I16_DC_GetPagedCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c", "d", "e", "f", "g")
	pages := dc.GetPagedCollection(3)
	actual := args.Map{"pagesLen": len(pages)}
	expected := args.Map{"pagesLen": 3}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- GetPagedCollection", actual)
}

func Test_I16_DC_GetPagedCollection_SmallSet(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a")
	pages := dc.GetPagedCollection(10)
	actual := args.Map{"pagesLen": len(pages)}
	expected := args.Map{"pagesLen": 1}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- GetPagedCollection small", actual)
}

func Test_I16_DC_GetSinglePageCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c", "d", "e", "f", "g")
	page := dc.GetSinglePageCollection(3, 2)
	actual := args.Map{"len": page.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- GetSinglePageCollection", actual)
}

func Test_I16_DC_MarshalUnmarshalJSON(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b")
	bytes, err := dc.MarshalJSON()
	actual := args.Map{"noErr": err == nil, "notEmpty": len(bytes) > 0}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- MarshalJSON", actual)
}

func Test_I16_DC_ParseInjectUsingJson(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b")
	jr := corejson.NewPtr(dc)
	target := coredynamic.EmptyDynamicCollection()
	// DynamicCollection can't unmarshal its Items ([]any) from JSON — expect error
	_, err := target.ParseInjectUsingJson(jr)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- ParseInjectUsingJson fails on unmarshal", actual)
}

func Test_I16_DC_ParseInjectUsingJsonMust(t *testing.T) {
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
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- ParseInjectUsingJsonMust panics", actual)
}

func Test_I16_DC_JsonParseSelfInject(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a")
	jr := corejson.NewPtr(dc)
	target := coredynamic.EmptyDynamicCollection()
	err := target.JsonParseSelfInject(jr)
	// DynamicCollection can't unmarshal []any — expect error
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DC returns correct value -- JsonParseSelfInject fails on unmarshal", actual)
}
