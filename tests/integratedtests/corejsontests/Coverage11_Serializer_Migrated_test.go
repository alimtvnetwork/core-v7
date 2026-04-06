package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Migrated from Coverage02_Serializer_test.go ──

func Test_C02_Serializer_Apply(t *testing.T) {
	r := corejson.Serialize.Apply("hello")
	actual := args.Map{"result": r.HasError() || r.JsonString() != `"hello"`}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C02_Serializer_StringsApply(t *testing.T) {
	r := corejson.Serialize.StringsApply([]string{"a", "b"})
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
}

func Test_C02_Serializer_FromBytes(t *testing.T) {
	r := corejson.Serialize.FromBytes([]byte(`"test"`))
	actual := args.Map{"hasError": r.HasError()}
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
}

func Test_C02_Serializer_FromStrings(t *testing.T) {
	r := corejson.Serialize.FromStrings([]string{"a"})
	actual := args.Map{"hasError": r.HasError()}
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
}

func Test_C02_Serializer_FromStringsSpread(t *testing.T) {
	r := corejson.Serialize.FromStringsSpread("a", "b")
	actual := args.Map{"hasError": r.HasError()}
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
}

func Test_C02_Serializer_FromString(t *testing.T) {
	r := corejson.Serialize.FromString("hello")
	actual := args.Map{"hasError": r.HasError()}
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
}

func Test_C02_Serializer_FromInteger(t *testing.T) {
	r := corejson.Serialize.FromInteger(42)
	actual := args.Map{"hasError": r.HasError()}
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
}

func Test_C02_Serializer_FromInteger64(t *testing.T) {
	r := corejson.Serialize.FromInteger64(64)
	actual := args.Map{"hasError": r.HasError()}
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
}

func Test_C02_Serializer_FromBool(t *testing.T) {
	r := corejson.Serialize.FromBool(true)
	actual := args.Map{"result": r.HasError() || r.JsonString() != "true"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C02_Serializer_FromIntegers(t *testing.T) {
	r := corejson.Serialize.FromIntegers([]int{1, 2, 3})
	actual := args.Map{"hasError": r.HasError()}
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
}

func Test_C02_Serializer_UsingAnyPtr(t *testing.T) {
	r := corejson.Serialize.UsingAnyPtr("x")
	actual := args.Map{"hasError": r.HasError()}
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
	ch := make(chan int)
	r2 := corejson.Serialize.UsingAnyPtr(ch)
	actual := args.Map{"result": r2.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C02_Serializer_UsingAny(t *testing.T) {
	r := corejson.Serialize.UsingAny("x")
	actual := args.Map{"hasError": r.HasError()}
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "result has no error", actual)
}

func Test_C02_Serializer_Raw(t *testing.T) {
	b, err := corejson.Serialize.Raw("x")
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C02_Serializer_Marshal(t *testing.T) {
	b, err := corejson.Serialize.Marshal("x")
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C02_Serializer_ToBytesErr(t *testing.T) {
	b, err := corejson.Serialize.ToBytesErr("x")
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C02_Serializer_ToBytesMust(t *testing.T) {
	b := corejson.Serialize.ToBytesMust("x")
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C02_Serializer_ToSafeBytesMust(t *testing.T) {
	b := corejson.Serialize.ToSafeBytesMust("x")
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C02_Serializer_ToSafeBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToSafeBytesSwallowErr("x")
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C02_Serializer_ToBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToBytesSwallowErr("x")
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C02_Serializer_ToString(t *testing.T) {
	s := corejson.Serialize.ToString("hello")
	actual := args.Map{"result": s != `"hello"`}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C02_Serializer_ToStringMust(t *testing.T) {
	s := corejson.Serialize.ToStringMust("x")
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_C02_Serializer_ToStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToStringErr("x")
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C02_Serializer_ToPrettyStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToPrettyStringErr(map[string]int{"a": 1})
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C02_Serializer_ToPrettyStringIncludingErr(t *testing.T) {
	s := corejson.Serialize.ToPrettyStringIncludingErr(map[string]int{"a": 1})
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_C02_Serializer_Pretty(t *testing.T) {
	s := corejson.Serialize.Pretty(map[string]int{"a": 1})
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_C02_Serializer_ApplyMust(t *testing.T) {
	r := corejson.Serialize.ApplyMust("x")
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
}
