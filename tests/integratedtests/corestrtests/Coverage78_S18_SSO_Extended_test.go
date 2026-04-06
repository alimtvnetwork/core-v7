package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ========================================
// S18: SimpleStringOnce extended
//   Split, Clone, JSON, Serialize,
//   newSimpleStringOnceCreator
// ========================================

func Test_C78_SSO_LinesSimpleSlice(t *testing.T) {
	safeTest(t, "Test_C78_SSO_LinesSimpleSlice", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a\nb\nc")

		// Act
		result := sso.LinesSimpleSlice()

		// Assert
		if result.Length() != 3 {
			t.Errorf("expected 3, got %d", result.Length())
		}
	})
}

func Test_C78_SSO_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_C78_SSO_SimpleSlice", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a:b:c")

		// Act
		result := sso.SimpleSlice(":")

		// Assert
		if result.Length() != 3 {
			t.Errorf("expected 3, got %d", result.Length())
		}
	})
}

func Test_C78_SSO_Split(t *testing.T) {
	safeTest(t, "Test_C78_SSO_Split", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a,b,c")

		// Act
		result := sso.Split(",")

		// Assert
		if len(result) != 3 {
			t.Errorf("expected 3, got %d", len(result))
		}
	})
}

func Test_C78_SSO_SplitLeftRight(t *testing.T) {
	safeTest(t, "Test_C78_SSO_SplitLeftRight", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("key=value")

		// Act
		left, right := sso.SplitLeftRight("=")

		// Assert
		if left != "key" || right != "value" {
			t.Errorf("expected key/value, got %s/%s", left, right)
		}
	})
}

func Test_C78_SSO_SplitLeftRight_NoSep(t *testing.T) {
	safeTest(t, "Test_C78_SSO_SplitLeftRight_NoSep", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("noseparator")

		// Act
		left, right := sso.SplitLeftRight("=")

		// Assert
		if left != "noseparator" || right != "" {
			t.Errorf("expected noseparator/'', got %s/%s", left, right)
		}
	})
}

func Test_C78_SSO_SplitLeftRightTrim(t *testing.T) {
	safeTest(t, "Test_C78_SSO_SplitLeftRightTrim", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init(" key = value ")

		// Act
		left, right := sso.SplitLeftRightTrim("=")

		// Assert
		if left != "key" || right != "value" {
			t.Errorf("expected key/value, got '%s'/'%s'", left, right)
		}
	})
}

func Test_C78_SSO_SplitLeftRightTrim_NoSep(t *testing.T) {
	safeTest(t, "Test_C78_SSO_SplitLeftRightTrim_NoSep", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init(" nosep ")

		// Act
		left, right := sso.SplitLeftRightTrim("=")

		// Assert
		if left != "nosep" || right != "" {
			t.Errorf("expected nosep/'', got '%s'/'%s'", left, right)
		}
	})
}

func Test_C78_SSO_SplitNonEmpty(t *testing.T) {
	safeTest(t, "Test_C78_SSO_SplitNonEmpty", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a::b::c")

		// Act
		result := sso.SplitNonEmpty("::")

		// Assert
		if len(result) < 3 {
			t.Errorf("expected at least 3, got %d", len(result))
		}
	})
}

func Test_C78_SSO_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C78_SSO_SplitTrimNonWhitespace", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a , , b")

		// Act
		result := sso.SplitTrimNonWhitespace(",")

		// Assert
		if len(result) < 2 {
			t.Errorf("expected at least 2 non-whitespace items, got %d", len(result))
		}
	})
}

func Test_C78_SSO_ClonePtr(t *testing.T) {
	safeTest(t, "Test_C78_SSO_ClonePtr", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("hello")

		// Act
		cloned := sso.ClonePtr()

		// Assert
		if cloned == nil || cloned.Value() != "hello" {
			t.Error("clone mismatch")
		}
	})
}

func Test_C78_SSO_ClonePtr_Nil(t *testing.T) {
	safeTest(t, "Test_C78_SSO_ClonePtr_Nil", func() {
		// Arrange
		var sso *corestr.SimpleStringOnce

		// Act
		cloned := sso.ClonePtr()

		// Assert
		if cloned != nil {
			t.Error("expected nil")
		}
	})
}

func Test_C78_SSO_Clone(t *testing.T) {
	safeTest(t, "Test_C78_SSO_Clone", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act
		cloned := sso.Clone()

		// Assert
		if cloned.Value() != "x" {
			t.Error("clone mismatch")
		}
	})
}

func Test_C78_SSO_CloneUsingNewVal(t *testing.T) {
	safeTest(t, "Test_C78_SSO_CloneUsingNewVal", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("old")

		// Act
		cloned := sso.CloneUsingNewVal("new")

		// Assert
		if cloned.Value() != "new" {
			t.Errorf("expected 'new', got '%s'", cloned.Value())
		}
		if !cloned.IsInitialized() {
			t.Error("expected initialized from source")
		}
	})
}

func Test_C78_SSO_Dispose(t *testing.T) {
	safeTest(t, "Test_C78_SSO_Dispose", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("val")

		// Act
		sso.Dispose()

		// Assert
		if sso.Value() != "" {
			t.Error("expected empty after dispose")
		}
	})
}

func Test_C78_SSO_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_C78_SSO_Dispose_Nil", func() {
		// Arrange
		var sso *corestr.SimpleStringOnce

		// Act — should not panic
		sso.Dispose()
	})
}

func Test_C78_SSO_String(t *testing.T) {
	safeTest(t, "Test_C78_SSO_String", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("hello")

		// Act & Assert
		if sso.String() != "hello" {
			t.Error("String mismatch")
		}
	})
}

func Test_C78_SSO_String_Nil(t *testing.T) {
	safeTest(t, "Test_C78_SSO_String_Nil", func() {
		// Arrange
		var sso *corestr.SimpleStringOnce

		// Act & Assert
		if sso.String() != "" {
			t.Error("expected empty for nil")
		}
	})
}

func Test_C78_SSO_StringPtr(t *testing.T) {
	safeTest(t, "Test_C78_SSO_StringPtr", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("val")

		// Act
		result := sso.StringPtr()

		// Assert
		if result == nil || *result != "val" {
			t.Error("StringPtr mismatch")
		}
	})
}

func Test_C78_SSO_StringPtr_Nil(t *testing.T) {
	safeTest(t, "Test_C78_SSO_StringPtr_Nil", func() {
		// Arrange
		var sso *corestr.SimpleStringOnce

		// Act
		result := sso.StringPtr()

		// Assert
		if result == nil || *result != "" {
			t.Error("expected empty string ptr for nil")
		}
	})
}

func Test_C78_SSO_JsonModel(t *testing.T) {
	safeTest(t, "Test_C78_SSO_JsonModel", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("test")

		// Act
		model := sso.JsonModel()

		// Assert
		if model.Value != "test" || !model.IsInitialize {
			t.Error("model mismatch")
		}
	})
}

func Test_C78_SSO_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C78_SSO_JsonModelAny", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("x")

		// Act & Assert
		if sso.JsonModelAny() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C78_SSO_MarshalUnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C78_SSO_MarshalUnmarshalJSON", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("hello")

		// Act
		bytes, err := sso.MarshalJSON()
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		target := corestr.New.SimpleStringOnce.CreatePtr("", false)
		err = target.UnmarshalJSON(bytes)

		// Assert
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}
		if target.Value() != "hello" {
			t.Errorf("expected 'hello', got '%s'", target.Value())
		}
	})
}

func Test_C78_SSO_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C78_SSO_Json_JsonPtr", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act
		jsonResult := sso.Json()
		jsonPtrResult := sso.JsonPtr()

		// Assert
		if jsonResult.HasError() {
			t.Error("json error")
		}
		if jsonPtrResult.HasError() {
			t.Error("jsonPtr error")
		}
	})
}

func Test_C78_SSO_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C78_SSO_ParseInjectUsingJson", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")
		jsonResult := sso.JsonPtr()
		target := corestr.New.SimpleStringOnce.CreatePtr("", false)

		// Act
		result, err := target.ParseInjectUsingJson(jsonResult)

		// Assert
		if err != nil {
			t.Errorf("error: %v", err)
		}
		if result.Value() != "hello" {
			t.Error("value mismatch")
		}
	})
}

func Test_C78_SSO_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C78_SSO_ParseInjectUsingJsonMust", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("test")
		jsonResult := sso.JsonPtr()
		target := corestr.New.SimpleStringOnce.CreatePtr("", false)

		// Act
		result := target.ParseInjectUsingJsonMust(jsonResult)

		// Assert
		if result.Value() != "test" {
			t.Error("value mismatch")
		}
	})
}

func Test_C78_SSO_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C78_SSO_AsJsonContractsBinder", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("x")

		// Act & Assert
		if sso.AsJsonContractsBinder() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C78_SSO_AsJsoner(t *testing.T) {
	safeTest(t, "Test_C78_SSO_AsJsoner", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("x")

		// Act & Assert
		if sso.AsJsoner() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C78_SSO_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C78_SSO_JsonParseSelfInject", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("val")
		jsonResult := sso.JsonPtr()
		target := corestr.New.SimpleStringOnce.CreatePtr("", false)

		// Act
		err := target.JsonParseSelfInject(jsonResult)

		// Assert
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

func Test_C78_SSO_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_C78_SSO_AsJsonParseSelfInjector", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("x")

		// Act & Assert
		if sso.AsJsonParseSelfInjector() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C78_SSO_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_C78_SSO_AsJsonMarshaller", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("x")

		// Act & Assert
		if sso.AsJsonMarshaller() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C78_SSO_Serialize(t *testing.T) {
	safeTest(t, "Test_C78_SSO_Serialize", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("hello")

		// Act
		bytes, err := sso.Serialize()

		// Assert
		if err != nil {
			t.Errorf("error: %v", err)
		}
		if len(bytes) == 0 {
			t.Error("expected non-empty bytes")
		}
	})
}

func Test_C78_SSO_Deserialize(t *testing.T) {
	safeTest(t, "Test_C78_SSO_Deserialize", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("data")

		// Act
		var target corestr.SimpleStringOnceModel
		err := sso.Deserialize(&target)

		// Assert
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

// --- newSimpleStringOnceCreator ---

func Test_C78_NewSSO_Any(t *testing.T) {
	safeTest(t, "Test_C78_NewSSO_Any", func() {
		// Arrange & Act
		sso := corestr.New.SimpleStringOnce.Any(false, 42, true)

		// Assert
		if sso.Value() != "42" {
			t.Errorf("expected '42', got '%s'", sso.Value())
		}
		if !sso.IsInitialized() {
			t.Error("expected initialized")
		}
	})
}

func Test_C78_NewSSO_Uninitialized(t *testing.T) {
	safeTest(t, "Test_C78_NewSSO_Uninitialized", func() {
		// Arrange & Act
		sso := corestr.New.SimpleStringOnce.Uninitialized("val")

		// Assert
		if sso.Value() != "val" {
			t.Error("value mismatch")
		}
		if sso.IsInitialized() {
			t.Error("expected uninitialized")
		}
	})
}

func Test_C78_NewSSO_Init(t *testing.T) {
	safeTest(t, "Test_C78_NewSSO_Init", func() {
		// Arrange & Act
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Assert
		if sso.Value() != "x" || !sso.IsInitialized() {
			t.Error("Init mismatch")
		}
	})
}

func Test_C78_NewSSO_InitPtr(t *testing.T) {
	safeTest(t, "Test_C78_NewSSO_InitPtr", func() {
		// Arrange & Act
		sso := corestr.New.SimpleStringOnce.InitPtr("x")

		// Assert
		if sso == nil || sso.Value() != "x" {
			t.Error("InitPtr mismatch")
		}
	})
}

func Test_C78_NewSSO_Create(t *testing.T) {
	safeTest(t, "Test_C78_NewSSO_Create", func() {
		// Arrange & Act
		sso := corestr.New.SimpleStringOnce.Create("val", true)

		// Assert
		if sso.Value() != "val" || !sso.IsInitialized() {
			t.Error("Create mismatch")
		}
	})
}

func Test_C78_NewSSO_CreatePtr(t *testing.T) {
	safeTest(t, "Test_C78_NewSSO_CreatePtr", func() {
		// Arrange & Act
		sso := corestr.New.SimpleStringOnce.CreatePtr("val", false)

		// Assert
		if sso == nil || sso.Value() != "val" || sso.IsInitialized() {
			t.Error("CreatePtr mismatch")
		}
	})
}

func Test_C78_NewSSO_Empty(t *testing.T) {
	safeTest(t, "Test_C78_NewSSO_Empty", func() {
		// Arrange & Act
		sso := corestr.New.SimpleStringOnce.Empty()

		// Assert
		if sso.Value() != "" || sso.IsInitialized() {
			t.Error("Empty mismatch")
		}
	})
}
