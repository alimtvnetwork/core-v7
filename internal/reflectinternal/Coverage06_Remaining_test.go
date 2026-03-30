package reflectinternal

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ======= reflectGetter =======

func TestReflectGetter_PublicValuesMapStruct(t *testing.T) {
	type s struct{ Name string }
	m, err := ReflectGetter.PublicValuesMapStruct(s{Name: "test"})
	if err != nil || m["Name"] != "test" {
		t.Fatal("unexpected")
	}
	_, err2 := ReflectGetter.PublicValuesMapStruct(nil)
	if err2 == nil {
		t.Fatal("expected error for nil")
	}
}

func TestReflectGetter_FieldNameWithValuesMap(t *testing.T) {
	type s struct{ Name string }
	m, err := ReflectGetter.FieldNameWithValuesMap(s{Name: "test"})
	if err != nil || m["Name"] != "test" {
		t.Fatal("unexpected")
	}
	_, err2 := ReflectGetter.FieldNameWithValuesMap(nil)
	if err2 == nil {
		t.Fatal("expected error for nil")
	}
}

func TestReflectGetter_FieldNamesMap(t *testing.T) {
	type s struct{ Name string }
	m, err := ReflectGetter.FieldNamesMap(s{Name: "test"})
	if err != nil || !m["Name"] {
		t.Fatal("unexpected")
	}
	_, err2 := ReflectGetter.FieldNamesMap(nil)
	if err2 == nil {
		t.Fatal("expected error for nil")
	}
}

func TestReflectGetter_StructFieldsMap(t *testing.T) {
	type s struct{ Name string }
	m := ReflectGetter.StructFieldsMap(s{Name: "test"})
	if _, ok := m["Name"]; !ok {
		t.Fatal("expected Name")
	}
	m2 := ReflectGetter.StructFieldsMap(nil)
	if len(m2) != 0 {
		t.Fatal("expected empty for nil")
	}
}

func TestReflectGetter_NullFieldsMap(t *testing.T) {
	type s struct {
		Name *string
		Age  int
	}
	m := ReflectGetter.NullFieldsMap(s{})
	if !m["Name"] {
		t.Fatal("expected Name to be null")
	}
	m2 := ReflectGetter.NullFieldsMap(nil)
	if len(m2) != 0 {
		t.Fatal("expected empty for nil")
	}
}

func TestReflectGetter_NullOrZeroFieldsMap(t *testing.T) {
	type s struct {
		Name *string
		Age  int
	}
	m := ReflectGetter.NullOrZeroFieldsMap(s{})
	if len(m) == 0 {
		t.Fatal("expected non-empty")
	}
	m2 := ReflectGetter.NullOrZeroFieldsMap(nil)
	if len(m2) != 0 {
		t.Fatal("expected empty for nil")
	}
}

// ======= reflectGetUsingReflectValue =======

func TestRVGetter_PublicValuesMapStruct(t *testing.T) {
	type s struct{ Name string }
	m, err := ReflectGetterUsingReflectValue.PublicValuesMapStruct(reflect.ValueOf(s{Name: "test"}))
	if err != nil || m["Name"] != "test" {
		t.Fatal("unexpected")
	}
	// not struct
	_, err2 := ReflectGetterUsingReflectValue.PublicValuesMapStruct(reflect.ValueOf(42))
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func TestRVGetter_FieldNameWithTypeMap(t *testing.T) {
	type s struct{ Name string }
	m := ReflectGetterUsingReflectValue.FieldNameWithTypeMap(reflect.ValueOf(s{}))
	if _, ok := m["Name"]; !ok {
		t.Fatal("expected Name")
	}
	// ptr
	m2 := ReflectGetterUsingReflectValue.FieldNameWithTypeMap(reflect.ValueOf(&s{}))
	if _, ok := m2["Name"]; !ok {
		t.Fatal("expected Name through ptr")
	}
	// invalid
	m3 := ReflectGetterUsingReflectValue.FieldNameWithTypeMap(reflect.ValueOf(42))
	if m3 != nil {
		t.Fatal("expected nil for non-struct")
	}
}

func TestRVGetter_FieldNamesMap(t *testing.T) {
	type s struct{ Name string }
	m, err := ReflectGetterUsingReflectValue.FieldNamesMap(reflect.ValueOf(s{}))
	if err != nil || !m["Name"] {
		t.Fatal("unexpected")
	}
	// not struct
	_, err2 := ReflectGetterUsingReflectValue.FieldNamesMap(reflect.ValueOf(42))
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func TestRVGetter_StructFieldsMap(t *testing.T) {
	type s struct{ Name string }
	m := ReflectGetterUsingReflectValue.StructFieldsMap(reflect.ValueOf(s{}))
	if _, ok := m["Name"]; !ok {
		t.Fatal("expected Name")
	}
	// not struct
	m2 := ReflectGetterUsingReflectValue.StructFieldsMap(reflect.ValueOf(42))
	if m2 != nil {
		t.Fatal("expected nil")
	}
}

func TestRVGetter_FieldNameWithValuesMap(t *testing.T) {
	type s struct {
		Name string
		age  int //nolint
	}
	m, err := ReflectGetterUsingReflectValue.FieldNameWithValuesMap(reflect.ValueOf(s{Name: "test", age: 5}))
	if err != nil {
		t.Fatal(err)
	}
	if m["Name"] != "test" {
		t.Fatal("unexpected")
	}
}

// ======= reflectTypeConverter =======

func TestReflectType_SafeName(t *testing.T) {
	s := ReflectType.SafeName("hello")
	if s != "string" {
		t.Fatal("expected string")
	}
	s2 := ReflectType.SafeName(nil)
	if s2 != "" {
		t.Fatal("expected empty")
	}
}

func TestReflectType_Name(t *testing.T) {
	s := ReflectType.Name("hello")
	if s != "string" {
		t.Fatal("expected string")
	}
	s2 := ReflectType.Name(nil)
	if s2 != "" {
		t.Fatal("expected empty")
	}
}

func TestReflectType_NameUsingFmt(t *testing.T) {
	s := ReflectType.NameUsingFmt("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestReflectType_Names(t *testing.T) {
	n := ReflectType.Names(true, "a", 1)
	if len(n) != 2 {
		t.Fatal("expected 2")
	}
	n2 := ReflectType.Names(false, "a")
	if len(n2) != 1 {
		t.Fatal("expected 1")
	}
	n3 := ReflectType.Names(true)
	if len(n3) != 0 {
		t.Fatal("expected 0")
	}
}

func TestReflectType_TypeNamesString(t *testing.T) {
	s := ReflectType.TypeNamesString(true, "a", 1)
	if s == "" {
		t.Fatal("expected non-empty")
	}
	s2 := ReflectType.TypeNamesString(true)
	if s2 != "" {
		t.Fatal("expected empty")
	}
}

func TestReflectType_NamesReferenceString(t *testing.T) {
	s := ReflectType.NamesReferenceString(true, "a")
	if s == "" {
		t.Fatal("expected non-empty")
	}
	s2 := ReflectType.NamesReferenceString(true)
	if s2 != "" {
		t.Fatal("expected empty")
	}
}

func TestReflectType_NamesUsingReflectType(t *testing.T) {
	rt := reflect.TypeOf("hello")
	n := ReflectType.NamesUsingReflectType(true, rt)
	if len(n) != 1 {
		t.Fatal("expected 1")
	}
	n2 := ReflectType.NamesUsingReflectType(false, rt)
	if len(n2) != 1 {
		t.Fatal("expected 1")
	}
	n3 := ReflectType.NamesUsingReflectType(true)
	if len(n3) != 0 {
		t.Fatal("expected 0")
	}
}

func TestReflectType_NamesStringUsingReflectType(t *testing.T) {
	s := ReflectType.NamesStringUsingReflectType(true, reflect.TypeOf("hello"))
	if s == "" {
		t.Fatal("expected non-empty")
	}
	s2 := ReflectType.NamesStringUsingReflectType(true)
	if s2 != "" {
		t.Fatal("expected empty")
	}
}

func TestReflectType_SliceFirstItemTypeName(t *testing.T) {
	s := ReflectType.SliceFirstItemTypeName([]string{})
	if s != "string" {
		t.Fatal("expected string, got:", s)
	}
	s2 := ReflectType.SliceFirstItemTypeName(nil)
	if s2 != "" {
		t.Fatal("expected empty")
	}
}

func TestReflectType_SafeTypeNameOfSliceOrSingle(t *testing.T) {
	s := ReflectType.SafeTypeNameOfSliceOrSingle(true, "hello")
	if s != "string" {
		t.Fatal("expected string")
	}
	s2 := ReflectType.SafeTypeNameOfSliceOrSingle(false, []string{})
	if s2 != "string" {
		t.Fatal("expected string")
	}
}

// ======= reflectUtils =======

func TestUtils_MaxLimit(t *testing.T) {
	if Utils.MaxLimit(10, 5) != 5 {
		t.Fatal("expected 5")
	}
	if Utils.MaxLimit(3, 5) != 3 {
		t.Fatal("expected 3")
	}
	if Utils.MaxLimit(10, -1) != 10 {
		t.Fatal("expected 10 for -1")
	}
}

func TestUtils_AppendArgs(t *testing.T) {
	r := Utils.AppendArgs("first", []any{"a", "b"})
	if len(r) != 3 {
		t.Fatal("expected 3")
	}
	r2 := Utils.AppendArgs("first", []any{})
	if len(r2) != 1 {
		t.Fatal("expected 1")
	}
}

func TestUtils_VerifyReflectTypesAny(t *testing.T) {
	ok, err := Utils.VerifyReflectTypesAny([]any{"a", 1}, []any{"b", 2})
	if !ok || err != nil {
		t.Fatal("same types should match")
	}

	ok2, err2 := Utils.VerifyReflectTypesAny([]any{"a"}, []any{1})
	if ok2 || err2 == nil {
		t.Fatal("different types should not match")
	}

	ok3, err3 := Utils.VerifyReflectTypesAny([]any{"a"}, []any{"b", "c"})
	if ok3 || err3 == nil {
		t.Fatal("different lengths should not match")
	}
}

func TestUtils_VerifyReflectTypes(t *testing.T) {
	ok, err := Utils.VerifyReflectTypes("test",
		[]reflect.Type{reflect.TypeOf("")},
		[]reflect.Type{reflect.TypeOf("")},
	)
	if !ok || err != nil {
		t.Fatal("same types should match")
	}

	ok2, err2 := Utils.VerifyReflectTypes("test",
		[]reflect.Type{reflect.TypeOf("")},
		[]reflect.Type{reflect.TypeOf(1)},
	)
	if ok2 || err2 == nil {
		t.Fatal("different types should not match")
	}
}

func TestUtils_PkgNameOnly(t *testing.T) {
	fn := func() {}
	p := Utils.PkgNameOnly(fn)
	_ = p
}

func TestUtils_FullNameToPkgName(t *testing.T) {
	p := Utils.FullNameToPkgName("github.com/alimtvnetwork/core/codestack.TestFunc")
	_ = p
}

func TestUtils_IsReflectTypeMatch(t *testing.T) {
	ok, _ := Utils.IsReflectTypeMatch(reflect.TypeOf(""), reflect.TypeOf(""))
	if !ok {
		t.Fatal("should match")
	}
}

// ======= Looper =======

func TestLooper_FieldsFor(t *testing.T) {
	type s struct{ Name string }
	err := Looper.FieldsFor(s{Name: "test"}, func(f *reflectmodel.FieldProcessor) error {
		return nil
	})
	_ = err // just exercise it
}

func TestLooper_FieldNames(t *testing.T) {
	type s struct{ Name string }
	names, err := Looper.FieldNames(s{})
	if err != nil || len(names) != 1 {
		t.Fatal("unexpected")
	}
}

func TestLooper_FieldsMap(t *testing.T) {
	type s struct{ Name string }
	m, err := Looper.FieldsMap(s{})
	if err != nil || len(m) != 1 {
		t.Fatal("unexpected")
	}
}

func TestLooper_Slice(t *testing.T) {
	err := Looper.Slice([]int{1, 2}, func(total, index int, item any) error {
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	err2 := Looper.Slice(nil, func(total, index int, item any) error {
		return nil
	})
	if err2 != nil {
		t.Fatal(err2)
	}
}

func TestLooper_Map(t *testing.T) {
	err := Looper.Map(map[string]int{"a": 1}, func(total, index int, key, value any) error {
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	err2 := Looper.Map(nil, func(total, index int, key, value any) error {
		return nil
	})
	if err2 != nil {
		t.Fatal(err2)
	}
}

func TestLooper_ReducePointer(t *testing.T) {
	s := "hello"
	r := Looper.ReducePointer(&s, 3)
	if r == nil || !r.IsValid {
		t.Fatal("expected valid")
	}
}

func TestLooper_ReducePointerDefault(t *testing.T) {
	r := Looper.ReducePointerDefault("hello")
	if r == nil || !r.IsValid {
		t.Fatal("expected valid")
	}
}

// ======= reflectPath =======

func TestPath_CurDir(t *testing.T) {
	d := Path.CurDir()
	if d == "" {
		t.Fatal("expected non-empty")
	}
}

func TestPath_CurDirSkipStack(t *testing.T) {
	d := Path.CurDirSkipStack(0)
	if d == "" {
		t.Fatal("expected non-empty")
	}
}

func TestPath_RepoDir(t *testing.T) {
	d := Path.RepoDir()
	_ = d // may be empty in CI
}

// ======= FileWithLine =======

func TestFileWithLine_All(t *testing.T) {
	f := &FileWithLine{FilePath: "/tmp/test.go", Line: 42}
	if f.FullFilePath() != "/tmp/test.go" {
		t.Fatal("unexpected")
	}
	if f.LineNumber() != 42 {
		t.Fatal("unexpected")
	}
	if f.IsNil() {
		t.Fatal("should not be nil")
	}
	if !f.IsNotNil() {
		t.Fatal("should be not nil")
	}
	if f.String() == "" {
		t.Fatal("expected non-empty")
	}
	_ = f.FileWithLine()
	_ = f.JsonModel()
	_ = f.JsonModelAny()
	_ = f.JsonString()

	var nilF *FileWithLine
	if nilF.String() != "" {
		t.Fatal("expected empty for nil")
	}

	fwl := FileWithLine{FilePath: "/tmp/test.go", Line: 10}
	_ = fwl.StringUsingFmt(func(fl FileWithLine) string {
		return fl.FilePath
	})
}
