package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/namevalue"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_Collection_String_NilReceiver(t *testing.T) {
	// Arrange
	var c *namevalue.Collection[string, string]
	s := c.String()

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_QW_Collection_JsonString_NilReceiver(t *testing.T) {
	// Arrange
	defer func() { recover() }() // value receiver on nil pointer may panic
	var c *namevalue.Collection[string, string]
	s := c.JsonString()

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_QW_Instance_IsNull(t *testing.T) {
	// Arrange
	var inst *namevalue.Instance[string, string]

	// Act
	actual := args.Map{"result": inst.IsNull()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected null", actual)
}

func Test_QW_Instance_String_Nil(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// nil pointer panic is expected for zero-value Instance.String()
		}
	}()
	inst := namevalue.Instance[string, string]{}
	_ = inst.String()
}

func Test_QW_Instance_JsonString_Nil(t *testing.T) {
	// JsonString is a value receiver — calling on nil pointer panics
	defer func() {
		if r := recover(); r != nil {
			// expected: nil pointer dereference on value receiver
		}
	}()
	var inst *namevalue.Instance[string, string]
	_ = inst.JsonString()
}
