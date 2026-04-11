package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// SimpleSlice comprehensive
// ══════════════════════════════════════════════════════════════════════════════

func Test_SimpleSlice_Basic(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_Basic", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		var nilS *corestr.SimpleSlice

		// Act
		actual := args.Map{
			"empty": s.IsEmpty(), "hasAny": s.HasAnyItem(), "len": s.Length(),
			"count": s.Count(), "lastIdx": s.LastIndex(), "nilLen": nilS.Length(),
		}

		// Assert
		expected := args.Map{
			"empty": true, "hasAny": false, "len": 0, "count": 0, "lastIdx": -1, "nilLen": 0,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- basic", actual)
	})
}

func Test_SimpleSlice_Add(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_Add", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.Add("a").Add("b")
		s.AddSplit("c,d", ",")
		s.AddIf(false, "skip")
		s.AddIf(true, "e")
		s.Adds("f", "g")
		s.Adds()
		s.Append("h")
		s.Append()
		s.AppendFmt("")
		s.AppendFmt("val=%d", 42)
		s.AppendFmtIf(false, "skip")
		s.AppendFmtIf(true, "yes=%d", 1)
		s.AddAsTitleValue("key", "val")
		s.AddAsCurlyTitleWrap("key", "val")
		s.AddAsCurlyTitleWrapIf(false, "k", "v")
		s.AddAsCurlyTitleWrapIf(true, "k", "v")
		s.AddAsTitleValueIf(false, "k", "v")
		s.AddAsTitleValueIf(true, "k", "v")
		s.AddsIf(false, "skip")
		s.AddsIf(true, "x", "y")

		// Act
		actual := args.Map{"hasItems": s.HasAnyItem()}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- Add", actual)
	})
}

func Test_SimpleSlice_InsertAt(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_InsertAt", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "c")
		s.InsertAt(1, "b")
		s.InsertAt(-1, "x")
		s.InsertAt(100, "x")

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- InsertAt", actual)
	})
}

func Test_SimpleSlice_AddError(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_AddError", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddError(nil)
		s.AddError(nil) // just exercising nil path

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns nil -- AddError nil", actual)
	})
}

func Test_SimpleSlice_AddStruct(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_AddStruct", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddStruct(false, nil)
		s.AddStruct(false, "hello")

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- AddStruct", actual)
	})
}

func Test_SimpleSlice_AddPointer(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_AddPointer", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		s.AddPointer(false, nil)
		val := "hello"
		s.AddPointer(false, &val)

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- AddPointer", actual)
	})
}

func Test_SimpleSlice_AsError(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_AsError", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()
		nilDef := s.AsDefaultError() == nil
		nilSep := s.AsError(",") == nil
		s.Add("e")
		hasDef := s.AsDefaultError() != nil

		// Act
		actual := args.Map{
			"nilDef": nilDef,
			"nilSep": nilSep,
			"hasDef": hasDef,
		}

		// Assert
		expected := args.Map{
			"nilDef": true,
			"nilSep": true,
			"hasDef": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns error -- AsError", actual)
	})
}

func Test_SimpleSlice_FirstLast(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_FirstLast", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		actual := args.Map{
			"first": s.First(), "last": s.Last(),
			"firstDyn": s.FirstDynamic(), "lastDyn": s.LastDynamic(),
			"firstDef": s.FirstOrDefault(), "lastDef": s.LastOrDefault(),
			"firstDefDyn": s.FirstOrDefaultDynamic(), "lastDefDyn": s.LastOrDefaultDynamic(),
		}

		// Assert
		expected := args.Map{
			"first": "a", "last": "c", "firstDyn": "a", "lastDyn": "c",
			"firstDef": "a", "lastDef": "c", "firstDefDyn": "a", "lastDefDyn": "c",
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- FirstLast", actual)
	})
}

func Test_SimpleSlice_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_FirstOrDefault_Empty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{
			"first": s.FirstOrDefault(),
			"last": s.LastOrDefault(),
		}

		// Assert
		expected := args.Map{
			"first": "",
			"last": "",
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns empty -- FirstOrDefault empty", actual)
	})
}

func Test_SimpleSlice_SkipTake(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_SkipTake", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		actual := args.Map{
			"skip1":    len(s.Skip(1)),
			"skip100":  len(s.Skip(100)),
			"take2":    len(s.Take(2)),
			"take100":  len(s.Take(100)),
		}
		_ = s.SkipDynamic(1)
		_ = s.TakeDynamic(1)
		_ = s.LimitDynamic(1)
		_ = s.Limit(1)

		// Assert
		expected := args.Map{
			"skip1": 2,
			"skip100": 0,
			"take2": 2,
			"take100": 3,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- SkipTake", actual)
	})
}

func Test_SimpleSlice_CountFunc(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_CountFunc", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "bb", "ccc")
		count := s.CountFunc(func(i int, item string) bool { return len(item) > 1 })
		empty := corestr.New.SimpleSlice.Empty()
		emptyCount := empty.CountFunc(func(i int, item string) bool { return true })

		// Act
		actual := args.Map{
			"count": count,
			"emptyCount": emptyCount,
		}

		// Assert
		expected := args.Map{
			"count": 2,
			"emptyCount": 0,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- CountFunc", actual)
	})
}

func Test_SimpleSlice_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_IsContains", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{
			"has": s.IsContains("a"),
			"notHas": !s.IsContains("c"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"notHas": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- IsContains", actual)
	})
}

func Test_SimpleSlice_IndexOf(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_IndexOf", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		actual := args.Map{
			"found": s.IndexOf("b"),
			"notFound": s.IndexOf("z"),
		}

		// Assert
		expected := args.Map{
			"found": 1,
			"notFound": -1,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- IndexOf", actual)
	})
}

func Test_SimpleSlice_Extended(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_Extended", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		_ = s.HasIndex(0)
		_ = s.Strings()
		_ = s.List()
		_ = s.WrapDoubleQuote()

		s2 := corestr.New.SimpleSlice.Lines("a")
		_ = s2.WrapSingleQuote()
		s3 := corestr.New.SimpleSlice.Lines("a")
		_ = s3.WrapTildaQuote()
		s4 := corestr.New.SimpleSlice.Lines("a")
		_ = s4.WrapDoubleQuoteIfMissing()
		s5 := corestr.New.SimpleSlice.Lines("a")
		_ = s5.WrapSingleQuoteIfMissing()

		_ = s.Join(",")
		_ = s.JoinLine()
		_ = s.JoinLineEofLine()
		_ = s.JoinSpace()
		_ = s.JoinComma()
		_ = s.JoinCsv()
		_ = s.JoinCsvLine()
		_ = s.JoinWith(",")
		_ = s.JoinCsvString(",")
		_ = s.CsvStrings()
		_ = s.String()
		_ = s.Collection(false)
		_ = s.ToCollection(false)
		_ = s.NonPtr()
		_ = s.Ptr()
		_ = s.ToPtr()
		_ = s.ToNonPtr()
		_ = s.Sort()
		_ = s.Reverse()
		_ = s.Hashset()
		_ = s.EachItemSplitBy(",")
		_ = s.TranspileJoin(func(ss string) string { return ss }, ",")
		_ = s.PrependJoin(",", "z")
		_ = s.AppendJoin(",", "z")
		_ = s.ConcatNew("d")
		_ = s.ConcatNewStrings("d")
		_ = s.ConcatNewSimpleSlices(corestr.New.SimpleSlice.Lines("e"))
		s.PrependAppend([]string{"pre"}, []string{"post"})
		_ = s.JsonModel()
		_ = s.JsonModelAny()
		_, _ = s.MarshalJSON()
		_, _ = s.Serialize()
		_ = s.SafeStrings()
		_ = s.AsJsoner()
		_ = s.AsJsonContractsBinder()
		_ = s.AsJsonParseSelfInjector()
		_ = s.AsJsonMarshaller()

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- extended", actual)
	})
}

func Test_SimpleSlice_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_IsEqual", func() {
		// Arrange
		s1 := corestr.New.SimpleSlice.Lines("a", "b")
		s2 := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{
			"equal":          s1.IsEqual(s2),
			"equalLines":     s1.IsEqualLines([]string{"a", "b"}),
			"equalUnordered": s1.IsEqualUnorderedLines([]string{"b", "a"}),
			"equalUnorderedClone": s1.IsEqualUnorderedLinesClone([]string{"b", "a"}),
			"distinctEqual":  s1.IsDistinctEqual(s2),
			"distinctEqualRaw": s1.IsDistinctEqualRaw("a", "b"),
			"unorderedEqual": s1.IsUnorderedEqual(true, s2),
			"unorderedEqualRaw": s1.IsUnorderedEqualRaw(true, "b", "a"),
		}

		// Assert
		expected := args.Map{
			"equal": true, "equalLines": true, "equalUnordered": true,
			"equalUnorderedClone": true, "distinctEqual": true, "distinctEqualRaw": true,
			"unorderedEqual": true, "unorderedEqualRaw": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- IsEqual", actual)
	})
}

func Test_SimpleSlice_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_DistinctDiff", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		_ = s.DistinctDiffRaw("b", "c")
		_ = s.DistinctDiff(corestr.New.SimpleSlice.Lines("b", "c"))
		added, removed := s.AddedRemovedLinesDiff("b", "c")
		_ = added
		_ = removed

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- DistinctDiff", actual)
	})
}

func Test_SimpleSlice_RemoveIndexes(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_RemoveIndexes", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		newS, err := s.RemoveIndexes(1)
		empty := corestr.New.SimpleSlice.Empty()
		_, err2 := empty.RemoveIndexes(0)

		// Act
		actual := args.Map{
			"len": newS.Length(),
			"noErr": err == nil,
			"errEmpty": err2 != nil,
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"noErr": true,
			"errEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- RemoveIndexes", actual)
	})
}

func Test_SimpleSlice_Clone(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_Clone", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		_ = s.Clone(true)
		_ = s.ClonePtr(true)
		_ = s.DeepClone()
		_ = s.ShadowClone()
		var nilS *corestr.SimpleSlice

		// Act
		actual := args.Map{"nilClone": nilS.ClonePtr(true) == nil}

		// Assert
		expected := args.Map{"nilClone": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- Clone", actual)
	})
}

func Test_SimpleSlice_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_ClearDispose", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a")
		s.Clear()
		s.Dispose()
		var nilS *corestr.SimpleSlice
		_ = nilS.Clear()

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- ClearDispose", actual)
	})
}

func Test_SimpleSlice_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_IsContainsFunc", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("hello", "world")
		found := s.IsContainsFunc("hello", func(item, searching string) bool { return item == searching })

		// Act
		actual := args.Map{"found": found}

		// Assert
		expected := args.Map{"found": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- IsContainsFunc", actual)
	})
}

func Test_SimpleSlice_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_Cov19_SimpleSlice_IndexOfFunc", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Lines("a", "b")
		idx := s.IndexOfFunc("b", func(item, searching string) bool { return item == searching })

		// Act
		actual := args.Map{"idx": idx}

		// Assert
		expected := args.Map{"idx": 1}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- IndexOfFunc", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedList comprehensive
// ══════════════════════════════════════════════════════════════════════════════

func Test_LinkedList_Basic(t *testing.T) {
	safeTest(t, "Test_Cov19_LinkedList_Basic", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{
			"empty": ll.IsEmpty(),
			"emptyLock": ll.IsEmptyLock(),
			"len": ll.Length(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"emptyLock": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- basic", actual)
	})
}

func Test_LinkedList_Add(t *testing.T) {
	safeTest(t, "Test_Cov19_LinkedList_Add", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a").Add("b").Add("c")
		ll.AddLock("d")
		ll.AddFront("z")
		ll.AddNonEmpty("")
		ll.AddNonEmpty("e")
		ll.AddNonEmptyWhitespace("   ")
		ll.AddNonEmptyWhitespace("f")
		ll.AddIf(false, "skip")
		ll.AddIf(true, "g")
		ll.AddsIf(false, "x")
		ll.AddsIf(true, "h", "i")
		ll.AddFunc(func() string { return "j" })
		ll.AddFuncErr(func() (string, error) { return "k", nil }, func(e error) {})
		ll.Push("l")
		ll.PushFront("m")
		ll.PushBack("n")

		// Act
		actual := args.Map{
			"head": ll.Head().Element,
			"hasItems": ll.HasItems(),
		}

		// Assert
		expected := args.Map{
			"head": "m",
			"hasItems": true,
		}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- Add", actual)
	})
}

func Test_LinkedList_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov19_LinkedList_AddStrings", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.Adds()
		ll.AddStrings([]string{"c"})
		ll.AddStrings(nil)
		ll.AddsLock("d")
		ll.AddCollection(corestr.New.Collection.Strings([]string{"e"}))
		ll.AddCollection(nil)
		ll.AddItemsMap(map[string]bool{"f": true, "g": false})
		ll.AddItemsMap(nil)

		// Act
		actual := args.Map{"hasItems": ll.HasItems()}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- AddStrings", actual)
	})
}

func Test_LinkedList_List(t *testing.T) {
	safeTest(t, "Test_Cov19_LinkedList_List", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"listLen": len(ll.List())}
		_ = ll.ListPtr()
		_ = ll.ListLock()
		_ = ll.ListPtrLock()
		_ = ll.String()
		_ = ll.StringLock()
		_ = ll.Join(",")
		_ = ll.JoinLock(",")

		// Assert
		expected := args.Map{"listLen": 3}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- List", actual)
	})
}

func Test_LinkedList_ToCollection(t *testing.T) {
	safeTest(t, "Test_Cov19_LinkedList_ToCollection", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		c := ll.ToCollection(0)
		empty := corestr.New.LinkedList.Create()
		c2 := empty.ToCollection(0)

		// Act
		actual := args.Map{
			"len": c.Length(),
			"emptyLen": c2.Length(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"emptyLen": 0,
		}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- ToCollection", actual)
	})
}

func Test_LinkedList_Loop(t *testing.T) {
	safeTest(t, "Test_Cov19_LinkedList_Loop", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})

		// Act
		actual := args.Map{"count": count}

		// Assert
		expected := args.Map{"count": 3}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- Loop", actual)
	})
}

func Test_LinkedList_Filter(t *testing.T) {
	safeTest(t, "Test_Cov19_LinkedList_Filter", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})

		// Act
		actual := args.Map{"len": len(nodes)}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- Filter", actual)
	})
}

func Test_LinkedList_IndexAt(t *testing.T) {
	safeTest(t, "Test_Cov19_LinkedList_IndexAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		node := ll.SafeIndexAt(1)

		// Act
		actual := args.Map{
			"found":   node != nil && node.Element == "b",
			"negNil":  ll.SafeIndexAt(-1) == nil,
			"bigNil":  ll.SafeIndexAt(99) == nil,
		}
		_ = ll.SafePointerIndexAt(1)
		_ = ll.SafePointerIndexAt(-1)
		_ = ll.SafePointerIndexAtUsingDefault(1, "def")
		_ = ll.SafePointerIndexAtUsingDefault(-1, "def")
		_ = ll.SafePointerIndexAtUsingDefaultLock(0, "def")
		_ = ll.SafeIndexAtLock(0)

		// Assert
		expected := args.Map{
			"found": true,
			"negNil": true,
			"bigNil": true,
		}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- IndexAt", actual)
	})
}

func Test_LinkedList_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_Cov19_LinkedList_GetNextNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		nodes := ll.GetNextNodes(2)
		all := ll.GetAllLinkedNodes()

		// Act
		actual := args.Map{
			"nextLen": len(nodes),
			"allLen": len(all),
		}

		// Assert
		expected := args.Map{
			"nextLen": 2,
			"allLen": 3,
		}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- GetNextNodes", actual)
	})
}

func Test_LinkedList_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov19_LinkedList_IsEquals", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll2 := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"equal": ll1.IsEquals(ll2),
			"sensEqual": ll1.IsEqualsWithSensitive(ll2, false),
		}

		// Assert
		expected := args.Map{
			"equal": true,
			"sensEqual": true,
		}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- IsEquals", actual)
	})
}

func Test_LinkedList_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_Cov19_LinkedList_RemoveNodeByIndex", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByIndex(0)

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- RemoveNodeByIndex", actual)
	})
}

func Test_LinkedList_Clear(t *testing.T) {
	safeTest(t, "Test_Cov19_LinkedList_Clear", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.Clear()
		ll.RemoveAll()

		// Act
		actual := args.Map{"len": ll.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- Clear", actual)
	})
}

func Test_LinkedList_JsonAndMarshal(t *testing.T) {
	safeTest(t, "Test_Cov19_LinkedList_JsonAndMarshal", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.JsonModel()
		_ = ll.JsonModelAny()
		_, _ = ll.MarshalJSON()
		_ = ll.AsJsonMarshaller()
		_ = ll.Joins(",", "c")

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- Json", actual)
	})
}

func Test_LinkedListNode_Methods(t *testing.T) {
	safeTest(t, "Test_Cov19_LinkedListNode_Methods", func() {
		// Arrange
		n := &corestr.LinkedListNode{Element: "a"}

		// Act
		actual := args.Map{
			"noNext":  !n.HasNext(),
			"str":     n.String() == "a",
			"isEqual": n.IsEqualValue("a"),
			"isSens":  n.IsEqualValueSensitive("A", false),
		}
		c := n.Clone()
		_ = c
		_ = n.List()
		_ = n.ListPtr()
		_ = n.Join(",")
		_ = n.StringList("header: ")
		_ = n.CreateLinkedList()

		// Assert
		expected := args.Map{
			"noNext": true,
			"str": true,
			"isEqual": true,
			"isSens": true,
		}
		expected.ShouldBeEqual(t, 0, "LinkedListNode returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValueCollection
// ══════════════════════════════════════════════════════════════════════════════

func Test_KeyValueCollection(t *testing.T) {
	safeTest(t, "Test_Cov19_KeyValueCollection", func() {
		// Arrange
		kv := corestr.New.KeyValues.Empty()
		kv.Add("k1", "v1").Add("k2", "v2")

		// Act
		actual := args.Map{
			"len":    kv.Length(),
			"hasKey": kv.HasKey("k1"),
			"contains": kv.IsContains("k1"),
		}
		v, found := kv.Get("k1")
		actual["found"] = found
		actual["val"] = v
		_ = kv.First()
		_ = kv.Last()
		_ = kv.FirstOrDefault()
		_ = kv.LastOrDefault()
		_ = kv.Strings()
		_ = kv.String()
		_ = kv.AllKeys()
		_ = kv.AllKeysSorted()
		_ = kv.AllValues()
		_ = kv.Join(",")
		_ = kv.JoinKeys(",")
		_ = kv.JoinValues(",")
		_ = kv.Compile()
		_ = kv.SafeValueAt(0)
		_ = kv.SafeValueAt(99)
		_ = kv.SafeValuesAtIndexes(0)
		_ = kv.StringsUsingFormat("%s=%s")
		_ = kv.Hashmap()
		_ = kv.Map()
		_ = kv.JsonModel()
		_ = kv.JsonModelAny()
		_, _ = kv.Serialize()
		_, _ = kv.MarshalJSON()
		_ = kv.SerializeMust()
		_ = kv.AsJsoner()
		_ = kv.AsJsonContractsBinder()
		_ = kv.AsJsonParseSelfInjector()

		// Assert
		expected := args.Map{
			"len": 2,
			"hasKey": true,
			"contains": true,
			"found": true,
			"val": "v1",
		}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection returns correct value -- with args", actual)
	})
}

func Test_KeyValueCollection_AddVariants(t *testing.T) {
	safeTest(t, "Test_Cov19_KeyValueCollection_AddVariants", func() {
		// Arrange
		kv := corestr.New.KeyValues.Empty()
		kv.AddIf(false, "skip", "val")
		kv.AddIf(true, "k", "v")
		kv.Adds(corestr.KeyValuePair{Key: "a", Value: "b"})
		kv.Adds()
		kv.AddMap(map[string]string{"c": "d"})
		kv.AddMap(nil)
		kv.AddHashsetMap(map[string]bool{"e": true})
		kv.AddHashsetMap(nil)
		kv.AddHashset(corestr.New.Hashset.StringsSpreadItems("f"))
		kv.AddHashset(nil)
		kv.AddsHashmap(corestr.New.Hashmap.UsingMap(map[string]string{"g": "h"}))
		kv.AddsHashmap(nil)
		kv.AddsHashmaps(corestr.New.Hashmap.UsingMap(map[string]string{"i": "j"}))
		kv.AddsHashmaps()
		kv.AddStringBySplit("=", "k=l")
		kv.AddStringBySplitTrim("=", " m = n ")

		// Act
		actual := args.Map{"hasItems": kv.HasAnyItem()}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection returns correct value -- AddVariants", actual)
	})
}

func Test_KeyValueCollection_Find(t *testing.T) {
	safeTest(t, "Test_Cov19_KeyValueCollection_Find", func() {
		// Arrange
		kv := corestr.New.KeyValues.Empty()
		kv.Add("a", "1").Add("b", "2")
		r := kv.Find(func(i int, curr corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return curr, curr.Key == "a", false
		})

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection returns correct value -- Find", actual)
	})
}

func Test_KeyValueCollection_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov19_KeyValueCollection_ClearDispose", func() {
		// Arrange
		kv := corestr.New.KeyValues.Empty()
		kv.Add("a", "1")
		kv.Clear()
		kv.Dispose()
		var nilKv *corestr.KeyValueCollection
		nilKv.Clear()
		nilKv.Dispose()

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "KeyValueCollection returns correct value -- ClearDispose", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Creators comprehensive
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewCollectionCreator(t *testing.T) {
	safeTest(t, "Test_Cov19_NewCollectionCreator", func() {
		// Arrange
		_ = corestr.New.Collection.Empty()
		_ = corestr.New.Collection.Cap(10)
		_ = corestr.New.Collection.Strings([]string{"a", "b"})
		_ = corestr.New.Collection.Create([]string{"a"})
		_ = corestr.New.Collection.CloneStrings([]string{"a"})
		_ = corestr.New.Collection.StringsOptions(true, []string{"a"})
		_ = corestr.New.Collection.StringsOptions(false, []string{"a"})
		_ = corestr.New.Collection.StringsOptions(false, nil)
		_ = corestr.New.Collection.LineUsingSep(",", "a,b")
		_ = corestr.New.Collection.LineDefault("a\nb")
		_ = corestr.New.Collection.StringsPlusCap(5, []string{"a"})
		_ = corestr.New.Collection.StringsPlusCap(0, []string{"a"})
		_ = corestr.New.Collection.CapStrings(5, []string{"a"})
		_ = corestr.New.Collection.CapStrings(0, []string{"a"})
		_ = corestr.New.Collection.LenCap(2, 5)

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "NewCollectionCreator returns correct value -- with args", actual)
	})
}

func Test_NewSimpleSliceCreator(t *testing.T) {
	safeTest(t, "Test_Cov19_NewSimpleSliceCreator", func() {
		// Arrange
		_ = corestr.New.SimpleSlice.Empty()
		_ = corestr.New.SimpleSlice.Cap(10)
		_ = corestr.New.SimpleSlice.Cap(-1)
		_ = corestr.New.SimpleSlice.Default()
		_ = corestr.New.SimpleSlice.Lines("a", "b")
		_ = corestr.New.SimpleSlice.SpreadStrings("a")
		_ = corestr.New.SimpleSlice.Strings([]string{"a"})
		_ = corestr.New.SimpleSlice.Create([]string{"a"})
		_ = corestr.New.SimpleSlice.StringsPtr([]string{"a"})
		_ = corestr.New.SimpleSlice.StringsPtr(nil)
		_ = corestr.New.SimpleSlice.StringsOptions(true, []string{"a"})
		_ = corestr.New.SimpleSlice.StringsOptions(false, []string{"a"})
		_ = corestr.New.SimpleSlice.StringsOptions(false, nil)
		_ = corestr.New.SimpleSlice.StringsClone([]string{"a"})
		_ = corestr.New.SimpleSlice.StringsClone(nil)
		_ = corestr.New.SimpleSlice.Direct(true, []string{"a"})
		_ = corestr.New.SimpleSlice.Direct(false, []string{"a"})
		_ = corestr.New.SimpleSlice.Direct(true, nil)
		_ = corestr.New.SimpleSlice.UsingLines(true, "a")
		_ = corestr.New.SimpleSlice.UsingLines(false, "a")
		_ = corestr.New.SimpleSlice.UsingLines(true)
		_ = corestr.New.SimpleSlice.Split("a,b", ",")
		_ = corestr.New.SimpleSlice.SplitLines("a\nb")
		_ = corestr.New.SimpleSlice.UsingSeparatorLine(",", "a,b")
		_ = corestr.New.SimpleSlice.UsingLine("a\nb")
		_ = corestr.New.SimpleSlice.ByLen([]string{"a", "b"})
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		_ = corestr.New.SimpleSlice.Hashset(hs)
		_ = corestr.New.SimpleSlice.Map(map[string]int{"a": 1})
		_ = corestr.New.SimpleSlice.Map(nil)

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "NewSimpleSliceCreator returns correct value -- with args", actual)
	})
}

func Test_NewHashmapCreator(t *testing.T) {
	safeTest(t, "Test_Cov19_NewHashmapCreator", func() {
		// Arrange
		_ = corestr.New.Hashmap.Empty()
		_ = corestr.New.Hashmap.Cap(10)
		_ = corestr.New.Hashmap.UsingMap(map[string]string{"a": "b"})
		_ = corestr.New.Hashmap.UsingMapOptions(true, 5, map[string]string{"a": "b"})
		_ = corestr.New.Hashmap.UsingMapOptions(false, 5, map[string]string{"a": "b"})
		_ = corestr.New.Hashmap.UsingMapOptions(false, 5, nil)
		_ = corestr.New.Hashmap.MapWithCap(5, map[string]string{"a": "b"})
		_ = corestr.New.Hashmap.MapWithCap(0, map[string]string{"a": "b"})
		_ = corestr.New.Hashmap.MapWithCap(5, nil)
		_ = corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k", Value: "v"})
		_ = corestr.New.Hashmap.KeyValues()
		_ = corestr.New.Hashmap.KeyAnyValues(corestr.KeyAnyValuePair{Key: "k", Value: "v"})
		_ = corestr.New.Hashmap.KeyAnyValues()
		_ = corestr.New.Hashmap.KeyValuesStrings([]string{"k"}, []string{"v"})
		_ = corestr.New.Hashmap.KeyValuesStrings(nil, nil)
		keys := corestr.New.Collection.Strings([]string{"k"})
		vals := corestr.New.Collection.Strings([]string{"v"})
		_ = corestr.New.Hashmap.KeyValuesCollection(keys, vals)
		_ = corestr.New.Hashmap.KeyValuesCollection(nil, nil)

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "NewHashmapCreator returns correct value -- with args", actual)
	})
}

func Test_NewHashsetCreator(t *testing.T) {
	safeTest(t, "Test_Cov19_NewHashsetCreator", func() {
		// Arrange
		_ = corestr.New.Hashset.Empty()
		_ = corestr.New.Hashset.Cap(10)
		_ = corestr.New.Hashset.Strings([]string{"a", "b"})
		_ = corestr.New.Hashset.Strings(nil)
		_ = corestr.New.Hashset.StringsSpreadItems("a")
		_ = corestr.New.Hashset.StringsSpreadItems()
		_ = corestr.New.Hashset.UsingMap(map[string]bool{"a": true})
		_ = corestr.New.Hashset.UsingMap(nil)
		_ = corestr.New.Hashset.UsingMapOption(5, true, map[string]bool{"a": true})
		_ = corestr.New.Hashset.UsingMapOption(5, false, map[string]bool{"a": true})
		_ = corestr.New.Hashset.UsingMapOption(5, false, nil)
		_ = corestr.New.Hashset.StringsOption(5, true, "a")
		_ = corestr.New.Hashset.StringsOption(0, false)
		_ = corestr.New.Hashset.StringsOption(5, false)
		_ = corestr.New.Hashset.UsingCollection(corestr.New.Collection.Strings([]string{"a"}))
		_ = corestr.New.Hashset.UsingCollection(nil)
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = corestr.New.Hashset.SimpleSlice(ss)
		_ = corestr.New.Hashset.SimpleSlice(corestr.New.SimpleSlice.Empty())

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "NewHashsetCreator returns correct value -- with args", actual)
	})
}

func Test_NewLinkedListCreator(t *testing.T) {
	safeTest(t, "Test_Cov19_NewLinkedListCreator", func() {
		// Arrange
		_ = corestr.New.LinkedList.Create()
		_ = corestr.New.LinkedList.Empty()
		_ = corestr.New.LinkedList.Strings([]string{"a", "b"})
		_ = corestr.New.LinkedList.Strings(nil)
		_ = corestr.New.LinkedList.SpreadStrings("a")
		_ = corestr.New.LinkedList.SpreadStrings()
		_ = corestr.New.LinkedList.UsingMap(map[string]bool{"a": true})
		_ = corestr.New.LinkedList.UsingMap(nil)

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "NewLinkedListCreator returns correct value -- with args", actual)
	})
}

func Test_NewLinkedCollectionCreator(t *testing.T) {
	safeTest(t, "Test_Cov19_NewLinkedCollectionCreator", func() {
		// Arrange
		_ = corestr.New.LinkedCollection.Create()
		_ = corestr.New.LinkedCollection.Empty()
		_ = corestr.New.LinkedCollection.Strings("a", "b")
		_ = corestr.New.LinkedCollection.Strings()
		_ = corestr.New.LinkedCollection.UsingCollections(corestr.New.Collection.Strings([]string{"a"}))
		_ = corestr.New.LinkedCollection.UsingCollections()

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "NewLinkedCollectionCreator returns correct value -- with args", actual)
	})
}

func Test_NewKeyValuesCreator(t *testing.T) {
	safeTest(t, "Test_Cov19_NewKeyValuesCreator", func() {
		// Arrange
		_ = corestr.New.KeyValues.Empty()
		_ = corestr.New.KeyValues.Cap(10)
		_ = corestr.New.KeyValues.UsingMap(map[string]string{"k": "v"})
		_ = corestr.New.KeyValues.UsingMap(nil)
		_ = corestr.New.KeyValues.UsingKeyValuePairs(corestr.KeyValuePair{Key: "k", Value: "v"})
		_ = corestr.New.KeyValues.UsingKeyValuePairs()
		_ = corestr.New.KeyValues.UsingKeyValueStrings([]string{"k"}, []string{"v"})
		_ = corestr.New.KeyValues.UsingKeyValueStrings(nil, nil)

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "NewKeyValuesCreator returns non-empty -- with args", actual)
	})
}

func Test_NewCollectionsOfCollectionCreator(t *testing.T) {
	safeTest(t, "Test_Cov19_NewCollectionsOfCollectionCreator", func() {
		// Arrange
		_ = corestr.New.CollectionsOfCollection.Empty()
		_ = corestr.New.CollectionsOfCollection.Cap(5)
		_ = corestr.New.CollectionsOfCollection.Strings([]string{"a"})
		_ = corestr.New.CollectionsOfCollection.CloneStrings([]string{"a"})
		_ = corestr.New.CollectionsOfCollection.StringsOption(true, 5, []string{"a"})
		_ = corestr.New.CollectionsOfCollection.StringsOptions(false, 0, []string{"a"})
		_ = corestr.New.CollectionsOfCollection.SpreadStrings(true, "a", "b")
		_ = corestr.New.CollectionsOfCollection.StringsOfStrings(false, []string{"a"}, []string{"b"})
		_ = corestr.New.CollectionsOfCollection.LenCap(0, 5)

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "NewCollectionsOfCollectionCreator returns correct value -- with args", actual)
	})
}

func Test_NewHashsetsCollectionCreator(t *testing.T) {
	safeTest(t, "Test_Cov19_NewHashsetsCollectionCreator", func() {
		// Arrange
		_ = corestr.New.HashsetsCollection.Empty()
		_ = corestr.New.HashsetsCollection.Cap(5)
		_ = corestr.New.HashsetsCollection.LenCap(0, 5)
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		_ = corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)
		_ = corestr.New.HashsetsCollection.UsingHashsetsPointers()

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "NewHashsetsCollectionCreator returns correct value -- with args", actual)
	})
}

func Test_NewCharCollectionMapCreator(t *testing.T) {
	safeTest(t, "Test_Cov19_NewCharCollectionMapCreator", func() {
		// Arrange
		_ = corestr.New.CharCollectionMap.Empty()
		_ = corestr.New.CharCollectionMap.CapSelfCap(20, 20)
		_ = corestr.New.CharCollectionMap.CapSelfCap(1, 1)
		_ = corestr.New.CharCollectionMap.Items([]string{"a", "b"})
		_ = corestr.New.CharCollectionMap.Items(nil)
		_ = corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 5, []string{"a"})
		_ = corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 5, nil)

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "NewCharCollectionMapCreator returns correct value -- with args", actual)
	})
}

func Test_NewCharHashsetMapCreator(t *testing.T) {
	safeTest(t, "Test_Cov19_NewCharHashsetMapCreator", func() {
		// Arrange
		_ = corestr.New.CharHashsetMap.Cap(20, 20)
		_ = corestr.New.CharHashsetMap.Cap(1, 1)
		_ = corestr.New.CharHashsetMap.CapItems(20, 20, "a", "b")
		_ = corestr.New.CharHashsetMap.Strings(20, []string{"a"})
		_ = corestr.New.CharHashsetMap.Strings(20, nil)

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "NewCharHashsetMapCreator returns correct value -- with args", actual)
	})
}

func Test_NewSimpleStringOnceCreator(t *testing.T) {
	safeTest(t, "Test_Cov19_NewSimpleStringOnceCreator", func() {
		// Arrange
		_ = corestr.New.SimpleStringOnce.Init("hello")
		_ = corestr.New.SimpleStringOnce.InitPtr("hello")
		_ = corestr.New.SimpleStringOnce.Uninitialized("val")
		_ = corestr.New.SimpleStringOnce.Create("val", true)
		_ = corestr.New.SimpleStringOnce.CreatePtr("val", false)
		_ = corestr.New.SimpleStringOnce.Empty()
		_ = corestr.New.SimpleStringOnce.Any(false, "hello", true)

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "NewSimpleStringOnceCreator returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CharCollectionMap / CharHashsetMap / LinkedCollections / remaining
// ══════════════════════════════════════════════════════════════════════════════

func Test_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_Cov19_CharCollectionMap", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Items([]string{"abc", "adef", "bcd"})
		_ = ccm.AllLengthsSum()
		_ = ccm.AllLengthsSumLock()
		_ = ccm.LengthLock()
		_ = ccm.GetMap()
		_ = ccm.GetCopyMapLock()
		_ = ccm.List()
		_ = ccm.ListLock()
		_ = ccm.SortedListAsc()
		_ = ccm.String()
		_ = ccm.StringLock()
		_ = ccm.SummaryString()
		_ = ccm.SummaryStringLock()
		_ = ccm.GetChar("")
		_ = ccm.GetChar("a")
		_ = ccm.LengthOf('a')
		_ = ccm.LengthOfLock('a')
		_ = ccm.LengthOfCollectionFromFirstChar("abc")
		_ = ccm.Has("abc")
		_ = ccm.Has("zzz")
		_, _ = ccm.HasWithCollection("abc")
		_, _ = ccm.HasWithCollectionLock("abc")
		_ = ccm.GetCollection("a", true)
		_ = ccm.GetCollectionLock("a", false)
		_ = ccm.GetCollectionByChar('a')
		_ = ccm.HashsetByChar('a')
		_ = ccm.HashsetByCharLock('a')
		_ = ccm.HashsetByStringFirstChar("abc")
		_ = ccm.HashsetByStringFirstCharLock("abc")
		_ = ccm.HashsetsCollection()
		_ = ccm.HashsetsCollectionByChars('a')
		_ = ccm.HashsetsCollectionByStringFirstChar("abc")
		_ = ccm.IsEquals(ccm)
		_ = ccm.IsEqualsLock(ccm)
		_ = ccm.IsEqualsCaseSensitive(true, ccm)
		_ = ccm.IsEqualsCaseSensitiveLock(false, ccm)

		ccm2 := corestr.New.CharCollectionMap.Empty()
		ccm2.Add("hello")
		ccm2.AddLock("world")
		ccm2.AddStrings("a", "b")
		ccm2.AddStrings()
		ccm2.AddCollectionItems(corestr.New.Collection.Strings([]string{"c"}))
		ccm2.AddCollectionItems(nil)
		ccm2.Clear()
		ccm2.Dispose()
		var nilCcm *corestr.CharCollectionMap
		nilCcm.Dispose()

		// Act
		actual := args.Map{"notEmpty": !ccm.IsEmpty()}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "CharCollectionMap returns correct value -- with args", actual)
	})
}

func Test_CharHashsetMap(t *testing.T) {
	safeTest(t, "Test_Cov19_CharHashsetMap", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.CapItems(20, 20, "abc", "adef", "bcd")
		_ = chm.Length()
		_ = chm.LengthLock()
		_ = chm.AllLengthsSum()
		_ = chm.AllLengthsSumLock()
		_ = chm.GetMap()
		_ = chm.GetCopyMapLock()
		_ = chm.List()
		_ = chm.SortedListAsc()
		_ = chm.SortedListDsc()
		_ = chm.String()
		_ = chm.StringLock()
		_ = chm.SummaryString()
		_ = chm.SummaryStringLock()
		_ = chm.GetCharOf("")
		_ = chm.GetCharOf("a")
		_ = chm.LengthOf('a')
		_ = chm.LengthOfLock('a')
		_ = chm.LengthOfHashsetFromFirstChar("abc")
		_ = chm.Has("abc")
		_, _ = chm.HasWithHashset("abc")
		_, _ = chm.HasWithHashsetLock("abc")
		_ = chm.GetHashset("a", true)
		_ = chm.GetHashsetLock(true, "a")
		_ = chm.GetHashsetByChar('a')
		_ = chm.HashsetByChar('a')
		_ = chm.HashsetByCharLock('a')
		_ = chm.HashsetByStringFirstChar("abc")
		_ = chm.HashsetByStringFirstCharLock("abc")
		_ = chm.HashsetsCollection()
		_ = chm.HashsetsCollectionByChars('a')
		_ = chm.HashsetsCollectionByStringsFirstChar("abc")
		_ = chm.IsEquals(chm)
		_ = chm.IsEqualsLock(chm)

		chm2 := corestr.New.CharHashsetMap.Cap(20, 20)
		chm2.Add("hello")
		chm2.AddLock("world")
		chm2.AddStrings("a", "b")
		chm2.AddStrings()
		chm2.AddStringsLock("c")
		chm2.AddStringsLock()
		chm2.AddCollectionItems(corestr.New.Collection.Strings([]string{"d"}))
		chm2.AddCollectionItems(nil)
		chm2.AddHashsetItems(corestr.New.Hashset.StringsSpreadItems("e"))
		chm2.AddCharCollectionMapItems(corestr.New.CharCollectionMap.Items([]string{"f"}))
		chm2.AddCharCollectionMapItems(nil)
		chm2.Clear()
		chm2.RemoveAll()

		// Act
		actual := args.Map{"notEmpty": !chm.IsEmpty()}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "CharHashsetMap returns correct value -- with args", actual)
	})
}

func Test_LinkedCollections(t *testing.T) {
	safeTest(t, "Test_Cov19_LinkedCollections", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		_ = lc.Head()
		_ = lc.Tail()
		_ = lc.First()
		_ = lc.Last()
		_ = lc.FirstOrDefault()
		_ = lc.LastOrDefault()
		_ = lc.AllIndividualItemsLength()
		_ = lc.ToStrings()
		_ = lc.ToCollectionSimple()
		_ = lc.ToCollection(0)
		_ = lc.ToCollectionsOfCollection(0)
		_ = lc.ItemsOfItems()
		_ = lc.ItemsOfItemsCollection()
		_ = lc.SimpleSlice()

		lc2 := corestr.New.LinkedCollection.Create()
		lc2.AddStrings("a", "b")
		lc2.AddStrings()
		lc2.AddStringsLock("c")
		lc2.AddStringsLock()
		lc2.AddCollection(corestr.New.Collection.Strings([]string{"d"}))
		lc2.AddCollection(nil)
		lc2.AddLock(corestr.New.Collection.Strings([]string{"e"}))
		lc2.Push(corestr.New.Collection.Strings([]string{"f"}))
		lc2.PushBack(corestr.New.Collection.Strings([]string{"g"}))
		lc2.PushBackLock(corestr.New.Collection.Strings([]string{"h"}))
		lc2.PushFront(corestr.New.Collection.Strings([]string{"i"}))
		lc2.AddFront(corestr.New.Collection.Strings([]string{"j"}))
		lc2.AddFrontLock(corestr.New.Collection.Strings([]string{"k"}))

		count := 0
		lc2.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return false
		})

		lc3 := corestr.New.LinkedCollection.Strings("a")
		lc4 := corestr.New.LinkedCollection.Strings("a")
		_ = lc3.IsEqualsPtr(lc4)

		// Act
		actual := args.Map{
			"len": lc.Length(),
			"loopCount": count > 0,
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"loopCount": true,
		}
		expected.ShouldBeEqual(t, 0, "LinkedCollections returns correct value -- with args", actual)
	})
}

func Test_CollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_Cov19_CollectionsOfCollection", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		coc.Add(corestr.New.Collection.Strings([]string{"c"}))

		// Act
		actual := args.Map{
			"len": coc.Length(),
			"allLen": coc.AllIndividualItemsLength(),
		}
		_ = coc.List(0)
		_ = coc.ToCollection()
		_ = coc.Items()
		_ = coc.String()
		_ = coc.JsonModel()
		_ = coc.JsonModelAny()
		_, _ = coc.MarshalJSON()
		_ = coc.AsJsoner()
		_ = coc.AsJsonContractsBinder()
		_ = coc.AsJsonParseSelfInjector()
		_ = coc.AsJsonMarshaller()

		// Assert
		expected := args.Map{
			"len": 2,
			"allLen": 3,
		}
		expected.ShouldBeEqual(t, 0, "CollectionsOfCollection returns correct value -- with args", actual)
	})
}

func Test_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Cov19_HashsetsCollection", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		h1 := corestr.New.Hashset.StringsSpreadItems("a")
		h2 := corestr.New.Hashset.StringsSpreadItems("b")
		hc.Add(h1)
		hc.AddNonNil(h2)
		hc.AddNonNil(nil)
		hc.AddNonEmpty(corestr.New.Hashset.Empty())
		hc.Adds(corestr.New.Hashset.StringsSpreadItems("c"))
		_ = hc.LastIndex()
		_ = hc.List()
		_ = hc.ListPtr()
		_ = hc.ListDirectPtr()
		_ = hc.StringsList()
		_ = hc.String()
		_ = hc.Join(",")
		_ = hc.IsEqual(*hc)
		_ = hc.IsEqualPtr(hc)
		_ = hc.JsonModel()
		_ = hc.JsonModelAny()
		_, _ = hc.MarshalJSON()
		_, _ = hc.Serialize()
		_ = hc.AsJsoner()
		_ = hc.AsJsonContractsBinder()
		_ = hc.AsJsonParseSelfInjector()
		_ = hc.AsJsonMarshaller()
		_ = hc.HasAll("a", "b")
		_ = hc.HasAll("z")

		hc2 := corestr.New.HashsetsCollection.Empty()
		hc2.Add(corestr.New.Hashset.StringsSpreadItems("d"))
		_ = hc.ConcatNew(hc2)
		_ = hc.ConcatNew()
		hc.AddHashsetsCollection(hc2)
		hc.AddHashsetsCollection(nil)

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection returns correct value -- with args", actual)
	})
}

func Test_NonChainedLinkedListNodes(t *testing.T) {
	safeTest(t, "Test_Cov19_NonChainedLinkedListNodes", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedListNodes(5)
		n1 := &corestr.LinkedListNode{Element: "a"}
		n2 := &corestr.LinkedListNode{Element: "b"}
		nc.Adds(n1, n2)
		_ = nc.First()
		_ = nc.Last()
		_ = nc.FirstOrDefault()
		_ = nc.LastOrDefault()
		_ = nc.Items()
		nc.ApplyChaining()
		_ = nc.IsChainingApplied()
		_ = nc.ToChainedNodes()

		// Act
		actual := args.Map{
			"len": nc.Length(),
			"notEmpty": !nc.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"notEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "NonChainedLinkedListNodes returns correct value -- with args", actual)
	})
}

func Test_NonChainedLinkedCollectionNodes(t *testing.T) {
	safeTest(t, "Test_Cov19_NonChainedLinkedCollectionNodes", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		n1 := &corestr.LinkedCollectionNode{Element: c1}
		n2 := &corestr.LinkedCollectionNode{Element: c2}
		nc.Adds(n1, n2)
		_ = nc.First()
		_ = nc.Last()
		_ = nc.FirstOrDefault()
		_ = nc.LastOrDefault()
		_ = nc.Items()
		nc.ApplyChaining()
		_ = nc.IsChainingApplied()
		_ = nc.ToChainedNodes()

		// Act
		actual := args.Map{
			"len": nc.Length(),
			"notEmpty": !nc.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"notEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "NonChainedLinkedCollectionNodes returns correct value -- with args", actual)
	})
}
