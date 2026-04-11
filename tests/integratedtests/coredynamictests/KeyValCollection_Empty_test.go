package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// KeyValCollection — constructors & basic
// ═══════════════════════════════════════════

func Test_KeyValCollection_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyKeyValCollection()
	var nilC *coredynamic.KeyValCollection

	// Act
	actual := args.Map{
		"len":     c.Length(),
		"isEmpty": c.IsEmpty(),
		"hasAny":  c.HasAnyItem(),
		"nilLen":  nilC.Length(),
		"nilStr":  nilC.String(),
	}

	// Assert
	expected := args.Map{
		"len": 0, "isEmpty": true, "hasAny": false,
		"nilLen": 0, "nilStr": "",
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns empty -- Empty", actual)
}

func Test_KeyValCollection_Add(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	c.Add(coredynamic.KeyVal{Key: "b", Value: 2})

	// Act
	actual := args.Map{
		"len":    c.Length(),
		"hasAny": c.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"hasAny": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Add", actual)
}

func Test_KeyValCollection_AddPtr(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.AddPtr(&coredynamic.KeyVal{Key: "a", Value: 1})
	c.AddPtr(nil) // should skip

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- AddPtr", actual)
}

func Test_KeyValCollection_AddMany(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.AddMany(
		coredynamic.KeyVal{Key: "a", Value: 1},
		coredynamic.KeyVal{Key: "b", Value: 2},
	)
	c.AddMany() // empty — should be no-op

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- AddMany", actual)
}

func Test_KeyValCollection_AddManyPtr(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.AddManyPtr(
		&coredynamic.KeyVal{Key: "a", Value: 1},
		nil,
		&coredynamic.KeyVal{Key: "b", Value: 2},
	)
	c.AddManyPtr() // empty — should be no-op

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- AddManyPtr", actual)
}

func Test_KeyValCollection_Items(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	var nilC *coredynamic.KeyValCollection

	// Act
	actual := args.Map{
		"items":    len(c.Items()),
		"nilItems": nilC.Items() == nil,
	}

	// Assert
	expected := args.Map{
		"items": 1,
		"nilItems": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Items", actual)
}

func Test_KeyValCollection_AllKeys_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	keys := c.AllKeys()
	sorted := c.AllKeysSorted()
	empty := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{
		"keysLen":    len(keys),
		"sortedFirst": sorted[0],
		"emptyKeys":  len(empty.AllKeys()),
	}

	// Assert
	expected := args.Map{
		"keysLen": 2,
		"sortedFirst": "a",
		"emptyKeys": 0,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- AllKeys", actual)
}

func Test_KeyValCollection_AllValues_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 42})
	vals := c.AllValues()
	empty := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{
		"valsLen":  len(vals),
		"emptyLen": len(empty.AllValues()),
	}

	// Assert
	expected := args.Map{
		"valsLen": 1,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns non-empty -- AllValues", actual)
}

func Test_KeyValCollection_MapAnyItems(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	c.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	m := c.MapAnyItems()
	empty := coredynamic.EmptyKeyValCollection()
	emptyM := empty.MapAnyItems()

	// Act
	actual := args.Map{
		"mapLen":   m.Length(),
		"hasA":     m.HasKey("a"),
		"emptyLen": emptyM.Length(),
	}

	// Assert
	expected := args.Map{
		"mapLen": 2,
		"hasA": true,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- MapAnyItems", actual)
}

func Test_KeyValCollection_String_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})

	// Act
	actual := args.Map{"notEmpty": c.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- String", actual)
}

func Test_KeyValCollection_Json(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	jsonResult := c.Json()
	jsonPtr := c.JsonPtr()
	model := c.JsonModel()
	modelAny := c.JsonModelAny()
	// Json() uses JsonModel() and now serializes with exported Items payload.

	// Act
	actual := args.Map{
		"jsonOk":     jsonResult.JsonString() != "",
		"ptrNotNil":  jsonPtr != nil,
		"modelNN":    model != nil,
		"modelAnyNN": modelAny != nil,
	}

	// Assert
	expected := args.Map{
		"jsonOk": true, "ptrNotNil": true,
		"modelNN": true, "modelAnyNN": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Json", actual)
}

func Test_KeyValCollection_JsonString(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	js, err := c.JsonString()
	// JsonString now returns a non-empty value from JsonModel().

	// Act
	actual := args.Map{
		"jsEmpty": js == "",
		"errNil":  err == nil,
	}

	// Assert
	expected := args.Map{
		"jsEmpty": false,
		"errNil": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- JsonString", actual)
}

func Test_KeyValCollection_Serialize(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	bytes, err := c.Serialize()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"errNil":   err == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"errNil": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Serialize", actual)
}

func Test_KeyValCollection_Clone(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	cloned := c.Clone()
	clonedP := cloned.Ptr()
	clonedPtr := c.ClonePtr()
	var nilC *coredynamic.KeyValCollection
	nilClone := nilC.ClonePtr()

	// Act
	actual := args.Map{
		"cloneLen":    clonedP.Length(),
		"clonePtrLen": clonedPtr.Length(),
		"nilCloneNil": nilClone == nil,
	}

	// Assert
	expected := args.Map{
		"cloneLen": 1,
		"clonePtrLen": 1,
		"nilCloneNil": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Clone", actual)
}

func Test_KeyValCollection_Paging(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(10)
	for i := 0; i < 5; i++ {
		c.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	pages := c.GetPagesSize(2)
	paged := c.GetPagedCollection(2)
	single := c.GetSinglePageCollection(2, 1)

	// Act
	actual := args.Map{
		"pages":     pages,
		"pagedLen":  len(paged),
		"singleLen": single.Length(),
	}

	// Assert
	expected := args.Map{
		"pages": 3,
		"pagedLen": 3,
		"singleLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Paging", actual)
}

func Test_KeyValCollection_Paging_SmallSet(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	pages := c.GetPagesSize(0)
	paged := c.GetPagedCollection(10)
	single := c.GetSinglePageCollection(10, 1)

	// Act
	actual := args.Map{
		"zeroPage":   pages,
		"pagedSelf":  len(paged),
		"singleSelf": single.Length(),
	}

	// Assert
	expected := args.Map{
		"zeroPage": 0,
		"pagedSelf": 1,
		"singleSelf": 1,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Paging small set", actual)
}

func Test_KeyValCollection_JsonMapResults(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: "hello"})
	mr, err := c.JsonMapResults()
	emptyC := coredynamic.EmptyKeyValCollection()
	emptyMR, emptyErr := emptyC.JsonMapResults()

	// Act
	actual := args.Map{
		"mrNotNil":      mr != nil,
		"errNil":        err == nil,
		"emptyMRNotNil": emptyMR != nil,
		"emptyErrNil":   emptyErr == nil,
	}

	// Assert
	expected := args.Map{
		"mrNotNil": true, "errNil": true,
		"emptyMRNotNil": true, "emptyErrNil": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- JsonMapResults", actual)
}

func Test_KeyValCollection_JsonResultsCollection(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: "hello"})
	rc := c.JsonResultsCollection()
	rpc := c.JsonResultsPtrCollection()

	// Act
	actual := args.Map{
		"rcNotNil":  rc != nil,
		"rpcNotNil": rpc != nil,
	}

	// Assert
	expected := args.Map{
		"rcNotNil": true,
		"rpcNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- JsonResultsCollection", actual)
}

func Test_KeyValCollection_ParseJson(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	jsonResult := c.Json()
	jsonPtr := &jsonResult

	target := coredynamic.EmptyKeyValCollection()
	parsed, err := target.ParseInjectUsingJson(jsonPtr)

	// Act
	actual := args.Map{
		"parsedNotNil": parsed != nil,
		"errNil":       err == nil,
	}

	// Assert
	expected := args.Map{
		"parsedNotNil": true,
		"errNil": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- ParseInjectUsingJson", actual)
}

func Test_KeyValCollection_JsonParseSelfInject(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	jsonResult := c.Json()
	jsonPtr := &jsonResult

	target := coredynamic.EmptyKeyValCollection()
	err := target.JsonParseSelfInject(jsonPtr)

	// Act
	actual := args.Map{"errNil": err == nil}

	// Assert
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- JsonParseSelfInject", actual)
}

func Test_KeyValCollection_NonPtrPtr(t *testing.T) {
	// Arrange
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	nonPtr := c.NonPtr()
	nonPtrPtr := nonPtr.Ptr()
	ptr := c.Ptr()

	// Act
	actual := args.Map{
		"nonPtrLen": nonPtrPtr.Length(),
		"ptrLen":    ptr.Length(),
		"same":      ptr == c,
	}

	// Assert
	expected := args.Map{
		"nonPtrLen": 1,
		"ptrLen": 1,
		"same": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- NonPtr/Ptr", actual)
}

// ═══════════════════════════════════════════
// DynamicCollection — constructors & basic
// ═══════════════════════════════════════════

func Test_DynamicCollection_Empty_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	var nilDC *coredynamic.DynamicCollection

	// Act
	actual := args.Map{
		"len":     dc.Length(),
		"count":   dc.Count(),
		"isEmpty": dc.IsEmpty(),
		"hasAny":  dc.HasAnyItem(),
		"lastIdx": dc.LastIndex(),
		"nilLen":  nilDC.Length(),
		"nilEmpty": nilDC.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"len": 0, "count": 0, "isEmpty": true, "hasAny": false,
		"lastIdx": -1, "nilLen": 0, "nilEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns empty -- Empty", actual)
}

func Test_DynamicCollection_Add_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.Add(coredynamic.NewDynamic("a", true))
	dc.AddPtr(coredynamic.NewDynamicPtr("b", true))
	dc.AddPtr(nil) // skip nil
	dc.AddAny("c", true)
	dc.AddAnyNonNull("d", true)
	dc.AddAnyNonNull(nil, true) // skip nil

	// Act
	actual := args.Map{
		"len":    dc.Length(),
		"hasIdx": dc.HasIndex(3),
		"noIdx":  dc.HasIndex(10),
	}

	// Assert
	expected := args.Map{
		"len": 4,
		"hasIdx": true,
		"noIdx": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Add", actual)
}

func Test_DynamicCollection_AddAnyMany_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnyMany("a", "b", "c")
	dc.AddAnyMany() // nil — no-op

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AddAnyMany", actual)
}

func Test_DynamicCollection_AddManyPtr_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddManyPtr(
		coredynamic.NewDynamicPtr("a", true),
		nil,
		coredynamic.NewDynamicPtr("b", true),
	)
	dc.AddManyPtr() // nil — no-op

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AddManyPtr", actual)
}

func Test_DynamicCollection_FirstLastOrDefault(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAny("first", true)
	dc.AddAny("last", true)
	emptyDC := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{
		"first":          dc.First().Value(),
		"last":           dc.Last().Value(),
		"firstDyn":       dc.FirstDynamic() != nil,
		"lastDyn":        dc.LastDynamic() != nil,
		"firstOrDef":     dc.FirstOrDefault() != nil,
		"lastOrDef":      dc.LastOrDefault() != nil,
		"firstOrDefDyn":  dc.FirstOrDefaultDynamic() != nil,
		"lastOrDefDyn":   dc.LastOrDefaultDynamic() != nil,
		"emptyFirstDef":  emptyDC.FirstOrDefault() == nil,
		"emptyLastDef":   emptyDC.LastOrDefault() == nil,
	}

	// Assert
	expected := args.Map{
		"first": "first", "last": "last",
		"firstDyn": true, "lastDyn": true,
		"firstOrDef": true, "lastOrDef": true,
		"firstOrDefDyn": true, "lastOrDefDyn": true,
		"emptyFirstDef": true, "emptyLastDef": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- FirstLast", actual)
}

func Test_DynamicCollection_At_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAny("hello", true)
	d := dc.At(0)

	// Act
	actual := args.Map{"val": d.Value()}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- At", actual)
}

func Test_DynamicCollection_Items_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAny("a", true)
	var nilDC *coredynamic.DynamicCollection

	// Act
	actual := args.Map{
		"itemsLen": len(dc.Items()),
		"nilItems": len(nilDC.Items()),
	}

	// Assert
	expected := args.Map{
		"itemsLen": 1,
		"nilItems": 0,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Items", actual)
}

func Test_DynamicCollection_SkipTakeLimitSlice(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(10)
	dc.AddAnyMany("a", "b", "c", "d", "e")

	// Act
	actual := args.Map{
		"skipLen":      len(dc.Skip(2)),
		"skipDynLen":   dc.SkipDynamic(2) != nil,
		"skipCol":      dc.SkipCollection(2).Length(),
		"takeLen":      len(dc.Take(2)),
		"takeDynLen":   dc.TakeDynamic(2) != nil,
		"takeCol":      dc.TakeCollection(2).Length(),
		"limitLen":     len(dc.Limit(3)),
		"limitDynLen":  dc.LimitDynamic(3) != nil,
		"limitCol":     dc.LimitCollection(3).Length(),
		"safeLimitCol": dc.SafeLimitCollection(100).Length(),
	}

	// Assert
	expected := args.Map{
		"skipLen": 3, "skipDynLen": true, "skipCol": 3,
		"takeLen": 2, "takeDynLen": true, "takeCol": 2,
		"limitLen": 3, "limitDynLen": true, "limitCol": 3,
		"safeLimitCol": 5,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Skip/Take/Limit", actual)
}

func Test_DynamicCollection_RemoveAt_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnyMany("a", "b", "c")
	ok := dc.RemoveAt(1)
	fail := dc.RemoveAt(100)

	// Act
	actual := args.Map{
		"ok":     ok,
		"fail":   fail,
		"newLen": dc.Length(),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"fail": false,
		"newLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- RemoveAt", actual)
}

func Test_DynamicCollection_AnyItems_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnyMany("a", "b")
	items := dc.AnyItems()
	emptyDC := coredynamic.EmptyDynamicCollection()
	emptyItems := emptyDC.AnyItems()

	// Act
	actual := args.Map{
		"itemsLen": len(items),
		"emptyLen": len(emptyItems),
	}

	// Assert
	expected := args.Map{
		"itemsLen": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AnyItems", actual)
}

func Test_DynamicCollection_AnyItemsCollection_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnyMany("a", "b")
	ac := dc.AnyItemsCollection()
	emptyDC := coredynamic.EmptyDynamicCollection()
	emptyAC := emptyDC.AnyItemsCollection()

	// Act
	actual := args.Map{
		"acLen":    ac.Length(),
		"emptyLen": emptyAC.Length(),
	}

	// Assert
	expected := args.Map{
		"acLen": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AnyItemsCollection", actual)
}

func Test_DynamicCollection_ListStrings_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnyMany("hello", "world")
	strs := dc.ListStrings()
	strsPtr := dc.ListStringsPtr()

	// Act
	actual := args.Map{
		"strsLen":    len(strs),
		"strsPtrLen": len(strsPtr),
	}

	// Assert
	expected := args.Map{
		"strsLen": 2,
		"strsPtrLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- ListStrings", actual)
}

func Test_DynamicCollection_Strings_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnyMany("a", "b")
	strs := dc.Strings()
	str := dc.String()
	emptyDC := coredynamic.EmptyDynamicCollection()
	emptyStrs := emptyDC.Strings()

	// Act
	actual := args.Map{
		"strsLen":   len(strs),
		"strNotEmpty": str != "",
		"emptyLen":  len(emptyStrs),
	}

	// Assert
	expected := args.Map{
		"strsLen": 2,
		"strNotEmpty": true,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Strings", actual)
}

func Test_DynamicCollection_Loop_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnyMany("a", "b", "c")
	count := 0
	dc.Loop(func(index int, dynamicItem *coredynamic.Dynamic) (isBreak bool) {
		count++
		return false
	})
	emptyDC := coredynamic.EmptyDynamicCollection()
	emptyCount := 0
	emptyDC.Loop(func(index int, dynamicItem *coredynamic.Dynamic) (isBreak bool) {
		emptyCount++
		return false
	})

	// Act
	actual := args.Map{
		"count": count,
		"emptyCount": emptyCount,
	}

	// Assert
	expected := args.Map{
		"count": 3,
		"emptyCount": 0,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Loop", actual)
}

func Test_DynamicCollection_Loop_Break_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnyMany("a", "b", "c")
	count := 0
	dc.Loop(func(index int, dynamicItem *coredynamic.Dynamic) (isBreak bool) {
		count++
		return true // break on first
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Loop break", actual)
}

func Test_DynamicCollection_Json_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAny("hello", true)
	jsonResult := dc.Json()
	jsonPtr := dc.JsonPtr()
	model := dc.JsonModel()
	modelAny := dc.JsonModelAny()
	js, jsErr := dc.JsonString()
	jsMust := dc.JsonStringMust()

	// Act
	actual := args.Map{
		"jsonOk":     jsonResult.HasError() == false,
		"ptrNotNil":  jsonPtr != nil,
		"modelItems": len(model.Items) > 0,
		"modelAnyNN": modelAny != nil,
		"jsNotEmpty": js != "",
		"jsErrNil":   jsErr == nil,
		"mustNE":     jsMust != "",
	}

	// Assert
	expected := args.Map{
		"jsonOk": true, "ptrNotNil": true,
		"modelItems": true, "modelAnyNN": true,
		"jsNotEmpty": true, "jsErrNil": true, "mustNE": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Json", actual)
}

func Test_DynamicCollection_JsonResultsCollection_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAny("hello", true)
	rc := dc.JsonResultsCollection()
	rpc := dc.JsonResultsPtrCollection()
	emptyDC := coredynamic.EmptyDynamicCollection()
	emptyRC := emptyDC.JsonResultsCollection()

	// Act
	actual := args.Map{
		"rcNotNil":    rc != nil,
		"rpcNotNil":   rpc != nil,
		"emptyRCNotNil": emptyRC != nil,
	}

	// Assert
	expected := args.Map{
		"rcNotNil": true,
		"rpcNotNil": true,
		"emptyRCNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- JsonResultsCollection", actual)
}

func Test_DynamicCollection_Paging(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(10)
	dc.AddAnyMany("a", "b", "c", "d", "e")
	pages := dc.GetPagesSize(2)
	paged := dc.GetPagedCollection(2)
	single := dc.GetSinglePageCollection(2, 1)

	// Act
	actual := args.Map{
		"pages":     pages,
		"pagedLen":  len(paged),
		"singleLen": single.Length(),
	}

	// Assert
	expected := args.Map{
		"pages": 3,
		"pagedLen": 3,
		"singleLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Paging", actual)
}

func Test_DynamicCollection_Paging_SmallSet(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAny("a", true)
	pages := dc.GetPagesSize(0)
	paged := dc.GetPagedCollection(10)
	single := dc.GetSinglePageCollection(10, 1)

	// Act
	actual := args.Map{
		"zeroPage":   pages,
		"pagedSelf":  len(paged),
		"singleSelf": single.Length(),
	}

	// Assert
	expected := args.Map{
		"zeroPage": 0,
		"pagedSelf": 1,
		"singleSelf": 1,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Paging small set", actual)
}

func Test_DynamicCollection_ParseJson(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAny("hello", true)
	jsonResult := dc.Json()
	jsonPtr := &jsonResult

	target := coredynamic.EmptyDynamicCollection()
	parsed, err := target.ParseInjectUsingJson(jsonPtr)

	// Act
	actual := args.Map{
		"parsedNotNil": parsed != nil,
		"errNil":       err == nil,
	}

	// Assert
	expected := args.Map{
		"parsedNotNil": true,
		"errNil": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- ParseInjectUsingJson", actual)
}

func Test_DynamicCollection_JsonParseSelfInject_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAny("hello", true)
	jsonResult := dc.Json()
	jsonPtr := &jsonResult

	target := coredynamic.EmptyDynamicCollection()
	err := target.JsonParseSelfInject(jsonPtr)

	// Act
	actual := args.Map{"errNil": err == nil}

	// Assert
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- JsonParseSelfInject", actual)
}

// ═══════════════════════════════════════════
// MapAnyItems — constructors & basic
// ═══════════════════════════════════════════

func Test_MapAnyItems_Constructors(t *testing.T) {
	// Arrange
	empty := coredynamic.EmptyMapAnyItems()
	m := coredynamic.NewMapAnyItems(5)
	fromItems := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	fromItemsEmpty := coredynamic.NewMapAnyItemsUsingItems(nil)
	var nilM *coredynamic.MapAnyItems

	// Act
	actual := args.Map{
		"emptyLen":      empty.Length(),
		"mLen":          m.Length(),
		"fromItemsLen":  fromItems.Length(),
		"fromEmptyLen":  fromItemsEmpty.Length(),
		"nilLen":        nilM.Length(),
		"isEmpty":       empty.IsEmpty(),
		"hasAny":        fromItems.HasAnyItem(),
		"nilHasKey":     nilM.HasKey("x"),
	}

	// Assert
	expected := args.Map{
		"emptyLen": 0, "mLen": 0, "fromItemsLen": 1,
		"fromEmptyLen": 0, "nilLen": 0, "isEmpty": true,
		"hasAny": true, "nilHasKey": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Constructors", actual)
}

func Test_MapAnyItems_AddSet(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(5)
	isNew1 := m.Add("a", 1)
	isNew2 := m.Add("a", 2) // overwrite
	isNew3 := m.Set("b", 3)

	// Act
	actual := args.Map{
		"isNew1": isNew1, "isNew2": isNew2, "isNew3": isNew3,
		"len": m.Length(), "hasA": m.HasKey("a"), "hasB": m.HasKey("b"),
	}

	// Assert
	expected := args.Map{
		"isNew1": true, "isNew2": false, "isNew3": true,
		"len": 2, "hasA": true, "hasB": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Add/Set", actual)
}

func Test_MapAnyItems_Get_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 42})
	val, has := m.Get("a")
	_, hasMissing := m.Get("z")
	getVal := m.GetValue("a")
	getMissing := m.GetValue("z")

	// Act
	actual := args.Map{
		"val": val, "has": has, "hasMissing": hasMissing,
		"getVal": getVal, "getMissing": getMissing == nil,
	}

	// Assert
	expected := args.Map{
		"val": 42, "has": true, "hasMissing": false,
		"getVal": 42, "getMissing": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Get", actual)
}

func Test_MapAnyItems_AllKeys_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"b": 2, "a": 1})
	keys := m.AllKeys()
	sorted := m.AllKeysSorted()
	vals := m.AllValues()
	emptyM := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{
		"keysLen":    len(keys),
		"sortedFirst": sorted[0],
		"valsLen":    len(vals),
		"emptyKeys":  len(emptyM.AllKeys()),
		"emptyVals":  len(emptyM.AllValues()),
	}

	// Assert
	expected := args.Map{
		"keysLen": 2, "sortedFirst": "a", "valsLen": 2,
		"emptyKeys": 0, "emptyVals": 0,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- AllKeys", actual)
}

func Test_MapAnyItems_AddMapResult_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(5)
	m.Add("a", 1)
	m.AddMapResult(map[string]any{"b": 2, "c": 3})
	m.AddMapResult(nil) // no-op

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- AddMapResult", actual)
}

func Test_MapAnyItems_AddManyMapResults(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(5)
	m.AddManyMapResultsUsingOption(true, map[string]any{"a": 1}, map[string]any{"b": 2})
	m.AddManyMapResultsUsingOption(true) // empty — no-op

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- AddManyMapResults", actual)
}

func Test_MapAnyItems_GetNewMapUsingKeys_FromKeyValCollectionEmpt(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2, "c": 3})
	sub := m.GetNewMapUsingKeys(false, "a", "c")
	emptyKeys := m.GetNewMapUsingKeys(false)

	// Act
	actual := args.Map{
		"subLen":   sub.Length(),
		"emptyLen": emptyKeys.Length(),
	}

	// Assert
	expected := args.Map{
		"subLen": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetNewMapUsingKeys", actual)
}

func Test_MapAnyItems_Json(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	jsonResult := m.Json()
	jsonPtr := m.JsonPtr()
	model := m.JsonModel()
	modelAny := m.JsonModelAny()
	js, jsErr := m.JsonString()
	jsMust := m.JsonStringMust()

	// Act
	actual := args.Map{
		"jsonOk":     jsonResult.JsonString() != "",
		"ptrNotNil":  jsonPtr != nil,
		"modelNN":    model != nil,
		"modelAnyNN": modelAny != nil,
		"jsNotEmpty": js != "",
		"jsErrNil":   jsErr == nil,
		"mustNE":     jsMust != "",
	}

	// Assert
	expected := args.Map{
		"jsonOk": true, "ptrNotNil": true,
		"modelNN": true, "modelAnyNN": true,
		"jsNotEmpty": true, "jsErrNil": true, "mustNE": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Json", actual)
}

func Test_MapAnyItems_Strings(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	strs := m.Strings()
	str := m.String()

	// Act
	actual := args.Map{
		"strsLen":    len(strs),
		"strNotEmpty": str != "",
	}

	// Assert
	expected := args.Map{
		"strsLen": 1,
		"strNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Strings", actual)
}

func Test_MapAnyItems_ClearDisposeDeepClear(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	m.Clear()
	cleared := m.Length()
	m.Add("x", 1)
	m.DeepClear()
	deepCleared := m.Length()

	m2 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m2.Dispose()
	disposed := m2.Items == nil

	var nilM *coredynamic.MapAnyItems
	nilM.Clear()    // should not panic
	nilM.DeepClear() // should not panic
	nilM.Dispose()   // should not panic

	// Act
	actual := args.Map{
		"cleared":     cleared,
		"deepCleared": deepCleared,
		"disposed":    disposed,
	}

	// Assert
	expected := args.Map{
		"cleared": 0,
		"deepCleared": 0,
		"disposed": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Clear/DeepClear/Dispose", actual)
}

func Test_MapAnyItems_IsEqual(t *testing.T) {
	// Arrange
	m1 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	m2 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	m3 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 3})
	m4 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	var nilM *coredynamic.MapAnyItems

	// Act
	actual := args.Map{
		"equal":       m1.IsEqual(m2),
		"notEqual":    m1.IsEqual(m3),
		"diffLen":     m1.IsEqual(m4),
		"nilBothEq":   nilM.IsEqual(nil),
		"nilOneNotEq": nilM.IsEqual(m1),
		"rawEqual":    m1.IsEqualRaw(map[string]any{"a": 1, "b": 2}),
		"rawNotEqual": m1.IsEqualRaw(map[string]any{"a": 1, "b": 3}),
		"rawMissingKey": m1.IsEqualRaw(map[string]any{"a": 1, "c": 2}),
	}

	// Assert
	expected := args.Map{
		"equal": true, "notEqual": false, "diffLen": false,
		"nilBothEq": true, "nilOneNotEq": false,
		"rawEqual": true, "rawNotEqual": false, "rawMissingKey": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- IsEqual", actual)
}

func Test_MapAnyItems_ClonePtr(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	cloned, err := m.ClonePtr()
	var nilM *coredynamic.MapAnyItems
	nilClone, nilErr := nilM.ClonePtr()

	// Act
	actual := args.Map{
		"clonedLen":    cloned.Length(),
		"errNil":       err == nil,
		"nilCloneNil":  nilClone == nil,
		"nilErrNotNil": nilErr != nil,
	}

	// Assert
	expected := args.Map{
		"clonedLen": 1, "errNil": true,
		"nilCloneNil": true, "nilErrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- ClonePtr", actual)
}

func Test_MapAnyItems_MapAnyItemsSelf(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	self := m.MapAnyItems()

	// Act
	actual := args.Map{"same": self == m}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- self-reference", actual)
}

func Test_MapAnyItems_Paging(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(10)
	for i := 0; i < 5; i++ {
		m.Add("k"+string(rune('a'+i)), i)
	}
	pages := m.GetPagesSize(2)
	paged := m.GetPagedCollection(2)

	// Act
	actual := args.Map{
		"pages":    pages,
		"pagedLen": len(paged),
	}

	// Assert
	expected := args.Map{
		"pages": 3,
		"pagedLen": 3,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Paging", actual)
}

func Test_MapAnyItems_Paging_SmallSet(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	pages := m.GetPagesSize(0)
	paged := m.GetPagedCollection(10)

	// Act
	actual := args.Map{
		"zeroPage":  pages,
		"pagedSelf": len(paged),
	}

	// Assert
	expected := args.Map{
		"zeroPage": 0,
		"pagedSelf": 1,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Paging small set", actual)
}

func Test_MapAnyItems_ParseJson(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	jsonResult := m.Json()
	jsonPtr := &jsonResult

	target := coredynamic.EmptyMapAnyItems()
	parsed, err := target.ParseInjectUsingJson(jsonPtr)

	// Act
	actual := args.Map{
		"parsedNotNil": parsed != nil,
		"errNil":       err == nil,
	}

	// Assert
	expected := args.Map{
		"parsedNotNil": true,
		"errNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- ParseInjectUsingJson", actual)
}

func Test_MapAnyItems_JsonParseSelfInject(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	jsonResult := m.Json()
	jsonPtr := &jsonResult

	target := coredynamic.EmptyMapAnyItems()
	err := target.JsonParseSelfInject(jsonPtr)

	// Act
	actual := args.Map{"errNil": err == nil}

	// Assert
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonParseSelfInject", actual)
}

func Test_MapAnyItems_JsonResultsCollections(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "hello"})
	rc := m.JsonResultsCollection()
	rpc := m.JsonResultsPtrCollection()
	mr, mrErr := m.JsonMapResults()
	emptyM := coredynamic.EmptyMapAnyItems()
	emptyRC := emptyM.JsonResultsCollection()
	emptyMR, emptyMRErr := emptyM.JsonMapResults()

	// Act
	actual := args.Map{
		"rcNotNil":      rc != nil,
		"rpcNotNil":     rpc != nil,
		"mrNotNil":      mr != nil,
		"mrErrNil":      mrErr == nil,
		"emptyRCNotNil": emptyRC != nil,
		"emptyMRNotNil": emptyMR != nil,
		"emptyMRErrNil": emptyMRErr == nil,
	}

	// Assert
	expected := args.Map{
		"rcNotNil": true, "rpcNotNil": true,
		"mrNotNil": true, "mrErrNil": true,
		"emptyRCNotNil": true, "emptyMRNotNil": true, "emptyMRErrNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonResultsCollections", actual)
}

func Test_MapAnyItems_JsonResultOfKey(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "hello"})
	found := m.JsonResultOfKey("a")
	missing := m.JsonResultOfKey("z")

	// Act
	actual := args.Map{
		"foundHasBytes":  len(found.Bytes) > 0,
		"missingHasErr":  missing.HasError(),
	}

	// Assert
	expected := args.Map{
		"foundHasBytes": true,
		"missingHasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonResultOfKey", actual)
}

func Test_MapAnyItems_JsonResultOfKeys(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "hello", "b": 42})
	results := m.JsonResultOfKeys("a", "b")
	emptyResults := m.JsonResultOfKeys()

	// Act
	actual := args.Map{
		"resultsNN": results != nil,
		"emptyNN":   emptyResults != nil,
	}

	// Assert
	expected := args.Map{
		"resultsNN": true,
		"emptyNN": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonResultOfKeys", actual)
}

func Test_MapAnyItems_AddJsonResultPtr(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(5)
	m.AddJsonResultPtr("a", nil) // nil — should skip
	jr := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"x": 1}).JsonPtr()
	m.AddJsonResultPtr("b", jr) // non-nil

	// Act
	actual := args.Map{
		"noA": m.HasKey("a"),
		"hasB": m.HasKey("b"),
	}

	// Assert
	expected := args.Map{
		"noA": false,
		"hasB": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- AddJsonResultPtr", actual)
}

func Test_MapAnyItems_RawMapStringAnyDiff(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	diff := m.RawMapStringAnyDiff()
	var nilM *coredynamic.MapAnyItems
	nilDiff := nilM.RawMapStringAnyDiff()

	// Act
	actual := args.Map{
		"diffLen":    len(diff),
		"nilDiffLen": len(nilDiff),
	}

	// Assert
	expected := args.Map{
		"diffLen": 1,
		"nilDiffLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- RawMapStringAnyDiff", actual)
}

func Test_MapAnyItems_Deserialize(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"name": "test"})
	var result string
	err := m.Deserialize("name", &result)
	errMissing := m.Deserialize("missing", &result)

	// Act
	actual := args.Map{
		"result":     result,
		"errNil":     err == nil,
		"missingErr": errMissing != nil,
	}

	// Assert
	expected := args.Map{
		"result": "test",
		"errNil": true,
		"missingErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Deserialize", actual)
}

func Test_MapAnyItems_GetFieldsMap(t *testing.T) {
	// Arrange
	inner := map[string]any{"x": 1}
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"data": inner})
	_, _, found := m.GetFieldsMap("data")
	_, notFound := m.GetSafeFieldsMap("missing")

	// Act
	actual := args.Map{
		"found":    found,
		"notFound": notFound,
	}

	// Assert
	expected := args.Map{
		"found": true,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetFieldsMap", actual)
}

func Test_MapAnyItems_GetManyItemsRefs_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(5)
	err := m.GetManyItemsRefs() // empty — should return nil

	// Act
	actual := args.Map{"errNil": err == nil}

	// Assert
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- GetManyItemsRefs empty", actual)
}

func Test_MapAnyItems_HasAnyChanges(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	changed := m.HasAnyChanges(true, map[string]any{"a": 2})
	notChanged := m.HasAnyChanges(true, map[string]any{"a": 1})

	// Act
	actual := args.Map{
		"changed":    changed,
		"notChanged": notChanged,
	}

	// Assert
	expected := args.Map{
		"changed": true,
		"notChanged": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- HasAnyChanges", actual)
}
