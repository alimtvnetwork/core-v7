package corecsvtests

import (
	"fmt"
	"reflect"

	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/issetter"
)

var (
	defaultTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]string{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}
	arrangeInterfaceArrayTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]interface{}{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	arrangeFmtStringerTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]fmt.Stringer{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	anyItemsToCsvStringSingleQuoteTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be displayed as csv. " +
					"On all true options, it will look like format: '%s', eg. '%s', '%s', '%s'...",
				ArrangeInput: []interface{}{
					1,
					2,
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
					"'1', " +
						"'2', " +
						"'alim', " +
						"'created', " +
						"'{curly}', " +
						"'which wraps', " +
						"'', " +
						"'any string to', " +
						"'curly', " +
						"'even empty ones', " +
						"'and', '{curly ones}', " +
						"'{left curly exists', " +
						"'right curly exists}'",
				},
				VerifyTypeOf: arrangeInterfaceArrayTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	anyItemsToCsvStringDoubleQuoteTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be displayed as csv. " +
					"It will look like format: \"%s\", eg. \"%s\", \"%s\", \"%s\"...",
				ArrangeInput: []interface{}{
					1,
					2,
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
					"\"1\", " +
						"\"2\", " +
						"\"alim\", " +
						"\"created\", " +
						"\"{curly}\", " +
						"\"which wraps\", " +
						"\"\", " +
						"\"any string to\", " +
						"\"curly\", " +
						"\"even empty ones\", " +
						"\"and\", " +
						"\"{curly ones}\", " +
						"\"{left curly exists\", " +
						"\"right curly exists}\"",
				},
				VerifyTypeOf: arrangeInterfaceArrayTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
	anyItemsToCsvStringNoQuoteTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be displayed as csv. " +
					"It will look like format: %s, eg. %s, %s, %s...",
				ArrangeInput: []interface{}{
					1,
					2,
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
					"1, 2, " +
						"alim, created, {curly}, which wraps, , " +
						"any string to, curly, even empty ones, " +
						"and, {curly ones}, " +
						"{left curly exists, right curly exists}",
				},
				VerifyTypeOf: arrangeInterfaceArrayTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
	stringsToCsvStringSingleQuoteTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be displayed as csv. " +
					"On all true options, it will look like format: '%s', eg. '%s', '%s', '%s'...",
				ArrangeInput: []string{
					"1",
					"2",
					"alim",
					"created",
					"{curly}",
					"",
					"any string to",
					"and",
					"{curly ones}",
					"{left curly exists",
					"right curly exists}",
				},
				ExpectedInput: []string{
					"'1', " +
						"'2', " +
						"'alim', " +
						"'created', " +
						"'{curly}', " +
						"'', " +
						"'any string to', " +
						"'and', " +
						"'{curly ones}', " +
						"'{left curly exists', " +
						"'right curly exists}'",
				},
				VerifyTypeOf: defaultTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
	stringsToCsvStringDoubleQuoteTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be displayed as csv. " +
					"Double quote format: \"%s\", eg. \"%s\", \"%s\", \"%s\"...",
				ArrangeInput: []string{
					"1",
					"2",
					"alim",
					"created",
					"{curly}",
					"",
					"any string to",
					"and",
					"{curly ones}",
					"{left curly exists",
					"right curly exists}",
				},
				ExpectedInput: []string{
					"\"1\", \"2\", \"alim\", \"created\", \"{curly}\", \"\", \"any string to\", \"and\", \"{curly ones}\", \"{left curly exists\", \"right curly exists}\"",
				},
				VerifyTypeOf: defaultTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
	stringsToCsvStringNoQuoteTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be displayed as csv. " +
					"No quote format: %s, eg. %s, %s, %s...",
				ArrangeInput: []string{
					"1",
					"2",
					"alim",
					"created",
					"{curly}",
					"",
					"any string to",
					"and",
					"{curly ones}",
					"{left curly exists",
					"right curly exists}",
				},
				ExpectedInput: []string{
					"1, 2, alim, created, {curly}, , any string to, and, {curly ones}, {left curly exists, right curly exists}",
				},
				VerifyTypeOf: defaultTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	rangeNamesWithValuesIndexesTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be displayed as Name[%d]. " +
					"format: SomeString[Index], eg. SomeString[0]",
				ArrangeInput: []string{
					"some val at 0",
					"some val at 1",
					"some val at 2",
					"Alim Ul Karim",
					"Where It is",
					"",
				},
				ExpectedInput: []string{
					"some val at 0[0]",
					"some val at 1[1]",
					"some val at 2[2]",
					"Alim Ul Karim[3]",
					"Where It is[4]",
					"[5]",
				},
				VerifyTypeOf: defaultTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	rangeNamesWithValuesIndexesStringTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will be displayed as Name[%d]. " +
					"format: SomeString[Index], eg. SomeString[0]",
				ArrangeInput: []string{
					"some val at 0",
					"some val at 1",
					"some val at 2",
					"Alim Ul Karim",
					"Where It is",
					"",
				},
				ExpectedInput: []string{
					"some val at 0[0], some val at 1[1], some val at 2[2], Alim Ul Karim[3], Where It is[4], [5]",
				},
				VerifyTypeOf: defaultTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	stringersTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given strings will join like Csv single quote join.",
				ArrangeInput: []fmt.Stringer{
					coretests.SomeString{
						Value: "some value",
					},
					coretests.SomeString{
						Value: "alim",
					},
					coretests.SomeString{},
					coretests.SomeString{
						Value: "this is stringer",
					},
				},
				ExpectedInput: []string{
					"'some value', 'alim', '', 'this is stringer'",
				},
				VerifyTypeOf: arrangeFmtStringerTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	anyToTypesCsvStringsSingleQuoteTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given items types will be converted into string with a single quote.",
				ArrangeInput: []interface{}{
					coretests.SomeString{
						Value: "some value",
					},
					&coretests.SomeString{
						Value: "alim",
					},
					nil,
					"Hello",
					1,
					[]string{},
				},
				ExpectedInput: []string{
					"'coretests.SomeString'",
					"'*coretests.SomeString'",
					"'<nil>'",
					"'string'",
					"'int'",
					"'[]string'",
				},
				VerifyTypeOf: arrangeInterfaceArrayTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	anyToTypesCsvStringsDoubleQuoteTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given items types will be converted into string with a double quote.",
				ArrangeInput: []interface{}{
					coretests.SomeString{
						Value: "some value",
					},
					&coretests.SomeString{
						Value: "alim",
					},
					nil,
					"Hello",
					1,
					[]string{},
				},
				ExpectedInput: []string{
					"\"coretests.SomeString\"",
					"\"*coretests.SomeString\"",
					"\"<nil>\"",
					"\"string\"",
					"\"int\"",
					"\"[]string\"",
				},
				VerifyTypeOf: arrangeInterfaceArrayTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	anyToTypesCsvStringsNoQuoteTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Given items types will be converted into string without any quote.",
				ArrangeInput: []interface{}{
					coretests.SomeString{
						Value: "some value",
					},
					&coretests.SomeString{
						Value: "alim",
					},
					nil,
					"Hello",
					1,
					[]string{},
				},
				ExpectedInput: []string{
					"coretests.SomeString",
					"*coretests.SomeString",
					"<nil>",
					"string",
					"int",
					"[]string",
				},
				VerifyTypeOf: arrangeInterfaceArrayTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
)
