package corestr

import (
	"testing"
)

// Internal tests for unexported functions — must remain in source package.

// ── reflectInterfaceVal ──

func TestReflectInterfaceVal_Nil(t *testing.T) {
	if reflectInterfaceVal(nil) != nil {
		t.Fatal("expected nil")
	}
}

func TestReflectInterfaceVal_Value(t *testing.T) {
	v := reflectInterfaceVal(42)
	if v != 42 {
		t.Fatal("expected 42")
	}
}

func TestReflectInterfaceVal_Ptr(t *testing.T) {
	val := "hello"
	v := reflectInterfaceVal(&val)
	if v != "hello" {
		t.Fatal("expected hello")
	}
}

// ── isCollectionPrecheckEqual ──

func TestIsCollectionPrecheckEqual_BothNil(t *testing.T) {
	var a, b *Collection
	result, handled := isCollectionPrecheckEqual(a, b)
	if !handled || !result {
		t.Fatal("expected true")
	}
}

func TestIsCollectionPrecheckEqual_OneNilLeft(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	result, handled := isCollectionPrecheckEqual(nil, c)
	if r := result; r || !handled {
		t.Fatal("expected false, true")
	}
}

func TestIsCollectionPrecheckEqual_OneNilRight(t *testing.T) {
	a := New.Collection.Strings([]string{"a"})
	result, handled := isCollectionPrecheckEqual(a, nil)
	if handled && result {
		t.Fatal("expected false")
	}
}

func TestIsCollectionPrecheckEqual_SamePtr(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	r, h := isCollectionPrecheckEqual(c, c)
	if !r || !h {
		t.Fatal("expected true, true")
	}
}

func TestIsCollectionPrecheckEqual_BothEmpty(t *testing.T) {
	e1 := New.Collection.Empty()
	e2 := New.Collection.Empty()
	r, h := isCollectionPrecheckEqual(e1, e2)
	if !r || !h {
		t.Fatal("expected true, true")
	}
}

func TestIsCollectionPrecheckEqual_OneEmpty(t *testing.T) {
	e := New.Collection.Empty()
	c := New.Collection.Strings([]string{"a"})
	r, h := isCollectionPrecheckEqual(e, c)
	if r || !h {
		t.Fatal("expected false, true")
	}
}

func TestIsCollectionPrecheckEqual_DiffLength(t *testing.T) {
	c1 := New.Collection.Strings([]string{"a"})
	c2 := New.Collection.Strings([]string{"a", "b"})
	r, h := isCollectionPrecheckEqual(c1, c2)
	if r || !h {
		t.Fatal("expected false, true")
	}
}

func TestIsCollectionPrecheckEqual_SameLenNotHandled(t *testing.T) {
	c1 := New.Collection.Strings([]string{"a"})
	c2 := New.Collection.Strings([]string{"b"})
	_, h := isCollectionPrecheckEqual(c1, c2)
	if h {
		t.Fatal("expected not handled")
	}
}
