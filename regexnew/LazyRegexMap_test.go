package regexnew

import (
	"regexp"
	"testing"
)

func Test_LazyRegexMap_IsEmpty_NilReceiver(t *testing.T) {
	var m *lazyRegexMap
	if !m.IsEmpty() {
		t.Fatal("expected true for nil receiver")
	}
}

func Test_LazyRegexMap_IsEmptyLock_EmptyMap(t *testing.T) {
	m := &lazyRegexMap{items: map[string]*LazyRegex{}}
	if !m.IsEmptyLock() {
		t.Fatal("expected true for empty map")
	}
}

func Test_LazyRegexMap_HasAnyItem_NilReceiver(t *testing.T) {
	var m *lazyRegexMap
	if m.HasAnyItem() {
		t.Fatal("expected false for nil receiver")
	}
}

func Test_LazyRegexMap_HasAnyItemLock_WithItems(t *testing.T) {
	m := &lazyRegexMap{items: map[string]*LazyRegex{
		"test": {},
	}}
	if !m.HasAnyItemLock() {
		t.Fatal("expected true for non-empty map")
	}
}

func Test_LazyRegexMap_Length_NilReceiver(t *testing.T) {
	var m *lazyRegexMap
	if m.Length() != 0 {
		t.Fatal("expected 0 for nil receiver")
	}
}

func Test_LazyRegexMap_LengthLock(t *testing.T) {
	m := &lazyRegexMap{items: map[string]*LazyRegex{
		"a": {},
		"b": {},
	}}
	if m.LengthLock() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_LazyRegexMap_Has(t *testing.T) {
	m := &lazyRegexMap{items: map[string]*LazyRegex{
		`^\d+$`: {},
	}}
	if !m.Has(`^\d+$`) {
		t.Fatal("expected true")
	}
	if m.Has("missing") {
		t.Fatal("expected false")
	}
}

func Test_LazyRegexMap_HasLock(t *testing.T) {
	m := &lazyRegexMap{items: map[string]*LazyRegex{
		`^\d+$`: {},
	}}
	if !m.HasLock(`^\d+$`) {
		t.Fatal("expected true")
	}
}

func Test_LazyRegexMap_CreateOrExistingLockIf_NoLock(t *testing.T) {
	m := &lazyRegexMap{items: map[string]*LazyRegex{}}
	lr, isExisting := m.CreateOrExistingLockIf(false, `^\d+$`)
	if isExisting {
		t.Fatal("expected new")
	}
	if lr == nil {
		t.Fatal("expected non-nil LazyRegex")
	}
}

func Test_LazyRegexMap_CreateOrExistingLockIf_WithLock(t *testing.T) {
	m := &lazyRegexMap{items: map[string]*LazyRegex{}}
	lr, _ := m.CreateOrExistingLockIf(true, `^\d+$`)
	if lr == nil {
		t.Fatal("expected non-nil LazyRegex")
	}
	// Call again — should get existing
	lr2, isExisting := m.CreateOrExistingLockIf(true, `^\d+$`)
	if !isExisting {
		t.Fatal("expected existing")
	}
	if lr2 == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_LazyRegexMap_createLazyRegex(t *testing.T) {
	m := &lazyRegexMap{items: map[string]*LazyRegex{}}
	lr := m.createLazyRegex(`^\d+$`, func(pattern string) (*regexp.Regexp, error) {
		return regexp.Compile(pattern)
	})
	if lr == nil {
		t.Fatal("expected non-nil")
	}
	if lr.pattern != `^\d+$` {
		t.Fatal("pattern mismatch")
	}
}

func Test_PrettyJson_Nil(t *testing.T) {
	result := prettyJson(nil)
	if result != "" {
		t.Fatal("expected empty string for nil")
	}
}

func Test_PrettyJson_ValidObject(t *testing.T) {
	result := prettyJson(map[string]string{"key": "value"})
	if result == "" {
		t.Fatal("expected non-empty string")
	}
}

func Test_RegExMatchValidationError_NilRegex(t *testing.T) {
	err := regExMatchValidationError("abc", "test", nil, nil)
	if err == nil {
		t.Fatal("expected error for nil regex")
	}
}
