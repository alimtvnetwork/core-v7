package corestr

import (
	"testing"
)

// Internal tests for unexported 'next' field on linked nodes — must remain in source package.

// ── LinkedListNode.next (from C08, C18) ──

func TestLinkedList_AppendChainOfNodes_NewList(t *testing.T) {
	ll := New.LinkedList.Create()
	n1 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	ll.AppendChainOfNodes(n1)
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestLinkedList_AppendChainOfNodes_Existing(t *testing.T) {
	ll := New.LinkedList.SpreadStrings("a")
	chain := &LinkedListNode{Element: "b", next: &LinkedListNode{Element: "c"}}
	ll.AppendChainOfNodes(chain)
	if ll.Length() != 3 {
		t.Fatalf("expected 3, got %d", ll.Length())
	}
}

func TestLinkedListNode_EndOfChain(t *testing.T) {
	n1 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	end, length := n1.EndOfChain()
	if end.Element != "b" || length != 2 {
		t.Fatal("unexpected")
	}
}

func TestLinkedListNode_EndOfChain_ThreeNodes(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b", next: &LinkedListNode{Element: "c"}}}
	end, length := node.EndOfChain()
	if end.Element != "c" || length != 3 {
		t.Fatal("expected c, 3")
	}
}

func TestLinkedListNode_HasNext(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	if !node.HasNext() {
		t.Fatal("expected has next")
	}
	if node.Next().HasNext() {
		t.Fatal("expected no next")
	}
}

func TestLinkedListNode_Clone(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	cloned := node.Clone()
	if cloned.Element != "a" || cloned.HasNext() {
		t.Fatal("expected a without next")
	}
}

func TestLinkedListNode_List(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	list := node.List()
	if len(list) != 2 {
		t.Fatal("expected 2")
	}
}

func TestLinkedListNode_Join(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	if node.Join(",") != "a,b" {
		t.Fatal("expected a,b")
	}
}

func TestLinkedListNode_IsEqual(t *testing.T) {
	n1 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	n2 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	if !n1.IsEqual(n2) {
		t.Fatal("expected equal")
	}
}

func TestLinkedListNode_IsChainEqual(t *testing.T) {
	n1 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	n2 := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	if !n1.IsChainEqual(n2, true) {
		t.Fatal("expected chain equal")
	}
}

func TestLinkedListNode_CreateLinkedList(t *testing.T) {
	n := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	ll := n.CreateLinkedList()
	if ll.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestLinkedListNode_LoopEndOfChain(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b", next: &LinkedListNode{Element: "c"}}}
	count := 0
	end, length := node.LoopEndOfChain(func(arg *LinkedListProcessorParameter) bool {
		count++
		return false
	})
	if end.Element != "c" || length != 3 || count != 3 {
		t.Fatal("expected c, 3, 3")
	}
}

func TestLinkedListNode_LoopEndOfChain_Break(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b"}}
	end, length := node.LoopEndOfChain(func(arg *LinkedListProcessorParameter) bool {
		return true
	})
	if end.Element != "a" || length != 1 {
		t.Fatal("expected a, 1")
	}
}

func TestLinkedListNode_LoopEndOfChain_BreakSecond(t *testing.T) {
	node := &LinkedListNode{Element: "a", next: &LinkedListNode{Element: "b", next: &LinkedListNode{Element: "c"}}}
	end, length := node.LoopEndOfChain(func(arg *LinkedListProcessorParameter) bool {
		return arg.Index == 1
	})
	if end.Element != "b" || length != 2 {
		t.Fatal("expected b, 2")
	}
}

// ── LinkedCollectionNode.next, NonChainedLinkedCollectionNodes.items (from C19) ──

func TestLinkedCollections_AppendChainOfNodes_Existing(t *testing.T) {
	lc := New.LinkedCollection.Strings("a")
	chain := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"b"}),
		next: &LinkedCollectionNode{
			Element: New.Collection.Strings([]string{"c"}),
		},
	}
	lc.AppendChainOfNodes(chain)
	if lc.Length() != 3 {
		t.Fatalf("expected 3, got %d", lc.Length())
	}
}

func TestLinkedCollections_AppendChainOfNodes_Empty(t *testing.T) {
	lc := New.LinkedCollection.Create()
	chain := &LinkedCollectionNode{Element: New.Collection.Strings([]string{"a"})}
	lc.AppendChainOfNodes(chain)
	if lc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestLinkedCollectionNode_HasNext(t *testing.T) {
	node := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	if !node.HasNext() {
		t.Fatal("expected has next")
	}
}

func TestLinkedCollectionNode_EndOfChain(t *testing.T) {
	node := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next: &LinkedCollectionNode{
			Element: New.Collection.Strings([]string{"b"}),
		},
	}
	end, length := node.EndOfChain()
	if length != 2 || end.Element.List()[0] != "b" {
		t.Fatal("expected b, 2")
	}
}

func TestLinkedCollectionNode_Clone(t *testing.T) {
	node := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	cloned := node.Clone()
	if cloned.HasNext() {
		t.Fatal("expected no next")
	}
}

func TestLinkedCollectionNode_List(t *testing.T) {
	node := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	list := node.List()
	if len(list) != 2 {
		t.Fatal("expected 2")
	}
}

func TestLinkedCollectionNode_IsChainEqual(t *testing.T) {
	n1 := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	n2 := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	if !n1.IsChainEqual(n2) {
		t.Fatal("expected chain equal")
	}
}

func TestLinkedCollectionNode_CreateLinkedList(t *testing.T) {
	n := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	lc := n.CreateLinkedList()
	if lc.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestLinkedCollectionNode_LoopEndOfChain(t *testing.T) {
	node := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next: &LinkedCollectionNode{
			Element: New.Collection.Strings([]string{"b"}),
		},
	}
	count := 0
	end, length := node.LoopEndOfChain(func(arg *LinkedCollectionProcessorParameter) bool {
		count++
		return false
	})
	if length != 2 || count != 2 || end == nil {
		t.Fatal("expected 2, 2")
	}
}

func TestLinkedCollectionNode_LoopEndOfChain_Break(t *testing.T) {
	node := &LinkedCollectionNode{
		Element: New.Collection.Strings([]string{"a"}),
		next:    &LinkedCollectionNode{Element: New.Collection.Strings([]string{"b"})},
	}
	end, length := node.LoopEndOfChain(func(arg *LinkedCollectionProcessorParameter) bool {
		return true
	})
	if length != 1 || end == nil {
		t.Fatal("expected 1")
	}
}

func TestNonChainedLinkedCollectionNodes_Basic(t *testing.T) {
	nc := &NonChainedLinkedCollectionNodes{
		items: []*LinkedCollectionNode{
			{Element: New.Collection.Strings([]string{"a"})},
			{Element: New.Collection.Strings([]string{"b"})},
		},
	}
	if nc.IsEmpty() || nc.Length() != 2 || !nc.HasItems() {
		t.Fatal("expected 2 items")
	}
}

func TestNonChainedLinkedCollectionNodes_ApplyChaining(t *testing.T) {
	nc := &NonChainedLinkedCollectionNodes{
		items: []*LinkedCollectionNode{
			{Element: New.Collection.Strings([]string{"a"})},
			{Element: New.Collection.Strings([]string{"b"})},
		},
	}
	nc.ApplyChaining()
	if !nc.IsChainingApplied() {
		t.Fatal("expected chaining applied")
	}
}
