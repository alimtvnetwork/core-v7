package coreoncetests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// Ptr Constructors
// ==========================================================================

var covPtrConstructorTestCases = []coretestcases.CaseV1{
	{
		Title:         "NewAnyOncePtr returns non-nil",
		ArrangeInput:  args.Map{"type": "any"},
		ExpectedInput: args.Map{"notNil": true},
	},
	{
		Title:         "NewBoolOncePtr returns non-nil",
		ArrangeInput:  args.Map{"type": "bool"},
		ExpectedInput: args.Map{"notNil": true},
	},
	{
		Title:         "NewByteOncePtr returns non-nil",
		ArrangeInput:  args.Map{"type": "byte"},
		ExpectedInput: args.Map{"notNil": true},
	},
	{
		Title:         "NewStringOncePtr returns non-nil",
		ArrangeInput:  args.Map{"type": "string"},
		ExpectedInput: args.Map{"notNil": true},
	},
	{
		Title:         "NewIntegerOncePtr returns non-nil",
		ArrangeInput:  args.Map{"type": "integer"},
		ExpectedInput: args.Map{"notNil": true},
	},
	{
		Title:         "NewErrorOncePtr returns non-nil",
		ArrangeInput:  args.Map{"type": "error"},
		ExpectedInput: args.Map{"notNil": true},
	},
	{
		Title:         "NewAnyErrorOncePtr returns non-nil",
		ArrangeInput:  args.Map{"type": "anyError"},
		ExpectedInput: args.Map{"notNil": true},
	},
	{
		Title:         "NewBytesOncePtr returns non-nil",
		ArrangeInput:  args.Map{"type": "bytes"},
		ExpectedInput: args.Map{"notNil": true},
	},
	{
		Title:         "NewBytesErrorOncePtr returns non-nil",
		ArrangeInput:  args.Map{"type": "bytesError"},
		ExpectedInput: args.Map{"notNil": true},
	},
	{
		Title:         "NewStringsOncePtr returns non-nil",
		ArrangeInput:  args.Map{"type": "strings"},
		ExpectedInput: args.Map{"notNil": true},
	},
	{
		Title:         "NewIntegersOncePtr returns non-nil",
		ArrangeInput:  args.Map{"type": "integers"},
		ExpectedInput: args.Map{"notNil": true},
	},
	{
		Title:         "NewMapStringStringOncePtr returns non-nil",
		ArrangeInput:  args.Map{"type": "mapSS"},
		ExpectedInput: args.Map{"notNil": true},
	},
}

func Test_PtrConstructors_Coverage(t *testing.T) {
	for caseIndex, testCase := range covPtrConstructorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		typeName, _ := input.GetAsString("type")
		notNil := false

		switch typeName {
		case "any":
			notNil = coreonce.NewAnyOncePtr(func() any { return nil }) != nil
		case "bool":
			notNil = coreonce.NewBoolOncePtr(func() bool { return false }) != nil
		case "byte":
			notNil = coreonce.NewByteOncePtr(func() byte { return 0 }) != nil
		case "string":
			notNil = coreonce.NewStringOncePtr(func() string { return "" }) != nil
		case "integer":
			notNil = coreonce.NewIntegerOncePtr(func() int { return 0 }) != nil
		case "error":
			notNil = coreonce.NewErrorOncePtr(func() error { return nil }) != nil
		case "anyError":
			notNil = coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil }) != nil
		case "bytes":
			notNil = coreonce.NewBytesOncePtr(func() []byte { return nil }) != nil
		case "bytesError":
			notNil = coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, nil }) != nil
		case "strings":
			notNil = coreonce.NewStringsOncePtr(func() []string { return nil }) != nil
		case "integers":
			notNil = coreonce.NewIntegersOncePtr(func() []int { return nil }) != nil
		case "mapSS":
			notNil = coreonce.NewMapStringStringOncePtr(func() map[string]string { return nil }) != nil
		}

		// Act
		actual := args.Map{"notNil": notNil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// StringOnce — comprehensive method coverage
// ==========================================================================

var covStringOnceTestCases = []coretestcases.CaseV1{
	{
		Title:        "StringOnce methods with value 'hello world'",
		ArrangeInput: args.Map{"value": "hello world"},
		ExpectedInput: args.Map{
			"valuePtr":    true, "isEqual": true, "hasPrefix": true,
			"isStartsWith": true, "hasSuffix": true, "isEndsWith": true,
			"isContains": true, "isEmpty": false, "isEmptyWs": false,
			"bytesLen":   11, "splitLen": 2, "string": "hello world",
		},
	},
	{
		Title:        "StringOnce empty string",
		ArrangeInput: args.Map{"value": ""},
		ExpectedInput: args.Map{
			"isEmpty": true, "isEmptyWs": true,
		},
	},
}

func Test_StringOnce_Methods_Coverage(t *testing.T) {
	for caseIndex, testCase := range covStringOnceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsString("value")
		so := coreonce.NewStringOnce(func() string { return val })

		// Act
		actual := args.Map{
			"isEmpty":   so.IsEmpty(),
			"isEmptyWs": so.IsEmptyOrWhitespace(),
		}

		if val != "" {
			actual["valuePtr"] = so.ValuePtr() != nil
			actual["isEqual"] = so.IsEqual("hello world")
			actual["hasPrefix"] = so.HasPrefix("hello")
			actual["isStartsWith"] = so.IsStartsWith("hello")
			actual["hasSuffix"] = so.HasSuffix("world")
			actual["isEndsWith"] = so.IsEndsWith("world")
			actual["isContains"] = so.IsContains("lo wo")
			actual["bytesLen"] = len(so.Bytes())
			actual["splitLen"] = len(so.SplitBy(" "))
			actual["string"] = so.String()
		}

		// Assert
		expected := testCase.ExpectedInput.(args.Map)
		filtered := args.Map{}
		for k := range expected {
			if v, has := actual[k]; has {
				filtered[k] = v
			}
		}

		testCase.ShouldBeEqualMap(t, caseIndex, filtered)
	}
}

// ==========================================================================
// StringOnce — SplitLeftRight edge cases
// ==========================================================================

var covStringSplitTestCases = []coretestcases.CaseV1{
	{
		Title:         "SplitLeftRight with separator",
		ArrangeInput:  args.Map{
			"value": "key=value",
			"sep": "=",
		},
		ExpectedInput: args.Map{
			"left": "key",
			"right": "value",
		},
	},
	{
		Title:         "SplitLeftRight no separator",
		ArrangeInput:  args.Map{
			"value": "nosep",
			"sep": "=",
		},
		ExpectedInput: args.Map{
			"left": "nosep",
			"right": "",
		},
	},
	{
		Title:         "SplitLeftRightTrim with spaces",
		ArrangeInput:  args.Map{
			"value": " key = value ",
			"sep": "=",
			"trim": true,
		},
		ExpectedInput: args.Map{
			"left": "key",
			"right": "value",
		},
	},
}

func Test_StringOnce_SplitLeftRight_Coverage(t *testing.T) {
	for caseIndex, testCase := range covStringSplitTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsString("value")
		sep, _ := input.GetAsString("sep")
		doTrim := input.GetAsBoolDefault("trim", false)
		so := coreonce.NewStringOnce(func() string { return val })

		var left, right string
		if doTrim {
			left, right = so.SplitLeftRightTrim(sep)
		} else {
			left, right = so.SplitLeftRight(sep)
		}

		// Act
		actual := args.Map{
			"left": left,
			"right": right,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// StringOnce — Error, Serialize, MarshalJSON, UnmarshalJSON
// ==========================================================================

var covStringOnceSerializeTestCases = []coretestcases.CaseV1{
	{
		Title:         "StringOnce Serialize and JSON roundtrip",
		ArrangeInput:  args.Map{"value": "test"},
		ExpectedInput: args.Map{
			"serializeOk": true,
			"marshalOk": true,
			"unmarshalOk": true,
			"errorOk": true,
		},
	},
}

func Test_StringOnce_Serialize_Coverage(t *testing.T) {
	for caseIndex, testCase := range covStringOnceSerializeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsString("value")
		so := coreonce.NewStringOnce(func() string { return val })

		_, serErr := so.Serialize()
		marshalBytes, marshalErr := so.MarshalJSON()
		unmarshalErr := so.UnmarshalJSON(marshalBytes)
		errVal := so.Error()

		// Act
		actual := args.Map{
			"serializeOk":  serErr == nil,
			"marshalOk":    marshalErr == nil,
			"unmarshalOk":  unmarshalErr == nil,
			"errorOk":      errVal != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// StringsOnce — comprehensive coverage
// ==========================================================================

var covStringsOnceTestCases = []coretestcases.CaseV1{
	{
		Title:        "StringsOnce with values",
		ArrangeInput: args.Map{"values": []string{"b", "a", "c"}},
		ExpectedInput: args.Map{
			"length": 3, "hasAnyItem": true, "isEmpty": false,
			"hasAll": true, "contains": true, "has": true,
			"sortedFirst": "a", "rangesLen": 3,
			"csvNotEmpty": true, "csvLinesNotEmpty": true,
			"safeLen": 3, "isEqual": true,
		},
	},
	{
		Title:        "StringsOnce not equal",
		ArrangeInput: args.Map{
			"values": []string{"a"},
			"compare": []string{"b"},
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
}

func Test_StringsOnce_Coverage(t *testing.T) {
	for caseIndex, testCase := range covStringsOnceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		values := input["values"].([]string)
		so := coreonce.NewStringsOnce(func() []string { return values })

		// Act
		actual := args.Map{}

		// Assert
		expected := testCase.ExpectedInput.(args.Map)

		if _, has := expected["length"]; has {
			actual["length"] = so.Length()
			actual["hasAnyItem"] = so.HasAnyItem()
			actual["isEmpty"] = so.IsEmpty()
			actual["hasAll"] = so.HasAll("a", "b")
			actual["contains"] = so.IsContains("a")
			actual["has"] = so.Has("b")
			sorted := so.Sorted()
			actual["sortedFirst"] = sorted[0]
			actual["rangesLen"] = len(so.RangesMap())
			actual["csvNotEmpty"] = so.Csv() != ""
			actual["csvLinesNotEmpty"] = len(so.CsvLines()) > 0
			actual["safeLen"] = len(so.SafeStrings())
			actual["isEqual"] = so.IsEqual("b", "a", "c")
			// also cover UniqueMap, UniqueMapLock, Strings, List, Values, Execute, String
			_ = so.UniqueMap()
			_ = so.UniqueMapLock()
			_ = so.Strings()
			_ = so.List()
			_ = so.Values()
			_ = so.Execute()
			_ = so.String()
			_ = so.JsonStringMust()
		}

		if cmpRaw, has := input["compare"]; has {
			cmp := cmpRaw.([]string)
			actual["isEqual"] = so.IsEqual(cmp...)
		}

		filtered := args.Map{}
		for k := range expected {
			if v, has := actual[k]; has {
				filtered[k] = v
			}
		}
		testCase.ShouldBeEqualMap(t, caseIndex, filtered)
	}
}

// ==========================================================================
// StringsOnce — IsEqual edge cases
// ==========================================================================

var covStringsOnceIsEqualTestCases = []coretestcases.CaseV1{
	{
		Title:         "StringsOnce IsEqual different length returns false",
		ArrangeInput:  args.Map{
			"values": []string{"a"},
			"compare": []string{"a", "b"},
		},
		ExpectedInput: args.Map{"isEqual": false},
	},
	{
		Title:         "StringsOnce IsEqual different value returns false",
		ArrangeInput:  args.Map{
			"values": []string{"a"},
			"compare": []string{"b"},
		},
		ExpectedInput: args.Map{"isEqual": false},
	},
}

func Test_StringsOnce_IsEqual_Coverage(t *testing.T) {
	for caseIndex, testCase := range covStringsOnceIsEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		values := input["values"].([]string)
		compare := input["compare"].([]string)
		so := coreonce.NewStringsOnce(func() []string { return values })

		// Act
		actual := args.Map{"isEqual": so.IsEqual(compare...)}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// IntegersOnce — comprehensive coverage
// ==========================================================================

var covIntegersOnceTestCases = []coretestcases.CaseV1{
	{
		Title:        "IntegersOnce with values",
		ArrangeInput: args.Map{"values": []int{3, 1, 2}},
		ExpectedInput: args.Map{
			"isEmpty": false, "sortedFirst": 1,
			"rangesLen": 3, "rangesBoolLen": 3,
			"uniqueLen": 3, "isEqual": true,
		},
	},
	{
		Title:         "IntegersOnce IsEqual different returns false",
		ArrangeInput:  args.Map{
			"values": []int{1},
			"compare": []int{2},
		},
		ExpectedInput: args.Map{"isEqual": false},
	},
	{
		Title:         "IntegersOnce IsEqual different length returns false",
		ArrangeInput:  args.Map{
			"values": []int{1},
			"compare": []int{1, 2},
		},
		ExpectedInput: args.Map{"isEqual": false},
	},
}

func Test_IntegersOnce_Coverage(t *testing.T) {
	for caseIndex, testCase := range covIntegersOnceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		values := input["values"].([]int)
		io := coreonce.NewIntegersOnce(func() []int { return values })

		// Act
		actual := args.Map{}

		// Assert
		expected := testCase.ExpectedInput.(args.Map)

		if _, has := expected["isEmpty"]; has {
			actual["isEmpty"] = io.IsEmpty()
			sorted := io.Sorted()
			actual["sortedFirst"] = sorted[0]
			actual["rangesLen"] = len(io.RangesMap())
			actual["rangesBoolLen"] = len(io.RangesBoolMap())
			actual["uniqueLen"] = len(io.UniqueMap())
			actual["isEqual"] = io.IsEqual(3, 1, 2)
			_ = io.Values()
			_ = io.Execute()
			_ = io.Integers()
			_ = io.Slice()
			_ = io.List()
			_ = io.String()
		}

		if cmpRaw, has := input["compare"]; has {
			cmp := cmpRaw.([]int)
			actual["isEqual"] = io.IsEqual(cmp...)
		}

		filtered := args.Map{}
		for k := range expected {
			if v, has := actual[k]; has {
				filtered[k] = v
			}
		}
		testCase.ShouldBeEqualMap(t, caseIndex, filtered)
	}
}

// ==========================================================================
// MapStringStringOnce — comprehensive coverage
// ==========================================================================

var covMapStringStringOnceTestCases = []coretestcases.CaseV1{
	{
		Title:        "MapStringStringOnce with values",
		ArrangeInput: args.Map{"map": map[string]string{"a": "1", "b": "2"}},
		ExpectedInput: args.Map{
			"length": 2, "hasAnyItem": true, "isEmpty": false,
			"hasAll": true, "contains": true, "isMissing": true,
			"getValue": "1", "hasValue": true,
			"allKeysLen": 2, "allValuesLen": 2,
			"allKeysSortedFirst": "a", "allValuesSortedLen": 2,
			"stringsLen": 2, "isEqual": true,
		},
	},
}

func Test_MapStringStringOnce_Coverage(t *testing.T) {
	for caseIndex, testCase := range covMapStringStringOnceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		m := input["map"].(map[string]string)
		mso := coreonce.NewMapStringStringOnce(func() map[string]string { return m })

		// Act
		actual := args.Map{
			"length":     mso.Length(),
			"hasAnyItem": mso.HasAnyItem(),
			"isEmpty":    mso.IsEmpty(),
			"hasAll":     mso.HasAll("a", "b"),
			"contains":   mso.IsContains("a"),
			"isMissing":  mso.IsMissing("c"),
			"getValue":   mso.GetValue("a"),
		}

		_, hasValue := mso.GetValueWithStatus("a")
		actual["hasValue"] = hasValue
		actual["allKeysLen"] = len(mso.AllKeys())
		actual["allValuesLen"] = len(mso.AllValues())
		sortedKeys := mso.AllKeysSorted()
		actual["allKeysSortedFirst"] = sortedKeys[0]
		actual["allValuesSortedLen"] = len(mso.AllValuesSorted())
		actual["stringsLen"] = len(mso.Strings())
		actual["isEqual"] = mso.IsEqual(map[string]string{"a": "1", "b": "2"})

		// also cover aliases
		_ = mso.List()
		_ = mso.ItemsMap()
		_ = mso.Values()
		_ = mso.ValuesPtr()
		_ = mso.Execute()
		_ = mso.Has("a")
		_ = mso.String()
		_ = mso.JsonStringMust()
		_, _ = mso.Serialize()
		_, _ = mso.MarshalJSON()

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// MapStringStringOnce — IsEqual edge cases
// ==========================================================================

var covMapSSIsEqualTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsEqual different length returns false",
		ArrangeInput:  args.Map{"map": map[string]string{"a": "1"}},
		ExpectedInput: args.Map{"isEqual": false},
	},
	{
		Title:         "IsEqual missing key returns false",
		ArrangeInput:  args.Map{
			"map": map[string]string{"a": "1"},
			"right": map[string]string{"b": "1"},
		},
		ExpectedInput: args.Map{"isEqual": false},
	},
	{
		Title:         "IsEqual different value returns false",
		ArrangeInput:  args.Map{
			"map": map[string]string{"a": "1"},
			"right": map[string]string{"a": "2"},
		},
		ExpectedInput: args.Map{"isEqual": false},
	},
}

func Test_MapStringStringOnce_IsEqual_Coverage(t *testing.T) {
	for caseIndex, testCase := range covMapSSIsEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		m := input["map"].(map[string]string)
		mso := coreonce.NewMapStringStringOnce(func() map[string]string { return m })

		var right map[string]string
		if r, has := input["right"]; has {
			right = r.(map[string]string)
		} else {
			right = map[string]string{"a": "1", "b": "2"}
		}

		// Act
		actual := args.Map{"isEqual": mso.IsEqual(right)}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// MapStringStringOnce — UnmarshalJSON
// ==========================================================================

var covMapSSUnmarshalTestCases = []coretestcases.CaseV1{
	{
		Title:         "UnmarshalJSON valid JSON",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{"ok": true},
	},
}

func Test_MapStringStringOnce_UnmarshalJSON_Coverage(t *testing.T) {
	for caseIndex, testCase := range covMapSSUnmarshalTestCases {
		// Arrange
		mso := coreonce.NewMapStringStringOnce(func() map[string]string { return nil })
		err := mso.UnmarshalJSON([]byte(`{"k":"v"}`))

		// Act
		actual := args.Map{"ok": err == nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// ByteOnce — additional methods
// ==========================================================================

var covByteOnceMethodsTestCases = []coretestcases.CaseV1{
	{
		Title:        "ByteOnce methods with value 42",
		ArrangeInput: args.Map{"value": byte(42)},
		ExpectedInput: args.Map{
			"intVal": 42, "isEmpty": false, "isZero": false,
			"isNegative": false, "isPositive": true, "string": "42",
		},
	},
	{
		Title:        "ByteOnce zero value",
		ArrangeInput: args.Map{"value": byte(0)},
		ExpectedInput: args.Map{
			"isEmpty": true, "isZero": true,
		},
	},
}

func Test_ByteOnce_Methods_Coverage(t *testing.T) {
	for caseIndex, testCase := range covByteOnceMethodsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val := input["value"].(byte)
		bo := coreonce.NewByteOnce(func() byte { return val })

		// Act
		actual := args.Map{
			"isEmpty": bo.IsEmpty(),
			"isZero":  bo.IsZero(),
		}

		// Assert
		expected := testCase.ExpectedInput.(args.Map)

		if _, has := expected["intVal"]; has {
			actual["intVal"] = bo.Int()
			actual["isNegative"] = bo.IsNegative()
			actual["isPositive"] = bo.IsPositive()
			actual["string"] = bo.String()
			_, _ = bo.MarshalJSON()
			_, _ = bo.Serialize()
		}

		filtered := args.Map{}
		for k := range expected {
			if v, has := actual[k]; has {
				filtered[k] = v
			}
		}
		testCase.ShouldBeEqualMap(t, caseIndex, filtered)
	}
}

// ==========================================================================
// IntegerOnce — additional methods
// ==========================================================================

var covIntegerOnceMethodsTestCases = []coretestcases.CaseV1{
	{
		Title:        "IntegerOnce positive methods",
		ArrangeInput: args.Map{"value": 5},
		ExpectedInput: args.Map{
			"isAbove3": true, "isAboveEq5": true,
			"isLessThan10": true, "isLessThanEq5": true,
			"isPositive": true, "isNegative": false,
			"string": "5",
		},
	},
}

func Test_IntegerOnce_Methods_Coverage(t *testing.T) {
	for caseIndex, testCase := range covIntegerOnceMethodsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val := input.GetAsIntDefault("value", 0)
		io := coreonce.NewIntegerOnce(func() int { return val })

		// Act
		actual := args.Map{
			"isAbove3":      io.IsAbove(3),
			"isAboveEq5":    io.IsAboveEqual(5),
			"isLessThan10":  io.IsLessThan(10),
			"isLessThanEq5": io.IsLessThanEqual(5),
			"isPositive":    io.IsPositive(),
			"isNegative":    io.IsNegative(),
			"string":        io.String(),
		}
		_, _ = io.MarshalJSON()

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// ErrorOnce — additional methods
// ==========================================================================

var covErrorOnceMethodsTestCases = []coretestcases.CaseV1{
	{
		Title:        "ErrorOnce with error -- all status methods",
		ArrangeInput: args.Map{"hasError": true},
		ExpectedInput: args.Map{
			"hasError": true, "isEmpty": false, "hasAnyItem": true,
			"isDefined": true, "isInvalid": true, "isValid": false,
			"isSuccess": false, "isFailed": true, "isNull": false,
			"messageNotEmpty": true, "isMessageEqual": true,
		},
	},
	{
		Title:        "ErrorOnce nil -- MarshalJSON",
		ArrangeInput: args.Map{"hasError": false},
		ExpectedInput: args.Map{
			"marshalOk": true,
			"isMessageEqual": false,
		},
	},
}

func Test_ErrorOnce_Methods_Coverage(t *testing.T) {
	for caseIndex, testCase := range covErrorOnceMethodsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		hasErr := input.GetAsBoolDefault("hasError", false)

		// Act
		actual := args.Map{}

		if hasErr {
			eo := coreonce.NewErrorOnce(func() error { return fmt.Errorf("test-error") })
			actual["hasError"] = eo.HasError()
			actual["isEmpty"] = eo.IsEmpty()
			actual["hasAnyItem"] = eo.HasAnyItem()
			actual["isDefined"] = eo.IsDefined()
			actual["isInvalid"] = eo.IsInvalid()
			actual["isValid"] = eo.IsValid()
			actual["isSuccess"] = eo.IsSuccess()
			actual["isFailed"] = eo.IsFailed()
			actual["isNull"] = eo.IsNull()
			actual["messageNotEmpty"] = eo.Message() != ""
			actual["isMessageEqual"] = eo.IsMessageEqual("test-error")
		} else {
			eo := coreonce.NewErrorOnce(func() error { return nil })
			_, marshalErr := eo.MarshalJSON()
			actual["marshalOk"] = marshalErr == nil
			actual["isMessageEqual"] = eo.IsMessageEqual("anything")
		}

		// Assert
		expected := testCase.ExpectedInput.(args.Map)
		filtered := args.Map{}
		for k := range expected {
			if v, has := actual[k]; has {
				filtered[k] = v
			}
		}
		testCase.ShouldBeEqualMap(t, caseIndex, filtered)
	}
}

// ==========================================================================
// AnyErrorOnce — additional methods
// ==========================================================================

var covAnyErrorOnceMethodsTestCases = []coretestcases.CaseV1{
	{
		Title:        "AnyErrorOnce success -- all status methods",
		ArrangeInput: args.Map{"hasError": false},
		ExpectedInput: args.Map{
			"isEmpty": false, "hasAnyItem": true, "isDefined": true,
			"isInvalid": false, "isValid": true, "isSuccess": true,
			"isFailed": false, "safeStringOk": true,
			"valueStringMustOk": true,
		},
	},
	{
		Title:        "AnyErrorOnce nil value -- IsEmpty true",
		ArrangeInput: args.Map{"nilValue": true},
		ExpectedInput: args.Map{
			"isEmpty": true, "isStringEmptyWs": true,
		},
	},
}

func Test_AnyErrorOnce_Methods_Coverage(t *testing.T) {
	for caseIndex, testCase := range covAnyErrorOnceMethodsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal := input.GetAsBoolDefault("nilValue", false)

		// Act
		actual := args.Map{}

		if isNilVal {
			aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return nil, nil })
			actual["isEmpty"] = aeo.IsEmpty()
			actual["isStringEmptyWs"] = aeo.IsStringEmptyOrWhitespace()
		} else {
			aeo := coreonce.NewAnyErrorOnce(func() (any, error) { return "val", nil })
			actual["isEmpty"] = aeo.IsEmpty()
			actual["hasAnyItem"] = aeo.HasAnyItem()
			actual["isDefined"] = aeo.IsDefined()
			actual["isInvalid"] = aeo.IsInvalid()
			actual["isValid"] = aeo.IsValid()
			actual["isSuccess"] = aeo.IsSuccess()
			actual["isFailed"] = aeo.IsFailed()
			actual["safeStringOk"] = aeo.SafeString() != ""
			actual["valueStringMustOk"] = aeo.ValueStringMust() != ""
		}

		// Assert
		expected := testCase.ExpectedInput.(args.Map)
		filtered := args.Map{}
		for k := range expected {
			if v, has := actual[k]; has {
				filtered[k] = v
			}
		}
		testCase.ShouldBeEqualMap(t, caseIndex, filtered)
	}
}

// ==========================================================================
// BytesOnce — nil initializerFunc, empty checks
// ==========================================================================

var covBytesOnceTestCases = []coretestcases.CaseV1{
	{
		Title:         "BytesOnce nil initializerFunc is empty",
		ArrangeInput:  args.Map{"nilInit": true},
		ExpectedInput: args.Map{
			"isEmpty": true,
			"length": 0,
		},
	},
	{
		Title:         "BytesOnce with data",
		ArrangeInput:  args.Map{"value": "abc"},
		ExpectedInput: args.Map{
			"isEmpty": false,
			"length": 3,
			"string": "abc",
		},
	},
}

func Test_BytesOnce_Coverage(t *testing.T) {
	for caseIndex, testCase := range covBytesOnceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nilInit := input.GetAsBoolDefault("nilInit", false)

		// Act
		actual := args.Map{}

		if nilInit {
			bo := &coreonce.BytesOnce{}
			actual["isEmpty"] = bo.IsEmpty()
			actual["length"] = bo.Length()
		} else {
			val, _ := input.GetAsString("value")
			bo := coreonce.NewBytesOnce(func() []byte { return []byte(val) })
			actual["isEmpty"] = bo.IsEmpty()
			actual["length"] = bo.Length()
			actual["string"] = bo.String()
			_ = bo.Execute()
			_, _ = bo.MarshalJSON()
			_ = bo.UnmarshalJSON([]byte(`"dGVzdA=="`))
		}

		// Assert
		expected := testCase.ExpectedInput.(args.Map)
		filtered := args.Map{}
		for k := range expected {
			if v, has := actual[k]; has {
				filtered[k] = v
			}
		}
		testCase.ShouldBeEqualMap(t, caseIndex, filtered)
	}
}

// ==========================================================================
// BytesErrorOnce — HasIssuesOrEmpty, HasSafeItems, Deserialize
// ==========================================================================

var covBytesErrorOnceTestCases = []coretestcases.CaseV1{
	{
		Title:        "BytesErrorOnce with valid JSON bytes",
		ArrangeInput: args.Map{"hasError": false},
		ExpectedInput: args.Map{
			"hasIssues": false, "hasSafeItems": true,
			"deserializeOk": true, "isBytesEmpty": false,
			"isNull": false, "isStringEmpty": false,
			"isStringEmptyWs": false, "marshalOk": true,
		},
	},
	{
		Title:        "BytesErrorOnce with error",
		ArrangeInput: args.Map{"hasError": true},
		ExpectedInput: args.Map{
			"hasIssues": true, "hasSafeItems": false,
			"isEmptyBytes": true,
		},
	},
}

func Test_BytesErrorOnce_Coverage(t *testing.T) {
	for caseIndex, testCase := range covBytesErrorOnceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		hasErr := input.GetAsBoolDefault("hasError", false)

		// Act
		actual := args.Map{}

		if hasErr {
			beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return nil, fmt.Errorf("fail") })
			actual["hasIssues"] = beo.HasIssuesOrEmpty()
			actual["hasSafeItems"] = beo.HasSafeItems()
			actual["isEmptyBytes"] = beo.IsEmptyBytes()
		} else {
			beo := coreonce.NewBytesErrorOnce(func() ([]byte, error) { return []byte(`"hello"`), nil })
			actual["hasIssues"] = beo.HasIssuesOrEmpty()
			actual["hasSafeItems"] = beo.HasSafeItems()

			var target string
			desErr := beo.Deserialize(&target)
			actual["deserializeOk"] = desErr == nil
			actual["isBytesEmpty"] = beo.IsBytesEmpty()
			actual["isNull"] = beo.IsNull()
			actual["isStringEmpty"] = beo.IsStringEmpty()
			actual["isStringEmptyWs"] = beo.IsStringEmptyOrWhitespace()
			_, marshalErr := beo.MarshalJSON()
			actual["marshalOk"] = marshalErr == nil
			_, _ = beo.ValueWithError()
			_ = beo.HasAnyItem()
			_ = beo.IsDefined()
			_ = beo.IsInvalid()
			_ = beo.IsValid()
			_ = beo.IsSuccess()
			_ = beo.IsFailed()
			_ = beo.IsInitialized()
			_ = beo.String()
			_, _ = beo.Serialize()
		}

		// Assert
		expected := testCase.ExpectedInput.(args.Map)
		filtered := args.Map{}
		for k := range expected {
			if v, has := actual[k]; has {
				filtered[k] = v
			}
		}
		testCase.ShouldBeEqualMap(t, caseIndex, filtered)
	}
}

// ==========================================================================
// StringsOnce — Serialize, MarshalJSON, UnmarshalJSON
// ==========================================================================

var covStringsOnceSerializeTestCases = []coretestcases.CaseV1{
	{
		Title:         "StringsOnce Serialize and JSON roundtrip",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{
			"serializeOk": true,
			"marshalOk": true,
			"unmarshalOk": true,
		},
	},
}

func Test_StringsOnce_Serialize_Coverage(t *testing.T) {
	for caseIndex, testCase := range covStringsOnceSerializeTestCases {
		// Arrange
		so := coreonce.NewStringsOnce(func() []string { return []string{"a", "b"} })
		_, serErr := so.Serialize()
		marshalBytes, marshalErr := so.MarshalJSON()
		unmarshalErr := so.UnmarshalJSON(marshalBytes)

		// Act
		actual := args.Map{
			"serializeOk":  serErr == nil,
			"marshalOk":    marshalErr == nil,
			"unmarshalOk":  unmarshalErr == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// IntegersOnce — Serialize, MarshalJSON, UnmarshalJSON
// ==========================================================================

var covIntegersOnceSerializeTestCases = []coretestcases.CaseV1{
	{
		Title:         "IntegersOnce Serialize and JSON roundtrip",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{
			"serializeOk": true,
			"marshalOk": true,
			"unmarshalOk": true,
		},
	},
}

func Test_IntegersOnce_Serialize_Coverage(t *testing.T) {
	for caseIndex, testCase := range covIntegersOnceSerializeTestCases {
		// Arrange
		io := coreonce.NewIntegersOnce(func() []int { return []int{1, 2} })
		_, serErr := io.Serialize()
		marshalBytes, marshalErr := io.MarshalJSON()
		unmarshalErr := io.UnmarshalJSON(marshalBytes)

		// Act
		actual := args.Map{
			"serializeOk":  serErr == nil,
			"marshalOk":    marshalErr == nil,
			"unmarshalOk":  unmarshalErr == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// IntegersOnce — empty checks
// ==========================================================================

var covIntegersOnceEmptyTestCases = []coretestcases.CaseV1{
	{
		Title:         "IntegersOnce empty returns zero for maps",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{
			"isEmpty": true,
			"isZero": true,
			"rangesLen": 0,
			"rangesBoolLen": 0,
			"uniqueLen": 0,
		},
	},
}

func Test_IntegersOnce_Empty_Coverage(t *testing.T) {
	for caseIndex, testCase := range covIntegersOnceEmptyTestCases {
		// Arrange
		io := coreonce.NewIntegersOnce(func() []int { return []int{} })

		// Act
		actual := args.Map{
			"isEmpty":       io.IsEmpty(),
			"isZero":        io.IsZero(),
			"rangesLen":     len(io.RangesMap()),
			"rangesBoolLen": len(io.RangesBoolMap()),
			"uniqueLen":     len(io.UniqueMap()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
