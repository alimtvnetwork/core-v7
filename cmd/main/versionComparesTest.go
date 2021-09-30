package main

func versionComparesTest() {
	testRanges := map[string]string{
		"v0.0.1": "v0.0.1",
		"v0.0.2": "v0.2.1",
		"v3.0":   "v0.2.1",
		"v4":     "v4.0",
	}

	for k, v := range testRanges {
		versionCompareTest(k, v)
	}
}
