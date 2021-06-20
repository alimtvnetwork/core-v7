package chmodhelper

// VerifyChmodMany - expectedHyphenedRwx should be 10 chars example "-rwxrwxrwx"
//
// Multiple files verification error will be returned as once.
func VerifyChmodPathsDirect(
	expectedHyphenedRwx string,
	isContinueOnError bool,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	return VerifyChmodPaths(
		&locations,
		expectedHyphenedRwx,
		isContinueOnError)
}
