package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_C23_MR_Length_Nil(t *testing.T) {
	// Arrange
	var mr *corejson.MapResults

	// Act
	actual := args.Map{"result": mr.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C23_MR_LastIndex(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{"result": mr.LastIndex() != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

func Test_C23_MR_IsEmpty(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{"result": mr.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C23_MR_HasAnyItem(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))

	// Act
	actual := args.Map{"result": mr.HasAnyItem()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C23_MR_AddSkipOnNil_Nil(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddSkipOnNil("k", nil)

	// Act
	actual := args.Map{"result": mr.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C23_MR_AddSkipOnNil_Valid(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	r := corejson.New("v")
	mr.AddSkipOnNil("k", r.Ptr())

	// Act
	actual := args.Map{"result": mr.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C23_MR_GetByKey_Found(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	r := mr.GetByKey("k")

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C23_MR_GetByKey_NotFound(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	r := mr.GetByKey("missing")

	// Act
	actual := args.Map{"result": r != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C23_MR_HasError_Yes(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Create(nil, errors.New("err"), ""))

	// Act
	actual := args.Map{"result": mr.HasError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C23_MR_HasError_No(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))

	// Act
	actual := args.Map{"result": mr.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C23_MR_AllErrors_Empty(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	errs, has := mr.AllErrors()

	// Act
	actual := args.Map{"result": has || len(errs) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C23_MR_AllErrors_WithErrors(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Create(nil, errors.New("err"), ""))
	errs, has := mr.AllErrors()

	// Act
	actual := args.Map{"result": has || len(errs) != 1}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C23_MR_GetErrorsStrings_Empty(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	s := mr.GetErrorsStrings()

	// Act
	actual := args.Map{"result": len(s) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C23_MR_GetErrorsStrings_WithErrors(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Create(nil, errors.New("err"), ""))
	s := mr.GetErrorsStrings()

	// Act
	actual := args.Map{"result": len(s) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C23_MR_GetErrorsStringsPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.GetErrorsStringsPtr()
}

func Test_C23_MR_GetErrorsAsSingleString(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.GetErrorsAsSingleString()
}

func Test_C23_MR_GetErrorsAsSingle(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	err := mr.GetErrorsAsSingle()
	_ = err
}

func Test_C23_MR_Unmarshal(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	var s string
	err := mr.Unmarshal("k", &s)
	_ = err
}

func Test_C23_MR_Deserialize(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	var s string
	_ = mr.Deserialize("k", &s)
}

func Test_C23_MR_DeserializeMust(t *testing.T) {
	defer func() { recover() }()
	mr := corejson.NewMapResults.Empty()
	var s string
	mr.DeserializeMust("missing", &s)
}

func Test_C23_MR_UnmarshalMany_Empty(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	err := mr.UnmarshalMany()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C23_MR_UnmarshalMany_Valid(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	var s string
	err := mr.UnmarshalMany(corejson.KeyAny{Key: "k", AnyInf: &s})
	_ = err
}

func Test_C23_MR_UnmarshalManySafe_Empty(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	err := mr.UnmarshalManySafe()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C23_MR_SafeUnmarshal(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	var s string
	_ = mr.SafeUnmarshal("k", &s)
}

func Test_C23_MR_SafeDeserialize(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.SafeDeserialize("k", nil)
}

func Test_C23_MR_SafeDeserializeMust(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.SafeDeserializeMust("k", nil)
}

func Test_C23_MR_InjectIntoAt(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New(map[string]string{"a": "b"}))
	target := corejson.Empty.MapResults()
	err := mr.InjectIntoAt("k", target)
	_ = err
}

func Test_C23_MR_AddPtr_Nil(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddPtr("k", nil)

	// Act
	actual := args.Map{"result": mr.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C23_MR_AddAny_Nil(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	err := mr.AddAny("k", nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C23_MR_AddAny_Valid(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	err := mr.AddAny("k", "v")

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C23_MR_AddAny_Error(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	err := mr.AddAny("k", make(chan int))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C23_MR_AddAnySkipOnNil_Nil(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	err := mr.AddAnySkipOnNil("k", nil)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C23_MR_AddAnySkipOnNil_Valid(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	err := mr.AddAnySkipOnNil("k", "v")

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_C23_MR_AddAnyNonEmptyNonError(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddAnyNonEmptyNonError("k", nil)
	mr.AddAnyNonEmptyNonError("k2", "v")
	_ = mr
}

func Test_C23_MR_AddAnyNonEmpty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddAnyNonEmpty("k", nil)
	mr.AddAnyNonEmpty("k2", "v")
}

func Test_C23_MR_AddKeyWithResult(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyWithResult(corejson.KeyWithResult{Key: "k", Result: corejson.New("v")})
}

func Test_C23_MR_AddKeyWithResultPtr_Nil(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyWithResultPtr(nil)

	// Act
	actual := args.Map{"result": mr.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C23_MR_AddKeysWithResults(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeysWithResults(corejson.KeyWithResult{Key: "k", Result: corejson.New("v")})
}

func Test_C23_MR_AddKeysWithResults_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeysWithResults()
}

func Test_C23_MR_AddKeysWithResultsPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeysWithResultsPtr()
}

func Test_C23_MR_AddKeyAnyInf(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyInf(corejson.KeyAny{Key: "k", AnyInf: "v"})
}

func Test_C23_MR_AddKeyAnyInfPtr_Nil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyInfPtr(nil)
}

func Test_C23_MR_AddKeyAnyInfPtr_Valid(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyInfPtr(&corejson.KeyAny{Key: "k", AnyInf: "v"})
}

func Test_C23_MR_AddKeyAnyItems(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyItems(corejson.KeyAny{Key: "k", AnyInf: "v"})
}

func Test_C23_MR_AddKeyAnyItems_Nil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyItems()
}

func Test_C23_MR_AddKeyAnyItemsPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyItemsPtr()
}

func Test_C23_MR_AddNonEmptyNonErrorPtr_Nil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddNonEmptyNonErrorPtr("k", nil)
}

func Test_C23_MR_AddNonEmptyNonErrorPtr_Error(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	r := corejson.NewResult.ErrorPtr(errors.New("err"))
	mr.AddNonEmptyNonErrorPtr("k", r)

	// Act
	actual := args.Map{"result": mr.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C23_MR_AddNonEmptyNonErrorPtr_Valid(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := corejson.New("v")
	mr.AddNonEmptyNonErrorPtr("k", r.Ptr())
}

func Test_C23_MR_AddMapResults_Nil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddMapResults(nil)
}

func Test_C23_MR_AddMapResults_Valid(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	other := corejson.NewMapResults.Empty()
	other.Add("k", corejson.New("v"))
	mr.AddMapResults(other)
}

func Test_C23_MR_AddMapAnyItems(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddMapAnyItems(map[string]any{"k": "v"})
}

func Test_C23_MR_AddMapAnyItems_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddMapAnyItems(map[string]any{})
}

func Test_C23_MR_AllKeys(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("a", corejson.New("1"))
	mr.Add("b", corejson.New("2"))
	keys := mr.AllKeys()

	// Act
	actual := args.Map{"result": len(keys) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C23_MR_AllKeys_Empty(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	keys := mr.AllKeys()

	// Act
	actual := args.Map{"result": len(keys) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C23_MR_AllKeysSorted(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("b", corejson.New("2"))
	mr.Add("a", corejson.New("1"))
	keys := mr.AllKeysSorted()

	// Act
	actual := args.Map{"result": keys[0] != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)
}

func Test_C23_MR_AllKeysSorted_Empty(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	keys := mr.AllKeysSorted()

	// Act
	actual := args.Map{"result": len(keys) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C23_MR_AllValues(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	vals := mr.AllValues()

	// Act
	actual := args.Map{"result": len(vals) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C23_MR_AllValues_Empty(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	vals := mr.AllValues()

	// Act
	actual := args.Map{"result": len(vals) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C23_MR_AllResultsCollection(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	rc := mr.AllResultsCollection()

	// Act
	actual := args.Map{"result": rc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C23_MR_AllResultsCollection_Empty(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	rc := mr.AllResultsCollection()

	// Act
	actual := args.Map{"result": rc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C23_MR_AllResults(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.AllResults()
}

func Test_C23_MR_GetStrings(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	s := mr.GetStrings()

	// Act
	actual := args.Map{"result": len(s) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C23_MR_GetStrings_Empty(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	s := mr.GetStrings()

	// Act
	actual := args.Map{"result": len(s) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C23_MR_GetStringsPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.GetStringsPtr()
}

func Test_C23_MR_AddJsoner_Nil(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddJsoner("k", nil)

	// Act
	actual := args.Map{"result": mr.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C23_MR_AddKeyWithJsoner(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := corejson.New("v")
	mr.AddKeyWithJsoner(corejson.KeyWithJsoner{Key: "k", Jsoner: &r})
}

func Test_C23_MR_AddKeysWithJsoners(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeysWithJsoners()
}

func Test_C23_MR_AddKeyWithJsonerPtr_Nil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyWithJsonerPtr(nil)
}

func Test_C23_MR_AddKeyWithJsonerPtr_NilJsoner(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyWithJsonerPtr(&corejson.KeyWithJsoner{Key: "k", Jsoner: nil})
}

func Test_C23_MR_AddMapResultsUsingCloneOption_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddMapResultsUsingCloneOption(false, false, map[string]corejson.Result{})
}

func Test_C23_MR_AddMapResultsUsingCloneOption_NoClone(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddMapResultsUsingCloneOption(false, false, map[string]corejson.Result{
		"k": corejson.New("v"),
	})
}

func Test_C23_MR_AddMapResultsUsingCloneOption_Clone(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddMapResultsUsingCloneOption(true, true, map[string]corejson.Result{
		"k": corejson.New("v"),
	})
}

func Test_C23_MR_GetPagesSize_Zero(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{"result": mr.GetPagesSize(0) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C23_MR_GetPagesSize_Valid(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	for i := 0; i < 5; i++ { mr.Add(string(rune('a'+i)), corejson.New(i)) }
	p := mr.GetPagesSize(2)

	// Act
	actual := args.Map{"result": p}

	// Assert
	expected := args.Map{"result": 3}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C23_MR_GetPagedCollection_Small(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("a", corejson.New("1"))
	result := mr.GetPagedCollection(5)

	// Act
	actual := args.Map{"result": len(result) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C23_MR_GetPagedCollection_Multi(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	for i := 0; i < 5; i++ { mr.Add(string(rune('a'+i)), corejson.New(i)) }
	result := mr.GetPagedCollection(2)

	// Act
	actual := args.Map{"result": len(result)}

	// Assert
	expected := args.Map{"result": 3}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C23_MR_GetNewMapUsingKeys_Empty(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	result := mr.GetNewMapUsingKeys(false)

	// Act
	actual := args.Map{"result": result.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C23_MR_GetNewMapUsingKeys_Valid(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("a", corejson.New("1"))
	mr.Add("b", corejson.New("2"))
	result := mr.GetNewMapUsingKeys(false, "a")

	// Act
	actual := args.Map{"result": result.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C23_MR_ResultCollection(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	rc := mr.ResultCollection()
	_ = rc
}

func Test_C23_MR_ResultCollection_Empty(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	rc := mr.ResultCollection()

	// Act
	actual := args.Map{"result": rc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C23_MR_JsonModel(t *testing.T) { _ = corejson.NewMapResults.Empty().JsonModel() }
	// Arrange
func Test_C23_MR_JsonModelAny(t *testing.T) { _ = corejson.NewMapResults.Empty().JsonModelAny() }

func Test_C23_MR_Clear(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.New("v"))
	mr.Clear()
}

func Test_C23_MR_Clear_Nil(t *testing.T) {
	var mr *corejson.MapResults
	mr.Clear()
}

func Test_C23_MR_Dispose(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Dispose()
}

func Test_C23_MR_Dispose_Nil(t *testing.T) {
	var mr *corejson.MapResults
	mr.Dispose()
}

func Test_C23_MR_Json(t *testing.T) { _ = corejson.NewMapResults.Empty().Json() }
func Test_C23_MR_JsonPtr(t *testing.T) { _ = corejson.NewMapResults.Empty().JsonPtr() }

func Test_C23_MR_ParseInjectUsingJson_Error(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	bad := corejson.NewResult.UsingString(`invalid`)
	_, err := mr.ParseInjectUsingJson(bad)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C23_MR_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() { recover() }()
	mr := corejson.NewMapResults.Empty()
	bad := corejson.NewResult.UsingString(`invalid`)
	mr.ParseInjectUsingJsonMust(bad)
}

func Test_C23_MR_AsJsonContractsBinder(t *testing.T) { _ = corejson.NewMapResults.Empty().AsJsonContractsBinder() }
func Test_C23_MR_AsJsoner(t *testing.T) { _ = corejson.NewMapResults.Empty().AsJsoner() }
func Test_C23_MR_AsJsonParseSelfInjector(t *testing.T) { _ = corejson.NewMapResults.Empty().AsJsonParseSelfInjector() }
func Test_C23_MR_JsonParseSelfInject(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := corejson.New(*mr)
	_ = mr.JsonParseSelfInject(&r)
}
