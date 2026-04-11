package coredynamictests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ReflectInterfaceVal non-ptr (line 20) ──

func Test_ReflectInterfaceVal_NonPtr_I29(t *testing.T) {
	// Arrange / Act
	result := coredynamic.ReflectInterfaceVal(42)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": 42}
	actual.ShouldBeEqual(t, 1, "ReflectInterfaceVal non-ptr", expected)
}

// ── ReflectSetFromTo: []byte → struct (line 159-167) ──

func Test_ReflectSetFromTo_BytesToStruct_I29(t *testing.T) {
	// Arrange
	type sample struct {
		Name string `json:"name"`
	}
	data, _ := json.Marshal(sample{Name: "test"})
	var target sample

	// Act
	err := coredynamic.ReflectSetFromTo(data, &target)

	// Assert
	actual := args.Map{
		"hasError": err != nil,
		"name":     target.Name,
	}
	expected := args.Map{
		"hasError": false,
		"name":     "test",
	}
	actual.ShouldBeEqual(t, 1, "ReflectSetFromTo bytes to struct", expected)
}

// ── ReflectSetFromTo: struct → *[]byte (line 174-180) ──

func Test_ReflectSetFromTo_StructToBytes_I29(t *testing.T) {
	// Arrange
	type sample struct {
		Name string `json:"name"`
	}
	from := sample{Name: "test"}
	var target []byte

	// Act
	err := coredynamic.ReflectSetFromTo(from, &target)

	// Assert
	actual := args.Map{
		"hasError":   err != nil,
		"hasContent": len(target) > 0,
	}
	expected := args.Map{
		"hasError":   false,
		"hasContent": true,
	}
	actual.ShouldBeEqual(t, 1, "ReflectSetFromTo struct to bytes", expected)
}

// ── SafeZeroSet non-pointer (line 18) ──

func Test_SafeZeroSet_NonPointer_I29(t *testing.T) {
	// Arrange
	val := reflect.ValueOf(42)

	// Act — should not panic on non-pointer
	coredynamic.SafeZeroSet(val)

	// Assert
	actual := args.Map{"noPanic": true}
	expected := args.Map{"noPanic": true}
	actual.ShouldBeEqual(t, 1, "SafeZeroSet non-pointer", expected)
}
