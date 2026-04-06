package corestrtests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
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
		actual := args.Map{"result": a.IsEqual(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S11_67_SimpleSlice_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_S11_67_SimpleSlice_IsEqual_BothNil", func() {
		var a *corestr.SimpleSlice
		var b *corestr.SimpleSlice
		actual := args.Map{"result": a.IsEqual(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S11_68_SimpleSlice_IsEqual_OneNil(t *testing.T) {
	safeTest(t, "Test_S11_68_SimpleSlice_IsEqual_OneNil", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"a"})
		var b *corestr.SimpleSlice
		actual := args.Map{"result": a.IsEqual(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S11_69_SimpleSlice_IsEqual_DiffLength(t *testing.T) {
	safeTest(t, "Test_S11_69_SimpleSlice_IsEqual_DiffLength", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"a"})
		b := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		actual := args.Map{"result": a.IsEqual(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S11_70_SimpleSlice_IsEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S11_70_SimpleSlice_IsEqual_BothEmpty", func() {
		a := corestr.Empty.SimpleSlice()
		b := corestr.Empty.SimpleSlice()
		actual := args.Map{"result": a.IsEqual(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S11_71_SimpleSlice_IsEqualLines(t *testing.T) {
	safeTest(t, "Test_S11_71_SimpleSlice_IsEqualLines", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		actual := args.Map{"result": ss.IsEqualLines([]string{"a", "b"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		actual := args.Map{"result": ss.IsEqualLines([]string{"a", "c"})}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S11_72_SimpleSlice_IsEqualLines_BothNil(t *testing.T) {
	safeTest(t, "Test_S11_72_SimpleSlice_IsEqualLines_BothNil", func() {
		var ss *corestr.SimpleSlice
		actual := args.Map{"result": ss.IsEqualLines(nil)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S11_73_SimpleSlice_IsEqualLines_OneNil(t *testing.T) {
	safeTest(t, "Test_S11_73_SimpleSlice_IsEqualLines_OneNil", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.IsEqualLines(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S11_74_SimpleSlice_IsEqualUnorderedLines(t *testing.T) {
	safeTest(t, "Test_S11_74_SimpleSlice_IsEqualUnorderedLines", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		actual := args.Map{"result": ss.IsEqualUnorderedLines([]string{"a", "b"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S11_75_SimpleSlice_IsEqualUnorderedLines_BothNil(t *testing.T) {
	safeTest(t, "Test_S11_75_SimpleSlice_IsEqualUnorderedLines_BothNil", func() {
		var ss *corestr.SimpleSlice
		actual := args.Map{"result": ss.IsEqualUnorderedLines(nil)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S11_76_SimpleSlice_IsEqualUnorderedLines_DiffLength(t *testing.T) {
	safeTest(t, "Test_S11_76_SimpleSlice_IsEqualUnorderedLines_DiffLength", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.IsEqualUnorderedLines([]string{"a", "b"})}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S11_77_SimpleSlice_IsEqualUnorderedLines_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S11_77_SimpleSlice_IsEqualUnorderedLines_BothEmpty", func() {
		ss := corestr.Empty.SimpleSlice()
		actual := args.Map{"result": ss.IsEqualUnorderedLines([]string{})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S11_78_SimpleSlice_IsEqualUnorderedLines_Mismatch(t *testing.T) {
	safeTest(t, "Test_S11_78_SimpleSlice_IsEqualUnorderedLines_Mismatch", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.IsEqualUnorderedLines([]string{"b"})}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S11_79_SimpleSlice_IsEqualUnorderedLinesClone(t *testing.T) {
	safeTest(t, "Test_S11_79_SimpleSlice_IsEqualUnorderedLinesClone", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		actual := args.Map{"result": ss.IsEqualUnorderedLinesClone([]string{"a", "b"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S11_80_SimpleSlice_IsEqualUnorderedLinesClone_BothNil(t *testing.T) {
	safeTest(t, "Test_S11_80_SimpleSlice_IsEqualUnorderedLinesClone_BothNil", func() {
		var ss *corestr.SimpleSlice
		actual := args.Map{"result": ss.IsEqualUnorderedLinesClone(nil)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S11_81_SimpleSlice_IsEqualUnorderedLinesClone_DiffLength(t *testing.T) {
	safeTest(t, "Test_S11_81_SimpleSlice_IsEqualUnorderedLinesClone_DiffLength", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.IsEqualUnorderedLinesClone([]string{"a", "b"})}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S11_82_SimpleSlice_IsEqualUnorderedLinesClone_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S11_82_SimpleSlice_IsEqualUnorderedLinesClone_BothEmpty", func() {
		ss := corestr.Empty.SimpleSlice()
		actual := args.Map{"result": ss.IsEqualUnorderedLinesClone([]string{})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S11_83_SimpleSlice_IsEqualUnorderedLinesClone_Mismatch(t *testing.T) {
	safeTest(t, "Test_S11_83_SimpleSlice_IsEqualUnorderedLinesClone_Mismatch", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.IsEqualUnorderedLinesClone([]string{"b"})}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S11_84_SimpleSlice_Collection(t *testing.T) {
	safeTest(t, "Test_S11_84_SimpleSlice_Collection", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.Collection(true).Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_85_SimpleSlice_NonPtr_Ptr(t *testing.T) {
	safeTest(t, "Test_S11_85_SimpleSlice_NonPtr_Ptr", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		_ = ss.NonPtr()
		actual := args.Map{"result": ss.Ptr() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S11_86_SimpleSlice_String(t *testing.T) {
	safeTest(t, "Test_S11_86_SimpleSlice_String", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		actual := args.Map{"result": corestr.Empty.SimpleSlice().String() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_S11_87_SimpleSlice_ConcatNewSimpleSlices(t *testing.T) {
	safeTest(t, "Test_S11_87_SimpleSlice_ConcatNewSimpleSlices", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"a"})
		b := corestr.New.SimpleSlice.Strings([]string{"b"})
		result := a.ConcatNewSimpleSlices(b)
		actual := args.Map{"result": result.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_S11_88_SimpleSlice_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_S11_88_SimpleSlice_ConcatNewStrings", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result := ss.ConcatNewStrings("b")
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S11_89_SimpleSlice_ConcatNewStrings_Nil(t *testing.T) {
	safeTest(t, "Test_S11_89_SimpleSlice_ConcatNewStrings_Nil", func() {
		var ss *corestr.SimpleSlice
		result := ss.ConcatNewStrings("b")
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_90_SimpleSlice_ConcatNew(t *testing.T) {
	safeTest(t, "Test_S11_90_SimpleSlice_ConcatNew", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result := ss.ConcatNew("b")
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S11_91_SimpleSlice_ToCollection(t *testing.T) {
	safeTest(t, "Test_S11_91_SimpleSlice_ToCollection", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.ToCollection(false).Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_92_SimpleSlice_CsvStrings(t *testing.T) {
	safeTest(t, "Test_S11_92_SimpleSlice_CsvStrings", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		csv := ss.CsvStrings()
		actual := args.Map{"result": len(csv) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": corestr.Empty.SimpleSlice().CsvStrings() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S11_93_SimpleSlice_JoinCsvString(t *testing.T) {
	safeTest(t, "Test_S11_93_SimpleSlice_JoinCsvString", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.JoinCsvString(",") == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		actual := args.Map{"result": corestr.Empty.SimpleSlice().JoinCsvString(",") != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_S11_94_SimpleSlice_JoinWith(t *testing.T) {
	safeTest(t, "Test_S11_94_SimpleSlice_JoinWith", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result := ss.JoinWith("|")
		actual := args.Map{"result": strings.HasPrefix(result, "|")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected prefix |", actual)
		actual := args.Map{"result": corestr.Empty.SimpleSlice().JoinWith("|") != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_S11_95_SimpleSlice_JsonModel(t *testing.T) {
	safeTest(t, "Test_S11_95_SimpleSlice_JsonModel", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": len(ss.JsonModel()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_96_SimpleSlice_Sort(t *testing.T) {
	safeTest(t, "Test_S11_96_SimpleSlice_Sort", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"c", "a", "b"})
		ss.Sort()
		actual := args.Map{"result": ss.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a first", actual)
	})
}

func Test_S11_97_SimpleSlice_Reverse(t *testing.T) {
	safeTest(t, "Test_S11_97_SimpleSlice_Reverse", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		ss.Reverse()
		actual := args.Map{"result": ss.First() != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c first", actual)
	})
}

func Test_S11_98_SimpleSlice_Reverse_Two(t *testing.T) {
	safeTest(t, "Test_S11_98_SimpleSlice_Reverse_Two", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		ss.Reverse()
		actual := args.Map{"result": ss.First() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_S11_99_SimpleSlice_Reverse_Single(t *testing.T) {
	safeTest(t, "Test_S11_99_SimpleSlice_Reverse_Single", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		ss.Reverse()
		actual := args.Map{"result": ss.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_S11_100_SimpleSlice_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_S11_100_SimpleSlice_MarshalJSON", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		data, err := ss.MarshalJSON()
		actual := args.Map{"result": err != nil || len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid JSON", actual)
	})
}

func Test_S11_101_SimpleSlice_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_S11_101_SimpleSlice_UnmarshalJSON", func() {
		ss := corestr.Empty.SimpleSlice()
		err := ss.UnmarshalJSON([]byte(`["a","b"]`))
		actual := args.Map{"result": err != nil || ss.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S11_102_SimpleSlice_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_S11_102_SimpleSlice_UnmarshalJSON_Invalid", func() {
		ss := corestr.Empty.SimpleSlice()
		err := ss.UnmarshalJSON([]byte(`invalid`))
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_S11_103_SimpleSlice_Json(t *testing.T) {
	safeTest(t, "Test_S11_103_SimpleSlice_Json", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		jsonResult := ss.Json()
		actual := args.Map{"result": jsonResult.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_S11_104_SimpleSlice_JsonPtr(t *testing.T) {
	safeTest(t, "Test_S11_104_SimpleSlice_JsonPtr", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.JsonPtr() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S11_105_SimpleSlice_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_S11_105_SimpleSlice_ParseInjectUsingJson", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		jsonResult := ss.JsonPtr()
		target := corestr.Empty.SimpleSlice()
		result, err := target.ParseInjectUsingJson(jsonResult)
		actual := args.Map{"result": err != nil || result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_106_SimpleSlice_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_S11_106_SimpleSlice_ParseInjectUsingJsonMust", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		jsonResult := ss.JsonPtr()
		target := corestr.Empty.SimpleSlice()
		result := target.ParseInjectUsingJsonMust(jsonResult)
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_107_SimpleSlice_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_S11_107_SimpleSlice_AsJsonContractsBinder", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S11_108_SimpleSlice_AsJsoner(t *testing.T) {
	safeTest(t, "Test_S11_108_SimpleSlice_AsJsoner", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.AsJsoner() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S11_109_SimpleSlice_ToPtr_ToNonPtr(t *testing.T) {
	safeTest(t, "Test_S11_109_SimpleSlice_ToPtr_ToNonPtr", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.ToPtr() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		_ = ss.ToNonPtr()
	})
}

func Test_S11_110_SimpleSlice_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_S11_110_SimpleSlice_JsonParseSelfInject", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		jsonResult := ss.JsonPtr()
		target := corestr.Empty.SimpleSlice()
		err := target.JsonParseSelfInject(jsonResult)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_S11_111_SimpleSlice_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_S11_111_SimpleSlice_AsJsonParseSelfInjector", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.AsJsonParseSelfInjector() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S11_112_SimpleSlice_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_S11_112_SimpleSlice_AsJsonMarshaller", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S11_113_SimpleSlice_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_S11_113_SimpleSlice_JsonModelAny", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.JsonModelAny() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S11_114_SimpleSlice_Clear(t *testing.T) {
	safeTest(t, "Test_S11_114_SimpleSlice_Clear", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		ss.Clear()
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S11_115_SimpleSlice_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_S11_115_SimpleSlice_Clear_Nil", func() {
		var ss *corestr.SimpleSlice
		actual := args.Map{"result": ss.Clear() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
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
		actual := args.Map{"result": cloned.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_119_SimpleSlice_ClonePtr(t *testing.T) {
	safeTest(t, "Test_S11_119_SimpleSlice_ClonePtr", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.ClonePtr(true).Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_120_SimpleSlice_ClonePtr_Nil(t *testing.T) {
	safeTest(t, "Test_S11_120_SimpleSlice_ClonePtr_Nil", func() {
		var ss *corestr.SimpleSlice
		actual := args.Map{"result": ss.ClonePtr(true) != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_S11_121_SimpleSlice_DeepClone(t *testing.T) {
	safeTest(t, "Test_S11_121_SimpleSlice_DeepClone", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.DeepClone().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_122_SimpleSlice_ShadowClone(t *testing.T) {
	safeTest(t, "Test_S11_122_SimpleSlice_ShadowClone", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.ShadowClone().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_123_SimpleSlice_IsDistinctEqualRaw(t *testing.T) {
	safeTest(t, "Test_S11_123_SimpleSlice_IsDistinctEqualRaw", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		actual := args.Map{"result": ss.IsDistinctEqualRaw("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S11_124_SimpleSlice_IsDistinctEqual(t *testing.T) {
	safeTest(t, "Test_S11_124_SimpleSlice_IsDistinctEqual", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"a"})
		b := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": a.IsDistinctEqual(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S11_125_SimpleSlice_IsUnorderedEqualRaw(t *testing.T) {
	safeTest(t, "Test_S11_125_SimpleSlice_IsUnorderedEqualRaw", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		actual := args.Map{"result": ss.IsUnorderedEqualRaw(true, "a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal with clone", actual)
		ss2 := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		actual := args.Map{"result": ss2.IsUnorderedEqualRaw(false, "a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal without clone", actual)
	})
}

func Test_S11_126_SimpleSlice_IsUnorderedEqualRaw_DiffLength(t *testing.T) {
	safeTest(t, "Test_S11_126_SimpleSlice_IsUnorderedEqualRaw_DiffLength", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.IsUnorderedEqualRaw(false, "a", "b")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S11_127_SimpleSlice_IsUnorderedEqualRaw_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S11_127_SimpleSlice_IsUnorderedEqualRaw_BothEmpty", func() {
		ss := corestr.Empty.SimpleSlice()
		actual := args.Map{"result": ss.IsUnorderedEqualRaw(false)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S11_128_SimpleSlice_IsUnorderedEqual(t *testing.T) {
	safeTest(t, "Test_S11_128_SimpleSlice_IsUnorderedEqual", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		b := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		actual := args.Map{"result": a.IsUnorderedEqual(true, b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S11_129_SimpleSlice_IsUnorderedEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_S11_129_SimpleSlice_IsUnorderedEqual_BothEmpty", func() {
		a := corestr.Empty.SimpleSlice()
		b := corestr.Empty.SimpleSlice()
		actual := args.Map{"result": a.IsUnorderedEqual(false, b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_S11_130_SimpleSlice_IsUnorderedEqual_NilRight(t *testing.T) {
	safeTest(t, "Test_S11_130_SimpleSlice_IsUnorderedEqual_NilRight", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": a.IsUnorderedEqual(false, nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_S11_131_SimpleSlice_IsEqualByFunc(t *testing.T) {
	safeTest(t, "Test_S11_131_SimpleSlice_IsEqualByFunc", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		result := ss.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "b")
		actual := args.Map{"result": result}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_S11_132_SimpleSlice_IsEqualByFunc_DiffLength(t *testing.T) {
	safeTest(t, "Test_S11_132_SimpleSlice_IsEqualByFunc_DiffLength", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.IsEqualByFunc(func(i int, l, r string) bool { return true }, "a", "b")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_S11_133_SimpleSlice_IsEqualByFunc_Empty(t *testing.T) {
	safeTest(t, "Test_S11_133_SimpleSlice_IsEqualByFunc_Empty", func() {
		ss := corestr.Empty.SimpleSlice()
		actual := args.Map{"result": ss.IsEqualByFunc(func(i int, l, r string) bool { return true })}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for both empty", actual)
	})
}

func Test_S11_134_SimpleSlice_IsEqualByFunc_Mismatch(t *testing.T) {
	safeTest(t, "Test_S11_134_SimpleSlice_IsEqualByFunc_Mismatch", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.IsEqualByFunc(func(i int, l, r string) bool { return false }, "a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_S11_135_SimpleSlice_IsEqualByFuncLinesSplit(t *testing.T) {
	safeTest(t, "Test_S11_135_SimpleSlice_IsEqualByFuncLinesSplit", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		result := ss.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return l == r })
		actual := args.Map{"result": result}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_S11_136_SimpleSlice_IsEqualByFuncLinesSplit_Trim(t *testing.T) {
	safeTest(t, "Test_S11_136_SimpleSlice_IsEqualByFuncLinesSplit_Trim", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{" a ", " b "})
		result := ss.IsEqualByFuncLinesSplit(true, ",", "a,b", func(i int, l, r string) bool { return l == r })
		actual := args.Map{"result": result}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true with trim", actual)
	})
}

func Test_S11_137_SimpleSlice_IsEqualByFuncLinesSplit_DiffLength(t *testing.T) {
	safeTest(t, "Test_S11_137_SimpleSlice_IsEqualByFuncLinesSplit_DiffLength", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return true })}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_S11_138_SimpleSlice_IsEqualByFuncLinesSplit_Empty(t *testing.T) {
	safeTest(t, "Test_S11_138_SimpleSlice_IsEqualByFuncLinesSplit_Empty", func() {
		ss := corestr.Empty.SimpleSlice()
		// strings.Split("", ",") returns [""] (length 1) which != 0, so returns false
		actual := args.Map{"result": ss.IsEqualByFuncLinesSplit(false, ",", "", func(i int, l, r string) bool { return true })}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty vs split-empty mismatch", actual)
	})
}

func Test_S11_139_SimpleSlice_IsEqualByFuncLinesSplit_Mismatch(t *testing.T) {
	safeTest(t, "Test_S11_139_SimpleSlice_IsEqualByFuncLinesSplit_Mismatch", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": ss.IsEqualByFuncLinesSplit(false, ",", "b", func(i int, l, r string) bool { return l == r })}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_S11_140_SimpleSlice_DistinctDiffRaw(t *testing.T) {
	safeTest(t, "Test_S11_140_SimpleSlice_DistinctDiffRaw", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		diff := ss.DistinctDiffRaw("b", "c")
		actual := args.Map{"result": len(diff) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S11_141_SimpleSlice_DistinctDiffRaw_BothNil(t *testing.T) {
	safeTest(t, "Test_S11_141_SimpleSlice_DistinctDiffRaw_BothNil", func() {
		var ss *corestr.SimpleSlice
		diff := ss.DistinctDiffRaw()
		actual := args.Map{"result": len(diff) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S11_142_SimpleSlice_DistinctDiffRaw_LeftNil(t *testing.T) {
	safeTest(t, "Test_S11_142_SimpleSlice_DistinctDiffRaw_LeftNil", func() {
		var ss *corestr.SimpleSlice
		diff := ss.DistinctDiffRaw("a")
		actual := args.Map{"result": len(diff) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_143_SimpleSlice_DistinctDiffRaw_RightNil(t *testing.T) {
	safeTest(t, "Test_S11_143_SimpleSlice_DistinctDiffRaw_RightNil", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		diff := ss.DistinctDiffRaw(nil...)
		actual := args.Map{"result": len(diff) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_144_SimpleSlice_AddedRemovedLinesDiff(t *testing.T) {
	safeTest(t, "Test_S11_144_SimpleSlice_AddedRemovedLinesDiff", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		added, removed := ss.AddedRemovedLinesDiff("b", "c")
		actual := args.Map{"result": len(added) != 1 || len(removed) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 added 1 removed", actual)
	})
}

func Test_S11_145_SimpleSlice_AddedRemovedLinesDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_S11_145_SimpleSlice_AddedRemovedLinesDiff_BothNil", func() {
		var ss *corestr.SimpleSlice
		added, removed := ss.AddedRemovedLinesDiff()
		actual := args.Map{"result": added != nil || removed != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_S11_146_SimpleSlice_RemoveIndexes(t *testing.T) {
	safeTest(t, "Test_S11_146_SimpleSlice_RemoveIndexes", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		result, err := ss.RemoveIndexes(1)
		actual := args.Map{"result": err != nil || result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S11_147_SimpleSlice_RemoveIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_S11_147_SimpleSlice_RemoveIndexes_Empty", func() {
		ss := corestr.Empty.SimpleSlice()
		_, err := ss.RemoveIndexes(0)
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_S11_148_SimpleSlice_RemoveIndexes_InvalidIndex(t *testing.T) {
	safeTest(t, "Test_S11_148_SimpleSlice_RemoveIndexes_InvalidIndex", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		_, err := ss.RemoveIndexes(5)
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error for invalid index", actual)
	})
}

func Test_S11_149_SimpleSlice_RemoveIndexes_AllRemoved(t *testing.T) {
	safeTest(t, "Test_S11_149_SimpleSlice_RemoveIndexes_AllRemoved", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result, err := ss.RemoveIndexes(0)
		actual := args.Map{"result": err != nil || result.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S11_150_SimpleSlice_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_S11_150_SimpleSlice_DistinctDiff", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		b := corestr.New.SimpleSlice.Strings([]string{"b", "c"})
		diff := a.DistinctDiff(b)
		actual := args.Map{"result": len(diff) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S11_151_SimpleSlice_DistinctDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_S11_151_SimpleSlice_DistinctDiff_BothNil", func() {
		var a *corestr.SimpleSlice
		diff := a.DistinctDiff(nil)
		actual := args.Map{"result": len(diff) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S11_152_SimpleSlice_DistinctDiff_LeftNil(t *testing.T) {
	safeTest(t, "Test_S11_152_SimpleSlice_DistinctDiff_LeftNil", func() {
		var a *corestr.SimpleSlice
		b := corestr.New.SimpleSlice.Strings([]string{"x"})
		diff := a.DistinctDiff(b)
		actual := args.Map{"result": len(diff) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_153_SimpleSlice_DistinctDiff_RightNil(t *testing.T) {
	safeTest(t, "Test_S11_153_SimpleSlice_DistinctDiff_RightNil", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"x"})
		diff := a.DistinctDiff(nil)
		actual := args.Map{"result": len(diff) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_154_SimpleSlice_Serialize(t *testing.T) {
	safeTest(t, "Test_S11_154_SimpleSlice_Serialize", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		data, err := ss.Serialize()
		actual := args.Map{"result": err != nil || len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid bytes", actual)
	})
}

func Test_S11_155_SimpleSlice_Deserialize(t *testing.T) {
	safeTest(t, "Test_S11_155_SimpleSlice_Deserialize", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		var target []string
		err := ss.Deserialize(&target)
		actual := args.Map{"result": err != nil || len(target) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_156_SimpleSlice_SafeStrings(t *testing.T) {
	safeTest(t, "Test_S11_156_SimpleSlice_SafeStrings", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual := args.Map{"result": len(ss.SafeStrings()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": len(corestr.Empty.SimpleSlice().SafeStrings()) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}
