package coretaskinfotests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretaskinfo"
)

func Test_Cov_ExcludingOptions_SetSecure_Nil(t *testing.T) {
	var opt *coretaskinfo.ExcludingOptions
	result := opt.SetSecure()
	if result == nil || !result.IsSecureText {
		t.Error("expected secure")
	}
}

func Test_Cov_ExcludingOptions_SetPlainText_Nil(t *testing.T) {
	var opt *coretaskinfo.ExcludingOptions
	result := opt.SetPlainText()
	if result == nil || result.IsSecureText {
		t.Error("expected plain text")
	}
}

func Test_Cov_ExcludingOptions_ClonePtr_Nil(t *testing.T) {
	var opt *coretaskinfo.ExcludingOptions
	result := opt.ClonePtr()
	if result == nil {
		t.Error("expected non-nil")
	}
}

func Test_Cov_ExcludingOptions_IsEmpty(t *testing.T) {
	opt := &coretaskinfo.ExcludingOptions{}
	if !opt.IsEmpty() {
		t.Error("expected empty")
	}
	opt.IsExcludeRootName = true
	if opt.IsEmpty() {
		t.Error("expected not empty")
	}
}

func Test_Cov_ExcludingOptions_IsZero(t *testing.T) {
	opt := &coretaskinfo.ExcludingOptions{}
	if !opt.IsZero() {
		t.Error("expected zero")
	}
}

func Test_Cov_ExcludingOptions_AllIncludes(t *testing.T) {
	var opt *coretaskinfo.ExcludingOptions
	if !opt.IsIncludeRootName() {
		t.Error("nil should include root name")
	}
	if !opt.IsIncludeDescription() {
		t.Error("nil should include desc")
	}
	if !opt.IsIncludeUrl() {
		t.Error("nil should include url")
	}
	if !opt.IsIncludeHintUrl() {
		t.Error("nil should include hint url")
	}
	if !opt.IsIncludeErrorUrl() {
		t.Error("nil should include error url")
	}
	if !opt.IsIncludeExampleUrl() {
		t.Error("nil should include example url")
	}
	if !opt.IsIncludeSingleExample() {
		t.Error("nil should include single example")
	}
	if !opt.IsIncludeExamples() {
		t.Error("nil should include examples")
	}
	if !opt.IsIncludeAdditionalErrorWrap() {
		t.Error("nil should include additional error wrap")
	}
	if !opt.IsIncludePayloads() {
		t.Error("nil should include payloads")
	}
}

func Test_Cov_ExcludingOptions_Clone(t *testing.T) {
	opt := coretaskinfo.ExcludingOptions{IsExcludeRootName: true}
	cloned := opt.Clone()
	if !cloned.IsExcludeRootName {
		t.Error("expected cloned")
	}
}

func Test_Cov_ExcludingOptions_ToPtr_ToNonPtr(t *testing.T) {
	opt := coretaskinfo.ExcludingOptions{IsExcludeRootName: true}
	p := opt.ToPtr()
	if p == nil {
		t.Error("expected non-nil")
	}
	np := opt.ToNonPtr()
	if !np.IsExcludeRootName {
		t.Error("expected copied")
	}
}
