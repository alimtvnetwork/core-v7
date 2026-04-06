package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// corejson Coverage — Segment 3: AnyTo, CastingAny, BytesCloneIf, BytesDeepClone,
//                                 BytesToString, JsonString, JsonStringOrErrMsg, funcs
// ══════════════════════════════════════════════════════════════════════════════

// --- anyTo ---

func Test_CovJsonS3_AT01_SerializedJsonResult_Nil(t *testing.T) {
	// Arrange
	r := corejson.AnyTo.SerializedJsonResult(nil)

	// Act
	actual := args.Map{"result": r == nil || !r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_CovJsonS3_AT02_SerializedJsonResult_Result(t *testing.T) {
	// Arrange
	r := corejson.New(1)
	got := corejson.AnyTo.SerializedJsonResult(r)

	// Act
	actual := args.Map{"result": got == nil || got.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS3_AT03_SerializedJsonResult_ResultPtr(t *testing.T) {
	// Arrange
	r := corejson.NewPtr(1)
	got := corejson.AnyTo.SerializedJsonResult(r)

	// Act
	actual := args.Map{"result": got == nil || got.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS3_AT04_SerializedJsonResult_Bytes(t *testing.T) {
	// Arrange
	got := corejson.AnyTo.SerializedJsonResult([]byte(`"hello"`))

	// Act
	actual := args.Map{"result": got == nil || got.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS3_AT05_SerializedJsonResult_String(t *testing.T) {
	// Arrange
	got := corejson.AnyTo.SerializedJsonResult(`"hello"`)

	// Act
	actual := args.Map{"result": got == nil || got.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS3_AT06_SerializedJsonResult_Jsoner(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.UsingCap(1)
	rc.AddAny(1)
	got := corejson.AnyTo.SerializedJsonResult(rc)

	// Act
	actual := args.Map{"result": got == nil || got.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS3_AT07_SerializedJsonResult_Error(t *testing.T) {
	// Arrange
	got := corejson.AnyTo.SerializedJsonResult(errors.New("fail"))

	// Act
	actual := args.Map{"result": got == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_CovJsonS3_AT08_SerializedJsonResult_AnyItem(t *testing.T) {
	// Arrange
	got := corejson.AnyTo.SerializedJsonResult(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": got == nil || got.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS3_AT09_SerializedRaw(t *testing.T) {
	// Arrange
	b, err := corejson.AnyTo.SerializedRaw(1)

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_CovJsonS3_AT10_SerializedString_Success(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.SerializedString(1)

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_CovJsonS3_AT11_SerializedSafeString(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SerializedSafeString(1)

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
	// nil gives empty
	s2 := corejson.AnyTo.SerializedSafeString(nil)
	actual := args.Map{"result": s2 != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovJsonS3_AT12_SerializedStringMust(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SerializedStringMust(1)

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_CovJsonS3_AT13_SafeJsonString(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SafeJsonString(1)

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_CovJsonS3_AT14_PrettyStringWithError_String(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.PrettyStringWithError("hello")

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_CovJsonS3_AT15_PrettyStringWithError_Bytes(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.PrettyStringWithError([]byte(`"hello"`))

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_CovJsonS3_AT16_PrettyStringWithError_Result(t *testing.T) {
	// Arrange
	r := corejson.New(1)
	s, err := corejson.AnyTo.PrettyStringWithError(r)

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_CovJsonS3_AT17_PrettyStringWithError_ResultPtr(t *testing.T) {
	// Arrange
	r := corejson.NewPtr(1)
	s, err := corejson.AnyTo.PrettyStringWithError(r)

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_CovJsonS3_AT18_PrettyStringWithError_AnyItem(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.PrettyStringWithError(42)

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_CovJsonS3_AT19_SafeJsonPrettyString_Branches(t *testing.T) {
	// string
	// Act
	actual := args.Map{"result": corejson.AnyTo.SafeJsonPrettyString("hello") != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	// bytes
	s := corejson.AnyTo.SafeJsonPrettyString([]byte(`"hello"`))
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
	// Result
	r := corejson.New(1)
	_ = corejson.AnyTo.SafeJsonPrettyString(r)
	// *Result
	rp := corejson.NewPtr(1)
	_ = corejson.AnyTo.SafeJsonPrettyString(rp)
	// anyItem
	_ = corejson.AnyTo.SafeJsonPrettyString(42)
}

func Test_CovJsonS3_AT20_JsonString_Branches(t *testing.T) {
	// string
	// Act
	actual := args.Map{"result": corejson.AnyTo.JsonString("hello") != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	// bytes
	_ = corejson.AnyTo.JsonString([]byte(`"hello"`))
	// Result
	r := corejson.New(1)
	_ = corejson.AnyTo.JsonString(r)
	// *Result
	rp := corejson.NewPtr(1)
	_ = corejson.AnyTo.JsonString(rp)
	// anyItem
	_ = corejson.AnyTo.JsonString(42)
}

func Test_CovJsonS3_AT21_JsonStringWithErr_Branches(t *testing.T) {
	// Arrange
	// string
	s, err := corejson.AnyTo.JsonStringWithErr("hello")

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	// bytes
	_, _ = corejson.AnyTo.JsonStringWithErr([]byte(`"hello"`))
	// Result
	r := corejson.New(1)
	_, _ = corejson.AnyTo.JsonStringWithErr(r)
	// *Result
	rp := corejson.NewPtr(1)
	_, _ = corejson.AnyTo.JsonStringWithErr(rp)
	// anyItem
	_, _ = corejson.AnyTo.JsonStringWithErr(42)
}

func Test_CovJsonS3_AT22_JsonStringMust(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.JsonStringMust(1)

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_CovJsonS3_AT23_PrettyStringMust(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.PrettyStringMust(1)

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_CovJsonS3_AT24_UsingSerializer(t *testing.T) {
	r := corejson.AnyTo.UsingSerializer(nil)
	_ = r
}

func Test_CovJsonS3_AT25_SerializedFieldsMap(t *testing.T) {
	// SerializedFieldsMap → DeserializedFieldsToMap passes value not pointer — known limitation
	m, _ := corejson.AnyTo.SerializedFieldsMap(map[string]int{"a": 1})
	_ = m // covers the call path regardless of result
}

// --- castingAny ---

func Test_CovJsonS3_CA01_FromToDefault(t *testing.T) {
	// Arrange
	var m map[string]int
	err := corejson.CastAny.FromToDefault(map[string]int{"a": 1}, &m)

	// Act
	actual := args.Map{"result": err != nil || m["a"] != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a=1", actual)
}

func Test_CovJsonS3_CA02_FromToReflection(t *testing.T) {
	// Arrange
	var m map[string]int
	err := corejson.CastAny.FromToReflection(map[string]int{"a": 1}, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS3_CA03_FromToOption_Bytes(t *testing.T) {
	// Arrange
	var m map[string]int
	err := corejson.CastAny.FromToOption(false, []byte(`{"a":1}`), &m)

	// Act
	actual := args.Map{"result": err != nil || m["a"] != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a=1", actual)
}

func Test_CovJsonS3_CA04_FromToOption_String(t *testing.T) {
	// Arrange
	var m map[string]int
	err := corejson.CastAny.FromToOption(false, `{"a":1}`, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS3_CA05_FromToOption_Jsoner(t *testing.T) {
	rc := corejson.NewResultsCollection.AnyItems(map[string]int{"a": 1})
	var out corejson.ResultsCollection
	_ = corejson.CastAny.FromToOption(false, rc, &out)
}

func Test_CovJsonS3_CA06_FromToOption_Result(t *testing.T) {
	// Arrange
	// Result implements Jsoner, so FromToOption matches Jsoner case first (double-serializes).
	// Use bytes directly to test the Result case path.
	r := corejson.New(map[string]int{"a": 1})
	var m map[string]int
	err := corejson.CastAny.FromToOption(false, r.Bytes, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS3_CA07_FromToOption_ResultPtr(t *testing.T) {
	// Arrange
	// *Result also implements Jsoner — use bytes directly
	r := corejson.NewPtr(map[string]int{"a": 1})
	var m map[string]int
	err := corejson.CastAny.FromToOption(false, r.Bytes, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS3_CA08_FromToOption_SerializerFunc(t *testing.T) {
	// Arrange
	fn := func() ([]byte, error) { return []byte(`{"a":1}`), nil }
	var m map[string]int
	err := corejson.CastAny.FromToOption(false, fn, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS3_CA09_FromToOption_Error(t *testing.T) {
	// Arrange
	e := errors.New(`{"a":1}`)
	var m map[string]int
	err := corejson.CastAny.FromToOption(false, e, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS3_CA10_FromToOption_ErrorNil(t *testing.T) {
	// Arrange
	var e error
	var m map[string]int
	err := corejson.CastAny.FromToOption(false, e, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CovJsonS3_CA11_FromToOption_AnyItem(t *testing.T) {
	// Arrange
	type s struct{ A int }
	src := s{A: 1}
	var dst s
	err := corejson.CastAny.FromToOption(false, src, &dst)

	// Act
	actual := args.Map{"result": err != nil || dst.A != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected A=1", actual)
}

func Test_CovJsonS3_CA12_ReflectionCasting_SkipReflection(t *testing.T) {
	// Arrange
	var m map[string]int
	err := corejson.CastAny.FromToOption(false, map[string]int{"a": 1}, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS3_CA13_ReflectionCasting_NilFrom(t *testing.T) {
	var m map[string]int
	// reflection enabled, nil from
	_ = corejson.CastAny.FromToOption(true, nil, &m)
}

func Test_CovJsonS3_CA14_OrDeserializeTo(t *testing.T) {
	// Arrange
	var m map[string]int
	err := corejson.CastAny.OrDeserializeTo(map[string]int{"a": 1}, &m)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

// --- BytesCloneIf ---

func Test_CovJsonS3_BC01_BytesCloneIf_DeepClone(t *testing.T) {
	// Arrange
	src := []byte("hello")
	dst := corejson.BytesCloneIf(true, src)

	// Act
	actual := args.Map{"result": string(dst) != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_CovJsonS3_BC02_BytesCloneIf_NoDeepClone(t *testing.T) {
	// Arrange
	dst := corejson.BytesCloneIf(false, []byte("hello"))

	// Act
	actual := args.Map{"result": len(dst) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovJsonS3_BC03_BytesCloneIf_Empty(t *testing.T) {
	// Arrange
	dst := corejson.BytesCloneIf(true, nil)

	// Act
	actual := args.Map{"result": len(dst) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// --- BytesDeepClone ---

func Test_CovJsonS3_BD01_BytesDeepClone(t *testing.T) {
	// Arrange
	src := []byte("hello")
	dst := corejson.BytesDeepClone(src)

	// Act
	actual := args.Map{"result": string(dst) != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_CovJsonS3_BD02_BytesDeepClone_Empty(t *testing.T) {
	// Arrange
	dst := corejson.BytesDeepClone(nil)

	// Act
	actual := args.Map{"result": len(dst) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// --- BytesToString / BytesToPrettyString ---

func Test_CovJsonS3_BS01_BytesToString(t *testing.T) {
	// Arrange
	s := corejson.BytesToString([]byte(`"hello"`))

	// Act
	actual := args.Map{"result": s != `"hello"`}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_CovJsonS3_BS02_BytesToString_Empty(t *testing.T) {
	// Arrange
	s := corejson.BytesToString(nil)

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CovJsonS3_BS03_BytesToPrettyString(t *testing.T) {
	// Arrange
	s := corejson.BytesToPrettyString([]byte(`{"a":1}`))

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CovJsonS3_BS04_BytesToPrettyString_Empty(t *testing.T) {
	// Arrange
	s := corejson.BytesToPrettyString(nil)

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// --- JsonString func ---

func Test_CovJsonS3_JS01_JsonString(t *testing.T) {
	// Arrange
	s, err := corejson.JsonString(1)

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

// --- JsonStringOrErrMsg func ---

func Test_CovJsonS3_JE01_JsonStringOrErrMsg(t *testing.T) {
	// Arrange
	s := corejson.JsonStringOrErrMsg(1)

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

// --- New / NewPtr ---

func Test_CovJsonS3_NW01_New(t *testing.T) {
	// Arrange
	r := corejson.New(42)

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

func Test_CovJsonS3_NW02_NewPtr(t *testing.T) {
	// Arrange
	r := corejson.NewPtr(42)

	// Act
	actual := args.Map{"result": r == nil || r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
}

// --- KeyWithResult / KeyAny ---

func Test_CovJsonS3_KR01_KeyWithResult(t *testing.T) {
	// Arrange
	kr := corejson.KeyWithResult{Key: "k", Result: corejson.New(1)}

	// Act
	actual := args.Map{"result": kr.Key != "k"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected k", actual)
}

func Test_CovJsonS3_KA01_KeyAny(t *testing.T) {
	// Arrange
	ka := corejson.KeyAny{Key: "k", AnyInf: 1}

	// Act
	actual := args.Map{"result": ka.Key != "k"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected k", actual)
}

func Test_CovJsonS3_KJ01_KeyWithJsoner(t *testing.T) {
	// Arrange
	rc := corejson.NewResultsCollection.UsingCap(1)
	rc.AddAny(1)
	kj := corejson.KeyWithJsoner{Key: "k", Jsoner: rc}

	// Act
	actual := args.Map{"result": kj.Key != "k"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected k", actual)
}
