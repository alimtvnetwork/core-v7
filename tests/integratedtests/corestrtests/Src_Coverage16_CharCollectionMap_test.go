package corestrtests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_SrcC16_Ccm_Creators_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC16_Ccm_Creators_Verification", func() {
		// Arrange
		tc := srcC16CcmCreatorsTestCase

		// Act
		actual := args.Map{
			"emptyIsEmpty":    !corestr.Empty.CharCollectionMap().HasItems(),
			"capNonNil":       corestr.New.CharCollectionMap.CapSelfCap(20, 5) != nil,
			"itemsLen":        corestr.New.CharCollectionMap.Items([]string{"apple", "avocado", "banana"}).Length(),
			"itemsEmptyEmpty": corestr.New.CharCollectionMap.Items([]string{}).IsEmpty(),
			"ptrCapLen":       corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 3, []string{"cat", "car", "dog"}).Length(),
			"ptrCapEmptyE":    !corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 3, []string{}).HasItems(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC16_Ccm_GetCharsGroups_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC16_Ccm_GetCharsGroups_Verification", func() {
		// Arrange
		tc := srcC16CcmGetCharsGroupsTestCase
		cm := corestr.New.CharCollectionMap.CapSelfCap(10, 5)

		// Act
		actual := args.Map{
			"groupsLen":   cm.GetCharsGroups([]string{"abc", "axy", "bcd"}).Length(),
			"groupsEmpty": cm.GetCharsGroups([]string{}) == cm,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC16_Ccm_Add_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC16_Ccm_Add_Verification", func() {
		// Arrange
		tc := srcC16CcmAddTestCase

		// Act
		cm1 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		cm1.Add("alpha"); cm1.Add("avocado"); cm1.Add("beta")
		cm2 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		cm2.AddStrings("x1", "x2", "y1")
		cm3 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		cm3.AddStrings()
		cm4 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		cm4.AddLock("hello"); cm4.AddLock("help")
		cm5 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		cm5.AddSameStartingCharItems('a', []string{"abc", "axy"}, false)
		cm5.AddSameStartingCharItems('a', []string{"azz"}, false)
		cm6 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		cm6.AddSameStartingCharItems('a', []string{}, false)
		actual := args.Map{
			"addLen":      cm1.Length(),
			"addSum":      cm1.AllLengthsSum(),
			"strLenX":     cm2.LengthOf('x'),
			"strEmptyE":   !cm3.HasItems(),
			"lockLen":     cm4.LengthLock(),
			"sameLen":     cm5.LengthOfCollectionFromFirstChar("a") - 1,
			"sameLenMore": cm5.LengthOfCollectionFromFirstChar("a"),
			"sameEmptyE":  !cm6.HasItems(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC16_Ccm_Has_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC16_Ccm_Has_Verification", func() {
		// Arrange
		tc := srcC16CcmHasTestCase
		cm := corestr.New.CharCollectionMap.Items([]string{"foo", "far", "bar"})
		cmE := corestr.Empty.CharCollectionMap()

		// Act
		hwcHas, hwcCol := corestr.New.CharCollectionMap.Items([]string{"foo", "far"}).HasWithCollection("foo")
		hwcMiss, _ := corestr.New.CharCollectionMap.Items([]string{"foo"}).HasWithCollection("zzz")
		hwcEH, _ := cmE.HasWithCollection("foo")
		hwcLH, _ := corestr.New.CharCollectionMap.Items([]string{"foo", "far"}).HasWithCollectionLock("foo")
		hwcLM, _ := corestr.New.CharCollectionMap.Items([]string{"foo"}).HasWithCollectionLock("zzz")
		hwcLE, _ := cmE.HasWithCollectionLock("x")
		actual := args.Map{
			"hasFoo":           cm.Has("foo"),
			"hasBaz":           cm.Has("baz"),
			"hasZzz":           cm.Has("zzz"),
			"hasEmptyAnything": cmE.Has("anything"),
			"hwcHas":           hwcHas,
			"hwcColNonE":       !hwcCol.IsEmpty(),
			"hwcMissHas":       hwcMiss,
			"hwcEmptyHas":      hwcEH,
			"hwcLockHas":       hwcLH,
			"hwcLockMiss":      hwcLM,
			"hwcLockEmpty":     hwcLE,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC16_Ccm_Length_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC16_Ccm_Length_Verification", func() {
		// Arrange
		tc := srcC16CcmLengthTestCase

		// Act
		cm2 := corestr.New.CharCollectionMap.Items([]string{"abc", "axy"})
		cmE := corestr.Empty.CharCollectionMap()
		cm1 := corestr.New.CharCollectionMap.Items([]string{"abc"})
		cm3 := corestr.New.CharCollectionMap.Items([]string{"a1", "a2", "b1"})
		cm4 := corestr.New.CharCollectionMap.Items([]string{"a1", "b1"})
		actual := args.Map{
			"lenOfA":     cm2.LengthOf('a'),
			"lenOfZ":     cm2.LengthOf('z'),
			"lenOfEmpty": cmE.LengthOf('a'),
			"lenOfLockA": cm1.LengthOfLock('a'),
			"lenOfLockZ": cm1.LengthOfLock('z'),
			"lenOfLockE": cmE.LengthOfLock('a'),
			"colFromA":   cm2.LengthOfCollectionFromFirstChar("a"),
			"colFromZ":   cm2.LengthOfCollectionFromFirstChar("z"),
			"allSum":     cm3.AllLengthsSum(),
			"allSumLock": cm4.AllLengthsSumLock(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC16_Ccm_State_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC16_Ccm_State_Verification", func() {
		// Arrange
		tc := srcC16CcmStateTestCase
		cmE := corestr.Empty.CharCollectionMap()

		// Act
		cmA := corestr.Empty.CharCollectionMap()
		cmA.Add("x")
		actual := args.Map{
			"emptyIsEmpty":  cmE.IsEmpty(),
			"emptyHasItems": cmE.HasItems(),
			"addedIsEmpty":  cmA.IsEmpty(),
			"addedHasItems": cmA.HasItems(),
			"emptyLock":     corestr.Empty.CharCollectionMap().IsEmptyLock(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC16_Ccm_Equals_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC16_Ccm_Equals_Verification", func() {
		// Arrange
		tc := srcC16CcmEqualsTestCase
		cm1 := corestr.New.CharCollectionMap.Items([]string{"abc", "xyz"})
		cm2 := corestr.New.CharCollectionMap.Items([]string{"abc", "xyz"})
		cmA := corestr.New.CharCollectionMap.Items([]string{"abc"})

		// Act
		actual := args.Map{
			"equalSame":       cm1.IsEquals(cm2),
			"equalNil":        cmA.IsEquals(nil),
			"equalSameRef":    cmA.IsEqualsCaseSensitive(true, cmA),
			"equalBothEmpty":  corestr.Empty.CharCollectionMap().IsEquals(corestr.Empty.CharCollectionMap()),
			"equalOneEmpty":   cmA.IsEquals(corestr.Empty.CharCollectionMap()),
			"equalDiffLen":    cm1.IsEquals(cmA),
			"equalDiffCont":   corestr.New.CharCollectionMap.Items([]string{"abc"}).IsEquals(corestr.New.CharCollectionMap.Items([]string{"axy"})),
			"equalLock":       corestr.New.CharCollectionMap.Items([]string{"abc"}).IsEqualsLock(corestr.New.CharCollectionMap.Items([]string{"abc"})),
			"equalCSLock":     corestr.New.CharCollectionMap.Items([]string{"abc"}).IsEqualsCaseSensitiveLock(true, corestr.New.CharCollectionMap.Items([]string{"abc"})),
			"equalMissingKey": corestr.New.CharCollectionMap.Items([]string{"abc"}).IsEqualsCaseSensitive(true, corestr.New.CharCollectionMap.Items([]string{"xyz"})),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC16_Ccm_GetColl_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC16_Ccm_GetColl_Verification", func() {
		// Arrange
		tc := srcC16CcmGetCollTestCase
		cm := corestr.New.CharCollectionMap.Items([]string{"abc", "axy"})

		// Act
		actual := args.Map{
			"getLen":      cm.GetCollection("a", false).Length(),
			"getMissNil":  cm.GetCollection("z", false) == nil,
			"getCreateNN": cm.GetCollection("z", true) != nil,
			"getLockNN":   corestr.New.CharCollectionMap.Items([]string{"abc"}).GetCollectionLock("a", false) != nil,
			"getByCharNN": corestr.New.CharCollectionMap.Items([]string{"abc"}).GetCollectionByChar('a') != nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC16_Ccm_AddSameCharsCol_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC16_Ccm_AddSameCharsCol_Verification", func() {
		// Arrange
		tc := srcC16CcmAddSameCharsColTestCase
		cm := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		col := corestr.New.Collection.Strings([]string{"abc", "axy"})

		// Act
		r := cm.AddSameCharsCollection("a", col)
		r2 := cm.AddSameCharsCollection("a", corestr.New.Collection.Strings([]string{"azz"}))
		actual := args.Map{
			"addLen":         r.Length(),
			"addMoreLen":     r2.Length(),
			"addNilNN":       corestr.New.CharCollectionMap.CapSelfCap(10, 5).AddSameCharsCollection("a", nil) != nil,
			"addExistNilNN":  corestr.New.CharCollectionMap.Items([]string{"abc"}).AddSameCharsCollection("a", nil) != nil,
			"lockNN":         corestr.New.CharCollectionMap.CapSelfCap(10, 5).AddSameCharsCollectionLock("a", corestr.New.Collection.Strings([]string{"abc"})) != nil,
			"lockNilNN":      corestr.New.CharCollectionMap.CapSelfCap(10, 5).AddSameCharsCollectionLock("a", nil) != nil,
			"lockExistNilNN": corestr.New.CharCollectionMap.Items([]string{"abc"}).AddSameCharsCollectionLock("a", nil) != nil,
			"lockAddExistNN": corestr.New.CharCollectionMap.Items([]string{"abc"}).AddSameCharsCollectionLock("a", corestr.New.Collection.Strings([]string{"axy"})) != nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC16_Ccm_AddColItemsHm_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC16_Ccm_AddColItemsHm_Verification", func() {
		// Arrange
		tc := srcC16CcmAddColItemsHmTestCase

		// Act
		cm1 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		cm1.AddCollectionItems(corestr.New.Collection.Strings([]string{"abc", "xyz"}))
		cm2 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		cm2.AddCollectionItems(nil)
		cm3 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		cm3.AddHashmapsValues(corestr.New.Hashmap.KeyValuesStrings([]string{"k1", "k2"}, []string{"alpha", "beta"}))
		cm4 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		cm4.AddHashmapsValues(nil)
		cm5 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		cm5.AddHashmapsKeysValuesBoth(corestr.New.Hashmap.KeyValuesStrings([]string{"k1"}, []string{"v1"}))
		cm6 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		cm6.AddHashmapsKeysValuesBoth(nil)
		cm7 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		cm7.AddHashmapsKeysOrValuesBothUsingFilter(func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Key, true, false
		}, corestr.New.Hashmap.KeyValuesStrings([]string{"k1", "k2"}, []string{"v1", "v2"}))
		cm8 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		cm8.AddHashmapsKeysOrValuesBothUsingFilter(nil, nil)
		cm9 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		cm9.AddHashmapsKeysOrValuesBothUsingFilter(func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Key, true, true
		}, corestr.New.Hashmap.KeyValuesStrings([]string{"k1", "k2"}, []string{"v1", "v2"}))
		cm10 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		cm10.AddCharHashsetMap(corestr.New.CharHashsetMap.CapItems(10, 5, "abc", "axy"))
		actual := args.Map{
			"colItemsLen":  cm1.Length(),
			"colItemsNilE": !cm2.HasItems(),
			"hmValsHasA":   cm3.Has("alpha"),
			"hmValsNilE":   !cm4.HasItems(),
			"hmBothHasK":   cm5.Has("k1"),
			"hmBothNilE":   !cm6.HasItems(),
			"hmFilterHasK": cm7.Has("k1"),
			"hmFilterNilE": !cm8.HasItems(),
			"hmFilterBreak": cm9.AllLengthsSum(),
			"charHsLen":    cm10.AllLengthsSum(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC16_Ccm_ResizeList_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC16_Ccm_ResizeList_Verification", func() {
		// Arrange
		tc := srcC16CcmResizeListTestCase

		// Act
		cmR := corestr.New.CharCollectionMap.Items([]string{"a1"}); cmR.Resize(100)
		cmNS := corestr.New.CharCollectionMap.Items([]string{"a1", "b1"}); cmNS.Resize(1)
		cmAL := corestr.New.CharCollectionMap.Items([]string{"a1"}); cmAL.AddLength(10, 20)
		cmALE := corestr.New.CharCollectionMap.Items([]string{"a1"}); cmALE.AddLength()
		actual := args.Map{
			"resizeHas":      cmR.Has("a1"),
			"resizeNoShrink": cmNS.Length(),
			"addLenHas":      cmAL.Has("a1"),
			"addLenEmptyOk":  cmALE.Has("a1"),
			"listLen":        len(corestr.New.CharCollectionMap.Items([]string{"abc", "xyz"}).List()),
			"listEmptyLen":   len(corestr.Empty.CharCollectionMap().List()),
			"listLockLen":    len(corestr.New.CharCollectionMap.Items([]string{"abc"}).ListLock()),
			"sortedFirst":    corestr.New.CharCollectionMap.Items([]string{"cherry", "apple", "banana"}).SortedListAsc()[0],
			"sortedEmptyLen": len(corestr.Empty.CharCollectionMap().SortedListAsc()),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC16_Ccm_MapString_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC16_Ccm_MapString_Verification", func() {
		// Arrange
		tc := srcC16CcmMapStringTestCase
		cm := corestr.New.CharCollectionMap.Items([]string{"abc"})

		// Act
		noPanic := !callPanicsSrcC16(func() {
			_ = cm.GetMap()
			_ = cm.GetCopyMapLock()
			_ = corestr.Empty.CharCollectionMap().GetCopyMapLock()
			_ = cm.String()
			_ = cm.SummaryString()
			_ = cm.StringLock()
			_ = cm.SummaryStringLock()
			cm.Print(false)
			cm.PrintLock(false)
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC16_Ccm_Hashset_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC16_Ccm_Hashset_Verification", func() {
		// Arrange
		tc := srcC16CcmHashsetTestCase
		cm := corestr.New.CharCollectionMap.Items([]string{"abc", "axy", "xyz"})

		// Act
		noPanic := !callPanicsSrcC16(func() {
			_ = cm.HashsetByChar('a')
			_ = cm.HashsetByChar('z')
			_ = cm.HashsetByCharLock('a')
			_ = cm.HashsetByCharLock('z')
			_ = cm.HashsetByStringFirstChar("abc")
			_ = cm.HashsetByStringFirstCharLock("abc")
			_ = cm.HashsetsCollection()
			_ = corestr.Empty.CharCollectionMap().HashsetsCollection()
			_ = cm.HashsetsCollectionByChars('a', 'x')
			_ = corestr.Empty.CharCollectionMap().HashsetsCollectionByChars('a')
			_ = cm.HashsetsCollectionByStringFirstChar("abc", "xyz")
			_ = corestr.Empty.CharCollectionMap().HashsetsCollectionByStringFirstChar("abc")
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC16_Ccm_Json_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC16_Ccm_Json_Verification", func() {
		// Arrange
		tc := srcC16CcmJsonTestCase
		cm := corestr.New.CharCollectionMap.Items([]string{"abc"})

		// Act
		noPanic := !callPanicsSrcC16(func() {
			_ = cm.JsonModel()
			_ = cm.JsonModelAny()
			data, _ := json.Marshal(cm)
			cm2 := corestr.Empty.CharCollectionMap()
			_ = json.Unmarshal(data, cm2)
			r := cm.Json()
			_ = r.Error == nil
			_ = cm.JsonPtr()
			cm3 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
			jr := cm.JsonPtr()
			_, _ = cm3.ParseInjectUsingJson(jr)
			cm4 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
			_ = cm4.ParseInjectUsingJsonMust(jr)
			cm5 := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
			_ = cm5.JsonParseSelfInject(jr)
			_ = cm.AsJsonContractsBinder()
			_ = cm.AsJsoner()
			_ = cm.AsJsonMarshaller()
			_ = cm.AsJsonParseSelfInjector()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC16_Ccm_ClearDataModel_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC16_Ccm_ClearDataModel_Verification", func() {
		// Arrange
		tc := srcC16CcmClearDataModelTestCase

		// Act
		cmC := corestr.New.CharCollectionMap.Items([]string{"abc", "xyz"})
		cmC.Clear()
		cmCE := corestr.Empty.CharCollectionMap()
		cmCE.Clear()
		cm := corestr.New.CharCollectionMap.Items([]string{"abc"})
		dm := corestr.NewCharCollectionMapDataModelUsing(cm)
		actual := args.Map{
			"clearHasItems": cmC.HasItems(),
			"clearEmptyOk":  !cmCE.HasItems(),
			"dataModelNN":   corestr.NewCharCollectionMapUsingDataModel(dm) != nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC16(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}
