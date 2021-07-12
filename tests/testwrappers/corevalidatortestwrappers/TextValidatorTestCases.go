package corevalidatortestwrappers

import (
	"gitlab.com/evatix-go/core/corevalidator"
	"gitlab.com/evatix-go/core/enums/stringcompareas"
)

var TextValidatorsTestCases = []TextValidatorsWrapper{
	{
		Header: "Comparing all flag to false, and comparing equal.",
		ComparingLines: []string{
			"alim      alim 2 alim 4",
		},
		Validators: corevalidator.TextValidators{
			Items: []corevalidator.TextValidator{
				{
					Search:   "   alim      alim 2 alim 3                 ",
					SearchAs: stringcompareas.Equal,
				},
			},
		},
		IsSkipOnContentsEmpty: false,
		IsCaseSensitive:       true,
		ExpectationLines: []string{
			"0 )\tExpectationLines failed: Failed match method [\"Equal\"], Index : [0]",
			"     Actual-Processed: `\"alim      alim 2 alim 4\"`",
			"   Expected-Processed: `\"   alim      alim 2 alim 3                 \"`",
		},
	},
	{
		Header: "Comparing all flag to false, and comparing equal.",
		ComparingLines: []string{
			"   alim      alim 2 alim 3                 ",
		},
		Validators: corevalidator.TextValidators{
			Items: []corevalidator.TextValidator{
				{
					Search:   "   alim      alim 2 alim 3                 ",
					SearchAs: stringcompareas.Equal,
				},
			},
		},
		IsSkipOnContentsEmpty: false,
		IsCaseSensitive:       true,
		ExpectationLines:      []string{},
	},
	{
		Header: "Trim compare spaced same text should not give an error.",
		ComparingLines: []string{
			"alim      alim 2 alim 3",
		},
		Validators: corevalidator.TextValidators{
			Items: []corevalidator.TextValidator{
				{
					Search: "   alim      alim 2 alim 3                 ",
					ValidatorCoreCondition: corevalidator.ValidatorCoreCondition{
						IsTrimCompare:        true,
						IsUniqueWordOnly:     false,
						IsNonEmptyWhitespace: false,
						IsSortStringsBySpace: false,
					},
					SearchAs: stringcompareas.Equal,
				},
			},
		},
		IsSkipOnContentsEmpty: false,
		IsCaseSensitive:       true,
		ExpectationLines:      []string{},
	},
	{
		Header: "IsTrimCompare with IsNonEmptyWhitespace true should should match the text and no error",
		ComparingLines: []string{
			"alim alim 2 alim 3",
		},
		Validators: corevalidator.TextValidators{
			Items: []corevalidator.TextValidator{
				{
					Search: "   alim      alim 2 alim 3                 ",
					ValidatorCoreCondition: corevalidator.ValidatorCoreCondition{
						IsTrimCompare:        true,
						IsUniqueWordOnly:     false,
						IsNonEmptyWhitespace: true,
						IsSortStringsBySpace: false,
					},
					SearchAs: stringcompareas.Equal,
				},
			},
		},
		IsSkipOnContentsEmpty: false,
		IsCaseSensitive:       true,
		ExpectationLines:      []string{},
	},
	{
		Header: "IsTrimCompare with IsNonEmptyWhitespace true different text and multiple search should give 2 errors",
		ComparingLines: []string{
			"alim alim 2 alim 4",
			"alim alim 2 alim 3",
			"alim alim 2 alim 5",
		},
		Validators: corevalidator.TextValidators{
			Items: []corevalidator.TextValidator{
				{
					Search: "   alim      alim 2 alim 3                 ",
					ValidatorCoreCondition: corevalidator.ValidatorCoreCondition{
						IsTrimCompare:        true,
						IsUniqueWordOnly:     false,
						IsNonEmptyWhitespace: true,
						IsSortStringsBySpace: false,
					},
					SearchAs: stringcompareas.Equal,
				},
			},
		},
		IsSkipOnContentsEmpty: false,
		IsCaseSensitive:       true,
		ExpectationLines: []string{
			"0 )\tExpectationLines failed: Failed match method [\"Equal\"], Index : [0]",
			"     Actual-Processed: `\"alim alim 2 alim 4\"`",
			"   Expected-Processed: `\"alim alim 2 alim 3\"`",
			"0 )\tExpectationLines failed: Failed match method [\"Equal\"], Index : [2]",
			"     Actual-Processed: `\"alim alim 2 alim 5\"`",
			"   Expected-Processed: `\"alim alim 2 alim 3\"`",
		},
	},
	{
		Header: "IsTrimCompare, IsSortStringsBySpace with IsNonEmptyWhitespace true " +
			"different text and multiple search should give 2 errors",
		ComparingLines: []string{
			"alim alim 2 alim 4",
			"alim alim 2 alim 3",
			"alim alim 2 alim 5",
		},
		Validators: corevalidator.TextValidators{
			Items: []corevalidator.TextValidator{
				{
					Search: "   alim      3 alim 2 alim                 ",
					ValidatorCoreCondition: corevalidator.ValidatorCoreCondition{
						IsTrimCompare:        true,
						IsUniqueWordOnly:     false,
						IsNonEmptyWhitespace: true,
						IsSortStringsBySpace: true,
					},
					SearchAs: stringcompareas.Equal,
				},
			},
		},
		IsSkipOnContentsEmpty: false,
		IsCaseSensitive:       true,
		ExpectationLines: []string{
			"0 )\tExpectationLines failed: Failed match method [\"Equal\"], Index : [0]",
			"     Actual-Processed: `\"2 4 alim alim alim\"`",
			"   Expected-Processed: `\"2 3 alim alim alim\"`",
			"0 )\tExpectationLines failed: Failed match method [\"Equal\"], Index : [2]",
			"     Actual-Processed: `\"2 5 alim alim alim\"`",
			"   Expected-Processed: `\"2 3 alim alim alim\"`",
		},
	},
	{
		Header: "All flags true different text and multiple search should give 2 errors",
		ComparingLines: []string{
			"alim alim 2 alim 4",
			"alim alim 2 alim 3",
			"alim alim 2 alim 5",
		},
		Validators: corevalidator.TextValidators{
			Items: []corevalidator.TextValidator{
				{
					Search: "   alim      alim 2 alim 3                 ",
					ValidatorCoreCondition: corevalidator.ValidatorCoreCondition{
						IsTrimCompare:        true,
						IsUniqueWordOnly:     true,
						IsNonEmptyWhitespace: true,
						IsSortStringsBySpace: true,
					},
					SearchAs: stringcompareas.Equal,
				},
			},
		},
		IsSkipOnContentsEmpty: false,
		IsCaseSensitive:       true,
		ExpectationLines: []string{
			"0 )\tExpectationLines failed: Failed match method [\"Equal\"], Index : [0]",
			"     Actual-Processed: `\"2 4 alim\"`",
			"   Expected-Processed: `\"2 3 alim\"`",
			"0 )\tExpectationLines failed: Failed match method [\"Equal\"], Index : [2]",
			"     Actual-Processed: `\"2 5 alim\"`",
			"   Expected-Processed: `\"2 3 alim\"`",
		},
	},
}
