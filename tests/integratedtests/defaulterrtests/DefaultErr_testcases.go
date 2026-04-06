package defaulterrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var defaultErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "DefaultErr returns non-nil error with message -- Marshalling",
		ArrangeInput: args.Map{
			"when":  "checking Marshalling error",
			"error": "Marshalling",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- UnMarshalling",
		ArrangeInput: args.Map{
			"when":  "checking UnMarshalling error",
			"error": "UnMarshalling",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- OutOfRange",
		ArrangeInput: args.Map{
			"when":  "checking OutOfRange error",
			"error": "OutOfRange",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- CannotProcessNilOrEmpty",
		ArrangeInput: args.Map{
			"when":  "checking CannotProcessNilOrEmpty error",
			"error": "CannotProcessNilOrEmpty",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- NegativeDataCannotProcess",
		ArrangeInput: args.Map{
			"when":  "checking NegativeDataCannotProcess error",
			"error": "NegativeDataCannotProcess",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- NilResult",
		ArrangeInput: args.Map{
			"when":  "checking NilResult error",
			"error": "NilResult",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- UnexpectedValue",
		ArrangeInput: args.Map{
			"when":  "checking UnexpectedValue error",
			"error": "UnexpectedValue",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- CannotRemoveFromEmptyCollection",
		ArrangeInput: args.Map{
			"when":  "checking CannotRemoveFromEmptyCollection error",
			"error": "CannotRemoveFromEmptyCollection",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- MarshallingFailedDueToNilOrEmpty",
		ArrangeInput: args.Map{
			"when":  "checking MarshallingFailedDueToNilOrEmpty error",
			"error": "MarshallingFailedDueToNilOrEmpty",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- UnmarshallingFailedDueToNilOrEmpty",
		ArrangeInput: args.Map{
			"when":  "checking UnmarshallingFailedDueToNilOrEmpty error",
			"error": "UnmarshallingFailedDueToNilOrEmpty",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- KeyNotExistInMap",
		ArrangeInput: args.Map{
			"when":  "checking KeyNotExistInMap error",
			"error": "KeyNotExistInMap",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
}
