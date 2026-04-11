package reflectmodel

import (
	"reflect"
	"testing"
)

func TestIsNull_Nil(t *testing.T) {
	if !isNull(nil) {
		t.Fatal("nil should be null")
	}
}

func TestIsNull_NilPtr(t *testing.T) {
	var s *string
	if !isNull(s) {
		t.Fatal("nil ptr should be null")
	}
}

func TestIsNull_Value(t *testing.T) {
	if isNull(42) {
		t.Fatal("int should not be null")
	}
}

func TestIsNull_NilMap(t *testing.T) {
	var m map[string]int
	if !isNull(m) {
		t.Fatal("nil map should be null")
	}
}

func TestIsNull_NilSlice(t *testing.T) {
	var s []int
	if !isNull(s) {
		t.Fatal("nil slice should be null")
	}
}

func TestIsNull_NilFunc(t *testing.T) {
	var fn func()
	if !isNull(fn) {
		t.Fatal("nil func should be null")
	}
}

func TestIsNull_NilChan(t *testing.T) {
	var ch chan int
	if !isNull(ch) {
		t.Fatal("nil chan should be null")
	}
}

// rvUtils internal tests

func TestRvUtils_IsNull(t *testing.T) {
	u := rvUtils{}
	if !u.IsNull(nil) {
		t.Fatal("nil should be null")
	}
	if u.IsNull(42) {
		t.Fatal("int should not be null")
	}
	var s *string
	if !u.IsNull(s) {
		t.Fatal("nil ptr should be null")
	}
}

func TestRvUtils_ArgsToReflectValues(t *testing.T) {
	u := rvUtils{}
	rv := u.ArgsToReflectValues([]any{"a", 1})
	if len(rv) != 2 {
		t.Fatal("expected 2")
	}
	rv2 := u.ArgsToReflectValues([]any{})
	if len(rv2) != 0 {
		t.Fatal("expected 0")
	}
}

func TestRvUtils_ReflectValuesToInterfaces(t *testing.T) {
	u := rvUtils{}
	rvs := []reflect.Value{reflect.ValueOf("a"), reflect.ValueOf(1)}
	ifs := u.ReflectValuesToInterfaces(rvs)
	if len(ifs) != 2 {
		t.Fatal("expected 2")
	}
	ifs2 := u.ReflectValuesToInterfaces([]reflect.Value{})
	if len(ifs2) != 0 {
		t.Fatal("expected 0")
	}
}

func TestRvUtils_ReflectValueToAnyValue(t *testing.T) {
	u := rvUtils{}
	v := u.ReflectValueToAnyValue(reflect.ValueOf("hello"))
	if v != "hello" {
		t.Fatal("unexpected")
	}
	s := "hello"
	v2 := u.ReflectValueToAnyValue(reflect.ValueOf(&s))
	if v2 != "hello" {
		t.Fatal("unexpected")
	}
}

func TestRvUtils_InterfacesToTypes(t *testing.T) {
	u := rvUtils{}
	types := u.InterfacesToTypes([]any{"a", 1})
	if len(types) != 2 {
		t.Fatal("expected 2")
	}
	types2 := u.InterfacesToTypes([]any{})
	if len(types2) != 0 {
		t.Fatal("expected 0")
	}
}

func TestRvUtils_InterfacesToTypesNamesWithValues(t *testing.T) {
	u := rvUtils{}
	names := u.InterfacesToTypesNamesWithValues([]any{"a", 1})
	if len(names) != 2 {
		t.Fatal("expected 2")
	}
	names2 := u.InterfacesToTypesNamesWithValues([]any{})
	if len(names2) != 0 {
		t.Fatal("expected 0")
	}
}

func TestRvUtils_IndexToPosition(t *testing.T) {
	u := rvUtils{}
	if u.IndexToPosition(0) != "1st" {
		t.Fatal("expected 1st")
	}
	if u.IndexToPosition(1) != "2nd" {
		t.Fatal("expected 2nd")
	}
	if u.IndexToPosition(2) != "3rd" {
		t.Fatal("expected 3rd")
	}
	if u.IndexToPosition(3) != "4th" {
		t.Fatal("expected 4th")
	}
}

func TestRvUtils_IsReflectTypeMatch(t *testing.T) {
	u := rvUtils{}
	ok, _ := u.IsReflectTypeMatch(reflect.TypeOf(""), reflect.TypeOf(""))
	if !ok {
		t.Fatal("should match")
	}
	ok2, err := u.IsReflectTypeMatch(reflect.TypeOf(""), reflect.TypeOf(1))
	if ok2 || err == nil {
		t.Fatal("should not match")
	}
}

func TestRvUtils_IsReflectTypeMatchAny(t *testing.T) {
	u := rvUtils{}
	ok, _ := u.IsReflectTypeMatchAny("a", "b")
	if !ok {
		t.Fatal("should match")
	}
	ok2, _ := u.IsReflectTypeMatchAny("a", 1)
	if ok2 {
		t.Fatal("should not match")
	}
}

func TestRvUtils_VerifyReflectTypes_LengthMismatch(t *testing.T) {
	u := rvUtils{}
	ok, err := u.VerifyReflectTypes("test",
		[]reflect.Type{reflect.TypeOf("")},
		[]reflect.Type{reflect.TypeOf(""), reflect.TypeOf(1)},
	)
	if ok || err == nil {
		t.Fatal("should fail on length mismatch")
	}
}

func TestRvUtils_VerifyReflectTypes_TypeMismatch(t *testing.T) {
	u := rvUtils{}
	ok, err := u.VerifyReflectTypes("test",
		[]reflect.Type{reflect.TypeOf("")},
		[]reflect.Type{reflect.TypeOf(1)},
	)
	if ok || err == nil {
		t.Fatal("should fail on type mismatch")
	}
}

func TestRvUtils_VerifyReflectTypes_Match(t *testing.T) {
	u := rvUtils{}
	ok, err := u.VerifyReflectTypes("test",
		[]reflect.Type{reflect.TypeOf("")},
		[]reflect.Type{reflect.TypeOf("")},
	)
	if !ok || err != nil {
		t.Fatal("should match")
	}
}

func TestRvUtils_WithSpaces_Empty(t *testing.T) {
	u := rvUtils{}
	result := u.WithSpaces(4)
	if len(result) != 0 {
		t.Fatal("expected empty")
	}
}

func TestRvUtils_WithSpaces_WithLines(t *testing.T) {
	u := rvUtils{}
	result := u.WithSpaces(2, "a", "b")
	if len(result) != 2 {
		t.Fatal("expected 2")
	}
}

func TestRvUtils_PrependWithSpaces(t *testing.T) {
	u := rvUtils{}
	result := u.PrependWithSpaces(4, []string{"existing"}, 2, "prepend")
	if len(result) != 2 {
		t.Fatal("expected 2")
	}
}

func TestRvUtils_PrependWithSpaces_NoSpaces(t *testing.T) {
	u := rvUtils{}
	result := u.PrependWithSpaces(0, []string{"existing"}, 0, "prepend")
	if len(result) != 2 {
		t.Fatal("expected 2")
	}
}
