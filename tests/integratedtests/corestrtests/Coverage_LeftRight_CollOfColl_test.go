package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight + LeftRightFromSplit — Segment 20
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovLR_01_NewLeftRight(t *testing.T) {
	safeTest(t, "Test_CovLR_01_NewLeftRight", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
		actual = args.Map{"result": lr.IsValid}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected valid", actual)
	})
}

func Test_CovLR_02_InvalidLeftRight(t *testing.T) {
	safeTest(t, "Test_CovLR_02_InvalidLeftRight", func() {
		// Arrange
		lr := corestr.InvalidLeftRight("msg")

		// Act
		actual := args.Map{"result": lr.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
		lr2 := corestr.InvalidLeftRightNoMessage()
		actual = args.Map{"result": lr2.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_CovLR_03_LeftRightUsingSlice(t *testing.T) {
	safeTest(t, "Test_CovLR_03_LeftRightUsingSlice", func() {
		// Arrange
		// 2 items
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})

		// Act
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
		// 1 item
		lr2 := corestr.LeftRightUsingSlice([]string{"a"})
		actual = args.Map{"result": lr2.Left != "a" || lr2.Right != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,empty", actual)
		// 0 items
		lr3 := corestr.LeftRightUsingSlice([]string{})
		actual = args.Map{"result": lr3.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
		// deprecated ptr
		lr4 := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
		actual = args.Map{"result": lr4.Left != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		lr5 := corestr.LeftRightUsingSlicePtr([]string{})
		actual = args.Map{"result": lr5.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_CovLR_04_LeftRightTrimmedUsingSlice(t *testing.T) {
	safeTest(t, "Test_CovLR_04_LeftRightTrimmedUsingSlice", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})

		// Act
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed a,b", actual)
		// nil
		lr2 := corestr.LeftRightTrimmedUsingSlice(nil)
		actual = args.Map{"result": lr2.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
		// 0 items
		lr3 := corestr.LeftRightTrimmedUsingSlice([]string{})
		actual = args.Map{"result": lr3.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
		// 1 item
		lr4 := corestr.LeftRightTrimmedUsingSlice([]string{"a"})
		actual = args.Map{"result": lr4.Left != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_CovLR_05_LeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_CovLR_05_LeftRightFromSplit", func() {
		// Arrange
		lr := corestr.LeftRightFromSplit("key=value", "=")

		// Act
		actual := args.Map{"result": lr.Left != "key" || lr.Right != "value"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected key,value", actual)
	})
}

func Test_CovLR_06_LeftRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_CovLR_06_LeftRightFromSplitTrimmed", func() {
		// Arrange
		lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")

		// Act
		actual := args.Map{"result": lr.Left != "key" || lr.Right != "value"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed, got '',''", actual)
	})
}

func Test_CovLR_07_LeftRightFromSplitFull(t *testing.T) {
	safeTest(t, "Test_CovLR_07_LeftRightFromSplitFull", func() {
		// Arrange
		lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")

		// Act
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b:c:d"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b:c:d got '',''", actual)
	})
}

func Test_CovLR_08_LeftRightFromSplitFullTrimmed(t *testing.T) {
	safeTest(t, "Test_CovLR_08_LeftRightFromSplitFullTrimmed", func() {
		// Arrange
		lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")

		// Act
		actual := args.Map{"result": lr.Left != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, got ''", actual)
	})
}

func Test_CovLR_09_LeftBytes_RightBytes(t *testing.T) {
	safeTest(t, "Test_CovLR_09_LeftBytes_RightBytes", func() {
		// Arrange
		lr := corestr.NewLeftRight("ab", "cd")

		// Act
		actual := args.Map{"result": len(lr.LeftBytes()) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual = args.Map{"result": len(lr.RightBytes()) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLR_10_LeftTrim_RightTrim(t *testing.T) {
	safeTest(t, "Test_CovLR_10_LeftTrim_RightTrim", func() {
		// Arrange
		lr := corestr.NewLeftRight(" a ", " b ")

		// Act
		actual := args.Map{"result": lr.LeftTrim() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": lr.RightTrim() != "b"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_CovLR_11_IsLeftEmpty_IsRightEmpty(t *testing.T) {
	safeTest(t, "Test_CovLR_11_IsLeftEmpty_IsRightEmpty", func() {
		// Arrange
		lr := corestr.NewLeftRight("", "b")

		// Act
		actual := args.Map{"result": lr.IsLeftEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": lr.IsRightEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovLR_12_IsLeftWhitespace_IsRightWhitespace(t *testing.T) {
	safeTest(t, "Test_CovLR_12_IsLeftWhitespace_IsRightWhitespace", func() {
		// Arrange
		lr := corestr.NewLeftRight("  ", "  ")

		// Act
		actual := args.Map{"result": lr.IsLeftWhitespace()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": lr.IsRightWhitespace()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovLR_13_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_CovLR_13_HasValidNonEmpty", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.HasValidNonEmptyLeft()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": lr.HasValidNonEmptyRight()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": lr.HasValidNonWhitespaceLeft()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": lr.HasValidNonWhitespaceRight()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": lr.HasSafeNonEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovLR_14_NonPtr_Ptr(t *testing.T) {
	safeTest(t, "Test_CovLR_14_NonPtr_Ptr", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		np := lr.NonPtr()

		// Act
		actual := args.Map{"result": np.Left != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		p := lr.Ptr()
		actual = args.Map{"result": p.Left != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_CovLR_15_IsLeftRegexMatch_IsRightRegexMatch(t *testing.T) {
	safeTest(t, "Test_CovLR_15_IsLeftRegexMatch_IsRightRegexMatch", func() {
		// Arrange
		lr := corestr.NewLeftRight("hello123", "world456")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{"result": lr.IsLeftRegexMatch(re)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": lr.IsRightRegexMatch(re)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": lr.IsLeftRegexMatch(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": lr.IsRightRegexMatch(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovLR_16_IsLeft_IsRight_Is(t *testing.T) {
	safeTest(t, "Test_CovLR_16_IsLeft_IsRight_Is", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.IsLeft("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": lr.IsRight("b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": lr.Is("a", "b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": lr.Is("x", "b")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovLR_17_IsEqual(t *testing.T) {
	safeTest(t, "Test_CovLR_17_IsEqual", func() {
		// Arrange
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr1.IsEqual(lr2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		lr3 := corestr.NewLeftRight("x", "b")
		actual = args.Map{"result": lr1.IsEqual(lr3)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
		actual = args.Map{"result": lr1.IsEqual(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
		var nilLR *corestr.LeftRight
		actual = args.Map{"result": nilLR.IsEqual(nil)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_CovLR_18_Clone_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_CovLR_18_Clone_Clear_Dispose", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		c := lr.Clone()

		// Act
		actual := args.Map{"result": c.Left != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		lr.Clear()
		lr2 := corestr.NewLeftRight("x", "y")
		lr2.Dispose()
		var nilLR *corestr.LeftRight
		nilLR.Clear()
		nilLR.Dispose()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionsOfCollection — Segment 20 Part 2
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovCoC_01_IsEmpty_HasItems_Length(t *testing.T) {
	safeTest(t, "Test_CovCoC_01_IsEmpty_HasItems_Length", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		actual := args.Map{"result": coc.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": coc.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
		actual = args.Map{"result": coc.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovCoC_02_Add_Adds_AddCollections(t *testing.T) {
	safeTest(t, "Test_CovCoC_02_Add_Adds_AddCollections", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		col := corestr.New.Collection.Strings([]string{"false", "a", "b"})
		coc.Add(col)

		// Act
		actual := args.Map{"result": coc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// Add empty skipped
		coc.Add(corestr.Empty.Collection())
		actual = args.Map{"result": coc.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		col2 := *corestr.New.Collection.Strings([]string{"false", "c"})
		coc.Adds(col2)
		actual = args.Map{"result": coc.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		coc.Adds()
		coc.AddCollections()
	})
}

func Test_CovCoC_03_AddStrings_AddsStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_CovCoC_03_AddStrings_AddsStringsOfStrings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{"a", "b"})

		// Act
		actual := args.Map{"result": coc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		coc.AddStrings(false, []string{})
		actual = args.Map{"result": coc.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		coc.AddsStringsOfStrings(false, []string{"c"}, []string{"d"})
		actual = args.Map{"result": coc.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		coc.AddsStringsOfStrings(false)
	})
}

func Test_CovCoC_04_AllIndividualItemsLength_Items_List(t *testing.T) {
	safeTest(t, "Test_CovCoC_04_AllIndividualItemsLength_Items_List", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		actual := args.Map{"result": coc.AllIndividualItemsLength() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		coc.AddStrings(false, []string{"a", "b"})
		coc.AddStrings(false, []string{"c"})
		actual = args.Map{"result": coc.AllIndividualItemsLength() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual = args.Map{"result": len(coc.Items()) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		list := coc.List(0)
		actual = args.Map{"result": len(list) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_CovCoC_05_ToCollection(t *testing.T) {
	safeTest(t, "Test_CovCoC_05_ToCollection", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{"a", "b"})
		col := coc.ToCollection()

		// Act
		actual := args.Map{"result": col.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovCoC_06_String(t *testing.T) {
	safeTest(t, "Test_CovCoC_06_String", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{"a"})
		_ = coc.String()
	})
}

func Test_CovCoC_07_JsonModel_MarshalUnmarshal(t *testing.T) {
	safeTest(t, "Test_CovCoC_07_JsonModel_MarshalUnmarshal", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{"a"})
		_ = coc.JsonModel()
		_ = coc.JsonModelAny()
		data, err := coc.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		coc2 := corestr.New.CollectionsOfCollection.Empty()
		err2 := coc2.UnmarshalJSON(data)
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		err3 := coc2.UnmarshalJSON([]byte("bad"))
		actual = args.Map{"result": err3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_CovCoC_08_Json_ParseInject(t *testing.T) {
	safeTest(t, "Test_CovCoC_08_Json_ParseInject", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{"a"})
		_ = coc.Json()
		jr := coc.JsonPtr()
		coc2 := corestr.New.CollectionsOfCollection.Empty()
		r, err := coc2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"result": err != nil || r == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_CovCoC_09_ParseInjectMust(t *testing.T) {
	safeTest(t, "Test_CovCoC_09_ParseInjectMust", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{"a"})
		jr := coc.JsonPtr()
		coc2 := corestr.New.CollectionsOfCollection.Empty()
		r := coc2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"result": r == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_CovCoC_10_JsonParseSelfInject_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_CovCoC_10_JsonParseSelfInject_AsInterfaces", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{"a"})
		jr := coc.JsonPtr()
		coc2 := corestr.New.CollectionsOfCollection.Empty()
		err := coc2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		_ = coc.AsJsonContractsBinder()
		_ = coc.AsJsoner()
		_ = coc.AsJsonParseSelfInjector()
		_ = coc.AsJsonMarshaller()
	})
}

func Test_CovCoC_11_Creators(t *testing.T) {
	safeTest(t, "Test_CovCoC_11_Creators", func() {
		// Arrange
		// LenCap
		c := corestr.New.CollectionsOfCollection.LenCap(0, 5)

		// Act
		actual := args.Map{"result": c.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// StringsOfStrings
		c2 := corestr.New.CollectionsOfCollection.StringsOfStrings(false, []string{"a"}, []string{"b"})
		actual = args.Map{"result": c2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// SpreadStrings
		c3 := corestr.New.CollectionsOfCollection.SpreadStrings(false, "a", "b")
		actual = args.Map{"result": c3.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// Strings
		c4 := corestr.New.CollectionsOfCollection.Strings([]string{"a"})
		actual = args.Map{"result": c4.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// CloneStrings
		c5 := corestr.New.CollectionsOfCollection.CloneStrings([]string{"a"})
		actual = args.Map{"result": c5.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// StringsOption
		c6 := corestr.New.CollectionsOfCollection.StringsOption(true, 5, []string{"a"})
		actual = args.Map{"result": c6.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// StringsOptions
		c7 := corestr.New.CollectionsOfCollection.StringsOptions(false, 5, []string{"a"})
		actual = args.Map{"result": c7.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovCoC_12_AddAsyncFuncItems(t *testing.T) {
	safeTest(t, "Test_CovCoC_12_AddAsyncFuncItems", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		// nil funcs
		coc.AddAsyncFuncItems(nil, false)
		// with funcs
		// We need a WaitGroup approach; skip if it requires external sync setup
	})
}
