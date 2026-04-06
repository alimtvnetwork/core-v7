package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ══════════════════════════════════════════════════════════════════════════════
// corejson Coverage — Segment 5 (Final): Remaining interface methods,
//                     edge branches, nil receivers, empty collections
// ══════════════════════════════════════════════════════════════════════════════

// --- Result edge cases not covered in seg4 ---

func Test_CovJsonS5_R01_New_NilInput(t *testing.T) {
	r := corejson.New(nil)
	// json.Marshal(nil) → "null" (4 bytes, no error) → HasSafeItems() = true
	if !r.HasSafeItems() {
		t.Fatal("expected true — null is valid JSON bytes")
	}
}

func Test_CovJsonS5_R02_NewPtr_NilInput(t *testing.T) {
	r := corejson.NewPtr(nil)
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovJsonS5_R03_Result_SafeBytes_Raw(t *testing.T) {
	r := corejson.New(1)
	_, _ = r.Raw()
	_ = r.SafeBytes()
	// with error
	re := corejson.Result{Error: errors.New("fail")}
	_, err := re.Raw()
	if err == nil {
		t.Fatal("expected error")
	}
	_ = re.SafeBytes()
}

func Test_CovJsonS5_R04_Result_PrettyJsonString(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	s := r.PrettyJsonString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
	var nr *corejson.Result
	s2 := nr.PrettyJsonString()
	if s2 != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovJsonS5_R05_Result_RawStringMust(t *testing.T) {
	r := corejson.New(1)
	s := r.RawStringMust()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovJsonS5_R06_Result_HasSafeItems_HasIssuesOrEmpty(t *testing.T) {
	r := corejson.New(1)
	if !r.HasSafeItems() {
		t.Fatal("expected has safe items")
	}
	if r.HasIssuesOrEmpty() {
		t.Fatal("expected no issues")
	}
}

func Test_CovJsonS5_R07_Result_HasError_IsEmptyError(t *testing.T) {
	r := corejson.New(1)
	if r.HasError() {
		t.Fatal("expected false")
	}
	if !r.IsEmptyError() {
		t.Fatal("expected true")
	}
	re := corejson.Result{Error: errors.New("fail")}
	if !re.HasError() {
		t.Fatal("expected true")
	}
	if re.IsEmptyError() {
		t.Fatal("expected false")
	}
}

func Test_CovJsonS5_R08_Result_IsAnyNull_Nil(t *testing.T) {
	var nr *corejson.Result
	if !nr.IsAnyNull() {
		t.Fatal("expected true")
	}
}

func Test_CovJsonS5_R09_Result_HandleError(t *testing.T) {
	r := corejson.New(1)
	r.HandleError() // should not panic
}

func Test_CovJsonS5_R10_Result_AsJsonContractsBinder(t *testing.T) {
	r := corejson.New(1)
	_ = r.AsJsonContractsBinder()
}

func Test_CovJsonS5_R11_Result_AsJsoner(t *testing.T) {
	r := corejson.New(1)
	_ = r.AsJsoner()
}

func Test_CovJsonS5_R12_Result_MeaningfulError_ValidResult(t *testing.T) {
	r := corejson.New(1)
	if r.MeaningfulError() != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJsonS5_R13_Result_HasIssuesOrEmpty_WithError(t *testing.T) {
	re := corejson.Result{Error: errors.New("fail")}
	if !re.HasIssuesOrEmpty() {
		t.Fatal("expected true")
	}
}

func Test_CovJsonS5_R14_Result_HasIssuesOrEmpty_Empty(t *testing.T) {
	r := corejson.Result{}
	if !r.HasIssuesOrEmpty() {
		t.Fatal("expected true")
	}
}

// --- emptyCreator ---

func Test_CovJsonS5_EC01_Empty_Result_ResultPtr(t *testing.T) {
	_ = corejson.Empty.Result()
	_ = corejson.Empty.ResultPtr()
}

func Test_CovJsonS5_EC02_Empty_ResultCollection(t *testing.T) {
	_ = corejson.Empty.ResultsCollection()
}

// --- Serialize ---

func Test_CovJsonS5_S01_Serialize_Raw(t *testing.T) {
	b, err := corejson.Serialize.Raw(map[string]int{"a": 1})
	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_CovJsonS5_S03_Serialize_Pretty(t *testing.T) {
	s := corejson.Serialize.Pretty(map[string]int{"a": 1})
	if len(s) == 0 {
		t.Fatal("expected string")
	}
}

// --- Deserialize ---

func Test_CovJsonS5_D01_Deserialize_Apply(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	var m map[string]int
	err := corejson.Deserialize.Apply(&r, &m)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS5_D02_Deserialize_UsingBytes(t *testing.T) {
	b, _ := corejson.Serialize.Raw(map[string]int{"a": 1})
	var m map[string]int
	err := corejson.Deserialize.UsingBytes(b, &m)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS5_D03_Deserialize_UsingBytesMust(t *testing.T) {
	b, _ := corejson.Serialize.Raw(map[string]int{"a": 1})
	var m map[string]int
	corejson.Deserialize.UsingBytesMust(b, &m)
	if m["a"] != 1 {
		t.Fatal("expected a=1")
	}
}

// --- CastAny ---

func Test_CovJsonS5_CA01_CastAny_FromToDefault(t *testing.T) {
	src := map[string]int{"a": 1}
	var dst map[string]int
	err := corejson.CastAny.FromToDefault(src, &dst)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS5_CA02_CastAny_FromToOption(t *testing.T) {
	src := map[string]int{"a": 1}
	var dst map[string]int
	err := corejson.CastAny.FromToOption(false, src, &dst)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS5_CA03_CastAny_OrDeserializeTo(t *testing.T) {
	src := map[string]int{"a": 1}
	var dst map[string]int
	err := corejson.CastAny.OrDeserializeTo(src, &dst)
	if err != nil {
		t.Fatal("expected no error")
	}
}

// --- BytesCollection remaining ---

func Test_CovJsonS5_BC01_BytesCollection_Basic(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(2)
	bc.Add([]byte("a"))
	bc.Add([]byte("b"))
	if bc.Length() != 2 {
		t.Fatal("expected 2")
	}
	if bc.IsEmpty() {
		t.Fatal("expected false")
	}
	if !bc.HasAnyItem() {
		t.Fatal("expected true")
	}
}

func Test_CovJsonS5_BC02_BytesCollection_AddMethods(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(5)
	bc.Add([]byte("a"))
	bc.Adds([]byte("b"), []byte("c"))
	bc.AddSkipOnNil(nil)
	bc.AddSkipOnNil([]byte("d"))
	bc.AddAny(1)
	bc.AddAny(nil)
	bc.AddAnyItems(1, 2)
}

func Test_CovJsonS5_BC03_BytesCollection_GetMethods(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(2)
	bc.Add([]byte("a"))
	_ = bc.GetAt(0)
	_ = bc.GetAtSafe(0)
	_ = bc.GetAtSafe(-1)
	_ = bc.GetAtSafe(10)
	_ = bc.FirstOrDefault()
	_ = bc.LastOrDefault()
	empty := corejson.NewBytesCollection.UsingCap(0)
	_ = empty.FirstOrDefault()
	_ = empty.LastOrDefault()
}
func Test_CovJsonS5_BC05_BytesCollection_Json(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(1)
	bc.Add([]byte("a"))
	_ = bc.Json()
	_ = bc.JsonPtr()
	_ = bc.JsonModel()
	_ = bc.JsonModelAny()
}

func Test_CovJsonS5_BC06_BytesCollection_ParseInject(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(1)
	bc.Add([]byte("a"))
	jr := bc.JsonPtr()
	bc2 := corejson.NewBytesCollection.UsingCap(0)
	_, _ = bc2.ParseInjectUsingJson(jr)
	bc3 := corejson.NewBytesCollection.UsingCap(0)
	_ = bc3.ParseInjectUsingJsonMust(jr)
}

func Test_CovJsonS5_BC07_BytesCollection_AsInterfaces(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(1)
	_ = bc.AsJsonContractsBinder()
	_ = bc.AsJsoner()
	_ = bc.AsJsonParseSelfInjector()
	_ = bc.JsonParseSelfInject(bc.JsonPtr())
}

func Test_CovJsonS5_BC08_BytesCollection_Paging(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		bc.Add([]byte{byte(i)})
	}
	if bc.GetPagesSize(3) != 4 {
		t.Fatal("expected 4")
	}
	_ = bc.GetPagedCollection(3)
	_ = bc.GetSinglePageCollection(3, 2)
	small := corejson.NewBytesCollection.UsingCap(1)
	small.Add([]byte("a"))
	_ = small.GetPagedCollection(10)
}

func Test_CovJsonS5_BC09_BytesCollection_Clear_Dispose(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(2)
	bc.Add([]byte("a"))
	bc.Clear()
	bc2 := corejson.NewBytesCollection.UsingCap(2)
	bc2.Add([]byte("a"))
	bc2.Dispose()
	var nilBC *corejson.BytesCollection
	nilBC.Clear()
	nilBC.Dispose()
}
func Test_CovJsonS5_BC11_BytesCollection_Take_Limit_Skip(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(5)
	for i := 0; i < 5; i++ {
		bc.Add([]byte{byte(i)})
	}
	_ = bc.Take(3)
	_ = bc.Limit(3)
	_ = bc.Limit(-1)
	_ = bc.Skip(2)
	empty := corejson.NewBytesCollection.UsingCap(0)
	_ = empty.Take(0)
	_ = empty.Limit(0)
	_ = empty.Skip(0)
}
func Test_CovJsonS5_F01_BytesCloneIf(t *testing.T) {
	b := []byte("hello")
	_ = corejson.BytesCloneIf(true, b)
	_ = corejson.BytesCloneIf(false, b)
	_ = corejson.BytesCloneIf(true, nil)
}

func Test_CovJsonS5_F02_BytesDeepClone(t *testing.T) {
	b := []byte("hello")
	_ = corejson.BytesDeepClone(b)
	_ = corejson.BytesDeepClone(nil)
}

func Test_CovJsonS5_F03_BytesToString(t *testing.T) {
	if corejson.BytesToString([]byte("hello")) != "hello" {
		t.Fatal("expected hello")
	}
	if corejson.BytesToString(nil) != "" {
		t.Fatal("expected empty")
	}
}

// --- KeyAny / KeyWithResult / KeyWithJsoner ---

func Test_CovJsonS5_K01_KeyAny(t *testing.T) {
	ka := corejson.KeyAny{Key: "k", AnyInf: 1}
	if ka.Key != "k" {
		t.Fatal("expected k")
	}
}

func Test_CovJsonS5_K02_KeyWithResult(t *testing.T) {
	r := corejson.New(1)
	kwr := corejson.KeyWithResult{Key: "k", Result: r}
	if kwr.Key != "k" {
		t.Fatal("expected k")
	}
}

func Test_CovJsonS5_K03_KeyWithJsoner(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(1)
	kwj := corejson.KeyWithJsoner{Key: "k", Jsoner: rc}
	if kwj.Key != "k" {
		t.Fatal("expected k")
	}
}

// --- newResultCreator ---

func Test_CovJsonS5_NRC01_UsingErrorStringPtr(t *testing.T) {
	s := `"hello"`
	r := corejson.NewResult.UsingErrorStringPtr(errors.New("fail"), &s, "test")
	if r.Error == nil {
		t.Fatal("expected error")
	}
}

func Test_CovJsonS5_NRC02_UsingBytesPtr(t *testing.T) {
	r := corejson.NewResult.UsingBytesPtr([]byte(`"hello"`))
	if r == nil {
		t.Fatal("expected result")
	}
}

// --- Creators for collections ---

func Test_CovJsonS5_NC01_NewMapResults_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	if mr.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovJsonS5_NC02_NewResultsPtrCollection_UsingCap(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingCap(5)
	if rpc.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovJsonS5_NC03_NewBytesCollection_Empty(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	if bc.Length() != 0 {
		t.Fatal("expected 0")
	}
}
