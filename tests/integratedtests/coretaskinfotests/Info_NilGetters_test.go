package coretaskinfotests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretaskinfo"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Info nil-safe getters ──

func Test_Info_NilGetters(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info

	// Act
	actual := args.Map{
		"name":        info.Name(),
		"isNull":      info.IsNull(),
		"isDefined":   info.IsDefined(),
		"hasAnyName":  info.HasAnyName(),
		"isEmpty":     info.IsEmpty(),
		"hasAnyItem":  info.HasAnyItem(),
		"safeName":    info.SafeName(),
		"safeDesc":    info.SafeDescription(),
		"safeUrl":     info.SafeUrl(),
		"safeHintUrl": info.SafeHintUrl(),
		"safeErrUrl":  info.SafeErrorUrl(),
		"safeExUrl":   info.SafeExampleUrl(),
		"safeCh":      info.SafeChainingExample(),
	}

	// Assert
	expected := args.Map{
		"name": "", "isNull": true, "isDefined": false,
		"hasAnyName": false, "isEmpty": true, "hasAnyItem": false,
		"safeName": "", "safeDesc": "", "safeUrl": "",
		"safeHintUrl": "", "safeErrUrl": "", "safeExUrl": "", "safeCh": "",
	}
	expected.ShouldBeEqual(t, 0, "Info returns nil -- nil getters", actual)
}

// ── Info with values ──

func Test_Info_DefinedGetters(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{
		RootName:      "test-task",
		Description:   "desc",
		Url:           "http://url",
		HintUrl:       "http://hint",
		ErrorUrl:      "http://err",
		ExampleUrl:    "http://ex",
		SingleExample: "example1",
		Examples:      []string{"e1", "e2"},
	}

	// Act
	actual := args.Map{
		"name":           info.Name(),
		"isNull":         info.IsNull(),
		"isDefined":      info.IsDefined(),
		"hasAnyName":     info.HasAnyName(),
		"hasDesc":        info.HasDescription(),
		"hasUrl":         info.HasUrl(),
		"hasHintUrl":     info.HasHintUrl(),
		"hasErrorUrl":    info.HasErrorUrl(),
		"hasExampleUrl":  info.HasExampleUrl(),
		"hasChainingEx":  info.HasChainingExample(),
		"hasExamples":    info.HasExamples(),
		"hasExcludeOpts": info.HasExcludeOptions(),
	}

	// Assert
	expected := args.Map{
		"name": "test-task", "isNull": false, "isDefined": true,
		"hasAnyName": true, "hasDesc": true, "hasUrl": true,
		"hasHintUrl": true, "hasErrorUrl": true, "hasExampleUrl": true,
		"hasChainingEx": true, "hasExamples": true, "hasExcludeOpts": false,
	}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- defined getters", actual)
}

func Test_Info_IsName(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{RootName: "test"}

	// Act
	actual := args.Map{
		"match":   info.IsName("test"),
		"noMatch": info.IsName("other"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- IsName", actual)
}

// ── IsInclude / IsExclude checks ──

func Test_Info_IsInclude_NoExcludeOptions(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{
		RootName:      "name",
		Description:   "desc",
		Url:           "url",
		HintUrl:       "hint",
		ErrorUrl:      "errUrl",
		ExampleUrl:    "exUrl",
		SingleExample: "single",
		Examples:      []string{"e1"},
	}

	// Act
	actual := args.Map{
		"includeRootName":    info.IsIncludeRootName(),
		"includeDesc":        info.IsIncludeDescription(),
		"includeUrl":         info.IsIncludeUrl(),
		"includeHintUrl":     info.IsIncludeHintUrl(),
		"includeErrorUrl":    info.IsIncludeErrorUrl(),
		"includeExampleUrl":  info.IsIncludeExampleUrl(),
		"includeSingleEx":    info.IsIncludeSingleExample(),
		"includeExamples":    info.IsIncludeExamples(),
		"includeAdditional":  info.IsIncludeAdditionalErrorWrap(),
		"isPlainText":        info.IsPlainText(),
		"isIncludePayloads":  info.IsIncludePayloads(),
		"isSecure":           info.IsSecure(),
		"isExcludePayload":   info.IsExcludePayload(),
		"excludeRootName":    info.IsExcludeRootName(),
		"excludeDescription": info.IsExcludeDescription(),
		"excludeUrl":         info.IsExcludeUrl(),
		"excludeHintUrl":     info.IsExcludeHintUrl(),
		"excludeErrorUrl":    info.IsExcludeErrorUrl(),
	}

	// Assert
	expected := args.Map{
		"includeRootName": true, "includeDesc": true, "includeUrl": true,
		"includeHintUrl": true, "includeErrorUrl": true, "includeExampleUrl": true,
		"includeSingleEx": true, "includeExamples": true, "includeAdditional": true,
		"isPlainText": true, "isIncludePayloads": true, "isSecure": false,
		"isExcludePayload": false, "excludeRootName": false, "excludeDescription": false,
		"excludeUrl": false, "excludeHintUrl": false, "excludeErrorUrl": false,
	}
	expected.ShouldBeEqual(t, 0, "Info returns empty -- IsInclude no exclude options", actual)
}

func Test_Info_NilIncludeAdditionalErrorWrap(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info

	// Act
	actual := args.Map{"result": info.IsIncludeAdditionalErrorWrap()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil Info IsIncludeAdditionalErrorWrap -- true", actual)
}

// ── Clone ──

func Test_Info_Clone(t *testing.T) {
	// Arrange
	info := coretaskinfo.Info{
		RootName:    "task",
		Description: "desc",
		Url:         "url",
	}
	cloned := info.Clone()

	// Act
	actual := args.Map{
		"name": cloned.RootName,
		"desc": cloned.Description,
		"url":  cloned.Url,
	}

	// Assert
	expected := args.Map{
		"name": "task",
		"desc": "desc",
		"url": "url",
	}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- Clone", actual)
}

func Test_Info_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info
	cloned := info.ClonePtr()

	// Act
	actual := args.Map{"isNil": cloned == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Info ClonePtr nil -- nil", actual)
}

func Test_Info_ClonePtr_Defined(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{RootName: "test"}
	cloned := info.ClonePtr()

	// Act
	actual := args.Map{"name": cloned.RootName}

	// Assert
	expected := args.Map{"name": "test"}
	expected.ShouldBeEqual(t, 0, "Info ClonePtr defined -- cloned", actual)
}

func Test_Info_ToPtr(t *testing.T) {
	// Arrange
	info := coretaskinfo.Info{RootName: "task"}
	ptr := info.ToPtr()

	// Act
	actual := args.Map{"name": ptr.RootName}

	// Assert
	expected := args.Map{"name": "task"}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- ToPtr", actual)
}

func Test_Info_ToNonPtr(t *testing.T) {
	// Arrange
	info := coretaskinfo.Info{RootName: "task"}
	nonPtr := info.ToNonPtr()

	// Act
	actual := args.Map{"name": nonPtr.RootName}

	// Assert
	expected := args.Map{"name": "task"}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- ToNonPtr", actual)
}

// ── SetSecure / SetPlain ──

func Test_Info_SetSecure_Nil(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info
	result := info.SetSecure()

	// Act
	actual := args.Map{"isSecure": result.IsSecure()}

	// Assert
	expected := args.Map{"isSecure": true}
	expected.ShouldBeEqual(t, 0, "Info SetSecure nil -- creates secure", actual)
}

func Test_Info_SetSecure_Defined(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{RootName: "task"}
	result := info.SetSecure()

	// Act
	actual := args.Map{
		"isSecure": result.IsSecure(),
		"name":     result.RootName,
	}

	// Assert
	expected := args.Map{
		"isSecure": true,
		"name": "task",
	}
	expected.ShouldBeEqual(t, 0, "Info SetSecure defined -- sets secure flag", actual)
}

func Test_Info_SetPlain_Nil(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info
	result := info.SetPlain()

	// Act
	actual := args.Map{"isPlain": result.IsPlainText()}

	// Assert
	expected := args.Map{"isPlain": true}
	expected.ShouldBeEqual(t, 0, "Info SetPlain nil -- creates plain", actual)
}

func Test_Info_SetPlain_Defined(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{RootName: "task"}
	info.SetSecure()
	result := info.SetPlain()

	// Act
	actual := args.Map{"isPlain": result.IsPlainText()}

	// Assert
	expected := args.Map{"isPlain": true}
	expected.ShouldBeEqual(t, 0, "Info SetPlain defined -- sets plain", actual)
}

// ── Options / ExamplesAsSlice ──

func Test_Info_Options_Nil(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info
	opts := info.Options()

	// Act
	actual := args.Map{"notNil": opts != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Info Options nil -- returns empty", actual)
}

func Test_Info_ExamplesAsSlice_Nil(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info
	slice := info.ExamplesAsSlice()

	// Act
	actual := args.Map{"notNil": slice != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Info ExamplesAsSlice nil -- returns empty", actual)
}

func Test_Info_ExamplesAsSlice_Defined(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{Examples: []string{"a", "b"}}
	slice := info.ExamplesAsSlice()

	// Act
	actual := args.Map{"len": slice.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- ExamplesAsSlice defined", actual)
}
