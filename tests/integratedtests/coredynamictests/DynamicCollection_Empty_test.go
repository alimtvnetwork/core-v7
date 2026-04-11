package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── DynamicCollection basic ops ──

func Test_DynamicCollection_Empty_FromDynamicCollectionEmp(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{
		"isEmpty": dc.IsEmpty(),
		"count":   dc.Count(),
		"hasAny":  dc.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
		"count":   0,
		"hasAny":  false,
	}
	expected.ShouldBeEqual(t, 0, "EmptyDynamicCollection returns empty -- new", actual)
}

func Test_DynamicCollection_AddAny_FromDynamicCollectionEmp(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello", true)
	dc.AddAny(42, true)

	// Act
	actual := args.Map{
		"count":     dc.Count(),
		"hasAny":    dc.HasAnyItem(),
		"isEmpty":   dc.IsEmpty(),
		"lastIndex": dc.LastIndex(),
		"hasIdx0":   dc.HasIndex(0),
		"hasIdx99":  dc.HasIndex(99),
	}

	// Assert
	expected := args.Map{
		"count":     2,
		"hasAny":    true,
		"isEmpty":   false,
		"lastIndex": 1,
		"hasIdx0":   true,
		"hasIdx99":  false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.AddAny works -- two items", actual)
}

func Test_DynamicCollection_FirstLast(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("first", true)
	dc.AddAny("last", true)

	// Act
	actual := args.Map{
		"first":          dc.First().Value(),
		"last":           dc.Last().Value(),
		"firstOrDefault": dc.FirstOrDefault().Value(),
		"lastOrDefault":  dc.LastOrDefault().Value(),
	}

	// Assert
	expected := args.Map{
		"first":          "first",
		"last":           "last",
		"firstOrDefault": "first",
		"lastOrDefault":  "last",
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection First/Last correct -- two items", actual)
}

func Test_DynamicCollection_SkipTakeLimit(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	dc.AddAny("b", true)
	dc.AddAny("c", true)

	// Act
	actual := args.Map{
		"skipLen":  len(dc.Skip(1)),
		"takeLen":  len(dc.Take(2)),
		"limitLen": len(dc.Limit(1)),
	}

	// Assert
	expected := args.Map{
		"skipLen":  2,
		"takeLen":  2,
		"limitLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Skip/Take/Limit correct -- 3 items", actual)
}

func Test_DynamicCollection_ListStrings_FromDynamicCollectionEmp(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello", true)
	dc.AddAny("world", true)
	strs := dc.ListStrings()

	// Act
	actual := args.Map{"len": len(strs)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.ListStrings returns 2 -- two items", actual)
}

func Test_DynamicCollection_RemoveAt_FromDynamicCollectionEmp(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	dc.AddAny("b", true)
	dc.AddAny("c", true)
	dc.RemoveAt(1)

	// Act
	actual := args.Map{"count": dc.Count()}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.RemoveAt removes middle -- 3 to 2", actual)
}

func Test_DynamicCollection_String_FromDynamicCollectionEmp(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello", true)
	s := dc.String()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.String returns non-empty -- single item", actual)
}

func Test_DynamicCollection_AddAnyNonNull_FromDynamicCollectionEmp(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyNonNull(nil, false)
	dc.AddAnyNonNull("hello", true)

	// Act
	actual := args.Map{"count": dc.Count()}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection.AddAnyNonNull skips nil -- 1 valid", actual)
}

// ── CollectionTypes (constructors take capacity int, not slices) ──

func Test_NewStringCollection(t *testing.T) {
	// Arrange
	sc := coredynamic.NewStringCollection(5)

	// Act
	actual := args.Map{"isEmpty": sc.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewStringCollection returns empty -- capacity 5", actual)
}

func Test_EmptyStringCollection(t *testing.T) {
	// Arrange
	sc := coredynamic.EmptyStringCollection()

	// Act
	actual := args.Map{"isEmpty": sc.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "EmptyStringCollection returns empty -- new", actual)
}

func Test_NewIntCollection(t *testing.T) {
	// Arrange
	ic := coredynamic.NewIntCollection(5)

	// Act
	actual := args.Map{"isEmpty": ic.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewIntCollection returns empty -- capacity 5", actual)
}

func Test_EmptyIntCollection(t *testing.T) {
	// Arrange
	ic := coredynamic.EmptyIntCollection()

	// Act
	actual := args.Map{"isEmpty": ic.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "EmptyIntCollection returns empty -- new", actual)
}

func Test_NewInt64Collection(t *testing.T) {
	// Arrange
	c := coredynamic.NewInt64Collection(5)

	// Act
	actual := args.Map{"isEmpty": c.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewInt64Collection returns empty -- capacity 5", actual)
}

func Test_NewByteCollection(t *testing.T) {
	// Arrange
	c := coredynamic.NewByteCollection(5)

	// Act
	actual := args.Map{"isEmpty": c.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewByteCollection returns empty -- capacity 5", actual)
}

func Test_NewBoolCollection(t *testing.T) {
	// Arrange
	c := coredynamic.NewBoolCollection(5)

	// Act
	actual := args.Map{"isEmpty": c.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewBoolCollection returns empty -- capacity 5", actual)
}

func Test_NewFloat64Collection(t *testing.T) {
	// Arrange
	c := coredynamic.NewFloat64Collection(5)

	// Act
	actual := args.Map{"isEmpty": c.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewFloat64Collection returns empty -- capacity 5", actual)
}

func Test_NewAnyMapCollection(t *testing.T) {
	// Arrange
	c := coredynamic.NewAnyMapCollection(5)

	// Act
	actual := args.Map{"isEmpty": c.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewAnyMapCollection returns empty -- capacity 5", actual)
}

func Test_NewStringMapCollection(t *testing.T) {
	// Arrange
	c := coredynamic.NewStringMapCollection(5)

	// Act
	actual := args.Map{"isEmpty": c.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewStringMapCollection returns empty -- capacity 5", actual)
}

// ── KeyVal ──

func Test_KeyVal_Basic(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "myKey", Value: "myVal"}

	// Act
	actual := args.Map{
		"key":       kv.Key,
		"val":       kv.Value,
		"isKeyNull": kv.IsKeyNull(),
	}

	// Assert
	expected := args.Map{
		"key":       "myKey",
		"val":       "myVal",
		"isKeyNull": false,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal basic getters -- key and val set", actual)
}

// ── KeyValCollection ──

func Test_KeyValCollection_Basic(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.KeyVal{Key: "k1", Value: "v1"})
	kvc.Add(coredynamic.KeyVal{Key: "k2", Value: "v2"})

	// Act
	actual := args.Map{
		"length":  kvc.Length(),
		"hasAny":  kvc.HasAnyItem(),
		"isEmpty": kvc.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"length":  2,
		"hasAny":  true,
		"isEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "KeyValCollection basic ops -- 2 items", actual)
}

func Test_KeyValCollection_AllKeys(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: "2"})
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: "1"})
	keys := kvc.AllKeys()

	// Act
	actual := args.Map{"len": len(keys)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyValCollection.AllKeys returns 2 -- two items", actual)
}

func Test_KeyValCollection_AllKeysSorted(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: "2"})
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: "1"})
	keys := kvc.AllKeysSorted()

	// Act
	actual := args.Map{"first": keys[0]}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "KeyValCollection.AllKeysSorted first is a -- sorted", actual)
}

func Test_KeyValCollection_AllValues(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	vals := kvc.AllValues()

	// Act
	actual := args.Map{"len": len(vals)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "KeyValCollection.AllValues returns 1 -- single item", actual)
}

func Test_KeyValCollection_String(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	s := kvc.String()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "KeyValCollection.String returns non-empty -- single item", actual)
}

// ── MapAnyItems ──

func Test_MapAnyItems_Basic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(5)
	m.Set("k1", "v1")
	m.Set("k2", 42)

	// Act
	actual := args.Map{
		"length":   m.Length(),
		"hasItems": m.HasAnyItem(),
		"isEmpty":  m.IsEmpty(),
		"hasK1":    m.HasKey("k1"),
		"hasX":     m.HasKey("x"),
	}

	// Assert
	expected := args.Map{
		"length":   2,
		"hasItems": true,
		"isEmpty":  false,
		"hasK1":    true,
		"hasX":     false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems basic ops -- 2 items", actual)
}

func Test_MapAnyItems_Get(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(5)
	m.Set("k1", "v1")
	val, has := m.Get("k1")

	// Act
	actual := args.Map{
		"val": val,
		"has": has,
	}

	// Assert
	expected := args.Map{
		"val": "v1",
		"has": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems.Get returns correct -- k1", actual)
}

func Test_MapAnyItems_AllKeysSorted(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(5)
	m.Set("c", 3)
	m.Set("a", 1)
	m.Set("b", 2)
	keys := m.AllKeysSorted()

	// Act
	actual := args.Map{
		"first": keys[0],
		"last": keys[2],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems.AllKeysSorted returns sorted -- 3 keys", actual)
}

func Test_MapAnyItems_String(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(5)
	m.Set("k", "v")
	s := m.String()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems.String returns non-empty -- single item", actual)
}

func Test_MapAnyItems_Nil(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems

	// Act
	actual := args.Map{
		"isEmpty": m.IsEmpty(),
		"length":  m.Length(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
		"length":  0,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems nil receiver safe -- isEmpty and length", actual)
}

// ── LeftRight ──

func Test_LeftRight_Basic(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "left", Right: "right"}

	// Act
	actual := args.Map{
		"isEmpty":  lr.IsEmpty(),
		"hasAny":   lr.HasAnyItem(),
		"hasLeft":  lr.HasLeft(),
		"hasRight": lr.HasRight(),
	}

	// Assert
	expected := args.Map{
		"isEmpty":  false,
		"hasAny":   true,
		"hasLeft":  true,
		"hasRight": true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight basic -- both set", actual)
}

func Test_LeftRight_Empty(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{}

	// Act
	actual := args.Map{"isEmpty": lr.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "LeftRight.IsEmpty returns true -- no values", actual)
}
