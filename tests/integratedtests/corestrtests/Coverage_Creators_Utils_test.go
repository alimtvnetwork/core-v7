package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// Creators + Utils — Segment 21
// ══════════════════════════════════════════════════════════════════════════════

// --- emptyCreator ---

func Test_CovEmpty_01_AllCreators(t *testing.T) {
	safeTest(t, "Test_CovEmpty_01_AllCreators", func() {
		_ = corestr.Empty.Collection()
		_ = corestr.Empty.LinkedList()
		_ = corestr.Empty.SimpleSlice()
		_ = corestr.Empty.KeyAnyValuePair()
		_ = corestr.Empty.KeyValuePair()
		_ = corestr.Empty.KeyValueCollection()
		_ = corestr.Empty.LinkedCollections()
		_ = corestr.Empty.LeftRight()
		_ = corestr.Empty.SimpleStringOnce()
		_ = corestr.Empty.SimpleStringOncePtr()
		_ = corestr.Empty.Hashset()
		_ = corestr.Empty.HashsetsCollection()
		_ = corestr.Empty.Hashmap()
		_ = corestr.Empty.CharCollectionMap()
		_ = corestr.Empty.KeyValuesCollection()
		_ = corestr.Empty.CollectionsOfCollection()
		_ = corestr.Empty.CharHashsetMap()
	})
}

// --- newSimpleStringOnceCreator ---

func Test_CovSSOCreator_01_Init_InitPtr(t *testing.T) {
	safeTest(t, "Test_CovSSOCreator_01_Init_InitPtr", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		if sso.Value() != "hello" {
			t.Fatal("expected hello")
		}
		ssoP := corestr.New.SimpleStringOnce.InitPtr("hello")
		if ssoP.Value() != "hello" {
			t.Fatal("expected hello")
		}
	})
}

func Test_CovSSOCreator_02_Uninitialized(t *testing.T) {
	safeTest(t, "Test_CovSSOCreator_02_Uninitialized", func() {
		sso := corestr.New.SimpleStringOnce.Uninitialized("hello")
		if sso.IsInitialized() {
			t.Fatal("expected uninitialized")
		}
	})
}

func Test_CovSSOCreator_03_Create_CreatePtr(t *testing.T) {
	safeTest(t, "Test_CovSSOCreator_03_Create_CreatePtr", func() {
		sso := corestr.New.SimpleStringOnce.Create("v", true)
		if sso.Value() != "v" {
			t.Fatal("expected v")
		}
		ssoP := corestr.New.SimpleStringOnce.CreatePtr("v", true)
		if ssoP.Value() != "v" {
			t.Fatal("expected v")
		}
	})
}

func Test_CovSSOCreator_04_Any(t *testing.T) {
	safeTest(t, "Test_CovSSOCreator_04_Any", func() {
		sso := corestr.New.SimpleStringOnce.Any(false, "test", true)
		if !sso.IsInitialized() {
			t.Fatal("expected initialized")
		}
	})
}

func Test_CovSSOCreator_05_Empty(t *testing.T) {
	safeTest(t, "Test_CovSSOCreator_05_Empty", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		if sso.IsInitialized() {
			t.Fatal("expected uninitialized")
		}
	})
}

// --- newKeyValuesCreator ---

func Test_CovKVCreator_01_Cap_Empty(t *testing.T) {
	safeTest(t, "Test_CovKVCreator_01_Cap_Empty", func() {
		kv := corestr.New.KeyValues.Cap(5)
		if kv.Length() != 0 {
			t.Fatal("expected 0")
		}
		kv2 := corestr.New.KeyValues.Empty()
		if kv2.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovKVCreator_02_UsingMap(t *testing.T) {
	safeTest(t, "Test_CovKVCreator_02_UsingMap", func() {
		kv := corestr.New.KeyValues.UsingMap(map[string]string{"a": "1"})
		if kv.Length() != 1 {
			t.Fatal("expected 1")
		}
		kv2 := corestr.New.KeyValues.UsingMap(map[string]string{})
		if kv2.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovKVCreator_03_UsingKeyValuePairs(t *testing.T) {
	safeTest(t, "Test_CovKVCreator_03_UsingKeyValuePairs", func() {
		kv := corestr.New.KeyValues.UsingKeyValuePairs(
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)
		if kv.Length() != 1 {
			t.Fatal("expected 1")
		}
		kv2 := corestr.New.KeyValues.UsingKeyValuePairs()
		if kv2.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovKVCreator_04_UsingKeyValueStrings(t *testing.T) {
	safeTest(t, "Test_CovKVCreator_04_UsingKeyValueStrings", func() {
		kv := corestr.New.KeyValues.UsingKeyValueStrings(
			[]string{"a", "b"}, []string{"1", "2"},
		)
		if kv.Length() != 2 {
			t.Fatal("expected 2")
		}
		kv2 := corestr.New.KeyValues.UsingKeyValueStrings([]string{}, []string{})
		if kv2.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

// --- newLinkedListCreator ---

func Test_CovLLCreator_01_Create_Empty(t *testing.T) {
	safeTest(t, "Test_CovLLCreator_01_Create_Empty", func() {
		ll := corestr.New.LinkedList.Create()
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
		ll2 := corestr.New.LinkedList.Empty()
		if ll2.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovLLCreator_02_Strings(t *testing.T) {
	safeTest(t, "Test_CovLLCreator_02_Strings", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
		ll2 := corestr.New.LinkedList.Strings([]string{})
		if ll2.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovLLCreator_03_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_CovLLCreator_03_SpreadStrings", func() {
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
		ll2 := corestr.New.LinkedList.SpreadStrings()
		if ll2.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovLLCreator_04_UsingMap(t *testing.T) {
	safeTest(t, "Test_CovLLCreator_04_UsingMap", func() {
		ll := corestr.New.LinkedList.UsingMap(map[string]bool{"a": true})
		if ll.Length() != 1 {
			t.Fatal("expected 1")
		}
		ll2 := corestr.New.LinkedList.UsingMap(nil)
		if ll2.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovLLCreator_05_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_CovLLCreator_05_PointerStringsPtr", func() {
		ll := corestr.New.LinkedList.PointerStringsPtr(nil)
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

// --- newLinkedListCollectionsCreator ---

func Test_CovLLCCreator_01_Create_Empty(t *testing.T) {
	safeTest(t, "Test_CovLLCCreator_01_Create_Empty", func() {
		lc := corestr.New.LinkedCollection.Create()
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
		lc2 := corestr.New.LinkedCollection.Empty()
		if lc2.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovLLCCreator_02_Strings(t *testing.T) {
	safeTest(t, "Test_CovLLCCreator_02_Strings", func() {
		lc := corestr.New.LinkedCollection.Strings("a", "b")
		if lc.Length() != 1 { // Strings creates one collection node containing both items
			t.Fatalf("expected 1, got %d", lc.Length())
		}
		lc2 := corestr.New.LinkedCollection.Strings()
		if lc2.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovLLCCreator_03_UsingCollections(t *testing.T) {
	safeTest(t, "Test_CovLLCCreator_03_UsingCollections", func() {
		col := corestr.New.Collection.Strings([]string{"false", "a"})
		lc := corestr.New.LinkedCollection.UsingCollections(col)
		if lc.Length() != 1 {
			t.Fatal("expected 1")
		}
		lc2 := corestr.New.LinkedCollection.UsingCollections(nil)
		_ = lc2
	})
}

func Test_CovLLCCreator_04_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_CovLLCCreator_04_PointerStringsPtr", func() {
		lc := corestr.New.LinkedCollection.PointerStringsPtr(nil)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

// --- newSimpleSliceCreator ---

func Test_CovSSCreator_01_Cap_Default_Empty(t *testing.T) {
	safeTest(t, "Test_CovSSCreator_01_Cap_Default_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
		ss2 := corestr.New.SimpleSlice.Cap(-1)
		if ss2.Length() != 0 {
			t.Fatal("expected 0")
		}
		ss3 := corestr.New.SimpleSlice.Default()
		if ss3.Length() != 0 {
			t.Fatal("expected 0")
		}
		ss4 := corestr.New.SimpleSlice.Empty()
		if ss4.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovSSCreator_02_Strings_Create_Lines(t *testing.T) {
	safeTest(t, "Test_CovSSCreator_02_Strings_Create_Lines", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		if ss.Length() != 2 {
			t.Fatal("expected 2")
		}
		ss2 := corestr.New.SimpleSlice.Create([]string{"a"})
		if ss2.Length() != 1 {
			t.Fatal("expected 1")
		}
		ss3 := corestr.New.SimpleSlice.Lines("a", "b")
		if ss3.Length() != 2 {
			t.Fatal("expected 2")
		}
		ss4 := corestr.New.SimpleSlice.SpreadStrings("a")
		if ss4.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovSSCreator_03_StringsPtr_StringsOptions_StringsClone(t *testing.T) {
	safeTest(t, "Test_CovSSCreator_03_StringsPtr_StringsOptions_StringsClone", func() {
		ss := corestr.New.SimpleSlice.StringsPtr([]string{"a"})
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
		ss2 := corestr.New.SimpleSlice.StringsPtr([]string{})
		if ss2.Length() != 0 {
			t.Fatal("expected 0")
		}
		ss3 := corestr.New.SimpleSlice.StringsOptions(true, []string{"a"})
		if ss3.Length() != 1 {
			t.Fatal("expected 1")
		}
		ss4 := corestr.New.SimpleSlice.StringsOptions(false, []string{"a"})
		if ss4.Length() != 1 {
			t.Fatal("expected 1")
		}
		ss5 := corestr.New.SimpleSlice.StringsOptions(true, []string{})
		if ss5.Length() != 0 {
			t.Fatal("expected 0")
		}
		ss6 := corestr.New.SimpleSlice.StringsClone([]string{"a"})
		if ss6.Length() != 1 {
			t.Fatal("expected 1")
		}
		ss7 := corestr.New.SimpleSlice.StringsClone(nil)
		if ss7.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovSSCreator_04_Direct_UsingLines(t *testing.T) {
	safeTest(t, "Test_CovSSCreator_04_Direct_UsingLines", func() {
		ss := corestr.New.SimpleSlice.Direct(true, []string{"a"})
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
		ss2 := corestr.New.SimpleSlice.Direct(false, []string{"a"})
		if ss2.Length() != 1 {
			t.Fatal("expected 1")
		}
		ss3 := corestr.New.SimpleSlice.Direct(true, nil)
		if ss3.Length() != 0 {
			t.Fatal("expected 0")
		}
		ss4 := corestr.New.SimpleSlice.UsingLines(true, "a", "b")
		if ss4.Length() != 2 {
			t.Fatal("expected 2")
		}
		ss5 := corestr.New.SimpleSlice.UsingLines(false, "a")
		if ss5.Length() != 1 {
			t.Fatal("expected 1")
		}
		ss6 := corestr.New.SimpleSlice.UsingLines(true)
		if ss6.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovSSCreator_05_Split_SplitLines_UsingSeparatorLine_UsingLine(t *testing.T) {
	safeTest(t, "Test_CovSSCreator_05_Split_SplitLines_UsingSeparatorLine_UsingLine", func() {
		ss := corestr.New.SimpleSlice.Split("a,b,c", ",")
		if ss.Length() != 3 {
			t.Fatal("expected 3")
		}
		ss2 := corestr.New.SimpleSlice.SplitLines("a\nb")
		if ss2.Length() != 2 {
			t.Fatal("expected 2")
		}
		ss3 := corestr.New.SimpleSlice.UsingSeparatorLine(",", "a,b")
		if ss3.Length() != 2 {
			t.Fatal("expected 2")
		}
		ss4 := corestr.New.SimpleSlice.UsingLine("a\nb")
		_ = ss4
	})
}

func Test_CovSSCreator_06_ByLen(t *testing.T) {
	safeTest(t, "Test_CovSSCreator_06_ByLen", func() {
		ss := corestr.New.SimpleSlice.ByLen([]string{"a", "b"})
		if ss.Length() != 0 {
			t.Fatal("expected 0 length, just capacity")
		}
	})
}

func Test_CovSSCreator_07_Hashset_Map(t *testing.T) {
	safeTest(t, "Test_CovSSCreator_07_Hashset_Map", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		ss := corestr.New.SimpleSlice.Hashset(hs)
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
		ssE := corestr.New.SimpleSlice.Hashset(corestr.New.Hashset.Empty())
		if ssE.Length() != 0 {
			t.Fatal("expected 0")
		}
		ssM := corestr.New.SimpleSlice.Map(map[string]string{"a": "1"})
		_ = ssM
		ssM2 := corestr.New.SimpleSlice.Map(nil)
		if ssM2.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovSSCreator_08_Deserialize(t *testing.T) {
	safeTest(t, "Test_CovSSCreator_08_Deserialize", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		data, _ := ss.Serialize()
		ss2, err := corestr.New.SimpleSlice.Deserialize(data)
		if err != nil || ss2.Length() != 2 {
			t.Fatal("unexpected error or wrong length")
		}
		_, err2 := corestr.New.SimpleSlice.Deserialize([]byte("bad"))
		if err2 == nil {
			t.Fatal("expected error")
		}
	})
}

// --- utils ---

func Test_CovUtils_01_WrapDoubleIfMissing(t *testing.T) {
	safeTest(t, "Test_CovUtils_01_WrapDoubleIfMissing", func() {
		u := corestr.StringUtils
		if u.WrapDoubleIfMissing("hello") != `"hello"` {
			t.Fatal("expected wrapped")
		}
		if u.WrapDoubleIfMissing(`"hello"`) != `"hello"` {
			t.Fatal("expected already wrapped")
		}
		if u.WrapDoubleIfMissing("") != `""` {
			t.Fatal("expected empty quotes")
		}
		if u.WrapDoubleIfMissing(`""`) != `""` {
			t.Fatal("expected empty quotes")
		}
	})
}

func Test_CovUtils_02_WrapSingleIfMissing(t *testing.T) {
	safeTest(t, "Test_CovUtils_02_WrapSingleIfMissing", func() {
		u := corestr.StringUtils
		if u.WrapSingleIfMissing("hello") != "'hello'" {
			t.Fatal("expected wrapped")
		}
		if u.WrapSingleIfMissing("'hello'") != "'hello'" {
			t.Fatal("expected already wrapped")
		}
		if u.WrapSingleIfMissing("") != "''" {
			t.Fatal("expected empty quotes")
		}
		if u.WrapSingleIfMissing("''") != "''" {
			t.Fatal("expected empty quotes")
		}
	})
}

func Test_CovUtils_03_WrapDouble_WrapSingle_WrapTilda(t *testing.T) {
	safeTest(t, "Test_CovUtils_03_WrapDouble_WrapSingle_WrapTilda", func() {
		u := corestr.StringUtils
		if u.WrapDouble("hello") != `"hello"` {
			t.Fatal("expected wrapped")
		}
		if u.WrapSingle("hello") != "'hello'" {
			t.Fatal("expected wrapped")
		}
		if u.WrapTilda("hello") != "`hello`" {
			t.Fatal("expected wrapped")
		}
	})
}
