package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// BytesCollection — Length / IsEmpty / HasAnyItem
// =============================================================================

func Test_BytesColl_NilLength(t *testing.T) {
	tc := bytesCollNilLengthTestCase

	// Arrange
	var c *corejson.BytesCollection

	// Act
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_IsEmpty(t *testing.T) {
	tc := bytesCollIsEmptyTrueTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{
		"isEmpty": c.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_HasAnyItem(t *testing.T) {
	tc := bytesCollHasAnyItemTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`))

	// Act
	actual := args.Map{
		"hasAny": c.HasAnyItem(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_LastIndex(t *testing.T) {
	tc := bytesCollLastIndexTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`))

	// Act
	actual := args.Map{
		"lastIndex": c.LastIndex(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// BytesCollection — FirstOrDefault / LastOrDefault
// =============================================================================

func Test_BytesColl_FirstOrDefault_Empty(t *testing.T) {
	tc := bytesCollFirstOrDefaultEmptyTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{
		"isNil": c.FirstOrDefault() == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_FirstOrDefault_HasItem(t *testing.T) {
	tc := bytesCollFirstOrDefaultHasItemTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"first"`))

	// Act
	actual := args.Map{
		"isNil": c.FirstOrDefault() == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_LastOrDefault_Empty(t *testing.T) {
	tc := bytesCollLastOrDefaultEmptyTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{
		"isNil": c.LastOrDefault() == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_LastOrDefault_HasItem(t *testing.T) {
	tc := bytesCollLastOrDefaultHasItemTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`))

	// Act
	actual := args.Map{
		"isNil": c.LastOrDefault() == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// BytesCollection — Take / Limit / Skip
// =============================================================================

func Test_BytesColl_Take_Empty(t *testing.T) {
	tc := bytesCollTakeEmptyTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	r := c.Take(5)
	actual := args.Map{
		"isEmpty": r.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_Take_Valid(t *testing.T) {
	tc := bytesCollTakeValidTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`)).Add([]byte(`"c"`))

	// Act
	r := c.Take(2)
	actual := args.Map{
		"length": r.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_Limit_TakeAll(t *testing.T) {
	tc := bytesCollLimitTakeAllTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`))

	// Act
	r := c.Limit(-1)
	actual := args.Map{
		"length": r.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_Limit_Empty(t *testing.T) {
	tc := bytesCollLimitEmptyTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	r := c.Limit(5)
	actual := args.Map{
		"isEmpty": r.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_Skip_Empty(t *testing.T) {
	tc := bytesCollSkipEmptyTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	r := c.Skip(1)
	actual := args.Map{
		"isEmpty": r.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_Skip_Valid(t *testing.T) {
	tc := bytesCollSkipValidTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`)).Add([]byte(`"c"`))

	// Act
	r := c.Skip(1)
	actual := args.Map{
		"length": r.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// BytesCollection — Add variants
// =============================================================================

func Test_BytesColl_AddSkipOnNil_Nil(t *testing.T) {
	tc := bytesCollAddSkipOnNilNilTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	c.AddSkipOnNil(nil)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_AddSkipOnNil_Valid(t *testing.T) {
	tc := bytesCollAddSkipOnNilValidTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	c.AddSkipOnNil([]byte(`"hi"`))
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_AddNonEmpty_Empty(t *testing.T) {
	tc := bytesCollAddNonEmptyEmptyTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	c.AddNonEmpty([]byte{})
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_AddNonEmpty_Valid(t *testing.T) {
	tc := bytesCollAddNonEmptyValidTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	c.AddNonEmpty([]byte(`"hi"`))
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_AddResultPtr_Skip(t *testing.T) {
	tc := bytesCollAddResultPtrSkipTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	r := &corejson.Result{}

	// Act
	c.AddResultPtr(r)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_AddResultPtr_Valid(t *testing.T) {
	tc := bytesCollAddResultPtrValidTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	r := corejson.NewPtr("hi")

	// Act
	c.AddResultPtr(r)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_AddResult_Skip(t *testing.T) {
	tc := bytesCollAddResultSkipTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	r := corejson.Result{}

	// Act
	c.AddResult(r)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_AddResult_Valid(t *testing.T) {
	tc := bytesCollAddResultValidTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	r := corejson.New("hi")

	// Act
	c.AddResult(r)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_AddPtr_Empty(t *testing.T) {
	tc := bytesCollAddPtrEmptyTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	c.AddPtr([]byte{})
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_AddPtr_Valid(t *testing.T) {
	tc := bytesCollAddPtrValidTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	c.AddPtr([]byte(`"hi"`))
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_Adds_Empty(t *testing.T) {
	tc := bytesCollAddsEmptyTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	c.Adds()
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_Adds_Valid(t *testing.T) {
	tc := bytesCollAddsValidTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	c.Adds([]byte(`"a"`), []byte(`"b"`))
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_AddsPtr_Nil(t *testing.T) {
	tc := bytesCollAddsPtrNilTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	c.AddsPtr(nil)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// BytesCollection — AddAny / AddAnyItems / AddSerializerFunc
// =============================================================================

func Test_BytesColl_AddAny_Valid(t *testing.T) {
	tc := bytesCollAddAnyValidTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	err := c.AddAny("hello")
	actual := args.Map{
		"hasError": err != nil,
		"length":   c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_AddAnyItems_Empty(t *testing.T) {
	tc := bytesCollAddAnyItemsEmptyTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	err := c.AddAnyItems()
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_AddAnyItems_Valid(t *testing.T) {
	tc := bytesCollAddAnyItemsValidTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	err := c.AddAnyItems("a", "b")
	actual := args.Map{
		"hasError": err != nil,
		"length":   c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_AddSerializerFunc_Nil(t *testing.T) {
	tc := bytesCollAddSerializerFuncNilTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	c.AddSerializerFunc(nil)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_AddSerializerFunctions_Empty(t *testing.T) {
	tc := bytesCollAddSerializerFunctionsEmptyTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	c.AddSerializerFunctions()
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// BytesCollection — GetAtSafe / GetResultAtSafe / GetAtSafeUsingLength
// =============================================================================

func Test_BytesColl_GetAtSafe_OutOfRange(t *testing.T) {
	tc := bytesCollGetAtSafeOutOfRangeTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{
		"isNil": c.GetAtSafe(5) == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_GetAtSafe_Valid(t *testing.T) {
	tc := bytesCollGetAtSafeValidTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"hi"`))

	// Act
	actual := args.Map{
		"isNil": c.GetAtSafe(0) == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_GetResultAtSafe_OutOfRange(t *testing.T) {
	tc := bytesCollGetResultAtSafeOutOfRangeTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{
		"isNil": c.GetResultAtSafe(5) == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_GetAtSafeUsingLength_Out(t *testing.T) {
	tc := bytesCollGetAtSafeUsingLengthOutTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{
		"isNil": c.GetAtSafeUsingLength(5, 0) == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_GetAtSafePtr_OutOfRange(t *testing.T) {
	tc := bytesCollGetAtSafePtrOutOfRangeTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{
		"isNil": c.GetAtSafePtr(5) == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// BytesCollection — Clear / Dispose / Clone / ClonePtr
// =============================================================================

func Test_BytesColl_Clear(t *testing.T) {
	tc := bytesCollClearTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`))

	// Act
	c.Clear()
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_Clear_Nil(t *testing.T) {
	tc := bytesCollClearNilTestCase

	// Arrange
	var c *corejson.BytesCollection

	// Act
	r := c.Clear()
	actual := args.Map{
		"isNil": r == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_Dispose(t *testing.T) {
	tc := bytesCollDisposeTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`))

	// Act
	c.Dispose()
	actual := args.Map{
		"nilItems": c.Items == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_Clone(t *testing.T) {
	tc := bytesCollCloneTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`))

	// Act
	cloned := c.Clone(false)
	actual := args.Map{
		"length": cloned.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_ClonePtr_Nil(t *testing.T) {
	tc := bytesCollClonePtrNilTestCase

	// Arrange
	var c *corejson.BytesCollection

	// Act
	r := c.ClonePtr(false)
	actual := args.Map{
		"isNil": r == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_ClonePtr_Valid(t *testing.T) {
	tc := bytesCollClonePtrValidTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`))

	// Act
	cloned := c.ClonePtr(true)
	actual := args.Map{
		"length": cloned.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// BytesCollection — Strings / Json / Paging / UnmarshalAt
// =============================================================================

func Test_BytesColl_Strings_Empty(t *testing.T) {
	tc := bytesCollStringsEmptyTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{
		"length": len(c.Strings()),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_Strings_Valid(t *testing.T) {
	tc := bytesCollStringsValidTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`))

	// Act
	actual := args.Map{
		"length": len(c.Strings()),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_Json(t *testing.T) {
	tc := bytesCollJsonTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`))

	// Act
	r := c.Json()
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_GetPagesSize_Zero(t *testing.T) {
	tc := bytesCollGetPagesSizeZeroTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{
		"pages": c.GetPagesSize(0),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_GetPagesSize_Valid(t *testing.T) {
	tc := bytesCollGetPagesSizeValidTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`)).Add([]byte(`"b"`)).Add([]byte(`"c"`))

	// Act
	actual := args.Map{
		"pages": c.GetPagesSize(2),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_UnmarshalAt(t *testing.T) {
	tc := bytesCollUnmarshalAtTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"hello"`))
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

func Test_BytesColl_GetAt(t *testing.T) {
	tc := bytesCollGetAtTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`))

	// Act
	actual := args.Map{
		"isNil": c.GetAt(0) == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_JsonResultAt(t *testing.T) {
	tc := bytesCollJsonResultAtTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`))

	// Act
	actual := args.Map{
		"isNil": c.JsonResultAt(0) == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// BytesCollection — AddBytesCollection / AddRawMapResults
// =============================================================================

func Test_BytesColl_AddBytesCollection_Empty(t *testing.T) {
	tc := bytesCollAddBytesCollectionEmptyTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	src := corejson.NewBytesCollection.Empty()

	// Act
	c.AddBytesCollection(src)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_AddBytesCollection_Valid(t *testing.T) {
	tc := bytesCollAddBytesCollectionValidTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	src := corejson.NewBytesCollection.Empty()
	src.Add([]byte(`"a"`))

	// Act
	c.AddBytesCollection(src)
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_AddRawMapResults_Empty(t *testing.T) {
	tc := bytesCollAddRawMapResultsEmptyTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	c.AddRawMapResults(map[string]corejson.Result{})
	actual := args.Map{
		"length": c.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// BytesCollection — ShadowClone / ParseInjectUsingJson / UnmarshalIntoSameIndex
// =============================================================================

func Test_BytesColl_ShadowClone(t *testing.T) {
	tc := bytesCollShadowCloneTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`))

	// Act
	cloned := c.ShadowClone()
	actual := args.Map{
		"length": cloned.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_ParseInjectUsingJson_Valid(t *testing.T) {
	tc := bytesCollParseInjectUsingJsonValidTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"a"`))
	jr := corejson.NewPtr(c)

	newC := corejson.NewBytesCollection.Empty()

	// Act
	_, err := newC.ParseInjectUsingJson(jr)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesColl_UnmarshalIntoSameIndex_Nil(t *testing.T) {
	tc := bytesCollUnmarshalIntoSameIndexNilTestCase

	// Arrange
	c := corejson.NewBytesCollection.Empty()

	// Act
	_, hasErr := c.UnmarshalIntoSameIndex(nil...)
	actual := args.Map{
		"hasAnyErr": hasErr,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
