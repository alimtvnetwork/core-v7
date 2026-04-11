package coreoncetests

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_BoolOnce_Extended_Verification(t *testing.T) {
	for caseIndex, testCase := range boolOnceExtendedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val := input.GetAsBoolDefault("value", false)
		boolOnce := coreonce.NewBoolOnce(func() bool { return val })

		// Act
		executeResult := boolOnce.Execute()
		_, serializeErr := boolOnce.Serialize()
		unmarshalErr := boolOnce.UnmarshalJSON([]byte("true"))

		actual := args.Map{
			"execute":     executeResult,
			"unmarshalOk": unmarshalErr == nil,
			"serializeOk": serializeErr == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ByteOnce_Extended_Verification(t *testing.T) {
	for caseIndex, testCase := range byteOnceExtendedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rawVal, _ := input.Get("value")
		val := rawVal.(byte)
		byteOnce := coreonce.NewByteOnce(func() byte { return val })

		// Act
		executeResult := byteOnce.Execute()
		unmarshalErr := byteOnce.UnmarshalJSON([]byte("42"))

		actual := args.Map{
			"execute":     executeResult,
			"unmarshalOk": unmarshalErr == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// Test_IntegerOnce_Extended_Verification tests boolean checks using a SEPARATE
// instance from serialize/unmarshal to avoid UnmarshalJSON corrupting innerData.
func Test_IntegerOnce_Extended_Verification(t *testing.T) {
	for caseIndex, testCase := range integerOnceExtendedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val := input.GetAsIntDefault("value", 0)

		// Use one instance for boolean checks only (no serialize/unmarshal)
		boolOnce := coreonce.NewIntegerOnce(func() int { return val })
		executeResult := boolOnce.Execute()
		isAboveEqualZero := boolOnce.IsAboveEqualZero()
		isLessThanEqZero := boolOnce.IsLessThanEqualZero()
		isInvalidIndex := boolOnce.IsInvalidIndex()
		isValidIndex := boolOnce.IsValidIndex()

		// Use a SEPARATE instance for serialize/unmarshal
		serOnce := coreonce.NewIntegerOnce(func() int { return val })
		_ = serOnce.Execute()
		_, serializeErr := serOnce.Serialize()
		unmarshalErr := serOnce.UnmarshalJSON([]byte("10"))

		actual := args.Map{
			"execute":          executeResult,
			"isAboveEqualZero": isAboveEqualZero,
			"isLessThanEqZero": isLessThanEqZero,
			"isInvalidIndex":   isInvalidIndex,
			"isValidIndex":     isValidIndex,
			"unmarshalOk":      unmarshalErr == nil,
			"serializeOk":      serializeErr == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// Test_ErrorOnce_Extended_Verification tests ErrorOnce with and without errors.
func Test_ErrorOnce_Extended_Verification(t *testing.T) {
	for caseIndex, testCase := range errorOnceExtendedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		hasErr := input.GetAsBoolDefault("hasError", false)

		var errOnce coreonce.ErrorOnce
		if hasErr {
			msg, _ := input.GetAsString("message")
			errOnce = coreonce.NewErrorOnce(func() error { return errors.New(msg) })
		} else {
			errOnce = coreonce.NewErrorOnce(func() error { return nil })
		}

		// Act
		executeResult := errOnce.Execute()
		isEmptyErr := errOnce.IsEmptyError()
		_, serializeErr := errOnce.Serialize()
		concatErr := errOnce.ConcatNew("extra")
		unmarshalErr := errOnce.UnmarshalJSON([]byte(`"hello"`))

		// String() on nil error does NOT panic — it returns ""
		canString := false
		func() {
			defer func() { recover() }()
			_ = errOnce.String()
			canString = true
		}()

		// handle no-panic for nil error
		func() {
			defer func() { recover() }()
			errOnce2 := coreonce.NewErrorOnce(func() error { return nil })
			errOnce2.HandleError()
			errOnce2.HandleErrorWith("msg")
		}()

		actual := args.Map{
			"hasError":    executeResult != nil,
			"isEmptyErr":  isEmptyErr,
			"executeOk":   true,
			"stringOk":    canString,
			"serializeOk": serializeErr == nil,
			"concatOk":    concatErr != nil,
			"unmarshalOk": unmarshalErr == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_AnyOnce_Extended_Verification(t *testing.T) {
	for caseIndex, testCase := range anyOnceExtendedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		expected := testCase.ExpectedInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)
		valueType, _ := input.GetAsString("valueType")

		actual := args.Map{}

		if isNil {
			// nil value path
			anyOnce := coreonce.NewAnyOnce(func() any { return nil })
			_ = anyOnce.Value()

			actual["isNull"] = anyOnce.IsNull()
			actual["isStringEmpty"] = anyOnce.IsStringEmpty()
			actual["isStringEmptyWs"] = anyOnce.IsStringEmptyOrWhitespace()
		} else if valueType == "map" {
			anyOnce := coreonce.NewAnyOnce(func() any { return map[string]any{"k": "v"} })
			_, ok := anyOnce.CastValueMapStringAnyMap()
			actual["castMapOk"] = ok
		} else if valueType == "strings" {
			anyOnce := coreonce.NewAnyOnce(func() any { return []string{"a", "b"} })
			_, ok := anyOnce.CastValueStrings()
			actual["castStringsOk"] = ok
		} else if valueType == "bytes" {
			anyOnce := coreonce.NewAnyOnce(func() any { return []byte("abc") })
			_, ok := anyOnce.CastValueBytes()
			actual["castBytesOk"] = ok
		} else if valueType == "hashmapMap" {
			anyOnce := coreonce.NewAnyOnce(func() any { return map[string]string{"a": "b"} })
			_, ok := anyOnce.CastValueHashmapMap()
			actual["castHashmapOk"] = ok
		} else {
			// string value path
			val, _ := input.GetAsString("value")
			anyOnce := coreonce.NewAnyOnce(func() any { return val })

			vs := anyOnce.ValueStringOnly()
			vsm := anyOnce.ValueStringMust()
			ss := anyOnce.SafeString()
			_, castOk := anyOnce.CastValueString()
			vo := anyOnce.ValueOnly()
			_, serErr := anyOnce.Serialize()
			sm := anyOnce.SerializeMust()
			var target string
			desErr := anyOnce.Deserialize(&target)

			actual["valueStringOk"] = vs != ""
			actual["valueStringMust"] = vsm != ""
			actual["safeStringOk"] = ss != ""
			actual["castStringOk"] = castOk
			actual["isInitialized"] = anyOnce.IsInitialized()
			actual["isStringEmpty"] = anyOnce.IsStringEmpty()
			actual["isStringEmptyWs"] = anyOnce.IsStringEmptyOrWhitespace()
			actual["deserializeOk"] = desErr == nil
			actual["serializeOk"] = serErr == nil
			actual["serializeMustOk"] = len(sm) > 0
			actual["valueOnlyOk"] = vo != nil
		}

		// filter to only expected keys
		filtered := args.Map{}
		for k := range expected {
			if v, has := actual[k]; has {
				filtered[k] = v
			}
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, filtered)
	}
}

func Test_AnyErrorOnce_Extended_Verification(t *testing.T) {
	for caseIndex, testCase := range anyErrorOnceExtendedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		expected := testCase.ExpectedInput.(args.Map)
		hasErr := input.GetAsBoolDefault("hasError", false)
		valueType, _ := input.GetAsString("valueType")

		actual := args.Map{}

		if valueType == "map" {
			aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return map[string]any{"k": "v"}, nil })
			_, _, ok := aeo.CastValueMapStringAnyMap()
			actual["castMapOk"] = ok
		} else if valueType == "strings" {
			aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return []string{"a"}, nil })
			_, _, ok := aeo.CastValueStrings()
			actual["castStringsOk"] = ok
		} else if valueType == "bytes" {
			aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return []byte("abc"), nil })
			_, _, ok := aeo.CastValueBytes()
			actual["castBytesOk"] = ok
		} else if valueType == "hashmapMap" {
			aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return map[string]string{"a": "b"}, nil })
			_, _, ok := aeo.CastValueHashmapMap()
			actual["castHashmapOk"] = ok
		} else if hasErr {
			aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return nil, errors.New("fail") })
			actual["isNull"] = aeo.IsNull()
			actual["hasError"] = aeo.HasError()
			actual["isStringEmpty"] = aeo.IsStringEmpty()
		} else {
			val, _ := input.GetAsString("value")
			aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return val, nil })

			_, vweErr := aeo.ValueWithError()
			em := aeo.ExecuteMust()
			vsm, _ := aeo.ValueString()
			_, _, castOk := aeo.CastValueString()
			serBytes, serErr := aeo.Serialize()
			smBytes := aeo.SerializeMust()
			skipBytes, skipErr := aeo.SerializeSkipExistingError()
			var target string
			desErr := aeo.Deserialize(&target)

			actual["valueWithErrorOk"] = vweErr == nil
			actual["executeMustOk"] = em != nil
			actual["valueStringMust"] = vsm != ""
			actual["isInitialized"] = aeo.IsInitialized()
			actual["isStringEmpty"] = aeo.IsStringEmpty()
			actual["isStringEmptyWs"] = aeo.IsStringEmptyOrWhitespace()
			actual["castStringOk"] = castOk
			actual["serializeOk"] = serErr == nil && len(serBytes) > 0
			actual["serializeMustOk"] = len(smBytes) > 0
			actual["serializeSkipOk"] = skipErr == nil && len(skipBytes) > 0
			actual["deserializeOk"] = desErr == nil
		}

		// filter actual to match expected keys
		filtered := args.Map{}
		for k := range expected {
			if v, has := actual[k]; has {
				filtered[k] = v
			}
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, filtered)
	}
}

func Test_BytesOnce_Serialize_Verification(t *testing.T) {
	for caseIndex, testCase := range bytesOnceSerializeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsString("value")
		bytesOnce := coreonce.NewBytesOnce(func() []byte { return []byte(val) })

		// Act
		_, serErr := bytesOnce.Serialize()

		actual := args.Map{
			"serializeOk": serErr == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// Ensure json.Marshal round-trip works for BoolOnce.
func Test_BoolOnce_MarshalJSON_RoundTrip(t *testing.T) {
	// Arrange
	boolOnce := coreonce.NewBoolOnce(func() bool { return true })

	// Act
	bytes, err := json.Marshal(&boolOnce)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MarshalJSON failed:", actual)
	actual = args.Map{"result": len(bytes) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MarshalJSON returned empty bytes", actual)
}
