package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Collection — additional methods ──

func Test_Cov3_Collection_Add(t *testing.T) {
	safeTest(t, "Test_Cov3_Collection_Add", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("hello")

		// Act
		actual := args.Map{
			"length": c.Length(),
			"hasAny": c.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"length": 1,
			"hasAny": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection Add returns 1 -- single item", actual)
	})
}

func Test_Cov3_Collection_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov3_Collection_AddStrings", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddStrings([]string{"ab", "cde"})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddStrings returns 2 -- ab+cde", actual)
	})
}

func Test_Cov3_Collection_AddIf(t *testing.T) {
	safeTest(t, "Test_Cov3_Collection_AddIf", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddIf(true, "yes")
		c.AddIf(false, "no")

		// Act
		actual := args.Map{"length": c.Length()}

		// Assert
		expected := args.Map{"length": 1}
		expected.ShouldBeEqual(t, 0, "Collection AddIf returns 1 -- conditional add", actual)
	})
}

// ── Hashmap — additional methods ──

func Test_Cov3_Hashmap_AddOrUpdate(t *testing.T) {
	safeTest(t, "Test_Cov3_Hashmap_AddOrUpdate", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("key", "value")
		val, found := h.Get("key")

		// Act
		actual := args.Map{
			"val": val,
			"found": found,
		}

		// Assert
		expected := args.Map{
			"val": "value",
			"found": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap AddOrUpdate and Get returns expected -- hit", actual)
	})
}

func Test_Cov3_Hashmap_Get_Miss(t *testing.T) {
	safeTest(t, "Test_Cov3_Hashmap_Get_Miss", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		_, found := h.Get("missing")

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": false}
		expected.ShouldBeEqual(t, 0, "Hashmap Get returns not found -- miss", actual)
	})
}

func Test_Cov3_Hashmap_Has(t *testing.T) {
	safeTest(t, "Test_Cov3_Hashmap_Has", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("key", "value")

		// Act
		actual := args.Map{
			"hasKey": h.Has("key"),
			"hasMissing": h.Has("missing"),
		}

		// Assert
		expected := args.Map{
			"hasKey": true,
			"hasMissing": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap Has returns expected -- hit and miss", actual)
	})
}

// ── Hashset — additional methods ──

func Test_Cov3_Hashset_Remove(t *testing.T) {
	safeTest(t, "Test_Cov3_Hashset_Remove", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.Adds("a", "b", "c")
		h.Remove("b")

		// Act
		actual := args.Map{
			"length": h.Length(),
			"hasB": h.Has("b"),
		}

		// Assert
		expected := args.Map{
			"length": 2,
			"hasB": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset Remove returns 2 -- removed b", actual)
	})
}

// ── LeftRight — additional methods ──

func Test_Cov3_LeftRight_Empty(t *testing.T) {
	safeTest(t, "Test_Cov3_LeftRight_Empty", func() {
		// Arrange
		lr := corestr.LeftRight{}

		// Act
		actual := args.Map{
			"isLeftEmpty": lr.IsLeftEmpty(),
			"isRightEmpty": lr.IsRightEmpty(),
		}

		// Assert
		expected := args.Map{
			"isLeftEmpty": true,
			"isRightEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight empty returns true -- all empty", actual)
	})
}

func Test_Cov3_LeftRight_PartialLeft(t *testing.T) {
	safeTest(t, "Test_Cov3_LeftRight_PartialLeft", func() {
		// Arrange
		lr := corestr.NewLeftRight("l", "")

		// Act
		actual := args.Map{
			"isLeftEmpty": lr.IsLeftEmpty(),
			"isRightEmpty": lr.IsRightEmpty(),
			"hasSafe": lr.HasSafeNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"isLeftEmpty": false,
			"isRightEmpty": true,
			"hasSafe": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight partial returns mixed -- only left", actual)
	})
}

// ── LeftMiddleRight — additional methods ──

func Test_Cov3_LeftMiddleRight_Empty(t *testing.T) {
	safeTest(t, "Test_Cov3_LeftMiddleRight_Empty", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("", "", "")

		// Act
		actual := args.Map{
			"isLeftEmpty": lmr.IsLeftEmpty(),
			"isMiddleEmpty": lmr.IsMiddleEmpty(),
			"isRightEmpty": lmr.IsRightEmpty(),
		}

		// Assert
		expected := args.Map{
			"isLeftEmpty": true,
			"isMiddleEmpty": true,
			"isRightEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight empty returns true -- all empty", actual)
	})
}

// ── KeyValuePair ──

func Test_Cov3_KeyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov3_KeyValuePair", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"key": kv.Key,
			"value": kv.Value,
		}

		// Assert
		expected := args.Map{
			"key": "k",
			"value": "v",
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns expected -- valid pair", actual)
	})
}

// ── KeyAnyValuePair ──

func Test_Cov3_KeyAnyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov3_KeyAnyValuePair", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: 42}

		// Act
		actual := args.Map{
			"key": kv.Key,
			"value": kv.Value,
		}

		// Assert
		expected := args.Map{
			"key": "k",
			"value": 42,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns expected -- valid pair", actual)
	})
}

// ── SimpleSlice — additional methods ──

func Test_Cov3_SimpleSlice_Clear(t *testing.T) {
	safeTest(t, "Test_Cov3_SimpleSlice_Clear", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)
		s.Adds("a", "b")
		s.Clear()

		// Act
		actual := args.Map{
			"isEmpty": s.IsEmpty(),
			"length": s.Length(),
		}

		// Assert
		expected := args.Map{
			"isEmpty": true,
			"length": 0,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice Clear returns empty -- after clear", actual)
	})
}

// ── NewValidValue ──

func Test_Cov3_NewValidValue(t *testing.T) {
	safeTest(t, "Test_Cov3_NewValidValue", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{
			"value": vv.Value,
			"isValid": vv.IsValid,
		}

		// Assert
		expected := args.Map{
			"value": "hello",
			"isValid": true,
		}
		expected.ShouldBeEqual(t, 0, "NewValidValue returns valid -- non-empty string", actual)
	})
}

func Test_Cov3_NewValidValue_Empty(t *testing.T) {
	safeTest(t, "Test_Cov3_NewValidValue_Empty", func() {
		// Arrange
		vv := corestr.NewValidValue("")

		// Act
		actual := args.Map{
			"value": vv.Value,
			"isValid": vv.IsValid,
		}

		// Assert
		expected := args.Map{
			"value": "",
			"isValid": true,
		}
		expected.ShouldBeEqual(t, 0, "NewValidValue always returns valid -- empty string", actual)
	})
}

// ── AllIndividualsLengthOfSimpleSlices ──

func Test_Cov3_AllIndividualsLengthOfSimpleSlices(t *testing.T) {
	safeTest(t, "Test_Cov3_AllIndividualsLengthOfSimpleSlices", func() {
		// Arrange
		s1 := corestr.New.SimpleSlice.Cap(5)
		s1.Adds("ab", "cde")
		s2 := corestr.New.SimpleSlice.Cap(5)
		s2.Add("f")

		// Act
		actual := args.Map{"result": corestr.AllIndividualsLengthOfSimpleSlices(s1, s2)}

		// Assert
		expected := args.Map{"result": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices returns 3 -- counts items not chars", actual)
	})
}

func Test_Cov3_AllIndividualsLengthOfSimpleSlices_NoArgs(t *testing.T) {
	safeTest(t, "Test_Cov3_AllIndividualsLengthOfSimpleSlices_NoArgs", func() {
		// Act
		actual := args.Map{"result": corestr.AllIndividualsLengthOfSimpleSlices()}

		// Assert
		expected := args.Map{"result": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices returns 0 -- no args", actual)
	})
}
