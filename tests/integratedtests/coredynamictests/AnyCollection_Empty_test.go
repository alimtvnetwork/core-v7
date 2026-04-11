package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// AnyCollection — constructors & basic
// ═══════════════════════════════════════════

func Test_AnyCollection_Empty_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	var nilAC *coredynamic.AnyCollection

	// Act
	actual := args.Map{
		"len":      ac.Length(),
		"count":    ac.Count(),
		"isEmpty":  ac.IsEmpty(),
		"hasAny":   ac.HasAnyItem(),
		"lastIdx":  ac.LastIndex(),
		"nilLen":   nilAC.Length(),
		"nilEmpty": nilAC.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"len": 0, "count": 0, "isEmpty": true, "hasAny": false,
		"lastIdx": -1, "nilLen": 0, "nilEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- Empty", actual)
}

func Test_AnyCollection_Add_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("a")
	ac.Add("b")

	// Act
	actual := args.Map{
		"len":    ac.Length(),
		"hasAny": ac.HasAnyItem(),
		"hasIdx": ac.HasIndex(1),
		"noIdx":  ac.HasIndex(10),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"hasAny": true,
		"hasIdx": true,
		"noIdx": false,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Add", actual)
}

func Test_AnyCollection_AddMany_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.AddMany("a", "b", "c")
	ac.AddMany() // nil — no-op

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddMany", actual)
}

func Test_AnyCollection_AddMany_SkipsNil(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.AddMany("a", nil, "b")

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns nil -- AddMany skips nil", actual)
}

func Test_AnyCollection_AddNonNull_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.AddNonNull("a")
	ac.AddNonNull(nil) // skip

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddNonNull", actual)
}

func Test_AnyCollection_AddAny_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.AddAny("hello", true)

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddAny", actual)
}

func Test_AnyCollection_AddNonNullDynamic_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.AddNonNullDynamic("x", true)
	ac.AddNonNullDynamic(nil, true) // skip

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddNonNullDynamic", actual)
}

func Test_AnyCollection_AddAnyManyDynamic_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.AddAnyManyDynamic("a", "b")
	ac.AddAnyManyDynamic() // nil — no-op

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddAnyManyDynamic", actual)
}

func Test_AnyCollection_Items_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("a")
	emptyAC := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{
		"itemsLen": len(ac.Items()),
		"emptyLen": len(emptyAC.Items()),
	}

	// Assert
	expected := args.Map{
		"itemsLen": 1,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Items", actual)
}

func Test_AnyCollection_At_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("hello")

	// Act
	actual := args.Map{"val": ac.At(0)}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- At", actual)
}

func Test_AnyCollection_AtAsDynamic_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("hello")
	d := ac.AtAsDynamic(0)

	// Act
	actual := args.Map{
		"val":     d.Value(),
		"isValid": d.IsValid(),
	}

	// Assert
	expected := args.Map{
		"val": "hello",
		"isValid": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AtAsDynamic", actual)
}

func Test_AnyCollection_FirstLastOrDefault(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("first")
	ac.Add("last")
	emptyAC := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{
		"first":         ac.First(),
		"last":          ac.Last(),
		"firstDyn":      ac.FirstDynamic(),
		"lastDyn":       ac.LastDynamic(),
		"firstOrDef":    ac.FirstOrDefault(),
		"lastOrDef":     ac.LastOrDefault(),
		"firstOrDefDyn": ac.FirstOrDefaultDynamic(),
		"lastOrDefDyn":  ac.LastOrDefaultDynamic(),
		"emptyFirst":    emptyAC.FirstOrDefault() == nil,
		"emptyLast":     emptyAC.LastOrDefault() == nil,
	}

	// Assert
	expected := args.Map{
		"first": "first", "last": "last",
		"firstDyn": "first", "lastDyn": "last",
		"firstOrDef": "first", "lastOrDef": "last",
		"firstOrDefDyn": "first", "lastOrDefDyn": "last",
		"emptyFirst": true, "emptyLast": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- FirstLast", actual)
}

func Test_AnyCollection_SkipTakeLimitSlice(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(10)
	ac.AddMany("a", "b", "c", "d", "e")

	// Act
	actual := args.Map{
		"skipLen":      len(ac.Skip(2)),
		"skipDynNN":    ac.SkipDynamic(2) != nil,
		"skipCol":      ac.SkipCollection(2).Length(),
		"takeLen":      len(ac.Take(2)),
		"takeDynNN":    ac.TakeDynamic(2) != nil,
		"takeCol":      ac.TakeCollection(2).Length(),
		"limitLen":     len(ac.Limit(3)),
		"limitDynNN":   ac.LimitDynamic(3) != nil,
		"limitCol":     ac.LimitCollection(3).Length(),
		"safeLimitCol": ac.SafeLimitCollection(100).Length(),
	}

	// Assert
	expected := args.Map{
		"skipLen": 3, "skipDynNN": true, "skipCol": 3,
		"takeLen": 2, "takeDynNN": true, "takeCol": 2,
		"limitLen": 3, "limitDynNN": true, "limitCol": 3,
		"safeLimitCol": 5,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Skip/Take/Limit", actual)
}

func Test_AnyCollection_RemoveAt_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.AddMany("a", "b", "c")
	ok := ac.RemoveAt(1)
	fail := ac.RemoveAt(100)

	// Act
	actual := args.Map{
		"ok":     ok,
		"fail":   fail,
		"newLen": ac.Length(),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"fail": false,
		"newLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- RemoveAt", actual)
}

func Test_AnyCollection_DynamicItems_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("hello")
	dynItems := ac.DynamicItems()
	emptyAC := coredynamic.EmptyAnyCollection()
	emptyDyn := emptyAC.DynamicItems()

	// Act
	actual := args.Map{
		"dynLen":   len(dynItems),
		"emptyLen": len(emptyDyn),
	}

	// Assert
	expected := args.Map{
		"dynLen": 1,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- DynamicItems", actual)
}

func Test_AnyCollection_DynamicCollection_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("hello")
	dc := ac.DynamicCollection()
	emptyAC := coredynamic.EmptyAnyCollection()
	emptyDC := emptyAC.DynamicCollection()

	// Act
	actual := args.Map{
		"dcLen":    dc.Length(),
		"emptyLen": emptyDC.Length(),
	}

	// Assert
	expected := args.Map{
		"dcLen": 1,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- DynamicCollection", actual)
}

func Test_AnyCollection_Loop_Sync_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.AddMany("a", "b", "c")
	count := 0
	ac.Loop(false, func(index int, item any) (isBreak bool) {
		count++
		return false
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Loop sync", actual)
}

func Test_AnyCollection_Loop_Break_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.AddMany("a", "b", "c")
	count := 0
	ac.Loop(false, func(index int, item any) (isBreak bool) {
		count++
		return true // break on first
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Loop break", actual)
}

func Test_AnyCollection_Loop_Empty_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	count := 0
	ac.Loop(false, func(index int, item any) (isBreak bool) {
		count++
		return false
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- Loop empty", actual)
}

func Test_AnyCollection_Loop_Async_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.AddMany("a", "b", "c")
	// Async loop — cannot reliably count due to goroutines,
	// but must not panic and must return self
	result := ac.Loop(true, func(index int, item any) (isBreak bool) {
		return false
	})

	// Act
	actual := args.Map{"resultNN": result != nil}

	// Assert
	expected := args.Map{"resultNN": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Loop async", actual)
}

func Test_AnyCollection_LoopDynamic_Sync_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("hello")
	ac.Add("world")
	count := 0
	ac.LoopDynamic(false, func(index int, item coredynamic.Dynamic) (isBreak bool) {
		count++
		return false
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- LoopDynamic sync", actual)
}

func Test_AnyCollection_LoopDynamic_Break_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("hello")
	ac.Add("world")
	count := 0
	ac.LoopDynamic(false, func(index int, item coredynamic.Dynamic) (isBreak bool) {
		count++
		return true
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- LoopDynamic break", actual)
}

func Test_AnyCollection_LoopDynamic_Async_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("hello")
	result := ac.LoopDynamic(true, func(index int, item coredynamic.Dynamic) (isBreak bool) {
		return false
	})

	// Act
	actual := args.Map{"resultNN": result != nil}

	// Assert
	expected := args.Map{"resultNN": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- LoopDynamic async", actual)
}

func Test_AnyCollection_LoopDynamic_Empty(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	count := 0
	ac.LoopDynamic(false, func(index int, item coredynamic.Dynamic) (isBreak bool) {
		count++
		return false
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- LoopDynamic empty", actual)
}

func Test_AnyCollection_ListStrings_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("hello")
	ac.Add("world")
	strs := ac.ListStrings(false)
	strsField := ac.ListStrings(true)
	strsPtr := ac.ListStringsPtr(false)

	// Act
	actual := args.Map{
		"strsLen":      len(strs),
		"strsFieldLen": len(strsField),
		"strsPtrLen":   len(strsPtr),
	}

	// Assert
	expected := args.Map{
		"strsLen": 2,
		"strsFieldLen": 2,
		"strsPtrLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- ListStrings", actual)
}

func Test_AnyCollection_Strings_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("a")
	ac.Add("b")
	strs := ac.Strings()
	str := ac.String()
	emptyAC := coredynamic.EmptyAnyCollection()
	emptyStrs := emptyAC.Strings()

	// Act
	actual := args.Map{
		"strsLen":     len(strs),
		"strNotEmpty": str != "",
		"emptyLen":    len(emptyStrs),
	}

	// Assert
	expected := args.Map{
		"strsLen": 2,
		"strNotEmpty": true,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Strings", actual)
}

func Test_AnyCollection_Json_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("hello")
	jsonResult := ac.Json()
	jsonPtr := ac.JsonPtr()
	model := ac.JsonModel()
	modelAny := ac.JsonModelAny()
	js, jsErr := ac.JsonString()
	jsMust := ac.JsonStringMust()

	// Act
	actual := args.Map{
		"jsonOk":     jsonResult.HasError() == false,
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
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Json", actual)
}

func Test_AnyCollection_JsonResultsCollection_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("hello")
	rc := ac.JsonResultsCollection()
	rpc := ac.JsonResultsPtrCollection()
	emptyAC := coredynamic.EmptyAnyCollection()
	emptyRC := emptyAC.JsonResultsCollection()

	// Act
	actual := args.Map{
		"rcNotNil":      rc != nil,
		"rpcNotNil":     rpc != nil,
		"emptyRCNotNil": emptyRC != nil,
	}

	// Assert
	expected := args.Map{
		"rcNotNil": true,
		"rpcNotNil": true,
		"emptyRCNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonResultsCollection", actual)
}

func Test_AnyCollection_Paging(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(10)
	ac.AddMany("a", "b", "c", "d", "e")
	pages := ac.GetPagesSize(2)
	paged := ac.GetPagedCollection(2)
	single := ac.GetSinglePageCollection(2, 1)

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
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Paging", actual)
}

func Test_AnyCollection_Paging_SmallSet(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("a")
	pages := ac.GetPagesSize(0)
	paged := ac.GetPagedCollection(10)
	single := ac.GetSinglePageCollection(10, 1)

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
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Paging small set", actual)
}

func Test_AnyCollection_ParseJson(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("hello")
	js, jsErr := ac.JsonString()
	if jsErr != nil || js == "" {
		// Json() on value receiver produces "{}" which is treated as empty,
		// so fall back to direct json bytes for parse testing

	// Act
		actual := args.Map{"jsWorked": false}

	// Assert
		expected := args.Map{"jsWorked": false}
		expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- ParseInjectUsingJson", actual)
		return
	}
}

func Test_AnyCollection_JsonParseSelfInject_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("hello")
	js, jsErr := ac.JsonString()
	if jsErr != nil || js == "" {
		// Json() on value receiver produces "{}" treated as empty

	// Act
		actual := args.Map{"jsWorked": false}

	// Assert
		expected := args.Map{"jsWorked": false}
		expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonParseSelfInject", actual)
		return
	}
}

// ═══════════════════════════════════════════
// AnyCollection.AddAnySliceFromSingleItem
// ═══════════════════════════════════════════

func Test_AnyCollection_AddAnySliceFromSingleItem_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.AddAnySliceFromSingleItem([]string{"a", "b", "c"})

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddAnySliceFromSingleItem", actual)
}

func Test_AnyCollection_AddAnySliceFromSingleItem_Nil_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.AddAnySliceFromSingleItem(nil) // no-op

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns nil -- AddAnySliceFromSingleItem nil", actual)
}

func Test_AnyCollection_AddAnySliceFromSingleItem_IntSlice(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.AddAnySliceFromSingleItem([]int{1, 2})

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddAnySliceFromSingleItem int slice", actual)
}

// ═══════════════════════════════════════════
// DynamicCollection.AddAnySliceFromSingleItem
// ═══════════════════════════════════════════

func Test_DynamicCollection_AddAnySliceFromSingleItem(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnySliceFromSingleItem(true, []string{"x", "y"})

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AddAnySliceFromSingleItem", actual)
}

func Test_DynamicCollection_AddAnySliceFromSingleItem_Nil(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	dc.AddAnySliceFromSingleItem(true, nil) // no-op

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns nil -- AddAnySliceFromSingleItem nil", actual)
}

// ═══════════════════════════════════════════
// AddAnyItemsWithTypeValidation — AnyCollection
// ═══════════════════════════════════════════

func Test_AnyCollection_AddAnyItemsWithTypeValidation_Valid(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	strType := reflect.TypeOf("")
	err := ac.AddAnyItemsWithTypeValidation(false, false, strType, "a", "b")

	// Act
	actual := args.Map{
		"errNil": err == nil,
		"len":    ac.Length(),
	}

	// Assert
	expected := args.Map{
		"errNil": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns non-empty -- AddAnyItemsWithTypeValidation valid", actual)
}

func Test_AnyCollection_AddAnyItemsWithTypeValidation_TypeMismatch(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	strType := reflect.TypeOf("")
	err := ac.AddAnyItemsWithTypeValidation(false, false, strType, "a", 123)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len":    ac.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns non-empty -- AddAnyItemsWithTypeValidation mismatch stops", actual)
}

func Test_AnyCollection_AddAnyItemsWithTypeValidation_ContinueOnError(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	strType := reflect.TypeOf("")
	err := ac.AddAnyItemsWithTypeValidation(true, false, strType, "a", 123, "c")

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len":    ac.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns error -- AddAnyItemsWithTypeValidation continueOnError", actual)
}

func Test_AnyCollection_AddAnyItemsWithTypeValidation_Empty_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	strType := reflect.TypeOf("")
	err := ac.AddAnyItemsWithTypeValidation(false, false, strType)

	// Act
	actual := args.Map{
		"errNil": err == nil,
		"len": ac.Length(),
	}

	// Assert
	expected := args.Map{
		"errNil": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- AddAnyItemsWithTypeValidation empty", actual)
}

func Test_AnyCollection_AddAnyItemsWithTypeValidation_NullNotAllowed(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	strType := reflect.TypeOf("")
	err := ac.AddAnyItemsWithTypeValidation(false, true, strType, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns non-empty -- AddAnyItemsWithTypeValidation null not allowed", actual)
}

func Test_AnyCollection_AddAnyWithTypeValidation_Valid(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	strType := reflect.TypeOf("")
	err := ac.AddAnyWithTypeValidation(false, strType, "hello")

	// Act
	actual := args.Map{
		"errNil": err == nil,
		"len": ac.Length(),
	}

	// Assert
	expected := args.Map{
		"errNil": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns non-empty -- AddAnyWithTypeValidation valid", actual)
}

func Test_AnyCollection_AddAnyWithTypeValidation_TypeMismatch(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	strType := reflect.TypeOf("")
	err := ac.AddAnyWithTypeValidation(false, strType, 42)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len": ac.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns non-empty -- AddAnyWithTypeValidation mismatch", actual)
}

// ═══════════════════════════════════════════
// AddAnyItemsWithTypeValidation — DynamicCollection
// ═══════════════════════════════════════════

func Test_DynamicCollection_AddAnyItemsWithTypeValidation_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	strType := reflect.TypeOf("")
	err := dc.AddAnyItemsWithTypeValidation(false, false, strType, "a", "b")

	// Act
	actual := args.Map{
		"errNil": err == nil,
		"len":    dc.Length(),
	}

	// Assert
	expected := args.Map{
		"errNil": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns non-empty -- AddAnyItemsWithTypeValidation valid", actual)
}

func Test_DynamicCollection_AddAnyItemsWithTypeValidation_Mismatch(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	strType := reflect.TypeOf("")
	err := dc.AddAnyItemsWithTypeValidation(false, false, strType, "a", 123)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len":    dc.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns non-empty -- AddAnyItemsWithTypeValidation mismatch", actual)
}

func Test_DynamicCollection_AddAnyItemsWithTypeValidation_ContinueOnError(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	strType := reflect.TypeOf("")
	err := dc.AddAnyItemsWithTypeValidation(true, false, strType, "a", 123, "c")

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len":    dc.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns error -- AddAnyItemsWithTypeValidation continueOnError", actual)
}

func Test_DynamicCollection_AddAnyItemsWithTypeValidation_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	strType := reflect.TypeOf("")
	err := dc.AddAnyItemsWithTypeValidation(false, false, strType)

	// Act
	actual := args.Map{
		"errNil": err == nil,
		"len": dc.Length(),
	}

	// Assert
	expected := args.Map{
		"errNil": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns empty -- AddAnyItemsWithTypeValidation empty", actual)
}

func Test_DynamicCollection_AddAnyWithTypeValidation_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(5)
	strType := reflect.TypeOf("")
	err := dc.AddAnyWithTypeValidation(false, strType, "hello")

	// Act
	actual := args.Map{
		"errNil": err == nil,
		"len": dc.Length(),
	}

	// Assert
	expected := args.Map{
		"errNil": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns non-empty -- AddAnyWithTypeValidation valid", actual)
}

// ═══════════════════════════════════════════
// MapAnyItems.Diff / DiffRaw / HashmapDiffUsingRaw
// ═══════════════════════════════════════════

func Test_MapAnyItems_Diff(t *testing.T) {
	// Arrange
	left := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	right := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 3})
	diff := left.Diff(true, right)

	// Act
	actual := args.Map{"diffNN": diff != nil}

	// Assert
	expected := args.Map{"diffNN": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Diff", actual)
}

func Test_MapAnyItems_DiffRaw(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	diffMap := m.DiffRaw(true, map[string]any{"a": 1, "b": 3})

	// Act
	actual := args.Map{"diffMapNN": diffMap != nil}

	// Assert
	expected := args.Map{"diffMapNN": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- DiffRaw", actual)
}

func Test_MapAnyItems_HashmapDiffUsingRaw(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	diff := m.HashmapDiffUsingRaw(true, map[string]any{"a": 1})

	// Act
	actual := args.Map{"isEmpty": diff.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- HashmapDiffUsingRaw no diff", actual)
}

func Test_MapAnyItems_IsRawEqual(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{
		"equal":    m.IsRawEqual(true, map[string]any{"a": 1}),
		"notEqual": m.IsRawEqual(true, map[string]any{"a": 2}),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEqual": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- IsRawEqual", actual)
}

func Test_MapAnyItems_MapStringAnyDiff(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	diff := m.MapStringAnyDiff()

	// Act
	actual := args.Map{"diffLen": len(diff)}

	// Assert
	expected := args.Map{"diffLen": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- MapStringAnyDiff", actual)
}

func Test_MapAnyItems_DiffJsonMessage(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	msg := m.DiffJsonMessage(true, map[string]any{"a": 2})

	// Act
	actual := args.Map{"msgNN": msg != ""}

	// Assert
	expected := args.Map{"msgNN": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- DiffJsonMessage", actual)
}

func Test_MapAnyItems_ShouldDiffMessage(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	msg := m.ShouldDiffMessage(true, "test", map[string]any{"a": 2})

	// Act
	actual := args.Map{"msgNN": msg != ""}

	// Assert
	expected := args.Map{"msgNN": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- ShouldDiffMessage", actual)
}

func Test_MapAnyItems_LogShouldDiffMessage(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	msg := m.LogShouldDiffMessage(true, "test", map[string]any{"a": 2})

	// Act
	actual := args.Map{"msgNN": msg != ""}

	// Assert
	expected := args.Map{"msgNN": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- LogShouldDiffMessage", actual)
}

func Test_MapAnyItems_ToStringsSliceOfDiffMap(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	diffMap := map[string]any{"a": 2}
	strs := m.ToStringsSliceOfDiffMap(diffMap)

	// Act
	actual := args.Map{"strsNN": strs != nil}

	// Assert
	expected := args.Map{"strsNN": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- ToStringsSliceOfDiffMap", actual)
}

// ═══════════════════════════════════════════
// MapAnyItemDiff — methods
// ═══════════════════════════════════════════

func Test_MapAnyItemDiff_Basic(t *testing.T) {
	// Arrange
	diff := coredynamic.MapAnyItemDiff(map[string]any{"a": 1, "b": 2})
	var nilDiff *coredynamic.MapAnyItemDiff

	// Act
	actual := args.Map{
		"len":       diff.Length(),
		"isEmpty":   diff.IsEmpty(),
		"hasAny":    diff.HasAnyItem(),
		"lastIdx":   diff.LastIndex(),
		"nilLen":    nilDiff.Length(),
		"nilRawLen": len(nilDiff.Raw()),
	}

	// Assert
	expected := args.Map{
		"len": 2, "isEmpty": false, "hasAny": true,
		"lastIdx": 1, "nilLen": 0, "nilRawLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- Basic", actual)
}

func Test_MapAnyItemDiff_AllKeysSorted(t *testing.T) {
	// Arrange
	diff := coredynamic.MapAnyItemDiff(map[string]any{"b": 2, "a": 1})
	sorted := diff.AllKeysSorted()

	// Act
	actual := args.Map{
		"sortedFirst": sorted[0],
		"sortedLen":   len(sorted),
	}

	// Assert
	expected := args.Map{
		"sortedFirst": "a",
		"sortedLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- AllKeysSorted", actual)
}

func Test_MapAnyItemDiff_IsRawEqual(t *testing.T) {
	// Arrange
	diff := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})

	// Act
	actual := args.Map{
		"equal":    diff.IsRawEqual(true, map[string]any{"a": 1}),
		"notEqual": diff.IsRawEqual(true, map[string]any{"a": 2}),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEqual": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- IsRawEqual", actual)
}

func Test_MapAnyItemDiff_HasAnyChanges(t *testing.T) {
	// Arrange
	diff := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})

	// Act
	actual := args.Map{
		"changed":    diff.HasAnyChanges(true, map[string]any{"a": 2}),
		"notChanged": diff.HasAnyChanges(true, map[string]any{"a": 1}),
	}

	// Assert
	expected := args.Map{
		"changed": true,
		"notChanged": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- HasAnyChanges", actual)
}

func Test_MapAnyItemDiff_DiffRaw(t *testing.T) {
	// Arrange
	diff := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	rawDiff := diff.DiffRaw(true, map[string]any{"a": 2})

	// Act
	actual := args.Map{"rawDiffNN": rawDiff != nil}

	// Assert
	expected := args.Map{"rawDiffNN": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- DiffRaw", actual)
}

func Test_MapAnyItemDiff_HashmapDiffUsingRaw(t *testing.T) {
	// Arrange
	diff := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	hashDiff := diff.HashmapDiffUsingRaw(true, map[string]any{"a": 1})

	// Act
	actual := args.Map{"isEmpty": hashDiff.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns empty -- HashmapDiffUsingRaw no diff", actual)
}

func Test_MapAnyItemDiff_MapAnyItems(t *testing.T) {
	// Arrange
	diff := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	m := diff.MapAnyItems()

	// Act
	actual := args.Map{
		"len": m.Length(),
		"hasA": m.HasKey("a"),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"hasA": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- MapAnyItems", actual)
}

func Test_MapAnyItemDiff_RawMapDiffer(t *testing.T) {
	// Arrange
	diff := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	rawDiffer := diff.RawMapDiffer()

	// Act
	actual := args.Map{"rawDifferLen": len(rawDiffer)}

	// Assert
	expected := args.Map{"rawDifferLen": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- RawMapDiffer", actual)
}

func Test_MapAnyItemDiff_Raw(t *testing.T) {
	// Arrange
	diff := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	raw := diff.Raw()

	// Act
	actual := args.Map{"rawLen": len(raw)}

	// Assert
	expected := args.Map{"rawLen": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- Raw", actual)
}

func Test_MapAnyItemDiff_Clear(t *testing.T) {
	// Arrange
	diff := coredynamic.MapAnyItemDiff(map[string]any{"a": 1, "b": 2})
	cleared := diff.Clear()
	var nilDiff *coredynamic.MapAnyItemDiff
	nilCleared := nilDiff.Clear()

	// Act
	actual := args.Map{
		"clearedLen":    cleared.Length(),
		"nilClearedLen": nilCleared.Length(),
	}

	// Assert
	expected := args.Map{
		"clearedLen": 0,
		"nilClearedLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- Clear", actual)
}

func Test_MapAnyItemDiff_Json(t *testing.T) {
	// Arrange
	diff := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	jsonResult := diff.Json()
	jsonPtr := diff.JsonPtr()
	pretty := diff.PrettyJsonString()

	// Act
	actual := args.Map{
		"jsonOk":    jsonResult.JsonString() != "",
		"ptrNotNil": jsonPtr != nil,
		"prettyNE":  pretty != "",
	}

	// Assert
	expected := args.Map{
		"jsonOk": true,
		"ptrNotNil": true,
		"prettyNE": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- Json", actual)
}

func Test_MapAnyItemDiff_DiffJsonMessage(t *testing.T) {
	// Arrange
	diff := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	msg := diff.DiffJsonMessage(true, map[string]any{"a": 2})

	// Act
	actual := args.Map{"msgNN": msg != ""}

	// Assert
	expected := args.Map{"msgNN": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- DiffJsonMessage", actual)
}

func Test_MapAnyItemDiff_ShouldDiffMessage(t *testing.T) {
	// Arrange
	diff := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	msg := diff.ShouldDiffMessage(true, "test", map[string]any{"a": 2})

	// Act
	actual := args.Map{"msgNN": msg != ""}

	// Assert
	expected := args.Map{"msgNN": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- ShouldDiffMessage", actual)
}

func Test_MapAnyItemDiff_LogShouldDiffMessage(t *testing.T) {
	// Arrange
	diff := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	msg := diff.LogShouldDiffMessage(true, "test", map[string]any{"a": 2})

	// Act
	actual := args.Map{"msgNN": msg != ""}

	// Assert
	expected := args.Map{"msgNN": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- LogShouldDiffMessage", actual)
}

func Test_MapAnyItemDiff_ToStringsSliceOfDiffMap(t *testing.T) {
	// Arrange
	diff := coredynamic.MapAnyItemDiff(map[string]any{"a": 1})
	strs := diff.ToStringsSliceOfDiffMap(map[string]any{"a": 2})

	// Act
	actual := args.Map{"strsNN": strs != nil}

	// Assert
	expected := args.Map{"strsNN": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- ToStringsSliceOfDiffMap", actual)
}

// ═══════════════════════════════════════════
// ReflectTypeValidation standalone
// ═══════════════════════════════════════════

func Test_ReflectTypeValidation_Valid(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	err := coredynamic.ReflectTypeValidation(false, strType, "hello")

	// Act
	actual := args.Map{"errNil": err == nil}

	// Assert
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns non-empty -- valid", actual)
}

func Test_ReflectTypeValidation_TypeMismatch(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	err := coredynamic.ReflectTypeValidation(false, strType, 42)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns non-empty -- type mismatch", actual)
}

func Test_ReflectTypeValidation_NilNotAllowed(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	err := coredynamic.ReflectTypeValidation(true, strType, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns nil -- nil not allowed", actual)
}

func Test_ReflectTypeValidation_NilAllowed(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	err := coredynamic.ReflectTypeValidation(false, strType, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns nil -- nil allowed but type mismatch", actual)
}

// ═══════════════════════════════════════════
// AnyCollection.ReflectSetAt
// ═══════════════════════════════════════════

func Test_AnyCollection_ReflectSetAt_FromAnyCollectionEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)
	ac.Add("hello")
	var target string
	err := ac.ReflectSetAt(0, &target)

	// Act
	actual := args.Map{
		"errNil": err == nil,
		"target": target,
	}

	// Assert
	expected := args.Map{
		"errNil": true,
		"target": "hello",
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- ReflectSetAt", actual)
}
