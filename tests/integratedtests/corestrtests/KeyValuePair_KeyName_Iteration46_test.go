package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════
// KeyValuePair
// ═══════════════════════════════════════════════════════════════

func Test_KeyValuePair_KeyName(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_KeyName", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KVP KeyName", Expected: "k", Actual: kv.KeyName(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_VariableName_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_VariableName", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KVP VariableName", Expected: "k", Actual: kv.VariableName(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_ValueString(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_ValueString", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KVP ValueString", Expected: "v", Actual: kv.ValueString(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_IsVariableNameEqual_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_IsVariableNameEqual", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KVP IsVarNameEqual", Expected: true, Actual: kv.IsVariableNameEqual("k"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_IsValueEqual_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_IsValueEqual", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KVP IsValueEqual", Expected: true, Actual: kv.IsValueEqual("v"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_IsKeyEmpty(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_IsKeyEmpty", func() {
		kv := corestr.KeyValuePair{Key: "", Value: "v"}
		tc := caseV1Compat{Name: "KVP IsKeyEmpty", Expected: true, Actual: kv.IsKeyEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_IsValueEmpty(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_IsValueEmpty", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: ""}
		tc := caseV1Compat{Name: "KVP IsValueEmpty", Expected: true, Actual: kv.IsValueEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_HasKey(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_HasKey", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KVP HasKey", Expected: true, Actual: kv.HasKey(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_HasValue(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_HasValue", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KVP HasValue", Expected: true, Actual: kv.HasValue(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_IsKeyValueEmpty(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_IsKeyValueEmpty", func() {
		kv := corestr.KeyValuePair{}
		tc := caseV1Compat{Name: "KVP IsKeyValueEmpty", Expected: true, Actual: kv.IsKeyValueEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_TrimKey(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_TrimKey", func() {
		kv := corestr.KeyValuePair{Key: " k ", Value: "v"}
		tc := caseV1Compat{Name: "KVP TrimKey", Expected: "k", Actual: kv.TrimKey(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_TrimValue(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_TrimValue", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: " v "}
		tc := caseV1Compat{Name: "KVP TrimValue", Expected: "v", Actual: kv.TrimValue(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_ValueBool_True(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_ValueBool_True", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "true"}
		tc := caseV1Compat{Name: "KVP ValueBool true", Expected: true, Actual: kv.ValueBool(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_ValueBool_Empty_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_ValueBool_Empty", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: ""}
		tc := caseV1Compat{Name: "KVP ValueBool empty", Expected: false, Actual: kv.ValueBool(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_ValueBool_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_ValueBool_Invalid", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "notbool"}
		tc := caseV1Compat{Name: "KVP ValueBool invalid", Expected: false, Actual: kv.ValueBool(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_ValueInt_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_ValueInt", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "42"}
		tc := caseV1Compat{Name: "KVP ValueInt", Expected: 42, Actual: kv.ValueInt(0), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_ValueInt_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_ValueInt_Invalid", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "abc"}
		tc := caseV1Compat{Name: "KVP ValueInt invalid", Expected: 99, Actual: kv.ValueInt(99), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_ValueDefInt", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "10"}
		tc := caseV1Compat{Name: "KVP ValueDefInt", Expected: 10, Actual: kv.ValueDefInt(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_ValueDefInt_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_ValueDefInt_Invalid", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "abc"}
		tc := caseV1Compat{Name: "KVP ValueDefInt invalid", Expected: 0, Actual: kv.ValueDefInt(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_ValueByte_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_ValueByte", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "65"}
		tc := caseV1Compat{Name: "KVP ValueByte", Expected: byte(65), Actual: kv.ValueByte(0), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_ValueByte_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_ValueByte_Invalid", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "abc"}
		tc := caseV1Compat{Name: "KVP ValueByte invalid", Expected: byte(5), Actual: kv.ValueByte(5), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_ValueDefByte", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "10"}
		tc := caseV1Compat{Name: "KVP ValueDefByte", Expected: byte(10), Actual: kv.ValueDefByte(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_ValueFloat64_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_ValueFloat64", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "3.14"}
		tc := caseV1Compat{Name: "KVP ValueFloat64", Expected: 3.14, Actual: kv.ValueFloat64(0), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_ValueFloat64_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_ValueFloat64_Invalid", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "abc"}
		tc := caseV1Compat{Name: "KVP ValueFloat64 invalid", Expected: 1.5, Actual: kv.ValueFloat64(1.5), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_ValueDefFloat64", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "2.5"}
		tc := caseV1Compat{Name: "KVP ValueDefFloat64", Expected: 2.5, Actual: kv.ValueDefFloat64(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_ValueValid_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_ValueValid", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kv.ValueValid()
		tc := caseV1Compat{Name: "KVP ValueValid", Expected: true, Actual: vv.IsValid, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_ValueValidOptions_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_ValueValidOptions", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kv.ValueValidOptions(false, "msg")
		tc := caseV1Compat{Name: "KVP ValueValidOptions", Expected: false, Actual: vv.IsValid, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_Is_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_Is", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KVP Is", Expected: true, Actual: kv.Is("k", "v"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_IsKey(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_IsKey", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KVP IsKey", Expected: true, Actual: kv.IsKey("k"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_IsVal(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_IsVal", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KVP IsVal", Expected: true, Actual: kv.IsVal("v"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_IsKeyValueAnyEmpty_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_IsKeyValueAnyEmpty", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: ""}
		tc := caseV1Compat{Name: "KVP IsKeyValueAnyEmpty", Expected: true, Actual: kv.IsKeyValueAnyEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_IsKeyValueAnyEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_IsKeyValueAnyEmpty_Nil", func() {
		var kv *corestr.KeyValuePair
		tc := caseV1Compat{Name: "KVP IsKeyValueAnyEmpty nil", Expected: true, Actual: kv.IsKeyValueAnyEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_FormatString_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_FormatString", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KVP FormatString", Expected: "k=v", Actual: kv.FormatString("%v=%v"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_String_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_String", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KVP String", Expected: true, Actual: len(kv.String()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_Serialize_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_Serialize", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		data, err := kv.Serialize()
		tc := caseV1Compat{Name: "KVP Serialize", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_SerializeMust_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_SerializeMust", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		data := kv.SerializeMust()
		tc := caseV1Compat{Name: "KVP SerializeMust", Expected: true, Actual: len(data) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_Compile_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_Compile", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KVP Compile", Expected: true, Actual: len(kv.Compile()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_Clear_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_Clear", func() {
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Clear()
		tc := caseV1Compat{Name: "KVP Clear", Expected: "", Actual: kv.Key, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_Clear_Nil", func() {
		var kv *corestr.KeyValuePair
		kv.Clear()
		tc := caseV1Compat{Name: "KVP Clear nil", Expected: true, Actual: true, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_Dispose_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_Dispose", func() {
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Dispose()
		tc := caseV1Compat{Name: "KVP Dispose", Expected: "", Actual: kv.Key, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValuePair_Dispose_Nil_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValuePair_Dispose_Nil", func() {
		var kv *corestr.KeyValuePair
		kv.Dispose()
		tc := caseV1Compat{Name: "KVP Dispose nil", Expected: true, Actual: true, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// KeyValueCollection
// ═══════════════════════════════════════════════════════════════

func Test_KeyValueCollection_Add_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_Add", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k1", "v1")
		tc := caseV1Compat{Name: "KVC Add", Expected: 1, Actual: kvc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AddIf_True(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AddIf_True", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddIf(true, "k", "v")
		tc := caseV1Compat{Name: "KVC AddIf true", Expected: 1, Actual: kvc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AddIf_False(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AddIf_False", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddIf(false, "k", "v")
		tc := caseV1Compat{Name: "KVC AddIf false", Expected: 0, Actual: kvc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_IsEmpty", func() {
		kvc := &corestr.KeyValueCollection{}
		tc := caseV1Compat{Name: "KVC IsEmpty", Expected: true, Actual: kvc.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_HasAnyItem", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		tc := caseV1Compat{Name: "KVC HasAnyItem", Expected: true, Actual: kvc.HasAnyItem(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_Count(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_Count", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		tc := caseV1Compat{Name: "KVC Count", Expected: 1, Actual: kvc.Count(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_LastIndex(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_LastIndex", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		tc := caseV1Compat{Name: "KVC LastIndex", Expected: 1, Actual: kvc.LastIndex(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_HasIndex_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_HasIndex", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		tc := caseV1Compat{Name: "KVC HasIndex", Expected: true, Actual: kvc.HasIndex(0), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_First(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_First", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		tc := caseV1Compat{Name: "KVC First", Expected: "a", Actual: kvc.First().Key, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_FirstOrDefault_Has(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_FirstOrDefault_Has", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		tc := caseV1Compat{Name: "KVC FirstOrDefault", Expected: "a", Actual: kvc.FirstOrDefault().Key, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_FirstOrDefault_Empty_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_FirstOrDefault_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		tc := caseV1Compat{Name: "KVC FirstOrDefault empty", Expected: true, Actual: kvc.FirstOrDefault() == nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_Last(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_Last", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		tc := caseV1Compat{Name: "KVC Last", Expected: "b", Actual: kvc.Last().Key, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_LastOrDefault_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		tc := caseV1Compat{Name: "KVC LastOrDefault empty", Expected: true, Actual: kvc.LastOrDefault() == nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_HasKey(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_HasKey", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		tc := caseV1Compat{Name: "KVC HasKey", Expected: true, Actual: kvc.HasKey("a"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_IsContains_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_IsContains", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		tc := caseV1Compat{Name: "KVC IsContains", Expected: true, Actual: kvc.IsContains("a"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_Get_Found(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_Get_Found", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		val, found := kvc.Get("a")
		tc := caseV1Compat{Name: "KVC Get found", Expected: "1", Actual: val, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "KVC Get found bool", Expected: true, Actual: found, Args: args.Map{}}
		tc2.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_Get_NotFound(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_Get_NotFound", func() {
		kvc := &corestr.KeyValueCollection{}
		_, found := kvc.Get("z")
		tc := caseV1Compat{Name: "KVC Get not found", Expected: false, Actual: found, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AllKeys(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AllKeys", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		tc := caseV1Compat{Name: "KVC AllKeys", Expected: 2, Actual: len(kvc.AllKeys()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AllValues(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AllValues", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		tc := caseV1Compat{Name: "KVC AllValues", Expected: 1, Actual: len(kvc.AllValues()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AllKeysSorted_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AllKeysSorted", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("b", "2")
		kvc.Add("a", "1")
		keys := kvc.AllKeysSorted()
		tc := caseV1Compat{Name: "KVC AllKeysSorted", Expected: "a", Actual: keys[0], Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_Adds_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_Adds", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Adds(corestr.KeyValuePair{Key: "a", Value: "1"}, corestr.KeyValuePair{Key: "b", Value: "2"})
		tc := caseV1Compat{Name: "KVC Adds", Expected: 2, Actual: kvc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_Adds_Empty_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_Adds_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Adds()
		tc := caseV1Compat{Name: "KVC Adds empty", Expected: 0, Actual: kvc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AddMap_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AddMap", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddMap(map[string]string{"a": "1"})
		tc := caseV1Compat{Name: "KVC AddMap", Expected: 1, Actual: kvc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AddMap_Nil_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AddMap_Nil", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddMap(nil)
		tc := caseV1Compat{Name: "KVC AddMap nil", Expected: 0, Actual: kvc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AddHashsetMap_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AddHashsetMap", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashsetMap(map[string]bool{"a": true})
		tc := caseV1Compat{Name: "KVC AddHashsetMap", Expected: 1, Actual: kvc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AddHashsetMap_Nil_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AddHashsetMap_Nil", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashsetMap(nil)
		tc := caseV1Compat{Name: "KVC AddHashsetMap nil", Expected: 0, Actual: kvc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AddHashset_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AddHashset", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashset(corestr.New.Hashset.StringsSpreadItems("a"))
		tc := caseV1Compat{Name: "KVC AddHashset", Expected: 1, Actual: kvc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AddHashset_Nil_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AddHashset_Nil", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashset(nil)
		tc := caseV1Compat{Name: "KVC AddHashset nil", Expected: 0, Actual: kvc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AddsHashmap_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AddsHashmap", func() {
		kvc := &corestr.KeyValueCollection{}
		hm := corestr.New.Hashmap.Cap(2)
		hm.AddOrUpdate("k", "v")
		kvc.AddsHashmap(hm)
		tc := caseV1Compat{Name: "KVC AddsHashmap", Expected: 1, Actual: kvc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AddsHashmap_Nil_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AddsHashmap_Nil", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddsHashmap(nil)
		tc := caseV1Compat{Name: "KVC AddsHashmap nil", Expected: 0, Actual: kvc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_Hashmap_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_Hashmap", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		hm := kvc.Hashmap()
		tc := caseV1Compat{Name: "KVC Hashmap", Expected: true, Actual: hm.Has("k"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_Map_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_Map", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		m := kvc.Map()
		tc := caseV1Compat{Name: "KVC Map", Expected: "v", Actual: m["k"], Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_Join(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_Join", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		tc := caseV1Compat{Name: "KVC Join", Expected: true, Actual: len(kvc.Join(",")) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_JoinKeys(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_JoinKeys", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		tc := caseV1Compat{Name: "KVC JoinKeys", Expected: "a,b", Actual: kvc.JoinKeys(","), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_JoinValues(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_JoinValues", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		tc := caseV1Compat{Name: "KVC JoinValues", Expected: "1", Actual: kvc.JoinValues(","), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_Strings(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_Strings", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		tc := caseV1Compat{Name: "KVC Strings", Expected: 1, Actual: len(kvc.Strings()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_Strings_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		tc := caseV1Compat{Name: "KVC Strings empty", Expected: 0, Actual: len(kvc.Strings()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_StringsUsingFormat_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_StringsUsingFormat", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		result := kvc.StringsUsingFormat("%v=%v")
		tc := caseV1Compat{Name: "KVC StringsUsingFormat", Expected: "k=v", Actual: result[0], Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_String_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_String", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		tc := caseV1Compat{Name: "KVC String", Expected: true, Actual: len(kvc.String()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_SafeValueAt_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_SafeValueAt", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		tc := caseV1Compat{Name: "KVC SafeValueAt", Expected: "v", Actual: kvc.SafeValueAt(0), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_SafeValueAt_OOB(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_SafeValueAt_OOB", func() {
		kvc := &corestr.KeyValueCollection{}
		tc := caseV1Compat{Name: "KVC SafeValueAt oob", Expected: "", Actual: kvc.SafeValueAt(0), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_SafeValuesAtIndexes_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_SafeValuesAtIndexes", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		result := kvc.SafeValuesAtIndexes(0, 1)
		tc := caseV1Compat{Name: "KVC SafeValuesAtIndexes", Expected: 2, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_Serialize_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_Serialize", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		data, err := kvc.Serialize()
		tc := caseV1Compat{Name: "KVC Serialize", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_SerializeMust_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_SerializeMust", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		data := kvc.SerializeMust()
		tc := caseV1Compat{Name: "KVC SerializeMust", Expected: true, Actual: len(data) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AddStringBySplit_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AddStringBySplit", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddStringBySplit("=", "key=value")
		tc := caseV1Compat{Name: "KVC AddStringBySplit", Expected: 1, Actual: kvc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AddStringBySplitTrim_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AddStringBySplitTrim", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddStringBySplitTrim("=", " key = value ")
		tc := caseV1Compat{Name: "KVC AddStringBySplitTrim", Expected: "key", Actual: kvc.First().Key, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_Find_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_Find", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		result := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, kv.Key == "a", false
		})
		tc := caseV1Compat{Name: "KVC Find", Expected: 1, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_Find_Empty(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_Find_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		result := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, false
		})
		tc := caseV1Compat{Name: "KVC Find empty", Expected: 0, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AddsHashmaps_FromKeyValuePairKeyNameI(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AddsHashmaps", func() {
		kvc := &corestr.KeyValueCollection{}
		hm := corestr.New.Hashmap.Cap(2)
		hm.AddOrUpdate("k", "v")
		kvc.AddsHashmaps(hm)
		tc := caseV1Compat{Name: "KVC AddsHashmaps", Expected: 1, Actual: kvc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyValueCollection_AddsHashmaps_Empty(t *testing.T) {
	safeTest(t, "Test_Cov50_KeyValueCollection_AddsHashmaps_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddsHashmaps()
		tc := caseV1Compat{Name: "KVC AddsHashmaps empty", Expected: 0, Actual: kvc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}
