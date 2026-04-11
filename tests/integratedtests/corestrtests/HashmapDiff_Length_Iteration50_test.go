package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════
// HashmapDiff
// ═══════════════════════════════════════════════════════════════

func Test_HashmapDiff_Length(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_Length", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		tc := caseV1Compat{Name: "HD Length", Expected: 1, Actual: hd.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_IsEmpty", func() {
		hd := corestr.HashmapDiff(map[string]string{})
		tc := caseV1Compat{Name: "HD IsEmpty", Expected: true, Actual: hd.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_HasAnyItem", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		tc := caseV1Compat{Name: "HD HasAnyItem", Expected: true, Actual: hd.HasAnyItem(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_LastIndex(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_LastIndex", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1", "b": "2"})
		tc := caseV1Compat{Name: "HD LastIndex", Expected: 1, Actual: hd.LastIndex(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_Raw(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_Raw", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		tc := caseV1Compat{Name: "HD Raw", Expected: "1", Actual: hd.Raw()["a"], Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_Raw_Nil(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_Raw_Nil", func() {
		var hd *corestr.HashmapDiff
		tc := caseV1Compat{Name: "HD Raw nil", Expected: 0, Actual: len(hd.Raw()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_AllKeysSorted_FromHashmapDiffLengthIte(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_AllKeysSorted", func() {
		hd := corestr.HashmapDiff(map[string]string{"b": "2", "a": "1"})
		keys := hd.AllKeysSorted()
		tc := caseV1Compat{Name: "HD AllKeysSorted", Expected: "a", Actual: keys[0], Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_MapAnyItems_FromHashmapDiffLengthIte(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_MapAnyItems", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		m := hd.MapAnyItems()
		tc := caseV1Compat{Name: "HD MapAnyItems", Expected: "1", Actual: m["a"], Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_MapAnyItems_Nil_FromHashmapDiffLengthIte(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_MapAnyItems_Nil", func() {
		var hd *corestr.HashmapDiff
		m := hd.MapAnyItems()
		tc := caseV1Compat{Name: "HD MapAnyItems nil", Expected: 0, Actual: len(m), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_IsRawEqual(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_IsRawEqual", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		tc := caseV1Compat{Name: "HD IsRawEqual", Expected: true, Actual: hd.IsRawEqual(map[string]string{"a": "1"}), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_IsRawEqual_Diff(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_IsRawEqual_Diff", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		tc := caseV1Compat{Name: "HD IsRawEqual diff", Expected: false, Actual: hd.IsRawEqual(map[string]string{"a": "2"}), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_HasAnyChanges(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_HasAnyChanges", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		tc := caseV1Compat{Name: "HD HasAnyChanges", Expected: true, Actual: hd.HasAnyChanges(map[string]string{"a": "2"}), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_DiffRaw(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_DiffRaw", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		diff := hd.DiffRaw(map[string]string{"a": "2"})
		tc := caseV1Compat{Name: "HD DiffRaw", Expected: true, Actual: len(diff) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_HashmapDiffUsingRaw(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_HashmapDiffUsingRaw", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		result := hd.HashmapDiffUsingRaw(map[string]string{"a": "2"})
		tc := caseV1Compat{Name: "HD HashmapDiffUsingRaw", Expected: true, Actual: result.HasAnyItem(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_HashmapDiffUsingRaw_Same(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_HashmapDiffUsingRaw_Same", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		result := hd.HashmapDiffUsingRaw(map[string]string{"a": "1"})
		tc := caseV1Compat{Name: "HD HashmapDiffUsingRaw same", Expected: true, Actual: result.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_DiffJsonMessage(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_DiffJsonMessage", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		msg := hd.DiffJsonMessage(map[string]string{"a": "2"})
		tc := caseV1Compat{Name: "HD DiffJsonMessage", Expected: true, Actual: len(msg) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_ToStringsSliceOfDiffMap(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_ToStringsSliceOfDiffMap", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		diff := hd.DiffRaw(map[string]string{"a": "2"})
		slice := hd.ToStringsSliceOfDiffMap(diff)
		tc := caseV1Compat{Name: "HD ToStringsSlice", Expected: true, Actual: len(slice) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_ShouldDiffMessage(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_ShouldDiffMessage", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		msg := hd.ShouldDiffMessage("test", map[string]string{"a": "2"})
		tc := caseV1Compat{Name: "HD ShouldDiffMessage", Expected: true, Actual: len(msg) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_Serialize", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		data, err := hd.Serialize()
		tc := caseV1Compat{Name: "HD Serialize", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_HashmapDiff_RawMapStringAnyDiff(t *testing.T) {
	safeTest(t, "Test_Cov54_HashmapDiff_RawMapStringAnyDiff", func() {
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		m := hd.RawMapStringAnyDiff()
		tc := caseV1Compat{Name: "HD RawMapStringAnyDiff", Expected: true, Actual: len(m) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// KeyAnyValuePair
// ═══════════════════════════════════════════════════════════════

func Test_KeyAnyValuePair_KeyName(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_KeyName", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KAV KeyName", Expected: "k", Actual: kav.KeyName(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_VariableName(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_VariableName", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KAV VariableName", Expected: "k", Actual: kav.VariableName(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_ValueAny(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_ValueAny", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
		tc := caseV1Compat{Name: "KAV ValueAny", Expected: 42, Actual: kav.ValueAny(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_IsVariableNameEqual(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_IsVariableNameEqual", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KAV IsVarNameEqual", Expected: true, Actual: kav.IsVariableNameEqual("k"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_IsValueNull_FromHashmapDiffLengthIte(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_IsValueNull", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: nil}
		tc := caseV1Compat{Name: "KAV IsValueNull", Expected: true, Actual: kav.IsValueNull(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_HasNonNull(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_HasNonNull", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KAV HasNonNull", Expected: true, Actual: kav.HasNonNull(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_HasValue(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_HasValue", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KAV HasValue", Expected: true, Actual: kav.HasValue(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_IsValueEmptyString(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_IsValueEmptyString", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: nil}
		tc := caseV1Compat{Name: "KAV IsValueEmptyString", Expected: true, Actual: kav.IsValueEmptyString(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_ValueString_FromHashmapDiffLengthIte(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_ValueString", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "hello"}
		tc := caseV1Compat{Name: "KAV ValueString", Expected: "hello", Actual: kav.ValueString(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_String(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_String", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KAV String", Expected: true, Actual: len(kav.String()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_Compile_FromHashmapDiffLengthIte(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_Compile", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KAV Compile", Expected: true, Actual: len(kav.Compile()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_Serialize", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		data, err := kav.Serialize()
		tc := caseV1Compat{Name: "KAV Serialize", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_SerializeMust(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_SerializeMust", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		data := kav.SerializeMust()
		tc := caseV1Compat{Name: "KAV SerializeMust", Expected: true, Actual: len(data) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_Json(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_Json", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		j := kav.Json()
		tc := caseV1Compat{Name: "KAV Json", Expected: true, Actual: j.HasSafeItems(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_AsJsoner", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KAV AsJsoner", Expected: true, Actual: kav.AsJsoner() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_AsJsonContractsBinder", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		tc := caseV1Compat{Name: "KAV AsJsonContractsBinder", Expected: true, Actual: kav.AsJsonContractsBinder() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_Clear(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_Clear", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kav.Clear()
		tc := caseV1Compat{Name: "KAV Clear", Expected: "", Actual: kav.Key, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_KeyAnyValuePair_Dispose_FromHashmapDiffLengthIte(t *testing.T) {
	safeTest(t, "Test_Cov54_KeyAnyValuePair_Dispose", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kav.Dispose()
		tc := caseV1Compat{Name: "KAV Dispose", Expected: "", Actual: kav.Key, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// CollectionsOfCollection
// ═══════════════════════════════════════════════════════════════

func Test_CollOfColl_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov54_CollOfColl_IsEmpty", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		tc := caseV1Compat{Name: "CoC IsEmpty", Expected: true, Actual: coc.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_CollOfColl_HasItems(t *testing.T) {
	safeTest(t, "Test_Cov54_CollOfColl_HasItems", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		coc.Adds(*col)
		tc := caseV1Compat{Name: "CoC HasItems", Expected: true, Actual: coc.HasItems(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_CollOfColl_Length(t *testing.T) {
	safeTest(t, "Test_Cov54_CollOfColl_Length", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		col := corestr.New.Collection.Strings([]string{"a"})
		coc.Adds(*col)
		tc := caseV1Compat{Name: "CoC Length", Expected: 1, Actual: coc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_CollOfColl_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_Cov54_CollOfColl_AllIndividualItemsLength", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		coc.Adds(*col)
		tc := caseV1Compat{Name: "CoC AllIndivLen", Expected: 2, Actual: coc.AllIndividualItemsLength(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_CollOfColl_Items(t *testing.T) {
	safeTest(t, "Test_Cov54_CollOfColl_Items", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		col := corestr.New.Collection.Strings([]string{"a"})
		coc.Adds(*col)
		tc := caseV1Compat{Name: "CoC Items", Expected: 1, Actual: len(coc.Items()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_CollOfColl_List(t *testing.T) {
	safeTest(t, "Test_Cov54_CollOfColl_List", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		coc.Adds(*col)
		tc := caseV1Compat{Name: "CoC List", Expected: 2, Actual: len(coc.List(0)), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_CollOfColl_ToCollection(t *testing.T) {
	safeTest(t, "Test_Cov54_CollOfColl_ToCollection", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		col := corestr.New.Collection.Strings([]string{"a"})
		coc.Adds(*col)
		result := coc.ToCollection()
		tc := caseV1Compat{Name: "CoC ToCollection", Expected: 1, Actual: result.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_CollOfColl_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov54_CollOfColl_AddStrings", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"x", "y"})
		tc := caseV1Compat{Name: "CoC AddStrings", Expected: 1, Actual: coc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_CollOfColl_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov54_CollOfColl_AddStrings_Empty", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{})
		tc := caseV1Compat{Name: "CoC AddStrings empty", Expected: 0, Actual: coc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_CollOfColl_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov54_CollOfColl_AsJsonContractsBinder", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		tc := caseV1Compat{Name: "CoC AsJsonContractsBinder", Expected: true, Actual: coc.AsJsonContractsBinder() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// SimpleStringOnce (core methods)
// ═══════════════════════════════════════════════════════════════

func Test_SSO_GetSetOnce(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_GetSetOnce", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		tc := caseV1Compat{Name: "SSO Value", Expected: "hello", Actual: sso.Value(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsInitialized(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_IsInitialized", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		tc := caseV1Compat{Name: "SSO IsInitialized", Expected: true, Actual: sso.IsInitialized(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsDefined(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_IsDefined", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		tc := caseV1Compat{Name: "SSO IsDefined", Expected: true, Actual: sso.IsDefined(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsUninitialized(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_IsUninitialized", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		tc := caseV1Compat{Name: "SSO IsUninitialized", Expected: true, Actual: sso.IsUninitialized(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Invalidate(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_Invalidate", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		sso.Invalidate()
		tc := caseV1Compat{Name: "SSO Invalidate", Expected: true, Actual: sso.IsUninitialized(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Reset(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_Reset", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		sso.Reset()
		tc := caseV1Compat{Name: "SSO Reset", Expected: true, Actual: sso.IsUninitialized(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsInvalid(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_IsInvalid", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		tc := caseV1Compat{Name: "SSO IsInvalid", Expected: true, Actual: sso.IsInvalid(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_ValueBytes(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_ValueBytes", func() {
		sso := corestr.New.SimpleStringOnce.Init("ab")
		tc := caseV1Compat{Name: "SSO ValueBytes", Expected: 2, Actual: len(sso.ValueBytes()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_SetOnUninitialized(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_SetOnUninitialized", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		err := sso.SetOnUninitialized("x")
		tc := caseV1Compat{Name: "SSO SetOnUninitialized", Expected: true, Actual: err == nil && sso.Value() == "x", Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_SetOnUninitialized_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_SetOnUninitialized_AlreadyInit", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		err := sso.SetOnUninitialized("y")
		tc := caseV1Compat{Name: "SSO SetOnUninitialized already", Expected: true, Actual: err != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_GetOnce(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_GetOnce", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		val := sso.GetOnce()
		tc := caseV1Compat{Name: "SSO GetOnce", Expected: "", Actual: val, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_GetOnceFunc(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_GetOnceFunc", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		val := sso.GetOnceFunc(func() string { return "computed" })
		tc := caseV1Compat{Name: "SSO GetOnceFunc", Expected: "computed", Actual: val, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_SetOnceIfUninitialized(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_SetOnceIfUninitialized", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		isSet := sso.SetOnceIfUninitialized("x")
		tc := caseV1Compat{Name: "SSO SetOnceIfUninitialized", Expected: true, Actual: isSet, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_IsEmpty", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		tc := caseV1Compat{Name: "SSO IsEmpty", Expected: true, Actual: sso.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Is(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_Is", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		tc := caseV1Compat{Name: "SSO Is", Expected: true, Actual: sso.Is("x"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_IsContains", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello world")
		tc := caseV1Compat{Name: "SSO IsContains", Expected: true, Actual: sso.IsContains("world"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Int(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_Int", func() {
		sso := corestr.New.SimpleStringOnce.Init("42")
		tc := caseV1Compat{Name: "SSO Int", Expected: 42, Actual: sso.Int(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Byte(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_Byte", func() {
		sso := corestr.New.SimpleStringOnce.Init("65")
		tc := caseV1Compat{Name: "SSO Byte", Expected: byte(65), Actual: sso.Byte(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Boolean(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_Boolean", func() {
		sso := corestr.New.SimpleStringOnce.Init("yes")
		tc := caseV1Compat{Name: "SSO Boolean", Expected: true, Actual: sso.Boolean(false), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_BooleanDefault(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_BooleanDefault", func() {
		sso := corestr.New.SimpleStringOnce.Init("true")
		tc := caseV1Compat{Name: "SSO BooleanDefault", Expected: true, Actual: sso.BooleanDefault(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_ConcatNew", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		result := sso.ConcatNew(" world")
		tc := caseV1Compat{Name: "SSO ConcatNew", Expected: "hello world", Actual: result.Value(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_HasSafeNonEmpty", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		tc := caseV1Compat{Name: "SSO HasSafeNonEmpty", Expected: true, Actual: sso.HasSafeNonEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_SafeValue(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_SafeValue", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		tc := caseV1Compat{Name: "SSO SafeValue", Expected: "x", Actual: sso.SafeValue(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_SafeValue_Uninit(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_SafeValue_Uninit", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		tc := caseV1Compat{Name: "SSO SafeValue uninit", Expected: "", Actual: sso.SafeValue(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_ValueInt(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_ValueInt", func() {
		sso := corestr.New.SimpleStringOnce.Init("10")
		tc := caseV1Compat{Name: "SSO ValueInt", Expected: 10, Actual: sso.ValueInt(0), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_ValueFloat64", func() {
		sso := corestr.New.SimpleStringOnce.Init("3.14")
		tc := caseV1Compat{Name: "SSO ValueFloat64", Expected: 3.14, Actual: sso.ValueFloat64(0), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_WithinRange_InRange(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_WithinRange_InRange", func() {
		sso := corestr.New.SimpleStringOnce.Init("5")
		val, inRange := sso.WithinRange(true, 0, 10)
		tc := caseV1Compat{Name: "SSO WithinRange in", Expected: true, Actual: inRange && val == 5, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_WithinRange_OutOfRange(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_WithinRange_OutOfRange", func() {
		sso := corestr.New.SimpleStringOnce.Init("20")
		val, inRange := sso.WithinRange(true, 0, 10)
		tc := caseV1Compat{Name: "SSO WithinRange out", Expected: true, Actual: !inRange && val == 10, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Int16(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_Int16", func() {
		sso := corestr.New.SimpleStringOnce.Init("100")
		tc := caseV1Compat{Name: "SSO Int16", Expected: int16(100), Actual: sso.Int16(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Int32(t *testing.T) {
	safeTest(t, "Test_Cov54_SSO_Int32", func() {
		sso := corestr.New.SimpleStringOnce.Init("100")
		tc := caseV1Compat{Name: "SSO Int32", Expected: int32(100), Actual: sso.Int32(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}
