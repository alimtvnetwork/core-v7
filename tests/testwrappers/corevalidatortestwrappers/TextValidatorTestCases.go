package corevalidatortestwrappers

import (
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
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
					ValidatorCoreCondition: corevalidator.ValidatorCoreCondition{
						IsTrimCompare:        false,
						IsUniqueWordOnly:     false,
						IsNonEmptyWhitespace: false,
						IsSortStringsBySpace: false,
					},
				},
			},
		},
		IsSkipOnContentsEmpty: false,
		IsCaseSensitive:       true,
		ExpectationLines: []string{
			"0 )\t  Expectation failed: Using CompareMethod `\"Equal\"`, Line Index: 0",
			"     Content-Processed:`\"alim      alim 2 alim 4\"`",
			"  SearchTerm-Processed:`\"   alim      alim 2 alim 3                 \"`",
			"            Additional:`\"Search Input: [`   alim      alim 2 alim 3                 `]," +
				" CompareMethod: [`Equal`], " +
				"IsTrimCompare: [`false`], " +
				"IsSplitByWhitespace: [`false`], " +
				"IsUniqueWordOnly: [`false`], " +
				"IsNonEmptyWhitespace: [`false`], " +
				"IsSortStringsBySpace: [`false`]\"`",
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
			"0 )\t  Expectation failed: Using CompareMethod `\"Equal\"`, Line Index: 0",
			"     Content-Processed:`\"alim alim 2 alim 4\"`",
			"  SearchTerm-Processed:`\"alim alim 2 alim 3\"`",
			"            Additional:`\"Search Input: [`   alim      alim 2 alim 3                 `], " +
				"CompareMethod: [`Equal`], " +
				"IsTrimCompare: [`true`], " +
				"IsSplitByWhitespace: [`true`], " +
				"IsUniqueWordOnly: [`false`], " +
				"IsNonEmptyWhitespace: [`true`], " +
				"IsSortStringsBySpace: [`false`]\"`",
			"0 )\t  Expectation failed: Using CompareMethod `\"Equal\"`, Line Index: 2",
			"     Content-Processed:`\"alim alim 2 alim 5\"`",
			"  SearchTerm-Processed:`\"alim alim 2 alim 3\"`",
			"            Additional:`\"Search Input: [`   alim      alim 2 alim 3                 `], " +
				"CompareMethod: [`Equal`], " +
				"IsTrimCompare: [`true`], " +
				"IsSplitByWhitespace: [`true`], " +
				"IsUniqueWordOnly: [`false`], " +
				"IsNonEmptyWhitespace: [`true`], " +
				"IsSortStringsBySpace: [`false`]\"`",
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
			"0 )\t  Expectation failed: Using CompareMethod `\"Equal\"`, Line Index: 0",
			"     Content-Processed:`\"2 4 alim alim alim\"`",
			"  SearchTerm-Processed:`\"2 3 alim alim alim\"`",
			"            Additional:`\"Search Input: [`   alim      3 alim 2 alim                 `], " +
				"CompareMethod: [`Equal`], " +
				"IsTrimCompare: [`true`], " +
				"IsSplitByWhitespace: [`true`], " +
				"IsUniqueWordOnly: [`false`], " +
				"IsNonEmptyWhitespace: [`true`], " +
				"IsSortStringsBySpace: [`true`]\"`",
			"0 )\t  Expectation failed: Using CompareMethod `\"Equal\"`, Line Index: 2",
			"     Content-Processed:`\"2 5 alim alim alim\"`",
			"  SearchTerm-Processed:`\"2 3 alim alim alim\"`",
			"            Additional:`\"Search Input: [`   alim      3 alim 2 alim                 `], " +
				"CompareMethod: [`Equal`], " +
				"IsTrimCompare: [`true`], " +
				"IsSplitByWhitespace: [`true`], " +
				"IsUniqueWordOnly: [`false`], " +
				"IsNonEmptyWhitespace: [`true`], " +
				"IsSortStringsBySpace: [`true`]\"`",
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
			"0 )\t  Expectation failed: Using CompareMethod `\"Equal\"`, Line Index: 0",
			"     Content-Processed:`\"2 4 alim\"`",
			"  SearchTerm-Processed:`\"2 3 alim\"`",
			"            Additional:`\"Search Input: [`   alim      alim 2 alim 3                 `], " +
				"CompareMethod: [`Equal`], " +
				"IsTrimCompare: [`true`], " +
				"IsSplitByWhitespace: [`true`], " +
				"IsUniqueWordOnly: [`true`], " +
				"IsNonEmptyWhitespace: [`true`], " +
				"IsSortStringsBySpace: [`true`]\"`",
			"0 )\t  Expectation failed: Using CompareMethod `\"Equal\"`, Line Index: 2",
			"     Content-Processed:`\"2 5 alim\"`",
			"  SearchTerm-Processed:`\"2 3 alim\"`",
			"            Additional:`\"Search Input: [`   alim      alim 2 alim 3                 `], " +
				"CompareMethod: [`Equal`], " +
				"IsTrimCompare: [`true`], " +
				"IsSplitByWhitespace: [`true`], " +
				"IsUniqueWordOnly: [`true`], " +
				"IsNonEmptyWhitespace: [`true`], " +
				"IsSortStringsBySpace: [`true`]\"`",
		},
	},
	{
		Header: "All flags true different text and multiple search NOT " +
			"equal method gives only one error for equal one.",
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
					SearchAs: stringcompareas.NotEqual,
				},
			},
		},
		IsSkipOnContentsEmpty: false,
		IsCaseSensitive:       true,
		ExpectationLines: []string{
			"0 )\t  Expectation failed: Using CompareMethod `\"NotEqual\"`, Line Index: 1",
			"     Content-Processed:`\"2 3 alim\"`",
			"  SearchTerm-Processed:`\"2 3 alim\"`",
			"            Additional:`\"Search Input: [`   alim      alim 2 alim 3                 `], " +
				"CompareMethod: [`NotEqual`], " +
				"IsTrimCompare: [`true`], " +
				"IsSplitByWhitespace: [`true`], " +
				"IsUniqueWordOnly: [`true`], " +
				"IsNonEmptyWhitespace: [`true`], " +
				"IsSortStringsBySpace: [`true`]\"`",
		},
	},
	{
		Header: "Nonsensitive - All flags true different text and multiple search NOT " +
			"equal method gives only one error for matching one for equal one.",
		ComparingLines: []string{
			"alim alim 2 alim 4",
			"Alim alim 2 Alim 3",
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
					SearchAs: stringcompareas.NotEqual,
				},
			},
		},
		IsSkipOnContentsEmpty: false,
		IsCaseSensitive:       false,
		ExpectationLines: []string{
			"0 )\t  Expectation failed: Using CompareMethod `\"NotEqual\"`, Line Index: 1",
			"     Content-Processed:`\"2 3 alim\"`",
			"  SearchTerm-Processed:`\"2 3 alim\"`",
			"            Additional:`\"Search Input: [`   alim      alim 2 alim 3                 `], " +
				"CompareMethod: [`NotEqual`], " +
				"IsTrimCompare: [`true`], " +
				"IsSplitByWhitespace: [`true`], " +
				"IsUniqueWordOnly: [`true`], " +
				"IsNonEmptyWhitespace: [`true`], " +
				"IsSortStringsBySpace: [`true`]\"`",
		},
	},
	{
		Header: "Nonsensitive - All flags true different text and multiple search NOT " +
			"equal method gives only one error for matching one for equal one.",
		ComparingLines: []string{
			"Alim alim 2 alim 4",
			"Alim alim 2 Alim 3",
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
					SearchAs: stringcompareas.NotEqual,
				},
			},
		},
		IsSkipOnContentsEmpty: false,
		IsCaseSensitive:       false,
		ExpectationLines: []string{
			"0 )\t  Expectation failed: Using CompareMethod `\"NotEqual\"`, Line Index: 1",
			"     Content-Processed:`\"2 3 alim\"`",
			"  SearchTerm-Processed:`\"2 3 alim\"`",
			"            Additional:`\"Search Input: [`   alim      alim 2 alim 3                 `], " +
				"CompareMethod: [`NotEqual`], " +
				"IsTrimCompare: [`true`], " +
				"IsSplitByWhitespace: [`true`], " +
				"IsUniqueWordOnly: [`true`], " +
				"IsNonEmptyWhitespace: [`true`], " +
				"IsSortStringsBySpace: [`true`]\"`",
		},
	},
	{
		Header: "Nonsensitive - All flags true different text and multiple search NOT " +
			"equal method gives only one error for matching one for equal one.",
		ComparingLines: []string{
			"Alim alim 2 alim 4",
			"Alim alim 2 Alim 3",
			"alim alim 2 alim 5",
		},
		Validators: corevalidator.TextValidators{
			Items: []corevalidator.TextValidator{
				{
					Search: "   alim      alim 2 Alim 3                 ",
					ValidatorCoreCondition: corevalidator.ValidatorCoreCondition{
						IsTrimCompare:        true,
						IsUniqueWordOnly:     false,
						IsNonEmptyWhitespace: true,
						IsSortStringsBySpace: true,
					},
					SearchAs: stringcompareas.NotEqual,
				},
			},
		},
		IsSkipOnContentsEmpty: false,
		IsCaseSensitive:       false,
		ExpectationLines: []string{
			"0 )\t  Expectation failed: Using CompareMethod `\"NotEqual\"`, Line Index: 1",
			"     Content-Processed:`\"2 3 Alim Alim alim\"`",
			"  SearchTerm-Processed:`\"2 3 Alim alim alim\"`",
			"            Additional:`\"Search Input: [`   alim      alim 2 Alim 3                 `], " +
				"CompareMethod: [`NotEqual`], " +
				"IsTrimCompare: [`true`], " +
				"IsSplitByWhitespace: [`true`], " +
				"IsUniqueWordOnly: [`false`], " +
				"IsNonEmptyWhitespace: [`true`], " +
				"IsSortStringsBySpace: [`true`]\"`",
		},
	},
}
