package corestr

import (
	"testing"
)

// Internal tests for unexported fields/constants — must remain in source package.

// ── Collection.items ──

func TestCollection_Dispose_ItemsNil(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	c.Dispose()
	if c.items != nil {
		t.Fatal("expected nil")
	}
}

// ── Hashmap.items ──

func TestHashmap_ParseInjectUsingJson_Unexported(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	jr := hm.Json()
	hm2 := &Hashmap{items: map[string]string{}}
	result, err := hm2.ParseInjectUsingJson(&jr)
	if err != nil || result.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func TestHashmap_ParseInjectUsingJsonMust_Unexported(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	jr := hm.Json()
	hm2 := &Hashmap{items: map[string]string{}}
	result := hm2.ParseInjectUsingJsonMust(&jr)
	if result.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func TestHashmap_JsonParseSelfInject_Unexported(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	jr := hm.Json()
	hm2 := &Hashmap{items: map[string]string{}}
	err := hm2.JsonParseSelfInject(&jr)
	if err != nil {
		t.Fatal("unexpected")
	}
}

func TestHashmap_Dispose_Unexported(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	hm.Dispose()
	if hm.items != nil {
		t.Fatal("expected nil")
	}
}

// ── Hashset.items, hasMapUpdated ──

func TestHashset_Dispose_Unexported(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	hs.Dispose()
	if hs.items != nil {
		t.Fatal("expected nil")
	}
}

func TestHashset_Join_Unexported(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	hs.hasMapUpdated = true
	s := hs.Join(",")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestHashset_NonEmptyJoins_Unexported(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	hs.hasMapUpdated = true
	s := hs.NonEmptyJoins(",")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestHashset_NonWhitespaceJoins_Unexported(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	hs.hasMapUpdated = true
	s := hs.NonWhitespaceJoins(",")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestHashset_JsonModel_Unexported(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	m := hs.JsonModel()
	if len(m) != 1 {
		t.Fatal("expected 1")
	}
}

func TestHashset_JsonModelAny_Unexported(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	_ = hs.JsonModelAny()
}

func TestHashset_MarshalJSON_Unexported(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	b, err := hs.MarshalJSON()
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func TestHashset_Json_Unexported(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	r := hs.Json()
	if r.HasError() {
		t.Fatal("unexpected")
	}
}

func TestHashset_ParseInjectUsingJson_Unexported(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	jr := hs.Json()
	hs2 := New.Hashset.Cap(5)
	_, err := hs2.ParseInjectUsingJson(&jr)
	if err != nil {
		t.Fatal("unexpected")
	}
}

func TestHashset_JoinLine_Unexported(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	hs.hasMapUpdated = true
	s := hs.JoinLine()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

// ── CharCollectionMap: emptyChar, items ──

func TestCharCollectionMap_GetChar_EmptyChar(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	if cm.GetChar("") != emptyChar {
		t.Fatal("expected emptyChar for empty string")
	}
}

func TestCharCollectionMap_Dispose_Unexported(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	cm.Dispose()
	if cm.items != nil {
		t.Fatal("expected nil items after dispose")
	}
}

// ── CharHashsetMap: emptyChar ──

func TestCharHashsetMap_GetChar_EmptyChar(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	if hsm.GetChar("") != emptyChar {
		t.Fatal("expected emptyChar")
	}
}

func TestCharHashsetMap_GetCharOf_EmptyChar(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	if hsm.GetCharOf("") != emptyChar {
		t.Fatal("expected emptyChar")
	}
}
