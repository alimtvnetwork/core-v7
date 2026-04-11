package corestrtests

import (
	"sort"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Collection creation and basic ops ──

func Test_Collection_NewEmpty(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_NewEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()

		// Act
		actual := args.Map{
			"notNil":  c != nil,
			"isEmpty": c.IsEmpty(),
			"length":  c.Length(),
			"count":   c.Count(),
		}

		// Assert
		expected := args.Map{
			"notNil":  true,
			"isEmpty": true,
			"length":  0,
			"count":   0,
		}
		expected.ShouldBeEqual(t, 0, "New.Collection.Empty returns empty -- new", actual)
	})
}

func Test_Collection_AddAndQuery(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_AddAndQuery", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.Add("hello")
		c.Add("world")

		// Act
		actual := args.Map{
			"length":     c.Length(),
			"hasItems":   c.HasItems(),
			"first":      c.First(),
			"last":       c.Last(),
			"lastIndex":  c.LastIndex(),
			"hasIndex0":  c.HasIndex(0),
			"hasIndex99": c.HasIndex(99),
		}

		// Assert
		expected := args.Map{
			"length":     2,
			"hasItems":   true,
			"first":      "hello",
			"last":       "world",
			"lastIndex":  1,
			"hasIndex0":  true,
			"hasIndex99": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection Add/query works -- two items", actual)
	})
}

func Test_Collection_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_FirstOrDefault_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()

		// Act
		actual := args.Map{"val": c.FirstOrDefault()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Collection.FirstOrDefault returns empty -- empty col", actual)
	})
}

func Test_Collection_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_LastOrDefault_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()

		// Act
		actual := args.Map{"val": c.LastOrDefault()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Collection.LastOrDefault returns empty -- empty col", actual)
	})
}

func Test_Collection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_AddNonEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmpty("")
		c.AddNonEmpty("hello")

		// Act
		actual := args.Map{"length": c.Length()}

		// Assert
		expected := args.Map{"length": 1}
		expected.ShouldBeEqual(t, 0, "Collection.AddNonEmpty skips empty -- one valid", actual)
	})
}

func Test_Collection_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_AddNonEmptyWhitespace", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyWhitespace("   ")
		c.AddNonEmptyWhitespace("hello")

		// Act
		actual := args.Map{"length": c.Length()}

		// Assert
		expected := args.Map{"length": 1}
		expected.ShouldBeEqual(t, 0, "Collection.AddNonEmptyWhitespace skips whitespace -- one valid", actual)
	})
}

func Test_Collection_Take(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_Take", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		taken := c.Take(2)

		// Act
		actual := args.Map{"len": taken.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection.Take returns 2 -- take 2 of 3", actual)
	})
}

func Test_Collection_Skip(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_Skip", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		skipped := c.Skip(1)

		// Act
		actual := args.Map{"len": skipped.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection.Skip returns 2 -- skip 1 of 3", actual)
	})
}

func Test_Collection_Reverse(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_Reverse", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		reversed := c.Reverse()

		// Act
		actual := args.Map{
			"first": reversed.First(),
			"last": reversed.Last(),
		}

		// Assert
		expected := args.Map{
			"first": "c",
			"last": "a",
		}
		expected.ShouldBeEqual(t, 0, "Collection.Reverse reverses -- 3 items", actual)
	})
}

func Test_Collection_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_IsEquals", func() {
		// Arrange
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"a", "b"})
		c3 := corestr.New.Collection.Strings([]string{"x", "y"})

		// Act
		actual := args.Map{
			"equal":    c1.IsEquals(c2),
			"notEqual": c1.IsEquals(c3),
		}

		// Assert
		expected := args.Map{
			"equal":    true,
			"notEqual": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection.IsEquals returns correct -- equal and not", actual)
	})
}

func Test_Collection_JsonString(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_JsonString", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		js := c.JsonString()

		// Act
		actual := args.Map{"hasContent": len(js) > 0}

		// Assert
		expected := args.Map{"hasContent": true}
		expected.ShouldBeEqual(t, 0, "Collection.JsonString returns content -- pointer receiver serialization", actual)
	})
}

func Test_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_RemoveAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.RemoveAt(1)

		// Act
		actual := args.Map{
			"length": c.Length(),
			"first": c.First(),
			"last": c.Last(),
		}

		// Assert
		expected := args.Map{
			"length": 2,
			"first": "a",
			"last": "c",
		}
		expected.ShouldBeEqual(t, 0, "Collection.RemoveAt removes middle -- 3 items", actual)
	})
}

func Test_Collection_AddIf(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_AddIf", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddIf(true, "yes")
		c.AddIf(false, "no")

		// Act
		actual := args.Map{"length": c.Length()}

		// Assert
		expected := args.Map{"length": 1}
		expected.ShouldBeEqual(t, 0, "Collection.AddIf conditionally adds -- one true", actual)
	})
}

func Test_Collection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_ConcatNew", func() {
		// Arrange
		c1 := corestr.New.Collection.Strings([]string{"a"})
		concat := c1.ConcatNew(0, "b")

		// Act
		actual := args.Map{"len": concat.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection.ConcatNew merges -- a + b", actual)
	})
}

func Test_Collection_InsertAt(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_InsertAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "c"})
		c.InsertAt(1, "b")

		// Act
		actual := args.Map{
			"len": c.Length(),
			"middle": c.IndexAt(1),
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"middle": "c",
		}
		expected.ShouldBeEqual(t, 0, "Collection.InsertAt at last index appends -- a,c,b", actual)
	})
}

func Test_Collection_UniqueList(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_UniqueList", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "a", "c", "b"})
		unique := c.UniqueList()

		// Act
		actual := args.Map{"len": len(unique)}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection.UniqueList returns 3 unique -- 5 items", actual)
	})
}

func Test_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_Filter", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"ab", "cd", "ae"})
		filtered := c.Filter(func(str string, index int) (string, bool, bool) {
			keep := len(str) > 0 && str[0] == 'a'
			return str, keep, false
		})

		// Act
		actual := args.Map{"len": len(filtered)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection.Filter returns 2 -- starts with a", actual)
	})
}

func Test_Collection_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_AddStrings", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddStrings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection.AddStrings adds all -- 3 strings", actual)
	})
}

func Test_Collection_ListStrings(t *testing.T) {
	safeTest(t, "Test_Cov6_Collection_ListStrings", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		list := c.ListStrings()

		// Act
		actual := args.Map{"len": len(list)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection.ListStrings returns 2 -- 2 items", actual)
	})
}

// ── Hashset ──

func Test_Hashset_NewAndBasic(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashset_NewAndBasic", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{
			"length":   hs.Length(),
			"hasItems": hs.HasItems(),
			"isEmpty":  hs.IsEmpty(),
			"hasA":     hs.Has("a"),
			"hasX":     hs.Has("x"),
			"missingX": hs.IsMissing("x"),
			"hasAny":   hs.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"length":   3,
			"hasItems": true,
			"isEmpty":  false,
			"hasA":     true,
			"hasX":     false,
			"missingX": true,
			"hasAny":   true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset basic ops -- 3 items", actual)
	})
}

func Test_Hashset_Add(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashset_Add", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.Add("hello")
		hs.Add("world")
		hs.Add("hello") // duplicate

		// Act
		actual := args.Map{"length": hs.Length()}

		// Assert
		expected := args.Map{"length": 2}
		expected.ShouldBeEqual(t, 0, "Hashset.Add deduplicates -- 2 unique", actual)
	})
}

func Test_Hashset_AddBool(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashset_AddBool", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		added1 := hs.AddBool("hello")
		added2 := hs.AddBool("hello")

		// Act
		actual := args.Map{
			"first": added1,
			"second": added2,
		}

		// Assert
		expected := args.Map{
			"first": false,
			"second": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset.AddBool returns isExist -- false then true", actual)
	})
}

func Test_Hashset_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashset_AddNonEmpty", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddNonEmpty("")
		hs.AddNonEmpty("hello")

		// Act
		actual := args.Map{"length": hs.Length()}

		// Assert
		expected := args.Map{"length": 1}
		expected.ShouldBeEqual(t, 0, "Hashset.AddNonEmpty skips empty -- one valid", actual)
	})
}

func Test_Hashset_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashset_AddStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddStrings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"length": hs.Length()}

		// Assert
		expected := args.Map{"length": 3}
		expected.ShouldBeEqual(t, 0, "Hashset.AddStrings adds all -- 3 strings", actual)
	})
}

func Test_Hashset_Contains(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashset_Contains", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"containsA": hs.Contains("a"),
			"containsX": hs.Contains("x"),
		}

		// Assert
		expected := args.Map{
			"containsA": true,
			"containsX": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset.Contains correct -- a yes, x no", actual)
	})
}

func Test_Hashset_SortedList(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashset_SortedList", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
		sorted := hs.SortedList()

		// Act
		actual := args.Map{
			"first": sorted[0],
			"last": sorted[2],
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"last": "c",
		}
		expected.ShouldBeEqual(t, 0, "Hashset.SortedList returns sorted -- 3 items", actual)
	})
}

func Test_Hashset_Remove(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashset_Remove", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
		hs.Remove("b")

		// Act
		actual := args.Map{
			"length": hs.Length(),
			"hasB": hs.Has("b"),
		}

		// Assert
		expected := args.Map{
			"length": 2,
			"hasB": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset.Remove removes item -- remove b", actual)
	})
}

func Test_Hashset_Clear(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashset_Clear", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		hs.Clear()

		// Act
		actual := args.Map{"isEmpty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"isEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashset.Clear empties -- after clear", actual)
	})
}

func Test_Hashset_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashset_IsEqual", func() {
		// Arrange
		hs1 := corestr.New.Hashset.Strings([]string{"a", "b"})
		hs2 := corestr.New.Hashset.Strings([]string{"b", "a"})
		hs3 := corestr.New.Hashset.Strings([]string{"x", "y"})

		// Act
		actual := args.Map{
			"equal":    hs1.IsEqual(hs2),
			"notEqual": hs1.IsEqual(hs3),
		}

		// Assert
		expected := args.Map{
			"equal":    true,
			"notEqual": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset.IsEqual correct -- same and different", actual)
	})
}

func Test_Hashset_Join(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashset_Join", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		joined := hs.JoinSorted(",")

		// Act
		actual := args.Map{"val": joined}

		// Assert
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Hashset.JoinSorted returns sorted csv -- 2 items", actual)
	})
}

func Test_Hashset_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashset_HasAllStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{
			"allPresent": hs.HasAllStrings([]string{"a", "b"}),
			"notAll":     hs.HasAllStrings([]string{"a", "x"}),
		}

		// Assert
		expected := args.Map{
			"allPresent": true,
			"notAll":     false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset.HasAllStrings correct -- all and partial", actual)
	})
}

func Test_Hashset_Filter(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashset_Filter", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"ab", "cd", "ae"})
		filtered := hs.Filter(func(s string) bool {
			return len(s) > 0 && s[0] == 'a'
		})

		// Act
		actual := args.Map{"len": filtered.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset.Filter returns 2 -- starts with a", actual)
	})
}

func Test_Hashset_String(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashset_String", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		s := hs.String()

		// Act
		actual := args.Map{"hasContent": len(s) > 0}

		// Assert
		expected := args.Map{"hasContent": true}
		expected.ShouldBeEqual(t, 0, "Hashset.String returns non-empty -- single item", actual)
	})
}

func Test_Hashset_Nil(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashset_Nil", func() {
		// Arrange
		var hs *corestr.Hashset

		// Act
		actual := args.Map{
			"isEmpty": hs.IsEmpty(),
			"length":  hs.Length(),
		}

		// Assert
		expected := args.Map{
			"isEmpty": true,
			"length":  0,
		}
		expected.ShouldBeEqual(t, 0, "Hashset nil receiver safe -- isEmpty and length", actual)
	})
}

// ── Hashmap ──

func Test_Hashmap_NewAndBasic(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashmap_NewAndBasic", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k1": "v1", "k2": "v2"})
		v1, found1 := hm.Get("k1")

		// Act
		actual := args.Map{
			"length":   hm.Length(),
			"hasItems": hm.HasItems(),
			"isEmpty":  hm.IsEmpty(),
			"hasK1":    hm.Has("k1"),
			"getK1":    v1,
			"foundK1":  found1,
		}

		// Assert
		expected := args.Map{
			"length":   2,
			"hasItems": true,
			"isEmpty":  false,
			"hasK1":    true,
			"getK1":    "v1",
			"foundK1":  true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap basic ops -- 2 items", actual)
	})
}

func Test_Hashmap_AddOrUpdate(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashmap_AddOrUpdate", func() {
		// Arrange
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k1", "v1")
		hm.AddOrUpdate("k1", "v2")
		v, _ := hm.Get("k1")

		// Act
		actual := args.Map{
			"length": hm.Length(),
			"val": v,
		}

		// Assert
		expected := args.Map{
			"length": 1,
			"val": "v2",
		}
		expected.ShouldBeEqual(t, 0, "Hashmap.AddOrUpdate updates -- key exists", actual)
	})
}

func Test_Hashmap_AllKeys(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashmap_AllKeys", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"c": "3", "a": "1", "b": "2"})
		keys := hm.AllKeys()
		sort.Strings(keys)

		// Act
		actual := args.Map{
			"first": keys[0],
			"last": keys[2],
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"last": "c",
		}
		expected.ShouldBeEqual(t, 0, "Hashmap.AllKeys returns all keys sorted -- 3 keys", actual)
	})
}

func Test_Hashmap_ValuesList(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashmap_ValuesList", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		vals := hm.ValuesList()

		// Act
		actual := args.Map{"len": len(vals)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap.ValuesList returns 1 -- single item", actual)
	})
}

func Test_Hashmap_Remove(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashmap_Remove", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k1": "v1", "k2": "v2"})
		hm.Remove("k1")

		// Act
		actual := args.Map{
			"length": hm.Length(),
			"hasK1": hm.Has("k1"),
		}

		// Assert
		expected := args.Map{
			"length": 1,
			"hasK1": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap.Remove removes key -- remove k1", actual)
	})
}

func Test_Hashmap_Clear(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashmap_Clear", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		hm.Clear()

		// Act
		actual := args.Map{"isEmpty": hm.IsEmpty()}

		// Assert
		expected := args.Map{"isEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap.Clear empties -- after clear", actual)
	})
}

func Test_Hashmap_Clone(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashmap_Clone", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		cloned := hm.Clone()
		v, _ := cloned.Get("k")

		// Act
		actual := args.Map{
			"sameLen": cloned.Length() == hm.Length(),
			"sameVal": v == "v",
		}

		// Assert
		expected := args.Map{
			"sameLen": true,
			"sameVal": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap.Clone preserves data -- clone roundtrip", actual)
	})
}

func Test_Hashmap_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashmap_IsEqual", func() {
		// Arrange
		hm1 := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		hm2 := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		hm3 := corestr.New.Hashmap.UsingMap(map[string]string{"k": "x"})

		// Act
		actual := args.Map{
			"equal":    hm1.IsEqual(*hm2),
			"notEqual": hm1.IsEqual(*hm3),
		}

		// Assert
		expected := args.Map{
			"equal":    true,
			"notEqual": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap.IsEqual correct -- same and different", actual)
	})
}

func Test_Hashmap_String(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashmap_String", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		s := hm.String()

		// Act
		actual := args.Map{"hasContent": len(s) > 0}

		// Assert
		expected := args.Map{"hasContent": true}
		expected.ShouldBeEqual(t, 0, "Hashmap.String returns non-empty -- single item", actual)
	})
}

func Test_Hashmap_Collection(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashmap_Collection", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
		col := hm.Collection()

		// Act
		actual := args.Map{"notNil": col != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashmap.Collection returns non-nil -- single item", actual)
	})
}

func Test_Hashmap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov6_Hashmap_Nil", func() {
		// Arrange
		var hm *corestr.Hashmap

		// Act
		actual := args.Map{
			"isEmpty": hm.IsEmpty(),
			"length":  hm.Length(),
		}

		// Assert
		expected := args.Map{
			"isEmpty": true,
			"length":  0,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap nil receiver safe -- isEmpty and length", actual)
	})
}

// ── StringUtils ──

func Test_StringUtils_WrapDouble(t *testing.T) {
	safeTest(t, "Test_Cov6_StringUtils_WrapDouble", func() {
		// Act
		actual := args.Map{"val": corestr.StringUtils.WrapDouble("hello")}

		// Assert
		expected := args.Map{"val": `"hello"`}
		expected.ShouldBeEqual(t, 0, "StringUtils.WrapDouble wraps correctly -- hello", actual)
	})
}

func Test_StringUtils_WrapSingle(t *testing.T) {
	safeTest(t, "Test_Cov6_StringUtils_WrapSingle", func() {
		// Act
		actual := args.Map{"val": corestr.StringUtils.WrapSingle("hello")}

		// Assert
		expected := args.Map{"val": "'hello'"}
		expected.ShouldBeEqual(t, 0, "StringUtils.WrapSingle wraps correctly -- hello", actual)
	})
}

func Test_StringUtils_WrapTilda(t *testing.T) {
	safeTest(t, "Test_Cov6_StringUtils_WrapTilda", func() {
		// Act
		actual := args.Map{"val": corestr.StringUtils.WrapTilda("hello")}

		// Assert
		expected := args.Map{"val": "`hello`"}
		expected.ShouldBeEqual(t, 0, "StringUtils.WrapTilda wraps correctly -- hello", actual)
	})
}
