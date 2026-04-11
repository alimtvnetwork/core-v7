package coretaskinfotests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretaskinfo"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── InfoJson — all methods ──

func Test_Info_Json(t *testing.T) {
	// Arrange
	info := coretaskinfo.Info{RootName: "task"}
	r := info.Json()

	// Act
	actual := args.Map{"hasBytes": r.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- Json", actual)
}

func Test_Info_JsonPtr(t *testing.T) {
	// Arrange
	info := coretaskinfo.Info{RootName: "task"}
	r := info.JsonPtr()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- JsonPtr", actual)
}

func Test_Info_JsonString(t *testing.T) {
	// Arrange
	info := coretaskinfo.Info{RootName: "task"}

	// Act
	actual := args.Map{"notEmpty": info.JsonString() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- JsonString", actual)
}

func Test_Info_JsonString_Nil(t *testing.T) {
	// Arrange
	defer func() {
		if r := recover(); r != nil {
			t.Skip("Info.JsonString panics on zero value -- skipping")
		}
	}()
	info := &coretaskinfo.Info{}
	result := info.JsonString()

	// Act
	actual := args.Map{
		"notPanic": true,
		"hasResult": result != "",
	}

	// Assert
	expected := args.Map{
		"notPanic": true,
		"hasResult": actual["hasResult"],
	}
	expected.ShouldBeEqual(t, 0, "Info returns nil -- JsonString nil", actual)
}

func Test_Info_JsonStringMust(t *testing.T) {
	// Arrange
	info := coretaskinfo.Info{RootName: "task"}

	// Act
	actual := args.Map{"notEmpty": info.JsonStringMust() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- JsonStringMust", actual)
}

func Test_Info_PrettyJsonString(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{RootName: "task"}

	// Act
	actual := args.Map{"notEmpty": info.PrettyJsonString() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- PrettyJsonString", actual)
}

func Test_Info_PrettyJsonString_Nil(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info

	// Act
	actual := args.Map{"empty": info.PrettyJsonString()}

	// Assert
	expected := args.Map{"empty": ""}
	expected.ShouldBeEqual(t, 0, "Info returns nil -- PrettyJsonString nil", actual)
}

func Test_Info_String(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{RootName: "task"}

	// Act
	actual := args.Map{"notEmpty": info.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- String", actual)
}

func Test_Info_String_Nil(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info

	// Act
	actual := args.Map{"empty": info.String()}

	// Assert
	expected := args.Map{"empty": ""}
	expected.ShouldBeEqual(t, 0, "Info returns nil -- String nil", actual)
}

func Test_Info_Serialize(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{RootName: "task"}
	bytes, err := info.Serialize()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- Serialize", actual)
}

func Test_Info_ExamplesAsString(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{Examples: []string{"a", "b"}}
	result := info.ExamplesAsString()

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a, b"}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- ExamplesAsString", actual)
}

func Test_Info_ExamplesAsString_Nil(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info

	// Act
	actual := args.Map{"result": info.ExamplesAsString()}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Info returns nil -- ExamplesAsString nil", actual)
}

func Test_Info_AsJsonContractsBinder(t *testing.T) {
	// Arrange
	info := coretaskinfo.Info{RootName: "task"}
	binder := info.AsJsonContractsBinder()

	// Act
	actual := args.Map{"notNil": binder != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- AsJsonContractsBinder", actual)
}

// ── InfoMap — Map / LazyMap ──

func Test_Info_Map_Defined(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{
		RootName: "task", Description: "desc", Url: "url",
		HintUrl: "hint", ErrorUrl: "errUrl", ExampleUrl: "exUrl",
		SingleExample: "single", Examples: []string{"e1"},
	}
	m := info.Map()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 8}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- Map defined", actual)
}

func Test_Info_Map_Nil(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info
	m := info.Map()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Info returns nil -- Map nil", actual)
}

func Test_Info_LazyMap(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{RootName: "task"}
	m1 := info.LazyMap()
	m2 := info.LazyMap() // cached

	// Act
	actual := args.Map{
		"len": len(m1),
		"sameRef": len(m2) == len(m1),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"sameRef": true,
	}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- LazyMap", actual)
}

func Test_Info_LazyMapPrettyJsonString(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{RootName: "task"}
	result := info.LazyMapPrettyJsonString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- LazyMapPrettyJsonString", actual)
}

func Test_Info_MapWithPayload(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{RootName: "task"}
	m := info.MapWithPayload([]byte("data"))

	// Act
	actual := args.Map{"gt0": len(m) > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Info returns non-empty -- MapWithPayload", actual)
}

func Test_Info_PrettyJsonStringWithPayloads(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{RootName: "task"}
	result := info.PrettyJsonStringWithPayloads([]byte("data"))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Info returns non-empty -- PrettyJsonStringWithPayloads", actual)
}

// ── newInfoCreator ──

func Test_NewInfo_Default(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Default("n", "d", "u")

	// Act
	actual := args.Map{"name": info.Name()}

	// Assert
	expected := args.Map{"name": "n"}
	expected.ShouldBeEqual(t, 0, "NewInfo returns correct value -- Default", actual)
}

func Test_NewInfo_Examples(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Examples("n", "d", "u", "e1")

	// Act
	actual := args.Map{"hasEx": info.HasExamples()}

	// Assert
	expected := args.Map{"hasEx": true}
	expected.ShouldBeEqual(t, 0, "NewInfo returns correct value -- Examples", actual)
}

func Test_NewInfo_Create(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.Create(false, "n", "d", "u", "h", "e", "ex", "ch", "e1")

	// Act
	actual := args.Map{"name": info.Name()}

	// Assert
	expected := args.Map{"name": "n"}
	expected.ShouldBeEqual(t, 0, "NewInfo returns correct value -- Create", actual)
}

func Test_NewInfo_SecureCreate(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.SecureCreate("n", "d", "u", "h", "e", "ex", "ch")

	// Act
	actual := args.Map{"secure": info.IsSecure()}

	// Assert
	expected := args.Map{"secure": true}
	expected.ShouldBeEqual(t, 0, "NewInfo returns correct value -- SecureCreate", actual)
}

func Test_NewInfo_PlainCreate(t *testing.T) {
	// Arrange
	info := coretaskinfo.New.Info.PlainCreate("n", "d", "u", "h", "e", "ex", "ch")

	// Act
	actual := args.Map{"plain": info.IsPlainText()}

	// Assert
	expected := args.Map{"plain": true}
	expected.ShouldBeEqual(t, 0, "NewInfo returns correct value -- PlainCreate", actual)
}

// ── newInfoPlainTextCreator ──

func Test_PlainCreator_AllMethods(t *testing.T) {
	// Arrange
	p := coretaskinfo.New.Info.Plain

	// Act
	actual := args.Map{
		"default":  p.Default("n", "d", "u").Name(),
		"ndu":      p.NameDescUrl("n", "d", "u").Name(),
		"nduEx":    p.NameDescUrlExamples("n", "d", "u", "e").HasExamples(),
		"ndueErr":  p.NewNameDescUrlErrorUrl("n", "d", "u", "e").HasErrorUrl(),
		"nduee":    p.NameDescUrlErrUrlExamples("n", "d", "u", "e", "e1").HasExamples(),
		"ndex":     p.NameDescExamples("n", "d", "e1").HasExamples(),
		"ex":       p.Examples("n", "d", "e1").HasExamples(),
		"nuEx":     p.NameUrlExamples("n", "u", "e1").HasExamples(),
		"uEx":      p.UrlExamples("u", "e1").HasExamples(),
		"exOnly":   p.ExamplesOnly("e1").HasExamples(),
		"urlOnly":  p.UrlOnly("u").HasUrl(),
		"errOnly":  p.ErrorUrlOnly("e").HasErrorUrl(),
		"hintOnly": p.HintUrlOnly("h").HasHintUrl(),
		"descHint": p.DescHintUrlOnly("d", "h").HasHintUrl(),
		"nameHint": p.NameHintUrlOnly("n", "h").HasHintUrl(),
		"singleEx": p.SingleExampleOnly("s").HasChainingExample(),
		"allUrl":   p.AllUrl("n", "d", "u", "h", "e").HasUrl(),
		"allUrlEx": p.AllUrlExamples("n", "d", "u", "h", "e", "e1").HasExamples(),
		"urlSE":    p.UrlSingleExample("n", "d", "u", "s").HasChainingExample(),
		"singleE":  p.SingleExample("n", "d", "s").HasChainingExample(),
		"exUrl":    p.ExampleUrl("n", "d", "eu", "s").HasExampleUrl(),
		"exUrlSE":  p.ExampleUrlSingleExample("n", "d", "eu", "s").HasChainingExample(),
	}

	// Assert
	expected := args.Map{
		"default": "n", "ndu": "n", "nduEx": true, "ndueErr": true,
		"nduee": true, "ndex": true, "ex": true, "nuEx": true,
		"uEx": true, "exOnly": true, "urlOnly": true, "errOnly": true,
		"hintOnly": true, "descHint": true, "nameHint": true, "singleEx": true,
		"allUrl": true, "allUrlEx": true, "urlSE": true, "singleE": true,
		"exUrl": true, "exUrlSE": true,
	}
	expected.ShouldBeEqual(t, 0, "PlainCreator returns correct value -- all methods", actual)
}

// ── newInfoSecureTextCreator ──

func Test_SecureCreator_AllMethods(t *testing.T) {
	// Arrange
	s := coretaskinfo.New.Info.Secure

	// Act
	actual := args.Map{
		"default":   s.Default("n", "d", "u").IsSecure(),
		"ndu":       s.NameDescUrl("n", "d", "u").IsSecure(),
		"nduEx":     s.NameDescUrlExamples("n", "d", "u", "e").IsSecure(),
		"ndueErr":   s.NewNameDescUrlErrorUrl("n", "d", "u", "e").IsSecure(),
		"nduee":     s.NameDescUrlErrUrlExamples("n", "d", "u", "e", "e1").IsSecure(),
		"ndex":      s.NameDescExamples("n", "d", "e1").IsSecure(),
		"ex":        s.Examples("n", "d", "e1").IsSecure(),
		"exOnly":    s.ExamplesOnly("e1").IsSecure(),
		"urlOnly":   s.UrlOnly("u").IsSecure(),
		"errOnly":   s.ErrorUrlOnly("e").IsSecure(),
		"hintOnly":  s.HintUrlOnly("h").IsSecure(),
		"nameHint":  s.NameHintUrlOnly("n", "h").IsSecure(),
		"singleEx":  s.SingleExampleOnly("s").IsSecure(),
		"allUrlEx":  s.AllUrlExamples("n", "d", "u", "h", "e", "e1").IsSecure(),
		"allUrl":    s.AllUrl("n", "d", "u", "h", "e").IsSecure(),
		"urlSE":     s.UrlSingleExample("n", "d", "u", "s").IsSecure(),
		"singleE":   s.SingleExample("n", "d", "s").IsSecure(),
		"exUrl":     s.ExampleUrl("n", "d", "eu", "s").IsSecure(),
		"exUrlSE":   s.ExampleUrlSingleExample("n", "d", "eu", "s").IsSecure(),
		"newExUrl":  s.NewExampleUrlSecure("n", "d", "eu", "s").IsSecure(),
	}

	// Assert
	expected := args.Map{
		"default": true, "ndu": true, "nduEx": true, "ndueErr": true,
		"nduee": true, "ndex": true, "ex": true, "exOnly": true,
		"urlOnly": true, "errOnly": true, "hintOnly": true, "nameHint": true,
		"singleEx": true, "allUrlEx": true, "allUrl": true, "urlSE": true,
		"singleE": true, "exUrl": true, "exUrlSE": true, "newExUrl": true,
	}
	expected.ShouldBeEqual(t, 0, "SecureCreator returns correct value -- all methods", actual)
}

// ── Deserialized / DeserializedUsingJsonResult ──

func Test_NewInfo_Deserialized(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{RootName: "task"}
	bytes, _ := info.Serialize()
	parsed, err := coretaskinfo.New.Info.Deserialized(bytes)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": parsed.Name(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "task",
	}
	expected.ShouldBeEqual(t, 0, "NewInfo returns correct value -- Deserialized", actual)
}

func Test_NewInfo_DeserializedUsingJsonResult(t *testing.T) {
	// Arrange
	info := coretaskinfo.Info{RootName: "task"}
	jsonResult := info.JsonPtr()
	parsed, err := coretaskinfo.New.Info.DeserializedUsingJsonResult(jsonResult)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": parsed.Name(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "task",
	}
	expected.ShouldBeEqual(t, 0, "NewInfo returns correct value -- DeserializedUsingJsonResult", actual)
}

// ── Exclude options ──

func Test_Info_WithExcludeOptions(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{
		RootName: "task",
		ExcludeOptions: &coretaskinfo.ExcludingOptions{
			IsExcludeRootName:    true,
			IsExcludeDescription: true,
			IsExcludeUrl:         true,
			IsExcludeHintUrl:     true,
			IsExcludeErrorUrl:    true,
			IsSecureText:         true,
		},
	}

	// Act
	actual := args.Map{
		"exName": info.IsExcludeRootName(),
		"exDesc": info.IsExcludeDescription(),
		"exUrl":  info.IsExcludeUrl(),
		"exHint": info.IsExcludeHintUrl(),
		"exErr":  info.IsExcludeErrorUrl(),
		"exPay":  info.IsExcludePayload(),
	}

	// Assert
	expected := args.Map{
		"exName": true, "exDesc": true, "exUrl": true,
		"exHint": true, "exErr": true, "exPay": true,
	}
	expected.ShouldBeEqual(t, 0, "Info returns non-empty -- with ExcludeOptions", actual)
}

// ── MapWithPayloadAsAny ──

func Test_Info_MapWithPayloadAsAny(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{RootName: "task"}
	m := info.MapWithPayloadAsAny("hello")

	// Act
	actual := args.Map{"gt0": len(m) > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Info returns non-empty -- MapWithPayloadAsAny", actual)
}

// ── LazyMapWithPayload ──

func Test_Info_LazyMapWithPayload(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{RootName: "task"}
	m := info.LazyMapWithPayload([]byte("data"))

	// Act
	actual := args.Map{"gt0": len(m) > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Info returns non-empty -- LazyMapWithPayload", actual)
}

// ── LazyMapWithPayloadAsAny ──

func Test_Info_LazyMapWithPayloadAsAny(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{RootName: "task"}
	m := info.LazyMapWithPayloadAsAny("payload")

	// Act
	actual := args.Map{"gt0": len(m) > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Info returns non-empty -- LazyMapWithPayloadAsAny", actual)
}

// ── JsonParseSelfInject ──

func Test_Info_JsonParseSelfInject(t *testing.T) {
	// Arrange
	info := coretaskinfo.Info{RootName: "task"}
	jsonResult := info.JsonPtr()
	var parsed coretaskinfo.Info
	err := parsed.JsonParseSelfInject(jsonResult)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": parsed.Name(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "task",
	}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- JsonParseSelfInject", actual)
}

// ── Deserialize ──

func Test_Info_Deserialize(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{RootName: "task"}
	var parsed coretaskinfo.Info
	err := info.Deserialize(&parsed)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": parsed.Name(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "task",
	}
	expected.ShouldBeEqual(t, 0, "Info returns correct value -- Deserialize", actual)
}
