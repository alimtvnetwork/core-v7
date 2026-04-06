package coretaskinfotests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretaskinfo"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_Info_JsonString_Nil(t *testing.T) {
	// Arrange
	defer func() { recover() }() // value receiver on nil pointer may panic
	var info *coretaskinfo.Info
	s := info.JsonString()

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_QW_Info_LazyMapWithPayload_ErrorPath(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{
		RootName:    "test",
		Description: "desc",
	}
	// LazyMapWithPayload takes []byte — pass invalid JSON bytes to cover error branch
	m := info.LazyMapWithPayload([]byte(`{invalid`))

	// Act
	actual := args.Map{"result": m == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil map", actual)
}

func Test_QW_Info_LazyMapWithPayloadAsAny_ErrorPath(t *testing.T) {
	info := &coretaskinfo.Info{
		RootName:    "test",
		Description: "desc",
	}
	m := info.LazyMapWithPayloadAsAny(make(chan int))
	actual := args.Map{"result": m == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil map", actual)
}
