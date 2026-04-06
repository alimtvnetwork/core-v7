package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

type cov9Mini struct {
	Name string `json:"name"`
}

func Test_Cov9_New_Valid(t *testing.T) {
	r := corejson.New(map[string]string{"k": "v"})
	actual := args.Map{"noErr": !r.HasError(), "hasBytes": r.HasBytes()}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "New returns non-empty -- valid", actual)
}

func Test_Cov9_NewPtr_Nil(t *testing.T) {
	r := corejson.NewPtr(nil)
	actual := args.Map{"notNil": r != nil, "hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"notNil": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "NewPtr returns nil -- nil", actual)
}

func Test_Cov9_BytesCloneIf_True(t *testing.T) {
	original := []byte(`"hello"`)
	cloned := corejson.BytesCloneIf(true, original)
	actual := args.Map{"len": len(cloned), "notSame": &cloned[0] != &original[0]}
	expected := args.Map{"len": 7, "notSame": true}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf returns non-empty -- true", actual)
}

func Test_Cov9_Deserialize_BytesTo_String(t *testing.T) {
	s, err := corejson.Deserialize.BytesTo.String([]byte(`"hello"`))
	actual := args.Map{"noErr": err == nil, "val": s}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize.BytesTo.String returns correct value -- with args", actual)
}

func Test_Cov9_Deserialize_BytesTo_MapStringString(t *testing.T) {
	m, err := corejson.Deserialize.BytesTo.MapStringString([]byte(`{"k":"v"}`))
	actual := args.Map{"noErr": err == nil, "len": len(m)}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "Deserialize.BytesTo.MapStringString returns correct value -- with args", actual)
}

func Test_Cov9_Deserialize_BytesTo_StringMust(t *testing.T) {
	s := corejson.Deserialize.BytesTo.StringMust([]byte(`"hello"`))
	actual := args.Map{"val": s}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize.BytesTo.StringMust returns correct value -- with args", actual)
}

func Test_Cov9_Deserialize_BytesTo_IntegerMust(t *testing.T) {
	i := corejson.Deserialize.BytesTo.IntegerMust([]byte(`42`))
	actual := args.Map{"val": i}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "Deserialize.BytesTo.IntegerMust returns correct value -- with args", actual)
}

func Test_Cov9_Serialize_Raw_Valid(t *testing.T) {
	b, err := corejson.Serialize.Raw(map[string]string{"k": "v"})
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize.Raw returns non-empty -- valid", actual)
}

func Test_Cov9_Empty_Result(t *testing.T) {
	r := corejson.Empty.Result()
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.Result returns empty -- with args", actual)
}

func Test_Cov9_Result_Clone_NilPtr(t *testing.T) {
	var r *corejson.Result
	cloned := r.ClonePtr(true)
	actual := args.Map{"nil": cloned == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_Cov9_CastAny_FromToDefault(t *testing.T) {
	// CastAny.FromToDefault serializes source then deserializes into target
	var casted map[string]string
	err := corejson.CastAny.FromToDefault(map[string]string{"k": "v"}, &casted)
	actual := args.Map{"noErr": err == nil, "len": len(casted)}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "CastAny.FromToDefault returns correct value -- with args", actual)
}

func Test_Cov9_Pretty_Bytes_SafeDefault(t *testing.T) {
	pretty := corejson.Pretty.Bytes.SafeDefault([]byte(`{"k":"v"}`))
	actual := args.Map{"notEmpty": pretty != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.Bytes.SafeDefault returns correct value -- with args", actual)
}

func Test_Cov9_Pretty_String_SafeDefault(t *testing.T) {
	pretty := corejson.Pretty.String.SafeDefault(`{"k":"v"}`)
	actual := args.Map{"notEmpty": pretty != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.String.SafeDefault returns correct value -- with args", actual)
}

func Test_Cov9_AnyTo_JsonString(t *testing.T) {
	jsonString := corejson.AnyTo.JsonString(cov9Mini{Name: "alice"})
	actual := args.Map{"notEmpty": jsonString != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.JsonString returns correct value -- with args", actual)
}

func Test_Cov9_AnyTo_PrettyStringWithError(t *testing.T) {
	pretty, err := corejson.AnyTo.PrettyStringWithError(map[string]string{"k": "v"})
	actual := args.Map{"notEmpty": pretty != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.PrettyStringWithError returns error -- with args", actual)
}
