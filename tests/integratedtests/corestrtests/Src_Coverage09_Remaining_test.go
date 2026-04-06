package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── SimpleStringOnce ──

func Test_SrcC09_SimpleStringOnce_Core_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleStringOnce_Core_Verification", func() {
		// Arrange
		tc := srcC09SimpleStringOnceCoreTestCase
		s := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{
			"value":             s.Value(),
			"isInitialized":     s.IsInitialized(),
			"isDefined":         s.IsDefined(),
			"isUninitialized":   s.IsUninitialized(),
			"isInvalid":         s.IsInvalid(),
			"safeValue":         s.SafeValue(),
			"isEmpty":           s.IsEmpty(),
			"isWhitespace":      s.IsWhitespace(),
			"trim":              s.Trim(),
			"hasValidNonEmpty":  s.HasValidNonEmpty(),
			"hasValidNonWS":     s.HasValidNonWhitespace(),
			"hasSafeNonEmpty":   s.HasSafeNonEmpty(),
			"isHello":           s.Is("hello"),
			"isWorld":           s.Is("world"),
			"isAnyOfHello":      s.IsAnyOf("hello"),
			"isAnyOfX":          s.IsAnyOf("x"),
			"isContainsHel":     s.IsContains("hel"),
			"isAnyContainsHel":  s.IsAnyContains("hel"),
			"isEqualNonSens":    s.IsEqualNonSensitive("HELLO"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_SimpleStringOnce_Set_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleStringOnce_Set_Verification", func() {
		// Arrange
		tc := srcC09SimpleStringOnceSetTestCase
		s := corestr.New.SimpleStringOnce.Uninitialized("")

		// Act
		err1 := s.SetOnUninitialized("val")
		err2 := s.SetOnUninitialized("val2")
		actual := args.Map{
			"firstErr":  err1 == nil,
			"secondErr": err2 == nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_SimpleStringOnce_GetSetOnce_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleStringOnce_GetSetOnce_Verification", func() {
		// Arrange
		tc := srcC09SimpleStringOnceGetSetOnceTestCase
		s := corestr.New.SimpleStringOnce.Uninitialized("")

		// Act
		actual := args.Map{
			"first":  s.GetSetOnce("first"),
			"second": s.GetSetOnce("second"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_SimpleStringOnce_GetOnce_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleStringOnce_GetOnce_Verification", func() {
		// Arrange
		tc := srcC09SimpleStringOnceGetOnceTestCase
		s := corestr.New.SimpleStringOnce.Uninitialized("")

		// Act
		actual := args.Map{
			"value": s.GetOnce(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_SimpleStringOnce_GetOnceFunc_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleStringOnce_GetOnceFunc_Verification", func() {
		// Arrange
		tc := srcC09SimpleStringOnceGetOnceFuncTestCase
		s := corestr.New.SimpleStringOnce.Uninitialized("")

		// Act
		v1 := s.GetOnceFunc(func() string { return "computed" })
		v2 := s.GetOnceFunc(func() string { return "other" })
		actual := args.Map{
			"first":  v1,
			"second": v2,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_SimpleStringOnce_SetOnceIf_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleStringOnce_SetOnceIf_Verification", func() {
		// Arrange
		tc := srcC09SimpleStringOnceSetOnceIfTestCase
		s := corestr.New.SimpleStringOnce.Uninitialized("")

		// Act
		actual := args.Map{
			"first":  s.SetOnceIfUninitialized("val"),
			"second": s.SetOnceIfUninitialized("val2"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_SimpleStringOnce_Invalidate_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleStringOnce_Invalidate_Verification", func() {
		// Arrange
		tc := srcC09SimpleStringOnceInvalidateTestCase
		s := corestr.New.SimpleStringOnce.Init("hello")
		s2 := corestr.New.SimpleStringOnce.Init("world")

		// Act
		s.Invalidate()
		s2.Reset()
		actual := args.Map{
			"afterInvalidate": s.IsInitialized(),
			"afterReset":      s2.IsInitialized(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_SimpleStringOnce_Conversions_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleStringOnce_Conversions_Verification", func() {
		// Arrange
		tc := srcC09SimpleStringOnceConversionsTestCase
		s := corestr.New.SimpleStringOnce.Init("42")
		s2 := corestr.New.SimpleStringOnce.Init("3.14")
		s3 := corestr.New.SimpleStringOnce.Init("true")
		s4 := corestr.New.SimpleStringOnce.Init("yes")
		s5 := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		noPanic := !callPanicsSrcC09(func() {
			_ = s5.Byte()
			_ = s5.Int16()
			_ = s5.Int32()
			_ = s5.ValueByte(0)
			_ = s5.IsSetter(false)
			_ = s5.IsSetter(true)
			_ = s.ValueBytes()
			_ = s.ValueBytesPtr()
		})
		actual := args.Map{
			"int":       s.Int(),
			"valInt":    s.ValueInt(0),
			"defInt":    s.ValueDefInt(),
			"byte":      s.Byte(),
			"valByte":   s.ValueByte(0),
			"defByte":   s.ValueDefByte(),
			"float64":   s2.ValueFloat64(0) != 0,
			"defFloat":  s2.ValueDefFloat64() != 0,
			"boolTrue":  s3.Boolean(false),
			"boolDef":   s3.BooleanDefault(),
			"isValBool": s3.IsValueBool(),
			"boolYes":   s4.Boolean(false),
			"intAbc":    s5.Int(),
			"noPanic":   noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_SimpleStringOnce_WithinRange_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleStringOnce_WithinRange_Verification", func() {
		// Arrange
		tc := srcC09SimpleStringOnceWithinRangeTestCase
		s := corestr.New.SimpleStringOnce.Init("50")

		// Act
		v, ok := s.WithinRange(true, 0, 100)
		v2, ok2 := s.WithinRangeDefault(0, 100)
		noPanic := !callPanicsSrcC09(func() {
			_, _ = s.Uint16()
			_, _ = s.Uint32()
		})
		actual := args.Map{
			"value":   v,
			"ok":      ok,
			"defVal":  v2,
			"defOk":   ok2,
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_SimpleStringOnce_Concat_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleStringOnce_Concat_Verification", func() {
		// Arrange
		tc := srcC09SimpleStringOnceConcatTestCase
		s := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		c := s.ConcatNew(" world")
		noPanic := !callPanicsSrcC09(func() {
			_ = s.ConcatNewUsingStrings(" ", "beautiful", "world")
		})
		actual := args.Map{
			"value":   c.Value(),
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_SimpleStringOnce_Split_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleStringOnce_Split_Verification", func() {
		// Arrange
		tc := srcC09SimpleStringOnceSplitTestCase
		s := corestr.New.SimpleStringOnce.Init("a,b,c")

		// Act
		noPanic := !callPanicsSrcC09(func() {
			_ = s.Split(",")
			_ = s.SplitNonEmpty(",")
			_ = s.SplitTrimNonWhitespace(",")
			_, _ = s.SplitLeftRight(",")
			_, _ = s.SplitLeftRightTrim(",")
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}
func Test_SrcC09_SimpleStringOnce_Json_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleStringOnce_Json_Verification", func() {
		// Arrange
		tc := srcC09SimpleStringOnceJsonTestCase
		s := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		noPanic := !callPanicsSrcC09(func() {
			_ = s.JsonModel()
			_ = s.JsonModelAny()
			_, _ = s.MarshalJSON()
			_, _ = s.Serialize()
			_ = s.AsJsoner()
			_ = s.AsJsonContractsBinder()
			_ = s.AsJsonParseSelfInjector()
			_ = s.AsJsonMarshaller()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── KeyValueCollection ──

func Test_SrcC09_KeyValueCollection_Core_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_KeyValueCollection_Core_Verification", func() {
		// Arrange
		tc := srcC09KeyValueCollectionCoreTestCase
		kv := corestr.New.KeyValues.Empty()

		// Act
		kv.Add("k1", "v1").Add("k2", "v2")
		v, found := kv.Get("k1")
		noPanic := !callPanicsSrcC09(func() {
			_ = kv.First()
			_ = kv.Last()
			_ = kv.FirstOrDefault()
			_ = kv.LastOrDefault()
			_ = kv.Strings()
			_ = kv.String()
			_ = kv.AllKeys()
			_ = kv.AllKeysSorted()
			_ = kv.AllValues()
			_ = kv.Join(",")
			_ = kv.JoinKeys(",")
			_ = kv.JoinValues(",")
			_ = kv.Compile()
		})
		actual := args.Map{
			"length":    kv.Length(),
			"count":     kv.Count(),
			"lastIndex": kv.LastIndex(),
			"hasIdx0":   kv.HasIndex(0),
			"hasIdx5":   kv.HasIndex(5),
			"hasKey":    kv.HasKey("k1"),
			"contains":  kv.IsContains("k1"),
			"getValue":  v,
			"found":     found,
			"noPanic":   noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_KeyValueCollection_Add_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_KeyValueCollection_Add_Verification", func() {
		// Arrange
		tc := srcC09KeyValueCollectionAddTestCase

		// Act
		noPanic := !callPanicsSrcC09(func() {
			kv := corestr.New.KeyValues.Empty()
			kv.AddIf(false, "skip", "val")
			kv.AddIf(true, "k", "v")
			kv.Adds(corestr.KeyValuePair{Key: "a", Value: "b"})
			kv.Adds()
			kv.AddMap(map[string]string{"c": "d"})
			kv.AddMap(nil)
			kv.AddHashsetMap(map[string]bool{"e": true})
			kv.AddHashsetMap(nil)
			kv.AddHashset(corestr.New.Hashset.StringsSpreadItems("f"))
			kv.AddHashset(nil)
			kv.AddsHashmap(corestr.New.Hashmap.UsingMap(map[string]string{"g": "h"}))
			kv.AddsHashmap(nil)
			kv.AddsHashmaps(corestr.New.Hashmap.UsingMap(map[string]string{"i": "j"}))
			kv.AddsHashmaps()
			kv.AddStringBySplit("=", "k=l")
			kv.AddStringBySplitTrim("=", " m = n ")
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_KeyValueCollection_Find_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_KeyValueCollection_Find_Verification", func() {
		// Arrange
		tc := srcC09KeyValueCollectionFindTestCase
		kv := corestr.New.KeyValues.Empty()
		kv.Add("a", "1").Add("b", "2")

		// Act
		r := kv.Find(func(i int, curr corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return curr, curr.Key == "a", false
		})
		actual := args.Map{
			"resultLen": len(r),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_KeyValueCollection_Safe_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_KeyValueCollection_Safe_Verification", func() {
		// Arrange
		tc := srcC09KeyValueCollectionSafeTestCase
		kv := corestr.New.KeyValues.Empty()
		kv.Add("a", "1")

		// Act
		noPanic := !callPanicsSrcC09(func() {
			_ = kv.SafeValuesAtIndexes(0)
			_ = kv.StringsUsingFormat("%s=%s")
			_ = kv.Hashmap()
			_ = kv.Map()
		})
		actual := args.Map{
			"safeVal0":  kv.SafeValueAt(0),
			"safeVal99": kv.SafeValueAt(99),
			"noPanic":   noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_KeyValueCollection_Json_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_KeyValueCollection_Json_Verification", func() {
		// Arrange
		tc := srcC09KeyValueCollectionJsonTestCase
		kv := corestr.New.KeyValues.Empty()
		kv.Add("a", "1")

		// Act
		noPanic := !callPanicsSrcC09(func() {
			_ = kv.JsonModel()
			_ = kv.JsonModelAny()
			_, _ = kv.Serialize()
			_, _ = kv.MarshalJSON()
			_ = kv.SerializeMust()
			_ = kv.AsJsoner()
			_ = kv.AsJsonContractsBinder()
			_ = kv.AsJsonParseSelfInjector()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_KeyValueCollection_Clear_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_KeyValueCollection_Clear_Verification", func() {
		// Arrange
		tc := srcC09KeyValueCollectionClearTestCase
		var nilKv *corestr.KeyValueCollection

		// Act
		noPanic := !callPanicsSrcC09(func() {
			kv := corestr.New.KeyValues.Empty()
			kv.Add("a", "1")
			kv.Clear()
			kv.Dispose()
			nilKv.Clear()
			nilKv.Dispose()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── NonChainedLinkedListNodes ──

func Test_SrcC09_NonChainedLLNodes_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_NonChainedLLNodes_Verification", func() {
		// Arrange
		tc := srcC09NonChainedLLNodesTestCase
		nc := corestr.NewNonChainedLinkedListNodes(5)

		// Act
		n1 := &corestr.LinkedListNode{Element: "a"}
		n2 := &corestr.LinkedListNode{Element: "b"}
		nc.Adds(n1, n2)
		chainingBefore := nc.IsChainingApplied()
		noPanic := !callPanicsSrcC09(func() {
			_ = nc.FirstOrDefault()
			_ = nc.LastOrDefault()
			_ = nc.Items()
			nc.ApplyChaining()
			_ = nc.ToChainedNodes()
		})
		actual := args.Map{
			"length":         nc.Length(),
			"firstElement":   nc.First().Element,
			"lastElement":    nc.Last().Element,
			"chainingBefore": chainingBefore,
			"chainingAfter":  nc.IsChainingApplied(),
			"noPanic":        noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── NonChainedLinkedCollectionNodes ──

func Test_SrcC09_NonChainedLCNodes_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_NonChainedLCNodes_Verification", func() {
		// Arrange
		tc := srcC09NonChainedLCNodesTestCase
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})

		// Act
		n1 := &corestr.LinkedCollectionNode{Element: c1}
		n2 := &corestr.LinkedCollectionNode{Element: c2}
		nc.Adds(n1, n2)
		noPanic := !callPanicsSrcC09(func() {
			_ = nc.First()
			_ = nc.Last()
			_ = nc.FirstOrDefault()
			_ = nc.LastOrDefault()
			_ = nc.Items()
			nc.ApplyChaining()
			_ = nc.ToChainedNodes()
		})
		actual := args.Map{
			"length":        nc.Length(),
			"chainingAfter": nc.IsChainingApplied(),
			"noPanic":       noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── CollectionsOfCollection ──

func Test_SrcC09_CollectionsOfCollection_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_CollectionsOfCollection_Verification", func() {
		// Arrange
		tc := srcC09CollectionsOfCollectionTestCase
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		coc.Add(corestr.New.Collection.Strings([]string{"c"}))

		// Act
		noPanic := !callPanicsSrcC09(func() {
			_ = coc.ToCollection()
			_ = coc.Items()
			_ = coc.String()
			_ = coc.JsonModel()
			_ = coc.JsonModelAny()
			_, _ = coc.MarshalJSON()
			_ = coc.AsJsoner()
			_ = coc.AsJsonContractsBinder()
			_ = coc.AsJsonParseSelfInjector()
			_ = coc.AsJsonMarshaller()
		})
		actual := args.Map{
			"length":     coc.Length(),
			"allItemLen": coc.AllIndividualItemsLength(),
			"listLen":    len(coc.List(0)),
			"noPanic":    noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}
func Test_SrcC09_HashsetsCollection_HasAll_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_HashsetsCollection_HasAll_Verification", func() {
		// Arrange
		tc := srcC09HashsetsCollectionHasAllTestCase
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.StringsSpreadItems("a", "b"))
		empty := corestr.New.HashsetsCollection.Empty()

		// Act
		actual := args.Map{
			"hasAB":    hc.HasAll("a", "b"),
			"hasZ":     hc.HasAll("z"),
			"emptyHas": empty.HasAll("a"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_HashsetsCollection_Concat_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_HashsetsCollection_Concat_Verification", func() {
		// Arrange
		tc := srcC09HashsetsCollectionConcatTestCase

		// Act
		noPanic := !callPanicsSrcC09(func() {
			hc := corestr.New.HashsetsCollection.Empty()
			hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
			hc2 := corestr.New.HashsetsCollection.Empty()
			hc2.Add(corestr.New.Hashset.StringsSpreadItems("b"))
			_ = hc.ConcatNew(hc2)
			_ = hc.ConcatNew()
			hc.AddHashsetsCollection(hc2)
			hc.AddHashsetsCollection(nil)
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── CharCollectionMap ──

func Test_SrcC09_CharCollectionMap_Methods_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_CharCollectionMap_Methods_Verification", func() {
		// Arrange
		tc := srcC09CharCollectionMapMethodsTestCase
		ccm := corestr.New.CharCollectionMap.Items([]string{"abc", "adef", "bcd"})

		// Act
		noPanic := !callPanicsSrcC09(func() {
			_ = ccm.AllLengthsSum()
			_ = ccm.AllLengthsSumLock()
			_ = ccm.LengthLock()
			_ = ccm.GetMap()
			_ = ccm.GetCopyMapLock()
			_ = ccm.List()
			_ = ccm.ListLock()
			_ = ccm.SortedListAsc()
			_ = ccm.String()
			_ = ccm.StringLock()
			_ = ccm.SummaryString()
			_ = ccm.SummaryStringLock()
			_ = ccm.LengthOf('a')
			_ = ccm.LengthOfLock('a')
			_ = ccm.LengthOfCollectionFromFirstChar("abc")
			_ = ccm.Has("abc")
			_ = ccm.Has("zzz")
			_, _ = ccm.HasWithCollection("abc")
			_, _ = ccm.HasWithCollectionLock("abc")
			_ = ccm.GetCollection("a", true)
			_ = ccm.GetCollectionLock("a", false)
			_ = ccm.GetCollectionByChar('a')
			_ = ccm.HashsetByChar('a')
			_ = ccm.HashsetByCharLock('a')
			_ = ccm.HashsetByStringFirstChar("abc")
			_ = ccm.HashsetByStringFirstCharLock("abc")
			_ = ccm.HashsetsCollection()
			_ = ccm.HashsetsCollectionByChars('a')
			_ = ccm.HashsetsCollectionByStringFirstChar("abc")
			_ = ccm.IsEquals(ccm)
			_ = ccm.IsEqualsLock(ccm)
			_ = ccm.IsEqualsCaseSensitive(true, ccm)
			_ = ccm.IsEqualsCaseSensitiveLock(false, ccm)
		})
		actual := args.Map{
			"isEmpty":  ccm.IsEmpty(),
			"noPanic":  noPanic,
			"getCharE": ccm.GetChar(""),
			"getCharA": ccm.GetChar("a"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_CharCollectionMap_Add_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_CharCollectionMap_Add_Verification", func() {
		// Arrange
		tc := srcC09CharCollectionMapAddTestCase

		// Act
		noPanic := !callPanicsSrcC09(func() {
			ccm := corestr.New.CharCollectionMap.Empty()
			ccm.Add("hello")
			ccm.AddLock("world")
			ccm.AddStrings("a", "b")
			ccm.AddStrings()
			ccm.AddCollectionItems(corestr.New.Collection.Strings([]string{"c"}))
			ccm.AddCollectionItems(nil)
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_CharCollectionMap_Clear_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_CharCollectionMap_Clear_Verification", func() {
		// Arrange
		tc := srcC09CharCollectionMapClearTestCase
		var nilCcm *corestr.CharCollectionMap

		// Act
		noPanic := !callPanicsSrcC09(func() {
			ccm := corestr.New.CharCollectionMap.Items([]string{"a"})
			ccm.Clear()
			ccm.Dispose()
			nilCcm.Dispose()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── CharHashsetMap ──

func Test_SrcC09_CharHashsetMap_Methods_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_CharHashsetMap_Methods_Verification", func() {
		// Arrange
		tc := srcC09CharHashsetMapMethodsTestCase
		chm := corestr.New.CharHashsetMap.CapItems(20, 20, "abc", "adef", "bcd")

		// Act
		noPanic := !callPanicsSrcC09(func() {
			_ = chm.Length()
			_ = chm.LengthLock()
			_ = chm.AllLengthsSum()
			_ = chm.AllLengthsSumLock()
			_ = chm.GetMap()
			_ = chm.GetCopyMapLock()
			_ = chm.List()
			_ = chm.SortedListAsc()
			_ = chm.SortedListDsc()
			_ = chm.String()
			_ = chm.StringLock()
			_ = chm.SummaryString()
			_ = chm.SummaryStringLock()
			_ = chm.GetCharOf("")
			_ = chm.GetCharOf("a")
			_ = chm.LengthOf('a')
			_ = chm.LengthOfLock('a')
			_ = chm.LengthOfHashsetFromFirstChar("abc")
			_ = chm.Has("abc")
			_, _ = chm.HasWithHashset("abc")
			_, _ = chm.HasWithHashsetLock("abc")
			_ = chm.GetHashset("a", true)
			_ = chm.GetHashsetLock(true, "a")
			_ = chm.GetHashsetByChar('a')
			_ = chm.HashsetByChar('a')
			_ = chm.HashsetByCharLock('a')
			_ = chm.HashsetByStringFirstChar("abc")
			_ = chm.HashsetByStringFirstCharLock("abc")
			_ = chm.HashsetsCollection()
			_ = chm.HashsetsCollectionByChars('a')
			_ = chm.HashsetsCollectionByStringsFirstChar("abc")
			_ = chm.IsEquals(chm)
			_ = chm.IsEqualsLock(chm)
		})
		actual := args.Map{
			"isEmpty": chm.IsEmpty(),
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_CharHashsetMap_Add_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_CharHashsetMap_Add_Verification", func() {
		// Arrange
		tc := srcC09CharHashsetMapAddTestCase

		// Act
		noPanic := !callPanicsSrcC09(func() {
			chm := corestr.New.CharHashsetMap.Cap(20, 20)
			chm.Add("hello")
			chm.AddLock("world")
			chm.AddStrings("a", "b")
			chm.AddStrings()
			chm.AddStringsLock("c")
			chm.AddStringsLock()
			chm.AddCollectionItems(corestr.New.Collection.Strings([]string{"d"}))
			chm.AddCollectionItems(nil)
			chm.AddHashsetItems(corestr.New.Hashset.StringsSpreadItems("e"))
			chm.AddCharCollectionMapItems(corestr.New.CharCollectionMap.Items([]string{"f"}))
			chm.AddCharCollectionMapItems(nil)
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_CharHashsetMap_Clear_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_CharHashsetMap_Clear_Verification", func() {
		// Arrange
		tc := srcC09CharHashsetMapClearTestCase

		// Act
		noPanic := !callPanicsSrcC09(func() {
			chm := corestr.New.CharHashsetMap.CapItems(20, 20, "a")
			chm.Clear()
			chm.RemoveAll()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── LinkedCollections ──

func Test_SrcC09_LinkedCollections_Basic_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_LinkedCollections_Basic_Verification", func() {
		// Arrange
		tc := srcC09LinkedCollectionsBasicTestCase
		lc := corestr.New.LinkedCollection.Create()

		// Act
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		noPanic := !callPanicsSrcC09(func() {
			_ = lc.Head()
			_ = lc.Tail()
			_ = lc.First()
			_ = lc.Last()
			_ = lc.FirstOrDefault()
			_ = lc.LastOrDefault()
			_ = lc.AllIndividualItemsLength()
			_ = lc.ToStrings()
			_ = lc.ToCollectionSimple()
			_ = lc.ToCollection(0)
			_ = lc.ToCollectionsOfCollection(0)
			_ = lc.ItemsOfItems()
			_ = lc.ItemsOfItemsCollection()
			_ = lc.SimpleSlice()
		})
		actual := args.Map{
			"length":  lc.Length(),
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_LinkedCollections_Add_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_LinkedCollections_Add_Verification", func() {
		// Arrange
		tc := srcC09LinkedCollectionsAddTestCase

		// Act
		noPanic := !callPanicsSrcC09(func() {
			lc := corestr.New.LinkedCollection.Create()
			lc.AddStrings("a", "b")
			lc.AddStrings()
			lc.AddStringsLock("c")
			lc.AddStringsLock()
			lc.AddCollection(corestr.New.Collection.Strings([]string{"d"}))
			lc.AddCollection(nil)
			lc.AddLock(corestr.New.Collection.Strings([]string{"e"}))
			lc.Push(corestr.New.Collection.Strings([]string{"f"}))
			lc.PushBack(corestr.New.Collection.Strings([]string{"g"}))
			lc.PushBackLock(corestr.New.Collection.Strings([]string{"h"}))
			lc.PushFront(corestr.New.Collection.Strings([]string{"i"}))
			lc.AddFront(corestr.New.Collection.Strings([]string{"j"}))
			lc.AddFrontLock(corestr.New.Collection.Strings([]string{"k"}))
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_LinkedCollections_Loop_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_LinkedCollections_Loop_Verification", func() {
		// Arrange
		tc := srcC09LinkedCollectionsLoopTestCase
		lc := corestr.New.LinkedCollection.Strings("a", "b")
		count := 0

		// Act
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return false
		})
		actual := args.Map{
			"countGt0": count > 0,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_LinkedCollections_Equals_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_LinkedCollections_Equals_Verification", func() {
		// Arrange
		tc := srcC09LinkedCollectionsEqualsTestCase
		lc1 := corestr.New.LinkedCollection.Strings("a")
		lc2 := corestr.New.LinkedCollection.Strings("a")

		// Act
		actual := args.Map{
			"isEquals": lc1.IsEqualsPtr(lc2),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── Collection remaining ──

func Test_SrcC09_Collection_Remaining_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_Collection_Remaining_Verification", func() {
		// Arrange
		tc := srcC09CollectionRemainingTestCase
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		noPanic := !callPanicsSrcC09(func() {
			_ = c.First()
			_ = c.Last()
			_ = c.FirstOrDefault()
			_ = c.LastOrDefault()
			_ = c.IndexAt(0)
			_ = c.SafeIndexAtUsingLength("def", 3, 0)
			_ = c.List()
			_ = c.HasItems()
			_ = c.Reverse()
			_ = c.GetPagesSize(2)
			_ = c.Take(2)
			_ = c.Skip(0)
			_ = c.UniqueList()
			_ = c.UniqueListLock()
			_ = c.UniqueBoolMap()
			_ = c.UniqueBoolMapLock()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_Collection_Filter_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_Collection_Filter_Verification", func() {
		// Arrange
		tc := srcC09CollectionFilterTestCase
		c := corestr.New.Collection.Strings([]string{"a", "bb", "ccc"})
		f := func(s string, i int) (string, bool, bool) { return s, len(s) > 1, false }

		// Act
		r := c.Filter(f)
		noPanic := !callPanicsSrcC09(func() {
			_ = c.FilteredCollection(f)
			_ = c.FilterLock(f)
			_ = c.FilteredCollectionLock(f)
		})
		actual := args.Map{
			"filterLen": len(r),
			"noPanic":   noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_Collection_AppendAnys_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_Collection_AppendAnys_Verification", func() {
		// Arrange
		tc := srcC09CollectionAppendAnysTestCase

		// Act
		noPanic := !callPanicsSrcC09(func() {
			c := corestr.New.Collection.Empty()
			c.AppendAnys("a", 42, nil)
			c.AppendAnys()
			c.AppendAnysLock("b")
			c.AppendAnysLock()
			c.AppendNonEmptyAnys("c", nil)
			c.AppendNonEmptyAnys()
			c.AddsNonEmpty("d", "", "e")
			c.AddsNonEmpty()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── SimpleSlice remaining ──

func Test_SrcC09_SimpleSlice_Remaining_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleSlice_Remaining_Verification", func() {
		// Arrange
		tc := srcC09SimpleSliceRemainingTestCase
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		noPanic := !callPanicsSrcC09(func() {
			_ = s.Join(",")
			_ = s.JoinLine()
			_ = s.JoinLineEofLine()
			_ = s.JoinSpace()
			_ = s.JoinComma()
			_ = s.JoinCsv()
			_ = s.JoinCsvLine()
			_ = s.JoinWith(",")
			_ = s.JoinCsvString(",")
			_ = s.CsvStrings()
			_ = s.String()
			_ = s.Collection(false)
			_ = s.ToCollection(false)
			_ = s.NonPtr()
			_ = s.Ptr()
			_ = s.ToPtr()
			_ = s.ToNonPtr()
			_ = s.Sort()
			_ = s.Reverse()
			_ = s.Hashset()
			_ = s.EachItemSplitBy(",")
			_ = s.TranspileJoin(func(ss string) string { return ss }, ",")
			_ = s.PrependJoin(",", "z")
			_ = s.AppendJoin(",", "z")
			_ = s.ConcatNew("d")
			_ = s.ConcatNewStrings("d")
			_ = s.ConcatNewSimpleSlices(corestr.New.SimpleSlice.Lines("e"))
			s.PrependAppend([]string{"pre"}, []string{"post"})
			_ = s.JsonModel()
			_ = s.JsonModelAny()
			_, _ = s.MarshalJSON()
			_, _ = s.Serialize()
			_ = s.SafeStrings()
			_ = s.AsJsoner()
			_ = s.AsJsonContractsBinder()
			_ = s.AsJsonParseSelfInjector()
			_ = s.AsJsonMarshaller()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_SimpleSlice_IsEqual_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleSlice_IsEqual_Verification", func() {
		// Arrange
		tc := srcC09SimpleSliceIsEqualTestCase
		s1 := corestr.New.SimpleSlice.Lines("a", "b")
		s2 := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{
			"isEqual":               s1.IsEqual(s2),
			"isEqualLines":          s1.IsEqualLines([]string{"a", "b"}),
			"isEqualUnordered":      s1.IsEqualUnorderedLines([]string{"b", "a"}),
			"isEqualUnorderedClone": s1.IsEqualUnorderedLinesClone([]string{"b", "a"}),
			"isDistinctEqual":       s1.IsDistinctEqual(s2),
			"isDistinctEqualRaw":    s1.IsDistinctEqualRaw("a", "b"),
			"isUnorderedEqual":      s1.IsUnorderedEqual(true, s2),
			"isUnorderedEqualRaw":   s1.IsUnorderedEqualRaw(true, "b", "a"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_SimpleSlice_DistinctDiff_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleSlice_DistinctDiff_Verification", func() {
		// Arrange
		tc := srcC09SimpleSliceDistinctDiffTestCase
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		noPanic := !callPanicsSrcC09(func() {
			_ = s.DistinctDiffRaw("b", "c")
			_ = s.DistinctDiff(corestr.New.SimpleSlice.Lines("b", "c"))
			_, _ = s.AddedRemovedLinesDiff("b", "c")
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC09_SimpleSlice_RemoveIndexes_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleSlice_RemoveIndexes_Verification", func() {
		// Arrange
		tc := srcC09SimpleSliceRemoveIndexesTestCase
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		empty := corestr.New.SimpleSlice.Empty()

		// Act
		newS, err := s.RemoveIndexes(1)
		_, err2 := empty.RemoveIndexes(0)
		actual := args.Map{
			"newLen":   newS.Length(),
			"noErr":    err == nil,
			"emptyErr": err2 != nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}
func Test_SrcC09_SimpleSlice_ClearDispose_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC09_SimpleSlice_ClearDispose_Verification", func() {
		// Arrange
		tc := srcC09SimpleSliceClearDisposeTestCase
		var nilS *corestr.SimpleSlice

		// Act
		noPanic := !callPanicsSrcC09(func() {
			s := corestr.New.SimpleSlice.Lines("a")
			s.Clear()
			s.Dispose()
		})
		actual := args.Map{
			"noPanic":  noPanic,
			"nilClear": nilS.Clear() == nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC09(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}
