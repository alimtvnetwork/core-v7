package corestrtests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_SrcC17_Creators_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC17_Creators_Verification", func() {
		// Arrange
		tc := srcC17CreatorsTestCase

		// Act
		actual := args.Map{
			"capNN":        corestr.New.CharHashsetMap.Cap(20, 10) != nil,
			"capEmpty":     !corestr.New.CharHashsetMap.Cap(20, 10).HasItems(),
			"capItemsLen":  corestr.New.CharHashsetMap.CapItems(20, 10, "abc", "axy", "xyz").Length(),
			"stringsLen":   corestr.New.CharHashsetMap.Strings(10, []string{"abc", "xyz"}).Length(),
			"stringsNilNN": corestr.New.CharHashsetMap.Strings(10, nil) != nil,
			"emptyNN":      corestr.Empty.CharHashsetMap() != nil,
			"emptyEmpty":   !corestr.Empty.CharHashsetMap().HasItems(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC17_GetCharsGroups_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC17_GetCharsGroups_Verification", func() {
		// Arrange
		tc := srcC17GetCharsGroupsTestCase
		hsm := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		actual := args.Map{
			"groupsLen":    hsm.GetCharsGroups("abc", "axy", "xyz").Length(),
			"emptySameRef": hsm.GetCharsGroups() == hsm,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC17_Add_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC17_Add_Verification", func() {
		// Arrange
		tc := srcC17AddTestCase

		// Act
		hsm1 := corestr.New.CharHashsetMap.Cap(10, 5)
		hsm1.Add("alpha"); hsm1.Add("avocado"); hsm1.Add("alpha"); hsm1.Add("beta")
		hsm2 := corestr.New.CharHashsetMap.Cap(10, 5)
		hsm2.AddStrings("x1", "x2", "y1")
		hsm3 := corestr.New.CharHashsetMap.Cap(10, 5)
		hsm3.AddStrings()
		hsm4 := corestr.New.CharHashsetMap.Cap(10, 5)
		hsm4.AddLock("hello"); hsm4.AddLock("help")
		hsm5 := corestr.New.CharHashsetMap.Cap(10, 5)
		hsm5.AddStringsLock("abc", "axy")
		hsm6 := corestr.New.CharHashsetMap.Cap(10, 5)
		hsm6.AddStringsLock()
		hsm7 := corestr.New.CharHashsetMap.Cap(10, 5)
		hsm7.AddSameStartingCharItems('a', []string{"abc", "axy"})
		hsm7.AddSameStartingCharItems('a', []string{"azz"})
		hsm8 := corestr.New.CharHashsetMap.Cap(10, 5)
		hsm8.AddSameStartingCharItems('a', []string{})
		actual := args.Map{
			"addLen":        hsm1.Length(),
			"addSum":        hsm1.AllLengthsSum(),
			"addStrLenX":    hsm2.LengthOf('x'),
			"addStrEmptyE":  !hsm3.HasItems(),
			"addLockLen":    hsm4.LengthLock(),
			"addStrLockLen": hsm5.Length(),
			"addStrLockEE":  !hsm6.HasItems(),
			"sameLen":       hsm7.LengthOfHashsetFromFirstChar("a") - 1,
			"sameMoreLen":   hsm7.LengthOfHashsetFromFirstChar("a"),
			"sameEmptyE":    !hsm8.HasItems(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC17_AddColl_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC17_AddColl_Verification", func() {
		// Arrange
		tc := srcC17AddCollTestCase

		// Act
		hsm1 := corestr.New.CharHashsetMap.Cap(10, 5)
		hsm1.AddCollectionItems(corestr.New.Collection.Strings([]string{"abc", "xyz"}))
		hsm2 := corestr.New.CharHashsetMap.Cap(10, 5)
		hsm2.AddCollectionItems(nil)
		hsm3 := corestr.New.CharHashsetMap.Cap(10, 5)
		hsm3.AddCharCollectionMapItems(corestr.New.CharCollectionMap.Items([]string{"abc", "xyz"}))
		hsm4 := corestr.New.CharHashsetMap.Cap(10, 5)
		hsm4.AddCharCollectionMapItems(nil)
		// async
		hsmA := corestr.New.CharHashsetMap.Cap(10, 5)
		done := make(chan bool, 1)
		hsmA.AddCollectionItemsAsyncLock(corestr.New.Collection.Strings([]string{"abc", "xyz"}), func(chm *corestr.CharHashsetMap) {
			done <- true
		})
		asyncDone := false
		select {
		case <-done:
			asyncDone = true
		case <-time.After(2 * time.Second):
		}
		hsmAN := corestr.New.CharHashsetMap.Cap(10, 5)
		hsmAN.AddCollectionItemsAsyncLock(nil, nil)
		actual := args.Map{
			"colLen":    hsm1.Length(),
			"colNilE":   !hsm2.HasItems(),
			"ccmLen":    hsm3.AllLengthsSum(),
			"ccmNilE":   !hsm4.HasItems(),
			"asyncDone": asyncDone,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC17_Has_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC17_Has_Verification", func() {
		// Arrange
		tc := srcC17HasTestCase
		hsm := corestr.New.CharHashsetMap.CapItems(10, 5, "foo", "far", "bar")
		hsmE := corestr.Empty.CharHashsetMap()

		// Act
		hwHas, hwHs := corestr.New.CharHashsetMap.CapItems(10, 5, "foo", "far").HasWithHashset("foo")
		hwMiss, _ := corestr.New.CharHashsetMap.CapItems(10, 5, "foo", "far").HasWithHashset("zzz")
		hwEH, hwEHs := hsmE.HasWithHashset("x")
		hwlH, hwlHs := corestr.New.CharHashsetMap.CapItems(10, 5, "foo").HasWithHashsetLock("foo")
		hwlM, _ := corestr.New.CharHashsetMap.CapItems(10, 5, "foo").HasWithHashsetLock("zzz")
		hwlE, _ := hsmE.HasWithHashsetLock("x")
		actual := args.Map{
			"hasFoo":      hsm.Has("foo"),
			"hasBaz":      hsm.Has("baz"),
			"hasZzz":      hsm.Has("zzz"),
			"hasEmptyX":   hsmE.Has("x"),
			"hwHasFoo":    hwHas,
			"hwHsNonE":    !hwHs.IsEmpty(),
			"hwMissHas":   hwMiss,
			"hwEmptyHas":  hwEH,
			"hwEmptyHsNN": hwEHs != nil,
			"hwlHasFoo":   hwlH,
			"hwlHsNN":     hwlHs != nil,
			"hwlMissHas":  hwlM,
			"hwlEmptyHas": hwlE,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC17_Length_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC17_Length_Verification", func() {
		// Arrange
		tc := srcC17LengthTestCase

		// Act
		actual := args.Map{
			"lenOfA":      corestr.New.CharHashsetMap.CapItems(10, 5, "abc", "axy").LengthOf('a'),
			"lenOfZ":      corestr.New.CharHashsetMap.CapItems(10, 5, "abc", "axy").LengthOf('z'),
			"lenOfEmptyA": corestr.Empty.CharHashsetMap().LengthOf('a'),
			"lockA":       corestr.New.CharHashsetMap.CapItems(10, 5, "abc").LengthOfLock('a'),
			"lockZ":       corestr.New.CharHashsetMap.CapItems(10, 5, "abc").LengthOfLock('z'),
			"lockEmptyA":  corestr.Empty.CharHashsetMap().LengthOfLock('a'),
			"allSum":      corestr.New.CharHashsetMap.CapItems(10, 5, "a1", "a2", "b1").AllLengthsSum(),
			"allSumEmpty": corestr.Empty.CharHashsetMap().AllLengthsSum(),
			"allSumLock":  corestr.New.CharHashsetMap.CapItems(10, 5, "a1", "b1").AllLengthsSumLock(),
			"allSumLockE": corestr.Empty.CharHashsetMap().AllLengthsSumLock(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC17_State_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC17_State_Verification", func() {
		// Arrange
		tc := srcC17StateTestCase
		hsmE := corestr.Empty.CharHashsetMap()

		// Act
		hsmA := corestr.Empty.CharHashsetMap()
		hsmA.Add("x")
		actual := args.Map{
			"emptyIsEmpty": hsmE.IsEmpty(),
			"emptyHasIt":   hsmE.HasItems(),
			"addedIsEmpty": hsmA.IsEmpty(),
			"addedHasIt":   hsmA.HasItems(),
			"emptyLock":    corestr.Empty.CharHashsetMap().IsEmptyLock(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC17_Equals_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC17_Equals_Verification", func() {
		// Arrange
		tc := srcC17EqualsTestCase

		// Act
		actual := args.Map{
			"equalSame":      corestr.New.CharHashsetMap.CapItems(10, 5, "abc", "xyz").IsEquals(corestr.New.CharHashsetMap.CapItems(10, 5, "abc", "xyz")),
			"equalNil":       corestr.New.CharHashsetMap.CapItems(10, 5, "abc").IsEquals(nil),
			"equalSameRef":   func() bool { h := corestr.New.CharHashsetMap.CapItems(10, 5, "abc"); return h.IsEquals(h) }(),
			"equalBothEmpty": corestr.Empty.CharHashsetMap().IsEquals(corestr.Empty.CharHashsetMap()),
			"equalOneEmpty":  corestr.New.CharHashsetMap.CapItems(10, 5, "abc").IsEquals(corestr.Empty.CharHashsetMap()),
			"equalDiffLen":   corestr.New.CharHashsetMap.CapItems(10, 5, "abc", "xyz").IsEquals(corestr.New.CharHashsetMap.CapItems(10, 5, "abc")),
			"equalDiffCont":  corestr.New.CharHashsetMap.CapItems(10, 5, "abc").IsEquals(corestr.New.CharHashsetMap.CapItems(10, 5, "axy")),
			"equalMissKey":   corestr.New.CharHashsetMap.CapItems(10, 5, "abc").IsEquals(corestr.New.CharHashsetMap.CapItems(10, 5, "xyz")),
			"equalLock":      corestr.New.CharHashsetMap.CapItems(10, 5, "abc").IsEqualsLock(corestr.New.CharHashsetMap.CapItems(10, 5, "abc")),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC17_GetHashset_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC17_GetHashset_Verification", func() {
		// Arrange
		tc := srcC17GetHashsetTestCase
		hsm := corestr.New.CharHashsetMap.CapItems(10, 5, "abc")

		// Act
		actual := args.Map{
			"getNN":          hsm.GetHashset("a", false) != nil,
			"getMissNil":     hsm.GetHashset("z", false) == nil,
			"getCreateNN":    hsm.GetHashset("z", true) != nil,
			"getLockNN":      corestr.New.CharHashsetMap.CapItems(10, 5, "abc").GetHashsetLock(false, "a") != nil,
			"getByCharNN":    corestr.New.CharHashsetMap.CapItems(10, 5, "abc").GetHashsetByChar('a') != nil,
			"hsByCharNN":     corestr.New.CharHashsetMap.CapItems(10, 5, "abc").HashsetByChar('a') != nil,
			"hsByCharLockNN": corestr.New.CharHashsetMap.CapItems(10, 5, "abc").HashsetByCharLock('a') != nil,
			"hsByCharLockZE": corestr.New.CharHashsetMap.CapItems(10, 5, "abc").HashsetByCharLock('z').IsEmpty(),
			"hsByStrNN":      corestr.New.CharHashsetMap.CapItems(10, 5, "abc").HashsetByStringFirstChar("abc") != nil,
			"hsByStrLockNN":  corestr.New.CharHashsetMap.CapItems(10, 5, "abc").HashsetByStringFirstCharLock("abc") != nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC17_AddSameChars_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC17_AddSameChars_Verification", func() {
		// Arrange
		tc := srcC17AddSameCharsTestCase

		// Act
		// Collection variants
		hsm1 := corestr.New.CharHashsetMap.Cap(10, 5)
		r1 := hsm1.AddSameCharsCollection("a", corestr.New.Collection.Strings([]string{"abc", "axy"}))
		r2 := hsm1.AddSameCharsCollection("a", corestr.New.Collection.Strings([]string{"azz"}))
		// Hashset variants
		hsm2 := corestr.New.CharHashsetMap.Cap(10, 5)
		rh1 := hsm2.AddSameCharsHashset("a", corestr.New.Hashset.Strings([]string{"abc", "axy"}))
		// Lock variants
		actual := args.Map{
			"colLen":         r1.Length(),
			"colMoreLen":     r2.Length(),
			"colNilNN":       corestr.New.CharHashsetMap.Cap(10, 5).AddSameCharsCollection("a", nil) != nil,
			"colExNilNN":     corestr.New.CharHashsetMap.CapItems(10, 5, "abc").AddSameCharsCollection("a", nil) != nil,
			"hsLen":          rh1.Length(),
			"hsNilNN":        corestr.New.CharHashsetMap.Cap(10, 5).AddSameCharsHashset("a", nil) != nil,
			"hsExNilNN":      corestr.New.CharHashsetMap.CapItems(10, 5, "abc").AddSameCharsHashset("a", nil) != nil,
			"hsAddExLen":     corestr.New.CharHashsetMap.CapItems(10, 5, "abc").AddSameCharsHashset("a", corestr.New.Hashset.Strings([]string{"axy"})).Length(),
			"colLockNN":      corestr.New.CharHashsetMap.Cap(10, 5).AddSameCharsCollectionLock("a", corestr.New.Collection.Strings([]string{"abc"})) != nil,
			"colLockNilNN":   corestr.New.CharHashsetMap.Cap(10, 5).AddSameCharsCollectionLock("a", nil) != nil,
			"colLockExNilNN": corestr.New.CharHashsetMap.CapItems(10, 5, "abc").AddSameCharsCollectionLock("a", nil) != nil,
			"colLockExAddNN": corestr.New.CharHashsetMap.CapItems(10, 5, "abc").AddSameCharsCollectionLock("a", corestr.New.Collection.Strings([]string{"axy"})) != nil,
			"hsLockNN":       corestr.New.CharHashsetMap.Cap(10, 5).AddHashsetLock("a", corestr.New.Hashset.Strings([]string{"abc"})) != nil,
			"hsLockNilNN":    corestr.New.CharHashsetMap.Cap(10, 5).AddHashsetLock("a", nil) != nil,
			"hsLockExNilNN":  corestr.New.CharHashsetMap.CapItems(10, 5, "abc").AddHashsetLock("a", nil) != nil,
			"hsLockExAddNN":  corestr.New.CharHashsetMap.CapItems(10, 5, "abc").AddHashsetLock("a", corestr.New.Hashset.Strings([]string{"axy"})) != nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC17_AddHashsetItems_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC17_AddHashsetItems_Verification", func() {
		// Arrange
		tc := srcC17AddHashsetItemsTestCase

		// Act
		hsm1 := corestr.New.CharHashsetMap.Cap(10, 5)
		hsm1.AddHashsetItems(corestr.New.Hashset.Strings([]string{"abc", "xyz"}))
		hsm2 := corestr.New.CharHashsetMap.Cap(10, 5)
		hsm2.AddHashsetItems(corestr.New.Hashset.Empty())
		hsm3 := corestr.New.CharHashsetMap.Cap(10, 5)
		hsm3.AddHashsetItemsAsyncLock(nil, nil)
		actual := args.Map{
			"itemsSum":   hsm1.AllLengthsSum(),
			"itemsEmE":   !hsm2.HasItems(),
			"asyncNilOk": true,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC17_HashsetsColl_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC17_HashsetsColl_Verification", func() {
		// Arrange
		tc := srcC17HashsetsCollTestCase

		// Act
		actual := args.Map{
			"hscNN":       corestr.New.CharHashsetMap.CapItems(10, 5, "abc", "xyz").HashsetsCollection() != nil,
			"hscEmptyNN":  corestr.Empty.CharHashsetMap().HashsetsCollection() != nil,
			"hscCharsNN":  corestr.New.CharHashsetMap.CapItems(10, 5, "abc", "xyz").HashsetsCollectionByChars('a', 'x') != nil,
			"hscCharsENN": corestr.Empty.CharHashsetMap().HashsetsCollectionByChars('a') != nil,
			"hscStrNN":    corestr.New.CharHashsetMap.CapItems(10, 5, "abc", "xyz").HashsetsCollectionByStringsFirstChar("abc", "xyz") != nil,
			"hscStrENN":   corestr.Empty.CharHashsetMap().HashsetsCollectionByStringsFirstChar("abc") != nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC17_ListSort_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC17_ListSort_Verification", func() {
		// Arrange
		tc := srcC17ListSortTestCase

		// Act
		asc := corestr.New.CharHashsetMap.CapItems(10, 5, "cherry", "apple", "banana").SortedListAsc()
		dsc := corestr.New.CharHashsetMap.CapItems(10, 5, "cherry", "apple", "banana").SortedListDsc()
		actual := args.Map{
			"listLen":  len(corestr.New.CharHashsetMap.CapItems(10, 5, "abc", "xyz").List()),
			"ascFirst": asc[0],
			"ascLen":   len(asc),
			"dscFirst": dsc[0],
			"dscLen":   len(dsc),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC17_MapStringPrint_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC17_MapStringPrint_Verification", func() {
		// Arrange
		tc := srcC17MapStringPrintTestCase
		hsm := corestr.New.CharHashsetMap.CapItems(10, 5, "abc")

		// Act
		noPanic := !callPanicsSrcC17(func() {
			_ = hsm.GetMap()
			_ = hsm.GetCopyMapLock()
			_ = corestr.Empty.CharHashsetMap().GetCopyMapLock()
			_ = hsm.String()
			_ = hsm.SummaryString()
			_ = hsm.StringLock()
			_ = hsm.SummaryStringLock()
			hsm.Print(false)
			hsm.PrintLock(false)
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC17_Json_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC17_Json_Verification", func() {
		// Arrange
		tc := srcC17JsonTestCase
		hsm := corestr.New.CharHashsetMap.CapItems(10, 5, "abc", "xyz")

		// Act
		noPanic := !callPanicsSrcC17(func() {
			_ = hsm.JsonModel()
			_ = hsm.JsonModelAny()
			data, _ := json.Marshal(hsm)
			hsm2 := corestr.New.CharHashsetMap.Cap(10, 5)
			_ = json.Unmarshal(data, hsm2)
			r := hsm.Json()
			_ = r.Error == nil
			jr := hsm.JsonPtr()
			hsm3 := corestr.New.CharHashsetMap.Cap(10, 5)
			_, _ = hsm3.ParseInjectUsingJson(jr)
			hsm4 := corestr.New.CharHashsetMap.Cap(10, 5)
			_ = hsm4.ParseInjectUsingJsonMust(jr)
			hsm5 := corestr.New.CharHashsetMap.Cap(10, 5)
			_ = hsm5.JsonParseSelfInject(jr)
			_ = hsm.AsJsonContractsBinder()
			_ = hsm.AsJsoner()
			_ = hsm.AsJsonMarshaller()
			_ = hsm.AsJsonParseSelfInjector()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC17_ClearDataModel_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC17_ClearDataModel_Verification", func() {
		// Arrange
		tc := srcC17ClearDataModelTestCase

		// Act
		hsmR := corestr.New.CharHashsetMap.CapItems(10, 5, "abc", "xyz")
		hsmR.RemoveAll()
		hsmC := corestr.New.CharHashsetMap.CapItems(10, 5, "abc")
		hsmC.Clear()
		hsmCE := corestr.Empty.CharHashsetMap()
		hsmCE.Clear()
		hsm := corestr.New.CharHashsetMap.CapItems(10, 5, "abc")
		dm := corestr.NewCharHashsetMapDataModelUsing(hsm)
		actual := args.Map{
			"removeAllE":   !hsmR.HasItems(),
			"clearE":       !hsmC.HasItems(),
			"clearEmptyOk": true,
			"dataModelNN":  corestr.NewCharHashsetMapUsingDataModel(dm) != nil,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC17(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}
