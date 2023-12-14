package coreversiontests

import (
	"reflect"

	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coreversion"
	"gitlab.com/auk-go/core/issetter"
)

var (
	arrangeLeftRightTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]args.LeftRight{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	argsFourTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]args.Four{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	argsFiveTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]args.Five{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	arrangeStringTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]string{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	defaultVersionTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]coreversion.Version{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	versionCreationTestCases = []testWrapper{
		{
			Title: "Create versions with different args and methods.",
			ArrangeInput: []coreversion.Version{
				coreversion.New.Invalid(),
				coreversion.New.Default("1.2.3.4"),
				coreversion.New.Default("5.3.6"),
				coreversion.New.Default("5.3"),
				coreversion.New.Default("9"),
				coreversion.New.Default("v1.2.3.4"),
				coreversion.New.Default("v5.3.6"),
				coreversion.New.Default("v5.3"),
				coreversion.New.Default("v9"),
				coreversion.New.Default(""),
			},
			ExpectedInput: []string{
				"0 : invalid - ",
				"1 : v1.2.3.4 (compact: 1.2.3.4, display: v1.2.3.4)",
				"2 : v5.3.6 (compact: 5.3.6, display: v5.3.6)",
				"3 : v5.3 (compact: 5.3, display: v5.3)",
				"4 : v9 (compact: 9, display: v9)",
				"5 : v1.2.3.4 (compact: 1.2.3.4, display: v1.2.3.4)",
				"6 : v5.3.6 (compact: 5.3.6, display: v5.3.6)",
				"7 : v5.3 (compact: 5.3, display: v5.3)",
				"8 : v9 (compact: 9, display: v9)",
				"9 : invalid - ",
			},
			VerifyTypeOf: defaultVersionTypeVerification,
			IsEnable:     issetter.True,
		},
	}

	versionCreationUsingStringTestCases = []testWrapper{
		{
			Title: "Create versions using string.",
			ArrangeInput: []string{
				"-1",
				"1.2.3.4",
				"5.3.6",
				"5.3",
				"9",
				"v1.2.3.4",
				"v5.3.6",
				"v5.3",
				"v9",
				"5.*.1",
				"7.*.*",
				"1.*.*.10",
				"-1.555.*.11",
				"v-1.576.*.12",
				"8v-1.581.*.13",
				"8v-1.x565.*.u14",
				"8v-1.*.*.u15",
				"v5.-5",
				"v10.-6",
				"v11...7",
				"v12...8",
				"xv12...9",
				"12..5",
				"13..6",
				"14..7",
				"",
			},
			ExpectedInput: []string{
				"0 : invalid - v-1 (raw: -1)",
				"1 : v1.2.3.4 (compact: 1.2.3.4, display: v1.2.3.4)",
				"2 : v5.3.6 (compact: 5.3.6, display: v5.3.6)",
				"3 : v5.3 (compact: 5.3, display: v5.3)",
				"4 : v9 (compact: 9, display: v9)",
				"5 : v1.2.3.4 (compact: 1.2.3.4, display: v1.2.3.4)",
				"6 : v5.3.6 (compact: 5.3.6, display: v5.3.6)",
				"7 : v5.3 (compact: 5.3, display: v5.3)",
				"8 : v9 (compact: 9, display: v9)",
				"9 : v5.0.1 (compact: 5.*.1, display: v5.*.1)",
				"10 : v7 (compact: 7.*.*, display: v7.*.*)",
				"11 : v1.0.0.10 (compact: 1.*.*.10, display: v1.*.*.10)",
				"12 : invalid - v0.555.0.11 (raw: -1.555.*.11)",
				"13 : invalid - v0.576.0.12 (raw: v-1.576.*.12)",
				"14 : invalid - v0.581.0.13 (raw: 8v-1.581.*.13)",
				"15 : invalid - v8v-1.x565.*.u14 (raw: 8v-1.x565.*.u14)",
				"16 : invalid - v8v-1.*.*.u15 (raw: 8v-1.*.*.u15)",
				"17 : invalid - v5 (raw: v5.-5)",
				"18 : invalid - v10 (raw: v10.-6)",
				"19 : v11.0.0.7 (compact: 11...7, display: v11...7)",
				"20 : v12.0.0.8 (compact: 12...8, display: v12...8)",
				"21 : invalid - v0.0.0.9 (raw: xv12...9)",
				"22 : v12.0.5 (compact: 12..5, display: v12..5)",
				"23 : v13.0.6 (compact: 13..6, display: v13..6)",
				"24 : v14.0.7 (compact: 14..7, display: v14..7)",
				"25 : invalid -  (raw: )",
			},
			VerifyTypeOf: arrangeStringTypeVerification,
			IsEnable:     issetter.True,
		},
	}

	comparisonStringTestCases = []testWrapper{
		{
			Title: "Versions comparisons - Left Greater",
			ArrangeInput: []args.LeftRight{
				{
					Left:   "1.2.5",
					Right:  "1.2.4",
					Expect: corecomparator.LeftGreater,
				},
				{
					Left:   "1.5.5",
					Right:  "1.*.8",
					Expect: corecomparator.LeftGreater,
				},
				{
					Left:   "5.2",
					Right:  "1.5",
					Expect: corecomparator.LeftGreater,
				},
				{
					Left:   "5.2",
					Right:  "5.2",
					Expect: corecomparator.LeftGreater,
				},
				{
					Left:   "5.2",
					Right:  "5.2",
					Expect: corecomparator.LeftGreaterEqual,
				},
			},
			ExpectedInput: []string{
				"0 : [ 1.2.5 ] > [ 1.2.4 ] | Expect: LeftGreater - true",
				"1 : [ 1.5.5 ] > [ 1.*.8 ] | Expect: LeftGreater - true",
				"2 : [ 5.2 ] > [ 1.5 ] | Expect: LeftGreater - true",
				"3 : [ 5.2 ] > [ 5.2 ] | Expect: LeftGreater - true",
				"4 : [ 5.2 ] >= [ 5.2 ] | Expect: LeftGreaterEqual - true",
			},
			VerifyTypeOf: arrangeLeftRightTypeVerification,
			IsEnable:     issetter.True,
		},
		{
			Title: "Versions comparisons - Left Less",
			ArrangeInput: []args.LeftRight{
				{
					Left:   "1.2",
					Right:  "1.2.1",
					Expect: corecomparator.LeftLess,
				},
				{
					Left:   "1.2",
					Right:  "1.2.1",
					Expect: corecomparator.LeftLess,
				},
				{
					Left:   "1.2",
					Right:  "1.5",
					Expect: corecomparator.LeftLess,
				},
				{
					Left:   "5.2",
					Right:  "5.1",
					Expect: corecomparator.LeftLess,
				},
				{
					Left:   "5.2",
					Right:  "5.1",
					Expect: corecomparator.LeftLessEqual,
				},
				{
					Left:   "*.2",
					Right:  "5.1",
					Expect: corecomparator.LeftLessEqual,
				},
				{
					Left:   "2.0.0.1",
					Right:  "2.0.0.5",
					Expect: corecomparator.LeftLess,
				},
				{
					Left:   "2.0.0.1",
					Right:  "2.0.0.1",
					Expect: corecomparator.LeftLessEqual,
				},
			},
			ExpectedInput: []string{
				"0 : [ 1.2 ] < [ 1.2.1 ] | Expect: LeftLess - true",
				"1 : [ 1.2 ] < [ 1.2.1 ] | Expect: LeftLess - true",
				"2 : [ 1.2 ] < [ 1.5 ] | Expect: LeftLess - true",
				"3 : [ 5.2 ] < [ 5.1 ] | Expect: LeftLess - false",
				"4 : [ 5.2 ] <= [ 5.1 ] | Expect: LeftLessEqual - false",
				"5 : [ *.2 ] <= [ 5.1 ] | Expect: LeftLessEqual - true",
				"6 : [ 2.0.0.1 ] < [ 2.0.0.5 ] | Expect: LeftLess - true",
				"7 : [ 2.0.0.1 ] <= [ 2.0.0.1 ] | Expect: LeftLessEqual - true",
			},
			VerifyTypeOf: arrangeLeftRightTypeVerification,
			IsEnable:     issetter.True,
		},
		{
			Title: "Versions comparisons - Equal",
			ArrangeInput: []args.LeftRight{
				{
					Left:   "v2.2",
					Right:  "v2.2.0",
					Expect: corecomparator.Equal,
				},
				{
					Left:   "v2.2",
					Right:  "2.2.0.0",
					Expect: corecomparator.Equal,
				},
				{
					Left:   "2.2",
					Right:  "v2.2",
					Expect: corecomparator.Equal,
				},
				{
					Left:   "52.2.1",
					Right:  "52",
					Expect: corecomparator.Equal,
				},
				{
					Left:   "v2.0",
					Right:  "2.0.0",
					Expect: corecomparator.Equal,
				},
				{
					Left:   "2.0.0.1",
					Right:  "v2.0.0.1",
					Expect: corecomparator.Equal,
				},
			},
			ExpectedInput: []string{
				"0 : [ v2.2 ] = [ v2.2.0 ] | Expect: Equal - true",
				"1 : [ v2.2 ] = [ 2.2.0.0 ] | Expect: Equal - true",
				"2 : [ 2.2 ] = [ v2.2 ] | Expect: Equal - true",
				"3 : [ 52.2.1 ] = [ 52 ] | Expect: Equal - false",
				"4 : [ v2.0 ] = [ 2.0.0 ] | Expect: Equal - true",
				"5 : [ 2.0.0.1 ] = [ v2.0.0.1 ] | Expect: Equal - true",
			},
			VerifyTypeOf: arrangeLeftRightTypeVerification,
			IsEnable:     issetter.True,
		},
		{
			Title: "Versions comparisons - Not Equal",
			ArrangeInput: []args.LeftRight{
				{
					Left:   "v2.2",
					Right:  "2.2.0",
					Expect: corecomparator.NotEqual,
				},
				{
					Left:   "v2.2",
					Right:  "2.2.0.0",
					Expect: corecomparator.NotEqual,
				},
				{
					Left:   "2.2",
					Right:  "v2.2",
					Expect: corecomparator.NotEqual,
				},
				{
					Left:   "2.0.0.1",
					Right:  "2.0.0.1",
					Expect: corecomparator.NotEqual,
				},
				{
					Left:   "v2.0",
					Right:  "2.0.0",
					Expect: corecomparator.NotEqual,
				},
				{
					Left:   "2.2.1",
					Right:  "2",
					Expect: corecomparator.NotEqual,
				},
				{
					Left:   "v52.1",
					Right:  "v52.2",
					Expect: corecomparator.NotEqual,
				},
				{
					Left:   "v0.3.0",
					Right:  "0.2.1",
					Expect: corecomparator.NotEqual,
				},
			},
			ExpectedInput: []string{
				"0 : [ v2.2 ] != [ 2.2.0 ] | Expect: NotEqual - false",
				"1 : [ v2.2 ] != [ 2.2.0.0 ] | Expect: NotEqual - false",
				"2 : [ 2.2 ] != [ v2.2 ] | Expect: NotEqual - false",
				"3 : [ 2.0.0.1 ] != [ 2.0.0.1 ] | Expect: NotEqual - false",
				"4 : [ v2.0 ] != [ 2.0.0 ] | Expect: NotEqual - false",
				"5 : [ 2.2.1 ] != [ 2 ] | Expect: NotEqual - true",
				"6 : [ v52.1 ] != [ v52.2 ] | Expect: NotEqual - true",
				"7 : [ v0.3.0 ] != [ 0.2.1 ] | Expect: NotEqual - true",
			},
			VerifyTypeOf: arrangeLeftRightTypeVerification,
			IsEnable:     issetter.True,
		},
	}

	jsonTestCases = []testWrapper{
		{
			Title: "Create versions json string.",
			ArrangeInput: []string{
				"-1",
				"v1.2.3.4",
				"v5.3.6",
				"5.3",
				"9",
				"v89.1.2   ",
				"v89.1  .2  ",
				someVersionV5.String(),
				"",
			},
			ExpectedInput: []string{
				"0 : [ -1 ] - {\"Compact\":\"-1\",\"Compiled\":\"v-1\",\"IsInvalid\":true}",
				"1 : [ v1.2.3.4 ] - {\"Compact\":\"1.2.3.4\",\"Compiled\":\"v1.2.3.4\",\"Major\":1,\"Minor\":2,\"Patch\":3,\"Build\":4}",
				"2 : [ v5.3.6 ] - {\"Compact\":\"5.3.6\",\"Compiled\":\"v5.3.6\",\"Major\":5,\"Minor\":3,\"Patch\":6}",
				"3 : [ 5.3 ] - {\"Compact\":\"5.3\",\"Compiled\":\"v5.3\",\"Major\":5,\"Minor\":3}",
				"4 : [ 9 ] - {\"Compact\":\"9\",\"Compiled\":\"v9\",\"Major\":9}",
				"5 : [ v89.1.2    ] - {\"Compact\":\"89.1.2\",\"Compiled\":\"v89.1.2\",\"Major\":89,\"Minor\":1,\"Patch\":2}",
				"6 : [ v89.1  .2   ] - {\"Compact\":\"89.1  .2\",\"Compiled\":\"v89.1.2\",\"Major\":89,\"Minor\":1,\"Patch\":2}",
				"7 : [ v5.8.1.5 ] - {\"Compact\":\"5.8.1.5\",\"Compiled\":\"v5.8.1.5\",\"Major\":5,\"Minor\":8,\"Patch\":1,\"Build\":5}",
				"8 : [  ] - {\"IsInvalid\":true,\"Major\":-1,\"Minor\":-1,\"Patch\":-1,\"Build\":-1}",
			},
			VerifyTypeOf: arrangeStringTypeVerification,
			IsEnable:     issetter.True,
		},
	}

	createFunc    = coreversion.New.Default
	someVersionV5 = createFunc("v5.8.1.5")

	versionTwoParamsVerificationTestCases = []testWrapper{
		{
			Title: "IsMajorBuildAtLeast - all matches the condition query.",
			ArrangeInput: []args.Four{
				{
					First:  5,    // major
					Second: 5,    // build
					Third:  true, // expect
					Fourth: someVersionV5.IsMajorBuildAtLeast,
				},
				{
					First:  4,    // major
					Second: 4,    // build
					Third:  true, // expect
					Fourth: someVersionV5.IsMajorBuildAtLeast,
				},
				{
					First:  4,    // major
					Second: 6,    // build
					Third:  true, // expect
					Fourth: someVersionV5.IsMajorBuildAtLeast,
				},
				{
					First:  6,     // major
					Second: 5,     // build
					Third:  false, // expect
					Fourth: someVersionV5.IsMajorBuildAtLeast,
				},
			},
			ExpectedInput: []string{
				"Testing for -> Version(v5.8.1.5)",
				"    0 : .IsMajorBuildAtLeast(5, 5) -> true | true - expected",
				"    1 : .IsMajorBuildAtLeast(4, 4) -> true | true - expected",
				"    2 : .IsMajorBuildAtLeast(4, 6) -> true | true - expected",
				"    3 : .IsMajorBuildAtLeast(6, 5) -> false | false - expected",
			},
			VerifyTypeOf: argsFourTypeVerification,
			IsEnable:     issetter.True,
		},
		{
			Title: "IsMajorMinorAtLeast - all matches the condition query.",
			ArrangeInput: []args.Four{
				{
					First:  5,    // major
					Second: 8,    // minor
					Third:  true, // expect
					Fourth: someVersionV5.IsMajorMinorAtLeast,
				},
				{
					First:  4,    // major
					Second: 1,    // minor
					Third:  true, // expect
					Fourth: someVersionV5.IsMajorMinorAtLeast,
				},
				{
					First:  5,    // major
					Second: 7,    // minor
					Third:  true, // expect
					Fourth: someVersionV5.IsMajorMinorAtLeast,
				},
				{
					First:  5,     // major
					Second: 9,     // minor
					Third:  false, // expect
					Fourth: someVersionV5.IsMajorMinorAtLeast,
				},
				{
					First:  6,     // major
					Second: 1,     // minor
					Third:  false, // expect
					Fourth: someVersionV5.IsMajorMinorAtLeast,
				},
			},
			ExpectedInput: []string{
				"Testing for -> Version(v5.8.1.5)",
				"    0 : .IsMajorMinorAtLeast(5, 8) -> true | true - expected",
				"    1 : .IsMajorMinorAtLeast(4, 1) -> true | true - expected",
				"    2 : .IsMajorMinorAtLeast(5, 7) -> true | true - expected",
				"    3 : .IsMajorMinorAtLeast(5, 9) -> false | false - expected",
				"    4 : .IsMajorMinorAtLeast(6, 1) -> false | false - expected",
			},
			VerifyTypeOf: argsFourTypeVerification,
			IsEnable:     issetter.True,
		},
	}

	versionThreeParamsVerificationTestCases = []testWrapper{
		{
			Title: "IsMajorMinorPatchAtLeast - all matches the condition query.",
			ArrangeInput: []args.Five{
				{
					First:  5,    // major
					Second: 5,    // build
					Third:  1,    // patch
					Fourth: true, // expect
					Fifth:  someVersionV5.IsMajorMinorPatchAtLeast,
				},
				{
					First:  4,    // major
					Second: 4,    // build
					Third:  1,    // patch
					Fourth: true, // expect
					Fifth:  someVersionV5.IsMajorMinorPatchAtLeast,
				},
				{
					First:  5,    // major
					Second: 8,    // build
					Third:  1,    // patch
					Fourth: true, // expect
					Fifth:  someVersionV5.IsMajorMinorPatchAtLeast,
				},
				{
					First:  5,     // major
					Second: 8,     // build
					Third:  2,     // patch
					Fourth: false, // expect
					Fifth:  someVersionV5.IsMajorMinorPatchAtLeast,
				},
				{
					First:  6,     // major
					Second: 5,     // build
					Third:  1,     // patch
					Fourth: false, // expect
					Fifth:  someVersionV5.IsMajorMinorPatchAtLeast,
				},
				{
					First:  7,     // major
					Second: 0,     // build
					Third:  0,     // patch
					Fourth: false, // expect
					Fifth:  someVersionV5.IsMajorMinorPatchAtLeast,
				},
			},
			ExpectedInput: []string{
				"Testing for -> Version(v5.8.1.5)",
				"    0 : .IsMajorMinorPatchAtLeast(5, 5, 1) -> true | true - expected",
				"    1 : .IsMajorMinorPatchAtLeast(4, 4, 1) -> true | true - expected",
				"    2 : .IsMajorMinorPatchAtLeast(5, 8, 1) -> true | true - expected",
				"    3 : .IsMajorMinorPatchAtLeast(5, 8, 2) -> false | false - expected",
				"    4 : .IsMajorMinorPatchAtLeast(6, 5, 1) -> false | false - expected",
				"    5 : .IsMajorMinorPatchAtLeast(7, 0, 0) -> false | false - expected",
			},
			VerifyTypeOf: argsFiveTypeVerification,
			IsEnable:     issetter.True,
		},
	}
)
