package coredynamictestwrappers

import (
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

var (
	ReflectSetFromToInvalidTestCases = []FromToTestWrapper{
		{
			Header: "(null, null) -- do nothing -- no error or exception panic" +
				"From `Null` to `Null` -- does nothing -- no error",
			From:             nil,
			To:               nil,
			IsExpectingError: false,
			HasPanic:         false,
			Validator: corevalidator.TextValidator{
				Search:    "",
				SearchAs:  stringcompareas.Equal,
				Condition: corevalidator.DefaultTrimCoreCondition,
			},
		},
		{
			Header: "(null, valid type - coretests.DraftType) -- should not panic, error returned -- " +
				"From `Null` to `coretests.DraftType`",
			From: nil,
			To: &coretests.DraftType{
				SampleString1: "Same data",
			},
			// ExpectedValue:    &ReflectSetFromToTestCasesDraftTypeExpected,
			IsExpectingError: true,
			HasPanic:         false,
			Validator: corevalidator.TextValidator{
				Search:    "Invalid : value cannot process it. `from` is nil, cannot set null or nil to destination.\"! Supported Types: https://t.ly/SGWUx,  Ref(s) { \"(FromType, ToType) = (<nil>, *coretests.DraftType)\" }",
				SearchAs:  stringcompareas.Equal,
				Condition: corevalidator.DefaultTrimCoreCondition,
			},
		},
		{
			Header: "(valid type - coretests.DraftType, null) -- should not panic, error returned -- " +
				"From `coretests.DraftType` to `Null`",
			From: &coretests.DraftType{
				SampleString1: "Same data",
			},
			To: nil,
			// ExpectedValue:    &ReflectSetFromToTestCasesDraftTypeExpected,
			IsExpectingError: true,
			HasPanic:         false,
			Validator: corevalidator.TextValidator{
				Search:    "Invalid : null pointer, cannot process it. \"destination pointer is null, cannot proceed further!\" Supported Types: https://t.ly/SGWUx,  Ref (s) { \"FromType\", \"*coretests.DraftType\", \"ToType\", \"<nil>\" }",
				SearchAs:  stringcompareas.Equal,
				Condition: corevalidator.DefaultTrimCoreCondition,
			},
		},
		{
			Header: "Pointer type to Value type (valid type - *coretests.DraftType, value type - coretests.DraftType) " +
				"-- should not panic, error returned -- " +
				"From `*coretests.DraftType` to `coretests.DraftType` (value type)",
			From: &coretests.DraftType{
				SampleString1: "Same data",
			},
			To: coretests.DraftType{
				SampleString1: "Same data",
			},
			IsExpectingError: true,
			HasPanic:         false,
			Validator: corevalidator.TextValidator{
				Search:    "Unexpected type error, which is unexpected. \"destination or toPointer must be a pointer to set!\" Supported Types: https://t.ly/SGWUx,  Ref (s) { \"FromType\", \"*coretests.DraftType\", \"ToType\", \"coretests.DraftType\" }",
				SearchAs:  stringcompareas.Equal,
				Condition: corevalidator.DefaultTrimCoreCondition,
			},
		},
		{
			Header: "Value type to value type is not valid - (valid type - coretests.DraftType, value type - coretests.DraftType) " +
				"-- error returned -- " +
				"From `coretests.DraftType` to `coretests.DraftType` (value type)",
			From: coretests.DraftType{
				SampleString1: "Same data",
			},
			To: coretests.DraftType{
				SampleString1: "Same data",
			},
			IsExpectingError: true,
			HasPanic:         false,
			Validator: corevalidator.TextValidator{
				Search:    "Unexpected type error, which is unexpected. \"destination or toPointer must be a pointer to set!\" Supported Types: https://t.ly/SGWUx,  Ref (s) { \"FromType\", \"coretests.DraftType\", \"ToType\", \"coretests.DraftType\" }",
				SearchAs:  stringcompareas.Equal,
				Condition: corevalidator.DefaultTrimCoreCondition,
			},
		},
		{
			Header: "Value type to Pointer is valid - (valid type - coretests.DraftType, pointer type - *coretests.DraftType) " +
				"-- works, no error -- " +
				"From `coretests.DraftType` to `*coretests.DraftType` (value type)",
			From: coretests.DraftType{
				SampleString1: "Same data",
			},
			To: &coretests.DraftType{
				SampleString1: "Same data",
			},
			IsExpectingError: false,
			HasPanic:         false,
			Validator:        corevalidator.EmptyValidator,
		},
	}
)
