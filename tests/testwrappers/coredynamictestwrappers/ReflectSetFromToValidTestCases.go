package coredynamictestwrappers

import (
	"gitlab.com/auk-go/core/coretests"
)

var (
	ReflectSetFromToValidTestCases = []FromToTestWrapper{
		{
			Header: "(null, null) -- do nothing -- " +
				"From `Null` to `Null` -- does nothing -- no error",
		},
		{
			Header: "(sameTypePointer, sameTypePointer) -- try reflection -- " +
				"From `*FromToTestWrapper{Expected}` " +
				"to   `*FromToTestWrapper{Sample data}` should set to Expected. ",
			From: &ReflectSetFromToTestCasesDraftTypeExpected,
			To: &coretests.DraftType{
				SampleString1: "Same data",
			},
			ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
		},
		{
			Header: "(sameTypeNonPointer, sameTypePointer) -- try reflection -- " +
				"From `FromToTestWrapper{Expected}` " +
				"to   `*FromToTestWrapper{Sample data}` should set to Expected.",
			From: ReflectSetFromToTestCasesDraftTypeExpected,
			To: &coretests.DraftType{
				SampleString1: "Sample data",
			},
			ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
		},
		{
			Header: "(*[]byte, otherType) -- try unmarshal, reflect -- " +
				"From `*[]bytes(FromToTestWrapper{Expected}` " +
				"to   `*FromToTestWrapper{Sample data}` should set to Expected.",
			From: ReflectSetFromToTestCasesDraftTypeExpected.JsonBytesPtr(),
			To: &coretests.DraftType{
				SampleString1: "Sample data",
			},
			ExpectedValue: &ReflectSetFromToTestCasesDraftTypeExpected,
		},
		{
			Header: "(otherType, *[]byte) -- try marshal, reflect -- " +
				"From `FromToTestWrapper{Expected}` " +
				"to   `*[]byte{}` should set to Expected.",
			From:          ReflectSetFromToTestCasesDraftTypeExpected.JsonBytesPtr(),
			To:            &[]byte{},
			ExpectedValue: ReflectSetFromToTestCasesDraftTypeExpected.JsonBytesPtr(),
		},
	}
)
