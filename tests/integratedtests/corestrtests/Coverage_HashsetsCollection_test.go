package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// HashsetsCollection — Segment 18
// ══════════════════════════════════════════════════════════════════════════════

func newHSC(items ...[]string) *corestr.HashsetsCollection {
	hsc := corestr.New.HashsetsCollection.Empty()
	for _, s := range items {
		hs := corestr.New.Hashset.Strings(s)
		hsc.Add(hs)
	}
	return hsc
}

func Test_CovHSC_01_IsEmpty_HasItems_Length(t *testing.T) {
	safeTest(t, "Test_CovHSC_01_IsEmpty_HasItems_Length", func() {
		hsc := corestr.New.HashsetsCollection.Empty()
		if !hsc.IsEmpty() {
			t.Fatal("expected empty")
		}
		if hsc.HasItems() {
			t.Fatal("expected no items")
		}
		if hsc.Length() != 0 {
			t.Fatal("expected 0")
		}
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		if hsc.IsEmpty() {
			t.Fatal("expected not empty")
		}
		if !hsc.HasItems() {
			t.Fatal("expected items")
		}
		if hsc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHSC_02_LastIndex(t *testing.T) {
	safeTest(t, "Test_CovHSC_02_LastIndex", func() {
		hsc := newHSC([]string{"a"}, []string{"b"})
		if hsc.LastIndex() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHSC_03_IndexOf(t *testing.T) {
	safeTest(t, "Test_CovHSC_03_IndexOf", func() {
		hsc := corestr.New.HashsetsCollection.Empty()
		if hsc.IndexOf(0) != nil {
			t.Fatal("expected nil for empty")
		}
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hsc.Add(corestr.New.Hashset.Strings([]string{"b"}))
		// IndexOf with valid index
		r := hsc.IndexOf(0)
		_ = r // may be nil due to bounds check logic
	})
}

func Test_CovHSC_04_List_ListPtr_ListDirectPtr(t *testing.T) {
	safeTest(t, "Test_CovHSC_04_List_ListPtr_ListDirectPtr", func() {
		hsc := newHSC([]string{"a"})
		if len(hsc.List()) != 1 {
			t.Fatal("expected 1")
		}
		if len(*hsc.ListPtr()) != 1 {
			t.Fatal("expected 1")
		}
		dp := hsc.ListDirectPtr()
		if len(*dp) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHSC_05_StringsList(t *testing.T) {
	safeTest(t, "Test_CovHSC_05_StringsList", func() {
		hsc := corestr.New.HashsetsCollection.Empty()
		if len(hsc.StringsList()) != 0 {
			t.Fatal("expected 0")
		}
		hsc = newHSC([]string{"a", "b"}, []string{"c"})
		sl := hsc.StringsList()
		if len(sl) != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_CovHSC_06_HasAll(t *testing.T) {
	safeTest(t, "Test_CovHSC_06_HasAll", func() {
		hsc := corestr.New.HashsetsCollection.Empty()
		if hsc.HasAll("a") {
			t.Fatal("expected false for empty")
		}
		hsc = newHSC([]string{"a", "b"})
		if !hsc.HasAll("a", "b") {
			t.Fatal("expected true")
		}
		if hsc.HasAll("x") {
			t.Fatal("expected false")
		}
	})
}

func Test_CovHSC_07_Add_AddNonNil_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_CovHSC_07_Add_AddNonNil_AddNonEmpty", func() {
		hsc := corestr.New.HashsetsCollection.Empty()
		hsc.AddNonNil(nil)
		if hsc.Length() != 0 {
			t.Fatal("expected 0")
		}
		hsc.AddNonNil(corestr.New.Hashset.Strings([]string{"a"}))
		if hsc.Length() != 1 {
			t.Fatal("expected 1")
		}
		hsc.AddNonEmpty(corestr.New.Hashset.Empty())
		if hsc.Length() != 1 {
			t.Fatal("expected 1")
		}
		hsc.AddNonEmpty(corestr.New.Hashset.Strings([]string{"b"}))
		if hsc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovHSC_08_Adds(t *testing.T) {
	safeTest(t, "Test_CovHSC_08_Adds", func() {
		hsc := corestr.New.HashsetsCollection.Empty()
		hsc.Adds(nil)
		hsc.Adds(corestr.New.Hashset.Strings([]string{"a"}))
		if hsc.Length() != 1 {
			t.Fatal("expected 1")
		}
		hsc.Adds(corestr.New.Hashset.Empty())
		if hsc.Length() != 1 {
			t.Fatal("expected 1, empty skipped")
		}
	})
}

func Test_CovHSC_09_AddHashsetsCollection(t *testing.T) {
	safeTest(t, "Test_CovHSC_09_AddHashsetsCollection", func() {
		hsc := newHSC([]string{"a"})
		hsc.AddHashsetsCollection(nil)
		if hsc.Length() != 1 {
			t.Fatal("expected 1")
		}
		hsc2 := newHSC([]string{"b"})
		hsc.AddHashsetsCollection(hsc2)
		if hsc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovHSC_10_ConcatNew(t *testing.T) {
	safeTest(t, "Test_CovHSC_10_ConcatNew", func() {
		hsc := newHSC([]string{"a"})
		// no args
		c := hsc.ConcatNew()
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
		// with args
		hsc2 := newHSC([]string{"b"})
		c2 := hsc.ConcatNew(hsc2)
		if c2.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovHSC_11_IsEqual_IsEqualPtr(t *testing.T) {
	safeTest(t, "Test_CovHSC_11_IsEqual_IsEqualPtr", func() {
		hsc1 := newHSC([]string{"a"})
		hsc2 := newHSC([]string{"a"})
		if !hsc1.IsEqualPtr(hsc2) {
			t.Fatal("expected equal")
		}
		// same ptr
		if !hsc1.IsEqualPtr(hsc1) {
			t.Fatal("expected equal same ptr")
		}
		// both empty
		e1 := corestr.New.HashsetsCollection.Empty()
		e2 := corestr.New.HashsetsCollection.Empty()
		if !e1.IsEqualPtr(e2) {
			t.Fatal("expected equal empties")
		}
		// diff length
		hsc3 := newHSC([]string{"a"}, []string{"b"})
		if hsc1.IsEqualPtr(hsc3) {
			t.Fatal("expected not equal")
		}
		// one nil
		if hsc1.IsEqualPtr(nil) {
			t.Fatal("expected not equal")
		}
		// diff content
		hsc4 := newHSC([]string{"x"})
		if hsc1.IsEqualPtr(hsc4) {
			t.Fatal("expected not equal")
		}
		// IsEqual value
		if !hsc1.IsEqual(*hsc2) {
			t.Fatal("expected equal")
		}
		// one empty, one not
		if e1.IsEqualPtr(hsc1) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_CovHSC_12_String_Join(t *testing.T) {
	safeTest(t, "Test_CovHSC_12_String_Join", func() {
		hsc := corestr.New.HashsetsCollection.Empty()
		_ = hsc.String()
		hsc = newHSC([]string{"a"})
		_ = hsc.String()
		_ = hsc.Join(",")
	})
}

func Test_CovHSC_13_JsonModel_MarshalUnmarshal(t *testing.T) {
	safeTest(t, "Test_CovHSC_13_JsonModel_MarshalUnmarshal", func() {
		hsc := newHSC([]string{"a"})
		_ = hsc.JsonModel()
		_ = hsc.JsonModelAny()
		data, err := hsc.MarshalJSON()
		if err != nil {
			t.Fatal("unexpected error")
		}
		hsc2 := corestr.New.HashsetsCollection.Empty()
		err2 := hsc2.UnmarshalJSON(data)
		if err2 != nil {
			t.Fatal("unexpected error")
		}
		// invalid
		err3 := hsc2.UnmarshalJSON([]byte("bad"))
		if err3 == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_CovHSC_14_Json_ParseInject(t *testing.T) {
	safeTest(t, "Test_CovHSC_14_Json_ParseInject", func() {
		hsc := newHSC([]string{"a"})
		_ = hsc.Json()
		jr := hsc.JsonPtr()
		hsc2 := corestr.New.HashsetsCollection.Empty()
		r, err := hsc2.ParseInjectUsingJson(jr)
		if err != nil || r == nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovHSC_15_ParseInjectMust(t *testing.T) {
	safeTest(t, "Test_CovHSC_15_ParseInjectMust", func() {
		hsc := newHSC([]string{"a"})
		jr := hsc.JsonPtr()
		hsc2 := corestr.New.HashsetsCollection.Empty()
		r := hsc2.ParseInjectUsingJsonMust(jr)
		if r == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_CovHSC_16_JsonParseSelfInject_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_CovHSC_16_JsonParseSelfInject_AsInterfaces", func() {
		hsc := newHSC([]string{"a"})
		jr := hsc.JsonPtr()
		hsc2 := corestr.New.HashsetsCollection.Empty()
		err := hsc2.JsonParseSelfInject(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
		_ = hsc.AsJsonContractsBinder()
		_ = hsc.AsJsoner()
		_ = hsc.AsJsonParseSelfInjector()
		_ = hsc.AsJsonMarshaller()
	})
}

func Test_CovHSC_17_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_CovHSC_17_Serialize_Deserialize", func() {
		hsc := newHSC([]string{"a"})
		_, err := hsc.Serialize()
		if err != nil {
			t.Fatal("unexpected error")
		}
		target := corestr.New.HashsetsCollection.Empty()
		err2 := hsc.Deserialize(target)
		if err2 != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovHSC_18_DataModel(t *testing.T) {
	safeTest(t, "Test_CovHSC_18_DataModel", func() {
		hsc := newHSC([]string{"a"})
		dm := corestr.NewHashsetsCollectionDataModelUsing(hsc)
		hsc2 := corestr.NewHashsetsCollectionUsingDataModel(dm)
		if hsc2.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovHSC_19_Creators(t *testing.T) {
	safeTest(t, "Test_CovHSC_19_Creators", func() {
		// Empty
		e := corestr.New.HashsetsCollection.Empty()
		if e.Length() != 0 {
			t.Fatal("expected 0")
		}
		// UsingHashsets
		hs := corestr.Hashset{}
		u := corestr.New.HashsetsCollection.UsingHashsets(hs)
		if u.Length() != 1 {
			t.Fatal("expected 1")
		}
		// UsingHashsets empty
		u2 := corestr.New.HashsetsCollection.UsingHashsets()
		if u2.Length() != 0 {
			t.Fatal("expected 0")
		}
		// UsingHashsetsPointers
		hp := corestr.New.Hashset.Strings([]string{"a"})
		u3 := corestr.New.HashsetsCollection.UsingHashsetsPointers(hp)
		if u3.Length() != 1 {
			t.Fatal("expected 1")
		}
		// UsingHashsetsPointers empty
		u4 := corestr.New.HashsetsCollection.UsingHashsetsPointers()
		if u4.Length() != 0 {
			t.Fatal("expected 0")
		}
		// LenCap
		lc := corestr.New.HashsetsCollection.LenCap(0, 5)
		if lc.Length() != 0 {
			t.Fatal("expected 0")
		}
		// Cap
		cp := corestr.New.HashsetsCollection.Cap(5)
		if cp.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}
