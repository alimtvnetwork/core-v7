package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ── Collection ──

func Test_Collection_BasicMethods(t *testing.T) {
	safeTest(t, "Test_Collection_BasicMethods", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{
			"length":     col.Length(),
			"count":      col.Count(),
			"lastIndex":  col.LastIndex(),
			"hasAny":     col.HasAnyItem(),
			"isEmpty":    col.IsEmpty(),
			"hasItems":   col.HasItems(),
			"hasIndex1":  col.HasIndex(1),
			"hasIndex10": col.HasIndex(10),
			"capacity":   col.Capacity() >= 3,
		}
		expected := args.Map{
			"length":     3,
			"count":      3,
			"lastIndex":  2,
			"hasAny":     true,
			"isEmpty":    false,
			"hasItems":   true,
			"hasIndex1":  true,
			"hasIndex10": false,
			"capacity":   true,
		}
		expected.ShouldBeEqual(t, 0, "Collection_BasicMethods returns correct value -- with args", actual)
	})
}

func Test_Collection_NilSafe(t *testing.T) {
	safeTest(t, "Test_Collection_NilSafe", func() {
		// Arrange
		var col *corestr.Collection

		// Act
		actual := args.Map{
			"length":  col.Length(),
			"isEmpty": col.IsEmpty(),
		}
		expected := args.Map{
			"length":  0,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection_NilSafe returns nil -- with args", actual)
	})
}

func Test_Collection_AddMethods(t *testing.T) {
	safeTest(t, "Test_Collection_AddMethods", func() {
		// Arrange
		col := corestr.New.Collection.Empty()

		// Act
		col.Add("a").
			AddNonEmpty("b").
			AddNonEmpty("").
			AddNonEmptyWhitespace("c").
			AddNonEmptyWhitespace("  ").
			AddIf(true, "d").
			AddIf(false, "e").
			Adds("f", "g")

		actual := args.Map{
			"length": col.Length(),
		}
		expected := args.Map{
			"length": 6,
		}
		expected.ShouldBeEqual(t, 0, "Collection_AddMethods returns correct value -- with args", actual)
	})
}

func Test_Collection_RemoveAt_FromCollectionBasicMetho(t *testing.T) {
	safeTest(t, "Test_Collection_RemoveAt", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		ok1 := col.RemoveAt(1)
		ok2 := col.RemoveAt(10)
		ok3 := col.RemoveAt(-1)

		actual := args.Map{
			"ok1":    ok1,
			"ok2":    ok2,
			"ok3":    ok3,
			"length": col.Length(),
		}
		expected := args.Map{
			"ok1":    true,
			"ok2":    false,
			"ok3":    false,
			"length": 2,
		}
		expected.ShouldBeEqual(t, 0, "Collection_RemoveAt returns correct value -- with args", actual)
	})
}

func Test_Collection_IsEquals_FromCollectionBasicMetho(t *testing.T) {
	safeTest(t, "Test_Collection_IsEquals", func() {
		// Arrange
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"a", "b"})
		c3 := corestr.New.Collection.Strings([]string{"a", "c"})
		c4 := corestr.New.Collection.Strings([]string{"A", "B"})

		// Act
		actual := args.Map{
			"sameVals":      c1.IsEquals(c2),
			"diffVals":      c1.IsEquals(c3),
			"insensitive":   c1.IsEqualsWithSensitive(false, c4),
			"sensitiveCase": c1.IsEqualsWithSensitive(true, c4),
		}
		expected := args.Map{
			"sameVals":      true,
			"diffVals":      false,
			"insensitive":   true,
			"sensitiveCase": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection_IsEquals returns correct value -- with args", actual)
	})
}

func Test_Collection_JoinAndString(t *testing.T) {
	safeTest(t, "Test_Collection_JoinAndString", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		joined := col.Join(",")
		strNotEmpty := col.String() != ""

		actual := args.Map{
			"joined":      joined,
			"strNotEmpty": strNotEmpty,
		}
		expected := args.Map{
			"joined":      "a,b,c",
			"strNotEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection_JoinAndString returns correct value -- with args", actual)
	})
}

func Test_Collection_AddError_FromCollectionBasicMetho(t *testing.T) {
	safeTest(t, "Test_Collection_AddError", func() {
		// Arrange
		col := corestr.New.Collection.Empty()

		// Act
		col.AddError(nil)
		col.AddError(errForTest)
		asErr := col.AsDefaultError()

		actual := args.Map{
			"length":   col.Length(),
			"hasError": asErr != nil,
		}
		expected := args.Map{
			"length":   1,
			"hasError": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection_AddError returns error -- with args", actual)
	})
}

func Test_Collection_ConcatNew_FromCollectionBasicMetho(t *testing.T) {
	safeTest(t, "Test_Collection_ConcatNew", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		newCol := col.ConcatNew(0, "c", "d")
		noAdd := col.ConcatNew(0)

		actual := args.Map{
			"newLen":   newCol.Length(),
			"noAddLen": noAdd.Length(),
		}
		expected := args.Map{
			"newLen":   4,
			"noAddLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "Collection_ConcatNew returns correct value -- with args", actual)
	})
}

func Test_Collection_EachItemSplitBy_FromCollectionBasicMetho(t *testing.T) {
	safeTest(t, "Test_Collection_EachItemSplitBy", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"a:b", "c:d"})

		// Act
		split := col.EachItemSplitBy(":")

		actual := args.Map{
			"splitLen": len(split),
		}
		expected := args.Map{
			"splitLen": 4,
		}
		expected.ShouldBeEqual(t, 0, "Collection_EachItemSplitBy returns correct value -- with args", actual)
	})
}

func Test_Collection_AddCollections_FromCollectionBasicMetho(t *testing.T) {
	safeTest(t, "Test_Collection_AddCollections", func() {
		// Arrange
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c3 := corestr.New.Collection.Strings([]string{"c"})

		// Act
		c1.AddCollections(c2, c3, nil)

		actual := args.Map{
			"length": c1.Length(),
		}
		expected := args.Map{
			"length": 3,
		}
		expected.ShouldBeEqual(t, 0, "Collection_AddCollections returns correct value -- with args", actual)
	})
}

func Test_Collection_AsError_FromCollectionBasicMetho(t *testing.T) {
	safeTest(t, "Test_Collection_AsError", func() {
		// Arrange
		empty := corestr.New.Collection.Empty()
		withItems := corestr.New.Collection.Strings([]string{"err1", "err2"})

		// Act
		actual := args.Map{
			"emptyErr": empty.AsError(",") == nil,
			"hasErr":   withItems.AsError(",") != nil,
		}
		expected := args.Map{
			"emptyErr": true,
			"hasErr":   true,
		}
		expected.ShouldBeEqual(t, 0, "Collection_AsError returns error -- with args", actual)
	})
}

// ── Hashset ──

func Test_Hashset_BasicMethods(t *testing.T) {
	safeTest(t, "Test_Hashset_BasicMethods", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{
			"isEmpty":    hs.IsEmpty(),
			"hasItems":   hs.HasItems(),
			"length":     hs.Length(),
			"hasA":       hs.Has("a"),
			"hasD":       hs.Has("d"),
			"containsB":  hs.Contains("b"),
			"isMissingD": hs.IsMissing("d"),
		}
		expected := args.Map{
			"isEmpty":    false,
			"hasItems":   true,
			"length":     3,
			"hasA":       true,
			"hasD":       false,
			"containsB":  true,
			"isMissingD": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset_BasicMethods returns correct value -- with args", actual)
	})
}

func Test_Hashset_AddMethods(t *testing.T) {
	safeTest(t, "Test_Hashset_AddMethods", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		hs.Add("a").
			AddNonEmpty("b").
			AddNonEmpty("").
			AddNonEmptyWhitespace("c").
			AddNonEmptyWhitespace("  ").
			AddIf(true, "d").
			AddIf(false, "e").
			Adds("f", "g")

		actual := args.Map{
			"length": hs.Length(),
		}
		expected := args.Map{
			"length": 6,
		}
		expected.ShouldBeEqual(t, 0, "Hashset_AddMethods returns correct value -- with args", actual)
	})
}

func Test_Hashset_AddBool(t *testing.T) {
	safeTest(t, "Test_Hashset_AddBool", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		firstAdd := hs.AddBool("a")
		secondAdd := hs.AddBool("a")

		actual := args.Map{
			"firstAdd":  firstAdd,
			"secondAdd": secondAdd,
		}
		expected := args.Map{
			"firstAdd":  false,
			"secondAdd": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset_AddBool returns correct value -- with args", actual)
	})
}

func Test_Hashset_NilSafe(t *testing.T) {
	safeTest(t, "Test_Hashset_NilSafe", func() {
		// Arrange
		var hs *corestr.Hashset

		// Act
		actual := args.Map{
			"isEmpty": hs.IsEmpty(),
		}
		expected := args.Map{
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset_NilSafe returns nil -- with args", actual)
	})
}

func Test_Hashset_ListAndSortedList(t *testing.T) {
	safeTest(t, "Test_Hashset_ListAndSortedList", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})

		// Act
		list := hs.List()
		sorted := hs.SortedList()

		actual := args.Map{
			"listLen":   len(list),
			"sortedLen": len(sorted),
			"sortedFirst": sorted[0],
		}
		expected := args.Map{
			"listLen":   3,
			"sortedLen": 3,
			"sortedFirst": "a",
		}
		expected.ShouldBeEqual(t, 0, "Hashset_ListAndSortedList returns correct value -- with args", actual)
	})
}

// ── Hashmap ──

func Test_Hashmap_BasicMethods(t *testing.T) {
	safeTest(t, "Test_Hashmap_BasicMethods", func() {
		// Arrange
		hm := corestr.New.Hashmap.KeyValues(
			corestr.KeyValuePair{Key: "k1", Value: "v1"},
			corestr.KeyValuePair{Key: "k2", Value: "v2"},
		)

		// Act
		actual := args.Map{
			"isEmpty":     hm.IsEmpty(),
			"hasItems":    hm.HasItems(),
			"length":      hm.Length(),
			"hasK1":       hm.Has("k1"),
			"hasK3":       hm.Has("k3"),
			"containsK2":  hm.Contains("k2"),
			"isMissingK3": hm.IsKeyMissing("k3"),
		}
		expected := args.Map{
			"isEmpty":     false,
			"hasItems":    true,
			"length":      2,
			"hasK1":       true,
			"hasK3":       false,
			"containsK2":  true,
			"isMissingK3": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap_BasicMethods returns correct value -- with args", actual)
	})
}

func Test_Hashmap_GetValue(t *testing.T) {
	safeTest(t, "Test_Hashmap_GetValue", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("key1", "val1")

		// Act
		val, found := hm.GetValue("key1")
		_, notFound := hm.GetValue("missing")

		actual := args.Map{
			"val":      val,
			"found":    found,
			"notFound": notFound,
		}
		expected := args.Map{
			"val":      "val1",
			"found":    true,
			"notFound": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap_GetValue returns correct value -- with args", actual)
	})
}

func Test_Hashmap_SetBySplitter_FromCollectionBasicMetho(t *testing.T) {
	safeTest(t, "Test_Hashmap_SetBySplitter", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()

		// Act
		isNew1 := hm.SetBySplitter("=", "key=value")
		isNew2 := hm.SetBySplitter("=", "noequals")

		actual := args.Map{
			"isNew1": isNew1,
			"isNew2": isNew2,
			"length": hm.Length(),
		}
		expected := args.Map{
			"isNew1": true,
			"isNew2": true,
			"length": 2,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap_SetBySplitter returns correct value -- with args", actual)
	})
}

func Test_Hashmap_NilSafe(t *testing.T) {
	safeTest(t, "Test_Hashmap_NilSafe", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act
		actual := args.Map{
			"isEmpty": hm.IsEmpty(),
		}
		expected := args.Map{
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap_NilSafe returns nil -- with args", actual)
	})
}

// ── SimpleSlice ──

func Test_SimpleSlice_BasicMethods(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_BasicMethods", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		actual := args.Map{
			"length":    ss.Length(),
			"count":     ss.Count(),
			"lastIndex": ss.LastIndex(),
			"hasAny":    ss.HasAnyItem(),
			"isEmpty":   ss.IsEmpty(),
			"first":     ss.First(),
			"last":      ss.Last(),
			"hasIndex1": ss.HasIndex(1),
		}
		expected := args.Map{
			"length":    3,
			"count":     3,
			"lastIndex": 2,
			"hasAny":    true,
			"isEmpty":   false,
			"first":     "a",
			"last":      "c",
			"hasIndex1": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice_BasicMethods returns correct value -- with args", actual)
	})
}

func Test_SimpleSlice_AddMethods(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_AddMethods", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.Add("a").
			AddIf(true, "b").
			AddIf(false, "c").
			Adds("d", "e").
			AddsIf(true, "f").
			AddsIf(false, "g")

		actual := args.Map{
			"length": ss.Length(),
		}
		expected := args.Map{
			"length": 5,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice_AddMethods returns correct value -- with args", actual)
	})
}

func Test_SimpleSlice_SkipTake_FromCollectionBasicMetho(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_SkipTake", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c", "d", "e")

		// Act
		skipped := ss.Skip(2)
		taken := ss.Take(3)
		limited := ss.Limit(2)

		actual := args.Map{
			"skippedLen": len(skipped),
			"takenLen":   len(taken),
			"limitedLen": len(limited),
		}
		expected := args.Map{
			"skippedLen": 3,
			"takenLen":   3,
			"limitedLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice_SkipTake returns correct value -- with args", actual)
	})
}

func Test_SimpleSlice_FirstLastOrDefault(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_FirstLastOrDefault", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		empty := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{
			"firstOrDef":      ss.FirstOrDefault(),
			"lastOrDef":       ss.LastOrDefault(),
			"emptyFirstOrDef": empty.FirstOrDefault(),
			"emptyLastOrDef":  empty.LastOrDefault(),
		}
		expected := args.Map{
			"firstOrDef":      "a",
			"lastOrDef":       "b",
			"emptyFirstOrDef": "",
			"emptyLastOrDef":  "",
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice_FirstLastOrDefault returns correct value -- with args", actual)
	})
}

func Test_SimpleSlice_IsContains_FromCollectionBasicMetho(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_IsContains", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("hello", "world")

		// Act
		actual := args.Map{
			"containsHello":   ss.IsContains("hello"),
			"containsMissing": ss.IsContains("missing"),
			"indexOf":         ss.IndexOf("world"),
			"indexOfMissing":  ss.IndexOf("missing"),
		}
		expected := args.Map{
			"containsHello":   true,
			"containsMissing": false,
			"indexOf":         1,
			"indexOfMissing":  -1,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice_IsContains returns correct value -- with args", actual)
	})
}

func Test_SimpleSlice_InsertAt_FromCollectionBasicMetho(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_InsertAt", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "c")

		// Act
		ss.InsertAt(1, "b")

		actual := args.Map{
			"length": ss.Length(),
			"item1":  ss.Strings()[1],
		}
		expected := args.Map{
			"length": 3,
			"item1":  "b",
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice_InsertAt returns correct value -- with args", actual)
	})
}

func Test_SimpleSlice_NilSafe(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_NilSafe", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		actual := args.Map{
			"length":  ss.Length(),
			"isEmpty": ss.IsEmpty(),
		}
		expected := args.Map{
			"length":  0,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice_NilSafe returns nil -- with args", actual)
	})
}

// ── LeftRight ──

func Test_LeftRight_BasicMethods(t *testing.T) {
	safeTest(t, "Test_LeftRight_BasicMethods", func() {
		// Arrange
		lr := corestr.NewLeftRight("hello", "world")

		// Act
		actual := args.Map{
			"isValid":        lr.IsValid,
			"left":           lr.Left,
			"right":          lr.Right,
			"isLeftEmpty":    lr.IsLeftEmpty(),
			"isRightEmpty":   lr.IsRightEmpty(),
			"hasSafeNonEmpty": lr.HasSafeNonEmpty(),
			"leftTrim":       lr.LeftTrim(),
			"rightTrim":      lr.RightTrim(),
		}
		expected := args.Map{
			"isValid":        true,
			"left":           "hello",
			"right":          "world",
			"isLeftEmpty":    false,
			"isRightEmpty":   false,
			"hasSafeNonEmpty": true,
			"leftTrim":       "hello",
			"rightTrim":      "world",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight_BasicMethods returns correct value -- with args", actual)
	})
}

func Test_LeftRight_UsingSlice_FromCollectionBasicMetho(t *testing.T) {
	safeTest(t, "Test_LeftRight_UsingSlice", func() {
		// Arrange & Act
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})
		lr1 := corestr.LeftRightUsingSlice([]string{"a"})
		lr0 := corestr.LeftRightUsingSlice([]string{})
		lrTrim := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})

		actual := args.Map{
			"lr_valid":     lr.IsValid,
			"lr_left":      lr.Left,
			"lr_right":     lr.Right,
			"lr1_valid":    lr1.IsValid,
			"lr1_right":    lr1.Right,
			"lr0_valid":    lr0.IsValid,
			"lrTrim_left":  lrTrim.Left,
			"lrTrim_right": lrTrim.Right,
		}
		expected := args.Map{
			"lr_valid":     true,
			"lr_left":      "a",
			"lr_right":     "b",
			"lr1_valid":    false,
			"lr1_right":    "",
			"lr0_valid":    false,
			"lrTrim_left":  "a",
			"lrTrim_right": "b",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight_UsingSlice returns correct value -- with args", actual)
	})
}

func Test_LeftRight_IsEqual_FromCollectionBasicMetho(t *testing.T) {
	safeTest(t, "Test_LeftRight_IsEqual", func() {
		// Arrange
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		lr3 := corestr.NewLeftRight("a", "c")

		// Act
		actual := args.Map{
			"same":    lr1.IsEqual(lr2),
			"diff":    lr1.IsEqual(lr3),
			"nilBoth": (*corestr.LeftRight)(nil).IsEqual(nil),
		}
		expected := args.Map{
			"same":    true,
			"diff":    false,
			"nilBoth": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight_IsEqual returns correct value -- with args", actual)
	})
}

// ── ValidValue ──

func Test_ValidValue_BasicMethods(t *testing.T) {
	safeTest(t, "Test_ValidValue_BasicMethods", func() {
		// Arrange
		vv := corestr.NewValidValue("42")

		// Act
		actual := args.Map{
			"value":       vv.Value,
			"isValid":     vv.IsValid,
			"isEmpty":     vv.IsEmpty(),
			"valueBool":   vv.ValueBool(),
			"valueInt":    vv.ValueInt(0),
			"valueDefInt": vv.ValueDefInt(),
			"valueByte":   int(vv.ValueDefByte()),
			"trim":        vv.Trim(),
		}
		expected := args.Map{
			"value":       "42",
			"isValid":     true,
			"isEmpty":     false,
			"valueBool":   false,
			"valueInt":    42,
			"valueDefInt": 42,
			"valueByte":   42,
			"trim":        "42",
		}
		expected.ShouldBeEqual(t, 0, "ValidValue_BasicMethods returns non-empty -- with args", actual)
	})
}

func Test_ValidValue_InvalidAndEmpty(t *testing.T) {
	safeTest(t, "Test_ValidValue_InvalidAndEmpty", func() {
		// Arrange
		invalid := corestr.InvalidValidValue("err msg")
		empty := corestr.NewValidValueEmpty()
		noMsg := corestr.InvalidValidValueNoMessage()

		// Act
		actual := args.Map{
			"invalidIsValid": invalid.IsValid,
			"invalidMsg":     invalid.Message,
			"emptyIsValid":   empty.IsValid,
			"emptyValue":     empty.Value,
			"noMsgIsValid":   noMsg.IsValid,
		}
		expected := args.Map{
			"invalidIsValid": false,
			"invalidMsg":     "err msg",
			"emptyIsValid":   true,
			"emptyValue":     "",
			"noMsgIsValid":   false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue_InvalidAndEmpty returns empty -- with args", actual)
	})
}

func Test_ValidValue_Clone_FromCollectionBasicMetho(t *testing.T) {
	safeTest(t, "Test_ValidValue_Clone", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		var nilVV *corestr.ValidValue

		// Act
		cloned := vv.Clone()
		nilClone := nilVV.Clone()

		actual := args.Map{
			"clonedVal": cloned.Value,
			"nilClone":  nilClone == nil,
		}
		expected := args.Map{
			"clonedVal": "hello",
			"nilClone":  true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue_Clone returns non-empty -- with args", actual)
	})
}

func Test_ValidValue_IsAnyOf_FromCollectionBasicMetho(t *testing.T) {
	safeTest(t, "Test_ValidValue_IsAnyOf", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{
			"found":    vv.IsAnyOf("hello", "world"),
			"notFound": vv.IsAnyOf("foo", "bar"),
			"empty":    vv.IsAnyOf(),
		}
		expected := args.Map{
			"found":    true,
			"notFound": false,
			"empty":    true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue_IsAnyOf returns non-empty -- with args", actual)
	})
}

// ── KeyValuePair ──

func Test_KeyValuePair_BasicMethods(t *testing.T) {
	safeTest(t, "Test_KeyValuePair_BasicMethods", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "name", Value: "42"}

		// Act
		actual := args.Map{
			"keyName":      kv.KeyName(),
			"valueString":  kv.ValueString(),
			"isKeyEmpty":   kv.IsKeyEmpty(),
			"isValueEmpty": kv.IsValueEmpty(),
			"hasKey":       kv.HasKey(),
			"hasValue":     kv.HasValue(),
			"valueBool":    kv.ValueBool(),
			"valueInt":     kv.ValueInt(0),
			"valueDefInt":  kv.ValueDefInt(),
			"valueByte":    int(kv.ValueDefByte()),
			"isKey":        kv.IsKey("name"),
			"isVal":        kv.IsVal("42"),
			"is":           kv.Is("name", "42"),
			"strNotEmpty":  kv.String() != "",
		}
		expected := args.Map{
			"keyName":      "name",
			"valueString":  "42",
			"isKeyEmpty":   false,
			"isValueEmpty": false,
			"hasKey":       true,
			"hasValue":     true,
			"valueBool":    false,
			"valueInt":     42,
			"valueDefInt":  42,
			"valueByte":    42,
			"isKey":        true,
			"isVal":        true,
			"is":           true,
			"strNotEmpty":  true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair_BasicMethods returns correct value -- with args", actual)
	})
}

// ── LeftMiddleRight ──

func Test_LeftMiddleRight_BasicMethods(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRight_BasicMethods", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{
			"isValid":        lmr.IsValid,
			"left":           lmr.Left,
			"middle":         lmr.Middle,
			"right":          lmr.Right,
			"hasSafeNonEmpty": lmr.HasSafeNonEmpty(),
			"isLeftEmpty":    lmr.IsLeftEmpty(),
			"isMiddleEmpty":  lmr.IsMiddleEmpty(),
			"isRightEmpty":   lmr.IsRightEmpty(),
			"isAll":          lmr.IsAll("a", "b", "c"),
		}
		expected := args.Map{
			"isValid":        true,
			"left":           "a",
			"middle":         "b",
			"right":          "c",
			"hasSafeNonEmpty": true,
			"isLeftEmpty":    false,
			"isMiddleEmpty":  false,
			"isRightEmpty":   false,
			"isAll":          true,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight_BasicMethods returns correct value -- with args", actual)
	})
}

// ── TextWithLineNumber ──

func Test_TextWithLineNumber_BasicMethods(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_BasicMethods", func() {
		// Arrange
		twl := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}
		var nilTwl *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"hasLineNumber":      twl.HasLineNumber(),
			"isInvalidLineNumber": twl.IsInvalidLineNumber(),
			"length":             twl.Length(),
			"isEmpty":            twl.IsEmpty(),
			"isEmptyText":        twl.IsEmptyText(),
			"nilLength":          nilTwl.Length(),
			"nilIsEmpty":         nilTwl.IsEmpty(),
		}
		expected := args.Map{
			"hasLineNumber":      true,
			"isInvalidLineNumber": false,
			"length":             5,
			"isEmpty":            false,
			"isEmptyText":        false,
			"nilLength":          0,
			"nilIsEmpty":         true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber_BasicMethods returns non-empty -- with args", actual)
	})
}

// ── ValidValues ──

func Test_ValidValues_BasicMethods(t *testing.T) {
	safeTest(t, "Test_ValidValues_BasicMethods", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b").AddFull(false, "c", "err")

		// Act
		actual := args.Map{
			"length":   vvs.Length(),
			"hasAny":   vvs.HasAnyItem(),
			"isEmpty":  vvs.IsEmpty(),
			"safeAt0":  vvs.SafeValueAt(0),
			"safeAt10": vvs.SafeValueAt(10),
			"strLen":   len(vvs.Strings()),
		}
		expected := args.Map{
			"length":   3,
			"hasAny":   true,
			"isEmpty":  false,
			"safeAt0":  "a",
			"safeAt10": "",
			"strLen":   3,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues_BasicMethods returns non-empty -- with args", actual)
	})
}

// ── ValueStatus ──

func Test_ValueStatus_Creation(t *testing.T) {
	safeTest(t, "Test_ValueStatus_Creation", func() {
		// Arrange
		invalid := corestr.InvalidValueStatus("msg")
		noMsg := corestr.InvalidValueStatusNoMessage()

		// Act
		cloned := invalid.Clone()

		actual := args.Map{
			"invalidIsValid": invalid.ValueValid.IsValid,
			"noMsgIsValid":   noMsg.ValueValid.IsValid,
			"clonedIndex":    cloned.Index,
		}
		expected := args.Map{
			"invalidIsValid": false,
			"noMsgIsValid":   false,
			"clonedIndex":    invalid.Index,
		}
		expected.ShouldBeEqual(t, 0, "ValueStatus_Creation returns non-empty -- with args", actual)
	})
}

// ── KeyValueCollection ──

func Test_KeyValueCollection_BasicMethods(t *testing.T) {
	safeTest(t, "Test_KeyValueCollection_BasicMethods", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")

		// Act
		actual := args.Map{
			"length":     kvc.Length(),
			"isEmpty":    kvc.IsEmpty(),
			"hasAny":     kvc.HasAnyItem(),
			"hasKey":     kvc.HasKey("k1"),
			"containsK2": kvc.IsContains("k2"),
			"keysLen":    len(kvc.AllKeys()),
			"valsLen":    len(kvc.AllValues()),
		}
		expected := args.Map{
			"length":     2,
			"isEmpty":    false,
			"hasAny":     true,
			"hasKey":     true,
			"containsK2": true,
			"keysLen":    2,
			"valsLen":    2,
		}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection_BasicMethods returns correct value -- with args", actual)
	})
}

func Test_KeyValueCollection_Get(t *testing.T) {
	safeTest(t, "Test_KeyValueCollection_Get", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("key", "val")

		// Act
		val, found := kvc.Get("key")
		_, notFound := kvc.Get("missing")

		actual := args.Map{
			"val":      val,
			"found":    found,
			"notFound": notFound,
		}
		expected := args.Map{
			"val":      "val",
			"found":    true,
			"notFound": false,
		}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection_Get returns correct value -- with args", actual)
	})
}

// ── LinkedList ──

func Test_LinkedList_BasicMethods(t *testing.T) {
	safeTest(t, "Test_LinkedList_BasicMethods", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("a").Add("b").Add("c")

		// Act
		actual := args.Map{
			"length":   ll.Length(),
			"isEmpty":  ll.IsEmpty(),
			"hasItems": ll.HasItems(),
			"headElem": ll.Head().Element,
			"tailElem": ll.Tail().Element,
		}
		expected := args.Map{
			"length":   3,
			"isEmpty":  false,
			"hasItems": true,
			"headElem": "a",
			"tailElem": "c",
		}
		expected.ShouldBeEqual(t, 0, "LinkedList_BasicMethods returns correct value -- with args", actual)
	})
}

func Test_LinkedList_AddFront(t *testing.T) {
	safeTest(t, "Test_LinkedList_AddFront", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("b").AddFront("a")

		// Act
		actual := args.Map{
			"length":   ll.Length(),
			"headElem": ll.Head().Element,
		}
		expected := args.Map{
			"length":   2,
			"headElem": "a",
		}
		expected.ShouldBeEqual(t, 0, "LinkedList_AddFront returns correct value -- with args", actual)
	})
}

func Test_LinkedList_IsEquals(t *testing.T) {
	safeTest(t, "Test_LinkedList_IsEquals", func() {
		// Arrange
		ll1 := corestr.Empty.LinkedList()
		ll1.Add("a").Add("b")
		ll2 := corestr.Empty.LinkedList()
		ll2.Add("a").Add("b")
		ll3 := corestr.Empty.LinkedList()
		ll3.Add("a").Add("c")

		// Act
		actual := args.Map{
			"same": ll1.IsEquals(ll2),
			"diff": ll1.IsEquals(ll3),
		}
		expected := args.Map{
			"same": true,
			"diff": false,
		}
		expected.ShouldBeEqual(t, 0, "LinkedList_IsEquals returns correct value -- with args", actual)
	})
}

// ── Creators ──

func Test_NewCreators(t *testing.T) {
	safeTest(t, "Test_NewCreators", func() {
		// Act
		col := corestr.New.Collection.Empty()
		ss := corestr.New.SimpleSlice.Empty()
		hs := corestr.New.Hashset.Empty()
		hm := corestr.New.Hashmap.Empty()
		colCap := corestr.New.Collection.Cap(10)
		colClone := corestr.New.Collection.CloneStrings([]string{"a", "b"})
		colLine := corestr.New.Collection.LineUsingSep(",", "a,b,c")

		actual := args.Map{
			"colEmpty":     col.IsEmpty(),
			"ssEmpty":      ss.IsEmpty(),
			"hsEmpty":      hs.IsEmpty(),
			"hmEmpty":      hm.IsEmpty(),
			"colCapEmpty":  colCap.IsEmpty(),
			"colCloneLen":  colClone.Length(),
			"colLineLen":   colLine.Length(),
		}
		expected := args.Map{
			"colEmpty":     true,
			"ssEmpty":      true,
			"hsEmpty":      true,
			"hmEmpty":      true,
			"colCapEmpty":  true,
			"colCloneLen":  2,
			"colLineLen":   3,
		}
		expected.ShouldBeEqual(t, 0, "NewCreators returns correct value -- with args", actual)
	})
}

func Test_EmptyCreators(t *testing.T) {
	safeTest(t, "Test_EmptyCreators", func() {
		// Act
		col := corestr.Empty.Collection()
		ll := corestr.Empty.LinkedList()
		ss := corestr.Empty.SimpleSlice()
		kv := corestr.Empty.KeyValuePair()
		kvCol := corestr.Empty.KeyValueCollection()
		lr := corestr.Empty.LeftRight()
		hs := corestr.Empty.Hashset()
		hm := corestr.Empty.Hashmap()

		actual := args.Map{
			"colEmpty":   col.IsEmpty(),
			"llEmpty":    ll.IsEmpty(),
			"ssEmpty":    ss.IsEmpty(),
			"kvKey":      kv.Key,
			"kvColEmpty": kvCol.IsEmpty(),
			"lrLeft":     lr.Left,
			"hsEmpty":    hs.IsEmpty(),
			"hmEmpty":    hm.IsEmpty(),
		}
		expected := args.Map{
			"colEmpty":   true,
			"llEmpty":    true,
			"ssEmpty":    true,
			"kvKey":      "",
			"kvColEmpty": true,
			"lrLeft":     "",
			"hsEmpty":    true,
			"hmEmpty":    true,
		}
		expected.ShouldBeEqual(t, 0, "EmptyCreators returns empty -- with args", actual)
	})
}
