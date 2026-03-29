package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── SimpleSlice (C15) ──

var srcC15SimpleSliceTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice Add AddSplit AddIf Adds Append AppendFmt IsEmpty Length LastIndex HasIndex List Strings return correct -- various",
	ExpectedInput: args.Map{
		"addLen":      2,
		"splitLen":    3,
		"addIfLen":    1,
		"addsLen":     2,
		"appendLen":   2,
		"fmtLen":      1,
		"fmtEmptyLen": 0,
		"isEmpty":     true,
		"length":      1,
		"lastIndex":   1,
		"hasIdx0":     true,
		"hasIdx5":     false,
		"listLen":     1,
		"stringsLen":  1,
	},
}

var srcC15AnyToStringTestCase = coretestcases.CaseV1{
	Title: "AnyToString returns non-empty -- int 42",
	ExpectedInput: args.Map{
		"nonEmpty": true,
	},
}

var srcC15AllIndLenSlicesTestCase = coretestcases.CaseV1{
	Title: "AllIndividualsLengthOfSimpleSlices returns correct -- 2 slices and nil",
	ExpectedInput: args.Map{
		"length": 3,
		"nilLen": 0,
	},
}

var srcC15TypesTestCase = coretestcases.CaseV1{
	Title: "ValidValues ValueStatus TextWithLineNumber LeftRight LeftMiddleRight KeyValuePair return correct -- struct fields",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC15KeyValueCollectionTestCase = coretestcases.CaseV1{
	Title: "KeyValueCollection Add IsEmpty return correct -- various",
	ExpectedInput: args.Map{
		"addLen":  1,
		"isEmpty": true,
	},
}

var srcC15HashsetsCollectionTestCase = coretestcases.CaseV1{
	Title: "HashsetsCollection IsEmpty Add return correct -- various",
	ExpectedInput: args.Map{
		"isEmpty": true,
		"addLen":  1,
	},
}

var srcC15DataModelTestCase = coretestcases.CaseV1{
	Title: "HashmapDataModel NewHashmapUsingDataModel NewHashmapsDataModelUsing return correct -- items",
	ExpectedInput: args.Map{
		"hmLen": 1,
		"dmLen": 1,
	},
}

var srcC15SimpleStringOnceTestCase = coretestcases.CaseV1{
	Title: "SimpleStringOnce returns non-empty -- test",
	ExpectedInput: args.Map{
		"nonEmpty": true,
	},
}

var srcC15CollectionCreatorsTestCase = coretestcases.CaseV1{
	Title: "Collection creators LenCap LineUsingSep LineDefault StringsPlusCap CapStrings CloneStrings return correct -- various",
	ExpectedInput: args.Map{
		"lenCapLen":    5,
		"lineLen":      3,
		"lineDefGe1":   true,
		"strPlusLen":   1,
		"capStrLen":    1,
		"cloneDeep":    "a",
	},
}

var srcC15HashmapCreatorsTestCase = coretestcases.CaseV1{
	Title: "Hashmap creators KeyAnyValues KeyValuesCollection KeyValuesStrings MapWithCap return correct -- various",
	ExpectedInput: args.Map{
		"anyLen":    1,
		"colLen":    1,
		"strLen":    1,
		"mapLen":    1,
	},
}

var srcC15HashsetCreatorsTestCase = coretestcases.CaseV1{
	Title: "Hashset creators StringsOption Empty return correct -- various",
	ExpectedInput: args.Map{
		"optLen":  2,
		"isEmpty": true,
	},
}

var srcC15CocCreatorsTestCase = coretestcases.CaseV1{
	Title: "CollectionsOfCollection creators Empty StringsOfStrings SpreadStrings CloneStrings StringsOptions return correct -- various",
	ExpectedInput: args.Map{
		"emptyIsEmpty":  true,
		"sosLen":        2,
		"spreadLen":     1,
		"cloneLen":      1,
		"optionsLen":    1,
	},
}

var srcC15CocJsonTestCase = coretestcases.CaseV1{
	Title: "CollectionsOfCollection MarshalJSON UnmarshalJSON Json ParseInjectUsingJson JsonParseSelfInject AddEmpty return correct -- various",
	ExpectedInput: args.Map{
		"noPanic":     true,
		"addEmptyLen": 0,
	},
}
