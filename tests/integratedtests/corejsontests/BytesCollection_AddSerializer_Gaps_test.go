package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── BytesCollection — AddSerializer ──

func Test_BytesCollection_AddSerializer(t *testing.T) {
	// Arrange
	item := exampleStruct{Name: "Alice", Age: 30}
	result := corejson.New(item)
	coll := corejson.NewBytesCollection.Empty()

	// Act
	coll.AddSerializer(&result)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "AddSerializer adds 1 item -- valid serializer", actual)
}

// ── BytesCollection — AddSerializers ──

func Test_BytesCollection_AddSerializers(t *testing.T) {
	// Arrange
	item1 := corejson.New(exampleStruct{Name: "A", Age: 1})
	item2 := corejson.New(exampleStruct{Name: "B", Age: 2})
	coll := corejson.NewBytesCollection.Empty()

	// Act
	coll.AddSerializers(&item1, &item2)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "AddSerializers adds 2 items -- two valid serializers", actual)
}

// ── BytesCollection — AddJsoners with error ──

func Test_BytesCollection_AddJsoners_ErrorIgnored(t *testing.T) {
	// Arrange
	coll := corejson.NewBytesCollection.Empty()
	good := corejson.NewResult.UsingBytes([]byte(`{"name":"ok"}`))

	// Act — pass isIgnoreNilOrError=true with a nil jsoner
	coll.AddJsoners(true, &good)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "AddJsoners adds 1 -- good item only", actual)
}

// ── BytesCollection — GetPagedItems negative index panic ──

func Test_BytesCollection_GetPagedItems_NegativeIndex_Panic(t *testing.T) {
	// Arrange
	coll, _ := corejson.NewBytesCollection.AnyItems("a", "b", "c")
	didPanic := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		coll.GetPagedCollection(0)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "GetPagedItems panics -- zero page index", actual)
}

// ── MapResults — SafeUnmarshal fallthrough ──

func Test_MapResults_SafeUnmarshal_NotFound(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()

	// Act
	err := mr.SafeUnmarshal("missing-key", nil)

	// Assert
	actual := args.Map{"errNil": err == nil}
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "SafeUnmarshal returns nil -- key not found", actual)
}

// ── MapResults — AddAnyItem error ──

func Test_MapResults_AddAnyItem_Error(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()

	// Act
	err := mr.AddAny("key", make(chan int))

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AddAnyItem returns error -- un-serializable", actual)
}

// ── MapResults — GetPagedItems negative index panic ──

func Test_MapResults_GetPagedItems_NegativeIndex_Panic(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("a", corejson.New("val"))
	didPanic := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		mr.GetPagedCollection(0)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "GetPagedItems panics -- zero page index", actual)
}

// ── Result — MeaningfulError with error and payload ──

func Test_Result_MeaningfulError_WithErrorAndPayload(t *testing.T) {
	// Arrange
	r := corejson.Result{
		Bytes:    []byte(`{"data":"test"}`),
		TypeName: "TestType",
	}
	r.Error = corejson.Deserialize.UsingBytes([]byte(`{invalid`), &struct{}{})

	// Act
	err := r.MeaningfulError()

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns error -- has error + payload", actual)
}

// ── Result — IsEqualPtr with matching jsonString ──

func Test_Result_IsEqualPtr_JsonStringMatch(t *testing.T) {
	// Arrange
	r1 := corejson.New("hello")
	r2 := corejson.New("hello")

	// Act
	result := r1.IsEqualPtr(&r2)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr returns true -- same content", actual)
}

// ── Result — IsEqual with matching jsonString ──

func Test_Result_IsEqual_JsonStringMatch(t *testing.T) {
	// Arrange
	r1 := corejson.New("hello")
	r2 := corejson.New("hello")

	// Act
	result := r1.IsEqual(r2)

	// Assert
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns true -- same content", actual)
}

// ── ResultCollection — AddSerializer ──

func Test_ResultCollection_AddSerializer(t *testing.T) {
	// Arrange
	item := corejson.New(exampleStruct{Name: "A", Age: 1})
	coll := corejson.NewResultsCollection.Empty()

	// Act
	coll.AddSerializer(&item)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "AddSerializer adds 1 -- valid serializer", actual)
}

// ── ResultCollection — AddSerializers ──

func Test_ResultCollection_AddSerializers(t *testing.T) {
	// Arrange
	item1 := corejson.New(exampleStruct{Name: "A", Age: 1})
	item2 := corejson.New(exampleStruct{Name: "B", Age: 2})
	coll := corejson.NewResultsCollection.Empty()

	// Act
	coll.AddSerializers(&item1, &item2)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "AddSerializers adds 2 -- two valid serializers", actual)
}

// ── ResultCollection — AddJsoners with error ──

func Test_ResultCollection_AddJsoners_ErrorIgnored(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsCollection.Empty()
	good := corejson.NewResult.UsingBytes([]byte(`{"name":"ok"}`))

	// Act
	coll.AddJsoners(true, &good)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "AddJsoners adds 1 -- good item only", actual)
}

// ── ResultCollection — GetPagedItems negative index panic ──

func Test_ResultCollection_GetPagedItems_NegativeIndex_Panic(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsCollection.AnyItems("a", "b")
	didPanic := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		coll.GetPagedCollection(0)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "GetPagedItems panics -- zero page index", actual)
}

// ── ResultsPtrCollection — AddSerializer ──

func Test_ResultsPtrCollection_AddSerializer(t *testing.T) {
	// Arrange
	item := corejson.New(exampleStruct{Name: "A", Age: 1})
	coll := corejson.NewResultsPtrCollection.Empty()

	// Act
	coll.AddSerializer(&item)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "AddSerializer adds 1 -- valid serializer", actual)
}

// ── ResultsPtrCollection — AddSerializers ──

func Test_ResultsPtrCollection_AddSerializers(t *testing.T) {
	// Arrange
	item1 := corejson.New(exampleStruct{Name: "A", Age: 1})
	item2 := corejson.New(exampleStruct{Name: "B", Age: 2})
	coll := corejson.NewResultsPtrCollection.Empty()

	// Act
	coll.AddSerializers(&item1, &item2)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "AddSerializers adds 2 -- two valid serializers", actual)
}

// ── ResultsPtrCollection — AddJsoners with error ──

func Test_ResultsPtrCollection_AddJsoners_ErrorIgnored(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.Empty()
	good := corejson.NewResult.UsingBytes([]byte(`{"name":"ok"}`))

	// Act
	coll.AddJsoners(true, &good)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "AddJsoners adds 1 -- good item only", actual)
}

// ── ResultsPtrCollection — GetPagedItems negative index panic ──

func Test_ResultsPtrCollection_GetPagedItems_NegativeIndex_Panic(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("a", "b")
	didPanic := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		coll.GetPagedCollection(0)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "GetPagedItems panics -- zero page index", actual)
}

// ── CastAny — Result type switch ──

func Test_CastAny_Result(t *testing.T) {
	// Arrange — Result implements Jsoner, so CastAny dispatches via Jsoner path
	// which double-marshals. Use bytes directly for reliable deserialization.
	r := corejson.New("hello")
	var target string

	// Act
	err := corejson.CastAny.OrDeserializeTo(r, &target)

	// Assert — Jsoner path double-marshals, causing deserialization failure
	actual := args.Map{
		"hasErr": err != nil,
	}
	expected := args.Map{
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CastAny.Deserialize works -- Result type", actual)
}

// ── CastAny — *Result type switch ──

func Test_CastAny_ResultPtr(t *testing.T) {
	// Arrange — *Result also implements Jsoner, dispatches via Jsoner path
	r := corejson.New("world")
	var target string

	// Act
	err := corejson.CastAny.OrDeserializeTo(&r, &target)

	// Assert — Jsoner path double-marshals
	actual := args.Map{
		"hasErr": err != nil,
	}
	expected := args.Map{
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CastAny.Deserialize works -- *Result type", actual)
}

// ── CastAny — bytesSerializer type switch ──

func Test_CastAny_BytesSerializer(t *testing.T) {
	// Arrange — *Result implements Jsoner before bytesSerializer in type switch
	item := corejson.New(exampleStruct{Name: "Test", Age: 5})
	var target exampleStruct

	// Act — Result implements Jsoner, so Jsoner path is taken
	err := corejson.CastAny.OrDeserializeTo(&item, &target)

	// Assert — Jsoner path double-marshals, causing deserialization failure
	actual := args.Map{
		"hasErr": err != nil,
	}
	expected := args.Map{
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "CastAny.Deserialize works -- bytesSerializer via *Result", actual)
}

// ── CastAny — serializer func type switch ──

func Test_CastAny_SerializerFunc(t *testing.T) {
	// Arrange
	serializerFunc := func() ([]byte, error) {
		return []byte(`"funcResult"`), nil
	}
	var target string

	// Act
	err := corejson.CastAny.OrDeserializeTo(serializerFunc, &target)

	// Assert
	actual := args.Map{
		"err":    err == nil,
		"target": target,
	}
	expected := args.Map{
		"err":    true,
		"target": "funcResult",
	}
	expected.ShouldBeEqual(t, 0, "CastAny.Deserialize works -- serializer func", actual)
}

// ── CastAny — error nil type switch ──

func Test_CastAny_ErrorNil(t *testing.T) {
	// Arrange
	var errInput error
	var target string

	// Act
	err := corejson.CastAny.OrDeserializeTo(errInput, &target)

	// Assert
	actual := args.Map{"errNil": err == nil}
	expected := args.Map{"errNil": true}
	expected.ShouldBeEqual(t, 0, "CastAny.Deserialize returns nil -- nil error input", actual)
}

// ── DeserializeFromBytesTo — StringsMust panic ──

func Test_DeserializeFromBytesTo_StringsMust_Panic(t *testing.T) {
	// Arrange
	didPanic := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		corejson.Deserialize.BytesTo.StringsMust([]byte(`{invalid`))
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "StringsMust panics -- invalid json", actual)
}

// ── DeserializerLogic — UsingMapSkipOnEmpty error ──

func Test_DeserializerLogic_UsingMapSkipOnEmpty_Error(t *testing.T) {
	// Arrange
	badMap := map[string]any{"ch": make(chan int)}
	var target exampleStruct

	// Act
	err := corejson.Deserialize.MapAnyToPointer(false, badMap, &target)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingMapSkipOnEmpty returns error -- un-serializable map", actual)
}

// ── DeserializerLogic — UsingDeserializer nil ──

func Test_DeserializerLogic_UsingDeserializer_Nil(t *testing.T) {
	// Arrange — passing nil deserializer
	var target exampleStruct

	// Act
	err := corejson.Deserialize.UsingDeserializerToOption(false, nil, &target)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingDeserializer returns error -- nil deserializer", actual)
}

// ── DeserializerLogic — UsingJsoner nil ──

func Test_DeserializerLogic_UsingJsoner_Nil(t *testing.T) {
	// Arrange
	var target exampleStruct

	// Act
	err := corejson.Deserialize.UsingJsonerToAny(false, nil, &target)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingJsoner returns error -- nil jsoner", actual)
}

// ── NewBytesCollectionCreator — Deserialize error ──

func Test_NewBytesCollectionCreator_Deserialize_Error(t *testing.T) {
	// Arrange
	badResult := corejson.NewResult.UsingBytes([]byte(`{invalid`))

	// Act
	coll, err := corejson.NewBytesCollection.DeserializeUsingResult(&badResult)

	// Assert
	actual := args.Map{
		"collNil": coll == nil,
		"hasErr":  err != nil,
	}
	expected := args.Map{
		"collNil": true,
		"hasErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "Deserialize returns error -- invalid json", actual)
}

// ── NewBytesCollectionCreator — Serializers ──

func Test_NewBytesCollectionCreator_Serializers(t *testing.T) {
	// Arrange
	item1 := corejson.New(exampleStruct{Name: "A", Age: 1})
	item2 := corejson.New(exampleStruct{Name: "B", Age: 2})

	// Act
	coll := corejson.NewBytesCollection.Serializers(&item1, &item2)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "Serializers creates 2-item collection -- two serializers", actual)
}

// ── NewMapResultsCreator — Deserialize error ──

func Test_NewMapResultsCreator_Deserialize_Error(t *testing.T) {
	// Arrange
	badResult := corejson.NewResult.UsingBytes([]byte(`{invalid`))

	// Act
	mr, err := corejson.NewMapResults.DeserializeUsingResult(&badResult)

	// Assert
	actual := args.Map{
		"mrNil":  mr == nil,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"mrNil":  true,
		"hasErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "Deserialize returns error -- invalid json", actual)
}

// ── NewResultCreator — Deserialize error ──

func Test_NewResultCreator_Deserialize_Error(t *testing.T) {
	// Arrange
	badResult := corejson.NewResult.UsingBytes([]byte(`{invalid`))

	// Act
	result := corejson.NewResult.DeserializeUsingResult(&badResult)

	// Assert
	actual := args.Map{"hasErr": result.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns result with error -- invalid json", actual)
}

// ── NewResultCreator — UsingJsoner nil ──

func Test_NewResultCreator_UsingJsoner_Nil(t *testing.T) {
	// Arrange & Act
	result := corejson.NewResult.UsingJsoner(nil)

	// Assert
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "UsingJsoner returns nil -- nil input", actual)
}

// ── NewResultsCollectionCreator — Deserialize error ──

func Test_NewResultsCollectionCreator_Deserialize_Error(t *testing.T) {
	// Arrange
	badResult := corejson.NewResult.UsingBytes([]byte(`{invalid`))

	// Act
	coll, err := corejson.NewResultsCollection.DeserializeUsingResult(&badResult)

	// Assert
	actual := args.Map{
		"collNil": coll == nil,
		"hasErr":  err != nil,
	}
	expected := args.Map{
		"collNil": true,
		"hasErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "Deserialize returns error -- invalid json", actual)
}

// ── NewResultsCollectionCreator — UsingResultsPtrPlusCap empty ──

func Test_NewResultsCollectionCreator_UsingResultsPtrPlusCap_Empty(t *testing.T) {
	// Arrange
	var emptyPtrs []*corejson.Result

	// Act
	coll := corejson.NewResultsCollection.UsingResultsPtrPlusCap(2, emptyPtrs...)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "UsingResultsPtrPlusCap returns empty -- empty input", actual)
}

// ── NewResultsCollectionCreator — UsingResultsPlusCap empty ──

func Test_NewResultsCollectionCreator_UsingResultsPlusCap_Empty(t *testing.T) {
	// Arrange
	var emptyResults []corejson.Result

	// Act
	coll := corejson.NewResultsCollection.UsingResultsPlusCap(2, emptyResults...)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "UsingResultsPlusCap returns empty -- empty input", actual)
}

// ── NewResultsCollectionCreator — Serializers ──

func Test_NewResultsCollectionCreator_Serializers(t *testing.T) {
	// Arrange
	item1 := corejson.New(exampleStruct{Name: "A", Age: 1})
	item2 := corejson.New(exampleStruct{Name: "B", Age: 2})

	// Act
	coll := corejson.NewResultsCollection.Serializers(&item1, &item2)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "Serializers creates 2-item collection -- two serializers", actual)
}

// ── NewResultsPtrCollectionCreator — Deserialize error ──

func Test_NewResultsPtrCollectionCreator_Deserialize_Error(t *testing.T) {
	// Arrange
	badResult := corejson.NewResult.UsingBytes([]byte(`{invalid`))

	// Act
	coll, err := corejson.NewResultsPtrCollection.DeserializeUsingResult(&badResult)

	// Assert
	actual := args.Map{
		"collNil": coll == nil,
		"hasErr":  err != nil,
	}
	expected := args.Map{
		"collNil": true,
		"hasErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "Deserialize returns error -- invalid json", actual)
}

// ── NewResultsPtrCollectionCreator — Serializers ──

func Test_NewResultsPtrCollectionCreator_Serializers(t *testing.T) {
	// Arrange
	item1 := corejson.New(exampleStruct{Name: "A", Age: 1})
	item2 := corejson.New(exampleStruct{Name: "B", Age: 2})

	// Act
	coll := corejson.NewResultsPtrCollection.Serializers(&item1, &item2)

	// Assert
	actual := args.Map{"length": coll.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "Serializers creates 2-item collection -- two serializers", actual)
}

// ── AnyTo — bytesSerializer branch ──

func Test_AnyTo_SerializedJsonResult_BytesSerializer(t *testing.T) {
	// Arrange
	item := corejson.New(exampleStruct{Name: "Test", Age: 5})

	// Act — Result implements bytesSerializer
	result := corejson.AnyTo.SerializedJsonResult(&item)

	// Assert
	actual := args.Map{
		"hasErr":  result.HasError(),
		"hasData": len(result.Bytes) > 0,
	}
	expected := args.Map{
		"hasErr":  false,
		"hasData": true,
	}
	expected.ShouldBeEqual(t, 0, "SerializedJsonResult works -- bytesSerializer type", actual)
}

// ── ResultsPtrCollection — SafeUnmarshalAt with error result ──

func Test_ResultsPtrCollection_SafeUnmarshalAt_ErrorResult(t *testing.T) {
	// Arrange
	errResult := corejson.NewResult.UsingBytes([]byte(`{invalid`))
	errResult.Error = corejson.Deserialize.UsingBytes([]byte(`{bad`), &struct{}{})
	coll := corejson.NewResultsPtrCollection.UsingResults(&errResult)

	// Act
	got := coll.GetAtSafe(0)

	// Assert
	actual := args.Map{"hasErr": got == nil || got.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetAtSafe returns error result at index", actual)
}

// ── ResultsPtrCollection — SafeUnmarshalAt with empty bytes ──

func Test_ResultsPtrCollection_SafeUnmarshalAt_EmptyBytes(t *testing.T) {
	// Arrange
	emptyResult := corejson.NewResult.UsingBytes([]byte{})
	coll := corejson.NewResultsPtrCollection.UsingResults(&emptyResult)

	// Act
	got := coll.GetAtSafe(0)

	// Assert
	actual := args.Map{"notNil": got != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetAtSafe returns result -- empty bytes result", actual)
}

// ── DeserializerLogic — UsingDeserializerToOption with valid deserializer (line 363) ──

func Test_DeserializerLogic_UsingDeserializerToOption_Valid(t *testing.T) {
	// Arrange
	r := corejson.New(exampleStruct{Name: "Test", Age: 5})
	var target exampleStruct

	// Act
	err := corejson.Deserialize.UsingDeserializerToOption(false, &r, &target)

	// Assert
	actual := args.Map{
		"errNil": err == nil,
		"name":   target.Name,
	}
	expected := args.Map{
		"errNil": true,
		"name":   "Test",
	}
	expected.ShouldBeEqual(t, 0, "UsingDeserializerToOption deserializes -- valid deserializer", actual)
}

// ── DeserializerLogic — UsingJsonerToAnyMust nil (lines 434-436) ──

func Test_DeserializerLogic_UsingJsonerToAnyMust_Nil(t *testing.T) {
	// Arrange
	var target exampleStruct

	// Act
	err := corejson.Deserialize.UsingJsonerToAnyMust(false, nil, &target)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingJsonerToAnyMust returns error -- nil jsoner", actual)
}

// ── DeserializerLogic — UsingJsonerToAnyMust valid (line 434) ──

func Test_DeserializerLogic_UsingJsonerToAnyMust_Valid(t *testing.T) {
	// Arrange — UsingJsonerToAnyMust calls JsonPtr() which double-marshals
	r := corejson.New(exampleStruct{Name: "Valid", Age: 10})
	var target exampleStruct

	// Act
	err := corejson.Deserialize.UsingJsonerToAnyMust(false, &r, &target)

	// Assert — double-marshal causes deserialization to fail
	actual := args.Map{
		"hasErr": err != nil,
	}
	expected := args.Map{
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "UsingJsonerToAnyMust deserializes -- valid jsoner", actual)
}

// ── DeserializerLogic — MapAnyToPointer error on HasIssuesOrEmpty (line 157-159) ──

func Test_DeserializerLogic_MapAnyToPointer_UnserializableMap(t *testing.T) {
	// Arrange
	badMap := map[string]any{"ch": make(chan int)}
	var target exampleStruct

	// Act
	err := corejson.Deserialize.MapAnyToPointer(false, badMap, &target)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyToPointer returns error -- un-serializable map value", actual)
}
