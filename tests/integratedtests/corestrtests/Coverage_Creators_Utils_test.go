package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
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
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{"result": sso.Value() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		ssoP := corestr.New.SimpleStringOnce.InitPtr("hello")
		actual = args.Map{"result": ssoP.Value() != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_CovSSOCreator_02_Uninitialized(t *testing.T) {
	safeTest(t, "Test_CovSSOCreator_02_Uninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Uninitialized("hello")

		// Act
		actual := args.Map{"result": sso.IsInitialized()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected uninitialized", actual)
	})
}

func Test_CovSSOCreator_03_Create_CreatePtr(t *testing.T) {
	safeTest(t, "Test_CovSSOCreator_03_Create_CreatePtr", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Create("v", true)

		// Act
		actual := args.Map{"result": sso.Value() != "v"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected v", actual)
		ssoP := corestr.New.SimpleStringOnce.CreatePtr("v", true)
		actual = args.Map{"result": ssoP.Value() != "v"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected v", actual)
	})
}

func Test_CovSSOCreator_04_Any(t *testing.T) {
	safeTest(t, "Test_CovSSOCreator_04_Any", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Any(false, "test", true)

		// Act
		actual := args.Map{"result": sso.IsInitialized()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected initialized", actual)
	})
}

func Test_CovSSOCreator_05_Empty(t *testing.T) {
	safeTest(t, "Test_CovSSOCreator_05_Empty", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		actual := args.Map{"result": sso.IsInitialized()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected uninitialized", actual)
	})
}

// --- newKeyValuesCreator ---

func Test_CovKVCreator_01_Cap_Empty(t *testing.T) {
	safeTest(t, "Test_CovKVCreator_01_Cap_Empty", func() {
		// Arrange
		kv := corestr.New.KeyValues.Cap(5)

		// Act
		actual := args.Map{"result": kv.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		kv2 := corestr.New.KeyValues.Empty()
		actual = args.Map{"result": kv2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovKVCreator_02_UsingMap(t *testing.T) {
	safeTest(t, "Test_CovKVCreator_02_UsingMap", func() {
		// Arrange
		kv := corestr.New.KeyValues.UsingMap(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"result": kv.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		kv2 := corestr.New.KeyValues.UsingMap(map[string]string{})
		actual = args.Map{"result": kv2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovKVCreator_03_UsingKeyValuePairs(t *testing.T) {
	safeTest(t, "Test_CovKVCreator_03_UsingKeyValuePairs", func() {
		// Arrange
		kv := corestr.New.KeyValues.UsingKeyValuePairs(
			corestr.KeyValuePair{Key: "a", Value: "1"},
		)

		// Act
		actual := args.Map{"result": kv.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		kv2 := corestr.New.KeyValues.UsingKeyValuePairs()
		actual = args.Map{"result": kv2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovKVCreator_04_UsingKeyValueStrings(t *testing.T) {
	safeTest(t, "Test_CovKVCreator_04_UsingKeyValueStrings", func() {
		// Arrange
		kv := corestr.New.KeyValues.UsingKeyValueStrings(
			[]string{"a", "b"}, []string{"1", "2"},
		)

		// Act
		actual := args.Map{"result": kv.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		kv2 := corestr.New.KeyValues.UsingKeyValueStrings([]string{}, []string{})
		actual = args.Map{"result": kv2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// --- newLinkedListCreator ---

func Test_CovLLCreator_01_Create_Empty(t *testing.T) {
	safeTest(t, "Test_CovLLCreator_01_Create_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ll2 := corestr.New.LinkedList.Empty()
		actual = args.Map{"result": ll2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovLLCreator_02_Strings(t *testing.T) {
	safeTest(t, "Test_CovLLCreator_02_Strings", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		ll2 := corestr.New.LinkedList.Strings([]string{})
		actual = args.Map{"result": ll2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovLLCreator_03_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_CovLLCreator_03_SpreadStrings", func() {
		// Arrange
		ll := corestr.New.LinkedList.SpreadStrings("a", "b")

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		ll2 := corestr.New.LinkedList.SpreadStrings()
		actual = args.Map{"result": ll2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovLLCreator_04_UsingMap(t *testing.T) {
	safeTest(t, "Test_CovLLCreator_04_UsingMap", func() {
		// Arrange
		ll := corestr.New.LinkedList.UsingMap(map[string]bool{"a": true})

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		ll2 := corestr.New.LinkedList.UsingMap(nil)
		actual = args.Map{"result": ll2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovLLCreator_05_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_CovLLCreator_05_PointerStringsPtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.PointerStringsPtr(nil)

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// --- newLinkedListCollectionsCreator ---

func Test_CovLLCCreator_01_Create_Empty(t *testing.T) {
	safeTest(t, "Test_CovLLCCreator_01_Create_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		lc2 := corestr.New.LinkedCollection.Empty()
		actual = args.Map{"result": lc2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovLLCCreator_02_Strings(t *testing.T) {
	safeTest(t, "Test_CovLLCCreator_02_Strings", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Strings("a", "b")

		// Act
		actualLen := args.Map{"result": lc.Length()}

		// Assert
		expectedLen := args.Map{"result": 1}
		expectedLen.ShouldBeEqual(t, 0, "Strings creates one collection node -- two items", actualLen)
		lc2 := corestr.New.LinkedCollection.Strings()
		actual := args.Map{"result": lc2.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovLLCCreator_03_UsingCollections(t *testing.T) {
	safeTest(t, "Test_CovLLCCreator_03_UsingCollections", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"false", "a"})
		lc := corestr.New.LinkedCollection.UsingCollections(col)

		// Act
		actual := args.Map{"result": lc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		lc2 := corestr.New.LinkedCollection.UsingCollections(nil)
		_ = lc2
	})
}

func Test_CovLLCCreator_04_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_CovLLCCreator_04_PointerStringsPtr", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.PointerStringsPtr(nil)

		// Act
		actual := args.Map{"result": lc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// --- newSimpleSliceCreator ---

func Test_CovSSCreator_01_Cap_Default_Empty(t *testing.T) {
	safeTest(t, "Test_CovSSCreator_01_Cap_Default_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		actual := args.Map{"result": ss.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss2 := corestr.New.SimpleSlice.Cap(-1)
		actual = args.Map{"result": ss2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss3 := corestr.New.SimpleSlice.Default()
		actual = args.Map{"result": ss3.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss4 := corestr.New.SimpleSlice.Empty()
		actual = args.Map{"result": ss4.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovSSCreator_02_Strings_Create_Lines(t *testing.T) {
	safeTest(t, "Test_CovSSCreator_02_Strings_Create_Lines", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ss.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		ss2 := corestr.New.SimpleSlice.Create([]string{"a"})
		actual = args.Map{"result": ss2.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		ss3 := corestr.New.SimpleSlice.Lines("a", "b")
		actual = args.Map{"result": ss3.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		ss4 := corestr.New.SimpleSlice.SpreadStrings("a")
		actual = args.Map{"result": ss4.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovSSCreator_03_StringsPtr_StringsOptions_StringsClone(t *testing.T) {
	safeTest(t, "Test_CovSSCreator_03_StringsPtr_StringsOptions_StringsClone", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.StringsPtr([]string{"a"})

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		ss2 := corestr.New.SimpleSlice.StringsPtr([]string{})
		actual = args.Map{"result": ss2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss3 := corestr.New.SimpleSlice.StringsOptions(true, []string{"a"})
		actual = args.Map{"result": ss3.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		ss4 := corestr.New.SimpleSlice.StringsOptions(false, []string{"a"})
		actual = args.Map{"result": ss4.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		ss5 := corestr.New.SimpleSlice.StringsOptions(true, []string{})
		actual = args.Map{"result": ss5.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss6 := corestr.New.SimpleSlice.StringsClone([]string{"a"})
		actual = args.Map{"result": ss6.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		ss7 := corestr.New.SimpleSlice.StringsClone(nil)
		actual = args.Map{"result": ss7.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovSSCreator_04_Direct_UsingLines(t *testing.T) {
	safeTest(t, "Test_CovSSCreator_04_Direct_UsingLines", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Direct(true, []string{"a"})

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		ss2 := corestr.New.SimpleSlice.Direct(false, []string{"a"})
		actual = args.Map{"result": ss2.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		ss3 := corestr.New.SimpleSlice.Direct(true, nil)
		actual = args.Map{"result": ss3.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss4 := corestr.New.SimpleSlice.UsingLines(true, "a", "b")
		actual = args.Map{"result": ss4.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		ss5 := corestr.New.SimpleSlice.UsingLines(false, "a")
		actual = args.Map{"result": ss5.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		ss6 := corestr.New.SimpleSlice.UsingLines(true)
		actual = args.Map{"result": ss6.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovSSCreator_05_Split_SplitLines_UsingSeparatorLine_UsingLine(t *testing.T) {
	safeTest(t, "Test_CovSSCreator_05_Split_SplitLines_UsingSeparatorLine_UsingLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Split("a,b,c", ",")

		// Act
		actual := args.Map{"result": ss.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		ss2 := corestr.New.SimpleSlice.SplitLines("a\nb")
		actual = args.Map{"result": ss2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		ss3 := corestr.New.SimpleSlice.UsingSeparatorLine(",", "a,b")
		actual = args.Map{"result": ss3.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		ss4 := corestr.New.SimpleSlice.UsingLine("a\nb")
		_ = ss4
	})
}

func Test_CovSSCreator_06_ByLen(t *testing.T) {
	safeTest(t, "Test_CovSSCreator_06_ByLen", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.ByLen([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ss.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 length, just capacity", actual)
	})
}

func Test_CovSSCreator_07_Hashset_Map(t *testing.T) {
	safeTest(t, "Test_CovSSCreator_07_Hashset_Map", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		ss := corestr.New.SimpleSlice.Hashset(hs)

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		ssE := corestr.New.SimpleSlice.Hashset(corestr.New.Hashset.Empty())
		actual = args.Map{"result": ssE.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ssM := corestr.New.SimpleSlice.Map(map[string]string{"a": "1"})
		_ = ssM
		ssM2 := corestr.New.SimpleSlice.Map(nil)
		actual = args.Map{"result": ssM2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovSSCreator_08_Deserialize(t *testing.T) {
	safeTest(t, "Test_CovSSCreator_08_Deserialize", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		data, _ := ss.Serialize()
		ss2, err := corestr.New.SimpleSlice.Deserialize(data)

		// Act
		actual := args.Map{"result": err != nil || ss2.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error or wrong length", actual)
		_, err2 := corestr.New.SimpleSlice.Deserialize([]byte("bad"))
		actual = args.Map{"result": err2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

// --- utils ---

func Test_CovUtils_01_WrapDoubleIfMissing(t *testing.T) {
	safeTest(t, "Test_CovUtils_01_WrapDoubleIfMissing", func() {
		// Arrange
		u := corestr.StringUtils

		// Act
		actual := args.Map{"result": u.WrapDoubleIfMissing("hello") != `"hello"`}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
		actual = args.Map{"result": u.WrapDoubleIfMissing(`"hello"`) != `"hello"`}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected already wrapped", actual)
		actual = args.Map{"result": u.WrapDoubleIfMissing("") != `""`}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty quotes", actual)
		actual = args.Map{"result": u.WrapDoubleIfMissing(`""`) != `""`}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty quotes", actual)
	})
}

func Test_CovUtils_02_WrapSingleIfMissing(t *testing.T) {
	safeTest(t, "Test_CovUtils_02_WrapSingleIfMissing", func() {
		// Arrange
		u := corestr.StringUtils

		// Act
		actual := args.Map{"result": u.WrapSingleIfMissing("hello") != "'hello'"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
		actual = args.Map{"result": u.WrapSingleIfMissing("'hello'") != "'hello'"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected already wrapped", actual)
		actual = args.Map{"result": u.WrapSingleIfMissing("") != "''"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty quotes", actual)
		actual = args.Map{"result": u.WrapSingleIfMissing("''") != "''"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty quotes", actual)
	})
}

func Test_CovUtils_03_WrapDouble_WrapSingle_WrapTilda(t *testing.T) {
	safeTest(t, "Test_CovUtils_03_WrapDouble_WrapSingle_WrapTilda", func() {
		// Arrange
		u := corestr.StringUtils

		// Act
		actual := args.Map{"result": u.WrapDouble("hello") != `"hello"`}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
		actual = args.Map{"result": u.WrapSingle("hello") != "'hello'"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
		actual = args.Map{"result": u.WrapTilda("hello") != "`hello`"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
	})
}
