package corepayloadtests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage26 — corepayload remaining gaps (Iteration 28)
//
// Targets:
//   - Attributes: IsEqual error-different, Clone err, CloneInterface HasError
//   - AttributesSetters: HandleErr, HandleError, MustBeEmptyError
//   - AttributesGetters: fallthrough paths
//   - PayloadWrapper: BasicError no-error, PayloadDeserializeToPayloadBinder HasError
//   - PayloadWrapper: SetPayloadDynamic/SetPayloadDynamicAny/SetAuthInfo/SetUserInfo nil guard
//   - PayloadWrapper: initializeAuthOnDemand, HandleError, IsEqualInterface cast fail
//   - PayloadWrapper: Error() with error
//   - TypedPayloadWrapper: HandleError, UnmarshalJSON nil/error, SerializeMust
//   - TypedPayloadWrapper: SetTypedData error, SetTypedDataMust, ClonePtr error, Clone error
//   - TypedPayloadWrapper: reparse error
//   - TypedPayloadCollection: Clone error, CloneMust, ClonePtr error
//   - TypedPayloadCollection: HasErrors, Errors, FirstError, MergedError
//   - TypedPayloadCollection: NewFromData error
//   - generic_helpers: error path
//   - newAttributesCreator: fallthrough
//   - newTypedPayloadWrapperCreator: error paths
//   - payloadProperties: line
//   - typed_collection_funcs: nil item skip
// ══════════════════════════════════════════════════════════════════════════════

// ---------- Attributes: IsEqual error-different ----------

func Test_I28_Attributes_IsEqual_ErrorDifferent(t *testing.T) {
	// Arrange
	a1 := corepayload.New.Attributes.All(
		errcore.New.BasicErr.Error(errcore.New.Err.Message("err1")),
		nil, nil, nil, nil, nil, nil,
	)
	a2 := corepayload.New.Attributes.All(
		errcore.New.BasicErr.Error(errcore.New.Err.Message("err2")),
		nil, nil, nil, nil, nil, nil,
	)

	// Act
	result := a1.IsEqual(a2)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- different errors", actual)
}

// ---------- AttributesSetters: HandleErr with error ----------

func Test_I28_Attributes_HandleErr_NoError(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.All(
		nil, nil, nil, nil, nil, nil, nil,
	)

	// Act & Assert — should not panic
	a.HandleErr()

	actual := args.Map{"completed": true}
	expected := args.Map{"completed": true}
	expected.ShouldBeEqual(t, 0, "HandleErr completes -- no error", actual)
}

// ---------- AttributesSetters: HandleError with no error ----------

func Test_I28_Attributes_HandleError_NoError(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.All(
		nil, nil, nil, nil, nil, nil, nil,
	)

	// Act & Assert — should not panic
	a.HandleError()

	actual := args.Map{"completed": true}
	expected := args.Map{"completed": true}
	expected.ShouldBeEqual(t, 0, "HandleError completes -- no error", actual)
}

// ---------- AttributesSetters: MustBeEmptyError no error ----------

func Test_I28_Attributes_MustBeEmptyError_NoError(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.All(
		nil, nil, nil, nil, nil, nil, nil,
	)

	// Act & Assert — should not panic because no error
	a.MustBeEmptyError()

	actual := args.Map{"completed": true}
	expected := args.Map{"completed": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmptyError completes -- no error", actual)
}

// ---------- PayloadWrapper: BasicError no error ----------

func Test_I28_PayloadWrapper_BasicError_NoError(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.UsingPayload([]byte(`{}`))

	// Act
	result := pw.BasicError()

	// Assert
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "BasicError returns nil -- no error", actual)
}

// ---------- PayloadWrapper: HandleError no error ----------

func Test_I28_PayloadWrapper_HandleError_NoError(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.UsingPayload([]byte(`{}`))

	// Act & Assert — should not panic
	pw.HandleError()

	actual := args.Map{"completed": true}
	expected := args.Map{"completed": true}
	expected.ShouldBeEqual(t, 0, "HandleError completes -- no error", actual)
}

// ---------- PayloadWrapper: IsEqualInterface cast fail ----------

func Test_I28_PayloadWrapper_IsEqualInterface_CastFail(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.UsingPayload([]byte(`{}`))

	// Act
	result := pw.IsEqualInterface("not-a-payload-wrapper")

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqualInterface returns false -- cast fail", actual)
}

// ---------- PayloadWrapper: Error with error ----------

func Test_I28_PayloadWrapper_Error_WithError(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.ErrorWrapper(
		errcore.New.BasicErr.Error(errcore.New.Err.Message("test-err")),
	)

	// Act
	result := pw.Error()

	// Assert
	actual := args.Map{"hasErr": result != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Error returns error -- with error", actual)
}

// ---------- TypedPayloadCollection: HasErrors/Errors/FirstError/MergedError ----------

func Test_I28_TypedPayloadCollection_ErrorMethods(t *testing.T) {
	// Arrange — create a typed collection from valid data first
	type simpleUser struct {
		Name string `json:"name"`
	}

	items := []simpleUser{
		{Name: "alice"},
		{Name: "bob"},
	}
	collection, err := corepayload.NewTypedPayloadCollectionFromData[simpleUser]("users", items)
	if err != nil {
		t.Fatalf("unexpected err creating collection: %v", err)
	}

	// Act
	hasErrors := collection.HasErrors()
	errs := collection.Errors()
	firstErr := collection.FirstError()
	mergedErr := collection.MergedError()

	// Assert
	actual := args.Map{
		"hasErrors":     hasErrors,
		"errCount":      len(errs),
		"firstErrNil":   firstErr == nil,
		"mergedErrNil":  mergedErr == nil,
	}
	expected := args.Map{
		"hasErrors":     false,
		"errCount":      0,
		"firstErrNil":   true,
		"mergedErrNil":  true,
	}
	expected.ShouldBeEqual(t, 0, "Error methods return clean -- no errors in collection", actual)
}

// ---------- TypedPayloadWrapper: UnmarshalJSON invalid data ----------

func Test_I28_TypedPayloadWrapper_UnmarshalJSON_InvalidData(t *testing.T) {
	// Arrange
	type simpleData struct {
		Val string `json:"val"`
	}
	tw := &corepayload.TypedPayloadWrapper[simpleData]{}

	// Act
	err := json.Unmarshal([]byte(`not-json`), tw)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON returns error -- invalid json", actual)
}

// ---------- TypedPayloadWrapper: HandleError no error ----------

func Test_I28_TypedPayloadWrapper_HandleError_NoError(t *testing.T) {
	// Arrange
	type simpleData struct {
		Val string `json:"val"`
	}
	data := simpleData{Val: "test"}
	tw, err := corepayload.TypedPayloadWrapperRecord[simpleData](data)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	// Act & Assert — should not panic
	tw.HandleError()

	actual := args.Map{"completed": true}
	expected := args.Map{"completed": true}
	expected.ShouldBeEqual(t, 0, "HandleError completes -- no error", actual)
}

// ---------- PayloadWrapper: PayloadDeserializeToPayloadBinder HasError ----------

func Test_I28_PayloadWrapper_PayloadDeserializeToPayloadBinder_HasError(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.ErrorWrapper(
		errcore.New.BasicErr.Error(errcore.New.Err.Message("deser-err")),
	)

	// Act
	result, err := pw.PayloadDeserializeToPayloadBinder()

	// Assert
	actual := args.Map{
		"resultNil": result == nil,
		"hasErr":    err != nil,
	}
	expected := args.Map{
		"resultNil": true,
		"hasErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadDeserializeToPayloadBinder returns error -- has error", actual)
}

// ---------- Attributes: CloneInterface HasError ----------

func Test_I28_Attributes_CloneInterface_WithError(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.All(
		errcore.New.BasicErr.Error(errcore.New.Err.Message("clone-err")),
		nil, nil, nil, nil, nil, nil,
	)

	// Act
	cloned := a.CloneInterface(true)

	// Assert
	actual := args.Map{"hasError": cloned.HasError()}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "CloneInterface preserves error -- with error", actual)
}
