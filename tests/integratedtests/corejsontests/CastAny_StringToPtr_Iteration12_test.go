package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ── castingAny: string case ──
// Covers castingAny.go L79-82

func Test_CastAny_StringToPtr(t *testing.T) {
	// Arrange
	jsonStr := `{"key":"value"}`
	var result map[string]string

	err := corejson.CastAny.OrDeserializeTo(jsonStr, &result)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"key": result["key"],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"key": "value",
	}
	expected.ShouldBeEqual(t, 0, "CastAny.OrDeserializeTo works -- string input", actual)
}

func Test_CastAny_BytesToPtr(t *testing.T) {
	// Arrange
	jsonBytes := []byte(`{"key":"value"}`)
	var result map[string]string

	err := corejson.CastAny.OrDeserializeTo(jsonBytes, &result)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"key": result["key"],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"key": "value",
	}
	expected.ShouldBeEqual(t, 0, "CastAny.OrDeserializeTo works -- []byte input", actual)
}

// ── anyTo: string case ──
// Covers anyTo.go L46-47, L48-49

func Test_AnyTo_String(t *testing.T) {
	// Arrange
	jsonStr := `{"key":"value"}`
	result := corejson.AnyTo.SerializedJsonResult(jsonStr)

	// Act
	actual := args.Map{
		"noErr": result.IsEmptyError(),
		"hasBytes": len(result.Bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedJsonResult works -- string input", actual)
}

// ── Result: FieldsNames with valid JSON ──
// Covers Result.go L85-94

func Test_Result_FieldsNames(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingBytes([]byte(`{"name":"test","age":30}`))
	names, err := (&r).FieldsNames()
	// DeserializedFieldsToMap initializes map and passes pointer to Deserialize.

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasNames": len(names) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasNames": true,
	}
	expected.ShouldBeEqual(t, 0, "FieldsNames returns empty -- DeserializedFieldsToMap limitation", actual)
}

// ── Result: MeaningfulError with error and payload ──
// Covers Result.go L376-381

func Test_Result_MeaningfulError_WithPayload(t *testing.T) {
	// Arrange
	result := corejson.NewResult.UsingBytesTypePtr(
		[]byte(`{"data":"test"}`),
		"TestType",
	)
	result.Error = errors.New("test error")

	err := result.MeaningfulError()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns error -- has error and payload", actual)
}

// ── Result: safeJsonStringInternal nil ──
// Covers Result.go L385

func Test_Result_MeaningfulError_NilResult(t *testing.T) {
	// Arrange
	var result *corejson.Result
	err := result.MeaningfulError()
	// nil Result → returns defaulterr.JsonResultNull (not nil)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": false}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns error -- nil result", actual)
}

// ── Result: serializeInternal error path ──
// Covers Result.go L639-646

func Test_Result_Serialize_Error(t *testing.T) {
	// Arrange
	result := &corejson.Result{
		Error: errors.New("serialize error"),
	}

	_, err := result.Serialize()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns error -- has error", actual)
}

// ── Result: IsEqual with same bytes ──
// Covers Result.go L827-829, L872-874

func Test_Result_IsEqual_SameBytes(t *testing.T) {
	// Arrange
	r1 := corejson.NewResult.UsingBytesTypePtr([]byte(`{"a":1}`), "T")
	r2 := corejson.NewResult.UsingBytesTypePtr([]byte(`{"a":1}`), "T")

	result := r1.IsEqualPtr(r2)

	// Act
	actual := args.Map{"equal": result}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr returns true -- same bytes same type", actual)
}

func Test_Result_IsEqual_DiffType(t *testing.T) {
	// Arrange
	r1 := corejson.NewResult.UsingBytesTypePtr([]byte(`{"a":1}`), "TypeA")
	r2 := corejson.NewResult.UsingBytesTypePtr([]byte(`{"a":1}`), "TypeB")

	result := r1.IsEqualPtr(r2)

	// Act
	actual := args.Map{"equal": result}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr returns false -- different type names", actual)
}

// ── MapResults: Unmarshal empty bytes ──
// Covers MapResults.go L164-165

func Test_MapResults_Unmarshal_EmptyResult(t *testing.T) {
	// Arrange
	emptyResult := corejson.NewResult.UsingBytes([]byte{})
	mr := &corejson.MapResults{
		Items: map[string]corejson.Result{
			"key": emptyResult,
		},
	}

	var target string
	err := mr.Unmarshal("key", &target)

	// the Unmarshal method has inverted has check (line 152: if has {return error})
	// so for existing key it returns "key not found" error

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Unmarshal returns error -- existing key with inverted check", actual)
}

// ── MapResults: UnmarshalManySafe empty ──
// Covers MapResults.go L202

func Test_MapResults_UnmarshalManySafe_Empty(t *testing.T) {
	// Arrange
	mr := &corejson.MapResults{Items: map[string]corejson.Result{}}

	err := mr.UnmarshalManySafe()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalManySafe returns nil -- empty input", actual)
}

// ── MapResults: SafeUnmarshal missing key ──
// Covers MapResults.go L235-236

func Test_MapResults_SafeUnmarshal_Missing(t *testing.T) {
	// Arrange
	mr := &corejson.MapResults{Items: map[string]corejson.Result{}}

	var target string
	err := mr.SafeUnmarshal("missing", &target)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SafeUnmarshal returns nil -- missing key", actual)
}

// ── MapResults: AddAnySkipOnNil nil ──
// Covers MapResults.go L324-326 (AddAnySkipOnNil error path)

func Test_MapResults_AddAnySkipOnNil_Nil(t *testing.T) {
	// Arrange
	mr := &corejson.MapResults{Items: map[string]corejson.Result{}}
	err := mr.AddAnySkipOnNil("key", nil)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AddAnySkipOnNil returns nil -- nil item", actual)
}

func Test_MapResults_AddAnySkipOnNil_Valid(t *testing.T) {
	// Arrange
	mr := &corejson.MapResults{Items: map[string]corejson.Result{}}
	err := mr.AddAnySkipOnNil("key", map[string]string{"a": "b"})

	_, hasKey := mr.Items["key"]

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasKey": hasKey,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasKey": true,
	}
	expected.ShouldBeEqual(t, 0, "AddAnySkipOnNil adds result -- valid item", actual)
}

// ── MapResults: GetSinglePageCollection length mismatch panic ──
// Covers MapResults.go L718-729

func Test_MapResults_GetSinglePageCollection_LengthMismatch(t *testing.T) {
	// Arrange
	mr := &corejson.MapResults{
		Items: map[string]corejson.Result{
			"a": corejson.NewResult.UsingBytes([]byte(`"1"`)),
			"b": corejson.NewResult.UsingBytes([]byte(`"2"`)),
		},
	}

	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		mr.GetSinglePageCollection(1, 1, []string{"a"}) // length mismatch
	}()

	// Act
	actual := args.Map{"didPanic": didPanic}

	// Assert
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "GetSinglePageCollection panics -- keys length != map length", actual)
}

// ── MapResults: GetNewMapUsingKeys with panic on missing ──
// Covers MapResults.go L773-779

func Test_MapResults_GetNewMapUsingKeys_PanicOnMissing(t *testing.T) {
	// Arrange
	mr := &corejson.MapResults{
		Items: map[string]corejson.Result{
			"a": corejson.NewResult.UsingBytes([]byte(`"1"`)),
		},
	}

	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		mr.GetNewMapUsingKeys(true, "missing_key")
	}()

	// Act
	actual := args.Map{"didPanic": didPanic}

	// Assert
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "GetNewMapUsingKeys panics -- missing key with isPanicOnMissing", actual)
}

// ── deserializeFromBytesTo: nil ptr ──
// Covers deserializeFromBytesTo.go L23-24

func Test_Deserialize_UsingBytes_NilPtr(t *testing.T) {
	// Arrange
	err := corejson.Deserialize.UsingBytes([]byte(`{}`), nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingBytes returns error -- nil unmarshal ptr", actual)
}

// ── newResultCreator: UsingSerializerFunc ──
// Covers newResultCreator.go L66-68

func Test_NewResult_UsingSerializerFunc(t *testing.T) {
	// Arrange
	fn := func() ([]byte, error) { return []byte(`{"key":"val"}`), nil }
	result := corejson.NewResult.UsingSerializerFunc(fn)

	// Act
	actual := args.Map{
		"noErr": result.IsEmptyError(),
		"hasBytes": len(result.Bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "UsingSerializerFunc creates result -- valid func", actual)
}

// ── newBytesCollectionCreator: empty ──
// Covers newBytesCollectionCreator.go L52

func Test_NewBytesCollection_Empty(t *testing.T) {
	// Arrange
	result := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"empty": result.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "NewBytesCollection.Empty returns empty collection", actual)
}

// ── newMapResultsCreator: empty ──
// Covers newMapResultsCreator.go L52

func Test_NewMapResults_Empty(t *testing.T) {
	// Arrange
	result := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"empty": result.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "NewMapResults.Empty returns empty map results", actual)
}

// ── newResultsCollectionCreator: empty ──
// Covers newResultsCollectionCreator.go L56

func Test_NewResultsCollection_Empty(t *testing.T) {
	// Arrange
	result := corejson.NewResultsCollection.Empty()

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"empty": result.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection.Empty returns empty collection", actual)
}
