package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// SimpleSlice — Segment 9: Remaining methods (L700-1317)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovSS2_01_Collection_ToCollection(t *testing.T) {
	safeTest(t, "Test_CovSS2_01_Collection_ToCollection", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		col := ss.Collection(false)
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
		col2 := ss.ToCollection(true)
		if col2.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovSS2_02_NonPtr_Ptr_ToPtr_ToNonPtr(t *testing.T) {
	safeTest(t, "Test_CovSS2_02_NonPtr_Ptr_ToPtr_ToNonPtr", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		_ = ss.NonPtr()
		_ = ss.Ptr()
		_ = ss.ToPtr()
		_ = ss.ToNonPtr()
	})
}

func Test_CovSS2_03_String(t *testing.T) {
	safeTest(t, "Test_CovSS2_03_String", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		s := ss.String()
		if s == "" {
			t.Fatal("expected non-empty")
		}
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		if e.String() != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_CovSS2_04_ConcatNewSimpleSlices(t *testing.T) {
	safeTest(t, "Test_CovSS2_04_ConcatNewSimpleSlices", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"a"})
		b := corestr.New.SimpleSlice.Strings([]string{"b"})
		r := a.ConcatNewSimpleSlices(b)
		if r.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovSS2_05_ConcatNewStrings_ConcatNew(t *testing.T) {
	safeTest(t, "Test_CovSS2_05_ConcatNewStrings_ConcatNew", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		r := ss.ConcatNewStrings("b", "c")
		if len(r) != 3 {
			t.Fatal("expected 3")
		}
		r2 := ss.ConcatNew("b")
		if r2.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovSS2_06_Sort_Reverse(t *testing.T) {
	safeTest(t, "Test_CovSS2_06_Sort_Reverse", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"c", "a", "b"})
		ss.Sort()
		if ss.First() != "a" {
			t.Fatal("expected a first")
		}
		ss.Reverse()
		if ss.First() != "c" {
			t.Fatal("expected c first after reverse")
		}
		// reverse 2 elements
		ss2 := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		ss2.Reverse()
		if ss2.First() != "b" {
			t.Fatal("expected b")
		}
		// reverse 1 element
		ss3 := corestr.New.SimpleSlice.Strings([]string{"a"})
		ss3.Reverse()
	})
}

func Test_CovSS2_07_JsonModel_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_CovSS2_07_JsonModel_JsonModelAny", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		m := ss.JsonModel()
		if len(m) != 1 {
			t.Fatal("expected 1")
		}
		_ = ss.JsonModelAny()
	})
}

func Test_CovSS2_08_MarshalJSON_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovSS2_08_MarshalJSON_UnmarshalJSON", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		data, err := ss.MarshalJSON()
		if err != nil {
			t.Fatal("unexpected error")
		}
		ss2 := corestr.New.SimpleSlice.Strings([]string{})
		err2 := ss2.UnmarshalJSON(data)
		if err2 != nil {
			t.Fatal("unexpected error")
		}
		if ss2.Length() != 2 {
			t.Fatal("expected 2")
		}
		// invalid
		err3 := ss2.UnmarshalJSON([]byte("invalid"))
		if err3 == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_CovSS2_09_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_CovSS2_09_Json_JsonPtr", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		_ = ss.Json()
		_ = ss.JsonPtr()
	})
}

func Test_CovSS2_10_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_CovSS2_10_ParseInjectUsingJson", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		jr := ss.JsonPtr()
		ss2 := corestr.New.SimpleSlice.Strings([]string{})
		r, err := ss2.ParseInjectUsingJson(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
		if r.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovSS2_11_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovSS2_11_ParseInjectUsingJsonMust", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		jr := ss.JsonPtr()
		ss2 := corestr.New.SimpleSlice.Strings([]string{})
		r := ss2.ParseInjectUsingJsonMust(jr)
		if r.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovSS2_12_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovSS2_12_JsonParseSelfInject", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		jr := ss.JsonPtr()
		ss2 := corestr.New.SimpleSlice.Strings([]string{})
		err := ss2.JsonParseSelfInject(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovSS2_13_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_CovSS2_13_AsInterfaces", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		_ = ss.AsJsonContractsBinder()
		_ = ss.AsJsoner()
		_ = ss.AsJsonParseSelfInjector()
		_ = ss.AsJsonMarshaller()
	})
}

func Test_CovSS2_14_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_CovSS2_14_Clear_Dispose", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		ss.Clear()
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
		ss2 := corestr.New.SimpleSlice.Strings([]string{"x"})
		ss2.Dispose()
	})
}

func Test_CovSS2_15_Clone_ClonePtr_DeepClone_ShadowClone(t *testing.T) {
	safeTest(t, "Test_CovSS2_15_Clone_ClonePtr_DeepClone_ShadowClone", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		c := ss.Clone(true)
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
		cp := ss.ClonePtr(true)
		if cp.Length() != 2 {
			t.Fatal("expected 2")
		}
		dc := ss.DeepClone()
		if dc.Length() != 2 {
			t.Fatal("expected 2")
		}
		sc := ss.ShadowClone()
		if sc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovSS2_16_IsDistinctEqualRaw_IsDistinctEqual(t *testing.T) {
	safeTest(t, "Test_CovSS2_16_IsDistinctEqualRaw_IsDistinctEqual", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		if !ss.IsDistinctEqualRaw("a", "b") {
			t.Fatal("expected true")
		}
		if ss.IsDistinctEqualRaw("a", "c") {
			t.Fatal("expected false")
		}
		other := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		if !ss.IsDistinctEqual(other) {
			t.Fatal("expected true")
		}
	})
}

func Test_CovSS2_17_IsUnorderedEqualRaw(t *testing.T) {
	safeTest(t, "Test_CovSS2_17_IsUnorderedEqualRaw", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		// with clone
		if !ss.IsUnorderedEqualRaw(true, "a", "b") {
			t.Fatal("expected true")
		}
		// without clone
		ss2 := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		if !ss2.IsUnorderedEqualRaw(false, "a", "b") {
			t.Fatal("expected true")
		}
		// diff length
		if ss.IsUnorderedEqualRaw(false, "a") {
			t.Fatal("expected false")
		}
		// both empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		if !e.IsUnorderedEqualRaw(false) {
			t.Fatal("expected true")
		}
	})
}

func Test_CovSS2_18_IsUnorderedEqual(t *testing.T) {
	safeTest(t, "Test_CovSS2_18_IsUnorderedEqual", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		other := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		if !ss.IsUnorderedEqual(true, other) {
			t.Fatal("expected true")
		}
		// both empty
		e1 := corestr.New.SimpleSlice.Strings([]string{})
		e2 := corestr.New.SimpleSlice.Strings([]string{})
		if !e1.IsUnorderedEqual(false, e2) {
			t.Fatal("expected true")
		}
		// nil right
		if ss.IsUnorderedEqual(false, nil) {
			t.Fatal("expected false")
		}
	})
}

func Test_CovSS2_19_IsEqualByFunc(t *testing.T) {
	safeTest(t, "Test_CovSS2_19_IsEqualByFunc", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		r := ss.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "b")
		if !r {
			t.Fatal("expected true")
		}
		// mismatch
		r2 := ss.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "c")
		if r2 {
			t.Fatal("expected false")
		}
		// diff length
		if ss.IsEqualByFunc(func(i int, l, r string) bool { return true }, "a") {
			t.Fatal("expected false")
		}
		// both empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		if !e.IsEqualByFunc(func(i int, l, r string) bool { return true }) {
			t.Fatal("expected true")
		}
	})
}

func Test_CovSS2_20_IsEqualByFuncLinesSplit(t *testing.T) {
	safeTest(t, "Test_CovSS2_20_IsEqualByFuncLinesSplit", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		r := ss.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return l == r })
		if !r {
			t.Fatal("expected true")
		}
		// with trim
		ss2 := corestr.New.SimpleSlice.Strings([]string{" a ", " b "})
		r2 := ss2.IsEqualByFuncLinesSplit(true, ",", " a , b ", func(i int, l, r string) bool { return l == r })
		if !r2 {
			t.Fatal("expected true")
		}
		// diff length
		if ss.IsEqualByFuncLinesSplit(false, ",", "a,b,c", func(i int, l, r string) bool { return true }) {
			t.Fatal("expected false")
		}
		// mismatch
		if ss.IsEqualByFuncLinesSplit(false, ",", "a,c", func(i int, l, r string) bool { return l == r }) {
			t.Fatal("expected false")
		}
		// empty — strings.Split("", ",") returns [""] (length 1) which != 0, so returns false
		e := corestr.New.SimpleSlice.Strings([]string{})
		if e.IsEqualByFuncLinesSplit(false, ",", "", func(i int, l, r string) bool { return true }) {
			t.Fatal("expected false for empty vs split-empty mismatch")
		}
	})
}

func Test_CovSS2_21_DistinctDiffRaw(t *testing.T) {
	safeTest(t, "Test_CovSS2_21_DistinctDiffRaw", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		r := ss.DistinctDiffRaw("b", "c")
		if len(r) != 2 {
			t.Fatalf("expected 2, got %d", len(r))
		}
		// nil right
		r2 := ss.DistinctDiffRaw()
		if len(r2) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovSS2_22_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_CovSS2_22_DistinctDiff", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		b := corestr.New.SimpleSlice.Strings([]string{"b", "c"})
		r := a.DistinctDiff(b)
		if len(r) != 2 {
			t.Fatalf("expected 2, got %d", len(r))
		}
		// nil
		r2 := a.DistinctDiff(nil)
		if len(r2) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovSS2_23_AddedRemovedLinesDiff(t *testing.T) {
	safeTest(t, "Test_CovSS2_23_AddedRemovedLinesDiff", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		added, removed := ss.AddedRemovedLinesDiff("b", "c")
		if len(added) != 1 {
			t.Fatalf("expected 1 added, got %d", len(added))
		}
		if len(removed) != 1 {
			t.Fatalf("expected 1 removed, got %d", len(removed))
		}
	})
}

func Test_CovSS2_24_RemoveIndexes(t *testing.T) {
	safeTest(t, "Test_CovSS2_24_RemoveIndexes", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		r, err := ss.RemoveIndexes(1)
		if err != nil {
			t.Fatal("unexpected error")
		}
		if r.Length() != 2 {
			t.Fatal("expected 2")
		}
		// invalid index
		_, err2 := ss.RemoveIndexes(99)
		if err2 == nil {
			t.Fatal("expected error for invalid index")
		}
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		_, err3 := e.RemoveIndexes(0)
		if err3 == nil {
			t.Fatal("expected error for empty slice")
		}
	})
}

func Test_CovSS2_25_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_CovSS2_25_Serialize_Deserialize", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		_, err := ss.Serialize()
		if err != nil {
			t.Fatal("unexpected error")
		}
		target := corestr.New.SimpleSlice.Strings([]string{})
		err2 := ss.Deserialize(target)
		if err2 != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovSS2_26_SafeStrings(t *testing.T) {
	safeTest(t, "Test_CovSS2_26_SafeStrings", func() {
		e := corestr.New.SimpleSlice.Strings([]string{})
		if len(e.SafeStrings()) != 0 {
			t.Fatal("expected 0")
		}
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if len(ss.SafeStrings()) != 1 {
			t.Fatal("expected 1")
		}
	})
}
