package corestrtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// LinkedList — Segment 13: Remaining methods + LinkedListNode (L600-1141, Node)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovLL2_01_ToCollection(t *testing.T) {
	safeTest(t, "Test_CovLL2_01_ToCollection", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		// empty
		col := ll.ToCollection(0)

		// Act
		actual := args.Map{"result": col.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ll.Adds("a", "b")
		col2 := ll.ToCollection(5)
		actual = args.Map{"result": col2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLL2_02_List_ListPtr_ListLock_ListPtrLock(t *testing.T) {
	safeTest(t, "Test_CovLL2_02_List_ListPtr_ListLock_ListPtrLock", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()

		// Act
		actual := args.Map{"result": len(ll.List()) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ll.Adds("a", "b")
		actual = args.Map{"result": len(ll.List()) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual = args.Map{"result": len(ll.ListPtr()) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual = args.Map{"result": len(ll.ListLock()) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual = args.Map{"result": len(ll.ListPtrLock()) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLL2_03_String_StringLock(t *testing.T) {
	safeTest(t, "Test_CovLL2_03_String_StringLock", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		s := ll.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		ll.Adds("a", "b")
		s2 := ll.String()
		actual = args.Map{"result": s2 == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		s3 := ll.StringLock()
		actual = args.Map{"result": s3 == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		// empty lock
		e := corestr.Empty.LinkedList()
		_ = e.StringLock()
	})
}

func Test_CovLL2_04_Join_JoinLock_Joins(t *testing.T) {
	safeTest(t, "Test_CovLL2_04_Join_JoinLock_Joins", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		j := ll.Join(",")

		// Act
		actual := args.Map{"result": j != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
		jl := ll.JoinLock(",")
		actual = args.Map{"result": jl != "a,b"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
		// Joins with extra items
		js := ll.Joins(",", "c")
		actual = args.Map{"result": js == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		// Joins empty list
		e := corestr.Empty.LinkedList()
		_ = e.Joins(",", "x")
	})
}

func Test_CovLL2_05_IsEquals(t *testing.T) {
	safeTest(t, "Test_CovLL2_05_IsEquals", func() {
		// Arrange
		a := corestr.Empty.LinkedList()
		a.Adds("a", "b")
		b := corestr.Empty.LinkedList()
		b.Adds("a", "b")

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		// nil
		actual = args.Map{"result": a.IsEquals(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// same pointer
		actual = args.Map{"result": a.IsEquals(a)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// diff length
		c := corestr.Empty.LinkedList()
		c.Add("a")
		actual = args.Map{"result": a.IsEquals(c)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// both empty
		e1 := corestr.Empty.LinkedList()
		e2 := corestr.Empty.LinkedList()
		actual = args.Map{"result": e1.IsEquals(e2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// one empty
		actual = args.Map{"result": a.IsEquals(e1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovLL2_06_IsEqualsWithSensitive(t *testing.T) {
	safeTest(t, "Test_CovLL2_06_IsEqualsWithSensitive", func() {
		// Arrange
		a := corestr.Empty.LinkedList()
		a.Adds("Hello", "World")
		b := corestr.Empty.LinkedList()
		b.Adds("hello", "world")

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(b, true)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for case sensitive", actual)
		actual = args.Map{"result": a.IsEqualsWithSensitive(b, false)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for case insensitive", actual)
	})
}

func Test_CovLL2_07_JsonModel_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_CovLL2_07_JsonModel_JsonModelAny", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		m := ll.JsonModel()

		// Act
		actual := args.Map{"result": len(m) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		_ = ll.JsonModelAny()
	})
}

func Test_CovLL2_08_MarshalJSON_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovLL2_08_MarshalJSON_UnmarshalJSON", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		data, err := ll.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		ll2 := corestr.Empty.LinkedList()
		err2 := ll2.UnmarshalJSON(data)
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual = args.Map{"result": ll2.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// invalid
		err3 := ll2.UnmarshalJSON([]byte("bad"))
		actual = args.Map{"result": err3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_CovLL2_09_Clear_RemoveAll(t *testing.T) {
	safeTest(t, "Test_CovLL2_09_Clear_RemoveAll", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		ll.Clear()

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// clear already empty
		ll.Clear()
		// RemoveAll
		ll.Add("x")
		ll.RemoveAll()
		actual = args.Map{"result": ll.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovLL2_10_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_CovLL2_10_Json_JsonPtr", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		_ = ll.Json()
		_ = ll.JsonPtr()
	})
}

func Test_CovLL2_11_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_CovLL2_11_ParseInjectUsingJson", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		jr := ll.JsonPtr()
		ll2 := corestr.Empty.LinkedList()
		r, err := ll2.ParseInjectUsingJson(jr)

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

func Test_CovLL2_12_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovLL2_12_ParseInjectUsingJsonMust", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		jr := ll.JsonPtr()
		ll2 := corestr.Empty.LinkedList()
		r := ll2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"result": r.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovLL2_13_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovLL2_13_JsonParseSelfInject", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		jr := ll.JsonPtr()
		ll2 := corestr.Empty.LinkedList()
		err := ll2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_CovLL2_14_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_CovLL2_14_AsJsonMarshaller", func() {
		ll := corestr.Empty.LinkedList()
		_ = ll.AsJsonMarshaller()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedListNode tests
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovLL2_15_Node_HasNext_Next(t *testing.T) {
	safeTest(t, "Test_CovLL2_15_Node_HasNext_Next", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		head := ll.Head()

		// Act
		actual := args.Map{"result": head.HasNext()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected next", actual)
		next := head.Next()
		actual = args.Map{"result": next.Element != "b"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
		actual = args.Map{"result": next.HasNext()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no next", actual)
	})
}

func Test_CovLL2_16_Node_EndOfChain(t *testing.T) {
	safeTest(t, "Test_CovLL2_16_Node_EndOfChain", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		end, length := ll.Head().EndOfChain()

		// Act
		actual := args.Map{"result": length != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual = args.Map{"result": end.Element != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c", actual)
	})
}

func Test_CovLL2_17_Node_LoopEndOfChain(t *testing.T) {
	safeTest(t, "Test_CovLL2_17_Node_LoopEndOfChain", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		count := 0
		end, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})

		// Act
		actual := args.Map{"result": length != 3 || count != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual = args.Map{"result": end.Element != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c", actual)
		// break early
		end2, length2 := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			return true
		})
		actual = args.Map{"result": length2 != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		_ = end2
		// break on second
		end3, length3 := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			return arg.Index == 1
		})
		actual = args.Map{"result": length3 != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		_ = end3
	})
}

func Test_CovLL2_18_Node_Clone(t *testing.T) {
	safeTest(t, "Test_CovLL2_18_Node_Clone", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		c := ll.Head().Clone()

		// Act
		actual := args.Map{"result": c.Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": c.HasNext()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no next", actual)
	})
}

func Test_CovLL2_19_Node_AddNext(t *testing.T) {
	safeTest(t, "Test_CovLL2_19_Node_AddNext", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		newNode := ll.Head().AddNext(ll, "b")

		// Act
		actual := args.Map{"result": newNode.Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
		actual = args.Map{"result": ll.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLL2_20_Node_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_CovLL2_20_Node_AddStringsToNode", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		ll.Head().AddStringsToNode(ll, false, []string{"b", "c"})

		// Act
		actual := args.Map{"result": ll.Length() < 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	})
}

func Test_CovLL2_21_Node_AddStringsPtrToNode(t *testing.T) {
	safeTest(t, "Test_CovLL2_21_Node_AddStringsPtrToNode", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		items := []string{"b"}
		ll.Head().AddStringsPtrToNode(ll, false, &items)
		// nil
		ll.Head().AddStringsPtrToNode(ll, false, nil)
	})
}

func Test_CovLL2_22_Node_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_CovLL2_22_Node_AddCollectionToNode", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		col := corestr.New.Collection.Strings([]string{"b", "c"})
		ll.Head().AddCollectionToNode(ll, true, col)

		// Act
		actual := args.Map{"result": ll.Length() < 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	})
}

func Test_CovLL2_23_Node_AddNextNode(t *testing.T) {
	safeTest(t, "Test_CovLL2_23_Node_AddNextNode", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		newNode := &corestr.LinkedListNode{Element: "b"}
		result := ll.Head().AddNextNode(ll, newNode)

		// Act
		actual := args.Map{"result": result.Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_CovLL2_24_Node_IsEqual(t *testing.T) {
	safeTest(t, "Test_CovLL2_24_Node_IsEqual", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		ll2 := corestr.Empty.LinkedList()
		ll2.Adds("a", "b")

		// Act
		actual := args.Map{"result": ll.Head().IsEqual(ll2.Head())}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		// same pointer
		actual = args.Map{"result": ll.Head().IsEqual(ll.Head())}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		// nil
		var nilNode *corestr.LinkedListNode
		actual = args.Map{"result": nilNode.IsEqual(nil)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": nilNode.IsEqual(ll.Head())}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// diff element
		ll3 := corestr.Empty.LinkedList()
		ll3.Adds("x", "b")
		actual = args.Map{"result": ll.Head().IsEqual(ll3.Head())}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovLL2_25_Node_IsChainEqual(t *testing.T) {
	safeTest(t, "Test_CovLL2_25_Node_IsChainEqual", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("Hello", "World")
		ll2 := corestr.Empty.LinkedList()
		ll2.Adds("hello", "world")
		// case sensitive

		// Act
		actual := args.Map{"result": ll.Head().IsChainEqual(ll2.Head(), true)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// case insensitive
		actual = args.Map{"result": ll.Head().IsChainEqual(ll2.Head(), false)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// nil both
		var n1 *corestr.LinkedListNode
		actual = args.Map{"result": n1.IsChainEqual(nil, true)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// one nil
		actual = args.Map{"result": n1.IsChainEqual(ll.Head(), true)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// same pointer
		actual = args.Map{"result": ll.Head().IsChainEqual(ll.Head(), true)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovLL2_26_Node_IsEqualSensitive(t *testing.T) {
	safeTest(t, "Test_CovLL2_26_Node_IsEqualSensitive", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("A")
		ll2 := corestr.Empty.LinkedList()
		ll2.Add("a")

		// Act
		actual := args.Map{"result": ll.Head().IsEqualSensitive(ll2.Head(), true)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": ll.Head().IsEqualSensitive(ll2.Head(), false)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// same pointer
		actual = args.Map{"result": ll.Head().IsEqualSensitive(ll.Head(), true)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// nil
		var n *corestr.LinkedListNode
		actual = args.Map{"result": n.IsEqualSensitive(nil, true)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": n.IsEqualSensitive(ll.Head(), true)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovLL2_27_Node_IsEqualValue_IsEqualValueSensitive(t *testing.T) {
	safeTest(t, "Test_CovLL2_27_Node_IsEqualValue_IsEqualValueSensitive", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("Hello")

		// Act
		actual := args.Map{"result": ll.Head().IsEqualValue("Hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ll.Head().IsEqualValue("hello")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": ll.Head().IsEqualValueSensitive("hello", false)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ll.Head().IsEqualValueSensitive("Hello", true)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovLL2_28_Node_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_CovLL2_28_Node_CreateLinkedList", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		newLL := ll.Head().CreateLinkedList()

		// Act
		actual := args.Map{"result": newLL.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_CovLL2_29_Node_String(t *testing.T) {
	safeTest(t, "Test_CovLL2_29_Node_String", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Add("hello")

		// Act
		actual := args.Map{"result": ll.Head().String() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_CovLL2_30_Node_List_ListPtr(t *testing.T) {
	safeTest(t, "Test_CovLL2_30_Node_List_ListPtr", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		items := ll.Head().List()

		// Act
		actual := args.Map{"result": len(items) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		items2 := ll.Head().ListPtr()
		actual = args.Map{"result": len(items2) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovLL2_31_Node_Join_StringList_Print(t *testing.T) {
	safeTest(t, "Test_CovLL2_31_Node_Join_StringList_Print", func() {
		// Arrange
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		j := ll.Head().Join(",")

		// Act
		actual := args.Map{"result": j != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
		sl := ll.Head().StringList("header:")
		actual = args.Map{"result": sl == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		// Print just calls slog, ensure no panic
		ll.Head().Print("test:")
	})
}

// suppress unused import
var _ = fmt.Sprintf

