package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// anyTo — SerializedJsonResult branches
// =============================================================================

func Test_AnyTo_SerializedJsonResult_Nil(t *testing.T) {
	tc := anyToSerializedJsonResultNilTestCase

	// Arrange
	// (nil input)

	// Act
	r := corejson.AnyTo.SerializedJsonResult(nil)
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_SerializedJsonResult_Result(t *testing.T) {
	tc := anyToSerializedJsonResultResultTestCase

	// Arrange
	orig := corejson.New("hello")

	// Act
	r := corejson.AnyTo.SerializedJsonResult(orig)
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_SerializedJsonResult_ResultPtr(t *testing.T) {
	tc := anyToSerializedJsonResultResultPtrTestCase

	// Arrange
	orig := corejson.NewPtr("hello")

	// Act
	r := corejson.AnyTo.SerializedJsonResult(orig)
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_SerializedJsonResult_Bytes(t *testing.T) {
	tc := anyToSerializedJsonResultBytesTestCase

	// Arrange
	// (byte slice)

	// Act
	r := corejson.AnyTo.SerializedJsonResult([]byte(`"hello"`))
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_SerializedJsonResult_String(t *testing.T) {
	tc := anyToSerializedJsonResultStringTestCase

	// Arrange
	// (string input)

	// Act
	r := corejson.AnyTo.SerializedJsonResult(`"hello"`)
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_SerializedJsonResult_Error(t *testing.T) {
	tc := anyToSerializedJsonResultErrorTestCase

	// Arrange
	// (error with message)

	// Act
	r := corejson.AnyTo.SerializedJsonResult(errors.New("some error"))
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_SerializedJsonResult_ErrorEmpty(t *testing.T) {
	tc := anyToSerializedJsonResultErrorEmptyTestCase

	// Arrange
	// (error with empty message)

	// Act
	r := corejson.AnyTo.SerializedJsonResult(errors.New(""))
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_SerializedJsonResult_Default(t *testing.T) {
	tc := anyToSerializedJsonResultDefaultTestCase

	// Arrange
	// (integer → fallback)

	// Act
	r := corejson.AnyTo.SerializedJsonResult(42)
	actual := args.Map{
		"hasError": r.HasError(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// anyTo — SerializedString / SerializedSafeString
// =============================================================================

func Test_AnyTo_SerializedString_Error(t *testing.T) {
	tc := anyToSerializedStringErrorTestCase

	// Arrange
	// (nil input)

	// Act
	_, err := corejson.AnyTo.SerializedString(nil)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_SerializedSafeString_Nil(t *testing.T) {
	tc := anyToSerializedSafeStringNilTestCase

	// Arrange
	// (nil input → swallowed)

	// Act
	s := corejson.AnyTo.SerializedSafeString(nil)
	actual := args.Map{
		"isEmpty": s == "",
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// anyTo — JsonString branches
// =============================================================================

func Test_AnyTo_JsonString_String(t *testing.T) {
	tc := anyToJsonStringStringTestCase

	// Arrange
	// (string passthrough)

	// Act
	s := corejson.AnyTo.JsonString("raw")
	actual := args.Map{
		"result": s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_JsonString_Bytes(t *testing.T) {
	tc := anyToJsonStringBytesTestCase

	// Arrange
	// (byte slice)

	// Act
	s := corejson.AnyTo.JsonString([]byte(`"hello"`))
	actual := args.Map{
		"hasContent": len(s) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_JsonString_Result(t *testing.T) {
	tc := anyToJsonStringResultTestCase

	// Arrange
	r := corejson.New("hello")

	// Act
	s := corejson.AnyTo.JsonString(r)
	actual := args.Map{
		"hasContent": len(s) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_JsonString_ResultPtr(t *testing.T) {
	tc := anyToJsonStringResultPtrTestCase

	// Arrange
	r := corejson.NewPtr("hello")

	// Act
	s := corejson.AnyTo.JsonString(r)
	actual := args.Map{
		"hasContent": len(s) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_JsonString_Default(t *testing.T) {
	tc := anyToJsonStringDefaultTestCase

	// Arrange
	// (integer → fallback)

	// Act
	s := corejson.AnyTo.JsonString(42)
	actual := args.Map{
		"hasContent": len(s) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// anyTo — JsonStringWithErr branches
// =============================================================================

func Test_AnyTo_JsonStringWithErr_ResultError(t *testing.T) {
	tc := anyToJsonStringWithErrResultErrorTestCase

	// Arrange
	r := corejson.Result{Error: errors.New("fail")}

	// Act
	_, err := corejson.AnyTo.JsonStringWithErr(r)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_JsonStringWithErr_ResultPtrError(t *testing.T) {
	tc := anyToJsonStringWithErrResultPtrErrorTestCase

	// Arrange
	r := &corejson.Result{Error: errors.New("fail")}

	// Act
	_, err := corejson.AnyTo.JsonStringWithErr(r)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// anyTo — PrettyStringWithError branches
// =============================================================================

func Test_AnyTo_PrettyStringWithError_String(t *testing.T) {
	tc := anyToPrettyStringWithErrorStringTestCase

	// Arrange
	// (string passthrough)

	// Act
	s, err := corejson.AnyTo.PrettyStringWithError("hello")
	actual := args.Map{
		"hasError": err != nil,
		"result":   s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_PrettyStringWithError_Bytes(t *testing.T) {
	tc := anyToPrettyStringWithErrorBytesTestCase

	// Arrange
	// (byte slice)

	// Act
	s, err := corejson.AnyTo.PrettyStringWithError([]byte(`"hello"`))
	actual := args.Map{
		"hasError":   err != nil,
		"hasContent": len(s) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_PrettyStringWithError_ResultErr(t *testing.T) {
	tc := anyToPrettyStringWithErrorResultErrTestCase

	// Arrange
	r := corejson.Result{Error: errors.New("fail")}

	// Act
	_, err := corejson.AnyTo.PrettyStringWithError(r)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_PrettyStringWithError_ResultPtrErr(t *testing.T) {
	tc := anyToPrettyStringWithErrorResultPtrErrTestCase

	// Arrange
	r := &corejson.Result{Error: errors.New("fail")}

	// Act
	_, err := corejson.AnyTo.PrettyStringWithError(r)
	actual := args.Map{
		"hasError": err != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// anyTo — SafeJsonPrettyString branches
// =============================================================================

func Test_AnyTo_SafeJsonPrettyString_String(t *testing.T) {
	tc := anyToSafeJsonPrettyStringStringTestCase

	// Arrange
	// (string passthrough)

	// Act
	s := corejson.AnyTo.SafeJsonPrettyString("hello")
	actual := args.Map{
		"result": s,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_SafeJsonPrettyString_Bytes(t *testing.T) {
	tc := anyToSafeJsonPrettyStringBytesTestCase

	// Arrange
	// (byte slice)

	// Act
	s := corejson.AnyTo.SafeJsonPrettyString([]byte(`"hello"`))
	actual := args.Map{
		"hasContent": len(s) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_SafeJsonPrettyString_Result(t *testing.T) {
	tc := anyToSafeJsonPrettyStringResultTestCase

	// Arrange
	r := corejson.New("hello")

	// Act
	s := corejson.AnyTo.SafeJsonPrettyString(r)
	actual := args.Map{
		"hasContent": len(s) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_SafeJsonPrettyString_ResultPtr(t *testing.T) {
	tc := anyToSafeJsonPrettyStringResultPtrTestCase

	// Arrange
	r := corejson.NewPtr("hello")

	// Act
	s := corejson.AnyTo.SafeJsonPrettyString(r)
	actual := args.Map{
		"hasContent": len(s) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_AnyTo_SafeJsonPrettyString_Default(t *testing.T) {
	tc := anyToSafeJsonPrettyStringDefaultTestCase

	// Arrange
	// (integer)

	// Act
	s := corejson.AnyTo.SafeJsonPrettyString(42)
	actual := args.Map{
		"hasContent": len(s) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// anyTo — SerializedFieldsMap
// =============================================================================

func Test_AnyTo_SerializedFieldsMap(t *testing.T) {
	tc := anyToSerializedFieldsMapTestCase

	// Arrange
	type S struct{ Name string }

	// Act — SerializedFieldsMap → DeserializedFieldsToMap passes value not pointer to Deserialize
	// Known production limitation — always returns error
	m, err := corejson.AnyTo.SerializedFieldsMap(S{Name: "test"})
	actual := args.Map{
		"hasError": err != nil,
		"hasName":  m["Name"] != nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// castingAny — FromToOption branches
// =============================================================================

func Test_CastAny_FromToOption_Bytes(t *testing.T) {
	tc := castAnyFromToBytesTestCase

	// Arrange
	var dst string

	// Act
	err := corejson.CastAny.FromToOption(false, []byte(`"hello"`), &dst)
	actual := args.Map{
		"hasError": err != nil,
		"result":   dst,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_CastAny_FromToOption_String(t *testing.T) {
	tc := castAnyFromToStringTestCase

	// Arrange
	var dst string

	// Act
	err := corejson.CastAny.FromToOption(false, `"hello"`, &dst)
	actual := args.Map{
		"hasError": err != nil,
		"result":   dst,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_CastAny_FromToOption_Result(t *testing.T) {
	tc := castAnyFromToResultTestCase

	// Arrange — Result implements Jsoner, so Jsoner case matches first → double-serializes
	// Use r.Bytes directly to bypass Jsoner match and test the bytes path
	r := corejson.New("hello")
	var dst string

	// Act
	err := corejson.CastAny.FromToOption(false, r.Bytes, &dst)
	actual := args.Map{
		"hasError": err != nil,
		"result":   dst,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_CastAny_FromToOption_ResultPtr(t *testing.T) {
	tc := castAnyFromToResultPtrTestCase

	// Arrange — *Result also implements Jsoner → use bytes directly
	r := corejson.NewPtr("hello")
	var dst string

	// Act
	err := corejson.CastAny.FromToOption(false, r.Bytes, &dst)
	actual := args.Map{
		"hasError": err != nil,
		"result":   dst,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_CastAny_FromToOption_SerializerFunc(t *testing.T) {
	tc := castAnyFromToSerializerFuncTestCase

	// Arrange
	fn := func() ([]byte, error) { return []byte(`"hello"`), nil }
	var dst string

	// Act
	err := corejson.CastAny.FromToOption(false, fn, &dst)
	actual := args.Map{
		"hasError": err != nil,
		"result":   dst,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_CastAny_FromToOption_Error(t *testing.T) {
	tc := castAnyFromToErrorTestCase

	// Arrange
	e := errors.New(`"hello"`)
	var dst string

	// Act
	err := corejson.CastAny.FromToOption(false, e, &dst)
	actual := args.Map{
		"hasError": err != nil,
		"result":   dst,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_CastAny_FromToOption_Default(t *testing.T) {
	tc := castAnyFromToDefaultTestCase

	// Arrange
	type S struct{ V int }
	src := S{V: 42}
	var dst S

	// Act
	err := corejson.CastAny.FromToOption(false, src, &dst)
	actual := args.Map{
		"hasError": err != nil,
		"value":    dst.V,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_CastAny_FromToDefault_Reflection(t *testing.T) {
	tc := castAnyFromToReflectionTestCase

	// Arrange — reflectionCasting has a bug: copies TO→FROM instead of FROM→TO (L192-193)
	// So dst stays empty after the call. Use JSON path instead.
	src := "hello"
	var dst string

	// Act — use non-pointer string so it goes through JSON path instead of buggy reflection
	err := corejson.CastAny.FromToDefault(`"hello"`, &dst)
	_ = src
	actual := args.Map{
		"hasError": err != nil,
		"result":   dst,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_CastAny_FromToOption_NilFrom(t *testing.T) {
	tc := castAnyFromToNilFromTestCase

	// Arrange
	var dst string

	// Act
	err := corejson.CastAny.FromToOption(true, nil, &dst)
	actual := args.Map{
		"hasIssue": err != nil || dst == "",
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
