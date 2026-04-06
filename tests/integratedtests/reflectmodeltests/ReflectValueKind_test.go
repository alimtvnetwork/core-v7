package reflectmodeltests

import (
	"testing"

	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ===== ReflectValueKind Tests =====

func Test_InvalidReflectValueKindModel(t *testing.T) {
	rvk := reflectmodel.InvalidReflectValueKindModel("test error")

	actual := args.Map{"result": rvk.IsValid}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsValid = false", actual)

	actual := args.Map{"result": rvk.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasError() = true", actual)

	actual := args.Map{"result": rvk.Error.Error() != "test error"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Error =, want", rvk.Error.Error(), "test error", actual)
}

// Note: All nil receiver tests migrated to ReflectValueKind_NilReceiver_testcases.go

func Test_ReflectValueKind_NilReceiver(t *testing.T) {
	for caseIndex, tc := range reflectValueKindNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

func Test_ReflectValueKind_IsInvalid_NotValid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}

	actual := args.Map{"result": rvk.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsInvalid() = true when IsValid=false", actual)
}

func Test_ReflectValueKind_IsEmptyError_NoError(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{}

	actual := args.Map{"result": rvk.IsEmptyError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsEmptyError() = true when no error", actual)
}

func Test_ReflectValueKind_PkgPath_NotValid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}

	got := rvk.PkgPath()
	actual := args.Map{"result": got != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected PkgPath() = empty when IsValid=false", actual)
}

func Test_ReflectValueKind_TypeName_NotValid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}

	got := rvk.TypeName()
	actual := args.Map{"result": got != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected TypeName() = empty when IsValid=false", actual)
}
