package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight + LeftRightFromSplit — Segment 20
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovLR_01_NewLeftRight(t *testing.T) {
	safeTest(t, "Test_CovLR_01_NewLeftRight", func() {
		lr := corestr.NewLeftRight("a", "b")
		if lr.Left != "a" || lr.Right != "b" {
			t.Fatal("expected a,b")
		}
		if !lr.IsValid {
			t.Fatal("expected valid")
		}
	})
}

func Test_CovLR_02_InvalidLeftRight(t *testing.T) {
	safeTest(t, "Test_CovLR_02_InvalidLeftRight", func() {
		lr := corestr.InvalidLeftRight("msg")
		if lr.IsValid {
			t.Fatal("expected invalid")
		}
		lr2 := corestr.InvalidLeftRightNoMessage()
		if lr2.IsValid {
			t.Fatal("expected invalid")
		}
	})
}

func Test_CovLR_03_LeftRightUsingSlice(t *testing.T) {
	safeTest(t, "Test_CovLR_03_LeftRightUsingSlice", func() {
		// 2 items
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})
		if lr.Left != "a" || lr.Right != "b" {
			t.Fatal("expected a,b")
		}
		// 1 item
		lr2 := corestr.LeftRightUsingSlice([]string{"a"})
		if lr2.Left != "a" || lr2.Right != "" {
			t.Fatal("expected a,empty")
		}
		// 0 items
		lr3 := corestr.LeftRightUsingSlice([]string{})
		if lr3.IsValid {
			t.Fatal("expected invalid")
		}
		// deprecated ptr
		lr4 := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
		if lr4.Left != "a" {
			t.Fatal("expected a")
		}
		lr5 := corestr.LeftRightUsingSlicePtr([]string{})
		if lr5.IsValid {
			t.Fatal("expected invalid")
		}
	})
}

func Test_CovLR_04_LeftRightTrimmedUsingSlice(t *testing.T) {
	safeTest(t, "Test_CovLR_04_LeftRightTrimmedUsingSlice", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
		if lr.Left != "a" || lr.Right != "b" {
			t.Fatal("expected trimmed a,b")
		}
		// nil
		lr2 := corestr.LeftRightTrimmedUsingSlice(nil)
		if lr2.IsValid {
			t.Fatal("expected invalid")
		}
		// 0 items
		lr3 := corestr.LeftRightTrimmedUsingSlice([]string{})
		if lr3.IsValid {
			t.Fatal("expected invalid")
		}
		// 1 item
		lr4 := corestr.LeftRightTrimmedUsingSlice([]string{"a"})
		if lr4.Left != "a" {
			t.Fatal("expected a")
		}
	})
}

func Test_CovLR_05_LeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_CovLR_05_LeftRightFromSplit", func() {
		lr := corestr.LeftRightFromSplit("key=value", "=")
		if lr.Left != "key" || lr.Right != "value" {
			t.Fatal("expected key,value")
		}
	})
}

func Test_CovLR_06_LeftRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_CovLR_06_LeftRightFromSplitTrimmed", func() {
		lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")
		if lr.Left != "key" || lr.Right != "value" {
			t.Fatalf("expected trimmed, got '%s','%s'", lr.Left, lr.Right)
		}
	})
}

func Test_CovLR_07_LeftRightFromSplitFull(t *testing.T) {
	safeTest(t, "Test_CovLR_07_LeftRightFromSplitFull", func() {
		lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")
		if lr.Left != "a" || lr.Right != "b:c:d" {
			t.Fatalf("expected a,b:c:d got '%s','%s'", lr.Left, lr.Right)
		}
	})
}

func Test_CovLR_08_LeftRightFromSplitFullTrimmed(t *testing.T) {
	safeTest(t, "Test_CovLR_08_LeftRightFromSplitFullTrimmed", func() {
		lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")
		if lr.Left != "a" {
			t.Fatalf("expected a, got '%s'", lr.Left)
		}
	})
}

func Test_CovLR_09_LeftBytes_RightBytes(t *testing.T) {
	safeTest(t, "Test_CovLR_09_LeftBytes_RightBytes", func() {
		lr := corestr.NewLeftRight("ab", "cd")
		if len(lr.LeftBytes()) != 2 {
			t.Fatal("expected 2")
		}
		if len(lr.RightBytes()) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLR_10_LeftTrim_RightTrim(t *testing.T) {
	safeTest(t, "Test_CovLR_10_LeftTrim_RightTrim", func() {
		lr := corestr.NewLeftRight(" a ", " b ")
		if lr.LeftTrim() != "a" {
			t.Fatal("expected a")
		}
		if lr.RightTrim() != "b" {
			t.Fatal("expected b")
		}
	})
}

func Test_CovLR_11_IsLeftEmpty_IsRightEmpty(t *testing.T) {
	safeTest(t, "Test_CovLR_11_IsLeftEmpty_IsRightEmpty", func() {
		lr := corestr.NewLeftRight("", "b")
		if !lr.IsLeftEmpty() {
			t.Fatal("expected true")
		}
		if lr.IsRightEmpty() {
			t.Fatal("expected false")
		}
	})
}

func Test_CovLR_12_IsLeftWhitespace_IsRightWhitespace(t *testing.T) {
	safeTest(t, "Test_CovLR_12_IsLeftWhitespace_IsRightWhitespace", func() {
		lr := corestr.NewLeftRight("  ", "  ")
		if !lr.IsLeftWhitespace() {
			t.Fatal("expected true")
		}
		if !lr.IsRightWhitespace() {
			t.Fatal("expected true")
		}
	})
}

func Test_CovLR_13_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_CovLR_13_HasValidNonEmpty", func() {
		lr := corestr.NewLeftRight("a", "b")
		if !lr.HasValidNonEmptyLeft() {
			t.Fatal("expected true")
		}
		if !lr.HasValidNonEmptyRight() {
			t.Fatal("expected true")
		}
		if !lr.HasValidNonWhitespaceLeft() {
			t.Fatal("expected true")
		}
		if !lr.HasValidNonWhitespaceRight() {
			t.Fatal("expected true")
		}
		if !lr.HasSafeNonEmpty() {
			t.Fatal("expected true")
		}
	})
}

func Test_CovLR_14_NonPtr_Ptr(t *testing.T) {
	safeTest(t, "Test_CovLR_14_NonPtr_Ptr", func() {
		lr := corestr.NewLeftRight("a", "b")
		np := lr.NonPtr()
		if np.Left != "a" {
			t.Fatal("expected a")
		}
		p := lr.Ptr()
		if p.Left != "a" {
			t.Fatal("expected a")
		}
	})
}

func Test_CovLR_15_IsLeftRegexMatch_IsRightRegexMatch(t *testing.T) {
	safeTest(t, "Test_CovLR_15_IsLeftRegexMatch_IsRightRegexMatch", func() {
		lr := corestr.NewLeftRight("hello123", "world456")
		re := regexp.MustCompile(`\d+`)
		if !lr.IsLeftRegexMatch(re) {
			t.Fatal("expected true")
		}
		if !lr.IsRightRegexMatch(re) {
			t.Fatal("expected true")
		}
		if lr.IsLeftRegexMatch(nil) {
			t.Fatal("expected false")
		}
		if lr.IsRightRegexMatch(nil) {
			t.Fatal("expected false")
		}
	})
}

func Test_CovLR_16_IsLeft_IsRight_Is(t *testing.T) {
	safeTest(t, "Test_CovLR_16_IsLeft_IsRight_Is", func() {
		lr := corestr.NewLeftRight("a", "b")
		if !lr.IsLeft("a") {
			t.Fatal("expected true")
		}
		if !lr.IsRight("b") {
			t.Fatal("expected true")
		}
		if !lr.Is("a", "b") {
			t.Fatal("expected true")
		}
		if lr.Is("x", "b") {
			t.Fatal("expected false")
		}
	})
}

func Test_CovLR_17_IsEqual(t *testing.T) {
	safeTest(t, "Test_CovLR_17_IsEqual", func() {
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		if !lr1.IsEqual(lr2) {
			t.Fatal("expected equal")
		}
		lr3 := corestr.NewLeftRight("x", "b")
		if lr1.IsEqual(lr3) {
			t.Fatal("expected not equal")
		}
		if lr1.IsEqual(nil) {
			t.Fatal("expected not equal")
		}
		var nilLR *corestr.LeftRight
		if !nilLR.IsEqual(nil) {
			t.Fatal("expected equal")
		}
	})
}

func Test_CovLR_18_Clone_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_CovLR_18_Clone_Clear_Dispose", func() {
		lr := corestr.NewLeftRight("a", "b")
		c := lr.Clone()
		if c.Left != "a" {
			t.Fatal("expected a")
		}
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
		coc := corestr.New.CollectionsOfCollection.Empty()
		if !coc.IsEmpty() {
			t.Fatal("expected empty")
		}
		if coc.HasItems() {
			t.Fatal("expected no items")
		}
		if coc.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovCoC_02_Add_Adds_AddCollections(t *testing.T) {
	safeTest(t, "Test_CovCoC_02_Add_Adds_AddCollections", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		col := corestr.New.Collection.Strings([]string{"false", "a", "b"})
		coc.Add(col)
		if coc.Length() != 1 {
			t.Fatal("expected 1")
		}
		// Add empty skipped
		coc.Add(corestr.Empty.Collection())
		if coc.Length() != 1 {
			t.Fatal("expected 1")
		}
		col2 := *corestr.New.Collection.Strings([]string{"false", "c"})
		coc.Adds(col2)
		if coc.Length() != 2 {
			t.Fatal("expected 2")
		}
		coc.Adds()
		coc.AddCollections()
	})
}

func Test_CovCoC_03_AddStrings_AddsStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_CovCoC_03_AddStrings_AddsStringsOfStrings", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{"a", "b"})
		if coc.Length() != 1 {
			t.Fatal("expected 1")
		}
		coc.AddStrings(false, []string{})
		if coc.Length() != 1 {
			t.Fatal("expected 1")
		}
		coc.AddsStringsOfStrings(false, []string{"c"}, []string{"d"})
		if coc.Length() != 3 {
			t.Fatal("expected 3")
		}
		coc.AddsStringsOfStrings(false)
	})
}

func Test_CovCoC_04_AllIndividualItemsLength_Items_List(t *testing.T) {
	safeTest(t, "Test_CovCoC_04_AllIndividualItemsLength_Items_List", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		if coc.AllIndividualItemsLength() != 0 {
			t.Fatal("expected 0")
		}
		coc.AddStrings(false, []string{"a", "b"})
		coc.AddStrings(false, []string{"c"})
		if coc.AllIndividualItemsLength() != 3 {
			t.Fatal("expected 3")
		}
		if len(coc.Items()) != 2 {
			t.Fatal("expected 2")
		}
		list := coc.List(0)
		if len(list) != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_CovCoC_05_ToCollection(t *testing.T) {
	safeTest(t, "Test_CovCoC_05_ToCollection", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{"a", "b"})
		col := coc.ToCollection()
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
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
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{"a"})
		_ = coc.JsonModel()
		_ = coc.JsonModelAny()
		data, err := coc.MarshalJSON()
		if err != nil {
			t.Fatal("unexpected error")
		}
		coc2 := corestr.New.CollectionsOfCollection.Empty()
		err2 := coc2.UnmarshalJSON(data)
		if err2 != nil {
			t.Fatal("unexpected error")
		}
		err3 := coc2.UnmarshalJSON([]byte("bad"))
		if err3 == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_CovCoC_08_Json_ParseInject(t *testing.T) {
	safeTest(t, "Test_CovCoC_08_Json_ParseInject", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{"a"})
		_ = coc.Json()
		jr := coc.JsonPtr()
		coc2 := corestr.New.CollectionsOfCollection.Empty()
		r, err := coc2.ParseInjectUsingJson(jr)
		if err != nil || r == nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovCoC_09_ParseInjectMust(t *testing.T) {
	safeTest(t, "Test_CovCoC_09_ParseInjectMust", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{"a"})
		jr := coc.JsonPtr()
		coc2 := corestr.New.CollectionsOfCollection.Empty()
		r := coc2.ParseInjectUsingJsonMust(jr)
		if r == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_CovCoC_10_JsonParseSelfInject_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_CovCoC_10_JsonParseSelfInject_AsInterfaces", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{"a"})
		jr := coc.JsonPtr()
		coc2 := corestr.New.CollectionsOfCollection.Empty()
		err := coc2.JsonParseSelfInject(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
		_ = coc.AsJsonContractsBinder()
		_ = coc.AsJsoner()
		_ = coc.AsJsonParseSelfInjector()
		_ = coc.AsJsonMarshaller()
	})
}

func Test_CovCoC_11_Creators(t *testing.T) {
	safeTest(t, "Test_CovCoC_11_Creators", func() {
		// LenCap
		c := corestr.New.CollectionsOfCollection.LenCap(0, 5)
		if c.Length() != 0 {
			t.Fatal("expected 0")
		}
		// StringsOfStrings
		c2 := corestr.New.CollectionsOfCollection.StringsOfStrings(false, []string{"a"}, []string{"b"})
		if c2.Length() != 2 {
			t.Fatal("expected 2")
		}
		// SpreadStrings
		c3 := corestr.New.CollectionsOfCollection.SpreadStrings(false, "a", "b")
		if c3.Length() != 1 {
			t.Fatal("expected 1")
		}
		// Strings
		c4 := corestr.New.CollectionsOfCollection.Strings([]string{"a"})
		if c4.Length() != 1 {
			t.Fatal("expected 1")
		}
		// CloneStrings
		c5 := corestr.New.CollectionsOfCollection.CloneStrings([]string{"a"})
		if c5.Length() != 1 {
			t.Fatal("expected 1")
		}
		// StringsOption
		c6 := corestr.New.CollectionsOfCollection.StringsOption(true, 5, []string{"a"})
		if c6.Length() != 1 {
			t.Fatal("expected 1")
		}
		// StringsOptions
		c7 := corestr.New.CollectionsOfCollection.StringsOptions(false, 5, []string{"a"})
		if c7.Length() != 1 {
			t.Fatal("expected 1")
		}
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
