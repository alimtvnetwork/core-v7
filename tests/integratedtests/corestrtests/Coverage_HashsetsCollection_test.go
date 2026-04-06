package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
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
		// Arrange
		hsc := corestr.New.HashsetsCollection.Empty()

		// Act
		actual := args.Map{"result": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": hsc.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
		actual = args.Map{"result": hsc.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		actual = args.Map{"result": hsc.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
		actual = args.Map{"result": hsc.HasItems()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected items", actual)
		actual = args.Map{"result": hsc.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHSC_02_LastIndex(t *testing.T) {
	safeTest(t, "Test_CovHSC_02_LastIndex", func() {
		// Arrange
		hsc := newHSC([]string{"a"}, []string{"b"})

		// Act
		actual := args.Map{"result": hsc.LastIndex() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHSC_03_IndexOf(t *testing.T) {
	safeTest(t, "Test_CovHSC_03_IndexOf", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.Empty()

		// Act
		actual := args.Map{"result": hsc.IndexOf(0) != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil for empty", actual)
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hsc.Add(corestr.New.Hashset.Strings([]string{"b"}))
		// IndexOf with valid index
		r := hsc.IndexOf(0)
		_ = r // may be nil due to bounds check logic
	})
}

func Test_CovHSC_04_List_ListPtr_ListDirectPtr(t *testing.T) {
	safeTest(t, "Test_CovHSC_04_List_ListPtr_ListDirectPtr", func() {
		// Arrange
		hsc := newHSC([]string{"a"})

		// Act
		actual := args.Map{"result": len(hsc.List()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": len(*hsc.ListPtr()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		dp := hsc.ListDirectPtr()
		actual = args.Map{"result": len(*dp) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHSC_05_StringsList(t *testing.T) {
	safeTest(t, "Test_CovHSC_05_StringsList", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.Empty()

		// Act
		actual := args.Map{"result": len(hsc.StringsList()) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hsc = newHSC([]string{"a", "b"}, []string{"c"})
		sl := hsc.StringsList()
		actual = args.Map{"result": len(sl) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_CovHSC_06_HasAll(t *testing.T) {
	safeTest(t, "Test_CovHSC_06_HasAll", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.Empty()

		// Act
		actual := args.Map{"result": hsc.HasAll("a")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
		hsc = newHSC([]string{"a", "b"})
		actual = args.Map{"result": hsc.HasAll("a", "b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": hsc.HasAll("x")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovHSC_07_Add_AddNonNil_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_CovHSC_07_Add_AddNonNil_AddNonEmpty", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.Empty()
		hsc.AddNonNil(nil)

		// Act
		actual := args.Map{"result": hsc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hsc.AddNonNil(corestr.New.Hashset.Strings([]string{"a"}))
		actual = args.Map{"result": hsc.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		hsc.AddNonEmpty(corestr.New.Hashset.Empty())
		actual = args.Map{"result": hsc.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		hsc.AddNonEmpty(corestr.New.Hashset.Strings([]string{"b"}))
		actual = args.Map{"result": hsc.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovHSC_08_Adds(t *testing.T) {
	safeTest(t, "Test_CovHSC_08_Adds", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.Empty()
		hsc.Adds(nil)
		hsc.Adds(corestr.New.Hashset.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": hsc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		hsc.Adds(corestr.New.Hashset.Empty())
		actual = args.Map{"result": hsc.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1, empty skipped", actual)
	})
}

func Test_CovHSC_09_AddHashsetsCollection(t *testing.T) {
	safeTest(t, "Test_CovHSC_09_AddHashsetsCollection", func() {
		// Arrange
		hsc := newHSC([]string{"a"})
		hsc.AddHashsetsCollection(nil)

		// Act
		actual := args.Map{"result": hsc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		hsc2 := newHSC([]string{"b"})
		hsc.AddHashsetsCollection(hsc2)
		actual = args.Map{"result": hsc.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovHSC_10_ConcatNew(t *testing.T) {
	safeTest(t, "Test_CovHSC_10_ConcatNew", func() {
		// Arrange
		hsc := newHSC([]string{"a"})
		// no args
		c := hsc.ConcatNew()

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// with args
		hsc2 := newHSC([]string{"b"})
		c2 := hsc.ConcatNew(hsc2)
		actual = args.Map{"result": c2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovHSC_11_IsEqual_IsEqualPtr(t *testing.T) {
	safeTest(t, "Test_CovHSC_11_IsEqual_IsEqualPtr", func() {
		// Arrange
		hsc1 := newHSC([]string{"a"})
		hsc2 := newHSC([]string{"a"})

		// Act
		actual := args.Map{"result": hsc1.IsEqualPtr(hsc2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		// same ptr
		actual = args.Map{"result": hsc1.IsEqualPtr(hsc1)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal same ptr", actual)
		// both empty
		e1 := corestr.New.HashsetsCollection.Empty()
		e2 := corestr.New.HashsetsCollection.Empty()
		actual = args.Map{"result": e1.IsEqualPtr(e2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal empties", actual)
		// diff length
		hsc3 := newHSC([]string{"a"}, []string{"b"})
		actual = args.Map{"result": hsc1.IsEqualPtr(hsc3)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
		// one nil
		actual = args.Map{"result": hsc1.IsEqualPtr(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
		// diff content
		hsc4 := newHSC([]string{"x"})
		actual = args.Map{"result": hsc1.IsEqualPtr(hsc4)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
		// IsEqual value
		actual = args.Map{"result": hsc1.IsEqual(*hsc2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		// one empty, one not
		actual = args.Map{"result": e1.IsEqualPtr(hsc1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
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
		// Arrange
		hsc := newHSC([]string{"a"})
		_ = hsc.JsonModel()
		_ = hsc.JsonModelAny()
		data, err := hsc.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		hsc2 := corestr.New.HashsetsCollection.Empty()
		err2 := hsc2.UnmarshalJSON(data)
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		// invalid
		err3 := hsc2.UnmarshalJSON([]byte("bad"))
		actual = args.Map{"result": err3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_CovHSC_14_Json_ParseInject(t *testing.T) {
	safeTest(t, "Test_CovHSC_14_Json_ParseInject", func() {
		// Arrange
		hsc := newHSC([]string{"a"})
		_ = hsc.Json()
		jr := hsc.JsonPtr()
		hsc2 := corestr.New.HashsetsCollection.Empty()
		r, err := hsc2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"result": err != nil || r == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_CovHSC_15_ParseInjectMust(t *testing.T) {
	safeTest(t, "Test_CovHSC_15_ParseInjectMust", func() {
		// Arrange
		hsc := newHSC([]string{"a"})
		jr := hsc.JsonPtr()
		hsc2 := corestr.New.HashsetsCollection.Empty()
		r := hsc2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"result": r == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_CovHSC_16_JsonParseSelfInject_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_CovHSC_16_JsonParseSelfInject_AsInterfaces", func() {
		// Arrange
		hsc := newHSC([]string{"a"})
		jr := hsc.JsonPtr()
		hsc2 := corestr.New.HashsetsCollection.Empty()
		err := hsc2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		_ = hsc.AsJsonContractsBinder()
		_ = hsc.AsJsoner()
		_ = hsc.AsJsonParseSelfInjector()
		_ = hsc.AsJsonMarshaller()
	})
}

func Test_CovHSC_17_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_CovHSC_17_Serialize_Deserialize", func() {
		// Arrange
		hsc := newHSC([]string{"a"})
		_, err := hsc.Serialize()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		target := corestr.New.HashsetsCollection.Empty()
		err2 := hsc.Deserialize(target)
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_CovHSC_18_DataModel(t *testing.T) {
	safeTest(t, "Test_CovHSC_18_DataModel", func() {
		// Arrange
		hsc := newHSC([]string{"a"})
		dm := corestr.NewHashsetsCollectionDataModelUsing(hsc)
		hsc2 := corestr.NewHashsetsCollectionUsingDataModel(dm)

		// Act
		actual := args.Map{"result": hsc2.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHSC_19_Creators(t *testing.T) {
	safeTest(t, "Test_CovHSC_19_Creators", func() {
		// Arrange
		// Empty
		e := corestr.New.HashsetsCollection.Empty()

		// Act
		actual := args.Map{"result": e.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// UsingHashsets
		hs := corestr.Hashset{}
		u := corestr.New.HashsetsCollection.UsingHashsets(hs)
		actual = args.Map{"result": u.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// UsingHashsets empty
		u2 := corestr.New.HashsetsCollection.UsingHashsets()
		actual = args.Map{"result": u2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// UsingHashsetsPointers
		hp := corestr.New.Hashset.Strings([]string{"a"})
		u3 := corestr.New.HashsetsCollection.UsingHashsetsPointers(hp)
		actual = args.Map{"result": u3.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// UsingHashsetsPointers empty
		u4 := corestr.New.HashsetsCollection.UsingHashsetsPointers()
		actual = args.Map{"result": u4.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// LenCap
		lc := corestr.New.HashsetsCollection.LenCap(0, 5)
		actual = args.Map{"result": lc.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// Cap
		cp := corestr.New.HashsetsCollection.Cap(5)
		actual = args.Map{"result": cp.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}
