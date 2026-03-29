package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Segment 11: Remaining methods (L700-1300)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovHM2_01_Items_SafeItems(t *testing.T) {
	safeTest(t, "Test_CovHM2_01_Items_SafeItems", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		if len(hm.Items()) != 1 {
			t.Fatal("expected 1")
		}
		if len(hm.SafeItems()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHM2_02_ItemsCopyLock(t *testing.T) {
	safeTest(t, "Test_CovHM2_02_ItemsCopyLock", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		cp := hm.ItemsCopyLock()
		if len(*cp) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHM2_03_ValuesCollection_ValuesHashset(t *testing.T) {
	safeTest(t, "Test_CovHM2_03_ValuesCollection_ValuesHashset", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "val1")
		col := hm.ValuesCollection()
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
		hs := hm.ValuesHashset()
		if hs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHM2_04_ValuesCollectionLock_ValuesHashsetLock(t *testing.T) {
	safeTest(t, "Test_CovHM2_04_ValuesCollectionLock_ValuesHashsetLock", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		_ = hm.ValuesCollectionLock()
		_ = hm.ValuesHashsetLock()
	})
}

func Test_CovHM2_05_ValuesList_ValuesListCopyLock(t *testing.T) {
	safeTest(t, "Test_CovHM2_05_ValuesList_ValuesListCopyLock", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		if len(hm.ValuesList()) != 1 {
			t.Fatal("expected 1")
		}
		if len(hm.ValuesListCopyLock()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHM2_06_KeysValuesCollection(t *testing.T) {
	safeTest(t, "Test_CovHM2_06_KeysValuesCollection", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		keys, vals := hm.KeysValuesCollection()
		if keys.Length() != 1 || vals.Length() != 1 {
			t.Fatal("expected 1 each")
		}
	})
}

func Test_CovHM2_07_KeysValuesList(t *testing.T) {
	safeTest(t, "Test_CovHM2_07_KeysValuesList", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		keys, vals := hm.KeysValuesList()
		if len(keys) != 1 || len(vals) != 1 {
			t.Fatal("expected 1 each")
		}
	})
}

func Test_CovHM2_08_KeysValuePairs(t *testing.T) {
	safeTest(t, "Test_CovHM2_08_KeysValuePairs", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		pairs := hm.KeysValuePairs()
		if len(pairs) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHM2_09_KeysValuePairsCollection(t *testing.T) {
	safeTest(t, "Test_CovHM2_09_KeysValuePairsCollection", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		kvc := hm.KeysValuePairsCollection()
		if kvc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHM2_10_KeysValuesListLock(t *testing.T) {
	safeTest(t, "Test_CovHM2_10_KeysValuesListLock", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		keys, vals := hm.KeysValuesListLock()
		if len(keys) != 1 || len(vals) != 1 {
			t.Fatal("expected 1 each")
		}
	})
}

func Test_CovHM2_11_AllKeys_Keys_KeysCollection_KeysLock(t *testing.T) {
	safeTest(t, "Test_CovHM2_11_AllKeys_Keys_KeysCollection_KeysLock", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		if len(hm.AllKeys()) != 1 {
			t.Fatal("expected 1")
		}
		if len(hm.Keys()) != 1 {
			t.Fatal("expected 1")
		}
		if hm.KeysCollection().Length() != 1 {
			t.Fatal("expected 1")
		}
		if len(hm.KeysLock()) != 1 {
			t.Fatal("expected 1")
		}
		// empty keys
		e := corestr.Empty.Hashmap()
		if len(e.AllKeys()) != 0 {
			t.Fatal("expected 0")
		}
		if len(e.KeysLock()) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovHM2_12_KeysToLower_ValuesToLower(t *testing.T) {
	safeTest(t, "Test_CovHM2_12_KeysToLower_ValuesToLower", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("ABC", "val")
		lower := hm.KeysToLower()
		if !lower.Has("abc") {
			t.Fatal("expected abc")
		}
		_ = hm.ValuesToLower()
	})
}

func Test_CovHM2_13_Length_LengthLock(t *testing.T) {
	safeTest(t, "Test_CovHM2_13_Length_LengthLock", func() {
		hm := corestr.Empty.Hashmap()
		if hm.Length() != 0 {
			t.Fatal("expected 0")
		}
		if hm.LengthLock() != 0 {
			t.Fatal("expected 0")
		}
		hm.AddOrUpdate("a", "1")
		if hm.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHM2_14_IsEqual_IsEqualPtr_IsEqualPtrLock(t *testing.T) {
	safeTest(t, "Test_CovHM2_14_IsEqual_IsEqualPtr_IsEqualPtrLock", func() {
		a := corestr.Empty.Hashmap()
		a.AddOrUpdate("a", "1")
		b := corestr.Empty.Hashmap()
		b.AddOrUpdate("a", "1")
		if !a.IsEqualPtr(b) {
			t.Fatal("expected equal")
		}
		if !a.IsEqualPtr(a) {
			t.Fatal("expected equal to self")
		}
		if a.IsEqualPtr(nil) {
			t.Fatal("expected false")
		}
		// both empty
		e1 := corestr.Empty.Hashmap()
		e2 := corestr.Empty.Hashmap()
		if !e1.IsEqualPtr(e2) {
			t.Fatal("expected true")
		}
		// one empty
		if a.IsEqualPtr(e1) {
			t.Fatal("expected false")
		}
		// diff length
		c := corestr.Empty.Hashmap()
		c.AddOrUpdate("a", "1")
		c.AddOrUpdate("b", "2")
		if a.IsEqualPtr(c) {
			t.Fatal("expected false")
		}
		// diff value
		d := corestr.Empty.Hashmap()
		d.AddOrUpdate("a", "99")
		if a.IsEqualPtr(d) {
			t.Fatal("expected false")
		}
		// IsEqual (value receiver)
		if !a.IsEqual(*b) {
			t.Fatal("expected equal")
		}
		// IsEqualPtrLock
		if !a.IsEqualPtrLock(b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_CovHM2_15_Remove_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_CovHM2_15_Remove_RemoveWithLock", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		hm.Remove("a")
		if hm.Has("a") {
			t.Fatal("expected removed")
		}
		hm.AddOrUpdate("b", "2")
		hm.RemoveWithLock("b")
		if hm.Has("b") {
			t.Fatal("expected removed")
		}
	})
}

func Test_CovHM2_16_String_StringLock(t *testing.T) {
	safeTest(t, "Test_CovHM2_16_String_StringLock", func() {
		hm := corestr.Empty.Hashmap()
		s := hm.String()
		if s == "" {
			t.Fatal("expected non-empty")
		}
		hm.AddOrUpdate("a", "1")
		_ = hm.String()
		_ = hm.StringLock()
		_ = corestr.Empty.Hashmap().StringLock()
	})
}

func Test_CovHM2_17_GetValuesExceptKeysInHashset(t *testing.T) {
	safeTest(t, "Test_CovHM2_17_GetValuesExceptKeysInHashset", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		except := corestr.New.Hashset.Empty()
		except.Add("a")
		r := hm.GetValuesExceptKeysInHashset(except)
		if len(r) != 1 {
			t.Fatal("expected 1")
		}
		// nil
		r2 := hm.GetValuesExceptKeysInHashset(nil)
		if len(r2) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovHM2_18_GetValuesKeysExcept(t *testing.T) {
	safeTest(t, "Test_CovHM2_18_GetValuesKeysExcept", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		r := hm.GetValuesKeysExcept([]string{"a"})
		if len(r) != 0 {
			t.Fatal("expected 0")
		}
		r2 := hm.GetValuesKeysExcept(nil)
		if len(r2) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHM2_19_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_CovHM2_19_GetAllExceptCollection", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		r := hm.GetAllExceptCollection(corestr.New.Collection.Strings([]string{"a"}))
		if len(r) != 0 {
			t.Fatal("expected 0")
		}
		r2 := hm.GetAllExceptCollection(nil)
		if len(r2) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHM2_20_Join_JoinKeys(t *testing.T) {
	safeTest(t, "Test_CovHM2_20_Join_JoinKeys", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		_ = hm.Join(",")
		_ = hm.JoinKeys(",")
	})
}

func Test_CovHM2_21_JsonModel_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_CovHM2_21_JsonModel_JsonModelAny", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		if len(hm.JsonModel()) != 1 {
			t.Fatal("expected 1")
		}
		_ = hm.JsonModelAny()
	})
}

func Test_CovHM2_22_MarshalJSON_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovHM2_22_MarshalJSON_UnmarshalJSON", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		data, err := hm.MarshalJSON()
		if err != nil {
			t.Fatal("unexpected error")
		}
		hm2 := corestr.Empty.Hashmap()
		err2 := hm2.UnmarshalJSON(data)
		if err2 != nil {
			t.Fatal("unexpected error")
		}
		// invalid
		err3 := hm2.UnmarshalJSON([]byte("bad"))
		if err3 == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_CovHM2_23_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_CovHM2_23_Json_JsonPtr", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		_ = hm.Json()
		_ = hm.JsonPtr()
	})
}

func Test_CovHM2_24_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_CovHM2_24_ParseInjectUsingJson", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		jr := hm.JsonPtr()
		hm2 := corestr.Empty.Hashmap()
		r, err := hm2.ParseInjectUsingJson(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
		if r.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHM2_25_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovHM2_25_ParseInjectUsingJsonMust", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		jr := hm.JsonPtr()
		hm2 := corestr.Empty.Hashmap()
		r := hm2.ParseInjectUsingJsonMust(jr)
		if r.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHM2_26_ToError_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_CovHM2_26_ToError_ToDefaultError", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("key", "val")
		e := hm.ToError(",")
		if e == nil {
			t.Fatal("expected error")
		}
		e2 := hm.ToDefaultError()
		if e2 == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_CovHM2_27_KeyValStringLines(t *testing.T) {
	safeTest(t, "Test_CovHM2_27_KeyValStringLines", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		lines := hm.KeyValStringLines()
		if len(lines) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHM2_28_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_CovHM2_28_Clear_Dispose", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		hm.Clear()
		if hm.Length() != 0 {
			t.Fatal("expected 0")
		}
		hm2 := corestr.Empty.Hashmap()
		hm2.Dispose()
	})
}

func Test_CovHM2_29_ToStringsUsingCompiler(t *testing.T) {
	safeTest(t, "Test_CovHM2_29_ToStringsUsingCompiler", func() {
		hm := corestr.Empty.Hashmap()
		// empty
		r := hm.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })
		if len(r) != 0 {
			t.Fatal("expected 0")
		}
		hm.AddOrUpdate("a", "1")
		r2 := hm.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })
		if len(r2) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHM2_30_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_CovHM2_30_AsInterfaces", func() {
		hm := corestr.Empty.Hashmap()
		_ = hm.AsJsoner()
		_ = hm.AsJsonContractsBinder()
		_ = hm.AsJsonParseSelfInjector()
		_ = hm.AsJsonMarshaller()
	})
}

func Test_CovHM2_31_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovHM2_31_JsonParseSelfInject", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		jr := hm.JsonPtr()
		hm2 := corestr.Empty.Hashmap()
		err := hm2.JsonParseSelfInject(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovHM2_32_Clone_ClonePtr(t *testing.T) {
	safeTest(t, "Test_CovHM2_32_Clone_ClonePtr", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		c := hm.Clone()
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
		cp := hm.ClonePtr()
		if cp.Length() != 1 {
			t.Fatal("expected 1")
		}
		// empty clone
		e := corestr.Empty.Hashmap()
		ec := e.Clone()
		if ec.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovHM2_33_Get_GetValue(t *testing.T) {
	safeTest(t, "Test_CovHM2_33_Get_GetValue", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		v, ok := hm.Get("a")
		if !ok || v != "1" {
			t.Fatal("expected found")
		}
		_, ok2 := hm.Get("z")
		if ok2 {
			t.Fatal("expected not found")
		}
		v2, _ := hm.GetValue("a")
		if v2 != "1" {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHM2_34_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_CovHM2_34_Serialize_Deserialize", func() {
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		_, err := hm.Serialize()
		if err != nil {
			t.Fatal("unexpected error")
		}
		target := corestr.Empty.Hashmap()
		err2 := hm.Deserialize(target)
		if err2 != nil {
			t.Fatal("unexpected error")
		}
	})
}
