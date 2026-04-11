package corestrtests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// AnyToString
// ═══════════════════════════════════════════

func Test_AnyToString_Empty(t *testing.T) {
	safeTest(t, "Test_Cov9_AnyToString_Empty", func() {
		// Act
		actual := args.Map{"result": corestr.AnyToString(false, "")}

		// Assert
		expected := args.Map{"result": ""}
		expected.ShouldBeEqual(t, 0, "AnyToString returns empty -- empty", actual)
	})
}

func Test_AnyToString_WithFieldName(t *testing.T) {
	safeTest(t, "Test_Cov9_AnyToString_WithFieldName", func() {
		// Arrange
		result := corestr.AnyToString(true, "hello")

		// Act
		actual := args.Map{"notEmpty": result != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- with field name", actual)
	})
}

func Test_AnyToString_WithoutFieldName(t *testing.T) {
	safeTest(t, "Test_Cov9_AnyToString_WithoutFieldName", func() {
		// Arrange
		result := corestr.AnyToString(false, "hello")

		// Act
		actual := args.Map{"notEmpty": result != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- without field name", actual)
	})
}

func Test_AnyToString_Pointer(t *testing.T) {
	safeTest(t, "Test_Cov9_AnyToString_Pointer", func() {
		// Arrange
		val := "hello"
		result := corestr.AnyToString(false, &val)

		// Act
		actual := args.Map{"notEmpty": result != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns correct value -- pointer", actual)
	})
}

// ═══════════════════════════════════════════
// CloneSlice / CloneSliceIf
// ═══════════════════════════════════════════

func Test_CloneSlice(t *testing.T) {
	safeTest(t, "Test_Cov9_CloneSlice", func() {
		// Act
		actual := args.Map{
			"nilLen":   len(corestr.CloneSlice(nil)),
			"emptyLen": len(corestr.CloneSlice([]string{})),
			"dataLen":  len(corestr.CloneSlice([]string{"a", "b"})),
		}

		// Assert
		expected := args.Map{
			"nilLen": 0,
			"emptyLen": 0,
			"dataLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns correct value -- with args", actual)
	})
}

func Test_CloneSliceIf_Clone(t *testing.T) {
	safeTest(t, "Test_Cov9_CloneSliceIf_Clone", func() {
		// Arrange
		result := corestr.CloneSliceIf(true, "a", "b")

		// Act
		actual := args.Map{
			"len": len(result),
			"first": result[0],
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns correct value -- clone", actual)
	})
}

func Test_CloneSliceIf_NoClone(t *testing.T) {
	safeTest(t, "Test_Cov9_CloneSliceIf_NoClone", func() {
		// Arrange
		result := corestr.CloneSliceIf(false, "a", "b")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns empty -- no clone", actual)
	})
}

func Test_CloneSliceIf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov9_CloneSliceIf_Empty", func() {
		// Arrange
		result := corestr.CloneSliceIf(true)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns empty -- empty", actual)
	})
}

// ═══════════════════════════════════════════
// Collection — comprehensive
// ═══════════════════════════════════════════

func Test_Collection_Basic(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_Basic", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a").Add("b").Add("c")

		// Act
		actual := args.Map{
			"len":      c.Length(),
			"count":    c.Count(),
			"cap":      c.Capacity() >= 5,
			"hasAny":   c.HasAnyItem(),
			"hasItems": c.HasItems(),
			"isEmpty":  c.IsEmpty(),
			"lastIdx":  c.LastIndex(),
			"hasIdx1":  c.HasIndex(1),
			"hasIdx5":  c.HasIndex(5),
		}

		// Assert
		expected := args.Map{
			"len": 3, "count": 3, "cap": true,
			"hasAny": true, "hasItems": true, "isEmpty": false,
			"lastIdx": 2, "hasIdx1": true, "hasIdx5": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- basic", actual)
	})
}

func Test_Collection_NilReceiver(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_NilReceiver", func() {
		// Arrange
		var c *corestr.Collection

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
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- nil receiver", actual)
	})
}

func Test_Collection_AddVariations(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_AddVariations", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.AddNonEmpty("a")
		c.AddNonEmpty("")
		c.AddNonEmptyWhitespace("b")
		c.AddNonEmptyWhitespace("  ")
		c.AddIf(true, "c")
		c.AddIf(false, "x")
		c.AddIfMany(true, "d", "e")
		c.AddIfMany(false, "y", "z")
		c.AddFunc(func() string { return "f" })
		c.AddError(errors.New("err"))
		c.AddError(nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 7} // a, b, c, d, e, f, err
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- add variations", actual)
	})
}

func Test_Collection_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_AddFuncErr", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
		c.AddFuncErr(func() (string, error) { return "", errors.New("fail") }, func(e error) {})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- AddFuncErr", actual)
	})
}

func Test_Collection_Adds(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_Adds", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b")
		c.AddStrings([]string{"c", "d"})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Adds", actual)
	})
}

func Test_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_RemoveAt", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b", "c")
		ok := c.RemoveAt(1)
		failNeg := c.RemoveAt(-1)
		failBig := c.RemoveAt(100)

		// Act
		actual := args.Map{
			"ok": ok,
			"failNeg": failNeg,
			"failBig": failBig,
			"len": c.Length(),
		}

		// Assert
		expected := args.Map{
			"ok": true,
			"failNeg": false,
			"failBig": false,
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- RemoveAt", actual)
	})
}

func Test_Collection_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_IsEquals", func() {
		// Arrange
		c1 := corestr.New.Collection.Cap(5)
		c1.Adds("a", "b")
		c2 := corestr.New.Collection.Cap(5)
		c2.Adds("a", "b")
		c3 := corestr.New.Collection.Cap(5)
		c3.Adds("a", "c")

		// Act
		actual := args.Map{
			"equal":      c1.IsEquals(c2),
			"notEqual":   c1.IsEquals(c3),
			"insensitive": c1.IsEqualsWithSensitive(false, c2),
		}

		// Assert
		expected := args.Map{
			"equal": true,
			"notEqual": false,
			"insensitive": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- IsEquals", actual)
	})
}

func Test_Collection_ListStrings(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_ListStrings", func() {
		// Arrange
		c := corestr.New.Collection.Cap(3)
		c.Adds("a", "b")
		items := c.ListStrings()
		itemsPtr := c.ListStringsPtr()

		// Act
		actual := args.Map{
			"len": len(items),
			"ptrLen": len(itemsPtr),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"ptrLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ListStrings", actual)
	})
}

func Test_Collection_AsError(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_AsError", func() {
		// Arrange
		c := corestr.New.Collection.Cap(3)
		nilErr := c.AsError("\n")
		c.Add("err1")
		hasErr := c.AsError("\n")
		defErr := c.AsDefaultError()

		// Act
		actual := args.Map{
			"nilErr": nilErr == nil,
			"hasErr": hasErr != nil,
			"defErr": defErr != nil,
		}

		// Assert
		expected := args.Map{
			"nilErr": true,
			"hasErr": true,
			"defErr": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- AsError", actual)
	})
}

func Test_Collection_ToError(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_ToError", func() {
		// Arrange
		c := corestr.New.Collection.Cap(3)
		c.Adds("e1", "e2")
		err := c.ToError("\n")
		defErr := c.ToDefaultError()

		// Act
		actual := args.Map{
			"errNotNil": err != nil,
			"defNotNil": defErr != nil,
		}

		// Assert
		expected := args.Map{
			"errNotNil": true,
			"defNotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- ToError", actual)
	})
}

func Test_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_EachItemSplitBy", func() {
		// Arrange
		c := corestr.New.Collection.Cap(3)
		c.Add("a,b").Add("c,d")
		result := c.EachItemSplitBy(",")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- EachItemSplitBy", actual)
	})
}

func Test_Collection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_ConcatNew", func() {
		// Arrange
		c := corestr.New.Collection.Cap(3)
		c.Adds("a", "b")
		newC := c.ConcatNew(0, "c", "d")

		// Act
		actual := args.Map{
			"origLen": c.Length(),
			"newLen": newC.Length(),
		}

		// Assert
		expected := args.Map{
			"origLen": 2,
			"newLen": 4,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ConcatNew", actual)
	})
}

func Test_Collection_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_ConcatNew_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(3)
		c.Adds("a")
		newC := c.ConcatNew(0)

		// Act
		actual := args.Map{"len": newC.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- ConcatNew empty", actual)
	})
}

func Test_Collection_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_AddCollection", func() {
		// Arrange
		c1 := corestr.New.Collection.Cap(3)
		c1.Add("a")
		c2 := corestr.New.Collection.Cap(3)
		c2.Adds("b", "c")
		c1.AddCollection(c2)

		// Act
		actual := args.Map{"len": c1.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddCollection", actual)
	})
}

func Test_Collection_AddCollections(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_AddCollections", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c1 := corestr.New.Collection.Cap(3)
		c1.Add("a")
		c2 := corestr.New.Collection.Cap(3)
		c2.Add("b")
		c.AddCollections(c1, c2)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddCollections", actual)
	})
}

func Test_Collection_LockVariations(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_LockVariations", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddLock("a")
		c.AddsLock("b", "c")
		lenLock := c.LengthLock()
		emptyLock := c.IsEmptyLock()

		// Act
		actual := args.Map{
			"len": lenLock,
			"empty": emptyLock,
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"empty": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- lock variations", actual)
	})
}

func Test_Collection_Json(t *testing.T) {
	safeTest(t, "Test_Cov9_Collection_Json", func() {
		// Arrange
		c := corestr.New.Collection.Cap(3)
		c.Adds("a", "b")
		jsonStr := c.JsonString()
		jsonStrMust := c.JsonStringMust()
		strJSON := c.StringJSON()

		// Act
		actual := args.Map{
			"jsonNotEmpty":     jsonStr != "",
			"jsonMustNotEmpty": jsonStrMust != "",
			"strJSONNotEmpty":  strJSON != "",
		}

		// Assert
		expected := args.Map{
			"jsonNotEmpty": true, "jsonMustNotEmpty": true, "strJSONNotEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JSON", actual)
	})
}

// ═══════════════════════════════════════════
// Hashmap — comprehensive
// ═══════════════════════════════════════════

func Test_Hashmap_Basic(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashmap_Basic", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("k1", "v1")
		h.Set("k2", "v2")

		// Act
		actual := args.Map{
			"len":     h.Length(),
			"hasK1":   h.Has("k1"),
			"hasK3":   h.Has("k3"),
			"missing": h.IsKeyMissing("k3"),
			"empty":   h.IsEmpty(),
			"hasItems": h.HasItems(),
		}

		// Assert
		expected := args.Map{
			"len": 2, "hasK1": true, "hasK3": false,
			"missing": true, "empty": false, "hasItems": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- basic", actual)
	})
}

func Test_Hashmap_Get(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashmap_Get", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("k1", "v1")
		v, found := h.Get("k1")
		_, notFound := h.Get("k2")

		// Act
		actual := args.Map{
			"v": v,
			"found": found,
			"notFound": notFound,
		}

		// Assert
		expected := args.Map{
			"v": "v1",
			"found": true,
			"notFound": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Get", actual)
	})
}

func Test_Hashmap_Contains(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashmap_Contains", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(3)
		h.Set("k1", "v1")

		// Act
		actual := args.Map{
			"contains":     h.Contains("k1"),
			"containsLock": h.ContainsLock("k1"),
			"hasLock":      h.HasLock("k1"),
			"missingLock":  h.IsKeyMissingLock("k2"),
		}

		// Assert
		expected := args.Map{
			"contains": true, "containsLock": true,
			"hasLock": true, "missingLock": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Contains", actual)
	})
}

func Test_Hashmap_SetTrim(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashmap_SetTrim", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(3)
		h.SetTrim("  key  ", "  val  ")

		// Act
		actual := args.Map{"has": h.Has("key")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- SetTrim", actual)
	})
}

func Test_Hashmap_SetBySplitter(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashmap_SetBySplitter", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(3)
		h.SetBySplitter("=", "key=value")
		h.SetBySplitter("=", "novalue")
		v1, _ := h.Get("key")
		v2, _ := h.Get("novalue")

		// Act
		actual := args.Map{
			"v1": v1,
			"v2": v2,
		}

		// Assert
		expected := args.Map{
			"v1": "value",
			"v2": "",
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- SetBySplitter", actual)
	})
}

func Test_Hashmap_AddOrUpdateVariations(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashmap_AddOrUpdateVariations", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(10)
		h.AddOrUpdateKeyStrValInt("intKey", 42)
		h.AddOrUpdateKeyStrValFloat("floatKey", 3.14)
		h.AddOrUpdateKeyStrValFloat64("float64Key", 2.71)
		h.AddOrUpdateKeyStrValAny("anyKey", "anyVal")
		h.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "kvKey", Value: "kvVal"})
		h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "kavKey", Value: 99})

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 6}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdate variations", actual)
	})
}

func Test_Hashmap_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashmap_HasAllStrings", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.Set("a", "1")
		h.Set("b", "2")

		// Act
		actual := args.Map{
			"hasAll":  h.HasAllStrings("a", "b"),
			"missing": h.HasAllStrings("a", "c"),
		}

		// Assert
		expected := args.Map{
			"hasAll": true,
			"missing": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- HasAllStrings", actual)
	})
}

func Test_Hashmap_EmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashmap_EmptyLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(3)

		// Act
		actual := args.Map{"emptyLock": h.IsEmptyLock()}

		// Assert
		expected := args.Map{"emptyLock": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- IsEmptyLock", actual)
	})
}

func Test_Hashmap_AddOrUpdateLock(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashmap_AddOrUpdateLock", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(3)
		h.AddOrUpdateLock("k", "v")

		// Act
		actual := args.Map{"has": h.Has("k")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateLock", actual)
	})
}

func Test_Hashmap_AddOrUpdateMap(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashmap_AddOrUpdateMap", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdateMap(map[string]string{"a": "1", "b": "2"})
		h.AddOrUpdateMap(nil)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateMap", actual)
	})
}

func Test_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashmap_AddOrUpdateHashmap", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(3)
		h1.Set("a", "1")
		h2 := corestr.New.Hashmap.Cap(3)
		h2.Set("b", "2")
		h1.AddOrUpdateHashmap(h2)
		h1.AddOrUpdateHashmap(nil)

		// Act
		actual := args.Map{"len": h1.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateHashmap", actual)
	})
}

func Test_Hashmap_Collection(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashmap_Collection", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(3)
		h.Set("a", "1")
		c := h.Collection()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Collection", actual)
	})
}

func Test_Hashmap_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashmap_ConcatNew", func() {
		// Arrange
		h1 := corestr.New.Hashmap.Cap(3)
		h1.Set("a", "1")
		h2 := corestr.New.Hashmap.Cap(3)
		h2.Set("b", "2")
		newH := h1.ConcatNew(true, h2)

		// Act
		actual := args.Map{"newLen": newH.Length()}

		// Assert
		expected := args.Map{"newLen": 2}
		expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- ConcatNew", actual)
	})
}

func Test_Hashmap_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashmap_ConcatNew_Empty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(3)
		h.Set("a", "1")
		newH := h.ConcatNew(true)

		// Act
		actual := args.Map{"len": newH.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- ConcatNew empty", actual)
	})
}

// ═══════════════════════════════════════════
// Hashset — comprehensive
// ═══════════════════════════════════════════

func Test_Hashset_Basic(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_Basic", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.Add("a").Add("b")

		// Act
		actual := args.Map{
			"len":      h.Length(),
			"has":      h.Has("a"),
			"contains": h.Contains("a"),
			"empty":    h.IsEmpty(),
			"hasItems": h.HasItems(),
		}

		// Assert
		expected := args.Map{
			"len": 2, "has": true, "contains": true,
			"empty": false, "hasItems": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- basic", actual)
	})
}

func Test_Hashset_AddVariations(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_AddVariations", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(10)
		h.AddNonEmpty("a")
		h.AddNonEmpty("")
		h.AddNonEmptyWhitespace("b")
		h.AddNonEmptyWhitespace("  ")
		h.AddIf(true, "c")
		h.AddIf(false, "x")
		h.AddIfMany(true, "d", "e")
		h.AddIfMany(false, "y", "z")
		h.AddFunc(func() string { return "f" })
		h.AddBool("g")
		h.AddBool("g") // duplicate

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 7} // a,b,c,d,e,f,g = 7
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- add variations", actual)
	})
}

func Test_Hashset_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_AddFuncErr", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
		h.AddFuncErr(func() (string, error) { return "", errors.New("fail") }, func(e error) {})

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns error -- AddFuncErr", actual)
	})
}

func Test_Hashset_Adds(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_Adds", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.Adds("a", "b", "c")
		h.AddStrings([]string{"d", "e"})
		h.AddStrings(nil) // should not panic
		h.Adds()         // empty

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 5}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Adds", actual)
	})
}

func Test_Hashset_List(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_List", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(3)
		h.Adds("b", "a")
		list := h.List()
		sorted := h.SortedList()

		// Act
		actual := args.Map{
			"listLen": len(list),
			"sortedLen": len(sorted),
			"first": sorted[0],
		}

		// Assert
		expected := args.Map{
			"listLen": 2,
			"sortedLen": 2,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- List/SortedList", actual)
	})
}

func Test_Hashset_EmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_EmptyLock", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(3)

		// Act
		actual := args.Map{"emptyLock": h.IsEmptyLock()}

		// Assert
		expected := args.Map{"emptyLock": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- IsEmptyLock", actual)
	})
}

func Test_Hashset_Resize(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_Resize", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(2)
		h.Adds("a", "b")
		h.Resize(10)

		// Act
		actual := args.Map{
			"len": h.Length(),
			"has": h.Has("a"),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"has": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Resize", actual)
	})
}

func Test_Hashset_Collection(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_Collection", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(3)
		h.Add("a")
		c := h.Collection()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Collection", actual)
	})
}

func Test_Hashset_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_IsEquals", func() {
		// Arrange
		h1 := corestr.New.Hashset.Cap(3)
		h1.Adds("a", "b")
		h2 := corestr.New.Hashset.Cap(3)
		h2.Adds("a", "b")

		// Act
		actual := args.Map{"equal": h1.IsEquals(h2)}

		// Assert
		expected := args.Map{"equal": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsEquals", actual)
	})
}

func Test_Hashset_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_AddHashsetItems", func() {
		// Arrange
		h1 := corestr.New.Hashset.Cap(3)
		h1.Add("a")
		h2 := corestr.New.Hashset.Cap(3)
		h2.Add("b")
		h1.AddHashsetItems(h2)
		h1.AddHashsetItems(nil)

		// Act
		actual := args.Map{"len": h1.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddHashsetItems", actual)
	})
}

func Test_Hashset_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_AddCollection", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		c := corestr.New.Collection.Cap(3)
		c.Adds("a", "b")
		h.AddCollection(c)
		h.AddCollection(nil)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddCollection", actual)
	})
}

func Test_Hashset_AddPtr(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_AddPtr", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(3)
		key := "hello"
		h.AddPtr(&key)

		// Act
		actual := args.Map{"has": h.Has("hello")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddPtr", actual)
	})
}

func Test_Hashset_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_AddItemsMap", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.AddItemsMap(map[string]bool{"a": true, "b": false, "c": true})
		h.AddItemsMap(nil)

		// Act
		actual := args.Map{"len": h.Length()}

		// Assert
		expected := args.Map{"len": 2} // a, c only (b=false)
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddItemsMap", actual)
	})
}

func Test_Hashset_ConcatNewHashsets(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_ConcatNewHashsets", func() {
		// Arrange
		h1 := corestr.New.Hashset.Cap(3)
		h1.Add("a")
		h2 := corestr.New.Hashset.Cap(3)
		h2.Add("b")
		newH := h1.ConcatNewHashsets(true, h2)

		// Act
		actual := args.Map{"len": newH.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ConcatNewHashsets", actual)
	})
}

func Test_Hashset_ConcatNewHashsets_Empty(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_ConcatNewHashsets_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(3)
		h.Add("a")
		newH := h.ConcatNewHashsets(true)

		// Act
		actual := args.Map{"len": newH.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- ConcatNewHashsets empty", actual)
	})
}

func Test_Hashset_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_ConcatNewStrings", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(3)
		h.Add("a")
		newH := h.ConcatNewStrings(true, []string{"b", "c"})

		// Act
		actual := args.Map{"len": newH.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ConcatNewStrings", actual)
	})
}

func Test_Hashset_ConcatNewStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov9_Hashset_ConcatNewStrings_Empty", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(3)
		h.Add("a")
		newH := h.ConcatNewStrings(true)

		// Act
		actual := args.Map{"len": newH.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- ConcatNewStrings empty", actual)
	})
}

// ═══════════════════════════════════════════
// SimpleSlice — comprehensive
// ═══════════════════════════════════════════

func Test_SimpleSlice_Basic(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_Basic", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.Add("a").Add("b").Add("c")

		// Act
		actual := args.Map{
			"len":       s.Length(),
			"count":     s.Count(),
			"isEmpty":   s.IsEmpty(),
			"hasAny":    s.HasAnyItem(),
			"lastIdx":   s.LastIndex(),
			"hasIdx1":   s.HasIndex(1),
			"first":     s.First(),
			"last":      s.Last(),
			"firstDyn":  s.FirstDynamic(),
			"lastDyn":   s.LastDynamic(),
			"firstOrDef": s.FirstOrDefault(),
			"lastOrDef":  s.LastOrDefault(),
		}

		// Assert
		expected := args.Map{
			"len": 3, "count": 3, "isEmpty": false, "hasAny": true,
			"lastIdx": 2, "hasIdx1": true,
			"first": "a", "last": "c",
			"firstDyn": "a", "lastDyn": "c",
			"firstOrDef": "a", "lastOrDef": "c",
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- basic", actual)
	})
}

func Test_SimpleSlice_FirstLastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_FirstLastOrDefault_Empty", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{
			"firstOrDef":    s.FirstOrDefault(),
			"lastOrDef":     s.LastOrDefault(),
			"firstOrDefDyn": fmt.Sprintf("%v", s.FirstOrDefaultDynamic()),
			"lastOrDefDyn":  fmt.Sprintf("%v", s.LastOrDefaultDynamic()),
		}

		// Assert
		expected := args.Map{
			"firstOrDef": "", "lastOrDef": "",
			"firstOrDefDyn": "", "lastOrDefDyn": "",
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns empty -- empty defaults", actual)
	})
}

func Test_SimpleSlice_AddVariations(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_AddVariations", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AddIf(true, "a")
		s.AddIf(false, "x")
		s.Adds("b", "c")
		s.Adds() // empty
		s.AddsIf(true, "d")
		s.AddsIf(false, "y")
		s.Append("e")
		s.Append() // empty
		s.AddSplit("f,g", ",")
		s.AddError(errors.New("err"))
		s.AddError(nil)

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 8} // a,b,c,d,e,f,g,err
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- add variations", actual)
	})
}

func Test_SimpleSlice_AppendFmt(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_AppendFmt", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AppendFmt("hello %s", "world")
		s.AppendFmt("", ) // skipped (empty format, no values)
		s.AppendFmtIf(true, "yes %d", 1)
		s.AppendFmtIf(false, "no %d", 2)

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- AppendFmt", actual)
	})
}

func Test_SimpleSlice_InsertAt(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_InsertAt", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "c"}
		s.InsertAt(1, "b")
		s.InsertAt(-1, "x") // out of range, no change

		// Act
		actual := args.Map{
			"len": s.Length(),
			"idx1": s[1],
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"idx1": "b",
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- InsertAt", actual)
	})
}

func Test_SimpleSlice_SkipTake(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_SkipTake", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b", "c"}
		skip := s.Skip(1)
		take := s.Take(2)
		skipDyn := s.SkipDynamic(1)
		takeDyn := s.TakeDynamic(2)
		limit := s.Limit(2)
		limitDyn := s.LimitDynamic(2)

		// Act
		actual := args.Map{
			"skipLen": len(skip), "takeLen": len(take),
			"skipDynNotNil": skipDyn != nil, "takeDynNotNil": takeDyn != nil,
			"limitLen": len(limit), "limitDynNotNil": limitDyn != nil,
		}

		// Assert
		expected := args.Map{
			"skipLen": 2, "takeLen": 2,
			"skipDynNotNil": true, "takeDynNotNil": true,
			"limitLen": 2, "limitDynNotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- Skip/Take", actual)
	})
}

func Test_SimpleSlice_SkipAll(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_SkipAll", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		skip := s.Skip(5)
		skipDyn := s.SkipDynamic(5)

		// Act
		actual := args.Map{
			"skipLen": len(skip),
			"skipDynNotNil": skipDyn != nil,
		}

		// Assert
		expected := args.Map{
			"skipLen": 0,
			"skipDynNotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- Skip all", actual)
	})
}

func Test_SimpleSlice_TakeAll(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_TakeAll", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		take := s.Take(5)
		takeDyn := s.TakeDynamic(5)

		// Act
		actual := args.Map{
			"takeLen": len(take),
			"takeDynNotNil": takeDyn != nil,
		}

		// Assert
		expected := args.Map{
			"takeLen": 1,
			"takeDynNotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- Take all", actual)
	})
}

func Test_SimpleSlice_IndexOf(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_IndexOf", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b", "c"}

		// Act
		actual := args.Map{
			"found":    s.IndexOf("b"),
			"notFound": s.IndexOf("z"),
		}

		// Assert
		expected := args.Map{
			"found": 1,
			"notFound": -1,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- IndexOf", actual)
	})
}

func Test_SimpleSlice_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_IsContains", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{
			"contains":    s.IsContains("a"),
			"notContains": s.IsContains("z"),
		}

		// Assert
		expected := args.Map{
			"contains": true,
			"notContains": false,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- IsContains", actual)
	})
}

func Test_SimpleSlice_CountFunc(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_CountFunc", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "bb", "ccc"}
		count := s.CountFunc(func(i int, item string) bool { return len(item) > 1 })

		// Act
		actual := args.Map{"count": count}

		// Assert
		expected := args.Map{"count": 2}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- CountFunc", actual)
	})
}

func Test_SimpleSlice_AsError(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_AsError", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		nilErr := s.AsError("\n")
		s.Add("err1")
		hasErr := s.AsDefaultError()

		// Act
		actual := args.Map{
			"nilErr": nilErr == nil,
			"hasErr": hasErr != nil,
		}

		// Assert
		expected := args.Map{
			"nilErr": true,
			"hasErr": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns error -- AsError", actual)
	})
}

func Test_SimpleSlice_WrapQuotes(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_WrapQuotes", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		d := s.WrapDoubleQuote()

		// Act
		actual := args.Map{"wrapped": (*d)[0]}

		// Assert
		expected := args.Map{"wrapped": `"a"`}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- WrapDoubleQuote", actual)
	})
}

func Test_SimpleSlice_Strings(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_Strings", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}
		strs := s.Strings()
		list := s.List()

		// Act
		actual := args.Map{
			"strsLen": len(strs),
			"listLen": len(list),
		}

		// Assert
		expected := args.Map{
			"strsLen": 2,
			"listLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- Strings/List", actual)
	})
}

func Test_SimpleSlice_Nil(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_Nil", func() {
		// Arrange
		var s *corestr.SimpleSlice

		// Act
		actual := args.Map{
			"len": s.Length(),
			"isEmpty": s.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns nil -- nil", actual)
	})
}

func Test_SimpleSlice_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_AddAsTitleValue", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AddAsTitleValue("Key", "Value")
		s.AddAsTitleValueIf(true, "K2", "V2")
		s.AddAsTitleValueIf(false, "K3", "V3")
		s.AddAsCurlyTitleWrap("K4", "V4")
		s.AddAsCurlyTitleWrapIf(true, "K5", "V5")
		s.AddAsCurlyTitleWrapIf(false, "K6", "V6")

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- title value", actual)
	})
}

func Test_SimpleSlice_AddStruct(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_AddStruct", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AddStruct(false, "hello")
		s.AddStruct(false, nil)
		s.AddPointer(false, "world")
		s.AddPointer(false, nil)

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- AddStruct/AddPointer", actual)
	})
}

func Test_SimpleSlice_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_IsContainsFunc", func() {
		// Arrange
		s := corestr.SimpleSlice{"abc", "def"}
		found := s.IsContainsFunc("abc", func(item, searching string) bool { return item == searching })
		notFound := s.IsContainsFunc("xyz", func(item, searching string) bool { return item == searching })

		// Act
		actual := args.Map{
			"found": found,
			"notFound": notFound,
		}

		// Assert
		expected := args.Map{
			"found": true,
			"notFound": false,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- IsContainsFunc", actual)
	})
}

func Test_SimpleSlice_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleSlice_IndexOfFunc", func() {
		// Arrange
		s := corestr.SimpleSlice{"abc", "def"}
		found := s.IndexOfFunc("def", func(item, searching string) bool { return item == searching })
		notFound := s.IndexOfFunc("xyz", func(item, searching string) bool { return item == searching })

		// Act
		actual := args.Map{
			"found": found,
			"notFound": notFound,
		}

		// Assert
		expected := args.Map{
			"found": 1,
			"notFound": -1,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- IndexOfFunc", actual)
	})
}

// ═══════════════════════════════════════════
// KeyValuePair
// ═══════════════════════════════════════════

func Test_KeyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov9_KeyValuePair", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "name", Value: "42"}

		// Act
		actual := args.Map{
			"keyName":     kv.KeyName(),
			"varName":     kv.VariableName(),
			"valueStr":    kv.ValueString(),
			"isVarEq":     kv.IsVariableNameEqual("name"),
			"isValEq":     kv.IsValueEqual("42"),
			"isKeyEmpty":  kv.IsKeyEmpty(),
			"isValEmpty":  kv.IsValueEmpty(),
			"hasKey":      kv.HasKey(),
			"hasValue":    kv.HasValue(),
			"isKVEmpty":   kv.IsKeyValueEmpty(),
			"isKVAnyEmp":  kv.IsKeyValueAnyEmpty(),
			"isKey":       kv.IsKey("name"),
			"isVal":       kv.IsVal("42"),
			"is":          kv.Is("name", "42"),
			"trimKey":     kv.TrimKey(),
			"trimValue":   kv.TrimValue(),
			"valueBool":   kv.ValueBool(),
			"valueInt":    kv.ValueInt(0),
			"valueDefInt": kv.ValueDefInt(),
			"valueByte":   kv.ValueByte(0),
			"defByte":     kv.ValueDefByte(),
			"valFloat":    kv.ValueFloat64(0.0),
			"defFloat":    kv.ValueDefFloat64(),
		}

		// Assert
		expected := args.Map{
			"keyName": "name", "varName": "name",
			"valueStr": "42", "isVarEq": true,
			"isValEq": true, "isKeyEmpty": false,
			"isValEmpty": false, "hasKey": true,
			"hasValue": true, "isKVEmpty": false,
			"isKVAnyEmp": false, "isKey": true, "isVal": true,
			"is": true, "trimKey": "name", "trimValue": "42",
			"valueBool": false, "valueInt": 42, "valueDefInt": 42,
			"valueByte": byte(42), "defByte": byte(42),
			"valFloat": 42.0, "defFloat": 42.0,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- with args", actual)
	})
}

func Test_KeyValuePair_Compile(t *testing.T) {
	safeTest(t, "Test_Cov9_KeyValuePair_Compile", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		compiled := kv.Compile()
		str := kv.String()
		format := kv.FormatString("%s=%s")

		// Act
		actual := args.Map{
			"compiled": compiled == str,
			"format":   format,
		}

		// Assert
		expected := args.Map{
			"compiled": true,
			"format": "k=v",
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- Compile", actual)
	})
}

func Test_KeyValuePair_ValueValid(t *testing.T) {
	safeTest(t, "Test_Cov9_KeyValuePair_ValueValid", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kv.ValueValid()
		vvo := kv.ValueValidOptions(false, "msg")

		// Act
		actual := args.Map{
			"vvIsValid": vv.IsValid,
			"vvoIsValid": vvo.IsValid,
			"vvoMsg":     vvo.Message,
		}

		// Assert
		expected := args.Map{
			"vvIsValid": true,
			"vvoIsValid": false,
			"vvoMsg": "msg",
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns non-empty -- ValueValid", actual)
	})
}

func Test_KeyValuePair_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov9_KeyValuePair_ClearDispose", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Clear()

		// Act
		actual := args.Map{
			"key": kv.Key,
			"val": kv.Value,
		}

		// Assert
		expected := args.Map{
			"key": "",
			"val": "",
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- Clear", actual)
	})
}

func Test_KeyValuePair_NilReceiver(t *testing.T) {
	safeTest(t, "Test_Cov9_KeyValuePair_NilReceiver", func() {
		// Arrange
		var kv *corestr.KeyValuePair
		kv.Clear()
		kv.Dispose()

		// Act
		actual := args.Map{
			"isKVAnyEmpty": kv.IsKeyValueAnyEmpty(),
			"isKey":        kv.IsKey("x"),
			"isVal":        kv.IsVal("x"),
			"is":           kv.Is("x", "y"),
		}

		// Assert
		expected := args.Map{
			"isKVAnyEmpty": true, "isKey": false, "isVal": false, "is": false,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns nil -- nil receiver", actual)
	})
}

func Test_KeyValuePair_ValueBoolTrue(t *testing.T) {
	safeTest(t, "Test_Cov9_KeyValuePair_ValueBoolTrue", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "true"}

		// Act
		actual := args.Map{"bool": kv.ValueBool()}

		// Assert
		expected := args.Map{"bool": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns non-empty -- ValueBool true", actual)
	})
}

func Test_KeyValuePair_ValueBoolEmpty(t *testing.T) {
	safeTest(t, "Test_Cov9_KeyValuePair_ValueBoolEmpty", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: ""}

		// Act
		actual := args.Map{"bool": kv.ValueBool()}

		// Assert
		expected := args.Map{"bool": false}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns empty -- ValueBool empty", actual)
	})
}

// ═══════════════════════════════════════════
// ValidValue
// ═══════════════════════════════════════════

func Test_ValidValue_Methods(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValue_Methods", func() {
		// Arrange
		vv := corestr.NewValidValue("42")

		// Act
		actual := args.Map{
			"value":     vv.Value,
			"isValid":   vv.IsValid,
			"isEmpty":   vv.IsEmpty(),
			"isWS":      vv.IsWhitespace(),
			"trim":      vv.Trim(),
			"hasValid":  vv.HasValidNonEmpty(),
			"hasValidWS": vv.HasValidNonWhitespace(),
			"hasSafe":   vv.HasSafeNonEmpty(),
			"is42":      vv.Is("42"),
			"isAnyOf":   vv.IsAnyOf("42", "99"),
			"contains":  vv.IsContains("4"),
			"anyContains": vv.IsAnyContains("4"),
			"nonSensitive": vv.IsEqualNonSensitive("42"),
			"valBool":   vv.ValueBool(),
			"valInt":    vv.ValueInt(0),
			"defInt":    vv.ValueDefInt(),
			"valByte":   vv.ValueByte(0),
			"defByte":   vv.ValueDefByte(),
			"valFloat":  vv.ValueFloat64(0.0),
			"defFloat":  vv.ValueDefFloat64(),
			"bytesLen":  len(vv.ValueBytesOnce()),
			"bytesLen2": len(vv.ValueBytesOnce()),
		}

		// Assert
		expected := args.Map{
			"value": "42", "isValid": true, "isEmpty": false,
			"isWS": false, "trim": "42",
			"hasValid": true, "hasValidWS": true, "hasSafe": true,
			"is42": true, "isAnyOf": true,
			"contains": true, "anyContains": true,
			"nonSensitive": true,
			"valBool": false, "valInt": 42, "defInt": 42,
			"valByte": byte(42), "defByte": byte(42),
			"valFloat": 42.0, "defFloat": 42.0,
			"bytesLen": 2, "bytesLen2": 2,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- methods", actual)
	})
}

func Test_ValidValue_Clone(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValue_Clone", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		cloned := vv.Clone()

		// Act
		actual := args.Map{
			"val": cloned.Value,
			"isValid": cloned.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"isValid": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Clone", actual)
	})
}

func Test_ValidValue_NilClone(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValue_NilClone", func() {
		// Arrange
		var vv *corestr.ValidValue
		cloned := vv.Clone()

		// Act
		actual := args.Map{"isNil": cloned == nil}

		// Assert
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns nil -- nil Clone", actual)
	})
}

func Test_ValidValue_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValue_ClearDispose", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		vv.Clear()

		// Act
		actual := args.Map{
			"val": vv.Value,
			"isValid": vv.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "",
			"isValid": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Clear", actual)
	})
}

func Test_ValidValue_NilString(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValue_NilString", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act
		actual := args.Map{
			"str": vv.String(),
			"fullStr": vv.FullString(),
		}

		// Assert
		expected := args.Map{
			"str": "",
			"fullStr": "",
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns nil -- nil String", actual)
	})
}

func Test_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValue_Split", func() {
		// Arrange
		vv := corestr.NewValidValue("a,b,c")
		parts := vv.Split(",")

		// Act
		actual := args.Map{"len": len(parts)}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Split", actual)
	})
}

func Test_ValidValue_IsAnyOf_EmptyValues(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValue_IsAnyOf_EmptyValues", func() {
		// Arrange
		vv := corestr.NewValidValue("anything")

		// Act
		actual := args.Map{"result": vv.IsAnyOf()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns empty -- IsAnyOf empty values", actual)
	})
}

func Test_ValidValue_IsAnyContains_EmptyValues(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValue_IsAnyContains_EmptyValues", func() {
		// Arrange
		vv := corestr.NewValidValue("anything")

		// Act
		actual := args.Map{"result": vv.IsAnyContains()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns empty -- IsAnyContains empty values", actual)
	})
}

func Test_ValidValue_RegexNil(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValue_RegexNil", func() {
		// Arrange
		vv := corestr.NewValidValue("test")

		// Act
		actual := args.Map{
			"matches":    vv.IsRegexMatches(nil),
			"findStr":    vv.RegexFindString(nil),
			"findAllLen": len(vv.RegexFindAllStrings(nil, -1)),
		}

		// Assert
		expected := args.Map{
			"matches": false,
			"findStr": "",
			"findAllLen": 0,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns nil -- regex nil", actual)
	})
}

func Test_ValidValue_Constructors(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValue_Constructors", func() {
		// Arrange
		empty := corestr.NewValidValueEmpty()
		invalid := corestr.InvalidValidValueNoMessage()
		usingAny := corestr.NewValidValueUsingAny(false, true, "hello")
		usingAnyAuto := corestr.NewValidValueUsingAnyAutoValid(false, "hello")

		// Act
		actual := args.Map{
			"emptyIsValid":   empty.IsValid,
			"invalidIsValid": invalid.IsValid,
			"usingAnyValue":  usingAny.Value != "",
			"usingAutoValue": usingAnyAuto.Value != "",
		}

		// Assert
		expected := args.Map{
			"emptyIsValid": true, "invalidIsValid": false,
			"usingAnyValue": true, "usingAutoValue": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- constructors", actual)
	})
}

// ═══════════════════════════════════════════
// TextWithLineNumber
// ═══════════════════════════════════════════

func Test_TextWithLineNumber(t *testing.T) {
	safeTest(t, "Test_Cov9_TextWithLineNumber", func() {
		// Arrange
		tl := corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}

		// Act
		actual := args.Map{
			"hasLine":     tl.HasLineNumber(),
			"isInvalidLn": tl.IsInvalidLineNumber(),
			"len":         tl.Length(),
			"isEmpty":     tl.IsEmpty(),
			"isEmptyText": tl.IsEmptyText(),
			"isEmptyBoth": tl.IsEmptyTextLineBoth(),
		}

		// Assert
		expected := args.Map{
			"hasLine": true, "isInvalidLn": false,
			"len": 5, "isEmpty": false, "isEmptyText": false,
			"isEmptyBoth": false,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns non-empty -- with args", actual)
	})
}

func Test_TextWithLineNumber_Nil(t *testing.T) {
	safeTest(t, "Test_Cov9_TextWithLineNumber_Nil", func() {
		// Arrange
		var tl *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"hasLine":     tl.HasLineNumber(),
			"isInvalidLn": tl.IsInvalidLineNumber(),
			"len":         tl.Length(),
			"isEmpty":     tl.IsEmpty(),
			"isEmptyText": tl.IsEmptyText(),
		}

		// Assert
		expected := args.Map{
			"hasLine": false, "isInvalidLn": true,
			"len": 0, "isEmpty": true, "isEmptyText": true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns nil -- nil", actual)
	})
}

// ═══════════════════════════════════════════
// ValueStatus
// ═══════════════════════════════════════════

func Test_ValueStatus(t *testing.T) {
	safeTest(t, "Test_Cov9_ValueStatus", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("bad")
		vsNoMsg := corestr.InvalidValueStatusNoMessage()
		cloned := vs.Clone()

		// Act
		actual := args.Map{
			"msg":        vs.ValueValid.Message,
			"noMsgValid": vsNoMsg.ValueValid.IsValid,
			"clonedIdx":  cloned.Index,
		}

		// Assert
		expected := args.Map{
			"msg": "bad", "noMsgValid": false, "clonedIdx": -1,
		}
		expected.ShouldBeEqual(t, 0, "ValueStatus returns non-empty -- with args", actual)
	})
}

// ═══════════════════════════════════════════
// HashmapDiff
// ═══════════════════════════════════════════

func Test_HashmapDiff_Basic(t *testing.T) {
	safeTest(t, "Test_Cov9_HashmapDiff_Basic", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1", "b": "2"}

		// Act
		actual := args.Map{
			"len":     hd.Length(),
			"isEmpty": hd.IsEmpty(),
			"hasAny":  hd.HasAnyItem(),
			"lastIdx": hd.LastIndex(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"isEmpty": false,
			"hasAny": true,
			"lastIdx": 1,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- basic", actual)
	})
}

func Test_HashmapDiff_Nil(t *testing.T) {
	safeTest(t, "Test_Cov9_HashmapDiff_Nil", func() {
		// Arrange
		var hd *corestr.HashmapDiff

		// Act
		actual := args.Map{
			"len": hd.Length(),
			"rawLen": len(hd.Raw()),
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"rawLen": 0,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns nil -- nil", actual)
	})
}

func Test_HashmapDiff_MapAnyItems(t *testing.T) {
	safeTest(t, "Test_Cov9_HashmapDiff_MapAnyItems", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1"}
		m := hd.MapAnyItems()

		// Act
		actual := args.Map{"len": len(m)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- MapAnyItems", actual)
	})
}

func Test_HashmapDiff_MapAnyItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov9_HashmapDiff_MapAnyItems_Nil", func() {
		// Arrange
		var hd *corestr.HashmapDiff
		m := hd.MapAnyItems()

		// Act
		actual := args.Map{"len": len(m)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns nil -- MapAnyItems nil", actual)
	})
}

func Test_HashmapDiff_Diff(t *testing.T) {
	safeTest(t, "Test_Cov9_HashmapDiff_Diff", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1", "b": "2"}
		isEqual := hd.IsRawEqual(map[string]string{"a": "1", "b": "2"})
		hasDiff := hd.HasAnyChanges(map[string]string{"a": "1", "b": "3"})
		diffMap := hd.DiffRaw(map[string]string{"a": "1", "b": "3"})
		diffHm := hd.HashmapDiffUsingRaw(map[string]string{"a": "1", "b": "3"})

		// Act
		actual := args.Map{
			"isEqual":    isEqual,
			"hasDiff":    hasDiff,
			"diffMapLen": len(diffMap),
			"diffHmLen":  diffHm.Length(),
		}

		// Assert
		expected := args.Map{
			"isEqual": true,
			"hasDiff": true,
			"diffMapLen": 1,
			"diffHmLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- Diff", actual)
	})
}

func Test_HashmapDiff_AllKeysSorted(t *testing.T) {
	safeTest(t, "Test_Cov9_HashmapDiff_AllKeysSorted", func() {
		// Arrange
		hd := corestr.HashmapDiff{"b": "2", "a": "1"}
		keys := hd.AllKeysSorted()

		// Act
		actual := args.Map{
			"first": keys[0],
			"second": keys[1],
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"second": "b",
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- AllKeysSorted", actual)
	})
}

// ═══════════════════════════════════════════
// LeftRight additional
// ═══════════════════════════════════════════

func Test_LeftRight_StringMethods(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftRight_StringMethods", func() {
		// Arrange
		lr := corestr.NewLeftRight("  left  ", "  right  ")

		// Act
		actual := args.Map{
			"leftTrim":          lr.LeftTrim(),
			"rightTrim":         lr.RightTrim(),
			"leftBytes":         len(lr.LeftBytes()),
			"rightBytes":        len(lr.RightBytes()),
			"isLeftEmpty":       lr.IsLeftEmpty(),
			"isRightEmpty":      lr.IsRightEmpty(),
			"isLeftWS":          lr.IsLeftWhitespace(),
			"isRightWS":         lr.IsRightWhitespace(),
			"hasValidLeft":      lr.HasValidNonEmptyLeft(),
			"hasValidRight":     lr.HasValidNonEmptyRight(),
			"hasValidWSLeft":    lr.HasValidNonWhitespaceLeft(),
			"hasValidWSRight":   lr.HasValidNonWhitespaceRight(),
			"hasSafeNonEmpty":   lr.HasSafeNonEmpty(),
			"isLeft":            lr.IsLeft("  left  "),
			"isRight":           lr.IsRight("  right  "),
			"is":                lr.Is("  left  ", "  right  "),
		}

		// Assert
		expected := args.Map{
			"leftTrim": "left", "rightTrim": "right",
			"leftBytes": 8, "rightBytes": 9,
			"isLeftEmpty": false, "isRightEmpty": false,
			"isLeftWS": false, "isRightWS": false,
			"hasValidLeft": true, "hasValidRight": true,
			"hasValidWSLeft": true, "hasValidWSRight": true,
			"hasSafeNonEmpty": true,
			"isLeft": true, "isRight": true, "is": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- string methods", actual)
	})
}

func Test_LeftRight_RegexMatch(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftRight_RegexMatch", func() {
		// Arrange
		lr := corestr.NewLeftRight("abc123", "def456")

		// Act
		actual := args.Map{
			"leftNilRegex":  lr.IsLeftRegexMatch(nil),
			"rightNilRegex": lr.IsRightRegexMatch(nil),
		}

		// Assert
		expected := args.Map{
			"leftNilRegex": false,
			"rightNilRegex": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns nil -- regex nil", actual)
	})
}

func Test_LeftRight_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftRight_IsEqual", func() {
		// Arrange
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		lr3 := corestr.NewLeftRight("x", "y")

		// Act
		actual := args.Map{
			"equal":    lr1.IsEqual(lr2),
			"notEqual": lr1.IsEqual(lr3),
			"bothNil":  (*corestr.LeftRight)(nil).IsEqual(nil),
			"oneNil":   lr1.IsEqual(nil),
		}

		// Assert
		expected := args.Map{
			"equal": true, "notEqual": false, "bothNil": true, "oneNil": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- IsEqual", actual)
	})
}

func Test_LeftRight_Clone(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftRight_Clone", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		cloned := lr.Clone()

		// Act
		actual := args.Map{
			"left": cloned.Left,
			"right": cloned.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Clone", actual)
	})
}

func Test_LeftRight_NonPtrPtr(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftRight_NonPtrPtr", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		nonPtr := lr.NonPtr()
		ptr := lr.Ptr()

		// Act
		actual := args.Map{
			"nonPtrLeft": nonPtr.Left,
			"ptrNotNil": ptr != nil,
		}

		// Assert
		expected := args.Map{
			"nonPtrLeft": "a",
			"ptrNotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- NonPtr/Ptr", actual)
	})
}

func Test_LeftRight_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftRight_ClearDispose", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		lr.Clear()

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "",
			"right": "",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Clear", actual)
	})
}

func Test_LeftRight_InvalidConstructors(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftRight_InvalidConstructors", func() {
		// Arrange
		lr1 := corestr.InvalidLeftRightNoMessage()
		lr2 := corestr.InvalidLeftRight("bad")

		// Act
		actual := args.Map{
			"valid1": lr1.IsValid,
			"valid2": lr2.IsValid,
			"msg": lr2.Message,
		}

		// Assert
		expected := args.Map{
			"valid1": false,
			"valid2": false,
			"msg": "bad",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns error -- invalid constructors", actual)
	})
}

func Test_LeftRight_UsingSlice(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftRight_UsingSlice", func() {
		// Arrange
		lr0 := corestr.LeftRightUsingSlice(nil)
		lr1 := corestr.LeftRightUsingSlice([]string{"only"})
		lr2 := corestr.LeftRightUsingSlice([]string{"a", "b"})

		// Act
		actual := args.Map{
			"lr0Valid": lr0.IsValid,
			"lr1Left":  lr1.Left,
			"lr1Valid": lr1.IsValid,
			"lr2Left":  lr2.Left,
			"lr2Right": lr2.Right,
			"lr2Valid": lr2.IsValid,
		}

		// Assert
		expected := args.Map{
			"lr0Valid": false,
			"lr1Left": "only", "lr1Valid": false,
			"lr2Left": "a", "lr2Right": "b", "lr2Valid": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- UsingSlice", actual)
	})
}

func Test_LeftRight_TrimmedUsingSlice(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftRight_TrimmedUsingSlice", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{"  a  ", "  b  "})
		lrNil := corestr.LeftRightTrimmedUsingSlice(nil)
		lrEmpty := corestr.LeftRightTrimmedUsingSlice([]string{})
		lr1 := corestr.LeftRightTrimmedUsingSlice([]string{"only"})

		// Act
		actual := args.Map{
			"left": lr.Left, "right": lr.Right,
			"nilValid": lrNil.IsValid, "emptyValid": lrEmpty.IsValid,
			"lr1Left": lr1.Left, "lr1Valid": lr1.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a", "right": "b",
			"nilValid": false, "emptyValid": false,
			"lr1Left": "only", "lr1Valid": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- TrimmedUsingSlice", actual)
	})
}

// ═══════════════════════════════════════════
// LeftMiddleRight additional
// ═══════════════════════════════════════════

func Test_LeftMiddleRight_Methods(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftMiddleRight_Methods", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("  l  ", "  m  ", "  r  ")

		// Act
		actual := args.Map{
			"leftTrim":       lmr.LeftTrim(),
			"middleTrim":     lmr.MiddleTrim(),
			"rightTrim":      lmr.RightTrim(),
			"leftBytes":      len(lmr.LeftBytes()),
			"middleBytes":    len(lmr.MiddleBytes()),
			"rightBytes":     len(lmr.RightBytes()),
			"isLeftEmpty":    lmr.IsLeftEmpty(),
			"isMiddleEmpty":  lmr.IsMiddleEmpty(),
			"isRightEmpty":   lmr.IsRightEmpty(),
			"isLeftWS":       lmr.IsLeftWhitespace(),
			"isMiddleWS":     lmr.IsMiddleWhitespace(),
			"isRightWS":      lmr.IsRightWhitespace(),
			"hasValidLeft":   lmr.HasValidNonEmptyLeft(),
			"hasValidMiddle": lmr.HasValidNonEmptyMiddle(),
			"hasValidRight":  lmr.HasValidNonEmptyRight(),
			"hasValidWSL":    lmr.HasValidNonWhitespaceLeft(),
			"hasValidWSM":    lmr.HasValidNonWhitespaceMiddle(),
			"hasValidWSR":    lmr.HasValidNonWhitespaceRight(),
			"hasSafe":        lmr.HasSafeNonEmpty(),
			"isAll":          lmr.IsAll("  l  ", "  m  ", "  r  "),
			"is":             lmr.Is("  l  ", "  r  "),
		}

		// Assert
		expected := args.Map{
			"leftTrim": "l", "middleTrim": "m", "rightTrim": "r",
			"leftBytes": 5, "middleBytes": 5, "rightBytes": 5,
			"isLeftEmpty": false, "isMiddleEmpty": false, "isRightEmpty": false,
			"isLeftWS": false, "isMiddleWS": false, "isRightWS": false,
			"hasValidLeft": true, "hasValidMiddle": true, "hasValidRight": true,
			"hasValidWSL": true, "hasValidWSM": true, "hasValidWSR": true,
			"hasSafe": true, "isAll": true, "is": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- methods", actual)
	})
}

func Test_LeftMiddleRight_Clone(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftMiddleRight_Clone", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		cloned := lmr.Clone()

		// Act
		actual := args.Map{
			"left": cloned.Left,
			"middle": cloned.Middle,
			"right": cloned.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"middle": "b",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- Clone", actual)
	})
}

func Test_LeftMiddleRight_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftMiddleRight_ToLeftRight", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"isValid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "c",
			"isValid": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- ToLeftRight", actual)
	})
}

func Test_LeftMiddleRight_InvalidConstructors(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftMiddleRight_InvalidConstructors", func() {
		// Arrange
		lmr1 := corestr.InvalidLeftMiddleRightNoMessage()
		lmr2 := corestr.InvalidLeftMiddleRight("bad")

		// Act
		actual := args.Map{
			"valid1": lmr1.IsValid,
			"valid2": lmr2.IsValid,
		}

		// Assert
		expected := args.Map{
			"valid1": false,
			"valid2": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns error -- invalid constructors", actual)
	})
}

func Test_LeftMiddleRight_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftMiddleRight_ClearDispose", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Clear()

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"middle": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "",
			"middle": "",
			"right": "",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- Clear", actual)
	})
}

// ═══════════════════════════════════════════
// LeftRightFromSplit / LeftMiddleRightFromSplit
// ═══════════════════════════════════════════

func Test_LeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftRightFromSplit", func() {
		// Arrange
		lr := corestr.LeftRightFromSplit("key=value", "=")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "key",
			"right": "value",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit returns correct value -- with args", actual)
	})
}

func Test_LeftRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftRightFromSplitTrimmed", func() {
		// Arrange
		lr := corestr.LeftRightFromSplitTrimmed("  key  =  value  ", "=")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "key",
			"right": "value",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitTrimmed returns correct value -- with args", actual)
	})
}

func Test_LeftRightFromSplitFull(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftRightFromSplitFull", func() {
		// Arrange
		lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b:c:d",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitFull returns correct value -- with args", actual)
	})
}

func Test_LeftRightFromSplitFullTrimmed(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftRightFromSplitFullTrimmed", func() {
		// Arrange
		lr := corestr.LeftRightFromSplitFullTrimmed("  a  :  b:c  ", ":")

		// Act
		actual := args.Map{"left": lr.Left}

		// Assert
		expected := args.Map{"left": "a"}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitFullTrimmed returns correct value -- with args", actual)
	})
}

func Test_LeftMiddleRightFromSplit(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftMiddleRightFromSplit", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"middle": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"middle": "b",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplit returns correct value -- with args", actual)
	})
}

func Test_LeftMiddleRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftMiddleRightFromSplitTrimmed", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplitTrimmed("  a  .  b  .  c  ", ".")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"middle": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"middle": "b",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitTrimmed returns correct value -- with args", actual)
	})
}

func Test_LeftMiddleRightFromSplitN(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftMiddleRightFromSplitN", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"middle": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"middle": "b",
			"right": "c:d:e",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitN returns correct value -- with args", actual)
	})
}

func Test_LeftMiddleRightFromSplitNTrimmed(t *testing.T) {
	safeTest(t, "Test_Cov9_LeftMiddleRightFromSplitNTrimmed", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed("  a  :  b  :  c:d  ", ":")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"middle": lmr.Middle,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"middle": "b",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitNTrimmed returns correct value -- with args", actual)
	})
}

// ═══════════════════════════════════════════
// SimpleStringOnce
// ═══════════════════════════════════════════

func Test_SimpleStringOnce_Basic(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_Basic", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		err := sso.SetOnUninitialized("hello")

		// Act
		actual := args.Map{
			"err":      err == nil,
			"value":    sso.Value(),
			"isInit":   sso.IsInitialized(),
			"isDef":    sso.IsDefined(),
			"isUninit": sso.IsUninitialized(),
			"isEmpty":  sso.IsEmpty(),
			"isWS":     sso.IsWhitespace(),
			"trim":     sso.Trim(),
			"hasValid": sso.HasValidNonEmpty(),
			"hasWS":    sso.HasValidNonWhitespace(),
			"hasSafe":  sso.HasSafeNonEmpty(),
			"is":       sso.Is("hello"),
			"isAnyOf":  sso.IsAnyOf("hello", "world"),
			"contains": sso.IsContains("ell"),
			"anyCont":  sso.IsAnyContains("ell"),
			"nonSens":  sso.IsEqualNonSensitive("HELLO"),
			"safeVal":  sso.SafeValue(),
		}

		// Assert
		expected := args.Map{
			"err": true, "value": "hello", "isInit": true,
			"isDef": true, "isUninit": false, "isEmpty": false,
			"isWS": false, "trim": "hello",
			"hasValid": true, "hasWS": true, "hasSafe": true,
			"is": true, "isAnyOf": true, "contains": true,
			"anyCont": true, "nonSens": true, "safeVal": "hello",
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- basic", actual)
	})
}

func Test_SimpleStringOnce_SetOnUninitialized_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_SetOnUninitialized_AlreadyInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.SetOnUninitialized("first")
		err := sso.SetOnUninitialized("second")

		// Act
		actual := args.Map{
			"hasErr": err != nil,
			"val": sso.Value(),
		}

		// Assert
		expected := args.Map{
			"hasErr": true,
			"val": "first",
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- already init", actual)
	})
}

func Test_SimpleStringOnce_GetSetOnce(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_GetSetOnce", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		val1 := sso.GetSetOnce("first")
		val2 := sso.GetSetOnce("second")

		// Act
		actual := args.Map{
			"val1": val1,
			"val2": val2,
		}

		// Assert
		expected := args.Map{
			"val1": "first",
			"val2": "first",
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- GetSetOnce", actual)
	})
}

func Test_SimpleStringOnce_GetOnce(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_GetOnce", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		val := sso.GetOnce()

		// Act
		actual := args.Map{
			"val": val,
			"isInit": sso.IsInitialized(),
		}

		// Assert
		expected := args.Map{
			"val": "",
			"isInit": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- GetOnce", actual)
	})
}

func Test_SimpleStringOnce_GetOnceFunc(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_GetOnceFunc", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		val := sso.GetOnceFunc(func() string { return "computed" })
		val2 := sso.GetOnceFunc(func() string { return "should not" })

		// Act
		actual := args.Map{
			"val": val,
			"val2": val2,
		}

		// Assert
		expected := args.Map{
			"val": "computed",
			"val2": "computed",
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- GetOnceFunc", actual)
	})
}

func Test_SimpleStringOnce_Invalidate(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_Invalidate", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.SetOnUninitialized("hello")
		sso.Invalidate()

		// Act
		actual := args.Map{
			"isInit": sso.IsInitialized(),
			"val": sso.Value(),
		}

		// Assert
		expected := args.Map{
			"isInit": false,
			"val": "",
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns error -- Invalidate", actual)
	})
}

func Test_SimpleStringOnce_Reset(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_Reset", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.SetOnUninitialized("hello")
		sso.Reset()

		// Act
		actual := args.Map{"isInit": sso.IsInitialized()}

		// Assert
		expected := args.Map{"isInit": false}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- Reset", actual)
	})
}

func Test_SimpleStringOnce_IsInvalid(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_IsInvalid", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		actual := args.Map{"isInvalid": sso.IsInvalid()}

		// Assert
		expected := args.Map{"isInvalid": true}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns error -- IsInvalid uninit", actual)
	})
}

func Test_SimpleStringOnce_IntTypes(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_IntTypes", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.SetOnUninitialized("42")

		// Act
		actual := args.Map{
			"int":        sso.Int(),
			"int16":      sso.Int16(),
			"int32":      sso.Int32(),
			"byte":       sso.Byte(),
			"valueInt":   sso.ValueInt(0),
			"defInt":     sso.ValueDefInt(),
			"valueByte":  sso.ValueByte(0),
			"defByte":    sso.ValueDefByte(),
			"valueFloat": sso.ValueFloat64(0),
			"defFloat":   sso.ValueDefFloat64(),
		}

		// Assert
		expected := args.Map{
			"int": 42, "int16": int16(42), "int32": int32(42),
			"byte": byte(42), "valueInt": 42, "defInt": 42,
			"valueByte": byte(42), "defByte": byte(42),
			"valueFloat": 42.0, "defFloat": 42.0,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- int types", actual)
	})
}

func Test_SimpleStringOnce_Boolean(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_Boolean", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.SetOnUninitialized("yes")

		// Act
		actual := args.Map{
			"bool":    sso.Boolean(false),
			"boolDef": sso.BooleanDefault(),
			"isVal":   sso.IsValueBool(),
		}

		// Assert
		expected := args.Map{
			"bool": true,
			"boolDef": true,
			"isVal": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- Boolean", actual)
	})
}

func Test_SimpleStringOnce_Boolean_Uninit(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_Boolean_Uninit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		actual := args.Map{"bool": sso.Boolean(true)}

		// Assert
		expected := args.Map{"bool": false}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- Boolean uninit", actual)
	})
}

func Test_SimpleStringOnce_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_ConcatNew", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.SetOnUninitialized("hello")
		newSso := sso.ConcatNew(" world")

		// Act
		actual := args.Map{"val": newSso.Value()}

		// Assert
		expected := args.Map{"val": "hello world"}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- ConcatNew", actual)
	})
}

func Test_SimpleStringOnce_ConcatNewUsingStrings(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_ConcatNewUsingStrings", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.SetOnUninitialized("a")
		newSso := sso.ConcatNewUsingStrings("-", "b", "c")

		// Act
		actual := args.Map{"val": newSso.Value()}

		// Assert
		expected := args.Map{"val": "a-b-c"}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- ConcatNewUsingStrings", actual)
	})
}

func Test_SimpleStringOnce_NonPtrPtr(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_NonPtrPtr", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.SetOnUninitialized("x")
		nonPtr := sso.NonPtr()
		ptr := sso.Ptr()

		// Act
		actual := args.Map{
			"nonPtrVal": nonPtr.Value(),
			"ptrNotNil": ptr != nil,
		}

		// Assert
		expected := args.Map{
			"nonPtrVal": "x",
			"ptrNotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- NonPtr/Ptr", actual)
	})
}

func Test_SimpleStringOnce_ValueBytes(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_ValueBytes", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.SetOnUninitialized("abc")

		// Act
		actual := args.Map{
			"len": len(sso.ValueBytes()),
			"lenPtr": len(sso.ValueBytesPtr()),
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"lenPtr": 3,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- ValueBytes", actual)
	})
}

func Test_SimpleStringOnce_SetOnceIfUninitialized(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_SetOnceIfUninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		isSet1 := sso.SetOnceIfUninitialized("first")
		isSet2 := sso.SetOnceIfUninitialized("second")

		// Act
		actual := args.Map{
			"isSet1": isSet1,
			"isSet2": isSet2,
			"val": sso.Value(),
		}

		// Assert
		expected := args.Map{
			"isSet1": true,
			"isSet2": false,
			"val": "first",
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- SetOnceIfUninitialized", actual)
	})
}

func Test_SimpleStringOnce_SetInitSetUnInit(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_SetInitSetUnInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		sso.SetInitialize()
		isInit := sso.IsInitialized()
		sso.SetUnInit()
		isUninit := sso.IsUninitialized()

		// Act
		actual := args.Map{
			"isInit": isInit,
			"isUninit": isUninit,
		}

		// Assert
		expected := args.Map{
			"isInit": true,
			"isUninit": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- SetInit/SetUnInit", actual)
	})
}

func Test_SimpleStringOnce_IsAnyOf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_IsAnyOf_Empty", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.SetOnUninitialized("x")

		// Act
		actual := args.Map{"result": sso.IsAnyOf()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns empty -- IsAnyOf empty", actual)
	})
}

func Test_SimpleStringOnce_IsAnyContains_Empty(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_IsAnyContains_Empty", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.SetOnUninitialized("x")

		// Act
		actual := args.Map{"result": sso.IsAnyContains()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns empty -- IsAnyContains empty", actual)
	})
}

func Test_SimpleStringOnce_RegexNil(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_RegexNil", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.SetOnUninitialized("test")

		// Act
		actual := args.Map{
			"matches": sso.IsRegexMatches(nil),
			"findStr": sso.RegexFindString(nil),
		}

		// Assert
		expected := args.Map{
			"matches": false,
			"findStr": "",
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns nil -- regex nil", actual)
	})
}

func Test_SimpleStringOnce_Uint16(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_Uint16", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.SetOnUninitialized("1000")
		val, inRange := sso.Uint16()

		// Act
		actual := args.Map{
			"val": val,
			"inRange": inRange,
		}

		// Assert
		expected := args.Map{
			"val": uint16(1000),
			"inRange": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- Uint16", actual)
	})
}

func Test_SimpleStringOnce_Uint32(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_Uint32", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.SetOnUninitialized("50000")
		val, inRange := sso.Uint32()

		// Act
		actual := args.Map{
			"val": val,
			"inRange": inRange,
		}

		// Assert
		expected := args.Map{
			"val": uint32(50000),
			"inRange": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- Uint32", actual)
	})
}

func Test_SimpleStringOnce_WithinRange(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_WithinRange", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.SetOnUninitialized("50")
		val, inRange := sso.WithinRange(true, 0, 100)
		valOut, outRange := sso.WithinRange(true, 60, 100)
		_, noBoundRange := sso.WithinRange(false, 60, 100)

		// Act
		actual := args.Map{
			"val": val, "inRange": inRange,
			"valOut": valOut, "outRange": outRange,
			"noBoundRange": noBoundRange,
		}

		// Assert
		expected := args.Map{
			"val": 50, "inRange": true,
			"valOut": 60, "outRange": false,
			"noBoundRange": false,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns non-empty -- WithinRange", actual)
	})
}

func Test_SimpleStringOnce_SafeValue_Uninit(t *testing.T) {
	safeTest(t, "Test_Cov9_SimpleStringOnce_SafeValue_Uninit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		actual := args.Map{"val": sso.SafeValue()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- SafeValue uninit", actual)
	})
}

// ═══════════════════════════════════════════
// KeyAnyValuePair
// ═══════════════════════════════════════════

func Test_KeyAnyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov9_KeyAnyValuePair", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "name", Value: "hello"}

		// Act
		actual := args.Map{
			"keyName":    kav.KeyName(),
			"varName":    kav.VariableName(),
			"valueAny":   kav.ValueAny(),
			"isVarEq":    kav.IsVariableNameEqual("name"),
			"isNull":     kav.IsValueNull(),
			"hasNonNull": kav.HasNonNull(),
			"hasValue":   kav.HasValue(),
			"isEmptyStr": kav.IsValueEmptyString(),
			"isWS":       kav.IsValueWhitespace(),
			"valStr":     kav.ValueString(),
		}

		// Assert
		expected := args.Map{
			"keyName": "name", "varName": "name",
			"valueAny": "hello", "isVarEq": true,
			"isNull": false, "hasNonNull": true,
			"hasValue": true, "isEmptyStr": false,
			"isWS": false, "valStr": "hello",
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- with args", actual)
	})
}

func Test_KeyAnyValuePair_NilValue(t *testing.T) {
	safeTest(t, "Test_Cov9_KeyAnyValuePair_NilValue", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: nil}

		// Act
		actual := args.Map{
			"isNull":     kav.IsValueNull(),
			"hasNonNull": kav.HasNonNull(),
		}

		// Assert
		expected := args.Map{
			"isNull": true,
			"hasNonNull": false,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns nil -- nil value", actual)
	})
}

func Test_KeyAnyValuePair_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov9_KeyAnyValuePair_ClearDispose", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kav.Clear()

		// Act
		actual := args.Map{
			"key": kav.Key,
			"valNil": kav.Value == nil,
		}

		// Assert
		expected := args.Map{
			"key": "",
			"valNil": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- Clear", actual)
	})
}

func Test_KeyAnyValuePair_NilReceiver(t *testing.T) {
	safeTest(t, "Test_Cov9_KeyAnyValuePair_NilReceiver", func() {
		// Arrange
		var kav *corestr.KeyAnyValuePair
		kav.Clear()
		kav.Dispose()

		// Act
		actual := args.Map{
			"isNull":     kav.IsValueNull(),
			"isEmptyStr": kav.IsValueEmptyString(),
			"isWS":       kav.IsValueWhitespace(),
		}

		// Assert
		expected := args.Map{
			"isNull": true,
			"isEmptyStr": true,
			"isWS": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns nil -- nil receiver", actual)
	})
}

func Test_KeyAnyValuePair_Compile(t *testing.T) {
	safeTest(t, "Test_Cov9_KeyAnyValuePair_Compile", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		compile := kav.Compile()
		str := kav.String()

		// Act
		actual := args.Map{"same": compile == str}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- Compile", actual)
	})
}

// ═══════════════════════════════════════════
// ValidValues
// ═══════════════════════════════════════════

func Test_ValidValues_Basic(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValues_Basic", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")

		// Act
		actual := args.Map{
			"len":      vvs.Length(),
			"count":    vvs.Count(),
			"isEmpty":  vvs.IsEmpty(),
			"hasAny":   vvs.HasAnyItem(),
			"lastIdx":  vvs.LastIndex(),
			"hasIdx0":  vvs.HasIndex(0),
			"safeVal0": vvs.SafeValueAt(0),
			"safeVal5": vvs.SafeValueAt(5),
		}

		// Assert
		expected := args.Map{
			"len": 2, "count": 2, "isEmpty": false,
			"hasAny": true, "lastIdx": 1,
			"hasIdx0": true, "safeVal0": "a", "safeVal5": "",
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- basic", actual)
	})
}

func Test_ValidValues_Strings(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValues_Strings", func() {
		// Arrange
		vvs := corestr.NewValidValues(3)
		vvs.Add("a").Add("b")
		strs := vvs.Strings()
		fullStrs := vvs.FullStrings()
		str := vvs.String()

		// Act
		actual := args.Map{
			"strsLen":     len(strs),
			"fullStrsLen": len(fullStrs),
			"strNotEmpty": str != "",
		}

		// Assert
		expected := args.Map{
			"strsLen": 2,
			"fullStrsLen": 2,
			"strNotEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- Strings", actual)
	})
}

func Test_ValidValues_EmptyConstructors(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValues_EmptyConstructors", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		actual := args.Map{
			"len": vvs.Length(),
			"isEmpty": vvs.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- empty", actual)
	})
}

func Test_ValidValues_UsingValues(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValues_UsingValues", func() {
		// Arrange
		vvs := corestr.NewValidValuesUsingValues(
			corestr.ValidValue{Value: "a", IsValid: true},
			corestr.ValidValue{Value: "b", IsValid: true},
		)

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- UsingValues", actual)
	})
}

func Test_ValidValues_UsingValues_Empty(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValues_UsingValues_Empty", func() {
		// Arrange
		vvs := corestr.NewValidValuesUsingValues()

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- UsingValues empty", actual)
	})
}

func Test_ValidValues_AddFull(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValues_AddFull", func() {
		// Arrange
		vvs := corestr.NewValidValues(3)
		vvs.AddFull(true, "val", "msg")

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- AddFull", actual)
	})
}

func Test_ValidValues_Find(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValues_Find", func() {
		// Arrange
		vvs := corestr.NewValidValues(3)
		vvs.Add("a").Add("b").Add("c")
		found := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, vv.Value == "b", false
		})

		// Act
		actual := args.Map{"len": len(found)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- Find", actual)
	})
}

func Test_ValidValues_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValues_SafeValuesAtIndexes", func() {
		// Arrange
		vvs := corestr.NewValidValues(3)
		vvs.Add("a").Add("b")
		vals := vvs.SafeValuesAtIndexes(0, 1, 5)

		// Act
		actual := args.Map{
			"first": vals[0],
			"second": vals[1],
			"third": vals[2],
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"second": "b",
			"third": "",
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- SafeValuesAtIndexes", actual)
	})
}

func Test_ValidValues_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValues_ConcatNew", func() {
		// Arrange
		vvs1 := corestr.NewValidValues(3)
		vvs1.Add("a")
		vvs2 := corestr.NewValidValues(3)
		vvs2.Add("b")
		newVvs := vvs1.ConcatNew(true, vvs2)

		// Act
		actual := args.Map{"len": newVvs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- ConcatNew", actual)
	})
}

func Test_ValidValues_ConcatNew_Empty_Clone(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValues_ConcatNew_Empty_Clone", func() {
		// Arrange
		vvs := corestr.NewValidValues(3)
		vvs.Add("a")
		newVvs := vvs.ConcatNew(true)

		// Act
		actual := args.Map{"len": newVvs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- ConcatNew empty clone", actual)
	})
}

func Test_ValidValues_ConcatNew_Empty_NoClone(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValues_ConcatNew_Empty_NoClone", func() {
		// Arrange
		vvs := corestr.NewValidValues(3)
		vvs.Add("a")
		sameVvs := vvs.ConcatNew(false)

		// Act
		actual := args.Map{"len": sameVvs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- ConcatNew empty no clone", actual)
	})
}

func Test_ValidValues_Hashmap(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValues_Hashmap", func() {
		// Arrange
		vvs := corestr.NewValidValues(3)
		vvs.Add("k1").Add("k2")
		hm := vvs.Hashmap()
		m := vvs.Map()

		// Act
		actual := args.Map{
			"hmLen": hm.Length(),
			"mapLen": len(m),
		}

		// Assert
		expected := args.Map{
			"hmLen": 2,
			"mapLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- Hashmap", actual)
	})
}

func Test_ValidValues_NilReceiver(t *testing.T) {
	safeTest(t, "Test_Cov9_ValidValues_NilReceiver", func() {
		// Arrange
		var vvs *corestr.ValidValues

		// Act
		actual := args.Map{
			"len": vvs.Length(),
			"isEmpty": vvs.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns nil -- nil receiver", actual)
	})
}

// ═══════════════════════════════════════════
// HashsetsCollection
// ═══════════════════════════════════════════

func Test_HashsetsCollection_Basic(t *testing.T) {
	safeTest(t, "Test_Cov9_HashsetsCollection_Basic", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 5)
		h1 := corestr.New.Hashset.Cap(3)
		h1.Adds("a", "b")
		hsc.Add(h1)

		// Act
		actual := args.Map{
			"len":      hsc.Length(),
			"isEmpty":  hsc.IsEmpty(),
			"hasItems": hsc.HasItems(),
			"lastIdx":  hsc.LastIndex(),
		}

		// Assert
		expected := args.Map{
			"len": 1, "isEmpty": false, "hasItems": true, "lastIdx": 0,
		}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection returns correct value -- basic", actual)
	})
}

func Test_HashsetsCollection_AddNonNilNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov9_HashsetsCollection_AddNonNilNonEmpty", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 5)
		h := corestr.New.Hashset.Cap(3)
		h.Add("a")
		hsc.AddNonNil(h)
		hsc.AddNonNil(nil)
		hsc.AddNonEmpty(h)

		// Act
		actual := args.Map{"len": hsc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection returns nil -- AddNonNil/NonEmpty", actual)
	})
}

func Test_HashsetsCollection_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov9_HashsetsCollection_IsEqual", func() {
		// Arrange
		h1 := corestr.New.Hashset.Cap(3)
		h1.Add("a")
		hsc1 := corestr.New.HashsetsCollection.LenCap(0, 3)
		hsc1.Add(h1)
		h2 := corestr.New.Hashset.Cap(3)
		h2.Add("a")
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 3)
		hsc2.Add(h2)

		// Act
		actual := args.Map{
			"equal":    hsc1.IsEqual(*hsc2),
			"equalPtr": hsc1.IsEqualPtr(hsc2),
		}

		// Assert
		expected := args.Map{
			"equal": true,
			"equalPtr": true,
		}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection returns correct value -- IsEqual", actual)
	})
}

func Test_HashsetsCollection_StringsList(t *testing.T) {
	safeTest(t, "Test_Cov9_HashsetsCollection_StringsList", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(3)
		h.Add("a")
		hsc := corestr.New.HashsetsCollection.LenCap(0, 3)
		hsc.Add(h)
		list := hsc.StringsList()

		// Act
		actual := args.Map{"len": len(list)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection returns correct value -- StringsList", actual)
	})
}

func Test_HashsetsCollection_NilReceiver(t *testing.T) {
	safeTest(t, "Test_Cov9_HashsetsCollection_NilReceiver", func() {
		// Arrange
		var hsc *corestr.HashsetsCollection

		// Act
		actual := args.Map{
			"len": hsc.Length(),
			"isEmpty": hsc.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection returns nil -- nil receiver", actual)
	})
}

func Test_HashsetsCollection_String(t *testing.T) {
	safeTest(t, "Test_Cov9_HashsetsCollection_String", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 3)
		emptyStr := hsc.String()
		h := corestr.New.Hashset.Cap(3)
		h.Add("a")
		hsc.Add(h)
		nonEmptyStr := hsc.String()

		// Act
		actual := args.Map{
			"emptyNotEmpty":    emptyStr != "",
			"nonEmptyNotEmpty": nonEmptyStr != "",
		}

		// Assert
		expected := args.Map{
			"emptyNotEmpty": true,
			"nonEmptyNotEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection returns correct value -- String", actual)
	})
}

func Test_HashsetsCollection_Join(t *testing.T) {
	safeTest(t, "Test_Cov9_HashsetsCollection_Join", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(3)
		h.Add("a")
		hsc := corestr.New.HashsetsCollection.LenCap(0, 3)
		hsc.Add(h)
		joined := hsc.Join(",")

		// Act
		actual := args.Map{"notEmpty": joined != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection returns correct value -- Join", actual)
	})
}
