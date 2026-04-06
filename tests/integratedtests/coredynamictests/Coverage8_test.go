package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// KeyValCollection — constructors & basic
// ═══════════════════════════════════════════

func Test_Cov8_KeyValCollection_Empty(t *testing.T) {
	c := coredynamic.EmptyKeyValCollection()
	var nilC *coredynamic.KeyValCollection
	actual := args.Map{
		"len":     c.Length(),
		"isEmpty": c.IsEmpty(),
		"hasAny":  c.HasAnyItem(),
		"nilLen":  nilC.Length(),
		"nilStr":  nilC.String(),
	}
	expected := args.Map{
		"len": 0, "isEmpty": true, "hasAny": false,
		"nilLen": 0, "nilStr": "",
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns empty -- Empty", actual)
}

func Test_Cov8_KeyValCollection_Add(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	c.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	actual := args.Map{
		"len":    c.Length(),
		"hasAny": c.HasAnyItem(),
	}
	expected := args.Map{"len": 2, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Add", actual)
}

func Test_Cov8_KeyValCollection_AddPtr(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.AddPtr(&coredynamic.KeyVal{Key: "a", Value: 1})
	c.AddPtr(nil) // should skip
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- AddPtr", actual)
}

func Test_Cov8_KeyValCollection_AddMany(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.AddMany(
		coredynamic.KeyVal{Key: "a", Value: 1},
		coredynamic.KeyVal{Key: "b", Value: 2},
	)
	c.AddMany() // empty — should be no-op
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- AddMany", actual)
}

func Test_Cov8_KeyValCollection_AddManyPtr(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.AddManyPtr(
		&coredynamic.KeyVal{Key: "a", Value: 1},
		nil,
		&coredynamic.KeyVal{Key: "b", Value: 2},
	)
	c.AddManyPtr() // empty — should be no-op
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- AddManyPtr", actual)
}

func Test_Cov8_KeyValCollection_Items(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	var nilC *coredynamic.KeyValCollection
	actual := args.Map{
		"items":    len(c.Items()),
		"nilItems": nilC.Items() == nil,
	}
	expected := args.Map{"items": 1, "nilItems": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Items", actual)
}

func Test_Cov8_KeyValCollection_AllKeys(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	keys := c.AllKeys()
	sorted := c.AllKeysSorted()
	empty := coredynamic.EmptyKeyValCollection()
	actual := args.Map{
		"keysLen":    len(keys),
		"sortedFirst": sorted[0],
		"emptyKeys":  len(empty.AllKeys()),
	}
	expected := args.Map{"keysLen": 2, "sortedFirst": "a", "emptyKeys": 0}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- AllKeys", actual)
}

func Test_Cov8_KeyValCollection_AllValues(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 42})
	vals := c.AllValues()
	empty := coredynamic.EmptyKeyValCollection()
	actual := args.Map{
		"valsLen":  len(vals),
		"emptyLen": len(empty.AllValues()),
	}
	expected := args.Map{"valsLen": 1, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns non-empty -- AllValues", actual)
}

func Test_Cov8_KeyValCollection_MapAnyItems(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	c.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	m := c.MapAnyItems()
	empty := coredynamic.EmptyKeyValCollection()
	emptyM := empty.MapAnyItems()
	actual := args.Map{
		"mapLen":   m.Length(),
		"hasA":     m.HasKey("a"),
		"emptyLen": emptyM.Length(),
	}
	expected := args.Map{"mapLen": 2, "hasA": true, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- MapAnyItems", actual)
}

func Test_Cov8_KeyValCollection_String(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	actual := args.Map{"notEmpty": c.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- String", actual)
}

func Test_Cov8_KeyValCollection_Json(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	jsonResult := c.Json()
	jsonPtr := c.JsonPtr()
	model := c.JsonModel()
	modelAny := c.JsonModelAny()
	// Json() uses JsonModel() and now serializes with exported Items payload.
	actual := args.Map{
		"jsonOk":     jsonResult.JsonString() != "",
		"ptrNotNil":  jsonPtr != nil,
		"modelNN":    model != nil,
		"modelAnyNN": modelAny != nil,
	}
	expected := args.Map{
		"jsonOk": true, "ptrNotNil": true,
		"modelNN": true, "modelAnyNN": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Json", actual)
}

func Test_Cov8_KeyValCollection_JsonString(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	js, err := c.JsonString()
	// JsonString now returns a non-empty value from JsonModel().
	actual := args.Map{
		"jsEmpty": js == "",
		"errNil":  err == nil,
	}
	expected := args.Map{"jsEmpty": false, "errNil": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- JsonString", actual)
}

func Test_Cov8_KeyValCollection_Serialize(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	bytes, err := c.Serialize()
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"errNil":   err == nil,
	}
	expected := args.Map{"hasBytes": true, "errNil": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Serialize", actual)
}
func Test_Cov8_KeyValCollection_Paging(t *testing.T) {
	c := coredynamic.NewKeyValCollection(10)
	for i := 0; i < 5; i++ {
		c.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	pages := c.GetPagesSize(2)
	paged := c.GetPagedCollection(2)
	single := c.GetSinglePageCollection(2, 1)
	actual := args.Map{
		"pages":     pages,
		"pagedLen":  len(paged),
		"singleLen": single.Length(),
	}
	expected := args.Map{"pages": 3, "pagedLen": 3, "singleLen": 2}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Paging", actual)
}

func Test_Cov8_KeyValCollection_Paging_SmallSet(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	pages := c.GetPagesSize(0)
	paged := c.GetPagedCollection(10)
	single := c.GetSinglePageCollection(10, 1)
	actual := args.Map{
		"zeroPage":   pages,
		"pagedSelf":  len(paged),
		"singleSelf": single.Length(),
	}
	expected := args.Map{"zeroPage": 0, "pagedSelf": 1, "singleSelf": 1}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- Paging small set", actual)
}

func Test_Cov8_KeyValCollection_JsonMapResults(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: "hello"})
	mr, err := c.JsonMapResults()
	emptyC := coredynamic.EmptyKeyValCollection()
	emptyMR, emptyErr := emptyC.JsonMapResults()
	actual := args.Map{
		"mrNotNil":      mr != nil,
		"errNil":        err == nil,
		"emptyMRNotNil": emptyMR != nil,
		"emptyErrNil":   emptyErr == nil,
	}
	expected := args.Map{
		"mrNotNil": true, "errNil": true,
		"emptyMRNotNil": true, "emptyErrNil": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- JsonMapResults", actual)
}

func Test_Cov8_KeyValCollection_JsonResultsCollection(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: "hello"})
	rc := c.JsonResultsCollection()
	rpc := c.JsonResultsPtrCollection()
	actual := args.Map{
		"rcNotNil":  rc != nil,
		"rpcNotNil": rpc != nil,
	}
	expected := args.Map{"rcNotNil": true, "rpcNotNil": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- JsonResultsCollection", actual)
}

func Test_Cov8_KeyValCollection_ParseJson(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	jsonResult := c.Json()
	jsonPtr := &jsonResult

	target := coredynamic.EmptyKeyValCollection()
	parsed, err := target.ParseInjectUsingJson(jsonPtr)
	actual := args.Map{
		"parsedNotNil": parsed != nil,
		"errNil":       err == nil,
	}
	expected := args.Map{"parsedNotNil": true, "errNil": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- ParseInjectUsingJson", actual)
}

func Test_Cov8_KeyValCollection_JsonParseSelfInject(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	jsonResult := c.Json()
	jsonPtr := &jsonResult

	target := coredynamic.EmptyKeyValCollection()
	err := target.JsonParseSelfInject(jsonPtr)
	actual := args.Map{"errNil": err == nil}
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- JsonParseSelfInject", actual)
}

func Test_Cov8_KeyValCollection_NonPtrPtr(t *testing.T) {
	c := coredynamic.NewKeyValCollection(5)
	c.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	nonPtr := c.NonPtr()
	nonPtrPtr := nonPtr.Ptr()
	ptr := c.Ptr()
	actual := args.Map{
		"nonPtrLen": nonPtrPtr.Length(),
		"ptrLen":    ptr.Length(),
		"same":      ptr == c,
	}
	expected := args.Map{"nonPtrLen": 1, "ptrLen": 1, "same": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection returns correct value -- NonPtr/Ptr", actual)
}

// ═══════════════════════════════════════════
// DynamicCollection — constructors & basic
// ═══════════════════════════════════════════

func Test_Cov8_DynamicCollection_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	var nilDC *coredynamic.DynamicCollection
	actual := args.Map{
		"len":     dc.Length(),
		"count":   dc.Count(),
		"isEmpty": dc.IsEmpty(),
		"hasAny":  dc.HasAnyItem(),
		"lastIdx": dc.LastIndex(),
		"nilLen":  nilDC.Length(),
		"nilEmpty": nilDC.IsEmpty(),
	}
	expected := args.Map{
		"len": 0, "count": 0, "isEmpty": true, "hasAny": false,
		"lastIdx": -1, "nilLen": 0, "nilEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns empty -- Empty", actual)
}

func Test_Cov8_DynamicCollection_Add(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.Add(coredynamic.NewDynamic("a", true))
	dc.AddPtr(coredynamic.NewDynamicPtr("b", true))
	dc.AddPtr(nil) // skip nil
	dc.AddAny("c", true)
	dc.AddAnyNonNull("d", true)
	dc.AddAnyNonNull(nil, true) // skip nil
	actual := args.Map{
		"len":    dc.Length(),
		"hasIdx": dc.HasIndex(3),
		"noIdx":  dc.HasIndex(10),
	}
	expected := args.Map{"len": 4, "hasIdx": true, "noIdx": false}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Add", actual)
}

func Test_Cov8_DynamicCollection_AddAnyMany(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnyMany("a", "b", "c")
	dc.AddAnyMany() // nil — no-op
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AddAnyMany", actual)
}

func Test_Cov8_DynamicCollection_AddManyPtr(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddManyPtr(
		coredynamic.NewDynamicPtr("a", true),
		nil,
		coredynamic.NewDynamicPtr("b", true),
	)
	dc.AddManyPtr() // nil — no-op
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AddManyPtr", actual)
}

func Test_Cov8_DynamicCollection_FirstLastOrDefault(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAny("first", true)
	dc.AddAny("last", true)
	emptyDC := coredynamic.EmptyDynamicCollection()
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
	expected := args.Map{
		"first": "first", "last": "last",
		"firstDyn": true, "lastDyn": true,
		"firstOrDef": true, "lastOrDef": true,
		"firstOrDefDyn": true, "lastOrDefDyn": true,
		"emptyFirstDef": true, "emptyLastDef": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- FirstLast", actual)
}

func Test_Cov8_DynamicCollection_At(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAny("hello", true)
	d := dc.At(0)
	actual := args.Map{"val": d.Value()}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- At", actual)
}

func Test_Cov8_DynamicCollection_Items(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAny("a", true)
	var nilDC *coredynamic.DynamicCollection
	actual := args.Map{
		"itemsLen": len(dc.Items()),
		"nilItems": len(nilDC.Items()),
	}
	expected := args.Map{"itemsLen": 1, "nilItems": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Items", actual)
}

func Test_Cov8_DynamicCollection_SkipTakeLimitSlice(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(10)
	dc.AddAnyMany("a", "b", "c", "d", "e")
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
	expected := args.Map{
		"skipLen": 3, "skipDynLen": true, "skipCol": 3,
		"takeLen": 2, "takeDynLen": true, "takeCol": 2,
		"limitLen": 3, "limitDynLen": true, "limitCol": 3,
		"safeLimitCol": 5,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Skip/Take/Limit", actual)
}

func Test_Cov8_DynamicCollection_RemoveAt(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnyMany("a", "b", "c")
	ok := dc.RemoveAt(1)
	fail := dc.RemoveAt(100)
	actual := args.Map{
		"ok":     ok,
		"fail":   fail,
		"newLen": dc.Length(),
	}
	expected := args.Map{"ok": true, "fail": false, "newLen": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- RemoveAt", actual)
}

func Test_Cov8_DynamicCollection_AnyItems(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnyMany("a", "b")
	items := dc.AnyItems()
	emptyDC := coredynamic.EmptyDynamicCollection()
	emptyItems := emptyDC.AnyItems()
	actual := args.Map{
		"itemsLen": len(items),
		"emptyLen": len(emptyItems),
	}
	expected := args.Map{"itemsLen": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AnyItems", actual)
}

func Test_Cov8_DynamicCollection_AnyItemsCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnyMany("a", "b")
	ac := dc.AnyItemsCollection()
	emptyDC := coredynamic.EmptyDynamicCollection()
	emptyAC := emptyDC.AnyItemsCollection()
	actual := args.Map{
		"acLen":    ac.Length(),
		"emptyLen": emptyAC.Length(),
	}
	expected := args.Map{"acLen": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AnyItemsCollection", actual)
}

func Test_Cov8_DynamicCollection_ListStrings(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnyMany("hello", "world")
	strs := dc.ListStrings()
	strsPtr := dc.ListStringsPtr()
	actual := args.Map{
		"strsLen":    len(strs),
		"strsPtrLen": len(strsPtr),
	}
	expected := args.Map{"strsLen": 2, "strsPtrLen": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- ListStrings", actual)
}

func Test_Cov8_DynamicCollection_Strings(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnyMany("a", "b")
	strs := dc.Strings()
	str := dc.String()
	emptyDC := coredynamic.EmptyDynamicCollection()
	emptyStrs := emptyDC.Strings()
	actual := args.Map{
		"strsLen":   len(strs),
		"strNotEmpty": str != "",
		"emptyLen":  len(emptyStrs),
	}
	expected := args.Map{"strsLen": 2, "strNotEmpty": true, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Strings", actual)
}

func Test_Cov8_DynamicCollection_Loop(t *testing.T) {
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
	actual := args.Map{"count": count, "emptyCount": emptyCount}
	expected := args.Map{"count": 3, "emptyCount": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Loop", actual)
}

func Test_Cov8_DynamicCollection_Loop_Break(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnyMany("a", "b", "c")
	count := 0
	dc.Loop(func(index int, dynamicItem *coredynamic.Dynamic) (isBreak bool) {
		count++
		return true // break on first
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Loop break", actual)
}

func Test_Cov8_DynamicCollection_Json(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAny("hello", true)
	jsonResult := dc.Json()
	jsonPtr := dc.JsonPtr()
	model := dc.JsonModel()
	modelAny := dc.JsonModelAny()
	js, jsErr := dc.JsonString()
	jsMust := dc.JsonStringMust()
	actual := args.Map{
		"jsonOk":     jsonResult.HasError() == false,
		"ptrNotNil":  jsonPtr != nil,
		"modelItems": len(model.Items) > 0,
		"modelAnyNN": modelAny != nil,
		"jsNotEmpty": js != "",
		"jsErrNil":   jsErr == nil,
		"mustNE":     jsMust != "",
	}
	expected := args.Map{
		"jsonOk": true, "ptrNotNil": true,
		"modelItems": true, "modelAnyNN": true,
		"jsNotEmpty": true, "jsErrNil": true, "mustNE": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Json", actual)
}

func Test_Cov8_DynamicCollection_JsonResultsCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAny("hello", true)
	rc := dc.JsonResultsCollection()
	rpc := dc.JsonResultsPtrCollection()
	emptyDC := coredynamic.EmptyDynamicCollection()
	emptyRC := emptyDC.JsonResultsCollection()
	actual := args.Map{
		"rcNotNil":    rc != nil,
		"rpcNotNil":   rpc != nil,
		"emptyRCNotNil": emptyRC != nil,
	}
	expected := args.Map{"rcNotNil": true, "rpcNotNil": true, "emptyRCNotNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- JsonResultsCollection", actual)
}

func Test_Cov8_DynamicCollection_Paging(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(10)
	dc.AddAnyMany("a", "b", "c", "d", "e")
	pages := dc.GetPagesSize(2)
	paged := dc.GetPagedCollection(2)
	single := dc.GetSinglePageCollection(2, 1)
	actual := args.Map{
		"pages":     pages,
		"pagedLen":  len(paged),
		"singleLen": single.Length(),
	}
	expected := args.Map{"pages": 3, "pagedLen": 3, "singleLen": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Paging", actual)
}

func Test_Cov8_DynamicCollection_Paging_SmallSet(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAny("a", true)
	pages := dc.GetPagesSize(0)
	paged := dc.GetPagedCollection(10)
	single := dc.GetSinglePageCollection(10, 1)
	actual := args.Map{
		"zeroPage":   pages,
		"pagedSelf":  len(paged),
		"singleSelf": single.Length(),
	}
	expected := args.Map{"zeroPage": 0, "pagedSelf": 1, "singleSelf": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Paging small set", actual)
}

func Test_Cov8_DynamicCollection_ParseJson(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAny("hello", true)
	jsonResult := dc.Json()
	jsonPtr := &jsonResult

	target := coredynamic.EmptyDynamicCollection()
	parsed, err := target.ParseInjectUsingJson(jsonPtr)
	actual := args.Map{
		"parsedNotNil": parsed != nil,
		"errNil":       err == nil,
	}
	expected := args.Map{"parsedNotNil": true, "errNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- ParseInjectUsingJson", actual)
}

func Test_Cov8_DynamicCollection_JsonParseSelfInject(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAny("hello", true)
	jsonResult := dc.Json()
	jsonPtr := &jsonResult

	target := coredynamic.EmptyDynamicCollection()
	err := target.JsonParseSelfInject(jsonPtr)
	actual := args.Map{"errNil": err == nil}
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- JsonParseSelfInject", actual)
}

// ═══════════════════════════════════════════
// MapAnyItems — constructors & basic
// ═══════════════════════════════════════════

func Test_Cov8_MapAnyItems_Constructors(t *testing.T) {
	empty := coredynamic.EmptyMapAnyItems()
	m := coredynamic.NewMapAnyItems(5)
	fromItems := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	fromItemsEmpty := coredynamic.NewMapAnyItemsUsingItems(nil)
	var nilM *coredynamic.MapAnyItems
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
	expected := args.Map{
		"emptyLen": 0, "mLen": 0, "fromItemsLen": 1,
		"fromEmptyLen": 0, "nilLen": 0, "isEmpty": true,
		"hasAny": true, "nilHasKey": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Constructors", actual)
}

func Test_Cov8_MapAnyItems_AddSet(t *testing.T) {
	m := coredynamic.NewMapAnyItems(5)
	isNew1 := m.Add("a", 1)
	isNew2 := m.Add("a", 2) // overwrite
	isNew3 := m.Set("b", 3)
	actual := args.Map{
		"isNew1": isNew1, "isNew2": isNew2, "isNew3": isNew3,
		"len": m.Length(), "hasA": m.HasKey("a"), "hasB": m.HasKey("b"),
	}
	expected := args.Map{
		"isNew1": true, "isNew2": false, "isNew3": true,
		"len": 2, "hasA": true, "hasB": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Add/Set", actual)
}

func Test_Cov8_MapAnyItems_Get(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 42})
	val, has := m.Get("a")
	_, hasMissing := m.Get("z")
	getVal := m.GetValue("a")
	getMissing := m.GetValue("z")
	actual := args.Map{
		"val": val, "has": has, "hasMissing": hasMissing,
		"getVal": getVal, "getMissing": getMissing == nil,
	}
	expected := args.Map{
		"val": 42, "has": true, "hasMissing": false,
		"getVal": 42, "getMissing": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Get", actual)
}

func Test_Cov8_MapAnyItems_AllKeys(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"b": 2, "a": 1})
	keys := m.AllKeys()
	sorted := m.AllKeysSorted()
	vals := m.AllValues()
	emptyM := coredynamic.EmptyMapAnyItems()
	actual := args.Map{
		"keysLen":    len(keys),
		"sortedFirst": sorted[0],
		"valsLen":    len(vals),
		"emptyKeys":  len(emptyM.AllKeys()),
		"emptyVals":  len(emptyM.AllValues()),
	}
	expected := args.Map{
		"keysLen": 2, "sortedFirst": "a", "valsLen": 2,
		"emptyKeys": 0, "emptyVals": 0,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- AllKeys", actual)
}

func Test_Cov8_MapAnyItems_AddMapResult(t *testing.T) {
	m := coredynamic.NewMapAnyItems(5)
	m.Add("a", 1)
	m.AddMapResult(map[string]any{"b": 2, "c": 3})
	m.AddMapResult(nil) // no-op
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- AddMapResult", actual)
}

func Test_Cov8_MapAnyItems_AddManyMapResults(t *testing.T) {
	m := coredynamic.NewMapAnyItems(5)
	m.AddManyMapResultsUsingOption(true, map[string]any{"a": 1}, map[string]any{"b": 2})
	m.AddManyMapResultsUsingOption(true) // empty — no-op
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- AddManyMapResults", actual)
}

func Test_Cov8_MapAnyItems_GetNewMapUsingKeys(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2, "c": 3})
	sub := m.GetNewMapUsingKeys(false, "a", "c")
	emptyKeys := m.GetNewMapUsingKeys(false)
	actual := args.Map{
		"subLen":   sub.Length(),
		"emptyLen": emptyKeys.Length(),
	}
	expected := args.Map{"subLen": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetNewMapUsingKeys", actual)
}

func Test_Cov8_MapAnyItems_Json(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	jsonResult := m.Json()
	jsonPtr := m.JsonPtr()
	model := m.JsonModel()
	modelAny := m.JsonModelAny()
	js, jsErr := m.JsonString()
	jsMust := m.JsonStringMust()
	actual := args.Map{
		"jsonOk":     jsonResult.JsonString() != "",
		"ptrNotNil":  jsonPtr != nil,
		"modelNN":    model != nil,
		"modelAnyNN": modelAny != nil,
		"jsNotEmpty": js != "",
		"jsErrNil":   jsErr == nil,
		"mustNE":     jsMust != "",
	}
	expected := args.Map{
		"jsonOk": true, "ptrNotNil": true,
		"modelNN": true, "modelAnyNN": true,
		"jsNotEmpty": true, "jsErrNil": true, "mustNE": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Json", actual)
}

func Test_Cov8_MapAnyItems_Strings(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	strs := m.Strings()
	str := m.String()
	actual := args.Map{
		"strsLen":    len(strs),
		"strNotEmpty": str != "",
	}
	expected := args.Map{"strsLen": 1, "strNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Strings", actual)
}

func Test_Cov8_MapAnyItems_ClearDisposeDeepClear(t *testing.T) {
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

	actual := args.Map{
		"cleared":     cleared,
		"deepCleared": deepCleared,
		"disposed":    disposed,
	}
	expected := args.Map{"cleared": 0, "deepCleared": 0, "disposed": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Clear/DeepClear/Dispose", actual)
}

func Test_Cov8_MapAnyItems_IsEqual(t *testing.T) {
	m1 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	m2 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	m3 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 3})
	m4 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	var nilM *coredynamic.MapAnyItems
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
	expected := args.Map{
		"equal": true, "notEqual": false, "diffLen": false,
		"nilBothEq": true, "nilOneNotEq": false,
		"rawEqual": true, "rawNotEqual": false, "rawMissingKey": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- IsEqual", actual)
}
func Test_Cov8_MapAnyItems_MapAnyItemsSelf(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	self := m.MapAnyItems()
	actual := args.Map{"same": self == m}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- self-reference", actual)
}

func Test_Cov8_MapAnyItems_Paging(t *testing.T) {
	m := coredynamic.NewMapAnyItems(10)
	for i := 0; i < 5; i++ {
		m.Add("k"+string(rune('a'+i)), i)
	}
	pages := m.GetPagesSize(2)
	paged := m.GetPagedCollection(2)
	actual := args.Map{
		"pages":    pages,
		"pagedLen": len(paged),
	}
	expected := args.Map{"pages": 3, "pagedLen": 3}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Paging", actual)
}

func Test_Cov8_MapAnyItems_Paging_SmallSet(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	pages := m.GetPagesSize(0)
	paged := m.GetPagedCollection(10)
	actual := args.Map{
		"zeroPage":  pages,
		"pagedSelf": len(paged),
	}
	expected := args.Map{"zeroPage": 0, "pagedSelf": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Paging small set", actual)
}

func Test_Cov8_MapAnyItems_ParseJson(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	jsonResult := m.Json()
	jsonPtr := &jsonResult

	target := coredynamic.EmptyMapAnyItems()
	parsed, err := target.ParseInjectUsingJson(jsonPtr)
	actual := args.Map{
		"parsedNotNil": parsed != nil,
		"errNil":       err == nil,
	}
	expected := args.Map{"parsedNotNil": true, "errNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- ParseInjectUsingJson", actual)
}

func Test_Cov8_MapAnyItems_JsonParseSelfInject(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	jsonResult := m.Json()
	jsonPtr := &jsonResult

	target := coredynamic.EmptyMapAnyItems()
	err := target.JsonParseSelfInject(jsonPtr)
	actual := args.Map{"errNil": err == nil}
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonParseSelfInject", actual)
}

func Test_Cov8_MapAnyItems_JsonResultsCollections(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "hello"})
	rc := m.JsonResultsCollection()
	rpc := m.JsonResultsPtrCollection()
	mr, mrErr := m.JsonMapResults()
	emptyM := coredynamic.EmptyMapAnyItems()
	emptyRC := emptyM.JsonResultsCollection()
	emptyMR, emptyMRErr := emptyM.JsonMapResults()
	actual := args.Map{
		"rcNotNil":      rc != nil,
		"rpcNotNil":     rpc != nil,
		"mrNotNil":      mr != nil,
		"mrErrNil":      mrErr == nil,
		"emptyRCNotNil": emptyRC != nil,
		"emptyMRNotNil": emptyMR != nil,
		"emptyMRErrNil": emptyMRErr == nil,
	}
	expected := args.Map{
		"rcNotNil": true, "rpcNotNil": true,
		"mrNotNil": true, "mrErrNil": true,
		"emptyRCNotNil": true, "emptyMRNotNil": true, "emptyMRErrNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonResultsCollections", actual)
}

func Test_Cov8_MapAnyItems_JsonResultOfKey(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "hello"})
	found := m.JsonResultOfKey("a")
	missing := m.JsonResultOfKey("z")
	actual := args.Map{
		"foundHasBytes":  len(found.Bytes) > 0,
		"missingHasErr":  missing.HasError(),
	}
	expected := args.Map{"foundHasBytes": true, "missingHasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonResultOfKey", actual)
}

func Test_Cov8_MapAnyItems_JsonResultOfKeys(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "hello", "b": 42})
	results := m.JsonResultOfKeys("a", "b")
	emptyResults := m.JsonResultOfKeys()
	actual := args.Map{
		"resultsNN": results != nil,
		"emptyNN":   emptyResults != nil,
	}
	expected := args.Map{"resultsNN": true, "emptyNN": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonResultOfKeys", actual)
}

func Test_Cov8_MapAnyItems_AddJsonResultPtr(t *testing.T) {
	m := coredynamic.NewMapAnyItems(5)
	m.AddJsonResultPtr("a", nil) // nil — should skip
	jr := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"x": 1}).JsonPtr()
	m.AddJsonResultPtr("b", jr) // non-nil
	actual := args.Map{
		"noA": m.HasKey("a"),
		"hasB": m.HasKey("b"),
	}
	expected := args.Map{"noA": false, "hasB": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- AddJsonResultPtr", actual)
}

func Test_Cov8_MapAnyItems_RawMapStringAnyDiff(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	diff := m.RawMapStringAnyDiff()
	var nilM *coredynamic.MapAnyItems
	nilDiff := nilM.RawMapStringAnyDiff()
	actual := args.Map{
		"diffLen":    len(diff),
		"nilDiffLen": len(nilDiff),
	}
	expected := args.Map{"diffLen": 1, "nilDiffLen": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- RawMapStringAnyDiff", actual)
}

func Test_Cov8_MapAnyItems_Deserialize(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"name": "test"})
	var result string
	err := m.Deserialize("name", &result)
	errMissing := m.Deserialize("missing", &result)
	actual := args.Map{
		"result":     result,
		"errNil":     err == nil,
		"missingErr": errMissing != nil,
	}
	expected := args.Map{"result": "test", "errNil": true, "missingErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Deserialize", actual)
}

func Test_Cov8_MapAnyItems_GetFieldsMap(t *testing.T) {
	inner := map[string]any{"x": 1}
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"data": inner})
	_, _, found := m.GetFieldsMap("data")
	_, notFound := m.GetSafeFieldsMap("missing")
	actual := args.Map{
		"found":    found,
		"notFound": notFound,
	}
	expected := args.Map{"found": true, "notFound": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetFieldsMap", actual)
}

func Test_Cov8_MapAnyItems_GetManyItemsRefs_Empty(t *testing.T) {
	m := coredynamic.NewMapAnyItems(5)
	err := m.GetManyItemsRefs() // empty — should return nil
	actual := args.Map{"errNil": err == nil}
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- GetManyItemsRefs empty", actual)
}

func Test_Cov8_MapAnyItems_HasAnyChanges(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	changed := m.HasAnyChanges(true, map[string]any{"a": 2})
	notChanged := m.HasAnyChanges(true, map[string]any{"a": 1})
	actual := args.Map{
		"changed":    changed,
		"notChanged": notChanged,
	}
	expected := args.Map{"changed": true, "notChanged": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- HasAnyChanges", actual)
}
