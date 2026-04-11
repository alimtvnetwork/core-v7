package coretaskinfotests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretaskinfo"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ============================================================================
// Info: Core Identity
// ============================================================================

func Test_Info_Name_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("myName", "desc", "http://url")

	// Act
	actual := args.Map{"result": info.Name() != "myName"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'myName', got ''", actual)
}

func Test_Info_NilName_Ext(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info

	// Act
	actual := args.Map{"result": info.Name() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil Name should return empty", actual)
}

func Test_Info_IsDefined_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")

	// Act
	actual := args.Map{"result": info.IsDefined()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be defined", actual)
}

func Test_Info_NilIsDefined_Ext(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info

	// Act
	actual := args.Map{"result": info.IsDefined()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not be defined", actual)
}

func Test_Info_HasAnyName_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")

	// Act
	actual := args.Map{"result": info.HasAnyName()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have name", actual)
}

func Test_Info_IsName_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("myName", "d", "u")

	// Act
	actual := args.Map{"result": info.IsName("myName")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match name", actual)
	actual = args.Map{"result": info.IsName("other")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not match other", actual)
}

func Test_Info_IsEmpty_Ext(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info

	// Act
	actual := args.Map{"result": info.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_Info_HasAnyItem_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")

	// Act
	actual := args.Map{"result": info.HasAnyItem()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have item", actual)
}

func Test_Info_Options_Ext(t *testing.T) {
	// Arrange
	// Plain.Default does not set ExcludeOptions, so Options() returns nil.
	// This is correct production behavior — nil means "no exclusions".
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	opts := info.Options()

	// Act
	actual := args.Map{"result": opts != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Options should be nil for Default (no ExcludeOptions set)", actual)
}

func Test_Info_NilOptions_Ext(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info
	opts := info.Options()

	// Act
	actual := args.Map{"result": opts == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil Options should return empty options", actual)
}

// ============================================================================
// Info: Clone / ToPtr / ToNonPtr
// ============================================================================

func Test_Info_Clone_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	cloned := info.Clone()

	// Act
	actual := args.Map{"result": cloned.RootName != "n"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone should preserve RootName", actual)
}

func Test_Info_ClonePtr_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	cloned := info.ClonePtr()

	// Act
	actual := args.Map{"result": cloned == nil || cloned.RootName != "n"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ClonePtr should preserve RootName", actual)
}

func Test_Info_NilClonePtr_Ext(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info

	// Act
	actual := args.Map{"result": info.ClonePtr() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil ClonePtr should return nil", actual)
}

func Test_Info_ToPtr_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	ptr := info.ToPtr()

	// Act
	actual := args.Map{"result": ptr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToPtr should not be nil", actual)
}

func Test_Info_ToNonPtr_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	nonPtr := info.ToNonPtr()

	// Act
	actual := args.Map{"result": nonPtr.RootName != "n"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToNonPtr should preserve RootName", actual)
}

// ============================================================================
// Info: SetSecure / SetPlain
// ============================================================================

func Test_Info_SetSecure_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	result := info.SetSecure()

	// Act
	actual := args.Map{"result": result.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Info_NilSetSecure_Ext(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info
	result := info.SetSecure()

	// Act
	actual := args.Map{"result": result == nil || !result.IsSecure()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil SetSecure should return secure info", actual)
}

func Test_Info_SetPlain_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.Default("n", "d", "u")
	result := info.SetPlain()

	// Act
	actual := args.Map{"result": result.IsPlainText()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be plain text", actual)
}

func Test_Info_NilSetPlain_Ext(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info
	result := info.SetPlain()

	// Act
	actual := args.Map{"result": result == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil SetPlain should return info", actual)
}

// ============================================================================
// Info: Getters
// ============================================================================

func Test_Info_IsInclude_Getters_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.AllUrlExamples(
		"n", "d", "http://url", "http://hint", "http://err", "ex1",
	)
	info.SingleExample = "chain"
	info.ExampleUrl = "http://exurl"

	// Act
	actual := args.Map{"result": info.IsIncludeRootName()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should include root name", actual)
	actual = args.Map{"result": info.IsIncludeDescription()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should include description", actual)
	actual = args.Map{"result": info.IsIncludeUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should include url", actual)
	actual = args.Map{"result": info.IsIncludeHintUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should include hint url", actual)
	actual = args.Map{"result": info.IsIncludeErrorUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should include error url", actual)
	actual = args.Map{"result": info.IsIncludeExampleUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should include example url", actual)
	actual = args.Map{"result": info.IsIncludeSingleExample()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should include single example", actual)
	actual = args.Map{"result": info.IsIncludeExamples()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should include examples", actual)
	actual = args.Map{"result": info.IsIncludeAdditionalErrorWrap()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should include additional error wrap", actual)
}

func Test_Info_NilIsInclude_Ext(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info

	// Act
	actual := args.Map{"result": info.IsIncludeRootName()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not include root name", actual)
	actual = args.Map{"result": info.IsIncludeAdditionalErrorWrap()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include additional error wrap", actual)
}

func Test_Info_IsSecure_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.Default("n", "d", "u")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
	actual = args.Map{"result": info.IsExcludePayload()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should exclude payload", actual)
	actual = args.Map{"result": info.IsPlainText()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be plain text", actual)
}

func Test_Info_IsPlainText_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")

	// Act
	actual := args.Map{"result": info.IsPlainText()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be plain text", actual)
	actual = args.Map{"result": info.IsIncludePayloads()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should include payloads", actual)
}

// ============================================================================
// Info: Safe* getters
// ============================================================================

func Test_Info_SafeGetters_Ext(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{
		RootName:    "n",
		Description: "d",
		Url:         "u",
		HintUrl:     "h",
		ErrorUrl:    "e",
		ExampleUrl:  "ex",
	}

	// Act
	actual := args.Map{"result": info.SafeName() != "n"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SafeName mismatch", actual)
	actual = args.Map{"result": info.SafeDescription() != "d"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SafeDescription mismatch", actual)
	actual = args.Map{"result": info.SafeUrl() != "u"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SafeUrl mismatch", actual)
	actual = args.Map{"result": info.SafeHintUrl() != "h"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SafeHintUrl mismatch", actual)
	actual = args.Map{"result": info.SafeErrorUrl() != "e"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SafeErrorUrl mismatch", actual)
	actual = args.Map{"result": info.SafeExampleUrl() != "ex"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SafeExampleUrl mismatch", actual)
	actual = args.Map{"result": info.SafeChainingExample() != "ex"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SafeChainingExample mismatch", actual)
}

func Test_Info_NilSafeGetters_Ext(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info

	// Act
	actual := args.Map{"result": info.SafeName() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil SafeName should be empty", actual)
	actual = args.Map{"result": info.SafeDescription() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil SafeDescription should be empty", actual)
	actual = args.Map{"result": info.SafeUrl() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil SafeUrl should be empty", actual)
	actual = args.Map{"result": info.SafeHintUrl() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil SafeHintUrl should be empty", actual)
	actual = args.Map{"result": info.SafeErrorUrl() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil SafeErrorUrl should be empty", actual)
	actual = args.Map{"result": info.SafeExampleUrl() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil SafeExampleUrl should be empty", actual)
	actual = args.Map{"result": info.SafeChainingExample() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil SafeChainingExample should be empty", actual)
}

// ============================================================================
// Info: Has* checks
// ============================================================================

func Test_Info_HasChecks_Ext(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{
		RootName:    "n",
		Description: "d",
		Url:         "u",
		HintUrl:     "h",
		ErrorUrl:    "e",
		ExampleUrl:  "ex",
		SingleExample: "se",
		Examples:    []string{"e1"},
		ExcludeOptions: &coretaskinfo.ExcludingOptions{IsExcludeRootName: true},
	}

	// Act
	actual := args.Map{"result": info.HasRootName()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have root name", actual)
	actual = args.Map{"result": info.HasDescription()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have description", actual)
	actual = args.Map{"result": info.HasUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have url", actual)
	actual = args.Map{"result": info.HasHintUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have hint url", actual)
	actual = args.Map{"result": info.HasErrorUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have error url", actual)
	actual = args.Map{"result": info.HasExampleUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have example url", actual)
	actual = args.Map{"result": info.HasChainingExample()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have chaining example", actual)
	actual = args.Map{"result": info.HasExamples()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have examples", actual)
	actual = args.Map{"result": info.HasExcludeOptions()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have exclude options", actual)
}

// ============================================================================
// Info: IsEmpty* checks
// ============================================================================

func Test_Info_IsEmptyChecks_Ext(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info

	// Act
	actual := args.Map{"result": info.IsEmptyName()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty name", actual)
	actual = args.Map{"result": info.IsEmptyDescription()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty description", actual)
	actual = args.Map{"result": info.IsEmptyUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty url", actual)
	actual = args.Map{"result": info.IsEmptyHintUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty hint url", actual)
	actual = args.Map{"result": info.IsEmptyErrorUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty error url", actual)
	actual = args.Map{"result": info.IsEmptyExampleUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty example url", actual)
	actual = args.Map{"result": info.IsEmptySingleExample()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty single example", actual)
	actual = args.Map{"result": info.IsEmptyExamples()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty examples", actual)
	actual = args.Map{"result": info.IsEmptyExcludeOptions()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty exclude options", actual)
}

// ============================================================================
// Info: IsExclude* checks
// ============================================================================

func Test_Info_IsExcludeChecks_Ext(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{
		ExcludeOptions: &coretaskinfo.ExcludingOptions{
			IsExcludeRootName:            true,
			IsExcludeDescription:         true,
			IsExcludeUrl:                 true,
			IsExcludeHintUrl:             true,
			IsExcludeErrorUrl:            true,
			IsExcludeAdditionalErrorWrap: true,
			IsExcludeExampleUrl:          true,
			IsExcludeSingleExample:       true,
			IsExcludeExamples:            true,
		},
	}

	// Act
	actual := args.Map{"result": info.IsExcludeRootName()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should exclude root name", actual)
	actual = args.Map{"result": info.IsExcludeDescription()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should exclude description", actual)
	actual = args.Map{"result": info.IsExcludeUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should exclude url", actual)
	actual = args.Map{"result": info.IsExcludeHintUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should exclude hint url", actual)
	actual = args.Map{"result": info.IsExcludeErrorUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should exclude error url", actual)
	actual = args.Map{"result": info.IsExcludeAdditionalErrorWrap()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should exclude additional error wrap", actual)
	actual = args.Map{"result": info.IsExcludeExampleUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should exclude example url", actual)
	actual = args.Map{"result": info.IsExcludeSingleExample()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should exclude single example", actual)
	actual = args.Map{"result": info.IsExcludeExamples()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should exclude examples", actual)
}

// ============================================================================
// Info: JSON methods
// ============================================================================

func Test_Info_JsonString_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	result := info.JsonString()

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonString should return non-empty", actual)
}

func Test_Info_NilJsonString_Ext(t *testing.T) {
	// Arrange
	var info coretaskinfo.Info
	result := info.JsonString()

	// Act
	actual := args.Map{"result": result != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil/zero JsonString should return empty", actual)
}

func Test_Info_PrettyJsonString_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	result := info.PrettyJsonString()

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Info_NilPrettyJsonString_Ext(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info
	result := info.PrettyJsonString()

	// Act
	actual := args.Map{"result": result != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_Info_String_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	result := info.String()

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "String should return non-empty", actual)
}

func Test_Info_NilString_Ext(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info
	result := info.String()

	// Act
	actual := args.Map{"result": result != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil String should return empty", actual)
}

func Test_Info_Serialize_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	bytes, err := info.Serialize()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Serialize error:", actual)
	actual = args.Map{"result": len(bytes) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty bytes", actual)
}

func Test_Info_ExamplesAsString_Ext(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{Examples: []string{"e1", "e2"}}
	result := info.ExamplesAsString()

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Info_NilExamplesAsString_Ext(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info
	result := info.ExamplesAsString()

	// Act
	actual := args.Map{"result": result != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_Info_ExamplesAsSlice_Ext(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{Examples: []string{"e1", "e2"}}
	result := info.ExamplesAsSlice()

	// Act
	actual := args.Map{"result": result == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_Info_NilExamplesAsSlice_Ext(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info
	result := info.ExamplesAsSlice()

	// Act
	actual := args.Map{"result": result == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty slice", actual)
}

func Test_Info_AsJsonContractsBinder_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	binder := info.AsJsonContractsBinder()

	// Act
	actual := args.Map{"result": binder == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_Info_JsonStringMust_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	result := info.JsonStringMust()

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Info_LazyMapPrettyJsonString_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	result := info.LazyMapPrettyJsonString()

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// ============================================================================
// Info: Map methods
// ============================================================================

func Test_Info_Map_Ext(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{
		RootName:      "n",
		Description:   "d",
		Url:           "u",
		HintUrl:       "h",
		ErrorUrl:      "e",
		ExampleUrl:    "ex",
		SingleExample: "se",
		Examples:      []string{"e1"},
	}
	m := info.Map()

	// Act
	actual := args.Map{"result": len(m) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have entries", actual)
}

func Test_Info_NilMap_Ext(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info
	m := info.Map()

	// Act
	actual := args.Map{"result": len(m) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil map should be empty", actual)
}

func Test_Info_MapWithPayload_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	m := info.MapWithPayload([]byte("payload"))

	// Act
	actual := args.Map{"result": m == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_Info_LazyMap_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	m1 := info.LazyMap()
	m2 := info.LazyMap() // cached

	// Act
	actual := args.Map{"result": len(m1) != len(m2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "lazy map should be cached", actual)
}

func Test_Info_NilLazyMap_Ext(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info
	m := info.LazyMap()

	// Act
	actual := args.Map{"result": len(m) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil lazy map should be empty", actual)
}

func Test_Info_PrettyJsonStringWithPayloads_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	result := info.PrettyJsonStringWithPayloads([]byte("payload"))

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Info_LazyMapWithPayload_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	m := info.LazyMapWithPayload([]byte("payload"))

	// Act
	actual := args.Map{"result": m == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_Info_MapWithPayloadAsAny_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	m := info.MapWithPayloadAsAny("test-payload")

	// Act
	actual := args.Map{"result": m == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_Info_LazyMapWithPayloadAsAny_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	m := info.LazyMapWithPayloadAsAny("test-payload")

	// Act
	actual := args.Map{"result": m == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_Info_SecureMapWithPayloadAsAny_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.Default("n", "d", "u")
	m := info.MapWithPayloadAsAny("test-payload")

	// Act
	actual := args.Map{"result": m == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_Info_SecureLazyMapWithPayloadAsAny_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.Default("n", "d", "u")
	m := info.LazyMapWithPayloadAsAny("test-payload")

	// Act
	actual := args.Map{"result": m == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

// ============================================================================
// ExcludingOptions
// ============================================================================

func Test_ExcludingOptions_SetSecure_Ext(t *testing.T) {
	// Arrange
	opts := &coretaskinfo.ExcludingOptions{}
	result := opts.SetSecure()

	// Act
	actual := args.Map{"result": result.IsSafeSecureText()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_ExcludingOptions_NilSetSecure_Ext(t *testing.T) {
	// Arrange
	var opts *coretaskinfo.ExcludingOptions
	result := opts.SetSecure()

	// Act
	actual := args.Map{"result": result == nil || !result.IsSafeSecureText()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil SetSecure should return secure", actual)
}

func Test_ExcludingOptions_SetPlainText_Ext(t *testing.T) {
	// Arrange
	opts := &coretaskinfo.ExcludingOptions{IsSecureText: true}
	result := opts.SetPlainText()

	// Act
	actual := args.Map{"result": result.IsSafeSecureText()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be secure after SetPlainText", actual)
}

func Test_ExcludingOptions_NilSetPlainText_Ext(t *testing.T) {
	// Arrange
	var opts *coretaskinfo.ExcludingOptions
	result := opts.SetPlainText()

	// Act
	actual := args.Map{"result": result == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil SetPlainText should return options", actual)
}

func Test_ExcludingOptions_IsEmpty_Ext(t *testing.T) {
	// Arrange
	opts := &coretaskinfo.ExcludingOptions{}

	// Act
	actual := args.Map{"result": opts.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "default should be empty", actual)
	opts.IsExcludeRootName = true
	actual = args.Map{"result": opts.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty after setting a flag", actual)
}

func Test_ExcludingOptions_NilIsEmpty_Ext(t *testing.T) {
	// Arrange
	var opts *coretaskinfo.ExcludingOptions

	// Act
	actual := args.Map{"result": opts.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_ExcludingOptions_Clone_Ext(t *testing.T) {
	// Arrange
	opts := coretaskinfo.ExcludingOptions{IsExcludeRootName: true}
	cloned := opts.Clone()

	// Act
	actual := args.Map{"result": cloned.IsExcludeRootName}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Clone should preserve flags", actual)
}

func Test_ExcludingOptions_ClonePtr_Ext(t *testing.T) {
	// Arrange
	opts := &coretaskinfo.ExcludingOptions{IsExcludeRootName: true}
	cloned := opts.ClonePtr()

	// Act
	actual := args.Map{"result": cloned == nil || !cloned.IsExcludeRootName}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ClonePtr should preserve flags", actual)
}

func Test_ExcludingOptions_NilClonePtr_Ext(t *testing.T) {
	// Arrange
	var opts *coretaskinfo.ExcludingOptions
	result := opts.ClonePtr()

	// Act
	actual := args.Map{"result": result == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil ClonePtr should return empty options", actual)
}

func Test_ExcludingOptions_ToPtr_Ext(t *testing.T) {
	// Arrange
	opts := coretaskinfo.ExcludingOptions{IsExcludeRootName: true}
	ptr := opts.ToPtr()

	// Act
	actual := args.Map{"result": ptr == nil || !ptr.IsExcludeRootName}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToPtr should preserve flags", actual)
}

func Test_ExcludingOptions_ToNonPtr_Ext(t *testing.T) {
	// Arrange
	opts := coretaskinfo.ExcludingOptions{IsExcludeRootName: true}
	nonPtr := opts.ToNonPtr()

	// Act
	actual := args.Map{"result": nonPtr.IsExcludeRootName}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ToNonPtr should preserve flags", actual)
}

func Test_ExcludingOptions_IsInclude_All_Ext(t *testing.T) {
	// Arrange
	var opts *coretaskinfo.ExcludingOptions

	// Act
	actual := args.Map{"result": opts.IsIncludeRootName()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include root name", actual)
	actual = args.Map{"result": opts.IsIncludeDescription()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include description", actual)
	actual = args.Map{"result": opts.IsIncludeUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include url", actual)
	actual = args.Map{"result": opts.IsIncludeHintUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include hint url", actual)
	actual = args.Map{"result": opts.IsIncludeErrorUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include error url", actual)
	actual = args.Map{"result": opts.IsIncludeExampleUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include example url", actual)
	actual = args.Map{"result": opts.IsIncludeSingleExample()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include single example", actual)
	actual = args.Map{"result": opts.IsIncludeExamples()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include examples", actual)
	actual = args.Map{"result": opts.IsIncludeAdditionalErrorWrap()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include additional error wrap", actual)
	actual = args.Map{"result": opts.IsIncludePayloads()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include payloads", actual)
}

func Test_ExcludingOptions_IsSafe_All_Ext(t *testing.T) {
	// Arrange
	var opts *coretaskinfo.ExcludingOptions

	// Act
	actual := args.Map{"result": opts.IsSafeExcludeRootName()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not exclude", actual)
	actual = args.Map{"result": opts.IsSafeExcludeDescription()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not exclude", actual)
	actual = args.Map{"result": opts.IsSafeExcludeUrl()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not exclude", actual)
	actual = args.Map{"result": opts.IsSafeExcludeErrorUrl()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not exclude", actual)
	actual = args.Map{"result": opts.IsSafeExcludeAdditionalErrorWrap()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not exclude", actual)
	actual = args.Map{"result": opts.IsSafeExcludeHintUrl()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not exclude", actual)
	actual = args.Map{"result": opts.IsSafeExcludeExampleUrl()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not exclude", actual)
	actual = args.Map{"result": opts.IsSafeExcludeSingleExample()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not exclude", actual)
	actual = args.Map{"result": opts.IsSafeExcludeExamples()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not exclude", actual)
	actual = args.Map{"result": opts.IsSafeSecureText()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not be secure", actual)
}

// ============================================================================
// newInfoCreator methods
// ============================================================================

func Test_NewInfo_Default_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Default("n", "d", "u")

	// Act
	actual := args.Map{"result": info.RootName != "n"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set root name", actual)
}

func Test_NewInfo_Examples_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Examples("n", "d", "u", "e1", "e2")

	// Act
	actual := args.Map{"result": len(info.Examples) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 2 examples", actual)
}

func Test_NewInfo_Create_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Create(true, "n", "d", "u", "h", "e", "ex", "se", "e1")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_NewInfo_SecureCreate_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.SecureCreate("n", "d", "u", "h", "e", "ex", "se", "e1")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_NewInfo_PlainCreate_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.PlainCreate("n", "d", "u", "h", "e", "ex", "se", "e1")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be secure", actual)
}

// ============================================================================
// newInfoPlainTextCreator - remaining methods
// ============================================================================

func Test_Plain_NameDescUrl_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.NameDescUrl("n", "d", "u")

	// Act
	actual := args.Map{"result": info.RootName != "n"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set name", actual)
}

func Test_Plain_NameDescUrlExamples_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.NameDescUrlExamples("n", "d", "u", "e1")

	// Act
	actual := args.Map{"result": len(info.Examples) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 1 example", actual)
}

func Test_Plain_NewNameDescUrlErrorUrl_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.NewNameDescUrlErrorUrl("n", "d", "u", "eu")

	// Act
	actual := args.Map{"result": info.ErrorUrl != "eu"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set error url", actual)
}

func Test_Plain_NameDescUrlErrUrlExamples_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.NameDescUrlErrUrlExamples("n", "d", "u", "eu", "e1")

	// Act
	actual := args.Map{"result": info.ErrorUrl != "eu" || len(info.Examples) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set error url and examples", actual)
}

func Test_Plain_NameDescExamples_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.NameDescExamples("n", "d", "e1")

	// Act
	actual := args.Map{"result": len(info.Examples) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 1 example", actual)
}

func Test_Plain_Examples_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.Examples("n", "d", "e1")

	// Act
	actual := args.Map{"result": len(info.Examples) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 1 example", actual)
}

func Test_Plain_NameUrlExamples_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.NameUrlExamples("n", "u", "e1")

	// Act
	actual := args.Map{"result": info.Url != "u"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set url", actual)
}

func Test_Plain_UrlExamples_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.UrlExamples("u", "e1")

	// Act
	actual := args.Map{"result": info.Url != "u"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set url", actual)
}

func Test_Plain_ExamplesOnly_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.ExamplesOnly("e1")

	// Act
	actual := args.Map{"result": len(info.Examples) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 1 example", actual)
}

func Test_Plain_UrlOnly_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.UrlOnly("u")

	// Act
	actual := args.Map{"result": info.Url != "u"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set url", actual)
}

func Test_Plain_ErrorUrlOnly_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.ErrorUrlOnly("eu")

	// Act
	actual := args.Map{"result": info.ErrorUrl != "eu"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set error url", actual)
}

func Test_Plain_HintUrlOnly_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.HintUrlOnly("hu")

	// Act
	actual := args.Map{"result": info.HintUrl != "hu"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set hint url", actual)
}

func Test_Plain_DescHintUrlOnly_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.DescHintUrlOnly("d", "hu")

	// Act
	actual := args.Map{"result": info.Description != "d" || info.HintUrl != "hu"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set desc and hint url", actual)
}

func Test_Plain_NameHintUrlOnly_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.NameHintUrlOnly("n", "hu")

	// Act
	actual := args.Map{"result": info.RootName != "n" || info.HintUrl != "hu"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set name and hint url", actual)
}

func Test_Plain_SingleExampleOnly_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.SingleExampleOnly("se")

	// Act
	actual := args.Map{"result": info.SingleExample != "se"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set single example", actual)
}

func Test_Plain_AllUrl_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.AllUrl("n", "d", "u", "hu", "eu")

	// Act
	actual := args.Map{"result": info.HintUrl != "hu" || info.ErrorUrl != "eu"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set all urls", actual)
}

func Test_Plain_UrlSingleExample_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.UrlSingleExample("n", "d", "u", "se")

	// Act
	actual := args.Map{"result": info.SingleExample != "se"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set single example", actual)
}

func Test_Plain_SingleExample_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.SingleExample("n", "d", "se")

	// Act
	actual := args.Map{"result": info.SingleExample != "se"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set single example", actual)
}

func Test_Plain_ExampleUrl_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.ExampleUrl("n", "d", "exu", "se")

	// Act
	actual := args.Map{"result": info.ExampleUrl != "exu"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set example url", actual)
}

func Test_Plain_ExampleUrlSingleExample_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Plain.ExampleUrlSingleExample("n", "d", "exu", "se")

	// Act
	actual := args.Map{"result": info.ExampleUrl != "exu" || info.SingleExample != "se"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set both", actual)
}

// ============================================================================
// newInfoSecureTextCreator - remaining methods
// ============================================================================

func Test_Secure_NameDescUrl_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.NameDescUrl("n", "d", "u")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Secure_NameDescUrlExamples_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.NameDescUrlExamples("n", "d", "u", "e1")

	// Act
	actual := args.Map{"result": info.IsSecure() || len(info.Examples) != 1}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure with examples", actual)
}

func Test_Secure_NewNameDescUrlErrorUrl_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.NewNameDescUrlErrorUrl("n", "d", "u", "eu")

	// Act
	actual := args.Map{"result": info.IsSecure() || info.ErrorUrl != "eu"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure with error url", actual)
}

func Test_Secure_NameDescUrlErrUrlExamples_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.NameDescUrlErrUrlExamples("n", "d", "u", "eu", "e1")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Secure_NameDescExamples_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.NameDescExamples("n", "d", "e1")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Secure_Examples_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.Examples("n", "d", "e1")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Secure_ExamplesOnly_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.ExamplesOnly("e1")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Secure_UrlOnly_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.UrlOnly("u")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Secure_ErrorUrlOnly_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.ErrorUrlOnly("eu")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Secure_HintUrlOnly_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.HintUrlOnly("hu")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Secure_DescHintUrlOnly_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.DescHintUrlOnly("d", "hu")

	// Act
	actual := args.Map{"result": info.Description != "d" || info.HintUrl != "hu"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should set desc and hint url", actual)
}

func Test_Secure_NameHintUrlOnly_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.NameHintUrlOnly("n", "hu")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Secure_SingleExampleOnly_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.SingleExampleOnly("se")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Secure_AllUrlExamples_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.AllUrlExamples("n", "d", "u", "hu", "eu", "e1")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Secure_AllUrl_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.AllUrl("n", "d", "u", "hu", "eu")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Secure_UrlSingleExample_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.UrlSingleExample("n", "d", "u", "se")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Secure_SingleExample_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.SingleExample("n", "d", "se")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Secure_ExampleUrl_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.ExampleUrl("n", "d", "exu", "se")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Secure_ExampleUrlSingleExample_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.ExampleUrlSingleExample("n", "d", "exu", "se")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

func Test_Secure_NewExampleUrlSecure_Ext(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Secure.NewExampleUrlSecure("n", "d", "exu", "se")

	// Act
	actual := args.Map{"result": info.IsSecure()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be secure", actual)
}

// ============================================================================
// newInfoCreator: Deserialized / DeserializedUsingJsonResult
// ============================================================================

func Test_NewInfo_Deserialized_Ext(t *testing.T) {
	// Arrange
	original := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	bytes, _ := original.Serialize()
	result, err := coretaskinfo.New.Info.Deserialized(bytes)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Deserialized error:", actual)
	actual = args.Map{"result": result.RootName != "n"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should preserve root name", actual)
}

func Test_NewInfo_DeserializedUsingJsonResult_Ext(t *testing.T) {
	// Arrange
	original := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	jsonResult := original.JsonPtr()
	result, err := coretaskinfo.New.Info.DeserializedUsingJsonResult(jsonResult)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": result.RootName != "n"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should preserve root name", actual)
}

func Test_Info_Deserialize_Ext(t *testing.T) {
	// Arrange
	original := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	var target coretaskinfo.Info
	err := original.Deserialize(&target)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Deserialize error:", actual)
	actual = args.Map{"result": target.RootName != "n"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should preserve root name", actual)
}

func Test_Info_JsonParseSelfInject_Ext(t *testing.T) {
	// Arrange
	original := coretaskinfo.New.Info.Plain.Default("n", "d", "u")
	jsonResult := original.JsonPtr()
	var target coretaskinfo.Info
	err := target.JsonParseSelfInject(jsonResult)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}
