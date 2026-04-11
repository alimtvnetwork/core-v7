package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// ResultsPtrCollection — Length / IsEmpty / HasAnyItem
// =============================================================================

func Test_PtrColl_NilLength(t *testing.T) {
	tc := ptrCollNilLengthTestCase

	// Arrange
	var c *corejson.ResultsPtrCollection

	// Act
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_IsEmpty(t *testing.T) {
	tc := ptrCollIsEmptyTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	actual := args.Map{
		"isEmpty": c.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_HasAnyItem(t *testing.T) {
	tc := ptrCollHasAnyItemTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()
	c.Add(corejson.NewPtr("a"))

	// Act
	actual := args.Map{
		"hasAny": c.HasAnyItem(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// ResultsPtrCollection — FirstOrDefault / LastOrDefault
// =============================================================================

func Test_PtrColl_FirstOrDefault_Empty(t *testing.T) {
	tc := ptrCollFirstOrDefaultEmptyTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	actual := args.Map{
		"isNil": c.FirstOrDefault() == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_FirstOrDefault_Valid(t *testing.T) {
	tc := ptrCollFirstOrDefaultValidTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()
	c.Add(corejson.NewPtr("a"))

	// Act
	actual := args.Map{
		"isNil": c.FirstOrDefault() == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_LastOrDefault_Empty(t *testing.T) {
	tc := ptrCollLastOrDefaultEmptyTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	actual := args.Map{
		"isNil": c.LastOrDefault() == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_LastOrDefault_Valid(t *testing.T) {
	tc := ptrCollLastOrDefaultValidTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()
	c.Add(corejson.NewPtr("a"))

	// Act
	actual := args.Map{
		"isNil": c.LastOrDefault() == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// ResultsPtrCollection — Take / Limit / Skip
// =============================================================================

func Test_PtrColl_Take_Empty(t *testing.T) {
	tc := ptrCollTakeEmptyTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	r := c.Take(5)
	actual := args.Map{
		"isEmpty": r.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_Take_Valid(t *testing.T) {
	tc := ptrCollTakeValidTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()
	c.Add(corejson.NewPtr("a")).Add(corejson.NewPtr("b")).Add(corejson.NewPtr("c"))

	// Act
	r := c.Take(2)
	actual := args.Map{
		"length": r.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_Limit_TakeAll(t *testing.T) {
	tc := ptrCollLimitTakeAllTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()
	c.Add(corejson.NewPtr("a")).Add(corejson.NewPtr("b"))

	// Act
	r := c.Limit(-1)
	actual := args.Map{
		"length": r.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_Skip_Empty(t *testing.T) {
	tc := ptrCollSkipEmptyTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	r := c.Skip(1)
	actual := args.Map{
		"isEmpty": r.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_Skip_Valid(t *testing.T) {
	tc := ptrCollSkipValidTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()
	c.Add(corejson.NewPtr("a")).Add(corejson.NewPtr("b")).Add(corejson.NewPtr("c"))

	// Act
	r := c.Skip(1)
	actual := args.Map{
		"length": r.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// ResultsPtrCollection — Add variants
// =============================================================================

func Test_PtrColl_AddSkipOnNil_Nil(t *testing.T) {
	tc := ptrCollAddSkipOnNilNilTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	c.AddSkipOnNil(nil)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_AddSkipOnNil_Valid(t *testing.T) {
	tc := ptrCollAddSkipOnNilValidTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	c.AddSkipOnNil(corejson.NewPtr("a"))
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_AddNonNilNonError_Nil(t *testing.T) {
	tc := ptrCollAddNonNilNonErrorNilTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	c.AddNonNilNonError(nil)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_AddNonNilNonError_Valid(t *testing.T) {
	tc := ptrCollAddNonNilNonErrorValidTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	c.AddNonNilNonError(corejson.NewPtr("a"))
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_AddResult(t *testing.T) {
	tc := ptrCollAddResultTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	c.AddResult(corejson.New("a"))
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_Adds_Nil(t *testing.T) {
	tc := ptrCollAddsNilTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act — Adds(nil) with variadic creates []*Result{nil}, which is not nil slice
	// The nil element gets appended, so length is 1
	c.Adds(nil)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_AddAny_Nil(t *testing.T) {
	tc := ptrCollAddAnyNilTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	c.AddAny(nil)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_AddAny_Valid(t *testing.T) {
	tc := ptrCollAddAnyValidTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	c.AddAny("hello")
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_AddAnyItems_Nil(t *testing.T) {
	tc := ptrCollAddAnyItemsNilTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	c.AddAnyItems(nil)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_AddResultsCollection_Nil(t *testing.T) {
	tc := ptrCollAddResultsCollectionNilTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	c.AddResultsCollection(nil)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// ResultsPtrCollection — HasError / AllErrors / GetErrorsStrings
// =============================================================================

func Test_PtrColl_HasError_False(t *testing.T) {
	tc := ptrCollHasErrorFalseTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()
	c.Add(corejson.NewPtr("a"))

	// Act
	actual := args.Map{
		"hasError": c.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_HasError_True(t *testing.T) {
	tc := ptrCollHasErrorTrueTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()
	c.Add(&corejson.Result{Error: errors.New("fail")})

	// Act
	actual := args.Map{
		"hasError": c.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_AllErrors_Empty(t *testing.T) {
	tc := ptrCollAllErrorsEmptyTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	errList, hasErr := c.AllErrors()
	actual := args.Map{
		"errorCount": len(errList),
		"hasAnyErr":  hasErr,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_GetErrorsStrings_Empty(t *testing.T) {
	tc := ptrCollGetErrorsStringsEmptyTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	actual := args.Map{
		"length": len(c.GetErrorsStrings()),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_GetErrorsAsSingleString(t *testing.T) {
	tc := ptrCollGetErrorsAsSingleStringTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()
	c.Add(&corejson.Result{Error: errors.New("fail")})

	// Act
	s := c.GetErrorsAsSingleString()
	actual := args.Map{
		"hasContent": len(s) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// ResultsPtrCollection — GetAtSafe / Clear / Dispose / Clone / GetStrings
// =============================================================================

func Test_PtrColl_GetAtSafe_OutOfRange(t *testing.T) {
	tc := ptrCollGetAtSafeOutOfRangeTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	actual := args.Map{
		"isNil": c.GetAtSafe(5) == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_GetAtSafeUsingLength_Out(t *testing.T) {
	tc := ptrCollGetAtSafeUsingLengthOutTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	actual := args.Map{
		"isNil": c.GetAtSafeUsingLength(5, 0) == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_Clear(t *testing.T) {
	tc := ptrCollClearTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()
	c.Add(corejson.NewPtr("a"))

	// Act
	c.Clear()
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_Clear_Nil(t *testing.T) {
	tc := ptrCollClearNilTestCase

	// Arrange
	var c *corejson.ResultsPtrCollection

	// Act
	r := c.Clear()
	actual := args.Map{
		"isNil": r == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_Dispose(t *testing.T) {
	tc := ptrCollDisposeTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()
	c.Add(corejson.NewPtr("a"))

	// Act
	c.Dispose()
	actual := args.Map{
		"nilItems": c.Items == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_GetStrings_Empty(t *testing.T) {
	tc := ptrCollGetStringsEmptyTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	actual := args.Map{
		"length": len(c.GetStrings()),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_GetPagesSize_Zero(t *testing.T) {
	tc := ptrCollGetPagesSizeZeroTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()

	// Act
	actual := args.Map{
		"pages": c.GetPagesSize(0),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_Clone_Nil(t *testing.T) {
	tc := ptrCollCloneNilTestCase

	// Arrange
	var c *corejson.ResultsPtrCollection

	// Act
	r := c.Clone(false)
	actual := args.Map{
		"isNil": r == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PtrColl_Json(t *testing.T) {
	tc := ptrCollJsonTestCase

	// Arrange
	c := corejson.NewResultsPtrCollection.Empty()
	c.Add(corejson.NewPtr("a"))

	// Act
	r := c.Json()
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// MapResults — Length / IsEmpty / HasAnyItem
// =============================================================================

func Test_MapResults_NilLength(t *testing.T) {
	tc := mapResultsNilLengthTestCase

	// Arrange
	var m *corejson.MapResults

	// Act
	actual := args.Map{
		"length": m.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_IsEmpty(t *testing.T) {
	tc := mapResultsIsEmptyTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{
		"isEmpty": m.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_HasAnyItem(t *testing.T) {
	tc := mapResultsHasAnyItemTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.New("v"))

	// Act
	actual := args.Map{
		"hasAny": m.HasAnyItem(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// MapResults — GetByKey / AddSkipOnNil / HasError / AllErrors
// =============================================================================

func Test_MapResults_GetByKey_Missing(t *testing.T) {
	tc := mapResultsGetByKeyMissingTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{
		"isNil": m.GetByKey("x") == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_GetByKey_Found(t *testing.T) {
	tc := mapResultsGetByKeyFoundTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.New("v"))

	// Act
	actual := args.Map{
		"isNil": m.GetByKey("k") == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_AddSkipOnNil_Nil(t *testing.T) {
	tc := mapResultsAddSkipOnNilNilTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	m.AddSkipOnNil("k", nil)
	actual := args.Map{
		"length": m.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_AddSkipOnNil_Valid(t *testing.T) {
	tc := mapResultsAddSkipOnNilValidTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()
	r := corejson.NewPtr("v")

	// Act
	m.AddSkipOnNil("k", r)
	actual := args.Map{
		"length": m.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_HasError_False(t *testing.T) {
	tc := mapResultsHasErrorFalseTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.New("v"))

	// Act
	actual := args.Map{
		"hasError": m.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_AllErrors_Empty(t *testing.T) {
	tc := mapResultsAllErrorsEmptyTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	errList, hasErr := m.AllErrors()
	actual := args.Map{
		"errorCount": len(errList),
		"hasAnyErr":  hasErr,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_GetErrorsStrings_Empty(t *testing.T) {
	tc := mapResultsGetErrorsStringsEmptyTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{
		"length": len(m.GetErrorsStrings()),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// MapResults — Add / AddPtr / AddAny / AddAnySkipOnNil
// =============================================================================

func Test_MapResults_Add(t *testing.T) {
	tc := mapResultsAddTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	m.Add("k", corejson.New("v"))
	actual := args.Map{
		"length": m.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_AddPtr_Nil(t *testing.T) {
	tc := mapResultsAddPtrNilTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	m.AddPtr("k", nil)
	actual := args.Map{
		"length": m.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_AddAny_Nil(t *testing.T) {
	tc := mapResultsAddAnyNilTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	err := m.AddAny("k", nil)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_AddAny_Valid(t *testing.T) {
	tc := mapResultsAddAnyValidTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	err := m.AddAny("k", "hello")
	actual := args.Map{
		"hasError": err != nil,
		"length":   m.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_AddAnySkipOnNil_Nil_FromPtrCollNilLengthResu(t *testing.T) {
	tc := mapResultsAddAnySkipOnNilNilTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	err := m.AddAnySkipOnNil("k", nil)
	actual := args.Map{
		"hasError": err != nil,
		"length":   m.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// MapResults — AllKeys / AllKeysSorted / AllValues / Clear / Dispose
// =============================================================================

func Test_MapResults_AllKeys_Empty(t *testing.T) {
	tc := mapResultsAllKeysEmptyTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{
		"length": len(m.AllKeys()),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_AllKeysSorted_Empty(t *testing.T) {
	tc := mapResultsAllKeysSortedEmptyTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{
		"length": len(m.AllKeysSorted()),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_AllValues_Empty(t *testing.T) {
	tc := mapResultsAllValuesEmptyTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{
		"length": len(m.AllValues()),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_Clear(t *testing.T) {
	tc := mapResultsClearTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.New("v"))

	// Act
	m.Clear()
	actual := args.Map{
		"length": m.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_Clear_Nil(t *testing.T) {
	tc := mapResultsClearNilTestCase

	// Arrange
	var m *corejson.MapResults

	// Act
	r := m.Clear()
	actual := args.Map{
		"isNil": r == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_Dispose(t *testing.T) {
	tc := mapResultsDisposeTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.New("v"))

	// Act
	m.Dispose()
	actual := args.Map{
		"nilItems": m.Items == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_Json(t *testing.T) {
	tc := mapResultsJsonTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.New("v"))

	// Act
	r := m.Json()
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_GetPagesSize_Zero(t *testing.T) {
	tc := mapResultsGetPagesSizeZeroTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{
		"pages": m.GetPagesSize(0),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_AddMapResults_Nil(t *testing.T) {
	tc := mapResultsAddMapResultsNilTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	m.AddMapResults(nil)
	actual := args.Map{
		"length": m.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_AddMapAnyItems_Empty(t *testing.T) {
	tc := mapResultsAddMapAnyItemsEmptyTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	m.AddMapAnyItems(map[string]any{})
	actual := args.Map{
		"length": m.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_ResultCollection_Empty(t *testing.T) {
	tc := mapResultsResultCollectionEmptyTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{
		"isEmpty": m.ResultCollection().IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_AllResultsCollection_Empty(t *testing.T) {
	tc := mapResultsAllResultsCollectionEmptyTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{
		"isEmpty": m.AllResultsCollection().IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_GetStrings_Empty(t *testing.T) {
	tc := mapResultsGetStringsEmptyTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{
		"length": len(m.GetStrings()),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_AddJsoner_Nil(t *testing.T) {
	tc := mapResultsAddJsonerNilTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	m.AddJsoner("k", nil)
	actual := args.Map{
		"length": m.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_AddNonEmptyNonErrorPtr_Nil(t *testing.T) {
	tc := mapResultsAddNonEmptyNonErrorPtrNilTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	m.AddNonEmptyNonErrorPtr("k", nil)
	actual := args.Map{
		"length": m.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_GetNewMapUsingKeys_Empty(t *testing.T) {
	tc := mapResultsGetNewMapUsingKeysEmptyTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	r := m.GetNewMapUsingKeys(false)
	actual := args.Map{
		"isEmpty": r.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_AddMapResultsUsingCloneOption_Empty(t *testing.T) {
	tc := mapResultsAddMapResultsUsingCloneOptionEmptyTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	m.AddMapResultsUsingCloneOption(false, false, map[string]corejson.Result{})
	actual := args.Map{
		"length": m.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapResults_AddKeysWithJsoners_Nil(t *testing.T) {
	tc := mapResultsAddKeysWithJsonersNilTestCase

	// Arrange
	m := corejson.NewMapResults.Empty()

	// Act
	r := m.AddKeysWithJsoners(nil...)
	actual := args.Map{
		"isNil": r == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
