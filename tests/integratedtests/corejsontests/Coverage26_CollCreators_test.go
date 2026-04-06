package corejsontests

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── newBytesCollectionCreator ──

func Test_C26_NBC_Empty(t *testing.T) { _ = corejson.NewBytesCollection.Empty() }
	// Arrange
func Test_C26_NBC_UsingCap(t *testing.T) { _ = corejson.NewBytesCollection.UsingCap(5) }

func Test_C26_NBC_UnmarshalUsingBytes(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	b, _ := json.Marshal(bc)
	out, err := corejson.NewBytesCollection.UnmarshalUsingBytes(b)
	_ = out; _ = err
}

func Test_C26_NBC_DeserializeUsingBytes_Error(t *testing.T) {
	_, err := corejson.NewBytesCollection.DeserializeUsingBytes([]byte(`invalid`))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C26_NBC_DeserializeUsingResult_HasIssue(t *testing.T) {
	_, err := corejson.NewBytesCollection.DeserializeUsingResult(&corejson.Result{})
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C26_NBC_DeserializeUsingResult_Valid(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	r := corejson.New(bc)
	_, err := corejson.NewBytesCollection.DeserializeUsingResult(r.Ptr())
	_ = err
}

func Test_C26_NBC_AnyItems(t *testing.T) {
	bc, err := corejson.NewBytesCollection.AnyItems("a", "b")
	actual := args.Map{"result": err != nil || bc.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C26_NBC_AnyItems_Error(t *testing.T) {
	_, err := corejson.NewBytesCollection.AnyItems(make(chan int))
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C26_NBC_JsonersPlusCap_Empty(t *testing.T) {
	bc := corejson.NewBytesCollection.JsonersPlusCap(true, 0)
	actual := args.Map{"result": bc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C26_NBC_Jsoners(t *testing.T) {
	bc := corejson.NewBytesCollection.Jsoners()
	actual := args.Map{"result": bc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C26_NBC_Serializers_Empty(t *testing.T) {
	bc := corejson.NewBytesCollection.Serializers()
	actual := args.Map{"result": bc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ── newResultsCollectionCreator ──

func Test_C26_NRC_Empty(t *testing.T) { _ = corejson.NewResultsCollection.Empty() }
func Test_C26_NRC_Default(t *testing.T) { _ = corejson.NewResultsCollection.Default() }

func Test_C26_NRC_UnmarshalUsingBytes(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	b, _ := json.Marshal(rc)
	out, err := corejson.NewResultsCollection.UnmarshalUsingBytes(b)
	_ = out; _ = err
}

func Test_C26_NRC_DeserializeUsingBytes_Error(t *testing.T) {
	_, err := corejson.NewResultsCollection.DeserializeUsingBytes([]byte(`invalid`))
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C26_NRC_DeserializeUsingResult_HasIssue(t *testing.T) {
	_, err := corejson.NewResultsCollection.DeserializeUsingResult(&corejson.Result{})
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C26_NRC_AnyItems(t *testing.T) {
	rc := corejson.NewResultsCollection.AnyItems("a", "b")
	actual := args.Map{"result": rc.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C26_NRC_AnyItemsPlusCap_Empty(t *testing.T) {
	rc := corejson.NewResultsCollection.AnyItemsPlusCap(5)
	actual := args.Map{"result": rc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C26_NRC_AnyItemsPlusCap_Valid(t *testing.T) {
	rc := corejson.NewResultsCollection.AnyItemsPlusCap(2, "a")
	_ = rc
}

func Test_C26_NRC_UsingJsonersOption_Nil(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingJsonersOption(true, 0, nil)
	actual := args.Map{"result": rc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C26_NRC_UsingJsonersNonNull(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingJsonersNonNull(0)
	_ = rc
}

func Test_C26_NRC_UsingJsoners(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingJsoners()
	_ = rc
}

func Test_C26_NRC_UsingResultsPtrPlusCap_Nil(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingResultsPtrPlusCap(5, nil)
	_ = rc
}

func Test_C26_NRC_UsingResultsPtrPlusCap_Empty(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingResultsPtrPlusCap(5)
	_ = rc
}

func Test_C26_NRC_UsingResultsPtrPlusCap_Valid(t *testing.T) {
	r := corejson.New("x")
	rc := corejson.NewResultsCollection.UsingResultsPtrPlusCap(0, r.Ptr())
	_ = rc
}

func Test_C26_NRC_UsingResultsPtr(t *testing.T) {
	r := corejson.New("x")
	rc := corejson.NewResultsCollection.UsingResultsPtr(r.Ptr())
	_ = rc
}

func Test_C26_NRC_UsingResultsPlusCap_Nil(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingResultsPlusCap(5)
	_ = rc
}

func Test_C26_NRC_UsingResultsPlusCap_Empty(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingResultsPlusCap(5)
	_ = rc
}

func Test_C26_NRC_UsingResultsPlusCap_Valid(t *testing.T) {
	r := corejson.New("x")
	rc := corejson.NewResultsCollection.UsingResultsPlusCap(0, r)
	_ = rc
}

func Test_C26_NRC_UsingResults(t *testing.T) {
	r := corejson.New("x")
	rc := corejson.NewResultsCollection.UsingResults(r)
	_ = rc
}

func Test_C26_NRC_Serializers_Empty(t *testing.T) {
	rc := corejson.NewResultsCollection.Serializers()
	actual := args.Map{"result": rc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C26_NRC_SerializerFunctions_Empty(t *testing.T) {
	rc := corejson.NewResultsCollection.SerializerFunctions()
	actual := args.Map{"result": rc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C26_NRC_SerializerFunctions_Valid(t *testing.T) {
	rc := corejson.NewResultsCollection.SerializerFunctions(func() ([]byte, error) {
		return json.Marshal("x")
	})
	_ = rc
}

// ── newResultsPtrCollectionCreator ──

func Test_C26_NRPC_Empty(t *testing.T) { _ = corejson.NewResultsPtrCollection.Empty() }
func Test_C26_NRPC_Default(t *testing.T) { _ = corejson.NewResultsPtrCollection.Default() }
func Test_C26_NRPC_UsingCap(t *testing.T) { _ = corejson.NewResultsPtrCollection.UsingCap(5) }

func Test_C26_NRPC_UnmarshalUsingBytes(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	b, _ := json.Marshal(rpc)
	_, err := corejson.NewResultsPtrCollection.UnmarshalUsingBytes(b)
	_ = err
}

func Test_C26_NRPC_DeserializeUsingBytes_Error(t *testing.T) {
	_, err := corejson.NewResultsPtrCollection.DeserializeUsingBytes([]byte(`invalid`))
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C26_NRPC_DeserializeUsingResult_HasIssue(t *testing.T) {
	_, err := corejson.NewResultsPtrCollection.DeserializeUsingResult(&corejson.Result{})
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C26_NRPC_AnyItemsPlusCap_Empty(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.AnyItemsPlusCap(5)
	_ = rpc
}

func Test_C26_NRPC_AnyItemsPlusCap_Valid(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.AnyItemsPlusCap(0, "a")
	_ = rpc
}

func Test_C26_NRPC_AnyItems(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.AnyItems("a")
	_ = rpc
}

func Test_C26_NRPC_UsingResultsPlusCap_Nil(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingResultsPlusCap(5, nil)
	_ = rpc
}

func Test_C26_NRPC_UsingResultsPlusCap_Empty(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.UsingResultsPlusCap(5)
	_ = rpc
}

func Test_C26_NRPC_UsingResultsPlusCap_Valid(t *testing.T) {
	r := corejson.New("x")
	rpc := corejson.NewResultsPtrCollection.UsingResultsPlusCap(0, r.Ptr())
	_ = rpc
}

func Test_C26_NRPC_UsingResults(t *testing.T) {
	r := corejson.New("x")
	rpc := corejson.NewResultsPtrCollection.UsingResults(r.Ptr())
	_ = rpc
}

func Test_C26_NRPC_JsonersPlusCap_Empty(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.JsonersPlusCap(true, 0)
	_ = rpc
}

func Test_C26_NRPC_Jsoners(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Jsoners()
	_ = rpc
}

func Test_C26_NRPC_Serializers_Empty(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Serializers()
	_ = rpc
}

// ── newMapResultsCreator ──

func Test_C26_NMR_Empty(t *testing.T) { _ = corejson.NewMapResults.Empty() }

func Test_C26_NMR_UnmarshalUsingBytes(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	b, _ := json.Marshal(mr)
	_, err := corejson.NewMapResults.UnmarshalUsingBytes(b)
	_ = err
}

func Test_C26_NMR_DeserializeUsingBytes_Error(t *testing.T) {
	_, err := corejson.NewMapResults.DeserializeUsingBytes([]byte(`invalid`))
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C26_NMR_DeserializeUsingResult_HasIssue(t *testing.T) {
	_, err := corejson.NewMapResults.DeserializeUsingResult(&corejson.Result{})
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C26_NMR_UsingKeyAnyItems_Empty(t *testing.T) {
	mr := corejson.NewMapResults.UsingKeyAnyItems(0)
	_ = mr
}

func Test_C26_NMR_UsingKeyAnyItems_Valid(t *testing.T) {
	mr := corejson.NewMapResults.UsingKeyAnyItems(0, corejson.KeyAny{Key: "k", AnyInf: "v"})
	_ = mr
}

func Test_C26_NMR_UsingMapOptions_Empty(t *testing.T) {
	mr := corejson.NewMapResults.UsingMapOptions(false, false, 0, map[string]corejson.Result{})
	_ = mr
}

func Test_C26_NMR_UsingMapOptions_NoChange(t *testing.T) {
	m := map[string]corejson.Result{"k": corejson.New("v")}
	mr := corejson.NewMapResults.UsingMapOptions(false, false, 0, m)
	_ = mr
}

func Test_C26_NMR_UsingMapOptions_Clone(t *testing.T) {
	m := map[string]corejson.Result{"k": corejson.New("v")}
	mr := corejson.NewMapResults.UsingMapOptions(true, true, 2, m)
	_ = mr
}

func Test_C26_NMR_UsingMapPlusCap(t *testing.T) {
	m := map[string]corejson.Result{"k": corejson.New("v")}
	mr := corejson.NewMapResults.UsingMapPlusCap(2, m)
	_ = mr
}

func Test_C26_NMR_UsingMapPlusCap_Empty(t *testing.T) {
	mr := corejson.NewMapResults.UsingMapPlusCap(2, map[string]corejson.Result{})
	_ = mr
}

func Test_C26_NMR_UsingMapPlusCapClone(t *testing.T) {
	m := map[string]corejson.Result{"k": corejson.New("v")}
	mr := corejson.NewMapResults.UsingMapPlusCapClone(2, m)
	_ = mr
}

func Test_C26_NMR_UsingMapPlusCapClone_Empty(t *testing.T) {
	mr := corejson.NewMapResults.UsingMapPlusCapClone(2, map[string]corejson.Result{})
	_ = mr
}

func Test_C26_NMR_UsingMapPlusCapDeepClone(t *testing.T) {
	m := map[string]corejson.Result{"k": corejson.New("v")}
	mr := corejson.NewMapResults.UsingMapPlusCapDeepClone(2, m)
	_ = mr
}

func Test_C26_NMR_UsingMapPlusCapDeepClone_Empty(t *testing.T) {
	mr := corejson.NewMapResults.UsingMapPlusCapDeepClone(2, map[string]corejson.Result{})
	_ = mr
}

func Test_C26_NMR_UsingMap(t *testing.T) {
	m := map[string]corejson.Result{"k": corejson.New("v")}
	mr := corejson.NewMapResults.UsingMap(m)
	_ = mr
}

func Test_C26_NMR_UsingMap_Empty(t *testing.T) {
	mr := corejson.NewMapResults.UsingMap(map[string]corejson.Result{})
	_ = mr
}

func Test_C26_NMR_UsingMapAnyItemsPlusCap_Empty(t *testing.T) {
	mr := corejson.NewMapResults.UsingMapAnyItemsPlusCap(0, map[string]any{})
	_ = mr
}

func Test_C26_NMR_UsingMapAnyItemsPlusCap_Valid(t *testing.T) {
	mr := corejson.NewMapResults.UsingMapAnyItemsPlusCap(0, map[string]any{"k": "v"})
	_ = mr
}

func Test_C26_NMR_UsingMapAnyItems(t *testing.T) {
	mr := corejson.NewMapResults.UsingMapAnyItems(map[string]any{"k": "v"})
	_ = mr
}

func Test_C26_NMR_UsingKeyWithResultsPlusCap_Nil(t *testing.T) {
	mr := corejson.NewMapResults.UsingKeyWithResultsPlusCap(5)
	_ = mr
}

func Test_C26_NMR_UsingKeyWithResultsPlusCap_Valid(t *testing.T) {
	mr := corejson.NewMapResults.UsingKeyWithResultsPlusCap(0, corejson.KeyWithResult{Key: "k", Result: corejson.New("v")})
	_ = mr
}

func Test_C26_NMR_UsingKeyWithResults(t *testing.T) {
	mr := corejson.NewMapResults.UsingKeyWithResults(corejson.KeyWithResult{Key: "k", Result: corejson.New("v")})
	_ = mr
}

func Test_C26_NMR_UsingKeyJsonersPlusCap_Nil(t *testing.T) {
	mr := corejson.NewMapResults.UsingKeyJsonersPlusCap(5)
	_ = mr
}

func Test_C26_NMR_UsingKeyJsonersPlusCap_Valid(t *testing.T) {
	r := corejson.New("v")
	mr := corejson.NewMapResults.UsingKeyJsonersPlusCap(0, corejson.KeyWithJsoner{Key: "k", Jsoner: &r})
	_ = mr
}

func Test_C26_NMR_UsingKeyJsoners(t *testing.T) {
	r := corejson.New("v")
	mr := corejson.NewMapResults.UsingKeyJsoners(corejson.KeyWithJsoner{Key: "k", Jsoner: &r})
	_ = mr
}

// ── ResultsCollection (key methods) ──

func Test_C26_RC_AddNonNilNonError_Nil(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddNonNilNonError(nil)
	actual := args.Map{"result": rc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C26_RC_AddNonNilNonError_Error(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	rc.AddNonNilNonError(r)
	actual := args.Map{"result": rc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C26_RC_AddNonNilNonError_Valid(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	r := corejson.New("x")
	rc.AddNonNilNonError(r.Ptr())
	actual := args.Map{"result": rc.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ── ResultsPtrCollection (key methods) ──

func Test_C26_RPC_AddNonNilNonError_Nil(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddNonNilNonError(nil)
	actual := args.Map{"result": rpc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C26_RPC_AddNonNilNonError_Valid(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := corejson.New("x")
	rpc.AddNonNilNonError(r.Ptr())
	actual := args.Map{"result": rpc.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}
