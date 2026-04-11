package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/smartystreets/goconvey/convey"
)

// =============================================================================
// HashsetsCollection
// =============================================================================

func Test_HashsetsCollection_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_IsEmpty", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		tc := coretestcases.CaseV1{
			Title:         "HashsetsCollection IsEmpty",
			ExpectedInput: args.Map{"IsEmpty": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsEmpty": hc.IsEmpty()})
	})
}

func Test_HashsetsCollection_HasItems(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_HasItems", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc.Add(hs)
		tc := coretestcases.CaseV1{
			Title:         "HashsetsCollection HasItems",
			ExpectedInput: args.Map{"HasItems": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"HasItems": hc.HasItems()})
	})
}

func Test_HashsetsCollection_Length(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_Length", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		tc := coretestcases.CaseV1{
			Title:         "HashsetsCollection Length",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": hc.Length()})
	})
}

func Test_HashsetsCollection_Add(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_Add", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc.Add(hs)
		tc := coretestcases.CaseV1{
			Title:         "HashsetsCollection Add",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": hc.Length()})
	})
}

func Test_HashsetsCollection_AddNonNil(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_AddNonNil", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.AddNonNil(nil)
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc.AddNonNil(hs)
		tc := coretestcases.CaseV1{
			Title:         "AddNonNil",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": hc.Length()})
	})
}

func Test_HashsetsCollection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_AddNonEmpty", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.AddNonEmpty(corestr.New.Hashset.Empty())
		hc.AddNonEmpty(corestr.New.Hashset.Strings([]string{"a"}))
		tc := coretestcases.CaseV1{
			Title:         "AddNonEmpty",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": hc.Length()})
	})
}

func Test_HashsetsCollection_Adds(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_Adds", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hs1 := corestr.New.Hashset.Strings([]string{"a"})
		hs2 := corestr.New.Hashset.Strings([]string{"b"})
		hc.Adds(hs1, hs2)
		tc := coretestcases.CaseV1{
			Title:         "Adds",
			ExpectedInput: args.Map{"Length": 2},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": hc.Length()})
	})
}

func Test_HashsetsCollection_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_Adds_Nil", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Adds(nil)
		tc := coretestcases.CaseV1{
			Title:         "Adds nil",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": hc.Length()})
	})
}

func Test_HashsetsCollection_StringsList(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_StringsList", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		result := hc.StringsList()
		tc := coretestcases.CaseV1{
			Title:         "StringsList",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(result)})
	})
}

func Test_HashsetsCollection_StringsList_Empty(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_StringsList_Empty", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		result := hc.StringsList()
		tc := coretestcases.CaseV1{
			Title:         "StringsList empty",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(result)})
	})
}

func Test_HashsetsCollection_HasAll_True(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_HasAll_True", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a", "b"}))
		result := hc.HasAll("a", "b")
		tc := coretestcases.CaseV1{
			Title:         "HasAll true",
			ExpectedInput: args.Map{"HasAll": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"HasAll": result})
	})
}

func Test_HashsetsCollection_HasAll_Empty(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_HasAll_Empty", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		result := hc.HasAll("a")
		tc := coretestcases.CaseV1{
			Title:         "HasAll empty",
			ExpectedInput: args.Map{"HasAll": false},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"HasAll": result})
	})
}

func Test_HashsetsCollection_ListDirectPtr(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_ListDirectPtr", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		result := hc.ListDirectPtr()
		convey.Convey("ListDirectPtr", t, func() {
			convey.So(result, convey.ShouldNotBeNil)
			convey.So(len(*result), convey.ShouldEqual, 1)
		})
	})
}

func Test_HashsetsCollection_AddHashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_AddHashsetsCollection", func() {
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc2 := corestr.New.HashsetsCollection.Empty()
		hc2.Add(corestr.New.Hashset.Strings([]string{"b"}))
		hc1.AddHashsetsCollection(hc2)
		tc := coretestcases.CaseV1{
			Title:         "AddHashsetsCollection",
			ExpectedInput: args.Map{"Length": 2},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": hc1.Length()})
	})
}

func Test_HashsetsCollection_AddHashsetsCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_AddHashsetsCollection_Nil", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.AddHashsetsCollection(nil)
		tc := coretestcases.CaseV1{
			Title:         "AddHashsetsCollection nil",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": hc.Length()})
	})
}

func Test_HashsetsCollection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_ConcatNew", func() {
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc2 := corestr.New.HashsetsCollection.Empty()
		hc2.Add(corestr.New.Hashset.Strings([]string{"b"}))
		result := hc1.ConcatNew(hc2)
		tc := coretestcases.CaseV1{
			Title:         "ConcatNew",
			ExpectedInput: args.Map{"Length": 2},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": result.Length()})
	})
}

func Test_HashsetsCollection_ConcatNew_NoArgs(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_ConcatNew_NoArgs", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		result := hc.ConcatNew()
		tc := coretestcases.CaseV1{
			Title:         "ConcatNew no args",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": result.Length()})
	})
}

func Test_HashsetsCollection_LastIndex(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_LastIndex", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		tc := coretestcases.CaseV1{
			Title:         "LastIndex",
			ExpectedInput: args.Map{"LastIndex": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"LastIndex": hc.LastIndex()})
	})
}

func Test_HashsetsCollection_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_IsEqual", func() {
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc2 := corestr.New.HashsetsCollection.Empty()
		hc2.Add(corestr.New.Hashset.Strings([]string{"a"}))
		tc := coretestcases.CaseV1{
			Title:         "IsEqual",
			ExpectedInput: args.Map{"IsEqual": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsEqual": hc1.IsEqualPtr(hc2)})
	})
}

func Test_HashsetsCollection_IsEqualPtr_DiffLength(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_IsEqualPtr_DiffLength", func() {
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc2 := corestr.New.HashsetsCollection.Empty()
		tc := coretestcases.CaseV1{
			Title:         "IsEqualPtr diff length",
			ExpectedInput: args.Map{"IsEqual": false},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsEqual": hc1.IsEqualPtr(hc2)})
	})
}

func Test_HashsetsCollection_IsEqualPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_IsEqualPtr_BothEmpty", func() {
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc2 := corestr.New.HashsetsCollection.Empty()
		tc := coretestcases.CaseV1{
			Title:         "IsEqualPtr both empty",
			ExpectedInput: args.Map{"IsEqual": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsEqual": hc1.IsEqualPtr(hc2)})
	})
}

func Test_HashsetsCollection_IsEqualPtr_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_IsEqualPtr_SamePtr", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		tc := coretestcases.CaseV1{
			Title:         "IsEqualPtr same ptr",
			ExpectedInput: args.Map{"IsEqual": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsEqual": hc.IsEqualPtr(hc)})
	})
}

func Test_HashsetsCollection_IsEqual_Val(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_IsEqual_Val", func() {
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc2 := corestr.New.HashsetsCollection.Empty()
		hc2.Add(corestr.New.Hashset.Strings([]string{"a"}))
		convey.Convey("IsEqual val", t, func() {
			convey.So(hc1.IsEqual(*hc2), convey.ShouldBeTrue)
		})
	})
}

func Test_HashsetsCollection_Json(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_Json", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		jsonResult := hc.JsonPtr()
		target := corestr.New.HashsetsCollection.Empty()
		result := target.ParseInjectUsingJsonMust(jsonResult)
		tc := coretestcases.CaseV1{
			Title:         "Json roundtrip",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": result.Length()})
	})
}

func Test_HashsetsCollection_String(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_String", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		convey.Convey("String empty", t, func() {
			convey.So(hc.String(), convey.ShouldContainSubstring, "No Element")
		})
	})
}

func Test_HashsetsCollection_String_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_String_NonEmpty", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		convey.Convey("String non-empty", t, func() {
			convey.So(hc.String(), convey.ShouldNotBeEmpty)
		})
	})
}

func Test_HashsetsCollection_Join(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_Join", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		convey.Convey("Join", t, func() {
			convey.So(hc.Join(","), convey.ShouldEqual, "a")
		})
	})
}

func Test_HashsetsCollection_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_Serialize", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		bytes, err := hc.Serialize()
		convey.Convey("Serialize", t, func() {
			convey.So(err, convey.ShouldBeNil)
			convey.So(bytes, convey.ShouldNotBeEmpty)
		})
	})
}

func Test_HashsetsCollection_Deserialize(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_Deserialize", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		var target corestr.HashsetsCollection
		err := hc.Deserialize(&target)
		convey.Convey("Deserialize", t, func() {
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func Test_HashsetsCollection_Interfaces(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_Interfaces", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		convey.Convey("Interfaces", t, func() {
			convey.So(hc.AsJsonContractsBinder(), convey.ShouldNotBeNil)
			convey.So(hc.AsJsoner(), convey.ShouldNotBeNil)
			convey.So(hc.AsJsonParseSelfInjector(), convey.ShouldNotBeNil)
			convey.So(hc.AsJsonMarshaller(), convey.ShouldNotBeNil)
		})
	})
}

func Test_HashsetsCollection_IndexOf(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_IndexOf", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		convey.Convey("IndexOf", t, func() {
			convey.So(hc.IndexOf(0), convey.ShouldNotBeNil)
		})
	})
}

func Test_HashsetsCollection_ListPtr(t *testing.T) {
	safeTest(t, "Test_Cov57_HashsetsCollection_ListPtr", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		convey.Convey("ListPtr", t, func() {
			convey.So(hc.ListPtr(), convey.ShouldNotBeNil)
			convey.So(hc.List(), convey.ShouldNotBeNil)
		})
	})
}

// =============================================================================
// KeyValuePair
// =============================================================================

func Test_KeyValuePair_Basics(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_Basics", func() {
		kv := corestr.KeyValuePair{Key: "name", Value: "test"}
		tc := coretestcases.CaseV1{
			Title:         "KVP basics",
			ExpectedInput: args.Map{
				"KeyName":      "name",
				"ValueString":  "test",
				"IsKey":        true,
				"IsVal":        true,
				"HasKey":       true,
				"HasValue":     true,
				"IsKeyEmpty":   false,
				"IsValueEmpty": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"KeyName":      kv.KeyName(),
			"ValueString":  kv.ValueString(),
			"IsKey":        kv.IsKey("name"),
			"IsVal":        kv.IsVal("test"),
			"HasKey":       kv.HasKey(),
			"HasValue":     kv.HasValue(),
			"IsKeyEmpty":   kv.IsKeyEmpty(),
			"IsValueEmpty": kv.IsValueEmpty(),
		})
	})
}

func Test_KeyValuePair_ValueConversions(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_ValueConversions", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "42"}
		tc := coretestcases.CaseV1{
			Title:         "KVP conversions",
			ExpectedInput: args.Map{
				"ValueInt":  42,
				"DefInt":    42,
				"DefByte":   byte(42),
				"Float":     42.0,
				"DefFloat":  42.0,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"ValueInt":  kv.ValueInt(0),
			"DefInt":    kv.ValueDefInt(),
			"DefByte":   kv.ValueDefByte(),
			"Float":     kv.ValueFloat64(0),
			"DefFloat":  kv.ValueDefFloat64(),
		})
	})
}

func Test_KeyValuePair_ValueBool(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_ValueBool", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "true"}
		tc := coretestcases.CaseV1{
			Title:         "KVP ValueBool",
			ExpectedInput: args.Map{"Bool": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Bool": kv.ValueBool()})
	})
}

func Test_KeyValuePair_ValueBool_Empty(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_ValueBool_Empty", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: ""}
		tc := coretestcases.CaseV1{
			Title:         "KVP ValueBool empty",
			ExpectedInput: args.Map{"Bool": false},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Bool": kv.ValueBool()})
	})
}

func Test_KeyValuePair_Is(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_Is", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := coretestcases.CaseV1{
			Title:         "KVP Is",
			ExpectedInput: args.Map{
				"Is": true,
				"IsNot": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Is": kv.Is("k", "v"),
			"IsNot": kv.Is("k", "x"),
		})
	})
}

func Test_KeyValuePair_Trim(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_Trim", func() {
		kv := corestr.KeyValuePair{Key: " k ", Value: " v "}
		tc := coretestcases.CaseV1{
			Title:         "KVP Trim",
			ExpectedInput: args.Map{
				"TrimKey": "k",
				"TrimVal": "v",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"TrimKey": kv.TrimKey(),
			"TrimVal": kv.TrimValue(),
		})
	})
}

func Test_KeyValuePair_ValueByte(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_ValueByte", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "100"}
		tc := coretestcases.CaseV1{
			Title:         "KVP ValueByte",
			ExpectedInput: args.Map{"Byte": byte(100)},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Byte": kv.ValueByte(0)})
	})
}

func Test_KeyValuePair_FormatString(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_FormatString", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		result := kv.FormatString("%s=%s")
		tc := coretestcases.CaseV1{
			Title:         "KVP FormatString",
			ExpectedInput: "k=v",
		}

		// Assert
		tc.ShouldBeEqual(t, 0, result)
	})
}

func Test_KeyValuePair_String(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_String", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		convey.Convey("KVP String", t, func() {
			convey.So(kv.String(), convey.ShouldNotBeEmpty)
		})
	})
}

func Test_KeyValuePair_Clear(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_Clear", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Clear()
		tc := coretestcases.CaseV1{
			Title:         "KVP Clear",
			ExpectedInput: args.Map{"IsEmpty": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsEmpty": kv.IsKeyValueEmpty()})
	})
}

func Test_KeyValuePair_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_Dispose", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Dispose()
		tc := coretestcases.CaseV1{
			Title:         "KVP Dispose",
			ExpectedInput: args.Map{"IsEmpty": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsEmpty": kv.IsKeyValueEmpty()})
	})
}

func Test_KeyValuePair_ValueValid(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_ValueValid", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kv.ValueValid()
		tc := coretestcases.CaseV1{
			Title:         "KVP ValueValid",
			ExpectedInput: args.Map{"IsValid": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsValid": vv.IsValid})
	})
}

func Test_KeyValuePair_ValueValidOptions(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_ValueValidOptions", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kv.ValueValidOptions(false, "err")
		tc := coretestcases.CaseV1{
			Title:         "KVP ValueValidOptions",
			ExpectedInput: args.Map{
				"IsValid": false,
				"Message": "err",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"IsValid": vv.IsValid,
			"Message": vv.Message,
		})
	})
}

func Test_KeyValuePair_IsKeyValueAnyEmpty(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_IsKeyValueAnyEmpty", func() {
		kv := corestr.KeyValuePair{Key: "", Value: "v"}
		tc := coretestcases.CaseV1{
			Title:         "IsKeyValueAnyEmpty",
			ExpectedInput: args.Map{"Result": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Result": kv.IsKeyValueAnyEmpty()})
	})
}

func Test_KeyValuePair_IsVariableNameEqual(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_IsVariableNameEqual", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := coretestcases.CaseV1{
			Title:         "IsVariableNameEqual",
			ExpectedInput: args.Map{"Result": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Result": kv.IsVariableNameEqual("k")})
	})
}

func Test_KeyValuePair_IsValueEqual(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_IsValueEqual", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := coretestcases.CaseV1{
			Title:         "IsValueEqual",
			ExpectedInput: args.Map{"Result": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Result": kv.IsValueEqual("v")})
	})
}

func Test_KeyValuePair_Json(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_Json", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		convey.Convey("KVP Json", t, func() {
			convey.So(kv.JsonPtr(), convey.ShouldNotBeNil)
		})
	})
}

func Test_KeyValuePair_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_Serialize", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		bytes, err := kv.Serialize()
		convey.Convey("KVP Serialize", t, func() {
			convey.So(err, convey.ShouldBeNil)
			convey.So(bytes, convey.ShouldNotBeEmpty)
		})
	})
}

func Test_KeyValuePair_SerializeMust(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_SerializeMust", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		convey.Convey("KVP SerializeMust", t, func() {
			convey.So(kv.SerializeMust(), convey.ShouldNotBeEmpty)
		})
	})
}

func Test_KeyValuePair_Compile(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_Compile", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		convey.Convey("KVP Compile", t, func() {
			convey.So(kv.Compile(), convey.ShouldEqual, kv.String())
		})
	})
}

func Test_KeyValuePair_VariableName(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValuePair_VariableName", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		tc := coretestcases.CaseV1{
			Title:         "VariableName",
			ExpectedInput: "k",
		}

		// Assert
		tc.ShouldBeEqual(t, 0, kv.VariableName())
	})
}

// =============================================================================
// KeyValueCollection
// =============================================================================

func Test_KeyValueCollection_Basics(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_Basics", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k1", "v1").Add("k2", "v2")
		tc := coretestcases.CaseV1{
			Title:         "KVC basics",
			ExpectedInput: args.Map{
				"Length":  2,
				"Count":  2,
				"HasAny": true,
				"HasKey": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Length":  kvc.Length(),
			"Count":  kvc.Count(),
			"HasAny": kvc.HasAnyItem(),
			"HasKey": kvc.HasKey("k1"),
		})
	})
}

func Test_KeyValueCollection_AddIf(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_AddIf", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddIf(true, "a", "b")
		kvc.AddIf(false, "c", "d")
		tc := coretestcases.CaseV1{
			Title:         "KVC AddIf",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": kvc.Length()})
	})
}

func Test_KeyValueCollection_FirstLast(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_FirstLast", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k1", "v1").Add("k2", "v2")
		tc := coretestcases.CaseV1{
			Title:         "KVC FirstLast",
			ExpectedInput: args.Map{
				"FirstKey": "k1",
				"LastKey":  "k2",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"FirstKey": kvc.First().Key,
			"LastKey":  kvc.Last().Key,
		})
	})
}

func Test_KeyValueCollection_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_FirstOrDefault_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		convey.Convey("FirstOrDefault empty", t, func() {
			convey.So(kvc.FirstOrDefault(), convey.ShouldBeNil)
			convey.So(kvc.LastOrDefault(), convey.ShouldBeNil)
		})
	})
}

func Test_KeyValueCollection_Find(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_Find", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")
		result := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, kv.Key == "a", false
		})
		tc := coretestcases.CaseV1{
			Title:         "KVC Find",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(result)})
	})
}

func Test_KeyValueCollection_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_SafeValueAt", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		tc := coretestcases.CaseV1{
			Title:         "SafeValueAt",
			ExpectedInput: args.Map{
				"Value": "v",
				"Empty": "",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Value": kvc.SafeValueAt(0),
			"Empty": kvc.SafeValueAt(99),
		})
	})
}

func Test_KeyValueCollection_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_SafeValuesAtIndexes", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		result := kvc.SafeValuesAtIndexes(0, 99)
		tc := coretestcases.CaseV1{
			Title:         "SafeValuesAtIndexes",
			ExpectedInput: args.Map{"Length": 2},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(result)})
	})
}

func Test_KeyValueCollection_StringsUsingFormat(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_StringsUsingFormat", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		result := kvc.StringsUsingFormat("%s=%s")
		tc := coretestcases.CaseV1{
			Title:         "StringsUsingFormat",
			ExpectedInput: "k=v",
		}

		// Assert
		tc.ShouldBeEqual(t, 0, result[0])
	})
}

func Test_KeyValueCollection_AllKeysValues(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_AllKeysValues", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k1", "v1").Add("k2", "v2")
		tc := coretestcases.CaseV1{
			Title:         "AllKeys/AllValues",
			ExpectedInput: args.Map{
				"KeysLen": 2,
				"ValsLen": 2,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"KeysLen": len(kvc.AllKeys()),
			"ValsLen": len(kvc.AllValues()),
		})
	})
}

func Test_KeyValueCollection_AllKeysSorted(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_AllKeysSorted", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("b", "2").Add("a", "1")
		result := kvc.AllKeysSorted()
		tc := coretestcases.CaseV1{
			Title:         "AllKeysSorted",
			ExpectedInput: args.Map{"First": "a"},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"First": result[0]})
	})
}

func Test_KeyValueCollection_JoinMethods(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_JoinMethods", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		convey.Convey("Join methods", t, func() {
			convey.So(kvc.Join(","), convey.ShouldNotBeEmpty)
			convey.So(kvc.JoinKeys(","), convey.ShouldEqual, "k")
			convey.So(kvc.JoinValues(","), convey.ShouldEqual, "v")
		})
	})
}

func Test_KeyValueCollection_AddMap(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_AddMap", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddMap(map[string]string{"k": "v"})
		tc := coretestcases.CaseV1{
			Title:         "AddMap",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": kvc.Length()})
	})
}

func Test_KeyValueCollection_AddMap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_AddMap_Nil", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddMap(nil)
		tc := coretestcases.CaseV1{
			Title:         "AddMap nil",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": kvc.Length()})
	})
}

func Test_KeyValueCollection_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_AddHashsetMap", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashsetMap(map[string]bool{"a": true})
		tc := coretestcases.CaseV1{
			Title:         "AddHashsetMap",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": kvc.Length()})
	})
}

func Test_KeyValueCollection_AddHashsetMap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_AddHashsetMap_Nil", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashsetMap(nil)
		tc := coretestcases.CaseV1{
			Title:         "AddHashsetMap nil",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": kvc.Length()})
	})
}

func Test_KeyValueCollection_AddHashset(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_AddHashset", func() {
		kvc := &corestr.KeyValueCollection{}
		hs := corestr.New.Hashset.Strings([]string{"a"})
		kvc.AddHashset(hs)
		tc := coretestcases.CaseV1{
			Title:         "AddHashset",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": kvc.Length()})
	})
}

func Test_KeyValueCollection_AddHashset_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_AddHashset_Nil", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashset(nil)
		tc := coretestcases.CaseV1{
			Title:         "AddHashset nil",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": kvc.Length()})
	})
}

func Test_KeyValueCollection_AddsHashmap(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_AddsHashmap", func() {
		kvc := &corestr.KeyValueCollection{}
		hm := corestr.New.Hashmap.Cap(1)
		hm.AddOrUpdate("k", "v")
		kvc.AddsHashmap(hm)
		tc := coretestcases.CaseV1{
			Title:         "AddsHashmap",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": kvc.Length()})
	})
}

func Test_KeyValueCollection_AddsHashmap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_AddsHashmap_Nil", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddsHashmap(nil)
		tc := coretestcases.CaseV1{
			Title:         "AddsHashmap nil",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": kvc.Length()})
	})
}

func Test_KeyValueCollection_AddsHashmaps(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_AddsHashmaps", func() {
		kvc := &corestr.KeyValueCollection{}
		hm := corestr.New.Hashmap.Cap(1)
		hm.AddOrUpdate("k", "v")
		kvc.AddsHashmaps(hm)
		tc := coretestcases.CaseV1{
			Title:         "AddsHashmaps",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": kvc.Length()})
	})
}

func Test_KeyValueCollection_AddsHashmaps_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_AddsHashmaps_Nil", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddsHashmaps()
		tc := coretestcases.CaseV1{
			Title:         "AddsHashmaps nil",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": kvc.Length()})
	})
}

func Test_KeyValueCollection_Hashmap(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_Hashmap", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		hm := kvc.Hashmap()
		tc := coretestcases.CaseV1{
			Title:         "Hashmap",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": hm.Length()})
	})
}

func Test_KeyValueCollection_Map(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_Map", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		m := kvc.Map()
		tc := coretestcases.CaseV1{
			Title:         "Map",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(m)})
	})
}

func Test_KeyValueCollection_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_IsContains", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		tc := coretestcases.CaseV1{
			Title:         "IsContains",
			ExpectedInput: args.Map{
				"Contains": true,
				"Missing": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Contains": kvc.IsContains("k"),
			"Missing": kvc.IsContains("x"),
		})
	})
}

func Test_KeyValueCollection_Get(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_Get", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		val, ok := kvc.Get("k")
		tc := coretestcases.CaseV1{
			Title:         "Get",
			ExpectedInput: args.Map{
				"Value": "v",
				"Ok": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Value": val,
			"Ok": ok,
		})
	})
}

func Test_KeyValueCollection_Get_Missing(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_Get_Missing", func() {
		kvc := &corestr.KeyValueCollection{}
		_, ok := kvc.Get("k")
		tc := coretestcases.CaseV1{
			Title:         "Get missing",
			ExpectedInput: args.Map{"Ok": false},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Ok": ok})
	})
}

func Test_KeyValueCollection_HasIndex(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_HasIndex", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		tc := coretestcases.CaseV1{
			Title:         "HasIndex",
			ExpectedInput: args.Map{
				"Has": true,
				"Missing": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Has": kvc.HasIndex(0),
			"Missing": kvc.HasIndex(99),
		})
	})
}

func Test_KeyValueCollection_Adds(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_Adds", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Adds(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)
		tc := coretestcases.CaseV1{
			Title:         "KVC Adds",
			ExpectedInput: args.Map{"Length": 2},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": kvc.Length()})
	})
}

func Test_KeyValueCollection_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_Adds_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Adds()
		tc := coretestcases.CaseV1{
			Title:         "KVC Adds empty",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": kvc.Length()})
	})
}

func Test_KeyValueCollection_AddStringBySplit(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_AddStringBySplit", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddStringBySplit("=", "key=value")
		tc := coretestcases.CaseV1{
			Title:         "AddStringBySplit",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": kvc.Length()})
	})
}

func Test_KeyValueCollection_AddStringBySplitTrim(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_AddStringBySplitTrim", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddStringBySplitTrim("=", " key = value ")
		tc := coretestcases.CaseV1{
			Title:         "AddStringBySplitTrim",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": kvc.Length()})
	})
}

func Test_KeyValueCollection_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_Serialize", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		bytes, err := kvc.Serialize()
		convey.Convey("KVC Serialize", t, func() {
			convey.So(err, convey.ShouldBeNil)
			convey.So(bytes, convey.ShouldNotBeEmpty)
		})
	})
}

func Test_KeyValueCollection_SerializeMust(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_SerializeMust", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		convey.Convey("KVC SerializeMust", t, func() {
			convey.So(kvc.SerializeMust(), convey.ShouldNotBeEmpty)
		})
	})
}

func Test_KeyValueCollection_String(t *testing.T) {
	safeTest(t, "Test_Cov57_KeyValueCollection_String", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		convey.Convey("KVC String", t, func() {
			convey.So(kvc.String(), convey.ShouldNotBeEmpty)
			convey.So(kvc.Compile(), convey.ShouldEqual, kvc.String())
		})
	})
}

// =============================================================================
// LeftRight
// =============================================================================

func Test_LeftRight_New(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_New", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := coretestcases.CaseV1{
			Title:         "LR New",
			ExpectedInput: args.Map{
				"Left": "a", "Right": "b", "IsValid": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lr.Left, "Right": lr.Right, "IsValid": lr.IsValid,
		})
	})
}

func Test_LeftRight_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_Invalid", func() {
		lr := corestr.InvalidLeftRight("err")
		tc := coretestcases.CaseV1{
			Title:         "LR Invalid",
			ExpectedInput: args.Map{"IsValid": false},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsValid": lr.IsValid})
	})
}

func Test_LeftRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_InvalidNoMessage", func() {
		lr := corestr.InvalidLeftRightNoMessage()
		tc := coretestcases.CaseV1{
			Title:         "LR InvalidNoMessage",
			ExpectedInput: args.Map{"IsValid": false},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsValid": lr.IsValid})
	})
}

func Test_LeftRight_Bytes(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_Bytes", func() {
		lr := corestr.NewLeftRight("a", "b")
		convey.Convey("LR Bytes", t, func() {
			convey.So(lr.LeftBytes(), convey.ShouldResemble, []byte("a"))
			convey.So(lr.RightBytes(), convey.ShouldResemble, []byte("b"))
		})
	})
}

func Test_LeftRight_Trim(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_Trim", func() {
		lr := corestr.NewLeftRight(" a ", " b ")
		tc := coretestcases.CaseV1{
			Title:         "LR Trim",
			ExpectedInput: args.Map{
				"Left": "a",
				"Right": "b",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lr.LeftTrim(),
			"Right": lr.RightTrim(),
		})
	})
}

func Test_LeftRight_EmptyChecks(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_EmptyChecks", func() {
		lr := corestr.NewLeftRight("", "b")
		tc := coretestcases.CaseV1{
			Title:         "LR empty checks",
			ExpectedInput: args.Map{
				"LeftEmpty":  true,
				"RightEmpty": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"LeftEmpty":  lr.IsLeftEmpty(),
			"RightEmpty": lr.IsRightEmpty(),
		})
	})
}

func Test_LeftRight_WhitespaceChecks(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_WhitespaceChecks", func() {
		lr := corestr.NewLeftRight("  ", "b")
		tc := coretestcases.CaseV1{
			Title:         "LR whitespace",
			ExpectedInput: args.Map{
				"LeftWS":  true,
				"RightWS": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"LeftWS":  lr.IsLeftWhitespace(),
			"RightWS": lr.IsRightWhitespace(),
		})
	})
}

func Test_LeftRight_ValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_ValidNonEmpty", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := coretestcases.CaseV1{
			Title:         "LR valid non-empty",
			ExpectedInput: args.Map{
				"Left": true, "Right": true, "Safe": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lr.HasValidNonEmptyLeft(), "Right": lr.HasValidNonEmptyRight(), "Safe": lr.HasSafeNonEmpty(),
		})
	})
}

func Test_LeftRight_ValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_ValidNonWhitespace", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := coretestcases.CaseV1{
			Title:         "LR valid non-ws",
			ExpectedInput: args.Map{
				"Left": true,
				"Right": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lr.HasValidNonWhitespaceLeft(), "Right": lr.HasValidNonWhitespaceRight(),
		})
	})
}

func Test_LeftRight_Is(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_Is", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := coretestcases.CaseV1{
			Title:         "LR Is",
			ExpectedInput: args.Map{
				"Is": true, "IsLeft": true, "IsRight": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Is": lr.Is("a", "b"), "IsLeft": lr.IsLeft("a"), "IsRight": lr.IsRight("b"),
		})
	})
}

func Test_LeftRight_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_IsEqual", func() {
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		tc := coretestcases.CaseV1{
			Title:         "LR IsEqual",
			ExpectedInput: args.Map{"Equal": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Equal": lr1.IsEqual(lr2)})
	})
}

func Test_LeftRight_Clone(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_Clone", func() {
		lr := corestr.NewLeftRight("a", "b")
		cloned := lr.Clone()
		tc := coretestcases.CaseV1{
			Title:         "LR Clone",
			ExpectedInput: args.Map{"Left": "a"},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Left": cloned.Left})
	})
}

func Test_LeftRight_NonPtr(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_NonPtr", func() {
		lr := corestr.NewLeftRight("a", "b")
		convey.Convey("NonPtr", t, func() {
			convey.So(lr.NonPtr().Left, convey.ShouldEqual, "a")
			convey.So(lr.Ptr(), convey.ShouldNotBeNil)
		})
	})
}

func Test_LeftRight_RegexMatch(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_RegexMatch", func() {
		lr := corestr.NewLeftRight("abc123", "xyz")
		re := regexp.MustCompile(`\d+`)
		tc := coretestcases.CaseV1{
			Title:         "LR Regex",
			ExpectedInput: args.Map{
				"LeftMatch": true,
				"RightMatch": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"LeftMatch": lr.IsLeftRegexMatch(re), "RightMatch": lr.IsRightRegexMatch(re),
		})
	})
}

func Test_LeftRight_RegexMatch_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_RegexMatch_Nil", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := coretestcases.CaseV1{
			Title:         "LR Regex nil",
			ExpectedInput: args.Map{
				"Left": false,
				"Right": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lr.IsLeftRegexMatch(nil), "Right": lr.IsRightRegexMatch(nil),
		})
	})
}

func Test_LeftRight_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_ClearDispose", func() {
		lr := corestr.NewLeftRight("a", "b")
		lr.Clear()
		tc := coretestcases.CaseV1{
			Title:         "LR Clear",
			ExpectedInput: args.Map{"LeftEmpty": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"LeftEmpty": lr.IsLeftEmpty()})
	})
}

func Test_LeftRight_FromSplit(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_FromSplit", func() {
		lr := corestr.LeftRightFromSplit("key=value", "=")
		tc := coretestcases.CaseV1{
			Title:         "LR FromSplit",
			ExpectedInput: args.Map{
				"Left": "key",
				"Right": "value",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lr.Left,
			"Right": lr.Right,
		})
	})
}

func Test_LeftRight_FromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_FromSplitTrimmed", func() {
		lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")
		tc := coretestcases.CaseV1{
			Title:         "LR FromSplitTrimmed",
			ExpectedInput: args.Map{
				"Left": "key",
				"Right": "value",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lr.Left,
			"Right": lr.Right,
		})
	})
}

func Test_LeftRight_FromSplitFull(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_FromSplitFull", func() {
		lr := corestr.LeftRightFromSplitFull("a:b:c", ":")
		tc := coretestcases.CaseV1{
			Title:         "LR FromSplitFull",
			ExpectedInput: args.Map{
				"Left": "a",
				"Right": "b:c",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lr.Left,
			"Right": lr.Right,
		})
	})
}

func Test_LeftRight_FromSplitFullTrimmed(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_FromSplitFullTrimmed", func() {
		lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")
		tc := coretestcases.CaseV1{
			Title:         "LR FromSplitFullTrimmed",
			ExpectedInput: args.Map{"Left": "a"},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Left": lr.Left})
	})
}

func Test_LeftRight_UsingSlice(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_UsingSlice", func() {
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})
		tc := coretestcases.CaseV1{
			Title:         "LR UsingSlice",
			ExpectedInput: args.Map{
				"Left": "a",
				"Right": "b",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lr.Left,
			"Right": lr.Right,
		})
	})
}

func Test_LeftRight_UsingSlice_Single(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_UsingSlice_Single", func() {
		lr := corestr.LeftRightUsingSlice([]string{"a"})
		tc := coretestcases.CaseV1{
			Title:         "LR UsingSlice single",
			ExpectedInput: args.Map{
				"Left": "a",
				"IsValid": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lr.Left,
			"IsValid": lr.IsValid,
		})
	})
}

func Test_LeftRight_UsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_UsingSlice_Empty", func() {
		lr := corestr.LeftRightUsingSlice([]string{})
		tc := coretestcases.CaseV1{
			Title:         "LR UsingSlice empty",
			ExpectedInput: args.Map{"IsValid": false},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsValid": lr.IsValid})
	})
}

func Test_LeftRight_UsingSlicePtr(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_UsingSlicePtr", func() {
		lr := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
		tc := coretestcases.CaseV1{
			Title:         "LR UsingSlicePtr",
			ExpectedInput: args.Map{"Left": "a"},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Left": lr.Left})
	})
}

func Test_LeftRight_TrimmedUsingSlice(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_TrimmedUsingSlice", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
		tc := coretestcases.CaseV1{
			Title:         "LR TrimmedUsingSlice",
			ExpectedInput: args.Map{
				"Left": "a",
				"Right": "b",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lr.Left,
			"Right": lr.Right,
		})
	})
}

func Test_LeftRight_TrimmedUsingSlice_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_TrimmedUsingSlice_Nil", func() {
		lr := corestr.LeftRightTrimmedUsingSlice(nil)
		tc := coretestcases.CaseV1{
			Title:         "LR TrimmedUsingSlice nil",
			ExpectedInput: args.Map{"IsValid": false},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsValid": lr.IsValid})
	})
}

func Test_LeftRight_TrimmedUsingSlice_Single(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftRight_TrimmedUsingSlice_Single", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a "})
		tc := coretestcases.CaseV1{
			Title:         "LR TrimmedUsingSlice single",
			ExpectedInput: args.Map{
				"Left": "a",
				"IsValid": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lr.Left,
			"IsValid": lr.IsValid,
		})
	})
}

// =============================================================================
// LeftMiddleRight
// =============================================================================

func Test_LeftMiddleRight_New(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_New", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		tc := coretestcases.CaseV1{
			Title:         "LMR New",
			ExpectedInput: args.Map{
				"Left": "a",
				"Middle": "b",
				"Right": "c",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lmr.Left,
			"Middle": lmr.Middle,
			"Right": lmr.Right,
		})
	})
}

func Test_LeftMiddleRight_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_Invalid", func() {
		lmr := corestr.InvalidLeftMiddleRight("err")
		tc := coretestcases.CaseV1{
			Title:         "LMR Invalid",
			ExpectedInput: args.Map{"IsValid": false},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsValid": lmr.IsValid})
	})
}

func Test_LeftMiddleRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_InvalidNoMessage", func() {
		lmr := corestr.InvalidLeftMiddleRightNoMessage()
		tc := coretestcases.CaseV1{
			Title:         "LMR InvalidNoMessage",
			ExpectedInput: args.Map{"IsValid": false},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsValid": lmr.IsValid})
	})
}

func Test_LeftMiddleRight_Bytes(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_Bytes", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		convey.Convey("LMR Bytes", t, func() {
			convey.So(lmr.LeftBytes(), convey.ShouldResemble, []byte("a"))
			convey.So(lmr.MiddleBytes(), convey.ShouldResemble, []byte("b"))
			convey.So(lmr.RightBytes(), convey.ShouldResemble, []byte("c"))
		})
	})
}

func Test_LeftMiddleRight_Trim(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_Trim", func() {
		lmr := corestr.NewLeftMiddleRight(" a ", " b ", " c ")
		tc := coretestcases.CaseV1{
			Title:         "LMR Trim",
			ExpectedInput: args.Map{
				"Left": "a",
				"Mid": "b",
				"Right": "c",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lmr.LeftTrim(),
			"Mid": lmr.MiddleTrim(),
			"Right": lmr.RightTrim(),
		})
	})
}

func Test_LeftMiddleRight_EmptyChecks(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_EmptyChecks", func() {
		lmr := corestr.NewLeftMiddleRight("", "b", "")
		tc := coretestcases.CaseV1{
			Title:         "LMR empty",
			ExpectedInput: args.Map{
				"LeftEmpty": true, "MidEmpty": false, "RightEmpty": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"LeftEmpty": lmr.IsLeftEmpty(), "MidEmpty": lmr.IsMiddleEmpty(), "RightEmpty": lmr.IsRightEmpty(),
		})
	})
}

func Test_LeftMiddleRight_WhitespaceChecks(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_WhitespaceChecks", func() {
		lmr := corestr.NewLeftMiddleRight("  ", "b", "  ")
		tc := coretestcases.CaseV1{
			Title:         "LMR ws",
			ExpectedInput: args.Map{
				"LeftWS": true, "MidWS": false, "RightWS": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"LeftWS": lmr.IsLeftWhitespace(), "MidWS": lmr.IsMiddleWhitespace(), "RightWS": lmr.IsRightWhitespace(),
		})
	})
}

func Test_LeftMiddleRight_ValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_ValidNonEmpty", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		tc := coretestcases.CaseV1{
			Title:         "LMR valid non-empty",
			ExpectedInput: args.Map{
				"Left": true, "Mid": true, "Right": true, "Safe": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lmr.HasValidNonEmptyLeft(), "Mid": lmr.HasValidNonEmptyMiddle(),
			"Right": lmr.HasValidNonEmptyRight(), "Safe": lmr.HasSafeNonEmpty(),
		})
	})
}

func Test_LeftMiddleRight_ValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_ValidNonWhitespace", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		tc := coretestcases.CaseV1{
			Title:         "LMR valid non-ws",
			ExpectedInput: args.Map{
				"Left": true,
				"Mid": true,
				"Right": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lmr.HasValidNonWhitespaceLeft(), "Mid": lmr.HasValidNonWhitespaceMiddle(),
			"Right": lmr.HasValidNonWhitespaceRight(),
		})
	})
}

func Test_LeftMiddleRight_IsAll(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_IsAll", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		tc := coretestcases.CaseV1{
			Title:         "LMR IsAll",
			ExpectedInput: args.Map{
				"IsAll": true,
				"Is": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"IsAll": lmr.IsAll("a", "b", "c"),
			"Is": lmr.Is("a", "c"),
		})
	})
}

func Test_LeftMiddleRight_Clone(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_Clone", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		cloned := lmr.Clone()
		tc := coretestcases.CaseV1{
			Title:         "LMR Clone",
			ExpectedInput: args.Map{"Left": "a"},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Left": cloned.Left})
	})
}

func Test_LeftMiddleRight_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_ToLeftRight", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()
		tc := coretestcases.CaseV1{
			Title:         "LMR ToLeftRight",
			ExpectedInput: args.Map{
				"Left": "a",
				"Right": "c",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lr.Left,
			"Right": lr.Right,
		})
	})
}

func Test_LeftMiddleRight_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_ClearDispose", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Dispose()
		tc := coretestcases.CaseV1{
			Title:         "LMR Dispose",
			ExpectedInput: args.Map{"LeftEmpty": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"LeftEmpty": lmr.IsLeftEmpty()})
	})
}

func Test_LeftMiddleRight_FromSplit(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_FromSplit", func() {
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
		tc := coretestcases.CaseV1{
			Title:         "LMR FromSplit",
			ExpectedInput: args.Map{
				"Left": "a",
				"Mid": "b",
				"Right": "c",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lmr.Left,
			"Mid": lmr.Middle,
			"Right": lmr.Right,
		})
	})
}

func Test_LeftMiddleRight_FromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_FromSplitTrimmed", func() {
		lmr := corestr.LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")
		tc := coretestcases.CaseV1{
			Title:         "LMR FromSplitTrimmed",
			ExpectedInput: args.Map{
				"Left": "a",
				"Mid": "b",
				"Right": "c",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lmr.Left,
			"Mid": lmr.Middle,
			"Right": lmr.Right,
		})
	})
}

func Test_LeftMiddleRight_FromSplitN(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_FromSplitN", func() {
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d", ":")
		tc := coretestcases.CaseV1{
			Title:         "LMR FromSplitN",
			ExpectedInput: args.Map{
				"Left": "a",
				"Mid": "b",
				"Right": "c:d",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lmr.Left,
			"Mid": lmr.Middle,
			"Right": lmr.Right,
		})
	})
}

func Test_LeftMiddleRight_FromSplitNTrimmed(t *testing.T) {
	safeTest(t, "Test_Cov57_LeftMiddleRight_FromSplitNTrimmed", func() {
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d ", ":")
		tc := coretestcases.CaseV1{
			Title:         "LMR FromSplitNTrimmed",
			ExpectedInput: args.Map{
				"Left": "a",
				"Mid": "b",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Left": lmr.Left,
			"Mid": lmr.Middle,
		})
	})
}

// =============================================================================
// ValidValue
// =============================================================================

func Test_ValidValue_New(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_New", func() {
		vv := corestr.NewValidValue("hello")
		tc := coretestcases.CaseV1{
			Title:         "VV New",
			ExpectedInput: args.Map{
				"Value": "hello",
				"IsValid": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Value": vv.Value,
			"IsValid": vv.IsValid,
		})
	})
}

func Test_ValidValue_Empty(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_Empty", func() {
		vv := corestr.NewValidValueEmpty()
		tc := coretestcases.CaseV1{
			Title:         "VV Empty",
			ExpectedInput: args.Map{
				"IsEmpty": true,
				"IsValid": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"IsEmpty": vv.IsEmpty(),
			"IsValid": vv.IsValid,
		})
	})
}

func Test_ValidValue_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_Invalid", func() {
		vv := corestr.InvalidValidValue("err")
		tc := coretestcases.CaseV1{
			Title:         "VV Invalid",
			ExpectedInput: args.Map{
				"IsValid": false,
				"Message": "err",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"IsValid": vv.IsValid,
			"Message": vv.Message,
		})
	})
}

func Test_ValidValue_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_InvalidNoMessage", func() {
		vv := corestr.InvalidValidValueNoMessage()
		tc := coretestcases.CaseV1{
			Title:         "VV InvalidNoMessage",
			ExpectedInput: args.Map{"IsValid": false},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsValid": vv.IsValid})
	})
}

func Test_ValidValue_Conversions(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_Conversions", func() {
		vv := corestr.NewValidValue("42")
		tc := coretestcases.CaseV1{
			Title:         "VV Conversions",
			ExpectedInput: args.Map{
				"Int": 42, "DefInt": 42, "Byte": byte(42), "Float": 42.0,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Int": vv.ValueInt(0), "DefInt": vv.ValueDefInt(),
			"Byte": vv.ValueByte(0), "Float": vv.ValueFloat64(0),
		})
	})
}

func Test_ValidValue_Bool(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_Bool", func() {
		vv := corestr.NewValidValue("true")
		tc := coretestcases.CaseV1{
			Title:         "VV Bool",
			ExpectedInput: args.Map{"Bool": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Bool": vv.ValueBool()})
	})
}

func Test_ValidValue_Bool_Empty(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_Bool_Empty", func() {
		vv := corestr.NewValidValue("")
		tc := coretestcases.CaseV1{
			Title:         "VV Bool empty",
			ExpectedInput: args.Map{"Bool": false},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Bool": vv.ValueBool()})
	})
}

func Test_ValidValue_WhitespaceChecks(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_WhitespaceChecks", func() {
		vv := corestr.NewValidValue("  ")
		tc := coretestcases.CaseV1{
			Title:         "VV ws",
			ExpectedInput: args.Map{
				"IsWS": true, "HasValidNonWS": false, "Trim": "",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"IsWS": vv.IsWhitespace(), "HasValidNonWS": vv.HasValidNonWhitespace(), "Trim": vv.Trim(),
		})
	})
}

func Test_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_HasValidNonEmpty", func() {
		vv := corestr.NewValidValue("a")
		tc := coretestcases.CaseV1{
			Title:         "VV HasValidNonEmpty",
			ExpectedInput: args.Map{
				"Result": true,
				"Safe": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Result": vv.HasValidNonEmpty(),
			"Safe": vv.HasSafeNonEmpty(),
		})
	})
}

func Test_ValidValue_Is(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_Is", func() {
		vv := corestr.NewValidValue("a")
		tc := coretestcases.CaseV1{
			Title:         "VV Is",
			ExpectedInput: args.Map{
				"Is": true,
				"IsNot": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Is": vv.Is("a"),
			"IsNot": vv.Is("b"),
		})
	})
}

func Test_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_IsAnyOf", func() {
		vv := corestr.NewValidValue("b")
		tc := coretestcases.CaseV1{
			Title:         "VV IsAnyOf",
			ExpectedInput: args.Map{
				"Found": true,
				"NotFound": false,
				"Empty": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Found": vv.IsAnyOf("a", "b"), "NotFound": vv.IsAnyOf("c"), "Empty": vv.IsAnyOf(),
		})
	})
}

func Test_ValidValue_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_IsContains", func() {
		vv := corestr.NewValidValue("hello world")
		tc := coretestcases.CaseV1{
			Title:         "VV IsContains",
			ExpectedInput: args.Map{"Contains": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Contains": vv.IsContains("world")})
	})
}

func Test_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_IsAnyContains", func() {
		vv := corestr.NewValidValue("hello world")
		tc := coretestcases.CaseV1{
			Title:         "VV IsAnyContains",
			ExpectedInput: args.Map{
				"Found": true, "NotFound": false, "Empty": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Found": vv.IsAnyContains("world"), "NotFound": vv.IsAnyContains("xyz"), "Empty": vv.IsAnyContains(),
		})
	})
}

func Test_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_IsEqualNonSensitive", func() {
		vv := corestr.NewValidValue("Hello")
		tc := coretestcases.CaseV1{
			Title:         "VV EqualFold",
			ExpectedInput: args.Map{"Result": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Result": vv.IsEqualNonSensitive("hello")})
	})
}

func Test_ValidValue_Regex(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_Regex", func() {
		vv := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)
		tc := coretestcases.CaseV1{
			Title:         "VV Regex",
			ExpectedInput: args.Map{
				"Matches": true, "FindStr": "123",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Matches": vv.IsRegexMatches(re), "FindStr": vv.RegexFindString(re),
		})
	})
}

func Test_ValidValue_Regex_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_Regex_Nil", func() {
		vv := corestr.NewValidValue("abc")
		tc := coretestcases.CaseV1{
			Title:         "VV Regex nil",
			ExpectedInput: args.Map{
				"Matches": false,
				"FindStr": "",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Matches": vv.IsRegexMatches(nil), "FindStr": vv.RegexFindString(nil),
		})
	})
}

func Test_ValidValue_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_RegexFindAllStrings", func() {
		vv := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		result := vv.RegexFindAllStrings(re, -1)
		tc := coretestcases.CaseV1{
			Title:         "VV FindAllStrings",
			ExpectedInput: args.Map{"Length": 3},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(result)})
	})
}

func Test_ValidValue_RegexFindAllStrings_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_RegexFindAllStrings_Nil", func() {
		vv := corestr.NewValidValue("abc")
		result := vv.RegexFindAllStrings(nil, -1)
		tc := coretestcases.CaseV1{
			Title:         "VV FindAllStrings nil",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(result)})
	})
}

func Test_ValidValue_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_RegexFindAllStringsWithFlag", func() {
		vv := corestr.NewValidValue("a1b2")
		re := regexp.MustCompile(`\d`)
		items, hasAny := vv.RegexFindAllStringsWithFlag(re, -1)
		tc := coretestcases.CaseV1{
			Title:         "VV FindAllWithFlag",
			ExpectedInput: args.Map{
				"Length": 2,
				"HasAny": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Length": len(items),
			"HasAny": hasAny,
		})
	})
}

func Test_ValidValue_RegexFindAllStringsWithFlag_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_RegexFindAllStringsWithFlag_Nil", func() {
		vv := corestr.NewValidValue("abc")
		_, hasAny := vv.RegexFindAllStringsWithFlag(nil, -1)
		tc := coretestcases.CaseV1{
			Title:         "VV FindAllWithFlag nil",
			ExpectedInput: args.Map{"HasAny": false},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"HasAny": hasAny})
	})
}

func Test_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_Split", func() {
		vv := corestr.NewValidValue("a,b,c")
		result := vv.Split(",")
		tc := coretestcases.CaseV1{
			Title:         "VV Split",
			ExpectedInput: args.Map{"Length": 3},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(result)})
	})
}

func Test_ValidValue_Clone(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_Clone", func() {
		vv := corestr.NewValidValue("a")
		cloned := vv.Clone()
		tc := coretestcases.CaseV1{
			Title:         "VV Clone",
			ExpectedInput: args.Map{"Value": "a"},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Value": cloned.Value})
	})
}

func Test_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_String", func() {
		vv := corestr.NewValidValue("a")
		tc := coretestcases.CaseV1{
			Title:         "VV String",
			ExpectedInput: "a",
		}

		// Assert
		tc.ShouldBeEqual(t, 0, vv.String())
	})
}

func Test_ValidValue_FullString(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_FullString", func() {
		vv := corestr.NewValidValue("a")
		convey.Convey("VV FullString", t, func() {
			convey.So(vv.FullString(), convey.ShouldNotBeEmpty)
		})
	})
}

func Test_ValidValue_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_ClearDispose", func() {
		vv := corestr.NewValidValue("a")
		vv.Dispose()
		tc := coretestcases.CaseV1{
			Title:         "VV Dispose",
			ExpectedInput: args.Map{
				"IsEmpty": true,
				"IsValid": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"IsEmpty": vv.IsEmpty(),
			"IsValid": vv.IsValid,
		})
	})
}

func Test_ValidValue_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_ValueBytesOnce", func() {
		vv := corestr.NewValidValue("abc")
		bytes1 := vv.ValueBytesOnce()
		bytes2 := vv.ValueBytesOnce()
		convey.Convey("VV ValueBytesOnce", t, func() {
			convey.So(bytes1, convey.ShouldResemble, []byte("abc"))
			convey.So(bytes2, convey.ShouldResemble, bytes1)
		})
	})
}

func Test_ValidValue_ValueBytesOncePtr(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_ValueBytesOncePtr", func() {
		vv := corestr.NewValidValue("abc")
		convey.Convey("VV ValueBytesOncePtr", t, func() {
			convey.So(vv.ValueBytesOncePtr(), convey.ShouldResemble, []byte("abc"))
		})
	})
}

func Test_ValidValue_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_Serialize", func() {
		vv := corestr.NewValidValue("a")
		bytes, err := vv.Serialize()
		convey.Convey("VV Serialize", t, func() {
			convey.So(err, convey.ShouldBeNil)
			convey.So(bytes, convey.ShouldNotBeEmpty)
		})
	})
}

func Test_ValidValue_NewUsingAny(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_NewUsingAny", func() {
		vv := corestr.NewValidValueUsingAny(false, true, "hello")
		tc := coretestcases.CaseV1{
			Title:         "VV UsingAny",
			ExpectedInput: args.Map{"IsValid": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsValid": vv.IsValid})
	})
}

func Test_ValidValue_DefByte_Overflow(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_DefByte_Overflow", func() {
		vv := corestr.NewValidValue("999")
		convey.Convey("VV DefByte overflow", t, func() {
			convey.So(vv.ValueDefByte(), convey.ShouldBeGreaterThan, 0)
		})
	})
}

func Test_ValidValue_DefFloat(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValue_DefFloat", func() {
		vv := corestr.NewValidValue("3.14")
		tc := coretestcases.CaseV1{
			Title:         "VV DefFloat",
			ExpectedInput: args.Map{"Float": 3.14},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Float": vv.ValueDefFloat64()})
	})
}

// =============================================================================
// ValidValues
// =============================================================================

func Test_ValidValues_Basics(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_Basics", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b")
		tc := coretestcases.CaseV1{
			Title:         "VVs basics",
			ExpectedInput: args.Map{
				"Length": 2, "Count": 2, "HasAny": true, "IsEmpty": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Length": vvs.Length(), "Count": vvs.Count(), "HasAny": vvs.HasAnyItem(), "IsEmpty": vvs.IsEmpty(),
		})
	})
}

func Test_ValidValues_AddFull(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_AddFull", func() {
		vvs := corestr.EmptyValidValues()
		vvs.AddFull(false, "val", "msg")
		tc := coretestcases.CaseV1{
			Title:         "VVs AddFull",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": vvs.Length()})
	})
}

func Test_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_SafeValueAt", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		tc := coretestcases.CaseV1{
			Title:         "VVs SafeValueAt",
			ExpectedInput: args.Map{
				"Value": "a",
				"Empty": "",
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Value": vvs.SafeValueAt(0),
			"Empty": vvs.SafeValueAt(99),
		})
	})
}

func Test_ValidValues_SafeValidValueAt(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_SafeValidValueAt", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		tc := coretestcases.CaseV1{
			Title:         "VVs SafeValidValueAt",
			ExpectedInput: args.Map{"Value": "a"},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Value": vvs.SafeValidValueAt(0)})
	})
}

func Test_ValidValues_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_SafeValuesAtIndexes", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		result := vvs.SafeValuesAtIndexes(0, 99)
		tc := coretestcases.CaseV1{
			Title:         "VVs SafeValuesAtIndexes",
			ExpectedInput: args.Map{"Length": 2},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(result)})
	})
}

func Test_ValidValues_SafeValidValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_SafeValidValuesAtIndexes", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		result := vvs.SafeValidValuesAtIndexes(0)
		tc := coretestcases.CaseV1{
			Title:         "VVs SafeValidValuesAtIndexes",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(result)})
	})
}

func Test_ValidValues_Strings(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_Strings", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		tc := coretestcases.CaseV1{
			Title:         "VVs Strings",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(vvs.Strings())})
	})
}

func Test_ValidValues_FullStrings(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_FullStrings", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		convey.Convey("VVs FullStrings", t, func() {
			convey.So(len(vvs.FullStrings()), convey.ShouldEqual, 1)
		})
	})
}

func Test_ValidValues_String(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_String", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		convey.Convey("VVs String", t, func() {
			convey.So(vvs.String(), convey.ShouldNotBeEmpty)
		})
	})
}

func Test_ValidValues_Find(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_Find", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b")
		result := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, vv.Value == "a", false
		})
		tc := coretestcases.CaseV1{
			Title:         "VVs Find",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(result)})
	})
}

func Test_ValidValues_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_ConcatNew", func() {
		vvs1 := corestr.EmptyValidValues()
		vvs1.Add("a")
		vvs2 := corestr.EmptyValidValues()
		vvs2.Add("b")
		result := vvs1.ConcatNew(false, vvs2)
		tc := coretestcases.CaseV1{
			Title:         "VVs ConcatNew",
			ExpectedInput: args.Map{"Length": 2},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": result.Length()})
	})
}

func Test_ValidValues_ConcatNew_CloneOnEmpty(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_ConcatNew_CloneOnEmpty", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		result := vvs.ConcatNew(true)
		tc := coretestcases.CaseV1{
			Title:         "VVs ConcatNew cloneOnEmpty",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": result.Length()})
	})
}

func Test_ValidValues_ConcatNew_NoClone(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_ConcatNew_NoClone", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		result := vvs.ConcatNew(false)
		convey.Convey("VVs ConcatNew noClone returns self", t, func() {
			convey.So(result, convey.ShouldEqual, vvs)
		})
	})
}

func Test_ValidValues_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_AddHashsetMap", func() {
		vvs := corestr.EmptyValidValues()
		vvs.AddHashsetMap(map[string]bool{"a": true})
		tc := coretestcases.CaseV1{
			Title:         "VVs AddHashsetMap",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": vvs.Length()})
	})
}

func Test_ValidValues_AddHashsetMap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_AddHashsetMap_Nil", func() {
		vvs := corestr.EmptyValidValues()
		vvs.AddHashsetMap(nil)
		tc := coretestcases.CaseV1{
			Title:         "VVs AddHashsetMap nil",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": vvs.Length()})
	})
}

func Test_ValidValues_AddHashset(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_AddHashset", func() {
		vvs := corestr.EmptyValidValues()
		hs := corestr.New.Hashset.Strings([]string{"a"})
		vvs.AddHashset(hs)
		tc := coretestcases.CaseV1{
			Title:         "VVs AddHashset",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": vvs.Length()})
	})
}

func Test_ValidValues_AddHashset_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_AddHashset_Nil", func() {
		vvs := corestr.EmptyValidValues()
		vvs.AddHashset(nil)
		tc := coretestcases.CaseV1{
			Title:         "VVs AddHashset nil",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": vvs.Length()})
	})
}

func Test_ValidValues_Hashmap(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_Hashmap", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		hm := vvs.Hashmap()
		tc := coretestcases.CaseV1{
			Title:         "VVs Hashmap",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": hm.Length()})
	})
}

func Test_ValidValues_Map(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_Map", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		m := vvs.Map()
		tc := coretestcases.CaseV1{
			Title:         "VVs Map",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(m)})
	})
}

func Test_ValidValues_NewUsingValues(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_NewUsingValues", func() {
		vvs := corestr.NewValidValuesUsingValues(
			corestr.ValidValue{Value: "a", IsValid: true},
		)
		tc := coretestcases.CaseV1{
			Title:         "VVs NewUsingValues",
			ExpectedInput: args.Map{"Length": 1},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": vvs.Length()})
	})
}

func Test_ValidValues_NewUsingValues_Empty(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_NewUsingValues_Empty", func() {
		vvs := corestr.NewValidValuesUsingValues()
		tc := coretestcases.CaseV1{
			Title:         "VVs NewUsingValues empty",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": vvs.Length()})
	})
}

func Test_ValidValues_HasIndex(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_HasIndex", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		tc := coretestcases.CaseV1{
			Title:         "VVs HasIndex",
			ExpectedInput: args.Map{
				"Has": true,
				"Missing": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Has": vvs.HasIndex(0),
			"Missing": vvs.HasIndex(99),
		})
	})
}

func Test_ValidValues_AddValidValues(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_AddValidValues", func() {
		vvs1 := corestr.EmptyValidValues()
		vvs1.Add("a")
		vvs2 := corestr.EmptyValidValues()
		vvs2.Add("b")
		vvs1.AddValidValues(vvs2)
		tc := coretestcases.CaseV1{
			Title:         "VVs AddValidValues",
			ExpectedInput: args.Map{"Length": 2},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": vvs1.Length()})
	})
}

func Test_ValidValues_AddValidValues_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_ValidValues_AddValidValues_Nil", func() {
		vvs := corestr.EmptyValidValues()
		vvs.AddValidValues(nil)
		tc := coretestcases.CaseV1{
			Title:         "VVs AddValidValues nil",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": vvs.Length()})
	})
}

// =============================================================================
// ValueStatus
// =============================================================================

func Test_ValueStatus_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov57_ValueStatus_Invalid", func() {
		vs := corestr.InvalidValueStatus("err")
		convey.Convey("VS Invalid", t, func() {
			convey.So(vs.ValueValid.IsValid, convey.ShouldBeFalse)
		})
	})
}

func Test_ValueStatus_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov57_ValueStatus_InvalidNoMessage", func() {
		vs := corestr.InvalidValueStatusNoMessage()
		convey.Convey("VS InvalidNoMessage", t, func() {
			convey.So(vs.ValueValid.IsValid, convey.ShouldBeFalse)
		})
	})
}

func Test_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_Cov57_ValueStatus_Clone", func() {
		vs := corestr.InvalidValueStatus("err")
		cloned := vs.Clone()
		convey.Convey("VS Clone", t, func() {
			convey.So(cloned.Index, convey.ShouldEqual, vs.Index)
		})
	})
}

// =============================================================================
// TextWithLineNumber
// =============================================================================

func Test_TextWithLineNumber(t *testing.T) {
	safeTest(t, "Test_Cov57_TextWithLineNumber", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
		tc := coretestcases.CaseV1{
			Title:         "TWL",
			ExpectedInput: args.Map{
				"HasLine": true, "IsInvalid": false, "Length": 5, "IsEmpty": false, "IsEmptyText": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"HasLine": twl.HasLineNumber(), "IsInvalid": twl.IsInvalidLineNumber(),
			"Length": twl.Length(), "IsEmpty": twl.IsEmpty(), "IsEmptyText": twl.IsEmptyText(),
		})
	})
}

func Test_TextWithLineNumber_Empty(t *testing.T) {
	safeTest(t, "Test_Cov57_TextWithLineNumber_Empty", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		tc := coretestcases.CaseV1{
			Title:         "TWL empty",
			ExpectedInput: args.Map{
				"IsEmpty": true, "IsEmptyBoth": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"IsEmpty": twl.IsEmpty(), "IsEmptyBoth": twl.IsEmptyTextLineBoth(),
		})
	})
}

// =============================================================================
// NonChainedLinkedListNodes
// =============================================================================

func Test_NonChainedLinkedListNodes(t *testing.T) {
	safeTest(t, "Test_Cov57_NonChainedLinkedListNodes", func() {
		nodes := corestr.NewNonChainedLinkedListNodes(2)
		tc := coretestcases.CaseV1{
			Title:         "NCLLN empty",
			ExpectedInput: args.Map{
				"IsEmpty": true, "Length": 0, "IsChainingApplied": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"IsEmpty": nodes.IsEmpty(), "Length": nodes.Length(), "IsChainingApplied": nodes.IsChainingApplied(),
		})
	})
}

func Test_NonChainedLinkedListNodes_Adds(t *testing.T) {
	safeTest(t, "Test_Cov57_NonChainedLinkedListNodes_Adds", func() {
		nodes := corestr.NewNonChainedLinkedListNodes(2)
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		nodes.Adds(ll.Head())
		tc := coretestcases.CaseV1{
			Title:         "NCLLN Adds",
			ExpectedInput: args.Map{
				"Length": 1,
				"HasItems": true,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"Length": nodes.Length(),
			"HasItems": nodes.HasItems(),
		})
	})
}

func Test_NonChainedLinkedListNodes_FirstLast(t *testing.T) {
	safeTest(t, "Test_Cov57_NonChainedLinkedListNodes_FirstLast", func() {
		nodes := corestr.NewNonChainedLinkedListNodes(2)
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		nodes.Adds(ll.Head(), ll.Head().Next())
		convey.Convey("NCLLN FirstLast", t, func() {
			convey.So(nodes.First(), convey.ShouldNotBeNil)
			convey.So(nodes.Last(), convey.ShouldNotBeNil)
			convey.So(nodes.FirstOrDefault(), convey.ShouldNotBeNil)
			convey.So(nodes.LastOrDefault(), convey.ShouldNotBeNil)
		})
	})
}

func Test_NonChainedLinkedListNodes_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov57_NonChainedLinkedListNodes_FirstOrDefault_Empty", func() {
		nodes := corestr.NewNonChainedLinkedListNodes(2)
		convey.Convey("NCLLN FirstOrDefault empty", t, func() {
			convey.So(nodes.FirstOrDefault(), convey.ShouldBeNil)
			convey.So(nodes.LastOrDefault(), convey.ShouldBeNil)
		})
	})
}

func Test_NonChainedLinkedListNodes_ApplyChaining(t *testing.T) {
	safeTest(t, "Test_Cov57_NonChainedLinkedListNodes_ApplyChaining", func() {
		nodes := corestr.NewNonChainedLinkedListNodes(2)
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		nodes.Adds(ll.Head(), ll.Head().Next())
		nodes.ApplyChaining()
		tc := coretestcases.CaseV1{
			Title:         "NCLLN ApplyChaining",
			ExpectedInput: args.Map{"IsChainingApplied": true},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"IsChainingApplied": nodes.IsChainingApplied()})
	})
}

func Test_NonChainedLinkedListNodes_Items(t *testing.T) {
	safeTest(t, "Test_Cov57_NonChainedLinkedListNodes_Items", func() {
		nodes := corestr.NewNonChainedLinkedListNodes(2)
		convey.Convey("NCLLN Items", t, func() {
			convey.So(nodes.Items(), convey.ShouldNotBeNil)
		})
	})
}

// =============================================================================
// NonChainedLinkedCollectionNodes
// =============================================================================

func Test_NonChainedLinkedCollectionNodes(t *testing.T) {
	safeTest(t, "Test_Cov57_NonChainedLinkedCollectionNodes", func() {
		nodes := corestr.NewNonChainedLinkedCollectionNodes(2)
		tc := coretestcases.CaseV1{
			Title:         "NCLCN empty",
			ExpectedInput: args.Map{
				"IsEmpty": true, "Length": 0, "IsChainingApplied": false,
			},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{
			"IsEmpty": nodes.IsEmpty(), "Length": nodes.Length(), "IsChainingApplied": nodes.IsChainingApplied(),
		})
	})
}

func Test_NonChainedLinkedCollectionNodes_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov57_NonChainedLinkedCollectionNodes_FirstOrDefault_Empty", func() {
		nodes := corestr.NewNonChainedLinkedCollectionNodes(2)
		convey.Convey("NCLCN FirstOrDefault empty", t, func() {
			convey.So(nodes.FirstOrDefault(), convey.ShouldBeNil)
			convey.So(nodes.LastOrDefault(), convey.ShouldBeNil)
		})
	})
}

func Test_NonChainedLinkedCollectionNodes_Items(t *testing.T) {
	safeTest(t, "Test_Cov57_NonChainedLinkedCollectionNodes_Items", func() {
		nodes := corestr.NewNonChainedLinkedCollectionNodes(2)
		convey.Convey("NCLCN Items", t, func() {
			convey.So(nodes.Items(), convey.ShouldNotBeNil)
		})
	})
}

// =============================================================================
// CloneSlice / CloneSliceIf
// =============================================================================

func Test_CloneSlice(t *testing.T) {
	safeTest(t, "Test_Cov57_CloneSlice", func() {
		result := corestr.CloneSlice([]string{"a", "b"})
		tc := coretestcases.CaseV1{
			Title:         "CloneSlice",
			ExpectedInput: args.Map{"Length": 2},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(result)})
	})
}

func Test_CloneSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov57_CloneSlice_Empty", func() {
		result := corestr.CloneSlice(nil)
		tc := coretestcases.CaseV1{
			Title:         "CloneSlice empty",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(result)})
	})
}

func Test_CloneSliceIf_True(t *testing.T) {
	safeTest(t, "Test_Cov57_CloneSliceIf_True", func() {
		result := corestr.CloneSliceIf(true, "a", "b")
		tc := coretestcases.CaseV1{
			Title:         "CloneSliceIf true",
			ExpectedInput: args.Map{"Length": 2},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(result)})
	})
}

func Test_CloneSliceIf_False(t *testing.T) {
	safeTest(t, "Test_Cov57_CloneSliceIf_False", func() {
		result := corestr.CloneSliceIf(false, "a", "b")
		tc := coretestcases.CaseV1{
			Title:         "CloneSliceIf false",
			ExpectedInput: args.Map{"Length": 2},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(result)})
	})
}

func Test_CloneSliceIf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov57_CloneSliceIf_Empty", func() {
		result := corestr.CloneSliceIf(true)
		tc := coretestcases.CaseV1{
			Title:         "CloneSliceIf empty",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": len(result)})
	})
}

// =============================================================================
// AnyToString
// =============================================================================

func Test_AnyToString(t *testing.T) {
	safeTest(t, "Test_Cov57_AnyToString", func() {
		result := corestr.AnyToString(false, "hello")
		convey.Convey("AnyToString", t, func() {
			convey.So(result, convey.ShouldNotBeEmpty)
		})
	})
}

func Test_AnyToString_WithFieldName(t *testing.T) {
	safeTest(t, "Test_Cov57_AnyToString_WithFieldName", func() {
		result := corestr.AnyToString(true, "hello")
		convey.Convey("AnyToString with field name", t, func() {
			convey.So(result, convey.ShouldNotBeEmpty)
		})
	})
}

func Test_AnyToString_Empty(t *testing.T) {
	safeTest(t, "Test_Cov57_AnyToString_Empty", func() {
		result := corestr.AnyToString(false, "")
		tc := coretestcases.CaseV1{
			Title:         "AnyToString empty",
			ExpectedInput: "",
		}

		// Assert
		tc.ShouldBeEqual(t, 0, result)
	})
}

// =============================================================================
// AllIndividualStringsOfStringsLength
// =============================================================================

func Test_AllIndividualStringsOfStringsLength(t *testing.T) {
	safeTest(t, "Test_Cov57_AllIndividualStringsOfStringsLength", func() {
		items := [][]string{{"a", "b"}, {"c"}}
		result := corestr.AllIndividualStringsOfStringsLength(&items)
		tc := coretestcases.CaseV1{
			Title:         "AllIndividualStringsOfStringsLength",
			ExpectedInput: args.Map{"Length": 3},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": result})
	})
}

func Test_AllIndividualStringsOfStringsLength_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_AllIndividualStringsOfStringsLength_Nil", func() {
		result := corestr.AllIndividualStringsOfStringsLength(nil)
		tc := coretestcases.CaseV1{
			Title:         "AllIndividualStringsOfStringsLength nil",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": result})
	})
}

// =============================================================================
// AllIndividualsLengthOfSimpleSlices
// =============================================================================

func Test_AllIndividualsLengthOfSimpleSlices(t *testing.T) {
	safeTest(t, "Test_Cov57_AllIndividualsLengthOfSimpleSlices", func() {
		ss1 := corestr.New.SimpleSlice.Strings([]string{"a"})
		ss2 := corestr.New.SimpleSlice.Strings([]string{"b", "c"})
		result := corestr.AllIndividualsLengthOfSimpleSlices(ss1, ss2)
		tc := coretestcases.CaseV1{
			Title:         "AllIndividualsLengthOfSimpleSlices",
			ExpectedInput: args.Map{"Length": 3},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": result})
	})
}

func Test_AllIndividualsLengthOfSimpleSlices_Nil(t *testing.T) {
	safeTest(t, "Test_Cov57_AllIndividualsLengthOfSimpleSlices_Nil", func() {
		result := corestr.AllIndividualsLengthOfSimpleSlices()
		tc := coretestcases.CaseV1{
			Title:         "AllIndividualsLengthOfSimpleSlices nil",
			ExpectedInput: args.Map{"Length": 0},
		}

		// Assert
		tc.ShouldBeEqualMap(t, 0, args.Map{"Length": result})
	})
}
