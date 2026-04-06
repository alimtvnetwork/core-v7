package corestrtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// LinkedList — Segment 13: Remaining methods + LinkedListNode (L600-1141, Node)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovLL2_01_ToCollection(t *testing.T) {
	safeTest(t, "Test_CovLL2_01_ToCollection", func() {
		ll := corestr.Empty.LinkedList()
		// empty
		col := ll.ToCollection(0)
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
		ll.Adds("a", "b")
		col2 := ll.ToCollection(5)
		if col2.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}
func Test_CovLL2_03_String_StringLock(t *testing.T) {
	safeTest(t, "Test_CovLL2_03_String_StringLock", func() {
		ll := corestr.Empty.LinkedList()
		s := ll.String()
		if s == "" {
			t.Fatal("expected non-empty")
		}
		ll.Adds("a", "b")
		s2 := ll.String()
		if s2 == "" {
			t.Fatal("expected non-empty")
		}
		s3 := ll.StringLock()
		if s3 == "" {
			t.Fatal("expected non-empty")
		}
		// empty lock
		e := corestr.Empty.LinkedList()
		_ = e.StringLock()
	})
}

func Test_CovLL2_04_Join_JoinLock_Joins(t *testing.T) {
	safeTest(t, "Test_CovLL2_04_Join_JoinLock_Joins", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		j := ll.Join(",")
		if j != "a,b" {
			t.Fatalf("expected 'a,b', got '%s'", j)
		}
		jl := ll.JoinLock(",")
		if jl != "a,b" {
			t.Fatal("expected a,b")
		}
		// Joins with extra items
		js := ll.Joins(",", "c")
		if js == "" {
			t.Fatal("expected non-empty")
		}
		// Joins empty list
		e := corestr.Empty.LinkedList()
		_ = e.Joins(",", "x")
	})
}

func Test_CovLL2_05_IsEquals(t *testing.T) {
	safeTest(t, "Test_CovLL2_05_IsEquals", func() {
		a := corestr.Empty.LinkedList()
		a.Adds("a", "b")
		b := corestr.Empty.LinkedList()
		b.Adds("a", "b")
		if !a.IsEquals(b) {
			t.Fatal("expected equal")
		}
		// nil
		if a.IsEquals(nil) {
			t.Fatal("expected false")
		}
		// same pointer
		if !a.IsEquals(a) {
			t.Fatal("expected true")
		}
		// diff length
		c := corestr.Empty.LinkedList()
		c.Add("a")
		if a.IsEquals(c) {
			t.Fatal("expected false")
		}
		// both empty
		e1 := corestr.Empty.LinkedList()
		e2 := corestr.Empty.LinkedList()
		if !e1.IsEquals(e2) {
			t.Fatal("expected true")
		}
		// one empty
		if a.IsEquals(e1) {
			t.Fatal("expected false")
		}
	})
}

func Test_CovLL2_06_IsEqualsWithSensitive(t *testing.T) {
	safeTest(t, "Test_CovLL2_06_IsEqualsWithSensitive", func() {
		a := corestr.Empty.LinkedList()
		a.Adds("Hello", "World")
		b := corestr.Empty.LinkedList()
		b.Adds("hello", "world")
		if a.IsEqualsWithSensitive(b, true) {
			t.Fatal("expected false for case sensitive")
		}
		if !a.IsEqualsWithSensitive(b, false) {
			t.Fatal("expected true for case insensitive")
		}
	})
}

func Test_CovLL2_07_JsonModel_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_CovLL2_07_JsonModel_JsonModelAny", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		m := ll.JsonModel()
		if len(m) != 2 {
			t.Fatal("expected 2")
		}
		_ = ll.JsonModelAny()
	})
}

func Test_CovLL2_08_MarshalJSON_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovLL2_08_MarshalJSON_UnmarshalJSON", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		data, err := ll.MarshalJSON()
		if err != nil {
			t.Fatal("unexpected error")
		}
		ll2 := corestr.Empty.LinkedList()
		err2 := ll2.UnmarshalJSON(data)
		if err2 != nil {
			t.Fatal("unexpected error")
		}
		if ll2.Length() != 2 {
			t.Fatal("expected 2")
		}
		// invalid
		err3 := ll2.UnmarshalJSON([]byte("bad"))
		if err3 == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_CovLL2_09_Clear_RemoveAll(t *testing.T) {
	safeTest(t, "Test_CovLL2_09_Clear_RemoveAll", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		ll.Clear()
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
		// clear already empty
		ll.Clear()
		// RemoveAll
		ll.Add("x")
		ll.RemoveAll()
		if ll.Length() != 0 {
			t.Fatal("expected 0")
		}
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
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		jr := ll.JsonPtr()
		ll2 := corestr.Empty.LinkedList()
		r, err := ll2.ParseInjectUsingJson(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
		if r.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLL2_12_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovLL2_12_ParseInjectUsingJsonMust", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		jr := ll.JsonPtr()
		ll2 := corestr.Empty.LinkedList()
		r := ll2.ParseInjectUsingJsonMust(jr)
		if r.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovLL2_13_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovLL2_13_JsonParseSelfInject", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		jr := ll.JsonPtr()
		ll2 := corestr.Empty.LinkedList()
		err := ll2.JsonParseSelfInject(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
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
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		head := ll.Head()
		if !head.HasNext() {
			t.Fatal("expected next")
		}
		next := head.Next()
		if next.Element != "b" {
			t.Fatal("expected b")
		}
		if next.HasNext() {
			t.Fatal("expected no next")
		}
	})
}

func Test_CovLL2_16_Node_EndOfChain(t *testing.T) {
	safeTest(t, "Test_CovLL2_16_Node_EndOfChain", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		end, length := ll.Head().EndOfChain()
		if length != 3 {
			t.Fatal("expected 3")
		}
		if end.Element != "c" {
			t.Fatal("expected c")
		}
	})
}

func Test_CovLL2_17_Node_LoopEndOfChain(t *testing.T) {
	safeTest(t, "Test_CovLL2_17_Node_LoopEndOfChain", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		count := 0
		end, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})
		if length != 3 || count != 3 {
			t.Fatal("expected 3")
		}
		if end.Element != "c" {
			t.Fatal("expected c")
		}
		// break early
		end2, length2 := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			return true
		})
		if length2 != 1 {
			t.Fatal("expected 1")
		}
		_ = end2
		// break on second
		end3, length3 := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			return arg.Index == 1
		})
		if length3 != 2 {
			t.Fatal("expected 2")
		}
		_ = end3
	})
}

func Test_CovLL2_18_Node_Clone(t *testing.T) {
	safeTest(t, "Test_CovLL2_18_Node_Clone", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		c := ll.Head().Clone()
		if c.Element != "a" {
			t.Fatal("expected a")
		}
		if c.HasNext() {
			t.Fatal("expected no next")
		}
	})
}

func Test_CovLL2_19_Node_AddNext(t *testing.T) {
	safeTest(t, "Test_CovLL2_19_Node_AddNext", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		newNode := ll.Head().AddNext(ll, "b")
		if newNode.Element != "b" {
			t.Fatal("expected b")
		}
		if ll.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovLL2_20_Node_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_CovLL2_20_Node_AddStringsToNode", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		ll.Head().AddStringsToNode(ll, false, []string{"b", "c"})
		if ll.Length() < 3 {
			t.Fatal("expected at least 3")
		}
	})
}
func Test_CovLL2_22_Node_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_CovLL2_22_Node_AddCollectionToNode", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		col := corestr.New.Collection.Strings([]string{"b", "c"})
		ll.Head().AddCollectionToNode(ll, true, col)
		if ll.Length() < 3 {
			t.Fatal("expected at least 3")
		}
	})
}

func Test_CovLL2_23_Node_AddNextNode(t *testing.T) {
	safeTest(t, "Test_CovLL2_23_Node_AddNextNode", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("a")
		newNode := &corestr.LinkedListNode{Element: "b"}
		result := ll.Head().AddNextNode(ll, newNode)
		if result.Element != "b" {
			t.Fatal("expected b")
		}
	})
}

func Test_CovLL2_24_Node_IsEqual(t *testing.T) {
	safeTest(t, "Test_CovLL2_24_Node_IsEqual", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		ll2 := corestr.Empty.LinkedList()
		ll2.Adds("a", "b")
		if !ll.Head().IsEqual(ll2.Head()) {
			t.Fatal("expected equal")
		}
		// same pointer
		if !ll.Head().IsEqual(ll.Head()) {
			t.Fatal("expected equal")
		}
		// nil
		var nilNode *corestr.LinkedListNode
		if !nilNode.IsEqual(nil) {
			t.Fatal("expected true")
		}
		if nilNode.IsEqual(ll.Head()) {
			t.Fatal("expected false")
		}
		// diff element
		ll3 := corestr.Empty.LinkedList()
		ll3.Adds("x", "b")
		if ll.Head().IsEqual(ll3.Head()) {
			t.Fatal("expected false")
		}
	})
}

func Test_CovLL2_25_Node_IsChainEqual(t *testing.T) {
	safeTest(t, "Test_CovLL2_25_Node_IsChainEqual", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("Hello", "World")
		ll2 := corestr.Empty.LinkedList()
		ll2.Adds("hello", "world")
		// case sensitive
		if ll.Head().IsChainEqual(ll2.Head(), true) {
			t.Fatal("expected false")
		}
		// case insensitive
		if !ll.Head().IsChainEqual(ll2.Head(), false) {
			t.Fatal("expected true")
		}
		// nil both
		var n1 *corestr.LinkedListNode
		if !n1.IsChainEqual(nil, true) {
			t.Fatal("expected true")
		}
		// one nil
		if n1.IsChainEqual(ll.Head(), true) {
			t.Fatal("expected false")
		}
		// same pointer
		if !ll.Head().IsChainEqual(ll.Head(), true) {
			t.Fatal("expected true")
		}
	})
}

func Test_CovLL2_26_Node_IsEqualSensitive(t *testing.T) {
	safeTest(t, "Test_CovLL2_26_Node_IsEqualSensitive", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("A")
		ll2 := corestr.Empty.LinkedList()
		ll2.Add("a")
		if ll.Head().IsEqualSensitive(ll2.Head(), true) {
			t.Fatal("expected false")
		}
		if !ll.Head().IsEqualSensitive(ll2.Head(), false) {
			t.Fatal("expected true")
		}
		// same pointer
		if !ll.Head().IsEqualSensitive(ll.Head(), true) {
			t.Fatal("expected true")
		}
		// nil
		var n *corestr.LinkedListNode
		if !n.IsEqualSensitive(nil, true) {
			t.Fatal("expected true")
		}
		if n.IsEqualSensitive(ll.Head(), true) {
			t.Fatal("expected false")
		}
	})
}

func Test_CovLL2_27_Node_IsEqualValue_IsEqualValueSensitive(t *testing.T) {
	safeTest(t, "Test_CovLL2_27_Node_IsEqualValue_IsEqualValueSensitive", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("Hello")
		if !ll.Head().IsEqualValue("Hello") {
			t.Fatal("expected true")
		}
		if ll.Head().IsEqualValue("hello") {
			t.Fatal("expected false")
		}
		if !ll.Head().IsEqualValueSensitive("hello", false) {
			t.Fatal("expected true")
		}
		if !ll.Head().IsEqualValueSensitive("Hello", true) {
			t.Fatal("expected true")
		}
	})
}

func Test_CovLL2_28_Node_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_CovLL2_28_Node_CreateLinkedList", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b", "c")
		newLL := ll.Head().CreateLinkedList()
		if newLL.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_CovLL2_29_Node_String(t *testing.T) {
	safeTest(t, "Test_CovLL2_29_Node_String", func() {
		ll := corestr.Empty.LinkedList()
		ll.Add("hello")
		if ll.Head().String() != "hello" {
			t.Fatal("expected hello")
		}
	})
}
func Test_CovLL2_31_Node_Join_StringList_Print(t *testing.T) {
	safeTest(t, "Test_CovLL2_31_Node_Join_StringList_Print", func() {
		ll := corestr.Empty.LinkedList()
		ll.Adds("a", "b")
		j := ll.Head().Join(",")
		if j != "a,b" {
			t.Fatalf("expected 'a,b', got '%s'", j)
		}
		sl := ll.Head().StringList("header:")
		if sl == "" {
			t.Fatal("expected non-empty")
		}
		// Print just calls slog, ensure no panic
		ll.Head().Print("test:")
	})
}

// suppress unused import
var _ = fmt.Sprintf

