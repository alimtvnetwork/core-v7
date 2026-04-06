package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── LinkedList ──

func Test_C33_LL_Empty(t *testing.T) {
	safeTest(t, "Test_C33_LL_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Empty()

		// Act
		actual := args.Map{"result": ll.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": ll.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
	})
}

func Test_C33_LL_Add(t *testing.T) {
	safeTest(t, "Test_C33_LL_Add", func() {
		// Arrange
		ll := corestr.New.LinkedList.Empty()
		ll.Add("a")
		ll.Add("b")

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C33_LL_Head(t *testing.T) {
	safeTest(t, "Test_C33_LL_Head", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ll.Head().Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C33_LL_Tail(t *testing.T) {
	safeTest(t, "Test_C33_LL_Tail", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ll.Tail().Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C33_LL_LengthLock(t *testing.T) {
	safeTest(t, "Test_C33_LL_LengthLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.LengthLock()
	})
}

func Test_C33_LL_IsEquals(t *testing.T) {
	safeTest(t, "Test_C33_LL_IsEquals", func() {
		// Arrange
		a := corestr.New.LinkedList.Strings([]string{"a"})
		b := corestr.New.LinkedList.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C33_LL_IsEqualsWithSensitive(t *testing.T) {
	safeTest(t, "Test_C33_LL_IsEqualsWithSensitive", func() {
		// Arrange
		a := corestr.New.LinkedList.Strings([]string{"A"})
		b := corestr.New.LinkedList.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(b, true)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": a.IsEqualsWithSensitive(b, false)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C33_LL_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_C33_LL_IsEquals_Nil", func() {
		// Arrange
		var a, b *corestr.LinkedList

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(b, true)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C33_LL_AddBack(t *testing.T) {
	safeTest(t, "Test_C33_LL_AddBack", func() {
		ll := corestr.New.LinkedList.Empty()
		ll.PushBack("a")
		ll.PushBack("b")
	})
}

func Test_C33_LL_AddLock(t *testing.T) {
	safeTest(t, "Test_C33_LL_AddLock", func() {
		ll := corestr.New.LinkedList.Empty()
		ll.AddLock("a")
	})
}

func Test_C33_LL_AddCollection(t *testing.T) {
	safeTest(t, "Test_C33_LL_AddCollection", func() {
		ll := corestr.New.LinkedList.Empty()
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		ll.AddCollection(c)
		ll.AddCollection(corestr.New.Collection.Empty())
	})
}

func Test_C33_LL_AddStrings(t *testing.T) {
	safeTest(t, "Test_C33_LL_AddStrings", func() {
		ll := corestr.New.LinkedList.Empty()
		ll.AddStrings([]string{"a"})
		ll.AddStrings(nil)
	})
}

func Test_C33_LL_Adds(t *testing.T) {
	safeTest(t, "Test_C33_LL_Adds", func() {
		ll := corestr.New.LinkedList.Empty()
		ll.Adds("a", "b")
	})
}

func Test_C33_LL_AddAfterNode(t *testing.T) {
	safeTest(t, "Test_C33_LL_AddAfterNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.AddAfterNode(ll.Head(), "x")
	})
}

func Test_C33_LL_Remove(t *testing.T) {
	safeTest(t, "Test_C33_LL_Remove", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByElementValue("b", true, false)
		ll.RemoveNodeByElementValue("a", true, false)
		ll.RemoveNodeByElementValue("c", true, false)
	})
}

func Test_C33_LL_RemoveAt(t *testing.T) {
	safeTest(t, "Test_C33_LL_RemoveAt", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByIndex(1)
		ll.RemoveNodeByIndex(0)
	})
}

func Test_C33_LL_List(t *testing.T) {
	safeTest(t, "Test_C33_LL_List", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		r := ll.List()

		// Act
		actual := args.Map{"result": len(r) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		ll2 := corestr.New.LinkedList.Empty()
		_ = ll2.List()
	})
}

func Test_C33_LL_ListLock(t *testing.T) {
	safeTest(t, "Test_C33_LL_ListLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.ListLock()
	})
}

func Test_C33_LL_ToCollection(t *testing.T) {
	safeTest(t, "Test_C33_LL_ToCollection", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		c := ll.ToCollection(0)

		// Act
		actual := args.Map{"result": c == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil collection", actual)
	})
}

func Test_C33_LL_Loop(t *testing.T) {
	safeTest(t, "Test_C33_LL_Loop", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			return false
		})
	})
}

func Test_C33_LL_String(t *testing.T) {
	safeTest(t, "Test_C33_LL_String", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.String()
		_ = corestr.New.LinkedList.Empty().String()
	})
}

func Test_C33_LL_StringLock(t *testing.T) {
	safeTest(t, "Test_C33_LL_StringLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.StringLock()
	})
}

func Test_C33_LL_Clear(t *testing.T) {
	safeTest(t, "Test_C33_LL_Clear", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.Clear()
	})
}

func Test_C33_LL_JsonMethods(t *testing.T) {
	safeTest(t, "Test_C33_LL_JsonMethods", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.Json()
		_ = ll.JsonPtr()
		_ = ll.JsonModel()
		_ = ll.JsonModelAny()
		_, _ = ll.MarshalJSON()
		_ = ll.AsJsonMarshaller()
	})
}

func Test_C33_LL_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_C33_LL_AddCollectionToNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.AddCollectionToNode(false, ll.Head(), corestr.New.Collection.Strings([]string{"b"}))
	})
}

func Test_C33_LL_AddStringsPtrToNode(t *testing.T) {
	safeTest(t, "Test_C33_LL_AddStringsPtrToNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		items := []string{"b"}
		ll.AddStringsPtrToNode(false, ll.Head(), &items)
	})
}

// ── newLinkedListCreator ──

func Test_C33_NLLC_Empty(t *testing.T)   { _ = corestr.New.LinkedList.Empty() }
func Test_C33_NLLC_Create(t *testing.T)  { _ = corestr.New.LinkedList.Create() }
func Test_C33_NLLC_Strings(t *testing.T) { _ = corestr.New.LinkedList.Strings([]string{"a"}) }
func Test_C33_NLLC_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_C33_NLLC_SpreadStrings", func() {
		// Arrange
		_ = corestr.New.LinkedList.SpreadStrings("a", "b")
		_ = corestr.New.LinkedList.SpreadStrings()
	})
}
func Test_C33_NLLC_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_C33_NLLC_PointerStringsPtr", func() {
		s := "a"
		_ = corestr.New.LinkedList.PointerStringsPtr(&[]*string{&s})
		_ = corestr.New.LinkedList.PointerStringsPtr(nil)
	})
}

// ── LinkedCollections ──

func Test_C33_LC_Empty(t *testing.T) {
	safeTest(t, "Test_C33_LC_Empty", func() {
		lc := corestr.New.LinkedCollection.Empty()

		// Act
		actual := args.Map{"result": lc.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C33_LC_Add(t *testing.T) {
	safeTest(t, "Test_C33_LC_Add", func() {
		lc := corestr.New.LinkedCollection.Empty()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C33_LC_AddLock(t *testing.T) {
	safeTest(t, "Test_C33_LC_AddLock", func() {
		lc := corestr.New.LinkedCollection.Empty()
		lc.AddLock(corestr.New.Collection.Strings([]string{"a"}))
	})
}

func Test_C33_LC_AddAsync(t *testing.T) {
	safeTest(t, "Test_C33_LC_AddAsync", func() {
		lc := corestr.New.LinkedCollection.Empty()
		var wg sync.WaitGroup
		wg.Add(1)
		lc.AddAsync(corestr.New.Collection.Strings([]string{"a"}), &wg)
		wg.Wait()
	})
}

func Test_C33_LC_AddAnother(t *testing.T) {
	safeTest(t, "Test_C33_LC_AddAnother", func() {
		lc := corestr.New.LinkedCollection.Empty()
		other := corestr.New.LinkedCollection.Strings("a")
		lc.AddAnother(other)
		lc.AddAnother(corestr.New.LinkedCollection.Empty())
	})
}

func Test_C33_LC_AddStrings(t *testing.T) {
	safeTest(t, "Test_C33_LC_AddStrings", func() {
		lc := corestr.New.LinkedCollection.Empty()
		lc.AddStrings("a")
	})
}

func Test_C33_LC_AddAfterNode(t *testing.T) {
	safeTest(t, "Test_C33_LC_AddAfterNode", func() {
		lc := corestr.New.LinkedCollection.Empty()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.AddAfterNode(lc.Head(), corestr.New.Collection.Strings([]string{"x"}))
	})
}

func Test_C33_LC_PushBack(t *testing.T) {
	safeTest(t, "Test_C33_LC_PushBack", func() {
		lc := corestr.New.LinkedCollection.Empty()
		lc.PushBack(corestr.New.Collection.Strings([]string{"a"}))
	})
}

func Test_C33_LC_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_C33_LC_AllIndividualItemsLength", func() {
		lc := corestr.New.LinkedCollection.Strings("a", "b")
		_ = lc.AllIndividualItemsLength()
	})
}

func Test_C33_LC_First(t *testing.T) {
	safeTest(t, "Test_C33_LC_First", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.First()
	})
}

func Test_C33_LC_Single(t *testing.T) {
	safeTest(t, "Test_C33_LC_Single", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.Single()
	})
}

func Test_C33_LC_Last(t *testing.T) {
	safeTest(t, "Test_C33_LC_Last", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.Last()
	})
}

func Test_C33_LC_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_C33_LC_FirstOrDefault", func() {
		lc := corestr.New.LinkedCollection.Empty()
		_ = lc.FirstOrDefault()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		_ = lc.FirstOrDefault()
	})
}

func Test_C33_LC_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_C33_LC_LastOrDefault", func() {
		lc := corestr.New.LinkedCollection.Empty()
		_ = lc.LastOrDefault()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		_ = lc.LastOrDefault()
	})
}

func Test_C33_LC_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_C33_LC_RemoveNodeByIndex", func() {
		lc := corestr.New.LinkedCollection.Empty()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.RemoveNodeByIndex(0)
	})
}

func Test_C33_LC_List(t *testing.T) {
	safeTest(t, "Test_C33_LC_List", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.List()
	})
}

func Test_C33_LC_Loop(t *testing.T) {
	safeTest(t, "Test_C33_LC_Loop", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			return false
		})
	})
}

func Test_C33_LC_String(t *testing.T) {
	safeTest(t, "Test_C33_LC_String", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.String()
		_ = corestr.New.LinkedCollection.Empty().String()
	})
}

func Test_C33_LC_Clear(t *testing.T) {
	safeTest(t, "Test_C33_LC_Clear", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		lc.Clear()
	})
}

func Test_C33_LC_JsonMethods(t *testing.T) {
	safeTest(t, "Test_C33_LC_JsonMethods", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.Json()
		_ = lc.JsonPtr()
		_ = lc.JsonModel()
		_ = lc.JsonModelAny()
		_, _ = lc.MarshalJSON()
		_ = lc.AsJsonContractsBinder()
		_ = lc.AsJsoner()
		_ = lc.AsJsonMarshaller()
		_ = lc.AsJsonParseSelfInjector()
	})
}

func Test_C33_LC_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_C33_LC_SimpleSlice", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.SimpleSlice()
	})
}

func Test_C33_LC_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_C33_LC_AddCollectionToNode", func() {
		lc := corestr.New.LinkedCollection.Empty()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.AddCollectionToNode(false, lc.Head(), corestr.New.Collection.Strings([]string{"x"}))
	})
}

// ── newLinkedListCollectionsCreator ──

func Test_C33_NLLCC_Empty(t *testing.T)  { _ = corestr.New.LinkedCollection.Empty() }
func Test_C33_NLLCC_Create(t *testing.T) { _ = corestr.New.LinkedCollection.Create() }
func Test_C33_NLLCC_Strings(t *testing.T) {
	safeTest(t, "Test_C33_NLLCC_Strings", func() {
		_ = corestr.New.LinkedCollection.Strings("a")
	})
}
func Test_C33_NLLCC_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_C33_NLLCC_PointerStringsPtr", func() {
		s := "a"
		_ = corestr.New.LinkedCollection.PointerStringsPtr(&[]*string{&s})
		_ = corestr.New.LinkedCollection.PointerStringsPtr(nil)
	})
}
func Test_C33_NLLCC_UsingCollections(t *testing.T) {
	safeTest(t, "Test_C33_NLLCC_UsingCollections", func() {
		_ = corestr.New.LinkedCollection.UsingCollections(
			corestr.New.Collection.Strings([]string{"a"}),
		)
		_ = corestr.New.LinkedCollection.UsingCollections()
	})
}

// ── LinkedListNode ──

func Test_C33_LLN_HasNext(t *testing.T) {
	safeTest(t, "Test_C33_LLN_HasNext", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		actual := args.Map{"result": ll.Head().HasNext()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C33_LLN_Next(t *testing.T) {
	safeTest(t, "Test_C33_LLN_Next", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		_ = ll.Head().Next()
	})
}

func Test_C33_LLN_EndOfChain(t *testing.T) {
	safeTest(t, "Test_C33_LLN_EndOfChain", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		_, l := ll.Head().EndOfChain()
		actual := args.Map{"result": l != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C33_LLN_LoopEndOfChain(t *testing.T) {
	safeTest(t, "Test_C33_LLN_LoopEndOfChain", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		_, _ = ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			return false
		})
	})
}

func Test_C33_LLN_Clone(t *testing.T) {
	safeTest(t, "Test_C33_LLN_Clone", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.Head().Clone()
	})
}

func Test_C33_LLN_String(t *testing.T) {
	safeTest(t, "Test_C33_LLN_String", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.Head().String()
	})
}

func Test_C33_LLN_AddNext(t *testing.T) {
	safeTest(t, "Test_C33_LLN_AddNext", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.Head().AddNext(ll, "x")
	})
}

func Test_C33_LLN_AddNextNode(t *testing.T) {
	safeTest(t, "Test_C33_LLN_AddNextNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.Head().AddNextNode(ll, &corestr.LinkedListNode{Element: "x"})
	})
}

func Test_C33_LLN_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_C33_LLN_AddStringsToNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.Head().AddStringsToNode(ll, false, []string{"b"})
	})
}

func Test_C33_LLN_AddStringsPtrToNode(t *testing.T) {
	safeTest(t, "Test_C33_LLN_AddStringsPtrToNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		items := []string{"b"}
		ll.Head().AddStringsPtrToNode(ll, false, &items)
	})
}

func Test_C33_LLN_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_C33_LLN_AddCollectionToNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.Head().AddCollectionToNode(ll, false, corestr.New.Collection.Strings([]string{"b"}))
	})
}

// ── LinkedCollectionNode ──

func Test_C33_LCN_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C33_LCN_IsEmpty", func() {
		node := &corestr.LinkedCollectionNode{}
		_ = node.IsEmpty()
	})
}

func Test_C33_LCN_HasElement(t *testing.T) {
	safeTest(t, "Test_C33_LCN_HasElement", func() {
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		_ = node.HasElement()
	})
}

func Test_C33_LCN_AddNext(t *testing.T) {
	safeTest(t, "Test_C33_LCN_AddNext", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		lc.Head().AddNext(lc, corestr.New.Collection.Strings([]string{"x"}))
	})
}

func Test_C33_LCN_AddNextNode(t *testing.T) {
	safeTest(t, "Test_C33_LCN_AddNextNode", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		newNode := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"x"})}
		lc.Head().AddNextNode(lc, newNode)
	})
}

func Test_C33_LCN_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_C33_LCN_AddStringsToNode", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		lc.Head().AddStringsToNode(lc, false, []string{"b"}, false)
	})
}

func Test_C33_LCN_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_C33_LCN_AddCollectionToNode", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		lc.Head().AddCollectionToNode(lc, false, corestr.New.Collection.Strings([]string{"b"}))
	})
}

func Test_C33_LCN_Clone(t *testing.T) {
	safeTest(t, "Test_C33_LCN_Clone", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.Head().Clone()
	})
}

func Test_C33_LCN_String(t *testing.T) {
	safeTest(t, "Test_C33_LCN_String", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.Head().String()
	})
}
