package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/namevalue"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage10 — namevalue final coverage gaps
// ══════════════════════════════════════════════════════════════════════════════

// --- AppendsIf ---

func Test_AppendsIf_True(t *testing.T) {
	// Arrange
	slice := []namevalue.StringAny{
		{Name: "a", Value: 1},
	}
	appending := namevalue.StringAny{Name: "b", Value: 2}

	// Act
	result := namevalue.AppendsIf(true, slice, appending)

	// Assert
	convey.Convey("AppendsIf true appends items", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}

func Test_AppendsIf_False(t *testing.T) {
	// Arrange
	slice := []namevalue.StringAny{
		{Name: "a", Value: 1},
	}

	// Act
	result := namevalue.AppendsIf(false, slice, namevalue.StringAny{Name: "b", Value: 2})

	// Assert
	convey.Convey("AppendsIf false returns original", t, func() {
		convey.So(len(result), convey.ShouldEqual, 1)
	})
}

func Test_AppendsIf_EmptyItems(t *testing.T) {
	// Arrange
	slice := []namevalue.StringAny{
		{Name: "a", Value: 1},
	}

	// Act
	result := namevalue.AppendsIf[string, any](true, slice)

	// Assert
	convey.Convey("AppendsIf true with no items returns original", t, func() {
		convey.So(len(result), convey.ShouldEqual, 1)
	})
}

// --- PrependsIf ---

func Test_PrependsIf_True(t *testing.T) {
	// Arrange
	slice := []namevalue.StringAny{
		{Name: "a", Value: 1},
	}
	prepending := namevalue.StringAny{Name: "b", Value: 2}

	// Act
	result := namevalue.PrependsIf(true, slice, prepending)

	// Assert
	convey.Convey("PrependsIf true prepends items", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
		convey.So(result[0].Name, convey.ShouldEqual, "b")
	})
}

func Test_PrependsIf_False(t *testing.T) {
	// Arrange
	slice := []namevalue.StringAny{
		{Name: "a", Value: 1},
	}

	// Act
	result := namevalue.PrependsIf(false, slice, namevalue.StringAny{Name: "b", Value: 2})

	// Assert
	convey.Convey("PrependsIf false returns original", t, func() {
		convey.So(len(result), convey.ShouldEqual, 1)
	})
}

func Test_PrependsIf_EmptyItems(t *testing.T) {
	// Arrange
	slice := []namevalue.StringAny{
		{Name: "a", Value: 1},
	}

	// Act
	result := namevalue.PrependsIf[string, any](true, slice)

	// Assert
	convey.Convey("PrependsIf true with no items returns original", t, func() {
		convey.So(len(result), convey.ShouldEqual, 1)
	})
}

// --- Collection methods ---

func Test_Collection_PrependUsingFuncIf_True(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})

	// Act
	col.PrependUsingFuncIf(true, func() []namevalue.StringAny {
		return []namevalue.StringAny{{Name: "b", Value: 2}}
	})

	// Assert
	convey.Convey("PrependUsingFuncIf true prepends from func", t, func() {
		convey.So(col.Length(), convey.ShouldEqual, 2)
	})
}

func Test_Collection_PrependUsingFuncIf_False(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})

	// Act
	col.PrependUsingFuncIf(false, func() []namevalue.StringAny {
		return []namevalue.StringAny{{Name: "b", Value: 2}}
	})

	// Assert
	convey.Convey("PrependUsingFuncIf false returns original", t, func() {
		convey.So(col.Length(), convey.ShouldEqual, 1)
	})
}

func Test_Collection_PrependUsingFuncIf_NilFunc(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})

	// Act
	col.PrependUsingFuncIf(true, nil)

	// Assert
	convey.Convey("PrependUsingFuncIf nil func returns original", t, func() {
		convey.So(col.Length(), convey.ShouldEqual, 1)
	})
}

func Test_Collection_AppendUsingFuncIf_True(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})

	// Act
	col.AppendUsingFuncIf(true, func() []namevalue.StringAny {
		return []namevalue.StringAny{{Name: "b", Value: 2}}
	})

	// Assert
	convey.Convey("AppendUsingFuncIf true appends from func", t, func() {
		convey.So(col.Length(), convey.ShouldEqual, 2)
	})
}

func Test_Collection_AppendUsingFuncIf_False(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})

	// Act
	col.AppendUsingFuncIf(false, func() []namevalue.StringAny {
		return []namevalue.StringAny{{Name: "b", Value: 2}}
	})

	// Assert
	convey.Convey("AppendUsingFuncIf false returns original", t, func() {
		convey.So(col.Length(), convey.ShouldEqual, 1)
	})
}

func Test_Collection_AppendUsingFuncIf_NilFunc(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})

	// Act
	col.AppendUsingFuncIf(true, nil)

	// Assert
	convey.Convey("AppendUsingFuncIf nil func returns original", t, func() {
		convey.So(col.Length(), convey.ShouldEqual, 1)
	})
}

func Test_Collection_AppendPrependIf_True(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})

	prepend := []namevalue.StringAny{{Name: "p", Value: 0}}
	append_ := []namevalue.StringAny{{Name: "z", Value: 9}}

	// Act
	col.AppendPrependIf(true, prepend, append_)

	// Assert
	convey.Convey("AppendPrependIf true prepends and appends", t, func() {
		convey.So(col.Length(), convey.ShouldEqual, 3)
	})
}

func Test_Collection_AppendPrependIf_False(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})

	// Act
	col.AppendPrependIf(false, nil, nil)

	// Assert
	convey.Convey("AppendPrependIf false returns original", t, func() {
		convey.So(col.Length(), convey.ShouldEqual, 1)
	})
}

func Test_Collection_AppendPrependIf_EmptySlices(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})

	// Act
	col.AppendPrependIf(true, []namevalue.StringAny{}, []namevalue.StringAny{})

	// Assert
	convey.Convey("AppendPrependIf with empty slices", t, func() {
		convey.So(col.Length(), convey.ShouldEqual, 1)
	})
}

func Test_Collection_AddsPtr(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	item := &namevalue.StringAny{Name: "a", Value: 1}

	// Act
	col.AddsPtr(item, nil, item)

	// Assert
	convey.Convey("AddsPtr adds non-nil items", t, func() {
		convey.So(col.Length(), convey.ShouldEqual, 2)
	})
}

func Test_Collection_AddsPtr_Empty(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()

	// Act
	col.AddsPtr()

	// Assert
	convey.Convey("AddsPtr with no items", t, func() {
		convey.So(col.Length(), convey.ShouldEqual, 0)
	})
}

func Test_Collection_CompiledLazyString(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})

	// Act
	str1 := col.CompiledLazyString()
	str2 := col.CompiledLazyString() // cached

	// Assert
	convey.Convey("CompiledLazyString caches", t, func() {
		convey.So(str1, convey.ShouldEqual, str2)
		convey.So(col.HasCompiledString(), convey.ShouldBeTrue)
	})
}

func Test_Collection_CompiledLazyString_Nil(t *testing.T) {
	// Arrange
	var col *namevalue.NameValuesCollection

	// Act
	result := col.CompiledLazyString()

	// Assert
	convey.Convey("CompiledLazyString nil returns empty", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

func Test_Collection_ConcatNew_AppendsIfFinalGaps(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})

	// Act
	result := col.ConcatNew(namevalue.StringAny{Name: "b", Value: 2})

	// Assert
	convey.Convey("ConcatNew creates new collection", t, func() {
		convey.So(result.Length(), convey.ShouldEqual, 2)
		convey.So(col.Length(), convey.ShouldEqual, 1)
	})
}

func Test_Collection_ConcatNewPtr_AppendsIfFinalGaps(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})
	item := &namevalue.StringAny{Name: "b", Value: 2}

	// Act
	result := col.ConcatNewPtr(item)

	// Assert
	convey.Convey("ConcatNewPtr creates new collection", t, func() {
		convey.So(result.Length(), convey.ShouldEqual, 2)
	})
}

func Test_Collection_IsEqualByString_DiffItems(t *testing.T) {
	// Arrange
	col1 := namevalue.NewCollection()
	col1.Add(namevalue.StringAny{Name: "a", Value: 1})

	col2 := namevalue.NewCollection()
	col2.Add(namevalue.StringAny{Name: "b", Value: 2})

	// Act
	result := col1.IsEqualByString(col2)

	// Assert
	convey.Convey("IsEqualByString returns false for diff items", t, func() {
		convey.So(result, convey.ShouldBeFalse)
	})
}

func Test_Collection_IsEqualByString_DiffLength(t *testing.T) {
	// Arrange
	col1 := namevalue.NewCollection()
	col1.Add(namevalue.StringAny{Name: "a", Value: 1})

	col2 := namevalue.NewCollection()

	// Act
	result := col1.IsEqualByString(col2)

	// Assert
	convey.Convey("IsEqualByString returns false for diff length", t, func() {
		convey.So(result, convey.ShouldBeFalse)
	})
}

func Test_Collection_IsEqualByString_BothNil(t *testing.T) {
	// Arrange
	var col1 *namevalue.NameValuesCollection
	var col2 *namevalue.NameValuesCollection

	// Act
	result := col1.IsEqualByString(col2)

	// Assert
	convey.Convey("IsEqualByString both nil returns true", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}

func Test_Collection_IsEqualByString_OneNil(t *testing.T) {
	// Arrange
	col1 := namevalue.NewCollection()
	var col2 *namevalue.NameValuesCollection

	// Act
	result := col1.IsEqualByString(col2)

	// Assert
	convey.Convey("IsEqualByString one nil returns false", t, func() {
		convey.So(result, convey.ShouldBeFalse)
	})
}

func Test_Collection_JsonStrings(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})

	// Act
	result := col.JsonStrings()

	// Assert
	convey.Convey("JsonStrings returns JSON per item", t, func() {
		convey.So(len(result), convey.ShouldEqual, 1)
	})
}

func Test_Collection_JoinJsonStrings(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})
	col.Add(namevalue.StringAny{Name: "b", Value: 2})

	// Act
	result := col.JoinJsonStrings(",")

	// Assert
	convey.Convey("JoinJsonStrings joins with separator", t, func() {
		convey.So(result, convey.ShouldContainSubstring, ",")
	})
}

func Test_Collection_JoinCsv(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})

	// Act
	result := col.JoinCsv()

	// Assert
	convey.Convey("JoinCsv returns CSV", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Collection_JoinCsvLine(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})

	// Act
	result := col.JoinCsvLine()

	// Assert
	convey.Convey("JoinCsvLine returns CSV line", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Collection_CsvStrings(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})

	// Act
	result := col.CsvStrings()

	// Assert
	convey.Convey("CsvStrings returns quoted strings", t, func() {
		convey.So(len(result), convey.ShouldEqual, 1)
	})
}

func Test_Collection_CsvStrings_Empty(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()

	// Act
	result := col.CsvStrings()

	// Assert
	convey.Convey("CsvStrings empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

func Test_Collection_Error(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: "fail"})

	// Act
	err := col.Error()

	// Assert
	convey.Convey("Error returns error with content", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_Collection_Error_Empty(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()

	// Act
	err := col.Error()

	// Assert
	convey.Convey("Error empty returns nil", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_Collection_ErrorUsingMessage_AppendsIfFinalGaps(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: "fail"})

	// Act
	err := col.ErrorUsingMessage("prefix:")

	// Assert
	convey.Convey("ErrorUsingMessage returns error with prefix", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
		convey.So(err.Error(), convey.ShouldStartWith, "prefix:")
	})
}

func Test_Collection_ErrorUsingMessage_Empty(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()

	// Act
	err := col.ErrorUsingMessage("prefix:")

	// Assert
	convey.Convey("ErrorUsingMessage empty returns nil", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_Collection_String_Cached(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})
	_ = col.CompiledLazyString() // cache it

	// Act
	result := col.String() // should use cached

	// Assert
	convey.Convey("String uses cached value", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Collection_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var col *namevalue.NameValuesCollection

	// Act
	result := col.ClonePtr()

	// Assert
	convey.Convey("ClonePtr nil returns nil", t, func() {
		convey.So(result, convey.ShouldBeNil)
	})
}

func Test_Collection_HasIndex(t *testing.T) {
	// Arrange
	col := namevalue.NewCollection()
	col.Add(namevalue.StringAny{Name: "a", Value: 1})

	// Act & Assert
	convey.Convey("HasIndex works correctly", t, func() {
		convey.So(col.HasIndex(0), convey.ShouldBeTrue)
		convey.So(col.HasIndex(1), convey.ShouldBeFalse)
	})
}

// --- Instance methods ---

func Test_Instance_JsonString(t *testing.T) {
	// Arrange
	inst := namevalue.StringAny{Name: "key", Value: "val"}

	// Act
	result := inst.JsonString()

	// Assert
	convey.Convey("Instance.JsonString returns JSON", t, func() {
		convey.So(result, convey.ShouldContainSubstring, "key")
	})
}

func Test_Instance_JsonString_Nil(t *testing.T) {
	// Arrange
	var inst *namevalue.StringAny

	// Act
	result := inst.JsonString()

	// Assert
	convey.Convey("Instance.JsonString nil returns empty", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

func Test_Instance_Dispose_FromAppendsIfTrueFinalGa(t *testing.T) {
	// Arrange
	inst := &namevalue.StringAny{Name: "key", Value: "val"}

	// Act
	inst.Dispose()

	// Assert
	convey.Convey("Instance.Dispose clears fields", t, func() {
		convey.So(inst.Name, convey.ShouldBeEmpty)
		convey.So(inst.Value, convey.ShouldBeNil)
	})
}

func Test_Instance_Dispose_Nil_FromAppendsIfTrueFinalGa(t *testing.T) {
	// Arrange
	var inst *namevalue.StringAny

	// Act & Assert
	convey.Convey("Instance.Dispose nil is safe", t, func() {
		convey.So(func() { inst.Dispose() }, convey.ShouldNotPanic)
	})
}

// --- NameValuesCollection factory methods ---

func Test_NewNameValuesCollection(t *testing.T) {
	// Act
	col := namevalue.NewNameValuesCollection(10)

	// Assert
	convey.Convey("NewNameValuesCollection creates empty with capacity", t, func() {
		convey.So(col.IsEmpty(), convey.ShouldBeTrue)
	})
}

func Test_NewCollection(t *testing.T) {
	// Act
	col := namevalue.NewCollection()

	// Assert
	convey.Convey("NewCollection creates empty", t, func() {
		convey.So(col.IsEmpty(), convey.ShouldBeTrue)
	})
}

func Test_NewNewNameValuesCollectionUsing_Clone(t *testing.T) {
	// Arrange
	items := []namevalue.StringAny{
		{Name: "a", Value: 1},
	}

	// Act
	col := namevalue.NewNewNameValuesCollectionUsing(true, items...)

	// Assert
	convey.Convey("NewNewNameValuesCollectionUsing clone creates new", t, func() {
		convey.So(col.Length(), convey.ShouldEqual, 1)
	})
}

func Test_NewNewNameValuesCollectionUsing_NoClone(t *testing.T) {
	// Arrange
	items := []namevalue.StringAny{
		{Name: "a", Value: 1},
	}

	// Act
	col := namevalue.NewNewNameValuesCollectionUsing(false, items...)

	// Assert
	convey.Convey("NewNewNameValuesCollectionUsing no clone", t, func() {
		convey.So(col.Length(), convey.ShouldEqual, 1)
	})
}

func Test_NewNewNameValuesCollectionUsing_Nil(t *testing.T) {
	// Act
	col := namevalue.NewNewNameValuesCollectionUsing(false)

	// Assert
	convey.Convey("NewNewNameValuesCollectionUsing nil returns empty", t, func() {
		convey.So(col.IsEmpty(), convey.ShouldBeTrue)
	})
}

func Test_EmptyNameValuesCollection(t *testing.T) {
	// Act
	col := namevalue.EmptyNameValuesCollection()

	// Assert
	convey.Convey("EmptyNameValuesCollection creates empty", t, func() {
		convey.So(col.IsEmpty(), convey.ShouldBeTrue)
	})
}

// --- Collection generic factory methods ---

func Test_NewGenericCollectionDefault(t *testing.T) {
	// Act
	col := namevalue.NewGenericCollectionDefault[string, int]()

	// Assert
	convey.Convey("NewGenericCollectionDefault creates empty", t, func() {
		convey.So(col.IsEmpty(), convey.ShouldBeTrue)
	})
}

func Test_EmptyGenericCollection(t *testing.T) {
	// Act
	col := namevalue.EmptyGenericCollection[string, int]()

	// Assert
	convey.Convey("EmptyGenericCollection creates empty", t, func() {
		convey.So(col.IsEmpty(), convey.ShouldBeTrue)
	})
}
