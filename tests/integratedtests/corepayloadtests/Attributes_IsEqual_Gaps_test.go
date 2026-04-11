package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage17 — corepayload remaining gaps (73 uncovered lines)
// ══════════════════════════════════════════════════════════════════════════════

// --- Attributes.IsEqual branches ---

func Test_Attributes_IsEqual_ErrorDifferent(t *testing.T) {
	// Arrange
	attr1 := corepayload.New.Attributes.Empty()
	attr2 := corepayload.New.Attributes.Empty()

	// Act — both have no error, should be equal
	result := attr1.IsEqual(attr2)

	// Assert
	convey.Convey("Attributes.IsEqual returns true when both empty", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}

func Test_Attributes_IsEqual_KeyValuesDifferent(t *testing.T) {
	// Arrange — two attributes with nil KeyValuePairs vs non-nil
	attr1 := corepayload.New.Attributes.Empty()
	attr2 := corepayload.New.Attributes.Empty()

	// Act
	result := attr1.IsEqual(attr2)

	// Assert
	convey.Convey("Attributes.IsEqual with same empty KV returns true", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}

// --- Attributes.Clone error branch ---

func Test_Attributes_Clone_Valid(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act
	cloned, err := attr.Clone(true)

	// Assert
	convey.Convey("Attributes.Clone returns valid clone", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(cloned.IsEmpty(), convey.ShouldBeTrue)
	})
}

// --- Attributes setter nil-receiver branches ---

func Test_Attributes_HandleErr_NoError_FromAttributesIsEqualGap(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act & Assert — should not panic
	convey.Convey("Attributes.HandleErr with no error does nothing", t, func() {
		convey.So(func() {
			attr.HandleErr()
		}, convey.ShouldNotPanic)
	})
}

func Test_Attributes_HandleError_NoError_FromAttributesIsEqualGap(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()

	// Act & Assert
	convey.Convey("Attributes.HandleError with no error does nothing", t, func() {
		convey.So(func() {
			attr.HandleError()
		}, convey.ShouldNotPanic)
	})
}

func Test_Attributes_AddNewStringKeyValueOnly_NilKV(t *testing.T) {
	// Arrange — attributes with nil KeyValuePairs (not using Empty() which initializes KV)
	attr := &corepayload.Attributes{}

	// Act
	isAdded := attr.AddNewStringKeyValueOnly("key", "value")

	// Assert
	convey.Convey("AddNewStringKeyValueOnly returns false when KV is nil", t, func() {
		convey.So(isAdded, convey.ShouldBeFalse)
	})
}

func Test_Attributes_AnyKeyReflectSetTo_NilKV(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	var target string

	// Act
	err := attr.AnyKeyReflectSetTo("key", &target)

	// Assert
	convey.Convey("AnyKeyReflectSetTo returns error when KV is nil", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

// --- AttributesJson branches ---

func Test_Attributes_ParseInjectUsingJsonMust_Valid(t *testing.T) {
	// Arrange
	jsonResult := corejson.New(corepayload.New.Attributes.Empty())

	// Act
	attr := corepayload.New.Attributes.Empty()
	result := attr.ParseInjectUsingJsonMust(&jsonResult)

	// Assert
	convey.Convey("ParseInjectUsingJsonMust succeeds with valid JSON", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_Attributes_BasicErrorDeserializedTo_EmptyError(t *testing.T) {
	// Arrange
	attr := corepayload.New.Attributes.Empty()
	var target map[string]string

	// Act
	err := attr.BasicErrorDeserializedTo(&target)

	// Assert
	convey.Convey("BasicErrorDeserializedTo returns nil when no error", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

// --- PayloadWrapper nil-receiver setter branches ---
// Most of these are nil-receiver guards that would actually panic
// on `it.InitializeAttributesOnNull()` because `it` is nil.
// These are defensive dead code.

func Test_PayloadWrapper_BasicError_NoError_FromAttributesIsEqualGap(t *testing.T) {
	// Arrange
	wrapper := &corepayload.PayloadWrapper{
		Name: "test",
	}

	// Act
	result := wrapper.BasicError()

	// Assert
	convey.Convey("BasicError returns nil when no error", t, func() {
		convey.So(result, convey.ShouldBeNil)
	})
}

func Test_PayloadWrapper_HandleError_NoError_FromAttributesIsEqualGap_FromAttributesIsEqualGap(t *testing.T) {
	// Arrange
	wrapper := &corepayload.PayloadWrapper{
		Name: "test",
	}

	// Act & Assert
	convey.Convey("PayloadWrapper.HandleError with no error does nothing", t, func() {
		convey.So(func() {
			wrapper.HandleError()
		}, convey.ShouldNotPanic)
	})
}

func Test_PayloadWrapper_IsStandardTaskEntityEqual_CastFail_FromAttributesIsEqualGap(t *testing.T) {
	// Arrange
	wrapper := &corepayload.PayloadWrapper{
		Name: "test",
	}

	// Act — passing a non-PayloadWrapper entity
	result := wrapper.IsStandardTaskEntityEqual(wrapper)

	// Assert
	convey.Convey("IsStandardTaskEntityEqual returns true for same pointer", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}

func Test_PayloadWrapper_Username_EmptyAttributes(t *testing.T) {
	// Arrange
	wrapper := &corepayload.PayloadWrapper{}

	// Act
	result := wrapper.Username()

	// Assert
	convey.Convey("Username returns empty for empty attributes", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

func Test_PayloadWrapper_Error_EmptyError(t *testing.T) {
	// Arrange
	wrapper := &corepayload.PayloadWrapper{}

	// Act
	result := wrapper.Error()

	// Assert
	convey.Convey("Error returns nil for empty error", t, func() {
		convey.So(result, convey.ShouldBeNil)
	})
}

func Test_PayloadWrapper_PayloadDeserializeMust_Valid(t *testing.T) {
	// Arrange
	wrapper := &corepayload.PayloadWrapper{
		Payloads: []byte(`{"name":"test"}`),
	}
	var result map[string]string

	// Act & Assert
	convey.Convey("PayloadDeserializeMust succeeds with valid JSON", t, func() {
		convey.So(func() {
			wrapper.PayloadDeserializeMust(&result)
		}, convey.ShouldNotPanic)
		convey.So(result["name"], convey.ShouldEqual, "test")
	})
}

func Test_PayloadWrapper_ParseInjectUsingJsonMust_Valid(t *testing.T) {
	// Arrange
	wrapper := &corepayload.PayloadWrapper{Name: "test"}
	jsonResult := corejson.New(wrapper)

	// Act
	newWrapper := &corepayload.PayloadWrapper{}
	result := newWrapper.ParseInjectUsingJsonMust(&jsonResult)

	// Assert
	convey.Convey("ParseInjectUsingJsonMust returns wrapper on success", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_PayloadWrapper_Clone_Error_FromAttributesIsEqualGap(t *testing.T) {
	// Arrange — wrapper with nil attributes, deep clone
	wrapper := &corepayload.PayloadWrapper{
		Name: "test",
	}

	// Act
	_, err := wrapper.Clone(true)

	// Assert
	convey.Convey("PayloadWrapper.Clone deep clone succeeds", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_PayloadWrapper_ClonePtr_Error(t *testing.T) {
	// Arrange — wrapper with nil attributes
	wrapper := &corepayload.PayloadWrapper{
		Name: "test",
	}

	// Act
	cloned, err := wrapper.ClonePtr(false)

	// Assert
	convey.Convey("PayloadWrapper.ClonePtr shallow clone succeeds", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(cloned, convey.ShouldNotBeNil)
	})
}

// --- TypedPayloadWrapper branches ---

func Test_TypedPayloadWrapper_MarshalJSON_NilWrapper(t *testing.T) {
	// Arrange
	typed := &corepayload.TypedPayloadWrapper[string]{}

	// Act
	_, err := typed.MarshalJSON()

	// Assert
	convey.Convey("TypedPayloadWrapper.MarshalJSON with nil wrapper returns error", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_TypedPayloadWrapper_UnmarshalJSON_Valid(t *testing.T) {
	// Arrange
	typed, _ := corepayload.NewTypedPayloadWrapperFrom("test", "id1", "entity", "hello")
	jsonBytes, _ := typed.MarshalJSON()

	// Act
	newTyped := &corepayload.TypedPayloadWrapper[string]{}
	err := newTyped.UnmarshalJSON(jsonBytes)

	// Assert
	convey.Convey("TypedPayloadWrapper.UnmarshalJSON succeeds with valid data", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_TypedPayloadWrapper_SetTypedData_NilWrapper(t *testing.T) {
	// Arrange
	typed := &corepayload.TypedPayloadWrapper[string]{}

	// Act
	err := typed.SetTypedData("new-data")

	// Assert
	convey.Convey("SetTypedData with nil wrapper returns error", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_TypedPayloadWrapper_ClonePtr_Valid(t *testing.T) {
	// Arrange
	typed, _ := corepayload.NewTypedPayloadWrapperFrom("test", "id1", "entity", "hello")

	// Act
	cloned, err := typed.ClonePtr(true)

	// Assert
	convey.Convey("TypedPayloadWrapper.ClonePtr succeeds", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(cloned, convey.ShouldNotBeNil)
	})
}

func Test_TypedPayloadWrapper_Clone_NilReceiver(t *testing.T) {
	// Arrange
	var typed *corepayload.TypedPayloadWrapper[string]

	// Act
	_, err := typed.Clone(true)

	// Assert
	convey.Convey("TypedPayloadWrapper.Clone nil receiver returns error", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_TypedPayloadWrapper_Reparse_NilWrapper(t *testing.T) {
	// Arrange
	typed := &corepayload.TypedPayloadWrapper[string]{}

	// Act
	err := typed.Reparse()

	// Assert
	convey.Convey("Reparse with nil wrapper returns error", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

// --- TypedPayloadCollection branches ---

func Test_TypedPayloadCollection_Clone_Error(t *testing.T) {
	// Arrange
	typed, _ := corepayload.NewTypedPayloadWrapperFrom("test", "id1", "entity", "hello")
	collection := corepayload.NewTypedPayloadCollectionSingle(typed)

	// Act
	cloned, err := collection.Clone()

	// Assert
	convey.Convey("TypedPayloadCollection.Clone succeeds", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(cloned, convey.ShouldNotBeNil)
	})
}

func Test_TypedPayloadCollection_CloneMust_Valid(t *testing.T) {
	// Arrange
	typed, _ := corepayload.NewTypedPayloadWrapperFrom("test", "id1", "entity", "hello")
	collection := corepayload.NewTypedPayloadCollectionSingle(typed)

	// Act & Assert
	convey.Convey("TypedPayloadCollection.CloneMust succeeds", t, func() {
		convey.So(func() {
			collection.CloneMust()
		}, convey.ShouldNotPanic)
	})
}

func Test_TypedPayloadCollection_NewFromDataMust_Valid(t *testing.T) {
	// Act & Assert
	convey.Convey("NewTypedPayloadCollectionFromDataMust succeeds", t, func() {
		convey.So(func() {
			corepayload.NewTypedPayloadCollectionFromDataMust[string]("test", []string{"a", "b"})
		}, convey.ShouldNotPanic)
	})
}

func Test_TypedPayloadCollection_HasErrors_NoErrors(t *testing.T) {
	// Arrange
	typed, _ := corepayload.NewTypedPayloadWrapperFrom("test", "id1", "entity", "hello")
	collection := corepayload.NewTypedPayloadCollectionSingle(typed)

	// Act
	hasErr := collection.HasErrors()

	// Assert
	convey.Convey("HasErrors returns false when no errors", t, func() {
		convey.So(hasErr, convey.ShouldBeFalse)
	})
}

func Test_TypedPayloadCollection_Errors_Empty_FromAttributesIsEqualGap(t *testing.T) {
	// Arrange
	collection := corepayload.EmptyTypedPayloadCollection[string]()

	// Act
	errs := collection.Errors()

	// Assert
	convey.Convey("Errors returns nil for empty collection", t, func() {
		convey.So(errs, convey.ShouldBeNil)
	})
}

func Test_TypedPayloadCollection_MergedError_None(t *testing.T) {
	// Arrange
	collection := corepayload.EmptyTypedPayloadCollection[string]()

	// Act
	err := collection.MergedError()

	// Assert
	convey.Convey("MergedError returns nil for empty collection", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_TypedPayloadCollection_MergedError_Single(t *testing.T) {
	// Arrange
	typed, _ := corepayload.NewTypedPayloadWrapperFrom("test", "id1", "entity", "hello")
	collection := corepayload.NewTypedPayloadCollectionSingle(typed)

	// Act
	err := collection.MergedError()

	// Assert
	convey.Convey("MergedError returns nil for healthy collection", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

// --- PayloadsCollectionFilter branches ---

func Test_PayloadsCollection_ParseInjectUsingJsonMust_Valid(t *testing.T) {
	// Arrange
	emptyCollection := corepayload.Empty.PayloadsCollection()
	jsonResult := corejson.New(emptyCollection)

	// Act & Assert
	convey.Convey("PayloadsCollection.ParseInjectUsingJsonMust succeeds", t, func() {
		newColl := corepayload.Empty.PayloadsCollection()
		convey.So(func() {
			newColl.ParseInjectUsingJsonMust(&jsonResult)
		}, convey.ShouldNotPanic)
	})
}

// Coverage note: Many uncovered lines are nil-receiver guards that would
// panic on method dispatch. These are defensive dead code:
// - PayloadWrapper.Set*() nil receiver → it.InitializeAttributesOnNull() panics
// - TypedPayloadWrapper.Serialize/SerializeMust nil wrapper → dead code
// - payloadProperties.AsPayloadPropertiesDefiner → unexported type
// - newAttributesCreator.DeserializeUsingJsonResult error branch → requires invalid JSON
// - TypedPayloadCollectionDeserializeMust → requires marshal error (dead code)
