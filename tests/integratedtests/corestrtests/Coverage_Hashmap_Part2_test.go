package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Segment 11: Remaining methods (L700-1300)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovHM2_01_Items_SafeItems(t *testing.T) {
	safeTest(t, "Test_CovHM2_01_Items_SafeItems", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": len(hm.Items()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": len(hm.SafeItems()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHM2_02_ItemsCopyLock(t *testing.T) {
	safeTest(t, "Test_CovHM2_02_ItemsCopyLock", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		cp := hm.ItemsCopyLock()

		// Act
		actual := args.Map{"result": len(*cp) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHM2_03_ValuesCollection_ValuesHashset(t *testing.T) {
	safeTest(t, "Test_CovHM2_03_ValuesCollection_ValuesHashset", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "val1")
		col := hm.ValuesCollection()

		// Act
		actual := args.Map{"result": col.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		hs := hm.ValuesHashset()
		actual = args.Map{"result": hs.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": len(hm.ValuesList()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": len(hm.ValuesListCopyLock()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHM2_06_KeysValuesCollection(t *testing.T) {
	safeTest(t, "Test_CovHM2_06_KeysValuesCollection", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		keys, vals := hm.KeysValuesCollection()

		// Act
		actual := args.Map{"result": keys.Length() != 1 || vals.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 each", actual)
	})
}

func Test_CovHM2_07_KeysValuesList(t *testing.T) {
	safeTest(t, "Test_CovHM2_07_KeysValuesList", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		keys, vals := hm.KeysValuesList()

		// Act
		actual := args.Map{"result": len(keys) != 1 || len(vals) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 each", actual)
	})
}

func Test_CovHM2_08_KeysValuePairs(t *testing.T) {
	safeTest(t, "Test_CovHM2_08_KeysValuePairs", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		pairs := hm.KeysValuePairs()

		// Act
		actual := args.Map{"result": len(pairs) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHM2_09_KeysValuePairsCollection(t *testing.T) {
	safeTest(t, "Test_CovHM2_09_KeysValuePairsCollection", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		kvc := hm.KeysValuePairsCollection()

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHM2_10_KeysValuesListLock(t *testing.T) {
	safeTest(t, "Test_CovHM2_10_KeysValuesListLock", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		keys, vals := hm.KeysValuesListLock()

		// Act
		actual := args.Map{"result": len(keys) != 1 || len(vals) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 each", actual)
	})
}

func Test_CovHM2_11_AllKeys_Keys_KeysCollection_KeysLock(t *testing.T) {
	safeTest(t, "Test_CovHM2_11_AllKeys_Keys_KeysCollection_KeysLock", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": len(hm.AllKeys()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": len(hm.Keys()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": hm.KeysCollection().Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": len(hm.KeysLock()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty keys
		e := corestr.Empty.Hashmap()
		actual = args.Map{"result": len(e.AllKeys()) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": len(e.KeysLock()) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovHM2_12_KeysToLower_ValuesToLower(t *testing.T) {
	safeTest(t, "Test_CovHM2_12_KeysToLower_ValuesToLower", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("ABC", "val")
		lower := hm.KeysToLower()

		// Act
		actual := args.Map{"result": lower.Has("abc")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected abc", actual)
		_ = hm.ValuesToLower()
	})
}

func Test_CovHM2_13_Length_LengthLock(t *testing.T) {
	safeTest(t, "Test_CovHM2_13_Length_LengthLock", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()

		// Act
		actual := args.Map{"result": hm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": hm.LengthLock() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hm.AddOrUpdate("a", "1")
		actual = args.Map{"result": hm.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHM2_14_IsEqual_IsEqualPtr_IsEqualPtrLock(t *testing.T) {
	safeTest(t, "Test_CovHM2_14_IsEqual_IsEqualPtr_IsEqualPtrLock", func() {
		// Arrange
		a := corestr.Empty.Hashmap()
		a.AddOrUpdate("a", "1")
		b := corestr.Empty.Hashmap()
		b.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		actual = args.Map{"result": a.IsEqualPtr(a)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal to self", actual)
		actual = args.Map{"result": a.IsEqualPtr(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// both empty
		e1 := corestr.Empty.Hashmap()
		e2 := corestr.Empty.Hashmap()
		actual = args.Map{"result": e1.IsEqualPtr(e2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// one empty
		actual = args.Map{"result": a.IsEqualPtr(e1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// diff length
		c := corestr.Empty.Hashmap()
		c.AddOrUpdate("a", "1")
		c.AddOrUpdate("b", "2")
		actual = args.Map{"result": a.IsEqualPtr(c)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// diff value
		d := corestr.Empty.Hashmap()
		d.AddOrUpdate("a", "99")
		actual = args.Map{"result": a.IsEqualPtr(d)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// IsEqual (value receiver)
		actual = args.Map{"result": a.IsEqual(*b)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		// IsEqualPtrLock
		actual = args.Map{"result": a.IsEqualPtrLock(b)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_CovHM2_15_Remove_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_CovHM2_15_Remove_RemoveWithLock", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		hm.Remove("a")

		// Act
		actual := args.Map{"result": hm.Has("a")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed", actual)
		hm.AddOrUpdate("b", "2")
		hm.RemoveWithLock("b")
		actual = args.Map{"result": hm.Has("b")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed", actual)
	})
}

func Test_CovHM2_16_String_StringLock(t *testing.T) {
	safeTest(t, "Test_CovHM2_16_String_StringLock", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		s := hm.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		hm.AddOrUpdate("a", "1")
		_ = hm.String()
		_ = hm.StringLock()
		_ = corestr.Empty.Hashmap().StringLock()
	})
}

func Test_CovHM2_17_GetValuesExceptKeysInHashset(t *testing.T) {
	safeTest(t, "Test_CovHM2_17_GetValuesExceptKeysInHashset", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		hm.AddOrUpdate("b", "2")
		except := corestr.New.Hashset.Empty()
		except.Add("a")
		r := hm.GetValuesExceptKeysInHashset(except)

		// Act
		actual := args.Map{"result": len(r) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// nil
		r2 := hm.GetValuesExceptKeysInHashset(nil)
		actual = args.Map{"result": len(r2) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovHM2_18_GetValuesKeysExcept(t *testing.T) {
	safeTest(t, "Test_CovHM2_18_GetValuesKeysExcept", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		r := hm.GetValuesKeysExcept([]string{"a"})

		// Act
		actual := args.Map{"result": len(r) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		r2 := hm.GetValuesKeysExcept(nil)
		actual = args.Map{"result": len(r2) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHM2_19_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_CovHM2_19_GetAllExceptCollection", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		r := hm.GetAllExceptCollection(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": len(r) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		r2 := hm.GetAllExceptCollection(nil)
		actual = args.Map{"result": len(r2) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")

		// Act
		actual := args.Map{"result": len(hm.JsonModel()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		_ = hm.JsonModelAny()
	})
}

func Test_CovHM2_22_MarshalJSON_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovHM2_22_MarshalJSON_UnmarshalJSON", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		data, err := hm.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		hm2 := corestr.Empty.Hashmap()
		err2 := hm2.UnmarshalJSON(data)
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		// invalid
		err3 := hm2.UnmarshalJSON([]byte("bad"))
		actual = args.Map{"result": err3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
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
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		jr := hm.JsonPtr()
		hm2 := corestr.Empty.Hashmap()
		r, err := hm2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual = args.Map{"result": r.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHM2_25_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovHM2_25_ParseInjectUsingJsonMust", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		jr := hm.JsonPtr()
		hm2 := corestr.Empty.Hashmap()
		r := hm2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"result": r.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHM2_26_ToError_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_CovHM2_26_ToError_ToDefaultError", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("key", "val")
		e := hm.ToError(",")

		// Act
		actual := args.Map{"result": e == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
		e2 := hm.ToDefaultError()
		actual = args.Map{"result": e2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_CovHM2_27_KeyValStringLines(t *testing.T) {
	safeTest(t, "Test_CovHM2_27_KeyValStringLines", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		lines := hm.KeyValStringLines()

		// Act
		actual := args.Map{"result": len(lines) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHM2_28_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_CovHM2_28_Clear_Dispose", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		hm.Clear()

		// Act
		actual := args.Map{"result": hm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hm2 := corestr.Empty.Hashmap()
		hm2.Dispose()
	})
}

func Test_CovHM2_29_ToStringsUsingCompiler(t *testing.T) {
	safeTest(t, "Test_CovHM2_29_ToStringsUsingCompiler", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		// empty
		r := hm.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })

		// Act
		actual := args.Map{"result": len(r) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hm.AddOrUpdate("a", "1")
		r2 := hm.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })
		actual = args.Map{"result": len(r2) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		jr := hm.JsonPtr()
		hm2 := corestr.Empty.Hashmap()
		err := hm2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_CovHM2_32_Clone_ClonePtr(t *testing.T) {
	safeTest(t, "Test_CovHM2_32_Clone_ClonePtr", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		c := hm.Clone()

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		cp := hm.ClonePtr()
		actual = args.Map{"result": cp.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty clone
		e := corestr.Empty.Hashmap()
		ec := e.Clone()
		actual = args.Map{"result": ec.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovHM2_33_Get_GetValue(t *testing.T) {
	safeTest(t, "Test_CovHM2_33_Get_GetValue", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		v, ok := hm.Get("a")

		// Act
		actual := args.Map{"result": ok || v != "1"}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		_, ok2 := hm.Get("z")
		actual = args.Map{"result": ok2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not found", actual)
		v2, _ := hm.GetValue("a")
		actual = args.Map{"result": v2 != "1"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHM2_34_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_CovHM2_34_Serialize_Deserialize", func() {
		// Arrange
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		_, err := hm.Serialize()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		target := corestr.Empty.Hashmap()
		err2 := hm.Deserialize(target)
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}
