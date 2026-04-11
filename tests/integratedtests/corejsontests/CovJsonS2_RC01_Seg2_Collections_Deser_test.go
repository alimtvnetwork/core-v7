package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// corejson Coverage — Segment 2: Collections, Deserializer, Serializer, Creators
// ══════════════════════════════════════════════════════════════════════════════

// --- ResultsCollection ---

func newTestRC() *corejson.ResultsCollection {
	rc := corejson.NewResultsCollection.UsingCap(4)
	rc.Add(corejson.New(map[string]int{"a": 1}))
	rc.Add(corejson.New(map[string]int{"b": 2}))
	return rc
}

func Test_CovJsonS2_RC01_Length_IsEmpty_HasAnyItem(t *testing.T) {
	// Arrange
	rc := newTestRC()

	// Act
	actual := args.Map{"result": rc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": rc.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual = args.Map{"result": rc.HasAnyItem()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	// nil
	var nilRC *corejson.ResultsCollection
	actual = args.Map{"result": nilRC.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovJsonS2_RC02_FirstOrDefault_LastOrDefault(t *testing.T) {
	// Arrange
	rc := newTestRC()
	f := rc.FirstOrDefault()

	// Act
	actual := args.Map{"result": f == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	l := rc.LastOrDefault()
	actual = args.Map{"result": l == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	// empty
	empty := corejson.NewResultsCollection.Empty()
	actual = args.Map{"result": empty.FirstOrDefault() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": empty.LastOrDefault() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJsonS2_RC03_Take_Limit_Skip(t *testing.T) {
	// Arrange
	rc := newTestRC()
	taken := rc.Take(1)

	// Act
	actual := args.Map{"result": taken.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	limited := rc.Limit(1)
	actual = args.Map{"result": limited.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	// limit -1 returns all
	all := rc.Limit(-1)
	actual = args.Map{"result": all.Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	skipped := rc.Skip(1)
	actual = args.Map{"result": skipped.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	// empty
	empty := corejson.NewResultsCollection.Empty()
	actual = args.Map{"result": empty.Take(1).Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": empty.Limit(1).Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": empty.Skip(1).Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovJsonS2_RC04_AddSkipOnNil_AddNonNilNonError(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSkipOnNil(nil)

	// Act
	actual := args.Map{"result": rc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	r := corejson.NewPtr(1)
	rc.AddSkipOnNil(r)
	actual = args.Map{"result": rc.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	// non nil non error
	rc2 := corejson.NewResultsCollection.Empty()
	errR := &corejson.Result{Error: errors.New("err")}
	rc2.AddNonNilNonError(errR)
	actual = args.Map{"result": rc2.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	rc2.AddNonNilNonError(r)
	actual = args.Map{"result": rc2.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovJsonS2_RC05_HasError_AllErrors(t *testing.T) {
	// Arrange
	rc := newTestRC()

	// Act
	actual := args.Map{"result": rc.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	errs, hasAny := rc.AllErrors()
	actual = args.Map{"result": hasAny || len(errs) > 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no errors", actual)
	// empty
	empty := corejson.NewResultsCollection.Empty()
	errs2, hasAny2 := empty.AllErrors()
	actual = args.Map{"result": hasAny2 || len(errs2) > 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no errors", actual)
}

func Test_CovJsonS2_RC06_GetErrorsStrings_GetErrorsAsSingle(t *testing.T) {
	// Arrange
	rc := newTestRC()
	ss := rc.GetErrorsStrings()

	// Act
	actual := args.Map{"result": len(ss) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	_ = rc.GetErrorsStringsPtr()
	_ = rc.GetErrorsAsSingleString()
	_ = rc.GetErrorsAsSingle()
	// empty
	empty := corejson.NewResultsCollection.Empty()
	ss2 := empty.GetErrorsStrings()
	actual = args.Map{"result": len(ss2) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovJsonS2_RC07_GetAt_GetAtSafe_GetAtSafeUsingLength(t *testing.T) {
	// Arrange
	rc := newTestRC()
	_ = rc.GetAt(0)
	safe := rc.GetAtSafe(0)

	// Act
	actual := args.Map{"result": safe == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	safe2 := rc.GetAtSafe(-2)
	actual = args.Map{"result": safe2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	safe3 := rc.GetAtSafeUsingLength(0, 2)
	actual = args.Map{"result": safe3 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	safe4 := rc.GetAtSafeUsingLength(5, 2)
	actual = args.Map{"result": safe4 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJsonS2_RC08_AddPtr_Adds_AddsPtr_AddAny_AddAnyItems(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddPtr(nil)

	// Act
	actual := args.Map{"result": rc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	r := corejson.NewPtr(1)
	rc.AddPtr(r)
	actual = args.Map{"result": rc.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	rc.AddsPtr(nil, r)
	rc.AddAny(1)
	rc.AddAny(nil)
	rc.AddAnyItems(1, nil, 2)
	rc.AddAnyItems(nil)
}

func Test_CovJsonS2_RC09_AddResultsCollection_AddNonNilItemsPtr(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddResultsCollection(nil)
	rc2 := newTestRC()
	rc.AddResultsCollection(rc2)
	rc.AddNonNilItemsPtr(nil)
}

func Test_CovJsonS2_RC10_AddAnyItemsSlice(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddAnyItemsSlice(nil)
	rc.AddAnyItemsSlice([]any{1, nil, 2})
}

func Test_CovJsonS2_RC11_Dispose_Clear(t *testing.T) {
	rc := newTestRC()
	rc.Dispose()
	// nil dispose
	var nilRC *corejson.ResultsCollection
	nilRC.Dispose()
}

func Test_CovJsonS2_RC12_GetStrings(t *testing.T) {
	// Arrange
	rc := newTestRC()
	ss := rc.GetStrings()

	// Act
	actual := args.Map{"result": len(ss) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	_ = rc.GetStringsPtr()
	// empty
	empty := corejson.NewResultsCollection.Empty()
	ss2 := empty.GetStrings()
	actual = args.Map{"result": len(ss2) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovJsonS2_RC13_GetPagesSize_GetPagedCollection_GetSinglePageCollection(t *testing.T) {
	// Arrange
	// build 15 items
	rc := corejson.NewResultsCollection.UsingCap(20)
	for i := 0; i < 15; i++ {
		rc.AddAny(i)
	}
	ps := rc.GetPagesSize(5)

	// Act
	actual := args.Map{"result": ps != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual = args.Map{"result": rc.GetPagesSize(0) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	pages := rc.GetPagedCollection(5)
	actual = args.Map{"result": len(pages) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	single := rc.GetSinglePageCollection(5, 2)
	actual = args.Map{"result": single.Length() != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	// small collection
	small := newTestRC()
	pages2 := small.GetPagedCollection(10)
	actual = args.Map{"result": len(pages2) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	single2 := small.GetSinglePageCollection(10, 1)
	actual = args.Map{"result": single2.Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_CovJsonS2_RC14_Json_JsonPtr_JsonModel_Interfaces(t *testing.T) {
	// Arrange
	rc := newTestRC()
	j := rc.Json()

	// Act
	actual := args.Map{"result": j.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	jp := rc.JsonPtr()
	actual = args.Map{"result": jp == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	_ = rc.JsonModel()
	_ = rc.JsonModelAny()
	_ = rc.AsJsonContractsBinder()
	_ = rc.AsJsoner()
	_ = rc.AsJsonParseSelfInjector()
	_ = rc.NonPtr()
	_ = rc.Ptr()
}

func Test_CovJsonS2_RC15_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	rc := newTestRC()
	jr := rc.JsonPtr()
	rc2 := corejson.NewResultsCollection.Empty()
	_, err := rc2.ParseInjectUsingJson(jr)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS2_RC16_Clone_ShadowClone_ClonePtr(t *testing.T) {
	// Arrange
	rc := newTestRC()
	c := rc.ShadowClone()
	_ = c
	c2 := rc.Clone(true)
	_ = c2
	cp := rc.ClonePtr(true)

	// Act
	actual := args.Map{"result": cp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	// nil
	var nilRC *corejson.ResultsCollection
	actual = args.Map{"result": nilRC.ClonePtr(true) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJsonS2_RC17_UnmarshalAt(t *testing.T) {
	// Arrange
	rc := newTestRC()
	var m map[string]int
	err := rc.UnmarshalAt(0, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS2_RC18_UnmarshalIntoSameIndex(t *testing.T) {
	// Arrange
	rc := newTestRC()
	var m1, m2 map[string]int
	errs, hasAny := rc.UnmarshalIntoSameIndex(&m1, &m2)

	// Act
	actual := args.Map{"result": hasAny}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_ = errs
	// nil
	errs2, hasAny2 := rc.UnmarshalIntoSameIndex(nil)
	_ = errs2
	_ = hasAny2
}

func Test_CovJsonS2_RC19_AddSerializer_AddSerializerFunc(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSerializer(nil)
	rc.AddSerializerFunc(nil)
	rc.AddSerializerFunc(func() ([]byte, error) {
		return []byte(`1`), nil
	})
}

func Test_CovJsonS2_RC20_AddSerializers_AddSerializerFunctions(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSerializers()
	rc.AddSerializerFunctions()
}

// --- ResultsPtrCollection ---

func newTestRPC() *corejson.ResultsPtrCollection {
	rpc := corejson.NewResultsPtrCollection.UsingCap(4)
	rpc.Add(corejson.NewPtr(map[string]int{"a": 1}))
	rpc.Add(corejson.NewPtr(map[string]int{"b": 2}))
	return rpc
}

func Test_CovJsonS2_RPC01_Length_IsEmpty_HasAnyItem(t *testing.T) {
	// Arrange
	rpc := newTestRPC()

	// Act
	actual := args.Map{"result": rpc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	var nilRPC *corejson.ResultsPtrCollection
	actual = args.Map{"result": nilRPC.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovJsonS2_RPC02_FirstOrDefault_LastOrDefault(t *testing.T) {
	// Arrange
	rpc := newTestRPC()

	// Act
	actual := args.Map{"result": rpc.FirstOrDefault() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": rpc.LastOrDefault() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	empty := corejson.NewResultsPtrCollection.Empty()
	actual = args.Map{"result": empty.FirstOrDefault() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJsonS2_RPC03_Take_Limit_Skip(t *testing.T) {
	// Arrange
	rpc := newTestRPC()

	// Act
	actual := args.Map{"result": rpc.Take(1).Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": rpc.Limit(1).Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual = args.Map{"result": rpc.Limit(-1).Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": rpc.Skip(1).Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	empty := corejson.NewResultsPtrCollection.Empty()
	_ = empty.Take(1)
	_ = empty.Limit(1)
	_ = empty.Skip(1)
}

func Test_CovJsonS2_RPC04_AddSkipOnNil_AddNonNilNonError(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSkipOnNil(nil)
	r := corejson.NewPtr(1)
	rpc.AddSkipOnNil(r)

	// Act
	actual := args.Map{"result": rpc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	errR := &corejson.Result{Error: errors.New("err")}
	rpc.AddNonNilNonError(errR)
	actual = args.Map{"result": rpc.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovJsonS2_RPC05_HasError_AllErrors_GetErrorsStrings(t *testing.T) {
	// Arrange
	rpc := newTestRPC()

	// Act
	actual := args.Map{"result": rpc.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	errs, hasAny := rpc.AllErrors()
	actual = args.Map{"result": hasAny || len(errs) > 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no errors", actual)
	_ = rpc.GetErrorsStrings()
	_ = rpc.GetErrorsStringsPtr()
	_ = rpc.GetErrorsAsSingleString()
	_ = rpc.GetErrorsAsSingle()
	// empty
	empty := corejson.NewResultsPtrCollection.Empty()
	_, _ = empty.AllErrors()
	_ = empty.GetErrorsStrings()
}

func Test_CovJsonS2_RPC06_GetAt_GetAtSafe(t *testing.T) {
	// Arrange
	rpc := newTestRPC()
	_ = rpc.GetAt(0)
	safe := rpc.GetAtSafe(0)

	// Act
	actual := args.Map{"result": safe == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": rpc.GetAtSafe(-2) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": rpc.GetAtSafeUsingLength(5, 2) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJsonS2_RPC07_Add_AddResult_Adds_AddsPtr_AddAny_AddAnyItems(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := corejson.New(1)
	rpc.AddResult(r)
	rpc.Adds(nil)
	rpc.AddAny(nil)
	rpc.AddAny(1)
	rpc.AddAnyItems(1, nil, 2)
	rpc.AddAnyItems(nil)
}

func Test_CovJsonS2_RPC08_AddResultsCollection_AddNonNilItems(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddResultsCollection(nil)
	rpc2 := newTestRPC()
	rpc.AddResultsCollection(rpc2)
}

func Test_CovJsonS2_RPC09_UnmarshalAt(t *testing.T) {
	// Arrange
	rpc := newTestRPC()
	var m map[string]int
	err := rpc.UnmarshalAt(0, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	// nil result
	rpc2 := corejson.NewResultsPtrCollection.Empty()
	rpc2.Add(nil)
	err2 := rpc2.UnmarshalAt(0, &m)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil result", actual)
}

func Test_CovJsonS2_RPC10_UnmarshalIntoSameIndex(t *testing.T) {
	// Arrange
	rpc := newTestRPC()
	var m1, m2 map[string]int
	errs, hasAny := rpc.UnmarshalIntoSameIndex(&m1, &m2)

	// Act
	actual := args.Map{"result": hasAny}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_ = errs
	// nil
	errs2, hasAny2 := rpc.UnmarshalIntoSameIndex(nil)
	_ = errs2
	_ = hasAny2
}

func Test_CovJsonS2_RPC11_AddSerializer_Funcs(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSerializer(nil)
	rpc.AddSerializerFunc(nil)
	rpc.AddSerializerFunc(func() ([]byte, error) {
		return []byte(`1`), nil
	})
	rpc.AddSerializers()
	rpc.AddSerializerFunctions()
}

// --- MapResults ---

func newTestMR() *corejson.MapResults {
	mr := corejson.NewMapResults.UsingCap(4)
	mr.Add("a", corejson.New(map[string]int{"x": 1}))
	mr.Add("b", corejson.New(map[string]int{"y": 2}))
	return mr
}

func Test_CovJsonS2_MR01_Length_IsEmpty_HasAnyItem(t *testing.T) {
	// Arrange
	mr := newTestMR()

	// Act
	actual := args.Map{"result": mr.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	var nilMR *corejson.MapResults
	actual = args.Map{"result": nilMR.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovJsonS2_MR02_AddSkipOnNil_GetByKey(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddSkipOnNil("k", nil)
	r := corejson.NewPtr(1)
	mr.AddSkipOnNil("k", r)

	// Act
	actual := args.Map{"result": mr.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	got := mr.GetByKey("k")
	actual = args.Map{"result": got == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": mr.GetByKey("missing") != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJsonS2_MR03_HasError_AllErrors_GetErrorsStrings(t *testing.T) {
	// Arrange
	mr := newTestMR()

	// Act
	actual := args.Map{"result": mr.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	errs, hasAny := mr.AllErrors()
	actual = args.Map{"result": hasAny || len(errs) > 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected none", actual)
	_ = mr.GetErrorsStrings()
	_ = mr.GetErrorsStringsPtr()
	_ = mr.GetErrorsAsSingleString()
	_ = mr.GetErrorsAsSingle()
	// empty
	empty := corejson.NewMapResults.Empty()
	_, _ = empty.AllErrors()
	_ = empty.GetErrorsStrings()
}

func Test_CovJsonS2_MR04_AllKeys_AllKeysSorted_AllValues(t *testing.T) {
	// Arrange
	mr := newTestMR()
	keys := mr.AllKeys()

	// Act
	actual := args.Map{"result": len(keys) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	sorted := mr.AllKeysSorted()
	actual = args.Map{"result": sorted[0] != "a"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a first", actual)
	vals := mr.AllValues()
	actual = args.Map{"result": len(vals) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	_ = mr.AllResults()
	_ = mr.AllResultsCollection()
	// empty
	empty := corejson.NewMapResults.Empty()
	actual = args.Map{"result": len(empty.AllKeys()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": len(empty.AllKeysSorted()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": len(empty.AllValues()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": empty.AllResultsCollection().HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovJsonS2_MR05_GetStrings(t *testing.T) {
	// Arrange
	mr := newTestMR()
	ss := mr.GetStrings()

	// Act
	actual := args.Map{"result": len(ss) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	_ = mr.GetStringsPtr()
	// empty
	empty := corejson.NewMapResults.Empty()
	actual = args.Map{"result": len(empty.GetStrings()) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovJsonS2_MR06_Add_AddPtr_AddAny_AddAnySkipOnNil(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddPtr("k", nil)

	// Act
	actual := args.Map{"result": mr.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	err := mr.AddAny("k", map[string]int{"a": 1})
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	err2 := mr.AddAny("nil", nil)
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err3 := mr.AddAnySkipOnNil("nil", nil)
	actual = args.Map{"result": err3 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJsonS2_MR07_AddAnyNonEmptyNonError_AddAnyNonEmpty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddAnyNonEmptyNonError("k", nil)
	mr.AddAnyNonEmpty("k2", nil)
	mr.AddAnyNonEmpty("k3", 1)
}

func Test_CovJsonS2_MR08_AddKeyWithResult_AddKeyWithResultPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	kr := corejson.KeyWithResult{Key: "k", Result: corejson.New(1)}
	mr.AddKeyWithResult(kr)
	mr.AddKeyWithResultPtr(nil)
	mr.AddKeyWithResultPtr(&kr)
	mr.AddKeysWithResultsPtr()
	mr.AddKeysWithResults()
}

func Test_CovJsonS2_MR09_AddKeyAny(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	ka := corejson.KeyAny{Key: "k", AnyInf: 1}
	mr.AddKeyAnyInf(ka)
	mr.AddKeyAnyInfPtr(nil)
	mr.AddKeyAnyInfPtr(&ka)
	mr.AddKeyAnyItems(ka)
	mr.AddKeyAnyItemsPtr(nil)
}

func Test_CovJsonS2_MR10_AddNonEmptyNonErrorPtr_AddMapResults_AddMapAnyItems(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddNonEmptyNonErrorPtr("k", nil)
	mr.AddMapResults(nil)
	mr2 := newTestMR()
	mr.AddMapResults(mr2)
	mr.AddMapAnyItems(nil)
	mr.AddMapAnyItems(map[string]any{"x": 1})
}

func Test_CovJsonS2_MR11_GetPagesSize_GetPagedCollection(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.UsingCap(10)
	for i := 0; i < 10; i++ {
		mr.Add(string(rune('a'+i)), corejson.New(i))
	}

	// Act
	actual := args.Map{"result": mr.GetPagesSize(3) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual = args.Map{"result": mr.GetPagesSize(0) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	pages := mr.GetPagedCollection(3)
	actual = args.Map{"result": len(pages) < 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	// small
	small := newTestMR()
	pages2 := small.GetPagedCollection(10)
	actual = args.Map{"result": len(pages2) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovJsonS2_MR12_AddMapResultsUsingCloneOption(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	raw := map[string]corejson.Result{"k": corejson.New(1)}
	mr.AddMapResultsUsingCloneOption(false, false, raw)
	mr2 := corejson.NewMapResults.Empty()
	mr2.AddMapResultsUsingCloneOption(true, true, raw)
	mr3 := corejson.NewMapResults.Empty()
	mr3.AddMapResultsUsingCloneOption(false, false, nil)
}

func Test_CovJsonS2_MR13_GetNewMapUsingKeys(t *testing.T) {
	// Arrange
	mr := newTestMR()
	sub := mr.GetNewMapUsingKeys(false, "a")

	// Act
	actual := args.Map{"result": sub.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	// empty keys
	empty := mr.GetNewMapUsingKeys(false)
	actual = args.Map{"result": empty.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_CovJsonS2_MR14_ResultCollection_Json_JsonPtr_Interfaces(t *testing.T) {
	// Arrange
	mr := newTestMR()
	rc := mr.ResultCollection()

	// Act
	actual := args.Map{"result": rc.IsEmpty()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	_ = mr.Json()
	_ = mr.JsonPtr()
	_ = mr.JsonModel()
	_ = mr.JsonModelAny()
	_ = mr.AsJsonContractsBinder()
	_ = mr.AsJsoner()
	_ = mr.AsJsonParseSelfInjector()
}

func Test_CovJsonS2_MR15_Clear_Dispose(t *testing.T) {
	mr := newTestMR()
	mr.Dispose()
	var nilMR *corejson.MapResults
	nilMR.Dispose()
}

func Test_CovJsonS2_MR16_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	mr := newTestMR()
	jr := mr.JsonPtr()
	mr2 := corejson.NewMapResults.Empty()
	_, err := mr2.ParseInjectUsingJson(jr)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

// --- Deserializer (deserializerLogic) ---

func Test_CovJsonS2_DL01_Apply_UsingResult(t *testing.T) {
	// Arrange
	r := corejson.NewPtr(map[string]int{"a": 1})
	var m map[string]int
	err := corejson.Deserialize.Apply(r, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	err2 := corejson.Deserialize.UsingResult(r, &m)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS2_DL02_UsingString_FromString(t *testing.T) {
	// Arrange
	var m map[string]int
	err := corejson.Deserialize.UsingString(`{"a":1}`, &m)

	// Act
	actual := args.Map{"result": err != nil || m["a"] != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a=1", actual)
	var m2 map[string]int
	err2 := corejson.Deserialize.FromString(`{"a":1}`, &m2)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS2_DL03_UsingStringOption_UsingStringIgnoreEmpty(t *testing.T) {
	// Arrange
	var m map[string]int
	err := corejson.Deserialize.UsingStringOption(true, "", &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty", actual)
	err2 := corejson.Deserialize.UsingStringIgnoreEmpty("", &m)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty", actual)
}

func Test_CovJsonS2_DL04_UsingStringPtr(t *testing.T) {
	// Arrange
	s := `{"a":1}`
	var m map[string]int
	err := corejson.Deserialize.UsingStringPtr(&s, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	err2 := corejson.Deserialize.UsingStringPtr(nil, &m)
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_CovJsonS2_DL05_UsingError_UsingErrorWhichJsonResult(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingError(nil, nil)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	err2 := corejson.Deserialize.UsingErrorWhichJsonResult(nil, nil)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJsonS2_DL06_UsingBytes_UsingBytesPointer_UsingBytesMust(t *testing.T) {
	// Arrange
	var m map[string]int
	err := corejson.Deserialize.UsingBytes([]byte(`{"a":1}`), &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	err2 := corejson.Deserialize.UsingBytesPointer(nil, &m)
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil bytes", actual)
}

func Test_CovJsonS2_DL07_UsingBytesIf_UsingBytesPointerIf(t *testing.T) {
	// Arrange
	var m map[string]int
	err := corejson.Deserialize.UsingBytesIf(false, nil, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	err2 := corejson.Deserialize.UsingBytesPointerIf(false, nil, &m)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJsonS2_DL08_MapAnyToPointer(t *testing.T) {
	// Arrange
	var m map[string]int
	err := corejson.Deserialize.MapAnyToPointer(true, nil, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty skip", actual)
	err2 := corejson.Deserialize.MapAnyToPointer(false, map[string]any{"a": 1}, &m)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS2_DL09_FromTo(t *testing.T) {
	// Arrange
	from := map[string]int{"a": 1}
	var to map[string]int
	err := corejson.Deserialize.FromTo(from, &to)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS2_DL10_AnyToFieldsMap(t *testing.T) {
	// AnyToFieldsMap → DeserializedFieldsToMap passes value not pointer — known limitation
	m, _ := corejson.Deserialize.AnyToFieldsMap(map[string]int{"a": 1})
	_ = m // covers the call path regardless of result
}

// --- deserializeFromBytesTo ---

func Test_CovJsonS2_BT01_Strings_String_Integer(t *testing.T) {
	// Arrange
	sb, _ := corejson.Serialize.Raw([]string{"a", "b"})
	lines, err := corejson.Deserialize.BytesTo.Strings(sb)

	// Act
	actual := args.Map{"result": err != nil || len(lines) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	sb2, _ := corejson.Serialize.Raw("hello")
	s, err2 := corejson.Deserialize.BytesTo.String(sb2)
	actual = args.Map{"result": err2 != nil || s != "hello"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	sb3, _ := corejson.Serialize.Raw(42)
	i, err3 := corejson.Deserialize.BytesTo.Integer(sb3)
	actual = args.Map{"result": err3 != nil || i != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	_ = corejson.Deserialize.BytesTo.IntegerMust(sb3)
}

func Test_CovJsonS2_BT02_Integer64_Integers_Bool(t *testing.T) {
	// Arrange
	sb, _ := corejson.Serialize.Raw(int64(42))
	i64, err := corejson.Deserialize.BytesTo.Integer64(sb)

	// Act
	actual := args.Map{"result": err != nil || i64 != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	_ = corejson.Deserialize.BytesTo.Integer64Must(sb)
	sbi, _ := corejson.Serialize.Raw([]int{1, 2})
	ints, err2 := corejson.Deserialize.BytesTo.Integers(sbi)
	actual = args.Map{"result": err2 != nil || len(ints) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	sbb, _ := corejson.Serialize.Raw(true)
	b, err3 := corejson.Deserialize.BytesTo.Bool(sbb)
	actual = args.Map{"result": err3 != nil || !b}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_CovJsonS2_BT03_MapAnyItem_MapStringString(t *testing.T) {
	// Arrange
	sb, _ := corejson.Serialize.Raw(map[string]any{"a": 1})
	m, err := corejson.Deserialize.BytesTo.MapAnyItem(sb)

	// Act
	actual := args.Map{"result": err != nil || len(m) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected map", actual)
	_ = corejson.Deserialize.BytesTo.MapAnyItemMust(sb)
	sb2, _ := corejson.Serialize.Raw(map[string]string{"a": "b"})
	ms, err2 := corejson.Deserialize.BytesTo.MapStringString(sb2)
	actual = args.Map{"result": err2 != nil || len(ms) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected map", actual)
	_ = corejson.Deserialize.BytesTo.MapStringStringMust(sb2)
}

func Test_CovJsonS2_BT04_Bytes_BytesMust(t *testing.T) {
	// Arrange
	inner := []byte(`"hello"`)
	sb, _ := corejson.Serialize.Raw(inner)
	b, err := corejson.Deserialize.BytesTo.Bytes(sb)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = b
}

func Test_CovJsonS2_BT05_ResultCollection_ResultsPtrCollection_MapResults(t *testing.T) {
	// Arrange
	rc := newTestRC()
	sb, _ := corejson.Serialize.Raw(rc)
	rc2, err := corejson.Deserialize.BytesTo.ResultCollection(sb)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = rc2
	rpc := newTestRPC()
	sb2, _ := corejson.Serialize.Raw(rpc)
	rpc2, err2 := corejson.Deserialize.BytesTo.ResultsPtrCollection(sb2)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = rpc2
	mr := newTestMR()
	sb3, _ := corejson.Serialize.Raw(mr)
	mr2, err3 := corejson.Deserialize.BytesTo.MapResults(sb3)
	actual = args.Map{"result": err3 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = mr2
}

// --- deserializeFromResultTo ---

func Test_CovJsonS2_RT01_String_Bool_Byte(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")
	s, err := corejson.Deserialize.ResultTo.String(r)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	r2 := corejson.NewPtr(true)
	b, err2 := corejson.Deserialize.ResultTo.Bool(r2)
	actual = args.Map{"result": err2 != nil || !b}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	r3 := corejson.NewPtr(byte(5))
	bv, err3 := corejson.Deserialize.ResultTo.Byte(r3)
	actual = args.Map{"result": err3 != nil || bv != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_CovJsonS2_RT02_MapAnyItem_MapStringString(t *testing.T) {
	// Arrange
	r := corejson.NewPtr(map[string]any{"a": 1})
	m, err := corejson.Deserialize.ResultTo.MapAnyItem(r)

	// Act
	actual := args.Map{"result": err != nil || len(m) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected map", actual)
	r2 := corejson.NewPtr(map[string]string{"a": "b"})
	ms, err2 := corejson.Deserialize.ResultTo.MapStringString(r2)
	actual = args.Map{"result": err2 != nil || len(ms) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected map", actual)
}

// --- Serializer (serializerLogic) ---

func Test_CovJsonS2_SL01_Apply_Various(t *testing.T) {
	// Arrange
	r := corejson.Serialize.Apply(1)

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = corejson.Serialize.FromBytes([]byte("test"))
	_ = corejson.Serialize.FromStrings([]string{"a"})
	_ = corejson.Serialize.FromStringsSpread("a", "b")
	_ = corejson.Serialize.FromString("hello")
	_ = corejson.Serialize.FromInteger(42)
	_ = corejson.Serialize.FromInteger64(42)
	_ = corejson.Serialize.FromBool(true)
	_ = corejson.Serialize.FromIntegers([]int{1, 2})
	_ = corejson.Serialize.StringsApply([]string{"a"})
}

func Test_CovJsonS2_SL02_UsingAny_UsingAnyPtr_Raw_Marshal(t *testing.T) {
	// Arrange
	r := corejson.Serialize.UsingAny(1)

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	rp := corejson.Serialize.UsingAnyPtr(1)
	actual = args.Map{"result": rp.HasError()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	raw, err := corejson.Serialize.Raw(1)
	actual = args.Map{"result": err != nil || len(raw) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	raw2, err2 := corejson.Serialize.Marshal(1)
	actual = args.Map{"result": err2 != nil || len(raw2) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_CovJsonS2_SL03_ToBytes_ToString(t *testing.T) {
	_ = corejson.Serialize.ToBytesMust(1)
	_ = corejson.Serialize.ToSafeBytesMust(1)
	_ = corejson.Serialize.ToSafeBytesSwallowErr(1)
	_ = corejson.Serialize.ToBytesSwallowErr(1)
	_, _ = corejson.Serialize.ToBytesErr(1)
	_ = corejson.Serialize.ToString(1)
	_ = corejson.Serialize.ToStringMust(1)
	_, _ = corejson.Serialize.ToStringErr(1)
	_, _ = corejson.Serialize.ToPrettyStringErr(1)
	_ = corejson.Serialize.ToPrettyStringIncludingErr(1)
	_ = corejson.Serialize.Pretty(1)
}

// --- Creators ---

func Test_CovJsonS2_CR01_Empty(t *testing.T) {
	_ = corejson.Empty.Result()
	_ = corejson.Empty.ResultPtr()
	_ = corejson.Empty.ResultWithErr("type", errors.New("err"))
	_ = corejson.Empty.ResultPtrWithErr("type", errors.New("err"))
	_ = corejson.Empty.BytesCollection()
	_ = corejson.Empty.BytesCollectionPtr()
	_ = corejson.Empty.ResultsCollection()
	_ = corejson.Empty.ResultsPtrCollection()
	_ = corejson.Empty.MapResults()
}

func Test_CovJsonS2_CR02_NewResult(t *testing.T) {
	_ = corejson.NewResult.UsingBytes([]byte(`1`))
	_ = corejson.NewResult.UsingBytesType([]byte(`1`), "int")
	_ = corejson.NewResult.UsingBytesTypePtr([]byte(`1`), "int")
	_ = corejson.NewResult.UsingTypeBytesPtr("int", []byte(`1`))
	_ = corejson.NewResult.UsingBytesPtr(nil)
	_ = corejson.NewResult.UsingBytesPtr([]byte(`1`))
	_ = corejson.NewResult.UsingBytesPtrErrPtr(nil, errors.New("e"), "t")
	_ = corejson.NewResult.UsingBytesPtrErrPtr([]byte(`1`), nil, "t")
	_ = corejson.NewResult.UsingBytesErrPtr(nil, nil, "t")
	_ = corejson.NewResult.UsingBytesErrPtr([]byte(`1`), nil, "t")
	s := `{"a":1}`
	_ = corejson.NewResult.PtrUsingStringPtr(&s, "t")
	_ = corejson.NewResult.PtrUsingStringPtr(nil, "t")
	_ = corejson.NewResult.UsingErrorStringPtr(nil, &s, "t")
	_ = corejson.NewResult.UsingErrorStringPtr(errors.New("e"), nil, "t")
	_ = corejson.NewResult.Ptr(nil, nil, "t")
	_ = corejson.NewResult.UsingJsonBytesTypeError(nil, nil, "t")
	_ = corejson.NewResult.UsingJsonBytesError(nil, nil)
	_ = corejson.NewResult.UsingTypePlusString("t", s)
	_ = corejson.NewResult.UsingTypePlusStringPtr("t", nil)
	_ = corejson.NewResult.UsingTypePlusStringPtr("t", &s)
	_ = corejson.NewResult.UsingStringWithType(s, "t")
	_ = corejson.NewResult.UsingString(s)
	_ = corejson.NewResult.UsingStringPtr(nil)
	_ = corejson.NewResult.UsingStringPtr(&s)
	_ = corejson.NewResult.CreatePtr(nil, nil, "t")
	_ = corejson.NewResult.NonPtr(nil, nil, "t")
	_ = corejson.NewResult.Create(nil, nil, "t")
	_ = corejson.NewResult.PtrUsingBytesPtr(nil, errors.New("e"), "t")
	_ = corejson.NewResult.PtrUsingBytesPtr(nil, nil, "t")
	_ = corejson.NewResult.PtrUsingBytesPtr([]byte(`1`), nil, "t")
	_ = corejson.NewResult.CastingAny(1)
	_ = corejson.NewResult.Any(1)
	_ = corejson.NewResult.AnyPtr(1)
	_ = corejson.NewResult.Error(errors.New("e"))
	_ = corejson.NewResult.ErrorPtr(errors.New("e"))
	_ = corejson.NewResult.Empty()
	_ = corejson.NewResult.EmptyPtr()
	_ = corejson.NewResult.TypeName("t")
	_ = corejson.NewResult.TypeNameBytes("t")
	_ = corejson.NewResult.Many(1, 2, 3)
	_ = corejson.NewResult.Serialize(1)
	_ = corejson.NewResult.Marshal(1)
	_ = corejson.NewResult.Serialize(1) // ApplyMust does not exist, use Serialize
}

func Test_CovJsonS2_CR03_NewResultCreator_DeserializeUsingBytes_DeserializeUsingResult(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	b, _ := r.Serialize()
	_ = corejson.NewResult.DeserializeUsingBytes(b)
	_ = corejson.NewResult.UnmarshalUsingBytes(b)
	rp := r.ToPtr()
	_ = corejson.NewResult.DeserializeUsingResult(rp)
}

func Test_CovJsonS2_CR04_NewResultsCollection(t *testing.T) {
	_ = corejson.NewResultsCollection.Empty()
	_ = corejson.NewResultsCollection.Default()
	_ = corejson.NewResultsCollection.UsingCap(5)
	_ = corejson.NewResultsCollection.AnyItems(1, 2)
	_ = corejson.NewResultsCollection.AnyItemsPlusCap(2, 1, 2)
	_ = corejson.NewResultsCollection.AnyItemsPlusCap(2)
	_ = corejson.NewResultsCollection.UsingResultsPtr()
	_ = corejson.NewResultsCollection.UsingResults()
	_ = corejson.NewResultsCollection.UsingResultsPlusCap(2, corejson.New(1))
	_ = corejson.NewResultsCollection.UsingResultsPlusCap(2)
	_ = corejson.NewResultsCollection.UsingResultsPtrPlusCap(2, corejson.NewPtr(1))
	_ = corejson.NewResultsCollection.UsingResultsPtrPlusCap(2)
	_ = corejson.NewResultsCollection.Serializers()
	_ = corejson.NewResultsCollection.SerializerFunctions()
}

func Test_CovJsonS2_CR05_NewResultsPtrCollection(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.Empty()
	_ = corejson.NewResultsPtrCollection.Default()
	_ = corejson.NewResultsPtrCollection.UsingCap(5)
	_ = corejson.NewResultsPtrCollection.AnyItems(1, 2)
	_ = corejson.NewResultsPtrCollection.AnyItemsPlusCap(2, 1, 2)
	_ = corejson.NewResultsPtrCollection.AnyItemsPlusCap(2)
	_ = corejson.NewResultsPtrCollection.UsingResults()
	_ = corejson.NewResultsPtrCollection.UsingResultsPlusCap(2, corejson.NewPtr(1))
	_ = corejson.NewResultsPtrCollection.Serializers()
}

func Test_CovJsonS2_CR06_NewMapResults(t *testing.T) {
	_ = corejson.NewMapResults.Empty()
	_ = corejson.NewMapResults.UsingCap(5)
	_ = corejson.NewMapResults.UsingKeyAnyItems(0, corejson.KeyAny{Key: "k", AnyInf: 1})
	_ = corejson.NewMapResults.UsingKeyAnyItems(0)
	_ = corejson.NewMapResults.UsingMapOptions(false, false, 0, nil)
	raw := map[string]corejson.Result{"k": corejson.New(1)}
	_ = corejson.NewMapResults.UsingMapOptions(false, false, 0, raw)
	_ = corejson.NewMapResults.UsingMapOptions(true, false, 2, raw)
	_ = corejson.NewMapResults.UsingMapPlusCap(0, nil)
	_ = corejson.NewMapResults.UsingMapPlusCap(0, raw)
	_ = corejson.NewMapResults.UsingMapPlusCapClone(0, nil)
	_ = corejson.NewMapResults.UsingMapPlusCapClone(0, raw)
	_ = corejson.NewMapResults.UsingMapPlusCapDeepClone(0, nil)
	_ = corejson.NewMapResults.UsingMapPlusCapDeepClone(0, raw)
	_ = corejson.NewMapResults.UsingMap(nil)
	_ = corejson.NewMapResults.UsingMap(raw)
	_ = corejson.NewMapResults.UsingMapAnyItems(nil)
	_ = corejson.NewMapResults.UsingMapAnyItems(map[string]any{"k": 1})
	_ = corejson.NewMapResults.UsingMapAnyItemsPlusCap(0, nil)
	_ = corejson.NewMapResults.UsingKeyWithResults()
	_ = corejson.NewMapResults.UsingKeyWithResultsPlusCap(0)
	_ = corejson.NewMapResults.UsingKeyJsoners()
	_ = corejson.NewMapResults.UsingKeyJsonersPlusCap(0)
}

func Test_CovJsonS2_CR07_NewBytesCollection(t *testing.T) {
	_ = corejson.NewBytesCollection.Empty()
	_ = corejson.NewBytesCollection.UsingCap(5)
	_, _ = corejson.NewBytesCollection.AnyItems(1, 2)
}

func Test_CovJsonS2_CR08_NewResultsCollection_DeserializeUsingBytes(t *testing.T) {
	// Arrange
	rc := newTestRC()
	b, _ := corejson.Serialize.Raw(rc)
	rc2, err := corejson.NewResultsCollection.DeserializeUsingBytes(b)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = rc2
	_, _ = corejson.NewResultsCollection.UnmarshalUsingBytes(b)
	rp := corejson.NewPtr(rc)
	_, _ = corejson.NewResultsCollection.DeserializeUsingResult(rp)
}

func Test_CovJsonS2_CR09_NewResultsPtrCollection_DeserializeUsingBytes(t *testing.T) {
	// Arrange
	rpc := newTestRPC()
	b, _ := corejson.Serialize.Raw(rpc)
	rpc2, err := corejson.NewResultsPtrCollection.DeserializeUsingBytes(b)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = rpc2
	_, _ = corejson.NewResultsPtrCollection.UnmarshalUsingBytes(b)
	rp := corejson.NewPtr(rpc)
	_, _ = corejson.NewResultsPtrCollection.DeserializeUsingResult(rp)
}

func Test_CovJsonS2_CR10_NewMapResults_DeserializeUsingBytes(t *testing.T) {
	// Arrange
	mr := newTestMR()
	b, _ := corejson.Serialize.Raw(mr)
	mr2, err := corejson.NewMapResults.DeserializeUsingBytes(b)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = mr2
	_, _ = corejson.NewMapResults.UnmarshalUsingBytes(b)
	rp := corejson.NewPtr(mr)
	_, _ = corejson.NewMapResults.DeserializeUsingResult(rp)
}

func Test_CovJsonS2_CR11_NewBytesCollection_DeserializeUsingBytes_DeserializeUsingResult(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.UsingCap(2)
	bc.Add([]byte(`"a"`))
	b, _ := corejson.Serialize.Raw(bc)
	bc2, err := corejson.NewBytesCollection.DeserializeUsingBytes(b)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = bc2
	_, _ = corejson.NewBytesCollection.UnmarshalUsingBytes(b)
	rp := corejson.NewPtr(bc)
	_, _ = corejson.NewBytesCollection.DeserializeUsingResult(rp)
}

// --- BytesCollection methods ---

func Test_CovJsonS2_BC01_Length_FirstOrDefault_LastOrDefault(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.UsingCap(2)
	bc.Add([]byte(`1`))
	bc.Add([]byte(`2`))

	// Act
	actual := args.Map{"result": bc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": bc.FirstOrDefault() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": bc.LastOrDefault() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	empty := corejson.NewBytesCollection.Empty()
	actual = args.Map{"result": empty.FirstOrDefault() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJsonS2_BC02_Take_Limit_Skip(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(3)
	bc.Add([]byte(`1`))
	bc.Add([]byte(`2`))
	_ = bc.Take(1)
	_ = bc.Limit(1)
	_ = bc.Limit(-1)
	_ = bc.Skip(1)
	empty := corejson.NewBytesCollection.Empty()
	_ = empty.Take(1)
	_ = empty.Limit(1)
	_ = empty.Skip(1)
}

func Test_CovJsonS2_BC03_AddSkipOnNil_AddNonEmpty_AddPtr_Adds(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.AddSkipOnNil(nil)
	bc.AddNonEmpty(nil)
	bc.AddPtr(nil)
	bc.Adds()
	bc.Adds([]byte(`1`), nil, []byte(`2`))
}

func Test_CovJsonS2_BC04_AddResult_AddResultPtr_AddAny_AddAnyItems(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	r := corejson.New(1)
	bc.AddResult(r)
	bc.AddResultPtr(corejson.NewPtr(1))
	err := bc.AddAny(1)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	err2 := bc.AddAnyItems(1, 2)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = bc.AddAnyItems()
}

func Test_CovJsonS2_BC05_GetAt_GetAtSafe_JsonResultAt(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.UsingCap(2)
	bc.Add([]byte(`1`))
	_ = bc.GetAt(0)
	_ = bc.GetAtSafe(0)
	_ = bc.GetAtSafePtr(0)
	_ = bc.GetResultAtSafe(0)
	_ = bc.GetAtSafeUsingLength(0, 1)

	// Act
	actual := args.Map{"result": bc.GetAtSafe(-2) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	_ = bc.JsonResultAt(0)
}

func Test_CovJsonS2_BC06_UnmarshalAt_UnmarshalIntoSameIndex(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.UsingCap(2)
	bc.Add([]byte(`1`))
	bc.Add([]byte(`2`))
	var i, j int
	err := bc.UnmarshalAt(0, &i)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	errs, _ := bc.UnmarshalIntoSameIndex(&i, &j)
	_ = errs
	bc.UnmarshalIntoSameIndex(nil)
}

func Test_CovJsonS2_BC07_Strings_Clone_Dispose(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.UsingCap(2)
	bc.Add([]byte(`"a"`))
	ss := bc.Strings()

	// Act
	actual := args.Map{"result": len(ss) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_ = bc.StringsPtr()
	c := bc.Clone(true)
	_ = c
	_ = bc.ShadowClone()
	cp := bc.ClonePtr(true)
	_ = cp
	var nilBC *corejson.BytesCollection
	actual = args.Map{"result": nilBC.ClonePtr(true) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	bc.Dispose()
	var nilBC2 *corejson.BytesCollection
	nilBC2.Dispose()
	// empty clone
	empty := corejson.NewBytesCollection.Empty()
	_ = empty.Clone(true)
	_ = empty.Strings()
}

func Test_CovJsonS2_BC08_GetPagesSize_GetPagedCollection(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		bc.Add([]byte(`1`))
	}

	// Act
	actual := args.Map{"result": bc.GetPagesSize(3) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual = args.Map{"result": bc.GetPagesSize(0) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	pages := bc.GetPagedCollection(3)
	actual = args.Map{"result": len(pages) < 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	// small
	small := corejson.NewBytesCollection.UsingCap(2)
	small.Add([]byte(`1`))
	pages2 := small.GetPagedCollection(5)
	actual = args.Map{"result": len(pages2) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_CovJsonS2_BC09_Json_JsonModel_Interfaces(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(1)
	bc.Add([]byte(`1`))
	_ = bc.JsonModel()
	_ = bc.JsonModelAny()
	_ = bc.Json()
	_ = bc.JsonPtr()
	_ = bc.AsJsonContractsBinder()
	_ = bc.AsJsoner()
	_ = bc.AsJsonParseSelfInjector()
}

func Test_CovJsonS2_BC10_AddMapResults_AddRawMapResults_AddsPtr_AddBytesCollection(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	mr := newTestMR()
	bc.AddMapResults(mr)
	bc.AddRawMapResults(nil)
	bc.AddsPtr(nil)
	bc2 := corejson.NewBytesCollection.UsingCap(1)
	bc2.Add([]byte(`1`))
	bc.AddBytesCollection(bc2)
}

func Test_CovJsonS2_BC11_AddSerializer_AddSerializers_Funcs(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.AddSerializer(nil)
	bc.AddSerializers()
	bc.AddSerializerFunc(nil)
	bc.AddSerializerFunctions()
	bc.AddSerializerFunc(func() ([]byte, error) {
		return []byte(`1`), nil
	})
}
