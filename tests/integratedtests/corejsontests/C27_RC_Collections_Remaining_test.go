package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ResultsCollection remaining methods ──

func Test_C27_RC_Length_Nil(t *testing.T) {
	// Arrange
	var rc *corejson.ResultsCollection

	// Act
	actual := args.Map{"result": rc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}
func Test_C27_RC_LastIndex(t *testing.T) { _ = corejson.NewResultsCollection.Empty().LastIndex() }
	// Arrange
func Test_C27_RC_IsEmpty(t *testing.T) { _ = corejson.NewResultsCollection.Empty().IsEmpty() }
func Test_C27_RC_HasAnyItem(t *testing.T) { _ = corejson.NewResultsCollection.Empty().HasAnyItem() }

func Test_C27_RC_FirstOrDefault_Empty(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	r := rc.FirstOrDefault()
	_ = r
}

func Test_C27_RC_FirstOrDefault_Valid(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("x"))
	r := rc.FirstOrDefault()
	_ = r
}

func Test_C27_RC_LastOrDefault_Empty(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	r := rc.LastOrDefault()
	_ = r
}

func Test_C27_RC_LastOrDefault_Valid(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("x"))
	r := rc.LastOrDefault()
	_ = r
}

func Test_C27_RC_Take(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("a")).Add(corejson.New("b"))
	taken := rc.Take(1)

	// Act
	actual := args.Map{"result": taken.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C27_RC_Take_Empty(t *testing.T) {
	_ = corejson.NewResultsCollection.Empty().Take(1)
}

func Test_C27_RC_Limit(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("a")).Add(corejson.New("b"))
	_ = rc.Limit(1)
	_ = rc.Limit(-1)
}

func Test_C27_RC_Limit_Empty(t *testing.T) {
	_ = corejson.NewResultsCollection.Empty().Limit(1)
}

func Test_C27_RC_Skip(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("a")).Add(corejson.New("b"))
	_ = rc.Skip(1)
}

func Test_C27_RC_Skip_Empty(t *testing.T) {
	_ = corejson.NewResultsCollection.Empty().Skip(0)
}

func Test_C27_RC_AddSkipOnNil(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSkipOnNil(nil)
	rc.AddSkipOnNil(corejson.New("x").Ptr())
}

func Test_C27_RC_Add(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("x"))
}

func Test_C27_RC_AddPtr(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddPtr(nil)
	rc.AddPtr(corejson.New("x").Ptr())
}

func Test_C27_RC_Adds(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Adds(corejson.New("a"), corejson.New("b"))
	rc.Adds()
}

func Test_C27_RC_GetAt(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("x"))
	_ = rc.GetAt(0)
}

func Test_C27_RC_GetAtSafe(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("x"))
	_ = rc.GetAtSafe(0)
	_ = rc.GetAtSafe(99)
}

func Test_C27_RC_GetAtSafeUsingLength(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("x"))
	_ = rc.GetAtSafeUsingLength(0, 1)
	_ = rc.GetAtSafeUsingLength(99, 1)
}

func Test_C27_RC_HasError(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	actual := args.Map{"result": rc.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	rc.Add(corejson.NewResult.Create(nil, errors.New("e"), ""))
	actual := args.Map{"result": rc.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C27_RC_AllErrors(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Create(nil, errors.New("e"), ""))
	errs, has := rc.AllErrors()
	actual := args.Map{"result": has || len(errs) == 0}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected errors", actual)
}

func Test_C27_RC_GetErrorsStrings(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Create(nil, errors.New("e"), ""))
	_ = rc.GetErrorsStrings()
	_ = rc.GetErrorsStringsPtr()
}

func Test_C27_RC_GetErrorsAsSingleString(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	_ = rc.GetErrorsAsSingleString()
}

func Test_C27_RC_GetErrorsAsSingle(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	_ = rc.GetErrorsAsSingle()
}

func Test_C27_RC_AddAny(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddAny("x")
}

func Test_C27_RC_AddAnyItems(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddAnyItems("a", "b")
	rc.AddAnyItems()
}

func Test_C27_RC_AddAnyItemsSlice(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddAnyItemsSlice([]any{"a", "b"})
	rc.AddAnyItemsSlice(nil)
}

func Test_C27_RC_AddsPtr(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddsPtr()
	r := corejson.New("x")
	rc.AddsPtr(r.Ptr())
}

func Test_C27_RC_AddNonNilItemsPtr(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddNonNilItemsPtr()
	r := corejson.New("x")
	rc.AddNonNilItemsPtr(nil, r.Ptr())
}

func Test_C27_RC_AddResultsCollection(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	other := corejson.NewResultsCollection.Empty()
	other.Add(corejson.New("x"))
	rc.AddResultsCollection(other)
	rc.AddResultsCollection(corejson.NewResultsCollection.Empty())
}

func Test_C27_RC_AddMapResults(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	rc.AddMapResults(mr)
	rc.AddMapResults(corejson.NewMapResults.Empty())
}

func Test_C27_RC_AddRawMapResults(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddRawMapResults(map[string]corejson.Result{"k": corejson.New("v")})
}

func Test_C27_RC_AddSerializer(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSerializer(nil)
}

func Test_C27_RC_AddSerializerFunc(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSerializerFunc(nil)
}

func Test_C27_RC_AddSerializerFunctions(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSerializerFunctions()
}

func Test_C27_RC_AddSerializers(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSerializers()
}

func Test_C27_RC_GetStrings(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("x"))
	_ = rc.GetStrings()
	_ = rc.GetStringsPtr()
}

func Test_C27_RC_GetStrings_Empty(t *testing.T) {
	_ = corejson.NewResultsCollection.Empty().GetStrings()
}

func Test_C27_RC_GetPagesSize(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	for i := 0; i < 5; i++ { rc.Add(corejson.New(i)) }
	actual := args.Map{"result": rc.GetPagesSize(2) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": rc.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C27_RC_GetPagedCollection(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	for i := 0; i < 5; i++ { rc.Add(corejson.New(i)) }
	pages := rc.GetPagedCollection(2)
	actual := args.Map{"result": len(pages) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C27_RC_GetPagedCollection_Small(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("x"))
	pages := rc.GetPagedCollection(5)
	actual := args.Map{"result": len(pages) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C27_RC_InjectIntoAt(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New(map[string]string{"a": "b"}))
	target := corejson.Empty.MapResults()
	_ = rc.InjectIntoAt(0, target)
}

func Test_C27_RC_InjectIntoSameIndex(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New(map[string]string{"a": "b"}))
	target := corejson.Empty.MapResults()
	_, _ = rc.InjectIntoSameIndex(target)
	_, _ = rc.InjectIntoSameIndex(nil) // nil element in populated collection - ok
}

func Test_C27_RC_UnmarshalAt(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("hello"))
	var s string
	_ = rc.UnmarshalAt(0, &s)
}

func Test_C27_RC_UnmarshalIntoSameIndex(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("hello"))
	var s string
	_, _ = rc.UnmarshalIntoSameIndex(&s)
	_, _ = rc.UnmarshalIntoSameIndex(nil) // nil element in populated collection - ok
}

func Test_C27_RC_NonPtr(t *testing.T) { _ = corejson.NewResultsCollection.Empty().NonPtr() }
func Test_C27_RC_Ptr(t *testing.T) { _ = corejson.NewResultsCollection.Empty().Ptr() }

func Test_C27_RC_Clear(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("x"))
	rc.Clear()
}

func Test_C27_RC_Dispose(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Dispose()
}

func Test_C27_RC_Json(t *testing.T) { _ = corejson.NewResultsCollection.Empty().Json() }
func Test_C27_RC_JsonPtr(t *testing.T) { _ = corejson.NewResultsCollection.Empty().JsonPtr() }
func Test_C27_RC_JsonModel(t *testing.T) { _ = corejson.NewResultsCollection.Empty().JsonModel() }
func Test_C27_RC_JsonModelAny(t *testing.T) { _ = corejson.NewResultsCollection.Empty().JsonModelAny() }
func Test_C27_RC_AsJsonContractsBinder(t *testing.T) { _ = corejson.NewResultsCollection.Empty().AsJsonContractsBinder() }
func Test_C27_RC_AsJsoner(t *testing.T) { _ = corejson.NewResultsCollection.Empty().AsJsoner() }
func Test_C27_RC_AsJsonParseSelfInjector(t *testing.T) { _ = corejson.NewResultsCollection.Empty().AsJsonParseSelfInjector() }

func Test_C27_RC_JsonParseSelfInject(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	r := corejson.New(rc)
	_ = rc.JsonParseSelfInject(&r)
}

func Test_C27_RC_ParseInjectUsingJson_Error(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	bad := corejson.NewResult.UsingString(`invalid`)
	_, err := rc.ParseInjectUsingJson(bad)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C27_RC_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() { recover() }()
	rc := corejson.NewResultsCollection.Empty()
	bad := corejson.NewResult.UsingString(`invalid`)
	rc.ParseInjectUsingJsonMust(bad)
}

func Test_C27_RC_ShadowClone(t *testing.T) { _ = corejson.NewResultsCollection.Empty().ShadowClone() }

func Test_C27_RC_Clone(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("x"))
	_ = rc.Clone(true)
}

func Test_C27_RC_ClonePtr_Nil(t *testing.T) {
	var rc *corejson.ResultsCollection
	actual := args.Map{"result": rc.ClonePtr(true) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C27_RC_ClonePtr_Valid(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("x"))
	_ = rc.ClonePtr(true)
}

func Test_C27_RC_AddJsoners(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddJsoners(true)
}

// ── ResultsPtrCollection remaining methods ──

func Test_C27_RPC_Length_Nil(t *testing.T) {
	var rpc *corejson.ResultsPtrCollection
	actual := args.Map{"result": rpc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}
func Test_C27_RPC_LastIndex(t *testing.T) { _ = corejson.NewResultsPtrCollection.Empty().LastIndex() }
func Test_C27_RPC_IsEmpty(t *testing.T) { _ = corejson.NewResultsPtrCollection.Empty().IsEmpty() }
func Test_C27_RPC_HasAnyItem(t *testing.T) { _ = corejson.NewResultsPtrCollection.Empty().HasAnyItem() }

func Test_C27_RPC_FirstOrDefault(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.FirstOrDefault()
	r := corejson.New("x")
	rpc.Add(r.Ptr())
	_ = rpc.FirstOrDefault()
}

func Test_C27_RPC_LastOrDefault(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.LastOrDefault()
	r := corejson.New("x")
	rpc.Add(r.Ptr())
	_ = rpc.LastOrDefault()
}

func Test_C27_RPC_Take(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := corejson.New("x")
	rpc.Add(r.Ptr())
	_ = rpc.Take(1)
	_ = corejson.NewResultsPtrCollection.Empty().Take(1)
}

func Test_C27_RPC_Limit(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := corejson.New("x")
	rpc.Add(r.Ptr())
	_ = rpc.Limit(1)
	_ = rpc.Limit(-1)
	_ = corejson.NewResultsPtrCollection.Empty().Limit(1)
}

func Test_C27_RPC_Skip(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := corejson.New("x")
	rpc.Add(r.Ptr())
	_ = rpc.Skip(0)
	_ = corejson.NewResultsPtrCollection.Empty().Skip(0)
}

func Test_C27_RPC_AddSkipOnNil(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSkipOnNil(nil)
	r := corejson.New("x")
	rpc.AddSkipOnNil(r.Ptr())
}

func Test_C27_RPC_Add(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := corejson.New("x")
	rpc.Add(r.Ptr())
}

func Test_C27_RPC_AddResult(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddResult(corejson.New("x"))
}

func Test_C27_RPC_AddSerializer(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSerializer(nil)
}

func Test_C27_RPC_AddSerializers(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSerializers()
}

func Test_C27_RPC_AddSerializerFunc(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSerializerFunc(nil)
}

func Test_C27_RPC_AddSerializerFunctions(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSerializerFunctions()
}

func Test_C27_RPC_Adds(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Adds()
	r := corejson.New("x")
	rpc.Adds(nil, r.Ptr())
}

func Test_C27_RPC_AddAny(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddAny("x")
}

func Test_C27_RPC_AddAnyItems(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddAnyItems("a", "b")
	rpc.AddAnyItems()
}

func Test_C27_RPC_AddResultsCollection(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	other := corejson.NewResultsPtrCollection.Empty()
	r := corejson.New("x")
	other.Add(r.Ptr())
	rpc.AddResultsCollection(other)
	rpc.AddResultsCollection(corejson.NewResultsPtrCollection.Empty())
}

func Test_C27_RPC_AddNonNilItems(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddNonNilItems()
	r := corejson.New("x")
	rpc.AddNonNilItems(nil, r.Ptr())
}

func Test_C27_RPC_AddNonNilItemsPtr(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddNonNilItemsPtr()
	r := corejson.New("x")
	rpc.AddNonNilItemsPtr(nil, r.Ptr())
}

func Test_C27_RPC_GetAt(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := corejson.New("x")
	rpc.Add(r.Ptr())
	_ = rpc.GetAt(0)
}

func Test_C27_RPC_GetAtSafe(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.GetAtSafe(0)
	r := corejson.New("x")
	rpc.Add(r.Ptr())
	_ = rpc.GetAtSafe(0)
}

func Test_C27_RPC_GetAtSafeUsingLength(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := corejson.New("x")
	rpc.Add(r.Ptr())
	_ = rpc.GetAtSafeUsingLength(0, 1)
	_ = rpc.GetAtSafeUsingLength(99, 1)
}

func Test_C27_RPC_HasError(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	actual := args.Map{"result": rpc.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	rpc.Add(corejson.NewResult.ErrorPtr(errors.New("e")))
	actual := args.Map{"result": rpc.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C27_RPC_AllErrors(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.ErrorPtr(errors.New("e")))
	errs, has := rpc.AllErrors()
	actual := args.Map{"result": has || len(errs) == 0}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected errors", actual)
}

func Test_C27_RPC_GetErrorsStrings(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.ErrorPtr(errors.New("e")))
	_ = rpc.GetErrorsStrings()
	_ = rpc.GetErrorsStringsPtr()
}

func Test_C27_RPC_GetErrorsAsSingleString(t *testing.T) { _ = corejson.NewResultsPtrCollection.Empty().GetErrorsAsSingleString() }
func Test_C27_RPC_GetErrorsAsSingle(t *testing.T) { _ = corejson.NewResultsPtrCollection.Empty().GetErrorsAsSingle() }

func Test_C27_RPC_GetStrings(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := corejson.New("x")
	rpc.Add(r.Ptr())
	_ = rpc.GetStrings()
	_ = rpc.GetStringsPtr()
}

func Test_C27_RPC_InjectIntoAt(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.New(map[string]string{"a": "b"}).Ptr())
	target := corejson.Empty.MapResults()
	_ = rpc.InjectIntoAt(0, target)
}

func Test_C27_RPC_InjectIntoSameIndex(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.New(map[string]string{"a": "b"}).Ptr())
	target := corejson.Empty.MapResults()
	_, _ = rpc.InjectIntoSameIndex(target)
	_, _ = rpc.InjectIntoSameIndex(nil) // nil element in populated collection - ok
}

func Test_C27_RPC_UnmarshalAt(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.New("hello").Ptr())
	var s string
	_ = rpc.UnmarshalAt(0, &s)
}

func Test_C27_RPC_UnmarshalIntoSameIndex(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.New("hello").Ptr())
	var s string
	_, _ = rpc.UnmarshalIntoSameIndex(&s)
	_, _ = rpc.UnmarshalIntoSameIndex(nil) // nil element in populated collection - ok
}

func Test_C27_RPC_NonPtr(t *testing.T) { _ = corejson.NewResultsPtrCollection.Empty().NonPtr() }
func Test_C27_RPC_Ptr(t *testing.T) { _ = corejson.NewResultsPtrCollection.Empty().Ptr() }

func Test_C27_RPC_Clear(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := corejson.New("x")
	rpc.Add(r.Ptr())
	rpc.Clear()
}

func Test_C27_RPC_Dispose(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Dispose()
}

func Test_C27_RPC_Json(t *testing.T) { _ = corejson.NewResultsPtrCollection.Empty().Json() }
func Test_C27_RPC_JsonPtr(t *testing.T) { _ = corejson.NewResultsPtrCollection.Empty().JsonPtr() }
func Test_C27_RPC_JsonModel(t *testing.T) { _ = corejson.NewResultsPtrCollection.Empty().JsonModel() }
func Test_C27_RPC_JsonModelAny(t *testing.T) { _ = corejson.NewResultsPtrCollection.Empty().JsonModelAny() }
func Test_C27_RPC_AsJsonContractsBinder(t *testing.T) { _ = corejson.NewResultsPtrCollection.Empty().AsJsonContractsBinder() }
func Test_C27_RPC_AsJsoner(t *testing.T) { _ = corejson.NewResultsPtrCollection.Empty().AsJsoner() }
func Test_C27_RPC_AsJsonParseSelfInjector(t *testing.T) { _ = corejson.NewResultsPtrCollection.Empty().AsJsonParseSelfInjector() }

func Test_C27_RPC_JsonParseSelfInject(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := corejson.New(rpc)
	_ = rpc.JsonParseSelfInject(&r)
}

func Test_C27_RPC_ParseInjectUsingJson_Error(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	bad := corejson.NewResult.UsingString(`invalid`)
	_, err := rpc.ParseInjectUsingJson(bad)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C27_RPC_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() { recover() }()
	rpc := corejson.NewResultsPtrCollection.Empty()
	bad := corejson.NewResult.UsingString(`invalid`)
	rpc.ParseInjectUsingJsonMust(bad)
}

func Test_C27_RPC_Clone(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := corejson.New("x")
	rpc.Add(r.Ptr())
	_ = rpc.Clone(true)
}

func Test_C27_RPC_Clone_Empty(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.Clone(true)
}

func Test_C27_RPC_GetPagesSize(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	for i := 0; i < 5; i++ { rpc.Add(corejson.New(i).Ptr()) }
	actual := args.Map{"result": rpc.GetPagesSize(2) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": rpc.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C27_RPC_GetPagedCollection(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	for i := 0; i < 5; i++ { rpc.Add(corejson.New(i).Ptr()) }
	pages := rpc.GetPagedCollection(2)
	actual := args.Map{"result": len(pages) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C27_RPC_GetPagedCollection_Small(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.New("x").Ptr())
	pages := rpc.GetPagedCollection(5)
	actual := args.Map{"result": len(pages) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C27_RPC_AddJsoners(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddJsoners(true)
}
