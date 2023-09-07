package coredynamictestwrappers

import (
	"gitlab.com/auk-go/core/coretests"
)

var (
	ReflectSetFromToTestCasesDraftTypeInput = coretests.DraftType{
		SampleString1: "Same data",
		SampleString2: "",
		SampleInteger: 0,
	}
	ReflectSetFromToTestCasesDraftTypeExpected = coretests.DraftType{
		SampleString1: "Expected",
		SampleString2: "",
		SampleInteger: 0,
	}

	nilBytes        []byte  = nil
	nilBytesPointer *[]byte = nil

	ReflectSetFromToValidTestCases = []ReflectSetFromToTestWrapper{
		{
			Header: "(null, null) -- do nothing -- " +
				"From `Null` to `Null` -- does nothing -- no error",
		},
		{
			Header: "(sameTypePointer, sameTypePointer) -- try reflection -- " +
				"From `*ReflectSetFromToTestWrapper{Expected}` " +
				"to   `*ReflectSetFromToTestWrapper{Sample data}` should set to Expected. ",
			From: &ReflectSetFromToTestCasesDraftTypeExpected,
			To: &coretests.DraftType{
				SampleString1: "Same data",
			},
			ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
		},
		{
			Header: "(sameTypeNonPointer, sameTypePointer) -- try reflection -- " +
				"From `ReflectSetFromToTestWrapper{Expected}` " +
				"to   `*ReflectSetFromToTestWrapper{Sample data}` should set to Expected.",
			From: ReflectSetFromToTestCasesDraftTypeExpected,
			To: &coretests.DraftType{
				SampleString1: "Sample data",
			},
			ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
		},
		{
			Header: "(*[]byte, otherType) -- try unmarshal, reflect -- " +
				"From `*[]bytes(ReflectSetFromToTestWrapper{Expected}` " +
				"to   `*ReflectSetFromToTestWrapper{Sample data}` should set to Expected.",
			From: ReflectSetFromToTestCasesDraftTypeExpected.JsonBytesPtr(),
			To: &coretests.DraftType{
				SampleString1: "Sample data",
			},
			ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
		},
		{
			Header: "(otherType, *[]byte) -- try marshal, reflect -- " +
				"From `ReflectSetFromToTestWrapper{Expected}` " +
				"to   `*[]byte{}` should set to Expected.",
			From:          ReflectSetFromToTestCasesDraftTypeExpected.JsonBytesPtr(),
			To:            &[]byte{},
			ExpectedValue: ReflectSetFromToTestCasesDraftTypeExpected.JsonBytesPtr(),
		},
	}
)
