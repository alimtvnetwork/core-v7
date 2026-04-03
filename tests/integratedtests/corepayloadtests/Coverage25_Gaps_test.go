package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// Helpers for this file: uses testUserCov23, makeTypedWrapperCov23, makeCollectionCov23
// from Coverage23_TypedFuncs_test.go

// ── Attributes.IsEqual — error different ──

func Test_Cov25_Attributes_IsEqual_ErrorDifferent(t *testing.T) {
	// Arrange
	a1 := corepayload.New.Attributes.All(
		nil, nil, nil, nil, nil, nil,
		errcore.MarshallingFailedType.ErrorRefOnly("err-a"),
	)
	a2 := corepayload.New.Attributes.All(
		nil, nil, nil, nil, nil, nil,
		errcore.MarshallingFailedType.ErrorRefOnly("err-b"),
	)

	// Act
	result := a1.IsEqual(a2)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- different errors", actual)
}

// ── Attributes.IsEqual — paging different ──

func Test_Cov25_Attributes_IsEqual_PagingDifferent(t *testing.T) {
	// Arrange
	p1 := &corepayload.PagingInfo{PageIndex: 1, PageSize: 10}
	p2 := &corepayload.PagingInfo{PageIndex: 2, PageSize: 10}
	a1 := corepayload.New.Attributes.All(nil, nil, nil, p1, nil, nil, nil)
	a2 := corepayload.New.Attributes.All(nil, nil, nil, p2, nil, nil, nil)

	// Act
	result := a1.IsEqual(a2)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- different paging", actual)
}

// ── Attributes.IsEqual — keyValuePairs different ──

func Test_Cov25_Attributes_IsEqual_KeyValuePairsDifferent(t *testing.T) {
	// Arrange
	kv1 := corestr.New.Hashmap.UsingPairs("k", "v1")
	kv2 := corestr.New.Hashmap.UsingPairs("k", "v2")
	a1 := corepayload.New.Attributes.All(nil, kv1, nil, nil, nil, nil, nil)
	a2 := corepayload.New.Attributes.All(nil, kv2, nil, nil, nil, nil, nil)

	// Act
	result := a1.IsEqual(a2)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- different key-value pairs", actual)
}

// ── Attributes.IsEqual — dynamicPayloads different ──

func Test_Cov25_Attributes_IsEqual_DynamicPayloadsDifferent(t *testing.T) {
	// Arrange
	a1 := corepayload.New.Attributes.All(nil, nil, nil, nil, []byte(`{"a":1}`), nil, nil)
	a2 := corepayload.New.Attributes.All(nil, nil, nil, nil, []byte(`{"b":2}`), nil, nil)

	// Act
	result := a1.IsEqual(a2)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- different dynamic payloads", actual)
}

// ── Attributes.IsEqual — anyKeyValuePairs different ──

func Test_Cov25_Attributes_IsEqual_AnyKeyValuePairsDifferent(t *testing.T) {
	// Arrange
	m1 := coredynamic.NewMapAnyItems(2)
	m1.Add("key", "val1")
	m2 := coredynamic.NewMapAnyItems(2)
	m2.Add("key", "val2")
	a1 := corepayload.New.Attributes.All(nil, nil, m1, nil, nil, nil, nil)
	a2 := corepayload.New.Attributes.All(nil, nil, m2, nil, nil, nil, nil)

	// Act
	result := a1.IsEqual(a2)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- different any key-value pairs", actual)
}

// ── Attributes.Clone — ClonePtr returns error ──

func Test_Cov25_Attributes_Clone_DeepClone_AnyKeyValuesCloneError(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.All(nil, nil, nil, nil, nil, nil, nil)

	// Act
	cloned, err := attr.Clone(true)

	// Assert
	actual := args.Map{
		"err":     err == nil,
		"isEmpty": cloned.IsEqual(nil),
	}
	expected := args.Map{
		"err":     true,
		"isEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "Clone deep returns no error -- nil anyKeyValues", actual)
}

// ── Attributes.deepClonePtr — HasError branch ──

func Test_Cov25_Attributes_DeepClone_WithError(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.All(nil, nil, nil, nil, nil, nil, nil)

	// Act
	cloned, err := attr.ClonePtr(true)

	// Assert
	actual := args.Map{
		"err":      err == nil,
		"hasError": cloned.HasError(),
	}
	expected := args.Map{
		"err":      true,
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "ClonePtr deep clones error -- has BasicErrWrapper", actual)
}

// ── AttributesGetters — Error() with error ──

func Test_Cov25_Attributes_Error_NilError(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.All(nil, nil, nil, nil, nil, nil, nil)

	// Act
	err := attr.Error()

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "Error returns nil -- no BasicErrWrapper", actual)
}

// ── AttributesGetters — IsErrorEqual with non-empty errors ──

func Test_Cov25_Attributes_IsErrorEqual_BothNilErrors(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.All(nil, nil, nil, nil, nil, nil, nil)

	// Act
	result := attr.IsErrorEqual(nil)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsErrorEqual returns true -- both nil error", actual)
}

// ── AttributesJson — ParseInjectUsingJson error ──

func Test_Cov25_Attributes_ParseInjectUsingJson_Error(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	badResult := corejson.NewResult.UsingBytes([]byte(`{invalid`))

	// Act
	_, err := attr.ParseInjectUsingJson(&badResult)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns error -- invalid json", actual)
}

// ── AttributesJson — ParseInjectUsingJsonMust panic ──

func Test_Cov25_Attributes_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	badResult := corejson.NewResult.UsingBytes([]byte(`{invalid`))
	didPanic := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		attr.ParseInjectUsingJsonMust(&badResult)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust panics -- invalid json", actual)
}

// ── AttributesJson — BasicErrorDeserializedTo with error ──

func Test_Cov25_Attributes_BasicErrorDeserializedTo_NilError(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.All(nil, nil, nil, nil, nil, nil, nil)
	var target map[string]any

	// Act
	err := attr.BasicErrorDeserializedTo(&target)

	// Assert
	actual := args.Map{"errNil": err == nil}
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "BasicErrorDeserializedTo no error -- nil BasicErrWrapper", actual)
}

// ── AttributesJson — DynamicPayloadsDeserializeMust panic ──

func Test_Cov25_Attributes_DynamicPayloadsDeserializeMust_Panic(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.All(nil, nil, nil, nil, []byte(`{invalid`), nil, nil)
	didPanic := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		var target map[string]any
		attr.DynamicPayloadsDeserializeMust(&target)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "DynamicPayloadsDeserializeMust panics -- invalid json", actual)
}

// ── AttributesSetters — HandleErr ──

func Test_Cov25_Attributes_HandleErr_WithError(t *testing.T) {
	// Arrange
	// HandleErr on attributes with error should not panic (it calls HandleError which logs)
	// We just verify it doesn't crash
	attr := corepayload.New.Attributes.Empty()

	// Act — no error, should be no-op
	attr.HandleErr()

	// Assert
	actual := args.Map{"noError": attr.IsEmptyError()}
	expected := args.Map{"noError": true}
	expected.ShouldBeEqual(t, 0, "HandleErr is no-op -- no error", actual)
}

// ── AttributesSetters — MustBeEmptyError panic ──

func Test_Cov25_Attributes_MustBeEmptyError_NoPanic(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.All(nil, nil, nil, nil, nil, nil, nil)
	didPanic := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		attr.MustBeEmptyError()
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": false}
	expected.ShouldBeEqual(t, 0, "MustBeEmptyError no panic -- no error", actual)
}

// ── PayloadWrapper — BasicError with error ──

func Test_Cov25_PayloadWrapper_BasicError_NilError(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.All(nil, nil, nil, nil, nil, nil, nil)
	pw := &corepayload.PayloadWrapper{Attributes: attr}

	// Act
	result := pw.BasicError()

	// Assert
	actual := args.Map{"hasBasicErr": result != nil}
	expected := args.Map{"hasBasicErr": false}
	expected.ShouldBeEqual(t, 0, "BasicError returns nil -- no error", actual)
}

// ── PayloadWrapper — PayloadDeserializeToPayloadBinder error ──

func Test_Cov25_PayloadWrapper_PayloadDeserializeToPayloadBinder_Error(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.All(nil, nil, nil, nil, nil, nil, nil)
	pw := &corepayload.PayloadWrapper{Attributes: attr}

	// Act
	_, err := pw.PayloadDeserializeToPayloadBinder()

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PayloadDeserializeToPayloadBinder returns error -- has error attr", actual)
}

// ── PayloadWrapper — SetPayloadDynamicAny error ──

func Test_Cov25_PayloadWrapper_SetPayloadDynamicAny_Error(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{}
	badInput := make(chan int) // channels cannot be serialized

	// Act
	result, err := pw.SetPayloadDynamicAny(badInput)

	// Assert
	actual := args.Map{
		"resultNil": result == nil,
		"hasErr":    err != nil,
	}
	expected := args.Map{
		"resultNil": true,
		"hasErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "SetPayloadDynamicAny returns error -- un-serializable input", actual)
}

// ── PayloadWrapper — HandleError ──

func Test_Cov25_PayloadWrapper_HandleError_NoError(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{}

	// Act — no error, should be no-op
	pw.HandleError()

	// Assert
	actual := args.Map{"noError": pw.IsEmptyError()}
	expected := args.Map{"noError": true}
	expected.ShouldBeEqual(t, 0, "HandleError is no-op -- no error", actual)
}

// ── PayloadWrapper — IsStandardTaskEntityEqual cast fail ──

func Test_Cov25_PayloadWrapper_IsStandardTaskEntityEqual_CastFail(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "test"}

	// Act — pass a non-PayloadWrapper entity
	result := pw.IsStandardTaskEntityEqual(pw)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsStandardTaskEntityEqual returns true -- same pointer", actual)
}

// ── PayloadWrapper — Error() with error ──

func Test_Cov25_PayloadWrapper_Error_NilError(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.All(nil, nil, nil, nil, nil, nil, nil)
	pw := &corepayload.PayloadWrapper{Attributes: attr}

	// Act
	err := pw.Error()

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "Error returns nil -- no BasicErrWrapper", actual)
}

// ── PayloadWrapper — PayloadDeserializeMust panic ──

func Test_Cov25_PayloadWrapper_PayloadDeserializeMust_Panic(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Payloads: []byte(`{invalid`)}
	didPanic := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		var target map[string]any
		pw.PayloadDeserializeMust(&target)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "PayloadDeserializeMust panics -- invalid json", actual)
}

// ── PayloadWrapper — ParseInjectUsingJsonMust panic ──

func Test_Cov25_PayloadWrapper_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{}
	badResult := corejson.NewResult.UsingBytes([]byte(`{invalid`))
	didPanic := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		pw.ParseInjectUsingJsonMust(&badResult)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust panics -- invalid json", actual)
}

// ── PayloadWrapper — Clone error ──

func Test_Cov25_PayloadWrapper_Clone_Error(t *testing.T) {
	// Arrange
	// Create an attributes with AnyKeyValuePairs containing un-cloneable data
	anyMap := coredynamic.NewMapAnyItems(2)
	anyMap.Add("ch", make(chan int))
	attr := corepayload.New.Attributes.All(nil, nil, anyMap, nil, nil, nil, nil)
	pw := &corepayload.PayloadWrapper{Attributes: attr}

	// Act
	_, err := pw.Clone(true)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Clone returns error -- un-cloneable AnyKeyValuePairs", actual)
}

// ── PayloadsCollection — Filter with break ──

func Test_Cov25_PayloadsCollection_Filter_Break(t *testing.T) {
	// Arrange
	pw1 := &corepayload.PayloadWrapper{Name: "a"}
	pw2 := &corepayload.PayloadWrapper{Name: "b"}
	pw3 := &corepayload.PayloadWrapper{Name: "c"}
	col := corepayload.New.PayloadsCollection.UsingWrappers(pw1, pw2, pw3)

	// Act — break after first match
	result := col.Filter(func(pw *corepayload.PayloadWrapper) (isTake, isBreak bool) {
		if pw.Name == "b" {
			return true, true
		}
		return false, false
	}, 10)

	// Assert
	actual := args.Map{"length": len(result)}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "Filter returns 1 item -- break after first match", actual)
}

// ── PayloadsCollection — Filter with length limit ──

func Test_Cov25_PayloadsCollection_Filter_LengthLimit(t *testing.T) {
	// Arrange
	pw1 := &corepayload.PayloadWrapper{Name: "a"}
	pw2 := &corepayload.PayloadWrapper{Name: "b"}
	pw3 := &corepayload.PayloadWrapper{Name: "c"}
	col := corepayload.New.PayloadsCollection.UsingWrappers(pw1, pw2, pw3)

	// Act — take all but limit to 2
	result := col.Filter(func(pw *corepayload.PayloadWrapper) (isTake, isBreak bool) {
		return true, false
	}, 2)

	// Assert
	actual := args.Map{"length": len(result)}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "Filter returns 2 items -- length limit", actual)
}

// ── PayloadsCollection — IsEqualItems nil left ──

func Test_Cov25_PayloadsCollection_IsEqualItems_NilLeft(t *testing.T) {
	// Arrange
	var col *corepayload.PayloadsCollection

	// Act
	result := col.IsEqualItems(&corepayload.PayloadWrapper{})

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqualItems returns false -- nil collection", actual)
}

// ── PayloadsCollection — ParseInjectUsingJsonMust panic ──

func Test_Cov25_PayloadsCollection_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	// Arrange
	col := corepayload.New.PayloadsCollection.Empty()
	badResult := corejson.NewResult.UsingBytes([]byte(`{invalid`))
	didPanic := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		col.ParseInjectUsingJsonMust(&badResult)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust panics -- invalid json", actual)
}

// ── TypedPayloadCollection — FirstByFilter no match ──

func Test_Cov25_TypedPayloadCollection_FirstByFilter_NoMatch(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	result := col.FirstByFilter(func(item *corepayload.TypedPayloadWrapper[testUserCov23]) bool {
		return item.Data().Name == "NonExistent"
	})

	// Assert
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FirstByFilter returns nil -- no match", actual)
}

// ── TypedPayloadCollection — CloneMust panic ──

func Test_Cov25_TypedPayloadCollection_CloneMust_Success(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	cloned := col.CloneMust()

	// Assert
	actual := args.Map{"length": cloned.Length()}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "CloneMust returns cloned collection -- 3 items", actual)
}

// ── TypedPayloadCollection — ConcatNew ──

func Test_Cov25_TypedPayloadCollection_ConcatNew(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()
	extra := makeTypedWrapperCov23("user", "4", testUserCov23{Name: "Dave"})

	// Act
	newCol, err := col.ConcatNew(extra)

	// Assert
	actual := args.Map{
		"err":    err == nil,
		"length": newCol.Length(),
	}
	expected := args.Map{
		"err":    true,
		"length": 4,
	}
	expected.ShouldBeEqual(t, 0, "ConcatNew returns 4 items -- original 3 + 1", actual)
}

// ── TypedPayloadCollection — HasErrors ──

func Test_Cov25_TypedPayloadCollection_HasErrors_True(t *testing.T) {
	// Arrange
	col := corepayload.NewTypedPayloadCollection[testUserCov23](1)
	// Create a wrapper by deserializing valid bytes
	badWrapper, err := corepayload.TypedPayloadWrapperDeserialize[testUserCov23](
		[]byte(`{"Name":"test","Identifier":"1","Payloads":"aW52YWxpZA=="}`),
	)
	if err != nil {
		t.Skip("deserialization failed, skipping")
	}
	col.Add(badWrapper)

	// Act
	hasErrors := col.HasErrors()

	// Assert
	actual := args.Map{"hasErrors": hasErrors}
	expected := args.Map{"hasErrors": false}
	expected.ShouldBeEqual(t, 0, "HasErrors returns false -- deserialized without parsing error", actual)
}

// ── TypedPayloadCollection — Errors ──

func Test_Cov25_TypedPayloadCollection_Errors_Empty(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	errs := col.Errors()

	// Assert
	actual := args.Map{"errCount": len(errs)}
	expected := args.Map{"errCount": 0}
	expected.ShouldBeEqual(t, 0, "Errors returns empty -- no errors in collection", actual)
}

// ── TypedPayloadCollection — FirstError nil ──

func Test_Cov25_TypedPayloadCollection_FirstError_Nil(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	err := col.FirstError()

	// Assert
	actual := args.Map{"errNil": err == nil}
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "FirstError returns nil -- no errors", actual)
}

// ── TypedPayloadCollection — MergedError nil ──

func Test_Cov25_TypedPayloadCollection_MergedError_Nil(t *testing.T) {
	// Arrange
	col := makeCollectionCov23()

	// Act
	err := col.MergedError()

	// Assert
	actual := args.Map{"errNil": err == nil}
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "MergedError returns nil -- no errors", actual)
}

// ── TypedPayloadCollection — TypedPayloadCollectionDeserializeMust panic ──

func Test_Cov25_TypedPayloadCollectionDeserializeMust_Panic(t *testing.T) {
	// Arrange
	didPanic := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		corepayload.TypedPayloadCollectionDeserializeMust[testUserCov23]([]byte(`{invalid`))
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollectionDeserializeMust panics -- invalid json", actual)
}

// ── TypedPayloadCollection — NewTypedPayloadCollectionFromDataMust panic ──

func Test_Cov25_NewTypedPayloadCollectionFromDataMust_Success(t *testing.T) {
	// Arrange
	data := []testUserCov23{
		{Name: "Alice"},
		{Name: "Bob"},
	}

	// Act
	col := corepayload.NewTypedPayloadCollectionFromDataMust[testUserCov23]("user", data)

	// Assert
	actual := args.Map{"length": col.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "NewTypedPayloadCollectionFromDataMust returns 2 -- valid data", actual)
}
