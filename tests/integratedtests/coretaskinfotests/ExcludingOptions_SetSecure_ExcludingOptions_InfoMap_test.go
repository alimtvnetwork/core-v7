package coretaskinfotests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretaskinfo"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_ExcludingOptions_SetSecure_Nil(t *testing.T) {
	// Arrange
	var opt *coretaskinfo.ExcludingOptions
	result := opt.SetSecure()

	// Act
	actual := args.Map{"result": result == nil || !result.IsSecureText}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected secure", actual)
}

func Test_ExcludingOptions_SetPlainText_Nil(t *testing.T) {
	// Arrange
	var opt *coretaskinfo.ExcludingOptions
	result := opt.SetPlainText()

	// Act
	actual := args.Map{"result": result == nil || result.IsSecureText}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected plain text", actual)
}

func Test_ExcludingOptions_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var opt *coretaskinfo.ExcludingOptions
	result := opt.ClonePtr()

	// Act
	actual := args.Map{"result": result == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_ExcludingOptions_IsEmpty(t *testing.T) {
	// Arrange
	opt := &coretaskinfo.ExcludingOptions{}

	// Act
	actual := args.Map{"result": opt.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	opt.IsExcludeRootName = true
	actual = args.Map{"result": opt.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)
}

func Test_ExcludingOptions_IsZero(t *testing.T) {
	// Arrange
	opt := &coretaskinfo.ExcludingOptions{}

	// Act
	actual := args.Map{"result": opt.IsZero()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected zero", actual)
}

func Test_ExcludingOptions_AllIncludes(t *testing.T) {
	// Arrange
	var opt *coretaskinfo.ExcludingOptions

	// Act
	actual := args.Map{"result": opt.IsIncludeRootName()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include root name", actual)
	actual = args.Map{"result": opt.IsIncludeDescription()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include desc", actual)
	actual = args.Map{"result": opt.IsIncludeUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include url", actual)
	actual = args.Map{"result": opt.IsIncludeHintUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include hint url", actual)
	actual = args.Map{"result": opt.IsIncludeErrorUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include error url", actual)
	actual = args.Map{"result": opt.IsIncludeExampleUrl()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include example url", actual)
	actual = args.Map{"result": opt.IsIncludeSingleExample()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include single example", actual)
	actual = args.Map{"result": opt.IsIncludeExamples()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include examples", actual)
	actual = args.Map{"result": opt.IsIncludeAdditionalErrorWrap()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include additional error wrap", actual)
	actual = args.Map{"result": opt.IsIncludePayloads()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should include payloads", actual)
}

func Test_ExcludingOptions_Clone(t *testing.T) {
	// Arrange
	opt := coretaskinfo.ExcludingOptions{IsExcludeRootName: true}
	cloned := opt.Clone()

	// Act
	actual := args.Map{"result": cloned.IsExcludeRootName}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected cloned", actual)
}

func Test_ExcludingOptions_ToPtr_ToNonPtr(t *testing.T) {
	// Arrange
	opt := coretaskinfo.ExcludingOptions{IsExcludeRootName: true}
	p := opt.ToPtr()

	// Act
	actual := args.Map{"result": p == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	np := opt.ToNonPtr()
	actual = args.Map{"result": np.IsExcludeRootName}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected copied", actual)
}
