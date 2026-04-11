package coreoncetests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// BytesErrorOnce — Core
// =============================================================================

func Test_BytesErrorOnce_Core(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceCoreTestCases {
		// Arrange
		initBytes := tc.InitBytes
		initErr := tc.InitErr
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			return initBytes, initErr
		})

		// Act
		val, err := once.Value()
		actual := args.Map{
			"stringValue":  string(val),
			"noError":      err == nil,
			"length":       once.Length(),
			"hasAnyItem":   once.HasAnyItem(),
			"isEmpty":      once.IsEmpty(),
			"isEmptyBytes": once.IsEmptyBytes(),
			"isBytesEmpty": once.IsBytesEmpty(),
			"isNull":       once.IsNull(),
			"isDefined":    once.IsDefined(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// BytesErrorOnce — Caching
// =============================================================================

func Test_BytesErrorOnce_Caching(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceCachingTestCases {
		// Arrange
		initBytes := tc.InitBytes
		initErr := tc.InitErr

		if initErr != nil {
			once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
				return initBytes, initErr
			})

			// Act
			val, err := once.Value()
			actual := args.Map{
				"emptyValue":   string(val),
				"hasError":     val == nil,
				"errorMessage": err.Error(),
			}

			// Assert
			tc.Case.ShouldBeEqualMap(t, caseIndex, actual)

			continue
		}

		callCount := 0
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			callCount++

			return initBytes, nil
		})

		// Act
		r1, _ := once.Value()
		r2, _ := once.Value()

		actual = args.Map{
			"value1":         string(r1),
			"value2":         string(r2),
			"value1EqValue2": string(r1) == string(r2),
			"executeEqValue": func() bool {
				ev, _ := once.Execute()
				vv, _ := once.Value()
				return string(ev) == string(vv)
			}(),
			"callCount": callCount,
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// BytesErrorOnce — Execute
// =============================================================================

func Test_BytesErrorOnce_Execute(t *testing.T) {
	tc := bytesErrorOnceExecuteTestCase

	// Arrange
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return tc.InitBytes, nil
	})

	// Act
	v1, _ := once.Execute()
	v2, _ := once.Value()

	actual := args.Map{
		"executeEqValue": string(v1) == string(v2),
	}

	// Assert
	tc.Case.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// BytesErrorOnce — ValueOnly
// =============================================================================

func Test_BytesErrorOnce_ValueOnly(t *testing.T) {
	tc := bytesErrorOnceValueOnlyTestCase

	// Arrange
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return tc.InitBytes, nil
	})

	// Act
	actual := args.Map{
		"valueOnlyResult": string(once.ValueOnly()),
	}

	// Assert
	tc.Case.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// BytesErrorOnce — ValueWithError
// =============================================================================

func Test_BytesErrorOnce_ValueWithError(t *testing.T) {
	tc := bytesErrorOnceValueWithErrorTestCase

	// Arrange
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return tc.InitBytes, nil
	})

	// Act
	v, e := once.ValueWithError()

	actual := args.Map{
		"valueWithErrorResult": string(v),
		"noError":              e == nil,
	}

	// Assert
	tc.Case.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// BytesErrorOnce — Error State
// =============================================================================

func Test_BytesErrorOnce_ErrorState(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceErrorStateTestCases {
		// Arrange
		initBytes := tc.InitBytes
		initErr := tc.InitErr
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			return initBytes, initErr
		})

		// Act
		actual := args.Map{
			"hasError":  once.HasError(),
			"isValid":   once.IsEmptyError(),
			"isSuccess": once.IsValid(),
			"isEmpty":   once.IsSuccess(),
			"isInvalid": once.IsInvalid(),
			"isFailed":  once.IsFailed(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// BytesErrorOnce — HasIssues
// =============================================================================

func Test_BytesErrorOnce_HasIssues(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceHasIssuesTestCases {
		// Arrange
		var once *coreonce.BytesErrorOnce

		if tc.IsNilReceiver {
			once = nil
		} else {
			initBytes := tc.InitBytes
			initErr := tc.InitErr
			once = coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
				return initBytes, initErr
			})
		}

		// Act
		actual := args.Map{
			"hasIssuesOrEmpty": once.HasIssuesOrEmpty(),
			"hasSafeItems":     once.HasSafeItems(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// BytesErrorOnce — String
// =============================================================================

func Test_BytesErrorOnce_String(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceStringTestCases {
		// Arrange
		initBytes := tc.InitBytes
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			return initBytes, nil
		})

		// Act
		actual := args.Map{
			"stringValue":               once.String(),
			"isStringEmpty":             once.IsStringEmpty(),
			"isStringEmptyOrWhitespace": once.IsStringEmptyOrWhitespace(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// BytesErrorOnce — Deserialize
// =============================================================================

func Test_BytesErrorOnce_Deserialize(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceDeserializeTestCases {
		// Arrange
		initJson := tc.InitJson
		initErr := tc.InitErr
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			return []byte(initJson), initErr
		})

		// Act
		var actual args.Map

		if tc.IsMust {
			var result map[string]string
			panicked := callPanics(func() { once.DeserializeMust(&result) })
			actual = args.Map{"didPanic": panicked}

			if !panicked {
				actual["deserializedKey"] = result["key"]
			}
		} else if initErr != nil {
			var result map[string]string
			err := once.Deserialize(&result)
			actual = args.Map{
				"hasSourceError":      err != nil,
				"hasDeserializeError": strings.Contains(err.Error(), "existing error cannot deserialize"),
				"errorsMatch":         strings.Contains(err.Error(), initErr.Error()),
			}
		} else if initJson == "not-json" {
			var result map[string]string
			err = once.Deserialize(&result)
			actual = args.Map{
				"hasError":    err != nil,
				"isJsonError": strings.Contains(err.Error(), "deserialize failed"),
			}
		} else {
			var result map[string]string
			err = once.Deserialize(&result)
			actual = args.Map{
				"noError":          err == nil,
				"deserializedName": result["name"],
			}
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// BytesErrorOnce — MarshalJSON
// =============================================================================

func Test_BytesErrorOnce_MarshalJSON(t *testing.T) {
	tc := bytesErrorOnceMarshalJSONTestCase

	// Arrange
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return tc.InitBytes, nil
	})

	// Act
	data, err := once.MarshalJSON()

	actual := args.Map{
		"noError":        err == nil,
		"marshaledValue": string(data),
	}

	// Assert
	tc.Case.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// BytesErrorOnce — Serialize
// =============================================================================

func Test_BytesErrorOnce_Serialize(t *testing.T) {
	tc := bytesErrorOnceSerializeTestCase

	// Arrange
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return tc.InitBytes, nil
	})

	// Act
	data, err := once.Serialize()

	actual := args.Map{
		"noError":         err == nil,
		"serializedValue": string(data),
	}

	// Assert
	tc.Case.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// BytesErrorOnce — SerializeMust
// =============================================================================

func Test_BytesErrorOnce_SerializeMust(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceSerializeMustTestCases {
		// Arrange
		initBytes := tc.InitBytes
		initErr := tc.InitErr
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			return initBytes, initErr
		})

		// Act
		var result []byte
		panicked := callPanics(func() { result = once.SerializeMust() })
		actual := args.Map{"didPanic": panicked}

		if !panicked {
			actual["serializedValue"] = string(result)
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// BytesErrorOnce — Lifecycle (panic guards + IsInitialized)
// =============================================================================

func Test_BytesErrorOnce_Lifecycle(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceLifecycleTestCases {
		// Arrange
		initBytes := tc.InitBytes
		initErr := tc.InitErr
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			return initBytes, initErr
		})

		// Act
		actual := args.Map{
			"handleErrorPanicked":       callPanics(func() { once.HandleError() }),
			"mustBeEmptyErrorPanicked":  callPanics(func() { once.MustBeEmptyError() }),
			"mustHaveSafeItemsPanicked": callPanics(func() { once.MustHaveSafeItems() }),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// BytesErrorOnce — IsInitialized
// =============================================================================

func Test_BytesErrorOnce_IsInitialized(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceInitializedTestCases {
		// Arrange
		initBytes := tc.InitBytes
		once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
			return initBytes, nil
		})

		// Act
		beforeInit := once.IsInitialized()
		_, _ = once.Value()
		afterInit := once.IsInitialized()

		actual := args.Map{
			"isInitializedBefore": beforeInit,
			"isInitializedAfter":  afterInit,
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// BytesErrorOnce — Constructor
// =============================================================================

func Test_BytesErrorOnce_Constructor(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceConstructorTestCases {
		// Arrange
		initBytes := tc.InitBytes
		once := coreonce.NewBytesErrorOnce(func() ([]byte, error) {
			return initBytes, nil
		})

		// Act
		v, e := once.Value()
		actual := args.Map{
			"value":     string(v),
			"isCorrect": e == nil,
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// Ensure fmt is used
var _ = fmt.Sprintf
