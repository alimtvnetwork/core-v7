package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Collection — additional methods ──

func Test_Cov3_Collection_Add(t *testing.T) {
	safeTest(t, "Test_Cov3_Collection_Add", func() {
		c := corestr.New.Collection.Cap(5)
		c.Add("hello")
		actual := args.Map{"length": c.Length(), "hasAny": c.HasAnyItem()}
		expected := args.Map{"length": 1, "hasAny": true}
		expected.ShouldBeEqual(t, 0, "Collection Add returns 1 -- single item", actual)
	})
}

func Test_Cov3_Collection_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov3_Collection_AddStrings", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddStrings([]string{"ab", "cde"})
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddStrings returns 2 -- ab+cde", actual)
	})
}

func Test_Cov3_Collection_AddIf(t *testing.T) {
	safeTest(t, "Test_Cov3_Collection_AddIf", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddIf(true, "yes")
		c.AddIf(false, "no")
		actual := args.Map{"length": c.Length()}
		expected := args.Map{"length": 1}
		expected.ShouldBeEqual(t, 0, "Collection AddIf returns 1 -- conditional add", actual)
	})
}

// ── Hashmap — additional methods ──

func Test_Cov3_Hashmap_AddOrUpdate(t *testing.T) {
	safeTest(t, "Test_Cov3_Hashmap_AddOrUpdate", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("key", "value")
		val, found := h.Get("key")
		actual := args.Map{"val": val, "found": found}
		expected := args.Map{"val": "value", "found": true}
		expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdate and Get returns expected -- hit", actual)
	})
}

func Test_Cov3_Hashmap_Get_Miss(t *testing.T) {
	safeTest(t, "Test_Cov3_Hashmap_Get_Miss", func() {
		h := corestr.New.Hashmap.Cap(5)
		_, found := h.Get("missing")
		actual := args.Map{"found": found}
		expected := args.Map{"found": false}
		expected.ShouldBeEqual(t, 0, "Hashmap Get returns not found -- miss", actual)
	})
}

func Test_Cov3_Hashmap_Has(t *testing.T) {
	safeTest(t, "Test_Cov3_Hashmap_Has", func() {
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("key", "value")
		actual := args.Map{"hasKey": h.Has("key"), "hasMissing": h.Has("missing")}
		expected := args.Map{"hasKey": true, "hasMissing": false}
		expected.ShouldBeEqual(t, 0, "Hashmap Has returns expected -- hit and miss", actual)
	})
}

// ── Hashset — additional methods ──

func Test_Cov3_Hashset_Remove(t *testing.T) {
	safeTest(t, "Test_Cov3_Hashset_Remove", func() {
		h := corestr.New.Hashset.Cap(5)
		h.Adds("a", "b", "c")
		h.Remove("b")
		actual := args.Map{"length": h.Length(), "hasB": h.Has("b")}
		expected := args.Map{"length": 2, "hasB": false}
		expected.ShouldBeEqual(t, 0, "Hashset Remove returns 2 -- removed b", actual)
	})
}

// ── LeftRight — additional methods ──

func Test_Cov3_LeftRight_Empty(t *testing.T) {
	safeTest(t, "Test_Cov3_LeftRight_Empty", func() {
		lr := corestr.LeftRight{}
		actual := args.Map{"isLeftEmpty": lr.IsLeftEmpty(), "isRightEmpty": lr.IsRightEmpty()}
		expected := args.Map{"isLeftEmpty": true, "isRightEmpty": true}
		expected.ShouldBeEqual(t, 0, "LeftRight empty returns true -- all empty", actual)
	})
}

func Test_Cov3_LeftRight_PartialLeft(t *testing.T) {
	safeTest(t, "Test_Cov3_LeftRight_PartialLeft", func() {
		lr := corestr.NewLeftRight("l", "")
		actual := args.Map{"isLeftEmpty": lr.IsLeftEmpty(), "isRightEmpty": lr.IsRightEmpty(), "hasSafe": lr.HasSafeNonEmpty()}
		expected := args.Map{"isLeftEmpty": false, "isRightEmpty": true, "hasSafe": false}
		expected.ShouldBeEqual(t, 0, "LeftRight partial returns mixed -- only left", actual)
	})
}

// ── LeftMiddleRight — additional methods ──

func Test_Cov3_LeftMiddleRight_Empty(t *testing.T) {
	safeTest(t, "Test_Cov3_LeftMiddleRight_Empty", func() {
		lmr := corestr.NewLeftMiddleRight("", "", "")
		actual := args.Map{"isLeftEmpty": lmr.IsLeftEmpty(), "isMiddleEmpty": lmr.IsMiddleEmpty(), "isRightEmpty": lmr.IsRightEmpty()}
		expected := args.Map{"isLeftEmpty": true, "isMiddleEmpty": true, "isRightEmpty": true}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight empty returns true -- all empty", actual)
	})
}

// ── KeyValuePair ──

func Test_Cov3_KeyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov3_KeyValuePair", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"key": kv.Key, "value": kv.Value}
		expected := args.Map{"key": "k", "value": "v"}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns expected -- valid pair", actual)
	})
}

// ── KeyAnyValuePair ──

func Test_Cov3_KeyAnyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov3_KeyAnyValuePair", func() {
		kv := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		actual := args.Map{"key": kv.Key, "value": kv.Value}
		expected := args.Map{"key": "k", "value": 42}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns expected -- valid pair", actual)
	})
}

// ── SimpleSlice — additional methods ──

func Test_Cov3_SimpleSlice_Clear(t *testing.T) {
	safeTest(t, "Test_Cov3_SimpleSlice_Clear", func() {
		s := corestr.New.SimpleSlice.Cap(5)
		s.Adds("a", "b")
		s.Clear()
		actual := args.Map{"isEmpty": s.IsEmpty(), "length": s.Length()}
		expected := args.Map{"isEmpty": true, "length": 0}
		expected.ShouldBeEqual(t, 0, "SimpleSlice Clear returns empty -- after clear", actual)
	})
}

// ── NewValidValue ──

func Test_Cov3_NewValidValue(t *testing.T) {
	safeTest(t, "Test_Cov3_NewValidValue", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{"value": vv.Value, "isValid": vv.IsValid}
		expected := args.Map{"value": "hello", "isValid": true}
		expected.ShouldBeEqual(t, 0, "NewValidValue returns valid -- non-empty string", actual)
	})
}

func Test_Cov3_NewValidValue_Empty(t *testing.T) {
	safeTest(t, "Test_Cov3_NewValidValue_Empty", func() {
		vv := corestr.NewValidValue("")
		actual := args.Map{"value": vv.Value, "isValid": vv.IsValid}
		expected := args.Map{"value": "", "isValid": true}
		expected.ShouldBeEqual(t, 0, "NewValidValue always returns valid -- empty string", actual)
	})
}

// ── AllIndividualsLengthOfSimpleSlices ──

func Test_Cov3_AllIndividualsLengthOfSimpleSlices(t *testing.T) {
	safeTest(t, "Test_Cov3_AllIndividualsLengthOfSimpleSlices", func() {
		s1 := corestr.New.SimpleSlice.Cap(5)
		s1.Adds("ab", "cde")
		s2 := corestr.New.SimpleSlice.Cap(5)
		s2.Add("f")
		actual := args.Map{"result": corestr.AllIndividualsLengthOfSimpleSlices(s1, s2)}
		expected := args.Map{"result": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices returns 3 -- counts items not chars", actual)
	})
}

func Test_Cov3_AllIndividualsLengthOfSimpleSlices_NoArgs(t *testing.T) {
	safeTest(t, "Test_Cov3_AllIndividualsLengthOfSimpleSlices_NoArgs", func() {
		actual := args.Map{"result": corestr.AllIndividualsLengthOfSimpleSlices()}
		expected := args.Map{"result": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices returns 0 -- no args", actual)
	})
}
