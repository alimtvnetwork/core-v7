package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Collection extended (C11) ──

func Test_SrcC11_Collection_JsonString_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_Collection_JsonString_Verification", func() {
		// Arrange
		tc := srcC11CollectionJsonStringTestCase
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"jsonString":     c.JsonString() != "",
			"jsonStringMust": c.JsonStringMust() != "",
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC11_Collection_State_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_Collection_State_Verification", func() {
		// Arrange
		tc := srcC11CollectionStateTestCase
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"hasAnyItem": c.HasAnyItem(),
			"lastIndex":  c.LastIndex(),
			"hasIdx1":    c.HasIndex(1),
			"hasIdx5":    c.HasIndex(5),
			"hasIdxNeg":  c.HasIndex(-1),
			"count":      c2.Count(),
			"lengthLock": c2.LengthLock(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC11_Collection_ListMethods_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_Collection_ListMethods_Verification", func() {
		// Arrange
		tc := srcC11CollectionListMethodsTestCase
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"listStrPtrLen": len(c.ListStringsPtr()),
			"listStrLen":    len(c.ListStrings()),
			"stringJSON":    c.StringJSON() != "",
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC11_Collection_RemoveAt_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_Collection_RemoveAt_Verification", func() {
		// Arrange
		tc := srcC11CollectionRemoveAtTestCase
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		ok := c.RemoveAt(1)
		actual := args.Map{
			"ok":        ok,
			"newLength": c.Length(),
			"fail100":   c.RemoveAt(100),
			"failNeg":   c.RemoveAt(-1),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC11_Collection_Capacity_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_Collection_Capacity_Verification", func() {
		// Arrange
		tc := srcC11CollectionCapacityTestCase
		c := corestr.New.Collection.Cap(10)
		empty := &corestr.Collection{}

		// Act
		actual := args.Map{
			"capGt0":   c.Capacity() > 0,
			"emptyCap": empty.Capacity(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC11_Collection_Equals_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_Collection_Equals_Verification", func() {
		// Arrange
		tc := srcC11CollectionEqualsTestCase
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})
		diff := corestr.New.Collection.Strings([]string{"b"})
		diffLen := corestr.New.Collection.Strings([]string{"a"})
		e1 := corestr.New.Collection.Cap(0)
		e2 := corestr.New.Collection.Cap(0)
		hA := corestr.New.Collection.Strings([]string{"Hello"})
		hB := corestr.New.Collection.Strings([]string{"hello"})

		// Act
		actual := args.Map{
			"equalSame":      a.IsEquals(b),
			"equalDiff":      a.IsEquals(diff),
			"equalBothEmpty": e1.IsEquals(e2),
			"equalDiffLen":   diffLen.IsEquals(a),
			"insensitive":    hA.IsEqualsWithSensitive(false, hB),
			"isEmptyLock":    e1.IsEmptyLock(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC11_Collection_HasItems_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_Collection_HasItems_Verification", func() {
		// Arrange
		tc := srcC11CollectionHasItemsTestCase
		c := corestr.New.Collection.Strings([]string{"a"})
		var nilC *corestr.Collection

		// Act
		actual := args.Map{
			"hasItems":    c.HasItems(),
			"nilHasItems": nilC.HasItems(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC11_Collection_AddMethods_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_Collection_AddMethods_Verification", func() {
		// Arrange
		tc := srcC11CollectionAddMethodsTestCase

		// Act
		cLock := corestr.New.Collection.Cap(5)
		cLock.AddLock("a")

		cNE := corestr.New.Collection.Cap(5)
		cNE.AddNonEmpty("")
		cNE.AddNonEmpty("a")

		cNEWS := corestr.New.Collection.Cap(5)
		cNEWS.AddNonEmptyWhitespace("  ")
		cNEWS.AddNonEmptyWhitespace("a")

		cErr := corestr.New.Collection.Cap(5)
		cErr.AddError(nil)

		cIfMany := corestr.New.Collection.Cap(5)
		cIfMany.AddIfMany(true, "a", "b")
		cIfMany.AddIfMany(false, "c")

		cFunc := corestr.New.Collection.Cap(5)
		cFunc.AddFunc(func() string { return "x" })

		cAL := corestr.New.Collection.Cap(5)
		cAL.AddsLock("a", "b")

		cAS := corestr.New.Collection.Cap(5)
		cAS.AddStrings([]string{"a", "b"})

		cAC := corestr.New.Collection.Cap(5)
		cAC.AddCollection(corestr.New.Collection.Strings([]string{"a"}))

		cACE := corestr.New.Collection.Cap(5)
		cACE.AddCollection(corestr.New.Collection.Cap(0))

		cACs := corestr.New.Collection.Cap(5)
		cACs.AddCollections(
			corestr.New.Collection.Strings([]string{"a"}),
			corestr.New.Collection.Strings([]string{"b"}),
		)

		actual := args.Map{
			"addLockLen":     cLock.Length(),
			"addNonEmpty":    cNE.Length(),
			"addNonEmptyWS": cNEWS.Length(),
			"addErrorNil":   cErr.Length(),
			"addIfMany":     cIfMany.Length(),
			"addFunc":       cFunc.Length(),
			"addsLock":      cAL.Length(),
			"addStrings":    cAS.Length(),
			"addCol":        cAC.Length(),
			"addColEmpty":   cACE.Length(),
			"addCols":       cACs.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC11_Collection_Error_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_Collection_Error_Verification", func() {
		// Arrange
		tc := srcC11CollectionErrorTestCase
		c := corestr.New.Collection.Strings([]string{"err1"})
		empty := corestr.New.Collection.Cap(0)

		// Act
		noPanic := !callPanicsSrcC11(func() {
			_ = c.ToError("\n")
			_ = c.ToDefaultError()
		})
		actual := args.Map{
			"asDefaultErr": c.AsDefaultError() != nil,
			"asErrorNil":   empty.AsError("\n") == nil,
			"noPanic":      noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC11_Collection_EachSplitConcat_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_Collection_EachSplitConcat_Verification", func() {
		// Arrange
		tc := srcC11CollectionEachSplitConcatTestCase
		c := corestr.New.Collection.Strings([]string{"a,b", "c,d"})
		c2 := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"splitLen":       len(c.EachItemSplitBy(",")),
			"concatLen":      c2.ConcatNew(0, "b", "c").Length(),
			"concatEmptyLen": c2.ConcatNew(0).Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── Hashmap (C11) ──

func Test_SrcC11_Hashmap_Methods_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_Hashmap_Methods_Verification", func() {
		// Arrange
		tc := srcC11HashmapMethodsTestCase

		// Act
		noPanic := !callPanicsSrcC11(func() {
			hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k", Value: "v"})
			_ = hm.HasItems()
			_ = hm.Collection()
			_ = corestr.New.Hashmap.Cap(5).IsEmptyLock()
			hm2 := corestr.New.Hashmap.Cap(5)
			hm2.AddOrUpdateKeyStrValInt("k", 42)
			hm2.AddOrUpdateKeyStrValFloat("f", 3.14)
			hm2.AddOrUpdateKeyStrValFloat64("f64", 3.14)
			hm2.AddOrUpdateKeyStrValAny("a", 42)
			hm2.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "av", Value: 42})
			_ = hm2.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "kv", Value: "v"})
			_ = hm2.AddOrUpdateKeyVal(corestr.KeyValuePair{Key: "kv", Value: "v2"})
			_ = hm.Contains("k")
			_ = hm.ContainsLock("k")
			_ = hm.IsKeyMissing("k")
			_ = hm.IsKeyMissing("x")
			_ = hm.IsKeyMissingLock("k")
			_ = hm.HasLock("k")
			hm3 := corestr.New.Hashmap.KeyValues(
				corestr.KeyValuePair{Key: "a", Value: "1"},
				corestr.KeyValuePair{Key: "b", Value: "2"},
			)
			_ = hm3.HasAllStrings("a", "b")
			_ = hm3.HasAllStrings("a", "c")
			hm4 := corestr.New.Hashmap.Cap(5)
			hm4.AddOrUpdateLock("k", "v")
			hm4.AddOrUpdateKeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
			hm4.AddOrUpdateKeyAnyValues(corestr.KeyAnyValuePair{Key: "b", Value: 1})
			hm4.AddOrUpdateMap(map[string]string{"c": "3"})
			hm4.AddsOrUpdates(corestr.KeyValuePair{Key: "d", Value: "4"})
			hm5 := corestr.New.Hashmap.Cap(5)
			hm5.AddOrUpdateHashmap(corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "x", Value: "1"}))
			hm5.AddOrUpdateHashmap(nil)
			hm5.AddOrUpdateCollection(
				corestr.New.Collection.Strings([]string{"a", "b"}),
				corestr.New.Collection.Strings([]string{"1", "2"}),
			)
			hm5.AddOrUpdateCollection(nil, nil)
			hm5.AddOrUpdateCollection(
				corestr.New.Collection.Strings([]string{"a"}),
				corestr.New.Collection.Strings([]string{"1", "2"}),
			)
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC11_Hashmap_Set_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_Hashmap_Set_Verification", func() {
		// Arrange
		tc := srcC11HashmapSetTestCase

		// Act
		hm1 := corestr.New.Hashmap.Cap(5)
		hm1.Set("k", "v")
		hm2 := corestr.New.Hashmap.Cap(5)
		hm2.SetTrim("  k  ", "  v  ")
		hm3 := corestr.New.Hashmap.Cap(5)
		hm3.SetBySplitter("=", "key=value")
		splitVal, _ := hm3.Get("key")
		hm4 := corestr.New.Hashmap.Cap(5)
		hm4.SetBySplitter("=", "keyonly")
		noSplitVal, _ := hm4.Get("keyonly")
		actual := args.Map{
			"setHas":     hm1.Has("k"),
			"setTrimHas": hm2.Has("k"),
			"splitVal":   splitVal,
			"noSplitVal": noSplitVal,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC11_Hashmap_Concat_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_Hashmap_Concat_Verification", func() {
		// Arrange
		tc := srcC11HashmapConcatTestCase
		hm1 := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		hm2 := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "b", Value: "2"})

		// Act
		actual := args.Map{
			"concatGe2":    hm1.ConcatNew(false, hm2).Length() >= 2,
			"concatEGe1":   hm1.ConcatNew(true).Length() >= 1,
			"concatMapGe2": hm1.ConcatNewUsingMaps(false, map[string]string{"b": "2"}).Length() >= 2,
			"concatMEGe1":  hm1.ConcatNewUsingMaps(true).Length() >= 1,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── Hashset (C11) ──

func Test_SrcC11_Hashset_HasItems_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_Hashset_HasItems_Verification", func() {
		// Arrange
		tc := srcC11HashsetHasItemsTestCase
		hs := corestr.New.Hashset.Strings([]string{"a"})
		var nilHs *corestr.Hashset

		// Act
		actual := args.Map{
			"hasItems":    hs.HasItems(),
			"nilHasItems": nilHs.HasItems(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── SimpleSlice (C11) ──

func Test_SrcC11_SimpleSlice_HasItems_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_SimpleSlice_HasItems_Verification", func() {
		// Arrange
		tc := srcC11SimpleSliceHasItemsTestCase
		ss := corestr.New.SimpleSlice.SpreadStrings("a")

		// Act
		actual := args.Map{
			"hasAnyItem": ss.HasAnyItem(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── Types (C11) ──

func Test_SrcC11_ValidValue_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_ValidValue_Verification", func() {
		// Arrange
		tc := srcC11ValidValueTestCase

		// Act
		vv := corestr.NewValidValue("hello")
		iv := corestr.InvalidValidValue("x")
		actual := args.Map{
			"validIsValid":   vv.IsValid,
			"validValue":     vv.Value,
			"invalidIsValid": iv.IsValid,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC11_LeftRight_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_LeftRight_Verification", func() {
		// Arrange
		tc := srcC11LeftRightTestCase

		// Act
		lr := corestr.NewLeftRight("key", "val")
		lrNo := corestr.NewLeftRight("nosplit", "")
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{
			"lrLeft":    lr.Left,
			"lrRight":   lr.Right,
			"noSplitL":  lrNo.Left,
			"lmrLeft":   lmr.Left,
			"lmrMiddle": lmr.Middle,
			"lmrRight":  lmr.Right,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── CollectionsOfCollection (C11) ──

func Test_SrcC11_Coc_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_Coc_Verification", func() {
		// Arrange
		tc := srcC11CocTestCase

		// Act
		noPanic := !callPanicsSrcC11(func() {
			coc := corestr.New.CollectionsOfCollection.Cap(5)
			_ = coc.HasItems()
			c := corestr.New.Collection.Strings([]string{"a"})
			coc.Add(c)
			_ = coc.HasItems()
			_ = coc.AllIndividualItemsLength()
			_ = coc.Items()
			_ = coc.List(0)
			_ = coc.ToCollection()
			_ = coc.String()
			_ = coc.JsonModel()
			_ = coc.JsonModelAny()
			_ = coc.AsJsoner()
			_ = coc.AsJsonParseSelfInjector()
			_ = coc.AsJsonMarshaller()
			_ = coc.AsJsonContractsBinder()
			coc2 := corestr.New.CollectionsOfCollection.Cap(5)
			coc2.AddStrings(false, []string{"a", "b"})
			coc2.AddStrings(false, nil)
			coc2.AddsStringsOfStrings(false, []string{"a"}, []string{"b"})
			c2 := *corestr.New.Collection.Strings([]string{"a"})
			coc2.Adds(c2)
			coc2.AddCollections(c2)
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── CloneSlice (C11) ──

func Test_SrcC11_CloneSlice_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_CloneSlice_Verification", func() {
		// Arrange
		tc := srcC11CloneSliceTestCase

		// Act
		orig := []string{"a", "b"}
		cloned := corestr.CloneSlice(orig)
		orig[0] = "X"
		actual := args.Map{
			"deepClone":    cloned[0],
			"nilLen":       len(corestr.CloneSlice(nil)),
			"cloneIfTrue":  len(corestr.CloneSliceIf(true, "a", "b")),
			"cloneIfFalse": len(corestr.CloneSliceIf(false, "a")),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── KeyAnyValuePair (C11) ──

func Test_SrcC11_KeyAnyValuePair_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_KeyAnyValuePair_Verification", func() {
		// Arrange
		tc := srcC11KeyAnyValuePairTestCase

		// Act
		kv := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		actual := args.Map{
			"nonEmpty": kv.ValueString() != "",
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── AllIndividualStringsOfStringsLength (C11) ──

func Test_SrcC11_AllIndStrLen_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC11_AllIndStrLen_Verification", func() {
		// Arrange
		tc := srcC11AllIndStrLenTestCase

		// Act
		input := [][]string{{"a", "b"}, {"c"}}
		actual := args.Map{
			"length": corestr.AllIndividualStringsOfStringsLength(&input),
			"nilLen": corestr.AllIndividualStringsOfStringsLength(nil),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC11(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}
