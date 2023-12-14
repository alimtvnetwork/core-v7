package integratedtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var (
	toStringsTestCases = []coretestcases.CaseV1{
		{
			Title: "giving string - output split to lines by newlines",
			ArrangeInput: args.Map{
				"any": "some string contains\nnewline\nin between",
			},
			ExpectedInput: []string{
				"some string contains",
				"newline",
				"in between",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving []string or slice string - outputs as is.",
			ArrangeInput: args.Map{
				"any": []string{
					"having exact lines will output",
					"as the lines",
					"were.",
					"no change.",
				},
			},
			ExpectedInput: []string{
				"having exact lines will output",
				"as the lines",
				"were.",
				"no change.",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving []string{} outputs as it is - empty string has no issues.",
			ArrangeInput: args.Map{
				"any": []string{},
			},
			ExpectedInput: []string{},
			VerifyTypeOf:  commonType,
		},
		{
			Title: "giving []interface - json convert and returns as it is.",
			ArrangeInput: args.Map{
				"any": []interface{}{
					"passed []interface, which is",
					"any but lines of any",
					"gets no converted and",
					"returns as it is",
				},
			},
			ExpectedInput: []string{
				"passed []interface, which is",
				"any but lines of any",
				"gets no converted and",
				"returns as it is",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving map[string]interface{} - converts to lines and returns sorted lines.",
			ArrangeInput: args.Map{
				"any": map[string]interface{}{
					"line 1": "passed map[string]interface, which is",
					"line 2": "any but keys as is but converts",
					"line 3": "value to SmartJSON and",
					"line 4": map[string]interface{}{
						"sub line 1": "returns",
						"sub line 2": -5,
					},
					"line 5": []string{
						"some line 1",
						"some line 2",
					},
					"line 6": []interface{}{
						args.One{
							First:  "line 6.1 first",
							Expect: "line 6.1 expect",
						},
						"some line 2",
					},
				},
			},
			ExpectedInput: []string{
				"line 1 : passed map[string]interface, which is",
				"line 2 : any but keys as is but converts",
				"line 3 : value to SmartJSON and",
				"line 4 : {\"sub line 1\":\"returns\",\"sub line 2\":-5}",
				"line 5 : some line 1\nsome line 2",
				"line 6 : [{\"FirstItem\":\"line 6.1 first\",\"Expect\":\"line 6.1 expect\"},\"some line 2\"]",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving map[interface{}]interface{} - converts to lines and returns sorted lines.",
			ArrangeInput: args.Map{
				"any": map[interface{}]interface{}{
					0:        "it is 0",
					1:        []string{"it is 1"},
					"line 1": "passed map[interface{}]interface{}, which is",
					"line 2": "converts both keys and values to",
					"line 3": "SmartJSON and returns it.",
					"line 4": map[string]interface{}{
						"sub line 1": "returns",
						"sub line 2": -5,
					},
					"line 5": []string{
						"some line 1",
						"some line 2",
					},
					"line 6": []interface{}{
						args.One{
							First:  "line 6.1 first",
							Expect: "line 6.1 expect",
						},
						"some line 2",
					},
					args.One{
						First: "line 7 - key",
					}: args.One{
						First:  "line 7 - value",
						Expect: "line 7 - value.expect",
					},
				},
			},
			ExpectedInput: []string{
				"0 : it is 0",
				"1 : it is 1",
				"line 1 : passed map[interface{}]interface{}, which is",
				"line 2 : converts both keys and values to",
				"line 3 : SmartJSON and returns it.",
				"line 4 : {\"sub line 1\":\"returns\",\"sub line 2\":-5}",
				"line 5 : some line 1\nsome line 2",
				"line 6 : [{\"FirstItem\":\"line 6.1 first\",\"Expect\":\"line 6.1 expect\"},\"some line 2\"]",
				"{\"FirstItem\":\"line 7 - key\"} : {\"FirstItem\":\"line 7 - value\",\"Expect\":\"line 7 - value.expect\"}",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving map[string]string - converts to lines and returns sorted lines.",
			ArrangeInput: args.Map{
				"any": map[string]string{
					"line 1": "passed map[string]string, which is",
					"line 2": "any but keys as is but converts",
					"line 3": "value to as is and",
					"line 4": "returns simple line",
				},
			},
			ExpectedInput: []string{
				"line 1 : passed map[string]string, which is",
				"line 2 : any but keys as is but converts",
				"line 3 : value to as is and",
				"line 4 : returns simple line",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving map[string]int - converts to lines and returns sorted lines.",
			ArrangeInput: args.Map{
				"any": map[string]int{
					"line 1": 1,
					"line 2": 2,
					"line 3": 3,
					"line 4": 4,
				},
			},
			ExpectedInput: []string{
				"line 1 : 1",
				"line 2 : 2",
				"line 3 : 3",
				"line 4 : 4",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving map[string]int - converts to lines and returns sorted lines.",
			ArrangeInput: args.Map{
				"any": map[int]string{
					1: "line 1",
					2: "line 2",
					3: "line 3",
					4: "line 4",
				},
			},
			ExpectedInput: []string{
				"1 : line 1",
				"2 : line 2",
				"3 : line 3",
				"4 : line 4",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving int - gives []string { int }",
			ArrangeInput: args.Map{
				"any": 321,
			},
			ExpectedInput: []string{
				"321",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving byte - gives []string { byte }",
			ArrangeInput: args.Map{
				"any": byte(156),
			},
			ExpectedInput: []string{
				"156",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving bool - gives []string { bool }",
			ArrangeInput: args.Map{
				"any": true,
			},
			ExpectedInput: []string{
				"true",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "giving args.One - converts to Smart JSON.",
			ArrangeInput: args.Map{
				"any": args.One{
					First: []string{
						"line 1",
						"line 2",
					},
					Expect: []string{
						"expect 1",
						"expect 2",
					},
				},
			},
			ExpectedInput: []string{
				"{",
				"  \"FirstItem\": [",
				"    \"line 1\",",
				"    \"line 2\"",
				"  ],",
				"  \"Expect\": [",
				"    \"expect 1\",",
				"    \"expect 2\"",
				"  ]",
				"}",
			},
			VerifyTypeOf: commonType,
		},
	}
)
