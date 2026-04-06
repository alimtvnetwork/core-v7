package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
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
		actual := args.Map{"result": result.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C78_SSO_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_C78_SSO_SimpleSlice", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a:b:c")

		// Act
		result := sso.SimpleSlice(":")

		// Assert
		actual := args.Map{"result": result.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C78_SSO_Split(t *testing.T) {
	safeTest(t, "Test_C78_SSO_Split", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a,b,c")

		// Act
		result := sso.Split(",")

		// Assert
		actual := args.Map{"result": len(result) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C78_SSO_SplitLeftRight(t *testing.T) {
	safeTest(t, "Test_C78_SSO_SplitLeftRight", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("key=value")

		// Act
		left, right := sso.SplitLeftRight("=")

		// Assert
		actual := args.Map{"result": left != "key" || right != "value"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected key/value, got/", actual)
	})
}

func Test_C78_SSO_SplitLeftRight_NoSep(t *testing.T) {
	safeTest(t, "Test_C78_SSO_SplitLeftRight_NoSep", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("noseparator")

		// Act
		left, right := sso.SplitLeftRight("=")

		// Assert
		actual := args.Map{"result": left != "noseparator" || right != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected noseparator/'', got/", actual)
	})
}

func Test_C78_SSO_SplitLeftRightTrim(t *testing.T) {
	safeTest(t, "Test_C78_SSO_SplitLeftRightTrim", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init(" key = value ")

		// Act
		left, right := sso.SplitLeftRightTrim("=")

		// Assert
		actual := args.Map{"result": left != "key" || right != "value"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected key/value, got ''/''", actual)
	})
}

func Test_C78_SSO_SplitLeftRightTrim_NoSep(t *testing.T) {
	safeTest(t, "Test_C78_SSO_SplitLeftRightTrim_NoSep", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init(" nosep ")

		// Act
		left, right := sso.SplitLeftRightTrim("=")

		// Assert
		actual := args.Map{"result": left != "nosep" || right != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nosep/'', got ''/''", actual)
	})
}

func Test_C78_SSO_SplitNonEmpty(t *testing.T) {
	safeTest(t, "Test_C78_SSO_SplitNonEmpty", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a::b::c")

		// Act
		result := sso.SplitNonEmpty("::")

		// Assert
		actual := args.Map{"result": len(result) < 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	})
}

func Test_C78_SSO_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C78_SSO_SplitTrimNonWhitespace", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a , , b")

		// Act
		result := sso.SplitTrimNonWhitespace(",")

		// Assert
		actual := args.Map{"result": len(result) < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2 non-whitespace items", actual)
	})
}

func Test_C78_SSO_ClonePtr(t *testing.T) {
	safeTest(t, "Test_C78_SSO_ClonePtr", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("hello")

		// Act
		cloned := sso.ClonePtr()

		// Assert
		actual := args.Map{"result": cloned == nil || cloned.Value() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "clone mismatch", actual)
	})
}

func Test_C78_SSO_ClonePtr_Nil(t *testing.T) {
	safeTest(t, "Test_C78_SSO_ClonePtr_Nil", func() {
		// Arrange
		var sso *corestr.SimpleStringOnce

		// Act
		cloned := sso.ClonePtr()

		// Assert
		actual := args.Map{"result": cloned != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C78_SSO_Clone(t *testing.T) {
	safeTest(t, "Test_C78_SSO_Clone", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act
		cloned := sso.Clone()

		// Assert
		actual := args.Map{"result": cloned.Value() != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "clone mismatch", actual)
	})
}

func Test_C78_SSO_CloneUsingNewVal(t *testing.T) {
	safeTest(t, "Test_C78_SSO_CloneUsingNewVal", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("old")

		// Act
		cloned := sso.CloneUsingNewVal("new")

		// Assert
		actual := args.Map{"result": cloned.Value() != "new"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'new', got ''", actual)
		actual := args.Map{"result": cloned.IsInitialized()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected initialized from source", actual)
	})
}

func Test_C78_SSO_Dispose(t *testing.T) {
	safeTest(t, "Test_C78_SSO_Dispose", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("val")

		// Act
		sso.Dispose()

		// Assert
		actual := args.Map{"result": sso.Value() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty after dispose", actual)
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
		actual := args.Map{"result": sso.String() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "String mismatch", actual)
	})
}

func Test_C78_SSO_String_Nil(t *testing.T) {
	safeTest(t, "Test_C78_SSO_String_Nil", func() {
		// Arrange
		var sso *corestr.SimpleStringOnce

		// Act & Assert
		actual := args.Map{"result": sso.String() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	})
}

func Test_C78_SSO_StringPtr(t *testing.T) {
	safeTest(t, "Test_C78_SSO_StringPtr", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("val")

		// Act
		result := sso.StringPtr()

		// Assert
		actual := args.Map{"result": result == nil || *result != "val"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "StringPtr mismatch", actual)
	})
}

func Test_C78_SSO_StringPtr_Nil(t *testing.T) {
	safeTest(t, "Test_C78_SSO_StringPtr_Nil", func() {
		// Arrange
		var sso *corestr.SimpleStringOnce

		// Act
		result := sso.StringPtr()

		// Assert
		actual := args.Map{"result": result == nil || *result != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty string ptr for nil", actual)
	})
}

func Test_C78_SSO_JsonModel(t *testing.T) {
	safeTest(t, "Test_C78_SSO_JsonModel", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("test")

		// Act
		model := sso.JsonModel()

		// Assert
		actual := args.Map{"result": model.Value != "test" || !model.IsInitialize}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "model mismatch", actual)
	})
}

func Test_C78_SSO_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C78_SSO_JsonModelAny", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("x")

		// Act & Assert
		actual := args.Map{"result": sso.JsonModelAny() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C78_SSO_MarshalUnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C78_SSO_MarshalUnmarshalJSON", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("hello")

		// Act
		bytes, err := sso.MarshalJSON()
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "marshal error:", actual)

		target := corestr.New.SimpleStringOnce.CreatePtr("", false)
		err = target.UnmarshalJSON(bytes)

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unmarshal error:", actual)
		actual := args.Map{"result": target.Value() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'hello', got ''", actual)
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
		actual := args.Map{"result": jsonResult.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "json error", actual)
		actual := args.Map{"result": jsonPtrResult.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "jsonPtr error", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
		actual := args.Map{"result": result.Value() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "value mismatch", actual)
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
		actual := args.Map{"result": result.Value() != "test"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "value mismatch", actual)
	})
}

func Test_C78_SSO_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C78_SSO_AsJsonContractsBinder", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("x")

		// Act & Assert
		actual := args.Map{"result": sso.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C78_SSO_AsJsoner(t *testing.T) {
	safeTest(t, "Test_C78_SSO_AsJsoner", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("x")

		// Act & Assert
		actual := args.Map{"result": sso.AsJsoner() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C78_SSO_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_C78_SSO_AsJsonParseSelfInjector", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("x")

		// Act & Assert
		actual := args.Map{"result": sso.AsJsonParseSelfInjector() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C78_SSO_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_C78_SSO_AsJsonMarshaller", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("x")

		// Act & Assert
		actual := args.Map{"result": sso.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C78_SSO_Serialize(t *testing.T) {
	safeTest(t, "Test_C78_SSO_Serialize", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.InitPtr("hello")

		// Act
		bytes, err := sso.Serialize()

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
		actual := args.Map{"result": len(bytes) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty bytes", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

// --- newSimpleStringOnceCreator ---

func Test_C78_NewSSO_Any(t *testing.T) {
	safeTest(t, "Test_C78_NewSSO_Any", func() {
		// Arrange & Act
		sso := corestr.New.SimpleStringOnce.Any(false, 42, true)

		// Assert
		actual := args.Map{"result": sso.Value() != "42"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected '42', got ''", actual)
		actual := args.Map{"result": sso.IsInitialized()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected initialized", actual)
	})
}

func Test_C78_NewSSO_Uninitialized(t *testing.T) {
	safeTest(t, "Test_C78_NewSSO_Uninitialized", func() {
		// Arrange & Act
		sso := corestr.New.SimpleStringOnce.Uninitialized("val")

		// Assert
		actual := args.Map{"result": sso.Value() != "val"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "value mismatch", actual)
		actual := args.Map{"result": sso.IsInitialized()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected uninitialized", actual)
	})
}

func Test_C78_NewSSO_Init(t *testing.T) {
	safeTest(t, "Test_C78_NewSSO_Init", func() {
		// Arrange & Act
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Assert
		actual := args.Map{"result": sso.Value() != "x" || !sso.IsInitialized()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Init mismatch", actual)
	})
}

func Test_C78_NewSSO_InitPtr(t *testing.T) {
	safeTest(t, "Test_C78_NewSSO_InitPtr", func() {
		// Arrange & Act
		sso := corestr.New.SimpleStringOnce.InitPtr("x")

		// Assert
		actual := args.Map{"result": sso == nil || sso.Value() != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "InitPtr mismatch", actual)
	})
}

func Test_C78_NewSSO_Create(t *testing.T) {
	safeTest(t, "Test_C78_NewSSO_Create", func() {
		// Arrange & Act
		sso := corestr.New.SimpleStringOnce.Create("val", true)

		// Assert
		actual := args.Map{"result": sso.Value() != "val" || !sso.IsInitialized()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Create mismatch", actual)
	})
}

func Test_C78_NewSSO_CreatePtr(t *testing.T) {
	safeTest(t, "Test_C78_NewSSO_CreatePtr", func() {
		// Arrange & Act
		sso := corestr.New.SimpleStringOnce.CreatePtr("val", false)

		// Assert
		actual := args.Map{"result": sso == nil || sso.Value() != "val" || sso.IsInitialized()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CreatePtr mismatch", actual)
	})
}

func Test_C78_NewSSO_Empty(t *testing.T) {
	safeTest(t, "Test_C78_NewSSO_Empty", func() {
		// Arrange & Act
		sso := corestr.New.SimpleStringOnce.Empty()

		// Assert
		actual := args.Map{"result": sso.Value() != "" || sso.IsInitialized()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty mismatch", actual)
	})
}
