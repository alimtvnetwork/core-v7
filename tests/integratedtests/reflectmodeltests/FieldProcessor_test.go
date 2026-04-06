package reflectmodeltests

import (
	"reflect"
	"testing"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ===== FieldProcessor Tests =====

func Test_FieldProcessor_IsFieldType_Match(t *testing.T) {
	fp := newFieldProcessor("Name", 0)
	actual := args.Map{"result": fp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "failed to create FieldProcessor for Name", actual)

	actual := args.Map{"result": fp.IsFieldType(reflect.TypeOf(""))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsFieldType(string) = true for Name field", actual)
}

func Test_FieldProcessor_IsFieldType_NoMatch(t *testing.T) {
	fp := newFieldProcessor("Name", 0)
	actual := args.Map{"result": fp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "failed to create FieldProcessor for Name", actual)

	actual := args.Map{"result": fp.IsFieldType(reflect.TypeOf(0))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsFieldType(int) = false for Name (string) field", actual)
}

// Note: IsFieldType nil receiver test migrated to FieldProcessor_NilReceiver_testcases.go

func Test_FieldProcessor_IsFieldKind_Match(t *testing.T) {
	fp := newFieldProcessor("Age", 1)
	actual := args.Map{"result": fp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "failed to create FieldProcessor for Age", actual)

	actual := args.Map{"result": fp.IsFieldKind(reflect.Int)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsFieldKind(Int) = true for Age field", actual)
}

func Test_FieldProcessor_IsFieldKind_NoMatch(t *testing.T) {
	fp := newFieldProcessor("Age", 1)
	actual := args.Map{"result": fp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "failed to create FieldProcessor for Age", actual)

	actual := args.Map{"result": fp.IsFieldKind(reflect.String)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsFieldKind(String) = false for Age (int) field", actual)
}

// Note: IsFieldKind nil receiver test migrated to FieldProcessor_NilReceiver_testcases.go

func Test_FieldProcessor_NilReceiver(t *testing.T) {
	for caseIndex, tc := range fieldProcessorNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

func Test_FieldProcessor_BoolField(t *testing.T) {
	fp := newFieldProcessor("Active", 2)
	actual := args.Map{"result": fp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "failed to create FieldProcessor for Active", actual)

	actual := args.Map{"result": fp.IsFieldKind(reflect.Bool)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected Active field to be Bool kind", actual)

	actual := args.Map{"result": fp.IsFieldType(reflect.TypeOf(true))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected Active field to match bool type", actual)
}

func Test_FieldProcessor_FieldData(t *testing.T) {
	fp := newFieldProcessor("Name", 0)
	actual := args.Map{"result": fp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "failed to create FieldProcessor for Name", actual)

	actual := args.Map{"result": fp.Name != "Name"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Name =, want", fp.Name, "Name", actual)

	actual := args.Map{"result": fp.Index != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Index =, want 0", actual)

	actual := args.Map{"result": fp.Field.Name != "Name"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Field.Name =, want", fp.Field.Name, "Name", actual)
}
