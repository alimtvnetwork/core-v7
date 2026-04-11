package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// callPanics recovers from panics and returns whether one occurred
func callPanics(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}

// =============================================================================
// deserializerLogic — Apply / UsingResult / ApplyMust
// =============================================================================

func Test_Deserialize_Apply(t *testing.T) {
	tc := deserializeApplyTestCase

	// Arrange
	r := corejson.NewPtr("hello")
	var s string

	// Act
	err := corejson.Deserialize.Apply(r, &s)
	actual := args.Map{
		"hasError": err != nil,
		"result":   s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingResult(t *testing.T) {
	tc := deserializeUsingResultTestCase

	// Arrange
	r := corejson.NewPtr("hello")
	var s string

	// Act
	err := corejson.Deserialize.UsingResult(r, &s)
	actual := args.Map{
		"hasError": err != nil,
		"result":   s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_ApplyMust_Panics(t *testing.T) {
	tc := deserializeApplyMustPanicsTestCase

	// Arrange
	r := &corejson.Result{Error: invalidResultForPanic}
	var s string

	// Act
	panicked := callPanics(func() {
		corejson.Deserialize.ApplyMust(r, &s)
	})
	actual := args.Map{
		"panicked": panicked,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// deserializerLogic — UsingString / FromString / FromStringMust
// =============================================================================

func Test_Deserialize_UsingString(t *testing.T) {
	tc := deserializeUsingStringTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingString(`"hello"`, &s)
	actual := args.Map{
		"hasError": err != nil,
		"result":   s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_FromString(t *testing.T) {
	tc := deserializeFromStringTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.FromString(`"hello"`, &s)
	actual := args.Map{
		"hasError": err != nil,
		"result":   s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_FromStringMust_Panics(t *testing.T) {
	tc := deserializeFromStringMustPanicsTestCase

	// Arrange
	var s string

	// Act
	panicked := callPanics(func() {
		corejson.Deserialize.FromStringMust(`bad`, &s)
	})
	actual := args.Map{
		"panicked": panicked,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// deserializerLogic — UsingStringPtr / UsingStringOption / UsingStringIgnoreEmpty
// =============================================================================

func Test_Deserialize_UsingStringPtr_Nil(t *testing.T) {
	tc := deserializeUsingStringPtrNilTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingStringPtr(nil, &s)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingStringPtr_Valid(t *testing.T) {
	tc := deserializeUsingStringPtrValidTestCase

	// Arrange
	js := `"hello"`
	var s string

	// Act
	err := corejson.Deserialize.UsingStringPtr(&js, &s)
	actual := args.Map{
		"hasError": err != nil,
		"result":   s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingStringOption_Skip(t *testing.T) {
	tc := deserializeUsingStringOptionSkipTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingStringOption(true, "", &s)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingStringOption_Process(t *testing.T) {
	tc := deserializeUsingStringOptionProcessTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingStringOption(false, `"hello"`, &s)
	actual := args.Map{
		"hasError": err != nil,
		"result":   s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingStringIgnoreEmpty(t *testing.T) {
	tc := deserializeUsingStringIgnoreEmptyTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &s)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// deserializerLogic — UsingError / UsingBytes / UsingBytesMust
// =============================================================================

func Test_Deserialize_UsingError_Nil(t *testing.T) {
	tc := deserializeUsingErrorNilTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingError(nil, &s)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingBytes_Valid(t *testing.T) {
	tc := deserializeUsingBytesValidTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingBytes([]byte(`"hello"`), &s)
	actual := args.Map{
		"hasError": err != nil,
		"result":   s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingBytes_Invalid(t *testing.T) {
	tc := deserializeUsingBytesInvalidTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingBytes([]byte(`bad`), &s)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingBytesMust_Panics(t *testing.T) {
	tc := deserializeUsingBytesMustPanicsTestCase

	// Arrange
	var s string

	// Act
	panicked := callPanics(func() {
		corejson.Deserialize.UsingBytesMust([]byte(`bad`), &s)
	})
	actual := args.Map{
		"panicked": panicked,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// deserializerLogic — UsingBytesPointer / UsingBytesIf / UsingSafeBytesMust
// =============================================================================

func Test_Deserialize_UsingBytesPointer_Nil(t *testing.T) {
	tc := deserializeUsingBytesPointerNilTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingBytesPointer(nil, &s)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingBytesPointer_Valid(t *testing.T) {
	tc := deserializeUsingBytesPointerValidTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingBytesPointer([]byte(`"hello"`), &s)
	actual := args.Map{
		"hasError": err != nil,
		"result":   s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingBytesIf_Skip(t *testing.T) {
	tc := deserializeUsingBytesIfSkipTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingBytesIf(false, []byte(`"x"`), &s)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingBytesIf_Process(t *testing.T) {
	tc := deserializeUsingBytesIfProcessTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingBytesIf(true, []byte(`"hello"`), &s)
	actual := args.Map{
		"hasError": err != nil,
		"result":   s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingSafeBytesMust_Empty(t *testing.T) {
	tc := deserializeUsingSafeBytesMustEmptyTestCase

	// Arrange
	var s string

	// Act
	corejson.Deserialize.UsingSafeBytesMust([]byte{}, &s)
	actual := args.Map{
		"result": s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingSafeBytesMust_Valid(t *testing.T) {
	tc := deserializeUsingSafeBytesMustValidTestCase

	// Arrange
	var s string

	// Act
	corejson.Deserialize.UsingSafeBytesMust([]byte(`"hello"`), &s)
	actual := args.Map{
		"result": s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// deserializerLogic — MapAnyToPointer / FromTo
// =============================================================================

func Test_Deserialize_MapAnyToPointer_Skip(t *testing.T) {
	tc := deserializeMapAnyToPointerSkipTestCase

	// Arrange
	var s deserializeTestStruct

	// Act
	err := corejson.Deserialize.MapAnyToPointer(true, map[string]any{}, &s)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_MapAnyToPointer_Valid(t *testing.T) {
	tc := deserializeMapAnyToPointerValidTestCase

	// Arrange
	var s deserializeTestStruct

	// Act
	err := corejson.Deserialize.MapAnyToPointer(false, map[string]any{"Name": "test"}, &s)
	actual := args.Map{
		"hasError": err != nil,
		"name":     s.Name,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_FromTo(t *testing.T) {
	tc := deserializeFromToTestCase

	// Arrange
	var s string

	// Act
	// "hello" is not valid JSON — FromTo → CastAny.FromToDefault → case string → json.Unmarshal fails
	// Use quoted JSON string instead
	err := corejson.Deserialize.FromTo(`"hello"`, &s)
	actual := args.Map{
		"hasError": err != nil,
		"result":   s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// deserializerLogic — UsingDeserializerFuncDefined
// =============================================================================

func Test_Deserialize_UsingDeserializerFuncDefined_Nil(t *testing.T) {
	tc := deserializeUsingDeserializerFuncDefinedNilTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingDeserializerFuncDefined(nil, &s)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingDeserializerFuncDefined_Valid(t *testing.T) {
	tc := deserializeUsingDeserializerFuncDefinedValidTestCase

	// Arrange
	var s string
	fn := func(toPtr any) error {
		*(toPtr.(*string)) = "hello"
		return nil
	}

	// Act
	err := corejson.Deserialize.UsingDeserializerFuncDefined(fn, &s)
	actual := args.Map{
		"hasError": err != nil,
		"result":   s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// deserializerLogic — UsingJsonerToAny / UsingDeserializerToOption / UsingDeserializerDefined
// =============================================================================

func Test_Deserialize_UsingJsonerToAny_Skip(t *testing.T) {
	tc := deserializeUsingJsonerToAnySkipTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingJsonerToAny(true, nil, &s)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingJsonerToAny_NotSkip(t *testing.T) {
	tc := deserializeUsingJsonerToAnyNotSkipTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingJsonerToAny(false, nil, &s)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingDeserializerToOption_Skip(t *testing.T) {
	tc := deserializeUsingDeserializerToOptionSkipTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingDeserializerToOption(true, nil, &s)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingDeserializerToOption_NotSkip(t *testing.T) {
	tc := deserializeUsingDeserializerToOptionNotSkipTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingDeserializerToOption(false, nil, &s)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Deserialize_UsingDeserializerDefined_Nil(t *testing.T) {
	tc := deserializeUsingDeserializerDefinedNilTestCase

	// Arrange
	var s string

	// Act
	err := corejson.Deserialize.UsingDeserializerDefined(nil, &s)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// deserializerLogic — ResultPtr
// =============================================================================

func Test_Deserialize_ResultPtr_Invalid(t *testing.T) {
	tc := deserializeResultPtrInvalidTestCase

	// Arrange
	// (invalid bytes)

	// Act
	_, err := corejson.Deserialize.ResultPtr([]byte(`bad`))
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// deserializeFromBytesTo — String / Strings / Integer / Bool / Map
// =============================================================================

func Test_BytesTo_String_Verification(t *testing.T) {
	for caseIndex, tc := range bytesToStringTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		b, _ := input.Get("bytes")

		// Act
		s, err := corejson.Deserialize.BytesTo.String(b.([]byte))
		actual := args.Map{
			"hasError": err != nil,
			"result":   s,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BytesTo_StringMust(t *testing.T) {
	tc := bytesToStringMustTestCase

	// Arrange
	// (valid json bytes)

	// Act
	s := corejson.Deserialize.BytesTo.StringMust([]byte(`"hello"`))
	actual := args.Map{
		"result": s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesTo_StringMust_Panics(t *testing.T) {
	tc := bytesToStringMustPanicsTestCase

	// Arrange
	// (invalid json bytes)

	// Act
	panicked := callPanics(func() {
		corejson.Deserialize.BytesTo.StringMust([]byte(`bad`))
	})
	actual := args.Map{
		"panicked": panicked,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesTo_Strings(t *testing.T) {
	tc := bytesToStringsTestCase

	// Arrange
	// (valid json array)

	// Act
	ss, err := corejson.Deserialize.BytesTo.Strings([]byte(`["a","b"]`))
	actual := args.Map{
		"hasError": err != nil,
		"length":   len(ss),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesTo_Integer(t *testing.T) {
	tc := bytesToIntegerTestCase

	// Arrange
	// (valid json int)

	// Act
	v, err := corejson.Deserialize.BytesTo.Integer([]byte(`42`))
	actual := args.Map{
		"hasError": err != nil,
		"result":   v,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesTo_Integer64(t *testing.T) {
	tc := bytesToInteger64TestCase

	// Arrange
	// (valid json int64)

	// Act
	v, err := corejson.Deserialize.BytesTo.Integer64([]byte(`99`))
	actual := args.Map{
		"hasError": err != nil,
		"result":   v,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesTo_Bool(t *testing.T) {
	tc := bytesToBoolTestCase

	// Arrange
	// (valid json bool)

	// Act
	v, err := corejson.Deserialize.BytesTo.Bool([]byte(`true`))
	actual := args.Map{
		"hasError": err != nil,
		"result":   v,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesTo_MapAnyItem(t *testing.T) {
	tc := bytesToMapAnyItemTestCase

	// Arrange
	// (valid json object)

	// Act
	m, err := corejson.Deserialize.BytesTo.MapAnyItem([]byte(`{"a":1}`))
	actual := args.Map{
		"hasError": err != nil,
		"hasKeyA":  m["a"] != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytesTo_MapStringString(t *testing.T) {
	tc := bytesToMapStringStringTestCase

	// Arrange
	// (valid json string map)

	// Act
	m, err := corejson.Deserialize.BytesTo.MapStringString([]byte(`{"a":"b"}`))
	actual := args.Map{
		"hasError": err != nil,
		"valueA":   m["a"],
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// deserializeFromResultTo — String / Bool / Byte / MapAnyItem / MapStringString
// =============================================================================

func Test_ResultTo_String(t *testing.T) {
	tc := resultToStringTestCase

	// Arrange
	r := corejson.NewPtr("hello")

	// Act
	s, err := corejson.Deserialize.ResultTo.String(r)
	actual := args.Map{
		"hasError": err != nil,
		"result":   s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultTo_Bool(t *testing.T) {
	tc := resultToBoolTestCase

	// Arrange
	r := corejson.NewPtr(true)

	// Act
	v, err := corejson.Deserialize.ResultTo.Bool(r)
	actual := args.Map{
		"hasError": err != nil,
		"result":   v,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultTo_Byte(t *testing.T) {
	tc := resultToByteTestCase

	// Arrange
	r := corejson.NewPtr(byte(65))

	// Act
	v, err := corejson.Deserialize.ResultTo.Byte(r)
	actual := args.Map{
		"hasError": err != nil,
		"result":   v,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultTo_MapAnyItem(t *testing.T) {
	tc := resultToMapAnyItemTestCase

	// Arrange
	r := corejson.NewPtr(map[string]any{"a": 1})

	// Act
	m, err := corejson.Deserialize.ResultTo.MapAnyItem(r)
	actual := args.Map{
		"hasError": err != nil,
		"hasKeyA":  m["a"] != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultTo_MapStringString(t *testing.T) {
	tc := resultToMapStringStringTestCase

	// Arrange
	r := corejson.NewPtr(map[string]string{"a": "b"})

	// Act
	m, err := corejson.Deserialize.ResultTo.MapStringString(r)
	actual := args.Map{
		"hasError": err != nil,
		"valueA":   m["a"],
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultTo_ResultCollection_Invalid(t *testing.T) {
	tc := resultToResultCollectionInvalidTestCase

	// Arrange
	r := &corejson.Result{Bytes: []byte(`bad`)}

	// Act
	_, err := corejson.Deserialize.ResultTo.ResultCollection(r)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultTo_ResultsPtrCollection_Invalid(t *testing.T) {
	tc := resultToResultsPtrCollectionInvalidTestCase

	// Arrange
	r := &corejson.Result{Bytes: []byte(`bad`)}

	// Act
	_, err := corejson.Deserialize.ResultTo.ResultsPtrCollection(r)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ResultTo_MapResults_Invalid(t *testing.T) {
	tc := resultToMapResultsInvalidTestCase

	// Arrange
	r := &corejson.Result{Bytes: []byte(`bad`)}

	// Act
	_, err := corejson.Deserialize.ResultTo.MapResults(r)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// suppress unused import for errors
var _ = errors.New
