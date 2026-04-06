package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// corejson Coverage — Segment 5 (Final): Remaining interface methods,
//                     edge branches, nil receivers, empty collections
// ══════════════════════════════════════════════════════════════════════════════

// --- Result edge cases not covered in seg4 ---

func Test_CovJsonS5_R01_New_NilInput(t *testing.T) {
	// Arrange
	r := corejson.New(nil)
	// json.Marshal(nil) → "null" (4 bytes, no error) → HasSafeItems() = true

	// Act
	actual := args.Map{"result": r.HasSafeItems()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true — null is valid JSON bytes", actual)
}

func Test_CovJsonS5_R02_NewPtr_NilInput(t *testing.T) {
	// Arrange
	r := corejson.NewPtr(nil)

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovJsonS5_R03_Result_SafeBytes_Raw(t *testing.T) {
	// Arrange
	r := corejson.New(1)
	_, _ = r.Raw()
	_ = r.SafeBytes()
	// with error
	re := corejson.Result{Error: errors.New("fail")}
	_, err := re.Raw()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	_ = re.SafeBytes()
}

func Test_CovJsonS5_R04_Result_PrettyJsonString(t *testing.T) {
	// Arrange
	r := corejson.New(map[string]int{"a": 1})
	s := r.PrettyJsonString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	var nr *corejson.Result
	s2 := nr.PrettyJsonString()
	actual := args.Map{"result": s2 != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovJsonS5_R05_Result_RawStringMust(t *testing.T) {
	// Arrange
	r := corejson.New(1)
	s := r.RawStringMust()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovJsonS5_R06_Result_HasSafeItems_HasIssuesOrEmpty(t *testing.T) {
	// Arrange
	r := corejson.New(1)

	// Act
	actual := args.Map{"result": r.HasSafeItems()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected has safe items", actual)
	actual := args.Map{"result": r.HasIssuesOrEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no issues", actual)
}

func Test_CovJsonS5_R07_Result_HasError_IsEmptyError(t *testing.T) {
	// Arrange
	r := corejson.New(1)

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": r.IsEmptyError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	re := corejson.Result{Error: errors.New("fail")}
	actual := args.Map{"result": re.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": re.IsEmptyError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_CovJsonS5_R08_Result_IsAnyNull_Nil(t *testing.T) {
	// Arrange
	var nr *corejson.Result

	// Act
	actual := args.Map{"result": nr.IsAnyNull()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
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
	// Arrange
	r := corejson.New(1)

	// Act
	actual := args.Map{"result": r.MeaningfulError() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJsonS5_R13_Result_HasIssuesOrEmpty_WithError(t *testing.T) {
	// Arrange
	re := corejson.Result{Error: errors.New("fail")}

	// Act
	actual := args.Map{"result": re.HasIssuesOrEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovJsonS5_R14_Result_HasIssuesOrEmpty_Empty(t *testing.T) {
	// Arrange
	r := corejson.Result{}

	// Act
	actual := args.Map{"result": r.HasIssuesOrEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
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
	// Arrange
	b, err := corejson.Serialize.Raw(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_CovJsonS5_S03_Serialize_Pretty(t *testing.T) {
	// Arrange
	s := corejson.Serialize.Pretty(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": len(s) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

// --- Deserialize ---

func Test_CovJsonS5_D01_Deserialize_Apply(t *testing.T) {
	// Arrange
	r := corejson.New(map[string]int{"a": 1})
	var m map[string]int
	err := corejson.Deserialize.Apply(&r, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS5_D02_Deserialize_UsingBytes(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw(map[string]int{"a": 1})
	var m map[string]int
	err := corejson.Deserialize.UsingBytes(b, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS5_D03_Deserialize_UsingBytesMust(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw(map[string]int{"a": 1})
	var m map[string]int
	corejson.Deserialize.UsingBytesMust(b, &m)

	// Act
	actual := args.Map{"result": m["a"] != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a=1", actual)
}

// --- CastAny ---

func Test_CovJsonS5_CA01_CastAny_FromToDefault(t *testing.T) {
	// Arrange
	src := map[string]int{"a": 1}
	var dst map[string]int
	err := corejson.CastAny.FromToDefault(src, &dst)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS5_CA02_CastAny_FromToOption(t *testing.T) {
	// Arrange
	src := map[string]int{"a": 1}
	var dst map[string]int
	err := corejson.CastAny.FromToOption(false, src, &dst)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS5_CA03_CastAny_OrDeserializeTo(t *testing.T) {
	// Arrange
	src := map[string]int{"a": 1}
	var dst map[string]int
	err := corejson.CastAny.OrDeserializeTo(src, &dst)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

// --- BytesCollection remaining ---

func Test_CovJsonS5_BC01_BytesCollection_Basic(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.UsingCap(2)
	bc.Add([]byte("a"))
	bc.Add([]byte("b"))

	// Act
	actual := args.Map{"result": bc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": bc.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": bc.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
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

func Test_CovJsonS5_BC04_BytesCollection_Strings(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(2)
	bc.Add([]byte("a"))
	_ = bc.Strings()
	_ = bc.StringsPtr()
	empty := corejson.NewBytesCollection.UsingCap(0)
	_ = empty.Strings()
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
	// Arrange
	bc := corejson.NewBytesCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		bc.Add([]byte{byte(i)})
	}

	// Act
	actual := args.Map{"result": bc.GetPagesSize(3) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
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

func Test_CovJsonS5_BC10_BytesCollection_Clone(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.UsingCap(2)
	bc.Add([]byte("a"))
	_ = bc.Clone(true)
	_ = bc.Clone(false)
	_ = bc.ClonePtr(true)
	_ = bc.ClonePtr(false)
	var nilBC *corejson.BytesCollection

	// Act
	actual := args.Map{"result": nilBC.ClonePtr(true) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
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

func Test_CovJsonS5_BC12_BytesCollection_AddPtr(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(1)
	bc.AddPtr([]byte("a"))
	_ = bc.StringsPtr()
}

// --- funcs ---

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
	// Act
	actual := args.Map{"result": corejson.BytesToString([]byte("hello")) != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	actual := args.Map{"result": corejson.BytesToString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// --- KeyAny / KeyWithResult / KeyWithJsoner ---

func Test_CovJsonS5_K01_KeyAny(t *testing.T) {
	// Arrange
	ka := corejson.KeyAny{Key: "k", AnyInf: 1}

	// Act
	actual := args.Map{"result": ka.Key != "k"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected k", actual)
}

func Test_CovJsonS5_K02_KeyWithResult(t *testing.T) {
	// Arrange
	r := corejson.New(1)
	kwr := corejson.KeyWithResult{Key: "k", Result: r}

	// Act
	actual := args.Map{"result": kwr.Key != "k"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected k", actual)
}

func Test_CovJsonS5_K03_KeyWithJsoner(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.UsingCap(1)
	kwj := corejson.KeyWithJsoner{Key: "k", Jsoner: rc}

	// Act
	actual := args.Map{"result": kwj.Key != "k"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected k", actual)
}

// --- newResultCreator ---

func Test_CovJsonS5_NRC01_UsingErrorStringPtr(t *testing.T) {
	// Arrange
	s := `"hello"`
	r := corejson.NewResult.UsingErrorStringPtr(errors.New("fail"), &s, "test")

	// Act
	actual := args.Map{"result": r.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CovJsonS5_NRC02_UsingBytesPtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingBytesPtr([]byte(`"hello"`))

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected result", actual)
}

// --- Creators for collections ---

func Test_CovJsonS5_NC01_NewMapResults_Empty(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{"result": mr.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovJsonS5_NC02_NewResultsPtrCollection_UsingCap(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.UsingCap(5)

	// Act
	actual := args.Map{"result": rpc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovJsonS5_NC03_NewBytesCollection_Empty(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{"result": bc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}
