package corestrtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// SimpleStringOnce — JSON/Serialization coverage gaps (15 methods)
// =============================================================================

func Test_Cov65_SSO_Json(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_Json", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		r := sso.Json()
		actual := args.Map{"noErr": r.Error == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "SSO Json no error", actual)
	})
}

func Test_Cov65_SSO_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_JsonPtr", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		r := sso.JsonPtr()
		actual := args.Map{"nonNil": r != nil, "noErr": r.Error == nil}
		expected := args.Map{"nonNil": true, "noErr": true}
		expected.ShouldBeEqual(t, 0, "SSO JsonPtr", actual)
	})
}

func Test_Cov65_SSO_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_JsonModel", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		m := sso.JsonModel()
		actual := args.Map{"val": m.Value, "init": m.IsInitialize}
		expected := args.Map{"val": "hello", "init": true}
		expected.ShouldBeEqual(t, 0, "SSO JsonModel", actual)
	})
}

func Test_Cov65_SSO_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_JsonModelAny", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		actual := args.Map{"nonNil": sso.JsonModelAny() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "SSO JsonModelAny", actual)
	})
}

func Test_Cov65_SSO_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_MarshalJSON", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		b, err := sso.MarshalJSON()
		actual := args.Map{"noErr": err == nil, "nonEmpty": len(b) > 0}
		expected := args.Map{"noErr": true, "nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SSO MarshalJSON", actual)
	})
}

func Test_Cov65_SSO_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_UnmarshalJSON", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		b, _ := sso.MarshalJSON()
		sso2 := corestr.New.SimpleStringOnce.Init("")
		err := sso2.UnmarshalJSON(b)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "SSO UnmarshalJSON", actual)
	})
}

func Test_Cov65_SSO_UnmarshalJSON_Error(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_UnmarshalJSON_Error", func() {
		sso := corestr.New.SimpleStringOnce.Init("")
		err := sso.UnmarshalJSON([]byte("invalid"))
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "SSO UnmarshalJSON error", actual)
	})
}

func Test_Cov65_SSO_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_ParseInjectUsingJson", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		jr := sso.JsonPtr()
		sso2 := corestr.New.SimpleStringOnce.Init("")
		r, err := sso2.ParseInjectUsingJson(jr)
		actual := args.Map{"noErr": err == nil, "nonNil": r != nil}
		expected := args.Map{"noErr": true, "nonNil": true}
		expected.ShouldBeEqual(t, 0, "SSO ParseInjectUsingJson", actual)
	})
}

func Test_Cov65_SSO_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_ParseInjectUsingJson_Error", func() {
		sso := corestr.New.SimpleStringOnce.Init("")
		jr := &corejson.Result{Error: errors.New("fail")}
		_, err := sso.ParseInjectUsingJson(jr)
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "SSO ParseInjectUsingJson error", actual)
	})
}

func Test_Cov65_SSO_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_ParseInjectUsingJsonMust", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		jr := sso.JsonPtr()
		sso2 := corestr.New.SimpleStringOnce.Init("")
		r := sso2.ParseInjectUsingJsonMust(jr)
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "SSO ParseInjectUsingJsonMust", actual)
	})
}

func Test_Cov65_SSO_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_ParseInjectUsingJsonMust_Panics", func() {
		sso := corestr.New.SimpleStringOnce.Init("")
		jr := &corejson.Result{Error: errors.New("fail")}
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			sso.ParseInjectUsingJsonMust(jr)
		}()
		actual := args.Map{"panicked": panicked}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "SSO ParseInjectUsingJsonMust panics", actual)
	})
}

func Test_Cov65_SSO_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_JsonParseSelfInject", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		jr := sso.JsonPtr()
		sso2 := corestr.New.SimpleStringOnce.Init("")
		err := sso2.JsonParseSelfInject(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "SSO JsonParseSelfInject", actual)
	})
}

func Test_Cov65_SSO_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_AsJsonContractsBinder", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		actual := args.Map{"nonNil": sso.AsJsonContractsBinder() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "SSO AsJsonContractsBinder", actual)
	})
}

func Test_Cov65_SSO_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_AsJsoner", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		actual := args.Map{"nonNil": sso.AsJsoner() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "SSO AsJsoner", actual)
	})
}

func Test_Cov65_SSO_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_AsJsonParseSelfInjector", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		actual := args.Map{"nonNil": sso.AsJsonParseSelfInjector() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "SSO AsJsonParseSelfInjector", actual)
	})
}

func Test_Cov65_SSO_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_AsJsonMarshaller", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		actual := args.Map{"nonNil": sso.AsJsonMarshaller() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "SSO AsJsonMarshaller", actual)
	})
}

func Test_Cov65_SSO_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_Serialize", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		b, err := sso.Serialize()
		actual := args.Map{"noErr": err == nil, "nonEmpty": len(b) > 0}
		expected := args.Map{"noErr": true, "nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SSO Serialize", actual)
	})
}

func Test_Cov65_SSO_Deserialize(t *testing.T) {
	safeTest(t, "Test_Cov65_SSO_Deserialize", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		var target corestr.SimpleStringOnceModel
		err := sso.Deserialize(&target)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "SSO Deserialize", actual)
	})
}

// =============================================================================
// LinkedCollectionNode — Next, isNextEqual, isNextChainEqual
// =============================================================================

func Test_Cov65_LCN_Next(t *testing.T) {
	safeTest(t, "Test_Cov65_LCN_Next", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		head := lc.Head()
		actual := args.Map{"hasNext": head.Next() != nil}
		expected := args.Map{"hasNext": true}
		expected.ShouldBeEqual(t, 0, "LCN Next", actual)
	})
}

func Test_Cov65_LCN_Next_Nil(t *testing.T) {
	safeTest(t, "Test_Cov65_LCN_Next_Nil", func() {
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		head := lc.Head()
		actual := args.Map{"noNext": head.Next() == nil}
		expected := args.Map{"noNext": true}
		expected.ShouldBeEqual(t, 0, "LCN Next nil tail", actual)
	})
}

func Test_Cov65_LCN_IsChainEqual_SameChain(t *testing.T) {
	safeTest(t, "Test_Cov65_LCN_IsChainEqual_SameChain", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc1.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2.Add(corestr.New.Collection.Strings([]string{"b"}))
		actual := args.Map{"eq": lc1.Head().IsChainEqual(lc2.Head())}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "LCN IsChainEqual same", actual)
	})
}

func Test_Cov65_LCN_IsChainEqual_Different(t *testing.T) {
	safeTest(t, "Test_Cov65_LCN_IsChainEqual_Different", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc1.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2.Add(corestr.New.Collection.Strings([]string{"c"}))
		actual := args.Map{"eq": lc1.Head().IsChainEqual(lc2.Head())}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "LCN IsChainEqual diff", actual)
	})
}

func Test_Cov65_LCN_IsEqual_EqualElements(t *testing.T) {
	safeTest(t, "Test_Cov65_LCN_IsEqual_EqualElements", func() {
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"eq": lc1.Head().IsEqual(lc2.Head())}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "LCN IsEqual same elements", actual)
	})
}

// =============================================================================
// KeyValueCollection — KeysHashset (panics with "implement me")
// =============================================================================

func Test_Cov65_KVC_KeysHashset_Panics(t *testing.T) {
	safeTest(t, "Test_Cov65_KVC_KeysHashset_Panics", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			kvc.KeysHashset()
		}()
		actual := args.Map{"panicked": panicked}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "KVC KeysHashset panics", actual)
	})
}

// =============================================================================
// CollectionsOfCollection — JsonPtr
// =============================================================================

func Test_Cov65_COC_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov65_COC_JsonPtr", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		r := coc.JsonPtr()
		actual := args.Map{"nonNil": r != nil, "noErr": r.Error == nil}
		expected := args.Map{"nonNil": true, "noErr": true}
		expected.ShouldBeEqual(t, 0, "COC JsonPtr", actual)
	})
}

// =============================================================================
// HashsetsCollection — JsonPtr
// =============================================================================

func Test_Cov65_HC_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov65_HC_JsonPtr", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hs := corestr.New.Hashset.Strings([]string{"x", "y"})
		hc.Add(hs)
		r := hc.JsonPtr()
		actual := args.Map{"nonNil": r != nil, "noErr": r.Error == nil}
		expected := args.Map{"nonNil": true, "noErr": true}
		expected.ShouldBeEqual(t, 0, "HC JsonPtr", actual)
	})
}
