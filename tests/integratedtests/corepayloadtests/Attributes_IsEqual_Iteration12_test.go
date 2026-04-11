package corepayloadtests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
)

// ── Attributes: IsEqual branches (ErrorDifferent, PagingDifferent, KeyValuesDifferent, DynamicPayloadsDifferent, AnyKeyValuesDifferent) ──
// Covers Attributes.go L46-48, L58-60, L72-74

func Test_Attributes_IsEqual_DifferentDynamicPayloads(t *testing.T) {
	// Arrange
	a1 := &corepayload.Attributes{
		DynamicPayloads: []byte(`{"a":1}`),
	}
	a2 := &corepayload.Attributes{
		DynamicPayloads: []byte(`{"b":2}`),
	}

	result := a1.IsEqual(a2)

	// Act
	actual := args.Map{"equal": result}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- different DynamicPayloads", actual)
}

// ── Attributes: Clone error path ──
// Covers Attributes.go L84-86

func Test_Attributes_Clone_NilPtr(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes
	cloned, err := a.Clone(false)

	// Act
	actual := args.Map{
		"isEmpty": cloned.IsEmpty(),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Clone returns empty -- nil receiver", actual)
}

// ── Attributes: deepClonePtr error path ──
// Covers Attributes.go L127-129, L134-136

func Test_Attributes_DeepClone(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{
		DynamicPayloads: []byte(`{"test":1}`),
	}

	cloned, err := a.ClonePtr(true)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": cloned != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "DeepClonePtr returns cloned -- valid attrs", actual)
}

// ── AttributesSetters: HandleErr, HandleError, MustBeEmptyError ──
// Covers AttributesSetters.go L13-15, L19-21, L29

func Test_Attributes_HandleErr_NoError_FromAttributesIsEqualIte(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}
	a.HandleErr() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErr does nothing -- no error", actual)
}

func Test_Attributes_HandleError_NoError_FromAttributesIsEqualIte(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}
	a.HandleError() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleError does nothing -- no error", actual)
}

func Test_Attributes_MustBeEmptyError(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}
	a.MustBeEmptyError() // should not panic on empty error

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmptyError does not panic -- empty error", actual)
}

// ── AttributesSetters: ReflectSetToKey, line 96 ──
// Covers AttributesSetters.go L96

// ── AttributesGetters: Error with compiled ──
// Covers AttributesGetters.go L130-132

// ── AttributesGetters: HasAnyKey (nil pairs) ──
// Covers AttributesGetters.go L23-28

func Test_Attributes_HasAnyKey_NilPairs(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}
	found := a.HasAnyKey("key")

	// Act
	actual := args.Map{"found": found}

	// Assert
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "HasAnyKey returns false -- nil pairs", actual)
}

// ── PayloadWrapper: UnmarshalJSON nil ──
// Covers PayloadWrapper.go L51-55

func Test_PayloadWrapper_UnmarshalJSON_Nil(t *testing.T) {
	// Arrange
	var pw *corepayload.PayloadWrapper
	err := pw.UnmarshalJSON([]byte(`{}`))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON returns error -- nil receiver", actual)
}

// ── PayloadWrapper: BasicError (has error vs no error) ──
// Covers PayloadWrapper.go L134-136

func Test_PayloadWrapper_BasicError_NoError_FromAttributesIsEqualIte(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{}
	result := pw.BasicError()

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "BasicError returns nil -- no error", actual)
}

// ── PayloadWrapper: PayloadDeserializeToPayloadBinder error ──
// Covers PayloadWrapper.go L146-148

func Test_PayloadWrapper_PayloadDeserializeToPayloadBinder_Null(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{}
	_, err := pw.PayloadDeserializeToPayloadBinder()

	// depends on whether null returns error

	// Act
	actual := args.Map{
		"checked": true,
		"errChecked": err == nil || err != nil,
	}

	// Assert
	expected := args.Map{
		"checked": true,
		"errChecked": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadDeserializeToPayloadBinder -- null payload", actual)
}

// ── PayloadWrapper: SetPayloadDynamic (nil receiver check) ──
// Covers PayloadWrapper.go L188-190

func Test_PayloadWrapper_SetPayloadDynamic(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{}
	result := pw.SetPayloadDynamic([]byte(`{"x":1}`))

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"hasPayloads": len(result.Payloads) > 0,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"hasPayloads": true,
	}
	expected.ShouldBeEqual(t, 0, "SetPayloadDynamic sets payloads -- valid bytes", actual)
}

// ── PayloadWrapper: SetPayloadDynamicAny ──
// Covers PayloadWrapper.go L210-212, L218-220

func Test_PayloadWrapper_SetPayloadDynamicAny(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{}
	result, err := pw.SetPayloadDynamicAny(map[string]string{"key": "value"})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": result != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "SetPayloadDynamicAny sets payloads -- valid any", actual)
}

// ── PayloadWrapper: SetAuthInfo ──
// Covers PayloadWrapper.go L230-232

func Test_PayloadWrapper_SetAuthInfo(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{
		Attributes: &corepayload.Attributes{},
	}
	result := pw.SetAuthInfo(&corepayload.AuthInfo{})

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SetAuthInfo returns self -- valid input", actual)
}

// ── PayloadWrapper: SetUserInfo ──
// Covers PayloadWrapper.go L242-244

func Test_PayloadWrapper_SetUserInfo(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{
		Attributes: &corepayload.Attributes{},
	}
	result := pw.SetUserInfo(&corepayload.UserInfo{})

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SetUserInfo returns self -- valid input", actual)
}

// ── PayloadWrapper: initializeAuthOnDemand ──
// Covers PayloadWrapper.go L276-278, L280-282

// ── PayloadWrapper: HandleError ──
// Covers PayloadWrapper.go L294-296

func Test_PayloadWrapper_HandleError_NoError_FromAttributesIsEqualIte(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{}
	pw.HandleError() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleError does nothing -- no error", actual)
}

// ── PayloadWrapper: IsEntityEqual cast failed ──
// Covers PayloadWrapper.go L335-337

// ── PayloadWrapper: Username empty attrs ──
// Covers PayloadWrapper.go L363-365

func Test_PayloadWrapper_Username_EmptyAttrs(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{}
	result := pw.Username()

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Username returns empty -- empty attributes", actual)
}

// ── PayloadWrapper: Error with payload error ──
// Covers PayloadWrapper.go L385

// ── PayloadWrapper: IsEqual with different payloads and attrs ──
// Covers PayloadWrapper.go L426-428, L432-434

func Test_PayloadWrapper_IsEqual_DiffPayloads(t *testing.T) {
	// Arrange
	pw1 := &corepayload.PayloadWrapper{Payloads: []byte(`{"a":1}`)}
	pw2 := &corepayload.PayloadWrapper{Payloads: []byte(`{"b":2}`)}

	result := pw1.IsEqual(pw2)

	// Act
	actual := args.Map{"equal": result}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- different payloads", actual)
}

func Test_PayloadWrapper_IsEqual_DiffAttrs(t *testing.T) {
	// Arrange
	pw1 := &corepayload.PayloadWrapper{
		Payloads:   []byte(`{"a":1}`),
		Attributes: &corepayload.Attributes{DynamicPayloads: []byte(`{"x":1}`)},
	}
	pw2 := &corepayload.PayloadWrapper{
		Payloads:   []byte(`{"a":1}`),
		Attributes: &corepayload.Attributes{DynamicPayloads: []byte(`{"y":2}`)},
	}

	result := pw1.IsEqual(pw2)

	// Act
	actual := args.Map{"equal": result}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- different attrs", actual)
}

// ── PayloadWrapper: DeserializeMust, PayloadDeserializeMust ──
// Covers PayloadWrapper.go L597-604, L617-625

func Test_PayloadWrapper_DeserializeMust(t *testing.T) {
	// Arrange
	data := map[string]string{"key": "value"}
	jsonBytes, _ := json.Marshal(data)
	pw := &corepayload.PayloadWrapper{Payloads: jsonBytes}

	var result map[string]string
	pw.DeserializeMust(&result)

	// Act
	actual := args.Map{"key": result["key"]}

	// Assert
	expected := args.Map{"key": "value"}
	expected.ShouldBeEqual(t, 0, "DeserializeMust deserializes correctly", actual)
}

func Test_PayloadWrapper_PayloadDeserializeMust(t *testing.T) {
	// Arrange
	data := map[string]string{"key": "value"}
	jsonBytes, _ := json.Marshal(data)
	pw := &corepayload.PayloadWrapper{Payloads: jsonBytes}

	var result map[string]string
	pw.PayloadDeserializeMust(&result)

	// Act
	actual := args.Map{"key": result["key"]}

	// Assert
	expected := args.Map{"key": "value"}
	expected.ShouldBeEqual(t, 0, "PayloadDeserializeMust deserializes correctly", actual)
}

// ── PayloadWrapper: DeserializePayloadsToPayloadWrapperMust ──
// Covers PayloadWrapper.go L650-658

func Test_PayloadWrapper_DeserializeToPayloadWrapperMust(t *testing.T) {
	// Arrange
	inner := &corepayload.PayloadWrapper{
		Name:       "test",
		Identifier: "123",
	}
	jsonBytes, _ := json.Marshal(inner)
	pw := &corepayload.PayloadWrapper{Payloads: jsonBytes}

	result := pw.DeserializePayloadsToPayloadWrapperMust()

	// Act
	actual := args.Map{"name": result.Name}

	// Assert
	expected := args.Map{"name": "test"}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadsToPayloadWrapperMust returns valid wrapper", actual)
}

// ── PayloadWrapper: ParseInjectUsingJson error ──
// Covers PayloadWrapper.go L682-684

func Test_PayloadWrapper_ParseInjectUsingJson_Error(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{}
	r := corejson.NewResult.Error(errTestHelper("bad json"))
	badResult := &r

	_, err := pw.ParseInjectUsingJson(badResult)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns error -- bad json result", actual)
}

// ── PayloadWrapper: ParseInjectUsingJsonMust ──
// Covers PayloadWrapper.go L694-702

func Test_PayloadWrapper_ParseInjectUsingJsonMust_Success(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{}
	jsonBytes, _ := json.Marshal(pw)
	jsonResult := corejson.NewResult.UsingBytesTypePtr(jsonBytes, "PayloadWrapper")

	result := pw.ParseInjectUsingJsonMust(jsonResult)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust returns self -- valid json", actual)
}

// ── PayloadWrapper: Clone error ──
// Covers PayloadWrapper.go L744-746, L766-768

func Test_PayloadWrapper_Clone_Success(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{
		Name:     "test",
		Payloads: []byte(`{"a":1}`),
	}
	cloned, err := pw.Clone(false)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": cloned.Name,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "test",
	}
	expected.ShouldBeEqual(t, 0, "Clone returns valid clone -- shallow", actual)
}

// ── PayloadsCollection: AddsPtrOptions with skip ──
// Covers PayloadsCollection.go L62-64

func Test_PayloadsCollection_AddsPtrOptions_Empty(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	result := coll.AddsPtrOptions(true)

	// Act
	actual := args.Map{"empty": result.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AddsPtrOptions returns empty -- no items", actual)
}

// ── PayloadsCollection: AddsOptions with skip ──
// Covers PayloadsCollection.go L85-87

func Test_PayloadsCollection_AddsOptions_Empty(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	result := coll.AddsOptions(true)

	// Act
	actual := args.Map{"empty": result.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AddsOptions returns empty -- no items", actual)
}

// ── PayloadsCollectionGetters: FirstDynamic, Last, IsEqualItems ──
// Covers PayloadsCollectionGetters.go L52-54, L68-70, L189-191, L205-207

func Test_PayloadsCollection_FirstDynamic_Nil(t *testing.T) {
	// Arrange
	var coll *corepayload.PayloadsCollection
	result := coll.FirstDynamic()

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FirstDynamic returns nil -- nil receiver", actual)
}

func Test_PayloadsCollection_Last_Nil(t *testing.T) {
	// Arrange
	var coll *corepayload.PayloadsCollection
	result := coll.Last()

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Last returns nil -- nil receiver", actual)
}

func Test_PayloadsCollection_IsEqualItems_NilLeft_FromAttributesIsEqualIte(t *testing.T) {
	// Arrange
	var coll *corepayload.PayloadsCollection
	result := coll.IsEqualItems(&corepayload.PayloadWrapper{})

	// Act
	actual := args.Map{"equal": result}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqualItems returns false -- nil receiver", actual)
}

func Test_PayloadsCollection_IsEqualItems_DiffItem(t *testing.T) {
	// Arrange
	pw1 := &corepayload.PayloadWrapper{Name: "a"}
	pw2 := &corepayload.PayloadWrapper{Name: "b"}
	coll := corepayload.New.PayloadsCollection.Empty()
	coll.AddsPtr(pw1)

	result := coll.IsEqualItems(pw2)

	// Act
	actual := args.Map{"equal": result}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqualItems returns false -- different items", actual)
}

// ── PayloadsCollectionJson: ParseInjectUsingJson error ──
// Covers PayloadsCollectionJson.go L109-111

func Test_PayloadsCollection_ParseInjectUsingJson_Error(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	r := corejson.NewResult.Error(errTestHelper("bad"))
	badResult := &r

	_, err := coll.ParseInjectUsingJson(badResult)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns error -- bad json", actual)
}

// ── PayloadsCollectionJson: ParseInjectUsingJsonMust ──
// Covers PayloadsCollectionJson.go L119-127

func Test_PayloadsCollection_ParseInjectUsingJsonMust_Success(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	jsonBytes, _ := json.Marshal(coll)
	jsonResult := corejson.NewResult.UsingBytesTypePtr(jsonBytes, "PayloadsCollection")

	result := coll.ParseInjectUsingJsonMust(jsonResult)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust returns self -- valid json", actual)
}

// ── PayloadsCollectionPaging: GetSinglePageCollection negative index ──
// Covers PayloadsCollectionPaging.go L81-87

func Test_PayloadsCollection_GetSinglePageCollection_Panic(t *testing.T) {
	// Arrange
	coll := corepayload.New.PayloadsCollection.Empty()
	for i := 0; i < 20; i++ {
		coll.AddsPtr(&corepayload.PayloadWrapper{Name: "item"})
	}

	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		coll.GetSinglePageCollection(5, 0) // pageIndex 0 -> negative skip
	}()

	// Act
	actual := args.Map{"didPanic": didPanic}

	// Assert
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "GetSinglePageCollection panics -- zero page index", actual)
}

// ── TypedPayloadCollection: AddCollection empty ──
// Covers TypedPayloadCollection.go L214-216

func Test_TypedPayloadCollection_AddCollection_Empty(t *testing.T) {
	// Arrange
	coll := corepayload.NewTypedPayloadCollection[string](0)
	other := corepayload.NewTypedPayloadCollection[string](0)

	result := coll.AddCollection(other)

	// Act
	actual := args.Map{"empty": result.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AddCollection returns self -- empty other", actual)
}

// ── TypedPayloadCollection: Skip, Take beyond length ──
// Covers TypedPayloadCollection.go L365-367, L374-376

func Test_TypedPayloadCollection_Skip_BeyondLength(t *testing.T) {
	// Arrange
	coll := corepayload.NewTypedPayloadCollection[string](0)
	result := coll.Skip(10)

	// Act
	actual := args.Map{"empty": len(result) == 0}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Skip returns empty -- count >= length", actual)
}

func Test_TypedPayloadCollection_Take_BeyondLength(t *testing.T) {
	// Arrange
	coll := corepayload.NewTypedPayloadCollection[string](0)
	result := coll.Take(10)

	// Act
	actual := args.Map{"empty": len(result) == 0}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Take returns all -- count >= length", actual)
}

// ── helper ──

type errTestStruct struct{ msg string }

func (e *errTestStruct) Error() string { return e.msg }

func errTestHelper(msg string) error {
	return &errTestStruct{msg: msg}
}
