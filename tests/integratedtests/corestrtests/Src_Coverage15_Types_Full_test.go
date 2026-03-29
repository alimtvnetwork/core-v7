package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_SrcC15_SimpleSlice_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC15_SimpleSlice_Verification", func() {
		// Arrange
		tc := srcC15SimpleSliceTestCase

		// Act
		ssAdd := corestr.New.SimpleSlice.SpreadStrings("a"); ssAdd.Add("b")
		ssSplit := corestr.New.SimpleSlice.Cap(5); ssSplit.AddSplit("a,b,c", ",")
		ssIf := corestr.New.SimpleSlice.Cap(5); ssIf.AddIf(true, "a"); ssIf.AddIf(false, "b")
		ssAdds := corestr.New.SimpleSlice.Cap(5); ssAdds.Adds("a", "b")
		ssApp := corestr.New.SimpleSlice.Cap(5); ssApp.Append("a", "b")
		ssFmt := corestr.New.SimpleSlice.Cap(5); ssFmt.AppendFmt("hello %s", "world")
		ssFmtE := corestr.New.SimpleSlice.Cap(5); ssFmtE.AppendFmt("")
		actual := args.Map{
			"addLen":      ssAdd.Length(),
			"splitLen":    ssSplit.Length(),
			"addIfLen":    ssIf.Length(),
			"addsLen":     ssAdds.Length(),
			"appendLen":   ssApp.Length(),
			"fmtLen":      ssFmt.Length(),
			"fmtEmptyLen": ssFmtE.Length(),
			"isEmpty":     corestr.New.SimpleSlice.Cap(0).IsEmpty(),
			"length":      corestr.New.SimpleSlice.SpreadStrings("a").Length(),
			"lastIndex":   corestr.New.SimpleSlice.SpreadStrings("a", "b").LastIndex(),
			"hasIdx0":     corestr.New.SimpleSlice.SpreadStrings("a").HasIndex(0),
			"hasIdx5":     corestr.New.SimpleSlice.SpreadStrings("a").HasIndex(5),
			"listLen":     len(corestr.New.SimpleSlice.SpreadStrings("a").List()),
			"stringsLen":  len(corestr.New.SimpleSlice.SpreadStrings("a").Strings()),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC15_AnyToString_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC15_AnyToString_Verification", func() {
		// Arrange
		tc := srcC15AnyToStringTestCase

		// Act
		actual := args.Map{
			"nonEmpty": corestr.AnyToString(false, 42) != "",
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC15_AllIndLenSlices_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC15_AllIndLenSlices_Verification", func() {
		// Arrange
		tc := srcC15AllIndLenSlicesTestCase

		// Act
		ss1 := corestr.New.SimpleSlice.SpreadStrings("a", "b")
		ss2 := corestr.New.SimpleSlice.SpreadStrings("c")
		actual := args.Map{
			"length": corestr.AllIndividualsLengthOfSimpleSlices(ss1, ss2),
			"nilLen": corestr.AllIndividualsLengthOfSimpleSlices(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC15_Types_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC15_Types_Verification", func() {
		// Arrange
		tc := srcC15TypesTestCase

		// Act
		noPanic := !callPanicsSrcC15(func() {
			_ = corestr.NewValidValuesUsingValues(*corestr.NewValidValue("a"))
			_ = corestr.InvalidValueStatus("test")
			_ = corestr.TextWithLineNumber{Text: "hello", LineNumber: 1}
			_ = corestr.NewLeftRight("a", "b")
			_ = corestr.NewLeftMiddleRight("a", "b", "c")
			_ = corestr.KeyValuePair{Key: "k", Value: "v"}
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC15_KeyValueCollection_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC15_KeyValueCollection_Verification", func() {
		// Arrange
		tc := srcC15KeyValueCollectionTestCase

		// Act
		kvc := corestr.New.KeyValues.Cap(5)
		kvc.Add("k", "v")
		actual := args.Map{
			"addLen":  kvc.Length(),
			"isEmpty": corestr.New.KeyValues.Cap(0).IsEmpty(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC15_HashsetsCollection_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC15_HashsetsCollection_Verification", func() {
		// Arrange
		tc := srcC15HashsetsCollectionTestCase

		// Act
		hsc := corestr.New.HashsetsCollection.Cap(5)
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		actual := args.Map{
			"isEmpty": corestr.New.HashsetsCollection.Cap(0).IsEmpty(),
			"addLen":  hsc.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC15_DataModel_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC15_DataModel_Verification", func() {
		// Arrange
		tc := srcC15DataModelTestCase

		// Act
		dm := &corestr.HashmapDataModel{Items: map[string]string{"a": "1"}}
		hm := corestr.NewHashmapUsingDataModel(dm)
		hm2 := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "a", Value: "1"})
		dm2 := corestr.NewHashmapsDataModelUsing(hm2)
		actual := args.Map{
			"hmLen": hm.Length(),
			"dmLen": len(dm2.Items),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC15_SimpleStringOnce_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC15_SimpleStringOnce_Verification", func() {
		// Arrange
		tc := srcC15SimpleStringOnceTestCase

		// Act
		sso := corestr.New.SimpleStringOnce.Init("test")
		actual := args.Map{
			"nonEmpty": !sso.IsEmpty(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC15_CollectionCreators_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC15_CollectionCreators_Verification", func() {
		// Arrange
		tc := srcC15CollectionCreatorsTestCase

		// Act
		orig := []string{"a", "b"}
		cloned := corestr.New.Collection.CloneStrings(orig)
		orig[0] = "X"
		actual := args.Map{
			"lenCapLen":  corestr.New.Collection.LenCap(5, 10).Length(),
			"lineLen":    corestr.New.Collection.LineUsingSep(",", "a,b,c").Length(),
			"lineDefGe1": corestr.New.Collection.LineDefault("a | b").Length() >= 1,
			"strPlusLen": corestr.New.Collection.StringsPlusCap(5, []string{"a"}).Length(),
			"capStrLen":  corestr.New.Collection.CapStrings(5, []string{"a"}).Length(),
			"cloneDeep":  cloned.First(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC15_HashmapCreators_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC15_HashmapCreators_Verification", func() {
		// Arrange
		tc := srcC15HashmapCreatorsTestCase

		// Act
		actual := args.Map{
			"anyLen": corestr.New.Hashmap.KeyAnyValues(corestr.KeyAnyValuePair{Key: "a", Value: 1}).Length(),
			"colLen": corestr.New.Hashmap.KeyValuesCollection(
				corestr.New.Collection.Strings([]string{"a"}),
				corestr.New.Collection.Strings([]string{"1"}),
			).Length(),
			"strLen": corestr.New.Hashmap.KeyValuesStrings([]string{"a"}, []string{"1"}).Length(),
			"mapLen": corestr.New.Hashmap.MapWithCap(5, map[string]string{"a": "1"}).Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC15_HashsetCreators_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC15_HashsetCreators_Verification", func() {
		// Arrange
		tc := srcC15HashsetCreatorsTestCase

		// Act
		actual := args.Map{
			"optLen":  corestr.New.Hashset.StringsOption(10, false, "a", "b").Length(),
			"isEmpty": corestr.New.Hashset.Empty().IsEmpty(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC15_CocCreators_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC15_CocCreators_Verification", func() {
		// Arrange
		tc := srcC15CocCreatorsTestCase

		// Act
		actual := args.Map{
			"emptyIsEmpty": corestr.New.CollectionsOfCollection.Empty().IsEmpty(),
			"sosLen":       corestr.New.CollectionsOfCollection.StringsOfStrings(false, []string{"a"}, []string{"b"}).Length(),
			"spreadLen":    corestr.New.CollectionsOfCollection.SpreadStrings(false, "a", "b").Length(),
			"cloneLen":     corestr.New.CollectionsOfCollection.CloneStrings([]string{"a"}).Length(),
			"optionsLen":   corestr.New.CollectionsOfCollection.StringsOptions(false, 5, []string{"a"}).Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC15_CocJson_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC15_CocJson_Verification", func() {
		// Arrange
		tc := srcC15CocJsonTestCase

		// Act
		noPanic := !callPanicsSrcC15(func() {
			coc := corestr.New.CollectionsOfCollection.Strings([]string{"a"})
			b, _ := coc.MarshalJSON()
			coc2 := &corestr.CollectionsOfCollection{}
			_ = coc2.UnmarshalJSON(b)
			r := coc.Json()
			_ = r.Error == nil
			coc3 := &corestr.CollectionsOfCollection{}
			_, _ = coc3.ParseInjectUsingJson(&r)
			coc4 := &corestr.CollectionsOfCollection{}
			_ = coc4.JsonParseSelfInject(&r)
		})
		cocE := corestr.New.CollectionsOfCollection.Cap(5)
		cocE.Add(corestr.New.Collection.Cap(0))
		actual := args.Map{
			"noPanic":     noPanic,
			"addEmptyLen": cocE.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC15(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}
