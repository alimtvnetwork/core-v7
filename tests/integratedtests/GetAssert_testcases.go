package integratedtests

import (
	"fmt"
	"reflect"

	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var (
	commonType = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf(args.Map{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	quickTestCases = []coretestcases.CaseV1{
		{
			Title: "Quick output as gherkins format",
			ArrangeInput: args.Map{
				"when":    "some title, or case when",
				"actual":  "actual rec",
				"expect":  "expected item",
				"counter": 3,
			},
			ExpectedInput: []string{
				"----------------------",
				"3 )  When: some title, or case when,",
				"   Actual: actual rec,",
				" Expected: expected item",
			},
			VerifyTypeOf: commonType,
		},
	}

	sortedArrayTestCases = []coretestcases.CaseV1{
		{
			Title: "SortedArray output verification",
			ArrangeInput: args.Map{
				"isPrint": false,
				"isSort":  true,
				"message": "some message alim, knows, who, do you know --- #alim",
			},
			ExpectedInput: []string{
				"#alim",
				"---",
				"alim,",
				"do",
				"know",
				"knows,",
				"message",
				"some",
				"who,",
				"you",
			},
			VerifyTypeOf: commonType,
		},
	}

	sortedMessageTestCases = []coretestcases.CaseV1{
		{
			Title: "SortedMessage output verification",
			ArrangeInput: args.Map{
				"isPrint": false,
				"message": "some message alim, knows, who, do you know --- #alim",
				"joiner":  " | ",
			},
			ExpectedInput: []string{
				"#alim | --- | alim, | do | know | knows, | message | some | who, | you",
			},
			VerifyTypeOf: commonType,
		},
	}

	stringsToSpaceStringTestCases = []coretestcases.CaseV1{
		{
			Title: "StringsToWithSpaceLines output verification",
			ArrangeInput: args.Map{
				"spaceCount": 4,
				"lines": []string{
					"#alim",
					"---",
					"alim,",
					"do",
					"",
					"know",
					"when",
					"you",
					"type,",
					"lines",
				},
			},
			ExpectedInput: []string{
				"    #alim",
				"    ---",
				"    alim,",
				"    do",
				"    ",
				"    know",
				"    when",
				"    you",
				"    type,",
				"    lines",
			},
			VerifyTypeOf: commonType,
		},
	}

	toStringsWithTabTestCases = []coretestcases.CaseV1{
		{
			Title: "given []string in any parameter with 4 spaces",
			ArrangeInput: args.Map{
				"spaceCount": 4,
				"any": []string{
					"#alim",
					"---",
					"some,",
					"any lines",
				},
			},
			ExpectedInput: []string{
				"    #alim",
				"    ---",
				"    some,",
				"    any lines",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "given []interface{} in any parameter with 4 spaces",
			ArrangeInput: args.Map{
				"spaceCount": 4,
				"any": []interface{}{
					"#alim",
					"---",
					args.Map{
						"key": args.One{
							First: []string{
								"line alim 1",
								"line alim 2",
							},
							Expect: "alim expect",
						},
					},
					"any lines",
					1,
					255,
					true,
					args.One{
						Expect: "alim expect",
					},
				},
			},
			ExpectedInput: []string{
				"    #alim",
				"    ---",
				"    {\"key\":{\"FirstItem\":[\"line alim 1\",\"line alim 2\"],\"Expect\":\"alim expect\"}}",
				"    any lines",
				"    1",
				"    255",
				"    true",
				"    {\"Expect\":\"alim expect\"}",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "given args.One in any parameter with 4 spaces - does Pretty JSON with spaces",
			ArrangeInput: args.Map{
				"spaceCount": 4,
				"any": args.One{
					First: []string{
						"line alim 1",
						"line alim 2",
					},
					Expect: "alim expect",
				},
			},
			ExpectedInput: []string{
				"    {",
				"      \"FirstItem\": [",
				"        \"line alim 1\",",
				"        \"line alim 2\"",
				"      ],",
				"      \"Expect\": \"alim expect\"",
				"    }",
			},
			VerifyTypeOf: commonType,
		},
	}

	stringsToSpaceStringUsingFuncTestCases = []coretestcases.CaseV1{
		{
			Title: "given lines 4 spaces - displays as {space}%d. 'line';",
			ArrangeInput: args.Map{
				"spaceCount": 4,
				"converterFunc": coretests.ToLineConverterFunc(
					func(i int, spacePrefix, line string) string {
						return fmt.Sprintf(
							"%s %d. '%s';",
							spacePrefix,
							i+1,
							line,
						)
					},
				),
				"lines": []string{
					"alim introduced",
					"new custom formatter",
					"lets see the format",
				},
			},
			ExpectedInput: []string{
				"     1. 'alim introduced';",
				"     2. 'new custom formatter';",
				"     3. 'lets see the format';",
			},
			VerifyTypeOf: commonType,
		},
	}

	anyToDoubleQuoteLinesTestCases = []coretestcases.CaseV1{
		{
			Title: "AnyToDoubleQuoteLines verification test",
			Parameters: &args.Holder{
				First: 4,
				Second: []string{
					"line 1",
					"line 2",
					"line 3",
					"line 4",
					"line 5",
					"line 6",
				},
			},
			ExpectedInput: []string{
				"    \"line 1\",",
				"    \"line 2\",",
				"    \"line 3\",",
				"    \"line 4\",",
				"    \"line 5\",",
				"    \"line 6\",",
			},
		},
		{
			Title: "AnyToDoubleQuoteLines => nil given doesn't panic",
			Parameters: &args.Holder{
				First:  4,
				Second: nil,
			},
			ExpectedInput: []string{},
		},
		{
			Title: "AnyToDoubleQuoteLines => empty slice returns valid result",
			Parameters: &args.Holder{
				First:  4,
				Second: []string{},
			},
			ExpectedInput: []string{},
		},
		{
			Title: "AnyToDoubleQuoteLines => map[string]string provides concat line.",
			Parameters: &args.Holder{
				First: 4,
				Second: map[string]string{
					"line 1": "some line 1",
					"line 2": "some line 2",
					"line 3": "some line 3",
				},
			},
			ExpectedInput: []string{
				"    \"line 1 : some line 1\",",
				"    \"line 2 : some line 2\",",
				"    \"line 3 : some line 3\",",
			},
		},
		{
			Title: "AnyToDoubleQuoteLines []interface{} any array provide nice lines",
			Parameters: &args.Holder{
				First: 4,
				Second: []interface{}{
					"line 1",
					2,
					args.Map{
						"some key": "some val",
					},
					[]string{
						"line 1",
						"line 2",
					},
					"line 3",
					"line 4",
				},
			},
			ExpectedInput: []string{
				"    \"line 1\",",
				"    \"2\",",
				"    \"{\\\"some key\\\":\\\"some val\\\"}\",",
				"    \"line 1\\nline 2\",",
				"    \"line 3\",",
				"    \"line 4\",",
			},
		},
	}

	convertLinesToDoubleQuoteThenStringTestCases = []coretestcases.CaseV1{
		{
			Title: "ConvertLinesToDoubleQuoteThenString => convert []string to double quote string lines then to a single one",
			Parameters: &args.Holder{
				First: 4,
				Second: []string{
					"line 1",
					"line 2",
					"line 3",
					"line 4",
					"line 5",
					"line 6",
				},
			},
			ExpectedInput: []string{
				"    \"line 1\",\n    \"line 2\",\n    \"line 3\",\n    \"line 4\",\n    \"line 5\",\n    \"line 6\",",
			},
		},

		{
			Title: "ConvertLinesToDoubleQuoteThenString => convert []string - empty slice return simple empty string",
			Parameters: &args.Holder{
				First:  4,
				Second: []string{},
			},
			ExpectedInput: []string{
				"",
			},
		},
	}
)
