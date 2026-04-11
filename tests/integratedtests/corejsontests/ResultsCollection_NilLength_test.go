package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// ResultsCollection — Length / IsEmpty / HasAnyItem
// =============================================================================

func Test_ResultsCollection_NilLength_FromResultsCollectionNil(t *testing.T) {
	tc := resultsCollectionNilLengthTestCase

	// Arrange
	var c *corejson.ResultsCollection

	// Act
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_IsEmpty(t *testing.T) {
	tc := resultsCollectionIsEmptyTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()

	// Act
	actual := args.Map{
		"isEmpty": c.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_HasAnyItem(t *testing.T) {
	tc := resultsCollectionHasAnyItemTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))

	// Act
	actual := args.Map{
		"hasAny": c.HasAnyItem(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// ResultsCollection — FirstOrDefault / LastOrDefault
// =============================================================================

func Test_ResultsCollection_FirstOrDefault_Empty_FromResultsCollectionNil(t *testing.T) {
	tc := resultsCollectionFirstOrDefaultEmptyTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()

	// Act
	actual := args.Map{
		"isNil": c.FirstOrDefault() == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_FirstOrDefault_HasItem(t *testing.T) {
	tc := resultsCollectionFirstOrDefaultHasItemTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("first"))

	// Act
	actual := args.Map{
		"isNil": c.FirstOrDefault() == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_LastOrDefault_Empty_FromResultsCollectionNil(t *testing.T) {
	tc := resultsCollectionLastOrDefaultEmptyTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()

	// Act
	actual := args.Map{
		"isNil": c.LastOrDefault() == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_LastOrDefault_HasItem(t *testing.T) {
	tc := resultsCollectionLastOrDefaultHasItemTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("last"))

	// Act
	actual := args.Map{
		"isNil": c.LastOrDefault() == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// ResultsCollection — Take / Limit / Skip
// =============================================================================

func Test_ResultsCollection_Take_Empty(t *testing.T) {
	tc := resultsCollectionTakeEmptyTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()

	// Act
	r := c.Take(5)
	actual := args.Map{
		"isEmpty": r.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_Take_Valid(t *testing.T) {
	tc := resultsCollectionTakeValidTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a")).Add(corejson.New("b")).Add(corejson.New("c"))

	// Act
	r := c.Take(2)
	actual := args.Map{
		"length": r.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_Limit_TakeAll(t *testing.T) {
	tc := resultsCollectionLimitTakeAllTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a")).Add(corejson.New("b"))

	// Act
	r := c.Limit(-1)
	actual := args.Map{
		"length": r.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_Skip_Empty(t *testing.T) {
	tc := resultsCollectionSkipEmptyTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()

	// Act
	r := c.Skip(1)
	actual := args.Map{
		"isEmpty": r.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_Skip_Valid(t *testing.T) {
	tc := resultsCollectionSkipValidTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a")).Add(corejson.New("b")).Add(corejson.New("c"))

	// Act
	r := c.Skip(1)
	actual := args.Map{
		"length": r.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// ResultsCollection — Add variants
// =============================================================================

func Test_ResultsCollection_AddSkipOnNil_Nil(t *testing.T) {
	tc := resultsCollectionAddSkipOnNilNilTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()

	// Act
	c.AddSkipOnNil(nil)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_AddSkipOnNil_Valid(t *testing.T) {
	tc := resultsCollectionAddSkipOnNilValidTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	r := corejson.NewPtr("hi")

	// Act
	c.AddSkipOnNil(r)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_AddNonNilNonError_Nil(t *testing.T) {
	tc := resultsCollectionAddNonNilNonErrorNilTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()

	// Act
	c.AddNonNilNonError(nil)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_AddNonNilNonError_Error(t *testing.T) {
	tc := resultsCollectionAddNonNilNonErrorErrTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	r := &corejson.Result{Error: errors.New("fail")}

	// Act
	c.AddNonNilNonError(r)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_AddNonNilNonError_Valid(t *testing.T) {
	tc := resultsCollectionAddNonNilNonErrorValidTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	r := corejson.NewPtr("hi")

	// Act
	c.AddNonNilNonError(r)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// ResultsCollection — HasError / AllErrors
// =============================================================================

func Test_ResultsCollection_HasError_False(t *testing.T) {
	tc := resultsCollectionHasErrorFalseTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))

	// Act
	actual := args.Map{
		"hasError": c.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_HasError_True(t *testing.T) {
	tc := resultsCollectionHasErrorTrueTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.Result{Error: errors.New("fail")})

	// Act
	actual := args.Map{
		"hasError": c.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_AllErrors_Empty(t *testing.T) {
	tc := resultsCollectionAllErrorsEmptyTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()

	// Act
	errList, hasErr := c.AllErrors()
	actual := args.Map{
		"errorCount": len(errList),
		"hasAnyErr":  hasErr,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_AllErrors_WithErr(t *testing.T) {
	tc := resultsCollectionAllErrorsWithErrTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.Result{Error: errors.New("fail")})
	c.Add(corejson.New("ok"))

	// Act
	errList, hasErr := c.AllErrors()
	actual := args.Map{
		"errorCount": len(errList),
		"hasAnyErr":  hasErr,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// ResultsCollection — GetAtSafe / UnmarshalAt / Paging
// =============================================================================

func Test_ResultsCollection_GetAtSafe_OutOfRange(t *testing.T) {
	tc := resultsCollectionGetAtSafeOutOfRangeTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()

	// Act
	actual := args.Map{
		"isNil": c.GetAtSafe(5) == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_UnmarshalAt_FromResultsCollectionNil(t *testing.T) {
	tc := resultsCollectionUnmarshalAtTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("hello"))
	var s string

	// Act
	err := c.UnmarshalAt(0, &s)
	actual := args.Map{
		"hasError": err != nil,
		"result":   s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_GetPagesSize_Zero(t *testing.T) {
	tc := resultsCollectionGetPagesSizeZeroTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()

	// Act
	actual := args.Map{
		"pages": c.GetPagesSize(0),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_GetPagesSize_Valid(t *testing.T) {
	tc := resultsCollectionGetPagesSizeValidTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a")).Add(corejson.New("b")).Add(corejson.New("c"))

	// Act
	actual := args.Map{
		"pages": c.GetPagesSize(2),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// ResultsCollection — Clear / Dispose / Clone / Json
// =============================================================================

func Test_ResultsCollection_Clear(t *testing.T) {
	tc := resultsCollectionClearTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))

	// Act
	c.Clear()
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_Clear_Nil(t *testing.T) {
	tc := resultsCollectionClearNilTestCase

	// Arrange
	var c *corejson.ResultsCollection

	// Act
	r := c.Clear()
	actual := args.Map{
		"isNil": r == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_Dispose_FromResultsCollectionNil(t *testing.T) {
	tc := resultsCollectionDisposeTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))

	// Act
	c.Dispose()
	actual := args.Map{
		"nilItems": c.Items == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_Clone_Deep(t *testing.T) {
	tc := resultsCollectionCloneTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))

	// Act — Clone correctly copies items (bug was fixed: checks it.Length() not newResults.Length())
	cloned := c.Clone(true)
	actual := args.Map{
		"length": cloned.Length(),
	}

	// Assert — expect 1 (clone preserves items)
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_ClonePtr_Nil_FromResultsCollectionNil(t *testing.T) {
	tc := resultsCollectionClonePtrNilTestCase

	// Arrange
	var c *corejson.ResultsCollection

	// Act
	r := c.ClonePtr(false)
	actual := args.Map{
		"isNil": r == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultsCollection_Json(t *testing.T) {
	tc := resultsCollectionJsonTestCase

	// Arrange
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))

	// Act
	r := c.Json()
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// newResultsCollectionCreator
// =============================================================================

func Test_NewResultsCollection_DeserializeUsingBytes_Invalid(t *testing.T) {
	tc := newResultsCollectionDeserializeInvalidTestCase

	// Arrange
	// (invalid bytes)

	// Act
	_, err := corejson.NewResultsCollection.DeserializeUsingBytes([]byte(`bad`))
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_NewResultsCollection_DeserializeUsingResult_Error(t *testing.T) {
	tc := newResultsCollectionDeserializeResultErrorTestCase

	// Arrange
	jr := &corejson.Result{Error: errors.New("fail")}

	// Act
	_, err := corejson.NewResultsCollection.DeserializeUsingResult(jr)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
