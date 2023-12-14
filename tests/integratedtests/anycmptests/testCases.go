package anycmptests

import (
	"reflect"

	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/issetter"
)

var (
	arrangeTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]args.Two{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	testCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "left and right is null checking, " +
					"only Equal if both null or same pointer, " +
					"NotEqual if one is null and another isn't." +
					"On both not null it is inconclusive.",
				ArrangeInput: []args.Two{
					{
						First:  nil,
						Second: nil,
					},
					{
						First:  1,
						Second: nil,
					},
					{
						First:  1,
						Second: 2,
					},
					{
						First:  &coretests.DraftType{},
						Second: nil,
					},
					{
						First:  nil,
						Second: &coretests.DraftType{},
					},
					{
						First:  &coretests.DraftType{},
						Second: &coretests.DraftType{},
					},
					{
						First:  arrangeTypeVerification,
						Second: arrangeTypeVerification,
					},
				},
				ExpectedInput: []string{
					"0 : Equal (<nil>, <nil>)",
					"1 : NotEqual (int, <nil>)",
					"2 : Inconclusive (int, int)",
					"3 : NotEqual (*coretests.DraftType, <nil>)",
					"4 : NotEqual (<nil>, *coretests.DraftType)",
					"5 : Inconclusive (*coretests.DraftType, *coretests.DraftType)",
					"6 : Equal (*coretests.VerifyTypeOf, *coretests.VerifyTypeOf)",
				},
				VerifyTypeOf: arrangeTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
)
