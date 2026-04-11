package corepayloadtests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage26 — corepayload remaining gaps (Iteration 28)
//
// API Reference (verified from source):
//   - corepayload.New.PayloadWrapper.Create(name, id, taskName, category, record)
//     returns (*PayloadWrapper, error) — TWO return values
// ══════════════════════════════════════════════════════════════════════════════

// ---------- AttributesSetters: HandleErr no error ----------

func Test_Attributes_HandleErr_NoError(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()

	// Act & Assert — should not panic
	a.HandleErr()

	actual := args.Map{"completed": true}
	expected := args.Map{"completed": true}
	expected.ShouldBeEqual(t, 0, "HandleErr completes -- no error", actual)
}

// ---------- AttributesSetters: HandleError no error ----------

func Test_Attributes_HandleError_NoError(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()

	// Act & Assert — should not panic
	a.HandleError()

	actual := args.Map{"completed": true}
	expected := args.Map{"completed": true}
	expected.ShouldBeEqual(t, 0, "HandleError completes -- no error", actual)
}

// ---------- AttributesSetters: MustBeEmptyError no error ----------

func Test_Attributes_MustBeEmptyError_NoError(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()

	// Act & Assert
	a.MustBeEmptyError()

	actual := args.Map{"completed": true}
	expected := args.Map{"completed": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmptyError completes -- no error", actual)
}

// ---------- PayloadWrapper: BasicError no error ----------

func Test_PayloadWrapper_BasicError_NoError(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.Empty()

	// Act
	result := pw.BasicError()

	// Assert
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "BasicError returns nil -- no error", actual)
}

// ---------- PayloadWrapper: HandleError no error ----------

func Test_PayloadWrapper_HandleError_NoError(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.Empty()

	// Act & Assert
	pw.HandleError()

	actual := args.Map{"completed": true}
	expected := args.Map{"completed": true}
	expected.ShouldBeEqual(t, 0, "HandleError completes -- no error", actual)
}

// ---------- PayloadWrapper: IsStandardTaskEntityEqual different wrapper ----------

func Test_PayloadWrapper_IsStandardTaskEntityEqual_Different(t *testing.T) {
	// Arrange
	pw1 := corepayload.New.PayloadWrapper.Empty()
	// Create returns (*PayloadWrapper, error) — handle BOTH return values
	pw2, err := corepayload.New.PayloadWrapper.Create("other", "id-99", "task", "cat", "data")
	if err != nil {
		panic(err)
	}

	// Act
	result := pw1.IsStandardTaskEntityEqual(pw2)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsStandardTaskEntityEqual returns false -- different wrapper", actual)
}

// ---------- PayloadWrapper: Error no error ----------

func Test_PayloadWrapper_Error_NoError(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.Empty()

	// Act
	result := pw.Error()

	// Assert
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Error returns nil -- no error", actual)
}

// ---------- TypedPayloadWrapper: HandleError no error ----------

func Test_TypedPayloadWrapper_HandleError_NoError(t *testing.T) {
	// Arrange
	type simpleData struct {
		Val string `json:"val"`
	}
	data := simpleData{Val: "test"}
	tw, err := corepayload.TypedPayloadWrapperRecord[simpleData]("test", "id-1", "task", "cat", data)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected err:", actual)

	// Act & Assert
	tw.HandleError()

	actual = args.Map{"completed": true}
	expected = args.Map{"completed": true}
	expected.ShouldBeEqual(t, 0, "HandleError completes -- no error", actual)
}

// ---------- TypedPayloadWrapper: UnmarshalJSON invalid data ----------

func Test_TypedPayloadWrapper_UnmarshalJSON_InvalidData(t *testing.T) {
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

// ---------- TypedPayloadCollection: HasErrors/Errors/FirstError/MergedError ----------

func Test_TypedPayloadCollection_ErrorMethods_NoErrors(t *testing.T) {
	// Arrange
	type simpleUser struct {
		Name string `json:"name"`
	}
	items := []simpleUser{{Name: "alice"}, {Name: "bob"}}
	collection, err := corepayload.NewTypedPayloadCollectionFromData[simpleUser]("users", items)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected err creating collection:", actual)

	// Act
	hasErrors := collection.HasErrors()
	errs := collection.Errors()
	firstErr := collection.FirstError()
	mergedErr := collection.MergedError()

	// Assert
	actual = args.Map{
		"hasErrors":    hasErrors,
		"errCount":     len(errs),
		"firstErrNil":  firstErr == nil,
		"mergedErrNil": mergedErr == nil,
	}
	expected = args.Map{
		"hasErrors":    false,
		"errCount":     0,
		"firstErrNil":  true,
		"mergedErrNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Error methods return clean -- no errors in collection", actual)
}

// ---------- TypedPayloadCollection: Clone ----------

func Test_TypedPayloadCollection_Clone(t *testing.T) {
	// Arrange
	type simpleUser struct {
		Name string `json:"name"`
	}
	items := []simpleUser{{Name: "alice"}}
	collection, err := corepayload.NewTypedPayloadCollectionFromData[simpleUser]("users", items)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected err:", actual)

	// Act
	cloned, err := collection.Clone()

	// Assert
	actual = args.Map{
		"errNil": err == nil,
		"length": cloned.Length(),
	}
	expected = args.Map{
		"errNil": true,
		"length": 1,
	}
	expected.ShouldBeEqual(t, 0, "Clone returns valid copy -- single item", actual)
}

// ---------- TypedPayloadWrapper: ClonePtr ----------

func Test_TypedPayloadWrapper_ClonePtr(t *testing.T) {
	// Arrange
	type simpleData struct {
		Val string `json:"val"`
	}
	data := simpleData{Val: "test"}
	tw, err := corepayload.TypedPayloadWrapperRecord[simpleData]("test", "id1", "task", "cat", data)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected err:", actual)

	// Act
	cloned, err := tw.ClonePtr(true)

	// Assert
	actual = args.Map{
		"errNil": err == nil,
		"notNil": cloned != nil,
	}
	expected = args.Map{
		"errNil": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns valid copy -- deep clone", actual)
}

// ---------- TypedPayloadWrapper: Clone ----------

func Test_TypedPayloadWrapper_Clone(t *testing.T) {
	// Arrange
	type simpleData struct {
		Val string `json:"val"`
	}
	data := simpleData{Val: "test"}
	tw, err := corepayload.TypedPayloadWrapperRecord[simpleData]("test", "id1", "task", "cat", data)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected err:", actual)

	// Act
	cloned, err := tw.Clone(true)

	// Assert
	actual = args.Map{
		"errNil": err == nil,
	}
	expected = args.Map{
		"errNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Clone returns valid copy -- deep clone", actual)
	_ = cloned
}

// ---------- TypedPayloadWrapper: SetTypedData ----------

func Test_TypedPayloadWrapper_SetTypedData(t *testing.T) {
	// Arrange
	type simpleData struct {
		Val string `json:"val"`
	}
	data := simpleData{Val: "initial"}
	tw, err := corepayload.TypedPayloadWrapperRecord[simpleData]("test", "id1", "task", "cat", data)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected err:", actual)

	// Act
	newData := simpleData{Val: "updated"}
	err = tw.SetTypedData(newData)

	// Assert
	actual = args.Map{
		"errNil":  err == nil,
		"updated": tw.TypedData().Val,
	}
	expected = args.Map{
		"errNil":  true,
		"updated": "updated",
	}
	expected.ShouldBeEqual(t, 0, "SetTypedData updates data -- valid data", actual)
}

// ---------- PayloadsCollectionFilter: empty items ----------

func Test_PayloadsCollection_FilterEmpty(t *testing.T) {
	// Arrange
	pc := corepayload.New.PayloadsCollection.UsingCap(0)

	// Act
	result := pc.Filter(func(pw *corepayload.PayloadWrapper) (isTake, isBreak bool) {
		return true, false
	})

	// Assert
	actual := args.Map{"isEmpty": len(result) == 0}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Filter returns empty -- empty collection", actual)
}
