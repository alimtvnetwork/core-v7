package chmodhelpertestwrappers

var VerifyRwxPartialChmodLocationsTestCases = []VerifyRwxPartialChmodLocationsWrapper{
	{
		Header:             "Missing Paths should NOT have error with it's location!",
		Locations:          SimpleLocations,
		IsContinueOnError:  true,
		IsSkipOnInvalid:    true,
		ExpectedPartialRwx: "-rwxrwx",
		ExpectationErrorMessage: "/temp/core/test-cases-2 - " +
			"Expect [\"rwxrwx***\"] != [\"rwxr-xr--\"] Actual\n" +
			"/temp/core/test-cases-3 - " +
			"Expect [\"rwxrwx***\"] != [\"rwxr-xr--\"] Actual",
	},
	{
		Header:                  "Missing Paths should NOT have error with it's location and all matches with WhatIsExpected RWX!",
		Locations:               SimpleLocations,
		IsContinueOnError:       true,
		IsSkipOnInvalid:         true,
		ExpectedPartialRwx:      "-rwx",
		ExpectationErrorMessage: "",
	},
	{
		Header:             "Missing Paths should have error with it's location!",
		Locations:          SimpleLocations,
		IsContinueOnError:  true,
		IsSkipOnInvalid:    false,
		ExpectedPartialRwx: "-rwxrwx-",
		ExpectationErrorMessage: "/temp/core/test-cases-2 - " +
			"Expect [\"rwxrwx-**\"] != [\"rwxr-xr--\"] Actual\n" +
			"/temp/core/test-cases-3 - " +
			"Expect [\"rwxrwx-**\"] != [\"rwxr-xr--\"] Actual\n" +
			"Path missing or having other access issues! Ref(s) { \"" +
			"[/temp/core/test-cases-3s " +
			"/temp/core/test-cases-3x]\" }",
	},
}
