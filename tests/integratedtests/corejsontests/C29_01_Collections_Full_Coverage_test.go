package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════
// ResultsCollection — all uncovered methods
// ═══════════════════════════════════════════════

func Test_C29_01_RC_Length(t *testing.T) {
	// Arrange
	var rc *corejson.ResultsCollection

	// Act
	actual := args.Map{"result": rc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	rc2 := &corejson.ResultsCollection{}
	actual = args.Map{"result": rc2.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C29_02_RC_LastIndex(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()

	// Act
	actual := args.Map{"result": rc.LastIndex() != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

func Test_C29_03_RC_IsEmpty_HasAnyItem(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()

	// Act
	actual := args.Map{"result": rc.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": rc.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C29_04_RC_FirstOrDefault(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()

	// Act
	actual := args.Map{"result": rc.FirstOrDefault() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	rc.Add(corejson.NewResult.Any("x"))
	actual = args.Map{"result": rc.FirstOrDefault() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C29_05_RC_LastOrDefault(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()

	// Act
	actual := args.Map{"result": rc.LastOrDefault() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	rc.Add(corejson.NewResult.Any("x"))
	actual = args.Map{"result": rc.LastOrDefault() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C29_06_RC_Take(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	taken := rc.Take(1)

	// Act
	actual := args.Map{"result": taken.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	rc.Add(corejson.NewResult.Any("a"))
	rc.Add(corejson.NewResult.Any("b"))
	taken = rc.Take(1)
	actual = args.Map{"result": taken.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_07_RC_Limit(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	l := rc.Limit(5)

	// Act
	actual := args.Map{"result": l.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	rc.Add(corejson.NewResult.Any("a"))
	rc.Add(corejson.NewResult.Any("b"))
	// TakeAllMinusOne is -1
	l = rc.Limit(-1)
	actual = args.Map{"result": l.Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	l = rc.Limit(1)
	actual = args.Map{"result": l.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_08_RC_Skip(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	s := rc.Skip(0)

	// Act
	actual := args.Map{"result": s.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	rc.Add(corejson.NewResult.Any("a"))
	rc.Add(corejson.NewResult.Any("b"))
	s = rc.Skip(1)
	actual = args.Map{"result": s.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_09_RC_AddSkipOnNil(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSkipOnNil(nil)

	// Act
	actual := args.Map{"result": rc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	r := corejson.NewResult.AnyPtr("x")
	rc.AddSkipOnNil(r)
	actual = args.Map{"result": rc.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_10_RC_AddNonNilNonError(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddNonNilNonError(nil)
	rc.AddNonNilNonError(&corejson.Result{Error: errors.New("e")})

	// Act
	actual := args.Map{"result": rc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	rc.AddNonNilNonError(corejson.NewResult.AnyPtr("x"))
	actual = args.Map{"result": rc.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_11_RC_GetAt(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	r := rc.GetAt(0)

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C29_12_RC_HasError(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()

	// Act
	actual := args.Map{"result": rc.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	rc.Add(corejson.NewResult.Error(errors.New("e")))
	actual = args.Map{"result": rc.HasError()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C29_13_RC_AllErrors(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	errs, has := rc.AllErrors()

	// Act
	actual := args.Map{"result": has || len(errs) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	rc.Add(corejson.NewResult.Any("x"))
	rc.Add(corejson.NewResult.Error(errors.New("e")))
	errs, has = rc.AllErrors()
	actual = args.Map{"result": has || len(errs) != 1}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C29_14_RC_GetErrorsStrings(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	s := rc.GetErrorsStrings()

	// Act
	actual := args.Map{"result": len(s) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	rc.Add(corejson.NewResult.Any("x"))
	rc.Add(corejson.NewResult.Error(errors.New("e")))
	s = rc.GetErrorsStrings()
	actual = args.Map{"result": len(s) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_15_RC_GetErrorsStringsPtr(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	s := rc.GetErrorsStringsPtr()
	_ = s
}

func Test_C29_16_RC_GetErrorsAsSingleString(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	s := rc.GetErrorsAsSingleString()
	_ = s
}

func Test_C29_17_RC_GetErrorsAsSingle(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	err := rc.GetErrorsAsSingle()
	_ = err
}

func Test_C29_18_RC_UnmarshalAt(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("hello"))
	var s string
	err := rc.UnmarshalAt(0, &s)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C29_19_RC_InjectIntoAt(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	r := corejson.NewResult.Any(corejson.Result{Bytes: []byte(`"test"`), TypeName: "T"})
	rc.Add(r)
	target := corejson.Result{}
	err := rc.InjectIntoAt(0, &target)
	_ = err
}

func Test_C29_20_RC_InjectIntoSameIndex(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	// Pass true nil variadic slice
	var nilSlice []corejson.JsonParseSelfInjector
	errs, has := rc.InjectIntoSameIndex(nilSlice...)

	// Act
	actual := args.Map{"result": has || len(errs) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C29_21_RC_UnmarshalIntoSameIndex(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	var nilSlice []any
	errs, has := rc.UnmarshalIntoSameIndex(nilSlice...)

	// Act
	actual := args.Map{"result": has || len(errs) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)

	rc.Add(corejson.NewResult.Any("hello"))
	rc.Add(corejson.NewResult.Error(errors.New("e")))
	rc.Add(corejson.NewResult.Any("world"))
	var s1 string
	var s3 string
	errs, has = rc.UnmarshalIntoSameIndex(&s1, nil, &s3)
	_ = errs
	_ = has
}

func Test_C29_22_RC_GetAtSafe(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	r := rc.GetAtSafe(0)

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	r = rc.GetAtSafe(-1)
	actual = args.Map{"result": r != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	r = rc.GetAtSafe(999)
	actual = args.Map{"result": r != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C29_23_RC_GetAtSafeUsingLength(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	r := rc.GetAtSafeUsingLength(0, 1)

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	r = rc.GetAtSafeUsingLength(5, 1)
	actual = args.Map{"result": r != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C29_24_RC_AddPtr(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddPtr(nil)

	// Act
	actual := args.Map{"result": rc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	rc.AddPtr(corejson.NewResult.AnyPtr("x"))
	actual = args.Map{"result": rc.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_25_RC_Add(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))

	// Act
	actual := args.Map{"result": rc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_26_RC_Adds(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.Adds()

	// Act
	actual := args.Map{"result": rc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	rc.Adds(corejson.NewResult.Any("a"), corejson.NewResult.Any("b"))
	actual = args.Map{"result": rc.Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C29_27_RC_AddSerializer(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSerializer(nil)

	// Act
	actual := args.Map{"result": rc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C29_28_RC_AddSerializers(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSerializers()

	// Act
	actual := args.Map{"result": rc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C29_29_RC_AddSerializerFunc(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSerializerFunc(nil)

	// Act
	actual := args.Map{"result": rc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	rc.AddSerializerFunc(func() ([]byte, error) {
		return []byte(`"x"`), nil
	})
	actual = args.Map{"result": rc.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_30_RC_AddSerializerFunctions(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddSerializerFunctions()

	// Act
	actual := args.Map{"result": rc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C29_31_RC_AddMapResults(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	mr := corejson.NewMapResults.Empty()
	rc.AddMapResults(mr)

	// Act
	actual := args.Map{"result": rc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	mr.Add("k", corejson.NewResult.Any("v"))
	rc.AddMapResults(mr)
	actual = args.Map{"result": rc.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_32_RC_AddRawMapResults(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddRawMapResults(nil)

	// Act
	actual := args.Map{"result": rc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C29_33_RC_AddsPtr(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddsPtr(nil, corejson.NewResult.AnyPtr("x"))

	// Act
	actual := args.Map{"result": rc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_34_RC_AddAny(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddAny(nil)
	rc.AddAny("hello")

	// Act
	actual := args.Map{"result": rc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_35_RC_AddAnyItems(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddAnyItems(nil, "a", nil, "b")

	// Act
	actual := args.Map{"result": rc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C29_36_RC_AddAnyItemsSlice(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddAnyItemsSlice(nil)
	rc.AddAnyItemsSlice([]any{nil, "a"})

	// Act
	actual := args.Map{"result": rc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_37_RC_AddResultsCollection(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddResultsCollection(nil)
	other := corejson.NewResultsCollection.Empty()
	other.Add(corejson.NewResult.Any("x"))
	rc.AddResultsCollection(other)

	// Act
	actual := args.Map{"result": rc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_38_RC_AddNonNilItemsPtr(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddNonNilItemsPtr()
	rc.AddNonNilItemsPtr(nil, corejson.NewResult.AnyPtr("x"))

	// Act
	actual := args.Map{"result": rc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_39_RC_NonPtr_Ptr(t *testing.T) {
	rc := corejson.ResultsCollection{}
	_ = rc.NonPtr()
	_ = rc.Ptr()
}

func Test_C29_40_RC_Clear(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	rc.Clear()

	// Act
	actual := args.Map{"result": rc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty after clear", actual)
}

func Test_C29_41_RC_Clear_Nil(t *testing.T) {
	var rc *corejson.ResultsCollection
	result := rc.Clear()
	_ = result
}

func Test_C29_42_RC_Dispose(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	rc.Dispose()
}

func Test_C29_43_RC_Dispose_Nil(t *testing.T) {
	var rc *corejson.ResultsCollection
	rc.Dispose()
}

func Test_C29_44_RC_GetStrings(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	s := rc.GetStrings()

	// Act
	actual := args.Map{"result": len(s) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	rc.Add(corejson.NewResult.Any("hello"))
	s = rc.GetStrings()
	actual = args.Map{"result": len(s) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_45_RC_GetStringsPtr(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	_ = rc.GetStringsPtr()
}

func Test_C29_46_RC_AddJsoners(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.AddJsoners(true)

	// Act
	actual := args.Map{"result": rc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C29_47_RC_GetPagesSize(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()

	// Act
	actual := args.Map{"result": rc.GetPagesSize(0) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": rc.GetPagesSize(-1) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	for i := 0; i < 5; i++ {
		rc.Add(corejson.NewResult.Any(i))
	}
	actual = args.Map{"result": rc.GetPagesSize(2) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C29_48_RC_GetPagedCollection(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	for i := 0; i < 5; i++ {
		rc.Add(corejson.NewResult.Any(i))
	}
	pages := rc.GetPagedCollection(2)

	// Act
	actual := args.Map{"result": len(pages) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 pages", actual)
}

func Test_C29_49_RC_GetPagedCollection_SmallSize(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	pages := rc.GetPagedCollection(10)

	// Act
	actual := args.Map{"result": len(pages) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 page", actual)
}

func Test_C29_50_RC_GetSinglePageCollection(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	for i := 0; i < 10; i++ {
		rc.Add(corejson.NewResult.Any(i))
	}
	page := rc.GetSinglePageCollection(3, 1)

	// Act
	actual := args.Map{"result": page.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	page = rc.GetSinglePageCollection(3, 4)
	actual = args.Map{"result": page.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_51_RC_GetSinglePageCollection_Small(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	page := rc.GetSinglePageCollection(10, 1)

	// Act
	actual := args.Map{"result": page.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C29_52_RC_JsonModel_JsonModelAny(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	_ = rc.JsonModel()
	_ = rc.JsonModelAny()
}

func Test_C29_53_RC_Json_JsonPtr(t *testing.T) {
	rc := corejson.ResultsCollection{}
	_ = rc.Json()
	_ = rc.JsonPtr()
}

func Test_C29_54_RC_ParseInjectUsingJson(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	serialized := rc.JsonPtr()
	rc2 := corejson.NewResultsCollection.Empty()
	_, err := rc2.ParseInjectUsingJson(serialized)
	_ = err
}

func Test_C29_55_RC_ParseInjectUsingJson_Fail(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	bad := &corejson.Result{Error: errors.New("fail")}
	_, err := rc.ParseInjectUsingJson(bad)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C29_56_RC_ParseInjectUsingJsonMust(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	serialized := rc.JsonPtr()
	rc2 := corejson.NewResultsCollection.Empty()
	_ = rc2.ParseInjectUsingJsonMust(serialized)
}

func Test_C29_57_RC_AsJsonContractsBinder(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	_ = rc.AsJsonContractsBinder()
}

func Test_C29_58_RC_AsJsoner(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	_ = rc.AsJsoner()
}

func Test_C29_59_RC_JsonParseSelfInject(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	serialized := rc.JsonPtr()
	rc2 := corejson.NewResultsCollection.Empty()
	err := rc2.JsonParseSelfInject(serialized)
	_ = err
}

func Test_C29_60_RC_AsJsonParseSelfInjector(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	_ = rc.AsJsonParseSelfInjector()
}

func Test_C29_61_RC_ShadowClone(t *testing.T) {
	rc := corejson.ResultsCollection{}
	_ = rc.ShadowClone()
}

func Test_C29_62_RC_Clone(t *testing.T) {
	rc := corejson.ResultsCollection{}
	_ = rc.Clone(false)
	rc.Items = []corejson.Result{corejson.NewResult.Any("x")}
	_ = rc.Clone(true)
}

func Test_C29_63_RC_ClonePtr(t *testing.T) {
	// Arrange
	var rc *corejson.ResultsCollection

	// Act
	actual := args.Map{"result": rc.ClonePtr(false) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	rc = corejson.NewResultsCollection.Empty()
	_ = rc.ClonePtr(false)
	rc.Add(corejson.NewResult.Any("x"))
	_ = rc.ClonePtr(true)
}

// ─── UnmarshalIntoSameIndex edge: empty json bytes item ───

func Test_C29_64_RC_UnmarshalIntoSameIndex_EmptyJsonBytes(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Create([]byte(`{}`), nil, "T"))
	var m map[string]any
	errs, has := rc.UnmarshalIntoSameIndex(&m)
	_ = errs
	_ = has
}

// ─── InjectIntoSameIndex with error result and valid injector ───

func Test_C29_65_RC_InjectIntoSameIndex_ErrorResult(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Error(errors.New("e")))
	r := corejson.Result{}
	errs, has := rc.InjectIntoSameIndex(&r)

	// Act
	actual := args.Map{"result": has}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected hasAnyError true", actual)
	_ = errs
}
