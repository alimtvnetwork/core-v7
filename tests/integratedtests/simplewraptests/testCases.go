package simplewraptests

import (
	"reflect"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coredata/corerange"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/issetter"
)

var (
	stringsSliceTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]string{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	curlyWrapIfEnabledValidTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be curly wrapped without checking. " +
					"Curly wrapped guaranteed, duplicate curly wrap is possible.",
				ArrangeInput: []string{
					"alim",
					"created",
					"{curly}",
					"which wraps",
					"",
					"any string to",
					"curly",
					"even empty ones",
					"and",
					"{curly ones}",
					"{left curly exists",
					"right curly exists}",
				},
				ExpectedInput: []string{
					"{alim}",
					"{created}",
					"{{curly}}",
					"{which wraps}",
					"{}",
					"{any string to}",
					"{curly}",
					"{even empty ones}",
					"{and}",
					"{{curly ones}}",
					"{{left curly exists}",
					"{right curly exists}}",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
	curlyWrapIfDisabledValidTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be curly wrapped without checking. " +
					"Curly wrapped guaranteed, duplicate curly wrap is possible.",
				ArrangeInput: []string{
					"alim",
					"created",
					"{curly}",
					"which wraps",
					"",
					"any string to",
					"curly",
					"even empty ones",
					"and",
					"{curly ones}",
					"{left curly exists",
					"right curly exists}",
				},
				ExpectedInput: []string{
					"alim",
					"created",
					"{curly}",
					"which wraps",
					"",
					"any string to",
					"curly",
					"even empty ones",
					"and",
					"{curly ones}",
					"{left curly exists",
					"right curly exists}",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	curlyWrapOptionsValidTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be curly wrapped with checking. " +
					"Curly wrapped guaranteed, no duplicate curly wrap possible.",
				ArrangeInput: []string{
					"alim",
					"created",
					"{curly}",
					"which wraps",
					"",
					"any string to",
					"curly",
					"even empty ones",
					"and",
					"{curly ones}",
					"{left curly exists",
					"right curly exists}",
				},
				ExpectedInput: []string{
					"{alim}",
					"{created}",
					"{curly}",
					"{which wraps}",
					"{}",
					"{any string to}",
					"{curly}",
					"{even empty ones}",
					"{and}",
					"{curly ones}",
					"{left curly exists}",
					"{right curly exists}",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
	parenthesisValidTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be parenthesis ( wrapped ) with no checking. " +
					"Parenthesis wrapped guaranteed, duplicate wrap possible.",
				ArrangeInput: []string{
					"alim",
					"created",
					"(parenthesis)",
					"which wraps",
					"",
					"any string to",
					"parenthesis",
					"even empty ones",
					"and",
					"(parenthesis ones)",
					"(left parenthesis exists",
					"right parenthesis exists)",
				},
				ExpectedInput: []string{
					"(alim)",
					"(created)",
					"((parenthesis))",
					"(which wraps)",
					"()",
					"(any string to)",
					"(parenthesis)",
					"(even empty ones)",
					"(and)",
					"((parenthesis ones))",
					"((left parenthesis exists)",
					"(right parenthesis exists))",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	parenthesisDisabledRemainsAsItIsTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be parenthesis ( wrapped ) with no checking. " +
					"Parenthesis wrapped guaranteed, duplicate wrap possible.",
				ArrangeInput: []string{
					"alim",
					"created",
					"(parenthesis)",
					"which wraps",
					"",
					"any string to",
					"parenthesis",
					"even empty ones",
					"and",
					"(parenthesis ones)",
					"(left parenthesis exists",
					"right parenthesis exists)",
				},
				ExpectedInput: []string{
					"alim",
					"created",
					"(parenthesis)",
					"which wraps",
					"",
					"any string to",
					"parenthesis",
					"even empty ones",
					"and",
					"(parenthesis ones)",
					"(left parenthesis exists",
					"right parenthesis exists)",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	squareBracketWrapTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be square bracket [ wrapped ] with no checking. " +
					"Square bracket wrapped guaranteed, duplicate wrap possible.",
				ArrangeInput: []string{
					"alim",
					"created",
					"[sq bracket]",
					"which wraps",
					"",
					"any string to",
					"parenthesis",
					"even empty ones",
					"and",
					"[square]",
					"[left sq exists",
					"right sq exists]",
				},
				ExpectedInput: []string{
					"[alim]",
					"[created]",
					"[[sq bracket]]",
					"[which wraps]",
					"[]",
					"[any string to]",
					"[parenthesis]",
					"[even empty ones]",
					"[and]",
					"[[square]]",
					"[[left sq exists]",
					"[right sq exists]]",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	squareBracketWrapDisabledTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be NOT square bracket [ wrapped ]. " +
					"Square bracket wrapped is NOT guaranteed.",
				ArrangeInput: []string{
					"alim",
					"created",
					"[sq bracket]",
					"which wraps",
					"",
					"any string to",
					"parenthesis",
					"even empty ones",
					"and",
					"[square]",
					"[left sq exists",
					"right sq exists]",
				},
				ExpectedInput: []string{
					"alim",
					"created",
					"[sq bracket]",
					"which wraps",
					"",
					"any string to",
					"parenthesis",
					"even empty ones",
					"and",
					"[square]",
					"[left sq exists",
					"right sq exists]",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	titleCurlyMetaTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Title curly meta should look like - " +
					"title: {some-value} ({meta data}).",
				ArrangeInput: []string{
					"my title",       // title
					"some \"value\"", // value
					corejson.Serialize.ToString(map[string]string{
						"some-map-key": "Some meta information", // meta
					}),
				},
				ExpectedInput: []string{
					"my title: {some \"value\"} ({\"some-map-key\":\"Some meta information\"})",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Title curly meta should look like - " +
					"eg. title: {some-value} ({meta data}).",
				ArrangeInput: []string{
					"my title",        // title
					"some2 \"value\"", // value
					corejson.Serialize.ToString(map[string]string{
						"1": "2-meta", // meta
					}),
				},
				ExpectedInput: []string{
					"my title: {some2 \"value\"} ({\"1\":\"2-meta\"})",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	titleCurlyTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Title curly value should look like - " +
					"format: %s: {%s} - eg. title: {value}.",
				ArrangeInput: []string{
					"my title",       // title
					"some \"value\"", // value
				},
				ExpectedInput: []string{
					"my title: {some \"value\"}",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Title curly value should look like - " +
					"format: %s: {%s} - eg. title: {value}.",
				ArrangeInput: []string{
					"my next title",   // title
					"some2 \"value\"", // value
				},
				ExpectedInput: []string{
					"my next title: {some2 \"value\"}",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	msgCsvItemsTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Title curly value should look like - " +
					"format: %s: (%s) - eg. title: (csv values, ...).",
				ArrangeInput: []interface{}{
					"my title", // title
					[]interface{}{
						1,
						"some csv string",
						"some \"value\"",
						"to curly {no}",
					},
				},
				ExpectedInput: []string{
					"my title (\"1\", \"some csv string\", \"some \"value\"\", \"to curly {no}\")",
				},
				VerifyTypeOf: &coretests.VerifyTypeOf{
					ArrangeInput:  reflect.TypeOf([]interface{}{}),
					ActualInput:   reflect.TypeOf([]string{}),
					ExpectedInput: reflect.TypeOf([]string{}),
				},
				IsEnable: issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Title curly value should look like - " +
					"format : %s: (%s) - eg. title: (csv values, ...).",
				ArrangeInput: []interface{}{
					"my title", // title
					[]interface{}{
						1,
						5,
						9,
						"",
						nil,
						corerange.MinMaxInt{
							Min: 5,
							Max: 25,
						},
					},
				},
				ExpectedInput: []string{
					"my title (\"1\", \"5\", \"9\", \"\", \"<nil>\", \"5-25\")",
				},
				VerifyTypeOf: &coretests.VerifyTypeOf{
					ArrangeInput:  reflect.TypeOf([]interface{}{}),
					ActualInput:   reflect.TypeOf([]string{}),
					ExpectedInput: reflect.TypeOf([]string{}),
				},
				IsEnable: issetter.True,
			},
		},
	}

	msgWrapsMsgTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "msg wrap msg for both exist - first msg (second msg).",
				ArrangeInput: []string{
					"first \"alim\" msg",  // msg 1
					"second \"alim\" msg", // msg 2
				},
				ExpectedInput: []string{
					"first \"alim\" msg (second \"alim\" msg)",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "msg wrap msg for first exist, second doesn't - first msg.",
				ArrangeInput: []string{
					"first \"alim\" only msg", // msg 1
					"",
				},
				ExpectedInput: []string{
					"first \"alim\" only msg",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "msg wrap msg for first empty, second exist - 2nd msg.",
				ArrangeInput: []string{
					"",                         // msg 1
					"second \"alim\" only msg", // msg 2
				},
				ExpectedInput: []string{
					"second \"alim\" only msg",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "msg wrap msg for both empty - returns empty.",
				ArrangeInput: []string{
					"", // msg 1
					"", // msg 2
				},
				ExpectedInput: []string{
					"",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	withBracketsTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Title square bracket wraps - " +
					"format : [%s] - Eg. [value]. " +
					"Doesn't verify existence and may have duplicate brackets",
				ArrangeInput: []string{
					"alim",
					"created",
					"[sq bracket]",
					"which wraps",
					"",
					"any string to",
					"parenthesis",
					"even empty ones",
					"and",
					"[square]",
					"[left sq exists",
					"right sq exists]",
				},
				ExpectedInput: []string{
					"[alim]",
					"[created]",
					"[[sq bracket]]",
					"[which wraps]",
					"[]",
					"[any string to]",
					"[parenthesis]",
					"[even empty ones]",
					"[and]",
					"[[square]]",
					"[[left sq exists]",
					"[right sq exists]]",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	withBracketsQuotationTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Title square bracket wraps - " +
					"format : [\"%s\"] - Eg. [\"value\"]. " +
					"Doesn't verify existence and may have duplicate brackets",
				ArrangeInput: []string{
					"alim",
					"created",
					"[sq bracket]",
					"which wraps",
					"",
					"any string to",
					"parenthesis",
					"even empty ones",
					"and",
					"[square]",
					"[left sq exists",
					"right sq exists]",
				},
				ExpectedInput: []string{
					"[\"alim\"]",
					"[\"created\"]",
					"[\"[sq bracket]\"]",
					"[\"which wraps\"]",
					"[\"\"]",
					"[\"any string to\"]",
					"[\"parenthesis\"]",
					"[\"even empty ones\"]",
					"[\"and\"]",
					"[\"[square]\"]",
					"[\"[left sq exists\"]",
					"[\"right sq exists]\"]",
				},
				VerifyTypeOf: stringsSliceTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
)
