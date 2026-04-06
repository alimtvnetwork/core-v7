package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// SliceValidators — collection basics
// ==========================================

func Test_SliceValidators_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	actual := args.Map{"result": v.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should be empty", actual)
	actual := args.Map{"result": v.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// (nil receiver tests migrated to SliceValidators_NilReceiver_testcases.go)

func Test_SliceValidators_WithItems(t *testing.T) {
	v := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}
	actual := args.Map{"result": v.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	actual := args.Map{"result": v.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ==========================================
// SliceValidators.IsMatch / IsValid
// ==========================================

func Test_SliceValidators_IsMatch_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	actual := args.Map{"result": v.IsMatch(true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should match", actual)
}

func Test_SliceValidators_IsValid_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	actual := args.Map{"result": v.IsValid(true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty IsValid should be true", actual)
}

func Test_SliceValidators_IsMatch_AllPass(t *testing.T) {
	v := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				Condition:     corevalidator.DefaultDisabledCoreCondition,
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a", "b"},
				ExpectedLines: []string{"a", "b"},
			},
		},
	}
	actual := args.Map{"result": v.IsMatch(true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "matching validators should return true", actual)
}

func Test_SliceValidators_IsMatch_OneFails(t *testing.T) {
	v := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				Condition:     corevalidator.DefaultDisabledCoreCondition,
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
			{
				Condition:     corevalidator.DefaultDisabledCoreCondition,
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"x"},
				ExpectedLines: []string{"y"},
			},
		},
	}
	actual := args.Map{"result": v.IsMatch(true)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one failing validator should return false", actual)
}

// (nil receiver test migrated to SliceValidators_NilReceiver_testcases.go)

// ==========================================
// SliceValidators.VerifyAll
// ==========================================

func Test_SliceValidators_VerifyAll_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test"}
	err := v.VerifyAll("header", params, false)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

func Test_SliceValidators_VerifyAll_Pass(t *testing.T) {
	v := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				Condition:     corevalidator.DefaultDisabledCoreCondition,
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test", IsCaseSensitive: true}
	err := v.VerifyAll("header", params, false)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "matching should pass:", actual)
}

// ==========================================
// SliceValidators.VerifyAllError
// ==========================================

func Test_SliceValidators_VerifyAllError_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test"}
	err := v.VerifyAllError(params)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

// ==========================================
// SliceValidators.VerifyFirst
// ==========================================

func Test_SliceValidators_VerifyFirst_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test"}
	err := v.VerifyFirst(params, false)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

func Test_SliceValidators_VerifyFirst_Pass(t *testing.T) {
	v := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				Condition:     corevalidator.DefaultDisabledCoreCondition,
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test", IsCaseSensitive: true}
	err := v.VerifyFirst(params, false)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "matching should pass:", actual)
}

// ==========================================
// SliceValidators.VerifyUpto
// ==========================================

func Test_SliceValidators_VerifyUpto_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test"}
	err := v.VerifyUpto(false, false, 1, params)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

// ==========================================
// SliceValidators.SetActualOnAll
// ==========================================

func Test_SliceValidators_SetActualOnAll_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	// should not panic
	v.SetActualOnAll("a", "b")
}

// ==========================================
// SliceValidators.VerifyAllErrorUsingActual
// ==========================================

func Test_SliceValidators_VerifyAllErrorUsingActual_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test"}
	err := v.VerifyAllErrorUsingActual(params, "a")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}
