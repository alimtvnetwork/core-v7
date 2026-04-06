package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_SrcC13_Hashmap_Diff_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_Diff_Verification", func() {
		// Arrange
		tc := srcC13HashmapDiffTestCase

		// Act
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		actual := args.Map{
			"diffRawGt0": len(hm.DiffRaw(map[string]string{"a": "2"})) > 0,
			"diffNonE":   !hm.Diff(corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "2"})).IsEmpty(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_HasAllColl_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_HasAllColl_Verification", func() {
		// Arrange
		tc := srcC13HashmapHasAllCollTestCase
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"}, corestr.KeyValuePair{Key: "b", Value: "2"})

		// Act
		actual := args.Map{
			"hasAll":   hm.HasAllCollectionItems(corestr.New.Collection.Strings([]string{"a", "b"})),
			"hasNil":   corestr.New.Hashmap.Cap(5).HasAllCollectionItems(nil),
			"hasEmpty": corestr.New.Hashmap.Cap(5).HasAllCollectionItems(corestr.New.Collection.Cap(0)),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_HasVariants_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_HasVariants_Verification", func() {
		// Arrange
		tc := srcC13HashmapHasVariantsTestCase
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"}, corestr.KeyValuePair{Key: "b", Value: "2"})

		// Act
		actual := args.Map{
			"hasAll":       hm.HasAll("a", "b"),
			"hasAllMiss":   hm.HasAll("a", "z"),
			"hasAnyItem":   hm.HasAnyItem(),
			"emptyAnyItem": corestr.New.Hashmap.Cap(0).HasAnyItem(),
			"hasAny":       corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"}).HasAny("a", "z"),
			"hasAnyMiss":   corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"}).HasAny("x", "y"),
			"hasWithLock":  corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"}).HasWithLock("a"),
			"hasWLMiss":    corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"}).HasWithLock("z"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_Filter_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_Filter_Verification", func() {
		// Arrange
		tc := srcC13HashmapFilterTestCase
		fAll := func(str string, i int) (string, bool, bool) { return str, true, false }
		fBreak := func(str string, i int) (string, bool, bool) { return str, true, true }
		fSkip := func(str string, i int) (string, bool, bool) { return str, false, false }
		hm2 := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "abc", Value: "1"}, corestr.KeyValuePair{Key: "def", Value: "2"})

		// Act
		actual := args.Map{
			"filterLen":   len(hm2.GetKeysFilteredItems(fAll)),
			"filterEmpty": len(corestr.New.Hashmap.Cap(0).GetKeysFilteredItems(fAll)),
			"filterBreak": len(corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"}, corestr.KeyValuePair{Key: "b", Value: "2"}).GetKeysFilteredItems(fBreak)),
			"filterSkip":  len(corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"}).GetKeysFilteredItems(fSkip)),
			"colNonEmpty": !corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"}).GetKeysFilteredCollection(fAll).IsEmpty(),
			"colEmpty":    corestr.New.Hashmap.Cap(0).GetKeysFilteredCollection(fAll).IsEmpty(),
			"colBreakLen": corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"}, corestr.KeyValuePair{Key: "b", Value: "2"}).GetKeysFilteredCollection(fBreak).Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_Items_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_Items_Verification", func() {
		// Arrange
		tc := srcC13HashmapItemsTestCase
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		var nilHm *corestr.Hashmap

		// Act
		actual := args.Map{
			"itemsLen":   len(hm.Items()),
			"safeLen":    len(hm.SafeItems()),
			"safeNilNil": nilHm.SafeItems() == nil,
			"copyLen":    len(*hm.ItemsCopyLock()),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_Values_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_Values_Verification", func() {
		// Arrange
		tc := srcC13HashmapValuesTestCase
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})

		// Act
		actual := args.Map{
			"colNonE":     !hm.ValuesCollection().IsEmpty(),
			"hsNonE":      !hm.ValuesHashset().IsEmpty(),
			"colLockNonE": !hm.ValuesCollectionLock().IsEmpty(),
			"hsLockNonE":  !hm.ValuesHashsetLock().IsEmpty(),
			"valsLen":     len(hm.ValuesList()),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_KeysValues_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_KeysValues_Verification", func() {
		// Arrange
		tc := srcC13HashmapKeysValuesTestCase
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})

		// Act
		keys, vals := hm.KeysValuesCollection()
		kvKeys, kvVals := hm.KeysValuesList()
		kvlKeys, kvlVals := hm.KeysValuesListLock()
		actual := args.Map{
			"kvColNonE":      !keys.IsEmpty() && !vals.IsEmpty(),
			"kvListLen":      len(kvKeys) + len(kvVals) - 1,
			"pairsLen":       len(hm.KeysValuePairs()),
			"pairsColNonNil": hm.KeysValuePairsCollection() != nil,
			"kvListLockLen":  len(kvlKeys) + len(kvlVals) - 1,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_Keys_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_Keys_Verification", func() {
		// Arrange
		tc := srcC13HashmapKeysTestCase
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})

		// Act
		actual := args.Map{
			"allKeysLen":    len(hm.AllKeys()),
			"allKeysEmpty":  len(corestr.New.Hashmap.Cap(0).AllKeys()),
			"keysLen":       len(hm.Keys()),
			"keysColNonE":   !hm.KeysCollection().IsEmpty(),
			"keysLockLen":   len(hm.KeysLock()),
			"keysLockEmpty": len(corestr.New.Hashmap.Cap(0).KeysLock()),
			"valsCopyLen":   len(hm.ValuesListCopyLock()),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}
func Test_SrcC13_Hashmap_Length_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_Length_Verification", func() {
		// Arrange
		tc := srcC13HashmapLengthTestCase
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		var nilHm *corestr.Hashmap

		// Act
		actual := args.Map{
			"length":     hm.Length(),
			"nilLength":  nilHm.Length(),
			"lockLength": hm.LengthLock(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_Equal_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_Equal_Verification", func() {
		// Arrange
		tc := srcC13HashmapEqualTestCase
		hm1 := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		hm2 := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		var nilHm1, nilHm2 *corestr.Hashmap

		// Act
		actual := args.Map{
			"equal":        hm1.IsEqual(*hm2),
			"ptrEqual":     hm1.IsEqualPtr(hm2),
			"bothNil":      nilHm1.IsEqualPtr(nilHm2),
			"oneNil":       corestr.New.Hashmap.Cap(5).IsEqualPtr(nil),
			"samePtr":      hm1.IsEqualPtr(hm1),
			"bothEmpty":    corestr.New.Hashmap.Cap(0).IsEqualPtr(corestr.New.Hashmap.Cap(0)),
			"diffLen":      hm1.IsEqualPtr(corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"}, corestr.KeyValuePair{Key: "b", Value: "2"})),
			"diffVal":      hm1.IsEqualPtr(corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "2"})),
			"ptrLockEqual": hm1.IsEqualPtrLock(hm2),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_Remove_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_Remove_Verification", func() {
		// Arrange
		tc := srcC13HashmapRemoveTestCase

		// Act
		hm1 := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		hm1.Remove("a")
		hm2 := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		hm2.RemoveWithLock("a")
		actual := args.Map{
			"removed":     !hm1.Has("a"),
			"removedLock": !hm2.Has("a"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_String_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_String_Verification", func() {
		// Arrange
		tc := srcC13HashmapStringTestCase
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		empty := corestr.New.Hashmap.Cap(0)

		// Act
		actual := args.Map{
			"strNonE":      hm.String() != "",
			"strEmptyNonE": empty.String() != "",
			"lockNonE":     hm.StringLock() != "",
			"lockEmptyNE":  empty.StringLock() != "",
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_Except_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_Except_Verification", func() {
		// Arrange
		tc := srcC13HashmapExceptTestCase
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"}, corestr.KeyValuePair{Key: "b", Value: "2"})
		hm1 := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})

		// Act
		actual := args.Map{
			"exceptHsLen":   len(hm.GetValuesExceptKeysInHashset(corestr.New.Hashset.Strings([]string{"a"}))),
			"exceptHsNil":   len(hm1.GetValuesExceptKeysInHashset(nil)),
			"exceptKeysLen": len(hm.GetValuesKeysExcept([]string{"a"})),
			"exceptKeysNil": len(hm1.GetValuesKeysExcept(nil)),
			"exceptColLen":  len(hm.GetAllExceptCollection(corestr.New.Collection.Strings([]string{"a"}))),
			"exceptColNil":  len(hm1.GetAllExceptCollection(nil)),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_JoinJson_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_JoinJson_Verification", func() {
		// Arrange
		tc := srcC13HashmapJoinJsonTestCase
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})

		// Act
		b, mErr := hm.MarshalJSON()
		hmU := &corestr.Hashmap{}
		uErr := hmU.UnmarshalJSON([]byte(`{"a":"1"}`))
		hmBad := &corestr.Hashmap{}
		badErr := hmBad.UnmarshalJSON([]byte(`invalid`))
		noPanic := !callPanicsSrcC13(func() {
			_ = hm.JsonModelAny()
		})
		actual := args.Map{
			"joinNonE":     hm.Join(",") != "",
			"joinKeysNonE": hm.JoinKeys(",") != "",
			"jsonModelLen": len(hm.JsonModel()),
			"marshalOk":    mErr == nil && len(b) > 0,
			"unmarshalLen": hmU.Length(),
			"unmarshalErr": badErr != nil,
			"jsonNoErr":    hm.Json().Error == nil,
			"jsonPtrNoErr": hm.JsonPtr().HasError() == false,
			"noPanic":      noPanic,
		}
		_ = uErr

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_Error_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_Error_Verification", func() {
		// Arrange
		tc := srcC13HashmapErrorTestCase
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})

		// Act
		actual := args.Map{
			"toErr":    hm.ToError(",") != nil,
			"toDefErr": hm.ToDefaultError() != nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_Misc_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_Misc_Verification", func() {
		// Arrange
		tc := srcC13HashmapMiscTestCase
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		var nilHm *corestr.Hashmap

		// Act
		hmC := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		hmC.Clear()
		sb, sErr := hm.Serialize()
		var target map[string]string
		dErr := hm.Deserialize(&target)
		noPanic := !callPanicsSrcC13(func() {
			_ = corestr.New.Hashmap.Cap(0).AsJsoner()
			_ = corestr.New.Hashmap.Cap(0).AsJsonContractsBinder()
			_ = corestr.New.Hashmap.Cap(0).AsJsonParseSelfInjector()
			_ = corestr.New.Hashmap.Cap(0).AsJsonMarshaller()
			nilHm.Dispose()
		})
		actual := args.Map{
			"kvLinesLen":    len(hm.KeyValStringLines()),
			"clearLen":      hmC.Length(),
			"clearNilOk":    nilHm.Clear() == nil,
			"serializeOk":   sErr == nil && len(sb) > 0,
			"deserializeOk": dErr == nil && len(target) == 1,
			"noPanic":       noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_Compiler_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_Compiler_Verification", func() {
		// Arrange
		tc := srcC13HashmapCompilerTestCase
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})

		// Act
		actual := args.Map{
			"compLen":   len(hm.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })),
			"compEmpty": len(corestr.New.Hashmap.Cap(0).ToStringsUsingCompiler(func(k, v string) string { return k })),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}
func Test_SrcC13_Hashmap_Get_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_Get_Verification", func() {
		// Arrange
		tc := srcC13HashmapGetTestCase
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})

		// Act
		v, f := hm.Get("a")
		vv, vf := hm.GetValue("a")
		actual := args.Map{
			"getVal":      v,
			"getFound":    f,
			"getValVal":   vv,
			"getValFound": vf,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_FilterVariants_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_FilterVariants_Verification", func() {
		// Arrange
		tc := srcC13HashmapFilterVariantsTestCase
		fAny := func(pair corestr.KeyAnyValuePair) (string, bool, bool) { return pair.ValueString(), true, false }
		fAnyBreak := func(pair corestr.KeyAnyValuePair) (string, bool, bool) { return pair.ValueString(), true, true }
		fAnySkip := func(pair corestr.KeyAnyValuePair) (string, bool, bool) { return "", false, false }
		fKV := func(pair corestr.KeyValuePair) (string, bool, bool) { return pair.Value, true, false }
		fKVBreak := func(pair corestr.KeyValuePair) (string, bool, bool) { return pair.Value, true, true }
		fKVSkip := func(pair corestr.KeyValuePair) (string, bool, bool) { return "", false, false }

		// Act
		mk := func() *corestr.Hashmap { return corestr.New.Hashmap.Cap(5) }
		h1 := mk(); h1.AddsOrUpdatesAnyUsingFilter(fAny, corestr.KeyAnyValuePair{Key: "a", Value: 1})
		h2 := mk(); h2.AddsOrUpdatesAnyUsingFilter(nil)
		h3 := mk(); h3.AddsOrUpdatesAnyUsingFilter(fAnyBreak, corestr.KeyAnyValuePair{Key: "a", Value: 1}, corestr.KeyAnyValuePair{Key: "b", Value: 2})
		h4 := mk(); h4.AddsOrUpdatesAnyUsingFilter(fAnySkip, corestr.KeyAnyValuePair{Key: "a", Value: 1})
		h5 := mk(); h5.AddsOrUpdatesAnyUsingFilterLock(fAny, corestr.KeyAnyValuePair{Key: "a", Value: 1})
		h6 := mk(); h6.AddsOrUpdatesAnyUsingFilterLock(nil)
		h7 := mk(); h7.AddsOrUpdatesAnyUsingFilterLock(fAnyBreak, corestr.KeyAnyValuePair{Key: "a", Value: 1}, corestr.KeyAnyValuePair{Key: "b", Value: 2})
		h8 := mk(); h8.AddsOrUpdatesUsingFilter(fKV, corestr.KeyValuePair{Key: "a", Value: "1"})
		h9 := mk(); h9.AddsOrUpdatesUsingFilter(nil)
		h10 := mk(); h10.AddsOrUpdatesUsingFilter(fKVBreak, corestr.KeyValuePair{Key: "a", Value: "1"}, corestr.KeyValuePair{Key: "b", Value: "2"})
		h11 := mk(); h11.AddsOrUpdatesUsingFilter(fKVSkip, corestr.KeyValuePair{Key: "a", Value: "1"})
		actual := args.Map{
			"anyFilter":       h1.Length(),
			"anyFilterNil":    h2.Length(),
			"anyFilterBreak":  h3.Length(),
			"anyFilterSkip":   h4.Length(),
			"anyFilterLock":   h5.Length(),
			"anyFilterLNil":   h6.Length(),
			"anyFilterLBreak": h7.Length(),
			"filter":          h8.Length(),
			"filterNil":       h9.Length(),
			"filterBreak":     h10.Length(),
			"filterSkip":      h11.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_Hashmap_WgLock_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_Hashmap_WgLock_Verification", func() {
		// Arrange
		tc := srcC13HashmapWgLockTestCase

		// Act
		hm1 := corestr.New.Hashmap.Cap(5)
		wg1 := sync.WaitGroup{}
		wg1.Add(1)
		hm1.AddOrUpdateWithWgLock("k", "v", &wg1)
		wg1.Wait()

		hm2 := corestr.New.Hashmap.Cap(5)
		wg2 := sync.WaitGroup{}
		wg2.Add(1)
		hm2.AddOrUpdateStringsPtrWgLock(&wg2, []string{"a"}, []string{"1"})
		wg2.Wait()

		hm3 := corestr.New.Hashmap.Cap(5)
		wg3 := sync.WaitGroup{}
		wg3.Add(1)
		hm3.AddOrUpdateStringsPtrWgLock(&wg3, []string{}, []string{})

		panicOnMismatch := callPanicsSrcC13(func() {
			hm4 := corestr.New.Hashmap.Cap(5)
			wg4 := sync.WaitGroup{}
			wg4.Add(1)
			hm4.AddOrUpdateStringsPtrWgLock(&wg4, []string{"a"}, []string{})
		})

		actual := args.Map{
			"wgLockHas":       hm1.Has("k"),
			"strPtrHas":       hm2.Has("a"),
			"strPtrEmpty":     hm3.Length(),
			"panicOnMismatch": panicOnMismatch,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC13_HashmapDiff_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC13_HashmapDiff_Verification", func() {
		// Arrange
		tc := srcC13HashmapDiffTypeTestCase

		// Act
		noPanic := !callPanicsSrcC13(func() {
			d := corestr.HashmapDiff(map[string]string{"a": "1"})
			_ = d.Length()
			_ = d.IsEmpty()
			_ = d.HasAnyItem()
			_ = d.LastIndex()
			_ = d.Raw()
			_ = d.MapAnyItems()
			_ = d.AllKeysSorted()
			_ = d.IsRawEqual(map[string]string{"a": "1"})
			_ = d.HasAnyChanges(map[string]string{"a": "2"})
			_ = d.DiffRaw(map[string]string{"a": "2"})
			_ = d.HashmapDiffUsingRaw(map[string]string{"a": "2"})
			_ = d.HashmapDiffUsingRaw(map[string]string{"a": "1"})
			_ = d.DiffJsonMessage(map[string]string{"a": "2"})
			_ = d.RawMapStringAnyDiff()
			_ = d.ShouldDiffMessage("test", map[string]string{"a": "2"})
			_ = d.LogShouldDiffMessage("test", map[string]string{"a": "2"})
			_ = d.ToStringsSliceOfDiffMap(map[string]string{"a": "changed"})
			_, _ = d.Serialize()
			var target map[string]string
			_ = d.Deserialize(&target)
			var nilD *corestr.HashmapDiff
			_ = nilD.Length()
			_ = nilD.Raw()
			_ = nilD.MapAnyItems()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC13(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}
