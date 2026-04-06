package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Migrated from Coverage02_Serializer_test.go ──

func Test_C02_Serializer_Apply(t *testing.T) {
	// Arrange
	r := corejson.Serialize.Apply("hello")

	// Act
	actual := args.Map{"result": r.HasError() || r.JsonString() != `"hello"`}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C02_Serializer_StringsApply(t *testing.T) {
	// Arrange
	r := corejson.Serialize.StringsApply([]string{"a", "b"})

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
}

func Test_C02_Serializer_FromBytes(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromBytes([]byte(`"test"`))

	// Act
	actual := args.Map{"hasError": r.HasError()}

	// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
}

func Test_C02_Serializer_FromStrings(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromStrings([]string{"a"})

	// Act
	actual := args.Map{"hasError": r.HasError()}

	// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
}

func Test_C02_Serializer_FromStringsSpread(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromStringsSpread("a", "b")

	// Act
	actual := args.Map{"hasError": r.HasError()}

	// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
}

func Test_C02_Serializer_FromString(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromString("hello")

	// Act
	actual := args.Map{"hasError": r.HasError()}

	// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
}

func Test_C02_Serializer_FromInteger(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromInteger(42)

	// Act
	actual := args.Map{"hasError": r.HasError()}

	// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
}

func Test_C02_Serializer_FromInteger64(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromInteger64(64)

	// Act
	actual := args.Map{"hasError": r.HasError()}

	// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
}

func Test_C02_Serializer_FromBool(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromBool(true)

	// Act
	actual := args.Map{"result": r.HasError() || r.JsonString() != "true"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C02_Serializer_FromIntegers(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromIntegers([]int{1, 2, 3})

	// Act
	actual := args.Map{"hasError": r.HasError()}

	// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
}

func Test_C02_Serializer_UsingAnyPtr(t *testing.T) {
	// Arrange
	r := corejson.Serialize.UsingAnyPtr("x")

	// Act
	actual := args.Map{"hasError": r.HasError()}

	// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
	ch := make(chan int)
	r2 := corejson.Serialize.UsingAnyPtr(ch)
	actual := args.Map{"result": r2.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C02_Serializer_UsingAny(t *testing.T) {
	// Arrange
	r := corejson.Serialize.UsingAny("x")

	// Act
	actual := args.Map{"hasError": r.HasError()}

	// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
}

func Test_C02_Serializer_Raw(t *testing.T) {
	// Arrange
	b, err := corejson.Serialize.Raw("x")

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C02_Serializer_Marshal(t *testing.T) {
	// Arrange
	b, err := corejson.Serialize.Marshal("x")

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C02_Serializer_ToBytesErr(t *testing.T) {
	// Arrange
	b, err := corejson.Serialize.ToBytesErr("x")

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C02_Serializer_ToBytesMust(t *testing.T) {
	// Arrange
	b := corejson.Serialize.ToBytesMust("x")

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C02_Serializer_ToSafeBytesMust(t *testing.T) {
	// Arrange
	b := corejson.Serialize.ToSafeBytesMust("x")

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C02_Serializer_ToSafeBytesSwallowErr(t *testing.T) {
	// Arrange
	b := corejson.Serialize.ToSafeBytesSwallowErr("x")

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C02_Serializer_ToBytesSwallowErr(t *testing.T) {
	// Arrange
	b := corejson.Serialize.ToBytesSwallowErr("x")

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C02_Serializer_ToString(t *testing.T) {
	// Arrange
	s := corejson.Serialize.ToString("hello")

	// Act
	actual := args.Map{"result": s != `"hello"`}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C02_Serializer_ToStringMust(t *testing.T) {
	// Arrange
	s := corejson.Serialize.ToStringMust("x")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_C02_Serializer_ToStringErr(t *testing.T) {
	// Arrange
	s, err := corejson.Serialize.ToStringErr("x")

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C02_Serializer_ToPrettyStringErr(t *testing.T) {
	// Arrange
	s, err := corejson.Serialize.ToPrettyStringErr(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C02_Serializer_ToPrettyStringIncludingErr(t *testing.T) {
	// Arrange
	s := corejson.Serialize.ToPrettyStringIncludingErr(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_C02_Serializer_Pretty(t *testing.T) {
	// Arrange
	s := corejson.Serialize.Pretty(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_C02_Serializer_ApplyMust(t *testing.T) {
	// Arrange
	r := corejson.Serialize.ApplyMust("x")

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
}
