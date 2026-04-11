package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Result ──

func Test_Result_Basic(t *testing.T) {
	// Arrange
	result := corejson.New(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"hasErr":    result.HasError(),
		"isEmpty":   result.IsEmpty(),
		"hasBytes":  len(result.Bytes) > 0,
		"typeName":  result.TypeName != "",
		"jsonStr":   result.JsonString() != "",
		"prettyStr": result.PrettyJsonString() != "",
	}

	// Assert
	expected := args.Map{
		"hasErr": false, "isEmpty": false, "hasBytes": true,
		"typeName": true, "jsonStr": true, "prettyStr": true,
	}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- basic", actual)
}

func Test_Result_Ptr(t *testing.T) {
	// Arrange
	result := corejson.NewPtr(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"noErr": !result.HasError(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "NewPtr returns correct value -- with args", actual)
}

func Test_Result_SafeBytes_FromResultBasic(t *testing.T) {
	// Arrange
	result := corejson.New(map[string]int{"a": 1})

	// Act
	actual := args.Map{"hasBytes": len(result.SafeBytes()) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "SafeBytes returns correct value -- with args", actual)
}

func Test_Result_Unmarshal_FromResultBasic(t *testing.T) {
	// Arrange
	result := corejson.New(map[string]int{"a": 1})
	var target map[string]int
	err := result.Unmarshal(&target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"a": target["a"],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"a": 1,
	}
	expected.ShouldBeEqual(t, 0, "Unmarshal returns correct value -- with args", actual)
}

func Test_Result_Deserialize_FromResultBasic(t *testing.T) {
	// Arrange
	result := corejson.New(map[string]int{"b": 2})
	var target map[string]int
	err := result.Deserialize(&target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"b": target["b"],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"b": 2,
	}
	expected.ShouldBeEqual(t, 0, "Deserialize returns correct value -- with args", actual)
}

func Test_Result_MeaningfulError_NoError(t *testing.T) {
	// Arrange
	result := corejson.New("hello")
	err := result.MeaningfulError()

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns nil -- nil", actual)
}

// ── Serialize ──

func Test_Serialize_Raw(t *testing.T) {
	// Arrange
	bytes, err := corejson.Serialize.Raw(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize.Raw returns correct value -- with args", actual)
}

func Test_Serialize_UsingAny(t *testing.T) {
	// Arrange
	result := corejson.Serialize.UsingAny(map[string]int{"a": 1})

	// Act
	actual := args.Map{"noErr": !result.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize.UsingAny returns correct value -- with args", actual)
}

// ── Deserialize ──

func Test_Deserialize_UsingBytes(t *testing.T) {
	// Arrange
	var target map[string]int
	err := corejson.Deserialize.UsingBytes([]byte(`{"a":1}`), &target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"a": target["a"],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"a": 1,
	}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingBytes returns correct value -- with args", actual)
}

func Test_Deserialize_UsingBytes_Invalid_FromResultBasic(t *testing.T) {
	// Arrange
	var target map[string]int
	err := corejson.Deserialize.UsingBytes([]byte("invalid"), &target)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingBytes returns error -- invalid", actual)
}

func Test_Deserialize_UsingResult_FromResultBasic(t *testing.T) {
	// Arrange
	result := corejson.New(map[string]int{"a": 1})
	var target map[string]int
	err := corejson.Deserialize.UsingResult(&result, &target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"a": target["a"],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"a": 1,
	}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingResult returns correct value -- with args", actual)
}

func Test_Deserialize_Apply_FromResultBasic(t *testing.T) {
	// Arrange
	result := corejson.New(map[string]int{"a": 1})
	var target map[string]int
	err := corejson.Deserialize.Apply(&result, &target)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.Apply returns correct value -- with args", actual)
}

// ── BytesDeepClone ──

func Test_BytesDeepClone_FromResultBasic(t *testing.T) {
	// Arrange
	original := []byte("hello")
	cloned := corejson.BytesDeepClone(original)
	original[0] = 'X'

	// Act
	actual := args.Map{"different": string(cloned) == "hello"}

	// Assert
	expected := args.Map{"different": true}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone returns correct value -- with args", actual)
}

func Test_BytesDeepClone_Nil(t *testing.T) {
	// Arrange
	cloned := corejson.BytesDeepClone(nil)

	// Act
	actual := args.Map{"isEmpty": len(cloned) == 0}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone returns nil -- nil", actual)
}

// ── BytesToString ──

func Test_BytesToString_FromResultBasic(t *testing.T) {
	// Arrange
	result := corejson.BytesToString([]byte("hello"))

	// Act
	actual := args.Map{"str": result}

	// Assert
	expected := args.Map{"str": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesToString returns correct value -- with args", actual)
}

func Test_BytesToPrettyString_FromResultBasic(t *testing.T) {
	// Arrange
	result := corejson.BytesToPrettyString([]byte(`{"a":1}`))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BytesToPrettyString returns correct value -- with args", actual)
}

// ── Empty ──

func Test_Empty_ResultPtr_FromResultBasic(t *testing.T) {
	// Arrange
	result := corejson.Empty.ResultPtr()

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"isEmpty": result.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Empty.ResultPtr returns empty -- with args", actual)
}

func Test_Empty_Result_FromResultBasic(t *testing.T) {
	// Arrange
	result := corejson.Empty.Result()

	// Act
	actual := args.Map{"isEmpty": result.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Empty.Result returns empty -- with args", actual)
}

// ── NewResult ──

func Test_NewResult_UsingTypeBytesPtr(t *testing.T) {
	// Arrange
	result := corejson.NewResult.UsingTypeBytesPtr("TestType", []byte(`{"a":1}`))

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"noErr": !result.HasError(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "NewResult.UsingTypeBytesPtr returns correct value -- with args", actual)
}

// ── BytesCloneIf ──

func Test_BytesCloneIf_True_FromResultBasic(t *testing.T) {
	// Arrange
	original := []byte("data")
	cloned := corejson.BytesCloneIf(true, original)
	original[0] = 'X'

	// Act
	actual := args.Map{"different": string(cloned) == "data"}

	// Assert
	expected := args.Map{"different": true}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf returns non-empty -- true", actual)
}

func Test_BytesCloneIf_False_FromResultBasic(t *testing.T) {
	// Arrange
	cloned := corejson.BytesCloneIf(false, []byte("data"))

	// Act
	actual := args.Map{"isEmpty": len(cloned) == 0}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf returns non-empty -- false", actual)
}

// ── AnyTo ──

func Test_AnyTo_SerializedJsonResult_FromResultBasic(t *testing.T) {
	// Arrange
	result := corejson.AnyTo.SerializedJsonResult(map[string]int{"a": 1})

	// Act
	actual := args.Map{"noErr": !result.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedJsonResult returns correct value -- with args", actual)
}

func Test_AnyTo_SerializedJsonResult_Bytes_FromResultBasic(t *testing.T) {
	// Arrange
	result := corejson.AnyTo.SerializedJsonResult([]byte(`{"a":1}`))

	// Act
	actual := args.Map{"noErr": !result.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedJsonResult returns correct value -- bytes", actual)
}

func Test_AnyTo_SerializedJsonResult_String_FromResultBasic(t *testing.T) {
	// Arrange
	result := corejson.AnyTo.SerializedJsonResult(`{"a":1}`)

	// Act
	actual := args.Map{"noErr": !result.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedJsonResult returns correct value -- string", actual)
}

// ── CastAny ──

func Test_CastAny_FromToDefault_FromResultBasic(t *testing.T) {
	// Arrange
	source := map[string]int{"a": 1}
	var target map[string]int
	err := corejson.CastAny.FromToDefault(source, &target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"a": target["a"],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"a": 1,
	}
	expected.ShouldBeEqual(t, 0, "CastAny.FromToDefault returns correct value -- with args", actual)
}

// ── BytesCollection ──

func Test_BytesCollection_Empty(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{
		"isEmpty": bc.IsEmpty(),
		"len": bc.Length(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "BytesCollection returns empty -- empty", actual)
}

func Test_BytesCollection_Add(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte("hello"))
	bc.Add([]byte("world"))

	// Act
	actual := args.Map{
		"len": bc.Length(),
		"hasAny": bc.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"hasAny": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesCollection returns correct value -- add", actual)
}

// ── ResultCollection ──

func Test_ResultCollection_Empty(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()

	// Act
	actual := args.Map{"isEmpty": rc.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "ResultCollection returns empty -- empty", actual)
}

func Test_ResultCollection_Add(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.Empty()
	result := corejson.New("hello")
	rc.Add(result)

	// Act
	actual := args.Map{"len": rc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultCollection returns correct value -- add", actual)
}

// ── JsonString / JsonStringer ──

func Test_JsonString_FromResultBasic(t *testing.T) {
	// Arrange
	js, err := corejson.JsonString(`{"a":1}`)

	// Act
	actual := args.Map{
		"str": js,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"str": `"{\"a\":1}"`,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct value -- with args", actual)
}

// ── MapResults ──

func Test_MapResults_Empty(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{"isEmpty": mr.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapResults returns empty -- empty", actual)
}

func Test_MapResults_Add_FromResultBasic(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	result := corejson.New("hello")
	mr.Add("key1", result)

	// Act
	actual := args.Map{
		"len": mr.Length(),
		"hasAny": mr.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"hasAny": true,
	}
	expected.ShouldBeEqual(t, 0, "MapResults returns correct value -- add", actual)
}
