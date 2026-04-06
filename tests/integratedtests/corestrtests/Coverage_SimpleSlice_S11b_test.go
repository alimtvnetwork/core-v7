package corestrtests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════
// S11b — SimpleSlice.go Lines 600-1317 — Equal, Clone, Diff, JSON
// ══════════════════════════════════════════════════════════════

func Test_S11_66_SimpleSlice_IsEqual(t *testing.T) {
	safeTest(t, "Test_S11_66_SimpleSlice_IsEqual", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		b := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		if !a.IsEqual(b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S11_67_SimpleSlice_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_S11_67_SimpleSlice_IsEqual_BothNil", func() {
		var a *corestr.SimpleSlice
		var b *corestr.SimpleSlice
		if !a.IsEqual(b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S11_68_SimpleSlice_IsEqual_OneNil(t *testing.T) {
	safeTest(t, "Test_S11_68_SimpleSlice_IsEqual_OneNil", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"a"})
		var b *corestr.SimpleSlice
		if a.IsEqual(b) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S11_69_SimpleSlice_IsEqual_DiffLength(t *testing.T) {
	safeTest(t, "Test_S11_69_SimpleSlice_IsEqual_DiffLength", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"a"})
		b := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		if a.IsEqual(b) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S11_70_SimpleSlice_IsEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S11_70_SimpleSlice_IsEqual_BothEmpty", func() {
		a := corestr.Empty.SimpleSlice()
		b := corestr.Empty.SimpleSlice()
		if !a.IsEqual(b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S11_71_SimpleSlice_IsEqualLines(t *testing.T) {
	safeTest(t, "Test_S11_71_SimpleSlice_IsEqualLines", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		if !ss.IsEqualLines([]string{"a", "b"}) {
			t.Fatal("expected equal")
		}
		if ss.IsEqualLines([]string{"a", "c"}) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S11_72_SimpleSlice_IsEqualLines_BothNil(t *testing.T) {
	safeTest(t, "Test_S11_72_SimpleSlice_IsEqualLines_BothNil", func() {
		var ss *corestr.SimpleSlice
		if !ss.IsEqualLines(nil) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S11_73_SimpleSlice_IsEqualLines_OneNil(t *testing.T) {
	safeTest(t, "Test_S11_73_SimpleSlice_IsEqualLines_OneNil", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.IsEqualLines(nil) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S11_74_SimpleSlice_IsEqualUnorderedLines(t *testing.T) {
	safeTest(t, "Test_S11_74_SimpleSlice_IsEqualUnorderedLines", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		if !ss.IsEqualUnorderedLines([]string{"a", "b"}) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S11_75_SimpleSlice_IsEqualUnorderedLines_BothNil(t *testing.T) {
	safeTest(t, "Test_S11_75_SimpleSlice_IsEqualUnorderedLines_BothNil", func() {
		var ss *corestr.SimpleSlice
		if !ss.IsEqualUnorderedLines(nil) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S11_76_SimpleSlice_IsEqualUnorderedLines_DiffLength(t *testing.T) {
	safeTest(t, "Test_S11_76_SimpleSlice_IsEqualUnorderedLines_DiffLength", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.IsEqualUnorderedLines([]string{"a", "b"}) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S11_77_SimpleSlice_IsEqualUnorderedLines_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S11_77_SimpleSlice_IsEqualUnorderedLines_BothEmpty", func() {
		ss := corestr.Empty.SimpleSlice()
		if !ss.IsEqualUnorderedLines([]string{}) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S11_78_SimpleSlice_IsEqualUnorderedLines_Mismatch(t *testing.T) {
	safeTest(t, "Test_S11_78_SimpleSlice_IsEqualUnorderedLines_Mismatch", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.IsEqualUnorderedLines([]string{"b"}) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S11_79_SimpleSlice_IsEqualUnorderedLinesClone(t *testing.T) {
	safeTest(t, "Test_S11_79_SimpleSlice_IsEqualUnorderedLinesClone", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		if !ss.IsEqualUnorderedLinesClone([]string{"a", "b"}) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S11_80_SimpleSlice_IsEqualUnorderedLinesClone_BothNil(t *testing.T) {
	safeTest(t, "Test_S11_80_SimpleSlice_IsEqualUnorderedLinesClone_BothNil", func() {
		var ss *corestr.SimpleSlice
		if !ss.IsEqualUnorderedLinesClone(nil) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S11_81_SimpleSlice_IsEqualUnorderedLinesClone_DiffLength(t *testing.T) {
	safeTest(t, "Test_S11_81_SimpleSlice_IsEqualUnorderedLinesClone_DiffLength", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.IsEqualUnorderedLinesClone([]string{"a", "b"}) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S11_82_SimpleSlice_IsEqualUnorderedLinesClone_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S11_82_SimpleSlice_IsEqualUnorderedLinesClone_BothEmpty", func() {
		ss := corestr.Empty.SimpleSlice()
		if !ss.IsEqualUnorderedLinesClone([]string{}) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S11_83_SimpleSlice_IsEqualUnorderedLinesClone_Mismatch(t *testing.T) {
	safeTest(t, "Test_S11_83_SimpleSlice_IsEqualUnorderedLinesClone_Mismatch", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.IsEqualUnorderedLinesClone([]string{"b"}) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S11_84_SimpleSlice_Collection(t *testing.T) {
	safeTest(t, "Test_S11_84_SimpleSlice_Collection", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.Collection(true).Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_85_SimpleSlice_NonPtr_Ptr(t *testing.T) {
	safeTest(t, "Test_S11_85_SimpleSlice_NonPtr_Ptr", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		_ = ss.NonPtr()
		if ss.Ptr() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S11_86_SimpleSlice_String(t *testing.T) {
	safeTest(t, "Test_S11_86_SimpleSlice_String", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.String() == "" {
			t.Fatal("expected non-empty")
		}
		if corestr.Empty.SimpleSlice().String() != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_S11_87_SimpleSlice_ConcatNewSimpleSlices(t *testing.T) {
	safeTest(t, "Test_S11_87_SimpleSlice_ConcatNewSimpleSlices", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"a"})
		b := corestr.New.SimpleSlice.Strings([]string{"b"})
		result := a.ConcatNewSimpleSlices(b)
		if result.Length() < 2 {
			t.Fatal("expected at least 2")
		}
	})
}

func Test_S11_88_SimpleSlice_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_S11_88_SimpleSlice_ConcatNewStrings", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result := ss.ConcatNewStrings("b")
		if len(result) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S11_89_SimpleSlice_ConcatNewStrings_Nil(t *testing.T) {
	safeTest(t, "Test_S11_89_SimpleSlice_ConcatNewStrings_Nil", func() {
		var ss *corestr.SimpleSlice
		result := ss.ConcatNewStrings("b")
		if len(result) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_90_SimpleSlice_ConcatNew(t *testing.T) {
	safeTest(t, "Test_S11_90_SimpleSlice_ConcatNew", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result := ss.ConcatNew("b")
		if result.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S11_91_SimpleSlice_ToCollection(t *testing.T) {
	safeTest(t, "Test_S11_91_SimpleSlice_ToCollection", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.ToCollection(false).Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_92_SimpleSlice_CsvStrings(t *testing.T) {
	safeTest(t, "Test_S11_92_SimpleSlice_CsvStrings", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		csv := ss.CsvStrings()
		if len(csv) != 1 {
			t.Fatal("expected 1")
		}
		if corestr.Empty.SimpleSlice().CsvStrings() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S11_93_SimpleSlice_JoinCsvString(t *testing.T) {
	safeTest(t, "Test_S11_93_SimpleSlice_JoinCsvString", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.JoinCsvString(",") == "" {
			t.Fatal("expected non-empty")
		}
		if corestr.Empty.SimpleSlice().JoinCsvString(",") != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_S11_94_SimpleSlice_JoinWith(t *testing.T) {
	safeTest(t, "Test_S11_94_SimpleSlice_JoinWith", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result := ss.JoinWith("|")
		if !strings.HasPrefix(result, "|") {
			t.Fatal("expected prefix |")
		}
		if corestr.Empty.SimpleSlice().JoinWith("|") != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_S11_95_SimpleSlice_JsonModel(t *testing.T) {
	safeTest(t, "Test_S11_95_SimpleSlice_JsonModel", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if len(ss.JsonModel()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_96_SimpleSlice_Sort(t *testing.T) {
	safeTest(t, "Test_S11_96_SimpleSlice_Sort", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"c", "a", "b"})
		ss.Sort()
		if ss.First() != "a" {
			t.Fatal("expected a first")
		}
	})
}

func Test_S11_97_SimpleSlice_Reverse(t *testing.T) {
	safeTest(t, "Test_S11_97_SimpleSlice_Reverse", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		ss.Reverse()
		if ss.First() != "c" {
			t.Fatal("expected c first")
		}
	})
}

func Test_S11_98_SimpleSlice_Reverse_Two(t *testing.T) {
	safeTest(t, "Test_S11_98_SimpleSlice_Reverse_Two", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		ss.Reverse()
		if ss.First() != "b" {
			t.Fatal("expected b")
		}
	})
}

func Test_S11_99_SimpleSlice_Reverse_Single(t *testing.T) {
	safeTest(t, "Test_S11_99_SimpleSlice_Reverse_Single", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		ss.Reverse()
		if ss.First() != "a" {
			t.Fatal("expected a")
		}
	})
}

func Test_S11_100_SimpleSlice_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_S11_100_SimpleSlice_MarshalJSON", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		data, err := ss.MarshalJSON()
		if err != nil || len(data) == 0 {
			t.Fatal("expected valid JSON")
		}
	})
}

func Test_S11_101_SimpleSlice_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_S11_101_SimpleSlice_UnmarshalJSON", func() {
		ss := corestr.Empty.SimpleSlice()
		err := ss.UnmarshalJSON([]byte(`["a","b"]`))
		if err != nil || ss.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S11_102_SimpleSlice_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_S11_102_SimpleSlice_UnmarshalJSON_Invalid", func() {
		ss := corestr.Empty.SimpleSlice()
		err := ss.UnmarshalJSON([]byte(`invalid`))
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_S11_103_SimpleSlice_Json(t *testing.T) {
	safeTest(t, "Test_S11_103_SimpleSlice_Json", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		jsonResult := ss.Json()
		if jsonResult.HasError() {
			t.Fatal("expected no error")
		}
	})
}

func Test_S11_104_SimpleSlice_JsonPtr(t *testing.T) {
	safeTest(t, "Test_S11_104_SimpleSlice_JsonPtr", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.JsonPtr() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S11_105_SimpleSlice_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_S11_105_SimpleSlice_ParseInjectUsingJson", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		jsonResult := ss.JsonPtr()
		target := corestr.Empty.SimpleSlice()
		result, err := target.ParseInjectUsingJson(jsonResult)
		if err != nil || result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_106_SimpleSlice_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_S11_106_SimpleSlice_ParseInjectUsingJsonMust", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		jsonResult := ss.JsonPtr()
		target := corestr.Empty.SimpleSlice()
		result := target.ParseInjectUsingJsonMust(jsonResult)
		if result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_107_SimpleSlice_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_S11_107_SimpleSlice_AsJsonContractsBinder", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.AsJsonContractsBinder() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S11_108_SimpleSlice_AsJsoner(t *testing.T) {
	safeTest(t, "Test_S11_108_SimpleSlice_AsJsoner", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.AsJsoner() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S11_109_SimpleSlice_ToPtr_ToNonPtr(t *testing.T) {
	safeTest(t, "Test_S11_109_SimpleSlice_ToPtr_ToNonPtr", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.ToPtr() == nil {
			t.Fatal("expected non-nil")
		}
		_ = ss.ToNonPtr()
	})
}

func Test_S11_110_SimpleSlice_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_S11_110_SimpleSlice_JsonParseSelfInject", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		jsonResult := ss.JsonPtr()
		target := corestr.Empty.SimpleSlice()
		err := target.JsonParseSelfInject(jsonResult)
		if err != nil {
			t.Fatal("expected no error")
		}
	})
}

func Test_S11_111_SimpleSlice_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_S11_111_SimpleSlice_AsJsonParseSelfInjector", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.AsJsonParseSelfInjector() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S11_112_SimpleSlice_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_S11_112_SimpleSlice_AsJsonMarshaller", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.AsJsonMarshaller() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S11_113_SimpleSlice_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_S11_113_SimpleSlice_JsonModelAny", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.JsonModelAny() == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S11_114_SimpleSlice_Clear(t *testing.T) {
	safeTest(t, "Test_S11_114_SimpleSlice_Clear", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		ss.Clear()
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S11_115_SimpleSlice_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_S11_115_SimpleSlice_Clear_Nil", func() {
		var ss *corestr.SimpleSlice
		if ss.Clear() != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_S11_116_SimpleSlice_Dispose(t *testing.T) {
	safeTest(t, "Test_S11_116_SimpleSlice_Dispose", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		ss.Dispose()
	})
}

func Test_S11_117_SimpleSlice_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_S11_117_SimpleSlice_Dispose_Nil", func() {
		var ss *corestr.SimpleSlice
		ss.Dispose()
	})
}

func Test_S11_118_SimpleSlice_Clone(t *testing.T) {
	safeTest(t, "Test_S11_118_SimpleSlice_Clone", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		cloned := ss.Clone(true)
		if cloned.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_119_SimpleSlice_ClonePtr(t *testing.T) {
	safeTest(t, "Test_S11_119_SimpleSlice_ClonePtr", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.ClonePtr(true).Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_120_SimpleSlice_ClonePtr_Nil(t *testing.T) {
	safeTest(t, "Test_S11_120_SimpleSlice_ClonePtr_Nil", func() {
		var ss *corestr.SimpleSlice
		if ss.ClonePtr(true) != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_S11_121_SimpleSlice_DeepClone(t *testing.T) {
	safeTest(t, "Test_S11_121_SimpleSlice_DeepClone", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.DeepClone().Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_122_SimpleSlice_ShadowClone(t *testing.T) {
	safeTest(t, "Test_S11_122_SimpleSlice_ShadowClone", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.ShadowClone().Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_123_SimpleSlice_IsDistinctEqualRaw(t *testing.T) {
	safeTest(t, "Test_S11_123_SimpleSlice_IsDistinctEqualRaw", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		if !ss.IsDistinctEqualRaw("a", "b") {
			t.Fatal("expected equal")
		}
	})
}

func Test_S11_124_SimpleSlice_IsDistinctEqual(t *testing.T) {
	safeTest(t, "Test_S11_124_SimpleSlice_IsDistinctEqual", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"a"})
		b := corestr.New.SimpleSlice.Strings([]string{"a"})
		if !a.IsDistinctEqual(b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S11_125_SimpleSlice_IsUnorderedEqualRaw(t *testing.T) {
	safeTest(t, "Test_S11_125_SimpleSlice_IsUnorderedEqualRaw", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		if !ss.IsUnorderedEqualRaw(true, "a", "b") {
			t.Fatal("expected equal with clone")
		}
		ss2 := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		if !ss2.IsUnorderedEqualRaw(false, "a", "b") {
			t.Fatal("expected equal without clone")
		}
	})
}

func Test_S11_126_SimpleSlice_IsUnorderedEqualRaw_DiffLength(t *testing.T) {
	safeTest(t, "Test_S11_126_SimpleSlice_IsUnorderedEqualRaw_DiffLength", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.IsUnorderedEqualRaw(false, "a", "b") {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S11_127_SimpleSlice_IsUnorderedEqualRaw_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S11_127_SimpleSlice_IsUnorderedEqualRaw_BothEmpty", func() {
		ss := corestr.Empty.SimpleSlice()
		if !ss.IsUnorderedEqualRaw(false) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S11_128_SimpleSlice_IsUnorderedEqual(t *testing.T) {
	safeTest(t, "Test_S11_128_SimpleSlice_IsUnorderedEqual", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		b := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		if !a.IsUnorderedEqual(true, b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S11_129_SimpleSlice_IsUnorderedEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S11_129_SimpleSlice_IsUnorderedEqual_BothEmpty", func() {
		a := corestr.Empty.SimpleSlice()
		b := corestr.Empty.SimpleSlice()
		if !a.IsUnorderedEqual(false, b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_S11_130_SimpleSlice_IsUnorderedEqual_NilRight(t *testing.T) {
	safeTest(t, "Test_S11_130_SimpleSlice_IsUnorderedEqual_NilRight", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"a"})
		if a.IsUnorderedEqual(false, nil) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_S11_131_SimpleSlice_IsEqualByFunc(t *testing.T) {
	safeTest(t, "Test_S11_131_SimpleSlice_IsEqualByFunc", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		result := ss.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "b")
		if !result {
			t.Fatal("expected true")
		}
	})
}

func Test_S11_132_SimpleSlice_IsEqualByFunc_DiffLength(t *testing.T) {
	safeTest(t, "Test_S11_132_SimpleSlice_IsEqualByFunc_DiffLength", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.IsEqualByFunc(func(i int, l, r string) bool { return true }, "a", "b") {
			t.Fatal("expected false")
		}
	})
}

func Test_S11_133_SimpleSlice_IsEqualByFunc_Empty(t *testing.T) {
	safeTest(t, "Test_S11_133_SimpleSlice_IsEqualByFunc_Empty", func() {
		ss := corestr.Empty.SimpleSlice()
		if !ss.IsEqualByFunc(func(i int, l, r string) bool { return true }) {
			t.Fatal("expected true for both empty")
		}
	})
}

func Test_S11_134_SimpleSlice_IsEqualByFunc_Mismatch(t *testing.T) {
	safeTest(t, "Test_S11_134_SimpleSlice_IsEqualByFunc_Mismatch", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.IsEqualByFunc(func(i int, l, r string) bool { return false }, "a") {
			t.Fatal("expected false")
		}
	})
}

func Test_S11_135_SimpleSlice_IsEqualByFuncLinesSplit(t *testing.T) {
	safeTest(t, "Test_S11_135_SimpleSlice_IsEqualByFuncLinesSplit", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		result := ss.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return l == r })
		if !result {
			t.Fatal("expected true")
		}
	})
}

func Test_S11_136_SimpleSlice_IsEqualByFuncLinesSplit_Trim(t *testing.T) {
	safeTest(t, "Test_S11_136_SimpleSlice_IsEqualByFuncLinesSplit_Trim", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{" a ", " b "})
		result := ss.IsEqualByFuncLinesSplit(true, ",", "a,b", func(i int, l, r string) bool { return l == r })
		if !result {
			t.Fatal("expected true with trim")
		}
	})
}

func Test_S11_137_SimpleSlice_IsEqualByFuncLinesSplit_DiffLength(t *testing.T) {
	safeTest(t, "Test_S11_137_SimpleSlice_IsEqualByFuncLinesSplit_DiffLength", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return true }) {
			t.Fatal("expected false")
		}
	})
}

func Test_S11_138_SimpleSlice_IsEqualByFuncLinesSplit_Empty(t *testing.T) {
	safeTest(t, "Test_S11_138_SimpleSlice_IsEqualByFuncLinesSplit_Empty", func() {
		ss := corestr.Empty.SimpleSlice()
		// strings.Split("", ",") returns [""] (length 1) which != 0, so returns false
		if ss.IsEqualByFuncLinesSplit(false, ",", "", func(i int, l, r string) bool { return true }) {
			t.Fatal("expected false for empty vs split-empty mismatch")
		}
	})
}

func Test_S11_139_SimpleSlice_IsEqualByFuncLinesSplit_Mismatch(t *testing.T) {
	safeTest(t, "Test_S11_139_SimpleSlice_IsEqualByFuncLinesSplit_Mismatch", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.IsEqualByFuncLinesSplit(false, ",", "b", func(i int, l, r string) bool { return l == r }) {
			t.Fatal("expected false")
		}
	})
}

func Test_S11_140_SimpleSlice_DistinctDiffRaw(t *testing.T) {
	safeTest(t, "Test_S11_140_SimpleSlice_DistinctDiffRaw", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		diff := ss.DistinctDiffRaw("b", "c")
		if len(diff) != 2 {
			t.Fatalf("expected 2, got %d", len(diff))
		}
	})
}

func Test_S11_141_SimpleSlice_DistinctDiffRaw_BothNil(t *testing.T) {
	safeTest(t, "Test_S11_141_SimpleSlice_DistinctDiffRaw_BothNil", func() {
		var ss *corestr.SimpleSlice
		diff := ss.DistinctDiffRaw()
		if len(diff) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S11_142_SimpleSlice_DistinctDiffRaw_LeftNil(t *testing.T) {
	safeTest(t, "Test_S11_142_SimpleSlice_DistinctDiffRaw_LeftNil", func() {
		var ss *corestr.SimpleSlice
		diff := ss.DistinctDiffRaw("a")
		if len(diff) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_143_SimpleSlice_DistinctDiffRaw_RightNil(t *testing.T) {
	safeTest(t, "Test_S11_143_SimpleSlice_DistinctDiffRaw_RightNil", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		diff := ss.DistinctDiffRaw(nil...)
		if len(diff) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_144_SimpleSlice_AddedRemovedLinesDiff(t *testing.T) {
	safeTest(t, "Test_S11_144_SimpleSlice_AddedRemovedLinesDiff", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		added, removed := ss.AddedRemovedLinesDiff("b", "c")
		if len(added) != 1 || len(removed) != 1 {
			t.Fatalf("expected 1 added 1 removed, got %d %d", len(added), len(removed))
		}
	})
}

func Test_S11_145_SimpleSlice_AddedRemovedLinesDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_S11_145_SimpleSlice_AddedRemovedLinesDiff_BothNil", func() {
		var ss *corestr.SimpleSlice
		added, removed := ss.AddedRemovedLinesDiff()
		if added != nil || removed != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_S11_146_SimpleSlice_RemoveIndexes(t *testing.T) {
	safeTest(t, "Test_S11_146_SimpleSlice_RemoveIndexes", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		result, err := ss.RemoveIndexes(1)
		if err != nil || result.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S11_147_SimpleSlice_RemoveIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_S11_147_SimpleSlice_RemoveIndexes_Empty", func() {
		ss := corestr.Empty.SimpleSlice()
		_, err := ss.RemoveIndexes(0)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_S11_148_SimpleSlice_RemoveIndexes_InvalidIndex(t *testing.T) {
	safeTest(t, "Test_S11_148_SimpleSlice_RemoveIndexes_InvalidIndex", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		_, err := ss.RemoveIndexes(5)
		if err == nil {
			t.Fatal("expected error for invalid index")
		}
	})
}

func Test_S11_149_SimpleSlice_RemoveIndexes_AllRemoved(t *testing.T) {
	safeTest(t, "Test_S11_149_SimpleSlice_RemoveIndexes_AllRemoved", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result, err := ss.RemoveIndexes(0)
		if err != nil || result.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S11_150_SimpleSlice_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_S11_150_SimpleSlice_DistinctDiff", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		b := corestr.New.SimpleSlice.Strings([]string{"b", "c"})
		diff := a.DistinctDiff(b)
		if len(diff) != 2 {
			t.Fatalf("expected 2, got %d", len(diff))
		}
	})
}

func Test_S11_151_SimpleSlice_DistinctDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_S11_151_SimpleSlice_DistinctDiff_BothNil", func() {
		var a *corestr.SimpleSlice
		diff := a.DistinctDiff(nil)
		if len(diff) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S11_152_SimpleSlice_DistinctDiff_LeftNil(t *testing.T) {
	safeTest(t, "Test_S11_152_SimpleSlice_DistinctDiff_LeftNil", func() {
		var a *corestr.SimpleSlice
		b := corestr.New.SimpleSlice.Strings([]string{"x"})
		diff := a.DistinctDiff(b)
		if len(diff) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_153_SimpleSlice_DistinctDiff_RightNil(t *testing.T) {
	safeTest(t, "Test_S11_153_SimpleSlice_DistinctDiff_RightNil", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"x"})
		diff := a.DistinctDiff(nil)
		if len(diff) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_154_SimpleSlice_Serialize(t *testing.T) {
	safeTest(t, "Test_S11_154_SimpleSlice_Serialize", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		data, err := ss.Serialize()
		if err != nil || len(data) == 0 {
			t.Fatal("expected valid bytes")
		}
	})
}

func Test_S11_155_SimpleSlice_Deserialize(t *testing.T) {
	safeTest(t, "Test_S11_155_SimpleSlice_Deserialize", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		var target []string
		err := ss.Deserialize(&target)
		if err != nil || len(target) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_156_SimpleSlice_SafeStrings(t *testing.T) {
	safeTest(t, "Test_S11_156_SimpleSlice_SafeStrings", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if len(ss.SafeStrings()) != 1 {
			t.Fatal("expected 1")
		}
		if len(corestr.Empty.SimpleSlice().SafeStrings()) != 0 {
			t.Fatal("expected 0")
		}
	})
}
