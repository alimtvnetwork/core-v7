package coretaskinfotests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretaskinfo"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Test: Info.Default creation
// ==========================================

func Test_Info_Default_Verification(t *testing.T) {
	for caseIndex, testCase := range infoDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")

		// Act
		info := coretaskinfo.New.Info.Default(nameStr, descStr, urlStr)
		actual := args.Map{
			"name":      info.SafeName(),
			"desc":      info.SafeDescription(),
			"url":       info.SafeUrl(),
			"isNull":    fmt.Sprintf("%v", info.IsNull()),
			"isDefined": fmt.Sprintf("%v", info.IsDefined()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Info.Examples with examples
// ==========================================

func Test_Info_ExamplesWithItems_Verification(t *testing.T) {
	for caseIndex, testCase := range infoExamplesWithItemsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")
		examples, _ := input.GetAsStrings("examples")

		// Act
		info := coretaskinfo.New.Info.Examples(nameStr, descStr, urlStr, examples...)
		actual := args.Map{
			"name":         info.SafeName(),
			"desc":         info.SafeDescription(),
			"url":          info.SafeUrl(),
			"isNull":       fmt.Sprintf("%v", info.IsNull()),
			"isDefined":    fmt.Sprintf("%v", info.IsDefined()),
			"hasExamples":  fmt.Sprintf("%v", info.HasExamples()),
			"exampleCount": fmt.Sprintf("%d", len(info.Examples)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Info.Examples with no examples
// ==========================================

func Test_Info_ExamplesEmpty_Verification(t *testing.T) {
	for caseIndex, testCase := range infoExamplesEmptyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")

		// Act
		info := coretaskinfo.New.Info.Examples(nameStr, descStr, urlStr)
		actual := args.Map{
			"name":         info.SafeName(),
			"desc":         info.SafeDescription(),
			"url":          info.SafeUrl(),
			"isNull":       fmt.Sprintf("%v", info.IsNull()),
			"isDefined":    fmt.Sprintf("%v", info.IsDefined()),
			"hasExamples":  fmt.Sprintf("%v", info.HasExamples()),
			"exampleCount": fmt.Sprintf("%d", len(info.Examples)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Nil info — SafeName
// ==========================================

func Test_Info_Nil_SafeName_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilSafeNameTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		result := info.SafeName()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================
// Test: Nil info — SafeDescription
// ==========================================

func Test_Info_Nil_SafeDescription_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilSafeDescriptionTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		result := info.SafeDescription()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================
// Test: Nil info — SafeUrl
// ==========================================

func Test_Info_Nil_SafeUrl_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilSafeUrlTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		result := info.SafeUrl()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================
// Test: Nil info — SafeHintUrl
// ==========================================

func Test_Info_Nil_SafeHintUrl_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilSafeHintUrlTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		result := info.SafeHintUrl()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================
// Test: Nil info — SafeErrorUrl
// ==========================================

func Test_Info_Nil_SafeErrorUrl_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilSafeErrorUrlTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		result := info.SafeErrorUrl()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================
// Test: Nil info — SafeExampleUrl
// ==========================================

func Test_Info_Nil_SafeExampleUrl_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilSafeExampleUrlTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		result := info.SafeExampleUrl()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================
// Test: Nil info — NullCheck
// ==========================================

func Test_Info_Nil_NullCheck_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilNullCheckTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		actual := args.Map{
			"isNull":    fmt.Sprintf("%v", info.IsNull()),
			"isDefined": fmt.Sprintf("%v", info.IsDefined()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Nil info — EmptyCheck
// ==========================================

func Test_Info_Nil_EmptyCheck_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilEmptyCheckTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		actual := args.Map{
			"isEmpty":    fmt.Sprintf("%v", info.IsEmpty()),
			"hasAnyItem": fmt.Sprintf("%v", info.HasAnyItem()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Nil info — ClonePtr
// ==========================================

func Test_Info_Nil_ClonePtr_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilClonePtrTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		cloned := info.ClonePtr()
		result := fmt.Sprintf("%v", cloned == nil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================
// Test: Nil info — PrettyJsonString
// ==========================================

func Test_Info_Nil_PrettyJsonString_Verification(t *testing.T) {
	for caseIndex, testCase := range infoNilPrettyJsonTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		result := info.PrettyJsonString()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================
// Test: Secure.Default creation
// ==========================================

func Test_Info_SecureDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSecureDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")

		// Act
		info := coretaskinfo.New.Info.Secure.Default(nameStr, descStr, urlStr)
		actual := args.Map{
			"name":             info.SafeName(),
			"desc":             info.SafeDescription(),
			"url":              info.SafeUrl(),
			"isSecure":         fmt.Sprintf("%v", info.IsSecure()),
			"isPlainText":      fmt.Sprintf("%v", info.IsPlainText()),
			"isExcludePayload": fmt.Sprintf("%v", info.IsExcludePayload()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Secure.NameDescUrlExamples creation
// ==========================================

func Test_Info_SecureExamples_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSecureExamplesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")
		examples, _ := input.GetAsStrings("examples")

		// Act
		info := coretaskinfo.New.Info.Secure.NameDescUrlExamples(
			nameStr, descStr, urlStr, examples...)
		actual := args.Map{
			"name":             info.SafeName(),
			"isSecure":         fmt.Sprintf("%v", info.IsSecure()),
			"isPlainText":      fmt.Sprintf("%v", info.IsPlainText()),
			"isExcludePayload": fmt.Sprintf("%v", info.IsExcludePayload()),
			"exampleCount":     fmt.Sprintf("%d", len(info.Examples)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: SetSecure on nil
// ==========================================

func Test_Info_SetSecureOnNil_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSetSecureOnNilTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		result := info.SetSecure()
		actual := args.Map{
			"isSecure":    fmt.Sprintf("%v", result.IsSecure()),
			"isPlainText": fmt.Sprintf("%v", result.IsPlainText()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: SetSecure on existing plain info
// ==========================================

func Test_Info_SetSecureOnExisting_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSetSecureOnExistingTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")

		// Act
		info := coretaskinfo.New.Info.Plain.Default(nameStr, "d", "u")
		info.SetSecure()
		actual := args.Map{
			"isSecure":    fmt.Sprintf("%v", info.IsSecure()),
			"isPlainText": fmt.Sprintf("%v", info.IsPlainText()),
			"name":        info.SafeName(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Plain.Default creation
// ==========================================

func Test_Info_PlainDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range infoPlainDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")

		// Act
		info := coretaskinfo.New.Info.Plain.Default(nameStr, descStr, urlStr)
		actual := args.Map{
			"name":              info.SafeName(),
			"desc":              info.SafeDescription(),
			"url":               info.SafeUrl(),
			"isSecure":          fmt.Sprintf("%v", info.IsSecure()),
			"isPlainText":       fmt.Sprintf("%v", info.IsPlainText()),
			"isIncludePayloads": fmt.Sprintf("%v", info.IsIncludePayloads()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Plain.AllUrlExamples creation
// ==========================================

func Test_Info_PlainAllUrlExamples_Verification(t *testing.T) {
	for caseIndex, testCase := range infoPlainAllUrlExamplesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")
		hintUrl, _ := input.GetAsString("hintUrl")
		errorUrl, _ := input.GetAsString("errorUrl")
		examples, _ := input.GetAsStrings("examples")

		// Act
		info := coretaskinfo.New.Info.Plain.AllUrlExamples(
			nameStr, descStr, urlStr, hintUrl, errorUrl, examples...)
		actual := args.Map{
			"name":         info.SafeName(),
			"desc":         info.SafeDescription(),
			"url":          info.SafeUrl(),
			"hintUrl":      info.SafeHintUrl(),
			"errorUrl":     info.SafeErrorUrl(),
			"isSecure":     fmt.Sprintf("%v", info.IsSecure()),
			"isPlainText":  fmt.Sprintf("%v", info.IsPlainText()),
			"exampleCount": fmt.Sprintf("%d", len(info.Examples)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: SetPlain on nil
// ==========================================

func Test_Info_SetPlainOnNil_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSetPlainOnNilTestCases {
		// Arrange
		var info *coretaskinfo.Info

		// Act
		result := info.SetPlain()
		actual := args.Map{
			"isSecure":    fmt.Sprintf("%v", result.IsSecure()),
			"isPlainText": fmt.Sprintf("%v", result.IsPlainText()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Serialize Default round-trip
// ==========================================

func Test_Info_SerializeDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSerializeDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")

		// Act
		original := coretaskinfo.New.Info.Default(nameStr, descStr, urlStr)
		jsonBytes := original.JsonPtr().Bytes
		deserialized, err := coretaskinfo.New.Info.Deserialized(jsonBytes)
		actual := args.Map{
			"name":     deserialized.SafeName(),
			"desc":     deserialized.SafeDescription(),
			"url":      deserialized.SafeUrl(),
			"noError":  fmt.Sprintf("%v", err == nil),
			"isSecure": fmt.Sprintf("%v", deserialized.IsSecure()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Serialize Secure round-trip
// ==========================================

func Test_Info_SerializeSecure_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSerializeSecureTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")

		// Act
		original := coretaskinfo.New.Info.Secure.Default(nameStr, descStr, urlStr)
		jsonBytes := original.JsonPtr().Bytes
		deserialized, err := coretaskinfo.New.Info.Deserialized(jsonBytes)
		actual := args.Map{
			"name":     deserialized.SafeName(),
			"desc":     deserialized.SafeDescription(),
			"url":      deserialized.SafeUrl(),
			"noError":  fmt.Sprintf("%v", err == nil),
			"isSecure": fmt.Sprintf("%v", deserialized.IsSecure()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Serialize with examples round-trip
// ==========================================

func Test_Info_SerializeExamples_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSerializeExamplesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")
		examples, _ := input.GetAsStrings("examples")

		// Act
		original := coretaskinfo.New.Info.Examples(nameStr, descStr, urlStr, examples...)
		jsonBytes := original.JsonPtr().Bytes
		deserialized, err := coretaskinfo.New.Info.Deserialized(jsonBytes)
		actual := args.Map{
			"name":         deserialized.SafeName(),
			"noError":      fmt.Sprintf("%v", err == nil),
			"exampleCount": fmt.Sprintf("%d", len(deserialized.Examples)),
		}
		for i, ex := range deserialized.Examples {
			actual[fmt.Sprintf("example%d", i)] = ex
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Serialize with all URLs round-trip
// ==========================================

func Test_Info_SerializeAllUrls_Verification(t *testing.T) {
	for caseIndex, testCase := range infoSerializeAllUrlsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")
		hintUrl, _ := input.GetAsString("hintUrl")
		errorUrl, _ := input.GetAsString("errorUrl")

		// Act
		original := coretaskinfo.New.Info.Plain.AllUrl(nameStr, descStr, urlStr, hintUrl, errorUrl)
		jsonBytes := original.JsonPtr().Bytes
		deserialized, err := coretaskinfo.New.Info.Deserialized(jsonBytes)
		actual := args.Map{
			"name":     deserialized.SafeName(),
			"url":      deserialized.SafeUrl(),
			"hintUrl":  deserialized.SafeHintUrl(),
			"errorUrl": deserialized.SafeErrorUrl(),
			"noError":  fmt.Sprintf("%v", err == nil),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Clone
// ==========================================

func Test_Info_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range infoCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nameStr, _ := input.GetAsString("name")
		descStr, _ := input.GetAsString("desc")
		urlStr, _ := input.GetAsString("url")
		newName, _ := input.GetAsString("newName")

		// Act
		original := coretaskinfo.New.Info.Default(nameStr, descStr, urlStr)
		cloned := original.Clone()
		cloned.RootName = newName
		actual := args.Map{
			"originalName": original.SafeName(),
			"clonedName":   cloned.SafeName(),
			"clonedDesc":   cloned.SafeDescription(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Field checks — populated
// ==========================================

func Test_Info_FieldChecks_Populated_Verification(t *testing.T) {
	for caseIndex, testCase := range infoFieldCheckPopulatedTestCases {
		// Arrange
		info := coretaskinfo.New.Info.Secure.AllUrlExamples(
			"name", "desc",
			"url", "hint", "err",
			"ex1", "ex2",
		)
		info.SingleExample = "single"

		// Act
		actual := args.Map{
			"hasRootName":        fmt.Sprintf("%v", info.HasRootName()),
			"hasDescription":     fmt.Sprintf("%v", info.HasDescription()),
			"hasUrl":             fmt.Sprintf("%v", info.HasUrl()),
			"hasHintUrl":         fmt.Sprintf("%v", info.HasHintUrl()),
			"hasErrorUrl":        fmt.Sprintf("%v", info.HasErrorUrl()),
			"hasExamples":        fmt.Sprintf("%v", info.HasExamples()),
			"hasChainingExample": fmt.Sprintf("%v", info.HasChainingExample()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Field checks — empty
// ==========================================

func Test_Info_FieldChecks_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range infoFieldCheckEmptyTestCases {
		// Arrange
		info := &coretaskinfo.Info{}

		// Act
		actual := args.Map{
			"hasRootName":        fmt.Sprintf("%v", info.HasRootName()),
			"hasDescription":     fmt.Sprintf("%v", info.HasDescription()),
			"hasUrl":             fmt.Sprintf("%v", info.HasUrl()),
			"hasHintUrl":         fmt.Sprintf("%v", info.HasHintUrl()),
			"hasErrorUrl":        fmt.Sprintf("%v", info.HasErrorUrl()),
			"hasExamples":        fmt.Sprintf("%v", info.HasExamples()),
			"hasChainingExample": fmt.Sprintf("%v", info.HasChainingExample()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
