package reflectmodel

import (
	"reflect"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// FieldProcessor — IsFieldType, IsFieldKind
// Covers FieldProcessor.go L16-22
// ══════════════════════════════════════════════════════════════════════════════

type sampleStruct struct {
	Name string
	Age  int
}

func Test_FieldProcessor_IsFieldType(t *testing.T) {
	st := reflect.TypeOf(sampleStruct{})
	field, _ := st.FieldByName("Name")
	fp := &FieldProcessor{
		Name:      "Name",
		Index:     0,
		Field:     field,
		FieldType: field.Type,
	}
	if !fp.IsFieldType(reflect.TypeOf("")) {
		t.Fatal("expected true for string type")
	}
	if fp.IsFieldType(reflect.TypeOf(0)) {
		t.Fatal("expected false for int type")
	}
}

func Test_FieldProcessor_IsFieldKind(t *testing.T) {
	st := reflect.TypeOf(sampleStruct{})
	field, _ := st.FieldByName("Age")
	fp := &FieldProcessor{
		Name:      "Age",
		Index:     1,
		Field:     field,
		FieldType: field.Type,
	}
	if !fp.IsFieldKind(reflect.Int) {
		t.Fatal("expected true for Int kind")
	}
	if fp.IsFieldKind(reflect.String) {
		t.Fatal("expected false for String kind")
	}
}

func Test_FieldProcessor_NilReceiver(t *testing.T) {
	var fp *FieldProcessor
	if fp.IsFieldType(reflect.TypeOf("")) {
		t.Fatal("nil receiver should return false")
	}
	if fp.IsFieldKind(reflect.String) {
		t.Fatal("nil receiver should return false")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — full coverage
// Covers MethodProcessor.go L24-332
// ══════════════════════════════════════════════════════════════════════════════

type methodHost struct{}

func (m methodHost) PublicMethod(s string) (string, error) {
	return s, nil
}

func (m methodHost) NoArgs() string {
	return "ok"
}

func newMethodProcessorInternal(name string) *MethodProcessor {
	t := reflect.TypeOf(methodHost{})
	method, ok := t.MethodByName(name)
	if !ok {
		return nil
	}
	return &MethodProcessor{
		Name:          name,
		Index:         method.Index,
		ReflectMethod: method,
	}
}

func Test_MethodProcessor_HasValidFunc(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	if !mp.HasValidFunc() {
		t.Fatal("expected valid")
	}
	var nilMp *MethodProcessor
	if nilMp.HasValidFunc() {
		t.Fatal("nil should not be valid")
	}
}

func Test_MethodProcessor_GetFuncName(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	if mp.GetFuncName() != "PublicMethod" {
		t.Fatal("wrong name")
	}
}

func Test_MethodProcessor_IsInvalid(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	if mp.IsInvalid() {
		t.Fatal("should not be invalid")
	}
	var nilMp *MethodProcessor
	if !nilMp.IsInvalid() {
		t.Fatal("nil should be invalid")
	}
}

func Test_MethodProcessor_Func(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	fn := mp.Func()
	if fn == nil {
		t.Fatal("expected non-nil func")
	}
}

func Test_MethodProcessor_Func_Nil(t *testing.T) {
	var mp *MethodProcessor
	fn := mp.Func()
	if fn != nil {
		t.Fatal("expected nil for nil receiver")
	}
}

func Test_MethodProcessor_ArgsCount(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	// Method on value receiver: first arg is the receiver itself
	count := mp.ArgsCount()
	if count != 2 { // receiver + string arg
		t.Fatalf("expected 2, got %d", count)
	}
}

func Test_MethodProcessor_ReturnLength(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	rl := mp.ReturnLength()
	if rl != 2 { // string, error
		t.Fatalf("expected 2, got %d", rl)
	}
}

func Test_MethodProcessor_ReturnLength_Nil(t *testing.T) {
	var mp *MethodProcessor
	if mp.ReturnLength() != -1 {
		t.Fatal("expected -1 for nil")
	}
}

func Test_MethodProcessor_IsPublicMethod(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	if !mp.IsPublicMethod() {
		t.Fatal("expected public")
	}
}

func Test_MethodProcessor_IsPrivateMethod(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	if mp.IsPrivateMethod() {
		t.Fatal("expected not private")
	}
}

func Test_MethodProcessor_ArgsLength(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	if mp.ArgsLength() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_MethodProcessor_Invoke(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	host := methodHost{}
	results, err := mp.Invoke(host, "test")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if len(results) != 2 {
		t.Fatal("expected 2 results")
	}
	if results[0] != "test" {
		t.Fatalf("expected 'test', got %v", results[0])
	}
}

func Test_MethodProcessor_Invoke_NilReceiver(t *testing.T) {
	var mp *MethodProcessor
	_, err := mp.Invoke()
	if err == nil {
		t.Fatal("expected error for nil")
	}
}

func Test_MethodProcessor_Invoke_ArgsMismatch(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	_, err := mp.Invoke() // missing args
	if err == nil {
		t.Fatal("expected error for args count mismatch")
	}
}

func Test_MethodProcessor_GetFirstResponseOfInvoke(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	host := methodHost{}
	result, err := mp.GetFirstResponseOfInvoke(host, "hello")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if result != "hello" {
		t.Fatal("expected 'hello'")
	}
}

func Test_MethodProcessor_GetFirstResponseOfInvoke_Error(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	_, err := mp.GetFirstResponseOfInvoke() // missing args
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_MethodProcessor_InvokeResultOfIndex(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	host := methodHost{}
	result, err := mp.InvokeResultOfIndex(0, host, "val")
	if err != nil {
		t.Fatal("unexpected error")
	}
	if result != "val" {
		t.Fatal("expected 'val'")
	}
}

func Test_MethodProcessor_InvokeFirstAndError(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	host := methodHost{}
	first, funcErr, procErr := mp.InvokeFirstAndError(host, "data")
	if procErr != nil {
		t.Fatal("unexpected processing error")
	}
	if funcErr != nil {
		t.Fatal("unexpected func error")
	}
	if first != "data" {
		t.Fatal("expected 'data'")
	}
}

func Test_MethodProcessor_InvokeFirstAndError_ProcessingError(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	_, _, procErr := mp.InvokeFirstAndError() // missing args
	if procErr == nil {
		t.Fatal("expected processing error")
	}
}

func Test_MethodProcessor_InvokeFirstAndError_TooFewReturns(t *testing.T) {
	mp := newMethodProcessorInternal("NoArgs") // returns only 1 value
	host := methodHost{}
	_, _, procErr := mp.InvokeFirstAndError(host)
	if procErr == nil {
		t.Fatal("expected error for too few return args")
	}
}

func Test_MethodProcessor_IsEqual_BothNil(t *testing.T) {
	var a, b *MethodProcessor
	if !a.IsEqual(b) {
		t.Fatal("both nil should be equal")
	}
}

func Test_MethodProcessor_IsEqual_OneNil(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	if mp.IsEqual(nil) {
		t.Fatal("non-nil should not equal nil")
	}
}

func Test_MethodProcessor_IsEqual_SamePtr(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	if !mp.IsEqual(mp) {
		t.Fatal("same pointer should be equal")
	}
}

func Test_MethodProcessor_IsEqual_SameMethod(t *testing.T) {
	a := newMethodProcessorInternal("PublicMethod")
	b := newMethodProcessorInternal("PublicMethod")
	if !a.IsEqual(b) {
		t.Fatal("same method should be equal")
	}
}

func Test_MethodProcessor_IsNotEqual(t *testing.T) {
	a := newMethodProcessorInternal("PublicMethod")
	b := newMethodProcessorInternal("NoArgs")
	if !a.IsNotEqual(b) {
		t.Fatal("different methods should not be equal")
	}
}

func Test_MethodProcessor_GetType(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	ty := mp.GetType()
	if ty == nil {
		t.Fatal("expected non-nil type")
	}
}

func Test_MethodProcessor_GetType_Nil(t *testing.T) {
	var mp *MethodProcessor
	ty := mp.GetType()
	if ty != nil {
		t.Fatal("expected nil for nil receiver")
	}
}

func Test_MethodProcessor_GetOutArgsTypes(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	out := mp.GetOutArgsTypes()
	if len(out) != 2 {
		t.Fatalf("expected 2, got %d", len(out))
	}
	// cached
	out2 := mp.GetOutArgsTypes()
	if len(out2) != 2 {
		t.Fatal("cached should return same")
	}
}

func Test_MethodProcessor_GetOutArgsTypes_Nil(t *testing.T) {
	var mp *MethodProcessor
	out := mp.GetOutArgsTypes()
	if len(out) != 0 {
		t.Fatal("expected empty for nil")
	}
}

func Test_MethodProcessor_GetInArgsTypes(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	in := mp.GetInArgsTypes()
	if len(in) != 2 {
		t.Fatalf("expected 2, got %d", len(in))
	}
	// cached
	in2 := mp.GetInArgsTypes()
	if len(in2) != 2 {
		t.Fatal("cached should return same")
	}
}

func Test_MethodProcessor_GetInArgsTypes_Nil(t *testing.T) {
	var mp *MethodProcessor
	in := mp.GetInArgsTypes()
	if len(in) != 0 {
		t.Fatal("expected empty for nil")
	}
}

func Test_MethodProcessor_GetInArgsTypesNames(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	names := mp.GetInArgsTypesNames()
	if len(names) != 2 {
		t.Fatalf("expected 2, got %d", len(names))
	}
	// cached
	names2 := mp.GetInArgsTypesNames()
	if len(names2) != 2 {
		t.Fatal("cached should return same")
	}
}

func Test_MethodProcessor_GetInArgsTypesNames_Nil(t *testing.T) {
	var mp *MethodProcessor
	names := mp.GetInArgsTypesNames()
	if len(names) != 0 {
		t.Fatal("expected empty for nil")
	}
}

func Test_MethodProcessor_GetInArgsTypesNames_NoArgs(t *testing.T) {
	mp := newMethodProcessorInternal("NoArgs")
	names := mp.GetInArgsTypesNames()
	// NoArgs has 1 in arg (receiver), so let's check
	if len(names) != 1 {
		t.Fatalf("expected 1 (receiver), got %d", len(names))
	}
}

func Test_MethodProcessor_VerifyInArgs(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	host := methodHost{}
	ok, err := mp.VerifyInArgs([]any{host, "test"})
	if !ok || err != nil {
		t.Fatal("expected valid args")
	}
}

func Test_MethodProcessor_VerifyOutArgs(t *testing.T) {
	mp := newMethodProcessorInternal("NoArgs")
	ok, err := mp.VerifyOutArgs([]any{"result"})
	if !ok || err != nil {
		t.Fatal("expected valid out args")
	}
}

func Test_MethodProcessor_ValidateMethodArgs_CountMismatch(t *testing.T) {
	mp := newMethodProcessorInternal("PublicMethod")
	err := mp.ValidateMethodArgs([]any{"too", "many", "args"})
	if err == nil {
		t.Fatal("expected error for args count mismatch")
	}
}
