package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// SimpleSlice — Segment 9: Remaining methods (L700-1317)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovSS2_01_Collection_ToCollection(t *testing.T) {
	safeTest(t, "Test_CovSS2_01_Collection_ToCollection", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		col := ss.Collection(false)

		// Act
		actual := args.Map{"result": col.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		col2 := ss.ToCollection(true)
		actual = args.Map{"result": col2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		s := ss.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": e.String() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_CovSS2_04_ConcatNewSimpleSlices(t *testing.T) {
	safeTest(t, "Test_CovSS2_04_ConcatNewSimpleSlices", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Strings([]string{"a"})
		b := corestr.New.SimpleSlice.Strings([]string{"b"})
		r := a.ConcatNewSimpleSlices(b)

		// Act
		actual := args.Map{"result": r.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovSS2_05_ConcatNewStrings_ConcatNew(t *testing.T) {
	safeTest(t, "Test_CovSS2_05_ConcatNewStrings_ConcatNew", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		r := ss.ConcatNewStrings("b", "c")

		// Act
		actual := args.Map{"result": len(r) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		r2 := ss.ConcatNew("b")
		actual = args.Map{"result": r2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovSS2_06_Sort_Reverse(t *testing.T) {
	safeTest(t, "Test_CovSS2_06_Sort_Reverse", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"c", "a", "b"})
		ss.Sort()

		// Act
		actual := args.Map{"result": ss.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a first", actual)
		ss.Reverse()
		actual = args.Map{"result": ss.First() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c first after reverse", actual)
		// reverse 2 elements
		ss2 := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		ss2.Reverse()
		actual = args.Map{"result": ss2.First() != "b"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
		// reverse 1 element
		ss3 := corestr.New.SimpleSlice.Strings([]string{"a"})
		ss3.Reverse()
	})
}

func Test_CovSS2_07_JsonModel_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_CovSS2_07_JsonModel_JsonModelAny", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		m := ss.JsonModel()

		// Act
		actual := args.Map{"result": len(m) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		_ = ss.JsonModelAny()
	})
}

func Test_CovSS2_08_MarshalJSON_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovSS2_08_MarshalJSON_UnmarshalJSON", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		data, err := ss.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		ss2 := corestr.New.SimpleSlice.Strings([]string{})
		err2 := ss2.UnmarshalJSON(data)
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual = args.Map{"result": ss2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// invalid
		err3 := ss2.UnmarshalJSON([]byte("invalid"))
		actual = args.Map{"result": err3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
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
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		jr := ss.JsonPtr()
		ss2 := corestr.New.SimpleSlice.Strings([]string{})
		r, err := ss2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual = args.Map{"result": r.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovSS2_11_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovSS2_11_ParseInjectUsingJsonMust", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		jr := ss.JsonPtr()
		ss2 := corestr.New.SimpleSlice.Strings([]string{})
		r := ss2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"result": r.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovSS2_12_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovSS2_12_JsonParseSelfInject", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		jr := ss.JsonPtr()
		ss2 := corestr.New.SimpleSlice.Strings([]string{})
		err := ss2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
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
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		ss.Clear()

		// Act
		actual := args.Map{"result": ss.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss2 := corestr.New.SimpleSlice.Strings([]string{"x"})
		ss2.Dispose()
	})
}

func Test_CovSS2_15_Clone_ClonePtr_DeepClone_ShadowClone(t *testing.T) {
	safeTest(t, "Test_CovSS2_15_Clone_ClonePtr_DeepClone_ShadowClone", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		c := ss.Clone(true)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		cp := ss.ClonePtr(true)
		actual = args.Map{"result": cp.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		dc := ss.DeepClone()
		actual = args.Map{"result": dc.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		sc := ss.ShadowClone()
		actual = args.Map{"result": sc.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovSS2_16_IsDistinctEqualRaw_IsDistinctEqual(t *testing.T) {
	safeTest(t, "Test_CovSS2_16_IsDistinctEqualRaw_IsDistinctEqual", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ss.IsDistinctEqualRaw("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ss.IsDistinctEqualRaw("a", "c")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		other := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		actual = args.Map{"result": ss.IsDistinctEqual(other)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovSS2_17_IsUnorderedEqualRaw(t *testing.T) {
	safeTest(t, "Test_CovSS2_17_IsUnorderedEqualRaw", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		// with clone

		// Act
		actual := args.Map{"result": ss.IsUnorderedEqualRaw(true, "a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// without clone
		ss2 := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		actual = args.Map{"result": ss2.IsUnorderedEqualRaw(false, "a", "b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// diff length
		actual = args.Map{"result": ss.IsUnorderedEqualRaw(false, "a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// both empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": e.IsUnorderedEqualRaw(false)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovSS2_18_IsUnorderedEqual(t *testing.T) {
	safeTest(t, "Test_CovSS2_18_IsUnorderedEqual", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		other := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ss.IsUnorderedEqual(true, other)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// both empty
		e1 := corestr.New.SimpleSlice.Strings([]string{})
		e2 := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": e1.IsUnorderedEqual(false, e2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// nil right
		actual = args.Map{"result": ss.IsUnorderedEqual(false, nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovSS2_19_IsEqualByFunc(t *testing.T) {
	safeTest(t, "Test_CovSS2_19_IsEqualByFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		r := ss.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "b")

		// Act
		actual := args.Map{"result": r}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// mismatch
		r2 := ss.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "c")
		actual = args.Map{"result": r2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// diff length
		actual = args.Map{"result": ss.IsEqualByFunc(func(i int, l, r string) bool { return true }, "a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// both empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": e.IsEqualByFunc(func(i int, l, r string) bool { return true })}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovSS2_20_IsEqualByFuncLinesSplit(t *testing.T) {
	safeTest(t, "Test_CovSS2_20_IsEqualByFuncLinesSplit", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		r := ss.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return l == r })

		// Act
		actual := args.Map{"result": r}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// with trim
		ss2 := corestr.New.SimpleSlice.Strings([]string{" a ", " b "})
		r2 := ss2.IsEqualByFuncLinesSplit(true, ",", " a , b ", func(i int, l, r string) bool { return l == r })
		actual = args.Map{"result": r2}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// diff length
		actual = args.Map{"result": ss.IsEqualByFuncLinesSplit(false, ",", "a,b,c", func(i int, l, r string) bool { return true })}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// mismatch
		actual = args.Map{"result": ss.IsEqualByFuncLinesSplit(false, ",", "a,c", func(i int, l, r string) bool { return l == r })}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// empty — strings.Split("", ",") returns [""] (length 1) which != 0, so returns false
		e := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": e.IsEqualByFuncLinesSplit(false, ",", "", func(i int, l, r string) bool { return true })}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty vs split-empty mismatch", actual)
	})
}

func Test_CovSS2_21_DistinctDiffRaw(t *testing.T) {
	safeTest(t, "Test_CovSS2_21_DistinctDiffRaw", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		r := ss.DistinctDiffRaw("b", "c")

		// Act
		actual := args.Map{"result": len(r) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// nil right
		r2 := ss.DistinctDiffRaw()
		actual = args.Map{"result": len(r2) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovSS2_22_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_CovSS2_22_DistinctDiff", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		b := corestr.New.SimpleSlice.Strings([]string{"b", "c"})
		r := a.DistinctDiff(b)

		// Act
		actual := args.Map{"result": len(r) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// nil
		r2 := a.DistinctDiff(nil)
		actual = args.Map{"result": len(r2) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovSS2_23_AddedRemovedLinesDiff(t *testing.T) {
	safeTest(t, "Test_CovSS2_23_AddedRemovedLinesDiff", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		added, removed := ss.AddedRemovedLinesDiff("b", "c")

		// Act
		actual := args.Map{"result": len(added) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 added", actual)
		actual = args.Map{"result": len(removed) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 removed", actual)
	})
}

func Test_CovSS2_24_RemoveIndexes(t *testing.T) {
	safeTest(t, "Test_CovSS2_24_RemoveIndexes", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		r, err := ss.RemoveIndexes(1)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual = args.Map{"result": r.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// invalid index
		_, err2 := ss.RemoveIndexes(99)
		actual = args.Map{"result": err2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error for invalid index", actual)
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		_, err3 := e.RemoveIndexes(0)
		actual = args.Map{"result": err3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error for empty slice", actual)
	})
}

func Test_CovSS2_25_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_CovSS2_25_Serialize_Deserialize", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		_, err := ss.Serialize()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		target := corestr.New.SimpleSlice.Strings([]string{})
		err2 := ss.Deserialize(target)
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_CovSS2_26_SafeStrings(t *testing.T) {
	safeTest(t, "Test_CovSS2_26_SafeStrings", func() {
		// Arrange
		e := corestr.New.SimpleSlice.Strings([]string{})

		// Act
		actual := args.Map{"result": len(e.SafeStrings()) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual = args.Map{"result": len(ss.SafeStrings()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}
